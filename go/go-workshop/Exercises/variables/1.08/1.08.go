// Changing Multiple Values at Once

package main

import (
	"fmt"
)

func main() {
	query, limit, offset := "bat", 10, 0
	query, limit, offset = "ball", offset, 30
	fmt.Println(query, limit, offset)
}
