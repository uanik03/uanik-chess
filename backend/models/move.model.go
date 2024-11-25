package models

import "time"

type Move struct {
    ID         uint      `gorm:"primaryKey"`
    GameID     uint      `gorm:"not null"`
    Game       Game      `gorm:"foreignKey:GameID;constraint:OnDelete:CASCADE"`
    MoveNumber int       `gorm:"not null"`
    Notation   string    `gorm:"type:varchar(10);not null"` // e.g., e4, Nf3, O-O
    PlayedByID uint      `gorm:"not null"`
    PlayedBy   User      `gorm:"foreignKey:PlayedByID;constraint:OnDelete:CASCADE"`
    PlayedAt   time.Time `gorm:"autoCreateTime"`
}
