// transform.go
import "strconv"

// example of a custom input processing
// unescapeToHex converts backslash-encoded chars \x5c, \x27 and \x22
// to strings x5c, x27, and x22

func unescapeToHex(s string) (ret string) {
	ret = ""
	for i := 0; i < len(s); i++ {
		if s[i] == 0x5c || s[i] == 0x27 || s[i] == 0x22 {
			ret += "/x"
			sOut := strconv.FormatInt(int64(s[i]), 16)
			ret += sOut
		} else {
			ret += string(s[i])
		}
	}
	return
}

