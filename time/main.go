package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const defaultTimeFormat = "20060102"

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	d1, err := time.Parse(defaultTimeFormat, "20190101")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	d2, err := time.Parse(defaultTimeFormat, "20190105")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	duration := int(d2.Sub(d1).Seconds() / 24 / 60 / 60)
	if num >= 1 {
		fmt.Println(d1.Format(defaultTimeFormat))
	}
	if num >= 3 {
		step := (duration % num) + 1

		for i := step; i < duration; i += step {
			fmt.Println(d1.AddDate(0, 0, i).Format(defaultTimeFormat))
		}
	}
	if num >= 2 {
		fmt.Println(d2.Format(defaultTimeFormat))
	}
}
