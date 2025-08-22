package database

import (
	"fmt"
	"log/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _db *gorm.DB

func init() {
	username := "blog"
	password := "123456"
	host := "localhost"
	port := 3306
	database := "blog"
	timeout := "10s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, database, timeout)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		slog.Info("database", "连接数据库失败, error="+err.Error())
	}
}
func GetDB() *gorm.DB {
	return _db
}
