package dal

import (
	"fmt"
	"github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/util"
	// 确保另一个包先初始化，先调用随机数函数
	_ "github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/config"
)

// 用于存储用户名和密码
var userDatabase = map[string]string{}

const FlagUserName = "CC11001100"

func init() {

	// 为用户生成一个随机密码
	randomPasswd := util.Random(100)
	// 不许偷看
	//println(randomPasswd)
	userDatabase[FlagUserName] = randomPasswd

}

// Register 注册用户
func Register(username, password string) error {
	if _, ok := userDatabase[username]; ok {
		return fmt.Errorf("user %s already exists", username)
	}
	userDatabase[username] = password
	return nil
}

// FindPasswd 查找用户密码
func FindPasswd(username string) (string, bool) {
	password, ok := userDatabase[username]
	return password, ok
}

// ListUsername 列出当前存在的所有用户
func ListUsername() []string {
	usernames := make([]string, 0, len(userDatabase))
	for username := range userDatabase {
		usernames = append(usernames, username)
	}
	return usernames
}
