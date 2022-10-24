package main
import "fmt"
func main() {
	myMap := make(map[string]int)

	myMap["banana"] = 20
	myMap["apple"]  = 10

	fmt.Printf("a\t%v\n", myMap)
	fmt.Printf("%T\t%#v\n", myMap["banana"], myMap["banana"])
}