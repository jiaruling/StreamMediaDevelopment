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

type SignedIn struct {
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

type SimpleSession struct {
	Username string `json:"username"`
	TTL      int64  `json:"ttl"`
}

// Data model
type User struct {
	Id        int
	LoginName string
	Pwd       string
}

type UserInfo struct {
	Id int `json:"id"`
}

type NewVideo struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

type Comment struct {
	Id      string `json:"id"`
	VideoId string `json:"video_id"`
	User    string `json:"user"`
	Content string `json:"content"`
}

type Comments struct {
	Comments []*Comment `json:"comments"`
}

type NewComment struct {
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}
