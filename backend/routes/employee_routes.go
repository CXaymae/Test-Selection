package routes

import (
	"employee-management-backend/controllers"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(router *gin.Engine) {
	employee := router.Group("/employees")
	{
		employee.GET("/", controllers.GetEmployees)
		employee.GET("/:id", controllers.GetEmployeeByID)
		employee.POST("/", controllers.CreateEmployee)
		employee.PUT("/:id", controllers.UpdateEmployee)
		employee.DELETE("/:id", controllers.DeleteEmployee)
	}
}
