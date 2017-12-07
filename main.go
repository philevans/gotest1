package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Println(math.Pi)
	fmt.Println(math.SqrtPi)
	var a chan int        	// Define a channel of type int
	a = make(chan int, 1) 	// Actually make the channel with a buffer size of 1 item
	a <- 4                	// Put an integer into the channel
	fmt.Println(<-a)      	// Retrieve the integer from the channel and print it on screen
	// another comment
}
