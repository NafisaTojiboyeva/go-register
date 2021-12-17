package main

import (
	"log"
	"os"
	// "net/http"
	c "github.com/NafisaTojiboyeva/go-register/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")

	router := gin.Default()

	router.POST("/login", c.Login)
	router.POST("/register", c.Register)
	router.POST("/confirm/:id", c.Confirm)
	router.POST("/resetpassword", c.ResetPassword)
	router.GET("/logout", c.Logout)

	router.Run(PORT)
}
