///Little arguementative script just for fun. Prints the opposite bool of the input.

package main

import "fmt"

func main() {
	var oppositeDay bool
	fmt.Println("True or False: Is today opposite day?")
	fmt.Scan(&oppositeDay)
	fmt.Print(!oppositeDay, "!")
}
