package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	WX     WxConfig     `json:"WX"`
	Redis  string       `json:"Redis"`
	MYSQL  MysqlConfig  `json:"MYSQL"`
	YOUZAN YouzanConfig `json:"YOUZAN"`
}

type WxConfig struct {
	WxAppID   string `json:"WxAppID"`
	WxSecret  string `json:"WxSecret"`
	WxToken   string `json:"WxToken"`
	WxAccount string `json:"WxAccount"`
	WxAesKey  string `json:"WxAesKey"`
}
type MysqlConfig struct {
	MYSQL_HOST string `json:"MYSQL_HOST"`
	MYSQL_PORT string `json:"MYSQL_PORT"`
	MYSQL_USER string `json:"MYSQL_USER"`
	MYSQL_PWD  string `json:"MYSQL_PWD"`
	MYSQL_DB   string `json:"MYSQL_DB"`
}
type YouzanConfig struct {
	Youzan_Client_Id     string `json:"Youzan_Client_Id"`
	Youzan_Client_Secret string `json:"Youzan_Client_Secret"`
	Youzan_Kdt_Id        int    `json:"Youzan_Kdt_Id"`
}

var SiteConfig *Config

//传入config.json文件路径
func Init(configFileFullPath string) {

	file, err := os.OpenFile(configFileFullPath, os.O_RDONLY, 0666)
	if err != nil {
		panic("config file missd! file path:" + configFileFullPath)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic("config file read error!" + err.Error())
	}
	SiteConfig = new(Config)
	err = json.Unmarshal(bytes, SiteConfig)
	if err != nil {
		panic("config file unmarshal error!" + err.Error())
	}
	return

}
