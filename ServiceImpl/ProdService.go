package ServiceImpl

import (
	"context"
	"grpc_demo/Services"
	"strconv"
)

type ProdService struct {
}

func newProd(id int32, pname string) *Services_pb.ProdModel {
	return &Services_pb.ProdModel{ProdID: id, ProdName: pname}
}

func (p *ProdService) GetProdList(ctx context.Context, in *Services_pb.ProdRequest, res *Services_pb.ProdListResponse) error {
	//time.Sleep(3*time.Second)
	models := make([]*Services_pb.ProdModel, 0)
	var i int32
	for i = 0; i < in.Size; i++ {
		models = append(models, newProd(100+i, "prodname"+strconv.Itoa(100+int(i))))
	}
	res.Data = models
	return nil
}

func (p *ProdService) GetProdDetail(ctx context.Context, in *Services_pb.ProdRequest, res *Services_pb.ProdDetailResponse) error {
	res.Data = newProd(100+in.ProdId, "test prodname"+strconv.Itoa(100+int(in.ProdId)))
	return nil
}

// Server API for ProdService service

//type ProdServiceHandler interface {
//	GetProdList(context.Context, *ProdRequest, *ProdListResponse) error
//}
