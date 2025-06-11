/*
 * @Author: nijineko
 * @Date: 2025-06-09 13:08:32
 * @LastEditTime: 2025-06-11 11:59:24
 * @LastEditors: nijineko
 * @Description: noa gin handle
 * @FilePath: \noa-gin\handle.go
 */
package noagin

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/noa-log/noa"
)

var (
	DEFAULT_LOG_SOURCE = "Gin"
)

/**
 * @description: log handle
 * @param {*noa.LogConfig} Log noa log instance
 * @return {gin.HandlerFunc} gin handler function
 */
func LogHandle(Log *noa.LogConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record request time
		RequestStartTime := time.Now()
		c.Next()
		RequestEndTime := time.Now()
		RequestTime := RequestEndTime.Sub(RequestStartTime)

		Log.Info(DEFAULT_LOG_SOURCE, fmt.Sprintf("%s| %13v | %15s |%s %#v",
			StatusCodeColor(c.Writer.Status()),
			RequestTime,
			c.ClientIP(),
			RequestMethodColor(c.Request.Method),
			c.Request.URL.Path,
		))
	}
}
