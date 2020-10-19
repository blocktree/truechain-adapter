/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"github.com/blocktree/openwallet/v2/openw"
	"testing"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
)

func testGetAssetsAccountBalance(tm *openw.WalletManager, walletID, accountID string) {
	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func testGetAssetsAccountTokenBalance(tm *openw.WalletManager, walletID, accountID string, contract openwallet.SmartContract) {
	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("token balance:", balance.Balance)
}

func testCreateTransactionStep(tm *openw.WalletManager, walletID, accountID, to, amount, feeRate string, contract *openwallet.SmartContract, extParam map[string]interface{}) (*openwallet.RawTransaction, error) {

	//err := tm.RefreshAssetsAccountBalance(testApp, accountID)
	//if err != nil {
	//	log.Error("RefreshAssetsAccountBalance failed, unexpected error:", err)
	//	return nil, err
	//}

	rawTx, err := tm.CreateTransaction(testApp, walletID, accountID, amount, to, feeRate, "", contract, extParam)

	if err != nil {
		log.Error("CreateTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTx, nil
}

func testCreateSummaryTransactionStep(
	tm *openw.WalletManager,
	walletID, accountID, summaryAddress, minTransfer, retainedBalance, feeRate string,
	start, limit int,
	contract *openwallet.SmartContract,
	feeSupportAccount *openwallet.FeesSupportAccount) ([]*openwallet.RawTransactionWithError, error) {

	rawTxArray, err := tm.CreateSummaryRawTransactionWithError(testApp, walletID, accountID, summaryAddress, minTransfer,
		retainedBalance, feeRate, start, limit, contract, feeSupportAccount)

	if err != nil {
		log.Error("CreateSummaryTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTxArray, nil
}

func testSignTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	_, err := tm.SignTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, "12345678", rawTx)
	if err != nil {
		log.Error("SignTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testVerifyTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	//log.Info("rawTx.Signatures:", rawTx.Signatures)

	_, err := tm.VerifyTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("VerifyTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testSubmitTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	tx, err := tm.SubmitTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("SubmitTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Std.Info("tx: %+v", tx)
	//log.Info("wxID:", tx.WxID)
	log.Info("txID:", rawTx.TxID)

	return rawTx, nil
}

func TestTransfer_TRUE(t *testing.T) {

	addrs := []string{
		//"0x1fc3a35dc5d6e02687515a723f1eb8ab7fada85d",
		//"0x227185a1068c0484c32a88174dba128c63cb8a0d",
		//"0x72cb4302c995edca324578b9a074e4ebf16d2eec",
		//"0x90c472898053e3b172e9fe7cee56adca3be6051a",
		//"0x983a8de787c408266921e9724f3a5a3ff5e26c70",
		//"0xb58cb5c64b9966cf84b4c8231ecf7d73ed4a225e",

		"0x09267e01f35c39f142ebc504bb4c4bcf617b5fef",	//fee
	}

	tm := testInitWalletManager()
	walletID := "WCBkGX2YgKuZndqhZxEnuJMWm9w95UcMNx"
	accountID := "FEWHN8m8Mwey64KenQvVuZGYjErD1Sm7krsqirFe63uC"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	//extParam := map[string]interface{}{
	//	"nonce": 20,
	//}

	for _, to := range addrs {
		rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "0.03", "", nil, nil)
		if err != nil {
			return
		}

		log.Std.Info("rawTx: %+v", rawTx)

		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}
	}
}

func TestTransfer_TRC20(t *testing.T) {

	addrs := []string{
		"0x1fc3a35dc5d6e02687515a723f1eb8ab7fada85d",
		"0x227185a1068c0484c32a88174dba128c63cb8a0d",
		"0x72cb4302c995edca324578b9a074e4ebf16d2eec",
		"0x90c472898053e3b172e9fe7cee56adca3be6051a",
		"0x983a8de787c408266921e9724f3a5a3ff5e26c70",
		"0xb58cb5c64b9966cf84b4c8231ecf7d73ed4a225e",
	}

	tm := testInitWalletManager()
	walletID := "WCBkGX2YgKuZndqhZxEnuJMWm9w95UcMNx"
	accountID := "FEWHN8m8Mwey64KenQvVuZGYjErD1Sm7krsqirFe63uC"

	contract := openwallet.SmartContract{
		Address:  "",
		Symbol:   "TRUE",
		Name:     "",
		Token:    "",
		Decimals: 18,
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	for _, to := range addrs {
		rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "0.00001", "", &contract, nil)
		if err != nil {
			return
		}

		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}
	}
}

func TestSummary_TRUE(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WBGYxZ6yEX582Mx8mGvygXevdLVc7NQnLM"
	accountID := "9EfTQiMEaKSMd1CjxMXRMMxukrwckxdBZpiEkS2B3avD"
	//accountID := "CxE3ds4JdTHXV1f2xSsE6qahgfReKR9iPmFPcBmTfaKP"
	//summaryAddress := "0xbb0d592280f170069821bcae6c861a5686b77c43"
	summaryAddress := "0xd35f9ea14d063af9b3567064fab567275b09f03d"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "0.000000002",
		0, 100, nil, nil)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTxWithErr := range rawTxArray {

		if rawTxWithErr.Error != nil {
			log.Error(rawTxWithErr.Error.Error())
			continue
		}

		_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		//_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
		//if err != nil {
		//	return
		//}
	}

}

func TestSummary_TRC20(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WBGYxZ6yEX582Mx8mGvygXevdLVc7NQnLM"
	accountID := "A1QQ8mLa2uGJhboZJjs1qkEt6zEGrWewfEfmnxs3tYhr"
	summaryAddress := "0x301088aa99ce02fc51887920d14749316b0644c5"

	feesSupport := openwallet.FeesSupportAccount{
		AccountID: "HGwLhPQvU1at7BUiHBf1Kss1bDboX1TzKbA54CK5W3H",
		FeesSupportScale: "2",
	}

	contract := openwallet.SmartContract{
		Address:  "",
		Symbol:   "TRUE",
		Name:     "",
		Token:    "",
		Decimals: 18,
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	list, err := tm.GetAddressList(testApp, walletID, accountID, 0, -1, false)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}

	addressLimit := 2

	//分页汇总交易
	for i := 0; i < len(list); i = i + addressLimit {
		rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
			summaryAddress, "", "", "",
			i, addressLimit, &contract, &feesSupport)
		if err != nil {
			log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
			return
		}

		//执行汇总交易
		for _, rawTxWithErr := range rawTxArray {

			if rawTxWithErr.Error != nil {
				log.Error(rawTxWithErr.Error.Error())
				continue
			}

			_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
			if err != nil {
				return
			}

			_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
			if err != nil {
				return
			}

			_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
			if err != nil {
				return
			}
		}
	}

}
