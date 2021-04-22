package artifacts

import (
	"testing"

	"gotest.tools/assert"
)

func TestHeapSort(t *testing.T) {
	arr := []int{0, 34, 23, 45, 234, 15, 87, 20}
	result := HeapSort(arr)
	expectation := []int{0, 15, 20, 23, 34, 45, 87, 234}
	assert.DeepEqual(t, result, expectation)
}
