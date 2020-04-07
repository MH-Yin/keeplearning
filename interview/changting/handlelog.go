package changting

import (
	"sort"
)

// CountNumOfAccess 返回指定范围内的访问次数.
func CountNumOfAccess(logWithTime []int, startTime, endTime int) int {
	sort.Ints(logWithTime)
	head, tail := 0, len(logWithTime)-1
	// 快速检查不符合区间的情况.
	if logWithTime[head] > endTime || logWithTime[tail] < startTime {
		return 0
	}
	// 获取大于等于startTime的最早时间请求的下标.
	lower := findLowerIndex(logWithTime, startTime)
	// 获取小与等于endTime的最晚时间请求的下标.
	upper := findUpperIndex(logWithTime, endTime)
	// 下标为-1的特殊情况已在快速检查时被处理，因此不与考虑.
	return upper - lower + 1
}

// findLowerIndex 查找第一个大于等于给定值的元素.
func findLowerIndex(arr []int, target int) int {
	if target <= arr[0] {
		return 0
	}
	head, tail := 0, len(arr)-1
	var mid int
	for head <= tail {
		mid = head + (tail-head)>>1
		if arr[mid] >= target {
			if mid == 0 || arr[mid-1] < target {
				return mid
			}
			tail = mid - 1
		} else {
			if arr[mid+1] > target {
				return mid + 1
			}
			head = mid + 1
		}
	}
	return -1
}

// findUpperIndex 查找最后一个小于等于给定值的元素下标.
func findUpperIndex(arr []int, target int) int {
	if target >= arr[len(arr)-1] {
		return len(arr) - 1
	}
	var mid int
	head, tail := 0, len(arr)-1
	for head <= tail {
		mid = head + (tail-head)>>1
		if arr[mid] <= target {
			if mid == len(arr)-1 || arr[mid+1] > target {
				return mid
			}
			head = mid + 1
		} else {
			if arr[mid-1] < target {
				return mid - 1
			}
			tail = mid - 1
		}
	}
	return -1
}
