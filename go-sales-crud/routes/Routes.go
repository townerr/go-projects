package routes

import (
	"go-sales-crud/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Home")
	})

	r.GET("/sales", controllers.GetAllSales)

	r.GET("/sale", controllers.GetSaleById)

	r.POST("/sales", controllers.CreateSale)

	r.DELETE("/sales", controllers.DeleteSaleById)
}
