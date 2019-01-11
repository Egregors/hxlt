package main

import "testing"

func TestIsPowerOfThree(t *testing.T) {
	if isPowerOfThree(0) != false {
		t.Error(0)
	}
	if isPowerOfThree(1) != true {
		t.Error(1)
	}
	if isPowerOfThree(2) != false {
		t.Error(2)
	}
	if isPowerOfThree(9) != true {
		t.Error(9)
	}
}
