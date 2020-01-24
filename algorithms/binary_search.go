package algorithms

var count = 0

func BinarySearch(ints []int, wanted int) int {
	if len(ints) == 0 {
		return -1
	}

	return binarySeach(ints, 0, len(ints)-1, wanted)
}

func binarySeach(ints []int, left int, right int, wanted int) int {
	for left <= right {
		middle := (left + right) >> 1
		if ints[middle] == wanted {
			return middle
		} else if ints[middle] > wanted {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}
