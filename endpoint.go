package main

import "fmt"

const (
	ENDPOINT_BASE = "https://api.twitter.com"

	ENDPOINT_TWEET_IDS = "/2/tweets"
	ENDPOINT_TWEET_ID  = "/2/tweets/"

	ENDPOINT_USER_IDS  = "/2/users"
	ENDPOINT_USER_ID   = "/2/users/"
	ENDPOINT_USERNAMES = "/2/users/by"
	ENDPOINT_USERNAME  = "/2/users/by/username/"
	ENDPOINT_AUTH_USER = "/2/users/me"
)

func getEndpointOauth2Token() string {
	return fmt.Sprintf("%s/oauth2/token", ENDPOINT_BASE)
}

func getEndpointSearchRecentTweets() string {
	return fmt.Sprintf("%s/2/tweets/search/recent", ENDPOINT_BASE)
}

func getEndpointSearchAllTweets() string {
	return fmt.Sprintf("%s/2/tweets/search/all", ENDPOINT_BASE)
}

func getEndpointCountsRecentTweets() string {
	return fmt.Sprintf("%s/2/tweets/counts/recent", ENDPOINT_BASE)
}

func getEndpointCountsAllTweets() string {
	return fmt.Sprintf("%s/2/tweets/counts/all", ENDPOINT_BASE)
}

func getEndpoint() string {
	return fmt.Sprintf("%s", ENDPOINT_BASE)
}
