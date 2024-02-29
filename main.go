package main

import (
	"baqi/config"
	"baqi/routes"
	"baqi/services"
	"fmt"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error loading config file: %s \n", err))
	}

	// 初始化数据库连接
	_, err = services.InitDB(cfg)
	if err != nil {
		panic(fmt.Errorf("Fatal error initializing database: %s \n", err))
	}

	_, err = services.InitRedis(cfg)
	if err != nil {
		panic(fmt.Errorf("Fatal error initializing redis: %s \n", err))
	}

	myLogger := services.NewLogger()
	myLogger.Info("app.log", "这是一条信息日志")
	myLogger.Warning("service.log", "这是一条警告日志")

	//路由
	r := routes.SetupRouter()
	r.Run(":8888")
}

///func soapHandler(c *gin.Context) {

/*	// 示例用法
	persons := []Person{
		{1, "Alice", 30},
		{2, "Bob", 25},
		{3, "Charlie", 35},
	}

	// 将 Name 字段作为键转换为映射
	nameMap, err := utils.ConvertArrayToMap(persons, "Name")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回 JSON
	c.JSON(http.StatusOK, gin.H{"nameMap": nameMap})
*/
/*jsonData, err := json.Marshal(nameMap)
if err != nil {
	fmt.Println("Error:", err)
	return
}*/
//c.Data(http.StatusOK, "application/json", jsonData)
//c.JSON(http.StatusOK, gin.H{"valid": false, "message": message})

//fmt.Println("Name Map:", nameMap)

/*	httpClient := &http.Client{
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
	return*/
/*valid, _ := response.Boolean("//checkVatResponse/valid")
if valid {
	message, _ := response.ValueOf("//checkVatResponse")
	c.JSON(http.StatusOK, gin.H{"valid": true, "message": message})
} else {
	message, _ := response.ValueOf("//checkVatResponse")
	c.JSON(http.StatusOK, gin.H{"valid": false, "message": message})
}*/
//}

/*func main() {
	r := gin.Default()
	r.GET("/checkVat", soapHandler)
	err := r.Run(":8888")
	if err != nil {
		log.Fatal(err)
	}
}*/

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
