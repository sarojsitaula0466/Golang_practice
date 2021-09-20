package main

import "fmt"

func main() {
	statePopulations := map[string]int{
		"california":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	fmt.Println(statePopulations)
	statePopulations["Georgia"] = 10310371
	// return order of map may vary, keep that in mind
	fmt.Println(statePopulations)
	// to delete we use delete in built method
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	// in order to check if you have misspelled key you use below commands

	/* _, ok := statePopulations["Oho"]
	fmt.Println(ok) */

	fmt.Println(len(statePopulations))

}
