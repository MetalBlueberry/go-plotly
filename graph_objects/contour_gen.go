package grob

var TraceTypeContour TraceType = "contour"

func (trace *Contour) GetType() TraceType {
	return TraceTypeContour
}

// Contour The data from which contour lines are computed is set in `z`. Data in `z` must be a {2D array} of numbers. Say that `z` has N rows and M columns, then by default, these N rows correspond to N y coordinates (set in `y` or auto-generated) and the M columns correspond to M x coordinates (set in `x` or auto-generated). By setting `transpose` to *true*, the above behavior is flipped.
type Contour struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autocolorscale
	// arrayOK: false
	// type: boolean
	// Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Autocontour
	// arrayOK: false
	// type: boolean
	// Determines whether or not the contour level attributes are picked by an algorithm. If *true*, the number of contour levels can be set in `ncontours`. If *false*, set the contour level attributes in `contours`.
	Autocontour Bool `json:"autocontour,omitempty"`

	// Coloraxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar
	// role: Object
	Colorbar *ContourColorbar `json:"colorbar,omitempty"`

	// Colorscale
	// default: %!s(<nil>)
	// type: colorscale
	// Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`zmin` and `zmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	Colorscale ColorScale `json:"colorscale,omitempty"`

	// Connectgaps
	// arrayOK: false
	// type: boolean
	// Determines whether or not gaps (i.e. {nan} or missing values) in the `z` data are filled in. It is defaulted to true if `z` is a one dimensional array otherwise it is defaulted to false.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Contours
	// role: Object
	Contours *ContourContours `json:"contours,omitempty"`

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

	// Fillcolor
	// arrayOK: false
	// type: color
	// Sets the fill color if `contours.type` is *constraint*. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor Color `json:"fillcolor,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ContourHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *ContourHoverlabel `json:"hoverlabel,omitempty"`

	// Hoverongaps
	// arrayOK: false
	// type: boolean
	// Determines whether or not gaps (i.e. {nan} or missing values) in the `z` data have hover labels associated with them.
	Hoverongaps Bool `json:"hoverongaps,omitempty"`

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
	// arrayOK: false
	// type: data_array
	// Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

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

	// Line
	// role: Object
	Line *ContourLine `json:"line,omitempty"`

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

	// Ncontours
	// arrayOK: false
	// type: integer
	// Sets the maximum number of contour levels. The actual number of contours will be chosen automatically to be less than or equal to the value of `ncontours`. Has an effect only if `autocontour` is *true* or if `contours.size` is missing.
	Ncontours int64 `json:"ncontours,omitempty"`

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
	Stream *ContourStream `json:"stream,omitempty"`

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
	Visible ContourVisible `json:"visible,omitempty"`

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

	// Xcalendar
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `x` date data.
	Xcalendar ContourXcalendar `json:"xcalendar,omitempty"`

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
	Xperiodalignment ContourXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Xtype
	// default: %!s(<nil>)
	// type: enumerated
	// If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
	Xtype ContourXtype `json:"xtype,omitempty"`

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

	// Ycalendar
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `y` date data.
	Ycalendar ContourYcalendar `json:"ycalendar,omitempty"`

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
	Yperiodalignment ContourYperiodalignment `json:"yperiodalignment,omitempty"`

	// Ysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Ytype
	// default: %!s(<nil>)
	// type: enumerated
	// If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
	Ytype ContourYtype `json:"ytype,omitempty"`

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

	// Zhoverformat
	// arrayOK: false
	// type: string
	// Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. See: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format
	Zhoverformat String `json:"zhoverformat,omitempty"`

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

	// Zsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

// ContourColorbarTickfont Sets the color bar's tick label font
type ContourColorbarTickfont struct {

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

// ContourColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type ContourColorbarTitleFont struct {

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

// ContourColorbarTitle
type ContourColorbarTitle struct {

	// Font
	// role: Object
	Font *ContourColorbarTitleFont `json:"font,omitempty"`

	// Side
	// default: top
	// type: enumerated
	// Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
	Side ContourColorbarTitleSide `json:"side,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// ContourColorbar
type ContourColorbar struct {

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
	Exponentformat ContourColorbarExponentformat `json:"exponentformat,omitempty"`

	// Len
	// arrayOK: false
	// type: number
	// Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends.
	Len float64 `json:"len,omitempty"`

	// Lenmode
	// default: fraction
	// type: enumerated
	// Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
	Lenmode ContourColorbarLenmode `json:"lenmode,omitempty"`

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
	Showexponent ContourColorbarShowexponent `json:"showexponent,omitempty"`

	// Showticklabels
	// arrayOK: false
	// type: boolean
	// Determines whether or not the tick labels are drawn.
	Showticklabels Bool `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix ContourColorbarShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix ContourColorbarShowticksuffix `json:"showticksuffix,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels.
	Thickness float64 `json:"thickness,omitempty"`

	// Thicknessmode
	// default: pixels
	// type: enumerated
	// Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
	Thicknessmode ContourColorbarThicknessmode `json:"thicknessmode,omitempty"`

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
	Tickfont *ContourColorbarTickfont `json:"tickfont,omitempty"`

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
	Ticklabelposition ContourColorbarTicklabelposition `json:"ticklabelposition,omitempty"`

	// Ticklen
	// arrayOK: false
	// type: number
	// Sets the tick length (in px).
	Ticklen float64 `json:"ticklen,omitempty"`

	// Tickmode
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
	Tickmode ContourColorbarTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

	// Ticks
	// default:
	// type: enumerated
	// Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
	Ticks ContourColorbarTicks `json:"ticks,omitempty"`

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
	Title *ContourColorbarTitle `json:"title,omitempty"`

	// X
	// arrayOK: false
	// type: number
	// Sets the x position of the color bar (in plot fraction).
	X float64 `json:"x,omitempty"`

	// Xanchor
	// default: left
	// type: enumerated
	// Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
	Xanchor ContourColorbarXanchor `json:"xanchor,omitempty"`

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
	Yanchor ContourColorbarYanchor `json:"yanchor,omitempty"`

	// Ypad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the y direction.
	Ypad float64 `json:"ypad,omitempty"`
}

// ContourContoursLabelfont Sets the font used for labeling the contour levels. The default color comes from the lines, if shown. The default family and size come from `layout.font`.
type ContourContoursLabelfont struct {

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

// ContourContours
type ContourContours struct {

	// Coloring
	// default: fill
	// type: enumerated
	// Determines the coloring method showing the contour values. If *fill*, coloring is done evenly between each contour level If *heatmap*, a heatmap gradient coloring is applied between each contour level. If *lines*, coloring is done on the contour lines. If *none*, no coloring is applied on this trace.
	Coloring ContourContoursColoring `json:"coloring,omitempty"`

	// End
	// arrayOK: false
	// type: number
	// Sets the end contour level value. Must be more than `contours.start`
	End float64 `json:"end,omitempty"`

	// Labelfont
	// role: Object
	Labelfont *ContourContoursLabelfont `json:"labelfont,omitempty"`

	// Labelformat
	// arrayOK: false
	// type: string
	// Sets the contour label formatting rule using d3 formatting mini-language which is very similar to Python, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format
	Labelformat String `json:"labelformat,omitempty"`

	// Operation
	// default: =
	// type: enumerated
	// Sets the constraint operation. *=* keeps regions equal to `value` *<* and *<=* keep regions less than `value` *>* and *>=* keep regions greater than `value` *[]*, *()*, *[)*, and *(]* keep regions inside `value[0]` to `value[1]` *][*, *)(*, *](*, *)[* keep regions outside `value[0]` to value[1]` Open vs. closed intervals make no difference to constraint display, but all versions are allowed for consistency with filter transforms.
	Operation ContourContoursOperation `json:"operation,omitempty"`

	// Showlabels
	// arrayOK: false
	// type: boolean
	// Determines whether to label the contour lines with their values.
	Showlabels Bool `json:"showlabels,omitempty"`

	// Showlines
	// arrayOK: false
	// type: boolean
	// Determines whether or not the contour lines are drawn. Has an effect only if `contours.coloring` is set to *fill*.
	Showlines Bool `json:"showlines,omitempty"`

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

	// Type
	// default: levels
	// type: enumerated
	// If `levels`, the data is represented as a contour plot with multiple levels displayed. If `constraint`, the data is represented as constraints with the invalid region shaded as specified by the `operation` and `value` parameters.
	Type ContourContoursType `json:"type,omitempty"`

	// Value
	// arrayOK: false
	// type: any
	// Sets the value or values of the constraint boundary. When `operation` is set to one of the comparison values (=,<,>=,>,<=) *value* is expected to be a number. When `operation` is set to one of the interval values ([],(),[),(],][,)(,](,)[) *value* is expected to be an array of two numbers where the first is the lower bound and the second is the upper bound.
	Value interface{} `json:"value,omitempty"`
}

// ContourHoverlabelFont Sets the font used in hover labels.
type ContourHoverlabelFont struct {

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

// ContourHoverlabel
type ContourHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align ContourHoverlabelAlign `json:"align,omitempty"`

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
	Font *ContourHoverlabelFont `json:"font,omitempty"`

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

// ContourLine
type ContourLine struct {

	// Color
	// arrayOK: false
	// type: color
	// Sets the color of the contour level. Has no effect if `contours.coloring` is set to *lines*.
	Color Color `json:"color,omitempty"`

	// Dash
	// arrayOK: false
	// type: string
	// Sets the dash style of lines. Set to a dash type string (*solid*, *dot*, *dash*, *longdash*, *dashdot*, or *longdashdot*) or a dash length list in px (eg *5px,10px,2px,2px*).
	Dash String `json:"dash,omitempty"`

	// Smoothing
	// arrayOK: false
	// type: number
	// Sets the amount of smoothing for the contour lines, where *0* corresponds to no smoothing.
	Smoothing float64 `json:"smoothing,omitempty"`

	// Width
	// arrayOK: false
	// type: number
	// Sets the contour line width in (in px) Defaults to *0.5* when `contours.type` is *levels*. Defaults to *2* when `contour.type` is *constraint*.
	Width float64 `json:"width,omitempty"`
}

// ContourStream
type ContourStream struct {

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

// ContourColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ContourColorbarExponentformat string

const (
	ContourColorbarExponentformatNone  ContourColorbarExponentformat = "none"
	ContourColorbarExponentformatE1    ContourColorbarExponentformat = "e"
	ContourColorbarExponentformatE2    ContourColorbarExponentformat = "E"
	ContourColorbarExponentformatPower ContourColorbarExponentformat = "power"
	ContourColorbarExponentformatSi    ContourColorbarExponentformat = "SI"
	ContourColorbarExponentformatB     ContourColorbarExponentformat = "B"
)

// ContourColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ContourColorbarLenmode string

const (
	ContourColorbarLenmodeFraction ContourColorbarLenmode = "fraction"
	ContourColorbarLenmodePixels   ContourColorbarLenmode = "pixels"
)

// ContourColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ContourColorbarShowexponent string

const (
	ContourColorbarShowexponentAll   ContourColorbarShowexponent = "all"
	ContourColorbarShowexponentFirst ContourColorbarShowexponent = "first"
	ContourColorbarShowexponentLast  ContourColorbarShowexponent = "last"
	ContourColorbarShowexponentNone  ContourColorbarShowexponent = "none"
)

// ContourColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ContourColorbarShowtickprefix string

const (
	ContourColorbarShowtickprefixAll   ContourColorbarShowtickprefix = "all"
	ContourColorbarShowtickprefixFirst ContourColorbarShowtickprefix = "first"
	ContourColorbarShowtickprefixLast  ContourColorbarShowtickprefix = "last"
	ContourColorbarShowtickprefixNone  ContourColorbarShowtickprefix = "none"
)

// ContourColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ContourColorbarShowticksuffix string

const (
	ContourColorbarShowticksuffixAll   ContourColorbarShowticksuffix = "all"
	ContourColorbarShowticksuffixFirst ContourColorbarShowticksuffix = "first"
	ContourColorbarShowticksuffixLast  ContourColorbarShowticksuffix = "last"
	ContourColorbarShowticksuffixNone  ContourColorbarShowticksuffix = "none"
)

// ContourColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ContourColorbarThicknessmode string

const (
	ContourColorbarThicknessmodeFraction ContourColorbarThicknessmode = "fraction"
	ContourColorbarThicknessmodePixels   ContourColorbarThicknessmode = "pixels"
)

// ContourColorbarTicklabelposition Determines where tick labels are drawn.
type ContourColorbarTicklabelposition string

const (
	ContourColorbarTicklabelpositionOutside       ContourColorbarTicklabelposition = "outside"
	ContourColorbarTicklabelpositionInside        ContourColorbarTicklabelposition = "inside"
	ContourColorbarTicklabelpositionOutsideTop    ContourColorbarTicklabelposition = "outside top"
	ContourColorbarTicklabelpositionInsideTop     ContourColorbarTicklabelposition = "inside top"
	ContourColorbarTicklabelpositionOutsideBottom ContourColorbarTicklabelposition = "outside bottom"
	ContourColorbarTicklabelpositionInsideBottom  ContourColorbarTicklabelposition = "inside bottom"
)

// ContourColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ContourColorbarTickmode string

const (
	ContourColorbarTickmodeAuto   ContourColorbarTickmode = "auto"
	ContourColorbarTickmodeLinear ContourColorbarTickmode = "linear"
	ContourColorbarTickmodeArray  ContourColorbarTickmode = "array"
)

// ContourColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ContourColorbarTicks string

const (
	ContourColorbarTicksOutside ContourColorbarTicks = "outside"
	ContourColorbarTicksInside  ContourColorbarTicks = "inside"
	ContourColorbarTicksEmpty   ContourColorbarTicks = ""
)

// ContourColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ContourColorbarTitleSide string

const (
	ContourColorbarTitleSideRight  ContourColorbarTitleSide = "right"
	ContourColorbarTitleSideTop    ContourColorbarTitleSide = "top"
	ContourColorbarTitleSideBottom ContourColorbarTitleSide = "bottom"
)

// ContourColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ContourColorbarXanchor string

const (
	ContourColorbarXanchorLeft   ContourColorbarXanchor = "left"
	ContourColorbarXanchorCenter ContourColorbarXanchor = "center"
	ContourColorbarXanchorRight  ContourColorbarXanchor = "right"
)

// ContourColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ContourColorbarYanchor string

const (
	ContourColorbarYanchorTop    ContourColorbarYanchor = "top"
	ContourColorbarYanchorMiddle ContourColorbarYanchor = "middle"
	ContourColorbarYanchorBottom ContourColorbarYanchor = "bottom"
)

// ContourContoursColoring Determines the coloring method showing the contour values. If *fill*, coloring is done evenly between each contour level If *heatmap*, a heatmap gradient coloring is applied between each contour level. If *lines*, coloring is done on the contour lines. If *none*, no coloring is applied on this trace.
type ContourContoursColoring string

const (
	ContourContoursColoringFill    ContourContoursColoring = "fill"
	ContourContoursColoringHeatmap ContourContoursColoring = "heatmap"
	ContourContoursColoringLines   ContourContoursColoring = "lines"
	ContourContoursColoringNone    ContourContoursColoring = "none"
)

// ContourContoursOperation Sets the constraint operation. *=* keeps regions equal to `value` *<* and *<=* keep regions less than `value` *>* and *>=* keep regions greater than `value` *[]*, *()*, *[)*, and *(]* keep regions inside `value[0]` to `value[1]` *][*, *)(*, *](*, *)[* keep regions outside `value[0]` to value[1]` Open vs. closed intervals make no difference to constraint display, but all versions are allowed for consistency with filter transforms.
type ContourContoursOperation string

const (
	ContourContoursOperationEq               ContourContoursOperation = "="
	ContourContoursOperationLt               ContourContoursOperation = "<"
	ContourContoursOperationGtEq             ContourContoursOperation = ">="
	ContourContoursOperationGt               ContourContoursOperation = ">"
	ContourContoursOperationLtEq             ContourContoursOperation = "<="
	ContourContoursOperationLbracketRbracket ContourContoursOperation = "[]"
	ContourContoursOperationLparRpar         ContourContoursOperation = "()"
	ContourContoursOperationLbracketRpar     ContourContoursOperation = "[)"
	ContourContoursOperationLparRbracket     ContourContoursOperation = "(]"
	ContourContoursOperationRbracketLbracket ContourContoursOperation = "]["
	ContourContoursOperationRparLpar         ContourContoursOperation = ")("
	ContourContoursOperationRbracketLpar     ContourContoursOperation = "]("
	ContourContoursOperationRparLbracket     ContourContoursOperation = ")["
)

// ContourContoursType If `levels`, the data is represented as a contour plot with multiple levels displayed. If `constraint`, the data is represented as constraints with the invalid region shaded as specified by the `operation` and `value` parameters.
type ContourContoursType string

const (
	ContourContoursTypeLevels     ContourContoursType = "levels"
	ContourContoursTypeConstraint ContourContoursType = "constraint"
)

// ContourHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ContourHoverlabelAlign string

const (
	ContourHoverlabelAlignLeft  ContourHoverlabelAlign = "left"
	ContourHoverlabelAlignRight ContourHoverlabelAlign = "right"
	ContourHoverlabelAlignAuto  ContourHoverlabelAlign = "auto"
)

// ContourVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ContourVisible interface{}

var (
	ContourVisibleTrue       ContourVisible = true
	ContourVisibleFalse      ContourVisible = false
	ContourVisibleLegendonly ContourVisible = "legendonly"
)

// ContourXcalendar Sets the calendar system to use with `x` date data.
type ContourXcalendar string

const (
	ContourXcalendarGregorian  ContourXcalendar = "gregorian"
	ContourXcalendarChinese    ContourXcalendar = "chinese"
	ContourXcalendarCoptic     ContourXcalendar = "coptic"
	ContourXcalendarDiscworld  ContourXcalendar = "discworld"
	ContourXcalendarEthiopian  ContourXcalendar = "ethiopian"
	ContourXcalendarHebrew     ContourXcalendar = "hebrew"
	ContourXcalendarIslamic    ContourXcalendar = "islamic"
	ContourXcalendarJulian     ContourXcalendar = "julian"
	ContourXcalendarMayan      ContourXcalendar = "mayan"
	ContourXcalendarNanakshahi ContourXcalendar = "nanakshahi"
	ContourXcalendarNepali     ContourXcalendar = "nepali"
	ContourXcalendarPersian    ContourXcalendar = "persian"
	ContourXcalendarJalali     ContourXcalendar = "jalali"
	ContourXcalendarTaiwan     ContourXcalendar = "taiwan"
	ContourXcalendarThai       ContourXcalendar = "thai"
	ContourXcalendarUmmalqura  ContourXcalendar = "ummalqura"
)

// ContourXperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
type ContourXperiodalignment string

const (
	ContourXperiodalignmentStart  ContourXperiodalignment = "start"
	ContourXperiodalignmentMiddle ContourXperiodalignment = "middle"
	ContourXperiodalignmentEnd    ContourXperiodalignment = "end"
)

// ContourXtype If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
type ContourXtype string

const (
	ContourXtypeArray  ContourXtype = "array"
	ContourXtypeScaled ContourXtype = "scaled"
)

// ContourYcalendar Sets the calendar system to use with `y` date data.
type ContourYcalendar string

const (
	ContourYcalendarGregorian  ContourYcalendar = "gregorian"
	ContourYcalendarChinese    ContourYcalendar = "chinese"
	ContourYcalendarCoptic     ContourYcalendar = "coptic"
	ContourYcalendarDiscworld  ContourYcalendar = "discworld"
	ContourYcalendarEthiopian  ContourYcalendar = "ethiopian"
	ContourYcalendarHebrew     ContourYcalendar = "hebrew"
	ContourYcalendarIslamic    ContourYcalendar = "islamic"
	ContourYcalendarJulian     ContourYcalendar = "julian"
	ContourYcalendarMayan      ContourYcalendar = "mayan"
	ContourYcalendarNanakshahi ContourYcalendar = "nanakshahi"
	ContourYcalendarNepali     ContourYcalendar = "nepali"
	ContourYcalendarPersian    ContourYcalendar = "persian"
	ContourYcalendarJalali     ContourYcalendar = "jalali"
	ContourYcalendarTaiwan     ContourYcalendar = "taiwan"
	ContourYcalendarThai       ContourYcalendar = "thai"
	ContourYcalendarUmmalqura  ContourYcalendar = "ummalqura"
)

// ContourYperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
type ContourYperiodalignment string

const (
	ContourYperiodalignmentStart  ContourYperiodalignment = "start"
	ContourYperiodalignmentMiddle ContourYperiodalignment = "middle"
	ContourYperiodalignmentEnd    ContourYperiodalignment = "end"
)

// ContourYtype If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
type ContourYtype string

const (
	ContourYtypeArray  ContourYtype = "array"
	ContourYtypeScaled ContourYtype = "scaled"
)

// ContourHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type ContourHoverinfo string

const (
	// Flags
	ContourHoverinfoX    ContourHoverinfo = "x"
	ContourHoverinfoY    ContourHoverinfo = "y"
	ContourHoverinfoZ    ContourHoverinfo = "z"
	ContourHoverinfoText ContourHoverinfo = "text"
	ContourHoverinfoName ContourHoverinfo = "name"

	// Extra
	ContourHoverinfoAll  ContourHoverinfo = "all"
	ContourHoverinfoNone ContourHoverinfo = "none"
	ContourHoverinfoSkip ContourHoverinfo = "skip"
)
