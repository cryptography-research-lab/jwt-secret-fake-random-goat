package util

import "math/rand"

// Random 随机生成长度为n的小写字母
func Random(n int) string {
	chars := make([]rune, n)
	for i := range chars {
		chars[i] = rune(rand.Int()%26) + 'a'
	}
	return string(chars)
}
