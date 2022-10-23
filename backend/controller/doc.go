package controller

import (
	"github.com/napatsakornboy/Project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

func CreateDoctor(c *gin.Context) {

	var doctor entity.DOCTOR

	if err := c.ShouldBindJSON(&doctor); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&doctor).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": doctor})

}

func GetDoctor(c *gin.Context) {
	var GetDoctor entity.DOCTOR
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM doctor WHERE id = ?", id).Scan(&GetDoctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetDoctor})
}

func ListDoctor(c *gin.Context) {

	var doctor []entity.DOCTOR
	if err := entity.DB().Raw("SELECT * FROM doctor").Scan(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doctor})

}
