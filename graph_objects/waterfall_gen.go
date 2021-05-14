package grob

var TraceTypeWaterfall TraceType = "waterfall"

func (trace *Waterfall) GetType() TraceType {
	return TraceTypeWaterfall
}
// Waterfall Draws waterfall trace which is useful graph to displays the contribution of various elements (either positive or negative) in a bar chart. The data visualized by the span of the bars is set in `y` if `orientation` is set th *v* (the default) and the labels are set in `x`. By setting `orientation` to *h*, the roles are interchanged.
type Waterfall struct {
    
    // Type 
    // is the type of the plot 
    Type TraceType `json:"type,omitempty"`
    
    // Alignmentgroup 
    // string 
    // Set several traces linked to the same position axis or matching axes to the same alignmentgroup. This controls whether bars compute their positional range dependently or independently. 
    Alignmentgroup String `json:"alignmentgroup,omitempty"`
    
    // Base 
    // number 
    // Sets where the bar base is drawn (in position axis units). 
    Base float64 `json:"base,omitempty"`
    
    // Cliponaxis 
    // boolean 
    // Determines whether the text nodes are clipped about the subplot axes. To show the text nodes above axis lines and tick labels, make sure to set `xaxis.layer` and `yaxis.layer` to *below traces*. 
    Cliponaxis Bool `json:"cliponaxis,omitempty"`
    
    // Connector 
    //   
    Connector *WaterfallConnector `json:"connector,omitempty"`
    
    // Constraintext 
    // enumerated Constrain the size of text inside or outside a bar to be no larger than the bar itself. 
    Constraintext WaterfallConstraintext `json:"constraintext,omitempty"`
    
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
    Decreasing *WaterfallDecreasing `json:"decreasing,omitempty"`
    
    // Dx 
    // number 
    // Sets the x coordinate step. See `x0` for more info. 
    Dx float64 `json:"dx,omitempty"`
    
    // Dy 
    // number 
    // Sets the y coordinate step. See `y0` for more info. 
    Dy float64 `json:"dy,omitempty"`
    
    // Hoverinfo 
    // flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired. 
    Hoverinfo WaterfallHoverinfo `json:"hoverinfo,omitempty"`
    
    // Hoverinfosrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hoverinfo . 
    Hoverinfosrc String `json:"hoverinfosrc,omitempty"`
    
    // Hoverlabel 
    //   
    Hoverlabel *WaterfallHoverlabel `json:"hoverlabel,omitempty"`
    
    // Hovertemplate 
    // string 
    // Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `initial`, `delta` and `final`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`. 
    Hovertemplate String `json:"hovertemplate,omitempty"`
    
    // Hovertemplatesrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hovertemplate . 
    Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`
    
    // Hovertext 
    // string 
    // Sets hover text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag. 
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
    Increasing *WaterfallIncreasing `json:"increasing,omitempty"`
    
    // Insidetextanchor 
    // enumerated Determines if texts are kept at center or start/end points in `textposition` *inside* mode. 
    Insidetextanchor WaterfallInsidetextanchor `json:"insidetextanchor,omitempty"`
    
    // Insidetextfont 
    //  Sets the font used for `text` lying inside the bar. 
    Insidetextfont *WaterfallInsidetextfont `json:"insidetextfont,omitempty"`
    
    // Legendgroup 
    // string 
    // Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items. 
    Legendgroup String `json:"legendgroup,omitempty"`
    
    // Measure 
    // data_array 
    // An array containing types of values. By default the values are considered as 'relative'. However; it is possible to use 'total' to compute the sums. Also 'absolute' could be applied to reset the computed total or to declare an initial value where needed. 
    Measure interface{} `json:"measure,omitempty"`
    
    // Measuresrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  measure . 
    Measuresrc String `json:"measuresrc,omitempty"`
    
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
    
    // Offset 
    // number 
    // Shifts the position where the bar is drawn (in position axis units). In *group* barmode, traces that set *offset* will be excluded and drawn in *overlay* mode instead. 
    Offset float64 `json:"offset,omitempty"`
    
    // Offsetgroup 
    // string 
    // Set several traces linked to the same position axis or matching axes to the same offsetgroup where bars of the same position coordinate will line up. 
    Offsetgroup String `json:"offsetgroup,omitempty"`
    
    // Offsetsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  offset . 
    Offsetsrc String `json:"offsetsrc,omitempty"`
    
    // Opacity 
    // number 
    // Sets the opacity of the trace. 
    Opacity float64 `json:"opacity,omitempty"`
    
    // Orientation 
    // enumerated Sets the orientation of the bars. With *v* (*h*), the value of the each bar spans along the vertical (horizontal). 
    Orientation WaterfallOrientation `json:"orientation,omitempty"`
    
    // Outsidetextfont 
    //  Sets the font used for `text` lying outside the bar. 
    Outsidetextfont *WaterfallOutsidetextfont `json:"outsidetextfont,omitempty"`
    
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
    Stream *WaterfallStream `json:"stream,omitempty"`
    
    // Text 
    // string 
    // Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels. 
    Text String `json:"text,omitempty"`
    
    // Textangle 
    // angle 
    // Sets the angle of the tick labels with respect to the bar. For example, a `tickangle` of -90 draws the tick labels vertically. With *auto* the texts may automatically be rotated to fit with the maximum size in bars. 
    Textangle float64 `json:"textangle,omitempty"`
    
    // Textfont 
    //  Sets the font used for `text`. 
    Textfont *WaterfallTextfont `json:"textfont,omitempty"`
    
    // Textinfo 
    // flaglist Determines which trace information appear on the graph. In the case of having multiple waterfalls, totals are computed separately (per trace). 
    Textinfo WaterfallTextinfo `json:"textinfo,omitempty"`
    
    // Textposition 
    // enumerated Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside. 
    Textposition WaterfallTextposition `json:"textposition,omitempty"`
    
    // Textpositionsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  textposition . 
    Textpositionsrc String `json:"textpositionsrc,omitempty"`
    
    // Textsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  text . 
    Textsrc String `json:"textsrc,omitempty"`
    
    // Texttemplate 
    // string 
    // Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `initial`, `delta`, `final` and `label`. 
    Texttemplate String `json:"texttemplate,omitempty"`
    
    // Texttemplatesrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  texttemplate . 
    Texttemplatesrc String `json:"texttemplatesrc,omitempty"`
    
    // Totals 
    //   
    Totals *WaterfallTotals `json:"totals,omitempty"`
    
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
    Visible WaterfallVisible `json:"visible,omitempty"`
    
    // Width 
    // number 
    // Sets the bar width (in position axis units). 
    Width float64 `json:"width,omitempty"`
    
    // Widthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  width . 
    Widthsrc String `json:"widthsrc,omitempty"`
    
    // X 
    // data_array 
    // Sets the x coordinates. 
    X interface{} `json:"x,omitempty"`
    
    // X0 
    // any 
    // Alternate to `x`. Builds a linear space of x coordinates. Use with `dx` where `x0` is the starting coordinate and `dx` the step. 
    X0 interface{} `json:"x0,omitempty"`
    
    // Xaxis 
    // subplotid 
    // Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on. 
    Xaxis String `json:"xaxis,omitempty"`
    
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
    Xperiodalignment WaterfallXperiodalignment `json:"xperiodalignment,omitempty"`
    
    // Xsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  x . 
    Xsrc String `json:"xsrc,omitempty"`
    
    // Y 
    // data_array 
    // Sets the y coordinates. 
    Y interface{} `json:"y,omitempty"`
    
    // Y0 
    // any 
    // Alternate to `y`. Builds a linear space of y coordinates. Use with `dy` where `y0` is the starting coordinate and `dy` the step. 
    Y0 interface{} `json:"y0,omitempty"`
    
    // Yaxis 
    // subplotid 
    // Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on. 
    Yaxis String `json:"yaxis,omitempty"`
    
    // Yperiod 
    // any 
    // Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the y axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer. 
    Yperiod interface{} `json:"yperiod,omitempty"`
    
    // Yperiod0 
    // any 
    // Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the y0 axis. When `y0period` is round number of weeks, the `y0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01. 
    Yperiod0 interface{} `json:"yperiod0,omitempty"`
    
    // Yperiodalignment 
    // enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis. 
    Yperiodalignment WaterfallYperiodalignment `json:"yperiodalignment,omitempty"`
    
    // Ysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  y . 
    Ysrc String `json:"ysrc,omitempty"`
    
}
// WaterfallConnectorLine 
type WaterfallConnectorLine struct {
    
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
// WaterfallConnector 
type WaterfallConnector struct {
    
    // Line 
    //   
    Line *WaterfallConnectorLine `json:"line,omitempty"`
    
    // Mode 
    // enumerated Sets the shape of connector lines. 
    Mode WaterfallConnectorMode `json:"mode,omitempty"`
    
    // Visible 
    // boolean 
    // Determines if connector lines are drawn.  
    Visible Bool `json:"visible,omitempty"`
    
}
// WaterfallDecreasingMarkerLine 
type WaterfallDecreasingMarkerLine struct {
    
    // Color 
    // color 
    // Sets the line color of all decreasing values. 
    Color Color `json:"color,omitempty"`
    
    // Width 
    // number 
    // Sets the line width of all decreasing values. 
    Width float64 `json:"width,omitempty"`
    
}
// WaterfallDecreasingMarker 
type WaterfallDecreasingMarker struct {
    
    // Color 
    // color 
    // Sets the marker color of all decreasing values. 
    Color Color `json:"color,omitempty"`
    
    // Line 
    //   
    Line *WaterfallDecreasingMarkerLine `json:"line,omitempty"`
    
}
// WaterfallDecreasing 
type WaterfallDecreasing struct {
    
    // Marker 
    //   
    Marker *WaterfallDecreasingMarker `json:"marker,omitempty"`
    
}
// WaterfallHoverlabelFont Sets the font used in hover labels.
type WaterfallHoverlabelFont struct {
    
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
// WaterfallHoverlabel 
type WaterfallHoverlabel struct {
    
    // Align 
    // enumerated Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines 
    Align WaterfallHoverlabelAlign `json:"align,omitempty"`
    
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
    Font *WaterfallHoverlabelFont `json:"font,omitempty"`
    
    // Namelength 
    // integer 
    // Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis. 
    Namelength int64 `json:"namelength,omitempty"`
    
    // Namelengthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  namelength . 
    Namelengthsrc String `json:"namelengthsrc,omitempty"`
    
}
// WaterfallIncreasingMarkerLine 
type WaterfallIncreasingMarkerLine struct {
    
    // Color 
    // color 
    // Sets the line color of all increasing values. 
    Color Color `json:"color,omitempty"`
    
    // Width 
    // number 
    // Sets the line width of all increasing values. 
    Width float64 `json:"width,omitempty"`
    
}
// WaterfallIncreasingMarker 
type WaterfallIncreasingMarker struct {
    
    // Color 
    // color 
    // Sets the marker color of all increasing values. 
    Color Color `json:"color,omitempty"`
    
    // Line 
    //   
    Line *WaterfallIncreasingMarkerLine `json:"line,omitempty"`
    
}
// WaterfallIncreasing 
type WaterfallIncreasing struct {
    
    // Marker 
    //   
    Marker *WaterfallIncreasingMarker `json:"marker,omitempty"`
    
}
// WaterfallInsidetextfont Sets the font used for `text` lying inside the bar.
type WaterfallInsidetextfont struct {
    
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
// WaterfallOutsidetextfont Sets the font used for `text` lying outside the bar.
type WaterfallOutsidetextfont struct {
    
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
// WaterfallStream 
type WaterfallStream struct {
    
    // Maxpoints 
    // number 
    // Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot. 
    Maxpoints float64 `json:"maxpoints,omitempty"`
    
    // Token 
    // string 
    // The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details. 
    Token String `json:"token,omitempty"`
    
}
// WaterfallTextfont Sets the font used for `text`.
type WaterfallTextfont struct {
    
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
// WaterfallTotalsMarkerLine 
type WaterfallTotalsMarkerLine struct {
    
    // Color 
    // color 
    // Sets the line color of all intermediate sums and total values. 
    Color Color `json:"color,omitempty"`
    
    // Width 
    // number 
    // Sets the line width of all intermediate sums and total values. 
    Width float64 `json:"width,omitempty"`
    
}
// WaterfallTotalsMarker 
type WaterfallTotalsMarker struct {
    
    // Color 
    // color 
    // Sets the marker color of all intermediate sums and total values. 
    Color Color `json:"color,omitempty"`
    
    // Line 
    //   
    Line *WaterfallTotalsMarkerLine `json:"line,omitempty"`
    
}
// WaterfallTotals 
type WaterfallTotals struct {
    
    // Marker 
    //   
    Marker *WaterfallTotalsMarker `json:"marker,omitempty"`
    
}
// WaterfallConnectorMode Sets the shape of connector lines.
type WaterfallConnectorMode string 

const (
    WaterfallConnectorModeSpanning WaterfallConnectorMode = "spanning"
    WaterfallConnectorModeBetween WaterfallConnectorMode = "between"
    
)
// WaterfallConstraintext Constrain the size of text inside or outside a bar to be no larger than the bar itself.
type WaterfallConstraintext string 

const (
    WaterfallConstraintextInside WaterfallConstraintext = "inside"
    WaterfallConstraintextOutside WaterfallConstraintext = "outside"
    WaterfallConstraintextBoth WaterfallConstraintext = "both"
    WaterfallConstraintextNone WaterfallConstraintext = "none"
    
)
// WaterfallHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type WaterfallHoverlabelAlign string 

const (
    WaterfallHoverlabelAlignLeft WaterfallHoverlabelAlign = "left"
    WaterfallHoverlabelAlignRight WaterfallHoverlabelAlign = "right"
    WaterfallHoverlabelAlignAuto WaterfallHoverlabelAlign = "auto"
    
)
// WaterfallInsidetextanchor Determines if texts are kept at center or start/end points in `textposition` *inside* mode.
type WaterfallInsidetextanchor string 

const (
    WaterfallInsidetextanchorEnd WaterfallInsidetextanchor = "end"
    WaterfallInsidetextanchorMiddle WaterfallInsidetextanchor = "middle"
    WaterfallInsidetextanchorStart WaterfallInsidetextanchor = "start"
    
)
// WaterfallOrientation Sets the orientation of the bars. With *v* (*h*), the value of the each bar spans along the vertical (horizontal).
type WaterfallOrientation string 

const (
    WaterfallOrientationV WaterfallOrientation = "v"
    WaterfallOrientationH WaterfallOrientation = "h"
    
)
// WaterfallTextposition Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside.
type WaterfallTextposition string 

const (
    WaterfallTextpositionInside WaterfallTextposition = "inside"
    WaterfallTextpositionOutside WaterfallTextposition = "outside"
    WaterfallTextpositionAuto WaterfallTextposition = "auto"
    WaterfallTextpositionNone WaterfallTextposition = "none"
    
)
// WaterfallVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type WaterfallVisible interface{} 

var (
    WaterfallVisibleTrue WaterfallVisible = true
    WaterfallVisibleFalse WaterfallVisible = false
    WaterfallVisibleLegendonly WaterfallVisible = "legendonly"
    
)
// WaterfallXperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
type WaterfallXperiodalignment string 

const (
    WaterfallXperiodalignmentStart WaterfallXperiodalignment = "start"
    WaterfallXperiodalignmentMiddle WaterfallXperiodalignment = "middle"
    WaterfallXperiodalignmentEnd WaterfallXperiodalignment = "end"
    
)
// WaterfallYperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
type WaterfallYperiodalignment string 

const (
    WaterfallYperiodalignmentStart WaterfallYperiodalignment = "start"
    WaterfallYperiodalignmentMiddle WaterfallYperiodalignment = "middle"
    WaterfallYperiodalignmentEnd WaterfallYperiodalignment = "end"
    
)
// WaterfallHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type WaterfallHoverinfo string 

const (
    // Flags
    WaterfallHoverinfoName WaterfallHoverinfo = "name"
    WaterfallHoverinfoX WaterfallHoverinfo = "x"
    WaterfallHoverinfoY WaterfallHoverinfo = "y"
    WaterfallHoverinfoText WaterfallHoverinfo = "text"
    WaterfallHoverinfoInitial WaterfallHoverinfo = "initial"
    WaterfallHoverinfoDelta WaterfallHoverinfo = "delta"
    WaterfallHoverinfoFinal WaterfallHoverinfo = "final"
    
    // Extra
    WaterfallHoverinfoAll WaterfallHoverinfo = "all"
    WaterfallHoverinfoNone WaterfallHoverinfo = "none"
    WaterfallHoverinfoSkip WaterfallHoverinfo = "skip"
    
)
// WaterfallTextinfo Determines which trace information appear on the graph. In the case of having multiple waterfalls, totals are computed separately (per trace).
type WaterfallTextinfo string 

const (
    // Flags
    WaterfallTextinfoLabel WaterfallTextinfo = "label"
    WaterfallTextinfoText WaterfallTextinfo = "text"
    WaterfallTextinfoInitial WaterfallTextinfo = "initial"
    WaterfallTextinfoDelta WaterfallTextinfo = "delta"
    WaterfallTextinfoFinal WaterfallTextinfo = "final"
    
    // Extra
    WaterfallTextinfoNone WaterfallTextinfo = "none"
    
)
