package snowflake

import (
	"fmt"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

func init() {
	go resetSequenceLoop()
}

// seq 序號
var seq uint32

// resetSequenceLoop 重置序號
func resetSequenceLoop() {
	for range time.Tick(time.Second) {
		atomic.StoreUint32(&seq, 0)
	}
}

// Get 取得
func Get() uint64 {
	s := fmt.Sprintf("%d%d%d", time.Now().Unix(), os.Getpid(), atomic.LoadUint32(&seq))
	atomic.StoreUint32(&seq, atomic.AddUint32(&seq, 1))
	n, _ := strconv.ParseUint(s, 10, 64)
	return n
}
