package grob

var TraceTypeScatter3d TraceType = "scatter3d"

func (trace *Scatter3d) GetType() TraceType {
	return TraceTypeScatter3d
}

// Scatter3d The data visualized as scatter point or lines in 3D dimension is set in `x`, `y`, `z`. Text (appearing either on the chart or on hover only) is via `text`. Bubble charts are achieved by setting `marker.size` and/or `marker.color` Projections are achieved via `projection`. Surface fills are achieved via `surfaceaxis`.
type Scatter3d struct {

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

	// ErrorX
	// role: Object
	ErrorX *Scatter3dErrorX `json:"error_x,omitempty"`

	// ErrorY
	// role: Object
	ErrorY *Scatter3dErrorY `json:"error_y,omitempty"`

	// ErrorZ
	// role: Object
	ErrorZ *Scatter3dErrorZ `json:"error_z,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo Scatter3dHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *Scatter3dHoverlabel `json:"hoverlabel,omitempty"`

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
	// Sets text elements associated with each (x,y,z) triplet. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y,z) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
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
	Line *Scatter3dLine `json:"line,omitempty"`

	// Marker
	// role: Object
	Marker *Scatter3dMarker `json:"marker,omitempty"`

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
	// default: lines+markers
	// type: flaglist
	// Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
	Mode Scatter3dMode `json:"mode,omitempty"`

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

	// Projection
	// role: Object
	Projection *Scatter3dProjection `json:"projection,omitempty"`

	// Scene
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's 3D coordinate system and a 3D scene. If *scene* (the default value), the (x,y,z) coordinates refer to `layout.scene`. If *scene2*, the (x,y,z) coordinates refer to `layout.scene2`, and so on.
	Scene String `json:"scene,omitempty"`

	// Showlegend
	// arrayOK: false
	// type: boolean
	// Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream
	// role: Object
	Stream *Scatter3dStream `json:"stream,omitempty"`

	// Surfaceaxis
	// default: %!s(float64=-1)
	// type: enumerated
	// If *-1*, the scatter points are not fill with a surface If *0*, *1*, *2*, the scatter points are filled with a Delaunay surface about the x, y, z respectively.
	Surfaceaxis Scatter3dSurfaceaxis `json:"surfaceaxis,omitempty"`

	// Surfacecolor
	// arrayOK: false
	// type: color
	// Sets the surface fill color.
	Surfacecolor Color `json:"surfacecolor,omitempty"`

	// Text
	// arrayOK: true
	// type: string
	// Sets text elements associated with each (x,y,z) triplet. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y,z) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text String `json:"text,omitempty"`

	// Textfont
	// role: Object
	Textfont *Scatter3dTextfont `json:"textfont,omitempty"`

	// Textposition
	// default: top center
	// type: enumerated
	// Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition Scatter3dTextposition `json:"textposition,omitempty"`

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

	// Visible
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible Scatter3dVisible `json:"visible,omitempty"`

	// X
	// arrayOK: false
	// type: data_array
	// Sets the x coordinates.
	X interface{} `json:"x,omitempty"`

	// Xcalendar
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `x` date data.
	Xcalendar Scatter3dXcalendar `json:"xcalendar,omitempty"`

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

	// Ycalendar
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `y` date data.
	Ycalendar Scatter3dYcalendar `json:"ycalendar,omitempty"`

	// Ysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Z
	// arrayOK: false
	// type: data_array
	// Sets the z coordinates.
	Z interface{} `json:"z,omitempty"`

	// Zcalendar
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `z` date data.
	Zcalendar Scatter3dZcalendar `json:"zcalendar,omitempty"`

	// Zsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

// Scatter3dErrorX
type Scatter3dErrorX struct {

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

	// CopyZstyle
	// arrayOK: false
	// type: boolean
	//
	CopyZstyle Bool `json:"copy_zstyle,omitempty"`

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
	Type Scatter3dErrorXType `json:"type,omitempty"`

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

// Scatter3dErrorY
type Scatter3dErrorY struct {

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

	// CopyZstyle
	// arrayOK: false
	// type: boolean
	//
	CopyZstyle Bool `json:"copy_zstyle,omitempty"`

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
	Type Scatter3dErrorYType `json:"type,omitempty"`

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

// Scatter3dErrorZ
type Scatter3dErrorZ struct {

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
	Type Scatter3dErrorZType `json:"type,omitempty"`

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

// Scatter3dHoverlabelFont Sets the font used in hover labels.
type Scatter3dHoverlabelFont struct {

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

// Scatter3dHoverlabel
type Scatter3dHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align Scatter3dHoverlabelAlign `json:"align,omitempty"`

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
	Font *Scatter3dHoverlabelFont `json:"font,omitempty"`

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

// Scatter3dLineColorbarTickfont Sets the color bar's tick label font
type Scatter3dLineColorbarTickfont struct {

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

// Scatter3dLineColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type Scatter3dLineColorbarTitleFont struct {

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

// Scatter3dLineColorbarTitle
type Scatter3dLineColorbarTitle struct {

	// Font
	// role: Object
	Font *Scatter3dLineColorbarTitleFont `json:"font,omitempty"`

	// Side
	// default: top
	// type: enumerated
	// Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
	Side Scatter3dLineColorbarTitleSide `json:"side,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// Scatter3dLineColorbar
type Scatter3dLineColorbar struct {

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
	Exponentformat Scatter3dLineColorbarExponentformat `json:"exponentformat,omitempty"`

	// Len
	// arrayOK: false
	// type: number
	// Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends.
	Len float64 `json:"len,omitempty"`

	// Lenmode
	// default: fraction
	// type: enumerated
	// Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
	Lenmode Scatter3dLineColorbarLenmode `json:"lenmode,omitempty"`

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
	Showexponent Scatter3dLineColorbarShowexponent `json:"showexponent,omitempty"`

	// Showticklabels
	// arrayOK: false
	// type: boolean
	// Determines whether or not the tick labels are drawn.
	Showticklabels Bool `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix Scatter3dLineColorbarShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix Scatter3dLineColorbarShowticksuffix `json:"showticksuffix,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels.
	Thickness float64 `json:"thickness,omitempty"`

	// Thicknessmode
	// default: pixels
	// type: enumerated
	// Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
	Thicknessmode Scatter3dLineColorbarThicknessmode `json:"thicknessmode,omitempty"`

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
	Tickfont *Scatter3dLineColorbarTickfont `json:"tickfont,omitempty"`

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
	Ticklabelposition Scatter3dLineColorbarTicklabelposition `json:"ticklabelposition,omitempty"`

	// Ticklen
	// arrayOK: false
	// type: number
	// Sets the tick length (in px).
	Ticklen float64 `json:"ticklen,omitempty"`

	// Tickmode
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
	Tickmode Scatter3dLineColorbarTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

	// Ticks
	// default:
	// type: enumerated
	// Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
	Ticks Scatter3dLineColorbarTicks `json:"ticks,omitempty"`

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
	Title *Scatter3dLineColorbarTitle `json:"title,omitempty"`

	// X
	// arrayOK: false
	// type: number
	// Sets the x position of the color bar (in plot fraction).
	X float64 `json:"x,omitempty"`

	// Xanchor
	// default: left
	// type: enumerated
	// Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
	Xanchor Scatter3dLineColorbarXanchor `json:"xanchor,omitempty"`

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
	Yanchor Scatter3dLineColorbarYanchor `json:"yanchor,omitempty"`

	// Ypad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the y direction.
	Ypad float64 `json:"ypad,omitempty"`
}

// Scatter3dLine
type Scatter3dLine struct {

	// Autocolorscale
	// arrayOK: false
	// type: boolean
	// Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `line.colorscale`. Has an effect only if in `line.color`is set to a numerical array. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Cauto
	// arrayOK: false
	// type: boolean
	// Determines whether or not the color domain is computed with respect to the input data (here in `line.color`) or the bounds set in `line.cmin` and `line.cmax`  Has an effect only if in `line.color`is set to a numerical array. Defaults to `false` when `line.cmin` and `line.cmax` are set by the user.
	Cauto Bool `json:"cauto,omitempty"`

	// Cmax
	// arrayOK: false
	// type: number
	// Sets the upper bound of the color domain. Has an effect only if in `line.color`is set to a numerical array. Value should have the same units as in `line.color` and if set, `line.cmin` must be set as well.
	Cmax float64 `json:"cmax,omitempty"`

	// Cmid
	// arrayOK: false
	// type: number
	// Sets the mid-point of the color domain by scaling `line.cmin` and/or `line.cmax` to be equidistant to this point. Has an effect only if in `line.color`is set to a numerical array. Value should have the same units as in `line.color`. Has no effect when `line.cauto` is `false`.
	Cmid float64 `json:"cmid,omitempty"`

	// Cmin
	// arrayOK: false
	// type: number
	// Sets the lower bound of the color domain. Has an effect only if in `line.color`is set to a numerical array. Value should have the same units as in `line.color` and if set, `line.cmax` must be set as well.
	Cmin float64 `json:"cmin,omitempty"`

	// Color
	// arrayOK: true
	// type: color
	// Sets thelinecolor. It accepts either a specific color or an array of numbers that are mapped to the colorscale relative to the max and min values of the array or relative to `line.cmin` and `line.cmax` if set.
	Color Color `json:"color,omitempty"`

	// Coloraxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar
	// role: Object
	Colorbar *Scatter3dLineColorbar `json:"colorbar,omitempty"`

	// Colorscale
	// default: %!s(<nil>)
	// type: colorscale
	// Sets the colorscale. Has an effect only if in `line.color`is set to a numerical array. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`line.cmin` and `line.cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	Colorscale ColorScale `json:"colorscale,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Dash
	// default: solid
	// type: enumerated
	// Sets the dash style of the lines.
	Dash Scatter3dLineDash `json:"dash,omitempty"`

	// Reversescale
	// arrayOK: false
	// type: boolean
	// Reverses the color mapping if true. Has an effect only if in `line.color`is set to a numerical array. If true, `line.cmin` will correspond to the last color in the array and `line.cmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Showscale
	// arrayOK: false
	// type: boolean
	// Determines whether or not a colorbar is displayed for this trace. Has an effect only if in `line.color`is set to a numerical array.
	Showscale Bool `json:"showscale,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the line width (in px).
	Width float64 `json:"width,omitempty"`
}

// Scatter3dMarkerColorbarTickfont Sets the color bar's tick label font
type Scatter3dMarkerColorbarTickfont struct {

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

// Scatter3dMarkerColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type Scatter3dMarkerColorbarTitleFont struct {

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

// Scatter3dMarkerColorbarTitle
type Scatter3dMarkerColorbarTitle struct {

	// Font
	// role: Object
	Font *Scatter3dMarkerColorbarTitleFont `json:"font,omitempty"`

	// Side
	// default: top
	// type: enumerated
	// Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
	Side Scatter3dMarkerColorbarTitleSide `json:"side,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// Scatter3dMarkerColorbar
type Scatter3dMarkerColorbar struct {

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
	Exponentformat Scatter3dMarkerColorbarExponentformat `json:"exponentformat,omitempty"`

	// Len
	// arrayOK: false
	// type: number
	// Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends.
	Len float64 `json:"len,omitempty"`

	// Lenmode
	// default: fraction
	// type: enumerated
	// Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
	Lenmode Scatter3dMarkerColorbarLenmode `json:"lenmode,omitempty"`

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
	Showexponent Scatter3dMarkerColorbarShowexponent `json:"showexponent,omitempty"`

	// Showticklabels
	// arrayOK: false
	// type: boolean
	// Determines whether or not the tick labels are drawn.
	Showticklabels Bool `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix Scatter3dMarkerColorbarShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix Scatter3dMarkerColorbarShowticksuffix `json:"showticksuffix,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels.
	Thickness float64 `json:"thickness,omitempty"`

	// Thicknessmode
	// default: pixels
	// type: enumerated
	// Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
	Thicknessmode Scatter3dMarkerColorbarThicknessmode `json:"thicknessmode,omitempty"`

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
	Tickfont *Scatter3dMarkerColorbarTickfont `json:"tickfont,omitempty"`

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
	Ticklabelposition Scatter3dMarkerColorbarTicklabelposition `json:"ticklabelposition,omitempty"`

	// Ticklen
	// arrayOK: false
	// type: number
	// Sets the tick length (in px).
	Ticklen float64 `json:"ticklen,omitempty"`

	// Tickmode
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
	Tickmode Scatter3dMarkerColorbarTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

	// Ticks
	// default:
	// type: enumerated
	// Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
	Ticks Scatter3dMarkerColorbarTicks `json:"ticks,omitempty"`

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
	Title *Scatter3dMarkerColorbarTitle `json:"title,omitempty"`

	// X
	// arrayOK: false
	// type: number
	// Sets the x position of the color bar (in plot fraction).
	X float64 `json:"x,omitempty"`

	// Xanchor
	// default: left
	// type: enumerated
	// Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
	Xanchor Scatter3dMarkerColorbarXanchor `json:"xanchor,omitempty"`

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
	Yanchor Scatter3dMarkerColorbarYanchor `json:"yanchor,omitempty"`

	// Ypad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the y direction.
	Ypad float64 `json:"ypad,omitempty"`
}

// Scatter3dMarkerLine
type Scatter3dMarkerLine struct {

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
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the lines bounding the marker points.
	Width float64 `json:"width,omitempty"`
}

// Scatter3dMarker
type Scatter3dMarker struct {

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
	Colorbar *Scatter3dMarkerColorbar `json:"colorbar,omitempty"`

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
	Line *Scatter3dMarkerLine `json:"line,omitempty"`

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the marker opacity. Note that the marker opacity for scatter3d traces must be a scalar value for performance reasons. To set a blending opacity value (i.e. which is not transparent), set *marker.color* to an rgba color and use its alpha channel.
	Opacity float64 `json:"opacity,omitempty"`

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
	Sizemode Scatter3dMarkerSizemode `json:"sizemode,omitempty"`

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
	// Sets the marker symbol type.
	Symbol Scatter3dMarkerSymbol `json:"symbol,omitempty"`

	// Symbolsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  symbol .
	Symbolsrc String `json:"symbolsrc,omitempty"`
}

// Scatter3dProjectionX
type Scatter3dProjectionX struct {

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the projection color.
	Opacity float64 `json:"opacity,omitempty"`

	// Scale
	// arrayOK: false
	// type: number
	// Sets the scale factor determining the size of the projection marker points.
	Scale float64 `json:"scale,omitempty"`

	// Show
	// arrayOK: false
	// type: boolean
	// Sets whether or not projections are shown along the x axis.
	Show Bool `json:"show,omitempty"`
}

// Scatter3dProjectionY
type Scatter3dProjectionY struct {

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the projection color.
	Opacity float64 `json:"opacity,omitempty"`

	// Scale
	// arrayOK: false
	// type: number
	// Sets the scale factor determining the size of the projection marker points.
	Scale float64 `json:"scale,omitempty"`

	// Show
	// arrayOK: false
	// type: boolean
	// Sets whether or not projections are shown along the y axis.
	Show Bool `json:"show,omitempty"`
}

// Scatter3dProjectionZ
type Scatter3dProjectionZ struct {

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the projection color.
	Opacity float64 `json:"opacity,omitempty"`

	// Scale
	// arrayOK: false
	// type: number
	// Sets the scale factor determining the size of the projection marker points.
	Scale float64 `json:"scale,omitempty"`

	// Show
	// arrayOK: false
	// type: boolean
	// Sets whether or not projections are shown along the z axis.
	Show Bool `json:"show,omitempty"`
}

// Scatter3dProjection
type Scatter3dProjection struct {

	// X
	// role: Object
	X *Scatter3dProjectionX `json:"x,omitempty"`

	// Y
	// role: Object
	Y *Scatter3dProjectionY `json:"y,omitempty"`

	// Z
	// role: Object
	Z *Scatter3dProjectionZ `json:"z,omitempty"`
}

// Scatter3dStream
type Scatter3dStream struct {

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

// Scatter3dTextfont
type Scatter3dTextfont struct {

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
	// arrayOK: false
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family String `json:"family,omitempty"`

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

// Scatter3dErrorXType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the square of the underlying data. If *data*, the bar lengths are set with data set `array`.
type Scatter3dErrorXType string

const (
	Scatter3dErrorXTypePercent  Scatter3dErrorXType = "percent"
	Scatter3dErrorXTypeConstant Scatter3dErrorXType = "constant"
	Scatter3dErrorXTypeSqrt     Scatter3dErrorXType = "sqrt"
	Scatter3dErrorXTypeData     Scatter3dErrorXType = "data"
)

// Scatter3dErrorYType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the square of the underlying data. If *data*, the bar lengths are set with data set `array`.
type Scatter3dErrorYType string

const (
	Scatter3dErrorYTypePercent  Scatter3dErrorYType = "percent"
	Scatter3dErrorYTypeConstant Scatter3dErrorYType = "constant"
	Scatter3dErrorYTypeSqrt     Scatter3dErrorYType = "sqrt"
	Scatter3dErrorYTypeData     Scatter3dErrorYType = "data"
)

// Scatter3dErrorZType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the square of the underlying data. If *data*, the bar lengths are set with data set `array`.
type Scatter3dErrorZType string

const (
	Scatter3dErrorZTypePercent  Scatter3dErrorZType = "percent"
	Scatter3dErrorZTypeConstant Scatter3dErrorZType = "constant"
	Scatter3dErrorZTypeSqrt     Scatter3dErrorZType = "sqrt"
	Scatter3dErrorZTypeData     Scatter3dErrorZType = "data"
)

// Scatter3dHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type Scatter3dHoverlabelAlign string

const (
	Scatter3dHoverlabelAlignLeft  Scatter3dHoverlabelAlign = "left"
	Scatter3dHoverlabelAlignRight Scatter3dHoverlabelAlign = "right"
	Scatter3dHoverlabelAlignAuto  Scatter3dHoverlabelAlign = "auto"
)

// Scatter3dLineColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type Scatter3dLineColorbarExponentformat string

const (
	Scatter3dLineColorbarExponentformatNone  Scatter3dLineColorbarExponentformat = "none"
	Scatter3dLineColorbarExponentformatE1    Scatter3dLineColorbarExponentformat = "e"
	Scatter3dLineColorbarExponentformatE2    Scatter3dLineColorbarExponentformat = "E"
	Scatter3dLineColorbarExponentformatPower Scatter3dLineColorbarExponentformat = "power"
	Scatter3dLineColorbarExponentformatSi    Scatter3dLineColorbarExponentformat = "SI"
	Scatter3dLineColorbarExponentformatB     Scatter3dLineColorbarExponentformat = "B"
)

// Scatter3dLineColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type Scatter3dLineColorbarLenmode string

const (
	Scatter3dLineColorbarLenmodeFraction Scatter3dLineColorbarLenmode = "fraction"
	Scatter3dLineColorbarLenmodePixels   Scatter3dLineColorbarLenmode = "pixels"
)

// Scatter3dLineColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type Scatter3dLineColorbarShowexponent string

const (
	Scatter3dLineColorbarShowexponentAll   Scatter3dLineColorbarShowexponent = "all"
	Scatter3dLineColorbarShowexponentFirst Scatter3dLineColorbarShowexponent = "first"
	Scatter3dLineColorbarShowexponentLast  Scatter3dLineColorbarShowexponent = "last"
	Scatter3dLineColorbarShowexponentNone  Scatter3dLineColorbarShowexponent = "none"
)

// Scatter3dLineColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type Scatter3dLineColorbarShowtickprefix string

const (
	Scatter3dLineColorbarShowtickprefixAll   Scatter3dLineColorbarShowtickprefix = "all"
	Scatter3dLineColorbarShowtickprefixFirst Scatter3dLineColorbarShowtickprefix = "first"
	Scatter3dLineColorbarShowtickprefixLast  Scatter3dLineColorbarShowtickprefix = "last"
	Scatter3dLineColorbarShowtickprefixNone  Scatter3dLineColorbarShowtickprefix = "none"
)

// Scatter3dLineColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type Scatter3dLineColorbarShowticksuffix string

const (
	Scatter3dLineColorbarShowticksuffixAll   Scatter3dLineColorbarShowticksuffix = "all"
	Scatter3dLineColorbarShowticksuffixFirst Scatter3dLineColorbarShowticksuffix = "first"
	Scatter3dLineColorbarShowticksuffixLast  Scatter3dLineColorbarShowticksuffix = "last"
	Scatter3dLineColorbarShowticksuffixNone  Scatter3dLineColorbarShowticksuffix = "none"
)

// Scatter3dLineColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type Scatter3dLineColorbarThicknessmode string

const (
	Scatter3dLineColorbarThicknessmodeFraction Scatter3dLineColorbarThicknessmode = "fraction"
	Scatter3dLineColorbarThicknessmodePixels   Scatter3dLineColorbarThicknessmode = "pixels"
)

// Scatter3dLineColorbarTicklabelposition Determines where tick labels are drawn.
type Scatter3dLineColorbarTicklabelposition string

const (
	Scatter3dLineColorbarTicklabelpositionOutside       Scatter3dLineColorbarTicklabelposition = "outside"
	Scatter3dLineColorbarTicklabelpositionInside        Scatter3dLineColorbarTicklabelposition = "inside"
	Scatter3dLineColorbarTicklabelpositionOutsideTop    Scatter3dLineColorbarTicklabelposition = "outside top"
	Scatter3dLineColorbarTicklabelpositionInsideTop     Scatter3dLineColorbarTicklabelposition = "inside top"
	Scatter3dLineColorbarTicklabelpositionOutsideBottom Scatter3dLineColorbarTicklabelposition = "outside bottom"
	Scatter3dLineColorbarTicklabelpositionInsideBottom  Scatter3dLineColorbarTicklabelposition = "inside bottom"
)

// Scatter3dLineColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type Scatter3dLineColorbarTickmode string

const (
	Scatter3dLineColorbarTickmodeAuto   Scatter3dLineColorbarTickmode = "auto"
	Scatter3dLineColorbarTickmodeLinear Scatter3dLineColorbarTickmode = "linear"
	Scatter3dLineColorbarTickmodeArray  Scatter3dLineColorbarTickmode = "array"
)

// Scatter3dLineColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type Scatter3dLineColorbarTicks string

const (
	Scatter3dLineColorbarTicksOutside Scatter3dLineColorbarTicks = "outside"
	Scatter3dLineColorbarTicksInside  Scatter3dLineColorbarTicks = "inside"
	Scatter3dLineColorbarTicksEmpty   Scatter3dLineColorbarTicks = ""
)

// Scatter3dLineColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type Scatter3dLineColorbarTitleSide string

const (
	Scatter3dLineColorbarTitleSideRight  Scatter3dLineColorbarTitleSide = "right"
	Scatter3dLineColorbarTitleSideTop    Scatter3dLineColorbarTitleSide = "top"
	Scatter3dLineColorbarTitleSideBottom Scatter3dLineColorbarTitleSide = "bottom"
)

// Scatter3dLineColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type Scatter3dLineColorbarXanchor string

const (
	Scatter3dLineColorbarXanchorLeft   Scatter3dLineColorbarXanchor = "left"
	Scatter3dLineColorbarXanchorCenter Scatter3dLineColorbarXanchor = "center"
	Scatter3dLineColorbarXanchorRight  Scatter3dLineColorbarXanchor = "right"
)

// Scatter3dLineColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type Scatter3dLineColorbarYanchor string

const (
	Scatter3dLineColorbarYanchorTop    Scatter3dLineColorbarYanchor = "top"
	Scatter3dLineColorbarYanchorMiddle Scatter3dLineColorbarYanchor = "middle"
	Scatter3dLineColorbarYanchorBottom Scatter3dLineColorbarYanchor = "bottom"
)

// Scatter3dLineDash Sets the dash style of the lines.
type Scatter3dLineDash string

const (
	Scatter3dLineDashSolid       Scatter3dLineDash = "solid"
	Scatter3dLineDashDot         Scatter3dLineDash = "dot"
	Scatter3dLineDashDash        Scatter3dLineDash = "dash"
	Scatter3dLineDashLongdash    Scatter3dLineDash = "longdash"
	Scatter3dLineDashDashdot     Scatter3dLineDash = "dashdot"
	Scatter3dLineDashLongdashdot Scatter3dLineDash = "longdashdot"
)

// Scatter3dMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type Scatter3dMarkerColorbarExponentformat string

const (
	Scatter3dMarkerColorbarExponentformatNone  Scatter3dMarkerColorbarExponentformat = "none"
	Scatter3dMarkerColorbarExponentformatE1    Scatter3dMarkerColorbarExponentformat = "e"
	Scatter3dMarkerColorbarExponentformatE2    Scatter3dMarkerColorbarExponentformat = "E"
	Scatter3dMarkerColorbarExponentformatPower Scatter3dMarkerColorbarExponentformat = "power"
	Scatter3dMarkerColorbarExponentformatSi    Scatter3dMarkerColorbarExponentformat = "SI"
	Scatter3dMarkerColorbarExponentformatB     Scatter3dMarkerColorbarExponentformat = "B"
)

// Scatter3dMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type Scatter3dMarkerColorbarLenmode string

const (
	Scatter3dMarkerColorbarLenmodeFraction Scatter3dMarkerColorbarLenmode = "fraction"
	Scatter3dMarkerColorbarLenmodePixels   Scatter3dMarkerColorbarLenmode = "pixels"
)

// Scatter3dMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type Scatter3dMarkerColorbarShowexponent string

const (
	Scatter3dMarkerColorbarShowexponentAll   Scatter3dMarkerColorbarShowexponent = "all"
	Scatter3dMarkerColorbarShowexponentFirst Scatter3dMarkerColorbarShowexponent = "first"
	Scatter3dMarkerColorbarShowexponentLast  Scatter3dMarkerColorbarShowexponent = "last"
	Scatter3dMarkerColorbarShowexponentNone  Scatter3dMarkerColorbarShowexponent = "none"
)

// Scatter3dMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type Scatter3dMarkerColorbarShowtickprefix string

const (
	Scatter3dMarkerColorbarShowtickprefixAll   Scatter3dMarkerColorbarShowtickprefix = "all"
	Scatter3dMarkerColorbarShowtickprefixFirst Scatter3dMarkerColorbarShowtickprefix = "first"
	Scatter3dMarkerColorbarShowtickprefixLast  Scatter3dMarkerColorbarShowtickprefix = "last"
	Scatter3dMarkerColorbarShowtickprefixNone  Scatter3dMarkerColorbarShowtickprefix = "none"
)

// Scatter3dMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type Scatter3dMarkerColorbarShowticksuffix string

const (
	Scatter3dMarkerColorbarShowticksuffixAll   Scatter3dMarkerColorbarShowticksuffix = "all"
	Scatter3dMarkerColorbarShowticksuffixFirst Scatter3dMarkerColorbarShowticksuffix = "first"
	Scatter3dMarkerColorbarShowticksuffixLast  Scatter3dMarkerColorbarShowticksuffix = "last"
	Scatter3dMarkerColorbarShowticksuffixNone  Scatter3dMarkerColorbarShowticksuffix = "none"
)

// Scatter3dMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type Scatter3dMarkerColorbarThicknessmode string

const (
	Scatter3dMarkerColorbarThicknessmodeFraction Scatter3dMarkerColorbarThicknessmode = "fraction"
	Scatter3dMarkerColorbarThicknessmodePixels   Scatter3dMarkerColorbarThicknessmode = "pixels"
)

// Scatter3dMarkerColorbarTicklabelposition Determines where tick labels are drawn.
type Scatter3dMarkerColorbarTicklabelposition string

const (
	Scatter3dMarkerColorbarTicklabelpositionOutside       Scatter3dMarkerColorbarTicklabelposition = "outside"
	Scatter3dMarkerColorbarTicklabelpositionInside        Scatter3dMarkerColorbarTicklabelposition = "inside"
	Scatter3dMarkerColorbarTicklabelpositionOutsideTop    Scatter3dMarkerColorbarTicklabelposition = "outside top"
	Scatter3dMarkerColorbarTicklabelpositionInsideTop     Scatter3dMarkerColorbarTicklabelposition = "inside top"
	Scatter3dMarkerColorbarTicklabelpositionOutsideBottom Scatter3dMarkerColorbarTicklabelposition = "outside bottom"
	Scatter3dMarkerColorbarTicklabelpositionInsideBottom  Scatter3dMarkerColorbarTicklabelposition = "inside bottom"
)

// Scatter3dMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type Scatter3dMarkerColorbarTickmode string

const (
	Scatter3dMarkerColorbarTickmodeAuto   Scatter3dMarkerColorbarTickmode = "auto"
	Scatter3dMarkerColorbarTickmodeLinear Scatter3dMarkerColorbarTickmode = "linear"
	Scatter3dMarkerColorbarTickmodeArray  Scatter3dMarkerColorbarTickmode = "array"
)

// Scatter3dMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type Scatter3dMarkerColorbarTicks string

const (
	Scatter3dMarkerColorbarTicksOutside Scatter3dMarkerColorbarTicks = "outside"
	Scatter3dMarkerColorbarTicksInside  Scatter3dMarkerColorbarTicks = "inside"
	Scatter3dMarkerColorbarTicksEmpty   Scatter3dMarkerColorbarTicks = ""
)

// Scatter3dMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type Scatter3dMarkerColorbarTitleSide string

const (
	Scatter3dMarkerColorbarTitleSideRight  Scatter3dMarkerColorbarTitleSide = "right"
	Scatter3dMarkerColorbarTitleSideTop    Scatter3dMarkerColorbarTitleSide = "top"
	Scatter3dMarkerColorbarTitleSideBottom Scatter3dMarkerColorbarTitleSide = "bottom"
)

// Scatter3dMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type Scatter3dMarkerColorbarXanchor string

const (
	Scatter3dMarkerColorbarXanchorLeft   Scatter3dMarkerColorbarXanchor = "left"
	Scatter3dMarkerColorbarXanchorCenter Scatter3dMarkerColorbarXanchor = "center"
	Scatter3dMarkerColorbarXanchorRight  Scatter3dMarkerColorbarXanchor = "right"
)

// Scatter3dMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type Scatter3dMarkerColorbarYanchor string

const (
	Scatter3dMarkerColorbarYanchorTop    Scatter3dMarkerColorbarYanchor = "top"
	Scatter3dMarkerColorbarYanchorMiddle Scatter3dMarkerColorbarYanchor = "middle"
	Scatter3dMarkerColorbarYanchorBottom Scatter3dMarkerColorbarYanchor = "bottom"
)

// Scatter3dMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type Scatter3dMarkerSizemode string

const (
	Scatter3dMarkerSizemodeDiameter Scatter3dMarkerSizemode = "diameter"
	Scatter3dMarkerSizemodeArea     Scatter3dMarkerSizemode = "area"
)

// Scatter3dMarkerSymbol Sets the marker symbol type.
type Scatter3dMarkerSymbol string

const (
	Scatter3dMarkerSymbolCircle      Scatter3dMarkerSymbol = "circle"
	Scatter3dMarkerSymbolCircleOpen  Scatter3dMarkerSymbol = "circle-open"
	Scatter3dMarkerSymbolSquare      Scatter3dMarkerSymbol = "square"
	Scatter3dMarkerSymbolSquareOpen  Scatter3dMarkerSymbol = "square-open"
	Scatter3dMarkerSymbolDiamond     Scatter3dMarkerSymbol = "diamond"
	Scatter3dMarkerSymbolDiamondOpen Scatter3dMarkerSymbol = "diamond-open"
	Scatter3dMarkerSymbolCross       Scatter3dMarkerSymbol = "cross"
	Scatter3dMarkerSymbolX           Scatter3dMarkerSymbol = "x"
)

// Scatter3dSurfaceaxis If *-1*, the scatter points are not fill with a surface If *0*, *1*, *2*, the scatter points are filled with a Delaunay surface about the x, y, z respectively.
type Scatter3dSurfaceaxis interface{}

var (
	Scatter3dSurfaceaxisNumberNegative1 Scatter3dSurfaceaxis = -1
	Scatter3dSurfaceaxisNumber0         Scatter3dSurfaceaxis = 0
	Scatter3dSurfaceaxisNumber1         Scatter3dSurfaceaxis = 1
	Scatter3dSurfaceaxisNumber2         Scatter3dSurfaceaxis = 2
)

// Scatter3dTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type Scatter3dTextposition string

const (
	Scatter3dTextpositionTopLeft      Scatter3dTextposition = "top left"
	Scatter3dTextpositionTopCenter    Scatter3dTextposition = "top center"
	Scatter3dTextpositionTopRight     Scatter3dTextposition = "top right"
	Scatter3dTextpositionMiddleLeft   Scatter3dTextposition = "middle left"
	Scatter3dTextpositionMiddleCenter Scatter3dTextposition = "middle center"
	Scatter3dTextpositionMiddleRight  Scatter3dTextposition = "middle right"
	Scatter3dTextpositionBottomLeft   Scatter3dTextposition = "bottom left"
	Scatter3dTextpositionBottomCenter Scatter3dTextposition = "bottom center"
	Scatter3dTextpositionBottomRight  Scatter3dTextposition = "bottom right"
)

// Scatter3dVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type Scatter3dVisible interface{}

var (
	Scatter3dVisibleTrue       Scatter3dVisible = true
	Scatter3dVisibleFalse      Scatter3dVisible = false
	Scatter3dVisibleLegendonly Scatter3dVisible = "legendonly"
)

// Scatter3dXcalendar Sets the calendar system to use with `x` date data.
type Scatter3dXcalendar string

const (
	Scatter3dXcalendarGregorian  Scatter3dXcalendar = "gregorian"
	Scatter3dXcalendarChinese    Scatter3dXcalendar = "chinese"
	Scatter3dXcalendarCoptic     Scatter3dXcalendar = "coptic"
	Scatter3dXcalendarDiscworld  Scatter3dXcalendar = "discworld"
	Scatter3dXcalendarEthiopian  Scatter3dXcalendar = "ethiopian"
	Scatter3dXcalendarHebrew     Scatter3dXcalendar = "hebrew"
	Scatter3dXcalendarIslamic    Scatter3dXcalendar = "islamic"
	Scatter3dXcalendarJulian     Scatter3dXcalendar = "julian"
	Scatter3dXcalendarMayan      Scatter3dXcalendar = "mayan"
	Scatter3dXcalendarNanakshahi Scatter3dXcalendar = "nanakshahi"
	Scatter3dXcalendarNepali     Scatter3dXcalendar = "nepali"
	Scatter3dXcalendarPersian    Scatter3dXcalendar = "persian"
	Scatter3dXcalendarJalali     Scatter3dXcalendar = "jalali"
	Scatter3dXcalendarTaiwan     Scatter3dXcalendar = "taiwan"
	Scatter3dXcalendarThai       Scatter3dXcalendar = "thai"
	Scatter3dXcalendarUmmalqura  Scatter3dXcalendar = "ummalqura"
)

// Scatter3dYcalendar Sets the calendar system to use with `y` date data.
type Scatter3dYcalendar string

const (
	Scatter3dYcalendarGregorian  Scatter3dYcalendar = "gregorian"
	Scatter3dYcalendarChinese    Scatter3dYcalendar = "chinese"
	Scatter3dYcalendarCoptic     Scatter3dYcalendar = "coptic"
	Scatter3dYcalendarDiscworld  Scatter3dYcalendar = "discworld"
	Scatter3dYcalendarEthiopian  Scatter3dYcalendar = "ethiopian"
	Scatter3dYcalendarHebrew     Scatter3dYcalendar = "hebrew"
	Scatter3dYcalendarIslamic    Scatter3dYcalendar = "islamic"
	Scatter3dYcalendarJulian     Scatter3dYcalendar = "julian"
	Scatter3dYcalendarMayan      Scatter3dYcalendar = "mayan"
	Scatter3dYcalendarNanakshahi Scatter3dYcalendar = "nanakshahi"
	Scatter3dYcalendarNepali     Scatter3dYcalendar = "nepali"
	Scatter3dYcalendarPersian    Scatter3dYcalendar = "persian"
	Scatter3dYcalendarJalali     Scatter3dYcalendar = "jalali"
	Scatter3dYcalendarTaiwan     Scatter3dYcalendar = "taiwan"
	Scatter3dYcalendarThai       Scatter3dYcalendar = "thai"
	Scatter3dYcalendarUmmalqura  Scatter3dYcalendar = "ummalqura"
)

// Scatter3dZcalendar Sets the calendar system to use with `z` date data.
type Scatter3dZcalendar string

const (
	Scatter3dZcalendarGregorian  Scatter3dZcalendar = "gregorian"
	Scatter3dZcalendarChinese    Scatter3dZcalendar = "chinese"
	Scatter3dZcalendarCoptic     Scatter3dZcalendar = "coptic"
	Scatter3dZcalendarDiscworld  Scatter3dZcalendar = "discworld"
	Scatter3dZcalendarEthiopian  Scatter3dZcalendar = "ethiopian"
	Scatter3dZcalendarHebrew     Scatter3dZcalendar = "hebrew"
	Scatter3dZcalendarIslamic    Scatter3dZcalendar = "islamic"
	Scatter3dZcalendarJulian     Scatter3dZcalendar = "julian"
	Scatter3dZcalendarMayan      Scatter3dZcalendar = "mayan"
	Scatter3dZcalendarNanakshahi Scatter3dZcalendar = "nanakshahi"
	Scatter3dZcalendarNepali     Scatter3dZcalendar = "nepali"
	Scatter3dZcalendarPersian    Scatter3dZcalendar = "persian"
	Scatter3dZcalendarJalali     Scatter3dZcalendar = "jalali"
	Scatter3dZcalendarTaiwan     Scatter3dZcalendar = "taiwan"
	Scatter3dZcalendarThai       Scatter3dZcalendar = "thai"
	Scatter3dZcalendarUmmalqura  Scatter3dZcalendar = "ummalqura"
)

// Scatter3dHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type Scatter3dHoverinfo string

const (
	// Flags
	Scatter3dHoverinfoX    Scatter3dHoverinfo = "x"
	Scatter3dHoverinfoY    Scatter3dHoverinfo = "y"
	Scatter3dHoverinfoZ    Scatter3dHoverinfo = "z"
	Scatter3dHoverinfoText Scatter3dHoverinfo = "text"
	Scatter3dHoverinfoName Scatter3dHoverinfo = "name"

	// Extra
	Scatter3dHoverinfoAll  Scatter3dHoverinfo = "all"
	Scatter3dHoverinfoNone Scatter3dHoverinfo = "none"
	Scatter3dHoverinfoSkip Scatter3dHoverinfo = "skip"
)

// Scatter3dMode Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
type Scatter3dMode string

const (
	// Flags
	Scatter3dModeLines   Scatter3dMode = "lines"
	Scatter3dModeMarkers Scatter3dMode = "markers"
	Scatter3dModeText    Scatter3dMode = "text"

	// Extra
	Scatter3dModeNone Scatter3dMode = "none"
)
