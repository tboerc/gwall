package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tboerc/gwall/server/controllers"
	"github.com/tboerc/gwall/server/messages"
	"github.com/tboerc/gwall/server/models"
	"github.com/tboerc/gwall/shared"
)

func main() {
	if err := models.Connect(); err != nil {
		log.Fatalln(messages.ErrDatabaseConn)
	}

	r := gin.Default()

	r.SetTrustedProxies(nil)

	v1 := r.Group("/v1")
	{
		v1.POST("/user/create", controllers.UserCreate)

		v1.POST("/user/login", controllers.UserLogin)

		v1.POST("/user/sync", controllers.UserSync)

		v1.POST("/friend/invite", controllers.FriendInvite)

		v1.POST("/friend/answer", controllers.FriendAnswer)

		v1.POST("/host", controllers.Host)
	}

	if err := r.Run(":" + shared.Getenv("PORT", "8080")); err != nil {
		log.Fatalln(messages.ErrServerRun)
	}
}
