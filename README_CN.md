# Noa Gin
Noa Log 的 Gin 集成模块，可以快读的将 Noa 集成到 Gin 框架中。

# 安装
```bash
go get -u github.com/noa-log/noa-gin
```

## 快速开始
```go
package main

import (
    "os"
    "github.com/gin-gonic/gin"
    "github.com/noa-log/noa"
    noagin "github.com/noa-log/noa-gin"
)

func main() {
    // 创建一个新的日志实例
    logger := noa.NewLog()

    // 创建一个 Gin 引擎实例
    gin.DefaultWriter = os.Stdin // 禁用Gin默认日志输出
    r := gin.Default()

    // 使用 Noa 中间件
    r.Use(noagin.LogHandle(logger))

    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, Noa!")
    })

    if err := r.Run(":8080"); err != nil {
        logger.Error("Gin", err)
    }
}
```

## 许可
本项目基于[Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0)协议开源。使用时请遵守协议的条款。