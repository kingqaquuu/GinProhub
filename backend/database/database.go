package database

import (
	"gorm.io/driver/mysql"
	"log"
	"gin-prohub/models"
	"gorm.io/gorm"
	
)
func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Fatal("数据库连接失败: %v", err)
	}
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	return db,nil
}