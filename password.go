package validate

import "regexp"

// 通过密码复杂度检测
func passPasswordComplexity(password string) bool {
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[@#$%^&+\-=\.\!]`).MatchString(password)

	if hasNumber || hasLower || hasUpper {
		return true
	}
	if (hasNumber && hasLower || hasNumber && hasUpper || hasNumber && hasLower && hasUpper) && len(password) >= 6 {
		return true
	}
	if hasNumber && hasLower && hasUpper && hasSpecial && len(password) >= 8 {
		return true
	}
	return false
}
