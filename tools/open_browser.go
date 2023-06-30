package tools

import (
	"fmt"
	"github.com/pkg/browser"
	"log"
)

// OpenBrowser 打开本地浏览器
func OpenBrowser(bind string, port int) {
	if bind == "0.0.0.0" {
		bind = "localhost"
	}

	url := fmt.Sprintf("http://%s:%d", bind, port)
	err := browser.OpenURL(url)
	if err != nil {
		log.Printf("请手动打开浏览器并访问 %s 进行操作\n", url)
	} else {
		log.Println("打开浏览器成功, 请在网页中进行操作")
	}
}
