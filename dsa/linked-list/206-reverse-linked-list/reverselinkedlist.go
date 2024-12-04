package reverselinkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// iterative
// time O(N) | space O(1)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// recursive
// time O(N) | space O(N)
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

// create new ListNode
// time O(N) | space O(N)
func reverseListCopy(head *ListNode) *ListNode {
	var cur *ListNode

	for head != nil {
		// 创建新结点，并将新结点的Next指向当前结点
		cur = &ListNode{
			Val:  head.Val,
			Next: cur,
		}
		head = head.Next
	}

	return cur
}
