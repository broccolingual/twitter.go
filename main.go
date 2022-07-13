package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

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

var BEARER_TOKEN string

func main() {
	// load .env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	BEARER_TOKEN = os.Getenv("BEARER_TOKEN")

	// goroutine settings
	var wg sync.WaitGroup
	// var s = semaphore.NewWeighted(10)

	const MAX_RESULTS = 100
	const NUM_OF_LOOP = 100

	endpoint := getEndpointSearchRecentTweets()
	keyword := os.Args[1]

	var next_token string = ""
	var elapsed_times [NUM_OF_LOOP]time.Duration

	tweets := []Tweet{}
	tweet_user_ids := []string{}
	m_cnt := 0

	for i := 0; i < NUM_OF_LOOP; i++ {
		now := time.Now()
		query := map[string]interface{}{"query": keyword, "max_results": MAX_RESULTS, "tweet.fields": "attachments,author_id,lang", "media.fields": "url", "expansions": "attachments.media_keys"}
		if next_token != "" {
			query["next_token"] = next_token
		}
		resp, _ := getRequest(endpoint, query)
		ar := mappingData(resp)
		data := ar.Data
		meta := ar.Meta
		// includes := ar.Includes

		// for _, m := range includes.Media {
		// 	if m.Url != "" && m.Type == "photo" {
		// 		wg.Add(1)
		// 		go downloadFromURL(m.Url, &wg, s)
		// 		m_cnt++
		// 	}
		// }

		for _, r := range data {
			byte, _ := json.Marshal(r.(map[string]interface{}))
			tweet := new(Tweet)
			json.Unmarshal(byte, &tweet)
			if tweet.Lang == "ja" {
				tweets = append(tweets, *tweet)
				tweet_user_ids = append(tweet_user_ids, tweet.Author_id)
			}
		}

		next_token = meta.Next_token

		elapsed_times[i] = time.Since(now)
	}

	// for i, t := range tweets {
	// 	dt, _ := time.Parse(time.RFC3339, t.Created_at)
	// 	fmt.Printf("%d: %s\nSource: %s, %s\n%s\n\n", i, t.Id, t.Source, t.Lang, dt)
	// }

	var sumTimeDuration time.Duration
	for _, t := range elapsed_times {
		sumTimeDuration += t
	}

	wg.Wait()

	fmt.Printf("\nKeyword: %s\nElapsed Time(Request&Download): %vs\nGot %d tweets, Downloaded %d files\n\n", keyword, sumTimeDuration.Seconds(), len(tweets), m_cnt)

	// tweet user count
	sort.Strings(tweet_user_ids)
	count_user_ids := make(map[string]int)
	keys := make([]string, 0, len(count_user_ids))
	before_ids := ""

	for _, id := range tweet_user_ids {
		if id != before_ids {
			count_user_ids[id] = 0
		}
		count_user_ids[id]++
		before_ids = id
	}

	for key := range count_user_ids {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return count_user_ids[keys[i]] > count_user_ids[keys[j]]
	})

	// make highest count user map/list
	highest_count_user_ids := make(map[string]int)
	h_l := []string{}
	cnt := 0
	for _, k := range keys {
		if cnt > 10 {
			break
		}
		// fmt.Println(k, count_user_ids[k])
		highest_count_user_ids[k] = count_user_ids[k]
		h_l = append(h_l, k)
		cnt++
	}

	users := []User{}
	u_endpoint := getEndpointUsersByIDs(h_l)
	query := map[string]interface{}{}
	resp, _ := getRequest(u_endpoint, query)
	ar := mappingData(resp)
	data := ar.Data
	for _, r := range data {
		byte, _ := json.Marshal(r.(map[string]interface{}))
		user := new(User)
		json.Unmarshal(byte, &user)
		users = append(users, *user)
	}

	for i, user := range users {
		fmt.Printf("%d: [%d] @%s (%s)\n", i, highest_count_user_ids[user.Id], user.Username, user.Name)
	}
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
	u, err := url.Parse(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+BEARER_TOKEN)
	req.Header.Set("Content-Type", "application/json")
	setQuery(req, query)
	return req
}

func getRequest(url string, query map[string]interface{}) ([]byte, error) {
	req := makeRequest(url, query)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 400:
		log.Fatal(nil)
		break
	case 200:
		break
	default:
		fmt.Fprintln(os.Stderr, "Response Error: ", resp.Status)
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
