package main_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TrafficMirror", func() {
	Context("GET /", func() {
		It("returns OK", func() {
			response := doGetRequest("/")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})
	})

	Context("When GET /ok without parameters", func() {
		It("returns a 404 error", func() {
			response := doGetRequest("/ok")

			Expect(response.StatusCode).To(Equal(http.StatusNotFound))
		})
	})

	Context("When GET /ok with less than 2 parameters", func() {
		It("returns a 404 error", func() {
			response := doGetRequest("/ok/1")

			Expect(response.StatusCode).To(Equal(http.StatusNotFound))
		})
	})

	Context("When GET /ok/latency/bytes with correct parameters", func() {
		response := doGetRequest("/ok/100/200")
		It("returns OK", func() {
			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns a body filled with zeros 'bytes' times", func() {
			defer response.Body.Close()
			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			Expect(len(bodyBytes)).To(Equal(200))

			bodyString := string(bodyBytes)
			Expect(bodyString).To(Equal(strings.Repeat("0", 200)))
		})
	})

	Context("When GET /ok/latency/bytes with NaN latency", func() {
		response := doGetRequest("/ok/foo/200")
		It("returns OK", func() {
			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns a body filled with zeros 'bytes' times", func() {
			defer response.Body.Close()
			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			Expect(len(bodyBytes)).To(Equal(200))

			bodyString := string(bodyBytes)
			Expect(bodyString).To(Equal(strings.Repeat("0", 200)))
		})
	})

	Context("When GET /ok/latency/bytes with negative latency", func() {
		response := doGetRequest("/ok/-100/200")
		It("returns OK", func() {
			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns a body filled with zeros 'bytes' times", func() {
			defer response.Body.Close()
			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			Expect(len(bodyBytes)).To(Equal(200))

			bodyString := string(bodyBytes)
			Expect(bodyString).To(Equal(strings.Repeat("0", 200)))
		})
	})

	Context("When GET /ok/latency/bytes with NaN bytes", func() {
		response := doGetRequest("/ok/100/foo")
		It("returns OK", func() {
			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an empty body", func() {
			defer response.Body.Close()
			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			Expect(len(bodyBytes)).To(Equal(0))

			bodyString := string(bodyBytes)
			Expect(bodyString).To(Equal(""))
		})
	})

	Context("When GET /ok/latency/bytes with negative bytes", func() {
		response := doGetRequest("/ok/100/-200")
		It("returns OK", func() {
			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an empty body", func() {
			defer response.Body.Close()
			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			Expect(len(bodyBytes)).To(Equal(0))

			bodyString := string(bodyBytes)
			Expect(bodyString).To(Equal(""))
		})
	})

	Context("When GET /error without parameters", func() {
		It("returns a 404 error", func() {
			response := doGetRequest("/error")

			Expect(response.StatusCode).To(Equal(http.StatusNotFound))
		})
	})

	Context("When GET /error with correct parameters", func() {
		It("returns an error with the code", func() {
			response := doGetRequest("/error/503")

			Expect(response.StatusCode).To(Equal(503))
		})
	})

	Context("When GET /error with NaN code", func() {
		It("returns a 404 error", func() {
			response := doGetRequest("/error/foo")

			Expect(response.StatusCode).To(Equal(http.StatusNotFound))
		})
	})

	Context("When GET /error with negative code", func() {
		It("returns a 404 error", func() {
			response := doGetRequest("/error/-200")

			Expect(response.StatusCode).To(Equal(http.StatusNotFound))
		})
	})

})
