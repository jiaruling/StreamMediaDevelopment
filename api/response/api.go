package response

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jiaruling/StreamMediaDevelopment/api/defs"
)

func SendErrorResponse(w http.ResponseWriter, errResp *defs.ErrorResponse) {
	w.WriteHeader(errResp.HttpSC)
	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
	return
}

func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
	return
}
