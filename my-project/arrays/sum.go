package main

func Sum(numbers []int) int {
	sum := 0 
	for _, number := range numbers {
		sum += number
	// within the array numbers, we iterate over each number and add it to sum
	// each iteration, sends index and value, we use _ to ignore index
	}
	return sum
}