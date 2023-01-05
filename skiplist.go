package mycode

import (
	"fmt"
	"math/rand"
	"time"
)

type Element struct {
	Key   int
	Value interface{}
	Level []*Element
}

type SkipList struct {
	Head     *Element
	Len      int
	MaxLevel int
}

func NewSkipList(level int) *SkipList {
	return &SkipList{
		Head: &Element{
			Level: make([]*Element, level),
		},
		MaxLevel: level,
	}
}

func (sl *SkipList) Search(key int) (*Element, bool) {
	cur := sl.Head
	for i := sl.MaxLevel - 1; i >= 0; i-- {
		for cur.Level[i] != nil && cur.Level[i].Key < key {
			cur = cur.Level[i]
		}
	}

	cur = cur.Level[0]
	if cur != nil && cur.Key == key {
		return cur, true
	}

	return nil, false
}

func (sl *SkipList) Insert(key int, val interface{}) *Element {
	ele, ok := sl.Search(key)
	if ok {
		ele.Value = val
		return ele
	}

	level := sl.RandomLevel()
	fmt.Printf("key:%v,level:%v\n", key, level)
	newElement := &Element{
		Key:   key,
		Value: val,
		Level: make([]*Element, sl.MaxLevel),
	}

	cur := sl.Head
	for i := level; i >= 0; i-- {
		for cur.Level[i] != nil && cur.Level[i].Key < key {
			cur = cur.Level[i]
		}

		newElement.Level[i] = cur.Level[i]
		cur.Level[i] = newElement
	}
	sl.Len++

	return newElement
}

func (sl *SkipList) Delete(key int) {
	cur, ok := sl.Search(key)
	if !ok {
		return
	}

	front := sl.Head
	for i := sl.MaxLevel - 1; i >= 0; i-- {
		for front.Level[i] != nil && front.Level[i].Key < key {
			front = front.Level[i]
		}

		if front.Level[i] != nil && front.Level[i] == cur {
			front.Level[i] = cur.Level[i]
		}
	}
}

// RandomLevel 随机层数，范围：[0,MaxLevel)
func (sl *SkipList) RandomLevel() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(sl.MaxLevel)
}

func (sl *SkipList) Print() {
	head := sl.Head
	for i := sl.MaxLevel - 1; i >= 0; i-- {
		cur := head.Level[i]
		fmt.Printf("%v: ", i)
		for cur != nil {
			fmt.Printf("(%v,%v) ", cur.Key, cur.Value)
			cur = cur.Level[i]
		}
		fmt.Println()
	}
}
