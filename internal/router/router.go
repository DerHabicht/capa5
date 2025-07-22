package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Router struct {
	engine *gin.Engine
	prefix *gin.RouterGroup
}

func NewRouter(prefix string) *Router {
	engine := gin.Default()
	p := engine.Group(prefix)

	return &Router{
		engine: engine,
		prefix: p,
	}
}

func (r *Router) Run() error {
	err := r.engine.Run(":8080")

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Router) RegisterHandler(m string, addr string, h gin.HandlerFunc) {
	r.prefix.Handle(m, addr, h)
}

func (r *Router) RegisterRestHandler(h RestHandler, keyParam string) {
	g := h.Group(r.prefix)

	urlByKey := fmt.Sprintf("/:%s", keyParam)

	g.OPTIONS("/", h.Options)
	g.HEAD("/", h.Head)
	g.POST("/", h.Create)
	g.GET("/", h.List)
	g.GET(urlByKey, h.Fetch)
	g.PUT(urlByKey, h.Replace)
	g.PATCH(urlByKey, h.Update)
	g.DELETE(urlByKey, h.Delete)
}
