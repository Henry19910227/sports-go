package tracker

import (
	"fmt"
	"runtime"
)

type tool struct {
}

func New() Tool {
	return &tool{}
}

func (t *tool) GoroutineID() int {
	// 獲取堆疊信息
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false)
	stack := string(buf[:n])

	// 提取 goroutine ID
	var goroutineID int
	fmt.Sscanf(stack, "goroutine %d ", &goroutineID)
	return goroutineID
}
