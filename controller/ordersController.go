package controller

import (
	"assignment2/database"
	helpers "assignment2/helper"
	"assignment2/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func CreateOrders(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Orders := models.Orders{}

	if contentType == appJSON {
		err := c.ShouldBindJSON(&Orders)
		if err != nil {
			panic(err.Error())
			return
		}
	} else {
		err := c.ShouldBind(&Orders)
		if err != nil {
			panic(err.Error())
			return
		}
	}

	err := db.Debug().Create(&Orders).Error

	if err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Success get data",
		"data":    Orders,
	})

}
func GetOrders(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	var Orders []models.Orders
	//var Item models.Item
	//var Order models.Item

	if contentType == appJSON {
		err := c.ShouldBindJSON(&Orders)
		if err != nil {
			panic(err.Error())
			return
		}
	} else {
		err := c.ShouldBind(&Orders)
		if err != nil {
			panic(err.Error())
			return
		}
	}

	//err := db.Debug().Find(&Orders).Error
	err := db.Preload("Items").Find(&Orders).Error

	if err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if len(Orders) == 0 {
		c.JSON(http.StatusCreated, gin.H{
			"status":  true,
			"message": "Data kosong",
			"data":    Orders,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status":  true,
			"message": "Success get data",
			"data":    Orders,
		})
	}

}

func UpdateOrders(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Orders := models.Orders{}
	orderId, _ := strconv.Atoi(c.Param("orderId"))

	if contentType == appJSON {
		err := c.ShouldBindJSON(&Orders)
		if err != nil {
			panic(err.Error())
			return
		}
	} else {
		err := c.ShouldBind(&Orders)
		if err != nil {
			panic(err.Error())
			return
		}
	}
	//var dataItem models.Item
	err := db.Where("order_id = ?", orderId).Preload("Items").First(&Orders).Error
	//db.Where("customer_id = ?", param["id"]).Preload("Contacts").First(&customer)
	if err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Success get data",
		"data":    Orders,
	})

}

func DeleteOrder(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Orders := models.Orders{}
	orderId, _ := strconv.Atoi(c.Param("orderId"))

	if contentType == appJSON {
		err := c.ShouldBindJSON(&Orders)
		if err != nil {
			panic(err.Error())
			return
		}
	} else {
		err := c.ShouldBind(&Orders)
		if err != nil {
			panic(err.Error())
			return
		}
	}
	Orders.OrderId = uint(orderId)

	err := db.Debug().Model(&Orders).Where("order_id = ?", orderId).Delete(&Orders, orderId).Error
	//err := db.Where("order_id = ?", orderId).Association("Items").Delete(&Orders).Error()

	if err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Success delete data",
	})

}
