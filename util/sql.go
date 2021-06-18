package util

// SpiderSqlMap 定义各个方法的sql的映射关系
var SpiderSqlMap = map[string]string{
	"getList": "select id, md5, title, abstract, url, path, size, ctime, atime from spider_list %s limit %s, %s",
	"addList": "insert into spider_list set md5='%s', title='%s', abstract='%s', url='%s', path='%s', size=%s, ctime=%s, atime=%s",
	"delList": "delete from spider_list where id=%s",
}
