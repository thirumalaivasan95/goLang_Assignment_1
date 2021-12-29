package main

import (
	"fmt"
)

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func input(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return input(x, err)
}

func main() {
	fmt.Print("num1: ")
	num1 := input([]int{}, nil)

	fmt.Print("num2: ")
	num2 := input([]int{}, nil)

	num2 = reverseInts(num2)

	for i, _ := range num1 {
		num1[i] = num1[i] + num2[i]
	}

	fmt.Println(num1)
}
