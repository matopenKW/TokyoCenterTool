package util

import (
	"Carfare/pkg"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
)

func GetConnection() (db *gorm.DB, err error) {

	conf, err := ini.Load(pkg.CONFIG_PATH)
	if err != nil {
		return nil, err
	}

	section := conf.Section("db")
	DBMS := section.Key("dbms").String()
	USER := section.Key("user").String()
	PASS := section.Key("pass").String()
	PROTOCOL := section.Key("protocol").String()
	DBNAME := section.Key("db_name").String()

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}
