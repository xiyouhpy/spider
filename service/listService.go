package service

import (
	"context"
	"fmt"
	"time"

	"github.com/xiyouhpy/spider/base"

	"github.com/sirupsen/logrus"
	"github.com/xiyouhpy/spider/model"
	"github.com/xiyouhpy/spider/util"
)

// spiderInfoRedisPrefix 详情信息 redis 前缀
const redisSpiderMd5Prefix = "spider_md5_%s"
const redisExpire = 86400 * 365

// GetList 获取抓取列表信息
func GetList(ctx context.Context, req *util.GetListReq) (interface{}, error) {
	var strCond string = ""
	if req.Ctime != "" {
		if strCond == "" {
			strCond = "where ctime > " + req.Ctime
		} else {
			strCond = strCond + " and ctime > " + req.Ctime
		}
	}
	if req.Atime != "" {
		if strCond == "" {
			strCond = "where atime > " + req.Atime
		} else {
			strCond = strCond + " and atime > " + req.Atime
		}
	}
	if req.Size != "" {
		if strCond == "" {
			strCond = "where size > " + req.Size
		} else {
			strCond = strCond + " and size > " + req.Size
		}
	}
	strSql := fmt.Sprintf(util.SpiderSqlMap["getList"], strCond, req.Offset, req.Limit)
	res, err := model.Select(ctx, strSql)
	if err != nil {
		logrus.Warnf("call select err, err:%s", err.Error())
		return nil, err
	}

	return res, nil
}

// AddList 添加抓取列表信息
func AddList(ctx context.Context, req *util.AddListReq) (interface{}, error) {
	if req.Title == "" || req.Url == "" || req.Path == "" || req.Ctime == "" || req.Size == "" || req.Md5 == "" {
		logrus.Warnf("AddList params err, title:%s, url:%s, path:%s, ctime:%s, size:%s, md5:%s", req.Title, req.Url, req.Path, req.Ctime, req.Size, req.Md5)
		return nil, base.ErrParamsError
	}

	// 1、查看是否已记录
	strKey := fmt.Sprintf(redisSpiderMd5Prefix, req.Md5)
	redisList, err := model.Get(ctx, strKey)
	if err == nil && redisList != 0 {
		return redisList, nil
	}

	// 2、添加数据库
	nowTime := fmt.Sprintf("%d", time.Now().Unix())
	strSql := fmt.Sprintf(util.SpiderSqlMap["addList"], req.Md5, req.Title, req.Abstract, req.Url, req.Path, req.Size, req.Ctime, nowTime)
	res, err := model.Insert(ctx, strSql)
	if err != nil {
		logrus.Warnf("call insert err, err:%s", err.Error())
		return nil, err
	}

	// 3、设置缓存
	err = model.Set(ctx, strKey, res, redisExpire)
	if err != nil {
		logrus.Warnf("call redis set err, err:%s", err.Error())
	}

	return res, nil
}

// DelList 删除抓取列表信息
func DelList(ctx context.Context, id string, md5 string) error {
	if id == "" {
		logrus.Warnf("DelList params err, id:%s", id)
		return base.ErrParamsError
	}

	// 1、删除数据库值
	strSql := fmt.Sprintf(util.SpiderSqlMap["delList"], id)
	_, err := model.Update(ctx, strSql)
	if err != nil {
		logrus.Warnf("call insert err, err:%s", err.Error())
		return err
	}

	// 2、删除 redis 值
	if md5 != "" {
		strKey := fmt.Sprintf(redisSpiderMd5Prefix, md5)
		err := model.Del(ctx, strKey)
		if err == nil {
			return nil
		}
	}

	return nil
}
