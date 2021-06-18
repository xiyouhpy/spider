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

// Select 查询数据
func Select(ctx context.Context, strSql string) (interface{}, error) {
	if util.DebugSwitch {
		logrus.Printf("Select sql:%s", strSql)
	}

	// 1、sql 预处理
	conn := getMysql()
	prepare, err := conn.Prepare(strSql)
	if err != nil {
		logrus.Warnf("prepare err, err:%s", err.Error())
		return nil, err
	}
	defer prepare.Close()

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
		err = rows.Scan(&list.Id, &list.Md5, &list.Title, &list.Abstract, &list.Url, &list.Path, &list.Size, &list.Ctime, &list.Atime)
		if err != nil {
			logrus.Warnf("scan err, err:%s", err.Error())
			continue
		}
		spiderList = append(spiderList, list)
	}

	return spiderList, nil
}

// Insert 插入数据
func Insert(ctx context.Context, strSql string) (interface{}, error) {
	if util.DebugSwitch {
		logrus.Printf("sql:%s", strSql)
	}

	// 1、sql 预处理
	conn := getMysql()
	prepare, err := conn.Prepare(strSql)
	if err != nil {
		logrus.Warnf("prepare err, err:%s", err.Error())
		return nil, err
	}
	defer prepare.Close()

	// 2、sql 执行
	ret, err := conn.ExecContext(ctx, strSql)
	if err != nil {
		logrus.Warnf("exec err, err:%s", err.Error())
		return nil, err
	}

	// 3、获取插入自增id
	id, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		logrus.Warnf("get insert id err, err:%s", err.Error())
		return nil, err
	}

	return id, nil
}

// Update 插入数据
func Update(ctx context.Context, strSql string) (interface{}, error) {
	if util.DebugSwitch {
		logrus.Printf("sql:%s", strSql)
	}

	// 1、sql 预处理
	conn := getMysql()
	prepare, err := conn.Prepare(strSql)
	if err != nil {
		logrus.Warnf("prepare err, err:%s", err.Error())
		return nil, err
	}
	defer prepare.Close()

	// 2、sql 执行
	ret, err := conn.ExecContext(ctx, strSql)
	if err != nil {
		logrus.Warnf("exec err, err:%s", err.Error())
		return nil, err
	}

	// 3、获取本次操作影响的行数
	num, err := ret.RowsAffected()
	if err != nil {
		logrus.Warnf("get RowsAffected err, err:%s", err.Error())
		return nil, err
	}

	return num, nil
}

// Delete 插入数据
func Delete(ctx context.Context, strSql string) (interface{}, error) {
	if util.DebugSwitch {
		logrus.Printf("sql:%s", strSql)
	}

	// 1、sql 预处理
	conn := getMysql()
	prepare, err := conn.Prepare(strSql)
	if err != nil {
		logrus.Warnf("prepare err, err:%s", err.Error())
		return nil, err
	}
	defer prepare.Close()

	// 2、sql 执行
	ret, err := conn.ExecContext(ctx, strSql)
	if err != nil {
		logrus.Warnf("exec err, err:%s", err.Error())
		return nil, err
	}

	// 3、获取本次操作影响的行数
	num, err := ret.RowsAffected()
	if err != nil {
		logrus.Warnf("get RowsAffected err, err:%s", err.Error())
		return nil, err
	}

	return num, nil
}
