package util

// DebugSwitch 是否开启debug模式
const DebugSwitch bool = true

// SpiderList spider_list 表结构定义
type SpiderList struct {
	Id       int64  `db:"id"`
	Md5      string `db:"md5"`
	Title    string `db:"title"`
	Abstract string `db:"abstract"`
	Url      string `db:"url"`
	Path     string `db:"path"`
	Size     int64  `db:"size"`
	Ctime    int64  `db:"ctime"`
	Atime    int64  `db:"atime"`
}

// GetListReq GetList 接口请求参数结构
type GetListReq struct {
	Atime  string
	Ctime  string
	Size   string
	Offset string
	Limit  string
}

// AddListReq AddList 接口请求参数结构
type AddListReq struct {
	Md5      string
	Title    string
	Abstract string
	Url      string
	Path     string
	Ctime    string
	Size     string
}
