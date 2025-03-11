package bbsService

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
)

// ArticleService 定义结构体 , 为 BBS 数据库 的Article 表 提供 CURD 服务
type ArticleService struct {
}

// CreateArticle 创建方法 -- 如果一个方法被 BbsServiceGroup 限定死 只能让 BbsServiceGroup 的实例 来调用 CreateArticle 方法
// @author: Dream
// @function : CreateArticle
// @description : 创建文章数据
// @param : bbs.Article
// @return : error
func (bbsArticle *ArticleService) CreateArticle(article *bbs.Article) error {

	// 【1】获取数据库的链接对象
	err := global.GVA_DB.Create(article).Error
	if err != nil {
		return err
	}

	return nil
}

//@author: Dream
//@function: UpdateArticle
//@description: 更新文章信息
//@param: e bbs.Article
//@return: err error

func (bbsArticle *ArticleService) UpdateArticle(article *bbs.Article) (err error) {
	err = global.GVA_DB.Save(article).Error
	return err
}

//@author: Dream
//@function: DeleteArticle
//@description: 删除文章对象
//@param: e bbs.Article
//@return: err error

func (bbsArticle *ArticleService) DeleteArticle(article *bbs.Article) (err error) {
	err = global.GVA_DB.Delete(article).Error
	return err
}

//@author: Dream
//@function: GetArticle
//@description: 根据 id 获取文章信息
//@param: id uint
//@return: article bbs.Article, err error

func (bbsArticle *ArticleService) GetArticle(id uint) (article *bbs.Article, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&article).Error
	return
}

//@author: Dream
//@function: GetArticleInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (bbsArticle *ArticleService) GetArticleInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&bbs.Article{})

	var ArticleList []example.ExaCustomer
	err = db.Count(&total).Error
	if err != nil {
		return ArticleList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&ArticleList).Error
	}
	return ArticleList, total, err
}
