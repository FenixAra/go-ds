package sort

import "errors"

const (
	MERGE_SORT = iota
	SELECTION_SORT
	QUICK_SORT
)

var (
	ErrSortTypeUnsupported = errors.New("Unsupported sort type")
)

type Sort interface {
	Sort([]int) []int
}

func NewSort(sort int) (Sort, error) {
	switch sort {
	case MERGE_SORT:
		return &mergeSort{}, nil
	case SELECTION_SORT:
		return &selectionSort{}, nil
	case QUICK_SORT:
		return &quickSort{}, nil
	default:
		return nil, ErrSortTypeUnsupported
	}
}
