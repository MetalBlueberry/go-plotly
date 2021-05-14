package grob

var TraceTypeOhlc TraceType = "ohlc"

func (trace *Ohlc) GetType() TraceType {
	return TraceTypeOhlc
}
// Ohlc The ohlc (short for Open-High-Low-Close) is a style of financial chart describing open, high, low and close for a given `x` coordinate (most likely time). The tip of the lines represent the `low` and `high` values and the horizontal segments represent the `open` and `close` values. Sample points where the close value is higher (lower) then the open value are called increasing (decreasing). By default, increasing items are drawn in green whereas decreasing are drawn in red.
type Ohlc struct {
    
    // Type 
    // is the type of the plot 
    Type TraceType `json:"type,omitempty"`
    
    // Close 
    // data_array 
    // Sets the close values. 
    Close interface{} `json:"close,omitempty"`
    
    // Closesrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  close . 
    Closesrc String `json:"closesrc,omitempty"`
    
    // Customdata 
    // data_array 
    // Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements 
    Customdata interface{} `json:"customdata,omitempty"`
    
    // Customdatasrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  customdata . 
    Customdatasrc String `json:"customdatasrc,omitempty"`
    
    // Decreasing 
    //   
    Decreasing *OhlcDecreasing `json:"decreasing,omitempty"`
    
    // High 
    // data_array 
    // Sets the high values. 
    High interface{} `json:"high,omitempty"`
    
    // Highsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  high . 
    Highsrc String `json:"highsrc,omitempty"`
    
    // Hoverinfo 
    // flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired. 
    Hoverinfo OhlcHoverinfo `json:"hoverinfo,omitempty"`
    
    // Hoverinfosrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hoverinfo . 
    Hoverinfosrc String `json:"hoverinfosrc,omitempty"`
    
    // Hoverlabel 
    //   
    Hoverlabel *OhlcHoverlabel `json:"hoverlabel,omitempty"`
    
    // Hovertext 
    // string 
    // Same as `text`. 
    Hovertext String `json:"hovertext,omitempty"`
    
    // Hovertextsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hovertext . 
    Hovertextsrc String `json:"hovertextsrc,omitempty"`
    
    // Ids 
    // data_array 
    // Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type. 
    Ids interface{} `json:"ids,omitempty"`
    
    // Idssrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  ids . 
    Idssrc String `json:"idssrc,omitempty"`
    
    // Increasing 
    //   
    Increasing *OhlcIncreasing `json:"increasing,omitempty"`
    
    // Legendgroup 
    // string 
    // Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items. 
    Legendgroup String `json:"legendgroup,omitempty"`
    
    // Line 
    //   
    Line *OhlcLine `json:"line,omitempty"`
    
    // Low 
    // data_array 
    // Sets the low values. 
    Low interface{} `json:"low,omitempty"`
    
    // Lowsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  low . 
    Lowsrc String `json:"lowsrc,omitempty"`
    
    // Meta 
    // any 
    // Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index. 
    Meta interface{} `json:"meta,omitempty"`
    
    // Metasrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  meta . 
    Metasrc String `json:"metasrc,omitempty"`
    
    // Name 
    // string 
    // Sets the trace name. The trace name appear as the legend item and on hover. 
    Name String `json:"name,omitempty"`
    
    // Opacity 
    // number 
    // Sets the opacity of the trace. 
    Opacity float64 `json:"opacity,omitempty"`
    
    // Open 
    // data_array 
    // Sets the open values. 
    Open interface{} `json:"open,omitempty"`
    
    // Opensrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  open . 
    Opensrc String `json:"opensrc,omitempty"`
    
    // Selectedpoints 
    // any 
    // Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect. 
    Selectedpoints interface{} `json:"selectedpoints,omitempty"`
    
    // Showlegend 
    // boolean 
    // Determines whether or not an item corresponding to this trace is shown in the legend. 
    Showlegend Bool `json:"showlegend,omitempty"`
    
    // Stream 
    //   
    Stream *OhlcStream `json:"stream,omitempty"`
    
    // Text 
    // string 
    // Sets hover text elements associated with each sample point. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to this trace's sample points. 
    Text String `json:"text,omitempty"`
    
    // Textsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  text . 
    Textsrc String `json:"textsrc,omitempty"`
    
    // Tickwidth 
    // number 
    // Sets the width of the open/close tick marks relative to the *x* minimal interval. 
    Tickwidth float64 `json:"tickwidth,omitempty"`
    
    // Transforms 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Transforms interface{} `json:"transforms,omitempty"`
    
    // Uid 
    // string 
    // Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions. 
    Uid String `json:"uid,omitempty"`
    
    // Uirevision 
    // any 
    // Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves. 
    Uirevision interface{} `json:"uirevision,omitempty"`
    
    // Visible 
    // enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible). 
    Visible OhlcVisible `json:"visible,omitempty"`
    
    // X 
    // data_array 
    // Sets the x coordinates. If absent, linear coordinate will be generated. 
    X interface{} `json:"x,omitempty"`
    
    // Xaxis 
    // subplotid 
    // Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on. 
    Xaxis String `json:"xaxis,omitempty"`
    
    // Xcalendar 
    // enumerated Sets the calendar system to use with `x` date data. 
    Xcalendar OhlcXcalendar `json:"xcalendar,omitempty"`
    
    // Xperiod 
    // any 
    // Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the x axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer. 
    Xperiod interface{} `json:"xperiod,omitempty"`
    
    // Xperiod0 
    // any 
    // Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the x0 axis. When `x0period` is round number of weeks, the `x0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01. 
    Xperiod0 interface{} `json:"xperiod0,omitempty"`
    
    // Xperiodalignment 
    // enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis. 
    Xperiodalignment OhlcXperiodalignment `json:"xperiodalignment,omitempty"`
    
    // Xsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  x . 
    Xsrc String `json:"xsrc,omitempty"`
    
    // Yaxis 
    // subplotid 
    // Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on. 
    Yaxis String `json:"yaxis,omitempty"`
    
}
// OhlcDecreasingLine 
type OhlcDecreasingLine struct {
    
    // Color 
    // color 
    // Sets the line color. 
    Color Color `json:"color,omitempty"`
    
    // Dash 
    // string 
    // Sets the dash style of lines. Set to a dash type string (*solid*, *dot*, *dash*, *longdash*, *dashdot*, or *longdashdot*) or a dash length list in px (eg *5px,10px,2px,2px*). 
    Dash String `json:"dash,omitempty"`
    
    // Width 
    // number 
    // Sets the line width (in px). 
    Width float64 `json:"width,omitempty"`
    
}
// OhlcDecreasing 
type OhlcDecreasing struct {
    
    // Line 
    //   
    Line *OhlcDecreasingLine `json:"line,omitempty"`
    
}
// OhlcHoverlabelFont Sets the font used in hover labels.
type OhlcHoverlabelFont struct {
    
    // Color 
    // color 
    //  
    Color Color `json:"color,omitempty"`
    
    // Colorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  color . 
    Colorsrc String `json:"colorsrc,omitempty"`
    
    // Family 
    // string 
    // HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*. 
    Family String `json:"family,omitempty"`
    
    // Familysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  family . 
    Familysrc String `json:"familysrc,omitempty"`
    
    // Size 
    // number 
    //  
    Size float64 `json:"size,omitempty"`
    
    // Sizesrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  size . 
    Sizesrc String `json:"sizesrc,omitempty"`
    
}
// OhlcHoverlabel 
type OhlcHoverlabel struct {
    
    // Align 
    // enumerated Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines 
    Align OhlcHoverlabelAlign `json:"align,omitempty"`
    
    // Alignsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  align . 
    Alignsrc String `json:"alignsrc,omitempty"`
    
    // Bgcolor 
    // color 
    // Sets the background color of the hover labels for this trace 
    Bgcolor Color `json:"bgcolor,omitempty"`
    
    // Bgcolorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  bgcolor . 
    Bgcolorsrc String `json:"bgcolorsrc,omitempty"`
    
    // Bordercolor 
    // color 
    // Sets the border color of the hover labels for this trace. 
    Bordercolor Color `json:"bordercolor,omitempty"`
    
    // Bordercolorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  bordercolor . 
    Bordercolorsrc String `json:"bordercolorsrc,omitempty"`
    
    // Font 
    //  Sets the font used in hover labels. 
    Font *OhlcHoverlabelFont `json:"font,omitempty"`
    
    // Namelength 
    // integer 
    // Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis. 
    Namelength int64 `json:"namelength,omitempty"`
    
    // Namelengthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  namelength . 
    Namelengthsrc String `json:"namelengthsrc,omitempty"`
    
    // Split 
    // boolean 
    // Show hover information (open, close, high, low) in separate labels. 
    Split Bool `json:"split,omitempty"`
    
}
// OhlcIncreasingLine 
type OhlcIncreasingLine struct {
    
    // Color 
    // color 
    // Sets the line color. 
    Color Color `json:"color,omitempty"`
    
    // Dash 
    // string 
    // Sets the dash style of lines. Set to a dash type string (*solid*, *dot*, *dash*, *longdash*, *dashdot*, or *longdashdot*) or a dash length list in px (eg *5px,10px,2px,2px*). 
    Dash String `json:"dash,omitempty"`
    
    // Width 
    // number 
    // Sets the line width (in px). 
    Width float64 `json:"width,omitempty"`
    
}
// OhlcIncreasing 
type OhlcIncreasing struct {
    
    // Line 
    //   
    Line *OhlcIncreasingLine `json:"line,omitempty"`
    
}
// OhlcLine 
type OhlcLine struct {
    
    // Dash 
    // string 
    // Sets the dash style of lines. Set to a dash type string (*solid*, *dot*, *dash*, *longdash*, *dashdot*, or *longdashdot*) or a dash length list in px (eg *5px,10px,2px,2px*). Note that this style setting can also be set per direction via `increasing.line.dash` and `decreasing.line.dash`. 
    Dash String `json:"dash,omitempty"`
    
    // Width 
    // number 
    // [object Object] Note that this style setting can also be set per direction via `increasing.line.width` and `decreasing.line.width`. 
    Width float64 `json:"width,omitempty"`
    
}
// OhlcStream 
type OhlcStream struct {
    
    // Maxpoints 
    // number 
    // Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot. 
    Maxpoints float64 `json:"maxpoints,omitempty"`
    
    // Token 
    // string 
    // The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details. 
    Token String `json:"token,omitempty"`
    
}
// OhlcHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type OhlcHoverlabelAlign string 

const (
    OhlcHoverlabelAlignLeft OhlcHoverlabelAlign = "left"
    OhlcHoverlabelAlignRight OhlcHoverlabelAlign = "right"
    OhlcHoverlabelAlignAuto OhlcHoverlabelAlign = "auto"
    
)
// OhlcVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type OhlcVisible interface{} 

var (
    OhlcVisibleTrue OhlcVisible = true
    OhlcVisibleFalse OhlcVisible = false
    OhlcVisibleLegendonly OhlcVisible = "legendonly"
    
)
// OhlcXcalendar Sets the calendar system to use with `x` date data.
type OhlcXcalendar string 

const (
    OhlcXcalendarGregorian OhlcXcalendar = "gregorian"
    OhlcXcalendarChinese OhlcXcalendar = "chinese"
    OhlcXcalendarCoptic OhlcXcalendar = "coptic"
    OhlcXcalendarDiscworld OhlcXcalendar = "discworld"
    OhlcXcalendarEthiopian OhlcXcalendar = "ethiopian"
    OhlcXcalendarHebrew OhlcXcalendar = "hebrew"
    OhlcXcalendarIslamic OhlcXcalendar = "islamic"
    OhlcXcalendarJulian OhlcXcalendar = "julian"
    OhlcXcalendarMayan OhlcXcalendar = "mayan"
    OhlcXcalendarNanakshahi OhlcXcalendar = "nanakshahi"
    OhlcXcalendarNepali OhlcXcalendar = "nepali"
    OhlcXcalendarPersian OhlcXcalendar = "persian"
    OhlcXcalendarJalali OhlcXcalendar = "jalali"
    OhlcXcalendarTaiwan OhlcXcalendar = "taiwan"
    OhlcXcalendarThai OhlcXcalendar = "thai"
    OhlcXcalendarUmmalqura OhlcXcalendar = "ummalqura"
    
)
// OhlcXperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
type OhlcXperiodalignment string 

const (
    OhlcXperiodalignmentStart OhlcXperiodalignment = "start"
    OhlcXperiodalignmentMiddle OhlcXperiodalignment = "middle"
    OhlcXperiodalignmentEnd OhlcXperiodalignment = "end"
    
)
// OhlcHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type OhlcHoverinfo string 

const (
    // Flags
    OhlcHoverinfoX OhlcHoverinfo = "x"
    OhlcHoverinfoY OhlcHoverinfo = "y"
    OhlcHoverinfoZ OhlcHoverinfo = "z"
    OhlcHoverinfoText OhlcHoverinfo = "text"
    OhlcHoverinfoName OhlcHoverinfo = "name"
    
    // Extra
    OhlcHoverinfoAll OhlcHoverinfo = "all"
    OhlcHoverinfoNone OhlcHoverinfo = "none"
    OhlcHoverinfoSkip OhlcHoverinfo = "skip"
    
)
