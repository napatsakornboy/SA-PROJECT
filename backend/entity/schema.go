package entity

import (
	"time"

	"gorm.io/gorm"
)

type DOCTOR struct {
	gorm.Model
	Name     string
	Title    string
	Password string
	BASKETS  []BASKET `gorm:"foreignKey:DOCTOR_ID"`
}

type BASKET struct {
	gorm.Model
	Amount   int
	Add_time time.Time

	DOCTOR_ID *uint
	DOCTOR    DOCTOR `gorm:"references:id"`

	WHERE_ID *uint
	WHERE    WHERE `gorm:"references:id"`

	MEDICINE_ID *uint
	MEDICINE    MEDICINE `gorm:"references:id"`

	Symtomp_ID *uint
	Symtomp    Symtomp `gorm:"references:id"`

	//รอแอดของพืชที่เป็น FK heck_ID
}

type WHERE struct {
	gorm.Model
	Name    string   `gorm:"uniqueIndex"`
	BASKETS []BASKET `gorm:"foreignKey:WHERE_ID"`
}

type MEDICINE struct {
	gorm.Model
	NameTH  string
	Name    string
	How     string
	So      string
	Unit    string
	BASKETS []BASKET `gorm:"foreignKey:MEDICINE_ID"`
}

type Symtomp struct {
	gorm.Model
	Check_Date  time.Time
	Temperature int
	Pressure    int
	Heart_rate   int
	Comment     string
	Medicine    string
	Mapb	     string
	Check		string
	LevelID    string
	BASKETS     []BASKET `gorm:"foreignKey:Symtomp_ID"`
}
