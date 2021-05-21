package grob

var TraceTypeSurface TraceType = "surface"

func (trace *Surface) GetType() TraceType {
	return TraceTypeSurface
}

// Surface The data the describes the coordinates of the surface is set in `z`. Data in `z` should be a {2D array}. Coordinates in `x` and `y` can either be 1D {arrays} or {2D arrays} (e.g. to graph parametric surfaces). If not provided in `x` and `y`, the x and y coordinates are assumed to be linear starting at 0 with a unit step. The color scale corresponds to the `z` values by default. For custom color scales, use `surfacecolor` which should be a {2D array}, where its bounds can be controlled using `cmin` and `cmax`.
type Surface struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autocolorscale
	// arrayOK: false
	// type: boolean
	// Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Cauto
	// arrayOK: false
	// type: boolean
	// Determines whether or not the color domain is computed with respect to the input data (here z or surfacecolor) or the bounds set in `cmin` and `cmax`  Defaults to `false` when `cmin` and `cmax` are set by the user.
	Cauto Bool `json:"cauto,omitempty"`

	// Cmax
	// arrayOK: false
	// type: number
	// Sets the upper bound of the color domain. Value should have the same units as z or surfacecolor and if set, `cmin` must be set as well.
	Cmax float64 `json:"cmax,omitempty"`

	// Cmid
	// arrayOK: false
	// type: number
	// Sets the mid-point of the color domain by scaling `cmin` and/or `cmax` to be equidistant to this point. Value should have the same units as z or surfacecolor. Has no effect when `cauto` is `false`.
	Cmid float64 `json:"cmid,omitempty"`

	// Cmin
	// arrayOK: false
	// type: number
	// Sets the lower bound of the color domain. Value should have the same units as z or surfacecolor and if set, `cmax` must be set as well.
	Cmin float64 `json:"cmin,omitempty"`

	// Coloraxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar
	// role: Object
	Colorbar *SurfaceColorbar `json:"colorbar,omitempty"`

	// Colorscale
	// default: %!s(<nil>)
	// type: colorscale
	// Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`cmin` and `cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	Colorscale ColorScale `json:"colorscale,omitempty"`

	// Connectgaps
	// arrayOK: false
	// type: boolean
	// Determines whether or not gaps (i.e. {nan} or missing values) in the `z` data are filled in.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Contours
	// role: Object
	Contours *SurfaceContours `json:"contours,omitempty"`

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

	// Hidesurface
	// arrayOK: false
	// type: boolean
	// Determines whether or not a surface is drawn. For example, set `hidesurface` to *false* `contours.x.show` to *true* and `contours.y.show` to *true* to draw a wire frame plot.
	Hidesurface Bool `json:"hidesurface,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo SurfaceHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *SurfaceHoverlabel `json:"hoverlabel,omitempty"`

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

	// Legendgroup
	// arrayOK: false
	// type: string
	// Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Lighting
	// role: Object
	Lighting *SurfaceLighting `json:"lighting,omitempty"`

	// Lightposition
	// role: Object
	Lightposition *SurfaceLightposition `json:"lightposition,omitempty"`

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
	// Sets the opacity of the surface. Please note that in the case of using high `opacity` values for example a value greater than or equal to 0.5 on two surfaces (and 0.25 with four surfaces), an overlay of multiple transparent surfaces may not perfectly be sorted in depth by the webgl API. This behavior may be improved in the near future and is subject to change.
	Opacity float64 `json:"opacity,omitempty"`

	// Opacityscale
	// arrayOK: false
	// type: any
	// Sets the opacityscale. The opacityscale must be an array containing arrays mapping a normalized value to an opacity value. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 1], [0.5, 0.2], [1, 1]]` means that higher/lower values would have higher opacity values and those in the middle would be more transparent Alternatively, `opacityscale` may be a palette name string of the following list: 'min', 'max', 'extremes' and 'uniform'. The default is 'uniform'.
	Opacityscale interface{} `json:"opacityscale,omitempty"`

	// Reversescale
	// arrayOK: false
	// type: boolean
	// Reverses the color mapping if true. If true, `cmin` will correspond to the last color in the array and `cmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Scene
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's 3D coordinate system and a 3D scene. If *scene* (the default value), the (x,y,z) coordinates refer to `layout.scene`. If *scene2*, the (x,y,z) coordinates refer to `layout.scene2`, and so on.
	Scene String `json:"scene,omitempty"`

	// Showlegend
	// arrayOK: false
	// type: boolean
	// Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale
	// arrayOK: false
	// type: boolean
	// Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Stream
	// role: Object
	Stream *SurfaceStream `json:"stream,omitempty"`

	// Surfacecolor
	// arrayOK: false
	// type: data_array
	// Sets the surface color values, used for setting a color scale independent of `z`.
	Surfacecolor interface{} `json:"surfacecolor,omitempty"`

	// Surfacecolorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  surfacecolor .
	Surfacecolorsrc String `json:"surfacecolorsrc,omitempty"`

	// Text
	// arrayOK: true
	// type: string
	// Sets the text elements associated with each z value. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text String `json:"text,omitempty"`

	// Textsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

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
	Visible SurfaceVisible `json:"visible,omitempty"`

	// X
	// arrayOK: false
	// type: data_array
	// Sets the x coordinates.
	X interface{} `json:"x,omitempty"`

	// Xcalendar
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `x` date data.
	Xcalendar SurfaceXcalendar `json:"xcalendar,omitempty"`

	// Xsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y
	// arrayOK: false
	// type: data_array
	// Sets the y coordinates.
	Y interface{} `json:"y,omitempty"`

	// Ycalendar
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `y` date data.
	Ycalendar SurfaceYcalendar `json:"ycalendar,omitempty"`

	// Ysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Z
	// arrayOK: false
	// type: data_array
	// Sets the z coordinates.
	Z interface{} `json:"z,omitempty"`

	// Zcalendar
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `z` date data.
	Zcalendar SurfaceZcalendar `json:"zcalendar,omitempty"`

	// Zsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

// SurfaceColorbarTickfont Sets the color bar's tick label font
type SurfaceColorbarTickfont struct {

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

// SurfaceColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type SurfaceColorbarTitleFont struct {

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

// SurfaceColorbarTitle
type SurfaceColorbarTitle struct {

	// Font
	// role: Object
	Font *SurfaceColorbarTitleFont `json:"font,omitempty"`

	// Side
	// default: top
	// type: enumerated
	// Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
	Side SurfaceColorbarTitleSide `json:"side,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// SurfaceColorbar
type SurfaceColorbar struct {

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
	Exponentformat SurfaceColorbarExponentformat `json:"exponentformat,omitempty"`

	// Len
	// arrayOK: false
	// type: number
	// Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends.
	Len float64 `json:"len,omitempty"`

	// Lenmode
	// default: fraction
	// type: enumerated
	// Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
	Lenmode SurfaceColorbarLenmode `json:"lenmode,omitempty"`

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
	Showexponent SurfaceColorbarShowexponent `json:"showexponent,omitempty"`

	// Showticklabels
	// arrayOK: false
	// type: boolean
	// Determines whether or not the tick labels are drawn.
	Showticklabels Bool `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix SurfaceColorbarShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix SurfaceColorbarShowticksuffix `json:"showticksuffix,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels.
	Thickness float64 `json:"thickness,omitempty"`

	// Thicknessmode
	// default: pixels
	// type: enumerated
	// Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
	Thicknessmode SurfaceColorbarThicknessmode `json:"thicknessmode,omitempty"`

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
	Tickfont *SurfaceColorbarTickfont `json:"tickfont,omitempty"`

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
	Ticklabelposition SurfaceColorbarTicklabelposition `json:"ticklabelposition,omitempty"`

	// Ticklen
	// arrayOK: false
	// type: number
	// Sets the tick length (in px).
	Ticklen float64 `json:"ticklen,omitempty"`

	// Tickmode
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
	Tickmode SurfaceColorbarTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

	// Ticks
	// default:
	// type: enumerated
	// Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
	Ticks SurfaceColorbarTicks `json:"ticks,omitempty"`

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
	Title *SurfaceColorbarTitle `json:"title,omitempty"`

	// X
	// arrayOK: false
	// type: number
	// Sets the x position of the color bar (in plot fraction).
	X float64 `json:"x,omitempty"`

	// Xanchor
	// default: left
	// type: enumerated
	// Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
	Xanchor SurfaceColorbarXanchor `json:"xanchor,omitempty"`

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
	Yanchor SurfaceColorbarYanchor `json:"yanchor,omitempty"`

	// Ypad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the y direction.
	Ypad float64 `json:"ypad,omitempty"`
}

// SurfaceContoursXProject
type SurfaceContoursXProject struct {

	// X
	// arrayOK: false
	// type: boolean
	// Determines whether or not these contour lines are projected on the x plane. If `highlight` is set to *true* (the default), the projected lines are shown on hover. If `show` is set to *true*, the projected lines are shown in permanence.
	X Bool `json:"x,omitempty"`

	// Y
	// arrayOK: false
	// type: boolean
	// Determines whether or not these contour lines are projected on the y plane. If `highlight` is set to *true* (the default), the projected lines are shown on hover. If `show` is set to *true*, the projected lines are shown in permanence.
	Y Bool `json:"y,omitempty"`

	// Z
	// arrayOK: false
	// type: boolean
	// Determines whether or not these contour lines are projected on the z plane. If `highlight` is set to *true* (the default), the projected lines are shown on hover. If `show` is set to *true*, the projected lines are shown in permanence.
	Z Bool `json:"z,omitempty"`
}

// SurfaceContoursX
type SurfaceContoursX struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the color of the contour lines.
	Color Color `json:"color,omitempty"`

	// End
	// arrayOK: false
	// type: number
	// Sets the end contour level value. Must be more than `contours.start`
	End float64 `json:"end,omitempty"`

	// Highlight
	// arrayOK: false
	// type: boolean
	// Determines whether or not contour lines about the x dimension are highlighted on hover.
	Highlight Bool `json:"highlight,omitempty"`

	// Highlightcolor
	// arrayOK: false
	// type: color
	// Sets the color of the highlighted contour lines.
	Highlightcolor Color `json:"highlightcolor,omitempty"`

	// Highlightwidth
	// arrayOK: false
	// type: number
	// Sets the width of the highlighted contour lines.
	Highlightwidth float64 `json:"highlightwidth,omitempty"`

	// Project
	// role: Object
	Project *SurfaceContoursXProject `json:"project,omitempty"`

	// Show
	// arrayOK: false
	// type: boolean
	// Determines whether or not contour lines about the x dimension are drawn.
	Show Bool `json:"show,omitempty"`

	// Size
	// arrayOK: false
	// type: number
	// Sets the step between each contour level. Must be positive.
	Size float64 `json:"size,omitempty"`

	// Start
	// arrayOK: false
	// type: number
	// Sets the starting contour level value. Must be less than `contours.end`
	Start float64 `json:"start,omitempty"`

	// Usecolormap
	// arrayOK: false
	// type: boolean
	// An alternate to *color*. Determines whether or not the contour lines are colored using the trace *colorscale*.
	Usecolormap Bool `json:"usecolormap,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the width of the contour lines.
	Width float64 `json:"width,omitempty"`
}

// SurfaceContoursYProject
type SurfaceContoursYProject struct {

	// X
	// arrayOK: false
	// type: boolean
	// Determines whether or not these contour lines are projected on the x plane. If `highlight` is set to *true* (the default), the projected lines are shown on hover. If `show` is set to *true*, the projected lines are shown in permanence.
	X Bool `json:"x,omitempty"`

	// Y
	// arrayOK: false
	// type: boolean
	// Determines whether or not these contour lines are projected on the y plane. If `highlight` is set to *true* (the default), the projected lines are shown on hover. If `show` is set to *true*, the projected lines are shown in permanence.
	Y Bool `json:"y,omitempty"`

	// Z
	// arrayOK: false
	// type: boolean
	// Determines whether or not these contour lines are projected on the z plane. If `highlight` is set to *true* (the default), the projected lines are shown on hover. If `show` is set to *true*, the projected lines are shown in permanence.
	Z Bool `json:"z,omitempty"`
}

// SurfaceContoursY
type SurfaceContoursY struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the color of the contour lines.
	Color Color `json:"color,omitempty"`

	// End
	// arrayOK: false
	// type: number
	// Sets the end contour level value. Must be more than `contours.start`
	End float64 `json:"end,omitempty"`

	// Highlight
	// arrayOK: false
	// type: boolean
	// Determines whether or not contour lines about the y dimension are highlighted on hover.
	Highlight Bool `json:"highlight,omitempty"`

	// Highlightcolor
	// arrayOK: false
	// type: color
	// Sets the color of the highlighted contour lines.
	Highlightcolor Color `json:"highlightcolor,omitempty"`

	// Highlightwidth
	// arrayOK: false
	// type: number
	// Sets the width of the highlighted contour lines.
	Highlightwidth float64 `json:"highlightwidth,omitempty"`

	// Project
	// role: Object
	Project *SurfaceContoursYProject `json:"project,omitempty"`

	// Show
	// arrayOK: false
	// type: boolean
	// Determines whether or not contour lines about the y dimension are drawn.
	Show Bool `json:"show,omitempty"`

	// Size
	// arrayOK: false
	// type: number
	// Sets the step between each contour level. Must be positive.
	Size float64 `json:"size,omitempty"`

	// Start
	// arrayOK: false
	// type: number
	// Sets the starting contour level value. Must be less than `contours.end`
	Start float64 `json:"start,omitempty"`

	// Usecolormap
	// arrayOK: false
	// type: boolean
	// An alternate to *color*. Determines whether or not the contour lines are colored using the trace *colorscale*.
	Usecolormap Bool `json:"usecolormap,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the width of the contour lines.
	Width float64 `json:"width,omitempty"`
}

// SurfaceContoursZProject
type SurfaceContoursZProject struct {

	// X
	// arrayOK: false
	// type: boolean
	// Determines whether or not these contour lines are projected on the x plane. If `highlight` is set to *true* (the default), the projected lines are shown on hover. If `show` is set to *true*, the projected lines are shown in permanence.
	X Bool `json:"x,omitempty"`

	// Y
	// arrayOK: false
	// type: boolean
	// Determines whether or not these contour lines are projected on the y plane. If `highlight` is set to *true* (the default), the projected lines are shown on hover. If `show` is set to *true*, the projected lines are shown in permanence.
	Y Bool `json:"y,omitempty"`

	// Z
	// arrayOK: false
	// type: boolean
	// Determines whether or not these contour lines are projected on the z plane. If `highlight` is set to *true* (the default), the projected lines are shown on hover. If `show` is set to *true*, the projected lines are shown in permanence.
	Z Bool `json:"z,omitempty"`
}

// SurfaceContoursZ
type SurfaceContoursZ struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the color of the contour lines.
	Color Color `json:"color,omitempty"`

	// End
	// arrayOK: false
	// type: number
	// Sets the end contour level value. Must be more than `contours.start`
	End float64 `json:"end,omitempty"`

	// Highlight
	// arrayOK: false
	// type: boolean
	// Determines whether or not contour lines about the z dimension are highlighted on hover.
	Highlight Bool `json:"highlight,omitempty"`

	// Highlightcolor
	// arrayOK: false
	// type: color
	// Sets the color of the highlighted contour lines.
	Highlightcolor Color `json:"highlightcolor,omitempty"`

	// Highlightwidth
	// arrayOK: false
	// type: number
	// Sets the width of the highlighted contour lines.
	Highlightwidth float64 `json:"highlightwidth,omitempty"`

	// Project
	// role: Object
	Project *SurfaceContoursZProject `json:"project,omitempty"`

	// Show
	// arrayOK: false
	// type: boolean
	// Determines whether or not contour lines about the z dimension are drawn.
	Show Bool `json:"show,omitempty"`

	// Size
	// arrayOK: false
	// type: number
	// Sets the step between each contour level. Must be positive.
	Size float64 `json:"size,omitempty"`

	// Start
	// arrayOK: false
	// type: number
	// Sets the starting contour level value. Must be less than `contours.end`
	Start float64 `json:"start,omitempty"`

	// Usecolormap
	// arrayOK: false
	// type: boolean
	// An alternate to *color*. Determines whether or not the contour lines are colored using the trace *colorscale*.
	Usecolormap Bool `json:"usecolormap,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the width of the contour lines.
	Width float64 `json:"width,omitempty"`
}

// SurfaceContours
type SurfaceContours struct {

	// X
	// role: Object
	X *SurfaceContoursX `json:"x,omitempty"`

	// Y
	// role: Object
	Y *SurfaceContoursY `json:"y,omitempty"`

	// Z
	// role: Object
	Z *SurfaceContoursZ `json:"z,omitempty"`
}

// SurfaceHoverlabelFont Sets the font used in hover labels.
type SurfaceHoverlabelFont struct {

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

// SurfaceHoverlabel
type SurfaceHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align SurfaceHoverlabelAlign `json:"align,omitempty"`

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
	Font *SurfaceHoverlabelFont `json:"font,omitempty"`

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

// SurfaceLighting
type SurfaceLighting struct {

	// Ambient
	// arrayOK: false
	// type: number
	// Ambient light increases overall color visibility but can wash out the image.
	Ambient float64 `json:"ambient,omitempty"`

	// Diffuse
	// arrayOK: false
	// type: number
	// Represents the extent that incident rays are reflected in a range of angles.
	Diffuse float64 `json:"diffuse,omitempty"`

	// Fresnel
	// arrayOK: false
	// type: number
	// Represents the reflectance as a dependency of the viewing angle; e.g. paper is reflective when viewing it from the edge of the paper (almost 90 degrees), causing shine.
	Fresnel float64 `json:"fresnel,omitempty"`

	// Roughness
	// arrayOK: false
	// type: number
	// Alters specular reflection; the rougher the surface, the wider and less contrasty the shine.
	Roughness float64 `json:"roughness,omitempty"`

	// Specular
	// arrayOK: false
	// type: number
	// Represents the level that incident rays are reflected in a single direction, causing shine.
	Specular float64 `json:"specular,omitempty"`
}

// SurfaceLightposition
type SurfaceLightposition struct {

	// X
	// arrayOK: false
	// type: number
	// Numeric vector, representing the X coordinate for each vertex.
	X float64 `json:"x,omitempty"`

	// Y
	// arrayOK: false
	// type: number
	// Numeric vector, representing the Y coordinate for each vertex.
	Y float64 `json:"y,omitempty"`

	// Z
	// arrayOK: false
	// type: number
	// Numeric vector, representing the Z coordinate for each vertex.
	Z float64 `json:"z,omitempty"`
}

// SurfaceStream
type SurfaceStream struct {

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

// SurfaceColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type SurfaceColorbarExponentformat string

const (
	SurfaceColorbarExponentformatNone  SurfaceColorbarExponentformat = "none"
	SurfaceColorbarExponentformatE1    SurfaceColorbarExponentformat = "e"
	SurfaceColorbarExponentformatE2    SurfaceColorbarExponentformat = "E"
	SurfaceColorbarExponentformatPower SurfaceColorbarExponentformat = "power"
	SurfaceColorbarExponentformatSi    SurfaceColorbarExponentformat = "SI"
	SurfaceColorbarExponentformatB     SurfaceColorbarExponentformat = "B"
)

// SurfaceColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type SurfaceColorbarLenmode string

const (
	SurfaceColorbarLenmodeFraction SurfaceColorbarLenmode = "fraction"
	SurfaceColorbarLenmodePixels   SurfaceColorbarLenmode = "pixels"
)

// SurfaceColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type SurfaceColorbarShowexponent string

const (
	SurfaceColorbarShowexponentAll   SurfaceColorbarShowexponent = "all"
	SurfaceColorbarShowexponentFirst SurfaceColorbarShowexponent = "first"
	SurfaceColorbarShowexponentLast  SurfaceColorbarShowexponent = "last"
	SurfaceColorbarShowexponentNone  SurfaceColorbarShowexponent = "none"
)

// SurfaceColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type SurfaceColorbarShowtickprefix string

const (
	SurfaceColorbarShowtickprefixAll   SurfaceColorbarShowtickprefix = "all"
	SurfaceColorbarShowtickprefixFirst SurfaceColorbarShowtickprefix = "first"
	SurfaceColorbarShowtickprefixLast  SurfaceColorbarShowtickprefix = "last"
	SurfaceColorbarShowtickprefixNone  SurfaceColorbarShowtickprefix = "none"
)

// SurfaceColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type SurfaceColorbarShowticksuffix string

const (
	SurfaceColorbarShowticksuffixAll   SurfaceColorbarShowticksuffix = "all"
	SurfaceColorbarShowticksuffixFirst SurfaceColorbarShowticksuffix = "first"
	SurfaceColorbarShowticksuffixLast  SurfaceColorbarShowticksuffix = "last"
	SurfaceColorbarShowticksuffixNone  SurfaceColorbarShowticksuffix = "none"
)

// SurfaceColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type SurfaceColorbarThicknessmode string

const (
	SurfaceColorbarThicknessmodeFraction SurfaceColorbarThicknessmode = "fraction"
	SurfaceColorbarThicknessmodePixels   SurfaceColorbarThicknessmode = "pixels"
)

// SurfaceColorbarTicklabelposition Determines where tick labels are drawn.
type SurfaceColorbarTicklabelposition string

const (
	SurfaceColorbarTicklabelpositionOutside       SurfaceColorbarTicklabelposition = "outside"
	SurfaceColorbarTicklabelpositionInside        SurfaceColorbarTicklabelposition = "inside"
	SurfaceColorbarTicklabelpositionOutsideTop    SurfaceColorbarTicklabelposition = "outside top"
	SurfaceColorbarTicklabelpositionInsideTop     SurfaceColorbarTicklabelposition = "inside top"
	SurfaceColorbarTicklabelpositionOutsideBottom SurfaceColorbarTicklabelposition = "outside bottom"
	SurfaceColorbarTicklabelpositionInsideBottom  SurfaceColorbarTicklabelposition = "inside bottom"
)

// SurfaceColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type SurfaceColorbarTickmode string

const (
	SurfaceColorbarTickmodeAuto   SurfaceColorbarTickmode = "auto"
	SurfaceColorbarTickmodeLinear SurfaceColorbarTickmode = "linear"
	SurfaceColorbarTickmodeArray  SurfaceColorbarTickmode = "array"
)

// SurfaceColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type SurfaceColorbarTicks string

const (
	SurfaceColorbarTicksOutside SurfaceColorbarTicks = "outside"
	SurfaceColorbarTicksInside  SurfaceColorbarTicks = "inside"
	SurfaceColorbarTicksEmpty   SurfaceColorbarTicks = ""
)

// SurfaceColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type SurfaceColorbarTitleSide string

const (
	SurfaceColorbarTitleSideRight  SurfaceColorbarTitleSide = "right"
	SurfaceColorbarTitleSideTop    SurfaceColorbarTitleSide = "top"
	SurfaceColorbarTitleSideBottom SurfaceColorbarTitleSide = "bottom"
)

// SurfaceColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type SurfaceColorbarXanchor string

const (
	SurfaceColorbarXanchorLeft   SurfaceColorbarXanchor = "left"
	SurfaceColorbarXanchorCenter SurfaceColorbarXanchor = "center"
	SurfaceColorbarXanchorRight  SurfaceColorbarXanchor = "right"
)

// SurfaceColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type SurfaceColorbarYanchor string

const (
	SurfaceColorbarYanchorTop    SurfaceColorbarYanchor = "top"
	SurfaceColorbarYanchorMiddle SurfaceColorbarYanchor = "middle"
	SurfaceColorbarYanchorBottom SurfaceColorbarYanchor = "bottom"
)

// SurfaceHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SurfaceHoverlabelAlign string

const (
	SurfaceHoverlabelAlignLeft  SurfaceHoverlabelAlign = "left"
	SurfaceHoverlabelAlignRight SurfaceHoverlabelAlign = "right"
	SurfaceHoverlabelAlignAuto  SurfaceHoverlabelAlign = "auto"
)

// SurfaceVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type SurfaceVisible interface{}

var (
	SurfaceVisibleTrue       SurfaceVisible = true
	SurfaceVisibleFalse      SurfaceVisible = false
	SurfaceVisibleLegendonly SurfaceVisible = "legendonly"
)

// SurfaceXcalendar Sets the calendar system to use with `x` date data.
type SurfaceXcalendar string

const (
	SurfaceXcalendarGregorian  SurfaceXcalendar = "gregorian"
	SurfaceXcalendarChinese    SurfaceXcalendar = "chinese"
	SurfaceXcalendarCoptic     SurfaceXcalendar = "coptic"
	SurfaceXcalendarDiscworld  SurfaceXcalendar = "discworld"
	SurfaceXcalendarEthiopian  SurfaceXcalendar = "ethiopian"
	SurfaceXcalendarHebrew     SurfaceXcalendar = "hebrew"
	SurfaceXcalendarIslamic    SurfaceXcalendar = "islamic"
	SurfaceXcalendarJulian     SurfaceXcalendar = "julian"
	SurfaceXcalendarMayan      SurfaceXcalendar = "mayan"
	SurfaceXcalendarNanakshahi SurfaceXcalendar = "nanakshahi"
	SurfaceXcalendarNepali     SurfaceXcalendar = "nepali"
	SurfaceXcalendarPersian    SurfaceXcalendar = "persian"
	SurfaceXcalendarJalali     SurfaceXcalendar = "jalali"
	SurfaceXcalendarTaiwan     SurfaceXcalendar = "taiwan"
	SurfaceXcalendarThai       SurfaceXcalendar = "thai"
	SurfaceXcalendarUmmalqura  SurfaceXcalendar = "ummalqura"
)

// SurfaceYcalendar Sets the calendar system to use with `y` date data.
type SurfaceYcalendar string

const (
	SurfaceYcalendarGregorian  SurfaceYcalendar = "gregorian"
	SurfaceYcalendarChinese    SurfaceYcalendar = "chinese"
	SurfaceYcalendarCoptic     SurfaceYcalendar = "coptic"
	SurfaceYcalendarDiscworld  SurfaceYcalendar = "discworld"
	SurfaceYcalendarEthiopian  SurfaceYcalendar = "ethiopian"
	SurfaceYcalendarHebrew     SurfaceYcalendar = "hebrew"
	SurfaceYcalendarIslamic    SurfaceYcalendar = "islamic"
	SurfaceYcalendarJulian     SurfaceYcalendar = "julian"
	SurfaceYcalendarMayan      SurfaceYcalendar = "mayan"
	SurfaceYcalendarNanakshahi SurfaceYcalendar = "nanakshahi"
	SurfaceYcalendarNepali     SurfaceYcalendar = "nepali"
	SurfaceYcalendarPersian    SurfaceYcalendar = "persian"
	SurfaceYcalendarJalali     SurfaceYcalendar = "jalali"
	SurfaceYcalendarTaiwan     SurfaceYcalendar = "taiwan"
	SurfaceYcalendarThai       SurfaceYcalendar = "thai"
	SurfaceYcalendarUmmalqura  SurfaceYcalendar = "ummalqura"
)

// SurfaceZcalendar Sets the calendar system to use with `z` date data.
type SurfaceZcalendar string

const (
	SurfaceZcalendarGregorian  SurfaceZcalendar = "gregorian"
	SurfaceZcalendarChinese    SurfaceZcalendar = "chinese"
	SurfaceZcalendarCoptic     SurfaceZcalendar = "coptic"
	SurfaceZcalendarDiscworld  SurfaceZcalendar = "discworld"
	SurfaceZcalendarEthiopian  SurfaceZcalendar = "ethiopian"
	SurfaceZcalendarHebrew     SurfaceZcalendar = "hebrew"
	SurfaceZcalendarIslamic    SurfaceZcalendar = "islamic"
	SurfaceZcalendarJulian     SurfaceZcalendar = "julian"
	SurfaceZcalendarMayan      SurfaceZcalendar = "mayan"
	SurfaceZcalendarNanakshahi SurfaceZcalendar = "nanakshahi"
	SurfaceZcalendarNepali     SurfaceZcalendar = "nepali"
	SurfaceZcalendarPersian    SurfaceZcalendar = "persian"
	SurfaceZcalendarJalali     SurfaceZcalendar = "jalali"
	SurfaceZcalendarTaiwan     SurfaceZcalendar = "taiwan"
	SurfaceZcalendarThai       SurfaceZcalendar = "thai"
	SurfaceZcalendarUmmalqura  SurfaceZcalendar = "ummalqura"
)

// SurfaceHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type SurfaceHoverinfo string

const (
	// Flags
	SurfaceHoverinfoX    SurfaceHoverinfo = "x"
	SurfaceHoverinfoY    SurfaceHoverinfo = "y"
	SurfaceHoverinfoZ    SurfaceHoverinfo = "z"
	SurfaceHoverinfoText SurfaceHoverinfo = "text"
	SurfaceHoverinfoName SurfaceHoverinfo = "name"

	// Extra
	SurfaceHoverinfoAll  SurfaceHoverinfo = "all"
	SurfaceHoverinfoNone SurfaceHoverinfo = "none"
	SurfaceHoverinfoSkip SurfaceHoverinfo = "skip"
)
