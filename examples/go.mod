module github.com/metalblueberry/plotly/examples

go 1.21

toolchain go1.21.1

require (
	github.com/MetalBlueberry/go-plotly v0.0.0-20200503142240-1276ab260dcb
	github.com/go-gota/gota v0.10.1
	github.com/lucasb-eyer/go-colorful v1.2.0
)

require gonum.org/v1/gonum v0.7.0 // indirect

replace github.com/MetalBlueberry/go-plotly => ./../
