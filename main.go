package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"grpc_demo/ServiceImpl"
	"grpc_demo/Services"
	"log"
)

//wrapper
type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	// 上面是我们加的
	// wrapper 调用父类方法
	return l.Client.Call(ctx, req, rsp)
}

func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)

	service := micro.NewService(
		micro.Name("microProdService"),
		micro.Address(":8000"),
		micro.Registry(consulReg), //注册到服务器上的consul中
		micro.WrapClient(logWrap),
		//micro.WrapClient(Wrapper.NewProdsWrapper),
	)
	service.Init() //加了这句就可以使用命令行的形式去设置我们一些启动的配置
	// Server API for ProdService service
	_ = Services_pb.RegisterProdServiceHandler(service.Server(), new(ServiceImpl.ProdService))
	//type ProdServiceHandler interface {
	//	GetProdList(context.Context, *ProdRequest, *ProdListResponse) error
	//}
	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
