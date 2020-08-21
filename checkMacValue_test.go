package ecpayaio

import (
	"testing"
)

func TestGenerateCheckMacValue(t *testing.T) {

	payload := map[string]string{
		"TradeDesc":         "促銷方案",
		"PaymentType":       "aio",
		"MerchantTradeDate": "2013/03/12 15:30:23",
		"MerchantTradeNo":   "ecpay20130312153023",
		"MerchantID":        "2000132",
		"ReturnURL":         "https://www.ecpay.com.tw/receive.php",
		"ItemName":          "Apple iphone 7 手機殼",
		"TotalAmount":       "1000",
		"ChoosePayment":     "ALL",
		"EncryptType":       "1",
	}

	hashKey := "5294y06JbISpM5x9"
	hashIV := "v77hoKGq4kWxNNIS"

	macValue := GenerateCheckMacValue(payload, hashKey, hashIV, EncryptTypeSHA256)

	expected := "CFA9BDE377361FBDD8F160274930E815D1A8A2E3E80CE7D404C45FC9A0A1E407"

	if macValue != expected {
		t.Errorf("expected: " + expected + " got: " + macValue)
	}

}
