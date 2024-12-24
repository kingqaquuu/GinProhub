package models
import (
    "gorm.io/gorm"
)
// Post 代表博客文章
type Post struct {
    gorm.Model

    UserID      uint   `json:"userId"`
    User        User   `gorm:"association_foreignkey:ID" json:"user,omitempty"`
    Title       string `gorm:"not null" json:"title"`
    Content     string `gorm:"type:text" json:"content"`
    Status      string `gorm:"default:'draft'" json:"status"` // 'draft', 'published'
    Views       uint   `gorm:"default:0" json:"views"`
    Comments    []Comment `gorm:"foreignkey:PostID"`
    Likes       []Like    `gorm:"foreignkey:PostID"`
    Categories  []Category `gorm:"many2many:post_categories;"`
    Tags        []Tag      `gorm:"many2many:post_tags;"`
}

// Comment 代表对文章的评论
type Comment struct {
    gorm.Model

    UserID      uint   `json:"userId"`
    User        User   `gorm:"association_foreignkey:ID" json:"user,omitempty"`
    PostID      uint   `json:"postId"`
    Post        Post   `gorm:"association_foreignkey:ID" json:"post,omitempty"`
    Content     string `gorm:"type:text" json:"content"`
    ParentID    *uint  `json:"parentId,omitempty"` // 用于支持评论的嵌套
    Likes       []Like `gorm:"foreignkey:CommentID"`
}

// Like 代表对文章或评论的点赞
type Like struct {
    gorm.Model

    UserID      uint `json:"userId"`
    User        User `gorm:"association_foreignkey:ID" json:"user,omitempty"`
    PostID      *uint `json:"postId,omitempty"`
    Post        Post  `gorm:"association_foreignkey:ID" json:"post,omitempty"`
    CommentID   *uint `json:"commentId,omitempty"`
    Comment     Comment `gorm:"association_foreignkey:ID" json:"comment,omitempty"`
}

// Category 代表文章的分类
type Category struct {
    gorm.Model
    Name        string    `gorm:"unique_index;not null" json:"name"`
    Posts       []Post    `gorm:"many2many:post_categories;"`
}

// Tag 代表文章的标签
type Tag struct {
    gorm.Model
    Name        string    `gorm:"unique_index;not null" json:"name"`
    Posts       []Post    `gorm:"many2many:post_tags;"`
}

// AfterCreate 钩子函数在创建文章后更新用户的文章列表
func (post *Post) AfterCreate(tx *gorm.DB) error {
    return tx.Model(&User{}).Where("id = ?", post.UserID).Update("Posts", gorm.Expr("array_append(posts, ?)", post.ID)).Error
}

// AfterDelete 钩子函数在删除文章后更新用户的文章列表
func (post *Post) AfterDelete(tx *gorm.DB) error {
    return tx.Model(&User{}).Where("id = ?", post.UserID).Update("Posts", gorm.Expr("array_remove(posts, ?)", post.ID)).Error
}

// BeforeDelete 钩子函数在删除评论前删除所有相关的点赞
func (comment *Comment) BeforeDelete(tx *gorm.DB) error {
    return tx.Where("comment_id = ?", comment.ID).Delete(Like{}).Error
}