package bbs

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// Article 文章表模型结构体
type Article struct {
	// 继承全局 公共字段
	global.GVA_MODEL
	Title        string `json:"title" gorm:"not null;index;comment :标题"`
	Img          string `json:"img" gorm:"not null;comment : 封面图"`
	Description  string `json:"description" gorm:"not null;comment : 描述"`
	Content      string `json:"content" gorm:"not null;comment : 文章内容 -- MD 格式"`
	HtmlContent  string `json:"htmlContent" gorm:"not null;comment : html内容 -- MD格式"`
	CategoryId   uint   `json:"categoryId" gorm:"not null;comment : 分类ID"`
	CategoryName string `json:"categoryName" gorm:"not null;comment : 分类名称"`
	ViewCount    uint   `json:"viewCount" gorm:"not null;comment : 浏览量"`
	Comments     uint   `json:"comments" gorm:"not null;comment : 评论数"`
	CommentsOpen bool   `json:"commentsOpen" gorm:"not null;comment : 是否开启评论 0 未开启 1 已开启"`
	Status       bool   `json:"status" gorm:"not null;comment : 是否发布 0 未发布 1 已发布"`
	IsDelete     bool   `json:"isDelete" gorm:"not null;comment : 是否删除 0 未删除 1 已删除"`
	UserId       uint   `json:"userId" gorm:"not null;comment : 文章发布者ID"`
	UserName     string `json:"userName" gorm:"not null;comment : 文章发布者用户名"`
	UserAvatar   string `json:"userAvatar" gorm:"not null;comment : 文章发布者头像"`
}

// TableName 定义Article结构体的方法，返回表名
func (Article) TableName() string {
	// 返回字符串"bbs_article"作为表名
	return "bbs_article"
}
