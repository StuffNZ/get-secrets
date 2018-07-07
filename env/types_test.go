package env_test

import (
	"bitbucket.org/mexisme/get-secrets/dotenv"
	. "bitbucket.org/mexisme/get-secrets/env"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The env type", func() {
	Describe("when adding EnvMaps", func() {
		It("copies the single provided", func() {
			envMap := Map{
				"A": "1",
				"B": "2",
			}

			e := New().WithEnvMap(envMap)
			Expect(e.EnvMaps).To(Equal(Maps{
				{
					"A": "1",
					"B": "2",
				},
			}))
		})

		It("copies several individually provided", func() {
			envMap1 := Map{
				"A": "1",
				"B": "2",
			}
			envMap2 := Map{
				"C": "3",
				"D": "4",
			}
			envMap3 := Map{
				"E": "5",
				"F": "6",
			}

			e := New().WithEnvMap(envMap1).WithEnvMap(envMap2).WithEnvMap(envMap3)
			Expect(e.EnvMaps).To(Equal(Maps{
				{
					"A": "1",
					"B": "2",
				},
				{
					"C": "3",
					"D": "4",
				},
				{
					"E": "5",
					"F": "6",
				},
			}))
		})

		It("copies a provided group", func() {
			envMap := Maps{
				{
					"A": "1",
					"B": "2",
				},
				{
					"C": "3",
					"D": "4",
				},
				{
					"E": "5",
					"F": "6",
				},
			}

			e := New().WithEnvMaps(&envMap)
			Expect(e.EnvMaps).To(Equal(Maps{
				{
					"A": "1",
					"B": "2",
				},
				{
					"C": "3",
					"D": "4",
				},
				{
					"E": "5",
					"F": "6",
				},
			}))
		})
	})

	Describe("when adding from os.Environ", func() {
		It("copies the single provided", func() {
			environ := []string{
				"AA=11",
				"BB=22",
			}

			e := New().WithOsEnviron(environ)
			Expect(e.EnvMaps).To(Equal(Maps{
				{
					"AA": "11",
					"BB": "22",
				},
			}))
		})
	})

	Describe("when adding from DotEnv", func() {
		It("copies the single provided", func() {
			env := dotenv.New()
			env.AddFromString("lol.env", `A=1
B=2
C=3`)

			e := New().WithDotEnvs(env)
			Expect(e.EnvMaps).To(Equal(Maps{
				{
					"A": "1",
					"B": "2",
					"C": "3",
				},
			}))
		})
	})
})
