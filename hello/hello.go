package main

import "fmt"

// Hello returns a greeting.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
