package middlewares

import (
	"fmt"
	"go_project_template/internal/pkg/apires"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
				logrus.Errorln(err)
				PrintStack()
				var Err *apires.ApiRes
				if e, ok := err.(*apires.ApiRes); ok {
					Err = e
				} else {
					apires.SystemException.MoreInfo = err
					Err = &apires.SystemException
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
