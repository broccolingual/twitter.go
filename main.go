package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
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

type Includes struct {
	Media []Media
	Tweet []Tweet
}

type ApiResponse struct {
	Data     []interface{}
	Meta     Meta
	Includes Includes
}

// var API_KEY string
// var API_SECRET_KEY string
// var ACCESS_TOKEN string
// var ACCESS_TOKEN_SECRET string
var BEARER_TOKEN string

func main() {
	loadEnv()
	// API_KEY = os.Getenv("API_KEY")
	// API_SECRET_KEY = os.Getenv("API_SECRET_KEY")
	// ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")
	// ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")
	BEARER_TOKEN = os.Getenv("BEARER_TOKEN") // TODO: Bearer Tokenの自動取得
	MAX_RESULTS := 100

	endpoint := getEndpointCountRecentTweets()
	keyword := `"valorant"`

	var next_token string
	var elapsed_times [1]time.Duration

	tweets := []Tweet{}
	media_urls := []string{}

	for i := 0; i < 1; i++ {
		now := time.Now()
		if i == 0 {
			query := map[string]interface{}{"query": keyword, "max_results": MAX_RESULTS, "tweet.fields": "attachments,source,lang,created_at", "media.fields": "url,variants", "expansions": "attachments.media_keys"}
			resp, _ := getRequest(endpoint, query)
			ar := mappingData(resp)
			data := ar.Data
			meta := ar.Meta
			includes := ar.Includes

			for _, m := range includes.Media {
				if m.Url != "" && m.Type == "photo" {
					media_urls = append(media_urls, m.Url)
				}
			}

			for _, r := range data {
				byte, _ := json.Marshal(r.(map[string]interface{}))
				tweet := new(Tweet)
				json.Unmarshal(byte, &tweet)
				tweets = append(tweets, *tweet)
			}

			next_token = meta.Next_token
		} else {
			query := map[string]interface{}{"query": keyword, "max_results": MAX_RESULTS, "next_token": next_token, "tweet.fields": "attachments,source,lang,created_at", "media.fields": "url,variants", "expansions": "attachments.media_keys"}
			resp, _ := getRequest(endpoint, query)
			ar := mappingData(resp)
			data := ar.Data
			meta := ar.Meta
			includes := ar.Includes

			for _, m := range includes.Media {
				if m.Url != "" && m.Type == "photo" {
					media_urls = append(media_urls, m.Url)
				}
			}

			for _, r := range data {
				byte, _ := json.Marshal(r.(map[string]interface{}))
				tweet := new(Tweet)
				json.Unmarshal(byte, &tweet)
				tweets = append(tweets, *tweet)
			}

			next_token = meta.Next_token
		}
		elapsed_times[i] = time.Since(now)
	}

	for i, t := range tweets {
		dt, _ := time.Parse(time.RFC3339, t.Created_at)
		fmt.Printf("%d: %s\nSource: %s, %s\n%s\n\n", i, t.Id, t.Source, t.Lang, dt)
	}

	var sumTimeDuration time.Duration
	for i, t := range elapsed_times {
		fmt.Printf("Req %d: %vs\n", i, t.Seconds())
		sumTimeDuration += t
	}
	now := time.Now()
	downloadParallel(media_urls)
	elapsed_time := time.Since(now)

	fmt.Printf("Elapsed Time(Request): %vs\nGot %d tweets\n", sumTimeDuration.Seconds(), len(tweets))
	fmt.Printf("Elapsed Time(Download): %vs\nDownloaded %d files\n", elapsed_time.Seconds(), len(media_urls))
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

func mappingData(body []byte) *ApiResponse {
	ar := new(ApiResponse)
	json.Unmarshal(body, ar)
	return ar
}
