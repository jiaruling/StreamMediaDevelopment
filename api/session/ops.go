package session

import "sync"

type SimpleSession struct {
	Username string `json:"username"`
	TTL      int64  `json:"ttl"`
}

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB() {

}

func GenerateNewSessionId(un string) string {
	return ""
}

func IsSessionExpired(sid string) (string, bool) {
	return "", false
}
