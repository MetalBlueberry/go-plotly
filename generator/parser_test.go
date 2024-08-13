package generator_test

import (
	"bytes"
	"log"

	"github.com/MetalBlueberry/go-plotly/generator"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	It("Should parse traces", func() {
		reader := bytes.NewReader(schema)

		root, err := generator.LoadSchema(reader)
		Expect(err).To(BeNil())

		log.Println(root)
	})
})
