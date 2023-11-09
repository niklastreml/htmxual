package handlers

import (
	"fmt"
	"strconv"

	"githbu.com/niklastreml/htmxual/pkg/components"
	"githbu.com/niklastreml/htmxual/pkg/pages"
	"githbu.com/niklastreml/htmxual/pkg/services"
	"github.com/gin-gonic/gin"
)

var _ Handler = (*CountHandler)(nil)

type CountHandler struct {
	cs *services.CounterService
}

func NewCountHandler(cs services.CounterService) *CountHandler {
	return &CountHandler{
		cs: &cs,
	}
}

func (ch *CountHandler) Register(r *gin.RouterGroup) {
	r.PUT("/increment", ch.Increment)
	r.PUT("/decrement", ch.Decrement)
	r.GET("/", ch.View)
}

func (ch *CountHandler) Increment(c *gin.Context) {
	ch.cs.Increment()

	c.String(200, strconv.Itoa(ch.cs.Count()))
}

func (ch *CountHandler) Decrement(c *gin.Context) {
	ch.cs.Decrement()
	c.String(200, strconv.Itoa(ch.cs.Count()))
}

func (ch *CountHandler) View(c *gin.Context) {

	fmt.Println("CountHandler.View")
	page := components.Layout(pages.Counter(ch.cs.Count()))
	c.HTML(200, "", page)
	c.Status(200)

}
