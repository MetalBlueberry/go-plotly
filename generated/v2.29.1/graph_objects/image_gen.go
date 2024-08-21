package grob

// Code generated by go-plotly/generator. DO NOT EDIT.

import (
	"encoding/json"
	"github.com/MetalBlueberry/go-plotly/pkg/types/arrayok"
	"github.com/MetalBlueberry/go-plotly/pkg/types/color"
)

var TraceTypeImage TraceType = "image"

func (trace *Image) GetType() TraceType {
	return TraceTypeImage
}

func (trace *Image) MarshalJSON() ([]byte, error) {
	// Define the custom JSON structure including the "type" field
	type Alias Image
	return json.Marshal(&struct {
		Type TraceType `json:"type"`
		*Alias
	}{
		Type:  trace.GetType(), // Add your desired default value here
		Alias: (*Alias)(trace), // Embed the original struct fields
	})
}

// Image Display an image, i.e. data on a 2D regular raster. By default, when an image is displayed in a subplot, its y axis will be reversed (ie. `autorange: 'reversed'`), constrained to the domain (ie. `constrain: 'domain'`) and it will have the same scale as its x axis (ie. `scaleanchor: 'x,`) in order for pixels to be rendered as squares.
type Image struct {

	// Colormodel
	// arrayOK: false
	// default: %!s(<nil>)
	// type: enumerated
	// Color model used to map the numerical color components described in `z` into colors. If `source` is specified, this attribute will be set to `rgba256` otherwise it defaults to `rgb`.
	Colormodel ImageColormodel `json:"colormodel,omitempty"`

	// Customdata
	// arrayOK: false
	// type: data_array
	// Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `customdata`.
	Customdatasrc string `json:"customdatasrc,omitempty"`

	// Dx
	// arrayOK: false
	// type: number
	// Set the pixel's horizontal size.
	Dx float64 `json:"dx,omitempty"`

	// Dy
	// arrayOK: false
	// type: number
	// Set the pixel's vertical size
	Dy float64 `json:"dy,omitempty"`

	// Hoverinfo
	// arrayOK: true
	// default: x+y+z+text+name
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo arrayok.Type[*ImageHoverinfo] `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hoverinfo`.
	Hoverinfosrc string `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// arrayOK: false
	// role: Object
	Hoverlabel *ImageHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate
	// arrayOK: true
	// type: string
	// Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}" as well as %{xother}, {%_xother}, {%_xother_}, {%xother_}. When showing info for several points, *xother* will be added to those with different x positions from the first point. An underscore before or after *(x|y)other* will add a space on that side, only when this field is shown. Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-format/tree/v1.4.5#d3-format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format/tree/v2.2.3#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. Finally, the template string has access to variables `z`, `color` and `colormodel`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate arrayok.Type[*string] `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hovertemplate`.
	Hovertemplatesrc string `json:"hovertemplatesrc,omitempty"`

	// Hovertext
	// arrayOK: false
	// type: data_array
	// Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hovertext`.
	Hovertextsrc string `json:"hovertextsrc,omitempty"`

	// Ids
	// arrayOK: false
	// type: data_array
	// Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `ids`.
	Idssrc string `json:"idssrc,omitempty"`

	// Legend
	// arrayOK: false
	// type: subplotid
	// Sets the reference to a legend to show this trace in. References to these legends are *legend*, *legend2*, *legend3*, etc. Settings for these legends are set in the layout, under `layout.legend`, `layout.legend2`, etc.
	Legend String `json:"legend,omitempty"`

	// Legendgrouptitle
	// arrayOK: false
	// role: Object
	Legendgrouptitle *ImageLegendgrouptitle `json:"legendgrouptitle,omitempty"`

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

	// Meta
	// arrayOK: true
	// type: any
	// Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta arrayok.Type[*interface{}] `json:"meta,omitempty"`

	// Metasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `meta`.
	Metasrc string `json:"metasrc,omitempty"`

	// Name
	// arrayOK: false
	// type: string
	// Sets the trace name. The trace name appears as the legend item and on hover.
	Name string `json:"name,omitempty"`

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Source
	// arrayOK: false
	// type: string
	// Specifies the data URI of the image to be visualized. The URI consists of "data:image/[<media subtype>][;base64],<data>"
	Source string `json:"source,omitempty"`

	// Stream
	// arrayOK: false
	// role: Object
	Stream *ImageStream `json:"stream,omitempty"`

	// Text
	// arrayOK: false
	// type: data_array
	// Sets the text elements associated with each z value.
	Text interface{} `json:"text,omitempty"`

	// Textsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `text`.
	Textsrc string `json:"textsrc,omitempty"`

	// Uid
	// arrayOK: false
	// type: string
	// Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid string `json:"uid,omitempty"`

	// Uirevision
	// arrayOK: false
	// type: any
	// Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible
	// arrayOK: false
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ImageVisible `json:"visible,omitempty"`

	// X0
	// arrayOK: false
	// type: any
	// Set the image's x position. The left edge of the image (or the right edge if the x axis is reversed or dx is negative) will be found at xmin=x0-dx/2
	X0 interface{} `json:"x0,omitempty"`

	// Xaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Y0
	// arrayOK: false
	// type: any
	// Set the image's y position. The top edge of the image (or the bottom edge if the y axis is NOT reversed or if dy is negative) will be found at ymin=y0-dy/2. By default when an image trace is included, the y axis will be reversed so that the image is right-side-up, but you can disable this by setting yaxis.autorange=true or by providing an explicit y axis range.
	Y0 interface{} `json:"y0,omitempty"`

	// Yaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Z
	// arrayOK: false
	// type: data_array
	// A 2-dimensional array in which each element is an array of 3 or 4 numbers representing a color.
	Z interface{} `json:"z,omitempty"`

	// Zmax
	// arrayOK: false
	// type: info_array
	// Array defining the higher bound for each color component. Note that the default value will depend on the colormodel. For the `rgb` colormodel, it is [255, 255, 255]. For the `rgba` colormodel, it is [255, 255, 255, 1]. For the `rgba256` colormodel, it is [255, 255, 255, 255]. For the `hsl` colormodel, it is [360, 100, 100]. For the `hsla` colormodel, it is [360, 100, 100, 1].
	Zmax interface{} `json:"zmax,omitempty"`

	// Zmin
	// arrayOK: false
	// type: info_array
	// Array defining the lower bound for each color component. Note that the default value will depend on the colormodel. For the `rgb` colormodel, it is [0, 0, 0]. For the `rgba` colormodel, it is [0, 0, 0, 0]. For the `rgba256` colormodel, it is [0, 0, 0, 0]. For the `hsl` colormodel, it is [0, 0, 0]. For the `hsla` colormodel, it is [0, 0, 0, 0].
	Zmin interface{} `json:"zmin,omitempty"`

	// Zsmooth
	// arrayOK: false
	// default: %!s(bool=false)
	// type: enumerated
	// Picks a smoothing algorithm used to smooth `z` data. This only applies for image traces that use the `source` attribute.
	Zsmooth ImageZsmooth `json:"zsmooth,omitempty"`

	// Zsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `z`.
	Zsrc string `json:"zsrc,omitempty"`
}

// ImageHoverlabelFont Sets the font used in hover labels.
type ImageHoverlabelFont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	Color arrayok.Type[*color.String] `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	Colorsrc string `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family arrayok.Type[*string] `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `family`.
	Familysrc string `json:"familysrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	Size arrayok.Type[*float64] `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `size`.
	Sizesrc string `json:"sizesrc,omitempty"`
}

// ImageHoverlabel
type ImageHoverlabel struct {

	// Align
	// arrayOK: true
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align arrayok.Type[*ImageHoverlabelAlign] `json:"align,omitempty"`

	// Alignsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `align`.
	Alignsrc string `json:"alignsrc,omitempty"`

	// Bgcolor
	// arrayOK: true
	// type: color
	// Sets the background color of the hover labels for this trace
	Bgcolor arrayok.Type[*color.String] `json:"bgcolor,omitempty"`

	// Bgcolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `bgcolor`.
	Bgcolorsrc string `json:"bgcolorsrc,omitempty"`

	// Bordercolor
	// arrayOK: true
	// type: color
	// Sets the border color of the hover labels for this trace.
	Bordercolor arrayok.Type[*color.String] `json:"bordercolor,omitempty"`

	// Bordercolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `bordercolor`.
	Bordercolorsrc string `json:"bordercolorsrc,omitempty"`

	// Font
	// arrayOK: false
	// role: Object
	Font *ImageHoverlabelFont `json:"font,omitempty"`

	// Namelength
	// arrayOK: true
	// type: integer
	// Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis.
	Namelength arrayok.Type[*int64] `json:"namelength,omitempty"`

	// Namelengthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `namelength`.
	Namelengthsrc string `json:"namelengthsrc,omitempty"`
}

// ImageLegendgrouptitleFont Sets this legend group's title font.
type ImageLegendgrouptitleFont struct {

	// Color
	// arrayOK: false
	// type: color
	//
	Color color.String `json:"color,omitempty"`

	// Family
	// arrayOK: false
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	Family string `json:"family,omitempty"`

	// Size
	// arrayOK: false
	// type: number
	//
	Size float64 `json:"size,omitempty"`
}

// ImageLegendgrouptitle
type ImageLegendgrouptitle struct {

	// Font
	// arrayOK: false
	// role: Object
	Font *ImageLegendgrouptitleFont `json:"font,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the legend group.
	Text string `json:"text,omitempty"`
}

// ImageStream
type ImageStream struct {

	// Maxpoints
	// arrayOK: false
	// type: number
	// Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot.
	Maxpoints float64 `json:"maxpoints,omitempty"`

	// Token
	// arrayOK: false
	// type: string
	// The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details.
	Token string `json:"token,omitempty"`
}

// ImageColormodel Color model used to map the numerical color components described in `z` into colors. If `source` is specified, this attribute will be set to `rgba256` otherwise it defaults to `rgb`.
type ImageColormodel string

const (
	ImageColormodelRgb     ImageColormodel = "rgb"
	ImageColormodelRgba    ImageColormodel = "rgba"
	ImageColormodelRgba256 ImageColormodel = "rgba256"
	ImageColormodelHsl     ImageColormodel = "hsl"
	ImageColormodelHsla    ImageColormodel = "hsla"
)

// ImageHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ImageHoverlabelAlign string

const (
	ImageHoverlabelAlignLeft  ImageHoverlabelAlign = "left"
	ImageHoverlabelAlignRight ImageHoverlabelAlign = "right"
	ImageHoverlabelAlignAuto  ImageHoverlabelAlign = "auto"
)

// ImageVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ImageVisible interface{}

var (
	ImageVisibleTrue       ImageVisible = true
	ImageVisibleFalse      ImageVisible = false
	ImageVisibleLegendonly ImageVisible = "legendonly"
)

// ImageZsmooth Picks a smoothing algorithm used to smooth `z` data. This only applies for image traces that use the `source` attribute.
type ImageZsmooth interface{}

var (
	ImageZsmoothFast  ImageZsmooth = "fast"
	ImageZsmoothFalse ImageZsmooth = false
)

// ImageHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type ImageHoverinfo string

const (
	// Flags
	ImageHoverinfoX     ImageHoverinfo = "x"
	ImageHoverinfoY     ImageHoverinfo = "y"
	ImageHoverinfoZ     ImageHoverinfo = "z"
	ImageHoverinfoColor ImageHoverinfo = "color"
	ImageHoverinfoName  ImageHoverinfo = "name"
	ImageHoverinfoText  ImageHoverinfo = "text"

	// Extra
	ImageHoverinfoAll  ImageHoverinfo = "all"
	ImageHoverinfoNone ImageHoverinfo = "none"
	ImageHoverinfoSkip ImageHoverinfo = "skip"
)
