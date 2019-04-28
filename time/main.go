package main

import (
	"fmt"
	"time"
)

const defaultTimeFormat = "20060102"

func main() {
	d1, err := time.Parse(defaultTimeFormat, "20190101")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	d2, err := time.Parse(defaultTimeFormat, "20190105")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	sub := d2.Sub(d1).Seconds() / 24 / 60 / 60
	fmt.Printf("%f\n", sub)
}
