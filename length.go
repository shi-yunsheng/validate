package validate

import (
	"strconv"
	"strings"
)

// 验证长度是否正确
func isValidLength(str string, length string) bool {
	if (strings.Contains(length, "[") && length[0] == '[') && (strings.Contains(length, "]") && length[len(length)-1] == ']') {

		rs := strings.Split((length[1 : len(length)-1]), ",")

		if len(rs) != 2 {
			panic("区间写法错误！")
		}

		max, e := strconv.ParseUint(rs[1], 0, 0)
		if e != nil && len(rs[1]) != 0 {
			panic("最大值只能是一个正整数！")
		}
		min, e := strconv.ParseUint(rs[0], 0, 0)
		if e != nil && len(rs[0]) != 0 {
			panic("最小值只能是一个正整数！")
		}

		if len(rs[1]) != 0 && uint64(len(str)) > max || len(rs[0]) != 0 && uint64(len(str)) < min {
			return false
		}
	} else {
		l, e := strconv.ParseUint(length, 0, 0)
		if e != nil {
			panic("%s长度只能是一个正整数或一个正整数区间！")
		}
		if uint64(len(str)) != l {
			return false
		}
	}

	return true
}
