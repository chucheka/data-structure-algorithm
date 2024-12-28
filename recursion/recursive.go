package recursion

func Factorial(n uint) uint {

	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Pow(n, x uint) uint {

	if x == 0 {
		return 1
	}

	if x == 1 {
		return 3
	}

	return n * Pow(n, x-1)
}
