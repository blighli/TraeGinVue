package database

import (
    "log"
    "traesys/models"
    "gorm.io/gorm"
    "github.com/glebarez/sqlite"  // 改用这个驱动
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("traesys.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect database:", err)
    }

    // 自动迁移数据库结构
    DB.AutoMigrate(&models.Employee{})
}