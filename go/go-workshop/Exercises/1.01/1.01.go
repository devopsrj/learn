package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1 // Generate a random number between 0 and then add 1 to get a number between 1 and 5
	x := rand.Intn(20)
	stars := strings.Repeat("*", r) // Use the string repeater to create a string with the number of stars
	wars := strings.Repeat("#", x)
	fmt.Println(stars)
	fmt.Println(wars)
}
