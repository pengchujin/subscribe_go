package database


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pengchujin/subscribe_go/config"
	"github.com/pengchujin/subscribe_go/models"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	conf := config.Get()
	db, err := gorm.Open("postgres", conf.DSN)
	if err == nil {
		db.DB().SetMaxIdleConns(conf.MaxIdleConn)
		DB = db
		db.AutoMigrate(&models.User{}, &models.Node{})
		return db, err
	}
	return nil, err
}