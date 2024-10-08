package grob

// Code generated by go-plotly/generator. DO NOT EDIT.

import (
	"encoding/json"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

var TraceTypePointcloud types.TraceType = "pointcloud"

func (t *Pointcloud) GetType() types.TraceType {
	return TraceTypePointcloud
}

func (t *Pointcloud) MarshalJSON() ([]byte, error) {
	// Define the custom JSON structure including the "type" field
	type Alias Pointcloud
	return json.Marshal(&struct {
		Type types.TraceType `json:"type"`
		*Alias
	}{
		Type:  t.GetType(), // Add your desired default value here
		Alias: (*Alias)(t), // Embed the original struct fields
	})
}

// Pointcloud *pointcloud* trace is deprecated! Please consider switching to the *scattergl* trace type. The data visualized as a point cloud set in `x` and `y` using the WebGl plotting engine.
type Pointcloud struct {

	// Customdata
	// arrayOK: false
	// type: data_array
	// Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.pointcloud.attributes.customdata
	Customdata *types.DataArrayType `json:"customdata,omitempty"`

	// Customdatasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `customdata`.
	// .schema.traces.pointcloud.attributes.customdatasrc
	Customdatasrc types.StringType `json:"customdatasrc,omitempty"`

	// Hoverinfo
	// arrayOK: true
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	// .schema.traces.pointcloud.attributes.hoverinfo
	Hoverinfo *types.ArrayOK[*PointcloudHoverinfo] `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hoverinfo`.
	// .schema.traces.pointcloud.attributes.hoverinfosrc
	Hoverinfosrc types.StringType `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// arrayOK: false
	// role: Object
	// .schema.traces.pointcloud.attributes.hoverlabel
	Hoverlabel *PointcloudHoverlabel `json:"hoverlabel,omitempty"`

	// Ids
	// arrayOK: false
	// type: data_array
	// Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.pointcloud.attributes.ids
	Ids *types.DataArrayType `json:"ids,omitempty"`

	// Idssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `ids`.
	// .schema.traces.pointcloud.attributes.idssrc
	Idssrc types.StringType `json:"idssrc,omitempty"`

	// Indices
	// arrayOK: false
	// type: data_array
	// A sequential value, 0..n, supply it to avoid creating this array inside plotting. If specified, it must be a typed `Int32Array` array. Its length must be equal to or greater than the number of points. For the best performance and memory use, create one large `indices` typed array that is guaranteed to be at least as long as the largest number of points during use, and reuse it on each `Plotly.restyle()` call.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.pointcloud.attributes.indices
	Indices *types.DataArrayType `json:"indices,omitempty"`

	// Indicessrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `indices`.
	// .schema.traces.pointcloud.attributes.indicessrc
	Indicessrc types.StringType `json:"indicessrc,omitempty"`

	// Legendgroup
	// arrayOK: false
	// type: string
	// Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	// .schema.traces.pointcloud.attributes.legendgroup
	Legendgroup types.StringType `json:"legendgroup,omitempty"`

	// Legendgrouptitle
	// arrayOK: false
	// role: Object
	// .schema.traces.pointcloud.attributes.legendgrouptitle
	Legendgrouptitle *PointcloudLegendgrouptitle `json:"legendgrouptitle,omitempty"`

	// Legendrank
	// arrayOK: false
	// type: number
	// Sets the legend rank for this trace. Items and groups with smaller ranks are presented on top/left side while with `*reversed* `legend.traceorder` they are on bottom/right side. The default legendrank is 1000, so that you can use ranks less than 1000 to place certain items before all unranked items, and ranks greater than 1000 to go after all unranked items.
	// .schema.traces.pointcloud.attributes.legendrank
	Legendrank types.NumberType `json:"legendrank,omitempty"`

	// Legendwidth
	// arrayOK: false
	// type: number
	// Sets the width (in px or fraction) of the legend for this trace.
	// .schema.traces.pointcloud.attributes.legendwidth
	Legendwidth types.NumberType `json:"legendwidth,omitempty"`

	// Marker
	// arrayOK: false
	// role: Object
	// .schema.traces.pointcloud.attributes.marker
	Marker *PointcloudMarker `json:"marker,omitempty"`

	// Meta
	// arrayOK: true
	// type: any
	// Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	// .schema.traces.pointcloud.attributes.meta
	Meta *types.ArrayOK[*interface{}] `json:"meta,omitempty"`

	// Metasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `meta`.
	// .schema.traces.pointcloud.attributes.metasrc
	Metasrc types.StringType `json:"metasrc,omitempty"`

	// Name
	// arrayOK: false
	// type: string
	// Sets the trace name. The trace name appear as the legend item and on hover.
	// .schema.traces.pointcloud.attributes.name
	Name types.StringType `json:"name,omitempty"`

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the opacity of the trace.
	// .schema.traces.pointcloud.attributes.opacity
	Opacity types.NumberType `json:"opacity,omitempty"`

	// Showlegend
	// arrayOK: false
	// type: boolean
	// Determines whether or not an item corresponding to this trace is shown in the legend.
	// .schema.traces.pointcloud.attributes.showlegend
	Showlegend types.BoolType `json:"showlegend,omitempty"`

	// Stream
	// arrayOK: false
	// role: Object
	// .schema.traces.pointcloud.attributes.stream
	Stream *PointcloudStream `json:"stream,omitempty"`

	// Text
	// arrayOK: true
	// type: string
	// Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	// .schema.traces.pointcloud.attributes.text
	Text *types.ArrayOK[*types.StringType] `json:"text,omitempty"`

	// Textsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `text`.
	// .schema.traces.pointcloud.attributes.textsrc
	Textsrc types.StringType `json:"textsrc,omitempty"`

	// Uid
	// arrayOK: false
	// type: string
	// Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	// .schema.traces.pointcloud.attributes.uid
	Uid types.StringType `json:"uid,omitempty"`

	// Uirevision
	// arrayOK: false
	// type: any
	// Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	// .schema.traces.pointcloud.attributes.uirevision
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible
	// arrayOK: false
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	// .schema.traces.pointcloud.attributes.visible
	Visible PointcloudVisible `json:"visible,omitempty"`

	// X
	// arrayOK: false
	// type: data_array
	// Sets the x coordinates.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.pointcloud.attributes.x
	X *types.DataArrayType `json:"x,omitempty"`

	// Xaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	// .schema.traces.pointcloud.attributes.xaxis
	Xaxis types.StringType `json:"xaxis,omitempty"`

	// Xbounds
	// arrayOK: false
	// type: data_array
	// Specify `xbounds` in the shape of `[xMin, xMax] to avoid looping through the `xy` typed array. Use it in conjunction with `xy` and `ybounds` for the performance benefits.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.pointcloud.attributes.xbounds
	Xbounds *types.DataArrayType `json:"xbounds,omitempty"`

	// Xboundssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `xbounds`.
	// .schema.traces.pointcloud.attributes.xboundssrc
	Xboundssrc types.StringType `json:"xboundssrc,omitempty"`

	// Xsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `x`.
	// .schema.traces.pointcloud.attributes.xsrc
	Xsrc types.StringType `json:"xsrc,omitempty"`

	// Xy
	// arrayOK: false
	// type: data_array
	// Faster alternative to specifying `x` and `y` separately. If supplied, it must be a typed `Float32Array` array that represents points such that `xy[i * 2] = x[i]` and `xy[i * 2 + 1] = y[i]`
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.pointcloud.attributes.xy
	Xy *types.DataArrayType `json:"xy,omitempty"`

	// Xysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `xy`.
	// .schema.traces.pointcloud.attributes.xysrc
	Xysrc types.StringType `json:"xysrc,omitempty"`

	// Y
	// arrayOK: false
	// type: data_array
	// Sets the y coordinates.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.pointcloud.attributes.y
	Y *types.DataArrayType `json:"y,omitempty"`

	// Yaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	// .schema.traces.pointcloud.attributes.yaxis
	Yaxis types.StringType `json:"yaxis,omitempty"`

	// Ybounds
	// arrayOK: false
	// type: data_array
	// Specify `ybounds` in the shape of `[yMin, yMax] to avoid looping through the `xy` typed array. Use it in conjunction with `xy` and `xbounds` for the performance benefits.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.pointcloud.attributes.ybounds
	Ybounds *types.DataArrayType `json:"ybounds,omitempty"`

	// Yboundssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `ybounds`.
	// .schema.traces.pointcloud.attributes.yboundssrc
	Yboundssrc types.StringType `json:"yboundssrc,omitempty"`

	// Ysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `y`.
	// .schema.traces.pointcloud.attributes.ysrc
	Ysrc types.StringType `json:"ysrc,omitempty"`
}

// PointcloudHoverlabelFont Sets the font used in hover labels.
type PointcloudHoverlabelFont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	// .schema.traces.pointcloud.attributes.hoverlabel.font.color
	Color *types.ArrayOK[*types.Color] `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	// .schema.traces.pointcloud.attributes.hoverlabel.font.colorsrc
	Colorsrc types.StringType `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	// .schema.traces.pointcloud.attributes.hoverlabel.font.family
	Family *types.ArrayOK[*types.StringType] `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `family`.
	// .schema.traces.pointcloud.attributes.hoverlabel.font.familysrc
	Familysrc types.StringType `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	// .schema.traces.pointcloud.attributes.hoverlabel.font.size
	Size *types.ArrayOK[*types.NumberType] `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `size`.
	// .schema.traces.pointcloud.attributes.hoverlabel.font.sizesrc
	Sizesrc types.StringType `json:"sizesrc,omitempty"`
}

// PointcloudHoverlabel
type PointcloudHoverlabel struct {

	// Align
	// arrayOK: true
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	// .schema.traces.pointcloud.attributes.hoverlabel.align
	Align *types.ArrayOK[*PointcloudHoverlabelAlign] `json:"align,omitempty"`

	// Alignsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `align`.
	// .schema.traces.pointcloud.attributes.hoverlabel.alignsrc
	Alignsrc types.StringType `json:"alignsrc,omitempty"`

	// Bgcolor
	// arrayOK: true
	// type: color
	// Sets the background color of the hover labels for this trace
	// .schema.traces.pointcloud.attributes.hoverlabel.bgcolor
	Bgcolor *types.ArrayOK[*types.Color] `json:"bgcolor,omitempty"`

	// Bgcolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `bgcolor`.
	// .schema.traces.pointcloud.attributes.hoverlabel.bgcolorsrc
	Bgcolorsrc types.StringType `json:"bgcolorsrc,omitempty"`

	// Bordercolor
	// arrayOK: true
	// type: color
	// Sets the border color of the hover labels for this trace.
	// .schema.traces.pointcloud.attributes.hoverlabel.bordercolor
	Bordercolor *types.ArrayOK[*types.Color] `json:"bordercolor,omitempty"`

	// Bordercolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `bordercolor`.
	// .schema.traces.pointcloud.attributes.hoverlabel.bordercolorsrc
	Bordercolorsrc types.StringType `json:"bordercolorsrc,omitempty"`

	// Font
	// arrayOK: false
	// role: Object
	// .schema.traces.pointcloud.attributes.hoverlabel.font
	Font *PointcloudHoverlabelFont `json:"font,omitempty"`

	// Namelength
	// arrayOK: true
	// type: integer
	// Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis.
	// .schema.traces.pointcloud.attributes.hoverlabel.namelength
	Namelength *types.ArrayOK[*types.IntegerType] `json:"namelength,omitempty"`

	// Namelengthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `namelength`.
	// .schema.traces.pointcloud.attributes.hoverlabel.namelengthsrc
	Namelengthsrc types.StringType `json:"namelengthsrc,omitempty"`
}

// PointcloudLegendgrouptitleFont Sets this legend group's title font.
type PointcloudLegendgrouptitleFont struct {

	// Color
	// arrayOK: false
	// type: color
	//
	// .schema.traces.pointcloud.attributes.legendgrouptitle.font.color
	Color types.Color `json:"color,omitempty"`

	// Family
	// arrayOK: false
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	// .schema.traces.pointcloud.attributes.legendgrouptitle.font.family
	Family types.StringType `json:"family,omitempty"`

	// Size
	// arrayOK: false
	// type: number
	//
	// .schema.traces.pointcloud.attributes.legendgrouptitle.font.size
	Size types.NumberType `json:"size,omitempty"`
}

// PointcloudLegendgrouptitle
type PointcloudLegendgrouptitle struct {

	// Font
	// arrayOK: false
	// role: Object
	// .schema.traces.pointcloud.attributes.legendgrouptitle.font
	Font *PointcloudLegendgrouptitleFont `json:"font,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the legend group.
	// .schema.traces.pointcloud.attributes.legendgrouptitle.text
	Text types.StringType `json:"text,omitempty"`
}

// PointcloudMarkerBorder
type PointcloudMarkerBorder struct {

	// Arearatio
	// arrayOK: false
	// type: number
	// Specifies what fraction of the marker area is covered with the border.
	// .schema.traces.pointcloud.attributes.marker.border.arearatio
	Arearatio types.NumberType `json:"arearatio,omitempty"`

	// Color
	// arrayOK: false
	// type: color
	// Sets the stroke color. It accepts a specific color. If the color is not fully opaque and there are hundreds of thousands of points, it may cause slower zooming and panning.
	// .schema.traces.pointcloud.attributes.marker.border.color
	Color types.Color `json:"color,omitempty"`
}

// PointcloudMarker
type PointcloudMarker struct {

	// Blend
	// arrayOK: false
	// type: boolean
	// Determines if colors are blended together for a translucency effect in case `opacity` is specified as a value less then `1`. Setting `blend` to `true` reduces zoom/pan speed if used with large numbers of points.
	// .schema.traces.pointcloud.attributes.marker.blend
	Blend types.BoolType `json:"blend,omitempty"`

	// Border
	// arrayOK: false
	// role: Object
	// .schema.traces.pointcloud.attributes.marker.border
	Border *PointcloudMarkerBorder `json:"border,omitempty"`

	// Color
	// arrayOK: false
	// type: color
	// Sets the marker fill color. It accepts a specific color. If the color is not fully opaque and there are hundreds of thousands of points, it may cause slower zooming and panning.
	// .schema.traces.pointcloud.attributes.marker.color
	Color types.Color `json:"color,omitempty"`

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the marker opacity. The default value is `1` (fully opaque). If the markers are not fully opaque and there are hundreds of thousands of points, it may cause slower zooming and panning. Opacity fades the color even if `blend` is left on `false` even if there is no translucency effect in that case.
	// .schema.traces.pointcloud.attributes.marker.opacity
	Opacity types.NumberType `json:"opacity,omitempty"`

	// Sizemax
	// arrayOK: false
	// type: number
	// Sets the maximum size (in px) of the rendered marker points. Effective when the `pointcloud` shows only few points.
	// .schema.traces.pointcloud.attributes.marker.sizemax
	Sizemax types.NumberType `json:"sizemax,omitempty"`

	// Sizemin
	// arrayOK: false
	// type: number
	// Sets the minimum size (in px) of the rendered marker points, effective when the `pointcloud` shows a million or more points.
	// .schema.traces.pointcloud.attributes.marker.sizemin
	Sizemin types.NumberType `json:"sizemin,omitempty"`
}

// PointcloudStream
type PointcloudStream struct {

	// Maxpoints
	// arrayOK: false
	// type: number
	// Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot.
	// .schema.traces.pointcloud.attributes.stream.maxpoints
	Maxpoints types.NumberType `json:"maxpoints,omitempty"`

	// Token
	// arrayOK: false
	// type: string
	// The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details.
	// .schema.traces.pointcloud.attributes.stream.token
	Token types.StringType `json:"token,omitempty"`
}

// PointcloudHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
// .schema.traces.pointcloud.attributes.hoverlabel.align
type PointcloudHoverlabelAlign string

const (
	PointcloudHoverlabelAlignLeft  PointcloudHoverlabelAlign = "left"
	PointcloudHoverlabelAlignRight PointcloudHoverlabelAlign = "right"
	PointcloudHoverlabelAlignAuto  PointcloudHoverlabelAlign = "auto"
)

// PointcloudVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
// .schema.traces.pointcloud.attributes.visible
type PointcloudVisible interface{}

var (
	PointcloudVisibleTrue       PointcloudVisible = true
	PointcloudVisibleFalse      PointcloudVisible = false
	PointcloudVisibleLegendonly PointcloudVisible = "legendonly"
)

// PointcloudHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
// .schema.traces.pointcloud.attributes.hoverinfo
type PointcloudHoverinfo string

const (
	// Flags
	PointcloudHoverinfoX    PointcloudHoverinfo = "x"
	PointcloudHoverinfoY    PointcloudHoverinfo = "y"
	PointcloudHoverinfoZ    PointcloudHoverinfo = "z"
	PointcloudHoverinfoText PointcloudHoverinfo = "text"
	PointcloudHoverinfoName PointcloudHoverinfo = "name"

	// Extra
	PointcloudHoverinfoAll  PointcloudHoverinfo = "all"
	PointcloudHoverinfoNone PointcloudHoverinfo = "none"
	PointcloudHoverinfoSkip PointcloudHoverinfo = "skip"
)
