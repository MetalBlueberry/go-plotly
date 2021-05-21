package grob

var TraceTypeScattercarpet TraceType = "scattercarpet"

func (trace *Scattercarpet) GetType() TraceType {
	return TraceTypeScattercarpet
}

// Scattercarpet Plots a scatter trace on either the first carpet axis or the carpet axis with a matching `carpet` attribute.
type Scattercarpet struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// A
	// arrayOK: false
	// type: data_array
	// Sets the a-axis coordinates.
	A interface{} `json:"a,omitempty"`

	// Asrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  a .
	Asrc String `json:"asrc,omitempty"`

	// B
	// arrayOK: false
	// type: data_array
	// Sets the b-axis coordinates.
	B interface{} `json:"b,omitempty"`

	// Bsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  b .
	Bsrc String `json:"bsrc,omitempty"`

	// Carpet
	// arrayOK: false
	// type: string
	// An identifier for this carpet, so that `scattercarpet` and `contourcarpet` traces can specify a carpet plot on which they lie
	Carpet String `json:"carpet,omitempty"`

	// Connectgaps
	// arrayOK: false
	// type: boolean
	// Determines whether or not gaps (i.e. {nan} or missing values) in the provided data arrays are connected.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Customdata
	// arrayOK: false
	// type: data_array
	// Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Fill
	// default: none
	// type: enumerated
	// Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. scatterternary has a subset of the options available to scatter. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other.
	Fill ScattercarpetFill `json:"fill,omitempty"`

	// Fillcolor
	// arrayOK: false
	// type: color
	// Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor Color `json:"fillcolor,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ScattercarpetHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *ScattercarpetHoverlabel `json:"hoverlabel,omitempty"`

	// Hoveron
	// default: %!s(<nil>)
	// type: flaglist
	// Do the hover effects highlight individual points (markers or line points) or do they highlight filled regions? If the fill is *toself* or *tonext* and there are no markers or text, then the default is *fills*, otherwise it is *points*.
	Hoveron ScattercarpetHoveron `json:"hoveron,omitempty"`

	// Hovertemplate
	// arrayOK: true
	// type: string
	// Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate String `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext
	// arrayOK: true
	// type: string
	// Sets hover text elements associated with each (a,b) point. If a single string, the same string appears over all the data points. If an array of strings, the items are mapped in order to the the data points in (a,b). To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext String `json:"hovertext,omitempty"`

	// Hovertextsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids
	// arrayOK: false
	// type: data_array
	// Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup
	// arrayOK: false
	// type: string
	// Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line
	// role: Object
	Line *ScattercarpetLine `json:"line,omitempty"`

	// Marker
	// role: Object
	Marker *ScattercarpetMarker `json:"marker,omitempty"`

	// Meta
	// arrayOK: true
	// type: any
	// Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Mode
	// default: markers
	// type: flaglist
	// Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
	Mode ScattercarpetMode `json:"mode,omitempty"`

	// Name
	// arrayOK: false
	// type: string
	// Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Selected
	// role: Object
	Selected *ScattercarpetSelected `json:"selected,omitempty"`

	// Selectedpoints
	// arrayOK: false
	// type: any
	// Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend
	// arrayOK: false
	// type: boolean
	// Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream
	// role: Object
	Stream *ScattercarpetStream `json:"stream,omitempty"`

	// Text
	// arrayOK: true
	// type: string
	// Sets text elements associated with each (a,b) point. If a single string, the same string appears over all the data points. If an array of strings, the items are mapped in order to the the data points in (a,b). If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text String `json:"text,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScattercarpetTextfont `json:"textfont,omitempty"`

	// Textposition
	// default: middle center
	// type: enumerated
	// Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition ScattercarpetTextposition `json:"textposition,omitempty"`

	// Textpositionsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate
	// arrayOK: true
	// type: string
	// Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `a`, `b` and `text`.
	Texttemplate String `json:"texttemplate,omitempty"`

	// Texttemplatesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Transforms
	// It's an items array and what goes inside it's... messy... check the docs
	// I will be happy if you want to contribute by implementing this
	// just raise an issue before you start so we do not overlap
	Transforms interface{} `json:"transforms,omitempty"`

	// Uid
	// arrayOK: false
	// type: string
	// Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision
	// arrayOK: false
	// type: any
	// Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected
	// role: Object
	Unselected *ScattercarpetUnselected `json:"unselected,omitempty"`

	// Visible
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ScattercarpetVisible `json:"visible,omitempty"`

	// Xaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Yaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`
}

// ScattercarpetHoverlabelFont Sets the font used in hover labels.
type ScattercarpetHoverlabelFont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color Color `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family String `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  family .
	Familysrc String `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size float64 `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  size .
	Sizesrc String `json:"sizesrc,omitempty"`
}

// ScattercarpetHoverlabel
type ScattercarpetHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align ScattercarpetHoverlabelAlign `json:"align,omitempty"`

	// Alignsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  align .
	Alignsrc String `json:"alignsrc,omitempty"`

	// Bgcolor
	// arrayOK: true
	// type: color
	// Sets the background color of the hover labels for this trace
	Bgcolor Color `json:"bgcolor,omitempty"`

	// Bgcolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  bgcolor .
	Bgcolorsrc String `json:"bgcolorsrc,omitempty"`

	// Bordercolor
	// arrayOK: true
	// type: color
	// Sets the border color of the hover labels for this trace.
	Bordercolor Color `json:"bordercolor,omitempty"`

	// Bordercolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  bordercolor .
	Bordercolorsrc String `json:"bordercolorsrc,omitempty"`

	// Font
	// role: Object
	Font *ScattercarpetHoverlabelFont `json:"font,omitempty"`

	// Namelength
	// arrayOK: true
	// type: integer
	// Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis.
	Namelength int64 `json:"namelength,omitempty"`

	// Namelengthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  namelength .
	Namelengthsrc String `json:"namelengthsrc,omitempty"`
}

// ScattercarpetLine
type ScattercarpetLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the line color.
	Color Color `json:"color,omitempty"`

	// Dash
	// arrayOK: false
	// type: string
	// Sets the dash style of lines. Set to a dash type string (*solid*, *dot*, *dash*, *longdash*, *dashdot*, or *longdashdot*) or a dash length list in px (eg *5px,10px,2px,2px*).
	Dash String `json:"dash,omitempty"`

	// Shape
	// default: linear
	// type: enumerated
	// Determines the line shape. With *spline* the lines are drawn using spline interpolation. The other available values correspond to step-wise line shapes.
	Shape ScattercarpetLineShape `json:"shape,omitempty"`

	// Smoothing
	// arrayOK: false
	// type: number
	// Has an effect only if `shape` is set to *spline* Sets the amount of smoothing. *0* corresponds to no smoothing (equivalent to a *linear* shape).
	Smoothing float64 `json:"smoothing,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the line width (in px).
	Width float64 `json:"width,omitempty"`
}

// ScattercarpetMarkerColorbarTickfont Sets the color bar's tick label font
type ScattercarpetMarkerColorbarTickfont struct {

	// Color
	// arrayOK: false
	// type: color
	//
	Color Color `json:"color,omitempty"`

	// Family
	// arrayOK: false
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family String `json:"family,omitempty"`

	// Size
	// arrayOK: false
	// type: number
	//
	Size float64 `json:"size,omitempty"`
}

// ScattercarpetMarkerColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type ScattercarpetMarkerColorbarTitleFont struct {

	// Color
	// arrayOK: false
	// type: color
	//
	Color Color `json:"color,omitempty"`

	// Family
	// arrayOK: false
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family String `json:"family,omitempty"`

	// Size
	// arrayOK: false
	// type: number
	//
	Size float64 `json:"size,omitempty"`
}

// ScattercarpetMarkerColorbarTitle
type ScattercarpetMarkerColorbarTitle struct {

	// Font
	// role: Object
	Font *ScattercarpetMarkerColorbarTitleFont `json:"font,omitempty"`

	// Side
	// default: top
	// type: enumerated
	// Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
	Side ScattercarpetMarkerColorbarTitleSide `json:"side,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// ScattercarpetMarkerColorbar
type ScattercarpetMarkerColorbar struct {

	// Bgcolor
	// arrayOK: false
	// type: color
	// Sets the color of padded area.
	Bgcolor Color `json:"bgcolor,omitempty"`

	// Bordercolor
	// arrayOK: false
	// type: color
	// Sets the axis line color.
	Bordercolor Color `json:"bordercolor,omitempty"`

	// Borderwidth
	// arrayOK: false
	// type: number
	// Sets the width (in px) or the border enclosing this color bar.
	Borderwidth float64 `json:"borderwidth,omitempty"`

	// Dtick
	// arrayOK: false
	// type: any
	// Sets the step in-between ticks on this axis. Use with `tick0`. Must be a positive number, or special strings available to *log* and *date* axes. If the axis `type` is *log*, then ticks are set every 10^(n*dtick) where n is the tick number. For example, to set a tick mark at 1, 10, 100, 1000, ... set dtick to 1. To set tick marks at 1, 100, 10000, ... set dtick to 2. To set tick marks at 1, 5, 25, 125, 625, 3125, ... set dtick to log_10(5), or 0.69897000433. *log* has several special values; *L<f>*, where `f` is a positive number, gives ticks linearly spaced in value (but not position). For example `tick0` = 0.1, `dtick` = *L0.5* will put ticks at 0.1, 0.6, 1.1, 1.6 etc. To show powers of 10 plus small digits between, use *D1* (all digits) or *D2* (only 2 and 5). `tick0` is ignored for *D1* and *D2*. If the axis `type` is *date*, then you must convert the time to milliseconds. For example, to set the interval between ticks to one day, set `dtick` to 86400000.0. *date* also has special values *M<n>* gives ticks spaced by a number of months. `n` must be a positive integer. To set ticks on the 15th of every third month, set `tick0` to *2000-01-15* and `dtick` to *M3*. To set ticks every 4 years, set `dtick` to *M48*
	Dtick interface{} `json:"dtick,omitempty"`

	// Exponentformat
	// default: B
	// type: enumerated
	// Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
	Exponentformat ScattercarpetMarkerColorbarExponentformat `json:"exponentformat,omitempty"`

	// Len
	// arrayOK: false
	// type: number
	// Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends.
	Len float64 `json:"len,omitempty"`

	// Lenmode
	// default: fraction
	// type: enumerated
	// Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
	Lenmode ScattercarpetMarkerColorbarLenmode `json:"lenmode,omitempty"`

	// Minexponent
	// arrayOK: false
	// type: number
	// Hide SI prefix for 10^n if |n| is below this number. This only has an effect when `tickformat` is *SI* or *B*.
	Minexponent float64 `json:"minexponent,omitempty"`

	// Nticks
	// arrayOK: false
	// type: integer
	// Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*.
	Nticks int64 `json:"nticks,omitempty"`

	// Outlinecolor
	// arrayOK: false
	// type: color
	// Sets the axis line color.
	Outlinecolor Color `json:"outlinecolor,omitempty"`

	// Outlinewidth
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the axis line.
	Outlinewidth float64 `json:"outlinewidth,omitempty"`

	// Separatethousands
	// arrayOK: false
	// type: boolean
	// If "true", even 4-digit integers are separated
	Separatethousands Bool `json:"separatethousands,omitempty"`

	// Showexponent
	// default: all
	// type: enumerated
	// If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
	Showexponent ScattercarpetMarkerColorbarShowexponent `json:"showexponent,omitempty"`

	// Showticklabels
	// arrayOK: false
	// type: boolean
	// Determines whether or not the tick labels are drawn.
	Showticklabels Bool `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix ScattercarpetMarkerColorbarShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix ScattercarpetMarkerColorbarShowticksuffix `json:"showticksuffix,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels.
	Thickness float64 `json:"thickness,omitempty"`

	// Thicknessmode
	// default: pixels
	// type: enumerated
	// Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
	Thicknessmode ScattercarpetMarkerColorbarThicknessmode `json:"thicknessmode,omitempty"`

	// Tick0
	// arrayOK: false
	// type: any
	// Sets the placement of the first tick on this axis. Use with `dtick`. If the axis `type` is *log*, then you must take the log of your starting tick (e.g. to set the starting tick to 100, set the `tick0` to 2) except when `dtick`=*L<f>* (see `dtick` for more info). If the axis `type` is *date*, it should be a date string, like date data. If the axis `type` is *category*, it should be a number, using the scale where each category is assigned a serial number from zero in the order it appears.
	Tick0 interface{} `json:"tick0,omitempty"`

	// Tickangle
	// arrayOK: false
	// type: angle
	// Sets the angle of the tick labels with respect to the horizontal. For example, a `tickangle` of -90 draws the tick labels vertically.
	Tickangle float64 `json:"tickangle,omitempty"`

	// Tickcolor
	// arrayOK: false
	// type: color
	// Sets the tick color.
	Tickcolor Color `json:"tickcolor,omitempty"`

	// Tickfont
	// role: Object
	Tickfont *ScattercarpetMarkerColorbarTickfont `json:"tickfont,omitempty"`

	// Tickformat
	// arrayOK: false
	// type: string
	// Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46*
	Tickformat String `json:"tickformat,omitempty"`

	// Tickformatstops
	// It's an items array and what goes inside it's... messy... check the docs
	// I will be happy if you want to contribute by implementing this
	// just raise an issue before you start so we do not overlap
	Tickformatstops interface{} `json:"tickformatstops,omitempty"`

	// Ticklabelposition
	// default: outside
	// type: enumerated
	// Determines where tick labels are drawn.
	Ticklabelposition ScattercarpetMarkerColorbarTicklabelposition `json:"ticklabelposition,omitempty"`

	// Ticklen
	// arrayOK: false
	// type: number
	// Sets the tick length (in px).
	Ticklen float64 `json:"ticklen,omitempty"`

	// Tickmode
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
	Tickmode ScattercarpetMarkerColorbarTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

	// Ticks
	// default:
	// type: enumerated
	// Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
	Ticks ScattercarpetMarkerColorbarTicks `json:"ticks,omitempty"`

	// Ticksuffix
	// arrayOK: false
	// type: string
	// Sets a tick label suffix.
	Ticksuffix String `json:"ticksuffix,omitempty"`

	// Ticktext
	// arrayOK: false
	// type: data_array
	// Sets the text displayed at the ticks position via `tickvals`. Only has an effect if `tickmode` is set to *array*. Used with `tickvals`.
	Ticktext interface{} `json:"ticktext,omitempty"`

	// Ticktextsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  ticktext .
	Ticktextsrc String `json:"ticktextsrc,omitempty"`

	// Tickvals
	// arrayOK: false
	// type: data_array
	// Sets the values at which ticks on this axis appear. Only has an effect if `tickmode` is set to *array*. Used with `ticktext`.
	Tickvals interface{} `json:"tickvals,omitempty"`

	// Tickvalssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  tickvals .
	Tickvalssrc String `json:"tickvalssrc,omitempty"`

	// Tickwidth
	// arrayOK: false
	// type: number
	// Sets the tick width (in px).
	Tickwidth float64 `json:"tickwidth,omitempty"`

	// Title
	// role: Object
	Title *ScattercarpetMarkerColorbarTitle `json:"title,omitempty"`

	// X
	// arrayOK: false
	// type: number
	// Sets the x position of the color bar (in plot fraction).
	X float64 `json:"x,omitempty"`

	// Xanchor
	// default: left
	// type: enumerated
	// Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
	Xanchor ScattercarpetMarkerColorbarXanchor `json:"xanchor,omitempty"`

	// Xpad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the x direction.
	Xpad float64 `json:"xpad,omitempty"`

	// Y
	// arrayOK: false
	// type: number
	// Sets the y position of the color bar (in plot fraction).
	Y float64 `json:"y,omitempty"`

	// Yanchor
	// default: middle
	// type: enumerated
	// Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
	Yanchor ScattercarpetMarkerColorbarYanchor `json:"yanchor,omitempty"`

	// Ypad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the y direction.
	Ypad float64 `json:"ypad,omitempty"`
}

// ScattercarpetMarkerGradient
type ScattercarpetMarkerGradient struct {

	// Color
	// arrayOK: true
	// type: color
	// Sets the final color of the gradient fill: the center color for radial, the right for horizontal, or the bottom for vertical.
	Color Color `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Type
	// default: none
	// type: enumerated
	// Sets the type of gradient used to fill the markers
	Type ScattercarpetMarkerGradientType `json:"type,omitempty"`

	// Typesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  type .
	Typesrc String `json:"typesrc,omitempty"`
}

// ScattercarpetMarkerLine
type ScattercarpetMarkerLine struct {

	// Autocolorscale
	// arrayOK: false
	// type: boolean
	// Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `marker.line.colorscale`. Has an effect only if in `marker.line.color`is set to a numerical array. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Cauto
	// arrayOK: false
	// type: boolean
	// Determines whether or not the color domain is computed with respect to the input data (here in `marker.line.color`) or the bounds set in `marker.line.cmin` and `marker.line.cmax`  Has an effect only if in `marker.line.color`is set to a numerical array. Defaults to `false` when `marker.line.cmin` and `marker.line.cmax` are set by the user.
	Cauto Bool `json:"cauto,omitempty"`

	// Cmax
	// arrayOK: false
	// type: number
	// Sets the upper bound of the color domain. Has an effect only if in `marker.line.color`is set to a numerical array. Value should have the same units as in `marker.line.color` and if set, `marker.line.cmin` must be set as well.
	Cmax float64 `json:"cmax,omitempty"`

	// Cmid
	// arrayOK: false
	// type: number
	// Sets the mid-point of the color domain by scaling `marker.line.cmin` and/or `marker.line.cmax` to be equidistant to this point. Has an effect only if in `marker.line.color`is set to a numerical array. Value should have the same units as in `marker.line.color`. Has no effect when `marker.line.cauto` is `false`.
	Cmid float64 `json:"cmid,omitempty"`

	// Cmin
	// arrayOK: false
	// type: number
	// Sets the lower bound of the color domain. Has an effect only if in `marker.line.color`is set to a numerical array. Value should have the same units as in `marker.line.color` and if set, `marker.line.cmax` must be set as well.
	Cmin float64 `json:"cmin,omitempty"`

	// Color
	// arrayOK: true
	// type: color
	// Sets themarker.linecolor. It accepts either a specific color or an array of numbers that are mapped to the colorscale relative to the max and min values of the array or relative to `marker.line.cmin` and `marker.line.cmax` if set.
	Color Color `json:"color,omitempty"`

	// Coloraxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorscale
	// default: %!s(<nil>)
	// type: colorscale
	// Sets the colorscale. Has an effect only if in `marker.line.color`is set to a numerical array. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`marker.line.cmin` and `marker.line.cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	Colorscale ColorScale `json:"colorscale,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Reversescale
	// arrayOK: false
	// type: boolean
	// Reverses the color mapping if true. Has an effect only if in `marker.line.color`is set to a numerical array. If true, `marker.line.cmin` will correspond to the last color in the array and `marker.line.cmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Width
	// arrayOK: true
	// type: number
	// Sets the width (in px) of the lines bounding the marker points.
	Width float64 `json:"width,omitempty"`

	// Widthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  width .
	Widthsrc String `json:"widthsrc,omitempty"`
}

// ScattercarpetMarker
type ScattercarpetMarker struct {

	// Autocolorscale
	// arrayOK: false
	// type: boolean
	// Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `marker.colorscale`. Has an effect only if in `marker.color`is set to a numerical array. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Cauto
	// arrayOK: false
	// type: boolean
	// Determines whether or not the color domain is computed with respect to the input data (here in `marker.color`) or the bounds set in `marker.cmin` and `marker.cmax`  Has an effect only if in `marker.color`is set to a numerical array. Defaults to `false` when `marker.cmin` and `marker.cmax` are set by the user.
	Cauto Bool `json:"cauto,omitempty"`

	// Cmax
	// arrayOK: false
	// type: number
	// Sets the upper bound of the color domain. Has an effect only if in `marker.color`is set to a numerical array. Value should have the same units as in `marker.color` and if set, `marker.cmin` must be set as well.
	Cmax float64 `json:"cmax,omitempty"`

	// Cmid
	// arrayOK: false
	// type: number
	// Sets the mid-point of the color domain by scaling `marker.cmin` and/or `marker.cmax` to be equidistant to this point. Has an effect only if in `marker.color`is set to a numerical array. Value should have the same units as in `marker.color`. Has no effect when `marker.cauto` is `false`.
	Cmid float64 `json:"cmid,omitempty"`

	// Cmin
	// arrayOK: false
	// type: number
	// Sets the lower bound of the color domain. Has an effect only if in `marker.color`is set to a numerical array. Value should have the same units as in `marker.color` and if set, `marker.cmax` must be set as well.
	Cmin float64 `json:"cmin,omitempty"`

	// Color
	// arrayOK: true
	// type: color
	// Sets themarkercolor. It accepts either a specific color or an array of numbers that are mapped to the colorscale relative to the max and min values of the array or relative to `marker.cmin` and `marker.cmax` if set.
	Color Color `json:"color,omitempty"`

	// Coloraxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar
	// role: Object
	Colorbar *ScattercarpetMarkerColorbar `json:"colorbar,omitempty"`

	// Colorscale
	// default: %!s(<nil>)
	// type: colorscale
	// Sets the colorscale. Has an effect only if in `marker.color`is set to a numerical array. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`marker.cmin` and `marker.cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	Colorscale ColorScale `json:"colorscale,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Gradient
	// role: Object
	Gradient *ScattercarpetMarkerGradient `json:"gradient,omitempty"`

	// Line
	// role: Object
	Line *ScattercarpetMarkerLine `json:"line,omitempty"`

	// Maxdisplayed
	// arrayOK: false
	// type: number
	// Sets a maximum number of points to be drawn on the graph. *0* corresponds to no limit.
	Maxdisplayed float64 `json:"maxdisplayed,omitempty"`

	// Opacity
	// arrayOK: true
	// type: number
	// Sets the marker opacity.
	Opacity float64 `json:"opacity,omitempty"`

	// Opacitysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  opacity .
	Opacitysrc String `json:"opacitysrc,omitempty"`

	// Reversescale
	// arrayOK: false
	// type: boolean
	// Reverses the color mapping if true. Has an effect only if in `marker.color`is set to a numerical array. If true, `marker.cmin` will correspond to the last color in the array and `marker.cmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Showscale
	// arrayOK: false
	// type: boolean
	// Determines whether or not a colorbar is displayed for this trace. Has an effect only if in `marker.color`is set to a numerical array.
	Showscale Bool `json:"showscale,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	// Sets the marker size (in px).
	Size float64 `json:"size,omitempty"`

	// Sizemin
	// arrayOK: false
	// type: number
	// Has an effect only if `marker.size` is set to a numerical array. Sets the minimum size (in px) of the rendered marker points.
	Sizemin float64 `json:"sizemin,omitempty"`

	// Sizemode
	// default: diameter
	// type: enumerated
	// Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
	Sizemode ScattercarpetMarkerSizemode `json:"sizemode,omitempty"`

	// Sizeref
	// arrayOK: false
	// type: number
	// Has an effect only if `marker.size` is set to a numerical array. Sets the scale factor used to determine the rendered size of marker points. Use with `sizemin` and `sizemode`.
	Sizeref float64 `json:"sizeref,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  size .
	Sizesrc String `json:"sizesrc,omitempty"`

	// Symbol
	// default: circle
	// type: enumerated
	// Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
	Symbol ScattercarpetMarkerSymbol `json:"symbol,omitempty"`

	// Symbolsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  symbol .
	Symbolsrc String `json:"symbolsrc,omitempty"`
}

// ScattercarpetSelectedMarker
type ScattercarpetSelectedMarker struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the marker color of selected points.
	Color Color `json:"color,omitempty"`

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the marker opacity of selected points.
	Opacity float64 `json:"opacity,omitempty"`

	// Size
	// arrayOK: false
	// type: number
	// Sets the marker size of selected points.
	Size float64 `json:"size,omitempty"`
}

// ScattercarpetSelectedTextfont
type ScattercarpetSelectedTextfont struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the text font color of selected points.
	Color Color `json:"color,omitempty"`
}

// ScattercarpetSelected
type ScattercarpetSelected struct {

	// Marker
	// role: Object
	Marker *ScattercarpetSelectedMarker `json:"marker,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScattercarpetSelectedTextfont `json:"textfont,omitempty"`
}

// ScattercarpetStream
type ScattercarpetStream struct {

	// Maxpoints
	// arrayOK: false
	// type: number
	// Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot.
	Maxpoints float64 `json:"maxpoints,omitempty"`

	// Token
	// arrayOK: false
	// type: string
	// The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details.
	Token String `json:"token,omitempty"`
}

// ScattercarpetTextfont Sets the text font.
type ScattercarpetTextfont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color Color `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family String `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  family .
	Familysrc String `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size float64 `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  size .
	Sizesrc String `json:"sizesrc,omitempty"`
}

// ScattercarpetUnselectedMarker
type ScattercarpetUnselectedMarker struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the marker color of unselected points, applied only when a selection exists.
	Color Color `json:"color,omitempty"`

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the marker opacity of unselected points, applied only when a selection exists.
	Opacity float64 `json:"opacity,omitempty"`

	// Size
	// arrayOK: false
	// type: number
	// Sets the marker size of unselected points, applied only when a selection exists.
	Size float64 `json:"size,omitempty"`
}

// ScattercarpetUnselectedTextfont
type ScattercarpetUnselectedTextfont struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the text font color of unselected points, applied only when a selection exists.
	Color Color `json:"color,omitempty"`
}

// ScattercarpetUnselected
type ScattercarpetUnselected struct {

	// Marker
	// role: Object
	Marker *ScattercarpetUnselectedMarker `json:"marker,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScattercarpetUnselectedTextfont `json:"textfont,omitempty"`
}

// ScattercarpetFill Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. scatterternary has a subset of the options available to scatter. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other.
type ScattercarpetFill string

const (
	ScattercarpetFillNone   ScattercarpetFill = "none"
	ScattercarpetFillToself ScattercarpetFill = "toself"
	ScattercarpetFillTonext ScattercarpetFill = "tonext"
)

// ScattercarpetHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScattercarpetHoverlabelAlign string

const (
	ScattercarpetHoverlabelAlignLeft  ScattercarpetHoverlabelAlign = "left"
	ScattercarpetHoverlabelAlignRight ScattercarpetHoverlabelAlign = "right"
	ScattercarpetHoverlabelAlignAuto  ScattercarpetHoverlabelAlign = "auto"
)

// ScattercarpetLineShape Determines the line shape. With *spline* the lines are drawn using spline interpolation. The other available values correspond to step-wise line shapes.
type ScattercarpetLineShape string

const (
	ScattercarpetLineShapeLinear ScattercarpetLineShape = "linear"
	ScattercarpetLineShapeSpline ScattercarpetLineShape = "spline"
)

// ScattercarpetMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScattercarpetMarkerColorbarExponentformat string

const (
	ScattercarpetMarkerColorbarExponentformatNone  ScattercarpetMarkerColorbarExponentformat = "none"
	ScattercarpetMarkerColorbarExponentformatE1    ScattercarpetMarkerColorbarExponentformat = "e"
	ScattercarpetMarkerColorbarExponentformatE2    ScattercarpetMarkerColorbarExponentformat = "E"
	ScattercarpetMarkerColorbarExponentformatPower ScattercarpetMarkerColorbarExponentformat = "power"
	ScattercarpetMarkerColorbarExponentformatSi    ScattercarpetMarkerColorbarExponentformat = "SI"
	ScattercarpetMarkerColorbarExponentformatB     ScattercarpetMarkerColorbarExponentformat = "B"
)

// ScattercarpetMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScattercarpetMarkerColorbarLenmode string

const (
	ScattercarpetMarkerColorbarLenmodeFraction ScattercarpetMarkerColorbarLenmode = "fraction"
	ScattercarpetMarkerColorbarLenmodePixels   ScattercarpetMarkerColorbarLenmode = "pixels"
)

// ScattercarpetMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScattercarpetMarkerColorbarShowexponent string

const (
	ScattercarpetMarkerColorbarShowexponentAll   ScattercarpetMarkerColorbarShowexponent = "all"
	ScattercarpetMarkerColorbarShowexponentFirst ScattercarpetMarkerColorbarShowexponent = "first"
	ScattercarpetMarkerColorbarShowexponentLast  ScattercarpetMarkerColorbarShowexponent = "last"
	ScattercarpetMarkerColorbarShowexponentNone  ScattercarpetMarkerColorbarShowexponent = "none"
)

// ScattercarpetMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScattercarpetMarkerColorbarShowtickprefix string

const (
	ScattercarpetMarkerColorbarShowtickprefixAll   ScattercarpetMarkerColorbarShowtickprefix = "all"
	ScattercarpetMarkerColorbarShowtickprefixFirst ScattercarpetMarkerColorbarShowtickprefix = "first"
	ScattercarpetMarkerColorbarShowtickprefixLast  ScattercarpetMarkerColorbarShowtickprefix = "last"
	ScattercarpetMarkerColorbarShowtickprefixNone  ScattercarpetMarkerColorbarShowtickprefix = "none"
)

// ScattercarpetMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScattercarpetMarkerColorbarShowticksuffix string

const (
	ScattercarpetMarkerColorbarShowticksuffixAll   ScattercarpetMarkerColorbarShowticksuffix = "all"
	ScattercarpetMarkerColorbarShowticksuffixFirst ScattercarpetMarkerColorbarShowticksuffix = "first"
	ScattercarpetMarkerColorbarShowticksuffixLast  ScattercarpetMarkerColorbarShowticksuffix = "last"
	ScattercarpetMarkerColorbarShowticksuffixNone  ScattercarpetMarkerColorbarShowticksuffix = "none"
)

// ScattercarpetMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScattercarpetMarkerColorbarThicknessmode string

const (
	ScattercarpetMarkerColorbarThicknessmodeFraction ScattercarpetMarkerColorbarThicknessmode = "fraction"
	ScattercarpetMarkerColorbarThicknessmodePixels   ScattercarpetMarkerColorbarThicknessmode = "pixels"
)

// ScattercarpetMarkerColorbarTicklabelposition Determines where tick labels are drawn.
type ScattercarpetMarkerColorbarTicklabelposition string

const (
	ScattercarpetMarkerColorbarTicklabelpositionOutside       ScattercarpetMarkerColorbarTicklabelposition = "outside"
	ScattercarpetMarkerColorbarTicklabelpositionInside        ScattercarpetMarkerColorbarTicklabelposition = "inside"
	ScattercarpetMarkerColorbarTicklabelpositionOutsideTop    ScattercarpetMarkerColorbarTicklabelposition = "outside top"
	ScattercarpetMarkerColorbarTicklabelpositionInsideTop     ScattercarpetMarkerColorbarTicklabelposition = "inside top"
	ScattercarpetMarkerColorbarTicklabelpositionOutsideBottom ScattercarpetMarkerColorbarTicklabelposition = "outside bottom"
	ScattercarpetMarkerColorbarTicklabelpositionInsideBottom  ScattercarpetMarkerColorbarTicklabelposition = "inside bottom"
)

// ScattercarpetMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScattercarpetMarkerColorbarTickmode string

const (
	ScattercarpetMarkerColorbarTickmodeAuto   ScattercarpetMarkerColorbarTickmode = "auto"
	ScattercarpetMarkerColorbarTickmodeLinear ScattercarpetMarkerColorbarTickmode = "linear"
	ScattercarpetMarkerColorbarTickmodeArray  ScattercarpetMarkerColorbarTickmode = "array"
)

// ScattercarpetMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScattercarpetMarkerColorbarTicks string

const (
	ScattercarpetMarkerColorbarTicksOutside ScattercarpetMarkerColorbarTicks = "outside"
	ScattercarpetMarkerColorbarTicksInside  ScattercarpetMarkerColorbarTicks = "inside"
	ScattercarpetMarkerColorbarTicksEmpty   ScattercarpetMarkerColorbarTicks = ""
)

// ScattercarpetMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScattercarpetMarkerColorbarTitleSide string

const (
	ScattercarpetMarkerColorbarTitleSideRight  ScattercarpetMarkerColorbarTitleSide = "right"
	ScattercarpetMarkerColorbarTitleSideTop    ScattercarpetMarkerColorbarTitleSide = "top"
	ScattercarpetMarkerColorbarTitleSideBottom ScattercarpetMarkerColorbarTitleSide = "bottom"
)

// ScattercarpetMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScattercarpetMarkerColorbarXanchor string

const (
	ScattercarpetMarkerColorbarXanchorLeft   ScattercarpetMarkerColorbarXanchor = "left"
	ScattercarpetMarkerColorbarXanchorCenter ScattercarpetMarkerColorbarXanchor = "center"
	ScattercarpetMarkerColorbarXanchorRight  ScattercarpetMarkerColorbarXanchor = "right"
)

// ScattercarpetMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScattercarpetMarkerColorbarYanchor string

const (
	ScattercarpetMarkerColorbarYanchorTop    ScattercarpetMarkerColorbarYanchor = "top"
	ScattercarpetMarkerColorbarYanchorMiddle ScattercarpetMarkerColorbarYanchor = "middle"
	ScattercarpetMarkerColorbarYanchorBottom ScattercarpetMarkerColorbarYanchor = "bottom"
)

// ScattercarpetMarkerGradientType Sets the type of gradient used to fill the markers
type ScattercarpetMarkerGradientType string

const (
	ScattercarpetMarkerGradientTypeRadial     ScattercarpetMarkerGradientType = "radial"
	ScattercarpetMarkerGradientTypeHorizontal ScattercarpetMarkerGradientType = "horizontal"
	ScattercarpetMarkerGradientTypeVertical   ScattercarpetMarkerGradientType = "vertical"
	ScattercarpetMarkerGradientTypeNone       ScattercarpetMarkerGradientType = "none"
)

// ScattercarpetMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScattercarpetMarkerSizemode string

const (
	ScattercarpetMarkerSizemodeDiameter ScattercarpetMarkerSizemode = "diameter"
	ScattercarpetMarkerSizemodeArea     ScattercarpetMarkerSizemode = "area"
)

// ScattercarpetMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScattercarpetMarkerSymbol interface{}

var (
	ScattercarpetMarkerSymbolNumber0                 ScattercarpetMarkerSymbol = 0
	ScattercarpetMarkerSymbol0                       ScattercarpetMarkerSymbol = "0"
	ScattercarpetMarkerSymbolCircle                  ScattercarpetMarkerSymbol = "circle"
	ScattercarpetMarkerSymbolNumber100               ScattercarpetMarkerSymbol = 100
	ScattercarpetMarkerSymbol100                     ScattercarpetMarkerSymbol = "100"
	ScattercarpetMarkerSymbolCircleOpen              ScattercarpetMarkerSymbol = "circle-open"
	ScattercarpetMarkerSymbolNumber200               ScattercarpetMarkerSymbol = 200
	ScattercarpetMarkerSymbol200                     ScattercarpetMarkerSymbol = "200"
	ScattercarpetMarkerSymbolCircleDot               ScattercarpetMarkerSymbol = "circle-dot"
	ScattercarpetMarkerSymbolNumber300               ScattercarpetMarkerSymbol = 300
	ScattercarpetMarkerSymbol300                     ScattercarpetMarkerSymbol = "300"
	ScattercarpetMarkerSymbolCircleOpenDot           ScattercarpetMarkerSymbol = "circle-open-dot"
	ScattercarpetMarkerSymbolNumber1                 ScattercarpetMarkerSymbol = 1
	ScattercarpetMarkerSymbol1                       ScattercarpetMarkerSymbol = "1"
	ScattercarpetMarkerSymbolSquare                  ScattercarpetMarkerSymbol = "square"
	ScattercarpetMarkerSymbolNumber101               ScattercarpetMarkerSymbol = 101
	ScattercarpetMarkerSymbol101                     ScattercarpetMarkerSymbol = "101"
	ScattercarpetMarkerSymbolSquareOpen              ScattercarpetMarkerSymbol = "square-open"
	ScattercarpetMarkerSymbolNumber201               ScattercarpetMarkerSymbol = 201
	ScattercarpetMarkerSymbol201                     ScattercarpetMarkerSymbol = "201"
	ScattercarpetMarkerSymbolSquareDot               ScattercarpetMarkerSymbol = "square-dot"
	ScattercarpetMarkerSymbolNumber301               ScattercarpetMarkerSymbol = 301
	ScattercarpetMarkerSymbol301                     ScattercarpetMarkerSymbol = "301"
	ScattercarpetMarkerSymbolSquareOpenDot           ScattercarpetMarkerSymbol = "square-open-dot"
	ScattercarpetMarkerSymbolNumber2                 ScattercarpetMarkerSymbol = 2
	ScattercarpetMarkerSymbol2                       ScattercarpetMarkerSymbol = "2"
	ScattercarpetMarkerSymbolDiamond                 ScattercarpetMarkerSymbol = "diamond"
	ScattercarpetMarkerSymbolNumber102               ScattercarpetMarkerSymbol = 102
	ScattercarpetMarkerSymbol102                     ScattercarpetMarkerSymbol = "102"
	ScattercarpetMarkerSymbolDiamondOpen             ScattercarpetMarkerSymbol = "diamond-open"
	ScattercarpetMarkerSymbolNumber202               ScattercarpetMarkerSymbol = 202
	ScattercarpetMarkerSymbol202                     ScattercarpetMarkerSymbol = "202"
	ScattercarpetMarkerSymbolDiamondDot              ScattercarpetMarkerSymbol = "diamond-dot"
	ScattercarpetMarkerSymbolNumber302               ScattercarpetMarkerSymbol = 302
	ScattercarpetMarkerSymbol302                     ScattercarpetMarkerSymbol = "302"
	ScattercarpetMarkerSymbolDiamondOpenDot          ScattercarpetMarkerSymbol = "diamond-open-dot"
	ScattercarpetMarkerSymbolNumber3                 ScattercarpetMarkerSymbol = 3
	ScattercarpetMarkerSymbol3                       ScattercarpetMarkerSymbol = "3"
	ScattercarpetMarkerSymbolCross                   ScattercarpetMarkerSymbol = "cross"
	ScattercarpetMarkerSymbolNumber103               ScattercarpetMarkerSymbol = 103
	ScattercarpetMarkerSymbol103                     ScattercarpetMarkerSymbol = "103"
	ScattercarpetMarkerSymbolCrossOpen               ScattercarpetMarkerSymbol = "cross-open"
	ScattercarpetMarkerSymbolNumber203               ScattercarpetMarkerSymbol = 203
	ScattercarpetMarkerSymbol203                     ScattercarpetMarkerSymbol = "203"
	ScattercarpetMarkerSymbolCrossDot                ScattercarpetMarkerSymbol = "cross-dot"
	ScattercarpetMarkerSymbolNumber303               ScattercarpetMarkerSymbol = 303
	ScattercarpetMarkerSymbol303                     ScattercarpetMarkerSymbol = "303"
	ScattercarpetMarkerSymbolCrossOpenDot            ScattercarpetMarkerSymbol = "cross-open-dot"
	ScattercarpetMarkerSymbolNumber4                 ScattercarpetMarkerSymbol = 4
	ScattercarpetMarkerSymbol4                       ScattercarpetMarkerSymbol = "4"
	ScattercarpetMarkerSymbolX                       ScattercarpetMarkerSymbol = "x"
	ScattercarpetMarkerSymbolNumber104               ScattercarpetMarkerSymbol = 104
	ScattercarpetMarkerSymbol104                     ScattercarpetMarkerSymbol = "104"
	ScattercarpetMarkerSymbolXOpen                   ScattercarpetMarkerSymbol = "x-open"
	ScattercarpetMarkerSymbolNumber204               ScattercarpetMarkerSymbol = 204
	ScattercarpetMarkerSymbol204                     ScattercarpetMarkerSymbol = "204"
	ScattercarpetMarkerSymbolXDot                    ScattercarpetMarkerSymbol = "x-dot"
	ScattercarpetMarkerSymbolNumber304               ScattercarpetMarkerSymbol = 304
	ScattercarpetMarkerSymbol304                     ScattercarpetMarkerSymbol = "304"
	ScattercarpetMarkerSymbolXOpenDot                ScattercarpetMarkerSymbol = "x-open-dot"
	ScattercarpetMarkerSymbolNumber5                 ScattercarpetMarkerSymbol = 5
	ScattercarpetMarkerSymbol5                       ScattercarpetMarkerSymbol = "5"
	ScattercarpetMarkerSymbolTriangleUp              ScattercarpetMarkerSymbol = "triangle-up"
	ScattercarpetMarkerSymbolNumber105               ScattercarpetMarkerSymbol = 105
	ScattercarpetMarkerSymbol105                     ScattercarpetMarkerSymbol = "105"
	ScattercarpetMarkerSymbolTriangleUpOpen          ScattercarpetMarkerSymbol = "triangle-up-open"
	ScattercarpetMarkerSymbolNumber205               ScattercarpetMarkerSymbol = 205
	ScattercarpetMarkerSymbol205                     ScattercarpetMarkerSymbol = "205"
	ScattercarpetMarkerSymbolTriangleUpDot           ScattercarpetMarkerSymbol = "triangle-up-dot"
	ScattercarpetMarkerSymbolNumber305               ScattercarpetMarkerSymbol = 305
	ScattercarpetMarkerSymbol305                     ScattercarpetMarkerSymbol = "305"
	ScattercarpetMarkerSymbolTriangleUpOpenDot       ScattercarpetMarkerSymbol = "triangle-up-open-dot"
	ScattercarpetMarkerSymbolNumber6                 ScattercarpetMarkerSymbol = 6
	ScattercarpetMarkerSymbol6                       ScattercarpetMarkerSymbol = "6"
	ScattercarpetMarkerSymbolTriangleDown            ScattercarpetMarkerSymbol = "triangle-down"
	ScattercarpetMarkerSymbolNumber106               ScattercarpetMarkerSymbol = 106
	ScattercarpetMarkerSymbol106                     ScattercarpetMarkerSymbol = "106"
	ScattercarpetMarkerSymbolTriangleDownOpen        ScattercarpetMarkerSymbol = "triangle-down-open"
	ScattercarpetMarkerSymbolNumber206               ScattercarpetMarkerSymbol = 206
	ScattercarpetMarkerSymbol206                     ScattercarpetMarkerSymbol = "206"
	ScattercarpetMarkerSymbolTriangleDownDot         ScattercarpetMarkerSymbol = "triangle-down-dot"
	ScattercarpetMarkerSymbolNumber306               ScattercarpetMarkerSymbol = 306
	ScattercarpetMarkerSymbol306                     ScattercarpetMarkerSymbol = "306"
	ScattercarpetMarkerSymbolTriangleDownOpenDot     ScattercarpetMarkerSymbol = "triangle-down-open-dot"
	ScattercarpetMarkerSymbolNumber7                 ScattercarpetMarkerSymbol = 7
	ScattercarpetMarkerSymbol7                       ScattercarpetMarkerSymbol = "7"
	ScattercarpetMarkerSymbolTriangleLeft            ScattercarpetMarkerSymbol = "triangle-left"
	ScattercarpetMarkerSymbolNumber107               ScattercarpetMarkerSymbol = 107
	ScattercarpetMarkerSymbol107                     ScattercarpetMarkerSymbol = "107"
	ScattercarpetMarkerSymbolTriangleLeftOpen        ScattercarpetMarkerSymbol = "triangle-left-open"
	ScattercarpetMarkerSymbolNumber207               ScattercarpetMarkerSymbol = 207
	ScattercarpetMarkerSymbol207                     ScattercarpetMarkerSymbol = "207"
	ScattercarpetMarkerSymbolTriangleLeftDot         ScattercarpetMarkerSymbol = "triangle-left-dot"
	ScattercarpetMarkerSymbolNumber307               ScattercarpetMarkerSymbol = 307
	ScattercarpetMarkerSymbol307                     ScattercarpetMarkerSymbol = "307"
	ScattercarpetMarkerSymbolTriangleLeftOpenDot     ScattercarpetMarkerSymbol = "triangle-left-open-dot"
	ScattercarpetMarkerSymbolNumber8                 ScattercarpetMarkerSymbol = 8
	ScattercarpetMarkerSymbol8                       ScattercarpetMarkerSymbol = "8"
	ScattercarpetMarkerSymbolTriangleRight           ScattercarpetMarkerSymbol = "triangle-right"
	ScattercarpetMarkerSymbolNumber108               ScattercarpetMarkerSymbol = 108
	ScattercarpetMarkerSymbol108                     ScattercarpetMarkerSymbol = "108"
	ScattercarpetMarkerSymbolTriangleRightOpen       ScattercarpetMarkerSymbol = "triangle-right-open"
	ScattercarpetMarkerSymbolNumber208               ScattercarpetMarkerSymbol = 208
	ScattercarpetMarkerSymbol208                     ScattercarpetMarkerSymbol = "208"
	ScattercarpetMarkerSymbolTriangleRightDot        ScattercarpetMarkerSymbol = "triangle-right-dot"
	ScattercarpetMarkerSymbolNumber308               ScattercarpetMarkerSymbol = 308
	ScattercarpetMarkerSymbol308                     ScattercarpetMarkerSymbol = "308"
	ScattercarpetMarkerSymbolTriangleRightOpenDot    ScattercarpetMarkerSymbol = "triangle-right-open-dot"
	ScattercarpetMarkerSymbolNumber9                 ScattercarpetMarkerSymbol = 9
	ScattercarpetMarkerSymbol9                       ScattercarpetMarkerSymbol = "9"
	ScattercarpetMarkerSymbolTriangleNe              ScattercarpetMarkerSymbol = "triangle-ne"
	ScattercarpetMarkerSymbolNumber109               ScattercarpetMarkerSymbol = 109
	ScattercarpetMarkerSymbol109                     ScattercarpetMarkerSymbol = "109"
	ScattercarpetMarkerSymbolTriangleNeOpen          ScattercarpetMarkerSymbol = "triangle-ne-open"
	ScattercarpetMarkerSymbolNumber209               ScattercarpetMarkerSymbol = 209
	ScattercarpetMarkerSymbol209                     ScattercarpetMarkerSymbol = "209"
	ScattercarpetMarkerSymbolTriangleNeDot           ScattercarpetMarkerSymbol = "triangle-ne-dot"
	ScattercarpetMarkerSymbolNumber309               ScattercarpetMarkerSymbol = 309
	ScattercarpetMarkerSymbol309                     ScattercarpetMarkerSymbol = "309"
	ScattercarpetMarkerSymbolTriangleNeOpenDot       ScattercarpetMarkerSymbol = "triangle-ne-open-dot"
	ScattercarpetMarkerSymbolNumber10                ScattercarpetMarkerSymbol = 10
	ScattercarpetMarkerSymbol10                      ScattercarpetMarkerSymbol = "10"
	ScattercarpetMarkerSymbolTriangleSe              ScattercarpetMarkerSymbol = "triangle-se"
	ScattercarpetMarkerSymbolNumber110               ScattercarpetMarkerSymbol = 110
	ScattercarpetMarkerSymbol110                     ScattercarpetMarkerSymbol = "110"
	ScattercarpetMarkerSymbolTriangleSeOpen          ScattercarpetMarkerSymbol = "triangle-se-open"
	ScattercarpetMarkerSymbolNumber210               ScattercarpetMarkerSymbol = 210
	ScattercarpetMarkerSymbol210                     ScattercarpetMarkerSymbol = "210"
	ScattercarpetMarkerSymbolTriangleSeDot           ScattercarpetMarkerSymbol = "triangle-se-dot"
	ScattercarpetMarkerSymbolNumber310               ScattercarpetMarkerSymbol = 310
	ScattercarpetMarkerSymbol310                     ScattercarpetMarkerSymbol = "310"
	ScattercarpetMarkerSymbolTriangleSeOpenDot       ScattercarpetMarkerSymbol = "triangle-se-open-dot"
	ScattercarpetMarkerSymbolNumber11                ScattercarpetMarkerSymbol = 11
	ScattercarpetMarkerSymbol11                      ScattercarpetMarkerSymbol = "11"
	ScattercarpetMarkerSymbolTriangleSw              ScattercarpetMarkerSymbol = "triangle-sw"
	ScattercarpetMarkerSymbolNumber111               ScattercarpetMarkerSymbol = 111
	ScattercarpetMarkerSymbol111                     ScattercarpetMarkerSymbol = "111"
	ScattercarpetMarkerSymbolTriangleSwOpen          ScattercarpetMarkerSymbol = "triangle-sw-open"
	ScattercarpetMarkerSymbolNumber211               ScattercarpetMarkerSymbol = 211
	ScattercarpetMarkerSymbol211                     ScattercarpetMarkerSymbol = "211"
	ScattercarpetMarkerSymbolTriangleSwDot           ScattercarpetMarkerSymbol = "triangle-sw-dot"
	ScattercarpetMarkerSymbolNumber311               ScattercarpetMarkerSymbol = 311
	ScattercarpetMarkerSymbol311                     ScattercarpetMarkerSymbol = "311"
	ScattercarpetMarkerSymbolTriangleSwOpenDot       ScattercarpetMarkerSymbol = "triangle-sw-open-dot"
	ScattercarpetMarkerSymbolNumber12                ScattercarpetMarkerSymbol = 12
	ScattercarpetMarkerSymbol12                      ScattercarpetMarkerSymbol = "12"
	ScattercarpetMarkerSymbolTriangleNw              ScattercarpetMarkerSymbol = "triangle-nw"
	ScattercarpetMarkerSymbolNumber112               ScattercarpetMarkerSymbol = 112
	ScattercarpetMarkerSymbol112                     ScattercarpetMarkerSymbol = "112"
	ScattercarpetMarkerSymbolTriangleNwOpen          ScattercarpetMarkerSymbol = "triangle-nw-open"
	ScattercarpetMarkerSymbolNumber212               ScattercarpetMarkerSymbol = 212
	ScattercarpetMarkerSymbol212                     ScattercarpetMarkerSymbol = "212"
	ScattercarpetMarkerSymbolTriangleNwDot           ScattercarpetMarkerSymbol = "triangle-nw-dot"
	ScattercarpetMarkerSymbolNumber312               ScattercarpetMarkerSymbol = 312
	ScattercarpetMarkerSymbol312                     ScattercarpetMarkerSymbol = "312"
	ScattercarpetMarkerSymbolTriangleNwOpenDot       ScattercarpetMarkerSymbol = "triangle-nw-open-dot"
	ScattercarpetMarkerSymbolNumber13                ScattercarpetMarkerSymbol = 13
	ScattercarpetMarkerSymbol13                      ScattercarpetMarkerSymbol = "13"
	ScattercarpetMarkerSymbolPentagon                ScattercarpetMarkerSymbol = "pentagon"
	ScattercarpetMarkerSymbolNumber113               ScattercarpetMarkerSymbol = 113
	ScattercarpetMarkerSymbol113                     ScattercarpetMarkerSymbol = "113"
	ScattercarpetMarkerSymbolPentagonOpen            ScattercarpetMarkerSymbol = "pentagon-open"
	ScattercarpetMarkerSymbolNumber213               ScattercarpetMarkerSymbol = 213
	ScattercarpetMarkerSymbol213                     ScattercarpetMarkerSymbol = "213"
	ScattercarpetMarkerSymbolPentagonDot             ScattercarpetMarkerSymbol = "pentagon-dot"
	ScattercarpetMarkerSymbolNumber313               ScattercarpetMarkerSymbol = 313
	ScattercarpetMarkerSymbol313                     ScattercarpetMarkerSymbol = "313"
	ScattercarpetMarkerSymbolPentagonOpenDot         ScattercarpetMarkerSymbol = "pentagon-open-dot"
	ScattercarpetMarkerSymbolNumber14                ScattercarpetMarkerSymbol = 14
	ScattercarpetMarkerSymbol14                      ScattercarpetMarkerSymbol = "14"
	ScattercarpetMarkerSymbolHexagon                 ScattercarpetMarkerSymbol = "hexagon"
	ScattercarpetMarkerSymbolNumber114               ScattercarpetMarkerSymbol = 114
	ScattercarpetMarkerSymbol114                     ScattercarpetMarkerSymbol = "114"
	ScattercarpetMarkerSymbolHexagonOpen             ScattercarpetMarkerSymbol = "hexagon-open"
	ScattercarpetMarkerSymbolNumber214               ScattercarpetMarkerSymbol = 214
	ScattercarpetMarkerSymbol214                     ScattercarpetMarkerSymbol = "214"
	ScattercarpetMarkerSymbolHexagonDot              ScattercarpetMarkerSymbol = "hexagon-dot"
	ScattercarpetMarkerSymbolNumber314               ScattercarpetMarkerSymbol = 314
	ScattercarpetMarkerSymbol314                     ScattercarpetMarkerSymbol = "314"
	ScattercarpetMarkerSymbolHexagonOpenDot          ScattercarpetMarkerSymbol = "hexagon-open-dot"
	ScattercarpetMarkerSymbolNumber15                ScattercarpetMarkerSymbol = 15
	ScattercarpetMarkerSymbol15                      ScattercarpetMarkerSymbol = "15"
	ScattercarpetMarkerSymbolHexagon2                ScattercarpetMarkerSymbol = "hexagon2"
	ScattercarpetMarkerSymbolNumber115               ScattercarpetMarkerSymbol = 115
	ScattercarpetMarkerSymbol115                     ScattercarpetMarkerSymbol = "115"
	ScattercarpetMarkerSymbolHexagon2Open            ScattercarpetMarkerSymbol = "hexagon2-open"
	ScattercarpetMarkerSymbolNumber215               ScattercarpetMarkerSymbol = 215
	ScattercarpetMarkerSymbol215                     ScattercarpetMarkerSymbol = "215"
	ScattercarpetMarkerSymbolHexagon2Dot             ScattercarpetMarkerSymbol = "hexagon2-dot"
	ScattercarpetMarkerSymbolNumber315               ScattercarpetMarkerSymbol = 315
	ScattercarpetMarkerSymbol315                     ScattercarpetMarkerSymbol = "315"
	ScattercarpetMarkerSymbolHexagon2OpenDot         ScattercarpetMarkerSymbol = "hexagon2-open-dot"
	ScattercarpetMarkerSymbolNumber16                ScattercarpetMarkerSymbol = 16
	ScattercarpetMarkerSymbol16                      ScattercarpetMarkerSymbol = "16"
	ScattercarpetMarkerSymbolOctagon                 ScattercarpetMarkerSymbol = "octagon"
	ScattercarpetMarkerSymbolNumber116               ScattercarpetMarkerSymbol = 116
	ScattercarpetMarkerSymbol116                     ScattercarpetMarkerSymbol = "116"
	ScattercarpetMarkerSymbolOctagonOpen             ScattercarpetMarkerSymbol = "octagon-open"
	ScattercarpetMarkerSymbolNumber216               ScattercarpetMarkerSymbol = 216
	ScattercarpetMarkerSymbol216                     ScattercarpetMarkerSymbol = "216"
	ScattercarpetMarkerSymbolOctagonDot              ScattercarpetMarkerSymbol = "octagon-dot"
	ScattercarpetMarkerSymbolNumber316               ScattercarpetMarkerSymbol = 316
	ScattercarpetMarkerSymbol316                     ScattercarpetMarkerSymbol = "316"
	ScattercarpetMarkerSymbolOctagonOpenDot          ScattercarpetMarkerSymbol = "octagon-open-dot"
	ScattercarpetMarkerSymbolNumber17                ScattercarpetMarkerSymbol = 17
	ScattercarpetMarkerSymbol17                      ScattercarpetMarkerSymbol = "17"
	ScattercarpetMarkerSymbolStar                    ScattercarpetMarkerSymbol = "star"
	ScattercarpetMarkerSymbolNumber117               ScattercarpetMarkerSymbol = 117
	ScattercarpetMarkerSymbol117                     ScattercarpetMarkerSymbol = "117"
	ScattercarpetMarkerSymbolStarOpen                ScattercarpetMarkerSymbol = "star-open"
	ScattercarpetMarkerSymbolNumber217               ScattercarpetMarkerSymbol = 217
	ScattercarpetMarkerSymbol217                     ScattercarpetMarkerSymbol = "217"
	ScattercarpetMarkerSymbolStarDot                 ScattercarpetMarkerSymbol = "star-dot"
	ScattercarpetMarkerSymbolNumber317               ScattercarpetMarkerSymbol = 317
	ScattercarpetMarkerSymbol317                     ScattercarpetMarkerSymbol = "317"
	ScattercarpetMarkerSymbolStarOpenDot             ScattercarpetMarkerSymbol = "star-open-dot"
	ScattercarpetMarkerSymbolNumber18                ScattercarpetMarkerSymbol = 18
	ScattercarpetMarkerSymbol18                      ScattercarpetMarkerSymbol = "18"
	ScattercarpetMarkerSymbolHexagram                ScattercarpetMarkerSymbol = "hexagram"
	ScattercarpetMarkerSymbolNumber118               ScattercarpetMarkerSymbol = 118
	ScattercarpetMarkerSymbol118                     ScattercarpetMarkerSymbol = "118"
	ScattercarpetMarkerSymbolHexagramOpen            ScattercarpetMarkerSymbol = "hexagram-open"
	ScattercarpetMarkerSymbolNumber218               ScattercarpetMarkerSymbol = 218
	ScattercarpetMarkerSymbol218                     ScattercarpetMarkerSymbol = "218"
	ScattercarpetMarkerSymbolHexagramDot             ScattercarpetMarkerSymbol = "hexagram-dot"
	ScattercarpetMarkerSymbolNumber318               ScattercarpetMarkerSymbol = 318
	ScattercarpetMarkerSymbol318                     ScattercarpetMarkerSymbol = "318"
	ScattercarpetMarkerSymbolHexagramOpenDot         ScattercarpetMarkerSymbol = "hexagram-open-dot"
	ScattercarpetMarkerSymbolNumber19                ScattercarpetMarkerSymbol = 19
	ScattercarpetMarkerSymbol19                      ScattercarpetMarkerSymbol = "19"
	ScattercarpetMarkerSymbolStarTriangleUp          ScattercarpetMarkerSymbol = "star-triangle-up"
	ScattercarpetMarkerSymbolNumber119               ScattercarpetMarkerSymbol = 119
	ScattercarpetMarkerSymbol119                     ScattercarpetMarkerSymbol = "119"
	ScattercarpetMarkerSymbolStarTriangleUpOpen      ScattercarpetMarkerSymbol = "star-triangle-up-open"
	ScattercarpetMarkerSymbolNumber219               ScattercarpetMarkerSymbol = 219
	ScattercarpetMarkerSymbol219                     ScattercarpetMarkerSymbol = "219"
	ScattercarpetMarkerSymbolStarTriangleUpDot       ScattercarpetMarkerSymbol = "star-triangle-up-dot"
	ScattercarpetMarkerSymbolNumber319               ScattercarpetMarkerSymbol = 319
	ScattercarpetMarkerSymbol319                     ScattercarpetMarkerSymbol = "319"
	ScattercarpetMarkerSymbolStarTriangleUpOpenDot   ScattercarpetMarkerSymbol = "star-triangle-up-open-dot"
	ScattercarpetMarkerSymbolNumber20                ScattercarpetMarkerSymbol = 20
	ScattercarpetMarkerSymbol20                      ScattercarpetMarkerSymbol = "20"
	ScattercarpetMarkerSymbolStarTriangleDown        ScattercarpetMarkerSymbol = "star-triangle-down"
	ScattercarpetMarkerSymbolNumber120               ScattercarpetMarkerSymbol = 120
	ScattercarpetMarkerSymbol120                     ScattercarpetMarkerSymbol = "120"
	ScattercarpetMarkerSymbolStarTriangleDownOpen    ScattercarpetMarkerSymbol = "star-triangle-down-open"
	ScattercarpetMarkerSymbolNumber220               ScattercarpetMarkerSymbol = 220
	ScattercarpetMarkerSymbol220                     ScattercarpetMarkerSymbol = "220"
	ScattercarpetMarkerSymbolStarTriangleDownDot     ScattercarpetMarkerSymbol = "star-triangle-down-dot"
	ScattercarpetMarkerSymbolNumber320               ScattercarpetMarkerSymbol = 320
	ScattercarpetMarkerSymbol320                     ScattercarpetMarkerSymbol = "320"
	ScattercarpetMarkerSymbolStarTriangleDownOpenDot ScattercarpetMarkerSymbol = "star-triangle-down-open-dot"
	ScattercarpetMarkerSymbolNumber21                ScattercarpetMarkerSymbol = 21
	ScattercarpetMarkerSymbol21                      ScattercarpetMarkerSymbol = "21"
	ScattercarpetMarkerSymbolStarSquare              ScattercarpetMarkerSymbol = "star-square"
	ScattercarpetMarkerSymbolNumber121               ScattercarpetMarkerSymbol = 121
	ScattercarpetMarkerSymbol121                     ScattercarpetMarkerSymbol = "121"
	ScattercarpetMarkerSymbolStarSquareOpen          ScattercarpetMarkerSymbol = "star-square-open"
	ScattercarpetMarkerSymbolNumber221               ScattercarpetMarkerSymbol = 221
	ScattercarpetMarkerSymbol221                     ScattercarpetMarkerSymbol = "221"
	ScattercarpetMarkerSymbolStarSquareDot           ScattercarpetMarkerSymbol = "star-square-dot"
	ScattercarpetMarkerSymbolNumber321               ScattercarpetMarkerSymbol = 321
	ScattercarpetMarkerSymbol321                     ScattercarpetMarkerSymbol = "321"
	ScattercarpetMarkerSymbolStarSquareOpenDot       ScattercarpetMarkerSymbol = "star-square-open-dot"
	ScattercarpetMarkerSymbolNumber22                ScattercarpetMarkerSymbol = 22
	ScattercarpetMarkerSymbol22                      ScattercarpetMarkerSymbol = "22"
	ScattercarpetMarkerSymbolStarDiamond             ScattercarpetMarkerSymbol = "star-diamond"
	ScattercarpetMarkerSymbolNumber122               ScattercarpetMarkerSymbol = 122
	ScattercarpetMarkerSymbol122                     ScattercarpetMarkerSymbol = "122"
	ScattercarpetMarkerSymbolStarDiamondOpen         ScattercarpetMarkerSymbol = "star-diamond-open"
	ScattercarpetMarkerSymbolNumber222               ScattercarpetMarkerSymbol = 222
	ScattercarpetMarkerSymbol222                     ScattercarpetMarkerSymbol = "222"
	ScattercarpetMarkerSymbolStarDiamondDot          ScattercarpetMarkerSymbol = "star-diamond-dot"
	ScattercarpetMarkerSymbolNumber322               ScattercarpetMarkerSymbol = 322
	ScattercarpetMarkerSymbol322                     ScattercarpetMarkerSymbol = "322"
	ScattercarpetMarkerSymbolStarDiamondOpenDot      ScattercarpetMarkerSymbol = "star-diamond-open-dot"
	ScattercarpetMarkerSymbolNumber23                ScattercarpetMarkerSymbol = 23
	ScattercarpetMarkerSymbol23                      ScattercarpetMarkerSymbol = "23"
	ScattercarpetMarkerSymbolDiamondTall             ScattercarpetMarkerSymbol = "diamond-tall"
	ScattercarpetMarkerSymbolNumber123               ScattercarpetMarkerSymbol = 123
	ScattercarpetMarkerSymbol123                     ScattercarpetMarkerSymbol = "123"
	ScattercarpetMarkerSymbolDiamondTallOpen         ScattercarpetMarkerSymbol = "diamond-tall-open"
	ScattercarpetMarkerSymbolNumber223               ScattercarpetMarkerSymbol = 223
	ScattercarpetMarkerSymbol223                     ScattercarpetMarkerSymbol = "223"
	ScattercarpetMarkerSymbolDiamondTallDot          ScattercarpetMarkerSymbol = "diamond-tall-dot"
	ScattercarpetMarkerSymbolNumber323               ScattercarpetMarkerSymbol = 323
	ScattercarpetMarkerSymbol323                     ScattercarpetMarkerSymbol = "323"
	ScattercarpetMarkerSymbolDiamondTallOpenDot      ScattercarpetMarkerSymbol = "diamond-tall-open-dot"
	ScattercarpetMarkerSymbolNumber24                ScattercarpetMarkerSymbol = 24
	ScattercarpetMarkerSymbol24                      ScattercarpetMarkerSymbol = "24"
	ScattercarpetMarkerSymbolDiamondWide             ScattercarpetMarkerSymbol = "diamond-wide"
	ScattercarpetMarkerSymbolNumber124               ScattercarpetMarkerSymbol = 124
	ScattercarpetMarkerSymbol124                     ScattercarpetMarkerSymbol = "124"
	ScattercarpetMarkerSymbolDiamondWideOpen         ScattercarpetMarkerSymbol = "diamond-wide-open"
	ScattercarpetMarkerSymbolNumber224               ScattercarpetMarkerSymbol = 224
	ScattercarpetMarkerSymbol224                     ScattercarpetMarkerSymbol = "224"
	ScattercarpetMarkerSymbolDiamondWideDot          ScattercarpetMarkerSymbol = "diamond-wide-dot"
	ScattercarpetMarkerSymbolNumber324               ScattercarpetMarkerSymbol = 324
	ScattercarpetMarkerSymbol324                     ScattercarpetMarkerSymbol = "324"
	ScattercarpetMarkerSymbolDiamondWideOpenDot      ScattercarpetMarkerSymbol = "diamond-wide-open-dot"
	ScattercarpetMarkerSymbolNumber25                ScattercarpetMarkerSymbol = 25
	ScattercarpetMarkerSymbol25                      ScattercarpetMarkerSymbol = "25"
	ScattercarpetMarkerSymbolHourglass               ScattercarpetMarkerSymbol = "hourglass"
	ScattercarpetMarkerSymbolNumber125               ScattercarpetMarkerSymbol = 125
	ScattercarpetMarkerSymbol125                     ScattercarpetMarkerSymbol = "125"
	ScattercarpetMarkerSymbolHourglassOpen           ScattercarpetMarkerSymbol = "hourglass-open"
	ScattercarpetMarkerSymbolNumber26                ScattercarpetMarkerSymbol = 26
	ScattercarpetMarkerSymbol26                      ScattercarpetMarkerSymbol = "26"
	ScattercarpetMarkerSymbolBowtie                  ScattercarpetMarkerSymbol = "bowtie"
	ScattercarpetMarkerSymbolNumber126               ScattercarpetMarkerSymbol = 126
	ScattercarpetMarkerSymbol126                     ScattercarpetMarkerSymbol = "126"
	ScattercarpetMarkerSymbolBowtieOpen              ScattercarpetMarkerSymbol = "bowtie-open"
	ScattercarpetMarkerSymbolNumber27                ScattercarpetMarkerSymbol = 27
	ScattercarpetMarkerSymbol27                      ScattercarpetMarkerSymbol = "27"
	ScattercarpetMarkerSymbolCircleCross             ScattercarpetMarkerSymbol = "circle-cross"
	ScattercarpetMarkerSymbolNumber127               ScattercarpetMarkerSymbol = 127
	ScattercarpetMarkerSymbol127                     ScattercarpetMarkerSymbol = "127"
	ScattercarpetMarkerSymbolCircleCrossOpen         ScattercarpetMarkerSymbol = "circle-cross-open"
	ScattercarpetMarkerSymbolNumber28                ScattercarpetMarkerSymbol = 28
	ScattercarpetMarkerSymbol28                      ScattercarpetMarkerSymbol = "28"
	ScattercarpetMarkerSymbolCircleX                 ScattercarpetMarkerSymbol = "circle-x"
	ScattercarpetMarkerSymbolNumber128               ScattercarpetMarkerSymbol = 128
	ScattercarpetMarkerSymbol128                     ScattercarpetMarkerSymbol = "128"
	ScattercarpetMarkerSymbolCircleXOpen             ScattercarpetMarkerSymbol = "circle-x-open"
	ScattercarpetMarkerSymbolNumber29                ScattercarpetMarkerSymbol = 29
	ScattercarpetMarkerSymbol29                      ScattercarpetMarkerSymbol = "29"
	ScattercarpetMarkerSymbolSquareCross             ScattercarpetMarkerSymbol = "square-cross"
	ScattercarpetMarkerSymbolNumber129               ScattercarpetMarkerSymbol = 129
	ScattercarpetMarkerSymbol129                     ScattercarpetMarkerSymbol = "129"
	ScattercarpetMarkerSymbolSquareCrossOpen         ScattercarpetMarkerSymbol = "square-cross-open"
	ScattercarpetMarkerSymbolNumber30                ScattercarpetMarkerSymbol = 30
	ScattercarpetMarkerSymbol30                      ScattercarpetMarkerSymbol = "30"
	ScattercarpetMarkerSymbolSquareX                 ScattercarpetMarkerSymbol = "square-x"
	ScattercarpetMarkerSymbolNumber130               ScattercarpetMarkerSymbol = 130
	ScattercarpetMarkerSymbol130                     ScattercarpetMarkerSymbol = "130"
	ScattercarpetMarkerSymbolSquareXOpen             ScattercarpetMarkerSymbol = "square-x-open"
	ScattercarpetMarkerSymbolNumber31                ScattercarpetMarkerSymbol = 31
	ScattercarpetMarkerSymbol31                      ScattercarpetMarkerSymbol = "31"
	ScattercarpetMarkerSymbolDiamondCross            ScattercarpetMarkerSymbol = "diamond-cross"
	ScattercarpetMarkerSymbolNumber131               ScattercarpetMarkerSymbol = 131
	ScattercarpetMarkerSymbol131                     ScattercarpetMarkerSymbol = "131"
	ScattercarpetMarkerSymbolDiamondCrossOpen        ScattercarpetMarkerSymbol = "diamond-cross-open"
	ScattercarpetMarkerSymbolNumber32                ScattercarpetMarkerSymbol = 32
	ScattercarpetMarkerSymbol32                      ScattercarpetMarkerSymbol = "32"
	ScattercarpetMarkerSymbolDiamondX                ScattercarpetMarkerSymbol = "diamond-x"
	ScattercarpetMarkerSymbolNumber132               ScattercarpetMarkerSymbol = 132
	ScattercarpetMarkerSymbol132                     ScattercarpetMarkerSymbol = "132"
	ScattercarpetMarkerSymbolDiamondXOpen            ScattercarpetMarkerSymbol = "diamond-x-open"
	ScattercarpetMarkerSymbolNumber33                ScattercarpetMarkerSymbol = 33
	ScattercarpetMarkerSymbol33                      ScattercarpetMarkerSymbol = "33"
	ScattercarpetMarkerSymbolCrossThin               ScattercarpetMarkerSymbol = "cross-thin"
	ScattercarpetMarkerSymbolNumber133               ScattercarpetMarkerSymbol = 133
	ScattercarpetMarkerSymbol133                     ScattercarpetMarkerSymbol = "133"
	ScattercarpetMarkerSymbolCrossThinOpen           ScattercarpetMarkerSymbol = "cross-thin-open"
	ScattercarpetMarkerSymbolNumber34                ScattercarpetMarkerSymbol = 34
	ScattercarpetMarkerSymbol34                      ScattercarpetMarkerSymbol = "34"
	ScattercarpetMarkerSymbolXThin                   ScattercarpetMarkerSymbol = "x-thin"
	ScattercarpetMarkerSymbolNumber134               ScattercarpetMarkerSymbol = 134
	ScattercarpetMarkerSymbol134                     ScattercarpetMarkerSymbol = "134"
	ScattercarpetMarkerSymbolXThinOpen               ScattercarpetMarkerSymbol = "x-thin-open"
	ScattercarpetMarkerSymbolNumber35                ScattercarpetMarkerSymbol = 35
	ScattercarpetMarkerSymbol35                      ScattercarpetMarkerSymbol = "35"
	ScattercarpetMarkerSymbolAsterisk                ScattercarpetMarkerSymbol = "asterisk"
	ScattercarpetMarkerSymbolNumber135               ScattercarpetMarkerSymbol = 135
	ScattercarpetMarkerSymbol135                     ScattercarpetMarkerSymbol = "135"
	ScattercarpetMarkerSymbolAsteriskOpen            ScattercarpetMarkerSymbol = "asterisk-open"
	ScattercarpetMarkerSymbolNumber36                ScattercarpetMarkerSymbol = 36
	ScattercarpetMarkerSymbol36                      ScattercarpetMarkerSymbol = "36"
	ScattercarpetMarkerSymbolHash                    ScattercarpetMarkerSymbol = "hash"
	ScattercarpetMarkerSymbolNumber136               ScattercarpetMarkerSymbol = 136
	ScattercarpetMarkerSymbol136                     ScattercarpetMarkerSymbol = "136"
	ScattercarpetMarkerSymbolHashOpen                ScattercarpetMarkerSymbol = "hash-open"
	ScattercarpetMarkerSymbolNumber236               ScattercarpetMarkerSymbol = 236
	ScattercarpetMarkerSymbol236                     ScattercarpetMarkerSymbol = "236"
	ScattercarpetMarkerSymbolHashDot                 ScattercarpetMarkerSymbol = "hash-dot"
	ScattercarpetMarkerSymbolNumber336               ScattercarpetMarkerSymbol = 336
	ScattercarpetMarkerSymbol336                     ScattercarpetMarkerSymbol = "336"
	ScattercarpetMarkerSymbolHashOpenDot             ScattercarpetMarkerSymbol = "hash-open-dot"
	ScattercarpetMarkerSymbolNumber37                ScattercarpetMarkerSymbol = 37
	ScattercarpetMarkerSymbol37                      ScattercarpetMarkerSymbol = "37"
	ScattercarpetMarkerSymbolYUp                     ScattercarpetMarkerSymbol = "y-up"
	ScattercarpetMarkerSymbolNumber137               ScattercarpetMarkerSymbol = 137
	ScattercarpetMarkerSymbol137                     ScattercarpetMarkerSymbol = "137"
	ScattercarpetMarkerSymbolYUpOpen                 ScattercarpetMarkerSymbol = "y-up-open"
	ScattercarpetMarkerSymbolNumber38                ScattercarpetMarkerSymbol = 38
	ScattercarpetMarkerSymbol38                      ScattercarpetMarkerSymbol = "38"
	ScattercarpetMarkerSymbolYDown                   ScattercarpetMarkerSymbol = "y-down"
	ScattercarpetMarkerSymbolNumber138               ScattercarpetMarkerSymbol = 138
	ScattercarpetMarkerSymbol138                     ScattercarpetMarkerSymbol = "138"
	ScattercarpetMarkerSymbolYDownOpen               ScattercarpetMarkerSymbol = "y-down-open"
	ScattercarpetMarkerSymbolNumber39                ScattercarpetMarkerSymbol = 39
	ScattercarpetMarkerSymbol39                      ScattercarpetMarkerSymbol = "39"
	ScattercarpetMarkerSymbolYLeft                   ScattercarpetMarkerSymbol = "y-left"
	ScattercarpetMarkerSymbolNumber139               ScattercarpetMarkerSymbol = 139
	ScattercarpetMarkerSymbol139                     ScattercarpetMarkerSymbol = "139"
	ScattercarpetMarkerSymbolYLeftOpen               ScattercarpetMarkerSymbol = "y-left-open"
	ScattercarpetMarkerSymbolNumber40                ScattercarpetMarkerSymbol = 40
	ScattercarpetMarkerSymbol40                      ScattercarpetMarkerSymbol = "40"
	ScattercarpetMarkerSymbolYRight                  ScattercarpetMarkerSymbol = "y-right"
	ScattercarpetMarkerSymbolNumber140               ScattercarpetMarkerSymbol = 140
	ScattercarpetMarkerSymbol140                     ScattercarpetMarkerSymbol = "140"
	ScattercarpetMarkerSymbolYRightOpen              ScattercarpetMarkerSymbol = "y-right-open"
	ScattercarpetMarkerSymbolNumber41                ScattercarpetMarkerSymbol = 41
	ScattercarpetMarkerSymbol41                      ScattercarpetMarkerSymbol = "41"
	ScattercarpetMarkerSymbolLineEw                  ScattercarpetMarkerSymbol = "line-ew"
	ScattercarpetMarkerSymbolNumber141               ScattercarpetMarkerSymbol = 141
	ScattercarpetMarkerSymbol141                     ScattercarpetMarkerSymbol = "141"
	ScattercarpetMarkerSymbolLineEwOpen              ScattercarpetMarkerSymbol = "line-ew-open"
	ScattercarpetMarkerSymbolNumber42                ScattercarpetMarkerSymbol = 42
	ScattercarpetMarkerSymbol42                      ScattercarpetMarkerSymbol = "42"
	ScattercarpetMarkerSymbolLineNs                  ScattercarpetMarkerSymbol = "line-ns"
	ScattercarpetMarkerSymbolNumber142               ScattercarpetMarkerSymbol = 142
	ScattercarpetMarkerSymbol142                     ScattercarpetMarkerSymbol = "142"
	ScattercarpetMarkerSymbolLineNsOpen              ScattercarpetMarkerSymbol = "line-ns-open"
	ScattercarpetMarkerSymbolNumber43                ScattercarpetMarkerSymbol = 43
	ScattercarpetMarkerSymbol43                      ScattercarpetMarkerSymbol = "43"
	ScattercarpetMarkerSymbolLineNe                  ScattercarpetMarkerSymbol = "line-ne"
	ScattercarpetMarkerSymbolNumber143               ScattercarpetMarkerSymbol = 143
	ScattercarpetMarkerSymbol143                     ScattercarpetMarkerSymbol = "143"
	ScattercarpetMarkerSymbolLineNeOpen              ScattercarpetMarkerSymbol = "line-ne-open"
	ScattercarpetMarkerSymbolNumber44                ScattercarpetMarkerSymbol = 44
	ScattercarpetMarkerSymbol44                      ScattercarpetMarkerSymbol = "44"
	ScattercarpetMarkerSymbolLineNw                  ScattercarpetMarkerSymbol = "line-nw"
	ScattercarpetMarkerSymbolNumber144               ScattercarpetMarkerSymbol = 144
	ScattercarpetMarkerSymbol144                     ScattercarpetMarkerSymbol = "144"
	ScattercarpetMarkerSymbolLineNwOpen              ScattercarpetMarkerSymbol = "line-nw-open"
	ScattercarpetMarkerSymbolNumber45                ScattercarpetMarkerSymbol = 45
	ScattercarpetMarkerSymbol45                      ScattercarpetMarkerSymbol = "45"
	ScattercarpetMarkerSymbolArrowUp                 ScattercarpetMarkerSymbol = "arrow-up"
	ScattercarpetMarkerSymbolNumber145               ScattercarpetMarkerSymbol = 145
	ScattercarpetMarkerSymbol145                     ScattercarpetMarkerSymbol = "145"
	ScattercarpetMarkerSymbolArrowUpOpen             ScattercarpetMarkerSymbol = "arrow-up-open"
	ScattercarpetMarkerSymbolNumber46                ScattercarpetMarkerSymbol = 46
	ScattercarpetMarkerSymbol46                      ScattercarpetMarkerSymbol = "46"
	ScattercarpetMarkerSymbolArrowDown               ScattercarpetMarkerSymbol = "arrow-down"
	ScattercarpetMarkerSymbolNumber146               ScattercarpetMarkerSymbol = 146
	ScattercarpetMarkerSymbol146                     ScattercarpetMarkerSymbol = "146"
	ScattercarpetMarkerSymbolArrowDownOpen           ScattercarpetMarkerSymbol = "arrow-down-open"
	ScattercarpetMarkerSymbolNumber47                ScattercarpetMarkerSymbol = 47
	ScattercarpetMarkerSymbol47                      ScattercarpetMarkerSymbol = "47"
	ScattercarpetMarkerSymbolArrowLeft               ScattercarpetMarkerSymbol = "arrow-left"
	ScattercarpetMarkerSymbolNumber147               ScattercarpetMarkerSymbol = 147
	ScattercarpetMarkerSymbol147                     ScattercarpetMarkerSymbol = "147"
	ScattercarpetMarkerSymbolArrowLeftOpen           ScattercarpetMarkerSymbol = "arrow-left-open"
	ScattercarpetMarkerSymbolNumber48                ScattercarpetMarkerSymbol = 48
	ScattercarpetMarkerSymbol48                      ScattercarpetMarkerSymbol = "48"
	ScattercarpetMarkerSymbolArrowRight              ScattercarpetMarkerSymbol = "arrow-right"
	ScattercarpetMarkerSymbolNumber148               ScattercarpetMarkerSymbol = 148
	ScattercarpetMarkerSymbol148                     ScattercarpetMarkerSymbol = "148"
	ScattercarpetMarkerSymbolArrowRightOpen          ScattercarpetMarkerSymbol = "arrow-right-open"
	ScattercarpetMarkerSymbolNumber49                ScattercarpetMarkerSymbol = 49
	ScattercarpetMarkerSymbol49                      ScattercarpetMarkerSymbol = "49"
	ScattercarpetMarkerSymbolArrowBarUp              ScattercarpetMarkerSymbol = "arrow-bar-up"
	ScattercarpetMarkerSymbolNumber149               ScattercarpetMarkerSymbol = 149
	ScattercarpetMarkerSymbol149                     ScattercarpetMarkerSymbol = "149"
	ScattercarpetMarkerSymbolArrowBarUpOpen          ScattercarpetMarkerSymbol = "arrow-bar-up-open"
	ScattercarpetMarkerSymbolNumber50                ScattercarpetMarkerSymbol = 50
	ScattercarpetMarkerSymbol50                      ScattercarpetMarkerSymbol = "50"
	ScattercarpetMarkerSymbolArrowBarDown            ScattercarpetMarkerSymbol = "arrow-bar-down"
	ScattercarpetMarkerSymbolNumber150               ScattercarpetMarkerSymbol = 150
	ScattercarpetMarkerSymbol150                     ScattercarpetMarkerSymbol = "150"
	ScattercarpetMarkerSymbolArrowBarDownOpen        ScattercarpetMarkerSymbol = "arrow-bar-down-open"
	ScattercarpetMarkerSymbolNumber51                ScattercarpetMarkerSymbol = 51
	ScattercarpetMarkerSymbol51                      ScattercarpetMarkerSymbol = "51"
	ScattercarpetMarkerSymbolArrowBarLeft            ScattercarpetMarkerSymbol = "arrow-bar-left"
	ScattercarpetMarkerSymbolNumber151               ScattercarpetMarkerSymbol = 151
	ScattercarpetMarkerSymbol151                     ScattercarpetMarkerSymbol = "151"
	ScattercarpetMarkerSymbolArrowBarLeftOpen        ScattercarpetMarkerSymbol = "arrow-bar-left-open"
	ScattercarpetMarkerSymbolNumber52                ScattercarpetMarkerSymbol = 52
	ScattercarpetMarkerSymbol52                      ScattercarpetMarkerSymbol = "52"
	ScattercarpetMarkerSymbolArrowBarRight           ScattercarpetMarkerSymbol = "arrow-bar-right"
	ScattercarpetMarkerSymbolNumber152               ScattercarpetMarkerSymbol = 152
	ScattercarpetMarkerSymbol152                     ScattercarpetMarkerSymbol = "152"
	ScattercarpetMarkerSymbolArrowBarRightOpen       ScattercarpetMarkerSymbol = "arrow-bar-right-open"
)

// ScattercarpetTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScattercarpetTextposition string

const (
	ScattercarpetTextpositionTopLeft      ScattercarpetTextposition = "top left"
	ScattercarpetTextpositionTopCenter    ScattercarpetTextposition = "top center"
	ScattercarpetTextpositionTopRight     ScattercarpetTextposition = "top right"
	ScattercarpetTextpositionMiddleLeft   ScattercarpetTextposition = "middle left"
	ScattercarpetTextpositionMiddleCenter ScattercarpetTextposition = "middle center"
	ScattercarpetTextpositionMiddleRight  ScattercarpetTextposition = "middle right"
	ScattercarpetTextpositionBottomLeft   ScattercarpetTextposition = "bottom left"
	ScattercarpetTextpositionBottomCenter ScattercarpetTextposition = "bottom center"
	ScattercarpetTextpositionBottomRight  ScattercarpetTextposition = "bottom right"
)

// ScattercarpetVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScattercarpetVisible interface{}

var (
	ScattercarpetVisibleTrue       ScattercarpetVisible = true
	ScattercarpetVisibleFalse      ScattercarpetVisible = false
	ScattercarpetVisibleLegendonly ScattercarpetVisible = "legendonly"
)

// ScattercarpetHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type ScattercarpetHoverinfo string

const (
	// Flags
	ScattercarpetHoverinfoA    ScattercarpetHoverinfo = "a"
	ScattercarpetHoverinfoB    ScattercarpetHoverinfo = "b"
	ScattercarpetHoverinfoText ScattercarpetHoverinfo = "text"
	ScattercarpetHoverinfoName ScattercarpetHoverinfo = "name"

	// Extra
	ScattercarpetHoverinfoAll  ScattercarpetHoverinfo = "all"
	ScattercarpetHoverinfoNone ScattercarpetHoverinfo = "none"
	ScattercarpetHoverinfoSkip ScattercarpetHoverinfo = "skip"
)

// ScattercarpetHoveron Do the hover effects highlight individual points (markers or line points) or do they highlight filled regions? If the fill is *toself* or *tonext* and there are no markers or text, then the default is *fills*, otherwise it is *points*.
type ScattercarpetHoveron string

const (
	// Flags
	ScattercarpetHoveronPoints ScattercarpetHoveron = "points"
	ScattercarpetHoveronFills  ScattercarpetHoveron = "fills"

	// Extra

)

// ScattercarpetMode Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
type ScattercarpetMode string

const (
	// Flags
	ScattercarpetModeLines   ScattercarpetMode = "lines"
	ScattercarpetModeMarkers ScattercarpetMode = "markers"
	ScattercarpetModeText    ScattercarpetMode = "text"

	// Extra
	ScattercarpetModeNone ScattercarpetMode = "none"
)
