package structsmethodsinterfaces

type Rectangle struct{
	Width float64
	Height float64
}

// Function to calculate Perimeter of Rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width);
}

// Function to calculate Area of Rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}