package main

import (
	"fmt"
	"math"
)

func f(x float64) (y float64) {
	// return 3(x^2) + 2x + 1
	return math.Pow(x, 2)*3 + x*2 + 1
}

func sadari_integration(start float64, end float64, width float64) (result float64) {
	var i float64
	result = 0.0
	for i = start; i < end; i += width {
		result += math.Abs((f(i)+f(i+width))*width) / 2
	}

	return result
}

func main() {
	var start, end float64

	fmt.Print("[*] start index : ")
	fmt.Scan(&start)
	fmt.Print("[*] end index : ")
	fmt.Scan(&end)

	result := sadari_integration(start, end, 0.0001)
	fmt.Println("[*] Integration success!")
	fmt.Println(" -> ", result)
}
