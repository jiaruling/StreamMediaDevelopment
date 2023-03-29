package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"

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
	stmtOut, err := dbConn.Prepare(`select ttl, username from sessions where session_id=?`)
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

func ListSession() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare(`select * FROM sessions`)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var id, ttlstr, username string
		if err := rows.Scan(&id, &ttlstr, &username); err != nil {
			log.Printf("list sessions error: %s", err)
			break
		}
		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err != nil {
			ss := &defs.SimpleSession{Username: username, TTL: ttl}
			m.Store(id, ss)
			log.Printf("session id: %s, ttl: %d", id, ss.TTL)
		}
	}
	return m, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id= ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	if _, err := stmtOut.Exec(sid); err != nil {
		return err
	}
	return nil
}
