package model

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/xiyouhpy/spider/util"
	//"github.com/garyburd/redigo/redis"
)

var redisConn *redis.Client
var redisOnce sync.Once

// getInstance 获取 redis 对象
func getInstance(ctx context.Context) *redis.Client {
	if redisConn != nil {
		return redisConn
	}

	redisOpt := redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DB:           0,
		DialTimeout:  1000 * time.Millisecond,
		ReadTimeout:  5000 * time.Millisecond,
		WriteTimeout: 5000 * time.Millisecond,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	}

	redisInfo := util.RedisConfInfo["redis_spider"]
	// 设置访问 host、ip
	if redisInfo.Host != "" && redisInfo.Port != "" {
		redisOpt.Addr = fmt.Sprintf("%s:%s", redisInfo.Host, redisInfo.Port)
	}
	// 设置访问 timeout
	if redisInfo.ConnTimeout > 0 && redisInfo.ReadTimeout > 0 && redisInfo.WriteTimeout > 0 {
		redisOpt.DialTimeout = time.Duration(redisInfo.ConnTimeout) * time.Millisecond
		redisOpt.ReadTimeout = time.Duration(redisInfo.ReadTimeout) * time.Millisecond
		redisOpt.WriteTimeout = time.Duration(redisInfo.WriteTimeout) * time.Millisecond
	}

	redisOnce.Do(func() {
		redisConn = redis.NewClient(&redisOpt)
	})

	return redisConn
}

// Set redis set 操作
func Set(ctx context.Context, key string, val interface{}, expire int) error {
	c := getInstance(ctx)
	err := c.SetEX(ctx, key, val, time.Duration(expire)*time.Second).Err()
	if err != nil {
		logrus.Warnf("set redis err, err:%s", err.Error())
		return err
	}

	return nil
}

// Get redis get 操作
func Get(ctx context.Context, key string) (interface{}, error) {
	c := getInstance(ctx)
	val, err := c.Get(ctx, key).Result()
	if err != nil {
		logrus.Warnf("get redis err, err:%s", err.Error())
		return nil, err
	}

	return val, nil
}

// Exists redis exists 操作
func Exists(ctx context.Context, key string) (int64, error) {
	c := getInstance(ctx)
	val, err := c.Exists(ctx, key).Result()
	if err != nil {
		logrus.Warnf("exists redis err, err:%s", err.Error())
		return val, err
	}

	return val, nil
}
