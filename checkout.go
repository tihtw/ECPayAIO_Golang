package ecpayaio

import (
	"bytes"
	"fmt"
	"html/template"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var tmpl = `<!DOCTYPE html>
<html>
	<head>
	</head>
	<body>
		<form id="apiFrom" action="{{.url}}" method="POST">
		{{- range $key, $value := .payload}}
			<input type="hidden" name="{{$key}}" value="{{$value}}" />
		{{- end}}
		</form>
		<script type="text/javascript">
			apiFrom.submit()
		</script>
	</body>
</html>`

type ChoosePayment string

var (
	ChoosePaymentAll     ChoosePayment = "ALL"
	ChoosePaymentCredit  ChoosePayment = "Credit"
	ChoosePaymentWebATM  ChoosePayment = "WebATM"
	ChoosePaymentATM     ChoosePayment = "ATM"
	ChoosePaymentCVS     ChoosePayment = "CVS"
	ChoosePaymentBarcode ChoosePayment = "BARCODE"
)

type Language string

var (
	LanguageZhTw Language = ""
	LanguageEn   Language = "ENG"
	LanguageKr   Language = "KOR"
	LanguageJa   Language = "JPN"
	LanguageZhCn Language = "CHI"
)

// n. 收銀台
type Checkout struct {
	payload map[string]string
	hashKey string
	hashIV  string
	host    string
}

var (
	HostStage      = "https://payment-stage.ecpay.com.tw/Cashier/AioCheckOut/V5"
	HostProduction = "https://payment-stage.ecpay.com.tw/Cashier/AioCheckOut/V5"
)

func NewCheckout(host string) *Checkout {

	ret := Checkout{
		payload: map[string]string{
			"PaymentType":       "aio",
			"MerchantTradeDate": time.Now().Format("2006/01/02 15:04:05"),
			"DeviceSource":      "",
			"ChoosePayment":     "ALL",
			"EncryptType":       "1",
		},
		host: host,
	}
	return &ret

}

func (c *Checkout) SetMerchantID(value string) {
	c.payload["MerchantID"] = value
}

func (c *Checkout) SetMerchantTradeNo(value string) {
	c.payload["MerchantTradeNo"] = value
}

func (c *Checkout) SetStoreID(value string) {
	c.payload["StoreID"] = value
}
func (c *Checkout) SetMerchantTradeDate(value string) {
	c.payload["MerchantTradeDate"] = value
}
func (c *Checkout) SetMerchantTradeDateByTime(t time.Time) {
	c.payload["MerchantTradeDate"] = t.Format("2006/01/12 15:04:05")
}

// Fixed aio
// func (c *Checkout) SetPaymenrType(t *PaymentType) {
// 	c.payload["PaymentType"] = t.string
// }

// 結帳總額
func (c *Checkout) SetTotalAmount(value int) {
	c.payload["TotalAmount"] = strconv.Itoa(value)
}

func (c *Checkout) SetTradeDesc(value string) {
	c.payload["TradeDesc"] = urlencoded(value)
}
func (c *Checkout) SetItemName(values []string) {
	c.payload["ItemName"] = strings.Join(values, "#")
}
func (c *Checkout) SetReturnURL(value string) {
	c.payload["ReturnURL"] = value
}
func (c *Checkout) SeChoosePayment(value ChoosePayment) {
	c.payload["ChoosePayment"] = string(value)
}
func (c *Checkout) SetClientBackURL(value string) {
	c.payload["ClientBackURL"] = value
	delete(c.payload, "OrderResultURL")
}
func (c *Checkout) SetItemURL(value string) {
	c.payload["ItemURL"] = value
}
func (c *Checkout) SetRemark(value string) {
	c.payload["Remark"] = value
}
func (c *Checkout) SetChooseSubPayment(value string) {
	c.payload["ChooseSubPayment"] = value
}

func (c *Checkout) SetOrderResultURL(value string) {
	c.payload["OrderResultURL"] = value
}

func (c *Checkout) SetNeedExtraPaidInfo(value bool) {
	sv := "Y"
	if !value {
		sv = "N"
	}
	c.payload["NeedExtraPaidInfo"] = sv
}

func (c *Checkout) SetIgnorePayment(values []ChoosePayment) {
	ss := []string{}
	for _, it := range values {
		ss = append(ss, string(it))
	}
	c.payload["IgnorePayment"] = strings.Join(ss, "#")
}
func (c *Checkout) SetPlatformID(value string) {
	c.payload["PlatformID"] = value
}
func (c *Checkout) SetInvoiceMark(value bool) {
	sv := "Y"
	if !value {
		sv = "N"
	}
	c.payload["InvoiceMark"] = sv
}

func (c *Checkout) SetCustonField1(value string) {
	c.payload["CustonField1"] = value
}
func (c *Checkout) SetCustonField2(value string) {
	c.payload["CustonField2"] = value
}
func (c *Checkout) SetCustonField3(value string) {
	c.payload["CustonField3"] = value
}
func (c *Checkout) SetCustonField4(value string) {
	c.payload["CustonField4"] = value
}
func (c *Checkout) SetLanguage(value Language) {
	c.payload["Language"] = string(value)
}

// If Choosen Payment including ATM
func (c *Checkout) SetExpireDate(value int) {
	c.payload["ExpireDate"] = strconv.Itoa(value)
}
func (c *Checkout) SetPaymentInfoURL(value string) {
	c.payload["PaymentInfoURL"] = value
}
func (c *Checkout) SetClientRedirectURL(value string) {
	c.payload["ClientRedirectURL"] = value
}

// If Choosen payment including credit
func (c *Checkout) SetBindingCard(bindingCard bool, merchantMemberID string) {
	sv := "1"
	if !bindingCard {
		sv = "0"
	}
	c.payload["BindingCard"] = sv
	c.payload["MerchantMemberID"] = merchantMemberID

}

func (c *Checkout) SetHashKey(value string) {
	c.hashKey = value
}

func (c *Checkout) SetHashIV(value string) {
	c.hashIV = value
}

func (c *Checkout) SetParameter(parameter string, value string) {
	c.payload[parameter] = value
}

func (c *Checkout) String() string {
	return fmt.Sprintln("%q", c.payload)

}

func (c *Checkout) GeneratePostForm() (string, error) {

	// hashKey := "5294y06JbISpM5x9"
	// hashIV := "v77hoKGq4kWxNNIS"

	macValue := GenerateCheckMacValue(c.payload, c.hashKey, c.hashIV, EncryptTypeSHA256)

	c.payload["CheckMacValue"] = macValue

	t := template.New("t")

	t, err := t.Parse(tmpl)
	if err != nil {
		// fmt.Println(err)
		return "", err
	}

	ret := ""
	w := bytes.NewBufferString(ret)
	err = t.Execute(w, map[string]interface{}{
		"payload": c.payload,
		"url":     c.host,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return w.String(), nil
}

func urlencoded(s string) string {
	return url.QueryEscape(s)
}
