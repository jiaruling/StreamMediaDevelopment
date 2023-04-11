package handlers

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/jiaruling/StreamMediaDevelopment/web/client"
	"github.com/jiaruling/StreamMediaDevelopment/web/config"
	"github.com/jiaruling/StreamMediaDevelopment/web/defs"

	"github.com/julienschmidt/httprouter"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

// done: 完成
func HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		p := &HomePage{Name: "lihaoya"}
		t, e := template.ParseFiles("./template/home.html")
		if e != nil {
			log.Printf("Parsing template home.html error: %s", e)
			return
		}

		t.Execute(w, p)
		return
	}

	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return
	}
}

// done: 完成
func UserHomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cname, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fname := r.FormValue("username")

	var p *UserPage
	if len(cname.Value) != 0 {
		p = &UserPage{Name: cname.Value}
	} else if len(fname) != 0 {
		p = &UserPage{Name: fname}
	}

	t, e := template.ParseFiles("./template/userhome.html")
	if e != nil {
		log.Printf("Parsing userhome.html error: %s", e)
		return
	}

	t.Execute(w, p)
}

// done: api透传
func ApiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(defs.ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apibody := &defs.ApiBody{}
	if err := json.Unmarshal(res, apibody); err != nil {
		re, _ := json.Marshal(defs.ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
		return
	}

	client.Request(apibody, w, r)
	defer r.Body.Close()
}

// done: proxy转发
func ProxyVideoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u, _ := url.Parse("http://" + config.GetLbAddr() + ":8003/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}

// done: proxy转发
func ProxyUploadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u, _ := url.Parse("http://" + config.GetLbAddr() + ":8003/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}
