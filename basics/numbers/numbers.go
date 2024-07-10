package numbers

func Add(x int, y int) int { // see how the method name needs to be capatlized to be used in other areas
	// it's by convention that upper case is exported to other packages
	// interal methods should be lower cased and not exposed anywhere else
	return x + y
}

func Subtract(x int, y int) int {
	return x - y
}
