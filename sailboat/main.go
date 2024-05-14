//This code will output a sailboat in the terminal by using go run main.go

package main

import (
	"fmt"
)

// main my main sailboat function
func main() {
	drawSailboat()
}

func drawSailboat() {
	fmt.Println("    /|    ")
	fmt.Println("   / |    ")
	fmt.Println("  /__|    ")
	fmt.Println("     |    ")
	fmt.Println("_____|____")
	fmt.Println("\\________/")

}
