package schema

import (
	"strings"

	"luminoso.dev/graze/graze/render"
)

func parseGemtext(input string) []render.RenderLine {
	var result []render.RenderLine

	lines := strings.Split(input, "\n")
	monospacemode := false
	for _, line := range lines {
		if strings.HasPrefix(line, "```") {
			monospacemode = !monospacemode
			continue
		}
		if monospacemode {
			result = append(result, render.NewRendLine(line, "", 6))
			continue
		}
		if strings.HasPrefix(line, "# ") {
			result = append(result, render.NewRendLine(strings.TrimPrefix(line, "# "), "", 1))
		} else if strings.HasPrefix(line, "## ") {
			result = append(result, render.NewRendLine(strings.TrimPrefix(line, "## "), "", 2))
		} else if strings.HasPrefix(line, "### ") {
			result = append(result, render.NewRendLine(strings.TrimPrefix(line, "### "), "", 3))
		} else if strings.HasPrefix(line, "> ") {
			result = append(result, render.NewRendLine(strings.TrimPrefix(line, "> "), "", 4))
		} else if strings.HasPrefix(line, "=> ") {
			clean := strings.TrimPrefix(line, "=> ")
			sp := strings.Split(clean, " ")
			if len(sp) > 2 {
				result = append(result, render.NewRendLine(strings.Join(sp[1:], " "), sp[0], 7))
			} else {
				result = append(result, render.NewRendLine(sp[0], sp[0], 7))
			}
		} else {
			result = append(result, render.NewRendLine(line, "", 0))
		}
	}

	return result
}

func parsePureText(input string) []render.RenderLine {
	var result []render.RenderLine
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		result = append(result, render.NewRendLine(line, "", 0))
	}
	return result
}
