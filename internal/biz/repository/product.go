package repository

import (
	"go_project/internal/biz/entity"
)

//go:generate mockgen -destination=../../../mocks/product_repository_mock.go -package=mocks go_project/internal/biz/repository ProductRepoInf

type ProductRepoInf interface {
	GetOneProduct(productkey string) (*entity.Product, error)
}
