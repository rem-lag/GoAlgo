package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	s := 0

	for _, i := range numbers {
		s = s + i
	}

	return s
}

func SumRecursive(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + SumRecursive(numbers[1:])
}
