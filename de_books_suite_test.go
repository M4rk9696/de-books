package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDeBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DeBooks Suite")
}
