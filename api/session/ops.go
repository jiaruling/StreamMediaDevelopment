package session

import (
	"sync"
	"time"

	"github.com/jiaruling/StreamMediaDevelopment/api/dbops"
	"github.com/jiaruling/StreamMediaDevelopment/api/defs"
	uuid "github.com/satori/go.uuid"
)

type SimpleSession struct {
	Username string `json:"username"`
	TTL      int64  `json:"ttl"`
}

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB() {
	r, err := dbops.ListSession()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id := uuid.NewV4().String()
	ct := time.Now().Unix()
	ttl := ct + 30*60 // Serverside session valid time: 30min
	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.AddNewSession(id, ttl, un)
	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	ct := time.Now().Unix()
	if ok {
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	} else {
		ss, err := dbops.RetrieveSession(sid)
		if err != nil || ss == nil {
			return "", true
		}
		if ss.TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		sessionMap.Store(sid, ss)
		return ss.Username, false
	}
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}
