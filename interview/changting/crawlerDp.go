package changting

import (
	"sort"
)

/*
	动态规划, 这里为了简化假定每个网页的加载时各不相同.
*/

// MinTimeOutDp 返回满足条件的最小超时时间(动态规划).
func MinTimeOutDp(page, pro int, resTimes []int, relation [][]int) int {
	return maxLoadingTimeDp(page, pro, resTimes, relation) + 1
}

func maxLoadingTimeDp(page, pro int, resTimes []int, relation [][]int) int {
	// minPage为满足要求的最小请求页面数.
	minPage := page * pro / 100
	// resTimeMap 表示响应时间与页面编号的映射关系
	resTimeMap := make(map[int]int, len(resTimes))
	for k, v := range resTimes {
		resTimeMap[v] = k
	}
	time := resTimes[0]
	notConnected := 0
	// 按响应时间排序
	sort.Ints(resTimes[1:])
	// 初始化变量
	isConnectToFirstDp = make([]bool, len(resTimes))
	isConnectToFirstDp[0] = true
	q := newQueue(len(resTimes), relation[0])

	// DP
	for i := 1; i < minPage; i++ {
		if q.find(resTimeMap[resTimes[i]], relation) {
			time = resTimes[i]
		} else {
			notConnected++
		}
	}

	if page-notConnected < minPage {
		return -2
	}
	return time
}

var isConnectToFirstDp []bool
// find 广度优先遍历判断节点i是否可达
func (q *queue) find(i int, relation [][]int) bool {
	if isConnectToFirstDp[i] == true {
		return true
	}
	for q.size > 0 {
		l := q.size
		for ; l > 0; l-- {
			page := q.Dequeue()
			for _, v := range relation[page] {
				if isConnectToFirstDp[v] == true {
					continue
				}
				isConnectToFirstDp[v] = true
				q.Enqueue(v)
			}
			if page == i {
				isConnectToFirstDp[page] = true
				return true
			}
		}
	}
	return false
}
