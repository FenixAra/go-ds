package sort

type quickSort struct {
}

// Sorting an integer array using quick sort
func (q *quickSort) Sort(a []int) []int {
	return qSort(a, 0, len(a)-1)
}

func qSort(a []int, l, h int) []int {
	if l < h {
		a, i := partition(a, l, h)
		a = qSort(a, l, i-1)
		a = qSort(a, i+1, h)
		return a
	} else {
		return a
	}
}

func partition(a []int, l, h int) ([]int, int) {
	pivot := a[h]
	i := (l - 1)

	for j := l; j < h; j++ {
		if a[j] < pivot {
			i++
			swap(&a[i], &a[j])
		}
	}

	swap(&a[i+1], &a[h])
	return a, i + 1
}
