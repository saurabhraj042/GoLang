package arrays

// Function to add integers of an Array and return it's sum
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// Function to return sum of given Slices as an Slice itself
func SumAll(numbersToSum ...[]int) (sums []int){
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

// Function to find the Sum of Tails of given Slices
func SumTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return
}