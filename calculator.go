package main

import "fmt"

func add(a, b int) int {
	//Insert code here
	result := a + b
	return result
}

func subtract(a, b int) int {
	//Insert code here
	result := a - b
	return result
}

func multiply(a, b int) int {
	//Insert code here
	result := a * b
	return result
}

func divide(a, b int) float32 {
	//Insert code here
	//consider for b = 0
	var x float32 = float32(a)
	var y float32 = float32(b)
	result := x / y
	return result
}

func main() {
	var a, b int
	var process string
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&a)
	fmt.Println("Enter process: (add, sub, mul, div)")
	fmt.Scanln(&process)
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&b)

	//Insert code here
	if a == 0 || b == 0 {
		fmt.Println("Cannot divide by 0")
	}else if process == "add" {
		ans := add(a, b)
		fmt.Println("The answer is :", ans)
	}else if process == "sub" {
		ans := subtract(a, b)
		fmt.Println("The answer is :", ans)
	}else if process == "mul" {
		ans := multiply(a, b)
		fmt.Println("The answer is :", ans)
	}else if process == "div" {
		ans := divide(a, b)
		fmt.Println("The answer is :", ans)
	}

	main()
}
