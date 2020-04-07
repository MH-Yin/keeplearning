package changting

import (
	"sort"
)

// TimeToDrink 返回等待接水的时间
func TimeToDrink(waitTimeArr []int, s int) int {
	return waitForWaterTime(waitTimeArr, s) + 1
}

// waitForWaterTime 返回在有s个水龙头情况，前方等待数人，每个人接水时间为waitTimeArr时, 新排队者需要等待多少时间.
func waitForWaterTime(waitTimeArr []int, s int) int {
	// 水龙头数大于等待人数，不用等待
	if len(waitTimeArr) < s {
		return 0
	}
	// 水龙头数等于等待人数，等待时间为所有接水人等待时间最小值.
	if len(waitTimeArr) == s {
		return minNumInArray(waitTimeArr)
	}
	// 只有一个水龙头时，等待时间为所有接水人等待时间之和.
	if s == 1 {
		return sumArray(waitTimeArr)
	}

	sort.Ints(waitTimeArr[:s])
	for i := s; i < len(waitTimeArr); i++ {
		waitTimeArr[0] += waitTimeArr[i]
		// 将新等待时间按顺序插入到制定位置.
		insertNumInSortedArray(waitTimeArr, s)
	}
	return waitTimeArr[0]
}

// insertNumInSortedArray 遍历的方式插入数据.
func insertNumInSortedArray(arr []int, area int) {
	for i := 0; i < area-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		} else {
			break
		}
	}
}

// helper
// minNumInArray 返回数组中最小值.
func minNumInArray(arr []int) int {
	min := 0
	for _, v := range arr {
		if min > v {
			min = v
		}
	}
	return min
}

// sumArray 返回数组中所有元素之和.
func sumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
