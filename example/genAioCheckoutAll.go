package main

import (
	"fmt"
	"github.com/tihtw/ECPayAIO_Golang"
	"time"
)

func main() {
	gen()
}

func gen() {

	hashKey := "5294y06JbISpM5x9"
	hashIV := "v77hoKGq4kWxNNIS"

	c := ecpayaio.NewCheckout(ecpayaio.HostStage)
	c.SetHashKey(hashKey)
	c.SetHashIV(hashIV)
	c.SetMerchantID("2000132")
	c.SetTradeDesc("促銷方案")
	c.SetTotalAmount(1000)
	c.SetItemName([]string{"Apple iphone 7 手機殼", "網紅小遙"})
	c.SetMerchantTradeNo(time.Now().Format("20060102150405"))
	c.SetReturnURL("https://fec50e836a10.ngrok.io/QAQ")

	c.SetChoosePayment(ecpayaio.ChoosePaymentCredit)

	form, err := c.GeneratePostForm()
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println(form)

}
