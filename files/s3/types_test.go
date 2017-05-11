package s3_test

import (
	// "net/url"
	. "build-dotenv/files/s3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Types", func() {
	It("", func() {
		Expect(
			(&Source{URLString: "s3://a/b"}).InitURLFromURLString().URL.Host).To(Equal("a"))
	})

	It("Good Base URL (no session)", func() {
		s3ish := (&Source{URLString: "s3://a/b"}).InitURLFromURLString().InitBucketPrefixFromURL()

		Expect(s3ish.Bucket).To(Equal("a"))
		Expect(s3ish.Prefix).To(Equal("b"))
	})

	It("Bad Base URL Panics", func() {
		Expect(func() {
			(&Source{URLString: "no-way"}).Init()
		}).To(Panic())
	})

	It("Bad Session Panics", func() {
		Skip("Have to mock / fake the AWS SDK?")
	})
})
