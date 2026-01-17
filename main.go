package main

import (
	"fmt"

	"github.com/nukopy/tdd-with-go/greeting"
	"github.com/nukopy/tdd-with-go/integers"
)

func main() {
	// greeting
	name := "Yosuke"
	language := "French"
	fmt.Println(greeting.Hello(name, language))

	// integers
	x := 3
	y := -4
	fmt.Printf("%d + %d = %d\n", x, y, integers.Add(x, y))

	// HTTP server
	// server := NewServer(8888)
	// fmt.Printf("Starting HTTP server on port %d...\n", 8888)
	// if err := server.ListenAndServe(); err != nil {
	// 	fmt.Printf("Error starting server: %v\n", err)
	// }
}
