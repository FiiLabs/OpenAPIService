package handler

import (
	"fmt"
	"github.com/FiiLabs/OpenAPIService/errors"
	"github.com/FiiLabs/OpenAPIService/libs/pool"
	"github.com/FiiLabs/OpenAPIService/models/req"
	"github.com/FiiLabs/OpenAPIService/response"
	Perm "github.com/bianjieai/iritamod-sdk-go/perm"
	"github.com/gin-gonic/gin"
	"github.com/irisnet/core-sdk-go/types"
	"net/http"
)

func AccountHandler(c *gin.Context)  {
	//name := c.PostForm("name")
	//operation_id := c.PostForm("operation_id")
	var req req.AccountCreateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println("ctx.ShouldBindJSON err: ", err)
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}
	client := pool.GetClient()
	cfg := pool.GetConfig()
	address,mem, err := client.Key.Add(req.Name, cfg.Server.Password)
	if err != nil {
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}
	addrAdmin, err :=client.Key.Show("admin", cfg.Server.Password)
	if err != nil {
		addrAdmin, err = client.Key.Recover("admin", cfg.Server.Password,cfg.Server.Mnemonic)
	}
	if err != nil {
		fmt.Println(fmt.Errorf("导入私钥失败: %s", err.Error()))
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}
	fmt.Println("address:", address)
	fmt.Println("addrAdmin:", addrAdmin)
	baseTx := types.BaseTx{
		From:     "admin",
		Password: cfg.Server.Password,
		Gas:      400000,
		Memo:     "",
		Mode:     types.Sync,
	}
	fee, _ := types.ParseDecCoins("8000000ugas")
	sendResult ,err := client.Bank.Send(address,fee,baseTx)
	if err != nil {
		fmt.Println(fmt.Errorf("send失败: %s", err.Error()))
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	} else {
		fmt.Println("send成功 TxHash：", sendResult.Hash)
	}
	var roles [1]Perm.Role
	roles[0] = Perm.RolePowerUser
	roleResult ,err :=client.Perm.AssignRoles(address, roles[:],baseTx)
	if err != nil {
		fmt.Println(fmt.Errorf("AssignRoles失败: %s", err.Error()))
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	} else {
		fmt.Println("AssignRoles成功 TxHash：", roleResult.Hash)
	}
	datax := map[string]interface{}{
		"account": address,
		"name": req.Name,
		"mnemonic":mem,
		"operation_id": req.OperationId,
	}
	data := map[string]interface{}{
		"data": datax,
	}
	c.JSONP(http.StatusOK, data)
}