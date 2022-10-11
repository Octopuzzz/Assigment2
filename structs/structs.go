package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	OrderedAt  time.Time
	Items       Item `gorm:"foreignKey:OrderID;references:ID"`
}

type Item struct {
	gorm.Model
	ItemCode    int64  `json:"itemcode"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
	OrderID     uint
}
