package persistence

import (
	"go_project_template/internal/biz"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) biz.ProductRepo {
	return &ProductRepo{
		db: db,
	}
}
func (PR *ProductRepo) GetProductByIds(ids []string) ([]biz.Product, error) {
	PR.db.Where("id in (?)", ids).Find(&biz.Product{})
	return []biz.Product{}, nil
}
