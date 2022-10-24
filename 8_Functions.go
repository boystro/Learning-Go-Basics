package main

import "fmt"

func main() {
	myMessage()
	printName("Rose")
	fmt.Printf("[%v] Returned from add(1,2)", add(1,2))
	prawn(1,10)
}

func myMessage() {
	fmt.Println("myMessage function executed successfully!!!");
}


func printName(name string) {
	fmt.Println("Hello", name)
}

// named return value
// for without name, use
// func add(num1 int, num2 int) int { statements... }
func add(num1 int, num2 int) (result int) {
	result = num1 + num2
	return result
}

func prawn(lo int, hi int) int {
	if hi <= lo {
		fmt.Println(hi)
		return lo
	}
	fmt.Println(hi)
	return prawn(lo, hi - 1)
}