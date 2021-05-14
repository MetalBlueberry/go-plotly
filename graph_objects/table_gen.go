package grob

var TraceTypeTable TraceType = "table"

func (trace *Table) GetType() TraceType {
	return TraceTypeTable
}
// Table Table view for detailed data viewing. The data are arranged in a grid of rows and columns. Most styling can be specified for columns, rows or individual cells. Table is using a column-major order, ie. the grid is represented as a vector of column vectors.
type Table struct {
    
    // Type 
    // is the type of the plot 
    Type TraceType `json:"type,omitempty"`
    
    // Cells 
    //   
    Cells *TableCells `json:"cells,omitempty"`
    
    // Columnorder 
    // data_array 
    // Specifies the rendered order of the data columns; for example, a value `2` at position `0` means that column index `0` in the data will be rendered as the third column, as columns have an index base of zero. 
    Columnorder interface{} `json:"columnorder,omitempty"`
    
    // Columnordersrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  columnorder . 
    Columnordersrc String `json:"columnordersrc,omitempty"`
    
    // Columnwidth 
    // number 
    // The width of columns expressed as a ratio. Columns fill the available width in proportion of their specified column widths. 
    Columnwidth float64 `json:"columnwidth,omitempty"`
    
    // Columnwidthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  columnwidth . 
    Columnwidthsrc String `json:"columnwidthsrc,omitempty"`
    
    // Customdata 
    // data_array 
    // Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements 
    Customdata interface{} `json:"customdata,omitempty"`
    
    // Customdatasrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  customdata . 
    Customdatasrc String `json:"customdatasrc,omitempty"`
    
    // Domain 
    //   
    Domain *TableDomain `json:"domain,omitempty"`
    
    // Header 
    //   
    Header *TableHeader `json:"header,omitempty"`
    
    // Hoverinfo 
    // flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired. 
    Hoverinfo TableHoverinfo `json:"hoverinfo,omitempty"`
    
    // Hoverinfosrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  hoverinfo . 
    Hoverinfosrc String `json:"hoverinfosrc,omitempty"`
    
    // Hoverlabel 
    //   
    Hoverlabel *TableHoverlabel `json:"hoverlabel,omitempty"`
    
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
    
    // Name 
    // string 
    // Sets the trace name. The trace name appear as the legend item and on hover. 
    Name String `json:"name,omitempty"`
    
    // Stream 
    //   
    Stream *TableStream `json:"stream,omitempty"`
    
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
    Visible TableVisible `json:"visible,omitempty"`
    
}
// TableCellsFill 
type TableCellsFill struct {
    
    // Color 
    // color 
    // Sets the cell fill color. It accepts either a specific color or an array of colors or a 2D array of colors. 
    Color Color `json:"color,omitempty"`
    
    // Colorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  color . 
    Colorsrc String `json:"colorsrc,omitempty"`
    
}
// TableCellsFont 
type TableCellsFont struct {
    
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
// TableCellsLine 
type TableCellsLine struct {
    
    // Color 
    // color 
    //  
    Color Color `json:"color,omitempty"`
    
    // Colorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  color . 
    Colorsrc String `json:"colorsrc,omitempty"`
    
    // Width 
    // number 
    //  
    Width float64 `json:"width,omitempty"`
    
    // Widthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  width . 
    Widthsrc String `json:"widthsrc,omitempty"`
    
}
// TableCells 
type TableCells struct {
    
    // Align 
    // enumerated Sets the horizontal alignment of the `text` within the box. Has an effect only if `text` spans two or more lines (i.e. `text` contains one or more <br> HTML tags) or if an explicit width is set to override the text width. 
    Align TableCellsAlign `json:"align,omitempty"`
    
    // Alignsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  align . 
    Alignsrc String `json:"alignsrc,omitempty"`
    
    // Fill 
    //   
    Fill *TableCellsFill `json:"fill,omitempty"`
    
    // Font 
    //   
    Font *TableCellsFont `json:"font,omitempty"`
    
    // Format 
    // data_array 
    // Sets the cell value formatting rule using d3 formatting mini-language which is similar to those of Python. See https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format 
    Format interface{} `json:"format,omitempty"`
    
    // Formatsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  format . 
    Formatsrc String `json:"formatsrc,omitempty"`
    
    // Height 
    // number 
    // The height of cells. 
    Height float64 `json:"height,omitempty"`
    
    // Line 
    //   
    Line *TableCellsLine `json:"line,omitempty"`
    
    // Prefix 
    // string 
    // Prefix for cell values. 
    Prefix String `json:"prefix,omitempty"`
    
    // Prefixsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  prefix . 
    Prefixsrc String `json:"prefixsrc,omitempty"`
    
    // Suffix 
    // string 
    // Suffix for cell values. 
    Suffix String `json:"suffix,omitempty"`
    
    // Suffixsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  suffix . 
    Suffixsrc String `json:"suffixsrc,omitempty"`
    
    // Values 
    // data_array 
    // Cell values. `values[m][n]` represents the value of the `n`th point in column `m`, therefore the `values[m]` vector length for all columns must be the same (longer vectors will be truncated). Each value must be a finite number or a string. 
    Values interface{} `json:"values,omitempty"`
    
    // Valuessrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  values . 
    Valuessrc String `json:"valuessrc,omitempty"`
    
}
// TableDomain 
type TableDomain struct {
    
    // Column 
    // integer 
    // If there is a layout grid, use the domain for this column in the grid for this table trace . 
    Column int64 `json:"column,omitempty"`
    
    // Row 
    // integer 
    // If there is a layout grid, use the domain for this row in the grid for this table trace . 
    Row int64 `json:"row,omitempty"`
    
    // X 
    // info_array 
    // Sets the horizontal domain of this table trace (in plot fraction). 
    X interface{} `json:"x,omitempty"`
    
    // Y 
    // info_array 
    // Sets the vertical domain of this table trace (in plot fraction). 
    Y interface{} `json:"y,omitempty"`
    
}
// TableHeaderFill 
type TableHeaderFill struct {
    
    // Color 
    // color 
    // Sets the cell fill color. It accepts either a specific color or an array of colors or a 2D array of colors. 
    Color Color `json:"color,omitempty"`
    
    // Colorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  color . 
    Colorsrc String `json:"colorsrc,omitempty"`
    
}
// TableHeaderFont 
type TableHeaderFont struct {
    
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
// TableHeaderLine 
type TableHeaderLine struct {
    
    // Color 
    // color 
    //  
    Color Color `json:"color,omitempty"`
    
    // Colorsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  color . 
    Colorsrc String `json:"colorsrc,omitempty"`
    
    // Width 
    // number 
    //  
    Width float64 `json:"width,omitempty"`
    
    // Widthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  width . 
    Widthsrc String `json:"widthsrc,omitempty"`
    
}
// TableHeader 
type TableHeader struct {
    
    // Align 
    // enumerated Sets the horizontal alignment of the `text` within the box. Has an effect only if `text` spans two or more lines (i.e. `text` contains one or more <br> HTML tags) or if an explicit width is set to override the text width. 
    Align TableHeaderAlign `json:"align,omitempty"`
    
    // Alignsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  align . 
    Alignsrc String `json:"alignsrc,omitempty"`
    
    // Fill 
    //   
    Fill *TableHeaderFill `json:"fill,omitempty"`
    
    // Font 
    //   
    Font *TableHeaderFont `json:"font,omitempty"`
    
    // Format 
    // data_array 
    // Sets the cell value formatting rule using d3 formatting mini-language which is similar to those of Python. See https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format 
    Format interface{} `json:"format,omitempty"`
    
    // Formatsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  format . 
    Formatsrc String `json:"formatsrc,omitempty"`
    
    // Height 
    // number 
    // The height of cells. 
    Height float64 `json:"height,omitempty"`
    
    // Line 
    //   
    Line *TableHeaderLine `json:"line,omitempty"`
    
    // Prefix 
    // string 
    // Prefix for cell values. 
    Prefix String `json:"prefix,omitempty"`
    
    // Prefixsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  prefix . 
    Prefixsrc String `json:"prefixsrc,omitempty"`
    
    // Suffix 
    // string 
    // Suffix for cell values. 
    Suffix String `json:"suffix,omitempty"`
    
    // Suffixsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  suffix . 
    Suffixsrc String `json:"suffixsrc,omitempty"`
    
    // Values 
    // data_array 
    // Header cell values. `values[m][n]` represents the value of the `n`th point in column `m`, therefore the `values[m]` vector length for all columns must be the same (longer vectors will be truncated). Each value must be a finite number or a string. 
    Values interface{} `json:"values,omitempty"`
    
    // Valuessrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  values . 
    Valuessrc String `json:"valuessrc,omitempty"`
    
}
// TableHoverlabelFont Sets the font used in hover labels.
type TableHoverlabelFont struct {
    
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
// TableHoverlabel 
type TableHoverlabel struct {
    
    // Align 
    // enumerated Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines 
    Align TableHoverlabelAlign `json:"align,omitempty"`
    
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
    Font *TableHoverlabelFont `json:"font,omitempty"`
    
    // Namelength 
    // integer 
    // Sets the default length (in number of characters) of the trace name in the hover labels for all traces. -1 shows the whole name regardless of length. 0-3 shows the first 0-3 characters, and an integer >3 will show the whole name if it is less than that many characters, but if it is longer, will truncate to `namelength - 3` characters and add an ellipsis. 
    Namelength int64 `json:"namelength,omitempty"`
    
    // Namelengthsrc 
    // string 
    // Sets the source reference on Chart Studio Cloud for  namelength . 
    Namelengthsrc String `json:"namelengthsrc,omitempty"`
    
}
// TableStream 
type TableStream struct {
    
    // Maxpoints 
    // number 
    // Sets the maximum number of points to keep on the plots from an incoming stream. If `maxpoints` is set to *50*, only the newest 50 points will be displayed on the plot. 
    Maxpoints float64 `json:"maxpoints,omitempty"`
    
    // Token 
    // string 
    // The stream id number links a data trace on a plot with a stream. See https://chart-studio.plotly.com/settings for more details. 
    Token String `json:"token,omitempty"`
    
}
// TableCellsAlign Sets the horizontal alignment of the `text` within the box. Has an effect only if `text` spans two or more lines (i.e. `text` contains one or more <br> HTML tags) or if an explicit width is set to override the text width.
type TableCellsAlign string 

const (
    TableCellsAlignLeft TableCellsAlign = "left"
    TableCellsAlignCenter TableCellsAlign = "center"
    TableCellsAlignRight TableCellsAlign = "right"
    
)
// TableHeaderAlign Sets the horizontal alignment of the `text` within the box. Has an effect only if `text` spans two or more lines (i.e. `text` contains one or more <br> HTML tags) or if an explicit width is set to override the text width.
type TableHeaderAlign string 

const (
    TableHeaderAlignLeft TableHeaderAlign = "left"
    TableHeaderAlignCenter TableHeaderAlign = "center"
    TableHeaderAlignRight TableHeaderAlign = "right"
    
)
// TableHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type TableHoverlabelAlign string 

const (
    TableHoverlabelAlignLeft TableHoverlabelAlign = "left"
    TableHoverlabelAlignRight TableHoverlabelAlign = "right"
    TableHoverlabelAlignAuto TableHoverlabelAlign = "auto"
    
)
// TableVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type TableVisible interface{} 

var (
    TableVisibleTrue TableVisible = true
    TableVisibleFalse TableVisible = false
    TableVisibleLegendonly TableVisible = "legendonly"
    
)
// TableHoverinfo Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type TableHoverinfo string 

const (
    // Flags
    TableHoverinfoX TableHoverinfo = "x"
    TableHoverinfoY TableHoverinfo = "y"
    TableHoverinfoZ TableHoverinfo = "z"
    TableHoverinfoText TableHoverinfo = "text"
    TableHoverinfoName TableHoverinfo = "name"
    
    // Extra
    TableHoverinfoAll TableHoverinfo = "all"
    TableHoverinfoNone TableHoverinfo = "none"
    TableHoverinfoSkip TableHoverinfo = "skip"
    
)
