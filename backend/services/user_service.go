package services
import (
    "gin-prohub/models"
    "gorm.io/gorm"
)
func CheckUserExists(username,email string,db *gorm.DB) bool{
	
    result:=db.Where("username = ? OR email = ?", username, email).First(&models.User{})
    if result.Error == nil {
        return true
    }

    // 如果错误是记录未找到，则用户不存在
    if result.Error == gorm.ErrRecordNotFound {
        return false
    }
    return false
}