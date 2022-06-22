package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/tboerc/gwall/server/models"
	pb "github.com/tboerc/gwall/shared/proto"
)

func UserCreate(c *gin.Context) {
	r := &pb.CreateUserRequest{}
	if err := c.MustBindWith(r, binding.ProtoBuf); err != nil {
		return
	}

	u := &models.User{Username: r.Username, Password: r.Password}

	if err := u.Create(); err != nil {
		c.ProtoBuf(http.StatusInternalServerError, &pb.CreateUserResponse{})
		return
	}

	c.ProtoBuf(http.StatusOK, &pb.CreateUserResponse{})
}

func UserLogin(c *gin.Context) {
	c.String(http.StatusOK, "PENDING")
}

func UserSync(c *gin.Context) {
	c.String(http.StatusOK, "PENDING")
}
