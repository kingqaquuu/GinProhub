package models

import (
	"time"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)
type User struct{
	gorm.Model
	//基本信息
	Username 	string `gorm:"unique_index;not_null;size: 50" json:"username"`
	Email 	 	string `gorm:"unique_index;not_null" json:"email"`
	Password 	string `gorm:"not_null;" json:"-"`

	//用户状态
	IsActive 	bool 	`gorm:"default:true" json:"isActive"`
	IsBanned 	bool	`gorm:"default:false" json:"isBanned"`
	BanReason 	string	`json:"banReason,omitempty"`
	LastLogin   time.Time `json:"lastLogin"`

	//用户角色和权限
	Role		string	`gorm:"default:'user'" json:"role"`
	Permissions	[]Permission `gorm:"many2many:user_permissions;"`

	//用户行为记录
	Posts       []Post `gorm:"foreignkey:UserID"`
    Comments    []Comment `gorm:"foreignkey:UserID"`
    Likes       []Like `gorm:"foreignkey:UserID"`		
}

type Permission struct {
    gorm.Model
    Name string `gorm:"unique_index;not null" json:"name"`
}
//此函数于用户创建前执行，来加密密码
func (user *User) BeforeCreate(tx *gorm.DB) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    result := tx.Model(&User{}).Where("username = ?", user.Username).Updates(map[string]interface{}{
        "Password": string(hashedPassword),
    })

    if result.Error != nil {
        return result.Error
    }
    return nil
}
//检查密码是否匹配
func (user *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    return err == nil
}
//设置密码
func (user *User) SetPassword(password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return nil
}
// Ban 封禁用户
func (user *User) Ban(reason string) {
    user.IsBanned = true
    user.BanReason = reason
}
// Unban 解封用户
func (user *User) Unban() {
    user.IsBanned = false
    user.BanReason = ""
}
// AddPermission 为用户添加权限
func (user *User) AddPermission(permission Permission) {
    user.Permissions = append(user.Permissions, permission)
}
// RemovePermission 为用户移除权限
func (user *User) RemovePermission(permission Permission) {
    for i, p := range user.Permissions {
        if p.ID == permission.ID {
            user.Permissions = append(user.Permissions[:i], user.Permissions[i+1:]...)
            break
        }
    }
}