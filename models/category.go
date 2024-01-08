package models

import (
	"gorm.io/gorm"
	"ginrestauth/database"	
)

type Category struct {
    ID      uint `gorm:"primaryKey"`
    OwnerID uint
    StoreID uint
	StoreName   string `gorm:"max:25500" json: "storeName"`
    Owner   User  `gorm:"foreignKey:OwnerID" json:"-"`
    Store   Store `gorm:"foreignKey:StoreID" json:"-"`
    Name    string
}


func (category *Category) BeforeSave(tx *gorm.DB) (err error) {
    var store Store

    // Fetch the associated Store data based on StoreID
    if err := tx.Model(&Store{}).Where("id = ?", category.StoreID).First(&store).Error; err != nil {
        return err
    }

    // Set the StoreName in the Category model
    category.StoreName = store.StoreName

    return nil
}


func (category *Category) Save() (*Category, error) {
	err := database.Database.Create(&category).Error
	if err != nil {
		return &Category{}, err
	}
	return category, nil
}

func (category *Category) CategoriesForStore(storeName string) ([]Category, error) {
	var categories []Category
	err := database.Database.Where("store_name", storeName).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}