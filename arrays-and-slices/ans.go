package ans

// Sum takes a array and return the sum of the values
func Sum(numbers []int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumAll takes a few slices and return a new slice, each value in the new one is the sum of the all values of the old slice
func SumAll(numberToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numberToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails take a few slices and return a new slice, each value in the new one is the sum of the tail(all elements except the first element) of the old slice
func SumAllTails(numberToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numberToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
