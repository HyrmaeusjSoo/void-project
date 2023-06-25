package request

import (
	"errors"
	"strconv"
	"void-project/global"

	"github.com/gin-gonic/gin"
)

// 分页请求 参数验证/转换
func PageQuery(c *gin.Context) (page, size int, err error) {
	page = 0
	size = 0
	err = nil

	pageStr := c.Query("page")
	if pageStr == "" {
		return page, size, errors.New("page参数无效")
	}
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		return
	}
	sizeStr := c.Query("size")
	if sizeStr == "" {
		size = global.Config.System.PageSize
		return
	}
	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		return
	}
	return
}
