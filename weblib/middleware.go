package weblib

import (
	"github.com/gin-gonic/gin"
	Services_pb "grpc_demo/Services"
)

func InitMiddleware(prodService Services_pb.ProdService) gin.HandlerFunc{
	return func(ctx *gin.Context){
		ctx.Keys = make(map[string]interface{})
		ctx.Keys["microProdService"] = prodService
		ctx.Next()
	}
}
