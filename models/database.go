package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=db user=postgres password=postgres dbname=spy_cats port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Cat{}, &Mission{}, &Target{})

	DB = database
}
