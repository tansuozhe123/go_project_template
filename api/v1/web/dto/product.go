package dto

type GetOneProductReq struct {
	ProductKey string `form:"productkey" binding:"required"`
}
