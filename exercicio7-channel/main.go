package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)
	channelBuffer := make(chan int, 100)
	go setList(channelBuffer)
	go setList(channel)
	for i := range channelBuffer {
		fmt.Println("Recebendo Bufer: INDEX -> ", i)

	}
	for i := range channel {
		fmt.Println("Recebendo: INDEX -> ", i)
		time.Sleep(time.Second)
	}
}

func setList(channel chan<- int) {
	for i := 0; i < 100; i++ {
		channel <- i
		fmt.Println("Enviado: INDEX -> ", i)
	}
	close(channel)

}
