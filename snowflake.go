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
func get() (*big.Int, error) {
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
	_, e := fmt.Sscan(s, i)
	if e != nil {
		return nil, e
	}
	return i, nil
}

// Get 取得數值
func Get() (uint64, error) {
	i, e := get()
	if e != nil {
		return 0, e
	}
	return i.Uint64(), nil
}

// Text 輸出 62 進制文字
func Text() (string, error) {
	i, e := get()
	if e != nil {
		return "", e
	}
	return i.Text(62), nil
}
