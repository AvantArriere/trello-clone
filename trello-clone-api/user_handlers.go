package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// SignUp - sign up handler
func SignUp(ctx *gin.Context) {
	// read request body
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	reqData := SignUpRequest{}

	// check if request parameter is valid
	if err := json.Unmarshal(body, &reqData); err != nil {
		ctx.String(http.StatusBadRequest, "The request parameter is not valid.")
		return
	} else if err := validator.New().Struct(reqData); err != nil {
		ctx.String(http.StatusBadRequest, "The request parameter is not valid.")
		return
	}

	// set user info
	user := User{
		Email:    reqData.Email,
		Username: reqData.Username,
	}
	if !user.EncryptPassword(reqData.Password) {
		ctx.String(http.StatusBadRequest, "password encryption failed.")
		return
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	// create user
	if err := mariaDB.DB.Save(&user).Error; err != nil {
		ctx.String(http.StatusInternalServerError, "internal server error has occured.")
		return
	}

	userDetail := UserDetail{
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// create user detail
	if err := mariaDB.DB.Save(&userDetail).Error; err != nil {
		mariaDB.DB.Where("id = ", user.ID).Delete(&user)
		ctx.String(http.StatusInternalServerError, "internal server error has occured.")
		return
	}

	// create confirm hash
	confirmHash := ConfirmHash{UserID: user.ID}
	confirmHash.SetHash()
	mariaDB.DB.Save(&confirmHash)

	// send confirm email must be added here
	/**/
	// sendEmail()

	ctx.Status(http.StatusCreated)
}

// func sendEmail() {
// 	var from = "leeyjeen@naver.com"
// 	var msg = []byte("dummy message")
// 	var recipients = []string{"leeyjeen@naver.com"}
// 	auth := smtp.PlainAuth(
// 		"",
// 		"leeyjeen@naver.com",
// 		"temppassword",
// 		"smtp.naver.com:587")
// 	if err := smtp.SendMail("smtp.naver.com:587", auth, from, recipients, msg); err != nil {
// 		log.Fatal(err)
// 	}
// }
