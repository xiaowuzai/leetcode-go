package addtwonumbers

import "github.com/xiaowuzai/leetcode-go/leetcode/types"

type ListNode = types.ListNode

/*
分析题意：

	不要被题目绕晕了，其实就是按照两个链表顺序相加数字，超过十就进一位。不要考虑反转链表啥的

	方法 1.顺序遍历
	方法 2.递归处理
*/
// 1.顺序遍历
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Next: l1} // 将 l1 作为最终链表。当 l2 长于 l1 的时候，就需要接到 l1 的末尾
	p := head
	carry := 0 // 用来记录进位的值。比如 8 + 9 = 17 需要进 1 位，所以这个值就为 1

	for l1 != nil || l2 != nil || carry != 0 {
		// 如果 l1 和 l2 都为空，那么carry 的值需要放在 l1 的末尾
		if l1 == nil && l2 == nil {
			p.Next = &ListNode{Val: carry}
			break
		}

		if l1 != nil && l2 != nil {
			p.Next = l1
			data := l1.Val + l2.Val + carry
			carry = data / 10
			data = data % 10

			p.Next.Val = data
			l1 = l1.Next
			l2 = l2.Next
			p = p.Next
			continue
		}

		// 如果 l1 为空，那么就从 l2 里取值
		if l1 == nil && l2 != nil {
			p.Next = l2
			data := l2.Val + carry

			carry = data / 10
			data = data % 10
			p.Next.Val = data

			// 如果为零，不需要进位，直接退出循环
			if carry == 0 {
				break
			}
			l2 = l2.Next
			p = p.Next
			continue
		}

		if l1 != nil && l2 == nil {
			p.Next = l1
			data := l1.Val + carry
			carry = data / 10
			data = data % 10
			p.Next.Val = data

			// 如果为零，不需要进位，直接退出循环
			if carry == 0 {
				break
			}
			l1 = l1.Next
			p = p.Next
			continue
		}

	}

	return head.Next
}

func addTwoNumbers_dg(l1 *ListNode, l2 *ListNode) *ListNode {
	return dg(l1, l2, 0)
}

// 两个链表当前要处理的节点，返回处理过的节点
func dg(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	// 递归终止条件
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	// 处理当前节点
	if l1 != nil {
		carry += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		carry += l2.Val
		l2 = l2.Next
	}

	p := &ListNode{Val: carry % 10}
	p.Next = dg(l1, l2, carry/10)

	return p
}
