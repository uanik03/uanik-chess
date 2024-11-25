package models

import "time"

type Game struct {
    ID            uint      `gorm:"primaryKey"`
    PlayerWhiteID uint      `gorm:"not null"`
    PlayerBlackID uint      `gorm:"not null"`
    PlayerWhite   User      `gorm:"foreignKey:PlayerWhiteID;constraint:OnDelete:CASCADE"`
    PlayerBlack   User      `gorm:"foreignKey:PlayerBlackID;constraint:OnDelete:CASCADE"`
    Status        string    `gorm:"type:varchar(20);default:'in_progress'"` // in_progress, completed, abandoned
    Result        string    `gorm:"type:varchar(20)"` // white, black, draw, undecided
    CreatedAt     time.Time `gorm:"autoCreateTime"`
    UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
