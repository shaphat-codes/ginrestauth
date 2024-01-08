package models

import (
	"gorm.io/gorm"
	"ginrestauth/database"
	
)

type Product struct {
    ID          uint `gorm:"primaryKey"`
    CategoryID  uint
    OwnerID     uint
    StoreID     uint
	StoreName   string `gorm:"max:25500" json: "storeName"`
    Category    Category `gorm:"foreignKey:CategoryID" json:"-"`
    Owner       User     `gorm:"foreignKey:OwnerID" json:"-"`
    Store       Store    `gorm:"foreignKey:StoreID" json:"-"`
    Name        string
    Description string
    Price       float64 `gorm:"type:decimal(10,2);default:null"`
    Thumbnail   string
    MoreImages  map[string]interface{} `gorm:"type:json"`
    Stock       uint `gorm:"default:1000000"`
    Quantity    uint `gorm:"default:1"`
}


func (product *Product) BeforeSave(tx *gorm.DB) error {
	var store Store

	   // Fetch the associated Store data based on StoreID
	if err := tx.Model(&Store{}).Where("id = ?", product.StoreID).First(&store).Error; err != nil {
        return err
    }

	product.StoreName = store.StoreName

	return nil
}


func (product *Product) Save() (*Product, error) {
	err := database.Database.Create(&product).Error
	if err != nil {
		return &Product{}, err
	}
	return product, nil
}

func (product *Product) ProductsForStore(storeName string) ([]Product, error) {
	var products []Product
	err := database.Database.Where("store_name", storeName).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (product *Product) ProductsForStoreAndCategory(storeName string, categoryID int) ([]Product, error) {
	var products []Product
	err := database.Database.Where("store_name = ? AND category_id = ?", storeName, categoryID).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}