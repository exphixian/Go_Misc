package main

import (
	"fmt"
)

func fibonacci(number uint) {

}

func main() {
	fmt.Println("Insert your number:")
	var number uint
	fmt.Scan(&number)

	sequence := []uint{}
	var latest uint
	sequence = append(sequence, 0, 1)
	i := 0

	if number == 0 {
		fmt.Println("Your Fibonacci sequence is: [0]")
	} else if number == 1 {
		fmt.Println("Your Fibonacci sequence is: [0, 1]")
	} else {
		for latest <= number {
			latest = sequence[i] + sequence[i+1]
			if latest <= number {
				sequence = append(sequence, latest)
				i++
			}
		}
		fmt.Printf("\nYour Fibonacci sequence is: %v\n", sequence)
	}
}
