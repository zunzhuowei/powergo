package tests

import "fmt"

func TestChannel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c // 从通道 c 中接收

	x := <-c
	y := <-c

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	fmt.Println(s)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

// 用了两个 goroutine 线程，所以x,y的值，根据线程的执行顺序，可能不一样
//17 -5 12
//-5 17 12

func TestChannel2() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}

}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
