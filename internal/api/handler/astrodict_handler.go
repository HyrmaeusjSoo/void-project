package handler

import (
	"chat/internal/api/response"
	"chat/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AstroDict struct{}

var adService = &service.AstroDictService{}

func (*AstroDict) Fetch(c *gin.Context) {
	adRes, err := adService.Fetch(c.Param("name"))
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}
	response.Success(c, adRes)
}
