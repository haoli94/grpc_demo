package Wrapper

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
)

//wrapper
type ProdsWrapper struct {
	client.Client
}

func (p *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	// 熔断代码
	// 1. 配置config
	configName := req.Service()+"."+req.Endpoint()
	configA := hystrix.CommandConfig{
		Timeout: 100,
	}
	// 2. 配置command
	hystrix.ConfigureCommand(configName, configA)
	// 3. 执行, 使用do, 调用上面的config, 同步
	return hystrix.Do(configName, func() error {
		// 上面是我们加的
		// wrapper 调用父类方法
		return p.Client.Call(ctx, req, rsp)
	}, nil)
}

func NewProdsWrapper(client client.Client) client.Client {
	return &ProdsWrapper{client}
}
