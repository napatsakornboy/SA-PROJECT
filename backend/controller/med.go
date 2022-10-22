package controller

import (
	"github.com/napatsakornboy/Project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

func CreateMed(c *gin.Context) {

	var med entity.MED

	if err := c.ShouldBindJSON(&med); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&med).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": med})

}

func GetMed(c *gin.Context) {
	var GetMed entity.MED
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM med WHERE id = ?", id).Scan(&GetMed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetMed})
}


func ListMed(c *gin.Context) {

	var med []entity.MED
	if err := entity.DB().Raw("SELECT * FROM med").Scan(&med).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": med})

}