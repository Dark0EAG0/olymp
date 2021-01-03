package main

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if search(matrix[i], target) {
			return true
		}
	}
	return false
}

func search(m []int, t int) bool {
	min := 0
	max := len(m) - 1

	for min <= max {
		mid := (min + max) / 2

		if m[mid] < t {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	if min == len(m) || m[min] != t {
		return false
	}

	return true
}
