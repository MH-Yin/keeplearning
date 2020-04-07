package changting

/*
	题目要求是编号从1-N，这里为了简化改成编号从零开始了。
*/

// MinTimeOut 返回满足条件的最小超时时间.
func MinTimeOut(page, pro int, resTimes []int, relation [][]int) int {
	return maxLoadingTime(page, pro, resTimes, relation) + 1
}

func maxLoadingTime(page, pro int, resTimes []int, relation [][]int) int {
	// minPage为满足要求的最小请求页面数.
	minPage := page * pro / 100
	isConnectToFirst = make([]bool, len(relation))
	isConnectToFirst[0] = true
	// 获取根页面到指定页面是否可达.
	findSinglePage(0, relation)
	var timeArr []int
	// 获取所有可达页面的响应时间数组.
	for k, connected := range isConnectToFirst {
		if connected {
			timeArr = append(timeArr, resTimes[k])
		}
	}
	// 可达页面个数不满足要求
	if len(timeArr) < minPage {
		return -2
	}
	// 返回按响应时间从小到大排列的第n个元素(minPage)
	return getRankNTerms(timeArr, 0, len(timeArr)-1, minPage)
}

var isConnectToFirst []bool

// findSinglePage 标注出从i页面出发，遍历所有可达页面并将所有可达页面标记为true.
func findSinglePage(i int, relation [][]int) {
	for _, v := range relation[i] {
		if isConnectToFirst[v] == true {
			continue
		} else {
			isConnectToFirst[v] = true
			findSinglePage(v, relation)
		}
	}
}

// getRankNTerms 返回从小到大第n个元素(快排思想).
func getRankNTerms(arr []int, head, tail, n int) int {
	i, j := head, head
	pivot := arr[tail]
	for ; j < tail; j++ {
		if arr[j] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[j], arr[i] = arr[i], arr[j]
	if i == n-1 {
		return arr[i]
	}

	if i > n-1 {
		return getRankNTerms(arr, head, i-1, n)
	}
	return getRankNTerms(arr, i+1, tail, n)
}
