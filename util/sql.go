package util

// DebugSwitch 是否开启debug模式
const DebugSwitch bool = true

// SpiderList 定义表中数据结构
type SpiderList struct {
	ID       int64  `db:"id"`
	Md5      string `db:"spider_md5"`
	Title    string `db:"spider_title"`
	Abstract string `db:"spider_abstract"`
	URL      string `db:"spider_url"`
	Path     string `db:"spider_path"`
	Size     int64  `db:"spider_size"`
	Ctime    int64  `db:"spider_ctime"`
	OpTime   int64  `db:"op_time"`
}

// SpiderSqlMap 定义各个方法的sql的映射关系
var SpiderSqlMap = map[string]string{
	"getListByCond": "select id, spider_md5, spider_title, spider_abstract, spider_url, spider_path, spider_size, spider_ctime, op_time from spider_list %s limit %s, %s",
}