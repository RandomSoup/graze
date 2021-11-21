package render

import "bytes"

/* Rendering */

type RenderLine struct {
	//the text that is displayed
	Text string
	//associated metadata (for example, link target for links)
	Meta string
	//the rendertype - governs how the text is rendered. A single number is used over flags for simplicity
	// 0 = normal text, 1 = header, 2 = subheader, 3 = subsubheader, 4 = quote, 5 = small, 6 = monospace, 7 = link
	Rendtype int
}

func NewRendLine(text, meta string, rendtype int) RenderLine {
	return RenderLine{
		Text:     text,
		Meta:     meta,
		Rendtype: rendtype,
	}
}

//from https://stackoverflow.com/questions/33633168/how-to-insert-a-character-every-x-characters-in-a-string-in-golang
func insertNth(s string, n int) string {
	var buffer bytes.Buffer
	var n_1 = n - 1
	var l_1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n_1 && i != l_1 {
			buffer.WriteRune('\n')
		}
	}
	return buffer.String()
}
