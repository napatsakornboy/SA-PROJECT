package controller

import (
	"github.com/napatsakornboy/Project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

func CreateSymtomp(c *gin.Context) {

	var symtomp entity.Symtomp

	if err := c.ShouldBindJSON(&symtomp); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&symtomp).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": symtomp})

}

func GetSymtomp(c *gin.Context) {
	var GetSymtomp entity.Symtomp
	id := c.Param("id")
	if err := entity.DB().Preload("Symptom").Raw("SELECT * FROM symtomps WHERE id = ?", id).Scan(&GetSymtomp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetSymtomp})
}

func ListSymtomp(c *gin.Context) {

	var symtomp []entity.Symtomp
	if err := entity.DB().Raw("SELECT * FROM symtomps").Scan(&symtomp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": symtomp})

}
