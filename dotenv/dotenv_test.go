package dotenv_test

import (
	. "build-dotenv/dotenv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The Dotenv type", func() {
	It("initialises an empty map", func() {
		env := New()

		Expect(env.Join()).To(BeEmpty())
	})

	It("initialises a simple map", func() {
		env := New()
		env.AddFromString("lol", `A=1
B=2
C=3`)

		Expect(env.Join()).To(Equal(map[string]string{
			"A": "1",
			"B": "2",
			"C": "3",
		}))
	})
})
