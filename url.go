package validate

import "net/url"

// 是否是URL
func IsURL(u string) bool {
	if len(u) >= 7 && !(u[:7] == "http://" || u[:8] == "https://") {
		u = "https://" + u
	}
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}
	newU, err := url.Parse(u)
	if err != nil || newU.Host == "" {
		return false
	}
	return true
}
