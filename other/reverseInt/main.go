package reverseint

import (
	"strconv"
)

func r(n int) int {
	if n == 0 {
		return 0
	}

	res := ""

	if n < 0 {
		n = -1 * n
		res = "-"
	}

	s := strconv.Itoa(n)

	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}

	r, _ := strconv.Atoi(res)
	return r
}
