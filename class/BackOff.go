package class

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//PreOrdMarginCalculation : This api will calculate and get the exposure margin and span margin of all the margin of the client.
func PreOrdMarginCalculation(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "PreOrdMarginCalculation"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodePreOrdMarCal + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"OrderRequestorCode":"` + ClientCode + `",
		"Exch":"N",
		"ExchType":"D",
		"ClientCode":"` + ClientCode + `",
		"ScripCode":"45609",
		"PlaceModifyCancel":"M",
		"TransactionType":"B",
		"AtMarket":"false",
		"LimitRate":650,
		"WithSL":"N",
		"SLTriggerRate":0,
		"IsSLTriggered":"N",
		"Volume":250,
		"OldTradedQty":0,
		"ProductType":"D",
		"ExchOrderId":"0",
		"ClientIP":"21.1.2",
		"AppSource":54
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
	// 1i4yexcaxvqpfnvuq1mlajjc - get this from the login cokiee
	req.Header.Add("Cookie", "IIFLMarcookie=1i4yexcaxvqpfnvuq1mlajjc")

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var PreOrdMarginCalculationRes PreOrdMarginCalculationResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &PreOrdMarginCalculationRes)
	fmt.Printf("Message: %s, StatusDescription: %s", PreOrdMarginCalculationRes.Body.Message, PreOrdMarginCalculationRes.Head.StatusDescription)

}

//BackoffMutualFundTransaction : This API is used to the give the information BackOffice client mutual fund transaction information
func BackoffMutualFundTransaction(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "BackoffMutualFundTransaction"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodeBackoffMutulFundTransaction + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `",
		"FromDate":"20190101",
		"ToDate":"20201001"
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
	// 1i4yexcaxvqpfnvuq1mlajjc - get this from the login cokiee
	req.Header.Add("Cookie", "IIFLMarcookie=1i4yexcaxvqpfnvuq1mlajjc")

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var BackoffMutualFundTransactionRes BackoffMutualFundTransactionResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &BackoffMutualFundTransactionRes)
	fmt.Printf("Message: %s, StatusDescription: %s", BackoffMutualFundTransactionRes.Body.Message, BackoffMutualFundTransactionRes.Head.StatusDescription)

}

//BackoffClientProfile : This API is used to the give the information BackOffice client profile data
func BackoffClientProfile(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "BackoffClientProfile"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodeBackoffClientProfile + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `",
		"FromDate":"20190101",
		"ToDate":"20201001"
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
	// 1i4yexcaxvqpfnvuq1mlajjc - get this from the login cokiee
	req.Header.Add("Cookie", "IIFLMarcookie=1i4yexcaxvqpfnvuq1mlajjc")

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var BackoffClientProfileRes BackoffClientProfileResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &BackoffClientProfileRes)
	fmt.Printf("Message: %s, StatusDescription: %s", BackoffClientProfileRes.Body.Message, BackoffClientProfileRes.Head.StatusDescription)

}

//BackoffEquitytransaction : This API is used to the give the information BackOffice client equity transaction information
func BackoffEquitytransaction(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "BackoffEquitytransaction"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodeBackoffEquitytransaction + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `",
		"FromDate":"20190101",
		"ToDate":"20201001"
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
	// 1i4yexcaxvqpfnvuq1mlajjc - get this from the login cokiee
	req.Header.Add("Cookie", "IIFLMarcookie=1i4yexcaxvqpfnvuq1mlajjc")

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var BackoffClientProfileRes BackoffClientProfileResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &BackoffClientProfileRes)
	fmt.Printf("Message: %s, StatusDescription: %s", BackoffClientProfileRes.Body.Message, BackoffClientProfileRes.Head.StatusDescription)

}

//BackoffFutureTransaction : This API is used to the give the information BackOffice client future transaction information
func BackoffFutureTransaction(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "BackoffFutureTransaction"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodeBackoffFutureTransaction + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `",
		"FromDate":"20190101",
		"ToDate":"20201001"
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
	// 1i4yexcaxvqpfnvuq1mlajjc - get this from the login cokiee
	req.Header.Add("Cookie", "IIFLMarcookie=1i4yexcaxvqpfnvuq1mlajjc")

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var BackoffFutureTransactionRes BackoffFutureTransactionResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &BackoffFutureTransactionRes)
	fmt.Printf("Message: %s, StatusDescription: %s", BackoffFutureTransactionRes.Body.Message, BackoffFutureTransactionRes.Head.StatusDescription)

}

//BackoffoptionTransaction : This API is used to the give the information BackOffice client option transaction information
func BackoffoptionTransaction(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "BackoffoptionTransaction"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodeBackoffoptionTransaction + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `",
		"FromDate":"20190101",
		"ToDate":"20201001"
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
	// 1i4yexcaxvqpfnvuq1mlajjc - get this from the login cokiee
	req.Header.Add("Cookie", "IIFLMarcookie=1i4yexcaxvqpfnvuq1mlajjc")

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var BackoffoptionTransactionRes BackoffoptionTransactionResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &BackoffoptionTransactionRes)
	fmt.Printf("Message: %s, StatusDescription: %s", BackoffoptionTransactionRes.Body.Message, BackoffoptionTransactionRes.Head.StatusDescription)

}

//BackoffLedger : This API is used to the give the information BackOffice ledger information
func BackoffLedger(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "BackoffLedger"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodeBackoffLedger + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `",
		"FromDate":"20190101",
		"ToDate":"20201001"
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
	// 1i4yexcaxvqpfnvuq1mlajjc - get this from the login cokiee
	req.Header.Add("Cookie", "IIFLMarcookie=1i4yexcaxvqpfnvuq1mlajjc")

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var BackoffLedgerRes BackoffLedgerResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &BackoffLedgerRes)
	fmt.Printf("Message: %s, StatusDescription: %s", BackoffLedgerRes.Body.Message, BackoffLedgerRes.Head.StatusDescription)

}

//BackoffDPTransaction : This API is used to the give the information BackOffice DP transaction
func BackoffDPTransaction(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "BackoffDPTransaction"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodeBackoffDPTransaction + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `",
		"FromDate":"20190101",
		"ToDate":"20201001"
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
	// 1i4yexcaxvqpfnvuq1mlajjc - get this from the login cokiee
	req.Header.Add("Cookie", "IIFLMarcookie=1i4yexcaxvqpfnvuq1mlajjc")

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var BackoffLedgerRes BackoffLedgerResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &BackoffLedgerRes)
	fmt.Printf("Message: %s, StatusDescription: %s", BackoffLedgerRes.Body.Message, BackoffLedgerRes.Head.StatusDescription)

}

//BackoffDPHolding : This API is used to the give the information BackOffice DP Holding information
func BackoffDPHolding(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "BackoffDPHolding"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodeBackoffDPHolding + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `",
		"FromDate":"20190101",
		"ToDate":"20201001"
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
	// 1i4yexcaxvqpfnvuq1mlajjc - get this from the login cokiee
	req.Header.Add("Cookie", "IIFLMarcookie=1i4yexcaxvqpfnvuq1mlajjc")

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var BackoffDPHoldingRes BackoffDPHoldingResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &BackoffDPHoldingRes)
	fmt.Printf("Message: %s, StatusDescription: %s", BackoffDPHoldingRes.Body.Message, BackoffDPHoldingRes.Head.StatusDescription)

}
