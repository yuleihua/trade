package api

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	cmn "github.com/yuleihua/trade/pkg/common"
	"github.com/yuleihua/trade/pkg/server"
	"github.com/yuleihua/trade/transfer/errcode"
	"github.com/yuleihua/trade/transfer/service"
	"github.com/yuleihua/trade/transfer/types"
)

func CustomerLoginHandler(c echo.Context) error {
	var req types.RequestLogin

	bind := cmn.NewBinder(0)
	if err := bind.Bind(&req, c); err != nil {
		log.Errorf("bind error, %#v, %v", c.Request(), err)
		result := errcode.InvalidRequestError{Message: "invalid body"}
		c.JSON(200, map[string]interface{}{
			"code": result.Code(),
			"msg":  result.Error(),
		})
		return nil
	}

	transfer := server.GetService(types.ServiceName)
	if transfer == nil {
		result := errcode.InvalidRequestError{Message: "no service"}
		c.JSON(200, map[string]interface{}{
			"code": result.Code(),
			"msg":  result.Error(),
		})
		return nil
	}

	if t, ok := transfer.(*service.Transfer); ok {
		token, err := t.Login(&req)
		if err != nil {
			result := errcode.InvalidInternalError{Message: "internal error"}
			c.JSON(200, map[string]interface{}{
				"code": result.Code(),
				"msg":  result.Error(),
			})
		}

		c.JSON(200, map[string]interface{}{
			"code": 0,
			"data": token,
		})
		return nil
	}

	result := errcode.InvalidRequestError{Message: "invalid service"}
	c.JSON(200, map[string]interface{}{
		"code": result.Code(),
		"msg":  result.Error(),
	})

	return nil
}
