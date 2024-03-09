package models

type Item struct {
	ID          uint   `gorm:"primaryKey;type:bigint" json:"-"`
	Code        string `gorm:"type:varchar(10)" json:"itemCode"`
	Description string `gorm:"type:varchar(50)" json:"description"`
	Quantity    uint   `gorm:"type:bigint" json:"quantity"`
	OrderID     uint   `gorm:"type:bigint" json:"-"`
}