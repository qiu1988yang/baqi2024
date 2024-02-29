package controllers

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, Gin!"})
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome to home page"})
}

func About(c *gin.Context) {
	c.JSON(200, gin.H{"message": "About us"})
}

func Contact(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Contact us"})
}
