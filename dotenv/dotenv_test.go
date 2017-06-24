package dotenv_test

import (
	. "build-dotenv/dotenv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The Dotenv type", func() {
	It("initialises an empty map", func() {
		env := New()

		Expect(env.Combine()).To(BeEmpty())
	})
})
