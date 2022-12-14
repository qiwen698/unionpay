package unionpay

import (
	"log"
	"testing"
)

// demo
var pay = NewUnionPay(UnionPay{
	Mode:    "dev",
	MerID:   "777290058110048",
	PfxPath: "./cert/acp_test_sign.pfx",
	PfxPwd:  "000000",
	BackURL: "http://localhost.com",
})

func TestAppConsume(t *testing.T) {
	if result, err := pay.AppConsume(1000, "2022120117335150401999", ""); err == nil {
		log.Printf("merId: %v ,orderId: %v ,Tn: %v", result["merId"], result["orderId"], result["tn"])
	} else {
		log.Println(err)
	}
}

func TestQuery(t *testing.T) {
	if result, err := pay.Query("2212136035003593"); err == nil {
		log.Println(result)
	} else {
		log.Println(err)
	}
}
