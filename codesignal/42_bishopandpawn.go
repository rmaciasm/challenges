import "math"

func bishopAndPawn(bishop string, pawn string) bool {
	xB, yB, xP, yP := float64(bishop[0]-'a'), float64(bishop[1]-'0'), float64(pawn[0]-'a'), float64(pawn[1]-'0')
	return math.Abs(xB-xP) == math.Abs(yB-yP)
}