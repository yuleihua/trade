package api

import (
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	cmn "github.com/yuleihua/trade/pkg/common"
	"github.com/yuleihua/trade/pkg/server"
	"github.com/yuleihua/trade/transfer/biz"
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
		cid, err := t.Login(&req)
		if err != nil {
			result := errcode.InvalidDatabaseError{Message: "no record by username"}
			c.JSON(200, map[string]interface{}{
				"code": result.Code(),
				"msg":  result.Error(),
			})
		}

		uid := strconv.FormatInt(cid, 10)
		token, err := biz.MakeJWTToken(uid, req.Name)
		if err != nil {
			log.Errorf("make token error, uid:%s, name:%s, %v", uid, req.Name, err)
			result := errcode.InvalidTokenError{Message: "invalid token"}
			c.JSON(200, map[string]interface{}{
				"code": result.Code(),
				"msg":  result.Error(),
			})
			return nil
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
