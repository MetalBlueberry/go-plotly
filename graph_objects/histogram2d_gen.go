package grob

var TraceTypeHistogram2d TraceType = "histogram2d"

func (trace *Histogram2d) GetType() TraceType {
	return TraceTypeHistogram2d
}

// Histogram2d The sample data from which statistics are computed is set in `x` and `y` (where `x` and `y` represent marginal distributions, binning is set in `xbins` and `ybins` in this case) or `z` (where `z` represent the 2D distribution and binning set, binning is set by `x` and `y` in this case). The resulting distribution is visualized as a heatmap.
type Histogram2d struct {

	// Type
	// is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autobinx
	// arrayOK: false
	// type: boolean
	// Obsolete: since v1.42 each bin attribute is auto-determined separately and `autobinx` is not needed. However, we accept `autobinx: true` or `false` and will update `xbins` accordingly before deleting `autobinx` from the trace.
	Autobinx Bool `json:"autobinx,omitempty"`

	// Autobiny
	// arrayOK: false
	// type: boolean
	// Obsolete: since v1.42 each bin attribute is auto-determined separately and `autobiny` is not needed. However, we accept `autobiny: true` or `false` and will update `ybins` accordingly before deleting `autobiny` from the trace.
	Autobiny Bool `json:"autobiny,omitempty"`

	// Autocolorscale
	// arrayOK: false
	// type: boolean
	// Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Bingroup
	// arrayOK: false
	// type: string
	// Set the `xbingroup` and `ybingroup` default prefix For example, setting a `bingroup` of *1* on two histogram2d traces will make them their x-bins and y-bins match separately.
	Bingroup String `json:"bingroup,omitempty"`

	// Coloraxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar
	// role: Object
	Colorbar *Histogram2dColorbar `json:"colorbar,omitempty"`

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

	// Histfunc
	// default: count
	// type: enumerated
	// Specifies the binning function used for this histogram trace. If *count*, the histogram values are computed by counting the number of values lying inside each bin. If *sum*, *avg*, *min*, *max*, the histogram values are computed using the sum, the average, the minimum or the maximum of the values lying inside each bin respectively.
	Histfunc Histogram2dHistfunc `json:"histfunc,omitempty"`

	// Histnorm
	// default:
	// type: enumerated
	// Specifies the type of normalization used for this histogram trace. If **, the span of each bar corresponds to the number of occurrences (i.e. the number of data points lying inside the bins). If *percent* / *probability*, the span of each bar corresponds to the percentage / fraction of occurrences with respect to the total number of sample points (here, the sum of all bin HEIGHTS equals 100% / 1). If *density*, the span of each bar corresponds to the number of occurrences in a bin divided by the size of the bin interval (here, the sum of all bin AREAS equals the total number of sample points). If *probability density*, the area of each bar corresponds to the probability that an event will fall into the corresponding bin (here, the sum of all bin AREAS equals 1).
	Histnorm Histogram2dHistnorm `json:"histnorm,omitempty"`

	// Hoverinfo
	// default: all
	// type: flaglist
	// Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo Histogram2dHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel
	// role: Object
	Hoverlabel *Histogram2dHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate
	// arrayOK: true
	// type: string
	// Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variable `z` Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate String `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

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

	// Marker
	// role: Object
	Marker *Histogram2dMarker `json:"marker,omitempty"`

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

	// Nbinsx
	// arrayOK: false
	// type: integer
	// Specifies the maximum number of desired bins. This value will be used in an algorithm that will decide the optimal bin size such that the histogram best visualizes the distribution of the data. Ignored if `xbins.size` is provided.
	Nbinsx int64 `json:"nbinsx,omitempty"`

	// Nbinsy
	// arrayOK: false
	// type: integer
	// Specifies the maximum number of desired bins. This value will be used in an algorithm that will decide the optimal bin size such that the histogram best visualizes the distribution of the data. Ignored if `ybins.size` is provided.
	Nbinsy int64 `json:"nbinsy,omitempty"`

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
	Stream *Histogram2dStream `json:"stream,omitempty"`

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
	Visible Histogram2dVisible `json:"visible,omitempty"`

	// X
	// arrayOK: false
	// type: data_array
	// Sets the sample data to be binned on the x axis.
	X interface{} `json:"x,omitempty"`

	// Xaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xbingroup
	// arrayOK: false
	// type: string
	// Set a group of histogram traces which will have compatible x-bin settings. Using `xbingroup`, histogram2d and histogram2dcontour traces  (on axes of the same axis type) can have compatible x-bin settings. Note that the same `xbingroup` value can be used to set (1D) histogram `bingroup`
	Xbingroup String `json:"xbingroup,omitempty"`

	// Xbins
	// role: Object
	Xbins *Histogram2dXbins `json:"xbins,omitempty"`

	// Xcalendar
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `x` date data.
	Xcalendar Histogram2dXcalendar `json:"xcalendar,omitempty"`

	// Xgap
	// arrayOK: false
	// type: number
	// Sets the horizontal gap (in pixels) between bricks.
	Xgap float64 `json:"xgap,omitempty"`

	// Xsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y
	// arrayOK: false
	// type: data_array
	// Sets the sample data to be binned on the y axis.
	Y interface{} `json:"y,omitempty"`

	// Yaxis
	// arrayOK: false
	// type: subplotid
	// Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ybingroup
	// arrayOK: false
	// type: string
	// Set a group of histogram traces which will have compatible y-bin settings. Using `ybingroup`, histogram2d and histogram2dcontour traces  (on axes of the same axis type) can have compatible y-bin settings. Note that the same `ybingroup` value can be used to set (1D) histogram `bingroup`
	Ybingroup String `json:"ybingroup,omitempty"`

	// Ybins
	// role: Object
	Ybins *Histogram2dYbins `json:"ybins,omitempty"`

	// Ycalendar
	// default: gregorian
	// type: enumerated
	// Sets the calendar system to use with `y` date data.
	Ycalendar Histogram2dYcalendar `json:"ycalendar,omitempty"`

	// Ygap
	// arrayOK: false
	// type: number
	// Sets the vertical gap (in pixels) between bricks.
	Ygap float64 `json:"ygap,omitempty"`

	// Ysrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Z
	// arrayOK: false
	// type: data_array
	// Sets the aggregation data.
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

	// Zsmooth
	// default: %!s(bool=false)
	// type: enumerated
	// Picks a smoothing algorithm use to smooth `z` data.
	Zsmooth Histogram2dZsmooth `json:"zsmooth,omitempty"`

	// Zsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

// Histogram2dColorbarTickfont Sets the color bar's tick label font
type Histogram2dColorbarTickfont struct {

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

// Histogram2dColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type Histogram2dColorbarTitleFont struct {

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

// Histogram2dColorbarTitle
type Histogram2dColorbarTitle struct {

	// Font
	// role: Object
	Font *Histogram2dColorbarTitleFont `json:"font,omitempty"`

	// Side
	// default: top
	// type: enumerated
	// Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
	Side Histogram2dColorbarTitleSide `json:"side,omitempty"`

	// Text
	// arrayOK: false
	// type: string
	// Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated.
	Text String `json:"text,omitempty"`
}

// Histogram2dColorbar
type Histogram2dColorbar struct {

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
	Exponentformat Histogram2dColorbarExponentformat `json:"exponentformat,omitempty"`

	// Len
	// arrayOK: false
	// type: number
	// Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends.
	Len float64 `json:"len,omitempty"`

	// Lenmode
	// default: fraction
	// type: enumerated
	// Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
	Lenmode Histogram2dColorbarLenmode `json:"lenmode,omitempty"`

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
	Showexponent Histogram2dColorbarShowexponent `json:"showexponent,omitempty"`

	// Showticklabels
	// arrayOK: false
	// type: boolean
	// Determines whether or not the tick labels are drawn.
	Showticklabels Bool `json:"showticklabels,omitempty"`

	// Showtickprefix
	// default: all
	// type: enumerated
	// If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
	Showtickprefix Histogram2dColorbarShowtickprefix `json:"showtickprefix,omitempty"`

	// Showticksuffix
	// default: all
	// type: enumerated
	// Same as `showtickprefix` but for tick suffixes.
	Showticksuffix Histogram2dColorbarShowticksuffix `json:"showticksuffix,omitempty"`

	// Thickness
	// arrayOK: false
	// type: number
	// Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels.
	Thickness float64 `json:"thickness,omitempty"`

	// Thicknessmode
	// default: pixels
	// type: enumerated
	// Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
	Thicknessmode Histogram2dColorbarThicknessmode `json:"thicknessmode,omitempty"`

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
	Tickfont *Histogram2dColorbarTickfont `json:"tickfont,omitempty"`

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
	Ticklabelposition Histogram2dColorbarTicklabelposition `json:"ticklabelposition,omitempty"`

	// Ticklen
	// arrayOK: false
	// type: number
	// Sets the tick length (in px).
	Ticklen float64 `json:"ticklen,omitempty"`

	// Tickmode
	// default: %!s(<nil>)
	// type: enumerated
	// Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
	Tickmode Histogram2dColorbarTickmode `json:"tickmode,omitempty"`

	// Tickprefix
	// arrayOK: false
	// type: string
	// Sets a tick label prefix.
	Tickprefix String `json:"tickprefix,omitempty"`

	// Ticks
	// default:
	// type: enumerated
	// Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
	Ticks Histogram2dColorbarTicks `json:"ticks,omitempty"`

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
	Title *Histogram2dColorbarTitle `json:"title,omitempty"`

	// X
	// arrayOK: false
	// type: number
	// Sets the x position of the color bar (in plot fraction).
	X float64 `json:"x,omitempty"`

	// Xanchor
	// default: left
	// type: enumerated
	// Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
	Xanchor Histogram2dColorbarXanchor `json:"xanchor,omitempty"`

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
	Yanchor Histogram2dColorbarYanchor `json:"yanchor,omitempty"`

	// Ypad
	// arrayOK: false
	// type: number
	// Sets the amount of padding (in px) along the y direction.
	Ypad float64 `json:"ypad,omitempty"`
}

// Histogram2dHoverlabelFont Sets the font used in hover labels.
type Histogram2dHoverlabelFont struct {

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

// Histogram2dHoverlabel
type Histogram2dHoverlabel struct {

	// Align
	// default: auto
	// type: enumerated
	// Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
	Align Histogram2dHoverlabelAlign `json:"align,omitempty"`

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
	Font *Histogram2dHoverlabelFont `json:"font,omitempty"`

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

// Histogram2dMarker
type Histogram2dMarker struct {

	// Color
	// arrayOK: false
	// type: data_array
	// Sets the aggregation data.
	Color interface{} `json:"color,omitempty"`

	// Colorsrc
	// arrayOK: false
	// type: string
	// Sets the source reference on Chart Studio Cloud for  color .
	Colorsrc String `json:"colorsrc,omitempty"`
}

// Histogram2dStream
type Histogram2dStream struct {

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

// Histogram2dXbins
type Histogram2dXbins struct {

	// End
	// arrayOK: false
	// type: any
	// Sets the end value for the x axis bins. The last bin may not end exactly at this value, we increment the bin edge by `size` from `start` until we reach or exceed `end`. Defaults to the maximum data value. Like `start`, for dates use a date string, and for category data `end` is based on the category serial numbers.
	End interface{} `json:"end,omitempty"`

	// Size
	// arrayOK: false
	// type: any
	// Sets the size of each x axis bin. Default behavior: If `nbinsx` is 0 or omitted, we choose a nice round bin size such that the number of bins is about the same as the typical number of samples in each bin. If `nbinsx` is provided, we choose a nice round bin size giving no more than that many bins. For date data, use milliseconds or *M<n>* for months, as in `axis.dtick`. For category data, the number of categories to bin together (always defaults to 1).
	Size interface{} `json:"size,omitempty"`

	// Start
	// arrayOK: false
	// type: any
	// Sets the starting value for the x axis bins. Defaults to the minimum data value, shifted down if necessary to make nice round values and to remove ambiguous bin edges. For example, if most of the data is integers we shift the bin edges 0.5 down, so a `size` of 5 would have a default `start` of -0.5, so it is clear that 0-4 are in the first bin, 5-9 in the second, but continuous data gets a start of 0 and bins [0,5), [5,10) etc. Dates behave similarly, and `start` should be a date string. For category data, `start` is based on the category serial numbers, and defaults to -0.5.
	Start interface{} `json:"start,omitempty"`
}

// Histogram2dYbins
type Histogram2dYbins struct {

	// End
	// arrayOK: false
	// type: any
	// Sets the end value for the y axis bins. The last bin may not end exactly at this value, we increment the bin edge by `size` from `start` until we reach or exceed `end`. Defaults to the maximum data value. Like `start`, for dates use a date string, and for category data `end` is based on the category serial numbers.
	End interface{} `json:"end,omitempty"`

	// Size
	// arrayOK: false
	// type: any
	// Sets the size of each y axis bin. Default behavior: If `nbinsy` is 0 or omitted, we choose a nice round bin size such that the number of bins is about the same as the typical number of samples in each bin. If `nbinsy` is provided, we choose a nice round bin size giving no more than that many bins. For date data, use milliseconds or *M<n>* for months, as in `axis.dtick`. For category data, the number of categories to bin together (always defaults to 1).
	Size interface{} `json:"size,omitempty"`

	// Start
	// arrayOK: false
	// type: any
	// Sets the starting value for the y axis bins. Defaults to the minimum data value, shifted down if necessary to make nice round values and to remove ambiguous bin edges. For example, if most of the data is integers we shift the bin edges 0.5 down, so a `size` of 5 would have a default `start` of -0.5, so it is clear that 0-4 are in the first bin, 5-9 in the second, but continuous data gets a start of 0 and bins [0,5), [5,10) etc. Dates behave similarly, and `start` should be a date string. For category data, `start` is based on the category serial numbers, and defaults to -0.5.
	Start interface{} `json:"start,omitempty"`
}

// Histogram2dColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type Histogram2dColorbarExponentformat string

const (
	Histogram2dColorbarExponentformatNone  Histogram2dColorbarExponentformat = "none"
	Histogram2dColorbarExponentformatE1    Histogram2dColorbarExponentformat = "e"
	Histogram2dColorbarExponentformatE2    Histogram2dColorbarExponentformat = "E"
	Histogram2dColorbarExponentformatPower Histogram2dColorbarExponentformat = "power"
	Histogram2dColorbarExponentformatSi    Histogram2dColorbarExponentformat = "SI"
	Histogram2dColorbarExponentformatB     Histogram2dColorbarExponentformat = "B"
)

// Histogram2dColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type Histogram2dColorbarLenmode string

const (
	Histogram2dColorbarLenmodeFraction Histogram2dColorbarLenmode = "fraction"
	Histogram2dColorbarLenmodePixels   Histogram2dColorbarLenmode = "pixels"
)

// Histogram2dColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type Histogram2dColorbarShowexponent string

const (
	Histogram2dColorbarShowexponentAll   Histogram2dColorbarShowexponent = "all"
	Histogram2dColorbarShowexponentFirst Histogram2dColorbarShowexponent = "first"
	Histogram2dColorbarShowexponentLast  Histogram2dColorbarShowexponent = "last"
	Histogram2dColorbarShowexponentNone  Histogram2dColorbarShowexponent = "none"
)

// Histogram2dColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type Histogram2dColorbarShowtickprefix string

const (
	Histogram2dColorbarShowtickprefixAll   Histogram2dColorbarShowtickprefix = "all"
	Histogram2dColorbarShowtickprefixFirst Histogram2dColorbarShowtickprefix = "first"
	Histogram2dColorbarShowtickprefixLast  Histogram2dColorbarShowtickprefix = "last"
	Histogram2dColorbarShowtickprefixNone  Histogram2dColorbarShowtickprefix = "none"
)

// Histogram2dColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type Histogram2dColorbarShowticksuffix string

const (
	Histogram2dColorbarShowticksuffixAll   Histogram2dColorbarShowticksuffix = "all"
	Histogram2dColorbarShowticksuffixFirst Histogram2dColorbarShowticksuffix = "first"
	Histogram2dColorbarShowticksuffixLast  Histogram2dColorbarShowticksuffix = "last"
	Histogram2dColorbarShowticksuffixNone  Histogram2dColorbarShowticksuffix = "none"
)

// Histogram2dColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type Histogram2dColorbarThicknessmode string

const (
	Histogram2dColorbarThicknessmodeFraction Histogram2dColorbarThicknessmode = "fraction"
	Histogram2dColorbarThicknessmodePixels   Histogram2dColorbarThicknessmode = "pixels"
)

// Histogram2dColorbarTicklabelposition Determines where tick labels are drawn.
type Histogram2dColorbarTicklabelposition string

const (
	Histogram2dColorbarTicklabelpositionOutside       Histogram2dColorbarTicklabelposition = "outside"
	Histogram2dColorbarTicklabelpositionInside        Histogram2dColorbarTicklabelposition = "inside"
	Histogram2dColorbarTicklabelpositionOutsideTop    Histogram2dColorbarTicklabelposition = "outside top"
	Histogram2dColorbarTicklabelpositionInsideTop     Histogram2dColorbarTicklabelposition = "inside top"
	Histogram2dColorbarTicklabelpositionOutsideBottom Histogram2dColorbarTicklabelposition = "outside bottom"
	Histogram2dColorbarTicklabelpositionInsideBottom  Histogram2dColorbarTicklabelposition = "inside bottom"
)

// Histogram2dColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type Histogram2dColorbarTickmode string

const (
	Histogram2dColorbarTickmodeAuto   Histogram2dColorbarTickmode = "auto"
	Histogram2dColorbarTickmodeLinear Histogram2dColorbarTickmode = "linear"
	Histogram2dColorbarTickmodeArray  Histogram2dColorbarTickmode = "array"
)

// Histogram2dColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type Histogram2dColorbarTicks string

const (
	Histogram2dColorbarTicksOutside Histogram2dColorbarTicks = "outside"
	Histogram2dColorbarTicksInside  Histogram2dColorbarTicks = "inside"
	Histogram2dColorbarTicksEmpty   Histogram2dColorbarTicks = ""
)

// Histogram2dColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type Histogram2dColorbarTitleSide string

const (
	Histogram2dColorbarTitleSideRight  Histogram2dColorbarTitleSide = "right"
	Histogram2dColorbarTitleSideTop    Histogram2dColorbarTitleSide = "top"
	Histogram2dColorbarTitleSideBottom Histogram2dColorbarTitleSide = "bottom"
)

// Histogram2dColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type Histogram2dColorbarXanchor string

const (
	Histogram2dColorbarXanchorLeft   Histogram2dColorbarXanchor = "left"
	Histogram2dColorbarXanchorCenter Histogram2dColorbarXanchor = "center"
	Histogram2dColorbarXanchorRight  Histogram2dColorbarXanchor = "right"
)

// Histogram2dColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type Histogram2dColorbarYanchor string

const (
	Histogram2dColorbarYanchorTop    Histogram2dColorbarYanchor = "top"
	Histogram2dColorbarYanchorMiddle Histogram2dColorbarYanchor = "middle"
	Histogram2dColorbarYanchorBottom Histogram2dColorbarYanchor = "bottom"
)

// Histogram2dHistfunc Specifies the binning function used for this histogram trace. If *count*, the histogram values are computed by counting the number of values lying inside each bin. If *sum*, *avg*, *min*, *max*, the histogram values are computed using the sum, the average, the minimum or the maximum of the values lying inside each bin respectively.
type Histogram2dHistfunc string

const (
	Histogram2dHistfuncCount Histogram2dHistfunc = "count"
	Histogram2dHistfuncSum   Histogram2dHistfunc = "sum"
	Histogram2dHistfuncAvg   Histogram2dHistfunc = "avg"
	Histogram2dHistfuncMin   Histogram2dHistfunc = "min"
	Histogram2dHistfuncMax   Histogram2dHistfunc = "max"
)

// Histogram2dHistnorm Specifies the type of normalization used for this histogram trace. If **, the span of each bar corresponds to the number of occurrences (i.e. the number of data points lying inside the bins). If *percent* / *probability*, the span of each bar corresponds to the percentage / fraction of occurrences with respect to the total number of sample points (here, the sum of all bin HEIGHTS equals 100% / 1). If *density*, the span of each bar corresponds to the number of occurrences in a bin divided by the size of the bin interval (here, the sum of all bin AREAS equals the total number of sample points). If *probability density*, the area of each bar corresponds to the probability that an event will fall into the corresponding bin (here, the sum of all bin AREAS equals 1).
type Histogram2dHistnorm string

const (
	Histogram2dHistnormEmpty              Histogram2dHistnorm = ""
	Histogram2dHistnormPercent            Histogram2dHistnorm = "percent"
	Histogram2dHistnormProbability        Histogram2dHistnorm = "probability"
	Histogram2dHistnormDensity            Histogram2dHistnorm = "density"
	Histogram2dHistnormProbabilityDensity Histogram2dHistnorm = "probability density"
)

// Histogram2dHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type Histogram2dHoverlabelAlign string

const (
	Histogram2dHoverlabelAlignLeft  Histogram2dHoverlabelAlign = "left"
	Histogram2dHoverlabelAlignRight Histogram2dHoverlabelAlign = "right"
	Histogram2dHoverlabelAlignAuto  Histogram2dHoverlabelAlign = "auto"
)

// Histogram2dVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type Histogram2dVisible interface{}

var (
	Histogram2dVisibleTrue       Histogram2dVisible = true
	Histogram2dVisibleFalse      Histogram2dVisible = false
	Histogram2dVisibleLegendonly Histogram2dVisible = "legendonly"
)

// Histogram2dXcalendar Sets the calendar system to use with `x` date data.
type Histogram2dXcalendar string

const (
	Histogram2dXcalendarGregorian  Histogram2dXcalendar = "gregorian"
	Histogram2dXcalendarChinese    Histogram2dXcalendar = "chinese"
	Histogram2dXcalendarCoptic     Histogram2dXcalendar = "coptic"
	Histogram2dXcalendarDiscworld  Histogram2dXcalendar = "discworld"
	Histogram2dXcalendarEthiopian  Histogram2dXcalendar = "ethiopian"
	Histogram2dXcalendarHebrew     Histogram2dXcalendar = "hebrew"
	Histogram2dXcalendarIslamic    Histogram2dXcalendar = "islamic"
	Histogram2dXcalendarJulian     Histogram2dXcalendar = "julian"
	Histogram2dXcalendarMayan      Histogram2dXcalendar = "mayan"
	Histogram2dXcalendarNanakshahi Histogram2dXcalendar = "nanakshahi"
	Histogram2dXcalendarNepali     Histogram2dXcalendar = "nepali"
	Histogram2dXcalendarPersian    Histogram2dXcalendar = "persian"
	Histogram2dXcalendarJalali     Histogram2dXcalendar = "jalali"
	Histogram2dXcalendarTaiwan     Histogram2dXcalendar = "taiwan"
	Histogram2dXcalendarThai       Histogram2dXcalendar = "thai"
	Histogram2dXcalendarUmmalqura  Histogram2dXcalendar = "ummalqura"
)

// Histogram2dYcalendar Sets the calendar system to use with `y` date data.
type Histogram2dYcalendar string

const (
	Histogram2dYcalendarGregorian  Histogram2dYcalendar = "gregorian"
	Histogram2dYcalendarChinese    Histogram2dYcalendar = "chinese"
	Histogram2dYcalendarCoptic     Histogram2dYcalendar = "coptic"
	Histogram2dYcalendarDiscworld  Histogram2dYcalendar = "discworld"
	Histogram2dYcalendarEthiopian  Histogram2dYcalendar = "ethiopian"
	Histogram2dYcalendarHebrew     Histogram2dYcalendar = "hebrew"
	Histogram2dYcalendarIslamic    Histogram2dYcalendar = "islamic"
	Histogram2dYcalendarJulian     Histogram2dYcalendar = "julian"
	Histogram2dYcalendarMayan      Histogram2dYcalendar = "mayan"
	Histogram2dYcalendarNanakshahi Histogram2dYcalendar = "nanakshahi"
	Histogram2dYcalendarNepali     Histogram2dYcalendar = "nepali"
	Histogram2dYcalendarPersian    Histogram2dYcalendar = "persian"
	Histogram2dYcalendarJalali     Histogram2dYcalendar = "jalali"
	Histogram2dYcalendarTaiwan     Histogram2dYcalendar = "taiwan"
	Histogram2dYcalendarThai       Histogram2dYcalendar = "thai"
	Histogram2dYcalendarUmmalqura  Histogram2dYcalendar = "ummalqura"
)

// Histogram2dZsmooth Picks a smoothing algorithm use to smooth `z` data.
type Histogram2dZsmooth interface{}

var (
	Histogram2dZsmoothFast  Histogram2dZsmooth = "fast"
	Histogram2dZsmoothBest  Histogram2dZsmooth = "best"
	Histogram2dZsmoothFalse Histogram2dZsmooth = false
)

// Histogram2dHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type Histogram2dHoverinfo string

const (
	// Flags
	Histogram2dHoverinfoX    Histogram2dHoverinfo = "x"
	Histogram2dHoverinfoY    Histogram2dHoverinfo = "y"
	Histogram2dHoverinfoZ    Histogram2dHoverinfo = "z"
	Histogram2dHoverinfoText Histogram2dHoverinfo = "text"
	Histogram2dHoverinfoName Histogram2dHoverinfo = "name"

	// Extra
	Histogram2dHoverinfoAll  Histogram2dHoverinfo = "all"
	Histogram2dHoverinfoNone Histogram2dHoverinfo = "none"
	Histogram2dHoverinfoSkip Histogram2dHoverinfo = "skip"
)
