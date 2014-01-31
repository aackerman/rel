package rel_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rel Suite")
}
