package controllers

import (
	"go-sales-crud/config"
	"go-sales-crud/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllSales(c *gin.Context) {
	sales := []models.Sale{}
	config.DB.Find(&sales)

	if len(sales) <= 0 {
		c.JSON(404, gin.H{
			"error": "Failed to get sales",
		})

		return
	}

	c.JSON(200, &sales)
}

func GetSaleById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	sale := models.Sale{}
	config.DB.First(&sale, id)

	if sale.Amount == "" {
		c.JSON(404, gin.H{
			"error": "Failed to get sale by id",
		})

		return
	}

	c.JSON(200, &sale)
}

func DeleteSaleById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	result := config.DB.Delete(&models.Sale{}, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Failed to delete sale",
		})

		return
	}

	c.JSON(200, gin.H{})
}

func CreateSale(c *gin.Context) {
	var body struct {
		UserId int
		Amount string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	sale := models.Sale{UserId: body.UserId, Amount: body.Amount}
	result := config.DB.Create(&sale)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Failed to create sale",
		})

		return
	}

	c.JSON(200, gin.H{})
}
