package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kazu1029/go-clean-arch/adapter/gateway"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error

	db, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/clean_sample")
	if err != nil {
		panic(err)
	}
	db.Table("users").CreateTable(&gateway.User{})
	return db
}

func CloseConn() {
	db.Close()
}
