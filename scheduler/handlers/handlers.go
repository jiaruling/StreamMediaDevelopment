package handlers

import (
	"io"
	"net/http"

	"github.com/jiaruling/StreamMediaDevelopment/scheduler/dbops"
	"github.com/jiaruling/StreamMediaDevelopment/scheduler/response"
	"github.com/julienschmidt/httprouter"
)

func HealthCheck(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "healthcheck")
}

func VidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		response.SendResponse(w, http.StatusBadRequest, "video id should not empty")
		return
	}
	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		response.SendResponse(w, http.StatusInternalServerError, "internal server errror")
		return
	}
	response.SendResponse(w, http.StatusOK, "success")
	return
}
