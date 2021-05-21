package grob

var TraceTypeScattergl TraceType = "scattergl"

func (trace *Scattergl) GetType() TraceType {
	return TraceTypeScattergl
}

// Scattergl The data visualized as scatter point or lines is set in `x` and `y` using the WebGL plotting engine. Bubble charts are achieved by setting `marker.size` and/or `marker.color` to a numerical arrays.
type Scattergl struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

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

	// Dx
	// arrayOK: false
	// type: number
	// Sets the x coordinate step. See `x0` for more info.
	Dx float64 `json:"dx,omitempty"`

	// Dy
	// arrayOK: false
	// type: number
	// Sets the y coordinate step. See `y0` for more info.
	Dy float64 `json:"dy,omitempty"`

	// ErrorX
	// role: Object
	ErrorX *ScatterglErrorX `json:"error_x,omitempty"`

	// ErrorY
	// role: Object
	ErrorY *ScatterglErrorY `json:"error_y,omitempty"`

	// Fill
	// default: none
	// type: enumerated
	// Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
	Fill ScatterglFill `json:"fill,omitempty"`

	// Fillcolor
	// arrayOK: false
	// type: color
	// Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor Color `json:"fillcolor,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ScatterglHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *ScatterglHoverlabel `json:"hoverlabel,omitempty"`

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
	// Sets hover text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
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
	Line *ScatterglLine `json:"line,omitempty"`

	// Marker
	// role: Object
	Marker *ScatterglMarker `json:"marker,omitempty"`

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
	// default: %!s(<nil>)
	// type: flaglist
	// Determines the drawing mode for this scatter trace.
	Mode ScatterglMode `json:"mode,omitempty"`

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
	Selected *ScatterglSelected `json:"selected,omitempty"`

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
	Stream *ScatterglStream `json:"stream,omitempty"`

	// Text
	// arrayOK: true
	// type: string
	// Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text String `json:"text,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScatterglTextfont `json:"textfont,omitempty"`

	// Textposition
	// default: middle center
	// type: enumerated
	// Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition ScatterglTextposition `json:"textposition,omitempty"`

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
	// Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.
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
	Unselected *ScatterglUnselected `json:"unselected,omitempty"`

	// Visible
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ScatterglVisible `json:"visible,omitempty"`

	// X
	// arrayOK: false
	// type: data_array
	// Sets the x coordinates.
	X interface{} `json:"x,omitempty"`

	// X0
	// arrayOK: false
	// type: any
	// Alternate to `x`. Builds a linear space of x coordinates. Use with `dx` where `x0` is the starting coordinate and `dx` the step.
	X0 interface{} `json:"x0,omitempty"`

	// Xaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xcalendar
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `x` date data.
	Xcalendar ScatterglXcalendar `json:"xcalendar,omitempty"`

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
	// default: middle
	// type: enumerated
	// Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
	Xperiodalignment ScatterglXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y
	// arrayOK: false
	// type: data_array
	// Sets the y coordinates.
	Y interface{} `json:"y,omitempty"`

	// Y0
	// arrayOK: false
	// type: any
	// Alternate to `y`. Builds a linear space of y coordinates. Use with `dy` where `y0` is the starting coordinate and `dy` the step.
	Y0 interface{} `json:"y0,omitempty"`

	// Yaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ycalendar
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `y` date data.
	Ycalendar ScatterglYcalendar `json:"ycalendar,omitempty"`

	// Yperiod
	// arrayOK: false
	// type: any
	// Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the y axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Yperiod interface{} `json:"yperiod,omitempty"`

	// Yperiod0
	// arrayOK: false
	// type: any
	// Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the y0 axis. When `y0period` is round number of weeks, the `y0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Yperiod0 interface{} `json:"yperiod0,omitempty"`

	// Yperiodalignment
	// default: middle
	// type: enumerated
	// Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
	Yperiodalignment ScatterglYperiodalignment `json:"yperiodalignment,omitempty"`

	// Ysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

// ScatterglErrorX
type ScatterglErrorX struct {

	// Array
	// arrayOK: false
	// type: data_array
	// Sets the data corresponding the length of each error bar. Values are plotted relative to the underlying data.
	Array interface{} `json:"array,omitempty"`

	// Arrayminus
	// arrayOK: false
	// type: data_array
	// Sets the data corresponding the length of each error bar in the bottom (left) direction for vertical (horizontal) bars Values are plotted relative to the underlying data.
	Arrayminus interface{} `json:"arrayminus,omitempty"`

	// Arrayminussrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  arrayminus .
	Arrayminussrc String `json:"arrayminussrc,omitempty"`

	// Arraysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  array .
	Arraysrc String `json:"arraysrc,omitempty"`

	// Color
	// arrayOK: false
	// type: color
	// Sets the stoke color of the error bars.
	Color Color `json:"color,omitempty"`

	// CopyYstyle
	// arrayOK: false
	// type: boolean
	//
	CopyYstyle Bool `json:"copy_ystyle,omitempty"`

	// Symmetric
	// arrayOK: false
	// type: boolean
	// Determines whether or not the error bars have the same length in both direction (top/bottom for vertical bars, left/right for horizontal bars.
	Symmetric Bool `json:"symmetric,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness (in px) of the error bars.
	Thickness float64 `json:"thickness,omitempty"`

	// Traceref
	// arrayOK: false
	// type: integer
	//
	Traceref int64 `json:"traceref,omitempty"`

	// Tracerefminus
	// arrayOK: false
	// type: integer
	//
	Tracerefminus int64 `json:"tracerefminus,omitempty"`

	// Type
	// default: %!s(<nil>)
	// type: enumerated
	// Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the square of the underlying data. If *data*, the bar lengths are set with data set `array`.
	Type ScatterglErrorXType `json:"type,omitempty"`

	// Value
	// arrayOK: false
	// type: number
	// Sets the value of either the percentage (if `type` is set to *percent*) or the constant (if `type` is set to *constant*) corresponding to the lengths of the error bars.
	Value float64 `json:"value,omitempty"`

	// Valueminus
	// arrayOK: false
	// type: number
	// Sets the value of either the percentage (if `type` is set to *percent*) or the constant (if `type` is set to *constant*) corresponding to the lengths of the error bars in the bottom (left) direction for vertical (horizontal) bars
	Valueminus float64 `json:"valueminus,omitempty"`

	// Visible
	// arrayOK: false
	// type: boolean
	// Determines whether or not this set of error bars is visible.
	Visible Bool `json:"visible,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the cross-bar at both ends of the error bars.
	Width float64 `json:"width,omitempty"`
}

// ScatterglErrorY
type ScatterglErrorY struct {

	// Array
	// arrayOK: false
	// type: data_array
	// Sets the data corresponding the length of each error bar. Values are plotted relative to the underlying data.
	Array interface{} `json:"array,omitempty"`

	// Arrayminus
	// arrayOK: false
	// type: data_array
	// Sets the data corresponding the length of each error bar in the bottom (left) direction for vertical (horizontal) bars Values are plotted relative to the underlying data.
	Arrayminus interface{} `json:"arrayminus,omitempty"`

	// Arrayminussrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  arrayminus .
	Arrayminussrc String `json:"arrayminussrc,omitempty"`

	// Arraysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  array .
	Arraysrc String `json:"arraysrc,omitempty"`

	// Color
	// arrayOK: false
	// type: color
	// Sets the stoke color of the error bars.
	Color Color `json:"color,omitempty"`

	// Symmetric
	// arrayOK: false
	// type: boolean
	// Determines whether or not the error bars have the same length in both direction (top/bottom for vertical bars, left/right for horizontal bars.
	Symmetric Bool `json:"symmetric,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness (in px) of the error bars.
	Thickness float64 `json:"thickness,omitempty"`

	// Traceref
	// arrayOK: false
	// type: integer
	//
	Traceref int64 `json:"traceref,omitempty"`

	// Tracerefminus
	// arrayOK: false
	// type: integer
	//
	Tracerefminus int64 `json:"tracerefminus,omitempty"`

	// Type
	// default: %!s(<nil>)
	// type: enumerated
	// Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the square of the underlying data. If *data*, the bar lengths are set with data set `array`.
	Type ScatterglErrorYType `json:"type,omitempty"`

	// Value
	// arrayOK: false
	// type: number
	// Sets the value of either the percentage (if `type` is set to *percent*) or the constant (if `type` is set to *constant*) corresponding to the lengths of the error bars.
	Value float64 `json:"value,omitempty"`

	// Valueminus
	// arrayOK: false
	// type: number
	// Sets the value of either the percentage (if `type` is set to *percent*) or the constant (if `type` is set to *constant*) corresponding to the lengths of the error bars in the bottom (left) direction for vertical (horizontal) bars
	Valueminus float64 `json:"valueminus,omitempty"`

	// Visible
	// arrayOK: false
	// type: boolean
	// Determines whether or not this set of error bars is visible.
	Visible Bool `json:"visible,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the cross-bar at both ends of the error bars.
	Width float64 `json:"width,omitempty"`
}

// ScatterglHoverlabelFont Sets the font used in hover labels.
type ScatterglHoverlabelFont struct {

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

// ScatterglHoverlabel
type ScatterglHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align ScatterglHoverlabelAlign `json:"align,omitempty"`

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
	Font *ScatterglHoverlabelFont `json:"font,omitempty"`

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

// ScatterglLine
type ScatterglLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the line color.
	Color Color `json:"color,omitempty"`

	// Dash
	// default: solid
	// type: enumerated
	// Sets the style of the lines.
	Dash ScatterglLineDash `json:"dash,omitempty"`

	// Shape
	// default: linear
	// type: enumerated
	// Determines the line shape. The values correspond to step-wise line shapes.
	Shape ScatterglLineShape `json:"shape,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the line width (in px).
	Width float64 `json:"width,omitempty"`
}

// ScatterglMarkerColorbarTickfont Sets the color bar's tick label font
type ScatterglMarkerColorbarTickfont struct {

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

// ScatterglMarkerColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type ScatterglMarkerColorbarTitleFont struct {

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

// ScatterglMarkerColorbarTitle
type ScatterglMarkerColorbarTitle struct {

	// Font
	// role: Object
	Font *ScatterglMarkerColorbarTitleFont `json:"font,omitempty"`

	// Side
	// default: top
	// type: enumerated
	// Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
	Side ScatterglMarkerColorbarTitleSide `json:"side,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// ScatterglMarkerColorbar
type ScatterglMarkerColorbar struct {

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
	Exponentformat ScatterglMarkerColorbarExponentformat `json:"exponentformat,omitempty"`

	// Len
	// arrayOK: false
	// type: number
	// Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends.
	Len float64 `json:"len,omitempty"`

	// Lenmode
	// default: fraction
	// type: enumerated
	// Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
	Lenmode ScatterglMarkerColorbarLenmode `json:"lenmode,omitempty"`

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
	Showexponent ScatterglMarkerColorbarShowexponent `json:"showexponent,omitempty"`

	// Showticklabels
	// arrayOK: false
	// type: boolean
	// Determines whether or not the tick labels are drawn.
	Showticklabels Bool `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix ScatterglMarkerColorbarShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix ScatterglMarkerColorbarShowticksuffix `json:"showticksuffix,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels.
	Thickness float64 `json:"thickness,omitempty"`

	// Thicknessmode
	// default: pixels
	// type: enumerated
	// Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
	Thicknessmode ScatterglMarkerColorbarThicknessmode `json:"thicknessmode,omitempty"`

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
	Tickfont *ScatterglMarkerColorbarTickfont `json:"tickfont,omitempty"`

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
	Ticklabelposition ScatterglMarkerColorbarTicklabelposition `json:"ticklabelposition,omitempty"`

	// Ticklen
	// arrayOK: false
	// type: number
	// Sets the tick length (in px).
	Ticklen float64 `json:"ticklen,omitempty"`

	// Tickmode
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
	Tickmode ScatterglMarkerColorbarTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

	// Ticks
	// default:
	// type: enumerated
	// Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
	Ticks ScatterglMarkerColorbarTicks `json:"ticks,omitempty"`

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
	Title *ScatterglMarkerColorbarTitle `json:"title,omitempty"`

	// X
	// arrayOK: false
	// type: number
	// Sets the x position of the color bar (in plot fraction).
	X float64 `json:"x,omitempty"`

	// Xanchor
	// default: left
	// type: enumerated
	// Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
	Xanchor ScatterglMarkerColorbarXanchor `json:"xanchor,omitempty"`

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
	Yanchor ScatterglMarkerColorbarYanchor `json:"yanchor,omitempty"`

	// Ypad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the y direction.
	Ypad float64 `json:"ypad,omitempty"`
}

// ScatterglMarkerLine
type ScatterglMarkerLine struct {

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

// ScatterglMarker
type ScatterglMarker struct {

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
	Colorbar *ScatterglMarkerColorbar `json:"colorbar,omitempty"`

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

	// Line
	// role: Object
	Line *ScatterglMarkerLine `json:"line,omitempty"`

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
	Sizemode ScatterglMarkerSizemode `json:"sizemode,omitempty"`

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
	Symbol ScatterglMarkerSymbol `json:"symbol,omitempty"`

	// Symbolsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  symbol .
	Symbolsrc String `json:"symbolsrc,omitempty"`
}

// ScatterglSelectedMarker
type ScatterglSelectedMarker struct {

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

// ScatterglSelectedTextfont
type ScatterglSelectedTextfont struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the text font color of selected points.
	Color Color `json:"color,omitempty"`
}

// ScatterglSelected
type ScatterglSelected struct {

	// Marker
	// role: Object
	Marker *ScatterglSelectedMarker `json:"marker,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScatterglSelectedTextfont `json:"textfont,omitempty"`
}

// ScatterglStream
type ScatterglStream struct {

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

// ScatterglTextfont Sets the text font.
type ScatterglTextfont struct {

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

// ScatterglUnselectedMarker
type ScatterglUnselectedMarker struct {

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

// ScatterglUnselectedTextfont
type ScatterglUnselectedTextfont struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the text font color of unselected points, applied only when a selection exists.
	Color Color `json:"color,omitempty"`
}

// ScatterglUnselected
type ScatterglUnselected struct {

	// Marker
	// role: Object
	Marker *ScatterglUnselectedMarker `json:"marker,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScatterglUnselectedTextfont `json:"textfont,omitempty"`
}

// ScatterglErrorXType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the square of the underlying data. If *data*, the bar lengths are set with data set `array`.
type ScatterglErrorXType string

const (
	ScatterglErrorXTypePercent  ScatterglErrorXType = "percent"
	ScatterglErrorXTypeConstant ScatterglErrorXType = "constant"
	ScatterglErrorXTypeSqrt     ScatterglErrorXType = "sqrt"
	ScatterglErrorXTypeData     ScatterglErrorXType = "data"
)

// ScatterglErrorYType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the square of the underlying data. If *data*, the bar lengths are set with data set `array`.
type ScatterglErrorYType string

const (
	ScatterglErrorYTypePercent  ScatterglErrorYType = "percent"
	ScatterglErrorYTypeConstant ScatterglErrorYType = "constant"
	ScatterglErrorYTypeSqrt     ScatterglErrorYType = "sqrt"
	ScatterglErrorYTypeData     ScatterglErrorYType = "data"
)

// ScatterglFill Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
type ScatterglFill string

const (
	ScatterglFillNone    ScatterglFill = "none"
	ScatterglFillTozeroy ScatterglFill = "tozeroy"
	ScatterglFillTozerox ScatterglFill = "tozerox"
	ScatterglFillTonexty ScatterglFill = "tonexty"
	ScatterglFillTonextx ScatterglFill = "tonextx"
	ScatterglFillToself  ScatterglFill = "toself"
	ScatterglFillTonext  ScatterglFill = "tonext"
)

// ScatterglHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterglHoverlabelAlign string

const (
	ScatterglHoverlabelAlignLeft  ScatterglHoverlabelAlign = "left"
	ScatterglHoverlabelAlignRight ScatterglHoverlabelAlign = "right"
	ScatterglHoverlabelAlignAuto  ScatterglHoverlabelAlign = "auto"
)

// ScatterglLineDash Sets the style of the lines.
type ScatterglLineDash string

const (
	ScatterglLineDashSolid       ScatterglLineDash = "solid"
	ScatterglLineDashDot         ScatterglLineDash = "dot"
	ScatterglLineDashDash        ScatterglLineDash = "dash"
	ScatterglLineDashLongdash    ScatterglLineDash = "longdash"
	ScatterglLineDashDashdot     ScatterglLineDash = "dashdot"
	ScatterglLineDashLongdashdot ScatterglLineDash = "longdashdot"
)

// ScatterglLineShape Determines the line shape. The values correspond to step-wise line shapes.
type ScatterglLineShape string

const (
	ScatterglLineShapeLinear ScatterglLineShape = "linear"
	ScatterglLineShapeHv     ScatterglLineShape = "hv"
	ScatterglLineShapeVh     ScatterglLineShape = "vh"
	ScatterglLineShapeHvh    ScatterglLineShape = "hvh"
	ScatterglLineShapeVhv    ScatterglLineShape = "vhv"
)

// ScatterglMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterglMarkerColorbarExponentformat string

const (
	ScatterglMarkerColorbarExponentformatNone  ScatterglMarkerColorbarExponentformat = "none"
	ScatterglMarkerColorbarExponentformatE1    ScatterglMarkerColorbarExponentformat = "e"
	ScatterglMarkerColorbarExponentformatE2    ScatterglMarkerColorbarExponentformat = "E"
	ScatterglMarkerColorbarExponentformatPower ScatterglMarkerColorbarExponentformat = "power"
	ScatterglMarkerColorbarExponentformatSi    ScatterglMarkerColorbarExponentformat = "SI"
	ScatterglMarkerColorbarExponentformatB     ScatterglMarkerColorbarExponentformat = "B"
)

// ScatterglMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterglMarkerColorbarLenmode string

const (
	ScatterglMarkerColorbarLenmodeFraction ScatterglMarkerColorbarLenmode = "fraction"
	ScatterglMarkerColorbarLenmodePixels   ScatterglMarkerColorbarLenmode = "pixels"
)

// ScatterglMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterglMarkerColorbarShowexponent string

const (
	ScatterglMarkerColorbarShowexponentAll   ScatterglMarkerColorbarShowexponent = "all"
	ScatterglMarkerColorbarShowexponentFirst ScatterglMarkerColorbarShowexponent = "first"
	ScatterglMarkerColorbarShowexponentLast  ScatterglMarkerColorbarShowexponent = "last"
	ScatterglMarkerColorbarShowexponentNone  ScatterglMarkerColorbarShowexponent = "none"
)

// ScatterglMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterglMarkerColorbarShowtickprefix string

const (
	ScatterglMarkerColorbarShowtickprefixAll   ScatterglMarkerColorbarShowtickprefix = "all"
	ScatterglMarkerColorbarShowtickprefixFirst ScatterglMarkerColorbarShowtickprefix = "first"
	ScatterglMarkerColorbarShowtickprefixLast  ScatterglMarkerColorbarShowtickprefix = "last"
	ScatterglMarkerColorbarShowtickprefixNone  ScatterglMarkerColorbarShowtickprefix = "none"
)

// ScatterglMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterglMarkerColorbarShowticksuffix string

const (
	ScatterglMarkerColorbarShowticksuffixAll   ScatterglMarkerColorbarShowticksuffix = "all"
	ScatterglMarkerColorbarShowticksuffixFirst ScatterglMarkerColorbarShowticksuffix = "first"
	ScatterglMarkerColorbarShowticksuffixLast  ScatterglMarkerColorbarShowticksuffix = "last"
	ScatterglMarkerColorbarShowticksuffixNone  ScatterglMarkerColorbarShowticksuffix = "none"
)

// ScatterglMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterglMarkerColorbarThicknessmode string

const (
	ScatterglMarkerColorbarThicknessmodeFraction ScatterglMarkerColorbarThicknessmode = "fraction"
	ScatterglMarkerColorbarThicknessmodePixels   ScatterglMarkerColorbarThicknessmode = "pixels"
)

// ScatterglMarkerColorbarTicklabelposition Determines where tick labels are drawn.
type ScatterglMarkerColorbarTicklabelposition string

const (
	ScatterglMarkerColorbarTicklabelpositionOutside       ScatterglMarkerColorbarTicklabelposition = "outside"
	ScatterglMarkerColorbarTicklabelpositionInside        ScatterglMarkerColorbarTicklabelposition = "inside"
	ScatterglMarkerColorbarTicklabelpositionOutsideTop    ScatterglMarkerColorbarTicklabelposition = "outside top"
	ScatterglMarkerColorbarTicklabelpositionInsideTop     ScatterglMarkerColorbarTicklabelposition = "inside top"
	ScatterglMarkerColorbarTicklabelpositionOutsideBottom ScatterglMarkerColorbarTicklabelposition = "outside bottom"
	ScatterglMarkerColorbarTicklabelpositionInsideBottom  ScatterglMarkerColorbarTicklabelposition = "inside bottom"
)

// ScatterglMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterglMarkerColorbarTickmode string

const (
	ScatterglMarkerColorbarTickmodeAuto   ScatterglMarkerColorbarTickmode = "auto"
	ScatterglMarkerColorbarTickmodeLinear ScatterglMarkerColorbarTickmode = "linear"
	ScatterglMarkerColorbarTickmodeArray  ScatterglMarkerColorbarTickmode = "array"
)

// ScatterglMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterglMarkerColorbarTicks string

const (
	ScatterglMarkerColorbarTicksOutside ScatterglMarkerColorbarTicks = "outside"
	ScatterglMarkerColorbarTicksInside  ScatterglMarkerColorbarTicks = "inside"
	ScatterglMarkerColorbarTicksEmpty   ScatterglMarkerColorbarTicks = ""
)

// ScatterglMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterglMarkerColorbarTitleSide string

const (
	ScatterglMarkerColorbarTitleSideRight  ScatterglMarkerColorbarTitleSide = "right"
	ScatterglMarkerColorbarTitleSideTop    ScatterglMarkerColorbarTitleSide = "top"
	ScatterglMarkerColorbarTitleSideBottom ScatterglMarkerColorbarTitleSide = "bottom"
)

// ScatterglMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterglMarkerColorbarXanchor string

const (
	ScatterglMarkerColorbarXanchorLeft   ScatterglMarkerColorbarXanchor = "left"
	ScatterglMarkerColorbarXanchorCenter ScatterglMarkerColorbarXanchor = "center"
	ScatterglMarkerColorbarXanchorRight  ScatterglMarkerColorbarXanchor = "right"
)

// ScatterglMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterglMarkerColorbarYanchor string

const (
	ScatterglMarkerColorbarYanchorTop    ScatterglMarkerColorbarYanchor = "top"
	ScatterglMarkerColorbarYanchorMiddle ScatterglMarkerColorbarYanchor = "middle"
	ScatterglMarkerColorbarYanchorBottom ScatterglMarkerColorbarYanchor = "bottom"
)

// ScatterglMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterglMarkerSizemode string

const (
	ScatterglMarkerSizemodeDiameter ScatterglMarkerSizemode = "diameter"
	ScatterglMarkerSizemodeArea     ScatterglMarkerSizemode = "area"
)

// ScatterglMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterglMarkerSymbol interface{}

var (
	ScatterglMarkerSymbolNumber0                 ScatterglMarkerSymbol = 0
	ScatterglMarkerSymbol0                       ScatterglMarkerSymbol = "0"
	ScatterglMarkerSymbolCircle                  ScatterglMarkerSymbol = "circle"
	ScatterglMarkerSymbolNumber100               ScatterglMarkerSymbol = 100
	ScatterglMarkerSymbol100                     ScatterglMarkerSymbol = "100"
	ScatterglMarkerSymbolCircleOpen              ScatterglMarkerSymbol = "circle-open"
	ScatterglMarkerSymbolNumber200               ScatterglMarkerSymbol = 200
	ScatterglMarkerSymbol200                     ScatterglMarkerSymbol = "200"
	ScatterglMarkerSymbolCircleDot               ScatterglMarkerSymbol = "circle-dot"
	ScatterglMarkerSymbolNumber300               ScatterglMarkerSymbol = 300
	ScatterglMarkerSymbol300                     ScatterglMarkerSymbol = "300"
	ScatterglMarkerSymbolCircleOpenDot           ScatterglMarkerSymbol = "circle-open-dot"
	ScatterglMarkerSymbolNumber1                 ScatterglMarkerSymbol = 1
	ScatterglMarkerSymbol1                       ScatterglMarkerSymbol = "1"
	ScatterglMarkerSymbolSquare                  ScatterglMarkerSymbol = "square"
	ScatterglMarkerSymbolNumber101               ScatterglMarkerSymbol = 101
	ScatterglMarkerSymbol101                     ScatterglMarkerSymbol = "101"
	ScatterglMarkerSymbolSquareOpen              ScatterglMarkerSymbol = "square-open"
	ScatterglMarkerSymbolNumber201               ScatterglMarkerSymbol = 201
	ScatterglMarkerSymbol201                     ScatterglMarkerSymbol = "201"
	ScatterglMarkerSymbolSquareDot               ScatterglMarkerSymbol = "square-dot"
	ScatterglMarkerSymbolNumber301               ScatterglMarkerSymbol = 301
	ScatterglMarkerSymbol301                     ScatterglMarkerSymbol = "301"
	ScatterglMarkerSymbolSquareOpenDot           ScatterglMarkerSymbol = "square-open-dot"
	ScatterglMarkerSymbolNumber2                 ScatterglMarkerSymbol = 2
	ScatterglMarkerSymbol2                       ScatterglMarkerSymbol = "2"
	ScatterglMarkerSymbolDiamond                 ScatterglMarkerSymbol = "diamond"
	ScatterglMarkerSymbolNumber102               ScatterglMarkerSymbol = 102
	ScatterglMarkerSymbol102                     ScatterglMarkerSymbol = "102"
	ScatterglMarkerSymbolDiamondOpen             ScatterglMarkerSymbol = "diamond-open"
	ScatterglMarkerSymbolNumber202               ScatterglMarkerSymbol = 202
	ScatterglMarkerSymbol202                     ScatterglMarkerSymbol = "202"
	ScatterglMarkerSymbolDiamondDot              ScatterglMarkerSymbol = "diamond-dot"
	ScatterglMarkerSymbolNumber302               ScatterglMarkerSymbol = 302
	ScatterglMarkerSymbol302                     ScatterglMarkerSymbol = "302"
	ScatterglMarkerSymbolDiamondOpenDot          ScatterglMarkerSymbol = "diamond-open-dot"
	ScatterglMarkerSymbolNumber3                 ScatterglMarkerSymbol = 3
	ScatterglMarkerSymbol3                       ScatterglMarkerSymbol = "3"
	ScatterglMarkerSymbolCross                   ScatterglMarkerSymbol = "cross"
	ScatterglMarkerSymbolNumber103               ScatterglMarkerSymbol = 103
	ScatterglMarkerSymbol103                     ScatterglMarkerSymbol = "103"
	ScatterglMarkerSymbolCrossOpen               ScatterglMarkerSymbol = "cross-open"
	ScatterglMarkerSymbolNumber203               ScatterglMarkerSymbol = 203
	ScatterglMarkerSymbol203                     ScatterglMarkerSymbol = "203"
	ScatterglMarkerSymbolCrossDot                ScatterglMarkerSymbol = "cross-dot"
	ScatterglMarkerSymbolNumber303               ScatterglMarkerSymbol = 303
	ScatterglMarkerSymbol303                     ScatterglMarkerSymbol = "303"
	ScatterglMarkerSymbolCrossOpenDot            ScatterglMarkerSymbol = "cross-open-dot"
	ScatterglMarkerSymbolNumber4                 ScatterglMarkerSymbol = 4
	ScatterglMarkerSymbol4                       ScatterglMarkerSymbol = "4"
	ScatterglMarkerSymbolX                       ScatterglMarkerSymbol = "x"
	ScatterglMarkerSymbolNumber104               ScatterglMarkerSymbol = 104
	ScatterglMarkerSymbol104                     ScatterglMarkerSymbol = "104"
	ScatterglMarkerSymbolXOpen                   ScatterglMarkerSymbol = "x-open"
	ScatterglMarkerSymbolNumber204               ScatterglMarkerSymbol = 204
	ScatterglMarkerSymbol204                     ScatterglMarkerSymbol = "204"
	ScatterglMarkerSymbolXDot                    ScatterglMarkerSymbol = "x-dot"
	ScatterglMarkerSymbolNumber304               ScatterglMarkerSymbol = 304
	ScatterglMarkerSymbol304                     ScatterglMarkerSymbol = "304"
	ScatterglMarkerSymbolXOpenDot                ScatterglMarkerSymbol = "x-open-dot"
	ScatterglMarkerSymbolNumber5                 ScatterglMarkerSymbol = 5
	ScatterglMarkerSymbol5                       ScatterglMarkerSymbol = "5"
	ScatterglMarkerSymbolTriangleUp              ScatterglMarkerSymbol = "triangle-up"
	ScatterglMarkerSymbolNumber105               ScatterglMarkerSymbol = 105
	ScatterglMarkerSymbol105                     ScatterglMarkerSymbol = "105"
	ScatterglMarkerSymbolTriangleUpOpen          ScatterglMarkerSymbol = "triangle-up-open"
	ScatterglMarkerSymbolNumber205               ScatterglMarkerSymbol = 205
	ScatterglMarkerSymbol205                     ScatterglMarkerSymbol = "205"
	ScatterglMarkerSymbolTriangleUpDot           ScatterglMarkerSymbol = "triangle-up-dot"
	ScatterglMarkerSymbolNumber305               ScatterglMarkerSymbol = 305
	ScatterglMarkerSymbol305                     ScatterglMarkerSymbol = "305"
	ScatterglMarkerSymbolTriangleUpOpenDot       ScatterglMarkerSymbol = "triangle-up-open-dot"
	ScatterglMarkerSymbolNumber6                 ScatterglMarkerSymbol = 6
	ScatterglMarkerSymbol6                       ScatterglMarkerSymbol = "6"
	ScatterglMarkerSymbolTriangleDown            ScatterglMarkerSymbol = "triangle-down"
	ScatterglMarkerSymbolNumber106               ScatterglMarkerSymbol = 106
	ScatterglMarkerSymbol106                     ScatterglMarkerSymbol = "106"
	ScatterglMarkerSymbolTriangleDownOpen        ScatterglMarkerSymbol = "triangle-down-open"
	ScatterglMarkerSymbolNumber206               ScatterglMarkerSymbol = 206
	ScatterglMarkerSymbol206                     ScatterglMarkerSymbol = "206"
	ScatterglMarkerSymbolTriangleDownDot         ScatterglMarkerSymbol = "triangle-down-dot"
	ScatterglMarkerSymbolNumber306               ScatterglMarkerSymbol = 306
	ScatterglMarkerSymbol306                     ScatterglMarkerSymbol = "306"
	ScatterglMarkerSymbolTriangleDownOpenDot     ScatterglMarkerSymbol = "triangle-down-open-dot"
	ScatterglMarkerSymbolNumber7                 ScatterglMarkerSymbol = 7
	ScatterglMarkerSymbol7                       ScatterglMarkerSymbol = "7"
	ScatterglMarkerSymbolTriangleLeft            ScatterglMarkerSymbol = "triangle-left"
	ScatterglMarkerSymbolNumber107               ScatterglMarkerSymbol = 107
	ScatterglMarkerSymbol107                     ScatterglMarkerSymbol = "107"
	ScatterglMarkerSymbolTriangleLeftOpen        ScatterglMarkerSymbol = "triangle-left-open"
	ScatterglMarkerSymbolNumber207               ScatterglMarkerSymbol = 207
	ScatterglMarkerSymbol207                     ScatterglMarkerSymbol = "207"
	ScatterglMarkerSymbolTriangleLeftDot         ScatterglMarkerSymbol = "triangle-left-dot"
	ScatterglMarkerSymbolNumber307               ScatterglMarkerSymbol = 307
	ScatterglMarkerSymbol307                     ScatterglMarkerSymbol = "307"
	ScatterglMarkerSymbolTriangleLeftOpenDot     ScatterglMarkerSymbol = "triangle-left-open-dot"
	ScatterglMarkerSymbolNumber8                 ScatterglMarkerSymbol = 8
	ScatterglMarkerSymbol8                       ScatterglMarkerSymbol = "8"
	ScatterglMarkerSymbolTriangleRight           ScatterglMarkerSymbol = "triangle-right"
	ScatterglMarkerSymbolNumber108               ScatterglMarkerSymbol = 108
	ScatterglMarkerSymbol108                     ScatterglMarkerSymbol = "108"
	ScatterglMarkerSymbolTriangleRightOpen       ScatterglMarkerSymbol = "triangle-right-open"
	ScatterglMarkerSymbolNumber208               ScatterglMarkerSymbol = 208
	ScatterglMarkerSymbol208                     ScatterglMarkerSymbol = "208"
	ScatterglMarkerSymbolTriangleRightDot        ScatterglMarkerSymbol = "triangle-right-dot"
	ScatterglMarkerSymbolNumber308               ScatterglMarkerSymbol = 308
	ScatterglMarkerSymbol308                     ScatterglMarkerSymbol = "308"
	ScatterglMarkerSymbolTriangleRightOpenDot    ScatterglMarkerSymbol = "triangle-right-open-dot"
	ScatterglMarkerSymbolNumber9                 ScatterglMarkerSymbol = 9
	ScatterglMarkerSymbol9                       ScatterglMarkerSymbol = "9"
	ScatterglMarkerSymbolTriangleNe              ScatterglMarkerSymbol = "triangle-ne"
	ScatterglMarkerSymbolNumber109               ScatterglMarkerSymbol = 109
	ScatterglMarkerSymbol109                     ScatterglMarkerSymbol = "109"
	ScatterglMarkerSymbolTriangleNeOpen          ScatterglMarkerSymbol = "triangle-ne-open"
	ScatterglMarkerSymbolNumber209               ScatterglMarkerSymbol = 209
	ScatterglMarkerSymbol209                     ScatterglMarkerSymbol = "209"
	ScatterglMarkerSymbolTriangleNeDot           ScatterglMarkerSymbol = "triangle-ne-dot"
	ScatterglMarkerSymbolNumber309               ScatterglMarkerSymbol = 309
	ScatterglMarkerSymbol309                     ScatterglMarkerSymbol = "309"
	ScatterglMarkerSymbolTriangleNeOpenDot       ScatterglMarkerSymbol = "triangle-ne-open-dot"
	ScatterglMarkerSymbolNumber10                ScatterglMarkerSymbol = 10
	ScatterglMarkerSymbol10                      ScatterglMarkerSymbol = "10"
	ScatterglMarkerSymbolTriangleSe              ScatterglMarkerSymbol = "triangle-se"
	ScatterglMarkerSymbolNumber110               ScatterglMarkerSymbol = 110
	ScatterglMarkerSymbol110                     ScatterglMarkerSymbol = "110"
	ScatterglMarkerSymbolTriangleSeOpen          ScatterglMarkerSymbol = "triangle-se-open"
	ScatterglMarkerSymbolNumber210               ScatterglMarkerSymbol = 210
	ScatterglMarkerSymbol210                     ScatterglMarkerSymbol = "210"
	ScatterglMarkerSymbolTriangleSeDot           ScatterglMarkerSymbol = "triangle-se-dot"
	ScatterglMarkerSymbolNumber310               ScatterglMarkerSymbol = 310
	ScatterglMarkerSymbol310                     ScatterglMarkerSymbol = "310"
	ScatterglMarkerSymbolTriangleSeOpenDot       ScatterglMarkerSymbol = "triangle-se-open-dot"
	ScatterglMarkerSymbolNumber11                ScatterglMarkerSymbol = 11
	ScatterglMarkerSymbol11                      ScatterglMarkerSymbol = "11"
	ScatterglMarkerSymbolTriangleSw              ScatterglMarkerSymbol = "triangle-sw"
	ScatterglMarkerSymbolNumber111               ScatterglMarkerSymbol = 111
	ScatterglMarkerSymbol111                     ScatterglMarkerSymbol = "111"
	ScatterglMarkerSymbolTriangleSwOpen          ScatterglMarkerSymbol = "triangle-sw-open"
	ScatterglMarkerSymbolNumber211               ScatterglMarkerSymbol = 211
	ScatterglMarkerSymbol211                     ScatterglMarkerSymbol = "211"
	ScatterglMarkerSymbolTriangleSwDot           ScatterglMarkerSymbol = "triangle-sw-dot"
	ScatterglMarkerSymbolNumber311               ScatterglMarkerSymbol = 311
	ScatterglMarkerSymbol311                     ScatterglMarkerSymbol = "311"
	ScatterglMarkerSymbolTriangleSwOpenDot       ScatterglMarkerSymbol = "triangle-sw-open-dot"
	ScatterglMarkerSymbolNumber12                ScatterglMarkerSymbol = 12
	ScatterglMarkerSymbol12                      ScatterglMarkerSymbol = "12"
	ScatterglMarkerSymbolTriangleNw              ScatterglMarkerSymbol = "triangle-nw"
	ScatterglMarkerSymbolNumber112               ScatterglMarkerSymbol = 112
	ScatterglMarkerSymbol112                     ScatterglMarkerSymbol = "112"
	ScatterglMarkerSymbolTriangleNwOpen          ScatterglMarkerSymbol = "triangle-nw-open"
	ScatterglMarkerSymbolNumber212               ScatterglMarkerSymbol = 212
	ScatterglMarkerSymbol212                     ScatterglMarkerSymbol = "212"
	ScatterglMarkerSymbolTriangleNwDot           ScatterglMarkerSymbol = "triangle-nw-dot"
	ScatterglMarkerSymbolNumber312               ScatterglMarkerSymbol = 312
	ScatterglMarkerSymbol312                     ScatterglMarkerSymbol = "312"
	ScatterglMarkerSymbolTriangleNwOpenDot       ScatterglMarkerSymbol = "triangle-nw-open-dot"
	ScatterglMarkerSymbolNumber13                ScatterglMarkerSymbol = 13
	ScatterglMarkerSymbol13                      ScatterglMarkerSymbol = "13"
	ScatterglMarkerSymbolPentagon                ScatterglMarkerSymbol = "pentagon"
	ScatterglMarkerSymbolNumber113               ScatterglMarkerSymbol = 113
	ScatterglMarkerSymbol113                     ScatterglMarkerSymbol = "113"
	ScatterglMarkerSymbolPentagonOpen            ScatterglMarkerSymbol = "pentagon-open"
	ScatterglMarkerSymbolNumber213               ScatterglMarkerSymbol = 213
	ScatterglMarkerSymbol213                     ScatterglMarkerSymbol = "213"
	ScatterglMarkerSymbolPentagonDot             ScatterglMarkerSymbol = "pentagon-dot"
	ScatterglMarkerSymbolNumber313               ScatterglMarkerSymbol = 313
	ScatterglMarkerSymbol313                     ScatterglMarkerSymbol = "313"
	ScatterglMarkerSymbolPentagonOpenDot         ScatterglMarkerSymbol = "pentagon-open-dot"
	ScatterglMarkerSymbolNumber14                ScatterglMarkerSymbol = 14
	ScatterglMarkerSymbol14                      ScatterglMarkerSymbol = "14"
	ScatterglMarkerSymbolHexagon                 ScatterglMarkerSymbol = "hexagon"
	ScatterglMarkerSymbolNumber114               ScatterglMarkerSymbol = 114
	ScatterglMarkerSymbol114                     ScatterglMarkerSymbol = "114"
	ScatterglMarkerSymbolHexagonOpen             ScatterglMarkerSymbol = "hexagon-open"
	ScatterglMarkerSymbolNumber214               ScatterglMarkerSymbol = 214
	ScatterglMarkerSymbol214                     ScatterglMarkerSymbol = "214"
	ScatterglMarkerSymbolHexagonDot              ScatterglMarkerSymbol = "hexagon-dot"
	ScatterglMarkerSymbolNumber314               ScatterglMarkerSymbol = 314
	ScatterglMarkerSymbol314                     ScatterglMarkerSymbol = "314"
	ScatterglMarkerSymbolHexagonOpenDot          ScatterglMarkerSymbol = "hexagon-open-dot"
	ScatterglMarkerSymbolNumber15                ScatterglMarkerSymbol = 15
	ScatterglMarkerSymbol15                      ScatterglMarkerSymbol = "15"
	ScatterglMarkerSymbolHexagon2                ScatterglMarkerSymbol = "hexagon2"
	ScatterglMarkerSymbolNumber115               ScatterglMarkerSymbol = 115
	ScatterglMarkerSymbol115                     ScatterglMarkerSymbol = "115"
	ScatterglMarkerSymbolHexagon2Open            ScatterglMarkerSymbol = "hexagon2-open"
	ScatterglMarkerSymbolNumber215               ScatterglMarkerSymbol = 215
	ScatterglMarkerSymbol215                     ScatterglMarkerSymbol = "215"
	ScatterglMarkerSymbolHexagon2Dot             ScatterglMarkerSymbol = "hexagon2-dot"
	ScatterglMarkerSymbolNumber315               ScatterglMarkerSymbol = 315
	ScatterglMarkerSymbol315                     ScatterglMarkerSymbol = "315"
	ScatterglMarkerSymbolHexagon2OpenDot         ScatterglMarkerSymbol = "hexagon2-open-dot"
	ScatterglMarkerSymbolNumber16                ScatterglMarkerSymbol = 16
	ScatterglMarkerSymbol16                      ScatterglMarkerSymbol = "16"
	ScatterglMarkerSymbolOctagon                 ScatterglMarkerSymbol = "octagon"
	ScatterglMarkerSymbolNumber116               ScatterglMarkerSymbol = 116
	ScatterglMarkerSymbol116                     ScatterglMarkerSymbol = "116"
	ScatterglMarkerSymbolOctagonOpen             ScatterglMarkerSymbol = "octagon-open"
	ScatterglMarkerSymbolNumber216               ScatterglMarkerSymbol = 216
	ScatterglMarkerSymbol216                     ScatterglMarkerSymbol = "216"
	ScatterglMarkerSymbolOctagonDot              ScatterglMarkerSymbol = "octagon-dot"
	ScatterglMarkerSymbolNumber316               ScatterglMarkerSymbol = 316
	ScatterglMarkerSymbol316                     ScatterglMarkerSymbol = "316"
	ScatterglMarkerSymbolOctagonOpenDot          ScatterglMarkerSymbol = "octagon-open-dot"
	ScatterglMarkerSymbolNumber17                ScatterglMarkerSymbol = 17
	ScatterglMarkerSymbol17                      ScatterglMarkerSymbol = "17"
	ScatterglMarkerSymbolStar                    ScatterglMarkerSymbol = "star"
	ScatterglMarkerSymbolNumber117               ScatterglMarkerSymbol = 117
	ScatterglMarkerSymbol117                     ScatterglMarkerSymbol = "117"
	ScatterglMarkerSymbolStarOpen                ScatterglMarkerSymbol = "star-open"
	ScatterglMarkerSymbolNumber217               ScatterglMarkerSymbol = 217
	ScatterglMarkerSymbol217                     ScatterglMarkerSymbol = "217"
	ScatterglMarkerSymbolStarDot                 ScatterglMarkerSymbol = "star-dot"
	ScatterglMarkerSymbolNumber317               ScatterglMarkerSymbol = 317
	ScatterglMarkerSymbol317                     ScatterglMarkerSymbol = "317"
	ScatterglMarkerSymbolStarOpenDot             ScatterglMarkerSymbol = "star-open-dot"
	ScatterglMarkerSymbolNumber18                ScatterglMarkerSymbol = 18
	ScatterglMarkerSymbol18                      ScatterglMarkerSymbol = "18"
	ScatterglMarkerSymbolHexagram                ScatterglMarkerSymbol = "hexagram"
	ScatterglMarkerSymbolNumber118               ScatterglMarkerSymbol = 118
	ScatterglMarkerSymbol118                     ScatterglMarkerSymbol = "118"
	ScatterglMarkerSymbolHexagramOpen            ScatterglMarkerSymbol = "hexagram-open"
	ScatterglMarkerSymbolNumber218               ScatterglMarkerSymbol = 218
	ScatterglMarkerSymbol218                     ScatterglMarkerSymbol = "218"
	ScatterglMarkerSymbolHexagramDot             ScatterglMarkerSymbol = "hexagram-dot"
	ScatterglMarkerSymbolNumber318               ScatterglMarkerSymbol = 318
	ScatterglMarkerSymbol318                     ScatterglMarkerSymbol = "318"
	ScatterglMarkerSymbolHexagramOpenDot         ScatterglMarkerSymbol = "hexagram-open-dot"
	ScatterglMarkerSymbolNumber19                ScatterglMarkerSymbol = 19
	ScatterglMarkerSymbol19                      ScatterglMarkerSymbol = "19"
	ScatterglMarkerSymbolStarTriangleUp          ScatterglMarkerSymbol = "star-triangle-up"
	ScatterglMarkerSymbolNumber119               ScatterglMarkerSymbol = 119
	ScatterglMarkerSymbol119                     ScatterglMarkerSymbol = "119"
	ScatterglMarkerSymbolStarTriangleUpOpen      ScatterglMarkerSymbol = "star-triangle-up-open"
	ScatterglMarkerSymbolNumber219               ScatterglMarkerSymbol = 219
	ScatterglMarkerSymbol219                     ScatterglMarkerSymbol = "219"
	ScatterglMarkerSymbolStarTriangleUpDot       ScatterglMarkerSymbol = "star-triangle-up-dot"
	ScatterglMarkerSymbolNumber319               ScatterglMarkerSymbol = 319
	ScatterglMarkerSymbol319                     ScatterglMarkerSymbol = "319"
	ScatterglMarkerSymbolStarTriangleUpOpenDot   ScatterglMarkerSymbol = "star-triangle-up-open-dot"
	ScatterglMarkerSymbolNumber20                ScatterglMarkerSymbol = 20
	ScatterglMarkerSymbol20                      ScatterglMarkerSymbol = "20"
	ScatterglMarkerSymbolStarTriangleDown        ScatterglMarkerSymbol = "star-triangle-down"
	ScatterglMarkerSymbolNumber120               ScatterglMarkerSymbol = 120
	ScatterglMarkerSymbol120                     ScatterglMarkerSymbol = "120"
	ScatterglMarkerSymbolStarTriangleDownOpen    ScatterglMarkerSymbol = "star-triangle-down-open"
	ScatterglMarkerSymbolNumber220               ScatterglMarkerSymbol = 220
	ScatterglMarkerSymbol220                     ScatterglMarkerSymbol = "220"
	ScatterglMarkerSymbolStarTriangleDownDot     ScatterglMarkerSymbol = "star-triangle-down-dot"
	ScatterglMarkerSymbolNumber320               ScatterglMarkerSymbol = 320
	ScatterglMarkerSymbol320                     ScatterglMarkerSymbol = "320"
	ScatterglMarkerSymbolStarTriangleDownOpenDot ScatterglMarkerSymbol = "star-triangle-down-open-dot"
	ScatterglMarkerSymbolNumber21                ScatterglMarkerSymbol = 21
	ScatterglMarkerSymbol21                      ScatterglMarkerSymbol = "21"
	ScatterglMarkerSymbolStarSquare              ScatterglMarkerSymbol = "star-square"
	ScatterglMarkerSymbolNumber121               ScatterglMarkerSymbol = 121
	ScatterglMarkerSymbol121                     ScatterglMarkerSymbol = "121"
	ScatterglMarkerSymbolStarSquareOpen          ScatterglMarkerSymbol = "star-square-open"
	ScatterglMarkerSymbolNumber221               ScatterglMarkerSymbol = 221
	ScatterglMarkerSymbol221                     ScatterglMarkerSymbol = "221"
	ScatterglMarkerSymbolStarSquareDot           ScatterglMarkerSymbol = "star-square-dot"
	ScatterglMarkerSymbolNumber321               ScatterglMarkerSymbol = 321
	ScatterglMarkerSymbol321                     ScatterglMarkerSymbol = "321"
	ScatterglMarkerSymbolStarSquareOpenDot       ScatterglMarkerSymbol = "star-square-open-dot"
	ScatterglMarkerSymbolNumber22                ScatterglMarkerSymbol = 22
	ScatterglMarkerSymbol22                      ScatterglMarkerSymbol = "22"
	ScatterglMarkerSymbolStarDiamond             ScatterglMarkerSymbol = "star-diamond"
	ScatterglMarkerSymbolNumber122               ScatterglMarkerSymbol = 122
	ScatterglMarkerSymbol122                     ScatterglMarkerSymbol = "122"
	ScatterglMarkerSymbolStarDiamondOpen         ScatterglMarkerSymbol = "star-diamond-open"
	ScatterglMarkerSymbolNumber222               ScatterglMarkerSymbol = 222
	ScatterglMarkerSymbol222                     ScatterglMarkerSymbol = "222"
	ScatterglMarkerSymbolStarDiamondDot          ScatterglMarkerSymbol = "star-diamond-dot"
	ScatterglMarkerSymbolNumber322               ScatterglMarkerSymbol = 322
	ScatterglMarkerSymbol322                     ScatterglMarkerSymbol = "322"
	ScatterglMarkerSymbolStarDiamondOpenDot      ScatterglMarkerSymbol = "star-diamond-open-dot"
	ScatterglMarkerSymbolNumber23                ScatterglMarkerSymbol = 23
	ScatterglMarkerSymbol23                      ScatterglMarkerSymbol = "23"
	ScatterglMarkerSymbolDiamondTall             ScatterglMarkerSymbol = "diamond-tall"
	ScatterglMarkerSymbolNumber123               ScatterglMarkerSymbol = 123
	ScatterglMarkerSymbol123                     ScatterglMarkerSymbol = "123"
	ScatterglMarkerSymbolDiamondTallOpen         ScatterglMarkerSymbol = "diamond-tall-open"
	ScatterglMarkerSymbolNumber223               ScatterglMarkerSymbol = 223
	ScatterglMarkerSymbol223                     ScatterglMarkerSymbol = "223"
	ScatterglMarkerSymbolDiamondTallDot          ScatterglMarkerSymbol = "diamond-tall-dot"
	ScatterglMarkerSymbolNumber323               ScatterglMarkerSymbol = 323
	ScatterglMarkerSymbol323                     ScatterglMarkerSymbol = "323"
	ScatterglMarkerSymbolDiamondTallOpenDot      ScatterglMarkerSymbol = "diamond-tall-open-dot"
	ScatterglMarkerSymbolNumber24                ScatterglMarkerSymbol = 24
	ScatterglMarkerSymbol24                      ScatterglMarkerSymbol = "24"
	ScatterglMarkerSymbolDiamondWide             ScatterglMarkerSymbol = "diamond-wide"
	ScatterglMarkerSymbolNumber124               ScatterglMarkerSymbol = 124
	ScatterglMarkerSymbol124                     ScatterglMarkerSymbol = "124"
	ScatterglMarkerSymbolDiamondWideOpen         ScatterglMarkerSymbol = "diamond-wide-open"
	ScatterglMarkerSymbolNumber224               ScatterglMarkerSymbol = 224
	ScatterglMarkerSymbol224                     ScatterglMarkerSymbol = "224"
	ScatterglMarkerSymbolDiamondWideDot          ScatterglMarkerSymbol = "diamond-wide-dot"
	ScatterglMarkerSymbolNumber324               ScatterglMarkerSymbol = 324
	ScatterglMarkerSymbol324                     ScatterglMarkerSymbol = "324"
	ScatterglMarkerSymbolDiamondWideOpenDot      ScatterglMarkerSymbol = "diamond-wide-open-dot"
	ScatterglMarkerSymbolNumber25                ScatterglMarkerSymbol = 25
	ScatterglMarkerSymbol25                      ScatterglMarkerSymbol = "25"
	ScatterglMarkerSymbolHourglass               ScatterglMarkerSymbol = "hourglass"
	ScatterglMarkerSymbolNumber125               ScatterglMarkerSymbol = 125
	ScatterglMarkerSymbol125                     ScatterglMarkerSymbol = "125"
	ScatterglMarkerSymbolHourglassOpen           ScatterglMarkerSymbol = "hourglass-open"
	ScatterglMarkerSymbolNumber26                ScatterglMarkerSymbol = 26
	ScatterglMarkerSymbol26                      ScatterglMarkerSymbol = "26"
	ScatterglMarkerSymbolBowtie                  ScatterglMarkerSymbol = "bowtie"
	ScatterglMarkerSymbolNumber126               ScatterglMarkerSymbol = 126
	ScatterglMarkerSymbol126                     ScatterglMarkerSymbol = "126"
	ScatterglMarkerSymbolBowtieOpen              ScatterglMarkerSymbol = "bowtie-open"
	ScatterglMarkerSymbolNumber27                ScatterglMarkerSymbol = 27
	ScatterglMarkerSymbol27                      ScatterglMarkerSymbol = "27"
	ScatterglMarkerSymbolCircleCross             ScatterglMarkerSymbol = "circle-cross"
	ScatterglMarkerSymbolNumber127               ScatterglMarkerSymbol = 127
	ScatterglMarkerSymbol127                     ScatterglMarkerSymbol = "127"
	ScatterglMarkerSymbolCircleCrossOpen         ScatterglMarkerSymbol = "circle-cross-open"
	ScatterglMarkerSymbolNumber28                ScatterglMarkerSymbol = 28
	ScatterglMarkerSymbol28                      ScatterglMarkerSymbol = "28"
	ScatterglMarkerSymbolCircleX                 ScatterglMarkerSymbol = "circle-x"
	ScatterglMarkerSymbolNumber128               ScatterglMarkerSymbol = 128
	ScatterglMarkerSymbol128                     ScatterglMarkerSymbol = "128"
	ScatterglMarkerSymbolCircleXOpen             ScatterglMarkerSymbol = "circle-x-open"
	ScatterglMarkerSymbolNumber29                ScatterglMarkerSymbol = 29
	ScatterglMarkerSymbol29                      ScatterglMarkerSymbol = "29"
	ScatterglMarkerSymbolSquareCross             ScatterglMarkerSymbol = "square-cross"
	ScatterglMarkerSymbolNumber129               ScatterglMarkerSymbol = 129
	ScatterglMarkerSymbol129                     ScatterglMarkerSymbol = "129"
	ScatterglMarkerSymbolSquareCrossOpen         ScatterglMarkerSymbol = "square-cross-open"
	ScatterglMarkerSymbolNumber30                ScatterglMarkerSymbol = 30
	ScatterglMarkerSymbol30                      ScatterglMarkerSymbol = "30"
	ScatterglMarkerSymbolSquareX                 ScatterglMarkerSymbol = "square-x"
	ScatterglMarkerSymbolNumber130               ScatterglMarkerSymbol = 130
	ScatterglMarkerSymbol130                     ScatterglMarkerSymbol = "130"
	ScatterglMarkerSymbolSquareXOpen             ScatterglMarkerSymbol = "square-x-open"
	ScatterglMarkerSymbolNumber31                ScatterglMarkerSymbol = 31
	ScatterglMarkerSymbol31                      ScatterglMarkerSymbol = "31"
	ScatterglMarkerSymbolDiamondCross            ScatterglMarkerSymbol = "diamond-cross"
	ScatterglMarkerSymbolNumber131               ScatterglMarkerSymbol = 131
	ScatterglMarkerSymbol131                     ScatterglMarkerSymbol = "131"
	ScatterglMarkerSymbolDiamondCrossOpen        ScatterglMarkerSymbol = "diamond-cross-open"
	ScatterglMarkerSymbolNumber32                ScatterglMarkerSymbol = 32
	ScatterglMarkerSymbol32                      ScatterglMarkerSymbol = "32"
	ScatterglMarkerSymbolDiamondX                ScatterglMarkerSymbol = "diamond-x"
	ScatterglMarkerSymbolNumber132               ScatterglMarkerSymbol = 132
	ScatterglMarkerSymbol132                     ScatterglMarkerSymbol = "132"
	ScatterglMarkerSymbolDiamondXOpen            ScatterglMarkerSymbol = "diamond-x-open"
	ScatterglMarkerSymbolNumber33                ScatterglMarkerSymbol = 33
	ScatterglMarkerSymbol33                      ScatterglMarkerSymbol = "33"
	ScatterglMarkerSymbolCrossThin               ScatterglMarkerSymbol = "cross-thin"
	ScatterglMarkerSymbolNumber133               ScatterglMarkerSymbol = 133
	ScatterglMarkerSymbol133                     ScatterglMarkerSymbol = "133"
	ScatterglMarkerSymbolCrossThinOpen           ScatterglMarkerSymbol = "cross-thin-open"
	ScatterglMarkerSymbolNumber34                ScatterglMarkerSymbol = 34
	ScatterglMarkerSymbol34                      ScatterglMarkerSymbol = "34"
	ScatterglMarkerSymbolXThin                   ScatterglMarkerSymbol = "x-thin"
	ScatterglMarkerSymbolNumber134               ScatterglMarkerSymbol = 134
	ScatterglMarkerSymbol134                     ScatterglMarkerSymbol = "134"
	ScatterglMarkerSymbolXThinOpen               ScatterglMarkerSymbol = "x-thin-open"
	ScatterglMarkerSymbolNumber35                ScatterglMarkerSymbol = 35
	ScatterglMarkerSymbol35                      ScatterglMarkerSymbol = "35"
	ScatterglMarkerSymbolAsterisk                ScatterglMarkerSymbol = "asterisk"
	ScatterglMarkerSymbolNumber135               ScatterglMarkerSymbol = 135
	ScatterglMarkerSymbol135                     ScatterglMarkerSymbol = "135"
	ScatterglMarkerSymbolAsteriskOpen            ScatterglMarkerSymbol = "asterisk-open"
	ScatterglMarkerSymbolNumber36                ScatterglMarkerSymbol = 36
	ScatterglMarkerSymbol36                      ScatterglMarkerSymbol = "36"
	ScatterglMarkerSymbolHash                    ScatterglMarkerSymbol = "hash"
	ScatterglMarkerSymbolNumber136               ScatterglMarkerSymbol = 136
	ScatterglMarkerSymbol136                     ScatterglMarkerSymbol = "136"
	ScatterglMarkerSymbolHashOpen                ScatterglMarkerSymbol = "hash-open"
	ScatterglMarkerSymbolNumber236               ScatterglMarkerSymbol = 236
	ScatterglMarkerSymbol236                     ScatterglMarkerSymbol = "236"
	ScatterglMarkerSymbolHashDot                 ScatterglMarkerSymbol = "hash-dot"
	ScatterglMarkerSymbolNumber336               ScatterglMarkerSymbol = 336
	ScatterglMarkerSymbol336                     ScatterglMarkerSymbol = "336"
	ScatterglMarkerSymbolHashOpenDot             ScatterglMarkerSymbol = "hash-open-dot"
	ScatterglMarkerSymbolNumber37                ScatterglMarkerSymbol = 37
	ScatterglMarkerSymbol37                      ScatterglMarkerSymbol = "37"
	ScatterglMarkerSymbolYUp                     ScatterglMarkerSymbol = "y-up"
	ScatterglMarkerSymbolNumber137               ScatterglMarkerSymbol = 137
	ScatterglMarkerSymbol137                     ScatterglMarkerSymbol = "137"
	ScatterglMarkerSymbolYUpOpen                 ScatterglMarkerSymbol = "y-up-open"
	ScatterglMarkerSymbolNumber38                ScatterglMarkerSymbol = 38
	ScatterglMarkerSymbol38                      ScatterglMarkerSymbol = "38"
	ScatterglMarkerSymbolYDown                   ScatterglMarkerSymbol = "y-down"
	ScatterglMarkerSymbolNumber138               ScatterglMarkerSymbol = 138
	ScatterglMarkerSymbol138                     ScatterglMarkerSymbol = "138"
	ScatterglMarkerSymbolYDownOpen               ScatterglMarkerSymbol = "y-down-open"
	ScatterglMarkerSymbolNumber39                ScatterglMarkerSymbol = 39
	ScatterglMarkerSymbol39                      ScatterglMarkerSymbol = "39"
	ScatterglMarkerSymbolYLeft                   ScatterglMarkerSymbol = "y-left"
	ScatterglMarkerSymbolNumber139               ScatterglMarkerSymbol = 139
	ScatterglMarkerSymbol139                     ScatterglMarkerSymbol = "139"
	ScatterglMarkerSymbolYLeftOpen               ScatterglMarkerSymbol = "y-left-open"
	ScatterglMarkerSymbolNumber40                ScatterglMarkerSymbol = 40
	ScatterglMarkerSymbol40                      ScatterglMarkerSymbol = "40"
	ScatterglMarkerSymbolYRight                  ScatterglMarkerSymbol = "y-right"
	ScatterglMarkerSymbolNumber140               ScatterglMarkerSymbol = 140
	ScatterglMarkerSymbol140                     ScatterglMarkerSymbol = "140"
	ScatterglMarkerSymbolYRightOpen              ScatterglMarkerSymbol = "y-right-open"
	ScatterglMarkerSymbolNumber41                ScatterglMarkerSymbol = 41
	ScatterglMarkerSymbol41                      ScatterglMarkerSymbol = "41"
	ScatterglMarkerSymbolLineEw                  ScatterglMarkerSymbol = "line-ew"
	ScatterglMarkerSymbolNumber141               ScatterglMarkerSymbol = 141
	ScatterglMarkerSymbol141                     ScatterglMarkerSymbol = "141"
	ScatterglMarkerSymbolLineEwOpen              ScatterglMarkerSymbol = "line-ew-open"
	ScatterglMarkerSymbolNumber42                ScatterglMarkerSymbol = 42
	ScatterglMarkerSymbol42                      ScatterglMarkerSymbol = "42"
	ScatterglMarkerSymbolLineNs                  ScatterglMarkerSymbol = "line-ns"
	ScatterglMarkerSymbolNumber142               ScatterglMarkerSymbol = 142
	ScatterglMarkerSymbol142                     ScatterglMarkerSymbol = "142"
	ScatterglMarkerSymbolLineNsOpen              ScatterglMarkerSymbol = "line-ns-open"
	ScatterglMarkerSymbolNumber43                ScatterglMarkerSymbol = 43
	ScatterglMarkerSymbol43                      ScatterglMarkerSymbol = "43"
	ScatterglMarkerSymbolLineNe                  ScatterglMarkerSymbol = "line-ne"
	ScatterglMarkerSymbolNumber143               ScatterglMarkerSymbol = 143
	ScatterglMarkerSymbol143                     ScatterglMarkerSymbol = "143"
	ScatterglMarkerSymbolLineNeOpen              ScatterglMarkerSymbol = "line-ne-open"
	ScatterglMarkerSymbolNumber44                ScatterglMarkerSymbol = 44
	ScatterglMarkerSymbol44                      ScatterglMarkerSymbol = "44"
	ScatterglMarkerSymbolLineNw                  ScatterglMarkerSymbol = "line-nw"
	ScatterglMarkerSymbolNumber144               ScatterglMarkerSymbol = 144
	ScatterglMarkerSymbol144                     ScatterglMarkerSymbol = "144"
	ScatterglMarkerSymbolLineNwOpen              ScatterglMarkerSymbol = "line-nw-open"
	ScatterglMarkerSymbolNumber45                ScatterglMarkerSymbol = 45
	ScatterglMarkerSymbol45                      ScatterglMarkerSymbol = "45"
	ScatterglMarkerSymbolArrowUp                 ScatterglMarkerSymbol = "arrow-up"
	ScatterglMarkerSymbolNumber145               ScatterglMarkerSymbol = 145
	ScatterglMarkerSymbol145                     ScatterglMarkerSymbol = "145"
	ScatterglMarkerSymbolArrowUpOpen             ScatterglMarkerSymbol = "arrow-up-open"
	ScatterglMarkerSymbolNumber46                ScatterglMarkerSymbol = 46
	ScatterglMarkerSymbol46                      ScatterglMarkerSymbol = "46"
	ScatterglMarkerSymbolArrowDown               ScatterglMarkerSymbol = "arrow-down"
	ScatterglMarkerSymbolNumber146               ScatterglMarkerSymbol = 146
	ScatterglMarkerSymbol146                     ScatterglMarkerSymbol = "146"
	ScatterglMarkerSymbolArrowDownOpen           ScatterglMarkerSymbol = "arrow-down-open"
	ScatterglMarkerSymbolNumber47                ScatterglMarkerSymbol = 47
	ScatterglMarkerSymbol47                      ScatterglMarkerSymbol = "47"
	ScatterglMarkerSymbolArrowLeft               ScatterglMarkerSymbol = "arrow-left"
	ScatterglMarkerSymbolNumber147               ScatterglMarkerSymbol = 147
	ScatterglMarkerSymbol147                     ScatterglMarkerSymbol = "147"
	ScatterglMarkerSymbolArrowLeftOpen           ScatterglMarkerSymbol = "arrow-left-open"
	ScatterglMarkerSymbolNumber48                ScatterglMarkerSymbol = 48
	ScatterglMarkerSymbol48                      ScatterglMarkerSymbol = "48"
	ScatterglMarkerSymbolArrowRight              ScatterglMarkerSymbol = "arrow-right"
	ScatterglMarkerSymbolNumber148               ScatterglMarkerSymbol = 148
	ScatterglMarkerSymbol148                     ScatterglMarkerSymbol = "148"
	ScatterglMarkerSymbolArrowRightOpen          ScatterglMarkerSymbol = "arrow-right-open"
	ScatterglMarkerSymbolNumber49                ScatterglMarkerSymbol = 49
	ScatterglMarkerSymbol49                      ScatterglMarkerSymbol = "49"
	ScatterglMarkerSymbolArrowBarUp              ScatterglMarkerSymbol = "arrow-bar-up"
	ScatterglMarkerSymbolNumber149               ScatterglMarkerSymbol = 149
	ScatterglMarkerSymbol149                     ScatterglMarkerSymbol = "149"
	ScatterglMarkerSymbolArrowBarUpOpen          ScatterglMarkerSymbol = "arrow-bar-up-open"
	ScatterglMarkerSymbolNumber50                ScatterglMarkerSymbol = 50
	ScatterglMarkerSymbol50                      ScatterglMarkerSymbol = "50"
	ScatterglMarkerSymbolArrowBarDown            ScatterglMarkerSymbol = "arrow-bar-down"
	ScatterglMarkerSymbolNumber150               ScatterglMarkerSymbol = 150
	ScatterglMarkerSymbol150                     ScatterglMarkerSymbol = "150"
	ScatterglMarkerSymbolArrowBarDownOpen        ScatterglMarkerSymbol = "arrow-bar-down-open"
	ScatterglMarkerSymbolNumber51                ScatterglMarkerSymbol = 51
	ScatterglMarkerSymbol51                      ScatterglMarkerSymbol = "51"
	ScatterglMarkerSymbolArrowBarLeft            ScatterglMarkerSymbol = "arrow-bar-left"
	ScatterglMarkerSymbolNumber151               ScatterglMarkerSymbol = 151
	ScatterglMarkerSymbol151                     ScatterglMarkerSymbol = "151"
	ScatterglMarkerSymbolArrowBarLeftOpen        ScatterglMarkerSymbol = "arrow-bar-left-open"
	ScatterglMarkerSymbolNumber52                ScatterglMarkerSymbol = 52
	ScatterglMarkerSymbol52                      ScatterglMarkerSymbol = "52"
	ScatterglMarkerSymbolArrowBarRight           ScatterglMarkerSymbol = "arrow-bar-right"
	ScatterglMarkerSymbolNumber152               ScatterglMarkerSymbol = 152
	ScatterglMarkerSymbol152                     ScatterglMarkerSymbol = "152"
	ScatterglMarkerSymbolArrowBarRightOpen       ScatterglMarkerSymbol = "arrow-bar-right-open"
)

// ScatterglTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterglTextposition string

const (
	ScatterglTextpositionTopLeft      ScatterglTextposition = "top left"
	ScatterglTextpositionTopCenter    ScatterglTextposition = "top center"
	ScatterglTextpositionTopRight     ScatterglTextposition = "top right"
	ScatterglTextpositionMiddleLeft   ScatterglTextposition = "middle left"
	ScatterglTextpositionMiddleCenter ScatterglTextposition = "middle center"
	ScatterglTextpositionMiddleRight  ScatterglTextposition = "middle right"
	ScatterglTextpositionBottomLeft   ScatterglTextposition = "bottom left"
	ScatterglTextpositionBottomCenter ScatterglTextposition = "bottom center"
	ScatterglTextpositionBottomRight  ScatterglTextposition = "bottom right"
)

// ScatterglVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterglVisible interface{}

var (
	ScatterglVisibleTrue       ScatterglVisible = true
	ScatterglVisibleFalse      ScatterglVisible = false
	ScatterglVisibleLegendonly ScatterglVisible = "legendonly"
)

// ScatterglXcalendar Sets the calendar system to use with `x` date data.
type ScatterglXcalendar string

const (
	ScatterglXcalendarGregorian  ScatterglXcalendar = "gregorian"
	ScatterglXcalendarChinese    ScatterglXcalendar = "chinese"
	ScatterglXcalendarCoptic     ScatterglXcalendar = "coptic"
	ScatterglXcalendarDiscworld  ScatterglXcalendar = "discworld"
	ScatterglXcalendarEthiopian  ScatterglXcalendar = "ethiopian"
	ScatterglXcalendarHebrew     ScatterglXcalendar = "hebrew"
	ScatterglXcalendarIslamic    ScatterglXcalendar = "islamic"
	ScatterglXcalendarJulian     ScatterglXcalendar = "julian"
	ScatterglXcalendarMayan      ScatterglXcalendar = "mayan"
	ScatterglXcalendarNanakshahi ScatterglXcalendar = "nanakshahi"
	ScatterglXcalendarNepali     ScatterglXcalendar = "nepali"
	ScatterglXcalendarPersian    ScatterglXcalendar = "persian"
	ScatterglXcalendarJalali     ScatterglXcalendar = "jalali"
	ScatterglXcalendarTaiwan     ScatterglXcalendar = "taiwan"
	ScatterglXcalendarThai       ScatterglXcalendar = "thai"
	ScatterglXcalendarUmmalqura  ScatterglXcalendar = "ummalqura"
)

// ScatterglXperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
type ScatterglXperiodalignment string

const (
	ScatterglXperiodalignmentStart  ScatterglXperiodalignment = "start"
	ScatterglXperiodalignmentMiddle ScatterglXperiodalignment = "middle"
	ScatterglXperiodalignmentEnd    ScatterglXperiodalignment = "end"
)

// ScatterglYcalendar Sets the calendar system to use with `y` date data.
type ScatterglYcalendar string

const (
	ScatterglYcalendarGregorian  ScatterglYcalendar = "gregorian"
	ScatterglYcalendarChinese    ScatterglYcalendar = "chinese"
	ScatterglYcalendarCoptic     ScatterglYcalendar = "coptic"
	ScatterglYcalendarDiscworld  ScatterglYcalendar = "discworld"
	ScatterglYcalendarEthiopian  ScatterglYcalendar = "ethiopian"
	ScatterglYcalendarHebrew     ScatterglYcalendar = "hebrew"
	ScatterglYcalendarIslamic    ScatterglYcalendar = "islamic"
	ScatterglYcalendarJulian     ScatterglYcalendar = "julian"
	ScatterglYcalendarMayan      ScatterglYcalendar = "mayan"
	ScatterglYcalendarNanakshahi ScatterglYcalendar = "nanakshahi"
	ScatterglYcalendarNepali     ScatterglYcalendar = "nepali"
	ScatterglYcalendarPersian    ScatterglYcalendar = "persian"
	ScatterglYcalendarJalali     ScatterglYcalendar = "jalali"
	ScatterglYcalendarTaiwan     ScatterglYcalendar = "taiwan"
	ScatterglYcalendarThai       ScatterglYcalendar = "thai"
	ScatterglYcalendarUmmalqura  ScatterglYcalendar = "ummalqura"
)

// ScatterglYperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
type ScatterglYperiodalignment string

const (
	ScatterglYperiodalignmentStart  ScatterglYperiodalignment = "start"
	ScatterglYperiodalignmentMiddle ScatterglYperiodalignment = "middle"
	ScatterglYperiodalignmentEnd    ScatterglYperiodalignment = "end"
)

// ScatterglHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type ScatterglHoverinfo string

const (
	// Flags
	ScatterglHoverinfoX    ScatterglHoverinfo = "x"
	ScatterglHoverinfoY    ScatterglHoverinfo = "y"
	ScatterglHoverinfoZ    ScatterglHoverinfo = "z"
	ScatterglHoverinfoText ScatterglHoverinfo = "text"
	ScatterglHoverinfoName ScatterglHoverinfo = "name"

	// Extra
	ScatterglHoverinfoAll  ScatterglHoverinfo = "all"
	ScatterglHoverinfoNone ScatterglHoverinfo = "none"
	ScatterglHoverinfoSkip ScatterglHoverinfo = "skip"
)

// ScatterglMode Determines the drawing mode for this scatter trace.
type ScatterglMode string

const (
	// Flags
	ScatterglModeLines   ScatterglMode = "lines"
	ScatterglModeMarkers ScatterglMode = "markers"
	ScatterglModeText    ScatterglMode = "text"

	// Extra
	ScatterglModeNone ScatterglMode = "none"
)
