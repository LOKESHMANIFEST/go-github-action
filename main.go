package main

import "fmt"

func main() {
	fmt.Println("hello")
	helloResult := hello()
	fmt.Printf("%s", helloResult)
}

func hello() string {
	return "hello 1 1 1 "
}
