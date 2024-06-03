package handler

import (
	"os"
	"void-project/internal/api/request"
	"void-project/internal/api/response"
	"void-project/internal/api/response/apierr"
	"void-project/internal/service"

	"github.com/gin-gonic/gin"
)

type Visitor struct {
	service *service.VisitorService
}

func NewVisitor() *Visitor {
	return &Visitor{service.NewVisitorService()}
}

func (v *Visitor) IP(c *gin.Context) {
	info, err := v.service.IPQuery(c.Param("ip"))
	if err != nil {
		response.FailError(c, apierr.FetchFailed, err)
		return
	}
	response.Success(c, info)
}

func (v *Visitor) FetchLog(c *gin.Context) {
	begin := request.GetQueryDate(c, "begin_date")
	end := request.GetQueryDate(c, "end_date")
	logs, err := v.service.ReadLog(begin, end)
	if err != nil {
		if os.IsNotExist(err) {
			response.FailError(c, apierr.RecordNotFound)
			return
		}
		response.FailError(c, apierr.FetchFailed, err)
		return
	}
	response.Success(c, logs)
}

func (v *Visitor) Stat(c *gin.Context) {
	begin := request.GetQueryDate(c, "begin_date")
	end := request.GetQueryDate(c, "end_date")
	stats, err := v.service.Stat(begin, end)
	if err != nil {
		if os.IsNotExist(err) {
			response.FailError(c, apierr.RecordNotFound)
			return
		}
		response.FailError(c, apierr.FetchFailed, err)
		return
	}
	response.Success(c, stats)
}
