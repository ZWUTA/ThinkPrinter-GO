package printer

import (
	"log/slog"
	"sync"
	"sync/atomic"
)

type WaitGroupCount struct {
	sync.WaitGroup
	count int64
}

func (wg *WaitGroupCount) Add(delta int) {
	atomic.AddInt64(&wg.count, int64(delta))
	wg.WaitGroup.Add(delta)
}

func (wg *WaitGroupCount) Done() {
	atomic.AddInt64(&wg.count, -1)
	wg.WaitGroup.Done()
}

func (wg *WaitGroupCount) GetCount() int {
	return int(atomic.LoadInt64(&wg.count))
}

var PrintQueue = make(chan string)
var WG WaitGroupCount

// PrintWorker 打印工作线程
func PrintWorker() {
	for {
		docPath := <-PrintQueue

		slog.Info("开始打印", "file", docPath)
		err := Print(docPath)
		if err != nil {
			slog.Error("打印失败", "file", docPath, "error", err)
			WG.Done()
			return
		}
		slog.Info("打印完成", "file", docPath)
		WG.Done()
	}
}
