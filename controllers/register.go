package controllers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/NafisaTojiboyeva/go-register/config"
	"github.com/NafisaTojiboyeva/go-register/models"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {

	var newUser models.RegisterInfo

	if err := ctx.BindJSON(&newUser); err != nil {
		log.Print(err)
		return
	}

	db, err := config.DB()

	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Can not connected to database"})
	}

	var user models.User

	isExists := db.Where("phone = ?", newUser.Phone).Find(&user)

	if isExists.RowsAffected != 1 {
		rand.Seed(time.Now().UnixNano())

		user := models.User{
			Fullname:   newUser.Fullname,
			Phone:      newUser.Phone,
			Password:   newUser.Password,
			SmsConfirm: strconv.Itoa(rand.Intn(1000000 - 100000)),
		}

		fmt.Println(user)

		SendConfirmCode(user.Phone, user.SmsConfirm)

		db.Select("Fullname", "Phone", "Password", "SmsConfirm").Create(&user)

	} else {

		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "This phone number is already exists!"})
	}
}
