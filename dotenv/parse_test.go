package dotenv_test

import (
	. "build-dotenv/dotenv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The Dotenv type, parsing functions,", func() {
	It("initialises a simple map", func() {
		env := New()
		env.AddFromString("lol", `A=1
B=2
C=3`)

		Expect(env.Combine()).To(Equal(map[string]string{
			"A": "1",
			"B": "2",
			"C": "3",
		}))
	})

	It("combines multiple envs", func() {
		env := New()
		env.AddFromString("lol2", `A=1
B=2
C=3`)
		env.AddFromString("lol", `D=4
E=5`)

		Expect(env.Combine()).To(Equal(map[string]string{
			"A": "1",
			"B": "2",
			"C": "3",
			"D": "4",
			"E": "5",
		}))
	})

	It("combines multiple envs based on lexical-order of the path", func() {
		env := New()
		env.AddFromString("lol2", `A=11
B=2
C=3`)
		env.AddFromString("lol", `A=1
D=4`)

		Expect(env.Combine()).To(Equal(map[string]string{
			"A": "11",
			"B": "2",
			"C": "3",
			"D": "4",
		}))
	})
})
