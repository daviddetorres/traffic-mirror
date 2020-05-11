package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

// NewRouter creates a new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	registerRouter(router)
	return router
}

func registerRouter(router *httprouter.Router) {
	router.GET("/ok/:milliseconds/:bytes", returnOk)
	router.GET("/error/:code", returnError)
	router.GET("/", returnHome)
}

func returnHome(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "<H1>Traffic Mirror</H1>\n")
	fmt.Fprintf(w, "<b>Usage:</b><br>\n")
	fmt.Fprintf(w, "/ok/milliseconds/bytes - Returns CODE:200 ")
	fmt.Fprintf(w, "with random latency from 0 to 'milliseconds' ")
	fmt.Fprintf(w, "and with a body filled with 0's 'bytes' times.<br>\n")
	fmt.Fprintf(w, "/error/code - Returns error with the specified 'code'.<br>\n")
	return
}

func returnOk(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	payloadSize, err := strconv.Atoi(ps.ByName("bytes"))
	if err != nil {
		log.Printf("Payload size must be an integer. Received %s", ps.ByName("bytes"))
		return
	}
	if payloadSize < 0 {
		payloadSize = 0
	}

	payload := strings.Repeat("0", payloadSize)

	latency, err := strconv.Atoi(ps.ByName("milliseconds"))
	if err != nil {
		log.Printf("Latency must be an integer. Received %s", ps.ByName("milliseconds"))
		latency = 1
	}
	if latency <= 0 {
		latency = 1
	}

	rand.Seed(time.Now().UnixNano())
	randomLatency := rand.Intn(latency)
	time.Sleep(time.Duration(randomLatency) * time.Millisecond)

	fmt.Fprintf(w, "%s", payload)
	return
}

func returnError(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	errorCode, err := strconv.Atoi(ps.ByName("code"))
	if err != nil {
		log.Printf("Payload size must be an integer. Received %s", ps.ByName("code"))
		errorCode = 404
	}
	if errorCode < 0 {
		errorCode = 404
	}
	w.WriteHeader(errorCode)

	fmt.Fprintf(w, "Error generated: %s\n", ps.ByName("code"))
	return
}

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
