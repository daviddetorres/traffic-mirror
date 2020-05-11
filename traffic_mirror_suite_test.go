package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/daviddetorres/traffic-mirror"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTrafficMirror(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TrafficMirror Suite")
}

func doGetRequest(path string) *http.Response {
	request, _ := http.NewRequest("GET", path, nil)

	recorder := httptest.NewRecorder()
	router := NewRouter()
	router.ServeHTTP(recorder, request)

	return recorder.Result()
}
