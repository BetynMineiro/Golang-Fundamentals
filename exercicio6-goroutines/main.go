package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go callProcessA(&wg)
	go callProcessB(&wg)
	go callProcessC(&wg)

	wg.Wait()
}

func callProcessA(wg *sync.WaitGroup) {
	fmt.Println("Inicializando Processo A")
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizando Processo A")
	wg.Done()
}

func callProcessB(wg *sync.WaitGroup) {
	fmt.Println("Inicializando Processo B")
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizando Processo B")
	wg.Done()
}

func callProcessC(wg *sync.WaitGroup) {
	fmt.Println("Inicializando Processo C")
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizando Processo C")
	wg.Done()
}
