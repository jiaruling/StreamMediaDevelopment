package dbops

import (
	"log"
	"time"

	"github.com/jiaruling/StreamMediaDevelopment/api/defs"
	uuid "github.com/satori/go.uuid"
)

func AddNewComments(vid string, aid int, content string) error {
	u1 := uuid.NewV4().String()
	ctime := time.Now().Format("2006-01-02T15:04:05")
	stmtIns, err := dbConn.Prepare(`INSERT INTO comments(id, video_id, user_id, content, time) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		log.Printf("AddComments %s", err.Error())
		return err
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(u1, vid, aid, content, ctime)
	if err != nil {
		return err
	}
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, user.username, comments.content FROM comments
									 INNER JOIN user ON comments.user_id = user.id
									 WHERE comments.video_id= ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)
									 ORDER BY comments.time DESC`)
	if err != nil {
		log.Printf("ListComments %s", err.Error())
		return nil, err
	}
	defer stmtOut.Close()
	var res []*defs.Comment
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{Id: id, VideoId: vid, User: name, Content: content}
		res = append(res, c)
	}
	return res, nil
}

func DeleteComments(vid string) error {
	stmtIns, err := dbConn.Prepare(`delete from comments where video_id=?;`)
	if err != nil {
		log.Printf("AddComments %s", err.Error())
		return err
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(vid)
	if err != nil {
		return err
	}
	return nil
}
