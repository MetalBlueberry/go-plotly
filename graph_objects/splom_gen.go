package grob

var TraceTypeSplom TraceType = "splom"

func (trace *Splom) GetType() TraceType {
	return TraceTypeSplom
}

// Splom Splom traces generate scatter plot matrix visualizations. Each splom `dimensions` items correspond to a generated axis. Values for each of those dimensions are set in `dimensions[i].values`. Splom traces support all `scattergl` marker style attributes. Specify `layout.grid` attributes and/or layout x-axis and y-axis attributes for more control over the axis positioning and style.
type Splom struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

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

	// Diagonal
	// role: Object
	Diagonal *SplomDiagonal `json:"diagonal,omitempty"`

	// Dimensions
	// It's an items array and what goes inside it's... messy... check the docs
	// I will be happy if you want to contribute by implementing this
	// just raise an issue before you start so we do not overlap
	Dimensions interface{} `json:"dimensions,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo SplomHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *SplomHoverlabel `json:"hoverlabel,omitempty"`

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
	// Same as `text`.
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

	// Marker
	// role: Object
	Marker *SplomMarker `json:"marker,omitempty"`

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
	Selected *SplomSelected `json:"selected,omitempty"`

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

	// Showlowerhalf
	// arrayOK: false
	// type: boolean
	// Determines whether or not subplots on the lower half from the diagonal are displayed.
	Showlowerhalf Bool `json:"showlowerhalf,omitempty"`

	// Showupperhalf
	// arrayOK: false
	// type: boolean
	// Determines whether or not subplots on the upper half from the diagonal are displayed.
	Showupperhalf Bool `json:"showupperhalf,omitempty"`

	// Stream
	// role: Object
	Stream *SplomStream `json:"stream,omitempty"`

	// Text
	// arrayOK: true
	// type: string
	// Sets text elements associated with each (x,y) pair to appear on hover. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates.
	Text String `json:"text,omitempty"`

	// Textsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

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
	Unselected *SplomUnselected `json:"unselected,omitempty"`

	// Visible
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible SplomVisible `json:"visible,omitempty"`

	// Xaxes
	// arrayOK: false
	// type: info_array
	// Sets the list of x axes corresponding to dimensions of this splom trace. By default, a splom will match the first N xaxes where N is the number of input dimensions. Note that, in case where `diagonal.visible` is false and `showupperhalf` or `showlowerhalf` is false, this splom trace will generate one less x-axis and one less y-axis.
	Xaxes interface{} `json:"xaxes,omitempty"`

	// Yaxes
	// arrayOK: false
	// type: info_array
	// Sets the list of y axes corresponding to dimensions of this splom trace. By default, a splom will match the first N yaxes where N is the number of input dimensions. Note that, in case where `diagonal.visible` is false and `showupperhalf` or `showlowerhalf` is false, this splom trace will generate one less x-axis and one less y-axis.
	Yaxes interface{} `json:"yaxes,omitempty"`
}

// SplomDiagonal
type SplomDiagonal struct {

	// Visible
	// arrayOK: false
	// type: boolean
	// Determines whether or not subplots on the diagonal are displayed.
	Visible Bool `json:"visible,omitempty"`
}

// SplomHoverlabelFont Sets the font used in hover labels.
type SplomHoverlabelFont struct {

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

// SplomHoverlabel
type SplomHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align SplomHoverlabelAlign `json:"align,omitempty"`

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
	Font *SplomHoverlabelFont `json:"font,omitempty"`

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

// SplomMarkerColorbarTickfont Sets the color bar's tick label font
type SplomMarkerColorbarTickfont struct {

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

// SplomMarkerColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type SplomMarkerColorbarTitleFont struct {

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

// SplomMarkerColorbarTitle
type SplomMarkerColorbarTitle struct {

	// Font
	// role: Object
	Font *SplomMarkerColorbarTitleFont `json:"font,omitempty"`

	// Side
	// default: top
	// type: enumerated
	// Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
	Side SplomMarkerColorbarTitleSide `json:"side,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// SplomMarkerColorbar
type SplomMarkerColorbar struct {

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
	Exponentformat SplomMarkerColorbarExponentformat `json:"exponentformat,omitempty"`

	// Len
	// arrayOK: false
	// type: number
	// Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends.
	Len float64 `json:"len,omitempty"`

	// Lenmode
	// default: fraction
	// type: enumerated
	// Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
	Lenmode SplomMarkerColorbarLenmode `json:"lenmode,omitempty"`

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
	Showexponent SplomMarkerColorbarShowexponent `json:"showexponent,omitempty"`

	// Showticklabels
	// arrayOK: false
	// type: boolean
	// Determines whether or not the tick labels are drawn.
	Showticklabels Bool `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix SplomMarkerColorbarShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix SplomMarkerColorbarShowticksuffix `json:"showticksuffix,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels.
	Thickness float64 `json:"thickness,omitempty"`

	// Thicknessmode
	// default: pixels
	// type: enumerated
	// Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
	Thicknessmode SplomMarkerColorbarThicknessmode `json:"thicknessmode,omitempty"`

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
	Tickfont *SplomMarkerColorbarTickfont `json:"tickfont,omitempty"`

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
	Ticklabelposition SplomMarkerColorbarTicklabelposition `json:"ticklabelposition,omitempty"`

	// Ticklen
	// arrayOK: false
	// type: number
	// Sets the tick length (in px).
	Ticklen float64 `json:"ticklen,omitempty"`

	// Tickmode
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
	Tickmode SplomMarkerColorbarTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

	// Ticks
	// default:
	// type: enumerated
	// Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
	Ticks SplomMarkerColorbarTicks `json:"ticks,omitempty"`

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
	Title *SplomMarkerColorbarTitle `json:"title,omitempty"`

	// X
	// arrayOK: false
	// type: number
	// Sets the x position of the color bar (in plot fraction).
	X float64 `json:"x,omitempty"`

	// Xanchor
	// default: left
	// type: enumerated
	// Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
	Xanchor SplomMarkerColorbarXanchor `json:"xanchor,omitempty"`

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
	Yanchor SplomMarkerColorbarYanchor `json:"yanchor,omitempty"`

	// Ypad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the y direction.
	Ypad float64 `json:"ypad,omitempty"`
}

// SplomMarkerLine
type SplomMarkerLine struct {

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

// SplomMarker
type SplomMarker struct {

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
	Colorbar *SplomMarkerColorbar `json:"colorbar,omitempty"`

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
	Line *SplomMarkerLine `json:"line,omitempty"`

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
	Sizemode SplomMarkerSizemode `json:"sizemode,omitempty"`

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
	Symbol SplomMarkerSymbol `json:"symbol,omitempty"`

	// Symbolsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  symbol .
	Symbolsrc String `json:"symbolsrc,omitempty"`
}

// SplomSelectedMarker
type SplomSelectedMarker struct {

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

// SplomSelected
type SplomSelected struct {

	// Marker
	// role: Object
	Marker *SplomSelectedMarker `json:"marker,omitempty"`
}

// SplomStream
type SplomStream struct {

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

// SplomUnselectedMarker
type SplomUnselectedMarker struct {

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

// SplomUnselected
type SplomUnselected struct {

	// Marker
	// role: Object
	Marker *SplomUnselectedMarker `json:"marker,omitempty"`
}

// SplomHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SplomHoverlabelAlign string

const (
	SplomHoverlabelAlignLeft  SplomHoverlabelAlign = "left"
	SplomHoverlabelAlignRight SplomHoverlabelAlign = "right"
	SplomHoverlabelAlignAuto  SplomHoverlabelAlign = "auto"
)

// SplomMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type SplomMarkerColorbarExponentformat string

const (
	SplomMarkerColorbarExponentformatNone  SplomMarkerColorbarExponentformat = "none"
	SplomMarkerColorbarExponentformatE1    SplomMarkerColorbarExponentformat = "e"
	SplomMarkerColorbarExponentformatE2    SplomMarkerColorbarExponentformat = "E"
	SplomMarkerColorbarExponentformatPower SplomMarkerColorbarExponentformat = "power"
	SplomMarkerColorbarExponentformatSi    SplomMarkerColorbarExponentformat = "SI"
	SplomMarkerColorbarExponentformatB     SplomMarkerColorbarExponentformat = "B"
)

// SplomMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type SplomMarkerColorbarLenmode string

const (
	SplomMarkerColorbarLenmodeFraction SplomMarkerColorbarLenmode = "fraction"
	SplomMarkerColorbarLenmodePixels   SplomMarkerColorbarLenmode = "pixels"
)

// SplomMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type SplomMarkerColorbarShowexponent string

const (
	SplomMarkerColorbarShowexponentAll   SplomMarkerColorbarShowexponent = "all"
	SplomMarkerColorbarShowexponentFirst SplomMarkerColorbarShowexponent = "first"
	SplomMarkerColorbarShowexponentLast  SplomMarkerColorbarShowexponent = "last"
	SplomMarkerColorbarShowexponentNone  SplomMarkerColorbarShowexponent = "none"
)

// SplomMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type SplomMarkerColorbarShowtickprefix string

const (
	SplomMarkerColorbarShowtickprefixAll   SplomMarkerColorbarShowtickprefix = "all"
	SplomMarkerColorbarShowtickprefixFirst SplomMarkerColorbarShowtickprefix = "first"
	SplomMarkerColorbarShowtickprefixLast  SplomMarkerColorbarShowtickprefix = "last"
	SplomMarkerColorbarShowtickprefixNone  SplomMarkerColorbarShowtickprefix = "none"
)

// SplomMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type SplomMarkerColorbarShowticksuffix string

const (
	SplomMarkerColorbarShowticksuffixAll   SplomMarkerColorbarShowticksuffix = "all"
	SplomMarkerColorbarShowticksuffixFirst SplomMarkerColorbarShowticksuffix = "first"
	SplomMarkerColorbarShowticksuffixLast  SplomMarkerColorbarShowticksuffix = "last"
	SplomMarkerColorbarShowticksuffixNone  SplomMarkerColorbarShowticksuffix = "none"
)

// SplomMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type SplomMarkerColorbarThicknessmode string

const (
	SplomMarkerColorbarThicknessmodeFraction SplomMarkerColorbarThicknessmode = "fraction"
	SplomMarkerColorbarThicknessmodePixels   SplomMarkerColorbarThicknessmode = "pixels"
)

// SplomMarkerColorbarTicklabelposition Determines where tick labels are drawn.
type SplomMarkerColorbarTicklabelposition string

const (
	SplomMarkerColorbarTicklabelpositionOutside       SplomMarkerColorbarTicklabelposition = "outside"
	SplomMarkerColorbarTicklabelpositionInside        SplomMarkerColorbarTicklabelposition = "inside"
	SplomMarkerColorbarTicklabelpositionOutsideTop    SplomMarkerColorbarTicklabelposition = "outside top"
	SplomMarkerColorbarTicklabelpositionInsideTop     SplomMarkerColorbarTicklabelposition = "inside top"
	SplomMarkerColorbarTicklabelpositionOutsideBottom SplomMarkerColorbarTicklabelposition = "outside bottom"
	SplomMarkerColorbarTicklabelpositionInsideBottom  SplomMarkerColorbarTicklabelposition = "inside bottom"
)

// SplomMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type SplomMarkerColorbarTickmode string

const (
	SplomMarkerColorbarTickmodeAuto   SplomMarkerColorbarTickmode = "auto"
	SplomMarkerColorbarTickmodeLinear SplomMarkerColorbarTickmode = "linear"
	SplomMarkerColorbarTickmodeArray  SplomMarkerColorbarTickmode = "array"
)

// SplomMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type SplomMarkerColorbarTicks string

const (
	SplomMarkerColorbarTicksOutside SplomMarkerColorbarTicks = "outside"
	SplomMarkerColorbarTicksInside  SplomMarkerColorbarTicks = "inside"
	SplomMarkerColorbarTicksEmpty   SplomMarkerColorbarTicks = ""
)

// SplomMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type SplomMarkerColorbarTitleSide string

const (
	SplomMarkerColorbarTitleSideRight  SplomMarkerColorbarTitleSide = "right"
	SplomMarkerColorbarTitleSideTop    SplomMarkerColorbarTitleSide = "top"
	SplomMarkerColorbarTitleSideBottom SplomMarkerColorbarTitleSide = "bottom"
)

// SplomMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type SplomMarkerColorbarXanchor string

const (
	SplomMarkerColorbarXanchorLeft   SplomMarkerColorbarXanchor = "left"
	SplomMarkerColorbarXanchorCenter SplomMarkerColorbarXanchor = "center"
	SplomMarkerColorbarXanchorRight  SplomMarkerColorbarXanchor = "right"
)

// SplomMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type SplomMarkerColorbarYanchor string

const (
	SplomMarkerColorbarYanchorTop    SplomMarkerColorbarYanchor = "top"
	SplomMarkerColorbarYanchorMiddle SplomMarkerColorbarYanchor = "middle"
	SplomMarkerColorbarYanchorBottom SplomMarkerColorbarYanchor = "bottom"
)

// SplomMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type SplomMarkerSizemode string

const (
	SplomMarkerSizemodeDiameter SplomMarkerSizemode = "diameter"
	SplomMarkerSizemodeArea     SplomMarkerSizemode = "area"
)

// SplomMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type SplomMarkerSymbol interface{}

var (
	SplomMarkerSymbolNumber0                 SplomMarkerSymbol = 0
	SplomMarkerSymbol0                       SplomMarkerSymbol = "0"
	SplomMarkerSymbolCircle                  SplomMarkerSymbol = "circle"
	SplomMarkerSymbolNumber100               SplomMarkerSymbol = 100
	SplomMarkerSymbol100                     SplomMarkerSymbol = "100"
	SplomMarkerSymbolCircleOpen              SplomMarkerSymbol = "circle-open"
	SplomMarkerSymbolNumber200               SplomMarkerSymbol = 200
	SplomMarkerSymbol200                     SplomMarkerSymbol = "200"
	SplomMarkerSymbolCircleDot               SplomMarkerSymbol = "circle-dot"
	SplomMarkerSymbolNumber300               SplomMarkerSymbol = 300
	SplomMarkerSymbol300                     SplomMarkerSymbol = "300"
	SplomMarkerSymbolCircleOpenDot           SplomMarkerSymbol = "circle-open-dot"
	SplomMarkerSymbolNumber1                 SplomMarkerSymbol = 1
	SplomMarkerSymbol1                       SplomMarkerSymbol = "1"
	SplomMarkerSymbolSquare                  SplomMarkerSymbol = "square"
	SplomMarkerSymbolNumber101               SplomMarkerSymbol = 101
	SplomMarkerSymbol101                     SplomMarkerSymbol = "101"
	SplomMarkerSymbolSquareOpen              SplomMarkerSymbol = "square-open"
	SplomMarkerSymbolNumber201               SplomMarkerSymbol = 201
	SplomMarkerSymbol201                     SplomMarkerSymbol = "201"
	SplomMarkerSymbolSquareDot               SplomMarkerSymbol = "square-dot"
	SplomMarkerSymbolNumber301               SplomMarkerSymbol = 301
	SplomMarkerSymbol301                     SplomMarkerSymbol = "301"
	SplomMarkerSymbolSquareOpenDot           SplomMarkerSymbol = "square-open-dot"
	SplomMarkerSymbolNumber2                 SplomMarkerSymbol = 2
	SplomMarkerSymbol2                       SplomMarkerSymbol = "2"
	SplomMarkerSymbolDiamond                 SplomMarkerSymbol = "diamond"
	SplomMarkerSymbolNumber102               SplomMarkerSymbol = 102
	SplomMarkerSymbol102                     SplomMarkerSymbol = "102"
	SplomMarkerSymbolDiamondOpen             SplomMarkerSymbol = "diamond-open"
	SplomMarkerSymbolNumber202               SplomMarkerSymbol = 202
	SplomMarkerSymbol202                     SplomMarkerSymbol = "202"
	SplomMarkerSymbolDiamondDot              SplomMarkerSymbol = "diamond-dot"
	SplomMarkerSymbolNumber302               SplomMarkerSymbol = 302
	SplomMarkerSymbol302                     SplomMarkerSymbol = "302"
	SplomMarkerSymbolDiamondOpenDot          SplomMarkerSymbol = "diamond-open-dot"
	SplomMarkerSymbolNumber3                 SplomMarkerSymbol = 3
	SplomMarkerSymbol3                       SplomMarkerSymbol = "3"
	SplomMarkerSymbolCross                   SplomMarkerSymbol = "cross"
	SplomMarkerSymbolNumber103               SplomMarkerSymbol = 103
	SplomMarkerSymbol103                     SplomMarkerSymbol = "103"
	SplomMarkerSymbolCrossOpen               SplomMarkerSymbol = "cross-open"
	SplomMarkerSymbolNumber203               SplomMarkerSymbol = 203
	SplomMarkerSymbol203                     SplomMarkerSymbol = "203"
	SplomMarkerSymbolCrossDot                SplomMarkerSymbol = "cross-dot"
	SplomMarkerSymbolNumber303               SplomMarkerSymbol = 303
	SplomMarkerSymbol303                     SplomMarkerSymbol = "303"
	SplomMarkerSymbolCrossOpenDot            SplomMarkerSymbol = "cross-open-dot"
	SplomMarkerSymbolNumber4                 SplomMarkerSymbol = 4
	SplomMarkerSymbol4                       SplomMarkerSymbol = "4"
	SplomMarkerSymbolX                       SplomMarkerSymbol = "x"
	SplomMarkerSymbolNumber104               SplomMarkerSymbol = 104
	SplomMarkerSymbol104                     SplomMarkerSymbol = "104"
	SplomMarkerSymbolXOpen                   SplomMarkerSymbol = "x-open"
	SplomMarkerSymbolNumber204               SplomMarkerSymbol = 204
	SplomMarkerSymbol204                     SplomMarkerSymbol = "204"
	SplomMarkerSymbolXDot                    SplomMarkerSymbol = "x-dot"
	SplomMarkerSymbolNumber304               SplomMarkerSymbol = 304
	SplomMarkerSymbol304                     SplomMarkerSymbol = "304"
	SplomMarkerSymbolXOpenDot                SplomMarkerSymbol = "x-open-dot"
	SplomMarkerSymbolNumber5                 SplomMarkerSymbol = 5
	SplomMarkerSymbol5                       SplomMarkerSymbol = "5"
	SplomMarkerSymbolTriangleUp              SplomMarkerSymbol = "triangle-up"
	SplomMarkerSymbolNumber105               SplomMarkerSymbol = 105
	SplomMarkerSymbol105                     SplomMarkerSymbol = "105"
	SplomMarkerSymbolTriangleUpOpen          SplomMarkerSymbol = "triangle-up-open"
	SplomMarkerSymbolNumber205               SplomMarkerSymbol = 205
	SplomMarkerSymbol205                     SplomMarkerSymbol = "205"
	SplomMarkerSymbolTriangleUpDot           SplomMarkerSymbol = "triangle-up-dot"
	SplomMarkerSymbolNumber305               SplomMarkerSymbol = 305
	SplomMarkerSymbol305                     SplomMarkerSymbol = "305"
	SplomMarkerSymbolTriangleUpOpenDot       SplomMarkerSymbol = "triangle-up-open-dot"
	SplomMarkerSymbolNumber6                 SplomMarkerSymbol = 6
	SplomMarkerSymbol6                       SplomMarkerSymbol = "6"
	SplomMarkerSymbolTriangleDown            SplomMarkerSymbol = "triangle-down"
	SplomMarkerSymbolNumber106               SplomMarkerSymbol = 106
	SplomMarkerSymbol106                     SplomMarkerSymbol = "106"
	SplomMarkerSymbolTriangleDownOpen        SplomMarkerSymbol = "triangle-down-open"
	SplomMarkerSymbolNumber206               SplomMarkerSymbol = 206
	SplomMarkerSymbol206                     SplomMarkerSymbol = "206"
	SplomMarkerSymbolTriangleDownDot         SplomMarkerSymbol = "triangle-down-dot"
	SplomMarkerSymbolNumber306               SplomMarkerSymbol = 306
	SplomMarkerSymbol306                     SplomMarkerSymbol = "306"
	SplomMarkerSymbolTriangleDownOpenDot     SplomMarkerSymbol = "triangle-down-open-dot"
	SplomMarkerSymbolNumber7                 SplomMarkerSymbol = 7
	SplomMarkerSymbol7                       SplomMarkerSymbol = "7"
	SplomMarkerSymbolTriangleLeft            SplomMarkerSymbol = "triangle-left"
	SplomMarkerSymbolNumber107               SplomMarkerSymbol = 107
	SplomMarkerSymbol107                     SplomMarkerSymbol = "107"
	SplomMarkerSymbolTriangleLeftOpen        SplomMarkerSymbol = "triangle-left-open"
	SplomMarkerSymbolNumber207               SplomMarkerSymbol = 207
	SplomMarkerSymbol207                     SplomMarkerSymbol = "207"
	SplomMarkerSymbolTriangleLeftDot         SplomMarkerSymbol = "triangle-left-dot"
	SplomMarkerSymbolNumber307               SplomMarkerSymbol = 307
	SplomMarkerSymbol307                     SplomMarkerSymbol = "307"
	SplomMarkerSymbolTriangleLeftOpenDot     SplomMarkerSymbol = "triangle-left-open-dot"
	SplomMarkerSymbolNumber8                 SplomMarkerSymbol = 8
	SplomMarkerSymbol8                       SplomMarkerSymbol = "8"
	SplomMarkerSymbolTriangleRight           SplomMarkerSymbol = "triangle-right"
	SplomMarkerSymbolNumber108               SplomMarkerSymbol = 108
	SplomMarkerSymbol108                     SplomMarkerSymbol = "108"
	SplomMarkerSymbolTriangleRightOpen       SplomMarkerSymbol = "triangle-right-open"
	SplomMarkerSymbolNumber208               SplomMarkerSymbol = 208
	SplomMarkerSymbol208                     SplomMarkerSymbol = "208"
	SplomMarkerSymbolTriangleRightDot        SplomMarkerSymbol = "triangle-right-dot"
	SplomMarkerSymbolNumber308               SplomMarkerSymbol = 308
	SplomMarkerSymbol308                     SplomMarkerSymbol = "308"
	SplomMarkerSymbolTriangleRightOpenDot    SplomMarkerSymbol = "triangle-right-open-dot"
	SplomMarkerSymbolNumber9                 SplomMarkerSymbol = 9
	SplomMarkerSymbol9                       SplomMarkerSymbol = "9"
	SplomMarkerSymbolTriangleNe              SplomMarkerSymbol = "triangle-ne"
	SplomMarkerSymbolNumber109               SplomMarkerSymbol = 109
	SplomMarkerSymbol109                     SplomMarkerSymbol = "109"
	SplomMarkerSymbolTriangleNeOpen          SplomMarkerSymbol = "triangle-ne-open"
	SplomMarkerSymbolNumber209               SplomMarkerSymbol = 209
	SplomMarkerSymbol209                     SplomMarkerSymbol = "209"
	SplomMarkerSymbolTriangleNeDot           SplomMarkerSymbol = "triangle-ne-dot"
	SplomMarkerSymbolNumber309               SplomMarkerSymbol = 309
	SplomMarkerSymbol309                     SplomMarkerSymbol = "309"
	SplomMarkerSymbolTriangleNeOpenDot       SplomMarkerSymbol = "triangle-ne-open-dot"
	SplomMarkerSymbolNumber10                SplomMarkerSymbol = 10
	SplomMarkerSymbol10                      SplomMarkerSymbol = "10"
	SplomMarkerSymbolTriangleSe              SplomMarkerSymbol = "triangle-se"
	SplomMarkerSymbolNumber110               SplomMarkerSymbol = 110
	SplomMarkerSymbol110                     SplomMarkerSymbol = "110"
	SplomMarkerSymbolTriangleSeOpen          SplomMarkerSymbol = "triangle-se-open"
	SplomMarkerSymbolNumber210               SplomMarkerSymbol = 210
	SplomMarkerSymbol210                     SplomMarkerSymbol = "210"
	SplomMarkerSymbolTriangleSeDot           SplomMarkerSymbol = "triangle-se-dot"
	SplomMarkerSymbolNumber310               SplomMarkerSymbol = 310
	SplomMarkerSymbol310                     SplomMarkerSymbol = "310"
	SplomMarkerSymbolTriangleSeOpenDot       SplomMarkerSymbol = "triangle-se-open-dot"
	SplomMarkerSymbolNumber11                SplomMarkerSymbol = 11
	SplomMarkerSymbol11                      SplomMarkerSymbol = "11"
	SplomMarkerSymbolTriangleSw              SplomMarkerSymbol = "triangle-sw"
	SplomMarkerSymbolNumber111               SplomMarkerSymbol = 111
	SplomMarkerSymbol111                     SplomMarkerSymbol = "111"
	SplomMarkerSymbolTriangleSwOpen          SplomMarkerSymbol = "triangle-sw-open"
	SplomMarkerSymbolNumber211               SplomMarkerSymbol = 211
	SplomMarkerSymbol211                     SplomMarkerSymbol = "211"
	SplomMarkerSymbolTriangleSwDot           SplomMarkerSymbol = "triangle-sw-dot"
	SplomMarkerSymbolNumber311               SplomMarkerSymbol = 311
	SplomMarkerSymbol311                     SplomMarkerSymbol = "311"
	SplomMarkerSymbolTriangleSwOpenDot       SplomMarkerSymbol = "triangle-sw-open-dot"
	SplomMarkerSymbolNumber12                SplomMarkerSymbol = 12
	SplomMarkerSymbol12                      SplomMarkerSymbol = "12"
	SplomMarkerSymbolTriangleNw              SplomMarkerSymbol = "triangle-nw"
	SplomMarkerSymbolNumber112               SplomMarkerSymbol = 112
	SplomMarkerSymbol112                     SplomMarkerSymbol = "112"
	SplomMarkerSymbolTriangleNwOpen          SplomMarkerSymbol = "triangle-nw-open"
	SplomMarkerSymbolNumber212               SplomMarkerSymbol = 212
	SplomMarkerSymbol212                     SplomMarkerSymbol = "212"
	SplomMarkerSymbolTriangleNwDot           SplomMarkerSymbol = "triangle-nw-dot"
	SplomMarkerSymbolNumber312               SplomMarkerSymbol = 312
	SplomMarkerSymbol312                     SplomMarkerSymbol = "312"
	SplomMarkerSymbolTriangleNwOpenDot       SplomMarkerSymbol = "triangle-nw-open-dot"
	SplomMarkerSymbolNumber13                SplomMarkerSymbol = 13
	SplomMarkerSymbol13                      SplomMarkerSymbol = "13"
	SplomMarkerSymbolPentagon                SplomMarkerSymbol = "pentagon"
	SplomMarkerSymbolNumber113               SplomMarkerSymbol = 113
	SplomMarkerSymbol113                     SplomMarkerSymbol = "113"
	SplomMarkerSymbolPentagonOpen            SplomMarkerSymbol = "pentagon-open"
	SplomMarkerSymbolNumber213               SplomMarkerSymbol = 213
	SplomMarkerSymbol213                     SplomMarkerSymbol = "213"
	SplomMarkerSymbolPentagonDot             SplomMarkerSymbol = "pentagon-dot"
	SplomMarkerSymbolNumber313               SplomMarkerSymbol = 313
	SplomMarkerSymbol313                     SplomMarkerSymbol = "313"
	SplomMarkerSymbolPentagonOpenDot         SplomMarkerSymbol = "pentagon-open-dot"
	SplomMarkerSymbolNumber14                SplomMarkerSymbol = 14
	SplomMarkerSymbol14                      SplomMarkerSymbol = "14"
	SplomMarkerSymbolHexagon                 SplomMarkerSymbol = "hexagon"
	SplomMarkerSymbolNumber114               SplomMarkerSymbol = 114
	SplomMarkerSymbol114                     SplomMarkerSymbol = "114"
	SplomMarkerSymbolHexagonOpen             SplomMarkerSymbol = "hexagon-open"
	SplomMarkerSymbolNumber214               SplomMarkerSymbol = 214
	SplomMarkerSymbol214                     SplomMarkerSymbol = "214"
	SplomMarkerSymbolHexagonDot              SplomMarkerSymbol = "hexagon-dot"
	SplomMarkerSymbolNumber314               SplomMarkerSymbol = 314
	SplomMarkerSymbol314                     SplomMarkerSymbol = "314"
	SplomMarkerSymbolHexagonOpenDot          SplomMarkerSymbol = "hexagon-open-dot"
	SplomMarkerSymbolNumber15                SplomMarkerSymbol = 15
	SplomMarkerSymbol15                      SplomMarkerSymbol = "15"
	SplomMarkerSymbolHexagon2                SplomMarkerSymbol = "hexagon2"
	SplomMarkerSymbolNumber115               SplomMarkerSymbol = 115
	SplomMarkerSymbol115                     SplomMarkerSymbol = "115"
	SplomMarkerSymbolHexagon2Open            SplomMarkerSymbol = "hexagon2-open"
	SplomMarkerSymbolNumber215               SplomMarkerSymbol = 215
	SplomMarkerSymbol215                     SplomMarkerSymbol = "215"
	SplomMarkerSymbolHexagon2Dot             SplomMarkerSymbol = "hexagon2-dot"
	SplomMarkerSymbolNumber315               SplomMarkerSymbol = 315
	SplomMarkerSymbol315                     SplomMarkerSymbol = "315"
	SplomMarkerSymbolHexagon2OpenDot         SplomMarkerSymbol = "hexagon2-open-dot"
	SplomMarkerSymbolNumber16                SplomMarkerSymbol = 16
	SplomMarkerSymbol16                      SplomMarkerSymbol = "16"
	SplomMarkerSymbolOctagon                 SplomMarkerSymbol = "octagon"
	SplomMarkerSymbolNumber116               SplomMarkerSymbol = 116
	SplomMarkerSymbol116                     SplomMarkerSymbol = "116"
	SplomMarkerSymbolOctagonOpen             SplomMarkerSymbol = "octagon-open"
	SplomMarkerSymbolNumber216               SplomMarkerSymbol = 216
	SplomMarkerSymbol216                     SplomMarkerSymbol = "216"
	SplomMarkerSymbolOctagonDot              SplomMarkerSymbol = "octagon-dot"
	SplomMarkerSymbolNumber316               SplomMarkerSymbol = 316
	SplomMarkerSymbol316                     SplomMarkerSymbol = "316"
	SplomMarkerSymbolOctagonOpenDot          SplomMarkerSymbol = "octagon-open-dot"
	SplomMarkerSymbolNumber17                SplomMarkerSymbol = 17
	SplomMarkerSymbol17                      SplomMarkerSymbol = "17"
	SplomMarkerSymbolStar                    SplomMarkerSymbol = "star"
	SplomMarkerSymbolNumber117               SplomMarkerSymbol = 117
	SplomMarkerSymbol117                     SplomMarkerSymbol = "117"
	SplomMarkerSymbolStarOpen                SplomMarkerSymbol = "star-open"
	SplomMarkerSymbolNumber217               SplomMarkerSymbol = 217
	SplomMarkerSymbol217                     SplomMarkerSymbol = "217"
	SplomMarkerSymbolStarDot                 SplomMarkerSymbol = "star-dot"
	SplomMarkerSymbolNumber317               SplomMarkerSymbol = 317
	SplomMarkerSymbol317                     SplomMarkerSymbol = "317"
	SplomMarkerSymbolStarOpenDot             SplomMarkerSymbol = "star-open-dot"
	SplomMarkerSymbolNumber18                SplomMarkerSymbol = 18
	SplomMarkerSymbol18                      SplomMarkerSymbol = "18"
	SplomMarkerSymbolHexagram                SplomMarkerSymbol = "hexagram"
	SplomMarkerSymbolNumber118               SplomMarkerSymbol = 118
	SplomMarkerSymbol118                     SplomMarkerSymbol = "118"
	SplomMarkerSymbolHexagramOpen            SplomMarkerSymbol = "hexagram-open"
	SplomMarkerSymbolNumber218               SplomMarkerSymbol = 218
	SplomMarkerSymbol218                     SplomMarkerSymbol = "218"
	SplomMarkerSymbolHexagramDot             SplomMarkerSymbol = "hexagram-dot"
	SplomMarkerSymbolNumber318               SplomMarkerSymbol = 318
	SplomMarkerSymbol318                     SplomMarkerSymbol = "318"
	SplomMarkerSymbolHexagramOpenDot         SplomMarkerSymbol = "hexagram-open-dot"
	SplomMarkerSymbolNumber19                SplomMarkerSymbol = 19
	SplomMarkerSymbol19                      SplomMarkerSymbol = "19"
	SplomMarkerSymbolStarTriangleUp          SplomMarkerSymbol = "star-triangle-up"
	SplomMarkerSymbolNumber119               SplomMarkerSymbol = 119
	SplomMarkerSymbol119                     SplomMarkerSymbol = "119"
	SplomMarkerSymbolStarTriangleUpOpen      SplomMarkerSymbol = "star-triangle-up-open"
	SplomMarkerSymbolNumber219               SplomMarkerSymbol = 219
	SplomMarkerSymbol219                     SplomMarkerSymbol = "219"
	SplomMarkerSymbolStarTriangleUpDot       SplomMarkerSymbol = "star-triangle-up-dot"
	SplomMarkerSymbolNumber319               SplomMarkerSymbol = 319
	SplomMarkerSymbol319                     SplomMarkerSymbol = "319"
	SplomMarkerSymbolStarTriangleUpOpenDot   SplomMarkerSymbol = "star-triangle-up-open-dot"
	SplomMarkerSymbolNumber20                SplomMarkerSymbol = 20
	SplomMarkerSymbol20                      SplomMarkerSymbol = "20"
	SplomMarkerSymbolStarTriangleDown        SplomMarkerSymbol = "star-triangle-down"
	SplomMarkerSymbolNumber120               SplomMarkerSymbol = 120
	SplomMarkerSymbol120                     SplomMarkerSymbol = "120"
	SplomMarkerSymbolStarTriangleDownOpen    SplomMarkerSymbol = "star-triangle-down-open"
	SplomMarkerSymbolNumber220               SplomMarkerSymbol = 220
	SplomMarkerSymbol220                     SplomMarkerSymbol = "220"
	SplomMarkerSymbolStarTriangleDownDot     SplomMarkerSymbol = "star-triangle-down-dot"
	SplomMarkerSymbolNumber320               SplomMarkerSymbol = 320
	SplomMarkerSymbol320                     SplomMarkerSymbol = "320"
	SplomMarkerSymbolStarTriangleDownOpenDot SplomMarkerSymbol = "star-triangle-down-open-dot"
	SplomMarkerSymbolNumber21                SplomMarkerSymbol = 21
	SplomMarkerSymbol21                      SplomMarkerSymbol = "21"
	SplomMarkerSymbolStarSquare              SplomMarkerSymbol = "star-square"
	SplomMarkerSymbolNumber121               SplomMarkerSymbol = 121
	SplomMarkerSymbol121                     SplomMarkerSymbol = "121"
	SplomMarkerSymbolStarSquareOpen          SplomMarkerSymbol = "star-square-open"
	SplomMarkerSymbolNumber221               SplomMarkerSymbol = 221
	SplomMarkerSymbol221                     SplomMarkerSymbol = "221"
	SplomMarkerSymbolStarSquareDot           SplomMarkerSymbol = "star-square-dot"
	SplomMarkerSymbolNumber321               SplomMarkerSymbol = 321
	SplomMarkerSymbol321                     SplomMarkerSymbol = "321"
	SplomMarkerSymbolStarSquareOpenDot       SplomMarkerSymbol = "star-square-open-dot"
	SplomMarkerSymbolNumber22                SplomMarkerSymbol = 22
	SplomMarkerSymbol22                      SplomMarkerSymbol = "22"
	SplomMarkerSymbolStarDiamond             SplomMarkerSymbol = "star-diamond"
	SplomMarkerSymbolNumber122               SplomMarkerSymbol = 122
	SplomMarkerSymbol122                     SplomMarkerSymbol = "122"
	SplomMarkerSymbolStarDiamondOpen         SplomMarkerSymbol = "star-diamond-open"
	SplomMarkerSymbolNumber222               SplomMarkerSymbol = 222
	SplomMarkerSymbol222                     SplomMarkerSymbol = "222"
	SplomMarkerSymbolStarDiamondDot          SplomMarkerSymbol = "star-diamond-dot"
	SplomMarkerSymbolNumber322               SplomMarkerSymbol = 322
	SplomMarkerSymbol322                     SplomMarkerSymbol = "322"
	SplomMarkerSymbolStarDiamondOpenDot      SplomMarkerSymbol = "star-diamond-open-dot"
	SplomMarkerSymbolNumber23                SplomMarkerSymbol = 23
	SplomMarkerSymbol23                      SplomMarkerSymbol = "23"
	SplomMarkerSymbolDiamondTall             SplomMarkerSymbol = "diamond-tall"
	SplomMarkerSymbolNumber123               SplomMarkerSymbol = 123
	SplomMarkerSymbol123                     SplomMarkerSymbol = "123"
	SplomMarkerSymbolDiamondTallOpen         SplomMarkerSymbol = "diamond-tall-open"
	SplomMarkerSymbolNumber223               SplomMarkerSymbol = 223
	SplomMarkerSymbol223                     SplomMarkerSymbol = "223"
	SplomMarkerSymbolDiamondTallDot          SplomMarkerSymbol = "diamond-tall-dot"
	SplomMarkerSymbolNumber323               SplomMarkerSymbol = 323
	SplomMarkerSymbol323                     SplomMarkerSymbol = "323"
	SplomMarkerSymbolDiamondTallOpenDot      SplomMarkerSymbol = "diamond-tall-open-dot"
	SplomMarkerSymbolNumber24                SplomMarkerSymbol = 24
	SplomMarkerSymbol24                      SplomMarkerSymbol = "24"
	SplomMarkerSymbolDiamondWide             SplomMarkerSymbol = "diamond-wide"
	SplomMarkerSymbolNumber124               SplomMarkerSymbol = 124
	SplomMarkerSymbol124                     SplomMarkerSymbol = "124"
	SplomMarkerSymbolDiamondWideOpen         SplomMarkerSymbol = "diamond-wide-open"
	SplomMarkerSymbolNumber224               SplomMarkerSymbol = 224
	SplomMarkerSymbol224                     SplomMarkerSymbol = "224"
	SplomMarkerSymbolDiamondWideDot          SplomMarkerSymbol = "diamond-wide-dot"
	SplomMarkerSymbolNumber324               SplomMarkerSymbol = 324
	SplomMarkerSymbol324                     SplomMarkerSymbol = "324"
	SplomMarkerSymbolDiamondWideOpenDot      SplomMarkerSymbol = "diamond-wide-open-dot"
	SplomMarkerSymbolNumber25                SplomMarkerSymbol = 25
	SplomMarkerSymbol25                      SplomMarkerSymbol = "25"
	SplomMarkerSymbolHourglass               SplomMarkerSymbol = "hourglass"
	SplomMarkerSymbolNumber125               SplomMarkerSymbol = 125
	SplomMarkerSymbol125                     SplomMarkerSymbol = "125"
	SplomMarkerSymbolHourglassOpen           SplomMarkerSymbol = "hourglass-open"
	SplomMarkerSymbolNumber26                SplomMarkerSymbol = 26
	SplomMarkerSymbol26                      SplomMarkerSymbol = "26"
	SplomMarkerSymbolBowtie                  SplomMarkerSymbol = "bowtie"
	SplomMarkerSymbolNumber126               SplomMarkerSymbol = 126
	SplomMarkerSymbol126                     SplomMarkerSymbol = "126"
	SplomMarkerSymbolBowtieOpen              SplomMarkerSymbol = "bowtie-open"
	SplomMarkerSymbolNumber27                SplomMarkerSymbol = 27
	SplomMarkerSymbol27                      SplomMarkerSymbol = "27"
	SplomMarkerSymbolCircleCross             SplomMarkerSymbol = "circle-cross"
	SplomMarkerSymbolNumber127               SplomMarkerSymbol = 127
	SplomMarkerSymbol127                     SplomMarkerSymbol = "127"
	SplomMarkerSymbolCircleCrossOpen         SplomMarkerSymbol = "circle-cross-open"
	SplomMarkerSymbolNumber28                SplomMarkerSymbol = 28
	SplomMarkerSymbol28                      SplomMarkerSymbol = "28"
	SplomMarkerSymbolCircleX                 SplomMarkerSymbol = "circle-x"
	SplomMarkerSymbolNumber128               SplomMarkerSymbol = 128
	SplomMarkerSymbol128                     SplomMarkerSymbol = "128"
	SplomMarkerSymbolCircleXOpen             SplomMarkerSymbol = "circle-x-open"
	SplomMarkerSymbolNumber29                SplomMarkerSymbol = 29
	SplomMarkerSymbol29                      SplomMarkerSymbol = "29"
	SplomMarkerSymbolSquareCross             SplomMarkerSymbol = "square-cross"
	SplomMarkerSymbolNumber129               SplomMarkerSymbol = 129
	SplomMarkerSymbol129                     SplomMarkerSymbol = "129"
	SplomMarkerSymbolSquareCrossOpen         SplomMarkerSymbol = "square-cross-open"
	SplomMarkerSymbolNumber30                SplomMarkerSymbol = 30
	SplomMarkerSymbol30                      SplomMarkerSymbol = "30"
	SplomMarkerSymbolSquareX                 SplomMarkerSymbol = "square-x"
	SplomMarkerSymbolNumber130               SplomMarkerSymbol = 130
	SplomMarkerSymbol130                     SplomMarkerSymbol = "130"
	SplomMarkerSymbolSquareXOpen             SplomMarkerSymbol = "square-x-open"
	SplomMarkerSymbolNumber31                SplomMarkerSymbol = 31
	SplomMarkerSymbol31                      SplomMarkerSymbol = "31"
	SplomMarkerSymbolDiamondCross            SplomMarkerSymbol = "diamond-cross"
	SplomMarkerSymbolNumber131               SplomMarkerSymbol = 131
	SplomMarkerSymbol131                     SplomMarkerSymbol = "131"
	SplomMarkerSymbolDiamondCrossOpen        SplomMarkerSymbol = "diamond-cross-open"
	SplomMarkerSymbolNumber32                SplomMarkerSymbol = 32
	SplomMarkerSymbol32                      SplomMarkerSymbol = "32"
	SplomMarkerSymbolDiamondX                SplomMarkerSymbol = "diamond-x"
	SplomMarkerSymbolNumber132               SplomMarkerSymbol = 132
	SplomMarkerSymbol132                     SplomMarkerSymbol = "132"
	SplomMarkerSymbolDiamondXOpen            SplomMarkerSymbol = "diamond-x-open"
	SplomMarkerSymbolNumber33                SplomMarkerSymbol = 33
	SplomMarkerSymbol33                      SplomMarkerSymbol = "33"
	SplomMarkerSymbolCrossThin               SplomMarkerSymbol = "cross-thin"
	SplomMarkerSymbolNumber133               SplomMarkerSymbol = 133
	SplomMarkerSymbol133                     SplomMarkerSymbol = "133"
	SplomMarkerSymbolCrossThinOpen           SplomMarkerSymbol = "cross-thin-open"
	SplomMarkerSymbolNumber34                SplomMarkerSymbol = 34
	SplomMarkerSymbol34                      SplomMarkerSymbol = "34"
	SplomMarkerSymbolXThin                   SplomMarkerSymbol = "x-thin"
	SplomMarkerSymbolNumber134               SplomMarkerSymbol = 134
	SplomMarkerSymbol134                     SplomMarkerSymbol = "134"
	SplomMarkerSymbolXThinOpen               SplomMarkerSymbol = "x-thin-open"
	SplomMarkerSymbolNumber35                SplomMarkerSymbol = 35
	SplomMarkerSymbol35                      SplomMarkerSymbol = "35"
	SplomMarkerSymbolAsterisk                SplomMarkerSymbol = "asterisk"
	SplomMarkerSymbolNumber135               SplomMarkerSymbol = 135
	SplomMarkerSymbol135                     SplomMarkerSymbol = "135"
	SplomMarkerSymbolAsteriskOpen            SplomMarkerSymbol = "asterisk-open"
	SplomMarkerSymbolNumber36                SplomMarkerSymbol = 36
	SplomMarkerSymbol36                      SplomMarkerSymbol = "36"
	SplomMarkerSymbolHash                    SplomMarkerSymbol = "hash"
	SplomMarkerSymbolNumber136               SplomMarkerSymbol = 136
	SplomMarkerSymbol136                     SplomMarkerSymbol = "136"
	SplomMarkerSymbolHashOpen                SplomMarkerSymbol = "hash-open"
	SplomMarkerSymbolNumber236               SplomMarkerSymbol = 236
	SplomMarkerSymbol236                     SplomMarkerSymbol = "236"
	SplomMarkerSymbolHashDot                 SplomMarkerSymbol = "hash-dot"
	SplomMarkerSymbolNumber336               SplomMarkerSymbol = 336
	SplomMarkerSymbol336                     SplomMarkerSymbol = "336"
	SplomMarkerSymbolHashOpenDot             SplomMarkerSymbol = "hash-open-dot"
	SplomMarkerSymbolNumber37                SplomMarkerSymbol = 37
	SplomMarkerSymbol37                      SplomMarkerSymbol = "37"
	SplomMarkerSymbolYUp                     SplomMarkerSymbol = "y-up"
	SplomMarkerSymbolNumber137               SplomMarkerSymbol = 137
	SplomMarkerSymbol137                     SplomMarkerSymbol = "137"
	SplomMarkerSymbolYUpOpen                 SplomMarkerSymbol = "y-up-open"
	SplomMarkerSymbolNumber38                SplomMarkerSymbol = 38
	SplomMarkerSymbol38                      SplomMarkerSymbol = "38"
	SplomMarkerSymbolYDown                   SplomMarkerSymbol = "y-down"
	SplomMarkerSymbolNumber138               SplomMarkerSymbol = 138
	SplomMarkerSymbol138                     SplomMarkerSymbol = "138"
	SplomMarkerSymbolYDownOpen               SplomMarkerSymbol = "y-down-open"
	SplomMarkerSymbolNumber39                SplomMarkerSymbol = 39
	SplomMarkerSymbol39                      SplomMarkerSymbol = "39"
	SplomMarkerSymbolYLeft                   SplomMarkerSymbol = "y-left"
	SplomMarkerSymbolNumber139               SplomMarkerSymbol = 139
	SplomMarkerSymbol139                     SplomMarkerSymbol = "139"
	SplomMarkerSymbolYLeftOpen               SplomMarkerSymbol = "y-left-open"
	SplomMarkerSymbolNumber40                SplomMarkerSymbol = 40
	SplomMarkerSymbol40                      SplomMarkerSymbol = "40"
	SplomMarkerSymbolYRight                  SplomMarkerSymbol = "y-right"
	SplomMarkerSymbolNumber140               SplomMarkerSymbol = 140
	SplomMarkerSymbol140                     SplomMarkerSymbol = "140"
	SplomMarkerSymbolYRightOpen              SplomMarkerSymbol = "y-right-open"
	SplomMarkerSymbolNumber41                SplomMarkerSymbol = 41
	SplomMarkerSymbol41                      SplomMarkerSymbol = "41"
	SplomMarkerSymbolLineEw                  SplomMarkerSymbol = "line-ew"
	SplomMarkerSymbolNumber141               SplomMarkerSymbol = 141
	SplomMarkerSymbol141                     SplomMarkerSymbol = "141"
	SplomMarkerSymbolLineEwOpen              SplomMarkerSymbol = "line-ew-open"
	SplomMarkerSymbolNumber42                SplomMarkerSymbol = 42
	SplomMarkerSymbol42                      SplomMarkerSymbol = "42"
	SplomMarkerSymbolLineNs                  SplomMarkerSymbol = "line-ns"
	SplomMarkerSymbolNumber142               SplomMarkerSymbol = 142
	SplomMarkerSymbol142                     SplomMarkerSymbol = "142"
	SplomMarkerSymbolLineNsOpen              SplomMarkerSymbol = "line-ns-open"
	SplomMarkerSymbolNumber43                SplomMarkerSymbol = 43
	SplomMarkerSymbol43                      SplomMarkerSymbol = "43"
	SplomMarkerSymbolLineNe                  SplomMarkerSymbol = "line-ne"
	SplomMarkerSymbolNumber143               SplomMarkerSymbol = 143
	SplomMarkerSymbol143                     SplomMarkerSymbol = "143"
	SplomMarkerSymbolLineNeOpen              SplomMarkerSymbol = "line-ne-open"
	SplomMarkerSymbolNumber44                SplomMarkerSymbol = 44
	SplomMarkerSymbol44                      SplomMarkerSymbol = "44"
	SplomMarkerSymbolLineNw                  SplomMarkerSymbol = "line-nw"
	SplomMarkerSymbolNumber144               SplomMarkerSymbol = 144
	SplomMarkerSymbol144                     SplomMarkerSymbol = "144"
	SplomMarkerSymbolLineNwOpen              SplomMarkerSymbol = "line-nw-open"
	SplomMarkerSymbolNumber45                SplomMarkerSymbol = 45
	SplomMarkerSymbol45                      SplomMarkerSymbol = "45"
	SplomMarkerSymbolArrowUp                 SplomMarkerSymbol = "arrow-up"
	SplomMarkerSymbolNumber145               SplomMarkerSymbol = 145
	SplomMarkerSymbol145                     SplomMarkerSymbol = "145"
	SplomMarkerSymbolArrowUpOpen             SplomMarkerSymbol = "arrow-up-open"
	SplomMarkerSymbolNumber46                SplomMarkerSymbol = 46
	SplomMarkerSymbol46                      SplomMarkerSymbol = "46"
	SplomMarkerSymbolArrowDown               SplomMarkerSymbol = "arrow-down"
	SplomMarkerSymbolNumber146               SplomMarkerSymbol = 146
	SplomMarkerSymbol146                     SplomMarkerSymbol = "146"
	SplomMarkerSymbolArrowDownOpen           SplomMarkerSymbol = "arrow-down-open"
	SplomMarkerSymbolNumber47                SplomMarkerSymbol = 47
	SplomMarkerSymbol47                      SplomMarkerSymbol = "47"
	SplomMarkerSymbolArrowLeft               SplomMarkerSymbol = "arrow-left"
	SplomMarkerSymbolNumber147               SplomMarkerSymbol = 147
	SplomMarkerSymbol147                     SplomMarkerSymbol = "147"
	SplomMarkerSymbolArrowLeftOpen           SplomMarkerSymbol = "arrow-left-open"
	SplomMarkerSymbolNumber48                SplomMarkerSymbol = 48
	SplomMarkerSymbol48                      SplomMarkerSymbol = "48"
	SplomMarkerSymbolArrowRight              SplomMarkerSymbol = "arrow-right"
	SplomMarkerSymbolNumber148               SplomMarkerSymbol = 148
	SplomMarkerSymbol148                     SplomMarkerSymbol = "148"
	SplomMarkerSymbolArrowRightOpen          SplomMarkerSymbol = "arrow-right-open"
	SplomMarkerSymbolNumber49                SplomMarkerSymbol = 49
	SplomMarkerSymbol49                      SplomMarkerSymbol = "49"
	SplomMarkerSymbolArrowBarUp              SplomMarkerSymbol = "arrow-bar-up"
	SplomMarkerSymbolNumber149               SplomMarkerSymbol = 149
	SplomMarkerSymbol149                     SplomMarkerSymbol = "149"
	SplomMarkerSymbolArrowBarUpOpen          SplomMarkerSymbol = "arrow-bar-up-open"
	SplomMarkerSymbolNumber50                SplomMarkerSymbol = 50
	SplomMarkerSymbol50                      SplomMarkerSymbol = "50"
	SplomMarkerSymbolArrowBarDown            SplomMarkerSymbol = "arrow-bar-down"
	SplomMarkerSymbolNumber150               SplomMarkerSymbol = 150
	SplomMarkerSymbol150                     SplomMarkerSymbol = "150"
	SplomMarkerSymbolArrowBarDownOpen        SplomMarkerSymbol = "arrow-bar-down-open"
	SplomMarkerSymbolNumber51                SplomMarkerSymbol = 51
	SplomMarkerSymbol51                      SplomMarkerSymbol = "51"
	SplomMarkerSymbolArrowBarLeft            SplomMarkerSymbol = "arrow-bar-left"
	SplomMarkerSymbolNumber151               SplomMarkerSymbol = 151
	SplomMarkerSymbol151                     SplomMarkerSymbol = "151"
	SplomMarkerSymbolArrowBarLeftOpen        SplomMarkerSymbol = "arrow-bar-left-open"
	SplomMarkerSymbolNumber52                SplomMarkerSymbol = 52
	SplomMarkerSymbol52                      SplomMarkerSymbol = "52"
	SplomMarkerSymbolArrowBarRight           SplomMarkerSymbol = "arrow-bar-right"
	SplomMarkerSymbolNumber152               SplomMarkerSymbol = 152
	SplomMarkerSymbol152                     SplomMarkerSymbol = "152"
	SplomMarkerSymbolArrowBarRightOpen       SplomMarkerSymbol = "arrow-bar-right-open"
)

// SplomVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type SplomVisible interface{}

var (
	SplomVisibleTrue       SplomVisible = true
	SplomVisibleFalse      SplomVisible = false
	SplomVisibleLegendonly SplomVisible = "legendonly"
)

// SplomHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type SplomHoverinfo string

const (
	// Flags
	SplomHoverinfoX    SplomHoverinfo = "x"
	SplomHoverinfoY    SplomHoverinfo = "y"
	SplomHoverinfoZ    SplomHoverinfo = "z"
	SplomHoverinfoText SplomHoverinfo = "text"
	SplomHoverinfoName SplomHoverinfo = "name"

	// Extra
	SplomHoverinfoAll  SplomHoverinfo = "all"
	SplomHoverinfoNone SplomHoverinfo = "none"
	SplomHoverinfoSkip SplomHoverinfo = "skip"
)
