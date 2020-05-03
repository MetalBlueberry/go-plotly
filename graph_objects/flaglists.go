package graph_objects

type AreaHoverinfo string

const (
	// Flags
	AreaHoverinfoX    AreaHoverinfo = "x"
	AreaHoverinfoY    AreaHoverinfo = "y"
	AreaHoverinfoZ    AreaHoverinfo = "z"
	AreaHoverinfoText AreaHoverinfo = "text"
	AreaHoverinfoName AreaHoverinfo = "name"
	// Extras
	AreaHoverinfoAll  AreaHoverinfo = "all"
	AreaHoverinfoNone AreaHoverinfo = "none"
	AreaHoverinfoSkip AreaHoverinfo = "skip"
)

type BarHoverinfo string

const (
	// Flags
	BarHoverinfoX    BarHoverinfo = "x"
	BarHoverinfoY    BarHoverinfo = "y"
	BarHoverinfoZ    BarHoverinfo = "z"
	BarHoverinfoText BarHoverinfo = "text"
	BarHoverinfoName BarHoverinfo = "name"
	// Extras
	BarHoverinfoAll  BarHoverinfo = "all"
	BarHoverinfoNone BarHoverinfo = "none"
	BarHoverinfoSkip BarHoverinfo = "skip"
)

type BarpolarHoverinfo string

const (
	// Flags
	BarpolarHoverinfoR     BarpolarHoverinfo = "r"
	BarpolarHoverinfoTheta BarpolarHoverinfo = "theta"
	BarpolarHoverinfoText  BarpolarHoverinfo = "text"
	BarpolarHoverinfoName  BarpolarHoverinfo = "name"
	// Extras
	BarpolarHoverinfoAll  BarpolarHoverinfo = "all"
	BarpolarHoverinfoNone BarpolarHoverinfo = "none"
	BarpolarHoverinfoSkip BarpolarHoverinfo = "skip"
)

type BoxHoverinfo string

const (
	// Flags
	BoxHoverinfoX    BoxHoverinfo = "x"
	BoxHoverinfoY    BoxHoverinfo = "y"
	BoxHoverinfoZ    BoxHoverinfo = "z"
	BoxHoverinfoText BoxHoverinfo = "text"
	BoxHoverinfoName BoxHoverinfo = "name"
	// Extras
	BoxHoverinfoAll  BoxHoverinfo = "all"
	BoxHoverinfoNone BoxHoverinfo = "none"
	BoxHoverinfoSkip BoxHoverinfo = "skip"
)

type BoxHoveron string

const (
	// Flags
	BoxHoveronBoxes  BoxHoveron = "boxes"
	BoxHoveronPoints BoxHoveron = "points"
	// Extras
)

type CandlestickHoverinfo string

const (
	// Flags
	CandlestickHoverinfoX    CandlestickHoverinfo = "x"
	CandlestickHoverinfoY    CandlestickHoverinfo = "y"
	CandlestickHoverinfoZ    CandlestickHoverinfo = "z"
	CandlestickHoverinfoText CandlestickHoverinfo = "text"
	CandlestickHoverinfoName CandlestickHoverinfo = "name"
	// Extras
	CandlestickHoverinfoAll  CandlestickHoverinfo = "all"
	CandlestickHoverinfoNone CandlestickHoverinfo = "none"
	CandlestickHoverinfoSkip CandlestickHoverinfo = "skip"
)

type ChoroplethHoverinfo string

const (
	// Flags
	ChoroplethHoverinfoLocation ChoroplethHoverinfo = "location"
	ChoroplethHoverinfoZ        ChoroplethHoverinfo = "z"
	ChoroplethHoverinfoText     ChoroplethHoverinfo = "text"
	ChoroplethHoverinfoName     ChoroplethHoverinfo = "name"
	// Extras
	ChoroplethHoverinfoAll  ChoroplethHoverinfo = "all"
	ChoroplethHoverinfoNone ChoroplethHoverinfo = "none"
	ChoroplethHoverinfoSkip ChoroplethHoverinfo = "skip"
)

type ChoroplethmapboxHoverinfo string

const (
	// Flags
	ChoroplethmapboxHoverinfoLocation ChoroplethmapboxHoverinfo = "location"
	ChoroplethmapboxHoverinfoZ        ChoroplethmapboxHoverinfo = "z"
	ChoroplethmapboxHoverinfoText     ChoroplethmapboxHoverinfo = "text"
	ChoroplethmapboxHoverinfoName     ChoroplethmapboxHoverinfo = "name"
	// Extras
	ChoroplethmapboxHoverinfoAll  ChoroplethmapboxHoverinfo = "all"
	ChoroplethmapboxHoverinfoNone ChoroplethmapboxHoverinfo = "none"
	ChoroplethmapboxHoverinfoSkip ChoroplethmapboxHoverinfo = "skip"
)

type ConeHoverinfo string

const (
	// Flags
	ConeHoverinfoX    ConeHoverinfo = "x"
	ConeHoverinfoY    ConeHoverinfo = "y"
	ConeHoverinfoZ    ConeHoverinfo = "z"
	ConeHoverinfoU    ConeHoverinfo = "u"
	ConeHoverinfoV    ConeHoverinfo = "v"
	ConeHoverinfoW    ConeHoverinfo = "w"
	ConeHoverinfoNorm ConeHoverinfo = "norm"
	ConeHoverinfoText ConeHoverinfo = "text"
	ConeHoverinfoName ConeHoverinfo = "name"
	// Extras
	ConeHoverinfoAll  ConeHoverinfo = "all"
	ConeHoverinfoNone ConeHoverinfo = "none"
	ConeHoverinfoSkip ConeHoverinfo = "skip"
)

type ContourHoverinfo string

const (
	// Flags
	ContourHoverinfoX    ContourHoverinfo = "x"
	ContourHoverinfoY    ContourHoverinfo = "y"
	ContourHoverinfoZ    ContourHoverinfo = "z"
	ContourHoverinfoText ContourHoverinfo = "text"
	ContourHoverinfoName ContourHoverinfo = "name"
	// Extras
	ContourHoverinfoAll  ContourHoverinfo = "all"
	ContourHoverinfoNone ContourHoverinfo = "none"
	ContourHoverinfoSkip ContourHoverinfo = "skip"
)

type DensitymapboxHoverinfo string

const (
	// Flags
	DensitymapboxHoverinfoLon  DensitymapboxHoverinfo = "lon"
	DensitymapboxHoverinfoLat  DensitymapboxHoverinfo = "lat"
	DensitymapboxHoverinfoZ    DensitymapboxHoverinfo = "z"
	DensitymapboxHoverinfoText DensitymapboxHoverinfo = "text"
	DensitymapboxHoverinfoName DensitymapboxHoverinfo = "name"
	// Extras
	DensitymapboxHoverinfoAll  DensitymapboxHoverinfo = "all"
	DensitymapboxHoverinfoNone DensitymapboxHoverinfo = "none"
	DensitymapboxHoverinfoSkip DensitymapboxHoverinfo = "skip"
)

type FunnelHoverinfo string

const (
	// Flags
	FunnelHoverinfoName            FunnelHoverinfo = "name"
	FunnelHoverinfoX               FunnelHoverinfo = "x"
	FunnelHoverinfoY               FunnelHoverinfo = "y"
	FunnelHoverinfoText            FunnelHoverinfo = "text"
	FunnelHoverinfoPercentInitial  FunnelHoverinfo = "percent initial"
	FunnelHoverinfoPercentPrevious FunnelHoverinfo = "percent previous"
	FunnelHoverinfoPercentTotal    FunnelHoverinfo = "percent total"
	// Extras
	FunnelHoverinfoAll  FunnelHoverinfo = "all"
	FunnelHoverinfoNone FunnelHoverinfo = "none"
	FunnelHoverinfoSkip FunnelHoverinfo = "skip"
)

type FunnelTextinfo string

const (
	// Flags
	FunnelTextinfoLabel           FunnelTextinfo = "label"
	FunnelTextinfoText            FunnelTextinfo = "text"
	FunnelTextinfoPercentInitial  FunnelTextinfo = "percent initial"
	FunnelTextinfoPercentPrevious FunnelTextinfo = "percent previous"
	FunnelTextinfoPercentTotal    FunnelTextinfo = "percent total"
	FunnelTextinfoValue           FunnelTextinfo = "value"
	// Extras
	FunnelTextinfoNone FunnelTextinfo = "none"
)

type FunnelareaHoverinfo string

const (
	// Flags
	FunnelareaHoverinfoLabel   FunnelareaHoverinfo = "label"
	FunnelareaHoverinfoText    FunnelareaHoverinfo = "text"
	FunnelareaHoverinfoValue   FunnelareaHoverinfo = "value"
	FunnelareaHoverinfoPercent FunnelareaHoverinfo = "percent"
	FunnelareaHoverinfoName    FunnelareaHoverinfo = "name"
	// Extras
	FunnelareaHoverinfoAll  FunnelareaHoverinfo = "all"
	FunnelareaHoverinfoNone FunnelareaHoverinfo = "none"
	FunnelareaHoverinfoSkip FunnelareaHoverinfo = "skip"
)

type FunnelareaTextinfo string

const (
	// Flags
	FunnelareaTextinfoLabel   FunnelareaTextinfo = "label"
	FunnelareaTextinfoText    FunnelareaTextinfo = "text"
	FunnelareaTextinfoValue   FunnelareaTextinfo = "value"
	FunnelareaTextinfoPercent FunnelareaTextinfo = "percent"
	// Extras
	FunnelareaTextinfoNone FunnelareaTextinfo = "none"
)

type HeatmapHoverinfo string

const (
	// Flags
	HeatmapHoverinfoX    HeatmapHoverinfo = "x"
	HeatmapHoverinfoY    HeatmapHoverinfo = "y"
	HeatmapHoverinfoZ    HeatmapHoverinfo = "z"
	HeatmapHoverinfoText HeatmapHoverinfo = "text"
	HeatmapHoverinfoName HeatmapHoverinfo = "name"
	// Extras
	HeatmapHoverinfoAll  HeatmapHoverinfo = "all"
	HeatmapHoverinfoNone HeatmapHoverinfo = "none"
	HeatmapHoverinfoSkip HeatmapHoverinfo = "skip"
)

type HeatmapglHoverinfo string

const (
	// Flags
	HeatmapglHoverinfoX    HeatmapglHoverinfo = "x"
	HeatmapglHoverinfoY    HeatmapglHoverinfo = "y"
	HeatmapglHoverinfoZ    HeatmapglHoverinfo = "z"
	HeatmapglHoverinfoText HeatmapglHoverinfo = "text"
	HeatmapglHoverinfoName HeatmapglHoverinfo = "name"
	// Extras
	HeatmapglHoverinfoAll  HeatmapglHoverinfo = "all"
	HeatmapglHoverinfoNone HeatmapglHoverinfo = "none"
	HeatmapglHoverinfoSkip HeatmapglHoverinfo = "skip"
)

type Histogram2dHoverinfo string

const (
	// Flags
	Histogram2dHoverinfoX    Histogram2dHoverinfo = "x"
	Histogram2dHoverinfoY    Histogram2dHoverinfo = "y"
	Histogram2dHoverinfoZ    Histogram2dHoverinfo = "z"
	Histogram2dHoverinfoText Histogram2dHoverinfo = "text"
	Histogram2dHoverinfoName Histogram2dHoverinfo = "name"
	// Extras
	Histogram2dHoverinfoAll  Histogram2dHoverinfo = "all"
	Histogram2dHoverinfoNone Histogram2dHoverinfo = "none"
	Histogram2dHoverinfoSkip Histogram2dHoverinfo = "skip"
)

type Histogram2dcontourHoverinfo string

const (
	// Flags
	Histogram2dcontourHoverinfoX    Histogram2dcontourHoverinfo = "x"
	Histogram2dcontourHoverinfoY    Histogram2dcontourHoverinfo = "y"
	Histogram2dcontourHoverinfoZ    Histogram2dcontourHoverinfo = "z"
	Histogram2dcontourHoverinfoText Histogram2dcontourHoverinfo = "text"
	Histogram2dcontourHoverinfoName Histogram2dcontourHoverinfo = "name"
	// Extras
	Histogram2dcontourHoverinfoAll  Histogram2dcontourHoverinfo = "all"
	Histogram2dcontourHoverinfoNone Histogram2dcontourHoverinfo = "none"
	Histogram2dcontourHoverinfoSkip Histogram2dcontourHoverinfo = "skip"
)

type HistogramHoverinfo string

const (
	// Flags
	HistogramHoverinfoX    HistogramHoverinfo = "x"
	HistogramHoverinfoY    HistogramHoverinfo = "y"
	HistogramHoverinfoZ    HistogramHoverinfo = "z"
	HistogramHoverinfoText HistogramHoverinfo = "text"
	HistogramHoverinfoName HistogramHoverinfo = "name"
	// Extras
	HistogramHoverinfoAll  HistogramHoverinfo = "all"
	HistogramHoverinfoNone HistogramHoverinfo = "none"
	HistogramHoverinfoSkip HistogramHoverinfo = "skip"
)

type ImageHoverinfo string

const (
	// Flags
	ImageHoverinfoX     ImageHoverinfo = "x"
	ImageHoverinfoY     ImageHoverinfo = "y"
	ImageHoverinfoZ     ImageHoverinfo = "z"
	ImageHoverinfoColor ImageHoverinfo = "color"
	ImageHoverinfoName  ImageHoverinfo = "name"
	ImageHoverinfoText  ImageHoverinfo = "text"
	// Extras
	ImageHoverinfoAll  ImageHoverinfo = "all"
	ImageHoverinfoNone ImageHoverinfo = "none"
	ImageHoverinfoSkip ImageHoverinfo = "skip"
)

type IndicatorMode string

const (
	// Flags
	IndicatorModeNumber IndicatorMode = "number"
	IndicatorModeDelta  IndicatorMode = "delta"
	IndicatorModeGauge  IndicatorMode = "gauge"
	// Extras
)

type IsosurfaceHoverinfo string

const (
	// Flags
	IsosurfaceHoverinfoX    IsosurfaceHoverinfo = "x"
	IsosurfaceHoverinfoY    IsosurfaceHoverinfo = "y"
	IsosurfaceHoverinfoZ    IsosurfaceHoverinfo = "z"
	IsosurfaceHoverinfoText IsosurfaceHoverinfo = "text"
	IsosurfaceHoverinfoName IsosurfaceHoverinfo = "name"
	// Extras
	IsosurfaceHoverinfoAll  IsosurfaceHoverinfo = "all"
	IsosurfaceHoverinfoNone IsosurfaceHoverinfo = "none"
	IsosurfaceHoverinfoSkip IsosurfaceHoverinfo = "skip"
)

type IsosurfaceSurfacePattern string

const (
	// Flags
	IsosurfaceSurfacePatternA IsosurfaceSurfacePattern = "A"
	IsosurfaceSurfacePatternB IsosurfaceSurfacePattern = "B"
	IsosurfaceSurfacePatternC IsosurfaceSurfacePattern = "C"
	IsosurfaceSurfacePatternD IsosurfaceSurfacePattern = "D"
	IsosurfaceSurfacePatternE IsosurfaceSurfacePattern = "E"
	// Extras
	IsosurfaceSurfacePatternAll  IsosurfaceSurfacePattern = "all"
	IsosurfaceSurfacePatternOdd  IsosurfaceSurfacePattern = "odd"
	IsosurfaceSurfacePatternEven IsosurfaceSurfacePattern = "even"
)

type LayoutClickmode string

const (
	// Flags
	LayoutClickmodeEvent  LayoutClickmode = "event"
	LayoutClickmodeSelect LayoutClickmode = "select"
	// Extras
	LayoutClickmodeNone LayoutClickmode = "none"
)

type LayoutLegendTraceorder string

const (
	// Flags
	LayoutLegendTraceorderReversed LayoutLegendTraceorder = "reversed"
	LayoutLegendTraceorderGrouped  LayoutLegendTraceorder = "grouped"
	// Extras
	LayoutLegendTraceorderNormal LayoutLegendTraceorder = "normal"
)

type LayoutXaxisSpikemode string

const (
	// Flags
	LayoutXaxisSpikemodeToaxis LayoutXaxisSpikemode = "toaxis"
	LayoutXaxisSpikemodeAcross LayoutXaxisSpikemode = "across"
	LayoutXaxisSpikemodeMarker LayoutXaxisSpikemode = "marker"
	// Extras
)

type LayoutYaxisSpikemode string

const (
	// Flags
	LayoutYaxisSpikemodeToaxis LayoutYaxisSpikemode = "toaxis"
	LayoutYaxisSpikemodeAcross LayoutYaxisSpikemode = "across"
	LayoutYaxisSpikemodeMarker LayoutYaxisSpikemode = "marker"
	// Extras
)

type Mesh3dHoverinfo string

const (
	// Flags
	Mesh3dHoverinfoX    Mesh3dHoverinfo = "x"
	Mesh3dHoverinfoY    Mesh3dHoverinfo = "y"
	Mesh3dHoverinfoZ    Mesh3dHoverinfo = "z"
	Mesh3dHoverinfoText Mesh3dHoverinfo = "text"
	Mesh3dHoverinfoName Mesh3dHoverinfo = "name"
	// Extras
	Mesh3dHoverinfoAll  Mesh3dHoverinfo = "all"
	Mesh3dHoverinfoNone Mesh3dHoverinfo = "none"
	Mesh3dHoverinfoSkip Mesh3dHoverinfo = "skip"
)

type OhlcHoverinfo string

const (
	// Flags
	OhlcHoverinfoX    OhlcHoverinfo = "x"
	OhlcHoverinfoY    OhlcHoverinfo = "y"
	OhlcHoverinfoZ    OhlcHoverinfo = "z"
	OhlcHoverinfoText OhlcHoverinfo = "text"
	OhlcHoverinfoName OhlcHoverinfo = "name"
	// Extras
	OhlcHoverinfoAll  OhlcHoverinfo = "all"
	OhlcHoverinfoNone OhlcHoverinfo = "none"
	OhlcHoverinfoSkip OhlcHoverinfo = "skip"
)

type ParcatsHoverinfo string

const (
	// Flags
	ParcatsHoverinfoCount       ParcatsHoverinfo = "count"
	ParcatsHoverinfoProbability ParcatsHoverinfo = "probability"
	// Extras
	ParcatsHoverinfoAll  ParcatsHoverinfo = "all"
	ParcatsHoverinfoNone ParcatsHoverinfo = "none"
	ParcatsHoverinfoSkip ParcatsHoverinfo = "skip"
)

type PieHoverinfo string

const (
	// Flags
	PieHoverinfoLabel   PieHoverinfo = "label"
	PieHoverinfoText    PieHoverinfo = "text"
	PieHoverinfoValue   PieHoverinfo = "value"
	PieHoverinfoPercent PieHoverinfo = "percent"
	PieHoverinfoName    PieHoverinfo = "name"
	// Extras
	PieHoverinfoAll  PieHoverinfo = "all"
	PieHoverinfoNone PieHoverinfo = "none"
	PieHoverinfoSkip PieHoverinfo = "skip"
)

type PieTextinfo string

const (
	// Flags
	PieTextinfoLabel   PieTextinfo = "label"
	PieTextinfoText    PieTextinfo = "text"
	PieTextinfoValue   PieTextinfo = "value"
	PieTextinfoPercent PieTextinfo = "percent"
	// Extras
	PieTextinfoNone PieTextinfo = "none"
)

type PointcloudHoverinfo string

const (
	// Flags
	PointcloudHoverinfoX    PointcloudHoverinfo = "x"
	PointcloudHoverinfoY    PointcloudHoverinfo = "y"
	PointcloudHoverinfoZ    PointcloudHoverinfo = "z"
	PointcloudHoverinfoText PointcloudHoverinfo = "text"
	PointcloudHoverinfoName PointcloudHoverinfo = "name"
	// Extras
	PointcloudHoverinfoAll  PointcloudHoverinfo = "all"
	PointcloudHoverinfoNone PointcloudHoverinfo = "none"
	PointcloudHoverinfoSkip PointcloudHoverinfo = "skip"
)

type SankeyHoverinfo string

const (
	// Flags
	// Extras
	SankeyHoverinfoAll  SankeyHoverinfo = "all"
	SankeyHoverinfoNone SankeyHoverinfo = "none"
	SankeyHoverinfoSkip SankeyHoverinfo = "skip"
)

type Scatter3dHoverinfo string

const (
	// Flags
	Scatter3dHoverinfoX    Scatter3dHoverinfo = "x"
	Scatter3dHoverinfoY    Scatter3dHoverinfo = "y"
	Scatter3dHoverinfoZ    Scatter3dHoverinfo = "z"
	Scatter3dHoverinfoText Scatter3dHoverinfo = "text"
	Scatter3dHoverinfoName Scatter3dHoverinfo = "name"
	// Extras
	Scatter3dHoverinfoAll  Scatter3dHoverinfo = "all"
	Scatter3dHoverinfoNone Scatter3dHoverinfo = "none"
	Scatter3dHoverinfoSkip Scatter3dHoverinfo = "skip"
)

type Scatter3dMode string

const (
	// Flags
	Scatter3dModeLines   Scatter3dMode = "lines"
	Scatter3dModeMarkers Scatter3dMode = "markers"
	Scatter3dModeText    Scatter3dMode = "text"
	// Extras
	Scatter3dModeNone Scatter3dMode = "none"
)

type ScatterHoverinfo string

const (
	// Flags
	ScatterHoverinfoX    ScatterHoverinfo = "x"
	ScatterHoverinfoY    ScatterHoverinfo = "y"
	ScatterHoverinfoZ    ScatterHoverinfo = "z"
	ScatterHoverinfoText ScatterHoverinfo = "text"
	ScatterHoverinfoName ScatterHoverinfo = "name"
	// Extras
	ScatterHoverinfoAll  ScatterHoverinfo = "all"
	ScatterHoverinfoNone ScatterHoverinfo = "none"
	ScatterHoverinfoSkip ScatterHoverinfo = "skip"
)

type ScatterHoveron string

const (
	// Flags
	ScatterHoveronPoints ScatterHoveron = "points"
	ScatterHoveronFills  ScatterHoveron = "fills"
	// Extras
)

type ScatterMode string

const (
	// Flags
	ScatterModeLines   ScatterMode = "lines"
	ScatterModeMarkers ScatterMode = "markers"
	ScatterModeText    ScatterMode = "text"
	// Extras
	ScatterModeNone ScatterMode = "none"
)

type ScattercarpetHoverinfo string

const (
	// Flags
	ScattercarpetHoverinfoA    ScattercarpetHoverinfo = "a"
	ScattercarpetHoverinfoB    ScattercarpetHoverinfo = "b"
	ScattercarpetHoverinfoText ScattercarpetHoverinfo = "text"
	ScattercarpetHoverinfoName ScattercarpetHoverinfo = "name"
	// Extras
	ScattercarpetHoverinfoAll  ScattercarpetHoverinfo = "all"
	ScattercarpetHoverinfoNone ScattercarpetHoverinfo = "none"
	ScattercarpetHoverinfoSkip ScattercarpetHoverinfo = "skip"
)

type ScattercarpetHoveron string

const (
	// Flags
	ScattercarpetHoveronPoints ScattercarpetHoveron = "points"
	ScattercarpetHoveronFills  ScattercarpetHoveron = "fills"
	// Extras
)

type ScattercarpetMode string

const (
	// Flags
	ScattercarpetModeLines   ScattercarpetMode = "lines"
	ScattercarpetModeMarkers ScattercarpetMode = "markers"
	ScattercarpetModeText    ScattercarpetMode = "text"
	// Extras
	ScattercarpetModeNone ScattercarpetMode = "none"
)

type ScattergeoHoverinfo string

const (
	// Flags
	ScattergeoHoverinfoLon      ScattergeoHoverinfo = "lon"
	ScattergeoHoverinfoLat      ScattergeoHoverinfo = "lat"
	ScattergeoHoverinfoLocation ScattergeoHoverinfo = "location"
	ScattergeoHoverinfoText     ScattergeoHoverinfo = "text"
	ScattergeoHoverinfoName     ScattergeoHoverinfo = "name"
	// Extras
	ScattergeoHoverinfoAll  ScattergeoHoverinfo = "all"
	ScattergeoHoverinfoNone ScattergeoHoverinfo = "none"
	ScattergeoHoverinfoSkip ScattergeoHoverinfo = "skip"
)

type ScattergeoMode string

const (
	// Flags
	ScattergeoModeLines   ScattergeoMode = "lines"
	ScattergeoModeMarkers ScattergeoMode = "markers"
	ScattergeoModeText    ScattergeoMode = "text"
	// Extras
	ScattergeoModeNone ScattergeoMode = "none"
)

type ScatterglHoverinfo string

const (
	// Flags
	ScatterglHoverinfoX    ScatterglHoverinfo = "x"
	ScatterglHoverinfoY    ScatterglHoverinfo = "y"
	ScatterglHoverinfoZ    ScatterglHoverinfo = "z"
	ScatterglHoverinfoText ScatterglHoverinfo = "text"
	ScatterglHoverinfoName ScatterglHoverinfo = "name"
	// Extras
	ScatterglHoverinfoAll  ScatterglHoverinfo = "all"
	ScatterglHoverinfoNone ScatterglHoverinfo = "none"
	ScatterglHoverinfoSkip ScatterglHoverinfo = "skip"
)

type ScatterglMode string

const (
	// Flags
	ScatterglModeLines   ScatterglMode = "lines"
	ScatterglModeMarkers ScatterglMode = "markers"
	ScatterglModeText    ScatterglMode = "text"
	// Extras
	ScatterglModeNone ScatterglMode = "none"
)

type ScattermapboxHoverinfo string

const (
	// Flags
	ScattermapboxHoverinfoLon  ScattermapboxHoverinfo = "lon"
	ScattermapboxHoverinfoLat  ScattermapboxHoverinfo = "lat"
	ScattermapboxHoverinfoText ScattermapboxHoverinfo = "text"
	ScattermapboxHoverinfoName ScattermapboxHoverinfo = "name"
	// Extras
	ScattermapboxHoverinfoAll  ScattermapboxHoverinfo = "all"
	ScattermapboxHoverinfoNone ScattermapboxHoverinfo = "none"
	ScattermapboxHoverinfoSkip ScattermapboxHoverinfo = "skip"
)

type ScattermapboxMode string

const (
	// Flags
	ScattermapboxModeLines   ScattermapboxMode = "lines"
	ScattermapboxModeMarkers ScattermapboxMode = "markers"
	ScattermapboxModeText    ScattermapboxMode = "text"
	// Extras
	ScattermapboxModeNone ScattermapboxMode = "none"
)

type ScatterpolarHoverinfo string

const (
	// Flags
	ScatterpolarHoverinfoR     ScatterpolarHoverinfo = "r"
	ScatterpolarHoverinfoTheta ScatterpolarHoverinfo = "theta"
	ScatterpolarHoverinfoText  ScatterpolarHoverinfo = "text"
	ScatterpolarHoverinfoName  ScatterpolarHoverinfo = "name"
	// Extras
	ScatterpolarHoverinfoAll  ScatterpolarHoverinfo = "all"
	ScatterpolarHoverinfoNone ScatterpolarHoverinfo = "none"
	ScatterpolarHoverinfoSkip ScatterpolarHoverinfo = "skip"
)

type ScatterpolarHoveron string

const (
	// Flags
	ScatterpolarHoveronPoints ScatterpolarHoveron = "points"
	ScatterpolarHoveronFills  ScatterpolarHoveron = "fills"
	// Extras
)

type ScatterpolarMode string

const (
	// Flags
	ScatterpolarModeLines   ScatterpolarMode = "lines"
	ScatterpolarModeMarkers ScatterpolarMode = "markers"
	ScatterpolarModeText    ScatterpolarMode = "text"
	// Extras
	ScatterpolarModeNone ScatterpolarMode = "none"
)

type ScatterpolarglHoverinfo string

const (
	// Flags
	ScatterpolarglHoverinfoR     ScatterpolarglHoverinfo = "r"
	ScatterpolarglHoverinfoTheta ScatterpolarglHoverinfo = "theta"
	ScatterpolarglHoverinfoText  ScatterpolarglHoverinfo = "text"
	ScatterpolarglHoverinfoName  ScatterpolarglHoverinfo = "name"
	// Extras
	ScatterpolarglHoverinfoAll  ScatterpolarglHoverinfo = "all"
	ScatterpolarglHoverinfoNone ScatterpolarglHoverinfo = "none"
	ScatterpolarglHoverinfoSkip ScatterpolarglHoverinfo = "skip"
)

type ScatterpolarglMode string

const (
	// Flags
	ScatterpolarglModeLines   ScatterpolarglMode = "lines"
	ScatterpolarglModeMarkers ScatterpolarglMode = "markers"
	ScatterpolarglModeText    ScatterpolarglMode = "text"
	// Extras
	ScatterpolarglModeNone ScatterpolarglMode = "none"
)

type ScatterternaryHoverinfo string

const (
	// Flags
	ScatterternaryHoverinfoA    ScatterternaryHoverinfo = "a"
	ScatterternaryHoverinfoB    ScatterternaryHoverinfo = "b"
	ScatterternaryHoverinfoC    ScatterternaryHoverinfo = "c"
	ScatterternaryHoverinfoText ScatterternaryHoverinfo = "text"
	ScatterternaryHoverinfoName ScatterternaryHoverinfo = "name"
	// Extras
	ScatterternaryHoverinfoAll  ScatterternaryHoverinfo = "all"
	ScatterternaryHoverinfoNone ScatterternaryHoverinfo = "none"
	ScatterternaryHoverinfoSkip ScatterternaryHoverinfo = "skip"
)

type ScatterternaryHoveron string

const (
	// Flags
	ScatterternaryHoveronPoints ScatterternaryHoveron = "points"
	ScatterternaryHoveronFills  ScatterternaryHoveron = "fills"
	// Extras
)

type ScatterternaryMode string

const (
	// Flags
	ScatterternaryModeLines   ScatterternaryMode = "lines"
	ScatterternaryModeMarkers ScatterternaryMode = "markers"
	ScatterternaryModeText    ScatterternaryMode = "text"
	// Extras
	ScatterternaryModeNone ScatterternaryMode = "none"
)

type SplomHoverinfo string

const (
	// Flags
	SplomHoverinfoX    SplomHoverinfo = "x"
	SplomHoverinfoY    SplomHoverinfo = "y"
	SplomHoverinfoZ    SplomHoverinfo = "z"
	SplomHoverinfoText SplomHoverinfo = "text"
	SplomHoverinfoName SplomHoverinfo = "name"
	// Extras
	SplomHoverinfoAll  SplomHoverinfo = "all"
	SplomHoverinfoNone SplomHoverinfo = "none"
	SplomHoverinfoSkip SplomHoverinfo = "skip"
)

type StreamtubeHoverinfo string

const (
	// Flags
	StreamtubeHoverinfoX          StreamtubeHoverinfo = "x"
	StreamtubeHoverinfoY          StreamtubeHoverinfo = "y"
	StreamtubeHoverinfoZ          StreamtubeHoverinfo = "z"
	StreamtubeHoverinfoU          StreamtubeHoverinfo = "u"
	StreamtubeHoverinfoV          StreamtubeHoverinfo = "v"
	StreamtubeHoverinfoW          StreamtubeHoverinfo = "w"
	StreamtubeHoverinfoNorm       StreamtubeHoverinfo = "norm"
	StreamtubeHoverinfoDivergence StreamtubeHoverinfo = "divergence"
	StreamtubeHoverinfoText       StreamtubeHoverinfo = "text"
	StreamtubeHoverinfoName       StreamtubeHoverinfo = "name"
	// Extras
	StreamtubeHoverinfoAll  StreamtubeHoverinfo = "all"
	StreamtubeHoverinfoNone StreamtubeHoverinfo = "none"
	StreamtubeHoverinfoSkip StreamtubeHoverinfo = "skip"
)

type SunburstCount string

const (
	// Flags
	SunburstCountBranches SunburstCount = "branches"
	SunburstCountLeaves   SunburstCount = "leaves"
	// Extras
)

type SunburstHoverinfo string

const (
	// Flags
	SunburstHoverinfoLabel         SunburstHoverinfo = "label"
	SunburstHoverinfoText          SunburstHoverinfo = "text"
	SunburstHoverinfoValue         SunburstHoverinfo = "value"
	SunburstHoverinfoName          SunburstHoverinfo = "name"
	SunburstHoverinfoCurrentPath   SunburstHoverinfo = "current path"
	SunburstHoverinfoPercentRoot   SunburstHoverinfo = "percent root"
	SunburstHoverinfoPercentEntry  SunburstHoverinfo = "percent entry"
	SunburstHoverinfoPercentParent SunburstHoverinfo = "percent parent"
	// Extras
	SunburstHoverinfoAll  SunburstHoverinfo = "all"
	SunburstHoverinfoNone SunburstHoverinfo = "none"
	SunburstHoverinfoSkip SunburstHoverinfo = "skip"
)

type SunburstTextinfo string

const (
	// Flags
	SunburstTextinfoLabel         SunburstTextinfo = "label"
	SunburstTextinfoText          SunburstTextinfo = "text"
	SunburstTextinfoValue         SunburstTextinfo = "value"
	SunburstTextinfoCurrentPath   SunburstTextinfo = "current path"
	SunburstTextinfoPercentRoot   SunburstTextinfo = "percent root"
	SunburstTextinfoPercentEntry  SunburstTextinfo = "percent entry"
	SunburstTextinfoPercentParent SunburstTextinfo = "percent parent"
	// Extras
	SunburstTextinfoNone SunburstTextinfo = "none"
)

type SurfaceHoverinfo string

const (
	// Flags
	SurfaceHoverinfoX    SurfaceHoverinfo = "x"
	SurfaceHoverinfoY    SurfaceHoverinfo = "y"
	SurfaceHoverinfoZ    SurfaceHoverinfo = "z"
	SurfaceHoverinfoText SurfaceHoverinfo = "text"
	SurfaceHoverinfoName SurfaceHoverinfo = "name"
	// Extras
	SurfaceHoverinfoAll  SurfaceHoverinfo = "all"
	SurfaceHoverinfoNone SurfaceHoverinfo = "none"
	SurfaceHoverinfoSkip SurfaceHoverinfo = "skip"
)

type TableHoverinfo string

const (
	// Flags
	TableHoverinfoX    TableHoverinfo = "x"
	TableHoverinfoY    TableHoverinfo = "y"
	TableHoverinfoZ    TableHoverinfo = "z"
	TableHoverinfoText TableHoverinfo = "text"
	TableHoverinfoName TableHoverinfo = "name"
	// Extras
	TableHoverinfoAll  TableHoverinfo = "all"
	TableHoverinfoNone TableHoverinfo = "none"
	TableHoverinfoSkip TableHoverinfo = "skip"
)

type TreemapCount string

const (
	// Flags
	TreemapCountBranches TreemapCount = "branches"
	TreemapCountLeaves   TreemapCount = "leaves"
	// Extras
)

type TreemapHoverinfo string

const (
	// Flags
	TreemapHoverinfoLabel         TreemapHoverinfo = "label"
	TreemapHoverinfoText          TreemapHoverinfo = "text"
	TreemapHoverinfoValue         TreemapHoverinfo = "value"
	TreemapHoverinfoName          TreemapHoverinfo = "name"
	TreemapHoverinfoCurrentPath   TreemapHoverinfo = "current path"
	TreemapHoverinfoPercentRoot   TreemapHoverinfo = "percent root"
	TreemapHoverinfoPercentEntry  TreemapHoverinfo = "percent entry"
	TreemapHoverinfoPercentParent TreemapHoverinfo = "percent parent"
	// Extras
	TreemapHoverinfoAll  TreemapHoverinfo = "all"
	TreemapHoverinfoNone TreemapHoverinfo = "none"
	TreemapHoverinfoSkip TreemapHoverinfo = "skip"
)

type TreemapTextinfo string

const (
	// Flags
	TreemapTextinfoLabel         TreemapTextinfo = "label"
	TreemapTextinfoText          TreemapTextinfo = "text"
	TreemapTextinfoValue         TreemapTextinfo = "value"
	TreemapTextinfoCurrentPath   TreemapTextinfo = "current path"
	TreemapTextinfoPercentRoot   TreemapTextinfo = "percent root"
	TreemapTextinfoPercentEntry  TreemapTextinfo = "percent entry"
	TreemapTextinfoPercentParent TreemapTextinfo = "percent parent"
	// Extras
	TreemapTextinfoNone TreemapTextinfo = "none"
)

type TreemapTilingFlip string

const (
	// Flags
	TreemapTilingFlipX TreemapTilingFlip = "x"
	TreemapTilingFlipY TreemapTilingFlip = "y"
	// Extras
)

type ViolinHoverinfo string

const (
	// Flags
	ViolinHoverinfoX    ViolinHoverinfo = "x"
	ViolinHoverinfoY    ViolinHoverinfo = "y"
	ViolinHoverinfoZ    ViolinHoverinfo = "z"
	ViolinHoverinfoText ViolinHoverinfo = "text"
	ViolinHoverinfoName ViolinHoverinfo = "name"
	// Extras
	ViolinHoverinfoAll  ViolinHoverinfo = "all"
	ViolinHoverinfoNone ViolinHoverinfo = "none"
	ViolinHoverinfoSkip ViolinHoverinfo = "skip"
)

type ViolinHoveron string

const (
	// Flags
	ViolinHoveronViolins ViolinHoveron = "violins"
	ViolinHoveronPoints  ViolinHoveron = "points"
	ViolinHoveronKde     ViolinHoveron = "kde"
	// Extras
	ViolinHoveronAll ViolinHoveron = "all"
)

type VolumeHoverinfo string

const (
	// Flags
	VolumeHoverinfoX    VolumeHoverinfo = "x"
	VolumeHoverinfoY    VolumeHoverinfo = "y"
	VolumeHoverinfoZ    VolumeHoverinfo = "z"
	VolumeHoverinfoText VolumeHoverinfo = "text"
	VolumeHoverinfoName VolumeHoverinfo = "name"
	// Extras
	VolumeHoverinfoAll  VolumeHoverinfo = "all"
	VolumeHoverinfoNone VolumeHoverinfo = "none"
	VolumeHoverinfoSkip VolumeHoverinfo = "skip"
)

type VolumeSurfacePattern string

const (
	// Flags
	VolumeSurfacePatternA VolumeSurfacePattern = "A"
	VolumeSurfacePatternB VolumeSurfacePattern = "B"
	VolumeSurfacePatternC VolumeSurfacePattern = "C"
	VolumeSurfacePatternD VolumeSurfacePattern = "D"
	VolumeSurfacePatternE VolumeSurfacePattern = "E"
	// Extras
	VolumeSurfacePatternAll  VolumeSurfacePattern = "all"
	VolumeSurfacePatternOdd  VolumeSurfacePattern = "odd"
	VolumeSurfacePatternEven VolumeSurfacePattern = "even"
)

type WaterfallHoverinfo string

const (
	// Flags
	WaterfallHoverinfoName    WaterfallHoverinfo = "name"
	WaterfallHoverinfoX       WaterfallHoverinfo = "x"
	WaterfallHoverinfoY       WaterfallHoverinfo = "y"
	WaterfallHoverinfoText    WaterfallHoverinfo = "text"
	WaterfallHoverinfoInitial WaterfallHoverinfo = "initial"
	WaterfallHoverinfoDelta   WaterfallHoverinfo = "delta"
	WaterfallHoverinfoFinal   WaterfallHoverinfo = "final"
	// Extras
	WaterfallHoverinfoAll  WaterfallHoverinfo = "all"
	WaterfallHoverinfoNone WaterfallHoverinfo = "none"
	WaterfallHoverinfoSkip WaterfallHoverinfo = "skip"
)

type WaterfallTextinfo string

const (
	// Flags
	WaterfallTextinfoLabel   WaterfallTextinfo = "label"
	WaterfallTextinfoText    WaterfallTextinfo = "text"
	WaterfallTextinfoInitial WaterfallTextinfo = "initial"
	WaterfallTextinfoDelta   WaterfallTextinfo = "delta"
	WaterfallTextinfoFinal   WaterfallTextinfo = "final"
	// Extras
	WaterfallTextinfoNone WaterfallTextinfo = "none"
)
