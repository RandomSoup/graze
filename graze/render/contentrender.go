package render

import rl "github.com/gen2brain/raylib-go/raylib"

//renders the main content pane
//returns a potential link to nav to
func CPRender(lines []RenderLine, x, y, linespace int32, fontsize float32, font rl.Font, linkRet *string, scrollOffset int) {
	linkRetl := ""
	origY := y
	y -= int32(scrollOffset * int(fontsize))
	for _, line := range lines {
		line.Text = insertNth(line.Text, int(rl.GetScreenWidth()/int(rl.MeasureText("A", int32(fontsize))))-8)
		shouldrender := true
		if y < origY {
			shouldrender = false
		}
		switch line.Rendtype {
		case 0:
			if shouldrender {
				rl.DrawTextEx(font, line.Text, rl.Vector2{float32(x), float32(y)}, fontsize, 3, rl.Black)
			}
			y += int32(fontsize) + linespace
		case 1:
			if shouldrender {
				rl.DrawTextEx(font, line.Text, rl.Vector2{float32(x), float32(y)}, fontsize+(fontsize/2), 3, rl.Black)
			}
			y += int32(fontsize) + int32((fontsize / 2)) + linespace
		case 2:
			if shouldrender {
				rl.DrawTextEx(font, line.Text, rl.Vector2{float32(x), float32(y)}, fontsize+(fontsize/4), 3, rl.Black)
			}
			y += int32(fontsize) + int32((fontsize / 4)) + linespace
		case 3:
			if shouldrender {
				rl.DrawTextEx(font, line.Text, rl.Vector2{float32(x), float32(y)}, fontsize+(fontsize/8), 3, rl.Black)
			}
			y += int32(fontsize) + int32((fontsize / 8)) + linespace
		case 4:
			if shouldrender {
				rl.DrawRectangle(x, y, int32(5), int32(fontsize), rl.LightGray)
				rl.DrawTextEx(font, line.Text, rl.Vector2{float32(x) + 7, float32(y)}, fontsize, 3, rl.Black)
			}
			y += int32(fontsize) + linespace
		case 5:
			if shouldrender {
				rl.DrawTextEx(font, line.Text, rl.Vector2{float32(x), float32(y)}, fontsize-(fontsize/4), 3, rl.Black)
			}
			y += int32(fontsize) - int32((fontsize / 4)) + linespace
		case 6:
			if shouldrender {
				rl.DrawTextEx(font, line.Text, rl.Vector2{float32(x), float32(y)}, fontsize, 3, rl.Beige)
			}
			y += int32(fontsize) + linespace
		case 7:
			if shouldrender {
				mousepos := rl.GetMousePosition()
				isInRange := mousepos.X > 5 && mousepos.X < float32(5+(rl.MeasureText(line.Text, int32(fontsize*1.5)))) && mousepos.Y > float32(y) && mousepos.Y < float32(y)+fontsize+3
				if isInRange {
					rl.DrawTextEx(font, line.Text, rl.Vector2{float32(x), float32(y)}, fontsize, 3, rl.DarkBlue)
					if rl.IsMouseButtonDown(rl.MouseLeftButton) {
						linkRetl = line.Meta
					}
				} else {
					rl.DrawTextEx(font, line.Text, rl.Vector2{float32(x), float32(y)}, fontsize, 3, rl.SkyBlue)
				}
				rl.DrawTextEx(font, line.Text, rl.Vector2{float32(x), float32(y)}, fontsize, 3, rl.SkyBlue)
			}
			y += int32(fontsize) + linespace
		case 8:
			if shouldrender {
				rl.DrawTextEx(font, line.Text, rl.Vector2{float32(x), float32(y)}, fontsize, 3, rl.Red)
			}
			y += int32(fontsize) + linespace
		}
	}

	//scrollbar, if it should exist
	// if y > int32(rl.GetScreenHeight()) {
	// 	*showScroll = true
	// } else {
	// 	*showScroll = false
	// }

	*linkRet = linkRetl
}
