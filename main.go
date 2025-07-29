package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sum(a float64, b float64) float64 {
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
		fmt.Println("\033[1;31mError! cannot divide by zero.")
		return 0
	} else {
		return a / b
	}
}

func correctchoice(input string, correctoptions []string) bool {
	for _, option := range correctoptions {
		if input == option {
			return true
		}
	}

	return false

}

func getnumbers() (float64, float64) {

	var num1, num2 float64
	for {
		fmt.Print("Enter first number: ")
		var a string
		fmt.Scanln(&a)

		num, err := strconv.ParseFloat(a, 64)
		if err != nil {
			fmt.Println("\033[1;31mInvalid Input, try again  \033[0m")
			continue
		}
		num1 = num

		break
	}
	for {
		fmt.Print("Enter second number: ")
		var b string
		fmt.Scanln(&b)

		num, err := strconv.ParseFloat(b, 64)
		if err != nil {
			fmt.Println("\033[1;31mInvalid Input, try again. \033[0m")
			continue
		}
		num2 = num
		break
	}

	return num1, num2

}

func avg() float64 {
	var numbers []float64
	for {
		fmt.Println("Enter the numbers. Press q to stop.: ")
		var input string
		fmt.Scanln(&input)

		if input == "q" {
			break
		}
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("\033[1;31mInvalid Input, try again.  \033[0m")
			continue
		}
		numbers = append(numbers, num)

	}
	var sum float64
	for _, value := range numbers {
		sum += value
	}
	if len(numbers) == 0 {
		fmt.Println("\033[1;31mincorrect input. \033[0m")
		return 0
	}
	return sum / float64(len(numbers))

}

func main() {
	correctoptions := []string{"1", "2", "3", "4", "5"}

	var a float64
	var b float64
	var op string

	fmt.Println("Welcome to Calculator! please pick the desired operation.")

	for {

		fmt.Println("\033[0m1 ======> Sum ")
		fmt.Println("\033[0m2 ======> Subtraction ")
		fmt.Println("\033[0m3 ======> Multiplication ")
		fmt.Println("\033[0m4 ======> Division ")
		fmt.Println("\033[0m5 ======> Average ")

		fmt.Scanln(&op)
		op = strings.ToLower(op)

		if op == "q" {
			fmt.Println("\033[1;31mQuitting. ")
			break
		}
		if !correctchoice(op, correctoptions) {
			fmt.Println("\033[1;31mInvalid options, please enter 1-5 or 'q' to quit. ")
			continue

		}

		// add a check; if op=5-> start a loop to take input; if input=q then exit loop
		if op != "5" {
			a, b = getnumbers()
		}

		switch op {
		case "1":
			result := sum(a, b)
			fmt.Println("\033[1;32m the sum is: ", result)

		case "2":
			result := sub(a, b)
			fmt.Println("\033[1;32m the sub is: ", result)

		case "3":
			result := prod(a, b)
			fmt.Println("\033[1;32m the prod is: ", result)

		case "4":
			result := div(a, b)
			fmt.Println("\033[1;32m the div is: ", result)

		case "5":
			result := avg()
			fmt.Println("\033[1;32m the avg is: ", result)

		default:
			fmt.Println("Please enter correct symbols or numbers.")
		}
		fmt.Println("\033[0mContinue? yes/no")
		var choice string

		fmt.Scanln(&choice)
		if choice == "no" {
			break

		}
	}
}
