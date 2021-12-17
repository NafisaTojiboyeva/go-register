package controllers

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/NafisaTojiboyeva/go-register/config"
	"github.com/NafisaTojiboyeva/go-register/models"

	"github.com/gin-gonic/gin"
)

func ResetPassword(ctx *gin.Context) {

	var newPasswordData models.ChangePasswordInfo

	if err := ctx.BindJSON(&newPasswordData); err != nil {
		log.Print(err)
		return
	}

	db, err := config.DB()

	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Can not connected to database"})
	}

	session, _ := store.Get(ctx.Request, "mysession")
	id := session.Values["user_id"]

	rand.Seed(time.Now().UnixNano())

	newPasswordData.SmsConfirm = strconv.Itoa(rand.Intn(1000000 - 100000))

	var user models.User

	result := db.Model(&user).Where("user_id = ? and phone = ?", id, newPasswordData.Phone).Updates(models.User{
		Password:   newPasswordData.Password,
		SmsConfirm: newPasswordData.SmsConfirm,
		IsVerified: false,
	})

	if result.RowsAffected == 1 {

		SendConfirmCode(newPasswordData.Phone, newPasswordData.SmsConfirm)

		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Please enter confirmation code!"})

	} else {

		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invailed data!"})
	}

}
