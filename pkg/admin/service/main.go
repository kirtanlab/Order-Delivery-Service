package service

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/admin"
	"github.com/shariarfaisal/order-ms/pkg/middleware"
	"gorm.io/gorm"
)

var db *gorm.DB

func Migration(db *gorm.DB) {
	if err := db.AutoMigrate(&admin.Admin{}, &admin.Role{}); err != nil {
	log.Printf("failed to migrate admin models: %v", err)
}

}

func Init(database *gorm.DB, r *gin.Engine) {
	db = database
	Migration(db)

	as := NewUserService(db) // admin services
	adminGroup := r.Group("/admin")
	{
		adminGroup.POST("/create", middleware.AdminAuth, as.createAdmin)
		adminGroup.POST("/login", as.loginAdmin)
		adminGroup.GET("/me", middleware.AdminAuth, as.getProfile)
	}

	rs := NewRoleService(db) // role services
	roleGroup := r.Group("/role")
	{
		roleGroup.POST("/create", rs.createRole)
	}

}
