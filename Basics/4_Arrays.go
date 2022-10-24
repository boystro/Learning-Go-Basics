package main;
import "fmt";

func main() {
	var arr1 = [3]int{1,2,3};
	arr2 := [...]int{4,5,6,7,8,9}; // Full initialisation
	arr3 := [5]int{1,2}; // Partial initialisation
	arr4 := [7]int{}; // No initialisation
	arr5 := [5]int{1:5,2:6}; // Specific Initialisation

	fmt.Println(arr1);
	fmt.Println(arr2);
	fmt.Println(arr3);
	fmt.Println(arr4);
	fmt.Println(arr5);

	fmt.Println("len[arr5] = ", len(arr5));
}