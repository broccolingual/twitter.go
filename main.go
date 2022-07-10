package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/broccolingual/twitter.go/model"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot load: %v", err)
	}
}

type error interface {
	Error() string
}

type Meta struct {
	Newest_id    string
	Oldest_id    string
	Result_count int
	Next_token   string
}

var API_KEY string
var API_SECRET_KEY string
var ACCESS_TOKEN string
var ACCESS_TOKEN_SECRET string
var BEARER_TOKEN string

func main() {
	loadEnv()
	API_KEY = os.Getenv("API_KEY")
	API_SECRET_KEY = os.Getenv("API_SECRET_KEY")
	ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")
	ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")
	BEARER_TOKEN = os.Getenv("BEARER_TOKEN")

	endpoint := "https://api.twitter.com/2/tweets/search/recent"
	query := map[string]interface{}{"query": "valorant", "max_results": 10}

	resp, _ := getRequest(endpoint, query)

	tweets, meta := mappingData(resp)

	for i, r := range tweets {
		byte, _ := json.Marshal(r.(map[string]interface{}))
		tweet := new(Tweet)
		json.Unmarshal(byte, &tweet)
		fmt.Println(i, tweet)
	}

	fmt.Println(meta)
}

func setHeader(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+BEARER_TOKEN)
	req.Header.Set("Content-Type", "application/json")
}

func setQuery(req *http.Request, query map[string]interface{}) {
	if query == nil {
		return
	}

	params := req.URL.Query()
	for k, v := range query {
		var value string
		switch v := v.(type) {
		case int:
			value = strconv.Itoa(v)
		default:
			value = v.(string)
		}
		params.Add(k, value)
	}
	req.URL.RawQuery = params.Encode()
}

func makeRequest(endpoint string, query map[string]interface{}) *http.Request {
	req, _ := http.NewRequest("GET", endpoint, nil)
	setHeader(req)
	setQuery(req, query)
	return req
}

func showRequestURL(req *http.Request) {
	fmt.Println(req.URL)
}

func getRequest(url string, query map[string]interface{}) ([]byte, error) {
	req := makeRequest(url, query)
	showRequestURL(req)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Request Error: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Fprintln(os.Stderr, "Response Error: ", err)
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	return body, nil
}

func mappingData(body []byte) ([]interface{}, Meta) {
	var respMap map[string]interface{}
	var metaData Meta
	json.Unmarshal(body, &respMap)
	dataList := respMap["data"].([]interface{})
	byte, _ := json.Marshal(respMap["meta"])
	json.Unmarshal(byte, &metaData)
	return dataList, metaData
}
