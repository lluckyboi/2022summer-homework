package queue

import (
	"errors"
	"reflect"
)

type Queue interface {
	push()
	pop()
	top()
	length()
}

type TSlice struct {
	I    []interface{}
	Head int
	Tail int
}

func (t *TSlice) Push(x interface{}) {
	if t.I[0] != nil {
		if reflect.TypeOf(t.I[0]) != reflect.TypeOf(x) {
			panic(errors.New("类型不一致"))
		}
	}
	t.I = append(t.I, nil)
	t.Tail++
	t.I[t.Tail] = x
}

func (t *TSlice) Pop() {
	t.Head++
}

func (t *TSlice) Top() interface{} {
	return t.I[t.Head]
}

func (t *TSlice) Length() int {
	return t.Tail - t.Head + 1
}
