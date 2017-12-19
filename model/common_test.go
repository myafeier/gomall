package model

import (
	"config"
)

func InitDbForTest() {
	if Db == nil {
		config.Init("/Users/xiafei/workspace/huiwanjia/gomall/src/etc/config.json")
		InitDb("mysql", config.SiteConfig.MYSQL.MYSQL_HOST, config.SiteConfig.MYSQL.MYSQL_PORT, config.SiteConfig.MYSQL.MYSQL_USER, config.SiteConfig.MYSQL.MYSQL_PWD, config.SiteConfig.MYSQL.MYSQL_DB)

	}

}
