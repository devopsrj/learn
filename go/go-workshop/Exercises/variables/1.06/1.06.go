// Declaring Multiple Variables with a Short Variable Declaration

package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	Debug, LogLevel, startUpTime := getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
}
