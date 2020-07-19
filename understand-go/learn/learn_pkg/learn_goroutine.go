package learn_pkg

import (
	"fmt"
	"time"
)

func Print_Goroutine() int {
	//go say("World!")
	say("Hello!")
	s := []int{1, 2, 3, 4, 5, 6, 7}
	// 带有缓冲区的channel 意味着值可以暂时保存在缓冲区中 如果缓冲区已经满，则go协程将阻塞;
	c := make(chan int, 3)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y) // 注意x,y的值不固定;
	return 0
}

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(10*time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s [] int, c chan int){
	sum := 0
	for _, v := range s	{
		sum += v
	}
	c <- sum // 把sum 发送到通道c中;
}