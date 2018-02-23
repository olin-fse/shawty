package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Insert Mapping", func() {
	It("inserts a valid mapping entry", func() {
		s := newTestStore()
		defer s.Close()

		_, err := s.CreateMapping("http://test.com", "test", false)
		Expect(err).To(BeNil())
	})
})
