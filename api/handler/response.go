package handler

import (
	"dousheng/pkg/errdeal"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpResponse(c *gin.Context, response *errdeal.Response) {
	c.JSON(http.StatusOK, errdeal.Response{
		StatusCode:    response.StatusCode,
		StatusMessage: response.StatusMessage,
		Data:          response.Data,
	})
}
