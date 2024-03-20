package middlenode

import "github.com/xiaowuzai/leetcode-go/leetcode/types"

type ListNode = types.ListNode

/*
	思考：
		方法一： 遍历链表，将节点存入切片中。找到切片的中间值，返回对应的节点即可。
			时间复杂度 O（N）,空间复杂度 O(N)
		方法二：遍历两次链表，第一次为了计数。第二次通过双指针的方式找到中间节点。
			时间复杂度 O（N）,空间复杂度 O(1)
		方法三：快慢指针
			 时间复杂度 O(N),空间复杂度 O（1）
*/

// 方法一：遍历链表并存储
func middleNode_save(head *ListNode) *ListNode {
	// 为空直接返回
	if head == nil {
		return head
	}

	// 存储节点
	var nodes []*ListNode
	var p = head
	for p != nil {
		nodes = append(nodes, p)
		p = p.Next
	}

	// 如果时偶数则取后面的元素。这里很特殊不用判断奇偶性。
	return nodes[len(nodes)/2]
}

func middleNode_pointer(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	n := 0
	p := head
	for p != nil {
		p = p.Next
		n++
	}

	n = n / 2
	p = head
	for i := 0; i < n; i++ {
		p = p.Next
	}
	return p
}

// 快慢指针
func middleNode_fast(head *ListNode) *ListNode {
	tmp := &ListNode{Next: head}
	fast, slow := tmp, tmp

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast == nil {
		return slow
	}
	return slow.Next
}
