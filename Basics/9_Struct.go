package main

import "fmt"

type Vector3 struct {
	x int
	y int
	z int
}

func main() {
	var v1 Vector3;

	v1.x = 0
	v1.y = 5
	v1.z = -2

	fmt.Println(v1.x)
	fmt.Println(v1.y)
	fmt.Println(v1.z)
}