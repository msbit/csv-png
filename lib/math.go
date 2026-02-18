package lib

func Max(x int, y int) int {
	if x < y {
		return y
	}

	return x
}

func Min(x int, y int) int {
	if x < y {
		return x
	}

	return y
}
