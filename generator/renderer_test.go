package generator_test

import (
	"bytes"
	"go/format"
	"io"
	"os"

	_ "embed"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/MetalBlueberry/go-plotly/generator"
	"github.com/MetalBlueberry/go-plotly/generator/mocks"
)

//go:embed schema.json
var schema []byte

type Creator struct{}

func (c Creator) Create(name string) (io.WriteCloser, error) {
	return os.Create(name)
}

var _ = Describe("Integration", func() {
	It("Should render traces", func() {
		root, err := generator.LoadSchema(bytes.NewReader(schema))
		Expect(err).To(BeNil())

		r, err := generator.NewRenderer(Creator{}, root)
		Expect(err).To(BeNil())

		err = r.WriteTraces("gen/")
		Expect(err).To(BeNil())

	})
	FIt("Should render layout", func() {
		root, err := generator.LoadSchema(bytes.NewReader(schema))
		Expect(err).To(BeNil())

		r, err := generator.NewRenderer(Creator{}, root)
		Expect(err).To(BeNil())

		err = r.WriteLayout("gen/")
		Expect(err).To(BeNil())
	})
	It("Should render config", func() {
		root, err := generator.LoadSchema(bytes.NewReader(schema))
		Expect(err).To(BeNil())

		r, err := generator.NewRenderer(Creator{}, root)
		Expect(err).To(BeNil())

		err = r.WriteConfig("gen/")
		Expect(err).To(BeNil())
	})
})

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

		mockCreator.EXPECT().Create(gomock.Eq("scatter.go")).Return(buf, nil).Times(1)

		root, err := generator.LoadSchema(bytes.NewReader(schema))
		Expect(err).To(BeNil())

		r, err := generator.NewRenderer(mockCreator, root)
		Expect(err).To(BeNil())

		err = r.CreateTrace("scatter")
		Expect(err).To(BeNil())

		formatted, err := format.Source(buf.Bytes())
		Expect(err).To(BeNil())

		Expect(string(formatted)).To(ContainSubstring(`package grob`))
		Expect(string(formatted)).To(ContainSubstring(`type Scatter struct`))

	})
})

type NopWriterCloser struct {
	*bytes.Buffer
}

func (_ NopWriterCloser) Close() error {
	return nil
}
