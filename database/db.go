package database

import (
	"log"

	"github.com/jinzhu/gorm"
)
var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}	

func Setup(){
	host:="localhost"
	port:="5432"
	dbName:="book"
	username:="postgres"
	password:="postgres"
	arg:= "host=" + host + " port="+port+" username="+username+" dbname="+dbName+" sslmode=disable password="+password
	db,err:=gorm.Open("postgres",arg)
	if err!=nil{
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	DB = db

}