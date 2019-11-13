package models

import (
	"time"
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
	"database/sql/driver"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type nodeType string

const (
	SS  nodeType = "SS"
	SSR nodeType = "SSR"
	V2RAY nodeType = "V2RAY"
)

func (p *nodeType) Scan(value interface{}) error {
	*p = nodeType(value.([]byte))
	return nil
}

func (p nodeType) Value() (driver.Value, error) {
	return string(p), nil
}

type Node struct {
	gorm.Model
	NodeID uint
	UUID uuid.UUID   `json:"id"`
	NodeType nodeType `sql:"type:nodeType"`
	Info postgres.Jsonb
	Serial int
	User User
	NodeCreatedAt time.Time  `json:"created_at"`
  NodeUpdatedAt time.Time `json:"updated_at"`
}

func (node *Node) BeforeCreate(scope *gorm.Scope) error {
  scope.SetColumn("NodeCreatedAt", time.Now())
  scope.SetColumn("UUID", uuid.NewV4().String())
  return nil
}

func (node *Node) BeforeUpdate(scope *gorm.Scope) error {
  scope.SetColumn("NodeUpdatedAt", time.Now())
  return nil
}

func (Node) TableName() string {
	return "node"
}