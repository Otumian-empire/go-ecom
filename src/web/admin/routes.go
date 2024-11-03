package admin

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup, db *sql.DB) *gin.RouterGroup {
	controller := NewController(db)

	router.POST("/login", controller.Login())
	router.POST("/create-admin", controller.CreatedAdmin())
	// router.POST("/", controller.UpdateProfile())
	// router.POST("/", controller.UpdateAdminRole())
	// router.POST("/", controller.ForgetPassword())
	// router.POST("/", controller.ResetPassword())
	// router.POST("/", controller.BlockUser())
	// router.POST("/", controller.BlockAdmin())

	return router
}
