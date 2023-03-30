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
	"github.com/julienschmidt/httprouter"
)

func PostComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !middleware.ValidateUser(w, r) {
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	cbody := &defs.NewComment{}
	if err := json.Unmarshal(reqBody, cbody); err != nil {
		log.Printf("%s", err)
		response.SendErrorResponse(w, &defs.ErrorRequestBodyParseFailed)
		return
	}

	vid := p.ByName("vid-id")
	if err := dbops.AddNewComments(vid, cbody.UserId, cbody.Content); err != nil {
		log.Printf("Error in PostComment: %s", err)
		response.SendErrorResponse(w, &defs.ErrorDBError)
	} else {
		response.SendNormalResponse(w, "ok", 201)
	}
}

func ShowComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !middleware.ValidateUser(w, r) {
		return
	}

	vid := p.ByName("vid-id")
	cm, err := dbops.ListComments(vid, 0, int(time.Now().Unix()))
	if err != nil {
		log.Printf("Error in ShowComments: %s", err)
		response.SendErrorResponse(w, &defs.ErrorDBError)
		return
	}

	cms := &defs.Comments{Comments: cm}
	if resp, err := json.Marshal(cms); err != nil {
		response.SendErrorResponse(w, &defs.ErrorSerializationError)
	} else {
		response.SendNormalResponse(w, string(resp), 200)
	}
}
