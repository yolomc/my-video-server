package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/yolomc/my-video-server/api/dbops"
	"github.com/yolomc/my-video-server/api/defs"
	"github.com/yolomc/my-video-server/api/gredis"
	"github.com/yolomc/my-video-server/api/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		utils.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		utils.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	sessionID := uuid.New().String()
	if err := gredis.Set(sessionID, ubody.Username); err != nil {
		utils.SendErrorResponse(w, defs.ErrorRedisError)
		return
	}

	su := &defs.SignedUp{Success: true, SessionId: sessionID}
	resp, err := json.Marshal(su)
	if err != nil {
		utils.SendErrorResponse(w, defs.ErrorServiceError)
	}

	utils.SendNormalResponse(w, string(resp), 201)
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("username")
	io.WriteString(w, username)
}
