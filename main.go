package main

import (
	"pratice/config"
	"pratice/controllers"

	"github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql"
	// // "gorm.io/gorm"
	// // "gorm.io/driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	router := gin.Default()
	router.GET("/Order/:id", inDB.GetOrder)
	router.GET("/Orders", inDB.GetOrders)
	router.POST("/Orders", inDB.CreateOrder)
	router.PUT("/Order/:id", inDB.UpdateOrder)
	router.DELETE("/Order/:id", inDB.DeleteOrder)
	router.Run(":3000")
}
