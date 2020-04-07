package xuke

import "fmt"

// ArrayToString 将数组转化为逗号隔开的字符串
func ArrayToString(arr []string) string {
	length := len(arr)
	switch length {
	case 0:
		return ""
	case 1:
		return arr[0]
	default:
	}
	var result string
	for k, v := range arr {
		result += v
		if k == length-2 {
			result += " and "
			break
		} else {
			result += ", "
		}
	}
	return result + arr[length-1]
}

// ArrayToStringWithLimit 将数组转化为逗号隔开的字符串, 加上条目限制
func ArrayToStringWithLimit(arr []string, limit int) string {
	length := len(arr)
	if length <= limit {
		return ArrayToString(arr)
	}
	if limit == 0 {
		return ""
	}
	var result string
	for i := 0; i < limit-1; i++ {
		result += arr[i] + ", "
	}
	result += arr[limit-1]
	return result + fmt.Sprintf(" and %d more", length-limit)
}
