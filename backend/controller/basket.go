package controller

import (
	"github.com/napatsakornboy/Project/entity"

	"github.com/gin-gonic/gin"

	"net/http"

)

func CreateBasket(c *gin.Context) {

	var BASKET entity.BASKET
	var WHERE entity.WHERE
	var	MEDICINE entity.MEDICINE
	var DOCTOR entity.DOCTOR
	var Symtomp entity.Symtomp

	if err := c.ShouldBindJSON(&BASKET); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", BASKET.WHERE_ID).First(&WHERE); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Where not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", BASKET.MEDICINE_ID).First(&MEDICINE); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Medicine not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", BASKET.DOCTOR_ID).First(&DOCTOR); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Doctor not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", BASKET.Symtomp_ID).First(&Symtomp); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Symtomp not found"})
		return
	}
	basket := entity.BASKET{
		Amount:       BASKET.Amount,
		Add_time:  BASKET.Add_time,
		MEDICINE: MEDICINE,
		WHERE:        WHERE,
		DOCTOR:    DOCTOR,
		Symtomp:  Symtomp,

	}
		
	if err := entity.DB().Create(&basket).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": basket})

}
	


func GetBasket(c *gin.Context) {
	var GetBasket entity.BASKET
	id := c.Param("id")	
	if err := entity.DB().Preload("DOCTOR_ID").Preload("MEDICINE_ID").Preload("WHERE_ID").Preload("Symtomp_ID").Raw("SELECT * FROM baskets WHERE id = ?", id).Find(&GetBasket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetBasket})
}

func ListBasket(c *gin.Context) {

	var basket []entity.BASKET
	if err := entity.DB().Raw("SELECT * FROM baskets").Scan(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": basket})

}
