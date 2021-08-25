package structsmethodsinterfaces

type Rectangle struct{
	Width float64
	Height float64
}

// Function to calculate Perimeter of Rectangle
func Perimeter(height, width float64) float64 {
	return 2 * (height + width);
}

// Function to calculate Area of Rectangle
func Area(height, width float64) float64 {
	return height * width
}