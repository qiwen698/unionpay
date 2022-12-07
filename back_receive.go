package unionpay

import (
	"net/http"
)

// BackNotifyReceive  post application/x-www-form-urlencoded

func (c *UnionPay) BackNotifyReceive(writer http.ResponseWriter, request *http.Request) (map[string]interface{}, error) {
	request.ParseForm()
	m := make(map[string]interface{})
	for k, v := range request.PostForm {
		m[k] = v[0]
	}
	return m, nil

}
