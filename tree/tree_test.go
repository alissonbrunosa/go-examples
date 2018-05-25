package tree

import "testing"

func TestInsert(t *testing.T) {
	tree := New()

  tree.Insert(10)
  
  // Insert to left branch
  tree.Insert(9)
  tree.Insert(8)

  // Insert to right branch
  tree.Insert(11)
  tree.Insert(12)

  // Root data
	if tree.Root.data == 10 {
	} else {
		t.Errorf("Root data should be 10, got %v", tree.Root.data)
	}

	

  // Left data
	left := tree.Root.left
  
  if left.data == 9 {
	} else {
		t.Errorf("Left node data should be 9, got %v", left.data)
	}


  // Most left data
  left = left.left

  if left.data == 8 {
  } else {
    t.Errorf("Left node data should be 9, got %v", left.data)
  }

  
  // Right data
  right := tree.Root.right
  
  if right.data == 11 {
  } else {
    t.Errorf("Right node data should be 11, got %v", right.data)
  }


  // Most right data
  right = right.right

  if right.data == 12 {
  } else {
    t.Errorf("Right node data should be 12, got %v", right.data)
  }
}
