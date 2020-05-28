package main

func Sum(numbers []int) (sum int) {
	for _, value := range numbers {
		sum += value
	}

	return sum
}

// ... this allows for a variable amount of parameters to be passed to the function
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, value := range numbersToSum {
		sums = append(sums, Sum(value))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, value := range numbersToSum {
		if len(value) > 1 {
			numbers := value[1:] // create a slice of the value array from index 1 to the end.
			sums = append(sums, Sum(numbers))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
