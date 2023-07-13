package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAuthUserId(c *gin.Context) (userId uint) {
	switch id := c.MustGet("userId").(type) {
	case int:
		userId = uint(id)
	case uint:
		userId = id
	case string:
		intUid, _ := strconv.Atoi(id)
		userId = uint(intUid)
	default:
		userId = 0
	}
	return
}
