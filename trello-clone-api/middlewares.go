package main

import (
	"log"
	"time"
	"fmt"
	"errors"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)



var identityKey = "email"
var authMiddleware *jwt.GinJWTMiddleware


func setAuthMiddleware() {
	// the jwt middleware
	var err error
	authMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Trello Clone REST API",
		Key:         []byte("secret key"), // to be changed
		Timeout:     time.Hour * 3,
		MaxRefresh:  time.Hour * 3,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if user, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: user.Email,
				}
			}
			fmt.Println("Payload Func error occured.")
			return nil
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Email: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			email := loginVals.Email
			password := loginVals.Password

			user := User{}
			notExist := mariaDB.DB.Where("email = ?", email).Find(&user).RecordNotFound()
			if notExist {
				return nil, errors.New("")
			}
			if password == user.Password {
				return user, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(map[string]interface{}); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}