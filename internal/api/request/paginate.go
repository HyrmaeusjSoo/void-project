package request

import (
	"errors"
	"strconv"
	"void-project/global"
	"void-project/internal/model/base"
	"void-project/pkg/logger"

	"github.com/gin-gonic/gin"
)

// 分页请求 参数验证/转换
func PageQuery(c *gin.Context) (pager *base.Pager, err error) {
	pager = &base.Pager{Page: 1, Size: global.Config.System.PageSize}
	if page := c.Query("page"); page != "" {
		pager.Page, err = strconv.Atoi(page)
		if err != nil {
			logger.LogError("[request-paginate-PageQuery:19] page参数无效：" + err.Error())
			return nil, errors.New("page参数无效：" + err.Error())
		}
	}
	if size := c.Query("size"); size != "" {
		pager.Size, err = strconv.Atoi(size)
		if err != nil {
			logger.LogError("[request-paginate-PageQuery:26] size参数无效：" + err.Error())
			return nil, errors.New("size参数无效：" + err.Error())
		}
	}

	return
}
