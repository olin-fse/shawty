package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

var _ = Describe("Utility Functions", func() {
	It("generates a random sequence of 5 characters", func() {
		code := RandSeq(5)
		Expect(code).To(HaveLen(5))
	})
})
