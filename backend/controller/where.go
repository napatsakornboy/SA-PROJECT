package controller

import (
	"github.com/napatsakornboy/Project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

func CreateWhere(c *gin.Context) {

	var where entity.WHERE

	if err := c.ShouldBindJSON(&where); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&where).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": where})

}

func GetWhere(c *gin.Context) {
	var GetWhere entity.WHERE
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM where WHERE id = ?", id).Scan(&GetWhere).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetWhere})
}


func ListWhere(c *gin.Context) {

	var where []entity.WHERE
	if err := entity.DB().Raw("SELECT * FROM where").Scan(&where).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": where})

}
