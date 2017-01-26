package lists

import (
	"testing"

	"github.com/FenixAra/go-ds/helpers"
)

func TestIntGetVal(t *testing.T) {
	var v int = 10
	list := &SinglyLinkedList{
		value: v,
		next:  nil,
	}

	nv := list.GetVal()
	helpers.AssertEqual(v, nv, t)
}

func TestStringGetVal(t *testing.T) {
	var v string = "Hello"
	list := &SinglyLinkedList{
		value: v,
		next:  nil,
	}

	nv := list.GetVal()
	helpers.AssertEqual(v, nv, t)
}

func TestFloatGetVal(t *testing.T) {
	var v float32 = 100.11
	list := &SinglyLinkedList{
		value: v,
		next:  nil,
	}

	nv := list.GetVal()
	helpers.AssertEqual(v, nv, t)
}

func TestIntSetVal(t *testing.T) {
	list := GetList(SINGLY_LINKED_LIST)
	var v int = 10

	list.SetVal(v)
	nv := list.GetVal()
	helpers.AssertEqual(v, nv, t)
}

func TestStringSetVal(t *testing.T) {
	list := GetList(SINGLY_LINKED_LIST)
	var v string = "Hello"

	list.SetVal(v)
	nv := list.GetVal()
	helpers.AssertEqual(v, nv, t)
}

func TestFloatSetVal(t *testing.T) {
	list := GetList(SINGLY_LINKED_LIST)
	var v float32 = 100.11

	list.SetVal(v)
	nv := list.GetVal()
	helpers.AssertEqual(v, nv, t)
}

func TestSetNext(t *testing.T) {
	list := GetList(SINGLY_LINKED_LIST)
	var v int = 10
	list.SetVal(v)
	next := GetList(SINGLY_LINKED_LIST)
	var nv int = 20
	next.SetVal(nv)

	list.SetNext(next)
	helpers.AssertDeepEqual(next, list.Next(), t)
}

func TestGetNext(t *testing.T) {
	list := GetList(SINGLY_LINKED_LIST)
	var v int = 10
	list.SetVal(v)
	next := GetList(SINGLY_LINKED_LIST)
	var nv int = 20
	next.SetVal(nv)

	list.SetNext(next)
	helpers.AssertDeepEqual(next, list.Next(), t)
}
