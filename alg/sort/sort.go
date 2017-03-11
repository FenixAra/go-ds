package sort

import "errors"

const (
	MERGE_SORT = iota
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
	default:
		return nil, ErrSortTypeUnsupported
	}
}
