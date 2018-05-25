package list

import "testing"

func TestAdd(t *testing.T) {
	list := New()
	list.Add(1)

	if list.Size() != 1 {
		t.Errorf("List's size should be 1, got %v \n", list.Size())
	}
}

func TestDeleteAt(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)

	list.DeleteAt(0)

	if list.Size() == 1 {
	} else {
		t.Error("It should have size == 1")
	}

	element, _ := list.Get(0)
	if element == 2 {
	} else {
		t.Errorf("First element after remove should be 2, got %v", element)
	}
}

func TestIncreaseCapacity(t *testing.T) {
	list := New()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	list.Add(11)

	if list.Size() == 11 {
	} else {
		t.Error("It must double list's size")
	}
}

func TestGet(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	element, _ := list.Get(1)
	if element != 2 {
		t.Errorf("Element at index 1 should be 2, got %v \n", element)
	}
}

func TestGetWithNegativeIndex(t *testing.T) {
	list := New()
	list.Add(1)

	_, err := list.Get(-1)
	if err == nil {
		t.Errorf("Should return a error")
	}
}
