package removeelements

import "github.com/xiaowuzai/leetcode-go/leetcode/types"

// 引入定义
type ListNode = types.ListNode

/*
方法 1： 循环遍历
*/
func removeElements(head *ListNode, val int) *ListNode {
	// 引入虚拟头节点,方便处理特殊情况
	tmpHead := &ListNode{Next: head}
	p := tmpHead

	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			continue
		}
		p = p.Next
	}
	return tmpHead.Next
}

/*
2.递归

	a. 明确递归的终止条件
	b. 明确递归是否有返回值？有的话是什么？
	c. 如何递归？
*/
func removeElements_2(head *ListNode, val int) *ListNode {

	// 终止条件
	if head == nil {
		return nil
	}

	// 递归处理后续节点,并返回后续节点
	head.Next = removeElements_2(head.Next, val)

	// 当前节点该如何处理？
	// 当前节点需要删除，直接跳过
	if head.Val == val {
		return head.Next
	}
	// 直接返回当前节点
	return head
}
