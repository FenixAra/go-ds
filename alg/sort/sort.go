package sort

import "errors"

const (
	MERGE_SORT = iota
	SELECTION_SORT
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
	default:
		return nil, ErrSortTypeUnsupported
	}
}
