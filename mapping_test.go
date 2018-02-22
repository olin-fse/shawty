package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

//func TestCart(t *testing.T) {
//	RegisterFailHandler(Fail)
//	RunSpecs(t, "Test Suite")
//}

var _ = Describe("Insert Mapping", func() {
	It("inserts a valid mapping entry", func() {
		s := newTestStore()
		defer s.Close()

		// Insert an item
		_, err := s.CreateMapping("http://test.com", "test", false)
		Expect(err).To(BeNil())
	})
})

func TestCreateItem(t *testing.T) {

}
