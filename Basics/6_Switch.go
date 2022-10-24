package main
import "fmt"

func main() {
	op   := "plus"

	switch op {
	case "+", "add", "plus":
		fmt.Println(5+5)
	case "-":
		fmt.Println(5-5)
	case "*":
		fmt.Println(5*5)
	case "/":
		fmt.Println(5/5)
	default:
		fmt.Println("[Error]")
	}
}