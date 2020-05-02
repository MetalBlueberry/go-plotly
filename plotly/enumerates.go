package plotly

// Enumerate section


// AreaHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type AreaHoverlabelAlign string

const (
    AreaHoverlabelAlign_left AreaHoverlabelAlign = "left" // left
    AreaHoverlabelAlign_right AreaHoverlabelAlign = "right" // right
    AreaHoverlabelAlign_auto AreaHoverlabelAlign = "auto" // auto
)

// AreaMarkerSymbol Area traces are deprecated! Please switch to the *barpolar* trace type. Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type AreaMarkerSymbol string

const (
    AreaMarkerSymbol0 AreaMarkerSymbol = "0" // 0
    AreaMarkerSymbol_circle AreaMarkerSymbol = "circle" // circle
    AreaMarkerSymbol100 AreaMarkerSymbol = "100" // 100
    AreaMarkerSymbol_circle_open AreaMarkerSymbol = "circle-open" // circle-open
    AreaMarkerSymbol200 AreaMarkerSymbol = "200" // 200
    AreaMarkerSymbol_circle_dot AreaMarkerSymbol = "circle-dot" // circle-dot
    AreaMarkerSymbol300 AreaMarkerSymbol = "300" // 300
    AreaMarkerSymbol_circle_open_dot AreaMarkerSymbol = "circle-open-dot" // circle-open-dot
    AreaMarkerSymbol1 AreaMarkerSymbol = "1" // 1
    AreaMarkerSymbol_square AreaMarkerSymbol = "square" // square
    AreaMarkerSymbol101 AreaMarkerSymbol = "101" // 101
    AreaMarkerSymbol_square_open AreaMarkerSymbol = "square-open" // square-open
    AreaMarkerSymbol201 AreaMarkerSymbol = "201" // 201
    AreaMarkerSymbol_square_dot AreaMarkerSymbol = "square-dot" // square-dot
    AreaMarkerSymbol301 AreaMarkerSymbol = "301" // 301
    AreaMarkerSymbol_square_open_dot AreaMarkerSymbol = "square-open-dot" // square-open-dot
    AreaMarkerSymbol2 AreaMarkerSymbol = "2" // 2
    AreaMarkerSymbol_diamond AreaMarkerSymbol = "diamond" // diamond
    AreaMarkerSymbol102 AreaMarkerSymbol = "102" // 102
    AreaMarkerSymbol_diamond_open AreaMarkerSymbol = "diamond-open" // diamond-open
    AreaMarkerSymbol202 AreaMarkerSymbol = "202" // 202
    AreaMarkerSymbol_diamond_dot AreaMarkerSymbol = "diamond-dot" // diamond-dot
    AreaMarkerSymbol302 AreaMarkerSymbol = "302" // 302
    AreaMarkerSymbol_diamond_open_dot AreaMarkerSymbol = "diamond-open-dot" // diamond-open-dot
    AreaMarkerSymbol3 AreaMarkerSymbol = "3" // 3
    AreaMarkerSymbol_cross AreaMarkerSymbol = "cross" // cross
    AreaMarkerSymbol103 AreaMarkerSymbol = "103" // 103
    AreaMarkerSymbol_cross_open AreaMarkerSymbol = "cross-open" // cross-open
    AreaMarkerSymbol203 AreaMarkerSymbol = "203" // 203
    AreaMarkerSymbol_cross_dot AreaMarkerSymbol = "cross-dot" // cross-dot
    AreaMarkerSymbol303 AreaMarkerSymbol = "303" // 303
    AreaMarkerSymbol_cross_open_dot AreaMarkerSymbol = "cross-open-dot" // cross-open-dot
    AreaMarkerSymbol4 AreaMarkerSymbol = "4" // 4
    AreaMarkerSymbol_x AreaMarkerSymbol = "x" // x
    AreaMarkerSymbol104 AreaMarkerSymbol = "104" // 104
    AreaMarkerSymbol_x_open AreaMarkerSymbol = "x-open" // x-open
    AreaMarkerSymbol204 AreaMarkerSymbol = "204" // 204
    AreaMarkerSymbol_x_dot AreaMarkerSymbol = "x-dot" // x-dot
    AreaMarkerSymbol304 AreaMarkerSymbol = "304" // 304
    AreaMarkerSymbol_x_open_dot AreaMarkerSymbol = "x-open-dot" // x-open-dot
    AreaMarkerSymbol5 AreaMarkerSymbol = "5" // 5
    AreaMarkerSymbol_triangle_up AreaMarkerSymbol = "triangle-up" // triangle-up
    AreaMarkerSymbol105 AreaMarkerSymbol = "105" // 105
    AreaMarkerSymbol_triangle_up_open AreaMarkerSymbol = "triangle-up-open" // triangle-up-open
    AreaMarkerSymbol205 AreaMarkerSymbol = "205" // 205
    AreaMarkerSymbol_triangle_up_dot AreaMarkerSymbol = "triangle-up-dot" // triangle-up-dot
    AreaMarkerSymbol305 AreaMarkerSymbol = "305" // 305
    AreaMarkerSymbol_triangle_up_open_dot AreaMarkerSymbol = "triangle-up-open-dot" // triangle-up-open-dot
    AreaMarkerSymbol6 AreaMarkerSymbol = "6" // 6
    AreaMarkerSymbol_triangle_down AreaMarkerSymbol = "triangle-down" // triangle-down
    AreaMarkerSymbol106 AreaMarkerSymbol = "106" // 106
    AreaMarkerSymbol_triangle_down_open AreaMarkerSymbol = "triangle-down-open" // triangle-down-open
    AreaMarkerSymbol206 AreaMarkerSymbol = "206" // 206
    AreaMarkerSymbol_triangle_down_dot AreaMarkerSymbol = "triangle-down-dot" // triangle-down-dot
    AreaMarkerSymbol306 AreaMarkerSymbol = "306" // 306
    AreaMarkerSymbol_triangle_down_open_dot AreaMarkerSymbol = "triangle-down-open-dot" // triangle-down-open-dot
    AreaMarkerSymbol7 AreaMarkerSymbol = "7" // 7
    AreaMarkerSymbol_triangle_left AreaMarkerSymbol = "triangle-left" // triangle-left
    AreaMarkerSymbol107 AreaMarkerSymbol = "107" // 107
    AreaMarkerSymbol_triangle_left_open AreaMarkerSymbol = "triangle-left-open" // triangle-left-open
    AreaMarkerSymbol207 AreaMarkerSymbol = "207" // 207
    AreaMarkerSymbol_triangle_left_dot AreaMarkerSymbol = "triangle-left-dot" // triangle-left-dot
    AreaMarkerSymbol307 AreaMarkerSymbol = "307" // 307
    AreaMarkerSymbol_triangle_left_open_dot AreaMarkerSymbol = "triangle-left-open-dot" // triangle-left-open-dot
    AreaMarkerSymbol8 AreaMarkerSymbol = "8" // 8
    AreaMarkerSymbol_triangle_right AreaMarkerSymbol = "triangle-right" // triangle-right
    AreaMarkerSymbol108 AreaMarkerSymbol = "108" // 108
    AreaMarkerSymbol_triangle_right_open AreaMarkerSymbol = "triangle-right-open" // triangle-right-open
    AreaMarkerSymbol208 AreaMarkerSymbol = "208" // 208
    AreaMarkerSymbol_triangle_right_dot AreaMarkerSymbol = "triangle-right-dot" // triangle-right-dot
    AreaMarkerSymbol308 AreaMarkerSymbol = "308" // 308
    AreaMarkerSymbol_triangle_right_open_dot AreaMarkerSymbol = "triangle-right-open-dot" // triangle-right-open-dot
    AreaMarkerSymbol9 AreaMarkerSymbol = "9" // 9
    AreaMarkerSymbol_triangle_ne AreaMarkerSymbol = "triangle-ne" // triangle-ne
    AreaMarkerSymbol109 AreaMarkerSymbol = "109" // 109
    AreaMarkerSymbol_triangle_ne_open AreaMarkerSymbol = "triangle-ne-open" // triangle-ne-open
    AreaMarkerSymbol209 AreaMarkerSymbol = "209" // 209
    AreaMarkerSymbol_triangle_ne_dot AreaMarkerSymbol = "triangle-ne-dot" // triangle-ne-dot
    AreaMarkerSymbol309 AreaMarkerSymbol = "309" // 309
    AreaMarkerSymbol_triangle_ne_open_dot AreaMarkerSymbol = "triangle-ne-open-dot" // triangle-ne-open-dot
    AreaMarkerSymbol10 AreaMarkerSymbol = "10" // 10
    AreaMarkerSymbol_triangle_se AreaMarkerSymbol = "triangle-se" // triangle-se
    AreaMarkerSymbol110 AreaMarkerSymbol = "110" // 110
    AreaMarkerSymbol_triangle_se_open AreaMarkerSymbol = "triangle-se-open" // triangle-se-open
    AreaMarkerSymbol210 AreaMarkerSymbol = "210" // 210
    AreaMarkerSymbol_triangle_se_dot AreaMarkerSymbol = "triangle-se-dot" // triangle-se-dot
    AreaMarkerSymbol310 AreaMarkerSymbol = "310" // 310
    AreaMarkerSymbol_triangle_se_open_dot AreaMarkerSymbol = "triangle-se-open-dot" // triangle-se-open-dot
    AreaMarkerSymbol11 AreaMarkerSymbol = "11" // 11
    AreaMarkerSymbol_triangle_sw AreaMarkerSymbol = "triangle-sw" // triangle-sw
    AreaMarkerSymbol111 AreaMarkerSymbol = "111" // 111
    AreaMarkerSymbol_triangle_sw_open AreaMarkerSymbol = "triangle-sw-open" // triangle-sw-open
    AreaMarkerSymbol211 AreaMarkerSymbol = "211" // 211
    AreaMarkerSymbol_triangle_sw_dot AreaMarkerSymbol = "triangle-sw-dot" // triangle-sw-dot
    AreaMarkerSymbol311 AreaMarkerSymbol = "311" // 311
    AreaMarkerSymbol_triangle_sw_open_dot AreaMarkerSymbol = "triangle-sw-open-dot" // triangle-sw-open-dot
    AreaMarkerSymbol12 AreaMarkerSymbol = "12" // 12
    AreaMarkerSymbol_triangle_nw AreaMarkerSymbol = "triangle-nw" // triangle-nw
    AreaMarkerSymbol112 AreaMarkerSymbol = "112" // 112
    AreaMarkerSymbol_triangle_nw_open AreaMarkerSymbol = "triangle-nw-open" // triangle-nw-open
    AreaMarkerSymbol212 AreaMarkerSymbol = "212" // 212
    AreaMarkerSymbol_triangle_nw_dot AreaMarkerSymbol = "triangle-nw-dot" // triangle-nw-dot
    AreaMarkerSymbol312 AreaMarkerSymbol = "312" // 312
    AreaMarkerSymbol_triangle_nw_open_dot AreaMarkerSymbol = "triangle-nw-open-dot" // triangle-nw-open-dot
    AreaMarkerSymbol13 AreaMarkerSymbol = "13" // 13
    AreaMarkerSymbol_pentagon AreaMarkerSymbol = "pentagon" // pentagon
    AreaMarkerSymbol113 AreaMarkerSymbol = "113" // 113
    AreaMarkerSymbol_pentagon_open AreaMarkerSymbol = "pentagon-open" // pentagon-open
    AreaMarkerSymbol213 AreaMarkerSymbol = "213" // 213
    AreaMarkerSymbol_pentagon_dot AreaMarkerSymbol = "pentagon-dot" // pentagon-dot
    AreaMarkerSymbol313 AreaMarkerSymbol = "313" // 313
    AreaMarkerSymbol_pentagon_open_dot AreaMarkerSymbol = "pentagon-open-dot" // pentagon-open-dot
    AreaMarkerSymbol14 AreaMarkerSymbol = "14" // 14
    AreaMarkerSymbol_hexagon AreaMarkerSymbol = "hexagon" // hexagon
    AreaMarkerSymbol114 AreaMarkerSymbol = "114" // 114
    AreaMarkerSymbol_hexagon_open AreaMarkerSymbol = "hexagon-open" // hexagon-open
    AreaMarkerSymbol214 AreaMarkerSymbol = "214" // 214
    AreaMarkerSymbol_hexagon_dot AreaMarkerSymbol = "hexagon-dot" // hexagon-dot
    AreaMarkerSymbol314 AreaMarkerSymbol = "314" // 314
    AreaMarkerSymbol_hexagon_open_dot AreaMarkerSymbol = "hexagon-open-dot" // hexagon-open-dot
    AreaMarkerSymbol15 AreaMarkerSymbol = "15" // 15
    AreaMarkerSymbol_hexagon2 AreaMarkerSymbol = "hexagon2" // hexagon2
    AreaMarkerSymbol115 AreaMarkerSymbol = "115" // 115
    AreaMarkerSymbol_hexagon2_open AreaMarkerSymbol = "hexagon2-open" // hexagon2-open
    AreaMarkerSymbol215 AreaMarkerSymbol = "215" // 215
    AreaMarkerSymbol_hexagon2_dot AreaMarkerSymbol = "hexagon2-dot" // hexagon2-dot
    AreaMarkerSymbol315 AreaMarkerSymbol = "315" // 315
    AreaMarkerSymbol_hexagon2_open_dot AreaMarkerSymbol = "hexagon2-open-dot" // hexagon2-open-dot
    AreaMarkerSymbol16 AreaMarkerSymbol = "16" // 16
    AreaMarkerSymbol_octagon AreaMarkerSymbol = "octagon" // octagon
    AreaMarkerSymbol116 AreaMarkerSymbol = "116" // 116
    AreaMarkerSymbol_octagon_open AreaMarkerSymbol = "octagon-open" // octagon-open
    AreaMarkerSymbol216 AreaMarkerSymbol = "216" // 216
    AreaMarkerSymbol_octagon_dot AreaMarkerSymbol = "octagon-dot" // octagon-dot
    AreaMarkerSymbol316 AreaMarkerSymbol = "316" // 316
    AreaMarkerSymbol_octagon_open_dot AreaMarkerSymbol = "octagon-open-dot" // octagon-open-dot
    AreaMarkerSymbol17 AreaMarkerSymbol = "17" // 17
    AreaMarkerSymbol_star AreaMarkerSymbol = "star" // star
    AreaMarkerSymbol117 AreaMarkerSymbol = "117" // 117
    AreaMarkerSymbol_star_open AreaMarkerSymbol = "star-open" // star-open
    AreaMarkerSymbol217 AreaMarkerSymbol = "217" // 217
    AreaMarkerSymbol_star_dot AreaMarkerSymbol = "star-dot" // star-dot
    AreaMarkerSymbol317 AreaMarkerSymbol = "317" // 317
    AreaMarkerSymbol_star_open_dot AreaMarkerSymbol = "star-open-dot" // star-open-dot
    AreaMarkerSymbol18 AreaMarkerSymbol = "18" // 18
    AreaMarkerSymbol_hexagram AreaMarkerSymbol = "hexagram" // hexagram
    AreaMarkerSymbol118 AreaMarkerSymbol = "118" // 118
    AreaMarkerSymbol_hexagram_open AreaMarkerSymbol = "hexagram-open" // hexagram-open
    AreaMarkerSymbol218 AreaMarkerSymbol = "218" // 218
    AreaMarkerSymbol_hexagram_dot AreaMarkerSymbol = "hexagram-dot" // hexagram-dot
    AreaMarkerSymbol318 AreaMarkerSymbol = "318" // 318
    AreaMarkerSymbol_hexagram_open_dot AreaMarkerSymbol = "hexagram-open-dot" // hexagram-open-dot
    AreaMarkerSymbol19 AreaMarkerSymbol = "19" // 19
    AreaMarkerSymbol_star_triangle_up AreaMarkerSymbol = "star-triangle-up" // star-triangle-up
    AreaMarkerSymbol119 AreaMarkerSymbol = "119" // 119
    AreaMarkerSymbol_star_triangle_up_open AreaMarkerSymbol = "star-triangle-up-open" // star-triangle-up-open
    AreaMarkerSymbol219 AreaMarkerSymbol = "219" // 219
    AreaMarkerSymbol_star_triangle_up_dot AreaMarkerSymbol = "star-triangle-up-dot" // star-triangle-up-dot
    AreaMarkerSymbol319 AreaMarkerSymbol = "319" // 319
    AreaMarkerSymbol_star_triangle_up_open_dot AreaMarkerSymbol = "star-triangle-up-open-dot" // star-triangle-up-open-dot
    AreaMarkerSymbol20 AreaMarkerSymbol = "20" // 20
    AreaMarkerSymbol_star_triangle_down AreaMarkerSymbol = "star-triangle-down" // star-triangle-down
    AreaMarkerSymbol120 AreaMarkerSymbol = "120" // 120
    AreaMarkerSymbol_star_triangle_down_open AreaMarkerSymbol = "star-triangle-down-open" // star-triangle-down-open
    AreaMarkerSymbol220 AreaMarkerSymbol = "220" // 220
    AreaMarkerSymbol_star_triangle_down_dot AreaMarkerSymbol = "star-triangle-down-dot" // star-triangle-down-dot
    AreaMarkerSymbol320 AreaMarkerSymbol = "320" // 320
    AreaMarkerSymbol_star_triangle_down_open_dot AreaMarkerSymbol = "star-triangle-down-open-dot" // star-triangle-down-open-dot
    AreaMarkerSymbol21 AreaMarkerSymbol = "21" // 21
    AreaMarkerSymbol_star_square AreaMarkerSymbol = "star-square" // star-square
    AreaMarkerSymbol121 AreaMarkerSymbol = "121" // 121
    AreaMarkerSymbol_star_square_open AreaMarkerSymbol = "star-square-open" // star-square-open
    AreaMarkerSymbol221 AreaMarkerSymbol = "221" // 221
    AreaMarkerSymbol_star_square_dot AreaMarkerSymbol = "star-square-dot" // star-square-dot
    AreaMarkerSymbol321 AreaMarkerSymbol = "321" // 321
    AreaMarkerSymbol_star_square_open_dot AreaMarkerSymbol = "star-square-open-dot" // star-square-open-dot
    AreaMarkerSymbol22 AreaMarkerSymbol = "22" // 22
    AreaMarkerSymbol_star_diamond AreaMarkerSymbol = "star-diamond" // star-diamond
    AreaMarkerSymbol122 AreaMarkerSymbol = "122" // 122
    AreaMarkerSymbol_star_diamond_open AreaMarkerSymbol = "star-diamond-open" // star-diamond-open
    AreaMarkerSymbol222 AreaMarkerSymbol = "222" // 222
    AreaMarkerSymbol_star_diamond_dot AreaMarkerSymbol = "star-diamond-dot" // star-diamond-dot
    AreaMarkerSymbol322 AreaMarkerSymbol = "322" // 322
    AreaMarkerSymbol_star_diamond_open_dot AreaMarkerSymbol = "star-diamond-open-dot" // star-diamond-open-dot
    AreaMarkerSymbol23 AreaMarkerSymbol = "23" // 23
    AreaMarkerSymbol_diamond_tall AreaMarkerSymbol = "diamond-tall" // diamond-tall
    AreaMarkerSymbol123 AreaMarkerSymbol = "123" // 123
    AreaMarkerSymbol_diamond_tall_open AreaMarkerSymbol = "diamond-tall-open" // diamond-tall-open
    AreaMarkerSymbol223 AreaMarkerSymbol = "223" // 223
    AreaMarkerSymbol_diamond_tall_dot AreaMarkerSymbol = "diamond-tall-dot" // diamond-tall-dot
    AreaMarkerSymbol323 AreaMarkerSymbol = "323" // 323
    AreaMarkerSymbol_diamond_tall_open_dot AreaMarkerSymbol = "diamond-tall-open-dot" // diamond-tall-open-dot
    AreaMarkerSymbol24 AreaMarkerSymbol = "24" // 24
    AreaMarkerSymbol_diamond_wide AreaMarkerSymbol = "diamond-wide" // diamond-wide
    AreaMarkerSymbol124 AreaMarkerSymbol = "124" // 124
    AreaMarkerSymbol_diamond_wide_open AreaMarkerSymbol = "diamond-wide-open" // diamond-wide-open
    AreaMarkerSymbol224 AreaMarkerSymbol = "224" // 224
    AreaMarkerSymbol_diamond_wide_dot AreaMarkerSymbol = "diamond-wide-dot" // diamond-wide-dot
    AreaMarkerSymbol324 AreaMarkerSymbol = "324" // 324
    AreaMarkerSymbol_diamond_wide_open_dot AreaMarkerSymbol = "diamond-wide-open-dot" // diamond-wide-open-dot
    AreaMarkerSymbol25 AreaMarkerSymbol = "25" // 25
    AreaMarkerSymbol_hourglass AreaMarkerSymbol = "hourglass" // hourglass
    AreaMarkerSymbol125 AreaMarkerSymbol = "125" // 125
    AreaMarkerSymbol_hourglass_open AreaMarkerSymbol = "hourglass-open" // hourglass-open
    AreaMarkerSymbol26 AreaMarkerSymbol = "26" // 26
    AreaMarkerSymbol_bowtie AreaMarkerSymbol = "bowtie" // bowtie
    AreaMarkerSymbol126 AreaMarkerSymbol = "126" // 126
    AreaMarkerSymbol_bowtie_open AreaMarkerSymbol = "bowtie-open" // bowtie-open
    AreaMarkerSymbol27 AreaMarkerSymbol = "27" // 27
    AreaMarkerSymbol_circle_cross AreaMarkerSymbol = "circle-cross" // circle-cross
    AreaMarkerSymbol127 AreaMarkerSymbol = "127" // 127
    AreaMarkerSymbol_circle_cross_open AreaMarkerSymbol = "circle-cross-open" // circle-cross-open
    AreaMarkerSymbol28 AreaMarkerSymbol = "28" // 28
    AreaMarkerSymbol_circle_x AreaMarkerSymbol = "circle-x" // circle-x
    AreaMarkerSymbol128 AreaMarkerSymbol = "128" // 128
    AreaMarkerSymbol_circle_x_open AreaMarkerSymbol = "circle-x-open" // circle-x-open
    AreaMarkerSymbol29 AreaMarkerSymbol = "29" // 29
    AreaMarkerSymbol_square_cross AreaMarkerSymbol = "square-cross" // square-cross
    AreaMarkerSymbol129 AreaMarkerSymbol = "129" // 129
    AreaMarkerSymbol_square_cross_open AreaMarkerSymbol = "square-cross-open" // square-cross-open
    AreaMarkerSymbol30 AreaMarkerSymbol = "30" // 30
    AreaMarkerSymbol_square_x AreaMarkerSymbol = "square-x" // square-x
    AreaMarkerSymbol130 AreaMarkerSymbol = "130" // 130
    AreaMarkerSymbol_square_x_open AreaMarkerSymbol = "square-x-open" // square-x-open
    AreaMarkerSymbol31 AreaMarkerSymbol = "31" // 31
    AreaMarkerSymbol_diamond_cross AreaMarkerSymbol = "diamond-cross" // diamond-cross
    AreaMarkerSymbol131 AreaMarkerSymbol = "131" // 131
    AreaMarkerSymbol_diamond_cross_open AreaMarkerSymbol = "diamond-cross-open" // diamond-cross-open
    AreaMarkerSymbol32 AreaMarkerSymbol = "32" // 32
    AreaMarkerSymbol_diamond_x AreaMarkerSymbol = "diamond-x" // diamond-x
    AreaMarkerSymbol132 AreaMarkerSymbol = "132" // 132
    AreaMarkerSymbol_diamond_x_open AreaMarkerSymbol = "diamond-x-open" // diamond-x-open
    AreaMarkerSymbol33 AreaMarkerSymbol = "33" // 33
    AreaMarkerSymbol_cross_thin AreaMarkerSymbol = "cross-thin" // cross-thin
    AreaMarkerSymbol133 AreaMarkerSymbol = "133" // 133
    AreaMarkerSymbol_cross_thin_open AreaMarkerSymbol = "cross-thin-open" // cross-thin-open
    AreaMarkerSymbol34 AreaMarkerSymbol = "34" // 34
    AreaMarkerSymbol_x_thin AreaMarkerSymbol = "x-thin" // x-thin
    AreaMarkerSymbol134 AreaMarkerSymbol = "134" // 134
    AreaMarkerSymbol_x_thin_open AreaMarkerSymbol = "x-thin-open" // x-thin-open
    AreaMarkerSymbol35 AreaMarkerSymbol = "35" // 35
    AreaMarkerSymbol_asterisk AreaMarkerSymbol = "asterisk" // asterisk
    AreaMarkerSymbol135 AreaMarkerSymbol = "135" // 135
    AreaMarkerSymbol_asterisk_open AreaMarkerSymbol = "asterisk-open" // asterisk-open
    AreaMarkerSymbol36 AreaMarkerSymbol = "36" // 36
    AreaMarkerSymbol_hash AreaMarkerSymbol = "hash" // hash
    AreaMarkerSymbol136 AreaMarkerSymbol = "136" // 136
    AreaMarkerSymbol_hash_open AreaMarkerSymbol = "hash-open" // hash-open
    AreaMarkerSymbol236 AreaMarkerSymbol = "236" // 236
    AreaMarkerSymbol_hash_dot AreaMarkerSymbol = "hash-dot" // hash-dot
    AreaMarkerSymbol336 AreaMarkerSymbol = "336" // 336
    AreaMarkerSymbol_hash_open_dot AreaMarkerSymbol = "hash-open-dot" // hash-open-dot
    AreaMarkerSymbol37 AreaMarkerSymbol = "37" // 37
    AreaMarkerSymbol_y_up AreaMarkerSymbol = "y-up" // y-up
    AreaMarkerSymbol137 AreaMarkerSymbol = "137" // 137
    AreaMarkerSymbol_y_up_open AreaMarkerSymbol = "y-up-open" // y-up-open
    AreaMarkerSymbol38 AreaMarkerSymbol = "38" // 38
    AreaMarkerSymbol_y_down AreaMarkerSymbol = "y-down" // y-down
    AreaMarkerSymbol138 AreaMarkerSymbol = "138" // 138
    AreaMarkerSymbol_y_down_open AreaMarkerSymbol = "y-down-open" // y-down-open
    AreaMarkerSymbol39 AreaMarkerSymbol = "39" // 39
    AreaMarkerSymbol_y_left AreaMarkerSymbol = "y-left" // y-left
    AreaMarkerSymbol139 AreaMarkerSymbol = "139" // 139
    AreaMarkerSymbol_y_left_open AreaMarkerSymbol = "y-left-open" // y-left-open
    AreaMarkerSymbol40 AreaMarkerSymbol = "40" // 40
    AreaMarkerSymbol_y_right AreaMarkerSymbol = "y-right" // y-right
    AreaMarkerSymbol140 AreaMarkerSymbol = "140" // 140
    AreaMarkerSymbol_y_right_open AreaMarkerSymbol = "y-right-open" // y-right-open
    AreaMarkerSymbol41 AreaMarkerSymbol = "41" // 41
    AreaMarkerSymbol_line_ew AreaMarkerSymbol = "line-ew" // line-ew
    AreaMarkerSymbol141 AreaMarkerSymbol = "141" // 141
    AreaMarkerSymbol_line_ew_open AreaMarkerSymbol = "line-ew-open" // line-ew-open
    AreaMarkerSymbol42 AreaMarkerSymbol = "42" // 42
    AreaMarkerSymbol_line_ns AreaMarkerSymbol = "line-ns" // line-ns
    AreaMarkerSymbol142 AreaMarkerSymbol = "142" // 142
    AreaMarkerSymbol_line_ns_open AreaMarkerSymbol = "line-ns-open" // line-ns-open
    AreaMarkerSymbol43 AreaMarkerSymbol = "43" // 43
    AreaMarkerSymbol_line_ne AreaMarkerSymbol = "line-ne" // line-ne
    AreaMarkerSymbol143 AreaMarkerSymbol = "143" // 143
    AreaMarkerSymbol_line_ne_open AreaMarkerSymbol = "line-ne-open" // line-ne-open
    AreaMarkerSymbol44 AreaMarkerSymbol = "44" // 44
    AreaMarkerSymbol_line_nw AreaMarkerSymbol = "line-nw" // line-nw
    AreaMarkerSymbol144 AreaMarkerSymbol = "144" // 144
    AreaMarkerSymbol_line_nw_open AreaMarkerSymbol = "line-nw-open" // line-nw-open
)

// AreaVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type AreaVisible interface{}

const (
    AreaVisibleTrue bool = true
    AreaVisibleFalse bool = false
    AreaVisibleLegendonly string = "legendonly"
)

// BarConstraintext Constrain the size of text inside or outside a bar to be no larger than the bar itself.
type BarConstraintext string

const (
    BarConstraintext_inside BarConstraintext = "inside" // inside
    BarConstraintext_outside BarConstraintext = "outside" // outside
    BarConstraintext_both BarConstraintext = "both" // both
    BarConstraintext_none BarConstraintext = "none" // none
)

// BarErrorXType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type BarErrorXType string

const (
    BarErrorXType_percent BarErrorXType = "percent" // percent
    BarErrorXType_constant BarErrorXType = "constant" // constant
    BarErrorXType_sqrt BarErrorXType = "sqrt" // sqrt
    BarErrorXType_data BarErrorXType = "data" // data
)

// BarErrorYType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type BarErrorYType string

const (
    BarErrorYType_percent BarErrorYType = "percent" // percent
    BarErrorYType_constant BarErrorYType = "constant" // constant
    BarErrorYType_sqrt BarErrorYType = "sqrt" // sqrt
    BarErrorYType_data BarErrorYType = "data" // data
)

// BarHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type BarHoverlabelAlign string

const (
    BarHoverlabelAlign_left BarHoverlabelAlign = "left" // left
    BarHoverlabelAlign_right BarHoverlabelAlign = "right" // right
    BarHoverlabelAlign_auto BarHoverlabelAlign = "auto" // auto
)

// BarInsidetextanchor Determines if texts are kept at center or start/end points in `textposition` *inside* mode.
type BarInsidetextanchor string

const (
    BarInsidetextanchor_end BarInsidetextanchor = "end" // end
    BarInsidetextanchor_middle BarInsidetextanchor = "middle" // middle
    BarInsidetextanchor_start BarInsidetextanchor = "start" // start
)

// BarMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type BarMarkerColorbarExponentformat string

const (
    BarMarkerColorbarExponentformat_none BarMarkerColorbarExponentformat = "none" // none
    BarMarkerColorbarExponentformat_e BarMarkerColorbarExponentformat = "e" // e
    BarMarkerColorbarExponentformat_E BarMarkerColorbarExponentformat = "E" // E
    BarMarkerColorbarExponentformat_power BarMarkerColorbarExponentformat = "power" // power
    BarMarkerColorbarExponentformat_SI BarMarkerColorbarExponentformat = "SI" // SI
    BarMarkerColorbarExponentformat_B BarMarkerColorbarExponentformat = "B" // B
)

// BarMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type BarMarkerColorbarLenmode string

const (
    BarMarkerColorbarLenmode_fraction BarMarkerColorbarLenmode = "fraction" // fraction
    BarMarkerColorbarLenmode_pixels BarMarkerColorbarLenmode = "pixels" // pixels
)

// BarMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type BarMarkerColorbarShowexponent string

const (
    BarMarkerColorbarShowexponent_all BarMarkerColorbarShowexponent = "all" // all
    BarMarkerColorbarShowexponent_first BarMarkerColorbarShowexponent = "first" // first
    BarMarkerColorbarShowexponent_last BarMarkerColorbarShowexponent = "last" // last
    BarMarkerColorbarShowexponent_none BarMarkerColorbarShowexponent = "none" // none
)

// BarMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type BarMarkerColorbarShowtickprefix string

const (
    BarMarkerColorbarShowtickprefix_all BarMarkerColorbarShowtickprefix = "all" // all
    BarMarkerColorbarShowtickprefix_first BarMarkerColorbarShowtickprefix = "first" // first
    BarMarkerColorbarShowtickprefix_last BarMarkerColorbarShowtickprefix = "last" // last
    BarMarkerColorbarShowtickprefix_none BarMarkerColorbarShowtickprefix = "none" // none
)

// BarMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type BarMarkerColorbarShowticksuffix string

const (
    BarMarkerColorbarShowticksuffix_all BarMarkerColorbarShowticksuffix = "all" // all
    BarMarkerColorbarShowticksuffix_first BarMarkerColorbarShowticksuffix = "first" // first
    BarMarkerColorbarShowticksuffix_last BarMarkerColorbarShowticksuffix = "last" // last
    BarMarkerColorbarShowticksuffix_none BarMarkerColorbarShowticksuffix = "none" // none
)

// BarMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type BarMarkerColorbarThicknessmode string

const (
    BarMarkerColorbarThicknessmode_fraction BarMarkerColorbarThicknessmode = "fraction" // fraction
    BarMarkerColorbarThicknessmode_pixels BarMarkerColorbarThicknessmode = "pixels" // pixels
)

// BarMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type BarMarkerColorbarTickmode string

const (
    BarMarkerColorbarTickmode_auto BarMarkerColorbarTickmode = "auto" // auto
    BarMarkerColorbarTickmode_linear BarMarkerColorbarTickmode = "linear" // linear
    BarMarkerColorbarTickmode_array BarMarkerColorbarTickmode = "array" // array
)

// BarMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type BarMarkerColorbarTicks string

const (
    BarMarkerColorbarTicks_outside BarMarkerColorbarTicks = "outside" // outside
    BarMarkerColorbarTicks_inside BarMarkerColorbarTicks = "inside" // inside
)

// BarMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type BarMarkerColorbarTitleSide string

const (
    BarMarkerColorbarTitleSide_right BarMarkerColorbarTitleSide = "right" // right
    BarMarkerColorbarTitleSide_top BarMarkerColorbarTitleSide = "top" // top
    BarMarkerColorbarTitleSide_bottom BarMarkerColorbarTitleSide = "bottom" // bottom
)

// BarMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type BarMarkerColorbarXanchor string

const (
    BarMarkerColorbarXanchor_left BarMarkerColorbarXanchor = "left" // left
    BarMarkerColorbarXanchor_center BarMarkerColorbarXanchor = "center" // center
    BarMarkerColorbarXanchor_right BarMarkerColorbarXanchor = "right" // right
)

// BarMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type BarMarkerColorbarYanchor string

const (
    BarMarkerColorbarYanchor_top BarMarkerColorbarYanchor = "top" // top
    BarMarkerColorbarYanchor_middle BarMarkerColorbarYanchor = "middle" // middle
    BarMarkerColorbarYanchor_bottom BarMarkerColorbarYanchor = "bottom" // bottom
)

// BarOrientation Sets the orientation of the bars. With *v* (*h*), the value of the each bar spans along the vertical (horizontal).
type BarOrientation string

const (
    BarOrientation_v BarOrientation = "v" // v
    BarOrientation_h BarOrientation = "h" // h
)

// BarTextposition Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside.
type BarTextposition string

const (
    BarTextposition_inside BarTextposition = "inside" // inside
    BarTextposition_outside BarTextposition = "outside" // outside
    BarTextposition_auto BarTextposition = "auto" // auto
    BarTextposition_none BarTextposition = "none" // none
)

// BarVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type BarVisible interface{}

const (
    BarVisibleTrue bool = true
    BarVisibleFalse bool = false
    BarVisibleLegendonly string = "legendonly"
)

// BarXcalendar Sets the calendar system to use with `x` date data.
type BarXcalendar string

const (
    BarXcalendar_gregorian BarXcalendar = "gregorian" // gregorian
    BarXcalendar_chinese BarXcalendar = "chinese" // chinese
    BarXcalendar_coptic BarXcalendar = "coptic" // coptic
    BarXcalendar_discworld BarXcalendar = "discworld" // discworld
    BarXcalendar_ethiopian BarXcalendar = "ethiopian" // ethiopian
    BarXcalendar_hebrew BarXcalendar = "hebrew" // hebrew
    BarXcalendar_islamic BarXcalendar = "islamic" // islamic
    BarXcalendar_julian BarXcalendar = "julian" // julian
    BarXcalendar_mayan BarXcalendar = "mayan" // mayan
    BarXcalendar_nanakshahi BarXcalendar = "nanakshahi" // nanakshahi
    BarXcalendar_nepali BarXcalendar = "nepali" // nepali
    BarXcalendar_persian BarXcalendar = "persian" // persian
    BarXcalendar_jalali BarXcalendar = "jalali" // jalali
    BarXcalendar_taiwan BarXcalendar = "taiwan" // taiwan
    BarXcalendar_thai BarXcalendar = "thai" // thai
    BarXcalendar_ummalqura BarXcalendar = "ummalqura" // ummalqura
)

// BarYcalendar Sets the calendar system to use with `y` date data.
type BarYcalendar string

const (
    BarYcalendar_gregorian BarYcalendar = "gregorian" // gregorian
    BarYcalendar_chinese BarYcalendar = "chinese" // chinese
    BarYcalendar_coptic BarYcalendar = "coptic" // coptic
    BarYcalendar_discworld BarYcalendar = "discworld" // discworld
    BarYcalendar_ethiopian BarYcalendar = "ethiopian" // ethiopian
    BarYcalendar_hebrew BarYcalendar = "hebrew" // hebrew
    BarYcalendar_islamic BarYcalendar = "islamic" // islamic
    BarYcalendar_julian BarYcalendar = "julian" // julian
    BarYcalendar_mayan BarYcalendar = "mayan" // mayan
    BarYcalendar_nanakshahi BarYcalendar = "nanakshahi" // nanakshahi
    BarYcalendar_nepali BarYcalendar = "nepali" // nepali
    BarYcalendar_persian BarYcalendar = "persian" // persian
    BarYcalendar_jalali BarYcalendar = "jalali" // jalali
    BarYcalendar_taiwan BarYcalendar = "taiwan" // taiwan
    BarYcalendar_thai BarYcalendar = "thai" // thai
    BarYcalendar_ummalqura BarYcalendar = "ummalqura" // ummalqura
)

// BarpolarHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type BarpolarHoverlabelAlign string

const (
    BarpolarHoverlabelAlign_left BarpolarHoverlabelAlign = "left" // left
    BarpolarHoverlabelAlign_right BarpolarHoverlabelAlign = "right" // right
    BarpolarHoverlabelAlign_auto BarpolarHoverlabelAlign = "auto" // auto
)

// BarpolarMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type BarpolarMarkerColorbarExponentformat string

const (
    BarpolarMarkerColorbarExponentformat_none BarpolarMarkerColorbarExponentformat = "none" // none
    BarpolarMarkerColorbarExponentformat_e BarpolarMarkerColorbarExponentformat = "e" // e
    BarpolarMarkerColorbarExponentformat_E BarpolarMarkerColorbarExponentformat = "E" // E
    BarpolarMarkerColorbarExponentformat_power BarpolarMarkerColorbarExponentformat = "power" // power
    BarpolarMarkerColorbarExponentformat_SI BarpolarMarkerColorbarExponentformat = "SI" // SI
    BarpolarMarkerColorbarExponentformat_B BarpolarMarkerColorbarExponentformat = "B" // B
)

// BarpolarMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type BarpolarMarkerColorbarLenmode string

const (
    BarpolarMarkerColorbarLenmode_fraction BarpolarMarkerColorbarLenmode = "fraction" // fraction
    BarpolarMarkerColorbarLenmode_pixels BarpolarMarkerColorbarLenmode = "pixels" // pixels
)

// BarpolarMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type BarpolarMarkerColorbarShowexponent string

const (
    BarpolarMarkerColorbarShowexponent_all BarpolarMarkerColorbarShowexponent = "all" // all
    BarpolarMarkerColorbarShowexponent_first BarpolarMarkerColorbarShowexponent = "first" // first
    BarpolarMarkerColorbarShowexponent_last BarpolarMarkerColorbarShowexponent = "last" // last
    BarpolarMarkerColorbarShowexponent_none BarpolarMarkerColorbarShowexponent = "none" // none
)

// BarpolarMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type BarpolarMarkerColorbarShowtickprefix string

const (
    BarpolarMarkerColorbarShowtickprefix_all BarpolarMarkerColorbarShowtickprefix = "all" // all
    BarpolarMarkerColorbarShowtickprefix_first BarpolarMarkerColorbarShowtickprefix = "first" // first
    BarpolarMarkerColorbarShowtickprefix_last BarpolarMarkerColorbarShowtickprefix = "last" // last
    BarpolarMarkerColorbarShowtickprefix_none BarpolarMarkerColorbarShowtickprefix = "none" // none
)

// BarpolarMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type BarpolarMarkerColorbarShowticksuffix string

const (
    BarpolarMarkerColorbarShowticksuffix_all BarpolarMarkerColorbarShowticksuffix = "all" // all
    BarpolarMarkerColorbarShowticksuffix_first BarpolarMarkerColorbarShowticksuffix = "first" // first
    BarpolarMarkerColorbarShowticksuffix_last BarpolarMarkerColorbarShowticksuffix = "last" // last
    BarpolarMarkerColorbarShowticksuffix_none BarpolarMarkerColorbarShowticksuffix = "none" // none
)

// BarpolarMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type BarpolarMarkerColorbarThicknessmode string

const (
    BarpolarMarkerColorbarThicknessmode_fraction BarpolarMarkerColorbarThicknessmode = "fraction" // fraction
    BarpolarMarkerColorbarThicknessmode_pixels BarpolarMarkerColorbarThicknessmode = "pixels" // pixels
)

// BarpolarMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type BarpolarMarkerColorbarTickmode string

const (
    BarpolarMarkerColorbarTickmode_auto BarpolarMarkerColorbarTickmode = "auto" // auto
    BarpolarMarkerColorbarTickmode_linear BarpolarMarkerColorbarTickmode = "linear" // linear
    BarpolarMarkerColorbarTickmode_array BarpolarMarkerColorbarTickmode = "array" // array
)

// BarpolarMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type BarpolarMarkerColorbarTicks string

const (
    BarpolarMarkerColorbarTicks_outside BarpolarMarkerColorbarTicks = "outside" // outside
    BarpolarMarkerColorbarTicks_inside BarpolarMarkerColorbarTicks = "inside" // inside
)

// BarpolarMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type BarpolarMarkerColorbarTitleSide string

const (
    BarpolarMarkerColorbarTitleSide_right BarpolarMarkerColorbarTitleSide = "right" // right
    BarpolarMarkerColorbarTitleSide_top BarpolarMarkerColorbarTitleSide = "top" // top
    BarpolarMarkerColorbarTitleSide_bottom BarpolarMarkerColorbarTitleSide = "bottom" // bottom
)

// BarpolarMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type BarpolarMarkerColorbarXanchor string

const (
    BarpolarMarkerColorbarXanchor_left BarpolarMarkerColorbarXanchor = "left" // left
    BarpolarMarkerColorbarXanchor_center BarpolarMarkerColorbarXanchor = "center" // center
    BarpolarMarkerColorbarXanchor_right BarpolarMarkerColorbarXanchor = "right" // right
)

// BarpolarMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type BarpolarMarkerColorbarYanchor string

const (
    BarpolarMarkerColorbarYanchor_top BarpolarMarkerColorbarYanchor = "top" // top
    BarpolarMarkerColorbarYanchor_middle BarpolarMarkerColorbarYanchor = "middle" // middle
    BarpolarMarkerColorbarYanchor_bottom BarpolarMarkerColorbarYanchor = "bottom" // bottom
)

// BarpolarThetaunit Sets the unit of input *theta* values. Has an effect only when on *linear* angular axes.
type BarpolarThetaunit string

const (
    BarpolarThetaunit_radians BarpolarThetaunit = "radians" // radians
    BarpolarThetaunit_degrees BarpolarThetaunit = "degrees" // degrees
    BarpolarThetaunit_gradians BarpolarThetaunit = "gradians" // gradians
)

// BarpolarVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type BarpolarVisible interface{}

const (
    BarpolarVisibleTrue bool = true
    BarpolarVisibleFalse bool = false
    BarpolarVisibleLegendonly string = "legendonly"
)

// BoxBoxmean If *true*, the mean of the box(es)' underlying distribution is drawn as a dashed line inside the box(es). If *sd* the standard deviation is also drawn. Defaults to *true* when `mean` is set. Defaults to *sd* when `sd` is set Otherwise defaults to *false*.
type BoxBoxmean interface{}

const (
    BoxBoxmeanTrue bool = true
    BoxBoxmeanSd string = "sd"
    BoxBoxmeanFalse bool = false
)

// BoxBoxpoints If *outliers*, only the sample points lying outside the whiskers are shown If *suspectedoutliers*, the outlier points are shown and points either less than 4*Q1-3*Q3 or greater than 4*Q3-3*Q1 are highlighted (see `outliercolor`) If *all*, all sample points are shown If *false*, only the box(es) are shown with no sample points Defaults to *suspectedoutliers* when `marker.outliercolor` or `marker.line.outliercolor` is set. Defaults to *all* under the q1/median/q3 signature. Otherwise defaults to *outliers*.
type BoxBoxpoints interface{}

const (
    BoxBoxpointsAll string = "all"
    BoxBoxpointsOutliers string = "outliers"
    BoxBoxpointsSuspectedoutliers string = "suspectedoutliers"
    BoxBoxpointsFalse bool = false
)

// BoxHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type BoxHoverlabelAlign string

const (
    BoxHoverlabelAlign_left BoxHoverlabelAlign = "left" // left
    BoxHoverlabelAlign_right BoxHoverlabelAlign = "right" // right
    BoxHoverlabelAlign_auto BoxHoverlabelAlign = "auto" // auto
)

// BoxMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type BoxMarkerSymbol string

const (
    BoxMarkerSymbol0 BoxMarkerSymbol = "0" // 0
    BoxMarkerSymbol_circle BoxMarkerSymbol = "circle" // circle
    BoxMarkerSymbol100 BoxMarkerSymbol = "100" // 100
    BoxMarkerSymbol_circle_open BoxMarkerSymbol = "circle-open" // circle-open
    BoxMarkerSymbol200 BoxMarkerSymbol = "200" // 200
    BoxMarkerSymbol_circle_dot BoxMarkerSymbol = "circle-dot" // circle-dot
    BoxMarkerSymbol300 BoxMarkerSymbol = "300" // 300
    BoxMarkerSymbol_circle_open_dot BoxMarkerSymbol = "circle-open-dot" // circle-open-dot
    BoxMarkerSymbol1 BoxMarkerSymbol = "1" // 1
    BoxMarkerSymbol_square BoxMarkerSymbol = "square" // square
    BoxMarkerSymbol101 BoxMarkerSymbol = "101" // 101
    BoxMarkerSymbol_square_open BoxMarkerSymbol = "square-open" // square-open
    BoxMarkerSymbol201 BoxMarkerSymbol = "201" // 201
    BoxMarkerSymbol_square_dot BoxMarkerSymbol = "square-dot" // square-dot
    BoxMarkerSymbol301 BoxMarkerSymbol = "301" // 301
    BoxMarkerSymbol_square_open_dot BoxMarkerSymbol = "square-open-dot" // square-open-dot
    BoxMarkerSymbol2 BoxMarkerSymbol = "2" // 2
    BoxMarkerSymbol_diamond BoxMarkerSymbol = "diamond" // diamond
    BoxMarkerSymbol102 BoxMarkerSymbol = "102" // 102
    BoxMarkerSymbol_diamond_open BoxMarkerSymbol = "diamond-open" // diamond-open
    BoxMarkerSymbol202 BoxMarkerSymbol = "202" // 202
    BoxMarkerSymbol_diamond_dot BoxMarkerSymbol = "diamond-dot" // diamond-dot
    BoxMarkerSymbol302 BoxMarkerSymbol = "302" // 302
    BoxMarkerSymbol_diamond_open_dot BoxMarkerSymbol = "diamond-open-dot" // diamond-open-dot
    BoxMarkerSymbol3 BoxMarkerSymbol = "3" // 3
    BoxMarkerSymbol_cross BoxMarkerSymbol = "cross" // cross
    BoxMarkerSymbol103 BoxMarkerSymbol = "103" // 103
    BoxMarkerSymbol_cross_open BoxMarkerSymbol = "cross-open" // cross-open
    BoxMarkerSymbol203 BoxMarkerSymbol = "203" // 203
    BoxMarkerSymbol_cross_dot BoxMarkerSymbol = "cross-dot" // cross-dot
    BoxMarkerSymbol303 BoxMarkerSymbol = "303" // 303
    BoxMarkerSymbol_cross_open_dot BoxMarkerSymbol = "cross-open-dot" // cross-open-dot
    BoxMarkerSymbol4 BoxMarkerSymbol = "4" // 4
    BoxMarkerSymbol_x BoxMarkerSymbol = "x" // x
    BoxMarkerSymbol104 BoxMarkerSymbol = "104" // 104
    BoxMarkerSymbol_x_open BoxMarkerSymbol = "x-open" // x-open
    BoxMarkerSymbol204 BoxMarkerSymbol = "204" // 204
    BoxMarkerSymbol_x_dot BoxMarkerSymbol = "x-dot" // x-dot
    BoxMarkerSymbol304 BoxMarkerSymbol = "304" // 304
    BoxMarkerSymbol_x_open_dot BoxMarkerSymbol = "x-open-dot" // x-open-dot
    BoxMarkerSymbol5 BoxMarkerSymbol = "5" // 5
    BoxMarkerSymbol_triangle_up BoxMarkerSymbol = "triangle-up" // triangle-up
    BoxMarkerSymbol105 BoxMarkerSymbol = "105" // 105
    BoxMarkerSymbol_triangle_up_open BoxMarkerSymbol = "triangle-up-open" // triangle-up-open
    BoxMarkerSymbol205 BoxMarkerSymbol = "205" // 205
    BoxMarkerSymbol_triangle_up_dot BoxMarkerSymbol = "triangle-up-dot" // triangle-up-dot
    BoxMarkerSymbol305 BoxMarkerSymbol = "305" // 305
    BoxMarkerSymbol_triangle_up_open_dot BoxMarkerSymbol = "triangle-up-open-dot" // triangle-up-open-dot
    BoxMarkerSymbol6 BoxMarkerSymbol = "6" // 6
    BoxMarkerSymbol_triangle_down BoxMarkerSymbol = "triangle-down" // triangle-down
    BoxMarkerSymbol106 BoxMarkerSymbol = "106" // 106
    BoxMarkerSymbol_triangle_down_open BoxMarkerSymbol = "triangle-down-open" // triangle-down-open
    BoxMarkerSymbol206 BoxMarkerSymbol = "206" // 206
    BoxMarkerSymbol_triangle_down_dot BoxMarkerSymbol = "triangle-down-dot" // triangle-down-dot
    BoxMarkerSymbol306 BoxMarkerSymbol = "306" // 306
    BoxMarkerSymbol_triangle_down_open_dot BoxMarkerSymbol = "triangle-down-open-dot" // triangle-down-open-dot
    BoxMarkerSymbol7 BoxMarkerSymbol = "7" // 7
    BoxMarkerSymbol_triangle_left BoxMarkerSymbol = "triangle-left" // triangle-left
    BoxMarkerSymbol107 BoxMarkerSymbol = "107" // 107
    BoxMarkerSymbol_triangle_left_open BoxMarkerSymbol = "triangle-left-open" // triangle-left-open
    BoxMarkerSymbol207 BoxMarkerSymbol = "207" // 207
    BoxMarkerSymbol_triangle_left_dot BoxMarkerSymbol = "triangle-left-dot" // triangle-left-dot
    BoxMarkerSymbol307 BoxMarkerSymbol = "307" // 307
    BoxMarkerSymbol_triangle_left_open_dot BoxMarkerSymbol = "triangle-left-open-dot" // triangle-left-open-dot
    BoxMarkerSymbol8 BoxMarkerSymbol = "8" // 8
    BoxMarkerSymbol_triangle_right BoxMarkerSymbol = "triangle-right" // triangle-right
    BoxMarkerSymbol108 BoxMarkerSymbol = "108" // 108
    BoxMarkerSymbol_triangle_right_open BoxMarkerSymbol = "triangle-right-open" // triangle-right-open
    BoxMarkerSymbol208 BoxMarkerSymbol = "208" // 208
    BoxMarkerSymbol_triangle_right_dot BoxMarkerSymbol = "triangle-right-dot" // triangle-right-dot
    BoxMarkerSymbol308 BoxMarkerSymbol = "308" // 308
    BoxMarkerSymbol_triangle_right_open_dot BoxMarkerSymbol = "triangle-right-open-dot" // triangle-right-open-dot
    BoxMarkerSymbol9 BoxMarkerSymbol = "9" // 9
    BoxMarkerSymbol_triangle_ne BoxMarkerSymbol = "triangle-ne" // triangle-ne
    BoxMarkerSymbol109 BoxMarkerSymbol = "109" // 109
    BoxMarkerSymbol_triangle_ne_open BoxMarkerSymbol = "triangle-ne-open" // triangle-ne-open
    BoxMarkerSymbol209 BoxMarkerSymbol = "209" // 209
    BoxMarkerSymbol_triangle_ne_dot BoxMarkerSymbol = "triangle-ne-dot" // triangle-ne-dot
    BoxMarkerSymbol309 BoxMarkerSymbol = "309" // 309
    BoxMarkerSymbol_triangle_ne_open_dot BoxMarkerSymbol = "triangle-ne-open-dot" // triangle-ne-open-dot
    BoxMarkerSymbol10 BoxMarkerSymbol = "10" // 10
    BoxMarkerSymbol_triangle_se BoxMarkerSymbol = "triangle-se" // triangle-se
    BoxMarkerSymbol110 BoxMarkerSymbol = "110" // 110
    BoxMarkerSymbol_triangle_se_open BoxMarkerSymbol = "triangle-se-open" // triangle-se-open
    BoxMarkerSymbol210 BoxMarkerSymbol = "210" // 210
    BoxMarkerSymbol_triangle_se_dot BoxMarkerSymbol = "triangle-se-dot" // triangle-se-dot
    BoxMarkerSymbol310 BoxMarkerSymbol = "310" // 310
    BoxMarkerSymbol_triangle_se_open_dot BoxMarkerSymbol = "triangle-se-open-dot" // triangle-se-open-dot
    BoxMarkerSymbol11 BoxMarkerSymbol = "11" // 11
    BoxMarkerSymbol_triangle_sw BoxMarkerSymbol = "triangle-sw" // triangle-sw
    BoxMarkerSymbol111 BoxMarkerSymbol = "111" // 111
    BoxMarkerSymbol_triangle_sw_open BoxMarkerSymbol = "triangle-sw-open" // triangle-sw-open
    BoxMarkerSymbol211 BoxMarkerSymbol = "211" // 211
    BoxMarkerSymbol_triangle_sw_dot BoxMarkerSymbol = "triangle-sw-dot" // triangle-sw-dot
    BoxMarkerSymbol311 BoxMarkerSymbol = "311" // 311
    BoxMarkerSymbol_triangle_sw_open_dot BoxMarkerSymbol = "triangle-sw-open-dot" // triangle-sw-open-dot
    BoxMarkerSymbol12 BoxMarkerSymbol = "12" // 12
    BoxMarkerSymbol_triangle_nw BoxMarkerSymbol = "triangle-nw" // triangle-nw
    BoxMarkerSymbol112 BoxMarkerSymbol = "112" // 112
    BoxMarkerSymbol_triangle_nw_open BoxMarkerSymbol = "triangle-nw-open" // triangle-nw-open
    BoxMarkerSymbol212 BoxMarkerSymbol = "212" // 212
    BoxMarkerSymbol_triangle_nw_dot BoxMarkerSymbol = "triangle-nw-dot" // triangle-nw-dot
    BoxMarkerSymbol312 BoxMarkerSymbol = "312" // 312
    BoxMarkerSymbol_triangle_nw_open_dot BoxMarkerSymbol = "triangle-nw-open-dot" // triangle-nw-open-dot
    BoxMarkerSymbol13 BoxMarkerSymbol = "13" // 13
    BoxMarkerSymbol_pentagon BoxMarkerSymbol = "pentagon" // pentagon
    BoxMarkerSymbol113 BoxMarkerSymbol = "113" // 113
    BoxMarkerSymbol_pentagon_open BoxMarkerSymbol = "pentagon-open" // pentagon-open
    BoxMarkerSymbol213 BoxMarkerSymbol = "213" // 213
    BoxMarkerSymbol_pentagon_dot BoxMarkerSymbol = "pentagon-dot" // pentagon-dot
    BoxMarkerSymbol313 BoxMarkerSymbol = "313" // 313
    BoxMarkerSymbol_pentagon_open_dot BoxMarkerSymbol = "pentagon-open-dot" // pentagon-open-dot
    BoxMarkerSymbol14 BoxMarkerSymbol = "14" // 14
    BoxMarkerSymbol_hexagon BoxMarkerSymbol = "hexagon" // hexagon
    BoxMarkerSymbol114 BoxMarkerSymbol = "114" // 114
    BoxMarkerSymbol_hexagon_open BoxMarkerSymbol = "hexagon-open" // hexagon-open
    BoxMarkerSymbol214 BoxMarkerSymbol = "214" // 214
    BoxMarkerSymbol_hexagon_dot BoxMarkerSymbol = "hexagon-dot" // hexagon-dot
    BoxMarkerSymbol314 BoxMarkerSymbol = "314" // 314
    BoxMarkerSymbol_hexagon_open_dot BoxMarkerSymbol = "hexagon-open-dot" // hexagon-open-dot
    BoxMarkerSymbol15 BoxMarkerSymbol = "15" // 15
    BoxMarkerSymbol_hexagon2 BoxMarkerSymbol = "hexagon2" // hexagon2
    BoxMarkerSymbol115 BoxMarkerSymbol = "115" // 115
    BoxMarkerSymbol_hexagon2_open BoxMarkerSymbol = "hexagon2-open" // hexagon2-open
    BoxMarkerSymbol215 BoxMarkerSymbol = "215" // 215
    BoxMarkerSymbol_hexagon2_dot BoxMarkerSymbol = "hexagon2-dot" // hexagon2-dot
    BoxMarkerSymbol315 BoxMarkerSymbol = "315" // 315
    BoxMarkerSymbol_hexagon2_open_dot BoxMarkerSymbol = "hexagon2-open-dot" // hexagon2-open-dot
    BoxMarkerSymbol16 BoxMarkerSymbol = "16" // 16
    BoxMarkerSymbol_octagon BoxMarkerSymbol = "octagon" // octagon
    BoxMarkerSymbol116 BoxMarkerSymbol = "116" // 116
    BoxMarkerSymbol_octagon_open BoxMarkerSymbol = "octagon-open" // octagon-open
    BoxMarkerSymbol216 BoxMarkerSymbol = "216" // 216
    BoxMarkerSymbol_octagon_dot BoxMarkerSymbol = "octagon-dot" // octagon-dot
    BoxMarkerSymbol316 BoxMarkerSymbol = "316" // 316
    BoxMarkerSymbol_octagon_open_dot BoxMarkerSymbol = "octagon-open-dot" // octagon-open-dot
    BoxMarkerSymbol17 BoxMarkerSymbol = "17" // 17
    BoxMarkerSymbol_star BoxMarkerSymbol = "star" // star
    BoxMarkerSymbol117 BoxMarkerSymbol = "117" // 117
    BoxMarkerSymbol_star_open BoxMarkerSymbol = "star-open" // star-open
    BoxMarkerSymbol217 BoxMarkerSymbol = "217" // 217
    BoxMarkerSymbol_star_dot BoxMarkerSymbol = "star-dot" // star-dot
    BoxMarkerSymbol317 BoxMarkerSymbol = "317" // 317
    BoxMarkerSymbol_star_open_dot BoxMarkerSymbol = "star-open-dot" // star-open-dot
    BoxMarkerSymbol18 BoxMarkerSymbol = "18" // 18
    BoxMarkerSymbol_hexagram BoxMarkerSymbol = "hexagram" // hexagram
    BoxMarkerSymbol118 BoxMarkerSymbol = "118" // 118
    BoxMarkerSymbol_hexagram_open BoxMarkerSymbol = "hexagram-open" // hexagram-open
    BoxMarkerSymbol218 BoxMarkerSymbol = "218" // 218
    BoxMarkerSymbol_hexagram_dot BoxMarkerSymbol = "hexagram-dot" // hexagram-dot
    BoxMarkerSymbol318 BoxMarkerSymbol = "318" // 318
    BoxMarkerSymbol_hexagram_open_dot BoxMarkerSymbol = "hexagram-open-dot" // hexagram-open-dot
    BoxMarkerSymbol19 BoxMarkerSymbol = "19" // 19
    BoxMarkerSymbol_star_triangle_up BoxMarkerSymbol = "star-triangle-up" // star-triangle-up
    BoxMarkerSymbol119 BoxMarkerSymbol = "119" // 119
    BoxMarkerSymbol_star_triangle_up_open BoxMarkerSymbol = "star-triangle-up-open" // star-triangle-up-open
    BoxMarkerSymbol219 BoxMarkerSymbol = "219" // 219
    BoxMarkerSymbol_star_triangle_up_dot BoxMarkerSymbol = "star-triangle-up-dot" // star-triangle-up-dot
    BoxMarkerSymbol319 BoxMarkerSymbol = "319" // 319
    BoxMarkerSymbol_star_triangle_up_open_dot BoxMarkerSymbol = "star-triangle-up-open-dot" // star-triangle-up-open-dot
    BoxMarkerSymbol20 BoxMarkerSymbol = "20" // 20
    BoxMarkerSymbol_star_triangle_down BoxMarkerSymbol = "star-triangle-down" // star-triangle-down
    BoxMarkerSymbol120 BoxMarkerSymbol = "120" // 120
    BoxMarkerSymbol_star_triangle_down_open BoxMarkerSymbol = "star-triangle-down-open" // star-triangle-down-open
    BoxMarkerSymbol220 BoxMarkerSymbol = "220" // 220
    BoxMarkerSymbol_star_triangle_down_dot BoxMarkerSymbol = "star-triangle-down-dot" // star-triangle-down-dot
    BoxMarkerSymbol320 BoxMarkerSymbol = "320" // 320
    BoxMarkerSymbol_star_triangle_down_open_dot BoxMarkerSymbol = "star-triangle-down-open-dot" // star-triangle-down-open-dot
    BoxMarkerSymbol21 BoxMarkerSymbol = "21" // 21
    BoxMarkerSymbol_star_square BoxMarkerSymbol = "star-square" // star-square
    BoxMarkerSymbol121 BoxMarkerSymbol = "121" // 121
    BoxMarkerSymbol_star_square_open BoxMarkerSymbol = "star-square-open" // star-square-open
    BoxMarkerSymbol221 BoxMarkerSymbol = "221" // 221
    BoxMarkerSymbol_star_square_dot BoxMarkerSymbol = "star-square-dot" // star-square-dot
    BoxMarkerSymbol321 BoxMarkerSymbol = "321" // 321
    BoxMarkerSymbol_star_square_open_dot BoxMarkerSymbol = "star-square-open-dot" // star-square-open-dot
    BoxMarkerSymbol22 BoxMarkerSymbol = "22" // 22
    BoxMarkerSymbol_star_diamond BoxMarkerSymbol = "star-diamond" // star-diamond
    BoxMarkerSymbol122 BoxMarkerSymbol = "122" // 122
    BoxMarkerSymbol_star_diamond_open BoxMarkerSymbol = "star-diamond-open" // star-diamond-open
    BoxMarkerSymbol222 BoxMarkerSymbol = "222" // 222
    BoxMarkerSymbol_star_diamond_dot BoxMarkerSymbol = "star-diamond-dot" // star-diamond-dot
    BoxMarkerSymbol322 BoxMarkerSymbol = "322" // 322
    BoxMarkerSymbol_star_diamond_open_dot BoxMarkerSymbol = "star-diamond-open-dot" // star-diamond-open-dot
    BoxMarkerSymbol23 BoxMarkerSymbol = "23" // 23
    BoxMarkerSymbol_diamond_tall BoxMarkerSymbol = "diamond-tall" // diamond-tall
    BoxMarkerSymbol123 BoxMarkerSymbol = "123" // 123
    BoxMarkerSymbol_diamond_tall_open BoxMarkerSymbol = "diamond-tall-open" // diamond-tall-open
    BoxMarkerSymbol223 BoxMarkerSymbol = "223" // 223
    BoxMarkerSymbol_diamond_tall_dot BoxMarkerSymbol = "diamond-tall-dot" // diamond-tall-dot
    BoxMarkerSymbol323 BoxMarkerSymbol = "323" // 323
    BoxMarkerSymbol_diamond_tall_open_dot BoxMarkerSymbol = "diamond-tall-open-dot" // diamond-tall-open-dot
    BoxMarkerSymbol24 BoxMarkerSymbol = "24" // 24
    BoxMarkerSymbol_diamond_wide BoxMarkerSymbol = "diamond-wide" // diamond-wide
    BoxMarkerSymbol124 BoxMarkerSymbol = "124" // 124
    BoxMarkerSymbol_diamond_wide_open BoxMarkerSymbol = "diamond-wide-open" // diamond-wide-open
    BoxMarkerSymbol224 BoxMarkerSymbol = "224" // 224
    BoxMarkerSymbol_diamond_wide_dot BoxMarkerSymbol = "diamond-wide-dot" // diamond-wide-dot
    BoxMarkerSymbol324 BoxMarkerSymbol = "324" // 324
    BoxMarkerSymbol_diamond_wide_open_dot BoxMarkerSymbol = "diamond-wide-open-dot" // diamond-wide-open-dot
    BoxMarkerSymbol25 BoxMarkerSymbol = "25" // 25
    BoxMarkerSymbol_hourglass BoxMarkerSymbol = "hourglass" // hourglass
    BoxMarkerSymbol125 BoxMarkerSymbol = "125" // 125
    BoxMarkerSymbol_hourglass_open BoxMarkerSymbol = "hourglass-open" // hourglass-open
    BoxMarkerSymbol26 BoxMarkerSymbol = "26" // 26
    BoxMarkerSymbol_bowtie BoxMarkerSymbol = "bowtie" // bowtie
    BoxMarkerSymbol126 BoxMarkerSymbol = "126" // 126
    BoxMarkerSymbol_bowtie_open BoxMarkerSymbol = "bowtie-open" // bowtie-open
    BoxMarkerSymbol27 BoxMarkerSymbol = "27" // 27
    BoxMarkerSymbol_circle_cross BoxMarkerSymbol = "circle-cross" // circle-cross
    BoxMarkerSymbol127 BoxMarkerSymbol = "127" // 127
    BoxMarkerSymbol_circle_cross_open BoxMarkerSymbol = "circle-cross-open" // circle-cross-open
    BoxMarkerSymbol28 BoxMarkerSymbol = "28" // 28
    BoxMarkerSymbol_circle_x BoxMarkerSymbol = "circle-x" // circle-x
    BoxMarkerSymbol128 BoxMarkerSymbol = "128" // 128
    BoxMarkerSymbol_circle_x_open BoxMarkerSymbol = "circle-x-open" // circle-x-open
    BoxMarkerSymbol29 BoxMarkerSymbol = "29" // 29
    BoxMarkerSymbol_square_cross BoxMarkerSymbol = "square-cross" // square-cross
    BoxMarkerSymbol129 BoxMarkerSymbol = "129" // 129
    BoxMarkerSymbol_square_cross_open BoxMarkerSymbol = "square-cross-open" // square-cross-open
    BoxMarkerSymbol30 BoxMarkerSymbol = "30" // 30
    BoxMarkerSymbol_square_x BoxMarkerSymbol = "square-x" // square-x
    BoxMarkerSymbol130 BoxMarkerSymbol = "130" // 130
    BoxMarkerSymbol_square_x_open BoxMarkerSymbol = "square-x-open" // square-x-open
    BoxMarkerSymbol31 BoxMarkerSymbol = "31" // 31
    BoxMarkerSymbol_diamond_cross BoxMarkerSymbol = "diamond-cross" // diamond-cross
    BoxMarkerSymbol131 BoxMarkerSymbol = "131" // 131
    BoxMarkerSymbol_diamond_cross_open BoxMarkerSymbol = "diamond-cross-open" // diamond-cross-open
    BoxMarkerSymbol32 BoxMarkerSymbol = "32" // 32
    BoxMarkerSymbol_diamond_x BoxMarkerSymbol = "diamond-x" // diamond-x
    BoxMarkerSymbol132 BoxMarkerSymbol = "132" // 132
    BoxMarkerSymbol_diamond_x_open BoxMarkerSymbol = "diamond-x-open" // diamond-x-open
    BoxMarkerSymbol33 BoxMarkerSymbol = "33" // 33
    BoxMarkerSymbol_cross_thin BoxMarkerSymbol = "cross-thin" // cross-thin
    BoxMarkerSymbol133 BoxMarkerSymbol = "133" // 133
    BoxMarkerSymbol_cross_thin_open BoxMarkerSymbol = "cross-thin-open" // cross-thin-open
    BoxMarkerSymbol34 BoxMarkerSymbol = "34" // 34
    BoxMarkerSymbol_x_thin BoxMarkerSymbol = "x-thin" // x-thin
    BoxMarkerSymbol134 BoxMarkerSymbol = "134" // 134
    BoxMarkerSymbol_x_thin_open BoxMarkerSymbol = "x-thin-open" // x-thin-open
    BoxMarkerSymbol35 BoxMarkerSymbol = "35" // 35
    BoxMarkerSymbol_asterisk BoxMarkerSymbol = "asterisk" // asterisk
    BoxMarkerSymbol135 BoxMarkerSymbol = "135" // 135
    BoxMarkerSymbol_asterisk_open BoxMarkerSymbol = "asterisk-open" // asterisk-open
    BoxMarkerSymbol36 BoxMarkerSymbol = "36" // 36
    BoxMarkerSymbol_hash BoxMarkerSymbol = "hash" // hash
    BoxMarkerSymbol136 BoxMarkerSymbol = "136" // 136
    BoxMarkerSymbol_hash_open BoxMarkerSymbol = "hash-open" // hash-open
    BoxMarkerSymbol236 BoxMarkerSymbol = "236" // 236
    BoxMarkerSymbol_hash_dot BoxMarkerSymbol = "hash-dot" // hash-dot
    BoxMarkerSymbol336 BoxMarkerSymbol = "336" // 336
    BoxMarkerSymbol_hash_open_dot BoxMarkerSymbol = "hash-open-dot" // hash-open-dot
    BoxMarkerSymbol37 BoxMarkerSymbol = "37" // 37
    BoxMarkerSymbol_y_up BoxMarkerSymbol = "y-up" // y-up
    BoxMarkerSymbol137 BoxMarkerSymbol = "137" // 137
    BoxMarkerSymbol_y_up_open BoxMarkerSymbol = "y-up-open" // y-up-open
    BoxMarkerSymbol38 BoxMarkerSymbol = "38" // 38
    BoxMarkerSymbol_y_down BoxMarkerSymbol = "y-down" // y-down
    BoxMarkerSymbol138 BoxMarkerSymbol = "138" // 138
    BoxMarkerSymbol_y_down_open BoxMarkerSymbol = "y-down-open" // y-down-open
    BoxMarkerSymbol39 BoxMarkerSymbol = "39" // 39
    BoxMarkerSymbol_y_left BoxMarkerSymbol = "y-left" // y-left
    BoxMarkerSymbol139 BoxMarkerSymbol = "139" // 139
    BoxMarkerSymbol_y_left_open BoxMarkerSymbol = "y-left-open" // y-left-open
    BoxMarkerSymbol40 BoxMarkerSymbol = "40" // 40
    BoxMarkerSymbol_y_right BoxMarkerSymbol = "y-right" // y-right
    BoxMarkerSymbol140 BoxMarkerSymbol = "140" // 140
    BoxMarkerSymbol_y_right_open BoxMarkerSymbol = "y-right-open" // y-right-open
    BoxMarkerSymbol41 BoxMarkerSymbol = "41" // 41
    BoxMarkerSymbol_line_ew BoxMarkerSymbol = "line-ew" // line-ew
    BoxMarkerSymbol141 BoxMarkerSymbol = "141" // 141
    BoxMarkerSymbol_line_ew_open BoxMarkerSymbol = "line-ew-open" // line-ew-open
    BoxMarkerSymbol42 BoxMarkerSymbol = "42" // 42
    BoxMarkerSymbol_line_ns BoxMarkerSymbol = "line-ns" // line-ns
    BoxMarkerSymbol142 BoxMarkerSymbol = "142" // 142
    BoxMarkerSymbol_line_ns_open BoxMarkerSymbol = "line-ns-open" // line-ns-open
    BoxMarkerSymbol43 BoxMarkerSymbol = "43" // 43
    BoxMarkerSymbol_line_ne BoxMarkerSymbol = "line-ne" // line-ne
    BoxMarkerSymbol143 BoxMarkerSymbol = "143" // 143
    BoxMarkerSymbol_line_ne_open BoxMarkerSymbol = "line-ne-open" // line-ne-open
    BoxMarkerSymbol44 BoxMarkerSymbol = "44" // 44
    BoxMarkerSymbol_line_nw BoxMarkerSymbol = "line-nw" // line-nw
    BoxMarkerSymbol144 BoxMarkerSymbol = "144" // 144
    BoxMarkerSymbol_line_nw_open BoxMarkerSymbol = "line-nw-open" // line-nw-open
)

// BoxOrientation Sets the orientation of the box(es). If *v* (*h*), the distribution is visualized along the vertical (horizontal).
type BoxOrientation string

const (
    BoxOrientation_v BoxOrientation = "v" // v
    BoxOrientation_h BoxOrientation = "h" // h
)

// BoxQuartilemethod Sets the method used to compute the sample's Q1 and Q3 quartiles. The *linear* method uses the 25th percentile for Q1 and 75th percentile for Q3 as computed using method #10 (listed on http://www.amstat.org/publications/jse/v14n3/langford.html). The *exclusive* method uses the median to divide the ordered dataset into two halves if the sample is odd, it does not include the median in either half - Q1 is then the median of the lower half and Q3 the median of the upper half. The *inclusive* method also uses the median to divide the ordered dataset into two halves but if the sample is odd, it includes the median in both halves - Q1 is then the median of the lower half and Q3 the median of the upper half.
type BoxQuartilemethod string

const (
    BoxQuartilemethod_linear BoxQuartilemethod = "linear" // linear
    BoxQuartilemethod_exclusive BoxQuartilemethod = "exclusive" // exclusive
    BoxQuartilemethod_inclusive BoxQuartilemethod = "inclusive" // inclusive
)

// BoxVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type BoxVisible interface{}

const (
    BoxVisibleTrue bool = true
    BoxVisibleFalse bool = false
    BoxVisibleLegendonly string = "legendonly"
)

// BoxXcalendar Sets the calendar system to use with `x` date data.
type BoxXcalendar string

const (
    BoxXcalendar_gregorian BoxXcalendar = "gregorian" // gregorian
    BoxXcalendar_chinese BoxXcalendar = "chinese" // chinese
    BoxXcalendar_coptic BoxXcalendar = "coptic" // coptic
    BoxXcalendar_discworld BoxXcalendar = "discworld" // discworld
    BoxXcalendar_ethiopian BoxXcalendar = "ethiopian" // ethiopian
    BoxXcalendar_hebrew BoxXcalendar = "hebrew" // hebrew
    BoxXcalendar_islamic BoxXcalendar = "islamic" // islamic
    BoxXcalendar_julian BoxXcalendar = "julian" // julian
    BoxXcalendar_mayan BoxXcalendar = "mayan" // mayan
    BoxXcalendar_nanakshahi BoxXcalendar = "nanakshahi" // nanakshahi
    BoxXcalendar_nepali BoxXcalendar = "nepali" // nepali
    BoxXcalendar_persian BoxXcalendar = "persian" // persian
    BoxXcalendar_jalali BoxXcalendar = "jalali" // jalali
    BoxXcalendar_taiwan BoxXcalendar = "taiwan" // taiwan
    BoxXcalendar_thai BoxXcalendar = "thai" // thai
    BoxXcalendar_ummalqura BoxXcalendar = "ummalqura" // ummalqura
)

// BoxYcalendar Sets the calendar system to use with `y` date data.
type BoxYcalendar string

const (
    BoxYcalendar_gregorian BoxYcalendar = "gregorian" // gregorian
    BoxYcalendar_chinese BoxYcalendar = "chinese" // chinese
    BoxYcalendar_coptic BoxYcalendar = "coptic" // coptic
    BoxYcalendar_discworld BoxYcalendar = "discworld" // discworld
    BoxYcalendar_ethiopian BoxYcalendar = "ethiopian" // ethiopian
    BoxYcalendar_hebrew BoxYcalendar = "hebrew" // hebrew
    BoxYcalendar_islamic BoxYcalendar = "islamic" // islamic
    BoxYcalendar_julian BoxYcalendar = "julian" // julian
    BoxYcalendar_mayan BoxYcalendar = "mayan" // mayan
    BoxYcalendar_nanakshahi BoxYcalendar = "nanakshahi" // nanakshahi
    BoxYcalendar_nepali BoxYcalendar = "nepali" // nepali
    BoxYcalendar_persian BoxYcalendar = "persian" // persian
    BoxYcalendar_jalali BoxYcalendar = "jalali" // jalali
    BoxYcalendar_taiwan BoxYcalendar = "taiwan" // taiwan
    BoxYcalendar_thai BoxYcalendar = "thai" // thai
    BoxYcalendar_ummalqura BoxYcalendar = "ummalqura" // ummalqura
)

// CandlestickHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type CandlestickHoverlabelAlign string

const (
    CandlestickHoverlabelAlign_left CandlestickHoverlabelAlign = "left" // left
    CandlestickHoverlabelAlign_right CandlestickHoverlabelAlign = "right" // right
    CandlestickHoverlabelAlign_auto CandlestickHoverlabelAlign = "auto" // auto
)

// CandlestickVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type CandlestickVisible interface{}

const (
    CandlestickVisibleTrue bool = true
    CandlestickVisibleFalse bool = false
    CandlestickVisibleLegendonly string = "legendonly"
)

// CandlestickXcalendar Sets the calendar system to use with `x` date data.
type CandlestickXcalendar string

const (
    CandlestickXcalendar_gregorian CandlestickXcalendar = "gregorian" // gregorian
    CandlestickXcalendar_chinese CandlestickXcalendar = "chinese" // chinese
    CandlestickXcalendar_coptic CandlestickXcalendar = "coptic" // coptic
    CandlestickXcalendar_discworld CandlestickXcalendar = "discworld" // discworld
    CandlestickXcalendar_ethiopian CandlestickXcalendar = "ethiopian" // ethiopian
    CandlestickXcalendar_hebrew CandlestickXcalendar = "hebrew" // hebrew
    CandlestickXcalendar_islamic CandlestickXcalendar = "islamic" // islamic
    CandlestickXcalendar_julian CandlestickXcalendar = "julian" // julian
    CandlestickXcalendar_mayan CandlestickXcalendar = "mayan" // mayan
    CandlestickXcalendar_nanakshahi CandlestickXcalendar = "nanakshahi" // nanakshahi
    CandlestickXcalendar_nepali CandlestickXcalendar = "nepali" // nepali
    CandlestickXcalendar_persian CandlestickXcalendar = "persian" // persian
    CandlestickXcalendar_jalali CandlestickXcalendar = "jalali" // jalali
    CandlestickXcalendar_taiwan CandlestickXcalendar = "taiwan" // taiwan
    CandlestickXcalendar_thai CandlestickXcalendar = "thai" // thai
    CandlestickXcalendar_ummalqura CandlestickXcalendar = "ummalqura" // ummalqura
)

// CarpetAaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type CarpetAaxisAutorange interface{}

const (
    CarpetAaxisAutorangeTrue bool = true
    CarpetAaxisAutorangeFalse bool = false
    CarpetAaxisAutorangeReversed string = "reversed"
)

// CarpetAaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`.
type CarpetAaxisCategoryorder string

const (
    CarpetAaxisCategoryorder_trace CarpetAaxisCategoryorder = "trace" // trace
    CarpetAaxisCategoryorder_categoryascending CarpetAaxisCategoryorder = "category ascending" // category ascending
    CarpetAaxisCategoryorder_categorydescending CarpetAaxisCategoryorder = "category descending" // category descending
    CarpetAaxisCategoryorder_array CarpetAaxisCategoryorder = "array" // array
)

// CarpetAaxisCheatertype <no value>
type CarpetAaxisCheatertype string

const (
    CarpetAaxisCheatertype_index CarpetAaxisCheatertype = "index" // index
    CarpetAaxisCheatertype_value CarpetAaxisCheatertype = "value" // value
)

// CarpetAaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type CarpetAaxisExponentformat string

const (
    CarpetAaxisExponentformat_none CarpetAaxisExponentformat = "none" // none
    CarpetAaxisExponentformat_e CarpetAaxisExponentformat = "e" // e
    CarpetAaxisExponentformat_E CarpetAaxisExponentformat = "E" // E
    CarpetAaxisExponentformat_power CarpetAaxisExponentformat = "power" // power
    CarpetAaxisExponentformat_SI CarpetAaxisExponentformat = "SI" // SI
    CarpetAaxisExponentformat_B CarpetAaxisExponentformat = "B" // B
)

// CarpetAaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data.
type CarpetAaxisRangemode string

const (
    CarpetAaxisRangemode_normal CarpetAaxisRangemode = "normal" // normal
    CarpetAaxisRangemode_tozero CarpetAaxisRangemode = "tozero" // tozero
    CarpetAaxisRangemode_nonnegative CarpetAaxisRangemode = "nonnegative" // nonnegative
)

// CarpetAaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type CarpetAaxisShowexponent string

const (
    CarpetAaxisShowexponent_all CarpetAaxisShowexponent = "all" // all
    CarpetAaxisShowexponent_first CarpetAaxisShowexponent = "first" // first
    CarpetAaxisShowexponent_last CarpetAaxisShowexponent = "last" // last
    CarpetAaxisShowexponent_none CarpetAaxisShowexponent = "none" // none
)

// CarpetAaxisShowticklabels Determines whether axis labels are drawn on the low side, the high side, both, or neither side of the axis.
type CarpetAaxisShowticklabels string

const (
    CarpetAaxisShowticklabels_start CarpetAaxisShowticklabels = "start" // start
    CarpetAaxisShowticklabels_end CarpetAaxisShowticklabels = "end" // end
    CarpetAaxisShowticklabels_both CarpetAaxisShowticklabels = "both" // both
    CarpetAaxisShowticklabels_none CarpetAaxisShowticklabels = "none" // none
)

// CarpetAaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type CarpetAaxisShowtickprefix string

const (
    CarpetAaxisShowtickprefix_all CarpetAaxisShowtickprefix = "all" // all
    CarpetAaxisShowtickprefix_first CarpetAaxisShowtickprefix = "first" // first
    CarpetAaxisShowtickprefix_last CarpetAaxisShowtickprefix = "last" // last
    CarpetAaxisShowtickprefix_none CarpetAaxisShowtickprefix = "none" // none
)

// CarpetAaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type CarpetAaxisShowticksuffix string

const (
    CarpetAaxisShowticksuffix_all CarpetAaxisShowticksuffix = "all" // all
    CarpetAaxisShowticksuffix_first CarpetAaxisShowticksuffix = "first" // first
    CarpetAaxisShowticksuffix_last CarpetAaxisShowticksuffix = "last" // last
    CarpetAaxisShowticksuffix_none CarpetAaxisShowticksuffix = "none" // none
)

// CarpetAaxisTickmode <no value>
type CarpetAaxisTickmode string

const (
    CarpetAaxisTickmode_linear CarpetAaxisTickmode = "linear" // linear
    CarpetAaxisTickmode_array CarpetAaxisTickmode = "array" // array
)

// CarpetAaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type CarpetAaxisType string

const (
    CarpetAaxisType__ CarpetAaxisType = "-" // -
    CarpetAaxisType_linear CarpetAaxisType = "linear" // linear
    CarpetAaxisType_date CarpetAaxisType = "date" // date
    CarpetAaxisType_category CarpetAaxisType = "category" // category
)

// CarpetBaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type CarpetBaxisAutorange interface{}

const (
    CarpetBaxisAutorangeTrue bool = true
    CarpetBaxisAutorangeFalse bool = false
    CarpetBaxisAutorangeReversed string = "reversed"
)

// CarpetBaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`.
type CarpetBaxisCategoryorder string

const (
    CarpetBaxisCategoryorder_trace CarpetBaxisCategoryorder = "trace" // trace
    CarpetBaxisCategoryorder_categoryascending CarpetBaxisCategoryorder = "category ascending" // category ascending
    CarpetBaxisCategoryorder_categorydescending CarpetBaxisCategoryorder = "category descending" // category descending
    CarpetBaxisCategoryorder_array CarpetBaxisCategoryorder = "array" // array
)

// CarpetBaxisCheatertype <no value>
type CarpetBaxisCheatertype string

const (
    CarpetBaxisCheatertype_index CarpetBaxisCheatertype = "index" // index
    CarpetBaxisCheatertype_value CarpetBaxisCheatertype = "value" // value
)

// CarpetBaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type CarpetBaxisExponentformat string

const (
    CarpetBaxisExponentformat_none CarpetBaxisExponentformat = "none" // none
    CarpetBaxisExponentformat_e CarpetBaxisExponentformat = "e" // e
    CarpetBaxisExponentformat_E CarpetBaxisExponentformat = "E" // E
    CarpetBaxisExponentformat_power CarpetBaxisExponentformat = "power" // power
    CarpetBaxisExponentformat_SI CarpetBaxisExponentformat = "SI" // SI
    CarpetBaxisExponentformat_B CarpetBaxisExponentformat = "B" // B
)

// CarpetBaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data.
type CarpetBaxisRangemode string

const (
    CarpetBaxisRangemode_normal CarpetBaxisRangemode = "normal" // normal
    CarpetBaxisRangemode_tozero CarpetBaxisRangemode = "tozero" // tozero
    CarpetBaxisRangemode_nonnegative CarpetBaxisRangemode = "nonnegative" // nonnegative
)

// CarpetBaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type CarpetBaxisShowexponent string

const (
    CarpetBaxisShowexponent_all CarpetBaxisShowexponent = "all" // all
    CarpetBaxisShowexponent_first CarpetBaxisShowexponent = "first" // first
    CarpetBaxisShowexponent_last CarpetBaxisShowexponent = "last" // last
    CarpetBaxisShowexponent_none CarpetBaxisShowexponent = "none" // none
)

// CarpetBaxisShowticklabels Determines whether axis labels are drawn on the low side, the high side, both, or neither side of the axis.
type CarpetBaxisShowticklabels string

const (
    CarpetBaxisShowticklabels_start CarpetBaxisShowticklabels = "start" // start
    CarpetBaxisShowticklabels_end CarpetBaxisShowticklabels = "end" // end
    CarpetBaxisShowticklabels_both CarpetBaxisShowticklabels = "both" // both
    CarpetBaxisShowticklabels_none CarpetBaxisShowticklabels = "none" // none
)

// CarpetBaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type CarpetBaxisShowtickprefix string

const (
    CarpetBaxisShowtickprefix_all CarpetBaxisShowtickprefix = "all" // all
    CarpetBaxisShowtickprefix_first CarpetBaxisShowtickprefix = "first" // first
    CarpetBaxisShowtickprefix_last CarpetBaxisShowtickprefix = "last" // last
    CarpetBaxisShowtickprefix_none CarpetBaxisShowtickprefix = "none" // none
)

// CarpetBaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type CarpetBaxisShowticksuffix string

const (
    CarpetBaxisShowticksuffix_all CarpetBaxisShowticksuffix = "all" // all
    CarpetBaxisShowticksuffix_first CarpetBaxisShowticksuffix = "first" // first
    CarpetBaxisShowticksuffix_last CarpetBaxisShowticksuffix = "last" // last
    CarpetBaxisShowticksuffix_none CarpetBaxisShowticksuffix = "none" // none
)

// CarpetBaxisTickmode <no value>
type CarpetBaxisTickmode string

const (
    CarpetBaxisTickmode_linear CarpetBaxisTickmode = "linear" // linear
    CarpetBaxisTickmode_array CarpetBaxisTickmode = "array" // array
)

// CarpetBaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type CarpetBaxisType string

const (
    CarpetBaxisType__ CarpetBaxisType = "-" // -
    CarpetBaxisType_linear CarpetBaxisType = "linear" // linear
    CarpetBaxisType_date CarpetBaxisType = "date" // date
    CarpetBaxisType_category CarpetBaxisType = "category" // category
)

// CarpetVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type CarpetVisible interface{}

const (
    CarpetVisibleTrue bool = true
    CarpetVisibleFalse bool = false
    CarpetVisibleLegendonly string = "legendonly"
)

// ChoroplethColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ChoroplethColorbarExponentformat string

const (
    ChoroplethColorbarExponentformat_none ChoroplethColorbarExponentformat = "none" // none
    ChoroplethColorbarExponentformat_e ChoroplethColorbarExponentformat = "e" // e
    ChoroplethColorbarExponentformat_E ChoroplethColorbarExponentformat = "E" // E
    ChoroplethColorbarExponentformat_power ChoroplethColorbarExponentformat = "power" // power
    ChoroplethColorbarExponentformat_SI ChoroplethColorbarExponentformat = "SI" // SI
    ChoroplethColorbarExponentformat_B ChoroplethColorbarExponentformat = "B" // B
)

// ChoroplethColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ChoroplethColorbarLenmode string

const (
    ChoroplethColorbarLenmode_fraction ChoroplethColorbarLenmode = "fraction" // fraction
    ChoroplethColorbarLenmode_pixels ChoroplethColorbarLenmode = "pixels" // pixels
)

// ChoroplethColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ChoroplethColorbarShowexponent string

const (
    ChoroplethColorbarShowexponent_all ChoroplethColorbarShowexponent = "all" // all
    ChoroplethColorbarShowexponent_first ChoroplethColorbarShowexponent = "first" // first
    ChoroplethColorbarShowexponent_last ChoroplethColorbarShowexponent = "last" // last
    ChoroplethColorbarShowexponent_none ChoroplethColorbarShowexponent = "none" // none
)

// ChoroplethColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ChoroplethColorbarShowtickprefix string

const (
    ChoroplethColorbarShowtickprefix_all ChoroplethColorbarShowtickprefix = "all" // all
    ChoroplethColorbarShowtickprefix_first ChoroplethColorbarShowtickprefix = "first" // first
    ChoroplethColorbarShowtickprefix_last ChoroplethColorbarShowtickprefix = "last" // last
    ChoroplethColorbarShowtickprefix_none ChoroplethColorbarShowtickprefix = "none" // none
)

// ChoroplethColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ChoroplethColorbarShowticksuffix string

const (
    ChoroplethColorbarShowticksuffix_all ChoroplethColorbarShowticksuffix = "all" // all
    ChoroplethColorbarShowticksuffix_first ChoroplethColorbarShowticksuffix = "first" // first
    ChoroplethColorbarShowticksuffix_last ChoroplethColorbarShowticksuffix = "last" // last
    ChoroplethColorbarShowticksuffix_none ChoroplethColorbarShowticksuffix = "none" // none
)

// ChoroplethColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ChoroplethColorbarThicknessmode string

const (
    ChoroplethColorbarThicknessmode_fraction ChoroplethColorbarThicknessmode = "fraction" // fraction
    ChoroplethColorbarThicknessmode_pixels ChoroplethColorbarThicknessmode = "pixels" // pixels
)

// ChoroplethColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ChoroplethColorbarTickmode string

const (
    ChoroplethColorbarTickmode_auto ChoroplethColorbarTickmode = "auto" // auto
    ChoroplethColorbarTickmode_linear ChoroplethColorbarTickmode = "linear" // linear
    ChoroplethColorbarTickmode_array ChoroplethColorbarTickmode = "array" // array
)

// ChoroplethColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ChoroplethColorbarTicks string

const (
    ChoroplethColorbarTicks_outside ChoroplethColorbarTicks = "outside" // outside
    ChoroplethColorbarTicks_inside ChoroplethColorbarTicks = "inside" // inside
)

// ChoroplethColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ChoroplethColorbarTitleSide string

const (
    ChoroplethColorbarTitleSide_right ChoroplethColorbarTitleSide = "right" // right
    ChoroplethColorbarTitleSide_top ChoroplethColorbarTitleSide = "top" // top
    ChoroplethColorbarTitleSide_bottom ChoroplethColorbarTitleSide = "bottom" // bottom
)

// ChoroplethColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ChoroplethColorbarXanchor string

const (
    ChoroplethColorbarXanchor_left ChoroplethColorbarXanchor = "left" // left
    ChoroplethColorbarXanchor_center ChoroplethColorbarXanchor = "center" // center
    ChoroplethColorbarXanchor_right ChoroplethColorbarXanchor = "right" // right
)

// ChoroplethColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ChoroplethColorbarYanchor string

const (
    ChoroplethColorbarYanchor_top ChoroplethColorbarYanchor = "top" // top
    ChoroplethColorbarYanchor_middle ChoroplethColorbarYanchor = "middle" // middle
    ChoroplethColorbarYanchor_bottom ChoroplethColorbarYanchor = "bottom" // bottom
)

// ChoroplethHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ChoroplethHoverlabelAlign string

const (
    ChoroplethHoverlabelAlign_left ChoroplethHoverlabelAlign = "left" // left
    ChoroplethHoverlabelAlign_right ChoroplethHoverlabelAlign = "right" // right
    ChoroplethHoverlabelAlign_auto ChoroplethHoverlabelAlign = "auto" // auto
)

// ChoroplethLocationmode Determines the set of locations used to match entries in `locations` to regions on the map. Values *ISO-3*, *USA-states*, *country names* correspond to features on the base map and value *geojson-id* corresponds to features from a custom GeoJSON linked to the `geojson` attribute.
type ChoroplethLocationmode string

const (
    ChoroplethLocationmode_ISO_3 ChoroplethLocationmode = "ISO-3" // ISO-3
    ChoroplethLocationmode_USA_states ChoroplethLocationmode = "USA-states" // USA-states
    ChoroplethLocationmode_countrynames ChoroplethLocationmode = "country names" // country names
    ChoroplethLocationmode_geojson_id ChoroplethLocationmode = "geojson-id" // geojson-id
)

// ChoroplethVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ChoroplethVisible interface{}

const (
    ChoroplethVisibleTrue bool = true
    ChoroplethVisibleFalse bool = false
    ChoroplethVisibleLegendonly string = "legendonly"
)

// ChoroplethmapboxColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ChoroplethmapboxColorbarExponentformat string

const (
    ChoroplethmapboxColorbarExponentformat_none ChoroplethmapboxColorbarExponentformat = "none" // none
    ChoroplethmapboxColorbarExponentformat_e ChoroplethmapboxColorbarExponentformat = "e" // e
    ChoroplethmapboxColorbarExponentformat_E ChoroplethmapboxColorbarExponentformat = "E" // E
    ChoroplethmapboxColorbarExponentformat_power ChoroplethmapboxColorbarExponentformat = "power" // power
    ChoroplethmapboxColorbarExponentformat_SI ChoroplethmapboxColorbarExponentformat = "SI" // SI
    ChoroplethmapboxColorbarExponentformat_B ChoroplethmapboxColorbarExponentformat = "B" // B
)

// ChoroplethmapboxColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ChoroplethmapboxColorbarLenmode string

const (
    ChoroplethmapboxColorbarLenmode_fraction ChoroplethmapboxColorbarLenmode = "fraction" // fraction
    ChoroplethmapboxColorbarLenmode_pixels ChoroplethmapboxColorbarLenmode = "pixels" // pixels
)

// ChoroplethmapboxColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ChoroplethmapboxColorbarShowexponent string

const (
    ChoroplethmapboxColorbarShowexponent_all ChoroplethmapboxColorbarShowexponent = "all" // all
    ChoroplethmapboxColorbarShowexponent_first ChoroplethmapboxColorbarShowexponent = "first" // first
    ChoroplethmapboxColorbarShowexponent_last ChoroplethmapboxColorbarShowexponent = "last" // last
    ChoroplethmapboxColorbarShowexponent_none ChoroplethmapboxColorbarShowexponent = "none" // none
)

// ChoroplethmapboxColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ChoroplethmapboxColorbarShowtickprefix string

const (
    ChoroplethmapboxColorbarShowtickprefix_all ChoroplethmapboxColorbarShowtickprefix = "all" // all
    ChoroplethmapboxColorbarShowtickprefix_first ChoroplethmapboxColorbarShowtickprefix = "first" // first
    ChoroplethmapboxColorbarShowtickprefix_last ChoroplethmapboxColorbarShowtickprefix = "last" // last
    ChoroplethmapboxColorbarShowtickprefix_none ChoroplethmapboxColorbarShowtickprefix = "none" // none
)

// ChoroplethmapboxColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ChoroplethmapboxColorbarShowticksuffix string

const (
    ChoroplethmapboxColorbarShowticksuffix_all ChoroplethmapboxColorbarShowticksuffix = "all" // all
    ChoroplethmapboxColorbarShowticksuffix_first ChoroplethmapboxColorbarShowticksuffix = "first" // first
    ChoroplethmapboxColorbarShowticksuffix_last ChoroplethmapboxColorbarShowticksuffix = "last" // last
    ChoroplethmapboxColorbarShowticksuffix_none ChoroplethmapboxColorbarShowticksuffix = "none" // none
)

// ChoroplethmapboxColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ChoroplethmapboxColorbarThicknessmode string

const (
    ChoroplethmapboxColorbarThicknessmode_fraction ChoroplethmapboxColorbarThicknessmode = "fraction" // fraction
    ChoroplethmapboxColorbarThicknessmode_pixels ChoroplethmapboxColorbarThicknessmode = "pixels" // pixels
)

// ChoroplethmapboxColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ChoroplethmapboxColorbarTickmode string

const (
    ChoroplethmapboxColorbarTickmode_auto ChoroplethmapboxColorbarTickmode = "auto" // auto
    ChoroplethmapboxColorbarTickmode_linear ChoroplethmapboxColorbarTickmode = "linear" // linear
    ChoroplethmapboxColorbarTickmode_array ChoroplethmapboxColorbarTickmode = "array" // array
)

// ChoroplethmapboxColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ChoroplethmapboxColorbarTicks string

const (
    ChoroplethmapboxColorbarTicks_outside ChoroplethmapboxColorbarTicks = "outside" // outside
    ChoroplethmapboxColorbarTicks_inside ChoroplethmapboxColorbarTicks = "inside" // inside
)

// ChoroplethmapboxColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ChoroplethmapboxColorbarTitleSide string

const (
    ChoroplethmapboxColorbarTitleSide_right ChoroplethmapboxColorbarTitleSide = "right" // right
    ChoroplethmapboxColorbarTitleSide_top ChoroplethmapboxColorbarTitleSide = "top" // top
    ChoroplethmapboxColorbarTitleSide_bottom ChoroplethmapboxColorbarTitleSide = "bottom" // bottom
)

// ChoroplethmapboxColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ChoroplethmapboxColorbarXanchor string

const (
    ChoroplethmapboxColorbarXanchor_left ChoroplethmapboxColorbarXanchor = "left" // left
    ChoroplethmapboxColorbarXanchor_center ChoroplethmapboxColorbarXanchor = "center" // center
    ChoroplethmapboxColorbarXanchor_right ChoroplethmapboxColorbarXanchor = "right" // right
)

// ChoroplethmapboxColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ChoroplethmapboxColorbarYanchor string

const (
    ChoroplethmapboxColorbarYanchor_top ChoroplethmapboxColorbarYanchor = "top" // top
    ChoroplethmapboxColorbarYanchor_middle ChoroplethmapboxColorbarYanchor = "middle" // middle
    ChoroplethmapboxColorbarYanchor_bottom ChoroplethmapboxColorbarYanchor = "bottom" // bottom
)

// ChoroplethmapboxHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ChoroplethmapboxHoverlabelAlign string

const (
    ChoroplethmapboxHoverlabelAlign_left ChoroplethmapboxHoverlabelAlign = "left" // left
    ChoroplethmapboxHoverlabelAlign_right ChoroplethmapboxHoverlabelAlign = "right" // right
    ChoroplethmapboxHoverlabelAlign_auto ChoroplethmapboxHoverlabelAlign = "auto" // auto
)

// ChoroplethmapboxVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ChoroplethmapboxVisible interface{}

const (
    ChoroplethmapboxVisibleTrue bool = true
    ChoroplethmapboxVisibleFalse bool = false
    ChoroplethmapboxVisibleLegendonly string = "legendonly"
)

// ConeAnchor Sets the cones' anchor with respect to their x/y/z positions. Note that *cm* denote the cone's center of mass which corresponds to 1/4 from the tail to tip.
type ConeAnchor string

const (
    ConeAnchor_tip ConeAnchor = "tip" // tip
    ConeAnchor_tail ConeAnchor = "tail" // tail
    ConeAnchor_cm ConeAnchor = "cm" // cm
    ConeAnchor_center ConeAnchor = "center" // center
)

// ConeColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ConeColorbarExponentformat string

const (
    ConeColorbarExponentformat_none ConeColorbarExponentformat = "none" // none
    ConeColorbarExponentformat_e ConeColorbarExponentformat = "e" // e
    ConeColorbarExponentformat_E ConeColorbarExponentformat = "E" // E
    ConeColorbarExponentformat_power ConeColorbarExponentformat = "power" // power
    ConeColorbarExponentformat_SI ConeColorbarExponentformat = "SI" // SI
    ConeColorbarExponentformat_B ConeColorbarExponentformat = "B" // B
)

// ConeColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ConeColorbarLenmode string

const (
    ConeColorbarLenmode_fraction ConeColorbarLenmode = "fraction" // fraction
    ConeColorbarLenmode_pixels ConeColorbarLenmode = "pixels" // pixels
)

// ConeColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ConeColorbarShowexponent string

const (
    ConeColorbarShowexponent_all ConeColorbarShowexponent = "all" // all
    ConeColorbarShowexponent_first ConeColorbarShowexponent = "first" // first
    ConeColorbarShowexponent_last ConeColorbarShowexponent = "last" // last
    ConeColorbarShowexponent_none ConeColorbarShowexponent = "none" // none
)

// ConeColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ConeColorbarShowtickprefix string

const (
    ConeColorbarShowtickprefix_all ConeColorbarShowtickprefix = "all" // all
    ConeColorbarShowtickprefix_first ConeColorbarShowtickprefix = "first" // first
    ConeColorbarShowtickprefix_last ConeColorbarShowtickprefix = "last" // last
    ConeColorbarShowtickprefix_none ConeColorbarShowtickprefix = "none" // none
)

// ConeColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ConeColorbarShowticksuffix string

const (
    ConeColorbarShowticksuffix_all ConeColorbarShowticksuffix = "all" // all
    ConeColorbarShowticksuffix_first ConeColorbarShowticksuffix = "first" // first
    ConeColorbarShowticksuffix_last ConeColorbarShowticksuffix = "last" // last
    ConeColorbarShowticksuffix_none ConeColorbarShowticksuffix = "none" // none
)

// ConeColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ConeColorbarThicknessmode string

const (
    ConeColorbarThicknessmode_fraction ConeColorbarThicknessmode = "fraction" // fraction
    ConeColorbarThicknessmode_pixels ConeColorbarThicknessmode = "pixels" // pixels
)

// ConeColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ConeColorbarTickmode string

const (
    ConeColorbarTickmode_auto ConeColorbarTickmode = "auto" // auto
    ConeColorbarTickmode_linear ConeColorbarTickmode = "linear" // linear
    ConeColorbarTickmode_array ConeColorbarTickmode = "array" // array
)

// ConeColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ConeColorbarTicks string

const (
    ConeColorbarTicks_outside ConeColorbarTicks = "outside" // outside
    ConeColorbarTicks_inside ConeColorbarTicks = "inside" // inside
)

// ConeColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ConeColorbarTitleSide string

const (
    ConeColorbarTitleSide_right ConeColorbarTitleSide = "right" // right
    ConeColorbarTitleSide_top ConeColorbarTitleSide = "top" // top
    ConeColorbarTitleSide_bottom ConeColorbarTitleSide = "bottom" // bottom
)

// ConeColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ConeColorbarXanchor string

const (
    ConeColorbarXanchor_left ConeColorbarXanchor = "left" // left
    ConeColorbarXanchor_center ConeColorbarXanchor = "center" // center
    ConeColorbarXanchor_right ConeColorbarXanchor = "right" // right
)

// ConeColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ConeColorbarYanchor string

const (
    ConeColorbarYanchor_top ConeColorbarYanchor = "top" // top
    ConeColorbarYanchor_middle ConeColorbarYanchor = "middle" // middle
    ConeColorbarYanchor_bottom ConeColorbarYanchor = "bottom" // bottom
)

// ConeHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ConeHoverlabelAlign string

const (
    ConeHoverlabelAlign_left ConeHoverlabelAlign = "left" // left
    ConeHoverlabelAlign_right ConeHoverlabelAlign = "right" // right
    ConeHoverlabelAlign_auto ConeHoverlabelAlign = "auto" // auto
)

// ConeSizemode Determines whether `sizeref` is set as a *scaled* (i.e unitless) scalar (normalized by the max u/v/w norm in the vector field) or as *absolute* value (in the same units as the vector field).
type ConeSizemode string

const (
    ConeSizemode_scaled ConeSizemode = "scaled" // scaled
    ConeSizemode_absolute ConeSizemode = "absolute" // absolute
)

// ConeVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ConeVisible interface{}

const (
    ConeVisibleTrue bool = true
    ConeVisibleFalse bool = false
    ConeVisibleLegendonly string = "legendonly"
)

// ContourColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ContourColorbarExponentformat string

const (
    ContourColorbarExponentformat_none ContourColorbarExponentformat = "none" // none
    ContourColorbarExponentformat_e ContourColorbarExponentformat = "e" // e
    ContourColorbarExponentformat_E ContourColorbarExponentformat = "E" // E
    ContourColorbarExponentformat_power ContourColorbarExponentformat = "power" // power
    ContourColorbarExponentformat_SI ContourColorbarExponentformat = "SI" // SI
    ContourColorbarExponentformat_B ContourColorbarExponentformat = "B" // B
)

// ContourColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ContourColorbarLenmode string

const (
    ContourColorbarLenmode_fraction ContourColorbarLenmode = "fraction" // fraction
    ContourColorbarLenmode_pixels ContourColorbarLenmode = "pixels" // pixels
)

// ContourColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ContourColorbarShowexponent string

const (
    ContourColorbarShowexponent_all ContourColorbarShowexponent = "all" // all
    ContourColorbarShowexponent_first ContourColorbarShowexponent = "first" // first
    ContourColorbarShowexponent_last ContourColorbarShowexponent = "last" // last
    ContourColorbarShowexponent_none ContourColorbarShowexponent = "none" // none
)

// ContourColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ContourColorbarShowtickprefix string

const (
    ContourColorbarShowtickprefix_all ContourColorbarShowtickprefix = "all" // all
    ContourColorbarShowtickprefix_first ContourColorbarShowtickprefix = "first" // first
    ContourColorbarShowtickprefix_last ContourColorbarShowtickprefix = "last" // last
    ContourColorbarShowtickprefix_none ContourColorbarShowtickprefix = "none" // none
)

// ContourColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ContourColorbarShowticksuffix string

const (
    ContourColorbarShowticksuffix_all ContourColorbarShowticksuffix = "all" // all
    ContourColorbarShowticksuffix_first ContourColorbarShowticksuffix = "first" // first
    ContourColorbarShowticksuffix_last ContourColorbarShowticksuffix = "last" // last
    ContourColorbarShowticksuffix_none ContourColorbarShowticksuffix = "none" // none
)

// ContourColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ContourColorbarThicknessmode string

const (
    ContourColorbarThicknessmode_fraction ContourColorbarThicknessmode = "fraction" // fraction
    ContourColorbarThicknessmode_pixels ContourColorbarThicknessmode = "pixels" // pixels
)

// ContourColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ContourColorbarTickmode string

const (
    ContourColorbarTickmode_auto ContourColorbarTickmode = "auto" // auto
    ContourColorbarTickmode_linear ContourColorbarTickmode = "linear" // linear
    ContourColorbarTickmode_array ContourColorbarTickmode = "array" // array
)

// ContourColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ContourColorbarTicks string

const (
    ContourColorbarTicks_outside ContourColorbarTicks = "outside" // outside
    ContourColorbarTicks_inside ContourColorbarTicks = "inside" // inside
)

// ContourColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ContourColorbarTitleSide string

const (
    ContourColorbarTitleSide_right ContourColorbarTitleSide = "right" // right
    ContourColorbarTitleSide_top ContourColorbarTitleSide = "top" // top
    ContourColorbarTitleSide_bottom ContourColorbarTitleSide = "bottom" // bottom
)

// ContourColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ContourColorbarXanchor string

const (
    ContourColorbarXanchor_left ContourColorbarXanchor = "left" // left
    ContourColorbarXanchor_center ContourColorbarXanchor = "center" // center
    ContourColorbarXanchor_right ContourColorbarXanchor = "right" // right
)

// ContourColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ContourColorbarYanchor string

const (
    ContourColorbarYanchor_top ContourColorbarYanchor = "top" // top
    ContourColorbarYanchor_middle ContourColorbarYanchor = "middle" // middle
    ContourColorbarYanchor_bottom ContourColorbarYanchor = "bottom" // bottom
)

// ContourContoursColoring Determines the coloring method showing the contour values. If *fill*, coloring is done evenly between each contour level If *heatmap*, a heatmap gradient coloring is applied between each contour level. If *lines*, coloring is done on the contour lines. If *none*, no coloring is applied on this trace.
type ContourContoursColoring string

const (
    ContourContoursColoring_fill ContourContoursColoring = "fill" // fill
    ContourContoursColoring_heatmap ContourContoursColoring = "heatmap" // heatmap
    ContourContoursColoring_lines ContourContoursColoring = "lines" // lines
    ContourContoursColoring_none ContourContoursColoring = "none" // none
)

// ContourContoursOperation Sets the constraint operation. *=* keeps regions equal to `value` *<* and *<=* keep regions less than `value` *>* and *>=* keep regions greater than `value` *[]*, *()*, *[)*, and *(]* keep regions inside `value[0]` to `value[1]` *][*, *)(*, *](*, *)[* keep regions outside `value[0]` to value[1]` Open vs. closed intervals make no difference to constraint display, but all versions are allowed for consistency with filter transforms.
type ContourContoursOperation string

const (
    ContourContoursOperation_eq ContourContoursOperation = "=" // =
    ContourContoursOperation_lt ContourContoursOperation = "<" // <
    ContourContoursOperation_gteq ContourContoursOperation = ">=" // >=
    ContourContoursOperation_gt ContourContoursOperation = ">" // >
    ContourContoursOperation_lteq ContourContoursOperation = "<=" // <=
    ContourContoursOperation__lbracket__rbracket_ ContourContoursOperation = "[]" // []
    ContourContoursOperation__lpar__rpar_ ContourContoursOperation = "()" // ()
    ContourContoursOperation__lbracket__rpar_ ContourContoursOperation = "[)" // [)
    ContourContoursOperation__lpar__rbracket_ ContourContoursOperation = "(]" // (]
    ContourContoursOperation__rbracket__lbracket_ ContourContoursOperation = "][" // ][
    ContourContoursOperation__rpar__lpar_ ContourContoursOperation = ")(" // )(
    ContourContoursOperation__rbracket__lpar_ ContourContoursOperation = "](" // ](
    ContourContoursOperation__rpar__lbracket_ ContourContoursOperation = ")[" // )[
)

// ContourContoursType If `levels`, the data is represented as a contour plot with multiple levels displayed. If `constraint`, the data is represented as constraints with the invalid region shaded as specified by the `operation` and `value` parameters.
type ContourContoursType string

const (
    ContourContoursType_levels ContourContoursType = "levels" // levels
    ContourContoursType_constraint ContourContoursType = "constraint" // constraint
)

// ContourHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ContourHoverlabelAlign string

const (
    ContourHoverlabelAlign_left ContourHoverlabelAlign = "left" // left
    ContourHoverlabelAlign_right ContourHoverlabelAlign = "right" // right
    ContourHoverlabelAlign_auto ContourHoverlabelAlign = "auto" // auto
)

// ContourVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ContourVisible interface{}

const (
    ContourVisibleTrue bool = true
    ContourVisibleFalse bool = false
    ContourVisibleLegendonly string = "legendonly"
)

// ContourXcalendar Sets the calendar system to use with `x` date data.
type ContourXcalendar string

const (
    ContourXcalendar_gregorian ContourXcalendar = "gregorian" // gregorian
    ContourXcalendar_chinese ContourXcalendar = "chinese" // chinese
    ContourXcalendar_coptic ContourXcalendar = "coptic" // coptic
    ContourXcalendar_discworld ContourXcalendar = "discworld" // discworld
    ContourXcalendar_ethiopian ContourXcalendar = "ethiopian" // ethiopian
    ContourXcalendar_hebrew ContourXcalendar = "hebrew" // hebrew
    ContourXcalendar_islamic ContourXcalendar = "islamic" // islamic
    ContourXcalendar_julian ContourXcalendar = "julian" // julian
    ContourXcalendar_mayan ContourXcalendar = "mayan" // mayan
    ContourXcalendar_nanakshahi ContourXcalendar = "nanakshahi" // nanakshahi
    ContourXcalendar_nepali ContourXcalendar = "nepali" // nepali
    ContourXcalendar_persian ContourXcalendar = "persian" // persian
    ContourXcalendar_jalali ContourXcalendar = "jalali" // jalali
    ContourXcalendar_taiwan ContourXcalendar = "taiwan" // taiwan
    ContourXcalendar_thai ContourXcalendar = "thai" // thai
    ContourXcalendar_ummalqura ContourXcalendar = "ummalqura" // ummalqura
)

// ContourXtype If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
type ContourXtype string

const (
    ContourXtype_array ContourXtype = "array" // array
    ContourXtype_scaled ContourXtype = "scaled" // scaled
)

// ContourYcalendar Sets the calendar system to use with `y` date data.
type ContourYcalendar string

const (
    ContourYcalendar_gregorian ContourYcalendar = "gregorian" // gregorian
    ContourYcalendar_chinese ContourYcalendar = "chinese" // chinese
    ContourYcalendar_coptic ContourYcalendar = "coptic" // coptic
    ContourYcalendar_discworld ContourYcalendar = "discworld" // discworld
    ContourYcalendar_ethiopian ContourYcalendar = "ethiopian" // ethiopian
    ContourYcalendar_hebrew ContourYcalendar = "hebrew" // hebrew
    ContourYcalendar_islamic ContourYcalendar = "islamic" // islamic
    ContourYcalendar_julian ContourYcalendar = "julian" // julian
    ContourYcalendar_mayan ContourYcalendar = "mayan" // mayan
    ContourYcalendar_nanakshahi ContourYcalendar = "nanakshahi" // nanakshahi
    ContourYcalendar_nepali ContourYcalendar = "nepali" // nepali
    ContourYcalendar_persian ContourYcalendar = "persian" // persian
    ContourYcalendar_jalali ContourYcalendar = "jalali" // jalali
    ContourYcalendar_taiwan ContourYcalendar = "taiwan" // taiwan
    ContourYcalendar_thai ContourYcalendar = "thai" // thai
    ContourYcalendar_ummalqura ContourYcalendar = "ummalqura" // ummalqura
)

// ContourYtype If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
type ContourYtype string

const (
    ContourYtype_array ContourYtype = "array" // array
    ContourYtype_scaled ContourYtype = "scaled" // scaled
)

// ContourcarpetAtype If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
type ContourcarpetAtype string

const (
    ContourcarpetAtype_array ContourcarpetAtype = "array" // array
    ContourcarpetAtype_scaled ContourcarpetAtype = "scaled" // scaled
)

// ContourcarpetBtype If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
type ContourcarpetBtype string

const (
    ContourcarpetBtype_array ContourcarpetBtype = "array" // array
    ContourcarpetBtype_scaled ContourcarpetBtype = "scaled" // scaled
)

// ContourcarpetColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ContourcarpetColorbarExponentformat string

const (
    ContourcarpetColorbarExponentformat_none ContourcarpetColorbarExponentformat = "none" // none
    ContourcarpetColorbarExponentformat_e ContourcarpetColorbarExponentformat = "e" // e
    ContourcarpetColorbarExponentformat_E ContourcarpetColorbarExponentformat = "E" // E
    ContourcarpetColorbarExponentformat_power ContourcarpetColorbarExponentformat = "power" // power
    ContourcarpetColorbarExponentformat_SI ContourcarpetColorbarExponentformat = "SI" // SI
    ContourcarpetColorbarExponentformat_B ContourcarpetColorbarExponentformat = "B" // B
)

// ContourcarpetColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ContourcarpetColorbarLenmode string

const (
    ContourcarpetColorbarLenmode_fraction ContourcarpetColorbarLenmode = "fraction" // fraction
    ContourcarpetColorbarLenmode_pixels ContourcarpetColorbarLenmode = "pixels" // pixels
)

// ContourcarpetColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ContourcarpetColorbarShowexponent string

const (
    ContourcarpetColorbarShowexponent_all ContourcarpetColorbarShowexponent = "all" // all
    ContourcarpetColorbarShowexponent_first ContourcarpetColorbarShowexponent = "first" // first
    ContourcarpetColorbarShowexponent_last ContourcarpetColorbarShowexponent = "last" // last
    ContourcarpetColorbarShowexponent_none ContourcarpetColorbarShowexponent = "none" // none
)

// ContourcarpetColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ContourcarpetColorbarShowtickprefix string

const (
    ContourcarpetColorbarShowtickprefix_all ContourcarpetColorbarShowtickprefix = "all" // all
    ContourcarpetColorbarShowtickprefix_first ContourcarpetColorbarShowtickprefix = "first" // first
    ContourcarpetColorbarShowtickprefix_last ContourcarpetColorbarShowtickprefix = "last" // last
    ContourcarpetColorbarShowtickprefix_none ContourcarpetColorbarShowtickprefix = "none" // none
)

// ContourcarpetColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ContourcarpetColorbarShowticksuffix string

const (
    ContourcarpetColorbarShowticksuffix_all ContourcarpetColorbarShowticksuffix = "all" // all
    ContourcarpetColorbarShowticksuffix_first ContourcarpetColorbarShowticksuffix = "first" // first
    ContourcarpetColorbarShowticksuffix_last ContourcarpetColorbarShowticksuffix = "last" // last
    ContourcarpetColorbarShowticksuffix_none ContourcarpetColorbarShowticksuffix = "none" // none
)

// ContourcarpetColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ContourcarpetColorbarThicknessmode string

const (
    ContourcarpetColorbarThicknessmode_fraction ContourcarpetColorbarThicknessmode = "fraction" // fraction
    ContourcarpetColorbarThicknessmode_pixels ContourcarpetColorbarThicknessmode = "pixels" // pixels
)

// ContourcarpetColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ContourcarpetColorbarTickmode string

const (
    ContourcarpetColorbarTickmode_auto ContourcarpetColorbarTickmode = "auto" // auto
    ContourcarpetColorbarTickmode_linear ContourcarpetColorbarTickmode = "linear" // linear
    ContourcarpetColorbarTickmode_array ContourcarpetColorbarTickmode = "array" // array
)

// ContourcarpetColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ContourcarpetColorbarTicks string

const (
    ContourcarpetColorbarTicks_outside ContourcarpetColorbarTicks = "outside" // outside
    ContourcarpetColorbarTicks_inside ContourcarpetColorbarTicks = "inside" // inside
)

// ContourcarpetColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ContourcarpetColorbarTitleSide string

const (
    ContourcarpetColorbarTitleSide_right ContourcarpetColorbarTitleSide = "right" // right
    ContourcarpetColorbarTitleSide_top ContourcarpetColorbarTitleSide = "top" // top
    ContourcarpetColorbarTitleSide_bottom ContourcarpetColorbarTitleSide = "bottom" // bottom
)

// ContourcarpetColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ContourcarpetColorbarXanchor string

const (
    ContourcarpetColorbarXanchor_left ContourcarpetColorbarXanchor = "left" // left
    ContourcarpetColorbarXanchor_center ContourcarpetColorbarXanchor = "center" // center
    ContourcarpetColorbarXanchor_right ContourcarpetColorbarXanchor = "right" // right
)

// ContourcarpetColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ContourcarpetColorbarYanchor string

const (
    ContourcarpetColorbarYanchor_top ContourcarpetColorbarYanchor = "top" // top
    ContourcarpetColorbarYanchor_middle ContourcarpetColorbarYanchor = "middle" // middle
    ContourcarpetColorbarYanchor_bottom ContourcarpetColorbarYanchor = "bottom" // bottom
)

// ContourcarpetContoursColoring Determines the coloring method showing the contour values. If *fill*, coloring is done evenly between each contour level If *lines*, coloring is done on the contour lines. If *none*, no coloring is applied on this trace.
type ContourcarpetContoursColoring string

const (
    ContourcarpetContoursColoring_fill ContourcarpetContoursColoring = "fill" // fill
    ContourcarpetContoursColoring_lines ContourcarpetContoursColoring = "lines" // lines
    ContourcarpetContoursColoring_none ContourcarpetContoursColoring = "none" // none
)

// ContourcarpetContoursOperation Sets the constraint operation. *=* keeps regions equal to `value` *<* and *<=* keep regions less than `value` *>* and *>=* keep regions greater than `value` *[]*, *()*, *[)*, and *(]* keep regions inside `value[0]` to `value[1]` *][*, *)(*, *](*, *)[* keep regions outside `value[0]` to value[1]` Open vs. closed intervals make no difference to constraint display, but all versions are allowed for consistency with filter transforms.
type ContourcarpetContoursOperation string

const (
    ContourcarpetContoursOperation_eq ContourcarpetContoursOperation = "=" // =
    ContourcarpetContoursOperation_lt ContourcarpetContoursOperation = "<" // <
    ContourcarpetContoursOperation_gteq ContourcarpetContoursOperation = ">=" // >=
    ContourcarpetContoursOperation_gt ContourcarpetContoursOperation = ">" // >
    ContourcarpetContoursOperation_lteq ContourcarpetContoursOperation = "<=" // <=
    ContourcarpetContoursOperation__lbracket__rbracket_ ContourcarpetContoursOperation = "[]" // []
    ContourcarpetContoursOperation__lpar__rpar_ ContourcarpetContoursOperation = "()" // ()
    ContourcarpetContoursOperation__lbracket__rpar_ ContourcarpetContoursOperation = "[)" // [)
    ContourcarpetContoursOperation__lpar__rbracket_ ContourcarpetContoursOperation = "(]" // (]
    ContourcarpetContoursOperation__rbracket__lbracket_ ContourcarpetContoursOperation = "][" // ][
    ContourcarpetContoursOperation__rpar__lpar_ ContourcarpetContoursOperation = ")(" // )(
    ContourcarpetContoursOperation__rbracket__lpar_ ContourcarpetContoursOperation = "](" // ](
    ContourcarpetContoursOperation__rpar__lbracket_ ContourcarpetContoursOperation = ")[" // )[
)

// ContourcarpetContoursType If `levels`, the data is represented as a contour plot with multiple levels displayed. If `constraint`, the data is represented as constraints with the invalid region shaded as specified by the `operation` and `value` parameters.
type ContourcarpetContoursType string

const (
    ContourcarpetContoursType_levels ContourcarpetContoursType = "levels" // levels
    ContourcarpetContoursType_constraint ContourcarpetContoursType = "constraint" // constraint
)

// ContourcarpetVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ContourcarpetVisible interface{}

const (
    ContourcarpetVisibleTrue bool = true
    ContourcarpetVisibleFalse bool = false
    ContourcarpetVisibleLegendonly string = "legendonly"
)

// DensitymapboxColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type DensitymapboxColorbarExponentformat string

const (
    DensitymapboxColorbarExponentformat_none DensitymapboxColorbarExponentformat = "none" // none
    DensitymapboxColorbarExponentformat_e DensitymapboxColorbarExponentformat = "e" // e
    DensitymapboxColorbarExponentformat_E DensitymapboxColorbarExponentformat = "E" // E
    DensitymapboxColorbarExponentformat_power DensitymapboxColorbarExponentformat = "power" // power
    DensitymapboxColorbarExponentformat_SI DensitymapboxColorbarExponentformat = "SI" // SI
    DensitymapboxColorbarExponentformat_B DensitymapboxColorbarExponentformat = "B" // B
)

// DensitymapboxColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type DensitymapboxColorbarLenmode string

const (
    DensitymapboxColorbarLenmode_fraction DensitymapboxColorbarLenmode = "fraction" // fraction
    DensitymapboxColorbarLenmode_pixels DensitymapboxColorbarLenmode = "pixels" // pixels
)

// DensitymapboxColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type DensitymapboxColorbarShowexponent string

const (
    DensitymapboxColorbarShowexponent_all DensitymapboxColorbarShowexponent = "all" // all
    DensitymapboxColorbarShowexponent_first DensitymapboxColorbarShowexponent = "first" // first
    DensitymapboxColorbarShowexponent_last DensitymapboxColorbarShowexponent = "last" // last
    DensitymapboxColorbarShowexponent_none DensitymapboxColorbarShowexponent = "none" // none
)

// DensitymapboxColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type DensitymapboxColorbarShowtickprefix string

const (
    DensitymapboxColorbarShowtickprefix_all DensitymapboxColorbarShowtickprefix = "all" // all
    DensitymapboxColorbarShowtickprefix_first DensitymapboxColorbarShowtickprefix = "first" // first
    DensitymapboxColorbarShowtickprefix_last DensitymapboxColorbarShowtickprefix = "last" // last
    DensitymapboxColorbarShowtickprefix_none DensitymapboxColorbarShowtickprefix = "none" // none
)

// DensitymapboxColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type DensitymapboxColorbarShowticksuffix string

const (
    DensitymapboxColorbarShowticksuffix_all DensitymapboxColorbarShowticksuffix = "all" // all
    DensitymapboxColorbarShowticksuffix_first DensitymapboxColorbarShowticksuffix = "first" // first
    DensitymapboxColorbarShowticksuffix_last DensitymapboxColorbarShowticksuffix = "last" // last
    DensitymapboxColorbarShowticksuffix_none DensitymapboxColorbarShowticksuffix = "none" // none
)

// DensitymapboxColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type DensitymapboxColorbarThicknessmode string

const (
    DensitymapboxColorbarThicknessmode_fraction DensitymapboxColorbarThicknessmode = "fraction" // fraction
    DensitymapboxColorbarThicknessmode_pixels DensitymapboxColorbarThicknessmode = "pixels" // pixels
)

// DensitymapboxColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type DensitymapboxColorbarTickmode string

const (
    DensitymapboxColorbarTickmode_auto DensitymapboxColorbarTickmode = "auto" // auto
    DensitymapboxColorbarTickmode_linear DensitymapboxColorbarTickmode = "linear" // linear
    DensitymapboxColorbarTickmode_array DensitymapboxColorbarTickmode = "array" // array
)

// DensitymapboxColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type DensitymapboxColorbarTicks string

const (
    DensitymapboxColorbarTicks_outside DensitymapboxColorbarTicks = "outside" // outside
    DensitymapboxColorbarTicks_inside DensitymapboxColorbarTicks = "inside" // inside
)

// DensitymapboxColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type DensitymapboxColorbarTitleSide string

const (
    DensitymapboxColorbarTitleSide_right DensitymapboxColorbarTitleSide = "right" // right
    DensitymapboxColorbarTitleSide_top DensitymapboxColorbarTitleSide = "top" // top
    DensitymapboxColorbarTitleSide_bottom DensitymapboxColorbarTitleSide = "bottom" // bottom
)

// DensitymapboxColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type DensitymapboxColorbarXanchor string

const (
    DensitymapboxColorbarXanchor_left DensitymapboxColorbarXanchor = "left" // left
    DensitymapboxColorbarXanchor_center DensitymapboxColorbarXanchor = "center" // center
    DensitymapboxColorbarXanchor_right DensitymapboxColorbarXanchor = "right" // right
)

// DensitymapboxColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type DensitymapboxColorbarYanchor string

const (
    DensitymapboxColorbarYanchor_top DensitymapboxColorbarYanchor = "top" // top
    DensitymapboxColorbarYanchor_middle DensitymapboxColorbarYanchor = "middle" // middle
    DensitymapboxColorbarYanchor_bottom DensitymapboxColorbarYanchor = "bottom" // bottom
)

// DensitymapboxHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type DensitymapboxHoverlabelAlign string

const (
    DensitymapboxHoverlabelAlign_left DensitymapboxHoverlabelAlign = "left" // left
    DensitymapboxHoverlabelAlign_right DensitymapboxHoverlabelAlign = "right" // right
    DensitymapboxHoverlabelAlign_auto DensitymapboxHoverlabelAlign = "auto" // auto
)

// DensitymapboxVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type DensitymapboxVisible interface{}

const (
    DensitymapboxVisibleTrue bool = true
    DensitymapboxVisibleFalse bool = false
    DensitymapboxVisibleLegendonly string = "legendonly"
)

// FunnelConstraintext Constrain the size of text inside or outside a bar to be no larger than the bar itself.
type FunnelConstraintext string

const (
    FunnelConstraintext_inside FunnelConstraintext = "inside" // inside
    FunnelConstraintext_outside FunnelConstraintext = "outside" // outside
    FunnelConstraintext_both FunnelConstraintext = "both" // both
    FunnelConstraintext_none FunnelConstraintext = "none" // none
)

// FunnelHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type FunnelHoverlabelAlign string

const (
    FunnelHoverlabelAlign_left FunnelHoverlabelAlign = "left" // left
    FunnelHoverlabelAlign_right FunnelHoverlabelAlign = "right" // right
    FunnelHoverlabelAlign_auto FunnelHoverlabelAlign = "auto" // auto
)

// FunnelInsidetextanchor Determines if texts are kept at center or start/end points in `textposition` *inside* mode.
type FunnelInsidetextanchor string

const (
    FunnelInsidetextanchor_end FunnelInsidetextanchor = "end" // end
    FunnelInsidetextanchor_middle FunnelInsidetextanchor = "middle" // middle
    FunnelInsidetextanchor_start FunnelInsidetextanchor = "start" // start
)

// FunnelMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type FunnelMarkerColorbarExponentformat string

const (
    FunnelMarkerColorbarExponentformat_none FunnelMarkerColorbarExponentformat = "none" // none
    FunnelMarkerColorbarExponentformat_e FunnelMarkerColorbarExponentformat = "e" // e
    FunnelMarkerColorbarExponentformat_E FunnelMarkerColorbarExponentformat = "E" // E
    FunnelMarkerColorbarExponentformat_power FunnelMarkerColorbarExponentformat = "power" // power
    FunnelMarkerColorbarExponentformat_SI FunnelMarkerColorbarExponentformat = "SI" // SI
    FunnelMarkerColorbarExponentformat_B FunnelMarkerColorbarExponentformat = "B" // B
)

// FunnelMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type FunnelMarkerColorbarLenmode string

const (
    FunnelMarkerColorbarLenmode_fraction FunnelMarkerColorbarLenmode = "fraction" // fraction
    FunnelMarkerColorbarLenmode_pixels FunnelMarkerColorbarLenmode = "pixels" // pixels
)

// FunnelMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type FunnelMarkerColorbarShowexponent string

const (
    FunnelMarkerColorbarShowexponent_all FunnelMarkerColorbarShowexponent = "all" // all
    FunnelMarkerColorbarShowexponent_first FunnelMarkerColorbarShowexponent = "first" // first
    FunnelMarkerColorbarShowexponent_last FunnelMarkerColorbarShowexponent = "last" // last
    FunnelMarkerColorbarShowexponent_none FunnelMarkerColorbarShowexponent = "none" // none
)

// FunnelMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type FunnelMarkerColorbarShowtickprefix string

const (
    FunnelMarkerColorbarShowtickprefix_all FunnelMarkerColorbarShowtickprefix = "all" // all
    FunnelMarkerColorbarShowtickprefix_first FunnelMarkerColorbarShowtickprefix = "first" // first
    FunnelMarkerColorbarShowtickprefix_last FunnelMarkerColorbarShowtickprefix = "last" // last
    FunnelMarkerColorbarShowtickprefix_none FunnelMarkerColorbarShowtickprefix = "none" // none
)

// FunnelMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type FunnelMarkerColorbarShowticksuffix string

const (
    FunnelMarkerColorbarShowticksuffix_all FunnelMarkerColorbarShowticksuffix = "all" // all
    FunnelMarkerColorbarShowticksuffix_first FunnelMarkerColorbarShowticksuffix = "first" // first
    FunnelMarkerColorbarShowticksuffix_last FunnelMarkerColorbarShowticksuffix = "last" // last
    FunnelMarkerColorbarShowticksuffix_none FunnelMarkerColorbarShowticksuffix = "none" // none
)

// FunnelMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type FunnelMarkerColorbarThicknessmode string

const (
    FunnelMarkerColorbarThicknessmode_fraction FunnelMarkerColorbarThicknessmode = "fraction" // fraction
    FunnelMarkerColorbarThicknessmode_pixels FunnelMarkerColorbarThicknessmode = "pixels" // pixels
)

// FunnelMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type FunnelMarkerColorbarTickmode string

const (
    FunnelMarkerColorbarTickmode_auto FunnelMarkerColorbarTickmode = "auto" // auto
    FunnelMarkerColorbarTickmode_linear FunnelMarkerColorbarTickmode = "linear" // linear
    FunnelMarkerColorbarTickmode_array FunnelMarkerColorbarTickmode = "array" // array
)

// FunnelMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type FunnelMarkerColorbarTicks string

const (
    FunnelMarkerColorbarTicks_outside FunnelMarkerColorbarTicks = "outside" // outside
    FunnelMarkerColorbarTicks_inside FunnelMarkerColorbarTicks = "inside" // inside
)

// FunnelMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type FunnelMarkerColorbarTitleSide string

const (
    FunnelMarkerColorbarTitleSide_right FunnelMarkerColorbarTitleSide = "right" // right
    FunnelMarkerColorbarTitleSide_top FunnelMarkerColorbarTitleSide = "top" // top
    FunnelMarkerColorbarTitleSide_bottom FunnelMarkerColorbarTitleSide = "bottom" // bottom
)

// FunnelMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type FunnelMarkerColorbarXanchor string

const (
    FunnelMarkerColorbarXanchor_left FunnelMarkerColorbarXanchor = "left" // left
    FunnelMarkerColorbarXanchor_center FunnelMarkerColorbarXanchor = "center" // center
    FunnelMarkerColorbarXanchor_right FunnelMarkerColorbarXanchor = "right" // right
)

// FunnelMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type FunnelMarkerColorbarYanchor string

const (
    FunnelMarkerColorbarYanchor_top FunnelMarkerColorbarYanchor = "top" // top
    FunnelMarkerColorbarYanchor_middle FunnelMarkerColorbarYanchor = "middle" // middle
    FunnelMarkerColorbarYanchor_bottom FunnelMarkerColorbarYanchor = "bottom" // bottom
)

// FunnelOrientation Sets the orientation of the funnels. With *v* (*h*), the value of the each bar spans along the vertical (horizontal). By default funnels are tend to be oriented horizontally; unless only *y* array is presented or orientation is set to *v*. Also regarding graphs including only 'horizontal' funnels, *autorange* on the *y-axis* are set to *reversed*.
type FunnelOrientation string

const (
    FunnelOrientation_v FunnelOrientation = "v" // v
    FunnelOrientation_h FunnelOrientation = "h" // h
)

// FunnelTextposition Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside.
type FunnelTextposition string

const (
    FunnelTextposition_inside FunnelTextposition = "inside" // inside
    FunnelTextposition_outside FunnelTextposition = "outside" // outside
    FunnelTextposition_auto FunnelTextposition = "auto" // auto
    FunnelTextposition_none FunnelTextposition = "none" // none
)

// FunnelVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type FunnelVisible interface{}

const (
    FunnelVisibleTrue bool = true
    FunnelVisibleFalse bool = false
    FunnelVisibleLegendonly string = "legendonly"
)

// FunnelareaHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type FunnelareaHoverlabelAlign string

const (
    FunnelareaHoverlabelAlign_left FunnelareaHoverlabelAlign = "left" // left
    FunnelareaHoverlabelAlign_right FunnelareaHoverlabelAlign = "right" // right
    FunnelareaHoverlabelAlign_auto FunnelareaHoverlabelAlign = "auto" // auto
)

// FunnelareaTextposition Specifies the location of the `textinfo`.
type FunnelareaTextposition string

const (
    FunnelareaTextposition_inside FunnelareaTextposition = "inside" // inside
    FunnelareaTextposition_none FunnelareaTextposition = "none" // none
)

// FunnelareaTitlePosition Specifies the location of the `title`. Note that the title's position used to be set by the now deprecated `titleposition` attribute.
type FunnelareaTitlePosition string

const (
    FunnelareaTitlePosition_topleft FunnelareaTitlePosition = "top left" // top left
    FunnelareaTitlePosition_topcenter FunnelareaTitlePosition = "top center" // top center
    FunnelareaTitlePosition_topright FunnelareaTitlePosition = "top right" // top right
)

// FunnelareaVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type FunnelareaVisible interface{}

const (
    FunnelareaVisibleTrue bool = true
    FunnelareaVisibleFalse bool = false
    FunnelareaVisibleLegendonly string = "legendonly"
)

// HeatmapColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type HeatmapColorbarExponentformat string

const (
    HeatmapColorbarExponentformat_none HeatmapColorbarExponentformat = "none" // none
    HeatmapColorbarExponentformat_e HeatmapColorbarExponentformat = "e" // e
    HeatmapColorbarExponentformat_E HeatmapColorbarExponentformat = "E" // E
    HeatmapColorbarExponentformat_power HeatmapColorbarExponentformat = "power" // power
    HeatmapColorbarExponentformat_SI HeatmapColorbarExponentformat = "SI" // SI
    HeatmapColorbarExponentformat_B HeatmapColorbarExponentformat = "B" // B
)

// HeatmapColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type HeatmapColorbarLenmode string

const (
    HeatmapColorbarLenmode_fraction HeatmapColorbarLenmode = "fraction" // fraction
    HeatmapColorbarLenmode_pixels HeatmapColorbarLenmode = "pixels" // pixels
)

// HeatmapColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type HeatmapColorbarShowexponent string

const (
    HeatmapColorbarShowexponent_all HeatmapColorbarShowexponent = "all" // all
    HeatmapColorbarShowexponent_first HeatmapColorbarShowexponent = "first" // first
    HeatmapColorbarShowexponent_last HeatmapColorbarShowexponent = "last" // last
    HeatmapColorbarShowexponent_none HeatmapColorbarShowexponent = "none" // none
)

// HeatmapColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type HeatmapColorbarShowtickprefix string

const (
    HeatmapColorbarShowtickprefix_all HeatmapColorbarShowtickprefix = "all" // all
    HeatmapColorbarShowtickprefix_first HeatmapColorbarShowtickprefix = "first" // first
    HeatmapColorbarShowtickprefix_last HeatmapColorbarShowtickprefix = "last" // last
    HeatmapColorbarShowtickprefix_none HeatmapColorbarShowtickprefix = "none" // none
)

// HeatmapColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type HeatmapColorbarShowticksuffix string

const (
    HeatmapColorbarShowticksuffix_all HeatmapColorbarShowticksuffix = "all" // all
    HeatmapColorbarShowticksuffix_first HeatmapColorbarShowticksuffix = "first" // first
    HeatmapColorbarShowticksuffix_last HeatmapColorbarShowticksuffix = "last" // last
    HeatmapColorbarShowticksuffix_none HeatmapColorbarShowticksuffix = "none" // none
)

// HeatmapColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type HeatmapColorbarThicknessmode string

const (
    HeatmapColorbarThicknessmode_fraction HeatmapColorbarThicknessmode = "fraction" // fraction
    HeatmapColorbarThicknessmode_pixels HeatmapColorbarThicknessmode = "pixels" // pixels
)

// HeatmapColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type HeatmapColorbarTickmode string

const (
    HeatmapColorbarTickmode_auto HeatmapColorbarTickmode = "auto" // auto
    HeatmapColorbarTickmode_linear HeatmapColorbarTickmode = "linear" // linear
    HeatmapColorbarTickmode_array HeatmapColorbarTickmode = "array" // array
)

// HeatmapColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type HeatmapColorbarTicks string

const (
    HeatmapColorbarTicks_outside HeatmapColorbarTicks = "outside" // outside
    HeatmapColorbarTicks_inside HeatmapColorbarTicks = "inside" // inside
)

// HeatmapColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type HeatmapColorbarTitleSide string

const (
    HeatmapColorbarTitleSide_right HeatmapColorbarTitleSide = "right" // right
    HeatmapColorbarTitleSide_top HeatmapColorbarTitleSide = "top" // top
    HeatmapColorbarTitleSide_bottom HeatmapColorbarTitleSide = "bottom" // bottom
)

// HeatmapColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type HeatmapColorbarXanchor string

const (
    HeatmapColorbarXanchor_left HeatmapColorbarXanchor = "left" // left
    HeatmapColorbarXanchor_center HeatmapColorbarXanchor = "center" // center
    HeatmapColorbarXanchor_right HeatmapColorbarXanchor = "right" // right
)

// HeatmapColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type HeatmapColorbarYanchor string

const (
    HeatmapColorbarYanchor_top HeatmapColorbarYanchor = "top" // top
    HeatmapColorbarYanchor_middle HeatmapColorbarYanchor = "middle" // middle
    HeatmapColorbarYanchor_bottom HeatmapColorbarYanchor = "bottom" // bottom
)

// HeatmapHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type HeatmapHoverlabelAlign string

const (
    HeatmapHoverlabelAlign_left HeatmapHoverlabelAlign = "left" // left
    HeatmapHoverlabelAlign_right HeatmapHoverlabelAlign = "right" // right
    HeatmapHoverlabelAlign_auto HeatmapHoverlabelAlign = "auto" // auto
)

// HeatmapVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type HeatmapVisible interface{}

const (
    HeatmapVisibleTrue bool = true
    HeatmapVisibleFalse bool = false
    HeatmapVisibleLegendonly string = "legendonly"
)

// HeatmapXcalendar Sets the calendar system to use with `x` date data.
type HeatmapXcalendar string

const (
    HeatmapXcalendar_gregorian HeatmapXcalendar = "gregorian" // gregorian
    HeatmapXcalendar_chinese HeatmapXcalendar = "chinese" // chinese
    HeatmapXcalendar_coptic HeatmapXcalendar = "coptic" // coptic
    HeatmapXcalendar_discworld HeatmapXcalendar = "discworld" // discworld
    HeatmapXcalendar_ethiopian HeatmapXcalendar = "ethiopian" // ethiopian
    HeatmapXcalendar_hebrew HeatmapXcalendar = "hebrew" // hebrew
    HeatmapXcalendar_islamic HeatmapXcalendar = "islamic" // islamic
    HeatmapXcalendar_julian HeatmapXcalendar = "julian" // julian
    HeatmapXcalendar_mayan HeatmapXcalendar = "mayan" // mayan
    HeatmapXcalendar_nanakshahi HeatmapXcalendar = "nanakshahi" // nanakshahi
    HeatmapXcalendar_nepali HeatmapXcalendar = "nepali" // nepali
    HeatmapXcalendar_persian HeatmapXcalendar = "persian" // persian
    HeatmapXcalendar_jalali HeatmapXcalendar = "jalali" // jalali
    HeatmapXcalendar_taiwan HeatmapXcalendar = "taiwan" // taiwan
    HeatmapXcalendar_thai HeatmapXcalendar = "thai" // thai
    HeatmapXcalendar_ummalqura HeatmapXcalendar = "ummalqura" // ummalqura
)

// HeatmapXtype If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
type HeatmapXtype string

const (
    HeatmapXtype_array HeatmapXtype = "array" // array
    HeatmapXtype_scaled HeatmapXtype = "scaled" // scaled
)

// HeatmapYcalendar Sets the calendar system to use with `y` date data.
type HeatmapYcalendar string

const (
    HeatmapYcalendar_gregorian HeatmapYcalendar = "gregorian" // gregorian
    HeatmapYcalendar_chinese HeatmapYcalendar = "chinese" // chinese
    HeatmapYcalendar_coptic HeatmapYcalendar = "coptic" // coptic
    HeatmapYcalendar_discworld HeatmapYcalendar = "discworld" // discworld
    HeatmapYcalendar_ethiopian HeatmapYcalendar = "ethiopian" // ethiopian
    HeatmapYcalendar_hebrew HeatmapYcalendar = "hebrew" // hebrew
    HeatmapYcalendar_islamic HeatmapYcalendar = "islamic" // islamic
    HeatmapYcalendar_julian HeatmapYcalendar = "julian" // julian
    HeatmapYcalendar_mayan HeatmapYcalendar = "mayan" // mayan
    HeatmapYcalendar_nanakshahi HeatmapYcalendar = "nanakshahi" // nanakshahi
    HeatmapYcalendar_nepali HeatmapYcalendar = "nepali" // nepali
    HeatmapYcalendar_persian HeatmapYcalendar = "persian" // persian
    HeatmapYcalendar_jalali HeatmapYcalendar = "jalali" // jalali
    HeatmapYcalendar_taiwan HeatmapYcalendar = "taiwan" // taiwan
    HeatmapYcalendar_thai HeatmapYcalendar = "thai" // thai
    HeatmapYcalendar_ummalqura HeatmapYcalendar = "ummalqura" // ummalqura
)

// HeatmapYtype If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
type HeatmapYtype string

const (
    HeatmapYtype_array HeatmapYtype = "array" // array
    HeatmapYtype_scaled HeatmapYtype = "scaled" // scaled
)

// HeatmapZsmooth Picks a smoothing algorithm use to smooth `z` data.
type HeatmapZsmooth interface{}

const (
    HeatmapZsmoothFast string = "fast"
    HeatmapZsmoothBest string = "best"
    HeatmapZsmoothFalse bool = false
)

// HeatmapglColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type HeatmapglColorbarExponentformat string

const (
    HeatmapglColorbarExponentformat_none HeatmapglColorbarExponentformat = "none" // none
    HeatmapglColorbarExponentformat_e HeatmapglColorbarExponentformat = "e" // e
    HeatmapglColorbarExponentformat_E HeatmapglColorbarExponentformat = "E" // E
    HeatmapglColorbarExponentformat_power HeatmapglColorbarExponentformat = "power" // power
    HeatmapglColorbarExponentformat_SI HeatmapglColorbarExponentformat = "SI" // SI
    HeatmapglColorbarExponentformat_B HeatmapglColorbarExponentformat = "B" // B
)

// HeatmapglColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type HeatmapglColorbarLenmode string

const (
    HeatmapglColorbarLenmode_fraction HeatmapglColorbarLenmode = "fraction" // fraction
    HeatmapglColorbarLenmode_pixels HeatmapglColorbarLenmode = "pixels" // pixels
)

// HeatmapglColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type HeatmapglColorbarShowexponent string

const (
    HeatmapglColorbarShowexponent_all HeatmapglColorbarShowexponent = "all" // all
    HeatmapglColorbarShowexponent_first HeatmapglColorbarShowexponent = "first" // first
    HeatmapglColorbarShowexponent_last HeatmapglColorbarShowexponent = "last" // last
    HeatmapglColorbarShowexponent_none HeatmapglColorbarShowexponent = "none" // none
)

// HeatmapglColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type HeatmapglColorbarShowtickprefix string

const (
    HeatmapglColorbarShowtickprefix_all HeatmapglColorbarShowtickprefix = "all" // all
    HeatmapglColorbarShowtickprefix_first HeatmapglColorbarShowtickprefix = "first" // first
    HeatmapglColorbarShowtickprefix_last HeatmapglColorbarShowtickprefix = "last" // last
    HeatmapglColorbarShowtickprefix_none HeatmapglColorbarShowtickprefix = "none" // none
)

// HeatmapglColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type HeatmapglColorbarShowticksuffix string

const (
    HeatmapglColorbarShowticksuffix_all HeatmapglColorbarShowticksuffix = "all" // all
    HeatmapglColorbarShowticksuffix_first HeatmapglColorbarShowticksuffix = "first" // first
    HeatmapglColorbarShowticksuffix_last HeatmapglColorbarShowticksuffix = "last" // last
    HeatmapglColorbarShowticksuffix_none HeatmapglColorbarShowticksuffix = "none" // none
)

// HeatmapglColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type HeatmapglColorbarThicknessmode string

const (
    HeatmapglColorbarThicknessmode_fraction HeatmapglColorbarThicknessmode = "fraction" // fraction
    HeatmapglColorbarThicknessmode_pixels HeatmapglColorbarThicknessmode = "pixels" // pixels
)

// HeatmapglColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type HeatmapglColorbarTickmode string

const (
    HeatmapglColorbarTickmode_auto HeatmapglColorbarTickmode = "auto" // auto
    HeatmapglColorbarTickmode_linear HeatmapglColorbarTickmode = "linear" // linear
    HeatmapglColorbarTickmode_array HeatmapglColorbarTickmode = "array" // array
)

// HeatmapglColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type HeatmapglColorbarTicks string

const (
    HeatmapglColorbarTicks_outside HeatmapglColorbarTicks = "outside" // outside
    HeatmapglColorbarTicks_inside HeatmapglColorbarTicks = "inside" // inside
)

// HeatmapglColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type HeatmapglColorbarTitleSide string

const (
    HeatmapglColorbarTitleSide_right HeatmapglColorbarTitleSide = "right" // right
    HeatmapglColorbarTitleSide_top HeatmapglColorbarTitleSide = "top" // top
    HeatmapglColorbarTitleSide_bottom HeatmapglColorbarTitleSide = "bottom" // bottom
)

// HeatmapglColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type HeatmapglColorbarXanchor string

const (
    HeatmapglColorbarXanchor_left HeatmapglColorbarXanchor = "left" // left
    HeatmapglColorbarXanchor_center HeatmapglColorbarXanchor = "center" // center
    HeatmapglColorbarXanchor_right HeatmapglColorbarXanchor = "right" // right
)

// HeatmapglColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type HeatmapglColorbarYanchor string

const (
    HeatmapglColorbarYanchor_top HeatmapglColorbarYanchor = "top" // top
    HeatmapglColorbarYanchor_middle HeatmapglColorbarYanchor = "middle" // middle
    HeatmapglColorbarYanchor_bottom HeatmapglColorbarYanchor = "bottom" // bottom
)

// HeatmapglHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type HeatmapglHoverlabelAlign string

const (
    HeatmapglHoverlabelAlign_left HeatmapglHoverlabelAlign = "left" // left
    HeatmapglHoverlabelAlign_right HeatmapglHoverlabelAlign = "right" // right
    HeatmapglHoverlabelAlign_auto HeatmapglHoverlabelAlign = "auto" // auto
)

// HeatmapglVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type HeatmapglVisible interface{}

const (
    HeatmapglVisibleTrue bool = true
    HeatmapglVisibleFalse bool = false
    HeatmapglVisibleLegendonly string = "legendonly"
)

// HeatmapglXtype If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
type HeatmapglXtype string

const (
    HeatmapglXtype_array HeatmapglXtype = "array" // array
    HeatmapglXtype_scaled HeatmapglXtype = "scaled" // scaled
)

// HeatmapglYtype If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
type HeatmapglYtype string

const (
    HeatmapglYtype_array HeatmapglYtype = "array" // array
    HeatmapglYtype_scaled HeatmapglYtype = "scaled" // scaled
)

// Histogram2dColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type Histogram2dColorbarExponentformat string

const (
    Histogram2dColorbarExponentformat_none Histogram2dColorbarExponentformat = "none" // none
    Histogram2dColorbarExponentformat_e Histogram2dColorbarExponentformat = "e" // e
    Histogram2dColorbarExponentformat_E Histogram2dColorbarExponentformat = "E" // E
    Histogram2dColorbarExponentformat_power Histogram2dColorbarExponentformat = "power" // power
    Histogram2dColorbarExponentformat_SI Histogram2dColorbarExponentformat = "SI" // SI
    Histogram2dColorbarExponentformat_B Histogram2dColorbarExponentformat = "B" // B
)

// Histogram2dColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type Histogram2dColorbarLenmode string

const (
    Histogram2dColorbarLenmode_fraction Histogram2dColorbarLenmode = "fraction" // fraction
    Histogram2dColorbarLenmode_pixels Histogram2dColorbarLenmode = "pixels" // pixels
)

// Histogram2dColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type Histogram2dColorbarShowexponent string

const (
    Histogram2dColorbarShowexponent_all Histogram2dColorbarShowexponent = "all" // all
    Histogram2dColorbarShowexponent_first Histogram2dColorbarShowexponent = "first" // first
    Histogram2dColorbarShowexponent_last Histogram2dColorbarShowexponent = "last" // last
    Histogram2dColorbarShowexponent_none Histogram2dColorbarShowexponent = "none" // none
)

// Histogram2dColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type Histogram2dColorbarShowtickprefix string

const (
    Histogram2dColorbarShowtickprefix_all Histogram2dColorbarShowtickprefix = "all" // all
    Histogram2dColorbarShowtickprefix_first Histogram2dColorbarShowtickprefix = "first" // first
    Histogram2dColorbarShowtickprefix_last Histogram2dColorbarShowtickprefix = "last" // last
    Histogram2dColorbarShowtickprefix_none Histogram2dColorbarShowtickprefix = "none" // none
)

// Histogram2dColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type Histogram2dColorbarShowticksuffix string

const (
    Histogram2dColorbarShowticksuffix_all Histogram2dColorbarShowticksuffix = "all" // all
    Histogram2dColorbarShowticksuffix_first Histogram2dColorbarShowticksuffix = "first" // first
    Histogram2dColorbarShowticksuffix_last Histogram2dColorbarShowticksuffix = "last" // last
    Histogram2dColorbarShowticksuffix_none Histogram2dColorbarShowticksuffix = "none" // none
)

// Histogram2dColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type Histogram2dColorbarThicknessmode string

const (
    Histogram2dColorbarThicknessmode_fraction Histogram2dColorbarThicknessmode = "fraction" // fraction
    Histogram2dColorbarThicknessmode_pixels Histogram2dColorbarThicknessmode = "pixels" // pixels
)

// Histogram2dColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type Histogram2dColorbarTickmode string

const (
    Histogram2dColorbarTickmode_auto Histogram2dColorbarTickmode = "auto" // auto
    Histogram2dColorbarTickmode_linear Histogram2dColorbarTickmode = "linear" // linear
    Histogram2dColorbarTickmode_array Histogram2dColorbarTickmode = "array" // array
)

// Histogram2dColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type Histogram2dColorbarTicks string

const (
    Histogram2dColorbarTicks_outside Histogram2dColorbarTicks = "outside" // outside
    Histogram2dColorbarTicks_inside Histogram2dColorbarTicks = "inside" // inside
)

// Histogram2dColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type Histogram2dColorbarTitleSide string

const (
    Histogram2dColorbarTitleSide_right Histogram2dColorbarTitleSide = "right" // right
    Histogram2dColorbarTitleSide_top Histogram2dColorbarTitleSide = "top" // top
    Histogram2dColorbarTitleSide_bottom Histogram2dColorbarTitleSide = "bottom" // bottom
)

// Histogram2dColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type Histogram2dColorbarXanchor string

const (
    Histogram2dColorbarXanchor_left Histogram2dColorbarXanchor = "left" // left
    Histogram2dColorbarXanchor_center Histogram2dColorbarXanchor = "center" // center
    Histogram2dColorbarXanchor_right Histogram2dColorbarXanchor = "right" // right
)

// Histogram2dColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type Histogram2dColorbarYanchor string

const (
    Histogram2dColorbarYanchor_top Histogram2dColorbarYanchor = "top" // top
    Histogram2dColorbarYanchor_middle Histogram2dColorbarYanchor = "middle" // middle
    Histogram2dColorbarYanchor_bottom Histogram2dColorbarYanchor = "bottom" // bottom
)

// Histogram2dHistfunc Specifies the binning function used for this histogram trace. If *count*, the histogram values are computed by counting the number of values lying inside each bin. If *sum*, *avg*, *min*, *max*, the histogram values are computed using the sum, the average, the minimum or the maximum of the values lying inside each bin respectively.
type Histogram2dHistfunc string

const (
    Histogram2dHistfunc_count Histogram2dHistfunc = "count" // count
    Histogram2dHistfunc_sum Histogram2dHistfunc = "sum" // sum
    Histogram2dHistfunc_avg Histogram2dHistfunc = "avg" // avg
    Histogram2dHistfunc_min Histogram2dHistfunc = "min" // min
    Histogram2dHistfunc_max Histogram2dHistfunc = "max" // max
)

// Histogram2dHistnorm Specifies the type of normalization used for this histogram trace. If **, the span of each bar corresponds to the number of occurrences (i.e. the number of data points lying inside the bins). If *percent* / *probability*, the span of each bar corresponds to the percentage / fraction of occurrences with respect to the total number of sample points (here, the sum of all bin HEIGHTS equals 100% / 1). If *density*, the span of each bar corresponds to the number of occurrences in a bin divided by the size of the bin interval (here, the sum of all bin AREAS equals the total number of sample points). If *probability density*, the area of each bar corresponds to the probability that an event will fall into the corresponding bin (here, the sum of all bin AREAS equals 1).
type Histogram2dHistnorm string

const (
    Histogram2dHistnorm_percent Histogram2dHistnorm = "percent" // percent
    Histogram2dHistnorm_probability Histogram2dHistnorm = "probability" // probability
    Histogram2dHistnorm_density Histogram2dHistnorm = "density" // density
    Histogram2dHistnorm_probabilitydensity Histogram2dHistnorm = "probability density" // probability density
)

// Histogram2dHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type Histogram2dHoverlabelAlign string

const (
    Histogram2dHoverlabelAlign_left Histogram2dHoverlabelAlign = "left" // left
    Histogram2dHoverlabelAlign_right Histogram2dHoverlabelAlign = "right" // right
    Histogram2dHoverlabelAlign_auto Histogram2dHoverlabelAlign = "auto" // auto
)

// Histogram2dVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type Histogram2dVisible interface{}

const (
    Histogram2dVisibleTrue bool = true
    Histogram2dVisibleFalse bool = false
    Histogram2dVisibleLegendonly string = "legendonly"
)

// Histogram2dXcalendar Sets the calendar system to use with `x` date data.
type Histogram2dXcalendar string

const (
    Histogram2dXcalendar_gregorian Histogram2dXcalendar = "gregorian" // gregorian
    Histogram2dXcalendar_chinese Histogram2dXcalendar = "chinese" // chinese
    Histogram2dXcalendar_coptic Histogram2dXcalendar = "coptic" // coptic
    Histogram2dXcalendar_discworld Histogram2dXcalendar = "discworld" // discworld
    Histogram2dXcalendar_ethiopian Histogram2dXcalendar = "ethiopian" // ethiopian
    Histogram2dXcalendar_hebrew Histogram2dXcalendar = "hebrew" // hebrew
    Histogram2dXcalendar_islamic Histogram2dXcalendar = "islamic" // islamic
    Histogram2dXcalendar_julian Histogram2dXcalendar = "julian" // julian
    Histogram2dXcalendar_mayan Histogram2dXcalendar = "mayan" // mayan
    Histogram2dXcalendar_nanakshahi Histogram2dXcalendar = "nanakshahi" // nanakshahi
    Histogram2dXcalendar_nepali Histogram2dXcalendar = "nepali" // nepali
    Histogram2dXcalendar_persian Histogram2dXcalendar = "persian" // persian
    Histogram2dXcalendar_jalali Histogram2dXcalendar = "jalali" // jalali
    Histogram2dXcalendar_taiwan Histogram2dXcalendar = "taiwan" // taiwan
    Histogram2dXcalendar_thai Histogram2dXcalendar = "thai" // thai
    Histogram2dXcalendar_ummalqura Histogram2dXcalendar = "ummalqura" // ummalqura
)

// Histogram2dYcalendar Sets the calendar system to use with `y` date data.
type Histogram2dYcalendar string

const (
    Histogram2dYcalendar_gregorian Histogram2dYcalendar = "gregorian" // gregorian
    Histogram2dYcalendar_chinese Histogram2dYcalendar = "chinese" // chinese
    Histogram2dYcalendar_coptic Histogram2dYcalendar = "coptic" // coptic
    Histogram2dYcalendar_discworld Histogram2dYcalendar = "discworld" // discworld
    Histogram2dYcalendar_ethiopian Histogram2dYcalendar = "ethiopian" // ethiopian
    Histogram2dYcalendar_hebrew Histogram2dYcalendar = "hebrew" // hebrew
    Histogram2dYcalendar_islamic Histogram2dYcalendar = "islamic" // islamic
    Histogram2dYcalendar_julian Histogram2dYcalendar = "julian" // julian
    Histogram2dYcalendar_mayan Histogram2dYcalendar = "mayan" // mayan
    Histogram2dYcalendar_nanakshahi Histogram2dYcalendar = "nanakshahi" // nanakshahi
    Histogram2dYcalendar_nepali Histogram2dYcalendar = "nepali" // nepali
    Histogram2dYcalendar_persian Histogram2dYcalendar = "persian" // persian
    Histogram2dYcalendar_jalali Histogram2dYcalendar = "jalali" // jalali
    Histogram2dYcalendar_taiwan Histogram2dYcalendar = "taiwan" // taiwan
    Histogram2dYcalendar_thai Histogram2dYcalendar = "thai" // thai
    Histogram2dYcalendar_ummalqura Histogram2dYcalendar = "ummalqura" // ummalqura
)

// Histogram2dZsmooth Picks a smoothing algorithm use to smooth `z` data.
type Histogram2dZsmooth interface{}

const (
    Histogram2dZsmoothFast string = "fast"
    Histogram2dZsmoothBest string = "best"
    Histogram2dZsmoothFalse bool = false
)

// Histogram2dcontourColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type Histogram2dcontourColorbarExponentformat string

const (
    Histogram2dcontourColorbarExponentformat_none Histogram2dcontourColorbarExponentformat = "none" // none
    Histogram2dcontourColorbarExponentformat_e Histogram2dcontourColorbarExponentformat = "e" // e
    Histogram2dcontourColorbarExponentformat_E Histogram2dcontourColorbarExponentformat = "E" // E
    Histogram2dcontourColorbarExponentformat_power Histogram2dcontourColorbarExponentformat = "power" // power
    Histogram2dcontourColorbarExponentformat_SI Histogram2dcontourColorbarExponentformat = "SI" // SI
    Histogram2dcontourColorbarExponentformat_B Histogram2dcontourColorbarExponentformat = "B" // B
)

// Histogram2dcontourColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type Histogram2dcontourColorbarLenmode string

const (
    Histogram2dcontourColorbarLenmode_fraction Histogram2dcontourColorbarLenmode = "fraction" // fraction
    Histogram2dcontourColorbarLenmode_pixels Histogram2dcontourColorbarLenmode = "pixels" // pixels
)

// Histogram2dcontourColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type Histogram2dcontourColorbarShowexponent string

const (
    Histogram2dcontourColorbarShowexponent_all Histogram2dcontourColorbarShowexponent = "all" // all
    Histogram2dcontourColorbarShowexponent_first Histogram2dcontourColorbarShowexponent = "first" // first
    Histogram2dcontourColorbarShowexponent_last Histogram2dcontourColorbarShowexponent = "last" // last
    Histogram2dcontourColorbarShowexponent_none Histogram2dcontourColorbarShowexponent = "none" // none
)

// Histogram2dcontourColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type Histogram2dcontourColorbarShowtickprefix string

const (
    Histogram2dcontourColorbarShowtickprefix_all Histogram2dcontourColorbarShowtickprefix = "all" // all
    Histogram2dcontourColorbarShowtickprefix_first Histogram2dcontourColorbarShowtickprefix = "first" // first
    Histogram2dcontourColorbarShowtickprefix_last Histogram2dcontourColorbarShowtickprefix = "last" // last
    Histogram2dcontourColorbarShowtickprefix_none Histogram2dcontourColorbarShowtickprefix = "none" // none
)

// Histogram2dcontourColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type Histogram2dcontourColorbarShowticksuffix string

const (
    Histogram2dcontourColorbarShowticksuffix_all Histogram2dcontourColorbarShowticksuffix = "all" // all
    Histogram2dcontourColorbarShowticksuffix_first Histogram2dcontourColorbarShowticksuffix = "first" // first
    Histogram2dcontourColorbarShowticksuffix_last Histogram2dcontourColorbarShowticksuffix = "last" // last
    Histogram2dcontourColorbarShowticksuffix_none Histogram2dcontourColorbarShowticksuffix = "none" // none
)

// Histogram2dcontourColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type Histogram2dcontourColorbarThicknessmode string

const (
    Histogram2dcontourColorbarThicknessmode_fraction Histogram2dcontourColorbarThicknessmode = "fraction" // fraction
    Histogram2dcontourColorbarThicknessmode_pixels Histogram2dcontourColorbarThicknessmode = "pixels" // pixels
)

// Histogram2dcontourColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type Histogram2dcontourColorbarTickmode string

const (
    Histogram2dcontourColorbarTickmode_auto Histogram2dcontourColorbarTickmode = "auto" // auto
    Histogram2dcontourColorbarTickmode_linear Histogram2dcontourColorbarTickmode = "linear" // linear
    Histogram2dcontourColorbarTickmode_array Histogram2dcontourColorbarTickmode = "array" // array
)

// Histogram2dcontourColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type Histogram2dcontourColorbarTicks string

const (
    Histogram2dcontourColorbarTicks_outside Histogram2dcontourColorbarTicks = "outside" // outside
    Histogram2dcontourColorbarTicks_inside Histogram2dcontourColorbarTicks = "inside" // inside
)

// Histogram2dcontourColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type Histogram2dcontourColorbarTitleSide string

const (
    Histogram2dcontourColorbarTitleSide_right Histogram2dcontourColorbarTitleSide = "right" // right
    Histogram2dcontourColorbarTitleSide_top Histogram2dcontourColorbarTitleSide = "top" // top
    Histogram2dcontourColorbarTitleSide_bottom Histogram2dcontourColorbarTitleSide = "bottom" // bottom
)

// Histogram2dcontourColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type Histogram2dcontourColorbarXanchor string

const (
    Histogram2dcontourColorbarXanchor_left Histogram2dcontourColorbarXanchor = "left" // left
    Histogram2dcontourColorbarXanchor_center Histogram2dcontourColorbarXanchor = "center" // center
    Histogram2dcontourColorbarXanchor_right Histogram2dcontourColorbarXanchor = "right" // right
)

// Histogram2dcontourColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type Histogram2dcontourColorbarYanchor string

const (
    Histogram2dcontourColorbarYanchor_top Histogram2dcontourColorbarYanchor = "top" // top
    Histogram2dcontourColorbarYanchor_middle Histogram2dcontourColorbarYanchor = "middle" // middle
    Histogram2dcontourColorbarYanchor_bottom Histogram2dcontourColorbarYanchor = "bottom" // bottom
)

// Histogram2dcontourContoursColoring Determines the coloring method showing the contour values. If *fill*, coloring is done evenly between each contour level If *heatmap*, a heatmap gradient coloring is applied between each contour level. If *lines*, coloring is done on the contour lines. If *none*, no coloring is applied on this trace.
type Histogram2dcontourContoursColoring string

const (
    Histogram2dcontourContoursColoring_fill Histogram2dcontourContoursColoring = "fill" // fill
    Histogram2dcontourContoursColoring_heatmap Histogram2dcontourContoursColoring = "heatmap" // heatmap
    Histogram2dcontourContoursColoring_lines Histogram2dcontourContoursColoring = "lines" // lines
    Histogram2dcontourContoursColoring_none Histogram2dcontourContoursColoring = "none" // none
)

// Histogram2dcontourContoursOperation Sets the constraint operation. *=* keeps regions equal to `value` *<* and *<=* keep regions less than `value` *>* and *>=* keep regions greater than `value` *[]*, *()*, *[)*, and *(]* keep regions inside `value[0]` to `value[1]` *][*, *)(*, *](*, *)[* keep regions outside `value[0]` to value[1]` Open vs. closed intervals make no difference to constraint display, but all versions are allowed for consistency with filter transforms.
type Histogram2dcontourContoursOperation string

const (
    Histogram2dcontourContoursOperation_eq Histogram2dcontourContoursOperation = "=" // =
    Histogram2dcontourContoursOperation_lt Histogram2dcontourContoursOperation = "<" // <
    Histogram2dcontourContoursOperation_gteq Histogram2dcontourContoursOperation = ">=" // >=
    Histogram2dcontourContoursOperation_gt Histogram2dcontourContoursOperation = ">" // >
    Histogram2dcontourContoursOperation_lteq Histogram2dcontourContoursOperation = "<=" // <=
    Histogram2dcontourContoursOperation__lbracket__rbracket_ Histogram2dcontourContoursOperation = "[]" // []
    Histogram2dcontourContoursOperation__lpar__rpar_ Histogram2dcontourContoursOperation = "()" // ()
    Histogram2dcontourContoursOperation__lbracket__rpar_ Histogram2dcontourContoursOperation = "[)" // [)
    Histogram2dcontourContoursOperation__lpar__rbracket_ Histogram2dcontourContoursOperation = "(]" // (]
    Histogram2dcontourContoursOperation__rbracket__lbracket_ Histogram2dcontourContoursOperation = "][" // ][
    Histogram2dcontourContoursOperation__rpar__lpar_ Histogram2dcontourContoursOperation = ")(" // )(
    Histogram2dcontourContoursOperation__rbracket__lpar_ Histogram2dcontourContoursOperation = "](" // ](
    Histogram2dcontourContoursOperation__rpar__lbracket_ Histogram2dcontourContoursOperation = ")[" // )[
)

// Histogram2dcontourContoursType If `levels`, the data is represented as a contour plot with multiple levels displayed. If `constraint`, the data is represented as constraints with the invalid region shaded as specified by the `operation` and `value` parameters.
type Histogram2dcontourContoursType string

const (
    Histogram2dcontourContoursType_levels Histogram2dcontourContoursType = "levels" // levels
    Histogram2dcontourContoursType_constraint Histogram2dcontourContoursType = "constraint" // constraint
)

// Histogram2dcontourHistfunc Specifies the binning function used for this histogram trace. If *count*, the histogram values are computed by counting the number of values lying inside each bin. If *sum*, *avg*, *min*, *max*, the histogram values are computed using the sum, the average, the minimum or the maximum of the values lying inside each bin respectively.
type Histogram2dcontourHistfunc string

const (
    Histogram2dcontourHistfunc_count Histogram2dcontourHistfunc = "count" // count
    Histogram2dcontourHistfunc_sum Histogram2dcontourHistfunc = "sum" // sum
    Histogram2dcontourHistfunc_avg Histogram2dcontourHistfunc = "avg" // avg
    Histogram2dcontourHistfunc_min Histogram2dcontourHistfunc = "min" // min
    Histogram2dcontourHistfunc_max Histogram2dcontourHistfunc = "max" // max
)

// Histogram2dcontourHistnorm Specifies the type of normalization used for this histogram trace. If **, the span of each bar corresponds to the number of occurrences (i.e. the number of data points lying inside the bins). If *percent* / *probability*, the span of each bar corresponds to the percentage / fraction of occurrences with respect to the total number of sample points (here, the sum of all bin HEIGHTS equals 100% / 1). If *density*, the span of each bar corresponds to the number of occurrences in a bin divided by the size of the bin interval (here, the sum of all bin AREAS equals the total number of sample points). If *probability density*, the area of each bar corresponds to the probability that an event will fall into the corresponding bin (here, the sum of all bin AREAS equals 1).
type Histogram2dcontourHistnorm string

const (
    Histogram2dcontourHistnorm_percent Histogram2dcontourHistnorm = "percent" // percent
    Histogram2dcontourHistnorm_probability Histogram2dcontourHistnorm = "probability" // probability
    Histogram2dcontourHistnorm_density Histogram2dcontourHistnorm = "density" // density
    Histogram2dcontourHistnorm_probabilitydensity Histogram2dcontourHistnorm = "probability density" // probability density
)

// Histogram2dcontourHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type Histogram2dcontourHoverlabelAlign string

const (
    Histogram2dcontourHoverlabelAlign_left Histogram2dcontourHoverlabelAlign = "left" // left
    Histogram2dcontourHoverlabelAlign_right Histogram2dcontourHoverlabelAlign = "right" // right
    Histogram2dcontourHoverlabelAlign_auto Histogram2dcontourHoverlabelAlign = "auto" // auto
)

// Histogram2dcontourVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type Histogram2dcontourVisible interface{}

const (
    Histogram2dcontourVisibleTrue bool = true
    Histogram2dcontourVisibleFalse bool = false
    Histogram2dcontourVisibleLegendonly string = "legendonly"
)

// Histogram2dcontourXcalendar Sets the calendar system to use with `x` date data.
type Histogram2dcontourXcalendar string

const (
    Histogram2dcontourXcalendar_gregorian Histogram2dcontourXcalendar = "gregorian" // gregorian
    Histogram2dcontourXcalendar_chinese Histogram2dcontourXcalendar = "chinese" // chinese
    Histogram2dcontourXcalendar_coptic Histogram2dcontourXcalendar = "coptic" // coptic
    Histogram2dcontourXcalendar_discworld Histogram2dcontourXcalendar = "discworld" // discworld
    Histogram2dcontourXcalendar_ethiopian Histogram2dcontourXcalendar = "ethiopian" // ethiopian
    Histogram2dcontourXcalendar_hebrew Histogram2dcontourXcalendar = "hebrew" // hebrew
    Histogram2dcontourXcalendar_islamic Histogram2dcontourXcalendar = "islamic" // islamic
    Histogram2dcontourXcalendar_julian Histogram2dcontourXcalendar = "julian" // julian
    Histogram2dcontourXcalendar_mayan Histogram2dcontourXcalendar = "mayan" // mayan
    Histogram2dcontourXcalendar_nanakshahi Histogram2dcontourXcalendar = "nanakshahi" // nanakshahi
    Histogram2dcontourXcalendar_nepali Histogram2dcontourXcalendar = "nepali" // nepali
    Histogram2dcontourXcalendar_persian Histogram2dcontourXcalendar = "persian" // persian
    Histogram2dcontourXcalendar_jalali Histogram2dcontourXcalendar = "jalali" // jalali
    Histogram2dcontourXcalendar_taiwan Histogram2dcontourXcalendar = "taiwan" // taiwan
    Histogram2dcontourXcalendar_thai Histogram2dcontourXcalendar = "thai" // thai
    Histogram2dcontourXcalendar_ummalqura Histogram2dcontourXcalendar = "ummalqura" // ummalqura
)

// Histogram2dcontourYcalendar Sets the calendar system to use with `y` date data.
type Histogram2dcontourYcalendar string

const (
    Histogram2dcontourYcalendar_gregorian Histogram2dcontourYcalendar = "gregorian" // gregorian
    Histogram2dcontourYcalendar_chinese Histogram2dcontourYcalendar = "chinese" // chinese
    Histogram2dcontourYcalendar_coptic Histogram2dcontourYcalendar = "coptic" // coptic
    Histogram2dcontourYcalendar_discworld Histogram2dcontourYcalendar = "discworld" // discworld
    Histogram2dcontourYcalendar_ethiopian Histogram2dcontourYcalendar = "ethiopian" // ethiopian
    Histogram2dcontourYcalendar_hebrew Histogram2dcontourYcalendar = "hebrew" // hebrew
    Histogram2dcontourYcalendar_islamic Histogram2dcontourYcalendar = "islamic" // islamic
    Histogram2dcontourYcalendar_julian Histogram2dcontourYcalendar = "julian" // julian
    Histogram2dcontourYcalendar_mayan Histogram2dcontourYcalendar = "mayan" // mayan
    Histogram2dcontourYcalendar_nanakshahi Histogram2dcontourYcalendar = "nanakshahi" // nanakshahi
    Histogram2dcontourYcalendar_nepali Histogram2dcontourYcalendar = "nepali" // nepali
    Histogram2dcontourYcalendar_persian Histogram2dcontourYcalendar = "persian" // persian
    Histogram2dcontourYcalendar_jalali Histogram2dcontourYcalendar = "jalali" // jalali
    Histogram2dcontourYcalendar_taiwan Histogram2dcontourYcalendar = "taiwan" // taiwan
    Histogram2dcontourYcalendar_thai Histogram2dcontourYcalendar = "thai" // thai
    Histogram2dcontourYcalendar_ummalqura Histogram2dcontourYcalendar = "ummalqura" // ummalqura
)

// HistogramCumulativeCurrentbin Only applies if cumulative is enabled. Sets whether the current bin is included, excluded, or has half of its value included in the current cumulative value. *include* is the default for compatibility with various other tools, however it introduces a half-bin bias to the results. *exclude* makes the opposite half-bin bias, and *half* removes it.
type HistogramCumulativeCurrentbin string

const (
    HistogramCumulativeCurrentbin_include HistogramCumulativeCurrentbin = "include" // include
    HistogramCumulativeCurrentbin_exclude HistogramCumulativeCurrentbin = "exclude" // exclude
    HistogramCumulativeCurrentbin_half HistogramCumulativeCurrentbin = "half" // half
)

// HistogramCumulativeDirection Only applies if cumulative is enabled. If *increasing* (default) we sum all prior bins, so the result increases from left to right. If *decreasing* we sum later bins so the result decreases from left to right.
type HistogramCumulativeDirection string

const (
    HistogramCumulativeDirection_increasing HistogramCumulativeDirection = "increasing" // increasing
    HistogramCumulativeDirection_decreasing HistogramCumulativeDirection = "decreasing" // decreasing
)

// HistogramErrorXType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type HistogramErrorXType string

const (
    HistogramErrorXType_percent HistogramErrorXType = "percent" // percent
    HistogramErrorXType_constant HistogramErrorXType = "constant" // constant
    HistogramErrorXType_sqrt HistogramErrorXType = "sqrt" // sqrt
    HistogramErrorXType_data HistogramErrorXType = "data" // data
)

// HistogramErrorYType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type HistogramErrorYType string

const (
    HistogramErrorYType_percent HistogramErrorYType = "percent" // percent
    HistogramErrorYType_constant HistogramErrorYType = "constant" // constant
    HistogramErrorYType_sqrt HistogramErrorYType = "sqrt" // sqrt
    HistogramErrorYType_data HistogramErrorYType = "data" // data
)

// HistogramHistfunc Specifies the binning function used for this histogram trace. If *count*, the histogram values are computed by counting the number of values lying inside each bin. If *sum*, *avg*, *min*, *max*, the histogram values are computed using the sum, the average, the minimum or the maximum of the values lying inside each bin respectively.
type HistogramHistfunc string

const (
    HistogramHistfunc_count HistogramHistfunc = "count" // count
    HistogramHistfunc_sum HistogramHistfunc = "sum" // sum
    HistogramHistfunc_avg HistogramHistfunc = "avg" // avg
    HistogramHistfunc_min HistogramHistfunc = "min" // min
    HistogramHistfunc_max HistogramHistfunc = "max" // max
)

// HistogramHistnorm Specifies the type of normalization used for this histogram trace. If **, the span of each bar corresponds to the number of occurrences (i.e. the number of data points lying inside the bins). If *percent* / *probability*, the span of each bar corresponds to the percentage / fraction of occurrences with respect to the total number of sample points (here, the sum of all bin HEIGHTS equals 100% / 1). If *density*, the span of each bar corresponds to the number of occurrences in a bin divided by the size of the bin interval (here, the sum of all bin AREAS equals the total number of sample points). If *probability density*, the area of each bar corresponds to the probability that an event will fall into the corresponding bin (here, the sum of all bin AREAS equals 1).
type HistogramHistnorm string

const (
    HistogramHistnorm_percent HistogramHistnorm = "percent" // percent
    HistogramHistnorm_probability HistogramHistnorm = "probability" // probability
    HistogramHistnorm_density HistogramHistnorm = "density" // density
    HistogramHistnorm_probabilitydensity HistogramHistnorm = "probability density" // probability density
)

// HistogramHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type HistogramHoverlabelAlign string

const (
    HistogramHoverlabelAlign_left HistogramHoverlabelAlign = "left" // left
    HistogramHoverlabelAlign_right HistogramHoverlabelAlign = "right" // right
    HistogramHoverlabelAlign_auto HistogramHoverlabelAlign = "auto" // auto
)

// HistogramMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type HistogramMarkerColorbarExponentformat string

const (
    HistogramMarkerColorbarExponentformat_none HistogramMarkerColorbarExponentformat = "none" // none
    HistogramMarkerColorbarExponentformat_e HistogramMarkerColorbarExponentformat = "e" // e
    HistogramMarkerColorbarExponentformat_E HistogramMarkerColorbarExponentformat = "E" // E
    HistogramMarkerColorbarExponentformat_power HistogramMarkerColorbarExponentformat = "power" // power
    HistogramMarkerColorbarExponentformat_SI HistogramMarkerColorbarExponentformat = "SI" // SI
    HistogramMarkerColorbarExponentformat_B HistogramMarkerColorbarExponentformat = "B" // B
)

// HistogramMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type HistogramMarkerColorbarLenmode string

const (
    HistogramMarkerColorbarLenmode_fraction HistogramMarkerColorbarLenmode = "fraction" // fraction
    HistogramMarkerColorbarLenmode_pixels HistogramMarkerColorbarLenmode = "pixels" // pixels
)

// HistogramMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type HistogramMarkerColorbarShowexponent string

const (
    HistogramMarkerColorbarShowexponent_all HistogramMarkerColorbarShowexponent = "all" // all
    HistogramMarkerColorbarShowexponent_first HistogramMarkerColorbarShowexponent = "first" // first
    HistogramMarkerColorbarShowexponent_last HistogramMarkerColorbarShowexponent = "last" // last
    HistogramMarkerColorbarShowexponent_none HistogramMarkerColorbarShowexponent = "none" // none
)

// HistogramMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type HistogramMarkerColorbarShowtickprefix string

const (
    HistogramMarkerColorbarShowtickprefix_all HistogramMarkerColorbarShowtickprefix = "all" // all
    HistogramMarkerColorbarShowtickprefix_first HistogramMarkerColorbarShowtickprefix = "first" // first
    HistogramMarkerColorbarShowtickprefix_last HistogramMarkerColorbarShowtickprefix = "last" // last
    HistogramMarkerColorbarShowtickprefix_none HistogramMarkerColorbarShowtickprefix = "none" // none
)

// HistogramMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type HistogramMarkerColorbarShowticksuffix string

const (
    HistogramMarkerColorbarShowticksuffix_all HistogramMarkerColorbarShowticksuffix = "all" // all
    HistogramMarkerColorbarShowticksuffix_first HistogramMarkerColorbarShowticksuffix = "first" // first
    HistogramMarkerColorbarShowticksuffix_last HistogramMarkerColorbarShowticksuffix = "last" // last
    HistogramMarkerColorbarShowticksuffix_none HistogramMarkerColorbarShowticksuffix = "none" // none
)

// HistogramMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type HistogramMarkerColorbarThicknessmode string

const (
    HistogramMarkerColorbarThicknessmode_fraction HistogramMarkerColorbarThicknessmode = "fraction" // fraction
    HistogramMarkerColorbarThicknessmode_pixels HistogramMarkerColorbarThicknessmode = "pixels" // pixels
)

// HistogramMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type HistogramMarkerColorbarTickmode string

const (
    HistogramMarkerColorbarTickmode_auto HistogramMarkerColorbarTickmode = "auto" // auto
    HistogramMarkerColorbarTickmode_linear HistogramMarkerColorbarTickmode = "linear" // linear
    HistogramMarkerColorbarTickmode_array HistogramMarkerColorbarTickmode = "array" // array
)

// HistogramMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type HistogramMarkerColorbarTicks string

const (
    HistogramMarkerColorbarTicks_outside HistogramMarkerColorbarTicks = "outside" // outside
    HistogramMarkerColorbarTicks_inside HistogramMarkerColorbarTicks = "inside" // inside
)

// HistogramMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type HistogramMarkerColorbarTitleSide string

const (
    HistogramMarkerColorbarTitleSide_right HistogramMarkerColorbarTitleSide = "right" // right
    HistogramMarkerColorbarTitleSide_top HistogramMarkerColorbarTitleSide = "top" // top
    HistogramMarkerColorbarTitleSide_bottom HistogramMarkerColorbarTitleSide = "bottom" // bottom
)

// HistogramMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type HistogramMarkerColorbarXanchor string

const (
    HistogramMarkerColorbarXanchor_left HistogramMarkerColorbarXanchor = "left" // left
    HistogramMarkerColorbarXanchor_center HistogramMarkerColorbarXanchor = "center" // center
    HistogramMarkerColorbarXanchor_right HistogramMarkerColorbarXanchor = "right" // right
)

// HistogramMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type HistogramMarkerColorbarYanchor string

const (
    HistogramMarkerColorbarYanchor_top HistogramMarkerColorbarYanchor = "top" // top
    HistogramMarkerColorbarYanchor_middle HistogramMarkerColorbarYanchor = "middle" // middle
    HistogramMarkerColorbarYanchor_bottom HistogramMarkerColorbarYanchor = "bottom" // bottom
)

// HistogramOrientation Sets the orientation of the bars. With *v* (*h*), the value of the each bar spans along the vertical (horizontal).
type HistogramOrientation string

const (
    HistogramOrientation_v HistogramOrientation = "v" // v
    HistogramOrientation_h HistogramOrientation = "h" // h
)

// HistogramVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type HistogramVisible interface{}

const (
    HistogramVisibleTrue bool = true
    HistogramVisibleFalse bool = false
    HistogramVisibleLegendonly string = "legendonly"
)

// HistogramXcalendar Sets the calendar system to use with `x` date data.
type HistogramXcalendar string

const (
    HistogramXcalendar_gregorian HistogramXcalendar = "gregorian" // gregorian
    HistogramXcalendar_chinese HistogramXcalendar = "chinese" // chinese
    HistogramXcalendar_coptic HistogramXcalendar = "coptic" // coptic
    HistogramXcalendar_discworld HistogramXcalendar = "discworld" // discworld
    HistogramXcalendar_ethiopian HistogramXcalendar = "ethiopian" // ethiopian
    HistogramXcalendar_hebrew HistogramXcalendar = "hebrew" // hebrew
    HistogramXcalendar_islamic HistogramXcalendar = "islamic" // islamic
    HistogramXcalendar_julian HistogramXcalendar = "julian" // julian
    HistogramXcalendar_mayan HistogramXcalendar = "mayan" // mayan
    HistogramXcalendar_nanakshahi HistogramXcalendar = "nanakshahi" // nanakshahi
    HistogramXcalendar_nepali HistogramXcalendar = "nepali" // nepali
    HistogramXcalendar_persian HistogramXcalendar = "persian" // persian
    HistogramXcalendar_jalali HistogramXcalendar = "jalali" // jalali
    HistogramXcalendar_taiwan HistogramXcalendar = "taiwan" // taiwan
    HistogramXcalendar_thai HistogramXcalendar = "thai" // thai
    HistogramXcalendar_ummalqura HistogramXcalendar = "ummalqura" // ummalqura
)

// HistogramYcalendar Sets the calendar system to use with `y` date data.
type HistogramYcalendar string

const (
    HistogramYcalendar_gregorian HistogramYcalendar = "gregorian" // gregorian
    HistogramYcalendar_chinese HistogramYcalendar = "chinese" // chinese
    HistogramYcalendar_coptic HistogramYcalendar = "coptic" // coptic
    HistogramYcalendar_discworld HistogramYcalendar = "discworld" // discworld
    HistogramYcalendar_ethiopian HistogramYcalendar = "ethiopian" // ethiopian
    HistogramYcalendar_hebrew HistogramYcalendar = "hebrew" // hebrew
    HistogramYcalendar_islamic HistogramYcalendar = "islamic" // islamic
    HistogramYcalendar_julian HistogramYcalendar = "julian" // julian
    HistogramYcalendar_mayan HistogramYcalendar = "mayan" // mayan
    HistogramYcalendar_nanakshahi HistogramYcalendar = "nanakshahi" // nanakshahi
    HistogramYcalendar_nepali HistogramYcalendar = "nepali" // nepali
    HistogramYcalendar_persian HistogramYcalendar = "persian" // persian
    HistogramYcalendar_jalali HistogramYcalendar = "jalali" // jalali
    HistogramYcalendar_taiwan HistogramYcalendar = "taiwan" // taiwan
    HistogramYcalendar_thai HistogramYcalendar = "thai" // thai
    HistogramYcalendar_ummalqura HistogramYcalendar = "ummalqura" // ummalqura
)

// ImageColormodel Color model used to map the numerical color components described in `z` into colors.
type ImageColormodel string

const (
    ImageColormodel_rgb ImageColormodel = "rgb" // rgb
    ImageColormodel_rgba ImageColormodel = "rgba" // rgba
    ImageColormodel_hsl ImageColormodel = "hsl" // hsl
    ImageColormodel_hsla ImageColormodel = "hsla" // hsla
)

// ImageHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ImageHoverlabelAlign string

const (
    ImageHoverlabelAlign_left ImageHoverlabelAlign = "left" // left
    ImageHoverlabelAlign_right ImageHoverlabelAlign = "right" // right
    ImageHoverlabelAlign_auto ImageHoverlabelAlign = "auto" // auto
)

// ImageVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ImageVisible interface{}

const (
    ImageVisibleTrue bool = true
    ImageVisibleFalse bool = false
    ImageVisibleLegendonly string = "legendonly"
)

// IndicatorAlign Sets the horizontal alignment of the `text` within the box. Note that this attribute has no effect if an angular gauge is displayed: in this case, it is always centered
type IndicatorAlign string

const (
    IndicatorAlign_left IndicatorAlign = "left" // left
    IndicatorAlign_center IndicatorAlign = "center" // center
    IndicatorAlign_right IndicatorAlign = "right" // right
)

// IndicatorDeltaPosition Sets the position of delta with respect to the number.
type IndicatorDeltaPosition string

const (
    IndicatorDeltaPosition_top IndicatorDeltaPosition = "top" // top
    IndicatorDeltaPosition_bottom IndicatorDeltaPosition = "bottom" // bottom
    IndicatorDeltaPosition_left IndicatorDeltaPosition = "left" // left
    IndicatorDeltaPosition_right IndicatorDeltaPosition = "right" // right
)

// IndicatorGaugeAxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type IndicatorGaugeAxisExponentformat string

const (
    IndicatorGaugeAxisExponentformat_none IndicatorGaugeAxisExponentformat = "none" // none
    IndicatorGaugeAxisExponentformat_e IndicatorGaugeAxisExponentformat = "e" // e
    IndicatorGaugeAxisExponentformat_E IndicatorGaugeAxisExponentformat = "E" // E
    IndicatorGaugeAxisExponentformat_power IndicatorGaugeAxisExponentformat = "power" // power
    IndicatorGaugeAxisExponentformat_SI IndicatorGaugeAxisExponentformat = "SI" // SI
    IndicatorGaugeAxisExponentformat_B IndicatorGaugeAxisExponentformat = "B" // B
)

// IndicatorGaugeAxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type IndicatorGaugeAxisShowexponent string

const (
    IndicatorGaugeAxisShowexponent_all IndicatorGaugeAxisShowexponent = "all" // all
    IndicatorGaugeAxisShowexponent_first IndicatorGaugeAxisShowexponent = "first" // first
    IndicatorGaugeAxisShowexponent_last IndicatorGaugeAxisShowexponent = "last" // last
    IndicatorGaugeAxisShowexponent_none IndicatorGaugeAxisShowexponent = "none" // none
)

// IndicatorGaugeAxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type IndicatorGaugeAxisShowtickprefix string

const (
    IndicatorGaugeAxisShowtickprefix_all IndicatorGaugeAxisShowtickprefix = "all" // all
    IndicatorGaugeAxisShowtickprefix_first IndicatorGaugeAxisShowtickprefix = "first" // first
    IndicatorGaugeAxisShowtickprefix_last IndicatorGaugeAxisShowtickprefix = "last" // last
    IndicatorGaugeAxisShowtickprefix_none IndicatorGaugeAxisShowtickprefix = "none" // none
)

// IndicatorGaugeAxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type IndicatorGaugeAxisShowticksuffix string

const (
    IndicatorGaugeAxisShowticksuffix_all IndicatorGaugeAxisShowticksuffix = "all" // all
    IndicatorGaugeAxisShowticksuffix_first IndicatorGaugeAxisShowticksuffix = "first" // first
    IndicatorGaugeAxisShowticksuffix_last IndicatorGaugeAxisShowticksuffix = "last" // last
    IndicatorGaugeAxisShowticksuffix_none IndicatorGaugeAxisShowticksuffix = "none" // none
)

// IndicatorGaugeAxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type IndicatorGaugeAxisTickmode string

const (
    IndicatorGaugeAxisTickmode_auto IndicatorGaugeAxisTickmode = "auto" // auto
    IndicatorGaugeAxisTickmode_linear IndicatorGaugeAxisTickmode = "linear" // linear
    IndicatorGaugeAxisTickmode_array IndicatorGaugeAxisTickmode = "array" // array
)

// IndicatorGaugeAxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type IndicatorGaugeAxisTicks string

const (
    IndicatorGaugeAxisTicks_outside IndicatorGaugeAxisTicks = "outside" // outside
    IndicatorGaugeAxisTicks_inside IndicatorGaugeAxisTicks = "inside" // inside
)

// IndicatorGaugeShape Set the shape of the gauge
type IndicatorGaugeShape string

const (
    IndicatorGaugeShape_angular IndicatorGaugeShape = "angular" // angular
    IndicatorGaugeShape_bullet IndicatorGaugeShape = "bullet" // bullet
)

// IndicatorTitleAlign Sets the horizontal alignment of the title. It defaults to `center` except for bullet charts for which it defaults to right.
type IndicatorTitleAlign string

const (
    IndicatorTitleAlign_left IndicatorTitleAlign = "left" // left
    IndicatorTitleAlign_center IndicatorTitleAlign = "center" // center
    IndicatorTitleAlign_right IndicatorTitleAlign = "right" // right
)

// IndicatorVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type IndicatorVisible interface{}

const (
    IndicatorVisibleTrue bool = true
    IndicatorVisibleFalse bool = false
    IndicatorVisibleLegendonly string = "legendonly"
)

// IsosurfaceColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type IsosurfaceColorbarExponentformat string

const (
    IsosurfaceColorbarExponentformat_none IsosurfaceColorbarExponentformat = "none" // none
    IsosurfaceColorbarExponentformat_e IsosurfaceColorbarExponentformat = "e" // e
    IsosurfaceColorbarExponentformat_E IsosurfaceColorbarExponentformat = "E" // E
    IsosurfaceColorbarExponentformat_power IsosurfaceColorbarExponentformat = "power" // power
    IsosurfaceColorbarExponentformat_SI IsosurfaceColorbarExponentformat = "SI" // SI
    IsosurfaceColorbarExponentformat_B IsosurfaceColorbarExponentformat = "B" // B
)

// IsosurfaceColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type IsosurfaceColorbarLenmode string

const (
    IsosurfaceColorbarLenmode_fraction IsosurfaceColorbarLenmode = "fraction" // fraction
    IsosurfaceColorbarLenmode_pixels IsosurfaceColorbarLenmode = "pixels" // pixels
)

// IsosurfaceColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type IsosurfaceColorbarShowexponent string

const (
    IsosurfaceColorbarShowexponent_all IsosurfaceColorbarShowexponent = "all" // all
    IsosurfaceColorbarShowexponent_first IsosurfaceColorbarShowexponent = "first" // first
    IsosurfaceColorbarShowexponent_last IsosurfaceColorbarShowexponent = "last" // last
    IsosurfaceColorbarShowexponent_none IsosurfaceColorbarShowexponent = "none" // none
)

// IsosurfaceColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type IsosurfaceColorbarShowtickprefix string

const (
    IsosurfaceColorbarShowtickprefix_all IsosurfaceColorbarShowtickprefix = "all" // all
    IsosurfaceColorbarShowtickprefix_first IsosurfaceColorbarShowtickprefix = "first" // first
    IsosurfaceColorbarShowtickprefix_last IsosurfaceColorbarShowtickprefix = "last" // last
    IsosurfaceColorbarShowtickprefix_none IsosurfaceColorbarShowtickprefix = "none" // none
)

// IsosurfaceColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type IsosurfaceColorbarShowticksuffix string

const (
    IsosurfaceColorbarShowticksuffix_all IsosurfaceColorbarShowticksuffix = "all" // all
    IsosurfaceColorbarShowticksuffix_first IsosurfaceColorbarShowticksuffix = "first" // first
    IsosurfaceColorbarShowticksuffix_last IsosurfaceColorbarShowticksuffix = "last" // last
    IsosurfaceColorbarShowticksuffix_none IsosurfaceColorbarShowticksuffix = "none" // none
)

// IsosurfaceColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type IsosurfaceColorbarThicknessmode string

const (
    IsosurfaceColorbarThicknessmode_fraction IsosurfaceColorbarThicknessmode = "fraction" // fraction
    IsosurfaceColorbarThicknessmode_pixels IsosurfaceColorbarThicknessmode = "pixels" // pixels
)

// IsosurfaceColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type IsosurfaceColorbarTickmode string

const (
    IsosurfaceColorbarTickmode_auto IsosurfaceColorbarTickmode = "auto" // auto
    IsosurfaceColorbarTickmode_linear IsosurfaceColorbarTickmode = "linear" // linear
    IsosurfaceColorbarTickmode_array IsosurfaceColorbarTickmode = "array" // array
)

// IsosurfaceColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type IsosurfaceColorbarTicks string

const (
    IsosurfaceColorbarTicks_outside IsosurfaceColorbarTicks = "outside" // outside
    IsosurfaceColorbarTicks_inside IsosurfaceColorbarTicks = "inside" // inside
)

// IsosurfaceColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type IsosurfaceColorbarTitleSide string

const (
    IsosurfaceColorbarTitleSide_right IsosurfaceColorbarTitleSide = "right" // right
    IsosurfaceColorbarTitleSide_top IsosurfaceColorbarTitleSide = "top" // top
    IsosurfaceColorbarTitleSide_bottom IsosurfaceColorbarTitleSide = "bottom" // bottom
)

// IsosurfaceColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type IsosurfaceColorbarXanchor string

const (
    IsosurfaceColorbarXanchor_left IsosurfaceColorbarXanchor = "left" // left
    IsosurfaceColorbarXanchor_center IsosurfaceColorbarXanchor = "center" // center
    IsosurfaceColorbarXanchor_right IsosurfaceColorbarXanchor = "right" // right
)

// IsosurfaceColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type IsosurfaceColorbarYanchor string

const (
    IsosurfaceColorbarYanchor_top IsosurfaceColorbarYanchor = "top" // top
    IsosurfaceColorbarYanchor_middle IsosurfaceColorbarYanchor = "middle" // middle
    IsosurfaceColorbarYanchor_bottom IsosurfaceColorbarYanchor = "bottom" // bottom
)

// IsosurfaceHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type IsosurfaceHoverlabelAlign string

const (
    IsosurfaceHoverlabelAlign_left IsosurfaceHoverlabelAlign = "left" // left
    IsosurfaceHoverlabelAlign_right IsosurfaceHoverlabelAlign = "right" // right
    IsosurfaceHoverlabelAlign_auto IsosurfaceHoverlabelAlign = "auto" // auto
)

// IsosurfaceVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type IsosurfaceVisible interface{}

const (
    IsosurfaceVisibleTrue bool = true
    IsosurfaceVisibleFalse bool = false
    IsosurfaceVisibleLegendonly string = "legendonly"
)

// LayoutAngularaxisTickorientation Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the orientation (from the paper perspective) of the angular axis tick labels.
type LayoutAngularaxisTickorientation string

const (
    LayoutAngularaxisTickorientation_horizontal LayoutAngularaxisTickorientation = "horizontal" // horizontal
    LayoutAngularaxisTickorientation_vertical LayoutAngularaxisTickorientation = "vertical" // vertical
)

// LayoutCalendar Sets the default calendar system to use for interpreting and displaying dates throughout the plot.
type LayoutCalendar string

const (
    LayoutCalendar_gregorian LayoutCalendar = "gregorian" // gregorian
    LayoutCalendar_chinese LayoutCalendar = "chinese" // chinese
    LayoutCalendar_coptic LayoutCalendar = "coptic" // coptic
    LayoutCalendar_discworld LayoutCalendar = "discworld" // discworld
    LayoutCalendar_ethiopian LayoutCalendar = "ethiopian" // ethiopian
    LayoutCalendar_hebrew LayoutCalendar = "hebrew" // hebrew
    LayoutCalendar_islamic LayoutCalendar = "islamic" // islamic
    LayoutCalendar_julian LayoutCalendar = "julian" // julian
    LayoutCalendar_mayan LayoutCalendar = "mayan" // mayan
    LayoutCalendar_nanakshahi LayoutCalendar = "nanakshahi" // nanakshahi
    LayoutCalendar_nepali LayoutCalendar = "nepali" // nepali
    LayoutCalendar_persian LayoutCalendar = "persian" // persian
    LayoutCalendar_jalali LayoutCalendar = "jalali" // jalali
    LayoutCalendar_taiwan LayoutCalendar = "taiwan" // taiwan
    LayoutCalendar_thai LayoutCalendar = "thai" // thai
    LayoutCalendar_ummalqura LayoutCalendar = "ummalqura" // ummalqura
)

// LayoutColoraxisColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutColoraxisColorbarExponentformat string

const (
    LayoutColoraxisColorbarExponentformat_none LayoutColoraxisColorbarExponentformat = "none" // none
    LayoutColoraxisColorbarExponentformat_e LayoutColoraxisColorbarExponentformat = "e" // e
    LayoutColoraxisColorbarExponentformat_E LayoutColoraxisColorbarExponentformat = "E" // E
    LayoutColoraxisColorbarExponentformat_power LayoutColoraxisColorbarExponentformat = "power" // power
    LayoutColoraxisColorbarExponentformat_SI LayoutColoraxisColorbarExponentformat = "SI" // SI
    LayoutColoraxisColorbarExponentformat_B LayoutColoraxisColorbarExponentformat = "B" // B
)

// LayoutColoraxisColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type LayoutColoraxisColorbarLenmode string

const (
    LayoutColoraxisColorbarLenmode_fraction LayoutColoraxisColorbarLenmode = "fraction" // fraction
    LayoutColoraxisColorbarLenmode_pixels LayoutColoraxisColorbarLenmode = "pixels" // pixels
)

// LayoutColoraxisColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutColoraxisColorbarShowexponent string

const (
    LayoutColoraxisColorbarShowexponent_all LayoutColoraxisColorbarShowexponent = "all" // all
    LayoutColoraxisColorbarShowexponent_first LayoutColoraxisColorbarShowexponent = "first" // first
    LayoutColoraxisColorbarShowexponent_last LayoutColoraxisColorbarShowexponent = "last" // last
    LayoutColoraxisColorbarShowexponent_none LayoutColoraxisColorbarShowexponent = "none" // none
)

// LayoutColoraxisColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutColoraxisColorbarShowtickprefix string

const (
    LayoutColoraxisColorbarShowtickprefix_all LayoutColoraxisColorbarShowtickprefix = "all" // all
    LayoutColoraxisColorbarShowtickprefix_first LayoutColoraxisColorbarShowtickprefix = "first" // first
    LayoutColoraxisColorbarShowtickprefix_last LayoutColoraxisColorbarShowtickprefix = "last" // last
    LayoutColoraxisColorbarShowtickprefix_none LayoutColoraxisColorbarShowtickprefix = "none" // none
)

// LayoutColoraxisColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutColoraxisColorbarShowticksuffix string

const (
    LayoutColoraxisColorbarShowticksuffix_all LayoutColoraxisColorbarShowticksuffix = "all" // all
    LayoutColoraxisColorbarShowticksuffix_first LayoutColoraxisColorbarShowticksuffix = "first" // first
    LayoutColoraxisColorbarShowticksuffix_last LayoutColoraxisColorbarShowticksuffix = "last" // last
    LayoutColoraxisColorbarShowticksuffix_none LayoutColoraxisColorbarShowticksuffix = "none" // none
)

// LayoutColoraxisColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type LayoutColoraxisColorbarThicknessmode string

const (
    LayoutColoraxisColorbarThicknessmode_fraction LayoutColoraxisColorbarThicknessmode = "fraction" // fraction
    LayoutColoraxisColorbarThicknessmode_pixels LayoutColoraxisColorbarThicknessmode = "pixels" // pixels
)

// LayoutColoraxisColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutColoraxisColorbarTickmode string

const (
    LayoutColoraxisColorbarTickmode_auto LayoutColoraxisColorbarTickmode = "auto" // auto
    LayoutColoraxisColorbarTickmode_linear LayoutColoraxisColorbarTickmode = "linear" // linear
    LayoutColoraxisColorbarTickmode_array LayoutColoraxisColorbarTickmode = "array" // array
)

// LayoutColoraxisColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutColoraxisColorbarTicks string

const (
    LayoutColoraxisColorbarTicks_outside LayoutColoraxisColorbarTicks = "outside" // outside
    LayoutColoraxisColorbarTicks_inside LayoutColoraxisColorbarTicks = "inside" // inside
)

// LayoutColoraxisColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type LayoutColoraxisColorbarTitleSide string

const (
    LayoutColoraxisColorbarTitleSide_right LayoutColoraxisColorbarTitleSide = "right" // right
    LayoutColoraxisColorbarTitleSide_top LayoutColoraxisColorbarTitleSide = "top" // top
    LayoutColoraxisColorbarTitleSide_bottom LayoutColoraxisColorbarTitleSide = "bottom" // bottom
)

// LayoutColoraxisColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type LayoutColoraxisColorbarXanchor string

const (
    LayoutColoraxisColorbarXanchor_left LayoutColoraxisColorbarXanchor = "left" // left
    LayoutColoraxisColorbarXanchor_center LayoutColoraxisColorbarXanchor = "center" // center
    LayoutColoraxisColorbarXanchor_right LayoutColoraxisColorbarXanchor = "right" // right
)

// LayoutColoraxisColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type LayoutColoraxisColorbarYanchor string

const (
    LayoutColoraxisColorbarYanchor_top LayoutColoraxisColorbarYanchor = "top" // top
    LayoutColoraxisColorbarYanchor_middle LayoutColoraxisColorbarYanchor = "middle" // middle
    LayoutColoraxisColorbarYanchor_bottom LayoutColoraxisColorbarYanchor = "bottom" // bottom
)

// LayoutDirection Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the direction corresponding to positive angles in legacy polar charts.
type LayoutDirection string

const (
    LayoutDirection_clockwise LayoutDirection = "clockwise" // clockwise
    LayoutDirection_counterclockwise LayoutDirection = "counterclockwise" // counterclockwise
)

// LayoutDragmode Determines the mode of drag interactions. *select* and *lasso* apply only to scatter traces with markers or text. *orbit* and *turntable* apply only to 3D scenes.
type LayoutDragmode interface{}

const (
    LayoutDragmodeZoom string = "zoom"
    LayoutDragmodePan string = "pan"
    LayoutDragmodeSelect string = "select"
    LayoutDragmodeLasso string = "lasso"
    LayoutDragmodeDrawclosedpath string = "drawclosedpath"
    LayoutDragmodeDrawopenpath string = "drawopenpath"
    LayoutDragmodeDrawline string = "drawline"
    LayoutDragmodeDrawrect string = "drawrect"
    LayoutDragmodeDrawcircle string = "drawcircle"
    LayoutDragmodeOrbit string = "orbit"
    LayoutDragmodeTurntable string = "turntable"
    LayoutDragmodeFalse bool = false
)

// LayoutGeoFitbounds Determines if this subplot's view settings are auto-computed to fit trace data. On scoped maps, setting `fitbounds` leads to `center.lon` and `center.lat` getting auto-filled. On maps with a non-clipped projection, setting `fitbounds` leads to `center.lon`, `center.lat`, and `projection.rotation.lon` getting auto-filled. On maps with a clipped projection, setting `fitbounds` leads to `center.lon`, `center.lat`, `projection.rotation.lon`, `projection.rotation.lat`, `lonaxis.range` and `lonaxis.range` getting auto-filled. If *locations*, only the trace's visible locations are considered in the `fitbounds` computations. If *geojson*, the entire trace input `geojson` (if provided) is considered in the `fitbounds` computations, Defaults to *false*.
type LayoutGeoFitbounds interface{}

const (
    LayoutGeoFitboundsFalse bool = false
    LayoutGeoFitboundsLocations string = "locations"
    LayoutGeoFitboundsGeojson string = "geojson"
)

// LayoutGeoProjectionType Sets the projection type.
type LayoutGeoProjectionType string

const (
    LayoutGeoProjectionType_equirectangular LayoutGeoProjectionType = "equirectangular" // equirectangular
    LayoutGeoProjectionType_mercator LayoutGeoProjectionType = "mercator" // mercator
    LayoutGeoProjectionType_orthographic LayoutGeoProjectionType = "orthographic" // orthographic
    LayoutGeoProjectionType_naturalearth LayoutGeoProjectionType = "natural earth" // natural earth
    LayoutGeoProjectionType_kavrayskiy7 LayoutGeoProjectionType = "kavrayskiy7" // kavrayskiy7
    LayoutGeoProjectionType_miller LayoutGeoProjectionType = "miller" // miller
    LayoutGeoProjectionType_robinson LayoutGeoProjectionType = "robinson" // robinson
    LayoutGeoProjectionType_eckert4 LayoutGeoProjectionType = "eckert4" // eckert4
    LayoutGeoProjectionType_azimuthalequalarea LayoutGeoProjectionType = "azimuthal equal area" // azimuthal equal area
    LayoutGeoProjectionType_azimuthalequidistant LayoutGeoProjectionType = "azimuthal equidistant" // azimuthal equidistant
    LayoutGeoProjectionType_conicequalarea LayoutGeoProjectionType = "conic equal area" // conic equal area
    LayoutGeoProjectionType_conicconformal LayoutGeoProjectionType = "conic conformal" // conic conformal
    LayoutGeoProjectionType_conicequidistant LayoutGeoProjectionType = "conic equidistant" // conic equidistant
    LayoutGeoProjectionType_gnomonic LayoutGeoProjectionType = "gnomonic" // gnomonic
    LayoutGeoProjectionType_stereographic LayoutGeoProjectionType = "stereographic" // stereographic
    LayoutGeoProjectionType_mollweide LayoutGeoProjectionType = "mollweide" // mollweide
    LayoutGeoProjectionType_hammer LayoutGeoProjectionType = "hammer" // hammer
    LayoutGeoProjectionType_transversemercator LayoutGeoProjectionType = "transverse mercator" // transverse mercator
    LayoutGeoProjectionType_albersusa LayoutGeoProjectionType = "albers usa" // albers usa
    LayoutGeoProjectionType_winkeltripel LayoutGeoProjectionType = "winkel tripel" // winkel tripel
    LayoutGeoProjectionType_aitoff LayoutGeoProjectionType = "aitoff" // aitoff
    LayoutGeoProjectionType_sinusoidal LayoutGeoProjectionType = "sinusoidal" // sinusoidal
)

// LayoutGeoResolution Sets the resolution of the base layers. The values have units of km/mm e.g. 110 corresponds to a scale ratio of 1:110,000,000.
type LayoutGeoResolution string

const (
    LayoutGeoResolution110 LayoutGeoResolution = "110" // 110
    LayoutGeoResolution50 LayoutGeoResolution = "50" // 50
)

// LayoutGeoScope Set the scope of the map.
type LayoutGeoScope string

const (
    LayoutGeoScope_world LayoutGeoScope = "world" // world
    LayoutGeoScope_usa LayoutGeoScope = "usa" // usa
    LayoutGeoScope_europe LayoutGeoScope = "europe" // europe
    LayoutGeoScope_asia LayoutGeoScope = "asia" // asia
    LayoutGeoScope_africa LayoutGeoScope = "africa" // africa
    LayoutGeoScope_northamerica LayoutGeoScope = "north america" // north america
    LayoutGeoScope_southamerica LayoutGeoScope = "south america" // south america
)

// LayoutGridPattern If no `subplots`, `xaxes`, or `yaxes` are given but we do have `rows` and `columns`, we can generate defaults using consecutive axis IDs, in two ways: *coupled* gives one x axis per column and one y axis per row. *independent* uses a new xy pair for each cell, left-to-right across each row then iterating rows according to `roworder`.
type LayoutGridPattern string

const (
    LayoutGridPattern_independent LayoutGridPattern = "independent" // independent
    LayoutGridPattern_coupled LayoutGridPattern = "coupled" // coupled
)

// LayoutGridRoworder Is the first row the top or the bottom? Note that columns are always enumerated from left to right.
type LayoutGridRoworder string

const (
    LayoutGridRoworder_toptobottom LayoutGridRoworder = "top to bottom" // top to bottom
    LayoutGridRoworder_bottomtotop LayoutGridRoworder = "bottom to top" // bottom to top
)

// LayoutGridXside Sets where the x axis labels and titles go. *bottom* means the very bottom of the grid. *bottom plot* is the lowest plot that each x axis is used in. *top* and *top plot* are similar.
type LayoutGridXside string

const (
    LayoutGridXside_bottom LayoutGridXside = "bottom" // bottom
    LayoutGridXside_bottomplot LayoutGridXside = "bottom plot" // bottom plot
    LayoutGridXside_topplot LayoutGridXside = "top plot" // top plot
    LayoutGridXside_top LayoutGridXside = "top" // top
)

// LayoutGridYside Sets where the y axis labels and titles go. *left* means the very left edge of the grid. *left plot* is the leftmost plot that each y axis is used in. *right* and *right plot* are similar.
type LayoutGridYside string

const (
    LayoutGridYside_left LayoutGridYside = "left" // left
    LayoutGridYside_leftplot LayoutGridYside = "left plot" // left plot
    LayoutGridYside_rightplot LayoutGridYside = "right plot" // right plot
    LayoutGridYside_right LayoutGridYside = "right" // right
)

// LayoutHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type LayoutHoverlabelAlign string

const (
    LayoutHoverlabelAlign_left LayoutHoverlabelAlign = "left" // left
    LayoutHoverlabelAlign_right LayoutHoverlabelAlign = "right" // right
    LayoutHoverlabelAlign_auto LayoutHoverlabelAlign = "auto" // auto
)

// LayoutHovermode Determines the mode of hover interactions. If *closest*, a single hoverlabel will appear for the *closest* point within the `hoverdistance`. If *x* (or *y*), multiple hoverlabels will appear for multiple points at the *closest* x- (or y-) coordinate within the `hoverdistance`, with the caveat that no more than one hoverlabel will appear per trace. If *x unified* (or *y unified*), a single hoverlabel will appear multiple points at the closest x- (or y-) coordinate within the `hoverdistance` with the caveat that no more than one hoverlabel will appear per trace. In this mode, spikelines are enabled by default perpendicular to the specified axis. If false, hover interactions are disabled. If `clickmode` includes the *select* flag, `hovermode` defaults to *closest*. If `clickmode` lacks the *select* flag, it defaults to *x* or *y* (depending on the trace's `orientation` value) for plots based on cartesian coordinates. For anything else the default value is *closest*.
type LayoutHovermode interface{}

const (
    LayoutHovermodeX string = "x"
    LayoutHovermodeY string = "y"
    LayoutHovermodeClosest string = "closest"
    LayoutHovermodeFalse bool = false
    LayoutHovermodeXUnified string = "x unified"
    LayoutHovermodeYUnified string = "y unified"
)

// LayoutLegendItemclick Determines the behavior on legend item click. *toggle* toggles the visibility of the item clicked on the graph. *toggleothers* makes the clicked item the sole visible item on the graph. *false* disable legend item click interactions.
type LayoutLegendItemclick interface{}

const (
    LayoutLegendItemclickToggle string = "toggle"
    LayoutLegendItemclickToggleothers string = "toggleothers"
    LayoutLegendItemclickFalse bool = false
)

// LayoutLegendItemdoubleclick Determines the behavior on legend item double-click. *toggle* toggles the visibility of the item clicked on the graph. *toggleothers* makes the clicked item the sole visible item on the graph. *false* disable legend item double-click interactions.
type LayoutLegendItemdoubleclick interface{}

const (
    LayoutLegendItemdoubleclickToggle string = "toggle"
    LayoutLegendItemdoubleclickToggleothers string = "toggleothers"
    LayoutLegendItemdoubleclickFalse bool = false
)

// LayoutLegendItemsizing Determines if the legend items symbols scale with their corresponding *trace* attributes or remain *constant* independent of the symbol size on the graph.
type LayoutLegendItemsizing string

const (
    LayoutLegendItemsizing_trace LayoutLegendItemsizing = "trace" // trace
    LayoutLegendItemsizing_constant LayoutLegendItemsizing = "constant" // constant
)

// LayoutLegendOrientation Sets the orientation of the legend.
type LayoutLegendOrientation string

const (
    LayoutLegendOrientation_v LayoutLegendOrientation = "v" // v
    LayoutLegendOrientation_h LayoutLegendOrientation = "h" // h
)

// LayoutLegendTitleSide Determines the location of legend's title with respect to the legend items. Defaulted to *top* with `orientation` is *h*. Defaulted to *left* with `orientation` is *v*. The *top left* options could be used to expand legend area in both x and y sides.
type LayoutLegendTitleSide string

const (
    LayoutLegendTitleSide_top LayoutLegendTitleSide = "top" // top
    LayoutLegendTitleSide_left LayoutLegendTitleSide = "left" // left
    LayoutLegendTitleSide_topleft LayoutLegendTitleSide = "top left" // top left
)

// LayoutLegendValign Sets the vertical alignment of the symbols with respect to their associated text.
type LayoutLegendValign string

const (
    LayoutLegendValign_top LayoutLegendValign = "top" // top
    LayoutLegendValign_middle LayoutLegendValign = "middle" // middle
    LayoutLegendValign_bottom LayoutLegendValign = "bottom" // bottom
)

// LayoutLegendXanchor Sets the legend's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the legend. Value *auto* anchors legends to the right for `x` values greater than or equal to 2/3, anchors legends to the left for `x` values less than or equal to 1/3 and anchors legends with respect to their center otherwise.
type LayoutLegendXanchor string

const (
    LayoutLegendXanchor_auto LayoutLegendXanchor = "auto" // auto
    LayoutLegendXanchor_left LayoutLegendXanchor = "left" // left
    LayoutLegendXanchor_center LayoutLegendXanchor = "center" // center
    LayoutLegendXanchor_right LayoutLegendXanchor = "right" // right
)

// LayoutLegendYanchor Sets the legend's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the legend. Value *auto* anchors legends at their bottom for `y` values less than or equal to 1/3, anchors legends to at their top for `y` values greater than or equal to 2/3 and anchors legends with respect to their middle otherwise.
type LayoutLegendYanchor string

const (
    LayoutLegendYanchor_auto LayoutLegendYanchor = "auto" // auto
    LayoutLegendYanchor_top LayoutLegendYanchor = "top" // top
    LayoutLegendYanchor_middle LayoutLegendYanchor = "middle" // middle
    LayoutLegendYanchor_bottom LayoutLegendYanchor = "bottom" // bottom
)

// LayoutModebarOrientation Sets the orientation of the modebar.
type LayoutModebarOrientation string

const (
    LayoutModebarOrientation_v LayoutModebarOrientation = "v" // v
    LayoutModebarOrientation_h LayoutModebarOrientation = "h" // h
)

// LayoutNewshapeDrawdirection When `dragmode` is set to *drawrect*, *drawline* or *drawcircle* this limits the drag to be horizontal, vertical or diagonal. Using *diagonal* there is no limit e.g. in drawing lines in any direction. *ortho* limits the draw to be either horizontal or vertical. *horizontal* allows horizontal extend. *vertical* allows vertical extend.
type LayoutNewshapeDrawdirection string

const (
    LayoutNewshapeDrawdirection_ortho LayoutNewshapeDrawdirection = "ortho" // ortho
    LayoutNewshapeDrawdirection_horizontal LayoutNewshapeDrawdirection = "horizontal" // horizontal
    LayoutNewshapeDrawdirection_vertical LayoutNewshapeDrawdirection = "vertical" // vertical
    LayoutNewshapeDrawdirection_diagonal LayoutNewshapeDrawdirection = "diagonal" // diagonal
)

// LayoutNewshapeFillrule Determines the path's interior. For more info please visit https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fill-rule
type LayoutNewshapeFillrule string

const (
    LayoutNewshapeFillrule_evenodd LayoutNewshapeFillrule = "evenodd" // evenodd
    LayoutNewshapeFillrule_nonzero LayoutNewshapeFillrule = "nonzero" // nonzero
)

// LayoutNewshapeLayer Specifies whether new shapes are drawn below or above traces.
type LayoutNewshapeLayer string

const (
    LayoutNewshapeLayer_below LayoutNewshapeLayer = "below" // below
    LayoutNewshapeLayer_above LayoutNewshapeLayer = "above" // above
)

// LayoutPolarAngularaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutPolarAngularaxisCategoryorder string

const (
    LayoutPolarAngularaxisCategoryorder_trace LayoutPolarAngularaxisCategoryorder = "trace" // trace
    LayoutPolarAngularaxisCategoryorder_categoryascending LayoutPolarAngularaxisCategoryorder = "category ascending" // category ascending
    LayoutPolarAngularaxisCategoryorder_categorydescending LayoutPolarAngularaxisCategoryorder = "category descending" // category descending
    LayoutPolarAngularaxisCategoryorder_array LayoutPolarAngularaxisCategoryorder = "array" // array
    LayoutPolarAngularaxisCategoryorder_totalascending LayoutPolarAngularaxisCategoryorder = "total ascending" // total ascending
    LayoutPolarAngularaxisCategoryorder_totaldescending LayoutPolarAngularaxisCategoryorder = "total descending" // total descending
    LayoutPolarAngularaxisCategoryorder_minascending LayoutPolarAngularaxisCategoryorder = "min ascending" // min ascending
    LayoutPolarAngularaxisCategoryorder_mindescending LayoutPolarAngularaxisCategoryorder = "min descending" // min descending
    LayoutPolarAngularaxisCategoryorder_maxascending LayoutPolarAngularaxisCategoryorder = "max ascending" // max ascending
    LayoutPolarAngularaxisCategoryorder_maxdescending LayoutPolarAngularaxisCategoryorder = "max descending" // max descending
    LayoutPolarAngularaxisCategoryorder_sumascending LayoutPolarAngularaxisCategoryorder = "sum ascending" // sum ascending
    LayoutPolarAngularaxisCategoryorder_sumdescending LayoutPolarAngularaxisCategoryorder = "sum descending" // sum descending
    LayoutPolarAngularaxisCategoryorder_meanascending LayoutPolarAngularaxisCategoryorder = "mean ascending" // mean ascending
    LayoutPolarAngularaxisCategoryorder_meandescending LayoutPolarAngularaxisCategoryorder = "mean descending" // mean descending
    LayoutPolarAngularaxisCategoryorder_medianascending LayoutPolarAngularaxisCategoryorder = "median ascending" // median ascending
    LayoutPolarAngularaxisCategoryorder_mediandescending LayoutPolarAngularaxisCategoryorder = "median descending" // median descending
)

// LayoutPolarAngularaxisDirection Sets the direction corresponding to positive angles.
type LayoutPolarAngularaxisDirection string

const (
    LayoutPolarAngularaxisDirection_counterclockwise LayoutPolarAngularaxisDirection = "counterclockwise" // counterclockwise
    LayoutPolarAngularaxisDirection_clockwise LayoutPolarAngularaxisDirection = "clockwise" // clockwise
)

// LayoutPolarAngularaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutPolarAngularaxisExponentformat string

const (
    LayoutPolarAngularaxisExponentformat_none LayoutPolarAngularaxisExponentformat = "none" // none
    LayoutPolarAngularaxisExponentformat_e LayoutPolarAngularaxisExponentformat = "e" // e
    LayoutPolarAngularaxisExponentformat_E LayoutPolarAngularaxisExponentformat = "E" // E
    LayoutPolarAngularaxisExponentformat_power LayoutPolarAngularaxisExponentformat = "power" // power
    LayoutPolarAngularaxisExponentformat_SI LayoutPolarAngularaxisExponentformat = "SI" // SI
    LayoutPolarAngularaxisExponentformat_B LayoutPolarAngularaxisExponentformat = "B" // B
)

// LayoutPolarAngularaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutPolarAngularaxisLayer string

const (
    LayoutPolarAngularaxisLayer_abovetraces LayoutPolarAngularaxisLayer = "above traces" // above traces
    LayoutPolarAngularaxisLayer_belowtraces LayoutPolarAngularaxisLayer = "below traces" // below traces
)

// LayoutPolarAngularaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutPolarAngularaxisShowexponent string

const (
    LayoutPolarAngularaxisShowexponent_all LayoutPolarAngularaxisShowexponent = "all" // all
    LayoutPolarAngularaxisShowexponent_first LayoutPolarAngularaxisShowexponent = "first" // first
    LayoutPolarAngularaxisShowexponent_last LayoutPolarAngularaxisShowexponent = "last" // last
    LayoutPolarAngularaxisShowexponent_none LayoutPolarAngularaxisShowexponent = "none" // none
)

// LayoutPolarAngularaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutPolarAngularaxisShowtickprefix string

const (
    LayoutPolarAngularaxisShowtickprefix_all LayoutPolarAngularaxisShowtickprefix = "all" // all
    LayoutPolarAngularaxisShowtickprefix_first LayoutPolarAngularaxisShowtickprefix = "first" // first
    LayoutPolarAngularaxisShowtickprefix_last LayoutPolarAngularaxisShowtickprefix = "last" // last
    LayoutPolarAngularaxisShowtickprefix_none LayoutPolarAngularaxisShowtickprefix = "none" // none
)

// LayoutPolarAngularaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutPolarAngularaxisShowticksuffix string

const (
    LayoutPolarAngularaxisShowticksuffix_all LayoutPolarAngularaxisShowticksuffix = "all" // all
    LayoutPolarAngularaxisShowticksuffix_first LayoutPolarAngularaxisShowticksuffix = "first" // first
    LayoutPolarAngularaxisShowticksuffix_last LayoutPolarAngularaxisShowticksuffix = "last" // last
    LayoutPolarAngularaxisShowticksuffix_none LayoutPolarAngularaxisShowticksuffix = "none" // none
)

// LayoutPolarAngularaxisThetaunit Sets the format unit of the formatted *theta* values. Has an effect only when `angularaxis.type` is *linear*.
type LayoutPolarAngularaxisThetaunit string

const (
    LayoutPolarAngularaxisThetaunit_radians LayoutPolarAngularaxisThetaunit = "radians" // radians
    LayoutPolarAngularaxisThetaunit_degrees LayoutPolarAngularaxisThetaunit = "degrees" // degrees
)

// LayoutPolarAngularaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutPolarAngularaxisTickmode string

const (
    LayoutPolarAngularaxisTickmode_auto LayoutPolarAngularaxisTickmode = "auto" // auto
    LayoutPolarAngularaxisTickmode_linear LayoutPolarAngularaxisTickmode = "linear" // linear
    LayoutPolarAngularaxisTickmode_array LayoutPolarAngularaxisTickmode = "array" // array
)

// LayoutPolarAngularaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutPolarAngularaxisTicks string

const (
    LayoutPolarAngularaxisTicks_outside LayoutPolarAngularaxisTicks = "outside" // outside
    LayoutPolarAngularaxisTicks_inside LayoutPolarAngularaxisTicks = "inside" // inside
)

// LayoutPolarAngularaxisType Sets the angular axis type. If *linear*, set `thetaunit` to determine the unit in which axis value are shown. If *category, use `period` to set the number of integer coordinates around polar axis.
type LayoutPolarAngularaxisType string

const (
    LayoutPolarAngularaxisType__ LayoutPolarAngularaxisType = "-" // -
    LayoutPolarAngularaxisType_linear LayoutPolarAngularaxisType = "linear" // linear
    LayoutPolarAngularaxisType_category LayoutPolarAngularaxisType = "category" // category
)

// LayoutPolarGridshape Determines if the radial axis grid lines and angular axis line are drawn as *circular* sectors or as *linear* (polygon) sectors. Has an effect only when the angular axis has `type` *category*. Note that `radialaxis.angle` is snapped to the angle of the closest vertex when `gridshape` is *circular* (so that radial axis scale is the same as the data scale).
type LayoutPolarGridshape string

const (
    LayoutPolarGridshape_circular LayoutPolarGridshape = "circular" // circular
    LayoutPolarGridshape_linear LayoutPolarGridshape = "linear" // linear
)

// LayoutPolarRadialaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutPolarRadialaxisAutorange interface{}

const (
    LayoutPolarRadialaxisAutorangeTrue bool = true
    LayoutPolarRadialaxisAutorangeFalse bool = false
    LayoutPolarRadialaxisAutorangeReversed string = "reversed"
)

// LayoutPolarRadialaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutPolarRadialaxisCalendar string

const (
    LayoutPolarRadialaxisCalendar_gregorian LayoutPolarRadialaxisCalendar = "gregorian" // gregorian
    LayoutPolarRadialaxisCalendar_chinese LayoutPolarRadialaxisCalendar = "chinese" // chinese
    LayoutPolarRadialaxisCalendar_coptic LayoutPolarRadialaxisCalendar = "coptic" // coptic
    LayoutPolarRadialaxisCalendar_discworld LayoutPolarRadialaxisCalendar = "discworld" // discworld
    LayoutPolarRadialaxisCalendar_ethiopian LayoutPolarRadialaxisCalendar = "ethiopian" // ethiopian
    LayoutPolarRadialaxisCalendar_hebrew LayoutPolarRadialaxisCalendar = "hebrew" // hebrew
    LayoutPolarRadialaxisCalendar_islamic LayoutPolarRadialaxisCalendar = "islamic" // islamic
    LayoutPolarRadialaxisCalendar_julian LayoutPolarRadialaxisCalendar = "julian" // julian
    LayoutPolarRadialaxisCalendar_mayan LayoutPolarRadialaxisCalendar = "mayan" // mayan
    LayoutPolarRadialaxisCalendar_nanakshahi LayoutPolarRadialaxisCalendar = "nanakshahi" // nanakshahi
    LayoutPolarRadialaxisCalendar_nepali LayoutPolarRadialaxisCalendar = "nepali" // nepali
    LayoutPolarRadialaxisCalendar_persian LayoutPolarRadialaxisCalendar = "persian" // persian
    LayoutPolarRadialaxisCalendar_jalali LayoutPolarRadialaxisCalendar = "jalali" // jalali
    LayoutPolarRadialaxisCalendar_taiwan LayoutPolarRadialaxisCalendar = "taiwan" // taiwan
    LayoutPolarRadialaxisCalendar_thai LayoutPolarRadialaxisCalendar = "thai" // thai
    LayoutPolarRadialaxisCalendar_ummalqura LayoutPolarRadialaxisCalendar = "ummalqura" // ummalqura
)

// LayoutPolarRadialaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutPolarRadialaxisCategoryorder string

const (
    LayoutPolarRadialaxisCategoryorder_trace LayoutPolarRadialaxisCategoryorder = "trace" // trace
    LayoutPolarRadialaxisCategoryorder_categoryascending LayoutPolarRadialaxisCategoryorder = "category ascending" // category ascending
    LayoutPolarRadialaxisCategoryorder_categorydescending LayoutPolarRadialaxisCategoryorder = "category descending" // category descending
    LayoutPolarRadialaxisCategoryorder_array LayoutPolarRadialaxisCategoryorder = "array" // array
    LayoutPolarRadialaxisCategoryorder_totalascending LayoutPolarRadialaxisCategoryorder = "total ascending" // total ascending
    LayoutPolarRadialaxisCategoryorder_totaldescending LayoutPolarRadialaxisCategoryorder = "total descending" // total descending
    LayoutPolarRadialaxisCategoryorder_minascending LayoutPolarRadialaxisCategoryorder = "min ascending" // min ascending
    LayoutPolarRadialaxisCategoryorder_mindescending LayoutPolarRadialaxisCategoryorder = "min descending" // min descending
    LayoutPolarRadialaxisCategoryorder_maxascending LayoutPolarRadialaxisCategoryorder = "max ascending" // max ascending
    LayoutPolarRadialaxisCategoryorder_maxdescending LayoutPolarRadialaxisCategoryorder = "max descending" // max descending
    LayoutPolarRadialaxisCategoryorder_sumascending LayoutPolarRadialaxisCategoryorder = "sum ascending" // sum ascending
    LayoutPolarRadialaxisCategoryorder_sumdescending LayoutPolarRadialaxisCategoryorder = "sum descending" // sum descending
    LayoutPolarRadialaxisCategoryorder_meanascending LayoutPolarRadialaxisCategoryorder = "mean ascending" // mean ascending
    LayoutPolarRadialaxisCategoryorder_meandescending LayoutPolarRadialaxisCategoryorder = "mean descending" // mean descending
    LayoutPolarRadialaxisCategoryorder_medianascending LayoutPolarRadialaxisCategoryorder = "median ascending" // median ascending
    LayoutPolarRadialaxisCategoryorder_mediandescending LayoutPolarRadialaxisCategoryorder = "median descending" // median descending
)

// LayoutPolarRadialaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutPolarRadialaxisExponentformat string

const (
    LayoutPolarRadialaxisExponentformat_none LayoutPolarRadialaxisExponentformat = "none" // none
    LayoutPolarRadialaxisExponentformat_e LayoutPolarRadialaxisExponentformat = "e" // e
    LayoutPolarRadialaxisExponentformat_E LayoutPolarRadialaxisExponentformat = "E" // E
    LayoutPolarRadialaxisExponentformat_power LayoutPolarRadialaxisExponentformat = "power" // power
    LayoutPolarRadialaxisExponentformat_SI LayoutPolarRadialaxisExponentformat = "SI" // SI
    LayoutPolarRadialaxisExponentformat_B LayoutPolarRadialaxisExponentformat = "B" // B
)

// LayoutPolarRadialaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutPolarRadialaxisLayer string

const (
    LayoutPolarRadialaxisLayer_abovetraces LayoutPolarRadialaxisLayer = "above traces" // above traces
    LayoutPolarRadialaxisLayer_belowtraces LayoutPolarRadialaxisLayer = "below traces" // below traces
)

// LayoutPolarRadialaxisRangemode If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. If *normal*, the range is computed in relation to the extrema of the input data (same behavior as for cartesian axes).
type LayoutPolarRadialaxisRangemode string

const (
    LayoutPolarRadialaxisRangemode_tozero LayoutPolarRadialaxisRangemode = "tozero" // tozero
    LayoutPolarRadialaxisRangemode_nonnegative LayoutPolarRadialaxisRangemode = "nonnegative" // nonnegative
    LayoutPolarRadialaxisRangemode_normal LayoutPolarRadialaxisRangemode = "normal" // normal
)

// LayoutPolarRadialaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutPolarRadialaxisShowexponent string

const (
    LayoutPolarRadialaxisShowexponent_all LayoutPolarRadialaxisShowexponent = "all" // all
    LayoutPolarRadialaxisShowexponent_first LayoutPolarRadialaxisShowexponent = "first" // first
    LayoutPolarRadialaxisShowexponent_last LayoutPolarRadialaxisShowexponent = "last" // last
    LayoutPolarRadialaxisShowexponent_none LayoutPolarRadialaxisShowexponent = "none" // none
)

// LayoutPolarRadialaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutPolarRadialaxisShowtickprefix string

const (
    LayoutPolarRadialaxisShowtickprefix_all LayoutPolarRadialaxisShowtickprefix = "all" // all
    LayoutPolarRadialaxisShowtickprefix_first LayoutPolarRadialaxisShowtickprefix = "first" // first
    LayoutPolarRadialaxisShowtickprefix_last LayoutPolarRadialaxisShowtickprefix = "last" // last
    LayoutPolarRadialaxisShowtickprefix_none LayoutPolarRadialaxisShowtickprefix = "none" // none
)

// LayoutPolarRadialaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutPolarRadialaxisShowticksuffix string

const (
    LayoutPolarRadialaxisShowticksuffix_all LayoutPolarRadialaxisShowticksuffix = "all" // all
    LayoutPolarRadialaxisShowticksuffix_first LayoutPolarRadialaxisShowticksuffix = "first" // first
    LayoutPolarRadialaxisShowticksuffix_last LayoutPolarRadialaxisShowticksuffix = "last" // last
    LayoutPolarRadialaxisShowticksuffix_none LayoutPolarRadialaxisShowticksuffix = "none" // none
)

// LayoutPolarRadialaxisSide Determines on which side of radial axis line the tick and tick labels appear.
type LayoutPolarRadialaxisSide string

const (
    LayoutPolarRadialaxisSide_clockwise LayoutPolarRadialaxisSide = "clockwise" // clockwise
    LayoutPolarRadialaxisSide_counterclockwise LayoutPolarRadialaxisSide = "counterclockwise" // counterclockwise
)

// LayoutPolarRadialaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutPolarRadialaxisTickmode string

const (
    LayoutPolarRadialaxisTickmode_auto LayoutPolarRadialaxisTickmode = "auto" // auto
    LayoutPolarRadialaxisTickmode_linear LayoutPolarRadialaxisTickmode = "linear" // linear
    LayoutPolarRadialaxisTickmode_array LayoutPolarRadialaxisTickmode = "array" // array
)

// LayoutPolarRadialaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutPolarRadialaxisTicks string

const (
    LayoutPolarRadialaxisTicks_outside LayoutPolarRadialaxisTicks = "outside" // outside
    LayoutPolarRadialaxisTicks_inside LayoutPolarRadialaxisTicks = "inside" // inside
)

// LayoutPolarRadialaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutPolarRadialaxisType string

const (
    LayoutPolarRadialaxisType__ LayoutPolarRadialaxisType = "-" // -
    LayoutPolarRadialaxisType_linear LayoutPolarRadialaxisType = "linear" // linear
    LayoutPolarRadialaxisType_log LayoutPolarRadialaxisType = "log" // log
    LayoutPolarRadialaxisType_date LayoutPolarRadialaxisType = "date" // date
    LayoutPolarRadialaxisType_category LayoutPolarRadialaxisType = "category" // category
)

// LayoutRadialaxisTickorientation Legacy polar charts are deprecated! Please switch to *polar* subplots. Sets the orientation (from the paper perspective) of the radial axis tick labels.
type LayoutRadialaxisTickorientation string

const (
    LayoutRadialaxisTickorientation_horizontal LayoutRadialaxisTickorientation = "horizontal" // horizontal
    LayoutRadialaxisTickorientation_vertical LayoutRadialaxisTickorientation = "vertical" // vertical
)

// LayoutSceneAspectmode If *cube*, this scene's axes are drawn as a cube, regardless of the axes' ranges. If *data*, this scene's axes are drawn in proportion with the axes' ranges. If *manual*, this scene's axes are drawn in proportion with the input of *aspectratio* (the default behavior if *aspectratio* is provided). If *auto*, this scene's axes are drawn using the results of *data* except when one axis is more than four times the size of the two others, where in that case the results of *cube* are used.
type LayoutSceneAspectmode string

const (
    LayoutSceneAspectmode_auto LayoutSceneAspectmode = "auto" // auto
    LayoutSceneAspectmode_cube LayoutSceneAspectmode = "cube" // cube
    LayoutSceneAspectmode_data LayoutSceneAspectmode = "data" // data
    LayoutSceneAspectmode_manual LayoutSceneAspectmode = "manual" // manual
)

// LayoutSceneCameraProjectionType Sets the projection type. The projection type could be either *perspective* or *orthographic*. The default is *perspective*.
type LayoutSceneCameraProjectionType string

const (
    LayoutSceneCameraProjectionType_perspective LayoutSceneCameraProjectionType = "perspective" // perspective
    LayoutSceneCameraProjectionType_orthographic LayoutSceneCameraProjectionType = "orthographic" // orthographic
)

// LayoutSceneDragmode Determines the mode of drag interactions for this scene.
type LayoutSceneDragmode interface{}

const (
    LayoutSceneDragmodeOrbit string = "orbit"
    LayoutSceneDragmodeTurntable string = "turntable"
    LayoutSceneDragmodeZoom string = "zoom"
    LayoutSceneDragmodePan string = "pan"
    LayoutSceneDragmodeFalse bool = false
)

// LayoutSceneHovermode Determines the mode of hover interactions for this scene.
type LayoutSceneHovermode interface{}

const (
    LayoutSceneHovermodeClosest string = "closest"
    LayoutSceneHovermodeFalse bool = false
)

// LayoutSceneXaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutSceneXaxisAutorange interface{}

const (
    LayoutSceneXaxisAutorangeTrue bool = true
    LayoutSceneXaxisAutorangeFalse bool = false
    LayoutSceneXaxisAutorangeReversed string = "reversed"
)

// LayoutSceneXaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutSceneXaxisCalendar string

const (
    LayoutSceneXaxisCalendar_gregorian LayoutSceneXaxisCalendar = "gregorian" // gregorian
    LayoutSceneXaxisCalendar_chinese LayoutSceneXaxisCalendar = "chinese" // chinese
    LayoutSceneXaxisCalendar_coptic LayoutSceneXaxisCalendar = "coptic" // coptic
    LayoutSceneXaxisCalendar_discworld LayoutSceneXaxisCalendar = "discworld" // discworld
    LayoutSceneXaxisCalendar_ethiopian LayoutSceneXaxisCalendar = "ethiopian" // ethiopian
    LayoutSceneXaxisCalendar_hebrew LayoutSceneXaxisCalendar = "hebrew" // hebrew
    LayoutSceneXaxisCalendar_islamic LayoutSceneXaxisCalendar = "islamic" // islamic
    LayoutSceneXaxisCalendar_julian LayoutSceneXaxisCalendar = "julian" // julian
    LayoutSceneXaxisCalendar_mayan LayoutSceneXaxisCalendar = "mayan" // mayan
    LayoutSceneXaxisCalendar_nanakshahi LayoutSceneXaxisCalendar = "nanakshahi" // nanakshahi
    LayoutSceneXaxisCalendar_nepali LayoutSceneXaxisCalendar = "nepali" // nepali
    LayoutSceneXaxisCalendar_persian LayoutSceneXaxisCalendar = "persian" // persian
    LayoutSceneXaxisCalendar_jalali LayoutSceneXaxisCalendar = "jalali" // jalali
    LayoutSceneXaxisCalendar_taiwan LayoutSceneXaxisCalendar = "taiwan" // taiwan
    LayoutSceneXaxisCalendar_thai LayoutSceneXaxisCalendar = "thai" // thai
    LayoutSceneXaxisCalendar_ummalqura LayoutSceneXaxisCalendar = "ummalqura" // ummalqura
)

// LayoutSceneXaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutSceneXaxisCategoryorder string

const (
    LayoutSceneXaxisCategoryorder_trace LayoutSceneXaxisCategoryorder = "trace" // trace
    LayoutSceneXaxisCategoryorder_categoryascending LayoutSceneXaxisCategoryorder = "category ascending" // category ascending
    LayoutSceneXaxisCategoryorder_categorydescending LayoutSceneXaxisCategoryorder = "category descending" // category descending
    LayoutSceneXaxisCategoryorder_array LayoutSceneXaxisCategoryorder = "array" // array
    LayoutSceneXaxisCategoryorder_totalascending LayoutSceneXaxisCategoryorder = "total ascending" // total ascending
    LayoutSceneXaxisCategoryorder_totaldescending LayoutSceneXaxisCategoryorder = "total descending" // total descending
    LayoutSceneXaxisCategoryorder_minascending LayoutSceneXaxisCategoryorder = "min ascending" // min ascending
    LayoutSceneXaxisCategoryorder_mindescending LayoutSceneXaxisCategoryorder = "min descending" // min descending
    LayoutSceneXaxisCategoryorder_maxascending LayoutSceneXaxisCategoryorder = "max ascending" // max ascending
    LayoutSceneXaxisCategoryorder_maxdescending LayoutSceneXaxisCategoryorder = "max descending" // max descending
    LayoutSceneXaxisCategoryorder_sumascending LayoutSceneXaxisCategoryorder = "sum ascending" // sum ascending
    LayoutSceneXaxisCategoryorder_sumdescending LayoutSceneXaxisCategoryorder = "sum descending" // sum descending
    LayoutSceneXaxisCategoryorder_meanascending LayoutSceneXaxisCategoryorder = "mean ascending" // mean ascending
    LayoutSceneXaxisCategoryorder_meandescending LayoutSceneXaxisCategoryorder = "mean descending" // mean descending
    LayoutSceneXaxisCategoryorder_medianascending LayoutSceneXaxisCategoryorder = "median ascending" // median ascending
    LayoutSceneXaxisCategoryorder_mediandescending LayoutSceneXaxisCategoryorder = "median descending" // median descending
)

// LayoutSceneXaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutSceneXaxisExponentformat string

const (
    LayoutSceneXaxisExponentformat_none LayoutSceneXaxisExponentformat = "none" // none
    LayoutSceneXaxisExponentformat_e LayoutSceneXaxisExponentformat = "e" // e
    LayoutSceneXaxisExponentformat_E LayoutSceneXaxisExponentformat = "E" // E
    LayoutSceneXaxisExponentformat_power LayoutSceneXaxisExponentformat = "power" // power
    LayoutSceneXaxisExponentformat_SI LayoutSceneXaxisExponentformat = "SI" // SI
    LayoutSceneXaxisExponentformat_B LayoutSceneXaxisExponentformat = "B" // B
)

// LayoutSceneXaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutSceneXaxisMirror interface{}

const (
    LayoutSceneXaxisMirrorTrue bool = true
    LayoutSceneXaxisMirrorTicks string = "ticks"
    LayoutSceneXaxisMirrorFalse bool = false
    LayoutSceneXaxisMirrorAll string = "all"
    LayoutSceneXaxisMirrorAllticks string = "allticks"
)

// LayoutSceneXaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutSceneXaxisRangemode string

const (
    LayoutSceneXaxisRangemode_normal LayoutSceneXaxisRangemode = "normal" // normal
    LayoutSceneXaxisRangemode_tozero LayoutSceneXaxisRangemode = "tozero" // tozero
    LayoutSceneXaxisRangemode_nonnegative LayoutSceneXaxisRangemode = "nonnegative" // nonnegative
)

// LayoutSceneXaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutSceneXaxisShowexponent string

const (
    LayoutSceneXaxisShowexponent_all LayoutSceneXaxisShowexponent = "all" // all
    LayoutSceneXaxisShowexponent_first LayoutSceneXaxisShowexponent = "first" // first
    LayoutSceneXaxisShowexponent_last LayoutSceneXaxisShowexponent = "last" // last
    LayoutSceneXaxisShowexponent_none LayoutSceneXaxisShowexponent = "none" // none
)

// LayoutSceneXaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutSceneXaxisShowtickprefix string

const (
    LayoutSceneXaxisShowtickprefix_all LayoutSceneXaxisShowtickprefix = "all" // all
    LayoutSceneXaxisShowtickprefix_first LayoutSceneXaxisShowtickprefix = "first" // first
    LayoutSceneXaxisShowtickprefix_last LayoutSceneXaxisShowtickprefix = "last" // last
    LayoutSceneXaxisShowtickprefix_none LayoutSceneXaxisShowtickprefix = "none" // none
)

// LayoutSceneXaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutSceneXaxisShowticksuffix string

const (
    LayoutSceneXaxisShowticksuffix_all LayoutSceneXaxisShowticksuffix = "all" // all
    LayoutSceneXaxisShowticksuffix_first LayoutSceneXaxisShowticksuffix = "first" // first
    LayoutSceneXaxisShowticksuffix_last LayoutSceneXaxisShowticksuffix = "last" // last
    LayoutSceneXaxisShowticksuffix_none LayoutSceneXaxisShowticksuffix = "none" // none
)

// LayoutSceneXaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutSceneXaxisTickmode string

const (
    LayoutSceneXaxisTickmode_auto LayoutSceneXaxisTickmode = "auto" // auto
    LayoutSceneXaxisTickmode_linear LayoutSceneXaxisTickmode = "linear" // linear
    LayoutSceneXaxisTickmode_array LayoutSceneXaxisTickmode = "array" // array
)

// LayoutSceneXaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutSceneXaxisTicks string

const (
    LayoutSceneXaxisTicks_outside LayoutSceneXaxisTicks = "outside" // outside
    LayoutSceneXaxisTicks_inside LayoutSceneXaxisTicks = "inside" // inside
)

// LayoutSceneXaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutSceneXaxisType string

const (
    LayoutSceneXaxisType__ LayoutSceneXaxisType = "-" // -
    LayoutSceneXaxisType_linear LayoutSceneXaxisType = "linear" // linear
    LayoutSceneXaxisType_log LayoutSceneXaxisType = "log" // log
    LayoutSceneXaxisType_date LayoutSceneXaxisType = "date" // date
    LayoutSceneXaxisType_category LayoutSceneXaxisType = "category" // category
)

// LayoutSceneYaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutSceneYaxisAutorange interface{}

const (
    LayoutSceneYaxisAutorangeTrue bool = true
    LayoutSceneYaxisAutorangeFalse bool = false
    LayoutSceneYaxisAutorangeReversed string = "reversed"
)

// LayoutSceneYaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutSceneYaxisCalendar string

const (
    LayoutSceneYaxisCalendar_gregorian LayoutSceneYaxisCalendar = "gregorian" // gregorian
    LayoutSceneYaxisCalendar_chinese LayoutSceneYaxisCalendar = "chinese" // chinese
    LayoutSceneYaxisCalendar_coptic LayoutSceneYaxisCalendar = "coptic" // coptic
    LayoutSceneYaxisCalendar_discworld LayoutSceneYaxisCalendar = "discworld" // discworld
    LayoutSceneYaxisCalendar_ethiopian LayoutSceneYaxisCalendar = "ethiopian" // ethiopian
    LayoutSceneYaxisCalendar_hebrew LayoutSceneYaxisCalendar = "hebrew" // hebrew
    LayoutSceneYaxisCalendar_islamic LayoutSceneYaxisCalendar = "islamic" // islamic
    LayoutSceneYaxisCalendar_julian LayoutSceneYaxisCalendar = "julian" // julian
    LayoutSceneYaxisCalendar_mayan LayoutSceneYaxisCalendar = "mayan" // mayan
    LayoutSceneYaxisCalendar_nanakshahi LayoutSceneYaxisCalendar = "nanakshahi" // nanakshahi
    LayoutSceneYaxisCalendar_nepali LayoutSceneYaxisCalendar = "nepali" // nepali
    LayoutSceneYaxisCalendar_persian LayoutSceneYaxisCalendar = "persian" // persian
    LayoutSceneYaxisCalendar_jalali LayoutSceneYaxisCalendar = "jalali" // jalali
    LayoutSceneYaxisCalendar_taiwan LayoutSceneYaxisCalendar = "taiwan" // taiwan
    LayoutSceneYaxisCalendar_thai LayoutSceneYaxisCalendar = "thai" // thai
    LayoutSceneYaxisCalendar_ummalqura LayoutSceneYaxisCalendar = "ummalqura" // ummalqura
)

// LayoutSceneYaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutSceneYaxisCategoryorder string

const (
    LayoutSceneYaxisCategoryorder_trace LayoutSceneYaxisCategoryorder = "trace" // trace
    LayoutSceneYaxisCategoryorder_categoryascending LayoutSceneYaxisCategoryorder = "category ascending" // category ascending
    LayoutSceneYaxisCategoryorder_categorydescending LayoutSceneYaxisCategoryorder = "category descending" // category descending
    LayoutSceneYaxisCategoryorder_array LayoutSceneYaxisCategoryorder = "array" // array
    LayoutSceneYaxisCategoryorder_totalascending LayoutSceneYaxisCategoryorder = "total ascending" // total ascending
    LayoutSceneYaxisCategoryorder_totaldescending LayoutSceneYaxisCategoryorder = "total descending" // total descending
    LayoutSceneYaxisCategoryorder_minascending LayoutSceneYaxisCategoryorder = "min ascending" // min ascending
    LayoutSceneYaxisCategoryorder_mindescending LayoutSceneYaxisCategoryorder = "min descending" // min descending
    LayoutSceneYaxisCategoryorder_maxascending LayoutSceneYaxisCategoryorder = "max ascending" // max ascending
    LayoutSceneYaxisCategoryorder_maxdescending LayoutSceneYaxisCategoryorder = "max descending" // max descending
    LayoutSceneYaxisCategoryorder_sumascending LayoutSceneYaxisCategoryorder = "sum ascending" // sum ascending
    LayoutSceneYaxisCategoryorder_sumdescending LayoutSceneYaxisCategoryorder = "sum descending" // sum descending
    LayoutSceneYaxisCategoryorder_meanascending LayoutSceneYaxisCategoryorder = "mean ascending" // mean ascending
    LayoutSceneYaxisCategoryorder_meandescending LayoutSceneYaxisCategoryorder = "mean descending" // mean descending
    LayoutSceneYaxisCategoryorder_medianascending LayoutSceneYaxisCategoryorder = "median ascending" // median ascending
    LayoutSceneYaxisCategoryorder_mediandescending LayoutSceneYaxisCategoryorder = "median descending" // median descending
)

// LayoutSceneYaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutSceneYaxisExponentformat string

const (
    LayoutSceneYaxisExponentformat_none LayoutSceneYaxisExponentformat = "none" // none
    LayoutSceneYaxisExponentformat_e LayoutSceneYaxisExponentformat = "e" // e
    LayoutSceneYaxisExponentformat_E LayoutSceneYaxisExponentformat = "E" // E
    LayoutSceneYaxisExponentformat_power LayoutSceneYaxisExponentformat = "power" // power
    LayoutSceneYaxisExponentformat_SI LayoutSceneYaxisExponentformat = "SI" // SI
    LayoutSceneYaxisExponentformat_B LayoutSceneYaxisExponentformat = "B" // B
)

// LayoutSceneYaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutSceneYaxisMirror interface{}

const (
    LayoutSceneYaxisMirrorTrue bool = true
    LayoutSceneYaxisMirrorTicks string = "ticks"
    LayoutSceneYaxisMirrorFalse bool = false
    LayoutSceneYaxisMirrorAll string = "all"
    LayoutSceneYaxisMirrorAllticks string = "allticks"
)

// LayoutSceneYaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutSceneYaxisRangemode string

const (
    LayoutSceneYaxisRangemode_normal LayoutSceneYaxisRangemode = "normal" // normal
    LayoutSceneYaxisRangemode_tozero LayoutSceneYaxisRangemode = "tozero" // tozero
    LayoutSceneYaxisRangemode_nonnegative LayoutSceneYaxisRangemode = "nonnegative" // nonnegative
)

// LayoutSceneYaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutSceneYaxisShowexponent string

const (
    LayoutSceneYaxisShowexponent_all LayoutSceneYaxisShowexponent = "all" // all
    LayoutSceneYaxisShowexponent_first LayoutSceneYaxisShowexponent = "first" // first
    LayoutSceneYaxisShowexponent_last LayoutSceneYaxisShowexponent = "last" // last
    LayoutSceneYaxisShowexponent_none LayoutSceneYaxisShowexponent = "none" // none
)

// LayoutSceneYaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutSceneYaxisShowtickprefix string

const (
    LayoutSceneYaxisShowtickprefix_all LayoutSceneYaxisShowtickprefix = "all" // all
    LayoutSceneYaxisShowtickprefix_first LayoutSceneYaxisShowtickprefix = "first" // first
    LayoutSceneYaxisShowtickprefix_last LayoutSceneYaxisShowtickprefix = "last" // last
    LayoutSceneYaxisShowtickprefix_none LayoutSceneYaxisShowtickprefix = "none" // none
)

// LayoutSceneYaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutSceneYaxisShowticksuffix string

const (
    LayoutSceneYaxisShowticksuffix_all LayoutSceneYaxisShowticksuffix = "all" // all
    LayoutSceneYaxisShowticksuffix_first LayoutSceneYaxisShowticksuffix = "first" // first
    LayoutSceneYaxisShowticksuffix_last LayoutSceneYaxisShowticksuffix = "last" // last
    LayoutSceneYaxisShowticksuffix_none LayoutSceneYaxisShowticksuffix = "none" // none
)

// LayoutSceneYaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutSceneYaxisTickmode string

const (
    LayoutSceneYaxisTickmode_auto LayoutSceneYaxisTickmode = "auto" // auto
    LayoutSceneYaxisTickmode_linear LayoutSceneYaxisTickmode = "linear" // linear
    LayoutSceneYaxisTickmode_array LayoutSceneYaxisTickmode = "array" // array
)

// LayoutSceneYaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutSceneYaxisTicks string

const (
    LayoutSceneYaxisTicks_outside LayoutSceneYaxisTicks = "outside" // outside
    LayoutSceneYaxisTicks_inside LayoutSceneYaxisTicks = "inside" // inside
)

// LayoutSceneYaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutSceneYaxisType string

const (
    LayoutSceneYaxisType__ LayoutSceneYaxisType = "-" // -
    LayoutSceneYaxisType_linear LayoutSceneYaxisType = "linear" // linear
    LayoutSceneYaxisType_log LayoutSceneYaxisType = "log" // log
    LayoutSceneYaxisType_date LayoutSceneYaxisType = "date" // date
    LayoutSceneYaxisType_category LayoutSceneYaxisType = "category" // category
)

// LayoutSceneZaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutSceneZaxisAutorange interface{}

const (
    LayoutSceneZaxisAutorangeTrue bool = true
    LayoutSceneZaxisAutorangeFalse bool = false
    LayoutSceneZaxisAutorangeReversed string = "reversed"
)

// LayoutSceneZaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutSceneZaxisCalendar string

const (
    LayoutSceneZaxisCalendar_gregorian LayoutSceneZaxisCalendar = "gregorian" // gregorian
    LayoutSceneZaxisCalendar_chinese LayoutSceneZaxisCalendar = "chinese" // chinese
    LayoutSceneZaxisCalendar_coptic LayoutSceneZaxisCalendar = "coptic" // coptic
    LayoutSceneZaxisCalendar_discworld LayoutSceneZaxisCalendar = "discworld" // discworld
    LayoutSceneZaxisCalendar_ethiopian LayoutSceneZaxisCalendar = "ethiopian" // ethiopian
    LayoutSceneZaxisCalendar_hebrew LayoutSceneZaxisCalendar = "hebrew" // hebrew
    LayoutSceneZaxisCalendar_islamic LayoutSceneZaxisCalendar = "islamic" // islamic
    LayoutSceneZaxisCalendar_julian LayoutSceneZaxisCalendar = "julian" // julian
    LayoutSceneZaxisCalendar_mayan LayoutSceneZaxisCalendar = "mayan" // mayan
    LayoutSceneZaxisCalendar_nanakshahi LayoutSceneZaxisCalendar = "nanakshahi" // nanakshahi
    LayoutSceneZaxisCalendar_nepali LayoutSceneZaxisCalendar = "nepali" // nepali
    LayoutSceneZaxisCalendar_persian LayoutSceneZaxisCalendar = "persian" // persian
    LayoutSceneZaxisCalendar_jalali LayoutSceneZaxisCalendar = "jalali" // jalali
    LayoutSceneZaxisCalendar_taiwan LayoutSceneZaxisCalendar = "taiwan" // taiwan
    LayoutSceneZaxisCalendar_thai LayoutSceneZaxisCalendar = "thai" // thai
    LayoutSceneZaxisCalendar_ummalqura LayoutSceneZaxisCalendar = "ummalqura" // ummalqura
)

// LayoutSceneZaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutSceneZaxisCategoryorder string

const (
    LayoutSceneZaxisCategoryorder_trace LayoutSceneZaxisCategoryorder = "trace" // trace
    LayoutSceneZaxisCategoryorder_categoryascending LayoutSceneZaxisCategoryorder = "category ascending" // category ascending
    LayoutSceneZaxisCategoryorder_categorydescending LayoutSceneZaxisCategoryorder = "category descending" // category descending
    LayoutSceneZaxisCategoryorder_array LayoutSceneZaxisCategoryorder = "array" // array
    LayoutSceneZaxisCategoryorder_totalascending LayoutSceneZaxisCategoryorder = "total ascending" // total ascending
    LayoutSceneZaxisCategoryorder_totaldescending LayoutSceneZaxisCategoryorder = "total descending" // total descending
    LayoutSceneZaxisCategoryorder_minascending LayoutSceneZaxisCategoryorder = "min ascending" // min ascending
    LayoutSceneZaxisCategoryorder_mindescending LayoutSceneZaxisCategoryorder = "min descending" // min descending
    LayoutSceneZaxisCategoryorder_maxascending LayoutSceneZaxisCategoryorder = "max ascending" // max ascending
    LayoutSceneZaxisCategoryorder_maxdescending LayoutSceneZaxisCategoryorder = "max descending" // max descending
    LayoutSceneZaxisCategoryorder_sumascending LayoutSceneZaxisCategoryorder = "sum ascending" // sum ascending
    LayoutSceneZaxisCategoryorder_sumdescending LayoutSceneZaxisCategoryorder = "sum descending" // sum descending
    LayoutSceneZaxisCategoryorder_meanascending LayoutSceneZaxisCategoryorder = "mean ascending" // mean ascending
    LayoutSceneZaxisCategoryorder_meandescending LayoutSceneZaxisCategoryorder = "mean descending" // mean descending
    LayoutSceneZaxisCategoryorder_medianascending LayoutSceneZaxisCategoryorder = "median ascending" // median ascending
    LayoutSceneZaxisCategoryorder_mediandescending LayoutSceneZaxisCategoryorder = "median descending" // median descending
)

// LayoutSceneZaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutSceneZaxisExponentformat string

const (
    LayoutSceneZaxisExponentformat_none LayoutSceneZaxisExponentformat = "none" // none
    LayoutSceneZaxisExponentformat_e LayoutSceneZaxisExponentformat = "e" // e
    LayoutSceneZaxisExponentformat_E LayoutSceneZaxisExponentformat = "E" // E
    LayoutSceneZaxisExponentformat_power LayoutSceneZaxisExponentformat = "power" // power
    LayoutSceneZaxisExponentformat_SI LayoutSceneZaxisExponentformat = "SI" // SI
    LayoutSceneZaxisExponentformat_B LayoutSceneZaxisExponentformat = "B" // B
)

// LayoutSceneZaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutSceneZaxisMirror interface{}

const (
    LayoutSceneZaxisMirrorTrue bool = true
    LayoutSceneZaxisMirrorTicks string = "ticks"
    LayoutSceneZaxisMirrorFalse bool = false
    LayoutSceneZaxisMirrorAll string = "all"
    LayoutSceneZaxisMirrorAllticks string = "allticks"
)

// LayoutSceneZaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutSceneZaxisRangemode string

const (
    LayoutSceneZaxisRangemode_normal LayoutSceneZaxisRangemode = "normal" // normal
    LayoutSceneZaxisRangemode_tozero LayoutSceneZaxisRangemode = "tozero" // tozero
    LayoutSceneZaxisRangemode_nonnegative LayoutSceneZaxisRangemode = "nonnegative" // nonnegative
)

// LayoutSceneZaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutSceneZaxisShowexponent string

const (
    LayoutSceneZaxisShowexponent_all LayoutSceneZaxisShowexponent = "all" // all
    LayoutSceneZaxisShowexponent_first LayoutSceneZaxisShowexponent = "first" // first
    LayoutSceneZaxisShowexponent_last LayoutSceneZaxisShowexponent = "last" // last
    LayoutSceneZaxisShowexponent_none LayoutSceneZaxisShowexponent = "none" // none
)

// LayoutSceneZaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutSceneZaxisShowtickprefix string

const (
    LayoutSceneZaxisShowtickprefix_all LayoutSceneZaxisShowtickprefix = "all" // all
    LayoutSceneZaxisShowtickprefix_first LayoutSceneZaxisShowtickprefix = "first" // first
    LayoutSceneZaxisShowtickprefix_last LayoutSceneZaxisShowtickprefix = "last" // last
    LayoutSceneZaxisShowtickprefix_none LayoutSceneZaxisShowtickprefix = "none" // none
)

// LayoutSceneZaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutSceneZaxisShowticksuffix string

const (
    LayoutSceneZaxisShowticksuffix_all LayoutSceneZaxisShowticksuffix = "all" // all
    LayoutSceneZaxisShowticksuffix_first LayoutSceneZaxisShowticksuffix = "first" // first
    LayoutSceneZaxisShowticksuffix_last LayoutSceneZaxisShowticksuffix = "last" // last
    LayoutSceneZaxisShowticksuffix_none LayoutSceneZaxisShowticksuffix = "none" // none
)

// LayoutSceneZaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutSceneZaxisTickmode string

const (
    LayoutSceneZaxisTickmode_auto LayoutSceneZaxisTickmode = "auto" // auto
    LayoutSceneZaxisTickmode_linear LayoutSceneZaxisTickmode = "linear" // linear
    LayoutSceneZaxisTickmode_array LayoutSceneZaxisTickmode = "array" // array
)

// LayoutSceneZaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutSceneZaxisTicks string

const (
    LayoutSceneZaxisTicks_outside LayoutSceneZaxisTicks = "outside" // outside
    LayoutSceneZaxisTicks_inside LayoutSceneZaxisTicks = "inside" // inside
)

// LayoutSceneZaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutSceneZaxisType string

const (
    LayoutSceneZaxisType__ LayoutSceneZaxisType = "-" // -
    LayoutSceneZaxisType_linear LayoutSceneZaxisType = "linear" // linear
    LayoutSceneZaxisType_log LayoutSceneZaxisType = "log" // log
    LayoutSceneZaxisType_date LayoutSceneZaxisType = "date" // date
    LayoutSceneZaxisType_category LayoutSceneZaxisType = "category" // category
)

// LayoutSelectdirection When `dragmode` is set to *select*, this limits the selection of the drag to horizontal, vertical or diagonal. *h* only allows horizontal selection, *v* only vertical, *d* only diagonal and *any* sets no limit.
type LayoutSelectdirection string

const (
    LayoutSelectdirection_h LayoutSelectdirection = "h" // h
    LayoutSelectdirection_v LayoutSelectdirection = "v" // v
    LayoutSelectdirection_d LayoutSelectdirection = "d" // d
    LayoutSelectdirection_any LayoutSelectdirection = "any" // any
)

// LayoutTernaryAaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutTernaryAaxisExponentformat string

const (
    LayoutTernaryAaxisExponentformat_none LayoutTernaryAaxisExponentformat = "none" // none
    LayoutTernaryAaxisExponentformat_e LayoutTernaryAaxisExponentformat = "e" // e
    LayoutTernaryAaxisExponentformat_E LayoutTernaryAaxisExponentformat = "E" // E
    LayoutTernaryAaxisExponentformat_power LayoutTernaryAaxisExponentformat = "power" // power
    LayoutTernaryAaxisExponentformat_SI LayoutTernaryAaxisExponentformat = "SI" // SI
    LayoutTernaryAaxisExponentformat_B LayoutTernaryAaxisExponentformat = "B" // B
)

// LayoutTernaryAaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutTernaryAaxisLayer string

const (
    LayoutTernaryAaxisLayer_abovetraces LayoutTernaryAaxisLayer = "above traces" // above traces
    LayoutTernaryAaxisLayer_belowtraces LayoutTernaryAaxisLayer = "below traces" // below traces
)

// LayoutTernaryAaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutTernaryAaxisShowexponent string

const (
    LayoutTernaryAaxisShowexponent_all LayoutTernaryAaxisShowexponent = "all" // all
    LayoutTernaryAaxisShowexponent_first LayoutTernaryAaxisShowexponent = "first" // first
    LayoutTernaryAaxisShowexponent_last LayoutTernaryAaxisShowexponent = "last" // last
    LayoutTernaryAaxisShowexponent_none LayoutTernaryAaxisShowexponent = "none" // none
)

// LayoutTernaryAaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutTernaryAaxisShowtickprefix string

const (
    LayoutTernaryAaxisShowtickprefix_all LayoutTernaryAaxisShowtickprefix = "all" // all
    LayoutTernaryAaxisShowtickprefix_first LayoutTernaryAaxisShowtickprefix = "first" // first
    LayoutTernaryAaxisShowtickprefix_last LayoutTernaryAaxisShowtickprefix = "last" // last
    LayoutTernaryAaxisShowtickprefix_none LayoutTernaryAaxisShowtickprefix = "none" // none
)

// LayoutTernaryAaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutTernaryAaxisShowticksuffix string

const (
    LayoutTernaryAaxisShowticksuffix_all LayoutTernaryAaxisShowticksuffix = "all" // all
    LayoutTernaryAaxisShowticksuffix_first LayoutTernaryAaxisShowticksuffix = "first" // first
    LayoutTernaryAaxisShowticksuffix_last LayoutTernaryAaxisShowticksuffix = "last" // last
    LayoutTernaryAaxisShowticksuffix_none LayoutTernaryAaxisShowticksuffix = "none" // none
)

// LayoutTernaryAaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutTernaryAaxisTickmode string

const (
    LayoutTernaryAaxisTickmode_auto LayoutTernaryAaxisTickmode = "auto" // auto
    LayoutTernaryAaxisTickmode_linear LayoutTernaryAaxisTickmode = "linear" // linear
    LayoutTernaryAaxisTickmode_array LayoutTernaryAaxisTickmode = "array" // array
)

// LayoutTernaryAaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutTernaryAaxisTicks string

const (
    LayoutTernaryAaxisTicks_outside LayoutTernaryAaxisTicks = "outside" // outside
    LayoutTernaryAaxisTicks_inside LayoutTernaryAaxisTicks = "inside" // inside
)

// LayoutTernaryBaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutTernaryBaxisExponentformat string

const (
    LayoutTernaryBaxisExponentformat_none LayoutTernaryBaxisExponentformat = "none" // none
    LayoutTernaryBaxisExponentformat_e LayoutTernaryBaxisExponentformat = "e" // e
    LayoutTernaryBaxisExponentformat_E LayoutTernaryBaxisExponentformat = "E" // E
    LayoutTernaryBaxisExponentformat_power LayoutTernaryBaxisExponentformat = "power" // power
    LayoutTernaryBaxisExponentformat_SI LayoutTernaryBaxisExponentformat = "SI" // SI
    LayoutTernaryBaxisExponentformat_B LayoutTernaryBaxisExponentformat = "B" // B
)

// LayoutTernaryBaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutTernaryBaxisLayer string

const (
    LayoutTernaryBaxisLayer_abovetraces LayoutTernaryBaxisLayer = "above traces" // above traces
    LayoutTernaryBaxisLayer_belowtraces LayoutTernaryBaxisLayer = "below traces" // below traces
)

// LayoutTernaryBaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutTernaryBaxisShowexponent string

const (
    LayoutTernaryBaxisShowexponent_all LayoutTernaryBaxisShowexponent = "all" // all
    LayoutTernaryBaxisShowexponent_first LayoutTernaryBaxisShowexponent = "first" // first
    LayoutTernaryBaxisShowexponent_last LayoutTernaryBaxisShowexponent = "last" // last
    LayoutTernaryBaxisShowexponent_none LayoutTernaryBaxisShowexponent = "none" // none
)

// LayoutTernaryBaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutTernaryBaxisShowtickprefix string

const (
    LayoutTernaryBaxisShowtickprefix_all LayoutTernaryBaxisShowtickprefix = "all" // all
    LayoutTernaryBaxisShowtickprefix_first LayoutTernaryBaxisShowtickprefix = "first" // first
    LayoutTernaryBaxisShowtickprefix_last LayoutTernaryBaxisShowtickprefix = "last" // last
    LayoutTernaryBaxisShowtickprefix_none LayoutTernaryBaxisShowtickprefix = "none" // none
)

// LayoutTernaryBaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutTernaryBaxisShowticksuffix string

const (
    LayoutTernaryBaxisShowticksuffix_all LayoutTernaryBaxisShowticksuffix = "all" // all
    LayoutTernaryBaxisShowticksuffix_first LayoutTernaryBaxisShowticksuffix = "first" // first
    LayoutTernaryBaxisShowticksuffix_last LayoutTernaryBaxisShowticksuffix = "last" // last
    LayoutTernaryBaxisShowticksuffix_none LayoutTernaryBaxisShowticksuffix = "none" // none
)

// LayoutTernaryBaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutTernaryBaxisTickmode string

const (
    LayoutTernaryBaxisTickmode_auto LayoutTernaryBaxisTickmode = "auto" // auto
    LayoutTernaryBaxisTickmode_linear LayoutTernaryBaxisTickmode = "linear" // linear
    LayoutTernaryBaxisTickmode_array LayoutTernaryBaxisTickmode = "array" // array
)

// LayoutTernaryBaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutTernaryBaxisTicks string

const (
    LayoutTernaryBaxisTicks_outside LayoutTernaryBaxisTicks = "outside" // outside
    LayoutTernaryBaxisTicks_inside LayoutTernaryBaxisTicks = "inside" // inside
)

// LayoutTernaryCaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutTernaryCaxisExponentformat string

const (
    LayoutTernaryCaxisExponentformat_none LayoutTernaryCaxisExponentformat = "none" // none
    LayoutTernaryCaxisExponentformat_e LayoutTernaryCaxisExponentformat = "e" // e
    LayoutTernaryCaxisExponentformat_E LayoutTernaryCaxisExponentformat = "E" // E
    LayoutTernaryCaxisExponentformat_power LayoutTernaryCaxisExponentformat = "power" // power
    LayoutTernaryCaxisExponentformat_SI LayoutTernaryCaxisExponentformat = "SI" // SI
    LayoutTernaryCaxisExponentformat_B LayoutTernaryCaxisExponentformat = "B" // B
)

// LayoutTernaryCaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutTernaryCaxisLayer string

const (
    LayoutTernaryCaxisLayer_abovetraces LayoutTernaryCaxisLayer = "above traces" // above traces
    LayoutTernaryCaxisLayer_belowtraces LayoutTernaryCaxisLayer = "below traces" // below traces
)

// LayoutTernaryCaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutTernaryCaxisShowexponent string

const (
    LayoutTernaryCaxisShowexponent_all LayoutTernaryCaxisShowexponent = "all" // all
    LayoutTernaryCaxisShowexponent_first LayoutTernaryCaxisShowexponent = "first" // first
    LayoutTernaryCaxisShowexponent_last LayoutTernaryCaxisShowexponent = "last" // last
    LayoutTernaryCaxisShowexponent_none LayoutTernaryCaxisShowexponent = "none" // none
)

// LayoutTernaryCaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutTernaryCaxisShowtickprefix string

const (
    LayoutTernaryCaxisShowtickprefix_all LayoutTernaryCaxisShowtickprefix = "all" // all
    LayoutTernaryCaxisShowtickprefix_first LayoutTernaryCaxisShowtickprefix = "first" // first
    LayoutTernaryCaxisShowtickprefix_last LayoutTernaryCaxisShowtickprefix = "last" // last
    LayoutTernaryCaxisShowtickprefix_none LayoutTernaryCaxisShowtickprefix = "none" // none
)

// LayoutTernaryCaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutTernaryCaxisShowticksuffix string

const (
    LayoutTernaryCaxisShowticksuffix_all LayoutTernaryCaxisShowticksuffix = "all" // all
    LayoutTernaryCaxisShowticksuffix_first LayoutTernaryCaxisShowticksuffix = "first" // first
    LayoutTernaryCaxisShowticksuffix_last LayoutTernaryCaxisShowticksuffix = "last" // last
    LayoutTernaryCaxisShowticksuffix_none LayoutTernaryCaxisShowticksuffix = "none" // none
)

// LayoutTernaryCaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutTernaryCaxisTickmode string

const (
    LayoutTernaryCaxisTickmode_auto LayoutTernaryCaxisTickmode = "auto" // auto
    LayoutTernaryCaxisTickmode_linear LayoutTernaryCaxisTickmode = "linear" // linear
    LayoutTernaryCaxisTickmode_array LayoutTernaryCaxisTickmode = "array" // array
)

// LayoutTernaryCaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutTernaryCaxisTicks string

const (
    LayoutTernaryCaxisTicks_outside LayoutTernaryCaxisTicks = "outside" // outside
    LayoutTernaryCaxisTicks_inside LayoutTernaryCaxisTicks = "inside" // inside
)

// LayoutTitleXanchor Sets the title's horizontal alignment with respect to its x position. *left* means that the title starts at x, *right* means that the title ends at x and *center* means that the title's center is at x. *auto* divides `xref` by three and calculates the `xanchor` value automatically based on the value of `x`.
type LayoutTitleXanchor string

const (
    LayoutTitleXanchor_auto LayoutTitleXanchor = "auto" // auto
    LayoutTitleXanchor_left LayoutTitleXanchor = "left" // left
    LayoutTitleXanchor_center LayoutTitleXanchor = "center" // center
    LayoutTitleXanchor_right LayoutTitleXanchor = "right" // right
)

// LayoutTitleXref Sets the container `x` refers to. *container* spans the entire `width` of the plot. *paper* refers to the width of the plotting area only.
type LayoutTitleXref string

const (
    LayoutTitleXref_container LayoutTitleXref = "container" // container
    LayoutTitleXref_paper LayoutTitleXref = "paper" // paper
)

// LayoutTitleYanchor Sets the title's vertical alignment with respect to its y position. *top* means that the title's cap line is at y, *bottom* means that the title's baseline is at y and *middle* means that the title's midline is at y. *auto* divides `yref` by three and calculates the `yanchor` value automatically based on the value of `y`.
type LayoutTitleYanchor string

const (
    LayoutTitleYanchor_auto LayoutTitleYanchor = "auto" // auto
    LayoutTitleYanchor_top LayoutTitleYanchor = "top" // top
    LayoutTitleYanchor_middle LayoutTitleYanchor = "middle" // middle
    LayoutTitleYanchor_bottom LayoutTitleYanchor = "bottom" // bottom
)

// LayoutTitleYref Sets the container `y` refers to. *container* spans the entire `height` of the plot. *paper* refers to the height of the plotting area only.
type LayoutTitleYref string

const (
    LayoutTitleYref_container LayoutTitleYref = "container" // container
    LayoutTitleYref_paper LayoutTitleYref = "paper" // paper
)

// LayoutTransitionEasing The easing function used for the transition
type LayoutTransitionEasing string

const (
    LayoutTransitionEasing_linear LayoutTransitionEasing = "linear" // linear
    LayoutTransitionEasing_quad LayoutTransitionEasing = "quad" // quad
    LayoutTransitionEasing_cubic LayoutTransitionEasing = "cubic" // cubic
    LayoutTransitionEasing_sin LayoutTransitionEasing = "sin" // sin
    LayoutTransitionEasing_exp LayoutTransitionEasing = "exp" // exp
    LayoutTransitionEasing_circle LayoutTransitionEasing = "circle" // circle
    LayoutTransitionEasing_elastic LayoutTransitionEasing = "elastic" // elastic
    LayoutTransitionEasing_back LayoutTransitionEasing = "back" // back
    LayoutTransitionEasing_bounce LayoutTransitionEasing = "bounce" // bounce
    LayoutTransitionEasing_linear_in LayoutTransitionEasing = "linear-in" // linear-in
    LayoutTransitionEasing_quad_in LayoutTransitionEasing = "quad-in" // quad-in
    LayoutTransitionEasing_cubic_in LayoutTransitionEasing = "cubic-in" // cubic-in
    LayoutTransitionEasing_sin_in LayoutTransitionEasing = "sin-in" // sin-in
    LayoutTransitionEasing_exp_in LayoutTransitionEasing = "exp-in" // exp-in
    LayoutTransitionEasing_circle_in LayoutTransitionEasing = "circle-in" // circle-in
    LayoutTransitionEasing_elastic_in LayoutTransitionEasing = "elastic-in" // elastic-in
    LayoutTransitionEasing_back_in LayoutTransitionEasing = "back-in" // back-in
    LayoutTransitionEasing_bounce_in LayoutTransitionEasing = "bounce-in" // bounce-in
    LayoutTransitionEasing_linear_out LayoutTransitionEasing = "linear-out" // linear-out
    LayoutTransitionEasing_quad_out LayoutTransitionEasing = "quad-out" // quad-out
    LayoutTransitionEasing_cubic_out LayoutTransitionEasing = "cubic-out" // cubic-out
    LayoutTransitionEasing_sin_out LayoutTransitionEasing = "sin-out" // sin-out
    LayoutTransitionEasing_exp_out LayoutTransitionEasing = "exp-out" // exp-out
    LayoutTransitionEasing_circle_out LayoutTransitionEasing = "circle-out" // circle-out
    LayoutTransitionEasing_elastic_out LayoutTransitionEasing = "elastic-out" // elastic-out
    LayoutTransitionEasing_back_out LayoutTransitionEasing = "back-out" // back-out
    LayoutTransitionEasing_bounce_out LayoutTransitionEasing = "bounce-out" // bounce-out
    LayoutTransitionEasing_linear_in_out LayoutTransitionEasing = "linear-in-out" // linear-in-out
    LayoutTransitionEasing_quad_in_out LayoutTransitionEasing = "quad-in-out" // quad-in-out
    LayoutTransitionEasing_cubic_in_out LayoutTransitionEasing = "cubic-in-out" // cubic-in-out
    LayoutTransitionEasing_sin_in_out LayoutTransitionEasing = "sin-in-out" // sin-in-out
    LayoutTransitionEasing_exp_in_out LayoutTransitionEasing = "exp-in-out" // exp-in-out
    LayoutTransitionEasing_circle_in_out LayoutTransitionEasing = "circle-in-out" // circle-in-out
    LayoutTransitionEasing_elastic_in_out LayoutTransitionEasing = "elastic-in-out" // elastic-in-out
    LayoutTransitionEasing_back_in_out LayoutTransitionEasing = "back-in-out" // back-in-out
    LayoutTransitionEasing_bounce_in_out LayoutTransitionEasing = "bounce-in-out" // bounce-in-out
)

// LayoutTransitionOrdering Determines whether the figure's layout or traces smoothly transitions during updates that make both traces and layout change.
type LayoutTransitionOrdering string

const (
    LayoutTransitionOrdering_layoutfirst LayoutTransitionOrdering = "layout first" // layout first
    LayoutTransitionOrdering_tracesfirst LayoutTransitionOrdering = "traces first" // traces first
)

// LayoutUniformtextMode Determines how the font size for various text elements are uniformed between each trace type. If the computed text sizes were smaller than the minimum size defined by `uniformtext.minsize` using *hide* option hides the text; and using *show* option shows the text without further downscaling. Please note that if the size defined by `minsize` is greater than the font size defined by trace, then the `minsize` is used.
type LayoutUniformtextMode interface{}

const (
    LayoutUniformtextModeFalse bool = false
    LayoutUniformtextModeHide string = "hide"
    LayoutUniformtextModeShow string = "show"
)

// LayoutXaxisAnchor If set to an opposite-letter axis id (e.g. `x2`, `y`), this axis is bound to the corresponding opposite-letter axis. If set to *free*, this axis' position is determined by `position`.
type LayoutXaxisAnchor string

const (
    LayoutXaxisAnchor_free LayoutXaxisAnchor = "free" // free
    LayoutXaxisAnchor_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisAnchor = "/^x([2-9]|[1-9][0-9]+)?$/" // /^x([2-9]|[1-9][0-9]+)?$/
    LayoutXaxisAnchor_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisAnchor = "/^y([2-9]|[1-9][0-9]+)?$/" // /^y([2-9]|[1-9][0-9]+)?$/
)

// LayoutXaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutXaxisAutorange interface{}

const (
    LayoutXaxisAutorangeTrue bool = true
    LayoutXaxisAutorangeFalse bool = false
    LayoutXaxisAutorangeReversed string = "reversed"
)

// LayoutXaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutXaxisCalendar string

const (
    LayoutXaxisCalendar_gregorian LayoutXaxisCalendar = "gregorian" // gregorian
    LayoutXaxisCalendar_chinese LayoutXaxisCalendar = "chinese" // chinese
    LayoutXaxisCalendar_coptic LayoutXaxisCalendar = "coptic" // coptic
    LayoutXaxisCalendar_discworld LayoutXaxisCalendar = "discworld" // discworld
    LayoutXaxisCalendar_ethiopian LayoutXaxisCalendar = "ethiopian" // ethiopian
    LayoutXaxisCalendar_hebrew LayoutXaxisCalendar = "hebrew" // hebrew
    LayoutXaxisCalendar_islamic LayoutXaxisCalendar = "islamic" // islamic
    LayoutXaxisCalendar_julian LayoutXaxisCalendar = "julian" // julian
    LayoutXaxisCalendar_mayan LayoutXaxisCalendar = "mayan" // mayan
    LayoutXaxisCalendar_nanakshahi LayoutXaxisCalendar = "nanakshahi" // nanakshahi
    LayoutXaxisCalendar_nepali LayoutXaxisCalendar = "nepali" // nepali
    LayoutXaxisCalendar_persian LayoutXaxisCalendar = "persian" // persian
    LayoutXaxisCalendar_jalali LayoutXaxisCalendar = "jalali" // jalali
    LayoutXaxisCalendar_taiwan LayoutXaxisCalendar = "taiwan" // taiwan
    LayoutXaxisCalendar_thai LayoutXaxisCalendar = "thai" // thai
    LayoutXaxisCalendar_ummalqura LayoutXaxisCalendar = "ummalqura" // ummalqura
)

// LayoutXaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutXaxisCategoryorder string

const (
    LayoutXaxisCategoryorder_trace LayoutXaxisCategoryorder = "trace" // trace
    LayoutXaxisCategoryorder_categoryascending LayoutXaxisCategoryorder = "category ascending" // category ascending
    LayoutXaxisCategoryorder_categorydescending LayoutXaxisCategoryorder = "category descending" // category descending
    LayoutXaxisCategoryorder_array LayoutXaxisCategoryorder = "array" // array
    LayoutXaxisCategoryorder_totalascending LayoutXaxisCategoryorder = "total ascending" // total ascending
    LayoutXaxisCategoryorder_totaldescending LayoutXaxisCategoryorder = "total descending" // total descending
    LayoutXaxisCategoryorder_minascending LayoutXaxisCategoryorder = "min ascending" // min ascending
    LayoutXaxisCategoryorder_mindescending LayoutXaxisCategoryorder = "min descending" // min descending
    LayoutXaxisCategoryorder_maxascending LayoutXaxisCategoryorder = "max ascending" // max ascending
    LayoutXaxisCategoryorder_maxdescending LayoutXaxisCategoryorder = "max descending" // max descending
    LayoutXaxisCategoryorder_sumascending LayoutXaxisCategoryorder = "sum ascending" // sum ascending
    LayoutXaxisCategoryorder_sumdescending LayoutXaxisCategoryorder = "sum descending" // sum descending
    LayoutXaxisCategoryorder_meanascending LayoutXaxisCategoryorder = "mean ascending" // mean ascending
    LayoutXaxisCategoryorder_meandescending LayoutXaxisCategoryorder = "mean descending" // mean descending
    LayoutXaxisCategoryorder_medianascending LayoutXaxisCategoryorder = "median ascending" // median ascending
    LayoutXaxisCategoryorder_mediandescending LayoutXaxisCategoryorder = "median descending" // median descending
)

// LayoutXaxisConstrain If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines how that happens: by increasing the *range* (default), or by decreasing the *domain*.
type LayoutXaxisConstrain string

const (
    LayoutXaxisConstrain_range LayoutXaxisConstrain = "range" // range
    LayoutXaxisConstrain_domain LayoutXaxisConstrain = "domain" // domain
)

// LayoutXaxisConstraintoward If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines which direction we push the originally specified plot area. Options are *left*, *center* (default), and *right* for x axes, and *top*, *middle* (default), and *bottom* for y axes.
type LayoutXaxisConstraintoward string

const (
    LayoutXaxisConstraintoward_left LayoutXaxisConstraintoward = "left" // left
    LayoutXaxisConstraintoward_center LayoutXaxisConstraintoward = "center" // center
    LayoutXaxisConstraintoward_right LayoutXaxisConstraintoward = "right" // right
    LayoutXaxisConstraintoward_top LayoutXaxisConstraintoward = "top" // top
    LayoutXaxisConstraintoward_middle LayoutXaxisConstraintoward = "middle" // middle
    LayoutXaxisConstraintoward_bottom LayoutXaxisConstraintoward = "bottom" // bottom
)

// LayoutXaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutXaxisExponentformat string

const (
    LayoutXaxisExponentformat_none LayoutXaxisExponentformat = "none" // none
    LayoutXaxisExponentformat_e LayoutXaxisExponentformat = "e" // e
    LayoutXaxisExponentformat_E LayoutXaxisExponentformat = "E" // E
    LayoutXaxisExponentformat_power LayoutXaxisExponentformat = "power" // power
    LayoutXaxisExponentformat_SI LayoutXaxisExponentformat = "SI" // SI
    LayoutXaxisExponentformat_B LayoutXaxisExponentformat = "B" // B
)

// LayoutXaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutXaxisLayer string

const (
    LayoutXaxisLayer_abovetraces LayoutXaxisLayer = "above traces" // above traces
    LayoutXaxisLayer_belowtraces LayoutXaxisLayer = "below traces" // below traces
)

// LayoutXaxisMatches If set to another axis id (e.g. `x2`, `y`), the range of this axis will match the range of the corresponding axis in data-coordinates space. Moreover, matching axes share auto-range values, category lists and histogram auto-bins. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden. Moreover, note that matching axes must have the same `type`.
type LayoutXaxisMatches string

const (
    LayoutXaxisMatches_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisMatches = "/^x([2-9]|[1-9][0-9]+)?$/" // /^x([2-9]|[1-9][0-9]+)?$/
    LayoutXaxisMatches_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisMatches = "/^y([2-9]|[1-9][0-9]+)?$/" // /^y([2-9]|[1-9][0-9]+)?$/
)

// LayoutXaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutXaxisMirror interface{}

const (
    LayoutXaxisMirrorTrue bool = true
    LayoutXaxisMirrorTicks string = "ticks"
    LayoutXaxisMirrorFalse bool = false
    LayoutXaxisMirrorAll string = "all"
    LayoutXaxisMirrorAllticks string = "allticks"
)

// LayoutXaxisOverlaying If set a same-letter axis id, this axis is overlaid on top of the corresponding same-letter axis, with traces and axes visible for both axes. If *false*, this axis does not overlay any same-letter axes. In this case, for axes with overlapping domains only the highest-numbered axis will be visible.
type LayoutXaxisOverlaying string

const (
    LayoutXaxisOverlaying_free LayoutXaxisOverlaying = "free" // free
    LayoutXaxisOverlaying_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisOverlaying = "/^x([2-9]|[1-9][0-9]+)?$/" // /^x([2-9]|[1-9][0-9]+)?$/
    LayoutXaxisOverlaying_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisOverlaying = "/^y([2-9]|[1-9][0-9]+)?$/" // /^y([2-9]|[1-9][0-9]+)?$/
)

// LayoutXaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutXaxisRangemode string

const (
    LayoutXaxisRangemode_normal LayoutXaxisRangemode = "normal" // normal
    LayoutXaxisRangemode_tozero LayoutXaxisRangemode = "tozero" // tozero
    LayoutXaxisRangemode_nonnegative LayoutXaxisRangemode = "nonnegative" // nonnegative
)

// LayoutXaxisRangeselectorXanchor Sets the range selector's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the range selector.
type LayoutXaxisRangeselectorXanchor string

const (
    LayoutXaxisRangeselectorXanchor_auto LayoutXaxisRangeselectorXanchor = "auto" // auto
    LayoutXaxisRangeselectorXanchor_left LayoutXaxisRangeselectorXanchor = "left" // left
    LayoutXaxisRangeselectorXanchor_center LayoutXaxisRangeselectorXanchor = "center" // center
    LayoutXaxisRangeselectorXanchor_right LayoutXaxisRangeselectorXanchor = "right" // right
)

// LayoutXaxisRangeselectorYanchor Sets the range selector's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the range selector.
type LayoutXaxisRangeselectorYanchor string

const (
    LayoutXaxisRangeselectorYanchor_auto LayoutXaxisRangeselectorYanchor = "auto" // auto
    LayoutXaxisRangeselectorYanchor_top LayoutXaxisRangeselectorYanchor = "top" // top
    LayoutXaxisRangeselectorYanchor_middle LayoutXaxisRangeselectorYanchor = "middle" // middle
    LayoutXaxisRangeselectorYanchor_bottom LayoutXaxisRangeselectorYanchor = "bottom" // bottom
)

// LayoutXaxisRangesliderYaxisRangemode Determines whether or not the range of this axis in the rangeslider use the same value than in the main plot when zooming in/out. If *auto*, the autorange will be used. If *fixed*, the `range` is used. If *match*, the current range of the corresponding y-axis on the main subplot is used.
type LayoutXaxisRangesliderYaxisRangemode string

const (
    LayoutXaxisRangesliderYaxisRangemode_auto LayoutXaxisRangesliderYaxisRangemode = "auto" // auto
    LayoutXaxisRangesliderYaxisRangemode_fixed LayoutXaxisRangesliderYaxisRangemode = "fixed" // fixed
    LayoutXaxisRangesliderYaxisRangemode_match LayoutXaxisRangesliderYaxisRangemode = "match" // match
)

// LayoutXaxisScaleanchor If set to another axis id (e.g. `x2`, `y`), the range of this axis changes together with the range of the corresponding axis such that the scale of pixels per unit is in a constant ratio. Both axes are still zoomable, but when you zoom one, the other will zoom the same amount, keeping a fixed midpoint. `constrain` and `constraintoward` determine how we enforce the constraint. You can chain these, ie `yaxis: {scaleanchor: *x*}, xaxis2: {scaleanchor: *y*}` but you can only link axes of the same `type`. The linked axis can have the opposite letter (to constrain the aspect ratio) or the same letter (to match scales across subplots). Loops (`yaxis: {scaleanchor: *x*}, xaxis: {scaleanchor: *y*}` or longer) are redundant and the last constraint encountered will be ignored to avoid possible inconsistent constraints via `scaleratio`. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden.
type LayoutXaxisScaleanchor string

const (
    LayoutXaxisScaleanchor_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisScaleanchor = "/^x([2-9]|[1-9][0-9]+)?$/" // /^x([2-9]|[1-9][0-9]+)?$/
    LayoutXaxisScaleanchor_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutXaxisScaleanchor = "/^y([2-9]|[1-9][0-9]+)?$/" // /^y([2-9]|[1-9][0-9]+)?$/
)

// LayoutXaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutXaxisShowexponent string

const (
    LayoutXaxisShowexponent_all LayoutXaxisShowexponent = "all" // all
    LayoutXaxisShowexponent_first LayoutXaxisShowexponent = "first" // first
    LayoutXaxisShowexponent_last LayoutXaxisShowexponent = "last" // last
    LayoutXaxisShowexponent_none LayoutXaxisShowexponent = "none" // none
)

// LayoutXaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutXaxisShowtickprefix string

const (
    LayoutXaxisShowtickprefix_all LayoutXaxisShowtickprefix = "all" // all
    LayoutXaxisShowtickprefix_first LayoutXaxisShowtickprefix = "first" // first
    LayoutXaxisShowtickprefix_last LayoutXaxisShowtickprefix = "last" // last
    LayoutXaxisShowtickprefix_none LayoutXaxisShowtickprefix = "none" // none
)

// LayoutXaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutXaxisShowticksuffix string

const (
    LayoutXaxisShowticksuffix_all LayoutXaxisShowticksuffix = "all" // all
    LayoutXaxisShowticksuffix_first LayoutXaxisShowticksuffix = "first" // first
    LayoutXaxisShowticksuffix_last LayoutXaxisShowticksuffix = "last" // last
    LayoutXaxisShowticksuffix_none LayoutXaxisShowticksuffix = "none" // none
)

// LayoutXaxisSide Determines whether a x (y) axis is positioned at the *bottom* (*left*) or *top* (*right*) of the plotting area.
type LayoutXaxisSide string

const (
    LayoutXaxisSide_top LayoutXaxisSide = "top" // top
    LayoutXaxisSide_bottom LayoutXaxisSide = "bottom" // bottom
    LayoutXaxisSide_left LayoutXaxisSide = "left" // left
    LayoutXaxisSide_right LayoutXaxisSide = "right" // right
)

// LayoutXaxisSpikesnap Determines whether spikelines are stuck to the cursor or to the closest datapoints.
type LayoutXaxisSpikesnap string

const (
    LayoutXaxisSpikesnap_data LayoutXaxisSpikesnap = "data" // data
    LayoutXaxisSpikesnap_cursor LayoutXaxisSpikesnap = "cursor" // cursor
    LayoutXaxisSpikesnap_hovereddata LayoutXaxisSpikesnap = "hovered data" // hovered data
)

// LayoutXaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutXaxisTickmode string

const (
    LayoutXaxisTickmode_auto LayoutXaxisTickmode = "auto" // auto
    LayoutXaxisTickmode_linear LayoutXaxisTickmode = "linear" // linear
    LayoutXaxisTickmode_array LayoutXaxisTickmode = "array" // array
)

// LayoutXaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutXaxisTicks string

const (
    LayoutXaxisTicks_outside LayoutXaxisTicks = "outside" // outside
    LayoutXaxisTicks_inside LayoutXaxisTicks = "inside" // inside
)

// LayoutXaxisTickson Determines where ticks and grid lines are drawn with respect to their corresponding tick labels. Only has an effect for axes of `type` *category* or *multicategory*. When set to *boundaries*, ticks and grid lines are drawn half a category to the left/bottom of labels.
type LayoutXaxisTickson string

const (
    LayoutXaxisTickson_labels LayoutXaxisTickson = "labels" // labels
    LayoutXaxisTickson_boundaries LayoutXaxisTickson = "boundaries" // boundaries
)

// LayoutXaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutXaxisType string

const (
    LayoutXaxisType__ LayoutXaxisType = "-" // -
    LayoutXaxisType_linear LayoutXaxisType = "linear" // linear
    LayoutXaxisType_log LayoutXaxisType = "log" // log
    LayoutXaxisType_date LayoutXaxisType = "date" // date
    LayoutXaxisType_category LayoutXaxisType = "category" // category
    LayoutXaxisType_multicategory LayoutXaxisType = "multicategory" // multicategory
)

// LayoutYaxisAnchor If set to an opposite-letter axis id (e.g. `x2`, `y`), this axis is bound to the corresponding opposite-letter axis. If set to *free*, this axis' position is determined by `position`.
type LayoutYaxisAnchor string

const (
    LayoutYaxisAnchor_free LayoutYaxisAnchor = "free" // free
    LayoutYaxisAnchor_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisAnchor = "/^x([2-9]|[1-9][0-9]+)?$/" // /^x([2-9]|[1-9][0-9]+)?$/
    LayoutYaxisAnchor_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisAnchor = "/^y([2-9]|[1-9][0-9]+)?$/" // /^y([2-9]|[1-9][0-9]+)?$/
)

// LayoutYaxisAutorange Determines whether or not the range of this axis is computed in relation to the input data. See `rangemode` for more info. If `range` is provided, then `autorange` is set to *false*.
type LayoutYaxisAutorange interface{}

const (
    LayoutYaxisAutorangeTrue bool = true
    LayoutYaxisAutorangeFalse bool = false
    LayoutYaxisAutorangeReversed string = "reversed"
)

// LayoutYaxisCalendar Sets the calendar system to use for `range` and `tick0` if this is a date axis. This does not set the calendar for interpreting data on this axis, that's specified in the trace or via the global `layout.calendar`
type LayoutYaxisCalendar string

const (
    LayoutYaxisCalendar_gregorian LayoutYaxisCalendar = "gregorian" // gregorian
    LayoutYaxisCalendar_chinese LayoutYaxisCalendar = "chinese" // chinese
    LayoutYaxisCalendar_coptic LayoutYaxisCalendar = "coptic" // coptic
    LayoutYaxisCalendar_discworld LayoutYaxisCalendar = "discworld" // discworld
    LayoutYaxisCalendar_ethiopian LayoutYaxisCalendar = "ethiopian" // ethiopian
    LayoutYaxisCalendar_hebrew LayoutYaxisCalendar = "hebrew" // hebrew
    LayoutYaxisCalendar_islamic LayoutYaxisCalendar = "islamic" // islamic
    LayoutYaxisCalendar_julian LayoutYaxisCalendar = "julian" // julian
    LayoutYaxisCalendar_mayan LayoutYaxisCalendar = "mayan" // mayan
    LayoutYaxisCalendar_nanakshahi LayoutYaxisCalendar = "nanakshahi" // nanakshahi
    LayoutYaxisCalendar_nepali LayoutYaxisCalendar = "nepali" // nepali
    LayoutYaxisCalendar_persian LayoutYaxisCalendar = "persian" // persian
    LayoutYaxisCalendar_jalali LayoutYaxisCalendar = "jalali" // jalali
    LayoutYaxisCalendar_taiwan LayoutYaxisCalendar = "taiwan" // taiwan
    LayoutYaxisCalendar_thai LayoutYaxisCalendar = "thai" // thai
    LayoutYaxisCalendar_ummalqura LayoutYaxisCalendar = "ummalqura" // ummalqura
)

// LayoutYaxisCategoryorder Specifies the ordering logic for the case of categorical variables. By default, plotly uses *trace*, which specifies the order that is present in the data supplied. Set `categoryorder` to *category ascending* or *category descending* if order should be determined by the alphanumerical order of the category names. Set `categoryorder` to *array* to derive the ordering from the attribute `categoryarray`. If a category is not found in the `categoryarray` array, the sorting behavior for that attribute will be identical to the *trace* mode. The unspecified categories will follow the categories in `categoryarray`. Set `categoryorder` to *total ascending* or *total descending* if order should be determined by the numerical order of the values. Similarly, the order can be determined by the min, max, sum, mean or median of all the values.
type LayoutYaxisCategoryorder string

const (
    LayoutYaxisCategoryorder_trace LayoutYaxisCategoryorder = "trace" // trace
    LayoutYaxisCategoryorder_categoryascending LayoutYaxisCategoryorder = "category ascending" // category ascending
    LayoutYaxisCategoryorder_categorydescending LayoutYaxisCategoryorder = "category descending" // category descending
    LayoutYaxisCategoryorder_array LayoutYaxisCategoryorder = "array" // array
    LayoutYaxisCategoryorder_totalascending LayoutYaxisCategoryorder = "total ascending" // total ascending
    LayoutYaxisCategoryorder_totaldescending LayoutYaxisCategoryorder = "total descending" // total descending
    LayoutYaxisCategoryorder_minascending LayoutYaxisCategoryorder = "min ascending" // min ascending
    LayoutYaxisCategoryorder_mindescending LayoutYaxisCategoryorder = "min descending" // min descending
    LayoutYaxisCategoryorder_maxascending LayoutYaxisCategoryorder = "max ascending" // max ascending
    LayoutYaxisCategoryorder_maxdescending LayoutYaxisCategoryorder = "max descending" // max descending
    LayoutYaxisCategoryorder_sumascending LayoutYaxisCategoryorder = "sum ascending" // sum ascending
    LayoutYaxisCategoryorder_sumdescending LayoutYaxisCategoryorder = "sum descending" // sum descending
    LayoutYaxisCategoryorder_meanascending LayoutYaxisCategoryorder = "mean ascending" // mean ascending
    LayoutYaxisCategoryorder_meandescending LayoutYaxisCategoryorder = "mean descending" // mean descending
    LayoutYaxisCategoryorder_medianascending LayoutYaxisCategoryorder = "median ascending" // median ascending
    LayoutYaxisCategoryorder_mediandescending LayoutYaxisCategoryorder = "median descending" // median descending
)

// LayoutYaxisConstrain If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines how that happens: by increasing the *range* (default), or by decreasing the *domain*.
type LayoutYaxisConstrain string

const (
    LayoutYaxisConstrain_range LayoutYaxisConstrain = "range" // range
    LayoutYaxisConstrain_domain LayoutYaxisConstrain = "domain" // domain
)

// LayoutYaxisConstraintoward If this axis needs to be compressed (either due to its own `scaleanchor` and `scaleratio` or those of the other axis), determines which direction we push the originally specified plot area. Options are *left*, *center* (default), and *right* for x axes, and *top*, *middle* (default), and *bottom* for y axes.
type LayoutYaxisConstraintoward string

const (
    LayoutYaxisConstraintoward_left LayoutYaxisConstraintoward = "left" // left
    LayoutYaxisConstraintoward_center LayoutYaxisConstraintoward = "center" // center
    LayoutYaxisConstraintoward_right LayoutYaxisConstraintoward = "right" // right
    LayoutYaxisConstraintoward_top LayoutYaxisConstraintoward = "top" // top
    LayoutYaxisConstraintoward_middle LayoutYaxisConstraintoward = "middle" // middle
    LayoutYaxisConstraintoward_bottom LayoutYaxisConstraintoward = "bottom" // bottom
)

// LayoutYaxisExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type LayoutYaxisExponentformat string

const (
    LayoutYaxisExponentformat_none LayoutYaxisExponentformat = "none" // none
    LayoutYaxisExponentformat_e LayoutYaxisExponentformat = "e" // e
    LayoutYaxisExponentformat_E LayoutYaxisExponentformat = "E" // E
    LayoutYaxisExponentformat_power LayoutYaxisExponentformat = "power" // power
    LayoutYaxisExponentformat_SI LayoutYaxisExponentformat = "SI" // SI
    LayoutYaxisExponentformat_B LayoutYaxisExponentformat = "B" // B
)

// LayoutYaxisLayer Sets the layer on which this axis is displayed. If *above traces*, this axis is displayed above all the subplot's traces If *below traces*, this axis is displayed below all the subplot's traces, but above the grid lines. Useful when used together with scatter-like traces with `cliponaxis` set to *false* to show markers and/or text nodes above this axis.
type LayoutYaxisLayer string

const (
    LayoutYaxisLayer_abovetraces LayoutYaxisLayer = "above traces" // above traces
    LayoutYaxisLayer_belowtraces LayoutYaxisLayer = "below traces" // below traces
)

// LayoutYaxisMatches If set to another axis id (e.g. `x2`, `y`), the range of this axis will match the range of the corresponding axis in data-coordinates space. Moreover, matching axes share auto-range values, category lists and histogram auto-bins. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden. Moreover, note that matching axes must have the same `type`.
type LayoutYaxisMatches string

const (
    LayoutYaxisMatches_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisMatches = "/^x([2-9]|[1-9][0-9]+)?$/" // /^x([2-9]|[1-9][0-9]+)?$/
    LayoutYaxisMatches_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisMatches = "/^y([2-9]|[1-9][0-9]+)?$/" // /^y([2-9]|[1-9][0-9]+)?$/
)

// LayoutYaxisMirror Determines if the axis lines or/and ticks are mirrored to the opposite side of the plotting area. If *true*, the axis lines are mirrored. If *ticks*, the axis lines and ticks are mirrored. If *false*, mirroring is disable. If *all*, axis lines are mirrored on all shared-axes subplots. If *allticks*, axis lines and ticks are mirrored on all shared-axes subplots.
type LayoutYaxisMirror interface{}

const (
    LayoutYaxisMirrorTrue bool = true
    LayoutYaxisMirrorTicks string = "ticks"
    LayoutYaxisMirrorFalse bool = false
    LayoutYaxisMirrorAll string = "all"
    LayoutYaxisMirrorAllticks string = "allticks"
)

// LayoutYaxisOverlaying If set a same-letter axis id, this axis is overlaid on top of the corresponding same-letter axis, with traces and axes visible for both axes. If *false*, this axis does not overlay any same-letter axes. In this case, for axes with overlapping domains only the highest-numbered axis will be visible.
type LayoutYaxisOverlaying string

const (
    LayoutYaxisOverlaying_free LayoutYaxisOverlaying = "free" // free
    LayoutYaxisOverlaying_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisOverlaying = "/^x([2-9]|[1-9][0-9]+)?$/" // /^x([2-9]|[1-9][0-9]+)?$/
    LayoutYaxisOverlaying_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisOverlaying = "/^y([2-9]|[1-9][0-9]+)?$/" // /^y([2-9]|[1-9][0-9]+)?$/
)

// LayoutYaxisRangemode If *normal*, the range is computed in relation to the extrema of the input data. If *tozero*`, the range extends to 0, regardless of the input data If *nonnegative*, the range is non-negative, regardless of the input data. Applies only to linear axes.
type LayoutYaxisRangemode string

const (
    LayoutYaxisRangemode_normal LayoutYaxisRangemode = "normal" // normal
    LayoutYaxisRangemode_tozero LayoutYaxisRangemode = "tozero" // tozero
    LayoutYaxisRangemode_nonnegative LayoutYaxisRangemode = "nonnegative" // nonnegative
)

// LayoutYaxisScaleanchor If set to another axis id (e.g. `x2`, `y`), the range of this axis changes together with the range of the corresponding axis such that the scale of pixels per unit is in a constant ratio. Both axes are still zoomable, but when you zoom one, the other will zoom the same amount, keeping a fixed midpoint. `constrain` and `constraintoward` determine how we enforce the constraint. You can chain these, ie `yaxis: {scaleanchor: *x*}, xaxis2: {scaleanchor: *y*}` but you can only link axes of the same `type`. The linked axis can have the opposite letter (to constrain the aspect ratio) or the same letter (to match scales across subplots). Loops (`yaxis: {scaleanchor: *x*}, xaxis: {scaleanchor: *y*}` or longer) are redundant and the last constraint encountered will be ignored to avoid possible inconsistent constraints via `scaleratio`. Note that setting axes simultaneously in both a `scaleanchor` and a `matches` constraint is currently forbidden.
type LayoutYaxisScaleanchor string

const (
    LayoutYaxisScaleanchor_slash_cape_x_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisScaleanchor = "/^x([2-9]|[1-9][0-9]+)?$/" // /^x([2-9]|[1-9][0-9]+)?$/
    LayoutYaxisScaleanchor_slash_cape_y_lpar__lbracket_2_9_rbracket_or_lbracket_1_9_rbracket__lbracket_0_9_rbracket_plus_rpar__question__dollar_slash LayoutYaxisScaleanchor = "/^y([2-9]|[1-9][0-9]+)?$/" // /^y([2-9]|[1-9][0-9]+)?$/
)

// LayoutYaxisShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type LayoutYaxisShowexponent string

const (
    LayoutYaxisShowexponent_all LayoutYaxisShowexponent = "all" // all
    LayoutYaxisShowexponent_first LayoutYaxisShowexponent = "first" // first
    LayoutYaxisShowexponent_last LayoutYaxisShowexponent = "last" // last
    LayoutYaxisShowexponent_none LayoutYaxisShowexponent = "none" // none
)

// LayoutYaxisShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type LayoutYaxisShowtickprefix string

const (
    LayoutYaxisShowtickprefix_all LayoutYaxisShowtickprefix = "all" // all
    LayoutYaxisShowtickprefix_first LayoutYaxisShowtickprefix = "first" // first
    LayoutYaxisShowtickprefix_last LayoutYaxisShowtickprefix = "last" // last
    LayoutYaxisShowtickprefix_none LayoutYaxisShowtickprefix = "none" // none
)

// LayoutYaxisShowticksuffix Same as `showtickprefix` but for tick suffixes.
type LayoutYaxisShowticksuffix string

const (
    LayoutYaxisShowticksuffix_all LayoutYaxisShowticksuffix = "all" // all
    LayoutYaxisShowticksuffix_first LayoutYaxisShowticksuffix = "first" // first
    LayoutYaxisShowticksuffix_last LayoutYaxisShowticksuffix = "last" // last
    LayoutYaxisShowticksuffix_none LayoutYaxisShowticksuffix = "none" // none
)

// LayoutYaxisSide Determines whether a x (y) axis is positioned at the *bottom* (*left*) or *top* (*right*) of the plotting area.
type LayoutYaxisSide string

const (
    LayoutYaxisSide_top LayoutYaxisSide = "top" // top
    LayoutYaxisSide_bottom LayoutYaxisSide = "bottom" // bottom
    LayoutYaxisSide_left LayoutYaxisSide = "left" // left
    LayoutYaxisSide_right LayoutYaxisSide = "right" // right
)

// LayoutYaxisSpikesnap Determines whether spikelines are stuck to the cursor or to the closest datapoints.
type LayoutYaxisSpikesnap string

const (
    LayoutYaxisSpikesnap_data LayoutYaxisSpikesnap = "data" // data
    LayoutYaxisSpikesnap_cursor LayoutYaxisSpikesnap = "cursor" // cursor
    LayoutYaxisSpikesnap_hovereddata LayoutYaxisSpikesnap = "hovered data" // hovered data
)

// LayoutYaxisTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type LayoutYaxisTickmode string

const (
    LayoutYaxisTickmode_auto LayoutYaxisTickmode = "auto" // auto
    LayoutYaxisTickmode_linear LayoutYaxisTickmode = "linear" // linear
    LayoutYaxisTickmode_array LayoutYaxisTickmode = "array" // array
)

// LayoutYaxisTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type LayoutYaxisTicks string

const (
    LayoutYaxisTicks_outside LayoutYaxisTicks = "outside" // outside
    LayoutYaxisTicks_inside LayoutYaxisTicks = "inside" // inside
)

// LayoutYaxisTickson Determines where ticks and grid lines are drawn with respect to their corresponding tick labels. Only has an effect for axes of `type` *category* or *multicategory*. When set to *boundaries*, ticks and grid lines are drawn half a category to the left/bottom of labels.
type LayoutYaxisTickson string

const (
    LayoutYaxisTickson_labels LayoutYaxisTickson = "labels" // labels
    LayoutYaxisTickson_boundaries LayoutYaxisTickson = "boundaries" // boundaries
)

// LayoutYaxisType Sets the axis type. By default, plotly attempts to determined the axis type by looking into the data of the traces that referenced the axis in question.
type LayoutYaxisType string

const (
    LayoutYaxisType__ LayoutYaxisType = "-" // -
    LayoutYaxisType_linear LayoutYaxisType = "linear" // linear
    LayoutYaxisType_log LayoutYaxisType = "log" // log
    LayoutYaxisType_date LayoutYaxisType = "date" // date
    LayoutYaxisType_category LayoutYaxisType = "category" // category
    LayoutYaxisType_multicategory LayoutYaxisType = "multicategory" // multicategory
)

// Mesh3dColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type Mesh3dColorbarExponentformat string

const (
    Mesh3dColorbarExponentformat_none Mesh3dColorbarExponentformat = "none" // none
    Mesh3dColorbarExponentformat_e Mesh3dColorbarExponentformat = "e" // e
    Mesh3dColorbarExponentformat_E Mesh3dColorbarExponentformat = "E" // E
    Mesh3dColorbarExponentformat_power Mesh3dColorbarExponentformat = "power" // power
    Mesh3dColorbarExponentformat_SI Mesh3dColorbarExponentformat = "SI" // SI
    Mesh3dColorbarExponentformat_B Mesh3dColorbarExponentformat = "B" // B
)

// Mesh3dColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type Mesh3dColorbarLenmode string

const (
    Mesh3dColorbarLenmode_fraction Mesh3dColorbarLenmode = "fraction" // fraction
    Mesh3dColorbarLenmode_pixels Mesh3dColorbarLenmode = "pixels" // pixels
)

// Mesh3dColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type Mesh3dColorbarShowexponent string

const (
    Mesh3dColorbarShowexponent_all Mesh3dColorbarShowexponent = "all" // all
    Mesh3dColorbarShowexponent_first Mesh3dColorbarShowexponent = "first" // first
    Mesh3dColorbarShowexponent_last Mesh3dColorbarShowexponent = "last" // last
    Mesh3dColorbarShowexponent_none Mesh3dColorbarShowexponent = "none" // none
)

// Mesh3dColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type Mesh3dColorbarShowtickprefix string

const (
    Mesh3dColorbarShowtickprefix_all Mesh3dColorbarShowtickprefix = "all" // all
    Mesh3dColorbarShowtickprefix_first Mesh3dColorbarShowtickprefix = "first" // first
    Mesh3dColorbarShowtickprefix_last Mesh3dColorbarShowtickprefix = "last" // last
    Mesh3dColorbarShowtickprefix_none Mesh3dColorbarShowtickprefix = "none" // none
)

// Mesh3dColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type Mesh3dColorbarShowticksuffix string

const (
    Mesh3dColorbarShowticksuffix_all Mesh3dColorbarShowticksuffix = "all" // all
    Mesh3dColorbarShowticksuffix_first Mesh3dColorbarShowticksuffix = "first" // first
    Mesh3dColorbarShowticksuffix_last Mesh3dColorbarShowticksuffix = "last" // last
    Mesh3dColorbarShowticksuffix_none Mesh3dColorbarShowticksuffix = "none" // none
)

// Mesh3dColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type Mesh3dColorbarThicknessmode string

const (
    Mesh3dColorbarThicknessmode_fraction Mesh3dColorbarThicknessmode = "fraction" // fraction
    Mesh3dColorbarThicknessmode_pixels Mesh3dColorbarThicknessmode = "pixels" // pixels
)

// Mesh3dColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type Mesh3dColorbarTickmode string

const (
    Mesh3dColorbarTickmode_auto Mesh3dColorbarTickmode = "auto" // auto
    Mesh3dColorbarTickmode_linear Mesh3dColorbarTickmode = "linear" // linear
    Mesh3dColorbarTickmode_array Mesh3dColorbarTickmode = "array" // array
)

// Mesh3dColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type Mesh3dColorbarTicks string

const (
    Mesh3dColorbarTicks_outside Mesh3dColorbarTicks = "outside" // outside
    Mesh3dColorbarTicks_inside Mesh3dColorbarTicks = "inside" // inside
)

// Mesh3dColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type Mesh3dColorbarTitleSide string

const (
    Mesh3dColorbarTitleSide_right Mesh3dColorbarTitleSide = "right" // right
    Mesh3dColorbarTitleSide_top Mesh3dColorbarTitleSide = "top" // top
    Mesh3dColorbarTitleSide_bottom Mesh3dColorbarTitleSide = "bottom" // bottom
)

// Mesh3dColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type Mesh3dColorbarXanchor string

const (
    Mesh3dColorbarXanchor_left Mesh3dColorbarXanchor = "left" // left
    Mesh3dColorbarXanchor_center Mesh3dColorbarXanchor = "center" // center
    Mesh3dColorbarXanchor_right Mesh3dColorbarXanchor = "right" // right
)

// Mesh3dColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type Mesh3dColorbarYanchor string

const (
    Mesh3dColorbarYanchor_top Mesh3dColorbarYanchor = "top" // top
    Mesh3dColorbarYanchor_middle Mesh3dColorbarYanchor = "middle" // middle
    Mesh3dColorbarYanchor_bottom Mesh3dColorbarYanchor = "bottom" // bottom
)

// Mesh3dDelaunayaxis Sets the Delaunay axis, which is the axis that is perpendicular to the surface of the Delaunay triangulation. It has an effect if `i`, `j`, `k` are not provided and `alphahull` is set to indicate Delaunay triangulation.
type Mesh3dDelaunayaxis string

const (
    Mesh3dDelaunayaxis_x Mesh3dDelaunayaxis = "x" // x
    Mesh3dDelaunayaxis_y Mesh3dDelaunayaxis = "y" // y
    Mesh3dDelaunayaxis_z Mesh3dDelaunayaxis = "z" // z
)

// Mesh3dHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type Mesh3dHoverlabelAlign string

const (
    Mesh3dHoverlabelAlign_left Mesh3dHoverlabelAlign = "left" // left
    Mesh3dHoverlabelAlign_right Mesh3dHoverlabelAlign = "right" // right
    Mesh3dHoverlabelAlign_auto Mesh3dHoverlabelAlign = "auto" // auto
)

// Mesh3dIntensitymode Determines the source of `intensity` values.
type Mesh3dIntensitymode string

const (
    Mesh3dIntensitymode_vertex Mesh3dIntensitymode = "vertex" // vertex
    Mesh3dIntensitymode_cell Mesh3dIntensitymode = "cell" // cell
)

// Mesh3dVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type Mesh3dVisible interface{}

const (
    Mesh3dVisibleTrue bool = true
    Mesh3dVisibleFalse bool = false
    Mesh3dVisibleLegendonly string = "legendonly"
)

// Mesh3dXcalendar Sets the calendar system to use with `x` date data.
type Mesh3dXcalendar string

const (
    Mesh3dXcalendar_gregorian Mesh3dXcalendar = "gregorian" // gregorian
    Mesh3dXcalendar_chinese Mesh3dXcalendar = "chinese" // chinese
    Mesh3dXcalendar_coptic Mesh3dXcalendar = "coptic" // coptic
    Mesh3dXcalendar_discworld Mesh3dXcalendar = "discworld" // discworld
    Mesh3dXcalendar_ethiopian Mesh3dXcalendar = "ethiopian" // ethiopian
    Mesh3dXcalendar_hebrew Mesh3dXcalendar = "hebrew" // hebrew
    Mesh3dXcalendar_islamic Mesh3dXcalendar = "islamic" // islamic
    Mesh3dXcalendar_julian Mesh3dXcalendar = "julian" // julian
    Mesh3dXcalendar_mayan Mesh3dXcalendar = "mayan" // mayan
    Mesh3dXcalendar_nanakshahi Mesh3dXcalendar = "nanakshahi" // nanakshahi
    Mesh3dXcalendar_nepali Mesh3dXcalendar = "nepali" // nepali
    Mesh3dXcalendar_persian Mesh3dXcalendar = "persian" // persian
    Mesh3dXcalendar_jalali Mesh3dXcalendar = "jalali" // jalali
    Mesh3dXcalendar_taiwan Mesh3dXcalendar = "taiwan" // taiwan
    Mesh3dXcalendar_thai Mesh3dXcalendar = "thai" // thai
    Mesh3dXcalendar_ummalqura Mesh3dXcalendar = "ummalqura" // ummalqura
)

// Mesh3dYcalendar Sets the calendar system to use with `y` date data.
type Mesh3dYcalendar string

const (
    Mesh3dYcalendar_gregorian Mesh3dYcalendar = "gregorian" // gregorian
    Mesh3dYcalendar_chinese Mesh3dYcalendar = "chinese" // chinese
    Mesh3dYcalendar_coptic Mesh3dYcalendar = "coptic" // coptic
    Mesh3dYcalendar_discworld Mesh3dYcalendar = "discworld" // discworld
    Mesh3dYcalendar_ethiopian Mesh3dYcalendar = "ethiopian" // ethiopian
    Mesh3dYcalendar_hebrew Mesh3dYcalendar = "hebrew" // hebrew
    Mesh3dYcalendar_islamic Mesh3dYcalendar = "islamic" // islamic
    Mesh3dYcalendar_julian Mesh3dYcalendar = "julian" // julian
    Mesh3dYcalendar_mayan Mesh3dYcalendar = "mayan" // mayan
    Mesh3dYcalendar_nanakshahi Mesh3dYcalendar = "nanakshahi" // nanakshahi
    Mesh3dYcalendar_nepali Mesh3dYcalendar = "nepali" // nepali
    Mesh3dYcalendar_persian Mesh3dYcalendar = "persian" // persian
    Mesh3dYcalendar_jalali Mesh3dYcalendar = "jalali" // jalali
    Mesh3dYcalendar_taiwan Mesh3dYcalendar = "taiwan" // taiwan
    Mesh3dYcalendar_thai Mesh3dYcalendar = "thai" // thai
    Mesh3dYcalendar_ummalqura Mesh3dYcalendar = "ummalqura" // ummalqura
)

// Mesh3dZcalendar Sets the calendar system to use with `z` date data.
type Mesh3dZcalendar string

const (
    Mesh3dZcalendar_gregorian Mesh3dZcalendar = "gregorian" // gregorian
    Mesh3dZcalendar_chinese Mesh3dZcalendar = "chinese" // chinese
    Mesh3dZcalendar_coptic Mesh3dZcalendar = "coptic" // coptic
    Mesh3dZcalendar_discworld Mesh3dZcalendar = "discworld" // discworld
    Mesh3dZcalendar_ethiopian Mesh3dZcalendar = "ethiopian" // ethiopian
    Mesh3dZcalendar_hebrew Mesh3dZcalendar = "hebrew" // hebrew
    Mesh3dZcalendar_islamic Mesh3dZcalendar = "islamic" // islamic
    Mesh3dZcalendar_julian Mesh3dZcalendar = "julian" // julian
    Mesh3dZcalendar_mayan Mesh3dZcalendar = "mayan" // mayan
    Mesh3dZcalendar_nanakshahi Mesh3dZcalendar = "nanakshahi" // nanakshahi
    Mesh3dZcalendar_nepali Mesh3dZcalendar = "nepali" // nepali
    Mesh3dZcalendar_persian Mesh3dZcalendar = "persian" // persian
    Mesh3dZcalendar_jalali Mesh3dZcalendar = "jalali" // jalali
    Mesh3dZcalendar_taiwan Mesh3dZcalendar = "taiwan" // taiwan
    Mesh3dZcalendar_thai Mesh3dZcalendar = "thai" // thai
    Mesh3dZcalendar_ummalqura Mesh3dZcalendar = "ummalqura" // ummalqura
)

// OhlcHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type OhlcHoverlabelAlign string

const (
    OhlcHoverlabelAlign_left OhlcHoverlabelAlign = "left" // left
    OhlcHoverlabelAlign_right OhlcHoverlabelAlign = "right" // right
    OhlcHoverlabelAlign_auto OhlcHoverlabelAlign = "auto" // auto
)

// OhlcVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type OhlcVisible interface{}

const (
    OhlcVisibleTrue bool = true
    OhlcVisibleFalse bool = false
    OhlcVisibleLegendonly string = "legendonly"
)

// OhlcXcalendar Sets the calendar system to use with `x` date data.
type OhlcXcalendar string

const (
    OhlcXcalendar_gregorian OhlcXcalendar = "gregorian" // gregorian
    OhlcXcalendar_chinese OhlcXcalendar = "chinese" // chinese
    OhlcXcalendar_coptic OhlcXcalendar = "coptic" // coptic
    OhlcXcalendar_discworld OhlcXcalendar = "discworld" // discworld
    OhlcXcalendar_ethiopian OhlcXcalendar = "ethiopian" // ethiopian
    OhlcXcalendar_hebrew OhlcXcalendar = "hebrew" // hebrew
    OhlcXcalendar_islamic OhlcXcalendar = "islamic" // islamic
    OhlcXcalendar_julian OhlcXcalendar = "julian" // julian
    OhlcXcalendar_mayan OhlcXcalendar = "mayan" // mayan
    OhlcXcalendar_nanakshahi OhlcXcalendar = "nanakshahi" // nanakshahi
    OhlcXcalendar_nepali OhlcXcalendar = "nepali" // nepali
    OhlcXcalendar_persian OhlcXcalendar = "persian" // persian
    OhlcXcalendar_jalali OhlcXcalendar = "jalali" // jalali
    OhlcXcalendar_taiwan OhlcXcalendar = "taiwan" // taiwan
    OhlcXcalendar_thai OhlcXcalendar = "thai" // thai
    OhlcXcalendar_ummalqura OhlcXcalendar = "ummalqura" // ummalqura
)

// ParcatsArrangement Sets the drag interaction mode for categories and dimensions. If `perpendicular`, the categories can only move along a line perpendicular to the paths. If `freeform`, the categories can freely move on the plane. If `fixed`, the categories and dimensions are stationary.
type ParcatsArrangement string

const (
    ParcatsArrangement_perpendicular ParcatsArrangement = "perpendicular" // perpendicular
    ParcatsArrangement_freeform ParcatsArrangement = "freeform" // freeform
    ParcatsArrangement_fixed ParcatsArrangement = "fixed" // fixed
)

// ParcatsHoveron Sets the hover interaction mode for the parcats diagram. If `category`, hover interaction take place per category. If `color`, hover interactions take place per color per category. If `dimension`, hover interactions take place across all categories per dimension.
type ParcatsHoveron string

const (
    ParcatsHoveron_category ParcatsHoveron = "category" // category
    ParcatsHoveron_color ParcatsHoveron = "color" // color
    ParcatsHoveron_dimension ParcatsHoveron = "dimension" // dimension
)

// ParcatsLineColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ParcatsLineColorbarExponentformat string

const (
    ParcatsLineColorbarExponentformat_none ParcatsLineColorbarExponentformat = "none" // none
    ParcatsLineColorbarExponentformat_e ParcatsLineColorbarExponentformat = "e" // e
    ParcatsLineColorbarExponentformat_E ParcatsLineColorbarExponentformat = "E" // E
    ParcatsLineColorbarExponentformat_power ParcatsLineColorbarExponentformat = "power" // power
    ParcatsLineColorbarExponentformat_SI ParcatsLineColorbarExponentformat = "SI" // SI
    ParcatsLineColorbarExponentformat_B ParcatsLineColorbarExponentformat = "B" // B
)

// ParcatsLineColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ParcatsLineColorbarLenmode string

const (
    ParcatsLineColorbarLenmode_fraction ParcatsLineColorbarLenmode = "fraction" // fraction
    ParcatsLineColorbarLenmode_pixels ParcatsLineColorbarLenmode = "pixels" // pixels
)

// ParcatsLineColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ParcatsLineColorbarShowexponent string

const (
    ParcatsLineColorbarShowexponent_all ParcatsLineColorbarShowexponent = "all" // all
    ParcatsLineColorbarShowexponent_first ParcatsLineColorbarShowexponent = "first" // first
    ParcatsLineColorbarShowexponent_last ParcatsLineColorbarShowexponent = "last" // last
    ParcatsLineColorbarShowexponent_none ParcatsLineColorbarShowexponent = "none" // none
)

// ParcatsLineColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ParcatsLineColorbarShowtickprefix string

const (
    ParcatsLineColorbarShowtickprefix_all ParcatsLineColorbarShowtickprefix = "all" // all
    ParcatsLineColorbarShowtickprefix_first ParcatsLineColorbarShowtickprefix = "first" // first
    ParcatsLineColorbarShowtickprefix_last ParcatsLineColorbarShowtickprefix = "last" // last
    ParcatsLineColorbarShowtickprefix_none ParcatsLineColorbarShowtickprefix = "none" // none
)

// ParcatsLineColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ParcatsLineColorbarShowticksuffix string

const (
    ParcatsLineColorbarShowticksuffix_all ParcatsLineColorbarShowticksuffix = "all" // all
    ParcatsLineColorbarShowticksuffix_first ParcatsLineColorbarShowticksuffix = "first" // first
    ParcatsLineColorbarShowticksuffix_last ParcatsLineColorbarShowticksuffix = "last" // last
    ParcatsLineColorbarShowticksuffix_none ParcatsLineColorbarShowticksuffix = "none" // none
)

// ParcatsLineColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ParcatsLineColorbarThicknessmode string

const (
    ParcatsLineColorbarThicknessmode_fraction ParcatsLineColorbarThicknessmode = "fraction" // fraction
    ParcatsLineColorbarThicknessmode_pixels ParcatsLineColorbarThicknessmode = "pixels" // pixels
)

// ParcatsLineColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ParcatsLineColorbarTickmode string

const (
    ParcatsLineColorbarTickmode_auto ParcatsLineColorbarTickmode = "auto" // auto
    ParcatsLineColorbarTickmode_linear ParcatsLineColorbarTickmode = "linear" // linear
    ParcatsLineColorbarTickmode_array ParcatsLineColorbarTickmode = "array" // array
)

// ParcatsLineColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ParcatsLineColorbarTicks string

const (
    ParcatsLineColorbarTicks_outside ParcatsLineColorbarTicks = "outside" // outside
    ParcatsLineColorbarTicks_inside ParcatsLineColorbarTicks = "inside" // inside
)

// ParcatsLineColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ParcatsLineColorbarTitleSide string

const (
    ParcatsLineColorbarTitleSide_right ParcatsLineColorbarTitleSide = "right" // right
    ParcatsLineColorbarTitleSide_top ParcatsLineColorbarTitleSide = "top" // top
    ParcatsLineColorbarTitleSide_bottom ParcatsLineColorbarTitleSide = "bottom" // bottom
)

// ParcatsLineColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ParcatsLineColorbarXanchor string

const (
    ParcatsLineColorbarXanchor_left ParcatsLineColorbarXanchor = "left" // left
    ParcatsLineColorbarXanchor_center ParcatsLineColorbarXanchor = "center" // center
    ParcatsLineColorbarXanchor_right ParcatsLineColorbarXanchor = "right" // right
)

// ParcatsLineColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ParcatsLineColorbarYanchor string

const (
    ParcatsLineColorbarYanchor_top ParcatsLineColorbarYanchor = "top" // top
    ParcatsLineColorbarYanchor_middle ParcatsLineColorbarYanchor = "middle" // middle
    ParcatsLineColorbarYanchor_bottom ParcatsLineColorbarYanchor = "bottom" // bottom
)

// ParcatsLineShape Sets the shape of the paths. If `linear`, paths are composed of straight lines. If `hspline`, paths are composed of horizontal curved splines
type ParcatsLineShape string

const (
    ParcatsLineShape_linear ParcatsLineShape = "linear" // linear
    ParcatsLineShape_hspline ParcatsLineShape = "hspline" // hspline
)

// ParcatsSortpaths Sets the path sorting algorithm. If `forward`, sort paths based on dimension categories from left to right. If `backward`, sort paths based on dimensions categories from right to left.
type ParcatsSortpaths string

const (
    ParcatsSortpaths_forward ParcatsSortpaths = "forward" // forward
    ParcatsSortpaths_backward ParcatsSortpaths = "backward" // backward
)

// ParcatsVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ParcatsVisible interface{}

const (
    ParcatsVisibleTrue bool = true
    ParcatsVisibleFalse bool = false
    ParcatsVisibleLegendonly string = "legendonly"
)

// ParcoordsLabelside Specifies the location of the `label`. *top* positions labels above, next to the title *bottom* positions labels below the graph Tilted labels with *labelangle* may be positioned better inside margins when `labelposition` is set to *bottom*.
type ParcoordsLabelside string

const (
    ParcoordsLabelside_top ParcoordsLabelside = "top" // top
    ParcoordsLabelside_bottom ParcoordsLabelside = "bottom" // bottom
)

// ParcoordsLineColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ParcoordsLineColorbarExponentformat string

const (
    ParcoordsLineColorbarExponentformat_none ParcoordsLineColorbarExponentformat = "none" // none
    ParcoordsLineColorbarExponentformat_e ParcoordsLineColorbarExponentformat = "e" // e
    ParcoordsLineColorbarExponentformat_E ParcoordsLineColorbarExponentformat = "E" // E
    ParcoordsLineColorbarExponentformat_power ParcoordsLineColorbarExponentformat = "power" // power
    ParcoordsLineColorbarExponentformat_SI ParcoordsLineColorbarExponentformat = "SI" // SI
    ParcoordsLineColorbarExponentformat_B ParcoordsLineColorbarExponentformat = "B" // B
)

// ParcoordsLineColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ParcoordsLineColorbarLenmode string

const (
    ParcoordsLineColorbarLenmode_fraction ParcoordsLineColorbarLenmode = "fraction" // fraction
    ParcoordsLineColorbarLenmode_pixels ParcoordsLineColorbarLenmode = "pixels" // pixels
)

// ParcoordsLineColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ParcoordsLineColorbarShowexponent string

const (
    ParcoordsLineColorbarShowexponent_all ParcoordsLineColorbarShowexponent = "all" // all
    ParcoordsLineColorbarShowexponent_first ParcoordsLineColorbarShowexponent = "first" // first
    ParcoordsLineColorbarShowexponent_last ParcoordsLineColorbarShowexponent = "last" // last
    ParcoordsLineColorbarShowexponent_none ParcoordsLineColorbarShowexponent = "none" // none
)

// ParcoordsLineColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ParcoordsLineColorbarShowtickprefix string

const (
    ParcoordsLineColorbarShowtickprefix_all ParcoordsLineColorbarShowtickprefix = "all" // all
    ParcoordsLineColorbarShowtickprefix_first ParcoordsLineColorbarShowtickprefix = "first" // first
    ParcoordsLineColorbarShowtickprefix_last ParcoordsLineColorbarShowtickprefix = "last" // last
    ParcoordsLineColorbarShowtickprefix_none ParcoordsLineColorbarShowtickprefix = "none" // none
)

// ParcoordsLineColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ParcoordsLineColorbarShowticksuffix string

const (
    ParcoordsLineColorbarShowticksuffix_all ParcoordsLineColorbarShowticksuffix = "all" // all
    ParcoordsLineColorbarShowticksuffix_first ParcoordsLineColorbarShowticksuffix = "first" // first
    ParcoordsLineColorbarShowticksuffix_last ParcoordsLineColorbarShowticksuffix = "last" // last
    ParcoordsLineColorbarShowticksuffix_none ParcoordsLineColorbarShowticksuffix = "none" // none
)

// ParcoordsLineColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ParcoordsLineColorbarThicknessmode string

const (
    ParcoordsLineColorbarThicknessmode_fraction ParcoordsLineColorbarThicknessmode = "fraction" // fraction
    ParcoordsLineColorbarThicknessmode_pixels ParcoordsLineColorbarThicknessmode = "pixels" // pixels
)

// ParcoordsLineColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ParcoordsLineColorbarTickmode string

const (
    ParcoordsLineColorbarTickmode_auto ParcoordsLineColorbarTickmode = "auto" // auto
    ParcoordsLineColorbarTickmode_linear ParcoordsLineColorbarTickmode = "linear" // linear
    ParcoordsLineColorbarTickmode_array ParcoordsLineColorbarTickmode = "array" // array
)

// ParcoordsLineColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ParcoordsLineColorbarTicks string

const (
    ParcoordsLineColorbarTicks_outside ParcoordsLineColorbarTicks = "outside" // outside
    ParcoordsLineColorbarTicks_inside ParcoordsLineColorbarTicks = "inside" // inside
)

// ParcoordsLineColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ParcoordsLineColorbarTitleSide string

const (
    ParcoordsLineColorbarTitleSide_right ParcoordsLineColorbarTitleSide = "right" // right
    ParcoordsLineColorbarTitleSide_top ParcoordsLineColorbarTitleSide = "top" // top
    ParcoordsLineColorbarTitleSide_bottom ParcoordsLineColorbarTitleSide = "bottom" // bottom
)

// ParcoordsLineColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ParcoordsLineColorbarXanchor string

const (
    ParcoordsLineColorbarXanchor_left ParcoordsLineColorbarXanchor = "left" // left
    ParcoordsLineColorbarXanchor_center ParcoordsLineColorbarXanchor = "center" // center
    ParcoordsLineColorbarXanchor_right ParcoordsLineColorbarXanchor = "right" // right
)

// ParcoordsLineColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ParcoordsLineColorbarYanchor string

const (
    ParcoordsLineColorbarYanchor_top ParcoordsLineColorbarYanchor = "top" // top
    ParcoordsLineColorbarYanchor_middle ParcoordsLineColorbarYanchor = "middle" // middle
    ParcoordsLineColorbarYanchor_bottom ParcoordsLineColorbarYanchor = "bottom" // bottom
)

// ParcoordsVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ParcoordsVisible interface{}

const (
    ParcoordsVisibleTrue bool = true
    ParcoordsVisibleFalse bool = false
    ParcoordsVisibleLegendonly string = "legendonly"
)

// PieDirection Specifies the direction at which succeeding sectors follow one another.
type PieDirection string

const (
    PieDirection_clockwise PieDirection = "clockwise" // clockwise
    PieDirection_counterclockwise PieDirection = "counterclockwise" // counterclockwise
)

// PieHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type PieHoverlabelAlign string

const (
    PieHoverlabelAlign_left PieHoverlabelAlign = "left" // left
    PieHoverlabelAlign_right PieHoverlabelAlign = "right" // right
    PieHoverlabelAlign_auto PieHoverlabelAlign = "auto" // auto
)

// PieInsidetextorientation Controls the orientation of the text inside chart sectors. When set to *auto*, text may be oriented in any direction in order to be as big as possible in the middle of a sector. The *horizontal* option orients text to be parallel with the bottom of the chart, and may make text smaller in order to achieve that goal. The *radial* option orients text along the radius of the sector. The *tangential* option orients text perpendicular to the radius of the sector.
type PieInsidetextorientation string

const (
    PieInsidetextorientation_horizontal PieInsidetextorientation = "horizontal" // horizontal
    PieInsidetextorientation_radial PieInsidetextorientation = "radial" // radial
    PieInsidetextorientation_tangential PieInsidetextorientation = "tangential" // tangential
    PieInsidetextorientation_auto PieInsidetextorientation = "auto" // auto
)

// PieTextposition Specifies the location of the `textinfo`.
type PieTextposition string

const (
    PieTextposition_inside PieTextposition = "inside" // inside
    PieTextposition_outside PieTextposition = "outside" // outside
    PieTextposition_auto PieTextposition = "auto" // auto
    PieTextposition_none PieTextposition = "none" // none
)

// PieTitlePosition Specifies the location of the `title`. Note that the title's position used to be set by the now deprecated `titleposition` attribute.
type PieTitlePosition string

const (
    PieTitlePosition_topleft PieTitlePosition = "top left" // top left
    PieTitlePosition_topcenter PieTitlePosition = "top center" // top center
    PieTitlePosition_topright PieTitlePosition = "top right" // top right
    PieTitlePosition_middlecenter PieTitlePosition = "middle center" // middle center
    PieTitlePosition_bottomleft PieTitlePosition = "bottom left" // bottom left
    PieTitlePosition_bottomcenter PieTitlePosition = "bottom center" // bottom center
    PieTitlePosition_bottomright PieTitlePosition = "bottom right" // bottom right
)

// PieVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type PieVisible interface{}

const (
    PieVisibleTrue bool = true
    PieVisibleFalse bool = false
    PieVisibleLegendonly string = "legendonly"
)

// PointcloudHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type PointcloudHoverlabelAlign string

const (
    PointcloudHoverlabelAlign_left PointcloudHoverlabelAlign = "left" // left
    PointcloudHoverlabelAlign_right PointcloudHoverlabelAlign = "right" // right
    PointcloudHoverlabelAlign_auto PointcloudHoverlabelAlign = "auto" // auto
)

// PointcloudVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type PointcloudVisible interface{}

const (
    PointcloudVisibleTrue bool = true
    PointcloudVisibleFalse bool = false
    PointcloudVisibleLegendonly string = "legendonly"
)

// SankeyArrangement If value is `snap` (the default), the node arrangement is assisted by automatic snapping of elements to preserve space between nodes specified via `nodepad`. If value is `perpendicular`, the nodes can only move along a line perpendicular to the flow. If value is `freeform`, the nodes can freely move on the plane. If value is `fixed`, the nodes are stationary.
type SankeyArrangement string

const (
    SankeyArrangement_snap SankeyArrangement = "snap" // snap
    SankeyArrangement_perpendicular SankeyArrangement = "perpendicular" // perpendicular
    SankeyArrangement_freeform SankeyArrangement = "freeform" // freeform
    SankeyArrangement_fixed SankeyArrangement = "fixed" // fixed
)

// SankeyHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SankeyHoverlabelAlign string

const (
    SankeyHoverlabelAlign_left SankeyHoverlabelAlign = "left" // left
    SankeyHoverlabelAlign_right SankeyHoverlabelAlign = "right" // right
    SankeyHoverlabelAlign_auto SankeyHoverlabelAlign = "auto" // auto
)

// SankeyLinkHoverinfo Determines which trace information appear when hovering links. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type SankeyLinkHoverinfo string

const (
    SankeyLinkHoverinfo_all SankeyLinkHoverinfo = "all" // all
    SankeyLinkHoverinfo_none SankeyLinkHoverinfo = "none" // none
    SankeyLinkHoverinfo_skip SankeyLinkHoverinfo = "skip" // skip
)

// SankeyLinkHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SankeyLinkHoverlabelAlign string

const (
    SankeyLinkHoverlabelAlign_left SankeyLinkHoverlabelAlign = "left" // left
    SankeyLinkHoverlabelAlign_right SankeyLinkHoverlabelAlign = "right" // right
    SankeyLinkHoverlabelAlign_auto SankeyLinkHoverlabelAlign = "auto" // auto
)

// SankeyNodeHoverinfo Determines which trace information appear when hovering nodes. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
type SankeyNodeHoverinfo string

const (
    SankeyNodeHoverinfo_all SankeyNodeHoverinfo = "all" // all
    SankeyNodeHoverinfo_none SankeyNodeHoverinfo = "none" // none
    SankeyNodeHoverinfo_skip SankeyNodeHoverinfo = "skip" // skip
)

// SankeyNodeHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SankeyNodeHoverlabelAlign string

const (
    SankeyNodeHoverlabelAlign_left SankeyNodeHoverlabelAlign = "left" // left
    SankeyNodeHoverlabelAlign_right SankeyNodeHoverlabelAlign = "right" // right
    SankeyNodeHoverlabelAlign_auto SankeyNodeHoverlabelAlign = "auto" // auto
)

// SankeyOrientation Sets the orientation of the Sankey diagram.
type SankeyOrientation string

const (
    SankeyOrientation_v SankeyOrientation = "v" // v
    SankeyOrientation_h SankeyOrientation = "h" // h
)

// SankeyVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type SankeyVisible interface{}

const (
    SankeyVisibleTrue bool = true
    SankeyVisibleFalse bool = false
    SankeyVisibleLegendonly string = "legendonly"
)

// Scatter3dErrorXType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type Scatter3dErrorXType string

const (
    Scatter3dErrorXType_percent Scatter3dErrorXType = "percent" // percent
    Scatter3dErrorXType_constant Scatter3dErrorXType = "constant" // constant
    Scatter3dErrorXType_sqrt Scatter3dErrorXType = "sqrt" // sqrt
    Scatter3dErrorXType_data Scatter3dErrorXType = "data" // data
)

// Scatter3dErrorYType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type Scatter3dErrorYType string

const (
    Scatter3dErrorYType_percent Scatter3dErrorYType = "percent" // percent
    Scatter3dErrorYType_constant Scatter3dErrorYType = "constant" // constant
    Scatter3dErrorYType_sqrt Scatter3dErrorYType = "sqrt" // sqrt
    Scatter3dErrorYType_data Scatter3dErrorYType = "data" // data
)

// Scatter3dErrorZType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type Scatter3dErrorZType string

const (
    Scatter3dErrorZType_percent Scatter3dErrorZType = "percent" // percent
    Scatter3dErrorZType_constant Scatter3dErrorZType = "constant" // constant
    Scatter3dErrorZType_sqrt Scatter3dErrorZType = "sqrt" // sqrt
    Scatter3dErrorZType_data Scatter3dErrorZType = "data" // data
)

// Scatter3dHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type Scatter3dHoverlabelAlign string

const (
    Scatter3dHoverlabelAlign_left Scatter3dHoverlabelAlign = "left" // left
    Scatter3dHoverlabelAlign_right Scatter3dHoverlabelAlign = "right" // right
    Scatter3dHoverlabelAlign_auto Scatter3dHoverlabelAlign = "auto" // auto
)

// Scatter3dLineColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type Scatter3dLineColorbarExponentformat string

const (
    Scatter3dLineColorbarExponentformat_none Scatter3dLineColorbarExponentformat = "none" // none
    Scatter3dLineColorbarExponentformat_e Scatter3dLineColorbarExponentformat = "e" // e
    Scatter3dLineColorbarExponentformat_E Scatter3dLineColorbarExponentformat = "E" // E
    Scatter3dLineColorbarExponentformat_power Scatter3dLineColorbarExponentformat = "power" // power
    Scatter3dLineColorbarExponentformat_SI Scatter3dLineColorbarExponentformat = "SI" // SI
    Scatter3dLineColorbarExponentformat_B Scatter3dLineColorbarExponentformat = "B" // B
)

// Scatter3dLineColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type Scatter3dLineColorbarLenmode string

const (
    Scatter3dLineColorbarLenmode_fraction Scatter3dLineColorbarLenmode = "fraction" // fraction
    Scatter3dLineColorbarLenmode_pixels Scatter3dLineColorbarLenmode = "pixels" // pixels
)

// Scatter3dLineColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type Scatter3dLineColorbarShowexponent string

const (
    Scatter3dLineColorbarShowexponent_all Scatter3dLineColorbarShowexponent = "all" // all
    Scatter3dLineColorbarShowexponent_first Scatter3dLineColorbarShowexponent = "first" // first
    Scatter3dLineColorbarShowexponent_last Scatter3dLineColorbarShowexponent = "last" // last
    Scatter3dLineColorbarShowexponent_none Scatter3dLineColorbarShowexponent = "none" // none
)

// Scatter3dLineColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type Scatter3dLineColorbarShowtickprefix string

const (
    Scatter3dLineColorbarShowtickprefix_all Scatter3dLineColorbarShowtickprefix = "all" // all
    Scatter3dLineColorbarShowtickprefix_first Scatter3dLineColorbarShowtickprefix = "first" // first
    Scatter3dLineColorbarShowtickprefix_last Scatter3dLineColorbarShowtickprefix = "last" // last
    Scatter3dLineColorbarShowtickprefix_none Scatter3dLineColorbarShowtickprefix = "none" // none
)

// Scatter3dLineColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type Scatter3dLineColorbarShowticksuffix string

const (
    Scatter3dLineColorbarShowticksuffix_all Scatter3dLineColorbarShowticksuffix = "all" // all
    Scatter3dLineColorbarShowticksuffix_first Scatter3dLineColorbarShowticksuffix = "first" // first
    Scatter3dLineColorbarShowticksuffix_last Scatter3dLineColorbarShowticksuffix = "last" // last
    Scatter3dLineColorbarShowticksuffix_none Scatter3dLineColorbarShowticksuffix = "none" // none
)

// Scatter3dLineColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type Scatter3dLineColorbarThicknessmode string

const (
    Scatter3dLineColorbarThicknessmode_fraction Scatter3dLineColorbarThicknessmode = "fraction" // fraction
    Scatter3dLineColorbarThicknessmode_pixels Scatter3dLineColorbarThicknessmode = "pixels" // pixels
)

// Scatter3dLineColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type Scatter3dLineColorbarTickmode string

const (
    Scatter3dLineColorbarTickmode_auto Scatter3dLineColorbarTickmode = "auto" // auto
    Scatter3dLineColorbarTickmode_linear Scatter3dLineColorbarTickmode = "linear" // linear
    Scatter3dLineColorbarTickmode_array Scatter3dLineColorbarTickmode = "array" // array
)

// Scatter3dLineColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type Scatter3dLineColorbarTicks string

const (
    Scatter3dLineColorbarTicks_outside Scatter3dLineColorbarTicks = "outside" // outside
    Scatter3dLineColorbarTicks_inside Scatter3dLineColorbarTicks = "inside" // inside
)

// Scatter3dLineColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type Scatter3dLineColorbarTitleSide string

const (
    Scatter3dLineColorbarTitleSide_right Scatter3dLineColorbarTitleSide = "right" // right
    Scatter3dLineColorbarTitleSide_top Scatter3dLineColorbarTitleSide = "top" // top
    Scatter3dLineColorbarTitleSide_bottom Scatter3dLineColorbarTitleSide = "bottom" // bottom
)

// Scatter3dLineColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type Scatter3dLineColorbarXanchor string

const (
    Scatter3dLineColorbarXanchor_left Scatter3dLineColorbarXanchor = "left" // left
    Scatter3dLineColorbarXanchor_center Scatter3dLineColorbarXanchor = "center" // center
    Scatter3dLineColorbarXanchor_right Scatter3dLineColorbarXanchor = "right" // right
)

// Scatter3dLineColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type Scatter3dLineColorbarYanchor string

const (
    Scatter3dLineColorbarYanchor_top Scatter3dLineColorbarYanchor = "top" // top
    Scatter3dLineColorbarYanchor_middle Scatter3dLineColorbarYanchor = "middle" // middle
    Scatter3dLineColorbarYanchor_bottom Scatter3dLineColorbarYanchor = "bottom" // bottom
)

// Scatter3dLineDash Sets the dash style of the lines.
type Scatter3dLineDash string

const (
    Scatter3dLineDash_solid Scatter3dLineDash = "solid" // solid
    Scatter3dLineDash_dot Scatter3dLineDash = "dot" // dot
    Scatter3dLineDash_dash Scatter3dLineDash = "dash" // dash
    Scatter3dLineDash_longdash Scatter3dLineDash = "longdash" // longdash
    Scatter3dLineDash_dashdot Scatter3dLineDash = "dashdot" // dashdot
    Scatter3dLineDash_longdashdot Scatter3dLineDash = "longdashdot" // longdashdot
)

// Scatter3dMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type Scatter3dMarkerColorbarExponentformat string

const (
    Scatter3dMarkerColorbarExponentformat_none Scatter3dMarkerColorbarExponentformat = "none" // none
    Scatter3dMarkerColorbarExponentformat_e Scatter3dMarkerColorbarExponentformat = "e" // e
    Scatter3dMarkerColorbarExponentformat_E Scatter3dMarkerColorbarExponentformat = "E" // E
    Scatter3dMarkerColorbarExponentformat_power Scatter3dMarkerColorbarExponentformat = "power" // power
    Scatter3dMarkerColorbarExponentformat_SI Scatter3dMarkerColorbarExponentformat = "SI" // SI
    Scatter3dMarkerColorbarExponentformat_B Scatter3dMarkerColorbarExponentformat = "B" // B
)

// Scatter3dMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type Scatter3dMarkerColorbarLenmode string

const (
    Scatter3dMarkerColorbarLenmode_fraction Scatter3dMarkerColorbarLenmode = "fraction" // fraction
    Scatter3dMarkerColorbarLenmode_pixels Scatter3dMarkerColorbarLenmode = "pixels" // pixels
)

// Scatter3dMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type Scatter3dMarkerColorbarShowexponent string

const (
    Scatter3dMarkerColorbarShowexponent_all Scatter3dMarkerColorbarShowexponent = "all" // all
    Scatter3dMarkerColorbarShowexponent_first Scatter3dMarkerColorbarShowexponent = "first" // first
    Scatter3dMarkerColorbarShowexponent_last Scatter3dMarkerColorbarShowexponent = "last" // last
    Scatter3dMarkerColorbarShowexponent_none Scatter3dMarkerColorbarShowexponent = "none" // none
)

// Scatter3dMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type Scatter3dMarkerColorbarShowtickprefix string

const (
    Scatter3dMarkerColorbarShowtickprefix_all Scatter3dMarkerColorbarShowtickprefix = "all" // all
    Scatter3dMarkerColorbarShowtickprefix_first Scatter3dMarkerColorbarShowtickprefix = "first" // first
    Scatter3dMarkerColorbarShowtickprefix_last Scatter3dMarkerColorbarShowtickprefix = "last" // last
    Scatter3dMarkerColorbarShowtickprefix_none Scatter3dMarkerColorbarShowtickprefix = "none" // none
)

// Scatter3dMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type Scatter3dMarkerColorbarShowticksuffix string

const (
    Scatter3dMarkerColorbarShowticksuffix_all Scatter3dMarkerColorbarShowticksuffix = "all" // all
    Scatter3dMarkerColorbarShowticksuffix_first Scatter3dMarkerColorbarShowticksuffix = "first" // first
    Scatter3dMarkerColorbarShowticksuffix_last Scatter3dMarkerColorbarShowticksuffix = "last" // last
    Scatter3dMarkerColorbarShowticksuffix_none Scatter3dMarkerColorbarShowticksuffix = "none" // none
)

// Scatter3dMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type Scatter3dMarkerColorbarThicknessmode string

const (
    Scatter3dMarkerColorbarThicknessmode_fraction Scatter3dMarkerColorbarThicknessmode = "fraction" // fraction
    Scatter3dMarkerColorbarThicknessmode_pixels Scatter3dMarkerColorbarThicknessmode = "pixels" // pixels
)

// Scatter3dMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type Scatter3dMarkerColorbarTickmode string

const (
    Scatter3dMarkerColorbarTickmode_auto Scatter3dMarkerColorbarTickmode = "auto" // auto
    Scatter3dMarkerColorbarTickmode_linear Scatter3dMarkerColorbarTickmode = "linear" // linear
    Scatter3dMarkerColorbarTickmode_array Scatter3dMarkerColorbarTickmode = "array" // array
)

// Scatter3dMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type Scatter3dMarkerColorbarTicks string

const (
    Scatter3dMarkerColorbarTicks_outside Scatter3dMarkerColorbarTicks = "outside" // outside
    Scatter3dMarkerColorbarTicks_inside Scatter3dMarkerColorbarTicks = "inside" // inside
)

// Scatter3dMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type Scatter3dMarkerColorbarTitleSide string

const (
    Scatter3dMarkerColorbarTitleSide_right Scatter3dMarkerColorbarTitleSide = "right" // right
    Scatter3dMarkerColorbarTitleSide_top Scatter3dMarkerColorbarTitleSide = "top" // top
    Scatter3dMarkerColorbarTitleSide_bottom Scatter3dMarkerColorbarTitleSide = "bottom" // bottom
)

// Scatter3dMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type Scatter3dMarkerColorbarXanchor string

const (
    Scatter3dMarkerColorbarXanchor_left Scatter3dMarkerColorbarXanchor = "left" // left
    Scatter3dMarkerColorbarXanchor_center Scatter3dMarkerColorbarXanchor = "center" // center
    Scatter3dMarkerColorbarXanchor_right Scatter3dMarkerColorbarXanchor = "right" // right
)

// Scatter3dMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type Scatter3dMarkerColorbarYanchor string

const (
    Scatter3dMarkerColorbarYanchor_top Scatter3dMarkerColorbarYanchor = "top" // top
    Scatter3dMarkerColorbarYanchor_middle Scatter3dMarkerColorbarYanchor = "middle" // middle
    Scatter3dMarkerColorbarYanchor_bottom Scatter3dMarkerColorbarYanchor = "bottom" // bottom
)

// Scatter3dMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type Scatter3dMarkerSizemode string

const (
    Scatter3dMarkerSizemode_diameter Scatter3dMarkerSizemode = "diameter" // diameter
    Scatter3dMarkerSizemode_area Scatter3dMarkerSizemode = "area" // area
)

// Scatter3dMarkerSymbol Sets the marker symbol type.
type Scatter3dMarkerSymbol string

const (
    Scatter3dMarkerSymbol_circle Scatter3dMarkerSymbol = "circle" // circle
    Scatter3dMarkerSymbol_circle_open Scatter3dMarkerSymbol = "circle-open" // circle-open
    Scatter3dMarkerSymbol_square Scatter3dMarkerSymbol = "square" // square
    Scatter3dMarkerSymbol_square_open Scatter3dMarkerSymbol = "square-open" // square-open
    Scatter3dMarkerSymbol_diamond Scatter3dMarkerSymbol = "diamond" // diamond
    Scatter3dMarkerSymbol_diamond_open Scatter3dMarkerSymbol = "diamond-open" // diamond-open
    Scatter3dMarkerSymbol_cross Scatter3dMarkerSymbol = "cross" // cross
    Scatter3dMarkerSymbol_x Scatter3dMarkerSymbol = "x" // x
)

// Scatter3dSurfaceaxis If *-1*, the scatter points are not fill with a surface If *0*, *1*, *2*, the scatter points are filled with a Delaunay surface about the x, y, z respectively.
type Scatter3dSurfaceaxis string

const (
    Scatter3dSurfaceaxis_1 Scatter3dSurfaceaxis = "-1" // -1
    Scatter3dSurfaceaxis0 Scatter3dSurfaceaxis = "0" // 0
    Scatter3dSurfaceaxis1 Scatter3dSurfaceaxis = "1" // 1
    Scatter3dSurfaceaxis2 Scatter3dSurfaceaxis = "2" // 2
)

// Scatter3dTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type Scatter3dTextposition string

const (
    Scatter3dTextposition_topleft Scatter3dTextposition = "top left" // top left
    Scatter3dTextposition_topcenter Scatter3dTextposition = "top center" // top center
    Scatter3dTextposition_topright Scatter3dTextposition = "top right" // top right
    Scatter3dTextposition_middleleft Scatter3dTextposition = "middle left" // middle left
    Scatter3dTextposition_middlecenter Scatter3dTextposition = "middle center" // middle center
    Scatter3dTextposition_middleright Scatter3dTextposition = "middle right" // middle right
    Scatter3dTextposition_bottomleft Scatter3dTextposition = "bottom left" // bottom left
    Scatter3dTextposition_bottomcenter Scatter3dTextposition = "bottom center" // bottom center
    Scatter3dTextposition_bottomright Scatter3dTextposition = "bottom right" // bottom right
)

// Scatter3dVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type Scatter3dVisible interface{}

const (
    Scatter3dVisibleTrue bool = true
    Scatter3dVisibleFalse bool = false
    Scatter3dVisibleLegendonly string = "legendonly"
)

// Scatter3dXcalendar Sets the calendar system to use with `x` date data.
type Scatter3dXcalendar string

const (
    Scatter3dXcalendar_gregorian Scatter3dXcalendar = "gregorian" // gregorian
    Scatter3dXcalendar_chinese Scatter3dXcalendar = "chinese" // chinese
    Scatter3dXcalendar_coptic Scatter3dXcalendar = "coptic" // coptic
    Scatter3dXcalendar_discworld Scatter3dXcalendar = "discworld" // discworld
    Scatter3dXcalendar_ethiopian Scatter3dXcalendar = "ethiopian" // ethiopian
    Scatter3dXcalendar_hebrew Scatter3dXcalendar = "hebrew" // hebrew
    Scatter3dXcalendar_islamic Scatter3dXcalendar = "islamic" // islamic
    Scatter3dXcalendar_julian Scatter3dXcalendar = "julian" // julian
    Scatter3dXcalendar_mayan Scatter3dXcalendar = "mayan" // mayan
    Scatter3dXcalendar_nanakshahi Scatter3dXcalendar = "nanakshahi" // nanakshahi
    Scatter3dXcalendar_nepali Scatter3dXcalendar = "nepali" // nepali
    Scatter3dXcalendar_persian Scatter3dXcalendar = "persian" // persian
    Scatter3dXcalendar_jalali Scatter3dXcalendar = "jalali" // jalali
    Scatter3dXcalendar_taiwan Scatter3dXcalendar = "taiwan" // taiwan
    Scatter3dXcalendar_thai Scatter3dXcalendar = "thai" // thai
    Scatter3dXcalendar_ummalqura Scatter3dXcalendar = "ummalqura" // ummalqura
)

// Scatter3dYcalendar Sets the calendar system to use with `y` date data.
type Scatter3dYcalendar string

const (
    Scatter3dYcalendar_gregorian Scatter3dYcalendar = "gregorian" // gregorian
    Scatter3dYcalendar_chinese Scatter3dYcalendar = "chinese" // chinese
    Scatter3dYcalendar_coptic Scatter3dYcalendar = "coptic" // coptic
    Scatter3dYcalendar_discworld Scatter3dYcalendar = "discworld" // discworld
    Scatter3dYcalendar_ethiopian Scatter3dYcalendar = "ethiopian" // ethiopian
    Scatter3dYcalendar_hebrew Scatter3dYcalendar = "hebrew" // hebrew
    Scatter3dYcalendar_islamic Scatter3dYcalendar = "islamic" // islamic
    Scatter3dYcalendar_julian Scatter3dYcalendar = "julian" // julian
    Scatter3dYcalendar_mayan Scatter3dYcalendar = "mayan" // mayan
    Scatter3dYcalendar_nanakshahi Scatter3dYcalendar = "nanakshahi" // nanakshahi
    Scatter3dYcalendar_nepali Scatter3dYcalendar = "nepali" // nepali
    Scatter3dYcalendar_persian Scatter3dYcalendar = "persian" // persian
    Scatter3dYcalendar_jalali Scatter3dYcalendar = "jalali" // jalali
    Scatter3dYcalendar_taiwan Scatter3dYcalendar = "taiwan" // taiwan
    Scatter3dYcalendar_thai Scatter3dYcalendar = "thai" // thai
    Scatter3dYcalendar_ummalqura Scatter3dYcalendar = "ummalqura" // ummalqura
)

// Scatter3dZcalendar Sets the calendar system to use with `z` date data.
type Scatter3dZcalendar string

const (
    Scatter3dZcalendar_gregorian Scatter3dZcalendar = "gregorian" // gregorian
    Scatter3dZcalendar_chinese Scatter3dZcalendar = "chinese" // chinese
    Scatter3dZcalendar_coptic Scatter3dZcalendar = "coptic" // coptic
    Scatter3dZcalendar_discworld Scatter3dZcalendar = "discworld" // discworld
    Scatter3dZcalendar_ethiopian Scatter3dZcalendar = "ethiopian" // ethiopian
    Scatter3dZcalendar_hebrew Scatter3dZcalendar = "hebrew" // hebrew
    Scatter3dZcalendar_islamic Scatter3dZcalendar = "islamic" // islamic
    Scatter3dZcalendar_julian Scatter3dZcalendar = "julian" // julian
    Scatter3dZcalendar_mayan Scatter3dZcalendar = "mayan" // mayan
    Scatter3dZcalendar_nanakshahi Scatter3dZcalendar = "nanakshahi" // nanakshahi
    Scatter3dZcalendar_nepali Scatter3dZcalendar = "nepali" // nepali
    Scatter3dZcalendar_persian Scatter3dZcalendar = "persian" // persian
    Scatter3dZcalendar_jalali Scatter3dZcalendar = "jalali" // jalali
    Scatter3dZcalendar_taiwan Scatter3dZcalendar = "taiwan" // taiwan
    Scatter3dZcalendar_thai Scatter3dZcalendar = "thai" // thai
    Scatter3dZcalendar_ummalqura Scatter3dZcalendar = "ummalqura" // ummalqura
)

// ScatterErrorXType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type ScatterErrorXType string

const (
    ScatterErrorXType_percent ScatterErrorXType = "percent" // percent
    ScatterErrorXType_constant ScatterErrorXType = "constant" // constant
    ScatterErrorXType_sqrt ScatterErrorXType = "sqrt" // sqrt
    ScatterErrorXType_data ScatterErrorXType = "data" // data
)

// ScatterErrorYType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type ScatterErrorYType string

const (
    ScatterErrorYType_percent ScatterErrorYType = "percent" // percent
    ScatterErrorYType_constant ScatterErrorYType = "constant" // constant
    ScatterErrorYType_sqrt ScatterErrorYType = "sqrt" // sqrt
    ScatterErrorYType_data ScatterErrorYType = "data" // data
)

// ScatterFill Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
type ScatterFill string

const (
    ScatterFill_none ScatterFill = "none" // none
    ScatterFill_tozeroy ScatterFill = "tozeroy" // tozeroy
    ScatterFill_tozerox ScatterFill = "tozerox" // tozerox
    ScatterFill_tonexty ScatterFill = "tonexty" // tonexty
    ScatterFill_tonextx ScatterFill = "tonextx" // tonextx
    ScatterFill_toself ScatterFill = "toself" // toself
    ScatterFill_tonext ScatterFill = "tonext" // tonext
)

// ScatterGroupnorm Only relevant when `stackgroup` is used, and only the first `groupnorm` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Sets the normalization for the sum of this `stackgroup`. With *fraction*, the value of each trace at each location is divided by the sum of all trace values at that location. *percent* is the same but multiplied by 100 to show percentages. If there are multiple subplots, or multiple `stackgroup`s on one subplot, each will be normalized within its own set.
type ScatterGroupnorm string

const (
    ScatterGroupnorm_fraction ScatterGroupnorm = "fraction" // fraction
    ScatterGroupnorm_percent ScatterGroupnorm = "percent" // percent
)

// ScatterHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterHoverlabelAlign string

const (
    ScatterHoverlabelAlign_left ScatterHoverlabelAlign = "left" // left
    ScatterHoverlabelAlign_right ScatterHoverlabelAlign = "right" // right
    ScatterHoverlabelAlign_auto ScatterHoverlabelAlign = "auto" // auto
)

// ScatterLineShape Determines the line shape. With *spline* the lines are drawn using spline interpolation. The other available values correspond to step-wise line shapes.
type ScatterLineShape string

const (
    ScatterLineShape_linear ScatterLineShape = "linear" // linear
    ScatterLineShape_spline ScatterLineShape = "spline" // spline
    ScatterLineShape_hv ScatterLineShape = "hv" // hv
    ScatterLineShape_vh ScatterLineShape = "vh" // vh
    ScatterLineShape_hvh ScatterLineShape = "hvh" // hvh
    ScatterLineShape_vhv ScatterLineShape = "vhv" // vhv
)

// ScatterMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterMarkerColorbarExponentformat string

const (
    ScatterMarkerColorbarExponentformat_none ScatterMarkerColorbarExponentformat = "none" // none
    ScatterMarkerColorbarExponentformat_e ScatterMarkerColorbarExponentformat = "e" // e
    ScatterMarkerColorbarExponentformat_E ScatterMarkerColorbarExponentformat = "E" // E
    ScatterMarkerColorbarExponentformat_power ScatterMarkerColorbarExponentformat = "power" // power
    ScatterMarkerColorbarExponentformat_SI ScatterMarkerColorbarExponentformat = "SI" // SI
    ScatterMarkerColorbarExponentformat_B ScatterMarkerColorbarExponentformat = "B" // B
)

// ScatterMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterMarkerColorbarLenmode string

const (
    ScatterMarkerColorbarLenmode_fraction ScatterMarkerColorbarLenmode = "fraction" // fraction
    ScatterMarkerColorbarLenmode_pixels ScatterMarkerColorbarLenmode = "pixels" // pixels
)

// ScatterMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterMarkerColorbarShowexponent string

const (
    ScatterMarkerColorbarShowexponent_all ScatterMarkerColorbarShowexponent = "all" // all
    ScatterMarkerColorbarShowexponent_first ScatterMarkerColorbarShowexponent = "first" // first
    ScatterMarkerColorbarShowexponent_last ScatterMarkerColorbarShowexponent = "last" // last
    ScatterMarkerColorbarShowexponent_none ScatterMarkerColorbarShowexponent = "none" // none
)

// ScatterMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterMarkerColorbarShowtickprefix string

const (
    ScatterMarkerColorbarShowtickprefix_all ScatterMarkerColorbarShowtickprefix = "all" // all
    ScatterMarkerColorbarShowtickprefix_first ScatterMarkerColorbarShowtickprefix = "first" // first
    ScatterMarkerColorbarShowtickprefix_last ScatterMarkerColorbarShowtickprefix = "last" // last
    ScatterMarkerColorbarShowtickprefix_none ScatterMarkerColorbarShowtickprefix = "none" // none
)

// ScatterMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterMarkerColorbarShowticksuffix string

const (
    ScatterMarkerColorbarShowticksuffix_all ScatterMarkerColorbarShowticksuffix = "all" // all
    ScatterMarkerColorbarShowticksuffix_first ScatterMarkerColorbarShowticksuffix = "first" // first
    ScatterMarkerColorbarShowticksuffix_last ScatterMarkerColorbarShowticksuffix = "last" // last
    ScatterMarkerColorbarShowticksuffix_none ScatterMarkerColorbarShowticksuffix = "none" // none
)

// ScatterMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterMarkerColorbarThicknessmode string

const (
    ScatterMarkerColorbarThicknessmode_fraction ScatterMarkerColorbarThicknessmode = "fraction" // fraction
    ScatterMarkerColorbarThicknessmode_pixels ScatterMarkerColorbarThicknessmode = "pixels" // pixels
)

// ScatterMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterMarkerColorbarTickmode string

const (
    ScatterMarkerColorbarTickmode_auto ScatterMarkerColorbarTickmode = "auto" // auto
    ScatterMarkerColorbarTickmode_linear ScatterMarkerColorbarTickmode = "linear" // linear
    ScatterMarkerColorbarTickmode_array ScatterMarkerColorbarTickmode = "array" // array
)

// ScatterMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterMarkerColorbarTicks string

const (
    ScatterMarkerColorbarTicks_outside ScatterMarkerColorbarTicks = "outside" // outside
    ScatterMarkerColorbarTicks_inside ScatterMarkerColorbarTicks = "inside" // inside
)

// ScatterMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterMarkerColorbarTitleSide string

const (
    ScatterMarkerColorbarTitleSide_right ScatterMarkerColorbarTitleSide = "right" // right
    ScatterMarkerColorbarTitleSide_top ScatterMarkerColorbarTitleSide = "top" // top
    ScatterMarkerColorbarTitleSide_bottom ScatterMarkerColorbarTitleSide = "bottom" // bottom
)

// ScatterMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterMarkerColorbarXanchor string

const (
    ScatterMarkerColorbarXanchor_left ScatterMarkerColorbarXanchor = "left" // left
    ScatterMarkerColorbarXanchor_center ScatterMarkerColorbarXanchor = "center" // center
    ScatterMarkerColorbarXanchor_right ScatterMarkerColorbarXanchor = "right" // right
)

// ScatterMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterMarkerColorbarYanchor string

const (
    ScatterMarkerColorbarYanchor_top ScatterMarkerColorbarYanchor = "top" // top
    ScatterMarkerColorbarYanchor_middle ScatterMarkerColorbarYanchor = "middle" // middle
    ScatterMarkerColorbarYanchor_bottom ScatterMarkerColorbarYanchor = "bottom" // bottom
)

// ScatterMarkerGradientType Sets the type of gradient used to fill the markers
type ScatterMarkerGradientType string

const (
    ScatterMarkerGradientType_radial ScatterMarkerGradientType = "radial" // radial
    ScatterMarkerGradientType_horizontal ScatterMarkerGradientType = "horizontal" // horizontal
    ScatterMarkerGradientType_vertical ScatterMarkerGradientType = "vertical" // vertical
    ScatterMarkerGradientType_none ScatterMarkerGradientType = "none" // none
)

// ScatterMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterMarkerSizemode string

const (
    ScatterMarkerSizemode_diameter ScatterMarkerSizemode = "diameter" // diameter
    ScatterMarkerSizemode_area ScatterMarkerSizemode = "area" // area
)

// ScatterMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterMarkerSymbol string

const (
    ScatterMarkerSymbol0 ScatterMarkerSymbol = "0" // 0
    ScatterMarkerSymbol_circle ScatterMarkerSymbol = "circle" // circle
    ScatterMarkerSymbol100 ScatterMarkerSymbol = "100" // 100
    ScatterMarkerSymbol_circle_open ScatterMarkerSymbol = "circle-open" // circle-open
    ScatterMarkerSymbol200 ScatterMarkerSymbol = "200" // 200
    ScatterMarkerSymbol_circle_dot ScatterMarkerSymbol = "circle-dot" // circle-dot
    ScatterMarkerSymbol300 ScatterMarkerSymbol = "300" // 300
    ScatterMarkerSymbol_circle_open_dot ScatterMarkerSymbol = "circle-open-dot" // circle-open-dot
    ScatterMarkerSymbol1 ScatterMarkerSymbol = "1" // 1
    ScatterMarkerSymbol_square ScatterMarkerSymbol = "square" // square
    ScatterMarkerSymbol101 ScatterMarkerSymbol = "101" // 101
    ScatterMarkerSymbol_square_open ScatterMarkerSymbol = "square-open" // square-open
    ScatterMarkerSymbol201 ScatterMarkerSymbol = "201" // 201
    ScatterMarkerSymbol_square_dot ScatterMarkerSymbol = "square-dot" // square-dot
    ScatterMarkerSymbol301 ScatterMarkerSymbol = "301" // 301
    ScatterMarkerSymbol_square_open_dot ScatterMarkerSymbol = "square-open-dot" // square-open-dot
    ScatterMarkerSymbol2 ScatterMarkerSymbol = "2" // 2
    ScatterMarkerSymbol_diamond ScatterMarkerSymbol = "diamond" // diamond
    ScatterMarkerSymbol102 ScatterMarkerSymbol = "102" // 102
    ScatterMarkerSymbol_diamond_open ScatterMarkerSymbol = "diamond-open" // diamond-open
    ScatterMarkerSymbol202 ScatterMarkerSymbol = "202" // 202
    ScatterMarkerSymbol_diamond_dot ScatterMarkerSymbol = "diamond-dot" // diamond-dot
    ScatterMarkerSymbol302 ScatterMarkerSymbol = "302" // 302
    ScatterMarkerSymbol_diamond_open_dot ScatterMarkerSymbol = "diamond-open-dot" // diamond-open-dot
    ScatterMarkerSymbol3 ScatterMarkerSymbol = "3" // 3
    ScatterMarkerSymbol_cross ScatterMarkerSymbol = "cross" // cross
    ScatterMarkerSymbol103 ScatterMarkerSymbol = "103" // 103
    ScatterMarkerSymbol_cross_open ScatterMarkerSymbol = "cross-open" // cross-open
    ScatterMarkerSymbol203 ScatterMarkerSymbol = "203" // 203
    ScatterMarkerSymbol_cross_dot ScatterMarkerSymbol = "cross-dot" // cross-dot
    ScatterMarkerSymbol303 ScatterMarkerSymbol = "303" // 303
    ScatterMarkerSymbol_cross_open_dot ScatterMarkerSymbol = "cross-open-dot" // cross-open-dot
    ScatterMarkerSymbol4 ScatterMarkerSymbol = "4" // 4
    ScatterMarkerSymbol_x ScatterMarkerSymbol = "x" // x
    ScatterMarkerSymbol104 ScatterMarkerSymbol = "104" // 104
    ScatterMarkerSymbol_x_open ScatterMarkerSymbol = "x-open" // x-open
    ScatterMarkerSymbol204 ScatterMarkerSymbol = "204" // 204
    ScatterMarkerSymbol_x_dot ScatterMarkerSymbol = "x-dot" // x-dot
    ScatterMarkerSymbol304 ScatterMarkerSymbol = "304" // 304
    ScatterMarkerSymbol_x_open_dot ScatterMarkerSymbol = "x-open-dot" // x-open-dot
    ScatterMarkerSymbol5 ScatterMarkerSymbol = "5" // 5
    ScatterMarkerSymbol_triangle_up ScatterMarkerSymbol = "triangle-up" // triangle-up
    ScatterMarkerSymbol105 ScatterMarkerSymbol = "105" // 105
    ScatterMarkerSymbol_triangle_up_open ScatterMarkerSymbol = "triangle-up-open" // triangle-up-open
    ScatterMarkerSymbol205 ScatterMarkerSymbol = "205" // 205
    ScatterMarkerSymbol_triangle_up_dot ScatterMarkerSymbol = "triangle-up-dot" // triangle-up-dot
    ScatterMarkerSymbol305 ScatterMarkerSymbol = "305" // 305
    ScatterMarkerSymbol_triangle_up_open_dot ScatterMarkerSymbol = "triangle-up-open-dot" // triangle-up-open-dot
    ScatterMarkerSymbol6 ScatterMarkerSymbol = "6" // 6
    ScatterMarkerSymbol_triangle_down ScatterMarkerSymbol = "triangle-down" // triangle-down
    ScatterMarkerSymbol106 ScatterMarkerSymbol = "106" // 106
    ScatterMarkerSymbol_triangle_down_open ScatterMarkerSymbol = "triangle-down-open" // triangle-down-open
    ScatterMarkerSymbol206 ScatterMarkerSymbol = "206" // 206
    ScatterMarkerSymbol_triangle_down_dot ScatterMarkerSymbol = "triangle-down-dot" // triangle-down-dot
    ScatterMarkerSymbol306 ScatterMarkerSymbol = "306" // 306
    ScatterMarkerSymbol_triangle_down_open_dot ScatterMarkerSymbol = "triangle-down-open-dot" // triangle-down-open-dot
    ScatterMarkerSymbol7 ScatterMarkerSymbol = "7" // 7
    ScatterMarkerSymbol_triangle_left ScatterMarkerSymbol = "triangle-left" // triangle-left
    ScatterMarkerSymbol107 ScatterMarkerSymbol = "107" // 107
    ScatterMarkerSymbol_triangle_left_open ScatterMarkerSymbol = "triangle-left-open" // triangle-left-open
    ScatterMarkerSymbol207 ScatterMarkerSymbol = "207" // 207
    ScatterMarkerSymbol_triangle_left_dot ScatterMarkerSymbol = "triangle-left-dot" // triangle-left-dot
    ScatterMarkerSymbol307 ScatterMarkerSymbol = "307" // 307
    ScatterMarkerSymbol_triangle_left_open_dot ScatterMarkerSymbol = "triangle-left-open-dot" // triangle-left-open-dot
    ScatterMarkerSymbol8 ScatterMarkerSymbol = "8" // 8
    ScatterMarkerSymbol_triangle_right ScatterMarkerSymbol = "triangle-right" // triangle-right
    ScatterMarkerSymbol108 ScatterMarkerSymbol = "108" // 108
    ScatterMarkerSymbol_triangle_right_open ScatterMarkerSymbol = "triangle-right-open" // triangle-right-open
    ScatterMarkerSymbol208 ScatterMarkerSymbol = "208" // 208
    ScatterMarkerSymbol_triangle_right_dot ScatterMarkerSymbol = "triangle-right-dot" // triangle-right-dot
    ScatterMarkerSymbol308 ScatterMarkerSymbol = "308" // 308
    ScatterMarkerSymbol_triangle_right_open_dot ScatterMarkerSymbol = "triangle-right-open-dot" // triangle-right-open-dot
    ScatterMarkerSymbol9 ScatterMarkerSymbol = "9" // 9
    ScatterMarkerSymbol_triangle_ne ScatterMarkerSymbol = "triangle-ne" // triangle-ne
    ScatterMarkerSymbol109 ScatterMarkerSymbol = "109" // 109
    ScatterMarkerSymbol_triangle_ne_open ScatterMarkerSymbol = "triangle-ne-open" // triangle-ne-open
    ScatterMarkerSymbol209 ScatterMarkerSymbol = "209" // 209
    ScatterMarkerSymbol_triangle_ne_dot ScatterMarkerSymbol = "triangle-ne-dot" // triangle-ne-dot
    ScatterMarkerSymbol309 ScatterMarkerSymbol = "309" // 309
    ScatterMarkerSymbol_triangle_ne_open_dot ScatterMarkerSymbol = "triangle-ne-open-dot" // triangle-ne-open-dot
    ScatterMarkerSymbol10 ScatterMarkerSymbol = "10" // 10
    ScatterMarkerSymbol_triangle_se ScatterMarkerSymbol = "triangle-se" // triangle-se
    ScatterMarkerSymbol110 ScatterMarkerSymbol = "110" // 110
    ScatterMarkerSymbol_triangle_se_open ScatterMarkerSymbol = "triangle-se-open" // triangle-se-open
    ScatterMarkerSymbol210 ScatterMarkerSymbol = "210" // 210
    ScatterMarkerSymbol_triangle_se_dot ScatterMarkerSymbol = "triangle-se-dot" // triangle-se-dot
    ScatterMarkerSymbol310 ScatterMarkerSymbol = "310" // 310
    ScatterMarkerSymbol_triangle_se_open_dot ScatterMarkerSymbol = "triangle-se-open-dot" // triangle-se-open-dot
    ScatterMarkerSymbol11 ScatterMarkerSymbol = "11" // 11
    ScatterMarkerSymbol_triangle_sw ScatterMarkerSymbol = "triangle-sw" // triangle-sw
    ScatterMarkerSymbol111 ScatterMarkerSymbol = "111" // 111
    ScatterMarkerSymbol_triangle_sw_open ScatterMarkerSymbol = "triangle-sw-open" // triangle-sw-open
    ScatterMarkerSymbol211 ScatterMarkerSymbol = "211" // 211
    ScatterMarkerSymbol_triangle_sw_dot ScatterMarkerSymbol = "triangle-sw-dot" // triangle-sw-dot
    ScatterMarkerSymbol311 ScatterMarkerSymbol = "311" // 311
    ScatterMarkerSymbol_triangle_sw_open_dot ScatterMarkerSymbol = "triangle-sw-open-dot" // triangle-sw-open-dot
    ScatterMarkerSymbol12 ScatterMarkerSymbol = "12" // 12
    ScatterMarkerSymbol_triangle_nw ScatterMarkerSymbol = "triangle-nw" // triangle-nw
    ScatterMarkerSymbol112 ScatterMarkerSymbol = "112" // 112
    ScatterMarkerSymbol_triangle_nw_open ScatterMarkerSymbol = "triangle-nw-open" // triangle-nw-open
    ScatterMarkerSymbol212 ScatterMarkerSymbol = "212" // 212
    ScatterMarkerSymbol_triangle_nw_dot ScatterMarkerSymbol = "triangle-nw-dot" // triangle-nw-dot
    ScatterMarkerSymbol312 ScatterMarkerSymbol = "312" // 312
    ScatterMarkerSymbol_triangle_nw_open_dot ScatterMarkerSymbol = "triangle-nw-open-dot" // triangle-nw-open-dot
    ScatterMarkerSymbol13 ScatterMarkerSymbol = "13" // 13
    ScatterMarkerSymbol_pentagon ScatterMarkerSymbol = "pentagon" // pentagon
    ScatterMarkerSymbol113 ScatterMarkerSymbol = "113" // 113
    ScatterMarkerSymbol_pentagon_open ScatterMarkerSymbol = "pentagon-open" // pentagon-open
    ScatterMarkerSymbol213 ScatterMarkerSymbol = "213" // 213
    ScatterMarkerSymbol_pentagon_dot ScatterMarkerSymbol = "pentagon-dot" // pentagon-dot
    ScatterMarkerSymbol313 ScatterMarkerSymbol = "313" // 313
    ScatterMarkerSymbol_pentagon_open_dot ScatterMarkerSymbol = "pentagon-open-dot" // pentagon-open-dot
    ScatterMarkerSymbol14 ScatterMarkerSymbol = "14" // 14
    ScatterMarkerSymbol_hexagon ScatterMarkerSymbol = "hexagon" // hexagon
    ScatterMarkerSymbol114 ScatterMarkerSymbol = "114" // 114
    ScatterMarkerSymbol_hexagon_open ScatterMarkerSymbol = "hexagon-open" // hexagon-open
    ScatterMarkerSymbol214 ScatterMarkerSymbol = "214" // 214
    ScatterMarkerSymbol_hexagon_dot ScatterMarkerSymbol = "hexagon-dot" // hexagon-dot
    ScatterMarkerSymbol314 ScatterMarkerSymbol = "314" // 314
    ScatterMarkerSymbol_hexagon_open_dot ScatterMarkerSymbol = "hexagon-open-dot" // hexagon-open-dot
    ScatterMarkerSymbol15 ScatterMarkerSymbol = "15" // 15
    ScatterMarkerSymbol_hexagon2 ScatterMarkerSymbol = "hexagon2" // hexagon2
    ScatterMarkerSymbol115 ScatterMarkerSymbol = "115" // 115
    ScatterMarkerSymbol_hexagon2_open ScatterMarkerSymbol = "hexagon2-open" // hexagon2-open
    ScatterMarkerSymbol215 ScatterMarkerSymbol = "215" // 215
    ScatterMarkerSymbol_hexagon2_dot ScatterMarkerSymbol = "hexagon2-dot" // hexagon2-dot
    ScatterMarkerSymbol315 ScatterMarkerSymbol = "315" // 315
    ScatterMarkerSymbol_hexagon2_open_dot ScatterMarkerSymbol = "hexagon2-open-dot" // hexagon2-open-dot
    ScatterMarkerSymbol16 ScatterMarkerSymbol = "16" // 16
    ScatterMarkerSymbol_octagon ScatterMarkerSymbol = "octagon" // octagon
    ScatterMarkerSymbol116 ScatterMarkerSymbol = "116" // 116
    ScatterMarkerSymbol_octagon_open ScatterMarkerSymbol = "octagon-open" // octagon-open
    ScatterMarkerSymbol216 ScatterMarkerSymbol = "216" // 216
    ScatterMarkerSymbol_octagon_dot ScatterMarkerSymbol = "octagon-dot" // octagon-dot
    ScatterMarkerSymbol316 ScatterMarkerSymbol = "316" // 316
    ScatterMarkerSymbol_octagon_open_dot ScatterMarkerSymbol = "octagon-open-dot" // octagon-open-dot
    ScatterMarkerSymbol17 ScatterMarkerSymbol = "17" // 17
    ScatterMarkerSymbol_star ScatterMarkerSymbol = "star" // star
    ScatterMarkerSymbol117 ScatterMarkerSymbol = "117" // 117
    ScatterMarkerSymbol_star_open ScatterMarkerSymbol = "star-open" // star-open
    ScatterMarkerSymbol217 ScatterMarkerSymbol = "217" // 217
    ScatterMarkerSymbol_star_dot ScatterMarkerSymbol = "star-dot" // star-dot
    ScatterMarkerSymbol317 ScatterMarkerSymbol = "317" // 317
    ScatterMarkerSymbol_star_open_dot ScatterMarkerSymbol = "star-open-dot" // star-open-dot
    ScatterMarkerSymbol18 ScatterMarkerSymbol = "18" // 18
    ScatterMarkerSymbol_hexagram ScatterMarkerSymbol = "hexagram" // hexagram
    ScatterMarkerSymbol118 ScatterMarkerSymbol = "118" // 118
    ScatterMarkerSymbol_hexagram_open ScatterMarkerSymbol = "hexagram-open" // hexagram-open
    ScatterMarkerSymbol218 ScatterMarkerSymbol = "218" // 218
    ScatterMarkerSymbol_hexagram_dot ScatterMarkerSymbol = "hexagram-dot" // hexagram-dot
    ScatterMarkerSymbol318 ScatterMarkerSymbol = "318" // 318
    ScatterMarkerSymbol_hexagram_open_dot ScatterMarkerSymbol = "hexagram-open-dot" // hexagram-open-dot
    ScatterMarkerSymbol19 ScatterMarkerSymbol = "19" // 19
    ScatterMarkerSymbol_star_triangle_up ScatterMarkerSymbol = "star-triangle-up" // star-triangle-up
    ScatterMarkerSymbol119 ScatterMarkerSymbol = "119" // 119
    ScatterMarkerSymbol_star_triangle_up_open ScatterMarkerSymbol = "star-triangle-up-open" // star-triangle-up-open
    ScatterMarkerSymbol219 ScatterMarkerSymbol = "219" // 219
    ScatterMarkerSymbol_star_triangle_up_dot ScatterMarkerSymbol = "star-triangle-up-dot" // star-triangle-up-dot
    ScatterMarkerSymbol319 ScatterMarkerSymbol = "319" // 319
    ScatterMarkerSymbol_star_triangle_up_open_dot ScatterMarkerSymbol = "star-triangle-up-open-dot" // star-triangle-up-open-dot
    ScatterMarkerSymbol20 ScatterMarkerSymbol = "20" // 20
    ScatterMarkerSymbol_star_triangle_down ScatterMarkerSymbol = "star-triangle-down" // star-triangle-down
    ScatterMarkerSymbol120 ScatterMarkerSymbol = "120" // 120
    ScatterMarkerSymbol_star_triangle_down_open ScatterMarkerSymbol = "star-triangle-down-open" // star-triangle-down-open
    ScatterMarkerSymbol220 ScatterMarkerSymbol = "220" // 220
    ScatterMarkerSymbol_star_triangle_down_dot ScatterMarkerSymbol = "star-triangle-down-dot" // star-triangle-down-dot
    ScatterMarkerSymbol320 ScatterMarkerSymbol = "320" // 320
    ScatterMarkerSymbol_star_triangle_down_open_dot ScatterMarkerSymbol = "star-triangle-down-open-dot" // star-triangle-down-open-dot
    ScatterMarkerSymbol21 ScatterMarkerSymbol = "21" // 21
    ScatterMarkerSymbol_star_square ScatterMarkerSymbol = "star-square" // star-square
    ScatterMarkerSymbol121 ScatterMarkerSymbol = "121" // 121
    ScatterMarkerSymbol_star_square_open ScatterMarkerSymbol = "star-square-open" // star-square-open
    ScatterMarkerSymbol221 ScatterMarkerSymbol = "221" // 221
    ScatterMarkerSymbol_star_square_dot ScatterMarkerSymbol = "star-square-dot" // star-square-dot
    ScatterMarkerSymbol321 ScatterMarkerSymbol = "321" // 321
    ScatterMarkerSymbol_star_square_open_dot ScatterMarkerSymbol = "star-square-open-dot" // star-square-open-dot
    ScatterMarkerSymbol22 ScatterMarkerSymbol = "22" // 22
    ScatterMarkerSymbol_star_diamond ScatterMarkerSymbol = "star-diamond" // star-diamond
    ScatterMarkerSymbol122 ScatterMarkerSymbol = "122" // 122
    ScatterMarkerSymbol_star_diamond_open ScatterMarkerSymbol = "star-diamond-open" // star-diamond-open
    ScatterMarkerSymbol222 ScatterMarkerSymbol = "222" // 222
    ScatterMarkerSymbol_star_diamond_dot ScatterMarkerSymbol = "star-diamond-dot" // star-diamond-dot
    ScatterMarkerSymbol322 ScatterMarkerSymbol = "322" // 322
    ScatterMarkerSymbol_star_diamond_open_dot ScatterMarkerSymbol = "star-diamond-open-dot" // star-diamond-open-dot
    ScatterMarkerSymbol23 ScatterMarkerSymbol = "23" // 23
    ScatterMarkerSymbol_diamond_tall ScatterMarkerSymbol = "diamond-tall" // diamond-tall
    ScatterMarkerSymbol123 ScatterMarkerSymbol = "123" // 123
    ScatterMarkerSymbol_diamond_tall_open ScatterMarkerSymbol = "diamond-tall-open" // diamond-tall-open
    ScatterMarkerSymbol223 ScatterMarkerSymbol = "223" // 223
    ScatterMarkerSymbol_diamond_tall_dot ScatterMarkerSymbol = "diamond-tall-dot" // diamond-tall-dot
    ScatterMarkerSymbol323 ScatterMarkerSymbol = "323" // 323
    ScatterMarkerSymbol_diamond_tall_open_dot ScatterMarkerSymbol = "diamond-tall-open-dot" // diamond-tall-open-dot
    ScatterMarkerSymbol24 ScatterMarkerSymbol = "24" // 24
    ScatterMarkerSymbol_diamond_wide ScatterMarkerSymbol = "diamond-wide" // diamond-wide
    ScatterMarkerSymbol124 ScatterMarkerSymbol = "124" // 124
    ScatterMarkerSymbol_diamond_wide_open ScatterMarkerSymbol = "diamond-wide-open" // diamond-wide-open
    ScatterMarkerSymbol224 ScatterMarkerSymbol = "224" // 224
    ScatterMarkerSymbol_diamond_wide_dot ScatterMarkerSymbol = "diamond-wide-dot" // diamond-wide-dot
    ScatterMarkerSymbol324 ScatterMarkerSymbol = "324" // 324
    ScatterMarkerSymbol_diamond_wide_open_dot ScatterMarkerSymbol = "diamond-wide-open-dot" // diamond-wide-open-dot
    ScatterMarkerSymbol25 ScatterMarkerSymbol = "25" // 25
    ScatterMarkerSymbol_hourglass ScatterMarkerSymbol = "hourglass" // hourglass
    ScatterMarkerSymbol125 ScatterMarkerSymbol = "125" // 125
    ScatterMarkerSymbol_hourglass_open ScatterMarkerSymbol = "hourglass-open" // hourglass-open
    ScatterMarkerSymbol26 ScatterMarkerSymbol = "26" // 26
    ScatterMarkerSymbol_bowtie ScatterMarkerSymbol = "bowtie" // bowtie
    ScatterMarkerSymbol126 ScatterMarkerSymbol = "126" // 126
    ScatterMarkerSymbol_bowtie_open ScatterMarkerSymbol = "bowtie-open" // bowtie-open
    ScatterMarkerSymbol27 ScatterMarkerSymbol = "27" // 27
    ScatterMarkerSymbol_circle_cross ScatterMarkerSymbol = "circle-cross" // circle-cross
    ScatterMarkerSymbol127 ScatterMarkerSymbol = "127" // 127
    ScatterMarkerSymbol_circle_cross_open ScatterMarkerSymbol = "circle-cross-open" // circle-cross-open
    ScatterMarkerSymbol28 ScatterMarkerSymbol = "28" // 28
    ScatterMarkerSymbol_circle_x ScatterMarkerSymbol = "circle-x" // circle-x
    ScatterMarkerSymbol128 ScatterMarkerSymbol = "128" // 128
    ScatterMarkerSymbol_circle_x_open ScatterMarkerSymbol = "circle-x-open" // circle-x-open
    ScatterMarkerSymbol29 ScatterMarkerSymbol = "29" // 29
    ScatterMarkerSymbol_square_cross ScatterMarkerSymbol = "square-cross" // square-cross
    ScatterMarkerSymbol129 ScatterMarkerSymbol = "129" // 129
    ScatterMarkerSymbol_square_cross_open ScatterMarkerSymbol = "square-cross-open" // square-cross-open
    ScatterMarkerSymbol30 ScatterMarkerSymbol = "30" // 30
    ScatterMarkerSymbol_square_x ScatterMarkerSymbol = "square-x" // square-x
    ScatterMarkerSymbol130 ScatterMarkerSymbol = "130" // 130
    ScatterMarkerSymbol_square_x_open ScatterMarkerSymbol = "square-x-open" // square-x-open
    ScatterMarkerSymbol31 ScatterMarkerSymbol = "31" // 31
    ScatterMarkerSymbol_diamond_cross ScatterMarkerSymbol = "diamond-cross" // diamond-cross
    ScatterMarkerSymbol131 ScatterMarkerSymbol = "131" // 131
    ScatterMarkerSymbol_diamond_cross_open ScatterMarkerSymbol = "diamond-cross-open" // diamond-cross-open
    ScatterMarkerSymbol32 ScatterMarkerSymbol = "32" // 32
    ScatterMarkerSymbol_diamond_x ScatterMarkerSymbol = "diamond-x" // diamond-x
    ScatterMarkerSymbol132 ScatterMarkerSymbol = "132" // 132
    ScatterMarkerSymbol_diamond_x_open ScatterMarkerSymbol = "diamond-x-open" // diamond-x-open
    ScatterMarkerSymbol33 ScatterMarkerSymbol = "33" // 33
    ScatterMarkerSymbol_cross_thin ScatterMarkerSymbol = "cross-thin" // cross-thin
    ScatterMarkerSymbol133 ScatterMarkerSymbol = "133" // 133
    ScatterMarkerSymbol_cross_thin_open ScatterMarkerSymbol = "cross-thin-open" // cross-thin-open
    ScatterMarkerSymbol34 ScatterMarkerSymbol = "34" // 34
    ScatterMarkerSymbol_x_thin ScatterMarkerSymbol = "x-thin" // x-thin
    ScatterMarkerSymbol134 ScatterMarkerSymbol = "134" // 134
    ScatterMarkerSymbol_x_thin_open ScatterMarkerSymbol = "x-thin-open" // x-thin-open
    ScatterMarkerSymbol35 ScatterMarkerSymbol = "35" // 35
    ScatterMarkerSymbol_asterisk ScatterMarkerSymbol = "asterisk" // asterisk
    ScatterMarkerSymbol135 ScatterMarkerSymbol = "135" // 135
    ScatterMarkerSymbol_asterisk_open ScatterMarkerSymbol = "asterisk-open" // asterisk-open
    ScatterMarkerSymbol36 ScatterMarkerSymbol = "36" // 36
    ScatterMarkerSymbol_hash ScatterMarkerSymbol = "hash" // hash
    ScatterMarkerSymbol136 ScatterMarkerSymbol = "136" // 136
    ScatterMarkerSymbol_hash_open ScatterMarkerSymbol = "hash-open" // hash-open
    ScatterMarkerSymbol236 ScatterMarkerSymbol = "236" // 236
    ScatterMarkerSymbol_hash_dot ScatterMarkerSymbol = "hash-dot" // hash-dot
    ScatterMarkerSymbol336 ScatterMarkerSymbol = "336" // 336
    ScatterMarkerSymbol_hash_open_dot ScatterMarkerSymbol = "hash-open-dot" // hash-open-dot
    ScatterMarkerSymbol37 ScatterMarkerSymbol = "37" // 37
    ScatterMarkerSymbol_y_up ScatterMarkerSymbol = "y-up" // y-up
    ScatterMarkerSymbol137 ScatterMarkerSymbol = "137" // 137
    ScatterMarkerSymbol_y_up_open ScatterMarkerSymbol = "y-up-open" // y-up-open
    ScatterMarkerSymbol38 ScatterMarkerSymbol = "38" // 38
    ScatterMarkerSymbol_y_down ScatterMarkerSymbol = "y-down" // y-down
    ScatterMarkerSymbol138 ScatterMarkerSymbol = "138" // 138
    ScatterMarkerSymbol_y_down_open ScatterMarkerSymbol = "y-down-open" // y-down-open
    ScatterMarkerSymbol39 ScatterMarkerSymbol = "39" // 39
    ScatterMarkerSymbol_y_left ScatterMarkerSymbol = "y-left" // y-left
    ScatterMarkerSymbol139 ScatterMarkerSymbol = "139" // 139
    ScatterMarkerSymbol_y_left_open ScatterMarkerSymbol = "y-left-open" // y-left-open
    ScatterMarkerSymbol40 ScatterMarkerSymbol = "40" // 40
    ScatterMarkerSymbol_y_right ScatterMarkerSymbol = "y-right" // y-right
    ScatterMarkerSymbol140 ScatterMarkerSymbol = "140" // 140
    ScatterMarkerSymbol_y_right_open ScatterMarkerSymbol = "y-right-open" // y-right-open
    ScatterMarkerSymbol41 ScatterMarkerSymbol = "41" // 41
    ScatterMarkerSymbol_line_ew ScatterMarkerSymbol = "line-ew" // line-ew
    ScatterMarkerSymbol141 ScatterMarkerSymbol = "141" // 141
    ScatterMarkerSymbol_line_ew_open ScatterMarkerSymbol = "line-ew-open" // line-ew-open
    ScatterMarkerSymbol42 ScatterMarkerSymbol = "42" // 42
    ScatterMarkerSymbol_line_ns ScatterMarkerSymbol = "line-ns" // line-ns
    ScatterMarkerSymbol142 ScatterMarkerSymbol = "142" // 142
    ScatterMarkerSymbol_line_ns_open ScatterMarkerSymbol = "line-ns-open" // line-ns-open
    ScatterMarkerSymbol43 ScatterMarkerSymbol = "43" // 43
    ScatterMarkerSymbol_line_ne ScatterMarkerSymbol = "line-ne" // line-ne
    ScatterMarkerSymbol143 ScatterMarkerSymbol = "143" // 143
    ScatterMarkerSymbol_line_ne_open ScatterMarkerSymbol = "line-ne-open" // line-ne-open
    ScatterMarkerSymbol44 ScatterMarkerSymbol = "44" // 44
    ScatterMarkerSymbol_line_nw ScatterMarkerSymbol = "line-nw" // line-nw
    ScatterMarkerSymbol144 ScatterMarkerSymbol = "144" // 144
    ScatterMarkerSymbol_line_nw_open ScatterMarkerSymbol = "line-nw-open" // line-nw-open
)

// ScatterOrientation Only relevant when `stackgroup` is used, and only the first `orientation` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Sets the stacking direction. With *v* (*h*), the y (x) values of subsequent traces are added. Also affects the default value of `fill`.
type ScatterOrientation string

const (
    ScatterOrientation_v ScatterOrientation = "v" // v
    ScatterOrientation_h ScatterOrientation = "h" // h
)

// ScatterStackgaps Only relevant when `stackgroup` is used, and only the first `stackgaps` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Determines how we handle locations at which other traces in this group have data but this one does not. With *infer zero* we insert a zero at these locations. With *interpolate* we linearly interpolate between existing values, and extrapolate a constant beyond the existing values.
type ScatterStackgaps string

const (
    ScatterStackgaps_inferzero ScatterStackgaps = "infer zero" // infer zero
    ScatterStackgaps_interpolate ScatterStackgaps = "interpolate" // interpolate
)

// ScatterTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterTextposition string

const (
    ScatterTextposition_topleft ScatterTextposition = "top left" // top left
    ScatterTextposition_topcenter ScatterTextposition = "top center" // top center
    ScatterTextposition_topright ScatterTextposition = "top right" // top right
    ScatterTextposition_middleleft ScatterTextposition = "middle left" // middle left
    ScatterTextposition_middlecenter ScatterTextposition = "middle center" // middle center
    ScatterTextposition_middleright ScatterTextposition = "middle right" // middle right
    ScatterTextposition_bottomleft ScatterTextposition = "bottom left" // bottom left
    ScatterTextposition_bottomcenter ScatterTextposition = "bottom center" // bottom center
    ScatterTextposition_bottomright ScatterTextposition = "bottom right" // bottom right
)

// ScatterVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterVisible interface{}

const (
    ScatterVisibleTrue bool = true
    ScatterVisibleFalse bool = false
    ScatterVisibleLegendonly string = "legendonly"
)

// ScatterXcalendar Sets the calendar system to use with `x` date data.
type ScatterXcalendar string

const (
    ScatterXcalendar_gregorian ScatterXcalendar = "gregorian" // gregorian
    ScatterXcalendar_chinese ScatterXcalendar = "chinese" // chinese
    ScatterXcalendar_coptic ScatterXcalendar = "coptic" // coptic
    ScatterXcalendar_discworld ScatterXcalendar = "discworld" // discworld
    ScatterXcalendar_ethiopian ScatterXcalendar = "ethiopian" // ethiopian
    ScatterXcalendar_hebrew ScatterXcalendar = "hebrew" // hebrew
    ScatterXcalendar_islamic ScatterXcalendar = "islamic" // islamic
    ScatterXcalendar_julian ScatterXcalendar = "julian" // julian
    ScatterXcalendar_mayan ScatterXcalendar = "mayan" // mayan
    ScatterXcalendar_nanakshahi ScatterXcalendar = "nanakshahi" // nanakshahi
    ScatterXcalendar_nepali ScatterXcalendar = "nepali" // nepali
    ScatterXcalendar_persian ScatterXcalendar = "persian" // persian
    ScatterXcalendar_jalali ScatterXcalendar = "jalali" // jalali
    ScatterXcalendar_taiwan ScatterXcalendar = "taiwan" // taiwan
    ScatterXcalendar_thai ScatterXcalendar = "thai" // thai
    ScatterXcalendar_ummalqura ScatterXcalendar = "ummalqura" // ummalqura
)

// ScatterYcalendar Sets the calendar system to use with `y` date data.
type ScatterYcalendar string

const (
    ScatterYcalendar_gregorian ScatterYcalendar = "gregorian" // gregorian
    ScatterYcalendar_chinese ScatterYcalendar = "chinese" // chinese
    ScatterYcalendar_coptic ScatterYcalendar = "coptic" // coptic
    ScatterYcalendar_discworld ScatterYcalendar = "discworld" // discworld
    ScatterYcalendar_ethiopian ScatterYcalendar = "ethiopian" // ethiopian
    ScatterYcalendar_hebrew ScatterYcalendar = "hebrew" // hebrew
    ScatterYcalendar_islamic ScatterYcalendar = "islamic" // islamic
    ScatterYcalendar_julian ScatterYcalendar = "julian" // julian
    ScatterYcalendar_mayan ScatterYcalendar = "mayan" // mayan
    ScatterYcalendar_nanakshahi ScatterYcalendar = "nanakshahi" // nanakshahi
    ScatterYcalendar_nepali ScatterYcalendar = "nepali" // nepali
    ScatterYcalendar_persian ScatterYcalendar = "persian" // persian
    ScatterYcalendar_jalali ScatterYcalendar = "jalali" // jalali
    ScatterYcalendar_taiwan ScatterYcalendar = "taiwan" // taiwan
    ScatterYcalendar_thai ScatterYcalendar = "thai" // thai
    ScatterYcalendar_ummalqura ScatterYcalendar = "ummalqura" // ummalqura
)

// ScattercarpetFill Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. scatterternary has a subset of the options available to scatter. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other.
type ScattercarpetFill string

const (
    ScattercarpetFill_none ScattercarpetFill = "none" // none
    ScattercarpetFill_toself ScattercarpetFill = "toself" // toself
    ScattercarpetFill_tonext ScattercarpetFill = "tonext" // tonext
)

// ScattercarpetHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScattercarpetHoverlabelAlign string

const (
    ScattercarpetHoverlabelAlign_left ScattercarpetHoverlabelAlign = "left" // left
    ScattercarpetHoverlabelAlign_right ScattercarpetHoverlabelAlign = "right" // right
    ScattercarpetHoverlabelAlign_auto ScattercarpetHoverlabelAlign = "auto" // auto
)

// ScattercarpetLineShape Determines the line shape. With *spline* the lines are drawn using spline interpolation. The other available values correspond to step-wise line shapes.
type ScattercarpetLineShape string

const (
    ScattercarpetLineShape_linear ScattercarpetLineShape = "linear" // linear
    ScattercarpetLineShape_spline ScattercarpetLineShape = "spline" // spline
)

// ScattercarpetMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScattercarpetMarkerColorbarExponentformat string

const (
    ScattercarpetMarkerColorbarExponentformat_none ScattercarpetMarkerColorbarExponentformat = "none" // none
    ScattercarpetMarkerColorbarExponentformat_e ScattercarpetMarkerColorbarExponentformat = "e" // e
    ScattercarpetMarkerColorbarExponentformat_E ScattercarpetMarkerColorbarExponentformat = "E" // E
    ScattercarpetMarkerColorbarExponentformat_power ScattercarpetMarkerColorbarExponentformat = "power" // power
    ScattercarpetMarkerColorbarExponentformat_SI ScattercarpetMarkerColorbarExponentformat = "SI" // SI
    ScattercarpetMarkerColorbarExponentformat_B ScattercarpetMarkerColorbarExponentformat = "B" // B
)

// ScattercarpetMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScattercarpetMarkerColorbarLenmode string

const (
    ScattercarpetMarkerColorbarLenmode_fraction ScattercarpetMarkerColorbarLenmode = "fraction" // fraction
    ScattercarpetMarkerColorbarLenmode_pixels ScattercarpetMarkerColorbarLenmode = "pixels" // pixels
)

// ScattercarpetMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScattercarpetMarkerColorbarShowexponent string

const (
    ScattercarpetMarkerColorbarShowexponent_all ScattercarpetMarkerColorbarShowexponent = "all" // all
    ScattercarpetMarkerColorbarShowexponent_first ScattercarpetMarkerColorbarShowexponent = "first" // first
    ScattercarpetMarkerColorbarShowexponent_last ScattercarpetMarkerColorbarShowexponent = "last" // last
    ScattercarpetMarkerColorbarShowexponent_none ScattercarpetMarkerColorbarShowexponent = "none" // none
)

// ScattercarpetMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScattercarpetMarkerColorbarShowtickprefix string

const (
    ScattercarpetMarkerColorbarShowtickprefix_all ScattercarpetMarkerColorbarShowtickprefix = "all" // all
    ScattercarpetMarkerColorbarShowtickprefix_first ScattercarpetMarkerColorbarShowtickprefix = "first" // first
    ScattercarpetMarkerColorbarShowtickprefix_last ScattercarpetMarkerColorbarShowtickprefix = "last" // last
    ScattercarpetMarkerColorbarShowtickprefix_none ScattercarpetMarkerColorbarShowtickprefix = "none" // none
)

// ScattercarpetMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScattercarpetMarkerColorbarShowticksuffix string

const (
    ScattercarpetMarkerColorbarShowticksuffix_all ScattercarpetMarkerColorbarShowticksuffix = "all" // all
    ScattercarpetMarkerColorbarShowticksuffix_first ScattercarpetMarkerColorbarShowticksuffix = "first" // first
    ScattercarpetMarkerColorbarShowticksuffix_last ScattercarpetMarkerColorbarShowticksuffix = "last" // last
    ScattercarpetMarkerColorbarShowticksuffix_none ScattercarpetMarkerColorbarShowticksuffix = "none" // none
)

// ScattercarpetMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScattercarpetMarkerColorbarThicknessmode string

const (
    ScattercarpetMarkerColorbarThicknessmode_fraction ScattercarpetMarkerColorbarThicknessmode = "fraction" // fraction
    ScattercarpetMarkerColorbarThicknessmode_pixels ScattercarpetMarkerColorbarThicknessmode = "pixels" // pixels
)

// ScattercarpetMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScattercarpetMarkerColorbarTickmode string

const (
    ScattercarpetMarkerColorbarTickmode_auto ScattercarpetMarkerColorbarTickmode = "auto" // auto
    ScattercarpetMarkerColorbarTickmode_linear ScattercarpetMarkerColorbarTickmode = "linear" // linear
    ScattercarpetMarkerColorbarTickmode_array ScattercarpetMarkerColorbarTickmode = "array" // array
)

// ScattercarpetMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScattercarpetMarkerColorbarTicks string

const (
    ScattercarpetMarkerColorbarTicks_outside ScattercarpetMarkerColorbarTicks = "outside" // outside
    ScattercarpetMarkerColorbarTicks_inside ScattercarpetMarkerColorbarTicks = "inside" // inside
)

// ScattercarpetMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScattercarpetMarkerColorbarTitleSide string

const (
    ScattercarpetMarkerColorbarTitleSide_right ScattercarpetMarkerColorbarTitleSide = "right" // right
    ScattercarpetMarkerColorbarTitleSide_top ScattercarpetMarkerColorbarTitleSide = "top" // top
    ScattercarpetMarkerColorbarTitleSide_bottom ScattercarpetMarkerColorbarTitleSide = "bottom" // bottom
)

// ScattercarpetMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScattercarpetMarkerColorbarXanchor string

const (
    ScattercarpetMarkerColorbarXanchor_left ScattercarpetMarkerColorbarXanchor = "left" // left
    ScattercarpetMarkerColorbarXanchor_center ScattercarpetMarkerColorbarXanchor = "center" // center
    ScattercarpetMarkerColorbarXanchor_right ScattercarpetMarkerColorbarXanchor = "right" // right
)

// ScattercarpetMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScattercarpetMarkerColorbarYanchor string

const (
    ScattercarpetMarkerColorbarYanchor_top ScattercarpetMarkerColorbarYanchor = "top" // top
    ScattercarpetMarkerColorbarYanchor_middle ScattercarpetMarkerColorbarYanchor = "middle" // middle
    ScattercarpetMarkerColorbarYanchor_bottom ScattercarpetMarkerColorbarYanchor = "bottom" // bottom
)

// ScattercarpetMarkerGradientType Sets the type of gradient used to fill the markers
type ScattercarpetMarkerGradientType string

const (
    ScattercarpetMarkerGradientType_radial ScattercarpetMarkerGradientType = "radial" // radial
    ScattercarpetMarkerGradientType_horizontal ScattercarpetMarkerGradientType = "horizontal" // horizontal
    ScattercarpetMarkerGradientType_vertical ScattercarpetMarkerGradientType = "vertical" // vertical
    ScattercarpetMarkerGradientType_none ScattercarpetMarkerGradientType = "none" // none
)

// ScattercarpetMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScattercarpetMarkerSizemode string

const (
    ScattercarpetMarkerSizemode_diameter ScattercarpetMarkerSizemode = "diameter" // diameter
    ScattercarpetMarkerSizemode_area ScattercarpetMarkerSizemode = "area" // area
)

// ScattercarpetMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScattercarpetMarkerSymbol string

const (
    ScattercarpetMarkerSymbol0 ScattercarpetMarkerSymbol = "0" // 0
    ScattercarpetMarkerSymbol_circle ScattercarpetMarkerSymbol = "circle" // circle
    ScattercarpetMarkerSymbol100 ScattercarpetMarkerSymbol = "100" // 100
    ScattercarpetMarkerSymbol_circle_open ScattercarpetMarkerSymbol = "circle-open" // circle-open
    ScattercarpetMarkerSymbol200 ScattercarpetMarkerSymbol = "200" // 200
    ScattercarpetMarkerSymbol_circle_dot ScattercarpetMarkerSymbol = "circle-dot" // circle-dot
    ScattercarpetMarkerSymbol300 ScattercarpetMarkerSymbol = "300" // 300
    ScattercarpetMarkerSymbol_circle_open_dot ScattercarpetMarkerSymbol = "circle-open-dot" // circle-open-dot
    ScattercarpetMarkerSymbol1 ScattercarpetMarkerSymbol = "1" // 1
    ScattercarpetMarkerSymbol_square ScattercarpetMarkerSymbol = "square" // square
    ScattercarpetMarkerSymbol101 ScattercarpetMarkerSymbol = "101" // 101
    ScattercarpetMarkerSymbol_square_open ScattercarpetMarkerSymbol = "square-open" // square-open
    ScattercarpetMarkerSymbol201 ScattercarpetMarkerSymbol = "201" // 201
    ScattercarpetMarkerSymbol_square_dot ScattercarpetMarkerSymbol = "square-dot" // square-dot
    ScattercarpetMarkerSymbol301 ScattercarpetMarkerSymbol = "301" // 301
    ScattercarpetMarkerSymbol_square_open_dot ScattercarpetMarkerSymbol = "square-open-dot" // square-open-dot
    ScattercarpetMarkerSymbol2 ScattercarpetMarkerSymbol = "2" // 2
    ScattercarpetMarkerSymbol_diamond ScattercarpetMarkerSymbol = "diamond" // diamond
    ScattercarpetMarkerSymbol102 ScattercarpetMarkerSymbol = "102" // 102
    ScattercarpetMarkerSymbol_diamond_open ScattercarpetMarkerSymbol = "diamond-open" // diamond-open
    ScattercarpetMarkerSymbol202 ScattercarpetMarkerSymbol = "202" // 202
    ScattercarpetMarkerSymbol_diamond_dot ScattercarpetMarkerSymbol = "diamond-dot" // diamond-dot
    ScattercarpetMarkerSymbol302 ScattercarpetMarkerSymbol = "302" // 302
    ScattercarpetMarkerSymbol_diamond_open_dot ScattercarpetMarkerSymbol = "diamond-open-dot" // diamond-open-dot
    ScattercarpetMarkerSymbol3 ScattercarpetMarkerSymbol = "3" // 3
    ScattercarpetMarkerSymbol_cross ScattercarpetMarkerSymbol = "cross" // cross
    ScattercarpetMarkerSymbol103 ScattercarpetMarkerSymbol = "103" // 103
    ScattercarpetMarkerSymbol_cross_open ScattercarpetMarkerSymbol = "cross-open" // cross-open
    ScattercarpetMarkerSymbol203 ScattercarpetMarkerSymbol = "203" // 203
    ScattercarpetMarkerSymbol_cross_dot ScattercarpetMarkerSymbol = "cross-dot" // cross-dot
    ScattercarpetMarkerSymbol303 ScattercarpetMarkerSymbol = "303" // 303
    ScattercarpetMarkerSymbol_cross_open_dot ScattercarpetMarkerSymbol = "cross-open-dot" // cross-open-dot
    ScattercarpetMarkerSymbol4 ScattercarpetMarkerSymbol = "4" // 4
    ScattercarpetMarkerSymbol_x ScattercarpetMarkerSymbol = "x" // x
    ScattercarpetMarkerSymbol104 ScattercarpetMarkerSymbol = "104" // 104
    ScattercarpetMarkerSymbol_x_open ScattercarpetMarkerSymbol = "x-open" // x-open
    ScattercarpetMarkerSymbol204 ScattercarpetMarkerSymbol = "204" // 204
    ScattercarpetMarkerSymbol_x_dot ScattercarpetMarkerSymbol = "x-dot" // x-dot
    ScattercarpetMarkerSymbol304 ScattercarpetMarkerSymbol = "304" // 304
    ScattercarpetMarkerSymbol_x_open_dot ScattercarpetMarkerSymbol = "x-open-dot" // x-open-dot
    ScattercarpetMarkerSymbol5 ScattercarpetMarkerSymbol = "5" // 5
    ScattercarpetMarkerSymbol_triangle_up ScattercarpetMarkerSymbol = "triangle-up" // triangle-up
    ScattercarpetMarkerSymbol105 ScattercarpetMarkerSymbol = "105" // 105
    ScattercarpetMarkerSymbol_triangle_up_open ScattercarpetMarkerSymbol = "triangle-up-open" // triangle-up-open
    ScattercarpetMarkerSymbol205 ScattercarpetMarkerSymbol = "205" // 205
    ScattercarpetMarkerSymbol_triangle_up_dot ScattercarpetMarkerSymbol = "triangle-up-dot" // triangle-up-dot
    ScattercarpetMarkerSymbol305 ScattercarpetMarkerSymbol = "305" // 305
    ScattercarpetMarkerSymbol_triangle_up_open_dot ScattercarpetMarkerSymbol = "triangle-up-open-dot" // triangle-up-open-dot
    ScattercarpetMarkerSymbol6 ScattercarpetMarkerSymbol = "6" // 6
    ScattercarpetMarkerSymbol_triangle_down ScattercarpetMarkerSymbol = "triangle-down" // triangle-down
    ScattercarpetMarkerSymbol106 ScattercarpetMarkerSymbol = "106" // 106
    ScattercarpetMarkerSymbol_triangle_down_open ScattercarpetMarkerSymbol = "triangle-down-open" // triangle-down-open
    ScattercarpetMarkerSymbol206 ScattercarpetMarkerSymbol = "206" // 206
    ScattercarpetMarkerSymbol_triangle_down_dot ScattercarpetMarkerSymbol = "triangle-down-dot" // triangle-down-dot
    ScattercarpetMarkerSymbol306 ScattercarpetMarkerSymbol = "306" // 306
    ScattercarpetMarkerSymbol_triangle_down_open_dot ScattercarpetMarkerSymbol = "triangle-down-open-dot" // triangle-down-open-dot
    ScattercarpetMarkerSymbol7 ScattercarpetMarkerSymbol = "7" // 7
    ScattercarpetMarkerSymbol_triangle_left ScattercarpetMarkerSymbol = "triangle-left" // triangle-left
    ScattercarpetMarkerSymbol107 ScattercarpetMarkerSymbol = "107" // 107
    ScattercarpetMarkerSymbol_triangle_left_open ScattercarpetMarkerSymbol = "triangle-left-open" // triangle-left-open
    ScattercarpetMarkerSymbol207 ScattercarpetMarkerSymbol = "207" // 207
    ScattercarpetMarkerSymbol_triangle_left_dot ScattercarpetMarkerSymbol = "triangle-left-dot" // triangle-left-dot
    ScattercarpetMarkerSymbol307 ScattercarpetMarkerSymbol = "307" // 307
    ScattercarpetMarkerSymbol_triangle_left_open_dot ScattercarpetMarkerSymbol = "triangle-left-open-dot" // triangle-left-open-dot
    ScattercarpetMarkerSymbol8 ScattercarpetMarkerSymbol = "8" // 8
    ScattercarpetMarkerSymbol_triangle_right ScattercarpetMarkerSymbol = "triangle-right" // triangle-right
    ScattercarpetMarkerSymbol108 ScattercarpetMarkerSymbol = "108" // 108
    ScattercarpetMarkerSymbol_triangle_right_open ScattercarpetMarkerSymbol = "triangle-right-open" // triangle-right-open
    ScattercarpetMarkerSymbol208 ScattercarpetMarkerSymbol = "208" // 208
    ScattercarpetMarkerSymbol_triangle_right_dot ScattercarpetMarkerSymbol = "triangle-right-dot" // triangle-right-dot
    ScattercarpetMarkerSymbol308 ScattercarpetMarkerSymbol = "308" // 308
    ScattercarpetMarkerSymbol_triangle_right_open_dot ScattercarpetMarkerSymbol = "triangle-right-open-dot" // triangle-right-open-dot
    ScattercarpetMarkerSymbol9 ScattercarpetMarkerSymbol = "9" // 9
    ScattercarpetMarkerSymbol_triangle_ne ScattercarpetMarkerSymbol = "triangle-ne" // triangle-ne
    ScattercarpetMarkerSymbol109 ScattercarpetMarkerSymbol = "109" // 109
    ScattercarpetMarkerSymbol_triangle_ne_open ScattercarpetMarkerSymbol = "triangle-ne-open" // triangle-ne-open
    ScattercarpetMarkerSymbol209 ScattercarpetMarkerSymbol = "209" // 209
    ScattercarpetMarkerSymbol_triangle_ne_dot ScattercarpetMarkerSymbol = "triangle-ne-dot" // triangle-ne-dot
    ScattercarpetMarkerSymbol309 ScattercarpetMarkerSymbol = "309" // 309
    ScattercarpetMarkerSymbol_triangle_ne_open_dot ScattercarpetMarkerSymbol = "triangle-ne-open-dot" // triangle-ne-open-dot
    ScattercarpetMarkerSymbol10 ScattercarpetMarkerSymbol = "10" // 10
    ScattercarpetMarkerSymbol_triangle_se ScattercarpetMarkerSymbol = "triangle-se" // triangle-se
    ScattercarpetMarkerSymbol110 ScattercarpetMarkerSymbol = "110" // 110
    ScattercarpetMarkerSymbol_triangle_se_open ScattercarpetMarkerSymbol = "triangle-se-open" // triangle-se-open
    ScattercarpetMarkerSymbol210 ScattercarpetMarkerSymbol = "210" // 210
    ScattercarpetMarkerSymbol_triangle_se_dot ScattercarpetMarkerSymbol = "triangle-se-dot" // triangle-se-dot
    ScattercarpetMarkerSymbol310 ScattercarpetMarkerSymbol = "310" // 310
    ScattercarpetMarkerSymbol_triangle_se_open_dot ScattercarpetMarkerSymbol = "triangle-se-open-dot" // triangle-se-open-dot
    ScattercarpetMarkerSymbol11 ScattercarpetMarkerSymbol = "11" // 11
    ScattercarpetMarkerSymbol_triangle_sw ScattercarpetMarkerSymbol = "triangle-sw" // triangle-sw
    ScattercarpetMarkerSymbol111 ScattercarpetMarkerSymbol = "111" // 111
    ScattercarpetMarkerSymbol_triangle_sw_open ScattercarpetMarkerSymbol = "triangle-sw-open" // triangle-sw-open
    ScattercarpetMarkerSymbol211 ScattercarpetMarkerSymbol = "211" // 211
    ScattercarpetMarkerSymbol_triangle_sw_dot ScattercarpetMarkerSymbol = "triangle-sw-dot" // triangle-sw-dot
    ScattercarpetMarkerSymbol311 ScattercarpetMarkerSymbol = "311" // 311
    ScattercarpetMarkerSymbol_triangle_sw_open_dot ScattercarpetMarkerSymbol = "triangle-sw-open-dot" // triangle-sw-open-dot
    ScattercarpetMarkerSymbol12 ScattercarpetMarkerSymbol = "12" // 12
    ScattercarpetMarkerSymbol_triangle_nw ScattercarpetMarkerSymbol = "triangle-nw" // triangle-nw
    ScattercarpetMarkerSymbol112 ScattercarpetMarkerSymbol = "112" // 112
    ScattercarpetMarkerSymbol_triangle_nw_open ScattercarpetMarkerSymbol = "triangle-nw-open" // triangle-nw-open
    ScattercarpetMarkerSymbol212 ScattercarpetMarkerSymbol = "212" // 212
    ScattercarpetMarkerSymbol_triangle_nw_dot ScattercarpetMarkerSymbol = "triangle-nw-dot" // triangle-nw-dot
    ScattercarpetMarkerSymbol312 ScattercarpetMarkerSymbol = "312" // 312
    ScattercarpetMarkerSymbol_triangle_nw_open_dot ScattercarpetMarkerSymbol = "triangle-nw-open-dot" // triangle-nw-open-dot
    ScattercarpetMarkerSymbol13 ScattercarpetMarkerSymbol = "13" // 13
    ScattercarpetMarkerSymbol_pentagon ScattercarpetMarkerSymbol = "pentagon" // pentagon
    ScattercarpetMarkerSymbol113 ScattercarpetMarkerSymbol = "113" // 113
    ScattercarpetMarkerSymbol_pentagon_open ScattercarpetMarkerSymbol = "pentagon-open" // pentagon-open
    ScattercarpetMarkerSymbol213 ScattercarpetMarkerSymbol = "213" // 213
    ScattercarpetMarkerSymbol_pentagon_dot ScattercarpetMarkerSymbol = "pentagon-dot" // pentagon-dot
    ScattercarpetMarkerSymbol313 ScattercarpetMarkerSymbol = "313" // 313
    ScattercarpetMarkerSymbol_pentagon_open_dot ScattercarpetMarkerSymbol = "pentagon-open-dot" // pentagon-open-dot
    ScattercarpetMarkerSymbol14 ScattercarpetMarkerSymbol = "14" // 14
    ScattercarpetMarkerSymbol_hexagon ScattercarpetMarkerSymbol = "hexagon" // hexagon
    ScattercarpetMarkerSymbol114 ScattercarpetMarkerSymbol = "114" // 114
    ScattercarpetMarkerSymbol_hexagon_open ScattercarpetMarkerSymbol = "hexagon-open" // hexagon-open
    ScattercarpetMarkerSymbol214 ScattercarpetMarkerSymbol = "214" // 214
    ScattercarpetMarkerSymbol_hexagon_dot ScattercarpetMarkerSymbol = "hexagon-dot" // hexagon-dot
    ScattercarpetMarkerSymbol314 ScattercarpetMarkerSymbol = "314" // 314
    ScattercarpetMarkerSymbol_hexagon_open_dot ScattercarpetMarkerSymbol = "hexagon-open-dot" // hexagon-open-dot
    ScattercarpetMarkerSymbol15 ScattercarpetMarkerSymbol = "15" // 15
    ScattercarpetMarkerSymbol_hexagon2 ScattercarpetMarkerSymbol = "hexagon2" // hexagon2
    ScattercarpetMarkerSymbol115 ScattercarpetMarkerSymbol = "115" // 115
    ScattercarpetMarkerSymbol_hexagon2_open ScattercarpetMarkerSymbol = "hexagon2-open" // hexagon2-open
    ScattercarpetMarkerSymbol215 ScattercarpetMarkerSymbol = "215" // 215
    ScattercarpetMarkerSymbol_hexagon2_dot ScattercarpetMarkerSymbol = "hexagon2-dot" // hexagon2-dot
    ScattercarpetMarkerSymbol315 ScattercarpetMarkerSymbol = "315" // 315
    ScattercarpetMarkerSymbol_hexagon2_open_dot ScattercarpetMarkerSymbol = "hexagon2-open-dot" // hexagon2-open-dot
    ScattercarpetMarkerSymbol16 ScattercarpetMarkerSymbol = "16" // 16
    ScattercarpetMarkerSymbol_octagon ScattercarpetMarkerSymbol = "octagon" // octagon
    ScattercarpetMarkerSymbol116 ScattercarpetMarkerSymbol = "116" // 116
    ScattercarpetMarkerSymbol_octagon_open ScattercarpetMarkerSymbol = "octagon-open" // octagon-open
    ScattercarpetMarkerSymbol216 ScattercarpetMarkerSymbol = "216" // 216
    ScattercarpetMarkerSymbol_octagon_dot ScattercarpetMarkerSymbol = "octagon-dot" // octagon-dot
    ScattercarpetMarkerSymbol316 ScattercarpetMarkerSymbol = "316" // 316
    ScattercarpetMarkerSymbol_octagon_open_dot ScattercarpetMarkerSymbol = "octagon-open-dot" // octagon-open-dot
    ScattercarpetMarkerSymbol17 ScattercarpetMarkerSymbol = "17" // 17
    ScattercarpetMarkerSymbol_star ScattercarpetMarkerSymbol = "star" // star
    ScattercarpetMarkerSymbol117 ScattercarpetMarkerSymbol = "117" // 117
    ScattercarpetMarkerSymbol_star_open ScattercarpetMarkerSymbol = "star-open" // star-open
    ScattercarpetMarkerSymbol217 ScattercarpetMarkerSymbol = "217" // 217
    ScattercarpetMarkerSymbol_star_dot ScattercarpetMarkerSymbol = "star-dot" // star-dot
    ScattercarpetMarkerSymbol317 ScattercarpetMarkerSymbol = "317" // 317
    ScattercarpetMarkerSymbol_star_open_dot ScattercarpetMarkerSymbol = "star-open-dot" // star-open-dot
    ScattercarpetMarkerSymbol18 ScattercarpetMarkerSymbol = "18" // 18
    ScattercarpetMarkerSymbol_hexagram ScattercarpetMarkerSymbol = "hexagram" // hexagram
    ScattercarpetMarkerSymbol118 ScattercarpetMarkerSymbol = "118" // 118
    ScattercarpetMarkerSymbol_hexagram_open ScattercarpetMarkerSymbol = "hexagram-open" // hexagram-open
    ScattercarpetMarkerSymbol218 ScattercarpetMarkerSymbol = "218" // 218
    ScattercarpetMarkerSymbol_hexagram_dot ScattercarpetMarkerSymbol = "hexagram-dot" // hexagram-dot
    ScattercarpetMarkerSymbol318 ScattercarpetMarkerSymbol = "318" // 318
    ScattercarpetMarkerSymbol_hexagram_open_dot ScattercarpetMarkerSymbol = "hexagram-open-dot" // hexagram-open-dot
    ScattercarpetMarkerSymbol19 ScattercarpetMarkerSymbol = "19" // 19
    ScattercarpetMarkerSymbol_star_triangle_up ScattercarpetMarkerSymbol = "star-triangle-up" // star-triangle-up
    ScattercarpetMarkerSymbol119 ScattercarpetMarkerSymbol = "119" // 119
    ScattercarpetMarkerSymbol_star_triangle_up_open ScattercarpetMarkerSymbol = "star-triangle-up-open" // star-triangle-up-open
    ScattercarpetMarkerSymbol219 ScattercarpetMarkerSymbol = "219" // 219
    ScattercarpetMarkerSymbol_star_triangle_up_dot ScattercarpetMarkerSymbol = "star-triangle-up-dot" // star-triangle-up-dot
    ScattercarpetMarkerSymbol319 ScattercarpetMarkerSymbol = "319" // 319
    ScattercarpetMarkerSymbol_star_triangle_up_open_dot ScattercarpetMarkerSymbol = "star-triangle-up-open-dot" // star-triangle-up-open-dot
    ScattercarpetMarkerSymbol20 ScattercarpetMarkerSymbol = "20" // 20
    ScattercarpetMarkerSymbol_star_triangle_down ScattercarpetMarkerSymbol = "star-triangle-down" // star-triangle-down
    ScattercarpetMarkerSymbol120 ScattercarpetMarkerSymbol = "120" // 120
    ScattercarpetMarkerSymbol_star_triangle_down_open ScattercarpetMarkerSymbol = "star-triangle-down-open" // star-triangle-down-open
    ScattercarpetMarkerSymbol220 ScattercarpetMarkerSymbol = "220" // 220
    ScattercarpetMarkerSymbol_star_triangle_down_dot ScattercarpetMarkerSymbol = "star-triangle-down-dot" // star-triangle-down-dot
    ScattercarpetMarkerSymbol320 ScattercarpetMarkerSymbol = "320" // 320
    ScattercarpetMarkerSymbol_star_triangle_down_open_dot ScattercarpetMarkerSymbol = "star-triangle-down-open-dot" // star-triangle-down-open-dot
    ScattercarpetMarkerSymbol21 ScattercarpetMarkerSymbol = "21" // 21
    ScattercarpetMarkerSymbol_star_square ScattercarpetMarkerSymbol = "star-square" // star-square
    ScattercarpetMarkerSymbol121 ScattercarpetMarkerSymbol = "121" // 121
    ScattercarpetMarkerSymbol_star_square_open ScattercarpetMarkerSymbol = "star-square-open" // star-square-open
    ScattercarpetMarkerSymbol221 ScattercarpetMarkerSymbol = "221" // 221
    ScattercarpetMarkerSymbol_star_square_dot ScattercarpetMarkerSymbol = "star-square-dot" // star-square-dot
    ScattercarpetMarkerSymbol321 ScattercarpetMarkerSymbol = "321" // 321
    ScattercarpetMarkerSymbol_star_square_open_dot ScattercarpetMarkerSymbol = "star-square-open-dot" // star-square-open-dot
    ScattercarpetMarkerSymbol22 ScattercarpetMarkerSymbol = "22" // 22
    ScattercarpetMarkerSymbol_star_diamond ScattercarpetMarkerSymbol = "star-diamond" // star-diamond
    ScattercarpetMarkerSymbol122 ScattercarpetMarkerSymbol = "122" // 122
    ScattercarpetMarkerSymbol_star_diamond_open ScattercarpetMarkerSymbol = "star-diamond-open" // star-diamond-open
    ScattercarpetMarkerSymbol222 ScattercarpetMarkerSymbol = "222" // 222
    ScattercarpetMarkerSymbol_star_diamond_dot ScattercarpetMarkerSymbol = "star-diamond-dot" // star-diamond-dot
    ScattercarpetMarkerSymbol322 ScattercarpetMarkerSymbol = "322" // 322
    ScattercarpetMarkerSymbol_star_diamond_open_dot ScattercarpetMarkerSymbol = "star-diamond-open-dot" // star-diamond-open-dot
    ScattercarpetMarkerSymbol23 ScattercarpetMarkerSymbol = "23" // 23
    ScattercarpetMarkerSymbol_diamond_tall ScattercarpetMarkerSymbol = "diamond-tall" // diamond-tall
    ScattercarpetMarkerSymbol123 ScattercarpetMarkerSymbol = "123" // 123
    ScattercarpetMarkerSymbol_diamond_tall_open ScattercarpetMarkerSymbol = "diamond-tall-open" // diamond-tall-open
    ScattercarpetMarkerSymbol223 ScattercarpetMarkerSymbol = "223" // 223
    ScattercarpetMarkerSymbol_diamond_tall_dot ScattercarpetMarkerSymbol = "diamond-tall-dot" // diamond-tall-dot
    ScattercarpetMarkerSymbol323 ScattercarpetMarkerSymbol = "323" // 323
    ScattercarpetMarkerSymbol_diamond_tall_open_dot ScattercarpetMarkerSymbol = "diamond-tall-open-dot" // diamond-tall-open-dot
    ScattercarpetMarkerSymbol24 ScattercarpetMarkerSymbol = "24" // 24
    ScattercarpetMarkerSymbol_diamond_wide ScattercarpetMarkerSymbol = "diamond-wide" // diamond-wide
    ScattercarpetMarkerSymbol124 ScattercarpetMarkerSymbol = "124" // 124
    ScattercarpetMarkerSymbol_diamond_wide_open ScattercarpetMarkerSymbol = "diamond-wide-open" // diamond-wide-open
    ScattercarpetMarkerSymbol224 ScattercarpetMarkerSymbol = "224" // 224
    ScattercarpetMarkerSymbol_diamond_wide_dot ScattercarpetMarkerSymbol = "diamond-wide-dot" // diamond-wide-dot
    ScattercarpetMarkerSymbol324 ScattercarpetMarkerSymbol = "324" // 324
    ScattercarpetMarkerSymbol_diamond_wide_open_dot ScattercarpetMarkerSymbol = "diamond-wide-open-dot" // diamond-wide-open-dot
    ScattercarpetMarkerSymbol25 ScattercarpetMarkerSymbol = "25" // 25
    ScattercarpetMarkerSymbol_hourglass ScattercarpetMarkerSymbol = "hourglass" // hourglass
    ScattercarpetMarkerSymbol125 ScattercarpetMarkerSymbol = "125" // 125
    ScattercarpetMarkerSymbol_hourglass_open ScattercarpetMarkerSymbol = "hourglass-open" // hourglass-open
    ScattercarpetMarkerSymbol26 ScattercarpetMarkerSymbol = "26" // 26
    ScattercarpetMarkerSymbol_bowtie ScattercarpetMarkerSymbol = "bowtie" // bowtie
    ScattercarpetMarkerSymbol126 ScattercarpetMarkerSymbol = "126" // 126
    ScattercarpetMarkerSymbol_bowtie_open ScattercarpetMarkerSymbol = "bowtie-open" // bowtie-open
    ScattercarpetMarkerSymbol27 ScattercarpetMarkerSymbol = "27" // 27
    ScattercarpetMarkerSymbol_circle_cross ScattercarpetMarkerSymbol = "circle-cross" // circle-cross
    ScattercarpetMarkerSymbol127 ScattercarpetMarkerSymbol = "127" // 127
    ScattercarpetMarkerSymbol_circle_cross_open ScattercarpetMarkerSymbol = "circle-cross-open" // circle-cross-open
    ScattercarpetMarkerSymbol28 ScattercarpetMarkerSymbol = "28" // 28
    ScattercarpetMarkerSymbol_circle_x ScattercarpetMarkerSymbol = "circle-x" // circle-x
    ScattercarpetMarkerSymbol128 ScattercarpetMarkerSymbol = "128" // 128
    ScattercarpetMarkerSymbol_circle_x_open ScattercarpetMarkerSymbol = "circle-x-open" // circle-x-open
    ScattercarpetMarkerSymbol29 ScattercarpetMarkerSymbol = "29" // 29
    ScattercarpetMarkerSymbol_square_cross ScattercarpetMarkerSymbol = "square-cross" // square-cross
    ScattercarpetMarkerSymbol129 ScattercarpetMarkerSymbol = "129" // 129
    ScattercarpetMarkerSymbol_square_cross_open ScattercarpetMarkerSymbol = "square-cross-open" // square-cross-open
    ScattercarpetMarkerSymbol30 ScattercarpetMarkerSymbol = "30" // 30
    ScattercarpetMarkerSymbol_square_x ScattercarpetMarkerSymbol = "square-x" // square-x
    ScattercarpetMarkerSymbol130 ScattercarpetMarkerSymbol = "130" // 130
    ScattercarpetMarkerSymbol_square_x_open ScattercarpetMarkerSymbol = "square-x-open" // square-x-open
    ScattercarpetMarkerSymbol31 ScattercarpetMarkerSymbol = "31" // 31
    ScattercarpetMarkerSymbol_diamond_cross ScattercarpetMarkerSymbol = "diamond-cross" // diamond-cross
    ScattercarpetMarkerSymbol131 ScattercarpetMarkerSymbol = "131" // 131
    ScattercarpetMarkerSymbol_diamond_cross_open ScattercarpetMarkerSymbol = "diamond-cross-open" // diamond-cross-open
    ScattercarpetMarkerSymbol32 ScattercarpetMarkerSymbol = "32" // 32
    ScattercarpetMarkerSymbol_diamond_x ScattercarpetMarkerSymbol = "diamond-x" // diamond-x
    ScattercarpetMarkerSymbol132 ScattercarpetMarkerSymbol = "132" // 132
    ScattercarpetMarkerSymbol_diamond_x_open ScattercarpetMarkerSymbol = "diamond-x-open" // diamond-x-open
    ScattercarpetMarkerSymbol33 ScattercarpetMarkerSymbol = "33" // 33
    ScattercarpetMarkerSymbol_cross_thin ScattercarpetMarkerSymbol = "cross-thin" // cross-thin
    ScattercarpetMarkerSymbol133 ScattercarpetMarkerSymbol = "133" // 133
    ScattercarpetMarkerSymbol_cross_thin_open ScattercarpetMarkerSymbol = "cross-thin-open" // cross-thin-open
    ScattercarpetMarkerSymbol34 ScattercarpetMarkerSymbol = "34" // 34
    ScattercarpetMarkerSymbol_x_thin ScattercarpetMarkerSymbol = "x-thin" // x-thin
    ScattercarpetMarkerSymbol134 ScattercarpetMarkerSymbol = "134" // 134
    ScattercarpetMarkerSymbol_x_thin_open ScattercarpetMarkerSymbol = "x-thin-open" // x-thin-open
    ScattercarpetMarkerSymbol35 ScattercarpetMarkerSymbol = "35" // 35
    ScattercarpetMarkerSymbol_asterisk ScattercarpetMarkerSymbol = "asterisk" // asterisk
    ScattercarpetMarkerSymbol135 ScattercarpetMarkerSymbol = "135" // 135
    ScattercarpetMarkerSymbol_asterisk_open ScattercarpetMarkerSymbol = "asterisk-open" // asterisk-open
    ScattercarpetMarkerSymbol36 ScattercarpetMarkerSymbol = "36" // 36
    ScattercarpetMarkerSymbol_hash ScattercarpetMarkerSymbol = "hash" // hash
    ScattercarpetMarkerSymbol136 ScattercarpetMarkerSymbol = "136" // 136
    ScattercarpetMarkerSymbol_hash_open ScattercarpetMarkerSymbol = "hash-open" // hash-open
    ScattercarpetMarkerSymbol236 ScattercarpetMarkerSymbol = "236" // 236
    ScattercarpetMarkerSymbol_hash_dot ScattercarpetMarkerSymbol = "hash-dot" // hash-dot
    ScattercarpetMarkerSymbol336 ScattercarpetMarkerSymbol = "336" // 336
    ScattercarpetMarkerSymbol_hash_open_dot ScattercarpetMarkerSymbol = "hash-open-dot" // hash-open-dot
    ScattercarpetMarkerSymbol37 ScattercarpetMarkerSymbol = "37" // 37
    ScattercarpetMarkerSymbol_y_up ScattercarpetMarkerSymbol = "y-up" // y-up
    ScattercarpetMarkerSymbol137 ScattercarpetMarkerSymbol = "137" // 137
    ScattercarpetMarkerSymbol_y_up_open ScattercarpetMarkerSymbol = "y-up-open" // y-up-open
    ScattercarpetMarkerSymbol38 ScattercarpetMarkerSymbol = "38" // 38
    ScattercarpetMarkerSymbol_y_down ScattercarpetMarkerSymbol = "y-down" // y-down
    ScattercarpetMarkerSymbol138 ScattercarpetMarkerSymbol = "138" // 138
    ScattercarpetMarkerSymbol_y_down_open ScattercarpetMarkerSymbol = "y-down-open" // y-down-open
    ScattercarpetMarkerSymbol39 ScattercarpetMarkerSymbol = "39" // 39
    ScattercarpetMarkerSymbol_y_left ScattercarpetMarkerSymbol = "y-left" // y-left
    ScattercarpetMarkerSymbol139 ScattercarpetMarkerSymbol = "139" // 139
    ScattercarpetMarkerSymbol_y_left_open ScattercarpetMarkerSymbol = "y-left-open" // y-left-open
    ScattercarpetMarkerSymbol40 ScattercarpetMarkerSymbol = "40" // 40
    ScattercarpetMarkerSymbol_y_right ScattercarpetMarkerSymbol = "y-right" // y-right
    ScattercarpetMarkerSymbol140 ScattercarpetMarkerSymbol = "140" // 140
    ScattercarpetMarkerSymbol_y_right_open ScattercarpetMarkerSymbol = "y-right-open" // y-right-open
    ScattercarpetMarkerSymbol41 ScattercarpetMarkerSymbol = "41" // 41
    ScattercarpetMarkerSymbol_line_ew ScattercarpetMarkerSymbol = "line-ew" // line-ew
    ScattercarpetMarkerSymbol141 ScattercarpetMarkerSymbol = "141" // 141
    ScattercarpetMarkerSymbol_line_ew_open ScattercarpetMarkerSymbol = "line-ew-open" // line-ew-open
    ScattercarpetMarkerSymbol42 ScattercarpetMarkerSymbol = "42" // 42
    ScattercarpetMarkerSymbol_line_ns ScattercarpetMarkerSymbol = "line-ns" // line-ns
    ScattercarpetMarkerSymbol142 ScattercarpetMarkerSymbol = "142" // 142
    ScattercarpetMarkerSymbol_line_ns_open ScattercarpetMarkerSymbol = "line-ns-open" // line-ns-open
    ScattercarpetMarkerSymbol43 ScattercarpetMarkerSymbol = "43" // 43
    ScattercarpetMarkerSymbol_line_ne ScattercarpetMarkerSymbol = "line-ne" // line-ne
    ScattercarpetMarkerSymbol143 ScattercarpetMarkerSymbol = "143" // 143
    ScattercarpetMarkerSymbol_line_ne_open ScattercarpetMarkerSymbol = "line-ne-open" // line-ne-open
    ScattercarpetMarkerSymbol44 ScattercarpetMarkerSymbol = "44" // 44
    ScattercarpetMarkerSymbol_line_nw ScattercarpetMarkerSymbol = "line-nw" // line-nw
    ScattercarpetMarkerSymbol144 ScattercarpetMarkerSymbol = "144" // 144
    ScattercarpetMarkerSymbol_line_nw_open ScattercarpetMarkerSymbol = "line-nw-open" // line-nw-open
)

// ScattercarpetTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScattercarpetTextposition string

const (
    ScattercarpetTextposition_topleft ScattercarpetTextposition = "top left" // top left
    ScattercarpetTextposition_topcenter ScattercarpetTextposition = "top center" // top center
    ScattercarpetTextposition_topright ScattercarpetTextposition = "top right" // top right
    ScattercarpetTextposition_middleleft ScattercarpetTextposition = "middle left" // middle left
    ScattercarpetTextposition_middlecenter ScattercarpetTextposition = "middle center" // middle center
    ScattercarpetTextposition_middleright ScattercarpetTextposition = "middle right" // middle right
    ScattercarpetTextposition_bottomleft ScattercarpetTextposition = "bottom left" // bottom left
    ScattercarpetTextposition_bottomcenter ScattercarpetTextposition = "bottom center" // bottom center
    ScattercarpetTextposition_bottomright ScattercarpetTextposition = "bottom right" // bottom right
)

// ScattercarpetVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScattercarpetVisible interface{}

const (
    ScattercarpetVisibleTrue bool = true
    ScattercarpetVisibleFalse bool = false
    ScattercarpetVisibleLegendonly string = "legendonly"
)

// ScattergeoFill Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape.
type ScattergeoFill string

const (
    ScattergeoFill_none ScattergeoFill = "none" // none
    ScattergeoFill_toself ScattergeoFill = "toself" // toself
)

// ScattergeoHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScattergeoHoverlabelAlign string

const (
    ScattergeoHoverlabelAlign_left ScattergeoHoverlabelAlign = "left" // left
    ScattergeoHoverlabelAlign_right ScattergeoHoverlabelAlign = "right" // right
    ScattergeoHoverlabelAlign_auto ScattergeoHoverlabelAlign = "auto" // auto
)

// ScattergeoLocationmode Determines the set of locations used to match entries in `locations` to regions on the map. Values *ISO-3*, *USA-states*, *country names* correspond to features on the base map and value *geojson-id* corresponds to features from a custom GeoJSON linked to the `geojson` attribute.
type ScattergeoLocationmode string

const (
    ScattergeoLocationmode_ISO_3 ScattergeoLocationmode = "ISO-3" // ISO-3
    ScattergeoLocationmode_USA_states ScattergeoLocationmode = "USA-states" // USA-states
    ScattergeoLocationmode_countrynames ScattergeoLocationmode = "country names" // country names
    ScattergeoLocationmode_geojson_id ScattergeoLocationmode = "geojson-id" // geojson-id
)

// ScattergeoMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScattergeoMarkerColorbarExponentformat string

const (
    ScattergeoMarkerColorbarExponentformat_none ScattergeoMarkerColorbarExponentformat = "none" // none
    ScattergeoMarkerColorbarExponentformat_e ScattergeoMarkerColorbarExponentformat = "e" // e
    ScattergeoMarkerColorbarExponentformat_E ScattergeoMarkerColorbarExponentformat = "E" // E
    ScattergeoMarkerColorbarExponentformat_power ScattergeoMarkerColorbarExponentformat = "power" // power
    ScattergeoMarkerColorbarExponentformat_SI ScattergeoMarkerColorbarExponentformat = "SI" // SI
    ScattergeoMarkerColorbarExponentformat_B ScattergeoMarkerColorbarExponentformat = "B" // B
)

// ScattergeoMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScattergeoMarkerColorbarLenmode string

const (
    ScattergeoMarkerColorbarLenmode_fraction ScattergeoMarkerColorbarLenmode = "fraction" // fraction
    ScattergeoMarkerColorbarLenmode_pixels ScattergeoMarkerColorbarLenmode = "pixels" // pixels
)

// ScattergeoMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScattergeoMarkerColorbarShowexponent string

const (
    ScattergeoMarkerColorbarShowexponent_all ScattergeoMarkerColorbarShowexponent = "all" // all
    ScattergeoMarkerColorbarShowexponent_first ScattergeoMarkerColorbarShowexponent = "first" // first
    ScattergeoMarkerColorbarShowexponent_last ScattergeoMarkerColorbarShowexponent = "last" // last
    ScattergeoMarkerColorbarShowexponent_none ScattergeoMarkerColorbarShowexponent = "none" // none
)

// ScattergeoMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScattergeoMarkerColorbarShowtickprefix string

const (
    ScattergeoMarkerColorbarShowtickprefix_all ScattergeoMarkerColorbarShowtickprefix = "all" // all
    ScattergeoMarkerColorbarShowtickprefix_first ScattergeoMarkerColorbarShowtickprefix = "first" // first
    ScattergeoMarkerColorbarShowtickprefix_last ScattergeoMarkerColorbarShowtickprefix = "last" // last
    ScattergeoMarkerColorbarShowtickprefix_none ScattergeoMarkerColorbarShowtickprefix = "none" // none
)

// ScattergeoMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScattergeoMarkerColorbarShowticksuffix string

const (
    ScattergeoMarkerColorbarShowticksuffix_all ScattergeoMarkerColorbarShowticksuffix = "all" // all
    ScattergeoMarkerColorbarShowticksuffix_first ScattergeoMarkerColorbarShowticksuffix = "first" // first
    ScattergeoMarkerColorbarShowticksuffix_last ScattergeoMarkerColorbarShowticksuffix = "last" // last
    ScattergeoMarkerColorbarShowticksuffix_none ScattergeoMarkerColorbarShowticksuffix = "none" // none
)

// ScattergeoMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScattergeoMarkerColorbarThicknessmode string

const (
    ScattergeoMarkerColorbarThicknessmode_fraction ScattergeoMarkerColorbarThicknessmode = "fraction" // fraction
    ScattergeoMarkerColorbarThicknessmode_pixels ScattergeoMarkerColorbarThicknessmode = "pixels" // pixels
)

// ScattergeoMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScattergeoMarkerColorbarTickmode string

const (
    ScattergeoMarkerColorbarTickmode_auto ScattergeoMarkerColorbarTickmode = "auto" // auto
    ScattergeoMarkerColorbarTickmode_linear ScattergeoMarkerColorbarTickmode = "linear" // linear
    ScattergeoMarkerColorbarTickmode_array ScattergeoMarkerColorbarTickmode = "array" // array
)

// ScattergeoMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScattergeoMarkerColorbarTicks string

const (
    ScattergeoMarkerColorbarTicks_outside ScattergeoMarkerColorbarTicks = "outside" // outside
    ScattergeoMarkerColorbarTicks_inside ScattergeoMarkerColorbarTicks = "inside" // inside
)

// ScattergeoMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScattergeoMarkerColorbarTitleSide string

const (
    ScattergeoMarkerColorbarTitleSide_right ScattergeoMarkerColorbarTitleSide = "right" // right
    ScattergeoMarkerColorbarTitleSide_top ScattergeoMarkerColorbarTitleSide = "top" // top
    ScattergeoMarkerColorbarTitleSide_bottom ScattergeoMarkerColorbarTitleSide = "bottom" // bottom
)

// ScattergeoMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScattergeoMarkerColorbarXanchor string

const (
    ScattergeoMarkerColorbarXanchor_left ScattergeoMarkerColorbarXanchor = "left" // left
    ScattergeoMarkerColorbarXanchor_center ScattergeoMarkerColorbarXanchor = "center" // center
    ScattergeoMarkerColorbarXanchor_right ScattergeoMarkerColorbarXanchor = "right" // right
)

// ScattergeoMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScattergeoMarkerColorbarYanchor string

const (
    ScattergeoMarkerColorbarYanchor_top ScattergeoMarkerColorbarYanchor = "top" // top
    ScattergeoMarkerColorbarYanchor_middle ScattergeoMarkerColorbarYanchor = "middle" // middle
    ScattergeoMarkerColorbarYanchor_bottom ScattergeoMarkerColorbarYanchor = "bottom" // bottom
)

// ScattergeoMarkerGradientType Sets the type of gradient used to fill the markers
type ScattergeoMarkerGradientType string

const (
    ScattergeoMarkerGradientType_radial ScattergeoMarkerGradientType = "radial" // radial
    ScattergeoMarkerGradientType_horizontal ScattergeoMarkerGradientType = "horizontal" // horizontal
    ScattergeoMarkerGradientType_vertical ScattergeoMarkerGradientType = "vertical" // vertical
    ScattergeoMarkerGradientType_none ScattergeoMarkerGradientType = "none" // none
)

// ScattergeoMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScattergeoMarkerSizemode string

const (
    ScattergeoMarkerSizemode_diameter ScattergeoMarkerSizemode = "diameter" // diameter
    ScattergeoMarkerSizemode_area ScattergeoMarkerSizemode = "area" // area
)

// ScattergeoMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScattergeoMarkerSymbol string

const (
    ScattergeoMarkerSymbol0 ScattergeoMarkerSymbol = "0" // 0
    ScattergeoMarkerSymbol_circle ScattergeoMarkerSymbol = "circle" // circle
    ScattergeoMarkerSymbol100 ScattergeoMarkerSymbol = "100" // 100
    ScattergeoMarkerSymbol_circle_open ScattergeoMarkerSymbol = "circle-open" // circle-open
    ScattergeoMarkerSymbol200 ScattergeoMarkerSymbol = "200" // 200
    ScattergeoMarkerSymbol_circle_dot ScattergeoMarkerSymbol = "circle-dot" // circle-dot
    ScattergeoMarkerSymbol300 ScattergeoMarkerSymbol = "300" // 300
    ScattergeoMarkerSymbol_circle_open_dot ScattergeoMarkerSymbol = "circle-open-dot" // circle-open-dot
    ScattergeoMarkerSymbol1 ScattergeoMarkerSymbol = "1" // 1
    ScattergeoMarkerSymbol_square ScattergeoMarkerSymbol = "square" // square
    ScattergeoMarkerSymbol101 ScattergeoMarkerSymbol = "101" // 101
    ScattergeoMarkerSymbol_square_open ScattergeoMarkerSymbol = "square-open" // square-open
    ScattergeoMarkerSymbol201 ScattergeoMarkerSymbol = "201" // 201
    ScattergeoMarkerSymbol_square_dot ScattergeoMarkerSymbol = "square-dot" // square-dot
    ScattergeoMarkerSymbol301 ScattergeoMarkerSymbol = "301" // 301
    ScattergeoMarkerSymbol_square_open_dot ScattergeoMarkerSymbol = "square-open-dot" // square-open-dot
    ScattergeoMarkerSymbol2 ScattergeoMarkerSymbol = "2" // 2
    ScattergeoMarkerSymbol_diamond ScattergeoMarkerSymbol = "diamond" // diamond
    ScattergeoMarkerSymbol102 ScattergeoMarkerSymbol = "102" // 102
    ScattergeoMarkerSymbol_diamond_open ScattergeoMarkerSymbol = "diamond-open" // diamond-open
    ScattergeoMarkerSymbol202 ScattergeoMarkerSymbol = "202" // 202
    ScattergeoMarkerSymbol_diamond_dot ScattergeoMarkerSymbol = "diamond-dot" // diamond-dot
    ScattergeoMarkerSymbol302 ScattergeoMarkerSymbol = "302" // 302
    ScattergeoMarkerSymbol_diamond_open_dot ScattergeoMarkerSymbol = "diamond-open-dot" // diamond-open-dot
    ScattergeoMarkerSymbol3 ScattergeoMarkerSymbol = "3" // 3
    ScattergeoMarkerSymbol_cross ScattergeoMarkerSymbol = "cross" // cross
    ScattergeoMarkerSymbol103 ScattergeoMarkerSymbol = "103" // 103
    ScattergeoMarkerSymbol_cross_open ScattergeoMarkerSymbol = "cross-open" // cross-open
    ScattergeoMarkerSymbol203 ScattergeoMarkerSymbol = "203" // 203
    ScattergeoMarkerSymbol_cross_dot ScattergeoMarkerSymbol = "cross-dot" // cross-dot
    ScattergeoMarkerSymbol303 ScattergeoMarkerSymbol = "303" // 303
    ScattergeoMarkerSymbol_cross_open_dot ScattergeoMarkerSymbol = "cross-open-dot" // cross-open-dot
    ScattergeoMarkerSymbol4 ScattergeoMarkerSymbol = "4" // 4
    ScattergeoMarkerSymbol_x ScattergeoMarkerSymbol = "x" // x
    ScattergeoMarkerSymbol104 ScattergeoMarkerSymbol = "104" // 104
    ScattergeoMarkerSymbol_x_open ScattergeoMarkerSymbol = "x-open" // x-open
    ScattergeoMarkerSymbol204 ScattergeoMarkerSymbol = "204" // 204
    ScattergeoMarkerSymbol_x_dot ScattergeoMarkerSymbol = "x-dot" // x-dot
    ScattergeoMarkerSymbol304 ScattergeoMarkerSymbol = "304" // 304
    ScattergeoMarkerSymbol_x_open_dot ScattergeoMarkerSymbol = "x-open-dot" // x-open-dot
    ScattergeoMarkerSymbol5 ScattergeoMarkerSymbol = "5" // 5
    ScattergeoMarkerSymbol_triangle_up ScattergeoMarkerSymbol = "triangle-up" // triangle-up
    ScattergeoMarkerSymbol105 ScattergeoMarkerSymbol = "105" // 105
    ScattergeoMarkerSymbol_triangle_up_open ScattergeoMarkerSymbol = "triangle-up-open" // triangle-up-open
    ScattergeoMarkerSymbol205 ScattergeoMarkerSymbol = "205" // 205
    ScattergeoMarkerSymbol_triangle_up_dot ScattergeoMarkerSymbol = "triangle-up-dot" // triangle-up-dot
    ScattergeoMarkerSymbol305 ScattergeoMarkerSymbol = "305" // 305
    ScattergeoMarkerSymbol_triangle_up_open_dot ScattergeoMarkerSymbol = "triangle-up-open-dot" // triangle-up-open-dot
    ScattergeoMarkerSymbol6 ScattergeoMarkerSymbol = "6" // 6
    ScattergeoMarkerSymbol_triangle_down ScattergeoMarkerSymbol = "triangle-down" // triangle-down
    ScattergeoMarkerSymbol106 ScattergeoMarkerSymbol = "106" // 106
    ScattergeoMarkerSymbol_triangle_down_open ScattergeoMarkerSymbol = "triangle-down-open" // triangle-down-open
    ScattergeoMarkerSymbol206 ScattergeoMarkerSymbol = "206" // 206
    ScattergeoMarkerSymbol_triangle_down_dot ScattergeoMarkerSymbol = "triangle-down-dot" // triangle-down-dot
    ScattergeoMarkerSymbol306 ScattergeoMarkerSymbol = "306" // 306
    ScattergeoMarkerSymbol_triangle_down_open_dot ScattergeoMarkerSymbol = "triangle-down-open-dot" // triangle-down-open-dot
    ScattergeoMarkerSymbol7 ScattergeoMarkerSymbol = "7" // 7
    ScattergeoMarkerSymbol_triangle_left ScattergeoMarkerSymbol = "triangle-left" // triangle-left
    ScattergeoMarkerSymbol107 ScattergeoMarkerSymbol = "107" // 107
    ScattergeoMarkerSymbol_triangle_left_open ScattergeoMarkerSymbol = "triangle-left-open" // triangle-left-open
    ScattergeoMarkerSymbol207 ScattergeoMarkerSymbol = "207" // 207
    ScattergeoMarkerSymbol_triangle_left_dot ScattergeoMarkerSymbol = "triangle-left-dot" // triangle-left-dot
    ScattergeoMarkerSymbol307 ScattergeoMarkerSymbol = "307" // 307
    ScattergeoMarkerSymbol_triangle_left_open_dot ScattergeoMarkerSymbol = "triangle-left-open-dot" // triangle-left-open-dot
    ScattergeoMarkerSymbol8 ScattergeoMarkerSymbol = "8" // 8
    ScattergeoMarkerSymbol_triangle_right ScattergeoMarkerSymbol = "triangle-right" // triangle-right
    ScattergeoMarkerSymbol108 ScattergeoMarkerSymbol = "108" // 108
    ScattergeoMarkerSymbol_triangle_right_open ScattergeoMarkerSymbol = "triangle-right-open" // triangle-right-open
    ScattergeoMarkerSymbol208 ScattergeoMarkerSymbol = "208" // 208
    ScattergeoMarkerSymbol_triangle_right_dot ScattergeoMarkerSymbol = "triangle-right-dot" // triangle-right-dot
    ScattergeoMarkerSymbol308 ScattergeoMarkerSymbol = "308" // 308
    ScattergeoMarkerSymbol_triangle_right_open_dot ScattergeoMarkerSymbol = "triangle-right-open-dot" // triangle-right-open-dot
    ScattergeoMarkerSymbol9 ScattergeoMarkerSymbol = "9" // 9
    ScattergeoMarkerSymbol_triangle_ne ScattergeoMarkerSymbol = "triangle-ne" // triangle-ne
    ScattergeoMarkerSymbol109 ScattergeoMarkerSymbol = "109" // 109
    ScattergeoMarkerSymbol_triangle_ne_open ScattergeoMarkerSymbol = "triangle-ne-open" // triangle-ne-open
    ScattergeoMarkerSymbol209 ScattergeoMarkerSymbol = "209" // 209
    ScattergeoMarkerSymbol_triangle_ne_dot ScattergeoMarkerSymbol = "triangle-ne-dot" // triangle-ne-dot
    ScattergeoMarkerSymbol309 ScattergeoMarkerSymbol = "309" // 309
    ScattergeoMarkerSymbol_triangle_ne_open_dot ScattergeoMarkerSymbol = "triangle-ne-open-dot" // triangle-ne-open-dot
    ScattergeoMarkerSymbol10 ScattergeoMarkerSymbol = "10" // 10
    ScattergeoMarkerSymbol_triangle_se ScattergeoMarkerSymbol = "triangle-se" // triangle-se
    ScattergeoMarkerSymbol110 ScattergeoMarkerSymbol = "110" // 110
    ScattergeoMarkerSymbol_triangle_se_open ScattergeoMarkerSymbol = "triangle-se-open" // triangle-se-open
    ScattergeoMarkerSymbol210 ScattergeoMarkerSymbol = "210" // 210
    ScattergeoMarkerSymbol_triangle_se_dot ScattergeoMarkerSymbol = "triangle-se-dot" // triangle-se-dot
    ScattergeoMarkerSymbol310 ScattergeoMarkerSymbol = "310" // 310
    ScattergeoMarkerSymbol_triangle_se_open_dot ScattergeoMarkerSymbol = "triangle-se-open-dot" // triangle-se-open-dot
    ScattergeoMarkerSymbol11 ScattergeoMarkerSymbol = "11" // 11
    ScattergeoMarkerSymbol_triangle_sw ScattergeoMarkerSymbol = "triangle-sw" // triangle-sw
    ScattergeoMarkerSymbol111 ScattergeoMarkerSymbol = "111" // 111
    ScattergeoMarkerSymbol_triangle_sw_open ScattergeoMarkerSymbol = "triangle-sw-open" // triangle-sw-open
    ScattergeoMarkerSymbol211 ScattergeoMarkerSymbol = "211" // 211
    ScattergeoMarkerSymbol_triangle_sw_dot ScattergeoMarkerSymbol = "triangle-sw-dot" // triangle-sw-dot
    ScattergeoMarkerSymbol311 ScattergeoMarkerSymbol = "311" // 311
    ScattergeoMarkerSymbol_triangle_sw_open_dot ScattergeoMarkerSymbol = "triangle-sw-open-dot" // triangle-sw-open-dot
    ScattergeoMarkerSymbol12 ScattergeoMarkerSymbol = "12" // 12
    ScattergeoMarkerSymbol_triangle_nw ScattergeoMarkerSymbol = "triangle-nw" // triangle-nw
    ScattergeoMarkerSymbol112 ScattergeoMarkerSymbol = "112" // 112
    ScattergeoMarkerSymbol_triangle_nw_open ScattergeoMarkerSymbol = "triangle-nw-open" // triangle-nw-open
    ScattergeoMarkerSymbol212 ScattergeoMarkerSymbol = "212" // 212
    ScattergeoMarkerSymbol_triangle_nw_dot ScattergeoMarkerSymbol = "triangle-nw-dot" // triangle-nw-dot
    ScattergeoMarkerSymbol312 ScattergeoMarkerSymbol = "312" // 312
    ScattergeoMarkerSymbol_triangle_nw_open_dot ScattergeoMarkerSymbol = "triangle-nw-open-dot" // triangle-nw-open-dot
    ScattergeoMarkerSymbol13 ScattergeoMarkerSymbol = "13" // 13
    ScattergeoMarkerSymbol_pentagon ScattergeoMarkerSymbol = "pentagon" // pentagon
    ScattergeoMarkerSymbol113 ScattergeoMarkerSymbol = "113" // 113
    ScattergeoMarkerSymbol_pentagon_open ScattergeoMarkerSymbol = "pentagon-open" // pentagon-open
    ScattergeoMarkerSymbol213 ScattergeoMarkerSymbol = "213" // 213
    ScattergeoMarkerSymbol_pentagon_dot ScattergeoMarkerSymbol = "pentagon-dot" // pentagon-dot
    ScattergeoMarkerSymbol313 ScattergeoMarkerSymbol = "313" // 313
    ScattergeoMarkerSymbol_pentagon_open_dot ScattergeoMarkerSymbol = "pentagon-open-dot" // pentagon-open-dot
    ScattergeoMarkerSymbol14 ScattergeoMarkerSymbol = "14" // 14
    ScattergeoMarkerSymbol_hexagon ScattergeoMarkerSymbol = "hexagon" // hexagon
    ScattergeoMarkerSymbol114 ScattergeoMarkerSymbol = "114" // 114
    ScattergeoMarkerSymbol_hexagon_open ScattergeoMarkerSymbol = "hexagon-open" // hexagon-open
    ScattergeoMarkerSymbol214 ScattergeoMarkerSymbol = "214" // 214
    ScattergeoMarkerSymbol_hexagon_dot ScattergeoMarkerSymbol = "hexagon-dot" // hexagon-dot
    ScattergeoMarkerSymbol314 ScattergeoMarkerSymbol = "314" // 314
    ScattergeoMarkerSymbol_hexagon_open_dot ScattergeoMarkerSymbol = "hexagon-open-dot" // hexagon-open-dot
    ScattergeoMarkerSymbol15 ScattergeoMarkerSymbol = "15" // 15
    ScattergeoMarkerSymbol_hexagon2 ScattergeoMarkerSymbol = "hexagon2" // hexagon2
    ScattergeoMarkerSymbol115 ScattergeoMarkerSymbol = "115" // 115
    ScattergeoMarkerSymbol_hexagon2_open ScattergeoMarkerSymbol = "hexagon2-open" // hexagon2-open
    ScattergeoMarkerSymbol215 ScattergeoMarkerSymbol = "215" // 215
    ScattergeoMarkerSymbol_hexagon2_dot ScattergeoMarkerSymbol = "hexagon2-dot" // hexagon2-dot
    ScattergeoMarkerSymbol315 ScattergeoMarkerSymbol = "315" // 315
    ScattergeoMarkerSymbol_hexagon2_open_dot ScattergeoMarkerSymbol = "hexagon2-open-dot" // hexagon2-open-dot
    ScattergeoMarkerSymbol16 ScattergeoMarkerSymbol = "16" // 16
    ScattergeoMarkerSymbol_octagon ScattergeoMarkerSymbol = "octagon" // octagon
    ScattergeoMarkerSymbol116 ScattergeoMarkerSymbol = "116" // 116
    ScattergeoMarkerSymbol_octagon_open ScattergeoMarkerSymbol = "octagon-open" // octagon-open
    ScattergeoMarkerSymbol216 ScattergeoMarkerSymbol = "216" // 216
    ScattergeoMarkerSymbol_octagon_dot ScattergeoMarkerSymbol = "octagon-dot" // octagon-dot
    ScattergeoMarkerSymbol316 ScattergeoMarkerSymbol = "316" // 316
    ScattergeoMarkerSymbol_octagon_open_dot ScattergeoMarkerSymbol = "octagon-open-dot" // octagon-open-dot
    ScattergeoMarkerSymbol17 ScattergeoMarkerSymbol = "17" // 17
    ScattergeoMarkerSymbol_star ScattergeoMarkerSymbol = "star" // star
    ScattergeoMarkerSymbol117 ScattergeoMarkerSymbol = "117" // 117
    ScattergeoMarkerSymbol_star_open ScattergeoMarkerSymbol = "star-open" // star-open
    ScattergeoMarkerSymbol217 ScattergeoMarkerSymbol = "217" // 217
    ScattergeoMarkerSymbol_star_dot ScattergeoMarkerSymbol = "star-dot" // star-dot
    ScattergeoMarkerSymbol317 ScattergeoMarkerSymbol = "317" // 317
    ScattergeoMarkerSymbol_star_open_dot ScattergeoMarkerSymbol = "star-open-dot" // star-open-dot
    ScattergeoMarkerSymbol18 ScattergeoMarkerSymbol = "18" // 18
    ScattergeoMarkerSymbol_hexagram ScattergeoMarkerSymbol = "hexagram" // hexagram
    ScattergeoMarkerSymbol118 ScattergeoMarkerSymbol = "118" // 118
    ScattergeoMarkerSymbol_hexagram_open ScattergeoMarkerSymbol = "hexagram-open" // hexagram-open
    ScattergeoMarkerSymbol218 ScattergeoMarkerSymbol = "218" // 218
    ScattergeoMarkerSymbol_hexagram_dot ScattergeoMarkerSymbol = "hexagram-dot" // hexagram-dot
    ScattergeoMarkerSymbol318 ScattergeoMarkerSymbol = "318" // 318
    ScattergeoMarkerSymbol_hexagram_open_dot ScattergeoMarkerSymbol = "hexagram-open-dot" // hexagram-open-dot
    ScattergeoMarkerSymbol19 ScattergeoMarkerSymbol = "19" // 19
    ScattergeoMarkerSymbol_star_triangle_up ScattergeoMarkerSymbol = "star-triangle-up" // star-triangle-up
    ScattergeoMarkerSymbol119 ScattergeoMarkerSymbol = "119" // 119
    ScattergeoMarkerSymbol_star_triangle_up_open ScattergeoMarkerSymbol = "star-triangle-up-open" // star-triangle-up-open
    ScattergeoMarkerSymbol219 ScattergeoMarkerSymbol = "219" // 219
    ScattergeoMarkerSymbol_star_triangle_up_dot ScattergeoMarkerSymbol = "star-triangle-up-dot" // star-triangle-up-dot
    ScattergeoMarkerSymbol319 ScattergeoMarkerSymbol = "319" // 319
    ScattergeoMarkerSymbol_star_triangle_up_open_dot ScattergeoMarkerSymbol = "star-triangle-up-open-dot" // star-triangle-up-open-dot
    ScattergeoMarkerSymbol20 ScattergeoMarkerSymbol = "20" // 20
    ScattergeoMarkerSymbol_star_triangle_down ScattergeoMarkerSymbol = "star-triangle-down" // star-triangle-down
    ScattergeoMarkerSymbol120 ScattergeoMarkerSymbol = "120" // 120
    ScattergeoMarkerSymbol_star_triangle_down_open ScattergeoMarkerSymbol = "star-triangle-down-open" // star-triangle-down-open
    ScattergeoMarkerSymbol220 ScattergeoMarkerSymbol = "220" // 220
    ScattergeoMarkerSymbol_star_triangle_down_dot ScattergeoMarkerSymbol = "star-triangle-down-dot" // star-triangle-down-dot
    ScattergeoMarkerSymbol320 ScattergeoMarkerSymbol = "320" // 320
    ScattergeoMarkerSymbol_star_triangle_down_open_dot ScattergeoMarkerSymbol = "star-triangle-down-open-dot" // star-triangle-down-open-dot
    ScattergeoMarkerSymbol21 ScattergeoMarkerSymbol = "21" // 21
    ScattergeoMarkerSymbol_star_square ScattergeoMarkerSymbol = "star-square" // star-square
    ScattergeoMarkerSymbol121 ScattergeoMarkerSymbol = "121" // 121
    ScattergeoMarkerSymbol_star_square_open ScattergeoMarkerSymbol = "star-square-open" // star-square-open
    ScattergeoMarkerSymbol221 ScattergeoMarkerSymbol = "221" // 221
    ScattergeoMarkerSymbol_star_square_dot ScattergeoMarkerSymbol = "star-square-dot" // star-square-dot
    ScattergeoMarkerSymbol321 ScattergeoMarkerSymbol = "321" // 321
    ScattergeoMarkerSymbol_star_square_open_dot ScattergeoMarkerSymbol = "star-square-open-dot" // star-square-open-dot
    ScattergeoMarkerSymbol22 ScattergeoMarkerSymbol = "22" // 22
    ScattergeoMarkerSymbol_star_diamond ScattergeoMarkerSymbol = "star-diamond" // star-diamond
    ScattergeoMarkerSymbol122 ScattergeoMarkerSymbol = "122" // 122
    ScattergeoMarkerSymbol_star_diamond_open ScattergeoMarkerSymbol = "star-diamond-open" // star-diamond-open
    ScattergeoMarkerSymbol222 ScattergeoMarkerSymbol = "222" // 222
    ScattergeoMarkerSymbol_star_diamond_dot ScattergeoMarkerSymbol = "star-diamond-dot" // star-diamond-dot
    ScattergeoMarkerSymbol322 ScattergeoMarkerSymbol = "322" // 322
    ScattergeoMarkerSymbol_star_diamond_open_dot ScattergeoMarkerSymbol = "star-diamond-open-dot" // star-diamond-open-dot
    ScattergeoMarkerSymbol23 ScattergeoMarkerSymbol = "23" // 23
    ScattergeoMarkerSymbol_diamond_tall ScattergeoMarkerSymbol = "diamond-tall" // diamond-tall
    ScattergeoMarkerSymbol123 ScattergeoMarkerSymbol = "123" // 123
    ScattergeoMarkerSymbol_diamond_tall_open ScattergeoMarkerSymbol = "diamond-tall-open" // diamond-tall-open
    ScattergeoMarkerSymbol223 ScattergeoMarkerSymbol = "223" // 223
    ScattergeoMarkerSymbol_diamond_tall_dot ScattergeoMarkerSymbol = "diamond-tall-dot" // diamond-tall-dot
    ScattergeoMarkerSymbol323 ScattergeoMarkerSymbol = "323" // 323
    ScattergeoMarkerSymbol_diamond_tall_open_dot ScattergeoMarkerSymbol = "diamond-tall-open-dot" // diamond-tall-open-dot
    ScattergeoMarkerSymbol24 ScattergeoMarkerSymbol = "24" // 24
    ScattergeoMarkerSymbol_diamond_wide ScattergeoMarkerSymbol = "diamond-wide" // diamond-wide
    ScattergeoMarkerSymbol124 ScattergeoMarkerSymbol = "124" // 124
    ScattergeoMarkerSymbol_diamond_wide_open ScattergeoMarkerSymbol = "diamond-wide-open" // diamond-wide-open
    ScattergeoMarkerSymbol224 ScattergeoMarkerSymbol = "224" // 224
    ScattergeoMarkerSymbol_diamond_wide_dot ScattergeoMarkerSymbol = "diamond-wide-dot" // diamond-wide-dot
    ScattergeoMarkerSymbol324 ScattergeoMarkerSymbol = "324" // 324
    ScattergeoMarkerSymbol_diamond_wide_open_dot ScattergeoMarkerSymbol = "diamond-wide-open-dot" // diamond-wide-open-dot
    ScattergeoMarkerSymbol25 ScattergeoMarkerSymbol = "25" // 25
    ScattergeoMarkerSymbol_hourglass ScattergeoMarkerSymbol = "hourglass" // hourglass
    ScattergeoMarkerSymbol125 ScattergeoMarkerSymbol = "125" // 125
    ScattergeoMarkerSymbol_hourglass_open ScattergeoMarkerSymbol = "hourglass-open" // hourglass-open
    ScattergeoMarkerSymbol26 ScattergeoMarkerSymbol = "26" // 26
    ScattergeoMarkerSymbol_bowtie ScattergeoMarkerSymbol = "bowtie" // bowtie
    ScattergeoMarkerSymbol126 ScattergeoMarkerSymbol = "126" // 126
    ScattergeoMarkerSymbol_bowtie_open ScattergeoMarkerSymbol = "bowtie-open" // bowtie-open
    ScattergeoMarkerSymbol27 ScattergeoMarkerSymbol = "27" // 27
    ScattergeoMarkerSymbol_circle_cross ScattergeoMarkerSymbol = "circle-cross" // circle-cross
    ScattergeoMarkerSymbol127 ScattergeoMarkerSymbol = "127" // 127
    ScattergeoMarkerSymbol_circle_cross_open ScattergeoMarkerSymbol = "circle-cross-open" // circle-cross-open
    ScattergeoMarkerSymbol28 ScattergeoMarkerSymbol = "28" // 28
    ScattergeoMarkerSymbol_circle_x ScattergeoMarkerSymbol = "circle-x" // circle-x
    ScattergeoMarkerSymbol128 ScattergeoMarkerSymbol = "128" // 128
    ScattergeoMarkerSymbol_circle_x_open ScattergeoMarkerSymbol = "circle-x-open" // circle-x-open
    ScattergeoMarkerSymbol29 ScattergeoMarkerSymbol = "29" // 29
    ScattergeoMarkerSymbol_square_cross ScattergeoMarkerSymbol = "square-cross" // square-cross
    ScattergeoMarkerSymbol129 ScattergeoMarkerSymbol = "129" // 129
    ScattergeoMarkerSymbol_square_cross_open ScattergeoMarkerSymbol = "square-cross-open" // square-cross-open
    ScattergeoMarkerSymbol30 ScattergeoMarkerSymbol = "30" // 30
    ScattergeoMarkerSymbol_square_x ScattergeoMarkerSymbol = "square-x" // square-x
    ScattergeoMarkerSymbol130 ScattergeoMarkerSymbol = "130" // 130
    ScattergeoMarkerSymbol_square_x_open ScattergeoMarkerSymbol = "square-x-open" // square-x-open
    ScattergeoMarkerSymbol31 ScattergeoMarkerSymbol = "31" // 31
    ScattergeoMarkerSymbol_diamond_cross ScattergeoMarkerSymbol = "diamond-cross" // diamond-cross
    ScattergeoMarkerSymbol131 ScattergeoMarkerSymbol = "131" // 131
    ScattergeoMarkerSymbol_diamond_cross_open ScattergeoMarkerSymbol = "diamond-cross-open" // diamond-cross-open
    ScattergeoMarkerSymbol32 ScattergeoMarkerSymbol = "32" // 32
    ScattergeoMarkerSymbol_diamond_x ScattergeoMarkerSymbol = "diamond-x" // diamond-x
    ScattergeoMarkerSymbol132 ScattergeoMarkerSymbol = "132" // 132
    ScattergeoMarkerSymbol_diamond_x_open ScattergeoMarkerSymbol = "diamond-x-open" // diamond-x-open
    ScattergeoMarkerSymbol33 ScattergeoMarkerSymbol = "33" // 33
    ScattergeoMarkerSymbol_cross_thin ScattergeoMarkerSymbol = "cross-thin" // cross-thin
    ScattergeoMarkerSymbol133 ScattergeoMarkerSymbol = "133" // 133
    ScattergeoMarkerSymbol_cross_thin_open ScattergeoMarkerSymbol = "cross-thin-open" // cross-thin-open
    ScattergeoMarkerSymbol34 ScattergeoMarkerSymbol = "34" // 34
    ScattergeoMarkerSymbol_x_thin ScattergeoMarkerSymbol = "x-thin" // x-thin
    ScattergeoMarkerSymbol134 ScattergeoMarkerSymbol = "134" // 134
    ScattergeoMarkerSymbol_x_thin_open ScattergeoMarkerSymbol = "x-thin-open" // x-thin-open
    ScattergeoMarkerSymbol35 ScattergeoMarkerSymbol = "35" // 35
    ScattergeoMarkerSymbol_asterisk ScattergeoMarkerSymbol = "asterisk" // asterisk
    ScattergeoMarkerSymbol135 ScattergeoMarkerSymbol = "135" // 135
    ScattergeoMarkerSymbol_asterisk_open ScattergeoMarkerSymbol = "asterisk-open" // asterisk-open
    ScattergeoMarkerSymbol36 ScattergeoMarkerSymbol = "36" // 36
    ScattergeoMarkerSymbol_hash ScattergeoMarkerSymbol = "hash" // hash
    ScattergeoMarkerSymbol136 ScattergeoMarkerSymbol = "136" // 136
    ScattergeoMarkerSymbol_hash_open ScattergeoMarkerSymbol = "hash-open" // hash-open
    ScattergeoMarkerSymbol236 ScattergeoMarkerSymbol = "236" // 236
    ScattergeoMarkerSymbol_hash_dot ScattergeoMarkerSymbol = "hash-dot" // hash-dot
    ScattergeoMarkerSymbol336 ScattergeoMarkerSymbol = "336" // 336
    ScattergeoMarkerSymbol_hash_open_dot ScattergeoMarkerSymbol = "hash-open-dot" // hash-open-dot
    ScattergeoMarkerSymbol37 ScattergeoMarkerSymbol = "37" // 37
    ScattergeoMarkerSymbol_y_up ScattergeoMarkerSymbol = "y-up" // y-up
    ScattergeoMarkerSymbol137 ScattergeoMarkerSymbol = "137" // 137
    ScattergeoMarkerSymbol_y_up_open ScattergeoMarkerSymbol = "y-up-open" // y-up-open
    ScattergeoMarkerSymbol38 ScattergeoMarkerSymbol = "38" // 38
    ScattergeoMarkerSymbol_y_down ScattergeoMarkerSymbol = "y-down" // y-down
    ScattergeoMarkerSymbol138 ScattergeoMarkerSymbol = "138" // 138
    ScattergeoMarkerSymbol_y_down_open ScattergeoMarkerSymbol = "y-down-open" // y-down-open
    ScattergeoMarkerSymbol39 ScattergeoMarkerSymbol = "39" // 39
    ScattergeoMarkerSymbol_y_left ScattergeoMarkerSymbol = "y-left" // y-left
    ScattergeoMarkerSymbol139 ScattergeoMarkerSymbol = "139" // 139
    ScattergeoMarkerSymbol_y_left_open ScattergeoMarkerSymbol = "y-left-open" // y-left-open
    ScattergeoMarkerSymbol40 ScattergeoMarkerSymbol = "40" // 40
    ScattergeoMarkerSymbol_y_right ScattergeoMarkerSymbol = "y-right" // y-right
    ScattergeoMarkerSymbol140 ScattergeoMarkerSymbol = "140" // 140
    ScattergeoMarkerSymbol_y_right_open ScattergeoMarkerSymbol = "y-right-open" // y-right-open
    ScattergeoMarkerSymbol41 ScattergeoMarkerSymbol = "41" // 41
    ScattergeoMarkerSymbol_line_ew ScattergeoMarkerSymbol = "line-ew" // line-ew
    ScattergeoMarkerSymbol141 ScattergeoMarkerSymbol = "141" // 141
    ScattergeoMarkerSymbol_line_ew_open ScattergeoMarkerSymbol = "line-ew-open" // line-ew-open
    ScattergeoMarkerSymbol42 ScattergeoMarkerSymbol = "42" // 42
    ScattergeoMarkerSymbol_line_ns ScattergeoMarkerSymbol = "line-ns" // line-ns
    ScattergeoMarkerSymbol142 ScattergeoMarkerSymbol = "142" // 142
    ScattergeoMarkerSymbol_line_ns_open ScattergeoMarkerSymbol = "line-ns-open" // line-ns-open
    ScattergeoMarkerSymbol43 ScattergeoMarkerSymbol = "43" // 43
    ScattergeoMarkerSymbol_line_ne ScattergeoMarkerSymbol = "line-ne" // line-ne
    ScattergeoMarkerSymbol143 ScattergeoMarkerSymbol = "143" // 143
    ScattergeoMarkerSymbol_line_ne_open ScattergeoMarkerSymbol = "line-ne-open" // line-ne-open
    ScattergeoMarkerSymbol44 ScattergeoMarkerSymbol = "44" // 44
    ScattergeoMarkerSymbol_line_nw ScattergeoMarkerSymbol = "line-nw" // line-nw
    ScattergeoMarkerSymbol144 ScattergeoMarkerSymbol = "144" // 144
    ScattergeoMarkerSymbol_line_nw_open ScattergeoMarkerSymbol = "line-nw-open" // line-nw-open
)

// ScattergeoTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScattergeoTextposition string

const (
    ScattergeoTextposition_topleft ScattergeoTextposition = "top left" // top left
    ScattergeoTextposition_topcenter ScattergeoTextposition = "top center" // top center
    ScattergeoTextposition_topright ScattergeoTextposition = "top right" // top right
    ScattergeoTextposition_middleleft ScattergeoTextposition = "middle left" // middle left
    ScattergeoTextposition_middlecenter ScattergeoTextposition = "middle center" // middle center
    ScattergeoTextposition_middleright ScattergeoTextposition = "middle right" // middle right
    ScattergeoTextposition_bottomleft ScattergeoTextposition = "bottom left" // bottom left
    ScattergeoTextposition_bottomcenter ScattergeoTextposition = "bottom center" // bottom center
    ScattergeoTextposition_bottomright ScattergeoTextposition = "bottom right" // bottom right
)

// ScattergeoVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScattergeoVisible interface{}

const (
    ScattergeoVisibleTrue bool = true
    ScattergeoVisibleFalse bool = false
    ScattergeoVisibleLegendonly string = "legendonly"
)

// ScatterglErrorXType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type ScatterglErrorXType string

const (
    ScatterglErrorXType_percent ScatterglErrorXType = "percent" // percent
    ScatterglErrorXType_constant ScatterglErrorXType = "constant" // constant
    ScatterglErrorXType_sqrt ScatterglErrorXType = "sqrt" // sqrt
    ScatterglErrorXType_data ScatterglErrorXType = "data" // data
)

// ScatterglErrorYType Determines the rule used to generate the error bars. If *constant`, the bar lengths are of a constant value. Set this constant in `value`. If *percent*, the bar lengths correspond to a percentage of underlying data. Set this percentage in `value`. If *sqrt*, the bar lengths correspond to the sqaure of the underlying data. If *data*, the bar lengths are set with data set `array`.
type ScatterglErrorYType string

const (
    ScatterglErrorYType_percent ScatterglErrorYType = "percent" // percent
    ScatterglErrorYType_constant ScatterglErrorYType = "constant" // constant
    ScatterglErrorYType_sqrt ScatterglErrorYType = "sqrt" // sqrt
    ScatterglErrorYType_data ScatterglErrorYType = "data" // data
)

// ScatterglFill Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
type ScatterglFill string

const (
    ScatterglFill_none ScatterglFill = "none" // none
    ScatterglFill_tozeroy ScatterglFill = "tozeroy" // tozeroy
    ScatterglFill_tozerox ScatterglFill = "tozerox" // tozerox
    ScatterglFill_tonexty ScatterglFill = "tonexty" // tonexty
    ScatterglFill_tonextx ScatterglFill = "tonextx" // tonextx
    ScatterglFill_toself ScatterglFill = "toself" // toself
    ScatterglFill_tonext ScatterglFill = "tonext" // tonext
)

// ScatterglHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterglHoverlabelAlign string

const (
    ScatterglHoverlabelAlign_left ScatterglHoverlabelAlign = "left" // left
    ScatterglHoverlabelAlign_right ScatterglHoverlabelAlign = "right" // right
    ScatterglHoverlabelAlign_auto ScatterglHoverlabelAlign = "auto" // auto
)

// ScatterglLineDash Sets the style of the lines.
type ScatterglLineDash string

const (
    ScatterglLineDash_solid ScatterglLineDash = "solid" // solid
    ScatterglLineDash_dot ScatterglLineDash = "dot" // dot
    ScatterglLineDash_dash ScatterglLineDash = "dash" // dash
    ScatterglLineDash_longdash ScatterglLineDash = "longdash" // longdash
    ScatterglLineDash_dashdot ScatterglLineDash = "dashdot" // dashdot
    ScatterglLineDash_longdashdot ScatterglLineDash = "longdashdot" // longdashdot
)

// ScatterglLineShape Determines the line shape. The values correspond to step-wise line shapes.
type ScatterglLineShape string

const (
    ScatterglLineShape_linear ScatterglLineShape = "linear" // linear
    ScatterglLineShape_hv ScatterglLineShape = "hv" // hv
    ScatterglLineShape_vh ScatterglLineShape = "vh" // vh
    ScatterglLineShape_hvh ScatterglLineShape = "hvh" // hvh
    ScatterglLineShape_vhv ScatterglLineShape = "vhv" // vhv
)

// ScatterglMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterglMarkerColorbarExponentformat string

const (
    ScatterglMarkerColorbarExponentformat_none ScatterglMarkerColorbarExponentformat = "none" // none
    ScatterglMarkerColorbarExponentformat_e ScatterglMarkerColorbarExponentformat = "e" // e
    ScatterglMarkerColorbarExponentformat_E ScatterglMarkerColorbarExponentformat = "E" // E
    ScatterglMarkerColorbarExponentformat_power ScatterglMarkerColorbarExponentformat = "power" // power
    ScatterglMarkerColorbarExponentformat_SI ScatterglMarkerColorbarExponentformat = "SI" // SI
    ScatterglMarkerColorbarExponentformat_B ScatterglMarkerColorbarExponentformat = "B" // B
)

// ScatterglMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterglMarkerColorbarLenmode string

const (
    ScatterglMarkerColorbarLenmode_fraction ScatterglMarkerColorbarLenmode = "fraction" // fraction
    ScatterglMarkerColorbarLenmode_pixels ScatterglMarkerColorbarLenmode = "pixels" // pixels
)

// ScatterglMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterglMarkerColorbarShowexponent string

const (
    ScatterglMarkerColorbarShowexponent_all ScatterglMarkerColorbarShowexponent = "all" // all
    ScatterglMarkerColorbarShowexponent_first ScatterglMarkerColorbarShowexponent = "first" // first
    ScatterglMarkerColorbarShowexponent_last ScatterglMarkerColorbarShowexponent = "last" // last
    ScatterglMarkerColorbarShowexponent_none ScatterglMarkerColorbarShowexponent = "none" // none
)

// ScatterglMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterglMarkerColorbarShowtickprefix string

const (
    ScatterglMarkerColorbarShowtickprefix_all ScatterglMarkerColorbarShowtickprefix = "all" // all
    ScatterglMarkerColorbarShowtickprefix_first ScatterglMarkerColorbarShowtickprefix = "first" // first
    ScatterglMarkerColorbarShowtickprefix_last ScatterglMarkerColorbarShowtickprefix = "last" // last
    ScatterglMarkerColorbarShowtickprefix_none ScatterglMarkerColorbarShowtickprefix = "none" // none
)

// ScatterglMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterglMarkerColorbarShowticksuffix string

const (
    ScatterglMarkerColorbarShowticksuffix_all ScatterglMarkerColorbarShowticksuffix = "all" // all
    ScatterglMarkerColorbarShowticksuffix_first ScatterglMarkerColorbarShowticksuffix = "first" // first
    ScatterglMarkerColorbarShowticksuffix_last ScatterglMarkerColorbarShowticksuffix = "last" // last
    ScatterglMarkerColorbarShowticksuffix_none ScatterglMarkerColorbarShowticksuffix = "none" // none
)

// ScatterglMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterglMarkerColorbarThicknessmode string

const (
    ScatterglMarkerColorbarThicknessmode_fraction ScatterglMarkerColorbarThicknessmode = "fraction" // fraction
    ScatterglMarkerColorbarThicknessmode_pixels ScatterglMarkerColorbarThicknessmode = "pixels" // pixels
)

// ScatterglMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterglMarkerColorbarTickmode string

const (
    ScatterglMarkerColorbarTickmode_auto ScatterglMarkerColorbarTickmode = "auto" // auto
    ScatterglMarkerColorbarTickmode_linear ScatterglMarkerColorbarTickmode = "linear" // linear
    ScatterglMarkerColorbarTickmode_array ScatterglMarkerColorbarTickmode = "array" // array
)

// ScatterglMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterglMarkerColorbarTicks string

const (
    ScatterglMarkerColorbarTicks_outside ScatterglMarkerColorbarTicks = "outside" // outside
    ScatterglMarkerColorbarTicks_inside ScatterglMarkerColorbarTicks = "inside" // inside
)

// ScatterglMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterglMarkerColorbarTitleSide string

const (
    ScatterglMarkerColorbarTitleSide_right ScatterglMarkerColorbarTitleSide = "right" // right
    ScatterglMarkerColorbarTitleSide_top ScatterglMarkerColorbarTitleSide = "top" // top
    ScatterglMarkerColorbarTitleSide_bottom ScatterglMarkerColorbarTitleSide = "bottom" // bottom
)

// ScatterglMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterglMarkerColorbarXanchor string

const (
    ScatterglMarkerColorbarXanchor_left ScatterglMarkerColorbarXanchor = "left" // left
    ScatterglMarkerColorbarXanchor_center ScatterglMarkerColorbarXanchor = "center" // center
    ScatterglMarkerColorbarXanchor_right ScatterglMarkerColorbarXanchor = "right" // right
)

// ScatterglMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterglMarkerColorbarYanchor string

const (
    ScatterglMarkerColorbarYanchor_top ScatterglMarkerColorbarYanchor = "top" // top
    ScatterglMarkerColorbarYanchor_middle ScatterglMarkerColorbarYanchor = "middle" // middle
    ScatterglMarkerColorbarYanchor_bottom ScatterglMarkerColorbarYanchor = "bottom" // bottom
)

// ScatterglMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterglMarkerSizemode string

const (
    ScatterglMarkerSizemode_diameter ScatterglMarkerSizemode = "diameter" // diameter
    ScatterglMarkerSizemode_area ScatterglMarkerSizemode = "area" // area
)

// ScatterglMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterglMarkerSymbol string

const (
    ScatterglMarkerSymbol0 ScatterglMarkerSymbol = "0" // 0
    ScatterglMarkerSymbol_circle ScatterglMarkerSymbol = "circle" // circle
    ScatterglMarkerSymbol100 ScatterglMarkerSymbol = "100" // 100
    ScatterglMarkerSymbol_circle_open ScatterglMarkerSymbol = "circle-open" // circle-open
    ScatterglMarkerSymbol200 ScatterglMarkerSymbol = "200" // 200
    ScatterglMarkerSymbol_circle_dot ScatterglMarkerSymbol = "circle-dot" // circle-dot
    ScatterglMarkerSymbol300 ScatterglMarkerSymbol = "300" // 300
    ScatterglMarkerSymbol_circle_open_dot ScatterglMarkerSymbol = "circle-open-dot" // circle-open-dot
    ScatterglMarkerSymbol1 ScatterglMarkerSymbol = "1" // 1
    ScatterglMarkerSymbol_square ScatterglMarkerSymbol = "square" // square
    ScatterglMarkerSymbol101 ScatterglMarkerSymbol = "101" // 101
    ScatterglMarkerSymbol_square_open ScatterglMarkerSymbol = "square-open" // square-open
    ScatterglMarkerSymbol201 ScatterglMarkerSymbol = "201" // 201
    ScatterglMarkerSymbol_square_dot ScatterglMarkerSymbol = "square-dot" // square-dot
    ScatterglMarkerSymbol301 ScatterglMarkerSymbol = "301" // 301
    ScatterglMarkerSymbol_square_open_dot ScatterglMarkerSymbol = "square-open-dot" // square-open-dot
    ScatterglMarkerSymbol2 ScatterglMarkerSymbol = "2" // 2
    ScatterglMarkerSymbol_diamond ScatterglMarkerSymbol = "diamond" // diamond
    ScatterglMarkerSymbol102 ScatterglMarkerSymbol = "102" // 102
    ScatterglMarkerSymbol_diamond_open ScatterglMarkerSymbol = "diamond-open" // diamond-open
    ScatterglMarkerSymbol202 ScatterglMarkerSymbol = "202" // 202
    ScatterglMarkerSymbol_diamond_dot ScatterglMarkerSymbol = "diamond-dot" // diamond-dot
    ScatterglMarkerSymbol302 ScatterglMarkerSymbol = "302" // 302
    ScatterglMarkerSymbol_diamond_open_dot ScatterglMarkerSymbol = "diamond-open-dot" // diamond-open-dot
    ScatterglMarkerSymbol3 ScatterglMarkerSymbol = "3" // 3
    ScatterglMarkerSymbol_cross ScatterglMarkerSymbol = "cross" // cross
    ScatterglMarkerSymbol103 ScatterglMarkerSymbol = "103" // 103
    ScatterglMarkerSymbol_cross_open ScatterglMarkerSymbol = "cross-open" // cross-open
    ScatterglMarkerSymbol203 ScatterglMarkerSymbol = "203" // 203
    ScatterglMarkerSymbol_cross_dot ScatterglMarkerSymbol = "cross-dot" // cross-dot
    ScatterglMarkerSymbol303 ScatterglMarkerSymbol = "303" // 303
    ScatterglMarkerSymbol_cross_open_dot ScatterglMarkerSymbol = "cross-open-dot" // cross-open-dot
    ScatterglMarkerSymbol4 ScatterglMarkerSymbol = "4" // 4
    ScatterglMarkerSymbol_x ScatterglMarkerSymbol = "x" // x
    ScatterglMarkerSymbol104 ScatterglMarkerSymbol = "104" // 104
    ScatterglMarkerSymbol_x_open ScatterglMarkerSymbol = "x-open" // x-open
    ScatterglMarkerSymbol204 ScatterglMarkerSymbol = "204" // 204
    ScatterglMarkerSymbol_x_dot ScatterglMarkerSymbol = "x-dot" // x-dot
    ScatterglMarkerSymbol304 ScatterglMarkerSymbol = "304" // 304
    ScatterglMarkerSymbol_x_open_dot ScatterglMarkerSymbol = "x-open-dot" // x-open-dot
    ScatterglMarkerSymbol5 ScatterglMarkerSymbol = "5" // 5
    ScatterglMarkerSymbol_triangle_up ScatterglMarkerSymbol = "triangle-up" // triangle-up
    ScatterglMarkerSymbol105 ScatterglMarkerSymbol = "105" // 105
    ScatterglMarkerSymbol_triangle_up_open ScatterglMarkerSymbol = "triangle-up-open" // triangle-up-open
    ScatterglMarkerSymbol205 ScatterglMarkerSymbol = "205" // 205
    ScatterglMarkerSymbol_triangle_up_dot ScatterglMarkerSymbol = "triangle-up-dot" // triangle-up-dot
    ScatterglMarkerSymbol305 ScatterglMarkerSymbol = "305" // 305
    ScatterglMarkerSymbol_triangle_up_open_dot ScatterglMarkerSymbol = "triangle-up-open-dot" // triangle-up-open-dot
    ScatterglMarkerSymbol6 ScatterglMarkerSymbol = "6" // 6
    ScatterglMarkerSymbol_triangle_down ScatterglMarkerSymbol = "triangle-down" // triangle-down
    ScatterglMarkerSymbol106 ScatterglMarkerSymbol = "106" // 106
    ScatterglMarkerSymbol_triangle_down_open ScatterglMarkerSymbol = "triangle-down-open" // triangle-down-open
    ScatterglMarkerSymbol206 ScatterglMarkerSymbol = "206" // 206
    ScatterglMarkerSymbol_triangle_down_dot ScatterglMarkerSymbol = "triangle-down-dot" // triangle-down-dot
    ScatterglMarkerSymbol306 ScatterglMarkerSymbol = "306" // 306
    ScatterglMarkerSymbol_triangle_down_open_dot ScatterglMarkerSymbol = "triangle-down-open-dot" // triangle-down-open-dot
    ScatterglMarkerSymbol7 ScatterglMarkerSymbol = "7" // 7
    ScatterglMarkerSymbol_triangle_left ScatterglMarkerSymbol = "triangle-left" // triangle-left
    ScatterglMarkerSymbol107 ScatterglMarkerSymbol = "107" // 107
    ScatterglMarkerSymbol_triangle_left_open ScatterglMarkerSymbol = "triangle-left-open" // triangle-left-open
    ScatterglMarkerSymbol207 ScatterglMarkerSymbol = "207" // 207
    ScatterglMarkerSymbol_triangle_left_dot ScatterglMarkerSymbol = "triangle-left-dot" // triangle-left-dot
    ScatterglMarkerSymbol307 ScatterglMarkerSymbol = "307" // 307
    ScatterglMarkerSymbol_triangle_left_open_dot ScatterglMarkerSymbol = "triangle-left-open-dot" // triangle-left-open-dot
    ScatterglMarkerSymbol8 ScatterglMarkerSymbol = "8" // 8
    ScatterglMarkerSymbol_triangle_right ScatterglMarkerSymbol = "triangle-right" // triangle-right
    ScatterglMarkerSymbol108 ScatterglMarkerSymbol = "108" // 108
    ScatterglMarkerSymbol_triangle_right_open ScatterglMarkerSymbol = "triangle-right-open" // triangle-right-open
    ScatterglMarkerSymbol208 ScatterglMarkerSymbol = "208" // 208
    ScatterglMarkerSymbol_triangle_right_dot ScatterglMarkerSymbol = "triangle-right-dot" // triangle-right-dot
    ScatterglMarkerSymbol308 ScatterglMarkerSymbol = "308" // 308
    ScatterglMarkerSymbol_triangle_right_open_dot ScatterglMarkerSymbol = "triangle-right-open-dot" // triangle-right-open-dot
    ScatterglMarkerSymbol9 ScatterglMarkerSymbol = "9" // 9
    ScatterglMarkerSymbol_triangle_ne ScatterglMarkerSymbol = "triangle-ne" // triangle-ne
    ScatterglMarkerSymbol109 ScatterglMarkerSymbol = "109" // 109
    ScatterglMarkerSymbol_triangle_ne_open ScatterglMarkerSymbol = "triangle-ne-open" // triangle-ne-open
    ScatterglMarkerSymbol209 ScatterglMarkerSymbol = "209" // 209
    ScatterglMarkerSymbol_triangle_ne_dot ScatterglMarkerSymbol = "triangle-ne-dot" // triangle-ne-dot
    ScatterglMarkerSymbol309 ScatterglMarkerSymbol = "309" // 309
    ScatterglMarkerSymbol_triangle_ne_open_dot ScatterglMarkerSymbol = "triangle-ne-open-dot" // triangle-ne-open-dot
    ScatterglMarkerSymbol10 ScatterglMarkerSymbol = "10" // 10
    ScatterglMarkerSymbol_triangle_se ScatterglMarkerSymbol = "triangle-se" // triangle-se
    ScatterglMarkerSymbol110 ScatterglMarkerSymbol = "110" // 110
    ScatterglMarkerSymbol_triangle_se_open ScatterglMarkerSymbol = "triangle-se-open" // triangle-se-open
    ScatterglMarkerSymbol210 ScatterglMarkerSymbol = "210" // 210
    ScatterglMarkerSymbol_triangle_se_dot ScatterglMarkerSymbol = "triangle-se-dot" // triangle-se-dot
    ScatterglMarkerSymbol310 ScatterglMarkerSymbol = "310" // 310
    ScatterglMarkerSymbol_triangle_se_open_dot ScatterglMarkerSymbol = "triangle-se-open-dot" // triangle-se-open-dot
    ScatterglMarkerSymbol11 ScatterglMarkerSymbol = "11" // 11
    ScatterglMarkerSymbol_triangle_sw ScatterglMarkerSymbol = "triangle-sw" // triangle-sw
    ScatterglMarkerSymbol111 ScatterglMarkerSymbol = "111" // 111
    ScatterglMarkerSymbol_triangle_sw_open ScatterglMarkerSymbol = "triangle-sw-open" // triangle-sw-open
    ScatterglMarkerSymbol211 ScatterglMarkerSymbol = "211" // 211
    ScatterglMarkerSymbol_triangle_sw_dot ScatterglMarkerSymbol = "triangle-sw-dot" // triangle-sw-dot
    ScatterglMarkerSymbol311 ScatterglMarkerSymbol = "311" // 311
    ScatterglMarkerSymbol_triangle_sw_open_dot ScatterglMarkerSymbol = "triangle-sw-open-dot" // triangle-sw-open-dot
    ScatterglMarkerSymbol12 ScatterglMarkerSymbol = "12" // 12
    ScatterglMarkerSymbol_triangle_nw ScatterglMarkerSymbol = "triangle-nw" // triangle-nw
    ScatterglMarkerSymbol112 ScatterglMarkerSymbol = "112" // 112
    ScatterglMarkerSymbol_triangle_nw_open ScatterglMarkerSymbol = "triangle-nw-open" // triangle-nw-open
    ScatterglMarkerSymbol212 ScatterglMarkerSymbol = "212" // 212
    ScatterglMarkerSymbol_triangle_nw_dot ScatterglMarkerSymbol = "triangle-nw-dot" // triangle-nw-dot
    ScatterglMarkerSymbol312 ScatterglMarkerSymbol = "312" // 312
    ScatterglMarkerSymbol_triangle_nw_open_dot ScatterglMarkerSymbol = "triangle-nw-open-dot" // triangle-nw-open-dot
    ScatterglMarkerSymbol13 ScatterglMarkerSymbol = "13" // 13
    ScatterglMarkerSymbol_pentagon ScatterglMarkerSymbol = "pentagon" // pentagon
    ScatterglMarkerSymbol113 ScatterglMarkerSymbol = "113" // 113
    ScatterglMarkerSymbol_pentagon_open ScatterglMarkerSymbol = "pentagon-open" // pentagon-open
    ScatterglMarkerSymbol213 ScatterglMarkerSymbol = "213" // 213
    ScatterglMarkerSymbol_pentagon_dot ScatterglMarkerSymbol = "pentagon-dot" // pentagon-dot
    ScatterglMarkerSymbol313 ScatterglMarkerSymbol = "313" // 313
    ScatterglMarkerSymbol_pentagon_open_dot ScatterglMarkerSymbol = "pentagon-open-dot" // pentagon-open-dot
    ScatterglMarkerSymbol14 ScatterglMarkerSymbol = "14" // 14
    ScatterglMarkerSymbol_hexagon ScatterglMarkerSymbol = "hexagon" // hexagon
    ScatterglMarkerSymbol114 ScatterglMarkerSymbol = "114" // 114
    ScatterglMarkerSymbol_hexagon_open ScatterglMarkerSymbol = "hexagon-open" // hexagon-open
    ScatterglMarkerSymbol214 ScatterglMarkerSymbol = "214" // 214
    ScatterglMarkerSymbol_hexagon_dot ScatterglMarkerSymbol = "hexagon-dot" // hexagon-dot
    ScatterglMarkerSymbol314 ScatterglMarkerSymbol = "314" // 314
    ScatterglMarkerSymbol_hexagon_open_dot ScatterglMarkerSymbol = "hexagon-open-dot" // hexagon-open-dot
    ScatterglMarkerSymbol15 ScatterglMarkerSymbol = "15" // 15
    ScatterglMarkerSymbol_hexagon2 ScatterglMarkerSymbol = "hexagon2" // hexagon2
    ScatterglMarkerSymbol115 ScatterglMarkerSymbol = "115" // 115
    ScatterglMarkerSymbol_hexagon2_open ScatterglMarkerSymbol = "hexagon2-open" // hexagon2-open
    ScatterglMarkerSymbol215 ScatterglMarkerSymbol = "215" // 215
    ScatterglMarkerSymbol_hexagon2_dot ScatterglMarkerSymbol = "hexagon2-dot" // hexagon2-dot
    ScatterglMarkerSymbol315 ScatterglMarkerSymbol = "315" // 315
    ScatterglMarkerSymbol_hexagon2_open_dot ScatterglMarkerSymbol = "hexagon2-open-dot" // hexagon2-open-dot
    ScatterglMarkerSymbol16 ScatterglMarkerSymbol = "16" // 16
    ScatterglMarkerSymbol_octagon ScatterglMarkerSymbol = "octagon" // octagon
    ScatterglMarkerSymbol116 ScatterglMarkerSymbol = "116" // 116
    ScatterglMarkerSymbol_octagon_open ScatterglMarkerSymbol = "octagon-open" // octagon-open
    ScatterglMarkerSymbol216 ScatterglMarkerSymbol = "216" // 216
    ScatterglMarkerSymbol_octagon_dot ScatterglMarkerSymbol = "octagon-dot" // octagon-dot
    ScatterglMarkerSymbol316 ScatterglMarkerSymbol = "316" // 316
    ScatterglMarkerSymbol_octagon_open_dot ScatterglMarkerSymbol = "octagon-open-dot" // octagon-open-dot
    ScatterglMarkerSymbol17 ScatterglMarkerSymbol = "17" // 17
    ScatterglMarkerSymbol_star ScatterglMarkerSymbol = "star" // star
    ScatterglMarkerSymbol117 ScatterglMarkerSymbol = "117" // 117
    ScatterglMarkerSymbol_star_open ScatterglMarkerSymbol = "star-open" // star-open
    ScatterglMarkerSymbol217 ScatterglMarkerSymbol = "217" // 217
    ScatterglMarkerSymbol_star_dot ScatterglMarkerSymbol = "star-dot" // star-dot
    ScatterglMarkerSymbol317 ScatterglMarkerSymbol = "317" // 317
    ScatterglMarkerSymbol_star_open_dot ScatterglMarkerSymbol = "star-open-dot" // star-open-dot
    ScatterglMarkerSymbol18 ScatterglMarkerSymbol = "18" // 18
    ScatterglMarkerSymbol_hexagram ScatterglMarkerSymbol = "hexagram" // hexagram
    ScatterglMarkerSymbol118 ScatterglMarkerSymbol = "118" // 118
    ScatterglMarkerSymbol_hexagram_open ScatterglMarkerSymbol = "hexagram-open" // hexagram-open
    ScatterglMarkerSymbol218 ScatterglMarkerSymbol = "218" // 218
    ScatterglMarkerSymbol_hexagram_dot ScatterglMarkerSymbol = "hexagram-dot" // hexagram-dot
    ScatterglMarkerSymbol318 ScatterglMarkerSymbol = "318" // 318
    ScatterglMarkerSymbol_hexagram_open_dot ScatterglMarkerSymbol = "hexagram-open-dot" // hexagram-open-dot
    ScatterglMarkerSymbol19 ScatterglMarkerSymbol = "19" // 19
    ScatterglMarkerSymbol_star_triangle_up ScatterglMarkerSymbol = "star-triangle-up" // star-triangle-up
    ScatterglMarkerSymbol119 ScatterglMarkerSymbol = "119" // 119
    ScatterglMarkerSymbol_star_triangle_up_open ScatterglMarkerSymbol = "star-triangle-up-open" // star-triangle-up-open
    ScatterglMarkerSymbol219 ScatterglMarkerSymbol = "219" // 219
    ScatterglMarkerSymbol_star_triangle_up_dot ScatterglMarkerSymbol = "star-triangle-up-dot" // star-triangle-up-dot
    ScatterglMarkerSymbol319 ScatterglMarkerSymbol = "319" // 319
    ScatterglMarkerSymbol_star_triangle_up_open_dot ScatterglMarkerSymbol = "star-triangle-up-open-dot" // star-triangle-up-open-dot
    ScatterglMarkerSymbol20 ScatterglMarkerSymbol = "20" // 20
    ScatterglMarkerSymbol_star_triangle_down ScatterglMarkerSymbol = "star-triangle-down" // star-triangle-down
    ScatterglMarkerSymbol120 ScatterglMarkerSymbol = "120" // 120
    ScatterglMarkerSymbol_star_triangle_down_open ScatterglMarkerSymbol = "star-triangle-down-open" // star-triangle-down-open
    ScatterglMarkerSymbol220 ScatterglMarkerSymbol = "220" // 220
    ScatterglMarkerSymbol_star_triangle_down_dot ScatterglMarkerSymbol = "star-triangle-down-dot" // star-triangle-down-dot
    ScatterglMarkerSymbol320 ScatterglMarkerSymbol = "320" // 320
    ScatterglMarkerSymbol_star_triangle_down_open_dot ScatterglMarkerSymbol = "star-triangle-down-open-dot" // star-triangle-down-open-dot
    ScatterglMarkerSymbol21 ScatterglMarkerSymbol = "21" // 21
    ScatterglMarkerSymbol_star_square ScatterglMarkerSymbol = "star-square" // star-square
    ScatterglMarkerSymbol121 ScatterglMarkerSymbol = "121" // 121
    ScatterglMarkerSymbol_star_square_open ScatterglMarkerSymbol = "star-square-open" // star-square-open
    ScatterglMarkerSymbol221 ScatterglMarkerSymbol = "221" // 221
    ScatterglMarkerSymbol_star_square_dot ScatterglMarkerSymbol = "star-square-dot" // star-square-dot
    ScatterglMarkerSymbol321 ScatterglMarkerSymbol = "321" // 321
    ScatterglMarkerSymbol_star_square_open_dot ScatterglMarkerSymbol = "star-square-open-dot" // star-square-open-dot
    ScatterglMarkerSymbol22 ScatterglMarkerSymbol = "22" // 22
    ScatterglMarkerSymbol_star_diamond ScatterglMarkerSymbol = "star-diamond" // star-diamond
    ScatterglMarkerSymbol122 ScatterglMarkerSymbol = "122" // 122
    ScatterglMarkerSymbol_star_diamond_open ScatterglMarkerSymbol = "star-diamond-open" // star-diamond-open
    ScatterglMarkerSymbol222 ScatterglMarkerSymbol = "222" // 222
    ScatterglMarkerSymbol_star_diamond_dot ScatterglMarkerSymbol = "star-diamond-dot" // star-diamond-dot
    ScatterglMarkerSymbol322 ScatterglMarkerSymbol = "322" // 322
    ScatterglMarkerSymbol_star_diamond_open_dot ScatterglMarkerSymbol = "star-diamond-open-dot" // star-diamond-open-dot
    ScatterglMarkerSymbol23 ScatterglMarkerSymbol = "23" // 23
    ScatterglMarkerSymbol_diamond_tall ScatterglMarkerSymbol = "diamond-tall" // diamond-tall
    ScatterglMarkerSymbol123 ScatterglMarkerSymbol = "123" // 123
    ScatterglMarkerSymbol_diamond_tall_open ScatterglMarkerSymbol = "diamond-tall-open" // diamond-tall-open
    ScatterglMarkerSymbol223 ScatterglMarkerSymbol = "223" // 223
    ScatterglMarkerSymbol_diamond_tall_dot ScatterglMarkerSymbol = "diamond-tall-dot" // diamond-tall-dot
    ScatterglMarkerSymbol323 ScatterglMarkerSymbol = "323" // 323
    ScatterglMarkerSymbol_diamond_tall_open_dot ScatterglMarkerSymbol = "diamond-tall-open-dot" // diamond-tall-open-dot
    ScatterglMarkerSymbol24 ScatterglMarkerSymbol = "24" // 24
    ScatterglMarkerSymbol_diamond_wide ScatterglMarkerSymbol = "diamond-wide" // diamond-wide
    ScatterglMarkerSymbol124 ScatterglMarkerSymbol = "124" // 124
    ScatterglMarkerSymbol_diamond_wide_open ScatterglMarkerSymbol = "diamond-wide-open" // diamond-wide-open
    ScatterglMarkerSymbol224 ScatterglMarkerSymbol = "224" // 224
    ScatterglMarkerSymbol_diamond_wide_dot ScatterglMarkerSymbol = "diamond-wide-dot" // diamond-wide-dot
    ScatterglMarkerSymbol324 ScatterglMarkerSymbol = "324" // 324
    ScatterglMarkerSymbol_diamond_wide_open_dot ScatterglMarkerSymbol = "diamond-wide-open-dot" // diamond-wide-open-dot
    ScatterglMarkerSymbol25 ScatterglMarkerSymbol = "25" // 25
    ScatterglMarkerSymbol_hourglass ScatterglMarkerSymbol = "hourglass" // hourglass
    ScatterglMarkerSymbol125 ScatterglMarkerSymbol = "125" // 125
    ScatterglMarkerSymbol_hourglass_open ScatterglMarkerSymbol = "hourglass-open" // hourglass-open
    ScatterglMarkerSymbol26 ScatterglMarkerSymbol = "26" // 26
    ScatterglMarkerSymbol_bowtie ScatterglMarkerSymbol = "bowtie" // bowtie
    ScatterglMarkerSymbol126 ScatterglMarkerSymbol = "126" // 126
    ScatterglMarkerSymbol_bowtie_open ScatterglMarkerSymbol = "bowtie-open" // bowtie-open
    ScatterglMarkerSymbol27 ScatterglMarkerSymbol = "27" // 27
    ScatterglMarkerSymbol_circle_cross ScatterglMarkerSymbol = "circle-cross" // circle-cross
    ScatterglMarkerSymbol127 ScatterglMarkerSymbol = "127" // 127
    ScatterglMarkerSymbol_circle_cross_open ScatterglMarkerSymbol = "circle-cross-open" // circle-cross-open
    ScatterglMarkerSymbol28 ScatterglMarkerSymbol = "28" // 28
    ScatterglMarkerSymbol_circle_x ScatterglMarkerSymbol = "circle-x" // circle-x
    ScatterglMarkerSymbol128 ScatterglMarkerSymbol = "128" // 128
    ScatterglMarkerSymbol_circle_x_open ScatterglMarkerSymbol = "circle-x-open" // circle-x-open
    ScatterglMarkerSymbol29 ScatterglMarkerSymbol = "29" // 29
    ScatterglMarkerSymbol_square_cross ScatterglMarkerSymbol = "square-cross" // square-cross
    ScatterglMarkerSymbol129 ScatterglMarkerSymbol = "129" // 129
    ScatterglMarkerSymbol_square_cross_open ScatterglMarkerSymbol = "square-cross-open" // square-cross-open
    ScatterglMarkerSymbol30 ScatterglMarkerSymbol = "30" // 30
    ScatterglMarkerSymbol_square_x ScatterglMarkerSymbol = "square-x" // square-x
    ScatterglMarkerSymbol130 ScatterglMarkerSymbol = "130" // 130
    ScatterglMarkerSymbol_square_x_open ScatterglMarkerSymbol = "square-x-open" // square-x-open
    ScatterglMarkerSymbol31 ScatterglMarkerSymbol = "31" // 31
    ScatterglMarkerSymbol_diamond_cross ScatterglMarkerSymbol = "diamond-cross" // diamond-cross
    ScatterglMarkerSymbol131 ScatterglMarkerSymbol = "131" // 131
    ScatterglMarkerSymbol_diamond_cross_open ScatterglMarkerSymbol = "diamond-cross-open" // diamond-cross-open
    ScatterglMarkerSymbol32 ScatterglMarkerSymbol = "32" // 32
    ScatterglMarkerSymbol_diamond_x ScatterglMarkerSymbol = "diamond-x" // diamond-x
    ScatterglMarkerSymbol132 ScatterglMarkerSymbol = "132" // 132
    ScatterglMarkerSymbol_diamond_x_open ScatterglMarkerSymbol = "diamond-x-open" // diamond-x-open
    ScatterglMarkerSymbol33 ScatterglMarkerSymbol = "33" // 33
    ScatterglMarkerSymbol_cross_thin ScatterglMarkerSymbol = "cross-thin" // cross-thin
    ScatterglMarkerSymbol133 ScatterglMarkerSymbol = "133" // 133
    ScatterglMarkerSymbol_cross_thin_open ScatterglMarkerSymbol = "cross-thin-open" // cross-thin-open
    ScatterglMarkerSymbol34 ScatterglMarkerSymbol = "34" // 34
    ScatterglMarkerSymbol_x_thin ScatterglMarkerSymbol = "x-thin" // x-thin
    ScatterglMarkerSymbol134 ScatterglMarkerSymbol = "134" // 134
    ScatterglMarkerSymbol_x_thin_open ScatterglMarkerSymbol = "x-thin-open" // x-thin-open
    ScatterglMarkerSymbol35 ScatterglMarkerSymbol = "35" // 35
    ScatterglMarkerSymbol_asterisk ScatterglMarkerSymbol = "asterisk" // asterisk
    ScatterglMarkerSymbol135 ScatterglMarkerSymbol = "135" // 135
    ScatterglMarkerSymbol_asterisk_open ScatterglMarkerSymbol = "asterisk-open" // asterisk-open
    ScatterglMarkerSymbol36 ScatterglMarkerSymbol = "36" // 36
    ScatterglMarkerSymbol_hash ScatterglMarkerSymbol = "hash" // hash
    ScatterglMarkerSymbol136 ScatterglMarkerSymbol = "136" // 136
    ScatterglMarkerSymbol_hash_open ScatterglMarkerSymbol = "hash-open" // hash-open
    ScatterglMarkerSymbol236 ScatterglMarkerSymbol = "236" // 236
    ScatterglMarkerSymbol_hash_dot ScatterglMarkerSymbol = "hash-dot" // hash-dot
    ScatterglMarkerSymbol336 ScatterglMarkerSymbol = "336" // 336
    ScatterglMarkerSymbol_hash_open_dot ScatterglMarkerSymbol = "hash-open-dot" // hash-open-dot
    ScatterglMarkerSymbol37 ScatterglMarkerSymbol = "37" // 37
    ScatterglMarkerSymbol_y_up ScatterglMarkerSymbol = "y-up" // y-up
    ScatterglMarkerSymbol137 ScatterglMarkerSymbol = "137" // 137
    ScatterglMarkerSymbol_y_up_open ScatterglMarkerSymbol = "y-up-open" // y-up-open
    ScatterglMarkerSymbol38 ScatterglMarkerSymbol = "38" // 38
    ScatterglMarkerSymbol_y_down ScatterglMarkerSymbol = "y-down" // y-down
    ScatterglMarkerSymbol138 ScatterglMarkerSymbol = "138" // 138
    ScatterglMarkerSymbol_y_down_open ScatterglMarkerSymbol = "y-down-open" // y-down-open
    ScatterglMarkerSymbol39 ScatterglMarkerSymbol = "39" // 39
    ScatterglMarkerSymbol_y_left ScatterglMarkerSymbol = "y-left" // y-left
    ScatterglMarkerSymbol139 ScatterglMarkerSymbol = "139" // 139
    ScatterglMarkerSymbol_y_left_open ScatterglMarkerSymbol = "y-left-open" // y-left-open
    ScatterglMarkerSymbol40 ScatterglMarkerSymbol = "40" // 40
    ScatterglMarkerSymbol_y_right ScatterglMarkerSymbol = "y-right" // y-right
    ScatterglMarkerSymbol140 ScatterglMarkerSymbol = "140" // 140
    ScatterglMarkerSymbol_y_right_open ScatterglMarkerSymbol = "y-right-open" // y-right-open
    ScatterglMarkerSymbol41 ScatterglMarkerSymbol = "41" // 41
    ScatterglMarkerSymbol_line_ew ScatterglMarkerSymbol = "line-ew" // line-ew
    ScatterglMarkerSymbol141 ScatterglMarkerSymbol = "141" // 141
    ScatterglMarkerSymbol_line_ew_open ScatterglMarkerSymbol = "line-ew-open" // line-ew-open
    ScatterglMarkerSymbol42 ScatterglMarkerSymbol = "42" // 42
    ScatterglMarkerSymbol_line_ns ScatterglMarkerSymbol = "line-ns" // line-ns
    ScatterglMarkerSymbol142 ScatterglMarkerSymbol = "142" // 142
    ScatterglMarkerSymbol_line_ns_open ScatterglMarkerSymbol = "line-ns-open" // line-ns-open
    ScatterglMarkerSymbol43 ScatterglMarkerSymbol = "43" // 43
    ScatterglMarkerSymbol_line_ne ScatterglMarkerSymbol = "line-ne" // line-ne
    ScatterglMarkerSymbol143 ScatterglMarkerSymbol = "143" // 143
    ScatterglMarkerSymbol_line_ne_open ScatterglMarkerSymbol = "line-ne-open" // line-ne-open
    ScatterglMarkerSymbol44 ScatterglMarkerSymbol = "44" // 44
    ScatterglMarkerSymbol_line_nw ScatterglMarkerSymbol = "line-nw" // line-nw
    ScatterglMarkerSymbol144 ScatterglMarkerSymbol = "144" // 144
    ScatterglMarkerSymbol_line_nw_open ScatterglMarkerSymbol = "line-nw-open" // line-nw-open
)

// ScatterglTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterglTextposition string

const (
    ScatterglTextposition_topleft ScatterglTextposition = "top left" // top left
    ScatterglTextposition_topcenter ScatterglTextposition = "top center" // top center
    ScatterglTextposition_topright ScatterglTextposition = "top right" // top right
    ScatterglTextposition_middleleft ScatterglTextposition = "middle left" // middle left
    ScatterglTextposition_middlecenter ScatterglTextposition = "middle center" // middle center
    ScatterglTextposition_middleright ScatterglTextposition = "middle right" // middle right
    ScatterglTextposition_bottomleft ScatterglTextposition = "bottom left" // bottom left
    ScatterglTextposition_bottomcenter ScatterglTextposition = "bottom center" // bottom center
    ScatterglTextposition_bottomright ScatterglTextposition = "bottom right" // bottom right
)

// ScatterglVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterglVisible interface{}

const (
    ScatterglVisibleTrue bool = true
    ScatterglVisibleFalse bool = false
    ScatterglVisibleLegendonly string = "legendonly"
)

// ScatterglXcalendar Sets the calendar system to use with `x` date data.
type ScatterglXcalendar string

const (
    ScatterglXcalendar_gregorian ScatterglXcalendar = "gregorian" // gregorian
    ScatterglXcalendar_chinese ScatterglXcalendar = "chinese" // chinese
    ScatterglXcalendar_coptic ScatterglXcalendar = "coptic" // coptic
    ScatterglXcalendar_discworld ScatterglXcalendar = "discworld" // discworld
    ScatterglXcalendar_ethiopian ScatterglXcalendar = "ethiopian" // ethiopian
    ScatterglXcalendar_hebrew ScatterglXcalendar = "hebrew" // hebrew
    ScatterglXcalendar_islamic ScatterglXcalendar = "islamic" // islamic
    ScatterglXcalendar_julian ScatterglXcalendar = "julian" // julian
    ScatterglXcalendar_mayan ScatterglXcalendar = "mayan" // mayan
    ScatterglXcalendar_nanakshahi ScatterglXcalendar = "nanakshahi" // nanakshahi
    ScatterglXcalendar_nepali ScatterglXcalendar = "nepali" // nepali
    ScatterglXcalendar_persian ScatterglXcalendar = "persian" // persian
    ScatterglXcalendar_jalali ScatterglXcalendar = "jalali" // jalali
    ScatterglXcalendar_taiwan ScatterglXcalendar = "taiwan" // taiwan
    ScatterglXcalendar_thai ScatterglXcalendar = "thai" // thai
    ScatterglXcalendar_ummalqura ScatterglXcalendar = "ummalqura" // ummalqura
)

// ScatterglYcalendar Sets the calendar system to use with `y` date data.
type ScatterglYcalendar string

const (
    ScatterglYcalendar_gregorian ScatterglYcalendar = "gregorian" // gregorian
    ScatterglYcalendar_chinese ScatterglYcalendar = "chinese" // chinese
    ScatterglYcalendar_coptic ScatterglYcalendar = "coptic" // coptic
    ScatterglYcalendar_discworld ScatterglYcalendar = "discworld" // discworld
    ScatterglYcalendar_ethiopian ScatterglYcalendar = "ethiopian" // ethiopian
    ScatterglYcalendar_hebrew ScatterglYcalendar = "hebrew" // hebrew
    ScatterglYcalendar_islamic ScatterglYcalendar = "islamic" // islamic
    ScatterglYcalendar_julian ScatterglYcalendar = "julian" // julian
    ScatterglYcalendar_mayan ScatterglYcalendar = "mayan" // mayan
    ScatterglYcalendar_nanakshahi ScatterglYcalendar = "nanakshahi" // nanakshahi
    ScatterglYcalendar_nepali ScatterglYcalendar = "nepali" // nepali
    ScatterglYcalendar_persian ScatterglYcalendar = "persian" // persian
    ScatterglYcalendar_jalali ScatterglYcalendar = "jalali" // jalali
    ScatterglYcalendar_taiwan ScatterglYcalendar = "taiwan" // taiwan
    ScatterglYcalendar_thai ScatterglYcalendar = "thai" // thai
    ScatterglYcalendar_ummalqura ScatterglYcalendar = "ummalqura" // ummalqura
)

// ScattermapboxFill Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape.
type ScattermapboxFill string

const (
    ScattermapboxFill_none ScattermapboxFill = "none" // none
    ScattermapboxFill_toself ScattermapboxFill = "toself" // toself
)

// ScattermapboxHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScattermapboxHoverlabelAlign string

const (
    ScattermapboxHoverlabelAlign_left ScattermapboxHoverlabelAlign = "left" // left
    ScattermapboxHoverlabelAlign_right ScattermapboxHoverlabelAlign = "right" // right
    ScattermapboxHoverlabelAlign_auto ScattermapboxHoverlabelAlign = "auto" // auto
)

// ScattermapboxMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScattermapboxMarkerColorbarExponentformat string

const (
    ScattermapboxMarkerColorbarExponentformat_none ScattermapboxMarkerColorbarExponentformat = "none" // none
    ScattermapboxMarkerColorbarExponentformat_e ScattermapboxMarkerColorbarExponentformat = "e" // e
    ScattermapboxMarkerColorbarExponentformat_E ScattermapboxMarkerColorbarExponentformat = "E" // E
    ScattermapboxMarkerColorbarExponentformat_power ScattermapboxMarkerColorbarExponentformat = "power" // power
    ScattermapboxMarkerColorbarExponentformat_SI ScattermapboxMarkerColorbarExponentformat = "SI" // SI
    ScattermapboxMarkerColorbarExponentformat_B ScattermapboxMarkerColorbarExponentformat = "B" // B
)

// ScattermapboxMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScattermapboxMarkerColorbarLenmode string

const (
    ScattermapboxMarkerColorbarLenmode_fraction ScattermapboxMarkerColorbarLenmode = "fraction" // fraction
    ScattermapboxMarkerColorbarLenmode_pixels ScattermapboxMarkerColorbarLenmode = "pixels" // pixels
)

// ScattermapboxMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScattermapboxMarkerColorbarShowexponent string

const (
    ScattermapboxMarkerColorbarShowexponent_all ScattermapboxMarkerColorbarShowexponent = "all" // all
    ScattermapboxMarkerColorbarShowexponent_first ScattermapboxMarkerColorbarShowexponent = "first" // first
    ScattermapboxMarkerColorbarShowexponent_last ScattermapboxMarkerColorbarShowexponent = "last" // last
    ScattermapboxMarkerColorbarShowexponent_none ScattermapboxMarkerColorbarShowexponent = "none" // none
)

// ScattermapboxMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScattermapboxMarkerColorbarShowtickprefix string

const (
    ScattermapboxMarkerColorbarShowtickprefix_all ScattermapboxMarkerColorbarShowtickprefix = "all" // all
    ScattermapboxMarkerColorbarShowtickprefix_first ScattermapboxMarkerColorbarShowtickprefix = "first" // first
    ScattermapboxMarkerColorbarShowtickprefix_last ScattermapboxMarkerColorbarShowtickprefix = "last" // last
    ScattermapboxMarkerColorbarShowtickprefix_none ScattermapboxMarkerColorbarShowtickprefix = "none" // none
)

// ScattermapboxMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScattermapboxMarkerColorbarShowticksuffix string

const (
    ScattermapboxMarkerColorbarShowticksuffix_all ScattermapboxMarkerColorbarShowticksuffix = "all" // all
    ScattermapboxMarkerColorbarShowticksuffix_first ScattermapboxMarkerColorbarShowticksuffix = "first" // first
    ScattermapboxMarkerColorbarShowticksuffix_last ScattermapboxMarkerColorbarShowticksuffix = "last" // last
    ScattermapboxMarkerColorbarShowticksuffix_none ScattermapboxMarkerColorbarShowticksuffix = "none" // none
)

// ScattermapboxMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScattermapboxMarkerColorbarThicknessmode string

const (
    ScattermapboxMarkerColorbarThicknessmode_fraction ScattermapboxMarkerColorbarThicknessmode = "fraction" // fraction
    ScattermapboxMarkerColorbarThicknessmode_pixels ScattermapboxMarkerColorbarThicknessmode = "pixels" // pixels
)

// ScattermapboxMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScattermapboxMarkerColorbarTickmode string

const (
    ScattermapboxMarkerColorbarTickmode_auto ScattermapboxMarkerColorbarTickmode = "auto" // auto
    ScattermapboxMarkerColorbarTickmode_linear ScattermapboxMarkerColorbarTickmode = "linear" // linear
    ScattermapboxMarkerColorbarTickmode_array ScattermapboxMarkerColorbarTickmode = "array" // array
)

// ScattermapboxMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScattermapboxMarkerColorbarTicks string

const (
    ScattermapboxMarkerColorbarTicks_outside ScattermapboxMarkerColorbarTicks = "outside" // outside
    ScattermapboxMarkerColorbarTicks_inside ScattermapboxMarkerColorbarTicks = "inside" // inside
)

// ScattermapboxMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScattermapboxMarkerColorbarTitleSide string

const (
    ScattermapboxMarkerColorbarTitleSide_right ScattermapboxMarkerColorbarTitleSide = "right" // right
    ScattermapboxMarkerColorbarTitleSide_top ScattermapboxMarkerColorbarTitleSide = "top" // top
    ScattermapboxMarkerColorbarTitleSide_bottom ScattermapboxMarkerColorbarTitleSide = "bottom" // bottom
)

// ScattermapboxMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScattermapboxMarkerColorbarXanchor string

const (
    ScattermapboxMarkerColorbarXanchor_left ScattermapboxMarkerColorbarXanchor = "left" // left
    ScattermapboxMarkerColorbarXanchor_center ScattermapboxMarkerColorbarXanchor = "center" // center
    ScattermapboxMarkerColorbarXanchor_right ScattermapboxMarkerColorbarXanchor = "right" // right
)

// ScattermapboxMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScattermapboxMarkerColorbarYanchor string

const (
    ScattermapboxMarkerColorbarYanchor_top ScattermapboxMarkerColorbarYanchor = "top" // top
    ScattermapboxMarkerColorbarYanchor_middle ScattermapboxMarkerColorbarYanchor = "middle" // middle
    ScattermapboxMarkerColorbarYanchor_bottom ScattermapboxMarkerColorbarYanchor = "bottom" // bottom
)

// ScattermapboxMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScattermapboxMarkerSizemode string

const (
    ScattermapboxMarkerSizemode_diameter ScattermapboxMarkerSizemode = "diameter" // diameter
    ScattermapboxMarkerSizemode_area ScattermapboxMarkerSizemode = "area" // area
)

// ScattermapboxTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScattermapboxTextposition string

const (
    ScattermapboxTextposition_topleft ScattermapboxTextposition = "top left" // top left
    ScattermapboxTextposition_topcenter ScattermapboxTextposition = "top center" // top center
    ScattermapboxTextposition_topright ScattermapboxTextposition = "top right" // top right
    ScattermapboxTextposition_middleleft ScattermapboxTextposition = "middle left" // middle left
    ScattermapboxTextposition_middlecenter ScattermapboxTextposition = "middle center" // middle center
    ScattermapboxTextposition_middleright ScattermapboxTextposition = "middle right" // middle right
    ScattermapboxTextposition_bottomleft ScattermapboxTextposition = "bottom left" // bottom left
    ScattermapboxTextposition_bottomcenter ScattermapboxTextposition = "bottom center" // bottom center
    ScattermapboxTextposition_bottomright ScattermapboxTextposition = "bottom right" // bottom right
)

// ScattermapboxVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScattermapboxVisible interface{}

const (
    ScattermapboxVisibleTrue bool = true
    ScattermapboxVisibleFalse bool = false
    ScattermapboxVisibleLegendonly string = "legendonly"
)

// ScatterpolarFill Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. scatterpolar has a subset of the options available to scatter. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other.
type ScatterpolarFill string

const (
    ScatterpolarFill_none ScatterpolarFill = "none" // none
    ScatterpolarFill_toself ScatterpolarFill = "toself" // toself
    ScatterpolarFill_tonext ScatterpolarFill = "tonext" // tonext
)

// ScatterpolarHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterpolarHoverlabelAlign string

const (
    ScatterpolarHoverlabelAlign_left ScatterpolarHoverlabelAlign = "left" // left
    ScatterpolarHoverlabelAlign_right ScatterpolarHoverlabelAlign = "right" // right
    ScatterpolarHoverlabelAlign_auto ScatterpolarHoverlabelAlign = "auto" // auto
)

// ScatterpolarLineShape Determines the line shape. With *spline* the lines are drawn using spline interpolation. The other available values correspond to step-wise line shapes.
type ScatterpolarLineShape string

const (
    ScatterpolarLineShape_linear ScatterpolarLineShape = "linear" // linear
    ScatterpolarLineShape_spline ScatterpolarLineShape = "spline" // spline
)

// ScatterpolarMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterpolarMarkerColorbarExponentformat string

const (
    ScatterpolarMarkerColorbarExponentformat_none ScatterpolarMarkerColorbarExponentformat = "none" // none
    ScatterpolarMarkerColorbarExponentformat_e ScatterpolarMarkerColorbarExponentformat = "e" // e
    ScatterpolarMarkerColorbarExponentformat_E ScatterpolarMarkerColorbarExponentformat = "E" // E
    ScatterpolarMarkerColorbarExponentformat_power ScatterpolarMarkerColorbarExponentformat = "power" // power
    ScatterpolarMarkerColorbarExponentformat_SI ScatterpolarMarkerColorbarExponentformat = "SI" // SI
    ScatterpolarMarkerColorbarExponentformat_B ScatterpolarMarkerColorbarExponentformat = "B" // B
)

// ScatterpolarMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterpolarMarkerColorbarLenmode string

const (
    ScatterpolarMarkerColorbarLenmode_fraction ScatterpolarMarkerColorbarLenmode = "fraction" // fraction
    ScatterpolarMarkerColorbarLenmode_pixels ScatterpolarMarkerColorbarLenmode = "pixels" // pixels
)

// ScatterpolarMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterpolarMarkerColorbarShowexponent string

const (
    ScatterpolarMarkerColorbarShowexponent_all ScatterpolarMarkerColorbarShowexponent = "all" // all
    ScatterpolarMarkerColorbarShowexponent_first ScatterpolarMarkerColorbarShowexponent = "first" // first
    ScatterpolarMarkerColorbarShowexponent_last ScatterpolarMarkerColorbarShowexponent = "last" // last
    ScatterpolarMarkerColorbarShowexponent_none ScatterpolarMarkerColorbarShowexponent = "none" // none
)

// ScatterpolarMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterpolarMarkerColorbarShowtickprefix string

const (
    ScatterpolarMarkerColorbarShowtickprefix_all ScatterpolarMarkerColorbarShowtickprefix = "all" // all
    ScatterpolarMarkerColorbarShowtickprefix_first ScatterpolarMarkerColorbarShowtickprefix = "first" // first
    ScatterpolarMarkerColorbarShowtickprefix_last ScatterpolarMarkerColorbarShowtickprefix = "last" // last
    ScatterpolarMarkerColorbarShowtickprefix_none ScatterpolarMarkerColorbarShowtickprefix = "none" // none
)

// ScatterpolarMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterpolarMarkerColorbarShowticksuffix string

const (
    ScatterpolarMarkerColorbarShowticksuffix_all ScatterpolarMarkerColorbarShowticksuffix = "all" // all
    ScatterpolarMarkerColorbarShowticksuffix_first ScatterpolarMarkerColorbarShowticksuffix = "first" // first
    ScatterpolarMarkerColorbarShowticksuffix_last ScatterpolarMarkerColorbarShowticksuffix = "last" // last
    ScatterpolarMarkerColorbarShowticksuffix_none ScatterpolarMarkerColorbarShowticksuffix = "none" // none
)

// ScatterpolarMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterpolarMarkerColorbarThicknessmode string

const (
    ScatterpolarMarkerColorbarThicknessmode_fraction ScatterpolarMarkerColorbarThicknessmode = "fraction" // fraction
    ScatterpolarMarkerColorbarThicknessmode_pixels ScatterpolarMarkerColorbarThicknessmode = "pixels" // pixels
)

// ScatterpolarMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterpolarMarkerColorbarTickmode string

const (
    ScatterpolarMarkerColorbarTickmode_auto ScatterpolarMarkerColorbarTickmode = "auto" // auto
    ScatterpolarMarkerColorbarTickmode_linear ScatterpolarMarkerColorbarTickmode = "linear" // linear
    ScatterpolarMarkerColorbarTickmode_array ScatterpolarMarkerColorbarTickmode = "array" // array
)

// ScatterpolarMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterpolarMarkerColorbarTicks string

const (
    ScatterpolarMarkerColorbarTicks_outside ScatterpolarMarkerColorbarTicks = "outside" // outside
    ScatterpolarMarkerColorbarTicks_inside ScatterpolarMarkerColorbarTicks = "inside" // inside
)

// ScatterpolarMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterpolarMarkerColorbarTitleSide string

const (
    ScatterpolarMarkerColorbarTitleSide_right ScatterpolarMarkerColorbarTitleSide = "right" // right
    ScatterpolarMarkerColorbarTitleSide_top ScatterpolarMarkerColorbarTitleSide = "top" // top
    ScatterpolarMarkerColorbarTitleSide_bottom ScatterpolarMarkerColorbarTitleSide = "bottom" // bottom
)

// ScatterpolarMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterpolarMarkerColorbarXanchor string

const (
    ScatterpolarMarkerColorbarXanchor_left ScatterpolarMarkerColorbarXanchor = "left" // left
    ScatterpolarMarkerColorbarXanchor_center ScatterpolarMarkerColorbarXanchor = "center" // center
    ScatterpolarMarkerColorbarXanchor_right ScatterpolarMarkerColorbarXanchor = "right" // right
)

// ScatterpolarMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterpolarMarkerColorbarYanchor string

const (
    ScatterpolarMarkerColorbarYanchor_top ScatterpolarMarkerColorbarYanchor = "top" // top
    ScatterpolarMarkerColorbarYanchor_middle ScatterpolarMarkerColorbarYanchor = "middle" // middle
    ScatterpolarMarkerColorbarYanchor_bottom ScatterpolarMarkerColorbarYanchor = "bottom" // bottom
)

// ScatterpolarMarkerGradientType Sets the type of gradient used to fill the markers
type ScatterpolarMarkerGradientType string

const (
    ScatterpolarMarkerGradientType_radial ScatterpolarMarkerGradientType = "radial" // radial
    ScatterpolarMarkerGradientType_horizontal ScatterpolarMarkerGradientType = "horizontal" // horizontal
    ScatterpolarMarkerGradientType_vertical ScatterpolarMarkerGradientType = "vertical" // vertical
    ScatterpolarMarkerGradientType_none ScatterpolarMarkerGradientType = "none" // none
)

// ScatterpolarMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterpolarMarkerSizemode string

const (
    ScatterpolarMarkerSizemode_diameter ScatterpolarMarkerSizemode = "diameter" // diameter
    ScatterpolarMarkerSizemode_area ScatterpolarMarkerSizemode = "area" // area
)

// ScatterpolarMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterpolarMarkerSymbol string

const (
    ScatterpolarMarkerSymbol0 ScatterpolarMarkerSymbol = "0" // 0
    ScatterpolarMarkerSymbol_circle ScatterpolarMarkerSymbol = "circle" // circle
    ScatterpolarMarkerSymbol100 ScatterpolarMarkerSymbol = "100" // 100
    ScatterpolarMarkerSymbol_circle_open ScatterpolarMarkerSymbol = "circle-open" // circle-open
    ScatterpolarMarkerSymbol200 ScatterpolarMarkerSymbol = "200" // 200
    ScatterpolarMarkerSymbol_circle_dot ScatterpolarMarkerSymbol = "circle-dot" // circle-dot
    ScatterpolarMarkerSymbol300 ScatterpolarMarkerSymbol = "300" // 300
    ScatterpolarMarkerSymbol_circle_open_dot ScatterpolarMarkerSymbol = "circle-open-dot" // circle-open-dot
    ScatterpolarMarkerSymbol1 ScatterpolarMarkerSymbol = "1" // 1
    ScatterpolarMarkerSymbol_square ScatterpolarMarkerSymbol = "square" // square
    ScatterpolarMarkerSymbol101 ScatterpolarMarkerSymbol = "101" // 101
    ScatterpolarMarkerSymbol_square_open ScatterpolarMarkerSymbol = "square-open" // square-open
    ScatterpolarMarkerSymbol201 ScatterpolarMarkerSymbol = "201" // 201
    ScatterpolarMarkerSymbol_square_dot ScatterpolarMarkerSymbol = "square-dot" // square-dot
    ScatterpolarMarkerSymbol301 ScatterpolarMarkerSymbol = "301" // 301
    ScatterpolarMarkerSymbol_square_open_dot ScatterpolarMarkerSymbol = "square-open-dot" // square-open-dot
    ScatterpolarMarkerSymbol2 ScatterpolarMarkerSymbol = "2" // 2
    ScatterpolarMarkerSymbol_diamond ScatterpolarMarkerSymbol = "diamond" // diamond
    ScatterpolarMarkerSymbol102 ScatterpolarMarkerSymbol = "102" // 102
    ScatterpolarMarkerSymbol_diamond_open ScatterpolarMarkerSymbol = "diamond-open" // diamond-open
    ScatterpolarMarkerSymbol202 ScatterpolarMarkerSymbol = "202" // 202
    ScatterpolarMarkerSymbol_diamond_dot ScatterpolarMarkerSymbol = "diamond-dot" // diamond-dot
    ScatterpolarMarkerSymbol302 ScatterpolarMarkerSymbol = "302" // 302
    ScatterpolarMarkerSymbol_diamond_open_dot ScatterpolarMarkerSymbol = "diamond-open-dot" // diamond-open-dot
    ScatterpolarMarkerSymbol3 ScatterpolarMarkerSymbol = "3" // 3
    ScatterpolarMarkerSymbol_cross ScatterpolarMarkerSymbol = "cross" // cross
    ScatterpolarMarkerSymbol103 ScatterpolarMarkerSymbol = "103" // 103
    ScatterpolarMarkerSymbol_cross_open ScatterpolarMarkerSymbol = "cross-open" // cross-open
    ScatterpolarMarkerSymbol203 ScatterpolarMarkerSymbol = "203" // 203
    ScatterpolarMarkerSymbol_cross_dot ScatterpolarMarkerSymbol = "cross-dot" // cross-dot
    ScatterpolarMarkerSymbol303 ScatterpolarMarkerSymbol = "303" // 303
    ScatterpolarMarkerSymbol_cross_open_dot ScatterpolarMarkerSymbol = "cross-open-dot" // cross-open-dot
    ScatterpolarMarkerSymbol4 ScatterpolarMarkerSymbol = "4" // 4
    ScatterpolarMarkerSymbol_x ScatterpolarMarkerSymbol = "x" // x
    ScatterpolarMarkerSymbol104 ScatterpolarMarkerSymbol = "104" // 104
    ScatterpolarMarkerSymbol_x_open ScatterpolarMarkerSymbol = "x-open" // x-open
    ScatterpolarMarkerSymbol204 ScatterpolarMarkerSymbol = "204" // 204
    ScatterpolarMarkerSymbol_x_dot ScatterpolarMarkerSymbol = "x-dot" // x-dot
    ScatterpolarMarkerSymbol304 ScatterpolarMarkerSymbol = "304" // 304
    ScatterpolarMarkerSymbol_x_open_dot ScatterpolarMarkerSymbol = "x-open-dot" // x-open-dot
    ScatterpolarMarkerSymbol5 ScatterpolarMarkerSymbol = "5" // 5
    ScatterpolarMarkerSymbol_triangle_up ScatterpolarMarkerSymbol = "triangle-up" // triangle-up
    ScatterpolarMarkerSymbol105 ScatterpolarMarkerSymbol = "105" // 105
    ScatterpolarMarkerSymbol_triangle_up_open ScatterpolarMarkerSymbol = "triangle-up-open" // triangle-up-open
    ScatterpolarMarkerSymbol205 ScatterpolarMarkerSymbol = "205" // 205
    ScatterpolarMarkerSymbol_triangle_up_dot ScatterpolarMarkerSymbol = "triangle-up-dot" // triangle-up-dot
    ScatterpolarMarkerSymbol305 ScatterpolarMarkerSymbol = "305" // 305
    ScatterpolarMarkerSymbol_triangle_up_open_dot ScatterpolarMarkerSymbol = "triangle-up-open-dot" // triangle-up-open-dot
    ScatterpolarMarkerSymbol6 ScatterpolarMarkerSymbol = "6" // 6
    ScatterpolarMarkerSymbol_triangle_down ScatterpolarMarkerSymbol = "triangle-down" // triangle-down
    ScatterpolarMarkerSymbol106 ScatterpolarMarkerSymbol = "106" // 106
    ScatterpolarMarkerSymbol_triangle_down_open ScatterpolarMarkerSymbol = "triangle-down-open" // triangle-down-open
    ScatterpolarMarkerSymbol206 ScatterpolarMarkerSymbol = "206" // 206
    ScatterpolarMarkerSymbol_triangle_down_dot ScatterpolarMarkerSymbol = "triangle-down-dot" // triangle-down-dot
    ScatterpolarMarkerSymbol306 ScatterpolarMarkerSymbol = "306" // 306
    ScatterpolarMarkerSymbol_triangle_down_open_dot ScatterpolarMarkerSymbol = "triangle-down-open-dot" // triangle-down-open-dot
    ScatterpolarMarkerSymbol7 ScatterpolarMarkerSymbol = "7" // 7
    ScatterpolarMarkerSymbol_triangle_left ScatterpolarMarkerSymbol = "triangle-left" // triangle-left
    ScatterpolarMarkerSymbol107 ScatterpolarMarkerSymbol = "107" // 107
    ScatterpolarMarkerSymbol_triangle_left_open ScatterpolarMarkerSymbol = "triangle-left-open" // triangle-left-open
    ScatterpolarMarkerSymbol207 ScatterpolarMarkerSymbol = "207" // 207
    ScatterpolarMarkerSymbol_triangle_left_dot ScatterpolarMarkerSymbol = "triangle-left-dot" // triangle-left-dot
    ScatterpolarMarkerSymbol307 ScatterpolarMarkerSymbol = "307" // 307
    ScatterpolarMarkerSymbol_triangle_left_open_dot ScatterpolarMarkerSymbol = "triangle-left-open-dot" // triangle-left-open-dot
    ScatterpolarMarkerSymbol8 ScatterpolarMarkerSymbol = "8" // 8
    ScatterpolarMarkerSymbol_triangle_right ScatterpolarMarkerSymbol = "triangle-right" // triangle-right
    ScatterpolarMarkerSymbol108 ScatterpolarMarkerSymbol = "108" // 108
    ScatterpolarMarkerSymbol_triangle_right_open ScatterpolarMarkerSymbol = "triangle-right-open" // triangle-right-open
    ScatterpolarMarkerSymbol208 ScatterpolarMarkerSymbol = "208" // 208
    ScatterpolarMarkerSymbol_triangle_right_dot ScatterpolarMarkerSymbol = "triangle-right-dot" // triangle-right-dot
    ScatterpolarMarkerSymbol308 ScatterpolarMarkerSymbol = "308" // 308
    ScatterpolarMarkerSymbol_triangle_right_open_dot ScatterpolarMarkerSymbol = "triangle-right-open-dot" // triangle-right-open-dot
    ScatterpolarMarkerSymbol9 ScatterpolarMarkerSymbol = "9" // 9
    ScatterpolarMarkerSymbol_triangle_ne ScatterpolarMarkerSymbol = "triangle-ne" // triangle-ne
    ScatterpolarMarkerSymbol109 ScatterpolarMarkerSymbol = "109" // 109
    ScatterpolarMarkerSymbol_triangle_ne_open ScatterpolarMarkerSymbol = "triangle-ne-open" // triangle-ne-open
    ScatterpolarMarkerSymbol209 ScatterpolarMarkerSymbol = "209" // 209
    ScatterpolarMarkerSymbol_triangle_ne_dot ScatterpolarMarkerSymbol = "triangle-ne-dot" // triangle-ne-dot
    ScatterpolarMarkerSymbol309 ScatterpolarMarkerSymbol = "309" // 309
    ScatterpolarMarkerSymbol_triangle_ne_open_dot ScatterpolarMarkerSymbol = "triangle-ne-open-dot" // triangle-ne-open-dot
    ScatterpolarMarkerSymbol10 ScatterpolarMarkerSymbol = "10" // 10
    ScatterpolarMarkerSymbol_triangle_se ScatterpolarMarkerSymbol = "triangle-se" // triangle-se
    ScatterpolarMarkerSymbol110 ScatterpolarMarkerSymbol = "110" // 110
    ScatterpolarMarkerSymbol_triangle_se_open ScatterpolarMarkerSymbol = "triangle-se-open" // triangle-se-open
    ScatterpolarMarkerSymbol210 ScatterpolarMarkerSymbol = "210" // 210
    ScatterpolarMarkerSymbol_triangle_se_dot ScatterpolarMarkerSymbol = "triangle-se-dot" // triangle-se-dot
    ScatterpolarMarkerSymbol310 ScatterpolarMarkerSymbol = "310" // 310
    ScatterpolarMarkerSymbol_triangle_se_open_dot ScatterpolarMarkerSymbol = "triangle-se-open-dot" // triangle-se-open-dot
    ScatterpolarMarkerSymbol11 ScatterpolarMarkerSymbol = "11" // 11
    ScatterpolarMarkerSymbol_triangle_sw ScatterpolarMarkerSymbol = "triangle-sw" // triangle-sw
    ScatterpolarMarkerSymbol111 ScatterpolarMarkerSymbol = "111" // 111
    ScatterpolarMarkerSymbol_triangle_sw_open ScatterpolarMarkerSymbol = "triangle-sw-open" // triangle-sw-open
    ScatterpolarMarkerSymbol211 ScatterpolarMarkerSymbol = "211" // 211
    ScatterpolarMarkerSymbol_triangle_sw_dot ScatterpolarMarkerSymbol = "triangle-sw-dot" // triangle-sw-dot
    ScatterpolarMarkerSymbol311 ScatterpolarMarkerSymbol = "311" // 311
    ScatterpolarMarkerSymbol_triangle_sw_open_dot ScatterpolarMarkerSymbol = "triangle-sw-open-dot" // triangle-sw-open-dot
    ScatterpolarMarkerSymbol12 ScatterpolarMarkerSymbol = "12" // 12
    ScatterpolarMarkerSymbol_triangle_nw ScatterpolarMarkerSymbol = "triangle-nw" // triangle-nw
    ScatterpolarMarkerSymbol112 ScatterpolarMarkerSymbol = "112" // 112
    ScatterpolarMarkerSymbol_triangle_nw_open ScatterpolarMarkerSymbol = "triangle-nw-open" // triangle-nw-open
    ScatterpolarMarkerSymbol212 ScatterpolarMarkerSymbol = "212" // 212
    ScatterpolarMarkerSymbol_triangle_nw_dot ScatterpolarMarkerSymbol = "triangle-nw-dot" // triangle-nw-dot
    ScatterpolarMarkerSymbol312 ScatterpolarMarkerSymbol = "312" // 312
    ScatterpolarMarkerSymbol_triangle_nw_open_dot ScatterpolarMarkerSymbol = "triangle-nw-open-dot" // triangle-nw-open-dot
    ScatterpolarMarkerSymbol13 ScatterpolarMarkerSymbol = "13" // 13
    ScatterpolarMarkerSymbol_pentagon ScatterpolarMarkerSymbol = "pentagon" // pentagon
    ScatterpolarMarkerSymbol113 ScatterpolarMarkerSymbol = "113" // 113
    ScatterpolarMarkerSymbol_pentagon_open ScatterpolarMarkerSymbol = "pentagon-open" // pentagon-open
    ScatterpolarMarkerSymbol213 ScatterpolarMarkerSymbol = "213" // 213
    ScatterpolarMarkerSymbol_pentagon_dot ScatterpolarMarkerSymbol = "pentagon-dot" // pentagon-dot
    ScatterpolarMarkerSymbol313 ScatterpolarMarkerSymbol = "313" // 313
    ScatterpolarMarkerSymbol_pentagon_open_dot ScatterpolarMarkerSymbol = "pentagon-open-dot" // pentagon-open-dot
    ScatterpolarMarkerSymbol14 ScatterpolarMarkerSymbol = "14" // 14
    ScatterpolarMarkerSymbol_hexagon ScatterpolarMarkerSymbol = "hexagon" // hexagon
    ScatterpolarMarkerSymbol114 ScatterpolarMarkerSymbol = "114" // 114
    ScatterpolarMarkerSymbol_hexagon_open ScatterpolarMarkerSymbol = "hexagon-open" // hexagon-open
    ScatterpolarMarkerSymbol214 ScatterpolarMarkerSymbol = "214" // 214
    ScatterpolarMarkerSymbol_hexagon_dot ScatterpolarMarkerSymbol = "hexagon-dot" // hexagon-dot
    ScatterpolarMarkerSymbol314 ScatterpolarMarkerSymbol = "314" // 314
    ScatterpolarMarkerSymbol_hexagon_open_dot ScatterpolarMarkerSymbol = "hexagon-open-dot" // hexagon-open-dot
    ScatterpolarMarkerSymbol15 ScatterpolarMarkerSymbol = "15" // 15
    ScatterpolarMarkerSymbol_hexagon2 ScatterpolarMarkerSymbol = "hexagon2" // hexagon2
    ScatterpolarMarkerSymbol115 ScatterpolarMarkerSymbol = "115" // 115
    ScatterpolarMarkerSymbol_hexagon2_open ScatterpolarMarkerSymbol = "hexagon2-open" // hexagon2-open
    ScatterpolarMarkerSymbol215 ScatterpolarMarkerSymbol = "215" // 215
    ScatterpolarMarkerSymbol_hexagon2_dot ScatterpolarMarkerSymbol = "hexagon2-dot" // hexagon2-dot
    ScatterpolarMarkerSymbol315 ScatterpolarMarkerSymbol = "315" // 315
    ScatterpolarMarkerSymbol_hexagon2_open_dot ScatterpolarMarkerSymbol = "hexagon2-open-dot" // hexagon2-open-dot
    ScatterpolarMarkerSymbol16 ScatterpolarMarkerSymbol = "16" // 16
    ScatterpolarMarkerSymbol_octagon ScatterpolarMarkerSymbol = "octagon" // octagon
    ScatterpolarMarkerSymbol116 ScatterpolarMarkerSymbol = "116" // 116
    ScatterpolarMarkerSymbol_octagon_open ScatterpolarMarkerSymbol = "octagon-open" // octagon-open
    ScatterpolarMarkerSymbol216 ScatterpolarMarkerSymbol = "216" // 216
    ScatterpolarMarkerSymbol_octagon_dot ScatterpolarMarkerSymbol = "octagon-dot" // octagon-dot
    ScatterpolarMarkerSymbol316 ScatterpolarMarkerSymbol = "316" // 316
    ScatterpolarMarkerSymbol_octagon_open_dot ScatterpolarMarkerSymbol = "octagon-open-dot" // octagon-open-dot
    ScatterpolarMarkerSymbol17 ScatterpolarMarkerSymbol = "17" // 17
    ScatterpolarMarkerSymbol_star ScatterpolarMarkerSymbol = "star" // star
    ScatterpolarMarkerSymbol117 ScatterpolarMarkerSymbol = "117" // 117
    ScatterpolarMarkerSymbol_star_open ScatterpolarMarkerSymbol = "star-open" // star-open
    ScatterpolarMarkerSymbol217 ScatterpolarMarkerSymbol = "217" // 217
    ScatterpolarMarkerSymbol_star_dot ScatterpolarMarkerSymbol = "star-dot" // star-dot
    ScatterpolarMarkerSymbol317 ScatterpolarMarkerSymbol = "317" // 317
    ScatterpolarMarkerSymbol_star_open_dot ScatterpolarMarkerSymbol = "star-open-dot" // star-open-dot
    ScatterpolarMarkerSymbol18 ScatterpolarMarkerSymbol = "18" // 18
    ScatterpolarMarkerSymbol_hexagram ScatterpolarMarkerSymbol = "hexagram" // hexagram
    ScatterpolarMarkerSymbol118 ScatterpolarMarkerSymbol = "118" // 118
    ScatterpolarMarkerSymbol_hexagram_open ScatterpolarMarkerSymbol = "hexagram-open" // hexagram-open
    ScatterpolarMarkerSymbol218 ScatterpolarMarkerSymbol = "218" // 218
    ScatterpolarMarkerSymbol_hexagram_dot ScatterpolarMarkerSymbol = "hexagram-dot" // hexagram-dot
    ScatterpolarMarkerSymbol318 ScatterpolarMarkerSymbol = "318" // 318
    ScatterpolarMarkerSymbol_hexagram_open_dot ScatterpolarMarkerSymbol = "hexagram-open-dot" // hexagram-open-dot
    ScatterpolarMarkerSymbol19 ScatterpolarMarkerSymbol = "19" // 19
    ScatterpolarMarkerSymbol_star_triangle_up ScatterpolarMarkerSymbol = "star-triangle-up" // star-triangle-up
    ScatterpolarMarkerSymbol119 ScatterpolarMarkerSymbol = "119" // 119
    ScatterpolarMarkerSymbol_star_triangle_up_open ScatterpolarMarkerSymbol = "star-triangle-up-open" // star-triangle-up-open
    ScatterpolarMarkerSymbol219 ScatterpolarMarkerSymbol = "219" // 219
    ScatterpolarMarkerSymbol_star_triangle_up_dot ScatterpolarMarkerSymbol = "star-triangle-up-dot" // star-triangle-up-dot
    ScatterpolarMarkerSymbol319 ScatterpolarMarkerSymbol = "319" // 319
    ScatterpolarMarkerSymbol_star_triangle_up_open_dot ScatterpolarMarkerSymbol = "star-triangle-up-open-dot" // star-triangle-up-open-dot
    ScatterpolarMarkerSymbol20 ScatterpolarMarkerSymbol = "20" // 20
    ScatterpolarMarkerSymbol_star_triangle_down ScatterpolarMarkerSymbol = "star-triangle-down" // star-triangle-down
    ScatterpolarMarkerSymbol120 ScatterpolarMarkerSymbol = "120" // 120
    ScatterpolarMarkerSymbol_star_triangle_down_open ScatterpolarMarkerSymbol = "star-triangle-down-open" // star-triangle-down-open
    ScatterpolarMarkerSymbol220 ScatterpolarMarkerSymbol = "220" // 220
    ScatterpolarMarkerSymbol_star_triangle_down_dot ScatterpolarMarkerSymbol = "star-triangle-down-dot" // star-triangle-down-dot
    ScatterpolarMarkerSymbol320 ScatterpolarMarkerSymbol = "320" // 320
    ScatterpolarMarkerSymbol_star_triangle_down_open_dot ScatterpolarMarkerSymbol = "star-triangle-down-open-dot" // star-triangle-down-open-dot
    ScatterpolarMarkerSymbol21 ScatterpolarMarkerSymbol = "21" // 21
    ScatterpolarMarkerSymbol_star_square ScatterpolarMarkerSymbol = "star-square" // star-square
    ScatterpolarMarkerSymbol121 ScatterpolarMarkerSymbol = "121" // 121
    ScatterpolarMarkerSymbol_star_square_open ScatterpolarMarkerSymbol = "star-square-open" // star-square-open
    ScatterpolarMarkerSymbol221 ScatterpolarMarkerSymbol = "221" // 221
    ScatterpolarMarkerSymbol_star_square_dot ScatterpolarMarkerSymbol = "star-square-dot" // star-square-dot
    ScatterpolarMarkerSymbol321 ScatterpolarMarkerSymbol = "321" // 321
    ScatterpolarMarkerSymbol_star_square_open_dot ScatterpolarMarkerSymbol = "star-square-open-dot" // star-square-open-dot
    ScatterpolarMarkerSymbol22 ScatterpolarMarkerSymbol = "22" // 22
    ScatterpolarMarkerSymbol_star_diamond ScatterpolarMarkerSymbol = "star-diamond" // star-diamond
    ScatterpolarMarkerSymbol122 ScatterpolarMarkerSymbol = "122" // 122
    ScatterpolarMarkerSymbol_star_diamond_open ScatterpolarMarkerSymbol = "star-diamond-open" // star-diamond-open
    ScatterpolarMarkerSymbol222 ScatterpolarMarkerSymbol = "222" // 222
    ScatterpolarMarkerSymbol_star_diamond_dot ScatterpolarMarkerSymbol = "star-diamond-dot" // star-diamond-dot
    ScatterpolarMarkerSymbol322 ScatterpolarMarkerSymbol = "322" // 322
    ScatterpolarMarkerSymbol_star_diamond_open_dot ScatterpolarMarkerSymbol = "star-diamond-open-dot" // star-diamond-open-dot
    ScatterpolarMarkerSymbol23 ScatterpolarMarkerSymbol = "23" // 23
    ScatterpolarMarkerSymbol_diamond_tall ScatterpolarMarkerSymbol = "diamond-tall" // diamond-tall
    ScatterpolarMarkerSymbol123 ScatterpolarMarkerSymbol = "123" // 123
    ScatterpolarMarkerSymbol_diamond_tall_open ScatterpolarMarkerSymbol = "diamond-tall-open" // diamond-tall-open
    ScatterpolarMarkerSymbol223 ScatterpolarMarkerSymbol = "223" // 223
    ScatterpolarMarkerSymbol_diamond_tall_dot ScatterpolarMarkerSymbol = "diamond-tall-dot" // diamond-tall-dot
    ScatterpolarMarkerSymbol323 ScatterpolarMarkerSymbol = "323" // 323
    ScatterpolarMarkerSymbol_diamond_tall_open_dot ScatterpolarMarkerSymbol = "diamond-tall-open-dot" // diamond-tall-open-dot
    ScatterpolarMarkerSymbol24 ScatterpolarMarkerSymbol = "24" // 24
    ScatterpolarMarkerSymbol_diamond_wide ScatterpolarMarkerSymbol = "diamond-wide" // diamond-wide
    ScatterpolarMarkerSymbol124 ScatterpolarMarkerSymbol = "124" // 124
    ScatterpolarMarkerSymbol_diamond_wide_open ScatterpolarMarkerSymbol = "diamond-wide-open" // diamond-wide-open
    ScatterpolarMarkerSymbol224 ScatterpolarMarkerSymbol = "224" // 224
    ScatterpolarMarkerSymbol_diamond_wide_dot ScatterpolarMarkerSymbol = "diamond-wide-dot" // diamond-wide-dot
    ScatterpolarMarkerSymbol324 ScatterpolarMarkerSymbol = "324" // 324
    ScatterpolarMarkerSymbol_diamond_wide_open_dot ScatterpolarMarkerSymbol = "diamond-wide-open-dot" // diamond-wide-open-dot
    ScatterpolarMarkerSymbol25 ScatterpolarMarkerSymbol = "25" // 25
    ScatterpolarMarkerSymbol_hourglass ScatterpolarMarkerSymbol = "hourglass" // hourglass
    ScatterpolarMarkerSymbol125 ScatterpolarMarkerSymbol = "125" // 125
    ScatterpolarMarkerSymbol_hourglass_open ScatterpolarMarkerSymbol = "hourglass-open" // hourglass-open
    ScatterpolarMarkerSymbol26 ScatterpolarMarkerSymbol = "26" // 26
    ScatterpolarMarkerSymbol_bowtie ScatterpolarMarkerSymbol = "bowtie" // bowtie
    ScatterpolarMarkerSymbol126 ScatterpolarMarkerSymbol = "126" // 126
    ScatterpolarMarkerSymbol_bowtie_open ScatterpolarMarkerSymbol = "bowtie-open" // bowtie-open
    ScatterpolarMarkerSymbol27 ScatterpolarMarkerSymbol = "27" // 27
    ScatterpolarMarkerSymbol_circle_cross ScatterpolarMarkerSymbol = "circle-cross" // circle-cross
    ScatterpolarMarkerSymbol127 ScatterpolarMarkerSymbol = "127" // 127
    ScatterpolarMarkerSymbol_circle_cross_open ScatterpolarMarkerSymbol = "circle-cross-open" // circle-cross-open
    ScatterpolarMarkerSymbol28 ScatterpolarMarkerSymbol = "28" // 28
    ScatterpolarMarkerSymbol_circle_x ScatterpolarMarkerSymbol = "circle-x" // circle-x
    ScatterpolarMarkerSymbol128 ScatterpolarMarkerSymbol = "128" // 128
    ScatterpolarMarkerSymbol_circle_x_open ScatterpolarMarkerSymbol = "circle-x-open" // circle-x-open
    ScatterpolarMarkerSymbol29 ScatterpolarMarkerSymbol = "29" // 29
    ScatterpolarMarkerSymbol_square_cross ScatterpolarMarkerSymbol = "square-cross" // square-cross
    ScatterpolarMarkerSymbol129 ScatterpolarMarkerSymbol = "129" // 129
    ScatterpolarMarkerSymbol_square_cross_open ScatterpolarMarkerSymbol = "square-cross-open" // square-cross-open
    ScatterpolarMarkerSymbol30 ScatterpolarMarkerSymbol = "30" // 30
    ScatterpolarMarkerSymbol_square_x ScatterpolarMarkerSymbol = "square-x" // square-x
    ScatterpolarMarkerSymbol130 ScatterpolarMarkerSymbol = "130" // 130
    ScatterpolarMarkerSymbol_square_x_open ScatterpolarMarkerSymbol = "square-x-open" // square-x-open
    ScatterpolarMarkerSymbol31 ScatterpolarMarkerSymbol = "31" // 31
    ScatterpolarMarkerSymbol_diamond_cross ScatterpolarMarkerSymbol = "diamond-cross" // diamond-cross
    ScatterpolarMarkerSymbol131 ScatterpolarMarkerSymbol = "131" // 131
    ScatterpolarMarkerSymbol_diamond_cross_open ScatterpolarMarkerSymbol = "diamond-cross-open" // diamond-cross-open
    ScatterpolarMarkerSymbol32 ScatterpolarMarkerSymbol = "32" // 32
    ScatterpolarMarkerSymbol_diamond_x ScatterpolarMarkerSymbol = "diamond-x" // diamond-x
    ScatterpolarMarkerSymbol132 ScatterpolarMarkerSymbol = "132" // 132
    ScatterpolarMarkerSymbol_diamond_x_open ScatterpolarMarkerSymbol = "diamond-x-open" // diamond-x-open
    ScatterpolarMarkerSymbol33 ScatterpolarMarkerSymbol = "33" // 33
    ScatterpolarMarkerSymbol_cross_thin ScatterpolarMarkerSymbol = "cross-thin" // cross-thin
    ScatterpolarMarkerSymbol133 ScatterpolarMarkerSymbol = "133" // 133
    ScatterpolarMarkerSymbol_cross_thin_open ScatterpolarMarkerSymbol = "cross-thin-open" // cross-thin-open
    ScatterpolarMarkerSymbol34 ScatterpolarMarkerSymbol = "34" // 34
    ScatterpolarMarkerSymbol_x_thin ScatterpolarMarkerSymbol = "x-thin" // x-thin
    ScatterpolarMarkerSymbol134 ScatterpolarMarkerSymbol = "134" // 134
    ScatterpolarMarkerSymbol_x_thin_open ScatterpolarMarkerSymbol = "x-thin-open" // x-thin-open
    ScatterpolarMarkerSymbol35 ScatterpolarMarkerSymbol = "35" // 35
    ScatterpolarMarkerSymbol_asterisk ScatterpolarMarkerSymbol = "asterisk" // asterisk
    ScatterpolarMarkerSymbol135 ScatterpolarMarkerSymbol = "135" // 135
    ScatterpolarMarkerSymbol_asterisk_open ScatterpolarMarkerSymbol = "asterisk-open" // asterisk-open
    ScatterpolarMarkerSymbol36 ScatterpolarMarkerSymbol = "36" // 36
    ScatterpolarMarkerSymbol_hash ScatterpolarMarkerSymbol = "hash" // hash
    ScatterpolarMarkerSymbol136 ScatterpolarMarkerSymbol = "136" // 136
    ScatterpolarMarkerSymbol_hash_open ScatterpolarMarkerSymbol = "hash-open" // hash-open
    ScatterpolarMarkerSymbol236 ScatterpolarMarkerSymbol = "236" // 236
    ScatterpolarMarkerSymbol_hash_dot ScatterpolarMarkerSymbol = "hash-dot" // hash-dot
    ScatterpolarMarkerSymbol336 ScatterpolarMarkerSymbol = "336" // 336
    ScatterpolarMarkerSymbol_hash_open_dot ScatterpolarMarkerSymbol = "hash-open-dot" // hash-open-dot
    ScatterpolarMarkerSymbol37 ScatterpolarMarkerSymbol = "37" // 37
    ScatterpolarMarkerSymbol_y_up ScatterpolarMarkerSymbol = "y-up" // y-up
    ScatterpolarMarkerSymbol137 ScatterpolarMarkerSymbol = "137" // 137
    ScatterpolarMarkerSymbol_y_up_open ScatterpolarMarkerSymbol = "y-up-open" // y-up-open
    ScatterpolarMarkerSymbol38 ScatterpolarMarkerSymbol = "38" // 38
    ScatterpolarMarkerSymbol_y_down ScatterpolarMarkerSymbol = "y-down" // y-down
    ScatterpolarMarkerSymbol138 ScatterpolarMarkerSymbol = "138" // 138
    ScatterpolarMarkerSymbol_y_down_open ScatterpolarMarkerSymbol = "y-down-open" // y-down-open
    ScatterpolarMarkerSymbol39 ScatterpolarMarkerSymbol = "39" // 39
    ScatterpolarMarkerSymbol_y_left ScatterpolarMarkerSymbol = "y-left" // y-left
    ScatterpolarMarkerSymbol139 ScatterpolarMarkerSymbol = "139" // 139
    ScatterpolarMarkerSymbol_y_left_open ScatterpolarMarkerSymbol = "y-left-open" // y-left-open
    ScatterpolarMarkerSymbol40 ScatterpolarMarkerSymbol = "40" // 40
    ScatterpolarMarkerSymbol_y_right ScatterpolarMarkerSymbol = "y-right" // y-right
    ScatterpolarMarkerSymbol140 ScatterpolarMarkerSymbol = "140" // 140
    ScatterpolarMarkerSymbol_y_right_open ScatterpolarMarkerSymbol = "y-right-open" // y-right-open
    ScatterpolarMarkerSymbol41 ScatterpolarMarkerSymbol = "41" // 41
    ScatterpolarMarkerSymbol_line_ew ScatterpolarMarkerSymbol = "line-ew" // line-ew
    ScatterpolarMarkerSymbol141 ScatterpolarMarkerSymbol = "141" // 141
    ScatterpolarMarkerSymbol_line_ew_open ScatterpolarMarkerSymbol = "line-ew-open" // line-ew-open
    ScatterpolarMarkerSymbol42 ScatterpolarMarkerSymbol = "42" // 42
    ScatterpolarMarkerSymbol_line_ns ScatterpolarMarkerSymbol = "line-ns" // line-ns
    ScatterpolarMarkerSymbol142 ScatterpolarMarkerSymbol = "142" // 142
    ScatterpolarMarkerSymbol_line_ns_open ScatterpolarMarkerSymbol = "line-ns-open" // line-ns-open
    ScatterpolarMarkerSymbol43 ScatterpolarMarkerSymbol = "43" // 43
    ScatterpolarMarkerSymbol_line_ne ScatterpolarMarkerSymbol = "line-ne" // line-ne
    ScatterpolarMarkerSymbol143 ScatterpolarMarkerSymbol = "143" // 143
    ScatterpolarMarkerSymbol_line_ne_open ScatterpolarMarkerSymbol = "line-ne-open" // line-ne-open
    ScatterpolarMarkerSymbol44 ScatterpolarMarkerSymbol = "44" // 44
    ScatterpolarMarkerSymbol_line_nw ScatterpolarMarkerSymbol = "line-nw" // line-nw
    ScatterpolarMarkerSymbol144 ScatterpolarMarkerSymbol = "144" // 144
    ScatterpolarMarkerSymbol_line_nw_open ScatterpolarMarkerSymbol = "line-nw-open" // line-nw-open
)

// ScatterpolarTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterpolarTextposition string

const (
    ScatterpolarTextposition_topleft ScatterpolarTextposition = "top left" // top left
    ScatterpolarTextposition_topcenter ScatterpolarTextposition = "top center" // top center
    ScatterpolarTextposition_topright ScatterpolarTextposition = "top right" // top right
    ScatterpolarTextposition_middleleft ScatterpolarTextposition = "middle left" // middle left
    ScatterpolarTextposition_middlecenter ScatterpolarTextposition = "middle center" // middle center
    ScatterpolarTextposition_middleright ScatterpolarTextposition = "middle right" // middle right
    ScatterpolarTextposition_bottomleft ScatterpolarTextposition = "bottom left" // bottom left
    ScatterpolarTextposition_bottomcenter ScatterpolarTextposition = "bottom center" // bottom center
    ScatterpolarTextposition_bottomright ScatterpolarTextposition = "bottom right" // bottom right
)

// ScatterpolarThetaunit Sets the unit of input *theta* values. Has an effect only when on *linear* angular axes.
type ScatterpolarThetaunit string

const (
    ScatterpolarThetaunit_radians ScatterpolarThetaunit = "radians" // radians
    ScatterpolarThetaunit_degrees ScatterpolarThetaunit = "degrees" // degrees
    ScatterpolarThetaunit_gradians ScatterpolarThetaunit = "gradians" // gradians
)

// ScatterpolarVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterpolarVisible interface{}

const (
    ScatterpolarVisibleTrue bool = true
    ScatterpolarVisibleFalse bool = false
    ScatterpolarVisibleLegendonly string = "legendonly"
)

// ScatterpolarglFill Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
type ScatterpolarglFill string

const (
    ScatterpolarglFill_none ScatterpolarglFill = "none" // none
    ScatterpolarglFill_tozeroy ScatterpolarglFill = "tozeroy" // tozeroy
    ScatterpolarglFill_tozerox ScatterpolarglFill = "tozerox" // tozerox
    ScatterpolarglFill_tonexty ScatterpolarglFill = "tonexty" // tonexty
    ScatterpolarglFill_tonextx ScatterpolarglFill = "tonextx" // tonextx
    ScatterpolarglFill_toself ScatterpolarglFill = "toself" // toself
    ScatterpolarglFill_tonext ScatterpolarglFill = "tonext" // tonext
)

// ScatterpolarglHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterpolarglHoverlabelAlign string

const (
    ScatterpolarglHoverlabelAlign_left ScatterpolarglHoverlabelAlign = "left" // left
    ScatterpolarglHoverlabelAlign_right ScatterpolarglHoverlabelAlign = "right" // right
    ScatterpolarglHoverlabelAlign_auto ScatterpolarglHoverlabelAlign = "auto" // auto
)

// ScatterpolarglLineDash Sets the style of the lines.
type ScatterpolarglLineDash string

const (
    ScatterpolarglLineDash_solid ScatterpolarglLineDash = "solid" // solid
    ScatterpolarglLineDash_dot ScatterpolarglLineDash = "dot" // dot
    ScatterpolarglLineDash_dash ScatterpolarglLineDash = "dash" // dash
    ScatterpolarglLineDash_longdash ScatterpolarglLineDash = "longdash" // longdash
    ScatterpolarglLineDash_dashdot ScatterpolarglLineDash = "dashdot" // dashdot
    ScatterpolarglLineDash_longdashdot ScatterpolarglLineDash = "longdashdot" // longdashdot
)

// ScatterpolarglLineShape Determines the line shape. The values correspond to step-wise line shapes.
type ScatterpolarglLineShape string

const (
    ScatterpolarglLineShape_linear ScatterpolarglLineShape = "linear" // linear
    ScatterpolarglLineShape_hv ScatterpolarglLineShape = "hv" // hv
    ScatterpolarglLineShape_vh ScatterpolarglLineShape = "vh" // vh
    ScatterpolarglLineShape_hvh ScatterpolarglLineShape = "hvh" // hvh
    ScatterpolarglLineShape_vhv ScatterpolarglLineShape = "vhv" // vhv
)

// ScatterpolarglMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterpolarglMarkerColorbarExponentformat string

const (
    ScatterpolarglMarkerColorbarExponentformat_none ScatterpolarglMarkerColorbarExponentformat = "none" // none
    ScatterpolarglMarkerColorbarExponentformat_e ScatterpolarglMarkerColorbarExponentformat = "e" // e
    ScatterpolarglMarkerColorbarExponentformat_E ScatterpolarglMarkerColorbarExponentformat = "E" // E
    ScatterpolarglMarkerColorbarExponentformat_power ScatterpolarglMarkerColorbarExponentformat = "power" // power
    ScatterpolarglMarkerColorbarExponentformat_SI ScatterpolarglMarkerColorbarExponentformat = "SI" // SI
    ScatterpolarglMarkerColorbarExponentformat_B ScatterpolarglMarkerColorbarExponentformat = "B" // B
)

// ScatterpolarglMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterpolarglMarkerColorbarLenmode string

const (
    ScatterpolarglMarkerColorbarLenmode_fraction ScatterpolarglMarkerColorbarLenmode = "fraction" // fraction
    ScatterpolarglMarkerColorbarLenmode_pixels ScatterpolarglMarkerColorbarLenmode = "pixels" // pixels
)

// ScatterpolarglMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterpolarglMarkerColorbarShowexponent string

const (
    ScatterpolarglMarkerColorbarShowexponent_all ScatterpolarglMarkerColorbarShowexponent = "all" // all
    ScatterpolarglMarkerColorbarShowexponent_first ScatterpolarglMarkerColorbarShowexponent = "first" // first
    ScatterpolarglMarkerColorbarShowexponent_last ScatterpolarglMarkerColorbarShowexponent = "last" // last
    ScatterpolarglMarkerColorbarShowexponent_none ScatterpolarglMarkerColorbarShowexponent = "none" // none
)

// ScatterpolarglMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterpolarglMarkerColorbarShowtickprefix string

const (
    ScatterpolarglMarkerColorbarShowtickprefix_all ScatterpolarglMarkerColorbarShowtickprefix = "all" // all
    ScatterpolarglMarkerColorbarShowtickprefix_first ScatterpolarglMarkerColorbarShowtickprefix = "first" // first
    ScatterpolarglMarkerColorbarShowtickprefix_last ScatterpolarglMarkerColorbarShowtickprefix = "last" // last
    ScatterpolarglMarkerColorbarShowtickprefix_none ScatterpolarglMarkerColorbarShowtickprefix = "none" // none
)

// ScatterpolarglMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterpolarglMarkerColorbarShowticksuffix string

const (
    ScatterpolarglMarkerColorbarShowticksuffix_all ScatterpolarglMarkerColorbarShowticksuffix = "all" // all
    ScatterpolarglMarkerColorbarShowticksuffix_first ScatterpolarglMarkerColorbarShowticksuffix = "first" // first
    ScatterpolarglMarkerColorbarShowticksuffix_last ScatterpolarglMarkerColorbarShowticksuffix = "last" // last
    ScatterpolarglMarkerColorbarShowticksuffix_none ScatterpolarglMarkerColorbarShowticksuffix = "none" // none
)

// ScatterpolarglMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterpolarglMarkerColorbarThicknessmode string

const (
    ScatterpolarglMarkerColorbarThicknessmode_fraction ScatterpolarglMarkerColorbarThicknessmode = "fraction" // fraction
    ScatterpolarglMarkerColorbarThicknessmode_pixels ScatterpolarglMarkerColorbarThicknessmode = "pixels" // pixels
)

// ScatterpolarglMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterpolarglMarkerColorbarTickmode string

const (
    ScatterpolarglMarkerColorbarTickmode_auto ScatterpolarglMarkerColorbarTickmode = "auto" // auto
    ScatterpolarglMarkerColorbarTickmode_linear ScatterpolarglMarkerColorbarTickmode = "linear" // linear
    ScatterpolarglMarkerColorbarTickmode_array ScatterpolarglMarkerColorbarTickmode = "array" // array
)

// ScatterpolarglMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterpolarglMarkerColorbarTicks string

const (
    ScatterpolarglMarkerColorbarTicks_outside ScatterpolarglMarkerColorbarTicks = "outside" // outside
    ScatterpolarglMarkerColorbarTicks_inside ScatterpolarglMarkerColorbarTicks = "inside" // inside
)

// ScatterpolarglMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterpolarglMarkerColorbarTitleSide string

const (
    ScatterpolarglMarkerColorbarTitleSide_right ScatterpolarglMarkerColorbarTitleSide = "right" // right
    ScatterpolarglMarkerColorbarTitleSide_top ScatterpolarglMarkerColorbarTitleSide = "top" // top
    ScatterpolarglMarkerColorbarTitleSide_bottom ScatterpolarglMarkerColorbarTitleSide = "bottom" // bottom
)

// ScatterpolarglMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterpolarglMarkerColorbarXanchor string

const (
    ScatterpolarglMarkerColorbarXanchor_left ScatterpolarglMarkerColorbarXanchor = "left" // left
    ScatterpolarglMarkerColorbarXanchor_center ScatterpolarglMarkerColorbarXanchor = "center" // center
    ScatterpolarglMarkerColorbarXanchor_right ScatterpolarglMarkerColorbarXanchor = "right" // right
)

// ScatterpolarglMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterpolarglMarkerColorbarYanchor string

const (
    ScatterpolarglMarkerColorbarYanchor_top ScatterpolarglMarkerColorbarYanchor = "top" // top
    ScatterpolarglMarkerColorbarYanchor_middle ScatterpolarglMarkerColorbarYanchor = "middle" // middle
    ScatterpolarglMarkerColorbarYanchor_bottom ScatterpolarglMarkerColorbarYanchor = "bottom" // bottom
)

// ScatterpolarglMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterpolarglMarkerSizemode string

const (
    ScatterpolarglMarkerSizemode_diameter ScatterpolarglMarkerSizemode = "diameter" // diameter
    ScatterpolarglMarkerSizemode_area ScatterpolarglMarkerSizemode = "area" // area
)

// ScatterpolarglMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterpolarglMarkerSymbol string

const (
    ScatterpolarglMarkerSymbol0 ScatterpolarglMarkerSymbol = "0" // 0
    ScatterpolarglMarkerSymbol_circle ScatterpolarglMarkerSymbol = "circle" // circle
    ScatterpolarglMarkerSymbol100 ScatterpolarglMarkerSymbol = "100" // 100
    ScatterpolarglMarkerSymbol_circle_open ScatterpolarglMarkerSymbol = "circle-open" // circle-open
    ScatterpolarglMarkerSymbol200 ScatterpolarglMarkerSymbol = "200" // 200
    ScatterpolarglMarkerSymbol_circle_dot ScatterpolarglMarkerSymbol = "circle-dot" // circle-dot
    ScatterpolarglMarkerSymbol300 ScatterpolarglMarkerSymbol = "300" // 300
    ScatterpolarglMarkerSymbol_circle_open_dot ScatterpolarglMarkerSymbol = "circle-open-dot" // circle-open-dot
    ScatterpolarglMarkerSymbol1 ScatterpolarglMarkerSymbol = "1" // 1
    ScatterpolarglMarkerSymbol_square ScatterpolarglMarkerSymbol = "square" // square
    ScatterpolarglMarkerSymbol101 ScatterpolarglMarkerSymbol = "101" // 101
    ScatterpolarglMarkerSymbol_square_open ScatterpolarglMarkerSymbol = "square-open" // square-open
    ScatterpolarglMarkerSymbol201 ScatterpolarglMarkerSymbol = "201" // 201
    ScatterpolarglMarkerSymbol_square_dot ScatterpolarglMarkerSymbol = "square-dot" // square-dot
    ScatterpolarglMarkerSymbol301 ScatterpolarglMarkerSymbol = "301" // 301
    ScatterpolarglMarkerSymbol_square_open_dot ScatterpolarglMarkerSymbol = "square-open-dot" // square-open-dot
    ScatterpolarglMarkerSymbol2 ScatterpolarglMarkerSymbol = "2" // 2
    ScatterpolarglMarkerSymbol_diamond ScatterpolarglMarkerSymbol = "diamond" // diamond
    ScatterpolarglMarkerSymbol102 ScatterpolarglMarkerSymbol = "102" // 102
    ScatterpolarglMarkerSymbol_diamond_open ScatterpolarglMarkerSymbol = "diamond-open" // diamond-open
    ScatterpolarglMarkerSymbol202 ScatterpolarglMarkerSymbol = "202" // 202
    ScatterpolarglMarkerSymbol_diamond_dot ScatterpolarglMarkerSymbol = "diamond-dot" // diamond-dot
    ScatterpolarglMarkerSymbol302 ScatterpolarglMarkerSymbol = "302" // 302
    ScatterpolarglMarkerSymbol_diamond_open_dot ScatterpolarglMarkerSymbol = "diamond-open-dot" // diamond-open-dot
    ScatterpolarglMarkerSymbol3 ScatterpolarglMarkerSymbol = "3" // 3
    ScatterpolarglMarkerSymbol_cross ScatterpolarglMarkerSymbol = "cross" // cross
    ScatterpolarglMarkerSymbol103 ScatterpolarglMarkerSymbol = "103" // 103
    ScatterpolarglMarkerSymbol_cross_open ScatterpolarglMarkerSymbol = "cross-open" // cross-open
    ScatterpolarglMarkerSymbol203 ScatterpolarglMarkerSymbol = "203" // 203
    ScatterpolarglMarkerSymbol_cross_dot ScatterpolarglMarkerSymbol = "cross-dot" // cross-dot
    ScatterpolarglMarkerSymbol303 ScatterpolarglMarkerSymbol = "303" // 303
    ScatterpolarglMarkerSymbol_cross_open_dot ScatterpolarglMarkerSymbol = "cross-open-dot" // cross-open-dot
    ScatterpolarglMarkerSymbol4 ScatterpolarglMarkerSymbol = "4" // 4
    ScatterpolarglMarkerSymbol_x ScatterpolarglMarkerSymbol = "x" // x
    ScatterpolarglMarkerSymbol104 ScatterpolarglMarkerSymbol = "104" // 104
    ScatterpolarglMarkerSymbol_x_open ScatterpolarglMarkerSymbol = "x-open" // x-open
    ScatterpolarglMarkerSymbol204 ScatterpolarglMarkerSymbol = "204" // 204
    ScatterpolarglMarkerSymbol_x_dot ScatterpolarglMarkerSymbol = "x-dot" // x-dot
    ScatterpolarglMarkerSymbol304 ScatterpolarglMarkerSymbol = "304" // 304
    ScatterpolarglMarkerSymbol_x_open_dot ScatterpolarglMarkerSymbol = "x-open-dot" // x-open-dot
    ScatterpolarglMarkerSymbol5 ScatterpolarglMarkerSymbol = "5" // 5
    ScatterpolarglMarkerSymbol_triangle_up ScatterpolarglMarkerSymbol = "triangle-up" // triangle-up
    ScatterpolarglMarkerSymbol105 ScatterpolarglMarkerSymbol = "105" // 105
    ScatterpolarglMarkerSymbol_triangle_up_open ScatterpolarglMarkerSymbol = "triangle-up-open" // triangle-up-open
    ScatterpolarglMarkerSymbol205 ScatterpolarglMarkerSymbol = "205" // 205
    ScatterpolarglMarkerSymbol_triangle_up_dot ScatterpolarglMarkerSymbol = "triangle-up-dot" // triangle-up-dot
    ScatterpolarglMarkerSymbol305 ScatterpolarglMarkerSymbol = "305" // 305
    ScatterpolarglMarkerSymbol_triangle_up_open_dot ScatterpolarglMarkerSymbol = "triangle-up-open-dot" // triangle-up-open-dot
    ScatterpolarglMarkerSymbol6 ScatterpolarglMarkerSymbol = "6" // 6
    ScatterpolarglMarkerSymbol_triangle_down ScatterpolarglMarkerSymbol = "triangle-down" // triangle-down
    ScatterpolarglMarkerSymbol106 ScatterpolarglMarkerSymbol = "106" // 106
    ScatterpolarglMarkerSymbol_triangle_down_open ScatterpolarglMarkerSymbol = "triangle-down-open" // triangle-down-open
    ScatterpolarglMarkerSymbol206 ScatterpolarglMarkerSymbol = "206" // 206
    ScatterpolarglMarkerSymbol_triangle_down_dot ScatterpolarglMarkerSymbol = "triangle-down-dot" // triangle-down-dot
    ScatterpolarglMarkerSymbol306 ScatterpolarglMarkerSymbol = "306" // 306
    ScatterpolarglMarkerSymbol_triangle_down_open_dot ScatterpolarglMarkerSymbol = "triangle-down-open-dot" // triangle-down-open-dot
    ScatterpolarglMarkerSymbol7 ScatterpolarglMarkerSymbol = "7" // 7
    ScatterpolarglMarkerSymbol_triangle_left ScatterpolarglMarkerSymbol = "triangle-left" // triangle-left
    ScatterpolarglMarkerSymbol107 ScatterpolarglMarkerSymbol = "107" // 107
    ScatterpolarglMarkerSymbol_triangle_left_open ScatterpolarglMarkerSymbol = "triangle-left-open" // triangle-left-open
    ScatterpolarglMarkerSymbol207 ScatterpolarglMarkerSymbol = "207" // 207
    ScatterpolarglMarkerSymbol_triangle_left_dot ScatterpolarglMarkerSymbol = "triangle-left-dot" // triangle-left-dot
    ScatterpolarglMarkerSymbol307 ScatterpolarglMarkerSymbol = "307" // 307
    ScatterpolarglMarkerSymbol_triangle_left_open_dot ScatterpolarglMarkerSymbol = "triangle-left-open-dot" // triangle-left-open-dot
    ScatterpolarglMarkerSymbol8 ScatterpolarglMarkerSymbol = "8" // 8
    ScatterpolarglMarkerSymbol_triangle_right ScatterpolarglMarkerSymbol = "triangle-right" // triangle-right
    ScatterpolarglMarkerSymbol108 ScatterpolarglMarkerSymbol = "108" // 108
    ScatterpolarglMarkerSymbol_triangle_right_open ScatterpolarglMarkerSymbol = "triangle-right-open" // triangle-right-open
    ScatterpolarglMarkerSymbol208 ScatterpolarglMarkerSymbol = "208" // 208
    ScatterpolarglMarkerSymbol_triangle_right_dot ScatterpolarglMarkerSymbol = "triangle-right-dot" // triangle-right-dot
    ScatterpolarglMarkerSymbol308 ScatterpolarglMarkerSymbol = "308" // 308
    ScatterpolarglMarkerSymbol_triangle_right_open_dot ScatterpolarglMarkerSymbol = "triangle-right-open-dot" // triangle-right-open-dot
    ScatterpolarglMarkerSymbol9 ScatterpolarglMarkerSymbol = "9" // 9
    ScatterpolarglMarkerSymbol_triangle_ne ScatterpolarglMarkerSymbol = "triangle-ne" // triangle-ne
    ScatterpolarglMarkerSymbol109 ScatterpolarglMarkerSymbol = "109" // 109
    ScatterpolarglMarkerSymbol_triangle_ne_open ScatterpolarglMarkerSymbol = "triangle-ne-open" // triangle-ne-open
    ScatterpolarglMarkerSymbol209 ScatterpolarglMarkerSymbol = "209" // 209
    ScatterpolarglMarkerSymbol_triangle_ne_dot ScatterpolarglMarkerSymbol = "triangle-ne-dot" // triangle-ne-dot
    ScatterpolarglMarkerSymbol309 ScatterpolarglMarkerSymbol = "309" // 309
    ScatterpolarglMarkerSymbol_triangle_ne_open_dot ScatterpolarglMarkerSymbol = "triangle-ne-open-dot" // triangle-ne-open-dot
    ScatterpolarglMarkerSymbol10 ScatterpolarglMarkerSymbol = "10" // 10
    ScatterpolarglMarkerSymbol_triangle_se ScatterpolarglMarkerSymbol = "triangle-se" // triangle-se
    ScatterpolarglMarkerSymbol110 ScatterpolarglMarkerSymbol = "110" // 110
    ScatterpolarglMarkerSymbol_triangle_se_open ScatterpolarglMarkerSymbol = "triangle-se-open" // triangle-se-open
    ScatterpolarglMarkerSymbol210 ScatterpolarglMarkerSymbol = "210" // 210
    ScatterpolarglMarkerSymbol_triangle_se_dot ScatterpolarglMarkerSymbol = "triangle-se-dot" // triangle-se-dot
    ScatterpolarglMarkerSymbol310 ScatterpolarglMarkerSymbol = "310" // 310
    ScatterpolarglMarkerSymbol_triangle_se_open_dot ScatterpolarglMarkerSymbol = "triangle-se-open-dot" // triangle-se-open-dot
    ScatterpolarglMarkerSymbol11 ScatterpolarglMarkerSymbol = "11" // 11
    ScatterpolarglMarkerSymbol_triangle_sw ScatterpolarglMarkerSymbol = "triangle-sw" // triangle-sw
    ScatterpolarglMarkerSymbol111 ScatterpolarglMarkerSymbol = "111" // 111
    ScatterpolarglMarkerSymbol_triangle_sw_open ScatterpolarglMarkerSymbol = "triangle-sw-open" // triangle-sw-open
    ScatterpolarglMarkerSymbol211 ScatterpolarglMarkerSymbol = "211" // 211
    ScatterpolarglMarkerSymbol_triangle_sw_dot ScatterpolarglMarkerSymbol = "triangle-sw-dot" // triangle-sw-dot
    ScatterpolarglMarkerSymbol311 ScatterpolarglMarkerSymbol = "311" // 311
    ScatterpolarglMarkerSymbol_triangle_sw_open_dot ScatterpolarglMarkerSymbol = "triangle-sw-open-dot" // triangle-sw-open-dot
    ScatterpolarglMarkerSymbol12 ScatterpolarglMarkerSymbol = "12" // 12
    ScatterpolarglMarkerSymbol_triangle_nw ScatterpolarglMarkerSymbol = "triangle-nw" // triangle-nw
    ScatterpolarglMarkerSymbol112 ScatterpolarglMarkerSymbol = "112" // 112
    ScatterpolarglMarkerSymbol_triangle_nw_open ScatterpolarglMarkerSymbol = "triangle-nw-open" // triangle-nw-open
    ScatterpolarglMarkerSymbol212 ScatterpolarglMarkerSymbol = "212" // 212
    ScatterpolarglMarkerSymbol_triangle_nw_dot ScatterpolarglMarkerSymbol = "triangle-nw-dot" // triangle-nw-dot
    ScatterpolarglMarkerSymbol312 ScatterpolarglMarkerSymbol = "312" // 312
    ScatterpolarglMarkerSymbol_triangle_nw_open_dot ScatterpolarglMarkerSymbol = "triangle-nw-open-dot" // triangle-nw-open-dot
    ScatterpolarglMarkerSymbol13 ScatterpolarglMarkerSymbol = "13" // 13
    ScatterpolarglMarkerSymbol_pentagon ScatterpolarglMarkerSymbol = "pentagon" // pentagon
    ScatterpolarglMarkerSymbol113 ScatterpolarglMarkerSymbol = "113" // 113
    ScatterpolarglMarkerSymbol_pentagon_open ScatterpolarglMarkerSymbol = "pentagon-open" // pentagon-open
    ScatterpolarglMarkerSymbol213 ScatterpolarglMarkerSymbol = "213" // 213
    ScatterpolarglMarkerSymbol_pentagon_dot ScatterpolarglMarkerSymbol = "pentagon-dot" // pentagon-dot
    ScatterpolarglMarkerSymbol313 ScatterpolarglMarkerSymbol = "313" // 313
    ScatterpolarglMarkerSymbol_pentagon_open_dot ScatterpolarglMarkerSymbol = "pentagon-open-dot" // pentagon-open-dot
    ScatterpolarglMarkerSymbol14 ScatterpolarglMarkerSymbol = "14" // 14
    ScatterpolarglMarkerSymbol_hexagon ScatterpolarglMarkerSymbol = "hexagon" // hexagon
    ScatterpolarglMarkerSymbol114 ScatterpolarglMarkerSymbol = "114" // 114
    ScatterpolarglMarkerSymbol_hexagon_open ScatterpolarglMarkerSymbol = "hexagon-open" // hexagon-open
    ScatterpolarglMarkerSymbol214 ScatterpolarglMarkerSymbol = "214" // 214
    ScatterpolarglMarkerSymbol_hexagon_dot ScatterpolarglMarkerSymbol = "hexagon-dot" // hexagon-dot
    ScatterpolarglMarkerSymbol314 ScatterpolarglMarkerSymbol = "314" // 314
    ScatterpolarglMarkerSymbol_hexagon_open_dot ScatterpolarglMarkerSymbol = "hexagon-open-dot" // hexagon-open-dot
    ScatterpolarglMarkerSymbol15 ScatterpolarglMarkerSymbol = "15" // 15
    ScatterpolarglMarkerSymbol_hexagon2 ScatterpolarglMarkerSymbol = "hexagon2" // hexagon2
    ScatterpolarglMarkerSymbol115 ScatterpolarglMarkerSymbol = "115" // 115
    ScatterpolarglMarkerSymbol_hexagon2_open ScatterpolarglMarkerSymbol = "hexagon2-open" // hexagon2-open
    ScatterpolarglMarkerSymbol215 ScatterpolarglMarkerSymbol = "215" // 215
    ScatterpolarglMarkerSymbol_hexagon2_dot ScatterpolarglMarkerSymbol = "hexagon2-dot" // hexagon2-dot
    ScatterpolarglMarkerSymbol315 ScatterpolarglMarkerSymbol = "315" // 315
    ScatterpolarglMarkerSymbol_hexagon2_open_dot ScatterpolarglMarkerSymbol = "hexagon2-open-dot" // hexagon2-open-dot
    ScatterpolarglMarkerSymbol16 ScatterpolarglMarkerSymbol = "16" // 16
    ScatterpolarglMarkerSymbol_octagon ScatterpolarglMarkerSymbol = "octagon" // octagon
    ScatterpolarglMarkerSymbol116 ScatterpolarglMarkerSymbol = "116" // 116
    ScatterpolarglMarkerSymbol_octagon_open ScatterpolarglMarkerSymbol = "octagon-open" // octagon-open
    ScatterpolarglMarkerSymbol216 ScatterpolarglMarkerSymbol = "216" // 216
    ScatterpolarglMarkerSymbol_octagon_dot ScatterpolarglMarkerSymbol = "octagon-dot" // octagon-dot
    ScatterpolarglMarkerSymbol316 ScatterpolarglMarkerSymbol = "316" // 316
    ScatterpolarglMarkerSymbol_octagon_open_dot ScatterpolarglMarkerSymbol = "octagon-open-dot" // octagon-open-dot
    ScatterpolarglMarkerSymbol17 ScatterpolarglMarkerSymbol = "17" // 17
    ScatterpolarglMarkerSymbol_star ScatterpolarglMarkerSymbol = "star" // star
    ScatterpolarglMarkerSymbol117 ScatterpolarglMarkerSymbol = "117" // 117
    ScatterpolarglMarkerSymbol_star_open ScatterpolarglMarkerSymbol = "star-open" // star-open
    ScatterpolarglMarkerSymbol217 ScatterpolarglMarkerSymbol = "217" // 217
    ScatterpolarglMarkerSymbol_star_dot ScatterpolarglMarkerSymbol = "star-dot" // star-dot
    ScatterpolarglMarkerSymbol317 ScatterpolarglMarkerSymbol = "317" // 317
    ScatterpolarglMarkerSymbol_star_open_dot ScatterpolarglMarkerSymbol = "star-open-dot" // star-open-dot
    ScatterpolarglMarkerSymbol18 ScatterpolarglMarkerSymbol = "18" // 18
    ScatterpolarglMarkerSymbol_hexagram ScatterpolarglMarkerSymbol = "hexagram" // hexagram
    ScatterpolarglMarkerSymbol118 ScatterpolarglMarkerSymbol = "118" // 118
    ScatterpolarglMarkerSymbol_hexagram_open ScatterpolarglMarkerSymbol = "hexagram-open" // hexagram-open
    ScatterpolarglMarkerSymbol218 ScatterpolarglMarkerSymbol = "218" // 218
    ScatterpolarglMarkerSymbol_hexagram_dot ScatterpolarglMarkerSymbol = "hexagram-dot" // hexagram-dot
    ScatterpolarglMarkerSymbol318 ScatterpolarglMarkerSymbol = "318" // 318
    ScatterpolarglMarkerSymbol_hexagram_open_dot ScatterpolarglMarkerSymbol = "hexagram-open-dot" // hexagram-open-dot
    ScatterpolarglMarkerSymbol19 ScatterpolarglMarkerSymbol = "19" // 19
    ScatterpolarglMarkerSymbol_star_triangle_up ScatterpolarglMarkerSymbol = "star-triangle-up" // star-triangle-up
    ScatterpolarglMarkerSymbol119 ScatterpolarglMarkerSymbol = "119" // 119
    ScatterpolarglMarkerSymbol_star_triangle_up_open ScatterpolarglMarkerSymbol = "star-triangle-up-open" // star-triangle-up-open
    ScatterpolarglMarkerSymbol219 ScatterpolarglMarkerSymbol = "219" // 219
    ScatterpolarglMarkerSymbol_star_triangle_up_dot ScatterpolarglMarkerSymbol = "star-triangle-up-dot" // star-triangle-up-dot
    ScatterpolarglMarkerSymbol319 ScatterpolarglMarkerSymbol = "319" // 319
    ScatterpolarglMarkerSymbol_star_triangle_up_open_dot ScatterpolarglMarkerSymbol = "star-triangle-up-open-dot" // star-triangle-up-open-dot
    ScatterpolarglMarkerSymbol20 ScatterpolarglMarkerSymbol = "20" // 20
    ScatterpolarglMarkerSymbol_star_triangle_down ScatterpolarglMarkerSymbol = "star-triangle-down" // star-triangle-down
    ScatterpolarglMarkerSymbol120 ScatterpolarglMarkerSymbol = "120" // 120
    ScatterpolarglMarkerSymbol_star_triangle_down_open ScatterpolarglMarkerSymbol = "star-triangle-down-open" // star-triangle-down-open
    ScatterpolarglMarkerSymbol220 ScatterpolarglMarkerSymbol = "220" // 220
    ScatterpolarglMarkerSymbol_star_triangle_down_dot ScatterpolarglMarkerSymbol = "star-triangle-down-dot" // star-triangle-down-dot
    ScatterpolarglMarkerSymbol320 ScatterpolarglMarkerSymbol = "320" // 320
    ScatterpolarglMarkerSymbol_star_triangle_down_open_dot ScatterpolarglMarkerSymbol = "star-triangle-down-open-dot" // star-triangle-down-open-dot
    ScatterpolarglMarkerSymbol21 ScatterpolarglMarkerSymbol = "21" // 21
    ScatterpolarglMarkerSymbol_star_square ScatterpolarglMarkerSymbol = "star-square" // star-square
    ScatterpolarglMarkerSymbol121 ScatterpolarglMarkerSymbol = "121" // 121
    ScatterpolarglMarkerSymbol_star_square_open ScatterpolarglMarkerSymbol = "star-square-open" // star-square-open
    ScatterpolarglMarkerSymbol221 ScatterpolarglMarkerSymbol = "221" // 221
    ScatterpolarglMarkerSymbol_star_square_dot ScatterpolarglMarkerSymbol = "star-square-dot" // star-square-dot
    ScatterpolarglMarkerSymbol321 ScatterpolarglMarkerSymbol = "321" // 321
    ScatterpolarglMarkerSymbol_star_square_open_dot ScatterpolarglMarkerSymbol = "star-square-open-dot" // star-square-open-dot
    ScatterpolarglMarkerSymbol22 ScatterpolarglMarkerSymbol = "22" // 22
    ScatterpolarglMarkerSymbol_star_diamond ScatterpolarglMarkerSymbol = "star-diamond" // star-diamond
    ScatterpolarglMarkerSymbol122 ScatterpolarglMarkerSymbol = "122" // 122
    ScatterpolarglMarkerSymbol_star_diamond_open ScatterpolarglMarkerSymbol = "star-diamond-open" // star-diamond-open
    ScatterpolarglMarkerSymbol222 ScatterpolarglMarkerSymbol = "222" // 222
    ScatterpolarglMarkerSymbol_star_diamond_dot ScatterpolarglMarkerSymbol = "star-diamond-dot" // star-diamond-dot
    ScatterpolarglMarkerSymbol322 ScatterpolarglMarkerSymbol = "322" // 322
    ScatterpolarglMarkerSymbol_star_diamond_open_dot ScatterpolarglMarkerSymbol = "star-diamond-open-dot" // star-diamond-open-dot
    ScatterpolarglMarkerSymbol23 ScatterpolarglMarkerSymbol = "23" // 23
    ScatterpolarglMarkerSymbol_diamond_tall ScatterpolarglMarkerSymbol = "diamond-tall" // diamond-tall
    ScatterpolarglMarkerSymbol123 ScatterpolarglMarkerSymbol = "123" // 123
    ScatterpolarglMarkerSymbol_diamond_tall_open ScatterpolarglMarkerSymbol = "diamond-tall-open" // diamond-tall-open
    ScatterpolarglMarkerSymbol223 ScatterpolarglMarkerSymbol = "223" // 223
    ScatterpolarglMarkerSymbol_diamond_tall_dot ScatterpolarglMarkerSymbol = "diamond-tall-dot" // diamond-tall-dot
    ScatterpolarglMarkerSymbol323 ScatterpolarglMarkerSymbol = "323" // 323
    ScatterpolarglMarkerSymbol_diamond_tall_open_dot ScatterpolarglMarkerSymbol = "diamond-tall-open-dot" // diamond-tall-open-dot
    ScatterpolarglMarkerSymbol24 ScatterpolarglMarkerSymbol = "24" // 24
    ScatterpolarglMarkerSymbol_diamond_wide ScatterpolarglMarkerSymbol = "diamond-wide" // diamond-wide
    ScatterpolarglMarkerSymbol124 ScatterpolarglMarkerSymbol = "124" // 124
    ScatterpolarglMarkerSymbol_diamond_wide_open ScatterpolarglMarkerSymbol = "diamond-wide-open" // diamond-wide-open
    ScatterpolarglMarkerSymbol224 ScatterpolarglMarkerSymbol = "224" // 224
    ScatterpolarglMarkerSymbol_diamond_wide_dot ScatterpolarglMarkerSymbol = "diamond-wide-dot" // diamond-wide-dot
    ScatterpolarglMarkerSymbol324 ScatterpolarglMarkerSymbol = "324" // 324
    ScatterpolarglMarkerSymbol_diamond_wide_open_dot ScatterpolarglMarkerSymbol = "diamond-wide-open-dot" // diamond-wide-open-dot
    ScatterpolarglMarkerSymbol25 ScatterpolarglMarkerSymbol = "25" // 25
    ScatterpolarglMarkerSymbol_hourglass ScatterpolarglMarkerSymbol = "hourglass" // hourglass
    ScatterpolarglMarkerSymbol125 ScatterpolarglMarkerSymbol = "125" // 125
    ScatterpolarglMarkerSymbol_hourglass_open ScatterpolarglMarkerSymbol = "hourglass-open" // hourglass-open
    ScatterpolarglMarkerSymbol26 ScatterpolarglMarkerSymbol = "26" // 26
    ScatterpolarglMarkerSymbol_bowtie ScatterpolarglMarkerSymbol = "bowtie" // bowtie
    ScatterpolarglMarkerSymbol126 ScatterpolarglMarkerSymbol = "126" // 126
    ScatterpolarglMarkerSymbol_bowtie_open ScatterpolarglMarkerSymbol = "bowtie-open" // bowtie-open
    ScatterpolarglMarkerSymbol27 ScatterpolarglMarkerSymbol = "27" // 27
    ScatterpolarglMarkerSymbol_circle_cross ScatterpolarglMarkerSymbol = "circle-cross" // circle-cross
    ScatterpolarglMarkerSymbol127 ScatterpolarglMarkerSymbol = "127" // 127
    ScatterpolarglMarkerSymbol_circle_cross_open ScatterpolarglMarkerSymbol = "circle-cross-open" // circle-cross-open
    ScatterpolarglMarkerSymbol28 ScatterpolarglMarkerSymbol = "28" // 28
    ScatterpolarglMarkerSymbol_circle_x ScatterpolarglMarkerSymbol = "circle-x" // circle-x
    ScatterpolarglMarkerSymbol128 ScatterpolarglMarkerSymbol = "128" // 128
    ScatterpolarglMarkerSymbol_circle_x_open ScatterpolarglMarkerSymbol = "circle-x-open" // circle-x-open
    ScatterpolarglMarkerSymbol29 ScatterpolarglMarkerSymbol = "29" // 29
    ScatterpolarglMarkerSymbol_square_cross ScatterpolarglMarkerSymbol = "square-cross" // square-cross
    ScatterpolarglMarkerSymbol129 ScatterpolarglMarkerSymbol = "129" // 129
    ScatterpolarglMarkerSymbol_square_cross_open ScatterpolarglMarkerSymbol = "square-cross-open" // square-cross-open
    ScatterpolarglMarkerSymbol30 ScatterpolarglMarkerSymbol = "30" // 30
    ScatterpolarglMarkerSymbol_square_x ScatterpolarglMarkerSymbol = "square-x" // square-x
    ScatterpolarglMarkerSymbol130 ScatterpolarglMarkerSymbol = "130" // 130
    ScatterpolarglMarkerSymbol_square_x_open ScatterpolarglMarkerSymbol = "square-x-open" // square-x-open
    ScatterpolarglMarkerSymbol31 ScatterpolarglMarkerSymbol = "31" // 31
    ScatterpolarglMarkerSymbol_diamond_cross ScatterpolarglMarkerSymbol = "diamond-cross" // diamond-cross
    ScatterpolarglMarkerSymbol131 ScatterpolarglMarkerSymbol = "131" // 131
    ScatterpolarglMarkerSymbol_diamond_cross_open ScatterpolarglMarkerSymbol = "diamond-cross-open" // diamond-cross-open
    ScatterpolarglMarkerSymbol32 ScatterpolarglMarkerSymbol = "32" // 32
    ScatterpolarglMarkerSymbol_diamond_x ScatterpolarglMarkerSymbol = "diamond-x" // diamond-x
    ScatterpolarglMarkerSymbol132 ScatterpolarglMarkerSymbol = "132" // 132
    ScatterpolarglMarkerSymbol_diamond_x_open ScatterpolarglMarkerSymbol = "diamond-x-open" // diamond-x-open
    ScatterpolarglMarkerSymbol33 ScatterpolarglMarkerSymbol = "33" // 33
    ScatterpolarglMarkerSymbol_cross_thin ScatterpolarglMarkerSymbol = "cross-thin" // cross-thin
    ScatterpolarglMarkerSymbol133 ScatterpolarglMarkerSymbol = "133" // 133
    ScatterpolarglMarkerSymbol_cross_thin_open ScatterpolarglMarkerSymbol = "cross-thin-open" // cross-thin-open
    ScatterpolarglMarkerSymbol34 ScatterpolarglMarkerSymbol = "34" // 34
    ScatterpolarglMarkerSymbol_x_thin ScatterpolarglMarkerSymbol = "x-thin" // x-thin
    ScatterpolarglMarkerSymbol134 ScatterpolarglMarkerSymbol = "134" // 134
    ScatterpolarglMarkerSymbol_x_thin_open ScatterpolarglMarkerSymbol = "x-thin-open" // x-thin-open
    ScatterpolarglMarkerSymbol35 ScatterpolarglMarkerSymbol = "35" // 35
    ScatterpolarglMarkerSymbol_asterisk ScatterpolarglMarkerSymbol = "asterisk" // asterisk
    ScatterpolarglMarkerSymbol135 ScatterpolarglMarkerSymbol = "135" // 135
    ScatterpolarglMarkerSymbol_asterisk_open ScatterpolarglMarkerSymbol = "asterisk-open" // asterisk-open
    ScatterpolarglMarkerSymbol36 ScatterpolarglMarkerSymbol = "36" // 36
    ScatterpolarglMarkerSymbol_hash ScatterpolarglMarkerSymbol = "hash" // hash
    ScatterpolarglMarkerSymbol136 ScatterpolarglMarkerSymbol = "136" // 136
    ScatterpolarglMarkerSymbol_hash_open ScatterpolarglMarkerSymbol = "hash-open" // hash-open
    ScatterpolarglMarkerSymbol236 ScatterpolarglMarkerSymbol = "236" // 236
    ScatterpolarglMarkerSymbol_hash_dot ScatterpolarglMarkerSymbol = "hash-dot" // hash-dot
    ScatterpolarglMarkerSymbol336 ScatterpolarglMarkerSymbol = "336" // 336
    ScatterpolarglMarkerSymbol_hash_open_dot ScatterpolarglMarkerSymbol = "hash-open-dot" // hash-open-dot
    ScatterpolarglMarkerSymbol37 ScatterpolarglMarkerSymbol = "37" // 37
    ScatterpolarglMarkerSymbol_y_up ScatterpolarglMarkerSymbol = "y-up" // y-up
    ScatterpolarglMarkerSymbol137 ScatterpolarglMarkerSymbol = "137" // 137
    ScatterpolarglMarkerSymbol_y_up_open ScatterpolarglMarkerSymbol = "y-up-open" // y-up-open
    ScatterpolarglMarkerSymbol38 ScatterpolarglMarkerSymbol = "38" // 38
    ScatterpolarglMarkerSymbol_y_down ScatterpolarglMarkerSymbol = "y-down" // y-down
    ScatterpolarglMarkerSymbol138 ScatterpolarglMarkerSymbol = "138" // 138
    ScatterpolarglMarkerSymbol_y_down_open ScatterpolarglMarkerSymbol = "y-down-open" // y-down-open
    ScatterpolarglMarkerSymbol39 ScatterpolarglMarkerSymbol = "39" // 39
    ScatterpolarglMarkerSymbol_y_left ScatterpolarglMarkerSymbol = "y-left" // y-left
    ScatterpolarglMarkerSymbol139 ScatterpolarglMarkerSymbol = "139" // 139
    ScatterpolarglMarkerSymbol_y_left_open ScatterpolarglMarkerSymbol = "y-left-open" // y-left-open
    ScatterpolarglMarkerSymbol40 ScatterpolarglMarkerSymbol = "40" // 40
    ScatterpolarglMarkerSymbol_y_right ScatterpolarglMarkerSymbol = "y-right" // y-right
    ScatterpolarglMarkerSymbol140 ScatterpolarglMarkerSymbol = "140" // 140
    ScatterpolarglMarkerSymbol_y_right_open ScatterpolarglMarkerSymbol = "y-right-open" // y-right-open
    ScatterpolarglMarkerSymbol41 ScatterpolarglMarkerSymbol = "41" // 41
    ScatterpolarglMarkerSymbol_line_ew ScatterpolarglMarkerSymbol = "line-ew" // line-ew
    ScatterpolarglMarkerSymbol141 ScatterpolarglMarkerSymbol = "141" // 141
    ScatterpolarglMarkerSymbol_line_ew_open ScatterpolarglMarkerSymbol = "line-ew-open" // line-ew-open
    ScatterpolarglMarkerSymbol42 ScatterpolarglMarkerSymbol = "42" // 42
    ScatterpolarglMarkerSymbol_line_ns ScatterpolarglMarkerSymbol = "line-ns" // line-ns
    ScatterpolarglMarkerSymbol142 ScatterpolarglMarkerSymbol = "142" // 142
    ScatterpolarglMarkerSymbol_line_ns_open ScatterpolarglMarkerSymbol = "line-ns-open" // line-ns-open
    ScatterpolarglMarkerSymbol43 ScatterpolarglMarkerSymbol = "43" // 43
    ScatterpolarglMarkerSymbol_line_ne ScatterpolarglMarkerSymbol = "line-ne" // line-ne
    ScatterpolarglMarkerSymbol143 ScatterpolarglMarkerSymbol = "143" // 143
    ScatterpolarglMarkerSymbol_line_ne_open ScatterpolarglMarkerSymbol = "line-ne-open" // line-ne-open
    ScatterpolarglMarkerSymbol44 ScatterpolarglMarkerSymbol = "44" // 44
    ScatterpolarglMarkerSymbol_line_nw ScatterpolarglMarkerSymbol = "line-nw" // line-nw
    ScatterpolarglMarkerSymbol144 ScatterpolarglMarkerSymbol = "144" // 144
    ScatterpolarglMarkerSymbol_line_nw_open ScatterpolarglMarkerSymbol = "line-nw-open" // line-nw-open
)

// ScatterpolarglTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterpolarglTextposition string

const (
    ScatterpolarglTextposition_topleft ScatterpolarglTextposition = "top left" // top left
    ScatterpolarglTextposition_topcenter ScatterpolarglTextposition = "top center" // top center
    ScatterpolarglTextposition_topright ScatterpolarglTextposition = "top right" // top right
    ScatterpolarglTextposition_middleleft ScatterpolarglTextposition = "middle left" // middle left
    ScatterpolarglTextposition_middlecenter ScatterpolarglTextposition = "middle center" // middle center
    ScatterpolarglTextposition_middleright ScatterpolarglTextposition = "middle right" // middle right
    ScatterpolarglTextposition_bottomleft ScatterpolarglTextposition = "bottom left" // bottom left
    ScatterpolarglTextposition_bottomcenter ScatterpolarglTextposition = "bottom center" // bottom center
    ScatterpolarglTextposition_bottomright ScatterpolarglTextposition = "bottom right" // bottom right
)

// ScatterpolarglThetaunit Sets the unit of input *theta* values. Has an effect only when on *linear* angular axes.
type ScatterpolarglThetaunit string

const (
    ScatterpolarglThetaunit_radians ScatterpolarglThetaunit = "radians" // radians
    ScatterpolarglThetaunit_degrees ScatterpolarglThetaunit = "degrees" // degrees
    ScatterpolarglThetaunit_gradians ScatterpolarglThetaunit = "gradians" // gradians
)

// ScatterpolarglVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterpolarglVisible interface{}

const (
    ScatterpolarglVisibleTrue bool = true
    ScatterpolarglVisibleFalse bool = false
    ScatterpolarglVisibleLegendonly string = "legendonly"
)

// ScatterternaryFill Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. scatterternary has a subset of the options available to scatter. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other.
type ScatterternaryFill string

const (
    ScatterternaryFill_none ScatterternaryFill = "none" // none
    ScatterternaryFill_toself ScatterternaryFill = "toself" // toself
    ScatterternaryFill_tonext ScatterternaryFill = "tonext" // tonext
)

// ScatterternaryHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ScatterternaryHoverlabelAlign string

const (
    ScatterternaryHoverlabelAlign_left ScatterternaryHoverlabelAlign = "left" // left
    ScatterternaryHoverlabelAlign_right ScatterternaryHoverlabelAlign = "right" // right
    ScatterternaryHoverlabelAlign_auto ScatterternaryHoverlabelAlign = "auto" // auto
)

// ScatterternaryLineShape Determines the line shape. With *spline* the lines are drawn using spline interpolation. The other available values correspond to step-wise line shapes.
type ScatterternaryLineShape string

const (
    ScatterternaryLineShape_linear ScatterternaryLineShape = "linear" // linear
    ScatterternaryLineShape_spline ScatterternaryLineShape = "spline" // spline
)

// ScatterternaryMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type ScatterternaryMarkerColorbarExponentformat string

const (
    ScatterternaryMarkerColorbarExponentformat_none ScatterternaryMarkerColorbarExponentformat = "none" // none
    ScatterternaryMarkerColorbarExponentformat_e ScatterternaryMarkerColorbarExponentformat = "e" // e
    ScatterternaryMarkerColorbarExponentformat_E ScatterternaryMarkerColorbarExponentformat = "E" // E
    ScatterternaryMarkerColorbarExponentformat_power ScatterternaryMarkerColorbarExponentformat = "power" // power
    ScatterternaryMarkerColorbarExponentformat_SI ScatterternaryMarkerColorbarExponentformat = "SI" // SI
    ScatterternaryMarkerColorbarExponentformat_B ScatterternaryMarkerColorbarExponentformat = "B" // B
)

// ScatterternaryMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type ScatterternaryMarkerColorbarLenmode string

const (
    ScatterternaryMarkerColorbarLenmode_fraction ScatterternaryMarkerColorbarLenmode = "fraction" // fraction
    ScatterternaryMarkerColorbarLenmode_pixels ScatterternaryMarkerColorbarLenmode = "pixels" // pixels
)

// ScatterternaryMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type ScatterternaryMarkerColorbarShowexponent string

const (
    ScatterternaryMarkerColorbarShowexponent_all ScatterternaryMarkerColorbarShowexponent = "all" // all
    ScatterternaryMarkerColorbarShowexponent_first ScatterternaryMarkerColorbarShowexponent = "first" // first
    ScatterternaryMarkerColorbarShowexponent_last ScatterternaryMarkerColorbarShowexponent = "last" // last
    ScatterternaryMarkerColorbarShowexponent_none ScatterternaryMarkerColorbarShowexponent = "none" // none
)

// ScatterternaryMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type ScatterternaryMarkerColorbarShowtickprefix string

const (
    ScatterternaryMarkerColorbarShowtickprefix_all ScatterternaryMarkerColorbarShowtickprefix = "all" // all
    ScatterternaryMarkerColorbarShowtickprefix_first ScatterternaryMarkerColorbarShowtickprefix = "first" // first
    ScatterternaryMarkerColorbarShowtickprefix_last ScatterternaryMarkerColorbarShowtickprefix = "last" // last
    ScatterternaryMarkerColorbarShowtickprefix_none ScatterternaryMarkerColorbarShowtickprefix = "none" // none
)

// ScatterternaryMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type ScatterternaryMarkerColorbarShowticksuffix string

const (
    ScatterternaryMarkerColorbarShowticksuffix_all ScatterternaryMarkerColorbarShowticksuffix = "all" // all
    ScatterternaryMarkerColorbarShowticksuffix_first ScatterternaryMarkerColorbarShowticksuffix = "first" // first
    ScatterternaryMarkerColorbarShowticksuffix_last ScatterternaryMarkerColorbarShowticksuffix = "last" // last
    ScatterternaryMarkerColorbarShowticksuffix_none ScatterternaryMarkerColorbarShowticksuffix = "none" // none
)

// ScatterternaryMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type ScatterternaryMarkerColorbarThicknessmode string

const (
    ScatterternaryMarkerColorbarThicknessmode_fraction ScatterternaryMarkerColorbarThicknessmode = "fraction" // fraction
    ScatterternaryMarkerColorbarThicknessmode_pixels ScatterternaryMarkerColorbarThicknessmode = "pixels" // pixels
)

// ScatterternaryMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type ScatterternaryMarkerColorbarTickmode string

const (
    ScatterternaryMarkerColorbarTickmode_auto ScatterternaryMarkerColorbarTickmode = "auto" // auto
    ScatterternaryMarkerColorbarTickmode_linear ScatterternaryMarkerColorbarTickmode = "linear" // linear
    ScatterternaryMarkerColorbarTickmode_array ScatterternaryMarkerColorbarTickmode = "array" // array
)

// ScatterternaryMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type ScatterternaryMarkerColorbarTicks string

const (
    ScatterternaryMarkerColorbarTicks_outside ScatterternaryMarkerColorbarTicks = "outside" // outside
    ScatterternaryMarkerColorbarTicks_inside ScatterternaryMarkerColorbarTicks = "inside" // inside
)

// ScatterternaryMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type ScatterternaryMarkerColorbarTitleSide string

const (
    ScatterternaryMarkerColorbarTitleSide_right ScatterternaryMarkerColorbarTitleSide = "right" // right
    ScatterternaryMarkerColorbarTitleSide_top ScatterternaryMarkerColorbarTitleSide = "top" // top
    ScatterternaryMarkerColorbarTitleSide_bottom ScatterternaryMarkerColorbarTitleSide = "bottom" // bottom
)

// ScatterternaryMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type ScatterternaryMarkerColorbarXanchor string

const (
    ScatterternaryMarkerColorbarXanchor_left ScatterternaryMarkerColorbarXanchor = "left" // left
    ScatterternaryMarkerColorbarXanchor_center ScatterternaryMarkerColorbarXanchor = "center" // center
    ScatterternaryMarkerColorbarXanchor_right ScatterternaryMarkerColorbarXanchor = "right" // right
)

// ScatterternaryMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type ScatterternaryMarkerColorbarYanchor string

const (
    ScatterternaryMarkerColorbarYanchor_top ScatterternaryMarkerColorbarYanchor = "top" // top
    ScatterternaryMarkerColorbarYanchor_middle ScatterternaryMarkerColorbarYanchor = "middle" // middle
    ScatterternaryMarkerColorbarYanchor_bottom ScatterternaryMarkerColorbarYanchor = "bottom" // bottom
)

// ScatterternaryMarkerGradientType Sets the type of gradient used to fill the markers
type ScatterternaryMarkerGradientType string

const (
    ScatterternaryMarkerGradientType_radial ScatterternaryMarkerGradientType = "radial" // radial
    ScatterternaryMarkerGradientType_horizontal ScatterternaryMarkerGradientType = "horizontal" // horizontal
    ScatterternaryMarkerGradientType_vertical ScatterternaryMarkerGradientType = "vertical" // vertical
    ScatterternaryMarkerGradientType_none ScatterternaryMarkerGradientType = "none" // none
)

// ScatterternaryMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type ScatterternaryMarkerSizemode string

const (
    ScatterternaryMarkerSizemode_diameter ScatterternaryMarkerSizemode = "diameter" // diameter
    ScatterternaryMarkerSizemode_area ScatterternaryMarkerSizemode = "area" // area
)

// ScatterternaryMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ScatterternaryMarkerSymbol string

const (
    ScatterternaryMarkerSymbol0 ScatterternaryMarkerSymbol = "0" // 0
    ScatterternaryMarkerSymbol_circle ScatterternaryMarkerSymbol = "circle" // circle
    ScatterternaryMarkerSymbol100 ScatterternaryMarkerSymbol = "100" // 100
    ScatterternaryMarkerSymbol_circle_open ScatterternaryMarkerSymbol = "circle-open" // circle-open
    ScatterternaryMarkerSymbol200 ScatterternaryMarkerSymbol = "200" // 200
    ScatterternaryMarkerSymbol_circle_dot ScatterternaryMarkerSymbol = "circle-dot" // circle-dot
    ScatterternaryMarkerSymbol300 ScatterternaryMarkerSymbol = "300" // 300
    ScatterternaryMarkerSymbol_circle_open_dot ScatterternaryMarkerSymbol = "circle-open-dot" // circle-open-dot
    ScatterternaryMarkerSymbol1 ScatterternaryMarkerSymbol = "1" // 1
    ScatterternaryMarkerSymbol_square ScatterternaryMarkerSymbol = "square" // square
    ScatterternaryMarkerSymbol101 ScatterternaryMarkerSymbol = "101" // 101
    ScatterternaryMarkerSymbol_square_open ScatterternaryMarkerSymbol = "square-open" // square-open
    ScatterternaryMarkerSymbol201 ScatterternaryMarkerSymbol = "201" // 201
    ScatterternaryMarkerSymbol_square_dot ScatterternaryMarkerSymbol = "square-dot" // square-dot
    ScatterternaryMarkerSymbol301 ScatterternaryMarkerSymbol = "301" // 301
    ScatterternaryMarkerSymbol_square_open_dot ScatterternaryMarkerSymbol = "square-open-dot" // square-open-dot
    ScatterternaryMarkerSymbol2 ScatterternaryMarkerSymbol = "2" // 2
    ScatterternaryMarkerSymbol_diamond ScatterternaryMarkerSymbol = "diamond" // diamond
    ScatterternaryMarkerSymbol102 ScatterternaryMarkerSymbol = "102" // 102
    ScatterternaryMarkerSymbol_diamond_open ScatterternaryMarkerSymbol = "diamond-open" // diamond-open
    ScatterternaryMarkerSymbol202 ScatterternaryMarkerSymbol = "202" // 202
    ScatterternaryMarkerSymbol_diamond_dot ScatterternaryMarkerSymbol = "diamond-dot" // diamond-dot
    ScatterternaryMarkerSymbol302 ScatterternaryMarkerSymbol = "302" // 302
    ScatterternaryMarkerSymbol_diamond_open_dot ScatterternaryMarkerSymbol = "diamond-open-dot" // diamond-open-dot
    ScatterternaryMarkerSymbol3 ScatterternaryMarkerSymbol = "3" // 3
    ScatterternaryMarkerSymbol_cross ScatterternaryMarkerSymbol = "cross" // cross
    ScatterternaryMarkerSymbol103 ScatterternaryMarkerSymbol = "103" // 103
    ScatterternaryMarkerSymbol_cross_open ScatterternaryMarkerSymbol = "cross-open" // cross-open
    ScatterternaryMarkerSymbol203 ScatterternaryMarkerSymbol = "203" // 203
    ScatterternaryMarkerSymbol_cross_dot ScatterternaryMarkerSymbol = "cross-dot" // cross-dot
    ScatterternaryMarkerSymbol303 ScatterternaryMarkerSymbol = "303" // 303
    ScatterternaryMarkerSymbol_cross_open_dot ScatterternaryMarkerSymbol = "cross-open-dot" // cross-open-dot
    ScatterternaryMarkerSymbol4 ScatterternaryMarkerSymbol = "4" // 4
    ScatterternaryMarkerSymbol_x ScatterternaryMarkerSymbol = "x" // x
    ScatterternaryMarkerSymbol104 ScatterternaryMarkerSymbol = "104" // 104
    ScatterternaryMarkerSymbol_x_open ScatterternaryMarkerSymbol = "x-open" // x-open
    ScatterternaryMarkerSymbol204 ScatterternaryMarkerSymbol = "204" // 204
    ScatterternaryMarkerSymbol_x_dot ScatterternaryMarkerSymbol = "x-dot" // x-dot
    ScatterternaryMarkerSymbol304 ScatterternaryMarkerSymbol = "304" // 304
    ScatterternaryMarkerSymbol_x_open_dot ScatterternaryMarkerSymbol = "x-open-dot" // x-open-dot
    ScatterternaryMarkerSymbol5 ScatterternaryMarkerSymbol = "5" // 5
    ScatterternaryMarkerSymbol_triangle_up ScatterternaryMarkerSymbol = "triangle-up" // triangle-up
    ScatterternaryMarkerSymbol105 ScatterternaryMarkerSymbol = "105" // 105
    ScatterternaryMarkerSymbol_triangle_up_open ScatterternaryMarkerSymbol = "triangle-up-open" // triangle-up-open
    ScatterternaryMarkerSymbol205 ScatterternaryMarkerSymbol = "205" // 205
    ScatterternaryMarkerSymbol_triangle_up_dot ScatterternaryMarkerSymbol = "triangle-up-dot" // triangle-up-dot
    ScatterternaryMarkerSymbol305 ScatterternaryMarkerSymbol = "305" // 305
    ScatterternaryMarkerSymbol_triangle_up_open_dot ScatterternaryMarkerSymbol = "triangle-up-open-dot" // triangle-up-open-dot
    ScatterternaryMarkerSymbol6 ScatterternaryMarkerSymbol = "6" // 6
    ScatterternaryMarkerSymbol_triangle_down ScatterternaryMarkerSymbol = "triangle-down" // triangle-down
    ScatterternaryMarkerSymbol106 ScatterternaryMarkerSymbol = "106" // 106
    ScatterternaryMarkerSymbol_triangle_down_open ScatterternaryMarkerSymbol = "triangle-down-open" // triangle-down-open
    ScatterternaryMarkerSymbol206 ScatterternaryMarkerSymbol = "206" // 206
    ScatterternaryMarkerSymbol_triangle_down_dot ScatterternaryMarkerSymbol = "triangle-down-dot" // triangle-down-dot
    ScatterternaryMarkerSymbol306 ScatterternaryMarkerSymbol = "306" // 306
    ScatterternaryMarkerSymbol_triangle_down_open_dot ScatterternaryMarkerSymbol = "triangle-down-open-dot" // triangle-down-open-dot
    ScatterternaryMarkerSymbol7 ScatterternaryMarkerSymbol = "7" // 7
    ScatterternaryMarkerSymbol_triangle_left ScatterternaryMarkerSymbol = "triangle-left" // triangle-left
    ScatterternaryMarkerSymbol107 ScatterternaryMarkerSymbol = "107" // 107
    ScatterternaryMarkerSymbol_triangle_left_open ScatterternaryMarkerSymbol = "triangle-left-open" // triangle-left-open
    ScatterternaryMarkerSymbol207 ScatterternaryMarkerSymbol = "207" // 207
    ScatterternaryMarkerSymbol_triangle_left_dot ScatterternaryMarkerSymbol = "triangle-left-dot" // triangle-left-dot
    ScatterternaryMarkerSymbol307 ScatterternaryMarkerSymbol = "307" // 307
    ScatterternaryMarkerSymbol_triangle_left_open_dot ScatterternaryMarkerSymbol = "triangle-left-open-dot" // triangle-left-open-dot
    ScatterternaryMarkerSymbol8 ScatterternaryMarkerSymbol = "8" // 8
    ScatterternaryMarkerSymbol_triangle_right ScatterternaryMarkerSymbol = "triangle-right" // triangle-right
    ScatterternaryMarkerSymbol108 ScatterternaryMarkerSymbol = "108" // 108
    ScatterternaryMarkerSymbol_triangle_right_open ScatterternaryMarkerSymbol = "triangle-right-open" // triangle-right-open
    ScatterternaryMarkerSymbol208 ScatterternaryMarkerSymbol = "208" // 208
    ScatterternaryMarkerSymbol_triangle_right_dot ScatterternaryMarkerSymbol = "triangle-right-dot" // triangle-right-dot
    ScatterternaryMarkerSymbol308 ScatterternaryMarkerSymbol = "308" // 308
    ScatterternaryMarkerSymbol_triangle_right_open_dot ScatterternaryMarkerSymbol = "triangle-right-open-dot" // triangle-right-open-dot
    ScatterternaryMarkerSymbol9 ScatterternaryMarkerSymbol = "9" // 9
    ScatterternaryMarkerSymbol_triangle_ne ScatterternaryMarkerSymbol = "triangle-ne" // triangle-ne
    ScatterternaryMarkerSymbol109 ScatterternaryMarkerSymbol = "109" // 109
    ScatterternaryMarkerSymbol_triangle_ne_open ScatterternaryMarkerSymbol = "triangle-ne-open" // triangle-ne-open
    ScatterternaryMarkerSymbol209 ScatterternaryMarkerSymbol = "209" // 209
    ScatterternaryMarkerSymbol_triangle_ne_dot ScatterternaryMarkerSymbol = "triangle-ne-dot" // triangle-ne-dot
    ScatterternaryMarkerSymbol309 ScatterternaryMarkerSymbol = "309" // 309
    ScatterternaryMarkerSymbol_triangle_ne_open_dot ScatterternaryMarkerSymbol = "triangle-ne-open-dot" // triangle-ne-open-dot
    ScatterternaryMarkerSymbol10 ScatterternaryMarkerSymbol = "10" // 10
    ScatterternaryMarkerSymbol_triangle_se ScatterternaryMarkerSymbol = "triangle-se" // triangle-se
    ScatterternaryMarkerSymbol110 ScatterternaryMarkerSymbol = "110" // 110
    ScatterternaryMarkerSymbol_triangle_se_open ScatterternaryMarkerSymbol = "triangle-se-open" // triangle-se-open
    ScatterternaryMarkerSymbol210 ScatterternaryMarkerSymbol = "210" // 210
    ScatterternaryMarkerSymbol_triangle_se_dot ScatterternaryMarkerSymbol = "triangle-se-dot" // triangle-se-dot
    ScatterternaryMarkerSymbol310 ScatterternaryMarkerSymbol = "310" // 310
    ScatterternaryMarkerSymbol_triangle_se_open_dot ScatterternaryMarkerSymbol = "triangle-se-open-dot" // triangle-se-open-dot
    ScatterternaryMarkerSymbol11 ScatterternaryMarkerSymbol = "11" // 11
    ScatterternaryMarkerSymbol_triangle_sw ScatterternaryMarkerSymbol = "triangle-sw" // triangle-sw
    ScatterternaryMarkerSymbol111 ScatterternaryMarkerSymbol = "111" // 111
    ScatterternaryMarkerSymbol_triangle_sw_open ScatterternaryMarkerSymbol = "triangle-sw-open" // triangle-sw-open
    ScatterternaryMarkerSymbol211 ScatterternaryMarkerSymbol = "211" // 211
    ScatterternaryMarkerSymbol_triangle_sw_dot ScatterternaryMarkerSymbol = "triangle-sw-dot" // triangle-sw-dot
    ScatterternaryMarkerSymbol311 ScatterternaryMarkerSymbol = "311" // 311
    ScatterternaryMarkerSymbol_triangle_sw_open_dot ScatterternaryMarkerSymbol = "triangle-sw-open-dot" // triangle-sw-open-dot
    ScatterternaryMarkerSymbol12 ScatterternaryMarkerSymbol = "12" // 12
    ScatterternaryMarkerSymbol_triangle_nw ScatterternaryMarkerSymbol = "triangle-nw" // triangle-nw
    ScatterternaryMarkerSymbol112 ScatterternaryMarkerSymbol = "112" // 112
    ScatterternaryMarkerSymbol_triangle_nw_open ScatterternaryMarkerSymbol = "triangle-nw-open" // triangle-nw-open
    ScatterternaryMarkerSymbol212 ScatterternaryMarkerSymbol = "212" // 212
    ScatterternaryMarkerSymbol_triangle_nw_dot ScatterternaryMarkerSymbol = "triangle-nw-dot" // triangle-nw-dot
    ScatterternaryMarkerSymbol312 ScatterternaryMarkerSymbol = "312" // 312
    ScatterternaryMarkerSymbol_triangle_nw_open_dot ScatterternaryMarkerSymbol = "triangle-nw-open-dot" // triangle-nw-open-dot
    ScatterternaryMarkerSymbol13 ScatterternaryMarkerSymbol = "13" // 13
    ScatterternaryMarkerSymbol_pentagon ScatterternaryMarkerSymbol = "pentagon" // pentagon
    ScatterternaryMarkerSymbol113 ScatterternaryMarkerSymbol = "113" // 113
    ScatterternaryMarkerSymbol_pentagon_open ScatterternaryMarkerSymbol = "pentagon-open" // pentagon-open
    ScatterternaryMarkerSymbol213 ScatterternaryMarkerSymbol = "213" // 213
    ScatterternaryMarkerSymbol_pentagon_dot ScatterternaryMarkerSymbol = "pentagon-dot" // pentagon-dot
    ScatterternaryMarkerSymbol313 ScatterternaryMarkerSymbol = "313" // 313
    ScatterternaryMarkerSymbol_pentagon_open_dot ScatterternaryMarkerSymbol = "pentagon-open-dot" // pentagon-open-dot
    ScatterternaryMarkerSymbol14 ScatterternaryMarkerSymbol = "14" // 14
    ScatterternaryMarkerSymbol_hexagon ScatterternaryMarkerSymbol = "hexagon" // hexagon
    ScatterternaryMarkerSymbol114 ScatterternaryMarkerSymbol = "114" // 114
    ScatterternaryMarkerSymbol_hexagon_open ScatterternaryMarkerSymbol = "hexagon-open" // hexagon-open
    ScatterternaryMarkerSymbol214 ScatterternaryMarkerSymbol = "214" // 214
    ScatterternaryMarkerSymbol_hexagon_dot ScatterternaryMarkerSymbol = "hexagon-dot" // hexagon-dot
    ScatterternaryMarkerSymbol314 ScatterternaryMarkerSymbol = "314" // 314
    ScatterternaryMarkerSymbol_hexagon_open_dot ScatterternaryMarkerSymbol = "hexagon-open-dot" // hexagon-open-dot
    ScatterternaryMarkerSymbol15 ScatterternaryMarkerSymbol = "15" // 15
    ScatterternaryMarkerSymbol_hexagon2 ScatterternaryMarkerSymbol = "hexagon2" // hexagon2
    ScatterternaryMarkerSymbol115 ScatterternaryMarkerSymbol = "115" // 115
    ScatterternaryMarkerSymbol_hexagon2_open ScatterternaryMarkerSymbol = "hexagon2-open" // hexagon2-open
    ScatterternaryMarkerSymbol215 ScatterternaryMarkerSymbol = "215" // 215
    ScatterternaryMarkerSymbol_hexagon2_dot ScatterternaryMarkerSymbol = "hexagon2-dot" // hexagon2-dot
    ScatterternaryMarkerSymbol315 ScatterternaryMarkerSymbol = "315" // 315
    ScatterternaryMarkerSymbol_hexagon2_open_dot ScatterternaryMarkerSymbol = "hexagon2-open-dot" // hexagon2-open-dot
    ScatterternaryMarkerSymbol16 ScatterternaryMarkerSymbol = "16" // 16
    ScatterternaryMarkerSymbol_octagon ScatterternaryMarkerSymbol = "octagon" // octagon
    ScatterternaryMarkerSymbol116 ScatterternaryMarkerSymbol = "116" // 116
    ScatterternaryMarkerSymbol_octagon_open ScatterternaryMarkerSymbol = "octagon-open" // octagon-open
    ScatterternaryMarkerSymbol216 ScatterternaryMarkerSymbol = "216" // 216
    ScatterternaryMarkerSymbol_octagon_dot ScatterternaryMarkerSymbol = "octagon-dot" // octagon-dot
    ScatterternaryMarkerSymbol316 ScatterternaryMarkerSymbol = "316" // 316
    ScatterternaryMarkerSymbol_octagon_open_dot ScatterternaryMarkerSymbol = "octagon-open-dot" // octagon-open-dot
    ScatterternaryMarkerSymbol17 ScatterternaryMarkerSymbol = "17" // 17
    ScatterternaryMarkerSymbol_star ScatterternaryMarkerSymbol = "star" // star
    ScatterternaryMarkerSymbol117 ScatterternaryMarkerSymbol = "117" // 117
    ScatterternaryMarkerSymbol_star_open ScatterternaryMarkerSymbol = "star-open" // star-open
    ScatterternaryMarkerSymbol217 ScatterternaryMarkerSymbol = "217" // 217
    ScatterternaryMarkerSymbol_star_dot ScatterternaryMarkerSymbol = "star-dot" // star-dot
    ScatterternaryMarkerSymbol317 ScatterternaryMarkerSymbol = "317" // 317
    ScatterternaryMarkerSymbol_star_open_dot ScatterternaryMarkerSymbol = "star-open-dot" // star-open-dot
    ScatterternaryMarkerSymbol18 ScatterternaryMarkerSymbol = "18" // 18
    ScatterternaryMarkerSymbol_hexagram ScatterternaryMarkerSymbol = "hexagram" // hexagram
    ScatterternaryMarkerSymbol118 ScatterternaryMarkerSymbol = "118" // 118
    ScatterternaryMarkerSymbol_hexagram_open ScatterternaryMarkerSymbol = "hexagram-open" // hexagram-open
    ScatterternaryMarkerSymbol218 ScatterternaryMarkerSymbol = "218" // 218
    ScatterternaryMarkerSymbol_hexagram_dot ScatterternaryMarkerSymbol = "hexagram-dot" // hexagram-dot
    ScatterternaryMarkerSymbol318 ScatterternaryMarkerSymbol = "318" // 318
    ScatterternaryMarkerSymbol_hexagram_open_dot ScatterternaryMarkerSymbol = "hexagram-open-dot" // hexagram-open-dot
    ScatterternaryMarkerSymbol19 ScatterternaryMarkerSymbol = "19" // 19
    ScatterternaryMarkerSymbol_star_triangle_up ScatterternaryMarkerSymbol = "star-triangle-up" // star-triangle-up
    ScatterternaryMarkerSymbol119 ScatterternaryMarkerSymbol = "119" // 119
    ScatterternaryMarkerSymbol_star_triangle_up_open ScatterternaryMarkerSymbol = "star-triangle-up-open" // star-triangle-up-open
    ScatterternaryMarkerSymbol219 ScatterternaryMarkerSymbol = "219" // 219
    ScatterternaryMarkerSymbol_star_triangle_up_dot ScatterternaryMarkerSymbol = "star-triangle-up-dot" // star-triangle-up-dot
    ScatterternaryMarkerSymbol319 ScatterternaryMarkerSymbol = "319" // 319
    ScatterternaryMarkerSymbol_star_triangle_up_open_dot ScatterternaryMarkerSymbol = "star-triangle-up-open-dot" // star-triangle-up-open-dot
    ScatterternaryMarkerSymbol20 ScatterternaryMarkerSymbol = "20" // 20
    ScatterternaryMarkerSymbol_star_triangle_down ScatterternaryMarkerSymbol = "star-triangle-down" // star-triangle-down
    ScatterternaryMarkerSymbol120 ScatterternaryMarkerSymbol = "120" // 120
    ScatterternaryMarkerSymbol_star_triangle_down_open ScatterternaryMarkerSymbol = "star-triangle-down-open" // star-triangle-down-open
    ScatterternaryMarkerSymbol220 ScatterternaryMarkerSymbol = "220" // 220
    ScatterternaryMarkerSymbol_star_triangle_down_dot ScatterternaryMarkerSymbol = "star-triangle-down-dot" // star-triangle-down-dot
    ScatterternaryMarkerSymbol320 ScatterternaryMarkerSymbol = "320" // 320
    ScatterternaryMarkerSymbol_star_triangle_down_open_dot ScatterternaryMarkerSymbol = "star-triangle-down-open-dot" // star-triangle-down-open-dot
    ScatterternaryMarkerSymbol21 ScatterternaryMarkerSymbol = "21" // 21
    ScatterternaryMarkerSymbol_star_square ScatterternaryMarkerSymbol = "star-square" // star-square
    ScatterternaryMarkerSymbol121 ScatterternaryMarkerSymbol = "121" // 121
    ScatterternaryMarkerSymbol_star_square_open ScatterternaryMarkerSymbol = "star-square-open" // star-square-open
    ScatterternaryMarkerSymbol221 ScatterternaryMarkerSymbol = "221" // 221
    ScatterternaryMarkerSymbol_star_square_dot ScatterternaryMarkerSymbol = "star-square-dot" // star-square-dot
    ScatterternaryMarkerSymbol321 ScatterternaryMarkerSymbol = "321" // 321
    ScatterternaryMarkerSymbol_star_square_open_dot ScatterternaryMarkerSymbol = "star-square-open-dot" // star-square-open-dot
    ScatterternaryMarkerSymbol22 ScatterternaryMarkerSymbol = "22" // 22
    ScatterternaryMarkerSymbol_star_diamond ScatterternaryMarkerSymbol = "star-diamond" // star-diamond
    ScatterternaryMarkerSymbol122 ScatterternaryMarkerSymbol = "122" // 122
    ScatterternaryMarkerSymbol_star_diamond_open ScatterternaryMarkerSymbol = "star-diamond-open" // star-diamond-open
    ScatterternaryMarkerSymbol222 ScatterternaryMarkerSymbol = "222" // 222
    ScatterternaryMarkerSymbol_star_diamond_dot ScatterternaryMarkerSymbol = "star-diamond-dot" // star-diamond-dot
    ScatterternaryMarkerSymbol322 ScatterternaryMarkerSymbol = "322" // 322
    ScatterternaryMarkerSymbol_star_diamond_open_dot ScatterternaryMarkerSymbol = "star-diamond-open-dot" // star-diamond-open-dot
    ScatterternaryMarkerSymbol23 ScatterternaryMarkerSymbol = "23" // 23
    ScatterternaryMarkerSymbol_diamond_tall ScatterternaryMarkerSymbol = "diamond-tall" // diamond-tall
    ScatterternaryMarkerSymbol123 ScatterternaryMarkerSymbol = "123" // 123
    ScatterternaryMarkerSymbol_diamond_tall_open ScatterternaryMarkerSymbol = "diamond-tall-open" // diamond-tall-open
    ScatterternaryMarkerSymbol223 ScatterternaryMarkerSymbol = "223" // 223
    ScatterternaryMarkerSymbol_diamond_tall_dot ScatterternaryMarkerSymbol = "diamond-tall-dot" // diamond-tall-dot
    ScatterternaryMarkerSymbol323 ScatterternaryMarkerSymbol = "323" // 323
    ScatterternaryMarkerSymbol_diamond_tall_open_dot ScatterternaryMarkerSymbol = "diamond-tall-open-dot" // diamond-tall-open-dot
    ScatterternaryMarkerSymbol24 ScatterternaryMarkerSymbol = "24" // 24
    ScatterternaryMarkerSymbol_diamond_wide ScatterternaryMarkerSymbol = "diamond-wide" // diamond-wide
    ScatterternaryMarkerSymbol124 ScatterternaryMarkerSymbol = "124" // 124
    ScatterternaryMarkerSymbol_diamond_wide_open ScatterternaryMarkerSymbol = "diamond-wide-open" // diamond-wide-open
    ScatterternaryMarkerSymbol224 ScatterternaryMarkerSymbol = "224" // 224
    ScatterternaryMarkerSymbol_diamond_wide_dot ScatterternaryMarkerSymbol = "diamond-wide-dot" // diamond-wide-dot
    ScatterternaryMarkerSymbol324 ScatterternaryMarkerSymbol = "324" // 324
    ScatterternaryMarkerSymbol_diamond_wide_open_dot ScatterternaryMarkerSymbol = "diamond-wide-open-dot" // diamond-wide-open-dot
    ScatterternaryMarkerSymbol25 ScatterternaryMarkerSymbol = "25" // 25
    ScatterternaryMarkerSymbol_hourglass ScatterternaryMarkerSymbol = "hourglass" // hourglass
    ScatterternaryMarkerSymbol125 ScatterternaryMarkerSymbol = "125" // 125
    ScatterternaryMarkerSymbol_hourglass_open ScatterternaryMarkerSymbol = "hourglass-open" // hourglass-open
    ScatterternaryMarkerSymbol26 ScatterternaryMarkerSymbol = "26" // 26
    ScatterternaryMarkerSymbol_bowtie ScatterternaryMarkerSymbol = "bowtie" // bowtie
    ScatterternaryMarkerSymbol126 ScatterternaryMarkerSymbol = "126" // 126
    ScatterternaryMarkerSymbol_bowtie_open ScatterternaryMarkerSymbol = "bowtie-open" // bowtie-open
    ScatterternaryMarkerSymbol27 ScatterternaryMarkerSymbol = "27" // 27
    ScatterternaryMarkerSymbol_circle_cross ScatterternaryMarkerSymbol = "circle-cross" // circle-cross
    ScatterternaryMarkerSymbol127 ScatterternaryMarkerSymbol = "127" // 127
    ScatterternaryMarkerSymbol_circle_cross_open ScatterternaryMarkerSymbol = "circle-cross-open" // circle-cross-open
    ScatterternaryMarkerSymbol28 ScatterternaryMarkerSymbol = "28" // 28
    ScatterternaryMarkerSymbol_circle_x ScatterternaryMarkerSymbol = "circle-x" // circle-x
    ScatterternaryMarkerSymbol128 ScatterternaryMarkerSymbol = "128" // 128
    ScatterternaryMarkerSymbol_circle_x_open ScatterternaryMarkerSymbol = "circle-x-open" // circle-x-open
    ScatterternaryMarkerSymbol29 ScatterternaryMarkerSymbol = "29" // 29
    ScatterternaryMarkerSymbol_square_cross ScatterternaryMarkerSymbol = "square-cross" // square-cross
    ScatterternaryMarkerSymbol129 ScatterternaryMarkerSymbol = "129" // 129
    ScatterternaryMarkerSymbol_square_cross_open ScatterternaryMarkerSymbol = "square-cross-open" // square-cross-open
    ScatterternaryMarkerSymbol30 ScatterternaryMarkerSymbol = "30" // 30
    ScatterternaryMarkerSymbol_square_x ScatterternaryMarkerSymbol = "square-x" // square-x
    ScatterternaryMarkerSymbol130 ScatterternaryMarkerSymbol = "130" // 130
    ScatterternaryMarkerSymbol_square_x_open ScatterternaryMarkerSymbol = "square-x-open" // square-x-open
    ScatterternaryMarkerSymbol31 ScatterternaryMarkerSymbol = "31" // 31
    ScatterternaryMarkerSymbol_diamond_cross ScatterternaryMarkerSymbol = "diamond-cross" // diamond-cross
    ScatterternaryMarkerSymbol131 ScatterternaryMarkerSymbol = "131" // 131
    ScatterternaryMarkerSymbol_diamond_cross_open ScatterternaryMarkerSymbol = "diamond-cross-open" // diamond-cross-open
    ScatterternaryMarkerSymbol32 ScatterternaryMarkerSymbol = "32" // 32
    ScatterternaryMarkerSymbol_diamond_x ScatterternaryMarkerSymbol = "diamond-x" // diamond-x
    ScatterternaryMarkerSymbol132 ScatterternaryMarkerSymbol = "132" // 132
    ScatterternaryMarkerSymbol_diamond_x_open ScatterternaryMarkerSymbol = "diamond-x-open" // diamond-x-open
    ScatterternaryMarkerSymbol33 ScatterternaryMarkerSymbol = "33" // 33
    ScatterternaryMarkerSymbol_cross_thin ScatterternaryMarkerSymbol = "cross-thin" // cross-thin
    ScatterternaryMarkerSymbol133 ScatterternaryMarkerSymbol = "133" // 133
    ScatterternaryMarkerSymbol_cross_thin_open ScatterternaryMarkerSymbol = "cross-thin-open" // cross-thin-open
    ScatterternaryMarkerSymbol34 ScatterternaryMarkerSymbol = "34" // 34
    ScatterternaryMarkerSymbol_x_thin ScatterternaryMarkerSymbol = "x-thin" // x-thin
    ScatterternaryMarkerSymbol134 ScatterternaryMarkerSymbol = "134" // 134
    ScatterternaryMarkerSymbol_x_thin_open ScatterternaryMarkerSymbol = "x-thin-open" // x-thin-open
    ScatterternaryMarkerSymbol35 ScatterternaryMarkerSymbol = "35" // 35
    ScatterternaryMarkerSymbol_asterisk ScatterternaryMarkerSymbol = "asterisk" // asterisk
    ScatterternaryMarkerSymbol135 ScatterternaryMarkerSymbol = "135" // 135
    ScatterternaryMarkerSymbol_asterisk_open ScatterternaryMarkerSymbol = "asterisk-open" // asterisk-open
    ScatterternaryMarkerSymbol36 ScatterternaryMarkerSymbol = "36" // 36
    ScatterternaryMarkerSymbol_hash ScatterternaryMarkerSymbol = "hash" // hash
    ScatterternaryMarkerSymbol136 ScatterternaryMarkerSymbol = "136" // 136
    ScatterternaryMarkerSymbol_hash_open ScatterternaryMarkerSymbol = "hash-open" // hash-open
    ScatterternaryMarkerSymbol236 ScatterternaryMarkerSymbol = "236" // 236
    ScatterternaryMarkerSymbol_hash_dot ScatterternaryMarkerSymbol = "hash-dot" // hash-dot
    ScatterternaryMarkerSymbol336 ScatterternaryMarkerSymbol = "336" // 336
    ScatterternaryMarkerSymbol_hash_open_dot ScatterternaryMarkerSymbol = "hash-open-dot" // hash-open-dot
    ScatterternaryMarkerSymbol37 ScatterternaryMarkerSymbol = "37" // 37
    ScatterternaryMarkerSymbol_y_up ScatterternaryMarkerSymbol = "y-up" // y-up
    ScatterternaryMarkerSymbol137 ScatterternaryMarkerSymbol = "137" // 137
    ScatterternaryMarkerSymbol_y_up_open ScatterternaryMarkerSymbol = "y-up-open" // y-up-open
    ScatterternaryMarkerSymbol38 ScatterternaryMarkerSymbol = "38" // 38
    ScatterternaryMarkerSymbol_y_down ScatterternaryMarkerSymbol = "y-down" // y-down
    ScatterternaryMarkerSymbol138 ScatterternaryMarkerSymbol = "138" // 138
    ScatterternaryMarkerSymbol_y_down_open ScatterternaryMarkerSymbol = "y-down-open" // y-down-open
    ScatterternaryMarkerSymbol39 ScatterternaryMarkerSymbol = "39" // 39
    ScatterternaryMarkerSymbol_y_left ScatterternaryMarkerSymbol = "y-left" // y-left
    ScatterternaryMarkerSymbol139 ScatterternaryMarkerSymbol = "139" // 139
    ScatterternaryMarkerSymbol_y_left_open ScatterternaryMarkerSymbol = "y-left-open" // y-left-open
    ScatterternaryMarkerSymbol40 ScatterternaryMarkerSymbol = "40" // 40
    ScatterternaryMarkerSymbol_y_right ScatterternaryMarkerSymbol = "y-right" // y-right
    ScatterternaryMarkerSymbol140 ScatterternaryMarkerSymbol = "140" // 140
    ScatterternaryMarkerSymbol_y_right_open ScatterternaryMarkerSymbol = "y-right-open" // y-right-open
    ScatterternaryMarkerSymbol41 ScatterternaryMarkerSymbol = "41" // 41
    ScatterternaryMarkerSymbol_line_ew ScatterternaryMarkerSymbol = "line-ew" // line-ew
    ScatterternaryMarkerSymbol141 ScatterternaryMarkerSymbol = "141" // 141
    ScatterternaryMarkerSymbol_line_ew_open ScatterternaryMarkerSymbol = "line-ew-open" // line-ew-open
    ScatterternaryMarkerSymbol42 ScatterternaryMarkerSymbol = "42" // 42
    ScatterternaryMarkerSymbol_line_ns ScatterternaryMarkerSymbol = "line-ns" // line-ns
    ScatterternaryMarkerSymbol142 ScatterternaryMarkerSymbol = "142" // 142
    ScatterternaryMarkerSymbol_line_ns_open ScatterternaryMarkerSymbol = "line-ns-open" // line-ns-open
    ScatterternaryMarkerSymbol43 ScatterternaryMarkerSymbol = "43" // 43
    ScatterternaryMarkerSymbol_line_ne ScatterternaryMarkerSymbol = "line-ne" // line-ne
    ScatterternaryMarkerSymbol143 ScatterternaryMarkerSymbol = "143" // 143
    ScatterternaryMarkerSymbol_line_ne_open ScatterternaryMarkerSymbol = "line-ne-open" // line-ne-open
    ScatterternaryMarkerSymbol44 ScatterternaryMarkerSymbol = "44" // 44
    ScatterternaryMarkerSymbol_line_nw ScatterternaryMarkerSymbol = "line-nw" // line-nw
    ScatterternaryMarkerSymbol144 ScatterternaryMarkerSymbol = "144" // 144
    ScatterternaryMarkerSymbol_line_nw_open ScatterternaryMarkerSymbol = "line-nw-open" // line-nw-open
)

// ScatterternaryTextposition Sets the positions of the `text` elements with respects to the (x,y) coordinates.
type ScatterternaryTextposition string

const (
    ScatterternaryTextposition_topleft ScatterternaryTextposition = "top left" // top left
    ScatterternaryTextposition_topcenter ScatterternaryTextposition = "top center" // top center
    ScatterternaryTextposition_topright ScatterternaryTextposition = "top right" // top right
    ScatterternaryTextposition_middleleft ScatterternaryTextposition = "middle left" // middle left
    ScatterternaryTextposition_middlecenter ScatterternaryTextposition = "middle center" // middle center
    ScatterternaryTextposition_middleright ScatterternaryTextposition = "middle right" // middle right
    ScatterternaryTextposition_bottomleft ScatterternaryTextposition = "bottom left" // bottom left
    ScatterternaryTextposition_bottomcenter ScatterternaryTextposition = "bottom center" // bottom center
    ScatterternaryTextposition_bottomright ScatterternaryTextposition = "bottom right" // bottom right
)

// ScatterternaryVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ScatterternaryVisible interface{}

const (
    ScatterternaryVisibleTrue bool = true
    ScatterternaryVisibleFalse bool = false
    ScatterternaryVisibleLegendonly string = "legendonly"
)

// SplomHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SplomHoverlabelAlign string

const (
    SplomHoverlabelAlign_left SplomHoverlabelAlign = "left" // left
    SplomHoverlabelAlign_right SplomHoverlabelAlign = "right" // right
    SplomHoverlabelAlign_auto SplomHoverlabelAlign = "auto" // auto
)

// SplomMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type SplomMarkerColorbarExponentformat string

const (
    SplomMarkerColorbarExponentformat_none SplomMarkerColorbarExponentformat = "none" // none
    SplomMarkerColorbarExponentformat_e SplomMarkerColorbarExponentformat = "e" // e
    SplomMarkerColorbarExponentformat_E SplomMarkerColorbarExponentformat = "E" // E
    SplomMarkerColorbarExponentformat_power SplomMarkerColorbarExponentformat = "power" // power
    SplomMarkerColorbarExponentformat_SI SplomMarkerColorbarExponentformat = "SI" // SI
    SplomMarkerColorbarExponentformat_B SplomMarkerColorbarExponentformat = "B" // B
)

// SplomMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type SplomMarkerColorbarLenmode string

const (
    SplomMarkerColorbarLenmode_fraction SplomMarkerColorbarLenmode = "fraction" // fraction
    SplomMarkerColorbarLenmode_pixels SplomMarkerColorbarLenmode = "pixels" // pixels
)

// SplomMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type SplomMarkerColorbarShowexponent string

const (
    SplomMarkerColorbarShowexponent_all SplomMarkerColorbarShowexponent = "all" // all
    SplomMarkerColorbarShowexponent_first SplomMarkerColorbarShowexponent = "first" // first
    SplomMarkerColorbarShowexponent_last SplomMarkerColorbarShowexponent = "last" // last
    SplomMarkerColorbarShowexponent_none SplomMarkerColorbarShowexponent = "none" // none
)

// SplomMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type SplomMarkerColorbarShowtickprefix string

const (
    SplomMarkerColorbarShowtickprefix_all SplomMarkerColorbarShowtickprefix = "all" // all
    SplomMarkerColorbarShowtickprefix_first SplomMarkerColorbarShowtickprefix = "first" // first
    SplomMarkerColorbarShowtickprefix_last SplomMarkerColorbarShowtickprefix = "last" // last
    SplomMarkerColorbarShowtickprefix_none SplomMarkerColorbarShowtickprefix = "none" // none
)

// SplomMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type SplomMarkerColorbarShowticksuffix string

const (
    SplomMarkerColorbarShowticksuffix_all SplomMarkerColorbarShowticksuffix = "all" // all
    SplomMarkerColorbarShowticksuffix_first SplomMarkerColorbarShowticksuffix = "first" // first
    SplomMarkerColorbarShowticksuffix_last SplomMarkerColorbarShowticksuffix = "last" // last
    SplomMarkerColorbarShowticksuffix_none SplomMarkerColorbarShowticksuffix = "none" // none
)

// SplomMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type SplomMarkerColorbarThicknessmode string

const (
    SplomMarkerColorbarThicknessmode_fraction SplomMarkerColorbarThicknessmode = "fraction" // fraction
    SplomMarkerColorbarThicknessmode_pixels SplomMarkerColorbarThicknessmode = "pixels" // pixels
)

// SplomMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type SplomMarkerColorbarTickmode string

const (
    SplomMarkerColorbarTickmode_auto SplomMarkerColorbarTickmode = "auto" // auto
    SplomMarkerColorbarTickmode_linear SplomMarkerColorbarTickmode = "linear" // linear
    SplomMarkerColorbarTickmode_array SplomMarkerColorbarTickmode = "array" // array
)

// SplomMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type SplomMarkerColorbarTicks string

const (
    SplomMarkerColorbarTicks_outside SplomMarkerColorbarTicks = "outside" // outside
    SplomMarkerColorbarTicks_inside SplomMarkerColorbarTicks = "inside" // inside
)

// SplomMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type SplomMarkerColorbarTitleSide string

const (
    SplomMarkerColorbarTitleSide_right SplomMarkerColorbarTitleSide = "right" // right
    SplomMarkerColorbarTitleSide_top SplomMarkerColorbarTitleSide = "top" // top
    SplomMarkerColorbarTitleSide_bottom SplomMarkerColorbarTitleSide = "bottom" // bottom
)

// SplomMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type SplomMarkerColorbarXanchor string

const (
    SplomMarkerColorbarXanchor_left SplomMarkerColorbarXanchor = "left" // left
    SplomMarkerColorbarXanchor_center SplomMarkerColorbarXanchor = "center" // center
    SplomMarkerColorbarXanchor_right SplomMarkerColorbarXanchor = "right" // right
)

// SplomMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type SplomMarkerColorbarYanchor string

const (
    SplomMarkerColorbarYanchor_top SplomMarkerColorbarYanchor = "top" // top
    SplomMarkerColorbarYanchor_middle SplomMarkerColorbarYanchor = "middle" // middle
    SplomMarkerColorbarYanchor_bottom SplomMarkerColorbarYanchor = "bottom" // bottom
)

// SplomMarkerSizemode Has an effect only if `marker.size` is set to a numerical array. Sets the rule for which the data in `size` is converted to pixels.
type SplomMarkerSizemode string

const (
    SplomMarkerSizemode_diameter SplomMarkerSizemode = "diameter" // diameter
    SplomMarkerSizemode_area SplomMarkerSizemode = "area" // area
)

// SplomMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type SplomMarkerSymbol string

const (
    SplomMarkerSymbol0 SplomMarkerSymbol = "0" // 0
    SplomMarkerSymbol_circle SplomMarkerSymbol = "circle" // circle
    SplomMarkerSymbol100 SplomMarkerSymbol = "100" // 100
    SplomMarkerSymbol_circle_open SplomMarkerSymbol = "circle-open" // circle-open
    SplomMarkerSymbol200 SplomMarkerSymbol = "200" // 200
    SplomMarkerSymbol_circle_dot SplomMarkerSymbol = "circle-dot" // circle-dot
    SplomMarkerSymbol300 SplomMarkerSymbol = "300" // 300
    SplomMarkerSymbol_circle_open_dot SplomMarkerSymbol = "circle-open-dot" // circle-open-dot
    SplomMarkerSymbol1 SplomMarkerSymbol = "1" // 1
    SplomMarkerSymbol_square SplomMarkerSymbol = "square" // square
    SplomMarkerSymbol101 SplomMarkerSymbol = "101" // 101
    SplomMarkerSymbol_square_open SplomMarkerSymbol = "square-open" // square-open
    SplomMarkerSymbol201 SplomMarkerSymbol = "201" // 201
    SplomMarkerSymbol_square_dot SplomMarkerSymbol = "square-dot" // square-dot
    SplomMarkerSymbol301 SplomMarkerSymbol = "301" // 301
    SplomMarkerSymbol_square_open_dot SplomMarkerSymbol = "square-open-dot" // square-open-dot
    SplomMarkerSymbol2 SplomMarkerSymbol = "2" // 2
    SplomMarkerSymbol_diamond SplomMarkerSymbol = "diamond" // diamond
    SplomMarkerSymbol102 SplomMarkerSymbol = "102" // 102
    SplomMarkerSymbol_diamond_open SplomMarkerSymbol = "diamond-open" // diamond-open
    SplomMarkerSymbol202 SplomMarkerSymbol = "202" // 202
    SplomMarkerSymbol_diamond_dot SplomMarkerSymbol = "diamond-dot" // diamond-dot
    SplomMarkerSymbol302 SplomMarkerSymbol = "302" // 302
    SplomMarkerSymbol_diamond_open_dot SplomMarkerSymbol = "diamond-open-dot" // diamond-open-dot
    SplomMarkerSymbol3 SplomMarkerSymbol = "3" // 3
    SplomMarkerSymbol_cross SplomMarkerSymbol = "cross" // cross
    SplomMarkerSymbol103 SplomMarkerSymbol = "103" // 103
    SplomMarkerSymbol_cross_open SplomMarkerSymbol = "cross-open" // cross-open
    SplomMarkerSymbol203 SplomMarkerSymbol = "203" // 203
    SplomMarkerSymbol_cross_dot SplomMarkerSymbol = "cross-dot" // cross-dot
    SplomMarkerSymbol303 SplomMarkerSymbol = "303" // 303
    SplomMarkerSymbol_cross_open_dot SplomMarkerSymbol = "cross-open-dot" // cross-open-dot
    SplomMarkerSymbol4 SplomMarkerSymbol = "4" // 4
    SplomMarkerSymbol_x SplomMarkerSymbol = "x" // x
    SplomMarkerSymbol104 SplomMarkerSymbol = "104" // 104
    SplomMarkerSymbol_x_open SplomMarkerSymbol = "x-open" // x-open
    SplomMarkerSymbol204 SplomMarkerSymbol = "204" // 204
    SplomMarkerSymbol_x_dot SplomMarkerSymbol = "x-dot" // x-dot
    SplomMarkerSymbol304 SplomMarkerSymbol = "304" // 304
    SplomMarkerSymbol_x_open_dot SplomMarkerSymbol = "x-open-dot" // x-open-dot
    SplomMarkerSymbol5 SplomMarkerSymbol = "5" // 5
    SplomMarkerSymbol_triangle_up SplomMarkerSymbol = "triangle-up" // triangle-up
    SplomMarkerSymbol105 SplomMarkerSymbol = "105" // 105
    SplomMarkerSymbol_triangle_up_open SplomMarkerSymbol = "triangle-up-open" // triangle-up-open
    SplomMarkerSymbol205 SplomMarkerSymbol = "205" // 205
    SplomMarkerSymbol_triangle_up_dot SplomMarkerSymbol = "triangle-up-dot" // triangle-up-dot
    SplomMarkerSymbol305 SplomMarkerSymbol = "305" // 305
    SplomMarkerSymbol_triangle_up_open_dot SplomMarkerSymbol = "triangle-up-open-dot" // triangle-up-open-dot
    SplomMarkerSymbol6 SplomMarkerSymbol = "6" // 6
    SplomMarkerSymbol_triangle_down SplomMarkerSymbol = "triangle-down" // triangle-down
    SplomMarkerSymbol106 SplomMarkerSymbol = "106" // 106
    SplomMarkerSymbol_triangle_down_open SplomMarkerSymbol = "triangle-down-open" // triangle-down-open
    SplomMarkerSymbol206 SplomMarkerSymbol = "206" // 206
    SplomMarkerSymbol_triangle_down_dot SplomMarkerSymbol = "triangle-down-dot" // triangle-down-dot
    SplomMarkerSymbol306 SplomMarkerSymbol = "306" // 306
    SplomMarkerSymbol_triangle_down_open_dot SplomMarkerSymbol = "triangle-down-open-dot" // triangle-down-open-dot
    SplomMarkerSymbol7 SplomMarkerSymbol = "7" // 7
    SplomMarkerSymbol_triangle_left SplomMarkerSymbol = "triangle-left" // triangle-left
    SplomMarkerSymbol107 SplomMarkerSymbol = "107" // 107
    SplomMarkerSymbol_triangle_left_open SplomMarkerSymbol = "triangle-left-open" // triangle-left-open
    SplomMarkerSymbol207 SplomMarkerSymbol = "207" // 207
    SplomMarkerSymbol_triangle_left_dot SplomMarkerSymbol = "triangle-left-dot" // triangle-left-dot
    SplomMarkerSymbol307 SplomMarkerSymbol = "307" // 307
    SplomMarkerSymbol_triangle_left_open_dot SplomMarkerSymbol = "triangle-left-open-dot" // triangle-left-open-dot
    SplomMarkerSymbol8 SplomMarkerSymbol = "8" // 8
    SplomMarkerSymbol_triangle_right SplomMarkerSymbol = "triangle-right" // triangle-right
    SplomMarkerSymbol108 SplomMarkerSymbol = "108" // 108
    SplomMarkerSymbol_triangle_right_open SplomMarkerSymbol = "triangle-right-open" // triangle-right-open
    SplomMarkerSymbol208 SplomMarkerSymbol = "208" // 208
    SplomMarkerSymbol_triangle_right_dot SplomMarkerSymbol = "triangle-right-dot" // triangle-right-dot
    SplomMarkerSymbol308 SplomMarkerSymbol = "308" // 308
    SplomMarkerSymbol_triangle_right_open_dot SplomMarkerSymbol = "triangle-right-open-dot" // triangle-right-open-dot
    SplomMarkerSymbol9 SplomMarkerSymbol = "9" // 9
    SplomMarkerSymbol_triangle_ne SplomMarkerSymbol = "triangle-ne" // triangle-ne
    SplomMarkerSymbol109 SplomMarkerSymbol = "109" // 109
    SplomMarkerSymbol_triangle_ne_open SplomMarkerSymbol = "triangle-ne-open" // triangle-ne-open
    SplomMarkerSymbol209 SplomMarkerSymbol = "209" // 209
    SplomMarkerSymbol_triangle_ne_dot SplomMarkerSymbol = "triangle-ne-dot" // triangle-ne-dot
    SplomMarkerSymbol309 SplomMarkerSymbol = "309" // 309
    SplomMarkerSymbol_triangle_ne_open_dot SplomMarkerSymbol = "triangle-ne-open-dot" // triangle-ne-open-dot
    SplomMarkerSymbol10 SplomMarkerSymbol = "10" // 10
    SplomMarkerSymbol_triangle_se SplomMarkerSymbol = "triangle-se" // triangle-se
    SplomMarkerSymbol110 SplomMarkerSymbol = "110" // 110
    SplomMarkerSymbol_triangle_se_open SplomMarkerSymbol = "triangle-se-open" // triangle-se-open
    SplomMarkerSymbol210 SplomMarkerSymbol = "210" // 210
    SplomMarkerSymbol_triangle_se_dot SplomMarkerSymbol = "triangle-se-dot" // triangle-se-dot
    SplomMarkerSymbol310 SplomMarkerSymbol = "310" // 310
    SplomMarkerSymbol_triangle_se_open_dot SplomMarkerSymbol = "triangle-se-open-dot" // triangle-se-open-dot
    SplomMarkerSymbol11 SplomMarkerSymbol = "11" // 11
    SplomMarkerSymbol_triangle_sw SplomMarkerSymbol = "triangle-sw" // triangle-sw
    SplomMarkerSymbol111 SplomMarkerSymbol = "111" // 111
    SplomMarkerSymbol_triangle_sw_open SplomMarkerSymbol = "triangle-sw-open" // triangle-sw-open
    SplomMarkerSymbol211 SplomMarkerSymbol = "211" // 211
    SplomMarkerSymbol_triangle_sw_dot SplomMarkerSymbol = "triangle-sw-dot" // triangle-sw-dot
    SplomMarkerSymbol311 SplomMarkerSymbol = "311" // 311
    SplomMarkerSymbol_triangle_sw_open_dot SplomMarkerSymbol = "triangle-sw-open-dot" // triangle-sw-open-dot
    SplomMarkerSymbol12 SplomMarkerSymbol = "12" // 12
    SplomMarkerSymbol_triangle_nw SplomMarkerSymbol = "triangle-nw" // triangle-nw
    SplomMarkerSymbol112 SplomMarkerSymbol = "112" // 112
    SplomMarkerSymbol_triangle_nw_open SplomMarkerSymbol = "triangle-nw-open" // triangle-nw-open
    SplomMarkerSymbol212 SplomMarkerSymbol = "212" // 212
    SplomMarkerSymbol_triangle_nw_dot SplomMarkerSymbol = "triangle-nw-dot" // triangle-nw-dot
    SplomMarkerSymbol312 SplomMarkerSymbol = "312" // 312
    SplomMarkerSymbol_triangle_nw_open_dot SplomMarkerSymbol = "triangle-nw-open-dot" // triangle-nw-open-dot
    SplomMarkerSymbol13 SplomMarkerSymbol = "13" // 13
    SplomMarkerSymbol_pentagon SplomMarkerSymbol = "pentagon" // pentagon
    SplomMarkerSymbol113 SplomMarkerSymbol = "113" // 113
    SplomMarkerSymbol_pentagon_open SplomMarkerSymbol = "pentagon-open" // pentagon-open
    SplomMarkerSymbol213 SplomMarkerSymbol = "213" // 213
    SplomMarkerSymbol_pentagon_dot SplomMarkerSymbol = "pentagon-dot" // pentagon-dot
    SplomMarkerSymbol313 SplomMarkerSymbol = "313" // 313
    SplomMarkerSymbol_pentagon_open_dot SplomMarkerSymbol = "pentagon-open-dot" // pentagon-open-dot
    SplomMarkerSymbol14 SplomMarkerSymbol = "14" // 14
    SplomMarkerSymbol_hexagon SplomMarkerSymbol = "hexagon" // hexagon
    SplomMarkerSymbol114 SplomMarkerSymbol = "114" // 114
    SplomMarkerSymbol_hexagon_open SplomMarkerSymbol = "hexagon-open" // hexagon-open
    SplomMarkerSymbol214 SplomMarkerSymbol = "214" // 214
    SplomMarkerSymbol_hexagon_dot SplomMarkerSymbol = "hexagon-dot" // hexagon-dot
    SplomMarkerSymbol314 SplomMarkerSymbol = "314" // 314
    SplomMarkerSymbol_hexagon_open_dot SplomMarkerSymbol = "hexagon-open-dot" // hexagon-open-dot
    SplomMarkerSymbol15 SplomMarkerSymbol = "15" // 15
    SplomMarkerSymbol_hexagon2 SplomMarkerSymbol = "hexagon2" // hexagon2
    SplomMarkerSymbol115 SplomMarkerSymbol = "115" // 115
    SplomMarkerSymbol_hexagon2_open SplomMarkerSymbol = "hexagon2-open" // hexagon2-open
    SplomMarkerSymbol215 SplomMarkerSymbol = "215" // 215
    SplomMarkerSymbol_hexagon2_dot SplomMarkerSymbol = "hexagon2-dot" // hexagon2-dot
    SplomMarkerSymbol315 SplomMarkerSymbol = "315" // 315
    SplomMarkerSymbol_hexagon2_open_dot SplomMarkerSymbol = "hexagon2-open-dot" // hexagon2-open-dot
    SplomMarkerSymbol16 SplomMarkerSymbol = "16" // 16
    SplomMarkerSymbol_octagon SplomMarkerSymbol = "octagon" // octagon
    SplomMarkerSymbol116 SplomMarkerSymbol = "116" // 116
    SplomMarkerSymbol_octagon_open SplomMarkerSymbol = "octagon-open" // octagon-open
    SplomMarkerSymbol216 SplomMarkerSymbol = "216" // 216
    SplomMarkerSymbol_octagon_dot SplomMarkerSymbol = "octagon-dot" // octagon-dot
    SplomMarkerSymbol316 SplomMarkerSymbol = "316" // 316
    SplomMarkerSymbol_octagon_open_dot SplomMarkerSymbol = "octagon-open-dot" // octagon-open-dot
    SplomMarkerSymbol17 SplomMarkerSymbol = "17" // 17
    SplomMarkerSymbol_star SplomMarkerSymbol = "star" // star
    SplomMarkerSymbol117 SplomMarkerSymbol = "117" // 117
    SplomMarkerSymbol_star_open SplomMarkerSymbol = "star-open" // star-open
    SplomMarkerSymbol217 SplomMarkerSymbol = "217" // 217
    SplomMarkerSymbol_star_dot SplomMarkerSymbol = "star-dot" // star-dot
    SplomMarkerSymbol317 SplomMarkerSymbol = "317" // 317
    SplomMarkerSymbol_star_open_dot SplomMarkerSymbol = "star-open-dot" // star-open-dot
    SplomMarkerSymbol18 SplomMarkerSymbol = "18" // 18
    SplomMarkerSymbol_hexagram SplomMarkerSymbol = "hexagram" // hexagram
    SplomMarkerSymbol118 SplomMarkerSymbol = "118" // 118
    SplomMarkerSymbol_hexagram_open SplomMarkerSymbol = "hexagram-open" // hexagram-open
    SplomMarkerSymbol218 SplomMarkerSymbol = "218" // 218
    SplomMarkerSymbol_hexagram_dot SplomMarkerSymbol = "hexagram-dot" // hexagram-dot
    SplomMarkerSymbol318 SplomMarkerSymbol = "318" // 318
    SplomMarkerSymbol_hexagram_open_dot SplomMarkerSymbol = "hexagram-open-dot" // hexagram-open-dot
    SplomMarkerSymbol19 SplomMarkerSymbol = "19" // 19
    SplomMarkerSymbol_star_triangle_up SplomMarkerSymbol = "star-triangle-up" // star-triangle-up
    SplomMarkerSymbol119 SplomMarkerSymbol = "119" // 119
    SplomMarkerSymbol_star_triangle_up_open SplomMarkerSymbol = "star-triangle-up-open" // star-triangle-up-open
    SplomMarkerSymbol219 SplomMarkerSymbol = "219" // 219
    SplomMarkerSymbol_star_triangle_up_dot SplomMarkerSymbol = "star-triangle-up-dot" // star-triangle-up-dot
    SplomMarkerSymbol319 SplomMarkerSymbol = "319" // 319
    SplomMarkerSymbol_star_triangle_up_open_dot SplomMarkerSymbol = "star-triangle-up-open-dot" // star-triangle-up-open-dot
    SplomMarkerSymbol20 SplomMarkerSymbol = "20" // 20
    SplomMarkerSymbol_star_triangle_down SplomMarkerSymbol = "star-triangle-down" // star-triangle-down
    SplomMarkerSymbol120 SplomMarkerSymbol = "120" // 120
    SplomMarkerSymbol_star_triangle_down_open SplomMarkerSymbol = "star-triangle-down-open" // star-triangle-down-open
    SplomMarkerSymbol220 SplomMarkerSymbol = "220" // 220
    SplomMarkerSymbol_star_triangle_down_dot SplomMarkerSymbol = "star-triangle-down-dot" // star-triangle-down-dot
    SplomMarkerSymbol320 SplomMarkerSymbol = "320" // 320
    SplomMarkerSymbol_star_triangle_down_open_dot SplomMarkerSymbol = "star-triangle-down-open-dot" // star-triangle-down-open-dot
    SplomMarkerSymbol21 SplomMarkerSymbol = "21" // 21
    SplomMarkerSymbol_star_square SplomMarkerSymbol = "star-square" // star-square
    SplomMarkerSymbol121 SplomMarkerSymbol = "121" // 121
    SplomMarkerSymbol_star_square_open SplomMarkerSymbol = "star-square-open" // star-square-open
    SplomMarkerSymbol221 SplomMarkerSymbol = "221" // 221
    SplomMarkerSymbol_star_square_dot SplomMarkerSymbol = "star-square-dot" // star-square-dot
    SplomMarkerSymbol321 SplomMarkerSymbol = "321" // 321
    SplomMarkerSymbol_star_square_open_dot SplomMarkerSymbol = "star-square-open-dot" // star-square-open-dot
    SplomMarkerSymbol22 SplomMarkerSymbol = "22" // 22
    SplomMarkerSymbol_star_diamond SplomMarkerSymbol = "star-diamond" // star-diamond
    SplomMarkerSymbol122 SplomMarkerSymbol = "122" // 122
    SplomMarkerSymbol_star_diamond_open SplomMarkerSymbol = "star-diamond-open" // star-diamond-open
    SplomMarkerSymbol222 SplomMarkerSymbol = "222" // 222
    SplomMarkerSymbol_star_diamond_dot SplomMarkerSymbol = "star-diamond-dot" // star-diamond-dot
    SplomMarkerSymbol322 SplomMarkerSymbol = "322" // 322
    SplomMarkerSymbol_star_diamond_open_dot SplomMarkerSymbol = "star-diamond-open-dot" // star-diamond-open-dot
    SplomMarkerSymbol23 SplomMarkerSymbol = "23" // 23
    SplomMarkerSymbol_diamond_tall SplomMarkerSymbol = "diamond-tall" // diamond-tall
    SplomMarkerSymbol123 SplomMarkerSymbol = "123" // 123
    SplomMarkerSymbol_diamond_tall_open SplomMarkerSymbol = "diamond-tall-open" // diamond-tall-open
    SplomMarkerSymbol223 SplomMarkerSymbol = "223" // 223
    SplomMarkerSymbol_diamond_tall_dot SplomMarkerSymbol = "diamond-tall-dot" // diamond-tall-dot
    SplomMarkerSymbol323 SplomMarkerSymbol = "323" // 323
    SplomMarkerSymbol_diamond_tall_open_dot SplomMarkerSymbol = "diamond-tall-open-dot" // diamond-tall-open-dot
    SplomMarkerSymbol24 SplomMarkerSymbol = "24" // 24
    SplomMarkerSymbol_diamond_wide SplomMarkerSymbol = "diamond-wide" // diamond-wide
    SplomMarkerSymbol124 SplomMarkerSymbol = "124" // 124
    SplomMarkerSymbol_diamond_wide_open SplomMarkerSymbol = "diamond-wide-open" // diamond-wide-open
    SplomMarkerSymbol224 SplomMarkerSymbol = "224" // 224
    SplomMarkerSymbol_diamond_wide_dot SplomMarkerSymbol = "diamond-wide-dot" // diamond-wide-dot
    SplomMarkerSymbol324 SplomMarkerSymbol = "324" // 324
    SplomMarkerSymbol_diamond_wide_open_dot SplomMarkerSymbol = "diamond-wide-open-dot" // diamond-wide-open-dot
    SplomMarkerSymbol25 SplomMarkerSymbol = "25" // 25
    SplomMarkerSymbol_hourglass SplomMarkerSymbol = "hourglass" // hourglass
    SplomMarkerSymbol125 SplomMarkerSymbol = "125" // 125
    SplomMarkerSymbol_hourglass_open SplomMarkerSymbol = "hourglass-open" // hourglass-open
    SplomMarkerSymbol26 SplomMarkerSymbol = "26" // 26
    SplomMarkerSymbol_bowtie SplomMarkerSymbol = "bowtie" // bowtie
    SplomMarkerSymbol126 SplomMarkerSymbol = "126" // 126
    SplomMarkerSymbol_bowtie_open SplomMarkerSymbol = "bowtie-open" // bowtie-open
    SplomMarkerSymbol27 SplomMarkerSymbol = "27" // 27
    SplomMarkerSymbol_circle_cross SplomMarkerSymbol = "circle-cross" // circle-cross
    SplomMarkerSymbol127 SplomMarkerSymbol = "127" // 127
    SplomMarkerSymbol_circle_cross_open SplomMarkerSymbol = "circle-cross-open" // circle-cross-open
    SplomMarkerSymbol28 SplomMarkerSymbol = "28" // 28
    SplomMarkerSymbol_circle_x SplomMarkerSymbol = "circle-x" // circle-x
    SplomMarkerSymbol128 SplomMarkerSymbol = "128" // 128
    SplomMarkerSymbol_circle_x_open SplomMarkerSymbol = "circle-x-open" // circle-x-open
    SplomMarkerSymbol29 SplomMarkerSymbol = "29" // 29
    SplomMarkerSymbol_square_cross SplomMarkerSymbol = "square-cross" // square-cross
    SplomMarkerSymbol129 SplomMarkerSymbol = "129" // 129
    SplomMarkerSymbol_square_cross_open SplomMarkerSymbol = "square-cross-open" // square-cross-open
    SplomMarkerSymbol30 SplomMarkerSymbol = "30" // 30
    SplomMarkerSymbol_square_x SplomMarkerSymbol = "square-x" // square-x
    SplomMarkerSymbol130 SplomMarkerSymbol = "130" // 130
    SplomMarkerSymbol_square_x_open SplomMarkerSymbol = "square-x-open" // square-x-open
    SplomMarkerSymbol31 SplomMarkerSymbol = "31" // 31
    SplomMarkerSymbol_diamond_cross SplomMarkerSymbol = "diamond-cross" // diamond-cross
    SplomMarkerSymbol131 SplomMarkerSymbol = "131" // 131
    SplomMarkerSymbol_diamond_cross_open SplomMarkerSymbol = "diamond-cross-open" // diamond-cross-open
    SplomMarkerSymbol32 SplomMarkerSymbol = "32" // 32
    SplomMarkerSymbol_diamond_x SplomMarkerSymbol = "diamond-x" // diamond-x
    SplomMarkerSymbol132 SplomMarkerSymbol = "132" // 132
    SplomMarkerSymbol_diamond_x_open SplomMarkerSymbol = "diamond-x-open" // diamond-x-open
    SplomMarkerSymbol33 SplomMarkerSymbol = "33" // 33
    SplomMarkerSymbol_cross_thin SplomMarkerSymbol = "cross-thin" // cross-thin
    SplomMarkerSymbol133 SplomMarkerSymbol = "133" // 133
    SplomMarkerSymbol_cross_thin_open SplomMarkerSymbol = "cross-thin-open" // cross-thin-open
    SplomMarkerSymbol34 SplomMarkerSymbol = "34" // 34
    SplomMarkerSymbol_x_thin SplomMarkerSymbol = "x-thin" // x-thin
    SplomMarkerSymbol134 SplomMarkerSymbol = "134" // 134
    SplomMarkerSymbol_x_thin_open SplomMarkerSymbol = "x-thin-open" // x-thin-open
    SplomMarkerSymbol35 SplomMarkerSymbol = "35" // 35
    SplomMarkerSymbol_asterisk SplomMarkerSymbol = "asterisk" // asterisk
    SplomMarkerSymbol135 SplomMarkerSymbol = "135" // 135
    SplomMarkerSymbol_asterisk_open SplomMarkerSymbol = "asterisk-open" // asterisk-open
    SplomMarkerSymbol36 SplomMarkerSymbol = "36" // 36
    SplomMarkerSymbol_hash SplomMarkerSymbol = "hash" // hash
    SplomMarkerSymbol136 SplomMarkerSymbol = "136" // 136
    SplomMarkerSymbol_hash_open SplomMarkerSymbol = "hash-open" // hash-open
    SplomMarkerSymbol236 SplomMarkerSymbol = "236" // 236
    SplomMarkerSymbol_hash_dot SplomMarkerSymbol = "hash-dot" // hash-dot
    SplomMarkerSymbol336 SplomMarkerSymbol = "336" // 336
    SplomMarkerSymbol_hash_open_dot SplomMarkerSymbol = "hash-open-dot" // hash-open-dot
    SplomMarkerSymbol37 SplomMarkerSymbol = "37" // 37
    SplomMarkerSymbol_y_up SplomMarkerSymbol = "y-up" // y-up
    SplomMarkerSymbol137 SplomMarkerSymbol = "137" // 137
    SplomMarkerSymbol_y_up_open SplomMarkerSymbol = "y-up-open" // y-up-open
    SplomMarkerSymbol38 SplomMarkerSymbol = "38" // 38
    SplomMarkerSymbol_y_down SplomMarkerSymbol = "y-down" // y-down
    SplomMarkerSymbol138 SplomMarkerSymbol = "138" // 138
    SplomMarkerSymbol_y_down_open SplomMarkerSymbol = "y-down-open" // y-down-open
    SplomMarkerSymbol39 SplomMarkerSymbol = "39" // 39
    SplomMarkerSymbol_y_left SplomMarkerSymbol = "y-left" // y-left
    SplomMarkerSymbol139 SplomMarkerSymbol = "139" // 139
    SplomMarkerSymbol_y_left_open SplomMarkerSymbol = "y-left-open" // y-left-open
    SplomMarkerSymbol40 SplomMarkerSymbol = "40" // 40
    SplomMarkerSymbol_y_right SplomMarkerSymbol = "y-right" // y-right
    SplomMarkerSymbol140 SplomMarkerSymbol = "140" // 140
    SplomMarkerSymbol_y_right_open SplomMarkerSymbol = "y-right-open" // y-right-open
    SplomMarkerSymbol41 SplomMarkerSymbol = "41" // 41
    SplomMarkerSymbol_line_ew SplomMarkerSymbol = "line-ew" // line-ew
    SplomMarkerSymbol141 SplomMarkerSymbol = "141" // 141
    SplomMarkerSymbol_line_ew_open SplomMarkerSymbol = "line-ew-open" // line-ew-open
    SplomMarkerSymbol42 SplomMarkerSymbol = "42" // 42
    SplomMarkerSymbol_line_ns SplomMarkerSymbol = "line-ns" // line-ns
    SplomMarkerSymbol142 SplomMarkerSymbol = "142" // 142
    SplomMarkerSymbol_line_ns_open SplomMarkerSymbol = "line-ns-open" // line-ns-open
    SplomMarkerSymbol43 SplomMarkerSymbol = "43" // 43
    SplomMarkerSymbol_line_ne SplomMarkerSymbol = "line-ne" // line-ne
    SplomMarkerSymbol143 SplomMarkerSymbol = "143" // 143
    SplomMarkerSymbol_line_ne_open SplomMarkerSymbol = "line-ne-open" // line-ne-open
    SplomMarkerSymbol44 SplomMarkerSymbol = "44" // 44
    SplomMarkerSymbol_line_nw SplomMarkerSymbol = "line-nw" // line-nw
    SplomMarkerSymbol144 SplomMarkerSymbol = "144" // 144
    SplomMarkerSymbol_line_nw_open SplomMarkerSymbol = "line-nw-open" // line-nw-open
)

// SplomVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type SplomVisible interface{}

const (
    SplomVisibleTrue bool = true
    SplomVisibleFalse bool = false
    SplomVisibleLegendonly string = "legendonly"
)

// StreamtubeColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type StreamtubeColorbarExponentformat string

const (
    StreamtubeColorbarExponentformat_none StreamtubeColorbarExponentformat = "none" // none
    StreamtubeColorbarExponentformat_e StreamtubeColorbarExponentformat = "e" // e
    StreamtubeColorbarExponentformat_E StreamtubeColorbarExponentformat = "E" // E
    StreamtubeColorbarExponentformat_power StreamtubeColorbarExponentformat = "power" // power
    StreamtubeColorbarExponentformat_SI StreamtubeColorbarExponentformat = "SI" // SI
    StreamtubeColorbarExponentformat_B StreamtubeColorbarExponentformat = "B" // B
)

// StreamtubeColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type StreamtubeColorbarLenmode string

const (
    StreamtubeColorbarLenmode_fraction StreamtubeColorbarLenmode = "fraction" // fraction
    StreamtubeColorbarLenmode_pixels StreamtubeColorbarLenmode = "pixels" // pixels
)

// StreamtubeColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type StreamtubeColorbarShowexponent string

const (
    StreamtubeColorbarShowexponent_all StreamtubeColorbarShowexponent = "all" // all
    StreamtubeColorbarShowexponent_first StreamtubeColorbarShowexponent = "first" // first
    StreamtubeColorbarShowexponent_last StreamtubeColorbarShowexponent = "last" // last
    StreamtubeColorbarShowexponent_none StreamtubeColorbarShowexponent = "none" // none
)

// StreamtubeColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type StreamtubeColorbarShowtickprefix string

const (
    StreamtubeColorbarShowtickprefix_all StreamtubeColorbarShowtickprefix = "all" // all
    StreamtubeColorbarShowtickprefix_first StreamtubeColorbarShowtickprefix = "first" // first
    StreamtubeColorbarShowtickprefix_last StreamtubeColorbarShowtickprefix = "last" // last
    StreamtubeColorbarShowtickprefix_none StreamtubeColorbarShowtickprefix = "none" // none
)

// StreamtubeColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type StreamtubeColorbarShowticksuffix string

const (
    StreamtubeColorbarShowticksuffix_all StreamtubeColorbarShowticksuffix = "all" // all
    StreamtubeColorbarShowticksuffix_first StreamtubeColorbarShowticksuffix = "first" // first
    StreamtubeColorbarShowticksuffix_last StreamtubeColorbarShowticksuffix = "last" // last
    StreamtubeColorbarShowticksuffix_none StreamtubeColorbarShowticksuffix = "none" // none
)

// StreamtubeColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type StreamtubeColorbarThicknessmode string

const (
    StreamtubeColorbarThicknessmode_fraction StreamtubeColorbarThicknessmode = "fraction" // fraction
    StreamtubeColorbarThicknessmode_pixels StreamtubeColorbarThicknessmode = "pixels" // pixels
)

// StreamtubeColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type StreamtubeColorbarTickmode string

const (
    StreamtubeColorbarTickmode_auto StreamtubeColorbarTickmode = "auto" // auto
    StreamtubeColorbarTickmode_linear StreamtubeColorbarTickmode = "linear" // linear
    StreamtubeColorbarTickmode_array StreamtubeColorbarTickmode = "array" // array
)

// StreamtubeColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type StreamtubeColorbarTicks string

const (
    StreamtubeColorbarTicks_outside StreamtubeColorbarTicks = "outside" // outside
    StreamtubeColorbarTicks_inside StreamtubeColorbarTicks = "inside" // inside
)

// StreamtubeColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type StreamtubeColorbarTitleSide string

const (
    StreamtubeColorbarTitleSide_right StreamtubeColorbarTitleSide = "right" // right
    StreamtubeColorbarTitleSide_top StreamtubeColorbarTitleSide = "top" // top
    StreamtubeColorbarTitleSide_bottom StreamtubeColorbarTitleSide = "bottom" // bottom
)

// StreamtubeColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type StreamtubeColorbarXanchor string

const (
    StreamtubeColorbarXanchor_left StreamtubeColorbarXanchor = "left" // left
    StreamtubeColorbarXanchor_center StreamtubeColorbarXanchor = "center" // center
    StreamtubeColorbarXanchor_right StreamtubeColorbarXanchor = "right" // right
)

// StreamtubeColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type StreamtubeColorbarYanchor string

const (
    StreamtubeColorbarYanchor_top StreamtubeColorbarYanchor = "top" // top
    StreamtubeColorbarYanchor_middle StreamtubeColorbarYanchor = "middle" // middle
    StreamtubeColorbarYanchor_bottom StreamtubeColorbarYanchor = "bottom" // bottom
)

// StreamtubeHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type StreamtubeHoverlabelAlign string

const (
    StreamtubeHoverlabelAlign_left StreamtubeHoverlabelAlign = "left" // left
    StreamtubeHoverlabelAlign_right StreamtubeHoverlabelAlign = "right" // right
    StreamtubeHoverlabelAlign_auto StreamtubeHoverlabelAlign = "auto" // auto
)

// StreamtubeVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type StreamtubeVisible interface{}

const (
    StreamtubeVisibleTrue bool = true
    StreamtubeVisibleFalse bool = false
    StreamtubeVisibleLegendonly string = "legendonly"
)

// SunburstBranchvalues Determines how the items in `values` are summed. When set to *total*, items in `values` are taken to be value of all its descendants. When set to *remainder*, items in `values` corresponding to the root and the branches sectors are taken to be the extra part not part of the sum of the values at their leaves.
type SunburstBranchvalues string

const (
    SunburstBranchvalues_remainder SunburstBranchvalues = "remainder" // remainder
    SunburstBranchvalues_total SunburstBranchvalues = "total" // total
)

// SunburstHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SunburstHoverlabelAlign string

const (
    SunburstHoverlabelAlign_left SunburstHoverlabelAlign = "left" // left
    SunburstHoverlabelAlign_right SunburstHoverlabelAlign = "right" // right
    SunburstHoverlabelAlign_auto SunburstHoverlabelAlign = "auto" // auto
)

// SunburstInsidetextorientation Controls the orientation of the text inside chart sectors. When set to *auto*, text may be oriented in any direction in order to be as big as possible in the middle of a sector. The *horizontal* option orients text to be parallel with the bottom of the chart, and may make text smaller in order to achieve that goal. The *radial* option orients text along the radius of the sector. The *tangential* option orients text perpendicular to the radius of the sector.
type SunburstInsidetextorientation string

const (
    SunburstInsidetextorientation_horizontal SunburstInsidetextorientation = "horizontal" // horizontal
    SunburstInsidetextorientation_radial SunburstInsidetextorientation = "radial" // radial
    SunburstInsidetextorientation_tangential SunburstInsidetextorientation = "tangential" // tangential
    SunburstInsidetextorientation_auto SunburstInsidetextorientation = "auto" // auto
)

// SunburstMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type SunburstMarkerColorbarExponentformat string

const (
    SunburstMarkerColorbarExponentformat_none SunburstMarkerColorbarExponentformat = "none" // none
    SunburstMarkerColorbarExponentformat_e SunburstMarkerColorbarExponentformat = "e" // e
    SunburstMarkerColorbarExponentformat_E SunburstMarkerColorbarExponentformat = "E" // E
    SunburstMarkerColorbarExponentformat_power SunburstMarkerColorbarExponentformat = "power" // power
    SunburstMarkerColorbarExponentformat_SI SunburstMarkerColorbarExponentformat = "SI" // SI
    SunburstMarkerColorbarExponentformat_B SunburstMarkerColorbarExponentformat = "B" // B
)

// SunburstMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type SunburstMarkerColorbarLenmode string

const (
    SunburstMarkerColorbarLenmode_fraction SunburstMarkerColorbarLenmode = "fraction" // fraction
    SunburstMarkerColorbarLenmode_pixels SunburstMarkerColorbarLenmode = "pixels" // pixels
)

// SunburstMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type SunburstMarkerColorbarShowexponent string

const (
    SunburstMarkerColorbarShowexponent_all SunburstMarkerColorbarShowexponent = "all" // all
    SunburstMarkerColorbarShowexponent_first SunburstMarkerColorbarShowexponent = "first" // first
    SunburstMarkerColorbarShowexponent_last SunburstMarkerColorbarShowexponent = "last" // last
    SunburstMarkerColorbarShowexponent_none SunburstMarkerColorbarShowexponent = "none" // none
)

// SunburstMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type SunburstMarkerColorbarShowtickprefix string

const (
    SunburstMarkerColorbarShowtickprefix_all SunburstMarkerColorbarShowtickprefix = "all" // all
    SunburstMarkerColorbarShowtickprefix_first SunburstMarkerColorbarShowtickprefix = "first" // first
    SunburstMarkerColorbarShowtickprefix_last SunburstMarkerColorbarShowtickprefix = "last" // last
    SunburstMarkerColorbarShowtickprefix_none SunburstMarkerColorbarShowtickprefix = "none" // none
)

// SunburstMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type SunburstMarkerColorbarShowticksuffix string

const (
    SunburstMarkerColorbarShowticksuffix_all SunburstMarkerColorbarShowticksuffix = "all" // all
    SunburstMarkerColorbarShowticksuffix_first SunburstMarkerColorbarShowticksuffix = "first" // first
    SunburstMarkerColorbarShowticksuffix_last SunburstMarkerColorbarShowticksuffix = "last" // last
    SunburstMarkerColorbarShowticksuffix_none SunburstMarkerColorbarShowticksuffix = "none" // none
)

// SunburstMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type SunburstMarkerColorbarThicknessmode string

const (
    SunburstMarkerColorbarThicknessmode_fraction SunburstMarkerColorbarThicknessmode = "fraction" // fraction
    SunburstMarkerColorbarThicknessmode_pixels SunburstMarkerColorbarThicknessmode = "pixels" // pixels
)

// SunburstMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type SunburstMarkerColorbarTickmode string

const (
    SunburstMarkerColorbarTickmode_auto SunburstMarkerColorbarTickmode = "auto" // auto
    SunburstMarkerColorbarTickmode_linear SunburstMarkerColorbarTickmode = "linear" // linear
    SunburstMarkerColorbarTickmode_array SunburstMarkerColorbarTickmode = "array" // array
)

// SunburstMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type SunburstMarkerColorbarTicks string

const (
    SunburstMarkerColorbarTicks_outside SunburstMarkerColorbarTicks = "outside" // outside
    SunburstMarkerColorbarTicks_inside SunburstMarkerColorbarTicks = "inside" // inside
)

// SunburstMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type SunburstMarkerColorbarTitleSide string

const (
    SunburstMarkerColorbarTitleSide_right SunburstMarkerColorbarTitleSide = "right" // right
    SunburstMarkerColorbarTitleSide_top SunburstMarkerColorbarTitleSide = "top" // top
    SunburstMarkerColorbarTitleSide_bottom SunburstMarkerColorbarTitleSide = "bottom" // bottom
)

// SunburstMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type SunburstMarkerColorbarXanchor string

const (
    SunburstMarkerColorbarXanchor_left SunburstMarkerColorbarXanchor = "left" // left
    SunburstMarkerColorbarXanchor_center SunburstMarkerColorbarXanchor = "center" // center
    SunburstMarkerColorbarXanchor_right SunburstMarkerColorbarXanchor = "right" // right
)

// SunburstMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type SunburstMarkerColorbarYanchor string

const (
    SunburstMarkerColorbarYanchor_top SunburstMarkerColorbarYanchor = "top" // top
    SunburstMarkerColorbarYanchor_middle SunburstMarkerColorbarYanchor = "middle" // middle
    SunburstMarkerColorbarYanchor_bottom SunburstMarkerColorbarYanchor = "bottom" // bottom
)

// SunburstVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type SunburstVisible interface{}

const (
    SunburstVisibleTrue bool = true
    SunburstVisibleFalse bool = false
    SunburstVisibleLegendonly string = "legendonly"
)

// SurfaceColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type SurfaceColorbarExponentformat string

const (
    SurfaceColorbarExponentformat_none SurfaceColorbarExponentformat = "none" // none
    SurfaceColorbarExponentformat_e SurfaceColorbarExponentformat = "e" // e
    SurfaceColorbarExponentformat_E SurfaceColorbarExponentformat = "E" // E
    SurfaceColorbarExponentformat_power SurfaceColorbarExponentformat = "power" // power
    SurfaceColorbarExponentformat_SI SurfaceColorbarExponentformat = "SI" // SI
    SurfaceColorbarExponentformat_B SurfaceColorbarExponentformat = "B" // B
)

// SurfaceColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type SurfaceColorbarLenmode string

const (
    SurfaceColorbarLenmode_fraction SurfaceColorbarLenmode = "fraction" // fraction
    SurfaceColorbarLenmode_pixels SurfaceColorbarLenmode = "pixels" // pixels
)

// SurfaceColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type SurfaceColorbarShowexponent string

const (
    SurfaceColorbarShowexponent_all SurfaceColorbarShowexponent = "all" // all
    SurfaceColorbarShowexponent_first SurfaceColorbarShowexponent = "first" // first
    SurfaceColorbarShowexponent_last SurfaceColorbarShowexponent = "last" // last
    SurfaceColorbarShowexponent_none SurfaceColorbarShowexponent = "none" // none
)

// SurfaceColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type SurfaceColorbarShowtickprefix string

const (
    SurfaceColorbarShowtickprefix_all SurfaceColorbarShowtickprefix = "all" // all
    SurfaceColorbarShowtickprefix_first SurfaceColorbarShowtickprefix = "first" // first
    SurfaceColorbarShowtickprefix_last SurfaceColorbarShowtickprefix = "last" // last
    SurfaceColorbarShowtickprefix_none SurfaceColorbarShowtickprefix = "none" // none
)

// SurfaceColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type SurfaceColorbarShowticksuffix string

const (
    SurfaceColorbarShowticksuffix_all SurfaceColorbarShowticksuffix = "all" // all
    SurfaceColorbarShowticksuffix_first SurfaceColorbarShowticksuffix = "first" // first
    SurfaceColorbarShowticksuffix_last SurfaceColorbarShowticksuffix = "last" // last
    SurfaceColorbarShowticksuffix_none SurfaceColorbarShowticksuffix = "none" // none
)

// SurfaceColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type SurfaceColorbarThicknessmode string

const (
    SurfaceColorbarThicknessmode_fraction SurfaceColorbarThicknessmode = "fraction" // fraction
    SurfaceColorbarThicknessmode_pixels SurfaceColorbarThicknessmode = "pixels" // pixels
)

// SurfaceColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type SurfaceColorbarTickmode string

const (
    SurfaceColorbarTickmode_auto SurfaceColorbarTickmode = "auto" // auto
    SurfaceColorbarTickmode_linear SurfaceColorbarTickmode = "linear" // linear
    SurfaceColorbarTickmode_array SurfaceColorbarTickmode = "array" // array
)

// SurfaceColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type SurfaceColorbarTicks string

const (
    SurfaceColorbarTicks_outside SurfaceColorbarTicks = "outside" // outside
    SurfaceColorbarTicks_inside SurfaceColorbarTicks = "inside" // inside
)

// SurfaceColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type SurfaceColorbarTitleSide string

const (
    SurfaceColorbarTitleSide_right SurfaceColorbarTitleSide = "right" // right
    SurfaceColorbarTitleSide_top SurfaceColorbarTitleSide = "top" // top
    SurfaceColorbarTitleSide_bottom SurfaceColorbarTitleSide = "bottom" // bottom
)

// SurfaceColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type SurfaceColorbarXanchor string

const (
    SurfaceColorbarXanchor_left SurfaceColorbarXanchor = "left" // left
    SurfaceColorbarXanchor_center SurfaceColorbarXanchor = "center" // center
    SurfaceColorbarXanchor_right SurfaceColorbarXanchor = "right" // right
)

// SurfaceColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type SurfaceColorbarYanchor string

const (
    SurfaceColorbarYanchor_top SurfaceColorbarYanchor = "top" // top
    SurfaceColorbarYanchor_middle SurfaceColorbarYanchor = "middle" // middle
    SurfaceColorbarYanchor_bottom SurfaceColorbarYanchor = "bottom" // bottom
)

// SurfaceHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type SurfaceHoverlabelAlign string

const (
    SurfaceHoverlabelAlign_left SurfaceHoverlabelAlign = "left" // left
    SurfaceHoverlabelAlign_right SurfaceHoverlabelAlign = "right" // right
    SurfaceHoverlabelAlign_auto SurfaceHoverlabelAlign = "auto" // auto
)

// SurfaceVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type SurfaceVisible interface{}

const (
    SurfaceVisibleTrue bool = true
    SurfaceVisibleFalse bool = false
    SurfaceVisibleLegendonly string = "legendonly"
)

// SurfaceXcalendar Sets the calendar system to use with `x` date data.
type SurfaceXcalendar string

const (
    SurfaceXcalendar_gregorian SurfaceXcalendar = "gregorian" // gregorian
    SurfaceXcalendar_chinese SurfaceXcalendar = "chinese" // chinese
    SurfaceXcalendar_coptic SurfaceXcalendar = "coptic" // coptic
    SurfaceXcalendar_discworld SurfaceXcalendar = "discworld" // discworld
    SurfaceXcalendar_ethiopian SurfaceXcalendar = "ethiopian" // ethiopian
    SurfaceXcalendar_hebrew SurfaceXcalendar = "hebrew" // hebrew
    SurfaceXcalendar_islamic SurfaceXcalendar = "islamic" // islamic
    SurfaceXcalendar_julian SurfaceXcalendar = "julian" // julian
    SurfaceXcalendar_mayan SurfaceXcalendar = "mayan" // mayan
    SurfaceXcalendar_nanakshahi SurfaceXcalendar = "nanakshahi" // nanakshahi
    SurfaceXcalendar_nepali SurfaceXcalendar = "nepali" // nepali
    SurfaceXcalendar_persian SurfaceXcalendar = "persian" // persian
    SurfaceXcalendar_jalali SurfaceXcalendar = "jalali" // jalali
    SurfaceXcalendar_taiwan SurfaceXcalendar = "taiwan" // taiwan
    SurfaceXcalendar_thai SurfaceXcalendar = "thai" // thai
    SurfaceXcalendar_ummalqura SurfaceXcalendar = "ummalqura" // ummalqura
)

// SurfaceYcalendar Sets the calendar system to use with `y` date data.
type SurfaceYcalendar string

const (
    SurfaceYcalendar_gregorian SurfaceYcalendar = "gregorian" // gregorian
    SurfaceYcalendar_chinese SurfaceYcalendar = "chinese" // chinese
    SurfaceYcalendar_coptic SurfaceYcalendar = "coptic" // coptic
    SurfaceYcalendar_discworld SurfaceYcalendar = "discworld" // discworld
    SurfaceYcalendar_ethiopian SurfaceYcalendar = "ethiopian" // ethiopian
    SurfaceYcalendar_hebrew SurfaceYcalendar = "hebrew" // hebrew
    SurfaceYcalendar_islamic SurfaceYcalendar = "islamic" // islamic
    SurfaceYcalendar_julian SurfaceYcalendar = "julian" // julian
    SurfaceYcalendar_mayan SurfaceYcalendar = "mayan" // mayan
    SurfaceYcalendar_nanakshahi SurfaceYcalendar = "nanakshahi" // nanakshahi
    SurfaceYcalendar_nepali SurfaceYcalendar = "nepali" // nepali
    SurfaceYcalendar_persian SurfaceYcalendar = "persian" // persian
    SurfaceYcalendar_jalali SurfaceYcalendar = "jalali" // jalali
    SurfaceYcalendar_taiwan SurfaceYcalendar = "taiwan" // taiwan
    SurfaceYcalendar_thai SurfaceYcalendar = "thai" // thai
    SurfaceYcalendar_ummalqura SurfaceYcalendar = "ummalqura" // ummalqura
)

// SurfaceZcalendar Sets the calendar system to use with `z` date data.
type SurfaceZcalendar string

const (
    SurfaceZcalendar_gregorian SurfaceZcalendar = "gregorian" // gregorian
    SurfaceZcalendar_chinese SurfaceZcalendar = "chinese" // chinese
    SurfaceZcalendar_coptic SurfaceZcalendar = "coptic" // coptic
    SurfaceZcalendar_discworld SurfaceZcalendar = "discworld" // discworld
    SurfaceZcalendar_ethiopian SurfaceZcalendar = "ethiopian" // ethiopian
    SurfaceZcalendar_hebrew SurfaceZcalendar = "hebrew" // hebrew
    SurfaceZcalendar_islamic SurfaceZcalendar = "islamic" // islamic
    SurfaceZcalendar_julian SurfaceZcalendar = "julian" // julian
    SurfaceZcalendar_mayan SurfaceZcalendar = "mayan" // mayan
    SurfaceZcalendar_nanakshahi SurfaceZcalendar = "nanakshahi" // nanakshahi
    SurfaceZcalendar_nepali SurfaceZcalendar = "nepali" // nepali
    SurfaceZcalendar_persian SurfaceZcalendar = "persian" // persian
    SurfaceZcalendar_jalali SurfaceZcalendar = "jalali" // jalali
    SurfaceZcalendar_taiwan SurfaceZcalendar = "taiwan" // taiwan
    SurfaceZcalendar_thai SurfaceZcalendar = "thai" // thai
    SurfaceZcalendar_ummalqura SurfaceZcalendar = "ummalqura" // ummalqura
)

// TableCellsAlign Sets the horizontal alignment of the `text` within the box. Has an effect only if `text` spans two or more lines (i.e. `text` contains one or more <br> HTML tags) or if an explicit width is set to override the text width.
type TableCellsAlign string

const (
    TableCellsAlign_left TableCellsAlign = "left" // left
    TableCellsAlign_center TableCellsAlign = "center" // center
    TableCellsAlign_right TableCellsAlign = "right" // right
)

// TableHeaderAlign Sets the horizontal alignment of the `text` within the box. Has an effect only if `text` spans two or more lines (i.e. `text` contains one or more <br> HTML tags) or if an explicit width is set to override the text width.
type TableHeaderAlign string

const (
    TableHeaderAlign_left TableHeaderAlign = "left" // left
    TableHeaderAlign_center TableHeaderAlign = "center" // center
    TableHeaderAlign_right TableHeaderAlign = "right" // right
)

// TableHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type TableHoverlabelAlign string

const (
    TableHoverlabelAlign_left TableHoverlabelAlign = "left" // left
    TableHoverlabelAlign_right TableHoverlabelAlign = "right" // right
    TableHoverlabelAlign_auto TableHoverlabelAlign = "auto" // auto
)

// TableVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type TableVisible interface{}

const (
    TableVisibleTrue bool = true
    TableVisibleFalse bool = false
    TableVisibleLegendonly string = "legendonly"
)

// TreemapBranchvalues Determines how the items in `values` are summed. When set to *total*, items in `values` are taken to be value of all its descendants. When set to *remainder*, items in `values` corresponding to the root and the branches sectors are taken to be the extra part not part of the sum of the values at their leaves.
type TreemapBranchvalues string

const (
    TreemapBranchvalues_remainder TreemapBranchvalues = "remainder" // remainder
    TreemapBranchvalues_total TreemapBranchvalues = "total" // total
)

// TreemapHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type TreemapHoverlabelAlign string

const (
    TreemapHoverlabelAlign_left TreemapHoverlabelAlign = "left" // left
    TreemapHoverlabelAlign_right TreemapHoverlabelAlign = "right" // right
    TreemapHoverlabelAlign_auto TreemapHoverlabelAlign = "auto" // auto
)

// TreemapMarkerColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type TreemapMarkerColorbarExponentformat string

const (
    TreemapMarkerColorbarExponentformat_none TreemapMarkerColorbarExponentformat = "none" // none
    TreemapMarkerColorbarExponentformat_e TreemapMarkerColorbarExponentformat = "e" // e
    TreemapMarkerColorbarExponentformat_E TreemapMarkerColorbarExponentformat = "E" // E
    TreemapMarkerColorbarExponentformat_power TreemapMarkerColorbarExponentformat = "power" // power
    TreemapMarkerColorbarExponentformat_SI TreemapMarkerColorbarExponentformat = "SI" // SI
    TreemapMarkerColorbarExponentformat_B TreemapMarkerColorbarExponentformat = "B" // B
)

// TreemapMarkerColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type TreemapMarkerColorbarLenmode string

const (
    TreemapMarkerColorbarLenmode_fraction TreemapMarkerColorbarLenmode = "fraction" // fraction
    TreemapMarkerColorbarLenmode_pixels TreemapMarkerColorbarLenmode = "pixels" // pixels
)

// TreemapMarkerColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type TreemapMarkerColorbarShowexponent string

const (
    TreemapMarkerColorbarShowexponent_all TreemapMarkerColorbarShowexponent = "all" // all
    TreemapMarkerColorbarShowexponent_first TreemapMarkerColorbarShowexponent = "first" // first
    TreemapMarkerColorbarShowexponent_last TreemapMarkerColorbarShowexponent = "last" // last
    TreemapMarkerColorbarShowexponent_none TreemapMarkerColorbarShowexponent = "none" // none
)

// TreemapMarkerColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type TreemapMarkerColorbarShowtickprefix string

const (
    TreemapMarkerColorbarShowtickprefix_all TreemapMarkerColorbarShowtickprefix = "all" // all
    TreemapMarkerColorbarShowtickprefix_first TreemapMarkerColorbarShowtickprefix = "first" // first
    TreemapMarkerColorbarShowtickprefix_last TreemapMarkerColorbarShowtickprefix = "last" // last
    TreemapMarkerColorbarShowtickprefix_none TreemapMarkerColorbarShowtickprefix = "none" // none
)

// TreemapMarkerColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type TreemapMarkerColorbarShowticksuffix string

const (
    TreemapMarkerColorbarShowticksuffix_all TreemapMarkerColorbarShowticksuffix = "all" // all
    TreemapMarkerColorbarShowticksuffix_first TreemapMarkerColorbarShowticksuffix = "first" // first
    TreemapMarkerColorbarShowticksuffix_last TreemapMarkerColorbarShowticksuffix = "last" // last
    TreemapMarkerColorbarShowticksuffix_none TreemapMarkerColorbarShowticksuffix = "none" // none
)

// TreemapMarkerColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type TreemapMarkerColorbarThicknessmode string

const (
    TreemapMarkerColorbarThicknessmode_fraction TreemapMarkerColorbarThicknessmode = "fraction" // fraction
    TreemapMarkerColorbarThicknessmode_pixels TreemapMarkerColorbarThicknessmode = "pixels" // pixels
)

// TreemapMarkerColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type TreemapMarkerColorbarTickmode string

const (
    TreemapMarkerColorbarTickmode_auto TreemapMarkerColorbarTickmode = "auto" // auto
    TreemapMarkerColorbarTickmode_linear TreemapMarkerColorbarTickmode = "linear" // linear
    TreemapMarkerColorbarTickmode_array TreemapMarkerColorbarTickmode = "array" // array
)

// TreemapMarkerColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type TreemapMarkerColorbarTicks string

const (
    TreemapMarkerColorbarTicks_outside TreemapMarkerColorbarTicks = "outside" // outside
    TreemapMarkerColorbarTicks_inside TreemapMarkerColorbarTicks = "inside" // inside
)

// TreemapMarkerColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type TreemapMarkerColorbarTitleSide string

const (
    TreemapMarkerColorbarTitleSide_right TreemapMarkerColorbarTitleSide = "right" // right
    TreemapMarkerColorbarTitleSide_top TreemapMarkerColorbarTitleSide = "top" // top
    TreemapMarkerColorbarTitleSide_bottom TreemapMarkerColorbarTitleSide = "bottom" // bottom
)

// TreemapMarkerColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type TreemapMarkerColorbarXanchor string

const (
    TreemapMarkerColorbarXanchor_left TreemapMarkerColorbarXanchor = "left" // left
    TreemapMarkerColorbarXanchor_center TreemapMarkerColorbarXanchor = "center" // center
    TreemapMarkerColorbarXanchor_right TreemapMarkerColorbarXanchor = "right" // right
)

// TreemapMarkerColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type TreemapMarkerColorbarYanchor string

const (
    TreemapMarkerColorbarYanchor_top TreemapMarkerColorbarYanchor = "top" // top
    TreemapMarkerColorbarYanchor_middle TreemapMarkerColorbarYanchor = "middle" // middle
    TreemapMarkerColorbarYanchor_bottom TreemapMarkerColorbarYanchor = "bottom" // bottom
)

// TreemapMarkerDepthfade Determines if the sector colors are faded towards the background from the leaves up to the headers. This option is unavailable when a `colorscale` is present, defaults to false when `marker.colors` is set, but otherwise defaults to true. When set to *reversed*, the fading direction is inverted, that is the top elements within hierarchy are drawn with fully saturated colors while the leaves are faded towards the background color.
type TreemapMarkerDepthfade interface{}

const (
    TreemapMarkerDepthfadeTrue bool = true
    TreemapMarkerDepthfadeFalse bool = false
    TreemapMarkerDepthfadeReversed string = "reversed"
)

// TreemapPathbarEdgeshape Determines which shape is used for edges between `barpath` labels.
type TreemapPathbarEdgeshape string

const (
    TreemapPathbarEdgeshape_gt TreemapPathbarEdgeshape = ">" // >
    TreemapPathbarEdgeshape_lt TreemapPathbarEdgeshape = "<" // <
    TreemapPathbarEdgeshape_or TreemapPathbarEdgeshape = "|" // |
    TreemapPathbarEdgeshape_slash TreemapPathbarEdgeshape = "/" // /
    TreemapPathbarEdgeshape_doublebackslash TreemapPathbarEdgeshape = "\\" // \
)

// TreemapPathbarSide Determines on which side of the the treemap the `pathbar` should be presented.
type TreemapPathbarSide string

const (
    TreemapPathbarSide_top TreemapPathbarSide = "top" // top
    TreemapPathbarSide_bottom TreemapPathbarSide = "bottom" // bottom
)

// TreemapTextposition Sets the positions of the `text` elements.
type TreemapTextposition string

const (
    TreemapTextposition_topleft TreemapTextposition = "top left" // top left
    TreemapTextposition_topcenter TreemapTextposition = "top center" // top center
    TreemapTextposition_topright TreemapTextposition = "top right" // top right
    TreemapTextposition_middleleft TreemapTextposition = "middle left" // middle left
    TreemapTextposition_middlecenter TreemapTextposition = "middle center" // middle center
    TreemapTextposition_middleright TreemapTextposition = "middle right" // middle right
    TreemapTextposition_bottomleft TreemapTextposition = "bottom left" // bottom left
    TreemapTextposition_bottomcenter TreemapTextposition = "bottom center" // bottom center
    TreemapTextposition_bottomright TreemapTextposition = "bottom right" // bottom right
)

// TreemapTilingPacking Determines d3 treemap solver. For more info please refer to https://github.com/d3/d3-hierarchy#treemap-tiling
type TreemapTilingPacking string

const (
    TreemapTilingPacking_squarify TreemapTilingPacking = "squarify" // squarify
    TreemapTilingPacking_binary TreemapTilingPacking = "binary" // binary
    TreemapTilingPacking_dice TreemapTilingPacking = "dice" // dice
    TreemapTilingPacking_slice TreemapTilingPacking = "slice" // slice
    TreemapTilingPacking_slice_dice TreemapTilingPacking = "slice-dice" // slice-dice
    TreemapTilingPacking_dice_slice TreemapTilingPacking = "dice-slice" // dice-slice
)

// TreemapVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type TreemapVisible interface{}

const (
    TreemapVisibleTrue bool = true
    TreemapVisibleFalse bool = false
    TreemapVisibleLegendonly string = "legendonly"
)

// ViolinHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type ViolinHoverlabelAlign string

const (
    ViolinHoverlabelAlign_left ViolinHoverlabelAlign = "left" // left
    ViolinHoverlabelAlign_right ViolinHoverlabelAlign = "right" // right
    ViolinHoverlabelAlign_auto ViolinHoverlabelAlign = "auto" // auto
)

// ViolinMarkerSymbol Sets the marker symbol type. Adding 100 is equivalent to appending *-open* to a symbol name. Adding 200 is equivalent to appending *-dot* to a symbol name. Adding 300 is equivalent to appending *-open-dot* or *dot-open* to a symbol name.
type ViolinMarkerSymbol string

const (
    ViolinMarkerSymbol0 ViolinMarkerSymbol = "0" // 0
    ViolinMarkerSymbol_circle ViolinMarkerSymbol = "circle" // circle
    ViolinMarkerSymbol100 ViolinMarkerSymbol = "100" // 100
    ViolinMarkerSymbol_circle_open ViolinMarkerSymbol = "circle-open" // circle-open
    ViolinMarkerSymbol200 ViolinMarkerSymbol = "200" // 200
    ViolinMarkerSymbol_circle_dot ViolinMarkerSymbol = "circle-dot" // circle-dot
    ViolinMarkerSymbol300 ViolinMarkerSymbol = "300" // 300
    ViolinMarkerSymbol_circle_open_dot ViolinMarkerSymbol = "circle-open-dot" // circle-open-dot
    ViolinMarkerSymbol1 ViolinMarkerSymbol = "1" // 1
    ViolinMarkerSymbol_square ViolinMarkerSymbol = "square" // square
    ViolinMarkerSymbol101 ViolinMarkerSymbol = "101" // 101
    ViolinMarkerSymbol_square_open ViolinMarkerSymbol = "square-open" // square-open
    ViolinMarkerSymbol201 ViolinMarkerSymbol = "201" // 201
    ViolinMarkerSymbol_square_dot ViolinMarkerSymbol = "square-dot" // square-dot
    ViolinMarkerSymbol301 ViolinMarkerSymbol = "301" // 301
    ViolinMarkerSymbol_square_open_dot ViolinMarkerSymbol = "square-open-dot" // square-open-dot
    ViolinMarkerSymbol2 ViolinMarkerSymbol = "2" // 2
    ViolinMarkerSymbol_diamond ViolinMarkerSymbol = "diamond" // diamond
    ViolinMarkerSymbol102 ViolinMarkerSymbol = "102" // 102
    ViolinMarkerSymbol_diamond_open ViolinMarkerSymbol = "diamond-open" // diamond-open
    ViolinMarkerSymbol202 ViolinMarkerSymbol = "202" // 202
    ViolinMarkerSymbol_diamond_dot ViolinMarkerSymbol = "diamond-dot" // diamond-dot
    ViolinMarkerSymbol302 ViolinMarkerSymbol = "302" // 302
    ViolinMarkerSymbol_diamond_open_dot ViolinMarkerSymbol = "diamond-open-dot" // diamond-open-dot
    ViolinMarkerSymbol3 ViolinMarkerSymbol = "3" // 3
    ViolinMarkerSymbol_cross ViolinMarkerSymbol = "cross" // cross
    ViolinMarkerSymbol103 ViolinMarkerSymbol = "103" // 103
    ViolinMarkerSymbol_cross_open ViolinMarkerSymbol = "cross-open" // cross-open
    ViolinMarkerSymbol203 ViolinMarkerSymbol = "203" // 203
    ViolinMarkerSymbol_cross_dot ViolinMarkerSymbol = "cross-dot" // cross-dot
    ViolinMarkerSymbol303 ViolinMarkerSymbol = "303" // 303
    ViolinMarkerSymbol_cross_open_dot ViolinMarkerSymbol = "cross-open-dot" // cross-open-dot
    ViolinMarkerSymbol4 ViolinMarkerSymbol = "4" // 4
    ViolinMarkerSymbol_x ViolinMarkerSymbol = "x" // x
    ViolinMarkerSymbol104 ViolinMarkerSymbol = "104" // 104
    ViolinMarkerSymbol_x_open ViolinMarkerSymbol = "x-open" // x-open
    ViolinMarkerSymbol204 ViolinMarkerSymbol = "204" // 204
    ViolinMarkerSymbol_x_dot ViolinMarkerSymbol = "x-dot" // x-dot
    ViolinMarkerSymbol304 ViolinMarkerSymbol = "304" // 304
    ViolinMarkerSymbol_x_open_dot ViolinMarkerSymbol = "x-open-dot" // x-open-dot
    ViolinMarkerSymbol5 ViolinMarkerSymbol = "5" // 5
    ViolinMarkerSymbol_triangle_up ViolinMarkerSymbol = "triangle-up" // triangle-up
    ViolinMarkerSymbol105 ViolinMarkerSymbol = "105" // 105
    ViolinMarkerSymbol_triangle_up_open ViolinMarkerSymbol = "triangle-up-open" // triangle-up-open
    ViolinMarkerSymbol205 ViolinMarkerSymbol = "205" // 205
    ViolinMarkerSymbol_triangle_up_dot ViolinMarkerSymbol = "triangle-up-dot" // triangle-up-dot
    ViolinMarkerSymbol305 ViolinMarkerSymbol = "305" // 305
    ViolinMarkerSymbol_triangle_up_open_dot ViolinMarkerSymbol = "triangle-up-open-dot" // triangle-up-open-dot
    ViolinMarkerSymbol6 ViolinMarkerSymbol = "6" // 6
    ViolinMarkerSymbol_triangle_down ViolinMarkerSymbol = "triangle-down" // triangle-down
    ViolinMarkerSymbol106 ViolinMarkerSymbol = "106" // 106
    ViolinMarkerSymbol_triangle_down_open ViolinMarkerSymbol = "triangle-down-open" // triangle-down-open
    ViolinMarkerSymbol206 ViolinMarkerSymbol = "206" // 206
    ViolinMarkerSymbol_triangle_down_dot ViolinMarkerSymbol = "triangle-down-dot" // triangle-down-dot
    ViolinMarkerSymbol306 ViolinMarkerSymbol = "306" // 306
    ViolinMarkerSymbol_triangle_down_open_dot ViolinMarkerSymbol = "triangle-down-open-dot" // triangle-down-open-dot
    ViolinMarkerSymbol7 ViolinMarkerSymbol = "7" // 7
    ViolinMarkerSymbol_triangle_left ViolinMarkerSymbol = "triangle-left" // triangle-left
    ViolinMarkerSymbol107 ViolinMarkerSymbol = "107" // 107
    ViolinMarkerSymbol_triangle_left_open ViolinMarkerSymbol = "triangle-left-open" // triangle-left-open
    ViolinMarkerSymbol207 ViolinMarkerSymbol = "207" // 207
    ViolinMarkerSymbol_triangle_left_dot ViolinMarkerSymbol = "triangle-left-dot" // triangle-left-dot
    ViolinMarkerSymbol307 ViolinMarkerSymbol = "307" // 307
    ViolinMarkerSymbol_triangle_left_open_dot ViolinMarkerSymbol = "triangle-left-open-dot" // triangle-left-open-dot
    ViolinMarkerSymbol8 ViolinMarkerSymbol = "8" // 8
    ViolinMarkerSymbol_triangle_right ViolinMarkerSymbol = "triangle-right" // triangle-right
    ViolinMarkerSymbol108 ViolinMarkerSymbol = "108" // 108
    ViolinMarkerSymbol_triangle_right_open ViolinMarkerSymbol = "triangle-right-open" // triangle-right-open
    ViolinMarkerSymbol208 ViolinMarkerSymbol = "208" // 208
    ViolinMarkerSymbol_triangle_right_dot ViolinMarkerSymbol = "triangle-right-dot" // triangle-right-dot
    ViolinMarkerSymbol308 ViolinMarkerSymbol = "308" // 308
    ViolinMarkerSymbol_triangle_right_open_dot ViolinMarkerSymbol = "triangle-right-open-dot" // triangle-right-open-dot
    ViolinMarkerSymbol9 ViolinMarkerSymbol = "9" // 9
    ViolinMarkerSymbol_triangle_ne ViolinMarkerSymbol = "triangle-ne" // triangle-ne
    ViolinMarkerSymbol109 ViolinMarkerSymbol = "109" // 109
    ViolinMarkerSymbol_triangle_ne_open ViolinMarkerSymbol = "triangle-ne-open" // triangle-ne-open
    ViolinMarkerSymbol209 ViolinMarkerSymbol = "209" // 209
    ViolinMarkerSymbol_triangle_ne_dot ViolinMarkerSymbol = "triangle-ne-dot" // triangle-ne-dot
    ViolinMarkerSymbol309 ViolinMarkerSymbol = "309" // 309
    ViolinMarkerSymbol_triangle_ne_open_dot ViolinMarkerSymbol = "triangle-ne-open-dot" // triangle-ne-open-dot
    ViolinMarkerSymbol10 ViolinMarkerSymbol = "10" // 10
    ViolinMarkerSymbol_triangle_se ViolinMarkerSymbol = "triangle-se" // triangle-se
    ViolinMarkerSymbol110 ViolinMarkerSymbol = "110" // 110
    ViolinMarkerSymbol_triangle_se_open ViolinMarkerSymbol = "triangle-se-open" // triangle-se-open
    ViolinMarkerSymbol210 ViolinMarkerSymbol = "210" // 210
    ViolinMarkerSymbol_triangle_se_dot ViolinMarkerSymbol = "triangle-se-dot" // triangle-se-dot
    ViolinMarkerSymbol310 ViolinMarkerSymbol = "310" // 310
    ViolinMarkerSymbol_triangle_se_open_dot ViolinMarkerSymbol = "triangle-se-open-dot" // triangle-se-open-dot
    ViolinMarkerSymbol11 ViolinMarkerSymbol = "11" // 11
    ViolinMarkerSymbol_triangle_sw ViolinMarkerSymbol = "triangle-sw" // triangle-sw
    ViolinMarkerSymbol111 ViolinMarkerSymbol = "111" // 111
    ViolinMarkerSymbol_triangle_sw_open ViolinMarkerSymbol = "triangle-sw-open" // triangle-sw-open
    ViolinMarkerSymbol211 ViolinMarkerSymbol = "211" // 211
    ViolinMarkerSymbol_triangle_sw_dot ViolinMarkerSymbol = "triangle-sw-dot" // triangle-sw-dot
    ViolinMarkerSymbol311 ViolinMarkerSymbol = "311" // 311
    ViolinMarkerSymbol_triangle_sw_open_dot ViolinMarkerSymbol = "triangle-sw-open-dot" // triangle-sw-open-dot
    ViolinMarkerSymbol12 ViolinMarkerSymbol = "12" // 12
    ViolinMarkerSymbol_triangle_nw ViolinMarkerSymbol = "triangle-nw" // triangle-nw
    ViolinMarkerSymbol112 ViolinMarkerSymbol = "112" // 112
    ViolinMarkerSymbol_triangle_nw_open ViolinMarkerSymbol = "triangle-nw-open" // triangle-nw-open
    ViolinMarkerSymbol212 ViolinMarkerSymbol = "212" // 212
    ViolinMarkerSymbol_triangle_nw_dot ViolinMarkerSymbol = "triangle-nw-dot" // triangle-nw-dot
    ViolinMarkerSymbol312 ViolinMarkerSymbol = "312" // 312
    ViolinMarkerSymbol_triangle_nw_open_dot ViolinMarkerSymbol = "triangle-nw-open-dot" // triangle-nw-open-dot
    ViolinMarkerSymbol13 ViolinMarkerSymbol = "13" // 13
    ViolinMarkerSymbol_pentagon ViolinMarkerSymbol = "pentagon" // pentagon
    ViolinMarkerSymbol113 ViolinMarkerSymbol = "113" // 113
    ViolinMarkerSymbol_pentagon_open ViolinMarkerSymbol = "pentagon-open" // pentagon-open
    ViolinMarkerSymbol213 ViolinMarkerSymbol = "213" // 213
    ViolinMarkerSymbol_pentagon_dot ViolinMarkerSymbol = "pentagon-dot" // pentagon-dot
    ViolinMarkerSymbol313 ViolinMarkerSymbol = "313" // 313
    ViolinMarkerSymbol_pentagon_open_dot ViolinMarkerSymbol = "pentagon-open-dot" // pentagon-open-dot
    ViolinMarkerSymbol14 ViolinMarkerSymbol = "14" // 14
    ViolinMarkerSymbol_hexagon ViolinMarkerSymbol = "hexagon" // hexagon
    ViolinMarkerSymbol114 ViolinMarkerSymbol = "114" // 114
    ViolinMarkerSymbol_hexagon_open ViolinMarkerSymbol = "hexagon-open" // hexagon-open
    ViolinMarkerSymbol214 ViolinMarkerSymbol = "214" // 214
    ViolinMarkerSymbol_hexagon_dot ViolinMarkerSymbol = "hexagon-dot" // hexagon-dot
    ViolinMarkerSymbol314 ViolinMarkerSymbol = "314" // 314
    ViolinMarkerSymbol_hexagon_open_dot ViolinMarkerSymbol = "hexagon-open-dot" // hexagon-open-dot
    ViolinMarkerSymbol15 ViolinMarkerSymbol = "15" // 15
    ViolinMarkerSymbol_hexagon2 ViolinMarkerSymbol = "hexagon2" // hexagon2
    ViolinMarkerSymbol115 ViolinMarkerSymbol = "115" // 115
    ViolinMarkerSymbol_hexagon2_open ViolinMarkerSymbol = "hexagon2-open" // hexagon2-open
    ViolinMarkerSymbol215 ViolinMarkerSymbol = "215" // 215
    ViolinMarkerSymbol_hexagon2_dot ViolinMarkerSymbol = "hexagon2-dot" // hexagon2-dot
    ViolinMarkerSymbol315 ViolinMarkerSymbol = "315" // 315
    ViolinMarkerSymbol_hexagon2_open_dot ViolinMarkerSymbol = "hexagon2-open-dot" // hexagon2-open-dot
    ViolinMarkerSymbol16 ViolinMarkerSymbol = "16" // 16
    ViolinMarkerSymbol_octagon ViolinMarkerSymbol = "octagon" // octagon
    ViolinMarkerSymbol116 ViolinMarkerSymbol = "116" // 116
    ViolinMarkerSymbol_octagon_open ViolinMarkerSymbol = "octagon-open" // octagon-open
    ViolinMarkerSymbol216 ViolinMarkerSymbol = "216" // 216
    ViolinMarkerSymbol_octagon_dot ViolinMarkerSymbol = "octagon-dot" // octagon-dot
    ViolinMarkerSymbol316 ViolinMarkerSymbol = "316" // 316
    ViolinMarkerSymbol_octagon_open_dot ViolinMarkerSymbol = "octagon-open-dot" // octagon-open-dot
    ViolinMarkerSymbol17 ViolinMarkerSymbol = "17" // 17
    ViolinMarkerSymbol_star ViolinMarkerSymbol = "star" // star
    ViolinMarkerSymbol117 ViolinMarkerSymbol = "117" // 117
    ViolinMarkerSymbol_star_open ViolinMarkerSymbol = "star-open" // star-open
    ViolinMarkerSymbol217 ViolinMarkerSymbol = "217" // 217
    ViolinMarkerSymbol_star_dot ViolinMarkerSymbol = "star-dot" // star-dot
    ViolinMarkerSymbol317 ViolinMarkerSymbol = "317" // 317
    ViolinMarkerSymbol_star_open_dot ViolinMarkerSymbol = "star-open-dot" // star-open-dot
    ViolinMarkerSymbol18 ViolinMarkerSymbol = "18" // 18
    ViolinMarkerSymbol_hexagram ViolinMarkerSymbol = "hexagram" // hexagram
    ViolinMarkerSymbol118 ViolinMarkerSymbol = "118" // 118
    ViolinMarkerSymbol_hexagram_open ViolinMarkerSymbol = "hexagram-open" // hexagram-open
    ViolinMarkerSymbol218 ViolinMarkerSymbol = "218" // 218
    ViolinMarkerSymbol_hexagram_dot ViolinMarkerSymbol = "hexagram-dot" // hexagram-dot
    ViolinMarkerSymbol318 ViolinMarkerSymbol = "318" // 318
    ViolinMarkerSymbol_hexagram_open_dot ViolinMarkerSymbol = "hexagram-open-dot" // hexagram-open-dot
    ViolinMarkerSymbol19 ViolinMarkerSymbol = "19" // 19
    ViolinMarkerSymbol_star_triangle_up ViolinMarkerSymbol = "star-triangle-up" // star-triangle-up
    ViolinMarkerSymbol119 ViolinMarkerSymbol = "119" // 119
    ViolinMarkerSymbol_star_triangle_up_open ViolinMarkerSymbol = "star-triangle-up-open" // star-triangle-up-open
    ViolinMarkerSymbol219 ViolinMarkerSymbol = "219" // 219
    ViolinMarkerSymbol_star_triangle_up_dot ViolinMarkerSymbol = "star-triangle-up-dot" // star-triangle-up-dot
    ViolinMarkerSymbol319 ViolinMarkerSymbol = "319" // 319
    ViolinMarkerSymbol_star_triangle_up_open_dot ViolinMarkerSymbol = "star-triangle-up-open-dot" // star-triangle-up-open-dot
    ViolinMarkerSymbol20 ViolinMarkerSymbol = "20" // 20
    ViolinMarkerSymbol_star_triangle_down ViolinMarkerSymbol = "star-triangle-down" // star-triangle-down
    ViolinMarkerSymbol120 ViolinMarkerSymbol = "120" // 120
    ViolinMarkerSymbol_star_triangle_down_open ViolinMarkerSymbol = "star-triangle-down-open" // star-triangle-down-open
    ViolinMarkerSymbol220 ViolinMarkerSymbol = "220" // 220
    ViolinMarkerSymbol_star_triangle_down_dot ViolinMarkerSymbol = "star-triangle-down-dot" // star-triangle-down-dot
    ViolinMarkerSymbol320 ViolinMarkerSymbol = "320" // 320
    ViolinMarkerSymbol_star_triangle_down_open_dot ViolinMarkerSymbol = "star-triangle-down-open-dot" // star-triangle-down-open-dot
    ViolinMarkerSymbol21 ViolinMarkerSymbol = "21" // 21
    ViolinMarkerSymbol_star_square ViolinMarkerSymbol = "star-square" // star-square
    ViolinMarkerSymbol121 ViolinMarkerSymbol = "121" // 121
    ViolinMarkerSymbol_star_square_open ViolinMarkerSymbol = "star-square-open" // star-square-open
    ViolinMarkerSymbol221 ViolinMarkerSymbol = "221" // 221
    ViolinMarkerSymbol_star_square_dot ViolinMarkerSymbol = "star-square-dot" // star-square-dot
    ViolinMarkerSymbol321 ViolinMarkerSymbol = "321" // 321
    ViolinMarkerSymbol_star_square_open_dot ViolinMarkerSymbol = "star-square-open-dot" // star-square-open-dot
    ViolinMarkerSymbol22 ViolinMarkerSymbol = "22" // 22
    ViolinMarkerSymbol_star_diamond ViolinMarkerSymbol = "star-diamond" // star-diamond
    ViolinMarkerSymbol122 ViolinMarkerSymbol = "122" // 122
    ViolinMarkerSymbol_star_diamond_open ViolinMarkerSymbol = "star-diamond-open" // star-diamond-open
    ViolinMarkerSymbol222 ViolinMarkerSymbol = "222" // 222
    ViolinMarkerSymbol_star_diamond_dot ViolinMarkerSymbol = "star-diamond-dot" // star-diamond-dot
    ViolinMarkerSymbol322 ViolinMarkerSymbol = "322" // 322
    ViolinMarkerSymbol_star_diamond_open_dot ViolinMarkerSymbol = "star-diamond-open-dot" // star-diamond-open-dot
    ViolinMarkerSymbol23 ViolinMarkerSymbol = "23" // 23
    ViolinMarkerSymbol_diamond_tall ViolinMarkerSymbol = "diamond-tall" // diamond-tall
    ViolinMarkerSymbol123 ViolinMarkerSymbol = "123" // 123
    ViolinMarkerSymbol_diamond_tall_open ViolinMarkerSymbol = "diamond-tall-open" // diamond-tall-open
    ViolinMarkerSymbol223 ViolinMarkerSymbol = "223" // 223
    ViolinMarkerSymbol_diamond_tall_dot ViolinMarkerSymbol = "diamond-tall-dot" // diamond-tall-dot
    ViolinMarkerSymbol323 ViolinMarkerSymbol = "323" // 323
    ViolinMarkerSymbol_diamond_tall_open_dot ViolinMarkerSymbol = "diamond-tall-open-dot" // diamond-tall-open-dot
    ViolinMarkerSymbol24 ViolinMarkerSymbol = "24" // 24
    ViolinMarkerSymbol_diamond_wide ViolinMarkerSymbol = "diamond-wide" // diamond-wide
    ViolinMarkerSymbol124 ViolinMarkerSymbol = "124" // 124
    ViolinMarkerSymbol_diamond_wide_open ViolinMarkerSymbol = "diamond-wide-open" // diamond-wide-open
    ViolinMarkerSymbol224 ViolinMarkerSymbol = "224" // 224
    ViolinMarkerSymbol_diamond_wide_dot ViolinMarkerSymbol = "diamond-wide-dot" // diamond-wide-dot
    ViolinMarkerSymbol324 ViolinMarkerSymbol = "324" // 324
    ViolinMarkerSymbol_diamond_wide_open_dot ViolinMarkerSymbol = "diamond-wide-open-dot" // diamond-wide-open-dot
    ViolinMarkerSymbol25 ViolinMarkerSymbol = "25" // 25
    ViolinMarkerSymbol_hourglass ViolinMarkerSymbol = "hourglass" // hourglass
    ViolinMarkerSymbol125 ViolinMarkerSymbol = "125" // 125
    ViolinMarkerSymbol_hourglass_open ViolinMarkerSymbol = "hourglass-open" // hourglass-open
    ViolinMarkerSymbol26 ViolinMarkerSymbol = "26" // 26
    ViolinMarkerSymbol_bowtie ViolinMarkerSymbol = "bowtie" // bowtie
    ViolinMarkerSymbol126 ViolinMarkerSymbol = "126" // 126
    ViolinMarkerSymbol_bowtie_open ViolinMarkerSymbol = "bowtie-open" // bowtie-open
    ViolinMarkerSymbol27 ViolinMarkerSymbol = "27" // 27
    ViolinMarkerSymbol_circle_cross ViolinMarkerSymbol = "circle-cross" // circle-cross
    ViolinMarkerSymbol127 ViolinMarkerSymbol = "127" // 127
    ViolinMarkerSymbol_circle_cross_open ViolinMarkerSymbol = "circle-cross-open" // circle-cross-open
    ViolinMarkerSymbol28 ViolinMarkerSymbol = "28" // 28
    ViolinMarkerSymbol_circle_x ViolinMarkerSymbol = "circle-x" // circle-x
    ViolinMarkerSymbol128 ViolinMarkerSymbol = "128" // 128
    ViolinMarkerSymbol_circle_x_open ViolinMarkerSymbol = "circle-x-open" // circle-x-open
    ViolinMarkerSymbol29 ViolinMarkerSymbol = "29" // 29
    ViolinMarkerSymbol_square_cross ViolinMarkerSymbol = "square-cross" // square-cross
    ViolinMarkerSymbol129 ViolinMarkerSymbol = "129" // 129
    ViolinMarkerSymbol_square_cross_open ViolinMarkerSymbol = "square-cross-open" // square-cross-open
    ViolinMarkerSymbol30 ViolinMarkerSymbol = "30" // 30
    ViolinMarkerSymbol_square_x ViolinMarkerSymbol = "square-x" // square-x
    ViolinMarkerSymbol130 ViolinMarkerSymbol = "130" // 130
    ViolinMarkerSymbol_square_x_open ViolinMarkerSymbol = "square-x-open" // square-x-open
    ViolinMarkerSymbol31 ViolinMarkerSymbol = "31" // 31
    ViolinMarkerSymbol_diamond_cross ViolinMarkerSymbol = "diamond-cross" // diamond-cross
    ViolinMarkerSymbol131 ViolinMarkerSymbol = "131" // 131
    ViolinMarkerSymbol_diamond_cross_open ViolinMarkerSymbol = "diamond-cross-open" // diamond-cross-open
    ViolinMarkerSymbol32 ViolinMarkerSymbol = "32" // 32
    ViolinMarkerSymbol_diamond_x ViolinMarkerSymbol = "diamond-x" // diamond-x
    ViolinMarkerSymbol132 ViolinMarkerSymbol = "132" // 132
    ViolinMarkerSymbol_diamond_x_open ViolinMarkerSymbol = "diamond-x-open" // diamond-x-open
    ViolinMarkerSymbol33 ViolinMarkerSymbol = "33" // 33
    ViolinMarkerSymbol_cross_thin ViolinMarkerSymbol = "cross-thin" // cross-thin
    ViolinMarkerSymbol133 ViolinMarkerSymbol = "133" // 133
    ViolinMarkerSymbol_cross_thin_open ViolinMarkerSymbol = "cross-thin-open" // cross-thin-open
    ViolinMarkerSymbol34 ViolinMarkerSymbol = "34" // 34
    ViolinMarkerSymbol_x_thin ViolinMarkerSymbol = "x-thin" // x-thin
    ViolinMarkerSymbol134 ViolinMarkerSymbol = "134" // 134
    ViolinMarkerSymbol_x_thin_open ViolinMarkerSymbol = "x-thin-open" // x-thin-open
    ViolinMarkerSymbol35 ViolinMarkerSymbol = "35" // 35
    ViolinMarkerSymbol_asterisk ViolinMarkerSymbol = "asterisk" // asterisk
    ViolinMarkerSymbol135 ViolinMarkerSymbol = "135" // 135
    ViolinMarkerSymbol_asterisk_open ViolinMarkerSymbol = "asterisk-open" // asterisk-open
    ViolinMarkerSymbol36 ViolinMarkerSymbol = "36" // 36
    ViolinMarkerSymbol_hash ViolinMarkerSymbol = "hash" // hash
    ViolinMarkerSymbol136 ViolinMarkerSymbol = "136" // 136
    ViolinMarkerSymbol_hash_open ViolinMarkerSymbol = "hash-open" // hash-open
    ViolinMarkerSymbol236 ViolinMarkerSymbol = "236" // 236
    ViolinMarkerSymbol_hash_dot ViolinMarkerSymbol = "hash-dot" // hash-dot
    ViolinMarkerSymbol336 ViolinMarkerSymbol = "336" // 336
    ViolinMarkerSymbol_hash_open_dot ViolinMarkerSymbol = "hash-open-dot" // hash-open-dot
    ViolinMarkerSymbol37 ViolinMarkerSymbol = "37" // 37
    ViolinMarkerSymbol_y_up ViolinMarkerSymbol = "y-up" // y-up
    ViolinMarkerSymbol137 ViolinMarkerSymbol = "137" // 137
    ViolinMarkerSymbol_y_up_open ViolinMarkerSymbol = "y-up-open" // y-up-open
    ViolinMarkerSymbol38 ViolinMarkerSymbol = "38" // 38
    ViolinMarkerSymbol_y_down ViolinMarkerSymbol = "y-down" // y-down
    ViolinMarkerSymbol138 ViolinMarkerSymbol = "138" // 138
    ViolinMarkerSymbol_y_down_open ViolinMarkerSymbol = "y-down-open" // y-down-open
    ViolinMarkerSymbol39 ViolinMarkerSymbol = "39" // 39
    ViolinMarkerSymbol_y_left ViolinMarkerSymbol = "y-left" // y-left
    ViolinMarkerSymbol139 ViolinMarkerSymbol = "139" // 139
    ViolinMarkerSymbol_y_left_open ViolinMarkerSymbol = "y-left-open" // y-left-open
    ViolinMarkerSymbol40 ViolinMarkerSymbol = "40" // 40
    ViolinMarkerSymbol_y_right ViolinMarkerSymbol = "y-right" // y-right
    ViolinMarkerSymbol140 ViolinMarkerSymbol = "140" // 140
    ViolinMarkerSymbol_y_right_open ViolinMarkerSymbol = "y-right-open" // y-right-open
    ViolinMarkerSymbol41 ViolinMarkerSymbol = "41" // 41
    ViolinMarkerSymbol_line_ew ViolinMarkerSymbol = "line-ew" // line-ew
    ViolinMarkerSymbol141 ViolinMarkerSymbol = "141" // 141
    ViolinMarkerSymbol_line_ew_open ViolinMarkerSymbol = "line-ew-open" // line-ew-open
    ViolinMarkerSymbol42 ViolinMarkerSymbol = "42" // 42
    ViolinMarkerSymbol_line_ns ViolinMarkerSymbol = "line-ns" // line-ns
    ViolinMarkerSymbol142 ViolinMarkerSymbol = "142" // 142
    ViolinMarkerSymbol_line_ns_open ViolinMarkerSymbol = "line-ns-open" // line-ns-open
    ViolinMarkerSymbol43 ViolinMarkerSymbol = "43" // 43
    ViolinMarkerSymbol_line_ne ViolinMarkerSymbol = "line-ne" // line-ne
    ViolinMarkerSymbol143 ViolinMarkerSymbol = "143" // 143
    ViolinMarkerSymbol_line_ne_open ViolinMarkerSymbol = "line-ne-open" // line-ne-open
    ViolinMarkerSymbol44 ViolinMarkerSymbol = "44" // 44
    ViolinMarkerSymbol_line_nw ViolinMarkerSymbol = "line-nw" // line-nw
    ViolinMarkerSymbol144 ViolinMarkerSymbol = "144" // 144
    ViolinMarkerSymbol_line_nw_open ViolinMarkerSymbol = "line-nw-open" // line-nw-open
)

// ViolinOrientation Sets the orientation of the violin(s). If *v* (*h*), the distribution is visualized along the vertical (horizontal).
type ViolinOrientation string

const (
    ViolinOrientation_v ViolinOrientation = "v" // v
    ViolinOrientation_h ViolinOrientation = "h" // h
)

// ViolinPoints If *outliers*, only the sample points lying outside the whiskers are shown If *suspectedoutliers*, the outlier points are shown and points either less than 4*Q1-3*Q3 or greater than 4*Q3-3*Q1 are highlighted (see `outliercolor`) If *all*, all sample points are shown If *false*, only the violins are shown with no sample points. Defaults to *suspectedoutliers* when `marker.outliercolor` or `marker.line.outliercolor` is set, otherwise defaults to *outliers*.
type ViolinPoints interface{}

const (
    ViolinPointsAll string = "all"
    ViolinPointsOutliers string = "outliers"
    ViolinPointsSuspectedoutliers string = "suspectedoutliers"
    ViolinPointsFalse bool = false
)

// ViolinScalemode Sets the metric by which the width of each violin is determined.*width* means each violin has the same (max) width*count* means the violins are scaled by the number of sample points makingup each violin.
type ViolinScalemode string

const (
    ViolinScalemode_width ViolinScalemode = "width" // width
    ViolinScalemode_count ViolinScalemode = "count" // count
)

// ViolinSide Determines on which side of the position value the density function making up one half of a violin is plotted. Useful when comparing two violin traces under *overlay* mode, where one trace has `side` set to *positive* and the other to *negative*.
type ViolinSide string

const (
    ViolinSide_both ViolinSide = "both" // both
    ViolinSide_positive ViolinSide = "positive" // positive
    ViolinSide_negative ViolinSide = "negative" // negative
)

// ViolinSpanmode Sets the method by which the span in data space where the density function will be computed. *soft* means the span goes from the sample's minimum value minus two bandwidths to the sample's maximum value plus two bandwidths. *hard* means the span goes from the sample's minimum to its maximum value. For custom span settings, use mode *manual* and fill in the `span` attribute.
type ViolinSpanmode string

const (
    ViolinSpanmode_soft ViolinSpanmode = "soft" // soft
    ViolinSpanmode_hard ViolinSpanmode = "hard" // hard
    ViolinSpanmode_manual ViolinSpanmode = "manual" // manual
)

// ViolinVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type ViolinVisible interface{}

const (
    ViolinVisibleTrue bool = true
    ViolinVisibleFalse bool = false
    ViolinVisibleLegendonly string = "legendonly"
)

// VolumeColorbarExponentformat Determines a formatting rule for the tick exponents. For example, consider the number 1,000,000,000. If *none*, it appears as 1,000,000,000. If *e*, 1e+9. If *E*, 1E+9. If *power*, 1x10^9 (with 9 in a super script). If *SI*, 1G. If *B*, 1B.
type VolumeColorbarExponentformat string

const (
    VolumeColorbarExponentformat_none VolumeColorbarExponentformat = "none" // none
    VolumeColorbarExponentformat_e VolumeColorbarExponentformat = "e" // e
    VolumeColorbarExponentformat_E VolumeColorbarExponentformat = "E" // E
    VolumeColorbarExponentformat_power VolumeColorbarExponentformat = "power" // power
    VolumeColorbarExponentformat_SI VolumeColorbarExponentformat = "SI" // SI
    VolumeColorbarExponentformat_B VolumeColorbarExponentformat = "B" // B
)

// VolumeColorbarLenmode Determines whether this color bar's length (i.e. the measure in the color variation direction) is set in units of plot *fraction* or in *pixels. Use `len` to set the value.
type VolumeColorbarLenmode string

const (
    VolumeColorbarLenmode_fraction VolumeColorbarLenmode = "fraction" // fraction
    VolumeColorbarLenmode_pixels VolumeColorbarLenmode = "pixels" // pixels
)

// VolumeColorbarShowexponent If *all*, all exponents are shown besides their significands. If *first*, only the exponent of the first tick is shown. If *last*, only the exponent of the last tick is shown. If *none*, no exponents appear.
type VolumeColorbarShowexponent string

const (
    VolumeColorbarShowexponent_all VolumeColorbarShowexponent = "all" // all
    VolumeColorbarShowexponent_first VolumeColorbarShowexponent = "first" // first
    VolumeColorbarShowexponent_last VolumeColorbarShowexponent = "last" // last
    VolumeColorbarShowexponent_none VolumeColorbarShowexponent = "none" // none
)

// VolumeColorbarShowtickprefix If *all*, all tick labels are displayed with a prefix. If *first*, only the first tick is displayed with a prefix. If *last*, only the last tick is displayed with a suffix. If *none*, tick prefixes are hidden.
type VolumeColorbarShowtickprefix string

const (
    VolumeColorbarShowtickprefix_all VolumeColorbarShowtickprefix = "all" // all
    VolumeColorbarShowtickprefix_first VolumeColorbarShowtickprefix = "first" // first
    VolumeColorbarShowtickprefix_last VolumeColorbarShowtickprefix = "last" // last
    VolumeColorbarShowtickprefix_none VolumeColorbarShowtickprefix = "none" // none
)

// VolumeColorbarShowticksuffix Same as `showtickprefix` but for tick suffixes.
type VolumeColorbarShowticksuffix string

const (
    VolumeColorbarShowticksuffix_all VolumeColorbarShowticksuffix = "all" // all
    VolumeColorbarShowticksuffix_first VolumeColorbarShowticksuffix = "first" // first
    VolumeColorbarShowticksuffix_last VolumeColorbarShowticksuffix = "last" // last
    VolumeColorbarShowticksuffix_none VolumeColorbarShowticksuffix = "none" // none
)

// VolumeColorbarThicknessmode Determines whether this color bar's thickness (i.e. the measure in the constant color direction) is set in units of plot *fraction* or in *pixels*. Use `thickness` to set the value.
type VolumeColorbarThicknessmode string

const (
    VolumeColorbarThicknessmode_fraction VolumeColorbarThicknessmode = "fraction" // fraction
    VolumeColorbarThicknessmode_pixels VolumeColorbarThicknessmode = "pixels" // pixels
)

// VolumeColorbarTickmode Sets the tick mode for this axis. If *auto*, the number of ticks is set via `nticks`. If *linear*, the placement of the ticks is determined by a starting position `tick0` and a tick step `dtick` (*linear* is the default value if `tick0` and `dtick` are provided). If *array*, the placement of the ticks is set via `tickvals` and the tick text is `ticktext`. (*array* is the default value if `tickvals` is provided).
type VolumeColorbarTickmode string

const (
    VolumeColorbarTickmode_auto VolumeColorbarTickmode = "auto" // auto
    VolumeColorbarTickmode_linear VolumeColorbarTickmode = "linear" // linear
    VolumeColorbarTickmode_array VolumeColorbarTickmode = "array" // array
)

// VolumeColorbarTicks Determines whether ticks are drawn or not. If **, this axis' ticks are not drawn. If *outside* (*inside*), this axis' are drawn outside (inside) the axis lines.
type VolumeColorbarTicks string

const (
    VolumeColorbarTicks_outside VolumeColorbarTicks = "outside" // outside
    VolumeColorbarTicks_inside VolumeColorbarTicks = "inside" // inside
)

// VolumeColorbarTitleSide Determines the location of color bar's title with respect to the color bar. Note that the title's location used to be set by the now deprecated `titleside` attribute.
type VolumeColorbarTitleSide string

const (
    VolumeColorbarTitleSide_right VolumeColorbarTitleSide = "right" // right
    VolumeColorbarTitleSide_top VolumeColorbarTitleSide = "top" // top
    VolumeColorbarTitleSide_bottom VolumeColorbarTitleSide = "bottom" // bottom
)

// VolumeColorbarXanchor Sets this color bar's horizontal position anchor. This anchor binds the `x` position to the *left*, *center* or *right* of the color bar.
type VolumeColorbarXanchor string

const (
    VolumeColorbarXanchor_left VolumeColorbarXanchor = "left" // left
    VolumeColorbarXanchor_center VolumeColorbarXanchor = "center" // center
    VolumeColorbarXanchor_right VolumeColorbarXanchor = "right" // right
)

// VolumeColorbarYanchor Sets this color bar's vertical position anchor This anchor binds the `y` position to the *top*, *middle* or *bottom* of the color bar.
type VolumeColorbarYanchor string

const (
    VolumeColorbarYanchor_top VolumeColorbarYanchor = "top" // top
    VolumeColorbarYanchor_middle VolumeColorbarYanchor = "middle" // middle
    VolumeColorbarYanchor_bottom VolumeColorbarYanchor = "bottom" // bottom
)

// VolumeHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type VolumeHoverlabelAlign string

const (
    VolumeHoverlabelAlign_left VolumeHoverlabelAlign = "left" // left
    VolumeHoverlabelAlign_right VolumeHoverlabelAlign = "right" // right
    VolumeHoverlabelAlign_auto VolumeHoverlabelAlign = "auto" // auto
)

// VolumeVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type VolumeVisible interface{}

const (
    VolumeVisibleTrue bool = true
    VolumeVisibleFalse bool = false
    VolumeVisibleLegendonly string = "legendonly"
)

// WaterfallConnectorMode Sets the shape of connector lines.
type WaterfallConnectorMode string

const (
    WaterfallConnectorMode_spanning WaterfallConnectorMode = "spanning" // spanning
    WaterfallConnectorMode_between WaterfallConnectorMode = "between" // between
)

// WaterfallConstraintext Constrain the size of text inside or outside a bar to be no larger than the bar itself.
type WaterfallConstraintext string

const (
    WaterfallConstraintext_inside WaterfallConstraintext = "inside" // inside
    WaterfallConstraintext_outside WaterfallConstraintext = "outside" // outside
    WaterfallConstraintext_both WaterfallConstraintext = "both" // both
    WaterfallConstraintext_none WaterfallConstraintext = "none" // none
)

// WaterfallHoverlabelAlign Sets the horizontal alignment of the text content within hover label box. Has an effect only if the hover label text spans more two or more lines
type WaterfallHoverlabelAlign string

const (
    WaterfallHoverlabelAlign_left WaterfallHoverlabelAlign = "left" // left
    WaterfallHoverlabelAlign_right WaterfallHoverlabelAlign = "right" // right
    WaterfallHoverlabelAlign_auto WaterfallHoverlabelAlign = "auto" // auto
)

// WaterfallInsidetextanchor Determines if texts are kept at center or start/end points in `textposition` *inside* mode.
type WaterfallInsidetextanchor string

const (
    WaterfallInsidetextanchor_end WaterfallInsidetextanchor = "end" // end
    WaterfallInsidetextanchor_middle WaterfallInsidetextanchor = "middle" // middle
    WaterfallInsidetextanchor_start WaterfallInsidetextanchor = "start" // start
)

// WaterfallOrientation Sets the orientation of the bars. With *v* (*h*), the value of the each bar spans along the vertical (horizontal).
type WaterfallOrientation string

const (
    WaterfallOrientation_v WaterfallOrientation = "v" // v
    WaterfallOrientation_h WaterfallOrientation = "h" // h
)

// WaterfallTextposition Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside.
type WaterfallTextposition string

const (
    WaterfallTextposition_inside WaterfallTextposition = "inside" // inside
    WaterfallTextposition_outside WaterfallTextposition = "outside" // outside
    WaterfallTextposition_auto WaterfallTextposition = "auto" // auto
    WaterfallTextposition_none WaterfallTextposition = "none" // none
)

// WaterfallVisible Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
type WaterfallVisible interface{}

const (
    WaterfallVisibleTrue bool = true
    WaterfallVisibleFalse bool = false
    WaterfallVisibleLegendonly string = "legendonly"
)
