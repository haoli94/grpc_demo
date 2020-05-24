package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	Services_pb "grpc_demo/Services"
	"grpc_demo/weblib"
)

func main() {

	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)

	micro_prodService := Services_pb.NewProdService("microProdService", micro.NewService(micro.Name("microProdService.client")).Client())
	httpServer := web.NewService(
		web.Name("httpProdService"),
		web.Address(":8001"),
		web.Handler(weblib.NewGinRouter(micro_prodService)),
		web.Registry(consulReg), //注册到服务器上的consul中
	)

	_ = httpServer.Init()
	_ = httpServer.Run()
}
