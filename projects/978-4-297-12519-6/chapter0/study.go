package main

import (
	"fmt"
)

func main() {
	src := []int{1, 2, 3, 4}
	fmt.Println(src, len(src), cap(src))

	src = append(src, 5)
	fmt.Println(src, len(src), cap(src))
}
