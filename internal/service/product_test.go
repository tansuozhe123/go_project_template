package service

import (
	"go_project_template/api/v1/web/dto"
	"go_project_template/internal/biz/entity"
	"go_project_template/mocks"
	"go_project_template/pkg/logger"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProductService_GetOneProduct_Success(t *testing.T) {
	logger.InitLogger("test", false, "", "")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockProductRepoInf(ctrl)
	var pk = "1"
	var req = dto.GetOneProductReq{
		ProductKey: pk,
	}
	var product = entity.Product{}
	repo.EXPECT().GetOneProduct(pk).Return(&product, nil)
	service := &ProductService{
		ProductRepo: repo,
	}
	res := service.GetOneProduct(req)
	assert.Equal(t, "success", res.Type)
}
func TestProductService_GetOneProduct_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockProductRepoInf(ctrl)
	var pk = "1"
	var req = dto.GetOneProductReq{
		ProductKey: pk,
	}
	repo.EXPECT().GetOneProduct(pk).Return(nil, nil)
	service := &ProductService{
		ProductRepo: repo,
	}
	res := service.GetOneProduct(req)
	assert.Equal(t, "error", res.Type)

}
