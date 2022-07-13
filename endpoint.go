package main

import (
	"fmt"
	"strings"
)

const (
	ENDPOINT_BASE = "https://api.twitter.com"
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

func getEndpointCountRecentTweets() string {
	return fmt.Sprintf("%s/2/tweets/counts/recent", ENDPOINT_BASE)
}

func getEndpointCountAllTweets() string {
	return fmt.Sprintf("%s/2/tweets/counts/all", ENDPOINT_BASE)
}

func getEndpointTweets(tweet_ids []string) string {
	return fmt.Sprintf("%s/2/tweets?ids=%s", ENDPOINT_BASE, strings.Join(tweet_ids, ","))
}

func getEndpointTweet(tweet_id string) string {
	return fmt.Sprintf("%s/2/tweets/%s", ENDPOINT_BASE, tweet_id)
}

func getEndpointReverseChronologicalTimeline(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/timelines/reverse_chronological", ENDPOINT_BASE, user_id)
}

func getEndpointUserTweetTimeline(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/tweets", ENDPOINT_BASE, user_id)
}

func getEndpointUserMentionTimeline(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/mentions", ENDPOINT_BASE, user_id)
}

func getEndpointRetweets(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/retweeted_by", ENDPOINT_BASE, user_id)
}

func getEndpointQuoteTweets(tweet_id string) string {
	return fmt.Sprintf("%s/2/tweets/%s/quote_tweets", ENDPOINT_BASE, tweet_id)
}

func getEndpointLikedUsers(tweet_id string) string {
	return fmt.Sprintf("%s/2/tweets/%s/liking_users", ENDPOINT_BASE, tweet_id)
}

func getEndpointLikedTweets(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/liked_tweets", ENDPOINT_BASE, user_id)
}

func getEndpointBookmarks(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/bookmarks", ENDPOINT_BASE, user_id)
}

func getEndpointUsersByIDs(user_ids []string) string {
	return fmt.Sprintf("%s/2/users?ids=%s", ENDPOINT_BASE, strings.Join(user_ids, ","))
}

func getEndpointUserByID(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s", ENDPOINT_BASE, user_id)
}

func getEndpointUsersByUsernames(usernames []string) string {
	return fmt.Sprintf("%s/2/users/by?usernames=%s", ENDPOINT_BASE, strings.Join(usernames, ","))
}

func getEndpointUserByUsername(username string) string {
	return fmt.Sprintf("%s/2/users/by/username/%s", ENDPOINT_BASE, username)
}

func getEndpointAuthUser() string {
	return fmt.Sprintf("%s/2/users/me", ENDPOINT_BASE)
}

func getEndpointFollowing(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/following", ENDPOINT_BASE, user_id)
}

func getEndpointFollowers(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/followers", ENDPOINT_BASE, user_id)
}

func getEndpointBlocking(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/blocking", ENDPOINT_BASE, user_id)
}

func getEndpointMuting(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/muting", ENDPOINT_BASE, user_id)
}
