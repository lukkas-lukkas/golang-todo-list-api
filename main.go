package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Sleeping!")

	time.Sleep(10 * time.Second)

	fmt.Println("Hello, World!")
}
