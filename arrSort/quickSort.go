package arrSort

import (
	"errors"
)

var SliceNilErr = errors.New("slice is nil")

// quick sort
func QuickSort(arr []int) error {
	if arr == nil {
		return SliceNilErr
	}
	if len(arr) < 2 {
		return nil
	}
	var l, r int
	r = len(arr) - 1
	pivot := arr[(l+r)/2]
	for l < r {
		// 从 pivot 的左边找到大于等于pivot的值
		for arr[l] < pivot {
			l++
		}
		// 从 pivot 的右边边找到小于等于pivot的值
		for arr[r] > pivot {
			r--
		}
		// 1 >= r 表明本次分解任务完成, break
		if l >= r {
			break
		}
		// 交换
		arr[l], arr[r] = arr[r], arr[l]
		// 优化
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	// 如果  1== r, 再移动下
	if l == r {
		l++
		r--
	}
	// 向左递归
	quickSort0(0, r, arr)
	// 向右递归
	quickSort0(l, len(arr)-1, arr)

	return nil
}

func quickSort0(l int, r int, arr []int) {
	if l > r {
		return
	}
	right := r
	left := l
	pivot := arr[(l+r)/2]
	for l < r {
		// 从 pivot 的左边找到大于等于pivot的值
		for arr[l] < pivot {
			l++
		}
		// 从 pivot 的右边边找到小于等于pivot的值
		for arr[r] > pivot {
			r--
		}
		// 1 >= r 表明本次分解任务完成, break
		if l >= r {
			break
		}
		// 交换
		arr[l], arr[r] = arr[r], arr[l]
		// 优化
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	// 如果  1== r, 再移动下
	if l == r {
		l++
		r--
	}

	// 向左递归
	quickSort0(left, r, arr)
	//向右递归
	quickSort0(l, right, arr)
}
