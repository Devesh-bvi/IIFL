package class

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type raw struct {
	language string
}

type options struct {
	raw
}

type body struct {
	mode string
	raw  string
	options
}

func loginRequestMobileForVendor() {

	e := body{
		mode: "raw",
		raw:  "{\r\n  \"head\": {\r\n    \"appName\": \"IIFLMarSHUBHAMR\",\r\n    \"appVer\": \"1.0\",\r\n    \"key\": \"ThG6LZT8JJUWi5mD8OMHYTdsOmFzkebF\",\r\n    \"osName\": \"Android\",\r\n    \"requestCode\": \"IIFLMarRQLoginForVendor\",\r\n    \"userId\": \"Ae86Ls5iwTrPOm1w6PL2Cg==\",\r\n    \"password\": \"MnB1tFNCzzLAf8ezdonXVw==\"\r\n  },\r\n  \"body\": {\r\n    \"Email_id\": \"shubham.rajak@iifl.com\",\r\n    \"LocalIP\": \"123.123.123.123\",\r\n    \"PublicIP\": \"123.123.123.12\",\r\n    \"ContactNumber\": \"7769941110\"\r\n  }\r\n}",
		options: options{
			raw: raw{language: "json"},
		},
	}

	jsonReq, err := json.Marshal(e)

	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://dataservice.iifl.in/openapi/prod/LoginRequestMobileForVendor", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("key", "Ocp-Apim-Subscription-Key")
	req.Header.Set("value", "fc714d8e5b82438a93a95baa493ff45b")
	req.Header.Set("type", "text")

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
