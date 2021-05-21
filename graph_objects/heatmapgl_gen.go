package grob

var TraceTypeHeatmapgl TraceType = "heatmapgl"

func (trace *Heatmapgl) GetType() TraceType {
	return TraceTypeHeatmapgl
}

// Heatmapgl WebGL version of the heatmap trace type.
type Heatmapgl struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autocolorscale
	// arrayOK: false
	// type: boolean
	// Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Coloraxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar
	// role: Object
	Colorbar *HeatmapglColorbar `json:"colorbar,omitempty"`

	// Colorscale
	// default: %!s(<nil>)
	// type: colorscale
	// Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`zmin` and `zmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	Colorscale ColorScale `json:"colorscale,omitempty"`

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
	Hoverinfo HeatmapglHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *HeatmapglHoverlabel `json:"hoverlabel,omitempty"`

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

	// Reversescale
	// arrayOK: false
	// type: boolean
	// Reverses the color mapping if true. If true, `zmin` will correspond to the last color in the array and `zmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Showscale
	// arrayOK: false
	// type: boolean
	// Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Stream
	// role: Object
	Stream *HeatmapglStream `json:"stream,omitempty"`

	// Text
	// arrayOK: false
	// type: data_array
	// Sets the text elements associated with each z value.
	Text interface{} `json:"text,omitempty"`

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

	// Transpose
	// arrayOK: false
	// type: boolean
	// Transposes the z data.
	Transpose Bool `json:"transpose,omitempty"`

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
	Visible HeatmapglVisible `json:"visible,omitempty"`

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

	// Xsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Xtype
	// default: %!s(<nil>)
	// type: enumerated
	// If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
	Xtype HeatmapglXtype `json:"xtype,omitempty"`

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

	// Ysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Ytype
	// default: %!s(<nil>)
	// type: enumerated
	// If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
	Ytype HeatmapglYtype `json:"ytype,omitempty"`

	// Z
	// arrayOK: false
	// type: data_array
	// Sets the z data.
	Z interface{} `json:"z,omitempty"`

	// Zauto
	// arrayOK: false
	// type: boolean
	// Determines whether or not the color domain is computed with respect to the input data (here in `z`) or the bounds set in `zmin` and `zmax`  Defaults to `false` when `zmin` and `zmax` are set by the user.
	Zauto Bool `json:"zauto,omitempty"`

	// Zmax
	// arrayOK: false
	// type: number
	// Sets the upper bound of the color domain. Value should have the same units as in `z` and if set, `zmin` must be set as well.
	Zmax float64 `json:"zmax,omitempty"`

	// Zmid
	// arrayOK: false
	// type: number
	// Sets the mid-point of the color domain by scaling `zmin` and/or `zmax` to be equidistant to this point. Value should have the same units as in `z`. Has no effect when `zauto` is `false`.
	Zmid float64 `json:"zmid,omitempty"`

	// Zmin
	// arrayOK: false
	// type: number
	// Sets the lower bound of the color domain. Value should have the same units as in `z` and if set, `zmax` must be set as well.
	Zmin float64 `json:"zmin,omitempty"`

	// Zsmooth
	// default: fast
	// type: enumerated
	// Picks a smoothing algorithm use to smooth `z` data.
	Zsmooth HeatmapglZsmooth `json:"zsmooth,omitempty"`

	// Zsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

// HeatmapglColorbarTickfont Sets the color bar's tick label font
type HeatmapglColorbarTickfont struct {

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

// HeatmapglColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type HeatmapglColorbarTitleFont struct {

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

// HeatmapglColorbarTitle
type HeatmapglColorbarTitle struct {

	// Font
	// role: Object
	Font *HeatmapglColorbarTitleFont `json:"font,omitempty"`

	// Side
	// default: top
	// type: enumerated
	// Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
	Side HeatmapglColorbarTitleSide `json:"side,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// HeatmapglColorbar
type HeatmapglColorbar struct {

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
	Exponentformat HeatmapglColorbarExponentformat `json:"exponentformat,omitempty"`

	// Len
	// arrayOK: false
	// type: number
	// Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends.
	Len float64 `json:"len,omitempty"`

	// Lenmode
	// default: fraction
	// type: enumerated
	// Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
	Lenmode HeatmapglColorbarLenmode `json:"lenmode,omitempty"`

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
	Showexponent HeatmapglColorbarShowexponent `json:"showexponent,omitempty"`

	// Showticklabels
	// arrayOK: false
	// type: boolean
	// Determines whether or not the tick labels are drawn.
	Showticklabels Bool `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix HeatmapglColorbarShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix HeatmapglColorbarShowticksuffix `json:"showticksuffix,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels.
	Thickness float64 `json:"thickness,omitempty"`

	// Thicknessmode
	// default: pixels
	// type: enumerated
	// Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
	Thicknessmode HeatmapglColorbarThicknessmode `json:"thicknessmode,omitempty"`

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
	Tickfont *HeatmapglColorbarTickfont `json:"tickfont,omitempty"`

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
	Ticklabelposition HeatmapglColorbarTicklabelposition `json:"ticklabelposition,omitempty"`

	// Ticklen
	// arrayOK: false
	// type: number
	// Sets the tick length (in px).
	Ticklen float64 `json:"ticklen,omitempty"`

	// Tickmode
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
	Tickmode HeatmapglColorbarTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

	// Ticks
	// default:
	// type: enumerated
	// Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
	Ticks HeatmapglColorbarTicks `json:"ticks,omitempty"`

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
	Title *HeatmapglColorbarTitle `json:"title,omitempty"`

	// X
	// arrayOK: false
	// type: number
	// Sets the x position of the color bar (in plot fraction).
	X float64 `json:"x,omitempty"`

	// Xanchor
	// default: left
	// type: enumerated
	// Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
	Xanchor HeatmapglColorbarXanchor `json:"xanchor,omitempty"`

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
	Yanchor HeatmapglColorbarYanchor `json:"yanchor,omitempty"`

	// Ypad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the y direction.
	Ypad float64 `json:"ypad,omitempty"`
}

// HeatmapglHoverlabelFont Sets the font used in hover labels.
type HeatmapglHoverlabelFont struct {

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

// HeatmapglHoverlabel
type HeatmapglHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align HeatmapglHoverlabelAlign `json:"align,omitempty"`

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
	Font *HeatmapglHoverlabelFont `json:"font,omitempty"`

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

// HeatmapglStream
type HeatmapglStream struct {

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

// HeatmapglColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type HeatmapglColorbarExponentformat string

const (
	HeatmapglColorbarExponentformatNone  HeatmapglColorbarExponentformat = "none"
	HeatmapglColorbarExponentformatE1    HeatmapglColorbarExponentformat = "e"
	HeatmapglColorbarExponentformatE2    HeatmapglColorbarExponentformat = "E"
	HeatmapglColorbarExponentformatPower HeatmapglColorbarExponentformat = "power"
	HeatmapglColorbarExponentformatSi    HeatmapglColorbarExponentformat = "SI"
	HeatmapglColorbarExponentformatB     HeatmapglColorbarExponentformat = "B"
)

// HeatmapglColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type HeatmapglColorbarLenmode string

const (
	HeatmapglColorbarLenmodeFraction HeatmapglColorbarLenmode = "fraction"
	HeatmapglColorbarLenmodePixels   HeatmapglColorbarLenmode = "pixels"
)

// HeatmapglColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type HeatmapglColorbarShowexponent string

const (
	HeatmapglColorbarShowexponentAll   HeatmapglColorbarShowexponent = "all"
	HeatmapglColorbarShowexponentFirst HeatmapglColorbarShowexponent = "first"
	HeatmapglColorbarShowexponentLast  HeatmapglColorbarShowexponent = "last"
	HeatmapglColorbarShowexponentNone  HeatmapglColorbarShowexponent = "none"
)

// HeatmapglColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type HeatmapglColorbarShowtickprefix string

const (
	HeatmapglColorbarShowtickprefixAll   HeatmapglColorbarShowtickprefix = "all"
	HeatmapglColorbarShowtickprefixFirst HeatmapglColorbarShowtickprefix = "first"
	HeatmapglColorbarShowtickprefixLast  HeatmapglColorbarShowtickprefix = "last"
	HeatmapglColorbarShowtickprefixNone  HeatmapglColorbarShowtickprefix = "none"
)

// HeatmapglColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type HeatmapglColorbarShowticksuffix string

const (
	HeatmapglColorbarShowticksuffixAll   HeatmapglColorbarShowticksuffix = "all"
	HeatmapglColorbarShowticksuffixFirst HeatmapglColorbarShowticksuffix = "first"
	HeatmapglColorbarShowticksuffixLast  HeatmapglColorbarShowticksuffix = "last"
	HeatmapglColorbarShowticksuffixNone  HeatmapglColorbarShowticksuffix = "none"
)

// HeatmapglColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type HeatmapglColorbarThicknessmode string

const (
	HeatmapglColorbarThicknessmodeFraction HeatmapglColorbarThicknessmode = "fraction"
	HeatmapglColorbarThicknessmodePixels   HeatmapglColorbarThicknessmode = "pixels"
)

// HeatmapglColorbarTicklabelposition Determines where tick labels are drawn.
type HeatmapglColorbarTicklabelposition string

const (
	HeatmapglColorbarTicklabelpositionOutside       HeatmapglColorbarTicklabelposition = "outside"
	HeatmapglColorbarTicklabelpositionInside        HeatmapglColorbarTicklabelposition = "inside"
	HeatmapglColorbarTicklabelpositionOutsideTop    HeatmapglColorbarTicklabelposition = "outside top"
	HeatmapglColorbarTicklabelpositionInsideTop     HeatmapglColorbarTicklabelposition = "inside top"
	HeatmapglColorbarTicklabelpositionOutsideBottom HeatmapglColorbarTicklabelposition = "outside bottom"
	HeatmapglColorbarTicklabelpositionInsideBottom  HeatmapglColorbarTicklabelposition = "inside bottom"
)

// HeatmapglColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type HeatmapglColorbarTickmode string

const (
	HeatmapglColorbarTickmodeAuto   HeatmapglColorbarTickmode = "auto"
	HeatmapglColorbarTickmodeLinear HeatmapglColorbarTickmode = "linear"
	HeatmapglColorbarTickmodeArray  HeatmapglColorbarTickmode = "array"
)

// HeatmapglColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type HeatmapglColorbarTicks string

const (
	HeatmapglColorbarTicksOutside HeatmapglColorbarTicks = "outside"
	HeatmapglColorbarTicksInside  HeatmapglColorbarTicks = "inside"
	HeatmapglColorbarTicksEmpty   HeatmapglColorbarTicks = ""
)

// HeatmapglColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type HeatmapglColorbarTitleSide string

const (
	HeatmapglColorbarTitleSideRight  HeatmapglColorbarTitleSide = "right"
	HeatmapglColorbarTitleSideTop    HeatmapglColorbarTitleSide = "top"
	HeatmapglColorbarTitleSideBottom HeatmapglColorbarTitleSide = "bottom"
)

// HeatmapglColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type HeatmapglColorbarXanchor string

const (
	HeatmapglColorbarXanchorLeft   HeatmapglColorbarXanchor = "left"
	HeatmapglColorbarXanchorCenter HeatmapglColorbarXanchor = "center"
	HeatmapglColorbarXanchorRight  HeatmapglColorbarXanchor = "right"
)

// HeatmapglColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type HeatmapglColorbarYanchor string

const (
	HeatmapglColorbarYanchorTop    HeatmapglColorbarYanchor = "top"
	HeatmapglColorbarYanchorMiddle HeatmapglColorbarYanchor = "middle"
	HeatmapglColorbarYanchorBottom HeatmapglColorbarYanchor = "bottom"
)

// HeatmapglHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type HeatmapglHoverlabelAlign string

const (
	HeatmapglHoverlabelAlignLeft  HeatmapglHoverlabelAlign = "left"
	HeatmapglHoverlabelAlignRight HeatmapglHoverlabelAlign = "right"
	HeatmapglHoverlabelAlignAuto  HeatmapglHoverlabelAlign = "auto"
)

// HeatmapglVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type HeatmapglVisible interface{}

var (
	HeatmapglVisibleTrue       HeatmapglVisible = true
	HeatmapglVisibleFalse      HeatmapglVisible = false
	HeatmapglVisibleLegendonly HeatmapglVisible = "legendonly"
)

// HeatmapglXtype If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
type HeatmapglXtype string

const (
	HeatmapglXtypeArray  HeatmapglXtype = "array"
	HeatmapglXtypeScaled HeatmapglXtype = "scaled"
)

// HeatmapglYtype If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
type HeatmapglYtype string

const (
	HeatmapglYtypeArray  HeatmapglYtype = "array"
	HeatmapglYtypeScaled HeatmapglYtype = "scaled"
)

// HeatmapglZsmooth Picks a smoothing algorithm use to smooth `z` data.
type HeatmapglZsmooth interface{}

var (
	HeatmapglZsmoothFast  HeatmapglZsmooth = "fast"
	HeatmapglZsmoothFalse HeatmapglZsmooth = false
)

// HeatmapglHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type HeatmapglHoverinfo string

const (
	// Flags
	HeatmapglHoverinfoX    HeatmapglHoverinfo = "x"
	HeatmapglHoverinfoY    HeatmapglHoverinfo = "y"
	HeatmapglHoverinfoZ    HeatmapglHoverinfo = "z"
	HeatmapglHoverinfoText HeatmapglHoverinfo = "text"
	HeatmapglHoverinfoName HeatmapglHoverinfo = "name"

	// Extra
	HeatmapglHoverinfoAll  HeatmapglHoverinfo = "all"
	HeatmapglHoverinfoNone HeatmapglHoverinfo = "none"
	HeatmapglHoverinfoSkip HeatmapglHoverinfo = "skip"
)
