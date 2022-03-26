package autoMigrate

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var db *gorm.DB

func InitDatabase(in *gorm.DB) {
	db = in
}

var migrateList = make([]interface{}, 0)

func AddToMigrateList(list ...interface{}) {
	migrateList = append(migrateList, list...)
}

func AutoMigrate() bool {
	return os.Getenv("AUTO-MIGRATE") == "TRUE"
}

func Migrate() {
	if !IsProd() && AutoMigrate() {
		err := db.Session(&gorm.Session{
			Logger: db.Logger.LogMode(logger.Warn),
		}).AutoMigrate(migrateList...)
		if err != nil {
			panic(err)
		}
	}
}
