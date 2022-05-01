package linear

import (
	"sync"
)

// DoubleNode 双链表结点
type DoubleNode struct {
	data *[]interface{} // 数据
	Prev *DoubleNode    // 上一个结点
	Next *DoubleNode    // 下一个结点
}

// DoubleList 双链表
type DoubleList struct {
	mutex    *sync.RWMutex
	Capacity uint        //最大容量
	Size     uint        //当前容量
	Head     *DoubleNode // 头结点
	Tail     *DoubleNode // 尾结点
}

// New 创建双链表对象
func New(capacity uint) *DoubleList {
	list := new(DoubleList)
	list.mutex = new(sync.RWMutex)
	list.Capacity = capacity
	list.Size = 0
	list.Head = nil
	list.Tail = nil
	return list
}

// AddHead 添加头部节点-----头部插入法
//实现思路
//先判断容量大小
//判断头部是否为空,
//
//如果为空则添加新节点
//如果不为空则更改现有的节点,并添加上
func (list *DoubleList) AddHead(node *DoubleNode) bool {
	//判断容量是否为0
	if list.Capacity == 0 {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	//判断头部节点是否为nil
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else { //存在头部节点
		list.Head.Prev = node //将旧头部节点上一个节点指向新节点
		node.Next = list.Head //新头部节点下一个节点指向旧头部节点
		list.Head = node      //设置新的头部节点
		list.Head.Prev = nil  //将新的头部节点上一个节点设置nil
	}
	list.Size++
	return true
}

// AddTail 添加尾部元素
//实现思路
//先判断容量大小
//判断尾部是否为空,
//
//如果为空则添加新节点
//如果不为空则更改现有的节点,并添加上
func (list *DoubleList) AddTail(node *DoubleNode) bool {
	//判断是否有容量,
	if list.Capacity == 0 {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	//判断尾部是否为空
	if list.Tail == nil {
		list.Tail = node
		list.Head = node
	} else {
		//旧的尾部下个节点指向新节点
		list.Tail.Next = node
		//追加新节点时,先将node的上节点指向旧的尾部节点
		node.Prev = list.Tail
		//设置新的尾部节点
		list.Tail = node
		//新的尾部下个节点设置为空
		list.Tail.Next = nil
	}
	//双链表大小+1
	list.Size++
	return true
}

// Insert 添加任意位置元素
//实现思路
//判断容量大小
//判断链表大小
//第一种: 如果插入的位置大于当前长度则尾部添加
//第二种: 如果插入的位置等于0则,头部添加
//第三种: 中间插入节点
//
//先取出要插入位置的节点,假调为C节点
//介于A, C之间, 插入一个B节点
//A的下节点应该是B, 即C的上节点的下节点是B
//B的上节点是C的上节点
//B的下节点是C
func (list *DoubleList) Insert(index uint, node *DoubleNode) bool {
	//容量满了
	if list.Size == list.Capacity {
		return false
	}
	//如果没有节点
	if list.Size == 0 {
		return list.AddTail(node)
	}
	//如果插入的位置大于当前长度则尾部添加
	if index > list.Size {
		return list.AddTail(node)
	}
	//如果插入的位置等于0则,头部添加
	if index == 0 {
		return list.AddHead(node)
	}
	//取出要插入位置的节点
	nextNode := list.Get(index)
	list.mutex.Lock()
	defer list.mutex.Unlock()
	//中间插入:
	//假设已有A, C节点, 现在要插入B节点
	// nextNode即是C节点,
	//A的下节点应该是B, 即C的上节点的下节点是B
	nextNode.Prev.Next = node
	//B的上节点是C的上节点
	node.Prev = nextNode.Prev
	//B的下节点是C
	node.Next = nextNode
	//C的上节点是B
	nextNode.Prev = node
	list.Size++
	return true
}

// Get 获取某一元素
func (list *DoubleList) Get(index uint) *DoubleNode {
	//如果是index = 0则返回头部
	if index == 0 {
		return list.Head
	}
	//如果超出或等于当前链大小,则返回尾部
	if index >= list.Size {
		return list.Tail
	}
	//如果中间,则需要循环index次数的链表
	var i uint
	node := list.Head
	for i = 1; i < index; i++ {
		node = node.Next
	}
	return node
}

// RemoveHead 删除头部节点
//实现思路
//判断头部是否为空
//将头部节点取出来
//判断头部是否有下一个节点
//
//没有下一个节点, 则说明只有一个节点, 删除本身, 无须移动指针位置
//如果有下一个节点, 则头部下一个节点即是头部节点.
func (list *DoubleList) RemoveHead() *DoubleNode {
	//判断头部节点是否为空
	if list.Head == nil {
		return nil
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	//将头部节点取出来
	node := list.Head
	//判断头部是否有下一个节点
	if node.Next != nil {
		list.Head = node.Next
		list.Head.Prev = nil
	} else { //如果没有下一个节点, 说明只有一个节点
		list.Head, list.Tail = nil, nil
	}
	list.Size--
	return node
}

// RemoveTail 删除尾部节点
//实现思路
//判断尾部节点是否为空
//取出尾部节点
//判断尾部节点的上一个节点是否存在
//
//不存在:则说明只有一个节点, 删除本身,无须移动指针位置
//如果存在上一个节点,则尾部的上一个节点即是尾部节点.
func (list *DoubleList) RemoveTail() *DoubleNode {
	//判断尾部节点是否为空
	if list.Tail == nil {
		return nil
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	//取出尾部节点
	node := list.Tail
	//判断尾部节点的上一个是否存在
	if node.Prev != nil {
		list.Tail = node.Prev
		list.Tail.Next = nil
	}
	list.Size--
	return node
}

// Remove 删除任意元素
//实现思路
//判断是否是头部节点
//判断是否是尾部节点
//否则为中间节点,需要移动上下节点的指针位置.并实现元素删除
//
//将上一个节点的下一节点指针指向下节点
//将下一个节点的上一节点指针指向上节点
func (list *DoubleList) Remove(node *DoubleNode) *DoubleNode {
	//判断是否是头部节点
	if node == list.Head {
		return list.RemoveHead()
	}
	//判断是否是尾部节点
	if node == list.Tail {
		return list.RemoveTail()
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	//节点为中间节点
	//则需要:
	//将上一个节点的下一节点指针指向下节点
	//将下一个节点的上一节点指针指向上节点
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	list.Size--
	return node
}
