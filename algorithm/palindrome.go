package algorithm

import "keeplearning/structure"

func isSinglyListPalindrome(list structure.SinglyList) bool {
	return true
}

func isDoublyListPalindrome(list structure.List) bool {
	if list.GetSize() <= 1 {
		return true
	}
	i, j := list.GetFirstNode(), list.GetLastNode()
	for ; i != j && i != nil; i, j = i.GetNextNode(), j.GetPreNode() {
		if i.GetValue() != j.GetValue() {
			return false
		}
	}
	return true
}
