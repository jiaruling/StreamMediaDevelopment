package defs

//requests
type UserCredential struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

// data model
type VideoInfo struct {
	Id          string `json:"id"`
	UserId      int    `json:"user_id"`
	Name        string `json:"name"`
	DisplayTime string `json:"display_time"`
}

type Comments struct {
	Id      string `json:"id"`
	VideoId string `json:"video_id"`
	User    string `json:"user"`
	Content string `json:"content"`
}

type SimpleSession struct {
	Username string `json:"username"`
	TTL      int64  `json:"ttl"`
}
