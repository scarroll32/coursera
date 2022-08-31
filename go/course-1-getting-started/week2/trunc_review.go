package main

import "fmt"

func main() {
	var floatNum float64
	fmt.Printf("Type a float value please:")
	fmt.Scan(&floatNum)
	intNum := int(floatNum)
	fmt.Printf("%d", intNum)
}

