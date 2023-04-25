package controller

import (
	"net/http"

	"go_project_template/api/v1/web/dto"
	"go_project_template/internal/service"
	"go_project_template/pkg/commonres"
	"go_project_template/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductCtroller struct {
	ProductService service.ProductServiceInf
}

var ProductCtrl *ProductCtroller

func init() {
	ProductCtrl = &ProductCtroller{
		ProductService: service.ProductSvc,
	}

}

// @Summary 获取单个产品
// @Tags 产品
// @Description 获取单个产品
// @Produce json
// @Param username query string true "用户名"
// @Success 200 {object} 	dtos.ApiRes "成功"
// @Router /product/one [get]
func (ctrl *ProductCtroller) GetOneProduct(c *gin.Context) {
	var req dto.GetOneProductReq
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Logger.Error("参数错误", zap.Error(err))
		c.JSON(http.StatusCreated, commonres.ParameterFormattedError)
		return
	}
	logger.Logger.Info("参数", zap.Any("req", req))
	//调用service
	res := ctrl.ProductService.GetOneProduct(req)

	//返回结果
	c.JSON(http.StatusOK, res)

}
