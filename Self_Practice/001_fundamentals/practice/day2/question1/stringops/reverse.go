package stringops

func reverseString(s string) string {
	if len(s) == 0 {
		return s
	}
	return reverseString(s[1:]) + string(s[0])
}
