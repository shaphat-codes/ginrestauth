package models

import (
	"gorm.io/gorm"
	"ginrestauth/database"
	
)

type Store struct {
    ID                uint   `gorm:"primaryKey"`
    OwnerID           uint
    Owner             User   `gorm:"foreignKey:OwnerID" json:"-"`
    StoreAccountType  string `gorm:"max:22500;default:'free';not null" json: "storeAccountType"`
    StoreName         string `gorm:"max:25500;uniqueIndex;not null" json: "storeName"`
    Twitter           string `gorm:"max:25555" json: "twitter"`
    Instagram         string `gorm:"max:25555" json: "instagram"`
    Media             string `gorm:"max:25500" json: "media"`
    DeliveryAreas     map[string]interface{} `gorm:"type:json" deliveryAreas`
    PhoneNumber       string `gorm:"max:255" json: "phoneNumber"`
    NationwideDelivery bool   `gorm:"default:true" json: "nationwideDelivery"`
}

func (store *Store) BeforeSave(*gorm.DB) error {
		var user User
		store.PhoneNumber = user.PhoneNumber
		return nil
}

func (store *Store) Save() (*Store, error) {
	err := database.Database.Create(&store).Error
	if err != nil {
		return &Store{}, err
	}
	return store, nil
}

func FindStoreById(id uint) (Store, error) {
	var store Store
	err := database.Database.Where("ID=?", id).Find(&store).Error
	if err != nil {
		return Store{}, err
	}
	return store, nil
}

func FindStoreByName(storeName string) (Store, error) {
	var store Store
	err := database.Database.Where("store_name=?", storeName).Find(&store).Error
	if err != nil {
		return Store{}, err
	}
	return store, nil
}

func (store *Store) Update(input *Store) (*Store, error) {
	err := database.Database.Model(&store).Updates(input).Error
	if err != nil {
		return &Store{}, err
	}	
	return store, nil
}

func (store *Store) Delete(id string) (*Store, error) {
	err := database.Database.Delete(&store, id).Error
	if err != nil {
		return &Store{}, err
	}
	return store, nil
}

// detail of an entry 
func (store *Store) Detail(storeName string) (*Store, error) {
	err := database.Database.First(&store, storeName).Error
	if err != nil {
		return &Store{}, err
	}
	return store, nil
}