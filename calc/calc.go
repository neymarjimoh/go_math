package calc

func Add(args  ...int) int {
	sum := 0

	for _, val := range args {
		sum += val
	}

	return sum
}

func Sub(a, b int) int {
	return a - b
}