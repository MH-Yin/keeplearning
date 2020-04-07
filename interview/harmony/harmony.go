package harmony

// NumberSearch Have the function NumberSearch(str) take the str parameter, search for all the numbers in the string,
// add them together, then return that final number divided by the total amount of letters in the string.
func NumberSearch(str string) int {
	var sum, divisor int
	for _, v := range str {
		// search number
		if v >= 48 && v <= 57 {
			sum += int(v - 48)
			continue
		}
		// search letter
		if v >= 65 && v <= 122 {
			divisor++
		}
	}
	r := sum % divisor
	d := sum / divisor
	// half adjust
	if r*2 > divisor {
		return d + 1
	}
	return d
}

// DistinctList Have the function DistinctList(arr) take the array of numbers stored in arr
// and determine the total number of duplicate entries.
func DistinctList(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	result, sentry := 0, arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] == sentry {
			result++
		} else {
			sentry = arr[i]
		}
	}
	return result
}

// MatrixSpiral Have the function MatrixSpiral(strArr) read the array of strings stored in strArr which will represent a 2D N matrix,
// and return the elements after printing them in a clockwise, spiral order.
func MatrixSpiral(strArr []string) string {
	splitS := make([][]string, len(strArr))
	// 将待处理数据转变成易处理的格式
	for i := range strArr {
		splitS[i] = handleString(strArr[i])
	}
	return matrixSpiral(splitS)
}

func handleString(s string) []string {
	var str []string
	var tmp string
	for _, v := range s {
		if string(v) == " " {
			continue
		}
		if string(v) == "[" || string(v) == "]" || string(v) == "," {
			if tmp != "" {
				str = append(str, tmp)
				tmp = ""
			}
			continue
		}
		tmp += string(v)
	}
	return str
}

func matrixSpiral(strArr [][]string) string {
	var x, y int
	var result string
	// 四种前进方式的边界条件
	dirSides := []int{len(strArr[0]) - 1, len(strArr) - 1, 0, 0}
	dir := 0
	for dirSides[2] <= dirSides[0] && dirSides[3] <= dirSides[1] {
		result += strArr[y][x] + ","
		switch dir {
		case 0:
			if x+1 > dirSides[0] {
				dir = 1
				y++
				dirSides[3]++
			} else {
				x++
			}
		case 1:
			if y+1 > dirSides[1] {
				dir = 2
				x--
				dirSides[0]--
			} else {
				y++
			}
		case 2:
			if x-1 < dirSides[2] {
				dir = 3
				y--
				dirSides[1]--
			} else {
				x--
			}
		case 3:
			if y-1 < dirSides[3] {
				dir = 0
				x++
				dirSides[2]++
			} else {
				y--
			}
		}
	}
	// 过滤掉逗号
	return result[:len(result)-1]
}
