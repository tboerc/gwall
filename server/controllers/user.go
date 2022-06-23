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
	if err := c.ShouldBindWith(r, binding.ProtoBuf); err != nil {
		c.ProtoBuf(http.StatusBadRequest, &pb.CreateUserResponse{})
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
	r := &pb.LoginRequest{}
	if err := c.ShouldBindWith(r, binding.ProtoBuf); err != nil {
		c.ProtoBuf(http.StatusBadRequest, &pb.LoginResponse{})
		return
	}

	u := &models.User{Username: r.Username, Password: r.Password}

	t, err := u.Login()
	if err != nil {
		c.ProtoBuf(http.StatusBadRequest, &pb.LoginResponse{Info: &pb.Info{Message: err.Error()}})
		return
	}

	c.ProtoBuf(http.StatusOK, &pb.LoginResponse{Token: t})
}

func UserSync(c *gin.Context) {
	c.String(http.StatusOK, "PENDING")
}
