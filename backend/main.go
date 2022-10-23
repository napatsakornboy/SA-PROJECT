package main

import (
	"github.com/gin-gonic/gin"
	"github.com/napatsakornboy/Project/controller"
	"github.com/napatsakornboy/Project/entity"
)

func main() {

	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/ListMedicine", controller.ListMedicine)
	r.POST("/CreateMedicine", controller.CreateMedicine)

	r.GET("/ListWhere", controller.ListWhere)
	r.POST("/CreateWhere", controller.CreateWhere)

	r.GET("/ListDoctor", controller.ListDoctor)
	r.POST("/CreateDoctor", controller.CreateDoctor)

	r.GET("/ListBasket", controller.ListBasket)
	r.POST("/CreateBasket", controller.CreateBasket)

	r.GET("/ListSymtomp", controller.ListSymtomp)
	r.POST("/CreateSymtomp", controller.CreateSymtomp)
	r.Run()

}
func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
