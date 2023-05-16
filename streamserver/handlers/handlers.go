package handlers

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jiaruling/StreamMediaDevelopment/streamserver/defs"
	"github.com/jiaruling/StreamMediaDevelopment/streamserver/oss"
	"github.com/jiaruling/StreamMediaDevelopment/streamserver/response"
	"github.com/julienschmidt/httprouter"
)

func HealthCheck(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "healthcheck")
}

// 视频文件读取
// 本地方案
func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := defs.VIDEO_DIR + "/" + vid

	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Open file err: %v", err)
		response.SendErrorResponse(w, http.StatusInternalServerError, "Internal error: ")
		return
	}
	defer video.Close()
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
}

// OSS方案
func StreamOssHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	targetUrl := "http://stream-media-videos.oss-cn-chengdu.aliyuncs.com/test1/" + p.ByName("vid-id")
	http.Redirect(w, r, targetUrl, 301)
}

// 视频文件上传
// 本地方案
func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, defs.MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(defs.MAX_UPLOAD_SIZE); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, "file is too big")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Form file err: %v", err)
		response.SendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file err: %v", err)
		response.SendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(defs.VIDEO_DIR+"/"+fn, data, 0666)
	if err != nil {
		log.Printf("Write file err: %v", err)
		response.SendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "uploaded successfully")
}

// OSS方案
func UploadOssHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, defs.MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(defs.MAX_UPLOAD_SIZE); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, "file is too big")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Form file err: %v", err)
		response.SendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file err: %v", err)
		response.SendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(defs.VIDEO_DIR+"/"+fn, data, 0666)
	if err != nil {
		log.Printf("Write file err: %v", err)
		response.SendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	// 上传至oss
	ossfn := "test1/" + fn
	byname := defs.VIDEO_DIR + "/" + fn
	bn := "stream-media-videos"
	ok := oss.UploadToOss(ossfn, byname, bn)
	if !ok {
		response.SendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	// 删除本地文件
	os.Remove(byname)

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "uploaded successfully")
}
