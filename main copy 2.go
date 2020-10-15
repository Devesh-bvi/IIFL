package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type body struct {
	Email_id      string
	LocalIP       string
	PublicIP      string
	ContactNumber string
}

func loginRequestMobileForVendor(w http.ResponseWriter, r *http.Request) {

	e := strings.NewReader(`{
		"Email_id":      "shubham.rajak@iifl.com",
		"LocalIP":       "123.123.123.123",
		"PublicIP":      "123.123.123.12",
		"ContactNumber": "7769941110",
	}`)

	//jsonReq, err := json.Marshal(e)

	/*if err != nil {
		log.Fatalln(err)
	}*/

	req, err := http.NewRequest("POST", "https://dataservice.iifl.in/openapi/prod/LoginRequestMobileForVendor", e)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("appName", "IIFLMarSHUBHAMR")
	req.Header.Set("appVer", "1.0")
	req.Header.Set("key", "YxjDEpOqxMSHJoXWb06CxjLNJxsZGNR0Z7Rg112orvaraf5UsWApy5tH8gCIEDO3")
	req.Header.Set("osName", "Android")
	req.Header.Set("requestCode", "IIFLMarRQLoginForVendor")
	req.Header.Set("userId", "Ae86Ls5iwTrPOm1w6PL2Cg==")
	req.Header.Set("password", "MnB1tFNCzzLAf8ezdonXVw==")
	req.Header.Set("Ocp-Apim-Subscription-Key", "fc714d8e5b82438a93a95baa493ff45b")

	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	json.NewEncoder(w).Encode(bodyString)
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	// replace http.HandleFunc with myRouter.HandleFunc

	myRouter.HandleFunc("/loginForVendor", loginRequestMobileForVendor)

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	//fmt.Println("Programmer = ", json.NewEncoder(w).Encode(e))
	handleRequests()
}
