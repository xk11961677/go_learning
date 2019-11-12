// select 类似linux io 中 select 获取channel
// 当多个channel有数据时,随机获取一个
// go 语言可以给定 default chan ,如果没有default 则会阻塞
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
