package redis

import (
	"context"
	redis_connector "github.com/go-redis/redis/v8"
	"github.com/iamwm/connector/exception"
	"github.com/iamwm/connector/store"
	"github.com/iamwm/connector/util/conf"
	"log"
	"time"
)

func InitRedisHandler(confPath string) bool {
	var globalStore = store.STORE
	log.Println("START INIT REDIS HANDLER")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	redisConfig, err := conf.InitRedisConfig(confPath)
	if err != nil {
		log.Println(err)
		log.Println(exception.ErrRedisConfig)
		return false
	}
	globalStore.RedisConf = redisConfig
	addr := redisConfig.Host + ":" + redisConfig.Port
	client := redis_connector.NewClient(&redis_connector.Options{
		Addr:     addr,
		Password: "",
		DB:       redisConfig.DB,
	})
	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Println("redis connect failed")
		return false
	}
	//defer client.Disconnect(ctx)

	globalStore.Redis = client
	log.Println("REDIS HANDLER INIT COMPLETE")
	return true
}
