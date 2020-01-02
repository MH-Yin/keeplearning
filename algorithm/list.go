package algorithm

import (
	"github.com/MH-Yin/keeplearning/structure"
)

func reverseSiglyList(list *structure.SinglyList) {
	var sentry = list.FirstNode
	var replace *structure.SinglyNode
	for ; sentry != nil; {
		next := sentry.NextNode
		sentry.NextNode = replace
		replace = sentry
		sentry = next
	}
	list.FirstNode = replace
}

func isSinglyListPalindrome(list *structure.SinglyList) bool {
	if list.Size <= 1 {
		return true
	}

	var fast, slow = list.FirstNode, list.FirstNode
	var replace *structure.SinglyNode
	for ; fast != nil && fast.NextNode != nil; {
		fast = fast.NextNode.NextNode
		next := slow.NextNode
		slow.NextNode = replace
		replace = slow
		slow = next
	}

	if fast != nil {
		slow = slow.NextNode
	}

	for ; slow != nil; {
		if slow.Value != replace.Value {
			return false
		}
		slow = slow.NextNode
		replace = replace.NextNode
	}
	return true
}

func isDoublyListPalindrome(list *structure.List) bool {
	if list.Size <= 1 {
		return true
	}
	i, j := list.FirstNode, list.LastNode
	for ; i != j && i != nil; i, j = i.NextNode, j.PreNode {
		if i.Value != j.Value {
			return false
		}
	}
	return true
}
