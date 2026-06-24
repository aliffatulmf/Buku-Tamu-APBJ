package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// default connection
func NewConnection(config *gorm.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("bukutamu.db"), config)
	if err != nil {
		os.Exit(1)
	}
	return db
}

// for testing only
func NewConnectionTesting(loc string, config *gorm.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(loc), config)
	if err != nil {
		os.Exit(1)
	}
	return db
}
