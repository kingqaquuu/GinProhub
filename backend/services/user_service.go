package services

import (
	"errors"
	"gin-prohub/models"
	"gin-prohub/utils/jwt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckUserExists(username, email string, db *gorm.DB) bool {

	result := db.Where("username = ? OR email = ?", username, email).First(&models.User{})
	if result.Error == nil {
		return true
	}

	// 如果错误是记录未找到，则用户不存在
	if result.Error == gorm.ErrRecordNotFound {
		return false
	}
	return false
}

func RegisterUser(c *gin.Context, db *gorm.DB) error {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return err
	}
	// 业务逻辑验证
    if user.Username == "" {
        return errors.New("用户名不能为空")
    }
    if user.Email == "" {
        return errors.New("邮箱不能为空")
    }
    if user.Password == "" {
        return errors.New("密码不能为空")
    }
    if len(user.Password) < 6 {
        return errors.New("密码长度不能小于6位")
    }

    // 检查用户是否存在
    if exists := CheckUserExists(user.Username, user.Email, db); exists {
        return errors.New("用户名或邮箱已存在")
    }
	// 创建用户（BeforeCreate 会自动处理密码加密）
    if err := db.Create(&user).Error; err != nil {
        return err
    }

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User registered successfully",
		"data":    user,
	})
	return nil
}

func LoginUser(c *gin.Context, db *gorm.DB) (gin.H,error) {
	//验证前端数据
	var loginData struct{
		Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
	}

	if err:=c.ShouldBindJSON(&loginData);err!=nil{
		return nil,err
	}
	//查询用户
    var user models.User
	if err:=db.Where("username = ?",loginData.Username).First(&user).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			return nil,errors.New("用户不存在")
		}
	}
	//验证密码

    if err := user.CheckPassword(loginData.Password); err != nil {
        return nil,errors.New("密码错误")
    }
    
	//生成token
	token,err:=jwt.GenerateToken(user.ID,user.Username)
	if err!=nil{
		return nil,err
	}
	//更新最后登录时间
	user.LastLogin=time.Now()
	if err:=db.Save(&user).Error;err!=nil{
		return nil,err
	}
	//返回登录成功响应
	return gin.H{
        "status": "success",
        "token": token,
        "user": gin.H{
            "id": user.ID,
            "username": user.Username,
            "email": user.Email,
            "last_login": user.LastLogin,
        },
    }, nil
}