package main
import "fmt"

type Vector3 struct {
	x, y, z float32
}

// Vector3 Method
func (v Vector3) Print() {
	fmt.Printf("[ Vector3 | x[%v] y[%v] z[%v] ]\n", v.x, v.y, v.z)
}

func (self Vector3) Add(other Vector3) Vector3 {
	return Vector3{ self.x + other.x, self.y + other.y, self.z + other.z }
}

func main() {
	v1 := Vector3{1,2,3}
	v2 := Vector3{5,6,7}
	v1.Print()
	v2.Print()
	v1.Add(v2).Print()
}