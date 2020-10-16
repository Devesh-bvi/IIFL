package class

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//Holding : This method is used to provide the client’s holdings as of beginning of the day.
func Holding(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "Holding"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodeHoldingV2 + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `"
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

	var HoldRes HoldResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &HoldRes)
	fmt.Printf("Message: %s, StatusDescription: %s", HoldRes.Body.Message, HoldRes.Head.StatusDescription)

}

//Margin : This API is used to fetch Client’s Available Margin details.
func Margin(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "Margin"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodeMarginV3 + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `"
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

	var MarginRes MarginResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &MarginRes)
	fmt.Printf("TimeStamp: %s, StatusDescription: %s", MarginRes.Body.TimeStamp, MarginRes.Head.StatusDescription)

}

//OrderBookV2 : This API is used to Fetch Order Book of a particular Client which will contain both Equity and Commodity Orders.
func OrderBookV2(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "OrderBookV2"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodeOrdBkV2 + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `"
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

	var OrderBookV2Res OrderBookV2Response

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &OrderBookV2Res)
	fmt.Printf("Reason: %s, StatusDescription: %s", OrderBookV2Res.Body.OrderBookDetail[0].Reason, OrderBookV2Res.Head.StatusDescription)

}

//TradeBook : This method is used to provide the daily trade details of the day.
func TradeBook(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "TradeBook"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodeTrdBkV1 + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `"
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

	var TradeBookRes TradeBookRespone

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &TradeBookRes)
	fmt.Printf("Message: %s, StatusDescription: %s", TradeBookRes.Body.Message, TradeBookRes.Head.StatusDescription)

}

//NetPosition : This method will provide the client’s Positions in both equity and commodity market across all exchanges and exchange segments.
func NetPosition(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "NetPosition"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodeNetPositionV4 + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `"
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

	var NetPositionRes NetPositionResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &NetPositionRes)
	fmt.Printf("Message: %s, StatusDescription: %s", NetPositionRes.Body.Message, NetPositionRes.Head.StatusDescription)

}

//NetPositionNetWise : This API will provide client Net Position report with additional parameter as multiplier.
func NetPositionNetWise(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "NetPositionNetWise"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
			"appName": "` + config.AppName + `",
			"appVer": "` + config.AppVer + `",
			"key": "` + config.Key + `",
			"osName": "` + config.OsName + `",
			"requestCode": "` + config.RequestCodeNPNWV1 + `",
			"userId":"` + config.UserID + `",
			"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `"
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
	// jhj5pipoc4xzda1f5hhj5kip - get this from the login cokiee
	req.Header.Add("Cookie", "IIFLMarcookie=jhj5pipoc4xzda1f5hhj5kip")

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var NetPositionNetWiseRes NetPositionNetWiseResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &NetPositionNetWiseRes)
	fmt.Printf("Message: %s, StatusDescription: %s", NetPositionNetWiseRes.Body.Message, NetPositionNetWiseRes.Head.StatusDescription)

}
