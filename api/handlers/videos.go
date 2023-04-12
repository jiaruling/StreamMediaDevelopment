package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jiaruling/StreamMediaDevelopment/api/dbops"
	"github.com/jiaruling/StreamMediaDevelopment/api/defs"
	"github.com/jiaruling/StreamMediaDevelopment/api/middleware"
	"github.com/jiaruling/StreamMediaDevelopment/api/response"
	"github.com/jiaruling/StreamMediaDevelopment/api/utils"
	"github.com/julienschmidt/httprouter"
)

func AddNewVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !middleware.ValidateUser(w, r) {
		log.Printf("Unathorized user \n")
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	nvbody := &defs.NewVideo{}
	if err := json.Unmarshal(res, nvbody); err != nil {
		log.Printf("%s", err)
		response.SendErrorResponse(w, &defs.ErrorRequestBodyParseFailed)
		return
	}

	vi, err := dbops.AddNewVideos(nvbody.UserId, nvbody.Name)
	log.Printf("Author id : %d, name: %s \n", nvbody.UserId, nvbody.Name)
	if err != nil {
		log.Printf("Error in AddNewVideo: %s", err)
		response.SendErrorResponse(w, &defs.ErrorDBError)
		return
	}

	if resp, err := json.Marshal(vi); err != nil {
		response.SendErrorResponse(w, &defs.ErrorSerializationError)
	} else {
		response.SendNormalResponse(w, string(resp), 201)
	}
}

func ListAllVideos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !middleware.ValidateUser(w, r) {
		return
	}

	uname := p.ByName("username")
	vs, err := dbops.ListVideoInfo(uname, 0, int(time.Now().Unix()))
	if err != nil {
		log.Printf("Error in ListAllvideos: %s", err)
		response.SendErrorResponse(w, &defs.ErrorDBError)
		return
	}

	vsi := &defs.VideosInfo{Videos: vs}
	if resp, err := json.Marshal(vsi); err != nil {
		response.SendErrorResponse(w, &defs.ErrorSerializationError)
	} else {
		response.SendNormalResponse(w, string(resp), 200)
	}
}

func DeleteVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !middleware.ValidateUser(w, r) {
		return
	}

	vid := p.ByName("vid-id")
	// 删除视频
	err := dbops.DeleteVideos(vid)
	if err != nil {
		log.Printf("Error in DeletVideo: %s", err)
		response.SendErrorResponse(w, &defs.ErrorDBError)
		return
	}
	// 删除评论
	dbops.DeleteComments(vid)

	go utils.SendDeleteVideoRequest(vid)
	response.SendNormalResponse(w, "", 204)
}
