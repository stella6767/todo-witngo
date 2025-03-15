package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/internal/handler"
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

	registerTodoHandler(todoHandler, router)

	router.GET("/json", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.JSON(http.StatusOK, data)
	})

	return router
}

func registerTodoHandler(handler *handler.TodoHandler, router *gin.Engine) {

	router.GET("/", handler.Index)

	json := router.Group("json")
	json.GET("/todos", handler.Test)
	//json.GET("/todos")
}

//func Router(
//	todoRepo repository.TodoRepository,
//	cfg config.Config,
//) {
//	// Gin 설정
//	router := gin.Default()
//	// 템플릿 등록
//	router.GET("/", func(c *gin.Context) {
//		//todos, _ := todoRepo.GetTodos(c.Request.Context(), 1) // 임시 유저ID
//		c.HTML(http.StatusOK, "", view.Test())
//	})
//
//	// API 라우트
//	api := router.Group(`/v1`)
//	{
//		todoHandler := handler.NewTodoHandler(todoRepo)
//		api.POST("/todos", todoHandler.CreateTodo)
//	}
//
//	// HTMX 에셋 제공
//	router.Static("/static", "./static")
//
//	log.Fatal(router.Run(":" + cfg.Port))
//}
