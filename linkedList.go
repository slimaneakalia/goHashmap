package hashMap

import "fmt"

type LinkedList struct{
	value interface{}
	next *LinkedList
}

type EqualityComparatorFunc func(interface{}, interface{}) bool

func (head *LinkedList) GetValue() interface{}{
	return head.value
}

func (head *LinkedList) AddValue(v interface{}) *LinkedList{
	return &LinkedList{
		value: v,
		next: head,
	}
}

func defaultEqualityComparator(a interface{}, b interface{}) bool {
	return a == b
}


func (head *LinkedList) FindValue(value interface{},
	equalityComparator EqualityComparatorFunc) *LinkedList{
	tmp := head
	eqCmp := equalityComparator
	if equalityComparator == nil {
		eqCmp = defaultEqualityComparator
	}

	for tmp != nil {
		if eqCmp(tmp.value, value) {
			return tmp
		}

		tmp = tmp.next
	}

	return nil
}

func (head *LinkedList) Print() string{
	str := ""
	tmp := head
	for tmp != nil {
		str += fmt.Sprintf("%v - ", tmp.value)
		tmp = tmp.next
	}

	return str
}