package handler

import (
	"fmt"
	"github.com/FiiLabs/OpenAPIService/errors"
	"github.com/FiiLabs/OpenAPIService/libs/pool"
	"github.com/FiiLabs/OpenAPIService/models/req"
	"github.com/FiiLabs/OpenAPIService/response"
	"github.com/gin-gonic/gin"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/irismod-sdk-go/nft"
	"net/http"
)

func NFTHandler(c *gin.Context) {
	var req req.NFTReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println("ctx.ShouldBindJSON err: ", err)
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}
	client := pool.GetClient()
	cfg := pool.GetConfig()
	baseTx := types.BaseTx{
		From:     req.UsrName,
		Password: cfg.Server.Password,
		Gas:      400000,
		Memo:     "",
		Mode:     types.Sync,
	}
	nftResult, err := client.NFT.MintNFT(nft.MintNFTRequest{
		Denom:req.Denom,
		ID:req.NFTId,
		Name:req.NFTName,
		URI:req.Uri,
		Data:req.Data,
		URIHash:req.UriHash,
		Recipient: req.Recipient,
	}, baseTx)
	if err != nil {
		fmt.Println(fmt.Errorf("NFT 创建失败: %s", err.Error()))
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	} else {
		fmt.Println("NFT 创建成功 TxHash：", nftResult.Hash)
	}
	data := map[string]interface{}{
		"hash":nftResult.Hash,
		"operation_id": req.OpId,
	}
	c.JSONP(http.StatusOK, data)
}

func NFTTransferHandler(c *gin.Context) {
	var req req.NFTTrfReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println("ctx.ShouldBindJSON err: ", err)
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}
	client := pool.GetClient()
	cfg := pool.GetConfig()
	baseTx := types.BaseTx{
		From:     req.UsrName,
		Password: cfg.Server.Password,
		Gas:      400000,
		Memo:     "",
		Mode:     types.Sync,
	}
	oldNFT,err:=client.NFT.QueryNFT(req.ClsId,req.NFTId)
	if err != nil {
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}
	nftResult, err := client.NFT.TransferNFT(nft.TransferNFTRequest{
		Denom:req.ClsId,
		ID:req.NFTId,
		URI:oldNFT.URI,
		Data:oldNFT.Data,
		Name:oldNFT.Name,
		Recipient:req.Recipient,
		URIHash:oldNFT.URIHash,
	}, baseTx)
	if err != nil {
		fmt.Println(fmt.Errorf("NFT 转移失败: %s", err.Error()))
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	} else {
		fmt.Println("NFT 转移成功 TxHash：", nftResult.Hash)
	}
	data := map[string]interface{}{
		"hash":nftResult.Hash,
		"operation_id": req.OpId,
	}
	c.JSONP(http.StatusOK, data)
}


func NFTEditHandler(c *gin.Context) {
	var req req.NFTEditReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println("ctx.ShouldBindJSON err: ", err)
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}
	client := pool.GetClient()
	cfg := pool.GetConfig()
	baseTx := types.BaseTx{
		From:     req.UsrName,
		Password: cfg.Server.Password,
		Gas:      400000,
		Memo:     "",
		Mode:     types.Sync,
	}
	_,err1:=client.NFT.QueryNFT(req.Denom,req.NFTId)
	if err1 != nil {
		e := errors.Wrap(err1)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}
	nftResult, err := client.NFT.EditNFT(nft.EditNFTRequest{
		Denom:req.Denom,
		ID:req.NFTId,
		URI:req.Uri,
		Data:req.Data,
		Name:req.NFTName,
		URIHash:req.UriHash,
	}, baseTx)
	if err != nil {
		fmt.Println(fmt.Errorf("NFT 编辑失败: %s", err.Error()))
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	} else {
		fmt.Println("NFT 编辑成功 TxHash：", nftResult.Hash)
	}
	data := map[string]interface{}{
		"hash":nftResult.Hash,
		"operation_id": req.OpId,
	}
	c.JSONP(http.StatusOK, data)
}
func NFTDeleteHandler(c *gin.Context) {
	var req req.NFTDelReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println("ctx.ShouldBindJSON err: ", err)
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}
	client := pool.GetClient()
	cfg := pool.GetConfig()
	baseTx := types.BaseTx{
		From:     req.UsrName,
		Password: cfg.Server.Password,
		Gas:      400000,
		Memo:     "",
		Mode:     types.Sync,
	}
	_,err1:=client.NFT.QueryNFT(req.ClsId,req.NFTId)
	if err1 != nil {
		e := errors.Wrap(err1)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}
	nftResult, err := client.NFT.BurnNFT(nft.BurnNFTRequest{
		Denom:req.ClsId,
		ID:req.NFTId,
	}, baseTx)
	if err != nil {
		fmt.Println(fmt.Errorf("NFT 删除失败: %s", err.Error()))
		e := errors.Wrap(err)
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	} else {
		fmt.Println("NFT 删除成功 TxHash：", nftResult.Hash)
	}
	data := map[string]interface{}{
		"hash":nftResult.Hash,
		"operation_id": req.OpId,
	}
	c.JSONP(http.StatusOK, data)
}
