package sort

type mergeSort struct {
}

// Sorting an integer array using merge sort
func (m *mergeSort) Sort(a []int) []int {
	return mSort(a, 0, len(a)-1)
}

func mSort(a []int, l, r int) []int {
	if l < r {
		m := l + (r-l)/2
		left := mSort(a, l, m)
		right := mSort(a, m+1, r)
		return merge(left, right)
	} else if l == r {
		return []int{a[l]}
	} else {
		return nil
	}
}

func merge(l []int, r []int) []int {
	var a []int
	j := 0
	for i := 0; i < len(l); i++ {
		if l[i] <= r[j] {
			a = append(a, l[i])
		} else {
			a = append(a, r[j])
			j++
			if j == len(r) {
				a = append(a, l[i:len(l)]...)
				break
			}
			i--
		}
	}

	if j < len(r) {
		a = append(a, r[j:len(r)]...)
	}
	return a
}
