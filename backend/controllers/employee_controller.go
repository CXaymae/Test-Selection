package controllers

import (
	"employee-management-backend/models"
	"employee-management-backend/services"
	"employee-management-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployees(c *gin.Context) {
	employees, err := services.GetAllEmployees()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error fetching employees")
		return
	}
	c.JSON(http.StatusOK, employees)
}

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	employee, err := services.GetEmployeeByID(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Employee not found")
		return
	}
	c.JSON(http.StatusOK, employee)
}

func CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}
	if err := services.AddEmployee(employee); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error creating employee")
		return
	}
	c.JSON(http.StatusCreated, employee)
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}
	if err := services.UpdateEmployee(id, employee); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error updating employee")
		return
	}
	c.JSON(http.StatusOK, employee)
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteEmployee(id); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error deleting employee")
		return
	}
	c.Status(http.StatusNoContent)
}
