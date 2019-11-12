// 优雅停止多个goroutine工作者
// 在Go1.7发布时，标准库增加了一个context包，
// 用来简化对于处理单个请求的多个Goroutine之间与请求域的数据、超时和退出等操作
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go Worker(ctx, &wg)
	}
	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}
