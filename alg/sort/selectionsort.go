package sort

type selectionSort struct {
}

// Sorting an integer array using selection sort
func (s *selectionSort) Sort(a []int) []int {
	for i := range a {
		j := minOfArray(a[i+1 : len(a)])
		if j == -1 {
			break
		}

		j += i + 1
		if a[i] > a[j] {
			a = swap(a, i, j)
		}
	}
	return a
}

func minOfArray(a []int) int {
	if len(a) == 0 {
		return -1
	}

	min := a[0]
	pos := 0
	for i, v := range a {
		if v < min {
			min = v
			pos = i
		}
	}
	return pos
}

func swap(a []int, i, j int) []int {
	t := a[i]
	a[i] = a[j]
	a[j] = t
	return a
}
