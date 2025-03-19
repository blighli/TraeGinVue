package controllers

import (
    "traesys/database"
    "traesys/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetEmployees(c *gin.Context) {
    var employees []models.Employee
    database.DB.Find(&employees)
    c.JSON(http.StatusOK, employees)
}

func CreateEmployee(c *gin.Context) {
    var employee models.Employee
    if err := c.ShouldBindJSON(&employee); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&employee)
    c.JSON(http.StatusOK, employee)
}

func UpdateEmployee(c *gin.Context) {
    id := c.Param("id")
    var employee models.Employee
    if err := database.DB.First(&employee, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
        return
    }
    
    if err := c.ShouldBindJSON(&employee); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    database.DB.Save(&employee)
    c.JSON(http.StatusOK, employee)
}

func DeleteEmployee(c *gin.Context) {
    id := c.Param("id")
    var employee models.Employee
    if err := database.DB.First(&employee, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
        return
    }
    
    database.DB.Delete(&employee)
    c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}