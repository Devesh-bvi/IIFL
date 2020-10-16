package class

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//OrderRequest :- This API is the backbone of all the trade dependent APIs and is the key method that you would need to place your order in the market.
func OrderRequest(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "OrderRequest"

	method := "POST"

	payload := strings.NewReader(`{
		"_ReqData":{
		"head":{
		"appName":"` + config.AppName + `",
		"appVer":"` + config.AppVer + `",
		"key":"` + config.Key + `",
		"osName":"` + config.OsName + `",
		"userId":"` + config.UserID + `",
		"password":"` + config.Password + `",
		"requestCode":"` + config.RequestCodeOrdReq + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `",
		"OrderFor":"P",
		"Exchange":"N",
		"ExchangeType":"C",
		"Price":190,
		"OrderID":1,
		"OrderType":"BUY",
		"Qty":1,
		"OrderDateTime":"/Date(1601880914379)/",
		"ScripCode":3045,
		"AtMarket":false,
		"RemoteOrderID":"s000220190",
		"ExchOrderID":"0",
		"DisQty":0,
		"IsStopLossOrder":false,
		"StopLossPrice":0,
		"IsVTD":false,
		"IOCOrder":false,
		"IsIntraday":false,
		"PublicIP":"192.168.84.215",
		"AHPlaced":"N",
		"ValidTillDate":"/Date(1602880914379)/",
		"iOrderValidity":0,
		"OrderRequesterCode":"96131461",
		"TradedQty":0
		}
		},
		"AppSource":54
		}`)
	//fmt.Println(payload)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Ocp-Apim-Subscription-Key", config.OcpKey)
	req.Header.Add("Content-Type", "application/json")
	// ibg0klmapcrilrwzra3ctne0 - get this from the login cokiee
	req.Header.Add("Cookie", "IIFLMarcookie=ibg0klmapcrilrwzra3ctne0")

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var OrderRes OrderResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &OrderRes)
	fmt.Printf("Message: %s, StatusDescription: %s", OrderRes.Body.Message, OrderRes.Head.StatusDescription)

}

//OrderStatus :- The purpose of this method is to navigate the status of an order placed by the client.
func OrderStatus(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "OrderStatus"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
		"appName":"` + config.AppName + `",
		"appVer":"` + config.AppVer + `",
		"key":"` + config.Key + `",
		"osName":"` + config.OsName + `",
		"userId":"` + config.UserID + `",
		"password":"` + config.Password + `",
		"requestCode":"` + config.RequestCodeOrdStatus + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `",
		"OrdStatusReqList":[
		{
		"Exch":"N",
		"ExchType":"C",
		"ScripCode":4717,
		"RemoteOrderID":"S123456789123456789"
		}
		]
		}
		}`)
	//fmt.Println(payload)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Ocp-Apim-Subscription-Key", config.OcpKey)
	req.Header.Add("Content-Type", "application/json")
	// ibg0klmapcrilrwzra3ctne0 - get this from the login cokiee
	req.Header.Add("Cookie", "IIFLMarcookie=ibg0klmapcrilrwzra3ctne0")

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var OrderStatusRes OrderStatusResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &OrderStatusRes)
	fmt.Printf("Message: %s, StatusDescription: %s", OrderStatusRes.Body.Message, OrderStatusRes.Head.StatusDescription)

}

//TradeInformation :- Trade Information method is used to Fetch Trade Information for a set of orders placed by the client.
func TradeInformation(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "TradeInformation"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
		"appName":"` + config.AppName + `",
		"appVer":"` + config.AppVer + `",
		"key":"` + config.Key + `",
		"osName":"` + config.OsName + `",
		"userId":"` + config.UserID + `",
		"password":"` + config.Password + `",
		"requestCode":"` + config.RequestCodeTrdInfo + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `",
		"TradeInformationList":[
		{
		"Exch":"B",
		"ExchType":"C",
		"ScripCode":500410,
		"ExchOrderID":"1557728588259000015"
		}
		]
		}
		}`)
	//fmt.Println(payload)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Ocp-Apim-Subscription-Key", config.OcpKey)
	req.Header.Add("Content-Type", "application/json")
	// ibg0klmapcrilrwzra3ctne0 - get this from the login cokiee
	req.Header.Add("Cookie", "IIFLMarcookie=ibg0klmapcrilrwzra3ctne0")

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var TradeInformationRes TradeInformationResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &TradeInformationRes)
	fmt.Printf("Message: %s, StatusDescription: %s", TradeInformationRes.Body.Message, TradeInformationRes.Head.StatusDescription)

}
