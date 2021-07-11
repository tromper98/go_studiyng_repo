package main

import (
	"fmt"
	"tempconv"
)

func main() {
	fmt.Printf("Брррр! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf(tempconv.CToF(32))
}
