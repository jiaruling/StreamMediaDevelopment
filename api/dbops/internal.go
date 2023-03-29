package dbops

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/jiaruling/StreamMediaDevelopment/api/defs"
)

func AddNewSession(sid string, ttl int64, username string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare(`INSERT INTO sessions(session_id,TTL, username) values (?, ?, ?)`)
	if err != nil {
		log.Printf("AddSessions %s", err.Error())
		return err
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(sid, ttlstr, username)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare(`select ttl, username from sessions where sid=?`)
	if err != nil {
		log.Printf("RetrieveSessions %s", err.Error())
		return nil, err
	}
	defer stmtOut.Close()
	var ttl, username string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &username)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = username
		return ss, nil
	} else {
		return nil, err
	}
}

func ListSession() {

}

func DeleteSession() {}
