package api

import (
	"context"
	"net/http"
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

func AddTransferHandler(c echo.Context) error {
	uid := biz.CheckJWTToken(c)
	if uid == "" {
		result := errcode.InvalidTokenError{Message: "invalid token"}
		c.JSON(200, map[string]interface{}{
			"code": result.Code(),
			"msg":  result.Error(),
		})
		return nil
	}
	log.Debugf("check token okay, uid: %s", uid)

	var req types.RequestTransfer
	bind := cmn.NewBinder(0)
	if err := bind.Bind(&req, c); err != nil {
		result := errcode.InvalidRequestError{Message: "invalid body"}
		c.JSON(200, map[string]interface{}{
			"code": result.Code(),
			"msg":  result.Error(),
		})
		return nil
	}
	log.Debugf("incoming request: %v", req)

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
		ctx := context.Background()
		response, err := t.AddTransfer(ctx, uid, &req)
		if err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": "1002",
				"msg":  err.Error(),
			})
			return nil
		}

		c.JSON(200, map[string]interface{}{
			"code": 0,
			"data": response,
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

func ConfirmTransferHandler(c echo.Context) error {
	uid := biz.CheckJWTToken(c)
	if uid == "" {
		result := errcode.InvalidTokenError{Message: "invalid token"}
		c.JSON(200, map[string]interface{}{
			"code": result.Code(),
			"msg":  result.Error(),
		})
		return nil
	}
	log.Debugf("check token okay, uid: %s", uid)

	var req types.RequestReceipt
	bind := cmn.NewBinder(0)
	if err := bind.Bind(&req, c); err != nil {
		result := errcode.InvalidRequestError{Message: "invalid body"}
		c.JSON(200, map[string]interface{}{
			"code": result.Code(),
			"msg":  result.Error(),
		})
		return nil
	}
	log.Debugf("incoming request: %v", req)

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
		ctx := context.Background()
		response, err := t.ConfirmTransfer(ctx, uid, &req)
		if err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": "1002",
				"msg":  err.Error(),
			})
			return nil
		}

		c.JSON(200, map[string]interface{}{
			"code": 0,
			"data": response,
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

func TransferDetailHandler(c echo.Context) error {
	uuid := c.Param("uuid")
	toAccount := c.Param("account")
	toName := c.Param("name")
	//position := c.Param("position")
	//pageSize := c.Param("pagesize")

	cid := biz.CheckJWTToken(c)
	uid, err := strconv.ParseInt(cid, 10, 64)
	if err != nil || uid == 0 {
		result := errcode.InvalidTokenError{Message: "invalid token"}
		c.JSON(200, map[string]interface{}{
			"code": result.Code(),
			"msg":  result.Error(),
		})
		return nil
	}
	log.Debugf("check token okay, uid: %s", cid)

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
		ctx := context.Background()
		response, err := t.GetTransfer(ctx, uid, toAccount, toName, uuid)
		if err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": "1002",
				"msg":  err.Error(),
			})
			return nil
		}

		c.JSON(200, map[string]interface{}{
			"code": 0,
			"data": response,
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
