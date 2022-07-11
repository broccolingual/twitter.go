package main

type Tweet struct {
	Id          string
	Text        string
	Attachments TweetAttachments
	Author_id   string
	// TODO: context_annotations
	Conversation_id string
	Created_at      string
	// TODO: entities
	Geo                 TweetGeo
	In_reply_to_user_id string
	Lang                string
	// TODO: non_public_metrics
	// TODO: organic_metrics
	Possibly_sensitive bool
	// TODO: promoted_metrics
	// TODO: public_metrics
	Referenced_tweets []TweetReferencedTweets
	Reply_settings    string
	Source            string
	// TODO: withheld
}

type TweetAttachments struct {
	Poll_ids   []string
	Media_keys []string
}

type TweetGeo struct {
	Coodinates TweetGeoCoodinates
	Place_id   string
}

type TweetGeoCoodinates struct {
	Type        string
	Coordinates []float64
}

type TweetReferencedTweets struct {
	Type string
	Id   string
}
