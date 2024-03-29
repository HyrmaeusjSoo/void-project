package handler

import (
	"void-project/internal/api/response"
	"void-project/internal/api/response/apierr"
	"void-project/internal/service"
	"void-project/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AstroDict struct {
	service          *service.AstroDictService
	translateService *service.TranslateService
}

func NewAstroDict() *AstroDict {
	return &AstroDict{service.NewAstroDictService(), service.NewTranslateService()}
}

// 从远程查询
func (ad *AstroDict) FetchRemote(c *gin.Context) {
	astro, err := ad.service.FetchRemote(c.Param("name"))
	if err != nil {
		response.FailError(c, apierr.FetchFailed, err)
		return
	}
	response.Success(c, astro)
}

// 查询
func (ad *AstroDict) Fetch(c *gin.Context) {
	astro, err := ad.service.Fetch(c.Param("name"))
	if err != nil {
		response.FailError(c, apierr.FetchFailed, err)
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
		response.FailError(c, apierr.SaveFailed, err)
		return
	}
	response.SuccessOk(c)
}

// 测试翻译
func (ad *AstroDict) Translate(c *gin.Context) {
	text := c.Query("text")
	if text == "" {
		response.FailError(c, apierr.MissingRequiredParameter, "text")
		return
	}
	target := c.Query("target")
	if target == "" {
		response.FailError(c, apierr.MissingRequiredParameter, "target")
		return
	}
	res, err := ad.translateService.Translate(text, c.Query("source"), target)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.FetchFailed, err)
		return
	}
	response.Success(c, res)
}
