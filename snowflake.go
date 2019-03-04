package snowflake

import (
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"
)

var (
	// seq 序號
	seq uint32

	// rwMux 變數鎖
	mux sync.Mutex
)

// Get 取得
func get() *big.Int {
	/*
		Unix timestamp_PID_seq		1545112028_12345_12345
									15451120281234512345
		Maximum uint64				18446744073709551615
	*/
	mux.Lock()
	s := fmt.Sprintf("%d%d%d", time.Now().Unix(), os.Getpid(), seq)
	seq++
	mux.Unlock()

	i := new(big.Int)
	fmt.Sscan(s, i)
	return i
}

// Uint64 取得數值
func Uint64() uint64 {
	return get().Uint64()
}

// Text 輸出 62 進制文字
func Text() string {
	return get().Text(62)
}
