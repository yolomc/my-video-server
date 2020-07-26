package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")

	t.Execute(w, nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid")
	vPath := VIDEO_DIR + vid + ".mp4"

	video, err := os.Open(vPath)
	if err != nil {
		log.Printf("open file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		log.Printf("The file is too large: %v", err)
		sendErrorResponse(w, http.StatusBadRequest, "The file is too large")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Get formfile error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	vid := p.ByName("vid")
	err = ioutil.WriteFile(VIDEO_DIR+vid+".mp4", data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Upload successfully")
}
