package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/payment"
	"bwastartup/transaction"
	"bwastartup/user"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	paymentService := payment.NewService()
	campaignService := campaign.NewService(campaignRepository)
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	// user, _ := userService.GetUserbyId(3)
	// input := transaction.CreateTransactionsInput{
	// 	CampaignID: 1,
	// 	Amount:     130000,
	// 	User:       user,
	// }
	// transactionService.CreateTransaction(input)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "text/plain", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Static("/images", "./images")
	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/checkemail", userHandler.CheckEmailAvailability)
	api.POST("/avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/user/fetch", authMiddleware(authService, userService), userHandler.FetchUser)

	//campaign
	api.GET("/campaign", campaignHandler.GetCampaign)
	api.GET("/campaign/:id", campaignHandler.GetCampaignDetail)
	api.POST("/campaign", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaign/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)
	api.GET("/campaign/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)

	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransactions)
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	router.Run()

	//input dari user
	//handler, mapping input dari user -> struct input
	//service : melakukan mapping dari struct input ke struct User
	//repository
	//db

}

//dengan handler kita tetap harus menggunakan gin.Context , namun tidak bisa langsung passing authservice & userservice,
//maka buat function mereturn function
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {

			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		arrayToken := strings.Split(authHeader, " ")
		var TokenString string
		if len(arrayToken) == 2 {
			TokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(TokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64)) //param dari service GetUserbyId adalah Int , maka dibuat int
		user, err := userService.GetUserbyId(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentuser", user)
	}
}
