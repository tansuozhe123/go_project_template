package service

type ProductServiceInf interface {
}
type ProductService struct {
}

var ProductSvc = &ProductService{}

func (svc *ProductService) GetProduct() {
}
