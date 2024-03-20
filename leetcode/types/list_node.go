package types

// 单链表节点定义
type ListNode struct {
	Next *ListNode
	Val  int
}

// 根据 切片转换成数组
func NewList(vals []int) *ListNode {
	tmpHead := &ListNode{}
	p := tmpHead
	for _, val := range vals {
		p.Next = &ListNode{Val: val}
		p = p.Next
	}
	return tmpHead.Next
}

// 将单链表中的节点转化为数组
func (l *ListNode) Values() []int {
	tmpHead := &ListNode{Next: l}
	p := tmpHead
	vals := make([]int, 0, 3)
	for p.Next != nil {
		vals = append(vals, p.Next.Val)
		p = p.Next
	}
	return vals
}
