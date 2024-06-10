package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a := slice[:5]
	b := slice[5:]
	t1, t2 := 0, 0
	for i := 0; i < len(a); i++ {
		t1 += a[i]
	}
	for i := 0; i < len(b); i++ {
		t2 += b[i]
	}

	fmt.Println(t1, t2)
}
