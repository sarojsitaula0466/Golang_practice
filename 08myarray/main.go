package main

import "fmt"

func main() {
	fmt.Println("Array in golang")

	var fruitList [4]string
	// In golang we need to mention how much data is coming, its compulsion
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	//fruitList[2]="Apple"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is :", fruitList)
	fmt.Println("Length :", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"} // another way to declare array
	fmt.Println("Veg list is", vegList)
}
