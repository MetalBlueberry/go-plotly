package grob

// Code generated by go-plotly/generator. DO NOT EDIT.

import (
	"encoding/json"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

var TraceTypeTable types.TraceType = "table"

func (t *Table) GetType() types.TraceType {
	return TraceTypeTable
}

func (t *Table) MarshalJSON() ([]byte, error) {
	// Define the custom JSON structure including the "type" field
	type Alias Table
	return json.Marshal(&struct {
		Type types.TraceType `json:"type"`
		*Alias
	}{
		Type:  t.GetType(), // Add your desired default value here
		Alias: (*Alias)(t), // Embed the original struct fields
	})
}

// Table Table view for detailed data viewing. The data are arranged in a grid of rows and columns. Most styling can be specified for columns, rows or individual cells. Table is using a column-major order, ie. the grid is represented as a vector of column vectors.
type Table struct {

	// Cells
	// arrayOK: false
	// role: Object
	Cells *TableCells `json:"cells,omitempty"`

	// Columnorder
	// arrayOK: false
	// type: data_array
	// Specifies the rendered order of the data columns; for example, a value `2` at position `0` means that column index `0` in the data will be rendered as the third column, as columns have an index base of zero.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	Columnorder *types.DataArrayType `json:"columnorder,omitempty"`

	// Columnordersrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `columnorder`.
	Columnordersrc types.StringType `json:"columnordersrc,omitempty"`

	// Columnwidth
	// arrayOK: true
	// type: number
	// The width of columns expressed as a ratio. Columns fill the available width in proportion of their specified column widths.
	Columnwidth *types.ArrayOK[*types.NumberType] `json:"columnwidth,omitempty"`

	// Columnwidthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `columnwidth`.
	Columnwidthsrc types.StringType `json:"columnwidthsrc,omitempty"`

	// Customdata
	// arrayOK: false
	// type: data_array
	// Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	Customdata *types.DataArrayType `json:"customdata,omitempty"`

	// Customdatasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `customdata`.
	Customdatasrc types.StringType `json:"customdatasrc,omitempty"`

	// Domain
	// arrayOK: false
	// role: Object
	Domain *TableDomain `json:"domain,omitempty"`

	// Header
	// arrayOK: false
	// role: Object
	Header *TableHeader `json:"header,omitempty"`

	// Hoverinfo
	// arrayOK: true
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo *types.ArrayOK[*TableHoverinfo] `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hoverinfo`.
	Hoverinfosrc types.StringType `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// arrayOK: false
	// role: Object
	Hoverlabel *TableHoverlabel `json:"hoverlabel,omitempty"`

	// Ids
	// arrayOK: false
	// type: data_array
	// Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	Ids *types.DataArrayType `json:"ids,omitempty"`

	// Idssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `ids`.
	Idssrc types.StringType `json:"idssrc,omitempty"`

	// Legendgrouptitle
	// arrayOK: false
	// role: Object
	Legendgrouptitle *TableLegendgrouptitle `json:"legendgrouptitle,omitempty"`

	// Legendrank
	// arrayOK: false
	// type: number
	// Sets the legend rank for this trace. Items and groups with smaller ranks are presented on top/left side while with `*reversed* `legend.traceorder` they are on bottom/right side. The default legendrank is 1000, so that you can use ranks less than 1000 to place certain items before all unranked items, and ranks greater than 1000 to go after all unranked items.
	Legendrank types.NumberType `json:"legendrank,omitempty"`

	// Legendwidth
	// arrayOK: false
	// type: number
	// Sets the width (in px or fraction) of the legend for this trace.
	Legendwidth types.NumberType `json:"legendwidth,omitempty"`

	// Meta
	// arrayOK: true
	// type: any
	// Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta *types.ArrayOK[*interface{}] `json:"meta,omitempty"`

	// Metasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `meta`.
	Metasrc types.StringType `json:"metasrc,omitempty"`

	// Name
	// arrayOK: false
	// type: string
	// Sets the trace name. The trace name appear as the legend item and on hover.
	Name types.StringType `json:"name,omitempty"`

	// Stream
	// arrayOK: false
	// role: Object
	Stream *TableStream `json:"stream,omitempty"`

	// Uid
	// arrayOK: false
	// type: string
	// Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid types.StringType `json:"uid,omitempty"`

	// Uirevision
	// arrayOK: false
	// type: any
	// Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible
	// arrayOK: false
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible TableVisible `json:"visible,omitempty"`
}

// TableCellsFill
type TableCellsFill struct {

	// Color
	// arrayOK: true
	// type: color
	// Sets the cell fill color. It accepts either a specific color or an array of colors or a 2D array of colors.
	Color *types.ArrayOK[*types.Color] `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	Colorsrc types.StringType `json:"colorsrc,omitempty"`
}

// TableCellsFont
type TableCellsFont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color *types.ArrayOK[*types.Color] `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	Colorsrc types.StringType `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family *types.ArrayOK[*types.StringType] `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `family`.
	Familysrc types.StringType `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size *types.ArrayOK[*types.NumberType] `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `size`.
	Sizesrc types.StringType `json:"sizesrc,omitempty"`
}

// TableCellsLine
type TableCellsLine struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color *types.ArrayOK[*types.Color] `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	Colorsrc types.StringType `json:"colorsrc,omitempty"`

	// Width
	// arrayOK: true
	// type: number
	//
	Width *types.ArrayOK[*types.NumberType] `json:"width,omitempty"`

	// Widthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `width`.
	Widthsrc types.StringType `json:"widthsrc,omitempty"`
}

// TableCells
type TableCells struct {

	// Align
	// arrayOK: true
	// default: center
	// type: enumerated
	// Sets the horizontal alignment of the `text` within the box. Has an effect only if `text` spans two or more lines (i.e. `text` contains one or more <br> HTML tags) or if an explicit width is set to override the text width.
	Align *types.ArrayOK[*TableCellsAlign] `json:"align,omitempty"`

	// Alignsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `align`.
	Alignsrc types.StringType `json:"alignsrc,omitempty"`

	// Fill
	// arrayOK: false
	// role: Object
	Fill *TableCellsFill `json:"fill,omitempty"`

	// Font
	// arrayOK: false
	// role: Object
	Font *TableCellsFont `json:"font,omitempty"`

	// Format
	// arrayOK: false
	// type: data_array
	// Sets the cell value formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-format/tree/v1.4.5#d3-format.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	Format *types.DataArrayType `json:"format,omitempty"`

	// Formatsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `format`.
	Formatsrc types.StringType `json:"formatsrc,omitempty"`

	// Height
	// arrayOK: false
	// type: number
	// The height of cells.
	Height types.NumberType `json:"height,omitempty"`

	// Line
	// arrayOK: false
	// role: Object
	Line *TableCellsLine `json:"line,omitempty"`

	// Prefix
	// arrayOK: true
	// type: string
	// Prefix for cell values.
	Prefix *types.ArrayOK[*types.StringType] `json:"prefix,omitempty"`

	// Prefixsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `prefix`.
	Prefixsrc types.StringType `json:"prefixsrc,omitempty"`

	// Suffix
	// arrayOK: true
	// type: string
	// Suffix for cell values.
	Suffix *types.ArrayOK[*types.StringType] `json:"suffix,omitempty"`

	// Suffixsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `suffix`.
	Suffixsrc types.StringType `json:"suffixsrc,omitempty"`

	// Values
	// arrayOK: false
	// type: data_array
	// Cell values. `values[m][n]` represents the value of the `n`th point in column `m`, therefore the `values[m]` vector length for all columns must be the same (longer vectors will be truncated). Each value must be a finite number or a string.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	Values *types.DataArrayType `json:"values,omitempty"`

	// Valuessrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `values`.
	Valuessrc types.StringType `json:"valuessrc,omitempty"`
}

// TableDomain
type TableDomain struct {

	// Column
	// arrayOK: false
	// type: integer
	// If there is a layout grid, use the domain for this column in the grid for this table trace .
	Column types.IntegerType `json:"column,omitempty"`

	// Row
	// arrayOK: false
	// type: integer
	// If there is a layout grid, use the domain for this row in the grid for this table trace .
	Row types.IntegerType `json:"row,omitempty"`

	// X
	// arrayOK: false
	// type: info_array
	// Sets the horizontal domain of this table trace (in plot fraction).
	X interface{} `json:"x,omitempty"`

	// Y
	// arrayOK: false
	// type: info_array
	// Sets the vertical domain of this table trace (in plot fraction).
	Y interface{} `json:"y,omitempty"`
}

// TableHeaderFill
type TableHeaderFill struct {

	// Color
	// arrayOK: true
	// type: color
	// Sets the cell fill color. It accepts either a specific color or an array of colors or a 2D array of colors.
	Color *types.ArrayOK[*types.Color] `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	Colorsrc types.StringType `json:"colorsrc,omitempty"`
}

// TableHeaderFont
type TableHeaderFont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color *types.ArrayOK[*types.Color] `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	Colorsrc types.StringType `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family *types.ArrayOK[*types.StringType] `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `family`.
	Familysrc types.StringType `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size *types.ArrayOK[*types.NumberType] `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `size`.
	Sizesrc types.StringType `json:"sizesrc,omitempty"`
}

// TableHeaderLine
type TableHeaderLine struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color *types.ArrayOK[*types.Color] `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	Colorsrc types.StringType `json:"colorsrc,omitempty"`

	// Width
	// arrayOK: true
	// type: number
	//
	Width *types.ArrayOK[*types.NumberType] `json:"width,omitempty"`

	// Widthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `width`.
	Widthsrc types.StringType `json:"widthsrc,omitempty"`
}

// TableHeader
type TableHeader struct {

	// Align
	// arrayOK: true
	// default: center
	// type: enumerated
	// Sets the horizontal alignment of the `text` within the box. Has an effect only if `text` spans two or more lines (i.e. `text` contains one or more <br> HTML tags) or if an explicit width is set to override the text width.
	Align *types.ArrayOK[*TableHeaderAlign] `json:"align,omitempty"`

	// Alignsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `align`.
	Alignsrc types.StringType `json:"alignsrc,omitempty"`

	// Fill
	// arrayOK: false
	// role: Object
	Fill *TableHeaderFill `json:"fill,omitempty"`

	// Font
	// arrayOK: false
	// role: Object
	Font *TableHeaderFont `json:"font,omitempty"`

	// Format
	// arrayOK: false
	// type: data_array
	// Sets the cell value formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-format/tree/v1.4.5#d3-format.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	Format *types.DataArrayType `json:"format,omitempty"`

	// Formatsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `format`.
	Formatsrc types.StringType `json:"formatsrc,omitempty"`

	// Height
	// arrayOK: false
	// type: number
	// The height of cells.
	Height types.NumberType `json:"height,omitempty"`

	// Line
	// arrayOK: false
	// role: Object
	Line *TableHeaderLine `json:"line,omitempty"`

	// Prefix
	// arrayOK: true
	// type: string
	// Prefix for cell values.
	Prefix *types.ArrayOK[*types.StringType] `json:"prefix,omitempty"`

	// Prefixsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `prefix`.
	Prefixsrc types.StringType `json:"prefixsrc,omitempty"`

	// Suffix
	// arrayOK: true
	// type: string
	// Suffix for cell values.
	Suffix *types.ArrayOK[*types.StringType] `json:"suffix,omitempty"`

	// Suffixsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `suffix`.
	Suffixsrc types.StringType `json:"suffixsrc,omitempty"`

	// Values
	// arrayOK: false
	// type: data_array
	// Header cell values. `values[m][n]` represents the value of the `n`th point in column `m`, therefore the `values[m]` vector length for all columns must be the same (longer vectors will be truncated). Each value must be a finite number or a string.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	Values *types.DataArrayType `json:"values,omitempty"`

	// Valuessrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `values`.
	Valuessrc types.StringType `json:"valuessrc,omitempty"`
}

// TableHoverlabelFont Sets the font used in hover labels.
type TableHoverlabelFont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color *types.ArrayOK[*types.Color] `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	Colorsrc types.StringType `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family *types.ArrayOK[*types.StringType] `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `family`.
	Familysrc types.StringType `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size *types.ArrayOK[*types.NumberType] `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `size`.
	Sizesrc types.StringType `json:"sizesrc,omitempty"`
}

// TableHoverlabel
type TableHoverlabel struct {

	// Align
	// arrayOK: true
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align *types.ArrayOK[*TableHoverlabelAlign] `json:"align,omitempty"`

	// Alignsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `align`.
	Alignsrc types.StringType `json:"alignsrc,omitempty"`

	// Bgcolor
	// arrayOK: true
	// type: color
	// Sets the background color of the hover labels for this trace
	Bgcolor *types.ArrayOK[*types.Color] `json:"bgcolor,omitempty"`

	// Bgcolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `bgcolor`.
	Bgcolorsrc types.StringType `json:"bgcolorsrc,omitempty"`

	// Bordercolor
	// arrayOK: true
	// type: color
	// Sets the border color of the hover labels for this trace.
	Bordercolor *types.ArrayOK[*types.Color] `json:"bordercolor,omitempty"`

	// Bordercolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `bordercolor`.
	Bordercolorsrc types.StringType `json:"bordercolorsrc,omitempty"`

	// Font
	// arrayOK: false
	// role: Object
	Font *TableHoverlabelFont `json:"font,omitempty"`

	// Namelength
	// arrayOK: true
	// type: integer
	// Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis.
	Namelength *types.ArrayOK[*types.IntegerType] `json:"namelength,omitempty"`

	// Namelengthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `namelength`.
	Namelengthsrc types.StringType `json:"namelengthsrc,omitempty"`
}

// TableLegendgrouptitleFont Sets this legend group's title font.
type TableLegendgrouptitleFont struct {

	// Color
	// arrayOK: false
	// type: color
	//
	Color types.Color `json:"color,omitempty"`

	// Family
	// arrayOK: false
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family types.StringType `json:"family,omitempty"`

	// Size
	// arrayOK: false
	// type: number
	//
	Size types.NumberType `json:"size,omitempty"`
}

// TableLegendgrouptitle
type TableLegendgrouptitle struct {

	// Font
	// arrayOK: false
	// role: Object
	Font *TableLegendgrouptitleFont `json:"font,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the legend group.
	Text types.StringType `json:"text,omitempty"`
}

// TableStream
type TableStream struct {

	// Maxpoints
	// arrayOK: false
	// type: number
	// Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot.
	Maxpoints types.NumberType `json:"maxpoints,omitempty"`

	// Token
	// arrayOK: false
	// type: string
	// The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details.
	Token types.StringType `json:"token,omitempty"`
}

// TableCellsAlign Sets the horizontal alignment of the `text` within the box. Has an effect only if `text` spans two or more lines (i.e. `text` contains one or more <br> HTML tags) or if an explicit width is set to override the text width.
type TableCellsAlign string

const (
	TableCellsAlignLeft   TableCellsAlign = "left"
	TableCellsAlignCenter TableCellsAlign = "center"
	TableCellsAlignRight  TableCellsAlign = "right"
)

// TableHeaderAlign Sets the horizontal alignment of the `text` within the box. Has an effect only if `text` spans two or more lines (i.e. `text` contains one or more <br> HTML tags) or if an explicit width is set to override the text width.
type TableHeaderAlign string

const (
	TableHeaderAlignLeft   TableHeaderAlign = "left"
	TableHeaderAlignCenter TableHeaderAlign = "center"
	TableHeaderAlignRight  TableHeaderAlign = "right"
)

// TableHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type TableHoverlabelAlign string

const (
	TableHoverlabelAlignLeft  TableHoverlabelAlign = "left"
	TableHoverlabelAlignRight TableHoverlabelAlign = "right"
	TableHoverlabelAlignAuto  TableHoverlabelAlign = "auto"
)

// TableVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type TableVisible interface{}

var (
	TableVisibleTrue       TableVisible = true
	TableVisibleFalse      TableVisible = false
	TableVisibleLegendonly TableVisible = "legendonly"
)

// TableHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type TableHoverinfo string

const (
	// Flags
	TableHoverinfoX    TableHoverinfo = "x"
	TableHoverinfoY    TableHoverinfo = "y"
	TableHoverinfoZ    TableHoverinfo = "z"
	TableHoverinfoText TableHoverinfo = "text"
	TableHoverinfoName TableHoverinfo = "name"

	// Extra
	TableHoverinfoAll  TableHoverinfo = "all"
	TableHoverinfoNone TableHoverinfo = "none"
	TableHoverinfoSkip TableHoverinfo = "skip"
)
