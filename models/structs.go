package models

type PollOption struct {
	Qid       int    `json:"qid" `
	Option_id int    `json:"option_id"`
	Option    string `json:"option"`
	Votes     int    `json:"votes"`
}

type PollQuestion struct {
	Qid      int          `json:"qid"`
	Question string       `json:"question"`
	Options  []PollOption `json:"options"`
}

type PollVote struct {
	Qid       int `json:"qid"`
	Option_id int `json:"option_id"`
}

type PollParticipation struct {
	PollVotes []PollVote `json:"poll_votes"`
}
