package dbcontext

import (
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func New(db *gorm.DB) *DB {
	return &DB{db}
}

func (db *DB) DB() *gorm.DB {
	return db.db
}
