package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpResponse(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}
