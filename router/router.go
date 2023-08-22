package router

import (
	"github.com/gin-gonic/gin"
	"initCake/pkg/aop"
	"initCake/pkg/ctx"
	"initCake/pkg/httpx"
)

type Router struct {
	HTTP httpx.Config
	Ctx  *ctx.Context
}

func New(httpConfig httpx.Config, ctx *ctx.Context) *Router {
	return &Router{
		HTTP: httpConfig,
		Ctx:  ctx,
	}
}

func (rt *Router) Config(r *gin.Engine) {
	r.Use(aop.Recovery())
	pagesPrefix := "/api/v1"
	pages := r.Group(pagesPrefix)
	{
		pages.GET("ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		//pages.GET("app/all", rt.applicationGetAll)
	}

}
