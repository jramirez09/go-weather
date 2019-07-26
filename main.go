package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, World!")
	}
}

// SayHello prints out someone's name.
func SayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}
