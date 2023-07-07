package di

import (
	"example.com/gin_wire/app/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	sampleController *controllers.SampleController
}

func NewRouter(
	sampleController *controllers.SampleController,
) *Router {
	return &Router{
		sampleController: sampleController,
	}
}

func (r *Router) GetEngine() *gin.Engine {
	engine := gin.Default()
	engine.GET("/sample", func(c *gin.Context){
		r.sampleController.Get(c)
	})
	return engine
}
