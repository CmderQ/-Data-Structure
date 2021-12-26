// 线性表顺序存储
// 的各种方法
package linear

import (
	"fmt"
)

// 逻辑结构: 线性表(具有相同数据元素是有限序列)
type SequenceList struct {
	data *[]interface{}
	len  int
	cap  int
}

// 初始化
func (s *SequenceList) InitList(InitSize int) {
	data := make([]interface{}, InitSize)
	s.data = &data
	s.len = 0
	s.cap = InitSize
}

// 清空
func (s *SequenceList) ClearList() {
	s.data = nil
	s.len = 0
}

//销毁
func (s *SequenceList) DestroyList() {
	s.data = nil
	s.len = 0
	s.cap = 0
}

// 长度
func (s *SequenceList) ListLength() int {
	return s.len
}

// 判断是否为空
func (s *SequenceList) ListEmpty() bool {
	if s.len != 0 {
		return false
	}
	return true
}

// 判断是否满
func (s *SequenceList) ListIsFull() bool {
	if s.len == s.cap {
		return true
	}

	return false
}

// 根据下标获取元素
func (s *SequenceList) ListGet(index int) (interface{}, bool) {
	if index < 0 || index > s.len {
		return nil, false
	}

	return (*s.data)[index], true
}

// 按值查找(查找第一个元素值等于e的元素,返回其位序)
func (s *SequenceList) ListLocal(elem interface{}) (int, bool) {
	for index := 0; index < s.len; index++ {
		if elem == (*s.data)[index] {
			return index, true
		}
	}

	return -1, false
}

// 获取当前元素的前一个元素
func (s *SequenceList) ListPriorElem(elem interface{}) (interface{}, bool) {
	index, _ := s.ListLocal(elem)
	//如果下标不存在，或者是第一个元素(则此时没有前驱元素)
	if index == -1 || index == 0 {
		return nil, false
	}

	return (*s.data)[index-1], true
}

// 获取当前元素的下一个元素
func (s *SequenceList) ListNextElem(elem interface{}) (interface{}, bool) {
	index, _ := s.ListLocal(elem)
	//如果下标不存在，或者是最后一个元素(则此时没有下一个元素)
	if index == -1 || index == s.len-1 {
		return nil, false
	}

	return (*s.data)[index+1], true
}

// 插入元素，index为插入的位置，elem为插入值
func (s *SequenceList) ListInsert(index int, elem interface{}) bool {
	// 判断下标是否正确，以及是否满了
	if index < 0 || index > s.cap || s.ListIsFull() {
		return false
	}

	// 先将index位置的元素及之后的元素后移一位
	for i := s.len - 1; i >= index; i-- {
		(*s.data)[i+1] = (*s.data)[i]
	}

	// 插入元素
	(*s.data)[index] = elem
	s.len++

	return true
}

// 删除指定位置的元素
func (s *SequenceList) ListDelete(index int) bool {
	// 判断下标是否正确，以及是否满了
	if index < 0 || index > s.cap || s.ListIsFull() {
		return false
	}

	//删除指定位置的元素后，后面的元素前移一位
	for i := index; i < s.len-1; i++ {
		(*s.data)[i] = (*s.data)[i+1]
	}
	s.len--

	return true
}

// 遍历输出元素
func (s *SequenceList) ListTraverse() {
	for index := 0; index < s.len; index++ {
		fmt.Println((*s.data)[index])
	}
}
