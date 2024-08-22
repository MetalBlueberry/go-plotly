package grob

// Code generated by go-plotly/generator. DO NOT EDIT.

import (
	"encoding/json"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

var TraceTypeOhlc types.TraceType = "ohlc"

func (t *Ohlc) GetType() types.TraceType {
	return TraceTypeOhlc
}

func (t *Ohlc) MarshalJSON() ([]byte, error) {
	// Define the custom JSON structure including the "type" field
	type Alias Ohlc
	return json.Marshal(&struct {
		Type types.TraceType `json:"type"`
		*Alias
	}{
		Type:  t.GetType(), // Add your desired default value here
		Alias: (*Alias)(t), // Embed the original struct fields
	})
}

// Ohlc The ohlc (short for Open-High-Low-Close) is a style of financial chart describing open, high, low and close for a given `x` coordinate (most likely time). The tip of the lines represent the `low` and `high` values and the horizontal segments represent the `open` and `close` values. Sample points where the close value is higher (lower) then the open value are called increasing (decreasing). By default, increasing items are drawn in green whereas decreasing are drawn in red.
type Ohlc struct {

	// Close
	// arrayOK: false
	// type: data_array
	// Sets the close values.
	Close interface{} `json:"close,omitempty"`

	// Closesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `close`.
	Closesrc types.String `json:"closesrc,omitempty"`

	// Customdata
	// arrayOK: false
	// type: data_array
	// Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `customdata`.
	Customdatasrc types.String `json:"customdatasrc,omitempty"`

	// Decreasing
	// arrayOK: false
	// role: Object
	Decreasing *OhlcDecreasing `json:"decreasing,omitempty"`

	// High
	// arrayOK: false
	// type: data_array
	// Sets the high values.
	High interface{} `json:"high,omitempty"`

	// Highsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `high`.
	Highsrc types.String `json:"highsrc,omitempty"`

	// Hoverinfo
	// arrayOK: true
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo *types.ArrayOK[*OhlcHoverinfo] `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hoverinfo`.
	Hoverinfosrc types.String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// arrayOK: false
	// role: Object
	Hoverlabel *OhlcHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertext
	// arrayOK: true
	// type: string
	// Same as `text`.
	Hovertext *types.ArrayOK[*types.String] `json:"hovertext,omitempty"`

	// Hovertextsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hovertext`.
	Hovertextsrc types.String `json:"hovertextsrc,omitempty"`

	// Ids
	// arrayOK: false
	// type: data_array
	// Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `ids`.
	Idssrc types.String `json:"idssrc,omitempty"`

	// Increasing
	// arrayOK: false
	// role: Object
	Increasing *OhlcIncreasing `json:"increasing,omitempty"`

	// Legend
	// arrayOK: false
	// type: subplotid
	// Sets the reference to a legend to show this trace in. References to these legends are *legend*, *legend2*, *legend3*, etc. Settings for these legends are set in the layout, under `layout.legend`, `layout.legend2`, etc.
	Legend types.String `json:"legend,omitempty"`

	// Legendgroup
	// arrayOK: false
	// type: string
	// Sets the legend group for this trace. Traces and shapes part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup types.String `json:"legendgroup,omitempty"`

	// Legendgrouptitle
	// arrayOK: false
	// role: Object
	Legendgrouptitle *OhlcLegendgrouptitle `json:"legendgrouptitle,omitempty"`

	// Legendrank
	// arrayOK: false
	// type: number
	// Sets the legend rank for this trace. Items and groups with smaller ranks are presented on top/left side while with *reversed* `legend.traceorder` they are on bottom/right side. The default legendrank is 1000, so that you can use ranks less than 1000 to place certain items before all unranked items, and ranks greater than 1000 to go after all unranked items. When having unranked or equal rank items shapes would be displayed after traces i.e. according to their order in data and layout.
	Legendrank float64 `json:"legendrank,omitempty"`

	// Legendwidth
	// arrayOK: false
	// type: number
	// Sets the width (in px or fraction) of the legend for this trace.
	Legendwidth float64 `json:"legendwidth,omitempty"`

	// Line
	// arrayOK: false
	// role: Object
	Line *OhlcLine `json:"line,omitempty"`

	// Low
	// arrayOK: false
	// type: data_array
	// Sets the low values.
	Low interface{} `json:"low,omitempty"`

	// Lowsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `low`.
	Lowsrc types.String `json:"lowsrc,omitempty"`

	// Meta
	// arrayOK: true
	// type: any
	// Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta *types.ArrayOK[*interface{}] `json:"meta,omitempty"`

	// Metasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `meta`.
	Metasrc types.String `json:"metasrc,omitempty"`

	// Name
	// arrayOK: false
	// type: string
	// Sets the trace name. The trace name appears as the legend item and on hover.
	Name types.String `json:"name,omitempty"`

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Open
	// arrayOK: false
	// type: data_array
	// Sets the open values.
	Open interface{} `json:"open,omitempty"`

	// Opensrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `open`.
	Opensrc types.String `json:"opensrc,omitempty"`

	// Selectedpoints
	// arrayOK: false
	// type: any
	// Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend
	// arrayOK: false
	// type: boolean
	// Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend types.Bool `json:"showlegend,omitempty"`

	// Stream
	// arrayOK: false
	// role: Object
	Stream *OhlcStream `json:"stream,omitempty"`

	// Text
	// arrayOK: true
	// type: string
	// Sets hover text elements associated with each sample point. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to this trace's sample points.
	Text *types.ArrayOK[*types.String] `json:"text,omitempty"`

	// Textsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `text`.
	Textsrc types.String `json:"textsrc,omitempty"`

	// Tickwidth
	// arrayOK: false
	// type: number
	// Sets the width of the open/close tick marks relative to the *x* minimal interval.
	Tickwidth float64 `json:"tickwidth,omitempty"`

	// Transforms
	// role: Object
	// items: OhlcTransform
	Transforms []OhlcTransform `json:"transforms,omitempty"`

	// Uid
	// arrayOK: false
	// type: string
	// Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid types.String `json:"uid,omitempty"`

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
	Visible OhlcVisible `json:"visible,omitempty"`

	// X
	// arrayOK: false
	// type: data_array
	// Sets the x coordinates. If absent, linear coordinate will be generated.
	X interface{} `json:"x,omitempty"`

	// Xaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis types.String `json:"xaxis,omitempty"`

	// Xcalendar
	// arrayOK: false
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `x` date data.
	Xcalendar OhlcXcalendar `json:"xcalendar,omitempty"`

	// Xhoverformat
	// arrayOK: false
	// type: string
	// Sets the hover text formatting rulefor `x`  using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-format/tree/v1.4.5#d3-format. And for dates see: https://github.com/d3/d3-time-format/tree/v2.2.3#locale_format. We add two items to d3's date formatter: *%h* for half of the year as a decimal number as well as *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46*By default the values are formatted using `xaxis.hoverformat`.
	Xhoverformat types.String `json:"xhoverformat,omitempty"`

	// Xperiod
	// arrayOK: false
	// type: any
	// Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the x axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Xperiod interface{} `json:"xperiod,omitempty"`

	// Xperiod0
	// arrayOK: false
	// type: any
	// Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the x0 axis. When `x0period` is round number of weeks, the `x0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Xperiod0 interface{} `json:"xperiod0,omitempty"`

	// Xperiodalignment
	// arrayOK: false
	// default: middle
	// type: enumerated
	// Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
	Xperiodalignment OhlcXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `x`.
	Xsrc types.String `json:"xsrc,omitempty"`

	// Yaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis types.String `json:"yaxis,omitempty"`

	// Yhoverformat
	// arrayOK: false
	// type: string
	// Sets the hover text formatting rulefor `y`  using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-format/tree/v1.4.5#d3-format. And for dates see: https://github.com/d3/d3-time-format/tree/v2.2.3#locale_format. We add two items to d3's date formatter: *%h* for half of the year as a decimal number as well as *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46*By default the values are formatted using `yaxis.hoverformat`.
	Yhoverformat types.String `json:"yhoverformat,omitempty"`

	// Zorder
	// arrayOK: false
	// type: integer
	// Sets the layer on which this trace is displayed, relative to other SVG traces on the same subplot. SVG traces with higher `zorder` appear in front of those with lower `zorder`.
	Zorder int64 `json:"zorder,omitempty"`
}

// OhlcDecreasingLine
type OhlcDecreasingLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the line color.
	Color types.Color `json:"color,omitempty"`

	// Dash
	// arrayOK: false
	// type: string
	// Sets the dash style of lines. Set to a dash type string (*solid*, *dot*, *dash*, *longdash*, *dashdot*, or *longdashdot*) or a dash length list in px (eg *5px,10px,2px,2px*).
	Dash types.String `json:"dash,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the line width (in px).
	Width float64 `json:"width,omitempty"`
}

// OhlcDecreasing
type OhlcDecreasing struct {

	// Line
	// arrayOK: false
	// role: Object
	Line *OhlcDecreasingLine `json:"line,omitempty"`
}

// OhlcHoverlabelFont Sets the font used in hover labels.
type OhlcHoverlabelFont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color *types.ArrayOK[*types.Color] `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	Colorsrc types.String `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family *types.ArrayOK[*types.String] `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `family`.
	Familysrc types.String `json:"familysrc,omitempty"`

	// Lineposition
	// arrayOK: true
	// default: none
	// type: flaglist
	// Sets the kind of decoration line(s) with text, such as an *under*, *over* or *through* as well as combinations e.g. *under+over*, etc.
	Lineposition *types.ArrayOK[*OhlcHoverlabelFontLineposition] `json:"lineposition,omitempty"`

	// Linepositionsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `lineposition`.
	Linepositionsrc types.String `json:"linepositionsrc,omitempty"`

	// Shadow
	// arrayOK: true
	// type: string
	// Sets the shape and color of the shadow behind text. *auto* places minimal shadow and applies contrast text font color. See https://developer.mozilla.org/en-US/docs/Web/CSS/text-shadow for additional options.
	Shadow *types.ArrayOK[*types.String] `json:"shadow,omitempty"`

	// Shadowsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `shadow`.
	Shadowsrc types.String `json:"shadowsrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size *types.ArrayOK[*float64] `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `size`.
	Sizesrc types.String `json:"sizesrc,omitempty"`

	// Style
	// arrayOK: true
	// default: normal
	// type: enumerated
	// Sets whether a font should be styled with a normal or italic face from its family.
	Style *types.ArrayOK[*OhlcHoverlabelFontStyle] `json:"style,omitempty"`

	// Stylesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `style`.
	Stylesrc types.String `json:"stylesrc,omitempty"`

	// Textcase
	// arrayOK: true
	// default: normal
	// type: enumerated
	// Sets capitalization of text. It can be used to make text appear in all-uppercase or all-lowercase, or with each word capitalized.
	Textcase *types.ArrayOK[*OhlcHoverlabelFontTextcase] `json:"textcase,omitempty"`

	// Textcasesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `textcase`.
	Textcasesrc types.String `json:"textcasesrc,omitempty"`

	// Variant
	// arrayOK: true
	// default: normal
	// type: enumerated
	// Sets the variant of the font.
	Variant *types.ArrayOK[*OhlcHoverlabelFontVariant] `json:"variant,omitempty"`

	// Variantsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `variant`.
	Variantsrc types.String `json:"variantsrc,omitempty"`

	// Weight
	// arrayOK: true
	// type: integer
	// Sets the weight (or boldness) of the font.
	Weight *types.ArrayOK[*int64] `json:"weight,omitempty"`

	// Weightsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `weight`.
	Weightsrc types.String `json:"weightsrc,omitempty"`
}

// OhlcHoverlabel
type OhlcHoverlabel struct {

	// Align
	// arrayOK: true
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align *types.ArrayOK[*OhlcHoverlabelAlign] `json:"align,omitempty"`

	// Alignsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `align`.
	Alignsrc types.String `json:"alignsrc,omitempty"`

	// Bgcolor
	// arrayOK: true
	// type: color
	// Sets the background color of the hover labels for this trace
	Bgcolor *types.ArrayOK[*types.Color] `json:"bgcolor,omitempty"`

	// Bgcolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `bgcolor`.
	Bgcolorsrc types.String `json:"bgcolorsrc,omitempty"`

	// Bordercolor
	// arrayOK: true
	// type: color
	// Sets the border color of the hover labels for this trace.
	Bordercolor *types.ArrayOK[*types.Color] `json:"bordercolor,omitempty"`

	// Bordercolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `bordercolor`.
	Bordercolorsrc types.String `json:"bordercolorsrc,omitempty"`

	// Font
	// arrayOK: false
	// role: Object
	Font *OhlcHoverlabelFont `json:"font,omitempty"`

	// Namelength
	// arrayOK: true
	// type: integer
	// Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis.
	Namelength *types.ArrayOK[*int64] `json:"namelength,omitempty"`

	// Namelengthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `namelength`.
	Namelengthsrc types.String `json:"namelengthsrc,omitempty"`

	// Split
	// arrayOK: false
	// type: boolean
	// Show hover information (open, close, high, low) in separate labels.
	Split types.Bool `json:"split,omitempty"`
}

// OhlcIncreasingLine
type OhlcIncreasingLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the line color.
	Color types.Color `json:"color,omitempty"`

	// Dash
	// arrayOK: false
	// type: string
	// Sets the dash style of lines. Set to a dash type string (*solid*, *dot*, *dash*, *longdash*, *dashdot*, or *longdashdot*) or a dash length list in px (eg *5px,10px,2px,2px*).
	Dash types.String `json:"dash,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the line width (in px).
	Width float64 `json:"width,omitempty"`
}

// OhlcIncreasing
type OhlcIncreasing struct {

	// Line
	// arrayOK: false
	// role: Object
	Line *OhlcIncreasingLine `json:"line,omitempty"`
}

// OhlcLegendgrouptitleFont Sets this legend group's title font.
type OhlcLegendgrouptitleFont struct {

	// Color
	// arrayOK: false
	// type: color
	//
	Color types.Color `json:"color,omitempty"`

	// Family
	// arrayOK: false
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family types.String `json:"family,omitempty"`

	// Lineposition
	// arrayOK: false
	// default: none
	// type: flaglist
	// Sets the kind of decoration line(s) with text, such as an *under*, *over* or *through* as well as combinations e.g. *under+over*, etc.
	Lineposition OhlcLegendgrouptitleFontLineposition `json:"lineposition,omitempty"`

	// Shadow
	// arrayOK: false
	// type: string
	// Sets the shape and color of the shadow behind text. *auto* places minimal shadow and applies contrast text font color. See https://developer.mozilla.org/en-US/docs/Web/CSS/text-shadow for additional options.
	Shadow types.String `json:"shadow,omitempty"`

	// Size
	// arrayOK: false
	// type: number
	//
	Size float64 `json:"size,omitempty"`

	// Style
	// arrayOK: false
	// default: normal
	// type: enumerated
	// Sets whether a font should be styled with a normal or italic face from its family.
	Style OhlcLegendgrouptitleFontStyle `json:"style,omitempty"`

	// Textcase
	// arrayOK: false
	// default: normal
	// type: enumerated
	// Sets capitalization of text. It can be used to make text appear in all-uppercase or all-lowercase, or with each word capitalized.
	Textcase OhlcLegendgrouptitleFontTextcase `json:"textcase,omitempty"`

	// Variant
	// arrayOK: false
	// default: normal
	// type: enumerated
	// Sets the variant of the font.
	Variant OhlcLegendgrouptitleFontVariant `json:"variant,omitempty"`

	// Weight
	// arrayOK: false
	// type: integer
	// Sets the weight (or boldness) of the font.
	Weight int64 `json:"weight,omitempty"`
}

// OhlcLegendgrouptitle
type OhlcLegendgrouptitle struct {

	// Font
	// arrayOK: false
	// role: Object
	Font *OhlcLegendgrouptitleFont `json:"font,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the legend group.
	Text types.String `json:"text,omitempty"`
}

// OhlcLine
type OhlcLine struct {

	// Dash
	// arrayOK: false
	// type: string
	// Sets the dash style of lines. Set to a dash type string (*solid*, *dot*, *dash*, *longdash*, *dashdot*, or *longdashdot*) or a dash length list in px (eg *5px,10px,2px,2px*). Note that this style setting can also be set per direction via `increasing.line.dash` and `decreasing.line.dash`.
	Dash types.String `json:"dash,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// [object Object] Note that this style setting can also be set per direction via `increasing.line.width` and `decreasing.line.width`.
	Width float64 `json:"width,omitempty"`
}

// OhlcStream
type OhlcStream struct {

	// Maxpoints
	// arrayOK: false
	// type: number
	// Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot.
	Maxpoints float64 `json:"maxpoints,omitempty"`

	// Token
	// arrayOK: false
	// type: string
	// The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details.
	Token types.String `json:"token,omitempty"`
}

// OhlcTransform WARNING: All transforms are deprecated and may be removed from the API in next major version. An array of operations that manipulate the trace data, for example filtering or sorting the data arrays.
type OhlcTransform struct {
}

// OhlcHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type OhlcHoverlabelAlign string

const (
	OhlcHoverlabelAlignLeft  OhlcHoverlabelAlign = "left"
	OhlcHoverlabelAlignRight OhlcHoverlabelAlign = "right"
	OhlcHoverlabelAlignAuto  OhlcHoverlabelAlign = "auto"
)

// OhlcHoverlabelFontStyle Sets whether a font should be styled with a normal or italic face from its family.
type OhlcHoverlabelFontStyle string

const (
	OhlcHoverlabelFontStyleNormal OhlcHoverlabelFontStyle = "normal"
	OhlcHoverlabelFontStyleItalic OhlcHoverlabelFontStyle = "italic"
)

// OhlcHoverlabelFontTextcase Sets capitalization of text. It can be used to make text appear in all-uppercase or all-lowercase, or with each word capitalized.
type OhlcHoverlabelFontTextcase string

const (
	OhlcHoverlabelFontTextcaseNormal   OhlcHoverlabelFontTextcase = "normal"
	OhlcHoverlabelFontTextcaseWordCaps OhlcHoverlabelFontTextcase = "word caps"
	OhlcHoverlabelFontTextcaseUpper    OhlcHoverlabelFontTextcase = "upper"
	OhlcHoverlabelFontTextcaseLower    OhlcHoverlabelFontTextcase = "lower"
)

// OhlcHoverlabelFontVariant Sets the variant of the font.
type OhlcHoverlabelFontVariant string

const (
	OhlcHoverlabelFontVariantNormal        OhlcHoverlabelFontVariant = "normal"
	OhlcHoverlabelFontVariantSmallCaps     OhlcHoverlabelFontVariant = "small-caps"
	OhlcHoverlabelFontVariantAllSmallCaps  OhlcHoverlabelFontVariant = "all-small-caps"
	OhlcHoverlabelFontVariantAllPetiteCaps OhlcHoverlabelFontVariant = "all-petite-caps"
	OhlcHoverlabelFontVariantPetiteCaps    OhlcHoverlabelFontVariant = "petite-caps"
	OhlcHoverlabelFontVariantUnicase       OhlcHoverlabelFontVariant = "unicase"
)

// OhlcLegendgrouptitleFontStyle Sets whether a font should be styled with a normal or italic face from its family.
type OhlcLegendgrouptitleFontStyle string

const (
	OhlcLegendgrouptitleFontStyleNormal OhlcLegendgrouptitleFontStyle = "normal"
	OhlcLegendgrouptitleFontStyleItalic OhlcLegendgrouptitleFontStyle = "italic"
)

// OhlcLegendgrouptitleFontTextcase Sets capitalization of text. It can be used to make text appear in all-uppercase or all-lowercase, or with each word capitalized.
type OhlcLegendgrouptitleFontTextcase string

const (
	OhlcLegendgrouptitleFontTextcaseNormal   OhlcLegendgrouptitleFontTextcase = "normal"
	OhlcLegendgrouptitleFontTextcaseWordCaps OhlcLegendgrouptitleFontTextcase = "word caps"
	OhlcLegendgrouptitleFontTextcaseUpper    OhlcLegendgrouptitleFontTextcase = "upper"
	OhlcLegendgrouptitleFontTextcaseLower    OhlcLegendgrouptitleFontTextcase = "lower"
)

// OhlcLegendgrouptitleFontVariant Sets the variant of the font.
type OhlcLegendgrouptitleFontVariant string

const (
	OhlcLegendgrouptitleFontVariantNormal        OhlcLegendgrouptitleFontVariant = "normal"
	OhlcLegendgrouptitleFontVariantSmallCaps     OhlcLegendgrouptitleFontVariant = "small-caps"
	OhlcLegendgrouptitleFontVariantAllSmallCaps  OhlcLegendgrouptitleFontVariant = "all-small-caps"
	OhlcLegendgrouptitleFontVariantAllPetiteCaps OhlcLegendgrouptitleFontVariant = "all-petite-caps"
	OhlcLegendgrouptitleFontVariantPetiteCaps    OhlcLegendgrouptitleFontVariant = "petite-caps"
	OhlcLegendgrouptitleFontVariantUnicase       OhlcLegendgrouptitleFontVariant = "unicase"
)

// OhlcVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type OhlcVisible interface{}

var (
	OhlcVisibleTrue       OhlcVisible = true
	OhlcVisibleFalse      OhlcVisible = false
	OhlcVisibleLegendonly OhlcVisible = "legendonly"
)

// OhlcXcalendar Sets the calendar system to use with `x` date data.
type OhlcXcalendar string

const (
	OhlcXcalendarChinese    OhlcXcalendar = "chinese"
	OhlcXcalendarCoptic     OhlcXcalendar = "coptic"
	OhlcXcalendarDiscworld  OhlcXcalendar = "discworld"
	OhlcXcalendarEthiopian  OhlcXcalendar = "ethiopian"
	OhlcXcalendarGregorian  OhlcXcalendar = "gregorian"
	OhlcXcalendarHebrew     OhlcXcalendar = "hebrew"
	OhlcXcalendarIslamic    OhlcXcalendar = "islamic"
	OhlcXcalendarJalali     OhlcXcalendar = "jalali"
	OhlcXcalendarJulian     OhlcXcalendar = "julian"
	OhlcXcalendarMayan      OhlcXcalendar = "mayan"
	OhlcXcalendarNanakshahi OhlcXcalendar = "nanakshahi"
	OhlcXcalendarNepali     OhlcXcalendar = "nepali"
	OhlcXcalendarPersian    OhlcXcalendar = "persian"
	OhlcXcalendarTaiwan     OhlcXcalendar = "taiwan"
	OhlcXcalendarThai       OhlcXcalendar = "thai"
	OhlcXcalendarUmmalqura  OhlcXcalendar = "ummalqura"
)

// OhlcXperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
type OhlcXperiodalignment string

const (
	OhlcXperiodalignmentStart  OhlcXperiodalignment = "start"
	OhlcXperiodalignmentMiddle OhlcXperiodalignment = "middle"
	OhlcXperiodalignmentEnd    OhlcXperiodalignment = "end"
)

// OhlcHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type OhlcHoverinfo string

const (
	// Flags
	OhlcHoverinfoX    OhlcHoverinfo = "x"
	OhlcHoverinfoY    OhlcHoverinfo = "y"
	OhlcHoverinfoZ    OhlcHoverinfo = "z"
	OhlcHoverinfoText OhlcHoverinfo = "text"
	OhlcHoverinfoName OhlcHoverinfo = "name"

	// Extra
	OhlcHoverinfoAll  OhlcHoverinfo = "all"
	OhlcHoverinfoNone OhlcHoverinfo = "none"
	OhlcHoverinfoSkip OhlcHoverinfo = "skip"
)

// OhlcHoverlabelFontLineposition Sets the kind of decoration line(s) with text, such as an *under*, *over* or *through* as well as combinations e.g. *under+over*, etc.
type OhlcHoverlabelFontLineposition string

const (
	// Flags
	OhlcHoverlabelFontLinepositionUnder   OhlcHoverlabelFontLineposition = "under"
	OhlcHoverlabelFontLinepositionOver    OhlcHoverlabelFontLineposition = "over"
	OhlcHoverlabelFontLinepositionThrough OhlcHoverlabelFontLineposition = "through"

	// Extra
	OhlcHoverlabelFontLinepositionNone OhlcHoverlabelFontLineposition = "none"
)

// OhlcLegendgrouptitleFontLineposition Sets the kind of decoration line(s) with text, such as an *under*, *over* or *through* as well as combinations e.g. *under+over*, etc.
type OhlcLegendgrouptitleFontLineposition string

const (
	// Flags
	OhlcLegendgrouptitleFontLinepositionUnder   OhlcLegendgrouptitleFontLineposition = "under"
	OhlcLegendgrouptitleFontLinepositionOver    OhlcLegendgrouptitleFontLineposition = "over"
	OhlcLegendgrouptitleFontLinepositionThrough OhlcLegendgrouptitleFontLineposition = "through"

	// Extra
	OhlcLegendgrouptitleFontLinepositionNone OhlcLegendgrouptitleFontLineposition = "none"
)
