package user

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"default:'customer'"` // admin o customer
	CreatedAt time.Time
	UpdatedAt time.Time
}
