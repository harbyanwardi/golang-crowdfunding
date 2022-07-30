package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// fmt.Println("Connect Database Success")

	// var users []user.User

	// db.Find(&users)

	// router := gin.Default()
	// router.GET("/handler", handler)
	// router.Run()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/checkemail", userHandler.CheckEmailAvailability)
	api.POST("/avatar", userHandler.UploadAvatar)

	router.Run()

	//input dari user
	//handler, mapping input dari user -> struct input
	//service : melakukan mapping dari struct input ke struct User
	//repository
	//db

}
