package main;
import "fmt";
func main() {
	num := 10;
	str := "sTriNg"
	// Prints the arguments
	fmt.Print("Hi", "World", num, "\n");

	// Prints the arguments and adds a new line
	fmt.Println("WAHOOOO", num);

	// Print Format - Similar to C's printf()
	// %v - Value
	// %T - type
	fmt.Printf("num : %T = %v\n", num, num);
	fmt.Printf("str : %T = %v\n", str, str);

	// Format Verbs
	fmt.Printf("%%v  = %v\n", str);
	fmt.Printf("%%#v = %#v\n", str);
	fmt.Printf("%%T  = %T\n", str);
	fmt.Printf("%%%% = %%\n");
}

/*
Integer Formatting Verbs
%b   - base 2
%o   - base 8
%O   - base 8  /starting with 0o
%x   - base 16 / lowercase
%X   - base 16 / uppercase
%#x  - base 16 / starting with 0x
%d   - base 10
%+d  - base 10 / always sign
%4d  - Pad with spaces / Width:4 /Right justify
%-4d - Pad with spaces / Width:4 /Left justify
%04d - Pad with zeroes / Width:4 /
*/

/*
String Formatting Verbs
Boolean Formatting Verbs
Float Formatting Verbs
*/