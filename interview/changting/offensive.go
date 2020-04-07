package changting

/*
	其实没有太好的思路，目前想到的就是先找到目标机直接的距离关系，然后根据先后访问顺序求出最少跳转路径
 */

// Offensive 寻找攻击路径
func Offensive(m int, targets []int, side [][2]int) {
	relation := make([][]int, m)
	for i := range side {
		relation[side[i][0]] = append(relation[side[i][0]], side[i][1])
		relation[side[i][1]] = append(relation[side[i][1]], side[i][0])
	}
	distances = make([][]int, m)

	// waysList := make([][]int, len(targets))

	for i := 0; i < m; i++ {
		q := newQueue(m, relation[i])
		distances[i] = make([]int, m)
		q.getDistance(0, relation)
	}
}

var distances [][]int

// 可以根据之前已经找到的距离优化。由于整体不是很清晰，时间也到了，就没继续进行了。。
func (q *queue) getDistance(from int, relation [][]int) {
	d := 0
	for q.size > 0 {
		d++
		l := q.size
		for ; l > 0; l-- {
			page := q.Dequeue()
			distances[from][page] = d
			for _, v := range relation[page] {
				if distances[from][page] != 0 {
					continue
				}
				distances[from][page] = d + 1
				q.Enqueue(v)
			}
		}
	}
}
