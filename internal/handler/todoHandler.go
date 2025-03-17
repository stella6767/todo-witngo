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
		Size: 10,
	}
	// 쿼리 파라미터 바인딩
	if err := c.ShouldBindQuery(&pageable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errUtils": err.Error()})
		return
	}

	todos := h.service.GetTodosByPage(c, pageable)
	page.Index(todos).Render(c.Request.Context(), c.Writer)
}

func (h *TodoHandler) UpdateTodoStatus(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.Error(errUtil.Wrap(err))
		return
	}

	todo, err := h.service.UpdateTodoStatus(c, int32(id))

	component.TodoComponent(todo).Render(c.Request.Context(), c.Writer)
}

//func (h *TodoHandler) CreateTodo(c *gin.Context) {
//	userID := c.GetInt("userID")
//	var req struct {
//		Title string `json:"title" binding:"required"`
//	}
//
//	//c.post
//
//	if errUtils := c.ShouldBind(&req); errUtils != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"errUtils": errUtils.Error()})
//		return
//	}
//
//	todo, errUtils := h.service.Create(c.Request.Context(), int32(userID), req.Title)
//	if errUtils != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"errUtils": "failed to create todo"})
//		return
//	}
//
//	c.Render(http.StatusOK, view.TodoList(todo))
//}
