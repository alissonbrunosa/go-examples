package algorithms

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := []int{5, 1, 4, 2, 8}
	expected := []int{1, 2, 4, 5, 8}

	BubbleSort(array)

	if !reflect.DeepEqual(array, expected) {
		t.Errorf("Expected sorted array to be %v, got %v \n", expected, array)
	}
}
