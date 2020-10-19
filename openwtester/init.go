package openwtester

import (
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
	"github.com/blocktree/truechain-adapter/truechain"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(truechain.Symbol, truechain.NewWalletManager())
}
