package bbsapi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	bbsService "github.com/flipped-aurora/gin-vue-admin/server/service/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleApi struct{}

// GetArticle
// @Tags      GetArticle
// @Summary   根据文章ID查询文章详情信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     bbsapi.GetArticle                                                true  "客户ID"
// @Success   200   {object}  response.Response{data=data,msg=string}  "获取单一客户信息,返回包括客户详情"
// @Router    /article/get?id=1 [get]
func (a *ArticleApi) GetArticle(c *gin.Context) {
	// 定义一个Article对象
	var articleObj bbs.Article
	// 从请求中绑定查询参数到Article对象
	err := c.ShouldBindQuery(&articleObj)
	// 如果绑定失败，返回错误信息并退出函数
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 验证Article对象的ID是否合法
	err = utils.Verify(articleObj.GVA_MODEL, utils.IdVerify)
	// 如果验证失败，返回错误信息并退出函数
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 创建一个新的ArticleService对象
	bbsService := new(bbsService.ArticleService)
	// 通过ArticleService获取Article对象的数据
	data, err := bbsService.GetArticle(articleObj.ID)
	// 如果获取失败，记录错误日志并返回错误信息
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	// 如果获取成功，返回数据和成功信息
	response.OkWithDetailed(data, "获取成功", c)
}
