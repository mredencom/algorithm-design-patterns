package array

import "errors"

type ArrayOperate interface {
	Len() uint
	Find(idx uint) (interface{}, error)
	Insert(idx uint, val interface{}) error
	Delete(idx uint) (interface{}, error)
}

type Array struct {
	data []interface{}
	len  uint
}

func NewArray(len uint) *Array {
	if len == 0 {
		return nil
	}
	return &Array{
		data: make([]interface{}, len, len),
		len:  len,
	}
}

func (a *Array) idxOutOfRange(idx uint) bool {
	if idx >= a.len {
		return true
	}
	return false
}

func (a *Array) Len() uint {
	return a.len
}

func (a *Array) Find(idx uint) (interface{}, error) {
	if a.idxOutOfRange(idx) {
		return nil, errors.New("数组越界")
	}
	return a.data[idx], nil
}

func (a *Array) Insert(idx uint, val interface{}) error {
	if a.len == uint(cap(a.data)) {
		return errors.New("数组满员了，装不下了")
	}
	if a.idxOutOfRange(idx) {
		return errors.New("数组越界")
	}
	// 需要把数据前面的向后移腾出位置，给新元素
	for i := a.len; i > idx; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[idx] = val
	a.len++
	return nil
}

func (a *Array) Delete(idx uint) (interface{}, error) {
	if a.idxOutOfRange(idx) {
		return nil, errors.New("数组越界")
	}
	val := a.data[idx]
	// 把数据元素前移
	for i := idx; i < a.len-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.len--
	return val, nil
}
