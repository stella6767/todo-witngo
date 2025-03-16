package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/internal/service"
	"todo-app/internal/view"
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
	//todos, _ := h.service.GetAllTodos(c.Request.Context())
	//view.Layout("Todos", view.TodoList(todos)).Render(c.Request.Context(), c.Writer)
	//c.HTML(http.StatusOK, "index", view.Index())

	//component.TodoComponent()

	view.Index().Render(c.Request.Context(), c.Writer)
}

//func (h *TodoHandler) CreateTodo(c *gin.Context) {
//	userID := c.GetInt("userID")
//	var req struct {
//		Title string `json:"title" binding:"required"`
//	}
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
//
//	c.Render(http.StatusOK, view.TodoList(todo))
//}
