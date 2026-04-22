package main

import "fmt"

// string keyword indicates the func returns string
func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}
