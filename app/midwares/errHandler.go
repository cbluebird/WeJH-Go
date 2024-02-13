package midwares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wejh-go/app/apiException"
	"wejh-go/config/logs"
)

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		fmt.Println(c.Errors)
		if length := len(c.Errors); length > 0 {
			e := c.Errors[length-1]
			err := e.Err
			if err != nil {
				var Err *apiException.Error
				if e, ok := err.(*apiException.Error); ok {
					Err = e
					if Err.StatusCode != http.StatusBadRequest {
						logs.WriteWarn(c, Err)
					} else {
						logs.WriteDebug(c, Err)
					}
				} else if e, ok := err.(error); ok {
					Err = apiException.OtherError(e.Error())
					logs.WriteError(c, err)
				} else {
					Err = apiException.ServerError
					logs.WriteError(c, err)
				}

				c.JSON(Err.StatusCode, Err)
				return
			}
		}

	}
}

// HandleNotFound
//
//	404处理
func HandleNotFound(c *gin.Context) {
	err := apiException.NotFound
	c.JSON(err.StatusCode, err)
	return
}
