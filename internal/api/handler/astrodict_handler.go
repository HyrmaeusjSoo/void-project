package handler

import (
	"net/http"
	"void-project/internal/api/response"
	"void-project/internal/service"

	"github.com/gin-gonic/gin"
)

type AstroDict struct {
	service *service.AstroDictService
}

func NewAstroDict() *AstroDict {
	return &AstroDict{service.NewAstroDictService()}
}

func (ad *AstroDict) Fetch(c *gin.Context) {
	astro, err := ad.service.Fetch(c.Param("name"))
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}
	response.Success(c, astro)
}
