package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	username := viper.GetString("datasource.username")
	passwd := viper.GetString("datasource.passwd")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	dbname := viper.GetString("datasource.dbname")
	charset := viper.GetString("datasource.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		username,
		passwd,
		host,
		port,
		dbname,
		charset)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},PrepareStmt: true})
	if err != nil {
		panic("failed to connect database,err :" + err.Error())
	}
	return db
}

func GetDb() *gorm.DB  {
	return DB
}