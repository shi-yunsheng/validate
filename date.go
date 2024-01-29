package validate

import "time"

// 是否日期时间格式
func IsDatetime(str string) bool {
	layouts := []string{"2006/01/02 15:04:05", "2006-01-02 15:04:05", "2006.01.02 15:04:05"}
	for _, layout := range layouts {
		_, err := time.Parse(layout, str)
		if err == nil {
			return true
		}
	}
	return false
}

// 是否日期格式
func IsDate(str string) bool {
	return IsDatetime(str + " 15:04:05")
}
