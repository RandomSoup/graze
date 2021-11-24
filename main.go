package main

import (
	"github.com/gdamore/tcell/v2"
	"luminoso.dev/graze/graze"
	"luminoso.dev/graze/graze/render"
	"os"
	"strings"
)

var (
	grazeCores   []graze.GrazeCore
	framecount   int
	scrollOffset int
	shouldScroll bool
	darkMode     bool
	tabs         []string
	currTab      int
)

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	//boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
	currTab = 0
	tabs = append(tabs, "home")
	grazeCores = append(grazeCores, graze.GrazeCore{})

	grazeCores[currTab].Init()
	framecount = 0

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		//log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		//log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	quit := func() {
		s.Fini()
		os.Exit(0)
	}
	for {
		framecount += 1
		if framecount > 256 {
			framecount = 0
		}
		// Update screen
		if grazeCores[currTab].ForceUpdate {
			s.Clear()
			grazeCores[currTab].ForceUpdate = false
		}
		render.SBRender(grazeCores[currTab].QBCurrentURL, grazeCores[currTab].SBStatus, grazeCores[currTab].SBStatusColor, 5, 0, 20, true, s)
		navlink := ""
		render.CPRender(grazeCores[currTab].RenderLines, int32(5), int32(4), int32(3), 18, &navlink, scrollOffset, darkMode, 900, s)
		s.Show()
		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			} else if ev.Key() == tcell.KeyCtrlL {
				//	s.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				//	s.Clear()
			}
			if ev.Key() == tcell.KeyBackspace {
				slen := len(grazeCores[currTab].QBCurrentURL)
				if slen > 0 {
					grazeCores[currTab].QBCurrentURL = grazeCores[currTab].QBCurrentURL[:slen-1]
				}
				s.Clear()
			} else if ev.Key() == tcell.KeyEnter {
				scrollOffset = 0
				grazeCores[currTab].TargetRenderFrames += 4
				grazeCores[currTab].SBStatus = "load"
				tabs[currTab] = strings.Split(grazeCores[currTab].QBCurrentURL, "/")[len(strings.Split(grazeCores[currTab].QBCurrentURL, "/"))-1]
				s.Clear()
				go grazeCores[currTab].Query()
			} else if ev.Rune() != 0 {
				grazeCores[currTab].QBCurrentURL += string(ev.Rune())
				s.Clear()
			}

		case *tcell.EventMouse:
			//_, _ := ev.Position()
			button := ev.Buttons()
			// Only process button events, not wheel events
			button &= tcell.ButtonMask(0xff)

			switch ev.Buttons() {
			case tcell.ButtonNone:
				break
			}
		}
	}
}
