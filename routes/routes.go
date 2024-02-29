package routes

import (
	"baqi/controllers"
	"baqi/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 创建 Gin 实例
	r := gin.Default()

	// 应用全局中间件
	r.Use(middleware.RequestLogger())

	// 示例路由
	r.GET("/", controllers.Index)

	// 定义路由分组并应用中间件
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.RequestLogger())
	{
		// 定义 "/api" 路由组的子路由
		apiGroup.GET("/home", controllers.Home)
		apiGroup.GET("/about", controllers.About)
	}
	// 单独定义路由，不在路由组中
	r.GET("/contact", controllers.Contact)
	return r
}
