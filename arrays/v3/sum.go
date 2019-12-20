package arrys

func Sum(arr []int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums[i] = 0
		} else {
			sums[i] = Sum(numbers[1:])
		}
	}
	return
}
