module github.com/metalblueberry/plotly/examples

go 1.22

require (
	github.com/MetalBlueberry/go-plotly v0.4.0
	github.com/go-gota/gota v0.12.0
	github.com/lucasb-eyer/go-colorful v1.2.0
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c
)

require (
	golang.org/x/exp v0.0.0-20240823005443-9b4947da3948 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.23.0 // indirect
	gonum.org/v1/gonum v0.15.0 // indirect
)

replace github.com/MetalBlueberry/go-plotly => ./../
