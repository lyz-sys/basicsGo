package arrSort

import "errors"

// bubble sort
func BubbleSort(arr []int) error {
	if arr == nil {
		return errors.New("slice is nil")
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
