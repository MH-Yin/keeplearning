package changting

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

// Access Log
func TestCountNumOfAccess(t *testing.T) {
	testArr := []int{1564900000, 1564900005, 1564900008, 1564900007, 1564900009}
	assert.Equal(t, 5, CountNumOfAccess(testArr, 0, 9999999999))
	assert.Equal(t, 1, CountNumOfAccess(testArr, 1564900008, 1564900008))
	assert.Equal(t, 0, CountNumOfAccess(testArr, 1564900006, 1564900006))
	assert.Equal(t, 3, CountNumOfAccess(testArr, 1564900007, 1564900009))
	assert.Equal(t, 3, CountNumOfAccess(testArr, 1564900000, 1564900007))

}

// 爬虫最小超时时间
func TestMinTimeOut(t *testing.T) {
	tests := []struct {
		except   int
		page     int
		pro      int
		resT     []int
		relation [][]int
	}{
		{101, 1, 100, []int{100}, [][]int{{}}},
		{501, 5, 100, []int{100, 200, 300, 400, 500}, [][]int{{1}, {2}, {3}, {4}, {}}},
		{-1, 5, 100, []int{1, 1, 1, 1, 1}, [][]int{{1}, {2}, {}, {4}, {}}},
		{401, 5, 80, []int{100, 200, 300, 400, 500}, [][]int{{1}, {2}, {3}, {4}, {}}},
	}

	for _, test := range tests {
		assert.Equal(t, test.except, MinTimeOut(test.page, test.pro, test.resT, test.relation))
		assert.Equal(t, test.except, MinTimeOutDp(test.page, test.pro, test.resT, test.relation))
	}
}

// RSA
func TestRSA(t *testing.T) {
	assert.Equal(t, 84859100, RSA(3, 123612763))
}

// 数组运算
func TestMaximumDividedByThree(t *testing.T) {
	assert.Equal(t, 3, MaximumDividedByThree([]int{3, 1, 2, 3, 1}))
	assert.Equal(t, 3, MaximumDividedByThree([]int{1, 1, 1, 1, 1, 2, 2}))
}

// monster要喝水
func TestTimeToDrink(t *testing.T) {
	assert.Equal(t, 3, TimeToDrink([]int{1, 1, 1, 1}, 2))
	assert.Equal(t, 6, TimeToDrink([]int{1, 2, 3, 4, 5, 2}, 3))
}
