package model

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/xiyouhpy/spider/util"
)

var (
	UserName = "root"
	PassWord = "root"
	NetWork  = "tcp"
	Host     = "localhost"
	Port     = "3306"
	Dbname   = "spider_info"
)

var db *sql.DB = nil

// getMysql 获取 mysql 对象
func getMysql() *sql.DB {
	if db != nil {
		return db
	}

	mysqlInfo := util.MysqlConfInfo["spider"]
	Dbname = "spider"
	if mysqlInfo.Host != "" && mysqlInfo.Port != "" {
		Host = mysqlInfo.Host
		Port = mysqlInfo.Port
	}
	if mysqlInfo.Username != "" && mysqlInfo.Password != "" {
		UserName = mysqlInfo.Username
		PassWord = mysqlInfo.Password
	}
	dbServer := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", UserName, PassWord, NetWork, Host, Port, Dbname)
	db, err := sql.Open("mysql", dbServer)
	if err != nil {
		logrus.Warnf("mysql conn err, err:%s", err.Error())
		return nil
	}

	if err := db.Ping(); err != nil {
		logrus.Warnf("ping mysql err, err:%s", err.Error())
		return nil
	}

	return db
}

// Select 查询单条数据示例
func Select(ctx context.Context, strSql string) (interface{}, error) {
	conn := getMysql()

	// 1、sql 预处理
	preSql, err := conn.Prepare(strSql)
	if err != nil {
		logrus.Warnf("prepare err, err:%s", err.Error())
		return nil, err
	}
	defer preSql.Close()

	// 2、sql 执行
	rows, err := conn.QueryContext(ctx, strSql)
	if err != nil {
		logrus.Warnf("query err, err:%s", err.Error())
		return nil, err
	}
	defer rows.Close()

	// 3、数据校验和整理
	spiderList := make([]*util.SpiderList, 0)
	for rows.Next() {
		list := new(util.SpiderList)
		err = rows.Scan(&list.ID, &list.Md5, &list.Title, &list.Abstract, &list.URL, &list.Path, &list.Size, &list.Ctime, &list.OpTime)
		if err != nil {
			logrus.Warnf("scan err, err:%s", err.Error())
			continue
		}
		spiderList = append(spiderList, list)
	}

	return spiderList, nil
}
