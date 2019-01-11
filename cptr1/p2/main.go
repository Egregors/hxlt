package diff

func diff(a, b int) int {
	if a == b {
		return 0
	}

	if a != 0 {
		if b < a {
			a, b = b, a
		}

		b = b - a
		a = 0

	}

	if b <= 180 {
		return b
	}
	return 360 - b

}
