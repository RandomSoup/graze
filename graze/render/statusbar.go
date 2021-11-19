package render

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//renders the status bar. Who woulda thought?
func SBRender(URLText, statusText string, statusColor int, x, y, size float32, font rl.Font, scrollTT bool) {
	parts := strings.Split(URLText, "://")
	if len(parts) != 2 {
		rl.DrawTextEx(font, URLText, rl.Vector2{x, y}, size, 2, rl.LightGray)
	} else {
		rl.DrawTextEx(font, parts[0]+"://", rl.Vector2{x, y}, size, 2, rl.LightGray)
		rl.DrawTextEx(font, parts[1], rl.Vector2{x + float32(rl.MeasureText(parts[0]+"://", int32(size))) + 5, y}, size, 2, rl.Blue)
	}

	statusPosVec := rl.Vector2{float32(rl.GetScreenWidth() - (int(rl.MeasureText(statusText, int32(size))) + 7)), y}
	//scrollTTPosVec := rl.Vector2{float32((rl.GetScreenWidth() / 2) - 80), float32(rl.GetScreenHeight() - 10)}
	//rl.DrawTextEx(font, "Scroll: Mousewheel", scrollTTPosVec, size/2, 2, rl.Black)

	switch statusColor {
	case 0:
		rl.DrawTextEx(font, statusText, statusPosVec, size, 2, rl.Black)
	case 1:
		rl.DrawTextEx(font, statusText, statusPosVec, size, 2, rl.Red)
	case 2:
		rl.DrawTextEx(font, statusText, statusPosVec, size, 2, rl.Orange)
	case 3:
		rl.DrawTextEx(font, statusText, statusPosVec, size, 2, rl.Green)
	default:
		rl.DrawTextEx(font, "Invalid Col", statusPosVec, size, 2, rl.Red)
	}

}
