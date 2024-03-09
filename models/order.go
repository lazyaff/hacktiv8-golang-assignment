package models

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey;type:bigint" json:"id"`
	OrderedAt    time.Time `gorm:"typeLdatetime" json:"orderedAt"`
	CustomerName string    `gorm:"type:varchar(50)" json:"customerName"`
	Items        []Item    `gorm:"foreignKey:OrderID" json:"items"`
}