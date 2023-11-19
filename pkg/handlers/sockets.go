package handlers

import (
	"bytes"
	"fmt"
	"io"

	"githbu.com/niklastreml/htmxual/pkg/components"
	"githbu.com/niklastreml/htmxual/pkg/pages"
	"githbu.com/niklastreml/htmxual/pkg/services"
	"github.com/gin-gonic/gin"
)

func NewSseHandler(ss *services.SseService) *SocketHandler {
	return &SocketHandler{
		ss: ss,
	}
}

type SocketHandler struct {
	ss *services.SseService
}

func (sh *SocketHandler) Register(r *gin.RouterGroup) {
	r.GET("/", sh.View)
	r.GET("/subscribe", sh.Subscribe)
	r.POST("/publish", sh.Publish)
}

func (sh *SocketHandler) View(c *gin.Context) {
	c.HTML(200, "", components.Layout(pages.Socket(sh.ss.Messages())))
}

func (sh *SocketHandler) Subscribe(c *gin.Context) {
	ch := sh.ss.Subscribe()
	defer sh.ss.Unsubscribe(ch)

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-ch; ok {
			fmt.Println(msg.User + ": " + msg.Body)
			b := bytes.NewBuffer([]byte{})
			pages.Chat(sh.ss.Messages()).Render(c, b)

			c.SSEvent("message", b.String())
			return true
		}
		return false
	})

}

func (sh *SocketHandler) Publish(c *gin.Context) {
	c.Request.ParseForm()
	user := c.Request.Form.Get("user")
	body := c.Request.Form.Get("body")

	msg := services.Message{
		User: user,
		Body: body,
	}

	sh.ss.Publish(msg)

	c.Status(201)
}
