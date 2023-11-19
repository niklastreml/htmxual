package handlers

import (
	"fmt"
	"time"

	"githbu.com/niklastreml/htmxual/pkg/components"
	"githbu.com/niklastreml/htmxual/pkg/pages"
	"githbu.com/niklastreml/htmxual/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func NewBasicAuthHandler(bas *services.BasicAuth) *BasicAuthHandler {
	return &BasicAuthHandler{
		bas: bas,
	}
}

type BasicAuthHandler struct {
	bas *services.BasicAuth
}

func (bah *BasicAuthHandler) Register(r *gin.RouterGroup) {
	authgroup := r.Group("/", bah.Middleware)
	authgroup.GET("/", bah.View)
	authgroup.POST("/logout", bah.Logout)

	r.POST("/login", bah.Login)
}

func (bah *BasicAuthHandler) View(c *gin.Context) {
	c.HTML(200, "", components.Layout(pages.Secure()))
}

func (bah *BasicAuthHandler) Login(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.AbortWithStatus(500)
		return
	}

	usr := c.Request.Form.Get("user")
	pass := c.Request.Form.Get("pass")
	isAuth := bah.bas.Users()[usr] == pass

	if isAuth && usr != "" && pass != "" {
		token, err := newSignedJwt(usr)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.SetCookie("token", token, 3600, "/", "", false, true)

		c.Redirect(302, "/basic")
		return
	}
	c.HTML(401, "", components.Layout(pages.Login()))
}

func (bah *BasicAuthHandler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Header("HX-Redirect", "/basic")
}

func (bah *BasicAuthHandler) Middleware(c *gin.Context) {
	// grab the users cookie
	token, err := c.Cookie("token")
	if err != nil {
		c.HTML(401, "", components.Layout(pages.Login()))
		c.AbortWithStatus(401)
		return
	}

	if token == "" {
		c.HTML(401, "", components.Layout(pages.Login()))
		c.AbortWithStatus(401)
	}

	_, err = validateSignedJwt(token)

	if err != nil {
		c.HTML(401, "", components.Layout(pages.Login()))
		c.AbortWithStatus(401)
	}

	c.Next()
}

func newSignedJwt(usr string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usr": usr,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	return token.SignedString([]byte("secret"))
}

func validateSignedJwt(tokenstring string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	return token, err
}
