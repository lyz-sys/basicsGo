package arrSort

// bubble sort
func BubbleSort(arr []int) error {
	if arr == nil {
		return SliceNilErr
	}
	if len(arr) < 2 {
		return nil
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return nil
}
