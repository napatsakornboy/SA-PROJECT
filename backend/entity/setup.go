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

	database, err := gorm.Open(sqlite.Open("Project.db"), &gorm.Config{Logger: SqlLogger{}})

	if err != nil {

		panic("failed to connect database")

	}
	// Migrate the schema

	database.AutoMigrate(
		&DOCTOR{},
		&WHERE{},
		&BASKET{},
		&MEDICINE{},
		&Symtomp{},
	)

	// database.Create(&Symtomp{Check: "2002-05-05",Temperature: 36, Pressure: 101, Heart_rate: 77 , Comment: "มีผื่นใส่ๆขึ้นตามตัว",  Mapb: "0001", LevelID: "9001", Medicine: "HALOPERIDOL5 MG.TAB"})
	// database.Create(&Symtomp{Check: "2002-06-05",Temperature: 37, Pressure: 115, Heart_rate: 98 , Comment: "ไอ มีน้ํามูก",  Mapb: "0003", LevelID: "9002", Medicine: "AMOXY + CLAVUL[ER][AMK]1gm. TAB"})
	// database.Create(&MEDICINE{Name: "HALOPERIDOL ",NameTH: "ยาระงับประสาทหูแว่ว ", How: "รับประทานครั้งละ1 เม็ด,วันละ 1 ครั้ง, หลังอาหาร, เช้า", So: "Tab", Unit: "MG."})
	// database.Create(&MEDICINE{Name: "AMOXY ", NameTH: "ยาปฏิชีวนะกลุ่มเพนิซิลลิน  ",How: "รับประทานครั้งละ1 เม็ด,วันละ 2 ครั้ง", So: "Tab", Unit: "gm."})
	// database.Create(&MEDICINE{Name: "PHARA",NameTH: "ยาพารา", How: "หลังอาหาร เช้า เย็น", So: "Tab", Unit: "MG."})
	// database.Create(&WHERE{Name: "ห้องผู้ป่วยใน"})
	// database.Create(&WHERE{Name: "ห้องผู้ป่วยนอก"})
	// database.Create(&WHERE{Name: "ห้องผู้ป่วย VIP"})
	// database.Create(&WHERE{Name: "ห้องฉุกเฉิน"})
	// database.Create(&WHERE{Name: "ไปรษณีย์"})
	// database.Create(&DOCTOR{Name: "นายสุจร สอนชัย", Title: "Senior", Password: "111111"})
	// database.Create(&DOCTOR{Name: "นายประวิตร คิดคด", Title: "Professior", Password: "22222"})
	// database.Create(&DOCTOR{Name: "นายประยุค คิดถึงจัง", Title: "Senior", Password: "33333"})
	
	db = database
}
