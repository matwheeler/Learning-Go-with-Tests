package main

func Sum(numbers []int) (sum int) {
	for _, value := range numbers {
		sum += value
	}

	return sum
}