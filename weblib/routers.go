package weblib

import (
	"github.com/gin-gonic/gin"
	Services_pb "grpc_demo/Services"
)

func NewGinRouter(prodService Services_pb.ProdService) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(InitMiddleware(prodService))
	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", GetProdListHandler)
		v1Group.Handle("GET", "/prod/:pid", GetProdDetailHandler)
	}
	return ginRouter
}
