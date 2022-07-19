package endpoint

import (
	"fmt"
	"strings"
)

const (
	ENDPOINT_BASE = "https://api.twitter.com"
)

func GetEndpointOauth2Token() string {
	return fmt.Sprintf("%s/oauth2/token", ENDPOINT_BASE)
}

func GetEndpointSearchRecentTweets() string {
	return fmt.Sprintf("%s/2/tweets/search/recent", ENDPOINT_BASE)
}

func GetEndpointSearchAllTweets() string {
	return fmt.Sprintf("%s/2/tweets/search/all", ENDPOINT_BASE)
}

func GetEndpointCountRecentTweets() string {
	return fmt.Sprintf("%s/2/tweets/counts/recent", ENDPOINT_BASE)
}

func GetEndpointCountAllTweets() string {
	return fmt.Sprintf("%s/2/tweets/counts/all", ENDPOINT_BASE)
}

func GetEndpointTweets(tweet_ids []string) string {
	return fmt.Sprintf("%s/2/tweets?ids=%s", ENDPOINT_BASE, strings.Join(tweet_ids, ","))
}

func GetEndpointTweet(tweet_id string) string {
	return fmt.Sprintf("%s/2/tweets/%s", ENDPOINT_BASE, tweet_id)
}

func GetEndpointReverseChronologicalTimeline(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/timelines/reverse_chronological", ENDPOINT_BASE, user_id)
}

func GetEndpointUserTweetTimeline(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/tweets", ENDPOINT_BASE, user_id)
}

func GetEndpointUserMentionTimeline(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/mentions", ENDPOINT_BASE, user_id)
}

func GetEndpointRetweets(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/retweeted_by", ENDPOINT_BASE, user_id)
}

func GetEndpointQuoteTweets(tweet_id string) string {
	return fmt.Sprintf("%s/2/tweets/%s/quote_tweets", ENDPOINT_BASE, tweet_id)
}

func GetEndpointLikedUsers(tweet_id string) string {
	return fmt.Sprintf("%s/2/tweets/%s/liking_users", ENDPOINT_BASE, tweet_id)
}

func GetEndpointLikedTweets(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/liked_tweets", ENDPOINT_BASE, user_id)
}

func GetEndpointBookmarks(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/bookmarks", ENDPOINT_BASE, user_id)
}

func GetEndpointUsersByIDs(user_ids []string) string {
	return fmt.Sprintf("%s/2/users?ids=%s", ENDPOINT_BASE, strings.Join(user_ids, ","))
}

func GetEndpointUserByID(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s", ENDPOINT_BASE, user_id)
}

func GetEndpointUsersByUsernames(usernames []string) string {
	return fmt.Sprintf("%s/2/users/by?usernames=%s", ENDPOINT_BASE, strings.Join(usernames, ","))
}

func GetEndpointUserByUsername(username string) string {
	return fmt.Sprintf("%s/2/users/by/username/%s", ENDPOINT_BASE, username)
}

func GetEndpointAuthUser() string {
	return fmt.Sprintf("%s/2/users/me", ENDPOINT_BASE)
}

func GetEndpointFollowing(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/following", ENDPOINT_BASE, user_id)
}

func GetEndpointFollowers(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/followers", ENDPOINT_BASE, user_id)
}

func GetEndpointBlocking(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/blocking", ENDPOINT_BASE, user_id)
}

func GetEndpointMuting(user_id string) string {
	return fmt.Sprintf("%s/2/users/%s/muting", ENDPOINT_BASE, user_id)
}

