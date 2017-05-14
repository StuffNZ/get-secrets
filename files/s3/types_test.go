package s3_test

import (
	// "net/url"
	. "build-dotenv/files/s3"
	// urlish "build-dotenv/files/s3/s3url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The S3 type", func() {
	It("", func() {
		Skip("Placeholder")
		Expect(true).To(BeTrue())
	})

	// It("stores the right S3 URL config", func() {
	// 	url := urlish.New(urlish.Config{URL: "s3://a/b"})
	// 	s3 := New(Config{Source: url}).
	// })

	It("panics with a bad Session", func() {
		Skip("Have to mock / fake the AWS SDK?")
	})

	// It("s with empty AWS Region", func() {
	// 	Expect
	// })
})
