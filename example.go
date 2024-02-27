package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tiaguinho/gosoap"
	"log"
	"net/http"
	"time"
)

func soapHandler(c *gin.Context) {
	httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	client, err := gosoap.SoapClient("http://wsgeoip.lavasoft.com/ipservice.asmx?WSDL", httpClient)
	params := gosoap.Params{
		"countryCode": c.Query("countryCode"),
		"vatNumber":   c.Query("vatNumber"),
	}
	response, err := client.Call("checkVat", params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": response})
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"valid": response, "message": "2222"})
	return
	/*valid, _ := response.Boolean("//checkVatResponse/valid")
	if valid {
		message, _ := response.ValueOf("//checkVatResponse")
		c.JSON(http.StatusOK, gin.H{"valid": true, "message": message})
	} else {
		message, _ := response.ValueOf("//checkVatResponse")
		c.JSON(http.StatusOK, gin.H{"valid": false, "message": message})
	}*/
}

func main() {
	r := gin.Default()

	r.GET("/checkVat", soapHandler)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

/*
func main() {
	r := gin.Default()
	gin.DisableConsoleColor()
	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong五2",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
*/
