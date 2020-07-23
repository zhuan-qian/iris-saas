package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strconv"
	"sync"
)

const (

	//用户token缓存	格式:iota_userToken => userId
	USERS_TOKEN_MAP = iota

	//管理员token缓存 格式:iota_adminToken => userId
	ADMINS_TOKEN_MAP

	//工人token缓存 格式:iota_wokerToken => userId
	WORKERS_TOKEN_MAP

	//用户手机短信30秒内重复发送判断缓存 格式:iota_phoneNum => userId
	USERS_SMS_SEND_REPEAT

	//es查询击穿丢失更新缓存 格式:iota_index_id
	ES_LOST_UPDATE
)

var (
	instance *redis.Client
	once     sync.Once
)

func RedisInit() *redis.Client {
	once.Do(func() {
		var (
			host     = os.Getenv("REDIS_HOST")
			port     = os.Getenv("REDIS_PORT")
			password = os.Getenv("REDIS_PASSWORD")
			db, _    = strconv.Atoi(os.Getenv("REDIS_DB"))
			address  = fmt.Sprintf("%s:%s", host, port)
		)

		instance = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       db,
		})

		pong, err := instance.Ping().Result()

		if err != nil {
			panic("Redis连接失败 原因:" + err.Error())
			return
		}

		println("Redis启动:" + pong)
	})
	return instance
}

//统一
func CacheKey(i int, s string) string {
	return fmt.Sprintf("%d_%s", i, s)
}
