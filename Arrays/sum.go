package arrays

// Function to add integers of an Array and return it's sum
func Sum(numbers [5]int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}