package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//we use defer to wait function to do all the things first and execute defer line before function return something
/*
1. function executes
2. defer
3. panic
4. return function

*/
func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll((res.Body))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
