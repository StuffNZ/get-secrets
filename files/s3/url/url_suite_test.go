package url_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUrl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Url Suite")
}
