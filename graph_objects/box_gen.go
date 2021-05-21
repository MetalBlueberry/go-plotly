package grob

var TraceTypeBox TraceType = "box"

func (trace *Box) GetType() TraceType {
	return TraceTypeBox
}

// Box Each box spans from quartile 1 (Q1) to quartile 3 (Q3). The second quartile (Q2, i.e. the median) is marked by a line inside the box. The fences grow outward from the boxes' edges, by default they span +/- 1.5 times the interquartile range (IQR: Q3-Q1), The sample mean and standard deviation as well as notches and the sample, outlier and suspected outliers points can be optionally added to the box plot. The values and positions corresponding to each boxes can be input using two signatures. The first signature expects users to supply the sample values in the `y` data array for vertical boxes (`x` for horizontal boxes). By supplying an `x` (`y`) array, one box per distinct `x` (`y`) value is drawn If no `x` (`y`) {array} is provided, a single box is drawn. In this case, the box is positioned with the trace `name` or with `x0` (`y0`) if provided. The second signature expects users to supply the boxes corresponding Q1, median and Q3 statistics in the `q1`, `median` and `q3` data arrays respectively. Other box features relying on statistics namely `lowerfence`, `upperfence`, `notchspan` can be set directly by the users. To have plotly compute them or to show sample points besides the boxes, users can set the `y` data array for vertical boxes (`x` for horizontal boxes) to a 2D array with the outer length corresponding to the number of boxes in the traces and the inner length corresponding the sample size.
type Box struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Alignmentgroup
	// arrayOK: false
	// type: string
	// Set several traces linked to the same position axis or matching axes to the same alignmentgroup. This controls whether bars compute their positional range dependently or independently.
	Alignmentgroup String `json:"alignmentgroup,omitempty"`

	// Boxmean
	// default: %!s(<nil>)
	// type: enumerated
	// If *true*, the mean of the box(es)' underlying distribution is drawn as a dashed line inside the box(es). If *sd* the standard deviation is also drawn. Defaults to *true* when `mean` is set. Defaults to *sd* when `sd` is set Otherwise defaults to *false*.
	Boxmean BoxBoxmean `json:"boxmean,omitempty"`

	// Boxpoints
	// default: %!s(<nil>)
	// type: enumerated
	// If *outliers*, only the sample points lying outside the whiskers are shown If *suspectedoutliers*, the outlier points are shown and points either less than 4*Q1-3*Q3 or greater than 4*Q3-3*Q1 are highlighted (see `outliercolor`) If *all*, all sample points are shown If *false*, only the box(es) are shown with no sample points Defaults to *suspectedoutliers* when `marker.outliercolor` or `marker.line.outliercolor` is set. Defaults to *all* under the q1/median/q3 signature. Otherwise defaults to *outliers*.
	Boxpoints BoxBoxpoints `json:"boxpoints,omitempty"`

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
	// Sets the x coordinate step for multi-box traces set using q1/median/q3.
	Dx float64 `json:"dx,omitempty"`

	// Dy
	// arrayOK: false
	// type: number
	// Sets the y coordinate step for multi-box traces set using q1/median/q3.
	Dy float64 `json:"dy,omitempty"`

	// Fillcolor
	// arrayOK: false
	// type: color
	// Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor Color `json:"fillcolor,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo BoxHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *BoxHoverlabel `json:"hoverlabel,omitempty"`

	// Hoveron
	// default: boxes+points
	// type: flaglist
	// Do the hover effects highlight individual boxes  or sample points or both?
	Hoveron BoxHoveron `json:"hoveron,omitempty"`

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
	// Sets the amount of jitter in the sample points drawn. If *0*, the sample points align along the distribution axis. If *1*, the sample points are drawn in a random jitter of width equal to the width of the box(es).
	Jitter float64 `json:"jitter,omitempty"`

	// Legendgroup
	// arrayOK: false
	// type: string
	// Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line
	// role: Object
	Line *BoxLine `json:"line,omitempty"`

	// Lowerfence
	// arrayOK: false
	// type: data_array
	// Sets the lower fence values. There should be as many items as the number of boxes desired. This attribute has effect only under the q1/median/q3 signature. If `lowerfence` is not provided but a sample (in `y` or `x`) is set, we compute the lower as the last sample point below 1.5 times the IQR.
	Lowerfence interface{} `json:"lowerfence,omitempty"`

	// Lowerfencesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  lowerfence .
	Lowerfencesrc String `json:"lowerfencesrc,omitempty"`

	// Marker
	// role: Object
	Marker *BoxMarker `json:"marker,omitempty"`

	// Mean
	// arrayOK: false
	// type: data_array
	// Sets the mean values. There should be as many items as the number of boxes desired. This attribute has effect only under the q1/median/q3 signature. If `mean` is not provided but a sample (in `y` or `x`) is set, we compute the mean for each box using the sample values.
	Mean interface{} `json:"mean,omitempty"`

	// Meansrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  mean .
	Meansrc String `json:"meansrc,omitempty"`

	// Median
	// arrayOK: false
	// type: data_array
	// Sets the median values. There should be as many items as the number of boxes desired.
	Median interface{} `json:"median,omitempty"`

	// Mediansrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  median .
	Mediansrc String `json:"mediansrc,omitempty"`

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
	// Sets the trace name. The trace name appear as the legend item and on hover. For box traces, the name will also be used for the position coordinate, if `x` and `x0` (`y` and `y0` if horizontal) are missing and the position axis is categorical
	Name String `json:"name,omitempty"`

	// Notched
	// arrayOK: false
	// type: boolean
	// Determines whether or not notches are drawn. Notches displays a confidence interval around the median. We compute the confidence interval as median +/- 1.57 * IQR / sqrt(N), where IQR is the interquartile range and N is the sample size. If two boxes' notches do not overlap there is 95% confidence their medians differ. See https://sites.google.com/site/davidsstatistics/home/notched-box-plots for more info. Defaults to *false* unless `notchwidth` or `notchspan` is set.
	Notched Bool `json:"notched,omitempty"`

	// Notchspan
	// arrayOK: false
	// type: data_array
	// Sets the notch span from the boxes' `median` values. There should be as many items as the number of boxes desired. This attribute has effect only under the q1/median/q3 signature. If `notchspan` is not provided but a sample (in `y` or `x`) is set, we compute it as 1.57 * IQR / sqrt(N), where N is the sample size.
	Notchspan interface{} `json:"notchspan,omitempty"`

	// Notchspansrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  notchspan .
	Notchspansrc String `json:"notchspansrc,omitempty"`

	// Notchwidth
	// arrayOK: false
	// type: number
	// Sets the width of the notches relative to the box' width. For example, with 0, the notches are as wide as the box(es).
	Notchwidth float64 `json:"notchwidth,omitempty"`

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
	// Sets the orientation of the box(es). If *v* (*h*), the distribution is visualized along the vertical (horizontal).
	Orientation BoxOrientation `json:"orientation,omitempty"`

	// Pointpos
	// arrayOK: false
	// type: number
	// Sets the position of the sample points in relation to the box(es). If *0*, the sample points are places over the center of the box(es). Positive (negative) values correspond to positions to the right (left) for vertical boxes and above (below) for horizontal boxes
	Pointpos float64 `json:"pointpos,omitempty"`

	// Q1
	// arrayOK: false
	// type: data_array
	// Sets the Quartile 1 values. There should be as many items as the number of boxes desired.
	Q1 interface{} `json:"q1,omitempty"`

	// Q1src
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  q1 .
	Q1src String `json:"q1src,omitempty"`

	// Q3
	// arrayOK: false
	// type: data_array
	// Sets the Quartile 3 values. There should be as many items as the number of boxes desired.
	Q3 interface{} `json:"q3,omitempty"`

	// Q3src
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  q3 .
	Q3src String `json:"q3src,omitempty"`

	// Quartilemethod
	// default: linear
	// type: enumerated
	// Sets the method used to compute the sample's Q1 and Q3 quartiles. The *linear* method uses the 25th percentile for Q1 and 75th percentile for Q3 as computed using method #10 (listed on http://www.amstat.org/publications/jse/v14n3/langford.html). The *exclusive* method uses the median to divide the ordered dataset into two halves if the sample is odd, it does not include the median in either half - Q1 is then the median of the lower half and Q3 the median of the upper half. The *inclusive* method also uses the median to divide the ordered dataset into two halves but if the sample is odd, it includes the median in both halves - Q1 is then the median of the lower half and Q3 the median of the upper half.
	Quartilemethod BoxQuartilemethod `json:"quartilemethod,omitempty"`

	// Sd
	// arrayOK: false
	// type: data_array
	// Sets the standard deviation values. There should be as many items as the number of boxes desired. This attribute has effect only under the q1/median/q3 signature. If `sd` is not provided but a sample (in `y` or `x`) is set, we compute the standard deviation for each box using the sample values.
	Sd interface{} `json:"sd,omitempty"`

	// Sdsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  sd .
	Sdsrc String `json:"sdsrc,omitempty"`

	// Selected
	// role: Object
	Selected *BoxSelected `json:"selected,omitempty"`

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
	Stream *BoxStream `json:"stream,omitempty"`

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
	Unselected *BoxUnselected `json:"unselected,omitempty"`

	// Upperfence
	// arrayOK: false
	// type: data_array
	// Sets the upper fence values. There should be as many items as the number of boxes desired. This attribute has effect only under the q1/median/q3 signature. If `upperfence` is not provided but a sample (in `y` or `x`) is set, we compute the lower as the last sample point above 1.5 times the IQR.
	Upperfence interface{} `json:"upperfence,omitempty"`

	// Upperfencesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  upperfence .
	Upperfencesrc String `json:"upperfencesrc,omitempty"`

	// Visible
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible BoxVisible `json:"visible,omitempty"`

	// Whiskerwidth
	// arrayOK: false
	// type: number
	// Sets the width of the whiskers relative to the box' width. For example, with 1, the whiskers are as wide as the box(es).
	Whiskerwidth float64 `json:"whiskerwidth,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the width of the box in data coordinate If *0* (default value) the width is automatically selected based on the positions of other box traces in the same subplot.
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

	// Xcalendar
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `x` date data.
	Xcalendar BoxXcalendar `json:"xcalendar,omitempty"`

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
	Xperiodalignment BoxXperiodalignment `json:"xperiodalignment,omitempty"`

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

	// Ycalendar
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `y` date data.
	Ycalendar BoxYcalendar `json:"ycalendar,omitempty"`

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
	Yperiodalignment BoxYperiodalignment `json:"yperiodalignment,omitempty"`

	// Ysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

// BoxHoverlabelFont Sets the font used in hover labels.
type BoxHoverlabelFont struct {

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

// BoxHoverlabel
type BoxHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align BoxHoverlabelAlign `json:"align,omitempty"`

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
	Font *BoxHoverlabelFont `json:"font,omitempty"`

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

// BoxLine
type BoxLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the color of line bounding the box(es).
	Color Color `json:"color,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the width (in px) of line bounding the box(es).
	Width float64 `json:"width,omitempty"`
}

// BoxMarkerLine
type BoxMarkerLine struct {

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

// BoxMarker
type BoxMarker struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets themarkercolor. It accepts either a specific color or an array of numbers that are mapped to the colorscale relative to the max and min values of the array or relative to `marker.cmin` and `marker.cmax` if set.
	Color Color `json:"color,omitempty"`

	// Line
	// role: Object
	Line *BoxMarkerLine `json:"line,omitempty"`

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
	Symbol BoxMarkerSymbol `json:"symbol,omitempty"`
}

// BoxSelectedMarker
type BoxSelectedMarker struct {

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

// BoxSelected
type BoxSelected struct {

	// Marker
	// role: Object
	Marker *BoxSelectedMarker `json:"marker,omitempty"`
}

// BoxStream
type BoxStream struct {

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

// BoxUnselectedMarker
type BoxUnselectedMarker struct {

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

// BoxUnselected
type BoxUnselected struct {

	// Marker
	// role: Object
	Marker *BoxUnselectedMarker `json:"marker,omitempty"`
}

// BoxBoxmean If *true*, the mean of the box(es)' underlying distribution is drawn as a dashed line inside the box(es). If *sd* the standard deviation is also drawn. Defaults to *true* when `mean` is set. Defaults to *sd* when `sd` is set Otherwise defaults to *false*.
type BoxBoxmean interface{}

var (
	BoxBoxmeanTrue  BoxBoxmean = true
	BoxBoxmeanSd    BoxBoxmean = "sd"
	BoxBoxmeanFalse BoxBoxmean = false
)

// BoxBoxpoints If *outliers*, only the sample points lying outside the whiskers are shown If *suspectedoutliers*, the outlier points are shown and points either less than 4*Q1-3*Q3 or greater than 4*Q3-3*Q1 are highlighted (see `outliercolor`) If *all*, all sample points are shown If *false*, only the box(es) are shown with no sample points Defaults to *suspectedoutliers* when `marker.outliercolor` or `marker.line.outliercolor` is set. Defaults to *all* under the q1/median/q3 signature. Otherwise defaults to *outliers*.
type BoxBoxpoints interface{}

var (
	BoxBoxpointsAll               BoxBoxpoints = "all"
	BoxBoxpointsOutliers          BoxBoxpoints = "outliers"
	BoxBoxpointsSuspectedoutliers BoxBoxpoints = "suspectedoutliers"
	BoxBoxpointsFalse             BoxBoxpoints = false
)

// BoxHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type BoxHoverlabelAlign string

const (
	BoxHoverlabelAlignLeft  BoxHoverlabelAlign = "left"
	BoxHoverlabelAlignRight BoxHoverlabelAlign = "right"
	BoxHoverlabelAlignAuto  BoxHoverlabelAlign = "auto"
)

// BoxMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type BoxMarkerSymbol interface{}

var (
	BoxMarkerSymbolNumber0                 BoxMarkerSymbol = 0
	BoxMarkerSymbol0                       BoxMarkerSymbol = "0"
	BoxMarkerSymbolCircle                  BoxMarkerSymbol = "circle"
	BoxMarkerSymbolNumber100               BoxMarkerSymbol = 100
	BoxMarkerSymbol100                     BoxMarkerSymbol = "100"
	BoxMarkerSymbolCircleOpen              BoxMarkerSymbol = "circle-open"
	BoxMarkerSymbolNumber200               BoxMarkerSymbol = 200
	BoxMarkerSymbol200                     BoxMarkerSymbol = "200"
	BoxMarkerSymbolCircleDot               BoxMarkerSymbol = "circle-dot"
	BoxMarkerSymbolNumber300               BoxMarkerSymbol = 300
	BoxMarkerSymbol300                     BoxMarkerSymbol = "300"
	BoxMarkerSymbolCircleOpenDot           BoxMarkerSymbol = "circle-open-dot"
	BoxMarkerSymbolNumber1                 BoxMarkerSymbol = 1
	BoxMarkerSymbol1                       BoxMarkerSymbol = "1"
	BoxMarkerSymbolSquare                  BoxMarkerSymbol = "square"
	BoxMarkerSymbolNumber101               BoxMarkerSymbol = 101
	BoxMarkerSymbol101                     BoxMarkerSymbol = "101"
	BoxMarkerSymbolSquareOpen              BoxMarkerSymbol = "square-open"
	BoxMarkerSymbolNumber201               BoxMarkerSymbol = 201
	BoxMarkerSymbol201                     BoxMarkerSymbol = "201"
	BoxMarkerSymbolSquareDot               BoxMarkerSymbol = "square-dot"
	BoxMarkerSymbolNumber301               BoxMarkerSymbol = 301
	BoxMarkerSymbol301                     BoxMarkerSymbol = "301"
	BoxMarkerSymbolSquareOpenDot           BoxMarkerSymbol = "square-open-dot"
	BoxMarkerSymbolNumber2                 BoxMarkerSymbol = 2
	BoxMarkerSymbol2                       BoxMarkerSymbol = "2"
	BoxMarkerSymbolDiamond                 BoxMarkerSymbol = "diamond"
	BoxMarkerSymbolNumber102               BoxMarkerSymbol = 102
	BoxMarkerSymbol102                     BoxMarkerSymbol = "102"
	BoxMarkerSymbolDiamondOpen             BoxMarkerSymbol = "diamond-open"
	BoxMarkerSymbolNumber202               BoxMarkerSymbol = 202
	BoxMarkerSymbol202                     BoxMarkerSymbol = "202"
	BoxMarkerSymbolDiamondDot              BoxMarkerSymbol = "diamond-dot"
	BoxMarkerSymbolNumber302               BoxMarkerSymbol = 302
	BoxMarkerSymbol302                     BoxMarkerSymbol = "302"
	BoxMarkerSymbolDiamondOpenDot          BoxMarkerSymbol = "diamond-open-dot"
	BoxMarkerSymbolNumber3                 BoxMarkerSymbol = 3
	BoxMarkerSymbol3                       BoxMarkerSymbol = "3"
	BoxMarkerSymbolCross                   BoxMarkerSymbol = "cross"
	BoxMarkerSymbolNumber103               BoxMarkerSymbol = 103
	BoxMarkerSymbol103                     BoxMarkerSymbol = "103"
	BoxMarkerSymbolCrossOpen               BoxMarkerSymbol = "cross-open"
	BoxMarkerSymbolNumber203               BoxMarkerSymbol = 203
	BoxMarkerSymbol203                     BoxMarkerSymbol = "203"
	BoxMarkerSymbolCrossDot                BoxMarkerSymbol = "cross-dot"
	BoxMarkerSymbolNumber303               BoxMarkerSymbol = 303
	BoxMarkerSymbol303                     BoxMarkerSymbol = "303"
	BoxMarkerSymbolCrossOpenDot            BoxMarkerSymbol = "cross-open-dot"
	BoxMarkerSymbolNumber4                 BoxMarkerSymbol = 4
	BoxMarkerSymbol4                       BoxMarkerSymbol = "4"
	BoxMarkerSymbolX                       BoxMarkerSymbol = "x"
	BoxMarkerSymbolNumber104               BoxMarkerSymbol = 104
	BoxMarkerSymbol104                     BoxMarkerSymbol = "104"
	BoxMarkerSymbolXOpen                   BoxMarkerSymbol = "x-open"
	BoxMarkerSymbolNumber204               BoxMarkerSymbol = 204
	BoxMarkerSymbol204                     BoxMarkerSymbol = "204"
	BoxMarkerSymbolXDot                    BoxMarkerSymbol = "x-dot"
	BoxMarkerSymbolNumber304               BoxMarkerSymbol = 304
	BoxMarkerSymbol304                     BoxMarkerSymbol = "304"
	BoxMarkerSymbolXOpenDot                BoxMarkerSymbol = "x-open-dot"
	BoxMarkerSymbolNumber5                 BoxMarkerSymbol = 5
	BoxMarkerSymbol5                       BoxMarkerSymbol = "5"
	BoxMarkerSymbolTriangleUp              BoxMarkerSymbol = "triangle-up"
	BoxMarkerSymbolNumber105               BoxMarkerSymbol = 105
	BoxMarkerSymbol105                     BoxMarkerSymbol = "105"
	BoxMarkerSymbolTriangleUpOpen          BoxMarkerSymbol = "triangle-up-open"
	BoxMarkerSymbolNumber205               BoxMarkerSymbol = 205
	BoxMarkerSymbol205                     BoxMarkerSymbol = "205"
	BoxMarkerSymbolTriangleUpDot           BoxMarkerSymbol = "triangle-up-dot"
	BoxMarkerSymbolNumber305               BoxMarkerSymbol = 305
	BoxMarkerSymbol305                     BoxMarkerSymbol = "305"
	BoxMarkerSymbolTriangleUpOpenDot       BoxMarkerSymbol = "triangle-up-open-dot"
	BoxMarkerSymbolNumber6                 BoxMarkerSymbol = 6
	BoxMarkerSymbol6                       BoxMarkerSymbol = "6"
	BoxMarkerSymbolTriangleDown            BoxMarkerSymbol = "triangle-down"
	BoxMarkerSymbolNumber106               BoxMarkerSymbol = 106
	BoxMarkerSymbol106                     BoxMarkerSymbol = "106"
	BoxMarkerSymbolTriangleDownOpen        BoxMarkerSymbol = "triangle-down-open"
	BoxMarkerSymbolNumber206               BoxMarkerSymbol = 206
	BoxMarkerSymbol206                     BoxMarkerSymbol = "206"
	BoxMarkerSymbolTriangleDownDot         BoxMarkerSymbol = "triangle-down-dot"
	BoxMarkerSymbolNumber306               BoxMarkerSymbol = 306
	BoxMarkerSymbol306                     BoxMarkerSymbol = "306"
	BoxMarkerSymbolTriangleDownOpenDot     BoxMarkerSymbol = "triangle-down-open-dot"
	BoxMarkerSymbolNumber7                 BoxMarkerSymbol = 7
	BoxMarkerSymbol7                       BoxMarkerSymbol = "7"
	BoxMarkerSymbolTriangleLeft            BoxMarkerSymbol = "triangle-left"
	BoxMarkerSymbolNumber107               BoxMarkerSymbol = 107
	BoxMarkerSymbol107                     BoxMarkerSymbol = "107"
	BoxMarkerSymbolTriangleLeftOpen        BoxMarkerSymbol = "triangle-left-open"
	BoxMarkerSymbolNumber207               BoxMarkerSymbol = 207
	BoxMarkerSymbol207                     BoxMarkerSymbol = "207"
	BoxMarkerSymbolTriangleLeftDot         BoxMarkerSymbol = "triangle-left-dot"
	BoxMarkerSymbolNumber307               BoxMarkerSymbol = 307
	BoxMarkerSymbol307                     BoxMarkerSymbol = "307"
	BoxMarkerSymbolTriangleLeftOpenDot     BoxMarkerSymbol = "triangle-left-open-dot"
	BoxMarkerSymbolNumber8                 BoxMarkerSymbol = 8
	BoxMarkerSymbol8                       BoxMarkerSymbol = "8"
	BoxMarkerSymbolTriangleRight           BoxMarkerSymbol = "triangle-right"
	BoxMarkerSymbolNumber108               BoxMarkerSymbol = 108
	BoxMarkerSymbol108                     BoxMarkerSymbol = "108"
	BoxMarkerSymbolTriangleRightOpen       BoxMarkerSymbol = "triangle-right-open"
	BoxMarkerSymbolNumber208               BoxMarkerSymbol = 208
	BoxMarkerSymbol208                     BoxMarkerSymbol = "208"
	BoxMarkerSymbolTriangleRightDot        BoxMarkerSymbol = "triangle-right-dot"
	BoxMarkerSymbolNumber308               BoxMarkerSymbol = 308
	BoxMarkerSymbol308                     BoxMarkerSymbol = "308"
	BoxMarkerSymbolTriangleRightOpenDot    BoxMarkerSymbol = "triangle-right-open-dot"
	BoxMarkerSymbolNumber9                 BoxMarkerSymbol = 9
	BoxMarkerSymbol9                       BoxMarkerSymbol = "9"
	BoxMarkerSymbolTriangleNe              BoxMarkerSymbol = "triangle-ne"
	BoxMarkerSymbolNumber109               BoxMarkerSymbol = 109
	BoxMarkerSymbol109                     BoxMarkerSymbol = "109"
	BoxMarkerSymbolTriangleNeOpen          BoxMarkerSymbol = "triangle-ne-open"
	BoxMarkerSymbolNumber209               BoxMarkerSymbol = 209
	BoxMarkerSymbol209                     BoxMarkerSymbol = "209"
	BoxMarkerSymbolTriangleNeDot           BoxMarkerSymbol = "triangle-ne-dot"
	BoxMarkerSymbolNumber309               BoxMarkerSymbol = 309
	BoxMarkerSymbol309                     BoxMarkerSymbol = "309"
	BoxMarkerSymbolTriangleNeOpenDot       BoxMarkerSymbol = "triangle-ne-open-dot"
	BoxMarkerSymbolNumber10                BoxMarkerSymbol = 10
	BoxMarkerSymbol10                      BoxMarkerSymbol = "10"
	BoxMarkerSymbolTriangleSe              BoxMarkerSymbol = "triangle-se"
	BoxMarkerSymbolNumber110               BoxMarkerSymbol = 110
	BoxMarkerSymbol110                     BoxMarkerSymbol = "110"
	BoxMarkerSymbolTriangleSeOpen          BoxMarkerSymbol = "triangle-se-open"
	BoxMarkerSymbolNumber210               BoxMarkerSymbol = 210
	BoxMarkerSymbol210                     BoxMarkerSymbol = "210"
	BoxMarkerSymbolTriangleSeDot           BoxMarkerSymbol = "triangle-se-dot"
	BoxMarkerSymbolNumber310               BoxMarkerSymbol = 310
	BoxMarkerSymbol310                     BoxMarkerSymbol = "310"
	BoxMarkerSymbolTriangleSeOpenDot       BoxMarkerSymbol = "triangle-se-open-dot"
	BoxMarkerSymbolNumber11                BoxMarkerSymbol = 11
	BoxMarkerSymbol11                      BoxMarkerSymbol = "11"
	BoxMarkerSymbolTriangleSw              BoxMarkerSymbol = "triangle-sw"
	BoxMarkerSymbolNumber111               BoxMarkerSymbol = 111
	BoxMarkerSymbol111                     BoxMarkerSymbol = "111"
	BoxMarkerSymbolTriangleSwOpen          BoxMarkerSymbol = "triangle-sw-open"
	BoxMarkerSymbolNumber211               BoxMarkerSymbol = 211
	BoxMarkerSymbol211                     BoxMarkerSymbol = "211"
	BoxMarkerSymbolTriangleSwDot           BoxMarkerSymbol = "triangle-sw-dot"
	BoxMarkerSymbolNumber311               BoxMarkerSymbol = 311
	BoxMarkerSymbol311                     BoxMarkerSymbol = "311"
	BoxMarkerSymbolTriangleSwOpenDot       BoxMarkerSymbol = "triangle-sw-open-dot"
	BoxMarkerSymbolNumber12                BoxMarkerSymbol = 12
	BoxMarkerSymbol12                      BoxMarkerSymbol = "12"
	BoxMarkerSymbolTriangleNw              BoxMarkerSymbol = "triangle-nw"
	BoxMarkerSymbolNumber112               BoxMarkerSymbol = 112
	BoxMarkerSymbol112                     BoxMarkerSymbol = "112"
	BoxMarkerSymbolTriangleNwOpen          BoxMarkerSymbol = "triangle-nw-open"
	BoxMarkerSymbolNumber212               BoxMarkerSymbol = 212
	BoxMarkerSymbol212                     BoxMarkerSymbol = "212"
	BoxMarkerSymbolTriangleNwDot           BoxMarkerSymbol = "triangle-nw-dot"
	BoxMarkerSymbolNumber312               BoxMarkerSymbol = 312
	BoxMarkerSymbol312                     BoxMarkerSymbol = "312"
	BoxMarkerSymbolTriangleNwOpenDot       BoxMarkerSymbol = "triangle-nw-open-dot"
	BoxMarkerSymbolNumber13                BoxMarkerSymbol = 13
	BoxMarkerSymbol13                      BoxMarkerSymbol = "13"
	BoxMarkerSymbolPentagon                BoxMarkerSymbol = "pentagon"
	BoxMarkerSymbolNumber113               BoxMarkerSymbol = 113
	BoxMarkerSymbol113                     BoxMarkerSymbol = "113"
	BoxMarkerSymbolPentagonOpen            BoxMarkerSymbol = "pentagon-open"
	BoxMarkerSymbolNumber213               BoxMarkerSymbol = 213
	BoxMarkerSymbol213                     BoxMarkerSymbol = "213"
	BoxMarkerSymbolPentagonDot             BoxMarkerSymbol = "pentagon-dot"
	BoxMarkerSymbolNumber313               BoxMarkerSymbol = 313
	BoxMarkerSymbol313                     BoxMarkerSymbol = "313"
	BoxMarkerSymbolPentagonOpenDot         BoxMarkerSymbol = "pentagon-open-dot"
	BoxMarkerSymbolNumber14                BoxMarkerSymbol = 14
	BoxMarkerSymbol14                      BoxMarkerSymbol = "14"
	BoxMarkerSymbolHexagon                 BoxMarkerSymbol = "hexagon"
	BoxMarkerSymbolNumber114               BoxMarkerSymbol = 114
	BoxMarkerSymbol114                     BoxMarkerSymbol = "114"
	BoxMarkerSymbolHexagonOpen             BoxMarkerSymbol = "hexagon-open"
	BoxMarkerSymbolNumber214               BoxMarkerSymbol = 214
	BoxMarkerSymbol214                     BoxMarkerSymbol = "214"
	BoxMarkerSymbolHexagonDot              BoxMarkerSymbol = "hexagon-dot"
	BoxMarkerSymbolNumber314               BoxMarkerSymbol = 314
	BoxMarkerSymbol314                     BoxMarkerSymbol = "314"
	BoxMarkerSymbolHexagonOpenDot          BoxMarkerSymbol = "hexagon-open-dot"
	BoxMarkerSymbolNumber15                BoxMarkerSymbol = 15
	BoxMarkerSymbol15                      BoxMarkerSymbol = "15"
	BoxMarkerSymbolHexagon2                BoxMarkerSymbol = "hexagon2"
	BoxMarkerSymbolNumber115               BoxMarkerSymbol = 115
	BoxMarkerSymbol115                     BoxMarkerSymbol = "115"
	BoxMarkerSymbolHexagon2Open            BoxMarkerSymbol = "hexagon2-open"
	BoxMarkerSymbolNumber215               BoxMarkerSymbol = 215
	BoxMarkerSymbol215                     BoxMarkerSymbol = "215"
	BoxMarkerSymbolHexagon2Dot             BoxMarkerSymbol = "hexagon2-dot"
	BoxMarkerSymbolNumber315               BoxMarkerSymbol = 315
	BoxMarkerSymbol315                     BoxMarkerSymbol = "315"
	BoxMarkerSymbolHexagon2OpenDot         BoxMarkerSymbol = "hexagon2-open-dot"
	BoxMarkerSymbolNumber16                BoxMarkerSymbol = 16
	BoxMarkerSymbol16                      BoxMarkerSymbol = "16"
	BoxMarkerSymbolOctagon                 BoxMarkerSymbol = "octagon"
	BoxMarkerSymbolNumber116               BoxMarkerSymbol = 116
	BoxMarkerSymbol116                     BoxMarkerSymbol = "116"
	BoxMarkerSymbolOctagonOpen             BoxMarkerSymbol = "octagon-open"
	BoxMarkerSymbolNumber216               BoxMarkerSymbol = 216
	BoxMarkerSymbol216                     BoxMarkerSymbol = "216"
	BoxMarkerSymbolOctagonDot              BoxMarkerSymbol = "octagon-dot"
	BoxMarkerSymbolNumber316               BoxMarkerSymbol = 316
	BoxMarkerSymbol316                     BoxMarkerSymbol = "316"
	BoxMarkerSymbolOctagonOpenDot          BoxMarkerSymbol = "octagon-open-dot"
	BoxMarkerSymbolNumber17                BoxMarkerSymbol = 17
	BoxMarkerSymbol17                      BoxMarkerSymbol = "17"
	BoxMarkerSymbolStar                    BoxMarkerSymbol = "star"
	BoxMarkerSymbolNumber117               BoxMarkerSymbol = 117
	BoxMarkerSymbol117                     BoxMarkerSymbol = "117"
	BoxMarkerSymbolStarOpen                BoxMarkerSymbol = "star-open"
	BoxMarkerSymbolNumber217               BoxMarkerSymbol = 217
	BoxMarkerSymbol217                     BoxMarkerSymbol = "217"
	BoxMarkerSymbolStarDot                 BoxMarkerSymbol = "star-dot"
	BoxMarkerSymbolNumber317               BoxMarkerSymbol = 317
	BoxMarkerSymbol317                     BoxMarkerSymbol = "317"
	BoxMarkerSymbolStarOpenDot             BoxMarkerSymbol = "star-open-dot"
	BoxMarkerSymbolNumber18                BoxMarkerSymbol = 18
	BoxMarkerSymbol18                      BoxMarkerSymbol = "18"
	BoxMarkerSymbolHexagram                BoxMarkerSymbol = "hexagram"
	BoxMarkerSymbolNumber118               BoxMarkerSymbol = 118
	BoxMarkerSymbol118                     BoxMarkerSymbol = "118"
	BoxMarkerSymbolHexagramOpen            BoxMarkerSymbol = "hexagram-open"
	BoxMarkerSymbolNumber218               BoxMarkerSymbol = 218
	BoxMarkerSymbol218                     BoxMarkerSymbol = "218"
	BoxMarkerSymbolHexagramDot             BoxMarkerSymbol = "hexagram-dot"
	BoxMarkerSymbolNumber318               BoxMarkerSymbol = 318
	BoxMarkerSymbol318                     BoxMarkerSymbol = "318"
	BoxMarkerSymbolHexagramOpenDot         BoxMarkerSymbol = "hexagram-open-dot"
	BoxMarkerSymbolNumber19                BoxMarkerSymbol = 19
	BoxMarkerSymbol19                      BoxMarkerSymbol = "19"
	BoxMarkerSymbolStarTriangleUp          BoxMarkerSymbol = "star-triangle-up"
	BoxMarkerSymbolNumber119               BoxMarkerSymbol = 119
	BoxMarkerSymbol119                     BoxMarkerSymbol = "119"
	BoxMarkerSymbolStarTriangleUpOpen      BoxMarkerSymbol = "star-triangle-up-open"
	BoxMarkerSymbolNumber219               BoxMarkerSymbol = 219
	BoxMarkerSymbol219                     BoxMarkerSymbol = "219"
	BoxMarkerSymbolStarTriangleUpDot       BoxMarkerSymbol = "star-triangle-up-dot"
	BoxMarkerSymbolNumber319               BoxMarkerSymbol = 319
	BoxMarkerSymbol319                     BoxMarkerSymbol = "319"
	BoxMarkerSymbolStarTriangleUpOpenDot   BoxMarkerSymbol = "star-triangle-up-open-dot"
	BoxMarkerSymbolNumber20                BoxMarkerSymbol = 20
	BoxMarkerSymbol20                      BoxMarkerSymbol = "20"
	BoxMarkerSymbolStarTriangleDown        BoxMarkerSymbol = "star-triangle-down"
	BoxMarkerSymbolNumber120               BoxMarkerSymbol = 120
	BoxMarkerSymbol120                     BoxMarkerSymbol = "120"
	BoxMarkerSymbolStarTriangleDownOpen    BoxMarkerSymbol = "star-triangle-down-open"
	BoxMarkerSymbolNumber220               BoxMarkerSymbol = 220
	BoxMarkerSymbol220                     BoxMarkerSymbol = "220"
	BoxMarkerSymbolStarTriangleDownDot     BoxMarkerSymbol = "star-triangle-down-dot"
	BoxMarkerSymbolNumber320               BoxMarkerSymbol = 320
	BoxMarkerSymbol320                     BoxMarkerSymbol = "320"
	BoxMarkerSymbolStarTriangleDownOpenDot BoxMarkerSymbol = "star-triangle-down-open-dot"
	BoxMarkerSymbolNumber21                BoxMarkerSymbol = 21
	BoxMarkerSymbol21                      BoxMarkerSymbol = "21"
	BoxMarkerSymbolStarSquare              BoxMarkerSymbol = "star-square"
	BoxMarkerSymbolNumber121               BoxMarkerSymbol = 121
	BoxMarkerSymbol121                     BoxMarkerSymbol = "121"
	BoxMarkerSymbolStarSquareOpen          BoxMarkerSymbol = "star-square-open"
	BoxMarkerSymbolNumber221               BoxMarkerSymbol = 221
	BoxMarkerSymbol221                     BoxMarkerSymbol = "221"
	BoxMarkerSymbolStarSquareDot           BoxMarkerSymbol = "star-square-dot"
	BoxMarkerSymbolNumber321               BoxMarkerSymbol = 321
	BoxMarkerSymbol321                     BoxMarkerSymbol = "321"
	BoxMarkerSymbolStarSquareOpenDot       BoxMarkerSymbol = "star-square-open-dot"
	BoxMarkerSymbolNumber22                BoxMarkerSymbol = 22
	BoxMarkerSymbol22                      BoxMarkerSymbol = "22"
	BoxMarkerSymbolStarDiamond             BoxMarkerSymbol = "star-diamond"
	BoxMarkerSymbolNumber122               BoxMarkerSymbol = 122
	BoxMarkerSymbol122                     BoxMarkerSymbol = "122"
	BoxMarkerSymbolStarDiamondOpen         BoxMarkerSymbol = "star-diamond-open"
	BoxMarkerSymbolNumber222               BoxMarkerSymbol = 222
	BoxMarkerSymbol222                     BoxMarkerSymbol = "222"
	BoxMarkerSymbolStarDiamondDot          BoxMarkerSymbol = "star-diamond-dot"
	BoxMarkerSymbolNumber322               BoxMarkerSymbol = 322
	BoxMarkerSymbol322                     BoxMarkerSymbol = "322"
	BoxMarkerSymbolStarDiamondOpenDot      BoxMarkerSymbol = "star-diamond-open-dot"
	BoxMarkerSymbolNumber23                BoxMarkerSymbol = 23
	BoxMarkerSymbol23                      BoxMarkerSymbol = "23"
	BoxMarkerSymbolDiamondTall             BoxMarkerSymbol = "diamond-tall"
	BoxMarkerSymbolNumber123               BoxMarkerSymbol = 123
	BoxMarkerSymbol123                     BoxMarkerSymbol = "123"
	BoxMarkerSymbolDiamondTallOpen         BoxMarkerSymbol = "diamond-tall-open"
	BoxMarkerSymbolNumber223               BoxMarkerSymbol = 223
	BoxMarkerSymbol223                     BoxMarkerSymbol = "223"
	BoxMarkerSymbolDiamondTallDot          BoxMarkerSymbol = "diamond-tall-dot"
	BoxMarkerSymbolNumber323               BoxMarkerSymbol = 323
	BoxMarkerSymbol323                     BoxMarkerSymbol = "323"
	BoxMarkerSymbolDiamondTallOpenDot      BoxMarkerSymbol = "diamond-tall-open-dot"
	BoxMarkerSymbolNumber24                BoxMarkerSymbol = 24
	BoxMarkerSymbol24                      BoxMarkerSymbol = "24"
	BoxMarkerSymbolDiamondWide             BoxMarkerSymbol = "diamond-wide"
	BoxMarkerSymbolNumber124               BoxMarkerSymbol = 124
	BoxMarkerSymbol124                     BoxMarkerSymbol = "124"
	BoxMarkerSymbolDiamondWideOpen         BoxMarkerSymbol = "diamond-wide-open"
	BoxMarkerSymbolNumber224               BoxMarkerSymbol = 224
	BoxMarkerSymbol224                     BoxMarkerSymbol = "224"
	BoxMarkerSymbolDiamondWideDot          BoxMarkerSymbol = "diamond-wide-dot"
	BoxMarkerSymbolNumber324               BoxMarkerSymbol = 324
	BoxMarkerSymbol324                     BoxMarkerSymbol = "324"
	BoxMarkerSymbolDiamondWideOpenDot      BoxMarkerSymbol = "diamond-wide-open-dot"
	BoxMarkerSymbolNumber25                BoxMarkerSymbol = 25
	BoxMarkerSymbol25                      BoxMarkerSymbol = "25"
	BoxMarkerSymbolHourglass               BoxMarkerSymbol = "hourglass"
	BoxMarkerSymbolNumber125               BoxMarkerSymbol = 125
	BoxMarkerSymbol125                     BoxMarkerSymbol = "125"
	BoxMarkerSymbolHourglassOpen           BoxMarkerSymbol = "hourglass-open"
	BoxMarkerSymbolNumber26                BoxMarkerSymbol = 26
	BoxMarkerSymbol26                      BoxMarkerSymbol = "26"
	BoxMarkerSymbolBowtie                  BoxMarkerSymbol = "bowtie"
	BoxMarkerSymbolNumber126               BoxMarkerSymbol = 126
	BoxMarkerSymbol126                     BoxMarkerSymbol = "126"
	BoxMarkerSymbolBowtieOpen              BoxMarkerSymbol = "bowtie-open"
	BoxMarkerSymbolNumber27                BoxMarkerSymbol = 27
	BoxMarkerSymbol27                      BoxMarkerSymbol = "27"
	BoxMarkerSymbolCircleCross             BoxMarkerSymbol = "circle-cross"
	BoxMarkerSymbolNumber127               BoxMarkerSymbol = 127
	BoxMarkerSymbol127                     BoxMarkerSymbol = "127"
	BoxMarkerSymbolCircleCrossOpen         BoxMarkerSymbol = "circle-cross-open"
	BoxMarkerSymbolNumber28                BoxMarkerSymbol = 28
	BoxMarkerSymbol28                      BoxMarkerSymbol = "28"
	BoxMarkerSymbolCircleX                 BoxMarkerSymbol = "circle-x"
	BoxMarkerSymbolNumber128               BoxMarkerSymbol = 128
	BoxMarkerSymbol128                     BoxMarkerSymbol = "128"
	BoxMarkerSymbolCircleXOpen             BoxMarkerSymbol = "circle-x-open"
	BoxMarkerSymbolNumber29                BoxMarkerSymbol = 29
	BoxMarkerSymbol29                      BoxMarkerSymbol = "29"
	BoxMarkerSymbolSquareCross             BoxMarkerSymbol = "square-cross"
	BoxMarkerSymbolNumber129               BoxMarkerSymbol = 129
	BoxMarkerSymbol129                     BoxMarkerSymbol = "129"
	BoxMarkerSymbolSquareCrossOpen         BoxMarkerSymbol = "square-cross-open"
	BoxMarkerSymbolNumber30                BoxMarkerSymbol = 30
	BoxMarkerSymbol30                      BoxMarkerSymbol = "30"
	BoxMarkerSymbolSquareX                 BoxMarkerSymbol = "square-x"
	BoxMarkerSymbolNumber130               BoxMarkerSymbol = 130
	BoxMarkerSymbol130                     BoxMarkerSymbol = "130"
	BoxMarkerSymbolSquareXOpen             BoxMarkerSymbol = "square-x-open"
	BoxMarkerSymbolNumber31                BoxMarkerSymbol = 31
	BoxMarkerSymbol31                      BoxMarkerSymbol = "31"
	BoxMarkerSymbolDiamondCross            BoxMarkerSymbol = "diamond-cross"
	BoxMarkerSymbolNumber131               BoxMarkerSymbol = 131
	BoxMarkerSymbol131                     BoxMarkerSymbol = "131"
	BoxMarkerSymbolDiamondCrossOpen        BoxMarkerSymbol = "diamond-cross-open"
	BoxMarkerSymbolNumber32                BoxMarkerSymbol = 32
	BoxMarkerSymbol32                      BoxMarkerSymbol = "32"
	BoxMarkerSymbolDiamondX                BoxMarkerSymbol = "diamond-x"
	BoxMarkerSymbolNumber132               BoxMarkerSymbol = 132
	BoxMarkerSymbol132                     BoxMarkerSymbol = "132"
	BoxMarkerSymbolDiamondXOpen            BoxMarkerSymbol = "diamond-x-open"
	BoxMarkerSymbolNumber33                BoxMarkerSymbol = 33
	BoxMarkerSymbol33                      BoxMarkerSymbol = "33"
	BoxMarkerSymbolCrossThin               BoxMarkerSymbol = "cross-thin"
	BoxMarkerSymbolNumber133               BoxMarkerSymbol = 133
	BoxMarkerSymbol133                     BoxMarkerSymbol = "133"
	BoxMarkerSymbolCrossThinOpen           BoxMarkerSymbol = "cross-thin-open"
	BoxMarkerSymbolNumber34                BoxMarkerSymbol = 34
	BoxMarkerSymbol34                      BoxMarkerSymbol = "34"
	BoxMarkerSymbolXThin                   BoxMarkerSymbol = "x-thin"
	BoxMarkerSymbolNumber134               BoxMarkerSymbol = 134
	BoxMarkerSymbol134                     BoxMarkerSymbol = "134"
	BoxMarkerSymbolXThinOpen               BoxMarkerSymbol = "x-thin-open"
	BoxMarkerSymbolNumber35                BoxMarkerSymbol = 35
	BoxMarkerSymbol35                      BoxMarkerSymbol = "35"
	BoxMarkerSymbolAsterisk                BoxMarkerSymbol = "asterisk"
	BoxMarkerSymbolNumber135               BoxMarkerSymbol = 135
	BoxMarkerSymbol135                     BoxMarkerSymbol = "135"
	BoxMarkerSymbolAsteriskOpen            BoxMarkerSymbol = "asterisk-open"
	BoxMarkerSymbolNumber36                BoxMarkerSymbol = 36
	BoxMarkerSymbol36                      BoxMarkerSymbol = "36"
	BoxMarkerSymbolHash                    BoxMarkerSymbol = "hash"
	BoxMarkerSymbolNumber136               BoxMarkerSymbol = 136
	BoxMarkerSymbol136                     BoxMarkerSymbol = "136"
	BoxMarkerSymbolHashOpen                BoxMarkerSymbol = "hash-open"
	BoxMarkerSymbolNumber236               BoxMarkerSymbol = 236
	BoxMarkerSymbol236                     BoxMarkerSymbol = "236"
	BoxMarkerSymbolHashDot                 BoxMarkerSymbol = "hash-dot"
	BoxMarkerSymbolNumber336               BoxMarkerSymbol = 336
	BoxMarkerSymbol336                     BoxMarkerSymbol = "336"
	BoxMarkerSymbolHashOpenDot             BoxMarkerSymbol = "hash-open-dot"
	BoxMarkerSymbolNumber37                BoxMarkerSymbol = 37
	BoxMarkerSymbol37                      BoxMarkerSymbol = "37"
	BoxMarkerSymbolYUp                     BoxMarkerSymbol = "y-up"
	BoxMarkerSymbolNumber137               BoxMarkerSymbol = 137
	BoxMarkerSymbol137                     BoxMarkerSymbol = "137"
	BoxMarkerSymbolYUpOpen                 BoxMarkerSymbol = "y-up-open"
	BoxMarkerSymbolNumber38                BoxMarkerSymbol = 38
	BoxMarkerSymbol38                      BoxMarkerSymbol = "38"
	BoxMarkerSymbolYDown                   BoxMarkerSymbol = "y-down"
	BoxMarkerSymbolNumber138               BoxMarkerSymbol = 138
	BoxMarkerSymbol138                     BoxMarkerSymbol = "138"
	BoxMarkerSymbolYDownOpen               BoxMarkerSymbol = "y-down-open"
	BoxMarkerSymbolNumber39                BoxMarkerSymbol = 39
	BoxMarkerSymbol39                      BoxMarkerSymbol = "39"
	BoxMarkerSymbolYLeft                   BoxMarkerSymbol = "y-left"
	BoxMarkerSymbolNumber139               BoxMarkerSymbol = 139
	BoxMarkerSymbol139                     BoxMarkerSymbol = "139"
	BoxMarkerSymbolYLeftOpen               BoxMarkerSymbol = "y-left-open"
	BoxMarkerSymbolNumber40                BoxMarkerSymbol = 40
	BoxMarkerSymbol40                      BoxMarkerSymbol = "40"
	BoxMarkerSymbolYRight                  BoxMarkerSymbol = "y-right"
	BoxMarkerSymbolNumber140               BoxMarkerSymbol = 140
	BoxMarkerSymbol140                     BoxMarkerSymbol = "140"
	BoxMarkerSymbolYRightOpen              BoxMarkerSymbol = "y-right-open"
	BoxMarkerSymbolNumber41                BoxMarkerSymbol = 41
	BoxMarkerSymbol41                      BoxMarkerSymbol = "41"
	BoxMarkerSymbolLineEw                  BoxMarkerSymbol = "line-ew"
	BoxMarkerSymbolNumber141               BoxMarkerSymbol = 141
	BoxMarkerSymbol141                     BoxMarkerSymbol = "141"
	BoxMarkerSymbolLineEwOpen              BoxMarkerSymbol = "line-ew-open"
	BoxMarkerSymbolNumber42                BoxMarkerSymbol = 42
	BoxMarkerSymbol42                      BoxMarkerSymbol = "42"
	BoxMarkerSymbolLineNs                  BoxMarkerSymbol = "line-ns"
	BoxMarkerSymbolNumber142               BoxMarkerSymbol = 142
	BoxMarkerSymbol142                     BoxMarkerSymbol = "142"
	BoxMarkerSymbolLineNsOpen              BoxMarkerSymbol = "line-ns-open"
	BoxMarkerSymbolNumber43                BoxMarkerSymbol = 43
	BoxMarkerSymbol43                      BoxMarkerSymbol = "43"
	BoxMarkerSymbolLineNe                  BoxMarkerSymbol = "line-ne"
	BoxMarkerSymbolNumber143               BoxMarkerSymbol = 143
	BoxMarkerSymbol143                     BoxMarkerSymbol = "143"
	BoxMarkerSymbolLineNeOpen              BoxMarkerSymbol = "line-ne-open"
	BoxMarkerSymbolNumber44                BoxMarkerSymbol = 44
	BoxMarkerSymbol44                      BoxMarkerSymbol = "44"
	BoxMarkerSymbolLineNw                  BoxMarkerSymbol = "line-nw"
	BoxMarkerSymbolNumber144               BoxMarkerSymbol = 144
	BoxMarkerSymbol144                     BoxMarkerSymbol = "144"
	BoxMarkerSymbolLineNwOpen              BoxMarkerSymbol = "line-nw-open"
	BoxMarkerSymbolNumber45                BoxMarkerSymbol = 45
	BoxMarkerSymbol45                      BoxMarkerSymbol = "45"
	BoxMarkerSymbolArrowUp                 BoxMarkerSymbol = "arrow-up"
	BoxMarkerSymbolNumber145               BoxMarkerSymbol = 145
	BoxMarkerSymbol145                     BoxMarkerSymbol = "145"
	BoxMarkerSymbolArrowUpOpen             BoxMarkerSymbol = "arrow-up-open"
	BoxMarkerSymbolNumber46                BoxMarkerSymbol = 46
	BoxMarkerSymbol46                      BoxMarkerSymbol = "46"
	BoxMarkerSymbolArrowDown               BoxMarkerSymbol = "arrow-down"
	BoxMarkerSymbolNumber146               BoxMarkerSymbol = 146
	BoxMarkerSymbol146                     BoxMarkerSymbol = "146"
	BoxMarkerSymbolArrowDownOpen           BoxMarkerSymbol = "arrow-down-open"
	BoxMarkerSymbolNumber47                BoxMarkerSymbol = 47
	BoxMarkerSymbol47                      BoxMarkerSymbol = "47"
	BoxMarkerSymbolArrowLeft               BoxMarkerSymbol = "arrow-left"
	BoxMarkerSymbolNumber147               BoxMarkerSymbol = 147
	BoxMarkerSymbol147                     BoxMarkerSymbol = "147"
	BoxMarkerSymbolArrowLeftOpen           BoxMarkerSymbol = "arrow-left-open"
	BoxMarkerSymbolNumber48                BoxMarkerSymbol = 48
	BoxMarkerSymbol48                      BoxMarkerSymbol = "48"
	BoxMarkerSymbolArrowRight              BoxMarkerSymbol = "arrow-right"
	BoxMarkerSymbolNumber148               BoxMarkerSymbol = 148
	BoxMarkerSymbol148                     BoxMarkerSymbol = "148"
	BoxMarkerSymbolArrowRightOpen          BoxMarkerSymbol = "arrow-right-open"
	BoxMarkerSymbolNumber49                BoxMarkerSymbol = 49
	BoxMarkerSymbol49                      BoxMarkerSymbol = "49"
	BoxMarkerSymbolArrowBarUp              BoxMarkerSymbol = "arrow-bar-up"
	BoxMarkerSymbolNumber149               BoxMarkerSymbol = 149
	BoxMarkerSymbol149                     BoxMarkerSymbol = "149"
	BoxMarkerSymbolArrowBarUpOpen          BoxMarkerSymbol = "arrow-bar-up-open"
	BoxMarkerSymbolNumber50                BoxMarkerSymbol = 50
	BoxMarkerSymbol50                      BoxMarkerSymbol = "50"
	BoxMarkerSymbolArrowBarDown            BoxMarkerSymbol = "arrow-bar-down"
	BoxMarkerSymbolNumber150               BoxMarkerSymbol = 150
	BoxMarkerSymbol150                     BoxMarkerSymbol = "150"
	BoxMarkerSymbolArrowBarDownOpen        BoxMarkerSymbol = "arrow-bar-down-open"
	BoxMarkerSymbolNumber51                BoxMarkerSymbol = 51
	BoxMarkerSymbol51                      BoxMarkerSymbol = "51"
	BoxMarkerSymbolArrowBarLeft            BoxMarkerSymbol = "arrow-bar-left"
	BoxMarkerSymbolNumber151               BoxMarkerSymbol = 151
	BoxMarkerSymbol151                     BoxMarkerSymbol = "151"
	BoxMarkerSymbolArrowBarLeftOpen        BoxMarkerSymbol = "arrow-bar-left-open"
	BoxMarkerSymbolNumber52                BoxMarkerSymbol = 52
	BoxMarkerSymbol52                      BoxMarkerSymbol = "52"
	BoxMarkerSymbolArrowBarRight           BoxMarkerSymbol = "arrow-bar-right"
	BoxMarkerSymbolNumber152               BoxMarkerSymbol = 152
	BoxMarkerSymbol152                     BoxMarkerSymbol = "152"
	BoxMarkerSymbolArrowBarRightOpen       BoxMarkerSymbol = "arrow-bar-right-open"
)

// BoxOrientation Sets the orientation of the box(es). If *v* (*h*), the distribution is visualized along the vertical (horizontal).
type BoxOrientation string

const (
	BoxOrientationV BoxOrientation = "v"
	BoxOrientationH BoxOrientation = "h"
)

// BoxQuartilemethod Sets the method used to compute the sample's Q1 and Q3 quartiles. The *linear* method uses the 25th percentile for Q1 and 75th percentile for Q3 as computed using method #10 (listed on http://www.amstat.org/publications/jse/v14n3/langford.html). The *exclusive* method uses the median to divide the ordered dataset into two halves if the sample is odd, it does not include the median in either half - Q1 is then the median of the lower half and Q3 the median of the upper half. The *inclusive* method also uses the median to divide the ordered dataset into two halves but if the sample is odd, it includes the median in both halves - Q1 is then the median of the lower half and Q3 the median of the upper half.
type BoxQuartilemethod string

const (
	BoxQuartilemethodLinear    BoxQuartilemethod = "linear"
	BoxQuartilemethodExclusive BoxQuartilemethod = "exclusive"
	BoxQuartilemethodInclusive BoxQuartilemethod = "inclusive"
)

// BoxVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type BoxVisible interface{}

var (
	BoxVisibleTrue       BoxVisible = true
	BoxVisibleFalse      BoxVisible = false
	BoxVisibleLegendonly BoxVisible = "legendonly"
)

// BoxXcalendar Sets the calendar system to use with `x` date data.
type BoxXcalendar string

const (
	BoxXcalendarGregorian  BoxXcalendar = "gregorian"
	BoxXcalendarChinese    BoxXcalendar = "chinese"
	BoxXcalendarCoptic     BoxXcalendar = "coptic"
	BoxXcalendarDiscworld  BoxXcalendar = "discworld"
	BoxXcalendarEthiopian  BoxXcalendar = "ethiopian"
	BoxXcalendarHebrew     BoxXcalendar = "hebrew"
	BoxXcalendarIslamic    BoxXcalendar = "islamic"
	BoxXcalendarJulian     BoxXcalendar = "julian"
	BoxXcalendarMayan      BoxXcalendar = "mayan"
	BoxXcalendarNanakshahi BoxXcalendar = "nanakshahi"
	BoxXcalendarNepali     BoxXcalendar = "nepali"
	BoxXcalendarPersian    BoxXcalendar = "persian"
	BoxXcalendarJalali     BoxXcalendar = "jalali"
	BoxXcalendarTaiwan     BoxXcalendar = "taiwan"
	BoxXcalendarThai       BoxXcalendar = "thai"
	BoxXcalendarUmmalqura  BoxXcalendar = "ummalqura"
)

// BoxXperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
type BoxXperiodalignment string

const (
	BoxXperiodalignmentStart  BoxXperiodalignment = "start"
	BoxXperiodalignmentMiddle BoxXperiodalignment = "middle"
	BoxXperiodalignmentEnd    BoxXperiodalignment = "end"
)

// BoxYcalendar Sets the calendar system to use with `y` date data.
type BoxYcalendar string

const (
	BoxYcalendarGregorian  BoxYcalendar = "gregorian"
	BoxYcalendarChinese    BoxYcalendar = "chinese"
	BoxYcalendarCoptic     BoxYcalendar = "coptic"
	BoxYcalendarDiscworld  BoxYcalendar = "discworld"
	BoxYcalendarEthiopian  BoxYcalendar = "ethiopian"
	BoxYcalendarHebrew     BoxYcalendar = "hebrew"
	BoxYcalendarIslamic    BoxYcalendar = "islamic"
	BoxYcalendarJulian     BoxYcalendar = "julian"
	BoxYcalendarMayan      BoxYcalendar = "mayan"
	BoxYcalendarNanakshahi BoxYcalendar = "nanakshahi"
	BoxYcalendarNepali     BoxYcalendar = "nepali"
	BoxYcalendarPersian    BoxYcalendar = "persian"
	BoxYcalendarJalali     BoxYcalendar = "jalali"
	BoxYcalendarTaiwan     BoxYcalendar = "taiwan"
	BoxYcalendarThai       BoxYcalendar = "thai"
	BoxYcalendarUmmalqura  BoxYcalendar = "ummalqura"
)

// BoxYperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
type BoxYperiodalignment string

const (
	BoxYperiodalignmentStart  BoxYperiodalignment = "start"
	BoxYperiodalignmentMiddle BoxYperiodalignment = "middle"
	BoxYperiodalignmentEnd    BoxYperiodalignment = "end"
)

// BoxHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type BoxHoverinfo string

const (
	// Flags
	BoxHoverinfoX    BoxHoverinfo = "x"
	BoxHoverinfoY    BoxHoverinfo = "y"
	BoxHoverinfoZ    BoxHoverinfo = "z"
	BoxHoverinfoText BoxHoverinfo = "text"
	BoxHoverinfoName BoxHoverinfo = "name"

	// Extra
	BoxHoverinfoAll  BoxHoverinfo = "all"
	BoxHoverinfoNone BoxHoverinfo = "none"
	BoxHoverinfoSkip BoxHoverinfo = "skip"
)

// BoxHoveron Do the hover effects highlight individual boxes  or sample points or both?
type BoxHoveron string

const (
	// Flags
	BoxHoveronBoxes  BoxHoveron = "boxes"
	BoxHoveronPoints BoxHoveron = "points"

	// Extra

)
