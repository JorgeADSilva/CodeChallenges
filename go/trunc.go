package main

import (
	"fmt"
	"os"
)
/*
Write a program which prompts the user to enter a 
floating point number and prints the integer which 
is a truncated version of the floating point number 
that was entered. Truncation is the process of removing 
the digits to the right of the decimal place.
*/

func main() {
	var x float64
	fmt.Print("Enter a floating point number ->")
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	i := int(x)
	fmt.Println(i)
}