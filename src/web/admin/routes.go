package admin

import (
	"database/sql"
	"otumian-empire/go-ecom/src/middleware"
	"otumian-empire/go-ecom/src/model"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup, db *sql.DB) *gin.RouterGroup {
	controller := NewController(db)
	dao := NewDao(db)

	router.POST("/login", controller.Login())

	router.Use(middleware.AuthorizeJWT[model.Admin](dao))

	router.POST("/", controller.CreatedAdmin())
	router.PATCH("/", controller.UpdateProfile())
	router.GET("/", controller.ReadProfile())
	// router.POST("/", controller.UpdateAdminRole())
	// router.POST("/", controller.ForgetPassword())
	// router.POST("/", controller.ResetPassword())
	// router.POST("/", controller.BlockUser())
	// router.POST("/", controller.BlockAdmin())

	return router
}
