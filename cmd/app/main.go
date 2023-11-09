package main

import (
	"githbu.com/niklastreml/htmxual/pkg/handlers"
	"githbu.com/niklastreml/htmxual/pkg/renderer"
	"githbu.com/niklastreml/htmxual/pkg/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.HTMLRender = &renderer.TemplRender{}

	g := r.Group("/count")

	cs := services.NewCounterService()
	handlers.NewCountHandler(*cs).Register(g)
	r.Run()
}
