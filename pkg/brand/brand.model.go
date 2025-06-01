package brand

import (
	"encoding/json"
	"errors"

	"log"

	"github.com/shariarfaisal/order-ms/pkg/hub"
	"gorm.io/gorm"
)

type BrandStatus string

const (
	BrandStatusPending  BrandStatus = "pending"
	BrandStatusActive   BrandStatus = "active"
	BrandStatusInactive BrandStatus = "inactive"
)

type BrandType string

const (
	BrandTypeStore      BrandType = "store"
	BrandTypeRestaurant BrandType = "restaurant"
	BrandTypeGrocery    BrandType = "grocery"
)

type OperatingTime struct {
	From struct {
		Hour   int `json:"hour"`
		Minute int `json:"minute"`
	} `json:"from,omitempty"`
	To struct {
		Hour   int `json:"hour"`
		Minute int `json:"minute"`
	} `json:"to,omitempty"`
}

type OperatingTimes map[string]interface{}

func (p *OperatingTimes) Scan(value interface{}) error {
	source, ok := value.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion OperatingTimes failed.")
	}

	return nil
}

type Brand struct {
	ID             uint           `json:"id,omitempty" gorm:"primaryKey"`
	Name           string         `json:"name,omitempty"`
	Slug           string         `json:"slug,omitempty"`
	Type           BrandType      `json:"type,omitempty"`
	Details        string         `json:"details,omitempty"`
	Phone          string         `json:"phone,omitempty"`
	Email          string         `json:"email,omitempty"`
	EmailVerified  bool           `json:"emailVerified,omitempty"`
	Logo           string         `json:"logo,omitempty"`
	BannerImage    string         `json:"bannerImage,omitempty"`
	Rating         float32        `json:"rating,omitempty"`
	PartnerId      uint           `json:"partnerId,omitempty"`
	Partner        Partner        `json:"-" gorm:"foreignKey:PartnerId"`
	Prefix         string         `json:"prefix,omitempty"`
	Status         BrandStatus    `json:"status,omitempty"`
	IsAvailable    bool           `json:"isAvailable,omitempty"`
	AddressId      uint           `json:"addressId,omitempty" gorm:"index"`
	Address        BrandAddress   `json:"-" gorm:"foreignKey:AddressId"`
	OperatingTimes OperatingTimes `json:"operatingTimes,omitempty" gorm:"type:jsonb"`
	HubId          uint           `json:"hubId,omitempty" gorm:"index"`
	Hub            hub.Hub        `json:"-" gorm:"foreignKey:HubId"`
}

type BrandAddress struct {
	ID         uint    `json:"id,omitempty" gorm:"primaryKey"`
	Address    string  `json:"address,omitempty"`
	Area       string  `json:"area,omitempty"`
	PostalCode string  `json:"postalCode,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
	Apartment  string  `json:"apartment,omitempty"`
	Floor      string  `json:"floor,omitempty"`
	RoadNo     string  `json:"roadNo,omitempty"`
}



func Migration(db *gorm.DB) {
	if err := db.AutoMigrate(&Brand{}); err != nil {
		log.Printf("failed to migrate Brand: %v", err)
	}
	if err := db.AutoMigrate(&BrandAddress{}); err != nil {
		log.Printf("failed to migrate BrandAddress: %v", err)
	}
	if err := db.AutoMigrate(&BrandCategory{}); err != nil {
		log.Printf("failed to migrate BrandCategory: %v", err)
	}
	if err := db.AutoMigrate(&Partner{}); err != nil {
		log.Printf("failed to migrate Partner: %v", err)
	}
	if err := db.AutoMigrate(&PartnerUser{}); err != nil {
		log.Printf("failed to migrate PartnerUser: %v", err)
	}
	if err := db.AutoMigrate(&ProductCategory{}); err != nil {
		log.Printf("failed to migrate ProductCategory: %v", err)
	}
	if err := db.AutoMigrate(&Product{}, &ProductVariant{}, &PurchaseProduct{}, &ProductDiscount{}); err != nil {
		log.Printf("failed to migrate product-related models: %v", err)
	}
}

