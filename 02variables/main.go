package main

import "fmt"

const LoginToken string = "asdfghjk" // public, can be used by any

func main() {
	var username string = "Saroj"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggidIn bool = true
	fmt.Println(isLoggidIn)
	fmt.Printf("Variable is of type: %T \n", isLoggidIn)

	var value int = 30
	fmt.Println(value)
	fmt.Printf("Variable is of type: %T \n", value)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
