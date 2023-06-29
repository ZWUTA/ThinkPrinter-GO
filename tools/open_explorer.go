package tools

import (
	"github.com/pkg/browser"
	"log"
)

// OpenExplorer 打开本地浏览器
func OpenExplorer(url string) {
	url = "http://" + url
	err := browser.OpenURL(url)
	if err != nil {
		log.Printf("请手动打开浏览器并访问 %s 进行操作\n", url)
	} else {
		log.Println("打开浏览器成功, 请在网页中进行操作")
	}
}
