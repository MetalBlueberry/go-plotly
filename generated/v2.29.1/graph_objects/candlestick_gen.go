package grob

// Code generated by go-plotly/generator. DO NOT EDIT.

import (
	"encoding/json"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

var TraceTypeCandlestick types.TraceType = "candlestick"

func (t *Candlestick) GetType() types.TraceType {
	return TraceTypeCandlestick
}

func (t *Candlestick) MarshalJSON() ([]byte, error) {
	// Define the custom JSON structure including the "type" field
	type Alias Candlestick
	return json.Marshal(&struct {
		Type types.TraceType `json:"type"`
		*Alias
	}{
		Type:  t.GetType(), // Add your desired default value here
		Alias: (*Alias)(t), // Embed the original struct fields
	})
}

// Candlestick The candlestick is a style of financial chart describing open, high, low and close for a given `x` coordinate (most likely time). The boxes represent the spread between the `open` and `close` values and the lines represent the spread between the `low` and `high` values Sample points where the close value is higher (lower) then the open value are called increasing (decreasing). By default, increasing candles are drawn in green whereas decreasing are drawn in red.
type Candlestick struct {

	// Close
	// arrayOK: false
	// type: data_array
	// Sets the close values.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.candlestick.attributes.close
	Close *types.DataArrayType `json:"close,omitempty"`

	// Closesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `close`.
	// .schema.traces.candlestick.attributes.closesrc
	Closesrc types.StringType `json:"closesrc,omitempty"`

	// Customdata
	// arrayOK: false
	// type: data_array
	// Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.candlestick.attributes.customdata
	Customdata *types.DataArrayType `json:"customdata,omitempty"`

	// Customdatasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `customdata`.
	// .schema.traces.candlestick.attributes.customdatasrc
	Customdatasrc types.StringType `json:"customdatasrc,omitempty"`

	// Decreasing
	// arrayOK: false
	// role: Object
	// .schema.traces.candlestick.attributes.decreasing
	Decreasing *CandlestickDecreasing `json:"decreasing,omitempty"`

	// High
	// arrayOK: false
	// type: data_array
	// Sets the high values.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.candlestick.attributes.high
	High *types.DataArrayType `json:"high,omitempty"`

	// Highsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `high`.
	// .schema.traces.candlestick.attributes.highsrc
	Highsrc types.StringType `json:"highsrc,omitempty"`

	// Hoverinfo
	// arrayOK: true
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	// .schema.traces.candlestick.attributes.hoverinfo
	Hoverinfo *types.ArrayOK[*CandlestickHoverinfo] `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hoverinfo`.
	// .schema.traces.candlestick.attributes.hoverinfosrc
	Hoverinfosrc types.StringType `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// arrayOK: false
	// role: Object
	// .schema.traces.candlestick.attributes.hoverlabel
	Hoverlabel *CandlestickHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertext
	// arrayOK: true
	// type: string
	// Same as `text`.
	// .schema.traces.candlestick.attributes.hovertext
	Hovertext *types.ArrayOK[*types.StringType] `json:"hovertext,omitempty"`

	// Hovertextsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hovertext`.
	// .schema.traces.candlestick.attributes.hovertextsrc
	Hovertextsrc types.StringType `json:"hovertextsrc,omitempty"`

	// Ids
	// arrayOK: false
	// type: data_array
	// Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.candlestick.attributes.ids
	Ids *types.DataArrayType `json:"ids,omitempty"`

	// Idssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `ids`.
	// .schema.traces.candlestick.attributes.idssrc
	Idssrc types.StringType `json:"idssrc,omitempty"`

	// Increasing
	// arrayOK: false
	// role: Object
	// .schema.traces.candlestick.attributes.increasing
	Increasing *CandlestickIncreasing `json:"increasing,omitempty"`

	// Legend
	// arrayOK: false
	// type: subplotid
	// Sets the reference to a legend to show this trace in. References to these legends are *legend*, *legend2*, *legend3*, etc. Settings for these legends are set in the layout, under `layout.legend`, `layout.legend2`, etc.
	// .schema.traces.candlestick.attributes.legend
	Legend types.StringType `json:"legend,omitempty"`

	// Legendgroup
	// arrayOK: false
	// type: string
	// Sets the legend group for this trace. Traces and shapes part of the same legend group hide/show at the same time when toggling legend items.
	// .schema.traces.candlestick.attributes.legendgroup
	Legendgroup types.StringType `json:"legendgroup,omitempty"`

	// Legendgrouptitle
	// arrayOK: false
	// role: Object
	// .schema.traces.candlestick.attributes.legendgrouptitle
	Legendgrouptitle *CandlestickLegendgrouptitle `json:"legendgrouptitle,omitempty"`

	// Legendrank
	// arrayOK: false
	// type: number
	// Sets the legend rank for this trace. Items and groups with smaller ranks are presented on top/left side while with *reversed* `legend.traceorder` they are on bottom/right side. The default legendrank is 1000, so that you can use ranks less than 1000 to place certain items before all unranked items, and ranks greater than 1000 to go after all unranked items. When having unranked or equal rank items shapes would be displayed after traces i.e. according to their order in data and layout.
	// .schema.traces.candlestick.attributes.legendrank
	Legendrank types.NumberType `json:"legendrank,omitempty"`

	// Legendwidth
	// arrayOK: false
	// type: number
	// Sets the width (in px or fraction) of the legend for this trace.
	// .schema.traces.candlestick.attributes.legendwidth
	Legendwidth types.NumberType `json:"legendwidth,omitempty"`

	// Line
	// arrayOK: false
	// role: Object
	// .schema.traces.candlestick.attributes.line
	Line *CandlestickLine `json:"line,omitempty"`

	// Low
	// arrayOK: false
	// type: data_array
	// Sets the low values.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.candlestick.attributes.low
	Low *types.DataArrayType `json:"low,omitempty"`

	// Lowsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `low`.
	// .schema.traces.candlestick.attributes.lowsrc
	Lowsrc types.StringType `json:"lowsrc,omitempty"`

	// Meta
	// arrayOK: true
	// type: any
	// Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	// .schema.traces.candlestick.attributes.meta
	Meta *types.ArrayOK[*interface{}] `json:"meta,omitempty"`

	// Metasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `meta`.
	// .schema.traces.candlestick.attributes.metasrc
	Metasrc types.StringType `json:"metasrc,omitempty"`

	// Name
	// arrayOK: false
	// type: string
	// Sets the trace name. The trace name appears as the legend item and on hover.
	// .schema.traces.candlestick.attributes.name
	Name types.StringType `json:"name,omitempty"`

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the opacity of the trace.
	// .schema.traces.candlestick.attributes.opacity
	Opacity types.NumberType `json:"opacity,omitempty"`

	// Open
	// arrayOK: false
	// type: data_array
	// Sets the open values.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.candlestick.attributes.open
	Open *types.DataArrayType `json:"open,omitempty"`

	// Opensrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `open`.
	// .schema.traces.candlestick.attributes.opensrc
	Opensrc types.StringType `json:"opensrc,omitempty"`

	// Selectedpoints
	// arrayOK: false
	// type: any
	// Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	// .schema.traces.candlestick.attributes.selectedpoints
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend
	// arrayOK: false
	// type: boolean
	// Determines whether or not an item corresponding to this trace is shown in the legend.
	// .schema.traces.candlestick.attributes.showlegend
	Showlegend types.BoolType `json:"showlegend,omitempty"`

	// Stream
	// arrayOK: false
	// role: Object
	// .schema.traces.candlestick.attributes.stream
	Stream *CandlestickStream `json:"stream,omitempty"`

	// Text
	// arrayOK: true
	// type: string
	// Sets hover text elements associated with each sample point. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to this trace's sample points.
	// .schema.traces.candlestick.attributes.text
	Text *types.ArrayOK[*types.StringType] `json:"text,omitempty"`

	// Textsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `text`.
	// .schema.traces.candlestick.attributes.textsrc
	Textsrc types.StringType `json:"textsrc,omitempty"`

	// Transforms
	// role: Object
	// items: CandlestickTransform
	// .schema.traces.candlestick.attributes.transforms
	Transforms []CandlestickTransform `json:"transforms,omitempty"`

	// Uid
	// arrayOK: false
	// type: string
	// Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	// .schema.traces.candlestick.attributes.uid
	Uid types.StringType `json:"uid,omitempty"`

	// Uirevision
	// arrayOK: false
	// type: any
	// Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	// .schema.traces.candlestick.attributes.uirevision
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible
	// arrayOK: false
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	// .schema.traces.candlestick.attributes.visible
	Visible CandlestickVisible `json:"visible,omitempty"`

	// Whiskerwidth
	// arrayOK: false
	// type: number
	// Sets the width of the whiskers relative to the box' width. For example, with 1, the whiskers are as wide as the box(es).
	// .schema.traces.candlestick.attributes.whiskerwidth
	Whiskerwidth types.NumberType `json:"whiskerwidth,omitempty"`

	// X
	// arrayOK: false
	// type: data_array
	// Sets the x coordinates. If absent, linear coordinate will be generated.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.candlestick.attributes.x
	X *types.DataArrayType `json:"x,omitempty"`

	// Xaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	// .schema.traces.candlestick.attributes.xaxis
	Xaxis types.StringType `json:"xaxis,omitempty"`

	// Xcalendar
	// arrayOK: false
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `x` date data.
	// .schema.traces.candlestick.attributes.xcalendar
	Xcalendar CandlestickXcalendar `json:"xcalendar,omitempty"`

	// Xhoverformat
	// arrayOK: false
	// type: string
	// Sets the hover text formatting rulefor `x`  using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-format/tree/v1.4.5#d3-format. And for dates see: https://github.com/d3/d3-time-format/tree/v2.2.3#locale_format. We add two items to d3's date formatter: *%h* for half of the year as a decimal number as well as *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46*By default the values are formatted using `xaxis.hoverformat`.
	// .schema.traces.candlestick.attributes.xhoverformat
	Xhoverformat types.StringType `json:"xhoverformat,omitempty"`

	// Xperiod
	// arrayOK: false
	// type: any
	// Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the x axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	// .schema.traces.candlestick.attributes.xperiod
	Xperiod interface{} `json:"xperiod,omitempty"`

	// Xperiod0
	// arrayOK: false
	// type: any
	// Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the x0 axis. When `x0period` is round number of weeks, the `x0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	// .schema.traces.candlestick.attributes.xperiod0
	Xperiod0 interface{} `json:"xperiod0,omitempty"`

	// Xperiodalignment
	// arrayOK: false
	// default: middle
	// type: enumerated
	// Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
	// .schema.traces.candlestick.attributes.xperiodalignment
	Xperiodalignment CandlestickXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `x`.
	// .schema.traces.candlestick.attributes.xsrc
	Xsrc types.StringType `json:"xsrc,omitempty"`

	// Yaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	// .schema.traces.candlestick.attributes.yaxis
	Yaxis types.StringType `json:"yaxis,omitempty"`

	// Yhoverformat
	// arrayOK: false
	// type: string
	// Sets the hover text formatting rulefor `y`  using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-format/tree/v1.4.5#d3-format. And for dates see: https://github.com/d3/d3-time-format/tree/v2.2.3#locale_format. We add two items to d3's date formatter: *%h* for half of the year as a decimal number as well as *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46*By default the values are formatted using `yaxis.hoverformat`.
	// .schema.traces.candlestick.attributes.yhoverformat
	Yhoverformat types.StringType `json:"yhoverformat,omitempty"`
}

// CandlestickDecreasingLine
type CandlestickDecreasingLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the color of line bounding the box(es).
	// .schema.traces.candlestick.attributes.decreasing.line.color
	Color types.Color `json:"color,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the width (in px) of line bounding the box(es).
	// .schema.traces.candlestick.attributes.decreasing.line.width
	Width types.NumberType `json:"width,omitempty"`
}

// CandlestickDecreasing
type CandlestickDecreasing struct {

	// Fillcolor
	// arrayOK: false
	// type: color
	// Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	// .schema.traces.candlestick.attributes.decreasing.fillcolor
	Fillcolor types.Color `json:"fillcolor,omitempty"`

	// Line
	// arrayOK: false
	// role: Object
	// .schema.traces.candlestick.attributes.decreasing.line
	Line *CandlestickDecreasingLine `json:"line,omitempty"`
}

// CandlestickHoverlabelFont Sets the font used in hover labels.
type CandlestickHoverlabelFont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	// .schema.traces.candlestick.attributes.hoverlabel.font.color
	Color *types.ArrayOK[*types.Color] `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	// .schema.traces.candlestick.attributes.hoverlabel.font.colorsrc
	Colorsrc types.StringType `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	// .schema.traces.candlestick.attributes.hoverlabel.font.family
	Family *types.ArrayOK[*types.StringType] `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `family`.
	// .schema.traces.candlestick.attributes.hoverlabel.font.familysrc
	Familysrc types.StringType `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	// .schema.traces.candlestick.attributes.hoverlabel.font.size
	Size *types.ArrayOK[*types.NumberType] `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `size`.
	// .schema.traces.candlestick.attributes.hoverlabel.font.sizesrc
	Sizesrc types.StringType `json:"sizesrc,omitempty"`
}

// CandlestickHoverlabel
type CandlestickHoverlabel struct {

	// Align
	// arrayOK: true
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	// .schema.traces.candlestick.attributes.hoverlabel.align
	Align *types.ArrayOK[*CandlestickHoverlabelAlign] `json:"align,omitempty"`

	// Alignsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `align`.
	// .schema.traces.candlestick.attributes.hoverlabel.alignsrc
	Alignsrc types.StringType `json:"alignsrc,omitempty"`

	// Bgcolor
	// arrayOK: true
	// type: color
	// Sets the background color of the hover labels for this trace
	// .schema.traces.candlestick.attributes.hoverlabel.bgcolor
	Bgcolor *types.ArrayOK[*types.Color] `json:"bgcolor,omitempty"`

	// Bgcolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `bgcolor`.
	// .schema.traces.candlestick.attributes.hoverlabel.bgcolorsrc
	Bgcolorsrc types.StringType `json:"bgcolorsrc,omitempty"`

	// Bordercolor
	// arrayOK: true
	// type: color
	// Sets the border color of the hover labels for this trace.
	// .schema.traces.candlestick.attributes.hoverlabel.bordercolor
	Bordercolor *types.ArrayOK[*types.Color] `json:"bordercolor,omitempty"`

	// Bordercolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `bordercolor`.
	// .schema.traces.candlestick.attributes.hoverlabel.bordercolorsrc
	Bordercolorsrc types.StringType `json:"bordercolorsrc,omitempty"`

	// Font
	// arrayOK: false
	// role: Object
	// .schema.traces.candlestick.attributes.hoverlabel.font
	Font *CandlestickHoverlabelFont `json:"font,omitempty"`

	// Namelength
	// arrayOK: true
	// type: integer
	// Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis.
	// .schema.traces.candlestick.attributes.hoverlabel.namelength
	Namelength *types.ArrayOK[*types.IntegerType] `json:"namelength,omitempty"`

	// Namelengthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `namelength`.
	// .schema.traces.candlestick.attributes.hoverlabel.namelengthsrc
	Namelengthsrc types.StringType `json:"namelengthsrc,omitempty"`

	// Split
	// arrayOK: false
	// type: boolean
	// Show hover information (open, close, high, low) in separate labels.
	// .schema.traces.candlestick.attributes.hoverlabel.split
	Split types.BoolType `json:"split,omitempty"`
}

// CandlestickIncreasingLine
type CandlestickIncreasingLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the color of line bounding the box(es).
	// .schema.traces.candlestick.attributes.increasing.line.color
	Color types.Color `json:"color,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the width (in px) of line bounding the box(es).
	// .schema.traces.candlestick.attributes.increasing.line.width
	Width types.NumberType `json:"width,omitempty"`
}

// CandlestickIncreasing
type CandlestickIncreasing struct {

	// Fillcolor
	// arrayOK: false
	// type: color
	// Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	// .schema.traces.candlestick.attributes.increasing.fillcolor
	Fillcolor types.Color `json:"fillcolor,omitempty"`

	// Line
	// arrayOK: false
	// role: Object
	// .schema.traces.candlestick.attributes.increasing.line
	Line *CandlestickIncreasingLine `json:"line,omitempty"`
}

// CandlestickLegendgrouptitleFont Sets this legend group's title font.
type CandlestickLegendgrouptitleFont struct {

	// Color
	// arrayOK: false
	// type: color
	//
	// .schema.traces.candlestick.attributes.legendgrouptitle.font.color
	Color types.Color `json:"color,omitempty"`

	// Family
	// arrayOK: false
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	// .schema.traces.candlestick.attributes.legendgrouptitle.font.family
	Family types.StringType `json:"family,omitempty"`

	// Size
	// arrayOK: false
	// type: number
	//
	// .schema.traces.candlestick.attributes.legendgrouptitle.font.size
	Size types.NumberType `json:"size,omitempty"`
}

// CandlestickLegendgrouptitle
type CandlestickLegendgrouptitle struct {

	// Font
	// arrayOK: false
	// role: Object
	// .schema.traces.candlestick.attributes.legendgrouptitle.font
	Font *CandlestickLegendgrouptitleFont `json:"font,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the legend group.
	// .schema.traces.candlestick.attributes.legendgrouptitle.text
	Text types.StringType `json:"text,omitempty"`
}

// CandlestickLine
type CandlestickLine struct {

	// Width
	// arrayOK: false
	// type: number
	// Sets the width (in px) of line bounding the box(es). Note that this style setting can also be set per direction via `increasing.line.width` and `decreasing.line.width`.
	// .schema.traces.candlestick.attributes.line.width
	Width types.NumberType `json:"width,omitempty"`
}

// CandlestickStream
type CandlestickStream struct {

	// Maxpoints
	// arrayOK: false
	// type: number
	// Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot.
	// .schema.traces.candlestick.attributes.stream.maxpoints
	Maxpoints types.NumberType `json:"maxpoints,omitempty"`

	// Token
	// arrayOK: false
	// type: string
	// The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details.
	// .schema.traces.candlestick.attributes.stream.token
	Token types.StringType `json:"token,omitempty"`
}

// CandlestickTransform WARNING: All transforms are deprecated and may be removed from the API in next major version. An array of operations that manipulate the trace data, for example filtering or sorting the data arrays.
type CandlestickTransform struct {
}

// CandlestickHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
// .schema.traces.candlestick.attributes.hoverlabel.align
type CandlestickHoverlabelAlign string

const (
	CandlestickHoverlabelAlignLeft  CandlestickHoverlabelAlign = "left"
	CandlestickHoverlabelAlignRight CandlestickHoverlabelAlign = "right"
	CandlestickHoverlabelAlignAuto  CandlestickHoverlabelAlign = "auto"
)

// CandlestickVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
// .schema.traces.candlestick.attributes.visible
type CandlestickVisible interface{}

var (
	CandlestickVisibleTrue       CandlestickVisible = true
	CandlestickVisibleFalse      CandlestickVisible = false
	CandlestickVisibleLegendonly CandlestickVisible = "legendonly"
)

// CandlestickXcalendar Sets the calendar system to use with `x` date data.
// .schema.traces.candlestick.attributes.xcalendar
type CandlestickXcalendar string

const (
	CandlestickXcalendarChinese    CandlestickXcalendar = "chinese"
	CandlestickXcalendarCoptic     CandlestickXcalendar = "coptic"
	CandlestickXcalendarDiscworld  CandlestickXcalendar = "discworld"
	CandlestickXcalendarEthiopian  CandlestickXcalendar = "ethiopian"
	CandlestickXcalendarGregorian  CandlestickXcalendar = "gregorian"
	CandlestickXcalendarHebrew     CandlestickXcalendar = "hebrew"
	CandlestickXcalendarIslamic    CandlestickXcalendar = "islamic"
	CandlestickXcalendarJalali     CandlestickXcalendar = "jalali"
	CandlestickXcalendarJulian     CandlestickXcalendar = "julian"
	CandlestickXcalendarMayan      CandlestickXcalendar = "mayan"
	CandlestickXcalendarNanakshahi CandlestickXcalendar = "nanakshahi"
	CandlestickXcalendarNepali     CandlestickXcalendar = "nepali"
	CandlestickXcalendarPersian    CandlestickXcalendar = "persian"
	CandlestickXcalendarTaiwan     CandlestickXcalendar = "taiwan"
	CandlestickXcalendarThai       CandlestickXcalendar = "thai"
	CandlestickXcalendarUmmalqura  CandlestickXcalendar = "ummalqura"
)

// CandlestickXperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
// .schema.traces.candlestick.attributes.xperiodalignment
type CandlestickXperiodalignment string

const (
	CandlestickXperiodalignmentStart  CandlestickXperiodalignment = "start"
	CandlestickXperiodalignmentMiddle CandlestickXperiodalignment = "middle"
	CandlestickXperiodalignmentEnd    CandlestickXperiodalignment = "end"
)

// CandlestickHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
// .schema.traces.candlestick.attributes.hoverinfo
type CandlestickHoverinfo string

const (
	// Flags
	CandlestickHoverinfoX    CandlestickHoverinfo = "x"
	CandlestickHoverinfoY    CandlestickHoverinfo = "y"
	CandlestickHoverinfoZ    CandlestickHoverinfo = "z"
	CandlestickHoverinfoText CandlestickHoverinfo = "text"
	CandlestickHoverinfoName CandlestickHoverinfo = "name"

	// Extra
	CandlestickHoverinfoAll  CandlestickHoverinfo = "all"
	CandlestickHoverinfoNone CandlestickHoverinfo = "none"
	CandlestickHoverinfoSkip CandlestickHoverinfo = "skip"
)
