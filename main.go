package main

import (
	"fmt"

	"aspect/demo"
)

func main() {
	x := demo.UseUnsafeInt()
	fmt.Printf("unsafe int: %v", x)
}
