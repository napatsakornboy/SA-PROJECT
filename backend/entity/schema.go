package entity

import (
	"time"

	"gorm.io/gorm"
)

type DOC struct {
	gorm.Model
	Name     string
	Title    string
	Password string
	BASKETS  []BASKET `gorm:foregionKey:DOC_ID`
}

type BASKET struct {
	gorm.Model
	Amount   int
	Add_time time.Time

	DOC_ID *uint
	DOC    DOC `gorm:"references:id"`

	WHERE_ID *uint
	WHERE    WHERE `gorm:"references:id"`

	MED_ID *uint
	MED    MED `gorm:"references:id"`

	//รอแอดของพืชที่เป็น FK heck_ID
}

type WHERE struct {
	gorm.Model
	Name    string   `gorm:"uniqueUndex"`
	BASKETS []BASKET `gorm:foregionKey:WHERE_ID`
}

type MED struct {
	gorm.Model
	Name    string
	How     string   `gorm:"uniqueUndex"`
	So      string   `gorm:"uniqueUndex"`
	Unit    string   `gorm:"uniqueUndex"`
	BASKETS []BASKET `gorm:foregionKey:MED_ID`
}
