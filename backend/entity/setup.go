package entity

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

//ใส่เพื่อดักข้อมูลที่ปลงจากตัวที่ไม่สามารถอ่านออก ให้เป็นภาษาของ sql ที่ให้สามารถอ่านรู้เรื่องเฉยๆ

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n=============================\n", sql)
}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("Projet.db"), &gorm.Config{Logger: SqlLogger{}})

	if err != nil {

		panic("failed to connect database")

	}
	// Migrate the schema

	database.AutoMigrate(
		&DOC{},
		&WHERE{},
		&BASKET{},
		&MED{},
	)
	db = database
}
