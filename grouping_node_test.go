package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
)

var _ = Describe("GroupingNode", func() {
	It("implements Predicator", func() {
		// compile time test
		var _ Predicator = GroupingNode{}
	})
})
