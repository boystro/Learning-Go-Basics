package main;
import "fmt";

// Global Variable
var gVar = "This is Global";

func main() {
	// Variables can be declared in the following ways
	var num1 int;
	num1 = 2000;
	var num2 int = 2000;
	var num3 = 2000;
	num4 := 2000; // Variable value MUST be initialized // Can ONLY be used INSIDE Functions

	// Other datatypes
	var s string = "Indranil";
	var f float32 = 2.2; // 32 - 64
	var n int64 = -300000000000; // int - 8 - 16 - 32 - 64
	var u uint64 = 300000000000; // 8 - 16 - 32 - 64
	var b bool = false;

	fmt.Println(gVar);
	fmt.Println(num1);
	fmt.Println(num2);
	fmt.Println(num3);
	fmt.Println(num4);
	fmt.Println(s);
	fmt.Println(f);
	fmt.Println(n);
	fmt.Println(u);	
	fmt.Println(b);

	// Multiple Variable Declaration
	var ma, mb = 6, "Hello";
	mc, md := 5, "Multi";

	fmt.Println(ma);
	fmt.Println(mb);
	fmt.Println(mc);
	fmt.Println(md);

	// Block Declaration
	var (
		ba int = 1;
		bb = "Hello";
	);

	fmt.Println(ba);
	fmt.Println(bb);

	// Constant Values
	const PI = 3.14;
	fmt.Println(PI);
}