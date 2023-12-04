package day3

type point struct {
	rowIndex int
	min      int
	max      int
}

type pointArr []point

func (s pointArr) contains(e point) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func inRange(min int, max int, min2 int, max2 int) bool {
	// check if any number is within the range of the other
	aInB := min >= min2 && min <= max2
	bInA := min2 >= min && min2 <= max

	return aInB || bInA
}
