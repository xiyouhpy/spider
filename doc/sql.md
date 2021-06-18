####
```
// spider 抓取数据库
CREATE DATABASE IF NOT EXISTS spider DEFAULT CHARACTER SET utf8;

// spider_list 抓取列表信息
CREATE TABLE `spider_list` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `md5` varchar(128) NOT NULL DEFAULT '' COMMENT 'md5',
    `title` varchar(128) NOT NULL DEFAULT '' COMMENT 'title',
    `abstract` varchar(64) NOT NULL DEFAULT '' COMMENT 'abstract',
    `url` varchar(128) NOT NULL DEFAULT '' COMMENT 'url',
    `path` varchar(256) NOT NULL DEFAULT '' COMMENT '本地保存路径',
    `size` int(11) NOT NULL DEFAULT 0 COMMENT '抓取大小',
    `ctime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
    `atime` int(11) NOT NULL DEFAULT 0 COMMENT '添加时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_md5` (`md5`),
    KEY `idx_ctime` (`ctime`),
    KEY `idx_size` (`size`),
    KEY `idx_atime` (`atime`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='spider_list'
```
