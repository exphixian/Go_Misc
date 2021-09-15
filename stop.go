//random script to practice go - will choose song lyrics based on console input.

package main

import "fmt"

func main() {
	var first int8
	fmt.Println("1 or 2? ")
	fmt.Scanln(&first)

	fmt.Println("STOP")
	if first == 1 {
		fmt.Print("Hammertime!")
	} else {
		fmt.Print("Collaborate and listen.")
	}
}
