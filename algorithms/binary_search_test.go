package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6}

	for index, element := range array {
		if idx := BinarySearch(array, element); index != idx {
			t.Errorf("Index of %v should be %v, got %v\n", element, index, idx)
		}
	}

	index := BinarySearch(array, 9)
	if index != -1 {
		t.Errorf("Expected index is -1, got %v", index)
	}
}
