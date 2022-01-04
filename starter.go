package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Who is your favorite Pokemon starter?\n")
	input.Scan()

	favStarter := input.Text()
	fmt.Println(favStarter + " is not as good as Pikachu!\n")

}
