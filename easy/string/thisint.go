package string

// 这里用的int的取值范围
const (
	MinInt = -1 << 31
	MaxInt = 1<<31 - 1

	OverFlowFlag = MaxInt / 10
)

func checkMaxFlag(n, elem int) bool {
	if n > OverFlowFlag || (n == OverFlowFlag && elem > 7) {
		return false
	}
	return true
}

func checkMinFlag(n, elem int) bool {
	if n > OverFlowFlag || (n == OverFlowFlag && elem > 8) {
		return false
	}
	return true
}

// CheckFlag 判断是否数字是否会越界（是否在int32的范围内）
// n 数字 未来将会*10
// elem 未来的个位上的数
// sign 正负数标记
func CheckFlag(n, elem, sign int) bool {
	if sign == -1 {
		return checkMinFlag(n, elem)
	}
	return checkMaxFlag(n, elem)
}

// isInt32 判断数字是否在int32范围内
func isInt32(n, elem int) bool {
	if n > 0 {
		return checkMaxFlag(n, elem)
	}
	return checkMinFlag(n, elem)
}

func isNum(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}
