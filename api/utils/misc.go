package utils

import (
	"log"
	"net/http"
)

func SendDeleteVideoRequest(id string) {
	addr := "127.0.0.1" + ":8002" // todo 有待优化
	url := "http://" + addr + "/video-delete-record/" + id
	_, err := http.Get(url)
	if err != nil {
		log.Printf("Sending deleting video request error: %s", err)
	}
}
