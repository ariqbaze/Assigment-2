package controllers

import (
	"Assigment-2/database"
	"Assigment-2/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	database.StartDB()

	var newOrder models.Order
	db := database.GetDB()

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	Order := models.Order{
		OrderedAt:    newOrder.OrderedAt,
		CustomerName: newOrder.CustomerName,
		Items:        newOrder.Items,
	}

	err := db.Create(&Order).Error
	fmt.Println(err)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// message := fmt.Sprintln("New User Data :", Order)
	ctx.JSON(http.StatusOK, gin.H{
		"data":    Order,
		"message": "berhasil menambahkan data",
		"success": true,
	})
}

func GetAllOrder(ctx *gin.Context) {
	database.StartDB()
	db := database.GetDB()
	order := []models.Order{}
	err := db.Preload("Items").Find(&order).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Printf("User Datas With Products")
	fmt.Printf("%+v \n", order)
	ctx.JSON(http.StatusOK, gin.H{
		"Data": order,
	})
}

func UpdateOrder(ctx *gin.Context) {
	database.StartDB()
	id := ctx.Param("orderId")
	db := database.GetDB()

	order := models.Order{}
	var updated_order models.Order
	if err := ctx.ShouldBindJSON(&updated_order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Model(&order).Where("order_id = ?", id).Updates(updated_order).Error

	if err != nil {
		fmt.Println("Error updating user data:", err)

		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    updated_order,
		"message": "berhasil menambahkan data",
		"success": true,
	})
}

func DeleteOrder(ctx *gin.Context) {
	database.StartDB()
	id := ctx.Param("orderId")
	db := database.GetDB()

	order := models.Order{}
	item := models.Item{}

	err := db.Model(&item).Where("order_id = ?", id).Delete(&item).Error
	if err != nil {
		fmt.Println("Error Deleting data:", err)

		return
	}
	err = db.Model(&order).Where("order_id = ?", id).Delete(&order).Error

	if err != nil {
		fmt.Println("Error Deleting data:", err)

		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    id,
		"message": "berhasil menghapus data",
		"success": true,
	})
}
