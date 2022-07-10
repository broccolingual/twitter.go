package main

import "time"

type User struct {
	Id          string
	Name        string
	Username    string
	Created_at  time.Time
	Description string
	// TODO: entities
	Location          string
	Pinned_tweet_id   string
	Profile_image_url string
	Protected         bool
	// TODO: public_metrics
	Url      string
	Verified bool
	// TODO: withheld
}
