package validate

import (
	"strings"
)

// 判断值是否存在
func existValue(data string, value string) bool {
	if (strings.Contains(data, "(") && data[0] == '(') && (strings.Contains(data, ")") && data[len(data)-1] == ')') {
		rs := strings.Split((data[1 : len(data)-1]), ",")

		for _, r := range rs {
			if r == value {
				return true
			}
		}
	} else {
		panic("不是一个有效枚举！")
	}

	return false
}
