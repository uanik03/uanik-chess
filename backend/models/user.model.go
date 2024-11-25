package models

import (
	"time"
	
)

type User struct {
	ID           uint      `gorm:"primaryKey"`
    Username     string    `gorm:"type:varchar(50);unique;not null"`
    Email        string    `gorm:"type:varchar(100);unique;not null"`
    Password string    `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

                                                                                                    