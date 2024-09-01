package config

import (
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/util"
)

// JwtSecret 随机生成一个jwt secret，目标就是把这个随机生成的secret给破解出来
var JwtSecret = []byte(util.Random(20))

func init() {
	print("JWT secret initialized: " + string(JwtSecret))
}
