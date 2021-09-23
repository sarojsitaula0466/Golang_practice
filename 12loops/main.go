package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//looping through collection
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	statePopulations := map[string]int{
		"california":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	//only to print values in Golang

	for _, v := range statePopulations {
		fmt.Println(v)
	}
}
