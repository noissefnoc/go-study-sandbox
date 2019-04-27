package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// initialize random seed
	rand.Seed(time.Now().UnixNano())

	// count command line arguments
	// if not specified, then raise error
	c := len(os.Args) - 1
	if c < 1 {
		fmt.Fprintf(os.Stderr, "[usage] %s choice1 choice2...", os.Args[0])
		os.Exit(1)
	}

	// print element which choice randomly
	fmt.Printf(os.Args[rand.Intn(c)+1])
}
