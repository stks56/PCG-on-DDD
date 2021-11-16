package infrastructure

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func GetDB() *gorm.DB {
	return db
}

func InitDB() {
	dsn := "pcg-on-ddd:pcg-on-ddd@tcp(127.0.0.1:3306)/pcg-on-ddd?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed init DB: ", err)
	}
}
