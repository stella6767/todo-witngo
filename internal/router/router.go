package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/internal/handler"
	"todo-app/internal/middleware"
)

func NewRouter(todoHandler *handler.TodoHandler) *gin.Engine {

	// Creates a router without any middleware by default
	router := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.Use(middleware.CustomLogger())
	router.Use(middleware.ErrorHandler)
	//router.Static("/assets", "./assets")

	registerTodoHandler(todoHandler, router)
	return router
}

func registerTodoHandler(handler *handler.TodoHandler, router *gin.Engine) {

	router.GET("/", handler.Index)
	router.GET("/todos", handler.GetTodosByPage)
	router.PUT("/todo/:id", handler.UpdateTodoStatus)
	router.DELETE("/todo/:id", handler.DeleteTodoById)
	router.POST("/todo", handler.CreateTodo)

	json := router.Group("json")
	json.GET("/test", handler.Test)
	json.GET("/test2", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.JSON(http.StatusOK, data)
	})

}
