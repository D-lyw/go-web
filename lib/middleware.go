package lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// MyLogger 自定义中间件
func MyLogger() gin.HandlerFunc {

	return func(context *gin.Context) {
		t := time.Now()
		fmt.Printf("Middleware Stage 1: %s\n", t)

		// 调用下一个中间件
		context.Next()

		afterTime := time.Now()
		fmt.Printf("Middleward Stage 2: %s\n", afterTime)
	}
}
