package changting

// MaximumDividedByThree  返回一个数组在不限次两两合并后最多能被3整除的个数.
func MaximumDividedByThree(arr []int) int {
	var result int
	// 分别记录余数为一和二的数目.
	remainderEqualOne, remainderEqualTwo := 0, 0
	for _, v := range arr {
		switch v % 3 {
		case 0:
			result++
		case 1:
			remainderEqualOne++
		case 2:
			remainderEqualTwo++
		}
	}

	// 先消耗所有的余数为1和2的组合，然后消化掉剩下一个.
	if remainderEqualOne >= remainderEqualTwo {
		result += remainderEqualTwo
		return result + (remainderEqualOne / 3)
	}
	result += remainderEqualOne
	return result + (remainderEqualTwo / 3)
}
