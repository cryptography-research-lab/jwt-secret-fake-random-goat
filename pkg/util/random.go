package util

import (
	"math/rand"
	"time"
)

// Random 随机生成长度为n的小写字母
func Random(n int) string {

	// 这里为了方便爆破，设置为了13位的毫秒，爆破空间小一些，毕竟只是个靶场爆太久ctfer会骂娘...
	rand.Seed(time.Now().UnixMilli())

	chars := make([]rune, n)
	for i := range chars {
		chars[i] = rune(rand.Int()%26) + 'a'
	}
	return string(chars)
}
