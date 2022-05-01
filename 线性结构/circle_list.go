package linear

// CircleList 双链表
type CircleList struct {
	prev  *CircleList // 前驱节点
	next  *CircleList // 后驱节点
	Value interface{} // 数据
}

// Init 初始化空的循环链表;初始的时候，都是指向自己
func (c *CircleList) Init() *CircleList {
	c.prev = c
	c.next = c
	return c
}

// NewCircleList 创建N个节点的循环链表
func NewCircleList(n int) *CircleList {
	if n <= 0 {
		return nil
	}

	// 创建初始化节点
	r := new(CircleList)

	// 创建第一个节点
	p := r
	for i := 1; i < n; i++ {
		// 创建下一个节点
		p.next = &CircleList{prev: p}

		// 修改当前节点为下一个节点
		p = p.next
	}

	// 最后一个节点的下一个节点为初始化节点
	p.next = r

	// 初始化节点的前一个节点为最后一个节点
	r.prev = p
	return r
}

// GetNext 获取下一个结点
func (c *CircleList) GetNext() *CircleList {
	if c.next == nil { // 如果下个结点不存在
		return c.Init()
	}

	return c.next
}

// GetPre 获取上一个结点
func (c *CircleList) GetPre() *CircleList {
	if c.prev == nil {
		return c.Init()
	}

	return c.prev
}

// IsElemExists 某个元素是否在环形链表中存在
func (c *CircleList) IsElemExists(elem interface{}) bool {
	if c.Value == nil {
		return false
	}
	p := c
	if p.Value != elem { // 如果不相等，则继续遍历
		if p.next == c.prev {
			return false
		}

		// 移动到下一个
		p = p.next
	}

	return true
}

// Len 获取双链表的长度
func (c *CircleList) Len() int {
	n := 0
	if c != nil {
		n = 1
		for p := c.GetNext(); p != c; p = p.next {
			n++
		}
	}
	return n
}
