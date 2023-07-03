package request

import (
	"errors"
	"strconv"
	"void-project/global"
	"void-project/pkg/logger"

	"github.com/gin-gonic/gin"
)

// 分页请求 参数验证/转换
func PageQuery(c *gin.Context) (page, size int, err error) {
	page = 1

	pageStr := c.Query("page")
	if pageStr == "" {
		return page, size, errors.New("page参数无效")
	}
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		logger.LogError("[request-paginate-PageQuery:22] page参数无效：" + err.Error())
		return page, size, errors.New("page参数无效：" + err.Error())
	}
	sizeStr := c.Query("size")
	if sizeStr == "" {
		size = global.Config.System.PageSize
		return
	}
	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		logger.LogError("[request-paginate-PageQuery:32] page参数无效：" + err.Error())
		return page, size, errors.New("page参数无效：" + err.Error())
	}
	return
}
