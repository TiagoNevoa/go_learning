package main

import (
	// Package bufio implements buffered Input/Output (Reader,Writer,Scanner).
	"bufio"
	// Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	"fmt"
	// Package math provides basic constants and mathematical functions.
	"math"
	// Package os provides a platform-independent interface to operating system functionality(error handling).
	"os"
	// Package strconv implements conversions to and from string representations of basic data types.
	"strconv"
	// Package strings implements simple functions to manipulate UTF-8 encoded strings.
	"strings"
)

func main() {

	// read line from terminal
	reader := bufio.NewReader(os.Stdin)
	// write on the terminal
	fmt.Print("Value 1: ")
	// save the line on var
	input1, _ := reader.ReadString('\n')
	// convert input String to float
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	// check if there is err
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	fmt.Print("Value 2: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}

	sum := float1 + float2
	sum = math.Round(sum*100) / 100
	fmt.Printf("The sum of %v and %v is %v\n\n", float1, float2, sum)
}
