package relate

import (
	"testing"
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
	"time"
)

func TestSidedness(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("homotopic sidedness test", func() {
		g.It("should test sidedness to a context geometry", func() {
			g.Timeout(1 * time.Hour)
			var cwkt = "POLYGON (( 278 307, 270 298, 274 286, 279 272, 301 274, 308 288, 311 304, 296 308, 278 307 ))"
			var wkt = "LINESTRING ( 155 171, 207 166, 253 175, 317 171, 367 182, 400 200, 428 249, 417 291, 383 324, 361 332, 333 347, 314 357, 257 383, 204 370, 176 337, 180 305, 214 295, 244 302, 281 332, 316 328, 331 306, 332 291, 315 265, 285 250, 247 261, 231 276, 195 264, 187 230, 216 215, 257 226, 273 217, 273 205, 240 197, 200 200, 178 193, 157 226, 156 246, 151 263, 120 264, 95 249, 89 261, 100 300, 116 359, 139 389, 172 413, 211 425, 256 430, 289 431, 348 427 )"
			var coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsTrue()

			cwkt = "POLYGON (( 221 347, 205 334, 221 322, 234 324, 237 342, 221 347 ))"

			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsFalse()

			cwkt = "POLYGON (( 221 347, 205 334, 221 322, 234 324, 237 342, 221 347 ))"
			wkt = "LINESTRING ( 155 171, 207 166, 253 175, 317 171, 366.3526316090967 196.8011373660995, 403.24889098928804 217.07380735521562, 449.06512516469047 258.02460073323016, 469.7432485535889 310.7335427049321, 431.6306289740506 348.8461622844704, 393.9234627942946 354.9279632812052, 371.21807240648457 309.92263590536743, 372.43443260583155 253.97006673540696, 415.4124929827577 249.91553273758373, 425.14337457753345 316.00443690210227, 404.465251188635 331.81711949361284, 359.0544704130149 294.10995331385686, 333 347, 365.94717820931436 412.50234605029505, 350.9454024173684 441.6949908346222, 269.44926906112164 441.6949908346222, 233.36391648049494 444.9386180328808, 205.38763189551466 402.77146445551926, 204 370, 176 337, 180 305, 214 295, 244 302, 281 332, 316 328, 331 306, 332 291, 315 265, 285 250, 247 261, 231 276, 195 264, 187 230, 216 215, 257 226, 273 217, 273 205, 240 197, 200 200, 178 193, 157 226, 156 246, 151 263, 120 264, 95 249, 89 261, 100 300, 116 359, 139 389, 172 413, 211 425, 256 430, 289 431, 348 427 )"
			coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsFalse()
			cwkt = "POLYGON (( 278 307, 270 298, 274 286, 279 272, 301 274, 308 288, 311 304, 296 308, 278 307 ))"
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsTrue()

			wkt = "LINESTRING ( 155 171, 207 166, 253 175, 317 171, 366.3526316090967 196.8011373660995, 403.24889098928804 217.07380735521562, 449.2399093677868 286.1290060157034, 451.783228702019 325.7276851722446, 415.8591984439269 355.71930167053, 418.6880722510337 402.11283210708103, 389.8335594185446 420.2176244725644, 350.9454024173684 441.6949908346222, 269.44926906112164 441.6949908346222, 233.36391648049494 444.9386180328808, 205.38763189551466 402.77146445551926, 204 370, 176 337, 180 305, 157 226, 95 249, 89 261, 100 300, 116 359, 139 389, 172 413, 211 425, 256 430, 289 431, 348 427 )"
			cwkt = "POLYGON (( 278 307, 270 298, 274 286, 279 272, 301 274, 308 288, 311 304, 296 308, 278 307 ))"
			coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsFalse()
			cwkt = "POLYGON (( 221 347, 205 334, 221 322, 234 324, 237 342, 221 347 ))"
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsTrue()

			coords = []*geom.Point{{2, 2}, {5, 2}, {7, 2}, {9, 2}}
			var lnconst = geom.NewLineString([]*geom.Point{{3, 2}, {6, 2}, {6.5, 2}})
			g.Assert(
				Homotopy(coords, lnconst),
			).IsTrue()
			var plyconst = geom.NewPolygon([]*geom.Point{{4, 1}, {4, 3}, {5, 3}, {5, 1}})
			g.Assert(
				Homotopy(coords, plyconst),
			).IsTrue()

			wkt = "LINESTRING ( 155 171, 207 166, 253 175, 317 171, 367 182, 400 200, 428 249, 417 291, 383 324, 361 332, 333 347, 314 357, 257 383, 204 370, 176 337, 180 305, 214 295, 244 302, 281 332, 316 328, 331 306, 332 291, 315 265, 285 250, 247 261, 231 276, 175.9815455950541 265.3822256568778, 187 230, 233.38568778979908 233.9203400309118, 265.9514992272025 234.4723029366305, 274.230942812983 207.97808346213282, 248.54273709311704 199.72323889171014, 215.17091190108192 192.52312210200918, 219.58661514683152 241.09585780525492, 203.57969088098918 278.07737248840795, 222.34642967542504 355.9041421947449, 211 425, 256 430, 289 431, 348 427 )"
			cwkt = "POLYGON (( 278 307, 268.711313755796 290.77251931993817, 235.45711654915712 296.5278817518552, 232.8337248840804 286.90877897990725, 277.40587928436713 278.46806414206173, 285.4987949850666 282.69941864135717, 288.58197836166926 290.77251931993817, 290.2378670788254 299.05196290571865, 278 307 ))"
			coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsFalse()

			wkt = "LINESTRING ( 155 171, 207 166, 253 175, 317 171, 367 182, 400 200, 428 249, 417 291, 383 324, 361 332, 333 347, 314 357, 257 383, 204 370, 176 337, 180 305, 205.2355795981453 296.29214837712516, 244 302, 281 332, 316 328, 331 306, 332 291, 315 265, 285 250, 247 261, 231 276, 175.9815455950541 265.3822256568778, 187 230, 233.38568778979908 233.9203400309118, 265.9514992272025 234.4723029366305, 274.230942812983 207.97808346213282, 248.54273709311704 199.72323889171014, 215.17091190108192 192.52312210200918, 219.58661514683152 241.09585780525492, 203.57969088098918 278.07737248840795, 225.65820710973728 312.29907264296753, 246.63279752704796 324.99421947449764, 255.46420401854718 354.24825347758883, 223.45035548686246 391.22976816074186, 233.93765069551782 421.58772797527047, 256 430, 289 431, 348 427 )"
			cwkt = "POLYGON (( 221 347, 205 334, 221 322, 234 324, 237 342, 221 347 ))"
			coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsFalse()

			wkt = "LINESTRING ( 155 171, 207 166, 253 175, 317 171, 367 182, 400 200, 428 249, 417 291, 383 324, 361 332, 333 347, 314 357, 257 383, 204 370, 176 337, 180 305, 205.2355795981453 296.29214837712516, 244 302, 281 332, 316 328, 331 306, 332 291, 315 265, 285 250, 247 261, 231 276, 175.9815455950541 265.3822256568778, 187 230, 233.38568778979908 233.9203400309118, 265.9514992272025 234.4723029366305, 274.230942812983 207.97808346213282, 248.54273709311704 199.72323889171014, 215.17091190108192 192.52312210200918, 219.58661514683152 241.09585780525492, 203.57969088098918 278.07737248840795, 225.65820710973728 312.29907264296753, 246.63279752704796 324.99421947449764, 255.46420401854718 354.24825347758883, 223.45035548686246 391.22976816074186, 233.93765069551782 421.58772797527047, 256 430, 289 431, 348 427 )"
			cwkt = "POLYGON (( 232.00298146975462 210.52416263821522, 232.00298146975462 220.68605060410567, 246.51996427816957 220.68605060410567, 246.51996427816957 210.52416263821522, 232.00298146975462 210.52416263821522 ))"
			coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsTrue()

			wkt = "LINESTRING ( 155 171, 207 166, 253 175, 317 171, 367 182, 400 200, 428 249, 417 291, 383 324, 361 332, 333 347, 314 357, 257 383, 204 370, 176 337, 180 305, 205.2355795981453 296.29214837712516, 244 302, 281 332, 316 328, 331 306, 332 291, 315 265, 285 250, 247 261, 231 276, 175.9815455950541 265.3822256568778, 187 230, 233.38568778979908 233.9203400309118, 265.9514992272025 234.4723029366305, 274.230942812983 207.97808346213282, 248.54273709311704 199.72323889171014, 215.17091190108192 192.52312210200918, 219.58661514683152 241.09585780525492, 203.57969088098918 278.07737248840795, 225.65820710973728 312.29907264296753, 246.63279752704796 324.99421947449764, 255.46420401854718 354.24825347758883, 223.45035548686246 391.22976816074186, 233.93765069551782 421.58772797527047, 256 430, 289 431, 348 427 )"
			cwkt = "POLYGON (( 211.34885547628429 235.3767011676454, 211.34885547628429 241.48063878287113, 213.69652378983264 241.48063878287113, 213.69652378983264 235.3767011676454, 211.34885547628429 235.3767011676454 ))"
			coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsTrue()

			wkt = "LINESTRING ( 155 171, 207 166, 253 175, 317 171, 367 182, 400 200, 428 249, 417 291, 383 324, 361 332, 333 347, 314 357, 257 383, 204 370, 176 337, 180 305, 205.2355795981453 296.29214837712516, 244 302, 281 332, 316 328, 331 306, 332 291, 315 265, 285 250, 247 261, 231 276, 175.9815455950541 265.3822256568778, 187 230, 233.38568778979908 233.9203400309118, 265.9514992272025 234.4723029366305, 274.230942812983 207.97808346213282, 248.54273709311704 199.72323889171014, 215.17091190108192 192.52312210200918, 203.26715129630935 235.02274990598545, 211.59588065682914 246.07018366916196, 217.04481846253088 253.29779029226893, 222.32389329119033 260.30008643805553, 226.53649692090931 265.8877886619315, 203.57969088098918 278.07737248840795, 225.65820710973728 312.29907264296753, 246.63279752704796 324.99421947449764, 255.46420401854718 354.24825347758883, 223.45035548686246 391.22976816074186, 233.93765069551782 421.58772797527047, 256 430, 289 431, 348 427 )"
			cwkt = "POLYGON (( 211.34885547628429 235.3767011676454, 211.34885547628429 241.48063878287113, 213.69652378983264 241.48063878287113, 213.69652378983264 235.3767011676454, 211.34885547628429 235.3767011676454 ))"
			coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsFalse()

			wkt = "LINESTRING ( 155 171, 207 166, 253 175, 317 171, 367 182, 400 200, 428 249, 417 291, 383 324, 361 332, 333 347, 314 357, 257 383, 204 370, 176 337, 180 305, 205.2355795981453 296.29214837712516, 244 302, 281 332, 316 328, 331 306, 332 291, 315 265, 285 250, 247 261, 231 276, 175.9815455950541 265.3822256568778, 187 230, 233.38568778979908 233.9203400309118, 265.9514992272025 234.4723029366305, 274.230942812983 207.97808346213282, 248.54273709311704 199.72323889171014, 215.17091190108192 192.52312210200918, 203.26715129630935 235.02274990598545, 211.59588065682914 246.07018366916196, 217.04481846253088 253.29779029226893, 222.32389329119033 260.30008643805553, 226.53649692090931 265.8877886619315, 348 427 )"
			cwkt = "POLYGON (( 211.34885547628429 235.3767011676454, 211.34885547628429 241.48063878287113, 213.69652378983264 241.48063878287113, 213.69652378983264 235.3767011676454, 211.34885547628429 235.3767011676454 ))"
			coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsFalse()

			wkt = "LINESTRING ( 155 171, 169.0211940381559 189.5980604858441, 178.36905451652774 201.99729511000572, 189.61729454361603 216.9172404309104, 155.33584906460158 235.98971111300912, 171.9249231022847 253.04608301090866, 184.81805947947288 266.3024063282993, 199.21035949720948 260.0405937778234, 216.46768369629606 252.53226438472433, 226.72184529093568 266.1336393496349, 253.77293550266072 261.8918612097899, 286.55244753078927 256.75183112348185, 298.89980243113513 277.3307559573916, 308.9755371420998 294.1236471423328, 311.6569923092747 322.2091261672002, 274.3587106022468 329.32036224961234, 286.70927127817396 345.7024530943654, 302.96002907406313 367.25786239875725, 265.3750851201071 385.89221758862226, 267.0360547209449 409.9762768007707, 292.87995239405126 407.8640164833156, 311.8206399330813 406.3159658301186, 333.85824574639133 408.2420254459906, 348 427 )"
			cwkt = "POLYGON (( 211.34885547628429 235.3767011676454, 211.34885547628429 241.48063878287113, 213.69652378983264 241.48063878287113, 213.69652378983264 235.3767011676454, 211.34885547628429 235.3767011676454 ))"
			coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsTrue()

			wkt = "LINESTRING ( 90 270, 120 270, 150 270, 190 220, 150 230, 180 250, 180 270, 200 250, 210 270, 230 260, 250 270, 280 270, 300 270, 310 290, 310 320, 400 270, 380 270, 340 270, 380 310, 410 270, 450 310, 430 270, 470 270 )"
			cwkt = "POLYGON (( 211.34885547628429 235.3767011676454, 211.34885547628429 241.48063878287113, 213.69652378983264 241.48063878287113, 213.69652378983264 235.3767011676454, 211.34885547628429 235.3767011676454 ))"
			coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsTrue()

			wkt = "LINESTRING ( 90 270, 470 270 )"
			cwkt = "POLYGON (( 211.34885547628429 235.3767011676454, 211.34885547628429 241.48063878287113, 213.69652378983264 241.48063878287113, 213.69652378983264 235.3767011676454, 211.34885547628429 235.3767011676454 ))"
			coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsTrue()

			wkt = "LINESTRING ( 90 270, 470 270 )"
			cwkt = "POLYGON (( 370 260, 370 300, 380 300, 380 260, 370 260 ))"
			coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			g.Assert(
				Homotopy(coords, geom.NewPolygonFromWKT(cwkt)),
			).IsTrue()

		})

	})

}

func TestSidness2(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("context neighbours", func() {
		g.It("should test context neighbours", func() {
			wkt := "LINESTRING ( 400 560, 520 660, 620 600, 720 640, 760 500, 680 460, 720 420, 780 400, 720 360, 580 340, 440 400, 380 440, 440 500, 400 560 )"
			ctx := geom.NewLineStringFromWKT(wkt)

			wkt_a := "LINESTRING ( 980 280, 981.3546694489717 324.82716460080917, 1006.1665140838959 336.10527579850196, 1026.4671142397428 362.0449315531954, 1017.4446252815886 410.5408097032744, 982.4824805687409 427.45797649981364, 947.5203358558933 470.31479905104624, 948.6481469756626 585.3515332675128, 921.5806801011998 611.2911890222063, 906.9191355441992 613.5468112617448, 843.7617128371196 666.5539338909009, 734.3640342194994 743.245090035212, 700 750, 650 750, 465.94498771441084 773.6959902689825, 412.93786508525466 739.8616566759041, 385.87039821079196 720.6888676398264, 361.0585535758678 717.3054342805185, 346.29722623487015 749.9852989139027, 315.5582431228852 768.0047028071353, 275.2795755968359 771.1845976118235, 246.660522354643 707.5867015180614, 265.00252427153964 657.6454550927426, 220.85197491460565 627.9662944103014, 156.8861009121442 574.9681901708519, 118.727363255887 505.01050446771376, 143.10655675849577 442.47257330884776, 122.96722299547113 361.9152382567492, 132.50690740953544 334.35614994945234, 193.01469673024513 224.45197494134325, 216.6987302454 210.9182415041119, 232.48808592216992 168.06141895287928, 300.1567531083267 120.69335192256952 )"
			wkt_b := "LINESTRING ( 310 540, 460 710, 540 950, 500 1140, 600 1080, 660 930 )"
			wkt_c := "LINESTRING ( 859.5510685138895 630.4639780582839, 850 650, 843.7617128371196 666.5539338909009, 826.8445460405803 677.8320450885938, 796.3936458068098 716.1776231607492, 791.8824013277327 726.3279232386727, 782.8599123695784 737.6060344363656, 734.3640342194994 743.245090035212, 700 750, 650 750, 618.1994888832636 772.5681791492133, 536.9970882598755 793.9965904248296, 500 800, 465.94498771441084 773.6959902689825, 412.93786508525466 739.8616566759041, 385.87039821079196 720.6888676398264, 361.0585535758678 717.3054342805185, 346.29722623487015 749.9852989139027, 315.5582431228852 768.0047028071353, 275.2795755968359 771.1845976118235, 246.660522354643 707.5867015180614, 265.00252427153964 657.6454550927426, 220.85197491460565 627.9662944103014, 156.8861009121442 574.9681901708519, 118.727363255887 505.01050446771376, 143.10655675849577 442.47257330884776, 114.31878332472097 397.14041836991646, 122.96722299547113 361.9152382567492, 132.50690740953544 334.35614994945234, 179.48096329301376 311.2934311635778, 162.56379649647457 295.5040754868079, 155.79692977785888 260.5419307739602, 193.01469673024513 224.45197494134325, 216.6987302454 210.9182415041119, 232.48808592216992 168.06141895287928, 300.1567531083267 120.69335192256952 )"
			wkt_d := "LINESTRING ( 859.5510685138895 630.4639780582839, 850 650, 843.7617128371196 666.5539338909009, 826.8445460405803 677.8320450885938, 796.3936458068098 716.1776231607492, 791.8824013277327 726.3279232386727, 782.8599123695784 737.6060344363656, 734.3640342194994 743.245090035212, 700 750, 650 750, 618.1994888832636 772.5681791492133, 536.9970882598755 793.9965904248296, 500 800, 465.94498771441084 773.6959902689825, 412.93786508525466 739.8616566759041, 385.87039821079196 720.6888676398264, 361.0585535758678 717.3054342805185, 346.29722623487015 749.9852989139027, 315.5582431228852 768.0047028071353, 275.2795755968359 771.1845976118235, 246.660522354643 707.5867015180614, 265.00252427153964 657.6454550927426, 220.85197491460565 627.9662944103014, 156.8861009121442 574.9681901708519, 118.727363255887 505.01050446771376, 190 510, 260 500, 310 440, 370 400, 450 280 )"

			pln_a := polyln(wkt_a).Coordinates
			pln_b := polyln(wkt_b).Coordinates
			pln_c := polyln(wkt_c).Coordinates
			pln_d := polyln(wkt_d).Coordinates

			g.Assert(Homotopy(pln_a, ctx.Geometry())).IsFalse()
			g.Assert(Homotopy(pln_b, ctx.Geometry())).IsTrue()
			g.Assert(Homotopy(pln_c, ctx.Geometry())).IsFalse()
			g.Assert(Homotopy(pln_d, ctx.Geometry())).IsFalse()

		})
	})

}
