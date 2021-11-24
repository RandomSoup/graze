package render

import (
	"github.com/gdamore/tcell/v2"
)

//renders the main content pane
//returns a potential link to nav to
func CPRender(lines []RenderLine, x, y, linespace int32, fontsize float32, linkRet *string, scrollOffset int, dark bool, lowYCap int, s tcell.Screen) {
	linkRetl := ""
	origY := y
	y -= int32(scrollOffset * int(fontsize))
	for _, line := range lines {
		shouldrender := true
		if y < origY {
			shouldrender = false
		}
		if y > int32(lowYCap) {
			shouldrender = false
		}

		if shouldrender {
			DrawText(s, int(x), int(y), tcell.StyleDefault.Foreground(tcell.ColorGray).Background(tcell.ColorWhite), line.Text)
		}
		y += 1
	}
	*linkRet = linkRetl
}
