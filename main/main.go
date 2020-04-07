package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	head, tail := 0, len(nums)-1
	for head <= tail {
		mid := head + (tail-head)>>1
		if nums[mid] == target {
			return mid
		}
		// 左边递增数列
		if nums[mid] > nums[head] {
			if nums[head] <= target {
				return searchInSortedArray(nums, target, head, mid-1)
			}
			head = mid + 1
		}
		// 右边递增数列
		if nums[mid] < nums[tail] {
			if nums[tail] >= target {
				return searchInSortedArray(nums, target, mid+1, tail)
			}
			tail = mid - 1
		}
	}
	return -1
}

func searchInSortedArray(nums []int, target, head, tail int) int {
	for head <= tail {
		mid := head + (tail-head)>>1
		if nums[mid] > target {
			tail = mid - 1
		} else {
			if nums[mid] == target {
				return mid
			}
			head = mid + 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	start, end := -1, -1
	lens := len(nums)
	head, tail := 0, lens-1
	for head <= tail {
		mid := head + ((tail - head) >> 1)
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				start = mid
			}

			if mid == lens-1 || nums[mid+1] != target {
				end = mid
			}

			if start == -1 {
				start = searchStartIndex(nums, target, head, mid-1)
			}

			if end == -1 {
				end = searchEndIndex(nums, target, mid+1, tail)
			}
			return []int{start, end}
		}

		if nums[mid] > target {
			tail = mid - 1
		} else {
			head = mid + 1
		}
	}
	return []int{start, end}
}

func searchStartIndex(nums []int, target, head, tail int) int {
	for head <= tail {
		mid := head + ((tail - head) >> 1)
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			tail = mid - 1
			continue
		}
		if nums[mid] > target {
			tail = mid - 1
		} else {
			head = mid + 1
		}
	}
	return -1
}

func searchEndIndex(nums []int, target, head, tail int) int {
	fmt.Println(target, head, tail)
	for head <= tail {
		mid := head + ((tail - head) >> 1)
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			head = mid + 1
			continue
		}
		if nums[mid] > target {
			tail = mid - 1
		} else {
			head = mid + 1
		}
	}
	return -1
}

func quickSort(arr []int, head, fin int) {
	if head > fin {
		return
	}
	var pivot = arr[fin]
	i, j := head, head
	for ; j <= fin; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	quickSort(arr, head, i-2)
	quickSort(arr, i, fin)
}

func reverseList(head *listNode) *listNode {
	var pre, tmp *listNode
	for head != nil {
		tmp = head.Next
		fmt.Println(tmp)
		head.Next = pre
		pre = head
		if head.Next == nil {
			return head
		}
		head = tmp
		fmt.Println(tmp)
	}
	return head
}

type listNode struct {
	Val  int
	Next *listNode
}
