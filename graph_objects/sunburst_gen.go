package grob

var TraceTypeSunburst TraceType = "sunburst"

func (trace *Sunburst) GetType() TraceType {
	return TraceTypeSunburst
}

// Sunburst Visualize hierarchal data spanning outward radially from root to leaves. The sunburst sectors are determined by the entries in *labels* or *ids* and in *parents*.
type Sunburst struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Branchvalues
	// default: remainder
	// type: enumerated
	// Determines how the items in `values` are summed. When set to *total*, items in `values` are taken to be value of all its descendants. When set to *remainder*, items in `values` corresponding to the root and the branches sectors are taken to be the extra part not part of the sum of the values at their leaves.
	Branchvalues SunburstBranchvalues `json:"branchvalues,omitempty"`

	// Count
	// default: leaves
	// type: flaglist
	// Determines default for `values` when it is not provided, by inferring a 1 for each of the *leaves* and/or *branches*, otherwise 0.
	Count SunburstCount `json:"count,omitempty"`

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

	// Domain
	// role: Object
	Domain *SunburstDomain `json:"domain,omitempty"`

	// Hoverinfo
	// default: label+text+value+name
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo SunburstHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *SunburstHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate
	// arrayOK: true
	// type: string
	// Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `currentPath`, `root`, `entry`, `percentRoot`, `percentEntry` and `percentParent`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate String `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext
	// arrayOK: true
	// type: string
	// Sets hover text elements associated with each sector. If a single string, the same string appears for all data points. If an array of string, the items are mapped in order of this trace's sectors. To be seen, trace `hoverinfo` must contain a *text* flag.
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

	// Insidetextfont
	// role: Object
	Insidetextfont *SunburstInsidetextfont `json:"insidetextfont,omitempty"`

	// Insidetextorientation
	// default: auto
	// type: enumerated
	// Controls the orientation of the text inside chart sectors. When set to *auto*, text may be oriented in any direction in order to be as big as possible in the middle of a sector. The *horizontal* option orients text to be parallel with the bottom of the chart, and may make text smaller in order to achieve that goal. The *radial* option orients text along the radius of the sector. The *tangential* option orients text perpendicular to the radius of the sector.
	Insidetextorientation SunburstInsidetextorientation `json:"insidetextorientation,omitempty"`

	// Labels
	// arrayOK: false
	// type: data_array
	// Sets the labels of each of the sectors.
	Labels interface{} `json:"labels,omitempty"`

	// Labelssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  labels .
	Labelssrc String `json:"labelssrc,omitempty"`

	// Leaf
	// role: Object
	Leaf *SunburstLeaf `json:"leaf,omitempty"`

	// Level
	// arrayOK: false
	// type: any
	// Sets the level from which this trace hierarchy is rendered. Set `level` to `''` to start from the root node in the hierarchy. Must be an "id" if `ids` is filled in, otherwise plotly attempts to find a matching item in `labels`.
	Level interface{} `json:"level,omitempty"`

	// Marker
	// role: Object
	Marker *SunburstMarker `json:"marker,omitempty"`

	// Maxdepth
	// arrayOK: false
	// type: integer
	// Sets the number of rendered sectors from any given `level`. Set `maxdepth` to *-1* to render all the levels in the hierarchy.
	Maxdepth int64 `json:"maxdepth,omitempty"`

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

	// Outsidetextfont
	// role: Object
	Outsidetextfont *SunburstOutsidetextfont `json:"outsidetextfont,omitempty"`

	// Parents
	// arrayOK: false
	// type: data_array
	// Sets the parent sectors for each of the sectors. Empty string items '' are understood to reference the root node in the hierarchy. If `ids` is filled, `parents` items are understood to be "ids" themselves. When `ids` is not set, plotly attempts to find matching items in `labels`, but beware they must be unique.
	Parents interface{} `json:"parents,omitempty"`

	// Parentssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  parents .
	Parentssrc String `json:"parentssrc,omitempty"`

	// Root
	// role: Object
	Root *SunburstRoot `json:"root,omitempty"`

	// Rotation
	// arrayOK: false
	// type: angle
	// Rotates the whole diagram counterclockwise by some angle. By default the first slice starts at 3 o'clock.
	Rotation float64 `json:"rotation,omitempty"`

	// Sort
	// arrayOK: false
	// type: boolean
	// Determines whether or not the sectors are reordered from largest to smallest.
	Sort Bool `json:"sort,omitempty"`

	// Stream
	// role: Object
	Stream *SunburstStream `json:"stream,omitempty"`

	// Text
	// arrayOK: false
	// type: data_array
	// Sets text elements associated with each sector. If trace `textinfo` contains a *text* flag, these elements will be seen on the chart. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont
	// role: Object
	Textfont *SunburstTextfont `json:"textfont,omitempty"`

	// Textinfo
	// default: %!s(<nil>)
	// type: flaglist
	// Determines which trace information appear on the graph.
	Textinfo SunburstTextinfo `json:"textinfo,omitempty"`

	// Textsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate
	// arrayOK: true
	// type: string
	// Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `currentPath`, `root`, `entry`, `percentRoot`, `percentEntry`, `percentParent`, `label` and `value`.
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

	// Values
	// arrayOK: false
	// type: data_array
	// Sets the values associated with each of the sectors. Use with `branchvalues` to determine how the values are summed.
	Values interface{} `json:"values,omitempty"`

	// Valuessrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  values .
	Valuessrc String `json:"valuessrc,omitempty"`

	// Visible
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible SunburstVisible `json:"visible,omitempty"`
}

// SunburstDomain
type SunburstDomain struct {

	// Column
	// arrayOK: false
	// type: integer
	// If there is a layout grid, use the domain for this column in the grid for this sunburst trace .
	Column int64 `json:"column,omitempty"`

	// Row
	// arrayOK: false
	// type: integer
	// If there is a layout grid, use the domain for this row in the grid for this sunburst trace .
	Row int64 `json:"row,omitempty"`

	// X
	// arrayOK: false
	// type: info_array
	// Sets the horizontal domain of this sunburst trace (in plot fraction).
	X interface{} `json:"x,omitempty"`

	// Y
	// arrayOK: false
	// type: info_array
	// Sets the vertical domain of this sunburst trace (in plot fraction).
	Y interface{} `json:"y,omitempty"`
}

// SunburstHoverlabelFont Sets the font used in hover labels.
type SunburstHoverlabelFont struct {

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

// SunburstHoverlabel
type SunburstHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align SunburstHoverlabelAlign `json:"align,omitempty"`

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
	Font *SunburstHoverlabelFont `json:"font,omitempty"`

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

// SunburstInsidetextfont Sets the font used for `textinfo` lying inside the sector.
type SunburstInsidetextfont struct {

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

// SunburstLeaf
type SunburstLeaf struct {

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the opacity of the leaves. With colorscale it is defaulted to 1; otherwise it is defaulted to 0.7
	Opacity float64 `json:"opacity,omitempty"`
}

// SunburstMarkerColorbarTickfont Sets the color bar's tick label font
type SunburstMarkerColorbarTickfont struct {

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

// SunburstMarkerColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type SunburstMarkerColorbarTitleFont struct {

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

// SunburstMarkerColorbarTitle
type SunburstMarkerColorbarTitle struct {

	// Font
	// role: Object
	Font *SunburstMarkerColorbarTitleFont `json:"font,omitempty"`

	// Side
	// default: top
	// type: enumerated
	// Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
	Side SunburstMarkerColorbarTitleSide `json:"side,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// SunburstMarkerColorbar
type SunburstMarkerColorbar struct {

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
	Exponentformat SunburstMarkerColorbarExponentformat `json:"exponentformat,omitempty"`

	// Len
	// arrayOK: false
	// type: number
	// Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends.
	Len float64 `json:"len,omitempty"`

	// Lenmode
	// default: fraction
	// type: enumerated
	// Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
	Lenmode SunburstMarkerColorbarLenmode `json:"lenmode,omitempty"`

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
	Showexponent SunburstMarkerColorbarShowexponent `json:"showexponent,omitempty"`

	// Showticklabels
	// arrayOK: false
	// type: boolean
	// Determines whether or not the tick labels are drawn.
	Showticklabels Bool `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix SunburstMarkerColorbarShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix SunburstMarkerColorbarShowticksuffix `json:"showticksuffix,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels.
	Thickness float64 `json:"thickness,omitempty"`

	// Thicknessmode
	// default: pixels
	// type: enumerated
	// Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
	Thicknessmode SunburstMarkerColorbarThicknessmode `json:"thicknessmode,omitempty"`

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
	Tickfont *SunburstMarkerColorbarTickfont `json:"tickfont,omitempty"`

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
	Ticklabelposition SunburstMarkerColorbarTicklabelposition `json:"ticklabelposition,omitempty"`

	// Ticklen
	// arrayOK: false
	// type: number
	// Sets the tick length (in px).
	Ticklen float64 `json:"ticklen,omitempty"`

	// Tickmode
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
	Tickmode SunburstMarkerColorbarTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

	// Ticks
	// default:
	// type: enumerated
	// Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
	Ticks SunburstMarkerColorbarTicks `json:"ticks,omitempty"`

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
	Title *SunburstMarkerColorbarTitle `json:"title,omitempty"`

	// X
	// arrayOK: false
	// type: number
	// Sets the x position of the color bar (in plot fraction).
	X float64 `json:"x,omitempty"`

	// Xanchor
	// default: left
	// type: enumerated
	// Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
	Xanchor SunburstMarkerColorbarXanchor `json:"xanchor,omitempty"`

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
	Yanchor SunburstMarkerColorbarYanchor `json:"yanchor,omitempty"`

	// Ypad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the y direction.
	Ypad float64 `json:"ypad,omitempty"`
}

// SunburstMarkerLine
type SunburstMarkerLine struct {

	// Color
	// arrayOK: true
	// type: color
	// Sets the color of the line enclosing each sector. Defaults to the `paper_bgcolor` value.
	Color Color `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Width
	// arrayOK: true
	// type: number
	// Sets the width (in px) of the line enclosing each sector.
	Width float64 `json:"width,omitempty"`

	// Widthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  width .
	Widthsrc String `json:"widthsrc,omitempty"`
}

// SunburstMarker
type SunburstMarker struct {

	// Autocolorscale
	// arrayOK: false
	// type: boolean
	// Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `marker.colorscale`. Has an effect only if colorsis set to a numerical array. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Cauto
	// arrayOK: false
	// type: boolean
	// Determines whether or not the color domain is computed with respect to the input data (here colors) or the bounds set in `marker.cmin` and `marker.cmax`  Has an effect only if colorsis set to a numerical array. Defaults to `false` when `marker.cmin` and `marker.cmax` are set by the user.
	Cauto Bool `json:"cauto,omitempty"`

	// Cmax
	// arrayOK: false
	// type: number
	// Sets the upper bound of the color domain. Has an effect only if colorsis set to a numerical array. Value should have the same units as colors and if set, `marker.cmin` must be set as well.
	Cmax float64 `json:"cmax,omitempty"`

	// Cmid
	// arrayOK: false
	// type: number
	// Sets the mid-point of the color domain by scaling `marker.cmin` and/or `marker.cmax` to be equidistant to this point. Has an effect only if colorsis set to a numerical array. Value should have the same units as colors. Has no effect when `marker.cauto` is `false`.
	Cmid float64 `json:"cmid,omitempty"`

	// Cmin
	// arrayOK: false
	// type: number
	// Sets the lower bound of the color domain. Has an effect only if colorsis set to a numerical array. Value should have the same units as colors and if set, `marker.cmax` must be set as well.
	Cmin float64 `json:"cmin,omitempty"`

	// Coloraxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar
	// role: Object
	Colorbar *SunburstMarkerColorbar `json:"colorbar,omitempty"`

	// Colors
	// arrayOK: false
	// type: data_array
	// Sets the color of each sector of this trace. If not specified, the default trace color set is used to pick the sector colors.
	Colors interface{} `json:"colors,omitempty"`

	// Colorscale
	// default: %!s(<nil>)
	// type: colorscale
	// Sets the colorscale. Has an effect only if colorsis set to a numerical array. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`marker.cmin` and `marker.cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	Colorscale ColorScale `json:"colorscale,omitempty"`

	// Colorssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  colors .
	Colorssrc String `json:"colorssrc,omitempty"`

	// Line
	// role: Object
	Line *SunburstMarkerLine `json:"line,omitempty"`

	// Reversescale
	// arrayOK: false
	// type: boolean
	// Reverses the color mapping if true. Has an effect only if colorsis set to a numerical array. If true, `marker.cmin` will correspond to the last color in the array and `marker.cmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Showscale
	// arrayOK: false
	// type: boolean
	// Determines whether or not a colorbar is displayed for this trace. Has an effect only if colorsis set to a numerical array.
	Showscale Bool `json:"showscale,omitempty"`
}

// SunburstOutsidetextfont Sets the font used for `textinfo` lying outside the sector. This option refers to the root of the hierarchy presented at the center of a sunburst graph. Please note that if a hierarchy has multiple root nodes, this option won't have any effect and `insidetextfont` would be used.
type SunburstOutsidetextfont struct {

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

// SunburstRoot
type SunburstRoot struct {

	// Color
	// arrayOK: false
	// type: color
	// sets the color of the root node for a sunburst or a treemap trace. this has no effect when a colorscale is used to set the markers.
	Color Color `json:"color,omitempty"`
}

// SunburstStream
type SunburstStream struct {

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

// SunburstTextfont Sets the font used for `textinfo`.
type SunburstTextfont struct {

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

// SunburstBranchvalues Determines how the items in `values` are summed. When set to *total*, items in `values` are taken to be value of all its descendants. When set to *remainder*, items in `values` corresponding to the root and the branches sectors are taken to be the extra part not part of the sum of the values at their leaves.
type SunburstBranchvalues string

const (
	SunburstBranchvaluesRemainder SunburstBranchvalues = "remainder"
	SunburstBranchvaluesTotal     SunburstBranchvalues = "total"
)

// SunburstHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SunburstHoverlabelAlign string

const (
	SunburstHoverlabelAlignLeft  SunburstHoverlabelAlign = "left"
	SunburstHoverlabelAlignRight SunburstHoverlabelAlign = "right"
	SunburstHoverlabelAlignAuto  SunburstHoverlabelAlign = "auto"
)

// SunburstInsidetextorientation Controls the orientation of the text inside chart sectors. When set to *auto*, text may be oriented in any direction in order to be as big as possible in the middle of a sector. The *horizontal* option orients text to be parallel with the bottom of the chart, and may make text smaller in order to achieve that goal. The *radial* option orients text along the radius of the sector. The *tangential* option orients text perpendicular to the radius of the sector.
type SunburstInsidetextorientation string

const (
	SunburstInsidetextorientationHorizontal SunburstInsidetextorientation = "horizontal"
	SunburstInsidetextorientationRadial     SunburstInsidetextorientation = "radial"
	SunburstInsidetextorientationTangential SunburstInsidetextorientation = "tangential"
	SunburstInsidetextorientationAuto       SunburstInsidetextorientation = "auto"
)

// SunburstMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type SunburstMarkerColorbarExponentformat string

const (
	SunburstMarkerColorbarExponentformatNone  SunburstMarkerColorbarExponentformat = "none"
	SunburstMarkerColorbarExponentformatE1    SunburstMarkerColorbarExponentformat = "e"
	SunburstMarkerColorbarExponentformatE2    SunburstMarkerColorbarExponentformat = "E"
	SunburstMarkerColorbarExponentformatPower SunburstMarkerColorbarExponentformat = "power"
	SunburstMarkerColorbarExponentformatSi    SunburstMarkerColorbarExponentformat = "SI"
	SunburstMarkerColorbarExponentformatB     SunburstMarkerColorbarExponentformat = "B"
)

// SunburstMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type SunburstMarkerColorbarLenmode string

const (
	SunburstMarkerColorbarLenmodeFraction SunburstMarkerColorbarLenmode = "fraction"
	SunburstMarkerColorbarLenmodePixels   SunburstMarkerColorbarLenmode = "pixels"
)

// SunburstMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type SunburstMarkerColorbarShowexponent string

const (
	SunburstMarkerColorbarShowexponentAll   SunburstMarkerColorbarShowexponent = "all"
	SunburstMarkerColorbarShowexponentFirst SunburstMarkerColorbarShowexponent = "first"
	SunburstMarkerColorbarShowexponentLast  SunburstMarkerColorbarShowexponent = "last"
	SunburstMarkerColorbarShowexponentNone  SunburstMarkerColorbarShowexponent = "none"
)

// SunburstMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type SunburstMarkerColorbarShowtickprefix string

const (
	SunburstMarkerColorbarShowtickprefixAll   SunburstMarkerColorbarShowtickprefix = "all"
	SunburstMarkerColorbarShowtickprefixFirst SunburstMarkerColorbarShowtickprefix = "first"
	SunburstMarkerColorbarShowtickprefixLast  SunburstMarkerColorbarShowtickprefix = "last"
	SunburstMarkerColorbarShowtickprefixNone  SunburstMarkerColorbarShowtickprefix = "none"
)

// SunburstMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type SunburstMarkerColorbarShowticksuffix string

const (
	SunburstMarkerColorbarShowticksuffixAll   SunburstMarkerColorbarShowticksuffix = "all"
	SunburstMarkerColorbarShowticksuffixFirst SunburstMarkerColorbarShowticksuffix = "first"
	SunburstMarkerColorbarShowticksuffixLast  SunburstMarkerColorbarShowticksuffix = "last"
	SunburstMarkerColorbarShowticksuffixNone  SunburstMarkerColorbarShowticksuffix = "none"
)

// SunburstMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type SunburstMarkerColorbarThicknessmode string

const (
	SunburstMarkerColorbarThicknessmodeFraction SunburstMarkerColorbarThicknessmode = "fraction"
	SunburstMarkerColorbarThicknessmodePixels   SunburstMarkerColorbarThicknessmode = "pixels"
)

// SunburstMarkerColorbarTicklabelposition Determines where tick labels are drawn.
type SunburstMarkerColorbarTicklabelposition string

const (
	SunburstMarkerColorbarTicklabelpositionOutside       SunburstMarkerColorbarTicklabelposition = "outside"
	SunburstMarkerColorbarTicklabelpositionInside        SunburstMarkerColorbarTicklabelposition = "inside"
	SunburstMarkerColorbarTicklabelpositionOutsideTop    SunburstMarkerColorbarTicklabelposition = "outside top"
	SunburstMarkerColorbarTicklabelpositionInsideTop     SunburstMarkerColorbarTicklabelposition = "inside top"
	SunburstMarkerColorbarTicklabelpositionOutsideBottom SunburstMarkerColorbarTicklabelposition = "outside bottom"
	SunburstMarkerColorbarTicklabelpositionInsideBottom  SunburstMarkerColorbarTicklabelposition = "inside bottom"
)

// SunburstMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type SunburstMarkerColorbarTickmode string

const (
	SunburstMarkerColorbarTickmodeAuto   SunburstMarkerColorbarTickmode = "auto"
	SunburstMarkerColorbarTickmodeLinear SunburstMarkerColorbarTickmode = "linear"
	SunburstMarkerColorbarTickmodeArray  SunburstMarkerColorbarTickmode = "array"
)

// SunburstMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type SunburstMarkerColorbarTicks string

const (
	SunburstMarkerColorbarTicksOutside SunburstMarkerColorbarTicks = "outside"
	SunburstMarkerColorbarTicksInside  SunburstMarkerColorbarTicks = "inside"
	SunburstMarkerColorbarTicksEmpty   SunburstMarkerColorbarTicks = ""
)

// SunburstMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type SunburstMarkerColorbarTitleSide string

const (
	SunburstMarkerColorbarTitleSideRight  SunburstMarkerColorbarTitleSide = "right"
	SunburstMarkerColorbarTitleSideTop    SunburstMarkerColorbarTitleSide = "top"
	SunburstMarkerColorbarTitleSideBottom SunburstMarkerColorbarTitleSide = "bottom"
)

// SunburstMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type SunburstMarkerColorbarXanchor string

const (
	SunburstMarkerColorbarXanchorLeft   SunburstMarkerColorbarXanchor = "left"
	SunburstMarkerColorbarXanchorCenter SunburstMarkerColorbarXanchor = "center"
	SunburstMarkerColorbarXanchorRight  SunburstMarkerColorbarXanchor = "right"
)

// SunburstMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type SunburstMarkerColorbarYanchor string

const (
	SunburstMarkerColorbarYanchorTop    SunburstMarkerColorbarYanchor = "top"
	SunburstMarkerColorbarYanchorMiddle SunburstMarkerColorbarYanchor = "middle"
	SunburstMarkerColorbarYanchorBottom SunburstMarkerColorbarYanchor = "bottom"
)

// SunburstVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type SunburstVisible interface{}

var (
	SunburstVisibleTrue       SunburstVisible = true
	SunburstVisibleFalse      SunburstVisible = false
	SunburstVisibleLegendonly SunburstVisible = "legendonly"
)

// SunburstCount Determines default for `values` when it is not provided, by inferring a 1 for each of the *leaves* and/or *branches*, otherwise 0.
type SunburstCount string

const (
	// Flags
	SunburstCountBranches SunburstCount = "branches"
	SunburstCountLeaves   SunburstCount = "leaves"

	// Extra

)

// SunburstHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type SunburstHoverinfo string

const (
	// Flags
	SunburstHoverinfoLabel         SunburstHoverinfo = "label"
	SunburstHoverinfoText          SunburstHoverinfo = "text"
	SunburstHoverinfoValue         SunburstHoverinfo = "value"
	SunburstHoverinfoName          SunburstHoverinfo = "name"
	SunburstHoverinfoCurrentPath   SunburstHoverinfo = "current path"
	SunburstHoverinfoPercentRoot   SunburstHoverinfo = "percent root"
	SunburstHoverinfoPercentEntry  SunburstHoverinfo = "percent entry"
	SunburstHoverinfoPercentParent SunburstHoverinfo = "percent parent"

	// Extra
	SunburstHoverinfoAll  SunburstHoverinfo = "all"
	SunburstHoverinfoNone SunburstHoverinfo = "none"
	SunburstHoverinfoSkip SunburstHoverinfo = "skip"
)

// SunburstTextinfo Determines which trace information appear on the graph.
type SunburstTextinfo string

const (
	// Flags
	SunburstTextinfoLabel         SunburstTextinfo = "label"
	SunburstTextinfoText          SunburstTextinfo = "text"
	SunburstTextinfoValue         SunburstTextinfo = "value"
	SunburstTextinfoCurrentPath   SunburstTextinfo = "current path"
	SunburstTextinfoPercentRoot   SunburstTextinfo = "percent root"
	SunburstTextinfoPercentEntry  SunburstTextinfo = "percent entry"
	SunburstTextinfoPercentParent SunburstTextinfo = "percent parent"

	// Extra
	SunburstTextinfoNone SunburstTextinfo = "none"
)
