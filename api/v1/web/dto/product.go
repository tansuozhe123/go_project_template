package dto

type GetProductReq struct {
	Username string `form:"username" binding:"required"`
	Ids      []int  `form:"ids" binding:"required"`
}
