package grob

var TraceTypeScatterpolargl TraceType = "scatterpolargl"

func (trace *Scatterpolargl) GetType() TraceType {
	return TraceTypeScatterpolargl
}

// Scatterpolargl The scatterpolargl trace type encompasses line charts, scatter charts, and bubble charts in polar coordinates using the WebGL plotting engine. The data visualized as scatter point or lines is set in `r` (radial) and `theta` (angular) coordinates Bubble charts are achieved by setting `marker.size` and/or `marker.color` to numerical arrays.
type Scatterpolargl struct {

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
	// Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
	Fill ScatterpolarglFill `json:"fill,omitempty"`

	// Fillcolor
	// arrayOK: false
	// type: color
	// Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor Color `json:"fillcolor,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ScatterpolarglHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *ScatterpolarglHoverlabel `json:"hoverlabel,omitempty"`

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
	Line *ScatterpolarglLine `json:"line,omitempty"`

	// Marker
	// role: Object
	Marker *ScatterpolarglMarker `json:"marker,omitempty"`

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
	Mode ScatterpolarglMode `json:"mode,omitempty"`

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
	Selected *ScatterpolarglSelected `json:"selected,omitempty"`

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
	Stream *ScatterpolarglStream `json:"stream,omitempty"`

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
	Textfont *ScatterpolarglTextfont `json:"textfont,omitempty"`

	// Textposition
	// default: middle center
	// type: enumerated
	// Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition ScatterpolarglTextposition `json:"textposition,omitempty"`

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
	Thetaunit ScatterpolarglThetaunit `json:"thetaunit,omitempty"`

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
	Unselected *ScatterpolarglUnselected `json:"unselected,omitempty"`

	// Visible
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ScatterpolarglVisible `json:"visible,omitempty"`
}

// ScatterpolarglHoverlabelFont Sets the font used in hover labels.
type ScatterpolarglHoverlabelFont struct {

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

// ScatterpolarglHoverlabel
type ScatterpolarglHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align ScatterpolarglHoverlabelAlign `json:"align,omitempty"`

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
	Font *ScatterpolarglHoverlabelFont `json:"font,omitempty"`

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

// ScatterpolarglLine
type ScatterpolarglLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the line color.
	Color Color `json:"color,omitempty"`

	// Dash
	// default: solid
	// type: enumerated
	// Sets the style of the lines.
	Dash ScatterpolarglLineDash `json:"dash,omitempty"`

	// Shape
	// default: linear
	// type: enumerated
	// Determines the line shape. The values correspond to step-wise line shapes.
	Shape ScatterpolarglLineShape `json:"shape,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the line width (in px).
	Width float64 `json:"width,omitempty"`
}

// ScatterpolarglMarkerColorbarTickfont Sets the color bar's tick label font
type ScatterpolarglMarkerColorbarTickfont struct {

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

// ScatterpolarglMarkerColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type ScatterpolarglMarkerColorbarTitleFont struct {

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

// ScatterpolarglMarkerColorbarTitle
type ScatterpolarglMarkerColorbarTitle struct {

	// Font
	// role: Object
	Font *ScatterpolarglMarkerColorbarTitleFont `json:"font,omitempty"`

	// Side
	// default: top
	// type: enumerated
	// Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
	Side ScatterpolarglMarkerColorbarTitleSide `json:"side,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// ScatterpolarglMarkerColorbar
type ScatterpolarglMarkerColorbar struct {

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
	Exponentformat ScatterpolarglMarkerColorbarExponentformat `json:"exponentformat,omitempty"`

	// Len
	// arrayOK: false
	// type: number
	// Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends.
	Len float64 `json:"len,omitempty"`

	// Lenmode
	// default: fraction
	// type: enumerated
	// Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
	Lenmode ScatterpolarglMarkerColorbarLenmode `json:"lenmode,omitempty"`

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
	Showexponent ScatterpolarglMarkerColorbarShowexponent `json:"showexponent,omitempty"`

	// Showticklabels
	// arrayOK: false
	// type: boolean
	// Determines whether or not the tick labels are drawn.
	Showticklabels Bool `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix ScatterpolarglMarkerColorbarShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix ScatterpolarglMarkerColorbarShowticksuffix `json:"showticksuffix,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels.
	Thickness float64 `json:"thickness,omitempty"`

	// Thicknessmode
	// default: pixels
	// type: enumerated
	// Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
	Thicknessmode ScatterpolarglMarkerColorbarThicknessmode `json:"thicknessmode,omitempty"`

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
	Tickfont *ScatterpolarglMarkerColorbarTickfont `json:"tickfont,omitempty"`

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
	Ticklabelposition ScatterpolarglMarkerColorbarTicklabelposition `json:"ticklabelposition,omitempty"`

	// Ticklen
	// arrayOK: false
	// type: number
	// Sets the tick length (in px).
	Ticklen float64 `json:"ticklen,omitempty"`

	// Tickmode
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
	Tickmode ScatterpolarglMarkerColorbarTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

	// Ticks
	// default:
	// type: enumerated
	// Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
	Ticks ScatterpolarglMarkerColorbarTicks `json:"ticks,omitempty"`

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
	Title *ScatterpolarglMarkerColorbarTitle `json:"title,omitempty"`

	// X
	// arrayOK: false
	// type: number
	// Sets the x position of the color bar (in plot fraction).
	X float64 `json:"x,omitempty"`

	// Xanchor
	// default: left
	// type: enumerated
	// Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
	Xanchor ScatterpolarglMarkerColorbarXanchor `json:"xanchor,omitempty"`

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
	Yanchor ScatterpolarglMarkerColorbarYanchor `json:"yanchor,omitempty"`

	// Ypad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the y direction.
	Ypad float64 `json:"ypad,omitempty"`
}

// ScatterpolarglMarkerLine
type ScatterpolarglMarkerLine struct {

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

// ScatterpolarglMarker
type ScatterpolarglMarker struct {

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
	Colorbar *ScatterpolarglMarkerColorbar `json:"colorbar,omitempty"`

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
	Line *ScatterpolarglMarkerLine `json:"line,omitempty"`

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
	Sizemode ScatterpolarglMarkerSizemode `json:"sizemode,omitempty"`

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
	Symbol ScatterpolarglMarkerSymbol `json:"symbol,omitempty"`

	// Symbolsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  symbol .
	Symbolsrc String `json:"symbolsrc,omitempty"`
}

// ScatterpolarglSelectedMarker
type ScatterpolarglSelectedMarker struct {

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

// ScatterpolarglSelectedTextfont
type ScatterpolarglSelectedTextfont struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the text font color of selected points.
	Color Color `json:"color,omitempty"`
}

// ScatterpolarglSelected
type ScatterpolarglSelected struct {

	// Marker
	// role: Object
	Marker *ScatterpolarglSelectedMarker `json:"marker,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScatterpolarglSelectedTextfont `json:"textfont,omitempty"`
}

// ScatterpolarglStream
type ScatterpolarglStream struct {

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

// ScatterpolarglTextfont Sets the text font.
type ScatterpolarglTextfont struct {

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

// ScatterpolarglUnselectedMarker
type ScatterpolarglUnselectedMarker struct {

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

// ScatterpolarglUnselectedTextfont
type ScatterpolarglUnselectedTextfont struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the text font color of unselected points, applied only when a selection exists.
	Color Color `json:"color,omitempty"`
}

// ScatterpolarglUnselected
type ScatterpolarglUnselected struct {

	// Marker
	// role: Object
	Marker *ScatterpolarglUnselectedMarker `json:"marker,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScatterpolarglUnselectedTextfont `json:"textfont,omitempty"`
}

// ScatterpolarglFill Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
type ScatterpolarglFill string

const (
	ScatterpolarglFillNone    ScatterpolarglFill = "none"
	ScatterpolarglFillTozeroy ScatterpolarglFill = "tozeroy"
	ScatterpolarglFillTozerox ScatterpolarglFill = "tozerox"
	ScatterpolarglFillTonexty ScatterpolarglFill = "tonexty"
	ScatterpolarglFillTonextx ScatterpolarglFill = "tonextx"
	ScatterpolarglFillToself  ScatterpolarglFill = "toself"
	ScatterpolarglFillTonext  ScatterpolarglFill = "tonext"
)

// ScatterpolarglHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterpolarglHoverlabelAlign string

const (
	ScatterpolarglHoverlabelAlignLeft  ScatterpolarglHoverlabelAlign = "left"
	ScatterpolarglHoverlabelAlignRight ScatterpolarglHoverlabelAlign = "right"
	ScatterpolarglHoverlabelAlignAuto  ScatterpolarglHoverlabelAlign = "auto"
)

// ScatterpolarglLineDash Sets the style of the lines.
type ScatterpolarglLineDash string

const (
	ScatterpolarglLineDashSolid       ScatterpolarglLineDash = "solid"
	ScatterpolarglLineDashDot         ScatterpolarglLineDash = "dot"
	ScatterpolarglLineDashDash        ScatterpolarglLineDash = "dash"
	ScatterpolarglLineDashLongdash    ScatterpolarglLineDash = "longdash"
	ScatterpolarglLineDashDashdot     ScatterpolarglLineDash = "dashdot"
	ScatterpolarglLineDashLongdashdot ScatterpolarglLineDash = "longdashdot"
)

// ScatterpolarglLineShape Determines the line shape. The values correspond to step-wise line shapes.
type ScatterpolarglLineShape string

const (
	ScatterpolarglLineShapeLinear ScatterpolarglLineShape = "linear"
	ScatterpolarglLineShapeHv     ScatterpolarglLineShape = "hv"
	ScatterpolarglLineShapeVh     ScatterpolarglLineShape = "vh"
	ScatterpolarglLineShapeHvh    ScatterpolarglLineShape = "hvh"
	ScatterpolarglLineShapeVhv    ScatterpolarglLineShape = "vhv"
)

// ScatterpolarglMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterpolarglMarkerColorbarExponentformat string

const (
	ScatterpolarglMarkerColorbarExponentformatNone  ScatterpolarglMarkerColorbarExponentformat = "none"
	ScatterpolarglMarkerColorbarExponentformatE1    ScatterpolarglMarkerColorbarExponentformat = "e"
	ScatterpolarglMarkerColorbarExponentformatE2    ScatterpolarglMarkerColorbarExponentformat = "E"
	ScatterpolarglMarkerColorbarExponentformatPower ScatterpolarglMarkerColorbarExponentformat = "power"
	ScatterpolarglMarkerColorbarExponentformatSi    ScatterpolarglMarkerColorbarExponentformat = "SI"
	ScatterpolarglMarkerColorbarExponentformatB     ScatterpolarglMarkerColorbarExponentformat = "B"
)

// ScatterpolarglMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterpolarglMarkerColorbarLenmode string

const (
	ScatterpolarglMarkerColorbarLenmodeFraction ScatterpolarglMarkerColorbarLenmode = "fraction"
	ScatterpolarglMarkerColorbarLenmodePixels   ScatterpolarglMarkerColorbarLenmode = "pixels"
)

// ScatterpolarglMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterpolarglMarkerColorbarShowexponent string

const (
	ScatterpolarglMarkerColorbarShowexponentAll   ScatterpolarglMarkerColorbarShowexponent = "all"
	ScatterpolarglMarkerColorbarShowexponentFirst ScatterpolarglMarkerColorbarShowexponent = "first"
	ScatterpolarglMarkerColorbarShowexponentLast  ScatterpolarglMarkerColorbarShowexponent = "last"
	ScatterpolarglMarkerColorbarShowexponentNone  ScatterpolarglMarkerColorbarShowexponent = "none"
)

// ScatterpolarglMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterpolarglMarkerColorbarShowtickprefix string

const (
	ScatterpolarglMarkerColorbarShowtickprefixAll   ScatterpolarglMarkerColorbarShowtickprefix = "all"
	ScatterpolarglMarkerColorbarShowtickprefixFirst ScatterpolarglMarkerColorbarShowtickprefix = "first"
	ScatterpolarglMarkerColorbarShowtickprefixLast  ScatterpolarglMarkerColorbarShowtickprefix = "last"
	ScatterpolarglMarkerColorbarShowtickprefixNone  ScatterpolarglMarkerColorbarShowtickprefix = "none"
)

// ScatterpolarglMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterpolarglMarkerColorbarShowticksuffix string

const (
	ScatterpolarglMarkerColorbarShowticksuffixAll   ScatterpolarglMarkerColorbarShowticksuffix = "all"
	ScatterpolarglMarkerColorbarShowticksuffixFirst ScatterpolarglMarkerColorbarShowticksuffix = "first"
	ScatterpolarglMarkerColorbarShowticksuffixLast  ScatterpolarglMarkerColorbarShowticksuffix = "last"
	ScatterpolarglMarkerColorbarShowticksuffixNone  ScatterpolarglMarkerColorbarShowticksuffix = "none"
)

// ScatterpolarglMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterpolarglMarkerColorbarThicknessmode string

const (
	ScatterpolarglMarkerColorbarThicknessmodeFraction ScatterpolarglMarkerColorbarThicknessmode = "fraction"
	ScatterpolarglMarkerColorbarThicknessmodePixels   ScatterpolarglMarkerColorbarThicknessmode = "pixels"
)

// ScatterpolarglMarkerColorbarTicklabelposition Determines where tick labels are drawn.
type ScatterpolarglMarkerColorbarTicklabelposition string

const (
	ScatterpolarglMarkerColorbarTicklabelpositionOutside       ScatterpolarglMarkerColorbarTicklabelposition = "outside"
	ScatterpolarglMarkerColorbarTicklabelpositionInside        ScatterpolarglMarkerColorbarTicklabelposition = "inside"
	ScatterpolarglMarkerColorbarTicklabelpositionOutsideTop    ScatterpolarglMarkerColorbarTicklabelposition = "outside top"
	ScatterpolarglMarkerColorbarTicklabelpositionInsideTop     ScatterpolarglMarkerColorbarTicklabelposition = "inside top"
	ScatterpolarglMarkerColorbarTicklabelpositionOutsideBottom ScatterpolarglMarkerColorbarTicklabelposition = "outside bottom"
	ScatterpolarglMarkerColorbarTicklabelpositionInsideBottom  ScatterpolarglMarkerColorbarTicklabelposition = "inside bottom"
)

// ScatterpolarglMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterpolarglMarkerColorbarTickmode string

const (
	ScatterpolarglMarkerColorbarTickmodeAuto   ScatterpolarglMarkerColorbarTickmode = "auto"
	ScatterpolarglMarkerColorbarTickmodeLinear ScatterpolarglMarkerColorbarTickmode = "linear"
	ScatterpolarglMarkerColorbarTickmodeArray  ScatterpolarglMarkerColorbarTickmode = "array"
)

// ScatterpolarglMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterpolarglMarkerColorbarTicks string

const (
	ScatterpolarglMarkerColorbarTicksOutside ScatterpolarglMarkerColorbarTicks = "outside"
	ScatterpolarglMarkerColorbarTicksInside  ScatterpolarglMarkerColorbarTicks = "inside"
	ScatterpolarglMarkerColorbarTicksEmpty   ScatterpolarglMarkerColorbarTicks = ""
)

// ScatterpolarglMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterpolarglMarkerColorbarTitleSide string

const (
	ScatterpolarglMarkerColorbarTitleSideRight  ScatterpolarglMarkerColorbarTitleSide = "right"
	ScatterpolarglMarkerColorbarTitleSideTop    ScatterpolarglMarkerColorbarTitleSide = "top"
	ScatterpolarglMarkerColorbarTitleSideBottom ScatterpolarglMarkerColorbarTitleSide = "bottom"
)

// ScatterpolarglMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterpolarglMarkerColorbarXanchor string

const (
	ScatterpolarglMarkerColorbarXanchorLeft   ScatterpolarglMarkerColorbarXanchor = "left"
	ScatterpolarglMarkerColorbarXanchorCenter ScatterpolarglMarkerColorbarXanchor = "center"
	ScatterpolarglMarkerColorbarXanchorRight  ScatterpolarglMarkerColorbarXanchor = "right"
)

// ScatterpolarglMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterpolarglMarkerColorbarYanchor string

const (
	ScatterpolarglMarkerColorbarYanchorTop    ScatterpolarglMarkerColorbarYanchor = "top"
	ScatterpolarglMarkerColorbarYanchorMiddle ScatterpolarglMarkerColorbarYanchor = "middle"
	ScatterpolarglMarkerColorbarYanchorBottom ScatterpolarglMarkerColorbarYanchor = "bottom"
)

// ScatterpolarglMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterpolarglMarkerSizemode string

const (
	ScatterpolarglMarkerSizemodeDiameter ScatterpolarglMarkerSizemode = "diameter"
	ScatterpolarglMarkerSizemodeArea     ScatterpolarglMarkerSizemode = "area"
)

// ScatterpolarglMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterpolarglMarkerSymbol interface{}

var (
	ScatterpolarglMarkerSymbolNumber0                 ScatterpolarglMarkerSymbol = 0
	ScatterpolarglMarkerSymbol0                       ScatterpolarglMarkerSymbol = "0"
	ScatterpolarglMarkerSymbolCircle                  ScatterpolarglMarkerSymbol = "circle"
	ScatterpolarglMarkerSymbolNumber100               ScatterpolarglMarkerSymbol = 100
	ScatterpolarglMarkerSymbol100                     ScatterpolarglMarkerSymbol = "100"
	ScatterpolarglMarkerSymbolCircleOpen              ScatterpolarglMarkerSymbol = "circle-open"
	ScatterpolarglMarkerSymbolNumber200               ScatterpolarglMarkerSymbol = 200
	ScatterpolarglMarkerSymbol200                     ScatterpolarglMarkerSymbol = "200"
	ScatterpolarglMarkerSymbolCircleDot               ScatterpolarglMarkerSymbol = "circle-dot"
	ScatterpolarglMarkerSymbolNumber300               ScatterpolarglMarkerSymbol = 300
	ScatterpolarglMarkerSymbol300                     ScatterpolarglMarkerSymbol = "300"
	ScatterpolarglMarkerSymbolCircleOpenDot           ScatterpolarglMarkerSymbol = "circle-open-dot"
	ScatterpolarglMarkerSymbolNumber1                 ScatterpolarglMarkerSymbol = 1
	ScatterpolarglMarkerSymbol1                       ScatterpolarglMarkerSymbol = "1"
	ScatterpolarglMarkerSymbolSquare                  ScatterpolarglMarkerSymbol = "square"
	ScatterpolarglMarkerSymbolNumber101               ScatterpolarglMarkerSymbol = 101
	ScatterpolarglMarkerSymbol101                     ScatterpolarglMarkerSymbol = "101"
	ScatterpolarglMarkerSymbolSquareOpen              ScatterpolarglMarkerSymbol = "square-open"
	ScatterpolarglMarkerSymbolNumber201               ScatterpolarglMarkerSymbol = 201
	ScatterpolarglMarkerSymbol201                     ScatterpolarglMarkerSymbol = "201"
	ScatterpolarglMarkerSymbolSquareDot               ScatterpolarglMarkerSymbol = "square-dot"
	ScatterpolarglMarkerSymbolNumber301               ScatterpolarglMarkerSymbol = 301
	ScatterpolarglMarkerSymbol301                     ScatterpolarglMarkerSymbol = "301"
	ScatterpolarglMarkerSymbolSquareOpenDot           ScatterpolarglMarkerSymbol = "square-open-dot"
	ScatterpolarglMarkerSymbolNumber2                 ScatterpolarglMarkerSymbol = 2
	ScatterpolarglMarkerSymbol2                       ScatterpolarglMarkerSymbol = "2"
	ScatterpolarglMarkerSymbolDiamond                 ScatterpolarglMarkerSymbol = "diamond"
	ScatterpolarglMarkerSymbolNumber102               ScatterpolarglMarkerSymbol = 102
	ScatterpolarglMarkerSymbol102                     ScatterpolarglMarkerSymbol = "102"
	ScatterpolarglMarkerSymbolDiamondOpen             ScatterpolarglMarkerSymbol = "diamond-open"
	ScatterpolarglMarkerSymbolNumber202               ScatterpolarglMarkerSymbol = 202
	ScatterpolarglMarkerSymbol202                     ScatterpolarglMarkerSymbol = "202"
	ScatterpolarglMarkerSymbolDiamondDot              ScatterpolarglMarkerSymbol = "diamond-dot"
	ScatterpolarglMarkerSymbolNumber302               ScatterpolarglMarkerSymbol = 302
	ScatterpolarglMarkerSymbol302                     ScatterpolarglMarkerSymbol = "302"
	ScatterpolarglMarkerSymbolDiamondOpenDot          ScatterpolarglMarkerSymbol = "diamond-open-dot"
	ScatterpolarglMarkerSymbolNumber3                 ScatterpolarglMarkerSymbol = 3
	ScatterpolarglMarkerSymbol3                       ScatterpolarglMarkerSymbol = "3"
	ScatterpolarglMarkerSymbolCross                   ScatterpolarglMarkerSymbol = "cross"
	ScatterpolarglMarkerSymbolNumber103               ScatterpolarglMarkerSymbol = 103
	ScatterpolarglMarkerSymbol103                     ScatterpolarglMarkerSymbol = "103"
	ScatterpolarglMarkerSymbolCrossOpen               ScatterpolarglMarkerSymbol = "cross-open"
	ScatterpolarglMarkerSymbolNumber203               ScatterpolarglMarkerSymbol = 203
	ScatterpolarglMarkerSymbol203                     ScatterpolarglMarkerSymbol = "203"
	ScatterpolarglMarkerSymbolCrossDot                ScatterpolarglMarkerSymbol = "cross-dot"
	ScatterpolarglMarkerSymbolNumber303               ScatterpolarglMarkerSymbol = 303
	ScatterpolarglMarkerSymbol303                     ScatterpolarglMarkerSymbol = "303"
	ScatterpolarglMarkerSymbolCrossOpenDot            ScatterpolarglMarkerSymbol = "cross-open-dot"
	ScatterpolarglMarkerSymbolNumber4                 ScatterpolarglMarkerSymbol = 4
	ScatterpolarglMarkerSymbol4                       ScatterpolarglMarkerSymbol = "4"
	ScatterpolarglMarkerSymbolX                       ScatterpolarglMarkerSymbol = "x"
	ScatterpolarglMarkerSymbolNumber104               ScatterpolarglMarkerSymbol = 104
	ScatterpolarglMarkerSymbol104                     ScatterpolarglMarkerSymbol = "104"
	ScatterpolarglMarkerSymbolXOpen                   ScatterpolarglMarkerSymbol = "x-open"
	ScatterpolarglMarkerSymbolNumber204               ScatterpolarglMarkerSymbol = 204
	ScatterpolarglMarkerSymbol204                     ScatterpolarglMarkerSymbol = "204"
	ScatterpolarglMarkerSymbolXDot                    ScatterpolarglMarkerSymbol = "x-dot"
	ScatterpolarglMarkerSymbolNumber304               ScatterpolarglMarkerSymbol = 304
	ScatterpolarglMarkerSymbol304                     ScatterpolarglMarkerSymbol = "304"
	ScatterpolarglMarkerSymbolXOpenDot                ScatterpolarglMarkerSymbol = "x-open-dot"
	ScatterpolarglMarkerSymbolNumber5                 ScatterpolarglMarkerSymbol = 5
	ScatterpolarglMarkerSymbol5                       ScatterpolarglMarkerSymbol = "5"
	ScatterpolarglMarkerSymbolTriangleUp              ScatterpolarglMarkerSymbol = "triangle-up"
	ScatterpolarglMarkerSymbolNumber105               ScatterpolarglMarkerSymbol = 105
	ScatterpolarglMarkerSymbol105                     ScatterpolarglMarkerSymbol = "105"
	ScatterpolarglMarkerSymbolTriangleUpOpen          ScatterpolarglMarkerSymbol = "triangle-up-open"
	ScatterpolarglMarkerSymbolNumber205               ScatterpolarglMarkerSymbol = 205
	ScatterpolarglMarkerSymbol205                     ScatterpolarglMarkerSymbol = "205"
	ScatterpolarglMarkerSymbolTriangleUpDot           ScatterpolarglMarkerSymbol = "triangle-up-dot"
	ScatterpolarglMarkerSymbolNumber305               ScatterpolarglMarkerSymbol = 305
	ScatterpolarglMarkerSymbol305                     ScatterpolarglMarkerSymbol = "305"
	ScatterpolarglMarkerSymbolTriangleUpOpenDot       ScatterpolarglMarkerSymbol = "triangle-up-open-dot"
	ScatterpolarglMarkerSymbolNumber6                 ScatterpolarglMarkerSymbol = 6
	ScatterpolarglMarkerSymbol6                       ScatterpolarglMarkerSymbol = "6"
	ScatterpolarglMarkerSymbolTriangleDown            ScatterpolarglMarkerSymbol = "triangle-down"
	ScatterpolarglMarkerSymbolNumber106               ScatterpolarglMarkerSymbol = 106
	ScatterpolarglMarkerSymbol106                     ScatterpolarglMarkerSymbol = "106"
	ScatterpolarglMarkerSymbolTriangleDownOpen        ScatterpolarglMarkerSymbol = "triangle-down-open"
	ScatterpolarglMarkerSymbolNumber206               ScatterpolarglMarkerSymbol = 206
	ScatterpolarglMarkerSymbol206                     ScatterpolarglMarkerSymbol = "206"
	ScatterpolarglMarkerSymbolTriangleDownDot         ScatterpolarglMarkerSymbol = "triangle-down-dot"
	ScatterpolarglMarkerSymbolNumber306               ScatterpolarglMarkerSymbol = 306
	ScatterpolarglMarkerSymbol306                     ScatterpolarglMarkerSymbol = "306"
	ScatterpolarglMarkerSymbolTriangleDownOpenDot     ScatterpolarglMarkerSymbol = "triangle-down-open-dot"
	ScatterpolarglMarkerSymbolNumber7                 ScatterpolarglMarkerSymbol = 7
	ScatterpolarglMarkerSymbol7                       ScatterpolarglMarkerSymbol = "7"
	ScatterpolarglMarkerSymbolTriangleLeft            ScatterpolarglMarkerSymbol = "triangle-left"
	ScatterpolarglMarkerSymbolNumber107               ScatterpolarglMarkerSymbol = 107
	ScatterpolarglMarkerSymbol107                     ScatterpolarglMarkerSymbol = "107"
	ScatterpolarglMarkerSymbolTriangleLeftOpen        ScatterpolarglMarkerSymbol = "triangle-left-open"
	ScatterpolarglMarkerSymbolNumber207               ScatterpolarglMarkerSymbol = 207
	ScatterpolarglMarkerSymbol207                     ScatterpolarglMarkerSymbol = "207"
	ScatterpolarglMarkerSymbolTriangleLeftDot         ScatterpolarglMarkerSymbol = "triangle-left-dot"
	ScatterpolarglMarkerSymbolNumber307               ScatterpolarglMarkerSymbol = 307
	ScatterpolarglMarkerSymbol307                     ScatterpolarglMarkerSymbol = "307"
	ScatterpolarglMarkerSymbolTriangleLeftOpenDot     ScatterpolarglMarkerSymbol = "triangle-left-open-dot"
	ScatterpolarglMarkerSymbolNumber8                 ScatterpolarglMarkerSymbol = 8
	ScatterpolarglMarkerSymbol8                       ScatterpolarglMarkerSymbol = "8"
	ScatterpolarglMarkerSymbolTriangleRight           ScatterpolarglMarkerSymbol = "triangle-right"
	ScatterpolarglMarkerSymbolNumber108               ScatterpolarglMarkerSymbol = 108
	ScatterpolarglMarkerSymbol108                     ScatterpolarglMarkerSymbol = "108"
	ScatterpolarglMarkerSymbolTriangleRightOpen       ScatterpolarglMarkerSymbol = "triangle-right-open"
	ScatterpolarglMarkerSymbolNumber208               ScatterpolarglMarkerSymbol = 208
	ScatterpolarglMarkerSymbol208                     ScatterpolarglMarkerSymbol = "208"
	ScatterpolarglMarkerSymbolTriangleRightDot        ScatterpolarglMarkerSymbol = "triangle-right-dot"
	ScatterpolarglMarkerSymbolNumber308               ScatterpolarglMarkerSymbol = 308
	ScatterpolarglMarkerSymbol308                     ScatterpolarglMarkerSymbol = "308"
	ScatterpolarglMarkerSymbolTriangleRightOpenDot    ScatterpolarglMarkerSymbol = "triangle-right-open-dot"
	ScatterpolarglMarkerSymbolNumber9                 ScatterpolarglMarkerSymbol = 9
	ScatterpolarglMarkerSymbol9                       ScatterpolarglMarkerSymbol = "9"
	ScatterpolarglMarkerSymbolTriangleNe              ScatterpolarglMarkerSymbol = "triangle-ne"
	ScatterpolarglMarkerSymbolNumber109               ScatterpolarglMarkerSymbol = 109
	ScatterpolarglMarkerSymbol109                     ScatterpolarglMarkerSymbol = "109"
	ScatterpolarglMarkerSymbolTriangleNeOpen          ScatterpolarglMarkerSymbol = "triangle-ne-open"
	ScatterpolarglMarkerSymbolNumber209               ScatterpolarglMarkerSymbol = 209
	ScatterpolarglMarkerSymbol209                     ScatterpolarglMarkerSymbol = "209"
	ScatterpolarglMarkerSymbolTriangleNeDot           ScatterpolarglMarkerSymbol = "triangle-ne-dot"
	ScatterpolarglMarkerSymbolNumber309               ScatterpolarglMarkerSymbol = 309
	ScatterpolarglMarkerSymbol309                     ScatterpolarglMarkerSymbol = "309"
	ScatterpolarglMarkerSymbolTriangleNeOpenDot       ScatterpolarglMarkerSymbol = "triangle-ne-open-dot"
	ScatterpolarglMarkerSymbolNumber10                ScatterpolarglMarkerSymbol = 10
	ScatterpolarglMarkerSymbol10                      ScatterpolarglMarkerSymbol = "10"
	ScatterpolarglMarkerSymbolTriangleSe              ScatterpolarglMarkerSymbol = "triangle-se"
	ScatterpolarglMarkerSymbolNumber110               ScatterpolarglMarkerSymbol = 110
	ScatterpolarglMarkerSymbol110                     ScatterpolarglMarkerSymbol = "110"
	ScatterpolarglMarkerSymbolTriangleSeOpen          ScatterpolarglMarkerSymbol = "triangle-se-open"
	ScatterpolarglMarkerSymbolNumber210               ScatterpolarglMarkerSymbol = 210
	ScatterpolarglMarkerSymbol210                     ScatterpolarglMarkerSymbol = "210"
	ScatterpolarglMarkerSymbolTriangleSeDot           ScatterpolarglMarkerSymbol = "triangle-se-dot"
	ScatterpolarglMarkerSymbolNumber310               ScatterpolarglMarkerSymbol = 310
	ScatterpolarglMarkerSymbol310                     ScatterpolarglMarkerSymbol = "310"
	ScatterpolarglMarkerSymbolTriangleSeOpenDot       ScatterpolarglMarkerSymbol = "triangle-se-open-dot"
	ScatterpolarglMarkerSymbolNumber11                ScatterpolarglMarkerSymbol = 11
	ScatterpolarglMarkerSymbol11                      ScatterpolarglMarkerSymbol = "11"
	ScatterpolarglMarkerSymbolTriangleSw              ScatterpolarglMarkerSymbol = "triangle-sw"
	ScatterpolarglMarkerSymbolNumber111               ScatterpolarglMarkerSymbol = 111
	ScatterpolarglMarkerSymbol111                     ScatterpolarglMarkerSymbol = "111"
	ScatterpolarglMarkerSymbolTriangleSwOpen          ScatterpolarglMarkerSymbol = "triangle-sw-open"
	ScatterpolarglMarkerSymbolNumber211               ScatterpolarglMarkerSymbol = 211
	ScatterpolarglMarkerSymbol211                     ScatterpolarglMarkerSymbol = "211"
	ScatterpolarglMarkerSymbolTriangleSwDot           ScatterpolarglMarkerSymbol = "triangle-sw-dot"
	ScatterpolarglMarkerSymbolNumber311               ScatterpolarglMarkerSymbol = 311
	ScatterpolarglMarkerSymbol311                     ScatterpolarglMarkerSymbol = "311"
	ScatterpolarglMarkerSymbolTriangleSwOpenDot       ScatterpolarglMarkerSymbol = "triangle-sw-open-dot"
	ScatterpolarglMarkerSymbolNumber12                ScatterpolarglMarkerSymbol = 12
	ScatterpolarglMarkerSymbol12                      ScatterpolarglMarkerSymbol = "12"
	ScatterpolarglMarkerSymbolTriangleNw              ScatterpolarglMarkerSymbol = "triangle-nw"
	ScatterpolarglMarkerSymbolNumber112               ScatterpolarglMarkerSymbol = 112
	ScatterpolarglMarkerSymbol112                     ScatterpolarglMarkerSymbol = "112"
	ScatterpolarglMarkerSymbolTriangleNwOpen          ScatterpolarglMarkerSymbol = "triangle-nw-open"
	ScatterpolarglMarkerSymbolNumber212               ScatterpolarglMarkerSymbol = 212
	ScatterpolarglMarkerSymbol212                     ScatterpolarglMarkerSymbol = "212"
	ScatterpolarglMarkerSymbolTriangleNwDot           ScatterpolarglMarkerSymbol = "triangle-nw-dot"
	ScatterpolarglMarkerSymbolNumber312               ScatterpolarglMarkerSymbol = 312
	ScatterpolarglMarkerSymbol312                     ScatterpolarglMarkerSymbol = "312"
	ScatterpolarglMarkerSymbolTriangleNwOpenDot       ScatterpolarglMarkerSymbol = "triangle-nw-open-dot"
	ScatterpolarglMarkerSymbolNumber13                ScatterpolarglMarkerSymbol = 13
	ScatterpolarglMarkerSymbol13                      ScatterpolarglMarkerSymbol = "13"
	ScatterpolarglMarkerSymbolPentagon                ScatterpolarglMarkerSymbol = "pentagon"
	ScatterpolarglMarkerSymbolNumber113               ScatterpolarglMarkerSymbol = 113
	ScatterpolarglMarkerSymbol113                     ScatterpolarglMarkerSymbol = "113"
	ScatterpolarglMarkerSymbolPentagonOpen            ScatterpolarglMarkerSymbol = "pentagon-open"
	ScatterpolarglMarkerSymbolNumber213               ScatterpolarglMarkerSymbol = 213
	ScatterpolarglMarkerSymbol213                     ScatterpolarglMarkerSymbol = "213"
	ScatterpolarglMarkerSymbolPentagonDot             ScatterpolarglMarkerSymbol = "pentagon-dot"
	ScatterpolarglMarkerSymbolNumber313               ScatterpolarglMarkerSymbol = 313
	ScatterpolarglMarkerSymbol313                     ScatterpolarglMarkerSymbol = "313"
	ScatterpolarglMarkerSymbolPentagonOpenDot         ScatterpolarglMarkerSymbol = "pentagon-open-dot"
	ScatterpolarglMarkerSymbolNumber14                ScatterpolarglMarkerSymbol = 14
	ScatterpolarglMarkerSymbol14                      ScatterpolarglMarkerSymbol = "14"
	ScatterpolarglMarkerSymbolHexagon                 ScatterpolarglMarkerSymbol = "hexagon"
	ScatterpolarglMarkerSymbolNumber114               ScatterpolarglMarkerSymbol = 114
	ScatterpolarglMarkerSymbol114                     ScatterpolarglMarkerSymbol = "114"
	ScatterpolarglMarkerSymbolHexagonOpen             ScatterpolarglMarkerSymbol = "hexagon-open"
	ScatterpolarglMarkerSymbolNumber214               ScatterpolarglMarkerSymbol = 214
	ScatterpolarglMarkerSymbol214                     ScatterpolarglMarkerSymbol = "214"
	ScatterpolarglMarkerSymbolHexagonDot              ScatterpolarglMarkerSymbol = "hexagon-dot"
	ScatterpolarglMarkerSymbolNumber314               ScatterpolarglMarkerSymbol = 314
	ScatterpolarglMarkerSymbol314                     ScatterpolarglMarkerSymbol = "314"
	ScatterpolarglMarkerSymbolHexagonOpenDot          ScatterpolarglMarkerSymbol = "hexagon-open-dot"
	ScatterpolarglMarkerSymbolNumber15                ScatterpolarglMarkerSymbol = 15
	ScatterpolarglMarkerSymbol15                      ScatterpolarglMarkerSymbol = "15"
	ScatterpolarglMarkerSymbolHexagon2                ScatterpolarglMarkerSymbol = "hexagon2"
	ScatterpolarglMarkerSymbolNumber115               ScatterpolarglMarkerSymbol = 115
	ScatterpolarglMarkerSymbol115                     ScatterpolarglMarkerSymbol = "115"
	ScatterpolarglMarkerSymbolHexagon2Open            ScatterpolarglMarkerSymbol = "hexagon2-open"
	ScatterpolarglMarkerSymbolNumber215               ScatterpolarglMarkerSymbol = 215
	ScatterpolarglMarkerSymbol215                     ScatterpolarglMarkerSymbol = "215"
	ScatterpolarglMarkerSymbolHexagon2Dot             ScatterpolarglMarkerSymbol = "hexagon2-dot"
	ScatterpolarglMarkerSymbolNumber315               ScatterpolarglMarkerSymbol = 315
	ScatterpolarglMarkerSymbol315                     ScatterpolarglMarkerSymbol = "315"
	ScatterpolarglMarkerSymbolHexagon2OpenDot         ScatterpolarglMarkerSymbol = "hexagon2-open-dot"
	ScatterpolarglMarkerSymbolNumber16                ScatterpolarglMarkerSymbol = 16
	ScatterpolarglMarkerSymbol16                      ScatterpolarglMarkerSymbol = "16"
	ScatterpolarglMarkerSymbolOctagon                 ScatterpolarglMarkerSymbol = "octagon"
	ScatterpolarglMarkerSymbolNumber116               ScatterpolarglMarkerSymbol = 116
	ScatterpolarglMarkerSymbol116                     ScatterpolarglMarkerSymbol = "116"
	ScatterpolarglMarkerSymbolOctagonOpen             ScatterpolarglMarkerSymbol = "octagon-open"
	ScatterpolarglMarkerSymbolNumber216               ScatterpolarglMarkerSymbol = 216
	ScatterpolarglMarkerSymbol216                     ScatterpolarglMarkerSymbol = "216"
	ScatterpolarglMarkerSymbolOctagonDot              ScatterpolarglMarkerSymbol = "octagon-dot"
	ScatterpolarglMarkerSymbolNumber316               ScatterpolarglMarkerSymbol = 316
	ScatterpolarglMarkerSymbol316                     ScatterpolarglMarkerSymbol = "316"
	ScatterpolarglMarkerSymbolOctagonOpenDot          ScatterpolarglMarkerSymbol = "octagon-open-dot"
	ScatterpolarglMarkerSymbolNumber17                ScatterpolarglMarkerSymbol = 17
	ScatterpolarglMarkerSymbol17                      ScatterpolarglMarkerSymbol = "17"
	ScatterpolarglMarkerSymbolStar                    ScatterpolarglMarkerSymbol = "star"
	ScatterpolarglMarkerSymbolNumber117               ScatterpolarglMarkerSymbol = 117
	ScatterpolarglMarkerSymbol117                     ScatterpolarglMarkerSymbol = "117"
	ScatterpolarglMarkerSymbolStarOpen                ScatterpolarglMarkerSymbol = "star-open"
	ScatterpolarglMarkerSymbolNumber217               ScatterpolarglMarkerSymbol = 217
	ScatterpolarglMarkerSymbol217                     ScatterpolarglMarkerSymbol = "217"
	ScatterpolarglMarkerSymbolStarDot                 ScatterpolarglMarkerSymbol = "star-dot"
	ScatterpolarglMarkerSymbolNumber317               ScatterpolarglMarkerSymbol = 317
	ScatterpolarglMarkerSymbol317                     ScatterpolarglMarkerSymbol = "317"
	ScatterpolarglMarkerSymbolStarOpenDot             ScatterpolarglMarkerSymbol = "star-open-dot"
	ScatterpolarglMarkerSymbolNumber18                ScatterpolarglMarkerSymbol = 18
	ScatterpolarglMarkerSymbol18                      ScatterpolarglMarkerSymbol = "18"
	ScatterpolarglMarkerSymbolHexagram                ScatterpolarglMarkerSymbol = "hexagram"
	ScatterpolarglMarkerSymbolNumber118               ScatterpolarglMarkerSymbol = 118
	ScatterpolarglMarkerSymbol118                     ScatterpolarglMarkerSymbol = "118"
	ScatterpolarglMarkerSymbolHexagramOpen            ScatterpolarglMarkerSymbol = "hexagram-open"
	ScatterpolarglMarkerSymbolNumber218               ScatterpolarglMarkerSymbol = 218
	ScatterpolarglMarkerSymbol218                     ScatterpolarglMarkerSymbol = "218"
	ScatterpolarglMarkerSymbolHexagramDot             ScatterpolarglMarkerSymbol = "hexagram-dot"
	ScatterpolarglMarkerSymbolNumber318               ScatterpolarglMarkerSymbol = 318
	ScatterpolarglMarkerSymbol318                     ScatterpolarglMarkerSymbol = "318"
	ScatterpolarglMarkerSymbolHexagramOpenDot         ScatterpolarglMarkerSymbol = "hexagram-open-dot"
	ScatterpolarglMarkerSymbolNumber19                ScatterpolarglMarkerSymbol = 19
	ScatterpolarglMarkerSymbol19                      ScatterpolarglMarkerSymbol = "19"
	ScatterpolarglMarkerSymbolStarTriangleUp          ScatterpolarglMarkerSymbol = "star-triangle-up"
	ScatterpolarglMarkerSymbolNumber119               ScatterpolarglMarkerSymbol = 119
	ScatterpolarglMarkerSymbol119                     ScatterpolarglMarkerSymbol = "119"
	ScatterpolarglMarkerSymbolStarTriangleUpOpen      ScatterpolarglMarkerSymbol = "star-triangle-up-open"
	ScatterpolarglMarkerSymbolNumber219               ScatterpolarglMarkerSymbol = 219
	ScatterpolarglMarkerSymbol219                     ScatterpolarglMarkerSymbol = "219"
	ScatterpolarglMarkerSymbolStarTriangleUpDot       ScatterpolarglMarkerSymbol = "star-triangle-up-dot"
	ScatterpolarglMarkerSymbolNumber319               ScatterpolarglMarkerSymbol = 319
	ScatterpolarglMarkerSymbol319                     ScatterpolarglMarkerSymbol = "319"
	ScatterpolarglMarkerSymbolStarTriangleUpOpenDot   ScatterpolarglMarkerSymbol = "star-triangle-up-open-dot"
	ScatterpolarglMarkerSymbolNumber20                ScatterpolarglMarkerSymbol = 20
	ScatterpolarglMarkerSymbol20                      ScatterpolarglMarkerSymbol = "20"
	ScatterpolarglMarkerSymbolStarTriangleDown        ScatterpolarglMarkerSymbol = "star-triangle-down"
	ScatterpolarglMarkerSymbolNumber120               ScatterpolarglMarkerSymbol = 120
	ScatterpolarglMarkerSymbol120                     ScatterpolarglMarkerSymbol = "120"
	ScatterpolarglMarkerSymbolStarTriangleDownOpen    ScatterpolarglMarkerSymbol = "star-triangle-down-open"
	ScatterpolarglMarkerSymbolNumber220               ScatterpolarglMarkerSymbol = 220
	ScatterpolarglMarkerSymbol220                     ScatterpolarglMarkerSymbol = "220"
	ScatterpolarglMarkerSymbolStarTriangleDownDot     ScatterpolarglMarkerSymbol = "star-triangle-down-dot"
	ScatterpolarglMarkerSymbolNumber320               ScatterpolarglMarkerSymbol = 320
	ScatterpolarglMarkerSymbol320                     ScatterpolarglMarkerSymbol = "320"
	ScatterpolarglMarkerSymbolStarTriangleDownOpenDot ScatterpolarglMarkerSymbol = "star-triangle-down-open-dot"
	ScatterpolarglMarkerSymbolNumber21                ScatterpolarglMarkerSymbol = 21
	ScatterpolarglMarkerSymbol21                      ScatterpolarglMarkerSymbol = "21"
	ScatterpolarglMarkerSymbolStarSquare              ScatterpolarglMarkerSymbol = "star-square"
	ScatterpolarglMarkerSymbolNumber121               ScatterpolarglMarkerSymbol = 121
	ScatterpolarglMarkerSymbol121                     ScatterpolarglMarkerSymbol = "121"
	ScatterpolarglMarkerSymbolStarSquareOpen          ScatterpolarglMarkerSymbol = "star-square-open"
	ScatterpolarglMarkerSymbolNumber221               ScatterpolarglMarkerSymbol = 221
	ScatterpolarglMarkerSymbol221                     ScatterpolarglMarkerSymbol = "221"
	ScatterpolarglMarkerSymbolStarSquareDot           ScatterpolarglMarkerSymbol = "star-square-dot"
	ScatterpolarglMarkerSymbolNumber321               ScatterpolarglMarkerSymbol = 321
	ScatterpolarglMarkerSymbol321                     ScatterpolarglMarkerSymbol = "321"
	ScatterpolarglMarkerSymbolStarSquareOpenDot       ScatterpolarglMarkerSymbol = "star-square-open-dot"
	ScatterpolarglMarkerSymbolNumber22                ScatterpolarglMarkerSymbol = 22
	ScatterpolarglMarkerSymbol22                      ScatterpolarglMarkerSymbol = "22"
	ScatterpolarglMarkerSymbolStarDiamond             ScatterpolarglMarkerSymbol = "star-diamond"
	ScatterpolarglMarkerSymbolNumber122               ScatterpolarglMarkerSymbol = 122
	ScatterpolarglMarkerSymbol122                     ScatterpolarglMarkerSymbol = "122"
	ScatterpolarglMarkerSymbolStarDiamondOpen         ScatterpolarglMarkerSymbol = "star-diamond-open"
	ScatterpolarglMarkerSymbolNumber222               ScatterpolarglMarkerSymbol = 222
	ScatterpolarglMarkerSymbol222                     ScatterpolarglMarkerSymbol = "222"
	ScatterpolarglMarkerSymbolStarDiamondDot          ScatterpolarglMarkerSymbol = "star-diamond-dot"
	ScatterpolarglMarkerSymbolNumber322               ScatterpolarglMarkerSymbol = 322
	ScatterpolarglMarkerSymbol322                     ScatterpolarglMarkerSymbol = "322"
	ScatterpolarglMarkerSymbolStarDiamondOpenDot      ScatterpolarglMarkerSymbol = "star-diamond-open-dot"
	ScatterpolarglMarkerSymbolNumber23                ScatterpolarglMarkerSymbol = 23
	ScatterpolarglMarkerSymbol23                      ScatterpolarglMarkerSymbol = "23"
	ScatterpolarglMarkerSymbolDiamondTall             ScatterpolarglMarkerSymbol = "diamond-tall"
	ScatterpolarglMarkerSymbolNumber123               ScatterpolarglMarkerSymbol = 123
	ScatterpolarglMarkerSymbol123                     ScatterpolarglMarkerSymbol = "123"
	ScatterpolarglMarkerSymbolDiamondTallOpen         ScatterpolarglMarkerSymbol = "diamond-tall-open"
	ScatterpolarglMarkerSymbolNumber223               ScatterpolarglMarkerSymbol = 223
	ScatterpolarglMarkerSymbol223                     ScatterpolarglMarkerSymbol = "223"
	ScatterpolarglMarkerSymbolDiamondTallDot          ScatterpolarglMarkerSymbol = "diamond-tall-dot"
	ScatterpolarglMarkerSymbolNumber323               ScatterpolarglMarkerSymbol = 323
	ScatterpolarglMarkerSymbol323                     ScatterpolarglMarkerSymbol = "323"
	ScatterpolarglMarkerSymbolDiamondTallOpenDot      ScatterpolarglMarkerSymbol = "diamond-tall-open-dot"
	ScatterpolarglMarkerSymbolNumber24                ScatterpolarglMarkerSymbol = 24
	ScatterpolarglMarkerSymbol24                      ScatterpolarglMarkerSymbol = "24"
	ScatterpolarglMarkerSymbolDiamondWide             ScatterpolarglMarkerSymbol = "diamond-wide"
	ScatterpolarglMarkerSymbolNumber124               ScatterpolarglMarkerSymbol = 124
	ScatterpolarglMarkerSymbol124                     ScatterpolarglMarkerSymbol = "124"
	ScatterpolarglMarkerSymbolDiamondWideOpen         ScatterpolarglMarkerSymbol = "diamond-wide-open"
	ScatterpolarglMarkerSymbolNumber224               ScatterpolarglMarkerSymbol = 224
	ScatterpolarglMarkerSymbol224                     ScatterpolarglMarkerSymbol = "224"
	ScatterpolarglMarkerSymbolDiamondWideDot          ScatterpolarglMarkerSymbol = "diamond-wide-dot"
	ScatterpolarglMarkerSymbolNumber324               ScatterpolarglMarkerSymbol = 324
	ScatterpolarglMarkerSymbol324                     ScatterpolarglMarkerSymbol = "324"
	ScatterpolarglMarkerSymbolDiamondWideOpenDot      ScatterpolarglMarkerSymbol = "diamond-wide-open-dot"
	ScatterpolarglMarkerSymbolNumber25                ScatterpolarglMarkerSymbol = 25
	ScatterpolarglMarkerSymbol25                      ScatterpolarglMarkerSymbol = "25"
	ScatterpolarglMarkerSymbolHourglass               ScatterpolarglMarkerSymbol = "hourglass"
	ScatterpolarglMarkerSymbolNumber125               ScatterpolarglMarkerSymbol = 125
	ScatterpolarglMarkerSymbol125                     ScatterpolarglMarkerSymbol = "125"
	ScatterpolarglMarkerSymbolHourglassOpen           ScatterpolarglMarkerSymbol = "hourglass-open"
	ScatterpolarglMarkerSymbolNumber26                ScatterpolarglMarkerSymbol = 26
	ScatterpolarglMarkerSymbol26                      ScatterpolarglMarkerSymbol = "26"
	ScatterpolarglMarkerSymbolBowtie                  ScatterpolarglMarkerSymbol = "bowtie"
	ScatterpolarglMarkerSymbolNumber126               ScatterpolarglMarkerSymbol = 126
	ScatterpolarglMarkerSymbol126                     ScatterpolarglMarkerSymbol = "126"
	ScatterpolarglMarkerSymbolBowtieOpen              ScatterpolarglMarkerSymbol = "bowtie-open"
	ScatterpolarglMarkerSymbolNumber27                ScatterpolarglMarkerSymbol = 27
	ScatterpolarglMarkerSymbol27                      ScatterpolarglMarkerSymbol = "27"
	ScatterpolarglMarkerSymbolCircleCross             ScatterpolarglMarkerSymbol = "circle-cross"
	ScatterpolarglMarkerSymbolNumber127               ScatterpolarglMarkerSymbol = 127
	ScatterpolarglMarkerSymbol127                     ScatterpolarglMarkerSymbol = "127"
	ScatterpolarglMarkerSymbolCircleCrossOpen         ScatterpolarglMarkerSymbol = "circle-cross-open"
	ScatterpolarglMarkerSymbolNumber28                ScatterpolarglMarkerSymbol = 28
	ScatterpolarglMarkerSymbol28                      ScatterpolarglMarkerSymbol = "28"
	ScatterpolarglMarkerSymbolCircleX                 ScatterpolarglMarkerSymbol = "circle-x"
	ScatterpolarglMarkerSymbolNumber128               ScatterpolarglMarkerSymbol = 128
	ScatterpolarglMarkerSymbol128                     ScatterpolarglMarkerSymbol = "128"
	ScatterpolarglMarkerSymbolCircleXOpen             ScatterpolarglMarkerSymbol = "circle-x-open"
	ScatterpolarglMarkerSymbolNumber29                ScatterpolarglMarkerSymbol = 29
	ScatterpolarglMarkerSymbol29                      ScatterpolarglMarkerSymbol = "29"
	ScatterpolarglMarkerSymbolSquareCross             ScatterpolarglMarkerSymbol = "square-cross"
	ScatterpolarglMarkerSymbolNumber129               ScatterpolarglMarkerSymbol = 129
	ScatterpolarglMarkerSymbol129                     ScatterpolarglMarkerSymbol = "129"
	ScatterpolarglMarkerSymbolSquareCrossOpen         ScatterpolarglMarkerSymbol = "square-cross-open"
	ScatterpolarglMarkerSymbolNumber30                ScatterpolarglMarkerSymbol = 30
	ScatterpolarglMarkerSymbol30                      ScatterpolarglMarkerSymbol = "30"
	ScatterpolarglMarkerSymbolSquareX                 ScatterpolarglMarkerSymbol = "square-x"
	ScatterpolarglMarkerSymbolNumber130               ScatterpolarglMarkerSymbol = 130
	ScatterpolarglMarkerSymbol130                     ScatterpolarglMarkerSymbol = "130"
	ScatterpolarglMarkerSymbolSquareXOpen             ScatterpolarglMarkerSymbol = "square-x-open"
	ScatterpolarglMarkerSymbolNumber31                ScatterpolarglMarkerSymbol = 31
	ScatterpolarglMarkerSymbol31                      ScatterpolarglMarkerSymbol = "31"
	ScatterpolarglMarkerSymbolDiamondCross            ScatterpolarglMarkerSymbol = "diamond-cross"
	ScatterpolarglMarkerSymbolNumber131               ScatterpolarglMarkerSymbol = 131
	ScatterpolarglMarkerSymbol131                     ScatterpolarglMarkerSymbol = "131"
	ScatterpolarglMarkerSymbolDiamondCrossOpen        ScatterpolarglMarkerSymbol = "diamond-cross-open"
	ScatterpolarglMarkerSymbolNumber32                ScatterpolarglMarkerSymbol = 32
	ScatterpolarglMarkerSymbol32                      ScatterpolarglMarkerSymbol = "32"
	ScatterpolarglMarkerSymbolDiamondX                ScatterpolarglMarkerSymbol = "diamond-x"
	ScatterpolarglMarkerSymbolNumber132               ScatterpolarglMarkerSymbol = 132
	ScatterpolarglMarkerSymbol132                     ScatterpolarglMarkerSymbol = "132"
	ScatterpolarglMarkerSymbolDiamondXOpen            ScatterpolarglMarkerSymbol = "diamond-x-open"
	ScatterpolarglMarkerSymbolNumber33                ScatterpolarglMarkerSymbol = 33
	ScatterpolarglMarkerSymbol33                      ScatterpolarglMarkerSymbol = "33"
	ScatterpolarglMarkerSymbolCrossThin               ScatterpolarglMarkerSymbol = "cross-thin"
	ScatterpolarglMarkerSymbolNumber133               ScatterpolarglMarkerSymbol = 133
	ScatterpolarglMarkerSymbol133                     ScatterpolarglMarkerSymbol = "133"
	ScatterpolarglMarkerSymbolCrossThinOpen           ScatterpolarglMarkerSymbol = "cross-thin-open"
	ScatterpolarglMarkerSymbolNumber34                ScatterpolarglMarkerSymbol = 34
	ScatterpolarglMarkerSymbol34                      ScatterpolarglMarkerSymbol = "34"
	ScatterpolarglMarkerSymbolXThin                   ScatterpolarglMarkerSymbol = "x-thin"
	ScatterpolarglMarkerSymbolNumber134               ScatterpolarglMarkerSymbol = 134
	ScatterpolarglMarkerSymbol134                     ScatterpolarglMarkerSymbol = "134"
	ScatterpolarglMarkerSymbolXThinOpen               ScatterpolarglMarkerSymbol = "x-thin-open"
	ScatterpolarglMarkerSymbolNumber35                ScatterpolarglMarkerSymbol = 35
	ScatterpolarglMarkerSymbol35                      ScatterpolarglMarkerSymbol = "35"
	ScatterpolarglMarkerSymbolAsterisk                ScatterpolarglMarkerSymbol = "asterisk"
	ScatterpolarglMarkerSymbolNumber135               ScatterpolarglMarkerSymbol = 135
	ScatterpolarglMarkerSymbol135                     ScatterpolarglMarkerSymbol = "135"
	ScatterpolarglMarkerSymbolAsteriskOpen            ScatterpolarglMarkerSymbol = "asterisk-open"
	ScatterpolarglMarkerSymbolNumber36                ScatterpolarglMarkerSymbol = 36
	ScatterpolarglMarkerSymbol36                      ScatterpolarglMarkerSymbol = "36"
	ScatterpolarglMarkerSymbolHash                    ScatterpolarglMarkerSymbol = "hash"
	ScatterpolarglMarkerSymbolNumber136               ScatterpolarglMarkerSymbol = 136
	ScatterpolarglMarkerSymbol136                     ScatterpolarglMarkerSymbol = "136"
	ScatterpolarglMarkerSymbolHashOpen                ScatterpolarglMarkerSymbol = "hash-open"
	ScatterpolarglMarkerSymbolNumber236               ScatterpolarglMarkerSymbol = 236
	ScatterpolarglMarkerSymbol236                     ScatterpolarglMarkerSymbol = "236"
	ScatterpolarglMarkerSymbolHashDot                 ScatterpolarglMarkerSymbol = "hash-dot"
	ScatterpolarglMarkerSymbolNumber336               ScatterpolarglMarkerSymbol = 336
	ScatterpolarglMarkerSymbol336                     ScatterpolarglMarkerSymbol = "336"
	ScatterpolarglMarkerSymbolHashOpenDot             ScatterpolarglMarkerSymbol = "hash-open-dot"
	ScatterpolarglMarkerSymbolNumber37                ScatterpolarglMarkerSymbol = 37
	ScatterpolarglMarkerSymbol37                      ScatterpolarglMarkerSymbol = "37"
	ScatterpolarglMarkerSymbolYUp                     ScatterpolarglMarkerSymbol = "y-up"
	ScatterpolarglMarkerSymbolNumber137               ScatterpolarglMarkerSymbol = 137
	ScatterpolarglMarkerSymbol137                     ScatterpolarglMarkerSymbol = "137"
	ScatterpolarglMarkerSymbolYUpOpen                 ScatterpolarglMarkerSymbol = "y-up-open"
	ScatterpolarglMarkerSymbolNumber38                ScatterpolarglMarkerSymbol = 38
	ScatterpolarglMarkerSymbol38                      ScatterpolarglMarkerSymbol = "38"
	ScatterpolarglMarkerSymbolYDown                   ScatterpolarglMarkerSymbol = "y-down"
	ScatterpolarglMarkerSymbolNumber138               ScatterpolarglMarkerSymbol = 138
	ScatterpolarglMarkerSymbol138                     ScatterpolarglMarkerSymbol = "138"
	ScatterpolarglMarkerSymbolYDownOpen               ScatterpolarglMarkerSymbol = "y-down-open"
	ScatterpolarglMarkerSymbolNumber39                ScatterpolarglMarkerSymbol = 39
	ScatterpolarglMarkerSymbol39                      ScatterpolarglMarkerSymbol = "39"
	ScatterpolarglMarkerSymbolYLeft                   ScatterpolarglMarkerSymbol = "y-left"
	ScatterpolarglMarkerSymbolNumber139               ScatterpolarglMarkerSymbol = 139
	ScatterpolarglMarkerSymbol139                     ScatterpolarglMarkerSymbol = "139"
	ScatterpolarglMarkerSymbolYLeftOpen               ScatterpolarglMarkerSymbol = "y-left-open"
	ScatterpolarglMarkerSymbolNumber40                ScatterpolarglMarkerSymbol = 40
	ScatterpolarglMarkerSymbol40                      ScatterpolarglMarkerSymbol = "40"
	ScatterpolarglMarkerSymbolYRight                  ScatterpolarglMarkerSymbol = "y-right"
	ScatterpolarglMarkerSymbolNumber140               ScatterpolarglMarkerSymbol = 140
	ScatterpolarglMarkerSymbol140                     ScatterpolarglMarkerSymbol = "140"
	ScatterpolarglMarkerSymbolYRightOpen              ScatterpolarglMarkerSymbol = "y-right-open"
	ScatterpolarglMarkerSymbolNumber41                ScatterpolarglMarkerSymbol = 41
	ScatterpolarglMarkerSymbol41                      ScatterpolarglMarkerSymbol = "41"
	ScatterpolarglMarkerSymbolLineEw                  ScatterpolarglMarkerSymbol = "line-ew"
	ScatterpolarglMarkerSymbolNumber141               ScatterpolarglMarkerSymbol = 141
	ScatterpolarglMarkerSymbol141                     ScatterpolarglMarkerSymbol = "141"
	ScatterpolarglMarkerSymbolLineEwOpen              ScatterpolarglMarkerSymbol = "line-ew-open"
	ScatterpolarglMarkerSymbolNumber42                ScatterpolarglMarkerSymbol = 42
	ScatterpolarglMarkerSymbol42                      ScatterpolarglMarkerSymbol = "42"
	ScatterpolarglMarkerSymbolLineNs                  ScatterpolarglMarkerSymbol = "line-ns"
	ScatterpolarglMarkerSymbolNumber142               ScatterpolarglMarkerSymbol = 142
	ScatterpolarglMarkerSymbol142                     ScatterpolarglMarkerSymbol = "142"
	ScatterpolarglMarkerSymbolLineNsOpen              ScatterpolarglMarkerSymbol = "line-ns-open"
	ScatterpolarglMarkerSymbolNumber43                ScatterpolarglMarkerSymbol = 43
	ScatterpolarglMarkerSymbol43                      ScatterpolarglMarkerSymbol = "43"
	ScatterpolarglMarkerSymbolLineNe                  ScatterpolarglMarkerSymbol = "line-ne"
	ScatterpolarglMarkerSymbolNumber143               ScatterpolarglMarkerSymbol = 143
	ScatterpolarglMarkerSymbol143                     ScatterpolarglMarkerSymbol = "143"
	ScatterpolarglMarkerSymbolLineNeOpen              ScatterpolarglMarkerSymbol = "line-ne-open"
	ScatterpolarglMarkerSymbolNumber44                ScatterpolarglMarkerSymbol = 44
	ScatterpolarglMarkerSymbol44                      ScatterpolarglMarkerSymbol = "44"
	ScatterpolarglMarkerSymbolLineNw                  ScatterpolarglMarkerSymbol = "line-nw"
	ScatterpolarglMarkerSymbolNumber144               ScatterpolarglMarkerSymbol = 144
	ScatterpolarglMarkerSymbol144                     ScatterpolarglMarkerSymbol = "144"
	ScatterpolarglMarkerSymbolLineNwOpen              ScatterpolarglMarkerSymbol = "line-nw-open"
	ScatterpolarglMarkerSymbolNumber45                ScatterpolarglMarkerSymbol = 45
	ScatterpolarglMarkerSymbol45                      ScatterpolarglMarkerSymbol = "45"
	ScatterpolarglMarkerSymbolArrowUp                 ScatterpolarglMarkerSymbol = "arrow-up"
	ScatterpolarglMarkerSymbolNumber145               ScatterpolarglMarkerSymbol = 145
	ScatterpolarglMarkerSymbol145                     ScatterpolarglMarkerSymbol = "145"
	ScatterpolarglMarkerSymbolArrowUpOpen             ScatterpolarglMarkerSymbol = "arrow-up-open"
	ScatterpolarglMarkerSymbolNumber46                ScatterpolarglMarkerSymbol = 46
	ScatterpolarglMarkerSymbol46                      ScatterpolarglMarkerSymbol = "46"
	ScatterpolarglMarkerSymbolArrowDown               ScatterpolarglMarkerSymbol = "arrow-down"
	ScatterpolarglMarkerSymbolNumber146               ScatterpolarglMarkerSymbol = 146
	ScatterpolarglMarkerSymbol146                     ScatterpolarglMarkerSymbol = "146"
	ScatterpolarglMarkerSymbolArrowDownOpen           ScatterpolarglMarkerSymbol = "arrow-down-open"
	ScatterpolarglMarkerSymbolNumber47                ScatterpolarglMarkerSymbol = 47
	ScatterpolarglMarkerSymbol47                      ScatterpolarglMarkerSymbol = "47"
	ScatterpolarglMarkerSymbolArrowLeft               ScatterpolarglMarkerSymbol = "arrow-left"
	ScatterpolarglMarkerSymbolNumber147               ScatterpolarglMarkerSymbol = 147
	ScatterpolarglMarkerSymbol147                     ScatterpolarglMarkerSymbol = "147"
	ScatterpolarglMarkerSymbolArrowLeftOpen           ScatterpolarglMarkerSymbol = "arrow-left-open"
	ScatterpolarglMarkerSymbolNumber48                ScatterpolarglMarkerSymbol = 48
	ScatterpolarglMarkerSymbol48                      ScatterpolarglMarkerSymbol = "48"
	ScatterpolarglMarkerSymbolArrowRight              ScatterpolarglMarkerSymbol = "arrow-right"
	ScatterpolarglMarkerSymbolNumber148               ScatterpolarglMarkerSymbol = 148
	ScatterpolarglMarkerSymbol148                     ScatterpolarglMarkerSymbol = "148"
	ScatterpolarglMarkerSymbolArrowRightOpen          ScatterpolarglMarkerSymbol = "arrow-right-open"
	ScatterpolarglMarkerSymbolNumber49                ScatterpolarglMarkerSymbol = 49
	ScatterpolarglMarkerSymbol49                      ScatterpolarglMarkerSymbol = "49"
	ScatterpolarglMarkerSymbolArrowBarUp              ScatterpolarglMarkerSymbol = "arrow-bar-up"
	ScatterpolarglMarkerSymbolNumber149               ScatterpolarglMarkerSymbol = 149
	ScatterpolarglMarkerSymbol149                     ScatterpolarglMarkerSymbol = "149"
	ScatterpolarglMarkerSymbolArrowBarUpOpen          ScatterpolarglMarkerSymbol = "arrow-bar-up-open"
	ScatterpolarglMarkerSymbolNumber50                ScatterpolarglMarkerSymbol = 50
	ScatterpolarglMarkerSymbol50                      ScatterpolarglMarkerSymbol = "50"
	ScatterpolarglMarkerSymbolArrowBarDown            ScatterpolarglMarkerSymbol = "arrow-bar-down"
	ScatterpolarglMarkerSymbolNumber150               ScatterpolarglMarkerSymbol = 150
	ScatterpolarglMarkerSymbol150                     ScatterpolarglMarkerSymbol = "150"
	ScatterpolarglMarkerSymbolArrowBarDownOpen        ScatterpolarglMarkerSymbol = "arrow-bar-down-open"
	ScatterpolarglMarkerSymbolNumber51                ScatterpolarglMarkerSymbol = 51
	ScatterpolarglMarkerSymbol51                      ScatterpolarglMarkerSymbol = "51"
	ScatterpolarglMarkerSymbolArrowBarLeft            ScatterpolarglMarkerSymbol = "arrow-bar-left"
	ScatterpolarglMarkerSymbolNumber151               ScatterpolarglMarkerSymbol = 151
	ScatterpolarglMarkerSymbol151                     ScatterpolarglMarkerSymbol = "151"
	ScatterpolarglMarkerSymbolArrowBarLeftOpen        ScatterpolarglMarkerSymbol = "arrow-bar-left-open"
	ScatterpolarglMarkerSymbolNumber52                ScatterpolarglMarkerSymbol = 52
	ScatterpolarglMarkerSymbol52                      ScatterpolarglMarkerSymbol = "52"
	ScatterpolarglMarkerSymbolArrowBarRight           ScatterpolarglMarkerSymbol = "arrow-bar-right"
	ScatterpolarglMarkerSymbolNumber152               ScatterpolarglMarkerSymbol = 152
	ScatterpolarglMarkerSymbol152                     ScatterpolarglMarkerSymbol = "152"
	ScatterpolarglMarkerSymbolArrowBarRightOpen       ScatterpolarglMarkerSymbol = "arrow-bar-right-open"
)

// ScatterpolarglTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterpolarglTextposition string

const (
	ScatterpolarglTextpositionTopLeft      ScatterpolarglTextposition = "top left"
	ScatterpolarglTextpositionTopCenter    ScatterpolarglTextposition = "top center"
	ScatterpolarglTextpositionTopRight     ScatterpolarglTextposition = "top right"
	ScatterpolarglTextpositionMiddleLeft   ScatterpolarglTextposition = "middle left"
	ScatterpolarglTextpositionMiddleCenter ScatterpolarglTextposition = "middle center"
	ScatterpolarglTextpositionMiddleRight  ScatterpolarglTextposition = "middle right"
	ScatterpolarglTextpositionBottomLeft   ScatterpolarglTextposition = "bottom left"
	ScatterpolarglTextpositionBottomCenter ScatterpolarglTextposition = "bottom center"
	ScatterpolarglTextpositionBottomRight  ScatterpolarglTextposition = "bottom right"
)

// ScatterpolarglThetaunit Sets the unit of input *theta* values. Has an effect only when on *linear* angular axes.
type ScatterpolarglThetaunit string

const (
	ScatterpolarglThetaunitRadians  ScatterpolarglThetaunit = "radians"
	ScatterpolarglThetaunitDegrees  ScatterpolarglThetaunit = "degrees"
	ScatterpolarglThetaunitGradians ScatterpolarglThetaunit = "gradians"
)

// ScatterpolarglVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterpolarglVisible interface{}

var (
	ScatterpolarglVisibleTrue       ScatterpolarglVisible = true
	ScatterpolarglVisibleFalse      ScatterpolarglVisible = false
	ScatterpolarglVisibleLegendonly ScatterpolarglVisible = "legendonly"
)

// ScatterpolarglHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type ScatterpolarglHoverinfo string

const (
	// Flags
	ScatterpolarglHoverinfoR     ScatterpolarglHoverinfo = "r"
	ScatterpolarglHoverinfoTheta ScatterpolarglHoverinfo = "theta"
	ScatterpolarglHoverinfoText  ScatterpolarglHoverinfo = "text"
	ScatterpolarglHoverinfoName  ScatterpolarglHoverinfo = "name"

	// Extra
	ScatterpolarglHoverinfoAll  ScatterpolarglHoverinfo = "all"
	ScatterpolarglHoverinfoNone ScatterpolarglHoverinfo = "none"
	ScatterpolarglHoverinfoSkip ScatterpolarglHoverinfo = "skip"
)

// ScatterpolarglMode Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
type ScatterpolarglMode string

const (
	// Flags
	ScatterpolarglModeLines   ScatterpolarglMode = "lines"
	ScatterpolarglModeMarkers ScatterpolarglMode = "markers"
	ScatterpolarglModeText    ScatterpolarglMode = "text"

	// Extra
	ScatterpolarglModeNone ScatterpolarglMode = "none"
)
