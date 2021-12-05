package main

func main() {
	ch := make(chan int)
	<-ch // 阻塞main goroutine, 信道c被锁
}
