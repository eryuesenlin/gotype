package gotype

// 定义链表
type LNode struct{
	Data interface{}
	Next *LNode
}

func NewNode() *LNode{
	return &LNode{
		Data: nil,
		Next: &LNode{},
	}
}