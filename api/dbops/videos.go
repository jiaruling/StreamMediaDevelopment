package dbops

import (
	"database/sql"
	"log"
	"time"

	"github.com/jiaruling/StreamMediaDevelopment/api/defs"
	uuid "github.com/satori/go.uuid"
)

func AddNewVideos(aid int, name string) (*defs.VideoInfo, error) {
	u1 := uuid.NewV4().String()
	ctime := time.Now().Format("2006-01-02T15:04:05")
	stmtIns, err := dbConn.Prepare(`INSERT INTO video(id, user_id, name, display_time) VALUES (?, ?, ?, ?)`)
	if err != nil {
		log.Printf("AddVideos %s", err.Error())
		return nil, err
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(u1, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: u1, UserId: aid, Name: name, DisplayTime: ctime}
	return res, nil
}

func DeleteVideos(vid string) error {
	stmtDel, err := dbConn.Prepare(`DELETE FROM video WHERE id = ?`)
	if err != nil {
		log.Printf("DeleteVideos %s", err.Error())
		return err
	}
	defer stmtDel.Close()
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	return nil
}

func GetVideos(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare(`select user_id, name, display_time from video where id=?`)
	if err != nil {
		log.Printf("GetVideos %s", err.Error())
		return nil, err
	}
	defer stmtOut.Close()
	var userId int
	var name, displayTime string
	err = stmtOut.QueryRow(vid).Scan(&userId, &name, &displayTime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	res := &defs.VideoInfo{Id: vid, UserId: userId, Name: name, DisplayTime: displayTime}
	return res, nil
}
