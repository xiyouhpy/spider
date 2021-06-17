package service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xiyouhpy/spider/model"
	"github.com/xiyouhpy/spider/util"
)

// spiderInfoRedisPrefix 详情信息 redis 前缀
const redisSpiderInfoPrefix = "spider_info_redis_%s"
const redisExpire = 86400

// GetList 获取抓取列表信息
func GetList(ctx context.Context, req *util.GetListReq) (interface{}, error) {
	var strCond string = ""
	if req.Ctime != "" {
		if strCond == "" {
			strCond = "where spider_ctime > " + req.Ctime
		} else {
			strCond = strCond + " and spider_ctime > " + req.Ctime
		}
	}
	if req.OPtime != "" {
		if strCond == "" {
			strCond = "where op_time > " + req.OPtime
		} else {
			strCond = strCond + " and op_time > " + req.OPtime
		}
	}
	if req.Size != "" {
		if strCond == "" {
			strCond = "where spider_size > " + req.Size
		} else {
			strCond = strCond + " and spider_size > " + req.Size
		}
	}
	strSql := fmt.Sprintf(util.SpiderSqlMap["getListByCond"], strCond, req.Offset, req.Limit)
	mysqlList, err := model.Select(ctx, strSql)
	if err != nil {
		logrus.Warnf("call select err, err:%s", err.Error())
		return nil, err
	}

	return mysqlList, nil
}

//// GetList 获取抓取列表信息
//func GetList(ctx context.Context, req *util.GetListReq) (interface{}, error) {
//
//	// 1、从 redis 中获取信息，获取失败再从数据库获取
//	strKey := fmt.Sprintf(redisSpiderInfoPrefix, "service_id")
//	redisList, err := model.Get(ctx, strKey)
//	if err != nil {
//		return redisList, nil
//	}
//
//	// 2、从 mysql 中获取信息
//	strSql := getListSql(req)
//	mysqlList, err := model.Select(ctx, strSql)
//	if err != nil {
//		return mysqlList, err
//	}
//
//	// 3、从 redis 中获取信息
//	err = model.Set(ctx, strKey, "test", redisExpire)
//	if err != nil {
//		logrus.Warnf("call redis set err, err:%s", err.Error())
//	}
//
//	return mysqlList, nil
//}
