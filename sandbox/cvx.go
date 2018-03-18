package main

import (
	"fmt"
	"github.com/intdxdt/geom"
)

func main() {

	var awkt = "LINESTRING ( 146 245, 132 317, 198 371, 267 397, 368 376, 447 339, 502 279, 523 214, 492.5217391304348 152.00000000000003, 366.8729756607012 131.37937372134223, 270.5033370195649 155.10039357880905, 218.60869565217388 210.69565217391306, 219.43513831516063 261.950201110032, 254.32469346346159 305.09066722753266, 304.69565217391306 313.304347826087, 355.5652173913044 307.6521739130435, 389.9130434782609 280.69565217391306, 403.3913043478261 218.5217391304348, 427.7391304347826 158.08695652173915, 536.8695652173914 108.08695652173915, 604.9215541470315 156.81201749892313, 642.1943603382012 183.49925740212734, 682 212, 647.5834615512252 255.89355628249535 )"
	var coordinates = geom.NewLineStringFromWKT(awkt).Coordinates()
	var ply = geom.NewPolygon(coordinates)
	fmt.Println(ply.WKT())

	//var wkt = "LINESTRING ( 146 245, 132 317, 198 371, 267 397, 368 376, 447 339, 502 279, 523 214, 489 168, 375.86733476421915 130.9143720672096, 292 140, 228 197, 219.43513831516063 261.950201110032, 268 322, 322 310, 376 276, 421 209, 473 124, 572 117, 640.130181640881 140.63700179377503, 682 212, 647.5834615512252 255.89355628249535 )"
	//var cwkt = "POLYGON (( 446 197, 446 225, 479 225, 479 197, 446 197 ))"
	//var cgs = ctx.NewContexts(ctx.New(geom.NewPolygonFromWKT(cwkt), -1, 0))
	//var coords = geom.NewLineStringFromWKT(wkt).Coordinates()
	//var h = relate.Homotopy(coords, cgs)
	//fmt.Println(h)

	var cwkt = "LINESTRING ( 120.40127879519089 -44.81956151009723, 132.92438563327013 -32.999999999999964, 197.19659735349703 -37.914933837429075, 237.3478260869564 -13.86956521739125, 332.5463137996218 -27.706994328922455, 382.8298676748582 -73.83175803402645, 386.2325141776937 -127.51795841209831, 350.69376181474473 -154.36105860113423, 307.2155009451795 -164.94706994328925, 237.6502835538751 -130.92060491493385, 231.34574089145985 -66.16939053389925, 290.9584120982986 -48.879017013232485, 337.0831758034026 -69.67296786389412, 330.2778827977315 -136.96975425330814, 266.1560344454161 -162.73645379615442, 174.73913043478245 -186.91304347826087 )"
	var ccoordinates = geom.NewLineStringFromWKT(cwkt).Coordinates()
	var cply = geom.NewPolygon(ccoordinates)
	fmt.Println(cply.WKT())

	var dwkt = "LINESTRING ( 109.29167991823921 -370.5295256949482, 135.59232921551876 -389.39738279951825, 173.32804342465897 -368.242512712576, 221.92706930006685 -338.5113439417383, 312.83583534935923 -329.36329201224976, 359.14784824239496 -364.8119932390178, 361.4348612247671 -432.2788762189958, 307.1183028934289 -481.4496553399967, 228.21635500159022 -465.44056446339175, 195.62642000278728 -408.2652399040884, 239.6514199134509 -376.8188113964715, 322.5556405244408 -373.9600451685064, 338.5647314010457 -400.83244771137896, 320.8403807876617 -425.41783727187936, 275.10012114021896 -421.41556455272814, 242.51018614141606 -370.5295256949482, 275.671874385812 -363.0967335022387, 295.1114847359752 -387.11036981714614, 275.671874385812 -406.5499801673093, 248.79947184293943 -410.5522528864605, 211.63551087939223 -407.6934866584954, 185.9066148277057 -410.5522528864605, 165.89525123194954 -437.4246554293331, 167.0387577231356 -475.73212288406637, 125.30077079484415 -496.8869929710086 )"
	var dcoordinates = geom.NewLineStringFromWKT(dwkt).Coordinates()
	var dply = geom.NewPolygon(dcoordinates)
	fmt.Println(dply.WKT())
}