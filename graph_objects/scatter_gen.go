package grob

var TraceTypeScatter TraceType = "scatter"

func (trace *Scatter) GetType() TraceType {
	return TraceTypeScatter
}

// Scatter The scatter trace type encompasses line charts, scatter charts, text charts, and bubble charts. The data visualized as scatter point or lines is set in `x` and `y`. Text (appearing either on the chart or on hover only) is via `text`. Bubble charts are achieved by setting `marker.size` and/or `marker.color` to numerical arrays.
type Scatter struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Cliponaxis
	// arrayOK: false
	// type: boolean
	// Determines whether or not markers and text nodes are clipped about the subplot axes. To show markers and text nodes above axis lines and tick labels, make sure to set `xaxis.layer` and `yaxis.layer` to *below traces*.
	Cliponaxis Bool `json:"cliponaxis,omitempty"`

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
	ErrorX *ScatterErrorX `json:"error_x,omitempty"`

	// ErrorY
	// role: Object
	ErrorY *ScatterErrorY `json:"error_y,omitempty"`

	// Fill
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
	Fill ScatterFill `json:"fill,omitempty"`

	// Fillcolor
	// arrayOK: false
	// type: color
	// Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor Color `json:"fillcolor,omitempty"`

	// Groupnorm
	// default:
	// type: enumerated
	// Only relevant when `stackgroup` is used, and only the first `groupnorm` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Sets the normalization for the sum of this `stackgroup`. With *fraction*, the value of each trace at each location is divided by the sum of all trace values at that location. *percent* is the same but multiplied by 100 to show percentages. If there are multiple subplots, or multiple `stackgroup`s on one subplot, each will be normalized within its own set.
	Groupnorm ScatterGroupnorm `json:"groupnorm,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ScatterHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *ScatterHoverlabel `json:"hoverlabel,omitempty"`

	// Hoveron
	// default: %!s(<nil>)
	// type: flaglist
	// Do the hover effects highlight individual points (markers or line points) or do they highlight filled regions? If the fill is *toself* or *tonext* and there are no markers or text, then the default is *fills*, otherwise it is *points*.
	Hoveron ScatterHoveron `json:"hoveron,omitempty"`

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
	Line *ScatterLine `json:"line,omitempty"`

	// Marker
	// role: Object
	Marker *ScatterMarker `json:"marker,omitempty"`

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
	// Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
	Mode ScatterMode `json:"mode,omitempty"`

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

	// Orientation
	// default: %!s(<nil>)
	// type: enumerated
	// Only relevant when `stackgroup` is used, and only the first `orientation` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Sets the stacking direction. With *v* (*h*), the y (x) values of subsequent traces are added. Also affects the default value of `fill`.
	Orientation ScatterOrientation `json:"orientation,omitempty"`

	// R
	// arrayOK: false
	// type: data_array
	// r coordinates in scatter traces are deprecated!Please switch to the *scatterpolar* trace type.Sets the radial coordinatesfor legacy polar chart only.
	R interface{} `json:"r,omitempty"`

	// Rsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  r .
	Rsrc String `json:"rsrc,omitempty"`

	// Selected
	// role: Object
	Selected *ScatterSelected `json:"selected,omitempty"`

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

	// Stackgaps
	// default: infer zero
	// type: enumerated
	// Only relevant when `stackgroup` is used, and only the first `stackgaps` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Determines how we handle locations at which other traces in this group have data but this one does not. With *infer zero* we insert a zero at these locations. With *interpolate* we linearly interpolate between existing values, and extrapolate a constant beyond the existing values.
	Stackgaps ScatterStackgaps `json:"stackgaps,omitempty"`

	// Stackgroup
	// arrayOK: false
	// type: string
	// Set several scatter traces (on the same subplot) to the same stackgroup in order to add their y values (or their x values if `orientation` is *h*). If blank or omitted this trace will not be stacked. Stacking also turns `fill` on by default, using *tonexty* (*tonextx*) if `orientation` is *h* (*v*) and sets the default `mode` to *lines* irrespective of point count. You can only stack on a numeric (linear or log) axis. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
	Stackgroup String `json:"stackgroup,omitempty"`

	// Stream
	// role: Object
	Stream *ScatterStream `json:"stream,omitempty"`

	// T
	// arrayOK: false
	// type: data_array
	// t coordinates in scatter traces are deprecated!Please switch to the *scatterpolar* trace type.Sets the angular coordinatesfor legacy polar chart only.
	T interface{} `json:"t,omitempty"`

	// Text
	// arrayOK: true
	// type: string
	// Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text String `json:"text,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScatterTextfont `json:"textfont,omitempty"`

	// Textposition
	// default: middle center
	// type: enumerated
	// Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition ScatterTextposition `json:"textposition,omitempty"`

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

	// Tsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  t .
	Tsrc String `json:"tsrc,omitempty"`

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
	Unselected *ScatterUnselected `json:"unselected,omitempty"`

	// Visible
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ScatterVisible `json:"visible,omitempty"`

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
	Xcalendar ScatterXcalendar `json:"xcalendar,omitempty"`

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
	Xperiodalignment ScatterXperiodalignment `json:"xperiodalignment,omitempty"`

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
	Ycalendar ScatterYcalendar `json:"ycalendar,omitempty"`

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
	Yperiodalignment ScatterYperiodalignment `json:"yperiodalignment,omitempty"`

	// Ysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

// ScatterErrorX
type ScatterErrorX struct {

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
	Type ScatterErrorXType `json:"type,omitempty"`

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

// ScatterErrorY
type ScatterErrorY struct {

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
	Type ScatterErrorYType `json:"type,omitempty"`

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

// ScatterHoverlabelFont Sets the font used in hover labels.
type ScatterHoverlabelFont struct {

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

// ScatterHoverlabel
type ScatterHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align ScatterHoverlabelAlign `json:"align,omitempty"`

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
	Font *ScatterHoverlabelFont `json:"font,omitempty"`

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

// ScatterLine
type ScatterLine struct {

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
	Shape ScatterLineShape `json:"shape,omitempty"`

	// Simplify
	// arrayOK: false
	// type: boolean
	// Simplifies lines by removing nearly-collinear points. When transitioning lines, it may be desirable to disable this so that the number of points along the resulting SVG path is unaffected.
	Simplify Bool `json:"simplify,omitempty"`

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

// ScatterMarkerColorbarTickfont Sets the color bar's tick label font
type ScatterMarkerColorbarTickfont struct {

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

// ScatterMarkerColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type ScatterMarkerColorbarTitleFont struct {

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

// ScatterMarkerColorbarTitle
type ScatterMarkerColorbarTitle struct {

	// Font
	// role: Object
	Font *ScatterMarkerColorbarTitleFont `json:"font,omitempty"`

	// Side
	// default: top
	// type: enumerated
	// Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
	Side ScatterMarkerColorbarTitleSide `json:"side,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// ScatterMarkerColorbar
type ScatterMarkerColorbar struct {

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
	Exponentformat ScatterMarkerColorbarExponentformat `json:"exponentformat,omitempty"`

	// Len
	// arrayOK: false
	// type: number
	// Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends.
	Len float64 `json:"len,omitempty"`

	// Lenmode
	// default: fraction
	// type: enumerated
	// Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
	Lenmode ScatterMarkerColorbarLenmode `json:"lenmode,omitempty"`

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
	Showexponent ScatterMarkerColorbarShowexponent `json:"showexponent,omitempty"`

	// Showticklabels
	// arrayOK: false
	// type: boolean
	// Determines whether or not the tick labels are drawn.
	Showticklabels Bool `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix ScatterMarkerColorbarShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix ScatterMarkerColorbarShowticksuffix `json:"showticksuffix,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels.
	Thickness float64 `json:"thickness,omitempty"`

	// Thicknessmode
	// default: pixels
	// type: enumerated
	// Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
	Thicknessmode ScatterMarkerColorbarThicknessmode `json:"thicknessmode,omitempty"`

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
	Tickfont *ScatterMarkerColorbarTickfont `json:"tickfont,omitempty"`

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
	Ticklabelposition ScatterMarkerColorbarTicklabelposition `json:"ticklabelposition,omitempty"`

	// Ticklen
	// arrayOK: false
	// type: number
	// Sets the tick length (in px).
	Ticklen float64 `json:"ticklen,omitempty"`

	// Tickmode
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
	Tickmode ScatterMarkerColorbarTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

	// Ticks
	// default:
	// type: enumerated
	// Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
	Ticks ScatterMarkerColorbarTicks `json:"ticks,omitempty"`

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
	Title *ScatterMarkerColorbarTitle `json:"title,omitempty"`

	// X
	// arrayOK: false
	// type: number
	// Sets the x position of the color bar (in plot fraction).
	X float64 `json:"x,omitempty"`

	// Xanchor
	// default: left
	// type: enumerated
	// Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
	Xanchor ScatterMarkerColorbarXanchor `json:"xanchor,omitempty"`

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
	Yanchor ScatterMarkerColorbarYanchor `json:"yanchor,omitempty"`

	// Ypad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the y direction.
	Ypad float64 `json:"ypad,omitempty"`
}

// ScatterMarkerGradient
type ScatterMarkerGradient struct {

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
	Type ScatterMarkerGradientType `json:"type,omitempty"`

	// Typesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  type .
	Typesrc String `json:"typesrc,omitempty"`
}

// ScatterMarkerLine
type ScatterMarkerLine struct {

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

// ScatterMarker
type ScatterMarker struct {

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
	Colorbar *ScatterMarkerColorbar `json:"colorbar,omitempty"`

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
	Gradient *ScatterMarkerGradient `json:"gradient,omitempty"`

	// Line
	// role: Object
	Line *ScatterMarkerLine `json:"line,omitempty"`

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
	Sizemode ScatterMarkerSizemode `json:"sizemode,omitempty"`

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
	Symbol ScatterMarkerSymbol `json:"symbol,omitempty"`

	// Symbolsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  symbol .
	Symbolsrc String `json:"symbolsrc,omitempty"`
}

// ScatterSelectedMarker
type ScatterSelectedMarker struct {

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

// ScatterSelectedTextfont
type ScatterSelectedTextfont struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the text font color of selected points.
	Color Color `json:"color,omitempty"`
}

// ScatterSelected
type ScatterSelected struct {

	// Marker
	// role: Object
	Marker *ScatterSelectedMarker `json:"marker,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScatterSelectedTextfont `json:"textfont,omitempty"`
}

// ScatterStream
type ScatterStream struct {

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

// ScatterTextfont Sets the text font.
type ScatterTextfont struct {

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

// ScatterUnselectedMarker
type ScatterUnselectedMarker struct {

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

// ScatterUnselectedTextfont
type ScatterUnselectedTextfont struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the text font color of unselected points, applied only when a selection exists.
	Color Color `json:"color,omitempty"`
}

// ScatterUnselected
type ScatterUnselected struct {

	// Marker
	// role: Object
	Marker *ScatterUnselectedMarker `json:"marker,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScatterUnselectedTextfont `json:"textfont,omitempty"`
}

// ScatterErrorXType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the square of the underlying data. If *data*, the bar lengths are set with data set `array`.
type ScatterErrorXType string

const (
	ScatterErrorXTypePercent  ScatterErrorXType = "percent"
	ScatterErrorXTypeConstant ScatterErrorXType = "constant"
	ScatterErrorXTypeSqrt     ScatterErrorXType = "sqrt"
	ScatterErrorXTypeData     ScatterErrorXType = "data"
)

// ScatterErrorYType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the square of the underlying data. If *data*, the bar lengths are set with data set `array`.
type ScatterErrorYType string

const (
	ScatterErrorYTypePercent  ScatterErrorYType = "percent"
	ScatterErrorYTypeConstant ScatterErrorYType = "constant"
	ScatterErrorYTypeSqrt     ScatterErrorYType = "sqrt"
	ScatterErrorYTypeData     ScatterErrorYType = "data"
)

// ScatterFill Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
type ScatterFill string

const (
	ScatterFillNone    ScatterFill = "none"
	ScatterFillTozeroy ScatterFill = "tozeroy"
	ScatterFillTozerox ScatterFill = "tozerox"
	ScatterFillTonexty ScatterFill = "tonexty"
	ScatterFillTonextx ScatterFill = "tonextx"
	ScatterFillToself  ScatterFill = "toself"
	ScatterFillTonext  ScatterFill = "tonext"
)

// ScatterGroupnorm Only relevant when `stackgroup` is used, and only the first `groupnorm` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Sets the normalization for the sum of this `stackgroup`. With *fraction*, the value of each trace at each location is divided by the sum of all trace values at that location. *percent* is the same but multiplied by 100 to show percentages. If there are multiple subplots, or multiple `stackgroup`s on one subplot, each will be normalized within its own set.
type ScatterGroupnorm string

const (
	ScatterGroupnormEmpty    ScatterGroupnorm = ""
	ScatterGroupnormFraction ScatterGroupnorm = "fraction"
	ScatterGroupnormPercent  ScatterGroupnorm = "percent"
)

// ScatterHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterHoverlabelAlign string

const (
	ScatterHoverlabelAlignLeft  ScatterHoverlabelAlign = "left"
	ScatterHoverlabelAlignRight ScatterHoverlabelAlign = "right"
	ScatterHoverlabelAlignAuto  ScatterHoverlabelAlign = "auto"
)

// ScatterLineShape Determines the line shape. With *spline* the lines are drawn using spline interpolation. The other available values correspond to step-wise line shapes.
type ScatterLineShape string

const (
	ScatterLineShapeLinear ScatterLineShape = "linear"
	ScatterLineShapeSpline ScatterLineShape = "spline"
	ScatterLineShapeHv     ScatterLineShape = "hv"
	ScatterLineShapeVh     ScatterLineShape = "vh"
	ScatterLineShapeHvh    ScatterLineShape = "hvh"
	ScatterLineShapeVhv    ScatterLineShape = "vhv"
)

// ScatterMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterMarkerColorbarExponentformat string

const (
	ScatterMarkerColorbarExponentformatNone  ScatterMarkerColorbarExponentformat = "none"
	ScatterMarkerColorbarExponentformatE1    ScatterMarkerColorbarExponentformat = "e"
	ScatterMarkerColorbarExponentformatE2    ScatterMarkerColorbarExponentformat = "E"
	ScatterMarkerColorbarExponentformatPower ScatterMarkerColorbarExponentformat = "power"
	ScatterMarkerColorbarExponentformatSi    ScatterMarkerColorbarExponentformat = "SI"
	ScatterMarkerColorbarExponentformatB     ScatterMarkerColorbarExponentformat = "B"
)

// ScatterMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterMarkerColorbarLenmode string

const (
	ScatterMarkerColorbarLenmodeFraction ScatterMarkerColorbarLenmode = "fraction"
	ScatterMarkerColorbarLenmodePixels   ScatterMarkerColorbarLenmode = "pixels"
)

// ScatterMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterMarkerColorbarShowexponent string

const (
	ScatterMarkerColorbarShowexponentAll   ScatterMarkerColorbarShowexponent = "all"
	ScatterMarkerColorbarShowexponentFirst ScatterMarkerColorbarShowexponent = "first"
	ScatterMarkerColorbarShowexponentLast  ScatterMarkerColorbarShowexponent = "last"
	ScatterMarkerColorbarShowexponentNone  ScatterMarkerColorbarShowexponent = "none"
)

// ScatterMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterMarkerColorbarShowtickprefix string

const (
	ScatterMarkerColorbarShowtickprefixAll   ScatterMarkerColorbarShowtickprefix = "all"
	ScatterMarkerColorbarShowtickprefixFirst ScatterMarkerColorbarShowtickprefix = "first"
	ScatterMarkerColorbarShowtickprefixLast  ScatterMarkerColorbarShowtickprefix = "last"
	ScatterMarkerColorbarShowtickprefixNone  ScatterMarkerColorbarShowtickprefix = "none"
)

// ScatterMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterMarkerColorbarShowticksuffix string

const (
	ScatterMarkerColorbarShowticksuffixAll   ScatterMarkerColorbarShowticksuffix = "all"
	ScatterMarkerColorbarShowticksuffixFirst ScatterMarkerColorbarShowticksuffix = "first"
	ScatterMarkerColorbarShowticksuffixLast  ScatterMarkerColorbarShowticksuffix = "last"
	ScatterMarkerColorbarShowticksuffixNone  ScatterMarkerColorbarShowticksuffix = "none"
)

// ScatterMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterMarkerColorbarThicknessmode string

const (
	ScatterMarkerColorbarThicknessmodeFraction ScatterMarkerColorbarThicknessmode = "fraction"
	ScatterMarkerColorbarThicknessmodePixels   ScatterMarkerColorbarThicknessmode = "pixels"
)

// ScatterMarkerColorbarTicklabelposition Determines where tick labels are drawn.
type ScatterMarkerColorbarTicklabelposition string

const (
	ScatterMarkerColorbarTicklabelpositionOutside       ScatterMarkerColorbarTicklabelposition = "outside"
	ScatterMarkerColorbarTicklabelpositionInside        ScatterMarkerColorbarTicklabelposition = "inside"
	ScatterMarkerColorbarTicklabelpositionOutsideTop    ScatterMarkerColorbarTicklabelposition = "outside top"
	ScatterMarkerColorbarTicklabelpositionInsideTop     ScatterMarkerColorbarTicklabelposition = "inside top"
	ScatterMarkerColorbarTicklabelpositionOutsideBottom ScatterMarkerColorbarTicklabelposition = "outside bottom"
	ScatterMarkerColorbarTicklabelpositionInsideBottom  ScatterMarkerColorbarTicklabelposition = "inside bottom"
)

// ScatterMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterMarkerColorbarTickmode string

const (
	ScatterMarkerColorbarTickmodeAuto   ScatterMarkerColorbarTickmode = "auto"
	ScatterMarkerColorbarTickmodeLinear ScatterMarkerColorbarTickmode = "linear"
	ScatterMarkerColorbarTickmodeArray  ScatterMarkerColorbarTickmode = "array"
)

// ScatterMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterMarkerColorbarTicks string

const (
	ScatterMarkerColorbarTicksOutside ScatterMarkerColorbarTicks = "outside"
	ScatterMarkerColorbarTicksInside  ScatterMarkerColorbarTicks = "inside"
	ScatterMarkerColorbarTicksEmpty   ScatterMarkerColorbarTicks = ""
)

// ScatterMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterMarkerColorbarTitleSide string

const (
	ScatterMarkerColorbarTitleSideRight  ScatterMarkerColorbarTitleSide = "right"
	ScatterMarkerColorbarTitleSideTop    ScatterMarkerColorbarTitleSide = "top"
	ScatterMarkerColorbarTitleSideBottom ScatterMarkerColorbarTitleSide = "bottom"
)

// ScatterMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterMarkerColorbarXanchor string

const (
	ScatterMarkerColorbarXanchorLeft   ScatterMarkerColorbarXanchor = "left"
	ScatterMarkerColorbarXanchorCenter ScatterMarkerColorbarXanchor = "center"
	ScatterMarkerColorbarXanchorRight  ScatterMarkerColorbarXanchor = "right"
)

// ScatterMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterMarkerColorbarYanchor string

const (
	ScatterMarkerColorbarYanchorTop    ScatterMarkerColorbarYanchor = "top"
	ScatterMarkerColorbarYanchorMiddle ScatterMarkerColorbarYanchor = "middle"
	ScatterMarkerColorbarYanchorBottom ScatterMarkerColorbarYanchor = "bottom"
)

// ScatterMarkerGradientType Sets the type of gradient used to fill the markers
type ScatterMarkerGradientType string

const (
	ScatterMarkerGradientTypeRadial     ScatterMarkerGradientType = "radial"
	ScatterMarkerGradientTypeHorizontal ScatterMarkerGradientType = "horizontal"
	ScatterMarkerGradientTypeVertical   ScatterMarkerGradientType = "vertical"
	ScatterMarkerGradientTypeNone       ScatterMarkerGradientType = "none"
)

// ScatterMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterMarkerSizemode string

const (
	ScatterMarkerSizemodeDiameter ScatterMarkerSizemode = "diameter"
	ScatterMarkerSizemodeArea     ScatterMarkerSizemode = "area"
)

// ScatterMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterMarkerSymbol interface{}

var (
	ScatterMarkerSymbolNumber0                 ScatterMarkerSymbol = 0
	ScatterMarkerSymbol0                       ScatterMarkerSymbol = "0"
	ScatterMarkerSymbolCircle                  ScatterMarkerSymbol = "circle"
	ScatterMarkerSymbolNumber100               ScatterMarkerSymbol = 100
	ScatterMarkerSymbol100                     ScatterMarkerSymbol = "100"
	ScatterMarkerSymbolCircleOpen              ScatterMarkerSymbol = "circle-open"
	ScatterMarkerSymbolNumber200               ScatterMarkerSymbol = 200
	ScatterMarkerSymbol200                     ScatterMarkerSymbol = "200"
	ScatterMarkerSymbolCircleDot               ScatterMarkerSymbol = "circle-dot"
	ScatterMarkerSymbolNumber300               ScatterMarkerSymbol = 300
	ScatterMarkerSymbol300                     ScatterMarkerSymbol = "300"
	ScatterMarkerSymbolCircleOpenDot           ScatterMarkerSymbol = "circle-open-dot"
	ScatterMarkerSymbolNumber1                 ScatterMarkerSymbol = 1
	ScatterMarkerSymbol1                       ScatterMarkerSymbol = "1"
	ScatterMarkerSymbolSquare                  ScatterMarkerSymbol = "square"
	ScatterMarkerSymbolNumber101               ScatterMarkerSymbol = 101
	ScatterMarkerSymbol101                     ScatterMarkerSymbol = "101"
	ScatterMarkerSymbolSquareOpen              ScatterMarkerSymbol = "square-open"
	ScatterMarkerSymbolNumber201               ScatterMarkerSymbol = 201
	ScatterMarkerSymbol201                     ScatterMarkerSymbol = "201"
	ScatterMarkerSymbolSquareDot               ScatterMarkerSymbol = "square-dot"
	ScatterMarkerSymbolNumber301               ScatterMarkerSymbol = 301
	ScatterMarkerSymbol301                     ScatterMarkerSymbol = "301"
	ScatterMarkerSymbolSquareOpenDot           ScatterMarkerSymbol = "square-open-dot"
	ScatterMarkerSymbolNumber2                 ScatterMarkerSymbol = 2
	ScatterMarkerSymbol2                       ScatterMarkerSymbol = "2"
	ScatterMarkerSymbolDiamond                 ScatterMarkerSymbol = "diamond"
	ScatterMarkerSymbolNumber102               ScatterMarkerSymbol = 102
	ScatterMarkerSymbol102                     ScatterMarkerSymbol = "102"
	ScatterMarkerSymbolDiamondOpen             ScatterMarkerSymbol = "diamond-open"
	ScatterMarkerSymbolNumber202               ScatterMarkerSymbol = 202
	ScatterMarkerSymbol202                     ScatterMarkerSymbol = "202"
	ScatterMarkerSymbolDiamondDot              ScatterMarkerSymbol = "diamond-dot"
	ScatterMarkerSymbolNumber302               ScatterMarkerSymbol = 302
	ScatterMarkerSymbol302                     ScatterMarkerSymbol = "302"
	ScatterMarkerSymbolDiamondOpenDot          ScatterMarkerSymbol = "diamond-open-dot"
	ScatterMarkerSymbolNumber3                 ScatterMarkerSymbol = 3
	ScatterMarkerSymbol3                       ScatterMarkerSymbol = "3"
	ScatterMarkerSymbolCross                   ScatterMarkerSymbol = "cross"
	ScatterMarkerSymbolNumber103               ScatterMarkerSymbol = 103
	ScatterMarkerSymbol103                     ScatterMarkerSymbol = "103"
	ScatterMarkerSymbolCrossOpen               ScatterMarkerSymbol = "cross-open"
	ScatterMarkerSymbolNumber203               ScatterMarkerSymbol = 203
	ScatterMarkerSymbol203                     ScatterMarkerSymbol = "203"
	ScatterMarkerSymbolCrossDot                ScatterMarkerSymbol = "cross-dot"
	ScatterMarkerSymbolNumber303               ScatterMarkerSymbol = 303
	ScatterMarkerSymbol303                     ScatterMarkerSymbol = "303"
	ScatterMarkerSymbolCrossOpenDot            ScatterMarkerSymbol = "cross-open-dot"
	ScatterMarkerSymbolNumber4                 ScatterMarkerSymbol = 4
	ScatterMarkerSymbol4                       ScatterMarkerSymbol = "4"
	ScatterMarkerSymbolX                       ScatterMarkerSymbol = "x"
	ScatterMarkerSymbolNumber104               ScatterMarkerSymbol = 104
	ScatterMarkerSymbol104                     ScatterMarkerSymbol = "104"
	ScatterMarkerSymbolXOpen                   ScatterMarkerSymbol = "x-open"
	ScatterMarkerSymbolNumber204               ScatterMarkerSymbol = 204
	ScatterMarkerSymbol204                     ScatterMarkerSymbol = "204"
	ScatterMarkerSymbolXDot                    ScatterMarkerSymbol = "x-dot"
	ScatterMarkerSymbolNumber304               ScatterMarkerSymbol = 304
	ScatterMarkerSymbol304                     ScatterMarkerSymbol = "304"
	ScatterMarkerSymbolXOpenDot                ScatterMarkerSymbol = "x-open-dot"
	ScatterMarkerSymbolNumber5                 ScatterMarkerSymbol = 5
	ScatterMarkerSymbol5                       ScatterMarkerSymbol = "5"
	ScatterMarkerSymbolTriangleUp              ScatterMarkerSymbol = "triangle-up"
	ScatterMarkerSymbolNumber105               ScatterMarkerSymbol = 105
	ScatterMarkerSymbol105                     ScatterMarkerSymbol = "105"
	ScatterMarkerSymbolTriangleUpOpen          ScatterMarkerSymbol = "triangle-up-open"
	ScatterMarkerSymbolNumber205               ScatterMarkerSymbol = 205
	ScatterMarkerSymbol205                     ScatterMarkerSymbol = "205"
	ScatterMarkerSymbolTriangleUpDot           ScatterMarkerSymbol = "triangle-up-dot"
	ScatterMarkerSymbolNumber305               ScatterMarkerSymbol = 305
	ScatterMarkerSymbol305                     ScatterMarkerSymbol = "305"
	ScatterMarkerSymbolTriangleUpOpenDot       ScatterMarkerSymbol = "triangle-up-open-dot"
	ScatterMarkerSymbolNumber6                 ScatterMarkerSymbol = 6
	ScatterMarkerSymbol6                       ScatterMarkerSymbol = "6"
	ScatterMarkerSymbolTriangleDown            ScatterMarkerSymbol = "triangle-down"
	ScatterMarkerSymbolNumber106               ScatterMarkerSymbol = 106
	ScatterMarkerSymbol106                     ScatterMarkerSymbol = "106"
	ScatterMarkerSymbolTriangleDownOpen        ScatterMarkerSymbol = "triangle-down-open"
	ScatterMarkerSymbolNumber206               ScatterMarkerSymbol = 206
	ScatterMarkerSymbol206                     ScatterMarkerSymbol = "206"
	ScatterMarkerSymbolTriangleDownDot         ScatterMarkerSymbol = "triangle-down-dot"
	ScatterMarkerSymbolNumber306               ScatterMarkerSymbol = 306
	ScatterMarkerSymbol306                     ScatterMarkerSymbol = "306"
	ScatterMarkerSymbolTriangleDownOpenDot     ScatterMarkerSymbol = "triangle-down-open-dot"
	ScatterMarkerSymbolNumber7                 ScatterMarkerSymbol = 7
	ScatterMarkerSymbol7                       ScatterMarkerSymbol = "7"
	ScatterMarkerSymbolTriangleLeft            ScatterMarkerSymbol = "triangle-left"
	ScatterMarkerSymbolNumber107               ScatterMarkerSymbol = 107
	ScatterMarkerSymbol107                     ScatterMarkerSymbol = "107"
	ScatterMarkerSymbolTriangleLeftOpen        ScatterMarkerSymbol = "triangle-left-open"
	ScatterMarkerSymbolNumber207               ScatterMarkerSymbol = 207
	ScatterMarkerSymbol207                     ScatterMarkerSymbol = "207"
	ScatterMarkerSymbolTriangleLeftDot         ScatterMarkerSymbol = "triangle-left-dot"
	ScatterMarkerSymbolNumber307               ScatterMarkerSymbol = 307
	ScatterMarkerSymbol307                     ScatterMarkerSymbol = "307"
	ScatterMarkerSymbolTriangleLeftOpenDot     ScatterMarkerSymbol = "triangle-left-open-dot"
	ScatterMarkerSymbolNumber8                 ScatterMarkerSymbol = 8
	ScatterMarkerSymbol8                       ScatterMarkerSymbol = "8"
	ScatterMarkerSymbolTriangleRight           ScatterMarkerSymbol = "triangle-right"
	ScatterMarkerSymbolNumber108               ScatterMarkerSymbol = 108
	ScatterMarkerSymbol108                     ScatterMarkerSymbol = "108"
	ScatterMarkerSymbolTriangleRightOpen       ScatterMarkerSymbol = "triangle-right-open"
	ScatterMarkerSymbolNumber208               ScatterMarkerSymbol = 208
	ScatterMarkerSymbol208                     ScatterMarkerSymbol = "208"
	ScatterMarkerSymbolTriangleRightDot        ScatterMarkerSymbol = "triangle-right-dot"
	ScatterMarkerSymbolNumber308               ScatterMarkerSymbol = 308
	ScatterMarkerSymbol308                     ScatterMarkerSymbol = "308"
	ScatterMarkerSymbolTriangleRightOpenDot    ScatterMarkerSymbol = "triangle-right-open-dot"
	ScatterMarkerSymbolNumber9                 ScatterMarkerSymbol = 9
	ScatterMarkerSymbol9                       ScatterMarkerSymbol = "9"
	ScatterMarkerSymbolTriangleNe              ScatterMarkerSymbol = "triangle-ne"
	ScatterMarkerSymbolNumber109               ScatterMarkerSymbol = 109
	ScatterMarkerSymbol109                     ScatterMarkerSymbol = "109"
	ScatterMarkerSymbolTriangleNeOpen          ScatterMarkerSymbol = "triangle-ne-open"
	ScatterMarkerSymbolNumber209               ScatterMarkerSymbol = 209
	ScatterMarkerSymbol209                     ScatterMarkerSymbol = "209"
	ScatterMarkerSymbolTriangleNeDot           ScatterMarkerSymbol = "triangle-ne-dot"
	ScatterMarkerSymbolNumber309               ScatterMarkerSymbol = 309
	ScatterMarkerSymbol309                     ScatterMarkerSymbol = "309"
	ScatterMarkerSymbolTriangleNeOpenDot       ScatterMarkerSymbol = "triangle-ne-open-dot"
	ScatterMarkerSymbolNumber10                ScatterMarkerSymbol = 10
	ScatterMarkerSymbol10                      ScatterMarkerSymbol = "10"
	ScatterMarkerSymbolTriangleSe              ScatterMarkerSymbol = "triangle-se"
	ScatterMarkerSymbolNumber110               ScatterMarkerSymbol = 110
	ScatterMarkerSymbol110                     ScatterMarkerSymbol = "110"
	ScatterMarkerSymbolTriangleSeOpen          ScatterMarkerSymbol = "triangle-se-open"
	ScatterMarkerSymbolNumber210               ScatterMarkerSymbol = 210
	ScatterMarkerSymbol210                     ScatterMarkerSymbol = "210"
	ScatterMarkerSymbolTriangleSeDot           ScatterMarkerSymbol = "triangle-se-dot"
	ScatterMarkerSymbolNumber310               ScatterMarkerSymbol = 310
	ScatterMarkerSymbol310                     ScatterMarkerSymbol = "310"
	ScatterMarkerSymbolTriangleSeOpenDot       ScatterMarkerSymbol = "triangle-se-open-dot"
	ScatterMarkerSymbolNumber11                ScatterMarkerSymbol = 11
	ScatterMarkerSymbol11                      ScatterMarkerSymbol = "11"
	ScatterMarkerSymbolTriangleSw              ScatterMarkerSymbol = "triangle-sw"
	ScatterMarkerSymbolNumber111               ScatterMarkerSymbol = 111
	ScatterMarkerSymbol111                     ScatterMarkerSymbol = "111"
	ScatterMarkerSymbolTriangleSwOpen          ScatterMarkerSymbol = "triangle-sw-open"
	ScatterMarkerSymbolNumber211               ScatterMarkerSymbol = 211
	ScatterMarkerSymbol211                     ScatterMarkerSymbol = "211"
	ScatterMarkerSymbolTriangleSwDot           ScatterMarkerSymbol = "triangle-sw-dot"
	ScatterMarkerSymbolNumber311               ScatterMarkerSymbol = 311
	ScatterMarkerSymbol311                     ScatterMarkerSymbol = "311"
	ScatterMarkerSymbolTriangleSwOpenDot       ScatterMarkerSymbol = "triangle-sw-open-dot"
	ScatterMarkerSymbolNumber12                ScatterMarkerSymbol = 12
	ScatterMarkerSymbol12                      ScatterMarkerSymbol = "12"
	ScatterMarkerSymbolTriangleNw              ScatterMarkerSymbol = "triangle-nw"
	ScatterMarkerSymbolNumber112               ScatterMarkerSymbol = 112
	ScatterMarkerSymbol112                     ScatterMarkerSymbol = "112"
	ScatterMarkerSymbolTriangleNwOpen          ScatterMarkerSymbol = "triangle-nw-open"
	ScatterMarkerSymbolNumber212               ScatterMarkerSymbol = 212
	ScatterMarkerSymbol212                     ScatterMarkerSymbol = "212"
	ScatterMarkerSymbolTriangleNwDot           ScatterMarkerSymbol = "triangle-nw-dot"
	ScatterMarkerSymbolNumber312               ScatterMarkerSymbol = 312
	ScatterMarkerSymbol312                     ScatterMarkerSymbol = "312"
	ScatterMarkerSymbolTriangleNwOpenDot       ScatterMarkerSymbol = "triangle-nw-open-dot"
	ScatterMarkerSymbolNumber13                ScatterMarkerSymbol = 13
	ScatterMarkerSymbol13                      ScatterMarkerSymbol = "13"
	ScatterMarkerSymbolPentagon                ScatterMarkerSymbol = "pentagon"
	ScatterMarkerSymbolNumber113               ScatterMarkerSymbol = 113
	ScatterMarkerSymbol113                     ScatterMarkerSymbol = "113"
	ScatterMarkerSymbolPentagonOpen            ScatterMarkerSymbol = "pentagon-open"
	ScatterMarkerSymbolNumber213               ScatterMarkerSymbol = 213
	ScatterMarkerSymbol213                     ScatterMarkerSymbol = "213"
	ScatterMarkerSymbolPentagonDot             ScatterMarkerSymbol = "pentagon-dot"
	ScatterMarkerSymbolNumber313               ScatterMarkerSymbol = 313
	ScatterMarkerSymbol313                     ScatterMarkerSymbol = "313"
	ScatterMarkerSymbolPentagonOpenDot         ScatterMarkerSymbol = "pentagon-open-dot"
	ScatterMarkerSymbolNumber14                ScatterMarkerSymbol = 14
	ScatterMarkerSymbol14                      ScatterMarkerSymbol = "14"
	ScatterMarkerSymbolHexagon                 ScatterMarkerSymbol = "hexagon"
	ScatterMarkerSymbolNumber114               ScatterMarkerSymbol = 114
	ScatterMarkerSymbol114                     ScatterMarkerSymbol = "114"
	ScatterMarkerSymbolHexagonOpen             ScatterMarkerSymbol = "hexagon-open"
	ScatterMarkerSymbolNumber214               ScatterMarkerSymbol = 214
	ScatterMarkerSymbol214                     ScatterMarkerSymbol = "214"
	ScatterMarkerSymbolHexagonDot              ScatterMarkerSymbol = "hexagon-dot"
	ScatterMarkerSymbolNumber314               ScatterMarkerSymbol = 314
	ScatterMarkerSymbol314                     ScatterMarkerSymbol = "314"
	ScatterMarkerSymbolHexagonOpenDot          ScatterMarkerSymbol = "hexagon-open-dot"
	ScatterMarkerSymbolNumber15                ScatterMarkerSymbol = 15
	ScatterMarkerSymbol15                      ScatterMarkerSymbol = "15"
	ScatterMarkerSymbolHexagon2                ScatterMarkerSymbol = "hexagon2"
	ScatterMarkerSymbolNumber115               ScatterMarkerSymbol = 115
	ScatterMarkerSymbol115                     ScatterMarkerSymbol = "115"
	ScatterMarkerSymbolHexagon2Open            ScatterMarkerSymbol = "hexagon2-open"
	ScatterMarkerSymbolNumber215               ScatterMarkerSymbol = 215
	ScatterMarkerSymbol215                     ScatterMarkerSymbol = "215"
	ScatterMarkerSymbolHexagon2Dot             ScatterMarkerSymbol = "hexagon2-dot"
	ScatterMarkerSymbolNumber315               ScatterMarkerSymbol = 315
	ScatterMarkerSymbol315                     ScatterMarkerSymbol = "315"
	ScatterMarkerSymbolHexagon2OpenDot         ScatterMarkerSymbol = "hexagon2-open-dot"
	ScatterMarkerSymbolNumber16                ScatterMarkerSymbol = 16
	ScatterMarkerSymbol16                      ScatterMarkerSymbol = "16"
	ScatterMarkerSymbolOctagon                 ScatterMarkerSymbol = "octagon"
	ScatterMarkerSymbolNumber116               ScatterMarkerSymbol = 116
	ScatterMarkerSymbol116                     ScatterMarkerSymbol = "116"
	ScatterMarkerSymbolOctagonOpen             ScatterMarkerSymbol = "octagon-open"
	ScatterMarkerSymbolNumber216               ScatterMarkerSymbol = 216
	ScatterMarkerSymbol216                     ScatterMarkerSymbol = "216"
	ScatterMarkerSymbolOctagonDot              ScatterMarkerSymbol = "octagon-dot"
	ScatterMarkerSymbolNumber316               ScatterMarkerSymbol = 316
	ScatterMarkerSymbol316                     ScatterMarkerSymbol = "316"
	ScatterMarkerSymbolOctagonOpenDot          ScatterMarkerSymbol = "octagon-open-dot"
	ScatterMarkerSymbolNumber17                ScatterMarkerSymbol = 17
	ScatterMarkerSymbol17                      ScatterMarkerSymbol = "17"
	ScatterMarkerSymbolStar                    ScatterMarkerSymbol = "star"
	ScatterMarkerSymbolNumber117               ScatterMarkerSymbol = 117
	ScatterMarkerSymbol117                     ScatterMarkerSymbol = "117"
	ScatterMarkerSymbolStarOpen                ScatterMarkerSymbol = "star-open"
	ScatterMarkerSymbolNumber217               ScatterMarkerSymbol = 217
	ScatterMarkerSymbol217                     ScatterMarkerSymbol = "217"
	ScatterMarkerSymbolStarDot                 ScatterMarkerSymbol = "star-dot"
	ScatterMarkerSymbolNumber317               ScatterMarkerSymbol = 317
	ScatterMarkerSymbol317                     ScatterMarkerSymbol = "317"
	ScatterMarkerSymbolStarOpenDot             ScatterMarkerSymbol = "star-open-dot"
	ScatterMarkerSymbolNumber18                ScatterMarkerSymbol = 18
	ScatterMarkerSymbol18                      ScatterMarkerSymbol = "18"
	ScatterMarkerSymbolHexagram                ScatterMarkerSymbol = "hexagram"
	ScatterMarkerSymbolNumber118               ScatterMarkerSymbol = 118
	ScatterMarkerSymbol118                     ScatterMarkerSymbol = "118"
	ScatterMarkerSymbolHexagramOpen            ScatterMarkerSymbol = "hexagram-open"
	ScatterMarkerSymbolNumber218               ScatterMarkerSymbol = 218
	ScatterMarkerSymbol218                     ScatterMarkerSymbol = "218"
	ScatterMarkerSymbolHexagramDot             ScatterMarkerSymbol = "hexagram-dot"
	ScatterMarkerSymbolNumber318               ScatterMarkerSymbol = 318
	ScatterMarkerSymbol318                     ScatterMarkerSymbol = "318"
	ScatterMarkerSymbolHexagramOpenDot         ScatterMarkerSymbol = "hexagram-open-dot"
	ScatterMarkerSymbolNumber19                ScatterMarkerSymbol = 19
	ScatterMarkerSymbol19                      ScatterMarkerSymbol = "19"
	ScatterMarkerSymbolStarTriangleUp          ScatterMarkerSymbol = "star-triangle-up"
	ScatterMarkerSymbolNumber119               ScatterMarkerSymbol = 119
	ScatterMarkerSymbol119                     ScatterMarkerSymbol = "119"
	ScatterMarkerSymbolStarTriangleUpOpen      ScatterMarkerSymbol = "star-triangle-up-open"
	ScatterMarkerSymbolNumber219               ScatterMarkerSymbol = 219
	ScatterMarkerSymbol219                     ScatterMarkerSymbol = "219"
	ScatterMarkerSymbolStarTriangleUpDot       ScatterMarkerSymbol = "star-triangle-up-dot"
	ScatterMarkerSymbolNumber319               ScatterMarkerSymbol = 319
	ScatterMarkerSymbol319                     ScatterMarkerSymbol = "319"
	ScatterMarkerSymbolStarTriangleUpOpenDot   ScatterMarkerSymbol = "star-triangle-up-open-dot"
	ScatterMarkerSymbolNumber20                ScatterMarkerSymbol = 20
	ScatterMarkerSymbol20                      ScatterMarkerSymbol = "20"
	ScatterMarkerSymbolStarTriangleDown        ScatterMarkerSymbol = "star-triangle-down"
	ScatterMarkerSymbolNumber120               ScatterMarkerSymbol = 120
	ScatterMarkerSymbol120                     ScatterMarkerSymbol = "120"
	ScatterMarkerSymbolStarTriangleDownOpen    ScatterMarkerSymbol = "star-triangle-down-open"
	ScatterMarkerSymbolNumber220               ScatterMarkerSymbol = 220
	ScatterMarkerSymbol220                     ScatterMarkerSymbol = "220"
	ScatterMarkerSymbolStarTriangleDownDot     ScatterMarkerSymbol = "star-triangle-down-dot"
	ScatterMarkerSymbolNumber320               ScatterMarkerSymbol = 320
	ScatterMarkerSymbol320                     ScatterMarkerSymbol = "320"
	ScatterMarkerSymbolStarTriangleDownOpenDot ScatterMarkerSymbol = "star-triangle-down-open-dot"
	ScatterMarkerSymbolNumber21                ScatterMarkerSymbol = 21
	ScatterMarkerSymbol21                      ScatterMarkerSymbol = "21"
	ScatterMarkerSymbolStarSquare              ScatterMarkerSymbol = "star-square"
	ScatterMarkerSymbolNumber121               ScatterMarkerSymbol = 121
	ScatterMarkerSymbol121                     ScatterMarkerSymbol = "121"
	ScatterMarkerSymbolStarSquareOpen          ScatterMarkerSymbol = "star-square-open"
	ScatterMarkerSymbolNumber221               ScatterMarkerSymbol = 221
	ScatterMarkerSymbol221                     ScatterMarkerSymbol = "221"
	ScatterMarkerSymbolStarSquareDot           ScatterMarkerSymbol = "star-square-dot"
	ScatterMarkerSymbolNumber321               ScatterMarkerSymbol = 321
	ScatterMarkerSymbol321                     ScatterMarkerSymbol = "321"
	ScatterMarkerSymbolStarSquareOpenDot       ScatterMarkerSymbol = "star-square-open-dot"
	ScatterMarkerSymbolNumber22                ScatterMarkerSymbol = 22
	ScatterMarkerSymbol22                      ScatterMarkerSymbol = "22"
	ScatterMarkerSymbolStarDiamond             ScatterMarkerSymbol = "star-diamond"
	ScatterMarkerSymbolNumber122               ScatterMarkerSymbol = 122
	ScatterMarkerSymbol122                     ScatterMarkerSymbol = "122"
	ScatterMarkerSymbolStarDiamondOpen         ScatterMarkerSymbol = "star-diamond-open"
	ScatterMarkerSymbolNumber222               ScatterMarkerSymbol = 222
	ScatterMarkerSymbol222                     ScatterMarkerSymbol = "222"
	ScatterMarkerSymbolStarDiamondDot          ScatterMarkerSymbol = "star-diamond-dot"
	ScatterMarkerSymbolNumber322               ScatterMarkerSymbol = 322
	ScatterMarkerSymbol322                     ScatterMarkerSymbol = "322"
	ScatterMarkerSymbolStarDiamondOpenDot      ScatterMarkerSymbol = "star-diamond-open-dot"
	ScatterMarkerSymbolNumber23                ScatterMarkerSymbol = 23
	ScatterMarkerSymbol23                      ScatterMarkerSymbol = "23"
	ScatterMarkerSymbolDiamondTall             ScatterMarkerSymbol = "diamond-tall"
	ScatterMarkerSymbolNumber123               ScatterMarkerSymbol = 123
	ScatterMarkerSymbol123                     ScatterMarkerSymbol = "123"
	ScatterMarkerSymbolDiamondTallOpen         ScatterMarkerSymbol = "diamond-tall-open"
	ScatterMarkerSymbolNumber223               ScatterMarkerSymbol = 223
	ScatterMarkerSymbol223                     ScatterMarkerSymbol = "223"
	ScatterMarkerSymbolDiamondTallDot          ScatterMarkerSymbol = "diamond-tall-dot"
	ScatterMarkerSymbolNumber323               ScatterMarkerSymbol = 323
	ScatterMarkerSymbol323                     ScatterMarkerSymbol = "323"
	ScatterMarkerSymbolDiamondTallOpenDot      ScatterMarkerSymbol = "diamond-tall-open-dot"
	ScatterMarkerSymbolNumber24                ScatterMarkerSymbol = 24
	ScatterMarkerSymbol24                      ScatterMarkerSymbol = "24"
	ScatterMarkerSymbolDiamondWide             ScatterMarkerSymbol = "diamond-wide"
	ScatterMarkerSymbolNumber124               ScatterMarkerSymbol = 124
	ScatterMarkerSymbol124                     ScatterMarkerSymbol = "124"
	ScatterMarkerSymbolDiamondWideOpen         ScatterMarkerSymbol = "diamond-wide-open"
	ScatterMarkerSymbolNumber224               ScatterMarkerSymbol = 224
	ScatterMarkerSymbol224                     ScatterMarkerSymbol = "224"
	ScatterMarkerSymbolDiamondWideDot          ScatterMarkerSymbol = "diamond-wide-dot"
	ScatterMarkerSymbolNumber324               ScatterMarkerSymbol = 324
	ScatterMarkerSymbol324                     ScatterMarkerSymbol = "324"
	ScatterMarkerSymbolDiamondWideOpenDot      ScatterMarkerSymbol = "diamond-wide-open-dot"
	ScatterMarkerSymbolNumber25                ScatterMarkerSymbol = 25
	ScatterMarkerSymbol25                      ScatterMarkerSymbol = "25"
	ScatterMarkerSymbolHourglass               ScatterMarkerSymbol = "hourglass"
	ScatterMarkerSymbolNumber125               ScatterMarkerSymbol = 125
	ScatterMarkerSymbol125                     ScatterMarkerSymbol = "125"
	ScatterMarkerSymbolHourglassOpen           ScatterMarkerSymbol = "hourglass-open"
	ScatterMarkerSymbolNumber26                ScatterMarkerSymbol = 26
	ScatterMarkerSymbol26                      ScatterMarkerSymbol = "26"
	ScatterMarkerSymbolBowtie                  ScatterMarkerSymbol = "bowtie"
	ScatterMarkerSymbolNumber126               ScatterMarkerSymbol = 126
	ScatterMarkerSymbol126                     ScatterMarkerSymbol = "126"
	ScatterMarkerSymbolBowtieOpen              ScatterMarkerSymbol = "bowtie-open"
	ScatterMarkerSymbolNumber27                ScatterMarkerSymbol = 27
	ScatterMarkerSymbol27                      ScatterMarkerSymbol = "27"
	ScatterMarkerSymbolCircleCross             ScatterMarkerSymbol = "circle-cross"
	ScatterMarkerSymbolNumber127               ScatterMarkerSymbol = 127
	ScatterMarkerSymbol127                     ScatterMarkerSymbol = "127"
	ScatterMarkerSymbolCircleCrossOpen         ScatterMarkerSymbol = "circle-cross-open"
	ScatterMarkerSymbolNumber28                ScatterMarkerSymbol = 28
	ScatterMarkerSymbol28                      ScatterMarkerSymbol = "28"
	ScatterMarkerSymbolCircleX                 ScatterMarkerSymbol = "circle-x"
	ScatterMarkerSymbolNumber128               ScatterMarkerSymbol = 128
	ScatterMarkerSymbol128                     ScatterMarkerSymbol = "128"
	ScatterMarkerSymbolCircleXOpen             ScatterMarkerSymbol = "circle-x-open"
	ScatterMarkerSymbolNumber29                ScatterMarkerSymbol = 29
	ScatterMarkerSymbol29                      ScatterMarkerSymbol = "29"
	ScatterMarkerSymbolSquareCross             ScatterMarkerSymbol = "square-cross"
	ScatterMarkerSymbolNumber129               ScatterMarkerSymbol = 129
	ScatterMarkerSymbol129                     ScatterMarkerSymbol = "129"
	ScatterMarkerSymbolSquareCrossOpen         ScatterMarkerSymbol = "square-cross-open"
	ScatterMarkerSymbolNumber30                ScatterMarkerSymbol = 30
	ScatterMarkerSymbol30                      ScatterMarkerSymbol = "30"
	ScatterMarkerSymbolSquareX                 ScatterMarkerSymbol = "square-x"
	ScatterMarkerSymbolNumber130               ScatterMarkerSymbol = 130
	ScatterMarkerSymbol130                     ScatterMarkerSymbol = "130"
	ScatterMarkerSymbolSquareXOpen             ScatterMarkerSymbol = "square-x-open"
	ScatterMarkerSymbolNumber31                ScatterMarkerSymbol = 31
	ScatterMarkerSymbol31                      ScatterMarkerSymbol = "31"
	ScatterMarkerSymbolDiamondCross            ScatterMarkerSymbol = "diamond-cross"
	ScatterMarkerSymbolNumber131               ScatterMarkerSymbol = 131
	ScatterMarkerSymbol131                     ScatterMarkerSymbol = "131"
	ScatterMarkerSymbolDiamondCrossOpen        ScatterMarkerSymbol = "diamond-cross-open"
	ScatterMarkerSymbolNumber32                ScatterMarkerSymbol = 32
	ScatterMarkerSymbol32                      ScatterMarkerSymbol = "32"
	ScatterMarkerSymbolDiamondX                ScatterMarkerSymbol = "diamond-x"
	ScatterMarkerSymbolNumber132               ScatterMarkerSymbol = 132
	ScatterMarkerSymbol132                     ScatterMarkerSymbol = "132"
	ScatterMarkerSymbolDiamondXOpen            ScatterMarkerSymbol = "diamond-x-open"
	ScatterMarkerSymbolNumber33                ScatterMarkerSymbol = 33
	ScatterMarkerSymbol33                      ScatterMarkerSymbol = "33"
	ScatterMarkerSymbolCrossThin               ScatterMarkerSymbol = "cross-thin"
	ScatterMarkerSymbolNumber133               ScatterMarkerSymbol = 133
	ScatterMarkerSymbol133                     ScatterMarkerSymbol = "133"
	ScatterMarkerSymbolCrossThinOpen           ScatterMarkerSymbol = "cross-thin-open"
	ScatterMarkerSymbolNumber34                ScatterMarkerSymbol = 34
	ScatterMarkerSymbol34                      ScatterMarkerSymbol = "34"
	ScatterMarkerSymbolXThin                   ScatterMarkerSymbol = "x-thin"
	ScatterMarkerSymbolNumber134               ScatterMarkerSymbol = 134
	ScatterMarkerSymbol134                     ScatterMarkerSymbol = "134"
	ScatterMarkerSymbolXThinOpen               ScatterMarkerSymbol = "x-thin-open"
	ScatterMarkerSymbolNumber35                ScatterMarkerSymbol = 35
	ScatterMarkerSymbol35                      ScatterMarkerSymbol = "35"
	ScatterMarkerSymbolAsterisk                ScatterMarkerSymbol = "asterisk"
	ScatterMarkerSymbolNumber135               ScatterMarkerSymbol = 135
	ScatterMarkerSymbol135                     ScatterMarkerSymbol = "135"
	ScatterMarkerSymbolAsteriskOpen            ScatterMarkerSymbol = "asterisk-open"
	ScatterMarkerSymbolNumber36                ScatterMarkerSymbol = 36
	ScatterMarkerSymbol36                      ScatterMarkerSymbol = "36"
	ScatterMarkerSymbolHash                    ScatterMarkerSymbol = "hash"
	ScatterMarkerSymbolNumber136               ScatterMarkerSymbol = 136
	ScatterMarkerSymbol136                     ScatterMarkerSymbol = "136"
	ScatterMarkerSymbolHashOpen                ScatterMarkerSymbol = "hash-open"
	ScatterMarkerSymbolNumber236               ScatterMarkerSymbol = 236
	ScatterMarkerSymbol236                     ScatterMarkerSymbol = "236"
	ScatterMarkerSymbolHashDot                 ScatterMarkerSymbol = "hash-dot"
	ScatterMarkerSymbolNumber336               ScatterMarkerSymbol = 336
	ScatterMarkerSymbol336                     ScatterMarkerSymbol = "336"
	ScatterMarkerSymbolHashOpenDot             ScatterMarkerSymbol = "hash-open-dot"
	ScatterMarkerSymbolNumber37                ScatterMarkerSymbol = 37
	ScatterMarkerSymbol37                      ScatterMarkerSymbol = "37"
	ScatterMarkerSymbolYUp                     ScatterMarkerSymbol = "y-up"
	ScatterMarkerSymbolNumber137               ScatterMarkerSymbol = 137
	ScatterMarkerSymbol137                     ScatterMarkerSymbol = "137"
	ScatterMarkerSymbolYUpOpen                 ScatterMarkerSymbol = "y-up-open"
	ScatterMarkerSymbolNumber38                ScatterMarkerSymbol = 38
	ScatterMarkerSymbol38                      ScatterMarkerSymbol = "38"
	ScatterMarkerSymbolYDown                   ScatterMarkerSymbol = "y-down"
	ScatterMarkerSymbolNumber138               ScatterMarkerSymbol = 138
	ScatterMarkerSymbol138                     ScatterMarkerSymbol = "138"
	ScatterMarkerSymbolYDownOpen               ScatterMarkerSymbol = "y-down-open"
	ScatterMarkerSymbolNumber39                ScatterMarkerSymbol = 39
	ScatterMarkerSymbol39                      ScatterMarkerSymbol = "39"
	ScatterMarkerSymbolYLeft                   ScatterMarkerSymbol = "y-left"
	ScatterMarkerSymbolNumber139               ScatterMarkerSymbol = 139
	ScatterMarkerSymbol139                     ScatterMarkerSymbol = "139"
	ScatterMarkerSymbolYLeftOpen               ScatterMarkerSymbol = "y-left-open"
	ScatterMarkerSymbolNumber40                ScatterMarkerSymbol = 40
	ScatterMarkerSymbol40                      ScatterMarkerSymbol = "40"
	ScatterMarkerSymbolYRight                  ScatterMarkerSymbol = "y-right"
	ScatterMarkerSymbolNumber140               ScatterMarkerSymbol = 140
	ScatterMarkerSymbol140                     ScatterMarkerSymbol = "140"
	ScatterMarkerSymbolYRightOpen              ScatterMarkerSymbol = "y-right-open"
	ScatterMarkerSymbolNumber41                ScatterMarkerSymbol = 41
	ScatterMarkerSymbol41                      ScatterMarkerSymbol = "41"
	ScatterMarkerSymbolLineEw                  ScatterMarkerSymbol = "line-ew"
	ScatterMarkerSymbolNumber141               ScatterMarkerSymbol = 141
	ScatterMarkerSymbol141                     ScatterMarkerSymbol = "141"
	ScatterMarkerSymbolLineEwOpen              ScatterMarkerSymbol = "line-ew-open"
	ScatterMarkerSymbolNumber42                ScatterMarkerSymbol = 42
	ScatterMarkerSymbol42                      ScatterMarkerSymbol = "42"
	ScatterMarkerSymbolLineNs                  ScatterMarkerSymbol = "line-ns"
	ScatterMarkerSymbolNumber142               ScatterMarkerSymbol = 142
	ScatterMarkerSymbol142                     ScatterMarkerSymbol = "142"
	ScatterMarkerSymbolLineNsOpen              ScatterMarkerSymbol = "line-ns-open"
	ScatterMarkerSymbolNumber43                ScatterMarkerSymbol = 43
	ScatterMarkerSymbol43                      ScatterMarkerSymbol = "43"
	ScatterMarkerSymbolLineNe                  ScatterMarkerSymbol = "line-ne"
	ScatterMarkerSymbolNumber143               ScatterMarkerSymbol = 143
	ScatterMarkerSymbol143                     ScatterMarkerSymbol = "143"
	ScatterMarkerSymbolLineNeOpen              ScatterMarkerSymbol = "line-ne-open"
	ScatterMarkerSymbolNumber44                ScatterMarkerSymbol = 44
	ScatterMarkerSymbol44                      ScatterMarkerSymbol = "44"
	ScatterMarkerSymbolLineNw                  ScatterMarkerSymbol = "line-nw"
	ScatterMarkerSymbolNumber144               ScatterMarkerSymbol = 144
	ScatterMarkerSymbol144                     ScatterMarkerSymbol = "144"
	ScatterMarkerSymbolLineNwOpen              ScatterMarkerSymbol = "line-nw-open"
	ScatterMarkerSymbolNumber45                ScatterMarkerSymbol = 45
	ScatterMarkerSymbol45                      ScatterMarkerSymbol = "45"
	ScatterMarkerSymbolArrowUp                 ScatterMarkerSymbol = "arrow-up"
	ScatterMarkerSymbolNumber145               ScatterMarkerSymbol = 145
	ScatterMarkerSymbol145                     ScatterMarkerSymbol = "145"
	ScatterMarkerSymbolArrowUpOpen             ScatterMarkerSymbol = "arrow-up-open"
	ScatterMarkerSymbolNumber46                ScatterMarkerSymbol = 46
	ScatterMarkerSymbol46                      ScatterMarkerSymbol = "46"
	ScatterMarkerSymbolArrowDown               ScatterMarkerSymbol = "arrow-down"
	ScatterMarkerSymbolNumber146               ScatterMarkerSymbol = 146
	ScatterMarkerSymbol146                     ScatterMarkerSymbol = "146"
	ScatterMarkerSymbolArrowDownOpen           ScatterMarkerSymbol = "arrow-down-open"
	ScatterMarkerSymbolNumber47                ScatterMarkerSymbol = 47
	ScatterMarkerSymbol47                      ScatterMarkerSymbol = "47"
	ScatterMarkerSymbolArrowLeft               ScatterMarkerSymbol = "arrow-left"
	ScatterMarkerSymbolNumber147               ScatterMarkerSymbol = 147
	ScatterMarkerSymbol147                     ScatterMarkerSymbol = "147"
	ScatterMarkerSymbolArrowLeftOpen           ScatterMarkerSymbol = "arrow-left-open"
	ScatterMarkerSymbolNumber48                ScatterMarkerSymbol = 48
	ScatterMarkerSymbol48                      ScatterMarkerSymbol = "48"
	ScatterMarkerSymbolArrowRight              ScatterMarkerSymbol = "arrow-right"
	ScatterMarkerSymbolNumber148               ScatterMarkerSymbol = 148
	ScatterMarkerSymbol148                     ScatterMarkerSymbol = "148"
	ScatterMarkerSymbolArrowRightOpen          ScatterMarkerSymbol = "arrow-right-open"
	ScatterMarkerSymbolNumber49                ScatterMarkerSymbol = 49
	ScatterMarkerSymbol49                      ScatterMarkerSymbol = "49"
	ScatterMarkerSymbolArrowBarUp              ScatterMarkerSymbol = "arrow-bar-up"
	ScatterMarkerSymbolNumber149               ScatterMarkerSymbol = 149
	ScatterMarkerSymbol149                     ScatterMarkerSymbol = "149"
	ScatterMarkerSymbolArrowBarUpOpen          ScatterMarkerSymbol = "arrow-bar-up-open"
	ScatterMarkerSymbolNumber50                ScatterMarkerSymbol = 50
	ScatterMarkerSymbol50                      ScatterMarkerSymbol = "50"
	ScatterMarkerSymbolArrowBarDown            ScatterMarkerSymbol = "arrow-bar-down"
	ScatterMarkerSymbolNumber150               ScatterMarkerSymbol = 150
	ScatterMarkerSymbol150                     ScatterMarkerSymbol = "150"
	ScatterMarkerSymbolArrowBarDownOpen        ScatterMarkerSymbol = "arrow-bar-down-open"
	ScatterMarkerSymbolNumber51                ScatterMarkerSymbol = 51
	ScatterMarkerSymbol51                      ScatterMarkerSymbol = "51"
	ScatterMarkerSymbolArrowBarLeft            ScatterMarkerSymbol = "arrow-bar-left"
	ScatterMarkerSymbolNumber151               ScatterMarkerSymbol = 151
	ScatterMarkerSymbol151                     ScatterMarkerSymbol = "151"
	ScatterMarkerSymbolArrowBarLeftOpen        ScatterMarkerSymbol = "arrow-bar-left-open"
	ScatterMarkerSymbolNumber52                ScatterMarkerSymbol = 52
	ScatterMarkerSymbol52                      ScatterMarkerSymbol = "52"
	ScatterMarkerSymbolArrowBarRight           ScatterMarkerSymbol = "arrow-bar-right"
	ScatterMarkerSymbolNumber152               ScatterMarkerSymbol = 152
	ScatterMarkerSymbol152                     ScatterMarkerSymbol = "152"
	ScatterMarkerSymbolArrowBarRightOpen       ScatterMarkerSymbol = "arrow-bar-right-open"
)

// ScatterOrientation Only relevant when `stackgroup` is used, and only the first `orientation` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Sets the stacking direction. With *v* (*h*), the y (x) values of subsequent traces are added. Also affects the default value of `fill`.
type ScatterOrientation string

const (
	ScatterOrientationV ScatterOrientation = "v"
	ScatterOrientationH ScatterOrientation = "h"
)

// ScatterStackgaps Only relevant when `stackgroup` is used, and only the first `stackgaps` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Determines how we handle locations at which other traces in this group have data but this one does not. With *infer zero* we insert a zero at these locations. With *interpolate* we linearly interpolate between existing values, and extrapolate a constant beyond the existing values.
type ScatterStackgaps string

const (
	ScatterStackgapsInferZero   ScatterStackgaps = "infer zero"
	ScatterStackgapsInterpolate ScatterStackgaps = "interpolate"
)

// ScatterTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterTextposition string

const (
	ScatterTextpositionTopLeft      ScatterTextposition = "top left"
	ScatterTextpositionTopCenter    ScatterTextposition = "top center"
	ScatterTextpositionTopRight     ScatterTextposition = "top right"
	ScatterTextpositionMiddleLeft   ScatterTextposition = "middle left"
	ScatterTextpositionMiddleCenter ScatterTextposition = "middle center"
	ScatterTextpositionMiddleRight  ScatterTextposition = "middle right"
	ScatterTextpositionBottomLeft   ScatterTextposition = "bottom left"
	ScatterTextpositionBottomCenter ScatterTextposition = "bottom center"
	ScatterTextpositionBottomRight  ScatterTextposition = "bottom right"
)

// ScatterVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterVisible interface{}

var (
	ScatterVisibleTrue       ScatterVisible = true
	ScatterVisibleFalse      ScatterVisible = false
	ScatterVisibleLegendonly ScatterVisible = "legendonly"
)

// ScatterXcalendar Sets the calendar system to use with `x` date data.
type ScatterXcalendar string

const (
	ScatterXcalendarGregorian  ScatterXcalendar = "gregorian"
	ScatterXcalendarChinese    ScatterXcalendar = "chinese"
	ScatterXcalendarCoptic     ScatterXcalendar = "coptic"
	ScatterXcalendarDiscworld  ScatterXcalendar = "discworld"
	ScatterXcalendarEthiopian  ScatterXcalendar = "ethiopian"
	ScatterXcalendarHebrew     ScatterXcalendar = "hebrew"
	ScatterXcalendarIslamic    ScatterXcalendar = "islamic"
	ScatterXcalendarJulian     ScatterXcalendar = "julian"
	ScatterXcalendarMayan      ScatterXcalendar = "mayan"
	ScatterXcalendarNanakshahi ScatterXcalendar = "nanakshahi"
	ScatterXcalendarNepali     ScatterXcalendar = "nepali"
	ScatterXcalendarPersian    ScatterXcalendar = "persian"
	ScatterXcalendarJalali     ScatterXcalendar = "jalali"
	ScatterXcalendarTaiwan     ScatterXcalendar = "taiwan"
	ScatterXcalendarThai       ScatterXcalendar = "thai"
	ScatterXcalendarUmmalqura  ScatterXcalendar = "ummalqura"
)

// ScatterXperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
type ScatterXperiodalignment string

const (
	ScatterXperiodalignmentStart  ScatterXperiodalignment = "start"
	ScatterXperiodalignmentMiddle ScatterXperiodalignment = "middle"
	ScatterXperiodalignmentEnd    ScatterXperiodalignment = "end"
)

// ScatterYcalendar Sets the calendar system to use with `y` date data.
type ScatterYcalendar string

const (
	ScatterYcalendarGregorian  ScatterYcalendar = "gregorian"
	ScatterYcalendarChinese    ScatterYcalendar = "chinese"
	ScatterYcalendarCoptic     ScatterYcalendar = "coptic"
	ScatterYcalendarDiscworld  ScatterYcalendar = "discworld"
	ScatterYcalendarEthiopian  ScatterYcalendar = "ethiopian"
	ScatterYcalendarHebrew     ScatterYcalendar = "hebrew"
	ScatterYcalendarIslamic    ScatterYcalendar = "islamic"
	ScatterYcalendarJulian     ScatterYcalendar = "julian"
	ScatterYcalendarMayan      ScatterYcalendar = "mayan"
	ScatterYcalendarNanakshahi ScatterYcalendar = "nanakshahi"
	ScatterYcalendarNepali     ScatterYcalendar = "nepali"
	ScatterYcalendarPersian    ScatterYcalendar = "persian"
	ScatterYcalendarJalali     ScatterYcalendar = "jalali"
	ScatterYcalendarTaiwan     ScatterYcalendar = "taiwan"
	ScatterYcalendarThai       ScatterYcalendar = "thai"
	ScatterYcalendarUmmalqura  ScatterYcalendar = "ummalqura"
)

// ScatterYperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
type ScatterYperiodalignment string

const (
	ScatterYperiodalignmentStart  ScatterYperiodalignment = "start"
	ScatterYperiodalignmentMiddle ScatterYperiodalignment = "middle"
	ScatterYperiodalignmentEnd    ScatterYperiodalignment = "end"
)

// ScatterHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type ScatterHoverinfo string

const (
	// Flags
	ScatterHoverinfoX    ScatterHoverinfo = "x"
	ScatterHoverinfoY    ScatterHoverinfo = "y"
	ScatterHoverinfoZ    ScatterHoverinfo = "z"
	ScatterHoverinfoText ScatterHoverinfo = "text"
	ScatterHoverinfoName ScatterHoverinfo = "name"

	// Extra
	ScatterHoverinfoAll  ScatterHoverinfo = "all"
	ScatterHoverinfoNone ScatterHoverinfo = "none"
	ScatterHoverinfoSkip ScatterHoverinfo = "skip"
)

// ScatterHoveron Do the hover effects highlight individual points (markers or line points) or do they highlight filled regions? If the fill is *toself* or *tonext* and there are no markers or text, then the default is *fills*, otherwise it is *points*.
type ScatterHoveron string

const (
	// Flags
	ScatterHoveronPoints ScatterHoveron = "points"
	ScatterHoveronFills  ScatterHoveron = "fills"

	// Extra

)

// ScatterMode Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
type ScatterMode string

const (
	// Flags
	ScatterModeLines   ScatterMode = "lines"
	ScatterModeMarkers ScatterMode = "markers"
	ScatterModeText    ScatterMode = "text"

	// Extra
	ScatterModeNone ScatterMode = "none"
)
