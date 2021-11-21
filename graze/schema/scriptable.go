package schema

//a placeholder handler that would call through to a script to impl a protocol

type ScriptableSchemaHandler struct {
	script       string
	targetSchema string
}

func (s ScriptableSchemaHandler) Name() string {
	return "[Script] " + s.targetSchema + " handler."
}

//TODO: Implement s.Query()
