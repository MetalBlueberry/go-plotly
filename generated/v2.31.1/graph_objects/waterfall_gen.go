package grob

// Code generated by go-plotly/generator. DO NOT EDIT.

var TraceTypeWaterfall TraceType = "waterfall"

func (trace *Waterfall) GetType() TraceType {
	return TraceTypeWaterfall
}

// Waterfall Draws waterfall trace which is useful graph to displays the contribution of various elements (either positive or negative) in a bar chart. The data visualized by the span of the bars is set in `y` if `orientation` is set to *v* (the default) and the labels are set in `x`. By setting `orientation` to *h*, the roles are interchanged.
type Waterfall struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Alignmentgroup
	// arrayOK: false
	// type: string
	// Set several traces linked to the same position axis or matching axes to the same alignmentgroup. This controls whether bars compute their positional range dependently or independently.
	Alignmentgroup String `json:"alignmentgroup,omitempty"`

	// Base
	// arrayOK: false
	// type: number
	// Sets where the bar base is drawn (in position axis units).
	Base float64 `json:"base,omitempty"`

	// Cliponaxis
	// arrayOK: false
	// type: boolean
	// Determines whether the text nodes are clipped about the subplot axes. To show the text nodes above axis lines and tick labels, make sure to set `xaxis.layer` and `yaxis.layer` to *below traces*.
	Cliponaxis Bool `json:"cliponaxis,omitempty"`

	// Connector
	// role: Object
	Connector *WaterfallConnector `json:"connector,omitempty"`

	// Constraintext
	// default: both
	// type: enumerated
	// Constrain the size of text inside or outside a bar to be no larger than the bar itself.
	Constraintext WaterfallConstraintext `json:"constraintext,omitempty"`

	// Customdata
	// arrayOK: false
	// type: data_array
	// Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `customdata`.
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Decreasing
	// role: Object
	Decreasing *WaterfallDecreasing `json:"decreasing,omitempty"`

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

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo interface{} `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hoverinfo`.
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *WaterfallHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate
	// arrayOK: true
	// type: string
	// Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}" as well as %{xother}, {%_xother}, {%_xother_}, {%xother_}. When showing info for several points, *xother* will be added to those with different x positions from the first point. An underscore before or after *(x|y)other* will add a space on that side, only when this field is shown. Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-format/tree/v1.4.5#d3-format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format/tree/v2.2.3#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. Finally, the template string has access to variables `initial`, `delta` and `final`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hovertemplate`.
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext
	// arrayOK: true
	// type: string
	// Sets hover text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hovertext`.
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids
	// arrayOK: false
	// type: data_array
	// Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `ids`.
	Idssrc String `json:"idssrc,omitempty"`

	// Increasing
	// role: Object
	Increasing *WaterfallIncreasing `json:"increasing,omitempty"`

	// Insidetextanchor
	// default: end
	// type: enumerated
	// Determines if texts are kept at center or start/end points in `textposition` *inside* mode.
	Insidetextanchor WaterfallInsidetextanchor `json:"insidetextanchor,omitempty"`

	// Insidetextfont
	// role: Object
	Insidetextfont *WaterfallInsidetextfont `json:"insidetextfont,omitempty"`

	// Legend
	// arrayOK: false
	// type: subplotid
	// Sets the reference to a legend to show this trace in. References to these legends are *legend*, *legend2*, *legend3*, etc. Settings for these legends are set in the layout, under `layout.legend`, `layout.legend2`, etc.
	Legend String `json:"legend,omitempty"`

	// Legendgroup
	// arrayOK: false
	// type: string
	// Sets the legend group for this trace. Traces and shapes part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Legendgrouptitle
	// role: Object
	Legendgrouptitle *WaterfallLegendgrouptitle `json:"legendgrouptitle,omitempty"`

	// Legendrank
	// arrayOK: false
	// type: number
	// Sets the legend rank for this trace. Items and groups with smaller ranks are presented on top/left side while with *reversed* `legend.traceorder` they are on bottom/right side. The default legendrank is 1000, so that you can use ranks less than 1000 to place certain items before all unranked items, and ranks greater than 1000 to go after all unranked items. When having unranked or equal rank items shapes would be displayed after traces i.e. according to their order in data and layout.
	Legendrank float64 `json:"legendrank,omitempty"`

	// Legendwidth
	// arrayOK: false
	// type: number
	// Sets the width (in px or fraction) of the legend for this trace.
	Legendwidth float64 `json:"legendwidth,omitempty"`

	// Measure
	// arrayOK: false
	// type: data_array
	// An array containing types of values. By default the values are considered as 'relative'. However; it is possible to use 'total' to compute the sums. Also 'absolute' could be applied to reset the computed total or to declare an initial value where needed.
	Measure interface{} `json:"measure,omitempty"`

	// Measuresrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `measure`.
	Measuresrc String `json:"measuresrc,omitempty"`

	// Meta
	// arrayOK: true
	// type: any
	// Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `meta`.
	Metasrc String `json:"metasrc,omitempty"`

	// Name
	// arrayOK: false
	// type: string
	// Sets the trace name. The trace name appears as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Offset
	// arrayOK: true
	// type: number
	// Shifts the position where the bar is drawn (in position axis units). In *group* barmode, traces that set *offset* will be excluded and drawn in *overlay* mode instead.
	Offset interface{} `json:"offset,omitempty"`

	// Offsetgroup
	// arrayOK: false
	// type: string
	// Set several traces linked to the same position axis or matching axes to the same offsetgroup where bars of the same position coordinate will line up.
	Offsetgroup String `json:"offsetgroup,omitempty"`

	// Offsetsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `offset`.
	Offsetsrc String `json:"offsetsrc,omitempty"`

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Orientation
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the orientation of the bars. With *v* (*h*), the value of the each bar spans along the vertical (horizontal).
	Orientation WaterfallOrientation `json:"orientation,omitempty"`

	// Outsidetextfont
	// role: Object
	Outsidetextfont *WaterfallOutsidetextfont `json:"outsidetextfont,omitempty"`

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
	Stream *WaterfallStream `json:"stream,omitempty"`

	// Text
	// arrayOK: true
	// type: string
	// Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textangle
	// arrayOK: false
	// type: angle
	// Sets the angle of the tick labels with respect to the bar. For example, a `tickangle` of -90 draws the tick labels vertically. With *auto* the texts may automatically be rotated to fit with the maximum size in bars.
	Textangle float64 `json:"textangle,omitempty"`

	// Textfont
	// role: Object
	Textfont *WaterfallTextfont `json:"textfont,omitempty"`

	// Textinfo
	// default: %!s(<nil>)
	// type: flaglist
	// Determines which trace information appear on the graph. In the case of having multiple waterfalls, totals are computed separately (per trace).
	Textinfo WaterfallTextinfo `json:"textinfo,omitempty"`

	// Textposition
	// default: auto
	// type: enumerated
	// Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside. If *none*, no text appears.
	Textposition interface{} `json:"textposition,omitempty"`

	// Textpositionsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `textposition`.
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `text`.
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate
	// arrayOK: true
	// type: string
	// Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-format/tree/v1.4.5#d3-format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format/tree/v2.2.3#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. Finally, the template string has access to variables `initial`, `delta`, `final` and `label`.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `texttemplate`.
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Totals
	// role: Object
	Totals *WaterfallTotals `json:"totals,omitempty"`

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
	Visible WaterfallVisible `json:"visible,omitempty"`

	// Width
	// arrayOK: true
	// type: number
	// Sets the bar width (in position axis units).
	Width interface{} `json:"width,omitempty"`

	// Widthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `width`.
	Widthsrc String `json:"widthsrc,omitempty"`

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

	// Xhoverformat
	// arrayOK: false
	// type: string
	// Sets the hover text formatting rulefor `x`  using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-format/tree/v1.4.5#d3-format. And for dates see: https://github.com/d3/d3-time-format/tree/v2.2.3#locale_format. We add two items to d3's date formatter: *%h* for half of the year as a decimal number as well as *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46*By default the values are formatted using `xaxis.hoverformat`.
	Xhoverformat String `json:"xhoverformat,omitempty"`

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
	Xperiodalignment WaterfallXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `x`.
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

	// Yhoverformat
	// arrayOK: false
	// type: string
	// Sets the hover text formatting rulefor `y`  using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-format/tree/v1.4.5#d3-format. And for dates see: https://github.com/d3/d3-time-format/tree/v2.2.3#locale_format. We add two items to d3's date formatter: *%h* for half of the year as a decimal number as well as *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46*By default the values are formatted using `yaxis.hoverformat`.
	Yhoverformat String `json:"yhoverformat,omitempty"`

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
	Yperiodalignment WaterfallYperiodalignment `json:"yperiodalignment,omitempty"`

	// Ysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `y`.
	Ysrc String `json:"ysrc,omitempty"`

	// Zorder
	// arrayOK: false
	// type: integer
	// Sets the layer on which this trace is displayed, relative to other SVG traces on the same subplot. SVG traces with higher `zorder` appear in front of those with lower `zorder`.
	Zorder int64 `json:"zorder,omitempty"`
}

// WaterfallConnectorLine
type WaterfallConnectorLine struct {

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

	// Width
	// arrayOK: false
	// type: number
	// Sets the line width (in px).
	Width float64 `json:"width,omitempty"`
}

// WaterfallConnector
type WaterfallConnector struct {

	// Line
	// role: Object
	Line *WaterfallConnectorLine `json:"line,omitempty"`

	// Mode
	// default: between
	// type: enumerated
	// Sets the shape of connector lines.
	Mode WaterfallConnectorMode `json:"mode,omitempty"`

	// Visible
	// arrayOK: false
	// type: boolean
	// Determines if connector lines are drawn.
	Visible Bool `json:"visible,omitempty"`
}

// WaterfallDecreasingMarkerLine
type WaterfallDecreasingMarkerLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the line color of all decreasing values.
	Color Color `json:"color,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the line width of all decreasing values.
	Width float64 `json:"width,omitempty"`
}

// WaterfallDecreasingMarker
type WaterfallDecreasingMarker struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the marker color of all decreasing values.
	Color Color `json:"color,omitempty"`

	// Line
	// role: Object
	Line *WaterfallDecreasingMarkerLine `json:"line,omitempty"`
}

// WaterfallDecreasing
type WaterfallDecreasing struct {

	// Marker
	// role: Object
	Marker *WaterfallDecreasingMarker `json:"marker,omitempty"`
}

// WaterfallHoverlabelFont Sets the font used in hover labels.
type WaterfallHoverlabelFont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color interface{} `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	Colorsrc String `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family interface{} `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `family`.
	Familysrc String `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size interface{} `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `size`.
	Sizesrc String `json:"sizesrc,omitempty"`
}

// WaterfallHoverlabel
type WaterfallHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align interface{} `json:"align,omitempty"`

	// Alignsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `align`.
	Alignsrc String `json:"alignsrc,omitempty"`

	// Bgcolor
	// arrayOK: true
	// type: color
	// Sets the background color of the hover labels for this trace
	Bgcolor interface{} `json:"bgcolor,omitempty"`

	// Bgcolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `bgcolor`.
	Bgcolorsrc String `json:"bgcolorsrc,omitempty"`

	// Bordercolor
	// arrayOK: true
	// type: color
	// Sets the border color of the hover labels for this trace.
	Bordercolor interface{} `json:"bordercolor,omitempty"`

	// Bordercolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `bordercolor`.
	Bordercolorsrc String `json:"bordercolorsrc,omitempty"`

	// Font
	// role: Object
	Font *WaterfallHoverlabelFont `json:"font,omitempty"`

	// Namelength
	// arrayOK: true
	// type: integer
	// Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis.
	Namelength interface{} `json:"namelength,omitempty"`

	// Namelengthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `namelength`.
	Namelengthsrc String `json:"namelengthsrc,omitempty"`
}

// WaterfallIncreasingMarkerLine
type WaterfallIncreasingMarkerLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the line color of all increasing values.
	Color Color `json:"color,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the line width of all increasing values.
	Width float64 `json:"width,omitempty"`
}

// WaterfallIncreasingMarker
type WaterfallIncreasingMarker struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the marker color of all increasing values.
	Color Color `json:"color,omitempty"`

	// Line
	// role: Object
	Line *WaterfallIncreasingMarkerLine `json:"line,omitempty"`
}

// WaterfallIncreasing
type WaterfallIncreasing struct {

	// Marker
	// role: Object
	Marker *WaterfallIncreasingMarker `json:"marker,omitempty"`
}

// WaterfallInsidetextfont Sets the font used for `text` lying inside the bar.
type WaterfallInsidetextfont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color interface{} `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	Colorsrc String `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family interface{} `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `family`.
	Familysrc String `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size interface{} `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `size`.
	Sizesrc String `json:"sizesrc,omitempty"`
}

// WaterfallLegendgrouptitleFont Sets this legend group's title font.
type WaterfallLegendgrouptitleFont struct {

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

// WaterfallLegendgrouptitle
type WaterfallLegendgrouptitle struct {

	// Font
	// role: Object
	Font *WaterfallLegendgrouptitleFont `json:"font,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the legend group.
	Text String `json:"text,omitempty"`
}

// WaterfallOutsidetextfont Sets the font used for `text` lying outside the bar.
type WaterfallOutsidetextfont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color interface{} `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	Colorsrc String `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family interface{} `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `family`.
	Familysrc String `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size interface{} `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `size`.
	Sizesrc String `json:"sizesrc,omitempty"`
}

// WaterfallStream
type WaterfallStream struct {

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

// WaterfallTextfont Sets the font used for `text`.
type WaterfallTextfont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color interface{} `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	Colorsrc String `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family interface{} `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `family`.
	Familysrc String `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size interface{} `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `size`.
	Sizesrc String `json:"sizesrc,omitempty"`
}

// WaterfallTotalsMarkerLine
type WaterfallTotalsMarkerLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the line color of all intermediate sums and total values.
	Color Color `json:"color,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the line width of all intermediate sums and total values.
	Width float64 `json:"width,omitempty"`
}

// WaterfallTotalsMarker
type WaterfallTotalsMarker struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the marker color of all intermediate sums and total values.
	Color Color `json:"color,omitempty"`

	// Line
	// role: Object
	Line *WaterfallTotalsMarkerLine `json:"line,omitempty"`
}

// WaterfallTotals
type WaterfallTotals struct {

	// Marker
	// role: Object
	Marker *WaterfallTotalsMarker `json:"marker,omitempty"`
}

// WaterfallConnectorMode Sets the shape of connector lines.
type WaterfallConnectorMode string

const (
	WaterfallConnectorModeSpanning WaterfallConnectorMode = "spanning"
	WaterfallConnectorModeBetween  WaterfallConnectorMode = "between"
)

// WaterfallConstraintext Constrain the size of text inside or outside a bar to be no larger than the bar itself.
type WaterfallConstraintext string

const (
	WaterfallConstraintextInside  WaterfallConstraintext = "inside"
	WaterfallConstraintextOutside WaterfallConstraintext = "outside"
	WaterfallConstraintextBoth    WaterfallConstraintext = "both"
	WaterfallConstraintextNone    WaterfallConstraintext = "none"
)

// WaterfallHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type WaterfallHoverlabelAlign string

const (
	WaterfallHoverlabelAlignLeft  WaterfallHoverlabelAlign = "left"
	WaterfallHoverlabelAlignRight WaterfallHoverlabelAlign = "right"
	WaterfallHoverlabelAlignAuto  WaterfallHoverlabelAlign = "auto"
)

// WaterfallInsidetextanchor Determines if texts are kept at center or start/end points in `textposition` *inside* mode.
type WaterfallInsidetextanchor string

const (
	WaterfallInsidetextanchorEnd    WaterfallInsidetextanchor = "end"
	WaterfallInsidetextanchorMiddle WaterfallInsidetextanchor = "middle"
	WaterfallInsidetextanchorStart  WaterfallInsidetextanchor = "start"
)

// WaterfallOrientation Sets the orientation of the bars. With *v* (*h*), the value of the each bar spans along the vertical (horizontal).
type WaterfallOrientation string

const (
	WaterfallOrientationV WaterfallOrientation = "v"
	WaterfallOrientationH WaterfallOrientation = "h"
)

// WaterfallTextposition Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside. If *none*, no text appears.
type WaterfallTextposition string

const (
	WaterfallTextpositionInside  WaterfallTextposition = "inside"
	WaterfallTextpositionOutside WaterfallTextposition = "outside"
	WaterfallTextpositionAuto    WaterfallTextposition = "auto"
	WaterfallTextpositionNone    WaterfallTextposition = "none"
)

// WaterfallVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type WaterfallVisible interface{}

var (
	WaterfallVisibleTrue       WaterfallVisible = true
	WaterfallVisibleFalse      WaterfallVisible = false
	WaterfallVisibleLegendonly WaterfallVisible = "legendonly"
)

// WaterfallXperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
type WaterfallXperiodalignment string

const (
	WaterfallXperiodalignmentStart  WaterfallXperiodalignment = "start"
	WaterfallXperiodalignmentMiddle WaterfallXperiodalignment = "middle"
	WaterfallXperiodalignmentEnd    WaterfallXperiodalignment = "end"
)

// WaterfallYperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
type WaterfallYperiodalignment string

const (
	WaterfallYperiodalignmentStart  WaterfallYperiodalignment = "start"
	WaterfallYperiodalignmentMiddle WaterfallYperiodalignment = "middle"
	WaterfallYperiodalignmentEnd    WaterfallYperiodalignment = "end"
)

// WaterfallHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type WaterfallHoverinfo string

const (
	// Flags
	WaterfallHoverinfoName    WaterfallHoverinfo = "name"
	WaterfallHoverinfoX       WaterfallHoverinfo = "x"
	WaterfallHoverinfoY       WaterfallHoverinfo = "y"
	WaterfallHoverinfoText    WaterfallHoverinfo = "text"
	WaterfallHoverinfoInitial WaterfallHoverinfo = "initial"
	WaterfallHoverinfoDelta   WaterfallHoverinfo = "delta"
	WaterfallHoverinfoFinal   WaterfallHoverinfo = "final"

	// Extra
	WaterfallHoverinfoAll  WaterfallHoverinfo = "all"
	WaterfallHoverinfoNone WaterfallHoverinfo = "none"
	WaterfallHoverinfoSkip WaterfallHoverinfo = "skip"
)

// WaterfallTextinfo Determines which trace information appear on the graph. In the case of having multiple waterfalls, totals are computed separately (per trace).
type WaterfallTextinfo string

const (
	// Flags
	WaterfallTextinfoLabel   WaterfallTextinfo = "label"
	WaterfallTextinfoText    WaterfallTextinfo = "text"
	WaterfallTextinfoInitial WaterfallTextinfo = "initial"
	WaterfallTextinfoDelta   WaterfallTextinfo = "delta"
	WaterfallTextinfoFinal   WaterfallTextinfo = "final"

	// Extra
	WaterfallTextinfoNone WaterfallTextinfo = "none"
)
