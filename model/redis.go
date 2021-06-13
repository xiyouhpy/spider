package model

import (
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
	"github.com/xiyouhpy/spider/util"
)

func getRedis(redisCluster string) {
	redisIP := util.RedisConfInfo[redisCluster].Ip
	redisPort := util.RedisConfInfo[redisCluster].Port
	conn, err := redis.Dial("tcp", redisIP+":"+redisPort)
	if err != nil {
		logrus.Warnf("Connect to redis error, err:%s", err)
		return
	}
	defer conn.Close()
}