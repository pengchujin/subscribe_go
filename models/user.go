package models 

import (
	"time"
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID int `gorm:"primary_key;AUTO_INCREMENT"`
	UUID  uuid.UUID   `json:"id"`
	Email string
	UserName string
	EncryptedPassword string
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Nodes []Node `gorm:"foreignkey:NodeID"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
  scope.SetColumn("CreatedAt", time.Now())
  scope.SetColumn("UUID", uuid.NewV4().String())
  return nil
}

func (user *User) BeforeUpdate(scope *gorm.Scope) error {
  scope.SetColumn("UpdatedAt", time.Now())
  return nil
}
