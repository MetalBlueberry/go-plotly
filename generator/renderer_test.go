package generator_test

import (
	"bytes"
	"go/format"

	_ "embed"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/MetalBlueberry/go-plotly/generator"
	"github.com/MetalBlueberry/go-plotly/generator/mocks"
)

//go:embed schema.json
var schema []byte

var _ = Describe("Renderer", func() {

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

type NopWriterCloser struct {
	*bytes.Buffer
}

func (_ NopWriterCloser) Close() error {
	return nil
}
