/*
 * Copyright 2018 The OpenWallet Authors
 * This file is part of the OpenWallet library.
 *
 * The OpenWallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OpenWallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package truechain

import (
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/quorum-adapter/quorum"
)

const (
	Symbol = "TRUE"
)

type WalletManager struct {
	*quorum.WalletManager
}

func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.WalletManager = quorum.NewWalletManager()
	wm.Config = quorum.NewConfig(Symbol)
	wm.Log = log.NewOWLogger(wm.Symbol())
	return &wm
}

//FullName 币种全名
func (wm *WalletManager) FullName() string {
	return "True"
}
