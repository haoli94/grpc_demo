# shellcheck disable=SC2164
cd Services/protos
protoc --micro_out=../ --go_out=../ Models.proto
  protoc --micro_out=../ --go_out=../ ProdService.proto
#protoc-go-inject-tag 对json转换传输的变换字段名声明
protoc-go-inject-tag -input=../ProdService.pb.go
cd .. && cd ..