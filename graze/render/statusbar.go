package render

import (
	"github.com/gdamore/tcell/v2"
	"strings"
)

//renders the status bar. Who woulda thought?
func SBRender(URLText, statusText string, statusColor int, x, y, size float32, dark bool, s tcell.Screen) {
	width, _ := s.Size()
	parts := strings.Split(URLText, "://")
	if len(parts) != 2 {
		DrawText(s, int(x), int(y), tcell.StyleDefault.Foreground(tcell.ColorGray).Background(tcell.ColorWhite), URLText)
	} else {
		DrawText(s, int(x), int(y), tcell.StyleDefault.Foreground(tcell.ColorGray).Background(tcell.ColorWhite), parts[0]+"://")
		DrawText(s, int(x)+len(parts[0]+"://"), int(y), tcell.StyleDefault.Foreground(tcell.ColorGray).Background(tcell.ColorBlue), parts[1])
	}
	switch statusColor {
	case 0:
		DrawText(s, width-(len(statusText)+7), int(y), tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite), statusText)
	case 1:
		DrawText(s, width-(len(statusText)+7), int(y), tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorRed), statusText)
	case 2:
		DrawText(s, width-(len(statusText)+7), int(y), tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorOrange), statusText)
	case 3:
		DrawText(s, width-(len(statusText)+7), int(y), tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorGreen), statusText)
	default:
		//	rl.DrawTextEx(font, "Invalid Col", statusPosVec, size, 2, rl.Red)
	}

}
