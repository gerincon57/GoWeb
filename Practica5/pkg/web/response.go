package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type response struct {
	Data interface{} `json:"data"`
}

// funciones
func ResponseOk(c *gin.Context, data interface{}, status int) {
	c.JSON(status, response{
		Data: data,
	})

}

func ResponseFail(c *gin.Context, status int, err error) {
	c.JSON(status, errorResponse{
		Status:  status,
		Code:    http.StatusText(status),
		Message: err.Error(),
	})
}
