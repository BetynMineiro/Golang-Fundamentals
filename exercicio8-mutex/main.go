package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	i := 0
	for x := 0; x < 100000; x++ {
		go func() {
			m.Lock()
			i++
			m.Unlock()
		}()
	}
	time.Sleep(time.Second * 5)
	fmt.Println("Com Lock Mutex -> ", i)
	i = 0
	for x := 0; x < 100000; x++ {
		go func() {
			i++
		}()
	}

	time.Sleep(time.Second * 5)
	fmt.Println("Sem Lock Mutex numero aleatorio  -> ", i)
}
