package service

//go:generate mockgen -destination=../../mocks/product_service_mock.go -package=mocks go_project_template/internal/service ProductServiceInf

import (
	"go_project_template/api/v1/web/dto"
	"go_project_template/internal/biz/repository"
	"go_project_template/internal/data/persistence"
	"go_project_template/internal/pkg/apires"
	"go_project_template/pkg/commonres"
)

type ProductServiceInf interface {
	GetOneProduct(req dto.GetOneProductReq) *commonres.CommonRes
}
type ProductService struct {
	ProductRepo repository.ProductRepoInf
}

var ProductSvc = &ProductService{
	ProductRepo: persistence.ProductPersi,
}

func (svc *ProductService) GetOneProduct(req dto.GetOneProductReq) *commonres.CommonRes {
	var res = commonres.NormalSucess
	if req.ProductKey == "" {
		return &commonres.NullParameter
	}
	//查找产品
	product, err := svc.ProductRepo.GetOneProduct(req.ProductKey)
	if err != nil {
		return &commonres.InternalServerError
	}
	//产品不存在
	if product == nil {
		return &apires.NoProduct
	}
	res.Data = product
	return &res
}
