
package graph_objects


type Layout struct {

    // _deprecated <no value> <no value>
    // Pending... _deprecated No valTyp <no value>
    // Activeshape <no value> <no value>
    Activeshape *LayoutActiveshape `json:"activeshape,omitempty"` // object
    // Angularaxis <no value> <no value>
    Angularaxis *LayoutAngularaxis `json:"angularaxis,omitempty"` // object
    // Annotations <no value> <no value>
    Annotations *LayoutAnnotations `json:"annotations,omitempty"` // object
    // Autosize boolean Determines whether or not a layout width or height that has been left undefined by the user is initialized on each relayout. Note that, regardless of this attribute, an undefined layout width or height is always initialized on the first call to plot.
    Autosize Bool `json:"autosize,omitempty"`
    
    // Bargap number Sets the gap (in plot fraction) between bars of adjacent location coordinates.
    Bargap float64 `json:"bargap,omitempty"`
    
    // Bargroupgap number Sets the gap (in plot fraction) between bars of the same location coordinate.
    Bargroupgap float64 `json:"bargroupgap,omitempty"`
    
    // Barmode enumerated Determines how bars at the same location coordinate are displayed on the graph. With *stack*, the bars are stacked on top of one another With *relative*, the bars are stacked on top of one another, with negative values below the axis, positive values above With *group*, the bars are plotted next to one another centered around the shared location. With *overlay*, the bars are plotted over one another, you might need to an *opacity* to see multiple bars.
    Barmode LayoutBarmode `json:"barmode,omitempty"`
    
    // Barnorm enumerated Sets the normalization for bar traces on the graph. With *fraction*, the value of each bar is divided by the sum of all values at that location coordinate. *percent* is the same but multiplied by 100 to show percentages.
    Barnorm LayoutBarnorm `json:"barnorm,omitempty"`
    
    // Boxgap number Sets the gap (in plot fraction) between boxes of adjacent location coordinates. Has no effect on traces that have *width* set.
    Boxgap float64 `json:"boxgap,omitempty"`
    
    // Boxgroupgap number Sets the gap (in plot fraction) between boxes of the same location coordinate. Has no effect on traces that have *width* set.
    Boxgroupgap float64 `json:"boxgroupgap,omitempty"`
    
    // Boxmode enumerated Determines how boxes at the same location coordinate are displayed on the graph. If *group*, the boxes are plotted next to one another centered around the shared location. If *overlay*, the boxes are plotted over one another, you might need to set *opacity* to see them multiple boxes. Has no effect on traces that have *width* set.
    Boxmode LayoutBoxmode `json:"boxmode,omitempty"`
    
    // Calendar enumerated Sets the default calendar system to use for interpreting and displaying dates throughout the plot.
    Calendar LayoutCalendar `json:"calendar,omitempty"`
    
    // Clickmode flaglist Determines the mode of single click interactions. *event* is the default value and emits the `plotly_click` event. In addition this mode emits the `plotly_selected` event in drag modes *lasso* and *select*, but with no event data attached (kept for compatibility reasons). The *select* flag enables selecting single data points via click. This mode also supports persistent selections, meaning that pressing Shift while clicking, adds to / subtracts from an existing selection. *select* with `hovermode`: *x* can be confusing, consider explicitly setting `hovermode`: *closest* when using this feature. Selection events are sent accordingly as long as *event* flag is set as well. When the *event* flag is missing, `plotly_click` and `plotly_selected` events are not fired.
    Clickmode LayoutClickmode `json:"clickmode,omitempty"`
    
    // Coloraxis <no value> 
    Coloraxis *LayoutColoraxis `json:"coloraxis,omitempty"` // object
    // Colorscale <no value> <no value>
    Colorscale *LayoutColorscale `json:"colorscale,omitempty"` // object
    // Colorway colorlist Sets the default trace colors.
    // Pending of type "colorlist"Colorway  `json:"colorway,omitempty"`
    
    // Datarevision any If provided, a changed value tells `Plotly.react` that one or more data arrays has changed. This way you can modify arrays in-place rather than making a complete new copy for an incremental change. If NOT provided, `Plotly.react` assumes that data arrays are being treated as immutable, thus any data array with a different identity from its predecessor contains new data.
    Datarevision interface{} `json:"datarevision,omitempty"`
    
    // Direction enumerated Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the direction corresponding to positive angles in legacy polar charts.
    Direction LayoutDirection `json:"direction,omitempty"`
    
    // Dragmode enumerated Determines the mode of drag interactions. *select* and *lasso* apply only to scatter traces with markers or text. *orbit* and *turntable* apply only to 3D scenes.
    Dragmode LayoutDragmode `json:"dragmode,omitempty"`
    
    // Editrevision any Controls persistence of user-driven changes in `editable: true` configuration, other than trace names and axis titles. Defaults to `layout.uirevision`.
    Editrevision interface{} `json:"editrevision,omitempty"`
    
    // Extendfunnelareacolors boolean If `true`, the funnelarea slice colors (whether given by `funnelareacolorway` or inherited from `colorway`) will be extended to three times its original length by first repeating every color 20% lighter then each color 20% darker. This is intended to reduce the likelihood of reusing the same color when you have many slices, but you can set `false` to disable. Colors provided in the trace, using `marker.colors`, are never extended.
    Extendfunnelareacolors Bool `json:"extendfunnelareacolors,omitempty"`
    
    // Extendpiecolors boolean If `true`, the pie slice colors (whether given by `piecolorway` or inherited from `colorway`) will be extended to three times its original length by first repeating every color 20% lighter then each color 20% darker. This is intended to reduce the likelihood of reusing the same color when you have many slices, but you can set `false` to disable. Colors provided in the trace, using `marker.colors`, are never extended.
    Extendpiecolors Bool `json:"extendpiecolors,omitempty"`
    
    // Extendsunburstcolors boolean If `true`, the sunburst slice colors (whether given by `sunburstcolorway` or inherited from `colorway`) will be extended to three times its original length by first repeating every color 20% lighter then each color 20% darker. This is intended to reduce the likelihood of reusing the same color when you have many slices, but you can set `false` to disable. Colors provided in the trace, using `marker.colors`, are never extended.
    Extendsunburstcolors Bool `json:"extendsunburstcolors,omitempty"`
    
    // Extendtreemapcolors boolean If `true`, the treemap slice colors (whether given by `treemapcolorway` or inherited from `colorway`) will be extended to three times its original length by first repeating every color 20% lighter then each color 20% darker. This is intended to reduce the likelihood of reusing the same color when you have many slices, but you can set `false` to disable. Colors provided in the trace, using `marker.colors`, are never extended.
    Extendtreemapcolors Bool `json:"extendtreemapcolors,omitempty"`
    
    // Font <no value> Sets the global font. Note that fonts used in traces and other layout components inherit from the global font.
    Font *LayoutFont `json:"font,omitempty"` // object
    // Funnelareacolorway colorlist Sets the default funnelarea slice colors. Defaults to the main `colorway` used for trace colors. If you specify a new list here it can still be extended with lighter and darker colors, see `extendfunnelareacolors`.
    // Pending of type "colorlist"Funnelareacolorway  `json:"funnelareacolorway,omitempty"`
    
    // Funnelgap number Sets the gap (in plot fraction) between bars of adjacent location coordinates.
    Funnelgap float64 `json:"funnelgap,omitempty"`
    
    // Funnelgroupgap number Sets the gap (in plot fraction) between bars of the same location coordinate.
    Funnelgroupgap float64 `json:"funnelgroupgap,omitempty"`
    
    // Funnelmode enumerated Determines how bars at the same location coordinate are displayed on the graph. With *stack*, the bars are stacked on top of one another With *group*, the bars are plotted next to one another centered around the shared location. With *overlay*, the bars are plotted over one another, you might need to an *opacity* to see multiple bars.
    Funnelmode LayoutFunnelmode `json:"funnelmode,omitempty"`
    
    // Geo <no value> <no value>
    Geo *LayoutGeo `json:"geo,omitempty"` // object
    // Grid <no value> <no value>
    Grid *LayoutGrid `json:"grid,omitempty"` // object
    // Height number Sets the plot's height (in px).
    Height float64 `json:"height,omitempty"`
    
    // Hiddenlabels data_array hiddenlabels is the funnelarea & pie chart analog of visible:'legendonly' but it can contain many labels, and can simultaneously hide slices from several pies/funnelarea charts
    Hiddenlabels interface{} `json:"hiddenlabels,omitempty"`
    
    // Hiddenlabelssrc string Sets the source reference on Chart Studio Cloud for  hiddenlabels .
    Hiddenlabelssrc String `json:"hiddenlabelssrc,omitempty"`
    
    // Hidesources boolean Determines whether or not a text link citing the data source is placed at the bottom-right cored of the figure. Has only an effect only on graphs that have been generated via forked graphs from the Chart Studio Cloud (at https://chart-studio.plotly.com or on-premise).
    Hidesources Bool `json:"hidesources,omitempty"`
    
    // Hoverdistance integer Sets the default distance (in pixels) to look for data to add hover labels (-1 means no cutoff, 0 means no looking for data). This is only a real distance for hovering on point-like objects, like scatter points. For area-like objects (bars, scatter fills, etc) hovering is on inside the area and off outside, but these objects will not supersede hover on point-like objects in case of conflict.
    Hoverdistance int64 `json:"hoverdistance,omitempty"`
    
    // Hoverlabel <no value> <no value>
    Hoverlabel *LayoutHoverlabel `json:"hoverlabel,omitempty"` // object
    // Hovermode enumerated Determines the mode of hover interactions. If *closest*, a single hoverlabel will appear for the *closest* point within the `hoverdistance`. If *x* (or *y*), multiple hoverlabels will appear for multiple points at the *closest* x- (or y-) coordinate within the `hoverdistance`, with the caveat that no more than one hoverlabel will appear per trace. If *x unified* (or *y unified*), a single hoverlabel will appear multiple points at the closest x- (or y-) coordinate within the `hoverdistance` with the caveat that no more than one hoverlabel will appear per trace. In this mode, spikelines are enabled by default perpendicular to the specified axis. If false, hover interactions are disabled. If `clickmode` includes the *select* flag, `hovermode` defaults to *closest*. If `clickmode` lacks the *select* flag, it defaults to *x* or *y* (depending on the trace's `orientation` value) for plots based on cartesian coordinates. For anything else the default value is *closest*.
    Hovermode LayoutHovermode `json:"hovermode,omitempty"`
    
    // Images <no value> <no value>
    Images *LayoutImages `json:"images,omitempty"` // object
    // Legend <no value> <no value>
    Legend *LayoutLegend `json:"legend,omitempty"` // object
    // Mapbox <no value> <no value>
    Mapbox *LayoutMapbox `json:"mapbox,omitempty"` // object
    // Margin <no value> <no value>
    Margin *LayoutMargin `json:"margin,omitempty"` // object
    // Meta any Assigns extra meta information that can be used in various `text` attributes. Attributes such as the graph, axis and colorbar `title.text`, annotation `text` `trace.name` in legend items, `rangeselector`, `updatemenus` and `sliders` `label` text all support `meta`. One can access `meta` fields using template strings: `%{meta[i]}` where `i` is the index of the `meta` item in question. `meta` can also be an object for example `{key: value}` which can be accessed %{meta[key]}.
    Meta interface{} `json:"meta,omitempty"`
    
    // Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
    Metasrc String `json:"metasrc,omitempty"`
    
    // Modebar <no value> <no value>
    Modebar *LayoutModebar `json:"modebar,omitempty"` // object
    // Newshape <no value> <no value>
    Newshape *LayoutNewshape `json:"newshape,omitempty"` // object
    // Orientation angle Legacy polar charts are deprecated! Please switch to *polar* subplots. Rotates the entire polar by the given angle in legacy polar charts.
    Orientation float64 `json:"orientation,omitempty"`
    
    // Paper_bgcolor color Sets the background color of the paper where the graph is drawn.
    Paper_bgcolor String `json:"paper_bgcolor,omitempty"`
    
    // Piecolorway colorlist Sets the default pie slice colors. Defaults to the main `colorway` used for trace colors. If you specify a new list here it can still be extended with lighter and darker colors, see `extendpiecolors`.
    // Pending of type "colorlist"Piecolorway  `json:"piecolorway,omitempty"`
    
    // Plot_bgcolor color Sets the background color of the plotting area in-between x and y axes.
    Plot_bgcolor String `json:"plot_bgcolor,omitempty"`
    
    // Polar <no value> <no value>
    Polar *LayoutPolar `json:"polar,omitempty"` // object
    // Radialaxis <no value> <no value>
    Radialaxis *LayoutRadialaxis `json:"radialaxis,omitempty"` // object
    // Scene <no value> <no value>
    Scene *LayoutScene `json:"scene,omitempty"` // object
    // Selectdirection enumerated When `dragmode` is set to *select*, this limits the selection of the drag to horizontal, vertical or diagonal. *h* only allows horizontal selection, *v* only vertical, *d* only diagonal and *any* sets no limit.
    Selectdirection LayoutSelectdirection `json:"selectdirection,omitempty"`
    
    // Selectionrevision any Controls persistence of user-driven changes in selected points from all traces.
    Selectionrevision interface{} `json:"selectionrevision,omitempty"`
    
    // Separators string Sets the decimal and thousand separators. For example, *. * puts a '.' before decimals and a space between thousands. In English locales, dflt is *.,* but other locales may alter this default.
    Separators String `json:"separators,omitempty"`
    
    // Shapes <no value> <no value>
    Shapes *LayoutShapes `json:"shapes,omitempty"` // object
    // Showlegend boolean Determines whether or not a legend is drawn. Default is `true` if there is a trace to show and any of these: a) Two or more traces would by default be shown in the legend. b) One pie trace is shown in the legend. c) One trace is explicitly given with `showlegend: true`.
    Showlegend Bool `json:"showlegend,omitempty"`
    
    // Sliders <no value> <no value>
    Sliders *LayoutSliders `json:"sliders,omitempty"` // object
    // Spikedistance integer Sets the default distance (in pixels) to look for data to draw spikelines to (-1 means no cutoff, 0 means no looking for data). As with hoverdistance, distance does not apply to area-like objects. In addition, some objects can be hovered on but will not generate spikelines, such as scatter fills.
    Spikedistance int64 `json:"spikedistance,omitempty"`
    
    // Sunburstcolorway colorlist Sets the default sunburst slice colors. Defaults to the main `colorway` used for trace colors. If you specify a new list here it can still be extended with lighter and darker colors, see `extendsunburstcolors`.
    // Pending of type "colorlist"Sunburstcolorway  `json:"sunburstcolorway,omitempty"`
    
    // Template any Default attributes to be applied to the plot. Templates can be created from existing plots using `Plotly.makeTemplate`, or created manually. They should be objects with format: `{layout: layoutTemplate, data: {[type]: [traceTemplate, ...]}, ...}` `layoutTemplate` and `traceTemplate` are objects matching the attribute structure of `layout` and a data trace.  Trace templates are applied cyclically to traces of each type. Container arrays (eg `annotations`) have special handling: An object ending in `defaults` (eg `annotationdefaults`) is applied to each array item. But if an item has a `templateitemname` key we look in the template array for an item with matching `name` and apply that instead. If no matching `name` is found we mark the item invisible. Any named template item not referenced is appended to the end of the array, so you can use this for a watermark annotation or a logo image, for example. To omit one of these items on the plot, make an item with matching `templateitemname` and `visible: false`.
    Template interface{} `json:"template,omitempty"`
    
    // Ternary <no value> <no value>
    Ternary *LayoutTernary `json:"ternary,omitempty"` // object
    // Title <no value> <no value>
    Title *LayoutTitle `json:"title,omitempty"` // object
    // Transition <no value> Sets transition options used during Plotly.react updates.
    Transition *LayoutTransition `json:"transition,omitempty"` // object
    // Treemapcolorway colorlist Sets the default treemap slice colors. Defaults to the main `colorway` used for trace colors. If you specify a new list here it can still be extended with lighter and darker colors, see `extendtreemapcolors`.
    // Pending of type "colorlist"Treemapcolorway  `json:"treemapcolorway,omitempty"`
    
    // Uirevision any Used to allow user interactions with the plot to persist after `Plotly.react` calls that are unaware of these interactions. If `uirevision` is omitted, or if it is given and it changed from the previous `Plotly.react` call, the exact new figure is used. If `uirevision` is truthy and did NOT change, any attribute that has been affected by user interactions and did not receive a different value in the new figure will keep the interaction value. `layout.uirevision` attribute serves as the default for `uirevision` attributes in various sub-containers. For finer control you can set these sub-attributes directly. For example, if your app separately controls the data on the x and y axes you might set `xaxis.uirevision=*time*` and `yaxis.uirevision=*cost*`. Then if only the y data is changed, you can update `yaxis.uirevision=*quantity*` and the y axis range will reset but the x axis range will retain any user-driven zoom.
    Uirevision interface{} `json:"uirevision,omitempty"`
    
    // Uniformtext <no value> <no value>
    Uniformtext *LayoutUniformtext `json:"uniformtext,omitempty"` // object
    // Updatemenus <no value> <no value>
    Updatemenus *LayoutUpdatemenus `json:"updatemenus,omitempty"` // object
    // Violingap number Sets the gap (in plot fraction) between violins of adjacent location coordinates. Has no effect on traces that have *width* set.
    Violingap float64 `json:"violingap,omitempty"`
    
    // Violingroupgap number Sets the gap (in plot fraction) between violins of the same location coordinate. Has no effect on traces that have *width* set.
    Violingroupgap float64 `json:"violingroupgap,omitempty"`
    
    // Violinmode enumerated Determines how violins at the same location coordinate are displayed on the graph. If *group*, the violins are plotted next to one another centered around the shared location. If *overlay*, the violins are plotted over one another, you might need to set *opacity* to see them multiple violins. Has no effect on traces that have *width* set.
    Violinmode LayoutViolinmode `json:"violinmode,omitempty"`
    
    // Waterfallgap number Sets the gap (in plot fraction) between bars of adjacent location coordinates.
    Waterfallgap float64 `json:"waterfallgap,omitempty"`
    
    // Waterfallgroupgap number Sets the gap (in plot fraction) between bars of the same location coordinate.
    Waterfallgroupgap float64 `json:"waterfallgroupgap,omitempty"`
    
    // Waterfallmode enumerated Determines how bars at the same location coordinate are displayed on the graph. With *group*, the bars are plotted next to one another centered around the shared location. With *overlay*, the bars are plotted over one another, you might need to an *opacity* to see multiple bars.
    Waterfallmode LayoutWaterfallmode `json:"waterfallmode,omitempty"`
    
    // Width number Sets the plot's width (in px).
    Width float64 `json:"width,omitempty"`
    
    // Xaxis <no value> <no value>
    Xaxis *LayoutXaxis `json:"xaxis,omitempty"` // object
    // Yaxis <no value> <no value>
    Yaxis *LayoutYaxis `json:"yaxis,omitempty"` // object


    // Exceptional hardcoded cases due to schema limitation
    // Xaxis2 see Xaxis prop
    Xaxis2 *LayoutXaxis `json:"xaxis2,omitempty"`  
    // Yaxis2 see Yaxis prop
    Yaxis2 *LayoutYaxis `json:"yaxis2,omitempty"` 

    // Xaxis3 see Xaxis prop
    Xaxis3 *LayoutXaxis `json:"xaxis3,omitempty"`  
    // Yaxis3 see Yaxis prop
    Yaxis3 *LayoutYaxis `json:"yaxis3,omitempty"` 

    // Xaxis4 see Xaxis prop
    Xaxis4 *LayoutXaxis `json:"xaxis4,omitempty"`  
    // Yaxis4 see Yaxis prop
    Yaxis4 *LayoutYaxis `json:"yaxis4,omitempty"` 

    // Xaxis5 see Xaxis prop
    Xaxis5 *LayoutXaxis `json:"xaxis5,omitempty"`  
    // Yaxis5 see Yaxis prop
    Yaxis5 *LayoutYaxis `json:"yaxis5,omitempty"` 

    // Xaxis6 see Xaxis prop
    Xaxis6 *LayoutXaxis `json:"xaxis6,omitempty"`  
    // Yaxis6 see Yaxis prop
    Yaxis6 *LayoutYaxis `json:"yaxis6,omitempty"` 
}

