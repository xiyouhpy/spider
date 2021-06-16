package service

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/xiyouhpy/spider/base"
	"github.com/xiyouhpy/spider/model"
)

// spiderInfoRedisPrefix 详情信息 redis 前缀
const redisSpiderInfoPrefix = "spider_info_redis_%s"
const redisExpire = 86400

// GetList 获取抓取列表信息
func GetList(ctx context.Context, req map[string]interface{}) (interface{}, error) {
	if req["service_id"] == "" {
		logrus.Warnf("params err, req:%+v", req)
		return nil, base.ErrParamsError
	}

	// 1、从 redis 中获取信息，获取失败再从数据库获取
	strKey := fmt.Sprintf(redisSpiderInfoPrefix, req["service_id"])
	redisList, err := model.Get(ctx, strKey)
	if err != nil {
		return redisList, nil
	}

	// 2、从 mysql 中获取信息
	strSql := "select id, spider_md5, spider_title, spider_abstract, spider_url, spider_path, spider_size, spider_ctime, op_time from spider_list"
	mysqlList, err := model.Select(ctx, strSql)
	if err != nil {
		return mysqlList, err
	}

	// 3、从 redis 中获取信息
	err = model.Set(ctx, strKey, "test", redisExpire)
	if err != nil {
		logrus.Warnf("call redis set err, err:%s", err.Error())
	}

	return mysqlList, nil
}
