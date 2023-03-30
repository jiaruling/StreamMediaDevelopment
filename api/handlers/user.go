package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jiaruling/StreamMediaDevelopment/api/dbops"
	"github.com/jiaruling/StreamMediaDevelopment/api/defs"
	"github.com/jiaruling/StreamMediaDevelopment/api/middleware"
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
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	log.Printf("%s", res)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		log.Printf("%s", err)
		response.SendErrorResponse(w, &defs.ErrorRequestBodyParseFailed)
		return
	}

	// Validate the request body
	uname := p.ByName("username")
	log.Printf("Login url name: %s", uname)
	log.Printf("Login body name: %s", ubody.Username)
	if uname != ubody.Username {
		response.SendErrorResponse(w, &defs.ErrorNotAuthUser)
		return
	}

	log.Printf("%s", ubody.Username)
	pwd, err := dbops.GetUserCredential(ubody.Username)
	log.Printf("Login pwd: %s", pwd)
	log.Printf("Login body pwd: %s", ubody.Pwd)
	if err != nil || len(pwd) == 0 || pwd != ubody.Pwd {
		response.SendErrorResponse(w, &defs.ErrorNotAuthUser)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	si := &defs.SignedIn{Success: true, SessionId: id}
	if resp, err := json.Marshal(si); err != nil {
		response.SendErrorResponse(w, &defs.ErrorSerializationError)
	} else {
		response.SendNormalResponse(w, string(resp), 200)
	}
}

func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !middleware.ValidateUser(w, r) {
		log.Printf("Unathorized user \n")
		return
	}

	uname := p.ByName("username")
	u, err := dbops.GetUser(uname)
	if err != nil {
		log.Printf("Error in GetUserInfo: %s", err)
		response.SendErrorResponse(w, &defs.ErrorDBError)
		return
	}

	ui := &defs.UserInfo{Id: u.Id}
	if resp, err := json.Marshal(ui); err != nil {
		response.SendErrorResponse(w, &defs.ErrorSerializationError)
	} else {
		response.SendNormalResponse(w, string(resp), 200)
	}
}
