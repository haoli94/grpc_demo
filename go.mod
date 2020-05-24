module grpc_demo

go 1.13

replace github.com/lucas-clemente/quic-go v0.15.2 => github.com/lucas-clemente/quic-go v0.14.1

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

// go get github.com/favadi/protoc-go-inject-tag
// go get github.com/afex/hystrix-go/hystrix 熔断

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/coreos/etcd v3.3.22+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/favadi/protoc-go-inject-tag v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v2.0.1+incompatible // indirect
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/micro/v2 v2.7.0 // indirect
	github.com/miekg/dns v1.1.29 // indirect
	github.com/nats-io/nats.go v1.10.0 // indirect
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2 // indirect
	golang.org/x/sys v0.0.0-20200519105757-fe76b779f299 // indirect
	golang.org/x/tools v0.0.0-20200522201501-cb1345f3a375 // indirect
	google.golang.org/genproto v0.0.0-20200521103424-e9a78aa275b7 // indirect
	google.golang.org/grpc v1.29.1 // indirect
	google.golang.org/protobuf v1.23.0
	gopkg.in/yaml.v2 v2.3.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
)
