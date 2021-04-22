package artifacts

import "fmt"

// Value to sort
type Value int32

// Sort entrypoint
// TBD: to introduce an interface
func Sort(data []Value, len int) {
	// TBD: what if len+1 overflows

	// sort(data, start, end) = merge(sort(data, start, mid), sort(data, mid+1, end))
	// stop if start >= end
	fmt.Println("dummy")
}

// sorting in the recursive way
func divideAndConquer(data []Value, start int, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	divideAndConquer(data, start, mid)
	divideAndConquer(data, mid+1, end)
	merge(data, start, mid, end)
}

//
func merge(data []Value, start int, mid int, end int) {
	// TBD
}

// swap by slice index
// func swap(data []Value, i int, j int) {
// 	// check if i & j valid
// 	if i > len(data) || j > len(data) {
// 		panic("out of index")
// 	}

// 	// swapping
// 	temp := data[i]
// 	data[i] = data[j]
// 	data[j] = temp
// }
