package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
			log.Errorf(err.Error())
			response := errUtil.MakeBaseResponse(statusCode)
			if !isBodyWritten {
				ctx.JSON(statusCode, response)
			}
			return
		}
		code := castedError.Code
		response := errUtil.MakeBaseResponse(code)

		if !isBodyWritten {
			///todo json 포맷이 아니라, htmx way 로
			ctx.JSON(statusCode, response)
		}
		return
	}
}
