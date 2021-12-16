package util

// InStringSlice 判断某个字符串是否在字符串切片中
func InStringSlice(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// InIntSlice 判断某个字符串是否在整型切片中
func InIntSlice(needle int, haystack []int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
