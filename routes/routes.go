package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rapando/gin-api/controllers"
)

// Router fn
func Router() *gin.Engine {
	router := gin.Default()
	grp1 := router.Group("/user-api")
	{
		grp1.GET("users", controllers.GetUsers)
		grp1.POST("user", controllers.CreateUser)
		// grp1.GET("user/:id", controllers.GetUserByID)
		// grp1.PUT("user/:id", controllers.UpdateUser)
		// grp1.DELETE("user/:id", controllers.DeleteUser)
	}
	return router
}
