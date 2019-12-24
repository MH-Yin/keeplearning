package algorithm

import "keeplearning/structure"

func isListPalindrome(list structure.List) bool {
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
