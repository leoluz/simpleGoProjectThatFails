package main

import (
	"fmt"

	"example.com/main/cmd/simple/prod"
	"example.com/sum"
)

func main() {
	num1 := 1
	num2 := 2
	sum := sum.Sum(num1, num2)
	prod := prod.Product(num1, num2)

	fmt.Println(sum)
	fmt.Println(prod)
}
