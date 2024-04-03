package deleteduplicates

import "github.com/xiaowuzai/leetcode-go/leetcode/types"

type ListNode = types.ListNode

/*
思考：
方法一： 循环遍历

	如果下一个节点的值与当前节点的值相同，则删除下一个节点
	时间复杂度: O(n),空间复杂度 O(1)
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p := head
	for p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
			continue
		}
		p = p.Next
	}
	return head
}
