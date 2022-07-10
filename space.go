package main

import "time"

type Space struct {
	Id                string
	State             string
	Created_at        time.Time
	Ended_at          time.Time
	Host_ids          []string
	Lang              string
	Is_ticketed       bool
	Invited_user_ids  []string
	Participant_count int
	Subscriber_count  int
	Scheduled_start   time.Time
	Speaker_ids       []string
	Started_at        time.Time
	Title             string
	Topic_ids         []string
	Updated_at        time.Time
}
