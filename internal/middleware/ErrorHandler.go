package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/internal/errUtils"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	isBodyWritten := ctx.Writer.Written()
	err := ctx.Errors.Last()
	if err != nil {
		castedError, ok := errUtil.CastApplicationError(err)
		statusCode := http.StatusInternalServerError // default 500 internal server errUtils
		if !ok {
			fmt.Errorf(err.Error())
			response := errUtil.MakeBaseResponse(statusCode)
			if !isBodyWritten {
				ctx.JSON(statusCode, response)
			}
			return
		}
		code := castedError.Code
		response := errUtil.MakeBaseResponse(code)

		// errUtils stack logging
		//switch {
		//// 40000 ~ 49999
		//case errUtil.ERROR_INVALID_PARAMS <= castedError.Code && errUtil.ERROR_INTERNAL_SERVER > castedError.Code:
		//	log.Warn(err.Error())
		//default:
		//	log.Error(err.Error())
		//}

		if !isBodyWritten {
			ctx.JSON(statusCode, response)
		}
		return
	}
}
