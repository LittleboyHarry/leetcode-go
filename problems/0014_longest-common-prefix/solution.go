package main

func longestCommonPrefix(strs []string) string {
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	default:
		var r []byte
		for i := 0; i < len(strs[0]) && isCommonIn(strs[1:], i, strs[0][i]); i++ {
			r = append(r, strs[0][i])
		}
		return string(r)
	}
}

func isCommonIn(strs []string, i int, u uint8) bool {
	for _, str := range strs {
		if i >= len(str) || str[i] != u {
			return false
		}
	}
	return true
}
