package reverselist

import "github.com/xiaowuzai/leetcode-go/leetcode/types"

type ListNode = types.ListNode

/*
思考解题方法：
 1. 递归

 2. 迭代
*/

func reverseList(head *ListNode) *ListNode {

	// 跟踪新链表的尾部, 新链表的头部就是 last节点
	var last *ListNode = nil

	cur, pre := head, head

	for cur != nil {
		// pre 先走一步，防止旧链表丢失
		pre = cur.Next

		// 当前节点指向后面的节点
		cur.Next = last

		// 新链表的尾部移动
		last = cur
		cur = pre // 该处理下一个节点了
	}

	return last
}
