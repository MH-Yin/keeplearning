package golang

import (
	"math"
	"testing"
)

func TestTestClosure(t *testing.T) {
	//fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}

func findMaxAverage(nums []int, k int) float64 {
	var result int
	for _, v := range nums[:k] {
		result += v
	}
	for i, j := 0, k; j < len(nums); i, j = i+1, j+1 {
		sub := nums[j] - nums[i]
		if sub >= 0 {
			result += sub
		}
	}
	return float64(result) / float64(k)
}

func isPalindrome(x int) bool {
	if x < 0 || x == 10 {
		return false
	}
	if x < 10 {
		return true
	}
	lens := int(math.Log10(float64(x)))
	maxDre := int(math.Pow(10, float64(lens)))
	for ; maxDre > 10; maxDre /= 10 {
		if (x/maxDre)%10 != x%10 {
			return false
		}
	}
	return true
}

// −231,  231 − 1
func reverse(x int) int {
	t := 2 << 31
	pos := 1
	if x < 0 {
		pos = -1
		x = -x
	}
	lens := int(math.Log10(float64(x)))
	maxDre := int(math.Pow(10, float64(lens)))
	var result int
	div := 10
	for {
		a, b := divmod(x, div)
		x = a
		result += b * maxDre
		maxDre = maxDre / 10
		if a < 10 {
			result += a
			if pos > 0 {
				t--
			}
			if result > t {
				return 0
			}
			return result * pos
		}
	}
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}
