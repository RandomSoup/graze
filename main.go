package main

import (
	//"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"luminoso.dev/graze/graze"
	"luminoso.dev/graze/graze/render"
)

var (
	grazeCore        graze.GrazeCore
	font             rl.Font
	framecount       int
	scrollOffset     int
	shouldScroll     bool
	darkMode         bool
	dmBackgroundCol  = color.RGBA{38, 38, 38, 255}
	dmBackgroundLAcc = color.RGBA{48, 48, 48, 255}
)

func main() {
	/* Initialize Main State */
	grazeCore = graze.GrazeCore{}
	grazeCore.Init()
	framecount = 0
	rl.InitWindow(800, 450, "Graze")
	rl.SetTargetFPS(60)
	font = rl.LoadFont("data/font.ttf")

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

		//at least updaet once every 64 frames
		if framecount%64 == 0 {
			grazeCore.TargetRenderFrames += 4
		}
		framecount += 1

		if framecount > 256 {
			framecount = 0
		}
		/* Keyboardy/Input things */
		key := rl.GetCharPressed()
		if key != 0 || rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsMouseButtonPressed(rl.MouseRightButton) || rl.GetMouseWheelMove() != 0 {
			grazeCore.TargetRenderFrames += 4
		}
		if rl.GetMouseDelta().X != 0 || rl.GetMouseDelta().Y != 0 {
			grazeCore.TargetRenderFrames += 1
		}

		if key >= 32 && key <= 125 {
			grazeCore.QBCurrentURL += string(key)

		}

		if rl.IsKeyDown(rl.KeyBackspace) && framecount%5 == 0 {
			grazeCore.TargetRenderFrames += 4
			slen := len(grazeCore.QBCurrentURL)
			if slen > 0 {
				grazeCore.QBCurrentURL = grazeCore.QBCurrentURL[:slen-1]
			}

		}

		if rl.IsKeyDown(rl.KeyUp) {
			grazeCore.TargetRenderFrames += 4
			scrollOffset += 1
		}
		if rl.IsKeyDown(rl.KeyDown) {
			grazeCore.TargetRenderFrames += 4
			scrollOffset -= 1
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyDown(rl.KeyD) {
			darkMode = !darkMode
			grazeCore.TargetRenderFrames += 4
		}

		if rl.IsKeyDown(rl.KeyEnter) && !grazeCore.QueryActive {
			grazeCore.TargetRenderFrames += 4
			grazeCore.SBStatus = "load"
			go grazeCore.Query()
		}

		/* Main GUI */

		rl.BeginDrawing()
		//fmt.Printf("[RDB] FC: %v | TRF: %v | shouldRender: %v | FPS:%v\n", framecount, grazeCore.TargetRenderFrames, grazeCore.TargetRenderFrames > 0, rl.GetFPS())
		//rl.DrawText(fmt.Sprintf("[RDB] FC: %v | TRF: %v | shouldRender: %v", framecount, grazeCore.TargetRenderFrames, grazeCore.TargetRenderFrames > 0), 5, int32(rl.GetScreenHeight()-30), 15, rl.Purple)
		if grazeCore.TargetRenderFrames > 0 {
			grazeCore.TargetRenderFrames -= 1
			if darkMode {
				rl.ClearBackground(dmBackgroundCol)
			} else {
				rl.ClearBackground(rl.RayWhite)
			}

			//Top Bar (status/util)
			render.SBRender(grazeCore.QBCurrentURL, grazeCore.SBStatus, grazeCore.SBStatusColor, 5, 0, 20, font, darkMode)
			if darkMode {
				rl.DrawLine(0, 23, int32(rl.GetScreenWidth()), 23, dmBackgroundLAcc)
			} else {
				rl.DrawLine(0, 23, int32(rl.GetScreenWidth()), 23, rl.LightGray)
			}
			navLink := ""
			render.CPRender(grazeCore.RenderLines, int32(5), int32(30), int32(3), 18, font, &navLink, scrollOffset, darkMode)
			if navLink != "" {
				grazeCore.QBCurrentURL = navLink
				grazeCore.SBStatus = "load"
				go grazeCore.Query()
			}
		}
		rl.EndDrawing()

	}

	rl.CloseWindow()
}
