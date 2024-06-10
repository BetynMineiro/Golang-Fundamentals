package main

import "fmt"

func main() {
	c := [5]int{1, 2, 3, 4, 5}
	t := 0
	for i := 0; i < len(c); i++ {
		t += c[i]
	}
	fmt.Println(t)
}
