package main

import "math"

func main() {
	//
}

func isPowerOfThree(num int) bool {
	if num == 0 {
		return false
	}

	b := 3.0
	s := math.Round(math.Log(float64(num)) / math.Log(float64(b)))

	if math.Pow(b, s) == float64(num) {
		return true
	}

	return false
}
