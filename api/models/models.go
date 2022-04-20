package models

type UserTwitterDetails struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UserTwitterBase struct {
	Data UserTwitterDetails `json:"data"`
}

type Pair struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

type PairList []Pair

type UserTweetData struct {
	LikeCount int    `json:"like_count"`
	Text      string `json:"text"`
}

type UserTweet struct {
	PublicMetrics UserTweetMetrics `json:"public_metrics"`
	ID            string           `json:"id"`
	Text          string           `json:"text"`
}

type UserTweetMetrics struct {
	RetweetCount int `json:"retweet_count"`
	ReplyCount   int `json:"reply_count"`
	LikeCount    int `json:"like_count"`
	QuoteCount   int `json:"quote_count"`
}

type UserTweetMeta struct {
	OldestID    string `json:"oldest_id"`
	NewestID    string `json:"newest_id"`
	ResultCount int    `json:"result_count"`
	NextToken   string `json:"next_token"`
}

type UserTweetList struct {
	Data []UserTweet   `json:"data"`
	Meta UserTweetMeta `json:"meta"`
}
