package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FriendInvite(c *gin.Context) {
	c.String(http.StatusOK, "PENDING")
}

func FriendAnswer(c *gin.Context) {
	c.String(http.StatusOK, "PENDING")
}
