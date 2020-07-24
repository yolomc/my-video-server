package middleware

import (
	"net/http"

	"github.com/yolomc/my-video-server/api/config"
	"github.com/yolomc/my-video-server/api/defs"
	"github.com/yolomc/my-video-server/api/gredis"
	"github.com/yolomc/my-video-server/api/utils"
)

func ValidateUserSession(r *http.Request) bool {
	sid := r.Header.Get(config.HTTP_HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, err := gredis.Get(sid)
	if err != nil || uname == "" {
		return false
	}

	r.Header.Add(config.HTTP_HEADER_FIELD_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(config.HTTP_HEADER_FIELD_UNAME)
	if uname == "" {
		utils.SendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
