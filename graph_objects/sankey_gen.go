package grob

var TraceTypeSankey TraceType = "sankey"

func (trace *Sankey) GetType() TraceType {
	return TraceTypeSankey
}

// Sankey Sankey plots for network flow data analysis. The nodes are specified in `nodes` and the links between sources and targets in `links`. The colors are set in `nodes[i].color` and `links[i].color`, otherwise defaults are used.
type Sankey struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Arrangement
	// default: snap
	// type: enumerated
	// If value is `snap` (the default), the node arrangement is assisted by automatic snapping of elements to preserve space between nodes specified via `nodepad`. If value is `perpendicular`, the nodes can only move along a line perpendicular to the flow. If value is `freeform`, the nodes can freely move on the plane. If value is `fixed`, the nodes are stationary.
	Arrangement SankeyArrangement `json:"arrangement,omitempty"`

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
	Domain *SankeyDomain `json:"domain,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired. Note that this attribute is superseded by `node.hoverinfo` and `node.hoverinfo` for nodes and links respectively.
	Hoverinfo SankeyHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *SankeyHoverlabel `json:"hoverlabel,omitempty"`

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

	// Link
	// role: Object
	Link *SankeyLink `json:"link,omitempty"`

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

	// Node
	// role: Object
	Node *SankeyNode `json:"node,omitempty"`

	// Orientation
	// default: h
	// type: enumerated
	// Sets the orientation of the Sankey diagram.
	Orientation SankeyOrientation `json:"orientation,omitempty"`

	// Selectedpoints
	// arrayOK: false
	// type: any
	// Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Stream
	// role: Object
	Stream *SankeyStream `json:"stream,omitempty"`

	// Textfont
	// role: Object
	Textfont *SankeyTextfont `json:"textfont,omitempty"`

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

	// Valueformat
	// arrayOK: false
	// type: string
	// Sets the value formatting rule using d3 formatting mini-language which is similar to those of Python. See https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format
	Valueformat String `json:"valueformat,omitempty"`

	// Valuesuffix
	// arrayOK: false
	// type: string
	// Adds a unit to follow the value in the hover tooltip. Add a space if a separation is necessary from the value.
	Valuesuffix String `json:"valuesuffix,omitempty"`

	// Visible
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible SankeyVisible `json:"visible,omitempty"`
}

// SankeyDomain
type SankeyDomain struct {

	// Column
	// arrayOK: false
	// type: integer
	// If there is a layout grid, use the domain for this column in the grid for this sankey trace .
	Column int64 `json:"column,omitempty"`

	// Row
	// arrayOK: false
	// type: integer
	// If there is a layout grid, use the domain for this row in the grid for this sankey trace .
	Row int64 `json:"row,omitempty"`

	// X
	// arrayOK: false
	// type: info_array
	// Sets the horizontal domain of this sankey trace (in plot fraction).
	X interface{} `json:"x,omitempty"`

	// Y
	// arrayOK: false
	// type: info_array
	// Sets the vertical domain of this sankey trace (in plot fraction).
	Y interface{} `json:"y,omitempty"`
}

// SankeyHoverlabelFont Sets the font used in hover labels.
type SankeyHoverlabelFont struct {

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

// SankeyHoverlabel
type SankeyHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align SankeyHoverlabelAlign `json:"align,omitempty"`

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
	Font *SankeyHoverlabelFont `json:"font,omitempty"`

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

// SankeyLinkHoverlabelFont Sets the font used in hover labels.
type SankeyLinkHoverlabelFont struct {

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

// SankeyLinkHoverlabel
type SankeyLinkHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align SankeyLinkHoverlabelAlign `json:"align,omitempty"`

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
	Font *SankeyLinkHoverlabelFont `json:"font,omitempty"`

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

// SankeyLinkLine
type SankeyLinkLine struct {

	// Color
	// arrayOK: true
	// type: color
	// Sets the color of the `line` around each `link`.
	Color Color `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Width
	// arrayOK: true
	// type: number
	// Sets the width (in px) of the `line` around each `link`.
	Width float64 `json:"width,omitempty"`

	// Widthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  width .
	Widthsrc String `json:"widthsrc,omitempty"`
}

// SankeyLink The links of the Sankey plot.
type SankeyLink struct {

	// Color
	// arrayOK: true
	// type: color
	// Sets the `link` color. It can be a single value, or an array for specifying color for each `link`. If `link.color` is omitted, then by default, a translucent grey link will be used.
	Color Color `json:"color,omitempty"`

	// Colorscales
	// It's an items array and what goes inside it's... messy... check the docs
	// I will be happy if you want to contribute by implementing this
	// just raise an issue before you start so we do not overlap
	Colorscales interface{} `json:"colorscales,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Customdata
	// arrayOK: false
	// type: data_array
	// Assigns extra data to each link.
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Hoverinfo
	// default: all
	// type: enumerated
	// Determines which trace information appear when hovering links. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo SankeyLinkHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *SankeyLinkHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate
	// arrayOK: true
	// type: string
	// Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `value` and `label`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate String `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Label
	// arrayOK: false
	// type: data_array
	// The shown name of the link.
	Label interface{} `json:"label,omitempty"`

	// Labelsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  label .
	Labelsrc String `json:"labelsrc,omitempty"`

	// Line
	// role: Object
	Line *SankeyLinkLine `json:"line,omitempty"`

	// Source
	// arrayOK: false
	// type: data_array
	// An integer number `[0..nodes.length - 1]` that represents the source node.
	Source interface{} `json:"source,omitempty"`

	// Sourcesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  source .
	Sourcesrc String `json:"sourcesrc,omitempty"`

	// Target
	// arrayOK: false
	// type: data_array
	// An integer number `[0..nodes.length - 1]` that represents the target node.
	Target interface{} `json:"target,omitempty"`

	// Targetsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  target .
	Targetsrc String `json:"targetsrc,omitempty"`

	// Value
	// arrayOK: false
	// type: data_array
	// A numeric value representing the flow volume value.
	Value interface{} `json:"value,omitempty"`

	// Valuesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  value .
	Valuesrc String `json:"valuesrc,omitempty"`
}

// SankeyNodeHoverlabelFont Sets the font used in hover labels.
type SankeyNodeHoverlabelFont struct {

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

// SankeyNodeHoverlabel
type SankeyNodeHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align SankeyNodeHoverlabelAlign `json:"align,omitempty"`

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
	Font *SankeyNodeHoverlabelFont `json:"font,omitempty"`

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

// SankeyNodeLine
type SankeyNodeLine struct {

	// Color
	// arrayOK: true
	// type: color
	// Sets the color of the `line` around each `node`.
	Color Color `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Width
	// arrayOK: true
	// type: number
	// Sets the width (in px) of the `line` around each `node`.
	Width float64 `json:"width,omitempty"`

	// Widthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  width .
	Widthsrc String `json:"widthsrc,omitempty"`
}

// SankeyNode The nodes of the Sankey plot.
type SankeyNode struct {

	// Color
	// arrayOK: true
	// type: color
	// Sets the `node` color. It can be a single value, or an array for specifying color for each `node`. If `node.color` is omitted, then the default `Plotly` color palette will be cycled through to have a variety of colors. These defaults are not fully opaque, to allow some visibility of what is beneath the node.
	Color Color `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`

	// Customdata
	// arrayOK: false
	// type: data_array
	// Assigns extra data to each node.
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Groups
	// arrayOK: false
	// type: info_array
	// Groups of nodes. Each group is defined by an array with the indices of the nodes it contains. Multiple groups can be specified.
	Groups interface{} `json:"groups,omitempty"`

	// Hoverinfo
	// default: all
	// type: enumerated
	// Determines which trace information appear when hovering nodes. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo SankeyNodeHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *SankeyNodeHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate
	// arrayOK: true
	// type: string
	// Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `value` and `label`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate String `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Label
	// arrayOK: false
	// type: data_array
	// The shown name of the node.
	Label interface{} `json:"label,omitempty"`

	// Labelsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  label .
	Labelsrc String `json:"labelsrc,omitempty"`

	// Line
	// role: Object
	Line *SankeyNodeLine `json:"line,omitempty"`

	// Pad
	// arrayOK: false
	// type: number
	// Sets the padding (in px) between the `nodes`.
	Pad float64 `json:"pad,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness (in px) of the `nodes`.
	Thickness float64 `json:"thickness,omitempty"`

	// X
	// arrayOK: false
	// type: data_array
	// The normalized horizontal position of the node.
	X interface{} `json:"x,omitempty"`

	// Xsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y
	// arrayOK: false
	// type: data_array
	// The normalized vertical position of the node.
	Y interface{} `json:"y,omitempty"`

	// Ysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

// SankeyStream
type SankeyStream struct {

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

// SankeyTextfont Sets the font for node labels
type SankeyTextfont struct {

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

// SankeyArrangement If value is `snap` (the default), the node arrangement is assisted by automatic snapping of elements to preserve space between nodes specified via `nodepad`. If value is `perpendicular`, the nodes can only move along a line perpendicular to the flow. If value is `freeform`, the nodes can freely move on the plane. If value is `fixed`, the nodes are stationary.
type SankeyArrangement string

const (
	SankeyArrangementSnap          SankeyArrangement = "snap"
	SankeyArrangementPerpendicular SankeyArrangement = "perpendicular"
	SankeyArrangementFreeform      SankeyArrangement = "freeform"
	SankeyArrangementFixed         SankeyArrangement = "fixed"
)

// SankeyHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SankeyHoverlabelAlign string

const (
	SankeyHoverlabelAlignLeft  SankeyHoverlabelAlign = "left"
	SankeyHoverlabelAlignRight SankeyHoverlabelAlign = "right"
	SankeyHoverlabelAlignAuto  SankeyHoverlabelAlign = "auto"
)

// SankeyLinkHoverinfo Determines which trace information appear when hovering links. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type SankeyLinkHoverinfo string

const (
	SankeyLinkHoverinfoAll  SankeyLinkHoverinfo = "all"
	SankeyLinkHoverinfoNone SankeyLinkHoverinfo = "none"
	SankeyLinkHoverinfoSkip SankeyLinkHoverinfo = "skip"
)

// SankeyLinkHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SankeyLinkHoverlabelAlign string

const (
	SankeyLinkHoverlabelAlignLeft  SankeyLinkHoverlabelAlign = "left"
	SankeyLinkHoverlabelAlignRight SankeyLinkHoverlabelAlign = "right"
	SankeyLinkHoverlabelAlignAuto  SankeyLinkHoverlabelAlign = "auto"
)

// SankeyNodeHoverinfo Determines which trace information appear when hovering nodes. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type SankeyNodeHoverinfo string

const (
	SankeyNodeHoverinfoAll  SankeyNodeHoverinfo = "all"
	SankeyNodeHoverinfoNone SankeyNodeHoverinfo = "none"
	SankeyNodeHoverinfoSkip SankeyNodeHoverinfo = "skip"
)

// SankeyNodeHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SankeyNodeHoverlabelAlign string

const (
	SankeyNodeHoverlabelAlignLeft  SankeyNodeHoverlabelAlign = "left"
	SankeyNodeHoverlabelAlignRight SankeyNodeHoverlabelAlign = "right"
	SankeyNodeHoverlabelAlignAuto  SankeyNodeHoverlabelAlign = "auto"
)

// SankeyOrientation Sets the orientation of the Sankey diagram.
type SankeyOrientation string

const (
	SankeyOrientationV SankeyOrientation = "v"
	SankeyOrientationH SankeyOrientation = "h"
)

// SankeyVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type SankeyVisible interface{}

var (
	SankeyVisibleTrue       SankeyVisible = true
	SankeyVisibleFalse      SankeyVisible = false
	SankeyVisibleLegendonly SankeyVisible = "legendonly"
)

// SankeyHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired. Note that this attribute is superseded by `node.hoverinfo` and `node.hoverinfo` for nodes and links respectively.
type SankeyHoverinfo string

const (
	// Flags

	// Extra
	SankeyHoverinfoAll  SankeyHoverinfo = "all"
	SankeyHoverinfoNone SankeyHoverinfo = "none"
	SankeyHoverinfoSkip SankeyHoverinfo = "skip"
)
