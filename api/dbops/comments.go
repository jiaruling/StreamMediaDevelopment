package dbops

import (
	"log"

	"github.com/jiaruling/StreamMediaDevelopment/api/defs"
	uuid "github.com/satori/go.uuid"
)

func AddNewComments(vid string, aid int, content string) error {
	u1 := uuid.NewV4().String()
	stmtIns, err := dbConn.Prepare(`INSERT INTO comments(id, video_id, user_id, content) VALUES (?, ?, ?, ?)`)
	if err != nil {
		log.Printf("AddComments %s", err.Error())
		return err
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(u1, vid, aid, content)
	if err != nil {
		return err
	}
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comments, error) {
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, user.username, comments.content FROM comments
									 INNER JOIN user ON comments.user_id = user.id
									 WHERE comments.video_id= ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)
	if err != nil {
		log.Printf("ListComments %s", err.Error())
		return nil, err
	}
	defer stmtOut.Close()
	var res []*defs.Comments
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comments{Id: id, VideoId: vid, User: name, Content: content}
		res = append(res, c)
	}
	return res, nil
}
