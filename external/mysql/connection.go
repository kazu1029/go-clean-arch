package mysql

import (
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error

	db, err = gorm.Open("mysql", "root:@tcp(db:3306)/clean_sample")
  if err != nil {
		panic(err)
	}
	db.Table("users").CreateTable(&gateway.User{})
	return db
}

func CloseConn() {
	db.Close()
}
