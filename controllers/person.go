package controllers

import (
	"net/http"
	"pratice/structs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) GetOrder(c *gin.Context) {
	var (
		Order  structs.Order
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&Order).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": Order,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		Orders []structs.Order
		result gin.H
	)
	idb.DB.Find(&Orders)
	if len(Orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": Orders,
			"count":  len(Orders),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateOrder(c *gin.Context) {
	var (
		Order  structs.Order
		result gin.H
	)
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	itemcode := c.PostForm("itemcode")
	description := c.PostForm("description")
	itemCode, _ := strconv.ParseInt(itemcode, 10, 64)
	Quantity := c.PostForm("quantity")
	qty, _ := strconv.ParseInt(Quantity, 10, 64)
	Order.Items.ItemCode = itemCode
	Order.Items.Quantity = qty
	Order.Items.Description = description
	Order.First_Name = first_name
	Order.Last_Name = last_name
	// idb.DB.Create(&Order)
	result = gin.H{
		"result": Order,
		"count":  itemcode,
		"qty":    first_name,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateOrder(c *gin.Context) {
	id := c.PostForm("id")
	var (
		Order    structs.Order
		newOrder structs.Order
		result   gin.H
	)
	err := idb.DB.First(&Order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	description := c.PostForm("description")
	itemcode := c.PostForm("itemcode")
	itemCode, _ := strconv.ParseInt(itemcode, 10, 64)
	Quantity := c.PostForm("quantity")
	qty, _ := strconv.ParseInt(Quantity, 10, 64)
	newOrder.First_Name = first_name
	newOrder.Last_Name = last_name
	newOrder.OrderedAt = time.Now()
	newOrder.Items.ItemCode = itemCode
	newOrder.Items.Quantity = qty
	newOrder.Items.Description = description
	err = idb.DB.Model(&Order.Items).Updates(&newOrder.Items).Error
	// err = idb.DB.Model(&Order).Updates(newOrder).Error
	if err != nil {
		result = gin.H{
			"result": "update Data failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteOrder(c *gin.Context) {
	var (
		Order  structs.Order
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&Order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&Order).Error
	if err != nil {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
