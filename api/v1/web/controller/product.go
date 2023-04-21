package controller

import (
	"net/http"

	"go_project_template/api/v1/web/dto"
	"go_project_template/internal/pkg/apires"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ProductCtroller struct {
}

var ProductCtrl *ProductCtroller

func init() {
	ProductCtrl = &ProductCtroller{}
}

// @Summary 测试
// @Tags 产品
// @Description 测试
// @Produce json
// @Param username query string true "用户名"
// @Success 200 {object} 	dtos.ApiRes "成功"
// @Router /product/ [get]
func (ctrl *ProductCtroller) GetProduct(c *gin.Context) {
	var req dto.GetProductReq
	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Infoln(err)
		c.JSON(http.StatusCreated, apires.NormalSucess)
		return
	}
	logrus.Infoln(req.Username)

	//返回结果
	c.JSON(http.StatusOK, apires.NormalSucess)
}
