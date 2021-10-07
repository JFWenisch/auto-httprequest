package main

import (
	"os"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := os.Getenv("AHTTP_URL")
	json := os.Getenv("AHTTP_BODY")
	method := os.Getenv("AHTTP_METHOD")
	contenttype := os.Getenv("AHTTP_CONTENTTYPE")

	fmt.Println("URL:>", url)
	fmt.Println("BODY:>", json)
	fmt.Println("METHOD:>", method)
	fmt.Println("TYPE:>", contenttype)
	var jsonStr = []byte(json)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", contenttype)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}