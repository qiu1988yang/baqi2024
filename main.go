package main

import (
	"baqi/config"
	"baqi/routes"
	"baqi/services"
	"fmt"
)

func main() {
	//初始配置
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
	r.Run(":" + cfg.PORT)
}
