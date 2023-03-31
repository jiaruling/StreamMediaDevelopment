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

func ListVideoInfo(uname string, from, to int) ([]*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare(`SELECT video.id, video.user_id, video.name, video.display_time FROM video 
		INNER JOIN user ON video.user_id = user.id
		WHERE user.username = ? AND video.create_time > FROM_UNIXTIME(?) AND video.create_time <= FROM_UNIXTIME(?) 
		ORDER BY video.create_time DESC`)

	var res []*defs.VideoInfo

	if err != nil {
		return res, err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(uname, from, to)
	if err != nil {
		log.Printf("%s", err)
		return res, err
	}

	for rows.Next() {
		var id, name, ctime string
		var aid int
		if err := rows.Scan(&id, &aid, &name, &ctime); err != nil {
			return res, err
		}

		vi := &defs.VideoInfo{Id: id, UserId: aid, Name: name, DisplayTime: ctime}
		res = append(res, vi)
	}

	return res, nil
}
