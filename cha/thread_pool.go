package cha

import (
	"fmt"
	"strconv"
	"time"
)

// 1. 常驻线程
// 2. 线程可扩展到一定范围

func GetThreadPool() {
	ch := make(chan int, 5)
	x := 0
	for {
		x++
		ch <- x
		go func(i int) {
			fmt.Println("开启一个routine" + strconv.Itoa(i))
			time.Sleep(time.Second)
			<-ch
		}(x)
	}
}

func run() {
	fmt.Println("running")
}
