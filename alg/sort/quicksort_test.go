package sort

import (
	"testing"

	"github.com/FenixAra/go-ds/helpers"
)

func TestQuickSort(t *testing.T) {
	a := []int{4, 6, 1, 3, 2, 10}
	mSort, _ := NewSort(QUICK_SORT)

	sortedArray := mSort.Sort(a)
	helpers.AssertDeepEqual([]int{1, 2, 3, 4, 6, 10}, sortedArray, t)
}

func TestQuickSort1(t *testing.T) {
	a := []int{7, 1, 8, 3, 0}
	mSort, _ := NewSort(QUICK_SORT)

	sortedArray := mSort.Sort(a)
	helpers.AssertDeepEqual([]int{0, 1, 3, 7, 8}, sortedArray, t)
}

func TestQuickSort2(t *testing.T) {
	a := []int{21, 11, 0, 33, 55, 44, 66, 99, 88, 77}
	mSort, _ := NewSort(QUICK_SORT)

	sortedArray := mSort.Sort(a)
	helpers.AssertDeepEqual([]int{0, 11, 21, 33, 44, 55, 66, 77, 88, 99}, sortedArray, t)
}

func BenchmarkQuickSort(b *testing.B) {
	a := helpers.RandomIntArray(1000)
	mSort, _ := NewSort(QUICK_SORT)

	for n := 0; n < b.N; n++ {
		mSort.Sort(a)
	}
}
