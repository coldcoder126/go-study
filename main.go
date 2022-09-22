package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

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
func main() {
	fmt.Println("begin")
	//GetThreadPool()
	rand.Seed(1)
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(5))
	}

}
