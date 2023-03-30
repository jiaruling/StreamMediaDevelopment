package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/jiaruling/StreamMediaDevelopment/api/dbops"
	"github.com/jiaruling/StreamMediaDevelopment/api/defs"
	"github.com/jiaruling/StreamMediaDevelopment/api/response"
	"github.com/jiaruling/StreamMediaDevelopment/api/session"
	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, &ubody); err != nil {
		response.SendErrorResponse(w, &defs.ErrorRequestBodyParseFailed)
		return
	}
	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		response.SendErrorResponse(w, &defs.ErrorDBError)
		return
	}
	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		response.SendErrorResponse(w, &defs.ErrorSerializationError)
	} else {
		response.SendNormalResponse(w, string(resp), 201)
	}
	return
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("username")
	io.WriteString(w, username)
}
