package main

import (
	"net/http"
	"testing"

	"github.com/difoul/go-rest-app/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRepositories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OSDF Tool")
}

/*
   Init The rooter
*/
var router = SetupRouter()
var AUTH = "Basic dGVzdDp0ZXN0"

var _ = Describe("Hello service", func() {
	Context("GET /v1/hello", func() {
		It("Retuns a status code of 200", func() {
			w := tests.PerformRequest(router, "GET", "/v1/hello", AUTH, "")
			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})
})

var _ = Describe("Quotes service", func() {
	Context("GET /v1/quotes/random", func() {
		It("Retuns a status code of 200", func() {
			w := tests.PerformRequest(router, "GET", "/v1/quotes/random", AUTH, "")
			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})
	Context("GET /v1/quotes/quote/:id", func() {
		Context("with a valid id", func() {
			It("Retuns a status code of 200", func() {
				w := tests.PerformRequest(router, "GET", "/v1/quotes/quote/0", AUTH, "")
				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})
		Context("with a valid id", func() {
			It("Retuns a status code of 404", func() {
				w := tests.PerformRequest(router, "GET", "/v1/quotes/quote/9", AUTH, "")
				Expect(w.Code).To(Equal(http.StatusNotFound))
			})
		})
		Context("with a non valid id", func() {
			It("Retuns a status code of 400", func() {
				w := tests.PerformRequest(router, "GET", "/v1/quotes/quote/abc", AUTH, "")
				Expect(w.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
