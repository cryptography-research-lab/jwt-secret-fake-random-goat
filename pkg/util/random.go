package util

import (
	"math/rand"
	"time"
)

// Random 随机生成长度为n的小写字母
func Random(n int) string {

	// 这里为了方便爆破，设置为了13位的毫秒，爆破空间小一些，毕竟只是个靶场爆太久ctfer会骂娘...
	// 实际环境中设置为 time.Now().UnixNano() 的情况可能更多一些，当然爆破速度会慢上很多倍，
	// 在实际的攻击时可以拿一个月甚至一年的时间来破解一个关键的secret，但这毕竟只是一个靶场耗时不宜过长，所以最好要让ctfer能够在几分钟内就能够算出seed
	rand.Seed(time.Now().UnixMilli())

	chars := make([]rune, n)
	for i := range chars {
		chars[i] = rune(rand.Int()%26) + 'a'
	}
	return string(chars)
}
