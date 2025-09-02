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

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	// use the append function which takes a slice and a new value,
	// then returns a new slice with all the items in it.
	for _, numbers := range numbersToSum {

		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		} else {	
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
