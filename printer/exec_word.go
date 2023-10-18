package printer

import (
	"os"
	"os/exec"
	"thinkprinter/models"
)

func Print(docPath string) error {
	_, err := os.Stat(docPath)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		}
	}

	cmd := exec.Command(models.C.Print.WordExePath, docPath,
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
