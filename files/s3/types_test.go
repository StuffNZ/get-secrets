package s3_test

import (
	. "build-dotenv/files/s3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Types", func() {
	It("Bad URL Panics", func() {
		Expect(func() {
			(&Source{Url: "no-way"}).Init()
		}).To(Panic())
	})
})
