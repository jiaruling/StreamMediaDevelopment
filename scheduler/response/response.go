package response

import (
	"io"
	"net/http"
)

func SendResponse(w http.ResponseWriter, sc int, response string) {
	w.WriteHeader(sc)
	io.WriteString(w, response)
}
