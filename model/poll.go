package model

import "time"

type Poll struct {
	Id               string
	Options          []PollOptions
	Duration_minutes int
	End_datetime     time.Time
	Voting_status    string
}

type PollOptions struct {
	Position int
	Label    string
	Votes    int
}
