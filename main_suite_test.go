package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBuildDotenv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BuildDotenv Suite")
}
