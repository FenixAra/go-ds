package lists

type ListType int

const (
	SINGLY_LINKED_LIST ListType = iota
)

type List interface {
	GetVal() interface{}
	SetVal(v interface{})
	Next() List
	SetNext(n List)
}

func GetList(t ListType) List {
	switch t {
	case SINGLY_LINKED_LIST:
		return &SinglyLinkedList{}
	}
	return nil
}
