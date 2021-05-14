package grob

var TraceTypeArea TraceType = "area"

func (trace *Area) GetType() TraceType {
	return TraceTypeArea
}
// Area 
type Area struct {
    
    // Type 
    // is the type of the plot 
    Type TraceType `json:"type,omitempty"`
    
    // Customdata 
    // data_array 
    // Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements 
    Customdata interface{} `json:"customdata,omitempty"`
    
    // Customdatasrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  customdata . 
    Customdatasrc String `json:"customdatasrc,omitempty"`
    
    // Hoverinfo 
    // flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired. 
    Hoverinfo AreaHoverinfo `json:"hoverinfo,omitempty"`
    
    // Hoverinfosrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hoverinfo . 
    Hoverinfosrc String `json:"hoverinfosrc,omitempty"`
    
    // Hoverlabel 
    //   
    Hoverlabel *AreaHoverlabel `json:"hoverlabel,omitempty"`
    
    // Ids 
    // data_array 
    // Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type. 
    Ids interface{} `json:"ids,omitempty"`
    
    // Idssrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  ids . 
    Idssrc String `json:"idssrc,omitempty"`
    
    // Legendgroup 
    // string 
    // Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items. 
    Legendgroup String `json:"legendgroup,omitempty"`
    
    // Marker 
    //   
    Marker *AreaMarker `json:"marker,omitempty"`
    
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
    
    // R 
    // data_array 
    // Area traces are deprecated! Please switch to the *barpolar* trace type. Sets the radial coordinates for legacy polar chart only. 
    R interface{} `json:"r,omitempty"`
    
    // Rsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  r . 
    Rsrc String `json:"rsrc,omitempty"`
    
    // Showlegend 
    // boolean 
    // Determines whether or not an item corresponding to this trace is shown in the legend. 
    Showlegend Bool `json:"showlegend,omitempty"`
    
    // Stream 
    //   
    Stream *AreaStream `json:"stream,omitempty"`
    
    // T 
    // data_array 
    // Area traces are deprecated! Please switch to the *barpolar* trace type. Sets the angular coordinates for legacy polar chart only. 
    T interface{} `json:"t,omitempty"`
    
    // Transforms 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Transforms interface{} `json:"transforms,omitempty"`
    
    // Tsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  t . 
    Tsrc String `json:"tsrc,omitempty"`
    
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
    Visible AreaVisible `json:"visible,omitempty"`
    
}
// AreaHoverlabelFont Sets the font used in hover labels.
type AreaHoverlabelFont struct {
    
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
// AreaHoverlabel 
type AreaHoverlabel struct {
    
    // Align 
    // enumerated Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines 
    Align AreaHoverlabelAlign `json:"align,omitempty"`
    
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
    Font *AreaHoverlabelFont `json:"font,omitempty"`
    
    // Namelength 
    // integer 
    // Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis. 
    Namelength int64 `json:"namelength,omitempty"`
    
    // Namelengthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  namelength . 
    Namelengthsrc String `json:"namelengthsrc,omitempty"`
    
}
// AreaMarker 
type AreaMarker struct {
    
    // Color 
    // color 
    // Area traces are deprecated! Please switch to the *barpolar* trace type. Sets themarkercolor. It accepts either a specific color or an array of numbers that are mapped to the colorscale relative to the max and min values of the array or relative to `marker.cmin` and `marker.cmax` if set. 
    Color Color `json:"color,omitempty"`
    
    // Colorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  color . 
    Colorsrc String `json:"colorsrc,omitempty"`
    
    // Opacity 
    // number 
    // Area traces are deprecated! Please switch to the *barpolar* trace type. Sets the marker opacity. 
    Opacity float64 `json:"opacity,omitempty"`
    
    // Opacitysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  opacity . 
    Opacitysrc String `json:"opacitysrc,omitempty"`
    
    // Size 
    // number 
    // Area traces are deprecated! Please switch to the *barpolar* trace type. Sets the marker size (in px). 
    Size float64 `json:"size,omitempty"`
    
    // Sizesrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  size . 
    Sizesrc String `json:"sizesrc,omitempty"`
    
    // Symbol 
    // enumerated Area traces are deprecated! Please switch to the *barpolar* trace type. Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name. 
    Symbol AreaMarkerSymbol `json:"symbol,omitempty"`
    
    // Symbolsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  symbol . 
    Symbolsrc String `json:"symbolsrc,omitempty"`
    
}
// AreaStream 
type AreaStream struct {
    
    // Maxpoints 
    // number 
    // Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot. 
    Maxpoints float64 `json:"maxpoints,omitempty"`
    
    // Token 
    // string 
    // The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details. 
    Token String `json:"token,omitempty"`
    
}
// AreaHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type AreaHoverlabelAlign string 

const (
    AreaHoverlabelAlignLeft AreaHoverlabelAlign = "left"
    AreaHoverlabelAlignRight AreaHoverlabelAlign = "right"
    AreaHoverlabelAlignAuto AreaHoverlabelAlign = "auto"
    
)
// AreaMarkerSymbol Area traces are deprecated! Please switch to the *barpolar* trace type. Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type AreaMarkerSymbol interface{} 

var (
    AreaMarkerSymbolNumber0 AreaMarkerSymbol = 0
    AreaMarkerSymbol0 AreaMarkerSymbol = "0"
    AreaMarkerSymbolCircle AreaMarkerSymbol = "circle"
    AreaMarkerSymbolNumber100 AreaMarkerSymbol = 100
    AreaMarkerSymbol100 AreaMarkerSymbol = "100"
    AreaMarkerSymbolCircleOpen AreaMarkerSymbol = "circle-open"
    AreaMarkerSymbolNumber200 AreaMarkerSymbol = 200
    AreaMarkerSymbol200 AreaMarkerSymbol = "200"
    AreaMarkerSymbolCircleDot AreaMarkerSymbol = "circle-dot"
    AreaMarkerSymbolNumber300 AreaMarkerSymbol = 300
    AreaMarkerSymbol300 AreaMarkerSymbol = "300"
    AreaMarkerSymbolCircleOpenDot AreaMarkerSymbol = "circle-open-dot"
    AreaMarkerSymbolNumber1 AreaMarkerSymbol = 1
    AreaMarkerSymbol1 AreaMarkerSymbol = "1"
    AreaMarkerSymbolSquare AreaMarkerSymbol = "square"
    AreaMarkerSymbolNumber101 AreaMarkerSymbol = 101
    AreaMarkerSymbol101 AreaMarkerSymbol = "101"
    AreaMarkerSymbolSquareOpen AreaMarkerSymbol = "square-open"
    AreaMarkerSymbolNumber201 AreaMarkerSymbol = 201
    AreaMarkerSymbol201 AreaMarkerSymbol = "201"
    AreaMarkerSymbolSquareDot AreaMarkerSymbol = "square-dot"
    AreaMarkerSymbolNumber301 AreaMarkerSymbol = 301
    AreaMarkerSymbol301 AreaMarkerSymbol = "301"
    AreaMarkerSymbolSquareOpenDot AreaMarkerSymbol = "square-open-dot"
    AreaMarkerSymbolNumber2 AreaMarkerSymbol = 2
    AreaMarkerSymbol2 AreaMarkerSymbol = "2"
    AreaMarkerSymbolDiamond AreaMarkerSymbol = "diamond"
    AreaMarkerSymbolNumber102 AreaMarkerSymbol = 102
    AreaMarkerSymbol102 AreaMarkerSymbol = "102"
    AreaMarkerSymbolDiamondOpen AreaMarkerSymbol = "diamond-open"
    AreaMarkerSymbolNumber202 AreaMarkerSymbol = 202
    AreaMarkerSymbol202 AreaMarkerSymbol = "202"
    AreaMarkerSymbolDiamondDot AreaMarkerSymbol = "diamond-dot"
    AreaMarkerSymbolNumber302 AreaMarkerSymbol = 302
    AreaMarkerSymbol302 AreaMarkerSymbol = "302"
    AreaMarkerSymbolDiamondOpenDot AreaMarkerSymbol = "diamond-open-dot"
    AreaMarkerSymbolNumber3 AreaMarkerSymbol = 3
    AreaMarkerSymbol3 AreaMarkerSymbol = "3"
    AreaMarkerSymbolCross AreaMarkerSymbol = "cross"
    AreaMarkerSymbolNumber103 AreaMarkerSymbol = 103
    AreaMarkerSymbol103 AreaMarkerSymbol = "103"
    AreaMarkerSymbolCrossOpen AreaMarkerSymbol = "cross-open"
    AreaMarkerSymbolNumber203 AreaMarkerSymbol = 203
    AreaMarkerSymbol203 AreaMarkerSymbol = "203"
    AreaMarkerSymbolCrossDot AreaMarkerSymbol = "cross-dot"
    AreaMarkerSymbolNumber303 AreaMarkerSymbol = 303
    AreaMarkerSymbol303 AreaMarkerSymbol = "303"
    AreaMarkerSymbolCrossOpenDot AreaMarkerSymbol = "cross-open-dot"
    AreaMarkerSymbolNumber4 AreaMarkerSymbol = 4
    AreaMarkerSymbol4 AreaMarkerSymbol = "4"
    AreaMarkerSymbolX AreaMarkerSymbol = "x"
    AreaMarkerSymbolNumber104 AreaMarkerSymbol = 104
    AreaMarkerSymbol104 AreaMarkerSymbol = "104"
    AreaMarkerSymbolXOpen AreaMarkerSymbol = "x-open"
    AreaMarkerSymbolNumber204 AreaMarkerSymbol = 204
    AreaMarkerSymbol204 AreaMarkerSymbol = "204"
    AreaMarkerSymbolXDot AreaMarkerSymbol = "x-dot"
    AreaMarkerSymbolNumber304 AreaMarkerSymbol = 304
    AreaMarkerSymbol304 AreaMarkerSymbol = "304"
    AreaMarkerSymbolXOpenDot AreaMarkerSymbol = "x-open-dot"
    AreaMarkerSymbolNumber5 AreaMarkerSymbol = 5
    AreaMarkerSymbol5 AreaMarkerSymbol = "5"
    AreaMarkerSymbolTriangleUp AreaMarkerSymbol = "triangle-up"
    AreaMarkerSymbolNumber105 AreaMarkerSymbol = 105
    AreaMarkerSymbol105 AreaMarkerSymbol = "105"
    AreaMarkerSymbolTriangleUpOpen AreaMarkerSymbol = "triangle-up-open"
    AreaMarkerSymbolNumber205 AreaMarkerSymbol = 205
    AreaMarkerSymbol205 AreaMarkerSymbol = "205"
    AreaMarkerSymbolTriangleUpDot AreaMarkerSymbol = "triangle-up-dot"
    AreaMarkerSymbolNumber305 AreaMarkerSymbol = 305
    AreaMarkerSymbol305 AreaMarkerSymbol = "305"
    AreaMarkerSymbolTriangleUpOpenDot AreaMarkerSymbol = "triangle-up-open-dot"
    AreaMarkerSymbolNumber6 AreaMarkerSymbol = 6
    AreaMarkerSymbol6 AreaMarkerSymbol = "6"
    AreaMarkerSymbolTriangleDown AreaMarkerSymbol = "triangle-down"
    AreaMarkerSymbolNumber106 AreaMarkerSymbol = 106
    AreaMarkerSymbol106 AreaMarkerSymbol = "106"
    AreaMarkerSymbolTriangleDownOpen AreaMarkerSymbol = "triangle-down-open"
    AreaMarkerSymbolNumber206 AreaMarkerSymbol = 206
    AreaMarkerSymbol206 AreaMarkerSymbol = "206"
    AreaMarkerSymbolTriangleDownDot AreaMarkerSymbol = "triangle-down-dot"
    AreaMarkerSymbolNumber306 AreaMarkerSymbol = 306
    AreaMarkerSymbol306 AreaMarkerSymbol = "306"
    AreaMarkerSymbolTriangleDownOpenDot AreaMarkerSymbol = "triangle-down-open-dot"
    AreaMarkerSymbolNumber7 AreaMarkerSymbol = 7
    AreaMarkerSymbol7 AreaMarkerSymbol = "7"
    AreaMarkerSymbolTriangleLeft AreaMarkerSymbol = "triangle-left"
    AreaMarkerSymbolNumber107 AreaMarkerSymbol = 107
    AreaMarkerSymbol107 AreaMarkerSymbol = "107"
    AreaMarkerSymbolTriangleLeftOpen AreaMarkerSymbol = "triangle-left-open"
    AreaMarkerSymbolNumber207 AreaMarkerSymbol = 207
    AreaMarkerSymbol207 AreaMarkerSymbol = "207"
    AreaMarkerSymbolTriangleLeftDot AreaMarkerSymbol = "triangle-left-dot"
    AreaMarkerSymbolNumber307 AreaMarkerSymbol = 307
    AreaMarkerSymbol307 AreaMarkerSymbol = "307"
    AreaMarkerSymbolTriangleLeftOpenDot AreaMarkerSymbol = "triangle-left-open-dot"
    AreaMarkerSymbolNumber8 AreaMarkerSymbol = 8
    AreaMarkerSymbol8 AreaMarkerSymbol = "8"
    AreaMarkerSymbolTriangleRight AreaMarkerSymbol = "triangle-right"
    AreaMarkerSymbolNumber108 AreaMarkerSymbol = 108
    AreaMarkerSymbol108 AreaMarkerSymbol = "108"
    AreaMarkerSymbolTriangleRightOpen AreaMarkerSymbol = "triangle-right-open"
    AreaMarkerSymbolNumber208 AreaMarkerSymbol = 208
    AreaMarkerSymbol208 AreaMarkerSymbol = "208"
    AreaMarkerSymbolTriangleRightDot AreaMarkerSymbol = "triangle-right-dot"
    AreaMarkerSymbolNumber308 AreaMarkerSymbol = 308
    AreaMarkerSymbol308 AreaMarkerSymbol = "308"
    AreaMarkerSymbolTriangleRightOpenDot AreaMarkerSymbol = "triangle-right-open-dot"
    AreaMarkerSymbolNumber9 AreaMarkerSymbol = 9
    AreaMarkerSymbol9 AreaMarkerSymbol = "9"
    AreaMarkerSymbolTriangleNe AreaMarkerSymbol = "triangle-ne"
    AreaMarkerSymbolNumber109 AreaMarkerSymbol = 109
    AreaMarkerSymbol109 AreaMarkerSymbol = "109"
    AreaMarkerSymbolTriangleNeOpen AreaMarkerSymbol = "triangle-ne-open"
    AreaMarkerSymbolNumber209 AreaMarkerSymbol = 209
    AreaMarkerSymbol209 AreaMarkerSymbol = "209"
    AreaMarkerSymbolTriangleNeDot AreaMarkerSymbol = "triangle-ne-dot"
    AreaMarkerSymbolNumber309 AreaMarkerSymbol = 309
    AreaMarkerSymbol309 AreaMarkerSymbol = "309"
    AreaMarkerSymbolTriangleNeOpenDot AreaMarkerSymbol = "triangle-ne-open-dot"
    AreaMarkerSymbolNumber10 AreaMarkerSymbol = 10
    AreaMarkerSymbol10 AreaMarkerSymbol = "10"
    AreaMarkerSymbolTriangleSe AreaMarkerSymbol = "triangle-se"
    AreaMarkerSymbolNumber110 AreaMarkerSymbol = 110
    AreaMarkerSymbol110 AreaMarkerSymbol = "110"
    AreaMarkerSymbolTriangleSeOpen AreaMarkerSymbol = "triangle-se-open"
    AreaMarkerSymbolNumber210 AreaMarkerSymbol = 210
    AreaMarkerSymbol210 AreaMarkerSymbol = "210"
    AreaMarkerSymbolTriangleSeDot AreaMarkerSymbol = "triangle-se-dot"
    AreaMarkerSymbolNumber310 AreaMarkerSymbol = 310
    AreaMarkerSymbol310 AreaMarkerSymbol = "310"
    AreaMarkerSymbolTriangleSeOpenDot AreaMarkerSymbol = "triangle-se-open-dot"
    AreaMarkerSymbolNumber11 AreaMarkerSymbol = 11
    AreaMarkerSymbol11 AreaMarkerSymbol = "11"
    AreaMarkerSymbolTriangleSw AreaMarkerSymbol = "triangle-sw"
    AreaMarkerSymbolNumber111 AreaMarkerSymbol = 111
    AreaMarkerSymbol111 AreaMarkerSymbol = "111"
    AreaMarkerSymbolTriangleSwOpen AreaMarkerSymbol = "triangle-sw-open"
    AreaMarkerSymbolNumber211 AreaMarkerSymbol = 211
    AreaMarkerSymbol211 AreaMarkerSymbol = "211"
    AreaMarkerSymbolTriangleSwDot AreaMarkerSymbol = "triangle-sw-dot"
    AreaMarkerSymbolNumber311 AreaMarkerSymbol = 311
    AreaMarkerSymbol311 AreaMarkerSymbol = "311"
    AreaMarkerSymbolTriangleSwOpenDot AreaMarkerSymbol = "triangle-sw-open-dot"
    AreaMarkerSymbolNumber12 AreaMarkerSymbol = 12
    AreaMarkerSymbol12 AreaMarkerSymbol = "12"
    AreaMarkerSymbolTriangleNw AreaMarkerSymbol = "triangle-nw"
    AreaMarkerSymbolNumber112 AreaMarkerSymbol = 112
    AreaMarkerSymbol112 AreaMarkerSymbol = "112"
    AreaMarkerSymbolTriangleNwOpen AreaMarkerSymbol = "triangle-nw-open"
    AreaMarkerSymbolNumber212 AreaMarkerSymbol = 212
    AreaMarkerSymbol212 AreaMarkerSymbol = "212"
    AreaMarkerSymbolTriangleNwDot AreaMarkerSymbol = "triangle-nw-dot"
    AreaMarkerSymbolNumber312 AreaMarkerSymbol = 312
    AreaMarkerSymbol312 AreaMarkerSymbol = "312"
    AreaMarkerSymbolTriangleNwOpenDot AreaMarkerSymbol = "triangle-nw-open-dot"
    AreaMarkerSymbolNumber13 AreaMarkerSymbol = 13
    AreaMarkerSymbol13 AreaMarkerSymbol = "13"
    AreaMarkerSymbolPentagon AreaMarkerSymbol = "pentagon"
    AreaMarkerSymbolNumber113 AreaMarkerSymbol = 113
    AreaMarkerSymbol113 AreaMarkerSymbol = "113"
    AreaMarkerSymbolPentagonOpen AreaMarkerSymbol = "pentagon-open"
    AreaMarkerSymbolNumber213 AreaMarkerSymbol = 213
    AreaMarkerSymbol213 AreaMarkerSymbol = "213"
    AreaMarkerSymbolPentagonDot AreaMarkerSymbol = "pentagon-dot"
    AreaMarkerSymbolNumber313 AreaMarkerSymbol = 313
    AreaMarkerSymbol313 AreaMarkerSymbol = "313"
    AreaMarkerSymbolPentagonOpenDot AreaMarkerSymbol = "pentagon-open-dot"
    AreaMarkerSymbolNumber14 AreaMarkerSymbol = 14
    AreaMarkerSymbol14 AreaMarkerSymbol = "14"
    AreaMarkerSymbolHexagon AreaMarkerSymbol = "hexagon"
    AreaMarkerSymbolNumber114 AreaMarkerSymbol = 114
    AreaMarkerSymbol114 AreaMarkerSymbol = "114"
    AreaMarkerSymbolHexagonOpen AreaMarkerSymbol = "hexagon-open"
    AreaMarkerSymbolNumber214 AreaMarkerSymbol = 214
    AreaMarkerSymbol214 AreaMarkerSymbol = "214"
    AreaMarkerSymbolHexagonDot AreaMarkerSymbol = "hexagon-dot"
    AreaMarkerSymbolNumber314 AreaMarkerSymbol = 314
    AreaMarkerSymbol314 AreaMarkerSymbol = "314"
    AreaMarkerSymbolHexagonOpenDot AreaMarkerSymbol = "hexagon-open-dot"
    AreaMarkerSymbolNumber15 AreaMarkerSymbol = 15
    AreaMarkerSymbol15 AreaMarkerSymbol = "15"
    AreaMarkerSymbolHexagon2 AreaMarkerSymbol = "hexagon2"
    AreaMarkerSymbolNumber115 AreaMarkerSymbol = 115
    AreaMarkerSymbol115 AreaMarkerSymbol = "115"
    AreaMarkerSymbolHexagon2Open AreaMarkerSymbol = "hexagon2-open"
    AreaMarkerSymbolNumber215 AreaMarkerSymbol = 215
    AreaMarkerSymbol215 AreaMarkerSymbol = "215"
    AreaMarkerSymbolHexagon2Dot AreaMarkerSymbol = "hexagon2-dot"
    AreaMarkerSymbolNumber315 AreaMarkerSymbol = 315
    AreaMarkerSymbol315 AreaMarkerSymbol = "315"
    AreaMarkerSymbolHexagon2OpenDot AreaMarkerSymbol = "hexagon2-open-dot"
    AreaMarkerSymbolNumber16 AreaMarkerSymbol = 16
    AreaMarkerSymbol16 AreaMarkerSymbol = "16"
    AreaMarkerSymbolOctagon AreaMarkerSymbol = "octagon"
    AreaMarkerSymbolNumber116 AreaMarkerSymbol = 116
    AreaMarkerSymbol116 AreaMarkerSymbol = "116"
    AreaMarkerSymbolOctagonOpen AreaMarkerSymbol = "octagon-open"
    AreaMarkerSymbolNumber216 AreaMarkerSymbol = 216
    AreaMarkerSymbol216 AreaMarkerSymbol = "216"
    AreaMarkerSymbolOctagonDot AreaMarkerSymbol = "octagon-dot"
    AreaMarkerSymbolNumber316 AreaMarkerSymbol = 316
    AreaMarkerSymbol316 AreaMarkerSymbol = "316"
    AreaMarkerSymbolOctagonOpenDot AreaMarkerSymbol = "octagon-open-dot"
    AreaMarkerSymbolNumber17 AreaMarkerSymbol = 17
    AreaMarkerSymbol17 AreaMarkerSymbol = "17"
    AreaMarkerSymbolStar AreaMarkerSymbol = "star"
    AreaMarkerSymbolNumber117 AreaMarkerSymbol = 117
    AreaMarkerSymbol117 AreaMarkerSymbol = "117"
    AreaMarkerSymbolStarOpen AreaMarkerSymbol = "star-open"
    AreaMarkerSymbolNumber217 AreaMarkerSymbol = 217
    AreaMarkerSymbol217 AreaMarkerSymbol = "217"
    AreaMarkerSymbolStarDot AreaMarkerSymbol = "star-dot"
    AreaMarkerSymbolNumber317 AreaMarkerSymbol = 317
    AreaMarkerSymbol317 AreaMarkerSymbol = "317"
    AreaMarkerSymbolStarOpenDot AreaMarkerSymbol = "star-open-dot"
    AreaMarkerSymbolNumber18 AreaMarkerSymbol = 18
    AreaMarkerSymbol18 AreaMarkerSymbol = "18"
    AreaMarkerSymbolHexagram AreaMarkerSymbol = "hexagram"
    AreaMarkerSymbolNumber118 AreaMarkerSymbol = 118
    AreaMarkerSymbol118 AreaMarkerSymbol = "118"
    AreaMarkerSymbolHexagramOpen AreaMarkerSymbol = "hexagram-open"
    AreaMarkerSymbolNumber218 AreaMarkerSymbol = 218
    AreaMarkerSymbol218 AreaMarkerSymbol = "218"
    AreaMarkerSymbolHexagramDot AreaMarkerSymbol = "hexagram-dot"
    AreaMarkerSymbolNumber318 AreaMarkerSymbol = 318
    AreaMarkerSymbol318 AreaMarkerSymbol = "318"
    AreaMarkerSymbolHexagramOpenDot AreaMarkerSymbol = "hexagram-open-dot"
    AreaMarkerSymbolNumber19 AreaMarkerSymbol = 19
    AreaMarkerSymbol19 AreaMarkerSymbol = "19"
    AreaMarkerSymbolStarTriangleUp AreaMarkerSymbol = "star-triangle-up"
    AreaMarkerSymbolNumber119 AreaMarkerSymbol = 119
    AreaMarkerSymbol119 AreaMarkerSymbol = "119"
    AreaMarkerSymbolStarTriangleUpOpen AreaMarkerSymbol = "star-triangle-up-open"
    AreaMarkerSymbolNumber219 AreaMarkerSymbol = 219
    AreaMarkerSymbol219 AreaMarkerSymbol = "219"
    AreaMarkerSymbolStarTriangleUpDot AreaMarkerSymbol = "star-triangle-up-dot"
    AreaMarkerSymbolNumber319 AreaMarkerSymbol = 319
    AreaMarkerSymbol319 AreaMarkerSymbol = "319"
    AreaMarkerSymbolStarTriangleUpOpenDot AreaMarkerSymbol = "star-triangle-up-open-dot"
    AreaMarkerSymbolNumber20 AreaMarkerSymbol = 20
    AreaMarkerSymbol20 AreaMarkerSymbol = "20"
    AreaMarkerSymbolStarTriangleDown AreaMarkerSymbol = "star-triangle-down"
    AreaMarkerSymbolNumber120 AreaMarkerSymbol = 120
    AreaMarkerSymbol120 AreaMarkerSymbol = "120"
    AreaMarkerSymbolStarTriangleDownOpen AreaMarkerSymbol = "star-triangle-down-open"
    AreaMarkerSymbolNumber220 AreaMarkerSymbol = 220
    AreaMarkerSymbol220 AreaMarkerSymbol = "220"
    AreaMarkerSymbolStarTriangleDownDot AreaMarkerSymbol = "star-triangle-down-dot"
    AreaMarkerSymbolNumber320 AreaMarkerSymbol = 320
    AreaMarkerSymbol320 AreaMarkerSymbol = "320"
    AreaMarkerSymbolStarTriangleDownOpenDot AreaMarkerSymbol = "star-triangle-down-open-dot"
    AreaMarkerSymbolNumber21 AreaMarkerSymbol = 21
    AreaMarkerSymbol21 AreaMarkerSymbol = "21"
    AreaMarkerSymbolStarSquare AreaMarkerSymbol = "star-square"
    AreaMarkerSymbolNumber121 AreaMarkerSymbol = 121
    AreaMarkerSymbol121 AreaMarkerSymbol = "121"
    AreaMarkerSymbolStarSquareOpen AreaMarkerSymbol = "star-square-open"
    AreaMarkerSymbolNumber221 AreaMarkerSymbol = 221
    AreaMarkerSymbol221 AreaMarkerSymbol = "221"
    AreaMarkerSymbolStarSquareDot AreaMarkerSymbol = "star-square-dot"
    AreaMarkerSymbolNumber321 AreaMarkerSymbol = 321
    AreaMarkerSymbol321 AreaMarkerSymbol = "321"
    AreaMarkerSymbolStarSquareOpenDot AreaMarkerSymbol = "star-square-open-dot"
    AreaMarkerSymbolNumber22 AreaMarkerSymbol = 22
    AreaMarkerSymbol22 AreaMarkerSymbol = "22"
    AreaMarkerSymbolStarDiamond AreaMarkerSymbol = "star-diamond"
    AreaMarkerSymbolNumber122 AreaMarkerSymbol = 122
    AreaMarkerSymbol122 AreaMarkerSymbol = "122"
    AreaMarkerSymbolStarDiamondOpen AreaMarkerSymbol = "star-diamond-open"
    AreaMarkerSymbolNumber222 AreaMarkerSymbol = 222
    AreaMarkerSymbol222 AreaMarkerSymbol = "222"
    AreaMarkerSymbolStarDiamondDot AreaMarkerSymbol = "star-diamond-dot"
    AreaMarkerSymbolNumber322 AreaMarkerSymbol = 322
    AreaMarkerSymbol322 AreaMarkerSymbol = "322"
    AreaMarkerSymbolStarDiamondOpenDot AreaMarkerSymbol = "star-diamond-open-dot"
    AreaMarkerSymbolNumber23 AreaMarkerSymbol = 23
    AreaMarkerSymbol23 AreaMarkerSymbol = "23"
    AreaMarkerSymbolDiamondTall AreaMarkerSymbol = "diamond-tall"
    AreaMarkerSymbolNumber123 AreaMarkerSymbol = 123
    AreaMarkerSymbol123 AreaMarkerSymbol = "123"
    AreaMarkerSymbolDiamondTallOpen AreaMarkerSymbol = "diamond-tall-open"
    AreaMarkerSymbolNumber223 AreaMarkerSymbol = 223
    AreaMarkerSymbol223 AreaMarkerSymbol = "223"
    AreaMarkerSymbolDiamondTallDot AreaMarkerSymbol = "diamond-tall-dot"
    AreaMarkerSymbolNumber323 AreaMarkerSymbol = 323
    AreaMarkerSymbol323 AreaMarkerSymbol = "323"
    AreaMarkerSymbolDiamondTallOpenDot AreaMarkerSymbol = "diamond-tall-open-dot"
    AreaMarkerSymbolNumber24 AreaMarkerSymbol = 24
    AreaMarkerSymbol24 AreaMarkerSymbol = "24"
    AreaMarkerSymbolDiamondWide AreaMarkerSymbol = "diamond-wide"
    AreaMarkerSymbolNumber124 AreaMarkerSymbol = 124
    AreaMarkerSymbol124 AreaMarkerSymbol = "124"
    AreaMarkerSymbolDiamondWideOpen AreaMarkerSymbol = "diamond-wide-open"
    AreaMarkerSymbolNumber224 AreaMarkerSymbol = 224
    AreaMarkerSymbol224 AreaMarkerSymbol = "224"
    AreaMarkerSymbolDiamondWideDot AreaMarkerSymbol = "diamond-wide-dot"
    AreaMarkerSymbolNumber324 AreaMarkerSymbol = 324
    AreaMarkerSymbol324 AreaMarkerSymbol = "324"
    AreaMarkerSymbolDiamondWideOpenDot AreaMarkerSymbol = "diamond-wide-open-dot"
    AreaMarkerSymbolNumber25 AreaMarkerSymbol = 25
    AreaMarkerSymbol25 AreaMarkerSymbol = "25"
    AreaMarkerSymbolHourglass AreaMarkerSymbol = "hourglass"
    AreaMarkerSymbolNumber125 AreaMarkerSymbol = 125
    AreaMarkerSymbol125 AreaMarkerSymbol = "125"
    AreaMarkerSymbolHourglassOpen AreaMarkerSymbol = "hourglass-open"
    AreaMarkerSymbolNumber26 AreaMarkerSymbol = 26
    AreaMarkerSymbol26 AreaMarkerSymbol = "26"
    AreaMarkerSymbolBowtie AreaMarkerSymbol = "bowtie"
    AreaMarkerSymbolNumber126 AreaMarkerSymbol = 126
    AreaMarkerSymbol126 AreaMarkerSymbol = "126"
    AreaMarkerSymbolBowtieOpen AreaMarkerSymbol = "bowtie-open"
    AreaMarkerSymbolNumber27 AreaMarkerSymbol = 27
    AreaMarkerSymbol27 AreaMarkerSymbol = "27"
    AreaMarkerSymbolCircleCross AreaMarkerSymbol = "circle-cross"
    AreaMarkerSymbolNumber127 AreaMarkerSymbol = 127
    AreaMarkerSymbol127 AreaMarkerSymbol = "127"
    AreaMarkerSymbolCircleCrossOpen AreaMarkerSymbol = "circle-cross-open"
    AreaMarkerSymbolNumber28 AreaMarkerSymbol = 28
    AreaMarkerSymbol28 AreaMarkerSymbol = "28"
    AreaMarkerSymbolCircleX AreaMarkerSymbol = "circle-x"
    AreaMarkerSymbolNumber128 AreaMarkerSymbol = 128
    AreaMarkerSymbol128 AreaMarkerSymbol = "128"
    AreaMarkerSymbolCircleXOpen AreaMarkerSymbol = "circle-x-open"
    AreaMarkerSymbolNumber29 AreaMarkerSymbol = 29
    AreaMarkerSymbol29 AreaMarkerSymbol = "29"
    AreaMarkerSymbolSquareCross AreaMarkerSymbol = "square-cross"
    AreaMarkerSymbolNumber129 AreaMarkerSymbol = 129
    AreaMarkerSymbol129 AreaMarkerSymbol = "129"
    AreaMarkerSymbolSquareCrossOpen AreaMarkerSymbol = "square-cross-open"
    AreaMarkerSymbolNumber30 AreaMarkerSymbol = 30
    AreaMarkerSymbol30 AreaMarkerSymbol = "30"
    AreaMarkerSymbolSquareX AreaMarkerSymbol = "square-x"
    AreaMarkerSymbolNumber130 AreaMarkerSymbol = 130
    AreaMarkerSymbol130 AreaMarkerSymbol = "130"
    AreaMarkerSymbolSquareXOpen AreaMarkerSymbol = "square-x-open"
    AreaMarkerSymbolNumber31 AreaMarkerSymbol = 31
    AreaMarkerSymbol31 AreaMarkerSymbol = "31"
    AreaMarkerSymbolDiamondCross AreaMarkerSymbol = "diamond-cross"
    AreaMarkerSymbolNumber131 AreaMarkerSymbol = 131
    AreaMarkerSymbol131 AreaMarkerSymbol = "131"
    AreaMarkerSymbolDiamondCrossOpen AreaMarkerSymbol = "diamond-cross-open"
    AreaMarkerSymbolNumber32 AreaMarkerSymbol = 32
    AreaMarkerSymbol32 AreaMarkerSymbol = "32"
    AreaMarkerSymbolDiamondX AreaMarkerSymbol = "diamond-x"
    AreaMarkerSymbolNumber132 AreaMarkerSymbol = 132
    AreaMarkerSymbol132 AreaMarkerSymbol = "132"
    AreaMarkerSymbolDiamondXOpen AreaMarkerSymbol = "diamond-x-open"
    AreaMarkerSymbolNumber33 AreaMarkerSymbol = 33
    AreaMarkerSymbol33 AreaMarkerSymbol = "33"
    AreaMarkerSymbolCrossThin AreaMarkerSymbol = "cross-thin"
    AreaMarkerSymbolNumber133 AreaMarkerSymbol = 133
    AreaMarkerSymbol133 AreaMarkerSymbol = "133"
    AreaMarkerSymbolCrossThinOpen AreaMarkerSymbol = "cross-thin-open"
    AreaMarkerSymbolNumber34 AreaMarkerSymbol = 34
    AreaMarkerSymbol34 AreaMarkerSymbol = "34"
    AreaMarkerSymbolXThin AreaMarkerSymbol = "x-thin"
    AreaMarkerSymbolNumber134 AreaMarkerSymbol = 134
    AreaMarkerSymbol134 AreaMarkerSymbol = "134"
    AreaMarkerSymbolXThinOpen AreaMarkerSymbol = "x-thin-open"
    AreaMarkerSymbolNumber35 AreaMarkerSymbol = 35
    AreaMarkerSymbol35 AreaMarkerSymbol = "35"
    AreaMarkerSymbolAsterisk AreaMarkerSymbol = "asterisk"
    AreaMarkerSymbolNumber135 AreaMarkerSymbol = 135
    AreaMarkerSymbol135 AreaMarkerSymbol = "135"
    AreaMarkerSymbolAsteriskOpen AreaMarkerSymbol = "asterisk-open"
    AreaMarkerSymbolNumber36 AreaMarkerSymbol = 36
    AreaMarkerSymbol36 AreaMarkerSymbol = "36"
    AreaMarkerSymbolHash AreaMarkerSymbol = "hash"
    AreaMarkerSymbolNumber136 AreaMarkerSymbol = 136
    AreaMarkerSymbol136 AreaMarkerSymbol = "136"
    AreaMarkerSymbolHashOpen AreaMarkerSymbol = "hash-open"
    AreaMarkerSymbolNumber236 AreaMarkerSymbol = 236
    AreaMarkerSymbol236 AreaMarkerSymbol = "236"
    AreaMarkerSymbolHashDot AreaMarkerSymbol = "hash-dot"
    AreaMarkerSymbolNumber336 AreaMarkerSymbol = 336
    AreaMarkerSymbol336 AreaMarkerSymbol = "336"
    AreaMarkerSymbolHashOpenDot AreaMarkerSymbol = "hash-open-dot"
    AreaMarkerSymbolNumber37 AreaMarkerSymbol = 37
    AreaMarkerSymbol37 AreaMarkerSymbol = "37"
    AreaMarkerSymbolYUp AreaMarkerSymbol = "y-up"
    AreaMarkerSymbolNumber137 AreaMarkerSymbol = 137
    AreaMarkerSymbol137 AreaMarkerSymbol = "137"
    AreaMarkerSymbolYUpOpen AreaMarkerSymbol = "y-up-open"
    AreaMarkerSymbolNumber38 AreaMarkerSymbol = 38
    AreaMarkerSymbol38 AreaMarkerSymbol = "38"
    AreaMarkerSymbolYDown AreaMarkerSymbol = "y-down"
    AreaMarkerSymbolNumber138 AreaMarkerSymbol = 138
    AreaMarkerSymbol138 AreaMarkerSymbol = "138"
    AreaMarkerSymbolYDownOpen AreaMarkerSymbol = "y-down-open"
    AreaMarkerSymbolNumber39 AreaMarkerSymbol = 39
    AreaMarkerSymbol39 AreaMarkerSymbol = "39"
    AreaMarkerSymbolYLeft AreaMarkerSymbol = "y-left"
    AreaMarkerSymbolNumber139 AreaMarkerSymbol = 139
    AreaMarkerSymbol139 AreaMarkerSymbol = "139"
    AreaMarkerSymbolYLeftOpen AreaMarkerSymbol = "y-left-open"
    AreaMarkerSymbolNumber40 AreaMarkerSymbol = 40
    AreaMarkerSymbol40 AreaMarkerSymbol = "40"
    AreaMarkerSymbolYRight AreaMarkerSymbol = "y-right"
    AreaMarkerSymbolNumber140 AreaMarkerSymbol = 140
    AreaMarkerSymbol140 AreaMarkerSymbol = "140"
    AreaMarkerSymbolYRightOpen AreaMarkerSymbol = "y-right-open"
    AreaMarkerSymbolNumber41 AreaMarkerSymbol = 41
    AreaMarkerSymbol41 AreaMarkerSymbol = "41"
    AreaMarkerSymbolLineEw AreaMarkerSymbol = "line-ew"
    AreaMarkerSymbolNumber141 AreaMarkerSymbol = 141
    AreaMarkerSymbol141 AreaMarkerSymbol = "141"
    AreaMarkerSymbolLineEwOpen AreaMarkerSymbol = "line-ew-open"
    AreaMarkerSymbolNumber42 AreaMarkerSymbol = 42
    AreaMarkerSymbol42 AreaMarkerSymbol = "42"
    AreaMarkerSymbolLineNs AreaMarkerSymbol = "line-ns"
    AreaMarkerSymbolNumber142 AreaMarkerSymbol = 142
    AreaMarkerSymbol142 AreaMarkerSymbol = "142"
    AreaMarkerSymbolLineNsOpen AreaMarkerSymbol = "line-ns-open"
    AreaMarkerSymbolNumber43 AreaMarkerSymbol = 43
    AreaMarkerSymbol43 AreaMarkerSymbol = "43"
    AreaMarkerSymbolLineNe AreaMarkerSymbol = "line-ne"
    AreaMarkerSymbolNumber143 AreaMarkerSymbol = 143
    AreaMarkerSymbol143 AreaMarkerSymbol = "143"
    AreaMarkerSymbolLineNeOpen AreaMarkerSymbol = "line-ne-open"
    AreaMarkerSymbolNumber44 AreaMarkerSymbol = 44
    AreaMarkerSymbol44 AreaMarkerSymbol = "44"
    AreaMarkerSymbolLineNw AreaMarkerSymbol = "line-nw"
    AreaMarkerSymbolNumber144 AreaMarkerSymbol = 144
    AreaMarkerSymbol144 AreaMarkerSymbol = "144"
    AreaMarkerSymbolLineNwOpen AreaMarkerSymbol = "line-nw-open"
    AreaMarkerSymbolNumber45 AreaMarkerSymbol = 45
    AreaMarkerSymbol45 AreaMarkerSymbol = "45"
    AreaMarkerSymbolArrowUp AreaMarkerSymbol = "arrow-up"
    AreaMarkerSymbolNumber145 AreaMarkerSymbol = 145
    AreaMarkerSymbol145 AreaMarkerSymbol = "145"
    AreaMarkerSymbolArrowUpOpen AreaMarkerSymbol = "arrow-up-open"
    AreaMarkerSymbolNumber46 AreaMarkerSymbol = 46
    AreaMarkerSymbol46 AreaMarkerSymbol = "46"
    AreaMarkerSymbolArrowDown AreaMarkerSymbol = "arrow-down"
    AreaMarkerSymbolNumber146 AreaMarkerSymbol = 146
    AreaMarkerSymbol146 AreaMarkerSymbol = "146"
    AreaMarkerSymbolArrowDownOpen AreaMarkerSymbol = "arrow-down-open"
    AreaMarkerSymbolNumber47 AreaMarkerSymbol = 47
    AreaMarkerSymbol47 AreaMarkerSymbol = "47"
    AreaMarkerSymbolArrowLeft AreaMarkerSymbol = "arrow-left"
    AreaMarkerSymbolNumber147 AreaMarkerSymbol = 147
    AreaMarkerSymbol147 AreaMarkerSymbol = "147"
    AreaMarkerSymbolArrowLeftOpen AreaMarkerSymbol = "arrow-left-open"
    AreaMarkerSymbolNumber48 AreaMarkerSymbol = 48
    AreaMarkerSymbol48 AreaMarkerSymbol = "48"
    AreaMarkerSymbolArrowRight AreaMarkerSymbol = "arrow-right"
    AreaMarkerSymbolNumber148 AreaMarkerSymbol = 148
    AreaMarkerSymbol148 AreaMarkerSymbol = "148"
    AreaMarkerSymbolArrowRightOpen AreaMarkerSymbol = "arrow-right-open"
    AreaMarkerSymbolNumber49 AreaMarkerSymbol = 49
    AreaMarkerSymbol49 AreaMarkerSymbol = "49"
    AreaMarkerSymbolArrowBarUp AreaMarkerSymbol = "arrow-bar-up"
    AreaMarkerSymbolNumber149 AreaMarkerSymbol = 149
    AreaMarkerSymbol149 AreaMarkerSymbol = "149"
    AreaMarkerSymbolArrowBarUpOpen AreaMarkerSymbol = "arrow-bar-up-open"
    AreaMarkerSymbolNumber50 AreaMarkerSymbol = 50
    AreaMarkerSymbol50 AreaMarkerSymbol = "50"
    AreaMarkerSymbolArrowBarDown AreaMarkerSymbol = "arrow-bar-down"
    AreaMarkerSymbolNumber150 AreaMarkerSymbol = 150
    AreaMarkerSymbol150 AreaMarkerSymbol = "150"
    AreaMarkerSymbolArrowBarDownOpen AreaMarkerSymbol = "arrow-bar-down-open"
    AreaMarkerSymbolNumber51 AreaMarkerSymbol = 51
    AreaMarkerSymbol51 AreaMarkerSymbol = "51"
    AreaMarkerSymbolArrowBarLeft AreaMarkerSymbol = "arrow-bar-left"
    AreaMarkerSymbolNumber151 AreaMarkerSymbol = 151
    AreaMarkerSymbol151 AreaMarkerSymbol = "151"
    AreaMarkerSymbolArrowBarLeftOpen AreaMarkerSymbol = "arrow-bar-left-open"
    AreaMarkerSymbolNumber52 AreaMarkerSymbol = 52
    AreaMarkerSymbol52 AreaMarkerSymbol = "52"
    AreaMarkerSymbolArrowBarRight AreaMarkerSymbol = "arrow-bar-right"
    AreaMarkerSymbolNumber152 AreaMarkerSymbol = 152
    AreaMarkerSymbol152 AreaMarkerSymbol = "152"
    AreaMarkerSymbolArrowBarRightOpen AreaMarkerSymbol = "arrow-bar-right-open"
    
)
// AreaVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type AreaVisible interface{} 

var (
    AreaVisibleTrue AreaVisible = true
    AreaVisibleFalse AreaVisible = false
    AreaVisibleLegendonly AreaVisible = "legendonly"
    
)
// AreaHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type AreaHoverinfo string 

const (
    // Flags
    AreaHoverinfoX AreaHoverinfo = "x"
    AreaHoverinfoY AreaHoverinfo = "y"
    AreaHoverinfoZ AreaHoverinfo = "z"
    AreaHoverinfoText AreaHoverinfo = "text"
    AreaHoverinfoName AreaHoverinfo = "name"
    
    // Extra
    AreaHoverinfoAll AreaHoverinfo = "all"
    AreaHoverinfoNone AreaHoverinfo = "none"
    AreaHoverinfoSkip AreaHoverinfo = "skip"
    
)
