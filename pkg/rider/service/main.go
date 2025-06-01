package service

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/rider"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	if err := db.AutoMigrate(&rider.Rider{}); err != nil {
	log.Printf("failed to migrate Rider model: %v", err)
}

}

var db *gorm.DB

func Init(database *gorm.DB, r *gin.Engine) {
	db = database
	Migration(db)
}
