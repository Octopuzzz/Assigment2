package config

import (
	"pratice/structs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dsn := "sql6524717:plfSJXNwhw@tcp(sql6.freemysqlhosting.net:3306)/sql6524717?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(structs.Order{}, structs.Item{})
	return db
}
