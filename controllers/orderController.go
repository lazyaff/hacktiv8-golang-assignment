package controllers

import (
	"assignment-2/database"
	"assignment-2/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// get all orders
func GetAllOrders(ctx *gin.Context) {
	// inisialisasi variabel
	db := database.GetDB()
	var orders []models.Order

	// query database get all orders
	err := db.Preload("Items").Find(&orders).Error
	if err != nil {
		// error handling
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// response
	ctx.JSON(200, orders)
}


// create order
func CreateOrder(ctx *gin.Context) {
	// inisialisasi variabel
	db := database.GetDB()
	var order models.Order

	// binding data
	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		// error handling
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// query database create order
	err = db.Create(&order).Error
	if err != nil {
		// error handling
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// response
	ctx.JSON(201, order)
}


// update order
func UpdateOrder(ctx *gin.Context)  {
	// inisialisasi variabel
	db := database.GetDB()
	var order models.Order
	orderID := ctx.Param("id")

	// convert id to uint
	id, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		// error handling
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// query database get order by id
	err = db.First(&order, id).Error
	if err != nil {
		// error handling
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	// binding data
	err = ctx.ShouldBindJSON(&order)
	if err != nil {
		// error handling
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// hapus item dari order
	err = db.Delete(&order.Items, "order_id = ?", order.ID).Error
	if err != nil {
		// error handling
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// query database update order
	err = db.Save(&order).Error
	if err != nil {
		// error handling
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// response
	ctx.JSON(200, order)
}

func DeleteOrder(ctx *gin.Context) {
	// inisialisasi variabel
	db := database.GetDB()
	var order models.Order
	orderID := ctx.Param("id")

	// convert id to uint
	id, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		// error handling
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// query database get order by id
	err = db.First(&order, id).Error
	if err != nil {
		// error handling
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	// query database delete item
	err = db.Delete(&order.Items, "order_id = ?", order.ID).Error
	if err != nil {
		// error handling
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// query database delete order
	err = db.Delete(&order).Error
	if err != nil {
		// error handling
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// response
	ctx.JSON(200, "Success delete")
}