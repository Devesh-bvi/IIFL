package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/pbkdf2"

	"github.com/gorilla/mux"
)

func loginRequestMobileForVendor(w http.ResponseWriter, r *http.Request) {

	encryptVendor("YxjDEpOqxMSHJoXWb06CxjLNJxsZGNR0Z7Rg112orvaraf5UsWApy5tH8gCIEDO3", "38wqMxgHEaM")

	url := "https://dataservice.iifl.in/openapi/prod/LoginRequestMobileForVendor"
	method := "POST"

	payload := strings.NewReader(`{
		"Email_id":      "shubham.rajak@iifl.com",
		"LocalIP":       "123.123.123.123",
		"PublicIP":      "123.123.123.12",
		"ContactNumber": "7769941110",
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("appName", "IIFLMarSHUBHAMR")
	req.Header.Add("appVer", "1.0")
	req.Header.Add("key", "ThG6LZT8JJUWi5mD8OMHYTdsOmFzkebF")
	req.Header.Add("osName", "Android")
	req.Header.Add("requestCode", "IIFLMarRQLoginForVendor")
	req.Header.Add("userId", "Ae86Ls5iwTrPOm1w6PL2Cg==")
	req.Header.Add("password", "MnB1tFNCzzLAf8ezdonXVw==")
	req.Header.Add("Ocp-Apim-Subscription-Key", "fc714d8e5b82438a93a95baa493ff45b")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "IIFLMarcookie=o3jly30fl55xncvnok4zhb44")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func encryptVendor(MyKey string, text string) {
	//_key, err := b64.StdEncoding.DecodeString(data)

	//keySize := 16
	_iv := []byte{83, 71, 26, 58, 54, 35, 22, 11, 83, 71, 26, 58, 54, 35, 22, 11}

	//keyBuilder := hashPassword([]byte(MyKey), _iv, keySize)
	keyBuilder := pbkdf2.Key([]byte(MyKey), _iv, 10000, 32, sha1.New)
	plaintext := []byte(text)

	block, err := aes.NewCipher(keyBuilder)

	encKey := keyBuilder[:block.BlockSize()]
	authKey := keyBuilder[block.BlockSize() : block.BlockSize()*2]

	if err != nil {
		fmt.Println(err)
	}

	ciphertext := make([]byte, len(plaintext))
	fmt.Println(ciphertext)
	blockSize := block.BlockSize()
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	ciphertext = append(ciphertext, padText...)

	mode := cipher.NewCBCEncrypter(block, _iv)

	mode.CryptBlocks(encKey, authKey)

	//return b64.StdEncoding.EncodeToString([]byte("Gocrypt_" + string(salt) + string(ciphertext)))
	fmt.Printf("%x\n", ciphertext)
}

func hashPassword(str, salt []byte, keySize int) []byte {
	return pbkdf2.Key(str, salt, 10000, keySize, sha256.New)
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
