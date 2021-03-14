package generator_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/MetalBlueberry/go-plotly/generator"
)

var _ = Describe("Renderer", func() {
	FIt("Should render traces", func() {

		schema, err := os.Open("schema.json")
		Expect(err).To(BeNil())
		defer schema.Close()

		root, err := generator.LoadSchema(schema)
		Expect(err).To(BeNil())

		r, err := generator.NewRenderer("templates/*.tmpl", root)
		Expect(err).To(BeNil())

		err = r.WriteTraces(os.Stdout)
		Expect(err).To(BeNil())

	})
})
