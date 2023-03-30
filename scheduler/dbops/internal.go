package dbops

import "log"

func ReadVideoDeletionRecord(count int) ([]string, error) {
	var ids []string
	stmtOut, err := dbConn.Prepare(`SELECT id from video_delete limit ?`)
	if err != nil {
		return ids, err
	}
	defer stmtOut.Close()
	row, err := stmtOut.Query(count)
	if err != nil {
		log.Printf("Query VideoDeletionRecord error: %v", err)
		return ids, err
	}
	for row.Next() {
		var id string
		if err := row.Scan(&id); err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func DelVideoDeletionRecord(vid string) error {
	stmtDel, err := dbConn.Prepare(`DELETE FROM video_delete WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("DELETE VideoDeletionRecord error: %v", err)
		return err
	}
	return nil
}
