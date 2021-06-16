package util

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
