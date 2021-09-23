package main

import "fmt"

func main() {
	/*
		interface is a type same as struct.
		struct holds data whereas interface does not hold data
		instead store method definition
	*/
	/*
		implicit is often used to refer to something that's
		done for you by other code behind the scenes.
		Explicit is the manual approach to accomplishing
		the change you wish to have by writing out the
		instructions to be done explicitly
	*/

	var t tank
	t = myvalue{10, 14}
	fmt.Println("Area of tank", t.Tarea())
	fmt.Println("Volume of tank", t.Volume())

}

type tank interface {

	// Methods
	Tarea() float64
	Volume() float64
}

type myvalue struct {
	radius float64
	height float64
}

// Implementing methods of
// the tank interface

// Creating an interface
/* type myinterface interface{

// Methods
fun1() int
fun2() float64
} */
func (m myvalue) Tarea() float64 {
	fmt.Println(m)
	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}

//here m is a variable set for myvalue struct which is {10,14}

func (m myvalue) Volume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}
