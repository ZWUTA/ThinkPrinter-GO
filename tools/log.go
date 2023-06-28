package tools

import (
	"log"
	"net/http"
)

// 重写io.Writer的Write函数函数，本质上就是遍历数组，比较巧妙
//
//	func (t *multiWriter) Write(p []byte) (n int, err error) {
//		for _, w := range t.writers {
//			n, err = w.Write(p)
//			if err != nil {
//				return
//			}
//			if n != len(p) {
//				err = ErrShortWrite
//				return
//			}
//		}
//		return len(p), nil
//	}
//
// https://blog.csdn.net/xmcy001122/article/details/119916227
func OutputLog(r *http.Request) {
	log.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
}
