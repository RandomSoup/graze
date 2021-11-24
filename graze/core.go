package graze

import (
	"strings"

	"luminoso.dev/graze/graze/render"
	"luminoso.dev/graze/graze/schema"
)

//core application state struct
type GrazeCore struct {
	SchemaTable map[string]schema.ISchemaHandler
	// Status Bar
	SBStatus      string
	SBStatusColor int
	// Query Bar (the text field for the address)
	QBCurrentURL string
	QueryActive  bool
	// Main rendering
	RenderLines        []render.RenderLine
	TargetRenderFrames int
	ForceUpdate        bool
}

//initial setup
func (g *GrazeCore) Init() {
	g.SchemaTable = schema.BuildSchemaTable()
	g.QBCurrentURL = "graze://home"
	g.SBStatus = "Idle"
	g.SBStatusColor = 0
	g.TargetRenderFrames = 8
	g.QueryActive = false
	//populate data for about page
	g.Query()
}

//the top-level query function. Passes through to the appropriate query handler
func (g *GrazeCore) Query() {
	g.QueryActive = true
	g.RenderLines = []render.RenderLine{}

	//*one* hardcoded thing: Schema list, because it requires access to grazeCore
	if g.QBCurrentURL == "graze://schemas" {
		g.SBStatus = "OK"
		g.SBStatusColor = 0
		for name, handler := range g.SchemaTable {
			g.RenderLines = append(g.RenderLines, render.NewRendLine(name+"://", "", 0))
			g.RenderLines = append(g.RenderLines, render.NewRendLine(handler.Name(), "", 6))
		}
		g.QueryActive = false
		return
	}

	urlParts := strings.Split(g.QBCurrentURL, "://")
	shandler, ok := g.SchemaTable[urlParts[0]]
	if ok {
		qresult := shandler.Query(urlParts[1])
		if qresult.ShouldRedir {
			g.QBCurrentURL = qresult.RedirTarget
			g.Query()
		} else {
			g.RenderLines = qresult.Contents
			g.SBStatus = qresult.Status
			g.SBStatusColor = qresult.StatusColor
		}
	} else {
		g.RenderLines = append(g.RenderLines, render.NewRendLine("Schema has no registered handler!", "", 0))
		g.SBStatus = "error"
		g.SBStatusColor = 1
	}
	g.TargetRenderFrames += 4
	g.QueryActive = false
	g.ForceUpdate = true
}
