package perfect

func isPerfect(n int) bool {

	if n <= 0 {
		return false
	}

	sum := 0
	for i := n - 1; i > 0; i-- {
		if n%i == 0 {
			sum += i
		}
	}

	return sum == n
}
