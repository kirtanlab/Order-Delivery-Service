package service

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/market"
	"github.com/shariarfaisal/order-ms/pkg/middleware"
	"gorm.io/gorm"
)

var db *gorm.DB

func Migration(db *gorm.DB) {
	if err := db.AutoMigrate(&market.Customer{}); err != nil {
	log.Fatalf("AutoMigrate failed for Customer: %v", err)
}

if err := db.AutoMigrate(&market.CustomerAddress{}); err != nil {
	log.Fatalf("AutoMigrate failed for CustomerAddress: %v", err)
}

if err := db.AutoMigrate(&market.Section{}, &market.SectionItem{}); err != nil {
	log.Fatalf("AutoMigrate failed for Section and SectionItem: %v", err)
}

if err := db.AutoMigrate(&market.Voucher{}); err != nil {
	log.Fatalf("AutoMigrate failed for Voucher: %v", err)
}

}

func Init(database *gorm.DB, r *gin.Engine) {
	db = database
	Migration(db)
	cs := NewCustomerService(db)
	customerGroup := r.Group("/customer")
	{
		customerGroup.POST("/signup", cs.signUp)
		customerGroup.POST("/login", cs.login)
		customerGroup.GET("/me", middleware.CustomerAuth, cs.getProfile)
	}

	// storeServices := NewStoreService(db)
	// sectionGroup := r.Group("/sections")
	// {
	// 	sectionGroup.POST("/create")
	// }

	sectionServices := NewSectionService(db)
	sectionGroup := r.Group("/sections")
	{
		sectionGroup.POST("/create", sectionServices.create)
		sectionGroup.GET("/:id", sectionServices.getById)
		sectionGroup.GET("/", sectionServices.getItems)
		sectionGroup.PUT("/:id", sectionServices.update)
		sectionGroup.DELETE("/:id", sectionServices.delete)
	}

	voucherGroup := r.Group("/vouchers")
	{
		voucherGroup.POST("/create")
	}
}
