/*
 * @Author: nijineko
 * @Date: 2025-06-09 13:08:32
 * @LastEditTime: 2025-06-09 13:54:54
 * @LastEditors: nijineko
 * @Description: noa gin handle
 * @FilePath: \noa-gin\handle_test.go
 */
package noagin

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/noa-log/noa"
)

func TestLogHandle(t *testing.T) {
	Log := noa.NewLog()

	gin.SetMode(gin.TestMode)
	gin.ForceConsoleColor()
	// gin.DefaultWriter = os.Stdin
	r := gin.Default()
	r.Use(LogHandle(Log))

	r.Any("/test", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	r.Any("/test/100", func(c *gin.Context) {
		c.String(100, "Hello, World!")
	})
	r.Any("/test/200", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	r.Any("/test/300", func(c *gin.Context) {
		c.String(300, "Multiple Choices")
	})
	r.Any("/test/400", func(c *gin.Context) {
		c.String(400, "Not Found")
	})
	r.Any("/test/500", func(c *gin.Context) {
		c.String(500, "Internal Server Error")
	})

	go func() {
		time.Sleep(time.Second) // Wait for the server to start

		RequestMethodList := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"}
		for _, Method := range RequestMethodList {
			Request, err := http.NewRequest(Method, "http://localhost:8080/test", nil)
			if err != nil {
				t.Errorf("Failed to make request: %v", err)
			}
			if _, err := http.DefaultClient.Do(Request); err != nil {
				t.Errorf("Failed to send request: %v", err)
			}
		}

		Request, err := http.NewRequest("GET", "http://localhost:8080/test/100", nil)
		if err != nil {
			t.Errorf("Failed to make request: %v", err)
		}
		if _, err := http.DefaultClient.Do(Request); err != nil {
			t.Errorf("Failed to send request: %v", err)
		}
		Request, err = http.NewRequest("GET", "http://localhost:8080/test/200", nil)
		if err != nil {
			t.Errorf("Failed to make request: %v", err)
		}
		if _, err := http.DefaultClient.Do(Request); err != nil {
			t.Errorf("Failed to send request: %v", err)
		}
		Request, err = http.NewRequest("GET", "http://localhost:8080/test/300", nil)
		if err != nil {
			t.Errorf("Failed to make request: %v", err)
		}
		if _, err := http.DefaultClient.Do(Request); err != nil {
			t.Errorf("Failed to send request: %v", err)
		}
		Request, err = http.NewRequest("GET", "http://localhost:8080/test/400", nil)
		if err != nil {
			t.Errorf("Failed to make request: %v", err)
		}
		if _, err := http.DefaultClient.Do(Request); err != nil {
			t.Errorf("Failed to send request: %v", err)
		}
		Request, err = http.NewRequest("GET", "http://localhost:8080/test/500", nil)
		if err != nil {
			t.Errorf("Failed to make request: %v", err)
		}
		if _, err := http.DefaultClient.Do(Request); err != nil {
			t.Errorf("Failed to send request: %v", err)
		}
	}()

	if err := r.Run("localhost:8080"); err != nil {
		t.Errorf("Failed to start server: %v", err)
	}
}
