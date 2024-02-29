package routes

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	// 创建 Gin 实例
	r := gin.Default()

	// 示例路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, Gin!"})
	})

	r.GET("/home", homeHandler)
	r.GET("/about", aboutHandler)
	r.GET("/contact", contactHandler)

	return r
	// 添加其他路由
	// ...
	return r
}

func homeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to home page",
	})
}

func aboutHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "About us",
	})
}

func contactHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Contact us",
	})
}
