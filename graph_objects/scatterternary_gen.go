package grob

var TraceTypeScatterternary TraceType = "scatterternary"

func (trace *Scatterternary) GetType() TraceType {
	return TraceTypeScatterternary
}

// Scatterternary Provides similar functionality to the *scatter* type but on a ternary phase diagram. The data is provided by at least two arrays out of `a`, `b`, `c` triplets.
type Scatterternary struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// A
	// arrayOK: false
	// type: data_array
	// Sets the quantity of component `a` in each data point. If `a`, `b`, and `c` are all provided, they need not be normalized, only the relative values matter. If only two arrays are provided they must be normalized to match `ternary<i>.sum`.
	A interface{} `json:"a,omitempty"`

	// Asrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  a .
	Asrc String `json:"asrc,omitempty"`

	// B
	// arrayOK: false
	// type: data_array
	// Sets the quantity of component `a` in each data point. If `a`, `b`, and `c` are all provided, they need not be normalized, only the relative values matter. If only two arrays are provided they must be normalized to match `ternary<i>.sum`.
	B interface{} `json:"b,omitempty"`

	// Bsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  b .
	Bsrc String `json:"bsrc,omitempty"`

	// C
	// arrayOK: false
	// type: data_array
	// Sets the quantity of component `a` in each data point. If `a`, `b`, and `c` are all provided, they need not be normalized, only the relative values matter. If only two arrays are provided they must be normalized to match `ternary<i>.sum`.
	C interface{} `json:"c,omitempty"`

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

	// Csrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  c .
	Csrc String `json:"csrc,omitempty"`

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
	Fill ScatterternaryFill `json:"fill,omitempty"`

	// Fillcolor
	// arrayOK: false
	// type: color
	// Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor Color `json:"fillcolor,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ScatterternaryHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *ScatterternaryHoverlabel `json:"hoverlabel,omitempty"`

	// Hoveron
	// default: %!s(<nil>)
	// type: flaglist
	// Do the hover effects highlight individual points (markers or line points) or do they highlight filled regions? If the fill is *toself* or *tonext* and there are no markers or text, then the default is *fills*, otherwise it is *points*.
	Hoveron ScatterternaryHoveron `json:"hoveron,omitempty"`

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
	// Sets hover text elements associated with each (a,b,c) point. If a single string, the same string appears over all the data points. If an array of strings, the items are mapped in order to the the data points in (a,b,c). To be seen, trace `hoverinfo` must contain a *text* flag.
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
	Line *ScatterternaryLine `json:"line,omitempty"`

	// Marker
	// role: Object
	Marker *ScatterternaryMarker `json:"marker,omitempty"`

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
	Mode ScatterternaryMode `json:"mode,omitempty"`

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
	Selected *ScatterternarySelected `json:"selected,omitempty"`

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
	Stream *ScatterternaryStream `json:"stream,omitempty"`

	// Subplot
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's data coordinates and a ternary subplot. If *ternary* (the default value), the data refer to `layout.ternary`. If *ternary2*, the data refer to `layout.ternary2`, and so on.
	Subplot String `json:"subplot,omitempty"`

	// Sum
	// arrayOK: false
	// type: number
	// The number each triplet should sum to, if only two of `a`, `b`, and `c` are provided. This overrides `ternary<i>.sum` to normalize this specific trace, but does not affect the values displayed on the axes. 0 (or missing) means to use ternary<i>.sum
	Sum float64 `json:"sum,omitempty"`

	// Text
	// arrayOK: true
	// type: string
	// Sets text elements associated with each (a,b,c) point. If a single string, the same string appears over all the data points. If an array of strings, the items are mapped in order to the the data points in (a,b,c). If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text String `json:"text,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScatterternaryTextfont `json:"textfont,omitempty"`

	// Textposition
	// default: middle center
	// type: enumerated
	// Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition ScatterternaryTextposition `json:"textposition,omitempty"`

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
	// Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `a`, `b`, `c` and `text`.
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
	Unselected *ScatterternaryUnselected `json:"unselected,omitempty"`

	// Visible
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ScatterternaryVisible `json:"visible,omitempty"`
}

// ScatterternaryHoverlabelFont Sets the font used in hover labels.
type ScatterternaryHoverlabelFont struct {

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

// ScatterternaryHoverlabel
type ScatterternaryHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align ScatterternaryHoverlabelAlign `json:"align,omitempty"`

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
	Font *ScatterternaryHoverlabelFont `json:"font,omitempty"`

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

// ScatterternaryLine
type ScatterternaryLine struct {

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
	Shape ScatterternaryLineShape `json:"shape,omitempty"`

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

// ScatterternaryMarkerColorbarTickfont Sets the color bar's tick label font
type ScatterternaryMarkerColorbarTickfont struct {

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

// ScatterternaryMarkerColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type ScatterternaryMarkerColorbarTitleFont struct {

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

// ScatterternaryMarkerColorbarTitle
type ScatterternaryMarkerColorbarTitle struct {

	// Font
	// role: Object
	Font *ScatterternaryMarkerColorbarTitleFont `json:"font,omitempty"`

	// Side
	// default: top
	// type: enumerated
	// Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
	Side ScatterternaryMarkerColorbarTitleSide `json:"side,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// ScatterternaryMarkerColorbar
type ScatterternaryMarkerColorbar struct {

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
	Exponentformat ScatterternaryMarkerColorbarExponentformat `json:"exponentformat,omitempty"`

	// Len
	// arrayOK: false
	// type: number
	// Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends.
	Len float64 `json:"len,omitempty"`

	// Lenmode
	// default: fraction
	// type: enumerated
	// Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
	Lenmode ScatterternaryMarkerColorbarLenmode `json:"lenmode,omitempty"`

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
	Showexponent ScatterternaryMarkerColorbarShowexponent `json:"showexponent,omitempty"`

	// Showticklabels
	// arrayOK: false
	// type: boolean
	// Determines whether or not the tick labels are drawn.
	Showticklabels Bool `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix ScatterternaryMarkerColorbarShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix ScatterternaryMarkerColorbarShowticksuffix `json:"showticksuffix,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels.
	Thickness float64 `json:"thickness,omitempty"`

	// Thicknessmode
	// default: pixels
	// type: enumerated
	// Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
	Thicknessmode ScatterternaryMarkerColorbarThicknessmode `json:"thicknessmode,omitempty"`

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
	Tickfont *ScatterternaryMarkerColorbarTickfont `json:"tickfont,omitempty"`

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
	Ticklabelposition ScatterternaryMarkerColorbarTicklabelposition `json:"ticklabelposition,omitempty"`

	// Ticklen
	// arrayOK: false
	// type: number
	// Sets the tick length (in px).
	Ticklen float64 `json:"ticklen,omitempty"`

	// Tickmode
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
	Tickmode ScatterternaryMarkerColorbarTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

	// Ticks
	// default:
	// type: enumerated
	// Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
	Ticks ScatterternaryMarkerColorbarTicks `json:"ticks,omitempty"`

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
	Title *ScatterternaryMarkerColorbarTitle `json:"title,omitempty"`

	// X
	// arrayOK: false
	// type: number
	// Sets the x position of the color bar (in plot fraction).
	X float64 `json:"x,omitempty"`

	// Xanchor
	// default: left
	// type: enumerated
	// Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
	Xanchor ScatterternaryMarkerColorbarXanchor `json:"xanchor,omitempty"`

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
	Yanchor ScatterternaryMarkerColorbarYanchor `json:"yanchor,omitempty"`

	// Ypad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the y direction.
	Ypad float64 `json:"ypad,omitempty"`
}

// ScatterternaryMarkerGradient
type ScatterternaryMarkerGradient struct {

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
	Type ScatterternaryMarkerGradientType `json:"type,omitempty"`

	// Typesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  type .
	Typesrc String `json:"typesrc,omitempty"`
}

// ScatterternaryMarkerLine
type ScatterternaryMarkerLine struct {

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

// ScatterternaryMarker
type ScatterternaryMarker struct {

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
	Colorbar *ScatterternaryMarkerColorbar `json:"colorbar,omitempty"`

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
	Gradient *ScatterternaryMarkerGradient `json:"gradient,omitempty"`

	// Line
	// role: Object
	Line *ScatterternaryMarkerLine `json:"line,omitempty"`

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
	Sizemode ScatterternaryMarkerSizemode `json:"sizemode,omitempty"`

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
	Symbol ScatterternaryMarkerSymbol `json:"symbol,omitempty"`

	// Symbolsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  symbol .
	Symbolsrc String `json:"symbolsrc,omitempty"`
}

// ScatterternarySelectedMarker
type ScatterternarySelectedMarker struct {

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

// ScatterternarySelectedTextfont
type ScatterternarySelectedTextfont struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the text font color of selected points.
	Color Color `json:"color,omitempty"`
}

// ScatterternarySelected
type ScatterternarySelected struct {

	// Marker
	// role: Object
	Marker *ScatterternarySelectedMarker `json:"marker,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScatterternarySelectedTextfont `json:"textfont,omitempty"`
}

// ScatterternaryStream
type ScatterternaryStream struct {

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

// ScatterternaryTextfont Sets the text font.
type ScatterternaryTextfont struct {

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

// ScatterternaryUnselectedMarker
type ScatterternaryUnselectedMarker struct {

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

// ScatterternaryUnselectedTextfont
type ScatterternaryUnselectedTextfont struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the text font color of unselected points, applied only when a selection exists.
	Color Color `json:"color,omitempty"`
}

// ScatterternaryUnselected
type ScatterternaryUnselected struct {

	// Marker
	// role: Object
	Marker *ScatterternaryUnselectedMarker `json:"marker,omitempty"`

	// Textfont
	// role: Object
	Textfont *ScatterternaryUnselectedTextfont `json:"textfont,omitempty"`
}

// ScatterternaryFill Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. scatterternary has a subset of the options available to scatter. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other.
type ScatterternaryFill string

const (
	ScatterternaryFillNone   ScatterternaryFill = "none"
	ScatterternaryFillToself ScatterternaryFill = "toself"
	ScatterternaryFillTonext ScatterternaryFill = "tonext"
)

// ScatterternaryHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterternaryHoverlabelAlign string

const (
	ScatterternaryHoverlabelAlignLeft  ScatterternaryHoverlabelAlign = "left"
	ScatterternaryHoverlabelAlignRight ScatterternaryHoverlabelAlign = "right"
	ScatterternaryHoverlabelAlignAuto  ScatterternaryHoverlabelAlign = "auto"
)

// ScatterternaryLineShape Determines the line shape. With *spline* the lines are drawn using spline interpolation. The other available values correspond to step-wise line shapes.
type ScatterternaryLineShape string

const (
	ScatterternaryLineShapeLinear ScatterternaryLineShape = "linear"
	ScatterternaryLineShapeSpline ScatterternaryLineShape = "spline"
)

// ScatterternaryMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterternaryMarkerColorbarExponentformat string

const (
	ScatterternaryMarkerColorbarExponentformatNone  ScatterternaryMarkerColorbarExponentformat = "none"
	ScatterternaryMarkerColorbarExponentformatE1    ScatterternaryMarkerColorbarExponentformat = "e"
	ScatterternaryMarkerColorbarExponentformatE2    ScatterternaryMarkerColorbarExponentformat = "E"
	ScatterternaryMarkerColorbarExponentformatPower ScatterternaryMarkerColorbarExponentformat = "power"
	ScatterternaryMarkerColorbarExponentformatSi    ScatterternaryMarkerColorbarExponentformat = "SI"
	ScatterternaryMarkerColorbarExponentformatB     ScatterternaryMarkerColorbarExponentformat = "B"
)

// ScatterternaryMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterternaryMarkerColorbarLenmode string

const (
	ScatterternaryMarkerColorbarLenmodeFraction ScatterternaryMarkerColorbarLenmode = "fraction"
	ScatterternaryMarkerColorbarLenmodePixels   ScatterternaryMarkerColorbarLenmode = "pixels"
)

// ScatterternaryMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterternaryMarkerColorbarShowexponent string

const (
	ScatterternaryMarkerColorbarShowexponentAll   ScatterternaryMarkerColorbarShowexponent = "all"
	ScatterternaryMarkerColorbarShowexponentFirst ScatterternaryMarkerColorbarShowexponent = "first"
	ScatterternaryMarkerColorbarShowexponentLast  ScatterternaryMarkerColorbarShowexponent = "last"
	ScatterternaryMarkerColorbarShowexponentNone  ScatterternaryMarkerColorbarShowexponent = "none"
)

// ScatterternaryMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterternaryMarkerColorbarShowtickprefix string

const (
	ScatterternaryMarkerColorbarShowtickprefixAll   ScatterternaryMarkerColorbarShowtickprefix = "all"
	ScatterternaryMarkerColorbarShowtickprefixFirst ScatterternaryMarkerColorbarShowtickprefix = "first"
	ScatterternaryMarkerColorbarShowtickprefixLast  ScatterternaryMarkerColorbarShowtickprefix = "last"
	ScatterternaryMarkerColorbarShowtickprefixNone  ScatterternaryMarkerColorbarShowtickprefix = "none"
)

// ScatterternaryMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterternaryMarkerColorbarShowticksuffix string

const (
	ScatterternaryMarkerColorbarShowticksuffixAll   ScatterternaryMarkerColorbarShowticksuffix = "all"
	ScatterternaryMarkerColorbarShowticksuffixFirst ScatterternaryMarkerColorbarShowticksuffix = "first"
	ScatterternaryMarkerColorbarShowticksuffixLast  ScatterternaryMarkerColorbarShowticksuffix = "last"
	ScatterternaryMarkerColorbarShowticksuffixNone  ScatterternaryMarkerColorbarShowticksuffix = "none"
)

// ScatterternaryMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterternaryMarkerColorbarThicknessmode string

const (
	ScatterternaryMarkerColorbarThicknessmodeFraction ScatterternaryMarkerColorbarThicknessmode = "fraction"
	ScatterternaryMarkerColorbarThicknessmodePixels   ScatterternaryMarkerColorbarThicknessmode = "pixels"
)

// ScatterternaryMarkerColorbarTicklabelposition Determines where tick labels are drawn.
type ScatterternaryMarkerColorbarTicklabelposition string

const (
	ScatterternaryMarkerColorbarTicklabelpositionOutside       ScatterternaryMarkerColorbarTicklabelposition = "outside"
	ScatterternaryMarkerColorbarTicklabelpositionInside        ScatterternaryMarkerColorbarTicklabelposition = "inside"
	ScatterternaryMarkerColorbarTicklabelpositionOutsideTop    ScatterternaryMarkerColorbarTicklabelposition = "outside top"
	ScatterternaryMarkerColorbarTicklabelpositionInsideTop     ScatterternaryMarkerColorbarTicklabelposition = "inside top"
	ScatterternaryMarkerColorbarTicklabelpositionOutsideBottom ScatterternaryMarkerColorbarTicklabelposition = "outside bottom"
	ScatterternaryMarkerColorbarTicklabelpositionInsideBottom  ScatterternaryMarkerColorbarTicklabelposition = "inside bottom"
)

// ScatterternaryMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterternaryMarkerColorbarTickmode string

const (
	ScatterternaryMarkerColorbarTickmodeAuto   ScatterternaryMarkerColorbarTickmode = "auto"
	ScatterternaryMarkerColorbarTickmodeLinear ScatterternaryMarkerColorbarTickmode = "linear"
	ScatterternaryMarkerColorbarTickmodeArray  ScatterternaryMarkerColorbarTickmode = "array"
)

// ScatterternaryMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterternaryMarkerColorbarTicks string

const (
	ScatterternaryMarkerColorbarTicksOutside ScatterternaryMarkerColorbarTicks = "outside"
	ScatterternaryMarkerColorbarTicksInside  ScatterternaryMarkerColorbarTicks = "inside"
	ScatterternaryMarkerColorbarTicksEmpty   ScatterternaryMarkerColorbarTicks = ""
)

// ScatterternaryMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterternaryMarkerColorbarTitleSide string

const (
	ScatterternaryMarkerColorbarTitleSideRight  ScatterternaryMarkerColorbarTitleSide = "right"
	ScatterternaryMarkerColorbarTitleSideTop    ScatterternaryMarkerColorbarTitleSide = "top"
	ScatterternaryMarkerColorbarTitleSideBottom ScatterternaryMarkerColorbarTitleSide = "bottom"
)

// ScatterternaryMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterternaryMarkerColorbarXanchor string

const (
	ScatterternaryMarkerColorbarXanchorLeft   ScatterternaryMarkerColorbarXanchor = "left"
	ScatterternaryMarkerColorbarXanchorCenter ScatterternaryMarkerColorbarXanchor = "center"
	ScatterternaryMarkerColorbarXanchorRight  ScatterternaryMarkerColorbarXanchor = "right"
)

// ScatterternaryMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterternaryMarkerColorbarYanchor string

const (
	ScatterternaryMarkerColorbarYanchorTop    ScatterternaryMarkerColorbarYanchor = "top"
	ScatterternaryMarkerColorbarYanchorMiddle ScatterternaryMarkerColorbarYanchor = "middle"
	ScatterternaryMarkerColorbarYanchorBottom ScatterternaryMarkerColorbarYanchor = "bottom"
)

// ScatterternaryMarkerGradientType Sets the type of gradient used to fill the markers
type ScatterternaryMarkerGradientType string

const (
	ScatterternaryMarkerGradientTypeRadial     ScatterternaryMarkerGradientType = "radial"
	ScatterternaryMarkerGradientTypeHorizontal ScatterternaryMarkerGradientType = "horizontal"
	ScatterternaryMarkerGradientTypeVertical   ScatterternaryMarkerGradientType = "vertical"
	ScatterternaryMarkerGradientTypeNone       ScatterternaryMarkerGradientType = "none"
)

// ScatterternaryMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterternaryMarkerSizemode string

const (
	ScatterternaryMarkerSizemodeDiameter ScatterternaryMarkerSizemode = "diameter"
	ScatterternaryMarkerSizemodeArea     ScatterternaryMarkerSizemode = "area"
)

// ScatterternaryMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterternaryMarkerSymbol interface{}

var (
	ScatterternaryMarkerSymbolNumber0                 ScatterternaryMarkerSymbol = 0
	ScatterternaryMarkerSymbol0                       ScatterternaryMarkerSymbol = "0"
	ScatterternaryMarkerSymbolCircle                  ScatterternaryMarkerSymbol = "circle"
	ScatterternaryMarkerSymbolNumber100               ScatterternaryMarkerSymbol = 100
	ScatterternaryMarkerSymbol100                     ScatterternaryMarkerSymbol = "100"
	ScatterternaryMarkerSymbolCircleOpen              ScatterternaryMarkerSymbol = "circle-open"
	ScatterternaryMarkerSymbolNumber200               ScatterternaryMarkerSymbol = 200
	ScatterternaryMarkerSymbol200                     ScatterternaryMarkerSymbol = "200"
	ScatterternaryMarkerSymbolCircleDot               ScatterternaryMarkerSymbol = "circle-dot"
	ScatterternaryMarkerSymbolNumber300               ScatterternaryMarkerSymbol = 300
	ScatterternaryMarkerSymbol300                     ScatterternaryMarkerSymbol = "300"
	ScatterternaryMarkerSymbolCircleOpenDot           ScatterternaryMarkerSymbol = "circle-open-dot"
	ScatterternaryMarkerSymbolNumber1                 ScatterternaryMarkerSymbol = 1
	ScatterternaryMarkerSymbol1                       ScatterternaryMarkerSymbol = "1"
	ScatterternaryMarkerSymbolSquare                  ScatterternaryMarkerSymbol = "square"
	ScatterternaryMarkerSymbolNumber101               ScatterternaryMarkerSymbol = 101
	ScatterternaryMarkerSymbol101                     ScatterternaryMarkerSymbol = "101"
	ScatterternaryMarkerSymbolSquareOpen              ScatterternaryMarkerSymbol = "square-open"
	ScatterternaryMarkerSymbolNumber201               ScatterternaryMarkerSymbol = 201
	ScatterternaryMarkerSymbol201                     ScatterternaryMarkerSymbol = "201"
	ScatterternaryMarkerSymbolSquareDot               ScatterternaryMarkerSymbol = "square-dot"
	ScatterternaryMarkerSymbolNumber301               ScatterternaryMarkerSymbol = 301
	ScatterternaryMarkerSymbol301                     ScatterternaryMarkerSymbol = "301"
	ScatterternaryMarkerSymbolSquareOpenDot           ScatterternaryMarkerSymbol = "square-open-dot"
	ScatterternaryMarkerSymbolNumber2                 ScatterternaryMarkerSymbol = 2
	ScatterternaryMarkerSymbol2                       ScatterternaryMarkerSymbol = "2"
	ScatterternaryMarkerSymbolDiamond                 ScatterternaryMarkerSymbol = "diamond"
	ScatterternaryMarkerSymbolNumber102               ScatterternaryMarkerSymbol = 102
	ScatterternaryMarkerSymbol102                     ScatterternaryMarkerSymbol = "102"
	ScatterternaryMarkerSymbolDiamondOpen             ScatterternaryMarkerSymbol = "diamond-open"
	ScatterternaryMarkerSymbolNumber202               ScatterternaryMarkerSymbol = 202
	ScatterternaryMarkerSymbol202                     ScatterternaryMarkerSymbol = "202"
	ScatterternaryMarkerSymbolDiamondDot              ScatterternaryMarkerSymbol = "diamond-dot"
	ScatterternaryMarkerSymbolNumber302               ScatterternaryMarkerSymbol = 302
	ScatterternaryMarkerSymbol302                     ScatterternaryMarkerSymbol = "302"
	ScatterternaryMarkerSymbolDiamondOpenDot          ScatterternaryMarkerSymbol = "diamond-open-dot"
	ScatterternaryMarkerSymbolNumber3                 ScatterternaryMarkerSymbol = 3
	ScatterternaryMarkerSymbol3                       ScatterternaryMarkerSymbol = "3"
	ScatterternaryMarkerSymbolCross                   ScatterternaryMarkerSymbol = "cross"
	ScatterternaryMarkerSymbolNumber103               ScatterternaryMarkerSymbol = 103
	ScatterternaryMarkerSymbol103                     ScatterternaryMarkerSymbol = "103"
	ScatterternaryMarkerSymbolCrossOpen               ScatterternaryMarkerSymbol = "cross-open"
	ScatterternaryMarkerSymbolNumber203               ScatterternaryMarkerSymbol = 203
	ScatterternaryMarkerSymbol203                     ScatterternaryMarkerSymbol = "203"
	ScatterternaryMarkerSymbolCrossDot                ScatterternaryMarkerSymbol = "cross-dot"
	ScatterternaryMarkerSymbolNumber303               ScatterternaryMarkerSymbol = 303
	ScatterternaryMarkerSymbol303                     ScatterternaryMarkerSymbol = "303"
	ScatterternaryMarkerSymbolCrossOpenDot            ScatterternaryMarkerSymbol = "cross-open-dot"
	ScatterternaryMarkerSymbolNumber4                 ScatterternaryMarkerSymbol = 4
	ScatterternaryMarkerSymbol4                       ScatterternaryMarkerSymbol = "4"
	ScatterternaryMarkerSymbolX                       ScatterternaryMarkerSymbol = "x"
	ScatterternaryMarkerSymbolNumber104               ScatterternaryMarkerSymbol = 104
	ScatterternaryMarkerSymbol104                     ScatterternaryMarkerSymbol = "104"
	ScatterternaryMarkerSymbolXOpen                   ScatterternaryMarkerSymbol = "x-open"
	ScatterternaryMarkerSymbolNumber204               ScatterternaryMarkerSymbol = 204
	ScatterternaryMarkerSymbol204                     ScatterternaryMarkerSymbol = "204"
	ScatterternaryMarkerSymbolXDot                    ScatterternaryMarkerSymbol = "x-dot"
	ScatterternaryMarkerSymbolNumber304               ScatterternaryMarkerSymbol = 304
	ScatterternaryMarkerSymbol304                     ScatterternaryMarkerSymbol = "304"
	ScatterternaryMarkerSymbolXOpenDot                ScatterternaryMarkerSymbol = "x-open-dot"
	ScatterternaryMarkerSymbolNumber5                 ScatterternaryMarkerSymbol = 5
	ScatterternaryMarkerSymbol5                       ScatterternaryMarkerSymbol = "5"
	ScatterternaryMarkerSymbolTriangleUp              ScatterternaryMarkerSymbol = "triangle-up"
	ScatterternaryMarkerSymbolNumber105               ScatterternaryMarkerSymbol = 105
	ScatterternaryMarkerSymbol105                     ScatterternaryMarkerSymbol = "105"
	ScatterternaryMarkerSymbolTriangleUpOpen          ScatterternaryMarkerSymbol = "triangle-up-open"
	ScatterternaryMarkerSymbolNumber205               ScatterternaryMarkerSymbol = 205
	ScatterternaryMarkerSymbol205                     ScatterternaryMarkerSymbol = "205"
	ScatterternaryMarkerSymbolTriangleUpDot           ScatterternaryMarkerSymbol = "triangle-up-dot"
	ScatterternaryMarkerSymbolNumber305               ScatterternaryMarkerSymbol = 305
	ScatterternaryMarkerSymbol305                     ScatterternaryMarkerSymbol = "305"
	ScatterternaryMarkerSymbolTriangleUpOpenDot       ScatterternaryMarkerSymbol = "triangle-up-open-dot"
	ScatterternaryMarkerSymbolNumber6                 ScatterternaryMarkerSymbol = 6
	ScatterternaryMarkerSymbol6                       ScatterternaryMarkerSymbol = "6"
	ScatterternaryMarkerSymbolTriangleDown            ScatterternaryMarkerSymbol = "triangle-down"
	ScatterternaryMarkerSymbolNumber106               ScatterternaryMarkerSymbol = 106
	ScatterternaryMarkerSymbol106                     ScatterternaryMarkerSymbol = "106"
	ScatterternaryMarkerSymbolTriangleDownOpen        ScatterternaryMarkerSymbol = "triangle-down-open"
	ScatterternaryMarkerSymbolNumber206               ScatterternaryMarkerSymbol = 206
	ScatterternaryMarkerSymbol206                     ScatterternaryMarkerSymbol = "206"
	ScatterternaryMarkerSymbolTriangleDownDot         ScatterternaryMarkerSymbol = "triangle-down-dot"
	ScatterternaryMarkerSymbolNumber306               ScatterternaryMarkerSymbol = 306
	ScatterternaryMarkerSymbol306                     ScatterternaryMarkerSymbol = "306"
	ScatterternaryMarkerSymbolTriangleDownOpenDot     ScatterternaryMarkerSymbol = "triangle-down-open-dot"
	ScatterternaryMarkerSymbolNumber7                 ScatterternaryMarkerSymbol = 7
	ScatterternaryMarkerSymbol7                       ScatterternaryMarkerSymbol = "7"
	ScatterternaryMarkerSymbolTriangleLeft            ScatterternaryMarkerSymbol = "triangle-left"
	ScatterternaryMarkerSymbolNumber107               ScatterternaryMarkerSymbol = 107
	ScatterternaryMarkerSymbol107                     ScatterternaryMarkerSymbol = "107"
	ScatterternaryMarkerSymbolTriangleLeftOpen        ScatterternaryMarkerSymbol = "triangle-left-open"
	ScatterternaryMarkerSymbolNumber207               ScatterternaryMarkerSymbol = 207
	ScatterternaryMarkerSymbol207                     ScatterternaryMarkerSymbol = "207"
	ScatterternaryMarkerSymbolTriangleLeftDot         ScatterternaryMarkerSymbol = "triangle-left-dot"
	ScatterternaryMarkerSymbolNumber307               ScatterternaryMarkerSymbol = 307
	ScatterternaryMarkerSymbol307                     ScatterternaryMarkerSymbol = "307"
	ScatterternaryMarkerSymbolTriangleLeftOpenDot     ScatterternaryMarkerSymbol = "triangle-left-open-dot"
	ScatterternaryMarkerSymbolNumber8                 ScatterternaryMarkerSymbol = 8
	ScatterternaryMarkerSymbol8                       ScatterternaryMarkerSymbol = "8"
	ScatterternaryMarkerSymbolTriangleRight           ScatterternaryMarkerSymbol = "triangle-right"
	ScatterternaryMarkerSymbolNumber108               ScatterternaryMarkerSymbol = 108
	ScatterternaryMarkerSymbol108                     ScatterternaryMarkerSymbol = "108"
	ScatterternaryMarkerSymbolTriangleRightOpen       ScatterternaryMarkerSymbol = "triangle-right-open"
	ScatterternaryMarkerSymbolNumber208               ScatterternaryMarkerSymbol = 208
	ScatterternaryMarkerSymbol208                     ScatterternaryMarkerSymbol = "208"
	ScatterternaryMarkerSymbolTriangleRightDot        ScatterternaryMarkerSymbol = "triangle-right-dot"
	ScatterternaryMarkerSymbolNumber308               ScatterternaryMarkerSymbol = 308
	ScatterternaryMarkerSymbol308                     ScatterternaryMarkerSymbol = "308"
	ScatterternaryMarkerSymbolTriangleRightOpenDot    ScatterternaryMarkerSymbol = "triangle-right-open-dot"
	ScatterternaryMarkerSymbolNumber9                 ScatterternaryMarkerSymbol = 9
	ScatterternaryMarkerSymbol9                       ScatterternaryMarkerSymbol = "9"
	ScatterternaryMarkerSymbolTriangleNe              ScatterternaryMarkerSymbol = "triangle-ne"
	ScatterternaryMarkerSymbolNumber109               ScatterternaryMarkerSymbol = 109
	ScatterternaryMarkerSymbol109                     ScatterternaryMarkerSymbol = "109"
	ScatterternaryMarkerSymbolTriangleNeOpen          ScatterternaryMarkerSymbol = "triangle-ne-open"
	ScatterternaryMarkerSymbolNumber209               ScatterternaryMarkerSymbol = 209
	ScatterternaryMarkerSymbol209                     ScatterternaryMarkerSymbol = "209"
	ScatterternaryMarkerSymbolTriangleNeDot           ScatterternaryMarkerSymbol = "triangle-ne-dot"
	ScatterternaryMarkerSymbolNumber309               ScatterternaryMarkerSymbol = 309
	ScatterternaryMarkerSymbol309                     ScatterternaryMarkerSymbol = "309"
	ScatterternaryMarkerSymbolTriangleNeOpenDot       ScatterternaryMarkerSymbol = "triangle-ne-open-dot"
	ScatterternaryMarkerSymbolNumber10                ScatterternaryMarkerSymbol = 10
	ScatterternaryMarkerSymbol10                      ScatterternaryMarkerSymbol = "10"
	ScatterternaryMarkerSymbolTriangleSe              ScatterternaryMarkerSymbol = "triangle-se"
	ScatterternaryMarkerSymbolNumber110               ScatterternaryMarkerSymbol = 110
	ScatterternaryMarkerSymbol110                     ScatterternaryMarkerSymbol = "110"
	ScatterternaryMarkerSymbolTriangleSeOpen          ScatterternaryMarkerSymbol = "triangle-se-open"
	ScatterternaryMarkerSymbolNumber210               ScatterternaryMarkerSymbol = 210
	ScatterternaryMarkerSymbol210                     ScatterternaryMarkerSymbol = "210"
	ScatterternaryMarkerSymbolTriangleSeDot           ScatterternaryMarkerSymbol = "triangle-se-dot"
	ScatterternaryMarkerSymbolNumber310               ScatterternaryMarkerSymbol = 310
	ScatterternaryMarkerSymbol310                     ScatterternaryMarkerSymbol = "310"
	ScatterternaryMarkerSymbolTriangleSeOpenDot       ScatterternaryMarkerSymbol = "triangle-se-open-dot"
	ScatterternaryMarkerSymbolNumber11                ScatterternaryMarkerSymbol = 11
	ScatterternaryMarkerSymbol11                      ScatterternaryMarkerSymbol = "11"
	ScatterternaryMarkerSymbolTriangleSw              ScatterternaryMarkerSymbol = "triangle-sw"
	ScatterternaryMarkerSymbolNumber111               ScatterternaryMarkerSymbol = 111
	ScatterternaryMarkerSymbol111                     ScatterternaryMarkerSymbol = "111"
	ScatterternaryMarkerSymbolTriangleSwOpen          ScatterternaryMarkerSymbol = "triangle-sw-open"
	ScatterternaryMarkerSymbolNumber211               ScatterternaryMarkerSymbol = 211
	ScatterternaryMarkerSymbol211                     ScatterternaryMarkerSymbol = "211"
	ScatterternaryMarkerSymbolTriangleSwDot           ScatterternaryMarkerSymbol = "triangle-sw-dot"
	ScatterternaryMarkerSymbolNumber311               ScatterternaryMarkerSymbol = 311
	ScatterternaryMarkerSymbol311                     ScatterternaryMarkerSymbol = "311"
	ScatterternaryMarkerSymbolTriangleSwOpenDot       ScatterternaryMarkerSymbol = "triangle-sw-open-dot"
	ScatterternaryMarkerSymbolNumber12                ScatterternaryMarkerSymbol = 12
	ScatterternaryMarkerSymbol12                      ScatterternaryMarkerSymbol = "12"
	ScatterternaryMarkerSymbolTriangleNw              ScatterternaryMarkerSymbol = "triangle-nw"
	ScatterternaryMarkerSymbolNumber112               ScatterternaryMarkerSymbol = 112
	ScatterternaryMarkerSymbol112                     ScatterternaryMarkerSymbol = "112"
	ScatterternaryMarkerSymbolTriangleNwOpen          ScatterternaryMarkerSymbol = "triangle-nw-open"
	ScatterternaryMarkerSymbolNumber212               ScatterternaryMarkerSymbol = 212
	ScatterternaryMarkerSymbol212                     ScatterternaryMarkerSymbol = "212"
	ScatterternaryMarkerSymbolTriangleNwDot           ScatterternaryMarkerSymbol = "triangle-nw-dot"
	ScatterternaryMarkerSymbolNumber312               ScatterternaryMarkerSymbol = 312
	ScatterternaryMarkerSymbol312                     ScatterternaryMarkerSymbol = "312"
	ScatterternaryMarkerSymbolTriangleNwOpenDot       ScatterternaryMarkerSymbol = "triangle-nw-open-dot"
	ScatterternaryMarkerSymbolNumber13                ScatterternaryMarkerSymbol = 13
	ScatterternaryMarkerSymbol13                      ScatterternaryMarkerSymbol = "13"
	ScatterternaryMarkerSymbolPentagon                ScatterternaryMarkerSymbol = "pentagon"
	ScatterternaryMarkerSymbolNumber113               ScatterternaryMarkerSymbol = 113
	ScatterternaryMarkerSymbol113                     ScatterternaryMarkerSymbol = "113"
	ScatterternaryMarkerSymbolPentagonOpen            ScatterternaryMarkerSymbol = "pentagon-open"
	ScatterternaryMarkerSymbolNumber213               ScatterternaryMarkerSymbol = 213
	ScatterternaryMarkerSymbol213                     ScatterternaryMarkerSymbol = "213"
	ScatterternaryMarkerSymbolPentagonDot             ScatterternaryMarkerSymbol = "pentagon-dot"
	ScatterternaryMarkerSymbolNumber313               ScatterternaryMarkerSymbol = 313
	ScatterternaryMarkerSymbol313                     ScatterternaryMarkerSymbol = "313"
	ScatterternaryMarkerSymbolPentagonOpenDot         ScatterternaryMarkerSymbol = "pentagon-open-dot"
	ScatterternaryMarkerSymbolNumber14                ScatterternaryMarkerSymbol = 14
	ScatterternaryMarkerSymbol14                      ScatterternaryMarkerSymbol = "14"
	ScatterternaryMarkerSymbolHexagon                 ScatterternaryMarkerSymbol = "hexagon"
	ScatterternaryMarkerSymbolNumber114               ScatterternaryMarkerSymbol = 114
	ScatterternaryMarkerSymbol114                     ScatterternaryMarkerSymbol = "114"
	ScatterternaryMarkerSymbolHexagonOpen             ScatterternaryMarkerSymbol = "hexagon-open"
	ScatterternaryMarkerSymbolNumber214               ScatterternaryMarkerSymbol = 214
	ScatterternaryMarkerSymbol214                     ScatterternaryMarkerSymbol = "214"
	ScatterternaryMarkerSymbolHexagonDot              ScatterternaryMarkerSymbol = "hexagon-dot"
	ScatterternaryMarkerSymbolNumber314               ScatterternaryMarkerSymbol = 314
	ScatterternaryMarkerSymbol314                     ScatterternaryMarkerSymbol = "314"
	ScatterternaryMarkerSymbolHexagonOpenDot          ScatterternaryMarkerSymbol = "hexagon-open-dot"
	ScatterternaryMarkerSymbolNumber15                ScatterternaryMarkerSymbol = 15
	ScatterternaryMarkerSymbol15                      ScatterternaryMarkerSymbol = "15"
	ScatterternaryMarkerSymbolHexagon2                ScatterternaryMarkerSymbol = "hexagon2"
	ScatterternaryMarkerSymbolNumber115               ScatterternaryMarkerSymbol = 115
	ScatterternaryMarkerSymbol115                     ScatterternaryMarkerSymbol = "115"
	ScatterternaryMarkerSymbolHexagon2Open            ScatterternaryMarkerSymbol = "hexagon2-open"
	ScatterternaryMarkerSymbolNumber215               ScatterternaryMarkerSymbol = 215
	ScatterternaryMarkerSymbol215                     ScatterternaryMarkerSymbol = "215"
	ScatterternaryMarkerSymbolHexagon2Dot             ScatterternaryMarkerSymbol = "hexagon2-dot"
	ScatterternaryMarkerSymbolNumber315               ScatterternaryMarkerSymbol = 315
	ScatterternaryMarkerSymbol315                     ScatterternaryMarkerSymbol = "315"
	ScatterternaryMarkerSymbolHexagon2OpenDot         ScatterternaryMarkerSymbol = "hexagon2-open-dot"
	ScatterternaryMarkerSymbolNumber16                ScatterternaryMarkerSymbol = 16
	ScatterternaryMarkerSymbol16                      ScatterternaryMarkerSymbol = "16"
	ScatterternaryMarkerSymbolOctagon                 ScatterternaryMarkerSymbol = "octagon"
	ScatterternaryMarkerSymbolNumber116               ScatterternaryMarkerSymbol = 116
	ScatterternaryMarkerSymbol116                     ScatterternaryMarkerSymbol = "116"
	ScatterternaryMarkerSymbolOctagonOpen             ScatterternaryMarkerSymbol = "octagon-open"
	ScatterternaryMarkerSymbolNumber216               ScatterternaryMarkerSymbol = 216
	ScatterternaryMarkerSymbol216                     ScatterternaryMarkerSymbol = "216"
	ScatterternaryMarkerSymbolOctagonDot              ScatterternaryMarkerSymbol = "octagon-dot"
	ScatterternaryMarkerSymbolNumber316               ScatterternaryMarkerSymbol = 316
	ScatterternaryMarkerSymbol316                     ScatterternaryMarkerSymbol = "316"
	ScatterternaryMarkerSymbolOctagonOpenDot          ScatterternaryMarkerSymbol = "octagon-open-dot"
	ScatterternaryMarkerSymbolNumber17                ScatterternaryMarkerSymbol = 17
	ScatterternaryMarkerSymbol17                      ScatterternaryMarkerSymbol = "17"
	ScatterternaryMarkerSymbolStar                    ScatterternaryMarkerSymbol = "star"
	ScatterternaryMarkerSymbolNumber117               ScatterternaryMarkerSymbol = 117
	ScatterternaryMarkerSymbol117                     ScatterternaryMarkerSymbol = "117"
	ScatterternaryMarkerSymbolStarOpen                ScatterternaryMarkerSymbol = "star-open"
	ScatterternaryMarkerSymbolNumber217               ScatterternaryMarkerSymbol = 217
	ScatterternaryMarkerSymbol217                     ScatterternaryMarkerSymbol = "217"
	ScatterternaryMarkerSymbolStarDot                 ScatterternaryMarkerSymbol = "star-dot"
	ScatterternaryMarkerSymbolNumber317               ScatterternaryMarkerSymbol = 317
	ScatterternaryMarkerSymbol317                     ScatterternaryMarkerSymbol = "317"
	ScatterternaryMarkerSymbolStarOpenDot             ScatterternaryMarkerSymbol = "star-open-dot"
	ScatterternaryMarkerSymbolNumber18                ScatterternaryMarkerSymbol = 18
	ScatterternaryMarkerSymbol18                      ScatterternaryMarkerSymbol = "18"
	ScatterternaryMarkerSymbolHexagram                ScatterternaryMarkerSymbol = "hexagram"
	ScatterternaryMarkerSymbolNumber118               ScatterternaryMarkerSymbol = 118
	ScatterternaryMarkerSymbol118                     ScatterternaryMarkerSymbol = "118"
	ScatterternaryMarkerSymbolHexagramOpen            ScatterternaryMarkerSymbol = "hexagram-open"
	ScatterternaryMarkerSymbolNumber218               ScatterternaryMarkerSymbol = 218
	ScatterternaryMarkerSymbol218                     ScatterternaryMarkerSymbol = "218"
	ScatterternaryMarkerSymbolHexagramDot             ScatterternaryMarkerSymbol = "hexagram-dot"
	ScatterternaryMarkerSymbolNumber318               ScatterternaryMarkerSymbol = 318
	ScatterternaryMarkerSymbol318                     ScatterternaryMarkerSymbol = "318"
	ScatterternaryMarkerSymbolHexagramOpenDot         ScatterternaryMarkerSymbol = "hexagram-open-dot"
	ScatterternaryMarkerSymbolNumber19                ScatterternaryMarkerSymbol = 19
	ScatterternaryMarkerSymbol19                      ScatterternaryMarkerSymbol = "19"
	ScatterternaryMarkerSymbolStarTriangleUp          ScatterternaryMarkerSymbol = "star-triangle-up"
	ScatterternaryMarkerSymbolNumber119               ScatterternaryMarkerSymbol = 119
	ScatterternaryMarkerSymbol119                     ScatterternaryMarkerSymbol = "119"
	ScatterternaryMarkerSymbolStarTriangleUpOpen      ScatterternaryMarkerSymbol = "star-triangle-up-open"
	ScatterternaryMarkerSymbolNumber219               ScatterternaryMarkerSymbol = 219
	ScatterternaryMarkerSymbol219                     ScatterternaryMarkerSymbol = "219"
	ScatterternaryMarkerSymbolStarTriangleUpDot       ScatterternaryMarkerSymbol = "star-triangle-up-dot"
	ScatterternaryMarkerSymbolNumber319               ScatterternaryMarkerSymbol = 319
	ScatterternaryMarkerSymbol319                     ScatterternaryMarkerSymbol = "319"
	ScatterternaryMarkerSymbolStarTriangleUpOpenDot   ScatterternaryMarkerSymbol = "star-triangle-up-open-dot"
	ScatterternaryMarkerSymbolNumber20                ScatterternaryMarkerSymbol = 20
	ScatterternaryMarkerSymbol20                      ScatterternaryMarkerSymbol = "20"
	ScatterternaryMarkerSymbolStarTriangleDown        ScatterternaryMarkerSymbol = "star-triangle-down"
	ScatterternaryMarkerSymbolNumber120               ScatterternaryMarkerSymbol = 120
	ScatterternaryMarkerSymbol120                     ScatterternaryMarkerSymbol = "120"
	ScatterternaryMarkerSymbolStarTriangleDownOpen    ScatterternaryMarkerSymbol = "star-triangle-down-open"
	ScatterternaryMarkerSymbolNumber220               ScatterternaryMarkerSymbol = 220
	ScatterternaryMarkerSymbol220                     ScatterternaryMarkerSymbol = "220"
	ScatterternaryMarkerSymbolStarTriangleDownDot     ScatterternaryMarkerSymbol = "star-triangle-down-dot"
	ScatterternaryMarkerSymbolNumber320               ScatterternaryMarkerSymbol = 320
	ScatterternaryMarkerSymbol320                     ScatterternaryMarkerSymbol = "320"
	ScatterternaryMarkerSymbolStarTriangleDownOpenDot ScatterternaryMarkerSymbol = "star-triangle-down-open-dot"
	ScatterternaryMarkerSymbolNumber21                ScatterternaryMarkerSymbol = 21
	ScatterternaryMarkerSymbol21                      ScatterternaryMarkerSymbol = "21"
	ScatterternaryMarkerSymbolStarSquare              ScatterternaryMarkerSymbol = "star-square"
	ScatterternaryMarkerSymbolNumber121               ScatterternaryMarkerSymbol = 121
	ScatterternaryMarkerSymbol121                     ScatterternaryMarkerSymbol = "121"
	ScatterternaryMarkerSymbolStarSquareOpen          ScatterternaryMarkerSymbol = "star-square-open"
	ScatterternaryMarkerSymbolNumber221               ScatterternaryMarkerSymbol = 221
	ScatterternaryMarkerSymbol221                     ScatterternaryMarkerSymbol = "221"
	ScatterternaryMarkerSymbolStarSquareDot           ScatterternaryMarkerSymbol = "star-square-dot"
	ScatterternaryMarkerSymbolNumber321               ScatterternaryMarkerSymbol = 321
	ScatterternaryMarkerSymbol321                     ScatterternaryMarkerSymbol = "321"
	ScatterternaryMarkerSymbolStarSquareOpenDot       ScatterternaryMarkerSymbol = "star-square-open-dot"
	ScatterternaryMarkerSymbolNumber22                ScatterternaryMarkerSymbol = 22
	ScatterternaryMarkerSymbol22                      ScatterternaryMarkerSymbol = "22"
	ScatterternaryMarkerSymbolStarDiamond             ScatterternaryMarkerSymbol = "star-diamond"
	ScatterternaryMarkerSymbolNumber122               ScatterternaryMarkerSymbol = 122
	ScatterternaryMarkerSymbol122                     ScatterternaryMarkerSymbol = "122"
	ScatterternaryMarkerSymbolStarDiamondOpen         ScatterternaryMarkerSymbol = "star-diamond-open"
	ScatterternaryMarkerSymbolNumber222               ScatterternaryMarkerSymbol = 222
	ScatterternaryMarkerSymbol222                     ScatterternaryMarkerSymbol = "222"
	ScatterternaryMarkerSymbolStarDiamondDot          ScatterternaryMarkerSymbol = "star-diamond-dot"
	ScatterternaryMarkerSymbolNumber322               ScatterternaryMarkerSymbol = 322
	ScatterternaryMarkerSymbol322                     ScatterternaryMarkerSymbol = "322"
	ScatterternaryMarkerSymbolStarDiamondOpenDot      ScatterternaryMarkerSymbol = "star-diamond-open-dot"
	ScatterternaryMarkerSymbolNumber23                ScatterternaryMarkerSymbol = 23
	ScatterternaryMarkerSymbol23                      ScatterternaryMarkerSymbol = "23"
	ScatterternaryMarkerSymbolDiamondTall             ScatterternaryMarkerSymbol = "diamond-tall"
	ScatterternaryMarkerSymbolNumber123               ScatterternaryMarkerSymbol = 123
	ScatterternaryMarkerSymbol123                     ScatterternaryMarkerSymbol = "123"
	ScatterternaryMarkerSymbolDiamondTallOpen         ScatterternaryMarkerSymbol = "diamond-tall-open"
	ScatterternaryMarkerSymbolNumber223               ScatterternaryMarkerSymbol = 223
	ScatterternaryMarkerSymbol223                     ScatterternaryMarkerSymbol = "223"
	ScatterternaryMarkerSymbolDiamondTallDot          ScatterternaryMarkerSymbol = "diamond-tall-dot"
	ScatterternaryMarkerSymbolNumber323               ScatterternaryMarkerSymbol = 323
	ScatterternaryMarkerSymbol323                     ScatterternaryMarkerSymbol = "323"
	ScatterternaryMarkerSymbolDiamondTallOpenDot      ScatterternaryMarkerSymbol = "diamond-tall-open-dot"
	ScatterternaryMarkerSymbolNumber24                ScatterternaryMarkerSymbol = 24
	ScatterternaryMarkerSymbol24                      ScatterternaryMarkerSymbol = "24"
	ScatterternaryMarkerSymbolDiamondWide             ScatterternaryMarkerSymbol = "diamond-wide"
	ScatterternaryMarkerSymbolNumber124               ScatterternaryMarkerSymbol = 124
	ScatterternaryMarkerSymbol124                     ScatterternaryMarkerSymbol = "124"
	ScatterternaryMarkerSymbolDiamondWideOpen         ScatterternaryMarkerSymbol = "diamond-wide-open"
	ScatterternaryMarkerSymbolNumber224               ScatterternaryMarkerSymbol = 224
	ScatterternaryMarkerSymbol224                     ScatterternaryMarkerSymbol = "224"
	ScatterternaryMarkerSymbolDiamondWideDot          ScatterternaryMarkerSymbol = "diamond-wide-dot"
	ScatterternaryMarkerSymbolNumber324               ScatterternaryMarkerSymbol = 324
	ScatterternaryMarkerSymbol324                     ScatterternaryMarkerSymbol = "324"
	ScatterternaryMarkerSymbolDiamondWideOpenDot      ScatterternaryMarkerSymbol = "diamond-wide-open-dot"
	ScatterternaryMarkerSymbolNumber25                ScatterternaryMarkerSymbol = 25
	ScatterternaryMarkerSymbol25                      ScatterternaryMarkerSymbol = "25"
	ScatterternaryMarkerSymbolHourglass               ScatterternaryMarkerSymbol = "hourglass"
	ScatterternaryMarkerSymbolNumber125               ScatterternaryMarkerSymbol = 125
	ScatterternaryMarkerSymbol125                     ScatterternaryMarkerSymbol = "125"
	ScatterternaryMarkerSymbolHourglassOpen           ScatterternaryMarkerSymbol = "hourglass-open"
	ScatterternaryMarkerSymbolNumber26                ScatterternaryMarkerSymbol = 26
	ScatterternaryMarkerSymbol26                      ScatterternaryMarkerSymbol = "26"
	ScatterternaryMarkerSymbolBowtie                  ScatterternaryMarkerSymbol = "bowtie"
	ScatterternaryMarkerSymbolNumber126               ScatterternaryMarkerSymbol = 126
	ScatterternaryMarkerSymbol126                     ScatterternaryMarkerSymbol = "126"
	ScatterternaryMarkerSymbolBowtieOpen              ScatterternaryMarkerSymbol = "bowtie-open"
	ScatterternaryMarkerSymbolNumber27                ScatterternaryMarkerSymbol = 27
	ScatterternaryMarkerSymbol27                      ScatterternaryMarkerSymbol = "27"
	ScatterternaryMarkerSymbolCircleCross             ScatterternaryMarkerSymbol = "circle-cross"
	ScatterternaryMarkerSymbolNumber127               ScatterternaryMarkerSymbol = 127
	ScatterternaryMarkerSymbol127                     ScatterternaryMarkerSymbol = "127"
	ScatterternaryMarkerSymbolCircleCrossOpen         ScatterternaryMarkerSymbol = "circle-cross-open"
	ScatterternaryMarkerSymbolNumber28                ScatterternaryMarkerSymbol = 28
	ScatterternaryMarkerSymbol28                      ScatterternaryMarkerSymbol = "28"
	ScatterternaryMarkerSymbolCircleX                 ScatterternaryMarkerSymbol = "circle-x"
	ScatterternaryMarkerSymbolNumber128               ScatterternaryMarkerSymbol = 128
	ScatterternaryMarkerSymbol128                     ScatterternaryMarkerSymbol = "128"
	ScatterternaryMarkerSymbolCircleXOpen             ScatterternaryMarkerSymbol = "circle-x-open"
	ScatterternaryMarkerSymbolNumber29                ScatterternaryMarkerSymbol = 29
	ScatterternaryMarkerSymbol29                      ScatterternaryMarkerSymbol = "29"
	ScatterternaryMarkerSymbolSquareCross             ScatterternaryMarkerSymbol = "square-cross"
	ScatterternaryMarkerSymbolNumber129               ScatterternaryMarkerSymbol = 129
	ScatterternaryMarkerSymbol129                     ScatterternaryMarkerSymbol = "129"
	ScatterternaryMarkerSymbolSquareCrossOpen         ScatterternaryMarkerSymbol = "square-cross-open"
	ScatterternaryMarkerSymbolNumber30                ScatterternaryMarkerSymbol = 30
	ScatterternaryMarkerSymbol30                      ScatterternaryMarkerSymbol = "30"
	ScatterternaryMarkerSymbolSquareX                 ScatterternaryMarkerSymbol = "square-x"
	ScatterternaryMarkerSymbolNumber130               ScatterternaryMarkerSymbol = 130
	ScatterternaryMarkerSymbol130                     ScatterternaryMarkerSymbol = "130"
	ScatterternaryMarkerSymbolSquareXOpen             ScatterternaryMarkerSymbol = "square-x-open"
	ScatterternaryMarkerSymbolNumber31                ScatterternaryMarkerSymbol = 31
	ScatterternaryMarkerSymbol31                      ScatterternaryMarkerSymbol = "31"
	ScatterternaryMarkerSymbolDiamondCross            ScatterternaryMarkerSymbol = "diamond-cross"
	ScatterternaryMarkerSymbolNumber131               ScatterternaryMarkerSymbol = 131
	ScatterternaryMarkerSymbol131                     ScatterternaryMarkerSymbol = "131"
	ScatterternaryMarkerSymbolDiamondCrossOpen        ScatterternaryMarkerSymbol = "diamond-cross-open"
	ScatterternaryMarkerSymbolNumber32                ScatterternaryMarkerSymbol = 32
	ScatterternaryMarkerSymbol32                      ScatterternaryMarkerSymbol = "32"
	ScatterternaryMarkerSymbolDiamondX                ScatterternaryMarkerSymbol = "diamond-x"
	ScatterternaryMarkerSymbolNumber132               ScatterternaryMarkerSymbol = 132
	ScatterternaryMarkerSymbol132                     ScatterternaryMarkerSymbol = "132"
	ScatterternaryMarkerSymbolDiamondXOpen            ScatterternaryMarkerSymbol = "diamond-x-open"
	ScatterternaryMarkerSymbolNumber33                ScatterternaryMarkerSymbol = 33
	ScatterternaryMarkerSymbol33                      ScatterternaryMarkerSymbol = "33"
	ScatterternaryMarkerSymbolCrossThin               ScatterternaryMarkerSymbol = "cross-thin"
	ScatterternaryMarkerSymbolNumber133               ScatterternaryMarkerSymbol = 133
	ScatterternaryMarkerSymbol133                     ScatterternaryMarkerSymbol = "133"
	ScatterternaryMarkerSymbolCrossThinOpen           ScatterternaryMarkerSymbol = "cross-thin-open"
	ScatterternaryMarkerSymbolNumber34                ScatterternaryMarkerSymbol = 34
	ScatterternaryMarkerSymbol34                      ScatterternaryMarkerSymbol = "34"
	ScatterternaryMarkerSymbolXThin                   ScatterternaryMarkerSymbol = "x-thin"
	ScatterternaryMarkerSymbolNumber134               ScatterternaryMarkerSymbol = 134
	ScatterternaryMarkerSymbol134                     ScatterternaryMarkerSymbol = "134"
	ScatterternaryMarkerSymbolXThinOpen               ScatterternaryMarkerSymbol = "x-thin-open"
	ScatterternaryMarkerSymbolNumber35                ScatterternaryMarkerSymbol = 35
	ScatterternaryMarkerSymbol35                      ScatterternaryMarkerSymbol = "35"
	ScatterternaryMarkerSymbolAsterisk                ScatterternaryMarkerSymbol = "asterisk"
	ScatterternaryMarkerSymbolNumber135               ScatterternaryMarkerSymbol = 135
	ScatterternaryMarkerSymbol135                     ScatterternaryMarkerSymbol = "135"
	ScatterternaryMarkerSymbolAsteriskOpen            ScatterternaryMarkerSymbol = "asterisk-open"
	ScatterternaryMarkerSymbolNumber36                ScatterternaryMarkerSymbol = 36
	ScatterternaryMarkerSymbol36                      ScatterternaryMarkerSymbol = "36"
	ScatterternaryMarkerSymbolHash                    ScatterternaryMarkerSymbol = "hash"
	ScatterternaryMarkerSymbolNumber136               ScatterternaryMarkerSymbol = 136
	ScatterternaryMarkerSymbol136                     ScatterternaryMarkerSymbol = "136"
	ScatterternaryMarkerSymbolHashOpen                ScatterternaryMarkerSymbol = "hash-open"
	ScatterternaryMarkerSymbolNumber236               ScatterternaryMarkerSymbol = 236
	ScatterternaryMarkerSymbol236                     ScatterternaryMarkerSymbol = "236"
	ScatterternaryMarkerSymbolHashDot                 ScatterternaryMarkerSymbol = "hash-dot"
	ScatterternaryMarkerSymbolNumber336               ScatterternaryMarkerSymbol = 336
	ScatterternaryMarkerSymbol336                     ScatterternaryMarkerSymbol = "336"
	ScatterternaryMarkerSymbolHashOpenDot             ScatterternaryMarkerSymbol = "hash-open-dot"
	ScatterternaryMarkerSymbolNumber37                ScatterternaryMarkerSymbol = 37
	ScatterternaryMarkerSymbol37                      ScatterternaryMarkerSymbol = "37"
	ScatterternaryMarkerSymbolYUp                     ScatterternaryMarkerSymbol = "y-up"
	ScatterternaryMarkerSymbolNumber137               ScatterternaryMarkerSymbol = 137
	ScatterternaryMarkerSymbol137                     ScatterternaryMarkerSymbol = "137"
	ScatterternaryMarkerSymbolYUpOpen                 ScatterternaryMarkerSymbol = "y-up-open"
	ScatterternaryMarkerSymbolNumber38                ScatterternaryMarkerSymbol = 38
	ScatterternaryMarkerSymbol38                      ScatterternaryMarkerSymbol = "38"
	ScatterternaryMarkerSymbolYDown                   ScatterternaryMarkerSymbol = "y-down"
	ScatterternaryMarkerSymbolNumber138               ScatterternaryMarkerSymbol = 138
	ScatterternaryMarkerSymbol138                     ScatterternaryMarkerSymbol = "138"
	ScatterternaryMarkerSymbolYDownOpen               ScatterternaryMarkerSymbol = "y-down-open"
	ScatterternaryMarkerSymbolNumber39                ScatterternaryMarkerSymbol = 39
	ScatterternaryMarkerSymbol39                      ScatterternaryMarkerSymbol = "39"
	ScatterternaryMarkerSymbolYLeft                   ScatterternaryMarkerSymbol = "y-left"
	ScatterternaryMarkerSymbolNumber139               ScatterternaryMarkerSymbol = 139
	ScatterternaryMarkerSymbol139                     ScatterternaryMarkerSymbol = "139"
	ScatterternaryMarkerSymbolYLeftOpen               ScatterternaryMarkerSymbol = "y-left-open"
	ScatterternaryMarkerSymbolNumber40                ScatterternaryMarkerSymbol = 40
	ScatterternaryMarkerSymbol40                      ScatterternaryMarkerSymbol = "40"
	ScatterternaryMarkerSymbolYRight                  ScatterternaryMarkerSymbol = "y-right"
	ScatterternaryMarkerSymbolNumber140               ScatterternaryMarkerSymbol = 140
	ScatterternaryMarkerSymbol140                     ScatterternaryMarkerSymbol = "140"
	ScatterternaryMarkerSymbolYRightOpen              ScatterternaryMarkerSymbol = "y-right-open"
	ScatterternaryMarkerSymbolNumber41                ScatterternaryMarkerSymbol = 41
	ScatterternaryMarkerSymbol41                      ScatterternaryMarkerSymbol = "41"
	ScatterternaryMarkerSymbolLineEw                  ScatterternaryMarkerSymbol = "line-ew"
	ScatterternaryMarkerSymbolNumber141               ScatterternaryMarkerSymbol = 141
	ScatterternaryMarkerSymbol141                     ScatterternaryMarkerSymbol = "141"
	ScatterternaryMarkerSymbolLineEwOpen              ScatterternaryMarkerSymbol = "line-ew-open"
	ScatterternaryMarkerSymbolNumber42                ScatterternaryMarkerSymbol = 42
	ScatterternaryMarkerSymbol42                      ScatterternaryMarkerSymbol = "42"
	ScatterternaryMarkerSymbolLineNs                  ScatterternaryMarkerSymbol = "line-ns"
	ScatterternaryMarkerSymbolNumber142               ScatterternaryMarkerSymbol = 142
	ScatterternaryMarkerSymbol142                     ScatterternaryMarkerSymbol = "142"
	ScatterternaryMarkerSymbolLineNsOpen              ScatterternaryMarkerSymbol = "line-ns-open"
	ScatterternaryMarkerSymbolNumber43                ScatterternaryMarkerSymbol = 43
	ScatterternaryMarkerSymbol43                      ScatterternaryMarkerSymbol = "43"
	ScatterternaryMarkerSymbolLineNe                  ScatterternaryMarkerSymbol = "line-ne"
	ScatterternaryMarkerSymbolNumber143               ScatterternaryMarkerSymbol = 143
	ScatterternaryMarkerSymbol143                     ScatterternaryMarkerSymbol = "143"
	ScatterternaryMarkerSymbolLineNeOpen              ScatterternaryMarkerSymbol = "line-ne-open"
	ScatterternaryMarkerSymbolNumber44                ScatterternaryMarkerSymbol = 44
	ScatterternaryMarkerSymbol44                      ScatterternaryMarkerSymbol = "44"
	ScatterternaryMarkerSymbolLineNw                  ScatterternaryMarkerSymbol = "line-nw"
	ScatterternaryMarkerSymbolNumber144               ScatterternaryMarkerSymbol = 144
	ScatterternaryMarkerSymbol144                     ScatterternaryMarkerSymbol = "144"
	ScatterternaryMarkerSymbolLineNwOpen              ScatterternaryMarkerSymbol = "line-nw-open"
	ScatterternaryMarkerSymbolNumber45                ScatterternaryMarkerSymbol = 45
	ScatterternaryMarkerSymbol45                      ScatterternaryMarkerSymbol = "45"
	ScatterternaryMarkerSymbolArrowUp                 ScatterternaryMarkerSymbol = "arrow-up"
	ScatterternaryMarkerSymbolNumber145               ScatterternaryMarkerSymbol = 145
	ScatterternaryMarkerSymbol145                     ScatterternaryMarkerSymbol = "145"
	ScatterternaryMarkerSymbolArrowUpOpen             ScatterternaryMarkerSymbol = "arrow-up-open"
	ScatterternaryMarkerSymbolNumber46                ScatterternaryMarkerSymbol = 46
	ScatterternaryMarkerSymbol46                      ScatterternaryMarkerSymbol = "46"
	ScatterternaryMarkerSymbolArrowDown               ScatterternaryMarkerSymbol = "arrow-down"
	ScatterternaryMarkerSymbolNumber146               ScatterternaryMarkerSymbol = 146
	ScatterternaryMarkerSymbol146                     ScatterternaryMarkerSymbol = "146"
	ScatterternaryMarkerSymbolArrowDownOpen           ScatterternaryMarkerSymbol = "arrow-down-open"
	ScatterternaryMarkerSymbolNumber47                ScatterternaryMarkerSymbol = 47
	ScatterternaryMarkerSymbol47                      ScatterternaryMarkerSymbol = "47"
	ScatterternaryMarkerSymbolArrowLeft               ScatterternaryMarkerSymbol = "arrow-left"
	ScatterternaryMarkerSymbolNumber147               ScatterternaryMarkerSymbol = 147
	ScatterternaryMarkerSymbol147                     ScatterternaryMarkerSymbol = "147"
	ScatterternaryMarkerSymbolArrowLeftOpen           ScatterternaryMarkerSymbol = "arrow-left-open"
	ScatterternaryMarkerSymbolNumber48                ScatterternaryMarkerSymbol = 48
	ScatterternaryMarkerSymbol48                      ScatterternaryMarkerSymbol = "48"
	ScatterternaryMarkerSymbolArrowRight              ScatterternaryMarkerSymbol = "arrow-right"
	ScatterternaryMarkerSymbolNumber148               ScatterternaryMarkerSymbol = 148
	ScatterternaryMarkerSymbol148                     ScatterternaryMarkerSymbol = "148"
	ScatterternaryMarkerSymbolArrowRightOpen          ScatterternaryMarkerSymbol = "arrow-right-open"
	ScatterternaryMarkerSymbolNumber49                ScatterternaryMarkerSymbol = 49
	ScatterternaryMarkerSymbol49                      ScatterternaryMarkerSymbol = "49"
	ScatterternaryMarkerSymbolArrowBarUp              ScatterternaryMarkerSymbol = "arrow-bar-up"
	ScatterternaryMarkerSymbolNumber149               ScatterternaryMarkerSymbol = 149
	ScatterternaryMarkerSymbol149                     ScatterternaryMarkerSymbol = "149"
	ScatterternaryMarkerSymbolArrowBarUpOpen          ScatterternaryMarkerSymbol = "arrow-bar-up-open"
	ScatterternaryMarkerSymbolNumber50                ScatterternaryMarkerSymbol = 50
	ScatterternaryMarkerSymbol50                      ScatterternaryMarkerSymbol = "50"
	ScatterternaryMarkerSymbolArrowBarDown            ScatterternaryMarkerSymbol = "arrow-bar-down"
	ScatterternaryMarkerSymbolNumber150               ScatterternaryMarkerSymbol = 150
	ScatterternaryMarkerSymbol150                     ScatterternaryMarkerSymbol = "150"
	ScatterternaryMarkerSymbolArrowBarDownOpen        ScatterternaryMarkerSymbol = "arrow-bar-down-open"
	ScatterternaryMarkerSymbolNumber51                ScatterternaryMarkerSymbol = 51
	ScatterternaryMarkerSymbol51                      ScatterternaryMarkerSymbol = "51"
	ScatterternaryMarkerSymbolArrowBarLeft            ScatterternaryMarkerSymbol = "arrow-bar-left"
	ScatterternaryMarkerSymbolNumber151               ScatterternaryMarkerSymbol = 151
	ScatterternaryMarkerSymbol151                     ScatterternaryMarkerSymbol = "151"
	ScatterternaryMarkerSymbolArrowBarLeftOpen        ScatterternaryMarkerSymbol = "arrow-bar-left-open"
	ScatterternaryMarkerSymbolNumber52                ScatterternaryMarkerSymbol = 52
	ScatterternaryMarkerSymbol52                      ScatterternaryMarkerSymbol = "52"
	ScatterternaryMarkerSymbolArrowBarRight           ScatterternaryMarkerSymbol = "arrow-bar-right"
	ScatterternaryMarkerSymbolNumber152               ScatterternaryMarkerSymbol = 152
	ScatterternaryMarkerSymbol152                     ScatterternaryMarkerSymbol = "152"
	ScatterternaryMarkerSymbolArrowBarRightOpen       ScatterternaryMarkerSymbol = "arrow-bar-right-open"
)

// ScatterternaryTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterternaryTextposition string

const (
	ScatterternaryTextpositionTopLeft      ScatterternaryTextposition = "top left"
	ScatterternaryTextpositionTopCenter    ScatterternaryTextposition = "top center"
	ScatterternaryTextpositionTopRight     ScatterternaryTextposition = "top right"
	ScatterternaryTextpositionMiddleLeft   ScatterternaryTextposition = "middle left"
	ScatterternaryTextpositionMiddleCenter ScatterternaryTextposition = "middle center"
	ScatterternaryTextpositionMiddleRight  ScatterternaryTextposition = "middle right"
	ScatterternaryTextpositionBottomLeft   ScatterternaryTextposition = "bottom left"
	ScatterternaryTextpositionBottomCenter ScatterternaryTextposition = "bottom center"
	ScatterternaryTextpositionBottomRight  ScatterternaryTextposition = "bottom right"
)

// ScatterternaryVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterternaryVisible interface{}

var (
	ScatterternaryVisibleTrue       ScatterternaryVisible = true
	ScatterternaryVisibleFalse      ScatterternaryVisible = false
	ScatterternaryVisibleLegendonly ScatterternaryVisible = "legendonly"
)

// ScatterternaryHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type ScatterternaryHoverinfo string

const (
	// Flags
	ScatterternaryHoverinfoA    ScatterternaryHoverinfo = "a"
	ScatterternaryHoverinfoB    ScatterternaryHoverinfo = "b"
	ScatterternaryHoverinfoC    ScatterternaryHoverinfo = "c"
	ScatterternaryHoverinfoText ScatterternaryHoverinfo = "text"
	ScatterternaryHoverinfoName ScatterternaryHoverinfo = "name"

	// Extra
	ScatterternaryHoverinfoAll  ScatterternaryHoverinfo = "all"
	ScatterternaryHoverinfoNone ScatterternaryHoverinfo = "none"
	ScatterternaryHoverinfoSkip ScatterternaryHoverinfo = "skip"
)

// ScatterternaryHoveron Do the hover effects highlight individual points (markers or line points) or do they highlight filled regions? If the fill is *toself* or *tonext* and there are no markers or text, then the default is *fills*, otherwise it is *points*.
type ScatterternaryHoveron string

const (
	// Flags
	ScatterternaryHoveronPoints ScatterternaryHoveron = "points"
	ScatterternaryHoveronFills  ScatterternaryHoveron = "fills"

	// Extra

)

// ScatterternaryMode Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
type ScatterternaryMode string

const (
	// Flags
	ScatterternaryModeLines   ScatterternaryMode = "lines"
	ScatterternaryModeMarkers ScatterternaryMode = "markers"
	ScatterternaryModeText    ScatterternaryMode = "text"

	// Extra
	ScatterternaryModeNone ScatterternaryMode = "none"
)
