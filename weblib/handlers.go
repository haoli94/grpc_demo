package weblib

import (
	"context"
	"github.com/gin-gonic/gin"
	Services_pb "grpc_demo/Services"
	"strconv"
)

func newProd(id int32, pname string) *Services_pb.ProdModel {
	return &Services_pb.ProdModel{ProdID: id, ProdName: pname}
}

// 降级功能尽量简单, 不要去调数据库, 就是在内存里写死的东西或者是redis里的东西, 读取文本文件等
// 定时任务生成静态数据, 业务逻辑尽量简单
func defaultProductsResp() (*Services_pb.ProdListResponse, error) {
	models := make([]*Services_pb.ProdModel, 0)
	var i int32
	for i = 0; i < 5; i++ {
		models = append(models, newProd(100+i, "prodname"+strconv.Itoa(100+int(i))))
	}
	res := &Services_pb.ProdListResponse{}
	res.Data = models
	return res, nil
}

func GetProdDetailHandler(ginCtx *gin.Context) {
	var pr Services_pb.ProdRequest
	micro_prodService := ginCtx.Keys["microProdService"].(Services_pb.ProdService)
	err := ginCtx.BindUri(&pr) //使用自定义的结构体解析post表单
	if err != nil {
		ginCtx.JSON(500, gin.H{
			"status": err.Error(),
		})
	} else {
		prodDetail, err := micro_prodService.GetProdDetail(context.Background(), &pr)
		if err != nil {
			ginCtx.JSON(500, gin.H{
				"status": err.Error(),
			})
		} else {
			ginCtx.JSON(200, gin.H{"data": prodDetail})
		}
	}
}

func GetProdListHandler(ginCtx *gin.Context) {
	var pr Services_pb.ProdRequest
	micro_prodService := ginCtx.Keys["microProdService"].(Services_pb.ProdService)
	err := ginCtx.Bind(&pr) //使用自定义的结构体解析post表单
	if err != nil {
		ginCtx.JSON(500, gin.H{
			"status": err.Error(),
		})
	} else {

		/**
		// 熔断代码
		// 1. 配置config
		configA := hystrix.CommandConfig{
			Timeout: 1200,
		}
		// 2. 配置command
		hystrix.ConfigureCommand("GetProds", configA)
		// 3. 执行, 使用do, 调用上面的config, 同步
		var prodList *Services_pb.ProdListResponse
		err := hystrix.Do("GetProds", func() error {
			prodList, err = micro_prodService.GetProdList(context.Background(), &pr)
			return err
		}, func(err error) error {
			prodList, err = defaultProductsResp()
			return err
		})
		// fallback是降级(服务)函数
		*/

		prodList, err := micro_prodService.GetProdList(context.Background(), &pr)
		if err != nil {
			ginCtx.JSON(500, gin.H{
				"status": err.Error(),
			})
		} else {
			ginCtx.JSON(200, gin.H{"data": prodList})
		}
	}
}
