package controller

import (
	"github.com/napatsakornboy/Project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

func CreateDoc(c *gin.Context) {

	var doc entity.DOC

	if err := c.ShouldBindJSON(&doc); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&doc).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": doc})

}

func GetDoc(c *gin.Context) {
	var GetDoc entity.DOC
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM doc WHERE id = ?", id).Scan(&GetDoc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetDoc})
}

func ListDoc(c *gin.Context) {

	var doc []entity.DOC
	if err := entity.DB().Raw("SELECT * FROM doc").Scan(&doc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doc})

}
