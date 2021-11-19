package schema

import (
	"luminoso.dev/graze/graze/gconst"
	"luminoso.dev/graze/graze/render"
)

/*
graze.go
implements the graze:// internal schema
(C) Luminoso 2021
*/

type GrazeSchemaHandler struct {
}

func (g GrazeSchemaHandler) Name() string {
	return "graze:// internal utility schema handler"
}

func (g GrazeSchemaHandler) Query(uri string) SchemaQueryResponse {
	grazeQueryResponse := SchemaQueryResponse{}
	grazeQueryResponse.ShouldRedir = false // it's internal about/similar pages. Why on earth would we redir

	switch uri {
	case "about":
		grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("Graze: The Simple Browser", "", 1))
		grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("-------------------------", "", 0))
		grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("Built on Go and Raylib", "", 0))
		grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("Graze "+gconst.VERSION, "", 6))
		if gconst.DEBUG_ENABLED {
			grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("Debug Mode Active (graze/gconst/const.go:DEBUG_ENABLED)", "", 8))
		}
		grazeQueryResponse.Status = "OK"
		grazeQueryResponse.StatusColor = 0
	case "home":
		grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("Graze: The Simple Browser", "", 1))
		grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("type a URL, and press enter to go there!", "", 0))
		grazeQueryResponse.Status = "OK"
		grazeQueryResponse.StatusColor = 0
	case "rendtest":
		for i := 0; i < 9; i++ {
			grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("Ask not what grapes can do for you", "", i))
			grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("!@#$%^&**().,{}[]1234567890-><-", "", i))
		}
		grazeQueryResponse.Status = "OK"
		grazeQueryResponse.StatusColor = 0
	default:
		grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("Unknown Page! How did you get here?", "", 1))
		grazeQueryResponse.Status = "Unknown"
		grazeQueryResponse.StatusColor = 1
	}

	return grazeQueryResponse
}
