package main

import (
	//"fmt"
	"fmt"
	"image/color"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	"luminoso.dev/graze/graze"
	"luminoso.dev/graze/graze/render"
)

var (
	grazeCores       []graze.GrazeCore
	font             rl.Font
	framecount       int
	scrollOffset     int
	shouldScroll     bool
	darkMode         bool
	dmBackgroundCol  = color.RGBA{38, 38, 38, 255}
	dmBackgroundLAcc = color.RGBA{48, 48, 48, 255}
	tabs             []string
	currTab          int
)

func main() {
	/* Initialize Main State */
	currTab = 0
	tabs = append(tabs, "home")
	grazeCores = append(grazeCores, graze.GrazeCore{})

	grazeCores[currTab].Init()
	framecount = 0
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 450, "Graze")
	rl.SetTargetFPS(60)
	font = rl.LoadFont("data/font.ttf")
	darkMode = true
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		framecount += 1

		if framecount > 256 {
			framecount = 0
		}
		/* Keyboardy/Input things */
		key := rl.GetCharPressed()
		if key >= 32 && key <= 125 {
			grazeCores[currTab].QBCurrentURL += string(key)
		}

		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeyOne) && len(tabs) > 1 {
			currTab = 1
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeyTwo) && len(tabs) > 2 {
			currTab = 2
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeyThree) && len(tabs) > 3 {
			currTab = 3
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeyFour) && len(tabs) > 4 {
			currTab = 4
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeyFive) && len(tabs) > 5 {
			currTab = 5
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeySix) && len(tabs) > 6 {
			currTab = 6
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeySeven) && len(tabs) > 7 {
			currTab = 7
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeyEight) && len(tabs) > 8 {
			currTab = 8
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeyNine) && len(tabs) > 9 {
			currTab = 9
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeyZero) && len(tabs) > 0 {
			currTab = 0
		}

		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeyN) && framecount%5 == 0 {
			if len(tabs) < 10 {
				newInd := len(tabs)
				grazeCores = append(grazeCores, graze.GrazeCore{})
				grazeCores[newInd].Init()
				currTab = newInd
				tabs = append(tabs, "home")
			}
		}

		if rl.IsKeyDown(rl.KeyBackspace) && framecount%5 == 0 {
			slen := len(grazeCores[currTab].QBCurrentURL)
			if slen > 0 {
				grazeCores[currTab].QBCurrentURL = grazeCores[currTab].QBCurrentURL[:slen-1]
			}
		}
		if rl.IsKeyDown(rl.KeyUp) {
			grazeCores[currTab].TargetRenderFrames += 4
			scrollOffset += 1
		}
		if rl.IsKeyDown(rl.KeyDown) {
			grazeCores[currTab].TargetRenderFrames += 4
			scrollOffset -= 1
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeyD) {
			darkMode = !darkMode
			grazeCores[currTab].TargetRenderFrames += 4
		}

		if rl.IsKeyDown(rl.KeyEnter) && !grazeCores[currTab].QueryActive {
			scrollOffset = 0
			grazeCores[currTab].TargetRenderFrames += 4
			grazeCores[currTab].SBStatus = "load"
			tabs[currTab] = strings.Split(grazeCores[currTab].QBCurrentURL, "/")[len(strings.Split(grazeCores[currTab].QBCurrentURL, "/"))-1]
			go grazeCores[currTab].Query()
		}

		/* Main GUI */

		rl.BeginDrawing()

		if darkMode {
			rl.ClearBackground(dmBackgroundCol)
		} else {
			rl.ClearBackground(rl.RayWhite)
		}

		//Top Bar (status/util)
		render.SBRender(grazeCores[currTab].QBCurrentURL, grazeCores[currTab].SBStatus, grazeCores[currTab].SBStatusColor, 5, 0, 20, font, darkMode)
		if darkMode {
			rl.DrawLine(0, 23, int32(rl.GetScreenWidth()), 23, dmBackgroundLAcc)
		} else {
			rl.DrawLine(0, 23, int32(rl.GetScreenWidth()), 23, rl.LightGray)
		}
		navLink := ""
		render.CPRender(grazeCores[currTab].RenderLines, int32(5), int32(30), int32(3), 18, font, &navLink, scrollOffset, darkMode, rl.GetScreenHeight()-48)
		if navLink != "" {
			scrollOffset = 0
			if strings.Contains(navLink, "piper://") {
				grazeCores[currTab].QBCurrentURL = navLink
			} else {
				grazeCores[currTab].QBCurrentURL += "" + navLink
			}
			grazeCores[currTab].SBStatus = "load"
			go grazeCores[currTab].Query()
		}
		//bottom bar (tabs)
		if darkMode {
			rl.DrawLine(0, int32(rl.GetScreenHeight()-20), int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()-20), dmBackgroundLAcc)
		} else {
			rl.DrawLine(0, int32(rl.GetScreenHeight()-20), int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()-20), rl.LightGray)
		}

		tbX := float32(0)
		for i, tab := range tabs {
			if i == currTab {
				rl.DrawRectangleRounded(rl.Rectangle{tbX + 5, float32(rl.GetScreenHeight() - 18), 20, 16}, 1, 12, rl.Orange)
			} else {
				rl.DrawRectangleRounded(rl.Rectangle{tbX + 5, float32(rl.GetScreenHeight() - 18), 20, 16}, 1, 12, rl.Yellow)
			}
			rl.DrawTextEx(font, fmt.Sprintf("^%d", i), rl.Vector2{tbX + 6, float32(rl.GetScreenHeight() - 16)}, 14, 3, rl.Black)
			if darkMode {
				rl.DrawTextEx(font, tab, rl.Vector2{tbX + 30, float32(rl.GetScreenHeight() - 16)}, 14, 3, rl.White)
			} else {
				rl.DrawTextEx(font, tab, rl.Vector2{tbX + 30, float32(rl.GetScreenHeight() - 16)}, 14, 3, rl.Black)
			}
			tbX += rl.MeasureTextEx(font, tab, 14, 3).X + 32
		}

		rl.EndDrawing()

	}

	rl.CloseWindow()
}
