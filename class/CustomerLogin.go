package class

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//LoginRequestMobileForVendor : his method allows the developer to login the client into the application
func LoginRequestMobileForVendor(w http.ResponseWriter, r *http.Request) {

	//	fmt.Printf("GetIP: %s", class.GetIP(r))
	//	fmt.Printf("GetIP2: %s", class.GetIPAddress())

	EmailID := r.URL.Query().Get("EmailID")
	ContactNumber := r.URL.Query().Get("ContactNumber")

	var IPadd string
	IPadd = GetIPAddress()

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "LoginRequestMobileForVendor"

	method := "POST"

	//"UserID": "Ae86Ls5iwTrPOm1w6PL2Cg==",
	//"Password": "MnB1tFNCzzLAf8ezdonXVw=="

	payload := strings.NewReader(`{
		"head": {
		  "appName": "` + config.AppName + `",
		  "appVer": "` + config.AppVer + `",
		  "key": "` + config.Key + `",
		  "osName": "` + config.OsName + `",
		  "requestCode": "` + config.RequestCode + `",
		  "userId": "Ae86Ls5iwTrPOm1w6PL2Cg==",
		  "password": "MnB1tFNCzzLAf8ezdonXVw=="
		},
		"body": {
		  "Email_id": "` + EmailID + `",
		  "LocalIP": "` + IPadd + `",
		  "PublicIP": "` + IPadd + `",
		  "ContactNumber": "` + ContactNumber + `"
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

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var LoginVendorRes LoginVendorResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &LoginVendorRes)
	fmt.Printf("Message: %s, TTManagerID: %s", LoginVendorRes.Body.Message, LoginVendorRes.Body.TTManagerID)

}

//LoginRequest : This API is used to fetch Client’s Available Margin details.
func LoginRequest(w http.ResponseWriter, r *http.Request) {

	var IPadd string
	IPadd = GetIPAddress()

	config := CheckConfig()
	//fmt.Println(domainame.AppName)
	url := config.ServiceURL + "LoginRequest"

	method := "POST"

	payload := strings.NewReader(`{
		"head":{
		"requestCode":"` + config.RequestCode + `",
		"key":"` + config.Key + `",
		"appVer":"` + config.AppVer + `",
		"appName": "` + config.AppName + `",
		"osName":"` + config.OsName + `",
		"userId":"` + config.UserID + `",
		"password":"` + config.Password + `"
		},
		"body":{
		"ClientCode":"y24giIyBtUNgutz1wYTM5w==",
		"Password":"49dKKNv3UFpwoK+st09O5A==",
		"HDSerialNumber":"asdf",
		"MACAddress":"asdf",
		"MachineID":"asfdasdf",
		"VersionNo":"` + config.VersionNo + `",
		"RequestNo":1,
		"My2PIN":"X8EUGQ2U6OzA+M6ejtB/uQ==",
		"ConnectionType":1,
		"LocalIP":"` + IPadd + `",
		"PublicIP":"` + IPadd + `"
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

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var session string
	var path string
	for _, cookie := range res.Cookies() {
		if cookie.Name == "IIFLMarcookie" {
			session = cookie.Value
			path = cookie.Path
		}
	}
	fmt.Printf("session: %s, path: %s", session, path)
	//fmt.Println(string(body))
	bodyString := string(body)
	json.NewEncoder(w).Encode(bodyString)

	var LoginClientRes LoginClientResponse

	// json to golang struct
	json.Unmarshal([]byte(bodyString), &LoginClientRes)
	fmt.Printf("\nClientName: %s, StatusDescription: %s", LoginClientRes.Body.ClientName, LoginClientRes.Head.StatusDescription)

}
