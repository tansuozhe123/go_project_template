package router

import (
	"go_project/api/v1/web/controller"
	"go_project/internal/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

// 路由注入
func InitRouter(r *gin.Engine) {
	r.Use(middlewares.PanicHandler()) // 全局panic错误处理
	ProductRouter := r.Group("/v1/product")
	{
		ProductRouter.GET("/one", controller.ProductCtrl.GetOneProduct)
	}

}
