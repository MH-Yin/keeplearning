package changting

import (
	"math/big"
)

// RSA 辗转相处法算 d
func RSA(e, c int) int {
	var phiN = 23332 * 10006
	// 递归(辗转相除).
	d := calD(e, phiN)
	// 大数运算.
	var result big.Int
	result.Exp(big.NewInt(int64(c)), big.NewInt(int64(d)), big.NewInt(int64(23333 * 10007)))
	return int(result.Int64())
}

func calD(a, b int) int {
	if a == 1 {
		return 1
	}
	tmpA, tmpB := a, b
	if a > b {
		a = a % b
	} else {
		b = b % a
	}
	return (tmpB*calK(a, b) + 1) / tmpA
}

func calK(a, b int) int {
	if b == 1 {
		return a - 1
	}
	tmpA, tmpB := a, b
	if a > b {
		a = a % b
	} else {
		b = b % a
	}
	return (tmpA*calD(a, b) - 1) / tmpB
}
