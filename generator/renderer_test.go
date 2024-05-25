package generator_test

import (
	"bytes"
	"go/format"

	_ "embed"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	// . "github.com/onsi/gomega/format"

	"github.com/MetalBlueberry/go-plotly/generator"
	"github.com/MetalBlueberry/go-plotly/generator/mocks"
)

// This test schema is used to make sure the generator works as expected
// It may be worth extending this to test with all schemas.
//
//go:embed test_schema.json
var schema []byte

var _ = Describe("Renderer", func() {

	Describe("A single trace", func() {
		var (
			ctrl        *gomock.Controller
			mockCreator *mocks.MockCreator
		)
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			mockCreator = mocks.NewMockCreator(ctrl)
		})
		AfterEach(func() {
			ctrl.Finish()
		})

		It("Should create package", func() {
			buf := NopWriterCloser{&bytes.Buffer{}}

			mockCreator.EXPECT().Create(gomock.Eq("scatter_gen.go")).Return(buf, nil).Times(1)

			root, err := generator.LoadSchema(bytes.NewReader(schema))
			Expect(err).To(BeNil())

			r, err := generator.NewRenderer(mockCreator, root)
			Expect(err).To(BeNil())

			err = r.CreateTrace(".", "scatter")
			Expect(err).To(BeNil())

			formatted, err := format.Source(buf.Bytes())
			Expect(err).To(BeNil())

			Expect(string(formatted)).To(ContainSubstring(`package grob`))
			// Type is defined
			Expect(string(formatted)).To(ContainSubstring(`type Scatter struct`))
			// Implements interface GetType()
			Expect(string(formatted)).To(ContainSubstring(`func (trace *Scatter) GetType() TraceType`))

		})
	})
	Describe("When writing", func() {
		var r *generator.Renderer

		BeforeEach(func() {
			root, err := generator.LoadSchema(bytes.NewReader(schema))
			Expect(err).To(BeNil())

			r, err = generator.NewRenderer(nil, root)
			Expect(err).To(BeNil())
		})

		Describe("The config", func() {
			It("Should be consistent", func() {

				original := &bytes.Buffer{}
				err := r.WriteConfig(original)
				Expect(err).To(BeNil())

				for i := 0; i < 10; i++ {
					attempt := &bytes.Buffer{}
					err = r.WriteConfig(attempt)
					Expect(err).To(BeNil())
					Expect(attempt).To(Equal(original))
				}

			})
		})
		Describe("The Layout", func() {
			// TruncatedDiff = false
			// MaxLength = 0
			It("Should be consistent", func() {

				original := &bytes.Buffer{}
				err := r.WriteLayout(original)
				Expect(err).To(BeNil())

				for i := 0; i < 10; i++ {
					attempt := &bytes.Buffer{}
					err = r.WriteLayout(attempt)
					Expect(err).To(BeNil())
					Expect(attempt.String()).To(Equal(original.String()))
				}

			})
		})
	})
})

type NopWriterCloser struct {
	*bytes.Buffer
}

func (_ NopWriterCloser) Close() error {
	return nil
}
