package main

import "fmt"

func main() {
	var totalBottles, bottlesLeft int
	fmt.Println("Pick a number between 1 and 100:")
	fmt.Scan(&totalBottles)
	bottlesLeft = totalBottles

	for i := 0; i < totalBottles; i++ {
		fmt.Printf("There are %v bottles of beer left - you share one around. \n", bottlesLeft)
		bottlesLeft--
	}
}
