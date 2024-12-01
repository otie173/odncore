package player

var (
	players map[string]Player
)

type Player struct {
	Nickname string
	X        float32
	Y        float32
	TargetX  float32
	TargetY  float32
}

func ChangePosition(nickname string, x, y, targetX, targetY float32) {
	if entry, ok := players[nickname]; ok {
		entry.X = x
		entry.Y = y
		entry.TargetX = targetX
		entry.TargetY = targetY

		players[nickname] = entry
	}
}
