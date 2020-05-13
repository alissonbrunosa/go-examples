package algorithms

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := []int{6, 5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5, 6}

	BubbleSort(array)

	if !reflect.DeepEqual(array, expected) {
		t.Errorf("Expected sorted array to be %v, got %v \n", expected, array)
	}
}
