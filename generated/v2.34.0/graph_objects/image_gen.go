package grob

// Code generated by go-plotly/generator. DO NOT EDIT.

import (
	"encoding/json"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

var TraceTypeImage types.TraceType = "image"

func (t *Image) GetType() types.TraceType {
	return TraceTypeImage
}

func (t *Image) MarshalJSON() ([]byte, error) {
	// Define the custom JSON structure including the "type" field
	type Alias Image
	return json.Marshal(&struct {
		Type types.TraceType `json:"type"`
		*Alias
	}{
		Type:  t.GetType(), // Add your desired default value here
		Alias: (*Alias)(t), // Embed the original struct fields
	})
}

// Image Display an image, i.e. data on a 2D regular raster. By default, when an image is displayed in a subplot, its y axis will be reversed (ie. `autorange: 'reversed'`), constrained to the domain (ie. `constrain: 'domain'`) and it will have the same scale as its x axis (ie. `scaleanchor: 'x,`) in order for pixels to be rendered as squares.
type Image struct {

	// Colormodel
	// arrayOK: false
	// default: %!s(<nil>)
	// type: enumerated
	// Color model used to map the numerical color components described in `z` into colors. If `source` is specified, this attribute will be set to `rgba256` otherwise it defaults to `rgb`.
	// .schema.traces.image.attributes.colormodel
	Colormodel ImageColormodel `json:"colormodel,omitempty"`

	// Customdata
	// arrayOK: false
	// type: data_array
	// Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.image.attributes.customdata
	Customdata *types.DataArrayType `json:"customdata,omitempty"`

	// Customdatasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `customdata`.
	// .schema.traces.image.attributes.customdatasrc
	Customdatasrc types.StringType `json:"customdatasrc,omitempty"`

	// Dx
	// arrayOK: false
	// type: number
	// Set the pixel's horizontal size.
	// .schema.traces.image.attributes.dx
	Dx types.NumberType `json:"dx,omitempty"`

	// Dy
	// arrayOK: false
	// type: number
	// Set the pixel's vertical size
	// .schema.traces.image.attributes.dy
	Dy types.NumberType `json:"dy,omitempty"`

	// Hoverinfo
	// arrayOK: true
	// default: x+y+z+text+name
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	// .schema.traces.image.attributes.hoverinfo
	Hoverinfo *types.ArrayOK[*ImageHoverinfo] `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hoverinfo`.
	// .schema.traces.image.attributes.hoverinfosrc
	Hoverinfosrc types.StringType `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// arrayOK: false
	// role: Object
	// .schema.traces.image.attributes.hoverlabel
	Hoverlabel *ImageHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate
	// arrayOK: true
	// type: string
	// Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}" as well as %{xother}, {%_xother}, {%_xother_}, {%xother_}. When showing info for several points, *xother* will be added to those with different x positions from the first point. An underscore before or after *(x|y)other* will add a space on that side, only when this field is shown. Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-format/tree/v1.4.5#d3-format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format/tree/v2.2.3#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. Finally, the template string has access to variables `z`, `color` and `colormodel`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	// .schema.traces.image.attributes.hovertemplate
	Hovertemplate *types.ArrayOK[*types.StringType] `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hovertemplate`.
	// .schema.traces.image.attributes.hovertemplatesrc
	Hovertemplatesrc types.StringType `json:"hovertemplatesrc,omitempty"`

	// Hovertext
	// arrayOK: false
	// type: data_array
	// Same as `text`.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.image.attributes.hovertext
	Hovertext *types.DataArrayType `json:"hovertext,omitempty"`

	// Hovertextsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `hovertext`.
	// .schema.traces.image.attributes.hovertextsrc
	Hovertextsrc types.StringType `json:"hovertextsrc,omitempty"`

	// Ids
	// arrayOK: false
	// type: data_array
	// Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.image.attributes.ids
	Ids *types.DataArrayType `json:"ids,omitempty"`

	// Idssrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `ids`.
	// .schema.traces.image.attributes.idssrc
	Idssrc types.StringType `json:"idssrc,omitempty"`

	// Legend
	// arrayOK: false
	// type: subplotid
	// Sets the reference to a legend to show this trace in. References to these legends are *legend*, *legend2*, *legend3*, etc. Settings for these legends are set in the layout, under `layout.legend`, `layout.legend2`, etc.
	// .schema.traces.image.attributes.legend
	Legend types.StringType `json:"legend,omitempty"`

	// Legendgrouptitle
	// arrayOK: false
	// role: Object
	// .schema.traces.image.attributes.legendgrouptitle
	Legendgrouptitle *ImageLegendgrouptitle `json:"legendgrouptitle,omitempty"`

	// Legendrank
	// arrayOK: false
	// type: number
	// Sets the legend rank for this trace. Items and groups with smaller ranks are presented on top/left side while with *reversed* `legend.traceorder` they are on bottom/right side. The default legendrank is 1000, so that you can use ranks less than 1000 to place certain items before all unranked items, and ranks greater than 1000 to go after all unranked items. When having unranked or equal rank items shapes would be displayed after traces i.e. according to their order in data and layout.
	// .schema.traces.image.attributes.legendrank
	Legendrank types.NumberType `json:"legendrank,omitempty"`

	// Legendwidth
	// arrayOK: false
	// type: number
	// Sets the width (in px or fraction) of the legend for this trace.
	// .schema.traces.image.attributes.legendwidth
	Legendwidth types.NumberType `json:"legendwidth,omitempty"`

	// Meta
	// arrayOK: true
	// type: any
	// Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	// .schema.traces.image.attributes.meta
	Meta *types.ArrayOK[*interface{}] `json:"meta,omitempty"`

	// Metasrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `meta`.
	// .schema.traces.image.attributes.metasrc
	Metasrc types.StringType `json:"metasrc,omitempty"`

	// Name
	// arrayOK: false
	// type: string
	// Sets the trace name. The trace name appears as the legend item and on hover.
	// .schema.traces.image.attributes.name
	Name types.StringType `json:"name,omitempty"`

	// Opacity
	// arrayOK: false
	// type: number
	// Sets the opacity of the trace.
	// .schema.traces.image.attributes.opacity
	Opacity types.NumberType `json:"opacity,omitempty"`

	// Source
	// arrayOK: false
	// type: string
	// Specifies the data URI of the image to be visualized. The URI consists of "data:image/[<media subtype>][;base64],<data>"
	// .schema.traces.image.attributes.source
	Source types.StringType `json:"source,omitempty"`

	// Stream
	// arrayOK: false
	// role: Object
	// .schema.traces.image.attributes.stream
	Stream *ImageStream `json:"stream,omitempty"`

	// Text
	// arrayOK: false
	// type: data_array
	// Sets the text elements associated with each z value.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.image.attributes.text
	Text *types.DataArrayType `json:"text,omitempty"`

	// Textsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `text`.
	// .schema.traces.image.attributes.textsrc
	Textsrc types.StringType `json:"textsrc,omitempty"`

	// Uid
	// arrayOK: false
	// type: string
	// Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	// .schema.traces.image.attributes.uid
	Uid types.StringType `json:"uid,omitempty"`

	// Uirevision
	// arrayOK: false
	// type: any
	// Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	// .schema.traces.image.attributes.uirevision
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible
	// arrayOK: false
	// default: %!s(bool=true)
	// type: enumerated
	// Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	// .schema.traces.image.attributes.visible
	Visible ImageVisible `json:"visible,omitempty"`

	// X0
	// arrayOK: false
	// type: any
	// Set the image's x position. The left edge of the image (or the right edge if the x axis is reversed or dx is negative) will be found at xmin=x0-dx/2
	// .schema.traces.image.attributes.x0
	X0 interface{} `json:"x0,omitempty"`

	// Xaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	// .schema.traces.image.attributes.xaxis
	Xaxis types.StringType `json:"xaxis,omitempty"`

	// Y0
	// arrayOK: false
	// type: any
	// Set the image's y position. The top edge of the image (or the bottom edge if the y axis is NOT reversed or if dy is negative) will be found at ymin=y0-dy/2. By default when an image trace is included, the y axis will be reversed so that the image is right-side-up, but you can disable this by setting yaxis.autorange=true or by providing an explicit y axis range.
	// .schema.traces.image.attributes.y0
	Y0 interface{} `json:"y0,omitempty"`

	// Yaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	// .schema.traces.image.attributes.yaxis
	Yaxis types.StringType `json:"yaxis,omitempty"`

	// Z
	// arrayOK: false
	// type: data_array
	// A 2-dimensional array in which each element is an array of 3 or 4 numbers representing a color.
	// use types.DataArray to pass any slice of data
	// use types.BDataArray to pass data in binary format as provided by numpy
	// .schema.traces.image.attributes.z
	Z *types.DataArrayType `json:"z,omitempty"`

	// Zmax
	// arrayOK: false
	// type: info_array
	// Array defining the higher bound for each color component. Note that the default value will depend on the colormodel. For the `rgb` colormodel, it is [255, 255, 255]. For the `rgba` colormodel, it is [255, 255, 255, 1]. For the `rgba256` colormodel, it is [255, 255, 255, 255]. For the `hsl` colormodel, it is [360, 100, 100]. For the `hsla` colormodel, it is [360, 100, 100, 1].
	// .schema.traces.image.attributes.zmax
	Zmax interface{} `json:"zmax,omitempty"`

	// Zmin
	// arrayOK: false
	// type: info_array
	// Array defining the lower bound for each color component. Note that the default value will depend on the colormodel. For the `rgb` colormodel, it is [0, 0, 0]. For the `rgba` colormodel, it is [0, 0, 0, 0]. For the `rgba256` colormodel, it is [0, 0, 0, 0]. For the `hsl` colormodel, it is [0, 0, 0]. For the `hsla` colormodel, it is [0, 0, 0, 0].
	// .schema.traces.image.attributes.zmin
	Zmin interface{} `json:"zmin,omitempty"`

	// Zorder
	// arrayOK: false
	// type: integer
	// Sets the layer on which this trace is displayed, relative to other SVG traces on the same subplot. SVG traces with higher `zorder` appear in front of those with lower `zorder`.
	// .schema.traces.image.attributes.zorder
	Zorder types.IntegerType `json:"zorder,omitempty"`

	// Zsmooth
	// arrayOK: false
	// default: %!s(bool=false)
	// type: enumerated
	// Picks a smoothing algorithm used to smooth `z` data. This only applies for image traces that use the `source` attribute.
	// .schema.traces.image.attributes.zsmooth
	Zsmooth ImageZsmooth `json:"zsmooth,omitempty"`

	// Zsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `z`.
	// .schema.traces.image.attributes.zsrc
	Zsrc types.StringType `json:"zsrc,omitempty"`
}

// ImageHoverlabelFont Sets the font used in hover labels.
type ImageHoverlabelFont struct {

	// Color
	// arrayOK: true
	// type: color
	//
	// .schema.traces.image.attributes.hoverlabel.font.color
	Color *types.ArrayOK[*types.Color] `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `color`.
	// .schema.traces.image.attributes.hoverlabel.font.colorsrc
	Colorsrc types.StringType `json:"colorsrc,omitempty"`

	// Family
	// arrayOK: true
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	// .schema.traces.image.attributes.hoverlabel.font.family
	Family *types.ArrayOK[*types.StringType] `json:"family,omitempty"`

	// Familysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `family`.
	// .schema.traces.image.attributes.hoverlabel.font.familysrc
	Familysrc types.StringType `json:"familysrc,omitempty"`

	// Lineposition
	// arrayOK: true
	// default: none
	// type: flaglist
	// Sets the kind of decoration line(s) with text, such as an *under*, *over* or *through* as well as combinations e.g. *under+over*, etc.
	// .schema.traces.image.attributes.hoverlabel.font.lineposition
	Lineposition *types.ArrayOK[*ImageHoverlabelFontLineposition] `json:"lineposition,omitempty"`

	// Linepositionsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `lineposition`.
	// .schema.traces.image.attributes.hoverlabel.font.linepositionsrc
	Linepositionsrc types.StringType `json:"linepositionsrc,omitempty"`

	// Shadow
	// arrayOK: true
	// type: string
	// Sets the shape and color of the shadow behind text. *auto* places minimal shadow and applies contrast text font color. See https://developer.mozilla.org/en-US/docs/Web/CSS/text-shadow for additional options.
	// .schema.traces.image.attributes.hoverlabel.font.shadow
	Shadow *types.ArrayOK[*types.StringType] `json:"shadow,omitempty"`

	// Shadowsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `shadow`.
	// .schema.traces.image.attributes.hoverlabel.font.shadowsrc
	Shadowsrc types.StringType `json:"shadowsrc,omitempty"`

	// Size
	// arrayOK: true
	// type: number
	//
	// .schema.traces.image.attributes.hoverlabel.font.size
	Size *types.ArrayOK[*types.NumberType] `json:"size,omitempty"`

	// Sizesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `size`.
	// .schema.traces.image.attributes.hoverlabel.font.sizesrc
	Sizesrc types.StringType `json:"sizesrc,omitempty"`

	// Style
	// arrayOK: true
	// default: normal
	// type: enumerated
	// Sets whether a font should be styled with a normal or italic face from its family.
	// .schema.traces.image.attributes.hoverlabel.font.style
	Style *types.ArrayOK[*ImageHoverlabelFontStyle] `json:"style,omitempty"`

	// Stylesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `style`.
	// .schema.traces.image.attributes.hoverlabel.font.stylesrc
	Stylesrc types.StringType `json:"stylesrc,omitempty"`

	// Textcase
	// arrayOK: true
	// default: normal
	// type: enumerated
	// Sets capitalization of text. It can be used to make text appear in all-uppercase or all-lowercase, or with each word capitalized.
	// .schema.traces.image.attributes.hoverlabel.font.textcase
	Textcase *types.ArrayOK[*ImageHoverlabelFontTextcase] `json:"textcase,omitempty"`

	// Textcasesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `textcase`.
	// .schema.traces.image.attributes.hoverlabel.font.textcasesrc
	Textcasesrc types.StringType `json:"textcasesrc,omitempty"`

	// Variant
	// arrayOK: true
	// default: normal
	// type: enumerated
	// Sets the variant of the font.
	// .schema.traces.image.attributes.hoverlabel.font.variant
	Variant *types.ArrayOK[*ImageHoverlabelFontVariant] `json:"variant,omitempty"`

	// Variantsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `variant`.
	// .schema.traces.image.attributes.hoverlabel.font.variantsrc
	Variantsrc types.StringType `json:"variantsrc,omitempty"`

	// Weight
	// arrayOK: true
	// type: integer
	// Sets the weight (or boldness) of the font.
	// .schema.traces.image.attributes.hoverlabel.font.weight
	Weight *types.ArrayOK[*types.IntegerType] `json:"weight,omitempty"`

	// Weightsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `weight`.
	// .schema.traces.image.attributes.hoverlabel.font.weightsrc
	Weightsrc types.StringType `json:"weightsrc,omitempty"`
}

// ImageHoverlabel
type ImageHoverlabel struct {

	// Align
	// arrayOK: true
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	// .schema.traces.image.attributes.hoverlabel.align
	Align *types.ArrayOK[*ImageHoverlabelAlign] `json:"align,omitempty"`

	// Alignsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `align`.
	// .schema.traces.image.attributes.hoverlabel.alignsrc
	Alignsrc types.StringType `json:"alignsrc,omitempty"`

	// Bgcolor
	// arrayOK: true
	// type: color
	// Sets the background color of the hover labels for this trace
	// .schema.traces.image.attributes.hoverlabel.bgcolor
	Bgcolor *types.ArrayOK[*types.Color] `json:"bgcolor,omitempty"`

	// Bgcolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `bgcolor`.
	// .schema.traces.image.attributes.hoverlabel.bgcolorsrc
	Bgcolorsrc types.StringType `json:"bgcolorsrc,omitempty"`

	// Bordercolor
	// arrayOK: true
	// type: color
	// Sets the border color of the hover labels for this trace.
	// .schema.traces.image.attributes.hoverlabel.bordercolor
	Bordercolor *types.ArrayOK[*types.Color] `json:"bordercolor,omitempty"`

	// Bordercolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `bordercolor`.
	// .schema.traces.image.attributes.hoverlabel.bordercolorsrc
	Bordercolorsrc types.StringType `json:"bordercolorsrc,omitempty"`

	// Font
	// arrayOK: false
	// role: Object
	// .schema.traces.image.attributes.hoverlabel.font
	Font *ImageHoverlabelFont `json:"font,omitempty"`

	// Namelength
	// arrayOK: true
	// type: integer
	// Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis.
	// .schema.traces.image.attributes.hoverlabel.namelength
	Namelength *types.ArrayOK[*types.IntegerType] `json:"namelength,omitempty"`

	// Namelengthsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for `namelength`.
	// .schema.traces.image.attributes.hoverlabel.namelengthsrc
	Namelengthsrc types.StringType `json:"namelengthsrc,omitempty"`
}

// ImageLegendgrouptitleFont Sets this legend group's title font.
type ImageLegendgrouptitleFont struct {

	// Color
	// arrayOK: false
	// type: color
	//
	// .schema.traces.image.attributes.legendgrouptitle.font.color
	Color types.Color `json:"color,omitempty"`

	// Family
	// arrayOK: false
	// type: string
	// HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*.
	// .schema.traces.image.attributes.legendgrouptitle.font.family
	Family types.StringType `json:"family,omitempty"`

	// Lineposition
	// arrayOK: false
	// default: none
	// type: flaglist
	// Sets the kind of decoration line(s) with text, such as an *under*, *over* or *through* as well as combinations e.g. *under+over*, etc.
	// .schema.traces.image.attributes.legendgrouptitle.font.lineposition
	Lineposition ImageLegendgrouptitleFontLineposition `json:"lineposition,omitempty"`

	// Shadow
	// arrayOK: false
	// type: string
	// Sets the shape and color of the shadow behind text. *auto* places minimal shadow and applies contrast text font color. See https://developer.mozilla.org/en-US/docs/Web/CSS/text-shadow for additional options.
	// .schema.traces.image.attributes.legendgrouptitle.font.shadow
	Shadow types.StringType `json:"shadow,omitempty"`

	// Size
	// arrayOK: false
	// type: number
	//
	// .schema.traces.image.attributes.legendgrouptitle.font.size
	Size types.NumberType `json:"size,omitempty"`

	// Style
	// arrayOK: false
	// default: normal
	// type: enumerated
	// Sets whether a font should be styled with a normal or italic face from its family.
	// .schema.traces.image.attributes.legendgrouptitle.font.style
	Style ImageLegendgrouptitleFontStyle `json:"style,omitempty"`

	// Textcase
	// arrayOK: false
	// default: normal
	// type: enumerated
	// Sets capitalization of text. It can be used to make text appear in all-uppercase or all-lowercase, or with each word capitalized.
	// .schema.traces.image.attributes.legendgrouptitle.font.textcase
	Textcase ImageLegendgrouptitleFontTextcase `json:"textcase,omitempty"`

	// Variant
	// arrayOK: false
	// default: normal
	// type: enumerated
	// Sets the variant of the font.
	// .schema.traces.image.attributes.legendgrouptitle.font.variant
	Variant ImageLegendgrouptitleFontVariant `json:"variant,omitempty"`

	// Weight
	// arrayOK: false
	// type: integer
	// Sets the weight (or boldness) of the font.
	// .schema.traces.image.attributes.legendgrouptitle.font.weight
	Weight types.IntegerType `json:"weight,omitempty"`
}

// ImageLegendgrouptitle
type ImageLegendgrouptitle struct {

	// Font
	// arrayOK: false
	// role: Object
	// .schema.traces.image.attributes.legendgrouptitle.font
	Font *ImageLegendgrouptitleFont `json:"font,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the legend group.
	// .schema.traces.image.attributes.legendgrouptitle.text
	Text types.StringType `json:"text,omitempty"`
}

// ImageStream
type ImageStream struct {

	// Maxpoints
	// arrayOK: false
	// type: number
	// Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot.
	// .schema.traces.image.attributes.stream.maxpoints
	Maxpoints types.NumberType `json:"maxpoints,omitempty"`

	// Token
	// arrayOK: false
	// type: string
	// The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details.
	// .schema.traces.image.attributes.stream.token
	Token types.StringType `json:"token,omitempty"`
}

// ImageColormodel Color model used to map the numerical color components described in `z` into colors. If `source` is specified, this attribute will be set to `rgba256` otherwise it defaults to `rgb`.
// .schema.traces.image.attributes.colormodel
type ImageColormodel string

const (
	ImageColormodelRgb     ImageColormodel = "rgb"
	ImageColormodelRgba    ImageColormodel = "rgba"
	ImageColormodelRgba256 ImageColormodel = "rgba256"
	ImageColormodelHsl     ImageColormodel = "hsl"
	ImageColormodelHsla    ImageColormodel = "hsla"
)

// ImageHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
// .schema.traces.image.attributes.hoverlabel.align
type ImageHoverlabelAlign string

const (
	ImageHoverlabelAlignLeft  ImageHoverlabelAlign = "left"
	ImageHoverlabelAlignRight ImageHoverlabelAlign = "right"
	ImageHoverlabelAlignAuto  ImageHoverlabelAlign = "auto"
)

// ImageHoverlabelFontStyle Sets whether a font should be styled with a normal or italic face from its family.
// .schema.traces.image.attributes.hoverlabel.font.style
type ImageHoverlabelFontStyle string

const (
	ImageHoverlabelFontStyleNormal ImageHoverlabelFontStyle = "normal"
	ImageHoverlabelFontStyleItalic ImageHoverlabelFontStyle = "italic"
)

// ImageHoverlabelFontTextcase Sets capitalization of text. It can be used to make text appear in all-uppercase or all-lowercase, or with each word capitalized.
// .schema.traces.image.attributes.hoverlabel.font.textcase
type ImageHoverlabelFontTextcase string

const (
	ImageHoverlabelFontTextcaseNormal   ImageHoverlabelFontTextcase = "normal"
	ImageHoverlabelFontTextcaseWordCaps ImageHoverlabelFontTextcase = "word caps"
	ImageHoverlabelFontTextcaseUpper    ImageHoverlabelFontTextcase = "upper"
	ImageHoverlabelFontTextcaseLower    ImageHoverlabelFontTextcase = "lower"
)

// ImageHoverlabelFontVariant Sets the variant of the font.
// .schema.traces.image.attributes.hoverlabel.font.variant
type ImageHoverlabelFontVariant string

const (
	ImageHoverlabelFontVariantNormal        ImageHoverlabelFontVariant = "normal"
	ImageHoverlabelFontVariantSmallCaps     ImageHoverlabelFontVariant = "small-caps"
	ImageHoverlabelFontVariantAllSmallCaps  ImageHoverlabelFontVariant = "all-small-caps"
	ImageHoverlabelFontVariantAllPetiteCaps ImageHoverlabelFontVariant = "all-petite-caps"
	ImageHoverlabelFontVariantPetiteCaps    ImageHoverlabelFontVariant = "petite-caps"
	ImageHoverlabelFontVariantUnicase       ImageHoverlabelFontVariant = "unicase"
)

// ImageLegendgrouptitleFontStyle Sets whether a font should be styled with a normal or italic face from its family.
// .schema.traces.image.attributes.legendgrouptitle.font.style
type ImageLegendgrouptitleFontStyle string

const (
	ImageLegendgrouptitleFontStyleNormal ImageLegendgrouptitleFontStyle = "normal"
	ImageLegendgrouptitleFontStyleItalic ImageLegendgrouptitleFontStyle = "italic"
)

// ImageLegendgrouptitleFontTextcase Sets capitalization of text. It can be used to make text appear in all-uppercase or all-lowercase, or with each word capitalized.
// .schema.traces.image.attributes.legendgrouptitle.font.textcase
type ImageLegendgrouptitleFontTextcase string

const (
	ImageLegendgrouptitleFontTextcaseNormal   ImageLegendgrouptitleFontTextcase = "normal"
	ImageLegendgrouptitleFontTextcaseWordCaps ImageLegendgrouptitleFontTextcase = "word caps"
	ImageLegendgrouptitleFontTextcaseUpper    ImageLegendgrouptitleFontTextcase = "upper"
	ImageLegendgrouptitleFontTextcaseLower    ImageLegendgrouptitleFontTextcase = "lower"
)

// ImageLegendgrouptitleFontVariant Sets the variant of the font.
// .schema.traces.image.attributes.legendgrouptitle.font.variant
type ImageLegendgrouptitleFontVariant string

const (
	ImageLegendgrouptitleFontVariantNormal        ImageLegendgrouptitleFontVariant = "normal"
	ImageLegendgrouptitleFontVariantSmallCaps     ImageLegendgrouptitleFontVariant = "small-caps"
	ImageLegendgrouptitleFontVariantAllSmallCaps  ImageLegendgrouptitleFontVariant = "all-small-caps"
	ImageLegendgrouptitleFontVariantAllPetiteCaps ImageLegendgrouptitleFontVariant = "all-petite-caps"
	ImageLegendgrouptitleFontVariantPetiteCaps    ImageLegendgrouptitleFontVariant = "petite-caps"
	ImageLegendgrouptitleFontVariantUnicase       ImageLegendgrouptitleFontVariant = "unicase"
)

// ImageVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
// .schema.traces.image.attributes.visible
type ImageVisible interface{}

var (
	ImageVisibleTrue       ImageVisible = true
	ImageVisibleFalse      ImageVisible = false
	ImageVisibleLegendonly ImageVisible = "legendonly"
)

// ImageZsmooth Picks a smoothing algorithm used to smooth `z` data. This only applies for image traces that use the `source` attribute.
// .schema.traces.image.attributes.zsmooth
type ImageZsmooth interface{}

var (
	ImageZsmoothFast  ImageZsmooth = "fast"
	ImageZsmoothFalse ImageZsmooth = false
)

// ImageHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
// .schema.traces.image.attributes.hoverinfo
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

// ImageHoverlabelFontLineposition Sets the kind of decoration line(s) with text, such as an *under*, *over* or *through* as well as combinations e.g. *under+over*, etc.
// .schema.traces.image.attributes.hoverlabel.font.lineposition
type ImageHoverlabelFontLineposition string

const (
	// Flags
	ImageHoverlabelFontLinepositionUnder   ImageHoverlabelFontLineposition = "under"
	ImageHoverlabelFontLinepositionOver    ImageHoverlabelFontLineposition = "over"
	ImageHoverlabelFontLinepositionThrough ImageHoverlabelFontLineposition = "through"

	// Extra
	ImageHoverlabelFontLinepositionNone ImageHoverlabelFontLineposition = "none"
)

// ImageLegendgrouptitleFontLineposition Sets the kind of decoration line(s) with text, such as an *under*, *over* or *through* as well as combinations e.g. *under+over*, etc.
// .schema.traces.image.attributes.legendgrouptitle.font.lineposition
type ImageLegendgrouptitleFontLineposition string

const (
	// Flags
	ImageLegendgrouptitleFontLinepositionUnder   ImageLegendgrouptitleFontLineposition = "under"
	ImageLegendgrouptitleFontLinepositionOver    ImageLegendgrouptitleFontLineposition = "over"
	ImageLegendgrouptitleFontLinepositionThrough ImageLegendgrouptitleFontLineposition = "through"

	// Extra
	ImageLegendgrouptitleFontLinepositionNone ImageLegendgrouptitleFontLineposition = "none"
)
