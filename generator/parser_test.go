package generator_test

import (
	"log"
	"os"

	"github.com/MetalBlueberry/go-plotly/generator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	It("Should parse traces", func() {
		schema, err := os.Open("schema.json")
		Expect(err).To(BeNil())
		defer schema.Close()

		root, err := generator.LoadSchema(schema)
		Expect(err).To(BeNil())

		log.Println(root)
	})
})
