package bbsrouter

import (
	bbsapi "github.com/flipped-aurora/gin-vue-admin/server/api/v1/bbs"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct{}

func (a *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	//articleRouter := Router.Group("article").Use(middleware.OperationRecord())
	//{
	//	customerRouter.POST("customer", exaCustomerApi.CreateExaCustomer)   // 创建客户
	//	customerRouter.PUT("customer", exaCustomerApi.UpdateExaCustomer)    // 更新客户
	//	customerRouter.DELETE("customer", exaCustomerApi.DeleteExaCustomer) // 删除客户
	//}

	articleRouterWithoutRecord := Router.Group("article")
	articleApi := new(bbsapi.ArticleApi)
	{
		articleRouterWithoutRecord.GET("article", articleApi.GetArticle) // 获取单一客户信息
		//customerRouterWithoutRecord.GET("customerList", exaCustomerApi.GetExaCustomerList) // 获取客户列表
	}
}
