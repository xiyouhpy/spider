####
```
// spider 抓取数据库
CREATE DATABASE IF NOT EXISTS spider DEFAULT CHARACTER SET utf8;

// spider_list 抓取列表信息
CREATE TABLE `spider_list` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `spider_md5` varchar(128) NOT NULL DEFAULT '' COMMENT 'md5',
    `spider_title` varchar(128) NOT NULL DEFAULT '' COMMENT 'title',
    `spider_abstract` varchar(64) NOT NULL DEFAULT '' COMMENT 'abstract',
    `spider_url` varchar(128) NOT NULL DEFAULT '' COMMENT 'url',
    `spider_path` varchar(256) NOT NULL DEFAULT '' COMMENT '本地保存路径',
    `spider_size` int(11) NOT NULL DEFAULT 0 COMMENT '抓取大小',
    `spider_ctime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
    `op_time` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_spider_md5` (`spider_md5`),
    KEY `idx_spider_ctime` (`spider_ctime`),
    KEY `idx_spider_size` (`spider_size`),
    KEY `idx_op_time` (`op_time`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='spider_list'
```
