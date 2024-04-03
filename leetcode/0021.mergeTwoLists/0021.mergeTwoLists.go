package mergetwolists

import (
	"github.com/xiaowuzai/leetcode-go/leetcode/types"
)

type ListNode = types.ListNode

/*
思考：有几种解题思路？
	1. 同时遍历两条链表，将较小的值追加到新链表
		时间复杂度 O(N)，空间复杂度 O(1)

*/
// 遍历
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	tmpHead := &ListNode{} // 返回 tmpHead.Next
	p := tmpHead           // 跟踪链表的尾部

	for list1 != nil || list2 != nil {
		// 有一个链表为空了，就把另一个放在末尾。跳出循环
		if list1 == nil {
			p.Next = list2
			break
		} else if list2 == nil {
			p.Next = list1
			break
		}

		// 如果都不为空，就找到最小的那个放在新链表的末尾
		if list1.Val <= list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}

		// 往后走一步，
		p = p.Next
	}
	return tmpHead.Next
}
