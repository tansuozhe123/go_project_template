package middlewares

import (
	"fmt"
	"go_project_template/pkg/commonres"
	"go_project_template/pkg/logger"
	"net/http"
	"runtime"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))
}

// panic错误统一处理中间件
func PanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Logger.Error("panic error: ", zap.Error(err.(error)))
				PrintStack()
				var Err *commonres.CommonRes
				if e, ok := err.(*commonres.CommonRes); ok {
					Err = e
				} else {
					commonres.SystemException.Message = err
					Err = &commonres.SystemException
				}
				// 记录一个错误的日志
				c.JSON(http.StatusCreated, Err)
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
