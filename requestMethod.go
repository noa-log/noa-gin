/*
 * @Author: nijineko
 * @Date: 2025-06-09 13:48:25
 * @LastEditTime: 2025-06-09 13:51:47
 * @LastEditors: nijineko
 * @Description: request method utility
 * @FilePath: \noa-gin\requestMethod.go
 */
package noagin

import (
	"fmt"

	"github.com/noa-log/colorize"
)

/**
 * @description: format request method with color
 * @param {string} method HTTP request method
 * @return {string} formatted request method with color
 */
func RequestMethodColor(method string) string {
	RequestMethodFormatStr := " %-7s "
	switch method {
	case "GET":
		return colorize.BlueBackground(fmt.Sprintf(RequestMethodFormatStr, method))
	case "HEAD":
		return colorize.MagentaBackground(fmt.Sprintf(RequestMethodFormatStr, method))
	case "POST":
		return colorize.CyanBackground(fmt.Sprintf(RequestMethodFormatStr, method))
	case "PUT":
		return colorize.BrightYellowBackground(fmt.Sprintf(RequestMethodFormatStr, method))
	case "DELETE":
		return colorize.RedBackground(fmt.Sprintf(RequestMethodFormatStr, method))
	case "CONNECT":
		return fmt.Sprintf(RequestMethodFormatStr, method)
	case "OPTIONS":
		return colorize.WhiteBackground(fmt.Sprintf(RequestMethodFormatStr, method))
	case "TRACE":
		return fmt.Sprintf(RequestMethodFormatStr, method)
	case "PATCH":
		return colorize.BrightGreenBackground(fmt.Sprintf(RequestMethodFormatStr, method))
	default:
		return fmt.Sprintf(RequestMethodFormatStr, method)
	}
}
