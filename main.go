package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"luminoso.dev/graze/graze"
	"luminoso.dev/graze/graze/render"
)

var (
	grazeCore    graze.GrazeCore
	font         rl.Font
	framecount   int
	scrollOffset int
	shouldScroll bool
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
		framecount += 1
		if framecount > 256 {
			framecount = 0
		}
		/* Keyboardy/Input things */
		key := rl.GetCharPressed()
		if key >= 32 && key <= 125 {
			grazeCore.QBCurrentURL += string(key)
		}

		if rl.IsKeyDown(rl.KeyBackspace) && framecount%4 == 0 {
			slen := len(grazeCore.QBCurrentURL)
			if slen > 0 {
				grazeCore.QBCurrentURL = grazeCore.QBCurrentURL[:slen-1]
			}
		}

		if rl.IsKeyDown(rl.KeyUp) {
			scrollOffset += 1
		}
		if rl.IsKeyDown(rl.KeyDown) {
			scrollOffset -= 1
		}

		if rl.IsKeyDown(rl.KeyEnter) && !grazeCore.QueryActive {
			grazeCore.SBStatus = "load"
			go grazeCore.Query()
		}

		if rl.IsKeyDown(rl.KeyRightBracket) {
			grazeCore.QBCurrentURL = "piper://localhost/test.gmi"
			go grazeCore.Query()
		}

		/* Main GUI */
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		//Top Bar (status/util)
		render.SBRender(grazeCore.QBCurrentURL, grazeCore.SBStatus, grazeCore.SBStatusColor, 5, 0, 20, font, shouldScroll)
		rl.DrawLine(0, 23, int32(rl.GetScreenWidth()), 23, rl.LightGray)
		navLink := ""
		render.CPRender(grazeCore.RenderLines, int32(5), int32(30), int32(3), 18, font, &navLink, scrollOffset, &shouldScroll)
		if navLink != "" {
			grazeCore.QBCurrentURL = navLink
			grazeCore.SBStatus = "load"
			go grazeCore.Query()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
