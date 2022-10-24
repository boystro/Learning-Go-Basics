package main
import "fmt"

// GO only has FOR loop

func main() {
	size := 5

	for i := 1; i <= size; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ");
		}
		fmt.Print("\n");
	}

	// Array iteration
	myArr := [...]string{"banana","orange","apple","mango"}

	for i := 0; i < len(myArr); i++ {
		fmt.Print(myArr[i], " ");
	}
	fmt.Print("\n");

	// We can change 'idx' to '_' if we do not wish to use it
	for idx, val := range myArr {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	// Personal Note : for can be used just like while
	k := 10
	for k > 0 {
		fmt.Print(k, " ")
		k--
	}
}