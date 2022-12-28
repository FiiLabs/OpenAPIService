package handler

import (
	"fmt"
	"github.com/FiiLabs/OpenAPIService/config"
	"github.com/FiiLabs/OpenAPIService/errors"
	"github.com/FiiLabs/OpenAPIService/models/req"
	"github.com/FiiLabs/OpenAPIService/response"
	"github.com/gin-gonic/gin"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/irismod-sdk-go/nft"
	"net/http"
)



func NFTClassHandler(c *gin.Context) {
	var req req.NFTClsReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println("ctx.ShouldBindJSON err: ", err)
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}
	client := config.GetConfigClient()
	cfg := config.GetConfig()
	baseTx := types.BaseTx{
		From:     req.UsrName,
		Password: cfg.Server.Password,
		Gas:      400000,
		Memo:     "",
		Mode:     types.Sync,
	}
	nftResult, err := client.NFT.IssueDenom(nft.IssueDenomRequest{
		ID: req.ClsId,
		Name:req.ClsName,
		Schema:req.Schema,
		Symbol:req.Symbol,
		Description:req.Description,
		Uri:req.Uri,
		UriHash:req.UriHash,
		Data:req.Data,
	}, baseTx)
	if err != nil {
		fmt.Println(fmt.Errorf("NFT 类别创建失败: %s", err.Error()))
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	} else {
		fmt.Println("NFT 类别创建成功 TxHash：", nftResult.Hash)
	}
	data := map[string]interface{}{
		"hash":nftResult.Hash,
		"operation_id": req.OpId,
	}
	c.JSONP(http.StatusOK, data)
}

func NFTClassTransferHandler(c *gin.Context) {
	var req req.NFTClsTrfReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println("ctx.ShouldBindJSON err: ", err)
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}
	client := config.GetConfigClient()
	cfg := config.GetConfig()
	baseTx := types.BaseTx{
		From:     req.UsrName,
		Password: cfg.Server.Password,
		Gas:      400000,
		Memo:     "",
		Mode:     types.Sync,
	}
	nftResult, err := client.NFT.TransferDenom(nft.TransferDenomRequest{
		ID: req.ClsId,
		Recipient:req.Recipient,
	}, baseTx)
	if err != nil {
		fmt.Println(fmt.Errorf("NFT 类别转移失败: %s", err.Error()))
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	} else {
		fmt.Println("NFT 类别转移成功 TxHash：", nftResult.Hash)
	}
	data := map[string]interface{}{
		"hash":nftResult.Hash,
		"operation_id": req.OpId,
	}
	c.JSONP(http.StatusOK, data)
}

