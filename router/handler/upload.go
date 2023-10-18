package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	. "thinkprinter/models"
	"thinkprinter/printer"
)

// Upload 批量上传文件
func Upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, ERR.
			WithCode(http.StatusBadRequest).
			WithMsg("上传文件失败"))
		return
	}
	files, ok := form.File["files"]
	if !ok || len(files) == 0 {
		return
	}
	username := c.GetString("username")
	dir := filepath.Join(os.TempDir(), "ThinkPrint", username)
	for _, file := range files {
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

	count := printer.WG.GetCount()
	// 推送到打印队列
	for _, file := range files {
		printer.PrintQueue <- filepath.Join(dir, file.Filename)
		printer.WG.Add(1)
	}
	c.JSON(http.StatusOK, OK.
		WithMsg("成功推送到打印队列").
		WithData("前方排队"+strconv.Itoa(count)+"个任务"))
}
