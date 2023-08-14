package nft

import (
	"server/api/logic"
	"server/api/schemas/resp"
)

var NftService = nftService{}

// walletService NFT服务类
type nftService struct{}

// MintSeaDrop 铸造NFT
func (ws nftService) MintSeaDrop(minter string, quantity uint) resp.NftTxResp {
	tx := logic.NftLogic.MintSeaDrop(minter, quantity)
	return resp.NftTxResp{
		Tx: tx,
	}
}

// TransferFrom NFT转账
func (ws nftService) TransferFrom(from string, to string, tokenId uint) resp.NftTxResp {
	tx := logic.NftLogic.TransferFrom(from, to, tokenId)
	return resp.NftTxResp{
		Tx: tx,
	}
}
