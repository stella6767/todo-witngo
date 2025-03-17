package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-app/internal/dto"
	errUtil "todo-app/internal/errUtils"
	"todo-app/internal/service"
	"todo-app/internal/view/component"
	"todo-app/internal/view/page"
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
		Size: 7,
	}
	// 쿼리 파라미터 바인딩
	if err := c.ShouldBindQuery(&pageable); err != nil {
		c.Error(errUtil.Wrap(err))
		return
	}
	todos, err := h.service.GetTodosByPage(c, pageable)

	if err != nil {
		c.Error(errUtil.Wrap(err))
		return
	}

	page.Index(todos).Render(c.Request.Context(), c.Writer)
}

func (h *TodoHandler) GetTodosByPage(c *gin.Context) {

	pageable := dto.Pageable{
		Page: 0,
		Size: 7,
	}

	// 쿼리 파라미터 바인딩
	if err := c.ShouldBindQuery(&pageable); err != nil {
		c.Error(errUtil.Wrap(err))
		return
	}
	todos, err := h.service.GetTodosByPage(c, pageable)
	if err != nil {
		c.Error(errUtil.Wrap(err))
		return
	}

	component.TodoContainer(todos).Render(c.Request.Context(), c.Writer)
}

func (h *TodoHandler) UpdateTodoStatus(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Error(errUtil.Wrap(err))
		return
	}
	todo, err := h.service.UpdateTodoStatus(c, int32(id))
	if err != nil {
		c.Error(errUtil.Wrap(err))
	}

	component.TodoComponent(todo).Render(c.Request.Context(), c.Writer)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {

	newTodo := c.PostForm("task")
	if newTodo == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "newTodo is required"})
		return
	}

	todo, err := h.service.CreateTodo(c, newTodo)
	if err != nil {
		c.Error(errUtil.Wrap(err))
		return
	}
	component.TodoComponent(*todo).Render(c.Request.Context(), c.Writer)
}

func (h *TodoHandler) DeleteTodoById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Error(errUtil.Wrap(err))
		return
	}
	err = h.service.DeleteTodoById(c, id)
	if err != nil {
		c.Error(errUtil.Wrap(err))
	}
}
