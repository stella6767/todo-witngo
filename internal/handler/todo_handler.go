package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/internal/db"
	"todo-app/internal/repository"
)

type TodoHandler struct {
	repo repository.TodoRepository
}

func NewTodoHandler(repo repository.TodoRepository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

func TodoList(todos []db.Todo) {

	for i := 0; i < len(todos); i++ {

		todo := todos[i]

		fmt.Print(todo.ID)
	}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {

	userID := c.GetInt("userID")
	var req struct {
		Title string `json:"title" binding:"required"`
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.repo.CreateTodo(c.Request.Context(), int32(userID), req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create todo"})
		return
	}

	c.Render(http.StatusOK, ui.TodoItem(todo))
}
