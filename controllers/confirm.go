package controllers

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/NafisaTojiboyeva/go-register/config"
	"github.com/NafisaTojiboyeva/go-register/models"

	"github.com/gin-gonic/gin"
)

func Confirm(ctx *gin.Context) {

	var code models.Code

	if err := ctx.Bind(&code); err != nil {
		log.Print(err)
		return
	}

	id := ctx.Param("id")

	db, err := config.DB()

	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Can not connected to database"})
	}

	var user models.User

	result := db.Where("user_id = ? and sms_confirm = ?", id, code.Code).Find(&user)

	if result.RowsAffected == 1 {

		db.Model(&user).Where("user_id = ?", id).Update("is_verified", true)

		session, _ := store.Get(ctx.Request, "mysession")

		session.Values["user_id"] = user.Id

		session.Values["fullname"] = user.Fullname

		session.Save(ctx.Request, ctx.Writer)

		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully registered! Welcome " + user.Fullname})
	} else {

		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invailed confirmation code!"})

	}

}
