package controller

import (
	"github.com/napatsakornboy/Project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

func CreateMedicine(c *gin.Context) {

	var medicine entity.MEDICINE

	if err := c.ShouldBindJSON(&medicine); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&medicine).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": medicine})

}

func GetMedicine(c *gin.Context) {
	var GetMedicine entity.MEDICINE
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM medicine WHERE id = ?", id).Scan(&GetMedicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetMedicine})
}

func ListMedicine(c *gin.Context) {

	var medicine []entity.MEDICINE
	if err := entity.DB().Raw("SELECT * FROM medicines").Scan(&medicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medicine})

}
