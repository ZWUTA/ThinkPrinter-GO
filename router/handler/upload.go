package handler

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	. "thinkprinter/models"
	"thinkprinter/printer"

	"github.com/gin-gonic/gin"
)

// Upload 批量上传文件
func Upload(c *gin.Context) {
	var allowedExt = []string{
		".docx",
		".doc",
		".pdf",
		".odt",
		".rtf",
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, ERR.
			WithCode(http.StatusBadRequest).
			WithMsg("表单无效"))
		return
	}
	files, ok := form.File["files"]
	if !ok || len(files) == 0 {
		c.JSON(http.StatusBadRequest, ERR.
			WithCode(http.StatusBadRequest).
			WithMsg("无文件上传"))
		return
	}
	username := c.GetString("username")
	dir := filepath.Join(os.TempDir(), "ThinkPrinter", username)
	for _, file := range files {
		// 检查文件后缀
		for _, ext := range allowedExt {
			if filepath.Ext(file.Filename) == ext {
				break
			}
			c.JSON(http.StatusBadRequest, ERR.
				WithCode(http.StatusBadRequest).
				WithMsg("非法文件格式，仅支持").
				WithData(allowedExt))
			return
		}

		slog.Info("上传文件", "username", username,
			"filename", file.Filename,
			"size", file.Size,
			"dir", dir)

		err := c.SaveUploadedFile(file, filepath.Join(dir, file.Filename))
		if err != nil {
			slog.Error("保存文件失败", "error", err)
			c.JSON(http.StatusInternalServerError, ERR.
				WithCode(http.StatusInternalServerError).
				WithMsg("保存文件失败"))
			return
		}
	}

	waiting := printer.WG.GetCount()
	count := 0
	// 推送到打印队列
	for _, file := range files {
		printer.PrintQueue <- filepath.Join(dir, file.Filename)
		printer.WG.Add(1)
		count++
	}
	c.JSON(http.StatusOK, OK.
		WithMsg("成功推送到打印队列").
		WithData(gin.H{
			"waiting": waiting,
			"count":   count,
		}))
}
