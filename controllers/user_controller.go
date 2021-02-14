package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rapando/gin-api/models"
)

// GetUsers controller
func GetUsers(c *gin.Context) {
	var user []models.User
	err := models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// CreateUser controller
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	if err := models.CreateUser(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
