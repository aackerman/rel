package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
)

var _ = Describe("InfixOperationNode", func() {
	It("implements Predicator", func() {
		// compile time test
		var _ Predicator = InfixOperationNode{}
	})
})
