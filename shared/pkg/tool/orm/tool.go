package orm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type tool struct {
	db *gorm.DB
}

func New() Tool {
	dns := fmt.Sprintf("%v:%v@tcp(%v)/%v", "henry", "aaaa8027", "127.0.0.1:3306", "game")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	return &tool{db: db.Debug()}
}

func (t *tool) DB() *gorm.DB {
	return t.db
}
