package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"todo-app/internal/dto"
	"todo-app/internal/service"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) Test(c *gin.Context) {
	c.String(http.StatusOK, "Test")
}

func (h *TodoHandler) Index(c *gin.Context) {

	pageable := dto.Pageable{
		Page: 0,
		Size: 10,
	}
	// 쿼리 파라미터 바인딩
	if err := c.ShouldBindQuery(&pageable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todos := h.service.GetTodosByPage(c, pageable)
	log.Println(todos)
	//view.Index(todos).Render(c.Request.Context(), c.Writer)

}

//func (h *TodoHandler) CreateTodo(c *gin.Context) {
//	userID := c.GetInt("userID")
//	var req struct {
//		Title string `json:"title" binding:"required"`
//	}
//
//	//c.post
//
//	if err := c.ShouldBind(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	todo, err := h.service.Create(c.Request.Context(), int32(userID), req.Title)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create todo"})
//		return
//	}
//
//	c.Render(http.StatusOK, view.TodoList(todo))
//}
