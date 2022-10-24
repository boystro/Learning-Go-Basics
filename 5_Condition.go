package main;
import "fmt";
func main() {

	// Leap Year Problem
	year := 2002

	if year%100 == 0 {
		if year%400 == 0 {
			fmt.Println("Leap Year");
		} else {
			fmt.Println("Not Leap Year");
		}
	} else {
		if year%4 == 0 {
			fmt.Println("Leap Year");
		} else {
			fmt.Println("Not Leap Year");
		}
	}

}