package unionpay

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// BackNotifyReceive  post application/x-www-form-urlencoded

func (c *UnionPay) BackNotifyReceive(writer http.ResponseWriter, request *http.Request) (map[string]string, error) {
	var (
		pass      bool
		err       error
		queryInfo map[string]string
	)
	err = request.ParseForm()
	if err != nil {
		fmt.Println("银联支付回调request.ParseForm,err:\n", err)
	}
	signParam := make(map[string]string)
	for k, v := range request.PostForm {
		signParam[k] = v[0]
	}
	resultMap := make(map[string]string)
	resultMap["response_text"] = "error"
	pass, err = SignVerify(signParam)
	if pass {
		fmt.Println("银联支付回调pass:\n", pass)
		//判断支付订单

		orderId := signParam["orderId"]
		queryInfo, err = c.Query(orderId)
		bytes, _ := json.Marshal(queryInfo)
		fmt.Println("银联支付订单查询info:\n", string(bytes))
		if err != nil {
			fmt.Println("银联支付订单查询err:", err)
			return resultMap, err
		}
		//判断交易状态
		if status, ok := queryInfo["origRespCode"]; ok && status == "00" {
			// 金额
			fmt.Println("银联支付回调金额：", signParam["txnAmt"])
			amount, _ := strconv.ParseFloat(signParam["txnAmt"], 64)
			fmt.Println("银联支付回调金额（分）：", amount)
			_, err = writer.Write([]byte("ok"))
			if err != nil {
				return resultMap, nil
			}
			// 对应订单号
			resultMap["order_id"] = signParam["orderId"]
			// 订单号、流水号
			resultMap["transaction_id"] = signParam["queryId"]
			resultMap["amount"] = signParam["txnAmt"]
			resultMap["response_text"] = "ok"

			return resultMap, nil
		} else {
			fmt.Println("status not success:", status)
		}
	} else {
		fmt.Printf("银联支付回调err:%v", err)
	}
	return resultMap, err

}
