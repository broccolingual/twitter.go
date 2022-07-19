package model

import "time"

type List struct {
	Id             string
	Name           string
	Created_at     time.Time
	Description    string
	Follower_count int
	Member_count   int
	Private        bool
	Owner_id       string
}
