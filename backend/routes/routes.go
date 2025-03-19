package routes

import (
    "traesys/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.GET("/employees", controllers.GetEmployees)
        api.POST("/employees", controllers.CreateEmployee)
        api.PUT("/employees/:id", controllers.UpdateEmployee)
        api.DELETE("/employees/:id", controllers.DeleteEmployee)
    }
}