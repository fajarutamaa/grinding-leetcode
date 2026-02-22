package compressstring

import "strconv"

func CompressString(str string) string {
	if len(str) == 0 {
		return str
	}

	compressed := ""
	count := 1

	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			count++
		} else {
			compressed += string(str[i-1]) + strconv.Itoa(count)
			count = 1
		}
	}

	compressed += string(str[len(str)-1]) + strconv.Itoa(count)

	if len(compressed) >= len(str) {
		return str
	}
	
	return compressed
}
