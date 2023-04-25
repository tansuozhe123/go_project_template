package controller

import (
	"encoding/json"
	"go_project/api/v1/web/dto"
	"go_project/mocks"
	"go_project/pkg/commonres"
	"go_project/pkg/logger"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProductCtroller_GetOneProduct(t *testing.T) {
	logger.InitLogger("test", false, "", "")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mocks.NewMockProductServiceInf(ctrl)
	reqPara := dto.GetOneProductReq{
		ProductKey: "1",
	}
	svc.EXPECT().GetOneProduct(reqPara).Return(&commonres.NormalSucess)
	ctl := ProductCtroller{
		ProductService: svc,
	}
	req := httptest.NewRequest("GET", "/v1/product/one?productkey=1", nil)
	// // 初始化响应
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctl.GetOneProduct(c)
	var resp commonres.CommonRes
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	assert.Equal(t, "success", resp.Type)

}
