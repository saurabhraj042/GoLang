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
func SumAll(numbers ...[]int) (sums []int){
	return
}