package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"column:id;primaryKey"` 
	Name  string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Age  int `gorm:"column:age"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}


func main() {
	dsn:= "root:@tcp(localhost:3306)/openapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	router := gin.Default()
	
	router.GET("/v1/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	router.GET("/v1/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User
		db.First(&user, id)
		c.JSON(http.StatusOK, user)
	})

	router.PUT("/v1/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User
		db.First(&user, id)
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&user)
		c.JSON(http.StatusOK, user)
	})

	router.DELETE("/v1/user/:id", func(c *gin.Context) {	
		id := c.Param("id")
		var user User
		db.First(&user, id)
		db.Delete(&user)
		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	})
	
	router.POST("/v1/user", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&user)
		c.JSON(http.StatusOK, user)
	})

	router.Run(":3000")

}
