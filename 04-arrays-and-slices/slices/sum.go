package slices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	sums := make([]int, 0, len(numbers))

	for _, part := range numbers {
		sums = append(sums, Sum(part))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	sums := make([]int, 0, len(numbers))

	for _, part := range numbers {
		if len(part) == 0 {
			sums = append(sums, 0)
		} else {
			tail := part[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
