package taskrunner

import (
	"errors"
	"log"
	"sync"

	"github.com/jiaruling/StreamMediaDevelopment/scheduler/dbops"
	"github.com/jiaruling/StreamMediaDevelopment/scheduler/oss"
)

var err error

func deletVideo(vid string) error {
	// err := os.Remove(VIDEO_DIR + "/" + vid)
	// if err != nil && !os.IsNotExist(err) {
	// 	return err
	// }
	ossfn := "test1/" + vid
	bn := "stream-media-videos"
	ok := oss.DeleteObject(ossfn, bn)
	if !ok {
		log.Printf("Deleting video error, oss operation failed")
		return errors.New("Deleting video error")
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("video clear dispatcher error: %v", err)
		return err
	}
	if len(res) == 0 {
		return errors.New("all tasks finished")
	}
	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExector(dc dataChan) error {
	errMap := &sync.Map{}
loop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				if err := deletVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break loop
		}
	}

	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}
