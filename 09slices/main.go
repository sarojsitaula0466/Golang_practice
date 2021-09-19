package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Apple", "Tomato", "peach"}
	fmt.Printf("Type of fruitlist is %T, %v\n", fruitList, fruitList) // []string, which means slices
	// to insert values in slices
	fruitList = append(fruitList, "mango", "banana") // it appends from last item
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:]) // it slice the slices from index 1 or neglects 0 index
	fmt.Println(fruitList)

	//fruitList = append(fruitList[1:3]) // it slice index 1 and 2 only
	//fmt.Println(fruitList)
	fruitList = append(fruitList[:3]) // it starts with default and slice 0,1 and 2 index
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// you can not add any value to the slices unless you use append method because slices give extra memory when append is used
	highScores = append(highScores, 555, 666, 321)
	sort.Ints(highScores)
	fmt.Println(highScores)

	//remove a value from slices based on index

	var courses = []string{"reactJs", "JavaScript", "swift", "python", "ruby"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
