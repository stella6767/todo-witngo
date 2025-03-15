package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"todo-app/internal/repository"
)

type TodoHandler struct {
	repo repository.TodoRepository
}

func NewTodoHandler(repo repository.TodoRepository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {

	//아 개불편하네..
	p := pgtype.Bool{Bool: false}

	value, err := p.Value()
	if err != nil {
		return
	}
	fmt.Print(value)

	//userID := c.GetInt("userID")
	//var req struct {
	//	Title string `json:"title" binding:"required"`
	//}
	//
	//if err := c.ShouldBind(&req); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//todo, err := h.repo.CreateTodo(c.Request.Context(), int32(userID), req.Title)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create todo"})
	//	return
	//}

	//c.Render(http.StatusOK, view.TodoList(todo))
}
