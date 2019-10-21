package main

import "fmt"

func main() {
	mb := "日本a語b"
	mbRune := []rune(mb)
	size := len(mbRune)

	fmt.Printf("mbRune = %d characters : ", size)
	for i := 0; i < size; i++ {
		fmt.Printf("%#U ", mbRune[i])
	}
	fmt.Print("\n")
}
