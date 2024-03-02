package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name  string
		head  *ListNode
		left  int
		right int
		want  *ListNode
	}{
		{
			"test1",
			&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
			2,
			4,
			&ListNode{Val: 10, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: nil}}}}},
		},
		{
			"test2",
			&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
			1,
			5,
			&ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}},
		},
		{
			"test3",
			&ListNode{Val: 1, Next: nil},
			1,
			1,
			&ListNode{Val: 1, Next: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.head, tt.left, tt.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", linkedListToString(got), linkedListToString(tt.want))
			}
		})
	}
}

func linkedListToString(head *ListNode) string {
	str := ""
	delimiter := "->"
	for head != nil {
		valStr := strconv.Itoa(head.Val)
		head = head.Next
		if head == nil {
			str += valStr
		} else {
			str += valStr + delimiter
		}
	}
	return str
}
