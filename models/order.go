package models

import (
	"gorm.io/gorm"
	"ginrestauth/database"	
	
)

type Order struct {
	ID           uint           `gorm:"primaryKey"`
	Products     map[string]interface{} `gorm:"type:json"`
	OwnerID      uint
	Owner        User           `gorm:"foreignKey:OwnerID"`
	Status       string         `gorm:"default:'pending';max:255"`
	Total        float64        `gorm:"type:decimal(10,2);default:null"`
	FullName     string         `gorm:"max:255"`
	PhoneNumber  string         `gorm:"max:255"`
	Region       string         `gorm:"max:255"`
	Location     string         `gorm:"max:255"`
	Notes        string         `gorm:"max:255"`
}

func (o *Order) BeforeSave(*gorm.DB) error {
	if len(o.Products) > 0 {
		var productTotals []float64

		for _, productInterface := range o.Products {
			product, ok := productInterface.(map[string]interface{})
			if !ok {
				continue
			}

			price, priceOK := product["price"].(float64)
			quantity, quantityOK := product["quantity"].(float64)

			if !priceOK || !quantityOK {
				// Handle the case where the type assertion for price or quantity fails
				continue
			}

			productTotal := price * quantity
			productTotals = append(productTotals, productTotal)
		}

		o.Total = sum(productTotals)
	} else {
		o.Total = 0
	}
	return nil
}

func sum(numbers []float64) float64 {
	var result float64
	for _, num := range numbers {
		result += num
	}
	return result
}

func (order *Order) Save() (*Order, error) {
	err := database.Database.Create(&order).Error
	if err != nil {
		return &Order{}, err
	}
	return order, nil
}
