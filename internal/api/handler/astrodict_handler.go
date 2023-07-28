package handler

import (
	"net/http"
	"void-project/internal/api/response"
	"void-project/internal/api/response/apierr"
	"void-project/internal/service"
	"void-project/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AstroDict struct {
	service *service.AstroDictService
}

func NewAstroDict() *AstroDict {
	return &AstroDict{service.NewAstroDictService()}
}

// 从远程查询
func (ad *AstroDict) FetchRemote(c *gin.Context) {
	astro, err := ad.service.FetchRemote(c.Param("name"))
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}
	response.Success(c, astro)
}

// 查询
func (ad *AstroDict) Fetch(c *gin.Context) {
	astro, err := ad.service.Fetch(c.Param("name"))
	if err != nil {
		response.FailError(c, apierr.FetchFailed)
		return
	}
	response.Success(c, astro)
}

// 同步到本地
func (ad *AstroDict) Sync(c *gin.Context) {
	lang := c.Param("lang")
	if lang != "ce" && lang != "ec" {
		response.FailError(c, apierr.InvalidParameter)
		return
	}
	err := ad.service.Sync(lang)
	if err != nil {
		logger.LogInfo(err)
		response.FailError(c, apierr.SaveFailed)
		return
	}
	response.SuccessOk(c)
}
