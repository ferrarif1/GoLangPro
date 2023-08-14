package routers

import (
	"server/api/schemas/req"
	"server/api/service/nft"
	"server/core"
	"server/core/response"
	"server/util"

	"github.com/gin-gonic/gin"
)

var NftGroup = core.Group("/nft")

func init() {
	group := NftGroup
	group.AddPOST("/mintSeaDrop", mintSeaDrop)
	group.AddPOST("/transferFrom", transferFrom)
}

// 铸造NFT
// 参数：minter: 接收者, quantity: 数量
func mintSeaDrop(c *gin.Context) {
	var mintReq req.NftMintSeaDropReq
	util.VerifyUtil.VerifyBody(c, &mintReq)
	response.OkWithData(c, nft.NftService.MintSeaDrop(mintReq.Minter, mintReq.Quantity))
}

// NFT转账
// 参数：from: 发送人, to: 接收人, tokenId: NFT ID
func transferFrom(c *gin.Context) {
	var transferFromReq req.NftTransferFromReq
	util.VerifyUtil.VerifyBody(c, &transferFromReq)
	response.OkWithData(c, nft.NftService.TransferFrom(transferFromReq.From, transferFromReq.To, transferFromReq.TokenId))
}
