package url_test

import (
	// "net/url"
	. "build-dotenv/files/s3/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The S3 URL type", func() {
	It("panics with an empty URL", func() {
		Expect(func() {
			New(Config{})
		}).To(Panic())
	})

	It("panics with an invalid URL", func() {
		Expect(func() {
			New(Config{URL: "..why::"})
		}).To(Panic())
	})

	It("returns correct Bucket and Prefix from a good Base URL", func() {
		s3ish := New(Config{URL: "s3://a/b"})

		Expect(s3ish.Bucket()).To(Equal("a"))
		Expect(s3ish.Prefix()).To(Equal("b"))
	})

	It("returns the correct Bucket when given a Bucket", func() {
		Expect(New(Config{Bucket: "Yerp"}).Bucket()).To(Equal("Yerp"))
	})

	It("returns the correct Prefix when given a Prefix", func() {
		Expect(New(Config{Bucket: "Yerp", Prefix: "Derp"}).Prefix()).To(Equal("Derp"))
	})

	It("returns a '/' Prefix when given an empty Prefix", func() {
		Expect(New(Config{Bucket: "Yerp"}).Prefix()).To(Equal("/"))
	})

	It("panics with an empty Bucket and 'full' Prefix", func() {
		Expect(func() {
			New(Config{Prefix: "Derp"})
		}).To(Panic())
	})
})
