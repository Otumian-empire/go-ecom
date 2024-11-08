package admin

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup, db *sql.DB) *gin.RouterGroup {
	controller := NewController(db)

	router.POST("/", controller.CreatedAdmin())
	router.POST("/login", controller.Login())
	router.PATCH("/", controller.UpdateProfile())
	// router.POST("/", controller.UpdateAdminRole())
	// router.POST("/", controller.ForgetPassword())
	// router.POST("/", controller.ResetPassword())
	// router.POST("/", controller.BlockUser())
	// router.POST("/", controller.BlockAdmin())

	return router
}
