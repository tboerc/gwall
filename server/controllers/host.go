package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Host(c *gin.Context) {
	c.String(http.StatusOK, "PENDING")
}
