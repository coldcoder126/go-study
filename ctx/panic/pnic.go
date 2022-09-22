package panic

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func father() {

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go son(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	ctx.Done()
	wg.Wait()
	fmt.Println("end")
}

func son(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	wg.Done()
}
