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
})
