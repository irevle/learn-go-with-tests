package main

import "fmt"

// string keyword indicates the func returns string
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Elv"))
}
