package main

import (
	"errors"
	"math"
)

const (
	initialMaxSize = 50
	resizeFactor = 2
)

type HashMap struct{
	dataSlice []*LinkedList
	usedColumns int
}

type ValueNode struct {
	key string
	value interface{}
}


func NewHashMap() *HashMap{
	return &HashMap{
		dataSlice: make([]*LinkedList, initialMaxSize),
	}
}

func (h *HashMap) Set(key string, value interface{}) (bool, error){
	if key == ""{
		return false, errors.New("Empty key")
	}

	valueNode := &ValueNode{
		key: key,
		value: value,
	}

	vNode, idx := h.getValueNode(key)
	if vNode != nil {
		vNode.value = value
	} else if h.dataSlice[idx] != nil {
		h.dataSlice[idx] = h.dataSlice[idx].AddValue(valueNode)
	} else {
		h.dataSlice[idx] = &LinkedList{ value: valueNode }
		h.usedColumns++
		if h.usedColumns == len(h.dataSlice) - 1{
			newSlice := make([]*LinkedList, len(h.dataSlice)*resizeFactor)
			copy(newSlice, h.dataSlice)
			h.dataSlice = newSlice
		}
	}

	return true, nil
}


func (h *HashMap) Get(key string) interface{} {
	vNode, _ := h.getValueNode(key)
	
	if vNode != nil {
		return vNode.value
	}

	return nil
}

func (h *HashMap) hash(key string) int{
	hash := 0
	for pos, char := range key {
		hash += int(char) * int(math.Pow(31, float64(len(key)-pos+1)))
	}

	return hash
}

func (h *HashMap) index(hash int) int{
	return hash % len(h.dataSlice)
}

func valueNodeKeyComparator(iVnode interface{}, iKey interface{}) bool{
	vNode := iVnode.(*ValueNode)
	key := iKey.(string)

	return vNode.key == key
}

func (h *HashMap) getValueNode(key string) (*ValueNode, int) {
	idx := h.index(h.hash(key))
	if h.dataSlice[idx] != nil {
		node := h.dataSlice[idx].FindValue(key, valueNodeKeyComparator)
		if node != nil {
			return node.GetValue().(*ValueNode), idx
		}
	}
	return nil, idx
}