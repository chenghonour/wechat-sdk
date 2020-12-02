package mch

import (
	"encoding/base64"
	"strconv"

	"github.com/shenghui0779/gochat/wx"
)

// TransferBalanceData 付款到零钱数据
type TransferBalanceData struct {
	// 必填参数
	PartnerTradeNO string // 商户订单号，需保持唯一性 (只能是字母或者数字，不能包含有其它字符)
	OpenID         string // 商户appid下，某用户的openid
	CheckName      string // NO_CHECK：不校验真实姓名；FORCE_CHECK：强校验真实姓名
	Amount         int    // 企业付款金额，单位：分
	Desc           string // 企业付款备注，必填。注意：备注中的敏感词会被转成字符*
	// 选填参数
	ReUserName     string // 收款用户真实姓名。如果check_name设置为FORCE_CHECK，则必填用户真实姓名
	DeviceInfo     string // 微信支付分配的终端设备号
	SpbillCreateIP string // 该IP同在商户平台设置的IP白名单中的IP没有关联，该IP可传用户端或者服务端的IP
}

// TransferBankCardData 付款到银行卡数据
type TransferBankCardData struct {
	// 必填参数
	PartnerTradeNO string // 商户订单号，需保持唯一（只允许数字[0~9]或字母[action~Z]和[a~z]，最短8位，最长32位）
	EncBankNO      string // 收款方银行卡号（采用标准RSA算法，公钥由微信侧提供）
	EncTrueName    string // 收款方用户名（采用标准RSA算法，公钥由微信侧提供）
	BankCode       string // 银行卡所在开户行编号，参考：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_4
	Amount         int    // 付款金额：RMB分（支付总额，不含手续费）注：大于0的整数
	// 选填参数
	Desc string // 企业付款到银行卡付款说明，即订单备注（UTF8编码，允许100个字符以内）
}

// TransferToBalance 付款到零钱【注意：当返回错误码为“SYSTEMERROR”时，请务必使用原商户订单号重试，否则可能造成重复支付等资金风险。】
func TransferToBalance(data *TransferBalanceData) wx.Action {
	return wx.NewMchAPI(TransferToBalanceURL, func(appid, mchid, apikey, nonce string) (wx.WXML, error) {
		body := wx.WXML{
			"mch_appid":        appid,
			"mchid":            mchid,
			"nonce_str":        nonce,
			"partner_trade_no": data.PartnerTradeNO,
			"openid":           data.OpenID,
			"check_name":       data.CheckName,
			"amount":           strconv.Itoa(data.Amount),
			"desc":             data.Desc,
		}

		if data.ReUserName != "" {
			body["re_user_name"] = data.ReUserName
		}

		if data.DeviceInfo != "" {
			body["device_info"] = data.DeviceInfo
		}

		if data.SpbillCreateIP != "" {
			body["spbill_create_ip"] = data.SpbillCreateIP
		}

		body["sign"] = wx.SignWithMD5(body, apikey, true)

		return body, nil
	}, true)
}

// QueryTransferBalanceOrder 查询付款到零钱订单
func QueryTransferBalanceOrder(partnerTradeNO string) wx.Action {
	return wx.NewMchAPI(TransferBalanceOrderQueryURL, func(appid, mchid, apikey, nonce string) (wx.WXML, error) {
		body := wx.WXML{
			"appid":            appid,
			"mch_id":           mchid,
			"partner_trade_no": partnerTradeNO,
			"nonce_str":        nonce,
		}

		body["sign"] = wx.SignWithMD5(body, apikey, true)

		return body, nil
	}, true)
}

// TransferToBankCard 付款到银行卡【注意：当返回错误码为“SYSTEMERROR”时，请务必使用原商户订单号重试，否则可能造成重复支付等资金风险。】
func TransferToBankCard(data *TransferBankCardData, publicKey []byte) wx.Action {
	return wx.NewMchAPI(TransferToBankCardURL, func(appid, mchid, apikey, nonce string) (wx.WXML, error) {
		body := wx.WXML{
			"mch_id":           mchid,
			"nonce_str":        nonce,
			"partner_trade_no": data.PartnerTradeNO,
			"bank_code":        data.BankCode,
			"amount":           strconv.Itoa(data.Amount),
		}

		// 收款方银行卡号加密
		b, err := wx.RSAEncrypt([]byte(data.EncBankNO), publicKey)

		if err != nil {
			return nil, err
		}

		body["enc_bank_no"] = base64.StdEncoding.EncodeToString(b)

		// 收款方用户名加密
		b, err = wx.RSAEncrypt([]byte(data.EncTrueName), publicKey)

		if err != nil {
			return nil, err
		}

		body["enc_true_name"] = base64.StdEncoding.EncodeToString(b)

		if data.Desc != "" {
			body["desc"] = data.Desc
		}

		body["sign"] = wx.SignWithMD5(body, apikey, true)

		return body, nil
	}, true)
}

// QueryTransferBankCardOrder 查询付款到银行卡订单
func QueryTransferBankCardOrder(partnerTradeNO string) wx.Action {
	return wx.NewMchAPI(TransferBankCardOrderQueryURL, func(appid, mchid, apikey, nonce string) (wx.WXML, error) {
		body := wx.WXML{
			"mch_id":           mchid,
			"partner_trade_no": partnerTradeNO,
			"nonce_str":        nonce,
		}

		body["sign"] = wx.SignWithMD5(body, apikey, true)

		return body, nil
	}, true)
}

// RSAPublicKey 获取RSA加密公钥
func RSAPublicKey() wx.Action {
	return wx.NewMchAPI(RSAPublicKeyURL, func(appid, mchid, apikey, nonce string) (wx.WXML, error) {
		body := wx.WXML{
			"mch_id":    mchid,
			"nonce_str": nonce,
			"sign_type": SignMD5,
		}

		body["sign"] = wx.SignWithMD5(body, apikey, true)

		return body, nil
	}, true)
}