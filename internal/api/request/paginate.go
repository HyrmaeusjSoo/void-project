package request

import (
	"strconv"
	"void-project/global"
	"void-project/internal/api/response"
	"void-project/internal/api/response/apierr"
	"void-project/internal/model/base"
	"void-project/pkg/logger"

	"github.com/gin-gonic/gin"
)

// 分页请求 参数验证/转换
func PageQuery(c *gin.Context) (pager base.Pager, err error) {
	pager = base.Pager{Page: 1, Size: global.Configs.System.PageSize}
	if page := c.Query("page"); page != "" {
		pager.Page, err = strconv.Atoi(page)
		if err != nil {
			logger.LogError("page参数无效：" + err.Error())
			err = apierr.InvalidPaginateParameter.Wrap("page参数无效")
			response.FailError(c, err)
			return pager, err
		}
	}
	if size := c.Query("size"); size != "" {
		pager.Size, err = strconv.Atoi(size)
		if err != nil {
			logger.LogError("size参数无效：" + err.Error())
			err = apierr.InvalidPaginateParameter.Wrap("size参数无效")
			response.FailError(c, err)
			return pager, err
		}
	}
	return
}

// cursor方式分页 参数转换
func CursorQuery(c *gin.Context) (cursor base.Cursor) {
	next := c.Query("next")
	if next != "" {
		cursor = base.Next(next).Decode()
	} else {
		cursor = base.Cursor{}
	}
	return
}
