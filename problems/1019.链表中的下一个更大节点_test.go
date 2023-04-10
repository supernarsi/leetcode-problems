package problems

import "testing"

func TestNextLargerNodes(t *testing.T) {
	t.Log(nextLargerNodes(&ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5}}}))
	t.Log(nextLargerNodes(&ListNode{Val: 2, Next: &ListNode{Val: 7, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}))
	t.Log(nextLargerNodes(&ListNode{Val: 1, Next: &ListNode{Val: 7, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1}}}}}}}}))
}
