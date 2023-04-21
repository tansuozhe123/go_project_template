package biz

type Product struct {
}

type ProductRepo interface {
	GetProductByIds(ids []string) ([]Product, error)
}
