package lists

type SinglyLinkedList struct {
	value interface{}
	next  *SinglyLinkedList
}

func (s *SinglyLinkedList) GetVal() interface{} {
	return s.value
}

func (s *SinglyLinkedList) SetVal(v interface{}) {
	s.value = v
}

func (s *SinglyLinkedList) Next() List {
	return s.next
}

func (s *SinglyLinkedList) SetNext(n List) {
	s.next = n.(*SinglyLinkedList)
}
