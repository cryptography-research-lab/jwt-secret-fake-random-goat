package config

import (
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/util"
)

// JwtSecret 随机生成一个jwt secret，目标就是把这个随机生成的secret给破解出来
// 通常情况下是只有第一次启动的时候才需要生成，然后存储在某个地方后面读取就可以了，但是这里为了不让狡猾的ctfer能够直接读取到，所以就只保留在内存中了
var JwtSecret = []byte(util.Random(20))

func init() {
	// 不许偷看，打印出来的不算
	//print("JWT secret initialized: " + string(JwtSecret))
}
