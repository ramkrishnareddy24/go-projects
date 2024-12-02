package config

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

func Connect() {
	d,err := gorm.Open("MySql","ramkrishna:Ramkrishna@24@/simplerest?chatset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d;
}

func GetDb() *gorm.DB {
	return db
}