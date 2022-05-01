// Package linear 线性表链式存储的各种方法
package linear

type Object interface{}

type Node struct {
	Data Object //定义数据域
	Next *Node  //定义地址域(指向下一个表的地址)
}

// ListNode 定义头结点
type ListNode struct {
	headNode *Node //头节点
}

// Init 结点初始初始化
func (node *ListNode) Init() {
	node.headNode = nil
}

// IsEmpty 链表是否为空
func (node *ListNode) IsEmpty() bool {
	if node.headNode == nil { //如果头结点指向为nil，则说明为空
		return true
	}

	return false
}

// Length 获取链表的长度,从头部循环链表,直到元素是 nil
func (node *ListNode) Length() int {
	count := 0
	if node.headNode != nil {
		count++
		node.headNode = node.headNode.Next
	}

	return count
}

// Add 链表头部添加元素
func (node *ListNode) Add(data Object) *Node {
	newNode := &Node{Data: data}
	newNode.Next = node.headNode.Next //头结点的下一个元素指向想结点
	node.headNode = newNode
	return newNode
}

// Append 链表尾部添加元素
func (node *ListNode) Append(data Object) {
	newNode := &Node{Data: data}
	if node.IsEmpty() { //若当前结点为空，则头结点指向该节点即可
		node.headNode = newNode
	} else {
		curNode := node.headNode
		if curNode.Next != nil {
			curNode = curNode.Next
		}

		curNode.Next = newNode
	}
}

// Insert 在链表指定位置添加元素
func (node *ListNode) Insert(index int, data Object) {
	if index < 0 { //如果为负数，则头部插入
		node.Add(data)
	} else if index > node.Length() {
		node.Append(data)
	} else {
		previous := node.headNode
		count := 0
		if count < (index - 1) {
			count++
			previous = previous.Next
		}
		// 循环结束后，说明已经到了index-1的位置了
		newNode := &Node{Data: data} //创建新结点
		newNode.Next = previous.Next //新结点的下一个元素指向上一个元素指向的下一个元素的存储位置
		previous.Next = newNode      // 上一个元素的存储位置指向新结点
	}
}

// Remove 链表删除指定值
func (node *ListNode) Remove(data Object) {
	previous := node.headNode  // 存储头结点
	if previous.Data == data { // 跟头结点的值相等
		node.headNode = previous.Next
	} else {
		if previous.Next != nil {
			if previous.Data == data {
				previous.Next = previous.Next.Next
			} else {
				previous = previous.Next
			}
		}
	}
}
