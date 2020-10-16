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
