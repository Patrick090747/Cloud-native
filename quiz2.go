package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		data := i
		fmt.Println("生产者生产数据:", data)
		time.Tick(time.Second)
		out <- data // 缓冲区写入数据
	}
	close(out) //写完关闭管道
}

func consumer(in <-chan int) {
	for data := range in {
		time.Tick(time.Second)
		fmt.Println("消费者得到数据：", data)
	}

}
func main() {
	ch := make(chan int) //无缓冲channel
	go producer(ch)      // 子go程作为生产者
	consumer(ch)         // 主go程作为消费者
}
