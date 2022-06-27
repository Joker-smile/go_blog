package util

import (
	"blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := strconv.Atoi(c.Query("page"))
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}
