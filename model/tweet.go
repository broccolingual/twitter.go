package model

import "time"

type Tweet struct {
	Id   string
	Text string
	// TODO: attachments
	Author_id string
	// TODO: context_annotations
	Conversation_id string
	Created_at      time.Time
	// TODO: entities
	// TODO: geo
	In_reply_to_user_id string
	Lang                string
	// TODO: non_public_metrics
	// TODO: organic_metrics
	Possibly_sensitive bool
	// TODO: promoted_metrics
	// TODO: public_metrics
	// TODO: referenced_tweets
	Reply_settings string
	Source         string
	// TODO: withheld
}
