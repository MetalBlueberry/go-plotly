package grob

var TraceTypeFunnelarea TraceType = "funnelarea"

func (trace *Funnelarea) GetType() TraceType {
	return TraceTypeFunnelarea
}
// Funnelarea Visualize stages in a process using area-encoded trapezoids. This trace can be used to show data in a part-to-whole representation similar to a "pie" trace, wherein each item appears in a single stage. See also the "funnel" trace type for a different approach to visualizing funnel data.
type Funnelarea struct {
    
    // Type 
    // is the type of the plot 
    Type TraceType `json:"type,omitempty"`
    
    // Aspectratio 
    // number 
    // Sets the ratio between height and width 
    Aspectratio float64 `json:"aspectratio,omitempty"`
    
    // Baseratio 
    // number 
    // Sets the ratio between bottom length and maximum top length. 
    Baseratio float64 `json:"baseratio,omitempty"`
    
    // Customdata 
    // data_array 
    // Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements 
    Customdata interface{} `json:"customdata,omitempty"`
    
    // Customdatasrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  customdata . 
    Customdatasrc String `json:"customdatasrc,omitempty"`
    
    // Dlabel 
    // number 
    // Sets the label step. See `label0` for more info. 
    Dlabel float64 `json:"dlabel,omitempty"`
    
    // Domain 
    //   
    Domain *FunnelareaDomain `json:"domain,omitempty"`
    
    // Hoverinfo 
    // flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired. 
    Hoverinfo FunnelareaHoverinfo `json:"hoverinfo,omitempty"`
    
    // Hoverinfosrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hoverinfo . 
    Hoverinfosrc String `json:"hoverinfosrc,omitempty"`
    
    // Hoverlabel 
    //   
    Hoverlabel *FunnelareaHoverlabel `json:"hoverlabel,omitempty"`
    
    // Hovertemplate 
    // string 
    // Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `label`, `color`, `value`, `text` and `percent`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`. 
    Hovertemplate String `json:"hovertemplate,omitempty"`
    
    // Hovertemplatesrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hovertemplate . 
    Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`
    
    // Hovertext 
    // string 
    // Sets hover text elements associated with each sector. If a single string, the same string appears for all data points. If an array of string, the items are mapped in order of this trace's sectors. To be seen, trace `hoverinfo` must contain a *text* flag. 
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
    
    // Insidetextfont 
    //  Sets the font used for `textinfo` lying inside the sector. 
    Insidetextfont *FunnelareaInsidetextfont `json:"insidetextfont,omitempty"`
    
    // Label0 
    // number 
    // Alternate to `labels`. Builds a numeric set of labels. Use with `dlabel` where `label0` is the starting label and `dlabel` the step. 
    Label0 float64 `json:"label0,omitempty"`
    
    // Labels 
    // data_array 
    // Sets the sector labels. If `labels` entries are duplicated, we sum associated `values` or simply count occurrences if `values` is not provided. For other array attributes (including color) we use the first non-empty entry among all occurrences of the label. 
    Labels interface{} `json:"labels,omitempty"`
    
    // Labelssrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  labels . 
    Labelssrc String `json:"labelssrc,omitempty"`
    
    // Legendgroup 
    // string 
    // Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items. 
    Legendgroup String `json:"legendgroup,omitempty"`
    
    // Marker 
    //   
    Marker *FunnelareaMarker `json:"marker,omitempty"`
    
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
    
    // Scalegroup 
    // string 
    // If there are multiple funnelareas that should be sized according to their totals, link them by providing a non-empty group id here shared by every trace in the same group. 
    Scalegroup String `json:"scalegroup,omitempty"`
    
    // Showlegend 
    // boolean 
    // Determines whether or not an item corresponding to this trace is shown in the legend. 
    Showlegend Bool `json:"showlegend,omitempty"`
    
    // Stream 
    //   
    Stream *FunnelareaStream `json:"stream,omitempty"`
    
    // Text 
    // data_array 
    // Sets text elements associated with each sector. If trace `textinfo` contains a *text* flag, these elements will be seen on the chart. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels. 
    Text interface{} `json:"text,omitempty"`
    
    // Textfont 
    //  Sets the font used for `textinfo`. 
    Textfont *FunnelareaTextfont `json:"textfont,omitempty"`
    
    // Textinfo 
    // flaglist Determines which trace information appear on the graph. 
    Textinfo FunnelareaTextinfo `json:"textinfo,omitempty"`
    
    // Textposition 
    // enumerated Specifies the location of the `textinfo`. 
    Textposition FunnelareaTextposition `json:"textposition,omitempty"`
    
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
    // Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `label`, `color`, `value`, `text` and `percent`. 
    Texttemplate String `json:"texttemplate,omitempty"`
    
    // Texttemplatesrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  texttemplate . 
    Texttemplatesrc String `json:"texttemplatesrc,omitempty"`
    
    // Title 
    //   
    Title *FunnelareaTitle `json:"title,omitempty"`
    
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
    
    // Values 
    // data_array 
    // Sets the values of the sectors. If omitted, we count occurrences of each label. 
    Values interface{} `json:"values,omitempty"`
    
    // Valuessrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  values . 
    Valuessrc String `json:"valuessrc,omitempty"`
    
    // Visible 
    // enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible). 
    Visible FunnelareaVisible `json:"visible,omitempty"`
    
}
// FunnelareaDomain 
type FunnelareaDomain struct {
    
    // Column 
    // integer 
    // If there is a layout grid, use the domain for this column in the grid for this funnelarea trace . 
    Column int64 `json:"column,omitempty"`
    
    // Row 
    // integer 
    // If there is a layout grid, use the domain for this row in the grid for this funnelarea trace . 
    Row int64 `json:"row,omitempty"`
    
    // X 
    // info_array 
    // Sets the horizontal domain of this funnelarea trace (in plot fraction). 
    X interface{} `json:"x,omitempty"`
    
    // Y 
    // info_array 
    // Sets the vertical domain of this funnelarea trace (in plot fraction). 
    Y interface{} `json:"y,omitempty"`
    
}
// FunnelareaHoverlabelFont Sets the font used in hover labels.
type FunnelareaHoverlabelFont struct {
    
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
// FunnelareaHoverlabel 
type FunnelareaHoverlabel struct {
    
    // Align 
    // enumerated Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines 
    Align FunnelareaHoverlabelAlign `json:"align,omitempty"`
    
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
    Font *FunnelareaHoverlabelFont `json:"font,omitempty"`
    
    // Namelength 
    // integer 
    // Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis. 
    Namelength int64 `json:"namelength,omitempty"`
    
    // Namelengthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  namelength . 
    Namelengthsrc String `json:"namelengthsrc,omitempty"`
    
}
// FunnelareaInsidetextfont Sets the font used for `textinfo` lying inside the sector.
type FunnelareaInsidetextfont struct {
    
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
// FunnelareaMarkerLine 
type FunnelareaMarkerLine struct {
    
    // Color 
    // color 
    // Sets the color of the line enclosing each sector. Defaults to the `paper_bgcolor` value. 
    Color Color `json:"color,omitempty"`
    
    // Colorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  color . 
    Colorsrc String `json:"colorsrc,omitempty"`
    
    // Width 
    // number 
    // Sets the width (in px) of the line enclosing each sector. 
    Width float64 `json:"width,omitempty"`
    
    // Widthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  width . 
    Widthsrc String `json:"widthsrc,omitempty"`
    
}
// FunnelareaMarker 
type FunnelareaMarker struct {
    
    // Colors 
    // data_array 
    // Sets the color of each sector. If not specified, the default trace color set is used to pick the sector colors. 
    Colors interface{} `json:"colors,omitempty"`
    
    // Colorssrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  colors . 
    Colorssrc String `json:"colorssrc,omitempty"`
    
    // Line 
    //   
    Line *FunnelareaMarkerLine `json:"line,omitempty"`
    
}
// FunnelareaStream 
type FunnelareaStream struct {
    
    // Maxpoints 
    // number 
    // Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot. 
    Maxpoints float64 `json:"maxpoints,omitempty"`
    
    // Token 
    // string 
    // The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details. 
    Token String `json:"token,omitempty"`
    
}
// FunnelareaTextfont Sets the font used for `textinfo`.
type FunnelareaTextfont struct {
    
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
// FunnelareaTitleFont Sets the font used for `title`. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type FunnelareaTitleFont struct {
    
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
// FunnelareaTitle 
type FunnelareaTitle struct {
    
    // Font 
    //  Sets the font used for `title`. Note that the title's font used to be set by the now deprecated `titlefont` attribute. 
    Font *FunnelareaTitleFont `json:"font,omitempty"`
    
    // Position 
    // enumerated Specifies the location of the `title`. Note that the title's position used to be set by the now deprecated `titleposition` attribute. 
    Position FunnelareaTitlePosition `json:"position,omitempty"`
    
    // Text 
    // string 
    // Sets the title of the chart. If it is empty, no title is displayed. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
}
// FunnelareaHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type FunnelareaHoverlabelAlign string 

const (
    FunnelareaHoverlabelAlignLeft FunnelareaHoverlabelAlign = "left"
    FunnelareaHoverlabelAlignRight FunnelareaHoverlabelAlign = "right"
    FunnelareaHoverlabelAlignAuto FunnelareaHoverlabelAlign = "auto"
    
)
// FunnelareaTextposition Specifies the location of the `textinfo`.
type FunnelareaTextposition string 

const (
    FunnelareaTextpositionInside FunnelareaTextposition = "inside"
    FunnelareaTextpositionNone FunnelareaTextposition = "none"
    
)
// FunnelareaTitlePosition Specifies the location of the `title`. Note that the title's position used to be set by the now deprecated `titleposition` attribute.
type FunnelareaTitlePosition string 

const (
    FunnelareaTitlePositionTopLeft FunnelareaTitlePosition = "top left"
    FunnelareaTitlePositionTopCenter FunnelareaTitlePosition = "top center"
    FunnelareaTitlePositionTopRight FunnelareaTitlePosition = "top right"
    
)
// FunnelareaVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type FunnelareaVisible interface{} 

var (
    FunnelareaVisibleTrue FunnelareaVisible = true
    FunnelareaVisibleFalse FunnelareaVisible = false
    FunnelareaVisibleLegendonly FunnelareaVisible = "legendonly"
    
)
// FunnelareaHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type FunnelareaHoverinfo string 

const (
    // Flags
    FunnelareaHoverinfoLabel FunnelareaHoverinfo = "label"
    FunnelareaHoverinfoText FunnelareaHoverinfo = "text"
    FunnelareaHoverinfoValue FunnelareaHoverinfo = "value"
    FunnelareaHoverinfoPercent FunnelareaHoverinfo = "percent"
    FunnelareaHoverinfoName FunnelareaHoverinfo = "name"
    
    // Extra
    FunnelareaHoverinfoAll FunnelareaHoverinfo = "all"
    FunnelareaHoverinfoNone FunnelareaHoverinfo = "none"
    FunnelareaHoverinfoSkip FunnelareaHoverinfo = "skip"
    
)
// FunnelareaTextinfo Determines which trace information appear on the graph.
type FunnelareaTextinfo string 

const (
    // Flags
    FunnelareaTextinfoLabel FunnelareaTextinfo = "label"
    FunnelareaTextinfoText FunnelareaTextinfo = "text"
    FunnelareaTextinfoValue FunnelareaTextinfo = "value"
    FunnelareaTextinfoPercent FunnelareaTextinfo = "percent"
    
    // Extra
    FunnelareaTextinfoNone FunnelareaTextinfo = "none"
    
)
