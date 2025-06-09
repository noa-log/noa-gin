# Noa Gin
The Gin integration module for Noa Log, allowing quick integration of Noa into the Gin framework.

# Installation
```bash
go get -u github.com/noa-log/noa-gin
```

## Quick Start
```go
package main

import (
    "os"
    "github.com/gin-gonic/gin"
    "github.com/noa-log/noa"
    noagin "github.com/noa-log/noa-gin"
)

func main() {
    // Create a new logger instance
    logger := noa.NewLog()

    // Create a Gin engine instance
    gin.DefaultWriter = os.Stdin // Disable Gin's default log output
    r := gin.Default()

    // Use Noa handle
    r.Use(noagin.LogHandle(logger))

    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, Noa!")
    })

    if err := r.Run(":8080"); err != nil {
        logger.Error("Gin", err)
    }
}
```

## License
This project is open-sourced under the [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0). Please comply with the terms when using it.