package dotenv_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDotenv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dotenv Suite")
}
