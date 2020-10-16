package class

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//MarketFeed :- This API is used to fetch the market feed of a particular scrip or a set of scrips
func MarketFeed(w http.ResponseWriter, r *http.Request) {

	ClientCode := r.URL.Query().Get("ClientCode")

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "MarketFeed"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
		"appName":"` + config.AppName + `",
		"appVer":"` + config.AppVer + `",
		"key":"` + config.Key + `",
		"osName":"` + config.OsName + `",
		"userId":"` + config.UserID + `",
		"password":"` + config.Password + `",
		"requestCode":"` + config.RequestCodeMarketFeed + `"
		},
		"body":{
		"ClientCode":"` + ClientCode + `",
		"Count":"2",
		"MarketFeedData":[
		{
		"Exch":"N",
		"ExchType":"C",
		"ScripCode":"2885"
		},
		{
		"Exch":"N",
		"ExchType":"C",
		"ScripCode":"22"
		}
		],
		"ClientLoginType":"0",
		"LastRequestTime":"/Date(1600248018615)/",
		"RefreshRate":"H"
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

	var MarketFeedRes MarketFeedResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &MarketFeedRes)
	fmt.Printf("TimeStamp: %s, StatusDescription: %s", MarketFeedRes.Body.TimeStamp, MarketFeedRes.Head.StatusDescription)

}
