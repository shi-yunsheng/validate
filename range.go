package validate

import (
	"strconv"
	"strings"
)

// 验证值是否在范围中
func isValidRange(value string, r string) bool {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic("值不是一个有效数！")
	}
	if (strings.Contains(r, "[") && r[0] == '[') && (strings.Contains(r, "]") && r[len(r)-1] == ']') {

		rs := strings.Split((r[1 : len(r)-1]), ",")

		if len(rs) != 2 {
			panic("区间写法错误！")
		}

		max, e := strconv.ParseFloat(rs[1], 64)
		if e != nil && len(rs[1]) != 0 {
			panic("最大值不是一个有效数！")
		}
		min, e := strconv.ParseFloat(rs[0], 64)
		if e != nil && len(rs[0]) != 0 {
			panic("最小值不是一个有效数！")
		}

		if len(rs[1]) != 0 && val > max || len(rs[0]) != 0 && val < min {
			return false
		}
	} else {
		panic("不是一个有效区间！")
	}

	return true
}
