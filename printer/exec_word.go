package printer

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/exec"
	"thinkprinter/models"
	"time"
)

func Print(docPath string) error {
	_, err := os.Stat(docPath)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			slog.Error("打印超时，任务已被强行停止", "file", docPath)
		}
	}()

	cmd := exec.CommandContext(ctx, models.C.Print.WordExePath, docPath,
		"/mFilePrintDefault",
		"/mFileCloseOrExit",
		"/x",
		"/q",
		"/n",
		"/safe",
		"/a",
	)

	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
