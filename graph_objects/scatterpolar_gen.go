package grob

var TraceTypeScatterpolar TraceType = "scatterpolar"

func (trace *Scatterpolar) GetType() TraceType {
	return TraceTypeScatterpolar
}

// Scatterpolar The scatterpolar trace type encompasses line charts, scatter charts, text charts, and bubble charts in polar coordinates. The data visualized as scatter point or lines is set in `r` (radial) and `theta` (angular) coordinates Text (appearing either on the chart or on hover only) is via `text`. Bubble charts are achieved by setting `marker.size` and/or `marker.color` to numerical arrays.
type Scatterpolar struct {

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

	// Dr
	// arrayOK: false
	// type: number
	// Sets the r coordinate step.
	Dr float64 `json:"dr,omitempty"`

	// Dtheta
	// arrayOK: false
	// type: number
	// Sets the theta coordinate step. By default, the `dtheta` step equals the subplot's period divided by the length of the `r` coordinates.
	Dtheta float64 `json:"dtheta,omitempty"`

	// Fill
	// default: none
	// type: enumerated
	// Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. scatterpolar has a subset of the options available to scatter. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other.
	Fill ScatterpolarFill `json:"fill,omitempty"`

	// Fillcolor
	// arrayOK: false
	// type: color
	// Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor Color `json:"fillcolor,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ScatterpolarHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *ScatterpolarHoverlabel `json:"hoverlabel,omitempty"`

	// Hoveron
	// default: %!s(<nil>)
	// type: flaglist
	// Do the hover effects highlight individual points (markers or line points) or do they highlight filled regions? If the fill is *toself* or *tonext* and there are no markers or text, then the default is *fills*, otherwise it is *points*.
	Hoveron ScatterpolarHoveron `json:"hoveron,omitempty"`

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
	Line *ScatterpolarLine `json:"line,omitempty"`

	// Marker
	// role: Object
	Marker *ScatterpolarMarker `json:"marker,omitempty"`

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
	Mode ScatterpolarMode `json:"mode,omitempty"`

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

	// R
	// arrayOK: false
	// type: data_array
	// Sets the radial coordinates
	R interface{} `json:"r,omitempty"`

	// R0
	// arrayOK: false
	// type: any
	// Alternate to `r`. Builds a linear space of r coordinates. Use with `dr` where `r0` is the starting coordinate and `dr` the step.
	R0 interface{} `json:"r0,omitempty"`

	// Rsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  r .
	Rsrc String `json:"rsrc,omitempty"`

	// Selected
	// role: Object
	Selected *ScatterpolarSelected `json:"selected,omitempty"`

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
	Stream *ScatterpolarStream `json:"stream,omitempty"`

	// Subplot
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's data coordinates and a polar subplot. If *polar* (the default value), the data refer to `layout.polar`. If *polar2*, the data refer to `layout.polar2`, and so on.
	Subplot String `json:"subplot,omitempty"`

	// Text
	// arrayOK: true
	// type: string
	// Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text String `json:"text,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScatterpolarTextfont `json:"textfont,omitempty"`

	// Textposition
	// default: middle center
	// type: enumerated
	// Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition ScatterpolarTextposition `json:"textposition,omitempty"`

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
	// Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `r`, `theta` and `text`.
	Texttemplate String `json:"texttemplate,omitempty"`

	// Texttemplatesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Theta
	// arrayOK: false
	// type: data_array
	// Sets the angular coordinates
	Theta interface{} `json:"theta,omitempty"`

	// Theta0
	// arrayOK: false
	// type: any
	// Alternate to `theta`. Builds a linear space of theta coordinates. Use with `dtheta` where `theta0` is the starting coordinate and `dtheta` the step.
	Theta0 interface{} `json:"theta0,omitempty"`

	// Thetasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  theta .
	Thetasrc String `json:"thetasrc,omitempty"`

	// Thetaunit
	// default: degrees
	// type: enumerated
	// Sets the unit of input *theta* values. Has an effect only when on *linear* angular axes.
	Thetaunit ScatterpolarThetaunit `json:"thetaunit,omitempty"`

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
	Unselected *ScatterpolarUnselected `json:"unselected,omitempty"`

	// Visible
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ScatterpolarVisible `json:"visible,omitempty"`
}

// ScatterpolarHoverlabelFont Sets the font used in hover labels.
type ScatterpolarHoverlabelFont struct {

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

// ScatterpolarHoverlabel
type ScatterpolarHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align ScatterpolarHoverlabelAlign `json:"align,omitempty"`

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
	Font *ScatterpolarHoverlabelFont `json:"font,omitempty"`

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

// ScatterpolarLine
type ScatterpolarLine struct {

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
	Shape ScatterpolarLineShape `json:"shape,omitempty"`

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

// ScatterpolarMarkerColorbarTickfont Sets the color bar's tick label font
type ScatterpolarMarkerColorbarTickfont struct {

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

// ScatterpolarMarkerColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type ScatterpolarMarkerColorbarTitleFont struct {

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

// ScatterpolarMarkerColorbarTitle
type ScatterpolarMarkerColorbarTitle struct {

	// Font
	// role: Object
	Font *ScatterpolarMarkerColorbarTitleFont `json:"font,omitempty"`

	// Side
	// default: top
	// type: enumerated
	// Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
	Side ScatterpolarMarkerColorbarTitleSide `json:"side,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// ScatterpolarMarkerColorbar
type ScatterpolarMarkerColorbar struct {

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
	Exponentformat ScatterpolarMarkerColorbarExponentformat `json:"exponentformat,omitempty"`

	// Len
	// arrayOK: false
	// type: number
	// Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends.
	Len float64 `json:"len,omitempty"`

	// Lenmode
	// default: fraction
	// type: enumerated
	// Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
	Lenmode ScatterpolarMarkerColorbarLenmode `json:"lenmode,omitempty"`

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
	Showexponent ScatterpolarMarkerColorbarShowexponent `json:"showexponent,omitempty"`

	// Showticklabels
	// arrayOK: false
	// type: boolean
	// Determines whether or not the tick labels are drawn.
	Showticklabels Bool `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix ScatterpolarMarkerColorbarShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix ScatterpolarMarkerColorbarShowticksuffix `json:"showticksuffix,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels.
	Thickness float64 `json:"thickness,omitempty"`

	// Thicknessmode
	// default: pixels
	// type: enumerated
	// Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
	Thicknessmode ScatterpolarMarkerColorbarThicknessmode `json:"thicknessmode,omitempty"`

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
	Tickfont *ScatterpolarMarkerColorbarTickfont `json:"tickfont,omitempty"`

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
	Ticklabelposition ScatterpolarMarkerColorbarTicklabelposition `json:"ticklabelposition,omitempty"`

	// Ticklen
	// arrayOK: false
	// type: number
	// Sets the tick length (in px).
	Ticklen float64 `json:"ticklen,omitempty"`

	// Tickmode
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
	Tickmode ScatterpolarMarkerColorbarTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

	// Ticks
	// default:
	// type: enumerated
	// Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
	Ticks ScatterpolarMarkerColorbarTicks `json:"ticks,omitempty"`

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
	Title *ScatterpolarMarkerColorbarTitle `json:"title,omitempty"`

	// X
	// arrayOK: false
	// type: number
	// Sets the x position of the color bar (in plot fraction).
	X float64 `json:"x,omitempty"`

	// Xanchor
	// default: left
	// type: enumerated
	// Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
	Xanchor ScatterpolarMarkerColorbarXanchor `json:"xanchor,omitempty"`

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
	Yanchor ScatterpolarMarkerColorbarYanchor `json:"yanchor,omitempty"`

	// Ypad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the y direction.
	Ypad float64 `json:"ypad,omitempty"`
}

// ScatterpolarMarkerGradient
type ScatterpolarMarkerGradient struct {

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
	Type ScatterpolarMarkerGradientType `json:"type,omitempty"`

	// Typesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  type .
	Typesrc String `json:"typesrc,omitempty"`
}

// ScatterpolarMarkerLine
type ScatterpolarMarkerLine struct {

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

// ScatterpolarMarker
type ScatterpolarMarker struct {

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
	Colorbar *ScatterpolarMarkerColorbar `json:"colorbar,omitempty"`

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
	Gradient *ScatterpolarMarkerGradient `json:"gradient,omitempty"`

	// Line
	// role: Object
	Line *ScatterpolarMarkerLine `json:"line,omitempty"`

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
	Sizemode ScatterpolarMarkerSizemode `json:"sizemode,omitempty"`

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
	Symbol ScatterpolarMarkerSymbol `json:"symbol,omitempty"`

	// Symbolsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  symbol .
	Symbolsrc String `json:"symbolsrc,omitempty"`
}

// ScatterpolarSelectedMarker
type ScatterpolarSelectedMarker struct {

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

// ScatterpolarSelectedTextfont
type ScatterpolarSelectedTextfont struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the text font color of selected points.
	Color Color `json:"color,omitempty"`
}

// ScatterpolarSelected
type ScatterpolarSelected struct {

	// Marker
	// role: Object
	Marker *ScatterpolarSelectedMarker `json:"marker,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScatterpolarSelectedTextfont `json:"textfont,omitempty"`
}

// ScatterpolarStream
type ScatterpolarStream struct {

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

// ScatterpolarTextfont Sets the text font.
type ScatterpolarTextfont struct {

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

// ScatterpolarUnselectedMarker
type ScatterpolarUnselectedMarker struct {

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

// ScatterpolarUnselectedTextfont
type ScatterpolarUnselectedTextfont struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the text font color of unselected points, applied only when a selection exists.
	Color Color `json:"color,omitempty"`
}

// ScatterpolarUnselected
type ScatterpolarUnselected struct {

	// Marker
	// role: Object
	Marker *ScatterpolarUnselectedMarker `json:"marker,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScatterpolarUnselectedTextfont `json:"textfont,omitempty"`
}

// ScatterpolarFill Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. scatterpolar has a subset of the options available to scatter. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other.
type ScatterpolarFill string

const (
	ScatterpolarFillNone   ScatterpolarFill = "none"
	ScatterpolarFillToself ScatterpolarFill = "toself"
	ScatterpolarFillTonext ScatterpolarFill = "tonext"
)

// ScatterpolarHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterpolarHoverlabelAlign string

const (
	ScatterpolarHoverlabelAlignLeft  ScatterpolarHoverlabelAlign = "left"
	ScatterpolarHoverlabelAlignRight ScatterpolarHoverlabelAlign = "right"
	ScatterpolarHoverlabelAlignAuto  ScatterpolarHoverlabelAlign = "auto"
)

// ScatterpolarLineShape Determines the line shape. With *spline* the lines are drawn using spline interpolation. The other available values correspond to step-wise line shapes.
type ScatterpolarLineShape string

const (
	ScatterpolarLineShapeLinear ScatterpolarLineShape = "linear"
	ScatterpolarLineShapeSpline ScatterpolarLineShape = "spline"
)

// ScatterpolarMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterpolarMarkerColorbarExponentformat string

const (
	ScatterpolarMarkerColorbarExponentformatNone  ScatterpolarMarkerColorbarExponentformat = "none"
	ScatterpolarMarkerColorbarExponentformatE1    ScatterpolarMarkerColorbarExponentformat = "e"
	ScatterpolarMarkerColorbarExponentformatE2    ScatterpolarMarkerColorbarExponentformat = "E"
	ScatterpolarMarkerColorbarExponentformatPower ScatterpolarMarkerColorbarExponentformat = "power"
	ScatterpolarMarkerColorbarExponentformatSi    ScatterpolarMarkerColorbarExponentformat = "SI"
	ScatterpolarMarkerColorbarExponentformatB     ScatterpolarMarkerColorbarExponentformat = "B"
)

// ScatterpolarMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterpolarMarkerColorbarLenmode string

const (
	ScatterpolarMarkerColorbarLenmodeFraction ScatterpolarMarkerColorbarLenmode = "fraction"
	ScatterpolarMarkerColorbarLenmodePixels   ScatterpolarMarkerColorbarLenmode = "pixels"
)

// ScatterpolarMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterpolarMarkerColorbarShowexponent string

const (
	ScatterpolarMarkerColorbarShowexponentAll   ScatterpolarMarkerColorbarShowexponent = "all"
	ScatterpolarMarkerColorbarShowexponentFirst ScatterpolarMarkerColorbarShowexponent = "first"
	ScatterpolarMarkerColorbarShowexponentLast  ScatterpolarMarkerColorbarShowexponent = "last"
	ScatterpolarMarkerColorbarShowexponentNone  ScatterpolarMarkerColorbarShowexponent = "none"
)

// ScatterpolarMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterpolarMarkerColorbarShowtickprefix string

const (
	ScatterpolarMarkerColorbarShowtickprefixAll   ScatterpolarMarkerColorbarShowtickprefix = "all"
	ScatterpolarMarkerColorbarShowtickprefixFirst ScatterpolarMarkerColorbarShowtickprefix = "first"
	ScatterpolarMarkerColorbarShowtickprefixLast  ScatterpolarMarkerColorbarShowtickprefix = "last"
	ScatterpolarMarkerColorbarShowtickprefixNone  ScatterpolarMarkerColorbarShowtickprefix = "none"
)

// ScatterpolarMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterpolarMarkerColorbarShowticksuffix string

const (
	ScatterpolarMarkerColorbarShowticksuffixAll   ScatterpolarMarkerColorbarShowticksuffix = "all"
	ScatterpolarMarkerColorbarShowticksuffixFirst ScatterpolarMarkerColorbarShowticksuffix = "first"
	ScatterpolarMarkerColorbarShowticksuffixLast  ScatterpolarMarkerColorbarShowticksuffix = "last"
	ScatterpolarMarkerColorbarShowticksuffixNone  ScatterpolarMarkerColorbarShowticksuffix = "none"
)

// ScatterpolarMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterpolarMarkerColorbarThicknessmode string

const (
	ScatterpolarMarkerColorbarThicknessmodeFraction ScatterpolarMarkerColorbarThicknessmode = "fraction"
	ScatterpolarMarkerColorbarThicknessmodePixels   ScatterpolarMarkerColorbarThicknessmode = "pixels"
)

// ScatterpolarMarkerColorbarTicklabelposition Determines where tick labels are drawn.
type ScatterpolarMarkerColorbarTicklabelposition string

const (
	ScatterpolarMarkerColorbarTicklabelpositionOutside       ScatterpolarMarkerColorbarTicklabelposition = "outside"
	ScatterpolarMarkerColorbarTicklabelpositionInside        ScatterpolarMarkerColorbarTicklabelposition = "inside"
	ScatterpolarMarkerColorbarTicklabelpositionOutsideTop    ScatterpolarMarkerColorbarTicklabelposition = "outside top"
	ScatterpolarMarkerColorbarTicklabelpositionInsideTop     ScatterpolarMarkerColorbarTicklabelposition = "inside top"
	ScatterpolarMarkerColorbarTicklabelpositionOutsideBottom ScatterpolarMarkerColorbarTicklabelposition = "outside bottom"
	ScatterpolarMarkerColorbarTicklabelpositionInsideBottom  ScatterpolarMarkerColorbarTicklabelposition = "inside bottom"
)

// ScatterpolarMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterpolarMarkerColorbarTickmode string

const (
	ScatterpolarMarkerColorbarTickmodeAuto   ScatterpolarMarkerColorbarTickmode = "auto"
	ScatterpolarMarkerColorbarTickmodeLinear ScatterpolarMarkerColorbarTickmode = "linear"
	ScatterpolarMarkerColorbarTickmodeArray  ScatterpolarMarkerColorbarTickmode = "array"
)

// ScatterpolarMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterpolarMarkerColorbarTicks string

const (
	ScatterpolarMarkerColorbarTicksOutside ScatterpolarMarkerColorbarTicks = "outside"
	ScatterpolarMarkerColorbarTicksInside  ScatterpolarMarkerColorbarTicks = "inside"
	ScatterpolarMarkerColorbarTicksEmpty   ScatterpolarMarkerColorbarTicks = ""
)

// ScatterpolarMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterpolarMarkerColorbarTitleSide string

const (
	ScatterpolarMarkerColorbarTitleSideRight  ScatterpolarMarkerColorbarTitleSide = "right"
	ScatterpolarMarkerColorbarTitleSideTop    ScatterpolarMarkerColorbarTitleSide = "top"
	ScatterpolarMarkerColorbarTitleSideBottom ScatterpolarMarkerColorbarTitleSide = "bottom"
)

// ScatterpolarMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterpolarMarkerColorbarXanchor string

const (
	ScatterpolarMarkerColorbarXanchorLeft   ScatterpolarMarkerColorbarXanchor = "left"
	ScatterpolarMarkerColorbarXanchorCenter ScatterpolarMarkerColorbarXanchor = "center"
	ScatterpolarMarkerColorbarXanchorRight  ScatterpolarMarkerColorbarXanchor = "right"
)

// ScatterpolarMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterpolarMarkerColorbarYanchor string

const (
	ScatterpolarMarkerColorbarYanchorTop    ScatterpolarMarkerColorbarYanchor = "top"
	ScatterpolarMarkerColorbarYanchorMiddle ScatterpolarMarkerColorbarYanchor = "middle"
	ScatterpolarMarkerColorbarYanchorBottom ScatterpolarMarkerColorbarYanchor = "bottom"
)

// ScatterpolarMarkerGradientType Sets the type of gradient used to fill the markers
type ScatterpolarMarkerGradientType string

const (
	ScatterpolarMarkerGradientTypeRadial     ScatterpolarMarkerGradientType = "radial"
	ScatterpolarMarkerGradientTypeHorizontal ScatterpolarMarkerGradientType = "horizontal"
	ScatterpolarMarkerGradientTypeVertical   ScatterpolarMarkerGradientType = "vertical"
	ScatterpolarMarkerGradientTypeNone       ScatterpolarMarkerGradientType = "none"
)

// ScatterpolarMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterpolarMarkerSizemode string

const (
	ScatterpolarMarkerSizemodeDiameter ScatterpolarMarkerSizemode = "diameter"
	ScatterpolarMarkerSizemodeArea     ScatterpolarMarkerSizemode = "area"
)

// ScatterpolarMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterpolarMarkerSymbol interface{}

var (
	ScatterpolarMarkerSymbolNumber0                 ScatterpolarMarkerSymbol = 0
	ScatterpolarMarkerSymbol0                       ScatterpolarMarkerSymbol = "0"
	ScatterpolarMarkerSymbolCircle                  ScatterpolarMarkerSymbol = "circle"
	ScatterpolarMarkerSymbolNumber100               ScatterpolarMarkerSymbol = 100
	ScatterpolarMarkerSymbol100                     ScatterpolarMarkerSymbol = "100"
	ScatterpolarMarkerSymbolCircleOpen              ScatterpolarMarkerSymbol = "circle-open"
	ScatterpolarMarkerSymbolNumber200               ScatterpolarMarkerSymbol = 200
	ScatterpolarMarkerSymbol200                     ScatterpolarMarkerSymbol = "200"
	ScatterpolarMarkerSymbolCircleDot               ScatterpolarMarkerSymbol = "circle-dot"
	ScatterpolarMarkerSymbolNumber300               ScatterpolarMarkerSymbol = 300
	ScatterpolarMarkerSymbol300                     ScatterpolarMarkerSymbol = "300"
	ScatterpolarMarkerSymbolCircleOpenDot           ScatterpolarMarkerSymbol = "circle-open-dot"
	ScatterpolarMarkerSymbolNumber1                 ScatterpolarMarkerSymbol = 1
	ScatterpolarMarkerSymbol1                       ScatterpolarMarkerSymbol = "1"
	ScatterpolarMarkerSymbolSquare                  ScatterpolarMarkerSymbol = "square"
	ScatterpolarMarkerSymbolNumber101               ScatterpolarMarkerSymbol = 101
	ScatterpolarMarkerSymbol101                     ScatterpolarMarkerSymbol = "101"
	ScatterpolarMarkerSymbolSquareOpen              ScatterpolarMarkerSymbol = "square-open"
	ScatterpolarMarkerSymbolNumber201               ScatterpolarMarkerSymbol = 201
	ScatterpolarMarkerSymbol201                     ScatterpolarMarkerSymbol = "201"
	ScatterpolarMarkerSymbolSquareDot               ScatterpolarMarkerSymbol = "square-dot"
	ScatterpolarMarkerSymbolNumber301               ScatterpolarMarkerSymbol = 301
	ScatterpolarMarkerSymbol301                     ScatterpolarMarkerSymbol = "301"
	ScatterpolarMarkerSymbolSquareOpenDot           ScatterpolarMarkerSymbol = "square-open-dot"
	ScatterpolarMarkerSymbolNumber2                 ScatterpolarMarkerSymbol = 2
	ScatterpolarMarkerSymbol2                       ScatterpolarMarkerSymbol = "2"
	ScatterpolarMarkerSymbolDiamond                 ScatterpolarMarkerSymbol = "diamond"
	ScatterpolarMarkerSymbolNumber102               ScatterpolarMarkerSymbol = 102
	ScatterpolarMarkerSymbol102                     ScatterpolarMarkerSymbol = "102"
	ScatterpolarMarkerSymbolDiamondOpen             ScatterpolarMarkerSymbol = "diamond-open"
	ScatterpolarMarkerSymbolNumber202               ScatterpolarMarkerSymbol = 202
	ScatterpolarMarkerSymbol202                     ScatterpolarMarkerSymbol = "202"
	ScatterpolarMarkerSymbolDiamondDot              ScatterpolarMarkerSymbol = "diamond-dot"
	ScatterpolarMarkerSymbolNumber302               ScatterpolarMarkerSymbol = 302
	ScatterpolarMarkerSymbol302                     ScatterpolarMarkerSymbol = "302"
	ScatterpolarMarkerSymbolDiamondOpenDot          ScatterpolarMarkerSymbol = "diamond-open-dot"
	ScatterpolarMarkerSymbolNumber3                 ScatterpolarMarkerSymbol = 3
	ScatterpolarMarkerSymbol3                       ScatterpolarMarkerSymbol = "3"
	ScatterpolarMarkerSymbolCross                   ScatterpolarMarkerSymbol = "cross"
	ScatterpolarMarkerSymbolNumber103               ScatterpolarMarkerSymbol = 103
	ScatterpolarMarkerSymbol103                     ScatterpolarMarkerSymbol = "103"
	ScatterpolarMarkerSymbolCrossOpen               ScatterpolarMarkerSymbol = "cross-open"
	ScatterpolarMarkerSymbolNumber203               ScatterpolarMarkerSymbol = 203
	ScatterpolarMarkerSymbol203                     ScatterpolarMarkerSymbol = "203"
	ScatterpolarMarkerSymbolCrossDot                ScatterpolarMarkerSymbol = "cross-dot"
	ScatterpolarMarkerSymbolNumber303               ScatterpolarMarkerSymbol = 303
	ScatterpolarMarkerSymbol303                     ScatterpolarMarkerSymbol = "303"
	ScatterpolarMarkerSymbolCrossOpenDot            ScatterpolarMarkerSymbol = "cross-open-dot"
	ScatterpolarMarkerSymbolNumber4                 ScatterpolarMarkerSymbol = 4
	ScatterpolarMarkerSymbol4                       ScatterpolarMarkerSymbol = "4"
	ScatterpolarMarkerSymbolX                       ScatterpolarMarkerSymbol = "x"
	ScatterpolarMarkerSymbolNumber104               ScatterpolarMarkerSymbol = 104
	ScatterpolarMarkerSymbol104                     ScatterpolarMarkerSymbol = "104"
	ScatterpolarMarkerSymbolXOpen                   ScatterpolarMarkerSymbol = "x-open"
	ScatterpolarMarkerSymbolNumber204               ScatterpolarMarkerSymbol = 204
	ScatterpolarMarkerSymbol204                     ScatterpolarMarkerSymbol = "204"
	ScatterpolarMarkerSymbolXDot                    ScatterpolarMarkerSymbol = "x-dot"
	ScatterpolarMarkerSymbolNumber304               ScatterpolarMarkerSymbol = 304
	ScatterpolarMarkerSymbol304                     ScatterpolarMarkerSymbol = "304"
	ScatterpolarMarkerSymbolXOpenDot                ScatterpolarMarkerSymbol = "x-open-dot"
	ScatterpolarMarkerSymbolNumber5                 ScatterpolarMarkerSymbol = 5
	ScatterpolarMarkerSymbol5                       ScatterpolarMarkerSymbol = "5"
	ScatterpolarMarkerSymbolTriangleUp              ScatterpolarMarkerSymbol = "triangle-up"
	ScatterpolarMarkerSymbolNumber105               ScatterpolarMarkerSymbol = 105
	ScatterpolarMarkerSymbol105                     ScatterpolarMarkerSymbol = "105"
	ScatterpolarMarkerSymbolTriangleUpOpen          ScatterpolarMarkerSymbol = "triangle-up-open"
	ScatterpolarMarkerSymbolNumber205               ScatterpolarMarkerSymbol = 205
	ScatterpolarMarkerSymbol205                     ScatterpolarMarkerSymbol = "205"
	ScatterpolarMarkerSymbolTriangleUpDot           ScatterpolarMarkerSymbol = "triangle-up-dot"
	ScatterpolarMarkerSymbolNumber305               ScatterpolarMarkerSymbol = 305
	ScatterpolarMarkerSymbol305                     ScatterpolarMarkerSymbol = "305"
	ScatterpolarMarkerSymbolTriangleUpOpenDot       ScatterpolarMarkerSymbol = "triangle-up-open-dot"
	ScatterpolarMarkerSymbolNumber6                 ScatterpolarMarkerSymbol = 6
	ScatterpolarMarkerSymbol6                       ScatterpolarMarkerSymbol = "6"
	ScatterpolarMarkerSymbolTriangleDown            ScatterpolarMarkerSymbol = "triangle-down"
	ScatterpolarMarkerSymbolNumber106               ScatterpolarMarkerSymbol = 106
	ScatterpolarMarkerSymbol106                     ScatterpolarMarkerSymbol = "106"
	ScatterpolarMarkerSymbolTriangleDownOpen        ScatterpolarMarkerSymbol = "triangle-down-open"
	ScatterpolarMarkerSymbolNumber206               ScatterpolarMarkerSymbol = 206
	ScatterpolarMarkerSymbol206                     ScatterpolarMarkerSymbol = "206"
	ScatterpolarMarkerSymbolTriangleDownDot         ScatterpolarMarkerSymbol = "triangle-down-dot"
	ScatterpolarMarkerSymbolNumber306               ScatterpolarMarkerSymbol = 306
	ScatterpolarMarkerSymbol306                     ScatterpolarMarkerSymbol = "306"
	ScatterpolarMarkerSymbolTriangleDownOpenDot     ScatterpolarMarkerSymbol = "triangle-down-open-dot"
	ScatterpolarMarkerSymbolNumber7                 ScatterpolarMarkerSymbol = 7
	ScatterpolarMarkerSymbol7                       ScatterpolarMarkerSymbol = "7"
	ScatterpolarMarkerSymbolTriangleLeft            ScatterpolarMarkerSymbol = "triangle-left"
	ScatterpolarMarkerSymbolNumber107               ScatterpolarMarkerSymbol = 107
	ScatterpolarMarkerSymbol107                     ScatterpolarMarkerSymbol = "107"
	ScatterpolarMarkerSymbolTriangleLeftOpen        ScatterpolarMarkerSymbol = "triangle-left-open"
	ScatterpolarMarkerSymbolNumber207               ScatterpolarMarkerSymbol = 207
	ScatterpolarMarkerSymbol207                     ScatterpolarMarkerSymbol = "207"
	ScatterpolarMarkerSymbolTriangleLeftDot         ScatterpolarMarkerSymbol = "triangle-left-dot"
	ScatterpolarMarkerSymbolNumber307               ScatterpolarMarkerSymbol = 307
	ScatterpolarMarkerSymbol307                     ScatterpolarMarkerSymbol = "307"
	ScatterpolarMarkerSymbolTriangleLeftOpenDot     ScatterpolarMarkerSymbol = "triangle-left-open-dot"
	ScatterpolarMarkerSymbolNumber8                 ScatterpolarMarkerSymbol = 8
	ScatterpolarMarkerSymbol8                       ScatterpolarMarkerSymbol = "8"
	ScatterpolarMarkerSymbolTriangleRight           ScatterpolarMarkerSymbol = "triangle-right"
	ScatterpolarMarkerSymbolNumber108               ScatterpolarMarkerSymbol = 108
	ScatterpolarMarkerSymbol108                     ScatterpolarMarkerSymbol = "108"
	ScatterpolarMarkerSymbolTriangleRightOpen       ScatterpolarMarkerSymbol = "triangle-right-open"
	ScatterpolarMarkerSymbolNumber208               ScatterpolarMarkerSymbol = 208
	ScatterpolarMarkerSymbol208                     ScatterpolarMarkerSymbol = "208"
	ScatterpolarMarkerSymbolTriangleRightDot        ScatterpolarMarkerSymbol = "triangle-right-dot"
	ScatterpolarMarkerSymbolNumber308               ScatterpolarMarkerSymbol = 308
	ScatterpolarMarkerSymbol308                     ScatterpolarMarkerSymbol = "308"
	ScatterpolarMarkerSymbolTriangleRightOpenDot    ScatterpolarMarkerSymbol = "triangle-right-open-dot"
	ScatterpolarMarkerSymbolNumber9                 ScatterpolarMarkerSymbol = 9
	ScatterpolarMarkerSymbol9                       ScatterpolarMarkerSymbol = "9"
	ScatterpolarMarkerSymbolTriangleNe              ScatterpolarMarkerSymbol = "triangle-ne"
	ScatterpolarMarkerSymbolNumber109               ScatterpolarMarkerSymbol = 109
	ScatterpolarMarkerSymbol109                     ScatterpolarMarkerSymbol = "109"
	ScatterpolarMarkerSymbolTriangleNeOpen          ScatterpolarMarkerSymbol = "triangle-ne-open"
	ScatterpolarMarkerSymbolNumber209               ScatterpolarMarkerSymbol = 209
	ScatterpolarMarkerSymbol209                     ScatterpolarMarkerSymbol = "209"
	ScatterpolarMarkerSymbolTriangleNeDot           ScatterpolarMarkerSymbol = "triangle-ne-dot"
	ScatterpolarMarkerSymbolNumber309               ScatterpolarMarkerSymbol = 309
	ScatterpolarMarkerSymbol309                     ScatterpolarMarkerSymbol = "309"
	ScatterpolarMarkerSymbolTriangleNeOpenDot       ScatterpolarMarkerSymbol = "triangle-ne-open-dot"
	ScatterpolarMarkerSymbolNumber10                ScatterpolarMarkerSymbol = 10
	ScatterpolarMarkerSymbol10                      ScatterpolarMarkerSymbol = "10"
	ScatterpolarMarkerSymbolTriangleSe              ScatterpolarMarkerSymbol = "triangle-se"
	ScatterpolarMarkerSymbolNumber110               ScatterpolarMarkerSymbol = 110
	ScatterpolarMarkerSymbol110                     ScatterpolarMarkerSymbol = "110"
	ScatterpolarMarkerSymbolTriangleSeOpen          ScatterpolarMarkerSymbol = "triangle-se-open"
	ScatterpolarMarkerSymbolNumber210               ScatterpolarMarkerSymbol = 210
	ScatterpolarMarkerSymbol210                     ScatterpolarMarkerSymbol = "210"
	ScatterpolarMarkerSymbolTriangleSeDot           ScatterpolarMarkerSymbol = "triangle-se-dot"
	ScatterpolarMarkerSymbolNumber310               ScatterpolarMarkerSymbol = 310
	ScatterpolarMarkerSymbol310                     ScatterpolarMarkerSymbol = "310"
	ScatterpolarMarkerSymbolTriangleSeOpenDot       ScatterpolarMarkerSymbol = "triangle-se-open-dot"
	ScatterpolarMarkerSymbolNumber11                ScatterpolarMarkerSymbol = 11
	ScatterpolarMarkerSymbol11                      ScatterpolarMarkerSymbol = "11"
	ScatterpolarMarkerSymbolTriangleSw              ScatterpolarMarkerSymbol = "triangle-sw"
	ScatterpolarMarkerSymbolNumber111               ScatterpolarMarkerSymbol = 111
	ScatterpolarMarkerSymbol111                     ScatterpolarMarkerSymbol = "111"
	ScatterpolarMarkerSymbolTriangleSwOpen          ScatterpolarMarkerSymbol = "triangle-sw-open"
	ScatterpolarMarkerSymbolNumber211               ScatterpolarMarkerSymbol = 211
	ScatterpolarMarkerSymbol211                     ScatterpolarMarkerSymbol = "211"
	ScatterpolarMarkerSymbolTriangleSwDot           ScatterpolarMarkerSymbol = "triangle-sw-dot"
	ScatterpolarMarkerSymbolNumber311               ScatterpolarMarkerSymbol = 311
	ScatterpolarMarkerSymbol311                     ScatterpolarMarkerSymbol = "311"
	ScatterpolarMarkerSymbolTriangleSwOpenDot       ScatterpolarMarkerSymbol = "triangle-sw-open-dot"
	ScatterpolarMarkerSymbolNumber12                ScatterpolarMarkerSymbol = 12
	ScatterpolarMarkerSymbol12                      ScatterpolarMarkerSymbol = "12"
	ScatterpolarMarkerSymbolTriangleNw              ScatterpolarMarkerSymbol = "triangle-nw"
	ScatterpolarMarkerSymbolNumber112               ScatterpolarMarkerSymbol = 112
	ScatterpolarMarkerSymbol112                     ScatterpolarMarkerSymbol = "112"
	ScatterpolarMarkerSymbolTriangleNwOpen          ScatterpolarMarkerSymbol = "triangle-nw-open"
	ScatterpolarMarkerSymbolNumber212               ScatterpolarMarkerSymbol = 212
	ScatterpolarMarkerSymbol212                     ScatterpolarMarkerSymbol = "212"
	ScatterpolarMarkerSymbolTriangleNwDot           ScatterpolarMarkerSymbol = "triangle-nw-dot"
	ScatterpolarMarkerSymbolNumber312               ScatterpolarMarkerSymbol = 312
	ScatterpolarMarkerSymbol312                     ScatterpolarMarkerSymbol = "312"
	ScatterpolarMarkerSymbolTriangleNwOpenDot       ScatterpolarMarkerSymbol = "triangle-nw-open-dot"
	ScatterpolarMarkerSymbolNumber13                ScatterpolarMarkerSymbol = 13
	ScatterpolarMarkerSymbol13                      ScatterpolarMarkerSymbol = "13"
	ScatterpolarMarkerSymbolPentagon                ScatterpolarMarkerSymbol = "pentagon"
	ScatterpolarMarkerSymbolNumber113               ScatterpolarMarkerSymbol = 113
	ScatterpolarMarkerSymbol113                     ScatterpolarMarkerSymbol = "113"
	ScatterpolarMarkerSymbolPentagonOpen            ScatterpolarMarkerSymbol = "pentagon-open"
	ScatterpolarMarkerSymbolNumber213               ScatterpolarMarkerSymbol = 213
	ScatterpolarMarkerSymbol213                     ScatterpolarMarkerSymbol = "213"
	ScatterpolarMarkerSymbolPentagonDot             ScatterpolarMarkerSymbol = "pentagon-dot"
	ScatterpolarMarkerSymbolNumber313               ScatterpolarMarkerSymbol = 313
	ScatterpolarMarkerSymbol313                     ScatterpolarMarkerSymbol = "313"
	ScatterpolarMarkerSymbolPentagonOpenDot         ScatterpolarMarkerSymbol = "pentagon-open-dot"
	ScatterpolarMarkerSymbolNumber14                ScatterpolarMarkerSymbol = 14
	ScatterpolarMarkerSymbol14                      ScatterpolarMarkerSymbol = "14"
	ScatterpolarMarkerSymbolHexagon                 ScatterpolarMarkerSymbol = "hexagon"
	ScatterpolarMarkerSymbolNumber114               ScatterpolarMarkerSymbol = 114
	ScatterpolarMarkerSymbol114                     ScatterpolarMarkerSymbol = "114"
	ScatterpolarMarkerSymbolHexagonOpen             ScatterpolarMarkerSymbol = "hexagon-open"
	ScatterpolarMarkerSymbolNumber214               ScatterpolarMarkerSymbol = 214
	ScatterpolarMarkerSymbol214                     ScatterpolarMarkerSymbol = "214"
	ScatterpolarMarkerSymbolHexagonDot              ScatterpolarMarkerSymbol = "hexagon-dot"
	ScatterpolarMarkerSymbolNumber314               ScatterpolarMarkerSymbol = 314
	ScatterpolarMarkerSymbol314                     ScatterpolarMarkerSymbol = "314"
	ScatterpolarMarkerSymbolHexagonOpenDot          ScatterpolarMarkerSymbol = "hexagon-open-dot"
	ScatterpolarMarkerSymbolNumber15                ScatterpolarMarkerSymbol = 15
	ScatterpolarMarkerSymbol15                      ScatterpolarMarkerSymbol = "15"
	ScatterpolarMarkerSymbolHexagon2                ScatterpolarMarkerSymbol = "hexagon2"
	ScatterpolarMarkerSymbolNumber115               ScatterpolarMarkerSymbol = 115
	ScatterpolarMarkerSymbol115                     ScatterpolarMarkerSymbol = "115"
	ScatterpolarMarkerSymbolHexagon2Open            ScatterpolarMarkerSymbol = "hexagon2-open"
	ScatterpolarMarkerSymbolNumber215               ScatterpolarMarkerSymbol = 215
	ScatterpolarMarkerSymbol215                     ScatterpolarMarkerSymbol = "215"
	ScatterpolarMarkerSymbolHexagon2Dot             ScatterpolarMarkerSymbol = "hexagon2-dot"
	ScatterpolarMarkerSymbolNumber315               ScatterpolarMarkerSymbol = 315
	ScatterpolarMarkerSymbol315                     ScatterpolarMarkerSymbol = "315"
	ScatterpolarMarkerSymbolHexagon2OpenDot         ScatterpolarMarkerSymbol = "hexagon2-open-dot"
	ScatterpolarMarkerSymbolNumber16                ScatterpolarMarkerSymbol = 16
	ScatterpolarMarkerSymbol16                      ScatterpolarMarkerSymbol = "16"
	ScatterpolarMarkerSymbolOctagon                 ScatterpolarMarkerSymbol = "octagon"
	ScatterpolarMarkerSymbolNumber116               ScatterpolarMarkerSymbol = 116
	ScatterpolarMarkerSymbol116                     ScatterpolarMarkerSymbol = "116"
	ScatterpolarMarkerSymbolOctagonOpen             ScatterpolarMarkerSymbol = "octagon-open"
	ScatterpolarMarkerSymbolNumber216               ScatterpolarMarkerSymbol = 216
	ScatterpolarMarkerSymbol216                     ScatterpolarMarkerSymbol = "216"
	ScatterpolarMarkerSymbolOctagonDot              ScatterpolarMarkerSymbol = "octagon-dot"
	ScatterpolarMarkerSymbolNumber316               ScatterpolarMarkerSymbol = 316
	ScatterpolarMarkerSymbol316                     ScatterpolarMarkerSymbol = "316"
	ScatterpolarMarkerSymbolOctagonOpenDot          ScatterpolarMarkerSymbol = "octagon-open-dot"
	ScatterpolarMarkerSymbolNumber17                ScatterpolarMarkerSymbol = 17
	ScatterpolarMarkerSymbol17                      ScatterpolarMarkerSymbol = "17"
	ScatterpolarMarkerSymbolStar                    ScatterpolarMarkerSymbol = "star"
	ScatterpolarMarkerSymbolNumber117               ScatterpolarMarkerSymbol = 117
	ScatterpolarMarkerSymbol117                     ScatterpolarMarkerSymbol = "117"
	ScatterpolarMarkerSymbolStarOpen                ScatterpolarMarkerSymbol = "star-open"
	ScatterpolarMarkerSymbolNumber217               ScatterpolarMarkerSymbol = 217
	ScatterpolarMarkerSymbol217                     ScatterpolarMarkerSymbol = "217"
	ScatterpolarMarkerSymbolStarDot                 ScatterpolarMarkerSymbol = "star-dot"
	ScatterpolarMarkerSymbolNumber317               ScatterpolarMarkerSymbol = 317
	ScatterpolarMarkerSymbol317                     ScatterpolarMarkerSymbol = "317"
	ScatterpolarMarkerSymbolStarOpenDot             ScatterpolarMarkerSymbol = "star-open-dot"
	ScatterpolarMarkerSymbolNumber18                ScatterpolarMarkerSymbol = 18
	ScatterpolarMarkerSymbol18                      ScatterpolarMarkerSymbol = "18"
	ScatterpolarMarkerSymbolHexagram                ScatterpolarMarkerSymbol = "hexagram"
	ScatterpolarMarkerSymbolNumber118               ScatterpolarMarkerSymbol = 118
	ScatterpolarMarkerSymbol118                     ScatterpolarMarkerSymbol = "118"
	ScatterpolarMarkerSymbolHexagramOpen            ScatterpolarMarkerSymbol = "hexagram-open"
	ScatterpolarMarkerSymbolNumber218               ScatterpolarMarkerSymbol = 218
	ScatterpolarMarkerSymbol218                     ScatterpolarMarkerSymbol = "218"
	ScatterpolarMarkerSymbolHexagramDot             ScatterpolarMarkerSymbol = "hexagram-dot"
	ScatterpolarMarkerSymbolNumber318               ScatterpolarMarkerSymbol = 318
	ScatterpolarMarkerSymbol318                     ScatterpolarMarkerSymbol = "318"
	ScatterpolarMarkerSymbolHexagramOpenDot         ScatterpolarMarkerSymbol = "hexagram-open-dot"
	ScatterpolarMarkerSymbolNumber19                ScatterpolarMarkerSymbol = 19
	ScatterpolarMarkerSymbol19                      ScatterpolarMarkerSymbol = "19"
	ScatterpolarMarkerSymbolStarTriangleUp          ScatterpolarMarkerSymbol = "star-triangle-up"
	ScatterpolarMarkerSymbolNumber119               ScatterpolarMarkerSymbol = 119
	ScatterpolarMarkerSymbol119                     ScatterpolarMarkerSymbol = "119"
	ScatterpolarMarkerSymbolStarTriangleUpOpen      ScatterpolarMarkerSymbol = "star-triangle-up-open"
	ScatterpolarMarkerSymbolNumber219               ScatterpolarMarkerSymbol = 219
	ScatterpolarMarkerSymbol219                     ScatterpolarMarkerSymbol = "219"
	ScatterpolarMarkerSymbolStarTriangleUpDot       ScatterpolarMarkerSymbol = "star-triangle-up-dot"
	ScatterpolarMarkerSymbolNumber319               ScatterpolarMarkerSymbol = 319
	ScatterpolarMarkerSymbol319                     ScatterpolarMarkerSymbol = "319"
	ScatterpolarMarkerSymbolStarTriangleUpOpenDot   ScatterpolarMarkerSymbol = "star-triangle-up-open-dot"
	ScatterpolarMarkerSymbolNumber20                ScatterpolarMarkerSymbol = 20
	ScatterpolarMarkerSymbol20                      ScatterpolarMarkerSymbol = "20"
	ScatterpolarMarkerSymbolStarTriangleDown        ScatterpolarMarkerSymbol = "star-triangle-down"
	ScatterpolarMarkerSymbolNumber120               ScatterpolarMarkerSymbol = 120
	ScatterpolarMarkerSymbol120                     ScatterpolarMarkerSymbol = "120"
	ScatterpolarMarkerSymbolStarTriangleDownOpen    ScatterpolarMarkerSymbol = "star-triangle-down-open"
	ScatterpolarMarkerSymbolNumber220               ScatterpolarMarkerSymbol = 220
	ScatterpolarMarkerSymbol220                     ScatterpolarMarkerSymbol = "220"
	ScatterpolarMarkerSymbolStarTriangleDownDot     ScatterpolarMarkerSymbol = "star-triangle-down-dot"
	ScatterpolarMarkerSymbolNumber320               ScatterpolarMarkerSymbol = 320
	ScatterpolarMarkerSymbol320                     ScatterpolarMarkerSymbol = "320"
	ScatterpolarMarkerSymbolStarTriangleDownOpenDot ScatterpolarMarkerSymbol = "star-triangle-down-open-dot"
	ScatterpolarMarkerSymbolNumber21                ScatterpolarMarkerSymbol = 21
	ScatterpolarMarkerSymbol21                      ScatterpolarMarkerSymbol = "21"
	ScatterpolarMarkerSymbolStarSquare              ScatterpolarMarkerSymbol = "star-square"
	ScatterpolarMarkerSymbolNumber121               ScatterpolarMarkerSymbol = 121
	ScatterpolarMarkerSymbol121                     ScatterpolarMarkerSymbol = "121"
	ScatterpolarMarkerSymbolStarSquareOpen          ScatterpolarMarkerSymbol = "star-square-open"
	ScatterpolarMarkerSymbolNumber221               ScatterpolarMarkerSymbol = 221
	ScatterpolarMarkerSymbol221                     ScatterpolarMarkerSymbol = "221"
	ScatterpolarMarkerSymbolStarSquareDot           ScatterpolarMarkerSymbol = "star-square-dot"
	ScatterpolarMarkerSymbolNumber321               ScatterpolarMarkerSymbol = 321
	ScatterpolarMarkerSymbol321                     ScatterpolarMarkerSymbol = "321"
	ScatterpolarMarkerSymbolStarSquareOpenDot       ScatterpolarMarkerSymbol = "star-square-open-dot"
	ScatterpolarMarkerSymbolNumber22                ScatterpolarMarkerSymbol = 22
	ScatterpolarMarkerSymbol22                      ScatterpolarMarkerSymbol = "22"
	ScatterpolarMarkerSymbolStarDiamond             ScatterpolarMarkerSymbol = "star-diamond"
	ScatterpolarMarkerSymbolNumber122               ScatterpolarMarkerSymbol = 122
	ScatterpolarMarkerSymbol122                     ScatterpolarMarkerSymbol = "122"
	ScatterpolarMarkerSymbolStarDiamondOpen         ScatterpolarMarkerSymbol = "star-diamond-open"
	ScatterpolarMarkerSymbolNumber222               ScatterpolarMarkerSymbol = 222
	ScatterpolarMarkerSymbol222                     ScatterpolarMarkerSymbol = "222"
	ScatterpolarMarkerSymbolStarDiamondDot          ScatterpolarMarkerSymbol = "star-diamond-dot"
	ScatterpolarMarkerSymbolNumber322               ScatterpolarMarkerSymbol = 322
	ScatterpolarMarkerSymbol322                     ScatterpolarMarkerSymbol = "322"
	ScatterpolarMarkerSymbolStarDiamondOpenDot      ScatterpolarMarkerSymbol = "star-diamond-open-dot"
	ScatterpolarMarkerSymbolNumber23                ScatterpolarMarkerSymbol = 23
	ScatterpolarMarkerSymbol23                      ScatterpolarMarkerSymbol = "23"
	ScatterpolarMarkerSymbolDiamondTall             ScatterpolarMarkerSymbol = "diamond-tall"
	ScatterpolarMarkerSymbolNumber123               ScatterpolarMarkerSymbol = 123
	ScatterpolarMarkerSymbol123                     ScatterpolarMarkerSymbol = "123"
	ScatterpolarMarkerSymbolDiamondTallOpen         ScatterpolarMarkerSymbol = "diamond-tall-open"
	ScatterpolarMarkerSymbolNumber223               ScatterpolarMarkerSymbol = 223
	ScatterpolarMarkerSymbol223                     ScatterpolarMarkerSymbol = "223"
	ScatterpolarMarkerSymbolDiamondTallDot          ScatterpolarMarkerSymbol = "diamond-tall-dot"
	ScatterpolarMarkerSymbolNumber323               ScatterpolarMarkerSymbol = 323
	ScatterpolarMarkerSymbol323                     ScatterpolarMarkerSymbol = "323"
	ScatterpolarMarkerSymbolDiamondTallOpenDot      ScatterpolarMarkerSymbol = "diamond-tall-open-dot"
	ScatterpolarMarkerSymbolNumber24                ScatterpolarMarkerSymbol = 24
	ScatterpolarMarkerSymbol24                      ScatterpolarMarkerSymbol = "24"
	ScatterpolarMarkerSymbolDiamondWide             ScatterpolarMarkerSymbol = "diamond-wide"
	ScatterpolarMarkerSymbolNumber124               ScatterpolarMarkerSymbol = 124
	ScatterpolarMarkerSymbol124                     ScatterpolarMarkerSymbol = "124"
	ScatterpolarMarkerSymbolDiamondWideOpen         ScatterpolarMarkerSymbol = "diamond-wide-open"
	ScatterpolarMarkerSymbolNumber224               ScatterpolarMarkerSymbol = 224
	ScatterpolarMarkerSymbol224                     ScatterpolarMarkerSymbol = "224"
	ScatterpolarMarkerSymbolDiamondWideDot          ScatterpolarMarkerSymbol = "diamond-wide-dot"
	ScatterpolarMarkerSymbolNumber324               ScatterpolarMarkerSymbol = 324
	ScatterpolarMarkerSymbol324                     ScatterpolarMarkerSymbol = "324"
	ScatterpolarMarkerSymbolDiamondWideOpenDot      ScatterpolarMarkerSymbol = "diamond-wide-open-dot"
	ScatterpolarMarkerSymbolNumber25                ScatterpolarMarkerSymbol = 25
	ScatterpolarMarkerSymbol25                      ScatterpolarMarkerSymbol = "25"
	ScatterpolarMarkerSymbolHourglass               ScatterpolarMarkerSymbol = "hourglass"
	ScatterpolarMarkerSymbolNumber125               ScatterpolarMarkerSymbol = 125
	ScatterpolarMarkerSymbol125                     ScatterpolarMarkerSymbol = "125"
	ScatterpolarMarkerSymbolHourglassOpen           ScatterpolarMarkerSymbol = "hourglass-open"
	ScatterpolarMarkerSymbolNumber26                ScatterpolarMarkerSymbol = 26
	ScatterpolarMarkerSymbol26                      ScatterpolarMarkerSymbol = "26"
	ScatterpolarMarkerSymbolBowtie                  ScatterpolarMarkerSymbol = "bowtie"
	ScatterpolarMarkerSymbolNumber126               ScatterpolarMarkerSymbol = 126
	ScatterpolarMarkerSymbol126                     ScatterpolarMarkerSymbol = "126"
	ScatterpolarMarkerSymbolBowtieOpen              ScatterpolarMarkerSymbol = "bowtie-open"
	ScatterpolarMarkerSymbolNumber27                ScatterpolarMarkerSymbol = 27
	ScatterpolarMarkerSymbol27                      ScatterpolarMarkerSymbol = "27"
	ScatterpolarMarkerSymbolCircleCross             ScatterpolarMarkerSymbol = "circle-cross"
	ScatterpolarMarkerSymbolNumber127               ScatterpolarMarkerSymbol = 127
	ScatterpolarMarkerSymbol127                     ScatterpolarMarkerSymbol = "127"
	ScatterpolarMarkerSymbolCircleCrossOpen         ScatterpolarMarkerSymbol = "circle-cross-open"
	ScatterpolarMarkerSymbolNumber28                ScatterpolarMarkerSymbol = 28
	ScatterpolarMarkerSymbol28                      ScatterpolarMarkerSymbol = "28"
	ScatterpolarMarkerSymbolCircleX                 ScatterpolarMarkerSymbol = "circle-x"
	ScatterpolarMarkerSymbolNumber128               ScatterpolarMarkerSymbol = 128
	ScatterpolarMarkerSymbol128                     ScatterpolarMarkerSymbol = "128"
	ScatterpolarMarkerSymbolCircleXOpen             ScatterpolarMarkerSymbol = "circle-x-open"
	ScatterpolarMarkerSymbolNumber29                ScatterpolarMarkerSymbol = 29
	ScatterpolarMarkerSymbol29                      ScatterpolarMarkerSymbol = "29"
	ScatterpolarMarkerSymbolSquareCross             ScatterpolarMarkerSymbol = "square-cross"
	ScatterpolarMarkerSymbolNumber129               ScatterpolarMarkerSymbol = 129
	ScatterpolarMarkerSymbol129                     ScatterpolarMarkerSymbol = "129"
	ScatterpolarMarkerSymbolSquareCrossOpen         ScatterpolarMarkerSymbol = "square-cross-open"
	ScatterpolarMarkerSymbolNumber30                ScatterpolarMarkerSymbol = 30
	ScatterpolarMarkerSymbol30                      ScatterpolarMarkerSymbol = "30"
	ScatterpolarMarkerSymbolSquareX                 ScatterpolarMarkerSymbol = "square-x"
	ScatterpolarMarkerSymbolNumber130               ScatterpolarMarkerSymbol = 130
	ScatterpolarMarkerSymbol130                     ScatterpolarMarkerSymbol = "130"
	ScatterpolarMarkerSymbolSquareXOpen             ScatterpolarMarkerSymbol = "square-x-open"
	ScatterpolarMarkerSymbolNumber31                ScatterpolarMarkerSymbol = 31
	ScatterpolarMarkerSymbol31                      ScatterpolarMarkerSymbol = "31"
	ScatterpolarMarkerSymbolDiamondCross            ScatterpolarMarkerSymbol = "diamond-cross"
	ScatterpolarMarkerSymbolNumber131               ScatterpolarMarkerSymbol = 131
	ScatterpolarMarkerSymbol131                     ScatterpolarMarkerSymbol = "131"
	ScatterpolarMarkerSymbolDiamondCrossOpen        ScatterpolarMarkerSymbol = "diamond-cross-open"
	ScatterpolarMarkerSymbolNumber32                ScatterpolarMarkerSymbol = 32
	ScatterpolarMarkerSymbol32                      ScatterpolarMarkerSymbol = "32"
	ScatterpolarMarkerSymbolDiamondX                ScatterpolarMarkerSymbol = "diamond-x"
	ScatterpolarMarkerSymbolNumber132               ScatterpolarMarkerSymbol = 132
	ScatterpolarMarkerSymbol132                     ScatterpolarMarkerSymbol = "132"
	ScatterpolarMarkerSymbolDiamondXOpen            ScatterpolarMarkerSymbol = "diamond-x-open"
	ScatterpolarMarkerSymbolNumber33                ScatterpolarMarkerSymbol = 33
	ScatterpolarMarkerSymbol33                      ScatterpolarMarkerSymbol = "33"
	ScatterpolarMarkerSymbolCrossThin               ScatterpolarMarkerSymbol = "cross-thin"
	ScatterpolarMarkerSymbolNumber133               ScatterpolarMarkerSymbol = 133
	ScatterpolarMarkerSymbol133                     ScatterpolarMarkerSymbol = "133"
	ScatterpolarMarkerSymbolCrossThinOpen           ScatterpolarMarkerSymbol = "cross-thin-open"
	ScatterpolarMarkerSymbolNumber34                ScatterpolarMarkerSymbol = 34
	ScatterpolarMarkerSymbol34                      ScatterpolarMarkerSymbol = "34"
	ScatterpolarMarkerSymbolXThin                   ScatterpolarMarkerSymbol = "x-thin"
	ScatterpolarMarkerSymbolNumber134               ScatterpolarMarkerSymbol = 134
	ScatterpolarMarkerSymbol134                     ScatterpolarMarkerSymbol = "134"
	ScatterpolarMarkerSymbolXThinOpen               ScatterpolarMarkerSymbol = "x-thin-open"
	ScatterpolarMarkerSymbolNumber35                ScatterpolarMarkerSymbol = 35
	ScatterpolarMarkerSymbol35                      ScatterpolarMarkerSymbol = "35"
	ScatterpolarMarkerSymbolAsterisk                ScatterpolarMarkerSymbol = "asterisk"
	ScatterpolarMarkerSymbolNumber135               ScatterpolarMarkerSymbol = 135
	ScatterpolarMarkerSymbol135                     ScatterpolarMarkerSymbol = "135"
	ScatterpolarMarkerSymbolAsteriskOpen            ScatterpolarMarkerSymbol = "asterisk-open"
	ScatterpolarMarkerSymbolNumber36                ScatterpolarMarkerSymbol = 36
	ScatterpolarMarkerSymbol36                      ScatterpolarMarkerSymbol = "36"
	ScatterpolarMarkerSymbolHash                    ScatterpolarMarkerSymbol = "hash"
	ScatterpolarMarkerSymbolNumber136               ScatterpolarMarkerSymbol = 136
	ScatterpolarMarkerSymbol136                     ScatterpolarMarkerSymbol = "136"
	ScatterpolarMarkerSymbolHashOpen                ScatterpolarMarkerSymbol = "hash-open"
	ScatterpolarMarkerSymbolNumber236               ScatterpolarMarkerSymbol = 236
	ScatterpolarMarkerSymbol236                     ScatterpolarMarkerSymbol = "236"
	ScatterpolarMarkerSymbolHashDot                 ScatterpolarMarkerSymbol = "hash-dot"
	ScatterpolarMarkerSymbolNumber336               ScatterpolarMarkerSymbol = 336
	ScatterpolarMarkerSymbol336                     ScatterpolarMarkerSymbol = "336"
	ScatterpolarMarkerSymbolHashOpenDot             ScatterpolarMarkerSymbol = "hash-open-dot"
	ScatterpolarMarkerSymbolNumber37                ScatterpolarMarkerSymbol = 37
	ScatterpolarMarkerSymbol37                      ScatterpolarMarkerSymbol = "37"
	ScatterpolarMarkerSymbolYUp                     ScatterpolarMarkerSymbol = "y-up"
	ScatterpolarMarkerSymbolNumber137               ScatterpolarMarkerSymbol = 137
	ScatterpolarMarkerSymbol137                     ScatterpolarMarkerSymbol = "137"
	ScatterpolarMarkerSymbolYUpOpen                 ScatterpolarMarkerSymbol = "y-up-open"
	ScatterpolarMarkerSymbolNumber38                ScatterpolarMarkerSymbol = 38
	ScatterpolarMarkerSymbol38                      ScatterpolarMarkerSymbol = "38"
	ScatterpolarMarkerSymbolYDown                   ScatterpolarMarkerSymbol = "y-down"
	ScatterpolarMarkerSymbolNumber138               ScatterpolarMarkerSymbol = 138
	ScatterpolarMarkerSymbol138                     ScatterpolarMarkerSymbol = "138"
	ScatterpolarMarkerSymbolYDownOpen               ScatterpolarMarkerSymbol = "y-down-open"
	ScatterpolarMarkerSymbolNumber39                ScatterpolarMarkerSymbol = 39
	ScatterpolarMarkerSymbol39                      ScatterpolarMarkerSymbol = "39"
	ScatterpolarMarkerSymbolYLeft                   ScatterpolarMarkerSymbol = "y-left"
	ScatterpolarMarkerSymbolNumber139               ScatterpolarMarkerSymbol = 139
	ScatterpolarMarkerSymbol139                     ScatterpolarMarkerSymbol = "139"
	ScatterpolarMarkerSymbolYLeftOpen               ScatterpolarMarkerSymbol = "y-left-open"
	ScatterpolarMarkerSymbolNumber40                ScatterpolarMarkerSymbol = 40
	ScatterpolarMarkerSymbol40                      ScatterpolarMarkerSymbol = "40"
	ScatterpolarMarkerSymbolYRight                  ScatterpolarMarkerSymbol = "y-right"
	ScatterpolarMarkerSymbolNumber140               ScatterpolarMarkerSymbol = 140
	ScatterpolarMarkerSymbol140                     ScatterpolarMarkerSymbol = "140"
	ScatterpolarMarkerSymbolYRightOpen              ScatterpolarMarkerSymbol = "y-right-open"
	ScatterpolarMarkerSymbolNumber41                ScatterpolarMarkerSymbol = 41
	ScatterpolarMarkerSymbol41                      ScatterpolarMarkerSymbol = "41"
	ScatterpolarMarkerSymbolLineEw                  ScatterpolarMarkerSymbol = "line-ew"
	ScatterpolarMarkerSymbolNumber141               ScatterpolarMarkerSymbol = 141
	ScatterpolarMarkerSymbol141                     ScatterpolarMarkerSymbol = "141"
	ScatterpolarMarkerSymbolLineEwOpen              ScatterpolarMarkerSymbol = "line-ew-open"
	ScatterpolarMarkerSymbolNumber42                ScatterpolarMarkerSymbol = 42
	ScatterpolarMarkerSymbol42                      ScatterpolarMarkerSymbol = "42"
	ScatterpolarMarkerSymbolLineNs                  ScatterpolarMarkerSymbol = "line-ns"
	ScatterpolarMarkerSymbolNumber142               ScatterpolarMarkerSymbol = 142
	ScatterpolarMarkerSymbol142                     ScatterpolarMarkerSymbol = "142"
	ScatterpolarMarkerSymbolLineNsOpen              ScatterpolarMarkerSymbol = "line-ns-open"
	ScatterpolarMarkerSymbolNumber43                ScatterpolarMarkerSymbol = 43
	ScatterpolarMarkerSymbol43                      ScatterpolarMarkerSymbol = "43"
	ScatterpolarMarkerSymbolLineNe                  ScatterpolarMarkerSymbol = "line-ne"
	ScatterpolarMarkerSymbolNumber143               ScatterpolarMarkerSymbol = 143
	ScatterpolarMarkerSymbol143                     ScatterpolarMarkerSymbol = "143"
	ScatterpolarMarkerSymbolLineNeOpen              ScatterpolarMarkerSymbol = "line-ne-open"
	ScatterpolarMarkerSymbolNumber44                ScatterpolarMarkerSymbol = 44
	ScatterpolarMarkerSymbol44                      ScatterpolarMarkerSymbol = "44"
	ScatterpolarMarkerSymbolLineNw                  ScatterpolarMarkerSymbol = "line-nw"
	ScatterpolarMarkerSymbolNumber144               ScatterpolarMarkerSymbol = 144
	ScatterpolarMarkerSymbol144                     ScatterpolarMarkerSymbol = "144"
	ScatterpolarMarkerSymbolLineNwOpen              ScatterpolarMarkerSymbol = "line-nw-open"
	ScatterpolarMarkerSymbolNumber45                ScatterpolarMarkerSymbol = 45
	ScatterpolarMarkerSymbol45                      ScatterpolarMarkerSymbol = "45"
	ScatterpolarMarkerSymbolArrowUp                 ScatterpolarMarkerSymbol = "arrow-up"
	ScatterpolarMarkerSymbolNumber145               ScatterpolarMarkerSymbol = 145
	ScatterpolarMarkerSymbol145                     ScatterpolarMarkerSymbol = "145"
	ScatterpolarMarkerSymbolArrowUpOpen             ScatterpolarMarkerSymbol = "arrow-up-open"
	ScatterpolarMarkerSymbolNumber46                ScatterpolarMarkerSymbol = 46
	ScatterpolarMarkerSymbol46                      ScatterpolarMarkerSymbol = "46"
	ScatterpolarMarkerSymbolArrowDown               ScatterpolarMarkerSymbol = "arrow-down"
	ScatterpolarMarkerSymbolNumber146               ScatterpolarMarkerSymbol = 146
	ScatterpolarMarkerSymbol146                     ScatterpolarMarkerSymbol = "146"
	ScatterpolarMarkerSymbolArrowDownOpen           ScatterpolarMarkerSymbol = "arrow-down-open"
	ScatterpolarMarkerSymbolNumber47                ScatterpolarMarkerSymbol = 47
	ScatterpolarMarkerSymbol47                      ScatterpolarMarkerSymbol = "47"
	ScatterpolarMarkerSymbolArrowLeft               ScatterpolarMarkerSymbol = "arrow-left"
	ScatterpolarMarkerSymbolNumber147               ScatterpolarMarkerSymbol = 147
	ScatterpolarMarkerSymbol147                     ScatterpolarMarkerSymbol = "147"
	ScatterpolarMarkerSymbolArrowLeftOpen           ScatterpolarMarkerSymbol = "arrow-left-open"
	ScatterpolarMarkerSymbolNumber48                ScatterpolarMarkerSymbol = 48
	ScatterpolarMarkerSymbol48                      ScatterpolarMarkerSymbol = "48"
	ScatterpolarMarkerSymbolArrowRight              ScatterpolarMarkerSymbol = "arrow-right"
	ScatterpolarMarkerSymbolNumber148               ScatterpolarMarkerSymbol = 148
	ScatterpolarMarkerSymbol148                     ScatterpolarMarkerSymbol = "148"
	ScatterpolarMarkerSymbolArrowRightOpen          ScatterpolarMarkerSymbol = "arrow-right-open"
	ScatterpolarMarkerSymbolNumber49                ScatterpolarMarkerSymbol = 49
	ScatterpolarMarkerSymbol49                      ScatterpolarMarkerSymbol = "49"
	ScatterpolarMarkerSymbolArrowBarUp              ScatterpolarMarkerSymbol = "arrow-bar-up"
	ScatterpolarMarkerSymbolNumber149               ScatterpolarMarkerSymbol = 149
	ScatterpolarMarkerSymbol149                     ScatterpolarMarkerSymbol = "149"
	ScatterpolarMarkerSymbolArrowBarUpOpen          ScatterpolarMarkerSymbol = "arrow-bar-up-open"
	ScatterpolarMarkerSymbolNumber50                ScatterpolarMarkerSymbol = 50
	ScatterpolarMarkerSymbol50                      ScatterpolarMarkerSymbol = "50"
	ScatterpolarMarkerSymbolArrowBarDown            ScatterpolarMarkerSymbol = "arrow-bar-down"
	ScatterpolarMarkerSymbolNumber150               ScatterpolarMarkerSymbol = 150
	ScatterpolarMarkerSymbol150                     ScatterpolarMarkerSymbol = "150"
	ScatterpolarMarkerSymbolArrowBarDownOpen        ScatterpolarMarkerSymbol = "arrow-bar-down-open"
	ScatterpolarMarkerSymbolNumber51                ScatterpolarMarkerSymbol = 51
	ScatterpolarMarkerSymbol51                      ScatterpolarMarkerSymbol = "51"
	ScatterpolarMarkerSymbolArrowBarLeft            ScatterpolarMarkerSymbol = "arrow-bar-left"
	ScatterpolarMarkerSymbolNumber151               ScatterpolarMarkerSymbol = 151
	ScatterpolarMarkerSymbol151                     ScatterpolarMarkerSymbol = "151"
	ScatterpolarMarkerSymbolArrowBarLeftOpen        ScatterpolarMarkerSymbol = "arrow-bar-left-open"
	ScatterpolarMarkerSymbolNumber52                ScatterpolarMarkerSymbol = 52
	ScatterpolarMarkerSymbol52                      ScatterpolarMarkerSymbol = "52"
	ScatterpolarMarkerSymbolArrowBarRight           ScatterpolarMarkerSymbol = "arrow-bar-right"
	ScatterpolarMarkerSymbolNumber152               ScatterpolarMarkerSymbol = 152
	ScatterpolarMarkerSymbol152                     ScatterpolarMarkerSymbol = "152"
	ScatterpolarMarkerSymbolArrowBarRightOpen       ScatterpolarMarkerSymbol = "arrow-bar-right-open"
)

// ScatterpolarTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterpolarTextposition string

const (
	ScatterpolarTextpositionTopLeft      ScatterpolarTextposition = "top left"
	ScatterpolarTextpositionTopCenter    ScatterpolarTextposition = "top center"
	ScatterpolarTextpositionTopRight     ScatterpolarTextposition = "top right"
	ScatterpolarTextpositionMiddleLeft   ScatterpolarTextposition = "middle left"
	ScatterpolarTextpositionMiddleCenter ScatterpolarTextposition = "middle center"
	ScatterpolarTextpositionMiddleRight  ScatterpolarTextposition = "middle right"
	ScatterpolarTextpositionBottomLeft   ScatterpolarTextposition = "bottom left"
	ScatterpolarTextpositionBottomCenter ScatterpolarTextposition = "bottom center"
	ScatterpolarTextpositionBottomRight  ScatterpolarTextposition = "bottom right"
)

// ScatterpolarThetaunit Sets the unit of input *theta* values. Has an effect only when on *linear* angular axes.
type ScatterpolarThetaunit string

const (
	ScatterpolarThetaunitRadians  ScatterpolarThetaunit = "radians"
	ScatterpolarThetaunitDegrees  ScatterpolarThetaunit = "degrees"
	ScatterpolarThetaunitGradians ScatterpolarThetaunit = "gradians"
)

// ScatterpolarVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterpolarVisible interface{}

var (
	ScatterpolarVisibleTrue       ScatterpolarVisible = true
	ScatterpolarVisibleFalse      ScatterpolarVisible = false
	ScatterpolarVisibleLegendonly ScatterpolarVisible = "legendonly"
)

// ScatterpolarHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type ScatterpolarHoverinfo string

const (
	// Flags
	ScatterpolarHoverinfoR     ScatterpolarHoverinfo = "r"
	ScatterpolarHoverinfoTheta ScatterpolarHoverinfo = "theta"
	ScatterpolarHoverinfoText  ScatterpolarHoverinfo = "text"
	ScatterpolarHoverinfoName  ScatterpolarHoverinfo = "name"

	// Extra
	ScatterpolarHoverinfoAll  ScatterpolarHoverinfo = "all"
	ScatterpolarHoverinfoNone ScatterpolarHoverinfo = "none"
	ScatterpolarHoverinfoSkip ScatterpolarHoverinfo = "skip"
)

// ScatterpolarHoveron Do the hover effects highlight individual points (markers or line points) or do they highlight filled regions? If the fill is *toself* or *tonext* and there are no markers or text, then the default is *fills*, otherwise it is *points*.
type ScatterpolarHoveron string

const (
	// Flags
	ScatterpolarHoveronPoints ScatterpolarHoveron = "points"
	ScatterpolarHoveronFills  ScatterpolarHoveron = "fills"

	// Extra

)

// ScatterpolarMode Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
type ScatterpolarMode string

const (
	// Flags
	ScatterpolarModeLines   ScatterpolarMode = "lines"
	ScatterpolarModeMarkers ScatterpolarMode = "markers"
	ScatterpolarModeText    ScatterpolarMode = "text"

	// Extra
	ScatterpolarModeNone ScatterpolarMode = "none"
)
