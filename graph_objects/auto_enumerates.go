package grob

// Enumerate section

// AreaHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type AreaHoverlabelAlign string

const (
	AreaHoverlabelAlign_left  AreaHoverlabelAlign = "left"
	AreaHoverlabelAlign_right AreaHoverlabelAlign = "right"
	AreaHoverlabelAlign_auto  AreaHoverlabelAlign = "auto"
)

// AreaMarkerSymbol Area traces are deprecated! Please switch to the *barpolar* trace type. Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type AreaMarkerSymbol string

const (
	AreaMarkerSymbol0                            AreaMarkerSymbol = "0"
	AreaMarkerSymbol_circle                      AreaMarkerSymbol = "circle"
	AreaMarkerSymbol100                          AreaMarkerSymbol = "100"
	AreaMarkerSymbol_circle_open                 AreaMarkerSymbol = "circle-open"
	AreaMarkerSymbol200                          AreaMarkerSymbol = "200"
	AreaMarkerSymbol_circle_dot                  AreaMarkerSymbol = "circle-dot"
	AreaMarkerSymbol300                          AreaMarkerSymbol = "300"
	AreaMarkerSymbol_circle_open_dot             AreaMarkerSymbol = "circle-open-dot"
	AreaMarkerSymbol1                            AreaMarkerSymbol = "1"
	AreaMarkerSymbol_square                      AreaMarkerSymbol = "square"
	AreaMarkerSymbol101                          AreaMarkerSymbol = "101"
	AreaMarkerSymbol_square_open                 AreaMarkerSymbol = "square-open"
	AreaMarkerSymbol201                          AreaMarkerSymbol = "201"
	AreaMarkerSymbol_square_dot                  AreaMarkerSymbol = "square-dot"
	AreaMarkerSymbol301                          AreaMarkerSymbol = "301"
	AreaMarkerSymbol_square_open_dot             AreaMarkerSymbol = "square-open-dot"
	AreaMarkerSymbol2                            AreaMarkerSymbol = "2"
	AreaMarkerSymbol_diamond                     AreaMarkerSymbol = "diamond"
	AreaMarkerSymbol102                          AreaMarkerSymbol = "102"
	AreaMarkerSymbol_diamond_open                AreaMarkerSymbol = "diamond-open"
	AreaMarkerSymbol202                          AreaMarkerSymbol = "202"
	AreaMarkerSymbol_diamond_dot                 AreaMarkerSymbol = "diamond-dot"
	AreaMarkerSymbol302                          AreaMarkerSymbol = "302"
	AreaMarkerSymbol_diamond_open_dot            AreaMarkerSymbol = "diamond-open-dot"
	AreaMarkerSymbol3                            AreaMarkerSymbol = "3"
	AreaMarkerSymbol_cross                       AreaMarkerSymbol = "cross"
	AreaMarkerSymbol103                          AreaMarkerSymbol = "103"
	AreaMarkerSymbol_cross_open                  AreaMarkerSymbol = "cross-open"
	AreaMarkerSymbol203                          AreaMarkerSymbol = "203"
	AreaMarkerSymbol_cross_dot                   AreaMarkerSymbol = "cross-dot"
	AreaMarkerSymbol303                          AreaMarkerSymbol = "303"
	AreaMarkerSymbol_cross_open_dot              AreaMarkerSymbol = "cross-open-dot"
	AreaMarkerSymbol4                            AreaMarkerSymbol = "4"
	AreaMarkerSymbol_x                           AreaMarkerSymbol = "x"
	AreaMarkerSymbol104                          AreaMarkerSymbol = "104"
	AreaMarkerSymbol_x_open                      AreaMarkerSymbol = "x-open"
	AreaMarkerSymbol204                          AreaMarkerSymbol = "204"
	AreaMarkerSymbol_x_dot                       AreaMarkerSymbol = "x-dot"
	AreaMarkerSymbol304                          AreaMarkerSymbol = "304"
	AreaMarkerSymbol_x_open_dot                  AreaMarkerSymbol = "x-open-dot"
	AreaMarkerSymbol5                            AreaMarkerSymbol = "5"
	AreaMarkerSymbol_triangle_up                 AreaMarkerSymbol = "triangle-up"
	AreaMarkerSymbol105                          AreaMarkerSymbol = "105"
	AreaMarkerSymbol_triangle_up_open            AreaMarkerSymbol = "triangle-up-open"
	AreaMarkerSymbol205                          AreaMarkerSymbol = "205"
	AreaMarkerSymbol_triangle_up_dot             AreaMarkerSymbol = "triangle-up-dot"
	AreaMarkerSymbol305                          AreaMarkerSymbol = "305"
	AreaMarkerSymbol_triangle_up_open_dot        AreaMarkerSymbol = "triangle-up-open-dot"
	AreaMarkerSymbol6                            AreaMarkerSymbol = "6"
	AreaMarkerSymbol_triangle_down               AreaMarkerSymbol = "triangle-down"
	AreaMarkerSymbol106                          AreaMarkerSymbol = "106"
	AreaMarkerSymbol_triangle_down_open          AreaMarkerSymbol = "triangle-down-open"
	AreaMarkerSymbol206                          AreaMarkerSymbol = "206"
	AreaMarkerSymbol_triangle_down_dot           AreaMarkerSymbol = "triangle-down-dot"
	AreaMarkerSymbol306                          AreaMarkerSymbol = "306"
	AreaMarkerSymbol_triangle_down_open_dot      AreaMarkerSymbol = "triangle-down-open-dot"
	AreaMarkerSymbol7                            AreaMarkerSymbol = "7"
	AreaMarkerSymbol_triangle_left               AreaMarkerSymbol = "triangle-left"
	AreaMarkerSymbol107                          AreaMarkerSymbol = "107"
	AreaMarkerSymbol_triangle_left_open          AreaMarkerSymbol = "triangle-left-open"
	AreaMarkerSymbol207                          AreaMarkerSymbol = "207"
	AreaMarkerSymbol_triangle_left_dot           AreaMarkerSymbol = "triangle-left-dot"
	AreaMarkerSymbol307                          AreaMarkerSymbol = "307"
	AreaMarkerSymbol_triangle_left_open_dot      AreaMarkerSymbol = "triangle-left-open-dot"
	AreaMarkerSymbol8                            AreaMarkerSymbol = "8"
	AreaMarkerSymbol_triangle_right              AreaMarkerSymbol = "triangle-right"
	AreaMarkerSymbol108                          AreaMarkerSymbol = "108"
	AreaMarkerSymbol_triangle_right_open         AreaMarkerSymbol = "triangle-right-open"
	AreaMarkerSymbol208                          AreaMarkerSymbol = "208"
	AreaMarkerSymbol_triangle_right_dot          AreaMarkerSymbol = "triangle-right-dot"
	AreaMarkerSymbol308                          AreaMarkerSymbol = "308"
	AreaMarkerSymbol_triangle_right_open_dot     AreaMarkerSymbol = "triangle-right-open-dot"
	AreaMarkerSymbol9                            AreaMarkerSymbol = "9"
	AreaMarkerSymbol_triangle_ne                 AreaMarkerSymbol = "triangle-ne"
	AreaMarkerSymbol109                          AreaMarkerSymbol = "109"
	AreaMarkerSymbol_triangle_ne_open            AreaMarkerSymbol = "triangle-ne-open"
	AreaMarkerSymbol209                          AreaMarkerSymbol = "209"
	AreaMarkerSymbol_triangle_ne_dot             AreaMarkerSymbol = "triangle-ne-dot"
	AreaMarkerSymbol309                          AreaMarkerSymbol = "309"
	AreaMarkerSymbol_triangle_ne_open_dot        AreaMarkerSymbol = "triangle-ne-open-dot"
	AreaMarkerSymbol10                           AreaMarkerSymbol = "10"
	AreaMarkerSymbol_triangle_se                 AreaMarkerSymbol = "triangle-se"
	AreaMarkerSymbol110                          AreaMarkerSymbol = "110"
	AreaMarkerSymbol_triangle_se_open            AreaMarkerSymbol = "triangle-se-open"
	AreaMarkerSymbol210                          AreaMarkerSymbol = "210"
	AreaMarkerSymbol_triangle_se_dot             AreaMarkerSymbol = "triangle-se-dot"
	AreaMarkerSymbol310                          AreaMarkerSymbol = "310"
	AreaMarkerSymbol_triangle_se_open_dot        AreaMarkerSymbol = "triangle-se-open-dot"
	AreaMarkerSymbol11                           AreaMarkerSymbol = "11"
	AreaMarkerSymbol_triangle_sw                 AreaMarkerSymbol = "triangle-sw"
	AreaMarkerSymbol111                          AreaMarkerSymbol = "111"
	AreaMarkerSymbol_triangle_sw_open            AreaMarkerSymbol = "triangle-sw-open"
	AreaMarkerSymbol211                          AreaMarkerSymbol = "211"
	AreaMarkerSymbol_triangle_sw_dot             AreaMarkerSymbol = "triangle-sw-dot"
	AreaMarkerSymbol311                          AreaMarkerSymbol = "311"
	AreaMarkerSymbol_triangle_sw_open_dot        AreaMarkerSymbol = "triangle-sw-open-dot"
	AreaMarkerSymbol12                           AreaMarkerSymbol = "12"
	AreaMarkerSymbol_triangle_nw                 AreaMarkerSymbol = "triangle-nw"
	AreaMarkerSymbol112                          AreaMarkerSymbol = "112"
	AreaMarkerSymbol_triangle_nw_open            AreaMarkerSymbol = "triangle-nw-open"
	AreaMarkerSymbol212                          AreaMarkerSymbol = "212"
	AreaMarkerSymbol_triangle_nw_dot             AreaMarkerSymbol = "triangle-nw-dot"
	AreaMarkerSymbol312                          AreaMarkerSymbol = "312"
	AreaMarkerSymbol_triangle_nw_open_dot        AreaMarkerSymbol = "triangle-nw-open-dot"
	AreaMarkerSymbol13                           AreaMarkerSymbol = "13"
	AreaMarkerSymbol_pentagon                    AreaMarkerSymbol = "pentagon"
	AreaMarkerSymbol113                          AreaMarkerSymbol = "113"
	AreaMarkerSymbol_pentagon_open               AreaMarkerSymbol = "pentagon-open"
	AreaMarkerSymbol213                          AreaMarkerSymbol = "213"
	AreaMarkerSymbol_pentagon_dot                AreaMarkerSymbol = "pentagon-dot"
	AreaMarkerSymbol313                          AreaMarkerSymbol = "313"
	AreaMarkerSymbol_pentagon_open_dot           AreaMarkerSymbol = "pentagon-open-dot"
	AreaMarkerSymbol14                           AreaMarkerSymbol = "14"
	AreaMarkerSymbol_hexagon                     AreaMarkerSymbol = "hexagon"
	AreaMarkerSymbol114                          AreaMarkerSymbol = "114"
	AreaMarkerSymbol_hexagon_open                AreaMarkerSymbol = "hexagon-open"
	AreaMarkerSymbol214                          AreaMarkerSymbol = "214"
	AreaMarkerSymbol_hexagon_dot                 AreaMarkerSymbol = "hexagon-dot"
	AreaMarkerSymbol314                          AreaMarkerSymbol = "314"
	AreaMarkerSymbol_hexagon_open_dot            AreaMarkerSymbol = "hexagon-open-dot"
	AreaMarkerSymbol15                           AreaMarkerSymbol = "15"
	AreaMarkerSymbol_hexagon2                    AreaMarkerSymbol = "hexagon2"
	AreaMarkerSymbol115                          AreaMarkerSymbol = "115"
	AreaMarkerSymbol_hexagon2_open               AreaMarkerSymbol = "hexagon2-open"
	AreaMarkerSymbol215                          AreaMarkerSymbol = "215"
	AreaMarkerSymbol_hexagon2_dot                AreaMarkerSymbol = "hexagon2-dot"
	AreaMarkerSymbol315                          AreaMarkerSymbol = "315"
	AreaMarkerSymbol_hexagon2_open_dot           AreaMarkerSymbol = "hexagon2-open-dot"
	AreaMarkerSymbol16                           AreaMarkerSymbol = "16"
	AreaMarkerSymbol_octagon                     AreaMarkerSymbol = "octagon"
	AreaMarkerSymbol116                          AreaMarkerSymbol = "116"
	AreaMarkerSymbol_octagon_open                AreaMarkerSymbol = "octagon-open"
	AreaMarkerSymbol216                          AreaMarkerSymbol = "216"
	AreaMarkerSymbol_octagon_dot                 AreaMarkerSymbol = "octagon-dot"
	AreaMarkerSymbol316                          AreaMarkerSymbol = "316"
	AreaMarkerSymbol_octagon_open_dot            AreaMarkerSymbol = "octagon-open-dot"
	AreaMarkerSymbol17                           AreaMarkerSymbol = "17"
	AreaMarkerSymbol_star                        AreaMarkerSymbol = "star"
	AreaMarkerSymbol117                          AreaMarkerSymbol = "117"
	AreaMarkerSymbol_star_open                   AreaMarkerSymbol = "star-open"
	AreaMarkerSymbol217                          AreaMarkerSymbol = "217"
	AreaMarkerSymbol_star_dot                    AreaMarkerSymbol = "star-dot"
	AreaMarkerSymbol317                          AreaMarkerSymbol = "317"
	AreaMarkerSymbol_star_open_dot               AreaMarkerSymbol = "star-open-dot"
	AreaMarkerSymbol18                           AreaMarkerSymbol = "18"
	AreaMarkerSymbol_hexagram                    AreaMarkerSymbol = "hexagram"
	AreaMarkerSymbol118                          AreaMarkerSymbol = "118"
	AreaMarkerSymbol_hexagram_open               AreaMarkerSymbol = "hexagram-open"
	AreaMarkerSymbol218                          AreaMarkerSymbol = "218"
	AreaMarkerSymbol_hexagram_dot                AreaMarkerSymbol = "hexagram-dot"
	AreaMarkerSymbol318                          AreaMarkerSymbol = "318"
	AreaMarkerSymbol_hexagram_open_dot           AreaMarkerSymbol = "hexagram-open-dot"
	AreaMarkerSymbol19                           AreaMarkerSymbol = "19"
	AreaMarkerSymbol_star_triangle_up            AreaMarkerSymbol = "star-triangle-up"
	AreaMarkerSymbol119                          AreaMarkerSymbol = "119"
	AreaMarkerSymbol_star_triangle_up_open       AreaMarkerSymbol = "star-triangle-up-open"
	AreaMarkerSymbol219                          AreaMarkerSymbol = "219"
	AreaMarkerSymbol_star_triangle_up_dot        AreaMarkerSymbol = "star-triangle-up-dot"
	AreaMarkerSymbol319                          AreaMarkerSymbol = "319"
	AreaMarkerSymbol_star_triangle_up_open_dot   AreaMarkerSymbol = "star-triangle-up-open-dot"
	AreaMarkerSymbol20                           AreaMarkerSymbol = "20"
	AreaMarkerSymbol_star_triangle_down          AreaMarkerSymbol = "star-triangle-down"
	AreaMarkerSymbol120                          AreaMarkerSymbol = "120"
	AreaMarkerSymbol_star_triangle_down_open     AreaMarkerSymbol = "star-triangle-down-open"
	AreaMarkerSymbol220                          AreaMarkerSymbol = "220"
	AreaMarkerSymbol_star_triangle_down_dot      AreaMarkerSymbol = "star-triangle-down-dot"
	AreaMarkerSymbol320                          AreaMarkerSymbol = "320"
	AreaMarkerSymbol_star_triangle_down_open_dot AreaMarkerSymbol = "star-triangle-down-open-dot"
	AreaMarkerSymbol21                           AreaMarkerSymbol = "21"
	AreaMarkerSymbol_star_square                 AreaMarkerSymbol = "star-square"
	AreaMarkerSymbol121                          AreaMarkerSymbol = "121"
	AreaMarkerSymbol_star_square_open            AreaMarkerSymbol = "star-square-open"
	AreaMarkerSymbol221                          AreaMarkerSymbol = "221"
	AreaMarkerSymbol_star_square_dot             AreaMarkerSymbol = "star-square-dot"
	AreaMarkerSymbol321                          AreaMarkerSymbol = "321"
	AreaMarkerSymbol_star_square_open_dot        AreaMarkerSymbol = "star-square-open-dot"
	AreaMarkerSymbol22                           AreaMarkerSymbol = "22"
	AreaMarkerSymbol_star_diamond                AreaMarkerSymbol = "star-diamond"
	AreaMarkerSymbol122                          AreaMarkerSymbol = "122"
	AreaMarkerSymbol_star_diamond_open           AreaMarkerSymbol = "star-diamond-open"
	AreaMarkerSymbol222                          AreaMarkerSymbol = "222"
	AreaMarkerSymbol_star_diamond_dot            AreaMarkerSymbol = "star-diamond-dot"
	AreaMarkerSymbol322                          AreaMarkerSymbol = "322"
	AreaMarkerSymbol_star_diamond_open_dot       AreaMarkerSymbol = "star-diamond-open-dot"
	AreaMarkerSymbol23                           AreaMarkerSymbol = "23"
	AreaMarkerSymbol_diamond_tall                AreaMarkerSymbol = "diamond-tall"
	AreaMarkerSymbol123                          AreaMarkerSymbol = "123"
	AreaMarkerSymbol_diamond_tall_open           AreaMarkerSymbol = "diamond-tall-open"
	AreaMarkerSymbol223                          AreaMarkerSymbol = "223"
	AreaMarkerSymbol_diamond_tall_dot            AreaMarkerSymbol = "diamond-tall-dot"
	AreaMarkerSymbol323                          AreaMarkerSymbol = "323"
	AreaMarkerSymbol_diamond_tall_open_dot       AreaMarkerSymbol = "diamond-tall-open-dot"
	AreaMarkerSymbol24                           AreaMarkerSymbol = "24"
	AreaMarkerSymbol_diamond_wide                AreaMarkerSymbol = "diamond-wide"
	AreaMarkerSymbol124                          AreaMarkerSymbol = "124"
	AreaMarkerSymbol_diamond_wide_open           AreaMarkerSymbol = "diamond-wide-open"
	AreaMarkerSymbol224                          AreaMarkerSymbol = "224"
	AreaMarkerSymbol_diamond_wide_dot            AreaMarkerSymbol = "diamond-wide-dot"
	AreaMarkerSymbol324                          AreaMarkerSymbol = "324"
	AreaMarkerSymbol_diamond_wide_open_dot       AreaMarkerSymbol = "diamond-wide-open-dot"
	AreaMarkerSymbol25                           AreaMarkerSymbol = "25"
	AreaMarkerSymbol_hourglass                   AreaMarkerSymbol = "hourglass"
	AreaMarkerSymbol125                          AreaMarkerSymbol = "125"
	AreaMarkerSymbol_hourglass_open              AreaMarkerSymbol = "hourglass-open"
	AreaMarkerSymbol26                           AreaMarkerSymbol = "26"
	AreaMarkerSymbol_bowtie                      AreaMarkerSymbol = "bowtie"
	AreaMarkerSymbol126                          AreaMarkerSymbol = "126"
	AreaMarkerSymbol_bowtie_open                 AreaMarkerSymbol = "bowtie-open"
	AreaMarkerSymbol27                           AreaMarkerSymbol = "27"
	AreaMarkerSymbol_circle_cross                AreaMarkerSymbol = "circle-cross"
	AreaMarkerSymbol127                          AreaMarkerSymbol = "127"
	AreaMarkerSymbol_circle_cross_open           AreaMarkerSymbol = "circle-cross-open"
	AreaMarkerSymbol28                           AreaMarkerSymbol = "28"
	AreaMarkerSymbol_circle_x                    AreaMarkerSymbol = "circle-x"
	AreaMarkerSymbol128                          AreaMarkerSymbol = "128"
	AreaMarkerSymbol_circle_x_open               AreaMarkerSymbol = "circle-x-open"
	AreaMarkerSymbol29                           AreaMarkerSymbol = "29"
	AreaMarkerSymbol_square_cross                AreaMarkerSymbol = "square-cross"
	AreaMarkerSymbol129                          AreaMarkerSymbol = "129"
	AreaMarkerSymbol_square_cross_open           AreaMarkerSymbol = "square-cross-open"
	AreaMarkerSymbol30                           AreaMarkerSymbol = "30"
	AreaMarkerSymbol_square_x                    AreaMarkerSymbol = "square-x"
	AreaMarkerSymbol130                          AreaMarkerSymbol = "130"
	AreaMarkerSymbol_square_x_open               AreaMarkerSymbol = "square-x-open"
	AreaMarkerSymbol31                           AreaMarkerSymbol = "31"
	AreaMarkerSymbol_diamond_cross               AreaMarkerSymbol = "diamond-cross"
	AreaMarkerSymbol131                          AreaMarkerSymbol = "131"
	AreaMarkerSymbol_diamond_cross_open          AreaMarkerSymbol = "diamond-cross-open"
	AreaMarkerSymbol32                           AreaMarkerSymbol = "32"
	AreaMarkerSymbol_diamond_x                   AreaMarkerSymbol = "diamond-x"
	AreaMarkerSymbol132                          AreaMarkerSymbol = "132"
	AreaMarkerSymbol_diamond_x_open              AreaMarkerSymbol = "diamond-x-open"
	AreaMarkerSymbol33                           AreaMarkerSymbol = "33"
	AreaMarkerSymbol_cross_thin                  AreaMarkerSymbol = "cross-thin"
	AreaMarkerSymbol133                          AreaMarkerSymbol = "133"
	AreaMarkerSymbol_cross_thin_open             AreaMarkerSymbol = "cross-thin-open"
	AreaMarkerSymbol34                           AreaMarkerSymbol = "34"
	AreaMarkerSymbol_x_thin                      AreaMarkerSymbol = "x-thin"
	AreaMarkerSymbol134                          AreaMarkerSymbol = "134"
	AreaMarkerSymbol_x_thin_open                 AreaMarkerSymbol = "x-thin-open"
	AreaMarkerSymbol35                           AreaMarkerSymbol = "35"
	AreaMarkerSymbol_asterisk                    AreaMarkerSymbol = "asterisk"
	AreaMarkerSymbol135                          AreaMarkerSymbol = "135"
	AreaMarkerSymbol_asterisk_open               AreaMarkerSymbol = "asterisk-open"
	AreaMarkerSymbol36                           AreaMarkerSymbol = "36"
	AreaMarkerSymbol_hash                        AreaMarkerSymbol = "hash"
	AreaMarkerSymbol136                          AreaMarkerSymbol = "136"
	AreaMarkerSymbol_hash_open                   AreaMarkerSymbol = "hash-open"
	AreaMarkerSymbol236                          AreaMarkerSymbol = "236"
	AreaMarkerSymbol_hash_dot                    AreaMarkerSymbol = "hash-dot"
	AreaMarkerSymbol336                          AreaMarkerSymbol = "336"
	AreaMarkerSymbol_hash_open_dot               AreaMarkerSymbol = "hash-open-dot"
	AreaMarkerSymbol37                           AreaMarkerSymbol = "37"
	AreaMarkerSymbol_y_up                        AreaMarkerSymbol = "y-up"
	AreaMarkerSymbol137                          AreaMarkerSymbol = "137"
	AreaMarkerSymbol_y_up_open                   AreaMarkerSymbol = "y-up-open"
	AreaMarkerSymbol38                           AreaMarkerSymbol = "38"
	AreaMarkerSymbol_y_down                      AreaMarkerSymbol = "y-down"
	AreaMarkerSymbol138                          AreaMarkerSymbol = "138"
	AreaMarkerSymbol_y_down_open                 AreaMarkerSymbol = "y-down-open"
	AreaMarkerSymbol39                           AreaMarkerSymbol = "39"
	AreaMarkerSymbol_y_left                      AreaMarkerSymbol = "y-left"
	AreaMarkerSymbol139                          AreaMarkerSymbol = "139"
	AreaMarkerSymbol_y_left_open                 AreaMarkerSymbol = "y-left-open"
	AreaMarkerSymbol40                           AreaMarkerSymbol = "40"
	AreaMarkerSymbol_y_right                     AreaMarkerSymbol = "y-right"
	AreaMarkerSymbol140                          AreaMarkerSymbol = "140"
	AreaMarkerSymbol_y_right_open                AreaMarkerSymbol = "y-right-open"
	AreaMarkerSymbol41                           AreaMarkerSymbol = "41"
	AreaMarkerSymbol_line_ew                     AreaMarkerSymbol = "line-ew"
	AreaMarkerSymbol141                          AreaMarkerSymbol = "141"
	AreaMarkerSymbol_line_ew_open                AreaMarkerSymbol = "line-ew-open"
	AreaMarkerSymbol42                           AreaMarkerSymbol = "42"
	AreaMarkerSymbol_line_ns                     AreaMarkerSymbol = "line-ns"
	AreaMarkerSymbol142                          AreaMarkerSymbol = "142"
	AreaMarkerSymbol_line_ns_open                AreaMarkerSymbol = "line-ns-open"
	AreaMarkerSymbol43                           AreaMarkerSymbol = "43"
	AreaMarkerSymbol_line_ne                     AreaMarkerSymbol = "line-ne"
	AreaMarkerSymbol143                          AreaMarkerSymbol = "143"
	AreaMarkerSymbol_line_ne_open                AreaMarkerSymbol = "line-ne-open"
	AreaMarkerSymbol44                           AreaMarkerSymbol = "44"
	AreaMarkerSymbol_line_nw                     AreaMarkerSymbol = "line-nw"
	AreaMarkerSymbol144                          AreaMarkerSymbol = "144"
	AreaMarkerSymbol_line_nw_open                AreaMarkerSymbol = "line-nw-open"
)

// AreaVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type AreaVisible interface{}

var (
	AreaVisible_True       AreaVisible = true
	AreaVisible_False      AreaVisible = false
	AreaVisible_legendonly AreaVisible = "legendonly"
)

// BarConstraintext Constrain the size of text inside or outside a bar to be no larger than the bar itself.
type BarConstraintext string

const (
	BarConstraintext_inside  BarConstraintext = "inside"
	BarConstraintext_outside BarConstraintext = "outside"
	BarConstraintext_both    BarConstraintext = "both"
	BarConstraintext_none    BarConstraintext = "none"
)

// BarErrorXType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type BarErrorXType string

const (
	BarErrorXType_percent  BarErrorXType = "percent"
	BarErrorXType_constant BarErrorXType = "constant"
	BarErrorXType_sqrt     BarErrorXType = "sqrt"
	BarErrorXType_data     BarErrorXType = "data"
)

// BarErrorYType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type BarErrorYType string

const (
	BarErrorYType_percent  BarErrorYType = "percent"
	BarErrorYType_constant BarErrorYType = "constant"
	BarErrorYType_sqrt     BarErrorYType = "sqrt"
	BarErrorYType_data     BarErrorYType = "data"
)

// BarHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type BarHoverlabelAlign string

const (
	BarHoverlabelAlign_left  BarHoverlabelAlign = "left"
	BarHoverlabelAlign_right BarHoverlabelAlign = "right"
	BarHoverlabelAlign_auto  BarHoverlabelAlign = "auto"
)

// BarInsidetextanchor Determines if texts are kept at center or start/end points in `textposition` *inside* mode.
type BarInsidetextanchor string

const (
	BarInsidetextanchor_end    BarInsidetextanchor = "end"
	BarInsidetextanchor_middle BarInsidetextanchor = "middle"
	BarInsidetextanchor_start  BarInsidetextanchor = "start"
)

// BarMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type BarMarkerColorbarExponentformat string

const (
	BarMarkerColorbarExponentformat_none  BarMarkerColorbarExponentformat = "none"
	BarMarkerColorbarExponentformat_e     BarMarkerColorbarExponentformat = "e"
	BarMarkerColorbarExponentformat_E     BarMarkerColorbarExponentformat = "E"
	BarMarkerColorbarExponentformat_power BarMarkerColorbarExponentformat = "power"
	BarMarkerColorbarExponentformat_SI    BarMarkerColorbarExponentformat = "SI"
	BarMarkerColorbarExponentformat_B     BarMarkerColorbarExponentformat = "B"
)

// BarMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type BarMarkerColorbarLenmode string

const (
	BarMarkerColorbarLenmode_fraction BarMarkerColorbarLenmode = "fraction"
	BarMarkerColorbarLenmode_pixels   BarMarkerColorbarLenmode = "pixels"
)

// BarMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type BarMarkerColorbarShowexponent string

const (
	BarMarkerColorbarShowexponent_all   BarMarkerColorbarShowexponent = "all"
	BarMarkerColorbarShowexponent_first BarMarkerColorbarShowexponent = "first"
	BarMarkerColorbarShowexponent_last  BarMarkerColorbarShowexponent = "last"
	BarMarkerColorbarShowexponent_none  BarMarkerColorbarShowexponent = "none"
)

// BarMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type BarMarkerColorbarShowtickprefix string

const (
	BarMarkerColorbarShowtickprefix_all   BarMarkerColorbarShowtickprefix = "all"
	BarMarkerColorbarShowtickprefix_first BarMarkerColorbarShowtickprefix = "first"
	BarMarkerColorbarShowtickprefix_last  BarMarkerColorbarShowtickprefix = "last"
	BarMarkerColorbarShowtickprefix_none  BarMarkerColorbarShowtickprefix = "none"
)

// BarMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type BarMarkerColorbarShowticksuffix string

const (
	BarMarkerColorbarShowticksuffix_all   BarMarkerColorbarShowticksuffix = "all"
	BarMarkerColorbarShowticksuffix_first BarMarkerColorbarShowticksuffix = "first"
	BarMarkerColorbarShowticksuffix_last  BarMarkerColorbarShowticksuffix = "last"
	BarMarkerColorbarShowticksuffix_none  BarMarkerColorbarShowticksuffix = "none"
)

// BarMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type BarMarkerColorbarThicknessmode string

const (
	BarMarkerColorbarThicknessmode_fraction BarMarkerColorbarThicknessmode = "fraction"
	BarMarkerColorbarThicknessmode_pixels   BarMarkerColorbarThicknessmode = "pixels"
)

// BarMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type BarMarkerColorbarTickmode string

const (
	BarMarkerColorbarTickmode_auto   BarMarkerColorbarTickmode = "auto"
	BarMarkerColorbarTickmode_linear BarMarkerColorbarTickmode = "linear"
	BarMarkerColorbarTickmode_array  BarMarkerColorbarTickmode = "array"
)

// BarMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type BarMarkerColorbarTicks string

const (
	BarMarkerColorbarTicks_outside BarMarkerColorbarTicks = "outside"
	BarMarkerColorbarTicks_inside  BarMarkerColorbarTicks = "inside"
)

// BarMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type BarMarkerColorbarTitleSide string

const (
	BarMarkerColorbarTitleSide_right  BarMarkerColorbarTitleSide = "right"
	BarMarkerColorbarTitleSide_top    BarMarkerColorbarTitleSide = "top"
	BarMarkerColorbarTitleSide_bottom BarMarkerColorbarTitleSide = "bottom"
)

// BarMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type BarMarkerColorbarXanchor string

const (
	BarMarkerColorbarXanchor_left   BarMarkerColorbarXanchor = "left"
	BarMarkerColorbarXanchor_center BarMarkerColorbarXanchor = "center"
	BarMarkerColorbarXanchor_right  BarMarkerColorbarXanchor = "right"
)

// BarMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type BarMarkerColorbarYanchor string

const (
	BarMarkerColorbarYanchor_top    BarMarkerColorbarYanchor = "top"
	BarMarkerColorbarYanchor_middle BarMarkerColorbarYanchor = "middle"
	BarMarkerColorbarYanchor_bottom BarMarkerColorbarYanchor = "bottom"
)

// BarOrientation Sets the orientation of the bars. With *v* (*h*), the value of the each bar spans along the vertical (horizontal).
type BarOrientation string

const (
	BarOrientation_v BarOrientation = "v"
	BarOrientation_h BarOrientation = "h"
)

// BarTextposition Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside.
type BarTextposition string

const (
	BarTextposition_inside  BarTextposition = "inside"
	BarTextposition_outside BarTextposition = "outside"
	BarTextposition_auto    BarTextposition = "auto"
	BarTextposition_none    BarTextposition = "none"
)

// BarVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type BarVisible interface{}

var (
	BarVisible_True       BarVisible = true
	BarVisible_False      BarVisible = false
	BarVisible_legendonly BarVisible = "legendonly"
)

// BarXcalendar Sets the calendar system to use with `x` date data.
type BarXcalendar string

const (
	BarXcalendar_gregorian  BarXcalendar = "gregorian"
	BarXcalendar_chinese    BarXcalendar = "chinese"
	BarXcalendar_coptic     BarXcalendar = "coptic"
	BarXcalendar_discworld  BarXcalendar = "discworld"
	BarXcalendar_ethiopian  BarXcalendar = "ethiopian"
	BarXcalendar_hebrew     BarXcalendar = "hebrew"
	BarXcalendar_islamic    BarXcalendar = "islamic"
	BarXcalendar_julian     BarXcalendar = "julian"
	BarXcalendar_mayan      BarXcalendar = "mayan"
	BarXcalendar_nanakshahi BarXcalendar = "nanakshahi"
	BarXcalendar_nepali     BarXcalendar = "nepali"
	BarXcalendar_persian    BarXcalendar = "persian"
	BarXcalendar_jalali     BarXcalendar = "jalali"
	BarXcalendar_taiwan     BarXcalendar = "taiwan"
	BarXcalendar_thai       BarXcalendar = "thai"
	BarXcalendar_ummalqura  BarXcalendar = "ummalqura"
)

// BarYcalendar Sets the calendar system to use with `y` date data.
type BarYcalendar string

const (
	BarYcalendar_gregorian  BarYcalendar = "gregorian"
	BarYcalendar_chinese    BarYcalendar = "chinese"
	BarYcalendar_coptic     BarYcalendar = "coptic"
	BarYcalendar_discworld  BarYcalendar = "discworld"
	BarYcalendar_ethiopian  BarYcalendar = "ethiopian"
	BarYcalendar_hebrew     BarYcalendar = "hebrew"
	BarYcalendar_islamic    BarYcalendar = "islamic"
	BarYcalendar_julian     BarYcalendar = "julian"
	BarYcalendar_mayan      BarYcalendar = "mayan"
	BarYcalendar_nanakshahi BarYcalendar = "nanakshahi"
	BarYcalendar_nepali     BarYcalendar = "nepali"
	BarYcalendar_persian    BarYcalendar = "persian"
	BarYcalendar_jalali     BarYcalendar = "jalali"
	BarYcalendar_taiwan     BarYcalendar = "taiwan"
	BarYcalendar_thai       BarYcalendar = "thai"
	BarYcalendar_ummalqura  BarYcalendar = "ummalqura"
)

// BarpolarHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type BarpolarHoverlabelAlign string

const (
	BarpolarHoverlabelAlign_left  BarpolarHoverlabelAlign = "left"
	BarpolarHoverlabelAlign_right BarpolarHoverlabelAlign = "right"
	BarpolarHoverlabelAlign_auto  BarpolarHoverlabelAlign = "auto"
)

// BarpolarMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type BarpolarMarkerColorbarExponentformat string

const (
	BarpolarMarkerColorbarExponentformat_none  BarpolarMarkerColorbarExponentformat = "none"
	BarpolarMarkerColorbarExponentformat_e     BarpolarMarkerColorbarExponentformat = "e"
	BarpolarMarkerColorbarExponentformat_E     BarpolarMarkerColorbarExponentformat = "E"
	BarpolarMarkerColorbarExponentformat_power BarpolarMarkerColorbarExponentformat = "power"
	BarpolarMarkerColorbarExponentformat_SI    BarpolarMarkerColorbarExponentformat = "SI"
	BarpolarMarkerColorbarExponentformat_B     BarpolarMarkerColorbarExponentformat = "B"
)

// BarpolarMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type BarpolarMarkerColorbarLenmode string

const (
	BarpolarMarkerColorbarLenmode_fraction BarpolarMarkerColorbarLenmode = "fraction"
	BarpolarMarkerColorbarLenmode_pixels   BarpolarMarkerColorbarLenmode = "pixels"
)

// BarpolarMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type BarpolarMarkerColorbarShowexponent string

const (
	BarpolarMarkerColorbarShowexponent_all   BarpolarMarkerColorbarShowexponent = "all"
	BarpolarMarkerColorbarShowexponent_first BarpolarMarkerColorbarShowexponent = "first"
	BarpolarMarkerColorbarShowexponent_last  BarpolarMarkerColorbarShowexponent = "last"
	BarpolarMarkerColorbarShowexponent_none  BarpolarMarkerColorbarShowexponent = "none"
)

// BarpolarMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type BarpolarMarkerColorbarShowtickprefix string

const (
	BarpolarMarkerColorbarShowtickprefix_all   BarpolarMarkerColorbarShowtickprefix = "all"
	BarpolarMarkerColorbarShowtickprefix_first BarpolarMarkerColorbarShowtickprefix = "first"
	BarpolarMarkerColorbarShowtickprefix_last  BarpolarMarkerColorbarShowtickprefix = "last"
	BarpolarMarkerColorbarShowtickprefix_none  BarpolarMarkerColorbarShowtickprefix = "none"
)

// BarpolarMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type BarpolarMarkerColorbarShowticksuffix string

const (
	BarpolarMarkerColorbarShowticksuffix_all   BarpolarMarkerColorbarShowticksuffix = "all"
	BarpolarMarkerColorbarShowticksuffix_first BarpolarMarkerColorbarShowticksuffix = "first"
	BarpolarMarkerColorbarShowticksuffix_last  BarpolarMarkerColorbarShowticksuffix = "last"
	BarpolarMarkerColorbarShowticksuffix_none  BarpolarMarkerColorbarShowticksuffix = "none"
)

// BarpolarMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type BarpolarMarkerColorbarThicknessmode string

const (
	BarpolarMarkerColorbarThicknessmode_fraction BarpolarMarkerColorbarThicknessmode = "fraction"
	BarpolarMarkerColorbarThicknessmode_pixels   BarpolarMarkerColorbarThicknessmode = "pixels"
)

// BarpolarMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type BarpolarMarkerColorbarTickmode string

const (
	BarpolarMarkerColorbarTickmode_auto   BarpolarMarkerColorbarTickmode = "auto"
	BarpolarMarkerColorbarTickmode_linear BarpolarMarkerColorbarTickmode = "linear"
	BarpolarMarkerColorbarTickmode_array  BarpolarMarkerColorbarTickmode = "array"
)

// BarpolarMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type BarpolarMarkerColorbarTicks string

const (
	BarpolarMarkerColorbarTicks_outside BarpolarMarkerColorbarTicks = "outside"
	BarpolarMarkerColorbarTicks_inside  BarpolarMarkerColorbarTicks = "inside"
)

// BarpolarMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type BarpolarMarkerColorbarTitleSide string

const (
	BarpolarMarkerColorbarTitleSide_right  BarpolarMarkerColorbarTitleSide = "right"
	BarpolarMarkerColorbarTitleSide_top    BarpolarMarkerColorbarTitleSide = "top"
	BarpolarMarkerColorbarTitleSide_bottom BarpolarMarkerColorbarTitleSide = "bottom"
)

// BarpolarMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type BarpolarMarkerColorbarXanchor string

const (
	BarpolarMarkerColorbarXanchor_left   BarpolarMarkerColorbarXanchor = "left"
	BarpolarMarkerColorbarXanchor_center BarpolarMarkerColorbarXanchor = "center"
	BarpolarMarkerColorbarXanchor_right  BarpolarMarkerColorbarXanchor = "right"
)

// BarpolarMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type BarpolarMarkerColorbarYanchor string

const (
	BarpolarMarkerColorbarYanchor_top    BarpolarMarkerColorbarYanchor = "top"
	BarpolarMarkerColorbarYanchor_middle BarpolarMarkerColorbarYanchor = "middle"
	BarpolarMarkerColorbarYanchor_bottom BarpolarMarkerColorbarYanchor = "bottom"
)

// BarpolarThetaunit Sets the unit of input *theta* values. Has an effect only when on *linear* angular axes.
type BarpolarThetaunit string

const (
	BarpolarThetaunit_radians  BarpolarThetaunit = "radians"
	BarpolarThetaunit_degrees  BarpolarThetaunit = "degrees"
	BarpolarThetaunit_gradians BarpolarThetaunit = "gradians"
)

// BarpolarVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type BarpolarVisible interface{}

var (
	BarpolarVisible_True       BarpolarVisible = true
	BarpolarVisible_False      BarpolarVisible = false
	BarpolarVisible_legendonly BarpolarVisible = "legendonly"
)

// BoxBoxmean If *true*, the mean of the box(es)' underlying distribution is drawn as a dashed line inside the box(es). If *sd* the standard deviation is also drawn. Defaults to *true* when `mean` is set. Defaults to *sd* when `sd` is set Otherwise defaults to *false*.
type BoxBoxmean interface{}

var (
	BoxBoxmean_True  BoxBoxmean = true
	BoxBoxmean_sd    BoxBoxmean = "sd"
	BoxBoxmean_False BoxBoxmean = false
)

// BoxBoxpoints If *outliers*, only the sample points lying outside the whiskers are shown If *suspectedoutliers*, the outlier points are shown and points either less than 4*Q1-3*Q3 or greater than 4*Q3-3*Q1 are highlighted (see `outliercolor`) If *all*, all sample points are shown If *false*, only the box(es) are shown with no sample points Defaults to *suspectedoutliers* when `marker.outliercolor` or `marker.line.outliercolor` is set. Defaults to *all* under the q1/median/q3 signature. Otherwise defaults to *outliers*.
type BoxBoxpoints interface{}

var (
	BoxBoxpoints_all               BoxBoxpoints = "all"
	BoxBoxpoints_outliers          BoxBoxpoints = "outliers"
	BoxBoxpoints_suspectedoutliers BoxBoxpoints = "suspectedoutliers"
	BoxBoxpoints_False             BoxBoxpoints = false
)

// BoxHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type BoxHoverlabelAlign string

const (
	BoxHoverlabelAlign_left  BoxHoverlabelAlign = "left"
	BoxHoverlabelAlign_right BoxHoverlabelAlign = "right"
	BoxHoverlabelAlign_auto  BoxHoverlabelAlign = "auto"
)

// BoxMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type BoxMarkerSymbol string

const (
	BoxMarkerSymbol0                            BoxMarkerSymbol = "0"
	BoxMarkerSymbol_circle                      BoxMarkerSymbol = "circle"
	BoxMarkerSymbol100                          BoxMarkerSymbol = "100"
	BoxMarkerSymbol_circle_open                 BoxMarkerSymbol = "circle-open"
	BoxMarkerSymbol200                          BoxMarkerSymbol = "200"
	BoxMarkerSymbol_circle_dot                  BoxMarkerSymbol = "circle-dot"
	BoxMarkerSymbol300                          BoxMarkerSymbol = "300"
	BoxMarkerSymbol_circle_open_dot             BoxMarkerSymbol = "circle-open-dot"
	BoxMarkerSymbol1                            BoxMarkerSymbol = "1"
	BoxMarkerSymbol_square                      BoxMarkerSymbol = "square"
	BoxMarkerSymbol101                          BoxMarkerSymbol = "101"
	BoxMarkerSymbol_square_open                 BoxMarkerSymbol = "square-open"
	BoxMarkerSymbol201                          BoxMarkerSymbol = "201"
	BoxMarkerSymbol_square_dot                  BoxMarkerSymbol = "square-dot"
	BoxMarkerSymbol301                          BoxMarkerSymbol = "301"
	BoxMarkerSymbol_square_open_dot             BoxMarkerSymbol = "square-open-dot"
	BoxMarkerSymbol2                            BoxMarkerSymbol = "2"
	BoxMarkerSymbol_diamond                     BoxMarkerSymbol = "diamond"
	BoxMarkerSymbol102                          BoxMarkerSymbol = "102"
	BoxMarkerSymbol_diamond_open                BoxMarkerSymbol = "diamond-open"
	BoxMarkerSymbol202                          BoxMarkerSymbol = "202"
	BoxMarkerSymbol_diamond_dot                 BoxMarkerSymbol = "diamond-dot"
	BoxMarkerSymbol302                          BoxMarkerSymbol = "302"
	BoxMarkerSymbol_diamond_open_dot            BoxMarkerSymbol = "diamond-open-dot"
	BoxMarkerSymbol3                            BoxMarkerSymbol = "3"
	BoxMarkerSymbol_cross                       BoxMarkerSymbol = "cross"
	BoxMarkerSymbol103                          BoxMarkerSymbol = "103"
	BoxMarkerSymbol_cross_open                  BoxMarkerSymbol = "cross-open"
	BoxMarkerSymbol203                          BoxMarkerSymbol = "203"
	BoxMarkerSymbol_cross_dot                   BoxMarkerSymbol = "cross-dot"
	BoxMarkerSymbol303                          BoxMarkerSymbol = "303"
	BoxMarkerSymbol_cross_open_dot              BoxMarkerSymbol = "cross-open-dot"
	BoxMarkerSymbol4                            BoxMarkerSymbol = "4"
	BoxMarkerSymbol_x                           BoxMarkerSymbol = "x"
	BoxMarkerSymbol104                          BoxMarkerSymbol = "104"
	BoxMarkerSymbol_x_open                      BoxMarkerSymbol = "x-open"
	BoxMarkerSymbol204                          BoxMarkerSymbol = "204"
	BoxMarkerSymbol_x_dot                       BoxMarkerSymbol = "x-dot"
	BoxMarkerSymbol304                          BoxMarkerSymbol = "304"
	BoxMarkerSymbol_x_open_dot                  BoxMarkerSymbol = "x-open-dot"
	BoxMarkerSymbol5                            BoxMarkerSymbol = "5"
	BoxMarkerSymbol_triangle_up                 BoxMarkerSymbol = "triangle-up"
	BoxMarkerSymbol105                          BoxMarkerSymbol = "105"
	BoxMarkerSymbol_triangle_up_open            BoxMarkerSymbol = "triangle-up-open"
	BoxMarkerSymbol205                          BoxMarkerSymbol = "205"
	BoxMarkerSymbol_triangle_up_dot             BoxMarkerSymbol = "triangle-up-dot"
	BoxMarkerSymbol305                          BoxMarkerSymbol = "305"
	BoxMarkerSymbol_triangle_up_open_dot        BoxMarkerSymbol = "triangle-up-open-dot"
	BoxMarkerSymbol6                            BoxMarkerSymbol = "6"
	BoxMarkerSymbol_triangle_down               BoxMarkerSymbol = "triangle-down"
	BoxMarkerSymbol106                          BoxMarkerSymbol = "106"
	BoxMarkerSymbol_triangle_down_open          BoxMarkerSymbol = "triangle-down-open"
	BoxMarkerSymbol206                          BoxMarkerSymbol = "206"
	BoxMarkerSymbol_triangle_down_dot           BoxMarkerSymbol = "triangle-down-dot"
	BoxMarkerSymbol306                          BoxMarkerSymbol = "306"
	BoxMarkerSymbol_triangle_down_open_dot      BoxMarkerSymbol = "triangle-down-open-dot"
	BoxMarkerSymbol7                            BoxMarkerSymbol = "7"
	BoxMarkerSymbol_triangle_left               BoxMarkerSymbol = "triangle-left"
	BoxMarkerSymbol107                          BoxMarkerSymbol = "107"
	BoxMarkerSymbol_triangle_left_open          BoxMarkerSymbol = "triangle-left-open"
	BoxMarkerSymbol207                          BoxMarkerSymbol = "207"
	BoxMarkerSymbol_triangle_left_dot           BoxMarkerSymbol = "triangle-left-dot"
	BoxMarkerSymbol307                          BoxMarkerSymbol = "307"
	BoxMarkerSymbol_triangle_left_open_dot      BoxMarkerSymbol = "triangle-left-open-dot"
	BoxMarkerSymbol8                            BoxMarkerSymbol = "8"
	BoxMarkerSymbol_triangle_right              BoxMarkerSymbol = "triangle-right"
	BoxMarkerSymbol108                          BoxMarkerSymbol = "108"
	BoxMarkerSymbol_triangle_right_open         BoxMarkerSymbol = "triangle-right-open"
	BoxMarkerSymbol208                          BoxMarkerSymbol = "208"
	BoxMarkerSymbol_triangle_right_dot          BoxMarkerSymbol = "triangle-right-dot"
	BoxMarkerSymbol308                          BoxMarkerSymbol = "308"
	BoxMarkerSymbol_triangle_right_open_dot     BoxMarkerSymbol = "triangle-right-open-dot"
	BoxMarkerSymbol9                            BoxMarkerSymbol = "9"
	BoxMarkerSymbol_triangle_ne                 BoxMarkerSymbol = "triangle-ne"
	BoxMarkerSymbol109                          BoxMarkerSymbol = "109"
	BoxMarkerSymbol_triangle_ne_open            BoxMarkerSymbol = "triangle-ne-open"
	BoxMarkerSymbol209                          BoxMarkerSymbol = "209"
	BoxMarkerSymbol_triangle_ne_dot             BoxMarkerSymbol = "triangle-ne-dot"
	BoxMarkerSymbol309                          BoxMarkerSymbol = "309"
	BoxMarkerSymbol_triangle_ne_open_dot        BoxMarkerSymbol = "triangle-ne-open-dot"
	BoxMarkerSymbol10                           BoxMarkerSymbol = "10"
	BoxMarkerSymbol_triangle_se                 BoxMarkerSymbol = "triangle-se"
	BoxMarkerSymbol110                          BoxMarkerSymbol = "110"
	BoxMarkerSymbol_triangle_se_open            BoxMarkerSymbol = "triangle-se-open"
	BoxMarkerSymbol210                          BoxMarkerSymbol = "210"
	BoxMarkerSymbol_triangle_se_dot             BoxMarkerSymbol = "triangle-se-dot"
	BoxMarkerSymbol310                          BoxMarkerSymbol = "310"
	BoxMarkerSymbol_triangle_se_open_dot        BoxMarkerSymbol = "triangle-se-open-dot"
	BoxMarkerSymbol11                           BoxMarkerSymbol = "11"
	BoxMarkerSymbol_triangle_sw                 BoxMarkerSymbol = "triangle-sw"
	BoxMarkerSymbol111                          BoxMarkerSymbol = "111"
	BoxMarkerSymbol_triangle_sw_open            BoxMarkerSymbol = "triangle-sw-open"
	BoxMarkerSymbol211                          BoxMarkerSymbol = "211"
	BoxMarkerSymbol_triangle_sw_dot             BoxMarkerSymbol = "triangle-sw-dot"
	BoxMarkerSymbol311                          BoxMarkerSymbol = "311"
	BoxMarkerSymbol_triangle_sw_open_dot        BoxMarkerSymbol = "triangle-sw-open-dot"
	BoxMarkerSymbol12                           BoxMarkerSymbol = "12"
	BoxMarkerSymbol_triangle_nw                 BoxMarkerSymbol = "triangle-nw"
	BoxMarkerSymbol112                          BoxMarkerSymbol = "112"
	BoxMarkerSymbol_triangle_nw_open            BoxMarkerSymbol = "triangle-nw-open"
	BoxMarkerSymbol212                          BoxMarkerSymbol = "212"
	BoxMarkerSymbol_triangle_nw_dot             BoxMarkerSymbol = "triangle-nw-dot"
	BoxMarkerSymbol312                          BoxMarkerSymbol = "312"
	BoxMarkerSymbol_triangle_nw_open_dot        BoxMarkerSymbol = "triangle-nw-open-dot"
	BoxMarkerSymbol13                           BoxMarkerSymbol = "13"
	BoxMarkerSymbol_pentagon                    BoxMarkerSymbol = "pentagon"
	BoxMarkerSymbol113                          BoxMarkerSymbol = "113"
	BoxMarkerSymbol_pentagon_open               BoxMarkerSymbol = "pentagon-open"
	BoxMarkerSymbol213                          BoxMarkerSymbol = "213"
	BoxMarkerSymbol_pentagon_dot                BoxMarkerSymbol = "pentagon-dot"
	BoxMarkerSymbol313                          BoxMarkerSymbol = "313"
	BoxMarkerSymbol_pentagon_open_dot           BoxMarkerSymbol = "pentagon-open-dot"
	BoxMarkerSymbol14                           BoxMarkerSymbol = "14"
	BoxMarkerSymbol_hexagon                     BoxMarkerSymbol = "hexagon"
	BoxMarkerSymbol114                          BoxMarkerSymbol = "114"
	BoxMarkerSymbol_hexagon_open                BoxMarkerSymbol = "hexagon-open"
	BoxMarkerSymbol214                          BoxMarkerSymbol = "214"
	BoxMarkerSymbol_hexagon_dot                 BoxMarkerSymbol = "hexagon-dot"
	BoxMarkerSymbol314                          BoxMarkerSymbol = "314"
	BoxMarkerSymbol_hexagon_open_dot            BoxMarkerSymbol = "hexagon-open-dot"
	BoxMarkerSymbol15                           BoxMarkerSymbol = "15"
	BoxMarkerSymbol_hexagon2                    BoxMarkerSymbol = "hexagon2"
	BoxMarkerSymbol115                          BoxMarkerSymbol = "115"
	BoxMarkerSymbol_hexagon2_open               BoxMarkerSymbol = "hexagon2-open"
	BoxMarkerSymbol215                          BoxMarkerSymbol = "215"
	BoxMarkerSymbol_hexagon2_dot                BoxMarkerSymbol = "hexagon2-dot"
	BoxMarkerSymbol315                          BoxMarkerSymbol = "315"
	BoxMarkerSymbol_hexagon2_open_dot           BoxMarkerSymbol = "hexagon2-open-dot"
	BoxMarkerSymbol16                           BoxMarkerSymbol = "16"
	BoxMarkerSymbol_octagon                     BoxMarkerSymbol = "octagon"
	BoxMarkerSymbol116                          BoxMarkerSymbol = "116"
	BoxMarkerSymbol_octagon_open                BoxMarkerSymbol = "octagon-open"
	BoxMarkerSymbol216                          BoxMarkerSymbol = "216"
	BoxMarkerSymbol_octagon_dot                 BoxMarkerSymbol = "octagon-dot"
	BoxMarkerSymbol316                          BoxMarkerSymbol = "316"
	BoxMarkerSymbol_octagon_open_dot            BoxMarkerSymbol = "octagon-open-dot"
	BoxMarkerSymbol17                           BoxMarkerSymbol = "17"
	BoxMarkerSymbol_star                        BoxMarkerSymbol = "star"
	BoxMarkerSymbol117                          BoxMarkerSymbol = "117"
	BoxMarkerSymbol_star_open                   BoxMarkerSymbol = "star-open"
	BoxMarkerSymbol217                          BoxMarkerSymbol = "217"
	BoxMarkerSymbol_star_dot                    BoxMarkerSymbol = "star-dot"
	BoxMarkerSymbol317                          BoxMarkerSymbol = "317"
	BoxMarkerSymbol_star_open_dot               BoxMarkerSymbol = "star-open-dot"
	BoxMarkerSymbol18                           BoxMarkerSymbol = "18"
	BoxMarkerSymbol_hexagram                    BoxMarkerSymbol = "hexagram"
	BoxMarkerSymbol118                          BoxMarkerSymbol = "118"
	BoxMarkerSymbol_hexagram_open               BoxMarkerSymbol = "hexagram-open"
	BoxMarkerSymbol218                          BoxMarkerSymbol = "218"
	BoxMarkerSymbol_hexagram_dot                BoxMarkerSymbol = "hexagram-dot"
	BoxMarkerSymbol318                          BoxMarkerSymbol = "318"
	BoxMarkerSymbol_hexagram_open_dot           BoxMarkerSymbol = "hexagram-open-dot"
	BoxMarkerSymbol19                           BoxMarkerSymbol = "19"
	BoxMarkerSymbol_star_triangle_up            BoxMarkerSymbol = "star-triangle-up"
	BoxMarkerSymbol119                          BoxMarkerSymbol = "119"
	BoxMarkerSymbol_star_triangle_up_open       BoxMarkerSymbol = "star-triangle-up-open"
	BoxMarkerSymbol219                          BoxMarkerSymbol = "219"
	BoxMarkerSymbol_star_triangle_up_dot        BoxMarkerSymbol = "star-triangle-up-dot"
	BoxMarkerSymbol319                          BoxMarkerSymbol = "319"
	BoxMarkerSymbol_star_triangle_up_open_dot   BoxMarkerSymbol = "star-triangle-up-open-dot"
	BoxMarkerSymbol20                           BoxMarkerSymbol = "20"
	BoxMarkerSymbol_star_triangle_down          BoxMarkerSymbol = "star-triangle-down"
	BoxMarkerSymbol120                          BoxMarkerSymbol = "120"
	BoxMarkerSymbol_star_triangle_down_open     BoxMarkerSymbol = "star-triangle-down-open"
	BoxMarkerSymbol220                          BoxMarkerSymbol = "220"
	BoxMarkerSymbol_star_triangle_down_dot      BoxMarkerSymbol = "star-triangle-down-dot"
	BoxMarkerSymbol320                          BoxMarkerSymbol = "320"
	BoxMarkerSymbol_star_triangle_down_open_dot BoxMarkerSymbol = "star-triangle-down-open-dot"
	BoxMarkerSymbol21                           BoxMarkerSymbol = "21"
	BoxMarkerSymbol_star_square                 BoxMarkerSymbol = "star-square"
	BoxMarkerSymbol121                          BoxMarkerSymbol = "121"
	BoxMarkerSymbol_star_square_open            BoxMarkerSymbol = "star-square-open"
	BoxMarkerSymbol221                          BoxMarkerSymbol = "221"
	BoxMarkerSymbol_star_square_dot             BoxMarkerSymbol = "star-square-dot"
	BoxMarkerSymbol321                          BoxMarkerSymbol = "321"
	BoxMarkerSymbol_star_square_open_dot        BoxMarkerSymbol = "star-square-open-dot"
	BoxMarkerSymbol22                           BoxMarkerSymbol = "22"
	BoxMarkerSymbol_star_diamond                BoxMarkerSymbol = "star-diamond"
	BoxMarkerSymbol122                          BoxMarkerSymbol = "122"
	BoxMarkerSymbol_star_diamond_open           BoxMarkerSymbol = "star-diamond-open"
	BoxMarkerSymbol222                          BoxMarkerSymbol = "222"
	BoxMarkerSymbol_star_diamond_dot            BoxMarkerSymbol = "star-diamond-dot"
	BoxMarkerSymbol322                          BoxMarkerSymbol = "322"
	BoxMarkerSymbol_star_diamond_open_dot       BoxMarkerSymbol = "star-diamond-open-dot"
	BoxMarkerSymbol23                           BoxMarkerSymbol = "23"
	BoxMarkerSymbol_diamond_tall                BoxMarkerSymbol = "diamond-tall"
	BoxMarkerSymbol123                          BoxMarkerSymbol = "123"
	BoxMarkerSymbol_diamond_tall_open           BoxMarkerSymbol = "diamond-tall-open"
	BoxMarkerSymbol223                          BoxMarkerSymbol = "223"
	BoxMarkerSymbol_diamond_tall_dot            BoxMarkerSymbol = "diamond-tall-dot"
	BoxMarkerSymbol323                          BoxMarkerSymbol = "323"
	BoxMarkerSymbol_diamond_tall_open_dot       BoxMarkerSymbol = "diamond-tall-open-dot"
	BoxMarkerSymbol24                           BoxMarkerSymbol = "24"
	BoxMarkerSymbol_diamond_wide                BoxMarkerSymbol = "diamond-wide"
	BoxMarkerSymbol124                          BoxMarkerSymbol = "124"
	BoxMarkerSymbol_diamond_wide_open           BoxMarkerSymbol = "diamond-wide-open"
	BoxMarkerSymbol224                          BoxMarkerSymbol = "224"
	BoxMarkerSymbol_diamond_wide_dot            BoxMarkerSymbol = "diamond-wide-dot"
	BoxMarkerSymbol324                          BoxMarkerSymbol = "324"
	BoxMarkerSymbol_diamond_wide_open_dot       BoxMarkerSymbol = "diamond-wide-open-dot"
	BoxMarkerSymbol25                           BoxMarkerSymbol = "25"
	BoxMarkerSymbol_hourglass                   BoxMarkerSymbol = "hourglass"
	BoxMarkerSymbol125                          BoxMarkerSymbol = "125"
	BoxMarkerSymbol_hourglass_open              BoxMarkerSymbol = "hourglass-open"
	BoxMarkerSymbol26                           BoxMarkerSymbol = "26"
	BoxMarkerSymbol_bowtie                      BoxMarkerSymbol = "bowtie"
	BoxMarkerSymbol126                          BoxMarkerSymbol = "126"
	BoxMarkerSymbol_bowtie_open                 BoxMarkerSymbol = "bowtie-open"
	BoxMarkerSymbol27                           BoxMarkerSymbol = "27"
	BoxMarkerSymbol_circle_cross                BoxMarkerSymbol = "circle-cross"
	BoxMarkerSymbol127                          BoxMarkerSymbol = "127"
	BoxMarkerSymbol_circle_cross_open           BoxMarkerSymbol = "circle-cross-open"
	BoxMarkerSymbol28                           BoxMarkerSymbol = "28"
	BoxMarkerSymbol_circle_x                    BoxMarkerSymbol = "circle-x"
	BoxMarkerSymbol128                          BoxMarkerSymbol = "128"
	BoxMarkerSymbol_circle_x_open               BoxMarkerSymbol = "circle-x-open"
	BoxMarkerSymbol29                           BoxMarkerSymbol = "29"
	BoxMarkerSymbol_square_cross                BoxMarkerSymbol = "square-cross"
	BoxMarkerSymbol129                          BoxMarkerSymbol = "129"
	BoxMarkerSymbol_square_cross_open           BoxMarkerSymbol = "square-cross-open"
	BoxMarkerSymbol30                           BoxMarkerSymbol = "30"
	BoxMarkerSymbol_square_x                    BoxMarkerSymbol = "square-x"
	BoxMarkerSymbol130                          BoxMarkerSymbol = "130"
	BoxMarkerSymbol_square_x_open               BoxMarkerSymbol = "square-x-open"
	BoxMarkerSymbol31                           BoxMarkerSymbol = "31"
	BoxMarkerSymbol_diamond_cross               BoxMarkerSymbol = "diamond-cross"
	BoxMarkerSymbol131                          BoxMarkerSymbol = "131"
	BoxMarkerSymbol_diamond_cross_open          BoxMarkerSymbol = "diamond-cross-open"
	BoxMarkerSymbol32                           BoxMarkerSymbol = "32"
	BoxMarkerSymbol_diamond_x                   BoxMarkerSymbol = "diamond-x"
	BoxMarkerSymbol132                          BoxMarkerSymbol = "132"
	BoxMarkerSymbol_diamond_x_open              BoxMarkerSymbol = "diamond-x-open"
	BoxMarkerSymbol33                           BoxMarkerSymbol = "33"
	BoxMarkerSymbol_cross_thin                  BoxMarkerSymbol = "cross-thin"
	BoxMarkerSymbol133                          BoxMarkerSymbol = "133"
	BoxMarkerSymbol_cross_thin_open             BoxMarkerSymbol = "cross-thin-open"
	BoxMarkerSymbol34                           BoxMarkerSymbol = "34"
	BoxMarkerSymbol_x_thin                      BoxMarkerSymbol = "x-thin"
	BoxMarkerSymbol134                          BoxMarkerSymbol = "134"
	BoxMarkerSymbol_x_thin_open                 BoxMarkerSymbol = "x-thin-open"
	BoxMarkerSymbol35                           BoxMarkerSymbol = "35"
	BoxMarkerSymbol_asterisk                    BoxMarkerSymbol = "asterisk"
	BoxMarkerSymbol135                          BoxMarkerSymbol = "135"
	BoxMarkerSymbol_asterisk_open               BoxMarkerSymbol = "asterisk-open"
	BoxMarkerSymbol36                           BoxMarkerSymbol = "36"
	BoxMarkerSymbol_hash                        BoxMarkerSymbol = "hash"
	BoxMarkerSymbol136                          BoxMarkerSymbol = "136"
	BoxMarkerSymbol_hash_open                   BoxMarkerSymbol = "hash-open"
	BoxMarkerSymbol236                          BoxMarkerSymbol = "236"
	BoxMarkerSymbol_hash_dot                    BoxMarkerSymbol = "hash-dot"
	BoxMarkerSymbol336                          BoxMarkerSymbol = "336"
	BoxMarkerSymbol_hash_open_dot               BoxMarkerSymbol = "hash-open-dot"
	BoxMarkerSymbol37                           BoxMarkerSymbol = "37"
	BoxMarkerSymbol_y_up                        BoxMarkerSymbol = "y-up"
	BoxMarkerSymbol137                          BoxMarkerSymbol = "137"
	BoxMarkerSymbol_y_up_open                   BoxMarkerSymbol = "y-up-open"
	BoxMarkerSymbol38                           BoxMarkerSymbol = "38"
	BoxMarkerSymbol_y_down                      BoxMarkerSymbol = "y-down"
	BoxMarkerSymbol138                          BoxMarkerSymbol = "138"
	BoxMarkerSymbol_y_down_open                 BoxMarkerSymbol = "y-down-open"
	BoxMarkerSymbol39                           BoxMarkerSymbol = "39"
	BoxMarkerSymbol_y_left                      BoxMarkerSymbol = "y-left"
	BoxMarkerSymbol139                          BoxMarkerSymbol = "139"
	BoxMarkerSymbol_y_left_open                 BoxMarkerSymbol = "y-left-open"
	BoxMarkerSymbol40                           BoxMarkerSymbol = "40"
	BoxMarkerSymbol_y_right                     BoxMarkerSymbol = "y-right"
	BoxMarkerSymbol140                          BoxMarkerSymbol = "140"
	BoxMarkerSymbol_y_right_open                BoxMarkerSymbol = "y-right-open"
	BoxMarkerSymbol41                           BoxMarkerSymbol = "41"
	BoxMarkerSymbol_line_ew                     BoxMarkerSymbol = "line-ew"
	BoxMarkerSymbol141                          BoxMarkerSymbol = "141"
	BoxMarkerSymbol_line_ew_open                BoxMarkerSymbol = "line-ew-open"
	BoxMarkerSymbol42                           BoxMarkerSymbol = "42"
	BoxMarkerSymbol_line_ns                     BoxMarkerSymbol = "line-ns"
	BoxMarkerSymbol142                          BoxMarkerSymbol = "142"
	BoxMarkerSymbol_line_ns_open                BoxMarkerSymbol = "line-ns-open"
	BoxMarkerSymbol43                           BoxMarkerSymbol = "43"
	BoxMarkerSymbol_line_ne                     BoxMarkerSymbol = "line-ne"
	BoxMarkerSymbol143                          BoxMarkerSymbol = "143"
	BoxMarkerSymbol_line_ne_open                BoxMarkerSymbol = "line-ne-open"
	BoxMarkerSymbol44                           BoxMarkerSymbol = "44"
	BoxMarkerSymbol_line_nw                     BoxMarkerSymbol = "line-nw"
	BoxMarkerSymbol144                          BoxMarkerSymbol = "144"
	BoxMarkerSymbol_line_nw_open                BoxMarkerSymbol = "line-nw-open"
)

// BoxOrientation Sets the orientation of the box(es). If *v* (*h*), the distribution is visualized along the vertical (horizontal).
type BoxOrientation string

const (
	BoxOrientation_v BoxOrientation = "v"
	BoxOrientation_h BoxOrientation = "h"
)

// BoxQuartilemethod Sets the method used to compute the sample's Q1 and Q3 quartiles. The *linear* method uses the 25th percentile for Q1 and 75th percentile for Q3 as computed using method #10 (listed on http://www.amstat.org/publications/jse/v14n3/langford.html). The *exclusive* method uses the median to divide the ordered dataset into two halves if the sample is odd, it does not include the median in either half - Q1 is then the median of the lower half and Q3 the median of the upper half. The *inclusive* method also uses the median to divide the ordered dataset into two halves but if the sample is odd, it includes the median in both halves - Q1 is then the median of the lower half and Q3 the median of the upper half.
type BoxQuartilemethod string

const (
	BoxQuartilemethod_linear    BoxQuartilemethod = "linear"
	BoxQuartilemethod_exclusive BoxQuartilemethod = "exclusive"
	BoxQuartilemethod_inclusive BoxQuartilemethod = "inclusive"
)

// BoxVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type BoxVisible interface{}

var (
	BoxVisible_True       BoxVisible = true
	BoxVisible_False      BoxVisible = false
	BoxVisible_legendonly BoxVisible = "legendonly"
)

// BoxXcalendar Sets the calendar system to use with `x` date data.
type BoxXcalendar string

const (
	BoxXcalendar_gregorian  BoxXcalendar = "gregorian"
	BoxXcalendar_chinese    BoxXcalendar = "chinese"
	BoxXcalendar_coptic     BoxXcalendar = "coptic"
	BoxXcalendar_discworld  BoxXcalendar = "discworld"
	BoxXcalendar_ethiopian  BoxXcalendar = "ethiopian"
	BoxXcalendar_hebrew     BoxXcalendar = "hebrew"
	BoxXcalendar_islamic    BoxXcalendar = "islamic"
	BoxXcalendar_julian     BoxXcalendar = "julian"
	BoxXcalendar_mayan      BoxXcalendar = "mayan"
	BoxXcalendar_nanakshahi BoxXcalendar = "nanakshahi"
	BoxXcalendar_nepali     BoxXcalendar = "nepali"
	BoxXcalendar_persian    BoxXcalendar = "persian"
	BoxXcalendar_jalali     BoxXcalendar = "jalali"
	BoxXcalendar_taiwan     BoxXcalendar = "taiwan"
	BoxXcalendar_thai       BoxXcalendar = "thai"
	BoxXcalendar_ummalqura  BoxXcalendar = "ummalqura"
)

// BoxYcalendar Sets the calendar system to use with `y` date data.
type BoxYcalendar string

const (
	BoxYcalendar_gregorian  BoxYcalendar = "gregorian"
	BoxYcalendar_chinese    BoxYcalendar = "chinese"
	BoxYcalendar_coptic     BoxYcalendar = "coptic"
	BoxYcalendar_discworld  BoxYcalendar = "discworld"
	BoxYcalendar_ethiopian  BoxYcalendar = "ethiopian"
	BoxYcalendar_hebrew     BoxYcalendar = "hebrew"
	BoxYcalendar_islamic    BoxYcalendar = "islamic"
	BoxYcalendar_julian     BoxYcalendar = "julian"
	BoxYcalendar_mayan      BoxYcalendar = "mayan"
	BoxYcalendar_nanakshahi BoxYcalendar = "nanakshahi"
	BoxYcalendar_nepali     BoxYcalendar = "nepali"
	BoxYcalendar_persian    BoxYcalendar = "persian"
	BoxYcalendar_jalali     BoxYcalendar = "jalali"
	BoxYcalendar_taiwan     BoxYcalendar = "taiwan"
	BoxYcalendar_thai       BoxYcalendar = "thai"
	BoxYcalendar_ummalqura  BoxYcalendar = "ummalqura"
)

// CandlestickHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type CandlestickHoverlabelAlign string

const (
	CandlestickHoverlabelAlign_left  CandlestickHoverlabelAlign = "left"
	CandlestickHoverlabelAlign_right CandlestickHoverlabelAlign = "right"
	CandlestickHoverlabelAlign_auto  CandlestickHoverlabelAlign = "auto"
)

// CandlestickVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type CandlestickVisible interface{}

var (
	CandlestickVisible_True       CandlestickVisible = true
	CandlestickVisible_False      CandlestickVisible = false
	CandlestickVisible_legendonly CandlestickVisible = "legendonly"
)

// CandlestickXcalendar Sets the calendar system to use with `x` date data.
type CandlestickXcalendar string

const (
	CandlestickXcalendar_gregorian  CandlestickXcalendar = "gregorian"
	CandlestickXcalendar_chinese    CandlestickXcalendar = "chinese"
	CandlestickXcalendar_coptic     CandlestickXcalendar = "coptic"
	CandlestickXcalendar_discworld  CandlestickXcalendar = "discworld"
	CandlestickXcalendar_ethiopian  CandlestickXcalendar = "ethiopian"
	CandlestickXcalendar_hebrew     CandlestickXcalendar = "hebrew"
	CandlestickXcalendar_islamic    CandlestickXcalendar = "islamic"
	CandlestickXcalendar_julian     CandlestickXcalendar = "julian"
	CandlestickXcalendar_mayan      CandlestickXcalendar = "mayan"
	CandlestickXcalendar_nanakshahi CandlestickXcalendar = "nanakshahi"
	CandlestickXcalendar_nepali     CandlestickXcalendar = "nepali"
	CandlestickXcalendar_persian    CandlestickXcalendar = "persian"
	CandlestickXcalendar_jalali     CandlestickXcalendar = "jalali"
	CandlestickXcalendar_taiwan     CandlestickXcalendar = "taiwan"
	CandlestickXcalendar_thai       CandlestickXcalendar = "thai"
	CandlestickXcalendar_ummalqura  CandlestickXcalendar = "ummalqura"
)

// CarpetAaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type CarpetAaxisAutorange interface{}

var (
	CarpetAaxisAutorange_True     CarpetAaxisAutorange = true
	CarpetAaxisAutorange_False    CarpetAaxisAutorange = false
	CarpetAaxisAutorange_reversed CarpetAaxisAutorange = "reversed"
)

// CarpetAaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`.
type CarpetAaxisCategoryorder string

const (
	CarpetAaxisCategoryorder_trace              CarpetAaxisCategoryorder = "trace"
	CarpetAaxisCategoryorder_categoryascending  CarpetAaxisCategoryorder = "category ascending"
	CarpetAaxisCategoryorder_categorydescending CarpetAaxisCategoryorder = "category descending"
	CarpetAaxisCategoryorder_array              CarpetAaxisCategoryorder = "array"
)

// CarpetAaxisCheatertype <no value>
type CarpetAaxisCheatertype string

const (
	CarpetAaxisCheatertype_index CarpetAaxisCheatertype = "index"
	CarpetAaxisCheatertype_value CarpetAaxisCheatertype = "value"
)

// CarpetAaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type CarpetAaxisExponentformat string

const (
	CarpetAaxisExponentformat_none  CarpetAaxisExponentformat = "none"
	CarpetAaxisExponentformat_e     CarpetAaxisExponentformat = "e"
	CarpetAaxisExponentformat_E     CarpetAaxisExponentformat = "E"
	CarpetAaxisExponentformat_power CarpetAaxisExponentformat = "power"
	CarpetAaxisExponentformat_SI    CarpetAaxisExponentformat = "SI"
	CarpetAaxisExponentformat_B     CarpetAaxisExponentformat = "B"
)

// CarpetAaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data.
type CarpetAaxisRangemode string

const (
	CarpetAaxisRangemode_normal      CarpetAaxisRangemode = "normal"
	CarpetAaxisRangemode_tozero      CarpetAaxisRangemode = "tozero"
	CarpetAaxisRangemode_nonnegative CarpetAaxisRangemode = "nonnegative"
)

// CarpetAaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type CarpetAaxisShowexponent string

const (
	CarpetAaxisShowexponent_all   CarpetAaxisShowexponent = "all"
	CarpetAaxisShowexponent_first CarpetAaxisShowexponent = "first"
	CarpetAaxisShowexponent_last  CarpetAaxisShowexponent = "last"
	CarpetAaxisShowexponent_none  CarpetAaxisShowexponent = "none"
)

// CarpetAaxisShowticklabels Determines whether axis labels are drawn on the low side, the high side, both, or neither side of the axis.
type CarpetAaxisShowticklabels string

const (
	CarpetAaxisShowticklabels_start CarpetAaxisShowticklabels = "start"
	CarpetAaxisShowticklabels_end   CarpetAaxisShowticklabels = "end"
	CarpetAaxisShowticklabels_both  CarpetAaxisShowticklabels = "both"
	CarpetAaxisShowticklabels_none  CarpetAaxisShowticklabels = "none"
)

// CarpetAaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type CarpetAaxisShowtickprefix string

const (
	CarpetAaxisShowtickprefix_all   CarpetAaxisShowtickprefix = "all"
	CarpetAaxisShowtickprefix_first CarpetAaxisShowtickprefix = "first"
	CarpetAaxisShowtickprefix_last  CarpetAaxisShowtickprefix = "last"
	CarpetAaxisShowtickprefix_none  CarpetAaxisShowtickprefix = "none"
)

// CarpetAaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type CarpetAaxisShowticksuffix string

const (
	CarpetAaxisShowticksuffix_all   CarpetAaxisShowticksuffix = "all"
	CarpetAaxisShowticksuffix_first CarpetAaxisShowticksuffix = "first"
	CarpetAaxisShowticksuffix_last  CarpetAaxisShowticksuffix = "last"
	CarpetAaxisShowticksuffix_none  CarpetAaxisShowticksuffix = "none"
)

// CarpetAaxisTickmode <no value>
type CarpetAaxisTickmode string

const (
	CarpetAaxisTickmode_linear CarpetAaxisTickmode = "linear"
	CarpetAaxisTickmode_array  CarpetAaxisTickmode = "array"
)

// CarpetAaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type CarpetAaxisType string

const (
	CarpetAaxisType__        CarpetAaxisType = "-"
	CarpetAaxisType_linear   CarpetAaxisType = "linear"
	CarpetAaxisType_date     CarpetAaxisType = "date"
	CarpetAaxisType_category CarpetAaxisType = "category"
)

// CarpetBaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type CarpetBaxisAutorange interface{}

var (
	CarpetBaxisAutorange_True     CarpetBaxisAutorange = true
	CarpetBaxisAutorange_False    CarpetBaxisAutorange = false
	CarpetBaxisAutorange_reversed CarpetBaxisAutorange = "reversed"
)

// CarpetBaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`.
type CarpetBaxisCategoryorder string

const (
	CarpetBaxisCategoryorder_trace              CarpetBaxisCategoryorder = "trace"
	CarpetBaxisCategoryorder_categoryascending  CarpetBaxisCategoryorder = "category ascending"
	CarpetBaxisCategoryorder_categorydescending CarpetBaxisCategoryorder = "category descending"
	CarpetBaxisCategoryorder_array              CarpetBaxisCategoryorder = "array"
)

// CarpetBaxisCheatertype <no value>
type CarpetBaxisCheatertype string

const (
	CarpetBaxisCheatertype_index CarpetBaxisCheatertype = "index"
	CarpetBaxisCheatertype_value CarpetBaxisCheatertype = "value"
)

// CarpetBaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type CarpetBaxisExponentformat string

const (
	CarpetBaxisExponentformat_none  CarpetBaxisExponentformat = "none"
	CarpetBaxisExponentformat_e     CarpetBaxisExponentformat = "e"
	CarpetBaxisExponentformat_E     CarpetBaxisExponentformat = "E"
	CarpetBaxisExponentformat_power CarpetBaxisExponentformat = "power"
	CarpetBaxisExponentformat_SI    CarpetBaxisExponentformat = "SI"
	CarpetBaxisExponentformat_B     CarpetBaxisExponentformat = "B"
)

// CarpetBaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data.
type CarpetBaxisRangemode string

const (
	CarpetBaxisRangemode_normal      CarpetBaxisRangemode = "normal"
	CarpetBaxisRangemode_tozero      CarpetBaxisRangemode = "tozero"
	CarpetBaxisRangemode_nonnegative CarpetBaxisRangemode = "nonnegative"
)

// CarpetBaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type CarpetBaxisShowexponent string

const (
	CarpetBaxisShowexponent_all   CarpetBaxisShowexponent = "all"
	CarpetBaxisShowexponent_first CarpetBaxisShowexponent = "first"
	CarpetBaxisShowexponent_last  CarpetBaxisShowexponent = "last"
	CarpetBaxisShowexponent_none  CarpetBaxisShowexponent = "none"
)

// CarpetBaxisShowticklabels Determines whether axis labels are drawn on the low side, the high side, both, or neither side of the axis.
type CarpetBaxisShowticklabels string

const (
	CarpetBaxisShowticklabels_start CarpetBaxisShowticklabels = "start"
	CarpetBaxisShowticklabels_end   CarpetBaxisShowticklabels = "end"
	CarpetBaxisShowticklabels_both  CarpetBaxisShowticklabels = "both"
	CarpetBaxisShowticklabels_none  CarpetBaxisShowticklabels = "none"
)

// CarpetBaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type CarpetBaxisShowtickprefix string

const (
	CarpetBaxisShowtickprefix_all   CarpetBaxisShowtickprefix = "all"
	CarpetBaxisShowtickprefix_first CarpetBaxisShowtickprefix = "first"
	CarpetBaxisShowtickprefix_last  CarpetBaxisShowtickprefix = "last"
	CarpetBaxisShowtickprefix_none  CarpetBaxisShowtickprefix = "none"
)

// CarpetBaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type CarpetBaxisShowticksuffix string

const (
	CarpetBaxisShowticksuffix_all   CarpetBaxisShowticksuffix = "all"
	CarpetBaxisShowticksuffix_first CarpetBaxisShowticksuffix = "first"
	CarpetBaxisShowticksuffix_last  CarpetBaxisShowticksuffix = "last"
	CarpetBaxisShowticksuffix_none  CarpetBaxisShowticksuffix = "none"
)

// CarpetBaxisTickmode <no value>
type CarpetBaxisTickmode string

const (
	CarpetBaxisTickmode_linear CarpetBaxisTickmode = "linear"
	CarpetBaxisTickmode_array  CarpetBaxisTickmode = "array"
)

// CarpetBaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type CarpetBaxisType string

const (
	CarpetBaxisType__        CarpetBaxisType = "-"
	CarpetBaxisType_linear   CarpetBaxisType = "linear"
	CarpetBaxisType_date     CarpetBaxisType = "date"
	CarpetBaxisType_category CarpetBaxisType = "category"
)

// CarpetVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type CarpetVisible interface{}

var (
	CarpetVisible_True       CarpetVisible = true
	CarpetVisible_False      CarpetVisible = false
	CarpetVisible_legendonly CarpetVisible = "legendonly"
)

// ChoroplethColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ChoroplethColorbarExponentformat string

const (
	ChoroplethColorbarExponentformat_none  ChoroplethColorbarExponentformat = "none"
	ChoroplethColorbarExponentformat_e     ChoroplethColorbarExponentformat = "e"
	ChoroplethColorbarExponentformat_E     ChoroplethColorbarExponentformat = "E"
	ChoroplethColorbarExponentformat_power ChoroplethColorbarExponentformat = "power"
	ChoroplethColorbarExponentformat_SI    ChoroplethColorbarExponentformat = "SI"
	ChoroplethColorbarExponentformat_B     ChoroplethColorbarExponentformat = "B"
)

// ChoroplethColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ChoroplethColorbarLenmode string

const (
	ChoroplethColorbarLenmode_fraction ChoroplethColorbarLenmode = "fraction"
	ChoroplethColorbarLenmode_pixels   ChoroplethColorbarLenmode = "pixels"
)

// ChoroplethColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ChoroplethColorbarShowexponent string

const (
	ChoroplethColorbarShowexponent_all   ChoroplethColorbarShowexponent = "all"
	ChoroplethColorbarShowexponent_first ChoroplethColorbarShowexponent = "first"
	ChoroplethColorbarShowexponent_last  ChoroplethColorbarShowexponent = "last"
	ChoroplethColorbarShowexponent_none  ChoroplethColorbarShowexponent = "none"
)

// ChoroplethColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ChoroplethColorbarShowtickprefix string

const (
	ChoroplethColorbarShowtickprefix_all   ChoroplethColorbarShowtickprefix = "all"
	ChoroplethColorbarShowtickprefix_first ChoroplethColorbarShowtickprefix = "first"
	ChoroplethColorbarShowtickprefix_last  ChoroplethColorbarShowtickprefix = "last"
	ChoroplethColorbarShowtickprefix_none  ChoroplethColorbarShowtickprefix = "none"
)

// ChoroplethColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ChoroplethColorbarShowticksuffix string

const (
	ChoroplethColorbarShowticksuffix_all   ChoroplethColorbarShowticksuffix = "all"
	ChoroplethColorbarShowticksuffix_first ChoroplethColorbarShowticksuffix = "first"
	ChoroplethColorbarShowticksuffix_last  ChoroplethColorbarShowticksuffix = "last"
	ChoroplethColorbarShowticksuffix_none  ChoroplethColorbarShowticksuffix = "none"
)

// ChoroplethColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ChoroplethColorbarThicknessmode string

const (
	ChoroplethColorbarThicknessmode_fraction ChoroplethColorbarThicknessmode = "fraction"
	ChoroplethColorbarThicknessmode_pixels   ChoroplethColorbarThicknessmode = "pixels"
)

// ChoroplethColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ChoroplethColorbarTickmode string

const (
	ChoroplethColorbarTickmode_auto   ChoroplethColorbarTickmode = "auto"
	ChoroplethColorbarTickmode_linear ChoroplethColorbarTickmode = "linear"
	ChoroplethColorbarTickmode_array  ChoroplethColorbarTickmode = "array"
)

// ChoroplethColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ChoroplethColorbarTicks string

const (
	ChoroplethColorbarTicks_outside ChoroplethColorbarTicks = "outside"
	ChoroplethColorbarTicks_inside  ChoroplethColorbarTicks = "inside"
)

// ChoroplethColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ChoroplethColorbarTitleSide string

const (
	ChoroplethColorbarTitleSide_right  ChoroplethColorbarTitleSide = "right"
	ChoroplethColorbarTitleSide_top    ChoroplethColorbarTitleSide = "top"
	ChoroplethColorbarTitleSide_bottom ChoroplethColorbarTitleSide = "bottom"
)

// ChoroplethColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ChoroplethColorbarXanchor string

const (
	ChoroplethColorbarXanchor_left   ChoroplethColorbarXanchor = "left"
	ChoroplethColorbarXanchor_center ChoroplethColorbarXanchor = "center"
	ChoroplethColorbarXanchor_right  ChoroplethColorbarXanchor = "right"
)

// ChoroplethColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ChoroplethColorbarYanchor string

const (
	ChoroplethColorbarYanchor_top    ChoroplethColorbarYanchor = "top"
	ChoroplethColorbarYanchor_middle ChoroplethColorbarYanchor = "middle"
	ChoroplethColorbarYanchor_bottom ChoroplethColorbarYanchor = "bottom"
)

// ChoroplethHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ChoroplethHoverlabelAlign string

const (
	ChoroplethHoverlabelAlign_left  ChoroplethHoverlabelAlign = "left"
	ChoroplethHoverlabelAlign_right ChoroplethHoverlabelAlign = "right"
	ChoroplethHoverlabelAlign_auto  ChoroplethHoverlabelAlign = "auto"
)

// ChoroplethLocationmode Determines the set of locations used to match entries in `locations` to regions on the map. Values *ISO-3*, *USA-states*, *country names* correspond to features on the base map and value *geojson-id* corresponds to features from a custom GeoJSON linked to the `geojson` attribute.
type ChoroplethLocationmode string

const (
	ChoroplethLocationmode_ISO_3        ChoroplethLocationmode = "ISO-3"
	ChoroplethLocationmode_USA_states   ChoroplethLocationmode = "USA-states"
	ChoroplethLocationmode_countrynames ChoroplethLocationmode = "country names"
	ChoroplethLocationmode_geojson_id   ChoroplethLocationmode = "geojson-id"
)

// ChoroplethVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ChoroplethVisible interface{}

var (
	ChoroplethVisible_True       ChoroplethVisible = true
	ChoroplethVisible_False      ChoroplethVisible = false
	ChoroplethVisible_legendonly ChoroplethVisible = "legendonly"
)

// ChoroplethmapboxColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ChoroplethmapboxColorbarExponentformat string

const (
	ChoroplethmapboxColorbarExponentformat_none  ChoroplethmapboxColorbarExponentformat = "none"
	ChoroplethmapboxColorbarExponentformat_e     ChoroplethmapboxColorbarExponentformat = "e"
	ChoroplethmapboxColorbarExponentformat_E     ChoroplethmapboxColorbarExponentformat = "E"
	ChoroplethmapboxColorbarExponentformat_power ChoroplethmapboxColorbarExponentformat = "power"
	ChoroplethmapboxColorbarExponentformat_SI    ChoroplethmapboxColorbarExponentformat = "SI"
	ChoroplethmapboxColorbarExponentformat_B     ChoroplethmapboxColorbarExponentformat = "B"
)

// ChoroplethmapboxColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ChoroplethmapboxColorbarLenmode string

const (
	ChoroplethmapboxColorbarLenmode_fraction ChoroplethmapboxColorbarLenmode = "fraction"
	ChoroplethmapboxColorbarLenmode_pixels   ChoroplethmapboxColorbarLenmode = "pixels"
)

// ChoroplethmapboxColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ChoroplethmapboxColorbarShowexponent string

const (
	ChoroplethmapboxColorbarShowexponent_all   ChoroplethmapboxColorbarShowexponent = "all"
	ChoroplethmapboxColorbarShowexponent_first ChoroplethmapboxColorbarShowexponent = "first"
	ChoroplethmapboxColorbarShowexponent_last  ChoroplethmapboxColorbarShowexponent = "last"
	ChoroplethmapboxColorbarShowexponent_none  ChoroplethmapboxColorbarShowexponent = "none"
)

// ChoroplethmapboxColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ChoroplethmapboxColorbarShowtickprefix string

const (
	ChoroplethmapboxColorbarShowtickprefix_all   ChoroplethmapboxColorbarShowtickprefix = "all"
	ChoroplethmapboxColorbarShowtickprefix_first ChoroplethmapboxColorbarShowtickprefix = "first"
	ChoroplethmapboxColorbarShowtickprefix_last  ChoroplethmapboxColorbarShowtickprefix = "last"
	ChoroplethmapboxColorbarShowtickprefix_none  ChoroplethmapboxColorbarShowtickprefix = "none"
)

// ChoroplethmapboxColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ChoroplethmapboxColorbarShowticksuffix string

const (
	ChoroplethmapboxColorbarShowticksuffix_all   ChoroplethmapboxColorbarShowticksuffix = "all"
	ChoroplethmapboxColorbarShowticksuffix_first ChoroplethmapboxColorbarShowticksuffix = "first"
	ChoroplethmapboxColorbarShowticksuffix_last  ChoroplethmapboxColorbarShowticksuffix = "last"
	ChoroplethmapboxColorbarShowticksuffix_none  ChoroplethmapboxColorbarShowticksuffix = "none"
)

// ChoroplethmapboxColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ChoroplethmapboxColorbarThicknessmode string

const (
	ChoroplethmapboxColorbarThicknessmode_fraction ChoroplethmapboxColorbarThicknessmode = "fraction"
	ChoroplethmapboxColorbarThicknessmode_pixels   ChoroplethmapboxColorbarThicknessmode = "pixels"
)

// ChoroplethmapboxColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ChoroplethmapboxColorbarTickmode string

const (
	ChoroplethmapboxColorbarTickmode_auto   ChoroplethmapboxColorbarTickmode = "auto"
	ChoroplethmapboxColorbarTickmode_linear ChoroplethmapboxColorbarTickmode = "linear"
	ChoroplethmapboxColorbarTickmode_array  ChoroplethmapboxColorbarTickmode = "array"
)

// ChoroplethmapboxColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ChoroplethmapboxColorbarTicks string

const (
	ChoroplethmapboxColorbarTicks_outside ChoroplethmapboxColorbarTicks = "outside"
	ChoroplethmapboxColorbarTicks_inside  ChoroplethmapboxColorbarTicks = "inside"
)

// ChoroplethmapboxColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ChoroplethmapboxColorbarTitleSide string

const (
	ChoroplethmapboxColorbarTitleSide_right  ChoroplethmapboxColorbarTitleSide = "right"
	ChoroplethmapboxColorbarTitleSide_top    ChoroplethmapboxColorbarTitleSide = "top"
	ChoroplethmapboxColorbarTitleSide_bottom ChoroplethmapboxColorbarTitleSide = "bottom"
)

// ChoroplethmapboxColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ChoroplethmapboxColorbarXanchor string

const (
	ChoroplethmapboxColorbarXanchor_left   ChoroplethmapboxColorbarXanchor = "left"
	ChoroplethmapboxColorbarXanchor_center ChoroplethmapboxColorbarXanchor = "center"
	ChoroplethmapboxColorbarXanchor_right  ChoroplethmapboxColorbarXanchor = "right"
)

// ChoroplethmapboxColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ChoroplethmapboxColorbarYanchor string

const (
	ChoroplethmapboxColorbarYanchor_top    ChoroplethmapboxColorbarYanchor = "top"
	ChoroplethmapboxColorbarYanchor_middle ChoroplethmapboxColorbarYanchor = "middle"
	ChoroplethmapboxColorbarYanchor_bottom ChoroplethmapboxColorbarYanchor = "bottom"
)

// ChoroplethmapboxHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ChoroplethmapboxHoverlabelAlign string

const (
	ChoroplethmapboxHoverlabelAlign_left  ChoroplethmapboxHoverlabelAlign = "left"
	ChoroplethmapboxHoverlabelAlign_right ChoroplethmapboxHoverlabelAlign = "right"
	ChoroplethmapboxHoverlabelAlign_auto  ChoroplethmapboxHoverlabelAlign = "auto"
)

// ChoroplethmapboxVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ChoroplethmapboxVisible interface{}

var (
	ChoroplethmapboxVisible_True       ChoroplethmapboxVisible = true
	ChoroplethmapboxVisible_False      ChoroplethmapboxVisible = false
	ChoroplethmapboxVisible_legendonly ChoroplethmapboxVisible = "legendonly"
)

// ConeAnchor Sets the cones' anchor with respect to their x/y/z positions. Note that *cm* denote the cone's center of mass which corresponds to 1/4 from the tail to tip.
type ConeAnchor string

const (
	ConeAnchor_tip    ConeAnchor = "tip"
	ConeAnchor_tail   ConeAnchor = "tail"
	ConeAnchor_cm     ConeAnchor = "cm"
	ConeAnchor_center ConeAnchor = "center"
)

// ConeColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ConeColorbarExponentformat string

const (
	ConeColorbarExponentformat_none  ConeColorbarExponentformat = "none"
	ConeColorbarExponentformat_e     ConeColorbarExponentformat = "e"
	ConeColorbarExponentformat_E     ConeColorbarExponentformat = "E"
	ConeColorbarExponentformat_power ConeColorbarExponentformat = "power"
	ConeColorbarExponentformat_SI    ConeColorbarExponentformat = "SI"
	ConeColorbarExponentformat_B     ConeColorbarExponentformat = "B"
)

// ConeColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ConeColorbarLenmode string

const (
	ConeColorbarLenmode_fraction ConeColorbarLenmode = "fraction"
	ConeColorbarLenmode_pixels   ConeColorbarLenmode = "pixels"
)

// ConeColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ConeColorbarShowexponent string

const (
	ConeColorbarShowexponent_all   ConeColorbarShowexponent = "all"
	ConeColorbarShowexponent_first ConeColorbarShowexponent = "first"
	ConeColorbarShowexponent_last  ConeColorbarShowexponent = "last"
	ConeColorbarShowexponent_none  ConeColorbarShowexponent = "none"
)

// ConeColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ConeColorbarShowtickprefix string

const (
	ConeColorbarShowtickprefix_all   ConeColorbarShowtickprefix = "all"
	ConeColorbarShowtickprefix_first ConeColorbarShowtickprefix = "first"
	ConeColorbarShowtickprefix_last  ConeColorbarShowtickprefix = "last"
	ConeColorbarShowtickprefix_none  ConeColorbarShowtickprefix = "none"
)

// ConeColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ConeColorbarShowticksuffix string

const (
	ConeColorbarShowticksuffix_all   ConeColorbarShowticksuffix = "all"
	ConeColorbarShowticksuffix_first ConeColorbarShowticksuffix = "first"
	ConeColorbarShowticksuffix_last  ConeColorbarShowticksuffix = "last"
	ConeColorbarShowticksuffix_none  ConeColorbarShowticksuffix = "none"
)

// ConeColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ConeColorbarThicknessmode string

const (
	ConeColorbarThicknessmode_fraction ConeColorbarThicknessmode = "fraction"
	ConeColorbarThicknessmode_pixels   ConeColorbarThicknessmode = "pixels"
)

// ConeColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ConeColorbarTickmode string

const (
	ConeColorbarTickmode_auto   ConeColorbarTickmode = "auto"
	ConeColorbarTickmode_linear ConeColorbarTickmode = "linear"
	ConeColorbarTickmode_array  ConeColorbarTickmode = "array"
)

// ConeColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ConeColorbarTicks string

const (
	ConeColorbarTicks_outside ConeColorbarTicks = "outside"
	ConeColorbarTicks_inside  ConeColorbarTicks = "inside"
)

// ConeColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ConeColorbarTitleSide string

const (
	ConeColorbarTitleSide_right  ConeColorbarTitleSide = "right"
	ConeColorbarTitleSide_top    ConeColorbarTitleSide = "top"
	ConeColorbarTitleSide_bottom ConeColorbarTitleSide = "bottom"
)

// ConeColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ConeColorbarXanchor string

const (
	ConeColorbarXanchor_left   ConeColorbarXanchor = "left"
	ConeColorbarXanchor_center ConeColorbarXanchor = "center"
	ConeColorbarXanchor_right  ConeColorbarXanchor = "right"
)

// ConeColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ConeColorbarYanchor string

const (
	ConeColorbarYanchor_top    ConeColorbarYanchor = "top"
	ConeColorbarYanchor_middle ConeColorbarYanchor = "middle"
	ConeColorbarYanchor_bottom ConeColorbarYanchor = "bottom"
)

// ConeHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ConeHoverlabelAlign string

const (
	ConeHoverlabelAlign_left  ConeHoverlabelAlign = "left"
	ConeHoverlabelAlign_right ConeHoverlabelAlign = "right"
	ConeHoverlabelAlign_auto  ConeHoverlabelAlign = "auto"
)

// ConeSizemode Determines whether `sizeref` is set as a *scaled* (i.e unitless) scalar (normalized by the max u/v/w norm in the vector field) or as *absolute* value (in the same units as the vector field).
type ConeSizemode string

const (
	ConeSizemode_scaled   ConeSizemode = "scaled"
	ConeSizemode_absolute ConeSizemode = "absolute"
)

// ConeVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ConeVisible interface{}

var (
	ConeVisible_True       ConeVisible = true
	ConeVisible_False      ConeVisible = false
	ConeVisible_legendonly ConeVisible = "legendonly"
)

// ConfigDisplaymodebar Determines the mode bar display mode. If *true*, the mode bar is always visible. If *false*, the mode bar is always hidden. If *hover*, the mode bar is visible while the mouse cursor is on the graph container.
type ConfigDisplaymodebar interface{}

var (
	ConfigDisplaymodebar_hover ConfigDisplaymodebar = "hover"
	ConfigDisplaymodebar_True  ConfigDisplaymodebar = true
	ConfigDisplaymodebar_False ConfigDisplaymodebar = false
)

// ConfigDoubleclick Sets the double click interaction mode. Has an effect only in cartesian plots. If *false*, double click is disable. If *reset*, double click resets the axis ranges to their initial values. If *autosize*, double click set the axis ranges to their autorange values. If *reset+autosize*, the odd double clicks resets the axis ranges to their initial values and even double clicks set the axis ranges to their autorange values.
type ConfigDoubleclick interface{}

var (
	ConfigDoubleclick_False             ConfigDoubleclick = false
	ConfigDoubleclick_reset             ConfigDoubleclick = "reset"
	ConfigDoubleclick_autosize          ConfigDoubleclick = "autosize"
	ConfigDoubleclick_resetplusautosize ConfigDoubleclick = "reset+autosize"
)

// ContourColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ContourColorbarExponentformat string

const (
	ContourColorbarExponentformat_none  ContourColorbarExponentformat = "none"
	ContourColorbarExponentformat_e     ContourColorbarExponentformat = "e"
	ContourColorbarExponentformat_E     ContourColorbarExponentformat = "E"
	ContourColorbarExponentformat_power ContourColorbarExponentformat = "power"
	ContourColorbarExponentformat_SI    ContourColorbarExponentformat = "SI"
	ContourColorbarExponentformat_B     ContourColorbarExponentformat = "B"
)

// ContourColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ContourColorbarLenmode string

const (
	ContourColorbarLenmode_fraction ContourColorbarLenmode = "fraction"
	ContourColorbarLenmode_pixels   ContourColorbarLenmode = "pixels"
)

// ContourColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ContourColorbarShowexponent string

const (
	ContourColorbarShowexponent_all   ContourColorbarShowexponent = "all"
	ContourColorbarShowexponent_first ContourColorbarShowexponent = "first"
	ContourColorbarShowexponent_last  ContourColorbarShowexponent = "last"
	ContourColorbarShowexponent_none  ContourColorbarShowexponent = "none"
)

// ContourColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ContourColorbarShowtickprefix string

const (
	ContourColorbarShowtickprefix_all   ContourColorbarShowtickprefix = "all"
	ContourColorbarShowtickprefix_first ContourColorbarShowtickprefix = "first"
	ContourColorbarShowtickprefix_last  ContourColorbarShowtickprefix = "last"
	ContourColorbarShowtickprefix_none  ContourColorbarShowtickprefix = "none"
)

// ContourColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ContourColorbarShowticksuffix string

const (
	ContourColorbarShowticksuffix_all   ContourColorbarShowticksuffix = "all"
	ContourColorbarShowticksuffix_first ContourColorbarShowticksuffix = "first"
	ContourColorbarShowticksuffix_last  ContourColorbarShowticksuffix = "last"
	ContourColorbarShowticksuffix_none  ContourColorbarShowticksuffix = "none"
)

// ContourColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ContourColorbarThicknessmode string

const (
	ContourColorbarThicknessmode_fraction ContourColorbarThicknessmode = "fraction"
	ContourColorbarThicknessmode_pixels   ContourColorbarThicknessmode = "pixels"
)

// ContourColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ContourColorbarTickmode string

const (
	ContourColorbarTickmode_auto   ContourColorbarTickmode = "auto"
	ContourColorbarTickmode_linear ContourColorbarTickmode = "linear"
	ContourColorbarTickmode_array  ContourColorbarTickmode = "array"
)

// ContourColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ContourColorbarTicks string

const (
	ContourColorbarTicks_outside ContourColorbarTicks = "outside"
	ContourColorbarTicks_inside  ContourColorbarTicks = "inside"
)

// ContourColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ContourColorbarTitleSide string

const (
	ContourColorbarTitleSide_right  ContourColorbarTitleSide = "right"
	ContourColorbarTitleSide_top    ContourColorbarTitleSide = "top"
	ContourColorbarTitleSide_bottom ContourColorbarTitleSide = "bottom"
)

// ContourColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ContourColorbarXanchor string

const (
	ContourColorbarXanchor_left   ContourColorbarXanchor = "left"
	ContourColorbarXanchor_center ContourColorbarXanchor = "center"
	ContourColorbarXanchor_right  ContourColorbarXanchor = "right"
)

// ContourColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ContourColorbarYanchor string

const (
	ContourColorbarYanchor_top    ContourColorbarYanchor = "top"
	ContourColorbarYanchor_middle ContourColorbarYanchor = "middle"
	ContourColorbarYanchor_bottom ContourColorbarYanchor = "bottom"
)

// ContourContoursColoring Determines the coloring method showing the contour values. If *fill*, coloring is done evenly between each contour level If *heatmap*, a heatmap gradient coloring is applied between each contour level. If *lines*, coloring is done on the contour lines. If *none*, no coloring is applied on this trace.
type ContourContoursColoring string

const (
	ContourContoursColoring_fill    ContourContoursColoring = "fill"
	ContourContoursColoring_heatmap ContourContoursColoring = "heatmap"
	ContourContoursColoring_lines   ContourContoursColoring = "lines"
	ContourContoursColoring_none    ContourContoursColoring = "none"
)

// ContourContoursOperation Sets the constraint operation. *=* keeps regions equal to `value` *<* and *<=* keep regions less than `value` *>* and *>=* keep regions greater than `value` *[]*, *()*, *[)*, and *(]* keep regions inside `value[0]` to `value[1]` *][*, *)(*, *](*, *)[* keep regions outside `value[0]` to value[1]` Open vs. closed intervals make no difference to constraint display, but all versions are allowed for consistency with filter transforms.
type ContourContoursOperation string

const (
	ContourContoursOperation_eq                   ContourContoursOperation = "="
	ContourContoursOperation_lt                   ContourContoursOperation = "<"
	ContourContoursOperation_gteq                 ContourContoursOperation = ">="
	ContourContoursOperation_gt                   ContourContoursOperation = ">"
	ContourContoursOperation_lteq                 ContourContoursOperation = "<="
	ContourContoursOperation__lbracket__rbracket_ ContourContoursOperation = "[]"
	ContourContoursOperation__lpar__rpar_         ContourContoursOperation = "()"
	ContourContoursOperation__lbracket__rpar_     ContourContoursOperation = "[)"
	ContourContoursOperation__lpar__rbracket_     ContourContoursOperation = "(]"
	ContourContoursOperation__rbracket__lbracket_ ContourContoursOperation = "]["
	ContourContoursOperation__rpar__lpar_         ContourContoursOperation = ")("
	ContourContoursOperation__rbracket__lpar_     ContourContoursOperation = "]("
	ContourContoursOperation__rpar__lbracket_     ContourContoursOperation = ")["
)

// ContourContoursType If `levels`, the data is represented as a contour plot with multiple levels displayed. If `constraint`, the data is represented as constraints with the invalid region shaded as specified by the `operation` and `value` parameters.
type ContourContoursType string

const (
	ContourContoursType_levels     ContourContoursType = "levels"
	ContourContoursType_constraint ContourContoursType = "constraint"
)

// ContourHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ContourHoverlabelAlign string

const (
	ContourHoverlabelAlign_left  ContourHoverlabelAlign = "left"
	ContourHoverlabelAlign_right ContourHoverlabelAlign = "right"
	ContourHoverlabelAlign_auto  ContourHoverlabelAlign = "auto"
)

// ContourVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ContourVisible interface{}

var (
	ContourVisible_True       ContourVisible = true
	ContourVisible_False      ContourVisible = false
	ContourVisible_legendonly ContourVisible = "legendonly"
)

// ContourXcalendar Sets the calendar system to use with `x` date data.
type ContourXcalendar string

const (
	ContourXcalendar_gregorian  ContourXcalendar = "gregorian"
	ContourXcalendar_chinese    ContourXcalendar = "chinese"
	ContourXcalendar_coptic     ContourXcalendar = "coptic"
	ContourXcalendar_discworld  ContourXcalendar = "discworld"
	ContourXcalendar_ethiopian  ContourXcalendar = "ethiopian"
	ContourXcalendar_hebrew     ContourXcalendar = "hebrew"
	ContourXcalendar_islamic    ContourXcalendar = "islamic"
	ContourXcalendar_julian     ContourXcalendar = "julian"
	ContourXcalendar_mayan      ContourXcalendar = "mayan"
	ContourXcalendar_nanakshahi ContourXcalendar = "nanakshahi"
	ContourXcalendar_nepali     ContourXcalendar = "nepali"
	ContourXcalendar_persian    ContourXcalendar = "persian"
	ContourXcalendar_jalali     ContourXcalendar = "jalali"
	ContourXcalendar_taiwan     ContourXcalendar = "taiwan"
	ContourXcalendar_thai       ContourXcalendar = "thai"
	ContourXcalendar_ummalqura  ContourXcalendar = "ummalqura"
)

// ContourXtype If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
type ContourXtype string

const (
	ContourXtype_array  ContourXtype = "array"
	ContourXtype_scaled ContourXtype = "scaled"
)

// ContourYcalendar Sets the calendar system to use with `y` date data.
type ContourYcalendar string

const (
	ContourYcalendar_gregorian  ContourYcalendar = "gregorian"
	ContourYcalendar_chinese    ContourYcalendar = "chinese"
	ContourYcalendar_coptic     ContourYcalendar = "coptic"
	ContourYcalendar_discworld  ContourYcalendar = "discworld"
	ContourYcalendar_ethiopian  ContourYcalendar = "ethiopian"
	ContourYcalendar_hebrew     ContourYcalendar = "hebrew"
	ContourYcalendar_islamic    ContourYcalendar = "islamic"
	ContourYcalendar_julian     ContourYcalendar = "julian"
	ContourYcalendar_mayan      ContourYcalendar = "mayan"
	ContourYcalendar_nanakshahi ContourYcalendar = "nanakshahi"
	ContourYcalendar_nepali     ContourYcalendar = "nepali"
	ContourYcalendar_persian    ContourYcalendar = "persian"
	ContourYcalendar_jalali     ContourYcalendar = "jalali"
	ContourYcalendar_taiwan     ContourYcalendar = "taiwan"
	ContourYcalendar_thai       ContourYcalendar = "thai"
	ContourYcalendar_ummalqura  ContourYcalendar = "ummalqura"
)

// ContourYtype If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
type ContourYtype string

const (
	ContourYtype_array  ContourYtype = "array"
	ContourYtype_scaled ContourYtype = "scaled"
)

// ContourcarpetAtype If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
type ContourcarpetAtype string

const (
	ContourcarpetAtype_array  ContourcarpetAtype = "array"
	ContourcarpetAtype_scaled ContourcarpetAtype = "scaled"
)

// ContourcarpetBtype If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
type ContourcarpetBtype string

const (
	ContourcarpetBtype_array  ContourcarpetBtype = "array"
	ContourcarpetBtype_scaled ContourcarpetBtype = "scaled"
)

// ContourcarpetColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ContourcarpetColorbarExponentformat string

const (
	ContourcarpetColorbarExponentformat_none  ContourcarpetColorbarExponentformat = "none"
	ContourcarpetColorbarExponentformat_e     ContourcarpetColorbarExponentformat = "e"
	ContourcarpetColorbarExponentformat_E     ContourcarpetColorbarExponentformat = "E"
	ContourcarpetColorbarExponentformat_power ContourcarpetColorbarExponentformat = "power"
	ContourcarpetColorbarExponentformat_SI    ContourcarpetColorbarExponentformat = "SI"
	ContourcarpetColorbarExponentformat_B     ContourcarpetColorbarExponentformat = "B"
)

// ContourcarpetColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ContourcarpetColorbarLenmode string

const (
	ContourcarpetColorbarLenmode_fraction ContourcarpetColorbarLenmode = "fraction"
	ContourcarpetColorbarLenmode_pixels   ContourcarpetColorbarLenmode = "pixels"
)

// ContourcarpetColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ContourcarpetColorbarShowexponent string

const (
	ContourcarpetColorbarShowexponent_all   ContourcarpetColorbarShowexponent = "all"
	ContourcarpetColorbarShowexponent_first ContourcarpetColorbarShowexponent = "first"
	ContourcarpetColorbarShowexponent_last  ContourcarpetColorbarShowexponent = "last"
	ContourcarpetColorbarShowexponent_none  ContourcarpetColorbarShowexponent = "none"
)

// ContourcarpetColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ContourcarpetColorbarShowtickprefix string

const (
	ContourcarpetColorbarShowtickprefix_all   ContourcarpetColorbarShowtickprefix = "all"
	ContourcarpetColorbarShowtickprefix_first ContourcarpetColorbarShowtickprefix = "first"
	ContourcarpetColorbarShowtickprefix_last  ContourcarpetColorbarShowtickprefix = "last"
	ContourcarpetColorbarShowtickprefix_none  ContourcarpetColorbarShowtickprefix = "none"
)

// ContourcarpetColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ContourcarpetColorbarShowticksuffix string

const (
	ContourcarpetColorbarShowticksuffix_all   ContourcarpetColorbarShowticksuffix = "all"
	ContourcarpetColorbarShowticksuffix_first ContourcarpetColorbarShowticksuffix = "first"
	ContourcarpetColorbarShowticksuffix_last  ContourcarpetColorbarShowticksuffix = "last"
	ContourcarpetColorbarShowticksuffix_none  ContourcarpetColorbarShowticksuffix = "none"
)

// ContourcarpetColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ContourcarpetColorbarThicknessmode string

const (
	ContourcarpetColorbarThicknessmode_fraction ContourcarpetColorbarThicknessmode = "fraction"
	ContourcarpetColorbarThicknessmode_pixels   ContourcarpetColorbarThicknessmode = "pixels"
)

// ContourcarpetColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ContourcarpetColorbarTickmode string

const (
	ContourcarpetColorbarTickmode_auto   ContourcarpetColorbarTickmode = "auto"
	ContourcarpetColorbarTickmode_linear ContourcarpetColorbarTickmode = "linear"
	ContourcarpetColorbarTickmode_array  ContourcarpetColorbarTickmode = "array"
)

// ContourcarpetColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ContourcarpetColorbarTicks string

const (
	ContourcarpetColorbarTicks_outside ContourcarpetColorbarTicks = "outside"
	ContourcarpetColorbarTicks_inside  ContourcarpetColorbarTicks = "inside"
)

// ContourcarpetColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ContourcarpetColorbarTitleSide string

const (
	ContourcarpetColorbarTitleSide_right  ContourcarpetColorbarTitleSide = "right"
	ContourcarpetColorbarTitleSide_top    ContourcarpetColorbarTitleSide = "top"
	ContourcarpetColorbarTitleSide_bottom ContourcarpetColorbarTitleSide = "bottom"
)

// ContourcarpetColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ContourcarpetColorbarXanchor string

const (
	ContourcarpetColorbarXanchor_left   ContourcarpetColorbarXanchor = "left"
	ContourcarpetColorbarXanchor_center ContourcarpetColorbarXanchor = "center"
	ContourcarpetColorbarXanchor_right  ContourcarpetColorbarXanchor = "right"
)

// ContourcarpetColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ContourcarpetColorbarYanchor string

const (
	ContourcarpetColorbarYanchor_top    ContourcarpetColorbarYanchor = "top"
	ContourcarpetColorbarYanchor_middle ContourcarpetColorbarYanchor = "middle"
	ContourcarpetColorbarYanchor_bottom ContourcarpetColorbarYanchor = "bottom"
)

// ContourcarpetContoursColoring Determines the coloring method showing the contour values. If *fill*, coloring is done evenly between each contour level If *lines*, coloring is done on the contour lines. If *none*, no coloring is applied on this trace.
type ContourcarpetContoursColoring string

const (
	ContourcarpetContoursColoring_fill  ContourcarpetContoursColoring = "fill"
	ContourcarpetContoursColoring_lines ContourcarpetContoursColoring = "lines"
	ContourcarpetContoursColoring_none  ContourcarpetContoursColoring = "none"
)

// ContourcarpetContoursOperation Sets the constraint operation. *=* keeps regions equal to `value` *<* and *<=* keep regions less than `value` *>* and *>=* keep regions greater than `value` *[]*, *()*, *[)*, and *(]* keep regions inside `value[0]` to `value[1]` *][*, *)(*, *](*, *)[* keep regions outside `value[0]` to value[1]` Open vs. closed intervals make no difference to constraint display, but all versions are allowed for consistency with filter transforms.
type ContourcarpetContoursOperation string

const (
	ContourcarpetContoursOperation_eq                   ContourcarpetContoursOperation = "="
	ContourcarpetContoursOperation_lt                   ContourcarpetContoursOperation = "<"
	ContourcarpetContoursOperation_gteq                 ContourcarpetContoursOperation = ">="
	ContourcarpetContoursOperation_gt                   ContourcarpetContoursOperation = ">"
	ContourcarpetContoursOperation_lteq                 ContourcarpetContoursOperation = "<="
	ContourcarpetContoursOperation__lbracket__rbracket_ ContourcarpetContoursOperation = "[]"
	ContourcarpetContoursOperation__lpar__rpar_         ContourcarpetContoursOperation = "()"
	ContourcarpetContoursOperation__lbracket__rpar_     ContourcarpetContoursOperation = "[)"
	ContourcarpetContoursOperation__lpar__rbracket_     ContourcarpetContoursOperation = "(]"
	ContourcarpetContoursOperation__rbracket__lbracket_ ContourcarpetContoursOperation = "]["
	ContourcarpetContoursOperation__rpar__lpar_         ContourcarpetContoursOperation = ")("
	ContourcarpetContoursOperation__rbracket__lpar_     ContourcarpetContoursOperation = "]("
	ContourcarpetContoursOperation__rpar__lbracket_     ContourcarpetContoursOperation = ")["
)

// ContourcarpetContoursType If `levels`, the data is represented as a contour plot with multiple levels displayed. If `constraint`, the data is represented as constraints with the invalid region shaded as specified by the `operation` and `value` parameters.
type ContourcarpetContoursType string

const (
	ContourcarpetContoursType_levels     ContourcarpetContoursType = "levels"
	ContourcarpetContoursType_constraint ContourcarpetContoursType = "constraint"
)

// ContourcarpetVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ContourcarpetVisible interface{}

var (
	ContourcarpetVisible_True       ContourcarpetVisible = true
	ContourcarpetVisible_False      ContourcarpetVisible = false
	ContourcarpetVisible_legendonly ContourcarpetVisible = "legendonly"
)

// DensitymapboxColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type DensitymapboxColorbarExponentformat string

const (
	DensitymapboxColorbarExponentformat_none  DensitymapboxColorbarExponentformat = "none"
	DensitymapboxColorbarExponentformat_e     DensitymapboxColorbarExponentformat = "e"
	DensitymapboxColorbarExponentformat_E     DensitymapboxColorbarExponentformat = "E"
	DensitymapboxColorbarExponentformat_power DensitymapboxColorbarExponentformat = "power"
	DensitymapboxColorbarExponentformat_SI    DensitymapboxColorbarExponentformat = "SI"
	DensitymapboxColorbarExponentformat_B     DensitymapboxColorbarExponentformat = "B"
)

// DensitymapboxColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type DensitymapboxColorbarLenmode string

const (
	DensitymapboxColorbarLenmode_fraction DensitymapboxColorbarLenmode = "fraction"
	DensitymapboxColorbarLenmode_pixels   DensitymapboxColorbarLenmode = "pixels"
)

// DensitymapboxColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type DensitymapboxColorbarShowexponent string

const (
	DensitymapboxColorbarShowexponent_all   DensitymapboxColorbarShowexponent = "all"
	DensitymapboxColorbarShowexponent_first DensitymapboxColorbarShowexponent = "first"
	DensitymapboxColorbarShowexponent_last  DensitymapboxColorbarShowexponent = "last"
	DensitymapboxColorbarShowexponent_none  DensitymapboxColorbarShowexponent = "none"
)

// DensitymapboxColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type DensitymapboxColorbarShowtickprefix string

const (
	DensitymapboxColorbarShowtickprefix_all   DensitymapboxColorbarShowtickprefix = "all"
	DensitymapboxColorbarShowtickprefix_first DensitymapboxColorbarShowtickprefix = "first"
	DensitymapboxColorbarShowtickprefix_last  DensitymapboxColorbarShowtickprefix = "last"
	DensitymapboxColorbarShowtickprefix_none  DensitymapboxColorbarShowtickprefix = "none"
)

// DensitymapboxColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type DensitymapboxColorbarShowticksuffix string

const (
	DensitymapboxColorbarShowticksuffix_all   DensitymapboxColorbarShowticksuffix = "all"
	DensitymapboxColorbarShowticksuffix_first DensitymapboxColorbarShowticksuffix = "first"
	DensitymapboxColorbarShowticksuffix_last  DensitymapboxColorbarShowticksuffix = "last"
	DensitymapboxColorbarShowticksuffix_none  DensitymapboxColorbarShowticksuffix = "none"
)

// DensitymapboxColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type DensitymapboxColorbarThicknessmode string

const (
	DensitymapboxColorbarThicknessmode_fraction DensitymapboxColorbarThicknessmode = "fraction"
	DensitymapboxColorbarThicknessmode_pixels   DensitymapboxColorbarThicknessmode = "pixels"
)

// DensitymapboxColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type DensitymapboxColorbarTickmode string

const (
	DensitymapboxColorbarTickmode_auto   DensitymapboxColorbarTickmode = "auto"
	DensitymapboxColorbarTickmode_linear DensitymapboxColorbarTickmode = "linear"
	DensitymapboxColorbarTickmode_array  DensitymapboxColorbarTickmode = "array"
)

// DensitymapboxColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type DensitymapboxColorbarTicks string

const (
	DensitymapboxColorbarTicks_outside DensitymapboxColorbarTicks = "outside"
	DensitymapboxColorbarTicks_inside  DensitymapboxColorbarTicks = "inside"
)

// DensitymapboxColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type DensitymapboxColorbarTitleSide string

const (
	DensitymapboxColorbarTitleSide_right  DensitymapboxColorbarTitleSide = "right"
	DensitymapboxColorbarTitleSide_top    DensitymapboxColorbarTitleSide = "top"
	DensitymapboxColorbarTitleSide_bottom DensitymapboxColorbarTitleSide = "bottom"
)

// DensitymapboxColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type DensitymapboxColorbarXanchor string

const (
	DensitymapboxColorbarXanchor_left   DensitymapboxColorbarXanchor = "left"
	DensitymapboxColorbarXanchor_center DensitymapboxColorbarXanchor = "center"
	DensitymapboxColorbarXanchor_right  DensitymapboxColorbarXanchor = "right"
)

// DensitymapboxColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type DensitymapboxColorbarYanchor string

const (
	DensitymapboxColorbarYanchor_top    DensitymapboxColorbarYanchor = "top"
	DensitymapboxColorbarYanchor_middle DensitymapboxColorbarYanchor = "middle"
	DensitymapboxColorbarYanchor_bottom DensitymapboxColorbarYanchor = "bottom"
)

// DensitymapboxHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type DensitymapboxHoverlabelAlign string

const (
	DensitymapboxHoverlabelAlign_left  DensitymapboxHoverlabelAlign = "left"
	DensitymapboxHoverlabelAlign_right DensitymapboxHoverlabelAlign = "right"
	DensitymapboxHoverlabelAlign_auto  DensitymapboxHoverlabelAlign = "auto"
)

// DensitymapboxVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type DensitymapboxVisible interface{}

var (
	DensitymapboxVisible_True       DensitymapboxVisible = true
	DensitymapboxVisible_False      DensitymapboxVisible = false
	DensitymapboxVisible_legendonly DensitymapboxVisible = "legendonly"
)

// FunnelConstraintext Constrain the size of text inside or outside a bar to be no larger than the bar itself.
type FunnelConstraintext string

const (
	FunnelConstraintext_inside  FunnelConstraintext = "inside"
	FunnelConstraintext_outside FunnelConstraintext = "outside"
	FunnelConstraintext_both    FunnelConstraintext = "both"
	FunnelConstraintext_none    FunnelConstraintext = "none"
)

// FunnelHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type FunnelHoverlabelAlign string

const (
	FunnelHoverlabelAlign_left  FunnelHoverlabelAlign = "left"
	FunnelHoverlabelAlign_right FunnelHoverlabelAlign = "right"
	FunnelHoverlabelAlign_auto  FunnelHoverlabelAlign = "auto"
)

// FunnelInsidetextanchor Determines if texts are kept at center or start/end points in `textposition` *inside* mode.
type FunnelInsidetextanchor string

const (
	FunnelInsidetextanchor_end    FunnelInsidetextanchor = "end"
	FunnelInsidetextanchor_middle FunnelInsidetextanchor = "middle"
	FunnelInsidetextanchor_start  FunnelInsidetextanchor = "start"
)

// FunnelMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type FunnelMarkerColorbarExponentformat string

const (
	FunnelMarkerColorbarExponentformat_none  FunnelMarkerColorbarExponentformat = "none"
	FunnelMarkerColorbarExponentformat_e     FunnelMarkerColorbarExponentformat = "e"
	FunnelMarkerColorbarExponentformat_E     FunnelMarkerColorbarExponentformat = "E"
	FunnelMarkerColorbarExponentformat_power FunnelMarkerColorbarExponentformat = "power"
	FunnelMarkerColorbarExponentformat_SI    FunnelMarkerColorbarExponentformat = "SI"
	FunnelMarkerColorbarExponentformat_B     FunnelMarkerColorbarExponentformat = "B"
)

// FunnelMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type FunnelMarkerColorbarLenmode string

const (
	FunnelMarkerColorbarLenmode_fraction FunnelMarkerColorbarLenmode = "fraction"
	FunnelMarkerColorbarLenmode_pixels   FunnelMarkerColorbarLenmode = "pixels"
)

// FunnelMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type FunnelMarkerColorbarShowexponent string

const (
	FunnelMarkerColorbarShowexponent_all   FunnelMarkerColorbarShowexponent = "all"
	FunnelMarkerColorbarShowexponent_first FunnelMarkerColorbarShowexponent = "first"
	FunnelMarkerColorbarShowexponent_last  FunnelMarkerColorbarShowexponent = "last"
	FunnelMarkerColorbarShowexponent_none  FunnelMarkerColorbarShowexponent = "none"
)

// FunnelMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type FunnelMarkerColorbarShowtickprefix string

const (
	FunnelMarkerColorbarShowtickprefix_all   FunnelMarkerColorbarShowtickprefix = "all"
	FunnelMarkerColorbarShowtickprefix_first FunnelMarkerColorbarShowtickprefix = "first"
	FunnelMarkerColorbarShowtickprefix_last  FunnelMarkerColorbarShowtickprefix = "last"
	FunnelMarkerColorbarShowtickprefix_none  FunnelMarkerColorbarShowtickprefix = "none"
)

// FunnelMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type FunnelMarkerColorbarShowticksuffix string

const (
	FunnelMarkerColorbarShowticksuffix_all   FunnelMarkerColorbarShowticksuffix = "all"
	FunnelMarkerColorbarShowticksuffix_first FunnelMarkerColorbarShowticksuffix = "first"
	FunnelMarkerColorbarShowticksuffix_last  FunnelMarkerColorbarShowticksuffix = "last"
	FunnelMarkerColorbarShowticksuffix_none  FunnelMarkerColorbarShowticksuffix = "none"
)

// FunnelMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type FunnelMarkerColorbarThicknessmode string

const (
	FunnelMarkerColorbarThicknessmode_fraction FunnelMarkerColorbarThicknessmode = "fraction"
	FunnelMarkerColorbarThicknessmode_pixels   FunnelMarkerColorbarThicknessmode = "pixels"
)

// FunnelMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type FunnelMarkerColorbarTickmode string

const (
	FunnelMarkerColorbarTickmode_auto   FunnelMarkerColorbarTickmode = "auto"
	FunnelMarkerColorbarTickmode_linear FunnelMarkerColorbarTickmode = "linear"
	FunnelMarkerColorbarTickmode_array  FunnelMarkerColorbarTickmode = "array"
)

// FunnelMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type FunnelMarkerColorbarTicks string

const (
	FunnelMarkerColorbarTicks_outside FunnelMarkerColorbarTicks = "outside"
	FunnelMarkerColorbarTicks_inside  FunnelMarkerColorbarTicks = "inside"
)

// FunnelMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type FunnelMarkerColorbarTitleSide string

const (
	FunnelMarkerColorbarTitleSide_right  FunnelMarkerColorbarTitleSide = "right"
	FunnelMarkerColorbarTitleSide_top    FunnelMarkerColorbarTitleSide = "top"
	FunnelMarkerColorbarTitleSide_bottom FunnelMarkerColorbarTitleSide = "bottom"
)

// FunnelMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type FunnelMarkerColorbarXanchor string

const (
	FunnelMarkerColorbarXanchor_left   FunnelMarkerColorbarXanchor = "left"
	FunnelMarkerColorbarXanchor_center FunnelMarkerColorbarXanchor = "center"
	FunnelMarkerColorbarXanchor_right  FunnelMarkerColorbarXanchor = "right"
)

// FunnelMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type FunnelMarkerColorbarYanchor string

const (
	FunnelMarkerColorbarYanchor_top    FunnelMarkerColorbarYanchor = "top"
	FunnelMarkerColorbarYanchor_middle FunnelMarkerColorbarYanchor = "middle"
	FunnelMarkerColorbarYanchor_bottom FunnelMarkerColorbarYanchor = "bottom"
)

// FunnelOrientation Sets the orientation of the funnels. With *v* (*h*), the value of the each bar spans along the vertical (horizontal). By default funnels are tend to be oriented horizontally; unless only *y* array is presented or orientation is set to *v*. Also regarding graphs including only 'horizontal' funnels, *autorange* on the *y-axis* are set to *reversed*.
type FunnelOrientation string

const (
	FunnelOrientation_v FunnelOrientation = "v"
	FunnelOrientation_h FunnelOrientation = "h"
)

// FunnelTextposition Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside.
type FunnelTextposition string

const (
	FunnelTextposition_inside  FunnelTextposition = "inside"
	FunnelTextposition_outside FunnelTextposition = "outside"
	FunnelTextposition_auto    FunnelTextposition = "auto"
	FunnelTextposition_none    FunnelTextposition = "none"
)

// FunnelVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type FunnelVisible interface{}

var (
	FunnelVisible_True       FunnelVisible = true
	FunnelVisible_False      FunnelVisible = false
	FunnelVisible_legendonly FunnelVisible = "legendonly"
)

// FunnelareaHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type FunnelareaHoverlabelAlign string

const (
	FunnelareaHoverlabelAlign_left  FunnelareaHoverlabelAlign = "left"
	FunnelareaHoverlabelAlign_right FunnelareaHoverlabelAlign = "right"
	FunnelareaHoverlabelAlign_auto  FunnelareaHoverlabelAlign = "auto"
)

// FunnelareaTextposition Specifies the location of the `textinfo`.
type FunnelareaTextposition string

const (
	FunnelareaTextposition_inside FunnelareaTextposition = "inside"
	FunnelareaTextposition_none   FunnelareaTextposition = "none"
)

// FunnelareaTitlePosition Specifies the location of the `title`. Note that the title's position used to be set by the now deprecated `titleposition` attribute.
type FunnelareaTitlePosition string

const (
	FunnelareaTitlePosition_topleft   FunnelareaTitlePosition = "top left"
	FunnelareaTitlePosition_topcenter FunnelareaTitlePosition = "top center"
	FunnelareaTitlePosition_topright  FunnelareaTitlePosition = "top right"
)

// FunnelareaVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type FunnelareaVisible interface{}

var (
	FunnelareaVisible_True       FunnelareaVisible = true
	FunnelareaVisible_False      FunnelareaVisible = false
	FunnelareaVisible_legendonly FunnelareaVisible = "legendonly"
)

// HeatmapColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type HeatmapColorbarExponentformat string

const (
	HeatmapColorbarExponentformat_none  HeatmapColorbarExponentformat = "none"
	HeatmapColorbarExponentformat_e     HeatmapColorbarExponentformat = "e"
	HeatmapColorbarExponentformat_E     HeatmapColorbarExponentformat = "E"
	HeatmapColorbarExponentformat_power HeatmapColorbarExponentformat = "power"
	HeatmapColorbarExponentformat_SI    HeatmapColorbarExponentformat = "SI"
	HeatmapColorbarExponentformat_B     HeatmapColorbarExponentformat = "B"
)

// HeatmapColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type HeatmapColorbarLenmode string

const (
	HeatmapColorbarLenmode_fraction HeatmapColorbarLenmode = "fraction"
	HeatmapColorbarLenmode_pixels   HeatmapColorbarLenmode = "pixels"
)

// HeatmapColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type HeatmapColorbarShowexponent string

const (
	HeatmapColorbarShowexponent_all   HeatmapColorbarShowexponent = "all"
	HeatmapColorbarShowexponent_first HeatmapColorbarShowexponent = "first"
	HeatmapColorbarShowexponent_last  HeatmapColorbarShowexponent = "last"
	HeatmapColorbarShowexponent_none  HeatmapColorbarShowexponent = "none"
)

// HeatmapColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type HeatmapColorbarShowtickprefix string

const (
	HeatmapColorbarShowtickprefix_all   HeatmapColorbarShowtickprefix = "all"
	HeatmapColorbarShowtickprefix_first HeatmapColorbarShowtickprefix = "first"
	HeatmapColorbarShowtickprefix_last  HeatmapColorbarShowtickprefix = "last"
	HeatmapColorbarShowtickprefix_none  HeatmapColorbarShowtickprefix = "none"
)

// HeatmapColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type HeatmapColorbarShowticksuffix string

const (
	HeatmapColorbarShowticksuffix_all   HeatmapColorbarShowticksuffix = "all"
	HeatmapColorbarShowticksuffix_first HeatmapColorbarShowticksuffix = "first"
	HeatmapColorbarShowticksuffix_last  HeatmapColorbarShowticksuffix = "last"
	HeatmapColorbarShowticksuffix_none  HeatmapColorbarShowticksuffix = "none"
)

// HeatmapColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type HeatmapColorbarThicknessmode string

const (
	HeatmapColorbarThicknessmode_fraction HeatmapColorbarThicknessmode = "fraction"
	HeatmapColorbarThicknessmode_pixels   HeatmapColorbarThicknessmode = "pixels"
)

// HeatmapColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type HeatmapColorbarTickmode string

const (
	HeatmapColorbarTickmode_auto   HeatmapColorbarTickmode = "auto"
	HeatmapColorbarTickmode_linear HeatmapColorbarTickmode = "linear"
	HeatmapColorbarTickmode_array  HeatmapColorbarTickmode = "array"
)

// HeatmapColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type HeatmapColorbarTicks string

const (
	HeatmapColorbarTicks_outside HeatmapColorbarTicks = "outside"
	HeatmapColorbarTicks_inside  HeatmapColorbarTicks = "inside"
)

// HeatmapColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type HeatmapColorbarTitleSide string

const (
	HeatmapColorbarTitleSide_right  HeatmapColorbarTitleSide = "right"
	HeatmapColorbarTitleSide_top    HeatmapColorbarTitleSide = "top"
	HeatmapColorbarTitleSide_bottom HeatmapColorbarTitleSide = "bottom"
)

// HeatmapColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type HeatmapColorbarXanchor string

const (
	HeatmapColorbarXanchor_left   HeatmapColorbarXanchor = "left"
	HeatmapColorbarXanchor_center HeatmapColorbarXanchor = "center"
	HeatmapColorbarXanchor_right  HeatmapColorbarXanchor = "right"
)

// HeatmapColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type HeatmapColorbarYanchor string

const (
	HeatmapColorbarYanchor_top    HeatmapColorbarYanchor = "top"
	HeatmapColorbarYanchor_middle HeatmapColorbarYanchor = "middle"
	HeatmapColorbarYanchor_bottom HeatmapColorbarYanchor = "bottom"
)

// HeatmapHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type HeatmapHoverlabelAlign string

const (
	HeatmapHoverlabelAlign_left  HeatmapHoverlabelAlign = "left"
	HeatmapHoverlabelAlign_right HeatmapHoverlabelAlign = "right"
	HeatmapHoverlabelAlign_auto  HeatmapHoverlabelAlign = "auto"
)

// HeatmapVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type HeatmapVisible interface{}

var (
	HeatmapVisible_True       HeatmapVisible = true
	HeatmapVisible_False      HeatmapVisible = false
	HeatmapVisible_legendonly HeatmapVisible = "legendonly"
)

// HeatmapXcalendar Sets the calendar system to use with `x` date data.
type HeatmapXcalendar string

const (
	HeatmapXcalendar_gregorian  HeatmapXcalendar = "gregorian"
	HeatmapXcalendar_chinese    HeatmapXcalendar = "chinese"
	HeatmapXcalendar_coptic     HeatmapXcalendar = "coptic"
	HeatmapXcalendar_discworld  HeatmapXcalendar = "discworld"
	HeatmapXcalendar_ethiopian  HeatmapXcalendar = "ethiopian"
	HeatmapXcalendar_hebrew     HeatmapXcalendar = "hebrew"
	HeatmapXcalendar_islamic    HeatmapXcalendar = "islamic"
	HeatmapXcalendar_julian     HeatmapXcalendar = "julian"
	HeatmapXcalendar_mayan      HeatmapXcalendar = "mayan"
	HeatmapXcalendar_nanakshahi HeatmapXcalendar = "nanakshahi"
	HeatmapXcalendar_nepali     HeatmapXcalendar = "nepali"
	HeatmapXcalendar_persian    HeatmapXcalendar = "persian"
	HeatmapXcalendar_jalali     HeatmapXcalendar = "jalali"
	HeatmapXcalendar_taiwan     HeatmapXcalendar = "taiwan"
	HeatmapXcalendar_thai       HeatmapXcalendar = "thai"
	HeatmapXcalendar_ummalqura  HeatmapXcalendar = "ummalqura"
)

// HeatmapXtype If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
type HeatmapXtype string

const (
	HeatmapXtype_array  HeatmapXtype = "array"
	HeatmapXtype_scaled HeatmapXtype = "scaled"
)

// HeatmapYcalendar Sets the calendar system to use with `y` date data.
type HeatmapYcalendar string

const (
	HeatmapYcalendar_gregorian  HeatmapYcalendar = "gregorian"
	HeatmapYcalendar_chinese    HeatmapYcalendar = "chinese"
	HeatmapYcalendar_coptic     HeatmapYcalendar = "coptic"
	HeatmapYcalendar_discworld  HeatmapYcalendar = "discworld"
	HeatmapYcalendar_ethiopian  HeatmapYcalendar = "ethiopian"
	HeatmapYcalendar_hebrew     HeatmapYcalendar = "hebrew"
	HeatmapYcalendar_islamic    HeatmapYcalendar = "islamic"
	HeatmapYcalendar_julian     HeatmapYcalendar = "julian"
	HeatmapYcalendar_mayan      HeatmapYcalendar = "mayan"
	HeatmapYcalendar_nanakshahi HeatmapYcalendar = "nanakshahi"
	HeatmapYcalendar_nepali     HeatmapYcalendar = "nepali"
	HeatmapYcalendar_persian    HeatmapYcalendar = "persian"
	HeatmapYcalendar_jalali     HeatmapYcalendar = "jalali"
	HeatmapYcalendar_taiwan     HeatmapYcalendar = "taiwan"
	HeatmapYcalendar_thai       HeatmapYcalendar = "thai"
	HeatmapYcalendar_ummalqura  HeatmapYcalendar = "ummalqura"
)

// HeatmapYtype If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
type HeatmapYtype string

const (
	HeatmapYtype_array  HeatmapYtype = "array"
	HeatmapYtype_scaled HeatmapYtype = "scaled"
)

// HeatmapZsmooth Picks a smoothing algorithm use to smooth `z` data.
type HeatmapZsmooth interface{}

var (
	HeatmapZsmooth_fast  HeatmapZsmooth = "fast"
	HeatmapZsmooth_best  HeatmapZsmooth = "best"
	HeatmapZsmooth_False HeatmapZsmooth = false
)

// HeatmapglColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type HeatmapglColorbarExponentformat string

const (
	HeatmapglColorbarExponentformat_none  HeatmapglColorbarExponentformat = "none"
	HeatmapglColorbarExponentformat_e     HeatmapglColorbarExponentformat = "e"
	HeatmapglColorbarExponentformat_E     HeatmapglColorbarExponentformat = "E"
	HeatmapglColorbarExponentformat_power HeatmapglColorbarExponentformat = "power"
	HeatmapglColorbarExponentformat_SI    HeatmapglColorbarExponentformat = "SI"
	HeatmapglColorbarExponentformat_B     HeatmapglColorbarExponentformat = "B"
)

// HeatmapglColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type HeatmapglColorbarLenmode string

const (
	HeatmapglColorbarLenmode_fraction HeatmapglColorbarLenmode = "fraction"
	HeatmapglColorbarLenmode_pixels   HeatmapglColorbarLenmode = "pixels"
)

// HeatmapglColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type HeatmapglColorbarShowexponent string

const (
	HeatmapglColorbarShowexponent_all   HeatmapglColorbarShowexponent = "all"
	HeatmapglColorbarShowexponent_first HeatmapglColorbarShowexponent = "first"
	HeatmapglColorbarShowexponent_last  HeatmapglColorbarShowexponent = "last"
	HeatmapglColorbarShowexponent_none  HeatmapglColorbarShowexponent = "none"
)

// HeatmapglColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type HeatmapglColorbarShowtickprefix string

const (
	HeatmapglColorbarShowtickprefix_all   HeatmapglColorbarShowtickprefix = "all"
	HeatmapglColorbarShowtickprefix_first HeatmapglColorbarShowtickprefix = "first"
	HeatmapglColorbarShowtickprefix_last  HeatmapglColorbarShowtickprefix = "last"
	HeatmapglColorbarShowtickprefix_none  HeatmapglColorbarShowtickprefix = "none"
)

// HeatmapglColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type HeatmapglColorbarShowticksuffix string

const (
	HeatmapglColorbarShowticksuffix_all   HeatmapglColorbarShowticksuffix = "all"
	HeatmapglColorbarShowticksuffix_first HeatmapglColorbarShowticksuffix = "first"
	HeatmapglColorbarShowticksuffix_last  HeatmapglColorbarShowticksuffix = "last"
	HeatmapglColorbarShowticksuffix_none  HeatmapglColorbarShowticksuffix = "none"
)

// HeatmapglColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type HeatmapglColorbarThicknessmode string

const (
	HeatmapglColorbarThicknessmode_fraction HeatmapglColorbarThicknessmode = "fraction"
	HeatmapglColorbarThicknessmode_pixels   HeatmapglColorbarThicknessmode = "pixels"
)

// HeatmapglColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type HeatmapglColorbarTickmode string

const (
	HeatmapglColorbarTickmode_auto   HeatmapglColorbarTickmode = "auto"
	HeatmapglColorbarTickmode_linear HeatmapglColorbarTickmode = "linear"
	HeatmapglColorbarTickmode_array  HeatmapglColorbarTickmode = "array"
)

// HeatmapglColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type HeatmapglColorbarTicks string

const (
	HeatmapglColorbarTicks_outside HeatmapglColorbarTicks = "outside"
	HeatmapglColorbarTicks_inside  HeatmapglColorbarTicks = "inside"
)

// HeatmapglColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type HeatmapglColorbarTitleSide string

const (
	HeatmapglColorbarTitleSide_right  HeatmapglColorbarTitleSide = "right"
	HeatmapglColorbarTitleSide_top    HeatmapglColorbarTitleSide = "top"
	HeatmapglColorbarTitleSide_bottom HeatmapglColorbarTitleSide = "bottom"
)

// HeatmapglColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type HeatmapglColorbarXanchor string

const (
	HeatmapglColorbarXanchor_left   HeatmapglColorbarXanchor = "left"
	HeatmapglColorbarXanchor_center HeatmapglColorbarXanchor = "center"
	HeatmapglColorbarXanchor_right  HeatmapglColorbarXanchor = "right"
)

// HeatmapglColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type HeatmapglColorbarYanchor string

const (
	HeatmapglColorbarYanchor_top    HeatmapglColorbarYanchor = "top"
	HeatmapglColorbarYanchor_middle HeatmapglColorbarYanchor = "middle"
	HeatmapglColorbarYanchor_bottom HeatmapglColorbarYanchor = "bottom"
)

// HeatmapglHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type HeatmapglHoverlabelAlign string

const (
	HeatmapglHoverlabelAlign_left  HeatmapglHoverlabelAlign = "left"
	HeatmapglHoverlabelAlign_right HeatmapglHoverlabelAlign = "right"
	HeatmapglHoverlabelAlign_auto  HeatmapglHoverlabelAlign = "auto"
)

// HeatmapglVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type HeatmapglVisible interface{}

var (
	HeatmapglVisible_True       HeatmapglVisible = true
	HeatmapglVisible_False      HeatmapglVisible = false
	HeatmapglVisible_legendonly HeatmapglVisible = "legendonly"
)

// HeatmapglXtype If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
type HeatmapglXtype string

const (
	HeatmapglXtype_array  HeatmapglXtype = "array"
	HeatmapglXtype_scaled HeatmapglXtype = "scaled"
)

// HeatmapglYtype If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
type HeatmapglYtype string

const (
	HeatmapglYtype_array  HeatmapglYtype = "array"
	HeatmapglYtype_scaled HeatmapglYtype = "scaled"
)

// Histogram2dColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type Histogram2dColorbarExponentformat string

const (
	Histogram2dColorbarExponentformat_none  Histogram2dColorbarExponentformat = "none"
	Histogram2dColorbarExponentformat_e     Histogram2dColorbarExponentformat = "e"
	Histogram2dColorbarExponentformat_E     Histogram2dColorbarExponentformat = "E"
	Histogram2dColorbarExponentformat_power Histogram2dColorbarExponentformat = "power"
	Histogram2dColorbarExponentformat_SI    Histogram2dColorbarExponentformat = "SI"
	Histogram2dColorbarExponentformat_B     Histogram2dColorbarExponentformat = "B"
)

// Histogram2dColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type Histogram2dColorbarLenmode string

const (
	Histogram2dColorbarLenmode_fraction Histogram2dColorbarLenmode = "fraction"
	Histogram2dColorbarLenmode_pixels   Histogram2dColorbarLenmode = "pixels"
)

// Histogram2dColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type Histogram2dColorbarShowexponent string

const (
	Histogram2dColorbarShowexponent_all   Histogram2dColorbarShowexponent = "all"
	Histogram2dColorbarShowexponent_first Histogram2dColorbarShowexponent = "first"
	Histogram2dColorbarShowexponent_last  Histogram2dColorbarShowexponent = "last"
	Histogram2dColorbarShowexponent_none  Histogram2dColorbarShowexponent = "none"
)

// Histogram2dColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type Histogram2dColorbarShowtickprefix string

const (
	Histogram2dColorbarShowtickprefix_all   Histogram2dColorbarShowtickprefix = "all"
	Histogram2dColorbarShowtickprefix_first Histogram2dColorbarShowtickprefix = "first"
	Histogram2dColorbarShowtickprefix_last  Histogram2dColorbarShowtickprefix = "last"
	Histogram2dColorbarShowtickprefix_none  Histogram2dColorbarShowtickprefix = "none"
)

// Histogram2dColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type Histogram2dColorbarShowticksuffix string

const (
	Histogram2dColorbarShowticksuffix_all   Histogram2dColorbarShowticksuffix = "all"
	Histogram2dColorbarShowticksuffix_first Histogram2dColorbarShowticksuffix = "first"
	Histogram2dColorbarShowticksuffix_last  Histogram2dColorbarShowticksuffix = "last"
	Histogram2dColorbarShowticksuffix_none  Histogram2dColorbarShowticksuffix = "none"
)

// Histogram2dColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type Histogram2dColorbarThicknessmode string

const (
	Histogram2dColorbarThicknessmode_fraction Histogram2dColorbarThicknessmode = "fraction"
	Histogram2dColorbarThicknessmode_pixels   Histogram2dColorbarThicknessmode = "pixels"
)

// Histogram2dColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type Histogram2dColorbarTickmode string

const (
	Histogram2dColorbarTickmode_auto   Histogram2dColorbarTickmode = "auto"
	Histogram2dColorbarTickmode_linear Histogram2dColorbarTickmode = "linear"
	Histogram2dColorbarTickmode_array  Histogram2dColorbarTickmode = "array"
)

// Histogram2dColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type Histogram2dColorbarTicks string

const (
	Histogram2dColorbarTicks_outside Histogram2dColorbarTicks = "outside"
	Histogram2dColorbarTicks_inside  Histogram2dColorbarTicks = "inside"
)

// Histogram2dColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type Histogram2dColorbarTitleSide string

const (
	Histogram2dColorbarTitleSide_right  Histogram2dColorbarTitleSide = "right"
	Histogram2dColorbarTitleSide_top    Histogram2dColorbarTitleSide = "top"
	Histogram2dColorbarTitleSide_bottom Histogram2dColorbarTitleSide = "bottom"
)

// Histogram2dColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type Histogram2dColorbarXanchor string

const (
	Histogram2dColorbarXanchor_left   Histogram2dColorbarXanchor = "left"
	Histogram2dColorbarXanchor_center Histogram2dColorbarXanchor = "center"
	Histogram2dColorbarXanchor_right  Histogram2dColorbarXanchor = "right"
)

// Histogram2dColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type Histogram2dColorbarYanchor string

const (
	Histogram2dColorbarYanchor_top    Histogram2dColorbarYanchor = "top"
	Histogram2dColorbarYanchor_middle Histogram2dColorbarYanchor = "middle"
	Histogram2dColorbarYanchor_bottom Histogram2dColorbarYanchor = "bottom"
)

// Histogram2dHistfunc Specifies the binning function used for this histogram trace. If *count*, the histogram values are computed by counting the number of values lying inside each bin. If *sum*, *avg*, *min*, *max*, the histogram values are computed using the sum, the average, the minimum or the maximum of the values lying inside each bin respectively.
type Histogram2dHistfunc string

const (
	Histogram2dHistfunc_count Histogram2dHistfunc = "count"
	Histogram2dHistfunc_sum   Histogram2dHistfunc = "sum"
	Histogram2dHistfunc_avg   Histogram2dHistfunc = "avg"
	Histogram2dHistfunc_min   Histogram2dHistfunc = "min"
	Histogram2dHistfunc_max   Histogram2dHistfunc = "max"
)

// Histogram2dHistnorm Specifies the type of normalization used for this histogram trace. If **, the span of each bar corresponds to the number of occurrences (i.e. the number of data points lying inside the bins). If *percent* / *probability*, the span of each bar corresponds to the percentage / fraction of occurrences with respect to the total number of sample points (here, the sum of all bin HEIGHTS equals 100% / 1). If *density*, the span of each bar corresponds to the number of occurrences in a bin divided by the size of the bin interval (here, the sum of all bin AREAS equals the total number of sample points). If *probability density*, the area of each bar corresponds to the probability that an event will fall into the corresponding bin (here, the sum of all bin AREAS equals 1).
type Histogram2dHistnorm string

const (
	Histogram2dHistnorm_percent            Histogram2dHistnorm = "percent"
	Histogram2dHistnorm_probability        Histogram2dHistnorm = "probability"
	Histogram2dHistnorm_density            Histogram2dHistnorm = "density"
	Histogram2dHistnorm_probabilitydensity Histogram2dHistnorm = "probability density"
)

// Histogram2dHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type Histogram2dHoverlabelAlign string

const (
	Histogram2dHoverlabelAlign_left  Histogram2dHoverlabelAlign = "left"
	Histogram2dHoverlabelAlign_right Histogram2dHoverlabelAlign = "right"
	Histogram2dHoverlabelAlign_auto  Histogram2dHoverlabelAlign = "auto"
)

// Histogram2dVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type Histogram2dVisible interface{}

var (
	Histogram2dVisible_True       Histogram2dVisible = true
	Histogram2dVisible_False      Histogram2dVisible = false
	Histogram2dVisible_legendonly Histogram2dVisible = "legendonly"
)

// Histogram2dXcalendar Sets the calendar system to use with `x` date data.
type Histogram2dXcalendar string

const (
	Histogram2dXcalendar_gregorian  Histogram2dXcalendar = "gregorian"
	Histogram2dXcalendar_chinese    Histogram2dXcalendar = "chinese"
	Histogram2dXcalendar_coptic     Histogram2dXcalendar = "coptic"
	Histogram2dXcalendar_discworld  Histogram2dXcalendar = "discworld"
	Histogram2dXcalendar_ethiopian  Histogram2dXcalendar = "ethiopian"
	Histogram2dXcalendar_hebrew     Histogram2dXcalendar = "hebrew"
	Histogram2dXcalendar_islamic    Histogram2dXcalendar = "islamic"
	Histogram2dXcalendar_julian     Histogram2dXcalendar = "julian"
	Histogram2dXcalendar_mayan      Histogram2dXcalendar = "mayan"
	Histogram2dXcalendar_nanakshahi Histogram2dXcalendar = "nanakshahi"
	Histogram2dXcalendar_nepali     Histogram2dXcalendar = "nepali"
	Histogram2dXcalendar_persian    Histogram2dXcalendar = "persian"
	Histogram2dXcalendar_jalali     Histogram2dXcalendar = "jalali"
	Histogram2dXcalendar_taiwan     Histogram2dXcalendar = "taiwan"
	Histogram2dXcalendar_thai       Histogram2dXcalendar = "thai"
	Histogram2dXcalendar_ummalqura  Histogram2dXcalendar = "ummalqura"
)

// Histogram2dYcalendar Sets the calendar system to use with `y` date data.
type Histogram2dYcalendar string

const (
	Histogram2dYcalendar_gregorian  Histogram2dYcalendar = "gregorian"
	Histogram2dYcalendar_chinese    Histogram2dYcalendar = "chinese"
	Histogram2dYcalendar_coptic     Histogram2dYcalendar = "coptic"
	Histogram2dYcalendar_discworld  Histogram2dYcalendar = "discworld"
	Histogram2dYcalendar_ethiopian  Histogram2dYcalendar = "ethiopian"
	Histogram2dYcalendar_hebrew     Histogram2dYcalendar = "hebrew"
	Histogram2dYcalendar_islamic    Histogram2dYcalendar = "islamic"
	Histogram2dYcalendar_julian     Histogram2dYcalendar = "julian"
	Histogram2dYcalendar_mayan      Histogram2dYcalendar = "mayan"
	Histogram2dYcalendar_nanakshahi Histogram2dYcalendar = "nanakshahi"
	Histogram2dYcalendar_nepali     Histogram2dYcalendar = "nepali"
	Histogram2dYcalendar_persian    Histogram2dYcalendar = "persian"
	Histogram2dYcalendar_jalali     Histogram2dYcalendar = "jalali"
	Histogram2dYcalendar_taiwan     Histogram2dYcalendar = "taiwan"
	Histogram2dYcalendar_thai       Histogram2dYcalendar = "thai"
	Histogram2dYcalendar_ummalqura  Histogram2dYcalendar = "ummalqura"
)

// Histogram2dZsmooth Picks a smoothing algorithm use to smooth `z` data.
type Histogram2dZsmooth interface{}

var (
	Histogram2dZsmooth_fast  Histogram2dZsmooth = "fast"
	Histogram2dZsmooth_best  Histogram2dZsmooth = "best"
	Histogram2dZsmooth_False Histogram2dZsmooth = false
)

// Histogram2dcontourColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type Histogram2dcontourColorbarExponentformat string

const (
	Histogram2dcontourColorbarExponentformat_none  Histogram2dcontourColorbarExponentformat = "none"
	Histogram2dcontourColorbarExponentformat_e     Histogram2dcontourColorbarExponentformat = "e"
	Histogram2dcontourColorbarExponentformat_E     Histogram2dcontourColorbarExponentformat = "E"
	Histogram2dcontourColorbarExponentformat_power Histogram2dcontourColorbarExponentformat = "power"
	Histogram2dcontourColorbarExponentformat_SI    Histogram2dcontourColorbarExponentformat = "SI"
	Histogram2dcontourColorbarExponentformat_B     Histogram2dcontourColorbarExponentformat = "B"
)

// Histogram2dcontourColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type Histogram2dcontourColorbarLenmode string

const (
	Histogram2dcontourColorbarLenmode_fraction Histogram2dcontourColorbarLenmode = "fraction"
	Histogram2dcontourColorbarLenmode_pixels   Histogram2dcontourColorbarLenmode = "pixels"
)

// Histogram2dcontourColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type Histogram2dcontourColorbarShowexponent string

const (
	Histogram2dcontourColorbarShowexponent_all   Histogram2dcontourColorbarShowexponent = "all"
	Histogram2dcontourColorbarShowexponent_first Histogram2dcontourColorbarShowexponent = "first"
	Histogram2dcontourColorbarShowexponent_last  Histogram2dcontourColorbarShowexponent = "last"
	Histogram2dcontourColorbarShowexponent_none  Histogram2dcontourColorbarShowexponent = "none"
)

// Histogram2dcontourColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type Histogram2dcontourColorbarShowtickprefix string

const (
	Histogram2dcontourColorbarShowtickprefix_all   Histogram2dcontourColorbarShowtickprefix = "all"
	Histogram2dcontourColorbarShowtickprefix_first Histogram2dcontourColorbarShowtickprefix = "first"
	Histogram2dcontourColorbarShowtickprefix_last  Histogram2dcontourColorbarShowtickprefix = "last"
	Histogram2dcontourColorbarShowtickprefix_none  Histogram2dcontourColorbarShowtickprefix = "none"
)

// Histogram2dcontourColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type Histogram2dcontourColorbarShowticksuffix string

const (
	Histogram2dcontourColorbarShowticksuffix_all   Histogram2dcontourColorbarShowticksuffix = "all"
	Histogram2dcontourColorbarShowticksuffix_first Histogram2dcontourColorbarShowticksuffix = "first"
	Histogram2dcontourColorbarShowticksuffix_last  Histogram2dcontourColorbarShowticksuffix = "last"
	Histogram2dcontourColorbarShowticksuffix_none  Histogram2dcontourColorbarShowticksuffix = "none"
)

// Histogram2dcontourColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type Histogram2dcontourColorbarThicknessmode string

const (
	Histogram2dcontourColorbarThicknessmode_fraction Histogram2dcontourColorbarThicknessmode = "fraction"
	Histogram2dcontourColorbarThicknessmode_pixels   Histogram2dcontourColorbarThicknessmode = "pixels"
)

// Histogram2dcontourColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type Histogram2dcontourColorbarTickmode string

const (
	Histogram2dcontourColorbarTickmode_auto   Histogram2dcontourColorbarTickmode = "auto"
	Histogram2dcontourColorbarTickmode_linear Histogram2dcontourColorbarTickmode = "linear"
	Histogram2dcontourColorbarTickmode_array  Histogram2dcontourColorbarTickmode = "array"
)

// Histogram2dcontourColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type Histogram2dcontourColorbarTicks string

const (
	Histogram2dcontourColorbarTicks_outside Histogram2dcontourColorbarTicks = "outside"
	Histogram2dcontourColorbarTicks_inside  Histogram2dcontourColorbarTicks = "inside"
)

// Histogram2dcontourColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type Histogram2dcontourColorbarTitleSide string

const (
	Histogram2dcontourColorbarTitleSide_right  Histogram2dcontourColorbarTitleSide = "right"
	Histogram2dcontourColorbarTitleSide_top    Histogram2dcontourColorbarTitleSide = "top"
	Histogram2dcontourColorbarTitleSide_bottom Histogram2dcontourColorbarTitleSide = "bottom"
)

// Histogram2dcontourColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type Histogram2dcontourColorbarXanchor string

const (
	Histogram2dcontourColorbarXanchor_left   Histogram2dcontourColorbarXanchor = "left"
	Histogram2dcontourColorbarXanchor_center Histogram2dcontourColorbarXanchor = "center"
	Histogram2dcontourColorbarXanchor_right  Histogram2dcontourColorbarXanchor = "right"
)

// Histogram2dcontourColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type Histogram2dcontourColorbarYanchor string

const (
	Histogram2dcontourColorbarYanchor_top    Histogram2dcontourColorbarYanchor = "top"
	Histogram2dcontourColorbarYanchor_middle Histogram2dcontourColorbarYanchor = "middle"
	Histogram2dcontourColorbarYanchor_bottom Histogram2dcontourColorbarYanchor = "bottom"
)

// Histogram2dcontourContoursColoring Determines the coloring method showing the contour values. If *fill*, coloring is done evenly between each contour level If *heatmap*, a heatmap gradient coloring is applied between each contour level. If *lines*, coloring is done on the contour lines. If *none*, no coloring is applied on this trace.
type Histogram2dcontourContoursColoring string

const (
	Histogram2dcontourContoursColoring_fill    Histogram2dcontourContoursColoring = "fill"
	Histogram2dcontourContoursColoring_heatmap Histogram2dcontourContoursColoring = "heatmap"
	Histogram2dcontourContoursColoring_lines   Histogram2dcontourContoursColoring = "lines"
	Histogram2dcontourContoursColoring_none    Histogram2dcontourContoursColoring = "none"
)

// Histogram2dcontourContoursOperation Sets the constraint operation. *=* keeps regions equal to `value` *<* and *<=* keep regions less than `value` *>* and *>=* keep regions greater than `value` *[]*, *()*, *[)*, and *(]* keep regions inside `value[0]` to `value[1]` *][*, *)(*, *](*, *)[* keep regions outside `value[0]` to value[1]` Open vs. closed intervals make no difference to constraint display, but all versions are allowed for consistency with filter transforms.
type Histogram2dcontourContoursOperation string

const (
	Histogram2dcontourContoursOperation_eq                   Histogram2dcontourContoursOperation = "="
	Histogram2dcontourContoursOperation_lt                   Histogram2dcontourContoursOperation = "<"
	Histogram2dcontourContoursOperation_gteq                 Histogram2dcontourContoursOperation = ">="
	Histogram2dcontourContoursOperation_gt                   Histogram2dcontourContoursOperation = ">"
	Histogram2dcontourContoursOperation_lteq                 Histogram2dcontourContoursOperation = "<="
	Histogram2dcontourContoursOperation__lbracket__rbracket_ Histogram2dcontourContoursOperation = "[]"
	Histogram2dcontourContoursOperation__lpar__rpar_         Histogram2dcontourContoursOperation = "()"
	Histogram2dcontourContoursOperation__lbracket__rpar_     Histogram2dcontourContoursOperation = "[)"
	Histogram2dcontourContoursOperation__lpar__rbracket_     Histogram2dcontourContoursOperation = "(]"
	Histogram2dcontourContoursOperation__rbracket__lbracket_ Histogram2dcontourContoursOperation = "]["
	Histogram2dcontourContoursOperation__rpar__lpar_         Histogram2dcontourContoursOperation = ")("
	Histogram2dcontourContoursOperation__rbracket__lpar_     Histogram2dcontourContoursOperation = "]("
	Histogram2dcontourContoursOperation__rpar__lbracket_     Histogram2dcontourContoursOperation = ")["
)

// Histogram2dcontourContoursType If `levels`, the data is represented as a contour plot with multiple levels displayed. If `constraint`, the data is represented as constraints with the invalid region shaded as specified by the `operation` and `value` parameters.
type Histogram2dcontourContoursType string

const (
	Histogram2dcontourContoursType_levels     Histogram2dcontourContoursType = "levels"
	Histogram2dcontourContoursType_constraint Histogram2dcontourContoursType = "constraint"
)

// Histogram2dcontourHistfunc Specifies the binning function used for this histogram trace. If *count*, the histogram values are computed by counting the number of values lying inside each bin. If *sum*, *avg*, *min*, *max*, the histogram values are computed using the sum, the average, the minimum or the maximum of the values lying inside each bin respectively.
type Histogram2dcontourHistfunc string

const (
	Histogram2dcontourHistfunc_count Histogram2dcontourHistfunc = "count"
	Histogram2dcontourHistfunc_sum   Histogram2dcontourHistfunc = "sum"
	Histogram2dcontourHistfunc_avg   Histogram2dcontourHistfunc = "avg"
	Histogram2dcontourHistfunc_min   Histogram2dcontourHistfunc = "min"
	Histogram2dcontourHistfunc_max   Histogram2dcontourHistfunc = "max"
)

// Histogram2dcontourHistnorm Specifies the type of normalization used for this histogram trace. If **, the span of each bar corresponds to the number of occurrences (i.e. the number of data points lying inside the bins). If *percent* / *probability*, the span of each bar corresponds to the percentage / fraction of occurrences with respect to the total number of sample points (here, the sum of all bin HEIGHTS equals 100% / 1). If *density*, the span of each bar corresponds to the number of occurrences in a bin divided by the size of the bin interval (here, the sum of all bin AREAS equals the total number of sample points). If *probability density*, the area of each bar corresponds to the probability that an event will fall into the corresponding bin (here, the sum of all bin AREAS equals 1).
type Histogram2dcontourHistnorm string

const (
	Histogram2dcontourHistnorm_percent            Histogram2dcontourHistnorm = "percent"
	Histogram2dcontourHistnorm_probability        Histogram2dcontourHistnorm = "probability"
	Histogram2dcontourHistnorm_density            Histogram2dcontourHistnorm = "density"
	Histogram2dcontourHistnorm_probabilitydensity Histogram2dcontourHistnorm = "probability density"
)

// Histogram2dcontourHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type Histogram2dcontourHoverlabelAlign string

const (
	Histogram2dcontourHoverlabelAlign_left  Histogram2dcontourHoverlabelAlign = "left"
	Histogram2dcontourHoverlabelAlign_right Histogram2dcontourHoverlabelAlign = "right"
	Histogram2dcontourHoverlabelAlign_auto  Histogram2dcontourHoverlabelAlign = "auto"
)

// Histogram2dcontourVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type Histogram2dcontourVisible interface{}

var (
	Histogram2dcontourVisible_True       Histogram2dcontourVisible = true
	Histogram2dcontourVisible_False      Histogram2dcontourVisible = false
	Histogram2dcontourVisible_legendonly Histogram2dcontourVisible = "legendonly"
)

// Histogram2dcontourXcalendar Sets the calendar system to use with `x` date data.
type Histogram2dcontourXcalendar string

const (
	Histogram2dcontourXcalendar_gregorian  Histogram2dcontourXcalendar = "gregorian"
	Histogram2dcontourXcalendar_chinese    Histogram2dcontourXcalendar = "chinese"
	Histogram2dcontourXcalendar_coptic     Histogram2dcontourXcalendar = "coptic"
	Histogram2dcontourXcalendar_discworld  Histogram2dcontourXcalendar = "discworld"
	Histogram2dcontourXcalendar_ethiopian  Histogram2dcontourXcalendar = "ethiopian"
	Histogram2dcontourXcalendar_hebrew     Histogram2dcontourXcalendar = "hebrew"
	Histogram2dcontourXcalendar_islamic    Histogram2dcontourXcalendar = "islamic"
	Histogram2dcontourXcalendar_julian     Histogram2dcontourXcalendar = "julian"
	Histogram2dcontourXcalendar_mayan      Histogram2dcontourXcalendar = "mayan"
	Histogram2dcontourXcalendar_nanakshahi Histogram2dcontourXcalendar = "nanakshahi"
	Histogram2dcontourXcalendar_nepali     Histogram2dcontourXcalendar = "nepali"
	Histogram2dcontourXcalendar_persian    Histogram2dcontourXcalendar = "persian"
	Histogram2dcontourXcalendar_jalali     Histogram2dcontourXcalendar = "jalali"
	Histogram2dcontourXcalendar_taiwan     Histogram2dcontourXcalendar = "taiwan"
	Histogram2dcontourXcalendar_thai       Histogram2dcontourXcalendar = "thai"
	Histogram2dcontourXcalendar_ummalqura  Histogram2dcontourXcalendar = "ummalqura"
)

// Histogram2dcontourYcalendar Sets the calendar system to use with `y` date data.
type Histogram2dcontourYcalendar string

const (
	Histogram2dcontourYcalendar_gregorian  Histogram2dcontourYcalendar = "gregorian"
	Histogram2dcontourYcalendar_chinese    Histogram2dcontourYcalendar = "chinese"
	Histogram2dcontourYcalendar_coptic     Histogram2dcontourYcalendar = "coptic"
	Histogram2dcontourYcalendar_discworld  Histogram2dcontourYcalendar = "discworld"
	Histogram2dcontourYcalendar_ethiopian  Histogram2dcontourYcalendar = "ethiopian"
	Histogram2dcontourYcalendar_hebrew     Histogram2dcontourYcalendar = "hebrew"
	Histogram2dcontourYcalendar_islamic    Histogram2dcontourYcalendar = "islamic"
	Histogram2dcontourYcalendar_julian     Histogram2dcontourYcalendar = "julian"
	Histogram2dcontourYcalendar_mayan      Histogram2dcontourYcalendar = "mayan"
	Histogram2dcontourYcalendar_nanakshahi Histogram2dcontourYcalendar = "nanakshahi"
	Histogram2dcontourYcalendar_nepali     Histogram2dcontourYcalendar = "nepali"
	Histogram2dcontourYcalendar_persian    Histogram2dcontourYcalendar = "persian"
	Histogram2dcontourYcalendar_jalali     Histogram2dcontourYcalendar = "jalali"
	Histogram2dcontourYcalendar_taiwan     Histogram2dcontourYcalendar = "taiwan"
	Histogram2dcontourYcalendar_thai       Histogram2dcontourYcalendar = "thai"
	Histogram2dcontourYcalendar_ummalqura  Histogram2dcontourYcalendar = "ummalqura"
)

// HistogramCumulativeCurrentbin Only applies if cumulative is enabled. Sets whether the current bin is included, excluded, or has half of its value included in the current cumulative value. *include* is the default for compatibility with various other tools, however it introduces a half-bin bias to the results. *exclude* makes the opposite half-bin bias, and *half* removes it.
type HistogramCumulativeCurrentbin string

const (
	HistogramCumulativeCurrentbin_include HistogramCumulativeCurrentbin = "include"
	HistogramCumulativeCurrentbin_exclude HistogramCumulativeCurrentbin = "exclude"
	HistogramCumulativeCurrentbin_half    HistogramCumulativeCurrentbin = "half"
)

// HistogramCumulativeDirection Only applies if cumulative is enabled. If *increasing* (default) we sum all prior bins, so the result increases from left to right. If *decreasing* we sum later bins so the result decreases from left to right.
type HistogramCumulativeDirection string

const (
	HistogramCumulativeDirection_increasing HistogramCumulativeDirection = "increasing"
	HistogramCumulativeDirection_decreasing HistogramCumulativeDirection = "decreasing"
)

// HistogramErrorXType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type HistogramErrorXType string

const (
	HistogramErrorXType_percent  HistogramErrorXType = "percent"
	HistogramErrorXType_constant HistogramErrorXType = "constant"
	HistogramErrorXType_sqrt     HistogramErrorXType = "sqrt"
	HistogramErrorXType_data     HistogramErrorXType = "data"
)

// HistogramErrorYType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type HistogramErrorYType string

const (
	HistogramErrorYType_percent  HistogramErrorYType = "percent"
	HistogramErrorYType_constant HistogramErrorYType = "constant"
	HistogramErrorYType_sqrt     HistogramErrorYType = "sqrt"
	HistogramErrorYType_data     HistogramErrorYType = "data"
)

// HistogramHistfunc Specifies the binning function used for this histogram trace. If *count*, the histogram values are computed by counting the number of values lying inside each bin. If *sum*, *avg*, *min*, *max*, the histogram values are computed using the sum, the average, the minimum or the maximum of the values lying inside each bin respectively.
type HistogramHistfunc string

const (
	HistogramHistfunc_count HistogramHistfunc = "count"
	HistogramHistfunc_sum   HistogramHistfunc = "sum"
	HistogramHistfunc_avg   HistogramHistfunc = "avg"
	HistogramHistfunc_min   HistogramHistfunc = "min"
	HistogramHistfunc_max   HistogramHistfunc = "max"
)

// HistogramHistnorm Specifies the type of normalization used for this histogram trace. If **, the span of each bar corresponds to the number of occurrences (i.e. the number of data points lying inside the bins). If *percent* / *probability*, the span of each bar corresponds to the percentage / fraction of occurrences with respect to the total number of sample points (here, the sum of all bin HEIGHTS equals 100% / 1). If *density*, the span of each bar corresponds to the number of occurrences in a bin divided by the size of the bin interval (here, the sum of all bin AREAS equals the total number of sample points). If *probability density*, the area of each bar corresponds to the probability that an event will fall into the corresponding bin (here, the sum of all bin AREAS equals 1).
type HistogramHistnorm string

const (
	HistogramHistnorm_percent            HistogramHistnorm = "percent"
	HistogramHistnorm_probability        HistogramHistnorm = "probability"
	HistogramHistnorm_density            HistogramHistnorm = "density"
	HistogramHistnorm_probabilitydensity HistogramHistnorm = "probability density"
)

// HistogramHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type HistogramHoverlabelAlign string

const (
	HistogramHoverlabelAlign_left  HistogramHoverlabelAlign = "left"
	HistogramHoverlabelAlign_right HistogramHoverlabelAlign = "right"
	HistogramHoverlabelAlign_auto  HistogramHoverlabelAlign = "auto"
)

// HistogramMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type HistogramMarkerColorbarExponentformat string

const (
	HistogramMarkerColorbarExponentformat_none  HistogramMarkerColorbarExponentformat = "none"
	HistogramMarkerColorbarExponentformat_e     HistogramMarkerColorbarExponentformat = "e"
	HistogramMarkerColorbarExponentformat_E     HistogramMarkerColorbarExponentformat = "E"
	HistogramMarkerColorbarExponentformat_power HistogramMarkerColorbarExponentformat = "power"
	HistogramMarkerColorbarExponentformat_SI    HistogramMarkerColorbarExponentformat = "SI"
	HistogramMarkerColorbarExponentformat_B     HistogramMarkerColorbarExponentformat = "B"
)

// HistogramMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type HistogramMarkerColorbarLenmode string

const (
	HistogramMarkerColorbarLenmode_fraction HistogramMarkerColorbarLenmode = "fraction"
	HistogramMarkerColorbarLenmode_pixels   HistogramMarkerColorbarLenmode = "pixels"
)

// HistogramMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type HistogramMarkerColorbarShowexponent string

const (
	HistogramMarkerColorbarShowexponent_all   HistogramMarkerColorbarShowexponent = "all"
	HistogramMarkerColorbarShowexponent_first HistogramMarkerColorbarShowexponent = "first"
	HistogramMarkerColorbarShowexponent_last  HistogramMarkerColorbarShowexponent = "last"
	HistogramMarkerColorbarShowexponent_none  HistogramMarkerColorbarShowexponent = "none"
)

// HistogramMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type HistogramMarkerColorbarShowtickprefix string

const (
	HistogramMarkerColorbarShowtickprefix_all   HistogramMarkerColorbarShowtickprefix = "all"
	HistogramMarkerColorbarShowtickprefix_first HistogramMarkerColorbarShowtickprefix = "first"
	HistogramMarkerColorbarShowtickprefix_last  HistogramMarkerColorbarShowtickprefix = "last"
	HistogramMarkerColorbarShowtickprefix_none  HistogramMarkerColorbarShowtickprefix = "none"
)

// HistogramMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type HistogramMarkerColorbarShowticksuffix string

const (
	HistogramMarkerColorbarShowticksuffix_all   HistogramMarkerColorbarShowticksuffix = "all"
	HistogramMarkerColorbarShowticksuffix_first HistogramMarkerColorbarShowticksuffix = "first"
	HistogramMarkerColorbarShowticksuffix_last  HistogramMarkerColorbarShowticksuffix = "last"
	HistogramMarkerColorbarShowticksuffix_none  HistogramMarkerColorbarShowticksuffix = "none"
)

// HistogramMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type HistogramMarkerColorbarThicknessmode string

const (
	HistogramMarkerColorbarThicknessmode_fraction HistogramMarkerColorbarThicknessmode = "fraction"
	HistogramMarkerColorbarThicknessmode_pixels   HistogramMarkerColorbarThicknessmode = "pixels"
)

// HistogramMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type HistogramMarkerColorbarTickmode string

const (
	HistogramMarkerColorbarTickmode_auto   HistogramMarkerColorbarTickmode = "auto"
	HistogramMarkerColorbarTickmode_linear HistogramMarkerColorbarTickmode = "linear"
	HistogramMarkerColorbarTickmode_array  HistogramMarkerColorbarTickmode = "array"
)

// HistogramMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type HistogramMarkerColorbarTicks string

const (
	HistogramMarkerColorbarTicks_outside HistogramMarkerColorbarTicks = "outside"
	HistogramMarkerColorbarTicks_inside  HistogramMarkerColorbarTicks = "inside"
)

// HistogramMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type HistogramMarkerColorbarTitleSide string

const (
	HistogramMarkerColorbarTitleSide_right  HistogramMarkerColorbarTitleSide = "right"
	HistogramMarkerColorbarTitleSide_top    HistogramMarkerColorbarTitleSide = "top"
	HistogramMarkerColorbarTitleSide_bottom HistogramMarkerColorbarTitleSide = "bottom"
)

// HistogramMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type HistogramMarkerColorbarXanchor string

const (
	HistogramMarkerColorbarXanchor_left   HistogramMarkerColorbarXanchor = "left"
	HistogramMarkerColorbarXanchor_center HistogramMarkerColorbarXanchor = "center"
	HistogramMarkerColorbarXanchor_right  HistogramMarkerColorbarXanchor = "right"
)

// HistogramMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type HistogramMarkerColorbarYanchor string

const (
	HistogramMarkerColorbarYanchor_top    HistogramMarkerColorbarYanchor = "top"
	HistogramMarkerColorbarYanchor_middle HistogramMarkerColorbarYanchor = "middle"
	HistogramMarkerColorbarYanchor_bottom HistogramMarkerColorbarYanchor = "bottom"
)

// HistogramOrientation Sets the orientation of the bars. With *v* (*h*), the value of the each bar spans along the vertical (horizontal).
type HistogramOrientation string

const (
	HistogramOrientation_v HistogramOrientation = "v"
	HistogramOrientation_h HistogramOrientation = "h"
)

// HistogramVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type HistogramVisible interface{}

var (
	HistogramVisible_True       HistogramVisible = true
	HistogramVisible_False      HistogramVisible = false
	HistogramVisible_legendonly HistogramVisible = "legendonly"
)

// HistogramXcalendar Sets the calendar system to use with `x` date data.
type HistogramXcalendar string

const (
	HistogramXcalendar_gregorian  HistogramXcalendar = "gregorian"
	HistogramXcalendar_chinese    HistogramXcalendar = "chinese"
	HistogramXcalendar_coptic     HistogramXcalendar = "coptic"
	HistogramXcalendar_discworld  HistogramXcalendar = "discworld"
	HistogramXcalendar_ethiopian  HistogramXcalendar = "ethiopian"
	HistogramXcalendar_hebrew     HistogramXcalendar = "hebrew"
	HistogramXcalendar_islamic    HistogramXcalendar = "islamic"
	HistogramXcalendar_julian     HistogramXcalendar = "julian"
	HistogramXcalendar_mayan      HistogramXcalendar = "mayan"
	HistogramXcalendar_nanakshahi HistogramXcalendar = "nanakshahi"
	HistogramXcalendar_nepali     HistogramXcalendar = "nepali"
	HistogramXcalendar_persian    HistogramXcalendar = "persian"
	HistogramXcalendar_jalali     HistogramXcalendar = "jalali"
	HistogramXcalendar_taiwan     HistogramXcalendar = "taiwan"
	HistogramXcalendar_thai       HistogramXcalendar = "thai"
	HistogramXcalendar_ummalqura  HistogramXcalendar = "ummalqura"
)

// HistogramYcalendar Sets the calendar system to use with `y` date data.
type HistogramYcalendar string

const (
	HistogramYcalendar_gregorian  HistogramYcalendar = "gregorian"
	HistogramYcalendar_chinese    HistogramYcalendar = "chinese"
	HistogramYcalendar_coptic     HistogramYcalendar = "coptic"
	HistogramYcalendar_discworld  HistogramYcalendar = "discworld"
	HistogramYcalendar_ethiopian  HistogramYcalendar = "ethiopian"
	HistogramYcalendar_hebrew     HistogramYcalendar = "hebrew"
	HistogramYcalendar_islamic    HistogramYcalendar = "islamic"
	HistogramYcalendar_julian     HistogramYcalendar = "julian"
	HistogramYcalendar_mayan      HistogramYcalendar = "mayan"
	HistogramYcalendar_nanakshahi HistogramYcalendar = "nanakshahi"
	HistogramYcalendar_nepali     HistogramYcalendar = "nepali"
	HistogramYcalendar_persian    HistogramYcalendar = "persian"
	HistogramYcalendar_jalali     HistogramYcalendar = "jalali"
	HistogramYcalendar_taiwan     HistogramYcalendar = "taiwan"
	HistogramYcalendar_thai       HistogramYcalendar = "thai"
	HistogramYcalendar_ummalqura  HistogramYcalendar = "ummalqura"
)

// ImageColormodel Color model used to map the numerical color components described in `z` into colors.
type ImageColormodel string

const (
	ImageColormodel_rgb  ImageColormodel = "rgb"
	ImageColormodel_rgba ImageColormodel = "rgba"
	ImageColormodel_hsl  ImageColormodel = "hsl"
	ImageColormodel_hsla ImageColormodel = "hsla"
)

// ImageHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ImageHoverlabelAlign string

const (
	ImageHoverlabelAlign_left  ImageHoverlabelAlign = "left"
	ImageHoverlabelAlign_right ImageHoverlabelAlign = "right"
	ImageHoverlabelAlign_auto  ImageHoverlabelAlign = "auto"
)

// ImageVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ImageVisible interface{}

var (
	ImageVisible_True       ImageVisible = true
	ImageVisible_False      ImageVisible = false
	ImageVisible_legendonly ImageVisible = "legendonly"
)

// IndicatorAlign Sets the horizontal alignment of the `text` within the box. Note that this attribute has no effect if an angular gauge is displayed: in this case, it is always centered
type IndicatorAlign string

const (
	IndicatorAlign_left   IndicatorAlign = "left"
	IndicatorAlign_center IndicatorAlign = "center"
	IndicatorAlign_right  IndicatorAlign = "right"
)

// IndicatorDeltaPosition Sets the position of delta with respect to the number.
type IndicatorDeltaPosition string

const (
	IndicatorDeltaPosition_top    IndicatorDeltaPosition = "top"
	IndicatorDeltaPosition_bottom IndicatorDeltaPosition = "bottom"
	IndicatorDeltaPosition_left   IndicatorDeltaPosition = "left"
	IndicatorDeltaPosition_right  IndicatorDeltaPosition = "right"
)

// IndicatorGaugeAxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type IndicatorGaugeAxisExponentformat string

const (
	IndicatorGaugeAxisExponentformat_none  IndicatorGaugeAxisExponentformat = "none"
	IndicatorGaugeAxisExponentformat_e     IndicatorGaugeAxisExponentformat = "e"
	IndicatorGaugeAxisExponentformat_E     IndicatorGaugeAxisExponentformat = "E"
	IndicatorGaugeAxisExponentformat_power IndicatorGaugeAxisExponentformat = "power"
	IndicatorGaugeAxisExponentformat_SI    IndicatorGaugeAxisExponentformat = "SI"
	IndicatorGaugeAxisExponentformat_B     IndicatorGaugeAxisExponentformat = "B"
)

// IndicatorGaugeAxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type IndicatorGaugeAxisShowexponent string

const (
	IndicatorGaugeAxisShowexponent_all   IndicatorGaugeAxisShowexponent = "all"
	IndicatorGaugeAxisShowexponent_first IndicatorGaugeAxisShowexponent = "first"
	IndicatorGaugeAxisShowexponent_last  IndicatorGaugeAxisShowexponent = "last"
	IndicatorGaugeAxisShowexponent_none  IndicatorGaugeAxisShowexponent = "none"
)

// IndicatorGaugeAxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type IndicatorGaugeAxisShowtickprefix string

const (
	IndicatorGaugeAxisShowtickprefix_all   IndicatorGaugeAxisShowtickprefix = "all"
	IndicatorGaugeAxisShowtickprefix_first IndicatorGaugeAxisShowtickprefix = "first"
	IndicatorGaugeAxisShowtickprefix_last  IndicatorGaugeAxisShowtickprefix = "last"
	IndicatorGaugeAxisShowtickprefix_none  IndicatorGaugeAxisShowtickprefix = "none"
)

// IndicatorGaugeAxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type IndicatorGaugeAxisShowticksuffix string

const (
	IndicatorGaugeAxisShowticksuffix_all   IndicatorGaugeAxisShowticksuffix = "all"
	IndicatorGaugeAxisShowticksuffix_first IndicatorGaugeAxisShowticksuffix = "first"
	IndicatorGaugeAxisShowticksuffix_last  IndicatorGaugeAxisShowticksuffix = "last"
	IndicatorGaugeAxisShowticksuffix_none  IndicatorGaugeAxisShowticksuffix = "none"
)

// IndicatorGaugeAxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type IndicatorGaugeAxisTickmode string

const (
	IndicatorGaugeAxisTickmode_auto   IndicatorGaugeAxisTickmode = "auto"
	IndicatorGaugeAxisTickmode_linear IndicatorGaugeAxisTickmode = "linear"
	IndicatorGaugeAxisTickmode_array  IndicatorGaugeAxisTickmode = "array"
)

// IndicatorGaugeAxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type IndicatorGaugeAxisTicks string

const (
	IndicatorGaugeAxisTicks_outside IndicatorGaugeAxisTicks = "outside"
	IndicatorGaugeAxisTicks_inside  IndicatorGaugeAxisTicks = "inside"
)

// IndicatorGaugeShape Set the shape of the gauge
type IndicatorGaugeShape string

const (
	IndicatorGaugeShape_angular IndicatorGaugeShape = "angular"
	IndicatorGaugeShape_bullet  IndicatorGaugeShape = "bullet"
)

// IndicatorTitleAlign Sets the horizontal alignment of the title. It defaults to `center` except for bullet charts for which it defaults to right.
type IndicatorTitleAlign string

const (
	IndicatorTitleAlign_left   IndicatorTitleAlign = "left"
	IndicatorTitleAlign_center IndicatorTitleAlign = "center"
	IndicatorTitleAlign_right  IndicatorTitleAlign = "right"
)

// IndicatorVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type IndicatorVisible interface{}

var (
	IndicatorVisible_True       IndicatorVisible = true
	IndicatorVisible_False      IndicatorVisible = false
	IndicatorVisible_legendonly IndicatorVisible = "legendonly"
)

// IsosurfaceColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type IsosurfaceColorbarExponentformat string

const (
	IsosurfaceColorbarExponentformat_none  IsosurfaceColorbarExponentformat = "none"
	IsosurfaceColorbarExponentformat_e     IsosurfaceColorbarExponentformat = "e"
	IsosurfaceColorbarExponentformat_E     IsosurfaceColorbarExponentformat = "E"
	IsosurfaceColorbarExponentformat_power IsosurfaceColorbarExponentformat = "power"
	IsosurfaceColorbarExponentformat_SI    IsosurfaceColorbarExponentformat = "SI"
	IsosurfaceColorbarExponentformat_B     IsosurfaceColorbarExponentformat = "B"
)

// IsosurfaceColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type IsosurfaceColorbarLenmode string

const (
	IsosurfaceColorbarLenmode_fraction IsosurfaceColorbarLenmode = "fraction"
	IsosurfaceColorbarLenmode_pixels   IsosurfaceColorbarLenmode = "pixels"
)

// IsosurfaceColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type IsosurfaceColorbarShowexponent string

const (
	IsosurfaceColorbarShowexponent_all   IsosurfaceColorbarShowexponent = "all"
	IsosurfaceColorbarShowexponent_first IsosurfaceColorbarShowexponent = "first"
	IsosurfaceColorbarShowexponent_last  IsosurfaceColorbarShowexponent = "last"
	IsosurfaceColorbarShowexponent_none  IsosurfaceColorbarShowexponent = "none"
)

// IsosurfaceColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type IsosurfaceColorbarShowtickprefix string

const (
	IsosurfaceColorbarShowtickprefix_all   IsosurfaceColorbarShowtickprefix = "all"
	IsosurfaceColorbarShowtickprefix_first IsosurfaceColorbarShowtickprefix = "first"
	IsosurfaceColorbarShowtickprefix_last  IsosurfaceColorbarShowtickprefix = "last"
	IsosurfaceColorbarShowtickprefix_none  IsosurfaceColorbarShowtickprefix = "none"
)

// IsosurfaceColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type IsosurfaceColorbarShowticksuffix string

const (
	IsosurfaceColorbarShowticksuffix_all   IsosurfaceColorbarShowticksuffix = "all"
	IsosurfaceColorbarShowticksuffix_first IsosurfaceColorbarShowticksuffix = "first"
	IsosurfaceColorbarShowticksuffix_last  IsosurfaceColorbarShowticksuffix = "last"
	IsosurfaceColorbarShowticksuffix_none  IsosurfaceColorbarShowticksuffix = "none"
)

// IsosurfaceColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type IsosurfaceColorbarThicknessmode string

const (
	IsosurfaceColorbarThicknessmode_fraction IsosurfaceColorbarThicknessmode = "fraction"
	IsosurfaceColorbarThicknessmode_pixels   IsosurfaceColorbarThicknessmode = "pixels"
)

// IsosurfaceColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type IsosurfaceColorbarTickmode string

const (
	IsosurfaceColorbarTickmode_auto   IsosurfaceColorbarTickmode = "auto"
	IsosurfaceColorbarTickmode_linear IsosurfaceColorbarTickmode = "linear"
	IsosurfaceColorbarTickmode_array  IsosurfaceColorbarTickmode = "array"
)

// IsosurfaceColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type IsosurfaceColorbarTicks string

const (
	IsosurfaceColorbarTicks_outside IsosurfaceColorbarTicks = "outside"
	IsosurfaceColorbarTicks_inside  IsosurfaceColorbarTicks = "inside"
)

// IsosurfaceColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type IsosurfaceColorbarTitleSide string

const (
	IsosurfaceColorbarTitleSide_right  IsosurfaceColorbarTitleSide = "right"
	IsosurfaceColorbarTitleSide_top    IsosurfaceColorbarTitleSide = "top"
	IsosurfaceColorbarTitleSide_bottom IsosurfaceColorbarTitleSide = "bottom"
)

// IsosurfaceColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type IsosurfaceColorbarXanchor string

const (
	IsosurfaceColorbarXanchor_left   IsosurfaceColorbarXanchor = "left"
	IsosurfaceColorbarXanchor_center IsosurfaceColorbarXanchor = "center"
	IsosurfaceColorbarXanchor_right  IsosurfaceColorbarXanchor = "right"
)

// IsosurfaceColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type IsosurfaceColorbarYanchor string

const (
	IsosurfaceColorbarYanchor_top    IsosurfaceColorbarYanchor = "top"
	IsosurfaceColorbarYanchor_middle IsosurfaceColorbarYanchor = "middle"
	IsosurfaceColorbarYanchor_bottom IsosurfaceColorbarYanchor = "bottom"
)

// IsosurfaceHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type IsosurfaceHoverlabelAlign string

const (
	IsosurfaceHoverlabelAlign_left  IsosurfaceHoverlabelAlign = "left"
	IsosurfaceHoverlabelAlign_right IsosurfaceHoverlabelAlign = "right"
	IsosurfaceHoverlabelAlign_auto  IsosurfaceHoverlabelAlign = "auto"
)

// IsosurfaceVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type IsosurfaceVisible interface{}

var (
	IsosurfaceVisible_True       IsosurfaceVisible = true
	IsosurfaceVisible_False      IsosurfaceVisible = false
	IsosurfaceVisible_legendonly IsosurfaceVisible = "legendonly"
)

// LayoutAngularaxisTickorientation Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the orientation (from the paper perspective) of the angular axis tick labels.
type LayoutAngularaxisTickorientation string

const (
	LayoutAngularaxisTickorientation_horizontal LayoutAngularaxisTickorientation = "horizontal"
	LayoutAngularaxisTickorientation_vertical   LayoutAngularaxisTickorientation = "vertical"
)

// LayoutBarmode Determines how bars at the same location coordinate are displayed on the graph. With *stack*, the bars are stacked on top of one another With *relative*, the bars are stacked on top of one another, with negative values below the axis, positive values above With *group*, the bars are plotted next to one another centered around the shared location. With *overlay*, the bars are plotted over one another, you might need to an *opacity* to see multiple bars.
type LayoutBarmode string

const (
	LayoutBarmode_stack    LayoutBarmode = "stack"
	LayoutBarmode_group    LayoutBarmode = "group"
	LayoutBarmode_overlay  LayoutBarmode = "overlay"
	LayoutBarmode_relative LayoutBarmode = "relative"
)

// LayoutBarnorm Sets the normalization for bar traces on the graph. With *fraction*, the value of each bar is divided by the sum of all values at that location coordinate. *percent* is the same but multiplied by 100 to show percentages.
type LayoutBarnorm string

const (
	LayoutBarnorm_fraction LayoutBarnorm = "fraction"
	LayoutBarnorm_percent  LayoutBarnorm = "percent"
)

// LayoutBoxmode Determines how boxes at the same location coordinate are displayed on the graph. If *group*, the boxes are plotted next to one another centered around the shared location. If *overlay*, the boxes are plotted over one another, you might need to set *opacity* to see them multiple boxes. Has no effect on traces that have *width* set.
type LayoutBoxmode string

const (
	LayoutBoxmode_group   LayoutBoxmode = "group"
	LayoutBoxmode_overlay LayoutBoxmode = "overlay"
)

// LayoutCalendar Sets the default calendar system to use for interpreting and displaying dates throughout the plot.
type LayoutCalendar string

const (
	LayoutCalendar_gregorian  LayoutCalendar = "gregorian"
	LayoutCalendar_chinese    LayoutCalendar = "chinese"
	LayoutCalendar_coptic     LayoutCalendar = "coptic"
	LayoutCalendar_discworld  LayoutCalendar = "discworld"
	LayoutCalendar_ethiopian  LayoutCalendar = "ethiopian"
	LayoutCalendar_hebrew     LayoutCalendar = "hebrew"
	LayoutCalendar_islamic    LayoutCalendar = "islamic"
	LayoutCalendar_julian     LayoutCalendar = "julian"
	LayoutCalendar_mayan      LayoutCalendar = "mayan"
	LayoutCalendar_nanakshahi LayoutCalendar = "nanakshahi"
	LayoutCalendar_nepali     LayoutCalendar = "nepali"
	LayoutCalendar_persian    LayoutCalendar = "persian"
	LayoutCalendar_jalali     LayoutCalendar = "jalali"
	LayoutCalendar_taiwan     LayoutCalendar = "taiwan"
	LayoutCalendar_thai       LayoutCalendar = "thai"
	LayoutCalendar_ummalqura  LayoutCalendar = "ummalqura"
)

// LayoutColoraxisColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutColoraxisColorbarExponentformat string

const (
	LayoutColoraxisColorbarExponentformat_none  LayoutColoraxisColorbarExponentformat = "none"
	LayoutColoraxisColorbarExponentformat_e     LayoutColoraxisColorbarExponentformat = "e"
	LayoutColoraxisColorbarExponentformat_E     LayoutColoraxisColorbarExponentformat = "E"
	LayoutColoraxisColorbarExponentformat_power LayoutColoraxisColorbarExponentformat = "power"
	LayoutColoraxisColorbarExponentformat_SI    LayoutColoraxisColorbarExponentformat = "SI"
	LayoutColoraxisColorbarExponentformat_B     LayoutColoraxisColorbarExponentformat = "B"
)

// LayoutColoraxisColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type LayoutColoraxisColorbarLenmode string

const (
	LayoutColoraxisColorbarLenmode_fraction LayoutColoraxisColorbarLenmode = "fraction"
	LayoutColoraxisColorbarLenmode_pixels   LayoutColoraxisColorbarLenmode = "pixels"
)

// LayoutColoraxisColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutColoraxisColorbarShowexponent string

const (
	LayoutColoraxisColorbarShowexponent_all   LayoutColoraxisColorbarShowexponent = "all"
	LayoutColoraxisColorbarShowexponent_first LayoutColoraxisColorbarShowexponent = "first"
	LayoutColoraxisColorbarShowexponent_last  LayoutColoraxisColorbarShowexponent = "last"
	LayoutColoraxisColorbarShowexponent_none  LayoutColoraxisColorbarShowexponent = "none"
)

// LayoutColoraxisColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutColoraxisColorbarShowtickprefix string

const (
	LayoutColoraxisColorbarShowtickprefix_all   LayoutColoraxisColorbarShowtickprefix = "all"
	LayoutColoraxisColorbarShowtickprefix_first LayoutColoraxisColorbarShowtickprefix = "first"
	LayoutColoraxisColorbarShowtickprefix_last  LayoutColoraxisColorbarShowtickprefix = "last"
	LayoutColoraxisColorbarShowtickprefix_none  LayoutColoraxisColorbarShowtickprefix = "none"
)

// LayoutColoraxisColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutColoraxisColorbarShowticksuffix string

const (
	LayoutColoraxisColorbarShowticksuffix_all   LayoutColoraxisColorbarShowticksuffix = "all"
	LayoutColoraxisColorbarShowticksuffix_first LayoutColoraxisColorbarShowticksuffix = "first"
	LayoutColoraxisColorbarShowticksuffix_last  LayoutColoraxisColorbarShowticksuffix = "last"
	LayoutColoraxisColorbarShowticksuffix_none  LayoutColoraxisColorbarShowticksuffix = "none"
)

// LayoutColoraxisColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type LayoutColoraxisColorbarThicknessmode string

const (
	LayoutColoraxisColorbarThicknessmode_fraction LayoutColoraxisColorbarThicknessmode = "fraction"
	LayoutColoraxisColorbarThicknessmode_pixels   LayoutColoraxisColorbarThicknessmode = "pixels"
)

// LayoutColoraxisColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutColoraxisColorbarTickmode string

const (
	LayoutColoraxisColorbarTickmode_auto   LayoutColoraxisColorbarTickmode = "auto"
	LayoutColoraxisColorbarTickmode_linear LayoutColoraxisColorbarTickmode = "linear"
	LayoutColoraxisColorbarTickmode_array  LayoutColoraxisColorbarTickmode = "array"
)

// LayoutColoraxisColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutColoraxisColorbarTicks string

const (
	LayoutColoraxisColorbarTicks_outside LayoutColoraxisColorbarTicks = "outside"
	LayoutColoraxisColorbarTicks_inside  LayoutColoraxisColorbarTicks = "inside"
)

// LayoutColoraxisColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type LayoutColoraxisColorbarTitleSide string

const (
	LayoutColoraxisColorbarTitleSide_right  LayoutColoraxisColorbarTitleSide = "right"
	LayoutColoraxisColorbarTitleSide_top    LayoutColoraxisColorbarTitleSide = "top"
	LayoutColoraxisColorbarTitleSide_bottom LayoutColoraxisColorbarTitleSide = "bottom"
)

// LayoutColoraxisColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type LayoutColoraxisColorbarXanchor string

const (
	LayoutColoraxisColorbarXanchor_left   LayoutColoraxisColorbarXanchor = "left"
	LayoutColoraxisColorbarXanchor_center LayoutColoraxisColorbarXanchor = "center"
	LayoutColoraxisColorbarXanchor_right  LayoutColoraxisColorbarXanchor = "right"
)

// LayoutColoraxisColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type LayoutColoraxisColorbarYanchor string

const (
	LayoutColoraxisColorbarYanchor_top    LayoutColoraxisColorbarYanchor = "top"
	LayoutColoraxisColorbarYanchor_middle LayoutColoraxisColorbarYanchor = "middle"
	LayoutColoraxisColorbarYanchor_bottom LayoutColoraxisColorbarYanchor = "bottom"
)

// LayoutDirection Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the direction corresponding to positive angles in legacy polar charts.
type LayoutDirection string

const (
	LayoutDirection_clockwise        LayoutDirection = "clockwise"
	LayoutDirection_counterclockwise LayoutDirection = "counterclockwise"
)

// LayoutDragmode Determines the mode of drag interactions. *select* and *lasso* apply only to scatter traces with markers or text. *orbit* and *turntable* apply only to 3D scenes.
type LayoutDragmode interface{}

var (
	LayoutDragmode_zoom           LayoutDragmode = "zoom"
	LayoutDragmode_pan            LayoutDragmode = "pan"
	LayoutDragmode_select         LayoutDragmode = "select"
	LayoutDragmode_lasso          LayoutDragmode = "lasso"
	LayoutDragmode_drawclosedpath LayoutDragmode = "drawclosedpath"
	LayoutDragmode_drawopenpath   LayoutDragmode = "drawopenpath"
	LayoutDragmode_drawline       LayoutDragmode = "drawline"
	LayoutDragmode_drawrect       LayoutDragmode = "drawrect"
	LayoutDragmode_drawcircle     LayoutDragmode = "drawcircle"
	LayoutDragmode_orbit          LayoutDragmode = "orbit"
	LayoutDragmode_turntable      LayoutDragmode = "turntable"
	LayoutDragmode_False          LayoutDragmode = false
)

// LayoutFunnelmode Determines how bars at the same location coordinate are displayed on the graph. With *stack*, the bars are stacked on top of one another With *group*, the bars are plotted next to one another centered around the shared location. With *overlay*, the bars are plotted over one another, you might need to an *opacity* to see multiple bars.
type LayoutFunnelmode string

const (
	LayoutFunnelmode_stack   LayoutFunnelmode = "stack"
	LayoutFunnelmode_group   LayoutFunnelmode = "group"
	LayoutFunnelmode_overlay LayoutFunnelmode = "overlay"
)

// LayoutGeoFitbounds Determines if this subplot's view settings are auto-computed to fit trace data. On scoped maps, setting `fitbounds` leads to `center.lon` and `center.lat` getting auto-filled. On maps with a non-clipped projection, setting `fitbounds` leads to `center.lon`, `center.lat`, and `projection.rotation.lon` getting auto-filled. On maps with a clipped projection, setting `fitbounds` leads to `center.lon`, `center.lat`, `projection.rotation.lon`, `projection.rotation.lat`, `lonaxis.range` and `lonaxis.range` getting auto-filled. If *locations*, only the trace's visible locations are considered in the `fitbounds` computations. If *geojson*, the entire trace input `geojson` (if provided) is considered in the `fitbounds` computations, Defaults to *false*.
type LayoutGeoFitbounds interface{}

var (
	LayoutGeoFitbounds_False     LayoutGeoFitbounds = false
	LayoutGeoFitbounds_locations LayoutGeoFitbounds = "locations"
	LayoutGeoFitbounds_geojson   LayoutGeoFitbounds = "geojson"
)

// LayoutGeoProjectionType Sets the projection type.
type LayoutGeoProjectionType string

const (
	LayoutGeoProjectionType_equirectangular      LayoutGeoProjectionType = "equirectangular"
	LayoutGeoProjectionType_mercator             LayoutGeoProjectionType = "mercator"
	LayoutGeoProjectionType_orthographic         LayoutGeoProjectionType = "orthographic"
	LayoutGeoProjectionType_naturalearth         LayoutGeoProjectionType = "natural earth"
	LayoutGeoProjectionType_kavrayskiy7          LayoutGeoProjectionType = "kavrayskiy7"
	LayoutGeoProjectionType_miller               LayoutGeoProjectionType = "miller"
	LayoutGeoProjectionType_robinson             LayoutGeoProjectionType = "robinson"
	LayoutGeoProjectionType_eckert4              LayoutGeoProjectionType = "eckert4"
	LayoutGeoProjectionType_azimuthalequalarea   LayoutGeoProjectionType = "azimuthal equal area"
	LayoutGeoProjectionType_azimuthalequidistant LayoutGeoProjectionType = "azimuthal equidistant"
	LayoutGeoProjectionType_conicequalarea       LayoutGeoProjectionType = "conic equal area"
	LayoutGeoProjectionType_conicconformal       LayoutGeoProjectionType = "conic conformal"
	LayoutGeoProjectionType_conicequidistant     LayoutGeoProjectionType = "conic equidistant"
	LayoutGeoProjectionType_gnomonic             LayoutGeoProjectionType = "gnomonic"
	LayoutGeoProjectionType_stereographic        LayoutGeoProjectionType = "stereographic"
	LayoutGeoProjectionType_mollweide            LayoutGeoProjectionType = "mollweide"
	LayoutGeoProjectionType_hammer               LayoutGeoProjectionType = "hammer"
	LayoutGeoProjectionType_transversemercator   LayoutGeoProjectionType = "transverse mercator"
	LayoutGeoProjectionType_albersusa            LayoutGeoProjectionType = "albers usa"
	LayoutGeoProjectionType_winkeltripel         LayoutGeoProjectionType = "winkel tripel"
	LayoutGeoProjectionType_aitoff               LayoutGeoProjectionType = "aitoff"
	LayoutGeoProjectionType_sinusoidal           LayoutGeoProjectionType = "sinusoidal"
)

// LayoutGeoResolution Sets the resolution of the base layers. The values have units of km/mm e.g. 110 corresponds to a scale ratio of 1:110,000,000.
type LayoutGeoResolution string

const (
	LayoutGeoResolution110 LayoutGeoResolution = "110"
	LayoutGeoResolution50  LayoutGeoResolution = "50"
)

// LayoutGeoScope Set the scope of the map.
type LayoutGeoScope string

const (
	LayoutGeoScope_world        LayoutGeoScope = "world"
	LayoutGeoScope_usa          LayoutGeoScope = "usa"
	LayoutGeoScope_europe       LayoutGeoScope = "europe"
	LayoutGeoScope_asia         LayoutGeoScope = "asia"
	LayoutGeoScope_africa       LayoutGeoScope = "africa"
	LayoutGeoScope_northamerica LayoutGeoScope = "north america"
	LayoutGeoScope_southamerica LayoutGeoScope = "south america"
)

// LayoutGridPattern If no `subplots`, `xaxes`, or `yaxes` are given but we do have `rows` and `columns`, we can generate defaults using consecutive axis IDs, in two ways: *coupled* gives one x axis per column and one y axis per row. *independent* uses a new xy pair for each cell, left-to-right across each row then iterating rows according to `roworder`.
type LayoutGridPattern string

const (
	LayoutGridPattern_independent LayoutGridPattern = "independent"
	LayoutGridPattern_coupled     LayoutGridPattern = "coupled"
)

// LayoutGridRoworder Is the first row the top or the bottom? Note that columns are always enumerated from left to right.
type LayoutGridRoworder string

const (
	LayoutGridRoworder_toptobottom LayoutGridRoworder = "top to bottom"
	LayoutGridRoworder_bottomtotop LayoutGridRoworder = "bottom to top"
)

// LayoutGridXside Sets where the x axis labels and titles go. *bottom* means the very bottom of the grid. *bottom plot* is the lowest plot that each x axis is used in. *top* and *top plot* are similar.
type LayoutGridXside string

const (
	LayoutGridXside_bottom     LayoutGridXside = "bottom"
	LayoutGridXside_bottomplot LayoutGridXside = "bottom plot"
	LayoutGridXside_topplot    LayoutGridXside = "top plot"
	LayoutGridXside_top        LayoutGridXside = "top"
)

// LayoutGridYside Sets where the y axis labels and titles go. *left* means the very left edge of the grid. *left plot* is the leftmost plot that each y axis is used in. *right* and *right plot* are similar.
type LayoutGridYside string

const (
	LayoutGridYside_left      LayoutGridYside = "left"
	LayoutGridYside_leftplot  LayoutGridYside = "left plot"
	LayoutGridYside_rightplot LayoutGridYside = "right plot"
	LayoutGridYside_right     LayoutGridYside = "right"
)

// LayoutHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type LayoutHoverlabelAlign string

const (
	LayoutHoverlabelAlign_left  LayoutHoverlabelAlign = "left"
	LayoutHoverlabelAlign_right LayoutHoverlabelAlign = "right"
	LayoutHoverlabelAlign_auto  LayoutHoverlabelAlign = "auto"
)

// LayoutHovermode Determines the mode of hover interactions. If *closest*, a single hoverlabel will appear for the *closest* point within the `hoverdistance`. If *x* (or *y*), multiple hoverlabels will appear for multiple points at the *closest* x- (or y-) coordinate within the `hoverdistance`, with the caveat that no more than one hoverlabel will appear per trace. If *x unified* (or *y unified*), a single hoverlabel will appear multiple points at the closest x- (or y-) coordinate within the `hoverdistance` with the caveat that no more than one hoverlabel will appear per trace. In this mode, spikelines are enabled by default perpendicular to the specified axis. If false, hover interactions are disabled. If `clickmode` includes the *select* flag, `hovermode` defaults to *closest*. If `clickmode` lacks the *select* flag, it defaults to *x* or *y* (depending on the trace's `orientation` value) for plots based on cartesian coordinates. For anything else the default value is *closest*.
type LayoutHovermode interface{}

var (
	LayoutHovermode_x        LayoutHovermode = "x"
	LayoutHovermode_y        LayoutHovermode = "y"
	LayoutHovermode_closest  LayoutHovermode = "closest"
	LayoutHovermode_False    LayoutHovermode = false
	LayoutHovermode_xunified LayoutHovermode = "x unified"
	LayoutHovermode_yunified LayoutHovermode = "y unified"
)

// LayoutLegendItemclick Determines the behavior on legend item click. *toggle* toggles the visibility of the item clicked on the graph. *toggleothers* makes the clicked item the sole visible item on the graph. *false* disable legend item click interactions.
type LayoutLegendItemclick interface{}

var (
	LayoutLegendItemclick_toggle       LayoutLegendItemclick = "toggle"
	LayoutLegendItemclick_toggleothers LayoutLegendItemclick = "toggleothers"
	LayoutLegendItemclick_False        LayoutLegendItemclick = false
)

// LayoutLegendItemdoubleclick Determines the behavior on legend item double-click. *toggle* toggles the visibility of the item clicked on the graph. *toggleothers* makes the clicked item the sole visible item on the graph. *false* disable legend item double-click interactions.
type LayoutLegendItemdoubleclick interface{}

var (
	LayoutLegendItemdoubleclick_toggle       LayoutLegendItemdoubleclick = "toggle"
	LayoutLegendItemdoubleclick_toggleothers LayoutLegendItemdoubleclick = "toggleothers"
	LayoutLegendItemdoubleclick_False        LayoutLegendItemdoubleclick = false
)

// LayoutLegendItemsizing Determines if the legend items symbols scale with their corresponding *trace* attributes or remain *constant* independent of the symbol size on the graph.
type LayoutLegendItemsizing string

const (
	LayoutLegendItemsizing_trace    LayoutLegendItemsizing = "trace"
	LayoutLegendItemsizing_constant LayoutLegendItemsizing = "constant"
)

// LayoutLegendOrientation Sets the orientation of the legend.
type LayoutLegendOrientation string

const (
	LayoutLegendOrientation_v LayoutLegendOrientation = "v"
	LayoutLegendOrientation_h LayoutLegendOrientation = "h"
)

// LayoutLegendTitleSide Determines the location of legend's title with respect to the legend items. Defaulted to *top* with `orientation` is *h*. Defaulted to *left* with `orientation` is *v*. The *top left* options could be used to expand legend area in both x and y sides.
type LayoutLegendTitleSide string

const (
	LayoutLegendTitleSide_top     LayoutLegendTitleSide = "top"
	LayoutLegendTitleSide_left    LayoutLegendTitleSide = "left"
	LayoutLegendTitleSide_topleft LayoutLegendTitleSide = "top left"
)

// LayoutLegendValign Sets the vertical alignment of the symbols with respect to their associated text.
type LayoutLegendValign string

const (
	LayoutLegendValign_top    LayoutLegendValign = "top"
	LayoutLegendValign_middle LayoutLegendValign = "middle"
	LayoutLegendValign_bottom LayoutLegendValign = "bottom"
)

// LayoutLegendXanchor Sets the legend's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the legend. Value *auto* anchors legends to the right for `x` values greater than or equal to 2/3, anchors legends to the left for `x` values less than or equal to 1/3 and anchors legends with respect to their center otherwise.
type LayoutLegendXanchor string

const (
	LayoutLegendXanchor_auto   LayoutLegendXanchor = "auto"
	LayoutLegendXanchor_left   LayoutLegendXanchor = "left"
	LayoutLegendXanchor_center LayoutLegendXanchor = "center"
	LayoutLegendXanchor_right  LayoutLegendXanchor = "right"
)

// LayoutLegendYanchor Sets the legend's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the legend. Value *auto* anchors legends at their bottom for `y` values less than or equal to 1/3, anchors legends to at their top for `y` values greater than or equal to 2/3 and anchors legends with respect to their middle otherwise.
type LayoutLegendYanchor string

const (
	LayoutLegendYanchor_auto   LayoutLegendYanchor = "auto"
	LayoutLegendYanchor_top    LayoutLegendYanchor = "top"
	LayoutLegendYanchor_middle LayoutLegendYanchor = "middle"
	LayoutLegendYanchor_bottom LayoutLegendYanchor = "bottom"
)

// LayoutModebarOrientation Sets the orientation of the modebar.
type LayoutModebarOrientation string

const (
	LayoutModebarOrientation_v LayoutModebarOrientation = "v"
	LayoutModebarOrientation_h LayoutModebarOrientation = "h"
)

// LayoutNewshapeDrawdirection When `dragmode` is set to *drawrect*, *drawline* or *drawcircle* this limits the drag to be horizontal, vertical or diagonal. Using *diagonal* there is no limit e.g. in drawing lines in any direction. *ortho* limits the draw to be either horizontal or vertical. *horizontal* allows horizontal extend. *vertical* allows vertical extend.
type LayoutNewshapeDrawdirection string

const (
	LayoutNewshapeDrawdirection_ortho      LayoutNewshapeDrawdirection = "ortho"
	LayoutNewshapeDrawdirection_horizontal LayoutNewshapeDrawdirection = "horizontal"
	LayoutNewshapeDrawdirection_vertical   LayoutNewshapeDrawdirection = "vertical"
	LayoutNewshapeDrawdirection_diagonal   LayoutNewshapeDrawdirection = "diagonal"
)

// LayoutNewshapeFillrule Determines the path's interior. For more info please visit https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fill-rule
type LayoutNewshapeFillrule string

const (
	LayoutNewshapeFillrule_evenodd LayoutNewshapeFillrule = "evenodd"
	LayoutNewshapeFillrule_nonzero LayoutNewshapeFillrule = "nonzero"
)

// LayoutNewshapeLayer Specifies whether new shapes are drawn below or above traces.
type LayoutNewshapeLayer string

const (
	LayoutNewshapeLayer_below LayoutNewshapeLayer = "below"
	LayoutNewshapeLayer_above LayoutNewshapeLayer = "above"
)

// LayoutPolarAngularaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutPolarAngularaxisCategoryorder string

const (
	LayoutPolarAngularaxisCategoryorder_trace              LayoutPolarAngularaxisCategoryorder = "trace"
	LayoutPolarAngularaxisCategoryorder_categoryascending  LayoutPolarAngularaxisCategoryorder = "category ascending"
	LayoutPolarAngularaxisCategoryorder_categorydescending LayoutPolarAngularaxisCategoryorder = "category descending"
	LayoutPolarAngularaxisCategoryorder_array              LayoutPolarAngularaxisCategoryorder = "array"
	LayoutPolarAngularaxisCategoryorder_totalascending     LayoutPolarAngularaxisCategoryorder = "total ascending"
	LayoutPolarAngularaxisCategoryorder_totaldescending    LayoutPolarAngularaxisCategoryorder = "total descending"
	LayoutPolarAngularaxisCategoryorder_minascending       LayoutPolarAngularaxisCategoryorder = "min ascending"
	LayoutPolarAngularaxisCategoryorder_mindescending      LayoutPolarAngularaxisCategoryorder = "min descending"
	LayoutPolarAngularaxisCategoryorder_maxascending       LayoutPolarAngularaxisCategoryorder = "max ascending"
	LayoutPolarAngularaxisCategoryorder_maxdescending      LayoutPolarAngularaxisCategoryorder = "max descending"
	LayoutPolarAngularaxisCategoryorder_sumascending       LayoutPolarAngularaxisCategoryorder = "sum ascending"
	LayoutPolarAngularaxisCategoryorder_sumdescending      LayoutPolarAngularaxisCategoryorder = "sum descending"
	LayoutPolarAngularaxisCategoryorder_meanascending      LayoutPolarAngularaxisCategoryorder = "mean ascending"
	LayoutPolarAngularaxisCategoryorder_meandescending     LayoutPolarAngularaxisCategoryorder = "mean descending"
	LayoutPolarAngularaxisCategoryorder_medianascending    LayoutPolarAngularaxisCategoryorder = "median ascending"
	LayoutPolarAngularaxisCategoryorder_mediandescending   LayoutPolarAngularaxisCategoryorder = "median descending"
)

// LayoutPolarAngularaxisDirection Sets the direction corresponding to positive angles.
type LayoutPolarAngularaxisDirection string

const (
	LayoutPolarAngularaxisDirection_counterclockwise LayoutPolarAngularaxisDirection = "counterclockwise"
	LayoutPolarAngularaxisDirection_clockwise        LayoutPolarAngularaxisDirection = "clockwise"
)

// LayoutPolarAngularaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutPolarAngularaxisExponentformat string

const (
	LayoutPolarAngularaxisExponentformat_none  LayoutPolarAngularaxisExponentformat = "none"
	LayoutPolarAngularaxisExponentformat_e     LayoutPolarAngularaxisExponentformat = "e"
	LayoutPolarAngularaxisExponentformat_E     LayoutPolarAngularaxisExponentformat = "E"
	LayoutPolarAngularaxisExponentformat_power LayoutPolarAngularaxisExponentformat = "power"
	LayoutPolarAngularaxisExponentformat_SI    LayoutPolarAngularaxisExponentformat = "SI"
	LayoutPolarAngularaxisExponentformat_B     LayoutPolarAngularaxisExponentformat = "B"
)

// LayoutPolarAngularaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutPolarAngularaxisLayer string

const (
	LayoutPolarAngularaxisLayer_abovetraces LayoutPolarAngularaxisLayer = "above traces"
	LayoutPolarAngularaxisLayer_belowtraces LayoutPolarAngularaxisLayer = "below traces"
)

// LayoutPolarAngularaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutPolarAngularaxisShowexponent string

const (
	LayoutPolarAngularaxisShowexponent_all   LayoutPolarAngularaxisShowexponent = "all"
	LayoutPolarAngularaxisShowexponent_first LayoutPolarAngularaxisShowexponent = "first"
	LayoutPolarAngularaxisShowexponent_last  LayoutPolarAngularaxisShowexponent = "last"
	LayoutPolarAngularaxisShowexponent_none  LayoutPolarAngularaxisShowexponent = "none"
)

// LayoutPolarAngularaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutPolarAngularaxisShowtickprefix string

const (
	LayoutPolarAngularaxisShowtickprefix_all   LayoutPolarAngularaxisShowtickprefix = "all"
	LayoutPolarAngularaxisShowtickprefix_first LayoutPolarAngularaxisShowtickprefix = "first"
	LayoutPolarAngularaxisShowtickprefix_last  LayoutPolarAngularaxisShowtickprefix = "last"
	LayoutPolarAngularaxisShowtickprefix_none  LayoutPolarAngularaxisShowtickprefix = "none"
)

// LayoutPolarAngularaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutPolarAngularaxisShowticksuffix string

const (
	LayoutPolarAngularaxisShowticksuffix_all   LayoutPolarAngularaxisShowticksuffix = "all"
	LayoutPolarAngularaxisShowticksuffix_first LayoutPolarAngularaxisShowticksuffix = "first"
	LayoutPolarAngularaxisShowticksuffix_last  LayoutPolarAngularaxisShowticksuffix = "last"
	LayoutPolarAngularaxisShowticksuffix_none  LayoutPolarAngularaxisShowticksuffix = "none"
)

// LayoutPolarAngularaxisThetaunit Sets the format unit of the formatted *theta* values. Has an effect only when `angularaxis.type` is *linear*.
type LayoutPolarAngularaxisThetaunit string

const (
	LayoutPolarAngularaxisThetaunit_radians LayoutPolarAngularaxisThetaunit = "radians"
	LayoutPolarAngularaxisThetaunit_degrees LayoutPolarAngularaxisThetaunit = "degrees"
)

// LayoutPolarAngularaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutPolarAngularaxisTickmode string

const (
	LayoutPolarAngularaxisTickmode_auto   LayoutPolarAngularaxisTickmode = "auto"
	LayoutPolarAngularaxisTickmode_linear LayoutPolarAngularaxisTickmode = "linear"
	LayoutPolarAngularaxisTickmode_array  LayoutPolarAngularaxisTickmode = "array"
)

// LayoutPolarAngularaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutPolarAngularaxisTicks string

const (
	LayoutPolarAngularaxisTicks_outside LayoutPolarAngularaxisTicks = "outside"
	LayoutPolarAngularaxisTicks_inside  LayoutPolarAngularaxisTicks = "inside"
)

// LayoutPolarAngularaxisType Sets the angular axis type. If *linear*, set `thetaunit` to determine the unit in which axis value are shown. If *category, use `period` to set the number of integer coordinates around polar axis.
type LayoutPolarAngularaxisType string

const (
	LayoutPolarAngularaxisType__        LayoutPolarAngularaxisType = "-"
	LayoutPolarAngularaxisType_linear   LayoutPolarAngularaxisType = "linear"
	LayoutPolarAngularaxisType_category LayoutPolarAngularaxisType = "category"
)

// LayoutPolarGridshape Determines if the radial axis grid lines and angular axis line are drawn as *circular* sectors or as *linear* (polygon) sectors. Has an effect only when the angular axis has `type` *category*. Note that `radialaxis.angle` is snapped to the angle of the closest vertex when `gridshape` is *circular* (so that radial axis scale is the same as the data scale).
type LayoutPolarGridshape string

const (
	LayoutPolarGridshape_circular LayoutPolarGridshape = "circular"
	LayoutPolarGridshape_linear   LayoutPolarGridshape = "linear"
)

// LayoutPolarRadialaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutPolarRadialaxisAutorange interface{}

var (
	LayoutPolarRadialaxisAutorange_True     LayoutPolarRadialaxisAutorange = true
	LayoutPolarRadialaxisAutorange_False    LayoutPolarRadialaxisAutorange = false
	LayoutPolarRadialaxisAutorange_reversed LayoutPolarRadialaxisAutorange = "reversed"
)

// LayoutPolarRadialaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutPolarRadialaxisCalendar string

const (
	LayoutPolarRadialaxisCalendar_gregorian  LayoutPolarRadialaxisCalendar = "gregorian"
	LayoutPolarRadialaxisCalendar_chinese    LayoutPolarRadialaxisCalendar = "chinese"
	LayoutPolarRadialaxisCalendar_coptic     LayoutPolarRadialaxisCalendar = "coptic"
	LayoutPolarRadialaxisCalendar_discworld  LayoutPolarRadialaxisCalendar = "discworld"
	LayoutPolarRadialaxisCalendar_ethiopian  LayoutPolarRadialaxisCalendar = "ethiopian"
	LayoutPolarRadialaxisCalendar_hebrew     LayoutPolarRadialaxisCalendar = "hebrew"
	LayoutPolarRadialaxisCalendar_islamic    LayoutPolarRadialaxisCalendar = "islamic"
	LayoutPolarRadialaxisCalendar_julian     LayoutPolarRadialaxisCalendar = "julian"
	LayoutPolarRadialaxisCalendar_mayan      LayoutPolarRadialaxisCalendar = "mayan"
	LayoutPolarRadialaxisCalendar_nanakshahi LayoutPolarRadialaxisCalendar = "nanakshahi"
	LayoutPolarRadialaxisCalendar_nepali     LayoutPolarRadialaxisCalendar = "nepali"
	LayoutPolarRadialaxisCalendar_persian    LayoutPolarRadialaxisCalendar = "persian"
	LayoutPolarRadialaxisCalendar_jalali     LayoutPolarRadialaxisCalendar = "jalali"
	LayoutPolarRadialaxisCalendar_taiwan     LayoutPolarRadialaxisCalendar = "taiwan"
	LayoutPolarRadialaxisCalendar_thai       LayoutPolarRadialaxisCalendar = "thai"
	LayoutPolarRadialaxisCalendar_ummalqura  LayoutPolarRadialaxisCalendar = "ummalqura"
)

// LayoutPolarRadialaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutPolarRadialaxisCategoryorder string

const (
	LayoutPolarRadialaxisCategoryorder_trace              LayoutPolarRadialaxisCategoryorder = "trace"
	LayoutPolarRadialaxisCategoryorder_categoryascending  LayoutPolarRadialaxisCategoryorder = "category ascending"
	LayoutPolarRadialaxisCategoryorder_categorydescending LayoutPolarRadialaxisCategoryorder = "category descending"
	LayoutPolarRadialaxisCategoryorder_array              LayoutPolarRadialaxisCategoryorder = "array"
	LayoutPolarRadialaxisCategoryorder_totalascending     LayoutPolarRadialaxisCategoryorder = "total ascending"
	LayoutPolarRadialaxisCategoryorder_totaldescending    LayoutPolarRadialaxisCategoryorder = "total descending"
	LayoutPolarRadialaxisCategoryorder_minascending       LayoutPolarRadialaxisCategoryorder = "min ascending"
	LayoutPolarRadialaxisCategoryorder_mindescending      LayoutPolarRadialaxisCategoryorder = "min descending"
	LayoutPolarRadialaxisCategoryorder_maxascending       LayoutPolarRadialaxisCategoryorder = "max ascending"
	LayoutPolarRadialaxisCategoryorder_maxdescending      LayoutPolarRadialaxisCategoryorder = "max descending"
	LayoutPolarRadialaxisCategoryorder_sumascending       LayoutPolarRadialaxisCategoryorder = "sum ascending"
	LayoutPolarRadialaxisCategoryorder_sumdescending      LayoutPolarRadialaxisCategoryorder = "sum descending"
	LayoutPolarRadialaxisCategoryorder_meanascending      LayoutPolarRadialaxisCategoryorder = "mean ascending"
	LayoutPolarRadialaxisCategoryorder_meandescending     LayoutPolarRadialaxisCategoryorder = "mean descending"
	LayoutPolarRadialaxisCategoryorder_medianascending    LayoutPolarRadialaxisCategoryorder = "median ascending"
	LayoutPolarRadialaxisCategoryorder_mediandescending   LayoutPolarRadialaxisCategoryorder = "median descending"
)

// LayoutPolarRadialaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutPolarRadialaxisExponentformat string

const (
	LayoutPolarRadialaxisExponentformat_none  LayoutPolarRadialaxisExponentformat = "none"
	LayoutPolarRadialaxisExponentformat_e     LayoutPolarRadialaxisExponentformat = "e"
	LayoutPolarRadialaxisExponentformat_E     LayoutPolarRadialaxisExponentformat = "E"
	LayoutPolarRadialaxisExponentformat_power LayoutPolarRadialaxisExponentformat = "power"
	LayoutPolarRadialaxisExponentformat_SI    LayoutPolarRadialaxisExponentformat = "SI"
	LayoutPolarRadialaxisExponentformat_B     LayoutPolarRadialaxisExponentformat = "B"
)

// LayoutPolarRadialaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutPolarRadialaxisLayer string

const (
	LayoutPolarRadialaxisLayer_abovetraces LayoutPolarRadialaxisLayer = "above traces"
	LayoutPolarRadialaxisLayer_belowtraces LayoutPolarRadialaxisLayer = "below traces"
)

// LayoutPolarRadialaxisRangemode If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. If *normal*, the range is computed in relation to the extrema of the input data (same behavior as for cartesian axes).
type LayoutPolarRadialaxisRangemode string

const (
	LayoutPolarRadialaxisRangemode_tozero      LayoutPolarRadialaxisRangemode = "tozero"
	LayoutPolarRadialaxisRangemode_nonnegative LayoutPolarRadialaxisRangemode = "nonnegative"
	LayoutPolarRadialaxisRangemode_normal      LayoutPolarRadialaxisRangemode = "normal"
)

// LayoutPolarRadialaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutPolarRadialaxisShowexponent string

const (
	LayoutPolarRadialaxisShowexponent_all   LayoutPolarRadialaxisShowexponent = "all"
	LayoutPolarRadialaxisShowexponent_first LayoutPolarRadialaxisShowexponent = "first"
	LayoutPolarRadialaxisShowexponent_last  LayoutPolarRadialaxisShowexponent = "last"
	LayoutPolarRadialaxisShowexponent_none  LayoutPolarRadialaxisShowexponent = "none"
)

// LayoutPolarRadialaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutPolarRadialaxisShowtickprefix string

const (
	LayoutPolarRadialaxisShowtickprefix_all   LayoutPolarRadialaxisShowtickprefix = "all"
	LayoutPolarRadialaxisShowtickprefix_first LayoutPolarRadialaxisShowtickprefix = "first"
	LayoutPolarRadialaxisShowtickprefix_last  LayoutPolarRadialaxisShowtickprefix = "last"
	LayoutPolarRadialaxisShowtickprefix_none  LayoutPolarRadialaxisShowtickprefix = "none"
)

// LayoutPolarRadialaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutPolarRadialaxisShowticksuffix string

const (
	LayoutPolarRadialaxisShowticksuffix_all   LayoutPolarRadialaxisShowticksuffix = "all"
	LayoutPolarRadialaxisShowticksuffix_first LayoutPolarRadialaxisShowticksuffix = "first"
	LayoutPolarRadialaxisShowticksuffix_last  LayoutPolarRadialaxisShowticksuffix = "last"
	LayoutPolarRadialaxisShowticksuffix_none  LayoutPolarRadialaxisShowticksuffix = "none"
)

// LayoutPolarRadialaxisSide Determines on which side of radial axis line the tick and tick labels appear.
type LayoutPolarRadialaxisSide string

const (
	LayoutPolarRadialaxisSide_clockwise        LayoutPolarRadialaxisSide = "clockwise"
	LayoutPolarRadialaxisSide_counterclockwise LayoutPolarRadialaxisSide = "counterclockwise"
)

// LayoutPolarRadialaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutPolarRadialaxisTickmode string

const (
	LayoutPolarRadialaxisTickmode_auto   LayoutPolarRadialaxisTickmode = "auto"
	LayoutPolarRadialaxisTickmode_linear LayoutPolarRadialaxisTickmode = "linear"
	LayoutPolarRadialaxisTickmode_array  LayoutPolarRadialaxisTickmode = "array"
)

// LayoutPolarRadialaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutPolarRadialaxisTicks string

const (
	LayoutPolarRadialaxisTicks_outside LayoutPolarRadialaxisTicks = "outside"
	LayoutPolarRadialaxisTicks_inside  LayoutPolarRadialaxisTicks = "inside"
)

// LayoutPolarRadialaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutPolarRadialaxisType string

const (
	LayoutPolarRadialaxisType__        LayoutPolarRadialaxisType = "-"
	LayoutPolarRadialaxisType_linear   LayoutPolarRadialaxisType = "linear"
	LayoutPolarRadialaxisType_log      LayoutPolarRadialaxisType = "log"
	LayoutPolarRadialaxisType_date     LayoutPolarRadialaxisType = "date"
	LayoutPolarRadialaxisType_category LayoutPolarRadialaxisType = "category"
)

// LayoutRadialaxisTickorientation Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the orientation (from the paper perspective) of the radial axis tick labels.
type LayoutRadialaxisTickorientation string

const (
	LayoutRadialaxisTickorientation_horizontal LayoutRadialaxisTickorientation = "horizontal"
	LayoutRadialaxisTickorientation_vertical   LayoutRadialaxisTickorientation = "vertical"
)

// LayoutSceneAspectmode If *cube*, this scene's axes are drawn as a cube, regardless of the axes' ranges. If *data*, this scene's axes are drawn in proportion with the axes' ranges. If *manual*, this scene's axes are drawn in proportion with the input of *aspectratio* (the default behavior if *aspectratio* is provided). If *auto*, this scene's axes are drawn using the results of *data* except when one axis is more than four times the size of the two others, where in that case the results of *cube* are used.
type LayoutSceneAspectmode string

const (
	LayoutSceneAspectmode_auto   LayoutSceneAspectmode = "auto"
	LayoutSceneAspectmode_cube   LayoutSceneAspectmode = "cube"
	LayoutSceneAspectmode_data   LayoutSceneAspectmode = "data"
	LayoutSceneAspectmode_manual LayoutSceneAspectmode = "manual"
)

// LayoutSceneCameraProjectionType Sets the projection type. The projection type could be either *perspective* or *orthographic*. The default is *perspective*.
type LayoutSceneCameraProjectionType string

const (
	LayoutSceneCameraProjectionType_perspective  LayoutSceneCameraProjectionType = "perspective"
	LayoutSceneCameraProjectionType_orthographic LayoutSceneCameraProjectionType = "orthographic"
)

// LayoutSceneDragmode Determines the mode of drag interactions for this scene.
type LayoutSceneDragmode interface{}

var (
	LayoutSceneDragmode_orbit     LayoutSceneDragmode = "orbit"
	LayoutSceneDragmode_turntable LayoutSceneDragmode = "turntable"
	LayoutSceneDragmode_zoom      LayoutSceneDragmode = "zoom"
	LayoutSceneDragmode_pan       LayoutSceneDragmode = "pan"
	LayoutSceneDragmode_False     LayoutSceneDragmode = false
)

// LayoutSceneHovermode Determines the mode of hover interactions for this scene.
type LayoutSceneHovermode interface{}

var (
	LayoutSceneHovermode_closest LayoutSceneHovermode = "closest"
	LayoutSceneHovermode_False   LayoutSceneHovermode = false
)

// LayoutSceneXaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutSceneXaxisAutorange interface{}

var (
	LayoutSceneXaxisAutorange_True     LayoutSceneXaxisAutorange = true
	LayoutSceneXaxisAutorange_False    LayoutSceneXaxisAutorange = false
	LayoutSceneXaxisAutorange_reversed LayoutSceneXaxisAutorange = "reversed"
)

// LayoutSceneXaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutSceneXaxisCalendar string

const (
	LayoutSceneXaxisCalendar_gregorian  LayoutSceneXaxisCalendar = "gregorian"
	LayoutSceneXaxisCalendar_chinese    LayoutSceneXaxisCalendar = "chinese"
	LayoutSceneXaxisCalendar_coptic     LayoutSceneXaxisCalendar = "coptic"
	LayoutSceneXaxisCalendar_discworld  LayoutSceneXaxisCalendar = "discworld"
	LayoutSceneXaxisCalendar_ethiopian  LayoutSceneXaxisCalendar = "ethiopian"
	LayoutSceneXaxisCalendar_hebrew     LayoutSceneXaxisCalendar = "hebrew"
	LayoutSceneXaxisCalendar_islamic    LayoutSceneXaxisCalendar = "islamic"
	LayoutSceneXaxisCalendar_julian     LayoutSceneXaxisCalendar = "julian"
	LayoutSceneXaxisCalendar_mayan      LayoutSceneXaxisCalendar = "mayan"
	LayoutSceneXaxisCalendar_nanakshahi LayoutSceneXaxisCalendar = "nanakshahi"
	LayoutSceneXaxisCalendar_nepali     LayoutSceneXaxisCalendar = "nepali"
	LayoutSceneXaxisCalendar_persian    LayoutSceneXaxisCalendar = "persian"
	LayoutSceneXaxisCalendar_jalali     LayoutSceneXaxisCalendar = "jalali"
	LayoutSceneXaxisCalendar_taiwan     LayoutSceneXaxisCalendar = "taiwan"
	LayoutSceneXaxisCalendar_thai       LayoutSceneXaxisCalendar = "thai"
	LayoutSceneXaxisCalendar_ummalqura  LayoutSceneXaxisCalendar = "ummalqura"
)

// LayoutSceneXaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutSceneXaxisCategoryorder string

const (
	LayoutSceneXaxisCategoryorder_trace              LayoutSceneXaxisCategoryorder = "trace"
	LayoutSceneXaxisCategoryorder_categoryascending  LayoutSceneXaxisCategoryorder = "category ascending"
	LayoutSceneXaxisCategoryorder_categorydescending LayoutSceneXaxisCategoryorder = "category descending"
	LayoutSceneXaxisCategoryorder_array              LayoutSceneXaxisCategoryorder = "array"
	LayoutSceneXaxisCategoryorder_totalascending     LayoutSceneXaxisCategoryorder = "total ascending"
	LayoutSceneXaxisCategoryorder_totaldescending    LayoutSceneXaxisCategoryorder = "total descending"
	LayoutSceneXaxisCategoryorder_minascending       LayoutSceneXaxisCategoryorder = "min ascending"
	LayoutSceneXaxisCategoryorder_mindescending      LayoutSceneXaxisCategoryorder = "min descending"
	LayoutSceneXaxisCategoryorder_maxascending       LayoutSceneXaxisCategoryorder = "max ascending"
	LayoutSceneXaxisCategoryorder_maxdescending      LayoutSceneXaxisCategoryorder = "max descending"
	LayoutSceneXaxisCategoryorder_sumascending       LayoutSceneXaxisCategoryorder = "sum ascending"
	LayoutSceneXaxisCategoryorder_sumdescending      LayoutSceneXaxisCategoryorder = "sum descending"
	LayoutSceneXaxisCategoryorder_meanascending      LayoutSceneXaxisCategoryorder = "mean ascending"
	LayoutSceneXaxisCategoryorder_meandescending     LayoutSceneXaxisCategoryorder = "mean descending"
	LayoutSceneXaxisCategoryorder_medianascending    LayoutSceneXaxisCategoryorder = "median ascending"
	LayoutSceneXaxisCategoryorder_mediandescending   LayoutSceneXaxisCategoryorder = "median descending"
)

// LayoutSceneXaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutSceneXaxisExponentformat string

const (
	LayoutSceneXaxisExponentformat_none  LayoutSceneXaxisExponentformat = "none"
	LayoutSceneXaxisExponentformat_e     LayoutSceneXaxisExponentformat = "e"
	LayoutSceneXaxisExponentformat_E     LayoutSceneXaxisExponentformat = "E"
	LayoutSceneXaxisExponentformat_power LayoutSceneXaxisExponentformat = "power"
	LayoutSceneXaxisExponentformat_SI    LayoutSceneXaxisExponentformat = "SI"
	LayoutSceneXaxisExponentformat_B     LayoutSceneXaxisExponentformat = "B"
)

// LayoutSceneXaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutSceneXaxisMirror interface{}

var (
	LayoutSceneXaxisMirror_True     LayoutSceneXaxisMirror = true
	LayoutSceneXaxisMirror_ticks    LayoutSceneXaxisMirror = "ticks"
	LayoutSceneXaxisMirror_False    LayoutSceneXaxisMirror = false
	LayoutSceneXaxisMirror_all      LayoutSceneXaxisMirror = "all"
	LayoutSceneXaxisMirror_allticks LayoutSceneXaxisMirror = "allticks"
)

// LayoutSceneXaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutSceneXaxisRangemode string

const (
	LayoutSceneXaxisRangemode_normal      LayoutSceneXaxisRangemode = "normal"
	LayoutSceneXaxisRangemode_tozero      LayoutSceneXaxisRangemode = "tozero"
	LayoutSceneXaxisRangemode_nonnegative LayoutSceneXaxisRangemode = "nonnegative"
)

// LayoutSceneXaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutSceneXaxisShowexponent string

const (
	LayoutSceneXaxisShowexponent_all   LayoutSceneXaxisShowexponent = "all"
	LayoutSceneXaxisShowexponent_first LayoutSceneXaxisShowexponent = "first"
	LayoutSceneXaxisShowexponent_last  LayoutSceneXaxisShowexponent = "last"
	LayoutSceneXaxisShowexponent_none  LayoutSceneXaxisShowexponent = "none"
)

// LayoutSceneXaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutSceneXaxisShowtickprefix string

const (
	LayoutSceneXaxisShowtickprefix_all   LayoutSceneXaxisShowtickprefix = "all"
	LayoutSceneXaxisShowtickprefix_first LayoutSceneXaxisShowtickprefix = "first"
	LayoutSceneXaxisShowtickprefix_last  LayoutSceneXaxisShowtickprefix = "last"
	LayoutSceneXaxisShowtickprefix_none  LayoutSceneXaxisShowtickprefix = "none"
)

// LayoutSceneXaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutSceneXaxisShowticksuffix string

const (
	LayoutSceneXaxisShowticksuffix_all   LayoutSceneXaxisShowticksuffix = "all"
	LayoutSceneXaxisShowticksuffix_first LayoutSceneXaxisShowticksuffix = "first"
	LayoutSceneXaxisShowticksuffix_last  LayoutSceneXaxisShowticksuffix = "last"
	LayoutSceneXaxisShowticksuffix_none  LayoutSceneXaxisShowticksuffix = "none"
)

// LayoutSceneXaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutSceneXaxisTickmode string

const (
	LayoutSceneXaxisTickmode_auto   LayoutSceneXaxisTickmode = "auto"
	LayoutSceneXaxisTickmode_linear LayoutSceneXaxisTickmode = "linear"
	LayoutSceneXaxisTickmode_array  LayoutSceneXaxisTickmode = "array"
)

// LayoutSceneXaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutSceneXaxisTicks string

const (
	LayoutSceneXaxisTicks_outside LayoutSceneXaxisTicks = "outside"
	LayoutSceneXaxisTicks_inside  LayoutSceneXaxisTicks = "inside"
)

// LayoutSceneXaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutSceneXaxisType string

const (
	LayoutSceneXaxisType__        LayoutSceneXaxisType = "-"
	LayoutSceneXaxisType_linear   LayoutSceneXaxisType = "linear"
	LayoutSceneXaxisType_log      LayoutSceneXaxisType = "log"
	LayoutSceneXaxisType_date     LayoutSceneXaxisType = "date"
	LayoutSceneXaxisType_category LayoutSceneXaxisType = "category"
)

// LayoutSceneYaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutSceneYaxisAutorange interface{}

var (
	LayoutSceneYaxisAutorange_True     LayoutSceneYaxisAutorange = true
	LayoutSceneYaxisAutorange_False    LayoutSceneYaxisAutorange = false
	LayoutSceneYaxisAutorange_reversed LayoutSceneYaxisAutorange = "reversed"
)

// LayoutSceneYaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutSceneYaxisCalendar string

const (
	LayoutSceneYaxisCalendar_gregorian  LayoutSceneYaxisCalendar = "gregorian"
	LayoutSceneYaxisCalendar_chinese    LayoutSceneYaxisCalendar = "chinese"
	LayoutSceneYaxisCalendar_coptic     LayoutSceneYaxisCalendar = "coptic"
	LayoutSceneYaxisCalendar_discworld  LayoutSceneYaxisCalendar = "discworld"
	LayoutSceneYaxisCalendar_ethiopian  LayoutSceneYaxisCalendar = "ethiopian"
	LayoutSceneYaxisCalendar_hebrew     LayoutSceneYaxisCalendar = "hebrew"
	LayoutSceneYaxisCalendar_islamic    LayoutSceneYaxisCalendar = "islamic"
	LayoutSceneYaxisCalendar_julian     LayoutSceneYaxisCalendar = "julian"
	LayoutSceneYaxisCalendar_mayan      LayoutSceneYaxisCalendar = "mayan"
	LayoutSceneYaxisCalendar_nanakshahi LayoutSceneYaxisCalendar = "nanakshahi"
	LayoutSceneYaxisCalendar_nepali     LayoutSceneYaxisCalendar = "nepali"
	LayoutSceneYaxisCalendar_persian    LayoutSceneYaxisCalendar = "persian"
	LayoutSceneYaxisCalendar_jalali     LayoutSceneYaxisCalendar = "jalali"
	LayoutSceneYaxisCalendar_taiwan     LayoutSceneYaxisCalendar = "taiwan"
	LayoutSceneYaxisCalendar_thai       LayoutSceneYaxisCalendar = "thai"
	LayoutSceneYaxisCalendar_ummalqura  LayoutSceneYaxisCalendar = "ummalqura"
)

// LayoutSceneYaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutSceneYaxisCategoryorder string

const (
	LayoutSceneYaxisCategoryorder_trace              LayoutSceneYaxisCategoryorder = "trace"
	LayoutSceneYaxisCategoryorder_categoryascending  LayoutSceneYaxisCategoryorder = "category ascending"
	LayoutSceneYaxisCategoryorder_categorydescending LayoutSceneYaxisCategoryorder = "category descending"
	LayoutSceneYaxisCategoryorder_array              LayoutSceneYaxisCategoryorder = "array"
	LayoutSceneYaxisCategoryorder_totalascending     LayoutSceneYaxisCategoryorder = "total ascending"
	LayoutSceneYaxisCategoryorder_totaldescending    LayoutSceneYaxisCategoryorder = "total descending"
	LayoutSceneYaxisCategoryorder_minascending       LayoutSceneYaxisCategoryorder = "min ascending"
	LayoutSceneYaxisCategoryorder_mindescending      LayoutSceneYaxisCategoryorder = "min descending"
	LayoutSceneYaxisCategoryorder_maxascending       LayoutSceneYaxisCategoryorder = "max ascending"
	LayoutSceneYaxisCategoryorder_maxdescending      LayoutSceneYaxisCategoryorder = "max descending"
	LayoutSceneYaxisCategoryorder_sumascending       LayoutSceneYaxisCategoryorder = "sum ascending"
	LayoutSceneYaxisCategoryorder_sumdescending      LayoutSceneYaxisCategoryorder = "sum descending"
	LayoutSceneYaxisCategoryorder_meanascending      LayoutSceneYaxisCategoryorder = "mean ascending"
	LayoutSceneYaxisCategoryorder_meandescending     LayoutSceneYaxisCategoryorder = "mean descending"
	LayoutSceneYaxisCategoryorder_medianascending    LayoutSceneYaxisCategoryorder = "median ascending"
	LayoutSceneYaxisCategoryorder_mediandescending   LayoutSceneYaxisCategoryorder = "median descending"
)

// LayoutSceneYaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutSceneYaxisExponentformat string

const (
	LayoutSceneYaxisExponentformat_none  LayoutSceneYaxisExponentformat = "none"
	LayoutSceneYaxisExponentformat_e     LayoutSceneYaxisExponentformat = "e"
	LayoutSceneYaxisExponentformat_E     LayoutSceneYaxisExponentformat = "E"
	LayoutSceneYaxisExponentformat_power LayoutSceneYaxisExponentformat = "power"
	LayoutSceneYaxisExponentformat_SI    LayoutSceneYaxisExponentformat = "SI"
	LayoutSceneYaxisExponentformat_B     LayoutSceneYaxisExponentformat = "B"
)

// LayoutSceneYaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutSceneYaxisMirror interface{}

var (
	LayoutSceneYaxisMirror_True     LayoutSceneYaxisMirror = true
	LayoutSceneYaxisMirror_ticks    LayoutSceneYaxisMirror = "ticks"
	LayoutSceneYaxisMirror_False    LayoutSceneYaxisMirror = false
	LayoutSceneYaxisMirror_all      LayoutSceneYaxisMirror = "all"
	LayoutSceneYaxisMirror_allticks LayoutSceneYaxisMirror = "allticks"
)

// LayoutSceneYaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutSceneYaxisRangemode string

const (
	LayoutSceneYaxisRangemode_normal      LayoutSceneYaxisRangemode = "normal"
	LayoutSceneYaxisRangemode_tozero      LayoutSceneYaxisRangemode = "tozero"
	LayoutSceneYaxisRangemode_nonnegative LayoutSceneYaxisRangemode = "nonnegative"
)

// LayoutSceneYaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutSceneYaxisShowexponent string

const (
	LayoutSceneYaxisShowexponent_all   LayoutSceneYaxisShowexponent = "all"
	LayoutSceneYaxisShowexponent_first LayoutSceneYaxisShowexponent = "first"
	LayoutSceneYaxisShowexponent_last  LayoutSceneYaxisShowexponent = "last"
	LayoutSceneYaxisShowexponent_none  LayoutSceneYaxisShowexponent = "none"
)

// LayoutSceneYaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutSceneYaxisShowtickprefix string

const (
	LayoutSceneYaxisShowtickprefix_all   LayoutSceneYaxisShowtickprefix = "all"
	LayoutSceneYaxisShowtickprefix_first LayoutSceneYaxisShowtickprefix = "first"
	LayoutSceneYaxisShowtickprefix_last  LayoutSceneYaxisShowtickprefix = "last"
	LayoutSceneYaxisShowtickprefix_none  LayoutSceneYaxisShowtickprefix = "none"
)

// LayoutSceneYaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutSceneYaxisShowticksuffix string

const (
	LayoutSceneYaxisShowticksuffix_all   LayoutSceneYaxisShowticksuffix = "all"
	LayoutSceneYaxisShowticksuffix_first LayoutSceneYaxisShowticksuffix = "first"
	LayoutSceneYaxisShowticksuffix_last  LayoutSceneYaxisShowticksuffix = "last"
	LayoutSceneYaxisShowticksuffix_none  LayoutSceneYaxisShowticksuffix = "none"
)

// LayoutSceneYaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutSceneYaxisTickmode string

const (
	LayoutSceneYaxisTickmode_auto   LayoutSceneYaxisTickmode = "auto"
	LayoutSceneYaxisTickmode_linear LayoutSceneYaxisTickmode = "linear"
	LayoutSceneYaxisTickmode_array  LayoutSceneYaxisTickmode = "array"
)

// LayoutSceneYaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutSceneYaxisTicks string

const (
	LayoutSceneYaxisTicks_outside LayoutSceneYaxisTicks = "outside"
	LayoutSceneYaxisTicks_inside  LayoutSceneYaxisTicks = "inside"
)

// LayoutSceneYaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutSceneYaxisType string

const (
	LayoutSceneYaxisType__        LayoutSceneYaxisType = "-"
	LayoutSceneYaxisType_linear   LayoutSceneYaxisType = "linear"
	LayoutSceneYaxisType_log      LayoutSceneYaxisType = "log"
	LayoutSceneYaxisType_date     LayoutSceneYaxisType = "date"
	LayoutSceneYaxisType_category LayoutSceneYaxisType = "category"
)

// LayoutSceneZaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutSceneZaxisAutorange interface{}

var (
	LayoutSceneZaxisAutorange_True     LayoutSceneZaxisAutorange = true
	LayoutSceneZaxisAutorange_False    LayoutSceneZaxisAutorange = false
	LayoutSceneZaxisAutorange_reversed LayoutSceneZaxisAutorange = "reversed"
)

// LayoutSceneZaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutSceneZaxisCalendar string

const (
	LayoutSceneZaxisCalendar_gregorian  LayoutSceneZaxisCalendar = "gregorian"
	LayoutSceneZaxisCalendar_chinese    LayoutSceneZaxisCalendar = "chinese"
	LayoutSceneZaxisCalendar_coptic     LayoutSceneZaxisCalendar = "coptic"
	LayoutSceneZaxisCalendar_discworld  LayoutSceneZaxisCalendar = "discworld"
	LayoutSceneZaxisCalendar_ethiopian  LayoutSceneZaxisCalendar = "ethiopian"
	LayoutSceneZaxisCalendar_hebrew     LayoutSceneZaxisCalendar = "hebrew"
	LayoutSceneZaxisCalendar_islamic    LayoutSceneZaxisCalendar = "islamic"
	LayoutSceneZaxisCalendar_julian     LayoutSceneZaxisCalendar = "julian"
	LayoutSceneZaxisCalendar_mayan      LayoutSceneZaxisCalendar = "mayan"
	LayoutSceneZaxisCalendar_nanakshahi LayoutSceneZaxisCalendar = "nanakshahi"
	LayoutSceneZaxisCalendar_nepali     LayoutSceneZaxisCalendar = "nepali"
	LayoutSceneZaxisCalendar_persian    LayoutSceneZaxisCalendar = "persian"
	LayoutSceneZaxisCalendar_jalali     LayoutSceneZaxisCalendar = "jalali"
	LayoutSceneZaxisCalendar_taiwan     LayoutSceneZaxisCalendar = "taiwan"
	LayoutSceneZaxisCalendar_thai       LayoutSceneZaxisCalendar = "thai"
	LayoutSceneZaxisCalendar_ummalqura  LayoutSceneZaxisCalendar = "ummalqura"
)

// LayoutSceneZaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutSceneZaxisCategoryorder string

const (
	LayoutSceneZaxisCategoryorder_trace              LayoutSceneZaxisCategoryorder = "trace"
	LayoutSceneZaxisCategoryorder_categoryascending  LayoutSceneZaxisCategoryorder = "category ascending"
	LayoutSceneZaxisCategoryorder_categorydescending LayoutSceneZaxisCategoryorder = "category descending"
	LayoutSceneZaxisCategoryorder_array              LayoutSceneZaxisCategoryorder = "array"
	LayoutSceneZaxisCategoryorder_totalascending     LayoutSceneZaxisCategoryorder = "total ascending"
	LayoutSceneZaxisCategoryorder_totaldescending    LayoutSceneZaxisCategoryorder = "total descending"
	LayoutSceneZaxisCategoryorder_minascending       LayoutSceneZaxisCategoryorder = "min ascending"
	LayoutSceneZaxisCategoryorder_mindescending      LayoutSceneZaxisCategoryorder = "min descending"
	LayoutSceneZaxisCategoryorder_maxascending       LayoutSceneZaxisCategoryorder = "max ascending"
	LayoutSceneZaxisCategoryorder_maxdescending      LayoutSceneZaxisCategoryorder = "max descending"
	LayoutSceneZaxisCategoryorder_sumascending       LayoutSceneZaxisCategoryorder = "sum ascending"
	LayoutSceneZaxisCategoryorder_sumdescending      LayoutSceneZaxisCategoryorder = "sum descending"
	LayoutSceneZaxisCategoryorder_meanascending      LayoutSceneZaxisCategoryorder = "mean ascending"
	LayoutSceneZaxisCategoryorder_meandescending     LayoutSceneZaxisCategoryorder = "mean descending"
	LayoutSceneZaxisCategoryorder_medianascending    LayoutSceneZaxisCategoryorder = "median ascending"
	LayoutSceneZaxisCategoryorder_mediandescending   LayoutSceneZaxisCategoryorder = "median descending"
)

// LayoutSceneZaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutSceneZaxisExponentformat string

const (
	LayoutSceneZaxisExponentformat_none  LayoutSceneZaxisExponentformat = "none"
	LayoutSceneZaxisExponentformat_e     LayoutSceneZaxisExponentformat = "e"
	LayoutSceneZaxisExponentformat_E     LayoutSceneZaxisExponentformat = "E"
	LayoutSceneZaxisExponentformat_power LayoutSceneZaxisExponentformat = "power"
	LayoutSceneZaxisExponentformat_SI    LayoutSceneZaxisExponentformat = "SI"
	LayoutSceneZaxisExponentformat_B     LayoutSceneZaxisExponentformat = "B"
)

// LayoutSceneZaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutSceneZaxisMirror interface{}

var (
	LayoutSceneZaxisMirror_True     LayoutSceneZaxisMirror = true
	LayoutSceneZaxisMirror_ticks    LayoutSceneZaxisMirror = "ticks"
	LayoutSceneZaxisMirror_False    LayoutSceneZaxisMirror = false
	LayoutSceneZaxisMirror_all      LayoutSceneZaxisMirror = "all"
	LayoutSceneZaxisMirror_allticks LayoutSceneZaxisMirror = "allticks"
)

// LayoutSceneZaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutSceneZaxisRangemode string

const (
	LayoutSceneZaxisRangemode_normal      LayoutSceneZaxisRangemode = "normal"
	LayoutSceneZaxisRangemode_tozero      LayoutSceneZaxisRangemode = "tozero"
	LayoutSceneZaxisRangemode_nonnegative LayoutSceneZaxisRangemode = "nonnegative"
)

// LayoutSceneZaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutSceneZaxisShowexponent string

const (
	LayoutSceneZaxisShowexponent_all   LayoutSceneZaxisShowexponent = "all"
	LayoutSceneZaxisShowexponent_first LayoutSceneZaxisShowexponent = "first"
	LayoutSceneZaxisShowexponent_last  LayoutSceneZaxisShowexponent = "last"
	LayoutSceneZaxisShowexponent_none  LayoutSceneZaxisShowexponent = "none"
)

// LayoutSceneZaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutSceneZaxisShowtickprefix string

const (
	LayoutSceneZaxisShowtickprefix_all   LayoutSceneZaxisShowtickprefix = "all"
	LayoutSceneZaxisShowtickprefix_first LayoutSceneZaxisShowtickprefix = "first"
	LayoutSceneZaxisShowtickprefix_last  LayoutSceneZaxisShowtickprefix = "last"
	LayoutSceneZaxisShowtickprefix_none  LayoutSceneZaxisShowtickprefix = "none"
)

// LayoutSceneZaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutSceneZaxisShowticksuffix string

const (
	LayoutSceneZaxisShowticksuffix_all   LayoutSceneZaxisShowticksuffix = "all"
	LayoutSceneZaxisShowticksuffix_first LayoutSceneZaxisShowticksuffix = "first"
	LayoutSceneZaxisShowticksuffix_last  LayoutSceneZaxisShowticksuffix = "last"
	LayoutSceneZaxisShowticksuffix_none  LayoutSceneZaxisShowticksuffix = "none"
)

// LayoutSceneZaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutSceneZaxisTickmode string

const (
	LayoutSceneZaxisTickmode_auto   LayoutSceneZaxisTickmode = "auto"
	LayoutSceneZaxisTickmode_linear LayoutSceneZaxisTickmode = "linear"
	LayoutSceneZaxisTickmode_array  LayoutSceneZaxisTickmode = "array"
)

// LayoutSceneZaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutSceneZaxisTicks string

const (
	LayoutSceneZaxisTicks_outside LayoutSceneZaxisTicks = "outside"
	LayoutSceneZaxisTicks_inside  LayoutSceneZaxisTicks = "inside"
)

// LayoutSceneZaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutSceneZaxisType string

const (
	LayoutSceneZaxisType__        LayoutSceneZaxisType = "-"
	LayoutSceneZaxisType_linear   LayoutSceneZaxisType = "linear"
	LayoutSceneZaxisType_log      LayoutSceneZaxisType = "log"
	LayoutSceneZaxisType_date     LayoutSceneZaxisType = "date"
	LayoutSceneZaxisType_category LayoutSceneZaxisType = "category"
)

// LayoutSelectdirection When `dragmode` is set to *select*, this limits the selection of the drag to horizontal, vertical or diagonal. *h* only allows horizontal selection, *v* only vertical, *d* only diagonal and *any* sets no limit.
type LayoutSelectdirection string

const (
	LayoutSelectdirection_h   LayoutSelectdirection = "h"
	LayoutSelectdirection_v   LayoutSelectdirection = "v"
	LayoutSelectdirection_d   LayoutSelectdirection = "d"
	LayoutSelectdirection_any LayoutSelectdirection = "any"
)

// LayoutTernaryAaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutTernaryAaxisExponentformat string

const (
	LayoutTernaryAaxisExponentformat_none  LayoutTernaryAaxisExponentformat = "none"
	LayoutTernaryAaxisExponentformat_e     LayoutTernaryAaxisExponentformat = "e"
	LayoutTernaryAaxisExponentformat_E     LayoutTernaryAaxisExponentformat = "E"
	LayoutTernaryAaxisExponentformat_power LayoutTernaryAaxisExponentformat = "power"
	LayoutTernaryAaxisExponentformat_SI    LayoutTernaryAaxisExponentformat = "SI"
	LayoutTernaryAaxisExponentformat_B     LayoutTernaryAaxisExponentformat = "B"
)

// LayoutTernaryAaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutTernaryAaxisLayer string

const (
	LayoutTernaryAaxisLayer_abovetraces LayoutTernaryAaxisLayer = "above traces"
	LayoutTernaryAaxisLayer_belowtraces LayoutTernaryAaxisLayer = "below traces"
)

// LayoutTernaryAaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutTernaryAaxisShowexponent string

const (
	LayoutTernaryAaxisShowexponent_all   LayoutTernaryAaxisShowexponent = "all"
	LayoutTernaryAaxisShowexponent_first LayoutTernaryAaxisShowexponent = "first"
	LayoutTernaryAaxisShowexponent_last  LayoutTernaryAaxisShowexponent = "last"
	LayoutTernaryAaxisShowexponent_none  LayoutTernaryAaxisShowexponent = "none"
)

// LayoutTernaryAaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutTernaryAaxisShowtickprefix string

const (
	LayoutTernaryAaxisShowtickprefix_all   LayoutTernaryAaxisShowtickprefix = "all"
	LayoutTernaryAaxisShowtickprefix_first LayoutTernaryAaxisShowtickprefix = "first"
	LayoutTernaryAaxisShowtickprefix_last  LayoutTernaryAaxisShowtickprefix = "last"
	LayoutTernaryAaxisShowtickprefix_none  LayoutTernaryAaxisShowtickprefix = "none"
)

// LayoutTernaryAaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutTernaryAaxisShowticksuffix string

const (
	LayoutTernaryAaxisShowticksuffix_all   LayoutTernaryAaxisShowticksuffix = "all"
	LayoutTernaryAaxisShowticksuffix_first LayoutTernaryAaxisShowticksuffix = "first"
	LayoutTernaryAaxisShowticksuffix_last  LayoutTernaryAaxisShowticksuffix = "last"
	LayoutTernaryAaxisShowticksuffix_none  LayoutTernaryAaxisShowticksuffix = "none"
)

// LayoutTernaryAaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutTernaryAaxisTickmode string

const (
	LayoutTernaryAaxisTickmode_auto   LayoutTernaryAaxisTickmode = "auto"
	LayoutTernaryAaxisTickmode_linear LayoutTernaryAaxisTickmode = "linear"
	LayoutTernaryAaxisTickmode_array  LayoutTernaryAaxisTickmode = "array"
)

// LayoutTernaryAaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutTernaryAaxisTicks string

const (
	LayoutTernaryAaxisTicks_outside LayoutTernaryAaxisTicks = "outside"
	LayoutTernaryAaxisTicks_inside  LayoutTernaryAaxisTicks = "inside"
)

// LayoutTernaryBaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutTernaryBaxisExponentformat string

const (
	LayoutTernaryBaxisExponentformat_none  LayoutTernaryBaxisExponentformat = "none"
	LayoutTernaryBaxisExponentformat_e     LayoutTernaryBaxisExponentformat = "e"
	LayoutTernaryBaxisExponentformat_E     LayoutTernaryBaxisExponentformat = "E"
	LayoutTernaryBaxisExponentformat_power LayoutTernaryBaxisExponentformat = "power"
	LayoutTernaryBaxisExponentformat_SI    LayoutTernaryBaxisExponentformat = "SI"
	LayoutTernaryBaxisExponentformat_B     LayoutTernaryBaxisExponentformat = "B"
)

// LayoutTernaryBaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutTernaryBaxisLayer string

const (
	LayoutTernaryBaxisLayer_abovetraces LayoutTernaryBaxisLayer = "above traces"
	LayoutTernaryBaxisLayer_belowtraces LayoutTernaryBaxisLayer = "below traces"
)

// LayoutTernaryBaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutTernaryBaxisShowexponent string

const (
	LayoutTernaryBaxisShowexponent_all   LayoutTernaryBaxisShowexponent = "all"
	LayoutTernaryBaxisShowexponent_first LayoutTernaryBaxisShowexponent = "first"
	LayoutTernaryBaxisShowexponent_last  LayoutTernaryBaxisShowexponent = "last"
	LayoutTernaryBaxisShowexponent_none  LayoutTernaryBaxisShowexponent = "none"
)

// LayoutTernaryBaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutTernaryBaxisShowtickprefix string

const (
	LayoutTernaryBaxisShowtickprefix_all   LayoutTernaryBaxisShowtickprefix = "all"
	LayoutTernaryBaxisShowtickprefix_first LayoutTernaryBaxisShowtickprefix = "first"
	LayoutTernaryBaxisShowtickprefix_last  LayoutTernaryBaxisShowtickprefix = "last"
	LayoutTernaryBaxisShowtickprefix_none  LayoutTernaryBaxisShowtickprefix = "none"
)

// LayoutTernaryBaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutTernaryBaxisShowticksuffix string

const (
	LayoutTernaryBaxisShowticksuffix_all   LayoutTernaryBaxisShowticksuffix = "all"
	LayoutTernaryBaxisShowticksuffix_first LayoutTernaryBaxisShowticksuffix = "first"
	LayoutTernaryBaxisShowticksuffix_last  LayoutTernaryBaxisShowticksuffix = "last"
	LayoutTernaryBaxisShowticksuffix_none  LayoutTernaryBaxisShowticksuffix = "none"
)

// LayoutTernaryBaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutTernaryBaxisTickmode string

const (
	LayoutTernaryBaxisTickmode_auto   LayoutTernaryBaxisTickmode = "auto"
	LayoutTernaryBaxisTickmode_linear LayoutTernaryBaxisTickmode = "linear"
	LayoutTernaryBaxisTickmode_array  LayoutTernaryBaxisTickmode = "array"
)

// LayoutTernaryBaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutTernaryBaxisTicks string

const (
	LayoutTernaryBaxisTicks_outside LayoutTernaryBaxisTicks = "outside"
	LayoutTernaryBaxisTicks_inside  LayoutTernaryBaxisTicks = "inside"
)

// LayoutTernaryCaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutTernaryCaxisExponentformat string

const (
	LayoutTernaryCaxisExponentformat_none  LayoutTernaryCaxisExponentformat = "none"
	LayoutTernaryCaxisExponentformat_e     LayoutTernaryCaxisExponentformat = "e"
	LayoutTernaryCaxisExponentformat_E     LayoutTernaryCaxisExponentformat = "E"
	LayoutTernaryCaxisExponentformat_power LayoutTernaryCaxisExponentformat = "power"
	LayoutTernaryCaxisExponentformat_SI    LayoutTernaryCaxisExponentformat = "SI"
	LayoutTernaryCaxisExponentformat_B     LayoutTernaryCaxisExponentformat = "B"
)

// LayoutTernaryCaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutTernaryCaxisLayer string

const (
	LayoutTernaryCaxisLayer_abovetraces LayoutTernaryCaxisLayer = "above traces"
	LayoutTernaryCaxisLayer_belowtraces LayoutTernaryCaxisLayer = "below traces"
)

// LayoutTernaryCaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutTernaryCaxisShowexponent string

const (
	LayoutTernaryCaxisShowexponent_all   LayoutTernaryCaxisShowexponent = "all"
	LayoutTernaryCaxisShowexponent_first LayoutTernaryCaxisShowexponent = "first"
	LayoutTernaryCaxisShowexponent_last  LayoutTernaryCaxisShowexponent = "last"
	LayoutTernaryCaxisShowexponent_none  LayoutTernaryCaxisShowexponent = "none"
)

// LayoutTernaryCaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutTernaryCaxisShowtickprefix string

const (
	LayoutTernaryCaxisShowtickprefix_all   LayoutTernaryCaxisShowtickprefix = "all"
	LayoutTernaryCaxisShowtickprefix_first LayoutTernaryCaxisShowtickprefix = "first"
	LayoutTernaryCaxisShowtickprefix_last  LayoutTernaryCaxisShowtickprefix = "last"
	LayoutTernaryCaxisShowtickprefix_none  LayoutTernaryCaxisShowtickprefix = "none"
)

// LayoutTernaryCaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutTernaryCaxisShowticksuffix string

const (
	LayoutTernaryCaxisShowticksuffix_all   LayoutTernaryCaxisShowticksuffix = "all"
	LayoutTernaryCaxisShowticksuffix_first LayoutTernaryCaxisShowticksuffix = "first"
	LayoutTernaryCaxisShowticksuffix_last  LayoutTernaryCaxisShowticksuffix = "last"
	LayoutTernaryCaxisShowticksuffix_none  LayoutTernaryCaxisShowticksuffix = "none"
)

// LayoutTernaryCaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutTernaryCaxisTickmode string

const (
	LayoutTernaryCaxisTickmode_auto   LayoutTernaryCaxisTickmode = "auto"
	LayoutTernaryCaxisTickmode_linear LayoutTernaryCaxisTickmode = "linear"
	LayoutTernaryCaxisTickmode_array  LayoutTernaryCaxisTickmode = "array"
)

// LayoutTernaryCaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutTernaryCaxisTicks string

const (
	LayoutTernaryCaxisTicks_outside LayoutTernaryCaxisTicks = "outside"
	LayoutTernaryCaxisTicks_inside  LayoutTernaryCaxisTicks = "inside"
)

// LayoutTitleXanchor Sets the title's horizontal alignment with respect to its x position. *left* means that the title starts at x, *right* means that the title ends at x and *center* means that the title's center is at x. *auto* divides `xref` by three and calculates the `xanchor` value automatically based on the value of `x`.
type LayoutTitleXanchor string

const (
	LayoutTitleXanchor_auto   LayoutTitleXanchor = "auto"
	LayoutTitleXanchor_left   LayoutTitleXanchor = "left"
	LayoutTitleXanchor_center LayoutTitleXanchor = "center"
	LayoutTitleXanchor_right  LayoutTitleXanchor = "right"
)

// LayoutTitleXref Sets the container `x` refers to. *container* spans the entire `width` of the plot. *paper* refers to the width of the plotting area only.
type LayoutTitleXref string

const (
	LayoutTitleXref_container LayoutTitleXref = "container"
	LayoutTitleXref_paper     LayoutTitleXref = "paper"
)

// LayoutTitleYanchor Sets the title's vertical alignment with respect to its y position. *top* means that the title's cap line is at y, *bottom* means that the title's baseline is at y and *middle* means that the title's midline is at y. *auto* divides `yref` by three and calculates the `yanchor` value automatically based on the value of `y`.
type LayoutTitleYanchor string

const (
	LayoutTitleYanchor_auto   LayoutTitleYanchor = "auto"
	LayoutTitleYanchor_top    LayoutTitleYanchor = "top"
	LayoutTitleYanchor_middle LayoutTitleYanchor = "middle"
	LayoutTitleYanchor_bottom LayoutTitleYanchor = "bottom"
)

// LayoutTitleYref Sets the container `y` refers to. *container* spans the entire `height` of the plot. *paper* refers to the height of the plotting area only.
type LayoutTitleYref string

const (
	LayoutTitleYref_container LayoutTitleYref = "container"
	LayoutTitleYref_paper     LayoutTitleYref = "paper"
)

// LayoutTransitionEasing The easing function used for the transition
type LayoutTransitionEasing string

const (
	LayoutTransitionEasing_linear         LayoutTransitionEasing = "linear"
	LayoutTransitionEasing_quad           LayoutTransitionEasing = "quad"
	LayoutTransitionEasing_cubic          LayoutTransitionEasing = "cubic"
	LayoutTransitionEasing_sin            LayoutTransitionEasing = "sin"
	LayoutTransitionEasing_exp            LayoutTransitionEasing = "exp"
	LayoutTransitionEasing_circle         LayoutTransitionEasing = "circle"
	LayoutTransitionEasing_elastic        LayoutTransitionEasing = "elastic"
	LayoutTransitionEasing_back           LayoutTransitionEasing = "back"
	LayoutTransitionEasing_bounce         LayoutTransitionEasing = "bounce"
	LayoutTransitionEasing_linear_in      LayoutTransitionEasing = "linear-in"
	LayoutTransitionEasing_quad_in        LayoutTransitionEasing = "quad-in"
	LayoutTransitionEasing_cubic_in       LayoutTransitionEasing = "cubic-in"
	LayoutTransitionEasing_sin_in         LayoutTransitionEasing = "sin-in"
	LayoutTransitionEasing_exp_in         LayoutTransitionEasing = "exp-in"
	LayoutTransitionEasing_circle_in      LayoutTransitionEasing = "circle-in"
	LayoutTransitionEasing_elastic_in     LayoutTransitionEasing = "elastic-in"
	LayoutTransitionEasing_back_in        LayoutTransitionEasing = "back-in"
	LayoutTransitionEasing_bounce_in      LayoutTransitionEasing = "bounce-in"
	LayoutTransitionEasing_linear_out     LayoutTransitionEasing = "linear-out"
	LayoutTransitionEasing_quad_out       LayoutTransitionEasing = "quad-out"
	LayoutTransitionEasing_cubic_out      LayoutTransitionEasing = "cubic-out"
	LayoutTransitionEasing_sin_out        LayoutTransitionEasing = "sin-out"
	LayoutTransitionEasing_exp_out        LayoutTransitionEasing = "exp-out"
	LayoutTransitionEasing_circle_out     LayoutTransitionEasing = "circle-out"
	LayoutTransitionEasing_elastic_out    LayoutTransitionEasing = "elastic-out"
	LayoutTransitionEasing_back_out       LayoutTransitionEasing = "back-out"
	LayoutTransitionEasing_bounce_out     LayoutTransitionEasing = "bounce-out"
	LayoutTransitionEasing_linear_in_out  LayoutTransitionEasing = "linear-in-out"
	LayoutTransitionEasing_quad_in_out    LayoutTransitionEasing = "quad-in-out"
	LayoutTransitionEasing_cubic_in_out   LayoutTransitionEasing = "cubic-in-out"
	LayoutTransitionEasing_sin_in_out     LayoutTransitionEasing = "sin-in-out"
	LayoutTransitionEasing_exp_in_out     LayoutTransitionEasing = "exp-in-out"
	LayoutTransitionEasing_circle_in_out  LayoutTransitionEasing = "circle-in-out"
	LayoutTransitionEasing_elastic_in_out LayoutTransitionEasing = "elastic-in-out"
	LayoutTransitionEasing_back_in_out    LayoutTransitionEasing = "back-in-out"
	LayoutTransitionEasing_bounce_in_out  LayoutTransitionEasing = "bounce-in-out"
)

// LayoutTransitionOrdering Determines whether the figure's layout or traces smoothly transitions during updates that make both traces and layout change.
type LayoutTransitionOrdering string

const (
	LayoutTransitionOrdering_layoutfirst LayoutTransitionOrdering = "layout first"
	LayoutTransitionOrdering_tracesfirst LayoutTransitionOrdering = "traces first"
)

// LayoutUniformtextMode Determines how the font size for various text elements are uniformed between each trace type. If the computed text sizes were smaller than the minimum size defined by `uniformtext.minsize` using *hide* option hides the text; and using *show* option shows the text without further downscaling. Please note that if the size defined by `minsize` is greater than the font size defined by trace, then the `minsize` is used.
type LayoutUniformtextMode interface{}

var (
	LayoutUniformtextMode_False LayoutUniformtextMode = false
	LayoutUniformtextMode_hide  LayoutUniformtextMode = "hide"
	LayoutUniformtextMode_show  LayoutUniformtextMode = "show"
)

// LayoutViolinmode Determines how violins at the same location coordinate are displayed on the graph. If *group*, the violins are plotted next to one another centered around the shared location. If *overlay*, the violins are plotted over one another, you might need to set *opacity* to see them multiple violins. Has no effect on traces that have *width* set.
type LayoutViolinmode string

const (
	LayoutViolinmode_group   LayoutViolinmode = "group"
	LayoutViolinmode_overlay LayoutViolinmode = "overlay"
)

// LayoutWaterfallmode Determines how bars at the same location coordinate are displayed on the graph. With *group*, the bars are plotted next to one another centered around the shared location. With *overlay*, the bars are plotted over one another, you might need to an *opacity* to see multiple bars.
type LayoutWaterfallmode string

const (
	LayoutWaterfallmode_group   LayoutWaterfallmode = "group"
	LayoutWaterfallmode_overlay LayoutWaterfallmode = "overlay"
)

// LayoutXaxisAnchor If set to an opposite-letter axis id (e.g. `x2`, `y`), this axis is bound to the corresponding opposite-letter axis. If set to *free*, this axis' position is determined by `position`.
type LayoutXaxisAnchor string

const (
	LayoutXaxisAnchor_free                                                                                                                       LayoutXaxisAnchor = "free"
	LayoutXaxisAnchor_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisAnchor = "/^x([2-9]|[1-9][0-9]+)?$/"
	LayoutXaxisAnchor_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisAnchor = "/^y([2-9]|[1-9][0-9]+)?$/"
)

// LayoutXaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutXaxisAutorange interface{}

var (
	LayoutXaxisAutorange_True     LayoutXaxisAutorange = true
	LayoutXaxisAutorange_False    LayoutXaxisAutorange = false
	LayoutXaxisAutorange_reversed LayoutXaxisAutorange = "reversed"
)

// LayoutXaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutXaxisCalendar string

const (
	LayoutXaxisCalendar_gregorian  LayoutXaxisCalendar = "gregorian"
	LayoutXaxisCalendar_chinese    LayoutXaxisCalendar = "chinese"
	LayoutXaxisCalendar_coptic     LayoutXaxisCalendar = "coptic"
	LayoutXaxisCalendar_discworld  LayoutXaxisCalendar = "discworld"
	LayoutXaxisCalendar_ethiopian  LayoutXaxisCalendar = "ethiopian"
	LayoutXaxisCalendar_hebrew     LayoutXaxisCalendar = "hebrew"
	LayoutXaxisCalendar_islamic    LayoutXaxisCalendar = "islamic"
	LayoutXaxisCalendar_julian     LayoutXaxisCalendar = "julian"
	LayoutXaxisCalendar_mayan      LayoutXaxisCalendar = "mayan"
	LayoutXaxisCalendar_nanakshahi LayoutXaxisCalendar = "nanakshahi"
	LayoutXaxisCalendar_nepali     LayoutXaxisCalendar = "nepali"
	LayoutXaxisCalendar_persian    LayoutXaxisCalendar = "persian"
	LayoutXaxisCalendar_jalali     LayoutXaxisCalendar = "jalali"
	LayoutXaxisCalendar_taiwan     LayoutXaxisCalendar = "taiwan"
	LayoutXaxisCalendar_thai       LayoutXaxisCalendar = "thai"
	LayoutXaxisCalendar_ummalqura  LayoutXaxisCalendar = "ummalqura"
)

// LayoutXaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutXaxisCategoryorder string

const (
	LayoutXaxisCategoryorder_trace              LayoutXaxisCategoryorder = "trace"
	LayoutXaxisCategoryorder_categoryascending  LayoutXaxisCategoryorder = "category ascending"
	LayoutXaxisCategoryorder_categorydescending LayoutXaxisCategoryorder = "category descending"
	LayoutXaxisCategoryorder_array              LayoutXaxisCategoryorder = "array"
	LayoutXaxisCategoryorder_totalascending     LayoutXaxisCategoryorder = "total ascending"
	LayoutXaxisCategoryorder_totaldescending    LayoutXaxisCategoryorder = "total descending"
	LayoutXaxisCategoryorder_minascending       LayoutXaxisCategoryorder = "min ascending"
	LayoutXaxisCategoryorder_mindescending      LayoutXaxisCategoryorder = "min descending"
	LayoutXaxisCategoryorder_maxascending       LayoutXaxisCategoryorder = "max ascending"
	LayoutXaxisCategoryorder_maxdescending      LayoutXaxisCategoryorder = "max descending"
	LayoutXaxisCategoryorder_sumascending       LayoutXaxisCategoryorder = "sum ascending"
	LayoutXaxisCategoryorder_sumdescending      LayoutXaxisCategoryorder = "sum descending"
	LayoutXaxisCategoryorder_meanascending      LayoutXaxisCategoryorder = "mean ascending"
	LayoutXaxisCategoryorder_meandescending     LayoutXaxisCategoryorder = "mean descending"
	LayoutXaxisCategoryorder_medianascending    LayoutXaxisCategoryorder = "median ascending"
	LayoutXaxisCategoryorder_mediandescending   LayoutXaxisCategoryorder = "median descending"
)

// LayoutXaxisConstrain If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines how that happens: by increasing the *range* (default), or by decreasing the *domain*.
type LayoutXaxisConstrain string

const (
	LayoutXaxisConstrain_range  LayoutXaxisConstrain = "range"
	LayoutXaxisConstrain_domain LayoutXaxisConstrain = "domain"
)

// LayoutXaxisConstraintoward If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines which direction we push the originally specified plot area. Options are *left*, *center* (default), and *right* for x axes, and *top*, *middle* (default), and *bottom* for y axes.
type LayoutXaxisConstraintoward string

const (
	LayoutXaxisConstraintoward_left   LayoutXaxisConstraintoward = "left"
	LayoutXaxisConstraintoward_center LayoutXaxisConstraintoward = "center"
	LayoutXaxisConstraintoward_right  LayoutXaxisConstraintoward = "right"
	LayoutXaxisConstraintoward_top    LayoutXaxisConstraintoward = "top"
	LayoutXaxisConstraintoward_middle LayoutXaxisConstraintoward = "middle"
	LayoutXaxisConstraintoward_bottom LayoutXaxisConstraintoward = "bottom"
)

// LayoutXaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutXaxisExponentformat string

const (
	LayoutXaxisExponentformat_none  LayoutXaxisExponentformat = "none"
	LayoutXaxisExponentformat_e     LayoutXaxisExponentformat = "e"
	LayoutXaxisExponentformat_E     LayoutXaxisExponentformat = "E"
	LayoutXaxisExponentformat_power LayoutXaxisExponentformat = "power"
	LayoutXaxisExponentformat_SI    LayoutXaxisExponentformat = "SI"
	LayoutXaxisExponentformat_B     LayoutXaxisExponentformat = "B"
)

// LayoutXaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutXaxisLayer string

const (
	LayoutXaxisLayer_abovetraces LayoutXaxisLayer = "above traces"
	LayoutXaxisLayer_belowtraces LayoutXaxisLayer = "below traces"
)

// LayoutXaxisMatches If set to another axis id (e.g. `x2`, `y`), the range of this axis will match the range of the corresponding axis in data-coordinates space. Moreover, matching axes share auto-range values, category lists and histogram auto-bins. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden. Moreover, note that matching axes must have the same `type`.
type LayoutXaxisMatches string

const (
	LayoutXaxisMatches_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisMatches = "/^x([2-9]|[1-9][0-9]+)?$/"
	LayoutXaxisMatches_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisMatches = "/^y([2-9]|[1-9][0-9]+)?$/"
)

// LayoutXaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutXaxisMirror interface{}

var (
	LayoutXaxisMirror_True     LayoutXaxisMirror = true
	LayoutXaxisMirror_ticks    LayoutXaxisMirror = "ticks"
	LayoutXaxisMirror_False    LayoutXaxisMirror = false
	LayoutXaxisMirror_all      LayoutXaxisMirror = "all"
	LayoutXaxisMirror_allticks LayoutXaxisMirror = "allticks"
)

// LayoutXaxisOverlaying If set a same-letter axis id, this axis is overlaid on top of the corresponding same-letter axis, with traces and axes visible for both axes. If *false*, this axis does not overlay any same-letter axes. In this case, for axes with overlapping domains only the highest-numbered axis will be visible.
type LayoutXaxisOverlaying string

const (
	LayoutXaxisOverlaying_free                                                                                                                       LayoutXaxisOverlaying = "free"
	LayoutXaxisOverlaying_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisOverlaying = "/^x([2-9]|[1-9][0-9]+)?$/"
	LayoutXaxisOverlaying_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisOverlaying = "/^y([2-9]|[1-9][0-9]+)?$/"
)

// LayoutXaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutXaxisRangemode string

const (
	LayoutXaxisRangemode_normal      LayoutXaxisRangemode = "normal"
	LayoutXaxisRangemode_tozero      LayoutXaxisRangemode = "tozero"
	LayoutXaxisRangemode_nonnegative LayoutXaxisRangemode = "nonnegative"
)

// LayoutXaxisRangeselectorXanchor Sets the range selector's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the range selector.
type LayoutXaxisRangeselectorXanchor string

const (
	LayoutXaxisRangeselectorXanchor_auto   LayoutXaxisRangeselectorXanchor = "auto"
	LayoutXaxisRangeselectorXanchor_left   LayoutXaxisRangeselectorXanchor = "left"
	LayoutXaxisRangeselectorXanchor_center LayoutXaxisRangeselectorXanchor = "center"
	LayoutXaxisRangeselectorXanchor_right  LayoutXaxisRangeselectorXanchor = "right"
)

// LayoutXaxisRangeselectorYanchor Sets the range selector's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the range selector.
type LayoutXaxisRangeselectorYanchor string

const (
	LayoutXaxisRangeselectorYanchor_auto   LayoutXaxisRangeselectorYanchor = "auto"
	LayoutXaxisRangeselectorYanchor_top    LayoutXaxisRangeselectorYanchor = "top"
	LayoutXaxisRangeselectorYanchor_middle LayoutXaxisRangeselectorYanchor = "middle"
	LayoutXaxisRangeselectorYanchor_bottom LayoutXaxisRangeselectorYanchor = "bottom"
)

// LayoutXaxisRangesliderYaxisRangemode Determines whether or not the range of this axis in the rangeslider use the same value than in the main plot when zooming in/out. If *auto*, the autorange will be used. If *fixed*, the `range` is used. If *match*, the current range of the corresponding y-axis on the main subplot is used.
type LayoutXaxisRangesliderYaxisRangemode string

const (
	LayoutXaxisRangesliderYaxisRangemode_auto  LayoutXaxisRangesliderYaxisRangemode = "auto"
	LayoutXaxisRangesliderYaxisRangemode_fixed LayoutXaxisRangesliderYaxisRangemode = "fixed"
	LayoutXaxisRangesliderYaxisRangemode_match LayoutXaxisRangesliderYaxisRangemode = "match"
)

// LayoutXaxisScaleanchor If set to another axis id (e.g. `x2`, `y`), the range of this axis changes together with the range of the corresponding axis such that the scale of pixels per unit is in a constant ratio. Both axes are still zoomable, but when you zoom one, the other will zoom the same amount, keeping a fixed midpoint. `constrain` and `constraintoward` determine how we enforce the constraint. You can chain these, ie `yaxis: {scaleanchor: *x*}, xaxis2: {scaleanchor: *y*}` but you can only link axes of the same `type`. The linked axis can have the opposite letter (to constrain the aspect ratio) or the same letter (to match scales across subplots). Loops (`yaxis: {scaleanchor: *x*}, xaxis: {scaleanchor: *y*}` or longer) are redundant and the last constraint encountered will be ignored to avoid possible inconsistent constraints via `scaleratio`. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden.
type LayoutXaxisScaleanchor string

const (
	LayoutXaxisScaleanchor_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisScaleanchor = "/^x([2-9]|[1-9][0-9]+)?$/"
	LayoutXaxisScaleanchor_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisScaleanchor = "/^y([2-9]|[1-9][0-9]+)?$/"
)

// LayoutXaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutXaxisShowexponent string

const (
	LayoutXaxisShowexponent_all   LayoutXaxisShowexponent = "all"
	LayoutXaxisShowexponent_first LayoutXaxisShowexponent = "first"
	LayoutXaxisShowexponent_last  LayoutXaxisShowexponent = "last"
	LayoutXaxisShowexponent_none  LayoutXaxisShowexponent = "none"
)

// LayoutXaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutXaxisShowtickprefix string

const (
	LayoutXaxisShowtickprefix_all   LayoutXaxisShowtickprefix = "all"
	LayoutXaxisShowtickprefix_first LayoutXaxisShowtickprefix = "first"
	LayoutXaxisShowtickprefix_last  LayoutXaxisShowtickprefix = "last"
	LayoutXaxisShowtickprefix_none  LayoutXaxisShowtickprefix = "none"
)

// LayoutXaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutXaxisShowticksuffix string

const (
	LayoutXaxisShowticksuffix_all   LayoutXaxisShowticksuffix = "all"
	LayoutXaxisShowticksuffix_first LayoutXaxisShowticksuffix = "first"
	LayoutXaxisShowticksuffix_last  LayoutXaxisShowticksuffix = "last"
	LayoutXaxisShowticksuffix_none  LayoutXaxisShowticksuffix = "none"
)

// LayoutXaxisSide Determines whether a x (y) axis is positioned at the *bottom* (*left*) or *top* (*right*) of the plotting area.
type LayoutXaxisSide string

const (
	LayoutXaxisSide_top    LayoutXaxisSide = "top"
	LayoutXaxisSide_bottom LayoutXaxisSide = "bottom"
	LayoutXaxisSide_left   LayoutXaxisSide = "left"
	LayoutXaxisSide_right  LayoutXaxisSide = "right"
)

// LayoutXaxisSpikesnap Determines whether spikelines are stuck to the cursor or to the closest datapoints.
type LayoutXaxisSpikesnap string

const (
	LayoutXaxisSpikesnap_data        LayoutXaxisSpikesnap = "data"
	LayoutXaxisSpikesnap_cursor      LayoutXaxisSpikesnap = "cursor"
	LayoutXaxisSpikesnap_hovereddata LayoutXaxisSpikesnap = "hovered data"
)

// LayoutXaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutXaxisTickmode string

const (
	LayoutXaxisTickmode_auto   LayoutXaxisTickmode = "auto"
	LayoutXaxisTickmode_linear LayoutXaxisTickmode = "linear"
	LayoutXaxisTickmode_array  LayoutXaxisTickmode = "array"
)

// LayoutXaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutXaxisTicks string

const (
	LayoutXaxisTicks_outside LayoutXaxisTicks = "outside"
	LayoutXaxisTicks_inside  LayoutXaxisTicks = "inside"
)

// LayoutXaxisTickson Determines where ticks and grid lines are drawn with respect to their corresponding tick labels. Only has an effect for axes of `type` *category* or *multicategory*. When set to *boundaries*, ticks and grid lines are drawn half a category to the left/bottom of labels.
type LayoutXaxisTickson string

const (
	LayoutXaxisTickson_labels     LayoutXaxisTickson = "labels"
	LayoutXaxisTickson_boundaries LayoutXaxisTickson = "boundaries"
)

// LayoutXaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutXaxisType string

const (
	LayoutXaxisType__             LayoutXaxisType = "-"
	LayoutXaxisType_linear        LayoutXaxisType = "linear"
	LayoutXaxisType_log           LayoutXaxisType = "log"
	LayoutXaxisType_date          LayoutXaxisType = "date"
	LayoutXaxisType_category      LayoutXaxisType = "category"
	LayoutXaxisType_multicategory LayoutXaxisType = "multicategory"
)

// LayoutYaxisAnchor If set to an opposite-letter axis id (e.g. `x2`, `y`), this axis is bound to the corresponding opposite-letter axis. If set to *free*, this axis' position is determined by `position`.
type LayoutYaxisAnchor string

const (
	LayoutYaxisAnchor_free                                                                                                                       LayoutYaxisAnchor = "free"
	LayoutYaxisAnchor_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisAnchor = "/^x([2-9]|[1-9][0-9]+)?$/"
	LayoutYaxisAnchor_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisAnchor = "/^y([2-9]|[1-9][0-9]+)?$/"
)

// LayoutYaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutYaxisAutorange interface{}

var (
	LayoutYaxisAutorange_True     LayoutYaxisAutorange = true
	LayoutYaxisAutorange_False    LayoutYaxisAutorange = false
	LayoutYaxisAutorange_reversed LayoutYaxisAutorange = "reversed"
)

// LayoutYaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutYaxisCalendar string

const (
	LayoutYaxisCalendar_gregorian  LayoutYaxisCalendar = "gregorian"
	LayoutYaxisCalendar_chinese    LayoutYaxisCalendar = "chinese"
	LayoutYaxisCalendar_coptic     LayoutYaxisCalendar = "coptic"
	LayoutYaxisCalendar_discworld  LayoutYaxisCalendar = "discworld"
	LayoutYaxisCalendar_ethiopian  LayoutYaxisCalendar = "ethiopian"
	LayoutYaxisCalendar_hebrew     LayoutYaxisCalendar = "hebrew"
	LayoutYaxisCalendar_islamic    LayoutYaxisCalendar = "islamic"
	LayoutYaxisCalendar_julian     LayoutYaxisCalendar = "julian"
	LayoutYaxisCalendar_mayan      LayoutYaxisCalendar = "mayan"
	LayoutYaxisCalendar_nanakshahi LayoutYaxisCalendar = "nanakshahi"
	LayoutYaxisCalendar_nepali     LayoutYaxisCalendar = "nepali"
	LayoutYaxisCalendar_persian    LayoutYaxisCalendar = "persian"
	LayoutYaxisCalendar_jalali     LayoutYaxisCalendar = "jalali"
	LayoutYaxisCalendar_taiwan     LayoutYaxisCalendar = "taiwan"
	LayoutYaxisCalendar_thai       LayoutYaxisCalendar = "thai"
	LayoutYaxisCalendar_ummalqura  LayoutYaxisCalendar = "ummalqura"
)

// LayoutYaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutYaxisCategoryorder string

const (
	LayoutYaxisCategoryorder_trace              LayoutYaxisCategoryorder = "trace"
	LayoutYaxisCategoryorder_categoryascending  LayoutYaxisCategoryorder = "category ascending"
	LayoutYaxisCategoryorder_categorydescending LayoutYaxisCategoryorder = "category descending"
	LayoutYaxisCategoryorder_array              LayoutYaxisCategoryorder = "array"
	LayoutYaxisCategoryorder_totalascending     LayoutYaxisCategoryorder = "total ascending"
	LayoutYaxisCategoryorder_totaldescending    LayoutYaxisCategoryorder = "total descending"
	LayoutYaxisCategoryorder_minascending       LayoutYaxisCategoryorder = "min ascending"
	LayoutYaxisCategoryorder_mindescending      LayoutYaxisCategoryorder = "min descending"
	LayoutYaxisCategoryorder_maxascending       LayoutYaxisCategoryorder = "max ascending"
	LayoutYaxisCategoryorder_maxdescending      LayoutYaxisCategoryorder = "max descending"
	LayoutYaxisCategoryorder_sumascending       LayoutYaxisCategoryorder = "sum ascending"
	LayoutYaxisCategoryorder_sumdescending      LayoutYaxisCategoryorder = "sum descending"
	LayoutYaxisCategoryorder_meanascending      LayoutYaxisCategoryorder = "mean ascending"
	LayoutYaxisCategoryorder_meandescending     LayoutYaxisCategoryorder = "mean descending"
	LayoutYaxisCategoryorder_medianascending    LayoutYaxisCategoryorder = "median ascending"
	LayoutYaxisCategoryorder_mediandescending   LayoutYaxisCategoryorder = "median descending"
)

// LayoutYaxisConstrain If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines how that happens: by increasing the *range* (default), or by decreasing the *domain*.
type LayoutYaxisConstrain string

const (
	LayoutYaxisConstrain_range  LayoutYaxisConstrain = "range"
	LayoutYaxisConstrain_domain LayoutYaxisConstrain = "domain"
)

// LayoutYaxisConstraintoward If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines which direction we push the originally specified plot area. Options are *left*, *center* (default), and *right* for x axes, and *top*, *middle* (default), and *bottom* for y axes.
type LayoutYaxisConstraintoward string

const (
	LayoutYaxisConstraintoward_left   LayoutYaxisConstraintoward = "left"
	LayoutYaxisConstraintoward_center LayoutYaxisConstraintoward = "center"
	LayoutYaxisConstraintoward_right  LayoutYaxisConstraintoward = "right"
	LayoutYaxisConstraintoward_top    LayoutYaxisConstraintoward = "top"
	LayoutYaxisConstraintoward_middle LayoutYaxisConstraintoward = "middle"
	LayoutYaxisConstraintoward_bottom LayoutYaxisConstraintoward = "bottom"
)

// LayoutYaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutYaxisExponentformat string

const (
	LayoutYaxisExponentformat_none  LayoutYaxisExponentformat = "none"
	LayoutYaxisExponentformat_e     LayoutYaxisExponentformat = "e"
	LayoutYaxisExponentformat_E     LayoutYaxisExponentformat = "E"
	LayoutYaxisExponentformat_power LayoutYaxisExponentformat = "power"
	LayoutYaxisExponentformat_SI    LayoutYaxisExponentformat = "SI"
	LayoutYaxisExponentformat_B     LayoutYaxisExponentformat = "B"
)

// LayoutYaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutYaxisLayer string

const (
	LayoutYaxisLayer_abovetraces LayoutYaxisLayer = "above traces"
	LayoutYaxisLayer_belowtraces LayoutYaxisLayer = "below traces"
)

// LayoutYaxisMatches If set to another axis id (e.g. `x2`, `y`), the range of this axis will match the range of the corresponding axis in data-coordinates space. Moreover, matching axes share auto-range values, category lists and histogram auto-bins. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden. Moreover, note that matching axes must have the same `type`.
type LayoutYaxisMatches string

const (
	LayoutYaxisMatches_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisMatches = "/^x([2-9]|[1-9][0-9]+)?$/"
	LayoutYaxisMatches_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisMatches = "/^y([2-9]|[1-9][0-9]+)?$/"
)

// LayoutYaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutYaxisMirror interface{}

var (
	LayoutYaxisMirror_True     LayoutYaxisMirror = true
	LayoutYaxisMirror_ticks    LayoutYaxisMirror = "ticks"
	LayoutYaxisMirror_False    LayoutYaxisMirror = false
	LayoutYaxisMirror_all      LayoutYaxisMirror = "all"
	LayoutYaxisMirror_allticks LayoutYaxisMirror = "allticks"
)

// LayoutYaxisOverlaying If set a same-letter axis id, this axis is overlaid on top of the corresponding same-letter axis, with traces and axes visible for both axes. If *false*, this axis does not overlay any same-letter axes. In this case, for axes with overlapping domains only the highest-numbered axis will be visible.
type LayoutYaxisOverlaying string

const (
	LayoutYaxisOverlaying_free                                                                                                                       LayoutYaxisOverlaying = "free"
	LayoutYaxisOverlaying_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisOverlaying = "/^x([2-9]|[1-9][0-9]+)?$/"
	LayoutYaxisOverlaying_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisOverlaying = "/^y([2-9]|[1-9][0-9]+)?$/"
)

// LayoutYaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutYaxisRangemode string

const (
	LayoutYaxisRangemode_normal      LayoutYaxisRangemode = "normal"
	LayoutYaxisRangemode_tozero      LayoutYaxisRangemode = "tozero"
	LayoutYaxisRangemode_nonnegative LayoutYaxisRangemode = "nonnegative"
)

// LayoutYaxisScaleanchor If set to another axis id (e.g. `x2`, `y`), the range of this axis changes together with the range of the corresponding axis such that the scale of pixels per unit is in a constant ratio. Both axes are still zoomable, but when you zoom one, the other will zoom the same amount, keeping a fixed midpoint. `constrain` and `constraintoward` determine how we enforce the constraint. You can chain these, ie `yaxis: {scaleanchor: *x*}, xaxis2: {scaleanchor: *y*}` but you can only link axes of the same `type`. The linked axis can have the opposite letter (to constrain the aspect ratio) or the same letter (to match scales across subplots). Loops (`yaxis: {scaleanchor: *x*}, xaxis: {scaleanchor: *y*}` or longer) are redundant and the last constraint encountered will be ignored to avoid possible inconsistent constraints via `scaleratio`. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden.
type LayoutYaxisScaleanchor string

const (
	LayoutYaxisScaleanchor_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisScaleanchor = "/^x([2-9]|[1-9][0-9]+)?$/"
	LayoutYaxisScaleanchor_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisScaleanchor = "/^y([2-9]|[1-9][0-9]+)?$/"
)

// LayoutYaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutYaxisShowexponent string

const (
	LayoutYaxisShowexponent_all   LayoutYaxisShowexponent = "all"
	LayoutYaxisShowexponent_first LayoutYaxisShowexponent = "first"
	LayoutYaxisShowexponent_last  LayoutYaxisShowexponent = "last"
	LayoutYaxisShowexponent_none  LayoutYaxisShowexponent = "none"
)

// LayoutYaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutYaxisShowtickprefix string

const (
	LayoutYaxisShowtickprefix_all   LayoutYaxisShowtickprefix = "all"
	LayoutYaxisShowtickprefix_first LayoutYaxisShowtickprefix = "first"
	LayoutYaxisShowtickprefix_last  LayoutYaxisShowtickprefix = "last"
	LayoutYaxisShowtickprefix_none  LayoutYaxisShowtickprefix = "none"
)

// LayoutYaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutYaxisShowticksuffix string

const (
	LayoutYaxisShowticksuffix_all   LayoutYaxisShowticksuffix = "all"
	LayoutYaxisShowticksuffix_first LayoutYaxisShowticksuffix = "first"
	LayoutYaxisShowticksuffix_last  LayoutYaxisShowticksuffix = "last"
	LayoutYaxisShowticksuffix_none  LayoutYaxisShowticksuffix = "none"
)

// LayoutYaxisSide Determines whether a x (y) axis is positioned at the *bottom* (*left*) or *top* (*right*) of the plotting area.
type LayoutYaxisSide string

const (
	LayoutYaxisSide_top    LayoutYaxisSide = "top"
	LayoutYaxisSide_bottom LayoutYaxisSide = "bottom"
	LayoutYaxisSide_left   LayoutYaxisSide = "left"
	LayoutYaxisSide_right  LayoutYaxisSide = "right"
)

// LayoutYaxisSpikesnap Determines whether spikelines are stuck to the cursor or to the closest datapoints.
type LayoutYaxisSpikesnap string

const (
	LayoutYaxisSpikesnap_data        LayoutYaxisSpikesnap = "data"
	LayoutYaxisSpikesnap_cursor      LayoutYaxisSpikesnap = "cursor"
	LayoutYaxisSpikesnap_hovereddata LayoutYaxisSpikesnap = "hovered data"
)

// LayoutYaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutYaxisTickmode string

const (
	LayoutYaxisTickmode_auto   LayoutYaxisTickmode = "auto"
	LayoutYaxisTickmode_linear LayoutYaxisTickmode = "linear"
	LayoutYaxisTickmode_array  LayoutYaxisTickmode = "array"
)

// LayoutYaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutYaxisTicks string

const (
	LayoutYaxisTicks_outside LayoutYaxisTicks = "outside"
	LayoutYaxisTicks_inside  LayoutYaxisTicks = "inside"
)

// LayoutYaxisTickson Determines where ticks and grid lines are drawn with respect to their corresponding tick labels. Only has an effect for axes of `type` *category* or *multicategory*. When set to *boundaries*, ticks and grid lines are drawn half a category to the left/bottom of labels.
type LayoutYaxisTickson string

const (
	LayoutYaxisTickson_labels     LayoutYaxisTickson = "labels"
	LayoutYaxisTickson_boundaries LayoutYaxisTickson = "boundaries"
)

// LayoutYaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutYaxisType string

const (
	LayoutYaxisType__             LayoutYaxisType = "-"
	LayoutYaxisType_linear        LayoutYaxisType = "linear"
	LayoutYaxisType_log           LayoutYaxisType = "log"
	LayoutYaxisType_date          LayoutYaxisType = "date"
	LayoutYaxisType_category      LayoutYaxisType = "category"
	LayoutYaxisType_multicategory LayoutYaxisType = "multicategory"
)

// Mesh3dColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type Mesh3dColorbarExponentformat string

const (
	Mesh3dColorbarExponentformat_none  Mesh3dColorbarExponentformat = "none"
	Mesh3dColorbarExponentformat_e     Mesh3dColorbarExponentformat = "e"
	Mesh3dColorbarExponentformat_E     Mesh3dColorbarExponentformat = "E"
	Mesh3dColorbarExponentformat_power Mesh3dColorbarExponentformat = "power"
	Mesh3dColorbarExponentformat_SI    Mesh3dColorbarExponentformat = "SI"
	Mesh3dColorbarExponentformat_B     Mesh3dColorbarExponentformat = "B"
)

// Mesh3dColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type Mesh3dColorbarLenmode string

const (
	Mesh3dColorbarLenmode_fraction Mesh3dColorbarLenmode = "fraction"
	Mesh3dColorbarLenmode_pixels   Mesh3dColorbarLenmode = "pixels"
)

// Mesh3dColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type Mesh3dColorbarShowexponent string

const (
	Mesh3dColorbarShowexponent_all   Mesh3dColorbarShowexponent = "all"
	Mesh3dColorbarShowexponent_first Mesh3dColorbarShowexponent = "first"
	Mesh3dColorbarShowexponent_last  Mesh3dColorbarShowexponent = "last"
	Mesh3dColorbarShowexponent_none  Mesh3dColorbarShowexponent = "none"
)

// Mesh3dColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type Mesh3dColorbarShowtickprefix string

const (
	Mesh3dColorbarShowtickprefix_all   Mesh3dColorbarShowtickprefix = "all"
	Mesh3dColorbarShowtickprefix_first Mesh3dColorbarShowtickprefix = "first"
	Mesh3dColorbarShowtickprefix_last  Mesh3dColorbarShowtickprefix = "last"
	Mesh3dColorbarShowtickprefix_none  Mesh3dColorbarShowtickprefix = "none"
)

// Mesh3dColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type Mesh3dColorbarShowticksuffix string

const (
	Mesh3dColorbarShowticksuffix_all   Mesh3dColorbarShowticksuffix = "all"
	Mesh3dColorbarShowticksuffix_first Mesh3dColorbarShowticksuffix = "first"
	Mesh3dColorbarShowticksuffix_last  Mesh3dColorbarShowticksuffix = "last"
	Mesh3dColorbarShowticksuffix_none  Mesh3dColorbarShowticksuffix = "none"
)

// Mesh3dColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type Mesh3dColorbarThicknessmode string

const (
	Mesh3dColorbarThicknessmode_fraction Mesh3dColorbarThicknessmode = "fraction"
	Mesh3dColorbarThicknessmode_pixels   Mesh3dColorbarThicknessmode = "pixels"
)

// Mesh3dColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type Mesh3dColorbarTickmode string

const (
	Mesh3dColorbarTickmode_auto   Mesh3dColorbarTickmode = "auto"
	Mesh3dColorbarTickmode_linear Mesh3dColorbarTickmode = "linear"
	Mesh3dColorbarTickmode_array  Mesh3dColorbarTickmode = "array"
)

// Mesh3dColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type Mesh3dColorbarTicks string

const (
	Mesh3dColorbarTicks_outside Mesh3dColorbarTicks = "outside"
	Mesh3dColorbarTicks_inside  Mesh3dColorbarTicks = "inside"
)

// Mesh3dColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type Mesh3dColorbarTitleSide string

const (
	Mesh3dColorbarTitleSide_right  Mesh3dColorbarTitleSide = "right"
	Mesh3dColorbarTitleSide_top    Mesh3dColorbarTitleSide = "top"
	Mesh3dColorbarTitleSide_bottom Mesh3dColorbarTitleSide = "bottom"
)

// Mesh3dColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type Mesh3dColorbarXanchor string

const (
	Mesh3dColorbarXanchor_left   Mesh3dColorbarXanchor = "left"
	Mesh3dColorbarXanchor_center Mesh3dColorbarXanchor = "center"
	Mesh3dColorbarXanchor_right  Mesh3dColorbarXanchor = "right"
)

// Mesh3dColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type Mesh3dColorbarYanchor string

const (
	Mesh3dColorbarYanchor_top    Mesh3dColorbarYanchor = "top"
	Mesh3dColorbarYanchor_middle Mesh3dColorbarYanchor = "middle"
	Mesh3dColorbarYanchor_bottom Mesh3dColorbarYanchor = "bottom"
)

// Mesh3dDelaunayaxis Sets the Delaunay axis, which is the axis that is perpendicular to the surface of the Delaunay triangulation. It has an effect if `i`, `j`, `k` are not provided and `alphahull` is set to indicate Delaunay triangulation.
type Mesh3dDelaunayaxis string

const (
	Mesh3dDelaunayaxis_x Mesh3dDelaunayaxis = "x"
	Mesh3dDelaunayaxis_y Mesh3dDelaunayaxis = "y"
	Mesh3dDelaunayaxis_z Mesh3dDelaunayaxis = "z"
)

// Mesh3dHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type Mesh3dHoverlabelAlign string

const (
	Mesh3dHoverlabelAlign_left  Mesh3dHoverlabelAlign = "left"
	Mesh3dHoverlabelAlign_right Mesh3dHoverlabelAlign = "right"
	Mesh3dHoverlabelAlign_auto  Mesh3dHoverlabelAlign = "auto"
)

// Mesh3dIntensitymode Determines the source of `intensity` values.
type Mesh3dIntensitymode string

const (
	Mesh3dIntensitymode_vertex Mesh3dIntensitymode = "vertex"
	Mesh3dIntensitymode_cell   Mesh3dIntensitymode = "cell"
)

// Mesh3dVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type Mesh3dVisible interface{}

var (
	Mesh3dVisible_True       Mesh3dVisible = true
	Mesh3dVisible_False      Mesh3dVisible = false
	Mesh3dVisible_legendonly Mesh3dVisible = "legendonly"
)

// Mesh3dXcalendar Sets the calendar system to use with `x` date data.
type Mesh3dXcalendar string

const (
	Mesh3dXcalendar_gregorian  Mesh3dXcalendar = "gregorian"
	Mesh3dXcalendar_chinese    Mesh3dXcalendar = "chinese"
	Mesh3dXcalendar_coptic     Mesh3dXcalendar = "coptic"
	Mesh3dXcalendar_discworld  Mesh3dXcalendar = "discworld"
	Mesh3dXcalendar_ethiopian  Mesh3dXcalendar = "ethiopian"
	Mesh3dXcalendar_hebrew     Mesh3dXcalendar = "hebrew"
	Mesh3dXcalendar_islamic    Mesh3dXcalendar = "islamic"
	Mesh3dXcalendar_julian     Mesh3dXcalendar = "julian"
	Mesh3dXcalendar_mayan      Mesh3dXcalendar = "mayan"
	Mesh3dXcalendar_nanakshahi Mesh3dXcalendar = "nanakshahi"
	Mesh3dXcalendar_nepali     Mesh3dXcalendar = "nepali"
	Mesh3dXcalendar_persian    Mesh3dXcalendar = "persian"
	Mesh3dXcalendar_jalali     Mesh3dXcalendar = "jalali"
	Mesh3dXcalendar_taiwan     Mesh3dXcalendar = "taiwan"
	Mesh3dXcalendar_thai       Mesh3dXcalendar = "thai"
	Mesh3dXcalendar_ummalqura  Mesh3dXcalendar = "ummalqura"
)

// Mesh3dYcalendar Sets the calendar system to use with `y` date data.
type Mesh3dYcalendar string

const (
	Mesh3dYcalendar_gregorian  Mesh3dYcalendar = "gregorian"
	Mesh3dYcalendar_chinese    Mesh3dYcalendar = "chinese"
	Mesh3dYcalendar_coptic     Mesh3dYcalendar = "coptic"
	Mesh3dYcalendar_discworld  Mesh3dYcalendar = "discworld"
	Mesh3dYcalendar_ethiopian  Mesh3dYcalendar = "ethiopian"
	Mesh3dYcalendar_hebrew     Mesh3dYcalendar = "hebrew"
	Mesh3dYcalendar_islamic    Mesh3dYcalendar = "islamic"
	Mesh3dYcalendar_julian     Mesh3dYcalendar = "julian"
	Mesh3dYcalendar_mayan      Mesh3dYcalendar = "mayan"
	Mesh3dYcalendar_nanakshahi Mesh3dYcalendar = "nanakshahi"
	Mesh3dYcalendar_nepali     Mesh3dYcalendar = "nepali"
	Mesh3dYcalendar_persian    Mesh3dYcalendar = "persian"
	Mesh3dYcalendar_jalali     Mesh3dYcalendar = "jalali"
	Mesh3dYcalendar_taiwan     Mesh3dYcalendar = "taiwan"
	Mesh3dYcalendar_thai       Mesh3dYcalendar = "thai"
	Mesh3dYcalendar_ummalqura  Mesh3dYcalendar = "ummalqura"
)

// Mesh3dZcalendar Sets the calendar system to use with `z` date data.
type Mesh3dZcalendar string

const (
	Mesh3dZcalendar_gregorian  Mesh3dZcalendar = "gregorian"
	Mesh3dZcalendar_chinese    Mesh3dZcalendar = "chinese"
	Mesh3dZcalendar_coptic     Mesh3dZcalendar = "coptic"
	Mesh3dZcalendar_discworld  Mesh3dZcalendar = "discworld"
	Mesh3dZcalendar_ethiopian  Mesh3dZcalendar = "ethiopian"
	Mesh3dZcalendar_hebrew     Mesh3dZcalendar = "hebrew"
	Mesh3dZcalendar_islamic    Mesh3dZcalendar = "islamic"
	Mesh3dZcalendar_julian     Mesh3dZcalendar = "julian"
	Mesh3dZcalendar_mayan      Mesh3dZcalendar = "mayan"
	Mesh3dZcalendar_nanakshahi Mesh3dZcalendar = "nanakshahi"
	Mesh3dZcalendar_nepali     Mesh3dZcalendar = "nepali"
	Mesh3dZcalendar_persian    Mesh3dZcalendar = "persian"
	Mesh3dZcalendar_jalali     Mesh3dZcalendar = "jalali"
	Mesh3dZcalendar_taiwan     Mesh3dZcalendar = "taiwan"
	Mesh3dZcalendar_thai       Mesh3dZcalendar = "thai"
	Mesh3dZcalendar_ummalqura  Mesh3dZcalendar = "ummalqura"
)

// OhlcHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type OhlcHoverlabelAlign string

const (
	OhlcHoverlabelAlign_left  OhlcHoverlabelAlign = "left"
	OhlcHoverlabelAlign_right OhlcHoverlabelAlign = "right"
	OhlcHoverlabelAlign_auto  OhlcHoverlabelAlign = "auto"
)

// OhlcVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type OhlcVisible interface{}

var (
	OhlcVisible_True       OhlcVisible = true
	OhlcVisible_False      OhlcVisible = false
	OhlcVisible_legendonly OhlcVisible = "legendonly"
)

// OhlcXcalendar Sets the calendar system to use with `x` date data.
type OhlcXcalendar string

const (
	OhlcXcalendar_gregorian  OhlcXcalendar = "gregorian"
	OhlcXcalendar_chinese    OhlcXcalendar = "chinese"
	OhlcXcalendar_coptic     OhlcXcalendar = "coptic"
	OhlcXcalendar_discworld  OhlcXcalendar = "discworld"
	OhlcXcalendar_ethiopian  OhlcXcalendar = "ethiopian"
	OhlcXcalendar_hebrew     OhlcXcalendar = "hebrew"
	OhlcXcalendar_islamic    OhlcXcalendar = "islamic"
	OhlcXcalendar_julian     OhlcXcalendar = "julian"
	OhlcXcalendar_mayan      OhlcXcalendar = "mayan"
	OhlcXcalendar_nanakshahi OhlcXcalendar = "nanakshahi"
	OhlcXcalendar_nepali     OhlcXcalendar = "nepali"
	OhlcXcalendar_persian    OhlcXcalendar = "persian"
	OhlcXcalendar_jalali     OhlcXcalendar = "jalali"
	OhlcXcalendar_taiwan     OhlcXcalendar = "taiwan"
	OhlcXcalendar_thai       OhlcXcalendar = "thai"
	OhlcXcalendar_ummalqura  OhlcXcalendar = "ummalqura"
)

// ParcatsArrangement Sets the drag interaction mode for categories and dimensions. If `perpendicular`, the categories can only move along a line perpendicular to the paths. If `freeform`, the categories can freely move on the plane. If `fixed`, the categories and dimensions are stationary.
type ParcatsArrangement string

const (
	ParcatsArrangement_perpendicular ParcatsArrangement = "perpendicular"
	ParcatsArrangement_freeform      ParcatsArrangement = "freeform"
	ParcatsArrangement_fixed         ParcatsArrangement = "fixed"
)

// ParcatsHoveron Sets the hover interaction mode for the parcats diagram. If `category`, hover interaction take place per category. If `color`, hover interactions take place per color per category. If `dimension`, hover interactions take place across all categories per dimension.
type ParcatsHoveron string

const (
	ParcatsHoveron_category  ParcatsHoveron = "category"
	ParcatsHoveron_color     ParcatsHoveron = "color"
	ParcatsHoveron_dimension ParcatsHoveron = "dimension"
)

// ParcatsLineColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ParcatsLineColorbarExponentformat string

const (
	ParcatsLineColorbarExponentformat_none  ParcatsLineColorbarExponentformat = "none"
	ParcatsLineColorbarExponentformat_e     ParcatsLineColorbarExponentformat = "e"
	ParcatsLineColorbarExponentformat_E     ParcatsLineColorbarExponentformat = "E"
	ParcatsLineColorbarExponentformat_power ParcatsLineColorbarExponentformat = "power"
	ParcatsLineColorbarExponentformat_SI    ParcatsLineColorbarExponentformat = "SI"
	ParcatsLineColorbarExponentformat_B     ParcatsLineColorbarExponentformat = "B"
)

// ParcatsLineColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ParcatsLineColorbarLenmode string

const (
	ParcatsLineColorbarLenmode_fraction ParcatsLineColorbarLenmode = "fraction"
	ParcatsLineColorbarLenmode_pixels   ParcatsLineColorbarLenmode = "pixels"
)

// ParcatsLineColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ParcatsLineColorbarShowexponent string

const (
	ParcatsLineColorbarShowexponent_all   ParcatsLineColorbarShowexponent = "all"
	ParcatsLineColorbarShowexponent_first ParcatsLineColorbarShowexponent = "first"
	ParcatsLineColorbarShowexponent_last  ParcatsLineColorbarShowexponent = "last"
	ParcatsLineColorbarShowexponent_none  ParcatsLineColorbarShowexponent = "none"
)

// ParcatsLineColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ParcatsLineColorbarShowtickprefix string

const (
	ParcatsLineColorbarShowtickprefix_all   ParcatsLineColorbarShowtickprefix = "all"
	ParcatsLineColorbarShowtickprefix_first ParcatsLineColorbarShowtickprefix = "first"
	ParcatsLineColorbarShowtickprefix_last  ParcatsLineColorbarShowtickprefix = "last"
	ParcatsLineColorbarShowtickprefix_none  ParcatsLineColorbarShowtickprefix = "none"
)

// ParcatsLineColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ParcatsLineColorbarShowticksuffix string

const (
	ParcatsLineColorbarShowticksuffix_all   ParcatsLineColorbarShowticksuffix = "all"
	ParcatsLineColorbarShowticksuffix_first ParcatsLineColorbarShowticksuffix = "first"
	ParcatsLineColorbarShowticksuffix_last  ParcatsLineColorbarShowticksuffix = "last"
	ParcatsLineColorbarShowticksuffix_none  ParcatsLineColorbarShowticksuffix = "none"
)

// ParcatsLineColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ParcatsLineColorbarThicknessmode string

const (
	ParcatsLineColorbarThicknessmode_fraction ParcatsLineColorbarThicknessmode = "fraction"
	ParcatsLineColorbarThicknessmode_pixels   ParcatsLineColorbarThicknessmode = "pixels"
)

// ParcatsLineColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ParcatsLineColorbarTickmode string

const (
	ParcatsLineColorbarTickmode_auto   ParcatsLineColorbarTickmode = "auto"
	ParcatsLineColorbarTickmode_linear ParcatsLineColorbarTickmode = "linear"
	ParcatsLineColorbarTickmode_array  ParcatsLineColorbarTickmode = "array"
)

// ParcatsLineColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ParcatsLineColorbarTicks string

const (
	ParcatsLineColorbarTicks_outside ParcatsLineColorbarTicks = "outside"
	ParcatsLineColorbarTicks_inside  ParcatsLineColorbarTicks = "inside"
)

// ParcatsLineColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ParcatsLineColorbarTitleSide string

const (
	ParcatsLineColorbarTitleSide_right  ParcatsLineColorbarTitleSide = "right"
	ParcatsLineColorbarTitleSide_top    ParcatsLineColorbarTitleSide = "top"
	ParcatsLineColorbarTitleSide_bottom ParcatsLineColorbarTitleSide = "bottom"
)

// ParcatsLineColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ParcatsLineColorbarXanchor string

const (
	ParcatsLineColorbarXanchor_left   ParcatsLineColorbarXanchor = "left"
	ParcatsLineColorbarXanchor_center ParcatsLineColorbarXanchor = "center"
	ParcatsLineColorbarXanchor_right  ParcatsLineColorbarXanchor = "right"
)

// ParcatsLineColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ParcatsLineColorbarYanchor string

const (
	ParcatsLineColorbarYanchor_top    ParcatsLineColorbarYanchor = "top"
	ParcatsLineColorbarYanchor_middle ParcatsLineColorbarYanchor = "middle"
	ParcatsLineColorbarYanchor_bottom ParcatsLineColorbarYanchor = "bottom"
)

// ParcatsLineShape Sets the shape of the paths. If `linear`, paths are composed of straight lines. If `hspline`, paths are composed of horizontal curved splines
type ParcatsLineShape string

const (
	ParcatsLineShape_linear  ParcatsLineShape = "linear"
	ParcatsLineShape_hspline ParcatsLineShape = "hspline"
)

// ParcatsSortpaths Sets the path sorting algorithm. If `forward`, sort paths based on dimension categories from left to right. If `backward`, sort paths based on dimensions categories from right to left.
type ParcatsSortpaths string

const (
	ParcatsSortpaths_forward  ParcatsSortpaths = "forward"
	ParcatsSortpaths_backward ParcatsSortpaths = "backward"
)

// ParcatsVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ParcatsVisible interface{}

var (
	ParcatsVisible_True       ParcatsVisible = true
	ParcatsVisible_False      ParcatsVisible = false
	ParcatsVisible_legendonly ParcatsVisible = "legendonly"
)

// ParcoordsLabelside Specifies the location of the `label`. *top* positions labels above, next to the title *bottom* positions labels below the graph Tilted labels with *labelangle* may be positioned better inside margins when `labelposition` is set to *bottom*.
type ParcoordsLabelside string

const (
	ParcoordsLabelside_top    ParcoordsLabelside = "top"
	ParcoordsLabelside_bottom ParcoordsLabelside = "bottom"
)

// ParcoordsLineColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ParcoordsLineColorbarExponentformat string

const (
	ParcoordsLineColorbarExponentformat_none  ParcoordsLineColorbarExponentformat = "none"
	ParcoordsLineColorbarExponentformat_e     ParcoordsLineColorbarExponentformat = "e"
	ParcoordsLineColorbarExponentformat_E     ParcoordsLineColorbarExponentformat = "E"
	ParcoordsLineColorbarExponentformat_power ParcoordsLineColorbarExponentformat = "power"
	ParcoordsLineColorbarExponentformat_SI    ParcoordsLineColorbarExponentformat = "SI"
	ParcoordsLineColorbarExponentformat_B     ParcoordsLineColorbarExponentformat = "B"
)

// ParcoordsLineColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ParcoordsLineColorbarLenmode string

const (
	ParcoordsLineColorbarLenmode_fraction ParcoordsLineColorbarLenmode = "fraction"
	ParcoordsLineColorbarLenmode_pixels   ParcoordsLineColorbarLenmode = "pixels"
)

// ParcoordsLineColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ParcoordsLineColorbarShowexponent string

const (
	ParcoordsLineColorbarShowexponent_all   ParcoordsLineColorbarShowexponent = "all"
	ParcoordsLineColorbarShowexponent_first ParcoordsLineColorbarShowexponent = "first"
	ParcoordsLineColorbarShowexponent_last  ParcoordsLineColorbarShowexponent = "last"
	ParcoordsLineColorbarShowexponent_none  ParcoordsLineColorbarShowexponent = "none"
)

// ParcoordsLineColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ParcoordsLineColorbarShowtickprefix string

const (
	ParcoordsLineColorbarShowtickprefix_all   ParcoordsLineColorbarShowtickprefix = "all"
	ParcoordsLineColorbarShowtickprefix_first ParcoordsLineColorbarShowtickprefix = "first"
	ParcoordsLineColorbarShowtickprefix_last  ParcoordsLineColorbarShowtickprefix = "last"
	ParcoordsLineColorbarShowtickprefix_none  ParcoordsLineColorbarShowtickprefix = "none"
)

// ParcoordsLineColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ParcoordsLineColorbarShowticksuffix string

const (
	ParcoordsLineColorbarShowticksuffix_all   ParcoordsLineColorbarShowticksuffix = "all"
	ParcoordsLineColorbarShowticksuffix_first ParcoordsLineColorbarShowticksuffix = "first"
	ParcoordsLineColorbarShowticksuffix_last  ParcoordsLineColorbarShowticksuffix = "last"
	ParcoordsLineColorbarShowticksuffix_none  ParcoordsLineColorbarShowticksuffix = "none"
)

// ParcoordsLineColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ParcoordsLineColorbarThicknessmode string

const (
	ParcoordsLineColorbarThicknessmode_fraction ParcoordsLineColorbarThicknessmode = "fraction"
	ParcoordsLineColorbarThicknessmode_pixels   ParcoordsLineColorbarThicknessmode = "pixels"
)

// ParcoordsLineColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ParcoordsLineColorbarTickmode string

const (
	ParcoordsLineColorbarTickmode_auto   ParcoordsLineColorbarTickmode = "auto"
	ParcoordsLineColorbarTickmode_linear ParcoordsLineColorbarTickmode = "linear"
	ParcoordsLineColorbarTickmode_array  ParcoordsLineColorbarTickmode = "array"
)

// ParcoordsLineColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ParcoordsLineColorbarTicks string

const (
	ParcoordsLineColorbarTicks_outside ParcoordsLineColorbarTicks = "outside"
	ParcoordsLineColorbarTicks_inside  ParcoordsLineColorbarTicks = "inside"
)

// ParcoordsLineColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ParcoordsLineColorbarTitleSide string

const (
	ParcoordsLineColorbarTitleSide_right  ParcoordsLineColorbarTitleSide = "right"
	ParcoordsLineColorbarTitleSide_top    ParcoordsLineColorbarTitleSide = "top"
	ParcoordsLineColorbarTitleSide_bottom ParcoordsLineColorbarTitleSide = "bottom"
)

// ParcoordsLineColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ParcoordsLineColorbarXanchor string

const (
	ParcoordsLineColorbarXanchor_left   ParcoordsLineColorbarXanchor = "left"
	ParcoordsLineColorbarXanchor_center ParcoordsLineColorbarXanchor = "center"
	ParcoordsLineColorbarXanchor_right  ParcoordsLineColorbarXanchor = "right"
)

// ParcoordsLineColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ParcoordsLineColorbarYanchor string

const (
	ParcoordsLineColorbarYanchor_top    ParcoordsLineColorbarYanchor = "top"
	ParcoordsLineColorbarYanchor_middle ParcoordsLineColorbarYanchor = "middle"
	ParcoordsLineColorbarYanchor_bottom ParcoordsLineColorbarYanchor = "bottom"
)

// ParcoordsVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ParcoordsVisible interface{}

var (
	ParcoordsVisible_True       ParcoordsVisible = true
	ParcoordsVisible_False      ParcoordsVisible = false
	ParcoordsVisible_legendonly ParcoordsVisible = "legendonly"
)

// PieDirection Specifies the direction at which succeeding sectors follow one another.
type PieDirection string

const (
	PieDirection_clockwise        PieDirection = "clockwise"
	PieDirection_counterclockwise PieDirection = "counterclockwise"
)

// PieHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type PieHoverlabelAlign string

const (
	PieHoverlabelAlign_left  PieHoverlabelAlign = "left"
	PieHoverlabelAlign_right PieHoverlabelAlign = "right"
	PieHoverlabelAlign_auto  PieHoverlabelAlign = "auto"
)

// PieInsidetextorientation Controls the orientation of the text inside chart sectors. When set to *auto*, text may be oriented in any direction in order to be as big as possible in the middle of a sector. The *horizontal* option orients text to be parallel with the bottom of the chart, and may make text smaller in order to achieve that goal. The *radial* option orients text along the radius of the sector. The *tangential* option orients text perpendicular to the radius of the sector.
type PieInsidetextorientation string

const (
	PieInsidetextorientation_horizontal PieInsidetextorientation = "horizontal"
	PieInsidetextorientation_radial     PieInsidetextorientation = "radial"
	PieInsidetextorientation_tangential PieInsidetextorientation = "tangential"
	PieInsidetextorientation_auto       PieInsidetextorientation = "auto"
)

// PieTextposition Specifies the location of the `textinfo`.
type PieTextposition string

const (
	PieTextposition_inside  PieTextposition = "inside"
	PieTextposition_outside PieTextposition = "outside"
	PieTextposition_auto    PieTextposition = "auto"
	PieTextposition_none    PieTextposition = "none"
)

// PieTitlePosition Specifies the location of the `title`. Note that the title's position used to be set by the now deprecated `titleposition` attribute.
type PieTitlePosition string

const (
	PieTitlePosition_topleft      PieTitlePosition = "top left"
	PieTitlePosition_topcenter    PieTitlePosition = "top center"
	PieTitlePosition_topright     PieTitlePosition = "top right"
	PieTitlePosition_middlecenter PieTitlePosition = "middle center"
	PieTitlePosition_bottomleft   PieTitlePosition = "bottom left"
	PieTitlePosition_bottomcenter PieTitlePosition = "bottom center"
	PieTitlePosition_bottomright  PieTitlePosition = "bottom right"
)

// PieVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type PieVisible interface{}

var (
	PieVisible_True       PieVisible = true
	PieVisible_False      PieVisible = false
	PieVisible_legendonly PieVisible = "legendonly"
)

// PointcloudHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type PointcloudHoverlabelAlign string

const (
	PointcloudHoverlabelAlign_left  PointcloudHoverlabelAlign = "left"
	PointcloudHoverlabelAlign_right PointcloudHoverlabelAlign = "right"
	PointcloudHoverlabelAlign_auto  PointcloudHoverlabelAlign = "auto"
)

// PointcloudVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type PointcloudVisible interface{}

var (
	PointcloudVisible_True       PointcloudVisible = true
	PointcloudVisible_False      PointcloudVisible = false
	PointcloudVisible_legendonly PointcloudVisible = "legendonly"
)

// SankeyArrangement If value is `snap` (the default), the node arrangement is assisted by automatic snapping of elements to preserve space between nodes specified via `nodepad`. If value is `perpendicular`, the nodes can only move along a line perpendicular to the flow. If value is `freeform`, the nodes can freely move on the plane. If value is `fixed`, the nodes are stationary.
type SankeyArrangement string

const (
	SankeyArrangement_snap          SankeyArrangement = "snap"
	SankeyArrangement_perpendicular SankeyArrangement = "perpendicular"
	SankeyArrangement_freeform      SankeyArrangement = "freeform"
	SankeyArrangement_fixed         SankeyArrangement = "fixed"
)

// SankeyHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SankeyHoverlabelAlign string

const (
	SankeyHoverlabelAlign_left  SankeyHoverlabelAlign = "left"
	SankeyHoverlabelAlign_right SankeyHoverlabelAlign = "right"
	SankeyHoverlabelAlign_auto  SankeyHoverlabelAlign = "auto"
)

// SankeyLinkHoverinfo Determines which trace information appear when hovering links. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type SankeyLinkHoverinfo string

const (
	SankeyLinkHoverinfo_all  SankeyLinkHoverinfo = "all"
	SankeyLinkHoverinfo_none SankeyLinkHoverinfo = "none"
	SankeyLinkHoverinfo_skip SankeyLinkHoverinfo = "skip"
)

// SankeyLinkHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SankeyLinkHoverlabelAlign string

const (
	SankeyLinkHoverlabelAlign_left  SankeyLinkHoverlabelAlign = "left"
	SankeyLinkHoverlabelAlign_right SankeyLinkHoverlabelAlign = "right"
	SankeyLinkHoverlabelAlign_auto  SankeyLinkHoverlabelAlign = "auto"
)

// SankeyNodeHoverinfo Determines which trace information appear when hovering nodes. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type SankeyNodeHoverinfo string

const (
	SankeyNodeHoverinfo_all  SankeyNodeHoverinfo = "all"
	SankeyNodeHoverinfo_none SankeyNodeHoverinfo = "none"
	SankeyNodeHoverinfo_skip SankeyNodeHoverinfo = "skip"
)

// SankeyNodeHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SankeyNodeHoverlabelAlign string

const (
	SankeyNodeHoverlabelAlign_left  SankeyNodeHoverlabelAlign = "left"
	SankeyNodeHoverlabelAlign_right SankeyNodeHoverlabelAlign = "right"
	SankeyNodeHoverlabelAlign_auto  SankeyNodeHoverlabelAlign = "auto"
)

// SankeyOrientation Sets the orientation of the Sankey diagram.
type SankeyOrientation string

const (
	SankeyOrientation_v SankeyOrientation = "v"
	SankeyOrientation_h SankeyOrientation = "h"
)

// SankeyVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type SankeyVisible interface{}

var (
	SankeyVisible_True       SankeyVisible = true
	SankeyVisible_False      SankeyVisible = false
	SankeyVisible_legendonly SankeyVisible = "legendonly"
)

// Scatter3dErrorXType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type Scatter3dErrorXType string

const (
	Scatter3dErrorXType_percent  Scatter3dErrorXType = "percent"
	Scatter3dErrorXType_constant Scatter3dErrorXType = "constant"
	Scatter3dErrorXType_sqrt     Scatter3dErrorXType = "sqrt"
	Scatter3dErrorXType_data     Scatter3dErrorXType = "data"
)

// Scatter3dErrorYType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type Scatter3dErrorYType string

const (
	Scatter3dErrorYType_percent  Scatter3dErrorYType = "percent"
	Scatter3dErrorYType_constant Scatter3dErrorYType = "constant"
	Scatter3dErrorYType_sqrt     Scatter3dErrorYType = "sqrt"
	Scatter3dErrorYType_data     Scatter3dErrorYType = "data"
)

// Scatter3dErrorZType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type Scatter3dErrorZType string

const (
	Scatter3dErrorZType_percent  Scatter3dErrorZType = "percent"
	Scatter3dErrorZType_constant Scatter3dErrorZType = "constant"
	Scatter3dErrorZType_sqrt     Scatter3dErrorZType = "sqrt"
	Scatter3dErrorZType_data     Scatter3dErrorZType = "data"
)

// Scatter3dHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type Scatter3dHoverlabelAlign string

const (
	Scatter3dHoverlabelAlign_left  Scatter3dHoverlabelAlign = "left"
	Scatter3dHoverlabelAlign_right Scatter3dHoverlabelAlign = "right"
	Scatter3dHoverlabelAlign_auto  Scatter3dHoverlabelAlign = "auto"
)

// Scatter3dLineColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type Scatter3dLineColorbarExponentformat string

const (
	Scatter3dLineColorbarExponentformat_none  Scatter3dLineColorbarExponentformat = "none"
	Scatter3dLineColorbarExponentformat_e     Scatter3dLineColorbarExponentformat = "e"
	Scatter3dLineColorbarExponentformat_E     Scatter3dLineColorbarExponentformat = "E"
	Scatter3dLineColorbarExponentformat_power Scatter3dLineColorbarExponentformat = "power"
	Scatter3dLineColorbarExponentformat_SI    Scatter3dLineColorbarExponentformat = "SI"
	Scatter3dLineColorbarExponentformat_B     Scatter3dLineColorbarExponentformat = "B"
)

// Scatter3dLineColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type Scatter3dLineColorbarLenmode string

const (
	Scatter3dLineColorbarLenmode_fraction Scatter3dLineColorbarLenmode = "fraction"
	Scatter3dLineColorbarLenmode_pixels   Scatter3dLineColorbarLenmode = "pixels"
)

// Scatter3dLineColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type Scatter3dLineColorbarShowexponent string

const (
	Scatter3dLineColorbarShowexponent_all   Scatter3dLineColorbarShowexponent = "all"
	Scatter3dLineColorbarShowexponent_first Scatter3dLineColorbarShowexponent = "first"
	Scatter3dLineColorbarShowexponent_last  Scatter3dLineColorbarShowexponent = "last"
	Scatter3dLineColorbarShowexponent_none  Scatter3dLineColorbarShowexponent = "none"
)

// Scatter3dLineColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type Scatter3dLineColorbarShowtickprefix string

const (
	Scatter3dLineColorbarShowtickprefix_all   Scatter3dLineColorbarShowtickprefix = "all"
	Scatter3dLineColorbarShowtickprefix_first Scatter3dLineColorbarShowtickprefix = "first"
	Scatter3dLineColorbarShowtickprefix_last  Scatter3dLineColorbarShowtickprefix = "last"
	Scatter3dLineColorbarShowtickprefix_none  Scatter3dLineColorbarShowtickprefix = "none"
)

// Scatter3dLineColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type Scatter3dLineColorbarShowticksuffix string

const (
	Scatter3dLineColorbarShowticksuffix_all   Scatter3dLineColorbarShowticksuffix = "all"
	Scatter3dLineColorbarShowticksuffix_first Scatter3dLineColorbarShowticksuffix = "first"
	Scatter3dLineColorbarShowticksuffix_last  Scatter3dLineColorbarShowticksuffix = "last"
	Scatter3dLineColorbarShowticksuffix_none  Scatter3dLineColorbarShowticksuffix = "none"
)

// Scatter3dLineColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type Scatter3dLineColorbarThicknessmode string

const (
	Scatter3dLineColorbarThicknessmode_fraction Scatter3dLineColorbarThicknessmode = "fraction"
	Scatter3dLineColorbarThicknessmode_pixels   Scatter3dLineColorbarThicknessmode = "pixels"
)

// Scatter3dLineColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type Scatter3dLineColorbarTickmode string

const (
	Scatter3dLineColorbarTickmode_auto   Scatter3dLineColorbarTickmode = "auto"
	Scatter3dLineColorbarTickmode_linear Scatter3dLineColorbarTickmode = "linear"
	Scatter3dLineColorbarTickmode_array  Scatter3dLineColorbarTickmode = "array"
)

// Scatter3dLineColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type Scatter3dLineColorbarTicks string

const (
	Scatter3dLineColorbarTicks_outside Scatter3dLineColorbarTicks = "outside"
	Scatter3dLineColorbarTicks_inside  Scatter3dLineColorbarTicks = "inside"
)

// Scatter3dLineColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type Scatter3dLineColorbarTitleSide string

const (
	Scatter3dLineColorbarTitleSide_right  Scatter3dLineColorbarTitleSide = "right"
	Scatter3dLineColorbarTitleSide_top    Scatter3dLineColorbarTitleSide = "top"
	Scatter3dLineColorbarTitleSide_bottom Scatter3dLineColorbarTitleSide = "bottom"
)

// Scatter3dLineColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type Scatter3dLineColorbarXanchor string

const (
	Scatter3dLineColorbarXanchor_left   Scatter3dLineColorbarXanchor = "left"
	Scatter3dLineColorbarXanchor_center Scatter3dLineColorbarXanchor = "center"
	Scatter3dLineColorbarXanchor_right  Scatter3dLineColorbarXanchor = "right"
)

// Scatter3dLineColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type Scatter3dLineColorbarYanchor string

const (
	Scatter3dLineColorbarYanchor_top    Scatter3dLineColorbarYanchor = "top"
	Scatter3dLineColorbarYanchor_middle Scatter3dLineColorbarYanchor = "middle"
	Scatter3dLineColorbarYanchor_bottom Scatter3dLineColorbarYanchor = "bottom"
)

// Scatter3dLineDash Sets the dash style of the lines.
type Scatter3dLineDash string

const (
	Scatter3dLineDash_solid       Scatter3dLineDash = "solid"
	Scatter3dLineDash_dot         Scatter3dLineDash = "dot"
	Scatter3dLineDash_dash        Scatter3dLineDash = "dash"
	Scatter3dLineDash_longdash    Scatter3dLineDash = "longdash"
	Scatter3dLineDash_dashdot     Scatter3dLineDash = "dashdot"
	Scatter3dLineDash_longdashdot Scatter3dLineDash = "longdashdot"
)

// Scatter3dMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type Scatter3dMarkerColorbarExponentformat string

const (
	Scatter3dMarkerColorbarExponentformat_none  Scatter3dMarkerColorbarExponentformat = "none"
	Scatter3dMarkerColorbarExponentformat_e     Scatter3dMarkerColorbarExponentformat = "e"
	Scatter3dMarkerColorbarExponentformat_E     Scatter3dMarkerColorbarExponentformat = "E"
	Scatter3dMarkerColorbarExponentformat_power Scatter3dMarkerColorbarExponentformat = "power"
	Scatter3dMarkerColorbarExponentformat_SI    Scatter3dMarkerColorbarExponentformat = "SI"
	Scatter3dMarkerColorbarExponentformat_B     Scatter3dMarkerColorbarExponentformat = "B"
)

// Scatter3dMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type Scatter3dMarkerColorbarLenmode string

const (
	Scatter3dMarkerColorbarLenmode_fraction Scatter3dMarkerColorbarLenmode = "fraction"
	Scatter3dMarkerColorbarLenmode_pixels   Scatter3dMarkerColorbarLenmode = "pixels"
)

// Scatter3dMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type Scatter3dMarkerColorbarShowexponent string

const (
	Scatter3dMarkerColorbarShowexponent_all   Scatter3dMarkerColorbarShowexponent = "all"
	Scatter3dMarkerColorbarShowexponent_first Scatter3dMarkerColorbarShowexponent = "first"
	Scatter3dMarkerColorbarShowexponent_last  Scatter3dMarkerColorbarShowexponent = "last"
	Scatter3dMarkerColorbarShowexponent_none  Scatter3dMarkerColorbarShowexponent = "none"
)

// Scatter3dMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type Scatter3dMarkerColorbarShowtickprefix string

const (
	Scatter3dMarkerColorbarShowtickprefix_all   Scatter3dMarkerColorbarShowtickprefix = "all"
	Scatter3dMarkerColorbarShowtickprefix_first Scatter3dMarkerColorbarShowtickprefix = "first"
	Scatter3dMarkerColorbarShowtickprefix_last  Scatter3dMarkerColorbarShowtickprefix = "last"
	Scatter3dMarkerColorbarShowtickprefix_none  Scatter3dMarkerColorbarShowtickprefix = "none"
)

// Scatter3dMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type Scatter3dMarkerColorbarShowticksuffix string

const (
	Scatter3dMarkerColorbarShowticksuffix_all   Scatter3dMarkerColorbarShowticksuffix = "all"
	Scatter3dMarkerColorbarShowticksuffix_first Scatter3dMarkerColorbarShowticksuffix = "first"
	Scatter3dMarkerColorbarShowticksuffix_last  Scatter3dMarkerColorbarShowticksuffix = "last"
	Scatter3dMarkerColorbarShowticksuffix_none  Scatter3dMarkerColorbarShowticksuffix = "none"
)

// Scatter3dMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type Scatter3dMarkerColorbarThicknessmode string

const (
	Scatter3dMarkerColorbarThicknessmode_fraction Scatter3dMarkerColorbarThicknessmode = "fraction"
	Scatter3dMarkerColorbarThicknessmode_pixels   Scatter3dMarkerColorbarThicknessmode = "pixels"
)

// Scatter3dMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type Scatter3dMarkerColorbarTickmode string

const (
	Scatter3dMarkerColorbarTickmode_auto   Scatter3dMarkerColorbarTickmode = "auto"
	Scatter3dMarkerColorbarTickmode_linear Scatter3dMarkerColorbarTickmode = "linear"
	Scatter3dMarkerColorbarTickmode_array  Scatter3dMarkerColorbarTickmode = "array"
)

// Scatter3dMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type Scatter3dMarkerColorbarTicks string

const (
	Scatter3dMarkerColorbarTicks_outside Scatter3dMarkerColorbarTicks = "outside"
	Scatter3dMarkerColorbarTicks_inside  Scatter3dMarkerColorbarTicks = "inside"
)

// Scatter3dMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type Scatter3dMarkerColorbarTitleSide string

const (
	Scatter3dMarkerColorbarTitleSide_right  Scatter3dMarkerColorbarTitleSide = "right"
	Scatter3dMarkerColorbarTitleSide_top    Scatter3dMarkerColorbarTitleSide = "top"
	Scatter3dMarkerColorbarTitleSide_bottom Scatter3dMarkerColorbarTitleSide = "bottom"
)

// Scatter3dMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type Scatter3dMarkerColorbarXanchor string

const (
	Scatter3dMarkerColorbarXanchor_left   Scatter3dMarkerColorbarXanchor = "left"
	Scatter3dMarkerColorbarXanchor_center Scatter3dMarkerColorbarXanchor = "center"
	Scatter3dMarkerColorbarXanchor_right  Scatter3dMarkerColorbarXanchor = "right"
)

// Scatter3dMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type Scatter3dMarkerColorbarYanchor string

const (
	Scatter3dMarkerColorbarYanchor_top    Scatter3dMarkerColorbarYanchor = "top"
	Scatter3dMarkerColorbarYanchor_middle Scatter3dMarkerColorbarYanchor = "middle"
	Scatter3dMarkerColorbarYanchor_bottom Scatter3dMarkerColorbarYanchor = "bottom"
)

// Scatter3dMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type Scatter3dMarkerSizemode string

const (
	Scatter3dMarkerSizemode_diameter Scatter3dMarkerSizemode = "diameter"
	Scatter3dMarkerSizemode_area     Scatter3dMarkerSizemode = "area"
)

// Scatter3dMarkerSymbol Sets the marker symbol type.
type Scatter3dMarkerSymbol string

const (
	Scatter3dMarkerSymbol_circle       Scatter3dMarkerSymbol = "circle"
	Scatter3dMarkerSymbol_circle_open  Scatter3dMarkerSymbol = "circle-open"
	Scatter3dMarkerSymbol_square       Scatter3dMarkerSymbol = "square"
	Scatter3dMarkerSymbol_square_open  Scatter3dMarkerSymbol = "square-open"
	Scatter3dMarkerSymbol_diamond      Scatter3dMarkerSymbol = "diamond"
	Scatter3dMarkerSymbol_diamond_open Scatter3dMarkerSymbol = "diamond-open"
	Scatter3dMarkerSymbol_cross        Scatter3dMarkerSymbol = "cross"
	Scatter3dMarkerSymbol_x            Scatter3dMarkerSymbol = "x"
)

// Scatter3dSurfaceaxis If *-1*, the scatter points are not fill with a surface If *0*, *1*, *2*, the scatter points are filled with a Delaunay surface about the x, y, z respectively.
type Scatter3dSurfaceaxis string

const (
	Scatter3dSurfaceaxis_1 Scatter3dSurfaceaxis = "-1"
	Scatter3dSurfaceaxis0  Scatter3dSurfaceaxis = "0"
	Scatter3dSurfaceaxis1  Scatter3dSurfaceaxis = "1"
	Scatter3dSurfaceaxis2  Scatter3dSurfaceaxis = "2"
)

// Scatter3dTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type Scatter3dTextposition string

const (
	Scatter3dTextposition_topleft      Scatter3dTextposition = "top left"
	Scatter3dTextposition_topcenter    Scatter3dTextposition = "top center"
	Scatter3dTextposition_topright     Scatter3dTextposition = "top right"
	Scatter3dTextposition_middleleft   Scatter3dTextposition = "middle left"
	Scatter3dTextposition_middlecenter Scatter3dTextposition = "middle center"
	Scatter3dTextposition_middleright  Scatter3dTextposition = "middle right"
	Scatter3dTextposition_bottomleft   Scatter3dTextposition = "bottom left"
	Scatter3dTextposition_bottomcenter Scatter3dTextposition = "bottom center"
	Scatter3dTextposition_bottomright  Scatter3dTextposition = "bottom right"
)

// Scatter3dVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type Scatter3dVisible interface{}

var (
	Scatter3dVisible_True       Scatter3dVisible = true
	Scatter3dVisible_False      Scatter3dVisible = false
	Scatter3dVisible_legendonly Scatter3dVisible = "legendonly"
)

// Scatter3dXcalendar Sets the calendar system to use with `x` date data.
type Scatter3dXcalendar string

const (
	Scatter3dXcalendar_gregorian  Scatter3dXcalendar = "gregorian"
	Scatter3dXcalendar_chinese    Scatter3dXcalendar = "chinese"
	Scatter3dXcalendar_coptic     Scatter3dXcalendar = "coptic"
	Scatter3dXcalendar_discworld  Scatter3dXcalendar = "discworld"
	Scatter3dXcalendar_ethiopian  Scatter3dXcalendar = "ethiopian"
	Scatter3dXcalendar_hebrew     Scatter3dXcalendar = "hebrew"
	Scatter3dXcalendar_islamic    Scatter3dXcalendar = "islamic"
	Scatter3dXcalendar_julian     Scatter3dXcalendar = "julian"
	Scatter3dXcalendar_mayan      Scatter3dXcalendar = "mayan"
	Scatter3dXcalendar_nanakshahi Scatter3dXcalendar = "nanakshahi"
	Scatter3dXcalendar_nepali     Scatter3dXcalendar = "nepali"
	Scatter3dXcalendar_persian    Scatter3dXcalendar = "persian"
	Scatter3dXcalendar_jalali     Scatter3dXcalendar = "jalali"
	Scatter3dXcalendar_taiwan     Scatter3dXcalendar = "taiwan"
	Scatter3dXcalendar_thai       Scatter3dXcalendar = "thai"
	Scatter3dXcalendar_ummalqura  Scatter3dXcalendar = "ummalqura"
)

// Scatter3dYcalendar Sets the calendar system to use with `y` date data.
type Scatter3dYcalendar string

const (
	Scatter3dYcalendar_gregorian  Scatter3dYcalendar = "gregorian"
	Scatter3dYcalendar_chinese    Scatter3dYcalendar = "chinese"
	Scatter3dYcalendar_coptic     Scatter3dYcalendar = "coptic"
	Scatter3dYcalendar_discworld  Scatter3dYcalendar = "discworld"
	Scatter3dYcalendar_ethiopian  Scatter3dYcalendar = "ethiopian"
	Scatter3dYcalendar_hebrew     Scatter3dYcalendar = "hebrew"
	Scatter3dYcalendar_islamic    Scatter3dYcalendar = "islamic"
	Scatter3dYcalendar_julian     Scatter3dYcalendar = "julian"
	Scatter3dYcalendar_mayan      Scatter3dYcalendar = "mayan"
	Scatter3dYcalendar_nanakshahi Scatter3dYcalendar = "nanakshahi"
	Scatter3dYcalendar_nepali     Scatter3dYcalendar = "nepali"
	Scatter3dYcalendar_persian    Scatter3dYcalendar = "persian"
	Scatter3dYcalendar_jalali     Scatter3dYcalendar = "jalali"
	Scatter3dYcalendar_taiwan     Scatter3dYcalendar = "taiwan"
	Scatter3dYcalendar_thai       Scatter3dYcalendar = "thai"
	Scatter3dYcalendar_ummalqura  Scatter3dYcalendar = "ummalqura"
)

// Scatter3dZcalendar Sets the calendar system to use with `z` date data.
type Scatter3dZcalendar string

const (
	Scatter3dZcalendar_gregorian  Scatter3dZcalendar = "gregorian"
	Scatter3dZcalendar_chinese    Scatter3dZcalendar = "chinese"
	Scatter3dZcalendar_coptic     Scatter3dZcalendar = "coptic"
	Scatter3dZcalendar_discworld  Scatter3dZcalendar = "discworld"
	Scatter3dZcalendar_ethiopian  Scatter3dZcalendar = "ethiopian"
	Scatter3dZcalendar_hebrew     Scatter3dZcalendar = "hebrew"
	Scatter3dZcalendar_islamic    Scatter3dZcalendar = "islamic"
	Scatter3dZcalendar_julian     Scatter3dZcalendar = "julian"
	Scatter3dZcalendar_mayan      Scatter3dZcalendar = "mayan"
	Scatter3dZcalendar_nanakshahi Scatter3dZcalendar = "nanakshahi"
	Scatter3dZcalendar_nepali     Scatter3dZcalendar = "nepali"
	Scatter3dZcalendar_persian    Scatter3dZcalendar = "persian"
	Scatter3dZcalendar_jalali     Scatter3dZcalendar = "jalali"
	Scatter3dZcalendar_taiwan     Scatter3dZcalendar = "taiwan"
	Scatter3dZcalendar_thai       Scatter3dZcalendar = "thai"
	Scatter3dZcalendar_ummalqura  Scatter3dZcalendar = "ummalqura"
)

// ScatterErrorXType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type ScatterErrorXType string

const (
	ScatterErrorXType_percent  ScatterErrorXType = "percent"
	ScatterErrorXType_constant ScatterErrorXType = "constant"
	ScatterErrorXType_sqrt     ScatterErrorXType = "sqrt"
	ScatterErrorXType_data     ScatterErrorXType = "data"
)

// ScatterErrorYType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type ScatterErrorYType string

const (
	ScatterErrorYType_percent  ScatterErrorYType = "percent"
	ScatterErrorYType_constant ScatterErrorYType = "constant"
	ScatterErrorYType_sqrt     ScatterErrorYType = "sqrt"
	ScatterErrorYType_data     ScatterErrorYType = "data"
)

// ScatterFill Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
type ScatterFill string

const (
	ScatterFill_none    ScatterFill = "none"
	ScatterFill_tozeroy ScatterFill = "tozeroy"
	ScatterFill_tozerox ScatterFill = "tozerox"
	ScatterFill_tonexty ScatterFill = "tonexty"
	ScatterFill_tonextx ScatterFill = "tonextx"
	ScatterFill_toself  ScatterFill = "toself"
	ScatterFill_tonext  ScatterFill = "tonext"
)

// ScatterGroupnorm Only relevant when `stackgroup` is used, and only the first `groupnorm` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Sets the normalization for the sum of this `stackgroup`. With *fraction*, the value of each trace at each location is divided by the sum of all trace values at that location. *percent* is the same but multiplied by 100 to show percentages. If there are multiple subplots, or multiple `stackgroup`s on one subplot, each will be normalized within its own set.
type ScatterGroupnorm string

const (
	ScatterGroupnorm_fraction ScatterGroupnorm = "fraction"
	ScatterGroupnorm_percent  ScatterGroupnorm = "percent"
)

// ScatterHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterHoverlabelAlign string

const (
	ScatterHoverlabelAlign_left  ScatterHoverlabelAlign = "left"
	ScatterHoverlabelAlign_right ScatterHoverlabelAlign = "right"
	ScatterHoverlabelAlign_auto  ScatterHoverlabelAlign = "auto"
)

// ScatterLineShape Determines the line shape. With *spline* the lines are drawn using spline interpolation. The other available values correspond to step-wise line shapes.
type ScatterLineShape string

const (
	ScatterLineShape_linear ScatterLineShape = "linear"
	ScatterLineShape_spline ScatterLineShape = "spline"
	ScatterLineShape_hv     ScatterLineShape = "hv"
	ScatterLineShape_vh     ScatterLineShape = "vh"
	ScatterLineShape_hvh    ScatterLineShape = "hvh"
	ScatterLineShape_vhv    ScatterLineShape = "vhv"
)

// ScatterMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterMarkerColorbarExponentformat string

const (
	ScatterMarkerColorbarExponentformat_none  ScatterMarkerColorbarExponentformat = "none"
	ScatterMarkerColorbarExponentformat_e     ScatterMarkerColorbarExponentformat = "e"
	ScatterMarkerColorbarExponentformat_E     ScatterMarkerColorbarExponentformat = "E"
	ScatterMarkerColorbarExponentformat_power ScatterMarkerColorbarExponentformat = "power"
	ScatterMarkerColorbarExponentformat_SI    ScatterMarkerColorbarExponentformat = "SI"
	ScatterMarkerColorbarExponentformat_B     ScatterMarkerColorbarExponentformat = "B"
)

// ScatterMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterMarkerColorbarLenmode string

const (
	ScatterMarkerColorbarLenmode_fraction ScatterMarkerColorbarLenmode = "fraction"
	ScatterMarkerColorbarLenmode_pixels   ScatterMarkerColorbarLenmode = "pixels"
)

// ScatterMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterMarkerColorbarShowexponent string

const (
	ScatterMarkerColorbarShowexponent_all   ScatterMarkerColorbarShowexponent = "all"
	ScatterMarkerColorbarShowexponent_first ScatterMarkerColorbarShowexponent = "first"
	ScatterMarkerColorbarShowexponent_last  ScatterMarkerColorbarShowexponent = "last"
	ScatterMarkerColorbarShowexponent_none  ScatterMarkerColorbarShowexponent = "none"
)

// ScatterMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterMarkerColorbarShowtickprefix string

const (
	ScatterMarkerColorbarShowtickprefix_all   ScatterMarkerColorbarShowtickprefix = "all"
	ScatterMarkerColorbarShowtickprefix_first ScatterMarkerColorbarShowtickprefix = "first"
	ScatterMarkerColorbarShowtickprefix_last  ScatterMarkerColorbarShowtickprefix = "last"
	ScatterMarkerColorbarShowtickprefix_none  ScatterMarkerColorbarShowtickprefix = "none"
)

// ScatterMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterMarkerColorbarShowticksuffix string

const (
	ScatterMarkerColorbarShowticksuffix_all   ScatterMarkerColorbarShowticksuffix = "all"
	ScatterMarkerColorbarShowticksuffix_first ScatterMarkerColorbarShowticksuffix = "first"
	ScatterMarkerColorbarShowticksuffix_last  ScatterMarkerColorbarShowticksuffix = "last"
	ScatterMarkerColorbarShowticksuffix_none  ScatterMarkerColorbarShowticksuffix = "none"
)

// ScatterMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterMarkerColorbarThicknessmode string

const (
	ScatterMarkerColorbarThicknessmode_fraction ScatterMarkerColorbarThicknessmode = "fraction"
	ScatterMarkerColorbarThicknessmode_pixels   ScatterMarkerColorbarThicknessmode = "pixels"
)

// ScatterMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterMarkerColorbarTickmode string

const (
	ScatterMarkerColorbarTickmode_auto   ScatterMarkerColorbarTickmode = "auto"
	ScatterMarkerColorbarTickmode_linear ScatterMarkerColorbarTickmode = "linear"
	ScatterMarkerColorbarTickmode_array  ScatterMarkerColorbarTickmode = "array"
)

// ScatterMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterMarkerColorbarTicks string

const (
	ScatterMarkerColorbarTicks_outside ScatterMarkerColorbarTicks = "outside"
	ScatterMarkerColorbarTicks_inside  ScatterMarkerColorbarTicks = "inside"
)

// ScatterMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterMarkerColorbarTitleSide string

const (
	ScatterMarkerColorbarTitleSide_right  ScatterMarkerColorbarTitleSide = "right"
	ScatterMarkerColorbarTitleSide_top    ScatterMarkerColorbarTitleSide = "top"
	ScatterMarkerColorbarTitleSide_bottom ScatterMarkerColorbarTitleSide = "bottom"
)

// ScatterMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterMarkerColorbarXanchor string

const (
	ScatterMarkerColorbarXanchor_left   ScatterMarkerColorbarXanchor = "left"
	ScatterMarkerColorbarXanchor_center ScatterMarkerColorbarXanchor = "center"
	ScatterMarkerColorbarXanchor_right  ScatterMarkerColorbarXanchor = "right"
)

// ScatterMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterMarkerColorbarYanchor string

const (
	ScatterMarkerColorbarYanchor_top    ScatterMarkerColorbarYanchor = "top"
	ScatterMarkerColorbarYanchor_middle ScatterMarkerColorbarYanchor = "middle"
	ScatterMarkerColorbarYanchor_bottom ScatterMarkerColorbarYanchor = "bottom"
)

// ScatterMarkerGradientType Sets the type of gradient used to fill the markers
type ScatterMarkerGradientType string

const (
	ScatterMarkerGradientType_radial     ScatterMarkerGradientType = "radial"
	ScatterMarkerGradientType_horizontal ScatterMarkerGradientType = "horizontal"
	ScatterMarkerGradientType_vertical   ScatterMarkerGradientType = "vertical"
	ScatterMarkerGradientType_none       ScatterMarkerGradientType = "none"
)

// ScatterMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterMarkerSizemode string

const (
	ScatterMarkerSizemode_diameter ScatterMarkerSizemode = "diameter"
	ScatterMarkerSizemode_area     ScatterMarkerSizemode = "area"
)

// ScatterMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterMarkerSymbol string

const (
	ScatterMarkerSymbol0                            ScatterMarkerSymbol = "0"
	ScatterMarkerSymbol_circle                      ScatterMarkerSymbol = "circle"
	ScatterMarkerSymbol100                          ScatterMarkerSymbol = "100"
	ScatterMarkerSymbol_circle_open                 ScatterMarkerSymbol = "circle-open"
	ScatterMarkerSymbol200                          ScatterMarkerSymbol = "200"
	ScatterMarkerSymbol_circle_dot                  ScatterMarkerSymbol = "circle-dot"
	ScatterMarkerSymbol300                          ScatterMarkerSymbol = "300"
	ScatterMarkerSymbol_circle_open_dot             ScatterMarkerSymbol = "circle-open-dot"
	ScatterMarkerSymbol1                            ScatterMarkerSymbol = "1"
	ScatterMarkerSymbol_square                      ScatterMarkerSymbol = "square"
	ScatterMarkerSymbol101                          ScatterMarkerSymbol = "101"
	ScatterMarkerSymbol_square_open                 ScatterMarkerSymbol = "square-open"
	ScatterMarkerSymbol201                          ScatterMarkerSymbol = "201"
	ScatterMarkerSymbol_square_dot                  ScatterMarkerSymbol = "square-dot"
	ScatterMarkerSymbol301                          ScatterMarkerSymbol = "301"
	ScatterMarkerSymbol_square_open_dot             ScatterMarkerSymbol = "square-open-dot"
	ScatterMarkerSymbol2                            ScatterMarkerSymbol = "2"
	ScatterMarkerSymbol_diamond                     ScatterMarkerSymbol = "diamond"
	ScatterMarkerSymbol102                          ScatterMarkerSymbol = "102"
	ScatterMarkerSymbol_diamond_open                ScatterMarkerSymbol = "diamond-open"
	ScatterMarkerSymbol202                          ScatterMarkerSymbol = "202"
	ScatterMarkerSymbol_diamond_dot                 ScatterMarkerSymbol = "diamond-dot"
	ScatterMarkerSymbol302                          ScatterMarkerSymbol = "302"
	ScatterMarkerSymbol_diamond_open_dot            ScatterMarkerSymbol = "diamond-open-dot"
	ScatterMarkerSymbol3                            ScatterMarkerSymbol = "3"
	ScatterMarkerSymbol_cross                       ScatterMarkerSymbol = "cross"
	ScatterMarkerSymbol103                          ScatterMarkerSymbol = "103"
	ScatterMarkerSymbol_cross_open                  ScatterMarkerSymbol = "cross-open"
	ScatterMarkerSymbol203                          ScatterMarkerSymbol = "203"
	ScatterMarkerSymbol_cross_dot                   ScatterMarkerSymbol = "cross-dot"
	ScatterMarkerSymbol303                          ScatterMarkerSymbol = "303"
	ScatterMarkerSymbol_cross_open_dot              ScatterMarkerSymbol = "cross-open-dot"
	ScatterMarkerSymbol4                            ScatterMarkerSymbol = "4"
	ScatterMarkerSymbol_x                           ScatterMarkerSymbol = "x"
	ScatterMarkerSymbol104                          ScatterMarkerSymbol = "104"
	ScatterMarkerSymbol_x_open                      ScatterMarkerSymbol = "x-open"
	ScatterMarkerSymbol204                          ScatterMarkerSymbol = "204"
	ScatterMarkerSymbol_x_dot                       ScatterMarkerSymbol = "x-dot"
	ScatterMarkerSymbol304                          ScatterMarkerSymbol = "304"
	ScatterMarkerSymbol_x_open_dot                  ScatterMarkerSymbol = "x-open-dot"
	ScatterMarkerSymbol5                            ScatterMarkerSymbol = "5"
	ScatterMarkerSymbol_triangle_up                 ScatterMarkerSymbol = "triangle-up"
	ScatterMarkerSymbol105                          ScatterMarkerSymbol = "105"
	ScatterMarkerSymbol_triangle_up_open            ScatterMarkerSymbol = "triangle-up-open"
	ScatterMarkerSymbol205                          ScatterMarkerSymbol = "205"
	ScatterMarkerSymbol_triangle_up_dot             ScatterMarkerSymbol = "triangle-up-dot"
	ScatterMarkerSymbol305                          ScatterMarkerSymbol = "305"
	ScatterMarkerSymbol_triangle_up_open_dot        ScatterMarkerSymbol = "triangle-up-open-dot"
	ScatterMarkerSymbol6                            ScatterMarkerSymbol = "6"
	ScatterMarkerSymbol_triangle_down               ScatterMarkerSymbol = "triangle-down"
	ScatterMarkerSymbol106                          ScatterMarkerSymbol = "106"
	ScatterMarkerSymbol_triangle_down_open          ScatterMarkerSymbol = "triangle-down-open"
	ScatterMarkerSymbol206                          ScatterMarkerSymbol = "206"
	ScatterMarkerSymbol_triangle_down_dot           ScatterMarkerSymbol = "triangle-down-dot"
	ScatterMarkerSymbol306                          ScatterMarkerSymbol = "306"
	ScatterMarkerSymbol_triangle_down_open_dot      ScatterMarkerSymbol = "triangle-down-open-dot"
	ScatterMarkerSymbol7                            ScatterMarkerSymbol = "7"
	ScatterMarkerSymbol_triangle_left               ScatterMarkerSymbol = "triangle-left"
	ScatterMarkerSymbol107                          ScatterMarkerSymbol = "107"
	ScatterMarkerSymbol_triangle_left_open          ScatterMarkerSymbol = "triangle-left-open"
	ScatterMarkerSymbol207                          ScatterMarkerSymbol = "207"
	ScatterMarkerSymbol_triangle_left_dot           ScatterMarkerSymbol = "triangle-left-dot"
	ScatterMarkerSymbol307                          ScatterMarkerSymbol = "307"
	ScatterMarkerSymbol_triangle_left_open_dot      ScatterMarkerSymbol = "triangle-left-open-dot"
	ScatterMarkerSymbol8                            ScatterMarkerSymbol = "8"
	ScatterMarkerSymbol_triangle_right              ScatterMarkerSymbol = "triangle-right"
	ScatterMarkerSymbol108                          ScatterMarkerSymbol = "108"
	ScatterMarkerSymbol_triangle_right_open         ScatterMarkerSymbol = "triangle-right-open"
	ScatterMarkerSymbol208                          ScatterMarkerSymbol = "208"
	ScatterMarkerSymbol_triangle_right_dot          ScatterMarkerSymbol = "triangle-right-dot"
	ScatterMarkerSymbol308                          ScatterMarkerSymbol = "308"
	ScatterMarkerSymbol_triangle_right_open_dot     ScatterMarkerSymbol = "triangle-right-open-dot"
	ScatterMarkerSymbol9                            ScatterMarkerSymbol = "9"
	ScatterMarkerSymbol_triangle_ne                 ScatterMarkerSymbol = "triangle-ne"
	ScatterMarkerSymbol109                          ScatterMarkerSymbol = "109"
	ScatterMarkerSymbol_triangle_ne_open            ScatterMarkerSymbol = "triangle-ne-open"
	ScatterMarkerSymbol209                          ScatterMarkerSymbol = "209"
	ScatterMarkerSymbol_triangle_ne_dot             ScatterMarkerSymbol = "triangle-ne-dot"
	ScatterMarkerSymbol309                          ScatterMarkerSymbol = "309"
	ScatterMarkerSymbol_triangle_ne_open_dot        ScatterMarkerSymbol = "triangle-ne-open-dot"
	ScatterMarkerSymbol10                           ScatterMarkerSymbol = "10"
	ScatterMarkerSymbol_triangle_se                 ScatterMarkerSymbol = "triangle-se"
	ScatterMarkerSymbol110                          ScatterMarkerSymbol = "110"
	ScatterMarkerSymbol_triangle_se_open            ScatterMarkerSymbol = "triangle-se-open"
	ScatterMarkerSymbol210                          ScatterMarkerSymbol = "210"
	ScatterMarkerSymbol_triangle_se_dot             ScatterMarkerSymbol = "triangle-se-dot"
	ScatterMarkerSymbol310                          ScatterMarkerSymbol = "310"
	ScatterMarkerSymbol_triangle_se_open_dot        ScatterMarkerSymbol = "triangle-se-open-dot"
	ScatterMarkerSymbol11                           ScatterMarkerSymbol = "11"
	ScatterMarkerSymbol_triangle_sw                 ScatterMarkerSymbol = "triangle-sw"
	ScatterMarkerSymbol111                          ScatterMarkerSymbol = "111"
	ScatterMarkerSymbol_triangle_sw_open            ScatterMarkerSymbol = "triangle-sw-open"
	ScatterMarkerSymbol211                          ScatterMarkerSymbol = "211"
	ScatterMarkerSymbol_triangle_sw_dot             ScatterMarkerSymbol = "triangle-sw-dot"
	ScatterMarkerSymbol311                          ScatterMarkerSymbol = "311"
	ScatterMarkerSymbol_triangle_sw_open_dot        ScatterMarkerSymbol = "triangle-sw-open-dot"
	ScatterMarkerSymbol12                           ScatterMarkerSymbol = "12"
	ScatterMarkerSymbol_triangle_nw                 ScatterMarkerSymbol = "triangle-nw"
	ScatterMarkerSymbol112                          ScatterMarkerSymbol = "112"
	ScatterMarkerSymbol_triangle_nw_open            ScatterMarkerSymbol = "triangle-nw-open"
	ScatterMarkerSymbol212                          ScatterMarkerSymbol = "212"
	ScatterMarkerSymbol_triangle_nw_dot             ScatterMarkerSymbol = "triangle-nw-dot"
	ScatterMarkerSymbol312                          ScatterMarkerSymbol = "312"
	ScatterMarkerSymbol_triangle_nw_open_dot        ScatterMarkerSymbol = "triangle-nw-open-dot"
	ScatterMarkerSymbol13                           ScatterMarkerSymbol = "13"
	ScatterMarkerSymbol_pentagon                    ScatterMarkerSymbol = "pentagon"
	ScatterMarkerSymbol113                          ScatterMarkerSymbol = "113"
	ScatterMarkerSymbol_pentagon_open               ScatterMarkerSymbol = "pentagon-open"
	ScatterMarkerSymbol213                          ScatterMarkerSymbol = "213"
	ScatterMarkerSymbol_pentagon_dot                ScatterMarkerSymbol = "pentagon-dot"
	ScatterMarkerSymbol313                          ScatterMarkerSymbol = "313"
	ScatterMarkerSymbol_pentagon_open_dot           ScatterMarkerSymbol = "pentagon-open-dot"
	ScatterMarkerSymbol14                           ScatterMarkerSymbol = "14"
	ScatterMarkerSymbol_hexagon                     ScatterMarkerSymbol = "hexagon"
	ScatterMarkerSymbol114                          ScatterMarkerSymbol = "114"
	ScatterMarkerSymbol_hexagon_open                ScatterMarkerSymbol = "hexagon-open"
	ScatterMarkerSymbol214                          ScatterMarkerSymbol = "214"
	ScatterMarkerSymbol_hexagon_dot                 ScatterMarkerSymbol = "hexagon-dot"
	ScatterMarkerSymbol314                          ScatterMarkerSymbol = "314"
	ScatterMarkerSymbol_hexagon_open_dot            ScatterMarkerSymbol = "hexagon-open-dot"
	ScatterMarkerSymbol15                           ScatterMarkerSymbol = "15"
	ScatterMarkerSymbol_hexagon2                    ScatterMarkerSymbol = "hexagon2"
	ScatterMarkerSymbol115                          ScatterMarkerSymbol = "115"
	ScatterMarkerSymbol_hexagon2_open               ScatterMarkerSymbol = "hexagon2-open"
	ScatterMarkerSymbol215                          ScatterMarkerSymbol = "215"
	ScatterMarkerSymbol_hexagon2_dot                ScatterMarkerSymbol = "hexagon2-dot"
	ScatterMarkerSymbol315                          ScatterMarkerSymbol = "315"
	ScatterMarkerSymbol_hexagon2_open_dot           ScatterMarkerSymbol = "hexagon2-open-dot"
	ScatterMarkerSymbol16                           ScatterMarkerSymbol = "16"
	ScatterMarkerSymbol_octagon                     ScatterMarkerSymbol = "octagon"
	ScatterMarkerSymbol116                          ScatterMarkerSymbol = "116"
	ScatterMarkerSymbol_octagon_open                ScatterMarkerSymbol = "octagon-open"
	ScatterMarkerSymbol216                          ScatterMarkerSymbol = "216"
	ScatterMarkerSymbol_octagon_dot                 ScatterMarkerSymbol = "octagon-dot"
	ScatterMarkerSymbol316                          ScatterMarkerSymbol = "316"
	ScatterMarkerSymbol_octagon_open_dot            ScatterMarkerSymbol = "octagon-open-dot"
	ScatterMarkerSymbol17                           ScatterMarkerSymbol = "17"
	ScatterMarkerSymbol_star                        ScatterMarkerSymbol = "star"
	ScatterMarkerSymbol117                          ScatterMarkerSymbol = "117"
	ScatterMarkerSymbol_star_open                   ScatterMarkerSymbol = "star-open"
	ScatterMarkerSymbol217                          ScatterMarkerSymbol = "217"
	ScatterMarkerSymbol_star_dot                    ScatterMarkerSymbol = "star-dot"
	ScatterMarkerSymbol317                          ScatterMarkerSymbol = "317"
	ScatterMarkerSymbol_star_open_dot               ScatterMarkerSymbol = "star-open-dot"
	ScatterMarkerSymbol18                           ScatterMarkerSymbol = "18"
	ScatterMarkerSymbol_hexagram                    ScatterMarkerSymbol = "hexagram"
	ScatterMarkerSymbol118                          ScatterMarkerSymbol = "118"
	ScatterMarkerSymbol_hexagram_open               ScatterMarkerSymbol = "hexagram-open"
	ScatterMarkerSymbol218                          ScatterMarkerSymbol = "218"
	ScatterMarkerSymbol_hexagram_dot                ScatterMarkerSymbol = "hexagram-dot"
	ScatterMarkerSymbol318                          ScatterMarkerSymbol = "318"
	ScatterMarkerSymbol_hexagram_open_dot           ScatterMarkerSymbol = "hexagram-open-dot"
	ScatterMarkerSymbol19                           ScatterMarkerSymbol = "19"
	ScatterMarkerSymbol_star_triangle_up            ScatterMarkerSymbol = "star-triangle-up"
	ScatterMarkerSymbol119                          ScatterMarkerSymbol = "119"
	ScatterMarkerSymbol_star_triangle_up_open       ScatterMarkerSymbol = "star-triangle-up-open"
	ScatterMarkerSymbol219                          ScatterMarkerSymbol = "219"
	ScatterMarkerSymbol_star_triangle_up_dot        ScatterMarkerSymbol = "star-triangle-up-dot"
	ScatterMarkerSymbol319                          ScatterMarkerSymbol = "319"
	ScatterMarkerSymbol_star_triangle_up_open_dot   ScatterMarkerSymbol = "star-triangle-up-open-dot"
	ScatterMarkerSymbol20                           ScatterMarkerSymbol = "20"
	ScatterMarkerSymbol_star_triangle_down          ScatterMarkerSymbol = "star-triangle-down"
	ScatterMarkerSymbol120                          ScatterMarkerSymbol = "120"
	ScatterMarkerSymbol_star_triangle_down_open     ScatterMarkerSymbol = "star-triangle-down-open"
	ScatterMarkerSymbol220                          ScatterMarkerSymbol = "220"
	ScatterMarkerSymbol_star_triangle_down_dot      ScatterMarkerSymbol = "star-triangle-down-dot"
	ScatterMarkerSymbol320                          ScatterMarkerSymbol = "320"
	ScatterMarkerSymbol_star_triangle_down_open_dot ScatterMarkerSymbol = "star-triangle-down-open-dot"
	ScatterMarkerSymbol21                           ScatterMarkerSymbol = "21"
	ScatterMarkerSymbol_star_square                 ScatterMarkerSymbol = "star-square"
	ScatterMarkerSymbol121                          ScatterMarkerSymbol = "121"
	ScatterMarkerSymbol_star_square_open            ScatterMarkerSymbol = "star-square-open"
	ScatterMarkerSymbol221                          ScatterMarkerSymbol = "221"
	ScatterMarkerSymbol_star_square_dot             ScatterMarkerSymbol = "star-square-dot"
	ScatterMarkerSymbol321                          ScatterMarkerSymbol = "321"
	ScatterMarkerSymbol_star_square_open_dot        ScatterMarkerSymbol = "star-square-open-dot"
	ScatterMarkerSymbol22                           ScatterMarkerSymbol = "22"
	ScatterMarkerSymbol_star_diamond                ScatterMarkerSymbol = "star-diamond"
	ScatterMarkerSymbol122                          ScatterMarkerSymbol = "122"
	ScatterMarkerSymbol_star_diamond_open           ScatterMarkerSymbol = "star-diamond-open"
	ScatterMarkerSymbol222                          ScatterMarkerSymbol = "222"
	ScatterMarkerSymbol_star_diamond_dot            ScatterMarkerSymbol = "star-diamond-dot"
	ScatterMarkerSymbol322                          ScatterMarkerSymbol = "322"
	ScatterMarkerSymbol_star_diamond_open_dot       ScatterMarkerSymbol = "star-diamond-open-dot"
	ScatterMarkerSymbol23                           ScatterMarkerSymbol = "23"
	ScatterMarkerSymbol_diamond_tall                ScatterMarkerSymbol = "diamond-tall"
	ScatterMarkerSymbol123                          ScatterMarkerSymbol = "123"
	ScatterMarkerSymbol_diamond_tall_open           ScatterMarkerSymbol = "diamond-tall-open"
	ScatterMarkerSymbol223                          ScatterMarkerSymbol = "223"
	ScatterMarkerSymbol_diamond_tall_dot            ScatterMarkerSymbol = "diamond-tall-dot"
	ScatterMarkerSymbol323                          ScatterMarkerSymbol = "323"
	ScatterMarkerSymbol_diamond_tall_open_dot       ScatterMarkerSymbol = "diamond-tall-open-dot"
	ScatterMarkerSymbol24                           ScatterMarkerSymbol = "24"
	ScatterMarkerSymbol_diamond_wide                ScatterMarkerSymbol = "diamond-wide"
	ScatterMarkerSymbol124                          ScatterMarkerSymbol = "124"
	ScatterMarkerSymbol_diamond_wide_open           ScatterMarkerSymbol = "diamond-wide-open"
	ScatterMarkerSymbol224                          ScatterMarkerSymbol = "224"
	ScatterMarkerSymbol_diamond_wide_dot            ScatterMarkerSymbol = "diamond-wide-dot"
	ScatterMarkerSymbol324                          ScatterMarkerSymbol = "324"
	ScatterMarkerSymbol_diamond_wide_open_dot       ScatterMarkerSymbol = "diamond-wide-open-dot"
	ScatterMarkerSymbol25                           ScatterMarkerSymbol = "25"
	ScatterMarkerSymbol_hourglass                   ScatterMarkerSymbol = "hourglass"
	ScatterMarkerSymbol125                          ScatterMarkerSymbol = "125"
	ScatterMarkerSymbol_hourglass_open              ScatterMarkerSymbol = "hourglass-open"
	ScatterMarkerSymbol26                           ScatterMarkerSymbol = "26"
	ScatterMarkerSymbol_bowtie                      ScatterMarkerSymbol = "bowtie"
	ScatterMarkerSymbol126                          ScatterMarkerSymbol = "126"
	ScatterMarkerSymbol_bowtie_open                 ScatterMarkerSymbol = "bowtie-open"
	ScatterMarkerSymbol27                           ScatterMarkerSymbol = "27"
	ScatterMarkerSymbol_circle_cross                ScatterMarkerSymbol = "circle-cross"
	ScatterMarkerSymbol127                          ScatterMarkerSymbol = "127"
	ScatterMarkerSymbol_circle_cross_open           ScatterMarkerSymbol = "circle-cross-open"
	ScatterMarkerSymbol28                           ScatterMarkerSymbol = "28"
	ScatterMarkerSymbol_circle_x                    ScatterMarkerSymbol = "circle-x"
	ScatterMarkerSymbol128                          ScatterMarkerSymbol = "128"
	ScatterMarkerSymbol_circle_x_open               ScatterMarkerSymbol = "circle-x-open"
	ScatterMarkerSymbol29                           ScatterMarkerSymbol = "29"
	ScatterMarkerSymbol_square_cross                ScatterMarkerSymbol = "square-cross"
	ScatterMarkerSymbol129                          ScatterMarkerSymbol = "129"
	ScatterMarkerSymbol_square_cross_open           ScatterMarkerSymbol = "square-cross-open"
	ScatterMarkerSymbol30                           ScatterMarkerSymbol = "30"
	ScatterMarkerSymbol_square_x                    ScatterMarkerSymbol = "square-x"
	ScatterMarkerSymbol130                          ScatterMarkerSymbol = "130"
	ScatterMarkerSymbol_square_x_open               ScatterMarkerSymbol = "square-x-open"
	ScatterMarkerSymbol31                           ScatterMarkerSymbol = "31"
	ScatterMarkerSymbol_diamond_cross               ScatterMarkerSymbol = "diamond-cross"
	ScatterMarkerSymbol131                          ScatterMarkerSymbol = "131"
	ScatterMarkerSymbol_diamond_cross_open          ScatterMarkerSymbol = "diamond-cross-open"
	ScatterMarkerSymbol32                           ScatterMarkerSymbol = "32"
	ScatterMarkerSymbol_diamond_x                   ScatterMarkerSymbol = "diamond-x"
	ScatterMarkerSymbol132                          ScatterMarkerSymbol = "132"
	ScatterMarkerSymbol_diamond_x_open              ScatterMarkerSymbol = "diamond-x-open"
	ScatterMarkerSymbol33                           ScatterMarkerSymbol = "33"
	ScatterMarkerSymbol_cross_thin                  ScatterMarkerSymbol = "cross-thin"
	ScatterMarkerSymbol133                          ScatterMarkerSymbol = "133"
	ScatterMarkerSymbol_cross_thin_open             ScatterMarkerSymbol = "cross-thin-open"
	ScatterMarkerSymbol34                           ScatterMarkerSymbol = "34"
	ScatterMarkerSymbol_x_thin                      ScatterMarkerSymbol = "x-thin"
	ScatterMarkerSymbol134                          ScatterMarkerSymbol = "134"
	ScatterMarkerSymbol_x_thin_open                 ScatterMarkerSymbol = "x-thin-open"
	ScatterMarkerSymbol35                           ScatterMarkerSymbol = "35"
	ScatterMarkerSymbol_asterisk                    ScatterMarkerSymbol = "asterisk"
	ScatterMarkerSymbol135                          ScatterMarkerSymbol = "135"
	ScatterMarkerSymbol_asterisk_open               ScatterMarkerSymbol = "asterisk-open"
	ScatterMarkerSymbol36                           ScatterMarkerSymbol = "36"
	ScatterMarkerSymbol_hash                        ScatterMarkerSymbol = "hash"
	ScatterMarkerSymbol136                          ScatterMarkerSymbol = "136"
	ScatterMarkerSymbol_hash_open                   ScatterMarkerSymbol = "hash-open"
	ScatterMarkerSymbol236                          ScatterMarkerSymbol = "236"
	ScatterMarkerSymbol_hash_dot                    ScatterMarkerSymbol = "hash-dot"
	ScatterMarkerSymbol336                          ScatterMarkerSymbol = "336"
	ScatterMarkerSymbol_hash_open_dot               ScatterMarkerSymbol = "hash-open-dot"
	ScatterMarkerSymbol37                           ScatterMarkerSymbol = "37"
	ScatterMarkerSymbol_y_up                        ScatterMarkerSymbol = "y-up"
	ScatterMarkerSymbol137                          ScatterMarkerSymbol = "137"
	ScatterMarkerSymbol_y_up_open                   ScatterMarkerSymbol = "y-up-open"
	ScatterMarkerSymbol38                           ScatterMarkerSymbol = "38"
	ScatterMarkerSymbol_y_down                      ScatterMarkerSymbol = "y-down"
	ScatterMarkerSymbol138                          ScatterMarkerSymbol = "138"
	ScatterMarkerSymbol_y_down_open                 ScatterMarkerSymbol = "y-down-open"
	ScatterMarkerSymbol39                           ScatterMarkerSymbol = "39"
	ScatterMarkerSymbol_y_left                      ScatterMarkerSymbol = "y-left"
	ScatterMarkerSymbol139                          ScatterMarkerSymbol = "139"
	ScatterMarkerSymbol_y_left_open                 ScatterMarkerSymbol = "y-left-open"
	ScatterMarkerSymbol40                           ScatterMarkerSymbol = "40"
	ScatterMarkerSymbol_y_right                     ScatterMarkerSymbol = "y-right"
	ScatterMarkerSymbol140                          ScatterMarkerSymbol = "140"
	ScatterMarkerSymbol_y_right_open                ScatterMarkerSymbol = "y-right-open"
	ScatterMarkerSymbol41                           ScatterMarkerSymbol = "41"
	ScatterMarkerSymbol_line_ew                     ScatterMarkerSymbol = "line-ew"
	ScatterMarkerSymbol141                          ScatterMarkerSymbol = "141"
	ScatterMarkerSymbol_line_ew_open                ScatterMarkerSymbol = "line-ew-open"
	ScatterMarkerSymbol42                           ScatterMarkerSymbol = "42"
	ScatterMarkerSymbol_line_ns                     ScatterMarkerSymbol = "line-ns"
	ScatterMarkerSymbol142                          ScatterMarkerSymbol = "142"
	ScatterMarkerSymbol_line_ns_open                ScatterMarkerSymbol = "line-ns-open"
	ScatterMarkerSymbol43                           ScatterMarkerSymbol = "43"
	ScatterMarkerSymbol_line_ne                     ScatterMarkerSymbol = "line-ne"
	ScatterMarkerSymbol143                          ScatterMarkerSymbol = "143"
	ScatterMarkerSymbol_line_ne_open                ScatterMarkerSymbol = "line-ne-open"
	ScatterMarkerSymbol44                           ScatterMarkerSymbol = "44"
	ScatterMarkerSymbol_line_nw                     ScatterMarkerSymbol = "line-nw"
	ScatterMarkerSymbol144                          ScatterMarkerSymbol = "144"
	ScatterMarkerSymbol_line_nw_open                ScatterMarkerSymbol = "line-nw-open"
)

// ScatterOrientation Only relevant when `stackgroup` is used, and only the first `orientation` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Sets the stacking direction. With *v* (*h*), the y (x) values of subsequent traces are added. Also affects the default value of `fill`.
type ScatterOrientation string

const (
	ScatterOrientation_v ScatterOrientation = "v"
	ScatterOrientation_h ScatterOrientation = "h"
)

// ScatterStackgaps Only relevant when `stackgroup` is used, and only the first `stackgaps` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Determines how we handle locations at which other traces in this group have data but this one does not. With *infer zero* we insert a zero at these locations. With *interpolate* we linearly interpolate between existing values, and extrapolate a constant beyond the existing values.
type ScatterStackgaps string

const (
	ScatterStackgaps_inferzero   ScatterStackgaps = "infer zero"
	ScatterStackgaps_interpolate ScatterStackgaps = "interpolate"
)

// ScatterTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterTextposition string

const (
	ScatterTextposition_topleft      ScatterTextposition = "top left"
	ScatterTextposition_topcenter    ScatterTextposition = "top center"
	ScatterTextposition_topright     ScatterTextposition = "top right"
	ScatterTextposition_middleleft   ScatterTextposition = "middle left"
	ScatterTextposition_middlecenter ScatterTextposition = "middle center"
	ScatterTextposition_middleright  ScatterTextposition = "middle right"
	ScatterTextposition_bottomleft   ScatterTextposition = "bottom left"
	ScatterTextposition_bottomcenter ScatterTextposition = "bottom center"
	ScatterTextposition_bottomright  ScatterTextposition = "bottom right"
)

// ScatterVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterVisible interface{}

var (
	ScatterVisible_True       ScatterVisible = true
	ScatterVisible_False      ScatterVisible = false
	ScatterVisible_legendonly ScatterVisible = "legendonly"
)

// ScatterXcalendar Sets the calendar system to use with `x` date data.
type ScatterXcalendar string

const (
	ScatterXcalendar_gregorian  ScatterXcalendar = "gregorian"
	ScatterXcalendar_chinese    ScatterXcalendar = "chinese"
	ScatterXcalendar_coptic     ScatterXcalendar = "coptic"
	ScatterXcalendar_discworld  ScatterXcalendar = "discworld"
	ScatterXcalendar_ethiopian  ScatterXcalendar = "ethiopian"
	ScatterXcalendar_hebrew     ScatterXcalendar = "hebrew"
	ScatterXcalendar_islamic    ScatterXcalendar = "islamic"
	ScatterXcalendar_julian     ScatterXcalendar = "julian"
	ScatterXcalendar_mayan      ScatterXcalendar = "mayan"
	ScatterXcalendar_nanakshahi ScatterXcalendar = "nanakshahi"
	ScatterXcalendar_nepali     ScatterXcalendar = "nepali"
	ScatterXcalendar_persian    ScatterXcalendar = "persian"
	ScatterXcalendar_jalali     ScatterXcalendar = "jalali"
	ScatterXcalendar_taiwan     ScatterXcalendar = "taiwan"
	ScatterXcalendar_thai       ScatterXcalendar = "thai"
	ScatterXcalendar_ummalqura  ScatterXcalendar = "ummalqura"
)

// ScatterYcalendar Sets the calendar system to use with `y` date data.
type ScatterYcalendar string

const (
	ScatterYcalendar_gregorian  ScatterYcalendar = "gregorian"
	ScatterYcalendar_chinese    ScatterYcalendar = "chinese"
	ScatterYcalendar_coptic     ScatterYcalendar = "coptic"
	ScatterYcalendar_discworld  ScatterYcalendar = "discworld"
	ScatterYcalendar_ethiopian  ScatterYcalendar = "ethiopian"
	ScatterYcalendar_hebrew     ScatterYcalendar = "hebrew"
	ScatterYcalendar_islamic    ScatterYcalendar = "islamic"
	ScatterYcalendar_julian     ScatterYcalendar = "julian"
	ScatterYcalendar_mayan      ScatterYcalendar = "mayan"
	ScatterYcalendar_nanakshahi ScatterYcalendar = "nanakshahi"
	ScatterYcalendar_nepali     ScatterYcalendar = "nepali"
	ScatterYcalendar_persian    ScatterYcalendar = "persian"
	ScatterYcalendar_jalali     ScatterYcalendar = "jalali"
	ScatterYcalendar_taiwan     ScatterYcalendar = "taiwan"
	ScatterYcalendar_thai       ScatterYcalendar = "thai"
	ScatterYcalendar_ummalqura  ScatterYcalendar = "ummalqura"
)

// ScattercarpetFill Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. scatterternary has a subset of the options available to scatter. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other.
type ScattercarpetFill string

const (
	ScattercarpetFill_none   ScattercarpetFill = "none"
	ScattercarpetFill_toself ScattercarpetFill = "toself"
	ScattercarpetFill_tonext ScattercarpetFill = "tonext"
)

// ScattercarpetHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScattercarpetHoverlabelAlign string

const (
	ScattercarpetHoverlabelAlign_left  ScattercarpetHoverlabelAlign = "left"
	ScattercarpetHoverlabelAlign_right ScattercarpetHoverlabelAlign = "right"
	ScattercarpetHoverlabelAlign_auto  ScattercarpetHoverlabelAlign = "auto"
)

// ScattercarpetLineShape Determines the line shape. With *spline* the lines are drawn using spline interpolation. The other available values correspond to step-wise line shapes.
type ScattercarpetLineShape string

const (
	ScattercarpetLineShape_linear ScattercarpetLineShape = "linear"
	ScattercarpetLineShape_spline ScattercarpetLineShape = "spline"
)

// ScattercarpetMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScattercarpetMarkerColorbarExponentformat string

const (
	ScattercarpetMarkerColorbarExponentformat_none  ScattercarpetMarkerColorbarExponentformat = "none"
	ScattercarpetMarkerColorbarExponentformat_e     ScattercarpetMarkerColorbarExponentformat = "e"
	ScattercarpetMarkerColorbarExponentformat_E     ScattercarpetMarkerColorbarExponentformat = "E"
	ScattercarpetMarkerColorbarExponentformat_power ScattercarpetMarkerColorbarExponentformat = "power"
	ScattercarpetMarkerColorbarExponentformat_SI    ScattercarpetMarkerColorbarExponentformat = "SI"
	ScattercarpetMarkerColorbarExponentformat_B     ScattercarpetMarkerColorbarExponentformat = "B"
)

// ScattercarpetMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScattercarpetMarkerColorbarLenmode string

const (
	ScattercarpetMarkerColorbarLenmode_fraction ScattercarpetMarkerColorbarLenmode = "fraction"
	ScattercarpetMarkerColorbarLenmode_pixels   ScattercarpetMarkerColorbarLenmode = "pixels"
)

// ScattercarpetMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScattercarpetMarkerColorbarShowexponent string

const (
	ScattercarpetMarkerColorbarShowexponent_all   ScattercarpetMarkerColorbarShowexponent = "all"
	ScattercarpetMarkerColorbarShowexponent_first ScattercarpetMarkerColorbarShowexponent = "first"
	ScattercarpetMarkerColorbarShowexponent_last  ScattercarpetMarkerColorbarShowexponent = "last"
	ScattercarpetMarkerColorbarShowexponent_none  ScattercarpetMarkerColorbarShowexponent = "none"
)

// ScattercarpetMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScattercarpetMarkerColorbarShowtickprefix string

const (
	ScattercarpetMarkerColorbarShowtickprefix_all   ScattercarpetMarkerColorbarShowtickprefix = "all"
	ScattercarpetMarkerColorbarShowtickprefix_first ScattercarpetMarkerColorbarShowtickprefix = "first"
	ScattercarpetMarkerColorbarShowtickprefix_last  ScattercarpetMarkerColorbarShowtickprefix = "last"
	ScattercarpetMarkerColorbarShowtickprefix_none  ScattercarpetMarkerColorbarShowtickprefix = "none"
)

// ScattercarpetMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScattercarpetMarkerColorbarShowticksuffix string

const (
	ScattercarpetMarkerColorbarShowticksuffix_all   ScattercarpetMarkerColorbarShowticksuffix = "all"
	ScattercarpetMarkerColorbarShowticksuffix_first ScattercarpetMarkerColorbarShowticksuffix = "first"
	ScattercarpetMarkerColorbarShowticksuffix_last  ScattercarpetMarkerColorbarShowticksuffix = "last"
	ScattercarpetMarkerColorbarShowticksuffix_none  ScattercarpetMarkerColorbarShowticksuffix = "none"
)

// ScattercarpetMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScattercarpetMarkerColorbarThicknessmode string

const (
	ScattercarpetMarkerColorbarThicknessmode_fraction ScattercarpetMarkerColorbarThicknessmode = "fraction"
	ScattercarpetMarkerColorbarThicknessmode_pixels   ScattercarpetMarkerColorbarThicknessmode = "pixels"
)

// ScattercarpetMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScattercarpetMarkerColorbarTickmode string

const (
	ScattercarpetMarkerColorbarTickmode_auto   ScattercarpetMarkerColorbarTickmode = "auto"
	ScattercarpetMarkerColorbarTickmode_linear ScattercarpetMarkerColorbarTickmode = "linear"
	ScattercarpetMarkerColorbarTickmode_array  ScattercarpetMarkerColorbarTickmode = "array"
)

// ScattercarpetMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScattercarpetMarkerColorbarTicks string

const (
	ScattercarpetMarkerColorbarTicks_outside ScattercarpetMarkerColorbarTicks = "outside"
	ScattercarpetMarkerColorbarTicks_inside  ScattercarpetMarkerColorbarTicks = "inside"
)

// ScattercarpetMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScattercarpetMarkerColorbarTitleSide string

const (
	ScattercarpetMarkerColorbarTitleSide_right  ScattercarpetMarkerColorbarTitleSide = "right"
	ScattercarpetMarkerColorbarTitleSide_top    ScattercarpetMarkerColorbarTitleSide = "top"
	ScattercarpetMarkerColorbarTitleSide_bottom ScattercarpetMarkerColorbarTitleSide = "bottom"
)

// ScattercarpetMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScattercarpetMarkerColorbarXanchor string

const (
	ScattercarpetMarkerColorbarXanchor_left   ScattercarpetMarkerColorbarXanchor = "left"
	ScattercarpetMarkerColorbarXanchor_center ScattercarpetMarkerColorbarXanchor = "center"
	ScattercarpetMarkerColorbarXanchor_right  ScattercarpetMarkerColorbarXanchor = "right"
)

// ScattercarpetMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScattercarpetMarkerColorbarYanchor string

const (
	ScattercarpetMarkerColorbarYanchor_top    ScattercarpetMarkerColorbarYanchor = "top"
	ScattercarpetMarkerColorbarYanchor_middle ScattercarpetMarkerColorbarYanchor = "middle"
	ScattercarpetMarkerColorbarYanchor_bottom ScattercarpetMarkerColorbarYanchor = "bottom"
)

// ScattercarpetMarkerGradientType Sets the type of gradient used to fill the markers
type ScattercarpetMarkerGradientType string

const (
	ScattercarpetMarkerGradientType_radial     ScattercarpetMarkerGradientType = "radial"
	ScattercarpetMarkerGradientType_horizontal ScattercarpetMarkerGradientType = "horizontal"
	ScattercarpetMarkerGradientType_vertical   ScattercarpetMarkerGradientType = "vertical"
	ScattercarpetMarkerGradientType_none       ScattercarpetMarkerGradientType = "none"
)

// ScattercarpetMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScattercarpetMarkerSizemode string

const (
	ScattercarpetMarkerSizemode_diameter ScattercarpetMarkerSizemode = "diameter"
	ScattercarpetMarkerSizemode_area     ScattercarpetMarkerSizemode = "area"
)

// ScattercarpetMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScattercarpetMarkerSymbol string

const (
	ScattercarpetMarkerSymbol0                            ScattercarpetMarkerSymbol = "0"
	ScattercarpetMarkerSymbol_circle                      ScattercarpetMarkerSymbol = "circle"
	ScattercarpetMarkerSymbol100                          ScattercarpetMarkerSymbol = "100"
	ScattercarpetMarkerSymbol_circle_open                 ScattercarpetMarkerSymbol = "circle-open"
	ScattercarpetMarkerSymbol200                          ScattercarpetMarkerSymbol = "200"
	ScattercarpetMarkerSymbol_circle_dot                  ScattercarpetMarkerSymbol = "circle-dot"
	ScattercarpetMarkerSymbol300                          ScattercarpetMarkerSymbol = "300"
	ScattercarpetMarkerSymbol_circle_open_dot             ScattercarpetMarkerSymbol = "circle-open-dot"
	ScattercarpetMarkerSymbol1                            ScattercarpetMarkerSymbol = "1"
	ScattercarpetMarkerSymbol_square                      ScattercarpetMarkerSymbol = "square"
	ScattercarpetMarkerSymbol101                          ScattercarpetMarkerSymbol = "101"
	ScattercarpetMarkerSymbol_square_open                 ScattercarpetMarkerSymbol = "square-open"
	ScattercarpetMarkerSymbol201                          ScattercarpetMarkerSymbol = "201"
	ScattercarpetMarkerSymbol_square_dot                  ScattercarpetMarkerSymbol = "square-dot"
	ScattercarpetMarkerSymbol301                          ScattercarpetMarkerSymbol = "301"
	ScattercarpetMarkerSymbol_square_open_dot             ScattercarpetMarkerSymbol = "square-open-dot"
	ScattercarpetMarkerSymbol2                            ScattercarpetMarkerSymbol = "2"
	ScattercarpetMarkerSymbol_diamond                     ScattercarpetMarkerSymbol = "diamond"
	ScattercarpetMarkerSymbol102                          ScattercarpetMarkerSymbol = "102"
	ScattercarpetMarkerSymbol_diamond_open                ScattercarpetMarkerSymbol = "diamond-open"
	ScattercarpetMarkerSymbol202                          ScattercarpetMarkerSymbol = "202"
	ScattercarpetMarkerSymbol_diamond_dot                 ScattercarpetMarkerSymbol = "diamond-dot"
	ScattercarpetMarkerSymbol302                          ScattercarpetMarkerSymbol = "302"
	ScattercarpetMarkerSymbol_diamond_open_dot            ScattercarpetMarkerSymbol = "diamond-open-dot"
	ScattercarpetMarkerSymbol3                            ScattercarpetMarkerSymbol = "3"
	ScattercarpetMarkerSymbol_cross                       ScattercarpetMarkerSymbol = "cross"
	ScattercarpetMarkerSymbol103                          ScattercarpetMarkerSymbol = "103"
	ScattercarpetMarkerSymbol_cross_open                  ScattercarpetMarkerSymbol = "cross-open"
	ScattercarpetMarkerSymbol203                          ScattercarpetMarkerSymbol = "203"
	ScattercarpetMarkerSymbol_cross_dot                   ScattercarpetMarkerSymbol = "cross-dot"
	ScattercarpetMarkerSymbol303                          ScattercarpetMarkerSymbol = "303"
	ScattercarpetMarkerSymbol_cross_open_dot              ScattercarpetMarkerSymbol = "cross-open-dot"
	ScattercarpetMarkerSymbol4                            ScattercarpetMarkerSymbol = "4"
	ScattercarpetMarkerSymbol_x                           ScattercarpetMarkerSymbol = "x"
	ScattercarpetMarkerSymbol104                          ScattercarpetMarkerSymbol = "104"
	ScattercarpetMarkerSymbol_x_open                      ScattercarpetMarkerSymbol = "x-open"
	ScattercarpetMarkerSymbol204                          ScattercarpetMarkerSymbol = "204"
	ScattercarpetMarkerSymbol_x_dot                       ScattercarpetMarkerSymbol = "x-dot"
	ScattercarpetMarkerSymbol304                          ScattercarpetMarkerSymbol = "304"
	ScattercarpetMarkerSymbol_x_open_dot                  ScattercarpetMarkerSymbol = "x-open-dot"
	ScattercarpetMarkerSymbol5                            ScattercarpetMarkerSymbol = "5"
	ScattercarpetMarkerSymbol_triangle_up                 ScattercarpetMarkerSymbol = "triangle-up"
	ScattercarpetMarkerSymbol105                          ScattercarpetMarkerSymbol = "105"
	ScattercarpetMarkerSymbol_triangle_up_open            ScattercarpetMarkerSymbol = "triangle-up-open"
	ScattercarpetMarkerSymbol205                          ScattercarpetMarkerSymbol = "205"
	ScattercarpetMarkerSymbol_triangle_up_dot             ScattercarpetMarkerSymbol = "triangle-up-dot"
	ScattercarpetMarkerSymbol305                          ScattercarpetMarkerSymbol = "305"
	ScattercarpetMarkerSymbol_triangle_up_open_dot        ScattercarpetMarkerSymbol = "triangle-up-open-dot"
	ScattercarpetMarkerSymbol6                            ScattercarpetMarkerSymbol = "6"
	ScattercarpetMarkerSymbol_triangle_down               ScattercarpetMarkerSymbol = "triangle-down"
	ScattercarpetMarkerSymbol106                          ScattercarpetMarkerSymbol = "106"
	ScattercarpetMarkerSymbol_triangle_down_open          ScattercarpetMarkerSymbol = "triangle-down-open"
	ScattercarpetMarkerSymbol206                          ScattercarpetMarkerSymbol = "206"
	ScattercarpetMarkerSymbol_triangle_down_dot           ScattercarpetMarkerSymbol = "triangle-down-dot"
	ScattercarpetMarkerSymbol306                          ScattercarpetMarkerSymbol = "306"
	ScattercarpetMarkerSymbol_triangle_down_open_dot      ScattercarpetMarkerSymbol = "triangle-down-open-dot"
	ScattercarpetMarkerSymbol7                            ScattercarpetMarkerSymbol = "7"
	ScattercarpetMarkerSymbol_triangle_left               ScattercarpetMarkerSymbol = "triangle-left"
	ScattercarpetMarkerSymbol107                          ScattercarpetMarkerSymbol = "107"
	ScattercarpetMarkerSymbol_triangle_left_open          ScattercarpetMarkerSymbol = "triangle-left-open"
	ScattercarpetMarkerSymbol207                          ScattercarpetMarkerSymbol = "207"
	ScattercarpetMarkerSymbol_triangle_left_dot           ScattercarpetMarkerSymbol = "triangle-left-dot"
	ScattercarpetMarkerSymbol307                          ScattercarpetMarkerSymbol = "307"
	ScattercarpetMarkerSymbol_triangle_left_open_dot      ScattercarpetMarkerSymbol = "triangle-left-open-dot"
	ScattercarpetMarkerSymbol8                            ScattercarpetMarkerSymbol = "8"
	ScattercarpetMarkerSymbol_triangle_right              ScattercarpetMarkerSymbol = "triangle-right"
	ScattercarpetMarkerSymbol108                          ScattercarpetMarkerSymbol = "108"
	ScattercarpetMarkerSymbol_triangle_right_open         ScattercarpetMarkerSymbol = "triangle-right-open"
	ScattercarpetMarkerSymbol208                          ScattercarpetMarkerSymbol = "208"
	ScattercarpetMarkerSymbol_triangle_right_dot          ScattercarpetMarkerSymbol = "triangle-right-dot"
	ScattercarpetMarkerSymbol308                          ScattercarpetMarkerSymbol = "308"
	ScattercarpetMarkerSymbol_triangle_right_open_dot     ScattercarpetMarkerSymbol = "triangle-right-open-dot"
	ScattercarpetMarkerSymbol9                            ScattercarpetMarkerSymbol = "9"
	ScattercarpetMarkerSymbol_triangle_ne                 ScattercarpetMarkerSymbol = "triangle-ne"
	ScattercarpetMarkerSymbol109                          ScattercarpetMarkerSymbol = "109"
	ScattercarpetMarkerSymbol_triangle_ne_open            ScattercarpetMarkerSymbol = "triangle-ne-open"
	ScattercarpetMarkerSymbol209                          ScattercarpetMarkerSymbol = "209"
	ScattercarpetMarkerSymbol_triangle_ne_dot             ScattercarpetMarkerSymbol = "triangle-ne-dot"
	ScattercarpetMarkerSymbol309                          ScattercarpetMarkerSymbol = "309"
	ScattercarpetMarkerSymbol_triangle_ne_open_dot        ScattercarpetMarkerSymbol = "triangle-ne-open-dot"
	ScattercarpetMarkerSymbol10                           ScattercarpetMarkerSymbol = "10"
	ScattercarpetMarkerSymbol_triangle_se                 ScattercarpetMarkerSymbol = "triangle-se"
	ScattercarpetMarkerSymbol110                          ScattercarpetMarkerSymbol = "110"
	ScattercarpetMarkerSymbol_triangle_se_open            ScattercarpetMarkerSymbol = "triangle-se-open"
	ScattercarpetMarkerSymbol210                          ScattercarpetMarkerSymbol = "210"
	ScattercarpetMarkerSymbol_triangle_se_dot             ScattercarpetMarkerSymbol = "triangle-se-dot"
	ScattercarpetMarkerSymbol310                          ScattercarpetMarkerSymbol = "310"
	ScattercarpetMarkerSymbol_triangle_se_open_dot        ScattercarpetMarkerSymbol = "triangle-se-open-dot"
	ScattercarpetMarkerSymbol11                           ScattercarpetMarkerSymbol = "11"
	ScattercarpetMarkerSymbol_triangle_sw                 ScattercarpetMarkerSymbol = "triangle-sw"
	ScattercarpetMarkerSymbol111                          ScattercarpetMarkerSymbol = "111"
	ScattercarpetMarkerSymbol_triangle_sw_open            ScattercarpetMarkerSymbol = "triangle-sw-open"
	ScattercarpetMarkerSymbol211                          ScattercarpetMarkerSymbol = "211"
	ScattercarpetMarkerSymbol_triangle_sw_dot             ScattercarpetMarkerSymbol = "triangle-sw-dot"
	ScattercarpetMarkerSymbol311                          ScattercarpetMarkerSymbol = "311"
	ScattercarpetMarkerSymbol_triangle_sw_open_dot        ScattercarpetMarkerSymbol = "triangle-sw-open-dot"
	ScattercarpetMarkerSymbol12                           ScattercarpetMarkerSymbol = "12"
	ScattercarpetMarkerSymbol_triangle_nw                 ScattercarpetMarkerSymbol = "triangle-nw"
	ScattercarpetMarkerSymbol112                          ScattercarpetMarkerSymbol = "112"
	ScattercarpetMarkerSymbol_triangle_nw_open            ScattercarpetMarkerSymbol = "triangle-nw-open"
	ScattercarpetMarkerSymbol212                          ScattercarpetMarkerSymbol = "212"
	ScattercarpetMarkerSymbol_triangle_nw_dot             ScattercarpetMarkerSymbol = "triangle-nw-dot"
	ScattercarpetMarkerSymbol312                          ScattercarpetMarkerSymbol = "312"
	ScattercarpetMarkerSymbol_triangle_nw_open_dot        ScattercarpetMarkerSymbol = "triangle-nw-open-dot"
	ScattercarpetMarkerSymbol13                           ScattercarpetMarkerSymbol = "13"
	ScattercarpetMarkerSymbol_pentagon                    ScattercarpetMarkerSymbol = "pentagon"
	ScattercarpetMarkerSymbol113                          ScattercarpetMarkerSymbol = "113"
	ScattercarpetMarkerSymbol_pentagon_open               ScattercarpetMarkerSymbol = "pentagon-open"
	ScattercarpetMarkerSymbol213                          ScattercarpetMarkerSymbol = "213"
	ScattercarpetMarkerSymbol_pentagon_dot                ScattercarpetMarkerSymbol = "pentagon-dot"
	ScattercarpetMarkerSymbol313                          ScattercarpetMarkerSymbol = "313"
	ScattercarpetMarkerSymbol_pentagon_open_dot           ScattercarpetMarkerSymbol = "pentagon-open-dot"
	ScattercarpetMarkerSymbol14                           ScattercarpetMarkerSymbol = "14"
	ScattercarpetMarkerSymbol_hexagon                     ScattercarpetMarkerSymbol = "hexagon"
	ScattercarpetMarkerSymbol114                          ScattercarpetMarkerSymbol = "114"
	ScattercarpetMarkerSymbol_hexagon_open                ScattercarpetMarkerSymbol = "hexagon-open"
	ScattercarpetMarkerSymbol214                          ScattercarpetMarkerSymbol = "214"
	ScattercarpetMarkerSymbol_hexagon_dot                 ScattercarpetMarkerSymbol = "hexagon-dot"
	ScattercarpetMarkerSymbol314                          ScattercarpetMarkerSymbol = "314"
	ScattercarpetMarkerSymbol_hexagon_open_dot            ScattercarpetMarkerSymbol = "hexagon-open-dot"
	ScattercarpetMarkerSymbol15                           ScattercarpetMarkerSymbol = "15"
	ScattercarpetMarkerSymbol_hexagon2                    ScattercarpetMarkerSymbol = "hexagon2"
	ScattercarpetMarkerSymbol115                          ScattercarpetMarkerSymbol = "115"
	ScattercarpetMarkerSymbol_hexagon2_open               ScattercarpetMarkerSymbol = "hexagon2-open"
	ScattercarpetMarkerSymbol215                          ScattercarpetMarkerSymbol = "215"
	ScattercarpetMarkerSymbol_hexagon2_dot                ScattercarpetMarkerSymbol = "hexagon2-dot"
	ScattercarpetMarkerSymbol315                          ScattercarpetMarkerSymbol = "315"
	ScattercarpetMarkerSymbol_hexagon2_open_dot           ScattercarpetMarkerSymbol = "hexagon2-open-dot"
	ScattercarpetMarkerSymbol16                           ScattercarpetMarkerSymbol = "16"
	ScattercarpetMarkerSymbol_octagon                     ScattercarpetMarkerSymbol = "octagon"
	ScattercarpetMarkerSymbol116                          ScattercarpetMarkerSymbol = "116"
	ScattercarpetMarkerSymbol_octagon_open                ScattercarpetMarkerSymbol = "octagon-open"
	ScattercarpetMarkerSymbol216                          ScattercarpetMarkerSymbol = "216"
	ScattercarpetMarkerSymbol_octagon_dot                 ScattercarpetMarkerSymbol = "octagon-dot"
	ScattercarpetMarkerSymbol316                          ScattercarpetMarkerSymbol = "316"
	ScattercarpetMarkerSymbol_octagon_open_dot            ScattercarpetMarkerSymbol = "octagon-open-dot"
	ScattercarpetMarkerSymbol17                           ScattercarpetMarkerSymbol = "17"
	ScattercarpetMarkerSymbol_star                        ScattercarpetMarkerSymbol = "star"
	ScattercarpetMarkerSymbol117                          ScattercarpetMarkerSymbol = "117"
	ScattercarpetMarkerSymbol_star_open                   ScattercarpetMarkerSymbol = "star-open"
	ScattercarpetMarkerSymbol217                          ScattercarpetMarkerSymbol = "217"
	ScattercarpetMarkerSymbol_star_dot                    ScattercarpetMarkerSymbol = "star-dot"
	ScattercarpetMarkerSymbol317                          ScattercarpetMarkerSymbol = "317"
	ScattercarpetMarkerSymbol_star_open_dot               ScattercarpetMarkerSymbol = "star-open-dot"
	ScattercarpetMarkerSymbol18                           ScattercarpetMarkerSymbol = "18"
	ScattercarpetMarkerSymbol_hexagram                    ScattercarpetMarkerSymbol = "hexagram"
	ScattercarpetMarkerSymbol118                          ScattercarpetMarkerSymbol = "118"
	ScattercarpetMarkerSymbol_hexagram_open               ScattercarpetMarkerSymbol = "hexagram-open"
	ScattercarpetMarkerSymbol218                          ScattercarpetMarkerSymbol = "218"
	ScattercarpetMarkerSymbol_hexagram_dot                ScattercarpetMarkerSymbol = "hexagram-dot"
	ScattercarpetMarkerSymbol318                          ScattercarpetMarkerSymbol = "318"
	ScattercarpetMarkerSymbol_hexagram_open_dot           ScattercarpetMarkerSymbol = "hexagram-open-dot"
	ScattercarpetMarkerSymbol19                           ScattercarpetMarkerSymbol = "19"
	ScattercarpetMarkerSymbol_star_triangle_up            ScattercarpetMarkerSymbol = "star-triangle-up"
	ScattercarpetMarkerSymbol119                          ScattercarpetMarkerSymbol = "119"
	ScattercarpetMarkerSymbol_star_triangle_up_open       ScattercarpetMarkerSymbol = "star-triangle-up-open"
	ScattercarpetMarkerSymbol219                          ScattercarpetMarkerSymbol = "219"
	ScattercarpetMarkerSymbol_star_triangle_up_dot        ScattercarpetMarkerSymbol = "star-triangle-up-dot"
	ScattercarpetMarkerSymbol319                          ScattercarpetMarkerSymbol = "319"
	ScattercarpetMarkerSymbol_star_triangle_up_open_dot   ScattercarpetMarkerSymbol = "star-triangle-up-open-dot"
	ScattercarpetMarkerSymbol20                           ScattercarpetMarkerSymbol = "20"
	ScattercarpetMarkerSymbol_star_triangle_down          ScattercarpetMarkerSymbol = "star-triangle-down"
	ScattercarpetMarkerSymbol120                          ScattercarpetMarkerSymbol = "120"
	ScattercarpetMarkerSymbol_star_triangle_down_open     ScattercarpetMarkerSymbol = "star-triangle-down-open"
	ScattercarpetMarkerSymbol220                          ScattercarpetMarkerSymbol = "220"
	ScattercarpetMarkerSymbol_star_triangle_down_dot      ScattercarpetMarkerSymbol = "star-triangle-down-dot"
	ScattercarpetMarkerSymbol320                          ScattercarpetMarkerSymbol = "320"
	ScattercarpetMarkerSymbol_star_triangle_down_open_dot ScattercarpetMarkerSymbol = "star-triangle-down-open-dot"
	ScattercarpetMarkerSymbol21                           ScattercarpetMarkerSymbol = "21"
	ScattercarpetMarkerSymbol_star_square                 ScattercarpetMarkerSymbol = "star-square"
	ScattercarpetMarkerSymbol121                          ScattercarpetMarkerSymbol = "121"
	ScattercarpetMarkerSymbol_star_square_open            ScattercarpetMarkerSymbol = "star-square-open"
	ScattercarpetMarkerSymbol221                          ScattercarpetMarkerSymbol = "221"
	ScattercarpetMarkerSymbol_star_square_dot             ScattercarpetMarkerSymbol = "star-square-dot"
	ScattercarpetMarkerSymbol321                          ScattercarpetMarkerSymbol = "321"
	ScattercarpetMarkerSymbol_star_square_open_dot        ScattercarpetMarkerSymbol = "star-square-open-dot"
	ScattercarpetMarkerSymbol22                           ScattercarpetMarkerSymbol = "22"
	ScattercarpetMarkerSymbol_star_diamond                ScattercarpetMarkerSymbol = "star-diamond"
	ScattercarpetMarkerSymbol122                          ScattercarpetMarkerSymbol = "122"
	ScattercarpetMarkerSymbol_star_diamond_open           ScattercarpetMarkerSymbol = "star-diamond-open"
	ScattercarpetMarkerSymbol222                          ScattercarpetMarkerSymbol = "222"
	ScattercarpetMarkerSymbol_star_diamond_dot            ScattercarpetMarkerSymbol = "star-diamond-dot"
	ScattercarpetMarkerSymbol322                          ScattercarpetMarkerSymbol = "322"
	ScattercarpetMarkerSymbol_star_diamond_open_dot       ScattercarpetMarkerSymbol = "star-diamond-open-dot"
	ScattercarpetMarkerSymbol23                           ScattercarpetMarkerSymbol = "23"
	ScattercarpetMarkerSymbol_diamond_tall                ScattercarpetMarkerSymbol = "diamond-tall"
	ScattercarpetMarkerSymbol123                          ScattercarpetMarkerSymbol = "123"
	ScattercarpetMarkerSymbol_diamond_tall_open           ScattercarpetMarkerSymbol = "diamond-tall-open"
	ScattercarpetMarkerSymbol223                          ScattercarpetMarkerSymbol = "223"
	ScattercarpetMarkerSymbol_diamond_tall_dot            ScattercarpetMarkerSymbol = "diamond-tall-dot"
	ScattercarpetMarkerSymbol323                          ScattercarpetMarkerSymbol = "323"
	ScattercarpetMarkerSymbol_diamond_tall_open_dot       ScattercarpetMarkerSymbol = "diamond-tall-open-dot"
	ScattercarpetMarkerSymbol24                           ScattercarpetMarkerSymbol = "24"
	ScattercarpetMarkerSymbol_diamond_wide                ScattercarpetMarkerSymbol = "diamond-wide"
	ScattercarpetMarkerSymbol124                          ScattercarpetMarkerSymbol = "124"
	ScattercarpetMarkerSymbol_diamond_wide_open           ScattercarpetMarkerSymbol = "diamond-wide-open"
	ScattercarpetMarkerSymbol224                          ScattercarpetMarkerSymbol = "224"
	ScattercarpetMarkerSymbol_diamond_wide_dot            ScattercarpetMarkerSymbol = "diamond-wide-dot"
	ScattercarpetMarkerSymbol324                          ScattercarpetMarkerSymbol = "324"
	ScattercarpetMarkerSymbol_diamond_wide_open_dot       ScattercarpetMarkerSymbol = "diamond-wide-open-dot"
	ScattercarpetMarkerSymbol25                           ScattercarpetMarkerSymbol = "25"
	ScattercarpetMarkerSymbol_hourglass                   ScattercarpetMarkerSymbol = "hourglass"
	ScattercarpetMarkerSymbol125                          ScattercarpetMarkerSymbol = "125"
	ScattercarpetMarkerSymbol_hourglass_open              ScattercarpetMarkerSymbol = "hourglass-open"
	ScattercarpetMarkerSymbol26                           ScattercarpetMarkerSymbol = "26"
	ScattercarpetMarkerSymbol_bowtie                      ScattercarpetMarkerSymbol = "bowtie"
	ScattercarpetMarkerSymbol126                          ScattercarpetMarkerSymbol = "126"
	ScattercarpetMarkerSymbol_bowtie_open                 ScattercarpetMarkerSymbol = "bowtie-open"
	ScattercarpetMarkerSymbol27                           ScattercarpetMarkerSymbol = "27"
	ScattercarpetMarkerSymbol_circle_cross                ScattercarpetMarkerSymbol = "circle-cross"
	ScattercarpetMarkerSymbol127                          ScattercarpetMarkerSymbol = "127"
	ScattercarpetMarkerSymbol_circle_cross_open           ScattercarpetMarkerSymbol = "circle-cross-open"
	ScattercarpetMarkerSymbol28                           ScattercarpetMarkerSymbol = "28"
	ScattercarpetMarkerSymbol_circle_x                    ScattercarpetMarkerSymbol = "circle-x"
	ScattercarpetMarkerSymbol128                          ScattercarpetMarkerSymbol = "128"
	ScattercarpetMarkerSymbol_circle_x_open               ScattercarpetMarkerSymbol = "circle-x-open"
	ScattercarpetMarkerSymbol29                           ScattercarpetMarkerSymbol = "29"
	ScattercarpetMarkerSymbol_square_cross                ScattercarpetMarkerSymbol = "square-cross"
	ScattercarpetMarkerSymbol129                          ScattercarpetMarkerSymbol = "129"
	ScattercarpetMarkerSymbol_square_cross_open           ScattercarpetMarkerSymbol = "square-cross-open"
	ScattercarpetMarkerSymbol30                           ScattercarpetMarkerSymbol = "30"
	ScattercarpetMarkerSymbol_square_x                    ScattercarpetMarkerSymbol = "square-x"
	ScattercarpetMarkerSymbol130                          ScattercarpetMarkerSymbol = "130"
	ScattercarpetMarkerSymbol_square_x_open               ScattercarpetMarkerSymbol = "square-x-open"
	ScattercarpetMarkerSymbol31                           ScattercarpetMarkerSymbol = "31"
	ScattercarpetMarkerSymbol_diamond_cross               ScattercarpetMarkerSymbol = "diamond-cross"
	ScattercarpetMarkerSymbol131                          ScattercarpetMarkerSymbol = "131"
	ScattercarpetMarkerSymbol_diamond_cross_open          ScattercarpetMarkerSymbol = "diamond-cross-open"
	ScattercarpetMarkerSymbol32                           ScattercarpetMarkerSymbol = "32"
	ScattercarpetMarkerSymbol_diamond_x                   ScattercarpetMarkerSymbol = "diamond-x"
	ScattercarpetMarkerSymbol132                          ScattercarpetMarkerSymbol = "132"
	ScattercarpetMarkerSymbol_diamond_x_open              ScattercarpetMarkerSymbol = "diamond-x-open"
	ScattercarpetMarkerSymbol33                           ScattercarpetMarkerSymbol = "33"
	ScattercarpetMarkerSymbol_cross_thin                  ScattercarpetMarkerSymbol = "cross-thin"
	ScattercarpetMarkerSymbol133                          ScattercarpetMarkerSymbol = "133"
	ScattercarpetMarkerSymbol_cross_thin_open             ScattercarpetMarkerSymbol = "cross-thin-open"
	ScattercarpetMarkerSymbol34                           ScattercarpetMarkerSymbol = "34"
	ScattercarpetMarkerSymbol_x_thin                      ScattercarpetMarkerSymbol = "x-thin"
	ScattercarpetMarkerSymbol134                          ScattercarpetMarkerSymbol = "134"
	ScattercarpetMarkerSymbol_x_thin_open                 ScattercarpetMarkerSymbol = "x-thin-open"
	ScattercarpetMarkerSymbol35                           ScattercarpetMarkerSymbol = "35"
	ScattercarpetMarkerSymbol_asterisk                    ScattercarpetMarkerSymbol = "asterisk"
	ScattercarpetMarkerSymbol135                          ScattercarpetMarkerSymbol = "135"
	ScattercarpetMarkerSymbol_asterisk_open               ScattercarpetMarkerSymbol = "asterisk-open"
	ScattercarpetMarkerSymbol36                           ScattercarpetMarkerSymbol = "36"
	ScattercarpetMarkerSymbol_hash                        ScattercarpetMarkerSymbol = "hash"
	ScattercarpetMarkerSymbol136                          ScattercarpetMarkerSymbol = "136"
	ScattercarpetMarkerSymbol_hash_open                   ScattercarpetMarkerSymbol = "hash-open"
	ScattercarpetMarkerSymbol236                          ScattercarpetMarkerSymbol = "236"
	ScattercarpetMarkerSymbol_hash_dot                    ScattercarpetMarkerSymbol = "hash-dot"
	ScattercarpetMarkerSymbol336                          ScattercarpetMarkerSymbol = "336"
	ScattercarpetMarkerSymbol_hash_open_dot               ScattercarpetMarkerSymbol = "hash-open-dot"
	ScattercarpetMarkerSymbol37                           ScattercarpetMarkerSymbol = "37"
	ScattercarpetMarkerSymbol_y_up                        ScattercarpetMarkerSymbol = "y-up"
	ScattercarpetMarkerSymbol137                          ScattercarpetMarkerSymbol = "137"
	ScattercarpetMarkerSymbol_y_up_open                   ScattercarpetMarkerSymbol = "y-up-open"
	ScattercarpetMarkerSymbol38                           ScattercarpetMarkerSymbol = "38"
	ScattercarpetMarkerSymbol_y_down                      ScattercarpetMarkerSymbol = "y-down"
	ScattercarpetMarkerSymbol138                          ScattercarpetMarkerSymbol = "138"
	ScattercarpetMarkerSymbol_y_down_open                 ScattercarpetMarkerSymbol = "y-down-open"
	ScattercarpetMarkerSymbol39                           ScattercarpetMarkerSymbol = "39"
	ScattercarpetMarkerSymbol_y_left                      ScattercarpetMarkerSymbol = "y-left"
	ScattercarpetMarkerSymbol139                          ScattercarpetMarkerSymbol = "139"
	ScattercarpetMarkerSymbol_y_left_open                 ScattercarpetMarkerSymbol = "y-left-open"
	ScattercarpetMarkerSymbol40                           ScattercarpetMarkerSymbol = "40"
	ScattercarpetMarkerSymbol_y_right                     ScattercarpetMarkerSymbol = "y-right"
	ScattercarpetMarkerSymbol140                          ScattercarpetMarkerSymbol = "140"
	ScattercarpetMarkerSymbol_y_right_open                ScattercarpetMarkerSymbol = "y-right-open"
	ScattercarpetMarkerSymbol41                           ScattercarpetMarkerSymbol = "41"
	ScattercarpetMarkerSymbol_line_ew                     ScattercarpetMarkerSymbol = "line-ew"
	ScattercarpetMarkerSymbol141                          ScattercarpetMarkerSymbol = "141"
	ScattercarpetMarkerSymbol_line_ew_open                ScattercarpetMarkerSymbol = "line-ew-open"
	ScattercarpetMarkerSymbol42                           ScattercarpetMarkerSymbol = "42"
	ScattercarpetMarkerSymbol_line_ns                     ScattercarpetMarkerSymbol = "line-ns"
	ScattercarpetMarkerSymbol142                          ScattercarpetMarkerSymbol = "142"
	ScattercarpetMarkerSymbol_line_ns_open                ScattercarpetMarkerSymbol = "line-ns-open"
	ScattercarpetMarkerSymbol43                           ScattercarpetMarkerSymbol = "43"
	ScattercarpetMarkerSymbol_line_ne                     ScattercarpetMarkerSymbol = "line-ne"
	ScattercarpetMarkerSymbol143                          ScattercarpetMarkerSymbol = "143"
	ScattercarpetMarkerSymbol_line_ne_open                ScattercarpetMarkerSymbol = "line-ne-open"
	ScattercarpetMarkerSymbol44                           ScattercarpetMarkerSymbol = "44"
	ScattercarpetMarkerSymbol_line_nw                     ScattercarpetMarkerSymbol = "line-nw"
	ScattercarpetMarkerSymbol144                          ScattercarpetMarkerSymbol = "144"
	ScattercarpetMarkerSymbol_line_nw_open                ScattercarpetMarkerSymbol = "line-nw-open"
)

// ScattercarpetTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScattercarpetTextposition string

const (
	ScattercarpetTextposition_topleft      ScattercarpetTextposition = "top left"
	ScattercarpetTextposition_topcenter    ScattercarpetTextposition = "top center"
	ScattercarpetTextposition_topright     ScattercarpetTextposition = "top right"
	ScattercarpetTextposition_middleleft   ScattercarpetTextposition = "middle left"
	ScattercarpetTextposition_middlecenter ScattercarpetTextposition = "middle center"
	ScattercarpetTextposition_middleright  ScattercarpetTextposition = "middle right"
	ScattercarpetTextposition_bottomleft   ScattercarpetTextposition = "bottom left"
	ScattercarpetTextposition_bottomcenter ScattercarpetTextposition = "bottom center"
	ScattercarpetTextposition_bottomright  ScattercarpetTextposition = "bottom right"
)

// ScattercarpetVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScattercarpetVisible interface{}

var (
	ScattercarpetVisible_True       ScattercarpetVisible = true
	ScattercarpetVisible_False      ScattercarpetVisible = false
	ScattercarpetVisible_legendonly ScattercarpetVisible = "legendonly"
)

// ScattergeoFill Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape.
type ScattergeoFill string

const (
	ScattergeoFill_none   ScattergeoFill = "none"
	ScattergeoFill_toself ScattergeoFill = "toself"
)

// ScattergeoHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScattergeoHoverlabelAlign string

const (
	ScattergeoHoverlabelAlign_left  ScattergeoHoverlabelAlign = "left"
	ScattergeoHoverlabelAlign_right ScattergeoHoverlabelAlign = "right"
	ScattergeoHoverlabelAlign_auto  ScattergeoHoverlabelAlign = "auto"
)

// ScattergeoLocationmode Determines the set of locations used to match entries in `locations` to regions on the map. Values *ISO-3*, *USA-states*, *country names* correspond to features on the base map and value *geojson-id* corresponds to features from a custom GeoJSON linked to the `geojson` attribute.
type ScattergeoLocationmode string

const (
	ScattergeoLocationmode_ISO_3        ScattergeoLocationmode = "ISO-3"
	ScattergeoLocationmode_USA_states   ScattergeoLocationmode = "USA-states"
	ScattergeoLocationmode_countrynames ScattergeoLocationmode = "country names"
	ScattergeoLocationmode_geojson_id   ScattergeoLocationmode = "geojson-id"
)

// ScattergeoMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScattergeoMarkerColorbarExponentformat string

const (
	ScattergeoMarkerColorbarExponentformat_none  ScattergeoMarkerColorbarExponentformat = "none"
	ScattergeoMarkerColorbarExponentformat_e     ScattergeoMarkerColorbarExponentformat = "e"
	ScattergeoMarkerColorbarExponentformat_E     ScattergeoMarkerColorbarExponentformat = "E"
	ScattergeoMarkerColorbarExponentformat_power ScattergeoMarkerColorbarExponentformat = "power"
	ScattergeoMarkerColorbarExponentformat_SI    ScattergeoMarkerColorbarExponentformat = "SI"
	ScattergeoMarkerColorbarExponentformat_B     ScattergeoMarkerColorbarExponentformat = "B"
)

// ScattergeoMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScattergeoMarkerColorbarLenmode string

const (
	ScattergeoMarkerColorbarLenmode_fraction ScattergeoMarkerColorbarLenmode = "fraction"
	ScattergeoMarkerColorbarLenmode_pixels   ScattergeoMarkerColorbarLenmode = "pixels"
)

// ScattergeoMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScattergeoMarkerColorbarShowexponent string

const (
	ScattergeoMarkerColorbarShowexponent_all   ScattergeoMarkerColorbarShowexponent = "all"
	ScattergeoMarkerColorbarShowexponent_first ScattergeoMarkerColorbarShowexponent = "first"
	ScattergeoMarkerColorbarShowexponent_last  ScattergeoMarkerColorbarShowexponent = "last"
	ScattergeoMarkerColorbarShowexponent_none  ScattergeoMarkerColorbarShowexponent = "none"
)

// ScattergeoMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScattergeoMarkerColorbarShowtickprefix string

const (
	ScattergeoMarkerColorbarShowtickprefix_all   ScattergeoMarkerColorbarShowtickprefix = "all"
	ScattergeoMarkerColorbarShowtickprefix_first ScattergeoMarkerColorbarShowtickprefix = "first"
	ScattergeoMarkerColorbarShowtickprefix_last  ScattergeoMarkerColorbarShowtickprefix = "last"
	ScattergeoMarkerColorbarShowtickprefix_none  ScattergeoMarkerColorbarShowtickprefix = "none"
)

// ScattergeoMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScattergeoMarkerColorbarShowticksuffix string

const (
	ScattergeoMarkerColorbarShowticksuffix_all   ScattergeoMarkerColorbarShowticksuffix = "all"
	ScattergeoMarkerColorbarShowticksuffix_first ScattergeoMarkerColorbarShowticksuffix = "first"
	ScattergeoMarkerColorbarShowticksuffix_last  ScattergeoMarkerColorbarShowticksuffix = "last"
	ScattergeoMarkerColorbarShowticksuffix_none  ScattergeoMarkerColorbarShowticksuffix = "none"
)

// ScattergeoMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScattergeoMarkerColorbarThicknessmode string

const (
	ScattergeoMarkerColorbarThicknessmode_fraction ScattergeoMarkerColorbarThicknessmode = "fraction"
	ScattergeoMarkerColorbarThicknessmode_pixels   ScattergeoMarkerColorbarThicknessmode = "pixels"
)

// ScattergeoMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScattergeoMarkerColorbarTickmode string

const (
	ScattergeoMarkerColorbarTickmode_auto   ScattergeoMarkerColorbarTickmode = "auto"
	ScattergeoMarkerColorbarTickmode_linear ScattergeoMarkerColorbarTickmode = "linear"
	ScattergeoMarkerColorbarTickmode_array  ScattergeoMarkerColorbarTickmode = "array"
)

// ScattergeoMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScattergeoMarkerColorbarTicks string

const (
	ScattergeoMarkerColorbarTicks_outside ScattergeoMarkerColorbarTicks = "outside"
	ScattergeoMarkerColorbarTicks_inside  ScattergeoMarkerColorbarTicks = "inside"
)

// ScattergeoMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScattergeoMarkerColorbarTitleSide string

const (
	ScattergeoMarkerColorbarTitleSide_right  ScattergeoMarkerColorbarTitleSide = "right"
	ScattergeoMarkerColorbarTitleSide_top    ScattergeoMarkerColorbarTitleSide = "top"
	ScattergeoMarkerColorbarTitleSide_bottom ScattergeoMarkerColorbarTitleSide = "bottom"
)

// ScattergeoMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScattergeoMarkerColorbarXanchor string

const (
	ScattergeoMarkerColorbarXanchor_left   ScattergeoMarkerColorbarXanchor = "left"
	ScattergeoMarkerColorbarXanchor_center ScattergeoMarkerColorbarXanchor = "center"
	ScattergeoMarkerColorbarXanchor_right  ScattergeoMarkerColorbarXanchor = "right"
)

// ScattergeoMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScattergeoMarkerColorbarYanchor string

const (
	ScattergeoMarkerColorbarYanchor_top    ScattergeoMarkerColorbarYanchor = "top"
	ScattergeoMarkerColorbarYanchor_middle ScattergeoMarkerColorbarYanchor = "middle"
	ScattergeoMarkerColorbarYanchor_bottom ScattergeoMarkerColorbarYanchor = "bottom"
)

// ScattergeoMarkerGradientType Sets the type of gradient used to fill the markers
type ScattergeoMarkerGradientType string

const (
	ScattergeoMarkerGradientType_radial     ScattergeoMarkerGradientType = "radial"
	ScattergeoMarkerGradientType_horizontal ScattergeoMarkerGradientType = "horizontal"
	ScattergeoMarkerGradientType_vertical   ScattergeoMarkerGradientType = "vertical"
	ScattergeoMarkerGradientType_none       ScattergeoMarkerGradientType = "none"
)

// ScattergeoMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScattergeoMarkerSizemode string

const (
	ScattergeoMarkerSizemode_diameter ScattergeoMarkerSizemode = "diameter"
	ScattergeoMarkerSizemode_area     ScattergeoMarkerSizemode = "area"
)

// ScattergeoMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScattergeoMarkerSymbol string

const (
	ScattergeoMarkerSymbol0                            ScattergeoMarkerSymbol = "0"
	ScattergeoMarkerSymbol_circle                      ScattergeoMarkerSymbol = "circle"
	ScattergeoMarkerSymbol100                          ScattergeoMarkerSymbol = "100"
	ScattergeoMarkerSymbol_circle_open                 ScattergeoMarkerSymbol = "circle-open"
	ScattergeoMarkerSymbol200                          ScattergeoMarkerSymbol = "200"
	ScattergeoMarkerSymbol_circle_dot                  ScattergeoMarkerSymbol = "circle-dot"
	ScattergeoMarkerSymbol300                          ScattergeoMarkerSymbol = "300"
	ScattergeoMarkerSymbol_circle_open_dot             ScattergeoMarkerSymbol = "circle-open-dot"
	ScattergeoMarkerSymbol1                            ScattergeoMarkerSymbol = "1"
	ScattergeoMarkerSymbol_square                      ScattergeoMarkerSymbol = "square"
	ScattergeoMarkerSymbol101                          ScattergeoMarkerSymbol = "101"
	ScattergeoMarkerSymbol_square_open                 ScattergeoMarkerSymbol = "square-open"
	ScattergeoMarkerSymbol201                          ScattergeoMarkerSymbol = "201"
	ScattergeoMarkerSymbol_square_dot                  ScattergeoMarkerSymbol = "square-dot"
	ScattergeoMarkerSymbol301                          ScattergeoMarkerSymbol = "301"
	ScattergeoMarkerSymbol_square_open_dot             ScattergeoMarkerSymbol = "square-open-dot"
	ScattergeoMarkerSymbol2                            ScattergeoMarkerSymbol = "2"
	ScattergeoMarkerSymbol_diamond                     ScattergeoMarkerSymbol = "diamond"
	ScattergeoMarkerSymbol102                          ScattergeoMarkerSymbol = "102"
	ScattergeoMarkerSymbol_diamond_open                ScattergeoMarkerSymbol = "diamond-open"
	ScattergeoMarkerSymbol202                          ScattergeoMarkerSymbol = "202"
	ScattergeoMarkerSymbol_diamond_dot                 ScattergeoMarkerSymbol = "diamond-dot"
	ScattergeoMarkerSymbol302                          ScattergeoMarkerSymbol = "302"
	ScattergeoMarkerSymbol_diamond_open_dot            ScattergeoMarkerSymbol = "diamond-open-dot"
	ScattergeoMarkerSymbol3                            ScattergeoMarkerSymbol = "3"
	ScattergeoMarkerSymbol_cross                       ScattergeoMarkerSymbol = "cross"
	ScattergeoMarkerSymbol103                          ScattergeoMarkerSymbol = "103"
	ScattergeoMarkerSymbol_cross_open                  ScattergeoMarkerSymbol = "cross-open"
	ScattergeoMarkerSymbol203                          ScattergeoMarkerSymbol = "203"
	ScattergeoMarkerSymbol_cross_dot                   ScattergeoMarkerSymbol = "cross-dot"
	ScattergeoMarkerSymbol303                          ScattergeoMarkerSymbol = "303"
	ScattergeoMarkerSymbol_cross_open_dot              ScattergeoMarkerSymbol = "cross-open-dot"
	ScattergeoMarkerSymbol4                            ScattergeoMarkerSymbol = "4"
	ScattergeoMarkerSymbol_x                           ScattergeoMarkerSymbol = "x"
	ScattergeoMarkerSymbol104                          ScattergeoMarkerSymbol = "104"
	ScattergeoMarkerSymbol_x_open                      ScattergeoMarkerSymbol = "x-open"
	ScattergeoMarkerSymbol204                          ScattergeoMarkerSymbol = "204"
	ScattergeoMarkerSymbol_x_dot                       ScattergeoMarkerSymbol = "x-dot"
	ScattergeoMarkerSymbol304                          ScattergeoMarkerSymbol = "304"
	ScattergeoMarkerSymbol_x_open_dot                  ScattergeoMarkerSymbol = "x-open-dot"
	ScattergeoMarkerSymbol5                            ScattergeoMarkerSymbol = "5"
	ScattergeoMarkerSymbol_triangle_up                 ScattergeoMarkerSymbol = "triangle-up"
	ScattergeoMarkerSymbol105                          ScattergeoMarkerSymbol = "105"
	ScattergeoMarkerSymbol_triangle_up_open            ScattergeoMarkerSymbol = "triangle-up-open"
	ScattergeoMarkerSymbol205                          ScattergeoMarkerSymbol = "205"
	ScattergeoMarkerSymbol_triangle_up_dot             ScattergeoMarkerSymbol = "triangle-up-dot"
	ScattergeoMarkerSymbol305                          ScattergeoMarkerSymbol = "305"
	ScattergeoMarkerSymbol_triangle_up_open_dot        ScattergeoMarkerSymbol = "triangle-up-open-dot"
	ScattergeoMarkerSymbol6                            ScattergeoMarkerSymbol = "6"
	ScattergeoMarkerSymbol_triangle_down               ScattergeoMarkerSymbol = "triangle-down"
	ScattergeoMarkerSymbol106                          ScattergeoMarkerSymbol = "106"
	ScattergeoMarkerSymbol_triangle_down_open          ScattergeoMarkerSymbol = "triangle-down-open"
	ScattergeoMarkerSymbol206                          ScattergeoMarkerSymbol = "206"
	ScattergeoMarkerSymbol_triangle_down_dot           ScattergeoMarkerSymbol = "triangle-down-dot"
	ScattergeoMarkerSymbol306                          ScattergeoMarkerSymbol = "306"
	ScattergeoMarkerSymbol_triangle_down_open_dot      ScattergeoMarkerSymbol = "triangle-down-open-dot"
	ScattergeoMarkerSymbol7                            ScattergeoMarkerSymbol = "7"
	ScattergeoMarkerSymbol_triangle_left               ScattergeoMarkerSymbol = "triangle-left"
	ScattergeoMarkerSymbol107                          ScattergeoMarkerSymbol = "107"
	ScattergeoMarkerSymbol_triangle_left_open          ScattergeoMarkerSymbol = "triangle-left-open"
	ScattergeoMarkerSymbol207                          ScattergeoMarkerSymbol = "207"
	ScattergeoMarkerSymbol_triangle_left_dot           ScattergeoMarkerSymbol = "triangle-left-dot"
	ScattergeoMarkerSymbol307                          ScattergeoMarkerSymbol = "307"
	ScattergeoMarkerSymbol_triangle_left_open_dot      ScattergeoMarkerSymbol = "triangle-left-open-dot"
	ScattergeoMarkerSymbol8                            ScattergeoMarkerSymbol = "8"
	ScattergeoMarkerSymbol_triangle_right              ScattergeoMarkerSymbol = "triangle-right"
	ScattergeoMarkerSymbol108                          ScattergeoMarkerSymbol = "108"
	ScattergeoMarkerSymbol_triangle_right_open         ScattergeoMarkerSymbol = "triangle-right-open"
	ScattergeoMarkerSymbol208                          ScattergeoMarkerSymbol = "208"
	ScattergeoMarkerSymbol_triangle_right_dot          ScattergeoMarkerSymbol = "triangle-right-dot"
	ScattergeoMarkerSymbol308                          ScattergeoMarkerSymbol = "308"
	ScattergeoMarkerSymbol_triangle_right_open_dot     ScattergeoMarkerSymbol = "triangle-right-open-dot"
	ScattergeoMarkerSymbol9                            ScattergeoMarkerSymbol = "9"
	ScattergeoMarkerSymbol_triangle_ne                 ScattergeoMarkerSymbol = "triangle-ne"
	ScattergeoMarkerSymbol109                          ScattergeoMarkerSymbol = "109"
	ScattergeoMarkerSymbol_triangle_ne_open            ScattergeoMarkerSymbol = "triangle-ne-open"
	ScattergeoMarkerSymbol209                          ScattergeoMarkerSymbol = "209"
	ScattergeoMarkerSymbol_triangle_ne_dot             ScattergeoMarkerSymbol = "triangle-ne-dot"
	ScattergeoMarkerSymbol309                          ScattergeoMarkerSymbol = "309"
	ScattergeoMarkerSymbol_triangle_ne_open_dot        ScattergeoMarkerSymbol = "triangle-ne-open-dot"
	ScattergeoMarkerSymbol10                           ScattergeoMarkerSymbol = "10"
	ScattergeoMarkerSymbol_triangle_se                 ScattergeoMarkerSymbol = "triangle-se"
	ScattergeoMarkerSymbol110                          ScattergeoMarkerSymbol = "110"
	ScattergeoMarkerSymbol_triangle_se_open            ScattergeoMarkerSymbol = "triangle-se-open"
	ScattergeoMarkerSymbol210                          ScattergeoMarkerSymbol = "210"
	ScattergeoMarkerSymbol_triangle_se_dot             ScattergeoMarkerSymbol = "triangle-se-dot"
	ScattergeoMarkerSymbol310                          ScattergeoMarkerSymbol = "310"
	ScattergeoMarkerSymbol_triangle_se_open_dot        ScattergeoMarkerSymbol = "triangle-se-open-dot"
	ScattergeoMarkerSymbol11                           ScattergeoMarkerSymbol = "11"
	ScattergeoMarkerSymbol_triangle_sw                 ScattergeoMarkerSymbol = "triangle-sw"
	ScattergeoMarkerSymbol111                          ScattergeoMarkerSymbol = "111"
	ScattergeoMarkerSymbol_triangle_sw_open            ScattergeoMarkerSymbol = "triangle-sw-open"
	ScattergeoMarkerSymbol211                          ScattergeoMarkerSymbol = "211"
	ScattergeoMarkerSymbol_triangle_sw_dot             ScattergeoMarkerSymbol = "triangle-sw-dot"
	ScattergeoMarkerSymbol311                          ScattergeoMarkerSymbol = "311"
	ScattergeoMarkerSymbol_triangle_sw_open_dot        ScattergeoMarkerSymbol = "triangle-sw-open-dot"
	ScattergeoMarkerSymbol12                           ScattergeoMarkerSymbol = "12"
	ScattergeoMarkerSymbol_triangle_nw                 ScattergeoMarkerSymbol = "triangle-nw"
	ScattergeoMarkerSymbol112                          ScattergeoMarkerSymbol = "112"
	ScattergeoMarkerSymbol_triangle_nw_open            ScattergeoMarkerSymbol = "triangle-nw-open"
	ScattergeoMarkerSymbol212                          ScattergeoMarkerSymbol = "212"
	ScattergeoMarkerSymbol_triangle_nw_dot             ScattergeoMarkerSymbol = "triangle-nw-dot"
	ScattergeoMarkerSymbol312                          ScattergeoMarkerSymbol = "312"
	ScattergeoMarkerSymbol_triangle_nw_open_dot        ScattergeoMarkerSymbol = "triangle-nw-open-dot"
	ScattergeoMarkerSymbol13                           ScattergeoMarkerSymbol = "13"
	ScattergeoMarkerSymbol_pentagon                    ScattergeoMarkerSymbol = "pentagon"
	ScattergeoMarkerSymbol113                          ScattergeoMarkerSymbol = "113"
	ScattergeoMarkerSymbol_pentagon_open               ScattergeoMarkerSymbol = "pentagon-open"
	ScattergeoMarkerSymbol213                          ScattergeoMarkerSymbol = "213"
	ScattergeoMarkerSymbol_pentagon_dot                ScattergeoMarkerSymbol = "pentagon-dot"
	ScattergeoMarkerSymbol313                          ScattergeoMarkerSymbol = "313"
	ScattergeoMarkerSymbol_pentagon_open_dot           ScattergeoMarkerSymbol = "pentagon-open-dot"
	ScattergeoMarkerSymbol14                           ScattergeoMarkerSymbol = "14"
	ScattergeoMarkerSymbol_hexagon                     ScattergeoMarkerSymbol = "hexagon"
	ScattergeoMarkerSymbol114                          ScattergeoMarkerSymbol = "114"
	ScattergeoMarkerSymbol_hexagon_open                ScattergeoMarkerSymbol = "hexagon-open"
	ScattergeoMarkerSymbol214                          ScattergeoMarkerSymbol = "214"
	ScattergeoMarkerSymbol_hexagon_dot                 ScattergeoMarkerSymbol = "hexagon-dot"
	ScattergeoMarkerSymbol314                          ScattergeoMarkerSymbol = "314"
	ScattergeoMarkerSymbol_hexagon_open_dot            ScattergeoMarkerSymbol = "hexagon-open-dot"
	ScattergeoMarkerSymbol15                           ScattergeoMarkerSymbol = "15"
	ScattergeoMarkerSymbol_hexagon2                    ScattergeoMarkerSymbol = "hexagon2"
	ScattergeoMarkerSymbol115                          ScattergeoMarkerSymbol = "115"
	ScattergeoMarkerSymbol_hexagon2_open               ScattergeoMarkerSymbol = "hexagon2-open"
	ScattergeoMarkerSymbol215                          ScattergeoMarkerSymbol = "215"
	ScattergeoMarkerSymbol_hexagon2_dot                ScattergeoMarkerSymbol = "hexagon2-dot"
	ScattergeoMarkerSymbol315                          ScattergeoMarkerSymbol = "315"
	ScattergeoMarkerSymbol_hexagon2_open_dot           ScattergeoMarkerSymbol = "hexagon2-open-dot"
	ScattergeoMarkerSymbol16                           ScattergeoMarkerSymbol = "16"
	ScattergeoMarkerSymbol_octagon                     ScattergeoMarkerSymbol = "octagon"
	ScattergeoMarkerSymbol116                          ScattergeoMarkerSymbol = "116"
	ScattergeoMarkerSymbol_octagon_open                ScattergeoMarkerSymbol = "octagon-open"
	ScattergeoMarkerSymbol216                          ScattergeoMarkerSymbol = "216"
	ScattergeoMarkerSymbol_octagon_dot                 ScattergeoMarkerSymbol = "octagon-dot"
	ScattergeoMarkerSymbol316                          ScattergeoMarkerSymbol = "316"
	ScattergeoMarkerSymbol_octagon_open_dot            ScattergeoMarkerSymbol = "octagon-open-dot"
	ScattergeoMarkerSymbol17                           ScattergeoMarkerSymbol = "17"
	ScattergeoMarkerSymbol_star                        ScattergeoMarkerSymbol = "star"
	ScattergeoMarkerSymbol117                          ScattergeoMarkerSymbol = "117"
	ScattergeoMarkerSymbol_star_open                   ScattergeoMarkerSymbol = "star-open"
	ScattergeoMarkerSymbol217                          ScattergeoMarkerSymbol = "217"
	ScattergeoMarkerSymbol_star_dot                    ScattergeoMarkerSymbol = "star-dot"
	ScattergeoMarkerSymbol317                          ScattergeoMarkerSymbol = "317"
	ScattergeoMarkerSymbol_star_open_dot               ScattergeoMarkerSymbol = "star-open-dot"
	ScattergeoMarkerSymbol18                           ScattergeoMarkerSymbol = "18"
	ScattergeoMarkerSymbol_hexagram                    ScattergeoMarkerSymbol = "hexagram"
	ScattergeoMarkerSymbol118                          ScattergeoMarkerSymbol = "118"
	ScattergeoMarkerSymbol_hexagram_open               ScattergeoMarkerSymbol = "hexagram-open"
	ScattergeoMarkerSymbol218                          ScattergeoMarkerSymbol = "218"
	ScattergeoMarkerSymbol_hexagram_dot                ScattergeoMarkerSymbol = "hexagram-dot"
	ScattergeoMarkerSymbol318                          ScattergeoMarkerSymbol = "318"
	ScattergeoMarkerSymbol_hexagram_open_dot           ScattergeoMarkerSymbol = "hexagram-open-dot"
	ScattergeoMarkerSymbol19                           ScattergeoMarkerSymbol = "19"
	ScattergeoMarkerSymbol_star_triangle_up            ScattergeoMarkerSymbol = "star-triangle-up"
	ScattergeoMarkerSymbol119                          ScattergeoMarkerSymbol = "119"
	ScattergeoMarkerSymbol_star_triangle_up_open       ScattergeoMarkerSymbol = "star-triangle-up-open"
	ScattergeoMarkerSymbol219                          ScattergeoMarkerSymbol = "219"
	ScattergeoMarkerSymbol_star_triangle_up_dot        ScattergeoMarkerSymbol = "star-triangle-up-dot"
	ScattergeoMarkerSymbol319                          ScattergeoMarkerSymbol = "319"
	ScattergeoMarkerSymbol_star_triangle_up_open_dot   ScattergeoMarkerSymbol = "star-triangle-up-open-dot"
	ScattergeoMarkerSymbol20                           ScattergeoMarkerSymbol = "20"
	ScattergeoMarkerSymbol_star_triangle_down          ScattergeoMarkerSymbol = "star-triangle-down"
	ScattergeoMarkerSymbol120                          ScattergeoMarkerSymbol = "120"
	ScattergeoMarkerSymbol_star_triangle_down_open     ScattergeoMarkerSymbol = "star-triangle-down-open"
	ScattergeoMarkerSymbol220                          ScattergeoMarkerSymbol = "220"
	ScattergeoMarkerSymbol_star_triangle_down_dot      ScattergeoMarkerSymbol = "star-triangle-down-dot"
	ScattergeoMarkerSymbol320                          ScattergeoMarkerSymbol = "320"
	ScattergeoMarkerSymbol_star_triangle_down_open_dot ScattergeoMarkerSymbol = "star-triangle-down-open-dot"
	ScattergeoMarkerSymbol21                           ScattergeoMarkerSymbol = "21"
	ScattergeoMarkerSymbol_star_square                 ScattergeoMarkerSymbol = "star-square"
	ScattergeoMarkerSymbol121                          ScattergeoMarkerSymbol = "121"
	ScattergeoMarkerSymbol_star_square_open            ScattergeoMarkerSymbol = "star-square-open"
	ScattergeoMarkerSymbol221                          ScattergeoMarkerSymbol = "221"
	ScattergeoMarkerSymbol_star_square_dot             ScattergeoMarkerSymbol = "star-square-dot"
	ScattergeoMarkerSymbol321                          ScattergeoMarkerSymbol = "321"
	ScattergeoMarkerSymbol_star_square_open_dot        ScattergeoMarkerSymbol = "star-square-open-dot"
	ScattergeoMarkerSymbol22                           ScattergeoMarkerSymbol = "22"
	ScattergeoMarkerSymbol_star_diamond                ScattergeoMarkerSymbol = "star-diamond"
	ScattergeoMarkerSymbol122                          ScattergeoMarkerSymbol = "122"
	ScattergeoMarkerSymbol_star_diamond_open           ScattergeoMarkerSymbol = "star-diamond-open"
	ScattergeoMarkerSymbol222                          ScattergeoMarkerSymbol = "222"
	ScattergeoMarkerSymbol_star_diamond_dot            ScattergeoMarkerSymbol = "star-diamond-dot"
	ScattergeoMarkerSymbol322                          ScattergeoMarkerSymbol = "322"
	ScattergeoMarkerSymbol_star_diamond_open_dot       ScattergeoMarkerSymbol = "star-diamond-open-dot"
	ScattergeoMarkerSymbol23                           ScattergeoMarkerSymbol = "23"
	ScattergeoMarkerSymbol_diamond_tall                ScattergeoMarkerSymbol = "diamond-tall"
	ScattergeoMarkerSymbol123                          ScattergeoMarkerSymbol = "123"
	ScattergeoMarkerSymbol_diamond_tall_open           ScattergeoMarkerSymbol = "diamond-tall-open"
	ScattergeoMarkerSymbol223                          ScattergeoMarkerSymbol = "223"
	ScattergeoMarkerSymbol_diamond_tall_dot            ScattergeoMarkerSymbol = "diamond-tall-dot"
	ScattergeoMarkerSymbol323                          ScattergeoMarkerSymbol = "323"
	ScattergeoMarkerSymbol_diamond_tall_open_dot       ScattergeoMarkerSymbol = "diamond-tall-open-dot"
	ScattergeoMarkerSymbol24                           ScattergeoMarkerSymbol = "24"
	ScattergeoMarkerSymbol_diamond_wide                ScattergeoMarkerSymbol = "diamond-wide"
	ScattergeoMarkerSymbol124                          ScattergeoMarkerSymbol = "124"
	ScattergeoMarkerSymbol_diamond_wide_open           ScattergeoMarkerSymbol = "diamond-wide-open"
	ScattergeoMarkerSymbol224                          ScattergeoMarkerSymbol = "224"
	ScattergeoMarkerSymbol_diamond_wide_dot            ScattergeoMarkerSymbol = "diamond-wide-dot"
	ScattergeoMarkerSymbol324                          ScattergeoMarkerSymbol = "324"
	ScattergeoMarkerSymbol_diamond_wide_open_dot       ScattergeoMarkerSymbol = "diamond-wide-open-dot"
	ScattergeoMarkerSymbol25                           ScattergeoMarkerSymbol = "25"
	ScattergeoMarkerSymbol_hourglass                   ScattergeoMarkerSymbol = "hourglass"
	ScattergeoMarkerSymbol125                          ScattergeoMarkerSymbol = "125"
	ScattergeoMarkerSymbol_hourglass_open              ScattergeoMarkerSymbol = "hourglass-open"
	ScattergeoMarkerSymbol26                           ScattergeoMarkerSymbol = "26"
	ScattergeoMarkerSymbol_bowtie                      ScattergeoMarkerSymbol = "bowtie"
	ScattergeoMarkerSymbol126                          ScattergeoMarkerSymbol = "126"
	ScattergeoMarkerSymbol_bowtie_open                 ScattergeoMarkerSymbol = "bowtie-open"
	ScattergeoMarkerSymbol27                           ScattergeoMarkerSymbol = "27"
	ScattergeoMarkerSymbol_circle_cross                ScattergeoMarkerSymbol = "circle-cross"
	ScattergeoMarkerSymbol127                          ScattergeoMarkerSymbol = "127"
	ScattergeoMarkerSymbol_circle_cross_open           ScattergeoMarkerSymbol = "circle-cross-open"
	ScattergeoMarkerSymbol28                           ScattergeoMarkerSymbol = "28"
	ScattergeoMarkerSymbol_circle_x                    ScattergeoMarkerSymbol = "circle-x"
	ScattergeoMarkerSymbol128                          ScattergeoMarkerSymbol = "128"
	ScattergeoMarkerSymbol_circle_x_open               ScattergeoMarkerSymbol = "circle-x-open"
	ScattergeoMarkerSymbol29                           ScattergeoMarkerSymbol = "29"
	ScattergeoMarkerSymbol_square_cross                ScattergeoMarkerSymbol = "square-cross"
	ScattergeoMarkerSymbol129                          ScattergeoMarkerSymbol = "129"
	ScattergeoMarkerSymbol_square_cross_open           ScattergeoMarkerSymbol = "square-cross-open"
	ScattergeoMarkerSymbol30                           ScattergeoMarkerSymbol = "30"
	ScattergeoMarkerSymbol_square_x                    ScattergeoMarkerSymbol = "square-x"
	ScattergeoMarkerSymbol130                          ScattergeoMarkerSymbol = "130"
	ScattergeoMarkerSymbol_square_x_open               ScattergeoMarkerSymbol = "square-x-open"
	ScattergeoMarkerSymbol31                           ScattergeoMarkerSymbol = "31"
	ScattergeoMarkerSymbol_diamond_cross               ScattergeoMarkerSymbol = "diamond-cross"
	ScattergeoMarkerSymbol131                          ScattergeoMarkerSymbol = "131"
	ScattergeoMarkerSymbol_diamond_cross_open          ScattergeoMarkerSymbol = "diamond-cross-open"
	ScattergeoMarkerSymbol32                           ScattergeoMarkerSymbol = "32"
	ScattergeoMarkerSymbol_diamond_x                   ScattergeoMarkerSymbol = "diamond-x"
	ScattergeoMarkerSymbol132                          ScattergeoMarkerSymbol = "132"
	ScattergeoMarkerSymbol_diamond_x_open              ScattergeoMarkerSymbol = "diamond-x-open"
	ScattergeoMarkerSymbol33                           ScattergeoMarkerSymbol = "33"
	ScattergeoMarkerSymbol_cross_thin                  ScattergeoMarkerSymbol = "cross-thin"
	ScattergeoMarkerSymbol133                          ScattergeoMarkerSymbol = "133"
	ScattergeoMarkerSymbol_cross_thin_open             ScattergeoMarkerSymbol = "cross-thin-open"
	ScattergeoMarkerSymbol34                           ScattergeoMarkerSymbol = "34"
	ScattergeoMarkerSymbol_x_thin                      ScattergeoMarkerSymbol = "x-thin"
	ScattergeoMarkerSymbol134                          ScattergeoMarkerSymbol = "134"
	ScattergeoMarkerSymbol_x_thin_open                 ScattergeoMarkerSymbol = "x-thin-open"
	ScattergeoMarkerSymbol35                           ScattergeoMarkerSymbol = "35"
	ScattergeoMarkerSymbol_asterisk                    ScattergeoMarkerSymbol = "asterisk"
	ScattergeoMarkerSymbol135                          ScattergeoMarkerSymbol = "135"
	ScattergeoMarkerSymbol_asterisk_open               ScattergeoMarkerSymbol = "asterisk-open"
	ScattergeoMarkerSymbol36                           ScattergeoMarkerSymbol = "36"
	ScattergeoMarkerSymbol_hash                        ScattergeoMarkerSymbol = "hash"
	ScattergeoMarkerSymbol136                          ScattergeoMarkerSymbol = "136"
	ScattergeoMarkerSymbol_hash_open                   ScattergeoMarkerSymbol = "hash-open"
	ScattergeoMarkerSymbol236                          ScattergeoMarkerSymbol = "236"
	ScattergeoMarkerSymbol_hash_dot                    ScattergeoMarkerSymbol = "hash-dot"
	ScattergeoMarkerSymbol336                          ScattergeoMarkerSymbol = "336"
	ScattergeoMarkerSymbol_hash_open_dot               ScattergeoMarkerSymbol = "hash-open-dot"
	ScattergeoMarkerSymbol37                           ScattergeoMarkerSymbol = "37"
	ScattergeoMarkerSymbol_y_up                        ScattergeoMarkerSymbol = "y-up"
	ScattergeoMarkerSymbol137                          ScattergeoMarkerSymbol = "137"
	ScattergeoMarkerSymbol_y_up_open                   ScattergeoMarkerSymbol = "y-up-open"
	ScattergeoMarkerSymbol38                           ScattergeoMarkerSymbol = "38"
	ScattergeoMarkerSymbol_y_down                      ScattergeoMarkerSymbol = "y-down"
	ScattergeoMarkerSymbol138                          ScattergeoMarkerSymbol = "138"
	ScattergeoMarkerSymbol_y_down_open                 ScattergeoMarkerSymbol = "y-down-open"
	ScattergeoMarkerSymbol39                           ScattergeoMarkerSymbol = "39"
	ScattergeoMarkerSymbol_y_left                      ScattergeoMarkerSymbol = "y-left"
	ScattergeoMarkerSymbol139                          ScattergeoMarkerSymbol = "139"
	ScattergeoMarkerSymbol_y_left_open                 ScattergeoMarkerSymbol = "y-left-open"
	ScattergeoMarkerSymbol40                           ScattergeoMarkerSymbol = "40"
	ScattergeoMarkerSymbol_y_right                     ScattergeoMarkerSymbol = "y-right"
	ScattergeoMarkerSymbol140                          ScattergeoMarkerSymbol = "140"
	ScattergeoMarkerSymbol_y_right_open                ScattergeoMarkerSymbol = "y-right-open"
	ScattergeoMarkerSymbol41                           ScattergeoMarkerSymbol = "41"
	ScattergeoMarkerSymbol_line_ew                     ScattergeoMarkerSymbol = "line-ew"
	ScattergeoMarkerSymbol141                          ScattergeoMarkerSymbol = "141"
	ScattergeoMarkerSymbol_line_ew_open                ScattergeoMarkerSymbol = "line-ew-open"
	ScattergeoMarkerSymbol42                           ScattergeoMarkerSymbol = "42"
	ScattergeoMarkerSymbol_line_ns                     ScattergeoMarkerSymbol = "line-ns"
	ScattergeoMarkerSymbol142                          ScattergeoMarkerSymbol = "142"
	ScattergeoMarkerSymbol_line_ns_open                ScattergeoMarkerSymbol = "line-ns-open"
	ScattergeoMarkerSymbol43                           ScattergeoMarkerSymbol = "43"
	ScattergeoMarkerSymbol_line_ne                     ScattergeoMarkerSymbol = "line-ne"
	ScattergeoMarkerSymbol143                          ScattergeoMarkerSymbol = "143"
	ScattergeoMarkerSymbol_line_ne_open                ScattergeoMarkerSymbol = "line-ne-open"
	ScattergeoMarkerSymbol44                           ScattergeoMarkerSymbol = "44"
	ScattergeoMarkerSymbol_line_nw                     ScattergeoMarkerSymbol = "line-nw"
	ScattergeoMarkerSymbol144                          ScattergeoMarkerSymbol = "144"
	ScattergeoMarkerSymbol_line_nw_open                ScattergeoMarkerSymbol = "line-nw-open"
)

// ScattergeoTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScattergeoTextposition string

const (
	ScattergeoTextposition_topleft      ScattergeoTextposition = "top left"
	ScattergeoTextposition_topcenter    ScattergeoTextposition = "top center"
	ScattergeoTextposition_topright     ScattergeoTextposition = "top right"
	ScattergeoTextposition_middleleft   ScattergeoTextposition = "middle left"
	ScattergeoTextposition_middlecenter ScattergeoTextposition = "middle center"
	ScattergeoTextposition_middleright  ScattergeoTextposition = "middle right"
	ScattergeoTextposition_bottomleft   ScattergeoTextposition = "bottom left"
	ScattergeoTextposition_bottomcenter ScattergeoTextposition = "bottom center"
	ScattergeoTextposition_bottomright  ScattergeoTextposition = "bottom right"
)

// ScattergeoVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScattergeoVisible interface{}

var (
	ScattergeoVisible_True       ScattergeoVisible = true
	ScattergeoVisible_False      ScattergeoVisible = false
	ScattergeoVisible_legendonly ScattergeoVisible = "legendonly"
)

// ScatterglErrorXType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type ScatterglErrorXType string

const (
	ScatterglErrorXType_percent  ScatterglErrorXType = "percent"
	ScatterglErrorXType_constant ScatterglErrorXType = "constant"
	ScatterglErrorXType_sqrt     ScatterglErrorXType = "sqrt"
	ScatterglErrorXType_data     ScatterglErrorXType = "data"
)

// ScatterglErrorYType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type ScatterglErrorYType string

const (
	ScatterglErrorYType_percent  ScatterglErrorYType = "percent"
	ScatterglErrorYType_constant ScatterglErrorYType = "constant"
	ScatterglErrorYType_sqrt     ScatterglErrorYType = "sqrt"
	ScatterglErrorYType_data     ScatterglErrorYType = "data"
)

// ScatterglFill Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
type ScatterglFill string

const (
	ScatterglFill_none    ScatterglFill = "none"
	ScatterglFill_tozeroy ScatterglFill = "tozeroy"
	ScatterglFill_tozerox ScatterglFill = "tozerox"
	ScatterglFill_tonexty ScatterglFill = "tonexty"
	ScatterglFill_tonextx ScatterglFill = "tonextx"
	ScatterglFill_toself  ScatterglFill = "toself"
	ScatterglFill_tonext  ScatterglFill = "tonext"
)

// ScatterglHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterglHoverlabelAlign string

const (
	ScatterglHoverlabelAlign_left  ScatterglHoverlabelAlign = "left"
	ScatterglHoverlabelAlign_right ScatterglHoverlabelAlign = "right"
	ScatterglHoverlabelAlign_auto  ScatterglHoverlabelAlign = "auto"
)

// ScatterglLineDash Sets the style of the lines.
type ScatterglLineDash string

const (
	ScatterglLineDash_solid       ScatterglLineDash = "solid"
	ScatterglLineDash_dot         ScatterglLineDash = "dot"
	ScatterglLineDash_dash        ScatterglLineDash = "dash"
	ScatterglLineDash_longdash    ScatterglLineDash = "longdash"
	ScatterglLineDash_dashdot     ScatterglLineDash = "dashdot"
	ScatterglLineDash_longdashdot ScatterglLineDash = "longdashdot"
)

// ScatterglLineShape Determines the line shape. The values correspond to step-wise line shapes.
type ScatterglLineShape string

const (
	ScatterglLineShape_linear ScatterglLineShape = "linear"
	ScatterglLineShape_hv     ScatterglLineShape = "hv"
	ScatterglLineShape_vh     ScatterglLineShape = "vh"
	ScatterglLineShape_hvh    ScatterglLineShape = "hvh"
	ScatterglLineShape_vhv    ScatterglLineShape = "vhv"
)

// ScatterglMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterglMarkerColorbarExponentformat string

const (
	ScatterglMarkerColorbarExponentformat_none  ScatterglMarkerColorbarExponentformat = "none"
	ScatterglMarkerColorbarExponentformat_e     ScatterglMarkerColorbarExponentformat = "e"
	ScatterglMarkerColorbarExponentformat_E     ScatterglMarkerColorbarExponentformat = "E"
	ScatterglMarkerColorbarExponentformat_power ScatterglMarkerColorbarExponentformat = "power"
	ScatterglMarkerColorbarExponentformat_SI    ScatterglMarkerColorbarExponentformat = "SI"
	ScatterglMarkerColorbarExponentformat_B     ScatterglMarkerColorbarExponentformat = "B"
)

// ScatterglMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterglMarkerColorbarLenmode string

const (
	ScatterglMarkerColorbarLenmode_fraction ScatterglMarkerColorbarLenmode = "fraction"
	ScatterglMarkerColorbarLenmode_pixels   ScatterglMarkerColorbarLenmode = "pixels"
)

// ScatterglMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterglMarkerColorbarShowexponent string

const (
	ScatterglMarkerColorbarShowexponent_all   ScatterglMarkerColorbarShowexponent = "all"
	ScatterglMarkerColorbarShowexponent_first ScatterglMarkerColorbarShowexponent = "first"
	ScatterglMarkerColorbarShowexponent_last  ScatterglMarkerColorbarShowexponent = "last"
	ScatterglMarkerColorbarShowexponent_none  ScatterglMarkerColorbarShowexponent = "none"
)

// ScatterglMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterglMarkerColorbarShowtickprefix string

const (
	ScatterglMarkerColorbarShowtickprefix_all   ScatterglMarkerColorbarShowtickprefix = "all"
	ScatterglMarkerColorbarShowtickprefix_first ScatterglMarkerColorbarShowtickprefix = "first"
	ScatterglMarkerColorbarShowtickprefix_last  ScatterglMarkerColorbarShowtickprefix = "last"
	ScatterglMarkerColorbarShowtickprefix_none  ScatterglMarkerColorbarShowtickprefix = "none"
)

// ScatterglMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterglMarkerColorbarShowticksuffix string

const (
	ScatterglMarkerColorbarShowticksuffix_all   ScatterglMarkerColorbarShowticksuffix = "all"
	ScatterglMarkerColorbarShowticksuffix_first ScatterglMarkerColorbarShowticksuffix = "first"
	ScatterglMarkerColorbarShowticksuffix_last  ScatterglMarkerColorbarShowticksuffix = "last"
	ScatterglMarkerColorbarShowticksuffix_none  ScatterglMarkerColorbarShowticksuffix = "none"
)

// ScatterglMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterglMarkerColorbarThicknessmode string

const (
	ScatterglMarkerColorbarThicknessmode_fraction ScatterglMarkerColorbarThicknessmode = "fraction"
	ScatterglMarkerColorbarThicknessmode_pixels   ScatterglMarkerColorbarThicknessmode = "pixels"
)

// ScatterglMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterglMarkerColorbarTickmode string

const (
	ScatterglMarkerColorbarTickmode_auto   ScatterglMarkerColorbarTickmode = "auto"
	ScatterglMarkerColorbarTickmode_linear ScatterglMarkerColorbarTickmode = "linear"
	ScatterglMarkerColorbarTickmode_array  ScatterglMarkerColorbarTickmode = "array"
)

// ScatterglMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterglMarkerColorbarTicks string

const (
	ScatterglMarkerColorbarTicks_outside ScatterglMarkerColorbarTicks = "outside"
	ScatterglMarkerColorbarTicks_inside  ScatterglMarkerColorbarTicks = "inside"
)

// ScatterglMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterglMarkerColorbarTitleSide string

const (
	ScatterglMarkerColorbarTitleSide_right  ScatterglMarkerColorbarTitleSide = "right"
	ScatterglMarkerColorbarTitleSide_top    ScatterglMarkerColorbarTitleSide = "top"
	ScatterglMarkerColorbarTitleSide_bottom ScatterglMarkerColorbarTitleSide = "bottom"
)

// ScatterglMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterglMarkerColorbarXanchor string

const (
	ScatterglMarkerColorbarXanchor_left   ScatterglMarkerColorbarXanchor = "left"
	ScatterglMarkerColorbarXanchor_center ScatterglMarkerColorbarXanchor = "center"
	ScatterglMarkerColorbarXanchor_right  ScatterglMarkerColorbarXanchor = "right"
)

// ScatterglMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterglMarkerColorbarYanchor string

const (
	ScatterglMarkerColorbarYanchor_top    ScatterglMarkerColorbarYanchor = "top"
	ScatterglMarkerColorbarYanchor_middle ScatterglMarkerColorbarYanchor = "middle"
	ScatterglMarkerColorbarYanchor_bottom ScatterglMarkerColorbarYanchor = "bottom"
)

// ScatterglMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterglMarkerSizemode string

const (
	ScatterglMarkerSizemode_diameter ScatterglMarkerSizemode = "diameter"
	ScatterglMarkerSizemode_area     ScatterglMarkerSizemode = "area"
)

// ScatterglMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterglMarkerSymbol string

const (
	ScatterglMarkerSymbol0                            ScatterglMarkerSymbol = "0"
	ScatterglMarkerSymbol_circle                      ScatterglMarkerSymbol = "circle"
	ScatterglMarkerSymbol100                          ScatterglMarkerSymbol = "100"
	ScatterglMarkerSymbol_circle_open                 ScatterglMarkerSymbol = "circle-open"
	ScatterglMarkerSymbol200                          ScatterglMarkerSymbol = "200"
	ScatterglMarkerSymbol_circle_dot                  ScatterglMarkerSymbol = "circle-dot"
	ScatterglMarkerSymbol300                          ScatterglMarkerSymbol = "300"
	ScatterglMarkerSymbol_circle_open_dot             ScatterglMarkerSymbol = "circle-open-dot"
	ScatterglMarkerSymbol1                            ScatterglMarkerSymbol = "1"
	ScatterglMarkerSymbol_square                      ScatterglMarkerSymbol = "square"
	ScatterglMarkerSymbol101                          ScatterglMarkerSymbol = "101"
	ScatterglMarkerSymbol_square_open                 ScatterglMarkerSymbol = "square-open"
	ScatterglMarkerSymbol201                          ScatterglMarkerSymbol = "201"
	ScatterglMarkerSymbol_square_dot                  ScatterglMarkerSymbol = "square-dot"
	ScatterglMarkerSymbol301                          ScatterglMarkerSymbol = "301"
	ScatterglMarkerSymbol_square_open_dot             ScatterglMarkerSymbol = "square-open-dot"
	ScatterglMarkerSymbol2                            ScatterglMarkerSymbol = "2"
	ScatterglMarkerSymbol_diamond                     ScatterglMarkerSymbol = "diamond"
	ScatterglMarkerSymbol102                          ScatterglMarkerSymbol = "102"
	ScatterglMarkerSymbol_diamond_open                ScatterglMarkerSymbol = "diamond-open"
	ScatterglMarkerSymbol202                          ScatterglMarkerSymbol = "202"
	ScatterglMarkerSymbol_diamond_dot                 ScatterglMarkerSymbol = "diamond-dot"
	ScatterglMarkerSymbol302                          ScatterglMarkerSymbol = "302"
	ScatterglMarkerSymbol_diamond_open_dot            ScatterglMarkerSymbol = "diamond-open-dot"
	ScatterglMarkerSymbol3                            ScatterglMarkerSymbol = "3"
	ScatterglMarkerSymbol_cross                       ScatterglMarkerSymbol = "cross"
	ScatterglMarkerSymbol103                          ScatterglMarkerSymbol = "103"
	ScatterglMarkerSymbol_cross_open                  ScatterglMarkerSymbol = "cross-open"
	ScatterglMarkerSymbol203                          ScatterglMarkerSymbol = "203"
	ScatterglMarkerSymbol_cross_dot                   ScatterglMarkerSymbol = "cross-dot"
	ScatterglMarkerSymbol303                          ScatterglMarkerSymbol = "303"
	ScatterglMarkerSymbol_cross_open_dot              ScatterglMarkerSymbol = "cross-open-dot"
	ScatterglMarkerSymbol4                            ScatterglMarkerSymbol = "4"
	ScatterglMarkerSymbol_x                           ScatterglMarkerSymbol = "x"
	ScatterglMarkerSymbol104                          ScatterglMarkerSymbol = "104"
	ScatterglMarkerSymbol_x_open                      ScatterglMarkerSymbol = "x-open"
	ScatterglMarkerSymbol204                          ScatterglMarkerSymbol = "204"
	ScatterglMarkerSymbol_x_dot                       ScatterglMarkerSymbol = "x-dot"
	ScatterglMarkerSymbol304                          ScatterglMarkerSymbol = "304"
	ScatterglMarkerSymbol_x_open_dot                  ScatterglMarkerSymbol = "x-open-dot"
	ScatterglMarkerSymbol5                            ScatterglMarkerSymbol = "5"
	ScatterglMarkerSymbol_triangle_up                 ScatterglMarkerSymbol = "triangle-up"
	ScatterglMarkerSymbol105                          ScatterglMarkerSymbol = "105"
	ScatterglMarkerSymbol_triangle_up_open            ScatterglMarkerSymbol = "triangle-up-open"
	ScatterglMarkerSymbol205                          ScatterglMarkerSymbol = "205"
	ScatterglMarkerSymbol_triangle_up_dot             ScatterglMarkerSymbol = "triangle-up-dot"
	ScatterglMarkerSymbol305                          ScatterglMarkerSymbol = "305"
	ScatterglMarkerSymbol_triangle_up_open_dot        ScatterglMarkerSymbol = "triangle-up-open-dot"
	ScatterglMarkerSymbol6                            ScatterglMarkerSymbol = "6"
	ScatterglMarkerSymbol_triangle_down               ScatterglMarkerSymbol = "triangle-down"
	ScatterglMarkerSymbol106                          ScatterglMarkerSymbol = "106"
	ScatterglMarkerSymbol_triangle_down_open          ScatterglMarkerSymbol = "triangle-down-open"
	ScatterglMarkerSymbol206                          ScatterglMarkerSymbol = "206"
	ScatterglMarkerSymbol_triangle_down_dot           ScatterglMarkerSymbol = "triangle-down-dot"
	ScatterglMarkerSymbol306                          ScatterglMarkerSymbol = "306"
	ScatterglMarkerSymbol_triangle_down_open_dot      ScatterglMarkerSymbol = "triangle-down-open-dot"
	ScatterglMarkerSymbol7                            ScatterglMarkerSymbol = "7"
	ScatterglMarkerSymbol_triangle_left               ScatterglMarkerSymbol = "triangle-left"
	ScatterglMarkerSymbol107                          ScatterglMarkerSymbol = "107"
	ScatterglMarkerSymbol_triangle_left_open          ScatterglMarkerSymbol = "triangle-left-open"
	ScatterglMarkerSymbol207                          ScatterglMarkerSymbol = "207"
	ScatterglMarkerSymbol_triangle_left_dot           ScatterglMarkerSymbol = "triangle-left-dot"
	ScatterglMarkerSymbol307                          ScatterglMarkerSymbol = "307"
	ScatterglMarkerSymbol_triangle_left_open_dot      ScatterglMarkerSymbol = "triangle-left-open-dot"
	ScatterglMarkerSymbol8                            ScatterglMarkerSymbol = "8"
	ScatterglMarkerSymbol_triangle_right              ScatterglMarkerSymbol = "triangle-right"
	ScatterglMarkerSymbol108                          ScatterglMarkerSymbol = "108"
	ScatterglMarkerSymbol_triangle_right_open         ScatterglMarkerSymbol = "triangle-right-open"
	ScatterglMarkerSymbol208                          ScatterglMarkerSymbol = "208"
	ScatterglMarkerSymbol_triangle_right_dot          ScatterglMarkerSymbol = "triangle-right-dot"
	ScatterglMarkerSymbol308                          ScatterglMarkerSymbol = "308"
	ScatterglMarkerSymbol_triangle_right_open_dot     ScatterglMarkerSymbol = "triangle-right-open-dot"
	ScatterglMarkerSymbol9                            ScatterglMarkerSymbol = "9"
	ScatterglMarkerSymbol_triangle_ne                 ScatterglMarkerSymbol = "triangle-ne"
	ScatterglMarkerSymbol109                          ScatterglMarkerSymbol = "109"
	ScatterglMarkerSymbol_triangle_ne_open            ScatterglMarkerSymbol = "triangle-ne-open"
	ScatterglMarkerSymbol209                          ScatterglMarkerSymbol = "209"
	ScatterglMarkerSymbol_triangle_ne_dot             ScatterglMarkerSymbol = "triangle-ne-dot"
	ScatterglMarkerSymbol309                          ScatterglMarkerSymbol = "309"
	ScatterglMarkerSymbol_triangle_ne_open_dot        ScatterglMarkerSymbol = "triangle-ne-open-dot"
	ScatterglMarkerSymbol10                           ScatterglMarkerSymbol = "10"
	ScatterglMarkerSymbol_triangle_se                 ScatterglMarkerSymbol = "triangle-se"
	ScatterglMarkerSymbol110                          ScatterglMarkerSymbol = "110"
	ScatterglMarkerSymbol_triangle_se_open            ScatterglMarkerSymbol = "triangle-se-open"
	ScatterglMarkerSymbol210                          ScatterglMarkerSymbol = "210"
	ScatterglMarkerSymbol_triangle_se_dot             ScatterglMarkerSymbol = "triangle-se-dot"
	ScatterglMarkerSymbol310                          ScatterglMarkerSymbol = "310"
	ScatterglMarkerSymbol_triangle_se_open_dot        ScatterglMarkerSymbol = "triangle-se-open-dot"
	ScatterglMarkerSymbol11                           ScatterglMarkerSymbol = "11"
	ScatterglMarkerSymbol_triangle_sw                 ScatterglMarkerSymbol = "triangle-sw"
	ScatterglMarkerSymbol111                          ScatterglMarkerSymbol = "111"
	ScatterglMarkerSymbol_triangle_sw_open            ScatterglMarkerSymbol = "triangle-sw-open"
	ScatterglMarkerSymbol211                          ScatterglMarkerSymbol = "211"
	ScatterglMarkerSymbol_triangle_sw_dot             ScatterglMarkerSymbol = "triangle-sw-dot"
	ScatterglMarkerSymbol311                          ScatterglMarkerSymbol = "311"
	ScatterglMarkerSymbol_triangle_sw_open_dot        ScatterglMarkerSymbol = "triangle-sw-open-dot"
	ScatterglMarkerSymbol12                           ScatterglMarkerSymbol = "12"
	ScatterglMarkerSymbol_triangle_nw                 ScatterglMarkerSymbol = "triangle-nw"
	ScatterglMarkerSymbol112                          ScatterglMarkerSymbol = "112"
	ScatterglMarkerSymbol_triangle_nw_open            ScatterglMarkerSymbol = "triangle-nw-open"
	ScatterglMarkerSymbol212                          ScatterglMarkerSymbol = "212"
	ScatterglMarkerSymbol_triangle_nw_dot             ScatterglMarkerSymbol = "triangle-nw-dot"
	ScatterglMarkerSymbol312                          ScatterglMarkerSymbol = "312"
	ScatterglMarkerSymbol_triangle_nw_open_dot        ScatterglMarkerSymbol = "triangle-nw-open-dot"
	ScatterglMarkerSymbol13                           ScatterglMarkerSymbol = "13"
	ScatterglMarkerSymbol_pentagon                    ScatterglMarkerSymbol = "pentagon"
	ScatterglMarkerSymbol113                          ScatterglMarkerSymbol = "113"
	ScatterglMarkerSymbol_pentagon_open               ScatterglMarkerSymbol = "pentagon-open"
	ScatterglMarkerSymbol213                          ScatterglMarkerSymbol = "213"
	ScatterglMarkerSymbol_pentagon_dot                ScatterglMarkerSymbol = "pentagon-dot"
	ScatterglMarkerSymbol313                          ScatterglMarkerSymbol = "313"
	ScatterglMarkerSymbol_pentagon_open_dot           ScatterglMarkerSymbol = "pentagon-open-dot"
	ScatterglMarkerSymbol14                           ScatterglMarkerSymbol = "14"
	ScatterglMarkerSymbol_hexagon                     ScatterglMarkerSymbol = "hexagon"
	ScatterglMarkerSymbol114                          ScatterglMarkerSymbol = "114"
	ScatterglMarkerSymbol_hexagon_open                ScatterglMarkerSymbol = "hexagon-open"
	ScatterglMarkerSymbol214                          ScatterglMarkerSymbol = "214"
	ScatterglMarkerSymbol_hexagon_dot                 ScatterglMarkerSymbol = "hexagon-dot"
	ScatterglMarkerSymbol314                          ScatterglMarkerSymbol = "314"
	ScatterglMarkerSymbol_hexagon_open_dot            ScatterglMarkerSymbol = "hexagon-open-dot"
	ScatterglMarkerSymbol15                           ScatterglMarkerSymbol = "15"
	ScatterglMarkerSymbol_hexagon2                    ScatterglMarkerSymbol = "hexagon2"
	ScatterglMarkerSymbol115                          ScatterglMarkerSymbol = "115"
	ScatterglMarkerSymbol_hexagon2_open               ScatterglMarkerSymbol = "hexagon2-open"
	ScatterglMarkerSymbol215                          ScatterglMarkerSymbol = "215"
	ScatterglMarkerSymbol_hexagon2_dot                ScatterglMarkerSymbol = "hexagon2-dot"
	ScatterglMarkerSymbol315                          ScatterglMarkerSymbol = "315"
	ScatterglMarkerSymbol_hexagon2_open_dot           ScatterglMarkerSymbol = "hexagon2-open-dot"
	ScatterglMarkerSymbol16                           ScatterglMarkerSymbol = "16"
	ScatterglMarkerSymbol_octagon                     ScatterglMarkerSymbol = "octagon"
	ScatterglMarkerSymbol116                          ScatterglMarkerSymbol = "116"
	ScatterglMarkerSymbol_octagon_open                ScatterglMarkerSymbol = "octagon-open"
	ScatterglMarkerSymbol216                          ScatterglMarkerSymbol = "216"
	ScatterglMarkerSymbol_octagon_dot                 ScatterglMarkerSymbol = "octagon-dot"
	ScatterglMarkerSymbol316                          ScatterglMarkerSymbol = "316"
	ScatterglMarkerSymbol_octagon_open_dot            ScatterglMarkerSymbol = "octagon-open-dot"
	ScatterglMarkerSymbol17                           ScatterglMarkerSymbol = "17"
	ScatterglMarkerSymbol_star                        ScatterglMarkerSymbol = "star"
	ScatterglMarkerSymbol117                          ScatterglMarkerSymbol = "117"
	ScatterglMarkerSymbol_star_open                   ScatterglMarkerSymbol = "star-open"
	ScatterglMarkerSymbol217                          ScatterglMarkerSymbol = "217"
	ScatterglMarkerSymbol_star_dot                    ScatterglMarkerSymbol = "star-dot"
	ScatterglMarkerSymbol317                          ScatterglMarkerSymbol = "317"
	ScatterglMarkerSymbol_star_open_dot               ScatterglMarkerSymbol = "star-open-dot"
	ScatterglMarkerSymbol18                           ScatterglMarkerSymbol = "18"
	ScatterglMarkerSymbol_hexagram                    ScatterglMarkerSymbol = "hexagram"
	ScatterglMarkerSymbol118                          ScatterglMarkerSymbol = "118"
	ScatterglMarkerSymbol_hexagram_open               ScatterglMarkerSymbol = "hexagram-open"
	ScatterglMarkerSymbol218                          ScatterglMarkerSymbol = "218"
	ScatterglMarkerSymbol_hexagram_dot                ScatterglMarkerSymbol = "hexagram-dot"
	ScatterglMarkerSymbol318                          ScatterglMarkerSymbol = "318"
	ScatterglMarkerSymbol_hexagram_open_dot           ScatterglMarkerSymbol = "hexagram-open-dot"
	ScatterglMarkerSymbol19                           ScatterglMarkerSymbol = "19"
	ScatterglMarkerSymbol_star_triangle_up            ScatterglMarkerSymbol = "star-triangle-up"
	ScatterglMarkerSymbol119                          ScatterglMarkerSymbol = "119"
	ScatterglMarkerSymbol_star_triangle_up_open       ScatterglMarkerSymbol = "star-triangle-up-open"
	ScatterglMarkerSymbol219                          ScatterglMarkerSymbol = "219"
	ScatterglMarkerSymbol_star_triangle_up_dot        ScatterglMarkerSymbol = "star-triangle-up-dot"
	ScatterglMarkerSymbol319                          ScatterglMarkerSymbol = "319"
	ScatterglMarkerSymbol_star_triangle_up_open_dot   ScatterglMarkerSymbol = "star-triangle-up-open-dot"
	ScatterglMarkerSymbol20                           ScatterglMarkerSymbol = "20"
	ScatterglMarkerSymbol_star_triangle_down          ScatterglMarkerSymbol = "star-triangle-down"
	ScatterglMarkerSymbol120                          ScatterglMarkerSymbol = "120"
	ScatterglMarkerSymbol_star_triangle_down_open     ScatterglMarkerSymbol = "star-triangle-down-open"
	ScatterglMarkerSymbol220                          ScatterglMarkerSymbol = "220"
	ScatterglMarkerSymbol_star_triangle_down_dot      ScatterglMarkerSymbol = "star-triangle-down-dot"
	ScatterglMarkerSymbol320                          ScatterglMarkerSymbol = "320"
	ScatterglMarkerSymbol_star_triangle_down_open_dot ScatterglMarkerSymbol = "star-triangle-down-open-dot"
	ScatterglMarkerSymbol21                           ScatterglMarkerSymbol = "21"
	ScatterglMarkerSymbol_star_square                 ScatterglMarkerSymbol = "star-square"
	ScatterglMarkerSymbol121                          ScatterglMarkerSymbol = "121"
	ScatterglMarkerSymbol_star_square_open            ScatterglMarkerSymbol = "star-square-open"
	ScatterglMarkerSymbol221                          ScatterglMarkerSymbol = "221"
	ScatterglMarkerSymbol_star_square_dot             ScatterglMarkerSymbol = "star-square-dot"
	ScatterglMarkerSymbol321                          ScatterglMarkerSymbol = "321"
	ScatterglMarkerSymbol_star_square_open_dot        ScatterglMarkerSymbol = "star-square-open-dot"
	ScatterglMarkerSymbol22                           ScatterglMarkerSymbol = "22"
	ScatterglMarkerSymbol_star_diamond                ScatterglMarkerSymbol = "star-diamond"
	ScatterglMarkerSymbol122                          ScatterglMarkerSymbol = "122"
	ScatterglMarkerSymbol_star_diamond_open           ScatterglMarkerSymbol = "star-diamond-open"
	ScatterglMarkerSymbol222                          ScatterglMarkerSymbol = "222"
	ScatterglMarkerSymbol_star_diamond_dot            ScatterglMarkerSymbol = "star-diamond-dot"
	ScatterglMarkerSymbol322                          ScatterglMarkerSymbol = "322"
	ScatterglMarkerSymbol_star_diamond_open_dot       ScatterglMarkerSymbol = "star-diamond-open-dot"
	ScatterglMarkerSymbol23                           ScatterglMarkerSymbol = "23"
	ScatterglMarkerSymbol_diamond_tall                ScatterglMarkerSymbol = "diamond-tall"
	ScatterglMarkerSymbol123                          ScatterglMarkerSymbol = "123"
	ScatterglMarkerSymbol_diamond_tall_open           ScatterglMarkerSymbol = "diamond-tall-open"
	ScatterglMarkerSymbol223                          ScatterglMarkerSymbol = "223"
	ScatterglMarkerSymbol_diamond_tall_dot            ScatterglMarkerSymbol = "diamond-tall-dot"
	ScatterglMarkerSymbol323                          ScatterglMarkerSymbol = "323"
	ScatterglMarkerSymbol_diamond_tall_open_dot       ScatterglMarkerSymbol = "diamond-tall-open-dot"
	ScatterglMarkerSymbol24                           ScatterglMarkerSymbol = "24"
	ScatterglMarkerSymbol_diamond_wide                ScatterglMarkerSymbol = "diamond-wide"
	ScatterglMarkerSymbol124                          ScatterglMarkerSymbol = "124"
	ScatterglMarkerSymbol_diamond_wide_open           ScatterglMarkerSymbol = "diamond-wide-open"
	ScatterglMarkerSymbol224                          ScatterglMarkerSymbol = "224"
	ScatterglMarkerSymbol_diamond_wide_dot            ScatterglMarkerSymbol = "diamond-wide-dot"
	ScatterglMarkerSymbol324                          ScatterglMarkerSymbol = "324"
	ScatterglMarkerSymbol_diamond_wide_open_dot       ScatterglMarkerSymbol = "diamond-wide-open-dot"
	ScatterglMarkerSymbol25                           ScatterglMarkerSymbol = "25"
	ScatterglMarkerSymbol_hourglass                   ScatterglMarkerSymbol = "hourglass"
	ScatterglMarkerSymbol125                          ScatterglMarkerSymbol = "125"
	ScatterglMarkerSymbol_hourglass_open              ScatterglMarkerSymbol = "hourglass-open"
	ScatterglMarkerSymbol26                           ScatterglMarkerSymbol = "26"
	ScatterglMarkerSymbol_bowtie                      ScatterglMarkerSymbol = "bowtie"
	ScatterglMarkerSymbol126                          ScatterglMarkerSymbol = "126"
	ScatterglMarkerSymbol_bowtie_open                 ScatterglMarkerSymbol = "bowtie-open"
	ScatterglMarkerSymbol27                           ScatterglMarkerSymbol = "27"
	ScatterglMarkerSymbol_circle_cross                ScatterglMarkerSymbol = "circle-cross"
	ScatterglMarkerSymbol127                          ScatterglMarkerSymbol = "127"
	ScatterglMarkerSymbol_circle_cross_open           ScatterglMarkerSymbol = "circle-cross-open"
	ScatterglMarkerSymbol28                           ScatterglMarkerSymbol = "28"
	ScatterglMarkerSymbol_circle_x                    ScatterglMarkerSymbol = "circle-x"
	ScatterglMarkerSymbol128                          ScatterglMarkerSymbol = "128"
	ScatterglMarkerSymbol_circle_x_open               ScatterglMarkerSymbol = "circle-x-open"
	ScatterglMarkerSymbol29                           ScatterglMarkerSymbol = "29"
	ScatterglMarkerSymbol_square_cross                ScatterglMarkerSymbol = "square-cross"
	ScatterglMarkerSymbol129                          ScatterglMarkerSymbol = "129"
	ScatterglMarkerSymbol_square_cross_open           ScatterglMarkerSymbol = "square-cross-open"
	ScatterglMarkerSymbol30                           ScatterglMarkerSymbol = "30"
	ScatterglMarkerSymbol_square_x                    ScatterglMarkerSymbol = "square-x"
	ScatterglMarkerSymbol130                          ScatterglMarkerSymbol = "130"
	ScatterglMarkerSymbol_square_x_open               ScatterglMarkerSymbol = "square-x-open"
	ScatterglMarkerSymbol31                           ScatterglMarkerSymbol = "31"
	ScatterglMarkerSymbol_diamond_cross               ScatterglMarkerSymbol = "diamond-cross"
	ScatterglMarkerSymbol131                          ScatterglMarkerSymbol = "131"
	ScatterglMarkerSymbol_diamond_cross_open          ScatterglMarkerSymbol = "diamond-cross-open"
	ScatterglMarkerSymbol32                           ScatterglMarkerSymbol = "32"
	ScatterglMarkerSymbol_diamond_x                   ScatterglMarkerSymbol = "diamond-x"
	ScatterglMarkerSymbol132                          ScatterglMarkerSymbol = "132"
	ScatterglMarkerSymbol_diamond_x_open              ScatterglMarkerSymbol = "diamond-x-open"
	ScatterglMarkerSymbol33                           ScatterglMarkerSymbol = "33"
	ScatterglMarkerSymbol_cross_thin                  ScatterglMarkerSymbol = "cross-thin"
	ScatterglMarkerSymbol133                          ScatterglMarkerSymbol = "133"
	ScatterglMarkerSymbol_cross_thin_open             ScatterglMarkerSymbol = "cross-thin-open"
	ScatterglMarkerSymbol34                           ScatterglMarkerSymbol = "34"
	ScatterglMarkerSymbol_x_thin                      ScatterglMarkerSymbol = "x-thin"
	ScatterglMarkerSymbol134                          ScatterglMarkerSymbol = "134"
	ScatterglMarkerSymbol_x_thin_open                 ScatterglMarkerSymbol = "x-thin-open"
	ScatterglMarkerSymbol35                           ScatterglMarkerSymbol = "35"
	ScatterglMarkerSymbol_asterisk                    ScatterglMarkerSymbol = "asterisk"
	ScatterglMarkerSymbol135                          ScatterglMarkerSymbol = "135"
	ScatterglMarkerSymbol_asterisk_open               ScatterglMarkerSymbol = "asterisk-open"
	ScatterglMarkerSymbol36                           ScatterglMarkerSymbol = "36"
	ScatterglMarkerSymbol_hash                        ScatterglMarkerSymbol = "hash"
	ScatterglMarkerSymbol136                          ScatterglMarkerSymbol = "136"
	ScatterglMarkerSymbol_hash_open                   ScatterglMarkerSymbol = "hash-open"
	ScatterglMarkerSymbol236                          ScatterglMarkerSymbol = "236"
	ScatterglMarkerSymbol_hash_dot                    ScatterglMarkerSymbol = "hash-dot"
	ScatterglMarkerSymbol336                          ScatterglMarkerSymbol = "336"
	ScatterglMarkerSymbol_hash_open_dot               ScatterglMarkerSymbol = "hash-open-dot"
	ScatterglMarkerSymbol37                           ScatterglMarkerSymbol = "37"
	ScatterglMarkerSymbol_y_up                        ScatterglMarkerSymbol = "y-up"
	ScatterglMarkerSymbol137                          ScatterglMarkerSymbol = "137"
	ScatterglMarkerSymbol_y_up_open                   ScatterglMarkerSymbol = "y-up-open"
	ScatterglMarkerSymbol38                           ScatterglMarkerSymbol = "38"
	ScatterglMarkerSymbol_y_down                      ScatterglMarkerSymbol = "y-down"
	ScatterglMarkerSymbol138                          ScatterglMarkerSymbol = "138"
	ScatterglMarkerSymbol_y_down_open                 ScatterglMarkerSymbol = "y-down-open"
	ScatterglMarkerSymbol39                           ScatterglMarkerSymbol = "39"
	ScatterglMarkerSymbol_y_left                      ScatterglMarkerSymbol = "y-left"
	ScatterglMarkerSymbol139                          ScatterglMarkerSymbol = "139"
	ScatterglMarkerSymbol_y_left_open                 ScatterglMarkerSymbol = "y-left-open"
	ScatterglMarkerSymbol40                           ScatterglMarkerSymbol = "40"
	ScatterglMarkerSymbol_y_right                     ScatterglMarkerSymbol = "y-right"
	ScatterglMarkerSymbol140                          ScatterglMarkerSymbol = "140"
	ScatterglMarkerSymbol_y_right_open                ScatterglMarkerSymbol = "y-right-open"
	ScatterglMarkerSymbol41                           ScatterglMarkerSymbol = "41"
	ScatterglMarkerSymbol_line_ew                     ScatterglMarkerSymbol = "line-ew"
	ScatterglMarkerSymbol141                          ScatterglMarkerSymbol = "141"
	ScatterglMarkerSymbol_line_ew_open                ScatterglMarkerSymbol = "line-ew-open"
	ScatterglMarkerSymbol42                           ScatterglMarkerSymbol = "42"
	ScatterglMarkerSymbol_line_ns                     ScatterglMarkerSymbol = "line-ns"
	ScatterglMarkerSymbol142                          ScatterglMarkerSymbol = "142"
	ScatterglMarkerSymbol_line_ns_open                ScatterglMarkerSymbol = "line-ns-open"
	ScatterglMarkerSymbol43                           ScatterglMarkerSymbol = "43"
	ScatterglMarkerSymbol_line_ne                     ScatterglMarkerSymbol = "line-ne"
	ScatterglMarkerSymbol143                          ScatterglMarkerSymbol = "143"
	ScatterglMarkerSymbol_line_ne_open                ScatterglMarkerSymbol = "line-ne-open"
	ScatterglMarkerSymbol44                           ScatterglMarkerSymbol = "44"
	ScatterglMarkerSymbol_line_nw                     ScatterglMarkerSymbol = "line-nw"
	ScatterglMarkerSymbol144                          ScatterglMarkerSymbol = "144"
	ScatterglMarkerSymbol_line_nw_open                ScatterglMarkerSymbol = "line-nw-open"
)

// ScatterglTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterglTextposition string

const (
	ScatterglTextposition_topleft      ScatterglTextposition = "top left"
	ScatterglTextposition_topcenter    ScatterglTextposition = "top center"
	ScatterglTextposition_topright     ScatterglTextposition = "top right"
	ScatterglTextposition_middleleft   ScatterglTextposition = "middle left"
	ScatterglTextposition_middlecenter ScatterglTextposition = "middle center"
	ScatterglTextposition_middleright  ScatterglTextposition = "middle right"
	ScatterglTextposition_bottomleft   ScatterglTextposition = "bottom left"
	ScatterglTextposition_bottomcenter ScatterglTextposition = "bottom center"
	ScatterglTextposition_bottomright  ScatterglTextposition = "bottom right"
)

// ScatterglVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterglVisible interface{}

var (
	ScatterglVisible_True       ScatterglVisible = true
	ScatterglVisible_False      ScatterglVisible = false
	ScatterglVisible_legendonly ScatterglVisible = "legendonly"
)

// ScatterglXcalendar Sets the calendar system to use with `x` date data.
type ScatterglXcalendar string

const (
	ScatterglXcalendar_gregorian  ScatterglXcalendar = "gregorian"
	ScatterglXcalendar_chinese    ScatterglXcalendar = "chinese"
	ScatterglXcalendar_coptic     ScatterglXcalendar = "coptic"
	ScatterglXcalendar_discworld  ScatterglXcalendar = "discworld"
	ScatterglXcalendar_ethiopian  ScatterglXcalendar = "ethiopian"
	ScatterglXcalendar_hebrew     ScatterglXcalendar = "hebrew"
	ScatterglXcalendar_islamic    ScatterglXcalendar = "islamic"
	ScatterglXcalendar_julian     ScatterglXcalendar = "julian"
	ScatterglXcalendar_mayan      ScatterglXcalendar = "mayan"
	ScatterglXcalendar_nanakshahi ScatterglXcalendar = "nanakshahi"
	ScatterglXcalendar_nepali     ScatterglXcalendar = "nepali"
	ScatterglXcalendar_persian    ScatterglXcalendar = "persian"
	ScatterglXcalendar_jalali     ScatterglXcalendar = "jalali"
	ScatterglXcalendar_taiwan     ScatterglXcalendar = "taiwan"
	ScatterglXcalendar_thai       ScatterglXcalendar = "thai"
	ScatterglXcalendar_ummalqura  ScatterglXcalendar = "ummalqura"
)

// ScatterglYcalendar Sets the calendar system to use with `y` date data.
type ScatterglYcalendar string

const (
	ScatterglYcalendar_gregorian  ScatterglYcalendar = "gregorian"
	ScatterglYcalendar_chinese    ScatterglYcalendar = "chinese"
	ScatterglYcalendar_coptic     ScatterglYcalendar = "coptic"
	ScatterglYcalendar_discworld  ScatterglYcalendar = "discworld"
	ScatterglYcalendar_ethiopian  ScatterglYcalendar = "ethiopian"
	ScatterglYcalendar_hebrew     ScatterglYcalendar = "hebrew"
	ScatterglYcalendar_islamic    ScatterglYcalendar = "islamic"
	ScatterglYcalendar_julian     ScatterglYcalendar = "julian"
	ScatterglYcalendar_mayan      ScatterglYcalendar = "mayan"
	ScatterglYcalendar_nanakshahi ScatterglYcalendar = "nanakshahi"
	ScatterglYcalendar_nepali     ScatterglYcalendar = "nepali"
	ScatterglYcalendar_persian    ScatterglYcalendar = "persian"
	ScatterglYcalendar_jalali     ScatterglYcalendar = "jalali"
	ScatterglYcalendar_taiwan     ScatterglYcalendar = "taiwan"
	ScatterglYcalendar_thai       ScatterglYcalendar = "thai"
	ScatterglYcalendar_ummalqura  ScatterglYcalendar = "ummalqura"
)

// ScattermapboxFill Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape.
type ScattermapboxFill string

const (
	ScattermapboxFill_none   ScattermapboxFill = "none"
	ScattermapboxFill_toself ScattermapboxFill = "toself"
)

// ScattermapboxHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScattermapboxHoverlabelAlign string

const (
	ScattermapboxHoverlabelAlign_left  ScattermapboxHoverlabelAlign = "left"
	ScattermapboxHoverlabelAlign_right ScattermapboxHoverlabelAlign = "right"
	ScattermapboxHoverlabelAlign_auto  ScattermapboxHoverlabelAlign = "auto"
)

// ScattermapboxMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScattermapboxMarkerColorbarExponentformat string

const (
	ScattermapboxMarkerColorbarExponentformat_none  ScattermapboxMarkerColorbarExponentformat = "none"
	ScattermapboxMarkerColorbarExponentformat_e     ScattermapboxMarkerColorbarExponentformat = "e"
	ScattermapboxMarkerColorbarExponentformat_E     ScattermapboxMarkerColorbarExponentformat = "E"
	ScattermapboxMarkerColorbarExponentformat_power ScattermapboxMarkerColorbarExponentformat = "power"
	ScattermapboxMarkerColorbarExponentformat_SI    ScattermapboxMarkerColorbarExponentformat = "SI"
	ScattermapboxMarkerColorbarExponentformat_B     ScattermapboxMarkerColorbarExponentformat = "B"
)

// ScattermapboxMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScattermapboxMarkerColorbarLenmode string

const (
	ScattermapboxMarkerColorbarLenmode_fraction ScattermapboxMarkerColorbarLenmode = "fraction"
	ScattermapboxMarkerColorbarLenmode_pixels   ScattermapboxMarkerColorbarLenmode = "pixels"
)

// ScattermapboxMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScattermapboxMarkerColorbarShowexponent string

const (
	ScattermapboxMarkerColorbarShowexponent_all   ScattermapboxMarkerColorbarShowexponent = "all"
	ScattermapboxMarkerColorbarShowexponent_first ScattermapboxMarkerColorbarShowexponent = "first"
	ScattermapboxMarkerColorbarShowexponent_last  ScattermapboxMarkerColorbarShowexponent = "last"
	ScattermapboxMarkerColorbarShowexponent_none  ScattermapboxMarkerColorbarShowexponent = "none"
)

// ScattermapboxMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScattermapboxMarkerColorbarShowtickprefix string

const (
	ScattermapboxMarkerColorbarShowtickprefix_all   ScattermapboxMarkerColorbarShowtickprefix = "all"
	ScattermapboxMarkerColorbarShowtickprefix_first ScattermapboxMarkerColorbarShowtickprefix = "first"
	ScattermapboxMarkerColorbarShowtickprefix_last  ScattermapboxMarkerColorbarShowtickprefix = "last"
	ScattermapboxMarkerColorbarShowtickprefix_none  ScattermapboxMarkerColorbarShowtickprefix = "none"
)

// ScattermapboxMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScattermapboxMarkerColorbarShowticksuffix string

const (
	ScattermapboxMarkerColorbarShowticksuffix_all   ScattermapboxMarkerColorbarShowticksuffix = "all"
	ScattermapboxMarkerColorbarShowticksuffix_first ScattermapboxMarkerColorbarShowticksuffix = "first"
	ScattermapboxMarkerColorbarShowticksuffix_last  ScattermapboxMarkerColorbarShowticksuffix = "last"
	ScattermapboxMarkerColorbarShowticksuffix_none  ScattermapboxMarkerColorbarShowticksuffix = "none"
)

// ScattermapboxMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScattermapboxMarkerColorbarThicknessmode string

const (
	ScattermapboxMarkerColorbarThicknessmode_fraction ScattermapboxMarkerColorbarThicknessmode = "fraction"
	ScattermapboxMarkerColorbarThicknessmode_pixels   ScattermapboxMarkerColorbarThicknessmode = "pixels"
)

// ScattermapboxMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScattermapboxMarkerColorbarTickmode string

const (
	ScattermapboxMarkerColorbarTickmode_auto   ScattermapboxMarkerColorbarTickmode = "auto"
	ScattermapboxMarkerColorbarTickmode_linear ScattermapboxMarkerColorbarTickmode = "linear"
	ScattermapboxMarkerColorbarTickmode_array  ScattermapboxMarkerColorbarTickmode = "array"
)

// ScattermapboxMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScattermapboxMarkerColorbarTicks string

const (
	ScattermapboxMarkerColorbarTicks_outside ScattermapboxMarkerColorbarTicks = "outside"
	ScattermapboxMarkerColorbarTicks_inside  ScattermapboxMarkerColorbarTicks = "inside"
)

// ScattermapboxMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScattermapboxMarkerColorbarTitleSide string

const (
	ScattermapboxMarkerColorbarTitleSide_right  ScattermapboxMarkerColorbarTitleSide = "right"
	ScattermapboxMarkerColorbarTitleSide_top    ScattermapboxMarkerColorbarTitleSide = "top"
	ScattermapboxMarkerColorbarTitleSide_bottom ScattermapboxMarkerColorbarTitleSide = "bottom"
)

// ScattermapboxMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScattermapboxMarkerColorbarXanchor string

const (
	ScattermapboxMarkerColorbarXanchor_left   ScattermapboxMarkerColorbarXanchor = "left"
	ScattermapboxMarkerColorbarXanchor_center ScattermapboxMarkerColorbarXanchor = "center"
	ScattermapboxMarkerColorbarXanchor_right  ScattermapboxMarkerColorbarXanchor = "right"
)

// ScattermapboxMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScattermapboxMarkerColorbarYanchor string

const (
	ScattermapboxMarkerColorbarYanchor_top    ScattermapboxMarkerColorbarYanchor = "top"
	ScattermapboxMarkerColorbarYanchor_middle ScattermapboxMarkerColorbarYanchor = "middle"
	ScattermapboxMarkerColorbarYanchor_bottom ScattermapboxMarkerColorbarYanchor = "bottom"
)

// ScattermapboxMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScattermapboxMarkerSizemode string

const (
	ScattermapboxMarkerSizemode_diameter ScattermapboxMarkerSizemode = "diameter"
	ScattermapboxMarkerSizemode_area     ScattermapboxMarkerSizemode = "area"
)

// ScattermapboxTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScattermapboxTextposition string

const (
	ScattermapboxTextposition_topleft      ScattermapboxTextposition = "top left"
	ScattermapboxTextposition_topcenter    ScattermapboxTextposition = "top center"
	ScattermapboxTextposition_topright     ScattermapboxTextposition = "top right"
	ScattermapboxTextposition_middleleft   ScattermapboxTextposition = "middle left"
	ScattermapboxTextposition_middlecenter ScattermapboxTextposition = "middle center"
	ScattermapboxTextposition_middleright  ScattermapboxTextposition = "middle right"
	ScattermapboxTextposition_bottomleft   ScattermapboxTextposition = "bottom left"
	ScattermapboxTextposition_bottomcenter ScattermapboxTextposition = "bottom center"
	ScattermapboxTextposition_bottomright  ScattermapboxTextposition = "bottom right"
)

// ScattermapboxVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScattermapboxVisible interface{}

var (
	ScattermapboxVisible_True       ScattermapboxVisible = true
	ScattermapboxVisible_False      ScattermapboxVisible = false
	ScattermapboxVisible_legendonly ScattermapboxVisible = "legendonly"
)

// ScatterpolarFill Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. scatterpolar has a subset of the options available to scatter. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other.
type ScatterpolarFill string

const (
	ScatterpolarFill_none   ScatterpolarFill = "none"
	ScatterpolarFill_toself ScatterpolarFill = "toself"
	ScatterpolarFill_tonext ScatterpolarFill = "tonext"
)

// ScatterpolarHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterpolarHoverlabelAlign string

const (
	ScatterpolarHoverlabelAlign_left  ScatterpolarHoverlabelAlign = "left"
	ScatterpolarHoverlabelAlign_right ScatterpolarHoverlabelAlign = "right"
	ScatterpolarHoverlabelAlign_auto  ScatterpolarHoverlabelAlign = "auto"
)

// ScatterpolarLineShape Determines the line shape. With *spline* the lines are drawn using spline interpolation. The other available values correspond to step-wise line shapes.
type ScatterpolarLineShape string

const (
	ScatterpolarLineShape_linear ScatterpolarLineShape = "linear"
	ScatterpolarLineShape_spline ScatterpolarLineShape = "spline"
)

// ScatterpolarMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterpolarMarkerColorbarExponentformat string

const (
	ScatterpolarMarkerColorbarExponentformat_none  ScatterpolarMarkerColorbarExponentformat = "none"
	ScatterpolarMarkerColorbarExponentformat_e     ScatterpolarMarkerColorbarExponentformat = "e"
	ScatterpolarMarkerColorbarExponentformat_E     ScatterpolarMarkerColorbarExponentformat = "E"
	ScatterpolarMarkerColorbarExponentformat_power ScatterpolarMarkerColorbarExponentformat = "power"
	ScatterpolarMarkerColorbarExponentformat_SI    ScatterpolarMarkerColorbarExponentformat = "SI"
	ScatterpolarMarkerColorbarExponentformat_B     ScatterpolarMarkerColorbarExponentformat = "B"
)

// ScatterpolarMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterpolarMarkerColorbarLenmode string

const (
	ScatterpolarMarkerColorbarLenmode_fraction ScatterpolarMarkerColorbarLenmode = "fraction"
	ScatterpolarMarkerColorbarLenmode_pixels   ScatterpolarMarkerColorbarLenmode = "pixels"
)

// ScatterpolarMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterpolarMarkerColorbarShowexponent string

const (
	ScatterpolarMarkerColorbarShowexponent_all   ScatterpolarMarkerColorbarShowexponent = "all"
	ScatterpolarMarkerColorbarShowexponent_first ScatterpolarMarkerColorbarShowexponent = "first"
	ScatterpolarMarkerColorbarShowexponent_last  ScatterpolarMarkerColorbarShowexponent = "last"
	ScatterpolarMarkerColorbarShowexponent_none  ScatterpolarMarkerColorbarShowexponent = "none"
)

// ScatterpolarMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterpolarMarkerColorbarShowtickprefix string

const (
	ScatterpolarMarkerColorbarShowtickprefix_all   ScatterpolarMarkerColorbarShowtickprefix = "all"
	ScatterpolarMarkerColorbarShowtickprefix_first ScatterpolarMarkerColorbarShowtickprefix = "first"
	ScatterpolarMarkerColorbarShowtickprefix_last  ScatterpolarMarkerColorbarShowtickprefix = "last"
	ScatterpolarMarkerColorbarShowtickprefix_none  ScatterpolarMarkerColorbarShowtickprefix = "none"
)

// ScatterpolarMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterpolarMarkerColorbarShowticksuffix string

const (
	ScatterpolarMarkerColorbarShowticksuffix_all   ScatterpolarMarkerColorbarShowticksuffix = "all"
	ScatterpolarMarkerColorbarShowticksuffix_first ScatterpolarMarkerColorbarShowticksuffix = "first"
	ScatterpolarMarkerColorbarShowticksuffix_last  ScatterpolarMarkerColorbarShowticksuffix = "last"
	ScatterpolarMarkerColorbarShowticksuffix_none  ScatterpolarMarkerColorbarShowticksuffix = "none"
)

// ScatterpolarMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterpolarMarkerColorbarThicknessmode string

const (
	ScatterpolarMarkerColorbarThicknessmode_fraction ScatterpolarMarkerColorbarThicknessmode = "fraction"
	ScatterpolarMarkerColorbarThicknessmode_pixels   ScatterpolarMarkerColorbarThicknessmode = "pixels"
)

// ScatterpolarMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterpolarMarkerColorbarTickmode string

const (
	ScatterpolarMarkerColorbarTickmode_auto   ScatterpolarMarkerColorbarTickmode = "auto"
	ScatterpolarMarkerColorbarTickmode_linear ScatterpolarMarkerColorbarTickmode = "linear"
	ScatterpolarMarkerColorbarTickmode_array  ScatterpolarMarkerColorbarTickmode = "array"
)

// ScatterpolarMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterpolarMarkerColorbarTicks string

const (
	ScatterpolarMarkerColorbarTicks_outside ScatterpolarMarkerColorbarTicks = "outside"
	ScatterpolarMarkerColorbarTicks_inside  ScatterpolarMarkerColorbarTicks = "inside"
)

// ScatterpolarMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterpolarMarkerColorbarTitleSide string

const (
	ScatterpolarMarkerColorbarTitleSide_right  ScatterpolarMarkerColorbarTitleSide = "right"
	ScatterpolarMarkerColorbarTitleSide_top    ScatterpolarMarkerColorbarTitleSide = "top"
	ScatterpolarMarkerColorbarTitleSide_bottom ScatterpolarMarkerColorbarTitleSide = "bottom"
)

// ScatterpolarMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterpolarMarkerColorbarXanchor string

const (
	ScatterpolarMarkerColorbarXanchor_left   ScatterpolarMarkerColorbarXanchor = "left"
	ScatterpolarMarkerColorbarXanchor_center ScatterpolarMarkerColorbarXanchor = "center"
	ScatterpolarMarkerColorbarXanchor_right  ScatterpolarMarkerColorbarXanchor = "right"
)

// ScatterpolarMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterpolarMarkerColorbarYanchor string

const (
	ScatterpolarMarkerColorbarYanchor_top    ScatterpolarMarkerColorbarYanchor = "top"
	ScatterpolarMarkerColorbarYanchor_middle ScatterpolarMarkerColorbarYanchor = "middle"
	ScatterpolarMarkerColorbarYanchor_bottom ScatterpolarMarkerColorbarYanchor = "bottom"
)

// ScatterpolarMarkerGradientType Sets the type of gradient used to fill the markers
type ScatterpolarMarkerGradientType string

const (
	ScatterpolarMarkerGradientType_radial     ScatterpolarMarkerGradientType = "radial"
	ScatterpolarMarkerGradientType_horizontal ScatterpolarMarkerGradientType = "horizontal"
	ScatterpolarMarkerGradientType_vertical   ScatterpolarMarkerGradientType = "vertical"
	ScatterpolarMarkerGradientType_none       ScatterpolarMarkerGradientType = "none"
)

// ScatterpolarMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterpolarMarkerSizemode string

const (
	ScatterpolarMarkerSizemode_diameter ScatterpolarMarkerSizemode = "diameter"
	ScatterpolarMarkerSizemode_area     ScatterpolarMarkerSizemode = "area"
)

// ScatterpolarMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterpolarMarkerSymbol string

const (
	ScatterpolarMarkerSymbol0                            ScatterpolarMarkerSymbol = "0"
	ScatterpolarMarkerSymbol_circle                      ScatterpolarMarkerSymbol = "circle"
	ScatterpolarMarkerSymbol100                          ScatterpolarMarkerSymbol = "100"
	ScatterpolarMarkerSymbol_circle_open                 ScatterpolarMarkerSymbol = "circle-open"
	ScatterpolarMarkerSymbol200                          ScatterpolarMarkerSymbol = "200"
	ScatterpolarMarkerSymbol_circle_dot                  ScatterpolarMarkerSymbol = "circle-dot"
	ScatterpolarMarkerSymbol300                          ScatterpolarMarkerSymbol = "300"
	ScatterpolarMarkerSymbol_circle_open_dot             ScatterpolarMarkerSymbol = "circle-open-dot"
	ScatterpolarMarkerSymbol1                            ScatterpolarMarkerSymbol = "1"
	ScatterpolarMarkerSymbol_square                      ScatterpolarMarkerSymbol = "square"
	ScatterpolarMarkerSymbol101                          ScatterpolarMarkerSymbol = "101"
	ScatterpolarMarkerSymbol_square_open                 ScatterpolarMarkerSymbol = "square-open"
	ScatterpolarMarkerSymbol201                          ScatterpolarMarkerSymbol = "201"
	ScatterpolarMarkerSymbol_square_dot                  ScatterpolarMarkerSymbol = "square-dot"
	ScatterpolarMarkerSymbol301                          ScatterpolarMarkerSymbol = "301"
	ScatterpolarMarkerSymbol_square_open_dot             ScatterpolarMarkerSymbol = "square-open-dot"
	ScatterpolarMarkerSymbol2                            ScatterpolarMarkerSymbol = "2"
	ScatterpolarMarkerSymbol_diamond                     ScatterpolarMarkerSymbol = "diamond"
	ScatterpolarMarkerSymbol102                          ScatterpolarMarkerSymbol = "102"
	ScatterpolarMarkerSymbol_diamond_open                ScatterpolarMarkerSymbol = "diamond-open"
	ScatterpolarMarkerSymbol202                          ScatterpolarMarkerSymbol = "202"
	ScatterpolarMarkerSymbol_diamond_dot                 ScatterpolarMarkerSymbol = "diamond-dot"
	ScatterpolarMarkerSymbol302                          ScatterpolarMarkerSymbol = "302"
	ScatterpolarMarkerSymbol_diamond_open_dot            ScatterpolarMarkerSymbol = "diamond-open-dot"
	ScatterpolarMarkerSymbol3                            ScatterpolarMarkerSymbol = "3"
	ScatterpolarMarkerSymbol_cross                       ScatterpolarMarkerSymbol = "cross"
	ScatterpolarMarkerSymbol103                          ScatterpolarMarkerSymbol = "103"
	ScatterpolarMarkerSymbol_cross_open                  ScatterpolarMarkerSymbol = "cross-open"
	ScatterpolarMarkerSymbol203                          ScatterpolarMarkerSymbol = "203"
	ScatterpolarMarkerSymbol_cross_dot                   ScatterpolarMarkerSymbol = "cross-dot"
	ScatterpolarMarkerSymbol303                          ScatterpolarMarkerSymbol = "303"
	ScatterpolarMarkerSymbol_cross_open_dot              ScatterpolarMarkerSymbol = "cross-open-dot"
	ScatterpolarMarkerSymbol4                            ScatterpolarMarkerSymbol = "4"
	ScatterpolarMarkerSymbol_x                           ScatterpolarMarkerSymbol = "x"
	ScatterpolarMarkerSymbol104                          ScatterpolarMarkerSymbol = "104"
	ScatterpolarMarkerSymbol_x_open                      ScatterpolarMarkerSymbol = "x-open"
	ScatterpolarMarkerSymbol204                          ScatterpolarMarkerSymbol = "204"
	ScatterpolarMarkerSymbol_x_dot                       ScatterpolarMarkerSymbol = "x-dot"
	ScatterpolarMarkerSymbol304                          ScatterpolarMarkerSymbol = "304"
	ScatterpolarMarkerSymbol_x_open_dot                  ScatterpolarMarkerSymbol = "x-open-dot"
	ScatterpolarMarkerSymbol5                            ScatterpolarMarkerSymbol = "5"
	ScatterpolarMarkerSymbol_triangle_up                 ScatterpolarMarkerSymbol = "triangle-up"
	ScatterpolarMarkerSymbol105                          ScatterpolarMarkerSymbol = "105"
	ScatterpolarMarkerSymbol_triangle_up_open            ScatterpolarMarkerSymbol = "triangle-up-open"
	ScatterpolarMarkerSymbol205                          ScatterpolarMarkerSymbol = "205"
	ScatterpolarMarkerSymbol_triangle_up_dot             ScatterpolarMarkerSymbol = "triangle-up-dot"
	ScatterpolarMarkerSymbol305                          ScatterpolarMarkerSymbol = "305"
	ScatterpolarMarkerSymbol_triangle_up_open_dot        ScatterpolarMarkerSymbol = "triangle-up-open-dot"
	ScatterpolarMarkerSymbol6                            ScatterpolarMarkerSymbol = "6"
	ScatterpolarMarkerSymbol_triangle_down               ScatterpolarMarkerSymbol = "triangle-down"
	ScatterpolarMarkerSymbol106                          ScatterpolarMarkerSymbol = "106"
	ScatterpolarMarkerSymbol_triangle_down_open          ScatterpolarMarkerSymbol = "triangle-down-open"
	ScatterpolarMarkerSymbol206                          ScatterpolarMarkerSymbol = "206"
	ScatterpolarMarkerSymbol_triangle_down_dot           ScatterpolarMarkerSymbol = "triangle-down-dot"
	ScatterpolarMarkerSymbol306                          ScatterpolarMarkerSymbol = "306"
	ScatterpolarMarkerSymbol_triangle_down_open_dot      ScatterpolarMarkerSymbol = "triangle-down-open-dot"
	ScatterpolarMarkerSymbol7                            ScatterpolarMarkerSymbol = "7"
	ScatterpolarMarkerSymbol_triangle_left               ScatterpolarMarkerSymbol = "triangle-left"
	ScatterpolarMarkerSymbol107                          ScatterpolarMarkerSymbol = "107"
	ScatterpolarMarkerSymbol_triangle_left_open          ScatterpolarMarkerSymbol = "triangle-left-open"
	ScatterpolarMarkerSymbol207                          ScatterpolarMarkerSymbol = "207"
	ScatterpolarMarkerSymbol_triangle_left_dot           ScatterpolarMarkerSymbol = "triangle-left-dot"
	ScatterpolarMarkerSymbol307                          ScatterpolarMarkerSymbol = "307"
	ScatterpolarMarkerSymbol_triangle_left_open_dot      ScatterpolarMarkerSymbol = "triangle-left-open-dot"
	ScatterpolarMarkerSymbol8                            ScatterpolarMarkerSymbol = "8"
	ScatterpolarMarkerSymbol_triangle_right              ScatterpolarMarkerSymbol = "triangle-right"
	ScatterpolarMarkerSymbol108                          ScatterpolarMarkerSymbol = "108"
	ScatterpolarMarkerSymbol_triangle_right_open         ScatterpolarMarkerSymbol = "triangle-right-open"
	ScatterpolarMarkerSymbol208                          ScatterpolarMarkerSymbol = "208"
	ScatterpolarMarkerSymbol_triangle_right_dot          ScatterpolarMarkerSymbol = "triangle-right-dot"
	ScatterpolarMarkerSymbol308                          ScatterpolarMarkerSymbol = "308"
	ScatterpolarMarkerSymbol_triangle_right_open_dot     ScatterpolarMarkerSymbol = "triangle-right-open-dot"
	ScatterpolarMarkerSymbol9                            ScatterpolarMarkerSymbol = "9"
	ScatterpolarMarkerSymbol_triangle_ne                 ScatterpolarMarkerSymbol = "triangle-ne"
	ScatterpolarMarkerSymbol109                          ScatterpolarMarkerSymbol = "109"
	ScatterpolarMarkerSymbol_triangle_ne_open            ScatterpolarMarkerSymbol = "triangle-ne-open"
	ScatterpolarMarkerSymbol209                          ScatterpolarMarkerSymbol = "209"
	ScatterpolarMarkerSymbol_triangle_ne_dot             ScatterpolarMarkerSymbol = "triangle-ne-dot"
	ScatterpolarMarkerSymbol309                          ScatterpolarMarkerSymbol = "309"
	ScatterpolarMarkerSymbol_triangle_ne_open_dot        ScatterpolarMarkerSymbol = "triangle-ne-open-dot"
	ScatterpolarMarkerSymbol10                           ScatterpolarMarkerSymbol = "10"
	ScatterpolarMarkerSymbol_triangle_se                 ScatterpolarMarkerSymbol = "triangle-se"
	ScatterpolarMarkerSymbol110                          ScatterpolarMarkerSymbol = "110"
	ScatterpolarMarkerSymbol_triangle_se_open            ScatterpolarMarkerSymbol = "triangle-se-open"
	ScatterpolarMarkerSymbol210                          ScatterpolarMarkerSymbol = "210"
	ScatterpolarMarkerSymbol_triangle_se_dot             ScatterpolarMarkerSymbol = "triangle-se-dot"
	ScatterpolarMarkerSymbol310                          ScatterpolarMarkerSymbol = "310"
	ScatterpolarMarkerSymbol_triangle_se_open_dot        ScatterpolarMarkerSymbol = "triangle-se-open-dot"
	ScatterpolarMarkerSymbol11                           ScatterpolarMarkerSymbol = "11"
	ScatterpolarMarkerSymbol_triangle_sw                 ScatterpolarMarkerSymbol = "triangle-sw"
	ScatterpolarMarkerSymbol111                          ScatterpolarMarkerSymbol = "111"
	ScatterpolarMarkerSymbol_triangle_sw_open            ScatterpolarMarkerSymbol = "triangle-sw-open"
	ScatterpolarMarkerSymbol211                          ScatterpolarMarkerSymbol = "211"
	ScatterpolarMarkerSymbol_triangle_sw_dot             ScatterpolarMarkerSymbol = "triangle-sw-dot"
	ScatterpolarMarkerSymbol311                          ScatterpolarMarkerSymbol = "311"
	ScatterpolarMarkerSymbol_triangle_sw_open_dot        ScatterpolarMarkerSymbol = "triangle-sw-open-dot"
	ScatterpolarMarkerSymbol12                           ScatterpolarMarkerSymbol = "12"
	ScatterpolarMarkerSymbol_triangle_nw                 ScatterpolarMarkerSymbol = "triangle-nw"
	ScatterpolarMarkerSymbol112                          ScatterpolarMarkerSymbol = "112"
	ScatterpolarMarkerSymbol_triangle_nw_open            ScatterpolarMarkerSymbol = "triangle-nw-open"
	ScatterpolarMarkerSymbol212                          ScatterpolarMarkerSymbol = "212"
	ScatterpolarMarkerSymbol_triangle_nw_dot             ScatterpolarMarkerSymbol = "triangle-nw-dot"
	ScatterpolarMarkerSymbol312                          ScatterpolarMarkerSymbol = "312"
	ScatterpolarMarkerSymbol_triangle_nw_open_dot        ScatterpolarMarkerSymbol = "triangle-nw-open-dot"
	ScatterpolarMarkerSymbol13                           ScatterpolarMarkerSymbol = "13"
	ScatterpolarMarkerSymbol_pentagon                    ScatterpolarMarkerSymbol = "pentagon"
	ScatterpolarMarkerSymbol113                          ScatterpolarMarkerSymbol = "113"
	ScatterpolarMarkerSymbol_pentagon_open               ScatterpolarMarkerSymbol = "pentagon-open"
	ScatterpolarMarkerSymbol213                          ScatterpolarMarkerSymbol = "213"
	ScatterpolarMarkerSymbol_pentagon_dot                ScatterpolarMarkerSymbol = "pentagon-dot"
	ScatterpolarMarkerSymbol313                          ScatterpolarMarkerSymbol = "313"
	ScatterpolarMarkerSymbol_pentagon_open_dot           ScatterpolarMarkerSymbol = "pentagon-open-dot"
	ScatterpolarMarkerSymbol14                           ScatterpolarMarkerSymbol = "14"
	ScatterpolarMarkerSymbol_hexagon                     ScatterpolarMarkerSymbol = "hexagon"
	ScatterpolarMarkerSymbol114                          ScatterpolarMarkerSymbol = "114"
	ScatterpolarMarkerSymbol_hexagon_open                ScatterpolarMarkerSymbol = "hexagon-open"
	ScatterpolarMarkerSymbol214                          ScatterpolarMarkerSymbol = "214"
	ScatterpolarMarkerSymbol_hexagon_dot                 ScatterpolarMarkerSymbol = "hexagon-dot"
	ScatterpolarMarkerSymbol314                          ScatterpolarMarkerSymbol = "314"
	ScatterpolarMarkerSymbol_hexagon_open_dot            ScatterpolarMarkerSymbol = "hexagon-open-dot"
	ScatterpolarMarkerSymbol15                           ScatterpolarMarkerSymbol = "15"
	ScatterpolarMarkerSymbol_hexagon2                    ScatterpolarMarkerSymbol = "hexagon2"
	ScatterpolarMarkerSymbol115                          ScatterpolarMarkerSymbol = "115"
	ScatterpolarMarkerSymbol_hexagon2_open               ScatterpolarMarkerSymbol = "hexagon2-open"
	ScatterpolarMarkerSymbol215                          ScatterpolarMarkerSymbol = "215"
	ScatterpolarMarkerSymbol_hexagon2_dot                ScatterpolarMarkerSymbol = "hexagon2-dot"
	ScatterpolarMarkerSymbol315                          ScatterpolarMarkerSymbol = "315"
	ScatterpolarMarkerSymbol_hexagon2_open_dot           ScatterpolarMarkerSymbol = "hexagon2-open-dot"
	ScatterpolarMarkerSymbol16                           ScatterpolarMarkerSymbol = "16"
	ScatterpolarMarkerSymbol_octagon                     ScatterpolarMarkerSymbol = "octagon"
	ScatterpolarMarkerSymbol116                          ScatterpolarMarkerSymbol = "116"
	ScatterpolarMarkerSymbol_octagon_open                ScatterpolarMarkerSymbol = "octagon-open"
	ScatterpolarMarkerSymbol216                          ScatterpolarMarkerSymbol = "216"
	ScatterpolarMarkerSymbol_octagon_dot                 ScatterpolarMarkerSymbol = "octagon-dot"
	ScatterpolarMarkerSymbol316                          ScatterpolarMarkerSymbol = "316"
	ScatterpolarMarkerSymbol_octagon_open_dot            ScatterpolarMarkerSymbol = "octagon-open-dot"
	ScatterpolarMarkerSymbol17                           ScatterpolarMarkerSymbol = "17"
	ScatterpolarMarkerSymbol_star                        ScatterpolarMarkerSymbol = "star"
	ScatterpolarMarkerSymbol117                          ScatterpolarMarkerSymbol = "117"
	ScatterpolarMarkerSymbol_star_open                   ScatterpolarMarkerSymbol = "star-open"
	ScatterpolarMarkerSymbol217                          ScatterpolarMarkerSymbol = "217"
	ScatterpolarMarkerSymbol_star_dot                    ScatterpolarMarkerSymbol = "star-dot"
	ScatterpolarMarkerSymbol317                          ScatterpolarMarkerSymbol = "317"
	ScatterpolarMarkerSymbol_star_open_dot               ScatterpolarMarkerSymbol = "star-open-dot"
	ScatterpolarMarkerSymbol18                           ScatterpolarMarkerSymbol = "18"
	ScatterpolarMarkerSymbol_hexagram                    ScatterpolarMarkerSymbol = "hexagram"
	ScatterpolarMarkerSymbol118                          ScatterpolarMarkerSymbol = "118"
	ScatterpolarMarkerSymbol_hexagram_open               ScatterpolarMarkerSymbol = "hexagram-open"
	ScatterpolarMarkerSymbol218                          ScatterpolarMarkerSymbol = "218"
	ScatterpolarMarkerSymbol_hexagram_dot                ScatterpolarMarkerSymbol = "hexagram-dot"
	ScatterpolarMarkerSymbol318                          ScatterpolarMarkerSymbol = "318"
	ScatterpolarMarkerSymbol_hexagram_open_dot           ScatterpolarMarkerSymbol = "hexagram-open-dot"
	ScatterpolarMarkerSymbol19                           ScatterpolarMarkerSymbol = "19"
	ScatterpolarMarkerSymbol_star_triangle_up            ScatterpolarMarkerSymbol = "star-triangle-up"
	ScatterpolarMarkerSymbol119                          ScatterpolarMarkerSymbol = "119"
	ScatterpolarMarkerSymbol_star_triangle_up_open       ScatterpolarMarkerSymbol = "star-triangle-up-open"
	ScatterpolarMarkerSymbol219                          ScatterpolarMarkerSymbol = "219"
	ScatterpolarMarkerSymbol_star_triangle_up_dot        ScatterpolarMarkerSymbol = "star-triangle-up-dot"
	ScatterpolarMarkerSymbol319                          ScatterpolarMarkerSymbol = "319"
	ScatterpolarMarkerSymbol_star_triangle_up_open_dot   ScatterpolarMarkerSymbol = "star-triangle-up-open-dot"
	ScatterpolarMarkerSymbol20                           ScatterpolarMarkerSymbol = "20"
	ScatterpolarMarkerSymbol_star_triangle_down          ScatterpolarMarkerSymbol = "star-triangle-down"
	ScatterpolarMarkerSymbol120                          ScatterpolarMarkerSymbol = "120"
	ScatterpolarMarkerSymbol_star_triangle_down_open     ScatterpolarMarkerSymbol = "star-triangle-down-open"
	ScatterpolarMarkerSymbol220                          ScatterpolarMarkerSymbol = "220"
	ScatterpolarMarkerSymbol_star_triangle_down_dot      ScatterpolarMarkerSymbol = "star-triangle-down-dot"
	ScatterpolarMarkerSymbol320                          ScatterpolarMarkerSymbol = "320"
	ScatterpolarMarkerSymbol_star_triangle_down_open_dot ScatterpolarMarkerSymbol = "star-triangle-down-open-dot"
	ScatterpolarMarkerSymbol21                           ScatterpolarMarkerSymbol = "21"
	ScatterpolarMarkerSymbol_star_square                 ScatterpolarMarkerSymbol = "star-square"
	ScatterpolarMarkerSymbol121                          ScatterpolarMarkerSymbol = "121"
	ScatterpolarMarkerSymbol_star_square_open            ScatterpolarMarkerSymbol = "star-square-open"
	ScatterpolarMarkerSymbol221                          ScatterpolarMarkerSymbol = "221"
	ScatterpolarMarkerSymbol_star_square_dot             ScatterpolarMarkerSymbol = "star-square-dot"
	ScatterpolarMarkerSymbol321                          ScatterpolarMarkerSymbol = "321"
	ScatterpolarMarkerSymbol_star_square_open_dot        ScatterpolarMarkerSymbol = "star-square-open-dot"
	ScatterpolarMarkerSymbol22                           ScatterpolarMarkerSymbol = "22"
	ScatterpolarMarkerSymbol_star_diamond                ScatterpolarMarkerSymbol = "star-diamond"
	ScatterpolarMarkerSymbol122                          ScatterpolarMarkerSymbol = "122"
	ScatterpolarMarkerSymbol_star_diamond_open           ScatterpolarMarkerSymbol = "star-diamond-open"
	ScatterpolarMarkerSymbol222                          ScatterpolarMarkerSymbol = "222"
	ScatterpolarMarkerSymbol_star_diamond_dot            ScatterpolarMarkerSymbol = "star-diamond-dot"
	ScatterpolarMarkerSymbol322                          ScatterpolarMarkerSymbol = "322"
	ScatterpolarMarkerSymbol_star_diamond_open_dot       ScatterpolarMarkerSymbol = "star-diamond-open-dot"
	ScatterpolarMarkerSymbol23                           ScatterpolarMarkerSymbol = "23"
	ScatterpolarMarkerSymbol_diamond_tall                ScatterpolarMarkerSymbol = "diamond-tall"
	ScatterpolarMarkerSymbol123                          ScatterpolarMarkerSymbol = "123"
	ScatterpolarMarkerSymbol_diamond_tall_open           ScatterpolarMarkerSymbol = "diamond-tall-open"
	ScatterpolarMarkerSymbol223                          ScatterpolarMarkerSymbol = "223"
	ScatterpolarMarkerSymbol_diamond_tall_dot            ScatterpolarMarkerSymbol = "diamond-tall-dot"
	ScatterpolarMarkerSymbol323                          ScatterpolarMarkerSymbol = "323"
	ScatterpolarMarkerSymbol_diamond_tall_open_dot       ScatterpolarMarkerSymbol = "diamond-tall-open-dot"
	ScatterpolarMarkerSymbol24                           ScatterpolarMarkerSymbol = "24"
	ScatterpolarMarkerSymbol_diamond_wide                ScatterpolarMarkerSymbol = "diamond-wide"
	ScatterpolarMarkerSymbol124                          ScatterpolarMarkerSymbol = "124"
	ScatterpolarMarkerSymbol_diamond_wide_open           ScatterpolarMarkerSymbol = "diamond-wide-open"
	ScatterpolarMarkerSymbol224                          ScatterpolarMarkerSymbol = "224"
	ScatterpolarMarkerSymbol_diamond_wide_dot            ScatterpolarMarkerSymbol = "diamond-wide-dot"
	ScatterpolarMarkerSymbol324                          ScatterpolarMarkerSymbol = "324"
	ScatterpolarMarkerSymbol_diamond_wide_open_dot       ScatterpolarMarkerSymbol = "diamond-wide-open-dot"
	ScatterpolarMarkerSymbol25                           ScatterpolarMarkerSymbol = "25"
	ScatterpolarMarkerSymbol_hourglass                   ScatterpolarMarkerSymbol = "hourglass"
	ScatterpolarMarkerSymbol125                          ScatterpolarMarkerSymbol = "125"
	ScatterpolarMarkerSymbol_hourglass_open              ScatterpolarMarkerSymbol = "hourglass-open"
	ScatterpolarMarkerSymbol26                           ScatterpolarMarkerSymbol = "26"
	ScatterpolarMarkerSymbol_bowtie                      ScatterpolarMarkerSymbol = "bowtie"
	ScatterpolarMarkerSymbol126                          ScatterpolarMarkerSymbol = "126"
	ScatterpolarMarkerSymbol_bowtie_open                 ScatterpolarMarkerSymbol = "bowtie-open"
	ScatterpolarMarkerSymbol27                           ScatterpolarMarkerSymbol = "27"
	ScatterpolarMarkerSymbol_circle_cross                ScatterpolarMarkerSymbol = "circle-cross"
	ScatterpolarMarkerSymbol127                          ScatterpolarMarkerSymbol = "127"
	ScatterpolarMarkerSymbol_circle_cross_open           ScatterpolarMarkerSymbol = "circle-cross-open"
	ScatterpolarMarkerSymbol28                           ScatterpolarMarkerSymbol = "28"
	ScatterpolarMarkerSymbol_circle_x                    ScatterpolarMarkerSymbol = "circle-x"
	ScatterpolarMarkerSymbol128                          ScatterpolarMarkerSymbol = "128"
	ScatterpolarMarkerSymbol_circle_x_open               ScatterpolarMarkerSymbol = "circle-x-open"
	ScatterpolarMarkerSymbol29                           ScatterpolarMarkerSymbol = "29"
	ScatterpolarMarkerSymbol_square_cross                ScatterpolarMarkerSymbol = "square-cross"
	ScatterpolarMarkerSymbol129                          ScatterpolarMarkerSymbol = "129"
	ScatterpolarMarkerSymbol_square_cross_open           ScatterpolarMarkerSymbol = "square-cross-open"
	ScatterpolarMarkerSymbol30                           ScatterpolarMarkerSymbol = "30"
	ScatterpolarMarkerSymbol_square_x                    ScatterpolarMarkerSymbol = "square-x"
	ScatterpolarMarkerSymbol130                          ScatterpolarMarkerSymbol = "130"
	ScatterpolarMarkerSymbol_square_x_open               ScatterpolarMarkerSymbol = "square-x-open"
	ScatterpolarMarkerSymbol31                           ScatterpolarMarkerSymbol = "31"
	ScatterpolarMarkerSymbol_diamond_cross               ScatterpolarMarkerSymbol = "diamond-cross"
	ScatterpolarMarkerSymbol131                          ScatterpolarMarkerSymbol = "131"
	ScatterpolarMarkerSymbol_diamond_cross_open          ScatterpolarMarkerSymbol = "diamond-cross-open"
	ScatterpolarMarkerSymbol32                           ScatterpolarMarkerSymbol = "32"
	ScatterpolarMarkerSymbol_diamond_x                   ScatterpolarMarkerSymbol = "diamond-x"
	ScatterpolarMarkerSymbol132                          ScatterpolarMarkerSymbol = "132"
	ScatterpolarMarkerSymbol_diamond_x_open              ScatterpolarMarkerSymbol = "diamond-x-open"
	ScatterpolarMarkerSymbol33                           ScatterpolarMarkerSymbol = "33"
	ScatterpolarMarkerSymbol_cross_thin                  ScatterpolarMarkerSymbol = "cross-thin"
	ScatterpolarMarkerSymbol133                          ScatterpolarMarkerSymbol = "133"
	ScatterpolarMarkerSymbol_cross_thin_open             ScatterpolarMarkerSymbol = "cross-thin-open"
	ScatterpolarMarkerSymbol34                           ScatterpolarMarkerSymbol = "34"
	ScatterpolarMarkerSymbol_x_thin                      ScatterpolarMarkerSymbol = "x-thin"
	ScatterpolarMarkerSymbol134                          ScatterpolarMarkerSymbol = "134"
	ScatterpolarMarkerSymbol_x_thin_open                 ScatterpolarMarkerSymbol = "x-thin-open"
	ScatterpolarMarkerSymbol35                           ScatterpolarMarkerSymbol = "35"
	ScatterpolarMarkerSymbol_asterisk                    ScatterpolarMarkerSymbol = "asterisk"
	ScatterpolarMarkerSymbol135                          ScatterpolarMarkerSymbol = "135"
	ScatterpolarMarkerSymbol_asterisk_open               ScatterpolarMarkerSymbol = "asterisk-open"
	ScatterpolarMarkerSymbol36                           ScatterpolarMarkerSymbol = "36"
	ScatterpolarMarkerSymbol_hash                        ScatterpolarMarkerSymbol = "hash"
	ScatterpolarMarkerSymbol136                          ScatterpolarMarkerSymbol = "136"
	ScatterpolarMarkerSymbol_hash_open                   ScatterpolarMarkerSymbol = "hash-open"
	ScatterpolarMarkerSymbol236                          ScatterpolarMarkerSymbol = "236"
	ScatterpolarMarkerSymbol_hash_dot                    ScatterpolarMarkerSymbol = "hash-dot"
	ScatterpolarMarkerSymbol336                          ScatterpolarMarkerSymbol = "336"
	ScatterpolarMarkerSymbol_hash_open_dot               ScatterpolarMarkerSymbol = "hash-open-dot"
	ScatterpolarMarkerSymbol37                           ScatterpolarMarkerSymbol = "37"
	ScatterpolarMarkerSymbol_y_up                        ScatterpolarMarkerSymbol = "y-up"
	ScatterpolarMarkerSymbol137                          ScatterpolarMarkerSymbol = "137"
	ScatterpolarMarkerSymbol_y_up_open                   ScatterpolarMarkerSymbol = "y-up-open"
	ScatterpolarMarkerSymbol38                           ScatterpolarMarkerSymbol = "38"
	ScatterpolarMarkerSymbol_y_down                      ScatterpolarMarkerSymbol = "y-down"
	ScatterpolarMarkerSymbol138                          ScatterpolarMarkerSymbol = "138"
	ScatterpolarMarkerSymbol_y_down_open                 ScatterpolarMarkerSymbol = "y-down-open"
	ScatterpolarMarkerSymbol39                           ScatterpolarMarkerSymbol = "39"
	ScatterpolarMarkerSymbol_y_left                      ScatterpolarMarkerSymbol = "y-left"
	ScatterpolarMarkerSymbol139                          ScatterpolarMarkerSymbol = "139"
	ScatterpolarMarkerSymbol_y_left_open                 ScatterpolarMarkerSymbol = "y-left-open"
	ScatterpolarMarkerSymbol40                           ScatterpolarMarkerSymbol = "40"
	ScatterpolarMarkerSymbol_y_right                     ScatterpolarMarkerSymbol = "y-right"
	ScatterpolarMarkerSymbol140                          ScatterpolarMarkerSymbol = "140"
	ScatterpolarMarkerSymbol_y_right_open                ScatterpolarMarkerSymbol = "y-right-open"
	ScatterpolarMarkerSymbol41                           ScatterpolarMarkerSymbol = "41"
	ScatterpolarMarkerSymbol_line_ew                     ScatterpolarMarkerSymbol = "line-ew"
	ScatterpolarMarkerSymbol141                          ScatterpolarMarkerSymbol = "141"
	ScatterpolarMarkerSymbol_line_ew_open                ScatterpolarMarkerSymbol = "line-ew-open"
	ScatterpolarMarkerSymbol42                           ScatterpolarMarkerSymbol = "42"
	ScatterpolarMarkerSymbol_line_ns                     ScatterpolarMarkerSymbol = "line-ns"
	ScatterpolarMarkerSymbol142                          ScatterpolarMarkerSymbol = "142"
	ScatterpolarMarkerSymbol_line_ns_open                ScatterpolarMarkerSymbol = "line-ns-open"
	ScatterpolarMarkerSymbol43                           ScatterpolarMarkerSymbol = "43"
	ScatterpolarMarkerSymbol_line_ne                     ScatterpolarMarkerSymbol = "line-ne"
	ScatterpolarMarkerSymbol143                          ScatterpolarMarkerSymbol = "143"
	ScatterpolarMarkerSymbol_line_ne_open                ScatterpolarMarkerSymbol = "line-ne-open"
	ScatterpolarMarkerSymbol44                           ScatterpolarMarkerSymbol = "44"
	ScatterpolarMarkerSymbol_line_nw                     ScatterpolarMarkerSymbol = "line-nw"
	ScatterpolarMarkerSymbol144                          ScatterpolarMarkerSymbol = "144"
	ScatterpolarMarkerSymbol_line_nw_open                ScatterpolarMarkerSymbol = "line-nw-open"
)

// ScatterpolarTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterpolarTextposition string

const (
	ScatterpolarTextposition_topleft      ScatterpolarTextposition = "top left"
	ScatterpolarTextposition_topcenter    ScatterpolarTextposition = "top center"
	ScatterpolarTextposition_topright     ScatterpolarTextposition = "top right"
	ScatterpolarTextposition_middleleft   ScatterpolarTextposition = "middle left"
	ScatterpolarTextposition_middlecenter ScatterpolarTextposition = "middle center"
	ScatterpolarTextposition_middleright  ScatterpolarTextposition = "middle right"
	ScatterpolarTextposition_bottomleft   ScatterpolarTextposition = "bottom left"
	ScatterpolarTextposition_bottomcenter ScatterpolarTextposition = "bottom center"
	ScatterpolarTextposition_bottomright  ScatterpolarTextposition = "bottom right"
)

// ScatterpolarThetaunit Sets the unit of input *theta* values. Has an effect only when on *linear* angular axes.
type ScatterpolarThetaunit string

const (
	ScatterpolarThetaunit_radians  ScatterpolarThetaunit = "radians"
	ScatterpolarThetaunit_degrees  ScatterpolarThetaunit = "degrees"
	ScatterpolarThetaunit_gradians ScatterpolarThetaunit = "gradians"
)

// ScatterpolarVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterpolarVisible interface{}

var (
	ScatterpolarVisible_True       ScatterpolarVisible = true
	ScatterpolarVisible_False      ScatterpolarVisible = false
	ScatterpolarVisible_legendonly ScatterpolarVisible = "legendonly"
)

// ScatterpolarglFill Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
type ScatterpolarglFill string

const (
	ScatterpolarglFill_none    ScatterpolarglFill = "none"
	ScatterpolarglFill_tozeroy ScatterpolarglFill = "tozeroy"
	ScatterpolarglFill_tozerox ScatterpolarglFill = "tozerox"
	ScatterpolarglFill_tonexty ScatterpolarglFill = "tonexty"
	ScatterpolarglFill_tonextx ScatterpolarglFill = "tonextx"
	ScatterpolarglFill_toself  ScatterpolarglFill = "toself"
	ScatterpolarglFill_tonext  ScatterpolarglFill = "tonext"
)

// ScatterpolarglHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterpolarglHoverlabelAlign string

const (
	ScatterpolarglHoverlabelAlign_left  ScatterpolarglHoverlabelAlign = "left"
	ScatterpolarglHoverlabelAlign_right ScatterpolarglHoverlabelAlign = "right"
	ScatterpolarglHoverlabelAlign_auto  ScatterpolarglHoverlabelAlign = "auto"
)

// ScatterpolarglLineDash Sets the style of the lines.
type ScatterpolarglLineDash string

const (
	ScatterpolarglLineDash_solid       ScatterpolarglLineDash = "solid"
	ScatterpolarglLineDash_dot         ScatterpolarglLineDash = "dot"
	ScatterpolarglLineDash_dash        ScatterpolarglLineDash = "dash"
	ScatterpolarglLineDash_longdash    ScatterpolarglLineDash = "longdash"
	ScatterpolarglLineDash_dashdot     ScatterpolarglLineDash = "dashdot"
	ScatterpolarglLineDash_longdashdot ScatterpolarglLineDash = "longdashdot"
)

// ScatterpolarglLineShape Determines the line shape. The values correspond to step-wise line shapes.
type ScatterpolarglLineShape string

const (
	ScatterpolarglLineShape_linear ScatterpolarglLineShape = "linear"
	ScatterpolarglLineShape_hv     ScatterpolarglLineShape = "hv"
	ScatterpolarglLineShape_vh     ScatterpolarglLineShape = "vh"
	ScatterpolarglLineShape_hvh    ScatterpolarglLineShape = "hvh"
	ScatterpolarglLineShape_vhv    ScatterpolarglLineShape = "vhv"
)

// ScatterpolarglMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterpolarglMarkerColorbarExponentformat string

const (
	ScatterpolarglMarkerColorbarExponentformat_none  ScatterpolarglMarkerColorbarExponentformat = "none"
	ScatterpolarglMarkerColorbarExponentformat_e     ScatterpolarglMarkerColorbarExponentformat = "e"
	ScatterpolarglMarkerColorbarExponentformat_E     ScatterpolarglMarkerColorbarExponentformat = "E"
	ScatterpolarglMarkerColorbarExponentformat_power ScatterpolarglMarkerColorbarExponentformat = "power"
	ScatterpolarglMarkerColorbarExponentformat_SI    ScatterpolarglMarkerColorbarExponentformat = "SI"
	ScatterpolarglMarkerColorbarExponentformat_B     ScatterpolarglMarkerColorbarExponentformat = "B"
)

// ScatterpolarglMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterpolarglMarkerColorbarLenmode string

const (
	ScatterpolarglMarkerColorbarLenmode_fraction ScatterpolarglMarkerColorbarLenmode = "fraction"
	ScatterpolarglMarkerColorbarLenmode_pixels   ScatterpolarglMarkerColorbarLenmode = "pixels"
)

// ScatterpolarglMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterpolarglMarkerColorbarShowexponent string

const (
	ScatterpolarglMarkerColorbarShowexponent_all   ScatterpolarglMarkerColorbarShowexponent = "all"
	ScatterpolarglMarkerColorbarShowexponent_first ScatterpolarglMarkerColorbarShowexponent = "first"
	ScatterpolarglMarkerColorbarShowexponent_last  ScatterpolarglMarkerColorbarShowexponent = "last"
	ScatterpolarglMarkerColorbarShowexponent_none  ScatterpolarglMarkerColorbarShowexponent = "none"
)

// ScatterpolarglMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterpolarglMarkerColorbarShowtickprefix string

const (
	ScatterpolarglMarkerColorbarShowtickprefix_all   ScatterpolarglMarkerColorbarShowtickprefix = "all"
	ScatterpolarglMarkerColorbarShowtickprefix_first ScatterpolarglMarkerColorbarShowtickprefix = "first"
	ScatterpolarglMarkerColorbarShowtickprefix_last  ScatterpolarglMarkerColorbarShowtickprefix = "last"
	ScatterpolarglMarkerColorbarShowtickprefix_none  ScatterpolarglMarkerColorbarShowtickprefix = "none"
)

// ScatterpolarglMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterpolarglMarkerColorbarShowticksuffix string

const (
	ScatterpolarglMarkerColorbarShowticksuffix_all   ScatterpolarglMarkerColorbarShowticksuffix = "all"
	ScatterpolarglMarkerColorbarShowticksuffix_first ScatterpolarglMarkerColorbarShowticksuffix = "first"
	ScatterpolarglMarkerColorbarShowticksuffix_last  ScatterpolarglMarkerColorbarShowticksuffix = "last"
	ScatterpolarglMarkerColorbarShowticksuffix_none  ScatterpolarglMarkerColorbarShowticksuffix = "none"
)

// ScatterpolarglMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterpolarglMarkerColorbarThicknessmode string

const (
	ScatterpolarglMarkerColorbarThicknessmode_fraction ScatterpolarglMarkerColorbarThicknessmode = "fraction"
	ScatterpolarglMarkerColorbarThicknessmode_pixels   ScatterpolarglMarkerColorbarThicknessmode = "pixels"
)

// ScatterpolarglMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterpolarglMarkerColorbarTickmode string

const (
	ScatterpolarglMarkerColorbarTickmode_auto   ScatterpolarglMarkerColorbarTickmode = "auto"
	ScatterpolarglMarkerColorbarTickmode_linear ScatterpolarglMarkerColorbarTickmode = "linear"
	ScatterpolarglMarkerColorbarTickmode_array  ScatterpolarglMarkerColorbarTickmode = "array"
)

// ScatterpolarglMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterpolarglMarkerColorbarTicks string

const (
	ScatterpolarglMarkerColorbarTicks_outside ScatterpolarglMarkerColorbarTicks = "outside"
	ScatterpolarglMarkerColorbarTicks_inside  ScatterpolarglMarkerColorbarTicks = "inside"
)

// ScatterpolarglMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterpolarglMarkerColorbarTitleSide string

const (
	ScatterpolarglMarkerColorbarTitleSide_right  ScatterpolarglMarkerColorbarTitleSide = "right"
	ScatterpolarglMarkerColorbarTitleSide_top    ScatterpolarglMarkerColorbarTitleSide = "top"
	ScatterpolarglMarkerColorbarTitleSide_bottom ScatterpolarglMarkerColorbarTitleSide = "bottom"
)

// ScatterpolarglMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterpolarglMarkerColorbarXanchor string

const (
	ScatterpolarglMarkerColorbarXanchor_left   ScatterpolarglMarkerColorbarXanchor = "left"
	ScatterpolarglMarkerColorbarXanchor_center ScatterpolarglMarkerColorbarXanchor = "center"
	ScatterpolarglMarkerColorbarXanchor_right  ScatterpolarglMarkerColorbarXanchor = "right"
)

// ScatterpolarglMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterpolarglMarkerColorbarYanchor string

const (
	ScatterpolarglMarkerColorbarYanchor_top    ScatterpolarglMarkerColorbarYanchor = "top"
	ScatterpolarglMarkerColorbarYanchor_middle ScatterpolarglMarkerColorbarYanchor = "middle"
	ScatterpolarglMarkerColorbarYanchor_bottom ScatterpolarglMarkerColorbarYanchor = "bottom"
)

// ScatterpolarglMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterpolarglMarkerSizemode string

const (
	ScatterpolarglMarkerSizemode_diameter ScatterpolarglMarkerSizemode = "diameter"
	ScatterpolarglMarkerSizemode_area     ScatterpolarglMarkerSizemode = "area"
)

// ScatterpolarglMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterpolarglMarkerSymbol string

const (
	ScatterpolarglMarkerSymbol0                            ScatterpolarglMarkerSymbol = "0"
	ScatterpolarglMarkerSymbol_circle                      ScatterpolarglMarkerSymbol = "circle"
	ScatterpolarglMarkerSymbol100                          ScatterpolarglMarkerSymbol = "100"
	ScatterpolarglMarkerSymbol_circle_open                 ScatterpolarglMarkerSymbol = "circle-open"
	ScatterpolarglMarkerSymbol200                          ScatterpolarglMarkerSymbol = "200"
	ScatterpolarglMarkerSymbol_circle_dot                  ScatterpolarglMarkerSymbol = "circle-dot"
	ScatterpolarglMarkerSymbol300                          ScatterpolarglMarkerSymbol = "300"
	ScatterpolarglMarkerSymbol_circle_open_dot             ScatterpolarglMarkerSymbol = "circle-open-dot"
	ScatterpolarglMarkerSymbol1                            ScatterpolarglMarkerSymbol = "1"
	ScatterpolarglMarkerSymbol_square                      ScatterpolarglMarkerSymbol = "square"
	ScatterpolarglMarkerSymbol101                          ScatterpolarglMarkerSymbol = "101"
	ScatterpolarglMarkerSymbol_square_open                 ScatterpolarglMarkerSymbol = "square-open"
	ScatterpolarglMarkerSymbol201                          ScatterpolarglMarkerSymbol = "201"
	ScatterpolarglMarkerSymbol_square_dot                  ScatterpolarglMarkerSymbol = "square-dot"
	ScatterpolarglMarkerSymbol301                          ScatterpolarglMarkerSymbol = "301"
	ScatterpolarglMarkerSymbol_square_open_dot             ScatterpolarglMarkerSymbol = "square-open-dot"
	ScatterpolarglMarkerSymbol2                            ScatterpolarglMarkerSymbol = "2"
	ScatterpolarglMarkerSymbol_diamond                     ScatterpolarglMarkerSymbol = "diamond"
	ScatterpolarglMarkerSymbol102                          ScatterpolarglMarkerSymbol = "102"
	ScatterpolarglMarkerSymbol_diamond_open                ScatterpolarglMarkerSymbol = "diamond-open"
	ScatterpolarglMarkerSymbol202                          ScatterpolarglMarkerSymbol = "202"
	ScatterpolarglMarkerSymbol_diamond_dot                 ScatterpolarglMarkerSymbol = "diamond-dot"
	ScatterpolarglMarkerSymbol302                          ScatterpolarglMarkerSymbol = "302"
	ScatterpolarglMarkerSymbol_diamond_open_dot            ScatterpolarglMarkerSymbol = "diamond-open-dot"
	ScatterpolarglMarkerSymbol3                            ScatterpolarglMarkerSymbol = "3"
	ScatterpolarglMarkerSymbol_cross                       ScatterpolarglMarkerSymbol = "cross"
	ScatterpolarglMarkerSymbol103                          ScatterpolarglMarkerSymbol = "103"
	ScatterpolarglMarkerSymbol_cross_open                  ScatterpolarglMarkerSymbol = "cross-open"
	ScatterpolarglMarkerSymbol203                          ScatterpolarglMarkerSymbol = "203"
	ScatterpolarglMarkerSymbol_cross_dot                   ScatterpolarglMarkerSymbol = "cross-dot"
	ScatterpolarglMarkerSymbol303                          ScatterpolarglMarkerSymbol = "303"
	ScatterpolarglMarkerSymbol_cross_open_dot              ScatterpolarglMarkerSymbol = "cross-open-dot"
	ScatterpolarglMarkerSymbol4                            ScatterpolarglMarkerSymbol = "4"
	ScatterpolarglMarkerSymbol_x                           ScatterpolarglMarkerSymbol = "x"
	ScatterpolarglMarkerSymbol104                          ScatterpolarglMarkerSymbol = "104"
	ScatterpolarglMarkerSymbol_x_open                      ScatterpolarglMarkerSymbol = "x-open"
	ScatterpolarglMarkerSymbol204                          ScatterpolarglMarkerSymbol = "204"
	ScatterpolarglMarkerSymbol_x_dot                       ScatterpolarglMarkerSymbol = "x-dot"
	ScatterpolarglMarkerSymbol304                          ScatterpolarglMarkerSymbol = "304"
	ScatterpolarglMarkerSymbol_x_open_dot                  ScatterpolarglMarkerSymbol = "x-open-dot"
	ScatterpolarglMarkerSymbol5                            ScatterpolarglMarkerSymbol = "5"
	ScatterpolarglMarkerSymbol_triangle_up                 ScatterpolarglMarkerSymbol = "triangle-up"
	ScatterpolarglMarkerSymbol105                          ScatterpolarglMarkerSymbol = "105"
	ScatterpolarglMarkerSymbol_triangle_up_open            ScatterpolarglMarkerSymbol = "triangle-up-open"
	ScatterpolarglMarkerSymbol205                          ScatterpolarglMarkerSymbol = "205"
	ScatterpolarglMarkerSymbol_triangle_up_dot             ScatterpolarglMarkerSymbol = "triangle-up-dot"
	ScatterpolarglMarkerSymbol305                          ScatterpolarglMarkerSymbol = "305"
	ScatterpolarglMarkerSymbol_triangle_up_open_dot        ScatterpolarglMarkerSymbol = "triangle-up-open-dot"
	ScatterpolarglMarkerSymbol6                            ScatterpolarglMarkerSymbol = "6"
	ScatterpolarglMarkerSymbol_triangle_down               ScatterpolarglMarkerSymbol = "triangle-down"
	ScatterpolarglMarkerSymbol106                          ScatterpolarglMarkerSymbol = "106"
	ScatterpolarglMarkerSymbol_triangle_down_open          ScatterpolarglMarkerSymbol = "triangle-down-open"
	ScatterpolarglMarkerSymbol206                          ScatterpolarglMarkerSymbol = "206"
	ScatterpolarglMarkerSymbol_triangle_down_dot           ScatterpolarglMarkerSymbol = "triangle-down-dot"
	ScatterpolarglMarkerSymbol306                          ScatterpolarglMarkerSymbol = "306"
	ScatterpolarglMarkerSymbol_triangle_down_open_dot      ScatterpolarglMarkerSymbol = "triangle-down-open-dot"
	ScatterpolarglMarkerSymbol7                            ScatterpolarglMarkerSymbol = "7"
	ScatterpolarglMarkerSymbol_triangle_left               ScatterpolarglMarkerSymbol = "triangle-left"
	ScatterpolarglMarkerSymbol107                          ScatterpolarglMarkerSymbol = "107"
	ScatterpolarglMarkerSymbol_triangle_left_open          ScatterpolarglMarkerSymbol = "triangle-left-open"
	ScatterpolarglMarkerSymbol207                          ScatterpolarglMarkerSymbol = "207"
	ScatterpolarglMarkerSymbol_triangle_left_dot           ScatterpolarglMarkerSymbol = "triangle-left-dot"
	ScatterpolarglMarkerSymbol307                          ScatterpolarglMarkerSymbol = "307"
	ScatterpolarglMarkerSymbol_triangle_left_open_dot      ScatterpolarglMarkerSymbol = "triangle-left-open-dot"
	ScatterpolarglMarkerSymbol8                            ScatterpolarglMarkerSymbol = "8"
	ScatterpolarglMarkerSymbol_triangle_right              ScatterpolarglMarkerSymbol = "triangle-right"
	ScatterpolarglMarkerSymbol108                          ScatterpolarglMarkerSymbol = "108"
	ScatterpolarglMarkerSymbol_triangle_right_open         ScatterpolarglMarkerSymbol = "triangle-right-open"
	ScatterpolarglMarkerSymbol208                          ScatterpolarglMarkerSymbol = "208"
	ScatterpolarglMarkerSymbol_triangle_right_dot          ScatterpolarglMarkerSymbol = "triangle-right-dot"
	ScatterpolarglMarkerSymbol308                          ScatterpolarglMarkerSymbol = "308"
	ScatterpolarglMarkerSymbol_triangle_right_open_dot     ScatterpolarglMarkerSymbol = "triangle-right-open-dot"
	ScatterpolarglMarkerSymbol9                            ScatterpolarglMarkerSymbol = "9"
	ScatterpolarglMarkerSymbol_triangle_ne                 ScatterpolarglMarkerSymbol = "triangle-ne"
	ScatterpolarglMarkerSymbol109                          ScatterpolarglMarkerSymbol = "109"
	ScatterpolarglMarkerSymbol_triangle_ne_open            ScatterpolarglMarkerSymbol = "triangle-ne-open"
	ScatterpolarglMarkerSymbol209                          ScatterpolarglMarkerSymbol = "209"
	ScatterpolarglMarkerSymbol_triangle_ne_dot             ScatterpolarglMarkerSymbol = "triangle-ne-dot"
	ScatterpolarglMarkerSymbol309                          ScatterpolarglMarkerSymbol = "309"
	ScatterpolarglMarkerSymbol_triangle_ne_open_dot        ScatterpolarglMarkerSymbol = "triangle-ne-open-dot"
	ScatterpolarglMarkerSymbol10                           ScatterpolarglMarkerSymbol = "10"
	ScatterpolarglMarkerSymbol_triangle_se                 ScatterpolarglMarkerSymbol = "triangle-se"
	ScatterpolarglMarkerSymbol110                          ScatterpolarglMarkerSymbol = "110"
	ScatterpolarglMarkerSymbol_triangle_se_open            ScatterpolarglMarkerSymbol = "triangle-se-open"
	ScatterpolarglMarkerSymbol210                          ScatterpolarglMarkerSymbol = "210"
	ScatterpolarglMarkerSymbol_triangle_se_dot             ScatterpolarglMarkerSymbol = "triangle-se-dot"
	ScatterpolarglMarkerSymbol310                          ScatterpolarglMarkerSymbol = "310"
	ScatterpolarglMarkerSymbol_triangle_se_open_dot        ScatterpolarglMarkerSymbol = "triangle-se-open-dot"
	ScatterpolarglMarkerSymbol11                           ScatterpolarglMarkerSymbol = "11"
	ScatterpolarglMarkerSymbol_triangle_sw                 ScatterpolarglMarkerSymbol = "triangle-sw"
	ScatterpolarglMarkerSymbol111                          ScatterpolarglMarkerSymbol = "111"
	ScatterpolarglMarkerSymbol_triangle_sw_open            ScatterpolarglMarkerSymbol = "triangle-sw-open"
	ScatterpolarglMarkerSymbol211                          ScatterpolarglMarkerSymbol = "211"
	ScatterpolarglMarkerSymbol_triangle_sw_dot             ScatterpolarglMarkerSymbol = "triangle-sw-dot"
	ScatterpolarglMarkerSymbol311                          ScatterpolarglMarkerSymbol = "311"
	ScatterpolarglMarkerSymbol_triangle_sw_open_dot        ScatterpolarglMarkerSymbol = "triangle-sw-open-dot"
	ScatterpolarglMarkerSymbol12                           ScatterpolarglMarkerSymbol = "12"
	ScatterpolarglMarkerSymbol_triangle_nw                 ScatterpolarglMarkerSymbol = "triangle-nw"
	ScatterpolarglMarkerSymbol112                          ScatterpolarglMarkerSymbol = "112"
	ScatterpolarglMarkerSymbol_triangle_nw_open            ScatterpolarglMarkerSymbol = "triangle-nw-open"
	ScatterpolarglMarkerSymbol212                          ScatterpolarglMarkerSymbol = "212"
	ScatterpolarglMarkerSymbol_triangle_nw_dot             ScatterpolarglMarkerSymbol = "triangle-nw-dot"
	ScatterpolarglMarkerSymbol312                          ScatterpolarglMarkerSymbol = "312"
	ScatterpolarglMarkerSymbol_triangle_nw_open_dot        ScatterpolarglMarkerSymbol = "triangle-nw-open-dot"
	ScatterpolarglMarkerSymbol13                           ScatterpolarglMarkerSymbol = "13"
	ScatterpolarglMarkerSymbol_pentagon                    ScatterpolarglMarkerSymbol = "pentagon"
	ScatterpolarglMarkerSymbol113                          ScatterpolarglMarkerSymbol = "113"
	ScatterpolarglMarkerSymbol_pentagon_open               ScatterpolarglMarkerSymbol = "pentagon-open"
	ScatterpolarglMarkerSymbol213                          ScatterpolarglMarkerSymbol = "213"
	ScatterpolarglMarkerSymbol_pentagon_dot                ScatterpolarglMarkerSymbol = "pentagon-dot"
	ScatterpolarglMarkerSymbol313                          ScatterpolarglMarkerSymbol = "313"
	ScatterpolarglMarkerSymbol_pentagon_open_dot           ScatterpolarglMarkerSymbol = "pentagon-open-dot"
	ScatterpolarglMarkerSymbol14                           ScatterpolarglMarkerSymbol = "14"
	ScatterpolarglMarkerSymbol_hexagon                     ScatterpolarglMarkerSymbol = "hexagon"
	ScatterpolarglMarkerSymbol114                          ScatterpolarglMarkerSymbol = "114"
	ScatterpolarglMarkerSymbol_hexagon_open                ScatterpolarglMarkerSymbol = "hexagon-open"
	ScatterpolarglMarkerSymbol214                          ScatterpolarglMarkerSymbol = "214"
	ScatterpolarglMarkerSymbol_hexagon_dot                 ScatterpolarglMarkerSymbol = "hexagon-dot"
	ScatterpolarglMarkerSymbol314                          ScatterpolarglMarkerSymbol = "314"
	ScatterpolarglMarkerSymbol_hexagon_open_dot            ScatterpolarglMarkerSymbol = "hexagon-open-dot"
	ScatterpolarglMarkerSymbol15                           ScatterpolarglMarkerSymbol = "15"
	ScatterpolarglMarkerSymbol_hexagon2                    ScatterpolarglMarkerSymbol = "hexagon2"
	ScatterpolarglMarkerSymbol115                          ScatterpolarglMarkerSymbol = "115"
	ScatterpolarglMarkerSymbol_hexagon2_open               ScatterpolarglMarkerSymbol = "hexagon2-open"
	ScatterpolarglMarkerSymbol215                          ScatterpolarglMarkerSymbol = "215"
	ScatterpolarglMarkerSymbol_hexagon2_dot                ScatterpolarglMarkerSymbol = "hexagon2-dot"
	ScatterpolarglMarkerSymbol315                          ScatterpolarglMarkerSymbol = "315"
	ScatterpolarglMarkerSymbol_hexagon2_open_dot           ScatterpolarglMarkerSymbol = "hexagon2-open-dot"
	ScatterpolarglMarkerSymbol16                           ScatterpolarglMarkerSymbol = "16"
	ScatterpolarglMarkerSymbol_octagon                     ScatterpolarglMarkerSymbol = "octagon"
	ScatterpolarglMarkerSymbol116                          ScatterpolarglMarkerSymbol = "116"
	ScatterpolarglMarkerSymbol_octagon_open                ScatterpolarglMarkerSymbol = "octagon-open"
	ScatterpolarglMarkerSymbol216                          ScatterpolarglMarkerSymbol = "216"
	ScatterpolarglMarkerSymbol_octagon_dot                 ScatterpolarglMarkerSymbol = "octagon-dot"
	ScatterpolarglMarkerSymbol316                          ScatterpolarglMarkerSymbol = "316"
	ScatterpolarglMarkerSymbol_octagon_open_dot            ScatterpolarglMarkerSymbol = "octagon-open-dot"
	ScatterpolarglMarkerSymbol17                           ScatterpolarglMarkerSymbol = "17"
	ScatterpolarglMarkerSymbol_star                        ScatterpolarglMarkerSymbol = "star"
	ScatterpolarglMarkerSymbol117                          ScatterpolarglMarkerSymbol = "117"
	ScatterpolarglMarkerSymbol_star_open                   ScatterpolarglMarkerSymbol = "star-open"
	ScatterpolarglMarkerSymbol217                          ScatterpolarglMarkerSymbol = "217"
	ScatterpolarglMarkerSymbol_star_dot                    ScatterpolarglMarkerSymbol = "star-dot"
	ScatterpolarglMarkerSymbol317                          ScatterpolarglMarkerSymbol = "317"
	ScatterpolarglMarkerSymbol_star_open_dot               ScatterpolarglMarkerSymbol = "star-open-dot"
	ScatterpolarglMarkerSymbol18                           ScatterpolarglMarkerSymbol = "18"
	ScatterpolarglMarkerSymbol_hexagram                    ScatterpolarglMarkerSymbol = "hexagram"
	ScatterpolarglMarkerSymbol118                          ScatterpolarglMarkerSymbol = "118"
	ScatterpolarglMarkerSymbol_hexagram_open               ScatterpolarglMarkerSymbol = "hexagram-open"
	ScatterpolarglMarkerSymbol218                          ScatterpolarglMarkerSymbol = "218"
	ScatterpolarglMarkerSymbol_hexagram_dot                ScatterpolarglMarkerSymbol = "hexagram-dot"
	ScatterpolarglMarkerSymbol318                          ScatterpolarglMarkerSymbol = "318"
	ScatterpolarglMarkerSymbol_hexagram_open_dot           ScatterpolarglMarkerSymbol = "hexagram-open-dot"
	ScatterpolarglMarkerSymbol19                           ScatterpolarglMarkerSymbol = "19"
	ScatterpolarglMarkerSymbol_star_triangle_up            ScatterpolarglMarkerSymbol = "star-triangle-up"
	ScatterpolarglMarkerSymbol119                          ScatterpolarglMarkerSymbol = "119"
	ScatterpolarglMarkerSymbol_star_triangle_up_open       ScatterpolarglMarkerSymbol = "star-triangle-up-open"
	ScatterpolarglMarkerSymbol219                          ScatterpolarglMarkerSymbol = "219"
	ScatterpolarglMarkerSymbol_star_triangle_up_dot        ScatterpolarglMarkerSymbol = "star-triangle-up-dot"
	ScatterpolarglMarkerSymbol319                          ScatterpolarglMarkerSymbol = "319"
	ScatterpolarglMarkerSymbol_star_triangle_up_open_dot   ScatterpolarglMarkerSymbol = "star-triangle-up-open-dot"
	ScatterpolarglMarkerSymbol20                           ScatterpolarglMarkerSymbol = "20"
	ScatterpolarglMarkerSymbol_star_triangle_down          ScatterpolarglMarkerSymbol = "star-triangle-down"
	ScatterpolarglMarkerSymbol120                          ScatterpolarglMarkerSymbol = "120"
	ScatterpolarglMarkerSymbol_star_triangle_down_open     ScatterpolarglMarkerSymbol = "star-triangle-down-open"
	ScatterpolarglMarkerSymbol220                          ScatterpolarglMarkerSymbol = "220"
	ScatterpolarglMarkerSymbol_star_triangle_down_dot      ScatterpolarglMarkerSymbol = "star-triangle-down-dot"
	ScatterpolarglMarkerSymbol320                          ScatterpolarglMarkerSymbol = "320"
	ScatterpolarglMarkerSymbol_star_triangle_down_open_dot ScatterpolarglMarkerSymbol = "star-triangle-down-open-dot"
	ScatterpolarglMarkerSymbol21                           ScatterpolarglMarkerSymbol = "21"
	ScatterpolarglMarkerSymbol_star_square                 ScatterpolarglMarkerSymbol = "star-square"
	ScatterpolarglMarkerSymbol121                          ScatterpolarglMarkerSymbol = "121"
	ScatterpolarglMarkerSymbol_star_square_open            ScatterpolarglMarkerSymbol = "star-square-open"
	ScatterpolarglMarkerSymbol221                          ScatterpolarglMarkerSymbol = "221"
	ScatterpolarglMarkerSymbol_star_square_dot             ScatterpolarglMarkerSymbol = "star-square-dot"
	ScatterpolarglMarkerSymbol321                          ScatterpolarglMarkerSymbol = "321"
	ScatterpolarglMarkerSymbol_star_square_open_dot        ScatterpolarglMarkerSymbol = "star-square-open-dot"
	ScatterpolarglMarkerSymbol22                           ScatterpolarglMarkerSymbol = "22"
	ScatterpolarglMarkerSymbol_star_diamond                ScatterpolarglMarkerSymbol = "star-diamond"
	ScatterpolarglMarkerSymbol122                          ScatterpolarglMarkerSymbol = "122"
	ScatterpolarglMarkerSymbol_star_diamond_open           ScatterpolarglMarkerSymbol = "star-diamond-open"
	ScatterpolarglMarkerSymbol222                          ScatterpolarglMarkerSymbol = "222"
	ScatterpolarglMarkerSymbol_star_diamond_dot            ScatterpolarglMarkerSymbol = "star-diamond-dot"
	ScatterpolarglMarkerSymbol322                          ScatterpolarglMarkerSymbol = "322"
	ScatterpolarglMarkerSymbol_star_diamond_open_dot       ScatterpolarglMarkerSymbol = "star-diamond-open-dot"
	ScatterpolarglMarkerSymbol23                           ScatterpolarglMarkerSymbol = "23"
	ScatterpolarglMarkerSymbol_diamond_tall                ScatterpolarglMarkerSymbol = "diamond-tall"
	ScatterpolarglMarkerSymbol123                          ScatterpolarglMarkerSymbol = "123"
	ScatterpolarglMarkerSymbol_diamond_tall_open           ScatterpolarglMarkerSymbol = "diamond-tall-open"
	ScatterpolarglMarkerSymbol223                          ScatterpolarglMarkerSymbol = "223"
	ScatterpolarglMarkerSymbol_diamond_tall_dot            ScatterpolarglMarkerSymbol = "diamond-tall-dot"
	ScatterpolarglMarkerSymbol323                          ScatterpolarglMarkerSymbol = "323"
	ScatterpolarglMarkerSymbol_diamond_tall_open_dot       ScatterpolarglMarkerSymbol = "diamond-tall-open-dot"
	ScatterpolarglMarkerSymbol24                           ScatterpolarglMarkerSymbol = "24"
	ScatterpolarglMarkerSymbol_diamond_wide                ScatterpolarglMarkerSymbol = "diamond-wide"
	ScatterpolarglMarkerSymbol124                          ScatterpolarglMarkerSymbol = "124"
	ScatterpolarglMarkerSymbol_diamond_wide_open           ScatterpolarglMarkerSymbol = "diamond-wide-open"
	ScatterpolarglMarkerSymbol224                          ScatterpolarglMarkerSymbol = "224"
	ScatterpolarglMarkerSymbol_diamond_wide_dot            ScatterpolarglMarkerSymbol = "diamond-wide-dot"
	ScatterpolarglMarkerSymbol324                          ScatterpolarglMarkerSymbol = "324"
	ScatterpolarglMarkerSymbol_diamond_wide_open_dot       ScatterpolarglMarkerSymbol = "diamond-wide-open-dot"
	ScatterpolarglMarkerSymbol25                           ScatterpolarglMarkerSymbol = "25"
	ScatterpolarglMarkerSymbol_hourglass                   ScatterpolarglMarkerSymbol = "hourglass"
	ScatterpolarglMarkerSymbol125                          ScatterpolarglMarkerSymbol = "125"
	ScatterpolarglMarkerSymbol_hourglass_open              ScatterpolarglMarkerSymbol = "hourglass-open"
	ScatterpolarglMarkerSymbol26                           ScatterpolarglMarkerSymbol = "26"
	ScatterpolarglMarkerSymbol_bowtie                      ScatterpolarglMarkerSymbol = "bowtie"
	ScatterpolarglMarkerSymbol126                          ScatterpolarglMarkerSymbol = "126"
	ScatterpolarglMarkerSymbol_bowtie_open                 ScatterpolarglMarkerSymbol = "bowtie-open"
	ScatterpolarglMarkerSymbol27                           ScatterpolarglMarkerSymbol = "27"
	ScatterpolarglMarkerSymbol_circle_cross                ScatterpolarglMarkerSymbol = "circle-cross"
	ScatterpolarglMarkerSymbol127                          ScatterpolarglMarkerSymbol = "127"
	ScatterpolarglMarkerSymbol_circle_cross_open           ScatterpolarglMarkerSymbol = "circle-cross-open"
	ScatterpolarglMarkerSymbol28                           ScatterpolarglMarkerSymbol = "28"
	ScatterpolarglMarkerSymbol_circle_x                    ScatterpolarglMarkerSymbol = "circle-x"
	ScatterpolarglMarkerSymbol128                          ScatterpolarglMarkerSymbol = "128"
	ScatterpolarglMarkerSymbol_circle_x_open               ScatterpolarglMarkerSymbol = "circle-x-open"
	ScatterpolarglMarkerSymbol29                           ScatterpolarglMarkerSymbol = "29"
	ScatterpolarglMarkerSymbol_square_cross                ScatterpolarglMarkerSymbol = "square-cross"
	ScatterpolarglMarkerSymbol129                          ScatterpolarglMarkerSymbol = "129"
	ScatterpolarglMarkerSymbol_square_cross_open           ScatterpolarglMarkerSymbol = "square-cross-open"
	ScatterpolarglMarkerSymbol30                           ScatterpolarglMarkerSymbol = "30"
	ScatterpolarglMarkerSymbol_square_x                    ScatterpolarglMarkerSymbol = "square-x"
	ScatterpolarglMarkerSymbol130                          ScatterpolarglMarkerSymbol = "130"
	ScatterpolarglMarkerSymbol_square_x_open               ScatterpolarglMarkerSymbol = "square-x-open"
	ScatterpolarglMarkerSymbol31                           ScatterpolarglMarkerSymbol = "31"
	ScatterpolarglMarkerSymbol_diamond_cross               ScatterpolarglMarkerSymbol = "diamond-cross"
	ScatterpolarglMarkerSymbol131                          ScatterpolarglMarkerSymbol = "131"
	ScatterpolarglMarkerSymbol_diamond_cross_open          ScatterpolarglMarkerSymbol = "diamond-cross-open"
	ScatterpolarglMarkerSymbol32                           ScatterpolarglMarkerSymbol = "32"
	ScatterpolarglMarkerSymbol_diamond_x                   ScatterpolarglMarkerSymbol = "diamond-x"
	ScatterpolarglMarkerSymbol132                          ScatterpolarglMarkerSymbol = "132"
	ScatterpolarglMarkerSymbol_diamond_x_open              ScatterpolarglMarkerSymbol = "diamond-x-open"
	ScatterpolarglMarkerSymbol33                           ScatterpolarglMarkerSymbol = "33"
	ScatterpolarglMarkerSymbol_cross_thin                  ScatterpolarglMarkerSymbol = "cross-thin"
	ScatterpolarglMarkerSymbol133                          ScatterpolarglMarkerSymbol = "133"
	ScatterpolarglMarkerSymbol_cross_thin_open             ScatterpolarglMarkerSymbol = "cross-thin-open"
	ScatterpolarglMarkerSymbol34                           ScatterpolarglMarkerSymbol = "34"
	ScatterpolarglMarkerSymbol_x_thin                      ScatterpolarglMarkerSymbol = "x-thin"
	ScatterpolarglMarkerSymbol134                          ScatterpolarglMarkerSymbol = "134"
	ScatterpolarglMarkerSymbol_x_thin_open                 ScatterpolarglMarkerSymbol = "x-thin-open"
	ScatterpolarglMarkerSymbol35                           ScatterpolarglMarkerSymbol = "35"
	ScatterpolarglMarkerSymbol_asterisk                    ScatterpolarglMarkerSymbol = "asterisk"
	ScatterpolarglMarkerSymbol135                          ScatterpolarglMarkerSymbol = "135"
	ScatterpolarglMarkerSymbol_asterisk_open               ScatterpolarglMarkerSymbol = "asterisk-open"
	ScatterpolarglMarkerSymbol36                           ScatterpolarglMarkerSymbol = "36"
	ScatterpolarglMarkerSymbol_hash                        ScatterpolarglMarkerSymbol = "hash"
	ScatterpolarglMarkerSymbol136                          ScatterpolarglMarkerSymbol = "136"
	ScatterpolarglMarkerSymbol_hash_open                   ScatterpolarglMarkerSymbol = "hash-open"
	ScatterpolarglMarkerSymbol236                          ScatterpolarglMarkerSymbol = "236"
	ScatterpolarglMarkerSymbol_hash_dot                    ScatterpolarglMarkerSymbol = "hash-dot"
	ScatterpolarglMarkerSymbol336                          ScatterpolarglMarkerSymbol = "336"
	ScatterpolarglMarkerSymbol_hash_open_dot               ScatterpolarglMarkerSymbol = "hash-open-dot"
	ScatterpolarglMarkerSymbol37                           ScatterpolarglMarkerSymbol = "37"
	ScatterpolarglMarkerSymbol_y_up                        ScatterpolarglMarkerSymbol = "y-up"
	ScatterpolarglMarkerSymbol137                          ScatterpolarglMarkerSymbol = "137"
	ScatterpolarglMarkerSymbol_y_up_open                   ScatterpolarglMarkerSymbol = "y-up-open"
	ScatterpolarglMarkerSymbol38                           ScatterpolarglMarkerSymbol = "38"
	ScatterpolarglMarkerSymbol_y_down                      ScatterpolarglMarkerSymbol = "y-down"
	ScatterpolarglMarkerSymbol138                          ScatterpolarglMarkerSymbol = "138"
	ScatterpolarglMarkerSymbol_y_down_open                 ScatterpolarglMarkerSymbol = "y-down-open"
	ScatterpolarglMarkerSymbol39                           ScatterpolarglMarkerSymbol = "39"
	ScatterpolarglMarkerSymbol_y_left                      ScatterpolarglMarkerSymbol = "y-left"
	ScatterpolarglMarkerSymbol139                          ScatterpolarglMarkerSymbol = "139"
	ScatterpolarglMarkerSymbol_y_left_open                 ScatterpolarglMarkerSymbol = "y-left-open"
	ScatterpolarglMarkerSymbol40                           ScatterpolarglMarkerSymbol = "40"
	ScatterpolarglMarkerSymbol_y_right                     ScatterpolarglMarkerSymbol = "y-right"
	ScatterpolarglMarkerSymbol140                          ScatterpolarglMarkerSymbol = "140"
	ScatterpolarglMarkerSymbol_y_right_open                ScatterpolarglMarkerSymbol = "y-right-open"
	ScatterpolarglMarkerSymbol41                           ScatterpolarglMarkerSymbol = "41"
	ScatterpolarglMarkerSymbol_line_ew                     ScatterpolarglMarkerSymbol = "line-ew"
	ScatterpolarglMarkerSymbol141                          ScatterpolarglMarkerSymbol = "141"
	ScatterpolarglMarkerSymbol_line_ew_open                ScatterpolarglMarkerSymbol = "line-ew-open"
	ScatterpolarglMarkerSymbol42                           ScatterpolarglMarkerSymbol = "42"
	ScatterpolarglMarkerSymbol_line_ns                     ScatterpolarglMarkerSymbol = "line-ns"
	ScatterpolarglMarkerSymbol142                          ScatterpolarglMarkerSymbol = "142"
	ScatterpolarglMarkerSymbol_line_ns_open                ScatterpolarglMarkerSymbol = "line-ns-open"
	ScatterpolarglMarkerSymbol43                           ScatterpolarglMarkerSymbol = "43"
	ScatterpolarglMarkerSymbol_line_ne                     ScatterpolarglMarkerSymbol = "line-ne"
	ScatterpolarglMarkerSymbol143                          ScatterpolarglMarkerSymbol = "143"
	ScatterpolarglMarkerSymbol_line_ne_open                ScatterpolarglMarkerSymbol = "line-ne-open"
	ScatterpolarglMarkerSymbol44                           ScatterpolarglMarkerSymbol = "44"
	ScatterpolarglMarkerSymbol_line_nw                     ScatterpolarglMarkerSymbol = "line-nw"
	ScatterpolarglMarkerSymbol144                          ScatterpolarglMarkerSymbol = "144"
	ScatterpolarglMarkerSymbol_line_nw_open                ScatterpolarglMarkerSymbol = "line-nw-open"
)

// ScatterpolarglTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterpolarglTextposition string

const (
	ScatterpolarglTextposition_topleft      ScatterpolarglTextposition = "top left"
	ScatterpolarglTextposition_topcenter    ScatterpolarglTextposition = "top center"
	ScatterpolarglTextposition_topright     ScatterpolarglTextposition = "top right"
	ScatterpolarglTextposition_middleleft   ScatterpolarglTextposition = "middle left"
	ScatterpolarglTextposition_middlecenter ScatterpolarglTextposition = "middle center"
	ScatterpolarglTextposition_middleright  ScatterpolarglTextposition = "middle right"
	ScatterpolarglTextposition_bottomleft   ScatterpolarglTextposition = "bottom left"
	ScatterpolarglTextposition_bottomcenter ScatterpolarglTextposition = "bottom center"
	ScatterpolarglTextposition_bottomright  ScatterpolarglTextposition = "bottom right"
)

// ScatterpolarglThetaunit Sets the unit of input *theta* values. Has an effect only when on *linear* angular axes.
type ScatterpolarglThetaunit string

const (
	ScatterpolarglThetaunit_radians  ScatterpolarglThetaunit = "radians"
	ScatterpolarglThetaunit_degrees  ScatterpolarglThetaunit = "degrees"
	ScatterpolarglThetaunit_gradians ScatterpolarglThetaunit = "gradians"
)

// ScatterpolarglVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterpolarglVisible interface{}

var (
	ScatterpolarglVisible_True       ScatterpolarglVisible = true
	ScatterpolarglVisible_False      ScatterpolarglVisible = false
	ScatterpolarglVisible_legendonly ScatterpolarglVisible = "legendonly"
)

// ScatterternaryFill Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. scatterternary has a subset of the options available to scatter. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other.
type ScatterternaryFill string

const (
	ScatterternaryFill_none   ScatterternaryFill = "none"
	ScatterternaryFill_toself ScatterternaryFill = "toself"
	ScatterternaryFill_tonext ScatterternaryFill = "tonext"
)

// ScatterternaryHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterternaryHoverlabelAlign string

const (
	ScatterternaryHoverlabelAlign_left  ScatterternaryHoverlabelAlign = "left"
	ScatterternaryHoverlabelAlign_right ScatterternaryHoverlabelAlign = "right"
	ScatterternaryHoverlabelAlign_auto  ScatterternaryHoverlabelAlign = "auto"
)

// ScatterternaryLineShape Determines the line shape. With *spline* the lines are drawn using spline interpolation. The other available values correspond to step-wise line shapes.
type ScatterternaryLineShape string

const (
	ScatterternaryLineShape_linear ScatterternaryLineShape = "linear"
	ScatterternaryLineShape_spline ScatterternaryLineShape = "spline"
)

// ScatterternaryMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterternaryMarkerColorbarExponentformat string

const (
	ScatterternaryMarkerColorbarExponentformat_none  ScatterternaryMarkerColorbarExponentformat = "none"
	ScatterternaryMarkerColorbarExponentformat_e     ScatterternaryMarkerColorbarExponentformat = "e"
	ScatterternaryMarkerColorbarExponentformat_E     ScatterternaryMarkerColorbarExponentformat = "E"
	ScatterternaryMarkerColorbarExponentformat_power ScatterternaryMarkerColorbarExponentformat = "power"
	ScatterternaryMarkerColorbarExponentformat_SI    ScatterternaryMarkerColorbarExponentformat = "SI"
	ScatterternaryMarkerColorbarExponentformat_B     ScatterternaryMarkerColorbarExponentformat = "B"
)

// ScatterternaryMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterternaryMarkerColorbarLenmode string

const (
	ScatterternaryMarkerColorbarLenmode_fraction ScatterternaryMarkerColorbarLenmode = "fraction"
	ScatterternaryMarkerColorbarLenmode_pixels   ScatterternaryMarkerColorbarLenmode = "pixels"
)

// ScatterternaryMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterternaryMarkerColorbarShowexponent string

const (
	ScatterternaryMarkerColorbarShowexponent_all   ScatterternaryMarkerColorbarShowexponent = "all"
	ScatterternaryMarkerColorbarShowexponent_first ScatterternaryMarkerColorbarShowexponent = "first"
	ScatterternaryMarkerColorbarShowexponent_last  ScatterternaryMarkerColorbarShowexponent = "last"
	ScatterternaryMarkerColorbarShowexponent_none  ScatterternaryMarkerColorbarShowexponent = "none"
)

// ScatterternaryMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterternaryMarkerColorbarShowtickprefix string

const (
	ScatterternaryMarkerColorbarShowtickprefix_all   ScatterternaryMarkerColorbarShowtickprefix = "all"
	ScatterternaryMarkerColorbarShowtickprefix_first ScatterternaryMarkerColorbarShowtickprefix = "first"
	ScatterternaryMarkerColorbarShowtickprefix_last  ScatterternaryMarkerColorbarShowtickprefix = "last"
	ScatterternaryMarkerColorbarShowtickprefix_none  ScatterternaryMarkerColorbarShowtickprefix = "none"
)

// ScatterternaryMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterternaryMarkerColorbarShowticksuffix string

const (
	ScatterternaryMarkerColorbarShowticksuffix_all   ScatterternaryMarkerColorbarShowticksuffix = "all"
	ScatterternaryMarkerColorbarShowticksuffix_first ScatterternaryMarkerColorbarShowticksuffix = "first"
	ScatterternaryMarkerColorbarShowticksuffix_last  ScatterternaryMarkerColorbarShowticksuffix = "last"
	ScatterternaryMarkerColorbarShowticksuffix_none  ScatterternaryMarkerColorbarShowticksuffix = "none"
)

// ScatterternaryMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterternaryMarkerColorbarThicknessmode string

const (
	ScatterternaryMarkerColorbarThicknessmode_fraction ScatterternaryMarkerColorbarThicknessmode = "fraction"
	ScatterternaryMarkerColorbarThicknessmode_pixels   ScatterternaryMarkerColorbarThicknessmode = "pixels"
)

// ScatterternaryMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterternaryMarkerColorbarTickmode string

const (
	ScatterternaryMarkerColorbarTickmode_auto   ScatterternaryMarkerColorbarTickmode = "auto"
	ScatterternaryMarkerColorbarTickmode_linear ScatterternaryMarkerColorbarTickmode = "linear"
	ScatterternaryMarkerColorbarTickmode_array  ScatterternaryMarkerColorbarTickmode = "array"
)

// ScatterternaryMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterternaryMarkerColorbarTicks string

const (
	ScatterternaryMarkerColorbarTicks_outside ScatterternaryMarkerColorbarTicks = "outside"
	ScatterternaryMarkerColorbarTicks_inside  ScatterternaryMarkerColorbarTicks = "inside"
)

// ScatterternaryMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterternaryMarkerColorbarTitleSide string

const (
	ScatterternaryMarkerColorbarTitleSide_right  ScatterternaryMarkerColorbarTitleSide = "right"
	ScatterternaryMarkerColorbarTitleSide_top    ScatterternaryMarkerColorbarTitleSide = "top"
	ScatterternaryMarkerColorbarTitleSide_bottom ScatterternaryMarkerColorbarTitleSide = "bottom"
)

// ScatterternaryMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterternaryMarkerColorbarXanchor string

const (
	ScatterternaryMarkerColorbarXanchor_left   ScatterternaryMarkerColorbarXanchor = "left"
	ScatterternaryMarkerColorbarXanchor_center ScatterternaryMarkerColorbarXanchor = "center"
	ScatterternaryMarkerColorbarXanchor_right  ScatterternaryMarkerColorbarXanchor = "right"
)

// ScatterternaryMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterternaryMarkerColorbarYanchor string

const (
	ScatterternaryMarkerColorbarYanchor_top    ScatterternaryMarkerColorbarYanchor = "top"
	ScatterternaryMarkerColorbarYanchor_middle ScatterternaryMarkerColorbarYanchor = "middle"
	ScatterternaryMarkerColorbarYanchor_bottom ScatterternaryMarkerColorbarYanchor = "bottom"
)

// ScatterternaryMarkerGradientType Sets the type of gradient used to fill the markers
type ScatterternaryMarkerGradientType string

const (
	ScatterternaryMarkerGradientType_radial     ScatterternaryMarkerGradientType = "radial"
	ScatterternaryMarkerGradientType_horizontal ScatterternaryMarkerGradientType = "horizontal"
	ScatterternaryMarkerGradientType_vertical   ScatterternaryMarkerGradientType = "vertical"
	ScatterternaryMarkerGradientType_none       ScatterternaryMarkerGradientType = "none"
)

// ScatterternaryMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterternaryMarkerSizemode string

const (
	ScatterternaryMarkerSizemode_diameter ScatterternaryMarkerSizemode = "diameter"
	ScatterternaryMarkerSizemode_area     ScatterternaryMarkerSizemode = "area"
)

// ScatterternaryMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterternaryMarkerSymbol string

const (
	ScatterternaryMarkerSymbol0                            ScatterternaryMarkerSymbol = "0"
	ScatterternaryMarkerSymbol_circle                      ScatterternaryMarkerSymbol = "circle"
	ScatterternaryMarkerSymbol100                          ScatterternaryMarkerSymbol = "100"
	ScatterternaryMarkerSymbol_circle_open                 ScatterternaryMarkerSymbol = "circle-open"
	ScatterternaryMarkerSymbol200                          ScatterternaryMarkerSymbol = "200"
	ScatterternaryMarkerSymbol_circle_dot                  ScatterternaryMarkerSymbol = "circle-dot"
	ScatterternaryMarkerSymbol300                          ScatterternaryMarkerSymbol = "300"
	ScatterternaryMarkerSymbol_circle_open_dot             ScatterternaryMarkerSymbol = "circle-open-dot"
	ScatterternaryMarkerSymbol1                            ScatterternaryMarkerSymbol = "1"
	ScatterternaryMarkerSymbol_square                      ScatterternaryMarkerSymbol = "square"
	ScatterternaryMarkerSymbol101                          ScatterternaryMarkerSymbol = "101"
	ScatterternaryMarkerSymbol_square_open                 ScatterternaryMarkerSymbol = "square-open"
	ScatterternaryMarkerSymbol201                          ScatterternaryMarkerSymbol = "201"
	ScatterternaryMarkerSymbol_square_dot                  ScatterternaryMarkerSymbol = "square-dot"
	ScatterternaryMarkerSymbol301                          ScatterternaryMarkerSymbol = "301"
	ScatterternaryMarkerSymbol_square_open_dot             ScatterternaryMarkerSymbol = "square-open-dot"
	ScatterternaryMarkerSymbol2                            ScatterternaryMarkerSymbol = "2"
	ScatterternaryMarkerSymbol_diamond                     ScatterternaryMarkerSymbol = "diamond"
	ScatterternaryMarkerSymbol102                          ScatterternaryMarkerSymbol = "102"
	ScatterternaryMarkerSymbol_diamond_open                ScatterternaryMarkerSymbol = "diamond-open"
	ScatterternaryMarkerSymbol202                          ScatterternaryMarkerSymbol = "202"
	ScatterternaryMarkerSymbol_diamond_dot                 ScatterternaryMarkerSymbol = "diamond-dot"
	ScatterternaryMarkerSymbol302                          ScatterternaryMarkerSymbol = "302"
	ScatterternaryMarkerSymbol_diamond_open_dot            ScatterternaryMarkerSymbol = "diamond-open-dot"
	ScatterternaryMarkerSymbol3                            ScatterternaryMarkerSymbol = "3"
	ScatterternaryMarkerSymbol_cross                       ScatterternaryMarkerSymbol = "cross"
	ScatterternaryMarkerSymbol103                          ScatterternaryMarkerSymbol = "103"
	ScatterternaryMarkerSymbol_cross_open                  ScatterternaryMarkerSymbol = "cross-open"
	ScatterternaryMarkerSymbol203                          ScatterternaryMarkerSymbol = "203"
	ScatterternaryMarkerSymbol_cross_dot                   ScatterternaryMarkerSymbol = "cross-dot"
	ScatterternaryMarkerSymbol303                          ScatterternaryMarkerSymbol = "303"
	ScatterternaryMarkerSymbol_cross_open_dot              ScatterternaryMarkerSymbol = "cross-open-dot"
	ScatterternaryMarkerSymbol4                            ScatterternaryMarkerSymbol = "4"
	ScatterternaryMarkerSymbol_x                           ScatterternaryMarkerSymbol = "x"
	ScatterternaryMarkerSymbol104                          ScatterternaryMarkerSymbol = "104"
	ScatterternaryMarkerSymbol_x_open                      ScatterternaryMarkerSymbol = "x-open"
	ScatterternaryMarkerSymbol204                          ScatterternaryMarkerSymbol = "204"
	ScatterternaryMarkerSymbol_x_dot                       ScatterternaryMarkerSymbol = "x-dot"
	ScatterternaryMarkerSymbol304                          ScatterternaryMarkerSymbol = "304"
	ScatterternaryMarkerSymbol_x_open_dot                  ScatterternaryMarkerSymbol = "x-open-dot"
	ScatterternaryMarkerSymbol5                            ScatterternaryMarkerSymbol = "5"
	ScatterternaryMarkerSymbol_triangle_up                 ScatterternaryMarkerSymbol = "triangle-up"
	ScatterternaryMarkerSymbol105                          ScatterternaryMarkerSymbol = "105"
	ScatterternaryMarkerSymbol_triangle_up_open            ScatterternaryMarkerSymbol = "triangle-up-open"
	ScatterternaryMarkerSymbol205                          ScatterternaryMarkerSymbol = "205"
	ScatterternaryMarkerSymbol_triangle_up_dot             ScatterternaryMarkerSymbol = "triangle-up-dot"
	ScatterternaryMarkerSymbol305                          ScatterternaryMarkerSymbol = "305"
	ScatterternaryMarkerSymbol_triangle_up_open_dot        ScatterternaryMarkerSymbol = "triangle-up-open-dot"
	ScatterternaryMarkerSymbol6                            ScatterternaryMarkerSymbol = "6"
	ScatterternaryMarkerSymbol_triangle_down               ScatterternaryMarkerSymbol = "triangle-down"
	ScatterternaryMarkerSymbol106                          ScatterternaryMarkerSymbol = "106"
	ScatterternaryMarkerSymbol_triangle_down_open          ScatterternaryMarkerSymbol = "triangle-down-open"
	ScatterternaryMarkerSymbol206                          ScatterternaryMarkerSymbol = "206"
	ScatterternaryMarkerSymbol_triangle_down_dot           ScatterternaryMarkerSymbol = "triangle-down-dot"
	ScatterternaryMarkerSymbol306                          ScatterternaryMarkerSymbol = "306"
	ScatterternaryMarkerSymbol_triangle_down_open_dot      ScatterternaryMarkerSymbol = "triangle-down-open-dot"
	ScatterternaryMarkerSymbol7                            ScatterternaryMarkerSymbol = "7"
	ScatterternaryMarkerSymbol_triangle_left               ScatterternaryMarkerSymbol = "triangle-left"
	ScatterternaryMarkerSymbol107                          ScatterternaryMarkerSymbol = "107"
	ScatterternaryMarkerSymbol_triangle_left_open          ScatterternaryMarkerSymbol = "triangle-left-open"
	ScatterternaryMarkerSymbol207                          ScatterternaryMarkerSymbol = "207"
	ScatterternaryMarkerSymbol_triangle_left_dot           ScatterternaryMarkerSymbol = "triangle-left-dot"
	ScatterternaryMarkerSymbol307                          ScatterternaryMarkerSymbol = "307"
	ScatterternaryMarkerSymbol_triangle_left_open_dot      ScatterternaryMarkerSymbol = "triangle-left-open-dot"
	ScatterternaryMarkerSymbol8                            ScatterternaryMarkerSymbol = "8"
	ScatterternaryMarkerSymbol_triangle_right              ScatterternaryMarkerSymbol = "triangle-right"
	ScatterternaryMarkerSymbol108                          ScatterternaryMarkerSymbol = "108"
	ScatterternaryMarkerSymbol_triangle_right_open         ScatterternaryMarkerSymbol = "triangle-right-open"
	ScatterternaryMarkerSymbol208                          ScatterternaryMarkerSymbol = "208"
	ScatterternaryMarkerSymbol_triangle_right_dot          ScatterternaryMarkerSymbol = "triangle-right-dot"
	ScatterternaryMarkerSymbol308                          ScatterternaryMarkerSymbol = "308"
	ScatterternaryMarkerSymbol_triangle_right_open_dot     ScatterternaryMarkerSymbol = "triangle-right-open-dot"
	ScatterternaryMarkerSymbol9                            ScatterternaryMarkerSymbol = "9"
	ScatterternaryMarkerSymbol_triangle_ne                 ScatterternaryMarkerSymbol = "triangle-ne"
	ScatterternaryMarkerSymbol109                          ScatterternaryMarkerSymbol = "109"
	ScatterternaryMarkerSymbol_triangle_ne_open            ScatterternaryMarkerSymbol = "triangle-ne-open"
	ScatterternaryMarkerSymbol209                          ScatterternaryMarkerSymbol = "209"
	ScatterternaryMarkerSymbol_triangle_ne_dot             ScatterternaryMarkerSymbol = "triangle-ne-dot"
	ScatterternaryMarkerSymbol309                          ScatterternaryMarkerSymbol = "309"
	ScatterternaryMarkerSymbol_triangle_ne_open_dot        ScatterternaryMarkerSymbol = "triangle-ne-open-dot"
	ScatterternaryMarkerSymbol10                           ScatterternaryMarkerSymbol = "10"
	ScatterternaryMarkerSymbol_triangle_se                 ScatterternaryMarkerSymbol = "triangle-se"
	ScatterternaryMarkerSymbol110                          ScatterternaryMarkerSymbol = "110"
	ScatterternaryMarkerSymbol_triangle_se_open            ScatterternaryMarkerSymbol = "triangle-se-open"
	ScatterternaryMarkerSymbol210                          ScatterternaryMarkerSymbol = "210"
	ScatterternaryMarkerSymbol_triangle_se_dot             ScatterternaryMarkerSymbol = "triangle-se-dot"
	ScatterternaryMarkerSymbol310                          ScatterternaryMarkerSymbol = "310"
	ScatterternaryMarkerSymbol_triangle_se_open_dot        ScatterternaryMarkerSymbol = "triangle-se-open-dot"
	ScatterternaryMarkerSymbol11                           ScatterternaryMarkerSymbol = "11"
	ScatterternaryMarkerSymbol_triangle_sw                 ScatterternaryMarkerSymbol = "triangle-sw"
	ScatterternaryMarkerSymbol111                          ScatterternaryMarkerSymbol = "111"
	ScatterternaryMarkerSymbol_triangle_sw_open            ScatterternaryMarkerSymbol = "triangle-sw-open"
	ScatterternaryMarkerSymbol211                          ScatterternaryMarkerSymbol = "211"
	ScatterternaryMarkerSymbol_triangle_sw_dot             ScatterternaryMarkerSymbol = "triangle-sw-dot"
	ScatterternaryMarkerSymbol311                          ScatterternaryMarkerSymbol = "311"
	ScatterternaryMarkerSymbol_triangle_sw_open_dot        ScatterternaryMarkerSymbol = "triangle-sw-open-dot"
	ScatterternaryMarkerSymbol12                           ScatterternaryMarkerSymbol = "12"
	ScatterternaryMarkerSymbol_triangle_nw                 ScatterternaryMarkerSymbol = "triangle-nw"
	ScatterternaryMarkerSymbol112                          ScatterternaryMarkerSymbol = "112"
	ScatterternaryMarkerSymbol_triangle_nw_open            ScatterternaryMarkerSymbol = "triangle-nw-open"
	ScatterternaryMarkerSymbol212                          ScatterternaryMarkerSymbol = "212"
	ScatterternaryMarkerSymbol_triangle_nw_dot             ScatterternaryMarkerSymbol = "triangle-nw-dot"
	ScatterternaryMarkerSymbol312                          ScatterternaryMarkerSymbol = "312"
	ScatterternaryMarkerSymbol_triangle_nw_open_dot        ScatterternaryMarkerSymbol = "triangle-nw-open-dot"
	ScatterternaryMarkerSymbol13                           ScatterternaryMarkerSymbol = "13"
	ScatterternaryMarkerSymbol_pentagon                    ScatterternaryMarkerSymbol = "pentagon"
	ScatterternaryMarkerSymbol113                          ScatterternaryMarkerSymbol = "113"
	ScatterternaryMarkerSymbol_pentagon_open               ScatterternaryMarkerSymbol = "pentagon-open"
	ScatterternaryMarkerSymbol213                          ScatterternaryMarkerSymbol = "213"
	ScatterternaryMarkerSymbol_pentagon_dot                ScatterternaryMarkerSymbol = "pentagon-dot"
	ScatterternaryMarkerSymbol313                          ScatterternaryMarkerSymbol = "313"
	ScatterternaryMarkerSymbol_pentagon_open_dot           ScatterternaryMarkerSymbol = "pentagon-open-dot"
	ScatterternaryMarkerSymbol14                           ScatterternaryMarkerSymbol = "14"
	ScatterternaryMarkerSymbol_hexagon                     ScatterternaryMarkerSymbol = "hexagon"
	ScatterternaryMarkerSymbol114                          ScatterternaryMarkerSymbol = "114"
	ScatterternaryMarkerSymbol_hexagon_open                ScatterternaryMarkerSymbol = "hexagon-open"
	ScatterternaryMarkerSymbol214                          ScatterternaryMarkerSymbol = "214"
	ScatterternaryMarkerSymbol_hexagon_dot                 ScatterternaryMarkerSymbol = "hexagon-dot"
	ScatterternaryMarkerSymbol314                          ScatterternaryMarkerSymbol = "314"
	ScatterternaryMarkerSymbol_hexagon_open_dot            ScatterternaryMarkerSymbol = "hexagon-open-dot"
	ScatterternaryMarkerSymbol15                           ScatterternaryMarkerSymbol = "15"
	ScatterternaryMarkerSymbol_hexagon2                    ScatterternaryMarkerSymbol = "hexagon2"
	ScatterternaryMarkerSymbol115                          ScatterternaryMarkerSymbol = "115"
	ScatterternaryMarkerSymbol_hexagon2_open               ScatterternaryMarkerSymbol = "hexagon2-open"
	ScatterternaryMarkerSymbol215                          ScatterternaryMarkerSymbol = "215"
	ScatterternaryMarkerSymbol_hexagon2_dot                ScatterternaryMarkerSymbol = "hexagon2-dot"
	ScatterternaryMarkerSymbol315                          ScatterternaryMarkerSymbol = "315"
	ScatterternaryMarkerSymbol_hexagon2_open_dot           ScatterternaryMarkerSymbol = "hexagon2-open-dot"
	ScatterternaryMarkerSymbol16                           ScatterternaryMarkerSymbol = "16"
	ScatterternaryMarkerSymbol_octagon                     ScatterternaryMarkerSymbol = "octagon"
	ScatterternaryMarkerSymbol116                          ScatterternaryMarkerSymbol = "116"
	ScatterternaryMarkerSymbol_octagon_open                ScatterternaryMarkerSymbol = "octagon-open"
	ScatterternaryMarkerSymbol216                          ScatterternaryMarkerSymbol = "216"
	ScatterternaryMarkerSymbol_octagon_dot                 ScatterternaryMarkerSymbol = "octagon-dot"
	ScatterternaryMarkerSymbol316                          ScatterternaryMarkerSymbol = "316"
	ScatterternaryMarkerSymbol_octagon_open_dot            ScatterternaryMarkerSymbol = "octagon-open-dot"
	ScatterternaryMarkerSymbol17                           ScatterternaryMarkerSymbol = "17"
	ScatterternaryMarkerSymbol_star                        ScatterternaryMarkerSymbol = "star"
	ScatterternaryMarkerSymbol117                          ScatterternaryMarkerSymbol = "117"
	ScatterternaryMarkerSymbol_star_open                   ScatterternaryMarkerSymbol = "star-open"
	ScatterternaryMarkerSymbol217                          ScatterternaryMarkerSymbol = "217"
	ScatterternaryMarkerSymbol_star_dot                    ScatterternaryMarkerSymbol = "star-dot"
	ScatterternaryMarkerSymbol317                          ScatterternaryMarkerSymbol = "317"
	ScatterternaryMarkerSymbol_star_open_dot               ScatterternaryMarkerSymbol = "star-open-dot"
	ScatterternaryMarkerSymbol18                           ScatterternaryMarkerSymbol = "18"
	ScatterternaryMarkerSymbol_hexagram                    ScatterternaryMarkerSymbol = "hexagram"
	ScatterternaryMarkerSymbol118                          ScatterternaryMarkerSymbol = "118"
	ScatterternaryMarkerSymbol_hexagram_open               ScatterternaryMarkerSymbol = "hexagram-open"
	ScatterternaryMarkerSymbol218                          ScatterternaryMarkerSymbol = "218"
	ScatterternaryMarkerSymbol_hexagram_dot                ScatterternaryMarkerSymbol = "hexagram-dot"
	ScatterternaryMarkerSymbol318                          ScatterternaryMarkerSymbol = "318"
	ScatterternaryMarkerSymbol_hexagram_open_dot           ScatterternaryMarkerSymbol = "hexagram-open-dot"
	ScatterternaryMarkerSymbol19                           ScatterternaryMarkerSymbol = "19"
	ScatterternaryMarkerSymbol_star_triangle_up            ScatterternaryMarkerSymbol = "star-triangle-up"
	ScatterternaryMarkerSymbol119                          ScatterternaryMarkerSymbol = "119"
	ScatterternaryMarkerSymbol_star_triangle_up_open       ScatterternaryMarkerSymbol = "star-triangle-up-open"
	ScatterternaryMarkerSymbol219                          ScatterternaryMarkerSymbol = "219"
	ScatterternaryMarkerSymbol_star_triangle_up_dot        ScatterternaryMarkerSymbol = "star-triangle-up-dot"
	ScatterternaryMarkerSymbol319                          ScatterternaryMarkerSymbol = "319"
	ScatterternaryMarkerSymbol_star_triangle_up_open_dot   ScatterternaryMarkerSymbol = "star-triangle-up-open-dot"
	ScatterternaryMarkerSymbol20                           ScatterternaryMarkerSymbol = "20"
	ScatterternaryMarkerSymbol_star_triangle_down          ScatterternaryMarkerSymbol = "star-triangle-down"
	ScatterternaryMarkerSymbol120                          ScatterternaryMarkerSymbol = "120"
	ScatterternaryMarkerSymbol_star_triangle_down_open     ScatterternaryMarkerSymbol = "star-triangle-down-open"
	ScatterternaryMarkerSymbol220                          ScatterternaryMarkerSymbol = "220"
	ScatterternaryMarkerSymbol_star_triangle_down_dot      ScatterternaryMarkerSymbol = "star-triangle-down-dot"
	ScatterternaryMarkerSymbol320                          ScatterternaryMarkerSymbol = "320"
	ScatterternaryMarkerSymbol_star_triangle_down_open_dot ScatterternaryMarkerSymbol = "star-triangle-down-open-dot"
	ScatterternaryMarkerSymbol21                           ScatterternaryMarkerSymbol = "21"
	ScatterternaryMarkerSymbol_star_square                 ScatterternaryMarkerSymbol = "star-square"
	ScatterternaryMarkerSymbol121                          ScatterternaryMarkerSymbol = "121"
	ScatterternaryMarkerSymbol_star_square_open            ScatterternaryMarkerSymbol = "star-square-open"
	ScatterternaryMarkerSymbol221                          ScatterternaryMarkerSymbol = "221"
	ScatterternaryMarkerSymbol_star_square_dot             ScatterternaryMarkerSymbol = "star-square-dot"
	ScatterternaryMarkerSymbol321                          ScatterternaryMarkerSymbol = "321"
	ScatterternaryMarkerSymbol_star_square_open_dot        ScatterternaryMarkerSymbol = "star-square-open-dot"
	ScatterternaryMarkerSymbol22                           ScatterternaryMarkerSymbol = "22"
	ScatterternaryMarkerSymbol_star_diamond                ScatterternaryMarkerSymbol = "star-diamond"
	ScatterternaryMarkerSymbol122                          ScatterternaryMarkerSymbol = "122"
	ScatterternaryMarkerSymbol_star_diamond_open           ScatterternaryMarkerSymbol = "star-diamond-open"
	ScatterternaryMarkerSymbol222                          ScatterternaryMarkerSymbol = "222"
	ScatterternaryMarkerSymbol_star_diamond_dot            ScatterternaryMarkerSymbol = "star-diamond-dot"
	ScatterternaryMarkerSymbol322                          ScatterternaryMarkerSymbol = "322"
	ScatterternaryMarkerSymbol_star_diamond_open_dot       ScatterternaryMarkerSymbol = "star-diamond-open-dot"
	ScatterternaryMarkerSymbol23                           ScatterternaryMarkerSymbol = "23"
	ScatterternaryMarkerSymbol_diamond_tall                ScatterternaryMarkerSymbol = "diamond-tall"
	ScatterternaryMarkerSymbol123                          ScatterternaryMarkerSymbol = "123"
	ScatterternaryMarkerSymbol_diamond_tall_open           ScatterternaryMarkerSymbol = "diamond-tall-open"
	ScatterternaryMarkerSymbol223                          ScatterternaryMarkerSymbol = "223"
	ScatterternaryMarkerSymbol_diamond_tall_dot            ScatterternaryMarkerSymbol = "diamond-tall-dot"
	ScatterternaryMarkerSymbol323                          ScatterternaryMarkerSymbol = "323"
	ScatterternaryMarkerSymbol_diamond_tall_open_dot       ScatterternaryMarkerSymbol = "diamond-tall-open-dot"
	ScatterternaryMarkerSymbol24                           ScatterternaryMarkerSymbol = "24"
	ScatterternaryMarkerSymbol_diamond_wide                ScatterternaryMarkerSymbol = "diamond-wide"
	ScatterternaryMarkerSymbol124                          ScatterternaryMarkerSymbol = "124"
	ScatterternaryMarkerSymbol_diamond_wide_open           ScatterternaryMarkerSymbol = "diamond-wide-open"
	ScatterternaryMarkerSymbol224                          ScatterternaryMarkerSymbol = "224"
	ScatterternaryMarkerSymbol_diamond_wide_dot            ScatterternaryMarkerSymbol = "diamond-wide-dot"
	ScatterternaryMarkerSymbol324                          ScatterternaryMarkerSymbol = "324"
	ScatterternaryMarkerSymbol_diamond_wide_open_dot       ScatterternaryMarkerSymbol = "diamond-wide-open-dot"
	ScatterternaryMarkerSymbol25                           ScatterternaryMarkerSymbol = "25"
	ScatterternaryMarkerSymbol_hourglass                   ScatterternaryMarkerSymbol = "hourglass"
	ScatterternaryMarkerSymbol125                          ScatterternaryMarkerSymbol = "125"
	ScatterternaryMarkerSymbol_hourglass_open              ScatterternaryMarkerSymbol = "hourglass-open"
	ScatterternaryMarkerSymbol26                           ScatterternaryMarkerSymbol = "26"
	ScatterternaryMarkerSymbol_bowtie                      ScatterternaryMarkerSymbol = "bowtie"
	ScatterternaryMarkerSymbol126                          ScatterternaryMarkerSymbol = "126"
	ScatterternaryMarkerSymbol_bowtie_open                 ScatterternaryMarkerSymbol = "bowtie-open"
	ScatterternaryMarkerSymbol27                           ScatterternaryMarkerSymbol = "27"
	ScatterternaryMarkerSymbol_circle_cross                ScatterternaryMarkerSymbol = "circle-cross"
	ScatterternaryMarkerSymbol127                          ScatterternaryMarkerSymbol = "127"
	ScatterternaryMarkerSymbol_circle_cross_open           ScatterternaryMarkerSymbol = "circle-cross-open"
	ScatterternaryMarkerSymbol28                           ScatterternaryMarkerSymbol = "28"
	ScatterternaryMarkerSymbol_circle_x                    ScatterternaryMarkerSymbol = "circle-x"
	ScatterternaryMarkerSymbol128                          ScatterternaryMarkerSymbol = "128"
	ScatterternaryMarkerSymbol_circle_x_open               ScatterternaryMarkerSymbol = "circle-x-open"
	ScatterternaryMarkerSymbol29                           ScatterternaryMarkerSymbol = "29"
	ScatterternaryMarkerSymbol_square_cross                ScatterternaryMarkerSymbol = "square-cross"
	ScatterternaryMarkerSymbol129                          ScatterternaryMarkerSymbol = "129"
	ScatterternaryMarkerSymbol_square_cross_open           ScatterternaryMarkerSymbol = "square-cross-open"
	ScatterternaryMarkerSymbol30                           ScatterternaryMarkerSymbol = "30"
	ScatterternaryMarkerSymbol_square_x                    ScatterternaryMarkerSymbol = "square-x"
	ScatterternaryMarkerSymbol130                          ScatterternaryMarkerSymbol = "130"
	ScatterternaryMarkerSymbol_square_x_open               ScatterternaryMarkerSymbol = "square-x-open"
	ScatterternaryMarkerSymbol31                           ScatterternaryMarkerSymbol = "31"
	ScatterternaryMarkerSymbol_diamond_cross               ScatterternaryMarkerSymbol = "diamond-cross"
	ScatterternaryMarkerSymbol131                          ScatterternaryMarkerSymbol = "131"
	ScatterternaryMarkerSymbol_diamond_cross_open          ScatterternaryMarkerSymbol = "diamond-cross-open"
	ScatterternaryMarkerSymbol32                           ScatterternaryMarkerSymbol = "32"
	ScatterternaryMarkerSymbol_diamond_x                   ScatterternaryMarkerSymbol = "diamond-x"
	ScatterternaryMarkerSymbol132                          ScatterternaryMarkerSymbol = "132"
	ScatterternaryMarkerSymbol_diamond_x_open              ScatterternaryMarkerSymbol = "diamond-x-open"
	ScatterternaryMarkerSymbol33                           ScatterternaryMarkerSymbol = "33"
	ScatterternaryMarkerSymbol_cross_thin                  ScatterternaryMarkerSymbol = "cross-thin"
	ScatterternaryMarkerSymbol133                          ScatterternaryMarkerSymbol = "133"
	ScatterternaryMarkerSymbol_cross_thin_open             ScatterternaryMarkerSymbol = "cross-thin-open"
	ScatterternaryMarkerSymbol34                           ScatterternaryMarkerSymbol = "34"
	ScatterternaryMarkerSymbol_x_thin                      ScatterternaryMarkerSymbol = "x-thin"
	ScatterternaryMarkerSymbol134                          ScatterternaryMarkerSymbol = "134"
	ScatterternaryMarkerSymbol_x_thin_open                 ScatterternaryMarkerSymbol = "x-thin-open"
	ScatterternaryMarkerSymbol35                           ScatterternaryMarkerSymbol = "35"
	ScatterternaryMarkerSymbol_asterisk                    ScatterternaryMarkerSymbol = "asterisk"
	ScatterternaryMarkerSymbol135                          ScatterternaryMarkerSymbol = "135"
	ScatterternaryMarkerSymbol_asterisk_open               ScatterternaryMarkerSymbol = "asterisk-open"
	ScatterternaryMarkerSymbol36                           ScatterternaryMarkerSymbol = "36"
	ScatterternaryMarkerSymbol_hash                        ScatterternaryMarkerSymbol = "hash"
	ScatterternaryMarkerSymbol136                          ScatterternaryMarkerSymbol = "136"
	ScatterternaryMarkerSymbol_hash_open                   ScatterternaryMarkerSymbol = "hash-open"
	ScatterternaryMarkerSymbol236                          ScatterternaryMarkerSymbol = "236"
	ScatterternaryMarkerSymbol_hash_dot                    ScatterternaryMarkerSymbol = "hash-dot"
	ScatterternaryMarkerSymbol336                          ScatterternaryMarkerSymbol = "336"
	ScatterternaryMarkerSymbol_hash_open_dot               ScatterternaryMarkerSymbol = "hash-open-dot"
	ScatterternaryMarkerSymbol37                           ScatterternaryMarkerSymbol = "37"
	ScatterternaryMarkerSymbol_y_up                        ScatterternaryMarkerSymbol = "y-up"
	ScatterternaryMarkerSymbol137                          ScatterternaryMarkerSymbol = "137"
	ScatterternaryMarkerSymbol_y_up_open                   ScatterternaryMarkerSymbol = "y-up-open"
	ScatterternaryMarkerSymbol38                           ScatterternaryMarkerSymbol = "38"
	ScatterternaryMarkerSymbol_y_down                      ScatterternaryMarkerSymbol = "y-down"
	ScatterternaryMarkerSymbol138                          ScatterternaryMarkerSymbol = "138"
	ScatterternaryMarkerSymbol_y_down_open                 ScatterternaryMarkerSymbol = "y-down-open"
	ScatterternaryMarkerSymbol39                           ScatterternaryMarkerSymbol = "39"
	ScatterternaryMarkerSymbol_y_left                      ScatterternaryMarkerSymbol = "y-left"
	ScatterternaryMarkerSymbol139                          ScatterternaryMarkerSymbol = "139"
	ScatterternaryMarkerSymbol_y_left_open                 ScatterternaryMarkerSymbol = "y-left-open"
	ScatterternaryMarkerSymbol40                           ScatterternaryMarkerSymbol = "40"
	ScatterternaryMarkerSymbol_y_right                     ScatterternaryMarkerSymbol = "y-right"
	ScatterternaryMarkerSymbol140                          ScatterternaryMarkerSymbol = "140"
	ScatterternaryMarkerSymbol_y_right_open                ScatterternaryMarkerSymbol = "y-right-open"
	ScatterternaryMarkerSymbol41                           ScatterternaryMarkerSymbol = "41"
	ScatterternaryMarkerSymbol_line_ew                     ScatterternaryMarkerSymbol = "line-ew"
	ScatterternaryMarkerSymbol141                          ScatterternaryMarkerSymbol = "141"
	ScatterternaryMarkerSymbol_line_ew_open                ScatterternaryMarkerSymbol = "line-ew-open"
	ScatterternaryMarkerSymbol42                           ScatterternaryMarkerSymbol = "42"
	ScatterternaryMarkerSymbol_line_ns                     ScatterternaryMarkerSymbol = "line-ns"
	ScatterternaryMarkerSymbol142                          ScatterternaryMarkerSymbol = "142"
	ScatterternaryMarkerSymbol_line_ns_open                ScatterternaryMarkerSymbol = "line-ns-open"
	ScatterternaryMarkerSymbol43                           ScatterternaryMarkerSymbol = "43"
	ScatterternaryMarkerSymbol_line_ne                     ScatterternaryMarkerSymbol = "line-ne"
	ScatterternaryMarkerSymbol143                          ScatterternaryMarkerSymbol = "143"
	ScatterternaryMarkerSymbol_line_ne_open                ScatterternaryMarkerSymbol = "line-ne-open"
	ScatterternaryMarkerSymbol44                           ScatterternaryMarkerSymbol = "44"
	ScatterternaryMarkerSymbol_line_nw                     ScatterternaryMarkerSymbol = "line-nw"
	ScatterternaryMarkerSymbol144                          ScatterternaryMarkerSymbol = "144"
	ScatterternaryMarkerSymbol_line_nw_open                ScatterternaryMarkerSymbol = "line-nw-open"
)

// ScatterternaryTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterternaryTextposition string

const (
	ScatterternaryTextposition_topleft      ScatterternaryTextposition = "top left"
	ScatterternaryTextposition_topcenter    ScatterternaryTextposition = "top center"
	ScatterternaryTextposition_topright     ScatterternaryTextposition = "top right"
	ScatterternaryTextposition_middleleft   ScatterternaryTextposition = "middle left"
	ScatterternaryTextposition_middlecenter ScatterternaryTextposition = "middle center"
	ScatterternaryTextposition_middleright  ScatterternaryTextposition = "middle right"
	ScatterternaryTextposition_bottomleft   ScatterternaryTextposition = "bottom left"
	ScatterternaryTextposition_bottomcenter ScatterternaryTextposition = "bottom center"
	ScatterternaryTextposition_bottomright  ScatterternaryTextposition = "bottom right"
)

// ScatterternaryVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterternaryVisible interface{}

var (
	ScatterternaryVisible_True       ScatterternaryVisible = true
	ScatterternaryVisible_False      ScatterternaryVisible = false
	ScatterternaryVisible_legendonly ScatterternaryVisible = "legendonly"
)

// SplomHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SplomHoverlabelAlign string

const (
	SplomHoverlabelAlign_left  SplomHoverlabelAlign = "left"
	SplomHoverlabelAlign_right SplomHoverlabelAlign = "right"
	SplomHoverlabelAlign_auto  SplomHoverlabelAlign = "auto"
)

// SplomMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type SplomMarkerColorbarExponentformat string

const (
	SplomMarkerColorbarExponentformat_none  SplomMarkerColorbarExponentformat = "none"
	SplomMarkerColorbarExponentformat_e     SplomMarkerColorbarExponentformat = "e"
	SplomMarkerColorbarExponentformat_E     SplomMarkerColorbarExponentformat = "E"
	SplomMarkerColorbarExponentformat_power SplomMarkerColorbarExponentformat = "power"
	SplomMarkerColorbarExponentformat_SI    SplomMarkerColorbarExponentformat = "SI"
	SplomMarkerColorbarExponentformat_B     SplomMarkerColorbarExponentformat = "B"
)

// SplomMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type SplomMarkerColorbarLenmode string

const (
	SplomMarkerColorbarLenmode_fraction SplomMarkerColorbarLenmode = "fraction"
	SplomMarkerColorbarLenmode_pixels   SplomMarkerColorbarLenmode = "pixels"
)

// SplomMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type SplomMarkerColorbarShowexponent string

const (
	SplomMarkerColorbarShowexponent_all   SplomMarkerColorbarShowexponent = "all"
	SplomMarkerColorbarShowexponent_first SplomMarkerColorbarShowexponent = "first"
	SplomMarkerColorbarShowexponent_last  SplomMarkerColorbarShowexponent = "last"
	SplomMarkerColorbarShowexponent_none  SplomMarkerColorbarShowexponent = "none"
)

// SplomMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type SplomMarkerColorbarShowtickprefix string

const (
	SplomMarkerColorbarShowtickprefix_all   SplomMarkerColorbarShowtickprefix = "all"
	SplomMarkerColorbarShowtickprefix_first SplomMarkerColorbarShowtickprefix = "first"
	SplomMarkerColorbarShowtickprefix_last  SplomMarkerColorbarShowtickprefix = "last"
	SplomMarkerColorbarShowtickprefix_none  SplomMarkerColorbarShowtickprefix = "none"
)

// SplomMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type SplomMarkerColorbarShowticksuffix string

const (
	SplomMarkerColorbarShowticksuffix_all   SplomMarkerColorbarShowticksuffix = "all"
	SplomMarkerColorbarShowticksuffix_first SplomMarkerColorbarShowticksuffix = "first"
	SplomMarkerColorbarShowticksuffix_last  SplomMarkerColorbarShowticksuffix = "last"
	SplomMarkerColorbarShowticksuffix_none  SplomMarkerColorbarShowticksuffix = "none"
)

// SplomMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type SplomMarkerColorbarThicknessmode string

const (
	SplomMarkerColorbarThicknessmode_fraction SplomMarkerColorbarThicknessmode = "fraction"
	SplomMarkerColorbarThicknessmode_pixels   SplomMarkerColorbarThicknessmode = "pixels"
)

// SplomMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type SplomMarkerColorbarTickmode string

const (
	SplomMarkerColorbarTickmode_auto   SplomMarkerColorbarTickmode = "auto"
	SplomMarkerColorbarTickmode_linear SplomMarkerColorbarTickmode = "linear"
	SplomMarkerColorbarTickmode_array  SplomMarkerColorbarTickmode = "array"
)

// SplomMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type SplomMarkerColorbarTicks string

const (
	SplomMarkerColorbarTicks_outside SplomMarkerColorbarTicks = "outside"
	SplomMarkerColorbarTicks_inside  SplomMarkerColorbarTicks = "inside"
)

// SplomMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type SplomMarkerColorbarTitleSide string

const (
	SplomMarkerColorbarTitleSide_right  SplomMarkerColorbarTitleSide = "right"
	SplomMarkerColorbarTitleSide_top    SplomMarkerColorbarTitleSide = "top"
	SplomMarkerColorbarTitleSide_bottom SplomMarkerColorbarTitleSide = "bottom"
)

// SplomMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type SplomMarkerColorbarXanchor string

const (
	SplomMarkerColorbarXanchor_left   SplomMarkerColorbarXanchor = "left"
	SplomMarkerColorbarXanchor_center SplomMarkerColorbarXanchor = "center"
	SplomMarkerColorbarXanchor_right  SplomMarkerColorbarXanchor = "right"
)

// SplomMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type SplomMarkerColorbarYanchor string

const (
	SplomMarkerColorbarYanchor_top    SplomMarkerColorbarYanchor = "top"
	SplomMarkerColorbarYanchor_middle SplomMarkerColorbarYanchor = "middle"
	SplomMarkerColorbarYanchor_bottom SplomMarkerColorbarYanchor = "bottom"
)

// SplomMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type SplomMarkerSizemode string

const (
	SplomMarkerSizemode_diameter SplomMarkerSizemode = "diameter"
	SplomMarkerSizemode_area     SplomMarkerSizemode = "area"
)

// SplomMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type SplomMarkerSymbol string

const (
	SplomMarkerSymbol0                            SplomMarkerSymbol = "0"
	SplomMarkerSymbol_circle                      SplomMarkerSymbol = "circle"
	SplomMarkerSymbol100                          SplomMarkerSymbol = "100"
	SplomMarkerSymbol_circle_open                 SplomMarkerSymbol = "circle-open"
	SplomMarkerSymbol200                          SplomMarkerSymbol = "200"
	SplomMarkerSymbol_circle_dot                  SplomMarkerSymbol = "circle-dot"
	SplomMarkerSymbol300                          SplomMarkerSymbol = "300"
	SplomMarkerSymbol_circle_open_dot             SplomMarkerSymbol = "circle-open-dot"
	SplomMarkerSymbol1                            SplomMarkerSymbol = "1"
	SplomMarkerSymbol_square                      SplomMarkerSymbol = "square"
	SplomMarkerSymbol101                          SplomMarkerSymbol = "101"
	SplomMarkerSymbol_square_open                 SplomMarkerSymbol = "square-open"
	SplomMarkerSymbol201                          SplomMarkerSymbol = "201"
	SplomMarkerSymbol_square_dot                  SplomMarkerSymbol = "square-dot"
	SplomMarkerSymbol301                          SplomMarkerSymbol = "301"
	SplomMarkerSymbol_square_open_dot             SplomMarkerSymbol = "square-open-dot"
	SplomMarkerSymbol2                            SplomMarkerSymbol = "2"
	SplomMarkerSymbol_diamond                     SplomMarkerSymbol = "diamond"
	SplomMarkerSymbol102                          SplomMarkerSymbol = "102"
	SplomMarkerSymbol_diamond_open                SplomMarkerSymbol = "diamond-open"
	SplomMarkerSymbol202                          SplomMarkerSymbol = "202"
	SplomMarkerSymbol_diamond_dot                 SplomMarkerSymbol = "diamond-dot"
	SplomMarkerSymbol302                          SplomMarkerSymbol = "302"
	SplomMarkerSymbol_diamond_open_dot            SplomMarkerSymbol = "diamond-open-dot"
	SplomMarkerSymbol3                            SplomMarkerSymbol = "3"
	SplomMarkerSymbol_cross                       SplomMarkerSymbol = "cross"
	SplomMarkerSymbol103                          SplomMarkerSymbol = "103"
	SplomMarkerSymbol_cross_open                  SplomMarkerSymbol = "cross-open"
	SplomMarkerSymbol203                          SplomMarkerSymbol = "203"
	SplomMarkerSymbol_cross_dot                   SplomMarkerSymbol = "cross-dot"
	SplomMarkerSymbol303                          SplomMarkerSymbol = "303"
	SplomMarkerSymbol_cross_open_dot              SplomMarkerSymbol = "cross-open-dot"
	SplomMarkerSymbol4                            SplomMarkerSymbol = "4"
	SplomMarkerSymbol_x                           SplomMarkerSymbol = "x"
	SplomMarkerSymbol104                          SplomMarkerSymbol = "104"
	SplomMarkerSymbol_x_open                      SplomMarkerSymbol = "x-open"
	SplomMarkerSymbol204                          SplomMarkerSymbol = "204"
	SplomMarkerSymbol_x_dot                       SplomMarkerSymbol = "x-dot"
	SplomMarkerSymbol304                          SplomMarkerSymbol = "304"
	SplomMarkerSymbol_x_open_dot                  SplomMarkerSymbol = "x-open-dot"
	SplomMarkerSymbol5                            SplomMarkerSymbol = "5"
	SplomMarkerSymbol_triangle_up                 SplomMarkerSymbol = "triangle-up"
	SplomMarkerSymbol105                          SplomMarkerSymbol = "105"
	SplomMarkerSymbol_triangle_up_open            SplomMarkerSymbol = "triangle-up-open"
	SplomMarkerSymbol205                          SplomMarkerSymbol = "205"
	SplomMarkerSymbol_triangle_up_dot             SplomMarkerSymbol = "triangle-up-dot"
	SplomMarkerSymbol305                          SplomMarkerSymbol = "305"
	SplomMarkerSymbol_triangle_up_open_dot        SplomMarkerSymbol = "triangle-up-open-dot"
	SplomMarkerSymbol6                            SplomMarkerSymbol = "6"
	SplomMarkerSymbol_triangle_down               SplomMarkerSymbol = "triangle-down"
	SplomMarkerSymbol106                          SplomMarkerSymbol = "106"
	SplomMarkerSymbol_triangle_down_open          SplomMarkerSymbol = "triangle-down-open"
	SplomMarkerSymbol206                          SplomMarkerSymbol = "206"
	SplomMarkerSymbol_triangle_down_dot           SplomMarkerSymbol = "triangle-down-dot"
	SplomMarkerSymbol306                          SplomMarkerSymbol = "306"
	SplomMarkerSymbol_triangle_down_open_dot      SplomMarkerSymbol = "triangle-down-open-dot"
	SplomMarkerSymbol7                            SplomMarkerSymbol = "7"
	SplomMarkerSymbol_triangle_left               SplomMarkerSymbol = "triangle-left"
	SplomMarkerSymbol107                          SplomMarkerSymbol = "107"
	SplomMarkerSymbol_triangle_left_open          SplomMarkerSymbol = "triangle-left-open"
	SplomMarkerSymbol207                          SplomMarkerSymbol = "207"
	SplomMarkerSymbol_triangle_left_dot           SplomMarkerSymbol = "triangle-left-dot"
	SplomMarkerSymbol307                          SplomMarkerSymbol = "307"
	SplomMarkerSymbol_triangle_left_open_dot      SplomMarkerSymbol = "triangle-left-open-dot"
	SplomMarkerSymbol8                            SplomMarkerSymbol = "8"
	SplomMarkerSymbol_triangle_right              SplomMarkerSymbol = "triangle-right"
	SplomMarkerSymbol108                          SplomMarkerSymbol = "108"
	SplomMarkerSymbol_triangle_right_open         SplomMarkerSymbol = "triangle-right-open"
	SplomMarkerSymbol208                          SplomMarkerSymbol = "208"
	SplomMarkerSymbol_triangle_right_dot          SplomMarkerSymbol = "triangle-right-dot"
	SplomMarkerSymbol308                          SplomMarkerSymbol = "308"
	SplomMarkerSymbol_triangle_right_open_dot     SplomMarkerSymbol = "triangle-right-open-dot"
	SplomMarkerSymbol9                            SplomMarkerSymbol = "9"
	SplomMarkerSymbol_triangle_ne                 SplomMarkerSymbol = "triangle-ne"
	SplomMarkerSymbol109                          SplomMarkerSymbol = "109"
	SplomMarkerSymbol_triangle_ne_open            SplomMarkerSymbol = "triangle-ne-open"
	SplomMarkerSymbol209                          SplomMarkerSymbol = "209"
	SplomMarkerSymbol_triangle_ne_dot             SplomMarkerSymbol = "triangle-ne-dot"
	SplomMarkerSymbol309                          SplomMarkerSymbol = "309"
	SplomMarkerSymbol_triangle_ne_open_dot        SplomMarkerSymbol = "triangle-ne-open-dot"
	SplomMarkerSymbol10                           SplomMarkerSymbol = "10"
	SplomMarkerSymbol_triangle_se                 SplomMarkerSymbol = "triangle-se"
	SplomMarkerSymbol110                          SplomMarkerSymbol = "110"
	SplomMarkerSymbol_triangle_se_open            SplomMarkerSymbol = "triangle-se-open"
	SplomMarkerSymbol210                          SplomMarkerSymbol = "210"
	SplomMarkerSymbol_triangle_se_dot             SplomMarkerSymbol = "triangle-se-dot"
	SplomMarkerSymbol310                          SplomMarkerSymbol = "310"
	SplomMarkerSymbol_triangle_se_open_dot        SplomMarkerSymbol = "triangle-se-open-dot"
	SplomMarkerSymbol11                           SplomMarkerSymbol = "11"
	SplomMarkerSymbol_triangle_sw                 SplomMarkerSymbol = "triangle-sw"
	SplomMarkerSymbol111                          SplomMarkerSymbol = "111"
	SplomMarkerSymbol_triangle_sw_open            SplomMarkerSymbol = "triangle-sw-open"
	SplomMarkerSymbol211                          SplomMarkerSymbol = "211"
	SplomMarkerSymbol_triangle_sw_dot             SplomMarkerSymbol = "triangle-sw-dot"
	SplomMarkerSymbol311                          SplomMarkerSymbol = "311"
	SplomMarkerSymbol_triangle_sw_open_dot        SplomMarkerSymbol = "triangle-sw-open-dot"
	SplomMarkerSymbol12                           SplomMarkerSymbol = "12"
	SplomMarkerSymbol_triangle_nw                 SplomMarkerSymbol = "triangle-nw"
	SplomMarkerSymbol112                          SplomMarkerSymbol = "112"
	SplomMarkerSymbol_triangle_nw_open            SplomMarkerSymbol = "triangle-nw-open"
	SplomMarkerSymbol212                          SplomMarkerSymbol = "212"
	SplomMarkerSymbol_triangle_nw_dot             SplomMarkerSymbol = "triangle-nw-dot"
	SplomMarkerSymbol312                          SplomMarkerSymbol = "312"
	SplomMarkerSymbol_triangle_nw_open_dot        SplomMarkerSymbol = "triangle-nw-open-dot"
	SplomMarkerSymbol13                           SplomMarkerSymbol = "13"
	SplomMarkerSymbol_pentagon                    SplomMarkerSymbol = "pentagon"
	SplomMarkerSymbol113                          SplomMarkerSymbol = "113"
	SplomMarkerSymbol_pentagon_open               SplomMarkerSymbol = "pentagon-open"
	SplomMarkerSymbol213                          SplomMarkerSymbol = "213"
	SplomMarkerSymbol_pentagon_dot                SplomMarkerSymbol = "pentagon-dot"
	SplomMarkerSymbol313                          SplomMarkerSymbol = "313"
	SplomMarkerSymbol_pentagon_open_dot           SplomMarkerSymbol = "pentagon-open-dot"
	SplomMarkerSymbol14                           SplomMarkerSymbol = "14"
	SplomMarkerSymbol_hexagon                     SplomMarkerSymbol = "hexagon"
	SplomMarkerSymbol114                          SplomMarkerSymbol = "114"
	SplomMarkerSymbol_hexagon_open                SplomMarkerSymbol = "hexagon-open"
	SplomMarkerSymbol214                          SplomMarkerSymbol = "214"
	SplomMarkerSymbol_hexagon_dot                 SplomMarkerSymbol = "hexagon-dot"
	SplomMarkerSymbol314                          SplomMarkerSymbol = "314"
	SplomMarkerSymbol_hexagon_open_dot            SplomMarkerSymbol = "hexagon-open-dot"
	SplomMarkerSymbol15                           SplomMarkerSymbol = "15"
	SplomMarkerSymbol_hexagon2                    SplomMarkerSymbol = "hexagon2"
	SplomMarkerSymbol115                          SplomMarkerSymbol = "115"
	SplomMarkerSymbol_hexagon2_open               SplomMarkerSymbol = "hexagon2-open"
	SplomMarkerSymbol215                          SplomMarkerSymbol = "215"
	SplomMarkerSymbol_hexagon2_dot                SplomMarkerSymbol = "hexagon2-dot"
	SplomMarkerSymbol315                          SplomMarkerSymbol = "315"
	SplomMarkerSymbol_hexagon2_open_dot           SplomMarkerSymbol = "hexagon2-open-dot"
	SplomMarkerSymbol16                           SplomMarkerSymbol = "16"
	SplomMarkerSymbol_octagon                     SplomMarkerSymbol = "octagon"
	SplomMarkerSymbol116                          SplomMarkerSymbol = "116"
	SplomMarkerSymbol_octagon_open                SplomMarkerSymbol = "octagon-open"
	SplomMarkerSymbol216                          SplomMarkerSymbol = "216"
	SplomMarkerSymbol_octagon_dot                 SplomMarkerSymbol = "octagon-dot"
	SplomMarkerSymbol316                          SplomMarkerSymbol = "316"
	SplomMarkerSymbol_octagon_open_dot            SplomMarkerSymbol = "octagon-open-dot"
	SplomMarkerSymbol17                           SplomMarkerSymbol = "17"
	SplomMarkerSymbol_star                        SplomMarkerSymbol = "star"
	SplomMarkerSymbol117                          SplomMarkerSymbol = "117"
	SplomMarkerSymbol_star_open                   SplomMarkerSymbol = "star-open"
	SplomMarkerSymbol217                          SplomMarkerSymbol = "217"
	SplomMarkerSymbol_star_dot                    SplomMarkerSymbol = "star-dot"
	SplomMarkerSymbol317                          SplomMarkerSymbol = "317"
	SplomMarkerSymbol_star_open_dot               SplomMarkerSymbol = "star-open-dot"
	SplomMarkerSymbol18                           SplomMarkerSymbol = "18"
	SplomMarkerSymbol_hexagram                    SplomMarkerSymbol = "hexagram"
	SplomMarkerSymbol118                          SplomMarkerSymbol = "118"
	SplomMarkerSymbol_hexagram_open               SplomMarkerSymbol = "hexagram-open"
	SplomMarkerSymbol218                          SplomMarkerSymbol = "218"
	SplomMarkerSymbol_hexagram_dot                SplomMarkerSymbol = "hexagram-dot"
	SplomMarkerSymbol318                          SplomMarkerSymbol = "318"
	SplomMarkerSymbol_hexagram_open_dot           SplomMarkerSymbol = "hexagram-open-dot"
	SplomMarkerSymbol19                           SplomMarkerSymbol = "19"
	SplomMarkerSymbol_star_triangle_up            SplomMarkerSymbol = "star-triangle-up"
	SplomMarkerSymbol119                          SplomMarkerSymbol = "119"
	SplomMarkerSymbol_star_triangle_up_open       SplomMarkerSymbol = "star-triangle-up-open"
	SplomMarkerSymbol219                          SplomMarkerSymbol = "219"
	SplomMarkerSymbol_star_triangle_up_dot        SplomMarkerSymbol = "star-triangle-up-dot"
	SplomMarkerSymbol319                          SplomMarkerSymbol = "319"
	SplomMarkerSymbol_star_triangle_up_open_dot   SplomMarkerSymbol = "star-triangle-up-open-dot"
	SplomMarkerSymbol20                           SplomMarkerSymbol = "20"
	SplomMarkerSymbol_star_triangle_down          SplomMarkerSymbol = "star-triangle-down"
	SplomMarkerSymbol120                          SplomMarkerSymbol = "120"
	SplomMarkerSymbol_star_triangle_down_open     SplomMarkerSymbol = "star-triangle-down-open"
	SplomMarkerSymbol220                          SplomMarkerSymbol = "220"
	SplomMarkerSymbol_star_triangle_down_dot      SplomMarkerSymbol = "star-triangle-down-dot"
	SplomMarkerSymbol320                          SplomMarkerSymbol = "320"
	SplomMarkerSymbol_star_triangle_down_open_dot SplomMarkerSymbol = "star-triangle-down-open-dot"
	SplomMarkerSymbol21                           SplomMarkerSymbol = "21"
	SplomMarkerSymbol_star_square                 SplomMarkerSymbol = "star-square"
	SplomMarkerSymbol121                          SplomMarkerSymbol = "121"
	SplomMarkerSymbol_star_square_open            SplomMarkerSymbol = "star-square-open"
	SplomMarkerSymbol221                          SplomMarkerSymbol = "221"
	SplomMarkerSymbol_star_square_dot             SplomMarkerSymbol = "star-square-dot"
	SplomMarkerSymbol321                          SplomMarkerSymbol = "321"
	SplomMarkerSymbol_star_square_open_dot        SplomMarkerSymbol = "star-square-open-dot"
	SplomMarkerSymbol22                           SplomMarkerSymbol = "22"
	SplomMarkerSymbol_star_diamond                SplomMarkerSymbol = "star-diamond"
	SplomMarkerSymbol122                          SplomMarkerSymbol = "122"
	SplomMarkerSymbol_star_diamond_open           SplomMarkerSymbol = "star-diamond-open"
	SplomMarkerSymbol222                          SplomMarkerSymbol = "222"
	SplomMarkerSymbol_star_diamond_dot            SplomMarkerSymbol = "star-diamond-dot"
	SplomMarkerSymbol322                          SplomMarkerSymbol = "322"
	SplomMarkerSymbol_star_diamond_open_dot       SplomMarkerSymbol = "star-diamond-open-dot"
	SplomMarkerSymbol23                           SplomMarkerSymbol = "23"
	SplomMarkerSymbol_diamond_tall                SplomMarkerSymbol = "diamond-tall"
	SplomMarkerSymbol123                          SplomMarkerSymbol = "123"
	SplomMarkerSymbol_diamond_tall_open           SplomMarkerSymbol = "diamond-tall-open"
	SplomMarkerSymbol223                          SplomMarkerSymbol = "223"
	SplomMarkerSymbol_diamond_tall_dot            SplomMarkerSymbol = "diamond-tall-dot"
	SplomMarkerSymbol323                          SplomMarkerSymbol = "323"
	SplomMarkerSymbol_diamond_tall_open_dot       SplomMarkerSymbol = "diamond-tall-open-dot"
	SplomMarkerSymbol24                           SplomMarkerSymbol = "24"
	SplomMarkerSymbol_diamond_wide                SplomMarkerSymbol = "diamond-wide"
	SplomMarkerSymbol124                          SplomMarkerSymbol = "124"
	SplomMarkerSymbol_diamond_wide_open           SplomMarkerSymbol = "diamond-wide-open"
	SplomMarkerSymbol224                          SplomMarkerSymbol = "224"
	SplomMarkerSymbol_diamond_wide_dot            SplomMarkerSymbol = "diamond-wide-dot"
	SplomMarkerSymbol324                          SplomMarkerSymbol = "324"
	SplomMarkerSymbol_diamond_wide_open_dot       SplomMarkerSymbol = "diamond-wide-open-dot"
	SplomMarkerSymbol25                           SplomMarkerSymbol = "25"
	SplomMarkerSymbol_hourglass                   SplomMarkerSymbol = "hourglass"
	SplomMarkerSymbol125                          SplomMarkerSymbol = "125"
	SplomMarkerSymbol_hourglass_open              SplomMarkerSymbol = "hourglass-open"
	SplomMarkerSymbol26                           SplomMarkerSymbol = "26"
	SplomMarkerSymbol_bowtie                      SplomMarkerSymbol = "bowtie"
	SplomMarkerSymbol126                          SplomMarkerSymbol = "126"
	SplomMarkerSymbol_bowtie_open                 SplomMarkerSymbol = "bowtie-open"
	SplomMarkerSymbol27                           SplomMarkerSymbol = "27"
	SplomMarkerSymbol_circle_cross                SplomMarkerSymbol = "circle-cross"
	SplomMarkerSymbol127                          SplomMarkerSymbol = "127"
	SplomMarkerSymbol_circle_cross_open           SplomMarkerSymbol = "circle-cross-open"
	SplomMarkerSymbol28                           SplomMarkerSymbol = "28"
	SplomMarkerSymbol_circle_x                    SplomMarkerSymbol = "circle-x"
	SplomMarkerSymbol128                          SplomMarkerSymbol = "128"
	SplomMarkerSymbol_circle_x_open               SplomMarkerSymbol = "circle-x-open"
	SplomMarkerSymbol29                           SplomMarkerSymbol = "29"
	SplomMarkerSymbol_square_cross                SplomMarkerSymbol = "square-cross"
	SplomMarkerSymbol129                          SplomMarkerSymbol = "129"
	SplomMarkerSymbol_square_cross_open           SplomMarkerSymbol = "square-cross-open"
	SplomMarkerSymbol30                           SplomMarkerSymbol = "30"
	SplomMarkerSymbol_square_x                    SplomMarkerSymbol = "square-x"
	SplomMarkerSymbol130                          SplomMarkerSymbol = "130"
	SplomMarkerSymbol_square_x_open               SplomMarkerSymbol = "square-x-open"
	SplomMarkerSymbol31                           SplomMarkerSymbol = "31"
	SplomMarkerSymbol_diamond_cross               SplomMarkerSymbol = "diamond-cross"
	SplomMarkerSymbol131                          SplomMarkerSymbol = "131"
	SplomMarkerSymbol_diamond_cross_open          SplomMarkerSymbol = "diamond-cross-open"
	SplomMarkerSymbol32                           SplomMarkerSymbol = "32"
	SplomMarkerSymbol_diamond_x                   SplomMarkerSymbol = "diamond-x"
	SplomMarkerSymbol132                          SplomMarkerSymbol = "132"
	SplomMarkerSymbol_diamond_x_open              SplomMarkerSymbol = "diamond-x-open"
	SplomMarkerSymbol33                           SplomMarkerSymbol = "33"
	SplomMarkerSymbol_cross_thin                  SplomMarkerSymbol = "cross-thin"
	SplomMarkerSymbol133                          SplomMarkerSymbol = "133"
	SplomMarkerSymbol_cross_thin_open             SplomMarkerSymbol = "cross-thin-open"
	SplomMarkerSymbol34                           SplomMarkerSymbol = "34"
	SplomMarkerSymbol_x_thin                      SplomMarkerSymbol = "x-thin"
	SplomMarkerSymbol134                          SplomMarkerSymbol = "134"
	SplomMarkerSymbol_x_thin_open                 SplomMarkerSymbol = "x-thin-open"
	SplomMarkerSymbol35                           SplomMarkerSymbol = "35"
	SplomMarkerSymbol_asterisk                    SplomMarkerSymbol = "asterisk"
	SplomMarkerSymbol135                          SplomMarkerSymbol = "135"
	SplomMarkerSymbol_asterisk_open               SplomMarkerSymbol = "asterisk-open"
	SplomMarkerSymbol36                           SplomMarkerSymbol = "36"
	SplomMarkerSymbol_hash                        SplomMarkerSymbol = "hash"
	SplomMarkerSymbol136                          SplomMarkerSymbol = "136"
	SplomMarkerSymbol_hash_open                   SplomMarkerSymbol = "hash-open"
	SplomMarkerSymbol236                          SplomMarkerSymbol = "236"
	SplomMarkerSymbol_hash_dot                    SplomMarkerSymbol = "hash-dot"
	SplomMarkerSymbol336                          SplomMarkerSymbol = "336"
	SplomMarkerSymbol_hash_open_dot               SplomMarkerSymbol = "hash-open-dot"
	SplomMarkerSymbol37                           SplomMarkerSymbol = "37"
	SplomMarkerSymbol_y_up                        SplomMarkerSymbol = "y-up"
	SplomMarkerSymbol137                          SplomMarkerSymbol = "137"
	SplomMarkerSymbol_y_up_open                   SplomMarkerSymbol = "y-up-open"
	SplomMarkerSymbol38                           SplomMarkerSymbol = "38"
	SplomMarkerSymbol_y_down                      SplomMarkerSymbol = "y-down"
	SplomMarkerSymbol138                          SplomMarkerSymbol = "138"
	SplomMarkerSymbol_y_down_open                 SplomMarkerSymbol = "y-down-open"
	SplomMarkerSymbol39                           SplomMarkerSymbol = "39"
	SplomMarkerSymbol_y_left                      SplomMarkerSymbol = "y-left"
	SplomMarkerSymbol139                          SplomMarkerSymbol = "139"
	SplomMarkerSymbol_y_left_open                 SplomMarkerSymbol = "y-left-open"
	SplomMarkerSymbol40                           SplomMarkerSymbol = "40"
	SplomMarkerSymbol_y_right                     SplomMarkerSymbol = "y-right"
	SplomMarkerSymbol140                          SplomMarkerSymbol = "140"
	SplomMarkerSymbol_y_right_open                SplomMarkerSymbol = "y-right-open"
	SplomMarkerSymbol41                           SplomMarkerSymbol = "41"
	SplomMarkerSymbol_line_ew                     SplomMarkerSymbol = "line-ew"
	SplomMarkerSymbol141                          SplomMarkerSymbol = "141"
	SplomMarkerSymbol_line_ew_open                SplomMarkerSymbol = "line-ew-open"
	SplomMarkerSymbol42                           SplomMarkerSymbol = "42"
	SplomMarkerSymbol_line_ns                     SplomMarkerSymbol = "line-ns"
	SplomMarkerSymbol142                          SplomMarkerSymbol = "142"
	SplomMarkerSymbol_line_ns_open                SplomMarkerSymbol = "line-ns-open"
	SplomMarkerSymbol43                           SplomMarkerSymbol = "43"
	SplomMarkerSymbol_line_ne                     SplomMarkerSymbol = "line-ne"
	SplomMarkerSymbol143                          SplomMarkerSymbol = "143"
	SplomMarkerSymbol_line_ne_open                SplomMarkerSymbol = "line-ne-open"
	SplomMarkerSymbol44                           SplomMarkerSymbol = "44"
	SplomMarkerSymbol_line_nw                     SplomMarkerSymbol = "line-nw"
	SplomMarkerSymbol144                          SplomMarkerSymbol = "144"
	SplomMarkerSymbol_line_nw_open                SplomMarkerSymbol = "line-nw-open"
)

// SplomVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type SplomVisible interface{}

var (
	SplomVisible_True       SplomVisible = true
	SplomVisible_False      SplomVisible = false
	SplomVisible_legendonly SplomVisible = "legendonly"
)

// StreamtubeColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type StreamtubeColorbarExponentformat string

const (
	StreamtubeColorbarExponentformat_none  StreamtubeColorbarExponentformat = "none"
	StreamtubeColorbarExponentformat_e     StreamtubeColorbarExponentformat = "e"
	StreamtubeColorbarExponentformat_E     StreamtubeColorbarExponentformat = "E"
	StreamtubeColorbarExponentformat_power StreamtubeColorbarExponentformat = "power"
	StreamtubeColorbarExponentformat_SI    StreamtubeColorbarExponentformat = "SI"
	StreamtubeColorbarExponentformat_B     StreamtubeColorbarExponentformat = "B"
)

// StreamtubeColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type StreamtubeColorbarLenmode string

const (
	StreamtubeColorbarLenmode_fraction StreamtubeColorbarLenmode = "fraction"
	StreamtubeColorbarLenmode_pixels   StreamtubeColorbarLenmode = "pixels"
)

// StreamtubeColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type StreamtubeColorbarShowexponent string

const (
	StreamtubeColorbarShowexponent_all   StreamtubeColorbarShowexponent = "all"
	StreamtubeColorbarShowexponent_first StreamtubeColorbarShowexponent = "first"
	StreamtubeColorbarShowexponent_last  StreamtubeColorbarShowexponent = "last"
	StreamtubeColorbarShowexponent_none  StreamtubeColorbarShowexponent = "none"
)

// StreamtubeColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type StreamtubeColorbarShowtickprefix string

const (
	StreamtubeColorbarShowtickprefix_all   StreamtubeColorbarShowtickprefix = "all"
	StreamtubeColorbarShowtickprefix_first StreamtubeColorbarShowtickprefix = "first"
	StreamtubeColorbarShowtickprefix_last  StreamtubeColorbarShowtickprefix = "last"
	StreamtubeColorbarShowtickprefix_none  StreamtubeColorbarShowtickprefix = "none"
)

// StreamtubeColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type StreamtubeColorbarShowticksuffix string

const (
	StreamtubeColorbarShowticksuffix_all   StreamtubeColorbarShowticksuffix = "all"
	StreamtubeColorbarShowticksuffix_first StreamtubeColorbarShowticksuffix = "first"
	StreamtubeColorbarShowticksuffix_last  StreamtubeColorbarShowticksuffix = "last"
	StreamtubeColorbarShowticksuffix_none  StreamtubeColorbarShowticksuffix = "none"
)

// StreamtubeColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type StreamtubeColorbarThicknessmode string

const (
	StreamtubeColorbarThicknessmode_fraction StreamtubeColorbarThicknessmode = "fraction"
	StreamtubeColorbarThicknessmode_pixels   StreamtubeColorbarThicknessmode = "pixels"
)

// StreamtubeColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type StreamtubeColorbarTickmode string

const (
	StreamtubeColorbarTickmode_auto   StreamtubeColorbarTickmode = "auto"
	StreamtubeColorbarTickmode_linear StreamtubeColorbarTickmode = "linear"
	StreamtubeColorbarTickmode_array  StreamtubeColorbarTickmode = "array"
)

// StreamtubeColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type StreamtubeColorbarTicks string

const (
	StreamtubeColorbarTicks_outside StreamtubeColorbarTicks = "outside"
	StreamtubeColorbarTicks_inside  StreamtubeColorbarTicks = "inside"
)

// StreamtubeColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type StreamtubeColorbarTitleSide string

const (
	StreamtubeColorbarTitleSide_right  StreamtubeColorbarTitleSide = "right"
	StreamtubeColorbarTitleSide_top    StreamtubeColorbarTitleSide = "top"
	StreamtubeColorbarTitleSide_bottom StreamtubeColorbarTitleSide = "bottom"
)

// StreamtubeColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type StreamtubeColorbarXanchor string

const (
	StreamtubeColorbarXanchor_left   StreamtubeColorbarXanchor = "left"
	StreamtubeColorbarXanchor_center StreamtubeColorbarXanchor = "center"
	StreamtubeColorbarXanchor_right  StreamtubeColorbarXanchor = "right"
)

// StreamtubeColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type StreamtubeColorbarYanchor string

const (
	StreamtubeColorbarYanchor_top    StreamtubeColorbarYanchor = "top"
	StreamtubeColorbarYanchor_middle StreamtubeColorbarYanchor = "middle"
	StreamtubeColorbarYanchor_bottom StreamtubeColorbarYanchor = "bottom"
)

// StreamtubeHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type StreamtubeHoverlabelAlign string

const (
	StreamtubeHoverlabelAlign_left  StreamtubeHoverlabelAlign = "left"
	StreamtubeHoverlabelAlign_right StreamtubeHoverlabelAlign = "right"
	StreamtubeHoverlabelAlign_auto  StreamtubeHoverlabelAlign = "auto"
)

// StreamtubeVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type StreamtubeVisible interface{}

var (
	StreamtubeVisible_True       StreamtubeVisible = true
	StreamtubeVisible_False      StreamtubeVisible = false
	StreamtubeVisible_legendonly StreamtubeVisible = "legendonly"
)

// SunburstBranchvalues Determines how the items in `values` are summed. When set to *total*, items in `values` are taken to be value of all its descendants. When set to *remainder*, items in `values` corresponding to the root and the branches sectors are taken to be the extra part not part of the sum of the values at their leaves.
type SunburstBranchvalues string

const (
	SunburstBranchvalues_remainder SunburstBranchvalues = "remainder"
	SunburstBranchvalues_total     SunburstBranchvalues = "total"
)

// SunburstHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SunburstHoverlabelAlign string

const (
	SunburstHoverlabelAlign_left  SunburstHoverlabelAlign = "left"
	SunburstHoverlabelAlign_right SunburstHoverlabelAlign = "right"
	SunburstHoverlabelAlign_auto  SunburstHoverlabelAlign = "auto"
)

// SunburstInsidetextorientation Controls the orientation of the text inside chart sectors. When set to *auto*, text may be oriented in any direction in order to be as big as possible in the middle of a sector. The *horizontal* option orients text to be parallel with the bottom of the chart, and may make text smaller in order to achieve that goal. The *radial* option orients text along the radius of the sector. The *tangential* option orients text perpendicular to the radius of the sector.
type SunburstInsidetextorientation string

const (
	SunburstInsidetextorientation_horizontal SunburstInsidetextorientation = "horizontal"
	SunburstInsidetextorientation_radial     SunburstInsidetextorientation = "radial"
	SunburstInsidetextorientation_tangential SunburstInsidetextorientation = "tangential"
	SunburstInsidetextorientation_auto       SunburstInsidetextorientation = "auto"
)

// SunburstMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type SunburstMarkerColorbarExponentformat string

const (
	SunburstMarkerColorbarExponentformat_none  SunburstMarkerColorbarExponentformat = "none"
	SunburstMarkerColorbarExponentformat_e     SunburstMarkerColorbarExponentformat = "e"
	SunburstMarkerColorbarExponentformat_E     SunburstMarkerColorbarExponentformat = "E"
	SunburstMarkerColorbarExponentformat_power SunburstMarkerColorbarExponentformat = "power"
	SunburstMarkerColorbarExponentformat_SI    SunburstMarkerColorbarExponentformat = "SI"
	SunburstMarkerColorbarExponentformat_B     SunburstMarkerColorbarExponentformat = "B"
)

// SunburstMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type SunburstMarkerColorbarLenmode string

const (
	SunburstMarkerColorbarLenmode_fraction SunburstMarkerColorbarLenmode = "fraction"
	SunburstMarkerColorbarLenmode_pixels   SunburstMarkerColorbarLenmode = "pixels"
)

// SunburstMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type SunburstMarkerColorbarShowexponent string

const (
	SunburstMarkerColorbarShowexponent_all   SunburstMarkerColorbarShowexponent = "all"
	SunburstMarkerColorbarShowexponent_first SunburstMarkerColorbarShowexponent = "first"
	SunburstMarkerColorbarShowexponent_last  SunburstMarkerColorbarShowexponent = "last"
	SunburstMarkerColorbarShowexponent_none  SunburstMarkerColorbarShowexponent = "none"
)

// SunburstMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type SunburstMarkerColorbarShowtickprefix string

const (
	SunburstMarkerColorbarShowtickprefix_all   SunburstMarkerColorbarShowtickprefix = "all"
	SunburstMarkerColorbarShowtickprefix_first SunburstMarkerColorbarShowtickprefix = "first"
	SunburstMarkerColorbarShowtickprefix_last  SunburstMarkerColorbarShowtickprefix = "last"
	SunburstMarkerColorbarShowtickprefix_none  SunburstMarkerColorbarShowtickprefix = "none"
)

// SunburstMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type SunburstMarkerColorbarShowticksuffix string

const (
	SunburstMarkerColorbarShowticksuffix_all   SunburstMarkerColorbarShowticksuffix = "all"
	SunburstMarkerColorbarShowticksuffix_first SunburstMarkerColorbarShowticksuffix = "first"
	SunburstMarkerColorbarShowticksuffix_last  SunburstMarkerColorbarShowticksuffix = "last"
	SunburstMarkerColorbarShowticksuffix_none  SunburstMarkerColorbarShowticksuffix = "none"
)

// SunburstMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type SunburstMarkerColorbarThicknessmode string

const (
	SunburstMarkerColorbarThicknessmode_fraction SunburstMarkerColorbarThicknessmode = "fraction"
	SunburstMarkerColorbarThicknessmode_pixels   SunburstMarkerColorbarThicknessmode = "pixels"
)

// SunburstMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type SunburstMarkerColorbarTickmode string

const (
	SunburstMarkerColorbarTickmode_auto   SunburstMarkerColorbarTickmode = "auto"
	SunburstMarkerColorbarTickmode_linear SunburstMarkerColorbarTickmode = "linear"
	SunburstMarkerColorbarTickmode_array  SunburstMarkerColorbarTickmode = "array"
)

// SunburstMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type SunburstMarkerColorbarTicks string

const (
	SunburstMarkerColorbarTicks_outside SunburstMarkerColorbarTicks = "outside"
	SunburstMarkerColorbarTicks_inside  SunburstMarkerColorbarTicks = "inside"
)

// SunburstMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type SunburstMarkerColorbarTitleSide string

const (
	SunburstMarkerColorbarTitleSide_right  SunburstMarkerColorbarTitleSide = "right"
	SunburstMarkerColorbarTitleSide_top    SunburstMarkerColorbarTitleSide = "top"
	SunburstMarkerColorbarTitleSide_bottom SunburstMarkerColorbarTitleSide = "bottom"
)

// SunburstMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type SunburstMarkerColorbarXanchor string

const (
	SunburstMarkerColorbarXanchor_left   SunburstMarkerColorbarXanchor = "left"
	SunburstMarkerColorbarXanchor_center SunburstMarkerColorbarXanchor = "center"
	SunburstMarkerColorbarXanchor_right  SunburstMarkerColorbarXanchor = "right"
)

// SunburstMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type SunburstMarkerColorbarYanchor string

const (
	SunburstMarkerColorbarYanchor_top    SunburstMarkerColorbarYanchor = "top"
	SunburstMarkerColorbarYanchor_middle SunburstMarkerColorbarYanchor = "middle"
	SunburstMarkerColorbarYanchor_bottom SunburstMarkerColorbarYanchor = "bottom"
)

// SunburstVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type SunburstVisible interface{}

var (
	SunburstVisible_True       SunburstVisible = true
	SunburstVisible_False      SunburstVisible = false
	SunburstVisible_legendonly SunburstVisible = "legendonly"
)

// SurfaceColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type SurfaceColorbarExponentformat string

const (
	SurfaceColorbarExponentformat_none  SurfaceColorbarExponentformat = "none"
	SurfaceColorbarExponentformat_e     SurfaceColorbarExponentformat = "e"
	SurfaceColorbarExponentformat_E     SurfaceColorbarExponentformat = "E"
	SurfaceColorbarExponentformat_power SurfaceColorbarExponentformat = "power"
	SurfaceColorbarExponentformat_SI    SurfaceColorbarExponentformat = "SI"
	SurfaceColorbarExponentformat_B     SurfaceColorbarExponentformat = "B"
)

// SurfaceColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type SurfaceColorbarLenmode string

const (
	SurfaceColorbarLenmode_fraction SurfaceColorbarLenmode = "fraction"
	SurfaceColorbarLenmode_pixels   SurfaceColorbarLenmode = "pixels"
)

// SurfaceColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type SurfaceColorbarShowexponent string

const (
	SurfaceColorbarShowexponent_all   SurfaceColorbarShowexponent = "all"
	SurfaceColorbarShowexponent_first SurfaceColorbarShowexponent = "first"
	SurfaceColorbarShowexponent_last  SurfaceColorbarShowexponent = "last"
	SurfaceColorbarShowexponent_none  SurfaceColorbarShowexponent = "none"
)

// SurfaceColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type SurfaceColorbarShowtickprefix string

const (
	SurfaceColorbarShowtickprefix_all   SurfaceColorbarShowtickprefix = "all"
	SurfaceColorbarShowtickprefix_first SurfaceColorbarShowtickprefix = "first"
	SurfaceColorbarShowtickprefix_last  SurfaceColorbarShowtickprefix = "last"
	SurfaceColorbarShowtickprefix_none  SurfaceColorbarShowtickprefix = "none"
)

// SurfaceColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type SurfaceColorbarShowticksuffix string

const (
	SurfaceColorbarShowticksuffix_all   SurfaceColorbarShowticksuffix = "all"
	SurfaceColorbarShowticksuffix_first SurfaceColorbarShowticksuffix = "first"
	SurfaceColorbarShowticksuffix_last  SurfaceColorbarShowticksuffix = "last"
	SurfaceColorbarShowticksuffix_none  SurfaceColorbarShowticksuffix = "none"
)

// SurfaceColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type SurfaceColorbarThicknessmode string

const (
	SurfaceColorbarThicknessmode_fraction SurfaceColorbarThicknessmode = "fraction"
	SurfaceColorbarThicknessmode_pixels   SurfaceColorbarThicknessmode = "pixels"
)

// SurfaceColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type SurfaceColorbarTickmode string

const (
	SurfaceColorbarTickmode_auto   SurfaceColorbarTickmode = "auto"
	SurfaceColorbarTickmode_linear SurfaceColorbarTickmode = "linear"
	SurfaceColorbarTickmode_array  SurfaceColorbarTickmode = "array"
)

// SurfaceColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type SurfaceColorbarTicks string

const (
	SurfaceColorbarTicks_outside SurfaceColorbarTicks = "outside"
	SurfaceColorbarTicks_inside  SurfaceColorbarTicks = "inside"
)

// SurfaceColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type SurfaceColorbarTitleSide string

const (
	SurfaceColorbarTitleSide_right  SurfaceColorbarTitleSide = "right"
	SurfaceColorbarTitleSide_top    SurfaceColorbarTitleSide = "top"
	SurfaceColorbarTitleSide_bottom SurfaceColorbarTitleSide = "bottom"
)

// SurfaceColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type SurfaceColorbarXanchor string

const (
	SurfaceColorbarXanchor_left   SurfaceColorbarXanchor = "left"
	SurfaceColorbarXanchor_center SurfaceColorbarXanchor = "center"
	SurfaceColorbarXanchor_right  SurfaceColorbarXanchor = "right"
)

// SurfaceColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type SurfaceColorbarYanchor string

const (
	SurfaceColorbarYanchor_top    SurfaceColorbarYanchor = "top"
	SurfaceColorbarYanchor_middle SurfaceColorbarYanchor = "middle"
	SurfaceColorbarYanchor_bottom SurfaceColorbarYanchor = "bottom"
)

// SurfaceHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SurfaceHoverlabelAlign string

const (
	SurfaceHoverlabelAlign_left  SurfaceHoverlabelAlign = "left"
	SurfaceHoverlabelAlign_right SurfaceHoverlabelAlign = "right"
	SurfaceHoverlabelAlign_auto  SurfaceHoverlabelAlign = "auto"
)

// SurfaceVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type SurfaceVisible interface{}

var (
	SurfaceVisible_True       SurfaceVisible = true
	SurfaceVisible_False      SurfaceVisible = false
	SurfaceVisible_legendonly SurfaceVisible = "legendonly"
)

// SurfaceXcalendar Sets the calendar system to use with `x` date data.
type SurfaceXcalendar string

const (
	SurfaceXcalendar_gregorian  SurfaceXcalendar = "gregorian"
	SurfaceXcalendar_chinese    SurfaceXcalendar = "chinese"
	SurfaceXcalendar_coptic     SurfaceXcalendar = "coptic"
	SurfaceXcalendar_discworld  SurfaceXcalendar = "discworld"
	SurfaceXcalendar_ethiopian  SurfaceXcalendar = "ethiopian"
	SurfaceXcalendar_hebrew     SurfaceXcalendar = "hebrew"
	SurfaceXcalendar_islamic    SurfaceXcalendar = "islamic"
	SurfaceXcalendar_julian     SurfaceXcalendar = "julian"
	SurfaceXcalendar_mayan      SurfaceXcalendar = "mayan"
	SurfaceXcalendar_nanakshahi SurfaceXcalendar = "nanakshahi"
	SurfaceXcalendar_nepali     SurfaceXcalendar = "nepali"
	SurfaceXcalendar_persian    SurfaceXcalendar = "persian"
	SurfaceXcalendar_jalali     SurfaceXcalendar = "jalali"
	SurfaceXcalendar_taiwan     SurfaceXcalendar = "taiwan"
	SurfaceXcalendar_thai       SurfaceXcalendar = "thai"
	SurfaceXcalendar_ummalqura  SurfaceXcalendar = "ummalqura"
)

// SurfaceYcalendar Sets the calendar system to use with `y` date data.
type SurfaceYcalendar string

const (
	SurfaceYcalendar_gregorian  SurfaceYcalendar = "gregorian"
	SurfaceYcalendar_chinese    SurfaceYcalendar = "chinese"
	SurfaceYcalendar_coptic     SurfaceYcalendar = "coptic"
	SurfaceYcalendar_discworld  SurfaceYcalendar = "discworld"
	SurfaceYcalendar_ethiopian  SurfaceYcalendar = "ethiopian"
	SurfaceYcalendar_hebrew     SurfaceYcalendar = "hebrew"
	SurfaceYcalendar_islamic    SurfaceYcalendar = "islamic"
	SurfaceYcalendar_julian     SurfaceYcalendar = "julian"
	SurfaceYcalendar_mayan      SurfaceYcalendar = "mayan"
	SurfaceYcalendar_nanakshahi SurfaceYcalendar = "nanakshahi"
	SurfaceYcalendar_nepali     SurfaceYcalendar = "nepali"
	SurfaceYcalendar_persian    SurfaceYcalendar = "persian"
	SurfaceYcalendar_jalali     SurfaceYcalendar = "jalali"
	SurfaceYcalendar_taiwan     SurfaceYcalendar = "taiwan"
	SurfaceYcalendar_thai       SurfaceYcalendar = "thai"
	SurfaceYcalendar_ummalqura  SurfaceYcalendar = "ummalqura"
)

// SurfaceZcalendar Sets the calendar system to use with `z` date data.
type SurfaceZcalendar string

const (
	SurfaceZcalendar_gregorian  SurfaceZcalendar = "gregorian"
	SurfaceZcalendar_chinese    SurfaceZcalendar = "chinese"
	SurfaceZcalendar_coptic     SurfaceZcalendar = "coptic"
	SurfaceZcalendar_discworld  SurfaceZcalendar = "discworld"
	SurfaceZcalendar_ethiopian  SurfaceZcalendar = "ethiopian"
	SurfaceZcalendar_hebrew     SurfaceZcalendar = "hebrew"
	SurfaceZcalendar_islamic    SurfaceZcalendar = "islamic"
	SurfaceZcalendar_julian     SurfaceZcalendar = "julian"
	SurfaceZcalendar_mayan      SurfaceZcalendar = "mayan"
	SurfaceZcalendar_nanakshahi SurfaceZcalendar = "nanakshahi"
	SurfaceZcalendar_nepali     SurfaceZcalendar = "nepali"
	SurfaceZcalendar_persian    SurfaceZcalendar = "persian"
	SurfaceZcalendar_jalali     SurfaceZcalendar = "jalali"
	SurfaceZcalendar_taiwan     SurfaceZcalendar = "taiwan"
	SurfaceZcalendar_thai       SurfaceZcalendar = "thai"
	SurfaceZcalendar_ummalqura  SurfaceZcalendar = "ummalqura"
)

// TableCellsAlign Sets the horizontal alignment of the `text` within the box. Has an effect only if `text` spans two or more lines (i.e. `text` contains one or more <br> HTML tags) or if an explicit width is set to override the text width.
type TableCellsAlign string

const (
	TableCellsAlign_left   TableCellsAlign = "left"
	TableCellsAlign_center TableCellsAlign = "center"
	TableCellsAlign_right  TableCellsAlign = "right"
)

// TableHeaderAlign Sets the horizontal alignment of the `text` within the box. Has an effect only if `text` spans two or more lines (i.e. `text` contains one or more <br> HTML tags) or if an explicit width is set to override the text width.
type TableHeaderAlign string

const (
	TableHeaderAlign_left   TableHeaderAlign = "left"
	TableHeaderAlign_center TableHeaderAlign = "center"
	TableHeaderAlign_right  TableHeaderAlign = "right"
)

// TableHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type TableHoverlabelAlign string

const (
	TableHoverlabelAlign_left  TableHoverlabelAlign = "left"
	TableHoverlabelAlign_right TableHoverlabelAlign = "right"
	TableHoverlabelAlign_auto  TableHoverlabelAlign = "auto"
)

// TableVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type TableVisible interface{}

var (
	TableVisible_True       TableVisible = true
	TableVisible_False      TableVisible = false
	TableVisible_legendonly TableVisible = "legendonly"
)

// TreemapBranchvalues Determines how the items in `values` are summed. When set to *total*, items in `values` are taken to be value of all its descendants. When set to *remainder*, items in `values` corresponding to the root and the branches sectors are taken to be the extra part not part of the sum of the values at their leaves.
type TreemapBranchvalues string

const (
	TreemapBranchvalues_remainder TreemapBranchvalues = "remainder"
	TreemapBranchvalues_total     TreemapBranchvalues = "total"
)

// TreemapHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type TreemapHoverlabelAlign string

const (
	TreemapHoverlabelAlign_left  TreemapHoverlabelAlign = "left"
	TreemapHoverlabelAlign_right TreemapHoverlabelAlign = "right"
	TreemapHoverlabelAlign_auto  TreemapHoverlabelAlign = "auto"
)

// TreemapMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type TreemapMarkerColorbarExponentformat string

const (
	TreemapMarkerColorbarExponentformat_none  TreemapMarkerColorbarExponentformat = "none"
	TreemapMarkerColorbarExponentformat_e     TreemapMarkerColorbarExponentformat = "e"
	TreemapMarkerColorbarExponentformat_E     TreemapMarkerColorbarExponentformat = "E"
	TreemapMarkerColorbarExponentformat_power TreemapMarkerColorbarExponentformat = "power"
	TreemapMarkerColorbarExponentformat_SI    TreemapMarkerColorbarExponentformat = "SI"
	TreemapMarkerColorbarExponentformat_B     TreemapMarkerColorbarExponentformat = "B"
)

// TreemapMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type TreemapMarkerColorbarLenmode string

const (
	TreemapMarkerColorbarLenmode_fraction TreemapMarkerColorbarLenmode = "fraction"
	TreemapMarkerColorbarLenmode_pixels   TreemapMarkerColorbarLenmode = "pixels"
)

// TreemapMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type TreemapMarkerColorbarShowexponent string

const (
	TreemapMarkerColorbarShowexponent_all   TreemapMarkerColorbarShowexponent = "all"
	TreemapMarkerColorbarShowexponent_first TreemapMarkerColorbarShowexponent = "first"
	TreemapMarkerColorbarShowexponent_last  TreemapMarkerColorbarShowexponent = "last"
	TreemapMarkerColorbarShowexponent_none  TreemapMarkerColorbarShowexponent = "none"
)

// TreemapMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type TreemapMarkerColorbarShowtickprefix string

const (
	TreemapMarkerColorbarShowtickprefix_all   TreemapMarkerColorbarShowtickprefix = "all"
	TreemapMarkerColorbarShowtickprefix_first TreemapMarkerColorbarShowtickprefix = "first"
	TreemapMarkerColorbarShowtickprefix_last  TreemapMarkerColorbarShowtickprefix = "last"
	TreemapMarkerColorbarShowtickprefix_none  TreemapMarkerColorbarShowtickprefix = "none"
)

// TreemapMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type TreemapMarkerColorbarShowticksuffix string

const (
	TreemapMarkerColorbarShowticksuffix_all   TreemapMarkerColorbarShowticksuffix = "all"
	TreemapMarkerColorbarShowticksuffix_first TreemapMarkerColorbarShowticksuffix = "first"
	TreemapMarkerColorbarShowticksuffix_last  TreemapMarkerColorbarShowticksuffix = "last"
	TreemapMarkerColorbarShowticksuffix_none  TreemapMarkerColorbarShowticksuffix = "none"
)

// TreemapMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type TreemapMarkerColorbarThicknessmode string

const (
	TreemapMarkerColorbarThicknessmode_fraction TreemapMarkerColorbarThicknessmode = "fraction"
	TreemapMarkerColorbarThicknessmode_pixels   TreemapMarkerColorbarThicknessmode = "pixels"
)

// TreemapMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type TreemapMarkerColorbarTickmode string

const (
	TreemapMarkerColorbarTickmode_auto   TreemapMarkerColorbarTickmode = "auto"
	TreemapMarkerColorbarTickmode_linear TreemapMarkerColorbarTickmode = "linear"
	TreemapMarkerColorbarTickmode_array  TreemapMarkerColorbarTickmode = "array"
)

// TreemapMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type TreemapMarkerColorbarTicks string

const (
	TreemapMarkerColorbarTicks_outside TreemapMarkerColorbarTicks = "outside"
	TreemapMarkerColorbarTicks_inside  TreemapMarkerColorbarTicks = "inside"
)

// TreemapMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type TreemapMarkerColorbarTitleSide string

const (
	TreemapMarkerColorbarTitleSide_right  TreemapMarkerColorbarTitleSide = "right"
	TreemapMarkerColorbarTitleSide_top    TreemapMarkerColorbarTitleSide = "top"
	TreemapMarkerColorbarTitleSide_bottom TreemapMarkerColorbarTitleSide = "bottom"
)

// TreemapMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type TreemapMarkerColorbarXanchor string

const (
	TreemapMarkerColorbarXanchor_left   TreemapMarkerColorbarXanchor = "left"
	TreemapMarkerColorbarXanchor_center TreemapMarkerColorbarXanchor = "center"
	TreemapMarkerColorbarXanchor_right  TreemapMarkerColorbarXanchor = "right"
)

// TreemapMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type TreemapMarkerColorbarYanchor string

const (
	TreemapMarkerColorbarYanchor_top    TreemapMarkerColorbarYanchor = "top"
	TreemapMarkerColorbarYanchor_middle TreemapMarkerColorbarYanchor = "middle"
	TreemapMarkerColorbarYanchor_bottom TreemapMarkerColorbarYanchor = "bottom"
)

// TreemapMarkerDepthfade Determines if the sector colors are faded towards the background from the leaves up to the headers. This option is unavailable when a `colorscale` is present, defaults to false when `marker.colors` is set, but otherwise defaults to true. When set to *reversed*, the fading direction is inverted, that is the top elements within hierarchy are drawn with fully saturated colors while the leaves are faded towards the background color.
type TreemapMarkerDepthfade interface{}

var (
	TreemapMarkerDepthfade_True     TreemapMarkerDepthfade = true
	TreemapMarkerDepthfade_False    TreemapMarkerDepthfade = false
	TreemapMarkerDepthfade_reversed TreemapMarkerDepthfade = "reversed"
)

// TreemapPathbarEdgeshape Determines which shape is used for edges between `barpath` labels.
type TreemapPathbarEdgeshape string

const (
	TreemapPathbarEdgeshape_gt              TreemapPathbarEdgeshape = ">"
	TreemapPathbarEdgeshape_lt              TreemapPathbarEdgeshape = "<"
	TreemapPathbarEdgeshape_or              TreemapPathbarEdgeshape = "|"
	TreemapPathbarEdgeshape_slash           TreemapPathbarEdgeshape = "/"
	TreemapPathbarEdgeshape_doublebackslash TreemapPathbarEdgeshape = "\\"
)

// TreemapPathbarSide Determines on which side of the the treemap the `pathbar` should be presented.
type TreemapPathbarSide string

const (
	TreemapPathbarSide_top    TreemapPathbarSide = "top"
	TreemapPathbarSide_bottom TreemapPathbarSide = "bottom"
)

// TreemapTextposition Sets the positions of the `text` elements.
type TreemapTextposition string

const (
	TreemapTextposition_topleft      TreemapTextposition = "top left"
	TreemapTextposition_topcenter    TreemapTextposition = "top center"
	TreemapTextposition_topright     TreemapTextposition = "top right"
	TreemapTextposition_middleleft   TreemapTextposition = "middle left"
	TreemapTextposition_middlecenter TreemapTextposition = "middle center"
	TreemapTextposition_middleright  TreemapTextposition = "middle right"
	TreemapTextposition_bottomleft   TreemapTextposition = "bottom left"
	TreemapTextposition_bottomcenter TreemapTextposition = "bottom center"
	TreemapTextposition_bottomright  TreemapTextposition = "bottom right"
)

// TreemapTilingPacking Determines d3 treemap solver. For more info please refer to https://github.com/d3/d3-hierarchy#treemap-tiling
type TreemapTilingPacking string

const (
	TreemapTilingPacking_squarify   TreemapTilingPacking = "squarify"
	TreemapTilingPacking_binary     TreemapTilingPacking = "binary"
	TreemapTilingPacking_dice       TreemapTilingPacking = "dice"
	TreemapTilingPacking_slice      TreemapTilingPacking = "slice"
	TreemapTilingPacking_slice_dice TreemapTilingPacking = "slice-dice"
	TreemapTilingPacking_dice_slice TreemapTilingPacking = "dice-slice"
)

// TreemapVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type TreemapVisible interface{}

var (
	TreemapVisible_True       TreemapVisible = true
	TreemapVisible_False      TreemapVisible = false
	TreemapVisible_legendonly TreemapVisible = "legendonly"
)

// ViolinHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ViolinHoverlabelAlign string

const (
	ViolinHoverlabelAlign_left  ViolinHoverlabelAlign = "left"
	ViolinHoverlabelAlign_right ViolinHoverlabelAlign = "right"
	ViolinHoverlabelAlign_auto  ViolinHoverlabelAlign = "auto"
)

// ViolinMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ViolinMarkerSymbol string

const (
	ViolinMarkerSymbol0                            ViolinMarkerSymbol = "0"
	ViolinMarkerSymbol_circle                      ViolinMarkerSymbol = "circle"
	ViolinMarkerSymbol100                          ViolinMarkerSymbol = "100"
	ViolinMarkerSymbol_circle_open                 ViolinMarkerSymbol = "circle-open"
	ViolinMarkerSymbol200                          ViolinMarkerSymbol = "200"
	ViolinMarkerSymbol_circle_dot                  ViolinMarkerSymbol = "circle-dot"
	ViolinMarkerSymbol300                          ViolinMarkerSymbol = "300"
	ViolinMarkerSymbol_circle_open_dot             ViolinMarkerSymbol = "circle-open-dot"
	ViolinMarkerSymbol1                            ViolinMarkerSymbol = "1"
	ViolinMarkerSymbol_square                      ViolinMarkerSymbol = "square"
	ViolinMarkerSymbol101                          ViolinMarkerSymbol = "101"
	ViolinMarkerSymbol_square_open                 ViolinMarkerSymbol = "square-open"
	ViolinMarkerSymbol201                          ViolinMarkerSymbol = "201"
	ViolinMarkerSymbol_square_dot                  ViolinMarkerSymbol = "square-dot"
	ViolinMarkerSymbol301                          ViolinMarkerSymbol = "301"
	ViolinMarkerSymbol_square_open_dot             ViolinMarkerSymbol = "square-open-dot"
	ViolinMarkerSymbol2                            ViolinMarkerSymbol = "2"
	ViolinMarkerSymbol_diamond                     ViolinMarkerSymbol = "diamond"
	ViolinMarkerSymbol102                          ViolinMarkerSymbol = "102"
	ViolinMarkerSymbol_diamond_open                ViolinMarkerSymbol = "diamond-open"
	ViolinMarkerSymbol202                          ViolinMarkerSymbol = "202"
	ViolinMarkerSymbol_diamond_dot                 ViolinMarkerSymbol = "diamond-dot"
	ViolinMarkerSymbol302                          ViolinMarkerSymbol = "302"
	ViolinMarkerSymbol_diamond_open_dot            ViolinMarkerSymbol = "diamond-open-dot"
	ViolinMarkerSymbol3                            ViolinMarkerSymbol = "3"
	ViolinMarkerSymbol_cross                       ViolinMarkerSymbol = "cross"
	ViolinMarkerSymbol103                          ViolinMarkerSymbol = "103"
	ViolinMarkerSymbol_cross_open                  ViolinMarkerSymbol = "cross-open"
	ViolinMarkerSymbol203                          ViolinMarkerSymbol = "203"
	ViolinMarkerSymbol_cross_dot                   ViolinMarkerSymbol = "cross-dot"
	ViolinMarkerSymbol303                          ViolinMarkerSymbol = "303"
	ViolinMarkerSymbol_cross_open_dot              ViolinMarkerSymbol = "cross-open-dot"
	ViolinMarkerSymbol4                            ViolinMarkerSymbol = "4"
	ViolinMarkerSymbol_x                           ViolinMarkerSymbol = "x"
	ViolinMarkerSymbol104                          ViolinMarkerSymbol = "104"
	ViolinMarkerSymbol_x_open                      ViolinMarkerSymbol = "x-open"
	ViolinMarkerSymbol204                          ViolinMarkerSymbol = "204"
	ViolinMarkerSymbol_x_dot                       ViolinMarkerSymbol = "x-dot"
	ViolinMarkerSymbol304                          ViolinMarkerSymbol = "304"
	ViolinMarkerSymbol_x_open_dot                  ViolinMarkerSymbol = "x-open-dot"
	ViolinMarkerSymbol5                            ViolinMarkerSymbol = "5"
	ViolinMarkerSymbol_triangle_up                 ViolinMarkerSymbol = "triangle-up"
	ViolinMarkerSymbol105                          ViolinMarkerSymbol = "105"
	ViolinMarkerSymbol_triangle_up_open            ViolinMarkerSymbol = "triangle-up-open"
	ViolinMarkerSymbol205                          ViolinMarkerSymbol = "205"
	ViolinMarkerSymbol_triangle_up_dot             ViolinMarkerSymbol = "triangle-up-dot"
	ViolinMarkerSymbol305                          ViolinMarkerSymbol = "305"
	ViolinMarkerSymbol_triangle_up_open_dot        ViolinMarkerSymbol = "triangle-up-open-dot"
	ViolinMarkerSymbol6                            ViolinMarkerSymbol = "6"
	ViolinMarkerSymbol_triangle_down               ViolinMarkerSymbol = "triangle-down"
	ViolinMarkerSymbol106                          ViolinMarkerSymbol = "106"
	ViolinMarkerSymbol_triangle_down_open          ViolinMarkerSymbol = "triangle-down-open"
	ViolinMarkerSymbol206                          ViolinMarkerSymbol = "206"
	ViolinMarkerSymbol_triangle_down_dot           ViolinMarkerSymbol = "triangle-down-dot"
	ViolinMarkerSymbol306                          ViolinMarkerSymbol = "306"
	ViolinMarkerSymbol_triangle_down_open_dot      ViolinMarkerSymbol = "triangle-down-open-dot"
	ViolinMarkerSymbol7                            ViolinMarkerSymbol = "7"
	ViolinMarkerSymbol_triangle_left               ViolinMarkerSymbol = "triangle-left"
	ViolinMarkerSymbol107                          ViolinMarkerSymbol = "107"
	ViolinMarkerSymbol_triangle_left_open          ViolinMarkerSymbol = "triangle-left-open"
	ViolinMarkerSymbol207                          ViolinMarkerSymbol = "207"
	ViolinMarkerSymbol_triangle_left_dot           ViolinMarkerSymbol = "triangle-left-dot"
	ViolinMarkerSymbol307                          ViolinMarkerSymbol = "307"
	ViolinMarkerSymbol_triangle_left_open_dot      ViolinMarkerSymbol = "triangle-left-open-dot"
	ViolinMarkerSymbol8                            ViolinMarkerSymbol = "8"
	ViolinMarkerSymbol_triangle_right              ViolinMarkerSymbol = "triangle-right"
	ViolinMarkerSymbol108                          ViolinMarkerSymbol = "108"
	ViolinMarkerSymbol_triangle_right_open         ViolinMarkerSymbol = "triangle-right-open"
	ViolinMarkerSymbol208                          ViolinMarkerSymbol = "208"
	ViolinMarkerSymbol_triangle_right_dot          ViolinMarkerSymbol = "triangle-right-dot"
	ViolinMarkerSymbol308                          ViolinMarkerSymbol = "308"
	ViolinMarkerSymbol_triangle_right_open_dot     ViolinMarkerSymbol = "triangle-right-open-dot"
	ViolinMarkerSymbol9                            ViolinMarkerSymbol = "9"
	ViolinMarkerSymbol_triangle_ne                 ViolinMarkerSymbol = "triangle-ne"
	ViolinMarkerSymbol109                          ViolinMarkerSymbol = "109"
	ViolinMarkerSymbol_triangle_ne_open            ViolinMarkerSymbol = "triangle-ne-open"
	ViolinMarkerSymbol209                          ViolinMarkerSymbol = "209"
	ViolinMarkerSymbol_triangle_ne_dot             ViolinMarkerSymbol = "triangle-ne-dot"
	ViolinMarkerSymbol309                          ViolinMarkerSymbol = "309"
	ViolinMarkerSymbol_triangle_ne_open_dot        ViolinMarkerSymbol = "triangle-ne-open-dot"
	ViolinMarkerSymbol10                           ViolinMarkerSymbol = "10"
	ViolinMarkerSymbol_triangle_se                 ViolinMarkerSymbol = "triangle-se"
	ViolinMarkerSymbol110                          ViolinMarkerSymbol = "110"
	ViolinMarkerSymbol_triangle_se_open            ViolinMarkerSymbol = "triangle-se-open"
	ViolinMarkerSymbol210                          ViolinMarkerSymbol = "210"
	ViolinMarkerSymbol_triangle_se_dot             ViolinMarkerSymbol = "triangle-se-dot"
	ViolinMarkerSymbol310                          ViolinMarkerSymbol = "310"
	ViolinMarkerSymbol_triangle_se_open_dot        ViolinMarkerSymbol = "triangle-se-open-dot"
	ViolinMarkerSymbol11                           ViolinMarkerSymbol = "11"
	ViolinMarkerSymbol_triangle_sw                 ViolinMarkerSymbol = "triangle-sw"
	ViolinMarkerSymbol111                          ViolinMarkerSymbol = "111"
	ViolinMarkerSymbol_triangle_sw_open            ViolinMarkerSymbol = "triangle-sw-open"
	ViolinMarkerSymbol211                          ViolinMarkerSymbol = "211"
	ViolinMarkerSymbol_triangle_sw_dot             ViolinMarkerSymbol = "triangle-sw-dot"
	ViolinMarkerSymbol311                          ViolinMarkerSymbol = "311"
	ViolinMarkerSymbol_triangle_sw_open_dot        ViolinMarkerSymbol = "triangle-sw-open-dot"
	ViolinMarkerSymbol12                           ViolinMarkerSymbol = "12"
	ViolinMarkerSymbol_triangle_nw                 ViolinMarkerSymbol = "triangle-nw"
	ViolinMarkerSymbol112                          ViolinMarkerSymbol = "112"
	ViolinMarkerSymbol_triangle_nw_open            ViolinMarkerSymbol = "triangle-nw-open"
	ViolinMarkerSymbol212                          ViolinMarkerSymbol = "212"
	ViolinMarkerSymbol_triangle_nw_dot             ViolinMarkerSymbol = "triangle-nw-dot"
	ViolinMarkerSymbol312                          ViolinMarkerSymbol = "312"
	ViolinMarkerSymbol_triangle_nw_open_dot        ViolinMarkerSymbol = "triangle-nw-open-dot"
	ViolinMarkerSymbol13                           ViolinMarkerSymbol = "13"
	ViolinMarkerSymbol_pentagon                    ViolinMarkerSymbol = "pentagon"
	ViolinMarkerSymbol113                          ViolinMarkerSymbol = "113"
	ViolinMarkerSymbol_pentagon_open               ViolinMarkerSymbol = "pentagon-open"
	ViolinMarkerSymbol213                          ViolinMarkerSymbol = "213"
	ViolinMarkerSymbol_pentagon_dot                ViolinMarkerSymbol = "pentagon-dot"
	ViolinMarkerSymbol313                          ViolinMarkerSymbol = "313"
	ViolinMarkerSymbol_pentagon_open_dot           ViolinMarkerSymbol = "pentagon-open-dot"
	ViolinMarkerSymbol14                           ViolinMarkerSymbol = "14"
	ViolinMarkerSymbol_hexagon                     ViolinMarkerSymbol = "hexagon"
	ViolinMarkerSymbol114                          ViolinMarkerSymbol = "114"
	ViolinMarkerSymbol_hexagon_open                ViolinMarkerSymbol = "hexagon-open"
	ViolinMarkerSymbol214                          ViolinMarkerSymbol = "214"
	ViolinMarkerSymbol_hexagon_dot                 ViolinMarkerSymbol = "hexagon-dot"
	ViolinMarkerSymbol314                          ViolinMarkerSymbol = "314"
	ViolinMarkerSymbol_hexagon_open_dot            ViolinMarkerSymbol = "hexagon-open-dot"
	ViolinMarkerSymbol15                           ViolinMarkerSymbol = "15"
	ViolinMarkerSymbol_hexagon2                    ViolinMarkerSymbol = "hexagon2"
	ViolinMarkerSymbol115                          ViolinMarkerSymbol = "115"
	ViolinMarkerSymbol_hexagon2_open               ViolinMarkerSymbol = "hexagon2-open"
	ViolinMarkerSymbol215                          ViolinMarkerSymbol = "215"
	ViolinMarkerSymbol_hexagon2_dot                ViolinMarkerSymbol = "hexagon2-dot"
	ViolinMarkerSymbol315                          ViolinMarkerSymbol = "315"
	ViolinMarkerSymbol_hexagon2_open_dot           ViolinMarkerSymbol = "hexagon2-open-dot"
	ViolinMarkerSymbol16                           ViolinMarkerSymbol = "16"
	ViolinMarkerSymbol_octagon                     ViolinMarkerSymbol = "octagon"
	ViolinMarkerSymbol116                          ViolinMarkerSymbol = "116"
	ViolinMarkerSymbol_octagon_open                ViolinMarkerSymbol = "octagon-open"
	ViolinMarkerSymbol216                          ViolinMarkerSymbol = "216"
	ViolinMarkerSymbol_octagon_dot                 ViolinMarkerSymbol = "octagon-dot"
	ViolinMarkerSymbol316                          ViolinMarkerSymbol = "316"
	ViolinMarkerSymbol_octagon_open_dot            ViolinMarkerSymbol = "octagon-open-dot"
	ViolinMarkerSymbol17                           ViolinMarkerSymbol = "17"
	ViolinMarkerSymbol_star                        ViolinMarkerSymbol = "star"
	ViolinMarkerSymbol117                          ViolinMarkerSymbol = "117"
	ViolinMarkerSymbol_star_open                   ViolinMarkerSymbol = "star-open"
	ViolinMarkerSymbol217                          ViolinMarkerSymbol = "217"
	ViolinMarkerSymbol_star_dot                    ViolinMarkerSymbol = "star-dot"
	ViolinMarkerSymbol317                          ViolinMarkerSymbol = "317"
	ViolinMarkerSymbol_star_open_dot               ViolinMarkerSymbol = "star-open-dot"
	ViolinMarkerSymbol18                           ViolinMarkerSymbol = "18"
	ViolinMarkerSymbol_hexagram                    ViolinMarkerSymbol = "hexagram"
	ViolinMarkerSymbol118                          ViolinMarkerSymbol = "118"
	ViolinMarkerSymbol_hexagram_open               ViolinMarkerSymbol = "hexagram-open"
	ViolinMarkerSymbol218                          ViolinMarkerSymbol = "218"
	ViolinMarkerSymbol_hexagram_dot                ViolinMarkerSymbol = "hexagram-dot"
	ViolinMarkerSymbol318                          ViolinMarkerSymbol = "318"
	ViolinMarkerSymbol_hexagram_open_dot           ViolinMarkerSymbol = "hexagram-open-dot"
	ViolinMarkerSymbol19                           ViolinMarkerSymbol = "19"
	ViolinMarkerSymbol_star_triangle_up            ViolinMarkerSymbol = "star-triangle-up"
	ViolinMarkerSymbol119                          ViolinMarkerSymbol = "119"
	ViolinMarkerSymbol_star_triangle_up_open       ViolinMarkerSymbol = "star-triangle-up-open"
	ViolinMarkerSymbol219                          ViolinMarkerSymbol = "219"
	ViolinMarkerSymbol_star_triangle_up_dot        ViolinMarkerSymbol = "star-triangle-up-dot"
	ViolinMarkerSymbol319                          ViolinMarkerSymbol = "319"
	ViolinMarkerSymbol_star_triangle_up_open_dot   ViolinMarkerSymbol = "star-triangle-up-open-dot"
	ViolinMarkerSymbol20                           ViolinMarkerSymbol = "20"
	ViolinMarkerSymbol_star_triangle_down          ViolinMarkerSymbol = "star-triangle-down"
	ViolinMarkerSymbol120                          ViolinMarkerSymbol = "120"
	ViolinMarkerSymbol_star_triangle_down_open     ViolinMarkerSymbol = "star-triangle-down-open"
	ViolinMarkerSymbol220                          ViolinMarkerSymbol = "220"
	ViolinMarkerSymbol_star_triangle_down_dot      ViolinMarkerSymbol = "star-triangle-down-dot"
	ViolinMarkerSymbol320                          ViolinMarkerSymbol = "320"
	ViolinMarkerSymbol_star_triangle_down_open_dot ViolinMarkerSymbol = "star-triangle-down-open-dot"
	ViolinMarkerSymbol21                           ViolinMarkerSymbol = "21"
	ViolinMarkerSymbol_star_square                 ViolinMarkerSymbol = "star-square"
	ViolinMarkerSymbol121                          ViolinMarkerSymbol = "121"
	ViolinMarkerSymbol_star_square_open            ViolinMarkerSymbol = "star-square-open"
	ViolinMarkerSymbol221                          ViolinMarkerSymbol = "221"
	ViolinMarkerSymbol_star_square_dot             ViolinMarkerSymbol = "star-square-dot"
	ViolinMarkerSymbol321                          ViolinMarkerSymbol = "321"
	ViolinMarkerSymbol_star_square_open_dot        ViolinMarkerSymbol = "star-square-open-dot"
	ViolinMarkerSymbol22                           ViolinMarkerSymbol = "22"
	ViolinMarkerSymbol_star_diamond                ViolinMarkerSymbol = "star-diamond"
	ViolinMarkerSymbol122                          ViolinMarkerSymbol = "122"
	ViolinMarkerSymbol_star_diamond_open           ViolinMarkerSymbol = "star-diamond-open"
	ViolinMarkerSymbol222                          ViolinMarkerSymbol = "222"
	ViolinMarkerSymbol_star_diamond_dot            ViolinMarkerSymbol = "star-diamond-dot"
	ViolinMarkerSymbol322                          ViolinMarkerSymbol = "322"
	ViolinMarkerSymbol_star_diamond_open_dot       ViolinMarkerSymbol = "star-diamond-open-dot"
	ViolinMarkerSymbol23                           ViolinMarkerSymbol = "23"
	ViolinMarkerSymbol_diamond_tall                ViolinMarkerSymbol = "diamond-tall"
	ViolinMarkerSymbol123                          ViolinMarkerSymbol = "123"
	ViolinMarkerSymbol_diamond_tall_open           ViolinMarkerSymbol = "diamond-tall-open"
	ViolinMarkerSymbol223                          ViolinMarkerSymbol = "223"
	ViolinMarkerSymbol_diamond_tall_dot            ViolinMarkerSymbol = "diamond-tall-dot"
	ViolinMarkerSymbol323                          ViolinMarkerSymbol = "323"
	ViolinMarkerSymbol_diamond_tall_open_dot       ViolinMarkerSymbol = "diamond-tall-open-dot"
	ViolinMarkerSymbol24                           ViolinMarkerSymbol = "24"
	ViolinMarkerSymbol_diamond_wide                ViolinMarkerSymbol = "diamond-wide"
	ViolinMarkerSymbol124                          ViolinMarkerSymbol = "124"
	ViolinMarkerSymbol_diamond_wide_open           ViolinMarkerSymbol = "diamond-wide-open"
	ViolinMarkerSymbol224                          ViolinMarkerSymbol = "224"
	ViolinMarkerSymbol_diamond_wide_dot            ViolinMarkerSymbol = "diamond-wide-dot"
	ViolinMarkerSymbol324                          ViolinMarkerSymbol = "324"
	ViolinMarkerSymbol_diamond_wide_open_dot       ViolinMarkerSymbol = "diamond-wide-open-dot"
	ViolinMarkerSymbol25                           ViolinMarkerSymbol = "25"
	ViolinMarkerSymbol_hourglass                   ViolinMarkerSymbol = "hourglass"
	ViolinMarkerSymbol125                          ViolinMarkerSymbol = "125"
	ViolinMarkerSymbol_hourglass_open              ViolinMarkerSymbol = "hourglass-open"
	ViolinMarkerSymbol26                           ViolinMarkerSymbol = "26"
	ViolinMarkerSymbol_bowtie                      ViolinMarkerSymbol = "bowtie"
	ViolinMarkerSymbol126                          ViolinMarkerSymbol = "126"
	ViolinMarkerSymbol_bowtie_open                 ViolinMarkerSymbol = "bowtie-open"
	ViolinMarkerSymbol27                           ViolinMarkerSymbol = "27"
	ViolinMarkerSymbol_circle_cross                ViolinMarkerSymbol = "circle-cross"
	ViolinMarkerSymbol127                          ViolinMarkerSymbol = "127"
	ViolinMarkerSymbol_circle_cross_open           ViolinMarkerSymbol = "circle-cross-open"
	ViolinMarkerSymbol28                           ViolinMarkerSymbol = "28"
	ViolinMarkerSymbol_circle_x                    ViolinMarkerSymbol = "circle-x"
	ViolinMarkerSymbol128                          ViolinMarkerSymbol = "128"
	ViolinMarkerSymbol_circle_x_open               ViolinMarkerSymbol = "circle-x-open"
	ViolinMarkerSymbol29                           ViolinMarkerSymbol = "29"
	ViolinMarkerSymbol_square_cross                ViolinMarkerSymbol = "square-cross"
	ViolinMarkerSymbol129                          ViolinMarkerSymbol = "129"
	ViolinMarkerSymbol_square_cross_open           ViolinMarkerSymbol = "square-cross-open"
	ViolinMarkerSymbol30                           ViolinMarkerSymbol = "30"
	ViolinMarkerSymbol_square_x                    ViolinMarkerSymbol = "square-x"
	ViolinMarkerSymbol130                          ViolinMarkerSymbol = "130"
	ViolinMarkerSymbol_square_x_open               ViolinMarkerSymbol = "square-x-open"
	ViolinMarkerSymbol31                           ViolinMarkerSymbol = "31"
	ViolinMarkerSymbol_diamond_cross               ViolinMarkerSymbol = "diamond-cross"
	ViolinMarkerSymbol131                          ViolinMarkerSymbol = "131"
	ViolinMarkerSymbol_diamond_cross_open          ViolinMarkerSymbol = "diamond-cross-open"
	ViolinMarkerSymbol32                           ViolinMarkerSymbol = "32"
	ViolinMarkerSymbol_diamond_x                   ViolinMarkerSymbol = "diamond-x"
	ViolinMarkerSymbol132                          ViolinMarkerSymbol = "132"
	ViolinMarkerSymbol_diamond_x_open              ViolinMarkerSymbol = "diamond-x-open"
	ViolinMarkerSymbol33                           ViolinMarkerSymbol = "33"
	ViolinMarkerSymbol_cross_thin                  ViolinMarkerSymbol = "cross-thin"
	ViolinMarkerSymbol133                          ViolinMarkerSymbol = "133"
	ViolinMarkerSymbol_cross_thin_open             ViolinMarkerSymbol = "cross-thin-open"
	ViolinMarkerSymbol34                           ViolinMarkerSymbol = "34"
	ViolinMarkerSymbol_x_thin                      ViolinMarkerSymbol = "x-thin"
	ViolinMarkerSymbol134                          ViolinMarkerSymbol = "134"
	ViolinMarkerSymbol_x_thin_open                 ViolinMarkerSymbol = "x-thin-open"
	ViolinMarkerSymbol35                           ViolinMarkerSymbol = "35"
	ViolinMarkerSymbol_asterisk                    ViolinMarkerSymbol = "asterisk"
	ViolinMarkerSymbol135                          ViolinMarkerSymbol = "135"
	ViolinMarkerSymbol_asterisk_open               ViolinMarkerSymbol = "asterisk-open"
	ViolinMarkerSymbol36                           ViolinMarkerSymbol = "36"
	ViolinMarkerSymbol_hash                        ViolinMarkerSymbol = "hash"
	ViolinMarkerSymbol136                          ViolinMarkerSymbol = "136"
	ViolinMarkerSymbol_hash_open                   ViolinMarkerSymbol = "hash-open"
	ViolinMarkerSymbol236                          ViolinMarkerSymbol = "236"
	ViolinMarkerSymbol_hash_dot                    ViolinMarkerSymbol = "hash-dot"
	ViolinMarkerSymbol336                          ViolinMarkerSymbol = "336"
	ViolinMarkerSymbol_hash_open_dot               ViolinMarkerSymbol = "hash-open-dot"
	ViolinMarkerSymbol37                           ViolinMarkerSymbol = "37"
	ViolinMarkerSymbol_y_up                        ViolinMarkerSymbol = "y-up"
	ViolinMarkerSymbol137                          ViolinMarkerSymbol = "137"
	ViolinMarkerSymbol_y_up_open                   ViolinMarkerSymbol = "y-up-open"
	ViolinMarkerSymbol38                           ViolinMarkerSymbol = "38"
	ViolinMarkerSymbol_y_down                      ViolinMarkerSymbol = "y-down"
	ViolinMarkerSymbol138                          ViolinMarkerSymbol = "138"
	ViolinMarkerSymbol_y_down_open                 ViolinMarkerSymbol = "y-down-open"
	ViolinMarkerSymbol39                           ViolinMarkerSymbol = "39"
	ViolinMarkerSymbol_y_left                      ViolinMarkerSymbol = "y-left"
	ViolinMarkerSymbol139                          ViolinMarkerSymbol = "139"
	ViolinMarkerSymbol_y_left_open                 ViolinMarkerSymbol = "y-left-open"
	ViolinMarkerSymbol40                           ViolinMarkerSymbol = "40"
	ViolinMarkerSymbol_y_right                     ViolinMarkerSymbol = "y-right"
	ViolinMarkerSymbol140                          ViolinMarkerSymbol = "140"
	ViolinMarkerSymbol_y_right_open                ViolinMarkerSymbol = "y-right-open"
	ViolinMarkerSymbol41                           ViolinMarkerSymbol = "41"
	ViolinMarkerSymbol_line_ew                     ViolinMarkerSymbol = "line-ew"
	ViolinMarkerSymbol141                          ViolinMarkerSymbol = "141"
	ViolinMarkerSymbol_line_ew_open                ViolinMarkerSymbol = "line-ew-open"
	ViolinMarkerSymbol42                           ViolinMarkerSymbol = "42"
	ViolinMarkerSymbol_line_ns                     ViolinMarkerSymbol = "line-ns"
	ViolinMarkerSymbol142                          ViolinMarkerSymbol = "142"
	ViolinMarkerSymbol_line_ns_open                ViolinMarkerSymbol = "line-ns-open"
	ViolinMarkerSymbol43                           ViolinMarkerSymbol = "43"
	ViolinMarkerSymbol_line_ne                     ViolinMarkerSymbol = "line-ne"
	ViolinMarkerSymbol143                          ViolinMarkerSymbol = "143"
	ViolinMarkerSymbol_line_ne_open                ViolinMarkerSymbol = "line-ne-open"
	ViolinMarkerSymbol44                           ViolinMarkerSymbol = "44"
	ViolinMarkerSymbol_line_nw                     ViolinMarkerSymbol = "line-nw"
	ViolinMarkerSymbol144                          ViolinMarkerSymbol = "144"
	ViolinMarkerSymbol_line_nw_open                ViolinMarkerSymbol = "line-nw-open"
)

// ViolinOrientation Sets the orientation of the violin(s). If *v* (*h*), the distribution is visualized along the vertical (horizontal).
type ViolinOrientation string

const (
	ViolinOrientation_v ViolinOrientation = "v"
	ViolinOrientation_h ViolinOrientation = "h"
)

// ViolinPoints If *outliers*, only the sample points lying outside the whiskers are shown If *suspectedoutliers*, the outlier points are shown and points either less than 4*Q1-3*Q3 or greater than 4*Q3-3*Q1 are highlighted (see `outliercolor`) If *all*, all sample points are shown If *false*, only the violins are shown with no sample points. Defaults to *suspectedoutliers* when `marker.outliercolor` or `marker.line.outliercolor` is set, otherwise defaults to *outliers*.
type ViolinPoints interface{}

var (
	ViolinPoints_all               ViolinPoints = "all"
	ViolinPoints_outliers          ViolinPoints = "outliers"
	ViolinPoints_suspectedoutliers ViolinPoints = "suspectedoutliers"
	ViolinPoints_False             ViolinPoints = false
)

// ViolinScalemode Sets the metric by which the width of each violin is determined.*width* means each violin has the same (max) width*count* means the violins are scaled by the number of sample points makingup each violin.
type ViolinScalemode string

const (
	ViolinScalemode_width ViolinScalemode = "width"
	ViolinScalemode_count ViolinScalemode = "count"
)

// ViolinSide Determines on which side of the position value the density function making up one half of a violin is plotted. Useful when comparing two violin traces under *overlay* mode, where one trace has `side` set to *positive* and the other to *negative*.
type ViolinSide string

const (
	ViolinSide_both     ViolinSide = "both"
	ViolinSide_positive ViolinSide = "positive"
	ViolinSide_negative ViolinSide = "negative"
)

// ViolinSpanmode Sets the method by which the span in data space where the density function will be computed. *soft* means the span goes from the sample's minimum value minus two bandwidths to the sample's maximum value plus two bandwidths. *hard* means the span goes from the sample's minimum to its maximum value. For custom span settings, use mode *manual* and fill in the `span` attribute.
type ViolinSpanmode string

const (
	ViolinSpanmode_soft   ViolinSpanmode = "soft"
	ViolinSpanmode_hard   ViolinSpanmode = "hard"
	ViolinSpanmode_manual ViolinSpanmode = "manual"
)

// ViolinVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ViolinVisible interface{}

var (
	ViolinVisible_True       ViolinVisible = true
	ViolinVisible_False      ViolinVisible = false
	ViolinVisible_legendonly ViolinVisible = "legendonly"
)

// VolumeColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type VolumeColorbarExponentformat string

const (
	VolumeColorbarExponentformat_none  VolumeColorbarExponentformat = "none"
	VolumeColorbarExponentformat_e     VolumeColorbarExponentformat = "e"
	VolumeColorbarExponentformat_E     VolumeColorbarExponentformat = "E"
	VolumeColorbarExponentformat_power VolumeColorbarExponentformat = "power"
	VolumeColorbarExponentformat_SI    VolumeColorbarExponentformat = "SI"
	VolumeColorbarExponentformat_B     VolumeColorbarExponentformat = "B"
)

// VolumeColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type VolumeColorbarLenmode string

const (
	VolumeColorbarLenmode_fraction VolumeColorbarLenmode = "fraction"
	VolumeColorbarLenmode_pixels   VolumeColorbarLenmode = "pixels"
)

// VolumeColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type VolumeColorbarShowexponent string

const (
	VolumeColorbarShowexponent_all   VolumeColorbarShowexponent = "all"
	VolumeColorbarShowexponent_first VolumeColorbarShowexponent = "first"
	VolumeColorbarShowexponent_last  VolumeColorbarShowexponent = "last"
	VolumeColorbarShowexponent_none  VolumeColorbarShowexponent = "none"
)

// VolumeColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type VolumeColorbarShowtickprefix string

const (
	VolumeColorbarShowtickprefix_all   VolumeColorbarShowtickprefix = "all"
	VolumeColorbarShowtickprefix_first VolumeColorbarShowtickprefix = "first"
	VolumeColorbarShowtickprefix_last  VolumeColorbarShowtickprefix = "last"
	VolumeColorbarShowtickprefix_none  VolumeColorbarShowtickprefix = "none"
)

// VolumeColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type VolumeColorbarShowticksuffix string

const (
	VolumeColorbarShowticksuffix_all   VolumeColorbarShowticksuffix = "all"
	VolumeColorbarShowticksuffix_first VolumeColorbarShowticksuffix = "first"
	VolumeColorbarShowticksuffix_last  VolumeColorbarShowticksuffix = "last"
	VolumeColorbarShowticksuffix_none  VolumeColorbarShowticksuffix = "none"
)

// VolumeColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type VolumeColorbarThicknessmode string

const (
	VolumeColorbarThicknessmode_fraction VolumeColorbarThicknessmode = "fraction"
	VolumeColorbarThicknessmode_pixels   VolumeColorbarThicknessmode = "pixels"
)

// VolumeColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type VolumeColorbarTickmode string

const (
	VolumeColorbarTickmode_auto   VolumeColorbarTickmode = "auto"
	VolumeColorbarTickmode_linear VolumeColorbarTickmode = "linear"
	VolumeColorbarTickmode_array  VolumeColorbarTickmode = "array"
)

// VolumeColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type VolumeColorbarTicks string

const (
	VolumeColorbarTicks_outside VolumeColorbarTicks = "outside"
	VolumeColorbarTicks_inside  VolumeColorbarTicks = "inside"
)

// VolumeColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type VolumeColorbarTitleSide string

const (
	VolumeColorbarTitleSide_right  VolumeColorbarTitleSide = "right"
	VolumeColorbarTitleSide_top    VolumeColorbarTitleSide = "top"
	VolumeColorbarTitleSide_bottom VolumeColorbarTitleSide = "bottom"
)

// VolumeColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type VolumeColorbarXanchor string

const (
	VolumeColorbarXanchor_left   VolumeColorbarXanchor = "left"
	VolumeColorbarXanchor_center VolumeColorbarXanchor = "center"
	VolumeColorbarXanchor_right  VolumeColorbarXanchor = "right"
)

// VolumeColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type VolumeColorbarYanchor string

const (
	VolumeColorbarYanchor_top    VolumeColorbarYanchor = "top"
	VolumeColorbarYanchor_middle VolumeColorbarYanchor = "middle"
	VolumeColorbarYanchor_bottom VolumeColorbarYanchor = "bottom"
)

// VolumeHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type VolumeHoverlabelAlign string

const (
	VolumeHoverlabelAlign_left  VolumeHoverlabelAlign = "left"
	VolumeHoverlabelAlign_right VolumeHoverlabelAlign = "right"
	VolumeHoverlabelAlign_auto  VolumeHoverlabelAlign = "auto"
)

// VolumeVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type VolumeVisible interface{}

var (
	VolumeVisible_True       VolumeVisible = true
	VolumeVisible_False      VolumeVisible = false
	VolumeVisible_legendonly VolumeVisible = "legendonly"
)

// WaterfallConnectorMode Sets the shape of connector lines.
type WaterfallConnectorMode string

const (
	WaterfallConnectorMode_spanning WaterfallConnectorMode = "spanning"
	WaterfallConnectorMode_between  WaterfallConnectorMode = "between"
)

// WaterfallConstraintext Constrain the size of text inside or outside a bar to be no larger than the bar itself.
type WaterfallConstraintext string

const (
	WaterfallConstraintext_inside  WaterfallConstraintext = "inside"
	WaterfallConstraintext_outside WaterfallConstraintext = "outside"
	WaterfallConstraintext_both    WaterfallConstraintext = "both"
	WaterfallConstraintext_none    WaterfallConstraintext = "none"
)

// WaterfallHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type WaterfallHoverlabelAlign string

const (
	WaterfallHoverlabelAlign_left  WaterfallHoverlabelAlign = "left"
	WaterfallHoverlabelAlign_right WaterfallHoverlabelAlign = "right"
	WaterfallHoverlabelAlign_auto  WaterfallHoverlabelAlign = "auto"
)

// WaterfallInsidetextanchor Determines if texts are kept at center or start/end points in `textposition` *inside* mode.
type WaterfallInsidetextanchor string

const (
	WaterfallInsidetextanchor_end    WaterfallInsidetextanchor = "end"
	WaterfallInsidetextanchor_middle WaterfallInsidetextanchor = "middle"
	WaterfallInsidetextanchor_start  WaterfallInsidetextanchor = "start"
)

// WaterfallOrientation Sets the orientation of the bars. With *v* (*h*), the value of the each bar spans along the vertical (horizontal).
type WaterfallOrientation string

const (
	WaterfallOrientation_v WaterfallOrientation = "v"
	WaterfallOrientation_h WaterfallOrientation = "h"
)

// WaterfallTextposition Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside.
type WaterfallTextposition string

const (
	WaterfallTextposition_inside  WaterfallTextposition = "inside"
	WaterfallTextposition_outside WaterfallTextposition = "outside"
	WaterfallTextposition_auto    WaterfallTextposition = "auto"
	WaterfallTextposition_none    WaterfallTextposition = "none"
)

// WaterfallVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type WaterfallVisible interface{}

var (
	WaterfallVisible_True       WaterfallVisible = true
	WaterfallVisible_False      WaterfallVisible = false
	WaterfallVisible_legendonly WaterfallVisible = "legendonly"
)
