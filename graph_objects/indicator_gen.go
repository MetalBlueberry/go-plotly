package grob

var TraceTypeIndicator TraceType = "indicator"

func (trace *Indicator) GetType() TraceType {
	return TraceTypeIndicator
}
// Indicator An indicator is used to visualize a single `value` along with some contextual information such as `steps` or a `threshold`, using a combination of three visual elements: a number, a delta, and/or a gauge. Deltas are taken with respect to a `reference`. Gauges can be either angular or bullet (aka linear) gauges.
type Indicator struct {
    
    // Type 
    // is the type of the plot 
    Type TraceType `json:"type,omitempty"`
    
    // Align 
    // enumerated Sets the horizontal alignment of the `text` within the box. Note that this attribute has no effect if an angular gauge is displayed: in this case, it is always centered 
    Align IndicatorAlign `json:"align,omitempty"`
    
    // Customdata 
    // data_array 
    // Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements 
    Customdata interface{} `json:"customdata,omitempty"`
    
    // Customdatasrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  customdata . 
    Customdatasrc String `json:"customdatasrc,omitempty"`
    
    // Delta 
    //   
    Delta *IndicatorDelta `json:"delta,omitempty"`
    
    // Domain 
    //   
    Domain *IndicatorDomain `json:"domain,omitempty"`
    
    // Gauge 
    //  The gauge of the Indicator plot. 
    Gauge *IndicatorGauge `json:"gauge,omitempty"`
    
    // Ids 
    // data_array 
    // Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type. 
    Ids interface{} `json:"ids,omitempty"`
    
    // Idssrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  ids . 
    Idssrc String `json:"idssrc,omitempty"`
    
    // Meta 
    // any 
    // Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index. 
    Meta interface{} `json:"meta,omitempty"`
    
    // Metasrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  meta . 
    Metasrc String `json:"metasrc,omitempty"`
    
    // Mode 
    // flaglist Determines how the value is displayed on the graph. `number` displays the value numerically in text. `delta` displays the difference to a reference value in text. Finally, `gauge` displays the value graphically on an axis. 
    Mode IndicatorMode `json:"mode,omitempty"`
    
    // Name 
    // string 
    // Sets the trace name. The trace name appear as the legend item and on hover. 
    Name String `json:"name,omitempty"`
    
    // Number 
    //   
    Number *IndicatorNumber `json:"number,omitempty"`
    
    // Stream 
    //   
    Stream *IndicatorStream `json:"stream,omitempty"`
    
    // Title 
    //   
    Title *IndicatorTitle `json:"title,omitempty"`
    
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
    
    // Value 
    // number 
    // Sets the number to be displayed. 
    Value float64 `json:"value,omitempty"`
    
    // Visible 
    // enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible). 
    Visible IndicatorVisible `json:"visible,omitempty"`
    
}
// IndicatorDeltaDecreasing 
type IndicatorDeltaDecreasing struct {
    
    // Color 
    // color 
    // Sets the color for increasing value. 
    Color Color `json:"color,omitempty"`
    
    // Symbol 
    // string 
    // Sets the symbol to display for increasing value 
    Symbol String `json:"symbol,omitempty"`
    
}
// IndicatorDeltaFont Set the font used to display the delta
type IndicatorDeltaFont struct {
    
    // Color 
    // color 
    //  
    Color Color `json:"color,omitempty"`
    
    // Family 
    // string 
    // HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*. 
    Family String `json:"family,omitempty"`
    
    // Size 
    // number 
    //  
    Size float64 `json:"size,omitempty"`
    
}
// IndicatorDeltaIncreasing 
type IndicatorDeltaIncreasing struct {
    
    // Color 
    // color 
    // Sets the color for increasing value. 
    Color Color `json:"color,omitempty"`
    
    // Symbol 
    // string 
    // Sets the symbol to display for increasing value 
    Symbol String `json:"symbol,omitempty"`
    
}
// IndicatorDelta 
type IndicatorDelta struct {
    
    // Decreasing 
    //   
    Decreasing *IndicatorDeltaDecreasing `json:"decreasing,omitempty"`
    
    // Font 
    //  Set the font used to display the delta 
    Font *IndicatorDeltaFont `json:"font,omitempty"`
    
    // Increasing 
    //   
    Increasing *IndicatorDeltaIncreasing `json:"increasing,omitempty"`
    
    // Position 
    // enumerated Sets the position of delta with respect to the number. 
    Position IndicatorDeltaPosition `json:"position,omitempty"`
    
    // Reference 
    // number 
    // Sets the reference value to compute the delta. By default, it is set to the current value. 
    Reference float64 `json:"reference,omitempty"`
    
    // Relative 
    // boolean 
    // Show relative change 
    Relative Bool `json:"relative,omitempty"`
    
    // Valueformat 
    // string 
    // Sets the value formatting rule using d3 formatting mini-language which is similar to those of Python. See https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format 
    Valueformat String `json:"valueformat,omitempty"`
    
}
// IndicatorDomain 
type IndicatorDomain struct {
    
    // Column 
    // integer 
    // If there is a layout grid, use the domain for this column in the grid for this indicator trace . 
    Column int64 `json:"column,omitempty"`
    
    // Row 
    // integer 
    // If there is a layout grid, use the domain for this row in the grid for this indicator trace . 
    Row int64 `json:"row,omitempty"`
    
    // X 
    // info_array 
    // Sets the horizontal domain of this indicator trace (in plot fraction). 
    X interface{} `json:"x,omitempty"`
    
    // Y 
    // info_array 
    // Sets the vertical domain of this indicator trace (in plot fraction). 
    Y interface{} `json:"y,omitempty"`
    
}
// IndicatorGaugeAxisTickfont Sets the color bar's tick label font
type IndicatorGaugeAxisTickfont struct {
    
    // Color 
    // color 
    //  
    Color Color `json:"color,omitempty"`
    
    // Family 
    // string 
    // HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*. 
    Family String `json:"family,omitempty"`
    
    // Size 
    // number 
    //  
    Size float64 `json:"size,omitempty"`
    
}
// IndicatorGaugeAxis 
type IndicatorGaugeAxis struct {
    
    // Dtick 
    // any 
    // Sets the step in-between ticks on this axis. Use with `tick0`. Must be a positive number, or special strings available to *log* and *date* axes. If the axis `type` is *log*, then ticks are set every 10^(n*dtick) where n is the tick number. For example, to set a tick mark at 1, 10, 100, 1000, ... set dtick to 1. To set tick marks at 1, 100, 10000, ... set dtick to 2. To set tick marks at 1, 5, 25, 125, 625, 3125, ... set dtick to log_10(5), or 0.69897000433. *log* has several special values; *L<f>*, where `f` is a positive number, gives ticks linearly spaced in value (but not position). For example `tick0` = 0.1, `dtick` = *L0.5* will put ticks at 0.1, 0.6, 1.1, 1.6 etc. To show powers of 10 plus small digits between, use *D1* (all digits) or *D2* (only 2 and 5). `tick0` is ignored for *D1* and *D2*. If the axis `type` is *date*, then you must convert the time to milliseconds. For example, to set the interval between ticks to one day, set `dtick` to 86400000.0. *date* also has special values *M<n>* gives ticks spaced by a number of months. `n` must be a positive integer. To set ticks on the 15th of every third month, set `tick0` to *2000-01-15* and `dtick` to *M3*. To set ticks every 4 years, set `dtick` to *M48* 
    Dtick interface{} `json:"dtick,omitempty"`
    
    // Exponentformat 
    // enumerated Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B. 
    Exponentformat IndicatorGaugeAxisExponentformat `json:"exponentformat,omitempty"`
    
    // Minexponent 
    // number 
    // Hide SI prefix for 10^n if |n| is below this number. This only has an effect when `tickformat` is *SI* or *B*. 
    Minexponent float64 `json:"minexponent,omitempty"`
    
    // Nticks 
    // integer 
    // Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*. 
    Nticks int64 `json:"nticks,omitempty"`
    
    // Range 
    // info_array 
    // Sets the range of this axis. 
    Range interface{} `json:"range,omitempty"`
    
    // Separatethousands 
    // boolean 
    // If "true", even 4-digit integers are separated 
    Separatethousands Bool `json:"separatethousands,omitempty"`
    
    // Showexponent 
    // enumerated If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear. 
    Showexponent IndicatorGaugeAxisShowexponent `json:"showexponent,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix IndicatorGaugeAxisShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix IndicatorGaugeAxisShowticksuffix `json:"showticksuffix,omitempty"`
    
    // Tick0 
    // any 
    // Sets the placement of the first tick on this axis. Use with `dtick`. If the axis `type` is *log*, then you must take the log of your starting tick (e.g. to set the starting tick to 100, set the `tick0` to 2) except when `dtick`=*L<f>* (see `dtick` for more info). If the axis `type` is *date*, it should be a date string, like date data. If the axis `type` is *category*, it should be a number, using the scale where each category is assigned a serial number from zero in the order it appears. 
    Tick0 interface{} `json:"tick0,omitempty"`
    
    // Tickangle 
    // angle 
    // Sets the angle of the tick labels with respect to the horizontal. For example, a `tickangle` of -90 draws the tick labels vertically. 
    Tickangle float64 `json:"tickangle,omitempty"`
    
    // Tickcolor 
    // color 
    // Sets the tick color. 
    Tickcolor Color `json:"tickcolor,omitempty"`
    
    // Tickfont 
    //  Sets the color bar's tick label font 
    Tickfont *IndicatorGaugeAxisTickfont `json:"tickfont,omitempty"`
    
    // Tickformat 
    // string 
    // Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Tickformat String `json:"tickformat,omitempty"`
    
    // Tickformatstops 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Tickformatstops interface{} `json:"tickformatstops,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode IndicatorGaugeAxisTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks IndicatorGaugeAxisTicks `json:"ticks,omitempty"`
    
    // Ticksuffix 
    // string 
    // Sets a tick label suffix. 
    Ticksuffix String `json:"ticksuffix,omitempty"`
    
    // Ticktext 
    // data_array 
    // Sets the text displayed at the ticks position via `tickvals`. Only has an effect if `tickmode` is set to *array*. Used with `tickvals`. 
    Ticktext interface{} `json:"ticktext,omitempty"`
    
    // Ticktextsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  ticktext . 
    Ticktextsrc String `json:"ticktextsrc,omitempty"`
    
    // Tickvals 
    // data_array 
    // Sets the values at which ticks on this axis appear. Only has an effect if `tickmode` is set to *array*. Used with `ticktext`. 
    Tickvals interface{} `json:"tickvals,omitempty"`
    
    // Tickvalssrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  tickvals . 
    Tickvalssrc String `json:"tickvalssrc,omitempty"`
    
    // Tickwidth 
    // number 
    // Sets the tick width (in px). 
    Tickwidth float64 `json:"tickwidth,omitempty"`
    
    // Visible 
    // boolean 
    // A single toggle to hide the axis while preserving interaction like dragging. Default is true when a cheater plot is present on the axis, otherwise false 
    Visible Bool `json:"visible,omitempty"`
    
}
// IndicatorGaugeBarLine 
type IndicatorGaugeBarLine struct {
    
    // Color 
    // color 
    // Sets the color of the line enclosing each sector. 
    Color Color `json:"color,omitempty"`
    
    // Width 
    // number 
    // Sets the width (in px) of the line enclosing each sector. 
    Width float64 `json:"width,omitempty"`
    
}
// IndicatorGaugeBar Set the appearance of the gauge's value
type IndicatorGaugeBar struct {
    
    // Color 
    // color 
    // Sets the background color of the arc. 
    Color Color `json:"color,omitempty"`
    
    // Line 
    //   
    Line *IndicatorGaugeBarLine `json:"line,omitempty"`
    
    // Thickness 
    // number 
    // Sets the thickness of the bar as a fraction of the total thickness of the gauge. 
    Thickness float64 `json:"thickness,omitempty"`
    
}
// IndicatorGaugeThresholdLine 
type IndicatorGaugeThresholdLine struct {
    
    // Color 
    // color 
    // Sets the color of the threshold line. 
    Color Color `json:"color,omitempty"`
    
    // Width 
    // number 
    // Sets the width (in px) of the threshold line. 
    Width float64 `json:"width,omitempty"`
    
}
// IndicatorGaugeThreshold 
type IndicatorGaugeThreshold struct {
    
    // Line 
    //   
    Line *IndicatorGaugeThresholdLine `json:"line,omitempty"`
    
    // Thickness 
    // number 
    // Sets the thickness of the threshold line as a fraction of the thickness of the gauge. 
    Thickness float64 `json:"thickness,omitempty"`
    
    // Value 
    // number 
    // Sets a treshold value drawn as a line. 
    Value float64 `json:"value,omitempty"`
    
}
// IndicatorGauge The gauge of the Indicator plot.
type IndicatorGauge struct {
    
    // Axis 
    //   
    Axis *IndicatorGaugeAxis `json:"axis,omitempty"`
    
    // Bar 
    //  Set the appearance of the gauge's value 
    Bar *IndicatorGaugeBar `json:"bar,omitempty"`
    
    // Bgcolor 
    // color 
    // Sets the gauge background color. 
    Bgcolor Color `json:"bgcolor,omitempty"`
    
    // Bordercolor 
    // color 
    // Sets the color of the border enclosing the gauge. 
    Bordercolor Color `json:"bordercolor,omitempty"`
    
    // Borderwidth 
    // number 
    // Sets the width (in px) of the border enclosing the gauge. 
    Borderwidth float64 `json:"borderwidth,omitempty"`
    
    // Shape 
    // enumerated Set the shape of the gauge 
    Shape IndicatorGaugeShape `json:"shape,omitempty"`
    
    // Steps 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Steps interface{} `json:"steps,omitempty"`
    
    // Threshold 
    //   
    Threshold *IndicatorGaugeThreshold `json:"threshold,omitempty"`
    
}
// IndicatorNumberFont Set the font used to display main number
type IndicatorNumberFont struct {
    
    // Color 
    // color 
    //  
    Color Color `json:"color,omitempty"`
    
    // Family 
    // string 
    // HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*. 
    Family String `json:"family,omitempty"`
    
    // Size 
    // number 
    //  
    Size float64 `json:"size,omitempty"`
    
}
// IndicatorNumber 
type IndicatorNumber struct {
    
    // Font 
    //  Set the font used to display main number 
    Font *IndicatorNumberFont `json:"font,omitempty"`
    
    // Prefix 
    // string 
    // Sets a prefix appearing before the number. 
    Prefix String `json:"prefix,omitempty"`
    
    // Suffix 
    // string 
    // Sets a suffix appearing next to the number. 
    Suffix String `json:"suffix,omitempty"`
    
    // Valueformat 
    // string 
    // Sets the value formatting rule using d3 formatting mini-language which is similar to those of Python. See https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format 
    Valueformat String `json:"valueformat,omitempty"`
    
}
// IndicatorStream 
type IndicatorStream struct {
    
    // Maxpoints 
    // number 
    // Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot. 
    Maxpoints float64 `json:"maxpoints,omitempty"`
    
    // Token 
    // string 
    // The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details. 
    Token String `json:"token,omitempty"`
    
}
// IndicatorTitleFont Set the font used to display the title
type IndicatorTitleFont struct {
    
    // Color 
    // color 
    //  
    Color Color `json:"color,omitempty"`
    
    // Family 
    // string 
    // HTML font family - the typeface that will be applied by the web browser. The web browser will only be able to apply a font if it is available on the system which it operates. Provide multiple font families, separated by commas, to indicate the preference in which to apply fonts if they aren't available on the system. The Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise) generates images on a server, where only a select number of fonts are installed and supported. These include *Arial*, *Balto*, *Courier New*, *Droid Sans*,, *Droid Serif*, *Droid Sans Mono*, *Gravitas One*, *Old Standard TT*, *Open Sans*, *Overpass*, *PT Sans Narrow*, *Raleway*, *Times New Roman*. 
    Family String `json:"family,omitempty"`
    
    // Size 
    // number 
    //  
    Size float64 `json:"size,omitempty"`
    
}
// IndicatorTitle 
type IndicatorTitle struct {
    
    // Align 
    // enumerated Sets the horizontal alignment of the title. It defaults to `center` except for bullet charts for which it defaults to right. 
    Align IndicatorTitleAlign `json:"align,omitempty"`
    
    // Font 
    //  Set the font used to display the title 
    Font *IndicatorTitleFont `json:"font,omitempty"`
    
    // Text 
    // string 
    // Sets the title of this indicator. 
    Text String `json:"text,omitempty"`
    
}
// IndicatorAlign Sets the horizontal alignment of the `text` within the box. Note that this attribute has no effect if an angular gauge is displayed: in this case, it is always centered
type IndicatorAlign string 

const (
    IndicatorAlignLeft IndicatorAlign = "left"
    IndicatorAlignCenter IndicatorAlign = "center"
    IndicatorAlignRight IndicatorAlign = "right"
    
)
// IndicatorDeltaPosition Sets the position of delta with respect to the number.
type IndicatorDeltaPosition string 

const (
    IndicatorDeltaPositionTop IndicatorDeltaPosition = "top"
    IndicatorDeltaPositionBottom IndicatorDeltaPosition = "bottom"
    IndicatorDeltaPositionLeft IndicatorDeltaPosition = "left"
    IndicatorDeltaPositionRight IndicatorDeltaPosition = "right"
    
)
// IndicatorGaugeAxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type IndicatorGaugeAxisExponentformat string 

const (
    IndicatorGaugeAxisExponentformatNone IndicatorGaugeAxisExponentformat = "none"
    IndicatorGaugeAxisExponentformatE1 IndicatorGaugeAxisExponentformat = "e"
    IndicatorGaugeAxisExponentformatE2 IndicatorGaugeAxisExponentformat = "E"
    IndicatorGaugeAxisExponentformatPower IndicatorGaugeAxisExponentformat = "power"
    IndicatorGaugeAxisExponentformatSi IndicatorGaugeAxisExponentformat = "SI"
    IndicatorGaugeAxisExponentformatB IndicatorGaugeAxisExponentformat = "B"
    
)
// IndicatorGaugeAxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type IndicatorGaugeAxisShowexponent string 

const (
    IndicatorGaugeAxisShowexponentAll IndicatorGaugeAxisShowexponent = "all"
    IndicatorGaugeAxisShowexponentFirst IndicatorGaugeAxisShowexponent = "first"
    IndicatorGaugeAxisShowexponentLast IndicatorGaugeAxisShowexponent = "last"
    IndicatorGaugeAxisShowexponentNone IndicatorGaugeAxisShowexponent = "none"
    
)
// IndicatorGaugeAxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type IndicatorGaugeAxisShowtickprefix string 

const (
    IndicatorGaugeAxisShowtickprefixAll IndicatorGaugeAxisShowtickprefix = "all"
    IndicatorGaugeAxisShowtickprefixFirst IndicatorGaugeAxisShowtickprefix = "first"
    IndicatorGaugeAxisShowtickprefixLast IndicatorGaugeAxisShowtickprefix = "last"
    IndicatorGaugeAxisShowtickprefixNone IndicatorGaugeAxisShowtickprefix = "none"
    
)
// IndicatorGaugeAxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type IndicatorGaugeAxisShowticksuffix string 

const (
    IndicatorGaugeAxisShowticksuffixAll IndicatorGaugeAxisShowticksuffix = "all"
    IndicatorGaugeAxisShowticksuffixFirst IndicatorGaugeAxisShowticksuffix = "first"
    IndicatorGaugeAxisShowticksuffixLast IndicatorGaugeAxisShowticksuffix = "last"
    IndicatorGaugeAxisShowticksuffixNone IndicatorGaugeAxisShowticksuffix = "none"
    
)
// IndicatorGaugeAxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type IndicatorGaugeAxisTickmode string 

const (
    IndicatorGaugeAxisTickmodeAuto IndicatorGaugeAxisTickmode = "auto"
    IndicatorGaugeAxisTickmodeLinear IndicatorGaugeAxisTickmode = "linear"
    IndicatorGaugeAxisTickmodeArray IndicatorGaugeAxisTickmode = "array"
    
)
// IndicatorGaugeAxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type IndicatorGaugeAxisTicks string 

const (
    IndicatorGaugeAxisTicksOutside IndicatorGaugeAxisTicks = "outside"
    IndicatorGaugeAxisTicksInside IndicatorGaugeAxisTicks = "inside"
    IndicatorGaugeAxisTicksEmpty IndicatorGaugeAxisTicks = ""
    
)
// IndicatorGaugeShape Set the shape of the gauge
type IndicatorGaugeShape string 

const (
    IndicatorGaugeShapeAngular IndicatorGaugeShape = "angular"
    IndicatorGaugeShapeBullet IndicatorGaugeShape = "bullet"
    
)
// IndicatorTitleAlign Sets the horizontal alignment of the title. It defaults to `center` except for bullet charts for which it defaults to right.
type IndicatorTitleAlign string 

const (
    IndicatorTitleAlignLeft IndicatorTitleAlign = "left"
    IndicatorTitleAlignCenter IndicatorTitleAlign = "center"
    IndicatorTitleAlignRight IndicatorTitleAlign = "right"
    
)
// IndicatorVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type IndicatorVisible interface{} 

var (
    IndicatorVisibleTrue IndicatorVisible = true
    IndicatorVisibleFalse IndicatorVisible = false
    IndicatorVisibleLegendonly IndicatorVisible = "legendonly"
    
)
// IndicatorMode Determines how the value is displayed on the graph. `number` displays the value numerically in text. `delta` displays the difference to a reference value in text. Finally, `gauge` displays the value graphically on an axis.
type IndicatorMode string 

const (
    // Flags
    IndicatorModeNumber IndicatorMode = "number"
    IndicatorModeDelta IndicatorMode = "delta"
    IndicatorModeGauge IndicatorMode = "gauge"
    
    // Extra
    
)
