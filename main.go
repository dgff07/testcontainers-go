package main

import (
	"fmt"

	"github.com/dgff07/testcontainers-go/internal"
)

func main() {
	fmt.Println("Hello, World!")
	calculator := internal.InitCalculator()
	calculator.Sub(1, 2)
}
