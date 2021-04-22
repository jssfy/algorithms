package artifacts

import "algorithms/utils"

const ROOT = 1

// ignore arr[0]
func HeapSort(arr []int) []int {
	arrLen := len(arr) - 1
	buildMaxHeap(arr, arrLen)
	for i := arrLen; i >= ROOT; i-- {
		utils.Swap(arr, ROOT, i)
		arrLen -= 1
		heapify(arr, ROOT, arrLen)
	}
	return arr
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= ROOT; i-- {
		heapify(arr, i, arrLen)
	}
}

// top down heapify
func heapify(arr []int, i, arrLen int) {
	left := 2 * i
	right := 2*i + 1
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		utils.Swap(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}
