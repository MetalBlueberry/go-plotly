package grob

var TraceTypeCarpet TraceType = "carpet"

func (trace *Carpet) GetType() TraceType {
	return TraceTypeCarpet
}

// Carpet The data describing carpet axis layout is set in `y` and (optionally) also `x`. If only `y` is present, `x` the plot is interpreted as a cheater plot and is filled in using the `y` values. `x` and `y` may either be 2D arrays matching with each dimension matching that of `a` and `b`, or they may be 1D arrays with total length equal to that of `a` and `b`.
type Carpet struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// A
	// arrayOK: false
	// type: data_array
	// An array containing values of the first parameter value
	A interface{} `json:"a,omitempty"`

	// A0
	// arrayOK: false
	// type: number
	// Alternate to `a`. Builds a linear space of a coordinates. Use with `da` where `a0` is the starting coordinate and `da` the step.
	A0 float64 `json:"a0,omitempty"`

	// Aaxis
	// role: Object
	Aaxis *CarpetAaxis `json:"aaxis,omitempty"`

	// Asrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  a .
	Asrc String `json:"asrc,omitempty"`

	// B
	// arrayOK: false
	// type: data_array
	// A two dimensional array of y coordinates at each carpet point.
	B interface{} `json:"b,omitempty"`

	// B0
	// arrayOK: false
	// type: number
	// Alternate to `b`. Builds a linear space of a coordinates. Use with `db` where `b0` is the starting coordinate and `db` the step.
	B0 float64 `json:"b0,omitempty"`

	// Baxis
	// role: Object
	Baxis *CarpetBaxis `json:"baxis,omitempty"`

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

	// Cheaterslope
	// arrayOK: false
	// type: number
	// The shift applied to each successive row of data in creating a cheater plot. Only used if `x` is been omitted.
	Cheaterslope float64 `json:"cheaterslope,omitempty"`

	// Color
	// arrayOK: false
	// type: color
	// Sets default for all colors associated with this axis all at once: line, font, tick, and grid colors. Grid color is lightened by blending this with the plot background Individual pieces can override this.
	Color Color `json:"color,omitempty"`

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

	// Da
	// arrayOK: false
	// type: number
	// Sets the a coordinate step. See `a0` for more info.
	Da float64 `json:"da,omitempty"`

	// Db
	// arrayOK: false
	// type: number
	// Sets the b coordinate step. See `b0` for more info.
	Db float64 `json:"db,omitempty"`

	// Font
	// role: Object
	Font *CarpetFont `json:"font,omitempty"`

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

	// Stream
	// role: Object
	Stream *CarpetStream `json:"stream,omitempty"`

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
	Visible CarpetVisible `json:"visible,omitempty"`

	// X
	// arrayOK: false
	// type: data_array
	// A two dimensional array of x coordinates at each carpet point. If omitted, the plot is a cheater plot and the xaxis is hidden by default.
	X interface{} `json:"x,omitempty"`

	// Xaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y
	// arrayOK: false
	// type: data_array
	// A two dimensional array of y coordinates at each carpet point.
	Y interface{} `json:"y,omitempty"`

	// Yaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

// CarpetAaxisTickfont Sets the tick font.
type CarpetAaxisTickfont struct {

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

// CarpetAaxisTitleFont Sets this axis' title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type CarpetAaxisTitleFont struct {

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

// CarpetAaxisTitle
type CarpetAaxisTitle struct {

	// Font
	// role: Object
	Font *CarpetAaxisTitleFont `json:"font,omitempty"`

	// Offset
	// arrayOK: false
	// type: number
	// An additional amount by which to offset the title from the tick labels, given in pixels. Note that this used to be set by the now deprecated `titleoffset` attribute.
	Offset float64 `json:"offset,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of this axis. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// CarpetAaxis
type CarpetAaxis struct {

	// Arraydtick
	// arrayOK: false
	// type: integer
	// The stride between grid lines along the axis
	Arraydtick int64 `json:"arraydtick,omitempty"`

	// Arraytick0
	// arrayOK: false
	// type: integer
	// The starting index of grid lines along the axis
	Arraytick0 int64 `json:"arraytick0,omitempty"`

	// Autorange
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
	Autorange CarpetAaxisAutorange `json:"autorange,omitempty"`

	// Autotypenumbers
	// default: convert types
	// type: enumerated
	// Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers.
	Autotypenumbers CarpetAaxisAutotypenumbers `json:"autotypenumbers,omitempty"`

	// Categoryarray
	// arrayOK: false
	// type: data_array
	// Sets the order in which categories on this axis appear. Only has an effect if `categoryorder` is set to *array*. Used with `categoryorder`.
	Categoryarray interface{} `json:"categoryarray,omitempty"`

	// Categoryarraysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  categoryarray .
	Categoryarraysrc String `json:"categoryarraysrc,omitempty"`

	// Categoryorder
	// default: trace
	// type: enumerated
	// Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`.
	Categoryorder CarpetAaxisCategoryorder `json:"categoryorder,omitempty"`

	// Cheatertype
	// default: value
	// type: enumerated
	//
	Cheatertype CarpetAaxisCheatertype `json:"cheatertype,omitempty"`

	// Color
	// arrayOK: false
	// type: color
	// Sets default for all colors associated with this axis all at once: line, font, tick, and grid colors. Grid color is lightened by blending this with the plot background Individual pieces can override this.
	Color Color `json:"color,omitempty"`

	// Dtick
	// arrayOK: false
	// type: number
	// The stride between grid lines along the axis
	Dtick float64 `json:"dtick,omitempty"`

	// Endline
	// arrayOK: false
	// type: boolean
	// Determines whether or not a line is drawn at along the final value of this axis. If *true*, the end line is drawn on top of the grid lines.
	Endline Bool `json:"endline,omitempty"`

	// Endlinecolor
	// arrayOK: false
	// type: color
	// Sets the line color of the end line.
	Endlinecolor Color `json:"endlinecolor,omitempty"`

	// Endlinewidth
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the end line.
	Endlinewidth float64 `json:"endlinewidth,omitempty"`

	// Exponentformat
	// default: B
	// type: enumerated
	// Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
	Exponentformat CarpetAaxisExponentformat `json:"exponentformat,omitempty"`

	// Fixedrange
	// arrayOK: false
	// type: boolean
	// Determines whether or not this axis is zoom-able. If true, then zoom is disabled.
	Fixedrange Bool `json:"fixedrange,omitempty"`

	// Gridcolor
	// arrayOK: false
	// type: color
	// Sets the axis line color.
	Gridcolor Color `json:"gridcolor,omitempty"`

	// Gridwidth
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the axis line.
	Gridwidth float64 `json:"gridwidth,omitempty"`

	// Labelpadding
	// arrayOK: false
	// type: integer
	// Extra padding between label and the axis
	Labelpadding int64 `json:"labelpadding,omitempty"`

	// Labelprefix
	// arrayOK: false
	// type: string
	// Sets a axis label prefix.
	Labelprefix String `json:"labelprefix,omitempty"`

	// Labelsuffix
	// arrayOK: false
	// type: string
	// Sets a axis label suffix.
	Labelsuffix String `json:"labelsuffix,omitempty"`

	// Linecolor
	// arrayOK: false
	// type: color
	// Sets the axis line color.
	Linecolor Color `json:"linecolor,omitempty"`

	// Linewidth
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the axis line.
	Linewidth float64 `json:"linewidth,omitempty"`

	// Minexponent
	// arrayOK: false
	// type: number
	// Hide SI prefix for 10^n if |n| is below this number
	Minexponent float64 `json:"minexponent,omitempty"`

	// Minorgridcolor
	// arrayOK: false
	// type: color
	// Sets the color of the grid lines.
	Minorgridcolor Color `json:"minorgridcolor,omitempty"`

	// Minorgridcount
	// arrayOK: false
	// type: integer
	// Sets the number of minor grid ticks per major grid tick
	Minorgridcount int64 `json:"minorgridcount,omitempty"`

	// Minorgridwidth
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the grid lines.
	Minorgridwidth float64 `json:"minorgridwidth,omitempty"`

	// Nticks
	// arrayOK: false
	// type: integer
	// Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*.
	Nticks int64 `json:"nticks,omitempty"`

	// Range
	// arrayOK: false
	// type: info_array
	// Sets the range of this axis. If the axis `type` is *log*, then you must take the log of your desired range (e.g. to set the range from 1 to 100, set the range from 0 to 2). If the axis `type` is *date*, it should be date strings, like date data, though Date objects and unix milliseconds will be accepted and converted to strings. If the axis `type` is *category*, it should be numbers, using the scale where each category is assigned a serial number from zero in the order it appears.
	Range interface{} `json:"range,omitempty"`

	// Rangemode
	// default: normal
	// type: enumerated
	// If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data.
	Rangemode CarpetAaxisRangemode `json:"rangemode,omitempty"`

	// Separatethousands
	// arrayOK: false
	// type: boolean
	// If "true", even 4-digit integers are separated
	Separatethousands Bool `json:"separatethousands,omitempty"`

	// Showexponent
	// default: all
	// type: enumerated
	// If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
	Showexponent CarpetAaxisShowexponent `json:"showexponent,omitempty"`

	// Showgrid
	// arrayOK: false
	// type: boolean
	// Determines whether or not grid lines are drawn. If *true*, the grid lines are drawn at every tick mark.
	Showgrid Bool `json:"showgrid,omitempty"`

	// Showline
	// arrayOK: false
	// type: boolean
	// Determines whether or not a line bounding this axis is drawn.
	Showline Bool `json:"showline,omitempty"`

	// Showticklabels
	// default: start
	// type: enumerated
	// Determines whether axis labels are drawn on the low side, the high side, both, or neither side of the axis.
	Showticklabels CarpetAaxisShowticklabels `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix CarpetAaxisShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix CarpetAaxisShowticksuffix `json:"showticksuffix,omitempty"`

	// Smoothing
	// arrayOK: false
	// type: number
	//
	Smoothing float64 `json:"smoothing,omitempty"`

	// Startline
	// arrayOK: false
	// type: boolean
	// Determines whether or not a line is drawn at along the starting value of this axis. If *true*, the start line is drawn on top of the grid lines.
	Startline Bool `json:"startline,omitempty"`

	// Startlinecolor
	// arrayOK: false
	// type: color
	// Sets the line color of the start line.
	Startlinecolor Color `json:"startlinecolor,omitempty"`

	// Startlinewidth
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the start line.
	Startlinewidth float64 `json:"startlinewidth,omitempty"`

	// Tick0
	// arrayOK: false
	// type: number
	// The starting index of grid lines along the axis
	Tick0 float64 `json:"tick0,omitempty"`

	// Tickangle
	// arrayOK: false
	// type: angle
	// Sets the angle of the tick labels with respect to the horizontal. For example, a `tickangle` of -90 draws the tick labels vertically.
	Tickangle float64 `json:"tickangle,omitempty"`

	// Tickfont
	// role: Object
	Tickfont *CarpetAaxisTickfont `json:"tickfont,omitempty"`

	// Tickformat
	// arrayOK: false
	// type: string
	// Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see:  We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46*
	Tickformat String `json:"tickformat,omitempty"`

	// Tickformatstops
	// It's an items array and what goes inside it's... messy... check the docs
	// I will be happy if you want to contribute by implementing this
	// just raise an issue before you start so we do not overlap
	Tickformatstops interface{} `json:"tickformatstops,omitempty"`

	// Tickmode
	// default: array
	// type: enumerated
	//
	Tickmode CarpetAaxisTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

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

	// Title
	// role: Object
	Title *CarpetAaxisTitle `json:"title,omitempty"`

	// Type
	// default: -
	// type: enumerated
	// Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
	Type CarpetAaxisType `json:"type,omitempty"`
}

// CarpetBaxisTickfont Sets the tick font.
type CarpetBaxisTickfont struct {

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

// CarpetBaxisTitleFont Sets this axis' title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type CarpetBaxisTitleFont struct {

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

// CarpetBaxisTitle
type CarpetBaxisTitle struct {

	// Font
	// role: Object
	Font *CarpetBaxisTitleFont `json:"font,omitempty"`

	// Offset
	// arrayOK: false
	// type: number
	// An additional amount by which to offset the title from the tick labels, given in pixels. Note that this used to be set by the now deprecated `titleoffset` attribute.
	Offset float64 `json:"offset,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of this axis. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// CarpetBaxis
type CarpetBaxis struct {

	// Arraydtick
	// arrayOK: false
	// type: integer
	// The stride between grid lines along the axis
	Arraydtick int64 `json:"arraydtick,omitempty"`

	// Arraytick0
	// arrayOK: false
	// type: integer
	// The starting index of grid lines along the axis
	Arraytick0 int64 `json:"arraytick0,omitempty"`

	// Autorange
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
	Autorange CarpetBaxisAutorange `json:"autorange,omitempty"`

	// Autotypenumbers
	// default: convert types
	// type: enumerated
	// Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers.
	Autotypenumbers CarpetBaxisAutotypenumbers `json:"autotypenumbers,omitempty"`

	// Categoryarray
	// arrayOK: false
	// type: data_array
	// Sets the order in which categories on this axis appear. Only has an effect if `categoryorder` is set to *array*. Used with `categoryorder`.
	Categoryarray interface{} `json:"categoryarray,omitempty"`

	// Categoryarraysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  categoryarray .
	Categoryarraysrc String `json:"categoryarraysrc,omitempty"`

	// Categoryorder
	// default: trace
	// type: enumerated
	// Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`.
	Categoryorder CarpetBaxisCategoryorder `json:"categoryorder,omitempty"`

	// Cheatertype
	// default: value
	// type: enumerated
	//
	Cheatertype CarpetBaxisCheatertype `json:"cheatertype,omitempty"`

	// Color
	// arrayOK: false
	// type: color
	// Sets default for all colors associated with this axis all at once: line, font, tick, and grid colors. Grid color is lightened by blending this with the plot background Individual pieces can override this.
	Color Color `json:"color,omitempty"`

	// Dtick
	// arrayOK: false
	// type: number
	// The stride between grid lines along the axis
	Dtick float64 `json:"dtick,omitempty"`

	// Endline
	// arrayOK: false
	// type: boolean
	// Determines whether or not a line is drawn at along the final value of this axis. If *true*, the end line is drawn on top of the grid lines.
	Endline Bool `json:"endline,omitempty"`

	// Endlinecolor
	// arrayOK: false
	// type: color
	// Sets the line color of the end line.
	Endlinecolor Color `json:"endlinecolor,omitempty"`

	// Endlinewidth
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the end line.
	Endlinewidth float64 `json:"endlinewidth,omitempty"`

	// Exponentformat
	// default: B
	// type: enumerated
	// Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
	Exponentformat CarpetBaxisExponentformat `json:"exponentformat,omitempty"`

	// Fixedrange
	// arrayOK: false
	// type: boolean
	// Determines whether or not this axis is zoom-able. If true, then zoom is disabled.
	Fixedrange Bool `json:"fixedrange,omitempty"`

	// Gridcolor
	// arrayOK: false
	// type: color
	// Sets the axis line color.
	Gridcolor Color `json:"gridcolor,omitempty"`

	// Gridwidth
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the axis line.
	Gridwidth float64 `json:"gridwidth,omitempty"`

	// Labelpadding
	// arrayOK: false
	// type: integer
	// Extra padding between label and the axis
	Labelpadding int64 `json:"labelpadding,omitempty"`

	// Labelprefix
	// arrayOK: false
	// type: string
	// Sets a axis label prefix.
	Labelprefix String `json:"labelprefix,omitempty"`

	// Labelsuffix
	// arrayOK: false
	// type: string
	// Sets a axis label suffix.
	Labelsuffix String `json:"labelsuffix,omitempty"`

	// Linecolor
	// arrayOK: false
	// type: color
	// Sets the axis line color.
	Linecolor Color `json:"linecolor,omitempty"`

	// Linewidth
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the axis line.
	Linewidth float64 `json:"linewidth,omitempty"`

	// Minexponent
	// arrayOK: false
	// type: number
	// Hide SI prefix for 10^n if |n| is below this number
	Minexponent float64 `json:"minexponent,omitempty"`

	// Minorgridcolor
	// arrayOK: false
	// type: color
	// Sets the color of the grid lines.
	Minorgridcolor Color `json:"minorgridcolor,omitempty"`

	// Minorgridcount
	// arrayOK: false
	// type: integer
	// Sets the number of minor grid ticks per major grid tick
	Minorgridcount int64 `json:"minorgridcount,omitempty"`

	// Minorgridwidth
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the grid lines.
	Minorgridwidth float64 `json:"minorgridwidth,omitempty"`

	// Nticks
	// arrayOK: false
	// type: integer
	// Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*.
	Nticks int64 `json:"nticks,omitempty"`

	// Range
	// arrayOK: false
	// type: info_array
	// Sets the range of this axis. If the axis `type` is *log*, then you must take the log of your desired range (e.g. to set the range from 1 to 100, set the range from 0 to 2). If the axis `type` is *date*, it should be date strings, like date data, though Date objects and unix milliseconds will be accepted and converted to strings. If the axis `type` is *category*, it should be numbers, using the scale where each category is assigned a serial number from zero in the order it appears.
	Range interface{} `json:"range,omitempty"`

	// Rangemode
	// default: normal
	// type: enumerated
	// If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data.
	Rangemode CarpetBaxisRangemode `json:"rangemode,omitempty"`

	// Separatethousands
	// arrayOK: false
	// type: boolean
	// If "true", even 4-digit integers are separated
	Separatethousands Bool `json:"separatethousands,omitempty"`

	// Showexponent
	// default: all
	// type: enumerated
	// If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
	Showexponent CarpetBaxisShowexponent `json:"showexponent,omitempty"`

	// Showgrid
	// arrayOK: false
	// type: boolean
	// Determines whether or not grid lines are drawn. If *true*, the grid lines are drawn at every tick mark.
	Showgrid Bool `json:"showgrid,omitempty"`

	// Showline
	// arrayOK: false
	// type: boolean
	// Determines whether or not a line bounding this axis is drawn.
	Showline Bool `json:"showline,omitempty"`

	// Showticklabels
	// default: start
	// type: enumerated
	// Determines whether axis labels are drawn on the low side, the high side, both, or neither side of the axis.
	Showticklabels CarpetBaxisShowticklabels `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix CarpetBaxisShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix CarpetBaxisShowticksuffix `json:"showticksuffix,omitempty"`

	// Smoothing
	// arrayOK: false
	// type: number
	//
	Smoothing float64 `json:"smoothing,omitempty"`

	// Startline
	// arrayOK: false
	// type: boolean
	// Determines whether or not a line is drawn at along the starting value of this axis. If *true*, the start line is drawn on top of the grid lines.
	Startline Bool `json:"startline,omitempty"`

	// Startlinecolor
	// arrayOK: false
	// type: color
	// Sets the line color of the start line.
	Startlinecolor Color `json:"startlinecolor,omitempty"`

	// Startlinewidth
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the start line.
	Startlinewidth float64 `json:"startlinewidth,omitempty"`

	// Tick0
	// arrayOK: false
	// type: number
	// The starting index of grid lines along the axis
	Tick0 float64 `json:"tick0,omitempty"`

	// Tickangle
	// arrayOK: false
	// type: angle
	// Sets the angle of the tick labels with respect to the horizontal. For example, a `tickangle` of -90 draws the tick labels vertically.
	Tickangle float64 `json:"tickangle,omitempty"`

	// Tickfont
	// role: Object
	Tickfont *CarpetBaxisTickfont `json:"tickfont,omitempty"`

	// Tickformat
	// arrayOK: false
	// type: string
	// Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see:  We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46*
	Tickformat String `json:"tickformat,omitempty"`

	// Tickformatstops
	// It's an items array and what goes inside it's... messy... check the docs
	// I will be happy if you want to contribute by implementing this
	// just raise an issue before you start so we do not overlap
	Tickformatstops interface{} `json:"tickformatstops,omitempty"`

	// Tickmode
	// default: array
	// type: enumerated
	//
	Tickmode CarpetBaxisTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

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

	// Title
	// role: Object
	Title *CarpetBaxisTitle `json:"title,omitempty"`

	// Type
	// default: -
	// type: enumerated
	// Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
	Type CarpetBaxisType `json:"type,omitempty"`
}

// CarpetFont The default font used for axis & tick labels on this carpet
type CarpetFont struct {

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

// CarpetStream
type CarpetStream struct {

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

// CarpetAaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type CarpetAaxisAutorange interface{}

var (
	CarpetAaxisAutorangeTrue     CarpetAaxisAutorange = true
	CarpetAaxisAutorangeFalse    CarpetAaxisAutorange = false
	CarpetAaxisAutorangeReversed CarpetAaxisAutorange = "reversed"
)

// CarpetAaxisAutotypenumbers Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers.
type CarpetAaxisAutotypenumbers string

const (
	CarpetAaxisAutotypenumbersConvertTypes CarpetAaxisAutotypenumbers = "convert types"
	CarpetAaxisAutotypenumbersStrict       CarpetAaxisAutotypenumbers = "strict"
)

// CarpetAaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`.
type CarpetAaxisCategoryorder string

const (
	CarpetAaxisCategoryorderTrace              CarpetAaxisCategoryorder = "trace"
	CarpetAaxisCategoryorderCategoryAscending  CarpetAaxisCategoryorder = "category ascending"
	CarpetAaxisCategoryorderCategoryDescending CarpetAaxisCategoryorder = "category descending"
	CarpetAaxisCategoryorderArray              CarpetAaxisCategoryorder = "array"
)

// CarpetAaxisCheatertype
type CarpetAaxisCheatertype string

const (
	CarpetAaxisCheatertypeIndex CarpetAaxisCheatertype = "index"
	CarpetAaxisCheatertypeValue CarpetAaxisCheatertype = "value"
)

// CarpetAaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type CarpetAaxisExponentformat string

const (
	CarpetAaxisExponentformatNone  CarpetAaxisExponentformat = "none"
	CarpetAaxisExponentformatE1    CarpetAaxisExponentformat = "e"
	CarpetAaxisExponentformatE2    CarpetAaxisExponentformat = "E"
	CarpetAaxisExponentformatPower CarpetAaxisExponentformat = "power"
	CarpetAaxisExponentformatSi    CarpetAaxisExponentformat = "SI"
	CarpetAaxisExponentformatB     CarpetAaxisExponentformat = "B"
)

// CarpetAaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data.
type CarpetAaxisRangemode string

const (
	CarpetAaxisRangemodeNormal      CarpetAaxisRangemode = "normal"
	CarpetAaxisRangemodeTozero      CarpetAaxisRangemode = "tozero"
	CarpetAaxisRangemodeNonnegative CarpetAaxisRangemode = "nonnegative"
)

// CarpetAaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type CarpetAaxisShowexponent string

const (
	CarpetAaxisShowexponentAll   CarpetAaxisShowexponent = "all"
	CarpetAaxisShowexponentFirst CarpetAaxisShowexponent = "first"
	CarpetAaxisShowexponentLast  CarpetAaxisShowexponent = "last"
	CarpetAaxisShowexponentNone  CarpetAaxisShowexponent = "none"
)

// CarpetAaxisShowticklabels Determines whether axis labels are drawn on the low side, the high side, both, or neither side of the axis.
type CarpetAaxisShowticklabels string

const (
	CarpetAaxisShowticklabelsStart CarpetAaxisShowticklabels = "start"
	CarpetAaxisShowticklabelsEnd   CarpetAaxisShowticklabels = "end"
	CarpetAaxisShowticklabelsBoth  CarpetAaxisShowticklabels = "both"
	CarpetAaxisShowticklabelsNone  CarpetAaxisShowticklabels = "none"
)

// CarpetAaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type CarpetAaxisShowtickprefix string

const (
	CarpetAaxisShowtickprefixAll   CarpetAaxisShowtickprefix = "all"
	CarpetAaxisShowtickprefixFirst CarpetAaxisShowtickprefix = "first"
	CarpetAaxisShowtickprefixLast  CarpetAaxisShowtickprefix = "last"
	CarpetAaxisShowtickprefixNone  CarpetAaxisShowtickprefix = "none"
)

// CarpetAaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type CarpetAaxisShowticksuffix string

const (
	CarpetAaxisShowticksuffixAll   CarpetAaxisShowticksuffix = "all"
	CarpetAaxisShowticksuffixFirst CarpetAaxisShowticksuffix = "first"
	CarpetAaxisShowticksuffixLast  CarpetAaxisShowticksuffix = "last"
	CarpetAaxisShowticksuffixNone  CarpetAaxisShowticksuffix = "none"
)

// CarpetAaxisTickmode
type CarpetAaxisTickmode string

const (
	CarpetAaxisTickmodeLinear CarpetAaxisTickmode = "linear"
	CarpetAaxisTickmodeArray  CarpetAaxisTickmode = "array"
)

// CarpetAaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type CarpetAaxisType string

const (
	CarpetAaxisTypeHyphenHyphen CarpetAaxisType = "-"
	CarpetAaxisTypeLinear       CarpetAaxisType = "linear"
	CarpetAaxisTypeDate         CarpetAaxisType = "date"
	CarpetAaxisTypeCategory     CarpetAaxisType = "category"
)

// CarpetBaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type CarpetBaxisAutorange interface{}

var (
	CarpetBaxisAutorangeTrue     CarpetBaxisAutorange = true
	CarpetBaxisAutorangeFalse    CarpetBaxisAutorange = false
	CarpetBaxisAutorangeReversed CarpetBaxisAutorange = "reversed"
)

// CarpetBaxisAutotypenumbers Using *strict* a numeric string in trace data is not converted to a number. Using *convert types* a numeric string in trace data may be treated as a number during automatic axis `type` detection. Defaults to layout.autotypenumbers.
type CarpetBaxisAutotypenumbers string

const (
	CarpetBaxisAutotypenumbersConvertTypes CarpetBaxisAutotypenumbers = "convert types"
	CarpetBaxisAutotypenumbersStrict       CarpetBaxisAutotypenumbers = "strict"
)

// CarpetBaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`.
type CarpetBaxisCategoryorder string

const (
	CarpetBaxisCategoryorderTrace              CarpetBaxisCategoryorder = "trace"
	CarpetBaxisCategoryorderCategoryAscending  CarpetBaxisCategoryorder = "category ascending"
	CarpetBaxisCategoryorderCategoryDescending CarpetBaxisCategoryorder = "category descending"
	CarpetBaxisCategoryorderArray              CarpetBaxisCategoryorder = "array"
)

// CarpetBaxisCheatertype
type CarpetBaxisCheatertype string

const (
	CarpetBaxisCheatertypeIndex CarpetBaxisCheatertype = "index"
	CarpetBaxisCheatertypeValue CarpetBaxisCheatertype = "value"
)

// CarpetBaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type CarpetBaxisExponentformat string

const (
	CarpetBaxisExponentformatNone  CarpetBaxisExponentformat = "none"
	CarpetBaxisExponentformatE1    CarpetBaxisExponentformat = "e"
	CarpetBaxisExponentformatE2    CarpetBaxisExponentformat = "E"
	CarpetBaxisExponentformatPower CarpetBaxisExponentformat = "power"
	CarpetBaxisExponentformatSi    CarpetBaxisExponentformat = "SI"
	CarpetBaxisExponentformatB     CarpetBaxisExponentformat = "B"
)

// CarpetBaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data.
type CarpetBaxisRangemode string

const (
	CarpetBaxisRangemodeNormal      CarpetBaxisRangemode = "normal"
	CarpetBaxisRangemodeTozero      CarpetBaxisRangemode = "tozero"
	CarpetBaxisRangemodeNonnegative CarpetBaxisRangemode = "nonnegative"
)

// CarpetBaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type CarpetBaxisShowexponent string

const (
	CarpetBaxisShowexponentAll   CarpetBaxisShowexponent = "all"
	CarpetBaxisShowexponentFirst CarpetBaxisShowexponent = "first"
	CarpetBaxisShowexponentLast  CarpetBaxisShowexponent = "last"
	CarpetBaxisShowexponentNone  CarpetBaxisShowexponent = "none"
)

// CarpetBaxisShowticklabels Determines whether axis labels are drawn on the low side, the high side, both, or neither side of the axis.
type CarpetBaxisShowticklabels string

const (
	CarpetBaxisShowticklabelsStart CarpetBaxisShowticklabels = "start"
	CarpetBaxisShowticklabelsEnd   CarpetBaxisShowticklabels = "end"
	CarpetBaxisShowticklabelsBoth  CarpetBaxisShowticklabels = "both"
	CarpetBaxisShowticklabelsNone  CarpetBaxisShowticklabels = "none"
)

// CarpetBaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type CarpetBaxisShowtickprefix string

const (
	CarpetBaxisShowtickprefixAll   CarpetBaxisShowtickprefix = "all"
	CarpetBaxisShowtickprefixFirst CarpetBaxisShowtickprefix = "first"
	CarpetBaxisShowtickprefixLast  CarpetBaxisShowtickprefix = "last"
	CarpetBaxisShowtickprefixNone  CarpetBaxisShowtickprefix = "none"
)

// CarpetBaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type CarpetBaxisShowticksuffix string

const (
	CarpetBaxisShowticksuffixAll   CarpetBaxisShowticksuffix = "all"
	CarpetBaxisShowticksuffixFirst CarpetBaxisShowticksuffix = "first"
	CarpetBaxisShowticksuffixLast  CarpetBaxisShowticksuffix = "last"
	CarpetBaxisShowticksuffixNone  CarpetBaxisShowticksuffix = "none"
)

// CarpetBaxisTickmode
type CarpetBaxisTickmode string

const (
	CarpetBaxisTickmodeLinear CarpetBaxisTickmode = "linear"
	CarpetBaxisTickmodeArray  CarpetBaxisTickmode = "array"
)

// CarpetBaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type CarpetBaxisType string

const (
	CarpetBaxisTypeHyphenHyphen CarpetBaxisType = "-"
	CarpetBaxisTypeLinear       CarpetBaxisType = "linear"
	CarpetBaxisTypeDate         CarpetBaxisType = "date"
	CarpetBaxisTypeCategory     CarpetBaxisType = "category"
)

// CarpetVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type CarpetVisible interface{}

var (
	CarpetVisibleTrue       CarpetVisible = true
	CarpetVisibleFalse      CarpetVisible = false
	CarpetVisibleLegendonly CarpetVisible = "legendonly"
)
