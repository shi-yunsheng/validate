package validate

import "net/mail"

// 是否是邮件格式
func IsMail(str string) bool {
	_, e := mail.ParseAddress(str)
	return e == nil
}
