package grob

var TraceTypeBar TraceType = "bar"

func (trace *Bar) GetType() TraceType {
	return TraceTypeBar
}
// Bar The data visualized by the span of the bars is set in `y` if `orientation` is set th *v* (the default) and the labels are set in `x`. By setting `orientation` to *h*, the roles are interchanged.
type Bar struct {
    
    // Type 
    // is the type of the plot 
    Type TraceType `json:"type,omitempty"`
    
    // Alignmentgroup 
    // string 
    // Set several traces linked to the same position axis or matching axes to the same alignmentgroup. This controls whether bars compute their positional range dependently or independently. 
    Alignmentgroup String `json:"alignmentgroup,omitempty"`
    
    // Base 
    // any 
    // Sets where the bar base is drawn (in position axis units). In *stack* or *relative* barmode, traces that set *base* will be excluded and drawn in *overlay* mode instead. 
    Base interface{} `json:"base,omitempty"`
    
    // Basesrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  base . 
    Basesrc String `json:"basesrc,omitempty"`
    
    // Cliponaxis 
    // boolean 
    // Determines whether the text nodes are clipped about the subplot axes. To show the text nodes above axis lines and tick labels, make sure to set `xaxis.layer` and `yaxis.layer` to *below traces*. 
    Cliponaxis Bool `json:"cliponaxis,omitempty"`
    
    // Constraintext 
    // enumerated Constrain the size of text inside or outside a bar to be no larger than the bar itself. 
    Constraintext BarConstraintext `json:"constraintext,omitempty"`
    
    // Customdata 
    // data_array 
    // Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements 
    Customdata interface{} `json:"customdata,omitempty"`
    
    // Customdatasrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  customdata . 
    Customdatasrc String `json:"customdatasrc,omitempty"`
    
    // Dx 
    // number 
    // Sets the x coordinate step. See `x0` for more info. 
    Dx float64 `json:"dx,omitempty"`
    
    // Dy 
    // number 
    // Sets the y coordinate step. See `y0` for more info. 
    Dy float64 `json:"dy,omitempty"`
    
    // ErrorX 
    //   
    ErrorX *BarErrorX `json:"error_x,omitempty"`
    
    // ErrorY 
    //   
    ErrorY *BarErrorY `json:"error_y,omitempty"`
    
    // Hoverinfo 
    // flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired. 
    Hoverinfo BarHoverinfo `json:"hoverinfo,omitempty"`
    
    // Hoverinfosrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hoverinfo . 
    Hoverinfosrc String `json:"hoverinfosrc,omitempty"`
    
    // Hoverlabel 
    //   
    Hoverlabel *BarHoverlabel `json:"hoverlabel,omitempty"`
    
    // Hovertemplate 
    // string 
    // Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `value` and `label`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`. 
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
    
    // Insidetextanchor 
    // enumerated Determines if texts are kept at center or start/end points in `textposition` *inside* mode. 
    Insidetextanchor BarInsidetextanchor `json:"insidetextanchor,omitempty"`
    
    // Insidetextfont 
    //  Sets the font used for `text` lying inside the bar. 
    Insidetextfont *BarInsidetextfont `json:"insidetextfont,omitempty"`
    
    // Legendgroup 
    // string 
    // Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items. 
    Legendgroup String `json:"legendgroup,omitempty"`
    
    // Marker 
    //   
    Marker *BarMarker `json:"marker,omitempty"`
    
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
    Orientation BarOrientation `json:"orientation,omitempty"`
    
    // Outsidetextfont 
    //  Sets the font used for `text` lying outside the bar. 
    Outsidetextfont *BarOutsidetextfont `json:"outsidetextfont,omitempty"`
    
    // R 
    // data_array 
    // r coordinates in scatter traces are deprecated!Please switch to the *scatterpolar* trace type.Sets the radial coordinatesfor legacy polar chart only. 
    R interface{} `json:"r,omitempty"`
    
    // Rsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  r . 
    Rsrc String `json:"rsrc,omitempty"`
    
    // Selected 
    //   
    Selected *BarSelected `json:"selected,omitempty"`
    
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
    Stream *BarStream `json:"stream,omitempty"`
    
    // T 
    // data_array 
    // t coordinates in scatter traces are deprecated!Please switch to the *scatterpolar* trace type.Sets the angular coordinatesfor legacy polar chart only. 
    T interface{} `json:"t,omitempty"`
    
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
    Textfont *BarTextfont `json:"textfont,omitempty"`
    
    // Textposition 
    // enumerated Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside. 
    Textposition BarTextposition `json:"textposition,omitempty"`
    
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
    // Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `value` and `label`. 
    Texttemplate String `json:"texttemplate,omitempty"`
    
    // Texttemplatesrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  texttemplate . 
    Texttemplatesrc String `json:"texttemplatesrc,omitempty"`
    
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
    
    // Unselected 
    //   
    Unselected *BarUnselected `json:"unselected,omitempty"`
    
    // Visible 
    // enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible). 
    Visible BarVisible `json:"visible,omitempty"`
    
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
    
    // Xcalendar 
    // enumerated Sets the calendar system to use with `x` date data. 
    Xcalendar BarXcalendar `json:"xcalendar,omitempty"`
    
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
    Xperiodalignment BarXperiodalignment `json:"xperiodalignment,omitempty"`
    
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
    
    // Ycalendar 
    // enumerated Sets the calendar system to use with `y` date data. 
    Ycalendar BarYcalendar `json:"ycalendar,omitempty"`
    
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
    Yperiodalignment BarYperiodalignment `json:"yperiodalignment,omitempty"`
    
    // Ysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  y . 
    Ysrc String `json:"ysrc,omitempty"`
    
}
// BarErrorX 
type BarErrorX struct {
    
    // Array 
    // data_array 
    // Sets the data corresponding the length of each error bar. Values are plotted relative to the underlying data. 
    Array interface{} `json:"array,omitempty"`
    
    // Arrayminus 
    // data_array 
    // Sets the data corresponding the length of each error bar in the bottom (left) direction for vertical (horizontal) bars Values are plotted relative to the underlying data. 
    Arrayminus interface{} `json:"arrayminus,omitempty"`
    
    // Arrayminussrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  arrayminus . 
    Arrayminussrc String `json:"arrayminussrc,omitempty"`
    
    // Arraysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  array . 
    Arraysrc String `json:"arraysrc,omitempty"`
    
    // Color 
    // color 
    // Sets the stoke color of the error bars. 
    Color Color `json:"color,omitempty"`
    
    // CopyYstyle 
    // boolean 
    //  
    CopyYstyle Bool `json:"copy_ystyle,omitempty"`
    
    // Symmetric 
    // boolean 
    // Determines whether or not the error bars have the same length in both direction (top/bottom for vertical bars, left/right for horizontal bars. 
    Symmetric Bool `json:"symmetric,omitempty"`
    
    // Thickness 
    // number 
    // Sets the thickness (in px) of the error bars. 
    Thickness float64 `json:"thickness,omitempty"`
    
    // Traceref 
    // integer 
    //  
    Traceref int64 `json:"traceref,omitempty"`
    
    // Tracerefminus 
    // integer 
    //  
    Tracerefminus int64 `json:"tracerefminus,omitempty"`
    
    // Type 
    // enumerated Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the square of the underlying data. If *data*, the bar lengths are set with data set `array`. 
    Type BarErrorXType `json:"type,omitempty"`
    
    // Value 
    // number 
    // Sets the value of either the percentage (if `type` is set to *percent*) or the constant (if `type` is set to *constant*) corresponding to the lengths of the error bars. 
    Value float64 `json:"value,omitempty"`
    
    // Valueminus 
    // number 
    // Sets the value of either the percentage (if `type` is set to *percent*) or the constant (if `type` is set to *constant*) corresponding to the lengths of the error bars in the bottom (left) direction for vertical (horizontal) bars 
    Valueminus float64 `json:"valueminus,omitempty"`
    
    // Visible 
    // boolean 
    // Determines whether or not this set of error bars is visible. 
    Visible Bool `json:"visible,omitempty"`
    
    // Width 
    // number 
    // Sets the width (in px) of the cross-bar at both ends of the error bars. 
    Width float64 `json:"width,omitempty"`
    
}
// BarErrorY 
type BarErrorY struct {
    
    // Array 
    // data_array 
    // Sets the data corresponding the length of each error bar. Values are plotted relative to the underlying data. 
    Array interface{} `json:"array,omitempty"`
    
    // Arrayminus 
    // data_array 
    // Sets the data corresponding the length of each error bar in the bottom (left) direction for vertical (horizontal) bars Values are plotted relative to the underlying data. 
    Arrayminus interface{} `json:"arrayminus,omitempty"`
    
    // Arrayminussrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  arrayminus . 
    Arrayminussrc String `json:"arrayminussrc,omitempty"`
    
    // Arraysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  array . 
    Arraysrc String `json:"arraysrc,omitempty"`
    
    // Color 
    // color 
    // Sets the stoke color of the error bars. 
    Color Color `json:"color,omitempty"`
    
    // Symmetric 
    // boolean 
    // Determines whether or not the error bars have the same length in both direction (top/bottom for vertical bars, left/right for horizontal bars. 
    Symmetric Bool `json:"symmetric,omitempty"`
    
    // Thickness 
    // number 
    // Sets the thickness (in px) of the error bars. 
    Thickness float64 `json:"thickness,omitempty"`
    
    // Traceref 
    // integer 
    //  
    Traceref int64 `json:"traceref,omitempty"`
    
    // Tracerefminus 
    // integer 
    //  
    Tracerefminus int64 `json:"tracerefminus,omitempty"`
    
    // Type 
    // enumerated Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the square of the underlying data. If *data*, the bar lengths are set with data set `array`. 
    Type BarErrorYType `json:"type,omitempty"`
    
    // Value 
    // number 
    // Sets the value of either the percentage (if `type` is set to *percent*) or the constant (if `type` is set to *constant*) corresponding to the lengths of the error bars. 
    Value float64 `json:"value,omitempty"`
    
    // Valueminus 
    // number 
    // Sets the value of either the percentage (if `type` is set to *percent*) or the constant (if `type` is set to *constant*) corresponding to the lengths of the error bars in the bottom (left) direction for vertical (horizontal) bars 
    Valueminus float64 `json:"valueminus,omitempty"`
    
    // Visible 
    // boolean 
    // Determines whether or not this set of error bars is visible. 
    Visible Bool `json:"visible,omitempty"`
    
    // Width 
    // number 
    // Sets the width (in px) of the cross-bar at both ends of the error bars. 
    Width float64 `json:"width,omitempty"`
    
}
// BarHoverlabelFont Sets the font used in hover labels.
type BarHoverlabelFont struct {
    
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
// BarHoverlabel 
type BarHoverlabel struct {
    
    // Align 
    // enumerated Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines 
    Align BarHoverlabelAlign `json:"align,omitempty"`
    
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
    Font *BarHoverlabelFont `json:"font,omitempty"`
    
    // Namelength 
    // integer 
    // Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis. 
    Namelength int64 `json:"namelength,omitempty"`
    
    // Namelengthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  namelength . 
    Namelengthsrc String `json:"namelengthsrc,omitempty"`
    
}
// BarInsidetextfont Sets the font used for `text` lying inside the bar.
type BarInsidetextfont struct {
    
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
// BarMarkerColorbarTickfont Sets the color bar's tick label font
type BarMarkerColorbarTickfont struct {
    
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
// BarMarkerColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type BarMarkerColorbarTitleFont struct {
    
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
// BarMarkerColorbarTitle 
type BarMarkerColorbarTitle struct {
    
    // Font 
    //  Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute. 
    Font *BarMarkerColorbarTitleFont `json:"font,omitempty"`
    
    // Side 
    // enumerated Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute. 
    Side BarMarkerColorbarTitleSide `json:"side,omitempty"`
    
    // Text 
    // string 
    // Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
}
// BarMarkerColorbar 
type BarMarkerColorbar struct {
    
    // Bgcolor 
    // color 
    // Sets the color of padded area. 
    Bgcolor Color `json:"bgcolor,omitempty"`
    
    // Bordercolor 
    // color 
    // Sets the axis line color. 
    Bordercolor Color `json:"bordercolor,omitempty"`
    
    // Borderwidth 
    // number 
    // Sets the width (in px) or the border enclosing this color bar. 
    Borderwidth float64 `json:"borderwidth,omitempty"`
    
    // Dtick 
    // any 
    // Sets the step in-between ticks on this axis. Use with `tick0`. Must be a positive number, or special strings available to *log* and *date* axes. If the axis `type` is *log*, then ticks are set every 10^(n*dtick) where n is the tick number. For example, to set a tick mark at 1, 10, 100, 1000, ... set dtick to 1. To set tick marks at 1, 100, 10000, ... set dtick to 2. To set tick marks at 1, 5, 25, 125, 625, 3125, ... set dtick to log_10(5), or 0.69897000433. *log* has several special values; *L<f>*, where `f` is a positive number, gives ticks linearly spaced in value (but not position). For example `tick0` = 0.1, `dtick` = *L0.5* will put ticks at 0.1, 0.6, 1.1, 1.6 etc. To show powers of 10 plus small digits between, use *D1* (all digits) or *D2* (only 2 and 5). `tick0` is ignored for *D1* and *D2*. If the axis `type` is *date*, then you must convert the time to milliseconds. For example, to set the interval between ticks to one day, set `dtick` to 86400000.0. *date* also has special values *M<n>* gives ticks spaced by a number of months. `n` must be a positive integer. To set ticks on the 15th of every third month, set `tick0` to *2000-01-15* and `dtick` to *M3*. To set ticks every 4 years, set `dtick` to *M48* 
    Dtick interface{} `json:"dtick,omitempty"`
    
    // Exponentformat 
    // enumerated Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B. 
    Exponentformat BarMarkerColorbarExponentformat `json:"exponentformat,omitempty"`
    
    // Len 
    // number 
    // Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends. 
    Len float64 `json:"len,omitempty"`
    
    // Lenmode 
    // enumerated Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value. 
    Lenmode BarMarkerColorbarLenmode `json:"lenmode,omitempty"`
    
    // Minexponent 
    // number 
    // Hide SI prefix for 10^n if |n| is below this number. This only has an effect when `tickformat` is *SI* or *B*. 
    Minexponent float64 `json:"minexponent,omitempty"`
    
    // Nticks 
    // integer 
    // Specifies the maximum number of ticks for the particular axis. The actual number of ticks will be chosen automatically to be less than or equal to `nticks`. Has an effect only if `tickmode` is set to *auto*. 
    Nticks int64 `json:"nticks,omitempty"`
    
    // Outlinecolor 
    // color 
    // Sets the axis line color. 
    Outlinecolor Color `json:"outlinecolor,omitempty"`
    
    // Outlinewidth 
    // number 
    // Sets the width (in px) of the axis line. 
    Outlinewidth float64 `json:"outlinewidth,omitempty"`
    
    // Separatethousands 
    // boolean 
    // If "true", even 4-digit integers are separated 
    Separatethousands Bool `json:"separatethousands,omitempty"`
    
    // Showexponent 
    // enumerated If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear. 
    Showexponent BarMarkerColorbarShowexponent `json:"showexponent,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix BarMarkerColorbarShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix BarMarkerColorbarShowticksuffix `json:"showticksuffix,omitempty"`
    
    // Thickness 
    // number 
    // Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels. 
    Thickness float64 `json:"thickness,omitempty"`
    
    // Thicknessmode 
    // enumerated Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value. 
    Thicknessmode BarMarkerColorbarThicknessmode `json:"thicknessmode,omitempty"`
    
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
    Tickfont *BarMarkerColorbarTickfont `json:"tickfont,omitempty"`
    
    // Tickformat 
    // string 
    // Sets the tick label formatting rule using d3 formatting mini-languages which are very similar to those in Python. For numbers, see: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format And for dates see: https://github.com/d3/d3-time-format#locale_format We add one item to d3's date formatter: *%{n}f* for fractional seconds with n digits. For example, *2016-10-13 09:15:23.456* with tickformat *%H~%M~%S.%2f* would display *09~15~23.46* 
    Tickformat String `json:"tickformat,omitempty"`
    
    // Tickformatstops 
    // It's an items array and what goes inside it's... messy... check the docs 
    // I will be happy if you want to contribute by implementing this 
    // just raise an issue before you start so we do not overlap 
    Tickformatstops interface{} `json:"tickformatstops,omitempty"`
    
    // Ticklabelposition 
    // enumerated Determines where tick labels are drawn. 
    Ticklabelposition BarMarkerColorbarTicklabelposition `json:"ticklabelposition,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode BarMarkerColorbarTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks BarMarkerColorbarTicks `json:"ticks,omitempty"`
    
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
    
    // Title 
    //   
    Title *BarMarkerColorbarTitle `json:"title,omitempty"`
    
    // X 
    // number 
    // Sets the x position of the color bar (in plot fraction). 
    X float64 `json:"x,omitempty"`
    
    // Xanchor 
    // enumerated Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar. 
    Xanchor BarMarkerColorbarXanchor `json:"xanchor,omitempty"`
    
    // Xpad 
    // number 
    // Sets the amount of padding (in px) along the x direction. 
    Xpad float64 `json:"xpad,omitempty"`
    
    // Y 
    // number 
    // Sets the y position of the color bar (in plot fraction). 
    Y float64 `json:"y,omitempty"`
    
    // Yanchor 
    // enumerated Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar. 
    Yanchor BarMarkerColorbarYanchor `json:"yanchor,omitempty"`
    
    // Ypad 
    // number 
    // Sets the amount of padding (in px) along the y direction. 
    Ypad float64 `json:"ypad,omitempty"`
    
}
// BarMarkerLine 
type BarMarkerLine struct {
    
    // Autocolorscale 
    // boolean 
    // Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `marker.line.colorscale`. Has an effect only if in `marker.line.color`is set to a numerical array. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed. 
    Autocolorscale Bool `json:"autocolorscale,omitempty"`
    
    // Cauto 
    // boolean 
    // Determines whether or not the color domain is computed with respect to the input data (here in `marker.line.color`) or the bounds set in `marker.line.cmin` and `marker.line.cmax`  Has an effect only if in `marker.line.color`is set to a numerical array. Defaults to `false` when `marker.line.cmin` and `marker.line.cmax` are set by the user. 
    Cauto Bool `json:"cauto,omitempty"`
    
    // Cmax 
    // number 
    // Sets the upper bound of the color domain. Has an effect only if in `marker.line.color`is set to a numerical array. Value should have the same units as in `marker.line.color` and if set, `marker.line.cmin` must be set as well. 
    Cmax float64 `json:"cmax,omitempty"`
    
    // Cmid 
    // number 
    // Sets the mid-point of the color domain by scaling `marker.line.cmin` and/or `marker.line.cmax` to be equidistant to this point. Has an effect only if in `marker.line.color`is set to a numerical array. Value should have the same units as in `marker.line.color`. Has no effect when `marker.line.cauto` is `false`. 
    Cmid float64 `json:"cmid,omitempty"`
    
    // Cmin 
    // number 
    // Sets the lower bound of the color domain. Has an effect only if in `marker.line.color`is set to a numerical array. Value should have the same units as in `marker.line.color` and if set, `marker.line.cmax` must be set as well. 
    Cmin float64 `json:"cmin,omitempty"`
    
    // Color 
    // color 
    // Sets themarker.linecolor. It accepts either a specific color or an array of numbers that are mapped to the colorscale relative to the max and min values of the array or relative to `marker.line.cmin` and `marker.line.cmax` if set. 
    Color Color `json:"color,omitempty"`
    
    // Coloraxis 
    // subplotid 
    // Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis. 
    Coloraxis String `json:"coloraxis,omitempty"`
    
    // Colorscale 
    // colorscale 
    // Sets the colorscale. Has an effect only if in `marker.line.color`is set to a numerical array. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`marker.line.cmin` and `marker.line.cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis. 
    // Default: <nil> 
    Colorscale ColorScale `json:"colorscale,omitempty"`
    
    // Colorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  color . 
    Colorsrc String `json:"colorsrc,omitempty"`
    
    // Reversescale 
    // boolean 
    // Reverses the color mapping if true. Has an effect only if in `marker.line.color`is set to a numerical array. If true, `marker.line.cmin` will correspond to the last color in the array and `marker.line.cmax` will correspond to the first color. 
    Reversescale Bool `json:"reversescale,omitempty"`
    
    // Width 
    // number 
    // Sets the width (in px) of the lines bounding the marker points. 
    Width float64 `json:"width,omitempty"`
    
    // Widthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  width . 
    Widthsrc String `json:"widthsrc,omitempty"`
    
}
// BarMarker 
type BarMarker struct {
    
    // Autocolorscale 
    // boolean 
    // Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `marker.colorscale`. Has an effect only if in `marker.color`is set to a numerical array. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed. 
    Autocolorscale Bool `json:"autocolorscale,omitempty"`
    
    // Cauto 
    // boolean 
    // Determines whether or not the color domain is computed with respect to the input data (here in `marker.color`) or the bounds set in `marker.cmin` and `marker.cmax`  Has an effect only if in `marker.color`is set to a numerical array. Defaults to `false` when `marker.cmin` and `marker.cmax` are set by the user. 
    Cauto Bool `json:"cauto,omitempty"`
    
    // Cmax 
    // number 
    // Sets the upper bound of the color domain. Has an effect only if in `marker.color`is set to a numerical array. Value should have the same units as in `marker.color` and if set, `marker.cmin` must be set as well. 
    Cmax float64 `json:"cmax,omitempty"`
    
    // Cmid 
    // number 
    // Sets the mid-point of the color domain by scaling `marker.cmin` and/or `marker.cmax` to be equidistant to this point. Has an effect only if in `marker.color`is set to a numerical array. Value should have the same units as in `marker.color`. Has no effect when `marker.cauto` is `false`. 
    Cmid float64 `json:"cmid,omitempty"`
    
    // Cmin 
    // number 
    // Sets the lower bound of the color domain. Has an effect only if in `marker.color`is set to a numerical array. Value should have the same units as in `marker.color` and if set, `marker.cmax` must be set as well. 
    Cmin float64 `json:"cmin,omitempty"`
    
    // Color 
    // color 
    // Sets themarkercolor. It accepts either a specific color or an array of numbers that are mapped to the colorscale relative to the max and min values of the array or relative to `marker.cmin` and `marker.cmax` if set. 
    Color Color `json:"color,omitempty"`
    
    // Coloraxis 
    // subplotid 
    // Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis. 
    Coloraxis String `json:"coloraxis,omitempty"`
    
    // Colorbar 
    //   
    Colorbar *BarMarkerColorbar `json:"colorbar,omitempty"`
    
    // Colorscale 
    // colorscale 
    // Sets the colorscale. Has an effect only if in `marker.color`is set to a numerical array. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`marker.cmin` and `marker.cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis. 
    // Default: <nil> 
    Colorscale ColorScale `json:"colorscale,omitempty"`
    
    // Colorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  color . 
    Colorsrc String `json:"colorsrc,omitempty"`
    
    // Line 
    //   
    Line *BarMarkerLine `json:"line,omitempty"`
    
    // Opacity 
    // number 
    // Sets the opacity of the bars. 
    Opacity float64 `json:"opacity,omitempty"`
    
    // Opacitysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  opacity . 
    Opacitysrc String `json:"opacitysrc,omitempty"`
    
    // Reversescale 
    // boolean 
    // Reverses the color mapping if true. Has an effect only if in `marker.color`is set to a numerical array. If true, `marker.cmin` will correspond to the last color in the array and `marker.cmax` will correspond to the first color. 
    Reversescale Bool `json:"reversescale,omitempty"`
    
    // Showscale 
    // boolean 
    // Determines whether or not a colorbar is displayed for this trace. Has an effect only if in `marker.color`is set to a numerical array. 
    Showscale Bool `json:"showscale,omitempty"`
    
}
// BarOutsidetextfont Sets the font used for `text` lying outside the bar.
type BarOutsidetextfont struct {
    
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
// BarSelectedMarker 
type BarSelectedMarker struct {
    
    // Color 
    // color 
    // Sets the marker color of selected points. 
    Color Color `json:"color,omitempty"`
    
    // Opacity 
    // number 
    // Sets the marker opacity of selected points. 
    Opacity float64 `json:"opacity,omitempty"`
    
}
// BarSelectedTextfont 
type BarSelectedTextfont struct {
    
    // Color 
    // color 
    // Sets the text font color of selected points. 
    Color Color `json:"color,omitempty"`
    
}
// BarSelected 
type BarSelected struct {
    
    // Marker 
    //   
    Marker *BarSelectedMarker `json:"marker,omitempty"`
    
    // Textfont 
    //   
    Textfont *BarSelectedTextfont `json:"textfont,omitempty"`
    
}
// BarStream 
type BarStream struct {
    
    // Maxpoints 
    // number 
    // Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot. 
    Maxpoints float64 `json:"maxpoints,omitempty"`
    
    // Token 
    // string 
    // The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details. 
    Token String `json:"token,omitempty"`
    
}
// BarTextfont Sets the font used for `text`.
type BarTextfont struct {
    
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
// BarUnselectedMarker 
type BarUnselectedMarker struct {
    
    // Color 
    // color 
    // Sets the marker color of unselected points, applied only when a selection exists. 
    Color Color `json:"color,omitempty"`
    
    // Opacity 
    // number 
    // Sets the marker opacity of unselected points, applied only when a selection exists. 
    Opacity float64 `json:"opacity,omitempty"`
    
}
// BarUnselectedTextfont 
type BarUnselectedTextfont struct {
    
    // Color 
    // color 
    // Sets the text font color of unselected points, applied only when a selection exists. 
    Color Color `json:"color,omitempty"`
    
}
// BarUnselected 
type BarUnselected struct {
    
    // Marker 
    //   
    Marker *BarUnselectedMarker `json:"marker,omitempty"`
    
    // Textfont 
    //   
    Textfont *BarUnselectedTextfont `json:"textfont,omitempty"`
    
}
// BarConstraintext Constrain the size of text inside or outside a bar to be no larger than the bar itself.
type BarConstraintext string 

const (
    BarConstraintextInside BarConstraintext = "inside"
    BarConstraintextOutside BarConstraintext = "outside"
    BarConstraintextBoth BarConstraintext = "both"
    BarConstraintextNone BarConstraintext = "none"
    
)
// BarErrorXType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the square of the underlying data. If *data*, the bar lengths are set with data set `array`.
type BarErrorXType string 

const (
    BarErrorXTypePercent BarErrorXType = "percent"
    BarErrorXTypeConstant BarErrorXType = "constant"
    BarErrorXTypeSqrt BarErrorXType = "sqrt"
    BarErrorXTypeData BarErrorXType = "data"
    
)
// BarErrorYType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the square of the underlying data. If *data*, the bar lengths are set with data set `array`.
type BarErrorYType string 

const (
    BarErrorYTypePercent BarErrorYType = "percent"
    BarErrorYTypeConstant BarErrorYType = "constant"
    BarErrorYTypeSqrt BarErrorYType = "sqrt"
    BarErrorYTypeData BarErrorYType = "data"
    
)
// BarHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type BarHoverlabelAlign string 

const (
    BarHoverlabelAlignLeft BarHoverlabelAlign = "left"
    BarHoverlabelAlignRight BarHoverlabelAlign = "right"
    BarHoverlabelAlignAuto BarHoverlabelAlign = "auto"
    
)
// BarInsidetextanchor Determines if texts are kept at center or start/end points in `textposition` *inside* mode.
type BarInsidetextanchor string 

const (
    BarInsidetextanchorEnd BarInsidetextanchor = "end"
    BarInsidetextanchorMiddle BarInsidetextanchor = "middle"
    BarInsidetextanchorStart BarInsidetextanchor = "start"
    
)
// BarMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type BarMarkerColorbarExponentformat string 

const (
    BarMarkerColorbarExponentformatNone BarMarkerColorbarExponentformat = "none"
    BarMarkerColorbarExponentformatE1 BarMarkerColorbarExponentformat = "e"
    BarMarkerColorbarExponentformatE2 BarMarkerColorbarExponentformat = "E"
    BarMarkerColorbarExponentformatPower BarMarkerColorbarExponentformat = "power"
    BarMarkerColorbarExponentformatSi BarMarkerColorbarExponentformat = "SI"
    BarMarkerColorbarExponentformatB BarMarkerColorbarExponentformat = "B"
    
)
// BarMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type BarMarkerColorbarLenmode string 

const (
    BarMarkerColorbarLenmodeFraction BarMarkerColorbarLenmode = "fraction"
    BarMarkerColorbarLenmodePixels BarMarkerColorbarLenmode = "pixels"
    
)
// BarMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type BarMarkerColorbarShowexponent string 

const (
    BarMarkerColorbarShowexponentAll BarMarkerColorbarShowexponent = "all"
    BarMarkerColorbarShowexponentFirst BarMarkerColorbarShowexponent = "first"
    BarMarkerColorbarShowexponentLast BarMarkerColorbarShowexponent = "last"
    BarMarkerColorbarShowexponentNone BarMarkerColorbarShowexponent = "none"
    
)
// BarMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type BarMarkerColorbarShowtickprefix string 

const (
    BarMarkerColorbarShowtickprefixAll BarMarkerColorbarShowtickprefix = "all"
    BarMarkerColorbarShowtickprefixFirst BarMarkerColorbarShowtickprefix = "first"
    BarMarkerColorbarShowtickprefixLast BarMarkerColorbarShowtickprefix = "last"
    BarMarkerColorbarShowtickprefixNone BarMarkerColorbarShowtickprefix = "none"
    
)
// BarMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type BarMarkerColorbarShowticksuffix string 

const (
    BarMarkerColorbarShowticksuffixAll BarMarkerColorbarShowticksuffix = "all"
    BarMarkerColorbarShowticksuffixFirst BarMarkerColorbarShowticksuffix = "first"
    BarMarkerColorbarShowticksuffixLast BarMarkerColorbarShowticksuffix = "last"
    BarMarkerColorbarShowticksuffixNone BarMarkerColorbarShowticksuffix = "none"
    
)
// BarMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type BarMarkerColorbarThicknessmode string 

const (
    BarMarkerColorbarThicknessmodeFraction BarMarkerColorbarThicknessmode = "fraction"
    BarMarkerColorbarThicknessmodePixels BarMarkerColorbarThicknessmode = "pixels"
    
)
// BarMarkerColorbarTicklabelposition Determines where tick labels are drawn.
type BarMarkerColorbarTicklabelposition string 

const (
    BarMarkerColorbarTicklabelpositionOutside BarMarkerColorbarTicklabelposition = "outside"
    BarMarkerColorbarTicklabelpositionInside BarMarkerColorbarTicklabelposition = "inside"
    BarMarkerColorbarTicklabelpositionOutsideTop BarMarkerColorbarTicklabelposition = "outside top"
    BarMarkerColorbarTicklabelpositionInsideTop BarMarkerColorbarTicklabelposition = "inside top"
    BarMarkerColorbarTicklabelpositionOutsideBottom BarMarkerColorbarTicklabelposition = "outside bottom"
    BarMarkerColorbarTicklabelpositionInsideBottom BarMarkerColorbarTicklabelposition = "inside bottom"
    
)
// BarMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type BarMarkerColorbarTickmode string 

const (
    BarMarkerColorbarTickmodeAuto BarMarkerColorbarTickmode = "auto"
    BarMarkerColorbarTickmodeLinear BarMarkerColorbarTickmode = "linear"
    BarMarkerColorbarTickmodeArray BarMarkerColorbarTickmode = "array"
    
)
// BarMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type BarMarkerColorbarTicks string 

const (
    BarMarkerColorbarTicksOutside BarMarkerColorbarTicks = "outside"
    BarMarkerColorbarTicksInside BarMarkerColorbarTicks = "inside"
    BarMarkerColorbarTicksEmpty BarMarkerColorbarTicks = ""
    
)
// BarMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type BarMarkerColorbarTitleSide string 

const (
    BarMarkerColorbarTitleSideRight BarMarkerColorbarTitleSide = "right"
    BarMarkerColorbarTitleSideTop BarMarkerColorbarTitleSide = "top"
    BarMarkerColorbarTitleSideBottom BarMarkerColorbarTitleSide = "bottom"
    
)
// BarMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type BarMarkerColorbarXanchor string 

const (
    BarMarkerColorbarXanchorLeft BarMarkerColorbarXanchor = "left"
    BarMarkerColorbarXanchorCenter BarMarkerColorbarXanchor = "center"
    BarMarkerColorbarXanchorRight BarMarkerColorbarXanchor = "right"
    
)
// BarMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type BarMarkerColorbarYanchor string 

const (
    BarMarkerColorbarYanchorTop BarMarkerColorbarYanchor = "top"
    BarMarkerColorbarYanchorMiddle BarMarkerColorbarYanchor = "middle"
    BarMarkerColorbarYanchorBottom BarMarkerColorbarYanchor = "bottom"
    
)
// BarOrientation Sets the orientation of the bars. With *v* (*h*), the value of the each bar spans along the vertical (horizontal).
type BarOrientation string 

const (
    BarOrientationV BarOrientation = "v"
    BarOrientationH BarOrientation = "h"
    
)
// BarTextposition Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside.
type BarTextposition string 

const (
    BarTextpositionInside BarTextposition = "inside"
    BarTextpositionOutside BarTextposition = "outside"
    BarTextpositionAuto BarTextposition = "auto"
    BarTextpositionNone BarTextposition = "none"
    
)
// BarVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type BarVisible interface{} 

var (
    BarVisibleTrue BarVisible = true
    BarVisibleFalse BarVisible = false
    BarVisibleLegendonly BarVisible = "legendonly"
    
)
// BarXcalendar Sets the calendar system to use with `x` date data.
type BarXcalendar string 

const (
    BarXcalendarGregorian BarXcalendar = "gregorian"
    BarXcalendarChinese BarXcalendar = "chinese"
    BarXcalendarCoptic BarXcalendar = "coptic"
    BarXcalendarDiscworld BarXcalendar = "discworld"
    BarXcalendarEthiopian BarXcalendar = "ethiopian"
    BarXcalendarHebrew BarXcalendar = "hebrew"
    BarXcalendarIslamic BarXcalendar = "islamic"
    BarXcalendarJulian BarXcalendar = "julian"
    BarXcalendarMayan BarXcalendar = "mayan"
    BarXcalendarNanakshahi BarXcalendar = "nanakshahi"
    BarXcalendarNepali BarXcalendar = "nepali"
    BarXcalendarPersian BarXcalendar = "persian"
    BarXcalendarJalali BarXcalendar = "jalali"
    BarXcalendarTaiwan BarXcalendar = "taiwan"
    BarXcalendarThai BarXcalendar = "thai"
    BarXcalendarUmmalqura BarXcalendar = "ummalqura"
    
)
// BarXperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
type BarXperiodalignment string 

const (
    BarXperiodalignmentStart BarXperiodalignment = "start"
    BarXperiodalignmentMiddle BarXperiodalignment = "middle"
    BarXperiodalignmentEnd BarXperiodalignment = "end"
    
)
// BarYcalendar Sets the calendar system to use with `y` date data.
type BarYcalendar string 

const (
    BarYcalendarGregorian BarYcalendar = "gregorian"
    BarYcalendarChinese BarYcalendar = "chinese"
    BarYcalendarCoptic BarYcalendar = "coptic"
    BarYcalendarDiscworld BarYcalendar = "discworld"
    BarYcalendarEthiopian BarYcalendar = "ethiopian"
    BarYcalendarHebrew BarYcalendar = "hebrew"
    BarYcalendarIslamic BarYcalendar = "islamic"
    BarYcalendarJulian BarYcalendar = "julian"
    BarYcalendarMayan BarYcalendar = "mayan"
    BarYcalendarNanakshahi BarYcalendar = "nanakshahi"
    BarYcalendarNepali BarYcalendar = "nepali"
    BarYcalendarPersian BarYcalendar = "persian"
    BarYcalendarJalali BarYcalendar = "jalali"
    BarYcalendarTaiwan BarYcalendar = "taiwan"
    BarYcalendarThai BarYcalendar = "thai"
    BarYcalendarUmmalqura BarYcalendar = "ummalqura"
    
)
// BarYperiodalignment Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
type BarYperiodalignment string 

const (
    BarYperiodalignmentStart BarYperiodalignment = "start"
    BarYperiodalignmentMiddle BarYperiodalignment = "middle"
    BarYperiodalignmentEnd BarYperiodalignment = "end"
    
)
// BarHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type BarHoverinfo string 

const (
    // Flags
    BarHoverinfoX BarHoverinfo = "x"
    BarHoverinfoY BarHoverinfo = "y"
    BarHoverinfoZ BarHoverinfo = "z"
    BarHoverinfoText BarHoverinfo = "text"
    BarHoverinfoName BarHoverinfo = "name"
    
    // Extra
    BarHoverinfoAll BarHoverinfo = "all"
    BarHoverinfoNone BarHoverinfo = "none"
    BarHoverinfoSkip BarHoverinfo = "skip"
    
)
