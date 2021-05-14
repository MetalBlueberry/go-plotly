package grob

var TraceTypePointcloud TraceType = "pointcloud"

func (trace *Pointcloud) GetType() TraceType {
	return TraceTypePointcloud
}
// Pointcloud The data visualized as a point cloud set in `x` and `y` using the WebGl plotting engine.
type Pointcloud struct {
    
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
    Hoverinfo PointcloudHoverinfo `json:"hoverinfo,omitempty"`
    
    // Hoverinfosrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hoverinfo . 
    Hoverinfosrc String `json:"hoverinfosrc,omitempty"`
    
    // Hoverlabel 
    //   
    Hoverlabel *PointcloudHoverlabel `json:"hoverlabel,omitempty"`
    
    // Ids 
    // data_array 
    // Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type. 
    Ids interface{} `json:"ids,omitempty"`
    
    // Idssrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  ids . 
    Idssrc String `json:"idssrc,omitempty"`
    
    // Indices 
    // data_array 
    // A sequential value, 0..n, supply it to avoid creating this array inside plotting. If specified, it must be a typed `Int32Array` array. Its length must be equal to or greater than the number of points. For the best performance and memory use, create one large `indices` typed array that is guaranteed to be at least as long as the largest number of points during use, and reuse it on each `Plotly.restyle()` call. 
    Indices interface{} `json:"indices,omitempty"`
    
    // Indicessrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  indices . 
    Indicessrc String `json:"indicessrc,omitempty"`
    
    // Legendgroup 
    // string 
    // Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items. 
    Legendgroup String `json:"legendgroup,omitempty"`
    
    // Marker 
    //   
    Marker *PointcloudMarker `json:"marker,omitempty"`
    
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
    
    // Showlegend 
    // boolean 
    // Determines whether or not an item corresponding to this trace is shown in the legend. 
    Showlegend Bool `json:"showlegend,omitempty"`
    
    // Stream 
    //   
    Stream *PointcloudStream `json:"stream,omitempty"`
    
    // Text 
    // string 
    // Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels. 
    Text String `json:"text,omitempty"`
    
    // Textsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  text . 
    Textsrc String `json:"textsrc,omitempty"`
    
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
    Visible PointcloudVisible `json:"visible,omitempty"`
    
    // X 
    // data_array 
    // Sets the x coordinates. 
    X interface{} `json:"x,omitempty"`
    
    // Xaxis 
    // subplotid 
    // Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on. 
    Xaxis String `json:"xaxis,omitempty"`
    
    // Xbounds 
    // data_array 
    // Specify `xbounds` in the shape of `[xMin, xMax] to avoid looping through the `xy` typed array. Use it in conjunction with `xy` and `ybounds` for the performance benefits. 
    Xbounds interface{} `json:"xbounds,omitempty"`
    
    // Xboundssrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  xbounds . 
    Xboundssrc String `json:"xboundssrc,omitempty"`
    
    // Xsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  x . 
    Xsrc String `json:"xsrc,omitempty"`
    
    // Xy 
    // data_array 
    // Faster alternative to specifying `x` and `y` separately. If supplied, it must be a typed `Float32Array` array that represents points such that `xy[i * 2] = x[i]` and `xy[i * 2 + 1] = y[i]` 
    Xy interface{} `json:"xy,omitempty"`
    
    // Xysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  xy . 
    Xysrc String `json:"xysrc,omitempty"`
    
    // Y 
    // data_array 
    // Sets the y coordinates. 
    Y interface{} `json:"y,omitempty"`
    
    // Yaxis 
    // subplotid 
    // Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on. 
    Yaxis String `json:"yaxis,omitempty"`
    
    // Ybounds 
    // data_array 
    // Specify `ybounds` in the shape of `[yMin, yMax] to avoid looping through the `xy` typed array. Use it in conjunction with `xy` and `xbounds` for the performance benefits. 
    Ybounds interface{} `json:"ybounds,omitempty"`
    
    // Yboundssrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  ybounds . 
    Yboundssrc String `json:"yboundssrc,omitempty"`
    
    // Ysrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  y . 
    Ysrc String `json:"ysrc,omitempty"`
    
}
// PointcloudHoverlabelFont Sets the font used in hover labels.
type PointcloudHoverlabelFont struct {
    
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
// PointcloudHoverlabel 
type PointcloudHoverlabel struct {
    
    // Align 
    // enumerated Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines 
    Align PointcloudHoverlabelAlign `json:"align,omitempty"`
    
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
    Font *PointcloudHoverlabelFont `json:"font,omitempty"`
    
    // Namelength 
    // integer 
    // Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis. 
    Namelength int64 `json:"namelength,omitempty"`
    
    // Namelengthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  namelength . 
    Namelengthsrc String `json:"namelengthsrc,omitempty"`
    
}
// PointcloudMarkerBorder 
type PointcloudMarkerBorder struct {
    
    // Arearatio 
    // number 
    // Specifies what fraction of the marker area is covered with the border. 
    Arearatio float64 `json:"arearatio,omitempty"`
    
    // Color 
    // color 
    // Sets the stroke color. It accepts a specific color. If the color is not fully opaque and there are hundreds of thousands of points, it may cause slower zooming and panning. 
    Color Color `json:"color,omitempty"`
    
}
// PointcloudMarker 
type PointcloudMarker struct {
    
    // Blend 
    // boolean 
    // Determines if colors are blended together for a translucency effect in case `opacity` is specified as a value less then `1`. Setting `blend` to `true` reduces zoom/pan speed if used with large numbers of points. 
    Blend Bool `json:"blend,omitempty"`
    
    // Border 
    //   
    Border *PointcloudMarkerBorder `json:"border,omitempty"`
    
    // Color 
    // color 
    // Sets the marker fill color. It accepts a specific color.If the color is not fully opaque and there are hundreds of thousandsof points, it may cause slower zooming and panning. 
    Color Color `json:"color,omitempty"`
    
    // Opacity 
    // number 
    // Sets the marker opacity. The default value is `1` (fully opaque). If the markers are not fully opaque and there are hundreds of thousands of points, it may cause slower zooming and panning. Opacity fades the color even if `blend` is left on `false` even if there is no translucency effect in that case. 
    Opacity float64 `json:"opacity,omitempty"`
    
    // Sizemax 
    // number 
    // Sets the maximum size (in px) of the rendered marker points. Effective when the `pointcloud` shows only few points. 
    Sizemax float64 `json:"sizemax,omitempty"`
    
    // Sizemin 
    // number 
    // Sets the minimum size (in px) of the rendered marker points, effective when the `pointcloud` shows a million or more points. 
    Sizemin float64 `json:"sizemin,omitempty"`
    
}
// PointcloudStream 
type PointcloudStream struct {
    
    // Maxpoints 
    // number 
    // Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot. 
    Maxpoints float64 `json:"maxpoints,omitempty"`
    
    // Token 
    // string 
    // The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details. 
    Token String `json:"token,omitempty"`
    
}
// PointcloudHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type PointcloudHoverlabelAlign string 

const (
    PointcloudHoverlabelAlignLeft PointcloudHoverlabelAlign = "left"
    PointcloudHoverlabelAlignRight PointcloudHoverlabelAlign = "right"
    PointcloudHoverlabelAlignAuto PointcloudHoverlabelAlign = "auto"
    
)
// PointcloudVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type PointcloudVisible interface{} 

var (
    PointcloudVisibleTrue PointcloudVisible = true
    PointcloudVisibleFalse PointcloudVisible = false
    PointcloudVisibleLegendonly PointcloudVisible = "legendonly"
    
)
// PointcloudHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type PointcloudHoverinfo string 

const (
    // Flags
    PointcloudHoverinfoX PointcloudHoverinfo = "x"
    PointcloudHoverinfoY PointcloudHoverinfo = "y"
    PointcloudHoverinfoZ PointcloudHoverinfo = "z"
    PointcloudHoverinfoText PointcloudHoverinfo = "text"
    PointcloudHoverinfoName PointcloudHoverinfo = "name"
    
    // Extra
    PointcloudHoverinfoAll PointcloudHoverinfo = "all"
    PointcloudHoverinfoNone PointcloudHoverinfo = "none"
    PointcloudHoverinfoSkip PointcloudHoverinfo = "skip"
    
)
