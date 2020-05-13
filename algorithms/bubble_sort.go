package algorithms

func BubbleSort(list []int) {
	size := len(list) - 1

	for i := 0; i <= size; i++ {
		swapped := false
		for j := 0; j < size-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
