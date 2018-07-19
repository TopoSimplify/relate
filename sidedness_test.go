package relate

import (
	"time"
	"testing"
	"github.com/intdxdt/geom"
	"github.com/franela/goblin"
	"github.com/TopoSimplify/ctx"
)

func TestSidedness(t *testing.T) {
	var g = goblin.Goblin(t)
	g.Describe("homotopic sidedness test", func() {
		g.It("should test sidedness to a context geometry - case 1", func() {
			g.Timeout(1 * time.Hour)
			var poly0 = "POLYGON (( 278 307, 270 298, 274 286, 279 272, 301 274, 308 288, 311 304, 296 308, 278 307 ))"
			var polyln = "LINESTRING ( 155 171, 207 166, 253 175, 317 171, 367 182, 400 200, 428 249, 417 291, 383 324, 361 332, 333 347, 314 357, 257 383, 204 370, 176 337, 180 305, 214 295, 244 302, 281 332, 316 328, 331 306, 332 291, 315 265, 285 250, 247 261, 231 276, 195 264, 187 230, 216 215, 257 226, 273 217, 273 205, 240 197, 200 200, 178 193, 157 226, 156 246, 151 263, 120 264, 95 249, 89 261, 100 300, 116 359, 139 389, 172 413, 211 425, 256 430, 289 431, 348 427 )"
			var coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			//using only poly0
			var ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms), ).IsTrue()

			//using poly1
			var poly1 = "POLYGON (( 221 347, 205 334, 221 322, 234 324, 237 342, 221 347 ))"
			ctxgeoms = ctx.NewContexts(ctx.New(geom.NewPolygonFromWKT(poly1), -1, 0))
			g.Assert(Homotopy(coords, ctxgeoms), ).IsTrue()

			//using poly0 and poly1
			ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries().Push(
				ctx.New(geom.NewPolygonFromWKT(poly1), -1, 0),
			)
			g.Assert(Homotopy(coords, ctxgeoms)).IsFalse()

			poly0 = "POLYGON (( 221 347, 205 334, 221 322, 234 324, 237 342, 221 347 ))"
			polyln = "LINESTRING ( 155 171, 207 166, 253 175, 317 171, 366.3526316090967 196.8011373660995, 403.24889098928804 217.07380735521562, 449.06512516469047 258.02460073323016, 469.7432485535889 310.7335427049321, 471.50392377693095 357.2445923621374, 459.17233787042085 384.55024686940965, 435.2315676771032 379.0896298976944, 421.41537343261 346.6608527245552, 392.74529568744737 330.93485457399595, 370.2087538305982 346.6746615851287, 375.07353147167777 371.93565702969966, 384.86292089142125 409.47595509141127, 401.0377185968734 428.5916251069456, 362.8344377501964 475.8341413005543, 303.51814850672173 477.88349751591124, 253.9395152835032 467.3480379559773, 233.36391648049494 444.9386180328808, 205.38763189551466 402.77146445551926, 204 370, 176 337, 180 305, 214 295, 191.4007581862021 285.8975596173291, 195 264, 187 230, 216 215, 257 226, 225.08021076634535 263.95613448800214, 248.30540712326405 294.76261255728286, 273.09263453166125 327.6410074616854, 308.7545836354981 374.9439036823187, 348 427 )"
			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms), ).IsFalse()

			poly0 = "POLYGON (( 278 307, 270 298, 274 286, 279 272, 301 274, 308 288, 311 304, 296 308, 278 307 ))"
			ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms), ).IsFalse()

			polyln = "LINESTRING ( 155 171, 189.63910305670066 144.08432169246322, 253.9395152835032 145.84597682196465, 317 171, 366.3526316090967 196.8011373660995, 403.24889098928804 217.07380735521562, 449.2399093677868 286.1290060157034, 451.783228702019 325.7276851722446, 415.8591984439269 355.71930167053, 418.6880722510337 402.11283210708103, 389.8335594185446 420.2176244725644, 350.83054740608236 477.91796873298597, 266.27110119001327 476.15631360348453, 164.0951036789298 449.7314866609629, 107.72213953488374 344.9130064556273, 116.53041518239094 298.2291455238391, 130.62365621840246 256.8302499805553, 151.76351777241973 229.52459547328303, 220.46806782297585 190.76818262425135, 263.6286184957611 215.4313544372715, 241.3878725848878 285.58702270327086, 217.39831804889542 317.49495464420494, 208.13648191646578 343.15135132612585, 214.3022748697208 392.47769495216613, 226.63386077623088 413.6175565061834, 242.7959277400187 428.5328808600021, 256 430, 289 431, 348 427 )"
			poly0 = "POLYGON (( 278 307, 270 298, 274 286, 279 272, 301 274, 308 288, 311 304, 296 308, 278 307 ))"
			ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries()
			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			g.Assert(Homotopy(coords, ctxgeoms), ).IsFalse()

			poly0 = "POLYGON (( 221 347, 205 334, 221 322, 234 324, 237 342, 221 347 ))"
			ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms), ).IsFalse()

			coords = []geom.Point{{2, 2}, {5, 2}, {7, 2}, {9, 2}}
			var lnconst = geom.NewLineString([]geom.Point{{3, 2}, {6, 2}, {6.5, 2}})
			ctxgeoms = ctx.New(lnconst, -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms), ).IsTrue()

			var plyconst = geom.NewPolygon([]geom.Point{{4, 1}, {4, 3}, {5, 3}, {5, 1}})
			ctxgeoms = ctx.New(plyconst, -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms), ).IsTrue()

			polyln = "LINESTRING ( 155 171, 207 166, 253 175, 317 171, 367 182, 400 200, 428 249, 417 291, 383 324, 361 332, 333 347, 314 357, 257 383, 204 370, 176 337, 180 305, 214 295, 244 302, 281 332, 316 328, 331 306, 332 291, 315 265, 285 250, 247 261, 231 276, 175.9815455950541 265.3822256568778, 187 230, 233.38568778979908 233.9203400309118, 265.9514992272025 234.4723029366305, 274.230942812983 207.97808346213282, 248.54273709311704 199.72323889171014, 215.17091190108192 192.52312210200918, 156.16765559617332 215.4313544372715, 129.74282865365174 249.78362946254956, 164.97593124368052 358.12541992688807, 211 425, 256 430, 289 431, 348 427 )"
			poly0 = "POLYGON (( 278 307, 268.711313755796 290.77251931993817, 235.45711654915712 296.5278817518552, 232.8337248840804 286.90877897990725, 277.40587928436713 278.46806414206173, 285.4987949850666 282.69941864135717, 288.58197836166926 290.77251931993817, 290.2378670788254 299.05196290571865, 278 307 ))"
			ctxgeoms = ctx.NewContexts(ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0))
			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			g.Assert(Homotopy(coords, ctxgeoms), ).IsFalse()

			poly0 = "POLYGON (( 221 347, 205 334, 221 322, 234 324, 237 342, 221 347 ))"
			ctxgeoms = ctx.NewContexts(ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0))
			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			g.Assert(Homotopy(coords, ctxgeoms), ).IsTrue()

			poly1 = "POLYGON (( 232.00298146975462 210.52416263821522, 232.00298146975462 220.68605060410567, 246.51996427816957 220.68605060410567, 246.51996427816957 210.52416263821522, 232.00298146975462 210.52416263821522 ))"
			ctxgeoms = ctx.NewContexts(ctx.New(geom.NewPolygonFromWKT(poly1), -1, 0))
			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			g.Assert(Homotopy(coords, ctxgeoms)).IsTrue()

			ctxgeoms.Push(ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0))
			g.Assert(Homotopy(coords, ctxgeoms)).IsFalse()

			poly1 = "POLYGON (( 225.21282919521113 245.6378900803107, 223.30879887780458 254.77723560386218, 232.44814440135602 257.06207198475005, 241.97464275611355 248.90012869080064, 225.21282919521113 245.6378900803107 ))"
			ctxgeoms = ctx.NewContexts(ctx.New(geom.NewPolygonFromWKT(poly1), -1, 0))
			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			g.Assert(Homotopy(coords, ctxgeoms), ).IsFalse()

			ctxgeoms.Push(ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0))
			g.Assert(Homotopy(coords, ctxgeoms)).IsFalse()

			polyln = "LINESTRING ( 155 171, 207 166, 253 175, 308.261560573073 200.97163319048545, 349.188755713404 221.82133637518243, 373.12730381435244 255.79863045394785, 386.25489470842086 295.9536143652161, 381.62162733404375 323.7532186114787, 355.36644554590686 380.1246383330668, 330.2777717925064 403.49279574550076, 322.32211827742316 392.9402190622814, 309.12410461200943 375.43404549572233, 257 383, 204 370, 176 337, 180 305, 214 295, 261.97977549245553 312.9006348500965, 284.39987101125786 342.6392071444664, 316 328, 331 306, 332 291, 315 265, 285 250, 247 261, 238.25818553022333 281.4357279571874, 215.94139273173917 251.83417896023434, 205.01673883978984 237.34344633671606, 233.38568778979908 233.9203400309118, 265.9514992272025 234.4723029366305, 274.230942812983 207.97808346213282, 248.54273709311704 199.72323889171014, 175.91990615765368 198.7486838153334, 190.09992303692636 217.55741086763288, 129.74282865365174 249.78362946254956, 164.97593124368052 358.12541992688807, 211 425, 256 430, 289 431, 348 427 )"
			poly0 = "POLYGON (( 211.34885547628429 235.3767011676454, 211.34885547628429 241.48063878287113, 213.69652378983264 241.48063878287113, 213.69652378983264 235.3767011676454, 211.34885547628429 235.3767011676454 ))"
			ctxgeoms = ctx.NewContexts(ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0))
			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			g.Assert(Homotopy(coords, ctxgeoms), ).IsFalse()

			polyln = "LINESTRING ( 155 171, 169.0211940381559 189.5980604858441, 178.36905451652774 201.99729511000572, 189.61729454361603 216.9172404309104, 155.33584906460158 235.98971111300912, 171.9249231022847 253.04608301090866, 184.81805947947288 266.3024063282993, 199.21035949720948 260.0405937778234, 216.46768369629606 252.53226438472433, 226.72184529093568 266.1336393496349, 253.77293550266072 261.8918612097899, 286.55244753078927 256.75183112348185, 298.89980243113513 277.3307559573916, 308.9755371420998 294.1236471423328, 311.6569923092747 322.2091261672002, 274.3587106022468 329.32036224961234, 286.70927127817396 345.7024530943654, 302.96002907406313 367.25786239875725, 265.3750851201071 385.89221758862226, 267.0360547209449 409.9762768007707, 292.87995239405126 407.8640164833156, 311.8206399330813 406.3159658301186, 333.85824574639133 408.2420254459906, 348 427 )"
			poly0 = "POLYGON (( 211.34885547628429 235.3767011676454, 211.34885547628429 241.48063878287113, 213.69652378983264 241.48063878287113, 213.69652378983264 235.3767011676454, 211.34885547628429 235.3767011676454 ))"
			ctxgeoms = ctx.NewContexts(ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0))

			coords = geom.NewLineStringFromWKT(polyln).Coordinates();
			g.Assert(Homotopy(coords, ctxgeoms), ).IsTrue()

			polyln = "LINESTRING ( 90 270, 120 270, 150 270, 136.36099104257372 230.88019642162786, 162.54902402818348 222.82234011836331, 180 250, 180 270, 200 250, 210 270, 230 260, 250 270, 280 270, 300 270, 310 290, 310 320, 323.7061500934744 270, 332.43549442201095 270, 365.33840766034115 270, 380 310, 410 270, 450 310, 430 270, 470 270 )"
			poly0 = "POLYGON (( 211.34885547628429 235.3767011676454, 211.34885547628429 241.48063878287113, 213.69652378983264 241.48063878287113, 213.69652378983264 235.3767011676454, 211.34885547628429 235.3767011676454 ))"
			ctxgeoms = ctx.NewContexts(ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0))

			coords = geom.NewLineStringFromWKT(polyln).Coordinates();
			g.Assert(Homotopy(coords, ctxgeoms), ).IsTrue()

			polyln = "LINESTRING ( 90 270, 470 270 )"
			poly0 = "POLYGON (( 211.34885547628429 235.3767011676454, 211.34885547628429 241.48063878287113, 213.69652378983264 241.48063878287113, 213.69652378983264 235.3767011676454, 211.34885547628429 235.3767011676454 ))"
			ctxgeoms = ctx.NewContexts(ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0))

			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			g.Assert(Homotopy(coords, ctxgeoms), ).IsTrue()

			polyln = "LINESTRING ( 90 270, 470 270 )"
			poly0 = "POLYGON (( 370 260, 370 300, 380 300, 380 260, 370 260 ))"
			ctxgeoms = ctx.NewContexts(ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0))

			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			g.Assert(Homotopy(coords, ctxgeoms), ).IsTrue()
		})

		g.It("should test sidedness to a context geometry - case 2", func() {
			var wkt = "LINESTRING ( 400 560, 520 660, 620 600, 720 640, 760 500, 680 460, 720 420, 780 400, 720 360, 580 340, 440 400, 380 440, 440 500, 400 560 )"
			var cgs = ctx.NewContexts(ctx.New(geom.NewLineStringFromWKT(wkt), -1, 0))

			var wkt_a = "LINESTRING ( 980 280, 981.3546694489717 324.82716460080917, 1006.1665140838959 336.10527579850196, 1026.4671142397428 362.0449315531954, 1017.4446252815886 410.5408097032744, 982.4824805687409 427.45797649981364, 947.5203358558933 470.31479905104624, 948.6481469756626 585.3515332675128, 921.5806801011998 611.2911890222063, 906.9191355441992 613.5468112617448, 843.7617128371196 666.5539338909009, 734.3640342194994 743.245090035212, 700 750, 650 750, 465.94498771441084 773.6959902689825, 412.93786508525466 739.8616566759041, 385.87039821079196 720.6888676398264, 361.0585535758678 717.3054342805185, 346.29722623487015 749.9852989139027, 315.5582431228852 768.0047028071353, 275.2795755968359 771.1845976118235, 246.660522354643 707.5867015180614, 265.00252427153964 657.6454550927426, 220.85197491460565 627.9662944103014, 156.8861009121442 574.9681901708519, 118.727363255887 505.01050446771376, 143.10655675849577 442.47257330884776, 122.96722299547113 361.9152382567492, 132.50690740953544 334.35614994945234, 193.01469673024513 224.45197494134325, 216.6987302454 210.9182415041119, 232.48808592216992 168.06141895287928, 300.1567531083267 120.69335192256952 )"
			var wkt_b = "LINESTRING ( 310 540, 460 710, 540 950, 500 1140, 600 1080, 660 930 )"
			var wkt_c = "LINESTRING ( 859.5510685138895 630.4639780582839, 850 650, 843.7617128371196 666.5539338909009, 826.8445460405803 677.8320450885938, 796.3936458068098 716.1776231607492, 791.8824013277327 726.3279232386727, 782.8599123695784 737.6060344363656, 734.3640342194994 743.245090035212, 700 750, 650 750, 618.1994888832636 772.5681791492133, 536.9970882598755 793.9965904248296, 500 800, 465.94498771441084 773.6959902689825, 412.93786508525466 739.8616566759041, 385.87039821079196 720.6888676398264, 361.0585535758678 717.3054342805185, 346.29722623487015 749.9852989139027, 315.5582431228852 768.0047028071353, 275.2795755968359 771.1845976118235, 246.660522354643 707.5867015180614, 265.00252427153964 657.6454550927426, 220.85197491460565 627.9662944103014, 156.8861009121442 574.9681901708519, 118.727363255887 505.01050446771376, 143.10655675849577 442.47257330884776, 114.31878332472097 397.14041836991646, 122.96722299547113 361.9152382567492, 132.50690740953544 334.35614994945234, 179.48096329301376 311.2934311635778, 162.56379649647457 295.5040754868079, 155.79692977785888 260.5419307739602, 193.01469673024513 224.45197494134325, 216.6987302454 210.9182415041119, 232.48808592216992 168.06141895287928, 300.1567531083267 120.69335192256952 )"
			var wkt_d = "LINESTRING ( 859.5510685138895 630.4639780582839, 850 650, 843.7617128371196 666.5539338909009, 826.8445460405803 677.8320450885938, 796.3936458068098 716.1776231607492, 791.8824013277327 726.3279232386727, 782.8599123695784 737.6060344363656, 734.3640342194994 743.245090035212, 700 750, 650 750, 618.1994888832636 772.5681791492133, 536.9970882598755 793.9965904248296, 500 800, 465.94498771441084 773.6959902689825, 412.93786508525466 739.8616566759041, 385.87039821079196 720.6888676398264, 361.0585535758678 717.3054342805185, 346.29722623487015 749.9852989139027, 315.5582431228852 768.0047028071353, 275.2795755968359 771.1845976118235, 246.660522354643 707.5867015180614, 265.00252427153964 657.6454550927426, 220.85197491460565 627.9662944103014, 156.8861009121442 574.9681901708519, 118.727363255887 505.01050446771376, 190 510, 260 500, 310 440, 370 400, 450 280 )"

			var pln_a = polyln(wkt_a).Coordinates
			var pln_b = polyln(wkt_b).Coordinates
			var pln_c = polyln(wkt_c).Coordinates
			var pln_d = polyln(wkt_d).Coordinates

			g.Assert(Homotopy(pln_a, cgs)).IsFalse()
			g.Assert(Homotopy(pln_b, cgs)).IsTrue()
			g.Assert(Homotopy(pln_c, cgs)).IsFalse()
			g.Assert(Homotopy(pln_d, cgs)).IsFalse()

		})

	})

}

func TestSidness2(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("context neighbours", func() {
		g.It("should test sidedness to a context geometry - case 3", func() {
			g.Timeout(1 * time.Hour)
			var polyln = "LINESTRING ( 400 560, 520 660, 620 600, 720 640, 760 500, 680 460, 720 420, 780 400, 720 360, 580 340, 440 400, 380 440, 440 500, 400 560 )"
			var poly0 = "POLYGON (( 527.213424901364 486.0891351812401, 527.213424901364 543.0444768497492, 573.5724239338715 543.0444768497492, 573.5724239338715 486.0891351812401, 527.213424901364 486.0891351812401 ))"
			var coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			//using only poly0
			var ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms)).IsFalse()

			polyln = "LINESTRING ( 400 560, 292.5633432073529 639.8060743748151, 276.47377443598197 797.6678482364749, 384.46641739350747 823.7951005649086, 444.93340971672137 771.6588422727477, 579.5499014458117 813.3441996335351, 645.7389406778434 778.5078631956236, 617.8698715275143 703.609739854114, 678.8334602938593 677.4824875256804, 607.7448975133443 621.1986472034275, 497.68451081671975 649.6134183753512, 492.45906035103303 583.4243791433195, 429.7536547627924 595.6170968965885, 400 560 )"
			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms)).IsTrue()

			poly0 = "POLYGON (( 419.30275383141895 687.9333884570539, 419.30275383141895 738.4460762920255, 461.1063575569127 738.4460762920255, 461.1063575569127 687.9333884570539, 419.30275383141895 687.9333884570539 ))"
			ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms)).IsFalse()

			polyln = "LINESTRING ( -184.33370092860702 170.51113898932573, -235.5790071667459 297.5103761881917, -211.07038244415773 371.0362503559562, -87.26103008104715 419.28056792962565, 13.738969918952847 398.28056792962565, 92.73896991895285 361.28056792962565, 323.66324786685686 335.3873416685552, 440.6362294973913 229.55464400283358, 183.29566991021554 70.24858330601049, 12.61194557965402 153.65994165096788, -83.75769306148223 177.3809615084347, -135.65233442887327 232.9762201035387, -134.82589176588652 284.2307690396577, -99.93633661758557 327.3712351571583, -49.56537790713409 335.58491575571264, 1.3041873102572197 329.93274184266915, 35.65201339721372 302.9762201035387, 49.13027426677897 240.80230706006046, 73.47810035373547 180.3675244513648, 153.21690320522094 153.80071304210654, 226.74277737298547 189.4496217295075, 276.8740552146431 169.39711059284446, 327.73896991895276 234.28056792962565, 293.3224314701781 278.174124212121 )"
			poly0 = "POLYGON (( 144.5093270074935 204.8742766590772, 144.5093270074935 235.1277696233169, 213.66016806861285 235.1277696233169, 213.66016806861285 204.8742766590772, 144.5093270074935 204.8742766590772 ))"
			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms)).IsTrue()

			polyln = "LINESTRING ( -200 -300, -213.09757729958832 -240.95296696640685, -213.67173452622256 -205.92937614171896, -170.60994252865547 -173.7765714502022, -126.39983607781993 -162.29342691751762, -84.76811597586375 -177.9637934727483, -35.63603004812814 -196.4376157005745, -33.960522589709235 -250.71363981918873, -55.204339975175664 -265.64172771167864, -98.80151775133675 -274.1496569368973, -124.10320717128302 -337.9855382675914, -191.09621058165152 -368.66831826914733, -270.7303986647175 -312.4873284532295, -280.36759550643706 -225.01123404377455, -238.11219396966646 -121.96736012077253, -128.39641454085856 -101.95164360335487, -70.1324278676656 -100.85860366765525 )"
			poly0 = "POLYGON (( -65.01216009495633 -156.27932633432664, -65.01216009495633 -140.05215379648916, -36.530399697148475 -140.05215379648916, -36.530399697148475 -156.27932633432664, -65.01216009495633 -156.27932633432664 ))"
			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms)).IsTrue()

			polyln = "LINESTRING ( 231.31696746126295 -168.4417854573867, 186.75351543760672 -142.37637012279532, 125.37366642389154 -138.59268080003204, 98.04702131504574 -157.0907174890969, 82.91226402399269 -191.56433131871776, 90.47964266951921 -257.14827957994765, 121.58997710112827 -265.5564780749771, 139.24719394069018 -245.79721161165787, 130.43051435450164 -213.6746310114183, 131.8610207147146 -189.83285834120238, 149.75744205947703 -180.21326335042795, 167.14684426663416 -187.4486810741808, 174.1786808187153 -216.63647446162548, 190.11679483561852 -215.94810695430323, 211.55770099794367 -218.0501565780606, 218.48147950375636 -211.34561416286152, 214.92098039595547 -195.7684305662325, 203.27876650877175 -191.0561058976105, 191.37802460987294 -172.64588470490142, 149.5039324906744 -159.315389323326, 121.84747619322391 -171.7131111118383, 109.818499208087 -203.75621913651048, 112.76136868134732 -235.28696349287102 )"
			poly0 = "POLYGON (( 143.45129318820491 -206.6990886097708, 143.45129318820491 -196.18884049098398, 152.27990160798586 -196.18884049098398, 152.27990160798586 -206.6990886097708, 143.45129318820491 -206.6990886097708 ))"
			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms)).IsFalse()

			poly0 = "POLYGON (( 152.25589072457 -241.1826298570605, 152.25589072457 -221.66128415149674, 160.2940918974492 -221.66128415149674, 160.2940918974492 -241.1826298570605, 152.25589072457 -241.1826298570605 ))"
			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms)).IsTrue()

			polyln = "LINESTRING ( 405.4864204322907 -229.58286228018753, 425.0781379743377 -253.83927447510288, 427.2549954790096 -290.2238927674759, 401.44182689507295 -322.07653073017025, 348.5458087250717 -353.04792453986676, 308.14979200910466 -358.9503939864028, 276.44764762700964 -332.55464952148867, 277.98476658722274 -321.3218571199315, 291.35689125877866 -316.03520318001404, 312.5035070184485 -333.4500632173892, 343.9124510144286 -333.13908357386464, 358.5284942600827 -310.4375695965721, 336.1379599263147 -299.5532820732126, 323.0768148982834 -285.8701777581321, 324.009753828857 -264.4125823549378, 334.27208206516735 -252.2843762574801, 354.7967385377881 -253.83927447510288, 385.5837232467191 -276.5407884523955, 399.88878684884867 -265.656500929036, 385.27274360319456 -244.50988516936619, 353.2418403201653 -235.4914755071541, 317.47918131484136 -242.95498695174342, 300.68628056451536 -263.79062306788865, 303.48509735623634 -300.79720064731083, 326.80857062057805 -317.59010139763683 )"
			poly0 = "POLYGON (( 337.9832718619027 -281.24881893305763, 337.9832718619027 -266.27080432520825, 348.46788208739736 -266.27080432520825, 348.46788208739736 -281.24881893305763, 337.9832718619027 -281.24881893305763 ))"
			coords = geom.NewLineStringFromWKT(polyln).Coordinates()
			ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms)).IsTrue()

			var poly1 = "POLYGON (( 334.51010905428546 -325.4448156599873, 334.51010905428546 -314.3741092107073, 339.5027805902353 -314.3741092107073, 339.5027805902353 -325.4448156599873, 334.51010905428546 -325.4448156599873 ))"
			ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly1), -1, 0).AsContextGeometries()
			g.Assert(Homotopy(coords, ctxgeoms)).IsTrue()

			ctxgeoms = ctx.New(geom.NewPolygonFromWKT(poly0), -1, 0).AsContextGeometries()
			ctxgeoms.Push(ctx.New(geom.NewPolygonFromWKT(poly1), -1, 0))
			g.Assert(Homotopy(coords, ctxgeoms)).IsTrue()
		})
	})
}
