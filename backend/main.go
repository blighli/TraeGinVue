package main

import (
    "traesys/database"
    "traesys/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // 初始化数据库
    database.InitDB()

    // 创建 Gin 实例
    r := gin.Default()

    // 配置 CORS
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    })

    // 设置路由
    routes.SetupRoutes(r)

    // 启动服务器
    r.Run(":8080")
}