package service

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/xiyouhpy/spider/base"

	"github.com/xiyouhpy/spider/model"
)

// GetList 获取抓取列表信息
func GetList(ctx context.Context, req map[string]interface{}) (interface{}, error) {
	if req["service_id"] == "" {
		logrus.Warnf("params err, req:%+v", req)
		return nil, base.ErrParamsError
	}
	strKey := fmt.Sprintf("%s", req["service_id"])

	// 1、判断 key 是否存在
	exists, err := model.Exists(ctx, strKey)
	if err != nil || exists == 0 {
		return nil, err
	}

	// 2、从 redis 中获取信息
	err = model.Set(ctx, strKey, "test", 0)
	if err != nil {
		return nil, err
	}

	// 3、从 redis 中获取信息
	getList, err := model.Get(ctx, strKey)
	if err != nil {
		return nil, err
	}

	return getList, nil
}
