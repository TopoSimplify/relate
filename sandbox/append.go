package main

import "fmt"

func reverse(coordinates []int) []int {
	for i, j := 0, len(coordinates)-1; i < j; i, j = i+1, j-1 {
		coordinates[i], coordinates[j] = coordinates[j], coordinates[i]
	}
	return coordinates
}

func main() {
	var ls = []int{1, 2, 3, 4, 5}
	reverse(ls)
	fmt.Println(ls)
}
