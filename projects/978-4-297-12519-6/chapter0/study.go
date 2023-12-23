package main

func main() {
	var empty struct{}

	var point struct {
		ID   string
		x, y int
	}

	var array [5]int

	arrayLiteral := [5]int{1, 2, 3, 4, 5}

	arrayInterface := [...]int{1, 2, 3, 4, 5}

	arrayIndex := [...]int{2: 1, 5: 5, 7: 13}
}
