package model

import (
	"time"

	"github.com/google/uuid"
	util "github.com/moluh/ginrest/util"
	"gorm.io/gorm"
)

type UserModel struct {
	ID        string `gorm:"primaryKey;"`
	Name      string `gorm:"type:varchar(255);unique;not null"`
	Surname   string `gorm:"type:varchar(255);unique;not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Active    bool   `gorm:"type:bool;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *UserModel) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.Password = util.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *UserModel) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
