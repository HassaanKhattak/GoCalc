package main

import "fmt"

func add(a float64, b float64) float64 {
	return a + b
}

func sub(a float64, b float64) float64 {
	return a - b
}

func prod(a float64, b float64) float64 {
	return a * b
}

func div(a float64, b float64) float64 {
	if b == 0 {
		fmt.Println("Error! cannot divide by zero.")
		return 0
	} else {
		return a / b
	}
}

func main() {
	fmt.Println(add(2.45, 13.2))
	fmt.Println(sub(1.29, 12.8))
	fmt.Println(prod(23.4, 2))
	fmt.Println(div(1.34, 23.4))

}
