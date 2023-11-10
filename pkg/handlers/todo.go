package handlers

import (
	"strconv"

	"githbu.com/niklastreml/htmxual/pkg/components"
	"githbu.com/niklastreml/htmxual/pkg/pages"
	"githbu.com/niklastreml/htmxual/pkg/services"
	"github.com/gin-gonic/gin"
)

var _ Handler = (*TodoHandler)(nil)

type TodoHandler struct {
	ts *services.TodoService
}

func NewTodoHandler(ts *services.TodoService) *TodoHandler {
	return &TodoHandler{
		ts: ts,
	}
}

func (t *TodoHandler) Register(r *gin.RouterGroup) {
	r.GET("/", t.View)
	r.POST("/", t.Add)
	r.PUT("/:id", t.Update)
	r.DELETE("/:id", t.Remove)
}

func (t *TodoHandler) View(c *gin.Context) {
	todos := t.ts.List()
	c.HTML(200, "", components.Layout(pages.TodoList(todos)))
}
func (t *TodoHandler) Add(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.AbortWithError(400, err)
		return
	}

	title := c.Request.FormValue("title")
	todo := services.Todo{
		Title: title,
	}

	id := t.ts.Add(todo)
	todo.Id = id
	c.HTML(200, "", pages.TodoItem(&todo))
}

func (t *TodoHandler) Update(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.AbortWithError(400, err)
		return
	}

	id := c.Param("id")
	title := c.Request.FormValue("title")
	done := c.Request.FormValue("done")
	doneBool := false

	idint, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	if done == "on" {
		doneBool = true
	}

	todo := t.ts.Get(idint)

	todo.Title = title
	todo.Done = doneBool

	c.HTML(200, "", pages.TodoItem(todo))

}
func (t *TodoHandler) Remove(c *gin.Context) {
	id := c.Param("id")

	idint, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	t.ts.Remove(idint)

	c.Status(200)
}
