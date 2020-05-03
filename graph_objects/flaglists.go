package graph_objects


type AreaHoverinfo string

const (
    AreaHoverinfoX AreaHoverinfo = "x"
    AreaHoverinfoY AreaHoverinfo = "y"
    AreaHoverinfoZ AreaHoverinfo = "z"
    AreaHoverinfoText AreaHoverinfo = "text"
    AreaHoverinfoName AreaHoverinfo = "name"
)

type BarHoverinfo string

const (
    BarHoverinfoX BarHoverinfo = "x"
    BarHoverinfoY BarHoverinfo = "y"
    BarHoverinfoZ BarHoverinfo = "z"
    BarHoverinfoText BarHoverinfo = "text"
    BarHoverinfoName BarHoverinfo = "name"
)

type BarpolarHoverinfo string

const (
    BarpolarHoverinfoR BarpolarHoverinfo = "r"
    BarpolarHoverinfoTheta BarpolarHoverinfo = "theta"
    BarpolarHoverinfoText BarpolarHoverinfo = "text"
    BarpolarHoverinfoName BarpolarHoverinfo = "name"
)

type BoxHoverinfo string

const (
    BoxHoverinfoX BoxHoverinfo = "x"
    BoxHoverinfoY BoxHoverinfo = "y"
    BoxHoverinfoZ BoxHoverinfo = "z"
    BoxHoverinfoText BoxHoverinfo = "text"
    BoxHoverinfoName BoxHoverinfo = "name"
)

type BoxHoveron string

const (
    BoxHoveronBoxes BoxHoveron = "boxes"
    BoxHoveronPoints BoxHoveron = "points"
)

type CandlestickHoverinfo string

const (
    CandlestickHoverinfoX CandlestickHoverinfo = "x"
    CandlestickHoverinfoY CandlestickHoverinfo = "y"
    CandlestickHoverinfoZ CandlestickHoverinfo = "z"
    CandlestickHoverinfoText CandlestickHoverinfo = "text"
    CandlestickHoverinfoName CandlestickHoverinfo = "name"
)

type ChoroplethHoverinfo string

const (
    ChoroplethHoverinfoLocation ChoroplethHoverinfo = "location"
    ChoroplethHoverinfoZ ChoroplethHoverinfo = "z"
    ChoroplethHoverinfoText ChoroplethHoverinfo = "text"
    ChoroplethHoverinfoName ChoroplethHoverinfo = "name"
)

type ChoroplethmapboxHoverinfo string

const (
    ChoroplethmapboxHoverinfoLocation ChoroplethmapboxHoverinfo = "location"
    ChoroplethmapboxHoverinfoZ ChoroplethmapboxHoverinfo = "z"
    ChoroplethmapboxHoverinfoText ChoroplethmapboxHoverinfo = "text"
    ChoroplethmapboxHoverinfoName ChoroplethmapboxHoverinfo = "name"
)

type ConeHoverinfo string

const (
    ConeHoverinfoX ConeHoverinfo = "x"
    ConeHoverinfoY ConeHoverinfo = "y"
    ConeHoverinfoZ ConeHoverinfo = "z"
    ConeHoverinfoU ConeHoverinfo = "u"
    ConeHoverinfoV ConeHoverinfo = "v"
    ConeHoverinfoW ConeHoverinfo = "w"
    ConeHoverinfoNorm ConeHoverinfo = "norm"
    ConeHoverinfoText ConeHoverinfo = "text"
    ConeHoverinfoName ConeHoverinfo = "name"
)

type ContourHoverinfo string

const (
    ContourHoverinfoX ContourHoverinfo = "x"
    ContourHoverinfoY ContourHoverinfo = "y"
    ContourHoverinfoZ ContourHoverinfo = "z"
    ContourHoverinfoText ContourHoverinfo = "text"
    ContourHoverinfoName ContourHoverinfo = "name"
)

type DensitymapboxHoverinfo string

const (
    DensitymapboxHoverinfoLon DensitymapboxHoverinfo = "lon"
    DensitymapboxHoverinfoLat DensitymapboxHoverinfo = "lat"
    DensitymapboxHoverinfoZ DensitymapboxHoverinfo = "z"
    DensitymapboxHoverinfoText DensitymapboxHoverinfo = "text"
    DensitymapboxHoverinfoName DensitymapboxHoverinfo = "name"
)

type FunnelHoverinfo string

const (
    FunnelHoverinfoName FunnelHoverinfo = "name"
    FunnelHoverinfoX FunnelHoverinfo = "x"
    FunnelHoverinfoY FunnelHoverinfo = "y"
    FunnelHoverinfoText FunnelHoverinfo = "text"
    FunnelHoverinfoPercentInitial FunnelHoverinfo = "percent initial"
    FunnelHoverinfoPercentPrevious FunnelHoverinfo = "percent previous"
    FunnelHoverinfoPercentTotal FunnelHoverinfo = "percent total"
)

type FunnelTextinfo string

const (
    FunnelTextinfoLabel FunnelTextinfo = "label"
    FunnelTextinfoText FunnelTextinfo = "text"
    FunnelTextinfoPercentInitial FunnelTextinfo = "percent initial"
    FunnelTextinfoPercentPrevious FunnelTextinfo = "percent previous"
    FunnelTextinfoPercentTotal FunnelTextinfo = "percent total"
    FunnelTextinfoValue FunnelTextinfo = "value"
)

type FunnelareaHoverinfo string

const (
    FunnelareaHoverinfoLabel FunnelareaHoverinfo = "label"
    FunnelareaHoverinfoText FunnelareaHoverinfo = "text"
    FunnelareaHoverinfoValue FunnelareaHoverinfo = "value"
    FunnelareaHoverinfoPercent FunnelareaHoverinfo = "percent"
    FunnelareaHoverinfoName FunnelareaHoverinfo = "name"
)

type FunnelareaTextinfo string

const (
    FunnelareaTextinfoLabel FunnelareaTextinfo = "label"
    FunnelareaTextinfoText FunnelareaTextinfo = "text"
    FunnelareaTextinfoValue FunnelareaTextinfo = "value"
    FunnelareaTextinfoPercent FunnelareaTextinfo = "percent"
)

type HeatmapHoverinfo string

const (
    HeatmapHoverinfoX HeatmapHoverinfo = "x"
    HeatmapHoverinfoY HeatmapHoverinfo = "y"
    HeatmapHoverinfoZ HeatmapHoverinfo = "z"
    HeatmapHoverinfoText HeatmapHoverinfo = "text"
    HeatmapHoverinfoName HeatmapHoverinfo = "name"
)

type HeatmapglHoverinfo string

const (
    HeatmapglHoverinfoX HeatmapglHoverinfo = "x"
    HeatmapglHoverinfoY HeatmapglHoverinfo = "y"
    HeatmapglHoverinfoZ HeatmapglHoverinfo = "z"
    HeatmapglHoverinfoText HeatmapglHoverinfo = "text"
    HeatmapglHoverinfoName HeatmapglHoverinfo = "name"
)

type Histogram2dHoverinfo string

const (
    Histogram2dHoverinfoX Histogram2dHoverinfo = "x"
    Histogram2dHoverinfoY Histogram2dHoverinfo = "y"
    Histogram2dHoverinfoZ Histogram2dHoverinfo = "z"
    Histogram2dHoverinfoText Histogram2dHoverinfo = "text"
    Histogram2dHoverinfoName Histogram2dHoverinfo = "name"
)

type Histogram2dcontourHoverinfo string

const (
    Histogram2dcontourHoverinfoX Histogram2dcontourHoverinfo = "x"
    Histogram2dcontourHoverinfoY Histogram2dcontourHoverinfo = "y"
    Histogram2dcontourHoverinfoZ Histogram2dcontourHoverinfo = "z"
    Histogram2dcontourHoverinfoText Histogram2dcontourHoverinfo = "text"
    Histogram2dcontourHoverinfoName Histogram2dcontourHoverinfo = "name"
)

type HistogramHoverinfo string

const (
    HistogramHoverinfoX HistogramHoverinfo = "x"
    HistogramHoverinfoY HistogramHoverinfo = "y"
    HistogramHoverinfoZ HistogramHoverinfo = "z"
    HistogramHoverinfoText HistogramHoverinfo = "text"
    HistogramHoverinfoName HistogramHoverinfo = "name"
)

type ImageHoverinfo string

const (
    ImageHoverinfoX ImageHoverinfo = "x"
    ImageHoverinfoY ImageHoverinfo = "y"
    ImageHoverinfoZ ImageHoverinfo = "z"
    ImageHoverinfoColor ImageHoverinfo = "color"
    ImageHoverinfoName ImageHoverinfo = "name"
    ImageHoverinfoText ImageHoverinfo = "text"
)

type IndicatorMode string

const (
    IndicatorModeNumber IndicatorMode = "number"
    IndicatorModeDelta IndicatorMode = "delta"
    IndicatorModeGauge IndicatorMode = "gauge"
)

type IsosurfaceHoverinfo string

const (
    IsosurfaceHoverinfoX IsosurfaceHoverinfo = "x"
    IsosurfaceHoverinfoY IsosurfaceHoverinfo = "y"
    IsosurfaceHoverinfoZ IsosurfaceHoverinfo = "z"
    IsosurfaceHoverinfoText IsosurfaceHoverinfo = "text"
    IsosurfaceHoverinfoName IsosurfaceHoverinfo = "name"
)

type IsosurfaceSurfacePattern string

const (
    IsosurfaceSurfacePatternA IsosurfaceSurfacePattern = "A"
    IsosurfaceSurfacePatternB IsosurfaceSurfacePattern = "B"
    IsosurfaceSurfacePatternC IsosurfaceSurfacePattern = "C"
    IsosurfaceSurfacePatternD IsosurfaceSurfacePattern = "D"
    IsosurfaceSurfacePatternE IsosurfaceSurfacePattern = "E"
)

type LayoutClickmode string

const (
    LayoutClickmodeEvent LayoutClickmode = "event"
    LayoutClickmodeSelect LayoutClickmode = "select"
)

type LayoutLegendTraceorder string

const (
    LayoutLegendTraceorderReversed LayoutLegendTraceorder = "reversed"
    LayoutLegendTraceorderGrouped LayoutLegendTraceorder = "grouped"
)

type LayoutXaxisSpikemode string

const (
    LayoutXaxisSpikemodeToaxis LayoutXaxisSpikemode = "toaxis"
    LayoutXaxisSpikemodeAcross LayoutXaxisSpikemode = "across"
    LayoutXaxisSpikemodeMarker LayoutXaxisSpikemode = "marker"
)

type LayoutYaxisSpikemode string

const (
    LayoutYaxisSpikemodeToaxis LayoutYaxisSpikemode = "toaxis"
    LayoutYaxisSpikemodeAcross LayoutYaxisSpikemode = "across"
    LayoutYaxisSpikemodeMarker LayoutYaxisSpikemode = "marker"
)

type Mesh3dHoverinfo string

const (
    Mesh3dHoverinfoX Mesh3dHoverinfo = "x"
    Mesh3dHoverinfoY Mesh3dHoverinfo = "y"
    Mesh3dHoverinfoZ Mesh3dHoverinfo = "z"
    Mesh3dHoverinfoText Mesh3dHoverinfo = "text"
    Mesh3dHoverinfoName Mesh3dHoverinfo = "name"
)

type OhlcHoverinfo string

const (
    OhlcHoverinfoX OhlcHoverinfo = "x"
    OhlcHoverinfoY OhlcHoverinfo = "y"
    OhlcHoverinfoZ OhlcHoverinfo = "z"
    OhlcHoverinfoText OhlcHoverinfo = "text"
    OhlcHoverinfoName OhlcHoverinfo = "name"
)

type ParcatsHoverinfo string

const (
    ParcatsHoverinfoCount ParcatsHoverinfo = "count"
    ParcatsHoverinfoProbability ParcatsHoverinfo = "probability"
)

type PieHoverinfo string

const (
    PieHoverinfoLabel PieHoverinfo = "label"
    PieHoverinfoText PieHoverinfo = "text"
    PieHoverinfoValue PieHoverinfo = "value"
    PieHoverinfoPercent PieHoverinfo = "percent"
    PieHoverinfoName PieHoverinfo = "name"
)

type PieTextinfo string

const (
    PieTextinfoLabel PieTextinfo = "label"
    PieTextinfoText PieTextinfo = "text"
    PieTextinfoValue PieTextinfo = "value"
    PieTextinfoPercent PieTextinfo = "percent"
)

type PointcloudHoverinfo string

const (
    PointcloudHoverinfoX PointcloudHoverinfo = "x"
    PointcloudHoverinfoY PointcloudHoverinfo = "y"
    PointcloudHoverinfoZ PointcloudHoverinfo = "z"
    PointcloudHoverinfoText PointcloudHoverinfo = "text"
    PointcloudHoverinfoName PointcloudHoverinfo = "name"
)

type SankeyHoverinfo string

const (
)

type Scatter3dHoverinfo string

const (
    Scatter3dHoverinfoX Scatter3dHoverinfo = "x"
    Scatter3dHoverinfoY Scatter3dHoverinfo = "y"
    Scatter3dHoverinfoZ Scatter3dHoverinfo = "z"
    Scatter3dHoverinfoText Scatter3dHoverinfo = "text"
    Scatter3dHoverinfoName Scatter3dHoverinfo = "name"
)

type Scatter3dMode string

const (
    Scatter3dModeLines Scatter3dMode = "lines"
    Scatter3dModeMarkers Scatter3dMode = "markers"
    Scatter3dModeText Scatter3dMode = "text"
)

type ScatterHoverinfo string

const (
    ScatterHoverinfoX ScatterHoverinfo = "x"
    ScatterHoverinfoY ScatterHoverinfo = "y"
    ScatterHoverinfoZ ScatterHoverinfo = "z"
    ScatterHoverinfoText ScatterHoverinfo = "text"
    ScatterHoverinfoName ScatterHoverinfo = "name"
)

type ScatterHoveron string

const (
    ScatterHoveronPoints ScatterHoveron = "points"
    ScatterHoveronFills ScatterHoveron = "fills"
)

type ScatterMode string

const (
    ScatterModeLines ScatterMode = "lines"
    ScatterModeMarkers ScatterMode = "markers"
    ScatterModeText ScatterMode = "text"
)

type ScattercarpetHoverinfo string

const (
    ScattercarpetHoverinfoA ScattercarpetHoverinfo = "a"
    ScattercarpetHoverinfoB ScattercarpetHoverinfo = "b"
    ScattercarpetHoverinfoText ScattercarpetHoverinfo = "text"
    ScattercarpetHoverinfoName ScattercarpetHoverinfo = "name"
)

type ScattercarpetHoveron string

const (
    ScattercarpetHoveronPoints ScattercarpetHoveron = "points"
    ScattercarpetHoveronFills ScattercarpetHoveron = "fills"
)

type ScattercarpetMode string

const (
    ScattercarpetModeLines ScattercarpetMode = "lines"
    ScattercarpetModeMarkers ScattercarpetMode = "markers"
    ScattercarpetModeText ScattercarpetMode = "text"
)

type ScattergeoHoverinfo string

const (
    ScattergeoHoverinfoLon ScattergeoHoverinfo = "lon"
    ScattergeoHoverinfoLat ScattergeoHoverinfo = "lat"
    ScattergeoHoverinfoLocation ScattergeoHoverinfo = "location"
    ScattergeoHoverinfoText ScattergeoHoverinfo = "text"
    ScattergeoHoverinfoName ScattergeoHoverinfo = "name"
)

type ScattergeoMode string

const (
    ScattergeoModeLines ScattergeoMode = "lines"
    ScattergeoModeMarkers ScattergeoMode = "markers"
    ScattergeoModeText ScattergeoMode = "text"
)

type ScatterglHoverinfo string

const (
    ScatterglHoverinfoX ScatterglHoverinfo = "x"
    ScatterglHoverinfoY ScatterglHoverinfo = "y"
    ScatterglHoverinfoZ ScatterglHoverinfo = "z"
    ScatterglHoverinfoText ScatterglHoverinfo = "text"
    ScatterglHoverinfoName ScatterglHoverinfo = "name"
)

type ScatterglMode string

const (
    ScatterglModeLines ScatterglMode = "lines"
    ScatterglModeMarkers ScatterglMode = "markers"
    ScatterglModeText ScatterglMode = "text"
)

type ScattermapboxHoverinfo string

const (
    ScattermapboxHoverinfoLon ScattermapboxHoverinfo = "lon"
    ScattermapboxHoverinfoLat ScattermapboxHoverinfo = "lat"
    ScattermapboxHoverinfoText ScattermapboxHoverinfo = "text"
    ScattermapboxHoverinfoName ScattermapboxHoverinfo = "name"
)

type ScattermapboxMode string

const (
    ScattermapboxModeLines ScattermapboxMode = "lines"
    ScattermapboxModeMarkers ScattermapboxMode = "markers"
    ScattermapboxModeText ScattermapboxMode = "text"
)

type ScatterpolarHoverinfo string

const (
    ScatterpolarHoverinfoR ScatterpolarHoverinfo = "r"
    ScatterpolarHoverinfoTheta ScatterpolarHoverinfo = "theta"
    ScatterpolarHoverinfoText ScatterpolarHoverinfo = "text"
    ScatterpolarHoverinfoName ScatterpolarHoverinfo = "name"
)

type ScatterpolarHoveron string

const (
    ScatterpolarHoveronPoints ScatterpolarHoveron = "points"
    ScatterpolarHoveronFills ScatterpolarHoveron = "fills"
)

type ScatterpolarMode string

const (
    ScatterpolarModeLines ScatterpolarMode = "lines"
    ScatterpolarModeMarkers ScatterpolarMode = "markers"
    ScatterpolarModeText ScatterpolarMode = "text"
)

type ScatterpolarglHoverinfo string

const (
    ScatterpolarglHoverinfoR ScatterpolarglHoverinfo = "r"
    ScatterpolarglHoverinfoTheta ScatterpolarglHoverinfo = "theta"
    ScatterpolarglHoverinfoText ScatterpolarglHoverinfo = "text"
    ScatterpolarglHoverinfoName ScatterpolarglHoverinfo = "name"
)

type ScatterpolarglMode string

const (
    ScatterpolarglModeLines ScatterpolarglMode = "lines"
    ScatterpolarglModeMarkers ScatterpolarglMode = "markers"
    ScatterpolarglModeText ScatterpolarglMode = "text"
)

type ScatterternaryHoverinfo string

const (
    ScatterternaryHoverinfoA ScatterternaryHoverinfo = "a"
    ScatterternaryHoverinfoB ScatterternaryHoverinfo = "b"
    ScatterternaryHoverinfoC ScatterternaryHoverinfo = "c"
    ScatterternaryHoverinfoText ScatterternaryHoverinfo = "text"
    ScatterternaryHoverinfoName ScatterternaryHoverinfo = "name"
)

type ScatterternaryHoveron string

const (
    ScatterternaryHoveronPoints ScatterternaryHoveron = "points"
    ScatterternaryHoveronFills ScatterternaryHoveron = "fills"
)

type ScatterternaryMode string

const (
    ScatterternaryModeLines ScatterternaryMode = "lines"
    ScatterternaryModeMarkers ScatterternaryMode = "markers"
    ScatterternaryModeText ScatterternaryMode = "text"
)

type SplomHoverinfo string

const (
    SplomHoverinfoX SplomHoverinfo = "x"
    SplomHoverinfoY SplomHoverinfo = "y"
    SplomHoverinfoZ SplomHoverinfo = "z"
    SplomHoverinfoText SplomHoverinfo = "text"
    SplomHoverinfoName SplomHoverinfo = "name"
)

type StreamtubeHoverinfo string

const (
    StreamtubeHoverinfoX StreamtubeHoverinfo = "x"
    StreamtubeHoverinfoY StreamtubeHoverinfo = "y"
    StreamtubeHoverinfoZ StreamtubeHoverinfo = "z"
    StreamtubeHoverinfoU StreamtubeHoverinfo = "u"
    StreamtubeHoverinfoV StreamtubeHoverinfo = "v"
    StreamtubeHoverinfoW StreamtubeHoverinfo = "w"
    StreamtubeHoverinfoNorm StreamtubeHoverinfo = "norm"
    StreamtubeHoverinfoDivergence StreamtubeHoverinfo = "divergence"
    StreamtubeHoverinfoText StreamtubeHoverinfo = "text"
    StreamtubeHoverinfoName StreamtubeHoverinfo = "name"
)

type SunburstCount string

const (
    SunburstCountBranches SunburstCount = "branches"
    SunburstCountLeaves SunburstCount = "leaves"
)

type SunburstHoverinfo string

const (
    SunburstHoverinfoLabel SunburstHoverinfo = "label"
    SunburstHoverinfoText SunburstHoverinfo = "text"
    SunburstHoverinfoValue SunburstHoverinfo = "value"
    SunburstHoverinfoName SunburstHoverinfo = "name"
    SunburstHoverinfoCurrentPath SunburstHoverinfo = "current path"
    SunburstHoverinfoPercentRoot SunburstHoverinfo = "percent root"
    SunburstHoverinfoPercentEntry SunburstHoverinfo = "percent entry"
    SunburstHoverinfoPercentParent SunburstHoverinfo = "percent parent"
)

type SunburstTextinfo string

const (
    SunburstTextinfoLabel SunburstTextinfo = "label"
    SunburstTextinfoText SunburstTextinfo = "text"
    SunburstTextinfoValue SunburstTextinfo = "value"
    SunburstTextinfoCurrentPath SunburstTextinfo = "current path"
    SunburstTextinfoPercentRoot SunburstTextinfo = "percent root"
    SunburstTextinfoPercentEntry SunburstTextinfo = "percent entry"
    SunburstTextinfoPercentParent SunburstTextinfo = "percent parent"
)

type SurfaceHoverinfo string

const (
    SurfaceHoverinfoX SurfaceHoverinfo = "x"
    SurfaceHoverinfoY SurfaceHoverinfo = "y"
    SurfaceHoverinfoZ SurfaceHoverinfo = "z"
    SurfaceHoverinfoText SurfaceHoverinfo = "text"
    SurfaceHoverinfoName SurfaceHoverinfo = "name"
)

type TableHoverinfo string

const (
    TableHoverinfoX TableHoverinfo = "x"
    TableHoverinfoY TableHoverinfo = "y"
    TableHoverinfoZ TableHoverinfo = "z"
    TableHoverinfoText TableHoverinfo = "text"
    TableHoverinfoName TableHoverinfo = "name"
)

type TreemapCount string

const (
    TreemapCountBranches TreemapCount = "branches"
    TreemapCountLeaves TreemapCount = "leaves"
)

type TreemapHoverinfo string

const (
    TreemapHoverinfoLabel TreemapHoverinfo = "label"
    TreemapHoverinfoText TreemapHoverinfo = "text"
    TreemapHoverinfoValue TreemapHoverinfo = "value"
    TreemapHoverinfoName TreemapHoverinfo = "name"
    TreemapHoverinfoCurrentPath TreemapHoverinfo = "current path"
    TreemapHoverinfoPercentRoot TreemapHoverinfo = "percent root"
    TreemapHoverinfoPercentEntry TreemapHoverinfo = "percent entry"
    TreemapHoverinfoPercentParent TreemapHoverinfo = "percent parent"
)

type TreemapTextinfo string

const (
    TreemapTextinfoLabel TreemapTextinfo = "label"
    TreemapTextinfoText TreemapTextinfo = "text"
    TreemapTextinfoValue TreemapTextinfo = "value"
    TreemapTextinfoCurrentPath TreemapTextinfo = "current path"
    TreemapTextinfoPercentRoot TreemapTextinfo = "percent root"
    TreemapTextinfoPercentEntry TreemapTextinfo = "percent entry"
    TreemapTextinfoPercentParent TreemapTextinfo = "percent parent"
)

type TreemapTilingFlip string

const (
    TreemapTilingFlipX TreemapTilingFlip = "x"
    TreemapTilingFlipY TreemapTilingFlip = "y"
)

type ViolinHoverinfo string

const (
    ViolinHoverinfoX ViolinHoverinfo = "x"
    ViolinHoverinfoY ViolinHoverinfo = "y"
    ViolinHoverinfoZ ViolinHoverinfo = "z"
    ViolinHoverinfoText ViolinHoverinfo = "text"
    ViolinHoverinfoName ViolinHoverinfo = "name"
)

type ViolinHoveron string

const (
    ViolinHoveronViolins ViolinHoveron = "violins"
    ViolinHoveronPoints ViolinHoveron = "points"
    ViolinHoveronKde ViolinHoveron = "kde"
)

type VolumeHoverinfo string

const (
    VolumeHoverinfoX VolumeHoverinfo = "x"
    VolumeHoverinfoY VolumeHoverinfo = "y"
    VolumeHoverinfoZ VolumeHoverinfo = "z"
    VolumeHoverinfoText VolumeHoverinfo = "text"
    VolumeHoverinfoName VolumeHoverinfo = "name"
)

type VolumeSurfacePattern string

const (
    VolumeSurfacePatternA VolumeSurfacePattern = "A"
    VolumeSurfacePatternB VolumeSurfacePattern = "B"
    VolumeSurfacePatternC VolumeSurfacePattern = "C"
    VolumeSurfacePatternD VolumeSurfacePattern = "D"
    VolumeSurfacePatternE VolumeSurfacePattern = "E"
)

type WaterfallHoverinfo string

const (
    WaterfallHoverinfoName WaterfallHoverinfo = "name"
    WaterfallHoverinfoX WaterfallHoverinfo = "x"
    WaterfallHoverinfoY WaterfallHoverinfo = "y"
    WaterfallHoverinfoText WaterfallHoverinfo = "text"
    WaterfallHoverinfoInitial WaterfallHoverinfo = "initial"
    WaterfallHoverinfoDelta WaterfallHoverinfo = "delta"
    WaterfallHoverinfoFinal WaterfallHoverinfo = "final"
)

type WaterfallTextinfo string

const (
    WaterfallTextinfoLabel WaterfallTextinfo = "label"
    WaterfallTextinfoText WaterfallTextinfo = "text"
    WaterfallTextinfoInitial WaterfallTextinfo = "initial"
    WaterfallTextinfoDelta WaterfallTextinfo = "delta"
    WaterfallTextinfoFinal WaterfallTextinfo = "final"
)
