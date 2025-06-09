/*
 * @Author: nijineko
 * @Date: 2025-06-09 13:46:06
 * @LastEditTime: 2025-06-09 13:47:15
 * @LastEditors: nijineko
 * @Description: status code utility
 * @FilePath: \noa-gin\statusCode.go
 */
package noagin

import (
	"fmt"

	"github.com/noa-log/colorize"
)

/**
 * @description: format status code with color
 * @param {int} statusCode HTTP status code
 * @return {string} formatted status code with color
 */
func StatusCodeColor(statusCode int) string {
	StatusCodeFormatStr := " %3d "
	switch statusCode / 100 % 10 {
	case 1:
		return colorize.WhiteBackground(fmt.Sprintf(StatusCodeFormatStr, statusCode))
	case 2:
		return colorize.GreenBackground(fmt.Sprintf(StatusCodeFormatStr, statusCode))
	case 3:
		return colorize.BrightWhiteBackground(fmt.Sprintf(StatusCodeFormatStr, statusCode))
	case 4:
		return colorize.YellowBackground(fmt.Sprintf(StatusCodeFormatStr, statusCode))
	case 5:
		return colorize.RedBackground(fmt.Sprintf(StatusCodeFormatStr, statusCode))
	default:
		return fmt.Sprintf(StatusCodeFormatStr, statusCode)
	}
}
