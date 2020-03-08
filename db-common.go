package main

import (
	"github.com/jinzhu/gorm"
	// mysql support for gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	log "github.com/sirupsen/logrus"
)

type omit *struct{}

// Context provides access to all DB operations
type Context struct {
	*gorm.DB
}
// CommonModel gives us a structure description
type CommonModel struct {
	ID uint `gorm:"primary_key" json:"id"`
}

// NewDB creates new instance of DB store
func NewDB(config *AppConfig) *Context {
	dbc := config.DB
	db, err := gorm.Open("sqlite3", dbc)

	if err != nil {
		log.Error(err)
		panic(1)
	}

	if config.Debug {
		db.LogMode(true)
	}

	db.AutoMigrate(&Country{})
	db.AutoMigrate(&Status{})
	db.AutoMigrate(&Contact{})
	db.AutoMigrate(&Activity{})
	db.AutoMigrate(&ActivityType{})
	db.AutoMigrate(&File{})
	return &Context{DB: db}
}
