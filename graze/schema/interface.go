package schema

import "luminoso.dev/graze/graze/render"

//the response a schema handler gives to a query.
type SchemaQueryResponse struct {
	//the document content to render
	Contents []render.RenderLine
	//a status to display in the status bar.
	//Integers are not used to avoid unnessecary switch complexity in the main renderloop
	Status string
	//decides the color the status is in. lets the schema do things such as show errors in red
	StatusColor int
	//flag & target for a redirect/rebrowse
	ShouldRedir bool
	RedirTarget string
}

type ISchemaHandler interface {
	//returns a user-friendly name for the schema handler (configuration files?/future proofing)
	Name() string
	//the big one! given a URI, return us the content we want - note this uri *does not* include the schema specifier
	//so, no piper:// or graze://, for example
	Query(string) SchemaQueryResponse
}

func BuildSchemaTable() map[string]ISchemaHandler {
	table := make(map[string]ISchemaHandler)
	//internal browser URLs
	table["graze"] = GrazeSchemaHandler{}
	//piper hander
	table["piper"] = PiperSchemaHandler{}
	// wIPC debugger
	table["wipc"] = PiperSchemaHandler{}
	return table
}
