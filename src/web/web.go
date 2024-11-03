package web

import (
	"database/sql"
	"otumian-empire/go-ecom/src/web/admin"

	"github.com/gin-gonic/gin"
)

func NewHandler(router *gin.RouterGroup, db *sql.DB) *gin.RouterGroup {

	adminRouter := router.Group("/admin")
	admin.Router(adminRouter, db)

	return router
}
