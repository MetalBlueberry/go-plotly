package grob

var TraceTypeScattergeo TraceType = "scattergeo"

func (trace *Scattergeo) GetType() TraceType {
	return TraceTypeScattergeo
}
// Scattergeo The data visualized as scatter point or lines on a geographic map is provided either by longitude/latitude pairs in `lon` and `lat` respectively or by geographic location IDs or names in `locations`.
type Scattergeo struct {
    
    // Type 
    // is the type of the plot 
    Type TraceType `json:"type,omitempty"`
    
    // Connectgaps 
    // boolean 
    // Determines whether or not gaps (i.e. {nan} or missing values) in the provided data arrays are connected. 
    Connectgaps Bool `json:"connectgaps,omitempty"`
    
    // Customdata 
    // data_array 
    // Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements 
    Customdata interface{} `json:"customdata,omitempty"`
    
    // Customdatasrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  customdata . 
    Customdatasrc String `json:"customdatasrc,omitempty"`
    
    // Featureidkey 
    // string 
    // Sets the key in GeoJSON features which is used as id to match the items included in the `locations` array. Only has an effect when `geojson` is set. Support nested property, for example *properties.name*. 
    Featureidkey String `json:"featureidkey,omitempty"`
    
    // Fill 
    // enumerated Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. 
    Fill ScattergeoFill `json:"fill,omitempty"`
    
    // Fillcolor 
    // color 
    // Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available. 
    Fillcolor Color `json:"fillcolor,omitempty"`
    
    // Geo 
    // subplotid 
    // Sets a reference between this trace's geospatial coordinates and a geographic map. If *geo* (the default value), the geospatial coordinates refer to `layout.geo`. If *geo2*, the geospatial coordinates refer to `layout.geo2`, and so on. 
    Geo String `json:"geo,omitempty"`
    
    // Geojson 
    // any 
    // Sets optional GeoJSON data associated with this trace. If not given, the features on the base map are used when `locations` is set. It can be set as a valid GeoJSON object or as a URL string. Note that we only accept GeoJSONs of type *FeatureCollection* or *Feature* with geometries of type *Polygon* or *MultiPolygon*. 
    Geojson interface{} `json:"geojson,omitempty"`
    
    // Hoverinfo 
    // flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired. 
    Hoverinfo ScattergeoHoverinfo `json:"hoverinfo,omitempty"`
    
    // Hoverinfosrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hoverinfo . 
    Hoverinfosrc String `json:"hoverinfosrc,omitempty"`
    
    // Hoverlabel 
    //   
    Hoverlabel *ScattergeoHoverlabel `json:"hoverlabel,omitempty"`
    
    // Hovertemplate 
    // string 
    // Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`. 
    Hovertemplate String `json:"hovertemplate,omitempty"`
    
    // Hovertemplatesrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hovertemplate . 
    Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`
    
    // Hovertext 
    // string 
    // Sets hover text elements associated with each (lon,lat) pair or item in `locations`. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (lon,lat) or `locations` coordinates. To be seen, trace `hoverinfo` must contain a *text* flag. 
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
    
    // Lat 
    // data_array 
    // Sets the latitude coordinates (in degrees North). 
    Lat interface{} `json:"lat,omitempty"`
    
    // Latsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  lat . 
    Latsrc String `json:"latsrc,omitempty"`
    
    // Legendgroup 
    // string 
    // Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items. 
    Legendgroup String `json:"legendgroup,omitempty"`
    
    // Line 
    //   
    Line *ScattergeoLine `json:"line,omitempty"`
    
    // Locationmode 
    // enumerated Determines the set of locations used to match entries in `locations` to regions on the map. Values *ISO-3*, *USA-states*, *country names* correspond to features on the base map and value *geojson-id* corresponds to features from a custom GeoJSON linked to the `geojson` attribute. 
    Locationmode ScattergeoLocationmode `json:"locationmode,omitempty"`
    
    // Locations 
    // data_array 
    // Sets the coordinates via location IDs or names. Coordinates correspond to the centroid of each location given. See `locationmode` for more info. 
    Locations interface{} `json:"locations,omitempty"`
    
    // Locationssrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  locations . 
    Locationssrc String `json:"locationssrc,omitempty"`
    
    // Lon 
    // data_array 
    // Sets the longitude coordinates (in degrees East). 
    Lon interface{} `json:"lon,omitempty"`
    
    // Lonsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  lon . 
    Lonsrc String `json:"lonsrc,omitempty"`
    
    // Marker 
    //   
    Marker *ScattergeoMarker `json:"marker,omitempty"`
    
    // Meta 
    // any 
    // Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index. 
    Meta interface{} `json:"meta,omitempty"`
    
    // Metasrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  meta . 
    Metasrc String `json:"metasrc,omitempty"`
    
    // Mode 
    // flaglist Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*. 
    Mode ScattergeoMode `json:"mode,omitempty"`
    
    // Name 
    // string 
    // Sets the trace name. The trace name appear as the legend item and on hover. 
    Name String `json:"name,omitempty"`
    
    // Opacity 
    // number 
    // Sets the opacity of the trace. 
    Opacity float64 `json:"opacity,omitempty"`
    
    // Selected 
    //   
    Selected *ScattergeoSelected `json:"selected,omitempty"`
    
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
    Stream *ScattergeoStream `json:"stream,omitempty"`
    
    // Text 
    // string 
    // Sets text elements associated with each (lon,lat) pair or item in `locations`. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (lon,lat) or `locations` coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels. 
    Text String `json:"text,omitempty"`
    
    // Textfont 
    //  Sets the text font. 
    Textfont *ScattergeoTextfont `json:"textfont,omitempty"`
    
    // Textposition 
    // enumerated Sets the positions of the `text` elements with respects to the (x,y) coordinates. 
    Textposition ScattergeoTextposition `json:"textposition,omitempty"`
    
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
    // Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `lat`, `lon`, `location` and `text`. 
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
    Unselected *ScattergeoUnselected `json:"unselected,omitempty"`
    
    // Visible 
    // enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible). 
    Visible ScattergeoVisible `json:"visible,omitempty"`
    
}
// ScattergeoHoverlabelFont Sets the font used in hover labels.
type ScattergeoHoverlabelFont struct {
    
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
// ScattergeoHoverlabel 
type ScattergeoHoverlabel struct {
    
    // Align 
    // enumerated Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines 
    Align ScattergeoHoverlabelAlign `json:"align,omitempty"`
    
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
    Font *ScattergeoHoverlabelFont `json:"font,omitempty"`
    
    // Namelength 
    // integer 
    // Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis. 
    Namelength int64 `json:"namelength,omitempty"`
    
    // Namelengthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  namelength . 
    Namelengthsrc String `json:"namelengthsrc,omitempty"`
    
}
// ScattergeoLine 
type ScattergeoLine struct {
    
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
// ScattergeoMarkerColorbarTickfont Sets the color bar's tick label font
type ScattergeoMarkerColorbarTickfont struct {
    
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
// ScattergeoMarkerColorbarTitleFont Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute.
type ScattergeoMarkerColorbarTitleFont struct {
    
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
// ScattergeoMarkerColorbarTitle 
type ScattergeoMarkerColorbarTitle struct {
    
    // Font 
    //  Sets this color bar's title font. Note that the title's font used to be set by the now deprecated `titlefont` attribute. 
    Font *ScattergeoMarkerColorbarTitleFont `json:"font,omitempty"`
    
    // Side 
    // enumerated Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute. 
    Side ScattergeoMarkerColorbarTitleSide `json:"side,omitempty"`
    
    // Text 
    // string 
    // Sets the title of the color bar. Note that before the existence of `title.text`, the title's contents used to be defined as the `title` attribute itself. This behavior has been deprecated. 
    Text String `json:"text,omitempty"`
    
}
// ScattergeoMarkerColorbar 
type ScattergeoMarkerColorbar struct {
    
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
    Exponentformat ScattergeoMarkerColorbarExponentformat `json:"exponentformat,omitempty"`
    
    // Len 
    // number 
    // Sets the length of the color bar This measure excludes the padding of both ends. That is, the color bar length is this length minus the padding on both ends. 
    Len float64 `json:"len,omitempty"`
    
    // Lenmode 
    // enumerated Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value. 
    Lenmode ScattergeoMarkerColorbarLenmode `json:"lenmode,omitempty"`
    
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
    Showexponent ScattergeoMarkerColorbarShowexponent `json:"showexponent,omitempty"`
    
    // Showticklabels 
    // boolean 
    // Determines whether or not the tick labels are drawn. 
    Showticklabels Bool `json:"showticklabels,omitempty"`
    
    // Showtickprefix 
    // enumerated If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden. 
    Showtickprefix ScattergeoMarkerColorbarShowtickprefix `json:"showtickprefix,omitempty"`
    
    // Showticksuffix 
    // enumerated Same as `showtickprefix` but for tick suffixes. 
    Showticksuffix ScattergeoMarkerColorbarShowticksuffix `json:"showticksuffix,omitempty"`
    
    // Thickness 
    // number 
    // Sets the thickness of the color bar This measure excludes the size of the padding, ticks and labels. 
    Thickness float64 `json:"thickness,omitempty"`
    
    // Thicknessmode 
    // enumerated Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value. 
    Thicknessmode ScattergeoMarkerColorbarThicknessmode `json:"thicknessmode,omitempty"`
    
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
    Tickfont *ScattergeoMarkerColorbarTickfont `json:"tickfont,omitempty"`
    
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
    Ticklabelposition ScattergeoMarkerColorbarTicklabelposition `json:"ticklabelposition,omitempty"`
    
    // Ticklen 
    // number 
    // Sets the tick length (in px). 
    Ticklen float64 `json:"ticklen,omitempty"`
    
    // Tickmode 
    // enumerated Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided). 
    Tickmode ScattergeoMarkerColorbarTickmode `json:"tickmode,omitempty"`
    
    // Tickprefix 
    // string 
    // Sets a tick label prefix. 
    Tickprefix String `json:"tickprefix,omitempty"`
    
    // Ticks 
    // enumerated Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines. 
    Ticks ScattergeoMarkerColorbarTicks `json:"ticks,omitempty"`
    
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
    Title *ScattergeoMarkerColorbarTitle `json:"title,omitempty"`
    
    // X 
    // number 
    // Sets the x position of the color bar (in plot fraction). 
    X float64 `json:"x,omitempty"`
    
    // Xanchor 
    // enumerated Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar. 
    Xanchor ScattergeoMarkerColorbarXanchor `json:"xanchor,omitempty"`
    
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
    Yanchor ScattergeoMarkerColorbarYanchor `json:"yanchor,omitempty"`
    
    // Ypad 
    // number 
    // Sets the amount of padding (in px) along the y direction. 
    Ypad float64 `json:"ypad,omitempty"`
    
}
// ScattergeoMarkerGradient 
type ScattergeoMarkerGradient struct {
    
    // Color 
    // color 
    // Sets the final color of the gradient fill: the center color for radial, the right for horizontal, or the bottom for vertical. 
    Color Color `json:"color,omitempty"`
    
    // Colorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  color . 
    Colorsrc String `json:"colorsrc,omitempty"`
    
    // Type 
    // enumerated Sets the type of gradient used to fill the markers 
    Type ScattergeoMarkerGradientType `json:"type,omitempty"`
    
    // Typesrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  type . 
    Typesrc String `json:"typesrc,omitempty"`
    
}
// ScattergeoMarkerLine 
type ScattergeoMarkerLine struct {
    
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
// ScattergeoMarker 
type ScattergeoMarker struct {
    
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
    Colorbar *ScattergeoMarkerColorbar `json:"colorbar,omitempty"`
    
    // Colorscale 
    // colorscale 
    // Sets the colorscale. Has an effect only if in `marker.color`is set to a numerical array. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`marker.cmin` and `marker.cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis. 
    // Default: <nil> 
    Colorscale ColorScale `json:"colorscale,omitempty"`
    
    // Colorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  color . 
    Colorsrc String `json:"colorsrc,omitempty"`
    
    // Gradient 
    //   
    Gradient *ScattergeoMarkerGradient `json:"gradient,omitempty"`
    
    // Line 
    //   
    Line *ScattergeoMarkerLine `json:"line,omitempty"`
    
    // Opacity 
    // number 
    // Sets the marker opacity. 
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
    
    // Size 
    // number 
    // Sets the marker size (in px). 
    Size float64 `json:"size,omitempty"`
    
    // Sizemin 
    // number 
    // Has an effect only if `marker.size` is set to a numerical array. Sets the minimum size (in px) of the rendered marker points. 
    Sizemin float64 `json:"sizemin,omitempty"`
    
    // Sizemode 
    // enumerated Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels. 
    Sizemode ScattergeoMarkerSizemode `json:"sizemode,omitempty"`
    
    // Sizeref 
    // number 
    // Has an effect only if `marker.size` is set to a numerical array. Sets the scale factor used to determine the rendered size of marker points. Use with `sizemin` and `sizemode`. 
    Sizeref float64 `json:"sizeref,omitempty"`
    
    // Sizesrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  size . 
    Sizesrc String `json:"sizesrc,omitempty"`
    
    // Symbol 
    // enumerated Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name. 
    Symbol ScattergeoMarkerSymbol `json:"symbol,omitempty"`
    
    // Symbolsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  symbol . 
    Symbolsrc String `json:"symbolsrc,omitempty"`
    
}
// ScattergeoSelectedMarker 
type ScattergeoSelectedMarker struct {
    
    // Color 
    // color 
    // Sets the marker color of selected points. 
    Color Color `json:"color,omitempty"`
    
    // Opacity 
    // number 
    // Sets the marker opacity of selected points. 
    Opacity float64 `json:"opacity,omitempty"`
    
    // Size 
    // number 
    // Sets the marker size of selected points. 
    Size float64 `json:"size,omitempty"`
    
}
// ScattergeoSelectedTextfont 
type ScattergeoSelectedTextfont struct {
    
    // Color 
    // color 
    // Sets the text font color of selected points. 
    Color Color `json:"color,omitempty"`
    
}
// ScattergeoSelected 
type ScattergeoSelected struct {
    
    // Marker 
    //   
    Marker *ScattergeoSelectedMarker `json:"marker,omitempty"`
    
    // Textfont 
    //   
    Textfont *ScattergeoSelectedTextfont `json:"textfont,omitempty"`
    
}
// ScattergeoStream 
type ScattergeoStream struct {
    
    // Maxpoints 
    // number 
    // Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot. 
    Maxpoints float64 `json:"maxpoints,omitempty"`
    
    // Token 
    // string 
    // The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details. 
    Token String `json:"token,omitempty"`
    
}
// ScattergeoTextfont Sets the text font.
type ScattergeoTextfont struct {
    
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
// ScattergeoUnselectedMarker 
type ScattergeoUnselectedMarker struct {
    
    // Color 
    // color 
    // Sets the marker color of unselected points, applied only when a selection exists. 
    Color Color `json:"color,omitempty"`
    
    // Opacity 
    // number 
    // Sets the marker opacity of unselected points, applied only when a selection exists. 
    Opacity float64 `json:"opacity,omitempty"`
    
    // Size 
    // number 
    // Sets the marker size of unselected points, applied only when a selection exists. 
    Size float64 `json:"size,omitempty"`
    
}
// ScattergeoUnselectedTextfont 
type ScattergeoUnselectedTextfont struct {
    
    // Color 
    // color 
    // Sets the text font color of unselected points, applied only when a selection exists. 
    Color Color `json:"color,omitempty"`
    
}
// ScattergeoUnselected 
type ScattergeoUnselected struct {
    
    // Marker 
    //   
    Marker *ScattergeoUnselectedMarker `json:"marker,omitempty"`
    
    // Textfont 
    //   
    Textfont *ScattergeoUnselectedTextfont `json:"textfont,omitempty"`
    
}
// ScattergeoFill Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape.
type ScattergeoFill string 

const (
    ScattergeoFillNone ScattergeoFill = "none"
    ScattergeoFillToself ScattergeoFill = "toself"
    
)
// ScattergeoHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScattergeoHoverlabelAlign string 

const (
    ScattergeoHoverlabelAlignLeft ScattergeoHoverlabelAlign = "left"
    ScattergeoHoverlabelAlignRight ScattergeoHoverlabelAlign = "right"
    ScattergeoHoverlabelAlignAuto ScattergeoHoverlabelAlign = "auto"
    
)
// ScattergeoLocationmode Determines the set of locations used to match entries in `locations` to regions on the map. Values *ISO-3*, *USA-states*, *country names* correspond to features on the base map and value *geojson-id* corresponds to features from a custom GeoJSON linked to the `geojson` attribute.
type ScattergeoLocationmode string 

const (
    ScattergeoLocationmodeIso3 ScattergeoLocationmode = "ISO-3"
    ScattergeoLocationmodeUsaStates ScattergeoLocationmode = "USA-states"
    ScattergeoLocationmodeCountryNames ScattergeoLocationmode = "country names"
    ScattergeoLocationmodeGeojsonId ScattergeoLocationmode = "geojson-id"
    
)
// ScattergeoMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScattergeoMarkerColorbarExponentformat string 

const (
    ScattergeoMarkerColorbarExponentformatNone ScattergeoMarkerColorbarExponentformat = "none"
    ScattergeoMarkerColorbarExponentformatE1 ScattergeoMarkerColorbarExponentformat = "e"
    ScattergeoMarkerColorbarExponentformatE2 ScattergeoMarkerColorbarExponentformat = "E"
    ScattergeoMarkerColorbarExponentformatPower ScattergeoMarkerColorbarExponentformat = "power"
    ScattergeoMarkerColorbarExponentformatSi ScattergeoMarkerColorbarExponentformat = "SI"
    ScattergeoMarkerColorbarExponentformatB ScattergeoMarkerColorbarExponentformat = "B"
    
)
// ScattergeoMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScattergeoMarkerColorbarLenmode string 

const (
    ScattergeoMarkerColorbarLenmodeFraction ScattergeoMarkerColorbarLenmode = "fraction"
    ScattergeoMarkerColorbarLenmodePixels ScattergeoMarkerColorbarLenmode = "pixels"
    
)
// ScattergeoMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScattergeoMarkerColorbarShowexponent string 

const (
    ScattergeoMarkerColorbarShowexponentAll ScattergeoMarkerColorbarShowexponent = "all"
    ScattergeoMarkerColorbarShowexponentFirst ScattergeoMarkerColorbarShowexponent = "first"
    ScattergeoMarkerColorbarShowexponentLast ScattergeoMarkerColorbarShowexponent = "last"
    ScattergeoMarkerColorbarShowexponentNone ScattergeoMarkerColorbarShowexponent = "none"
    
)
// ScattergeoMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScattergeoMarkerColorbarShowtickprefix string 

const (
    ScattergeoMarkerColorbarShowtickprefixAll ScattergeoMarkerColorbarShowtickprefix = "all"
    ScattergeoMarkerColorbarShowtickprefixFirst ScattergeoMarkerColorbarShowtickprefix = "first"
    ScattergeoMarkerColorbarShowtickprefixLast ScattergeoMarkerColorbarShowtickprefix = "last"
    ScattergeoMarkerColorbarShowtickprefixNone ScattergeoMarkerColorbarShowtickprefix = "none"
    
)
// ScattergeoMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScattergeoMarkerColorbarShowticksuffix string 

const (
    ScattergeoMarkerColorbarShowticksuffixAll ScattergeoMarkerColorbarShowticksuffix = "all"
    ScattergeoMarkerColorbarShowticksuffixFirst ScattergeoMarkerColorbarShowticksuffix = "first"
    ScattergeoMarkerColorbarShowticksuffixLast ScattergeoMarkerColorbarShowticksuffix = "last"
    ScattergeoMarkerColorbarShowticksuffixNone ScattergeoMarkerColorbarShowticksuffix = "none"
    
)
// ScattergeoMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScattergeoMarkerColorbarThicknessmode string 

const (
    ScattergeoMarkerColorbarThicknessmodeFraction ScattergeoMarkerColorbarThicknessmode = "fraction"
    ScattergeoMarkerColorbarThicknessmodePixels ScattergeoMarkerColorbarThicknessmode = "pixels"
    
)
// ScattergeoMarkerColorbarTicklabelposition Determines where tick labels are drawn.
type ScattergeoMarkerColorbarTicklabelposition string 

const (
    ScattergeoMarkerColorbarTicklabelpositionOutside ScattergeoMarkerColorbarTicklabelposition = "outside"
    ScattergeoMarkerColorbarTicklabelpositionInside ScattergeoMarkerColorbarTicklabelposition = "inside"
    ScattergeoMarkerColorbarTicklabelpositionOutsideTop ScattergeoMarkerColorbarTicklabelposition = "outside top"
    ScattergeoMarkerColorbarTicklabelpositionInsideTop ScattergeoMarkerColorbarTicklabelposition = "inside top"
    ScattergeoMarkerColorbarTicklabelpositionOutsideBottom ScattergeoMarkerColorbarTicklabelposition = "outside bottom"
    ScattergeoMarkerColorbarTicklabelpositionInsideBottom ScattergeoMarkerColorbarTicklabelposition = "inside bottom"
    
)
// ScattergeoMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScattergeoMarkerColorbarTickmode string 

const (
    ScattergeoMarkerColorbarTickmodeAuto ScattergeoMarkerColorbarTickmode = "auto"
    ScattergeoMarkerColorbarTickmodeLinear ScattergeoMarkerColorbarTickmode = "linear"
    ScattergeoMarkerColorbarTickmodeArray ScattergeoMarkerColorbarTickmode = "array"
    
)
// ScattergeoMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScattergeoMarkerColorbarTicks string 

const (
    ScattergeoMarkerColorbarTicksOutside ScattergeoMarkerColorbarTicks = "outside"
    ScattergeoMarkerColorbarTicksInside ScattergeoMarkerColorbarTicks = "inside"
    ScattergeoMarkerColorbarTicksEmpty ScattergeoMarkerColorbarTicks = ""
    
)
// ScattergeoMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScattergeoMarkerColorbarTitleSide string 

const (
    ScattergeoMarkerColorbarTitleSideRight ScattergeoMarkerColorbarTitleSide = "right"
    ScattergeoMarkerColorbarTitleSideTop ScattergeoMarkerColorbarTitleSide = "top"
    ScattergeoMarkerColorbarTitleSideBottom ScattergeoMarkerColorbarTitleSide = "bottom"
    
)
// ScattergeoMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScattergeoMarkerColorbarXanchor string 

const (
    ScattergeoMarkerColorbarXanchorLeft ScattergeoMarkerColorbarXanchor = "left"
    ScattergeoMarkerColorbarXanchorCenter ScattergeoMarkerColorbarXanchor = "center"
    ScattergeoMarkerColorbarXanchorRight ScattergeoMarkerColorbarXanchor = "right"
    
)
// ScattergeoMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScattergeoMarkerColorbarYanchor string 

const (
    ScattergeoMarkerColorbarYanchorTop ScattergeoMarkerColorbarYanchor = "top"
    ScattergeoMarkerColorbarYanchorMiddle ScattergeoMarkerColorbarYanchor = "middle"
    ScattergeoMarkerColorbarYanchorBottom ScattergeoMarkerColorbarYanchor = "bottom"
    
)
// ScattergeoMarkerGradientType Sets the type of gradient used to fill the markers
type ScattergeoMarkerGradientType string 

const (
    ScattergeoMarkerGradientTypeRadial ScattergeoMarkerGradientType = "radial"
    ScattergeoMarkerGradientTypeHorizontal ScattergeoMarkerGradientType = "horizontal"
    ScattergeoMarkerGradientTypeVertical ScattergeoMarkerGradientType = "vertical"
    ScattergeoMarkerGradientTypeNone ScattergeoMarkerGradientType = "none"
    
)
// ScattergeoMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScattergeoMarkerSizemode string 

const (
    ScattergeoMarkerSizemodeDiameter ScattergeoMarkerSizemode = "diameter"
    ScattergeoMarkerSizemodeArea ScattergeoMarkerSizemode = "area"
    
)
// ScattergeoMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScattergeoMarkerSymbol interface{} 

var (
    ScattergeoMarkerSymbolNumber0 ScattergeoMarkerSymbol = 0
    ScattergeoMarkerSymbol0 ScattergeoMarkerSymbol = "0"
    ScattergeoMarkerSymbolCircle ScattergeoMarkerSymbol = "circle"
    ScattergeoMarkerSymbolNumber100 ScattergeoMarkerSymbol = 100
    ScattergeoMarkerSymbol100 ScattergeoMarkerSymbol = "100"
    ScattergeoMarkerSymbolCircleOpen ScattergeoMarkerSymbol = "circle-open"
    ScattergeoMarkerSymbolNumber200 ScattergeoMarkerSymbol = 200
    ScattergeoMarkerSymbol200 ScattergeoMarkerSymbol = "200"
    ScattergeoMarkerSymbolCircleDot ScattergeoMarkerSymbol = "circle-dot"
    ScattergeoMarkerSymbolNumber300 ScattergeoMarkerSymbol = 300
    ScattergeoMarkerSymbol300 ScattergeoMarkerSymbol = "300"
    ScattergeoMarkerSymbolCircleOpenDot ScattergeoMarkerSymbol = "circle-open-dot"
    ScattergeoMarkerSymbolNumber1 ScattergeoMarkerSymbol = 1
    ScattergeoMarkerSymbol1 ScattergeoMarkerSymbol = "1"
    ScattergeoMarkerSymbolSquare ScattergeoMarkerSymbol = "square"
    ScattergeoMarkerSymbolNumber101 ScattergeoMarkerSymbol = 101
    ScattergeoMarkerSymbol101 ScattergeoMarkerSymbol = "101"
    ScattergeoMarkerSymbolSquareOpen ScattergeoMarkerSymbol = "square-open"
    ScattergeoMarkerSymbolNumber201 ScattergeoMarkerSymbol = 201
    ScattergeoMarkerSymbol201 ScattergeoMarkerSymbol = "201"
    ScattergeoMarkerSymbolSquareDot ScattergeoMarkerSymbol = "square-dot"
    ScattergeoMarkerSymbolNumber301 ScattergeoMarkerSymbol = 301
    ScattergeoMarkerSymbol301 ScattergeoMarkerSymbol = "301"
    ScattergeoMarkerSymbolSquareOpenDot ScattergeoMarkerSymbol = "square-open-dot"
    ScattergeoMarkerSymbolNumber2 ScattergeoMarkerSymbol = 2
    ScattergeoMarkerSymbol2 ScattergeoMarkerSymbol = "2"
    ScattergeoMarkerSymbolDiamond ScattergeoMarkerSymbol = "diamond"
    ScattergeoMarkerSymbolNumber102 ScattergeoMarkerSymbol = 102
    ScattergeoMarkerSymbol102 ScattergeoMarkerSymbol = "102"
    ScattergeoMarkerSymbolDiamondOpen ScattergeoMarkerSymbol = "diamond-open"
    ScattergeoMarkerSymbolNumber202 ScattergeoMarkerSymbol = 202
    ScattergeoMarkerSymbol202 ScattergeoMarkerSymbol = "202"
    ScattergeoMarkerSymbolDiamondDot ScattergeoMarkerSymbol = "diamond-dot"
    ScattergeoMarkerSymbolNumber302 ScattergeoMarkerSymbol = 302
    ScattergeoMarkerSymbol302 ScattergeoMarkerSymbol = "302"
    ScattergeoMarkerSymbolDiamondOpenDot ScattergeoMarkerSymbol = "diamond-open-dot"
    ScattergeoMarkerSymbolNumber3 ScattergeoMarkerSymbol = 3
    ScattergeoMarkerSymbol3 ScattergeoMarkerSymbol = "3"
    ScattergeoMarkerSymbolCross ScattergeoMarkerSymbol = "cross"
    ScattergeoMarkerSymbolNumber103 ScattergeoMarkerSymbol = 103
    ScattergeoMarkerSymbol103 ScattergeoMarkerSymbol = "103"
    ScattergeoMarkerSymbolCrossOpen ScattergeoMarkerSymbol = "cross-open"
    ScattergeoMarkerSymbolNumber203 ScattergeoMarkerSymbol = 203
    ScattergeoMarkerSymbol203 ScattergeoMarkerSymbol = "203"
    ScattergeoMarkerSymbolCrossDot ScattergeoMarkerSymbol = "cross-dot"
    ScattergeoMarkerSymbolNumber303 ScattergeoMarkerSymbol = 303
    ScattergeoMarkerSymbol303 ScattergeoMarkerSymbol = "303"
    ScattergeoMarkerSymbolCrossOpenDot ScattergeoMarkerSymbol = "cross-open-dot"
    ScattergeoMarkerSymbolNumber4 ScattergeoMarkerSymbol = 4
    ScattergeoMarkerSymbol4 ScattergeoMarkerSymbol = "4"
    ScattergeoMarkerSymbolX ScattergeoMarkerSymbol = "x"
    ScattergeoMarkerSymbolNumber104 ScattergeoMarkerSymbol = 104
    ScattergeoMarkerSymbol104 ScattergeoMarkerSymbol = "104"
    ScattergeoMarkerSymbolXOpen ScattergeoMarkerSymbol = "x-open"
    ScattergeoMarkerSymbolNumber204 ScattergeoMarkerSymbol = 204
    ScattergeoMarkerSymbol204 ScattergeoMarkerSymbol = "204"
    ScattergeoMarkerSymbolXDot ScattergeoMarkerSymbol = "x-dot"
    ScattergeoMarkerSymbolNumber304 ScattergeoMarkerSymbol = 304
    ScattergeoMarkerSymbol304 ScattergeoMarkerSymbol = "304"
    ScattergeoMarkerSymbolXOpenDot ScattergeoMarkerSymbol = "x-open-dot"
    ScattergeoMarkerSymbolNumber5 ScattergeoMarkerSymbol = 5
    ScattergeoMarkerSymbol5 ScattergeoMarkerSymbol = "5"
    ScattergeoMarkerSymbolTriangleUp ScattergeoMarkerSymbol = "triangle-up"
    ScattergeoMarkerSymbolNumber105 ScattergeoMarkerSymbol = 105
    ScattergeoMarkerSymbol105 ScattergeoMarkerSymbol = "105"
    ScattergeoMarkerSymbolTriangleUpOpen ScattergeoMarkerSymbol = "triangle-up-open"
    ScattergeoMarkerSymbolNumber205 ScattergeoMarkerSymbol = 205
    ScattergeoMarkerSymbol205 ScattergeoMarkerSymbol = "205"
    ScattergeoMarkerSymbolTriangleUpDot ScattergeoMarkerSymbol = "triangle-up-dot"
    ScattergeoMarkerSymbolNumber305 ScattergeoMarkerSymbol = 305
    ScattergeoMarkerSymbol305 ScattergeoMarkerSymbol = "305"
    ScattergeoMarkerSymbolTriangleUpOpenDot ScattergeoMarkerSymbol = "triangle-up-open-dot"
    ScattergeoMarkerSymbolNumber6 ScattergeoMarkerSymbol = 6
    ScattergeoMarkerSymbol6 ScattergeoMarkerSymbol = "6"
    ScattergeoMarkerSymbolTriangleDown ScattergeoMarkerSymbol = "triangle-down"
    ScattergeoMarkerSymbolNumber106 ScattergeoMarkerSymbol = 106
    ScattergeoMarkerSymbol106 ScattergeoMarkerSymbol = "106"
    ScattergeoMarkerSymbolTriangleDownOpen ScattergeoMarkerSymbol = "triangle-down-open"
    ScattergeoMarkerSymbolNumber206 ScattergeoMarkerSymbol = 206
    ScattergeoMarkerSymbol206 ScattergeoMarkerSymbol = "206"
    ScattergeoMarkerSymbolTriangleDownDot ScattergeoMarkerSymbol = "triangle-down-dot"
    ScattergeoMarkerSymbolNumber306 ScattergeoMarkerSymbol = 306
    ScattergeoMarkerSymbol306 ScattergeoMarkerSymbol = "306"
    ScattergeoMarkerSymbolTriangleDownOpenDot ScattergeoMarkerSymbol = "triangle-down-open-dot"
    ScattergeoMarkerSymbolNumber7 ScattergeoMarkerSymbol = 7
    ScattergeoMarkerSymbol7 ScattergeoMarkerSymbol = "7"
    ScattergeoMarkerSymbolTriangleLeft ScattergeoMarkerSymbol = "triangle-left"
    ScattergeoMarkerSymbolNumber107 ScattergeoMarkerSymbol = 107
    ScattergeoMarkerSymbol107 ScattergeoMarkerSymbol = "107"
    ScattergeoMarkerSymbolTriangleLeftOpen ScattergeoMarkerSymbol = "triangle-left-open"
    ScattergeoMarkerSymbolNumber207 ScattergeoMarkerSymbol = 207
    ScattergeoMarkerSymbol207 ScattergeoMarkerSymbol = "207"
    ScattergeoMarkerSymbolTriangleLeftDot ScattergeoMarkerSymbol = "triangle-left-dot"
    ScattergeoMarkerSymbolNumber307 ScattergeoMarkerSymbol = 307
    ScattergeoMarkerSymbol307 ScattergeoMarkerSymbol = "307"
    ScattergeoMarkerSymbolTriangleLeftOpenDot ScattergeoMarkerSymbol = "triangle-left-open-dot"
    ScattergeoMarkerSymbolNumber8 ScattergeoMarkerSymbol = 8
    ScattergeoMarkerSymbol8 ScattergeoMarkerSymbol = "8"
    ScattergeoMarkerSymbolTriangleRight ScattergeoMarkerSymbol = "triangle-right"
    ScattergeoMarkerSymbolNumber108 ScattergeoMarkerSymbol = 108
    ScattergeoMarkerSymbol108 ScattergeoMarkerSymbol = "108"
    ScattergeoMarkerSymbolTriangleRightOpen ScattergeoMarkerSymbol = "triangle-right-open"
    ScattergeoMarkerSymbolNumber208 ScattergeoMarkerSymbol = 208
    ScattergeoMarkerSymbol208 ScattergeoMarkerSymbol = "208"
    ScattergeoMarkerSymbolTriangleRightDot ScattergeoMarkerSymbol = "triangle-right-dot"
    ScattergeoMarkerSymbolNumber308 ScattergeoMarkerSymbol = 308
    ScattergeoMarkerSymbol308 ScattergeoMarkerSymbol = "308"
    ScattergeoMarkerSymbolTriangleRightOpenDot ScattergeoMarkerSymbol = "triangle-right-open-dot"
    ScattergeoMarkerSymbolNumber9 ScattergeoMarkerSymbol = 9
    ScattergeoMarkerSymbol9 ScattergeoMarkerSymbol = "9"
    ScattergeoMarkerSymbolTriangleNe ScattergeoMarkerSymbol = "triangle-ne"
    ScattergeoMarkerSymbolNumber109 ScattergeoMarkerSymbol = 109
    ScattergeoMarkerSymbol109 ScattergeoMarkerSymbol = "109"
    ScattergeoMarkerSymbolTriangleNeOpen ScattergeoMarkerSymbol = "triangle-ne-open"
    ScattergeoMarkerSymbolNumber209 ScattergeoMarkerSymbol = 209
    ScattergeoMarkerSymbol209 ScattergeoMarkerSymbol = "209"
    ScattergeoMarkerSymbolTriangleNeDot ScattergeoMarkerSymbol = "triangle-ne-dot"
    ScattergeoMarkerSymbolNumber309 ScattergeoMarkerSymbol = 309
    ScattergeoMarkerSymbol309 ScattergeoMarkerSymbol = "309"
    ScattergeoMarkerSymbolTriangleNeOpenDot ScattergeoMarkerSymbol = "triangle-ne-open-dot"
    ScattergeoMarkerSymbolNumber10 ScattergeoMarkerSymbol = 10
    ScattergeoMarkerSymbol10 ScattergeoMarkerSymbol = "10"
    ScattergeoMarkerSymbolTriangleSe ScattergeoMarkerSymbol = "triangle-se"
    ScattergeoMarkerSymbolNumber110 ScattergeoMarkerSymbol = 110
    ScattergeoMarkerSymbol110 ScattergeoMarkerSymbol = "110"
    ScattergeoMarkerSymbolTriangleSeOpen ScattergeoMarkerSymbol = "triangle-se-open"
    ScattergeoMarkerSymbolNumber210 ScattergeoMarkerSymbol = 210
    ScattergeoMarkerSymbol210 ScattergeoMarkerSymbol = "210"
    ScattergeoMarkerSymbolTriangleSeDot ScattergeoMarkerSymbol = "triangle-se-dot"
    ScattergeoMarkerSymbolNumber310 ScattergeoMarkerSymbol = 310
    ScattergeoMarkerSymbol310 ScattergeoMarkerSymbol = "310"
    ScattergeoMarkerSymbolTriangleSeOpenDot ScattergeoMarkerSymbol = "triangle-se-open-dot"
    ScattergeoMarkerSymbolNumber11 ScattergeoMarkerSymbol = 11
    ScattergeoMarkerSymbol11 ScattergeoMarkerSymbol = "11"
    ScattergeoMarkerSymbolTriangleSw ScattergeoMarkerSymbol = "triangle-sw"
    ScattergeoMarkerSymbolNumber111 ScattergeoMarkerSymbol = 111
    ScattergeoMarkerSymbol111 ScattergeoMarkerSymbol = "111"
    ScattergeoMarkerSymbolTriangleSwOpen ScattergeoMarkerSymbol = "triangle-sw-open"
    ScattergeoMarkerSymbolNumber211 ScattergeoMarkerSymbol = 211
    ScattergeoMarkerSymbol211 ScattergeoMarkerSymbol = "211"
    ScattergeoMarkerSymbolTriangleSwDot ScattergeoMarkerSymbol = "triangle-sw-dot"
    ScattergeoMarkerSymbolNumber311 ScattergeoMarkerSymbol = 311
    ScattergeoMarkerSymbol311 ScattergeoMarkerSymbol = "311"
    ScattergeoMarkerSymbolTriangleSwOpenDot ScattergeoMarkerSymbol = "triangle-sw-open-dot"
    ScattergeoMarkerSymbolNumber12 ScattergeoMarkerSymbol = 12
    ScattergeoMarkerSymbol12 ScattergeoMarkerSymbol = "12"
    ScattergeoMarkerSymbolTriangleNw ScattergeoMarkerSymbol = "triangle-nw"
    ScattergeoMarkerSymbolNumber112 ScattergeoMarkerSymbol = 112
    ScattergeoMarkerSymbol112 ScattergeoMarkerSymbol = "112"
    ScattergeoMarkerSymbolTriangleNwOpen ScattergeoMarkerSymbol = "triangle-nw-open"
    ScattergeoMarkerSymbolNumber212 ScattergeoMarkerSymbol = 212
    ScattergeoMarkerSymbol212 ScattergeoMarkerSymbol = "212"
    ScattergeoMarkerSymbolTriangleNwDot ScattergeoMarkerSymbol = "triangle-nw-dot"
    ScattergeoMarkerSymbolNumber312 ScattergeoMarkerSymbol = 312
    ScattergeoMarkerSymbol312 ScattergeoMarkerSymbol = "312"
    ScattergeoMarkerSymbolTriangleNwOpenDot ScattergeoMarkerSymbol = "triangle-nw-open-dot"
    ScattergeoMarkerSymbolNumber13 ScattergeoMarkerSymbol = 13
    ScattergeoMarkerSymbol13 ScattergeoMarkerSymbol = "13"
    ScattergeoMarkerSymbolPentagon ScattergeoMarkerSymbol = "pentagon"
    ScattergeoMarkerSymbolNumber113 ScattergeoMarkerSymbol = 113
    ScattergeoMarkerSymbol113 ScattergeoMarkerSymbol = "113"
    ScattergeoMarkerSymbolPentagonOpen ScattergeoMarkerSymbol = "pentagon-open"
    ScattergeoMarkerSymbolNumber213 ScattergeoMarkerSymbol = 213
    ScattergeoMarkerSymbol213 ScattergeoMarkerSymbol = "213"
    ScattergeoMarkerSymbolPentagonDot ScattergeoMarkerSymbol = "pentagon-dot"
    ScattergeoMarkerSymbolNumber313 ScattergeoMarkerSymbol = 313
    ScattergeoMarkerSymbol313 ScattergeoMarkerSymbol = "313"
    ScattergeoMarkerSymbolPentagonOpenDot ScattergeoMarkerSymbol = "pentagon-open-dot"
    ScattergeoMarkerSymbolNumber14 ScattergeoMarkerSymbol = 14
    ScattergeoMarkerSymbol14 ScattergeoMarkerSymbol = "14"
    ScattergeoMarkerSymbolHexagon ScattergeoMarkerSymbol = "hexagon"
    ScattergeoMarkerSymbolNumber114 ScattergeoMarkerSymbol = 114
    ScattergeoMarkerSymbol114 ScattergeoMarkerSymbol = "114"
    ScattergeoMarkerSymbolHexagonOpen ScattergeoMarkerSymbol = "hexagon-open"
    ScattergeoMarkerSymbolNumber214 ScattergeoMarkerSymbol = 214
    ScattergeoMarkerSymbol214 ScattergeoMarkerSymbol = "214"
    ScattergeoMarkerSymbolHexagonDot ScattergeoMarkerSymbol = "hexagon-dot"
    ScattergeoMarkerSymbolNumber314 ScattergeoMarkerSymbol = 314
    ScattergeoMarkerSymbol314 ScattergeoMarkerSymbol = "314"
    ScattergeoMarkerSymbolHexagonOpenDot ScattergeoMarkerSymbol = "hexagon-open-dot"
    ScattergeoMarkerSymbolNumber15 ScattergeoMarkerSymbol = 15
    ScattergeoMarkerSymbol15 ScattergeoMarkerSymbol = "15"
    ScattergeoMarkerSymbolHexagon2 ScattergeoMarkerSymbol = "hexagon2"
    ScattergeoMarkerSymbolNumber115 ScattergeoMarkerSymbol = 115
    ScattergeoMarkerSymbol115 ScattergeoMarkerSymbol = "115"
    ScattergeoMarkerSymbolHexagon2Open ScattergeoMarkerSymbol = "hexagon2-open"
    ScattergeoMarkerSymbolNumber215 ScattergeoMarkerSymbol = 215
    ScattergeoMarkerSymbol215 ScattergeoMarkerSymbol = "215"
    ScattergeoMarkerSymbolHexagon2Dot ScattergeoMarkerSymbol = "hexagon2-dot"
    ScattergeoMarkerSymbolNumber315 ScattergeoMarkerSymbol = 315
    ScattergeoMarkerSymbol315 ScattergeoMarkerSymbol = "315"
    ScattergeoMarkerSymbolHexagon2OpenDot ScattergeoMarkerSymbol = "hexagon2-open-dot"
    ScattergeoMarkerSymbolNumber16 ScattergeoMarkerSymbol = 16
    ScattergeoMarkerSymbol16 ScattergeoMarkerSymbol = "16"
    ScattergeoMarkerSymbolOctagon ScattergeoMarkerSymbol = "octagon"
    ScattergeoMarkerSymbolNumber116 ScattergeoMarkerSymbol = 116
    ScattergeoMarkerSymbol116 ScattergeoMarkerSymbol = "116"
    ScattergeoMarkerSymbolOctagonOpen ScattergeoMarkerSymbol = "octagon-open"
    ScattergeoMarkerSymbolNumber216 ScattergeoMarkerSymbol = 216
    ScattergeoMarkerSymbol216 ScattergeoMarkerSymbol = "216"
    ScattergeoMarkerSymbolOctagonDot ScattergeoMarkerSymbol = "octagon-dot"
    ScattergeoMarkerSymbolNumber316 ScattergeoMarkerSymbol = 316
    ScattergeoMarkerSymbol316 ScattergeoMarkerSymbol = "316"
    ScattergeoMarkerSymbolOctagonOpenDot ScattergeoMarkerSymbol = "octagon-open-dot"
    ScattergeoMarkerSymbolNumber17 ScattergeoMarkerSymbol = 17
    ScattergeoMarkerSymbol17 ScattergeoMarkerSymbol = "17"
    ScattergeoMarkerSymbolStar ScattergeoMarkerSymbol = "star"
    ScattergeoMarkerSymbolNumber117 ScattergeoMarkerSymbol = 117
    ScattergeoMarkerSymbol117 ScattergeoMarkerSymbol = "117"
    ScattergeoMarkerSymbolStarOpen ScattergeoMarkerSymbol = "star-open"
    ScattergeoMarkerSymbolNumber217 ScattergeoMarkerSymbol = 217
    ScattergeoMarkerSymbol217 ScattergeoMarkerSymbol = "217"
    ScattergeoMarkerSymbolStarDot ScattergeoMarkerSymbol = "star-dot"
    ScattergeoMarkerSymbolNumber317 ScattergeoMarkerSymbol = 317
    ScattergeoMarkerSymbol317 ScattergeoMarkerSymbol = "317"
    ScattergeoMarkerSymbolStarOpenDot ScattergeoMarkerSymbol = "star-open-dot"
    ScattergeoMarkerSymbolNumber18 ScattergeoMarkerSymbol = 18
    ScattergeoMarkerSymbol18 ScattergeoMarkerSymbol = "18"
    ScattergeoMarkerSymbolHexagram ScattergeoMarkerSymbol = "hexagram"
    ScattergeoMarkerSymbolNumber118 ScattergeoMarkerSymbol = 118
    ScattergeoMarkerSymbol118 ScattergeoMarkerSymbol = "118"
    ScattergeoMarkerSymbolHexagramOpen ScattergeoMarkerSymbol = "hexagram-open"
    ScattergeoMarkerSymbolNumber218 ScattergeoMarkerSymbol = 218
    ScattergeoMarkerSymbol218 ScattergeoMarkerSymbol = "218"
    ScattergeoMarkerSymbolHexagramDot ScattergeoMarkerSymbol = "hexagram-dot"
    ScattergeoMarkerSymbolNumber318 ScattergeoMarkerSymbol = 318
    ScattergeoMarkerSymbol318 ScattergeoMarkerSymbol = "318"
    ScattergeoMarkerSymbolHexagramOpenDot ScattergeoMarkerSymbol = "hexagram-open-dot"
    ScattergeoMarkerSymbolNumber19 ScattergeoMarkerSymbol = 19
    ScattergeoMarkerSymbol19 ScattergeoMarkerSymbol = "19"
    ScattergeoMarkerSymbolStarTriangleUp ScattergeoMarkerSymbol = "star-triangle-up"
    ScattergeoMarkerSymbolNumber119 ScattergeoMarkerSymbol = 119
    ScattergeoMarkerSymbol119 ScattergeoMarkerSymbol = "119"
    ScattergeoMarkerSymbolStarTriangleUpOpen ScattergeoMarkerSymbol = "star-triangle-up-open"
    ScattergeoMarkerSymbolNumber219 ScattergeoMarkerSymbol = 219
    ScattergeoMarkerSymbol219 ScattergeoMarkerSymbol = "219"
    ScattergeoMarkerSymbolStarTriangleUpDot ScattergeoMarkerSymbol = "star-triangle-up-dot"
    ScattergeoMarkerSymbolNumber319 ScattergeoMarkerSymbol = 319
    ScattergeoMarkerSymbol319 ScattergeoMarkerSymbol = "319"
    ScattergeoMarkerSymbolStarTriangleUpOpenDot ScattergeoMarkerSymbol = "star-triangle-up-open-dot"
    ScattergeoMarkerSymbolNumber20 ScattergeoMarkerSymbol = 20
    ScattergeoMarkerSymbol20 ScattergeoMarkerSymbol = "20"
    ScattergeoMarkerSymbolStarTriangleDown ScattergeoMarkerSymbol = "star-triangle-down"
    ScattergeoMarkerSymbolNumber120 ScattergeoMarkerSymbol = 120
    ScattergeoMarkerSymbol120 ScattergeoMarkerSymbol = "120"
    ScattergeoMarkerSymbolStarTriangleDownOpen ScattergeoMarkerSymbol = "star-triangle-down-open"
    ScattergeoMarkerSymbolNumber220 ScattergeoMarkerSymbol = 220
    ScattergeoMarkerSymbol220 ScattergeoMarkerSymbol = "220"
    ScattergeoMarkerSymbolStarTriangleDownDot ScattergeoMarkerSymbol = "star-triangle-down-dot"
    ScattergeoMarkerSymbolNumber320 ScattergeoMarkerSymbol = 320
    ScattergeoMarkerSymbol320 ScattergeoMarkerSymbol = "320"
    ScattergeoMarkerSymbolStarTriangleDownOpenDot ScattergeoMarkerSymbol = "star-triangle-down-open-dot"
    ScattergeoMarkerSymbolNumber21 ScattergeoMarkerSymbol = 21
    ScattergeoMarkerSymbol21 ScattergeoMarkerSymbol = "21"
    ScattergeoMarkerSymbolStarSquare ScattergeoMarkerSymbol = "star-square"
    ScattergeoMarkerSymbolNumber121 ScattergeoMarkerSymbol = 121
    ScattergeoMarkerSymbol121 ScattergeoMarkerSymbol = "121"
    ScattergeoMarkerSymbolStarSquareOpen ScattergeoMarkerSymbol = "star-square-open"
    ScattergeoMarkerSymbolNumber221 ScattergeoMarkerSymbol = 221
    ScattergeoMarkerSymbol221 ScattergeoMarkerSymbol = "221"
    ScattergeoMarkerSymbolStarSquareDot ScattergeoMarkerSymbol = "star-square-dot"
    ScattergeoMarkerSymbolNumber321 ScattergeoMarkerSymbol = 321
    ScattergeoMarkerSymbol321 ScattergeoMarkerSymbol = "321"
    ScattergeoMarkerSymbolStarSquareOpenDot ScattergeoMarkerSymbol = "star-square-open-dot"
    ScattergeoMarkerSymbolNumber22 ScattergeoMarkerSymbol = 22
    ScattergeoMarkerSymbol22 ScattergeoMarkerSymbol = "22"
    ScattergeoMarkerSymbolStarDiamond ScattergeoMarkerSymbol = "star-diamond"
    ScattergeoMarkerSymbolNumber122 ScattergeoMarkerSymbol = 122
    ScattergeoMarkerSymbol122 ScattergeoMarkerSymbol = "122"
    ScattergeoMarkerSymbolStarDiamondOpen ScattergeoMarkerSymbol = "star-diamond-open"
    ScattergeoMarkerSymbolNumber222 ScattergeoMarkerSymbol = 222
    ScattergeoMarkerSymbol222 ScattergeoMarkerSymbol = "222"
    ScattergeoMarkerSymbolStarDiamondDot ScattergeoMarkerSymbol = "star-diamond-dot"
    ScattergeoMarkerSymbolNumber322 ScattergeoMarkerSymbol = 322
    ScattergeoMarkerSymbol322 ScattergeoMarkerSymbol = "322"
    ScattergeoMarkerSymbolStarDiamondOpenDot ScattergeoMarkerSymbol = "star-diamond-open-dot"
    ScattergeoMarkerSymbolNumber23 ScattergeoMarkerSymbol = 23
    ScattergeoMarkerSymbol23 ScattergeoMarkerSymbol = "23"
    ScattergeoMarkerSymbolDiamondTall ScattergeoMarkerSymbol = "diamond-tall"
    ScattergeoMarkerSymbolNumber123 ScattergeoMarkerSymbol = 123
    ScattergeoMarkerSymbol123 ScattergeoMarkerSymbol = "123"
    ScattergeoMarkerSymbolDiamondTallOpen ScattergeoMarkerSymbol = "diamond-tall-open"
    ScattergeoMarkerSymbolNumber223 ScattergeoMarkerSymbol = 223
    ScattergeoMarkerSymbol223 ScattergeoMarkerSymbol = "223"
    ScattergeoMarkerSymbolDiamondTallDot ScattergeoMarkerSymbol = "diamond-tall-dot"
    ScattergeoMarkerSymbolNumber323 ScattergeoMarkerSymbol = 323
    ScattergeoMarkerSymbol323 ScattergeoMarkerSymbol = "323"
    ScattergeoMarkerSymbolDiamondTallOpenDot ScattergeoMarkerSymbol = "diamond-tall-open-dot"
    ScattergeoMarkerSymbolNumber24 ScattergeoMarkerSymbol = 24
    ScattergeoMarkerSymbol24 ScattergeoMarkerSymbol = "24"
    ScattergeoMarkerSymbolDiamondWide ScattergeoMarkerSymbol = "diamond-wide"
    ScattergeoMarkerSymbolNumber124 ScattergeoMarkerSymbol = 124
    ScattergeoMarkerSymbol124 ScattergeoMarkerSymbol = "124"
    ScattergeoMarkerSymbolDiamondWideOpen ScattergeoMarkerSymbol = "diamond-wide-open"
    ScattergeoMarkerSymbolNumber224 ScattergeoMarkerSymbol = 224
    ScattergeoMarkerSymbol224 ScattergeoMarkerSymbol = "224"
    ScattergeoMarkerSymbolDiamondWideDot ScattergeoMarkerSymbol = "diamond-wide-dot"
    ScattergeoMarkerSymbolNumber324 ScattergeoMarkerSymbol = 324
    ScattergeoMarkerSymbol324 ScattergeoMarkerSymbol = "324"
    ScattergeoMarkerSymbolDiamondWideOpenDot ScattergeoMarkerSymbol = "diamond-wide-open-dot"
    ScattergeoMarkerSymbolNumber25 ScattergeoMarkerSymbol = 25
    ScattergeoMarkerSymbol25 ScattergeoMarkerSymbol = "25"
    ScattergeoMarkerSymbolHourglass ScattergeoMarkerSymbol = "hourglass"
    ScattergeoMarkerSymbolNumber125 ScattergeoMarkerSymbol = 125
    ScattergeoMarkerSymbol125 ScattergeoMarkerSymbol = "125"
    ScattergeoMarkerSymbolHourglassOpen ScattergeoMarkerSymbol = "hourglass-open"
    ScattergeoMarkerSymbolNumber26 ScattergeoMarkerSymbol = 26
    ScattergeoMarkerSymbol26 ScattergeoMarkerSymbol = "26"
    ScattergeoMarkerSymbolBowtie ScattergeoMarkerSymbol = "bowtie"
    ScattergeoMarkerSymbolNumber126 ScattergeoMarkerSymbol = 126
    ScattergeoMarkerSymbol126 ScattergeoMarkerSymbol = "126"
    ScattergeoMarkerSymbolBowtieOpen ScattergeoMarkerSymbol = "bowtie-open"
    ScattergeoMarkerSymbolNumber27 ScattergeoMarkerSymbol = 27
    ScattergeoMarkerSymbol27 ScattergeoMarkerSymbol = "27"
    ScattergeoMarkerSymbolCircleCross ScattergeoMarkerSymbol = "circle-cross"
    ScattergeoMarkerSymbolNumber127 ScattergeoMarkerSymbol = 127
    ScattergeoMarkerSymbol127 ScattergeoMarkerSymbol = "127"
    ScattergeoMarkerSymbolCircleCrossOpen ScattergeoMarkerSymbol = "circle-cross-open"
    ScattergeoMarkerSymbolNumber28 ScattergeoMarkerSymbol = 28
    ScattergeoMarkerSymbol28 ScattergeoMarkerSymbol = "28"
    ScattergeoMarkerSymbolCircleX ScattergeoMarkerSymbol = "circle-x"
    ScattergeoMarkerSymbolNumber128 ScattergeoMarkerSymbol = 128
    ScattergeoMarkerSymbol128 ScattergeoMarkerSymbol = "128"
    ScattergeoMarkerSymbolCircleXOpen ScattergeoMarkerSymbol = "circle-x-open"
    ScattergeoMarkerSymbolNumber29 ScattergeoMarkerSymbol = 29
    ScattergeoMarkerSymbol29 ScattergeoMarkerSymbol = "29"
    ScattergeoMarkerSymbolSquareCross ScattergeoMarkerSymbol = "square-cross"
    ScattergeoMarkerSymbolNumber129 ScattergeoMarkerSymbol = 129
    ScattergeoMarkerSymbol129 ScattergeoMarkerSymbol = "129"
    ScattergeoMarkerSymbolSquareCrossOpen ScattergeoMarkerSymbol = "square-cross-open"
    ScattergeoMarkerSymbolNumber30 ScattergeoMarkerSymbol = 30
    ScattergeoMarkerSymbol30 ScattergeoMarkerSymbol = "30"
    ScattergeoMarkerSymbolSquareX ScattergeoMarkerSymbol = "square-x"
    ScattergeoMarkerSymbolNumber130 ScattergeoMarkerSymbol = 130
    ScattergeoMarkerSymbol130 ScattergeoMarkerSymbol = "130"
    ScattergeoMarkerSymbolSquareXOpen ScattergeoMarkerSymbol = "square-x-open"
    ScattergeoMarkerSymbolNumber31 ScattergeoMarkerSymbol = 31
    ScattergeoMarkerSymbol31 ScattergeoMarkerSymbol = "31"
    ScattergeoMarkerSymbolDiamondCross ScattergeoMarkerSymbol = "diamond-cross"
    ScattergeoMarkerSymbolNumber131 ScattergeoMarkerSymbol = 131
    ScattergeoMarkerSymbol131 ScattergeoMarkerSymbol = "131"
    ScattergeoMarkerSymbolDiamondCrossOpen ScattergeoMarkerSymbol = "diamond-cross-open"
    ScattergeoMarkerSymbolNumber32 ScattergeoMarkerSymbol = 32
    ScattergeoMarkerSymbol32 ScattergeoMarkerSymbol = "32"
    ScattergeoMarkerSymbolDiamondX ScattergeoMarkerSymbol = "diamond-x"
    ScattergeoMarkerSymbolNumber132 ScattergeoMarkerSymbol = 132
    ScattergeoMarkerSymbol132 ScattergeoMarkerSymbol = "132"
    ScattergeoMarkerSymbolDiamondXOpen ScattergeoMarkerSymbol = "diamond-x-open"
    ScattergeoMarkerSymbolNumber33 ScattergeoMarkerSymbol = 33
    ScattergeoMarkerSymbol33 ScattergeoMarkerSymbol = "33"
    ScattergeoMarkerSymbolCrossThin ScattergeoMarkerSymbol = "cross-thin"
    ScattergeoMarkerSymbolNumber133 ScattergeoMarkerSymbol = 133
    ScattergeoMarkerSymbol133 ScattergeoMarkerSymbol = "133"
    ScattergeoMarkerSymbolCrossThinOpen ScattergeoMarkerSymbol = "cross-thin-open"
    ScattergeoMarkerSymbolNumber34 ScattergeoMarkerSymbol = 34
    ScattergeoMarkerSymbol34 ScattergeoMarkerSymbol = "34"
    ScattergeoMarkerSymbolXThin ScattergeoMarkerSymbol = "x-thin"
    ScattergeoMarkerSymbolNumber134 ScattergeoMarkerSymbol = 134
    ScattergeoMarkerSymbol134 ScattergeoMarkerSymbol = "134"
    ScattergeoMarkerSymbolXThinOpen ScattergeoMarkerSymbol = "x-thin-open"
    ScattergeoMarkerSymbolNumber35 ScattergeoMarkerSymbol = 35
    ScattergeoMarkerSymbol35 ScattergeoMarkerSymbol = "35"
    ScattergeoMarkerSymbolAsterisk ScattergeoMarkerSymbol = "asterisk"
    ScattergeoMarkerSymbolNumber135 ScattergeoMarkerSymbol = 135
    ScattergeoMarkerSymbol135 ScattergeoMarkerSymbol = "135"
    ScattergeoMarkerSymbolAsteriskOpen ScattergeoMarkerSymbol = "asterisk-open"
    ScattergeoMarkerSymbolNumber36 ScattergeoMarkerSymbol = 36
    ScattergeoMarkerSymbol36 ScattergeoMarkerSymbol = "36"
    ScattergeoMarkerSymbolHash ScattergeoMarkerSymbol = "hash"
    ScattergeoMarkerSymbolNumber136 ScattergeoMarkerSymbol = 136
    ScattergeoMarkerSymbol136 ScattergeoMarkerSymbol = "136"
    ScattergeoMarkerSymbolHashOpen ScattergeoMarkerSymbol = "hash-open"
    ScattergeoMarkerSymbolNumber236 ScattergeoMarkerSymbol = 236
    ScattergeoMarkerSymbol236 ScattergeoMarkerSymbol = "236"
    ScattergeoMarkerSymbolHashDot ScattergeoMarkerSymbol = "hash-dot"
    ScattergeoMarkerSymbolNumber336 ScattergeoMarkerSymbol = 336
    ScattergeoMarkerSymbol336 ScattergeoMarkerSymbol = "336"
    ScattergeoMarkerSymbolHashOpenDot ScattergeoMarkerSymbol = "hash-open-dot"
    ScattergeoMarkerSymbolNumber37 ScattergeoMarkerSymbol = 37
    ScattergeoMarkerSymbol37 ScattergeoMarkerSymbol = "37"
    ScattergeoMarkerSymbolYUp ScattergeoMarkerSymbol = "y-up"
    ScattergeoMarkerSymbolNumber137 ScattergeoMarkerSymbol = 137
    ScattergeoMarkerSymbol137 ScattergeoMarkerSymbol = "137"
    ScattergeoMarkerSymbolYUpOpen ScattergeoMarkerSymbol = "y-up-open"
    ScattergeoMarkerSymbolNumber38 ScattergeoMarkerSymbol = 38
    ScattergeoMarkerSymbol38 ScattergeoMarkerSymbol = "38"
    ScattergeoMarkerSymbolYDown ScattergeoMarkerSymbol = "y-down"
    ScattergeoMarkerSymbolNumber138 ScattergeoMarkerSymbol = 138
    ScattergeoMarkerSymbol138 ScattergeoMarkerSymbol = "138"
    ScattergeoMarkerSymbolYDownOpen ScattergeoMarkerSymbol = "y-down-open"
    ScattergeoMarkerSymbolNumber39 ScattergeoMarkerSymbol = 39
    ScattergeoMarkerSymbol39 ScattergeoMarkerSymbol = "39"
    ScattergeoMarkerSymbolYLeft ScattergeoMarkerSymbol = "y-left"
    ScattergeoMarkerSymbolNumber139 ScattergeoMarkerSymbol = 139
    ScattergeoMarkerSymbol139 ScattergeoMarkerSymbol = "139"
    ScattergeoMarkerSymbolYLeftOpen ScattergeoMarkerSymbol = "y-left-open"
    ScattergeoMarkerSymbolNumber40 ScattergeoMarkerSymbol = 40
    ScattergeoMarkerSymbol40 ScattergeoMarkerSymbol = "40"
    ScattergeoMarkerSymbolYRight ScattergeoMarkerSymbol = "y-right"
    ScattergeoMarkerSymbolNumber140 ScattergeoMarkerSymbol = 140
    ScattergeoMarkerSymbol140 ScattergeoMarkerSymbol = "140"
    ScattergeoMarkerSymbolYRightOpen ScattergeoMarkerSymbol = "y-right-open"
    ScattergeoMarkerSymbolNumber41 ScattergeoMarkerSymbol = 41
    ScattergeoMarkerSymbol41 ScattergeoMarkerSymbol = "41"
    ScattergeoMarkerSymbolLineEw ScattergeoMarkerSymbol = "line-ew"
    ScattergeoMarkerSymbolNumber141 ScattergeoMarkerSymbol = 141
    ScattergeoMarkerSymbol141 ScattergeoMarkerSymbol = "141"
    ScattergeoMarkerSymbolLineEwOpen ScattergeoMarkerSymbol = "line-ew-open"
    ScattergeoMarkerSymbolNumber42 ScattergeoMarkerSymbol = 42
    ScattergeoMarkerSymbol42 ScattergeoMarkerSymbol = "42"
    ScattergeoMarkerSymbolLineNs ScattergeoMarkerSymbol = "line-ns"
    ScattergeoMarkerSymbolNumber142 ScattergeoMarkerSymbol = 142
    ScattergeoMarkerSymbol142 ScattergeoMarkerSymbol = "142"
    ScattergeoMarkerSymbolLineNsOpen ScattergeoMarkerSymbol = "line-ns-open"
    ScattergeoMarkerSymbolNumber43 ScattergeoMarkerSymbol = 43
    ScattergeoMarkerSymbol43 ScattergeoMarkerSymbol = "43"
    ScattergeoMarkerSymbolLineNe ScattergeoMarkerSymbol = "line-ne"
    ScattergeoMarkerSymbolNumber143 ScattergeoMarkerSymbol = 143
    ScattergeoMarkerSymbol143 ScattergeoMarkerSymbol = "143"
    ScattergeoMarkerSymbolLineNeOpen ScattergeoMarkerSymbol = "line-ne-open"
    ScattergeoMarkerSymbolNumber44 ScattergeoMarkerSymbol = 44
    ScattergeoMarkerSymbol44 ScattergeoMarkerSymbol = "44"
    ScattergeoMarkerSymbolLineNw ScattergeoMarkerSymbol = "line-nw"
    ScattergeoMarkerSymbolNumber144 ScattergeoMarkerSymbol = 144
    ScattergeoMarkerSymbol144 ScattergeoMarkerSymbol = "144"
    ScattergeoMarkerSymbolLineNwOpen ScattergeoMarkerSymbol = "line-nw-open"
    ScattergeoMarkerSymbolNumber45 ScattergeoMarkerSymbol = 45
    ScattergeoMarkerSymbol45 ScattergeoMarkerSymbol = "45"
    ScattergeoMarkerSymbolArrowUp ScattergeoMarkerSymbol = "arrow-up"
    ScattergeoMarkerSymbolNumber145 ScattergeoMarkerSymbol = 145
    ScattergeoMarkerSymbol145 ScattergeoMarkerSymbol = "145"
    ScattergeoMarkerSymbolArrowUpOpen ScattergeoMarkerSymbol = "arrow-up-open"
    ScattergeoMarkerSymbolNumber46 ScattergeoMarkerSymbol = 46
    ScattergeoMarkerSymbol46 ScattergeoMarkerSymbol = "46"
    ScattergeoMarkerSymbolArrowDown ScattergeoMarkerSymbol = "arrow-down"
    ScattergeoMarkerSymbolNumber146 ScattergeoMarkerSymbol = 146
    ScattergeoMarkerSymbol146 ScattergeoMarkerSymbol = "146"
    ScattergeoMarkerSymbolArrowDownOpen ScattergeoMarkerSymbol = "arrow-down-open"
    ScattergeoMarkerSymbolNumber47 ScattergeoMarkerSymbol = 47
    ScattergeoMarkerSymbol47 ScattergeoMarkerSymbol = "47"
    ScattergeoMarkerSymbolArrowLeft ScattergeoMarkerSymbol = "arrow-left"
    ScattergeoMarkerSymbolNumber147 ScattergeoMarkerSymbol = 147
    ScattergeoMarkerSymbol147 ScattergeoMarkerSymbol = "147"
    ScattergeoMarkerSymbolArrowLeftOpen ScattergeoMarkerSymbol = "arrow-left-open"
    ScattergeoMarkerSymbolNumber48 ScattergeoMarkerSymbol = 48
    ScattergeoMarkerSymbol48 ScattergeoMarkerSymbol = "48"
    ScattergeoMarkerSymbolArrowRight ScattergeoMarkerSymbol = "arrow-right"
    ScattergeoMarkerSymbolNumber148 ScattergeoMarkerSymbol = 148
    ScattergeoMarkerSymbol148 ScattergeoMarkerSymbol = "148"
    ScattergeoMarkerSymbolArrowRightOpen ScattergeoMarkerSymbol = "arrow-right-open"
    ScattergeoMarkerSymbolNumber49 ScattergeoMarkerSymbol = 49
    ScattergeoMarkerSymbol49 ScattergeoMarkerSymbol = "49"
    ScattergeoMarkerSymbolArrowBarUp ScattergeoMarkerSymbol = "arrow-bar-up"
    ScattergeoMarkerSymbolNumber149 ScattergeoMarkerSymbol = 149
    ScattergeoMarkerSymbol149 ScattergeoMarkerSymbol = "149"
    ScattergeoMarkerSymbolArrowBarUpOpen ScattergeoMarkerSymbol = "arrow-bar-up-open"
    ScattergeoMarkerSymbolNumber50 ScattergeoMarkerSymbol = 50
    ScattergeoMarkerSymbol50 ScattergeoMarkerSymbol = "50"
    ScattergeoMarkerSymbolArrowBarDown ScattergeoMarkerSymbol = "arrow-bar-down"
    ScattergeoMarkerSymbolNumber150 ScattergeoMarkerSymbol = 150
    ScattergeoMarkerSymbol150 ScattergeoMarkerSymbol = "150"
    ScattergeoMarkerSymbolArrowBarDownOpen ScattergeoMarkerSymbol = "arrow-bar-down-open"
    ScattergeoMarkerSymbolNumber51 ScattergeoMarkerSymbol = 51
    ScattergeoMarkerSymbol51 ScattergeoMarkerSymbol = "51"
    ScattergeoMarkerSymbolArrowBarLeft ScattergeoMarkerSymbol = "arrow-bar-left"
    ScattergeoMarkerSymbolNumber151 ScattergeoMarkerSymbol = 151
    ScattergeoMarkerSymbol151 ScattergeoMarkerSymbol = "151"
    ScattergeoMarkerSymbolArrowBarLeftOpen ScattergeoMarkerSymbol = "arrow-bar-left-open"
    ScattergeoMarkerSymbolNumber52 ScattergeoMarkerSymbol = 52
    ScattergeoMarkerSymbol52 ScattergeoMarkerSymbol = "52"
    ScattergeoMarkerSymbolArrowBarRight ScattergeoMarkerSymbol = "arrow-bar-right"
    ScattergeoMarkerSymbolNumber152 ScattergeoMarkerSymbol = 152
    ScattergeoMarkerSymbol152 ScattergeoMarkerSymbol = "152"
    ScattergeoMarkerSymbolArrowBarRightOpen ScattergeoMarkerSymbol = "arrow-bar-right-open"
    
)
// ScattergeoTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScattergeoTextposition string 

const (
    ScattergeoTextpositionTopLeft ScattergeoTextposition = "top left"
    ScattergeoTextpositionTopCenter ScattergeoTextposition = "top center"
    ScattergeoTextpositionTopRight ScattergeoTextposition = "top right"
    ScattergeoTextpositionMiddleLeft ScattergeoTextposition = "middle left"
    ScattergeoTextpositionMiddleCenter ScattergeoTextposition = "middle center"
    ScattergeoTextpositionMiddleRight ScattergeoTextposition = "middle right"
    ScattergeoTextpositionBottomLeft ScattergeoTextposition = "bottom left"
    ScattergeoTextpositionBottomCenter ScattergeoTextposition = "bottom center"
    ScattergeoTextpositionBottomRight ScattergeoTextposition = "bottom right"
    
)
// ScattergeoVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScattergeoVisible interface{} 

var (
    ScattergeoVisibleTrue ScattergeoVisible = true
    ScattergeoVisibleFalse ScattergeoVisible = false
    ScattergeoVisibleLegendonly ScattergeoVisible = "legendonly"
    
)
// ScattergeoHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type ScattergeoHoverinfo string 

const (
    // Flags
    ScattergeoHoverinfoLon ScattergeoHoverinfo = "lon"
    ScattergeoHoverinfoLat ScattergeoHoverinfo = "lat"
    ScattergeoHoverinfoLocation ScattergeoHoverinfo = "location"
    ScattergeoHoverinfoText ScattergeoHoverinfo = "text"
    ScattergeoHoverinfoName ScattergeoHoverinfo = "name"
    
    // Extra
    ScattergeoHoverinfoAll ScattergeoHoverinfo = "all"
    ScattergeoHoverinfoNone ScattergeoHoverinfo = "none"
    ScattergeoHoverinfoSkip ScattergeoHoverinfo = "skip"
    
)
// ScattergeoMode Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
type ScattergeoMode string 

const (
    // Flags
    ScattergeoModeLines ScattergeoMode = "lines"
    ScattergeoModeMarkers ScattergeoMode = "markers"
    ScattergeoModeText ScattergeoMode = "text"
    
    // Extra
    ScattergeoModeNone ScattergeoMode = "none"
    
)
