package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const defaultTimeFormat = "20060102"

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("invalid argument")
		os.Exit(1)
	}
	num, err := strconv.Atoi(os.Args[1])
	d1Str := os.Args[2]
	d2Str := os.Args[3]
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	d1, err := time.Parse(defaultTimeFormat, d1Str)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	d2, err := time.Parse(defaultTimeFormat, d2Str)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	duration := int(d2.Sub(d1).Seconds()/24/60/60) + 1
	if num >= 1 {
		fmt.Println(d1.Format(defaultTimeFormat))
	}
	if num >= 3 {
		step := (duration - 2) / (num - 2)

		for i := step; i < duration-1; i += step {
			fmt.Println(d1.AddDate(0, 0, i).Format(defaultTimeFormat))
		}
	}
	if num >= 2 {
		fmt.Println(d2.Format(defaultTimeFormat))
	}
}
