package grob

var TraceTypeViolin TraceType = "violin"

func (trace *Violin) GetType() TraceType {
	return TraceTypeViolin
}

// Violin In vertical (horizontal) violin plots, statistics are computed using `y` (`x`) values. By supplying an `x` (`y`) array, one violin per distinct x (y) value is drawn If no `x` (`y`) {array} is provided, a single violin is drawn. That violin position is then positioned with with `name` or with `x0` (`y0`) if provided.
type Violin struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Alignmentgroup
	// arrayOK: false
	// type: string
	// Set several traces linked to the same position axis or matching axes to the same alignmentgroup. This controls whether bars compute their positional range dependently or independently.
	Alignmentgroup String `json:"alignmentgroup,omitempty"`

	// Bandwidth
	// arrayOK: false
	// type: number
	// Sets the bandwidth used to compute the kernel density estimate. By default, the bandwidth is determined by Silverman's rule of thumb.
	Bandwidth float64 `json:"bandwidth,omitempty"`

	// Box
	// role: Object
	Box *ViolinBox `json:"box,omitempty"`

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

	// Fillcolor
	// arrayOK: false
	// type: color
	// Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor Color `json:"fillcolor,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ViolinHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *ViolinHoverlabel `json:"hoverlabel,omitempty"`

	// Hoveron
	// default: violins+points+kde
	// type: flaglist
	// Do the hover effects highlight individual violins or sample points or the kernel density estimate or any combination of them?
	Hoveron ViolinHoveron `json:"hoveron,omitempty"`

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

	// Jitter
	// arrayOK: false
	// type: number
	// Sets the amount of jitter in the sample points drawn. If *0*, the sample points align along the distribution axis. If *1*, the sample points are drawn in a random jitter of width equal to the width of the violins.
	Jitter float64 `json:"jitter,omitempty"`

	// Legendgroup
	// arrayOK: false
	// type: string
	// Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line
	// role: Object
	Line *ViolinLine `json:"line,omitempty"`

	// Marker
	// role: Object
	Marker *ViolinMarker `json:"marker,omitempty"`

	// Meanline
	// role: Object
	Meanline *ViolinMeanline `json:"meanline,omitempty"`

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
	// Sets the trace name. The trace name appear as the legend item and on hover. For violin traces, the name will also be used for the position coordinate, if `x` and `x0` (`y` and `y0` if horizontal) are missing and the position axis is categorical. Note that the trace name is also used as a default value for attribute `scalegroup` (please see its description for details).
	Name String `json:"name,omitempty"`

	// Offsetgroup
	// arrayOK: false
	// type: string
	// Set several traces linked to the same position axis or matching axes to the same offsetgroup where bars of the same position coordinate will line up.
	Offsetgroup String `json:"offsetgroup,omitempty"`

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Orientation
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the orientation of the violin(s). If *v* (*h*), the distribution is visualized along the vertical (horizontal).
	Orientation ViolinOrientation `json:"orientation,omitempty"`

	// Pointpos
	// arrayOK: false
	// type: number
	// Sets the position of the sample points in relation to the violins. If *0*, the sample points are places over the center of the violins. Positive (negative) values correspond to positions to the right (left) for vertical violins and above (below) for horizontal violins.
	Pointpos float64 `json:"pointpos,omitempty"`

	// Points
	// default: %!s(<nil>)
	// type: enumerated
	// If *outliers*, only the sample points lying outside the whiskers are shown If *suspectedoutliers*, the outlier points are shown and points either less than 4*Q1-3*Q3 or greater than 4*Q3-3*Q1 are highlighted (see `outliercolor`) If *all*, all sample points are shown If *false*, only the violins are shown with no sample points. Defaults to *suspectedoutliers* when `marker.outliercolor` or `marker.line.outliercolor` is set, otherwise defaults to *outliers*.
	Points ViolinPoints `json:"points,omitempty"`

	// Scalegroup
	// arrayOK: false
	// type: string
	// If there are multiple violins that should be sized according to to some metric (see `scalemode`), link them by providing a non-empty group id here shared by every trace in the same group. If a violin's `width` is undefined, `scalegroup` will default to the trace's name. In this case, violins with the same names will be linked together
	Scalegroup String `json:"scalegroup,omitempty"`

	// Scalemode
	// default: width
	// type: enumerated
	// Sets the metric by which the width of each violin is determined.*width* means each violin has the same (max) width*count* means the violins are scaled by the number of sample points makingup each violin.
	Scalemode ViolinScalemode `json:"scalemode,omitempty"`

	// Selected
	// role: Object
	Selected *ViolinSelected `json:"selected,omitempty"`

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

	// Side
	// default: both
	// type: enumerated
	// Determines on which side of the position value the density function making up one half of a violin is plotted. Useful when comparing two violin traces under *overlay* mode, where one trace has `side` set to *positive* and the other to *negative*.
	Side ViolinSide `json:"side,omitempty"`

	// Span
	// arrayOK: false
	// type: info_array
	// Sets the span in data space for which the density function will be computed. Has an effect only when `spanmode` is set to *manual*.
	Span interface{} `json:"span,omitempty"`

	// Spanmode
	// default: soft
	// type: enumerated
	// Sets the method by which the span in data space where the density function will be computed. *soft* means the span goes from the sample's minimum value minus two bandwidths to the sample's maximum value plus two bandwidths. *hard* means the span goes from the sample's minimum to its maximum value. For custom span settings, use mode *manual* and fill in the `span` attribute.
	Spanmode ViolinSpanmode `json:"spanmode,omitempty"`

	// Stream
	// role: Object
	Stream *ViolinStream `json:"stream,omitempty"`

	// Text
	// arrayOK: true
	// type: string
	// Sets the text elements associated with each sample value. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
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
	Unselected *ViolinUnselected `json:"unselected,omitempty"`

	// Visible
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ViolinVisible `json:"visible,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the width of the violin in data coordinates. If *0* (default value) the width is automatically selected based on the positions of other violin traces in the same subplot.
	Width float64 `json:"width,omitempty"`

	// X
	// arrayOK: false
	// type: data_array
	// Sets the x sample data or coordinates. See overview for more info.
	X interface{} `json:"x,omitempty"`

	// X0
	// arrayOK: false
	// type: any
	// Sets the x coordinate for single-box traces or the starting coordinate for multi-box traces set using q1/median/q3. See overview for more info.
	X0 interface{} `json:"x0,omitempty"`

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
	// Sets the y sample data or coordinates. See overview for more info.
	Y interface{} `json:"y,omitempty"`

	// Y0
	// arrayOK: false
	// type: any
	// Sets the y coordinate for single-box traces or the starting coordinate for multi-box traces set using q1/median/q3. See overview for more info.
	Y0 interface{} `json:"y0,omitempty"`

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

// ViolinBoxLine
type ViolinBoxLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the inner box plot bounding line color.
	Color Color `json:"color,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the inner box plot bounding line width.
	Width float64 `json:"width,omitempty"`
}

// ViolinBox
type ViolinBox struct {

	// Fillcolor
	// arrayOK: false
	// type: color
	// Sets the inner box plot fill color.
	Fillcolor Color `json:"fillcolor,omitempty"`

	// Line
	// role: Object
	Line *ViolinBoxLine `json:"line,omitempty"`

	// Visible
	// arrayOK: false
	// type: boolean
	// Determines if an miniature box plot is drawn inside the violins.
	Visible Bool `json:"visible,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the width of the inner box plots relative to the violins' width. For example, with 1, the inner box plots are as wide as the violins.
	Width float64 `json:"width,omitempty"`
}

// ViolinHoverlabelFont Sets the font used in hover labels.
type ViolinHoverlabelFont struct {

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

// ViolinHoverlabel
type ViolinHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align ViolinHoverlabelAlign `json:"align,omitempty"`

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
	Font *ViolinHoverlabelFont `json:"font,omitempty"`

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

// ViolinLine
type ViolinLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the color of line bounding the violin(s).
	Color Color `json:"color,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the width (in px) of line bounding the violin(s).
	Width float64 `json:"width,omitempty"`
}

// ViolinMarkerLine
type ViolinMarkerLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets themarker.linecolor. It accepts either a specific color or an array of numbers that are mapped to the colorscale relative to the max and min values of the array or relative to `marker.line.cmin` and `marker.line.cmax` if set.
	Color Color `json:"color,omitempty"`

	// Outliercolor
	// arrayOK: false
	// type: color
	// Sets the border line color of the outlier sample points. Defaults to marker.color
	Outliercolor Color `json:"outliercolor,omitempty"`

	// Outlierwidth
	// arrayOK: false
	// type: number
	// Sets the border line width (in px) of the outlier sample points.
	Outlierwidth float64 `json:"outlierwidth,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the width (in px) of the lines bounding the marker points.
	Width float64 `json:"width,omitempty"`
}

// ViolinMarker
type ViolinMarker struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets themarkercolor. It accepts either a specific color or an array of numbers that are mapped to the colorscale relative to the max and min values of the array or relative to `marker.cmin` and `marker.cmax` if set.
	Color Color `json:"color,omitempty"`

	// Line
	// role: Object
	Line *ViolinMarkerLine `json:"line,omitempty"`

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the marker opacity.
	Opacity float64 `json:"opacity,omitempty"`

	// Outliercolor
	// arrayOK: false
	// type: color
	// Sets the color of the outlier sample points.
	Outliercolor Color `json:"outliercolor,omitempty"`

	// Size
	// arrayOK: false
	// type: number
	// Sets the marker size (in px).
	Size float64 `json:"size,omitempty"`

	// Symbol
	// default: circle
	// type: enumerated
	// Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
	Symbol ViolinMarkerSymbol `json:"symbol,omitempty"`
}

// ViolinMeanline
type ViolinMeanline struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the mean line color.
	Color Color `json:"color,omitempty"`

	// Visible
	// arrayOK: false
	// type: boolean
	// Determines if a line corresponding to the sample's mean is shown inside the violins. If `box.visible` is turned on, the mean line is drawn inside the inner box. Otherwise, the mean line is drawn from one side of the violin to other.
	Visible Bool `json:"visible,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the mean line width.
	Width float64 `json:"width,omitempty"`
}

// ViolinSelectedMarker
type ViolinSelectedMarker struct {

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

// ViolinSelected
type ViolinSelected struct {

	// Marker
	// role: Object
	Marker *ViolinSelectedMarker `json:"marker,omitempty"`
}

// ViolinStream
type ViolinStream struct {

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

// ViolinUnselectedMarker
type ViolinUnselectedMarker struct {

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

// ViolinUnselected
type ViolinUnselected struct {

	// Marker
	// role: Object
	Marker *ViolinUnselectedMarker `json:"marker,omitempty"`
}

// ViolinHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ViolinHoverlabelAlign string

const (
	ViolinHoverlabelAlignLeft  ViolinHoverlabelAlign = "left"
	ViolinHoverlabelAlignRight ViolinHoverlabelAlign = "right"
	ViolinHoverlabelAlignAuto  ViolinHoverlabelAlign = "auto"
)

// ViolinMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ViolinMarkerSymbol interface{}

var (
	ViolinMarkerSymbolNumber0                 ViolinMarkerSymbol = 0
	ViolinMarkerSymbol0                       ViolinMarkerSymbol = "0"
	ViolinMarkerSymbolCircle                  ViolinMarkerSymbol = "circle"
	ViolinMarkerSymbolNumber100               ViolinMarkerSymbol = 100
	ViolinMarkerSymbol100                     ViolinMarkerSymbol = "100"
	ViolinMarkerSymbolCircleOpen              ViolinMarkerSymbol = "circle-open"
	ViolinMarkerSymbolNumber200               ViolinMarkerSymbol = 200
	ViolinMarkerSymbol200                     ViolinMarkerSymbol = "200"
	ViolinMarkerSymbolCircleDot               ViolinMarkerSymbol = "circle-dot"
	ViolinMarkerSymbolNumber300               ViolinMarkerSymbol = 300
	ViolinMarkerSymbol300                     ViolinMarkerSymbol = "300"
	ViolinMarkerSymbolCircleOpenDot           ViolinMarkerSymbol = "circle-open-dot"
	ViolinMarkerSymbolNumber1                 ViolinMarkerSymbol = 1
	ViolinMarkerSymbol1                       ViolinMarkerSymbol = "1"
	ViolinMarkerSymbolSquare                  ViolinMarkerSymbol = "square"
	ViolinMarkerSymbolNumber101               ViolinMarkerSymbol = 101
	ViolinMarkerSymbol101                     ViolinMarkerSymbol = "101"
	ViolinMarkerSymbolSquareOpen              ViolinMarkerSymbol = "square-open"
	ViolinMarkerSymbolNumber201               ViolinMarkerSymbol = 201
	ViolinMarkerSymbol201                     ViolinMarkerSymbol = "201"
	ViolinMarkerSymbolSquareDot               ViolinMarkerSymbol = "square-dot"
	ViolinMarkerSymbolNumber301               ViolinMarkerSymbol = 301
	ViolinMarkerSymbol301                     ViolinMarkerSymbol = "301"
	ViolinMarkerSymbolSquareOpenDot           ViolinMarkerSymbol = "square-open-dot"
	ViolinMarkerSymbolNumber2                 ViolinMarkerSymbol = 2
	ViolinMarkerSymbol2                       ViolinMarkerSymbol = "2"
	ViolinMarkerSymbolDiamond                 ViolinMarkerSymbol = "diamond"
	ViolinMarkerSymbolNumber102               ViolinMarkerSymbol = 102
	ViolinMarkerSymbol102                     ViolinMarkerSymbol = "102"
	ViolinMarkerSymbolDiamondOpen             ViolinMarkerSymbol = "diamond-open"
	ViolinMarkerSymbolNumber202               ViolinMarkerSymbol = 202
	ViolinMarkerSymbol202                     ViolinMarkerSymbol = "202"
	ViolinMarkerSymbolDiamondDot              ViolinMarkerSymbol = "diamond-dot"
	ViolinMarkerSymbolNumber302               ViolinMarkerSymbol = 302
	ViolinMarkerSymbol302                     ViolinMarkerSymbol = "302"
	ViolinMarkerSymbolDiamondOpenDot          ViolinMarkerSymbol = "diamond-open-dot"
	ViolinMarkerSymbolNumber3                 ViolinMarkerSymbol = 3
	ViolinMarkerSymbol3                       ViolinMarkerSymbol = "3"
	ViolinMarkerSymbolCross                   ViolinMarkerSymbol = "cross"
	ViolinMarkerSymbolNumber103               ViolinMarkerSymbol = 103
	ViolinMarkerSymbol103                     ViolinMarkerSymbol = "103"
	ViolinMarkerSymbolCrossOpen               ViolinMarkerSymbol = "cross-open"
	ViolinMarkerSymbolNumber203               ViolinMarkerSymbol = 203
	ViolinMarkerSymbol203                     ViolinMarkerSymbol = "203"
	ViolinMarkerSymbolCrossDot                ViolinMarkerSymbol = "cross-dot"
	ViolinMarkerSymbolNumber303               ViolinMarkerSymbol = 303
	ViolinMarkerSymbol303                     ViolinMarkerSymbol = "303"
	ViolinMarkerSymbolCrossOpenDot            ViolinMarkerSymbol = "cross-open-dot"
	ViolinMarkerSymbolNumber4                 ViolinMarkerSymbol = 4
	ViolinMarkerSymbol4                       ViolinMarkerSymbol = "4"
	ViolinMarkerSymbolX                       ViolinMarkerSymbol = "x"
	ViolinMarkerSymbolNumber104               ViolinMarkerSymbol = 104
	ViolinMarkerSymbol104                     ViolinMarkerSymbol = "104"
	ViolinMarkerSymbolXOpen                   ViolinMarkerSymbol = "x-open"
	ViolinMarkerSymbolNumber204               ViolinMarkerSymbol = 204
	ViolinMarkerSymbol204                     ViolinMarkerSymbol = "204"
	ViolinMarkerSymbolXDot                    ViolinMarkerSymbol = "x-dot"
	ViolinMarkerSymbolNumber304               ViolinMarkerSymbol = 304
	ViolinMarkerSymbol304                     ViolinMarkerSymbol = "304"
	ViolinMarkerSymbolXOpenDot                ViolinMarkerSymbol = "x-open-dot"
	ViolinMarkerSymbolNumber5                 ViolinMarkerSymbol = 5
	ViolinMarkerSymbol5                       ViolinMarkerSymbol = "5"
	ViolinMarkerSymbolTriangleUp              ViolinMarkerSymbol = "triangle-up"
	ViolinMarkerSymbolNumber105               ViolinMarkerSymbol = 105
	ViolinMarkerSymbol105                     ViolinMarkerSymbol = "105"
	ViolinMarkerSymbolTriangleUpOpen          ViolinMarkerSymbol = "triangle-up-open"
	ViolinMarkerSymbolNumber205               ViolinMarkerSymbol = 205
	ViolinMarkerSymbol205                     ViolinMarkerSymbol = "205"
	ViolinMarkerSymbolTriangleUpDot           ViolinMarkerSymbol = "triangle-up-dot"
	ViolinMarkerSymbolNumber305               ViolinMarkerSymbol = 305
	ViolinMarkerSymbol305                     ViolinMarkerSymbol = "305"
	ViolinMarkerSymbolTriangleUpOpenDot       ViolinMarkerSymbol = "triangle-up-open-dot"
	ViolinMarkerSymbolNumber6                 ViolinMarkerSymbol = 6
	ViolinMarkerSymbol6                       ViolinMarkerSymbol = "6"
	ViolinMarkerSymbolTriangleDown            ViolinMarkerSymbol = "triangle-down"
	ViolinMarkerSymbolNumber106               ViolinMarkerSymbol = 106
	ViolinMarkerSymbol106                     ViolinMarkerSymbol = "106"
	ViolinMarkerSymbolTriangleDownOpen        ViolinMarkerSymbol = "triangle-down-open"
	ViolinMarkerSymbolNumber206               ViolinMarkerSymbol = 206
	ViolinMarkerSymbol206                     ViolinMarkerSymbol = "206"
	ViolinMarkerSymbolTriangleDownDot         ViolinMarkerSymbol = "triangle-down-dot"
	ViolinMarkerSymbolNumber306               ViolinMarkerSymbol = 306
	ViolinMarkerSymbol306                     ViolinMarkerSymbol = "306"
	ViolinMarkerSymbolTriangleDownOpenDot     ViolinMarkerSymbol = "triangle-down-open-dot"
	ViolinMarkerSymbolNumber7                 ViolinMarkerSymbol = 7
	ViolinMarkerSymbol7                       ViolinMarkerSymbol = "7"
	ViolinMarkerSymbolTriangleLeft            ViolinMarkerSymbol = "triangle-left"
	ViolinMarkerSymbolNumber107               ViolinMarkerSymbol = 107
	ViolinMarkerSymbol107                     ViolinMarkerSymbol = "107"
	ViolinMarkerSymbolTriangleLeftOpen        ViolinMarkerSymbol = "triangle-left-open"
	ViolinMarkerSymbolNumber207               ViolinMarkerSymbol = 207
	ViolinMarkerSymbol207                     ViolinMarkerSymbol = "207"
	ViolinMarkerSymbolTriangleLeftDot         ViolinMarkerSymbol = "triangle-left-dot"
	ViolinMarkerSymbolNumber307               ViolinMarkerSymbol = 307
	ViolinMarkerSymbol307                     ViolinMarkerSymbol = "307"
	ViolinMarkerSymbolTriangleLeftOpenDot     ViolinMarkerSymbol = "triangle-left-open-dot"
	ViolinMarkerSymbolNumber8                 ViolinMarkerSymbol = 8
	ViolinMarkerSymbol8                       ViolinMarkerSymbol = "8"
	ViolinMarkerSymbolTriangleRight           ViolinMarkerSymbol = "triangle-right"
	ViolinMarkerSymbolNumber108               ViolinMarkerSymbol = 108
	ViolinMarkerSymbol108                     ViolinMarkerSymbol = "108"
	ViolinMarkerSymbolTriangleRightOpen       ViolinMarkerSymbol = "triangle-right-open"
	ViolinMarkerSymbolNumber208               ViolinMarkerSymbol = 208
	ViolinMarkerSymbol208                     ViolinMarkerSymbol = "208"
	ViolinMarkerSymbolTriangleRightDot        ViolinMarkerSymbol = "triangle-right-dot"
	ViolinMarkerSymbolNumber308               ViolinMarkerSymbol = 308
	ViolinMarkerSymbol308                     ViolinMarkerSymbol = "308"
	ViolinMarkerSymbolTriangleRightOpenDot    ViolinMarkerSymbol = "triangle-right-open-dot"
	ViolinMarkerSymbolNumber9                 ViolinMarkerSymbol = 9
	ViolinMarkerSymbol9                       ViolinMarkerSymbol = "9"
	ViolinMarkerSymbolTriangleNe              ViolinMarkerSymbol = "triangle-ne"
	ViolinMarkerSymbolNumber109               ViolinMarkerSymbol = 109
	ViolinMarkerSymbol109                     ViolinMarkerSymbol = "109"
	ViolinMarkerSymbolTriangleNeOpen          ViolinMarkerSymbol = "triangle-ne-open"
	ViolinMarkerSymbolNumber209               ViolinMarkerSymbol = 209
	ViolinMarkerSymbol209                     ViolinMarkerSymbol = "209"
	ViolinMarkerSymbolTriangleNeDot           ViolinMarkerSymbol = "triangle-ne-dot"
	ViolinMarkerSymbolNumber309               ViolinMarkerSymbol = 309
	ViolinMarkerSymbol309                     ViolinMarkerSymbol = "309"
	ViolinMarkerSymbolTriangleNeOpenDot       ViolinMarkerSymbol = "triangle-ne-open-dot"
	ViolinMarkerSymbolNumber10                ViolinMarkerSymbol = 10
	ViolinMarkerSymbol10                      ViolinMarkerSymbol = "10"
	ViolinMarkerSymbolTriangleSe              ViolinMarkerSymbol = "triangle-se"
	ViolinMarkerSymbolNumber110               ViolinMarkerSymbol = 110
	ViolinMarkerSymbol110                     ViolinMarkerSymbol = "110"
	ViolinMarkerSymbolTriangleSeOpen          ViolinMarkerSymbol = "triangle-se-open"
	ViolinMarkerSymbolNumber210               ViolinMarkerSymbol = 210
	ViolinMarkerSymbol210                     ViolinMarkerSymbol = "210"
	ViolinMarkerSymbolTriangleSeDot           ViolinMarkerSymbol = "triangle-se-dot"
	ViolinMarkerSymbolNumber310               ViolinMarkerSymbol = 310
	ViolinMarkerSymbol310                     ViolinMarkerSymbol = "310"
	ViolinMarkerSymbolTriangleSeOpenDot       ViolinMarkerSymbol = "triangle-se-open-dot"
	ViolinMarkerSymbolNumber11                ViolinMarkerSymbol = 11
	ViolinMarkerSymbol11                      ViolinMarkerSymbol = "11"
	ViolinMarkerSymbolTriangleSw              ViolinMarkerSymbol = "triangle-sw"
	ViolinMarkerSymbolNumber111               ViolinMarkerSymbol = 111
	ViolinMarkerSymbol111                     ViolinMarkerSymbol = "111"
	ViolinMarkerSymbolTriangleSwOpen          ViolinMarkerSymbol = "triangle-sw-open"
	ViolinMarkerSymbolNumber211               ViolinMarkerSymbol = 211
	ViolinMarkerSymbol211                     ViolinMarkerSymbol = "211"
	ViolinMarkerSymbolTriangleSwDot           ViolinMarkerSymbol = "triangle-sw-dot"
	ViolinMarkerSymbolNumber311               ViolinMarkerSymbol = 311
	ViolinMarkerSymbol311                     ViolinMarkerSymbol = "311"
	ViolinMarkerSymbolTriangleSwOpenDot       ViolinMarkerSymbol = "triangle-sw-open-dot"
	ViolinMarkerSymbolNumber12                ViolinMarkerSymbol = 12
	ViolinMarkerSymbol12                      ViolinMarkerSymbol = "12"
	ViolinMarkerSymbolTriangleNw              ViolinMarkerSymbol = "triangle-nw"
	ViolinMarkerSymbolNumber112               ViolinMarkerSymbol = 112
	ViolinMarkerSymbol112                     ViolinMarkerSymbol = "112"
	ViolinMarkerSymbolTriangleNwOpen          ViolinMarkerSymbol = "triangle-nw-open"
	ViolinMarkerSymbolNumber212               ViolinMarkerSymbol = 212
	ViolinMarkerSymbol212                     ViolinMarkerSymbol = "212"
	ViolinMarkerSymbolTriangleNwDot           ViolinMarkerSymbol = "triangle-nw-dot"
	ViolinMarkerSymbolNumber312               ViolinMarkerSymbol = 312
	ViolinMarkerSymbol312                     ViolinMarkerSymbol = "312"
	ViolinMarkerSymbolTriangleNwOpenDot       ViolinMarkerSymbol = "triangle-nw-open-dot"
	ViolinMarkerSymbolNumber13                ViolinMarkerSymbol = 13
	ViolinMarkerSymbol13                      ViolinMarkerSymbol = "13"
	ViolinMarkerSymbolPentagon                ViolinMarkerSymbol = "pentagon"
	ViolinMarkerSymbolNumber113               ViolinMarkerSymbol = 113
	ViolinMarkerSymbol113                     ViolinMarkerSymbol = "113"
	ViolinMarkerSymbolPentagonOpen            ViolinMarkerSymbol = "pentagon-open"
	ViolinMarkerSymbolNumber213               ViolinMarkerSymbol = 213
	ViolinMarkerSymbol213                     ViolinMarkerSymbol = "213"
	ViolinMarkerSymbolPentagonDot             ViolinMarkerSymbol = "pentagon-dot"
	ViolinMarkerSymbolNumber313               ViolinMarkerSymbol = 313
	ViolinMarkerSymbol313                     ViolinMarkerSymbol = "313"
	ViolinMarkerSymbolPentagonOpenDot         ViolinMarkerSymbol = "pentagon-open-dot"
	ViolinMarkerSymbolNumber14                ViolinMarkerSymbol = 14
	ViolinMarkerSymbol14                      ViolinMarkerSymbol = "14"
	ViolinMarkerSymbolHexagon                 ViolinMarkerSymbol = "hexagon"
	ViolinMarkerSymbolNumber114               ViolinMarkerSymbol = 114
	ViolinMarkerSymbol114                     ViolinMarkerSymbol = "114"
	ViolinMarkerSymbolHexagonOpen             ViolinMarkerSymbol = "hexagon-open"
	ViolinMarkerSymbolNumber214               ViolinMarkerSymbol = 214
	ViolinMarkerSymbol214                     ViolinMarkerSymbol = "214"
	ViolinMarkerSymbolHexagonDot              ViolinMarkerSymbol = "hexagon-dot"
	ViolinMarkerSymbolNumber314               ViolinMarkerSymbol = 314
	ViolinMarkerSymbol314                     ViolinMarkerSymbol = "314"
	ViolinMarkerSymbolHexagonOpenDot          ViolinMarkerSymbol = "hexagon-open-dot"
	ViolinMarkerSymbolNumber15                ViolinMarkerSymbol = 15
	ViolinMarkerSymbol15                      ViolinMarkerSymbol = "15"
	ViolinMarkerSymbolHexagon2                ViolinMarkerSymbol = "hexagon2"
	ViolinMarkerSymbolNumber115               ViolinMarkerSymbol = 115
	ViolinMarkerSymbol115                     ViolinMarkerSymbol = "115"
	ViolinMarkerSymbolHexagon2Open            ViolinMarkerSymbol = "hexagon2-open"
	ViolinMarkerSymbolNumber215               ViolinMarkerSymbol = 215
	ViolinMarkerSymbol215                     ViolinMarkerSymbol = "215"
	ViolinMarkerSymbolHexagon2Dot             ViolinMarkerSymbol = "hexagon2-dot"
	ViolinMarkerSymbolNumber315               ViolinMarkerSymbol = 315
	ViolinMarkerSymbol315                     ViolinMarkerSymbol = "315"
	ViolinMarkerSymbolHexagon2OpenDot         ViolinMarkerSymbol = "hexagon2-open-dot"
	ViolinMarkerSymbolNumber16                ViolinMarkerSymbol = 16
	ViolinMarkerSymbol16                      ViolinMarkerSymbol = "16"
	ViolinMarkerSymbolOctagon                 ViolinMarkerSymbol = "octagon"
	ViolinMarkerSymbolNumber116               ViolinMarkerSymbol = 116
	ViolinMarkerSymbol116                     ViolinMarkerSymbol = "116"
	ViolinMarkerSymbolOctagonOpen             ViolinMarkerSymbol = "octagon-open"
	ViolinMarkerSymbolNumber216               ViolinMarkerSymbol = 216
	ViolinMarkerSymbol216                     ViolinMarkerSymbol = "216"
	ViolinMarkerSymbolOctagonDot              ViolinMarkerSymbol = "octagon-dot"
	ViolinMarkerSymbolNumber316               ViolinMarkerSymbol = 316
	ViolinMarkerSymbol316                     ViolinMarkerSymbol = "316"
	ViolinMarkerSymbolOctagonOpenDot          ViolinMarkerSymbol = "octagon-open-dot"
	ViolinMarkerSymbolNumber17                ViolinMarkerSymbol = 17
	ViolinMarkerSymbol17                      ViolinMarkerSymbol = "17"
	ViolinMarkerSymbolStar                    ViolinMarkerSymbol = "star"
	ViolinMarkerSymbolNumber117               ViolinMarkerSymbol = 117
	ViolinMarkerSymbol117                     ViolinMarkerSymbol = "117"
	ViolinMarkerSymbolStarOpen                ViolinMarkerSymbol = "star-open"
	ViolinMarkerSymbolNumber217               ViolinMarkerSymbol = 217
	ViolinMarkerSymbol217                     ViolinMarkerSymbol = "217"
	ViolinMarkerSymbolStarDot                 ViolinMarkerSymbol = "star-dot"
	ViolinMarkerSymbolNumber317               ViolinMarkerSymbol = 317
	ViolinMarkerSymbol317                     ViolinMarkerSymbol = "317"
	ViolinMarkerSymbolStarOpenDot             ViolinMarkerSymbol = "star-open-dot"
	ViolinMarkerSymbolNumber18                ViolinMarkerSymbol = 18
	ViolinMarkerSymbol18                      ViolinMarkerSymbol = "18"
	ViolinMarkerSymbolHexagram                ViolinMarkerSymbol = "hexagram"
	ViolinMarkerSymbolNumber118               ViolinMarkerSymbol = 118
	ViolinMarkerSymbol118                     ViolinMarkerSymbol = "118"
	ViolinMarkerSymbolHexagramOpen            ViolinMarkerSymbol = "hexagram-open"
	ViolinMarkerSymbolNumber218               ViolinMarkerSymbol = 218
	ViolinMarkerSymbol218                     ViolinMarkerSymbol = "218"
	ViolinMarkerSymbolHexagramDot             ViolinMarkerSymbol = "hexagram-dot"
	ViolinMarkerSymbolNumber318               ViolinMarkerSymbol = 318
	ViolinMarkerSymbol318                     ViolinMarkerSymbol = "318"
	ViolinMarkerSymbolHexagramOpenDot         ViolinMarkerSymbol = "hexagram-open-dot"
	ViolinMarkerSymbolNumber19                ViolinMarkerSymbol = 19
	ViolinMarkerSymbol19                      ViolinMarkerSymbol = "19"
	ViolinMarkerSymbolStarTriangleUp          ViolinMarkerSymbol = "star-triangle-up"
	ViolinMarkerSymbolNumber119               ViolinMarkerSymbol = 119
	ViolinMarkerSymbol119                     ViolinMarkerSymbol = "119"
	ViolinMarkerSymbolStarTriangleUpOpen      ViolinMarkerSymbol = "star-triangle-up-open"
	ViolinMarkerSymbolNumber219               ViolinMarkerSymbol = 219
	ViolinMarkerSymbol219                     ViolinMarkerSymbol = "219"
	ViolinMarkerSymbolStarTriangleUpDot       ViolinMarkerSymbol = "star-triangle-up-dot"
	ViolinMarkerSymbolNumber319               ViolinMarkerSymbol = 319
	ViolinMarkerSymbol319                     ViolinMarkerSymbol = "319"
	ViolinMarkerSymbolStarTriangleUpOpenDot   ViolinMarkerSymbol = "star-triangle-up-open-dot"
	ViolinMarkerSymbolNumber20                ViolinMarkerSymbol = 20
	ViolinMarkerSymbol20                      ViolinMarkerSymbol = "20"
	ViolinMarkerSymbolStarTriangleDown        ViolinMarkerSymbol = "star-triangle-down"
	ViolinMarkerSymbolNumber120               ViolinMarkerSymbol = 120
	ViolinMarkerSymbol120                     ViolinMarkerSymbol = "120"
	ViolinMarkerSymbolStarTriangleDownOpen    ViolinMarkerSymbol = "star-triangle-down-open"
	ViolinMarkerSymbolNumber220               ViolinMarkerSymbol = 220
	ViolinMarkerSymbol220                     ViolinMarkerSymbol = "220"
	ViolinMarkerSymbolStarTriangleDownDot     ViolinMarkerSymbol = "star-triangle-down-dot"
	ViolinMarkerSymbolNumber320               ViolinMarkerSymbol = 320
	ViolinMarkerSymbol320                     ViolinMarkerSymbol = "320"
	ViolinMarkerSymbolStarTriangleDownOpenDot ViolinMarkerSymbol = "star-triangle-down-open-dot"
	ViolinMarkerSymbolNumber21                ViolinMarkerSymbol = 21
	ViolinMarkerSymbol21                      ViolinMarkerSymbol = "21"
	ViolinMarkerSymbolStarSquare              ViolinMarkerSymbol = "star-square"
	ViolinMarkerSymbolNumber121               ViolinMarkerSymbol = 121
	ViolinMarkerSymbol121                     ViolinMarkerSymbol = "121"
	ViolinMarkerSymbolStarSquareOpen          ViolinMarkerSymbol = "star-square-open"
	ViolinMarkerSymbolNumber221               ViolinMarkerSymbol = 221
	ViolinMarkerSymbol221                     ViolinMarkerSymbol = "221"
	ViolinMarkerSymbolStarSquareDot           ViolinMarkerSymbol = "star-square-dot"
	ViolinMarkerSymbolNumber321               ViolinMarkerSymbol = 321
	ViolinMarkerSymbol321                     ViolinMarkerSymbol = "321"
	ViolinMarkerSymbolStarSquareOpenDot       ViolinMarkerSymbol = "star-square-open-dot"
	ViolinMarkerSymbolNumber22                ViolinMarkerSymbol = 22
	ViolinMarkerSymbol22                      ViolinMarkerSymbol = "22"
	ViolinMarkerSymbolStarDiamond             ViolinMarkerSymbol = "star-diamond"
	ViolinMarkerSymbolNumber122               ViolinMarkerSymbol = 122
	ViolinMarkerSymbol122                     ViolinMarkerSymbol = "122"
	ViolinMarkerSymbolStarDiamondOpen         ViolinMarkerSymbol = "star-diamond-open"
	ViolinMarkerSymbolNumber222               ViolinMarkerSymbol = 222
	ViolinMarkerSymbol222                     ViolinMarkerSymbol = "222"
	ViolinMarkerSymbolStarDiamondDot          ViolinMarkerSymbol = "star-diamond-dot"
	ViolinMarkerSymbolNumber322               ViolinMarkerSymbol = 322
	ViolinMarkerSymbol322                     ViolinMarkerSymbol = "322"
	ViolinMarkerSymbolStarDiamondOpenDot      ViolinMarkerSymbol = "star-diamond-open-dot"
	ViolinMarkerSymbolNumber23                ViolinMarkerSymbol = 23
	ViolinMarkerSymbol23                      ViolinMarkerSymbol = "23"
	ViolinMarkerSymbolDiamondTall             ViolinMarkerSymbol = "diamond-tall"
	ViolinMarkerSymbolNumber123               ViolinMarkerSymbol = 123
	ViolinMarkerSymbol123                     ViolinMarkerSymbol = "123"
	ViolinMarkerSymbolDiamondTallOpen         ViolinMarkerSymbol = "diamond-tall-open"
	ViolinMarkerSymbolNumber223               ViolinMarkerSymbol = 223
	ViolinMarkerSymbol223                     ViolinMarkerSymbol = "223"
	ViolinMarkerSymbolDiamondTallDot          ViolinMarkerSymbol = "diamond-tall-dot"
	ViolinMarkerSymbolNumber323               ViolinMarkerSymbol = 323
	ViolinMarkerSymbol323                     ViolinMarkerSymbol = "323"
	ViolinMarkerSymbolDiamondTallOpenDot      ViolinMarkerSymbol = "diamond-tall-open-dot"
	ViolinMarkerSymbolNumber24                ViolinMarkerSymbol = 24
	ViolinMarkerSymbol24                      ViolinMarkerSymbol = "24"
	ViolinMarkerSymbolDiamondWide             ViolinMarkerSymbol = "diamond-wide"
	ViolinMarkerSymbolNumber124               ViolinMarkerSymbol = 124
	ViolinMarkerSymbol124                     ViolinMarkerSymbol = "124"
	ViolinMarkerSymbolDiamondWideOpen         ViolinMarkerSymbol = "diamond-wide-open"
	ViolinMarkerSymbolNumber224               ViolinMarkerSymbol = 224
	ViolinMarkerSymbol224                     ViolinMarkerSymbol = "224"
	ViolinMarkerSymbolDiamondWideDot          ViolinMarkerSymbol = "diamond-wide-dot"
	ViolinMarkerSymbolNumber324               ViolinMarkerSymbol = 324
	ViolinMarkerSymbol324                     ViolinMarkerSymbol = "324"
	ViolinMarkerSymbolDiamondWideOpenDot      ViolinMarkerSymbol = "diamond-wide-open-dot"
	ViolinMarkerSymbolNumber25                ViolinMarkerSymbol = 25
	ViolinMarkerSymbol25                      ViolinMarkerSymbol = "25"
	ViolinMarkerSymbolHourglass               ViolinMarkerSymbol = "hourglass"
	ViolinMarkerSymbolNumber125               ViolinMarkerSymbol = 125
	ViolinMarkerSymbol125                     ViolinMarkerSymbol = "125"
	ViolinMarkerSymbolHourglassOpen           ViolinMarkerSymbol = "hourglass-open"
	ViolinMarkerSymbolNumber26                ViolinMarkerSymbol = 26
	ViolinMarkerSymbol26                      ViolinMarkerSymbol = "26"
	ViolinMarkerSymbolBowtie                  ViolinMarkerSymbol = "bowtie"
	ViolinMarkerSymbolNumber126               ViolinMarkerSymbol = 126
	ViolinMarkerSymbol126                     ViolinMarkerSymbol = "126"
	ViolinMarkerSymbolBowtieOpen              ViolinMarkerSymbol = "bowtie-open"
	ViolinMarkerSymbolNumber27                ViolinMarkerSymbol = 27
	ViolinMarkerSymbol27                      ViolinMarkerSymbol = "27"
	ViolinMarkerSymbolCircleCross             ViolinMarkerSymbol = "circle-cross"
	ViolinMarkerSymbolNumber127               ViolinMarkerSymbol = 127
	ViolinMarkerSymbol127                     ViolinMarkerSymbol = "127"
	ViolinMarkerSymbolCircleCrossOpen         ViolinMarkerSymbol = "circle-cross-open"
	ViolinMarkerSymbolNumber28                ViolinMarkerSymbol = 28
	ViolinMarkerSymbol28                      ViolinMarkerSymbol = "28"
	ViolinMarkerSymbolCircleX                 ViolinMarkerSymbol = "circle-x"
	ViolinMarkerSymbolNumber128               ViolinMarkerSymbol = 128
	ViolinMarkerSymbol128                     ViolinMarkerSymbol = "128"
	ViolinMarkerSymbolCircleXOpen             ViolinMarkerSymbol = "circle-x-open"
	ViolinMarkerSymbolNumber29                ViolinMarkerSymbol = 29
	ViolinMarkerSymbol29                      ViolinMarkerSymbol = "29"
	ViolinMarkerSymbolSquareCross             ViolinMarkerSymbol = "square-cross"
	ViolinMarkerSymbolNumber129               ViolinMarkerSymbol = 129
	ViolinMarkerSymbol129                     ViolinMarkerSymbol = "129"
	ViolinMarkerSymbolSquareCrossOpen         ViolinMarkerSymbol = "square-cross-open"
	ViolinMarkerSymbolNumber30                ViolinMarkerSymbol = 30
	ViolinMarkerSymbol30                      ViolinMarkerSymbol = "30"
	ViolinMarkerSymbolSquareX                 ViolinMarkerSymbol = "square-x"
	ViolinMarkerSymbolNumber130               ViolinMarkerSymbol = 130
	ViolinMarkerSymbol130                     ViolinMarkerSymbol = "130"
	ViolinMarkerSymbolSquareXOpen             ViolinMarkerSymbol = "square-x-open"
	ViolinMarkerSymbolNumber31                ViolinMarkerSymbol = 31
	ViolinMarkerSymbol31                      ViolinMarkerSymbol = "31"
	ViolinMarkerSymbolDiamondCross            ViolinMarkerSymbol = "diamond-cross"
	ViolinMarkerSymbolNumber131               ViolinMarkerSymbol = 131
	ViolinMarkerSymbol131                     ViolinMarkerSymbol = "131"
	ViolinMarkerSymbolDiamondCrossOpen        ViolinMarkerSymbol = "diamond-cross-open"
	ViolinMarkerSymbolNumber32                ViolinMarkerSymbol = 32
	ViolinMarkerSymbol32                      ViolinMarkerSymbol = "32"
	ViolinMarkerSymbolDiamondX                ViolinMarkerSymbol = "diamond-x"
	ViolinMarkerSymbolNumber132               ViolinMarkerSymbol = 132
	ViolinMarkerSymbol132                     ViolinMarkerSymbol = "132"
	ViolinMarkerSymbolDiamondXOpen            ViolinMarkerSymbol = "diamond-x-open"
	ViolinMarkerSymbolNumber33                ViolinMarkerSymbol = 33
	ViolinMarkerSymbol33                      ViolinMarkerSymbol = "33"
	ViolinMarkerSymbolCrossThin               ViolinMarkerSymbol = "cross-thin"
	ViolinMarkerSymbolNumber133               ViolinMarkerSymbol = 133
	ViolinMarkerSymbol133                     ViolinMarkerSymbol = "133"
	ViolinMarkerSymbolCrossThinOpen           ViolinMarkerSymbol = "cross-thin-open"
	ViolinMarkerSymbolNumber34                ViolinMarkerSymbol = 34
	ViolinMarkerSymbol34                      ViolinMarkerSymbol = "34"
	ViolinMarkerSymbolXThin                   ViolinMarkerSymbol = "x-thin"
	ViolinMarkerSymbolNumber134               ViolinMarkerSymbol = 134
	ViolinMarkerSymbol134                     ViolinMarkerSymbol = "134"
	ViolinMarkerSymbolXThinOpen               ViolinMarkerSymbol = "x-thin-open"
	ViolinMarkerSymbolNumber35                ViolinMarkerSymbol = 35
	ViolinMarkerSymbol35                      ViolinMarkerSymbol = "35"
	ViolinMarkerSymbolAsterisk                ViolinMarkerSymbol = "asterisk"
	ViolinMarkerSymbolNumber135               ViolinMarkerSymbol = 135
	ViolinMarkerSymbol135                     ViolinMarkerSymbol = "135"
	ViolinMarkerSymbolAsteriskOpen            ViolinMarkerSymbol = "asterisk-open"
	ViolinMarkerSymbolNumber36                ViolinMarkerSymbol = 36
	ViolinMarkerSymbol36                      ViolinMarkerSymbol = "36"
	ViolinMarkerSymbolHash                    ViolinMarkerSymbol = "hash"
	ViolinMarkerSymbolNumber136               ViolinMarkerSymbol = 136
	ViolinMarkerSymbol136                     ViolinMarkerSymbol = "136"
	ViolinMarkerSymbolHashOpen                ViolinMarkerSymbol = "hash-open"
	ViolinMarkerSymbolNumber236               ViolinMarkerSymbol = 236
	ViolinMarkerSymbol236                     ViolinMarkerSymbol = "236"
	ViolinMarkerSymbolHashDot                 ViolinMarkerSymbol = "hash-dot"
	ViolinMarkerSymbolNumber336               ViolinMarkerSymbol = 336
	ViolinMarkerSymbol336                     ViolinMarkerSymbol = "336"
	ViolinMarkerSymbolHashOpenDot             ViolinMarkerSymbol = "hash-open-dot"
	ViolinMarkerSymbolNumber37                ViolinMarkerSymbol = 37
	ViolinMarkerSymbol37                      ViolinMarkerSymbol = "37"
	ViolinMarkerSymbolYUp                     ViolinMarkerSymbol = "y-up"
	ViolinMarkerSymbolNumber137               ViolinMarkerSymbol = 137
	ViolinMarkerSymbol137                     ViolinMarkerSymbol = "137"
	ViolinMarkerSymbolYUpOpen                 ViolinMarkerSymbol = "y-up-open"
	ViolinMarkerSymbolNumber38                ViolinMarkerSymbol = 38
	ViolinMarkerSymbol38                      ViolinMarkerSymbol = "38"
	ViolinMarkerSymbolYDown                   ViolinMarkerSymbol = "y-down"
	ViolinMarkerSymbolNumber138               ViolinMarkerSymbol = 138
	ViolinMarkerSymbol138                     ViolinMarkerSymbol = "138"
	ViolinMarkerSymbolYDownOpen               ViolinMarkerSymbol = "y-down-open"
	ViolinMarkerSymbolNumber39                ViolinMarkerSymbol = 39
	ViolinMarkerSymbol39                      ViolinMarkerSymbol = "39"
	ViolinMarkerSymbolYLeft                   ViolinMarkerSymbol = "y-left"
	ViolinMarkerSymbolNumber139               ViolinMarkerSymbol = 139
	ViolinMarkerSymbol139                     ViolinMarkerSymbol = "139"
	ViolinMarkerSymbolYLeftOpen               ViolinMarkerSymbol = "y-left-open"
	ViolinMarkerSymbolNumber40                ViolinMarkerSymbol = 40
	ViolinMarkerSymbol40                      ViolinMarkerSymbol = "40"
	ViolinMarkerSymbolYRight                  ViolinMarkerSymbol = "y-right"
	ViolinMarkerSymbolNumber140               ViolinMarkerSymbol = 140
	ViolinMarkerSymbol140                     ViolinMarkerSymbol = "140"
	ViolinMarkerSymbolYRightOpen              ViolinMarkerSymbol = "y-right-open"
	ViolinMarkerSymbolNumber41                ViolinMarkerSymbol = 41
	ViolinMarkerSymbol41                      ViolinMarkerSymbol = "41"
	ViolinMarkerSymbolLineEw                  ViolinMarkerSymbol = "line-ew"
	ViolinMarkerSymbolNumber141               ViolinMarkerSymbol = 141
	ViolinMarkerSymbol141                     ViolinMarkerSymbol = "141"
	ViolinMarkerSymbolLineEwOpen              ViolinMarkerSymbol = "line-ew-open"
	ViolinMarkerSymbolNumber42                ViolinMarkerSymbol = 42
	ViolinMarkerSymbol42                      ViolinMarkerSymbol = "42"
	ViolinMarkerSymbolLineNs                  ViolinMarkerSymbol = "line-ns"
	ViolinMarkerSymbolNumber142               ViolinMarkerSymbol = 142
	ViolinMarkerSymbol142                     ViolinMarkerSymbol = "142"
	ViolinMarkerSymbolLineNsOpen              ViolinMarkerSymbol = "line-ns-open"
	ViolinMarkerSymbolNumber43                ViolinMarkerSymbol = 43
	ViolinMarkerSymbol43                      ViolinMarkerSymbol = "43"
	ViolinMarkerSymbolLineNe                  ViolinMarkerSymbol = "line-ne"
	ViolinMarkerSymbolNumber143               ViolinMarkerSymbol = 143
	ViolinMarkerSymbol143                     ViolinMarkerSymbol = "143"
	ViolinMarkerSymbolLineNeOpen              ViolinMarkerSymbol = "line-ne-open"
	ViolinMarkerSymbolNumber44                ViolinMarkerSymbol = 44
	ViolinMarkerSymbol44                      ViolinMarkerSymbol = "44"
	ViolinMarkerSymbolLineNw                  ViolinMarkerSymbol = "line-nw"
	ViolinMarkerSymbolNumber144               ViolinMarkerSymbol = 144
	ViolinMarkerSymbol144                     ViolinMarkerSymbol = "144"
	ViolinMarkerSymbolLineNwOpen              ViolinMarkerSymbol = "line-nw-open"
	ViolinMarkerSymbolNumber45                ViolinMarkerSymbol = 45
	ViolinMarkerSymbol45                      ViolinMarkerSymbol = "45"
	ViolinMarkerSymbolArrowUp                 ViolinMarkerSymbol = "arrow-up"
	ViolinMarkerSymbolNumber145               ViolinMarkerSymbol = 145
	ViolinMarkerSymbol145                     ViolinMarkerSymbol = "145"
	ViolinMarkerSymbolArrowUpOpen             ViolinMarkerSymbol = "arrow-up-open"
	ViolinMarkerSymbolNumber46                ViolinMarkerSymbol = 46
	ViolinMarkerSymbol46                      ViolinMarkerSymbol = "46"
	ViolinMarkerSymbolArrowDown               ViolinMarkerSymbol = "arrow-down"
	ViolinMarkerSymbolNumber146               ViolinMarkerSymbol = 146
	ViolinMarkerSymbol146                     ViolinMarkerSymbol = "146"
	ViolinMarkerSymbolArrowDownOpen           ViolinMarkerSymbol = "arrow-down-open"
	ViolinMarkerSymbolNumber47                ViolinMarkerSymbol = 47
	ViolinMarkerSymbol47                      ViolinMarkerSymbol = "47"
	ViolinMarkerSymbolArrowLeft               ViolinMarkerSymbol = "arrow-left"
	ViolinMarkerSymbolNumber147               ViolinMarkerSymbol = 147
	ViolinMarkerSymbol147                     ViolinMarkerSymbol = "147"
	ViolinMarkerSymbolArrowLeftOpen           ViolinMarkerSymbol = "arrow-left-open"
	ViolinMarkerSymbolNumber48                ViolinMarkerSymbol = 48
	ViolinMarkerSymbol48                      ViolinMarkerSymbol = "48"
	ViolinMarkerSymbolArrowRight              ViolinMarkerSymbol = "arrow-right"
	ViolinMarkerSymbolNumber148               ViolinMarkerSymbol = 148
	ViolinMarkerSymbol148                     ViolinMarkerSymbol = "148"
	ViolinMarkerSymbolArrowRightOpen          ViolinMarkerSymbol = "arrow-right-open"
	ViolinMarkerSymbolNumber49                ViolinMarkerSymbol = 49
	ViolinMarkerSymbol49                      ViolinMarkerSymbol = "49"
	ViolinMarkerSymbolArrowBarUp              ViolinMarkerSymbol = "arrow-bar-up"
	ViolinMarkerSymbolNumber149               ViolinMarkerSymbol = 149
	ViolinMarkerSymbol149                     ViolinMarkerSymbol = "149"
	ViolinMarkerSymbolArrowBarUpOpen          ViolinMarkerSymbol = "arrow-bar-up-open"
	ViolinMarkerSymbolNumber50                ViolinMarkerSymbol = 50
	ViolinMarkerSymbol50                      ViolinMarkerSymbol = "50"
	ViolinMarkerSymbolArrowBarDown            ViolinMarkerSymbol = "arrow-bar-down"
	ViolinMarkerSymbolNumber150               ViolinMarkerSymbol = 150
	ViolinMarkerSymbol150                     ViolinMarkerSymbol = "150"
	ViolinMarkerSymbolArrowBarDownOpen        ViolinMarkerSymbol = "arrow-bar-down-open"
	ViolinMarkerSymbolNumber51                ViolinMarkerSymbol = 51
	ViolinMarkerSymbol51                      ViolinMarkerSymbol = "51"
	ViolinMarkerSymbolArrowBarLeft            ViolinMarkerSymbol = "arrow-bar-left"
	ViolinMarkerSymbolNumber151               ViolinMarkerSymbol = 151
	ViolinMarkerSymbol151                     ViolinMarkerSymbol = "151"
	ViolinMarkerSymbolArrowBarLeftOpen        ViolinMarkerSymbol = "arrow-bar-left-open"
	ViolinMarkerSymbolNumber52                ViolinMarkerSymbol = 52
	ViolinMarkerSymbol52                      ViolinMarkerSymbol = "52"
	ViolinMarkerSymbolArrowBarRight           ViolinMarkerSymbol = "arrow-bar-right"
	ViolinMarkerSymbolNumber152               ViolinMarkerSymbol = 152
	ViolinMarkerSymbol152                     ViolinMarkerSymbol = "152"
	ViolinMarkerSymbolArrowBarRightOpen       ViolinMarkerSymbol = "arrow-bar-right-open"
)

// ViolinOrientation Sets the orientation of the violin(s). If *v* (*h*), the distribution is visualized along the vertical (horizontal).
type ViolinOrientation string

const (
	ViolinOrientationV ViolinOrientation = "v"
	ViolinOrientationH ViolinOrientation = "h"
)

// ViolinPoints If *outliers*, only the sample points lying outside the whiskers are shown If *suspectedoutliers*, the outlier points are shown and points either less than 4*Q1-3*Q3 or greater than 4*Q3-3*Q1 are highlighted (see `outliercolor`) If *all*, all sample points are shown If *false*, only the violins are shown with no sample points. Defaults to *suspectedoutliers* when `marker.outliercolor` or `marker.line.outliercolor` is set, otherwise defaults to *outliers*.
type ViolinPoints interface{}

var (
	ViolinPointsAll               ViolinPoints = "all"
	ViolinPointsOutliers          ViolinPoints = "outliers"
	ViolinPointsSuspectedoutliers ViolinPoints = "suspectedoutliers"
	ViolinPointsFalse             ViolinPoints = false
)

// ViolinScalemode Sets the metric by which the width of each violin is determined.*width* means each violin has the same (max) width*count* means the violins are scaled by the number of sample points makingup each violin.
type ViolinScalemode string

const (
	ViolinScalemodeWidth ViolinScalemode = "width"
	ViolinScalemodeCount ViolinScalemode = "count"
)

// ViolinSide Determines on which side of the position value the density function making up one half of a violin is plotted. Useful when comparing two violin traces under *overlay* mode, where one trace has `side` set to *positive* and the other to *negative*.
type ViolinSide string

const (
	ViolinSideBoth     ViolinSide = "both"
	ViolinSidePositive ViolinSide = "positive"
	ViolinSideNegative ViolinSide = "negative"
)

// ViolinSpanmode Sets the method by which the span in data space where the density function will be computed. *soft* means the span goes from the sample's minimum value minus two bandwidths to the sample's maximum value plus two bandwidths. *hard* means the span goes from the sample's minimum to its maximum value. For custom span settings, use mode *manual* and fill in the `span` attribute.
type ViolinSpanmode string

const (
	ViolinSpanmodeSoft   ViolinSpanmode = "soft"
	ViolinSpanmodeHard   ViolinSpanmode = "hard"
	ViolinSpanmodeManual ViolinSpanmode = "manual"
)

// ViolinVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ViolinVisible interface{}

var (
	ViolinVisibleTrue       ViolinVisible = true
	ViolinVisibleFalse      ViolinVisible = false
	ViolinVisibleLegendonly ViolinVisible = "legendonly"
)

// ViolinHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type ViolinHoverinfo string

const (
	// Flags
	ViolinHoverinfoX    ViolinHoverinfo = "x"
	ViolinHoverinfoY    ViolinHoverinfo = "y"
	ViolinHoverinfoZ    ViolinHoverinfo = "z"
	ViolinHoverinfoText ViolinHoverinfo = "text"
	ViolinHoverinfoName ViolinHoverinfo = "name"

	// Extra
	ViolinHoverinfoAll  ViolinHoverinfo = "all"
	ViolinHoverinfoNone ViolinHoverinfo = "none"
	ViolinHoverinfoSkip ViolinHoverinfo = "skip"
)

// ViolinHoveron Do the hover effects highlight individual violins or sample points or the kernel density estimate or any combination of them?
type ViolinHoveron string

const (
	// Flags
	ViolinHoveronViolins ViolinHoveron = "violins"
	ViolinHoveronPoints  ViolinHoveron = "points"
	ViolinHoveronKde     ViolinHoveron = "kde"

	// Extra
	ViolinHoveronAll ViolinHoveron = "all"
)
